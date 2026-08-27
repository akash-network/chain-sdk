// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const bytesFromBase64 = (globalThis as any).Buffer
  ? (b64: string): Uint8Array => Uint8Array.from(globalThis.Buffer.from(b64, "base64"))
  : (b64: string): Uint8Array => {
      if ("fromBase64" in Uint8Array && typeof Uint8Array.fromBase64 === "function") {
        return Uint8Array.fromBase64(b64) as Uint8Array;
      }

      const bin = globalThis.atob(b64);
      const arr = new Uint8Array(bin.length);
      for (let i = 0; i < bin.length; ++i) {
        arr[i] = bin.charCodeAt(i);
      }
      return arr;
    };

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const base64FromBytes = (globalThis as any).Buffer
  ? (arr: Uint8Array): string => globalThis.Buffer.from(arr).toString("base64")
  : (arr: Uint8Array): string => {
      if ("toBase64" in arr && typeof arr.toBase64 === "function") {
        return arr.toBase64() as string;
      }

      const bin: string[] = [];
      arr.forEach((byte) => {
        bin.push(globalThis.String.fromCharCode(byte));
      });
      return globalThis.btoa(bin.join(""));
    };
