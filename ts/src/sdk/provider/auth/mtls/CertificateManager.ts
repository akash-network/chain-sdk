import { generateCertificate, parseCertificate } from "./x509.ts";

/**
 * Represents the PEM encoded certificate, public key, and private key.
 */
export interface CertificatePem {
  cert: string;
  publicKey: string;
  privateKey: string;
}

/**
 * Represents the information extracted from a certificate.
 */
export interface CertificateInfo {
  hSerial: string;
  sIssuer: string;
  sSubject: string;
  sNotBefore: string;
  sNotAfter: string;
  issuedOn: Date;
  expiresOn: Date;
}

/**
 * Options for specifying the validity range of a certificate.
 */
export interface ValidityRangeOptions {
  serial?: number;
  validFrom?: Date;
  validTo?: Date;
}

/**
 * Manages the creation and parsing of certificates.
 */
export class CertificateManager {
  /**
   * Parses a PEM encoded certificate and extracts its information.
   * @param certPEM - The PEM encoded certificate string.
   * @returns An object containing the certificate information.
   * @example
   * const certificateManager = new CertificateManager();
   * const pem = await certificateManager.generatePEM('exampleAddress');
   * const certInfo = await certificateManager.parsePem(pem.cert);
   * console.log(certInfo);
   */
  async parsePem(certPEM: string): Promise<CertificateInfo> {
    return parseCertificate(certPEM);
  }

  /**
   * Generates a PEM encoded certificate, public key, and private key.
   * @param address - The address to be used as the certificate's subject and issuer.
   * @param options - Optional validity range for the certificate.
   * @returns An object containing the PEM encoded certificate, public key, and private key.
   * @example
   * const certificateManager = new CertificateManager();
   * const pem = await certificateManager.generatePEM('exampleAddress');
   * console.log('Certificate:', pem.cert);
   * console.log('Public Key:', pem.publicKey);
   * console.log('Private Key:', pem.privateKey);
   */
  async generatePEM(address: string, options?: ValidityRangeOptions): Promise<CertificatePem> {
    return generateCertificate(address, options);
  }
}
