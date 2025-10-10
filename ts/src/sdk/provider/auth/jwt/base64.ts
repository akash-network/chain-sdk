export function base64UrlEncode(value: string | Uint8Array): string {
  const str = typeof value === "string" ? value : String.fromCharCode(...value);
  const base64 = btoa(str);
  return toBase64Url(base64);
}

/**
 * Converts a base64 encoded string to a base64url encoded string
 */
export function toBase64Url(base64Encoded: string): string {
  return base64Encoded.replace(/\+/g, "-").replace(/\//g, "_").replace(/=+$/, "");
}

export function base64UrlDecode(value: string): string {
  let str = value;
  // Convert from base64url → base64
  str = str.replace(/-/g, "+").replace(/_/g, "/");
  str = str.padEnd(str.length + (4 - (str.length % 4)) % 4, "=");

  return new TextDecoder().decode(Uint8Array.from(atob(str), (c) => c.charCodeAt(0)));
}

/**
 * Decode a base64 string
 * @param base64String The base64 string to decode
 * @returns The decoded object
 */
export function base64Decode(base64String: string): Record<string, unknown> {
  const decoded = atob(base64String);
  return JSON.parse(decoded);
}
