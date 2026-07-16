import { patches } from "./nodeCustomTypePatches.ts";
import type { MessageDesc } from "../../sdk/client/types.ts";
export const patched = <T extends MessageDesc>(messageDesc: T): T => {
  const patchMessage = patches[messageDesc.$type as keyof typeof patches] as any;
  if (!patchMessage) return messageDesc;
  const wrapped: T = {
    ...messageDesc,
    encode(message, writer) {
      return messageDesc.encode(patchMessage(message, 'encode'), writer);
    },
    decode(input, length) {
      return patchMessage(messageDesc.decode(input, length), 'decode');
    },
  };
  // fromJSON / toJSON / fromPartial are overridden only for representation-
  // changing messages, whose patch entry carries these normalizers. Plain
  // function patch entries (Dec/DecCoin) leave the generated behavior intact.
  // (assigned through an `any` view since `wrapped` has the generic type `T`.)
  const overrides = wrapped as any;
  if (patchMessage.fromJSON) {
    overrides.fromJSON = (object: any) => {
      // patchMessage.fromJSON normalizes any accepted input (base64, decimal
      // string, number, bigint) to the JS value. The generated fromJSON only
      // groks the wire-JSON (base64) form, so feed it a sanitized clone (custom
      // fields re-serialized) to fill non-custom fields, then overlay the JS values.
      const patchedFields = patchMessage.fromJSON(object);
      return { ...(messageDesc.fromJSON({ ...object, ...patchMessage.toJSON(patchedFields) }) as object), ...patchedFields };
    };
  }
  if (patchMessage.toJSON) {
    overrides.toJSON = (message: any) => ({ ...(messageDesc.toJSON(patchMessage(message, 'encode')) as object), ...patchMessage.toJSON(message) });
  }
  if (patchMessage.fromPartial) {
    overrides.fromPartial = (object: any) => patchMessage.fromPartial(messageDesc.fromPartial(object));
  }
  return wrapped;
};
