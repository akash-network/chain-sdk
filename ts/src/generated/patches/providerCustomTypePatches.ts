import { LegacyDec } from "../../encoding/customTypes/LegacyDec.ts";
import type * as _protos_cosmos_base_v1beta1_coin from "../protos/cosmos/base/v1beta1/coin.ts";

const p = {
  "cosmos.base.v1beta1.DecCoin"(value: _protos_cosmos_base_v1beta1_coin.DecCoin | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.amount != null) newValue.amount = LegacyDec[transformType](value.amount);
    return newValue;
  }
};

export const patches = p;
