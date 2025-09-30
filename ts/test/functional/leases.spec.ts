
import { describe, expect, it, afterAll } from "@jest/globals";
import Long from "long";
import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";

import { createChainNodeSDK } from "../../src/sdk/chain/server/index.ts";
import { MsgCreateDeployment } from "../../src/generated/protos/akash/deployment/v1beta4/deploymentmsg.ts";
import { MsgCreateLease } from "../../src/generated/protos/akash/market/v1beta5/leasemsg.ts";
import { MsgCloseDeployment } from "../../src/generated/protos/akash/deployment/v1beta4/deploymentmsg.ts";
import { BidID } from "../../src/generated/protos/akash/market/v1/bid.ts";
import { Storage } from "../../src/generated/protos/akash/base/resources/v1beta4/storage.ts";
import { Source } from "../../src/generated/protos/akash/base/deposit/v1/deposit.ts";
import { Coin, DecCoin } from "../../src/generated/protos/cosmos/base/v1beta1/coin.ts";

describe("Lease Operations", () => {
  const QUERY_RPC_URL = process.env.QUERY_RPC_URL || "http://rpc.dev.akash.pub:30090";
  const TX_RPC_URL = process.env.TX_RPC_URL || "https://rpc.testnet.akt.dev:443/rpc";
  const TEST_TIMEOUT = 60000;

  const createTestSDK = (wallet?: DirectSecp256k1HdWallet) => createChainNodeSDK({
    query: { baseUrl: QUERY_RPC_URL },
    tx: { baseUrl: TX_RPC_URL, signer: wallet || null as any },
  });

  const createResourceValue = (value: string): { val: Uint8Array } => ({
    val: new TextEncoder().encode(value)
  });

  const wait = (ms: number) => new Promise(resolve => setTimeout(resolve, ms));

  const cleanupDeployments = async () => {
    const testMnemonic = process.env.TEST_MNEMONIC;
    
    if (!testMnemonic) {
      console.log("Skipping deployment cleanup - TEST_MNEMONIC not set");
      return;
    }

    try {
      const wallet = await DirectSecp256k1HdWallet.fromMnemonic(testMnemonic, { prefix: "akash" });
      const [account] = await wallet.getAccounts();
      const sdk = createTestSDK(wallet);

      console.log(`\nCleaning up deployments for account: ${account.address}`);

      const deploymentsResponse = await sdk.akash.deployment.v1beta4.getDeployments({
        filters: {
          owner: account.address,
          state: "active",
          dseq: Long.UZERO
        },
        pagination: { limit: 100 }
      });

      if (!deploymentsResponse?.deployments || deploymentsResponse.deployments.length === 0) {
        console.log("No deployments found to clean up");
        return;
      }

      console.log(`Found ${deploymentsResponse.deployments.length} open deployments to clean up`);

      for (const deploymentResponse of deploymentsResponse.deployments) {
        const deployment = deploymentResponse.deployment;
        if (!deployment?.id) continue;

        console.log(`Processing deployment ${deployment.id.dseq} (state: ${deployment.state})`);

        try {
          const closeMessage: MsgCloseDeployment = {
            id: {
              owner: deployment.id.owner,
              dseq: deployment.id.dseq
            }
          };

          console.log(`Closing deployment ${deployment.id.owner}/${deployment.id.dseq}`);
          
          await sdk.akash.deployment.v1beta4.closeDeployment(closeMessage, {
            memo: "Test cleanup - closing deployment"
          });

          console.log(`Successfully closed deployment ${deployment.id.dseq}`);
          
          console.log("Waiting 6 seconds before next closure...");
          await wait(6000);
        } catch (error) {
          const errorMessage = error instanceof Error ? error.message : String(error);
          if (errorMessage.includes("Deployment closed") || errorMessage.includes("already closed")) {
            console.log(`Deployment ${deployment.id.dseq} is already closed, skipping`);
          } else {
            console.log(`Failed to close deployment ${deployment.id.dseq}:`, errorMessage);
          }
        }
      }

      console.log("Deployment cleanup completed");
    } catch (error) {
      console.log("Error during deployment cleanup:", error);
    }
  };

  // afterAll(async () => {
  //   await cleanupDeployments();
  // }, 120000);

  it("should create a deployment, wait for bids, select first bid and create a lease", async () => {
    const testMnemonic = process.env.TEST_MNEMONIC;
    
    if (!testMnemonic) {
      console.log("Skipping lease creation test - TEST_MNEMONIC environment variable not set");
      console.log("To run this test, set TEST_MNEMONIC with a funded testnet account mnemonic");
      return;
    }
    
    const wallet = await DirectSecp256k1HdWallet.fromMnemonic(testMnemonic, { prefix: "akash" });
    const [account] = await wallet.getAccounts();
    
    console.log(`Test Account Address: ${account.address}`);
    console.log(`To fund this account, send some AKT tokens to: ${account.address}`);
    
    const sdk = createTestSDK(wallet);

    console.log("Step 1: Creating deployment...");
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
    
    console.log("Deployment created successfully!");
    expect(deploymentResult).toBeDefined();
    console.log(deploymentResult);

    const deploymentId = {
      owner: account.address,
      dseq: deploymentMessage.id!.dseq
    };

    console.log("Step 2: Waiting for providers to create bids...");
    console.log(`Deployment ID: ${deploymentId.owner}/${deploymentId.dseq}`);
    let bidsResponse;
    let attempts = 0;
    const maxAttempts = 18;
    
    do {
      await wait(10000);
      attempts++;
      
      console.log(`Checking for bids (attempt ${attempts}/${maxAttempts})...`);
      console.log("Make sure your address is whitelisted on this network.");
      
      bidsResponse = await sdk.akash.market.v1beta5.getBids({
        filters: {
          owner: deploymentId.owner,
          dseq: deploymentId.dseq,
          gseq: 1,
          oseq: 1,
        }
      });
      
      console.log(`Found ${bidsResponse?.bids?.length || 0} bids`);
      
    } while ((!bidsResponse?.bids || bidsResponse.bids.length < 2) && attempts < maxAttempts);


    expect(bidsResponse?.bids).toBeDefined();
    expect(Array.isArray(bidsResponse?.bids)).toBe(true);
    
    if (bidsResponse!.bids!.length >= 2) {
      console.log(`Found ${bidsResponse!.bids!.length} bids for the deployment`);
    } else if (bidsResponse!.bids!.length === 1) {
      console.log(`Found only 1 bid, proceeding with single bid test`);
    } else {
      throw new Error(`No bids found after ${maxAttempts} attempts. Check deployment resources and pricing.`);
    }
    
    expect(bidsResponse!.bids!.length).toBeGreaterThan(0);
    
    bidsResponse!.bids!.forEach((bidResponse, index) => {
      const bid = bidResponse.bid;
      console.log(`  Bid ${index + 1}: Provider ${bid?.id?.provider}, Price: ${bid?.price?.amount}${bid?.price?.denom}`);
    });

    console.log("Step 4: Selecting the first bid...");
    const firstBid = bidsResponse!.bids![0]!.bid!;
    expect(firstBid).toBeDefined();
    expect(firstBid.id).toBeDefined();
    
    console.log(`Selected bid from provider: ${firstBid.id!.provider}`);

    console.log("Step 5: Creating lease from selected bid...");
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

    console.log("Step 6: Verifying lease creation...");
    expect(leaseResult).toBeDefined();
    console.log("Lease created successfully!");
    
    const leaseQuery = await sdk.akash.market.v1beta5.getLeases({
      filters: {
        owner: deploymentId.owner,
        dseq: deploymentId.dseq,
        gseq: 1,
        oseq: 1,
        provider: firstBid.id!.provider,
        state: "",
        bseq: 0
      }
    });

    expect(leaseQuery?.leases).toBeDefined();
    expect(Array.isArray(leaseQuery?.leases)).toBe(true);
    expect(leaseQuery!.leases!.length).toBeGreaterThan(0);
    
    const createdLease = leaseQuery!.leases![0]!.lease!;
    expect(createdLease.id?.owner).toBe(deploymentId.owner);
    expect(createdLease.id?.dseq.toString()).toBe(deploymentId.dseq.toString());
    expect(createdLease.id?.provider).toBe(firstBid.id!.provider);
    
    console.log("Lease verification completed successfully!");
    console.log(`Lease ID: ${createdLease.id?.owner}/${createdLease.id?.dseq}/${createdLease.id?.gseq}/${createdLease.id?.oseq}/${createdLease.id?.provider}`);
    console.log(`Lease State: ${createdLease.state}`);
    console.log(`Lease Price: ${createdLease.price?.amount}${createdLease.price?.denom}`);

  }, TEST_TIMEOUT);

  it("should query existing leases from the network", async () => {
    const sdk = createTestSDK();

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
    
    console.log(`Found ${response?.leases?.length || 0} leases`);
    
    if (response?.leases && response.leases.length > 0) {
      const lease = response.leases[0]?.lease;
      expect(lease?.id?.owner).toBeDefined();
      expect(lease?.id?.dseq).toBeDefined();
      expect(lease?.state).toBeDefined();
      
      console.log(`First lease: ${lease?.id?.owner}/${lease?.id?.dseq?.low} State: ${lease?.state}`);
    }
  }, 15000);

  it("should query existing bids from the network", async () => {
    const sdk = createTestSDK();

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
    
    console.log(`Found ${response?.bids?.length || 0} bids`);
    
    if (response?.bids && response.bids.length > 0) {
      const bid = response.bids[0]?.bid;
      expect(bid?.id?.owner).toBeDefined();
      expect(bid?.id?.dseq).toBeDefined();
      expect(bid?.state).toBeDefined();
      
      console.log(`First bid: ${bid?.id?.owner}/${bid?.id?.dseq?.low} Provider: ${bid?.id?.provider}, Price: ${bid?.price?.amount}${bid?.price?.denom}`);
    }
  }, 15000);

  it("should cleanup all deployments for the test account", async () => {
    await cleanupDeployments();
  }, 300000);
});
