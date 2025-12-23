/**
 * Tests for deployment operations against Go mock server
 *
 * These tests validate:
 * - TypeScript-Go protobuf encoding/decoding compatibility
 * - Transaction validation using Go ValidateBasic() logic
 * - Query endpoints (deployments, bids, leases)
 * - Message serialization consistency
 */

import { BinaryWriter } from "@bufbuild/protobuf/wire";
import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { afterAll, beforeAll, describe, expect, it } from "@jest/globals";
import Long from "long";
import path from "path";

import { Source } from "../../src/generated/protos/akash/base/deposit/v1/deposit.ts";
import { MsgCreateDeployment } from "../../src/generated/protos/akash/deployment/v1beta4/deploymentmsg.ts";
import { createChainNodeWebSDK } from "../../src/sdk/chain/createChainNodeWebSDK.ts";
import { getMessageType } from "../../src/sdk/getMessageType.ts";
import { startMockServer } from "../util/mockServer.ts";
import { createGatewayTxClient } from "../util/createGatewayTxClient.ts";

declare const jest: {
  setTimeout: (timeout: number) => void;
};

const TEST_MNEMONIC = "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about";

const createTestWallet = async () => 
  DirectSecp256k1HdWallet.fromMnemonic(TEST_MNEMONIC, { prefix: "akash" });

const createBaseResourceGroup = () => ({
  name: "test-group",
  requirements: {
    signedBy: { allOf: [], anyOf: [] },
    attributes: [],
  },
  resources: [{
    resource: {
      id: 1,
      cpu: { units: { val: new TextEncoder().encode("100") }, attributes: [] },
      memory: { quantity: { val: new TextEncoder().encode("134217728") }, attributes: [] },
      storage: [],
      gpu: { units: { val: new TextEncoder().encode("0") }, attributes: [] },
      endpoints: [],
    },
    count: 1,
    price: { denom: "uakt", amount: "1000" },
  }],
});

const createInvalidDeployment = (
  owner: string,
  dseq: number,
  overrides: Partial<MsgCreateDeployment> = {}
): MsgCreateDeployment => ({
  id: { owner, dseq: Long.fromNumber(dseq) },
  groups: [createBaseResourceGroup()],
  hash: new Uint8Array(32),
  deposit: {
    amount: { denom: "uakt", amount: "1000" },
    sources: [Source.balance],
  },
  ...overrides,
});

describe("Deployment Queries", () => {
  jest.setTimeout(180000);

  let mockServer: Awaited<ReturnType<typeof startMockServer>>;

  beforeAll(async () => {
    const dataDir = path.resolve(__dirname, "../../../go/testutil/mock/data");
    mockServer = await startMockServer(dataDir);
  }, 180000);

  afterAll(async () => {
    await mockServer.stop();
  }, 3000);

  const createTestSDK = (wallet?: DirectSecp256k1HdWallet) => {
    let txClient;

    if (wallet) {
      txClient = createGatewayTxClient({
        gatewayUrl: mockServer.gatewayUrl,
        signer: wallet,
        getMessageType,
      });
    }

    return createChainNodeWebSDK({
      query: { baseUrl: mockServer.gatewayUrl },
      tx: txClient ? { signer: txClient } : undefined,
    });
  };


  it("should query deployments from the network and decode Go-encoded responses", async () => {
    const sdk = createTestSDK();

    const queryParams = {
      pagination: {
        limit: 10,
      },
    };

    const response = await sdk.akash.deployment.v1beta4.getDeployments(queryParams);

    expect(response).toBeDefined();
    expect(response?.deployments).toBeDefined();
    expect(Array.isArray(response?.deployments)).toBe(true);
    expect(response.deployments.length).toBeGreaterThanOrEqual(0);
  });

  it("should query deployments with pagination", async () => {
    const sdk = createTestSDK();

    const response = await sdk.akash.deployment.v1beta4.getDeployments({
      pagination: { limit: 5, countTotal: true },
    });

    expect(response).toBeDefined();
    expect(response?.deployments).toBeDefined();
    expect(Array.isArray(response?.deployments)).toBe(true);

    if (response?.pagination) {
      expect(response.pagination).toHaveProperty("total");
    }
  });

  it("should create SDK instance with all modules", async () => {
    const sdk = createTestSDK();

    // Verify core SDK structure
    expect(typeof sdk.akash.deployment.v1beta4.getDeployments).toBe("function");
    expect(typeof sdk.akash.cert.v1.getCertificates).toBe("function");

    // Verify all modules are available
    expect(sdk.akash.deployment).toBeDefined();
    expect(sdk.akash.cert).toBeDefined();
    expect(sdk.akash.market).toBeDefined();
    expect(sdk.akash.provider).toBeDefined();
    expect(sdk.akash.escrow).toBeDefined();
  });

  it("should serialize MsgCreateDeployment consistently", () => {
    const createResourceValue = (value: string) => ({
      val: new TextEncoder().encode(value),
    });

    const deploymentRequest: MsgCreateDeployment = {
      id: {
        owner: "akash1test123456789abcdefghijklmnopqrstuvwxyz",
        dseq: Long.fromNumber(1234),
      },
      groups: [{
        name: "test-group",
        requirements: {
          signedBy: { allOf: [], anyOf: [] },
          attributes: [],
        },
        resources: [{
          resource: {
            id: 1,
            cpu: { units: createResourceValue("100"), attributes: [] },
            memory: { quantity: createResourceValue("134217728"), attributes: [] },
            storage: [{
              name: "main",
              quantity: createResourceValue("2147483648"),
              attributes: [],
            }],
            gpu: { units: createResourceValue("0"), attributes: [] },
            endpoints: [],
          },
          count: 1,
          price: { denom: "uakt", amount: "10000" },
        }],
      }],
      hash: new Uint8Array([0x01, 0x02, 0x03, 0x04]),
      deposit: {
        amount: { denom: "uakt", amount: "5000000" },
        sources: [],
      },
    };

    const writer = new BinaryWriter();
    MsgCreateDeployment.encode(deploymentRequest, writer);
    const encoded = writer.finish();
    const base64Encoded = Buffer.from(encoded).toString("base64");

    const expectedBase64 = "CjIKLWFrYXNoMXRlc3QxMjM0NTY3ODlhYmNkZWZnaGlqa2xtbm9wcXJzdHV2d3h5ehDSCRJcCgp0ZXN0LWdyb3VwEgIKABpKCjcIARIHCgUKAzEwMBoNCgsKCTEzNDIxNzcyOCIUCgRtYWluEgwKCjIxNDc0ODM2NDgqBQoDCgEwEAEaDQoEdWFrdBIFMTAwMDAaBAECAwQiEQoPCgR1YWt0Egc1MDAwMDAw";
    expect(base64Encoded).toBe(expectedBase64);

    const decoded = MsgCreateDeployment.decode(encoded);
    expect(decoded.id?.owner).toBe("akash1test123456789abcdefghijklmnopqrstuvwxyz");
    expect(decoded.id?.dseq.toNumber()).toBe(1234);
    expect(decoded.groups).toHaveLength(1);
    expect(decoded.groups[0]?.name).toBe("test-group");
  });


  describe("Transaction validation", () => {
    it("should reject deployment with empty groups", async () => {
      const wallet = await createTestWallet();
      const [account] = await wallet.getAccounts();
      const sdk = createTestSDK(wallet);

      const invalidDeployment = createInvalidDeployment(account.address, 999999, {
        groups: [],
      });

      await expect(
        sdk.akash.deployment.v1beta4.createDeployment(invalidDeployment, {
          memo: "Test invalid deployment",
        }),
      ).rejects.toThrow(/Invalid groups/i);
    });

    it("should reject deployment with empty hash", async () => {
      const wallet = await createTestWallet();
      const [account] = await wallet.getAccounts();
      const sdk = createTestSDK(wallet);

      const invalidDeployment = createInvalidDeployment(account.address, 999998, {
        hash: new Uint8Array(0),
      });

      await expect(
        sdk.akash.deployment.v1beta4.createDeployment(invalidDeployment, {
          memo: "Test invalid hash",
        }),
      ).rejects.toThrow(/Invalid: empty hash/i);
    });

    it("should reject deployment with invalid hash length", async () => {
      const wallet = await createTestWallet();
      const [account] = await wallet.getAccounts();
      const sdk = createTestSDK(wallet);

      const invalidDeployment = createInvalidDeployment(account.address, 999997, {
        hash: new Uint8Array(16),
      });

      await expect(
        sdk.akash.deployment.v1beta4.createDeployment(invalidDeployment, {
          memo: "Test invalid hash length",
        }),
      ).rejects.toThrow(/Invalid: deployment hash/i);
    });

    it("should reject deployment with negative price", async () => {
      const wallet = await createTestWallet();
      const [account] = await wallet.getAccounts();
      const sdk = createTestSDK(wallet);

      const baseGroup = createBaseResourceGroup();
      const groupWithNegativePrice = {
        ...baseGroup,
        resources: [{
          ...baseGroup.resources[0],
          resource: {
            ...baseGroup.resources[0].resource,
            storage: [{
              name: "main",
              quantity: { val: new TextEncoder().encode("1073741824") },
              attributes: [],
            }],
          },
          price: { denom: "uakt", amount: "-1" },
        }],
      };

      const invalidDeployment = createInvalidDeployment(account.address, 999996, {
        groups: [groupWithNegativePrice],
      });

      await expect(
        sdk.akash.deployment.v1beta4.createDeployment(invalidDeployment, {
          memo: "Test negative price",
        }),
      ).rejects.toThrow(/invalid price object/i);
    });
  });
});
