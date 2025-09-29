import { Code, ConnectError } from "@connectrpc/connect";

import type { MessageDesc, MethodDesc } from "../client/types.ts";

export function createSerialization(type: MessageDesc) {
  return {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    parse: (data: Uint8Array): any => type.decode(data),
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    serialize: (data: any) => type.encode(data).finish(),
  };
}

export function coerceTimeoutMs(timeoutMs: number | undefined, defaultTimeoutMs: number | undefined): number | undefined {
  const value = timeoutMs !== undefined && (timeoutMs <= 0 || Number.isNaN(timeoutMs)) ? undefined : timeoutMs;
  return value === undefined ? defaultTimeoutMs : value;
}

export function createMethodUrl(baseUrl: string | URL, method: MethodDesc): string {
  return baseUrl
    .toString()
    .replace(/\/?$/, `/${method.parent.typeName}/${method.name}`);
}

const TransportErrorCode = Code;
export class TransportError extends Error {
  static Code = TransportErrorCode;

  public readonly code: typeof TransportErrorCode[keyof typeof TransportErrorCode];
  public readonly metadata: Headers;
  public readonly cause?: unknown;

  /**
   * Convert any value - typically a caught error into a TransportError,
   * following these rules:
   * - If the value is already a TransportError, return it as is.
   * - If the value is an AbortError from the fetch API, return the message
   *   of the AbortError with code Canceled.
   * - For other Errors, return the error message with code Unknown by default.
   * - For other values, return the values String representation as a message,
   *   with the code Unknown by default.
   * The original value will be used for the "cause" property for the new
   * TransportError.
   */
  static from(cause: unknown, code = TransportError.Code.Unknown) {
    if (cause instanceof this) return cause;
    if (cause instanceof ConnectError) {
      return new TransportError(cause.message, cause.code, { cause, metadata: cause.metadata });
    }
    if (cause instanceof Error) {
      if (cause.name == "AbortError") {
        return new TransportError(cause.message, TransportErrorCode.Canceled);
      }
      return new TransportError(cause.message, code, { cause });
    }
    return new TransportError(String(cause), code, { cause });
  }

  /**
   * Create a new TransportError.
   * Outgoing details are only relevant for the server side - a service may
   * raise an error with details, and it is up to the protocol implementation
   * to encode and send the details along with error.
   */
  constructor(message: string, code = TransportError.Code.Unknown, options?: {
    metadata?: HeadersInit;
    cause?: unknown;
  }) {
    super(`[${stringifyCode(code)}] ${message}`, { cause: options?.cause });
    this.name = "TransportError";
    Object.setPrototypeOf(this, new.target.prototype);
    this.code = code;
    this.metadata = new Headers(options?.metadata ?? {});
    this.cause ??= options?.cause;
  }
}

function stringifyCode(value: typeof TransportErrorCode[keyof typeof TransportErrorCode]) {
  const name = TransportErrorCode[value];
  if (typeof name !== "string") return value.toString();
  return (name[0].toLowerCase() + name.slice(1).replace(/[A-Z]/g, (char) => `_${char.toLowerCase()}`));
}
