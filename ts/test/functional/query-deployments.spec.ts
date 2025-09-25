/**
 * Functional tests for querying deployments using the Akash Chain SDK
 * 
 * These tests demonstrate how to query live deployment data from the Akash network.
 */

import { describe, expect, it } from "@jest/globals";
import Long from "long";

import { createChainNodeSDK } from "../../src/sdk/chain/server/index.ts";
import type { QueryDeploymentsResponse } from "../../src/generated/protos/akash/deployment/v1beta4/query.ts";

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

  // Helper function to create pagination config
  const createPagination = (limit: number) => ({
    key: new Uint8Array(0),
    offset: Long.UZERO,
    limit: Long.fromNumber(limit),
    countTotal: false,
    reverse: false,
  });



  it("should query deployments from the network", async () => {
    const sdk = createTestSDK();

    const queryParams = {
      pagination: createPagination(10),
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
      pagination: { ...createPagination(5), countTotal: true },
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
      pagination: createPagination(1),
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
});
