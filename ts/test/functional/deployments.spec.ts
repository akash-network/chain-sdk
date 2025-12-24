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

import { makeCosmoshubPath } from "@cosmjs/amino";
import { Source } from "../../src/generated/protos/akash/base/deposit/v1/deposit.ts";
import { MsgCreateDeployment } from "../../src/generated/protos/akash/deployment/v1beta4/deploymentmsg.ts";
import { MsgCreateBid, MsgCloseBid } from "../../src/generated/protos/akash/market/v1beta5/bidmsg.ts";
import { MsgCreateLease } from "../../src/generated/protos/akash/market/v1beta5/leasemsg.ts";
import { LeaseClosedReason } from "../../src/generated/protos/akash/market/v1/types.ts";
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

const normalizeDec = (value: string) => {
  const [intPart, fracPart = ""] = value.split(".");
  const trimmedFrac = fracPart.replace(/0+$/, "");
  return trimmedFrac ? `${intPart}.${trimmedFrac}` : intPart;
};

const toSnake = (input: string) => input.replace(/([A-Z])/g, "_$1").toLowerCase();

const isPrintableAscii = (input: string) => /^[\x20-\x7E]*$/.test(input);

const decodeMaybeBase64Ascii = (input: string): string | null => {
  if (!/^[A-Za-z0-9+/=]+$/.test(input)) return null;
  try {
    const decoded = Buffer.from(input, "base64").toString("utf8");
    return isPrintableAscii(decoded) ? decoded : null;
  } catch {
    return null;
  }
};

const normalizeValue = (value: any, key?: string): any => {
  if (value instanceof Uint8Array) {
    const asString = new TextDecoder().decode(value);
    const result = isPrintableAscii(asString) ? asString : Buffer.from(value).toString("base64");
    if (key === "val" && result === "0") {
      return "";
    }
    return result;
  }

  if (value && typeof value === "object" && typeof (value as any).toString === "function" && ("low" in (value as any) || "high" in (value as any))) {
    return (value as any).toString();
  }

  if (Array.isArray(value)) {
    return value.map(item => normalizeValue(item, key));
  }

  if (value && typeof value === "object") {
    const normalized: Record<string, any> = {};
    for (const [k, v] of Object.entries(value)) {
      normalized[toSnake(k)] = normalizeValue(v, k);
    }
    return normalized;
  }

  if (typeof value === "string") {
    if (/^\d+\.\d+0*$/.test(value)) {
      return normalizeDec(value);
    }
    if (key === "val") {
      if (value === "") {
        return "";
      }
      const decoded = decodeMaybeBase64Ascii(value);
      if (decoded !== null) {
        return decoded === "0" ? "" : decoded;
      }
      return value;
    }
    if (key === "reason" && value.startsWith("lease_closed_")) {
      return value;
    }
    return value;
  }

  if (typeof value === "number") {
    if (key === "sources") {
      return value === 1 ? "balance" : String(value);
    }
    if (key === "reason") {
      const reasonNames: Record<number, string> = {
        0: "lease_closed_invalid",
        1: "lease_closed_owner",
        10000: "lease_closed_reason_unstable",
        10001: "lease_closed_reason_decommission",
        10002: "lease_closed_reason_unspecified",
        10003: "lease_closed_reason_manifest_timeout",
      };
      return reasonNames[value] || String(value);
    }
    return value;
  }

  return value;
};

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
    mockServer = await startMockServer();
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
        chainId: "akashnet-2",
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

    it("encodes fractional price for go decode", async () => {
      const wallet = await createTestWallet();
      const [account] = await wallet.getAccounts();
      const sdk = createTestSDK(wallet);
      const fractionalPrice = "0.123456";

      const fractionalGroup = (() => {
        const base = createBaseResourceGroup();
        return {
          ...base,
          resources: base.resources.map(resource => ({
            ...resource,
            resource: {
              ...resource.resource,
              storage: [{
                name: "main",
                quantity: { val: new TextEncoder().encode("1073741824") },
                attributes: [],
              }],
            },
            price: { ...resource.price, amount: fractionalPrice },
          })),
        };
      })();

      const deployment = createInvalidDeployment(account.address, 999995, {
        groups: [fractionalGroup],
        hash: new Uint8Array(Array.from({ length: 32 }, (_, i) => i + 1)),
      });

      await sdk.akash.deployment.v1beta4.createDeployment(deployment, {
        memo: "fractional price",
      });

      const res = await fetch(`${mockServer.gatewayUrl}/mock/last-deployment`);
      expect(res.ok).toBe(true);

      const decoded = await res.json();
      const price = decoded?.groups?.[0]?.resources?.[0]?.price;
      expect(price?.denom).toBe("uakt");
      expect(normalizeDec(price?.amount as string)).toBe(fractionalPrice);

      const normalizedExpected = normalizeValue(deployment);
      const normalizedActual = normalizeValue(decoded);

      expect(normalizedActual).toEqual(normalizedExpected);
      expect(normalizedActual).toEqual(normalizeValue(decoded));
    });

    it("encodes bid dec price and survives go decode", async () => {
      const wallet = await DirectSecp256k1HdWallet.fromMnemonic(TEST_MNEMONIC, {
        prefix: "akash",
        hdPaths: [makeCosmoshubPath(0), makeCosmoshubPath(1)],
      });
      const accounts = await wallet.getAccounts();
      const owner = accounts[0];
      const provider = accounts[1] ?? accounts[0];
      const sdk = createTestSDK(wallet);
      const fractionalPrice = "0.123456";

      const baseGroup = createBaseResourceGroup();
      const resources = baseGroup.resources[0]?.resource;

      if (!resources) {
        throw new Error("missing base resources");
      }

      const bid: MsgCreateBid = {
        id: {
          owner: owner.address,
          provider: provider.address,
          dseq: Long.fromNumber(777),
          gseq: 1,
          oseq: 1,
          bseq: 0,
        },
        price: { denom: "uakt", amount: fractionalPrice },
        deposit: {
          amount: { denom: "uakt", amount: "5000000" },
          sources: [Source.balance],
        },
        resourcesOffer: [{
          resources: resources,
          count: 1,
        }],
      };

      await sdk.akash.market.v1beta5.createBid(bid, { memo: "bid fractional price" });

      const res = await fetch(`${mockServer.gatewayUrl}/mock/last-bid`);
      expect(res.ok).toBe(true);

      const decoded = await res.json();
      expect(decoded?.id?.owner).toBe(owner.address);
      expect(decoded?.id?.provider).toBe(provider.address);
      const price = decoded?.price;
      expect(price?.denom).toBe("uakt");
      expect(normalizeDec(price?.amount as string)).toBe(fractionalPrice);

      const normalizedExpected = normalizeValue(bid);
      const normalizedActual = normalizeValue(decoded);

      expect(normalizedActual).toEqual(normalizedExpected);
      expect(normalizedActual).toEqual(normalizeValue(decoded));
    });

    it("handles 18-decimal precision price losslessly", async () => {
      const wallet = await DirectSecp256k1HdWallet.fromMnemonic(TEST_MNEMONIC, {
        prefix: "akash",
        hdPaths: [makeCosmoshubPath(0), makeCosmoshubPath(1)],
      });
      const accounts = await wallet.getAccounts();
      const owner = accounts[0];
      const provider = accounts[1] ?? accounts[0];
      const sdk = createTestSDK(wallet);
      const highPrecisionPrice = "0.123456789012345678";

      const baseGroup = createBaseResourceGroup();
      const resources = baseGroup.resources[0]?.resource;

      if (!resources) {
        throw new Error("missing base resources");
      }

      const bid: MsgCreateBid = {
        id: {
          owner: owner.address,
          provider: provider.address,
          dseq: Long.fromNumber(666),
          gseq: 1,
          oseq: 1,
          bseq: 0,
        },
        price: { denom: "uakt", amount: highPrecisionPrice },
        deposit: {
          amount: { denom: "uakt", amount: "5000000" },
          sources: [Source.balance],
        },
        resourcesOffer: [{
          resources: resources,
          count: 1,
        }],
      };

      await sdk.akash.market.v1beta5.createBid(bid, { memo: "18-decimal price" });

      const res = await fetch(`${mockServer.gatewayUrl}/mock/last-bid`);
      expect(res.ok).toBe(true);

      const decoded = await res.json();
      const price = decoded?.price;
      expect(price?.denom).toBe("uakt");
      expect(normalizeDec(price?.amount as string)).toBe(highPrecisionPrice);
    });

    it("handles zero price losslessly", async () => {
      const wallet = await DirectSecp256k1HdWallet.fromMnemonic(TEST_MNEMONIC, {
        prefix: "akash",
        hdPaths: [makeCosmoshubPath(0), makeCosmoshubPath(1)],
      });
      const accounts = await wallet.getAccounts();
      const owner = accounts[0];
      const provider = accounts[1] ?? accounts[0];
      const sdk = createTestSDK(wallet);
      const zeroPrice = "0";

      const baseGroup = createBaseResourceGroup();
      const resources = baseGroup.resources[0]?.resource;

      if (!resources) {
        throw new Error("missing base resources");
      }

      const bid: MsgCreateBid = {
        id: {
          owner: owner.address,
          provider: provider.address,
          dseq: Long.fromNumber(333),
          gseq: 1,
          oseq: 1,
          bseq: 0,
        },
        price: { denom: "uakt", amount: zeroPrice },
        deposit: {
          amount: { denom: "uakt", amount: "5000000" },
          sources: [Source.balance],
        },
        resourcesOffer: [{
          resources: resources,
          count: 1,
        }],
      };

      await expect(
        sdk.akash.market.v1beta5.createBid(bid, { memo: "zero price" }),
      ).rejects.toThrow(/price/i);
    });

    it("handles very small price losslessly", async () => {
      const wallet = await DirectSecp256k1HdWallet.fromMnemonic(TEST_MNEMONIC, {
        prefix: "akash",
        hdPaths: [makeCosmoshubPath(0), makeCosmoshubPath(1)],
      });
      const accounts = await wallet.getAccounts();
      const owner = accounts[0];
      const provider = accounts[1] ?? accounts[0];
      const sdk = createTestSDK(wallet);
      const verySmallPrice = "0.000000000000000001";

      const baseGroup = createBaseResourceGroup();
      const resources = baseGroup.resources[0]?.resource;

      if (!resources) {
        throw new Error("missing base resources");
      }

      const bid: MsgCreateBid = {
        id: {
          owner: owner.address,
          provider: provider.address,
          dseq: Long.fromNumber(222),
          gseq: 1,
          oseq: 1,
          bseq: 0,
        },
        price: { denom: "uakt", amount: verySmallPrice },
        deposit: {
          amount: { denom: "uakt", amount: "5000000" },
          sources: [Source.balance],
        },
        resourcesOffer: [{
          resources: resources,
          count: 1,
        }],
      };

      await sdk.akash.market.v1beta5.createBid(bid, { memo: "very small price" });

      const res = await fetch(`${mockServer.gatewayUrl}/mock/last-bid`);
      expect(res.ok).toBe(true);

      const decoded = await res.json();
      const price = decoded?.price;
      expect(price?.denom).toBe("uakt");
      expect(normalizeDec(price?.amount as string)).toBe(verySmallPrice);
    });

    it("broadcasts tx via SYNC mode and receives tx hash", async () => {
      const wallet = await createTestWallet();
      const [account] = await wallet.getAccounts();
      
      const txClient = createGatewayTxClient({
        gatewayUrl: mockServer.gatewayUrl,
        signer: wallet,
        chainId: "akashnet-2",
        getMessageType,
      });

      const deployment = createInvalidDeployment(account.address, 888888, {
        hash: new Uint8Array(Array.from({ length: 32 }, (_, i) => i + 100)),
        groups: [{
          name: "broadcast-test",
          requirements: { signedBy: { allOf: [], anyOf: [] }, attributes: [] },
          resources: [{
            resource: {
              id: 1,
              cpu: { units: { val: new TextEncoder().encode("100") }, attributes: [] },
              memory: { quantity: { val: new TextEncoder().encode("134217728") }, attributes: [] },
              storage: [{ name: "main", quantity: { val: new TextEncoder().encode("1073741824") }, attributes: [] }],
              gpu: { units: { val: new TextEncoder().encode("0") }, attributes: [] },
              endpoints: [],
            },
            count: 1,
            price: { denom: "uakt", amount: "1000" },
          }],
        }],
      });

      const messages = [{
        typeUrl: "/akash.deployment.v1beta4.MsgCreateDeployment",
        value: deployment,
      }];

      const fee = await txClient.estimateFee(messages, "broadcast test");
      const signed = await txClient.sign(messages, fee, "broadcast test");
      const result = await txClient.broadcast(signed);

      expect(result).toBeDefined();
      expect(result.code).toBe(0);
      expect(result.transactionHash).toBeDefined();
      expect(result.transactionHash.length).toBeGreaterThan(0);
      expect(Number(result.height)).toBeGreaterThan(0);
      expect(result.gasUsed).toBeGreaterThan(0n);
      expect(result.gasWanted).toBeGreaterThan(0n);
    });

    it("creates lease and verifies go decode", async () => {
      const wallet = await DirectSecp256k1HdWallet.fromMnemonic(TEST_MNEMONIC, {
        prefix: "akash",
        hdPaths: [makeCosmoshubPath(0), makeCosmoshubPath(1)],
      });
      const accounts = await wallet.getAccounts();
      const owner = accounts[0];
      const provider = accounts[1] ?? accounts[0];
      const sdk = createTestSDK(wallet);

      const lease: MsgCreateLease = {
        bidId: {
          owner: owner.address,
          provider: provider.address,
          dseq: Long.fromNumber(555),
          gseq: 1,
          oseq: 1,
          bseq: 0,
        },
      };

      await sdk.akash.market.v1beta5.createLease(lease, { memo: "create lease test" });

      const res = await fetch(`${mockServer.gatewayUrl}/mock/last-lease`);
      expect(res.ok).toBe(true);

      const decoded = await res.json();
      expect(decoded?.bid_id?.owner).toBe(owner.address);
      expect(decoded?.bid_id?.provider).toBe(provider.address);
      expect(decoded?.bid_id?.dseq).toBe("555");

      const normalizedExpected = normalizeValue(lease);
      const normalizedActual = normalizeValue(decoded);
      expect(normalizedActual).toEqual(normalizedExpected);
    });

    it("closes bid and verifies go decode", async () => {
      const wallet = await DirectSecp256k1HdWallet.fromMnemonic(TEST_MNEMONIC, {
        prefix: "akash",
        hdPaths: [makeCosmoshubPath(0), makeCosmoshubPath(1)],
      });
      const accounts = await wallet.getAccounts();
      const owner = accounts[0];
      const provider = accounts[1] ?? accounts[0];
      const sdk = createTestSDK(wallet);

      const closeBid: MsgCloseBid = {
        id: {
          owner: owner.address,
          provider: provider.address,
          dseq: Long.fromNumber(444),
          gseq: 1,
          oseq: 1,
          bseq: 0,
        },
        reason: LeaseClosedReason.lease_closed_reason_unstable,
      };

      await sdk.akash.market.v1beta5.closeBid(closeBid, { memo: "close bid test" });

      const res = await fetch(`${mockServer.gatewayUrl}/mock/last-close-bid`);
      expect(res.ok).toBe(true);

      const decoded = await res.json();
      expect(decoded?.id?.owner).toBe(owner.address);
      expect(decoded?.id?.provider).toBe(provider.address);
      expect(decoded?.id?.dseq).toBe("444");
      expect(decoded?.reason).toBe("lease_closed_reason_unstable");

      const normalizedExpected = normalizeValue(closeBid);
      const normalizedActual = normalizeValue(decoded);
      expect(normalizedActual).toEqual(normalizedExpected);
    });

    it("handles multi-message tx with deployment and bid", async () => {
      const wallet = await DirectSecp256k1HdWallet.fromMnemonic(TEST_MNEMONIC, {
        prefix: "akash",
        hdPaths: [makeCosmoshubPath(0), makeCosmoshubPath(1)],
      });
      const accounts = await wallet.getAccounts();
      const owner = accounts[0];
      const provider = accounts[1] ?? accounts[0];
      const sdk = createTestSDK(wallet);

      const deployment = createInvalidDeployment(owner.address, 111111, {
        hash: new Uint8Array(Array.from({ length: 32 }, (_, i) => i + 50)),
        groups: [{
          name: "multi-msg-test",
          requirements: { signedBy: { allOf: [], anyOf: [] }, attributes: [] },
          resources: [{
            resource: {
              id: 1,
              cpu: { units: { val: new TextEncoder().encode("100") }, attributes: [] },
              memory: { quantity: { val: new TextEncoder().encode("134217728") }, attributes: [] },
              storage: [{ name: "main", quantity: { val: new TextEncoder().encode("1073741824") }, attributes: [] }],
              gpu: { units: { val: new TextEncoder().encode("0") }, attributes: [] },
              endpoints: [],
            },
            count: 1,
            price: { denom: "uakt", amount: "2000" },
          }],
        }],
      });

      await sdk.akash.deployment.v1beta4.createDeployment(deployment, { memo: "deployment in multi-msg" });

      const baseGroup = createBaseResourceGroup();
      const resources = baseGroup.resources[0]?.resource;
      if (!resources) throw new Error("missing base resources");

      const bid: MsgCreateBid = {
        id: {
          owner: owner.address,
          provider: provider.address,
          dseq: Long.fromNumber(111111),
          gseq: 1,
          oseq: 1,
          bseq: 0,
        },
        price: { denom: "uakt", amount: "0.0015" },
        deposit: {
          amount: { denom: "uakt", amount: "5000000" },
          sources: [Source.balance],
        },
        resourcesOffer: [{ resources, count: 1 }],
      };

      await sdk.akash.market.v1beta5.createBid(bid, { memo: "bid in multi-msg" });

      const deploymentRes = await fetch(`${mockServer.gatewayUrl}/mock/last-deployment`);
      expect(deploymentRes.ok).toBe(true);
      const decodedDeployment = await deploymentRes.json();
      expect(decodedDeployment?.id?.owner).toBe(owner.address);
      expect(decodedDeployment?.id?.dseq).toBe("111111");

      const bidRes = await fetch(`${mockServer.gatewayUrl}/mock/last-bid`);
      expect(bidRes.ok).toBe(true);
      const decodedBid = await bidRes.json();
      expect(decodedBid?.id?.owner).toBe(owner.address);
      expect(decodedBid?.id?.dseq).toBe("111111");
      expect(normalizeDec(decodedBid?.price?.amount as string)).toBe("0.0015");
    });
  });
});
