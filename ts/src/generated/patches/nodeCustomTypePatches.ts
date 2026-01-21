import { LegacyDec } from "../../encoding/customTypes/LegacyDec.ts";
import type * as _protos_akash_bme_v1_types from "../protos/akash/bme/v1/types.ts";
import type * as _protos_akash_bme_v1_events from "../protos/akash/bme/v1/events.ts";
import type * as _protos_akash_bme_v1_genesis from "../protos/akash/bme/v1/genesis.ts";
import type * as _protos_akash_bme_v1_query from "../protos/akash/bme/v1/query.ts";
import type * as _protos_akash_deployment_v1beta4_resourceunit from "../protos/akash/deployment/v1beta4/resourceunit.ts";
import type * as _protos_cosmos_base_v1beta1_coin from "../protos/cosmos/base/v1beta1/coin.ts";
import type * as _protos_akash_escrow_types_v1_balance from "../protos/akash/escrow/types/v1/balance.ts";

const p = {
  "cosmos.base.v1beta1.DecCoin"(value: _protos_cosmos_base_v1beta1_coin.DecCoin | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.amount != null) newValue.amount = LegacyDec[transformType](value.amount);
    return newValue;
  },
  "akash.escrow.types.v1.Balance"(value: _protos_akash_escrow_types_v1_balance.Balance | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.amount != null) newValue.amount = LegacyDec[transformType](value.amount);
    return newValue;
  }
};

export const patches = p;
