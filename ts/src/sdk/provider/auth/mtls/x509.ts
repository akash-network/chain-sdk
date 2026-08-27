/**
 * Self-contained X.509 encode/decode (self-signed EC P-256 client cert).
 */

import { fromByteArray, toByteArray } from "base64-js";

import type { CertificateInfo, CertificatePem, ValidityRangeOptions } from "./CertificateManager.ts";

const KEY_ALGORITHM: EcKeyGenParams = { name: "ECDSA", namedCurve: "P-256" };
const SIGN_ALGORITHM: EcdsaParams = { name: "ECDSA", hash: "SHA-256" };
const OID = {
  ecdsaWithSHA256: "1.2.840.10045.4.3.2",
  commonName: "2.5.4.3",
  authVersion: "2.23.133.2.6",
  clientAuth: "1.3.6.1.5.5.7.3.2",
  keyUsage: "2.5.29.15",
  extKeyUsage: "2.5.29.37",
  basicConstraints: "2.5.29.19",
} as const;

// Auth protocol version carried in the certificate DN, matching the Akash Go
// reference (go/node/cert/v1/utils/key_pair_manager.go).
const AUTH_VERSION = "v0.0.1";

export async function generateCertificate(address: string, options?: ValidityRangeOptions): Promise<CertificatePem> {
  const notBefore = options?.validFrom ? new Date(options.validFrom) : new Date();
  const notAfter = options?.validTo ? new Date(options.validTo) : new Date(notBefore);
  if (!options?.validTo) notAfter.setFullYear(notBefore.getFullYear() + 1);
  if (notAfter <= notBefore) throw new Error("Certificate validTo must be after validFrom");

  const keys = await globalThis.crypto.subtle.generateKey(KEY_ALGORITHM, true, ["sign", "verify"]);
  const spki = new Uint8Array(await globalThis.crypto.subtle.exportKey("spki", keys.publicKey));
  const pkcs8 = new Uint8Array(await globalThis.crypto.subtle.exportKey("pkcs8", keys.privateKey));

  const sigAlg = seq(oid(OID.ecdsaWithSHA256));
  const name = buildName(address);
  const tbs = seq(
    explicit(0, integer([2])), // version v3
    integerFromBigInt(BigInt(options?.serial ?? Math.floor(Date.now() * 1000))),
    sigAlg,
    name, // issuer
    seq(time(notBefore), time(notAfter)),
    name, // subject (self-signed)
    [...spki], // SubjectPublicKeyInfo, straight from Web Crypto
    explicit(
      3,
      seq(
        extension(OID.keyUsage, true, bitString(4, [0x30])), // keyEncipherment + dataEncipherment
        extension(OID.extKeyUsage, false, seq(oid(OID.clientAuth))),
        extension(OID.basicConstraints, true, seq()), // empty => cA FALSE (leaf client cert)
      ),
    ),
  );

  const raw = new Uint8Array(await globalThis.crypto.subtle.sign(SIGN_ALGORITHM, keys.privateKey, Uint8Array.from(tbs)));
  const signature = seq(integer([...raw.slice(0, 32)]), integer([...raw.slice(32)])); // raw r||s -> DER SEQUENCE
  const certDer = seq(tbs, sigAlg, bitString(0, signature));

  return {
    cert: derToPem(certDer, "CERTIFICATE"),
    publicKey: derToPem(spki, "PUBLIC KEY").replaceAll("PUBLIC KEY", "EC PUBLIC KEY"),
    privateKey: derToPem(pkcs8, "PRIVATE KEY"),
  };
}

export function parseCertificate(certPEM: string): CertificateInfo {
  const tbs = children(read(pemToDer(certPEM), 0).content)[0]; // Certificate -> tbsCertificate
  const fields = children(tbs.content);
  let i = (fields[0].tag & 0xc0) === 0x80 ? 1 : 0; // skip version [0] if present
  const serial = fields[i++];
  i++; // signature algorithm
  const issuer = fields[i++];
  const validity = fields[i++];
  const subject = fields[i++];
  const [notBefore, notAfter] = children(validity.content);

  const issuedOn = parseTime(notBefore);
  const expiresOn = parseTime(notAfter);
  return {
    hSerial: toHex(serial.content),
    sIssuer: `/CN=${commonName(issuer)}`,
    sSubject: `/CN=${commonName(subject)}`,
    sNotBefore: dateToStr(issuedOn),
    sNotAfter: dateToStr(expiresOn),
    issuedOn,
    expiresOn,
  };
}

/**
 * Converts a Date object to a string in the format YYMMDDHHMMSSZ.
 * @param date - The date to convert.
 * @returns The formatted date string.
 * @example
 * const dateStr = dateToStr(new Date('2024-05-07T12:23:50.000Z'));
 * console.log(dateStr); // "240507122350Z"
 */
export function dateToStr(date: Date): string {
  const year = date.getUTCFullYear().toString().slice(2).padStart(2, "0");
  const month = (date.getUTCMonth() + 1).toString().padStart(2, "0");
  const day = date.getUTCDate().toString().padStart(2, "0");
  const hours = date.getUTCHours().toString().padStart(2, "0");
  const minutes = date.getUTCMinutes().toString().padStart(2, "0");
  const secs = date.getUTCSeconds().toString().padStart(2, "0");

  return `${year}${month}${day}${hours}${minutes}${secs}Z`;
}

/**
 * Converts a string in the format YYMMDDHHMMSSZ to a Date object.
 * @param str - The string to convert.
 * @returns The corresponding Date object.
 * @example
 * const date = strToDate("240507122350Z");
 * console.log(date.toISOString()); // "2024-05-07T12:23:50.000Z"
 */
export function strToDate(str: string): Date {
  const year = parseInt(`20${str.substring(0, 2)}`);
  const month = parseInt(str.substring(2, 4)) - 1;
  const day = parseInt(str.substring(4, 6));
  const hours = parseInt(str.substring(6, 8));
  const minutes = parseInt(str.substring(8, 10));
  const secs = parseInt(str.substring(10, 12));

  return new Date(Date.UTC(year, month, day, hours, minutes, secs));
}

// ---------- DER encoding ----------

const TAG = { BOOLEAN: 0x01, INTEGER: 0x02, BIT_STRING: 0x03, OCTET_STRING: 0x04, OID: 0x06, PRINTABLE: 0x13, SEQUENCE: 0x30, SET: 0x31, UTC_TIME: 0x17, GEN_TIME: 0x18 } as const;

function encodeLength(len: number): number[] {
  if (len < 0x80) return [len];
  const bytes: number[] = [];
  for (let n = len; n > 0; n = Math.floor(n / 256)) bytes.unshift(n & 0xff);
  return [0x80 | bytes.length, ...bytes];
}

function tlv(tag: number, content: number[]): number[] {
  return [tag, ...encodeLength(content.length), ...content];
}

const seq = (...items: number[][]): number[] => tlv(TAG.SEQUENCE, items.flat());
const set = (...items: number[][]): number[] => tlv(TAG.SET, items.flat());
const explicit = (tagNumber: number, content: number[]): number[] => tlv(0xa0 | tagNumber, content);

/** DER INTEGER from a big-endian magnitude: strip leading zeros, pad so it stays positive. */
function integer(magnitude: number[]): number[] {
  let bytes = magnitude;
  while (bytes.length > 1 && bytes[0] === 0) bytes = bytes.slice(1);
  if (bytes[0] & 0x80) bytes = [0x00, ...bytes];
  return tlv(TAG.INTEGER, bytes);
}

function integerFromBigInt(value: bigint): number[] {
  const bytes: number[] = [];
  for (let n = value; n > 0n; n >>= 8n) bytes.unshift(Number(n & 0xffn));
  return integer(bytes.length ? bytes : [0]);
}

function oid(dotted: string): number[] {
  const arcs = dotted.split(".").map(Number);
  const body = [40 * arcs[0] + arcs[1]];
  for (const arc of arcs.slice(2)) {
    const group = [arc & 0x7f];
    for (let n = Math.floor(arc / 128); n > 0; n = Math.floor(n / 128)) group.unshift((n & 0x7f) | 0x80);
    body.push(...group);
  }
  return tlv(TAG.OID, body);
}

const boolean = (value: boolean): number[] => tlv(TAG.BOOLEAN, [value ? 0xff : 0x00]);
const bitString = (unusedBits: number, data: number[]): number[] => tlv(TAG.BIT_STRING, [unusedBits, ...data]);
const octetString = (data: number[]): number[] => tlv(TAG.OCTET_STRING, data);
const printableString = (text: string): number[] => tlv(TAG.PRINTABLE, ascii(text));
const ascii = (text: string): number[] => Array.from(text, (c) => c.charCodeAt(0));

function time(date: Date): number[] {
  const d = new Date(date);
  d.setUTCMilliseconds(0);
  const pad = (n: number, len = 2) => n.toString().padStart(len, "0");
  const rest = `${pad(d.getUTCMonth() + 1)}${pad(d.getUTCDate())}${pad(d.getUTCHours())}${pad(d.getUTCMinutes())}${pad(d.getUTCSeconds())}Z`;
  return d.getUTCFullYear() >= 2050
    ? tlv(TAG.GEN_TIME, ascii(`${pad(d.getUTCFullYear(), 4)}${rest}`))
    : tlv(TAG.UTC_TIME, ascii(`${pad(d.getUTCFullYear() % 100)}${rest}`));
}

// Name ::= SEQUENCE OF RDN. Two RDNs, matching the Go reference: CN then authVersion.
const buildName = (cn: string): number[] =>
  seq(set(seq(oid(OID.commonName), printableString(cn))), set(seq(oid(OID.authVersion), printableString(AUTH_VERSION))));
const extension = (id: string, critical: boolean, value: number[]): number[] => seq(oid(id), ...(critical ? [boolean(true)] : []), octetString(value));

// ---------- DER decoding ----------

interface Node {
  tag: number;
  content: Uint8Array;
  end: number;
}

function read(buf: Uint8Array, offset: number): Node {
  if (offset + 2 > buf.length) throw new Error("Invalid DER: truncated header");
  let pos = offset + 1;
  let len = buf[pos++];
  if (len & 0x80) {
    let n = len & 0x7f;
    if (n === 0 || pos + n > buf.length) throw new Error("Invalid DER: invalid length");
    for (len = 0; n > 0; n--) len = len * 256 + buf[pos++];
  }
  if (pos + len > buf.length) throw new Error("Invalid DER: truncated content");
  return { tag: buf[offset], content: buf.subarray(pos, pos + len), end: pos + len };
}

function children(content: Uint8Array): Node[] {
  const nodes: Node[] = [];
  for (let off = 0; off < content.length;) {
    const node = read(content, off);
    nodes.push(node);
    off = node.end;
  }
  return nodes;
}

const commonNameOidContent = Uint8Array.from([0x55, 0x04, 0x03]); // DER content of 2.5.4.3

function commonName(name: Node): string {
  for (const rdn of children(name.content)) {
    for (const attribute of children(rdn.content)) {
      const [type, value] = children(attribute.content);
      if (type.content.length === 3 && type.content.every((b, i) => b === commonNameOidContent[i])) {
        return new TextDecoder().decode(value.content);
      }
    }
  }
  return "";
}

function parseTime(node: Node): Date {
  const s = String.fromCharCode(...node.content);
  const generalized = node.tag === TAG.GEN_TIME;
  const year = generalized ? +s.slice(0, 4) : +s.slice(0, 2) >= 50 ? 1900 + +s.slice(0, 2) : 2000 + +s.slice(0, 2);
  const o = generalized ? 4 : 2;
  return new Date(Date.UTC(year, +s.slice(o, o + 2) - 1, +s.slice(o + 2, o + 4), +s.slice(o + 4, o + 6), +s.slice(o + 6, o + 8), +s.slice(o + 8, o + 10)));
}

// ---------- PEM ----------

function derToPem(der: number[] | Uint8Array, label: string): string {
  const base64 = fromByteArray(der instanceof Uint8Array ? der : Uint8Array.from(der));
  const lines = base64.match(/.{1,64}/g) ?? [];
  return `-----BEGIN ${label}-----\r\n${lines.join("\r\n")}\r\n-----END ${label}-----\r\n`;
}

const pemToDer = (pem: string): Uint8Array => toByteArray(pem.replace(/-----[^-]+-----/g, "").replace(/\s+/g, ""));
const toHex = (bytes: Uint8Array): string => Array.from(bytes, (b) => b.toString(16).padStart(2, "0")).join("");
