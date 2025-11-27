
import { describe, expect, it, beforeAll } from "@jest/globals";
import Long from "long";
import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";

import { createChainNodeSDK } from "../../src/sdk/chain/createChainNodeSDK.ts";
import { createStargateClient } from "../../src/sdk/transport/tx/createStargateClient/createStargateClient.ts";
import { MsgCreateDeployment } from "../../src/generated/protos/akash/deployment/v1beta4/deploymentmsg.ts";
import { MsgCreateLease } from "../../src/generated/protos/akash/market/v1beta5/leasemsg.ts";
import { BidID } from "../../src/generated/protos/akash/market/v1/bid.ts";
import { Storage } from "../../src/generated/protos/akash/base/resources/v1beta4/storage.ts";
import { Source } from "../../src/generated/protos/akash/base/deposit/v1/deposit.ts";
import { Coin, DecCoin } from "../../src/generated/protos/cosmos/base/v1beta1/coin.ts";

declare const jest: {
  setTimeout: (timeout: number) => void;
};

describe("Lease Operations", () => {
  jest.setTimeout(60000);

  const QUERY_RPC_URL = process.env.QUERY_RPC_URL || process.env.TX_RPC_URL || "http://grpc.sandbox-2.aksh.pw:9090";
  const TX_RPC_URL = process.env.TX_RPC_URL || "https://rpc.sandbox-2.aksh.pw:443";

  const createTestSDK = async (wallet?: DirectSecp256k1HdWallet) => {
    const txClient = wallet ? await createStargateClient({
      baseUrl: TX_RPC_URL,
      signer: wallet,
    }) : undefined;
    
    return createChainNodeSDK({
      query: { baseUrl: QUERY_RPC_URL },
      tx: txClient ? { signer: txClient } : undefined,
    });
  };

  const createResourceValue = (value: string): { val: Uint8Array } => ({
    val: new TextEncoder().encode(value)
  });

  const wait = (ms: number) => new Promise(resolve => setTimeout(resolve, ms));


  it("should create a deployment, wait for bids, select first bid and create a lease", async () => {
    const testMnemonic = process.env.TEST_MNEMONIC;
    
    if (!testMnemonic) {
      return;
    }
    
    const wallet = await DirectSecp256k1HdWallet.fromMnemonic(testMnemonic, { prefix: "akash" });
    const [account] = await wallet.getAccounts();
    
    const sdk = await createTestSDK(wallet);

    const deploymentMessage: MsgCreateDeployment = {
      id: {
        owner: account.address,
        dseq: Long.fromNumber(Date.now()) // Use timestamp for uniqueness
      },
      groups: [{
        name: "web-service",
        requirements: {
          signedBy: {
            allOf: [],
            anyOf: []
          },
          attributes: []
        },
        resources: [{
          resource: {
            id: 1,
            cpu: {
              units: createResourceValue("1000"),
              attributes: []
            },
            memory: {
              quantity: createResourceValue("1073741824"),
              attributes: []
            },
            storage: [{
              name: "main",
              quantity: createResourceValue("2147483648"),
              attributes: []
            } as Storage],
            gpu: {
              units: createResourceValue("0"),
              attributes: []
            },
            endpoints: []
          },
          count: 1,
          price: {
            denom: "uakt",
            amount: "100000"
          } as DecCoin
        }]
      }],
      hash: new Uint8Array(32),
      deposit: {
        amount: {
          denom: "uakt",
          amount: "500000"
        } as Coin,
        sources: [Source.balance]
      }
    };

    const deploymentResult = await sdk.akash.deployment.v1beta4.createDeployment(deploymentMessage, {
      memo: "Test deployment for lease creation - Akash Chain SDK"
    });
    
    expect(deploymentResult).toBeDefined();

    const deploymentId = {
      owner: account.address,
      dseq: deploymentMessage.id!.dseq
    };

    let bidsResponse;
    let attempts = 0;
    const maxAttempts = 3;
    
    do {
      await wait(6000);
      attempts++;
      
      bidsResponse = await sdk.akash.market.v1beta5.getBids({
        filters: {
          owner: deploymentId.owner,
          dseq: deploymentId.dseq,
          gseq: 1,
          oseq: 1,
        }
      });
      
    } while ((!bidsResponse?.bids || bidsResponse.bids.length < 2) && attempts < maxAttempts);


    expect(bidsResponse?.bids).toBeDefined();
    expect(Array.isArray(bidsResponse?.bids)).toBe(true);
    
    if (bidsResponse!.bids!.length === 0) {
      throw new Error(`No bids found after ${maxAttempts} attempts. Check deployment resources and pricing.`);
    }
    
    expect(bidsResponse!.bids!.length).toBeGreaterThan(0);
    
    const firstBid = bidsResponse!.bids![0]!.bid!;
    expect(firstBid).toBeDefined();
    expect(firstBid.id).toBeDefined();

    const leaseMessage: MsgCreateLease = {
      bidId: {
        owner: firstBid.id!.owner,
        dseq: firstBid.id!.dseq,
        gseq: firstBid.id!.gseq,
        oseq: firstBid.id!.oseq,
        provider: firstBid.id!.provider,
        bseq: firstBid.id!.bseq
      } as BidID
    };

    const leaseResult = await sdk.akash.market.v1beta5.createLease(leaseMessage, {
      memo: "Test lease creation from bid - Akash Chain SDK"
    });

    expect(leaseResult).toBeDefined();
    
    const leaseQuery = await sdk.akash.market.v1beta5.getLeases({
      filters: {
        owner: deploymentId.owner,
        dseq: deploymentId.dseq,
        gseq: 1,
        oseq: 1,
        provider: firstBid.id!.provider,
      }
    });

    expect(leaseQuery?.leases).toBeDefined();
    expect(Array.isArray(leaseQuery?.leases)).toBe(true);
    expect(leaseQuery!.leases!.length).toBeGreaterThan(0);
    
    const createdLease = leaseQuery!.leases![0]!.lease!;
    expect(createdLease.id?.owner).toBe(deploymentId.owner);
    expect(createdLease.id?.dseq.toString()).toBe(deploymentId.dseq.toString());
    expect(createdLease.id?.provider).toBe(firstBid.id!.provider);
  });

  it("should query existing leases from the network", async () => {
    const sdk = await createTestSDK();

    const queryParams = {
      pagination: {
        limit: 10,
      },
    };

    const response = await sdk.akash.market.v1beta5.getLeases({
      filters: {
        owner: "",
        dseq: Long.UZERO,
        gseq: 0,
        oseq: 0,
        provider: "",
        state: "",
        bseq: 0
      },
      pagination: queryParams.pagination
    });
    
    expect(response?.leases).toBeDefined();
    expect(Array.isArray(response?.leases)).toBe(true);
    
    expect(response?.leases).toBeDefined();
    expect(response.leases.length).toBeGreaterThan(0);
    
    const lease = response.leases[0]?.lease;
    expect(lease?.id?.owner).toBeDefined();
    expect(lease?.id?.dseq).toBeDefined();
    expect(lease?.state).toBeDefined();
  });

  it("should query existing bids from the network", async () => {
    const sdk = await createTestSDK();

    const queryParams = {
      pagination: {
        limit: 10,
      },
    };

    const response = await sdk.akash.market.v1beta5.getBids({
      filters: {
        owner: "",
        dseq: Long.UZERO,
        gseq: 0,
        oseq: 0,
        provider: "",
        state: "",
        bseq: 0
      },
      pagination: queryParams.pagination
    });
    
    expect(response?.bids).toBeDefined();
    expect(Array.isArray(response?.bids)).toBe(true);

    expect(response?.bids).toBeDefined();
    expect(response.bids.length).toBeGreaterThan(0);
    
    const bid = response.bids[0]?.bid;
    expect(bid?.id?.owner).toBeDefined();
    expect(bid?.id?.dseq).toBeDefined();
    expect(bid?.state).toBeDefined();
  });
});
