import { createServiceLoader } from "../client/createServiceLoader.ts";

import type * as akash_audit_v1_query from "./protos/akash/audit/v1/query.ts";
import type * as akash_audit_v1_msg from "./protos/akash/audit/v1/msg.ts";
import type * as akash_cert_v1_query from "./protos/akash/cert/v1/query.ts";
import type * as akash_cert_v1_msg from "./protos/akash/cert/v1/msg.ts";
import type * as akash_deployment_v1beta4_query from "./protos/akash/deployment/v1beta4/query.ts";
import type * as akash_deployment_v1beta4_deploymentmsg from "./protos/akash/deployment/v1beta4/deploymentmsg.ts";
import type * as akash_deployment_v1beta4_groupmsg from "./protos/akash/deployment/v1beta4/groupmsg.ts";
import type * as akash_deployment_v1beta4_paramsmsg from "./protos/akash/deployment/v1beta4/paramsmsg.ts";
import type * as akash_escrow_v1_query from "./protos/akash/escrow/v1/query.ts";
import type * as akash_escrow_v1_msg from "./protos/akash/escrow/v1/msg.ts";
import type * as akash_market_v1beta5_query from "./protos/akash/market/v1beta5/query.ts";
import type * as akash_market_v1beta5_bidmsg from "./protos/akash/market/v1beta5/bidmsg.ts";
import type * as akash_market_v1beta5_leasemsg from "./protos/akash/market/v1beta5/leasemsg.ts";
import type * as akash_market_v1beta5_paramsmsg from "./protos/akash/market/v1beta5/paramsmsg.ts";
import type * as akash_provider_v1beta4_query from "./protos/akash/provider/v1beta4/query.ts";
import type * as akash_provider_v1beta4_msg from "./protos/akash/provider/v1beta4/msg.ts";
import type * as akash_take_v1_query from "./protos/akash/take/v1/query.ts";
import type * as akash_take_v1_paramsmsg from "./protos/akash/take/v1/paramsmsg.ts";
import { createClientFactory } from "../client/createClientFactory.ts";
import type { Transport, CallOptions, TxCallOptions } from "../transport/types.ts";
import type { SDKOptions } from "../sdk/types.ts";
import { withMetadata } from "../utils/sdkMetadata.ts";


export const serviceLoader= createServiceLoader([
  () => import("./protos/akash/audit/v1/query_akash.ts").then(m => m.Query),
  () => import("./protos/akash/audit/v1/service_akash.ts").then(m => m.Msg),
  () => import("./protos/akash/cert/v1/query_akash.ts").then(m => m.Query),
  () => import("./protos/akash/cert/v1/service_akash.ts").then(m => m.Msg),
  () => import("./protos/akash/deployment/v1beta4/query_akash.ts").then(m => m.Query),
  () => import("./protos/akash/deployment/v1beta4/service_akash.ts").then(m => m.Msg),
  () => import("./protos/akash/escrow/v1/query_akash.ts").then(m => m.Query),
  () => import("./protos/akash/escrow/v1/service_akash.ts").then(m => m.Msg),
  () => import("./protos/akash/market/v1beta5/query_akash.ts").then(m => m.Query),
  () => import("./protos/akash/market/v1beta5/service_akash.ts").then(m => m.Msg),
  () => import("./protos/akash/provider/v1beta4/query_akash.ts").then(m => m.Query),
  () => import("./protos/akash/provider/v1beta4/service_akash.ts").then(m => m.Msg),
  () => import("./protos/akash/take/v1/query_akash.ts").then(m => m.Query),
  () => import("./protos/akash/take/v1/service_akash.ts").then(m => m.Msg)
] as const);
export function createSDK(queryTransport: Transport, txTransport: Transport, options?: SDKOptions) {
  const getClient = createClientFactory<CallOptions>(queryTransport, options?.clientOptions);
  const getMsgClient = createClientFactory<TxCallOptions>(txTransport, options?.clientOptions);
  return {
    akash: {
      audit: {
        v1: {
          /**
           * getAllProvidersAttributes queries all providers.
           */
          getAllProvidersAttributes: withMetadata(async function getAllProvidersAttributes(input: akash_audit_v1_query.QueryAllProvidersAttributesRequest, options?: CallOptions) {
            const service = await serviceLoader.loadAt(0);
            return getClient(service).allProvidersAttributes(input, options);
          }, { path: [0, 0] }),
          /**
           * getProviderAttributes queries all provider signed attributes.
           */
          getProviderAttributes: withMetadata(async function getProviderAttributes(input: akash_audit_v1_query.QueryProviderAttributesRequest, options?: CallOptions) {
            const service = await serviceLoader.loadAt(0);
            return getClient(service).providerAttributes(input, options);
          }, { path: [0, 1] }),
          /**
           * getProviderAuditorAttributes queries provider signed attributes by specific auditor.
           */
          getProviderAuditorAttributes: withMetadata(async function getProviderAuditorAttributes(input: akash_audit_v1_query.QueryProviderAuditorRequest, options?: CallOptions) {
            const service = await serviceLoader.loadAt(0);
            return getClient(service).providerAuditorAttributes(input, options);
          }, { path: [0, 2] }),
          /**
           * getAuditorAttributes queries all providers signed by this auditor.
           */
          getAuditorAttributes: withMetadata(async function getAuditorAttributes(input: akash_audit_v1_query.QueryAuditorAttributesRequest, options?: CallOptions) {
            const service = await serviceLoader.loadAt(0);
            return getClient(service).auditorAttributes(input, options);
          }, { path: [0, 3] }),
          /**
           * signProviderAttributes defines a method that signs provider attributes.
           */
          signProviderAttributes: withMetadata(async function signProviderAttributes(input: akash_audit_v1_msg.MsgSignProviderAttributes, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(1);
            return getMsgClient(service).signProviderAttributes(input, options);
          }, { path: [1, 0] }),
          /**
           * deleteProviderAttributes defines a method that deletes provider attributes.
           */
          deleteProviderAttributes: withMetadata(async function deleteProviderAttributes(input: akash_audit_v1_msg.MsgDeleteProviderAttributes, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(1);
            return getMsgClient(service).deleteProviderAttributes(input, options);
          }, { path: [1, 1] })
        }
      },
      cert: {
        v1: {
          /**
           * getCertificates queries certificates on-chain.
           */
          getCertificates: withMetadata(async function getCertificates(input: akash_cert_v1_query.QueryCertificatesRequest, options?: CallOptions) {
            const service = await serviceLoader.loadAt(2);
            return getClient(service).certificates(input, options);
          }, { path: [2, 0] }),
          /**
           * createCertificate defines a method to create new certificate given proper inputs.
           */
          createCertificate: withMetadata(async function createCertificate(input: akash_cert_v1_msg.MsgCreateCertificate, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(3);
            return getMsgClient(service).createCertificate(input, options);
          }, { path: [3, 0] }),
          /**
           * revokeCertificate defines a method to revoke the certificate.
           */
          revokeCertificate: withMetadata(async function revokeCertificate(input: akash_cert_v1_msg.MsgRevokeCertificate, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(3);
            return getMsgClient(service).revokeCertificate(input, options);
          }, { path: [3, 1] })
        }
      },
      deployment: {
        v1beta4: {
          /**
           * getDeployments queries deployments.
           */
          getDeployments: withMetadata(async function getDeployments(input: akash_deployment_v1beta4_query.QueryDeploymentsRequest, options?: CallOptions) {
            const service = await serviceLoader.loadAt(4);
            return getClient(service).deployments(input, options);
          }, { path: [4, 0] }),
          /**
           * getDeployment queries deployment details.
           */
          getDeployment: withMetadata(async function getDeployment(input: akash_deployment_v1beta4_query.QueryDeploymentRequest, options?: CallOptions) {
            const service = await serviceLoader.loadAt(4);
            return getClient(service).deployment(input, options);
          }, { path: [4, 1] }),
          /**
           * getGroup queries group details.
           */
          getGroup: withMetadata(async function getGroup(input: akash_deployment_v1beta4_query.QueryGroupRequest, options?: CallOptions) {
            const service = await serviceLoader.loadAt(4);
            return getClient(service).group(input, options);
          }, { path: [4, 2] }),
          /**
           * getParams returns the total set of minting parameters.
           */
          getParams: withMetadata(async function getParams(input: akash_deployment_v1beta4_query.QueryParamsRequest = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(4);
            return getClient(service).params(input, options);
          }, { path: [4, 3] }),
          /**
           * createDeployment defines a method to create new deployment given proper inputs.
           */
          createDeployment: withMetadata(async function createDeployment(input: akash_deployment_v1beta4_deploymentmsg.MsgCreateDeployment, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(5);
            return getMsgClient(service).createDeployment(input, options);
          }, { path: [5, 0] }),
          /**
           * updateDeployment defines a method to update a deployment given proper inputs.
           */
          updateDeployment: withMetadata(async function updateDeployment(input: akash_deployment_v1beta4_deploymentmsg.MsgUpdateDeployment, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(5);
            return getMsgClient(service).updateDeployment(input, options);
          }, { path: [5, 1] }),
          /**
           * closeDeployment defines a method to close a deployment given proper inputs.
           */
          closeDeployment: withMetadata(async function closeDeployment(input: akash_deployment_v1beta4_deploymentmsg.MsgCloseDeployment, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(5);
            return getMsgClient(service).closeDeployment(input, options);
          }, { path: [5, 2] }),
          /**
           * closeGroup defines a method to close a group of a deployment given proper inputs.
           */
          closeGroup: withMetadata(async function closeGroup(input: akash_deployment_v1beta4_groupmsg.MsgCloseGroup, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(5);
            return getMsgClient(service).closeGroup(input, options);
          }, { path: [5, 3] }),
          /**
           * pauseGroup defines a method to pause a group of a deployment given proper inputs.
           */
          pauseGroup: withMetadata(async function pauseGroup(input: akash_deployment_v1beta4_groupmsg.MsgPauseGroup, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(5);
            return getMsgClient(service).pauseGroup(input, options);
          }, { path: [5, 4] }),
          /**
           * startGroup defines a method to start a group of a deployment given proper inputs.
           */
          startGroup: withMetadata(async function startGroup(input: akash_deployment_v1beta4_groupmsg.MsgStartGroup, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(5);
            return getMsgClient(service).startGroup(input, options);
          }, { path: [5, 5] }),
          /**
           * updateParams defines a governance operation for updating the x/deployment module
           * parameters. The authority is hard-coded to the x/gov module account.
           *
           * Since: akash v1.0.0
           */
          updateParams: withMetadata(async function updateParams(input: akash_deployment_v1beta4_paramsmsg.MsgUpdateParams, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(5);
            return getMsgClient(service).updateParams(input, options);
          }, { path: [5, 6] })
        }
      },
      escrow: {
        v1: {
          /**
           * getAccounts queries all accounts.
           */
          getAccounts: withMetadata(async function getAccounts(input: akash_escrow_v1_query.QueryAccountsRequest, options?: CallOptions) {
            const service = await serviceLoader.loadAt(6);
            return getClient(service).accounts(input, options);
          }, { path: [6, 0] }),
          /**
           * getPayments queries all payments.
           */
          getPayments: withMetadata(async function getPayments(input: akash_escrow_v1_query.QueryPaymentsRequest, options?: CallOptions) {
            const service = await serviceLoader.loadAt(6);
            return getClient(service).payments(input, options);
          }, { path: [6, 1] }),
          /**
           * accountDeposit deposits more funds into the escrow account.
           */
          accountDeposit: withMetadata(async function accountDeposit(input: akash_escrow_v1_msg.MsgAccountDeposit, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(7);
            return getMsgClient(service).accountDeposit(input, options);
          }, { path: [7, 0] })
        }
      },
      market: {
        v1beta5: {
          /**
           * getOrders queries orders with filters.
           */
          getOrders: withMetadata(async function getOrders(input: akash_market_v1beta5_query.QueryOrdersRequest, options?: CallOptions) {
            const service = await serviceLoader.loadAt(8);
            return getClient(service).orders(input, options);
          }, { path: [8, 0] }),
          /**
           * getOrder queries order details.
           */
          getOrder: withMetadata(async function getOrder(input: akash_market_v1beta5_query.QueryOrderRequest, options?: CallOptions) {
            const service = await serviceLoader.loadAt(8);
            return getClient(service).order(input, options);
          }, { path: [8, 1] }),
          /**
           * getBids queries bids with filters.
           */
          getBids: withMetadata(async function getBids(input: akash_market_v1beta5_query.QueryBidsRequest, options?: CallOptions) {
            const service = await serviceLoader.loadAt(8);
            return getClient(service).bids(input, options);
          }, { path: [8, 2] }),
          /**
           * getBid queries bid details.
           */
          getBid: withMetadata(async function getBid(input: akash_market_v1beta5_query.QueryBidRequest, options?: CallOptions) {
            const service = await serviceLoader.loadAt(8);
            return getClient(service).bid(input, options);
          }, { path: [8, 3] }),
          /**
           * getLeases queries leases with filters.
           */
          getLeases: withMetadata(async function getLeases(input: akash_market_v1beta5_query.QueryLeasesRequest, options?: CallOptions) {
            const service = await serviceLoader.loadAt(8);
            return getClient(service).leases(input, options);
          }, { path: [8, 4] }),
          /**
           * getLease queries lease details.
           */
          getLease: withMetadata(async function getLease(input: akash_market_v1beta5_query.QueryLeaseRequest, options?: CallOptions) {
            const service = await serviceLoader.loadAt(8);
            return getClient(service).lease(input, options);
          }, { path: [8, 5] }),
          /**
           * getParams returns the total set of minting parameters.
           */
          getParams: withMetadata(async function getParams(input: akash_market_v1beta5_query.QueryParamsRequest = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(8);
            return getClient(service).params(input, options);
          }, { path: [8, 6] }),
          /**
           * createBid defines a method to create a bid given proper inputs.
           */
          createBid: withMetadata(async function createBid(input: akash_market_v1beta5_bidmsg.MsgCreateBid, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(9);
            return getMsgClient(service).createBid(input, options);
          }, { path: [9, 0] }),
          /**
           * closeBid defines a method to close a bid given proper inputs.
           */
          closeBid: withMetadata(async function closeBid(input: akash_market_v1beta5_bidmsg.MsgCloseBid, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(9);
            return getMsgClient(service).closeBid(input, options);
          }, { path: [9, 1] }),
          /**
           * withdrawLease withdraws accrued funds from the lease payment
           */
          withdrawLease: withMetadata(async function withdrawLease(input: akash_market_v1beta5_leasemsg.MsgWithdrawLease, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(9);
            return getMsgClient(service).withdrawLease(input, options);
          }, { path: [9, 2] }),
          /**
           * createLease creates a new lease
           */
          createLease: withMetadata(async function createLease(input: akash_market_v1beta5_leasemsg.MsgCreateLease, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(9);
            return getMsgClient(service).createLease(input, options);
          }, { path: [9, 3] }),
          /**
           * closeLease defines a method to close an order given proper inputs.
           */
          closeLease: withMetadata(async function closeLease(input: akash_market_v1beta5_leasemsg.MsgCloseLease, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(9);
            return getMsgClient(service).closeLease(input, options);
          }, { path: [9, 4] }),
          /**
           * updateParams defines a governance operation for updating the x/market module
           * parameters. The authority is hard-coded to the x/gov module account.
           *
           * Since: akash v1.0.0
           */
          updateParams: withMetadata(async function updateParams(input: akash_market_v1beta5_paramsmsg.MsgUpdateParams, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(9);
            return getMsgClient(service).updateParams(input, options);
          }, { path: [9, 5] })
        }
      },
      provider: {
        v1beta4: {
          /**
           * getProviders queries providers
           */
          getProviders: withMetadata(async function getProviders(input: akash_provider_v1beta4_query.QueryProvidersRequest, options?: CallOptions) {
            const service = await serviceLoader.loadAt(10);
            return getClient(service).providers(input, options);
          }, { path: [10, 0] }),
          /**
           * getProvider queries provider details
           */
          getProvider: withMetadata(async function getProvider(input: akash_provider_v1beta4_query.QueryProviderRequest, options?: CallOptions) {
            const service = await serviceLoader.loadAt(10);
            return getClient(service).provider(input, options);
          }, { path: [10, 1] }),
          /**
           * createProvider defines a method that creates a provider given the proper inputs.
           */
          createProvider: withMetadata(async function createProvider(input: akash_provider_v1beta4_msg.MsgCreateProvider, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(11);
            return getMsgClient(service).createProvider(input, options);
          }, { path: [11, 0] }),
          /**
           * updateProvider defines a method that updates a provider given the proper inputs.
           */
          updateProvider: withMetadata(async function updateProvider(input: akash_provider_v1beta4_msg.MsgUpdateProvider, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(11);
            return getMsgClient(service).updateProvider(input, options);
          }, { path: [11, 1] }),
          /**
           * deleteProvider defines a method that deletes a provider given the proper inputs.
           */
          deleteProvider: withMetadata(async function deleteProvider(input: akash_provider_v1beta4_msg.MsgDeleteProvider, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(11);
            return getMsgClient(service).deleteProvider(input, options);
          }, { path: [11, 2] })
        }
      },
      take: {
        v1: {
          /**
           * getParams returns the total set of minting parameters.
           */
          getParams: withMetadata(async function getParams(input: akash_take_v1_query.QueryParamsRequest = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(12);
            return getClient(service).params(input, options);
          }, { path: [12, 0] }),
          /**
           * updateParams defines a governance operation for updating the x/market module
           * parameters. The authority is hard-coded to the x/gov module account.
           *
           * Since: akash v1.0.0
           */
          updateParams: withMetadata(async function updateParams(input: akash_take_v1_paramsmsg.MsgUpdateParams, options?: TxCallOptions) {
            const service = await serviceLoader.loadAt(13);
            return getMsgClient(service).updateParams(input, options);
          }, { path: [13, 0] })
        }
      }
    }
  };
}
