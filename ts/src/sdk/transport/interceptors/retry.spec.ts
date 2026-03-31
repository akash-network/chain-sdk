import { describe, expect, it, jest } from "@jest/globals";

import { TransportError } from "../TransportError.ts";
import { createRetryInterceptor, isRetryEnabled, type RetryOptions } from "./retry.ts";

describe("createRetryInterceptor", () => {
  const defaultOptions: RetryOptions = { maxAttempts: 3, maxDelayMs: 0 };

  it("returns response on successful first call", async () => {
    const response = createMockResponse();
    const next = createMockNext([() => response]);
    const interceptor = createRetryInterceptor(defaultOptions);
    const handler = interceptor(next);

    const result = await handler(mockRequest);

    expect(result).toBe(response);
    expect(next).toHaveBeenCalledTimes(1);
  });

  describe("retries on retriable TransportError codes", () => {
    it.each([
      TransportError.Code.Unavailable,
      TransportError.Code.DeadlineExceeded,
      TransportError.Code.Internal,
      TransportError.Code.Unknown,
    ])("retries on TransportError with code %s", async (code) => {
      const response = createMockResponse();
      const next = createMockNext([
        () => { throw new TransportError("fail", code); },
        () => response,
      ]);
      const interceptor = createRetryInterceptor(defaultOptions);
      const handler = interceptor(next);

      const result = await handler(mockRequest);

      expect(result).toBe(response);
      expect(next).toHaveBeenCalledTimes(2);
    });
  });

  describe("does not retry on non-retriable TransportError codes", () => {
    it.each([
      TransportError.Code.InvalidArgument,
      TransportError.Code.NotFound,
      TransportError.Code.AlreadyExists,
      TransportError.Code.PermissionDenied,
      TransportError.Code.Unauthenticated,
      TransportError.Code.Canceled,
    ])("does not retry on TransportError with code %s", async (code) => {
      const next = createMockNext([
        () => { throw new TransportError("fail", code); },
      ]);
      const interceptor = createRetryInterceptor(defaultOptions);
      const handler = interceptor(next);

      await expect(handler(mockRequest)).rejects.toThrow(TransportError);
      expect(next).toHaveBeenCalledTimes(1);
    });
  });

  it("does not retry on generic errors", async () => {
    const next = createMockNext([
      () => { throw new Error("generic"); },
    ]);
    const interceptor = createRetryInterceptor(defaultOptions);
    const handler = interceptor(next);

    await expect(handler(mockRequest)).rejects.toThrow("generic");
    expect(next).toHaveBeenCalledTimes(1);
  });

  it("caps maxAttempts at 3", async () => {
    const next = createMockNext([
      () => { throw new TransportError("fail", TransportError.Code.Unavailable); },
      () => { throw new TransportError("fail", TransportError.Code.Unavailable); },
      () => { throw new TransportError("fail", TransportError.Code.Unavailable); },
      () => createMockResponse(),
    ]);
    const interceptor = createRetryInterceptor({ maxAttempts: 10, maxDelayMs: 0 });
    const handler = interceptor(next);

    const result = await handler(mockRequest);

    expect(result).toBeDefined();
    // 1 initial + 3 retries = 4 total calls max
    expect(next).toHaveBeenCalledTimes(4);
  });

  it("throws after exhausting all retry attempts", async () => {
    const next = createMockNext([
      () => { throw new TransportError("fail", TransportError.Code.Unavailable); },
      () => { throw new TransportError("fail", TransportError.Code.Unavailable); },
      () => { throw new TransportError("fail", TransportError.Code.Unavailable); },
      () => { throw new TransportError("fail", TransportError.Code.Unavailable); },
    ]);
    const interceptor = createRetryInterceptor(defaultOptions);
    const handler = interceptor(next);

    await expect(handler(mockRequest)).rejects.toThrow(TransportError);
  });

  describe("retries on connection errors", () => {
    it.each([
      "ECONNREFUSED",
      "ECONNRESET",
      "ETIMEDOUT",
      "ESOCKETTIMEDOUT",
      "UND_ERR_SOCKET",
    ])("retries on error with code %s", async (code) => {
      const response = createMockResponse();
      const error = new Error("connection failed");
      (error as Error & { code: string }).code = code;
      const next = createMockNext([
        () => { throw error; },
        () => response,
      ]);
      const interceptor = createRetryInterceptor(defaultOptions);
      const handler = interceptor(next);

      const result = await handler(mockRequest);

      expect(result).toBe(response);
      expect(next).toHaveBeenCalledTimes(2);
    });

    it("retries on nested connection error in cause chain", async () => {
      const response = createMockResponse();
      const innerError = new Error("socket error");
      (innerError as Error & { code: string }).code = "ECONNRESET";
      const outerError = new Error("request failed", { cause: innerError });
      const next = createMockNext([
        () => { throw outerError; },
        () => response,
      ]);
      const interceptor = createRetryInterceptor(defaultOptions);
      const handler = interceptor(next);

      const result = await handler(mockRequest);

      expect(result).toBe(response);
      expect(next).toHaveBeenCalledTimes(2);
    });

    it("retries on connection error in aggregate errors list", async () => {
      const response = createMockResponse();
      const socketError = new Error("socket error");
      (socketError as Error & { code: string }).code = "ECONNREFUSED";
      const aggregateError = new Error("multiple failures");
      (aggregateError as Error & { errors: Error[] }).errors = [socketError];
      const next = createMockNext([
        () => { throw aggregateError; },
        () => response,
      ]);
      const interceptor = createRetryInterceptor(defaultOptions);
      const handler = interceptor(next);

      const result = await handler(mockRequest);

      expect(result).toBe(response);
      expect(next).toHaveBeenCalledTimes(2);
    });
  });

  it("uses lower of maxAttempts and 3", async () => {
    const next = createMockNext([
      () => { throw new TransportError("fail", TransportError.Code.Unavailable); },
      () => { throw new TransportError("fail", TransportError.Code.Unavailable); },
    ]);
    const interceptor = createRetryInterceptor({ maxAttempts: 1, maxDelayMs: 0 });
    const handler = interceptor(next);

    await expect(handler(mockRequest)).rejects.toThrow(TransportError);
    // 1 initial + 1 retry = 2 calls
    expect(next).toHaveBeenCalledTimes(2);
  });
});

describe("isRetryEnabled", () => {
  it("returns true for valid options with positive maxAttempts", () => {
    expect(isRetryEnabled({ maxAttempts: 1 })).toBe(true);
    expect(isRetryEnabled({ maxAttempts: 3 })).toBe(true);
  });

  it("returns false for undefined", () => {
    expect(isRetryEnabled(undefined)).toBe(false);
  });

  it("returns false for maxAttempts of 0", () => {
    expect(isRetryEnabled({ maxAttempts: 0 })).toBe(false);
  });

  it("returns false for negative maxAttempts", () => {
    expect(isRetryEnabled({ maxAttempts: -1 })).toBe(false);
  });

  it("returns false for NaN maxAttempts", () => {
    expect(isRetryEnabled({ maxAttempts: NaN })).toBe(false);
  });
});

const mockRequest = {} as Parameters<Parameters<ReturnType<typeof createRetryInterceptor>>[0]>[0];
type NextFn = Parameters<ReturnType<typeof createRetryInterceptor>>[0];
type ResponseType = Awaited<ReturnType<NextFn>>;

function createMockNext(results: (() => ResponseType)[]) {
  let callIndex = 0;
  const fn = jest.fn<NextFn>().mockImplementation(async () => {
    const result = results[callIndex++];
    if (!result) throw new Error("unexpected call");
    return result();
  });
  return fn;
}

function createMockResponse() {
  return {} as ResponseType;
}
