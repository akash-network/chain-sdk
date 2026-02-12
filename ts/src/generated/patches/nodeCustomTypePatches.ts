import { LegacyDec } from "../../encoding/customTypes/LegacyDec.ts";
import type * as _protos_akash_bme_v1_types from "../protos/akash/bme/v1/types.ts";
import type * as _protos_akash_bme_v1_events from "../protos/akash/bme/v1/events.ts";
import type * as _protos_akash_bme_v1_query from "../protos/akash/bme/v1/query.ts";
import type * as _protos_cosmos_base_v1beta1_coin from "../protos/cosmos/base/v1beta1/coin.ts";
import type * as _protos_akash_escrow_types_v1_balance from "../protos/akash/escrow/types/v1/balance.ts";
import type * as _protos_akash_oracle_v1_prices from "../protos/akash/oracle/v1/prices.ts";

const p = {
  "akash.bme.v1.CollateralRatio"(value: _protos_akash_bme_v1_types.CollateralRatio | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.ratio != null) newValue.ratio = LegacyDec[transformType](value.ratio);
    if (value.referencePrice != null) newValue.referencePrice = LegacyDec[transformType](value.referencePrice);
    return newValue;
  },
  "akash.bme.v1.CoinPrice"(value: _protos_akash_bme_v1_types.CoinPrice | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.price != null) newValue.price = LegacyDec[transformType](value.price);
    return newValue;
  },
  "akash.bme.v1.EventMintStatusChange"(value: _protos_akash_bme_v1_events.EventMintStatusChange | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.collateralRatio != null) newValue.collateralRatio = LegacyDec[transformType](value.collateralRatio);
    return newValue;
  },
  "akash.bme.v1.QueryStatusResponse"(value: _protos_akash_bme_v1_query.QueryStatusResponse | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.collateralRatio != null) newValue.collateralRatio = LegacyDec[transformType](value.collateralRatio);
    if (value.warnThreshold != null) newValue.warnThreshold = LegacyDec[transformType](value.warnThreshold);
    if (value.haltThreshold != null) newValue.haltThreshold = LegacyDec[transformType](value.haltThreshold);
    return newValue;
  },
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
  },
  "akash.oracle.v1.PriceDataState"(value: _protos_akash_oracle_v1_prices.PriceDataState | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.price != null) newValue.price = LegacyDec[transformType](value.price);
    return newValue;
  },
  "akash.oracle.v1.AggregatedPrice"(value: _protos_akash_oracle_v1_prices.AggregatedPrice | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.twap != null) newValue.twap = LegacyDec[transformType](value.twap);
    if (value.medianPrice != null) newValue.medianPrice = LegacyDec[transformType](value.medianPrice);
    if (value.minPrice != null) newValue.minPrice = LegacyDec[transformType](value.minPrice);
    if (value.maxPrice != null) newValue.maxPrice = LegacyDec[transformType](value.maxPrice);
    return newValue;
  }
};

export const patches = p;
