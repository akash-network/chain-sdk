import { createServiceLoader } from "../sdk/client/createServiceLoader.ts";

import type * as akash_audit_v1_query from "./protos/akash/audit/v1/query.ts";
import type * as akash_audit_v1_msg from "./protos/akash/audit/v1/msg.ts";
import type * as akash_bme_v1_query from "./protos/akash/bme/v1/query.ts";
import type * as akash_bme_v1_msgs from "./protos/akash/bme/v1/msgs.ts";
import type * as akash_cert_v1_query from "./protos/akash/cert/v1/query.ts";
import type * as akash_cert_v1_msg from "./protos/akash/cert/v1/msg.ts";
import type * as akash_deployment_v1beta4_query from "./protos/akash/deployment/v1beta4/query.ts";
import type * as akash_deployment_v1beta4_deploymentmsg from "./protos/akash/deployment/v1beta4/deploymentmsg.ts";
import type * as akash_deployment_v1beta4_groupmsg from "./protos/akash/deployment/v1beta4/groupmsg.ts";
import type * as akash_deployment_v1beta4_paramsmsg from "./protos/akash/deployment/v1beta4/paramsmsg.ts";
import type * as akash_discovery_v1_service from "./protos/akash/discovery/v1/service.ts";
import type * as akash_downtimedetector_v1beta1_query from "./protos/akash/downtimedetector/v1beta1/query.ts";
import type * as akash_epochs_v1beta1_query from "./protos/akash/epochs/v1beta1/query.ts";
import type * as akash_escrow_v1_query from "./protos/akash/escrow/v1/query.ts";
import type * as akash_escrow_v1_msg from "./protos/akash/escrow/v1/msg.ts";
import type * as akash_market_v1beta5_query from "./protos/akash/market/v1beta5/query.ts";
import type * as akash_market_v1beta5_bidmsg from "./protos/akash/market/v1beta5/bidmsg.ts";
import type * as akash_market_v1beta5_leasemsg from "./protos/akash/market/v1beta5/leasemsg.ts";
import type * as akash_market_v1beta5_paramsmsg from "./protos/akash/market/v1beta5/paramsmsg.ts";
import type * as akash_oracle_v1_prices from "./protos/akash/oracle/v1/prices.ts";
import type * as akash_oracle_v1_query from "./protos/akash/oracle/v1/query.ts";
import type * as akash_oracle_v1_msgs from "./protos/akash/oracle/v1/msgs.ts";
import type * as akash_oracle_v2_query from "./protos/akash/oracle/v2/query.ts";
import type * as akash_oracle_v2_msgs from "./protos/akash/oracle/v2/msgs.ts";
import type * as akash_provider_v1beta4_query from "./protos/akash/provider/v1beta4/query.ts";
import type * as akash_provider_v1beta4_msg from "./protos/akash/provider/v1beta4/msg.ts";
import type * as akash_provider_v1beta4_paramsmsg from "./protos/akash/provider/v1beta4/paramsmsg.ts";
import type * as akash_take_v1_query from "./protos/akash/take/v1/query.ts";
import type * as akash_take_v1_paramsmsg from "./protos/akash/take/v1/paramsmsg.ts";
import type * as akash_verification_v1_query from "./protos/akash/verification/v1/query.ts";
import type * as akash_verification_v1_msg from "./protos/akash/verification/v1/msg.ts";
import type * as akash_wasm_v1_query from "./protos/akash/wasm/v1/query.ts";
import type * as akash_wasm_v1_paramsmsg from "./protos/akash/wasm/v1/paramsmsg.ts";
import { createClientFactory } from "../sdk/client/createClientFactory.ts";
import type { Transport, CallOptions, TxCallOptions } from "../sdk/transport/types.ts";
import { withMetadata } from "../sdk/client/sdkMetadata.ts";
import type { DeepPartial, DeepSimplify } from "../encoding/typeEncodingHelpers.ts";


export const serviceLoader= createServiceLoader([
  () => import("./protos/akash/audit/v1/query_akash.ts").then(m => m.Query),
  () => import("./protos/akash/audit/v1/service_akash.ts").then(m => m.Msg),
  () => import("./protos/akash/bme/v1/query_akash.ts").then(m => m.Query),
  () => import("./protos/akash/bme/v1/service_akash.ts").then(m => m.Msg),
  () => import("./protos/akash/cert/v1/query_akash.ts").then(m => m.Query),
  () => import("./protos/akash/cert/v1/service_akash.ts").then(m => m.Msg),
  () => import("./protos/akash/deployment/v1beta4/query_akash.ts").then(m => m.Query),
  () => import("./protos/akash/deployment/v1beta4/service_akash.ts").then(m => m.Msg),
  () => import("./protos/akash/discovery/v1/service_akash.ts").then(m => m.Discovery),
  () => import("./protos/akash/downtimedetector/v1beta1/query_akash.ts").then(m => m.Query),
  () => import("./protos/akash/epochs/v1beta1/query_akash.ts").then(m => m.Query),
  () => import("./protos/akash/escrow/v1/query_akash.ts").then(m => m.Query),
  () => import("./protos/akash/escrow/v1/service_akash.ts").then(m => m.Msg),
  () => import("./protos/akash/market/v1beta5/query_akash.ts").then(m => m.Query),
  () => import("./protos/akash/market/v1beta5/service_akash.ts").then(m => m.Msg),
  () => import("./protos/akash/oracle/v1/query_akash.ts").then(m => m.Query),
  () => import("./protos/akash/oracle/v1/service_akash.ts").then(m => m.Msg),
  () => import("./protos/akash/oracle/v2/query_akash.ts").then(m => m.Query),
  () => import("./protos/akash/oracle/v2/service_akash.ts").then(m => m.Msg),
  () => import("./protos/akash/provider/v1beta4/query_akash.ts").then(m => m.Query),
  () => import("./protos/akash/provider/v1beta4/service_akash.ts").then(m => m.Msg),
  () => import("./protos/akash/take/v1/query_akash.ts").then(m => m.Query),
  () => import("./protos/akash/take/v1/service_akash.ts").then(m => m.Msg),
  () => import("./protos/akash/verification/v1/query_akash.ts").then(m => m.Query),
  () => import("./protos/akash/verification/v1/service_akash.ts").then(m => m.Msg),
  () => import("./protos/akash/wasm/v1/query_akash.ts").then(m => m.Query),
  () => import("./protos/akash/wasm/v1/service_akash.ts").then(m => m.Msg)
] as const);
export function createSDK(queryTransport: Transport, txTransport: Transport) {
  const getClient = createClientFactory<CallOptions>(queryTransport);
  const getMsgClient = createClientFactory<TxCallOptions>(txTransport);
  return {
    akash: {
      audit: {
        v1: {
          /**
           * getAllProvidersAttributes queries all providers.
           */
          getAllProvidersAttributes: withMetadata(async function getAllProvidersAttributes(input: DeepPartial<akash_audit_v1_query.QueryAllProvidersAttributesRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(0);
            return getClient(service).allProvidersAttributes(input, options);
          }, { path: [0, "allProvidersAttributes"], serviceLoader }),
          /**
           * getProviderAttributes queries all provider signed attributes.
           */
          getProviderAttributes: withMetadata(async function getProviderAttributes(input: DeepPartial<akash_audit_v1_query.QueryProviderAttributesRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(0);
            return getClient(service).providerAttributes(input, options);
          }, { path: [0, "providerAttributes"], serviceLoader }),
          /**
           * getProviderAuditorAttributes queries provider signed attributes by specific auditor.
           */
          getProviderAuditorAttributes: withMetadata(async function getProviderAuditorAttributes(input: DeepPartial<akash_audit_v1_query.QueryProviderAuditorRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(0);
            return getClient(service).providerAuditorAttributes(input, options);
          }, { path: [0, "providerAuditorAttributes"], serviceLoader }),
          /**
           * getAuditorAttributes queries all providers signed by this auditor.
           */
          getAuditorAttributes: withMetadata(async function getAuditorAttributes(input: DeepPartial<akash_audit_v1_query.QueryAuditorAttributesRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(0);
            return getClient(service).auditorAttributes(input, options);
          }, { path: [0, "auditorAttributes"], serviceLoader }),
          /**
           * signProviderAttributes defines a method that signs provider attributes.
           */
          signProviderAttributes: withMetadata(async function signProviderAttributes(input: DeepSimplify<akash_audit_v1_msg.MsgSignProviderAttributes>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(1);
            return getMsgClient(service).signProviderAttributes(input, options);
          }, { path: [1, "signProviderAttributes"], serviceLoader }),
          /**
           * deleteProviderAttributes defines a method that deletes provider attributes.
           */
          deleteProviderAttributes: withMetadata(async function deleteProviderAttributes(input: DeepSimplify<akash_audit_v1_msg.MsgDeleteProviderAttributes>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(1);
            return getMsgClient(service).deleteProviderAttributes(input, options);
          }, { path: [1, "deleteProviderAttributes"], serviceLoader })
        }
      },
      bme: {
        v1: {
          /**
           * getParams returns the module parameters
           */
          getParams: withMetadata(async function getParams(input: DeepPartial<akash_bme_v1_query.QueryParamsRequest> = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(2);
            return getClient(service).params(input, options);
          }, { path: [2, "params"], serviceLoader }),
          /**
           * getVaultState returns the current vault state
           */
          getVaultState: withMetadata(async function getVaultState(input: DeepPartial<akash_bme_v1_query.QueryVaultStateRequest> = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(2);
            return getClient(service).vaultState(input, options);
          }, { path: [2, "vaultState"], serviceLoader }),
          /**
           * getStatus returns the current circuit breaker status
           */
          getStatus: withMetadata(async function getStatus(input: DeepPartial<akash_bme_v1_query.QueryStatusRequest> = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(2);
            return getClient(service).status(input, options);
          }, { path: [2, "status"], serviceLoader }),
          /**
           * getLedgerRecords queries ledger records with optional filters for status, source, denom, to_denom
           */
          getLedgerRecords: withMetadata(async function getLedgerRecords(input: DeepPartial<akash_bme_v1_query.QueryLedgerRecordsRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(2);
            return getClient(service).ledgerRecords(input, options);
          }, { path: [2, "ledgerRecords"], serviceLoader }),
          /**
           * updateParams updates the module parameters.
           * This operation can only be performed through governance proposals.
           */
          updateParams: withMetadata(async function updateParams(input: DeepSimplify<akash_bme_v1_msgs.MsgUpdateParams>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(3);
            return getMsgClient(service).updateParams(input, options);
          }, { path: [3, "updateParams"], serviceLoader }),
          /**
           * burnMint allows users to burn one token and mint another at current oracle prices.
           * Typically used to burn unused ACT tokens back to AKT.
           * The operation may be delayed or rejected based on circuit breaker status.
           */
          burnMint: withMetadata(async function burnMint(input: DeepSimplify<akash_bme_v1_msgs.MsgBurnMint>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(3);
            return getMsgClient(service).burnMint(input, options);
          }, { path: [3, "burnMint"], serviceLoader }),
          /**
           * mintACT mints ACT tokens by burning the specified source token.
           * The mint amount is calculated based on current oracle prices and
           * the collateral ratio. May be halted if circuit breaker is triggered.
           */
          mintACT: withMetadata(async function mintACT(input: DeepSimplify<akash_bme_v1_msgs.MsgMintACT>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(3);
            return getMsgClient(service).mintACT(input, options);
          }, { path: [3, "mintACT"], serviceLoader }),
          /**
           * burnACT burns ACT tokens and mints the specified destination token.
           * The burn operation uses remint credits when available, otherwise
           * requires adequate collateral backing based on oracle prices.
           */
          burnACT: withMetadata(async function burnACT(input: DeepSimplify<akash_bme_v1_msgs.MsgBurnACT>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(3);
            return getMsgClient(service).burnACT(input, options);
          }, { path: [3, "burnACT"], serviceLoader }),
          /**
           * fundVault seeds the BME vault with AKT from a designated source (e.g., community pool).
           * This provides the initial volatility buffer required for burn/mint operations.
           * Can only be executed through governance proposals.
           */
          fundVault: withMetadata(async function fundVault(input: DeepSimplify<akash_bme_v1_msgs.MsgFundVault>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(3);
            return getMsgClient(service).fundVault(input, options);
          }, { path: [3, "fundVault"], serviceLoader })
        }
      },
      cert: {
        v1: {
          /**
           * getCertificates queries certificates on-chain.
           */
          getCertificates: withMetadata(async function getCertificates(input: DeepPartial<akash_cert_v1_query.QueryCertificatesRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(4);
            return getClient(service).certificates(input, options);
          }, { path: [4, "certificates"], serviceLoader }),
          /**
           * createCertificate defines a method to create new certificate given proper inputs.
           */
          createCertificate: withMetadata(async function createCertificate(input: DeepSimplify<akash_cert_v1_msg.MsgCreateCertificate>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(5);
            return getMsgClient(service).createCertificate(input, options);
          }, { path: [5, "createCertificate"], serviceLoader }),
          /**
           * revokeCertificate defines a method to revoke the certificate.
           */
          revokeCertificate: withMetadata(async function revokeCertificate(input: DeepSimplify<akash_cert_v1_msg.MsgRevokeCertificate>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(5);
            return getMsgClient(service).revokeCertificate(input, options);
          }, { path: [5, "revokeCertificate"], serviceLoader })
        }
      },
      deployment: {
        v1beta4: {
          /**
           * getDeployments queries deployments.
           */
          getDeployments: withMetadata(async function getDeployments(input: DeepPartial<akash_deployment_v1beta4_query.QueryDeploymentsRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(6);
            return getClient(service).deployments(input, options);
          }, { path: [6, "deployments"], serviceLoader }),
          /**
           * getDeployment queries deployment details.
           */
          getDeployment: withMetadata(async function getDeployment(input: DeepPartial<akash_deployment_v1beta4_query.QueryDeploymentRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(6);
            return getClient(service).deployment(input, options);
          }, { path: [6, "deployment"], serviceLoader }),
          /**
           * getGroup queries group details.
           */
          getGroup: withMetadata(async function getGroup(input: DeepPartial<akash_deployment_v1beta4_query.QueryGroupRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(6);
            return getClient(service).group(input, options);
          }, { path: [6, "group"], serviceLoader }),
          /**
           * getParams returns the total set of deployment parameters.
           */
          getParams: withMetadata(async function getParams(input: DeepPartial<akash_deployment_v1beta4_query.QueryParamsRequest> = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(6);
            return getClient(service).params(input, options);
          }, { path: [6, "params"], serviceLoader }),
          /**
           * createDeployment defines a method to create new deployment given proper inputs.
           */
          createDeployment: withMetadata(async function createDeployment(input: DeepSimplify<akash_deployment_v1beta4_deploymentmsg.MsgCreateDeployment>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(7);
            return getMsgClient(service).createDeployment(input, options);
          }, { path: [7, "createDeployment"], serviceLoader }),
          /**
           * updateDeployment defines a method to update a deployment given proper inputs.
           */
          updateDeployment: withMetadata(async function updateDeployment(input: DeepSimplify<akash_deployment_v1beta4_deploymentmsg.MsgUpdateDeployment>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(7);
            return getMsgClient(service).updateDeployment(input, options);
          }, { path: [7, "updateDeployment"], serviceLoader }),
          /**
           * closeDeployment defines a method to close a deployment given proper inputs.
           */
          closeDeployment: withMetadata(async function closeDeployment(input: DeepSimplify<akash_deployment_v1beta4_deploymentmsg.MsgCloseDeployment>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(7);
            return getMsgClient(service).closeDeployment(input, options);
          }, { path: [7, "closeDeployment"], serviceLoader }),
          /**
           * closeGroup defines a method to close a group of a deployment given proper inputs.
           */
          closeGroup: withMetadata(async function closeGroup(input: DeepSimplify<akash_deployment_v1beta4_groupmsg.MsgCloseGroup>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(7);
            return getMsgClient(service).closeGroup(input, options);
          }, { path: [7, "closeGroup"], serviceLoader }),
          /**
           * pauseGroup defines a method to pause a group of a deployment given proper inputs.
           */
          pauseGroup: withMetadata(async function pauseGroup(input: DeepSimplify<akash_deployment_v1beta4_groupmsg.MsgPauseGroup>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(7);
            return getMsgClient(service).pauseGroup(input, options);
          }, { path: [7, "pauseGroup"], serviceLoader }),
          /**
           * startGroup defines a method to start a group of a deployment given proper inputs.
           */
          startGroup: withMetadata(async function startGroup(input: DeepSimplify<akash_deployment_v1beta4_groupmsg.MsgStartGroup>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(7);
            return getMsgClient(service).startGroup(input, options);
          }, { path: [7, "startGroup"], serviceLoader }),
          /**
           * updateParams defines a governance operation for updating the x/deployment module
           * parameters. The authority is hard-coded to the x/gov module account.
           *
           * Since: akash v1.0.0
           */
          updateParams: withMetadata(async function updateParams(input: DeepSimplify<akash_deployment_v1beta4_paramsmsg.MsgUpdateParams>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(7);
            return getMsgClient(service).updateParams(input, options);
          }, { path: [7, "updateParams"], serviceLoader })
        }
      },
      discovery: {
        v1: {
          /**
           * getInfo returns the node's supported API versions and metadata.
           */
          getInfo: withMetadata(async function getInfo(input: DeepPartial<akash_discovery_v1_service.GetInfoRequest> = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(8);
            return getClient(service).getInfo(input, options);
          }, { path: [8, "getInfo"], serviceLoader })
        }
      },
      downtimedetector: {
        v1beta1: {
          /**
           * getRecoveredSinceDowntimeOfLength queries if the chain has recovered for a specified duration
           * since experiencing downtime of a given length
           */
          getRecoveredSinceDowntimeOfLength: withMetadata(async function getRecoveredSinceDowntimeOfLength(input: DeepPartial<akash_downtimedetector_v1beta1_query.RecoveredSinceDowntimeOfLengthRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(9);
            return getClient(service).recoveredSinceDowntimeOfLength(input, options);
          }, { path: [9, "recoveredSinceDowntimeOfLength"], serviceLoader })
        }
      },
      epochs: {
        v1beta1: {
          /**
           * getEpochInfos provide running epochInfos
           */
          getEpochInfos: withMetadata(async function getEpochInfos(input: DeepPartial<akash_epochs_v1beta1_query.QueryEpochInfosRequest> = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(10);
            return getClient(service).epochInfos(input, options);
          }, { path: [10, "epochInfos"], serviceLoader }),
          /**
           * getCurrentEpoch provide current epoch of specified identifier
           */
          getCurrentEpoch: withMetadata(async function getCurrentEpoch(input: DeepPartial<akash_epochs_v1beta1_query.QueryCurrentEpochRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(10);
            return getClient(service).currentEpoch(input, options);
          }, { path: [10, "currentEpoch"], serviceLoader })
        }
      },
      escrow: {
        v1: {
          /**
           * getAccounts queries all accounts.
           */
          getAccounts: withMetadata(async function getAccounts(input: DeepPartial<akash_escrow_v1_query.QueryAccountsRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(11);
            return getClient(service).accounts(input, options);
          }, { path: [11, "accounts"], serviceLoader }),
          /**
           * getPayments queries all payments.
           */
          getPayments: withMetadata(async function getPayments(input: DeepPartial<akash_escrow_v1_query.QueryPaymentsRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(11);
            return getClient(service).payments(input, options);
          }, { path: [11, "payments"], serviceLoader }),
          /**
           * accountDeposit deposits more funds into the escrow account.
           */
          accountDeposit: withMetadata(async function accountDeposit(input: DeepSimplify<akash_escrow_v1_msg.MsgAccountDeposit>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(12);
            return getMsgClient(service).accountDeposit(input, options);
          }, { path: [12, "accountDeposit"], serviceLoader })
        }
      },
      market: {
        v1beta5: {
          /**
           * getOrders queries orders with filters.
           */
          getOrders: withMetadata(async function getOrders(input: DeepPartial<akash_market_v1beta5_query.QueryOrdersRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(13);
            return getClient(service).orders(input, options);
          }, { path: [13, "orders"], serviceLoader }),
          /**
           * getOrder queries order details.
           */
          getOrder: withMetadata(async function getOrder(input: DeepPartial<akash_market_v1beta5_query.QueryOrderRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(13);
            return getClient(service).order(input, options);
          }, { path: [13, "order"], serviceLoader }),
          /**
           * getBids queries bids with filters.
           */
          getBids: withMetadata(async function getBids(input: DeepPartial<akash_market_v1beta5_query.QueryBidsRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(13);
            return getClient(service).bids(input, options);
          }, { path: [13, "bids"], serviceLoader }),
          /**
           * getBid queries bid details.
           */
          getBid: withMetadata(async function getBid(input: DeepPartial<akash_market_v1beta5_query.QueryBidRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(13);
            return getClient(service).bid(input, options);
          }, { path: [13, "bid"], serviceLoader }),
          /**
           * getLeases queries leases with filters.
           */
          getLeases: withMetadata(async function getLeases(input: DeepPartial<akash_market_v1beta5_query.QueryLeasesRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(13);
            return getClient(service).leases(input, options);
          }, { path: [13, "leases"], serviceLoader }),
          /**
           * getLease queries lease details.
           */
          getLease: withMetadata(async function getLease(input: DeepPartial<akash_market_v1beta5_query.QueryLeaseRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(13);
            return getClient(service).lease(input, options);
          }, { path: [13, "lease"], serviceLoader }),
          /**
           * getParams returns the total set of market parameters.
           */
          getParams: withMetadata(async function getParams(input: DeepPartial<akash_market_v1beta5_query.QueryParamsRequest> = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(13);
            return getClient(service).params(input, options);
          }, { path: [13, "params"], serviceLoader }),
          /**
           * createBid defines a method to create a bid given proper inputs.
           */
          createBid: withMetadata(async function createBid(input: DeepSimplify<akash_market_v1beta5_bidmsg.MsgCreateBid>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(14);
            return getMsgClient(service).createBid(input, options);
          }, { path: [14, "createBid"], serviceLoader }),
          /**
           * closeBid defines a method to close a bid given proper inputs.
           */
          closeBid: withMetadata(async function closeBid(input: DeepSimplify<akash_market_v1beta5_bidmsg.MsgCloseBid>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(14);
            return getMsgClient(service).closeBid(input, options);
          }, { path: [14, "closeBid"], serviceLoader }),
          /**
           * withdrawLease withdraws accrued funds from the lease payment
           */
          withdrawLease: withMetadata(async function withdrawLease(input: DeepSimplify<akash_market_v1beta5_leasemsg.MsgWithdrawLease>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(14);
            return getMsgClient(service).withdrawLease(input, options);
          }, { path: [14, "withdrawLease"], serviceLoader }),
          /**
           * createLease creates a new lease
           */
          createLease: withMetadata(async function createLease(input: DeepSimplify<akash_market_v1beta5_leasemsg.MsgCreateLease>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(14);
            return getMsgClient(service).createLease(input, options);
          }, { path: [14, "createLease"], serviceLoader }),
          /**
           * closeLease defines a method to close an order given proper inputs.
           */
          closeLease: withMetadata(async function closeLease(input: DeepSimplify<akash_market_v1beta5_leasemsg.MsgCloseLease>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(14);
            return getMsgClient(service).closeLease(input, options);
          }, { path: [14, "closeLease"], serviceLoader }),
          /**
           * leaseStartReclaim initiates the reclamation window on an active lease.
           */
          leaseStartReclaim: withMetadata(async function leaseStartReclaim(input: DeepSimplify<akash_market_v1beta5_leasemsg.MsgLeaseStartReclaim>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(14);
            return getMsgClient(service).leaseStartReclaim(input, options);
          }, { path: [14, "leaseStartReclaim"], serviceLoader }),
          /**
           * updateParams defines a governance operation for updating the x/market module
           * parameters. The authority is hard-coded to the x/gov module account.
           *
           * Since: akash v1.0.0
           */
          updateParams: withMetadata(async function updateParams(input: DeepSimplify<akash_market_v1beta5_paramsmsg.MsgUpdateParams>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(14);
            return getMsgClient(service).updateParams(input, options);
          }, { path: [14, "updateParams"], serviceLoader })
        }
      },
      oracle: {
        v1: {
          /**
           * getPrices query prices for specific denom
           */
          getPrices: withMetadata(async function getPrices(input: DeepPartial<akash_oracle_v1_prices.QueryPricesRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(15);
            return getClient(service).prices(input, options);
          }, { path: [15, "prices"], serviceLoader }),
          /**
           * getParams returns the total set of oracle parameters.
           */
          getParams: withMetadata(async function getParams(input: DeepPartial<akash_oracle_v1_query.QueryParamsRequest> = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(15);
            return getClient(service).params(input, options);
          }, { path: [15, "params"], serviceLoader }),
          /**
           * getAggregatedPrice queries the aggregated price for a given denom.
           */
          getAggregatedPrice: withMetadata(async function getAggregatedPrice(input: DeepPartial<akash_oracle_v1_query.QueryAggregatedPriceRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(15);
            return getClient(service).aggregatedPrice(input, options);
          }, { path: [15, "aggregatedPrice"], serviceLoader }),
          /**
           * addPriceEntry adds a new price entry for a denomination from an authorized source
           */
          addPriceEntry: withMetadata(async function addPriceEntry(input: DeepSimplify<akash_oracle_v1_msgs.MsgAddPriceEntry>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(16);
            return getMsgClient(service).addPriceEntry(input, options);
          }, { path: [16, "addPriceEntry"], serviceLoader }),
          /**
           * updateParams defines a governance operation for updating the x/wasm module
           * parameters. The authority is hard-coded to the x/gov module account.
           *
           * Since: akash v2.0.0
           */
          updateParams: withMetadata(async function updateParams(input: DeepSimplify<akash_oracle_v1_msgs.MsgUpdateParams>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(16);
            return getMsgClient(service).updateParams(input, options);
          }, { path: [16, "updateParams"], serviceLoader })
        },
        v2: {
          /**
           * getPrices query prices for specific denom
           */
          getPrices: withMetadata(async function getPrices(input: DeepPartial<akash_oracle_v2_query.QueryPricesRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(17);
            return getClient(service).prices(input, options);
          }, { path: [17, "prices"], serviceLoader }),
          /**
           * getParams returns the total set of oracle parameters.
           */
          getParams: withMetadata(async function getParams(input: DeepPartial<akash_oracle_v2_query.QueryParamsRequest> = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(17);
            return getClient(service).params(input, options);
          }, { path: [17, "params"], serviceLoader }),
          /**
           * getAggregatedPrice queries the aggregated price for a given denom.
           */
          getAggregatedPrice: withMetadata(async function getAggregatedPrice(input: DeepPartial<akash_oracle_v2_query.QueryAggregatedPriceRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(17);
            return getClient(service).aggregatedPrice(input, options);
          }, { path: [17, "aggregatedPrice"], serviceLoader }),
          /**
           * addPriceEntry adds a new price entry for a denomination from an authorized source
           */
          addPriceEntry: withMetadata(async function addPriceEntry(input: DeepSimplify<akash_oracle_v2_msgs.MsgAddPriceEntry>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(18);
            return getMsgClient(service).addPriceEntry(input, options);
          }, { path: [18, "addPriceEntry"], serviceLoader }),
          /**
           * updateParams defines a governance operation for updating the x/oracle module
           * parameters. The authority is hard-coded to the x/gov module account.
           *
           * Since: akash v2.0.0
           */
          updateParams: withMetadata(async function updateParams(input: DeepSimplify<akash_oracle_v2_msgs.MsgUpdateParams>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(18);
            return getMsgClient(service).updateParams(input, options);
          }, { path: [18, "updateParams"], serviceLoader })
        }
      },
      provider: {
        v1beta4: {
          /**
           * getProviders queries providers
           */
          getProviders: withMetadata(async function getProviders(input: DeepPartial<akash_provider_v1beta4_query.QueryProvidersRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(19);
            return getClient(service).providers(input, options);
          }, { path: [19, "providers"], serviceLoader }),
          /**
           * getProvider queries provider details
           */
          getProvider: withMetadata(async function getProvider(input: DeepPartial<akash_provider_v1beta4_query.QueryProviderRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(19);
            return getClient(service).provider(input, options);
          }, { path: [19, "provider"], serviceLoader }),
          /**
           * getProviderMaintenance queries a provider maintenance record.
           */
          getProviderMaintenance: withMetadata(async function getProviderMaintenance(input: DeepPartial<akash_provider_v1beta4_query.QueryProviderMaintenanceRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(19);
            return getClient(service).providerMaintenance(input, options);
          }, { path: [19, "providerMaintenance"], serviceLoader }),
          /**
           * getProviderMaintenances queries provider maintenance records.
           */
          getProviderMaintenances: withMetadata(async function getProviderMaintenances(input: DeepPartial<akash_provider_v1beta4_query.QueryProviderMaintenancesRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(19);
            return getClient(service).providerMaintenances(input, options);
          }, { path: [19, "providerMaintenances"], serviceLoader }),
          /**
           * getParams returns the x/provider ProviderMaintenanceParams.
           */
          getParams: withMetadata(async function getParams(input: DeepPartial<akash_provider_v1beta4_query.QueryParamsRequest> = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(19);
            return getClient(service).params(input, options);
          }, { path: [19, "params"], serviceLoader }),
          /**
           * getRegistration queries provider registration details.
           */
          getRegistration: withMetadata(async function getRegistration(input: DeepPartial<akash_provider_v1beta4_query.QueryRegistrationRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(19);
            return getClient(service).registration(input, options);
          }, { path: [19, "registration"], serviceLoader }),
          /**
           * createProvider defines a method that creates a provider given the proper inputs.
           */
          createProvider: withMetadata(async function createProvider(input: DeepSimplify<akash_provider_v1beta4_msg.MsgCreateProvider>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(20);
            return getMsgClient(service).createProvider(input, options);
          }, { path: [20, "createProvider"], serviceLoader }),
          /**
           * updateProvider defines a method that updates a provider given the proper inputs.
           */
          updateProvider: withMetadata(async function updateProvider(input: DeepSimplify<akash_provider_v1beta4_msg.MsgUpdateProvider>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(20);
            return getMsgClient(service).updateProvider(input, options);
          }, { path: [20, "updateProvider"], serviceLoader }),
          /**
           * deleteProvider defines a method that deletes a provider given the proper inputs.
           */
          deleteProvider: withMetadata(async function deleteProvider(input: DeepSimplify<akash_provider_v1beta4_msg.MsgDeleteProvider>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(20);
            return getMsgClient(service).deleteProvider(input, options);
          }, { path: [20, "deleteProvider"], serviceLoader }),
          /**
           * openProviderMaintenance opens a provider maintenance window.
           */
          openProviderMaintenance: withMetadata(async function openProviderMaintenance(input: DeepSimplify<akash_provider_v1beta4_msg.MsgOpenProviderMaintenance>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(20);
            return getMsgClient(service).openProviderMaintenance(input, options);
          }, { path: [20, "openProviderMaintenance"], serviceLoader }),
          /**
           * closeProviderMaintenance closes an open maintenance window.
           */
          closeProviderMaintenance: withMetadata(async function closeProviderMaintenance(input: DeepSimplify<akash_provider_v1beta4_msg.MsgCloseProviderMaintenance>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(20);
            return getMsgClient(service).closeProviderMaintenance(input, options);
          }, { path: [20, "closeProviderMaintenance"], serviceLoader }),
          /**
           * updateParams is a governance operation for updating the x/provider
           * parameters.
           *
           * Since: akash v1.0.0
           */
          updateParams: withMetadata(async function updateParams(input: DeepSimplify<akash_provider_v1beta4_paramsmsg.MsgUpdateParams>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(20);
            return getMsgClient(service).updateParams(input, options);
          }, { path: [20, "updateParams"], serviceLoader })
        }
      },
      take: {
        v1: {
          /**
           * getParams returns the total set of take parameters.
           */
          getParams: withMetadata(async function getParams(input: DeepPartial<akash_take_v1_query.QueryParamsRequest> = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(21);
            return getClient(service).params(input, options);
          }, { path: [21, "params"], serviceLoader }),
          /**
           * updateParams defines a governance operation for updating the x/market module
           * parameters. The authority is hard-coded to the x/gov module account.
           *
           * Since: akash v1.0.0
           */
          updateParams: withMetadata(async function updateParams(input: DeepSimplify<akash_take_v1_paramsmsg.MsgUpdateParams>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(22);
            return getMsgClient(service).updateParams(input, options);
          }, { path: [22, "updateParams"], serviceLoader })
        }
      },
      verification: {
        v1: {
          /**
           * getAuditor returns the on-chain record for a single auditor identified by
           * its bech32 address.
           */
          getAuditor: withMetadata(async function getAuditor(input: DeepPartial<akash_verification_v1_query.QueryAuditorRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(23);
            return getClient(service).auditor(input, options);
          }, { path: [23, "auditor"], serviceLoader }),
          /**
           * getAuditors returns a paginated list of auditor records, optionally filtered
           * by AuditorStatus.
           */
          getAuditors: withMetadata(async function getAuditors(input: DeepPartial<akash_verification_v1_query.QueryAuditorsRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(23);
            return getClient(service).auditors(input, options);
          }, { path: [23, "auditors"], serviceLoader }),
          /**
           * getAttestation returns the attestation record for a specific
           * (provider, auditor) pair.
           */
          getAttestation: withMetadata(async function getAttestation(input: DeepPartial<akash_verification_v1_query.QueryAttestationRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(23);
            return getClient(service).attestation(input, options);
          }, { path: [23, "attestation"], serviceLoader }),
          /**
           * getProviderAttestations returns a paginated list of all attestation records
           * for the given provider, optionally filtered by AttestationStatus.
           */
          getProviderAttestations: withMetadata(async function getProviderAttestations(input: DeepPartial<akash_verification_v1_query.QueryProviderAttestationsRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(23);
            return getClient(service).providerAttestations(input, options);
          }, { path: [23, "providerAttestations"], serviceLoader }),
          /**
           * getAuditorAttestations returns a paginated list of all attestation records
           * submitted by the given auditor. No status filter is applied; callers can
           * page through every attestation the auditor has emitted.
           */
          getAuditorAttestations: withMetadata(async function getAuditorAttestations(input: DeepPartial<akash_verification_v1_query.QueryAuditorAttestationsRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(23);
            return getClient(service).auditorAttestations(input, options);
          }, { path: [23, "auditorAttestations"], serviceLoader }),
          /**
           * getDiscrepancy returns a single discrepancy event by its numeric id.
           */
          getDiscrepancy: withMetadata(async function getDiscrepancy(input: DeepPartial<akash_verification_v1_query.QueryDiscrepancyRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(23);
            return getClient(service).discrepancy(input, options);
          }, { path: [23, "discrepancy"], serviceLoader }),
          /**
           * getDiscrepancies returns a paginated list of discrepancy events, optionally
           * filtered by DiscrepancyStatus.
           */
          getDiscrepancies: withMetadata(async function getDiscrepancies(input: DeepPartial<akash_verification_v1_query.QueryDiscrepanciesRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(23);
            return getClient(service).discrepancies(input, options);
          }, { path: [23, "discrepancies"], serviceLoader }),
          /**
           * getAuditEscrow returns a single audit-escrow record by its numeric id.
           */
          getAuditEscrow: withMetadata(async function getAuditEscrow(input: DeepPartial<akash_verification_v1_query.QueryAuditEscrowRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(23);
            return getClient(service).auditEscrow(input, options);
          }, { path: [23, "auditEscrow"], serviceLoader }),
          /**
           * getProviderAuditEscrows returns a paginated list of audit-escrow records
           * opened by the given provider, optionally filtered by AuditEscrowStatus.
           */
          getProviderAuditEscrows: withMetadata(async function getProviderAuditEscrows(input: DeepPartial<akash_verification_v1_query.QueryProviderAuditEscrowsRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(23);
            return getClient(service).providerAuditEscrows(input, options);
          }, { path: [23, "providerAuditEscrows"], serviceLoader }),
          /**
           * getProviderVerificationGrace returns the verification-grace record for the
           * given provider, if one is currently tracked.
           */
          getProviderVerificationGrace: withMetadata(async function getProviderVerificationGrace(input: DeepPartial<akash_verification_v1_query.QueryProviderVerificationGraceRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(23);
            return getClient(service).providerVerificationGrace(input, options);
          }, { path: [23, "providerVerificationGrace"], serviceLoader }),
          /**
           * getProviderBond returns the provider's bond record.
           */
          getProviderBond: withMetadata(async function getProviderBond(input: DeepPartial<akash_verification_v1_query.QueryProviderBondRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(23);
            return getClient(service).providerBond(input, options);
          }, { path: [23, "providerBond"], serviceLoader }),
          /**
           * getProviderSnapshot returns the most recent provider snapshot record posted
           * on-chain by the given provider.
           */
          getProviderSnapshot: withMetadata(async function getProviderSnapshot(input: DeepPartial<akash_verification_v1_query.QueryProviderSnapshotRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(23);
            return getClient(service).providerSnapshot(input, options);
          }, { path: [23, "providerSnapshot"], serviceLoader }),
          /**
           * getParams returns the current parameter set for the verification module.
           */
          getParams: withMetadata(async function getParams(input: DeepPartial<akash_verification_v1_query.QueryParamsRequest> = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(23);
            return getClient(service).params(input, options);
          }, { path: [23, "params"], serviceLoader }),
          /**
           * postAuditorBond posts (or tops up) an auditor's verification bond.
           */
          postAuditorBond: withMetadata(async function postAuditorBond(input: DeepSimplify<akash_verification_v1_msg.MsgPostAuditorBond>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(24);
            return getMsgClient(service).postAuditorBond(input, options);
          }, { path: [24, "postAuditorBond"], serviceLoader }),
          /**
           * openAuditEscrow opens a new audit escrow funding a pending attestation.
           */
          openAuditEscrow: withMetadata(async function openAuditEscrow(input: DeepSimplify<akash_verification_v1_msg.MsgOpenAuditEscrow>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(24);
            return getMsgClient(service).openAuditEscrow(input, options);
          }, { path: [24, "openAuditEscrow"], serviceLoader }),
          /**
           * cancelAuditEscrow cancels an open, unconsumed audit escrow before expiry
           * and returns the fee and provider deposit to the provider.
           */
          cancelAuditEscrow: withMetadata(async function cancelAuditEscrow(input: DeepSimplify<akash_verification_v1_msg.MsgCancelAuditEscrow>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(24);
            return getMsgClient(service).cancelAuditEscrow(input, options);
          }, { path: [24, "cancelAuditEscrow"], serviceLoader }),
          /**
           * settleAuditEscrow settles an unconsumed audit escrow with an explicit
           * reason, fault attribution, and evidence reference.
           */
          settleAuditEscrow: withMetadata(async function settleAuditEscrow(input: DeepSimplify<akash_verification_v1_msg.MsgSettleAuditEscrow>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(24);
            return getMsgClient(service).settleAuditEscrow(input, options);
          }, { path: [24, "settleAuditEscrow"], serviceLoader }),
          /**
           * submitAttestation submits an attestation about a provider; the first
           * valid submission against a matching open escrow consumes it.
           */
          submitAttestation: withMetadata(async function submitAttestation(input: DeepSimplify<akash_verification_v1_msg.MsgSubmitAttestation>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(24);
            return getMsgClient(service).submitAttestation(input, options);
          }, { path: [24, "submitAttestation"], serviceLoader }),
          /**
           * revokeAttestation revokes a previously submitted attestation with a
           * typed reason and evidence reference.
           */
          revokeAttestation: withMetadata(async function revokeAttestation(input: DeepSimplify<akash_verification_v1_msg.MsgRevokeAttestation>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(24);
            return getMsgClient(service).revokeAttestation(input, options);
          }, { path: [24, "revokeAttestation"], serviceLoader }),
          /**
           * removeAttestation voluntarily removes an attestation associated with the
           * signing provider.
           */
          removeAttestation: withMetadata(async function removeAttestation(input: DeepSimplify<akash_verification_v1_msg.MsgRemoveAttestation>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(24);
            return getMsgClient(service).removeAttestation(input, options);
          }, { path: [24, "removeAttestation"], serviceLoader }),
          /**
           * resignAuditor voluntarily exits the auditor role and begins unbonding
           * of any posted auditor bond.
           */
          resignAuditor: withMetadata(async function resignAuditor(input: DeepSimplify<akash_verification_v1_msg.MsgResignAuditor>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(24);
            return getMsgClient(service).resignAuditor(input, options);
          }, { path: [24, "resignAuditor"], serviceLoader }),
          /**
           * postProviderBond posts (or tops up) a provider's resource-scaled
           * verification bond.
           */
          postProviderBond: withMetadata(async function postProviderBond(input: DeepSimplify<akash_verification_v1_msg.MsgPostProviderBond>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(24);
            return getMsgClient(service).postProviderBond(input, options);
          }, { path: [24, "postProviderBond"], serviceLoader }),
          /**
           * withdrawProviderBond initiates withdrawal of part or all of a provider's
           * verification bond.
           */
          withdrawProviderBond: withMetadata(async function withdrawProviderBond(input: DeepSimplify<akash_verification_v1_msg.MsgWithdrawProviderBond>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(24);
            return getMsgClient(service).withdrawProviderBond(input, options);
          }, { path: [24, "withdrawProviderBond"], serviceLoader }),
          /**
           * postSnapshotHash posts the provider's most recent resource snapshot hash
           * and inline resource summary.
           */
          postSnapshotHash: withMetadata(async function postSnapshotHash(input: DeepSimplify<akash_verification_v1_msg.MsgPostSnapshotHash>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(24);
            return getMsgClient(service).postSnapshotHash(input, options);
          }, { path: [24, "postSnapshotHash"], serviceLoader }),
          /**
           * registerAuditor registers a new auditor with a maximum attestation tier;
           * governance only.
           */
          registerAuditor: withMetadata(async function registerAuditor(input: DeepSimplify<akash_verification_v1_msg.MsgRegisterAuditor>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(24);
            return getMsgClient(service).registerAuditor(input, options);
          }, { path: [24, "registerAuditor"], serviceLoader }),
          /**
           * renewAuditor renews an auditor's registration and resets the renewal
           * deadline; governance only.
           */
          renewAuditor: withMetadata(async function renewAuditor(input: DeepSimplify<akash_verification_v1_msg.MsgRenewAuditor>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(24);
            return getMsgClient(service).renewAuditor(input, options);
          }, { path: [24, "renewAuditor"], serviceLoader }),
          /**
           * removeAuditor removes an auditor from the active set; governance only.
           */
          removeAuditor: withMetadata(async function removeAuditor(input: DeepSimplify<akash_verification_v1_msg.MsgRemoveAuditor>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(24);
            return getMsgClient(service).removeAuditor(input, options);
          }, { path: [24, "removeAuditor"], serviceLoader }),
          /**
           * revokeProviderAttestation revokes a single attestation for a specific
           * provider/auditor pair; governance only.
           */
          revokeProviderAttestation: withMetadata(async function revokeProviderAttestation(input: DeepSimplify<akash_verification_v1_msg.MsgRevokeProviderAttestation>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(24);
            return getMsgClient(service).revokeProviderAttestation(input, options);
          }, { path: [24, "revokeProviderAttestation"], serviceLoader }),
          /**
           * revokeAllProviderAttestations revokes every active attestation for a
           * single provider; governance only.
           */
          revokeAllProviderAttestations: withMetadata(async function revokeAllProviderAttestations(input: DeepSimplify<akash_verification_v1_msg.MsgRevokeAllProviderAttestations>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(24);
            return getMsgClient(service).revokeAllProviderAttestations(input, options);
          }, { path: [24, "revokeAllProviderAttestations"], serviceLoader }),
          /**
           * revokeAuditorAttestations revokes every active attestation issued by a
           * single auditor; governance only.
           */
          revokeAuditorAttestations: withMetadata(async function revokeAuditorAttestations(input: DeepSimplify<akash_verification_v1_msg.MsgRevokeAuditorAttestations>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(24);
            return getMsgClient(service).revokeAuditorAttestations(input, options);
          }, { path: [24, "revokeAuditorAttestations"], serviceLoader }),
          /**
           * resolveDiscrepancy resolves a pending discrepancy between two auditors
           * over the same provider; governance only.
           */
          resolveDiscrepancy: withMetadata(async function resolveDiscrepancy(input: DeepSimplify<akash_verification_v1_msg.MsgResolveDiscrepancy>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(24);
            return getMsgClient(service).resolveDiscrepancy(input, options);
          }, { path: [24, "resolveDiscrepancy"], serviceLoader }),
          /**
           * slashProviderBond slashes a fraction of a provider's verification bond;
           * governance only.
           */
          slashProviderBond: withMetadata(async function slashProviderBond(input: DeepSimplify<akash_verification_v1_msg.MsgSlashProviderBond>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(24);
            return getMsgClient(service).slashProviderBond(input, options);
          }, { path: [24, "slashProviderBond"], serviceLoader }),
          /**
           * updateParams updates the x/verification module parameters; governance
           * only.
           */
          updateParams: withMetadata(async function updateParams(input: DeepSimplify<akash_verification_v1_msg.MsgUpdateParams>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(24);
            return getMsgClient(service).updateParams(input, options);
          }, { path: [24, "updateParams"], serviceLoader })
        }
      },
      wasm: {
        v1: {
          /**
           * getParams returns the total set of wasm parameters.
           */
          getParams: withMetadata(async function getParams(input: DeepPartial<akash_wasm_v1_query.QueryParamsRequest> = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(25);
            return getClient(service).params(input, options);
          }, { path: [25, "params"], serviceLoader }),
          /**
           * updateParams defines a governance operation for updating the x/wasm module
           * parameters. The authority is hard-coded to the x/gov module account.
           *
           * Since: akash v2.0.0
           */
          updateParams: withMetadata(async function updateParams(input: DeepSimplify<akash_wasm_v1_paramsmsg.MsgUpdateParams>, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(26);
            return getMsgClient(service).updateParams(input, options);
          }, { path: [26, "updateParams"], serviceLoader })
        }
      }
    }
  };
}
