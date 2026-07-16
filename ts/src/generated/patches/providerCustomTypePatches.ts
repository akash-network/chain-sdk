import { Int, bigIntFromJSON, bigIntFromPartial } from "../../encoding/customTypes/Int.ts";
import { encodeBinary, decodeBinary } from "../../encoding/binaryEncoding.ts";
import type * as _protos_akash_base_resources_v1beta4_resourcevalue from "../protos/akash/base/resources/v1beta4/resourcevalue.ts";
import { LegacyDec } from "../../encoding/customTypes/LegacyDec.ts";
import type * as _protos_cosmos_base_v1beta1_coin from "../protos/cosmos/base/v1beta1/coin.ts";

const p = {
  "akash.base.resources.v1beta4.ResourceValue": Object.assign(function(value: any, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.val != null) newValue.val = transformType === 'encode' ? encodeBinary(Int.encode(value.val)) : Int.decode(decodeBinary(value.val));
    return newValue;
  }, {
    fromJSON(object: any) {
      const newValue: any = {};
      newValue.val = bigIntFromJSON(object.val);
      return newValue;
    },
    toJSON(value: any) {
      const newValue: any = {};
      if (value.val != null) newValue.val = value.val.toString();
      return newValue;
    },
    fromPartial(newValue: any) {
      if (newValue.val != null) newValue.val = bigIntFromPartial(newValue.val);
      return newValue;
    },
  }),
  "cosmos.base.v1beta1.DecCoin"(value: _protos_cosmos_base_v1beta1_coin.DecCoin | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.amount != null) newValue.amount = LegacyDec[transformType](value.amount);
    return newValue;
  }
};

export const patches = p;

export const typeOverrides = {
  "akash.base.resources.v1beta4.ResourceValue": {
    "val": "bigint"
  }
} as const;
