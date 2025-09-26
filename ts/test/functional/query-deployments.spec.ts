/**
 * Functional tests for querying deployments using the Akash Chain SDK
 * 
 * These tests demonstrate how to query live deployment data from the Akash network.
 */

import { describe, expect, it } from "@jest/globals";
import Long from "long";
import { BinaryWriter } from "@bufbuild/protobuf/wire";

import { createChainNodeSDK } from "../../src/sdk/chain/server/index.ts";
import type { QueryDeploymentsResponse } from "../../src/generated/protos/akash/deployment/v1beta4/query.ts";
import { MsgCreateDeployment } from "../../src/generated/protos/akash/deployment/v1beta4/deploymentmsg.ts";
import { DeploymentID } from "../../src/generated/protos/akash/deployment/v1/deployment.ts";
import { GroupSpec } from "../../src/generated/protos/akash/deployment/v1beta4/groupspec.ts";
import { ResourceUnit } from "../../src/generated/protos/akash/deployment/v1beta4/resourceunit.ts";
import { Resources } from "../../src/generated/protos/akash/base/resources/v1beta4/resources.ts";
import { PlacementRequirements } from "../../src/generated/protos/akash/base/attributes/v1/attribute.ts";
import { Deposit } from "../../src/generated/protos/akash/base/deposit/v1/deposit.ts";
import { Coin, DecCoin } from "../../src/generated/protos/cosmos/base/v1beta1/coin.ts";

describe("Deployment Queries", () => {
  // Use the working configuration from your provided snippet
  // Query and TX endpoints are different!
  // Note: These are gRPC endpoints that need proper URL schemes
  const QUERY_RPC_URL = process.env.QUERY_RPC_URL || "http://rpc.dev.akash.pub:30090";
  const TX_RPC_URL = process.env.TX_RPC_URL || "https://testnetrpc.akashnet.net:443";
  const TEST_TIMEOUT = 15000;

  // Helper function to create SDK instance
  const createTestSDK = () => createChainNodeSDK({
    query: { baseUrl: QUERY_RPC_URL },
    tx: { baseUrl: TX_RPC_URL, signer: null as any },
  });

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
            storage: [],
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
    const expectedBase64 = "CjIKLWFrYXNoMXRlc3QxMjM0NTY3ODlhYmNkZWZnaGlqa2xtbm9wcXJzdHV2d3h5ehDSCRJFCgp0ZXN0LWdyb3VwEgIKABozCiEIARIHCgUKAzEwMBoNCgsKCTEzNDIxNzcyOCoFCgMKATAQARoMCgR1YWt0EgQxMDAwGgQBAgMEIhEKDwoEdWFrdBIHNTAwMDAwMA==";
    
    // Assert the serialization matches expected value
    expect(base64Encoded).toBe(expectedBase64);
    
    // Also verify we can decode it back
    const decoded = MsgCreateDeployment.decode(encoded);
    expect(decoded.id?.owner).toBe("akash1test123456789abcdefghijklmnopqrstuvwxyz");
    expect(decoded.id?.dseq.toNumber()).toBe(1234);
    expect(decoded.groups).toHaveLength(1);
    expect(decoded.groups[0]?.name).toBe("test-group");
  });
});
