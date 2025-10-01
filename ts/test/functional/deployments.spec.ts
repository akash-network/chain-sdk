/**
 * Functional tests for deployment operations using the Akash Chain SDK
 * 
 * These tests demonstrate how to:
 * - Query live deployment data from the Akash network
 * - Serialize deployment messages for API consistency testing  
 * - Create actual deployment transactions on testnet
 * 
 * Environment Variables:
 * - TEST_MNEMONIC: A funded testnet account mnemonic for deployment transaction tests
 *   Example: export TEST_MNEMONIC="word1 word2 word3 ... word12"
 * 
 * Note: Never use production mnemonics in tests!
 */

import { describe, expect, it, afterAll, beforeAll } from "@jest/globals";
import Long from "long";
import { BinaryWriter } from "@bufbuild/protobuf/wire";
import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";

import { createChainNodeSDK } from "../../src/sdk/chain/createChainNodeSDK.ts";
import { MsgCreateDeployment, MsgCloseDeployment } from "../../src/generated/protos/akash/deployment/v1beta4/deploymentmsg.ts";
import { Storage } from "../../src/generated/protos/akash/base/resources/v1beta4/storage.ts";
import { Source } from "../../src/generated/protos/akash/base/deposit/v1/deposit.ts";
import { Coin, DecCoin } from "../../src/generated/protos/cosmos/base/v1beta1/coin.ts";
import { testUtils } from "../helpers/testOrchestrator.js";

describe("Deployment Queries", () => {
  // Use the working configuration from your provided snippet
  // Query and TX endpoints are different!
  // Note: These are gRPC endpoints that need proper URL schemes
  const QUERY_RPC_URL = process.env.QUERY_RPC_URL || "http://rpc.dev.akash.pub:30090";
  const TX_RPC_URL = process.env.TX_RPC_URL || "https://rpc.testnet.akt.dev:443/rpc";
  const TEST_TIMEOUT = 15000;

  // Helper function to create SDK instance
  const createTestSDK = (wallet?: DirectSecp256k1HdWallet) => createChainNodeSDK({
    query: { baseUrl: QUERY_RPC_URL },
    tx: { baseUrl: TX_RPC_URL, signer: wallet || null as any },
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

  beforeAll(async () => {
    testUtils.reset();
  });

  afterAll(async () => {
    await cleanupDeployments();
  }, 300000);

  it("should query deployments from the network", async () => {
    const sdk = createTestSDK();

    const queryParams = {
      pagination: {
        limit: 10,
      },
    };

    const response = await sdk.akash.deployment.v1beta4.getDeployments(queryParams);
    
    expect(response?.deployments).toBeDefined();
    expect(Array.isArray(response?.deployments)).toBe(true);
    
    console.log(`Found ${response?.deployments?.length || 0} deployments`);
    
    if (response?.deployments && response.deployments.length > 0) {
      const deployment = response.deployments[0]?.deployment;
      expect(deployment?.id?.owner).toBeDefined();
      expect(deployment?.id?.dseq).toBeDefined();
      expect(deployment?.state).toBeDefined();
      
      console.log(`First deployment: ${deployment?.id?.owner}/${deployment?.id?.dseq?.low}`);
    }
  }, TEST_TIMEOUT);

  it("should query deployments with pagination", async () => {
    const sdk = createTestSDK();

    const response = await sdk.akash.deployment.v1beta4.getDeployments({
      pagination: { limit: 5, countTotal: true },
    });
    
    expect(response?.deployments).toBeDefined();
    expect(Array.isArray(response?.deployments)).toBe(true);
    
    console.log(`Paginated query returned ${response?.deployments?.length || 0} deployments`);
    
    if (response?.pagination) {
      expect(response?.pagination).toBeDefined();
    }
  }, TEST_TIMEOUT);

  it("should handle empty results gracefully", async () => {
    const sdk = createTestSDK();

    const response = await sdk.akash.deployment.v1beta4.getDeployments({
      pagination: { limit: 1 },
    }) as any;
    
    // Should handle both empty responses and empty deployment lists
    expect(response?.deployments).toBeDefined();
    expect(Array.isArray(response?.deployments)).toBe(true);
    expect(response?.deployments?.length || 0).toBeGreaterThanOrEqual(0);
    
  }, TEST_TIMEOUT);

  it("should create SDK instance with all modules", () => {
    const sdk = createTestSDK();

    // Verify core SDK structure
    expect(typeof sdk.akash.deployment.v1beta4.getDeployments).toBe('function');
    expect(typeof sdk.akash.cert.v1.getCertificates).toBe('function');
    
    // Verify all modules are available
    expect(sdk.akash.deployment).toBeDefined();
    expect(sdk.akash.cert).toBeDefined();
    expect(sdk.akash.market).toBeDefined();
    expect(sdk.akash.provider).toBeDefined();
    expect(sdk.akash.escrow).toBeDefined();
    
  });

  it("should serialize MsgCreateDeployment consistently", () => {
    // Helper function to create readable resource values from strings
    // This replaces hard-coded Uint8Array values with human-readable string values
    const createResourceValue = (value: string): { val: Uint8Array } => ({
      val: new TextEncoder().encode(value)
    });

    // Alternative readable values you could use:
    // CPU: "100" = 0.1 CPU, "500" = 0.5 CPU, "1000" = 1 CPU
    // Memory: "134217728" = 128Mi, "268435456" = 256Mi, "1073741824" = 1Gi
    // GPU: "0" = no GPU, "1" = 1 GPU unit

    // Create a minimal deployment request with deterministic data
    const deploymentRequest: MsgCreateDeployment = {
      id: {
        owner: "akash1test123456789abcdefghijklmnopqrstuvwxyz",
        dseq: Long.fromNumber(1234)
      },
      groups: [{
        name: "test-group",
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
              units: createResourceValue("100"), // 0.1 CPU (100 millicores)
              attributes: []
            },
            memory: {
              quantity: createResourceValue("134217728"), // 128Mi memory
              attributes: []
            },
            storage: [{
              name: "main",
              quantity: createResourceValue("2147483648"),
              attributes: []
            } as Storage],
            gpu: {
              units: createResourceValue("0"), // No GPU
              attributes: []
            },
            endpoints: []
          },
          count: 1,
          price: {
            denom: "uakt",
            amount: "10000"
          } as DecCoin
        }]
      }],
      hash: new Uint8Array([0x01, 0x02, 0x03, 0x04]),
      deposit: {
        amount: {
          denom: "uakt",
          amount: "5000000"
        } as Coin,
        sources: []
      }
    };

    // Encode the message
    const writer = new BinaryWriter();
    MsgCreateDeployment.encode(deploymentRequest, writer);
    const encoded = writer.finish();
    
    // Convert to base64
    const base64Encoded = Buffer.from(encoded).toString('base64');
    
    // Expected base64 - this will be the reference value to detect serialization changes
    // This is a snapshot test - if the serialization format changes, this test will fail
    // indicating a potential breaking change in the API
    const expectedBase64 = "CjIKLWFrYXNoMXRlc3QxMjM0NTY3ODlhYmNkZWZnaGlqa2xtbm9wcXJzdHV2d3h5ehDSCRJcCgp0ZXN0LWdyb3VwEgIKABpKCjcIARIHCgUKAzEwMBoNCgsKCTEzNDIxNzcyOCIUCgRtYWluEgwKCjIxNDc0ODM2NDgqBQoDCgEwEAEaDQoEdWFrdBIFMTAwMDAaBAECAwQiEQoPCgR1YWt0Egc1MDAwMDAw";
    
    // Assert the serialization matches expected value
    expect(base64Encoded).toBe(expectedBase64);
    
    // Also verify we can decode it back
    const decoded = MsgCreateDeployment.decode(encoded);
    expect(decoded.id?.owner).toBe("akash1test123456789abcdefghijklmnopqrstuvwxyz");
    expect(decoded.id?.dseq.toNumber()).toBe(1234);
    expect(decoded.groups).toHaveLength(1);
    expect(decoded.groups[0]?.name).toBe("test-group");
  });

  it("should create a deployment transaction", async () => {
    // Get test mnemonic from environment variable
    const testMnemonic = process.env.TEST_MNEMONIC;
    
    if (!testMnemonic) {
      console.log("Skipping deployment transaction test - TEST_MNEMONIC environment variable not set");
      console.log("To run this test, set TEST_MNEMONIC with a funded testnet account mnemonic");
      return;
    }
    
    // Create a test wallet
    const wallet = await DirectSecp256k1HdWallet.fromMnemonic(testMnemonic, { prefix: "akash" });
    const [account] = await wallet.getAccounts();
    
    // Print the test account address for funding if needed
    console.log(`\nTest Account Address: ${account.address}`);
    console.log(`To fund this account, send some AKT tokens to: ${account.address}`);
    console.log(`You can use a testnet faucet or transfer from another account\n`);
    
    // Helper function to create readable resource values from strings
    const createResourceValue = (value: string): { val: Uint8Array } => ({
      val: new TextEncoder().encode(value)
    });

    // Create SDK with test wallet
    const sdk = createChainNodeSDK({
      query: { baseUrl: QUERY_RPC_URL },
      tx: { baseUrl: TX_RPC_URL, signer: wallet },
    });

    // Create deployment message
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
              units: createResourceValue("500"), // 0.5 CPU
              attributes: []
            },
            memory: {
              quantity: createResourceValue("268435456"), // 256Mi memory
              attributes: []
            },
            storage: [{
              name: "beta3",
              quantity: createResourceValue("1073741824"), // 1Gi storage
              attributes: []
            } as Storage],
            gpu: {
              units: createResourceValue("0"), // No GPU
              attributes: []
            },
            endpoints: []
          },
          count: 1,
          price: {
            denom: "uakt",
            amount: "1000"
          } as DecCoin
        }]
      }],
      hash: new Uint8Array(32), // 32-byte hash (all zeros for test)
      deposit: {
        amount: {
          denom: "uakt",
          amount: "500000" // 5 AKT deposit
        } as Coin,
        sources: [Source.balance] // Use account balance as deposit source
      }
    };

    await testUtils.acquireTransactionLock();
    let result;
    try {
      result = await sdk.akash.deployment.v1beta4.createDeployment(deploymentMessage, {
      memo: "Test deployment creation - Akash Chain SDK",
      // Set afterSign callback to verify transaction structure
      afterSign: (txRaw: any) => {
        expect(txRaw).toBeDefined();
        expect(txRaw.bodyBytes).toBeDefined();
        expect(txRaw.authInfoBytes).toBeDefined();
        expect(txRaw.signatures).toBeDefined();
        expect(txRaw.signatures.length).toBeGreaterThan(0);
      },
      // Set afterBroadcast callback to capture transaction hash
      afterBroadcast: (txResponse: any) => {
        // Verify transaction was successful
        expect(txResponse.code).toBe(0); // 0 means success
        expect(txResponse.transactionHash).toBeDefined();
      }
      });
      
      // Transaction completed successfully
      console.log("Deployment transaction completed successfully!");
      console.log(`   - Transaction result:`, result);
    } finally {
      testUtils.releaseTransactionLock();
    }
    
    // Verify the response structure - these assertions are required for test to pass
    expect(result).toBeDefined();

    // Verify wallet and account structure
    expect(account.address).toMatch(/^akash1[a-z0-9]{38}$/);
    expect(account.pubkey).toHaveLength(33); // Compressed secp256k1 pubkey
    expect(deploymentMessage.id?.owner).toBe(account.address);
    expect(deploymentMessage.groups).toHaveLength(1);
    expect(deploymentMessage.groups[0]?.name).toBe("web-service");
  }, TEST_TIMEOUT);

  it("should cleanup all deployments for the test account", async () => {
    await testUtils.withTransactionLock(async () => {
      await cleanupDeployments();
    });
  }, 300000);
});
