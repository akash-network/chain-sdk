import type * as _protos_akash_deployment_v1beta4_resourceunit from "../protos/akash/deployment/v1beta4/resourceunit.ts";
import { LegacyDec } from "../../encoding/customTypes/LegacyDec.ts";
import type * as _protos_cosmos_base_v1beta1_coin from "../protos/cosmos/base/v1beta1/coin.ts";
import type * as _protos_akash_deployment_v1beta4_groupspec from "../protos/akash/deployment/v1beta4/groupspec.ts";
import type * as _protos_akash_deployment_v1beta4_deploymentmsg from "../protos/akash/deployment/v1beta4/deploymentmsg.ts";
import type * as _protos_akash_deployment_v1beta4_group from "../protos/akash/deployment/v1beta4/group.ts";
import type * as _protos_akash_deployment_v1beta4_genesis from "../protos/akash/deployment/v1beta4/genesis.ts";
import type * as _protos_akash_escrow_types_v1_balance from "../protos/akash/escrow/types/v1/balance.ts";
import type * as _protos_akash_escrow_types_v1_deposit from "../protos/akash/escrow/types/v1/deposit.ts";
import type * as _protos_akash_escrow_types_v1_account from "../protos/akash/escrow/types/v1/account.ts";
import type * as _protos_akash_deployment_v1beta4_query from "../protos/akash/deployment/v1beta4/query.ts";
import type * as _protos_akash_escrow_types_v1_payment from "../protos/akash/escrow/types/v1/payment.ts";
import type * as _protos_akash_escrow_v1_genesis from "../protos/akash/escrow/v1/genesis.ts";
import type * as _protos_akash_escrow_v1_query from "../protos/akash/escrow/v1/query.ts";
import type * as _protos_akash_market_v1_lease from "../protos/akash/market/v1/lease.ts";
import type * as _protos_akash_market_v1_event from "../protos/akash/market/v1/event.ts";
import type * as _protos_akash_market_v1beta5_bid from "../protos/akash/market/v1beta5/bid.ts";
import type * as _protos_akash_market_v1beta5_bidmsg from "../protos/akash/market/v1beta5/bidmsg.ts";
import type * as _protos_akash_market_v1beta5_order from "../protos/akash/market/v1beta5/order.ts";
import type * as _protos_akash_market_v1beta5_genesis from "../protos/akash/market/v1beta5/genesis.ts";
import type * as _protos_akash_market_v1beta5_query from "../protos/akash/market/v1beta5/query.ts";

const p = {
  "akash.deployment.v1beta4.ResourceUnit"(value: _protos_akash_deployment_v1beta4_resourceunit.ResourceUnit | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.price != null) newValue.price = p["cosmos.base.v1beta1.DecCoin"](value.price, transformType);
    return newValue;
  },
  "cosmos.base.v1beta1.DecCoin"(value: _protos_cosmos_base_v1beta1_coin.DecCoin | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.amount != null) newValue.amount = LegacyDec[transformType](value.amount);
    return newValue;
  },
  "akash.deployment.v1beta4.GroupSpec"(value: _protos_akash_deployment_v1beta4_groupspec.GroupSpec | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.resources) newValue.resources = value.resources.map((item) => p["akash.deployment.v1beta4.ResourceUnit"](item, transformType)!);
    return newValue;
  },
  "akash.deployment.v1beta4.MsgCreateDeployment"(value: _protos_akash_deployment_v1beta4_deploymentmsg.MsgCreateDeployment | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.groups) newValue.groups = value.groups.map((item) => p["akash.deployment.v1beta4.GroupSpec"](item, transformType)!);
    return newValue;
  },
  "akash.deployment.v1beta4.Group"(value: _protos_akash_deployment_v1beta4_group.Group | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.groupSpec != null) newValue.groupSpec = p["akash.deployment.v1beta4.GroupSpec"](value.groupSpec, transformType);
    return newValue;
  },
  "akash.deployment.v1beta4.GenesisDeployment"(value: _protos_akash_deployment_v1beta4_genesis.GenesisDeployment | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.groups) newValue.groups = value.groups.map((item) => p["akash.deployment.v1beta4.Group"](item, transformType)!);
    return newValue;
  },
  "akash.deployment.v1beta4.GenesisState"(value: _protos_akash_deployment_v1beta4_genesis.GenesisState | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.deployments) newValue.deployments = value.deployments.map((item) => p["akash.deployment.v1beta4.GenesisDeployment"](item, transformType)!);
    return newValue;
  },
  "akash.escrow.types.v1.Balance"(value: _protos_akash_escrow_types_v1_balance.Balance | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.amount != null) newValue.amount = LegacyDec[transformType](value.amount);
    return newValue;
  },
  "akash.escrow.types.v1.Depositor"(value: _protos_akash_escrow_types_v1_deposit.Depositor | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.balance != null) newValue.balance = p["cosmos.base.v1beta1.DecCoin"](value.balance, transformType);
    return newValue;
  },
  "akash.escrow.types.v1.AccountState"(value: _protos_akash_escrow_types_v1_account.AccountState | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.transferred) newValue.transferred = value.transferred.map((item) => p["cosmos.base.v1beta1.DecCoin"](item, transformType)!);
    if (value.funds) newValue.funds = value.funds.map((item) => p["akash.escrow.types.v1.Balance"](item, transformType)!);
    if (value.deposits) newValue.deposits = value.deposits.map((item) => p["akash.escrow.types.v1.Depositor"](item, transformType)!);
    return newValue;
  },
  "akash.escrow.types.v1.Account"(value: _protos_akash_escrow_types_v1_account.Account | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.state != null) newValue.state = p["akash.escrow.types.v1.AccountState"](value.state, transformType);
    return newValue;
  },
  "akash.deployment.v1beta4.QueryDeploymentsResponse"(value: _protos_akash_deployment_v1beta4_query.QueryDeploymentsResponse | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.deployments) newValue.deployments = value.deployments.map((item) => p["akash.deployment.v1beta4.QueryDeploymentResponse"](item, transformType)!);
    return newValue;
  },
  "akash.deployment.v1beta4.QueryDeploymentResponse"(value: _protos_akash_deployment_v1beta4_query.QueryDeploymentResponse | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.groups) newValue.groups = value.groups.map((item) => p["akash.deployment.v1beta4.Group"](item, transformType)!);
    if (value.escrowAccount != null) newValue.escrowAccount = p["akash.escrow.types.v1.Account"](value.escrowAccount, transformType);
    return newValue;
  },
  "akash.deployment.v1beta4.QueryGroupResponse"(value: _protos_akash_deployment_v1beta4_query.QueryGroupResponse | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.group != null) newValue.group = p["akash.deployment.v1beta4.Group"](value.group, transformType);
    return newValue;
  },
  "akash.escrow.types.v1.PaymentState"(value: _protos_akash_escrow_types_v1_payment.PaymentState | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.rate != null) newValue.rate = p["cosmos.base.v1beta1.DecCoin"](value.rate, transformType);
    if (value.balance != null) newValue.balance = p["cosmos.base.v1beta1.DecCoin"](value.balance, transformType);
    if (value.unsettled != null) newValue.unsettled = p["cosmos.base.v1beta1.DecCoin"](value.unsettled, transformType);
    return newValue;
  },
  "akash.escrow.types.v1.Payment"(value: _protos_akash_escrow_types_v1_payment.Payment | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.state != null) newValue.state = p["akash.escrow.types.v1.PaymentState"](value.state, transformType);
    return newValue;
  },
  "akash.escrow.v1.GenesisState"(value: _protos_akash_escrow_v1_genesis.GenesisState | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.accounts) newValue.accounts = value.accounts.map((item) => p["akash.escrow.types.v1.Account"](item, transformType)!);
    if (value.payments) newValue.payments = value.payments.map((item) => p["akash.escrow.types.v1.Payment"](item, transformType)!);
    return newValue;
  },
  "akash.escrow.v1.QueryAccountsResponse"(value: _protos_akash_escrow_v1_query.QueryAccountsResponse | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.accounts) newValue.accounts = value.accounts.map((item) => p["akash.escrow.types.v1.Account"](item, transformType)!);
    return newValue;
  },
  "akash.escrow.v1.QueryPaymentsResponse"(value: _protos_akash_escrow_v1_query.QueryPaymentsResponse | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.payments) newValue.payments = value.payments.map((item) => p["akash.escrow.types.v1.Payment"](item, transformType)!);
    return newValue;
  },
  "akash.market.v1.Lease"(value: _protos_akash_market_v1_lease.Lease | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.price != null) newValue.price = p["cosmos.base.v1beta1.DecCoin"](value.price, transformType);
    return newValue;
  },
  "akash.market.v1.EventBidCreated"(value: _protos_akash_market_v1_event.EventBidCreated | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.price != null) newValue.price = p["cosmos.base.v1beta1.DecCoin"](value.price, transformType);
    return newValue;
  },
  "akash.market.v1.EventLeaseCreated"(value: _protos_akash_market_v1_event.EventLeaseCreated | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.price != null) newValue.price = p["cosmos.base.v1beta1.DecCoin"](value.price, transformType);
    return newValue;
  },
  "akash.market.v1beta5.Bid"(value: _protos_akash_market_v1beta5_bid.Bid | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.price != null) newValue.price = p["cosmos.base.v1beta1.DecCoin"](value.price, transformType);
    return newValue;
  },
  "akash.market.v1beta5.MsgCreateBid"(value: _protos_akash_market_v1beta5_bidmsg.MsgCreateBid | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.price != null) newValue.price = p["cosmos.base.v1beta1.DecCoin"](value.price, transformType);
    return newValue;
  },
  "akash.market.v1beta5.Order"(value: _protos_akash_market_v1beta5_order.Order | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.spec != null) newValue.spec = p["akash.deployment.v1beta4.GroupSpec"](value.spec, transformType);
    return newValue;
  },
  "akash.market.v1beta5.GenesisState"(value: _protos_akash_market_v1beta5_genesis.GenesisState | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.orders) newValue.orders = value.orders.map((item) => p["akash.market.v1beta5.Order"](item, transformType)!);
    if (value.leases) newValue.leases = value.leases.map((item) => p["akash.market.v1.Lease"](item, transformType)!);
    if (value.bids) newValue.bids = value.bids.map((item) => p["akash.market.v1beta5.Bid"](item, transformType)!);
    return newValue;
  },
  "akash.market.v1beta5.QueryOrdersResponse"(value: _protos_akash_market_v1beta5_query.QueryOrdersResponse | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.orders) newValue.orders = value.orders.map((item) => p["akash.market.v1beta5.Order"](item, transformType)!);
    return newValue;
  },
  "akash.market.v1beta5.QueryOrderResponse"(value: _protos_akash_market_v1beta5_query.QueryOrderResponse | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.order != null) newValue.order = p["akash.market.v1beta5.Order"](value.order, transformType);
    return newValue;
  },
  "akash.market.v1beta5.QueryBidsResponse"(value: _protos_akash_market_v1beta5_query.QueryBidsResponse | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.bids) newValue.bids = value.bids.map((item) => p["akash.market.v1beta5.QueryBidResponse"](item, transformType)!);
    return newValue;
  },
  "akash.market.v1beta5.QueryBidResponse"(value: _protos_akash_market_v1beta5_query.QueryBidResponse | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.bid != null) newValue.bid = p["akash.market.v1beta5.Bid"](value.bid, transformType);
    if (value.escrowAccount != null) newValue.escrowAccount = p["akash.escrow.types.v1.Account"](value.escrowAccount, transformType);
    return newValue;
  },
  "akash.market.v1beta5.QueryLeasesResponse"(value: _protos_akash_market_v1beta5_query.QueryLeasesResponse | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.leases) newValue.leases = value.leases.map((item) => p["akash.market.v1beta5.QueryLeaseResponse"](item, transformType)!);
    return newValue;
  },
  "akash.market.v1beta5.QueryLeaseResponse"(value: _protos_akash_market_v1beta5_query.QueryLeaseResponse | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.lease != null) newValue.lease = p["akash.market.v1.Lease"](value.lease, transformType);
    if (value.escrowPayment != null) newValue.escrowPayment = p["akash.escrow.types.v1.Payment"](value.escrowPayment, transformType);
    return newValue;
  }
};

export const patches = p;
