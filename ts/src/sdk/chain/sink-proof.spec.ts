import { createHash } from "node:crypto";
import { readFileSync } from "node:fs";

import { describe, expect, it } from "vitest";

import { ID, RecordData, SettlementData } from "../../generated/protos/akash/sink/v1/types.ts";

type ProofVector = {
  name: string;
  owner: string;
  owner_hex: string;
  height: string;
  sequence: string;
  recipient: string;
  recipient_hex: string;
  amount: string;
  id_text: string;
  transfer_key_hex: string;
  record_value_hex: string;
  mint_proof_nonce_hex: string;
  height_le_hex: string;
  sequence_le_hex: string;
};
const vectors: ProofVector[] = JSON.parse(readFileSync(new URL("../../../../go/node/sink/v1/testdata/proof_vectors.json", import.meta.url), "utf8"));

describe("sink cross-language proof vectors", () => {
  for (const vector of vectors) {
    it(vector.name, () => {
      const id = ID.fromJSON({ owner: vector.owner, height: vector.height, sequence: vector.sequence });
      expect(ID.decode(ID.encode(id).finish())).toEqual(id);
      expect(`${id.owner}/${id.height}/${id.sequence}`).toBe(vector.id_text);
      const data = RecordData.fromPartial({ dest: vector.recipient, amount: { denom: "uakt", amount: vector.amount } });
      expect(RecordData.decode(RecordData.encode(data).finish())).toEqual(data);
      const suffix = Buffer.alloc(37);
      suffix[0] = 20;
      Buffer.from(vector.owner_hex, "hex").copy(suffix, 1);
      suffix.writeBigInt64BE(id.height, 21);
      suffix.writeBigUInt64BE(id.sequence, 29);
      const transfer = Buffer.concat([Buffer.from([0x02, 0]), suffix]);
      expect(transfer.toString("hex")).toBe(vector.transfer_key_hex);
      expect(Buffer.from(RecordData.encode(data).finish()).toString("hex")).toBe(vector.record_value_hex);
      expect(createHash("sha256").update(transfer).digest("hex")).toBe(vector.mint_proof_nonce_hex);
      const height = Buffer.alloc(8);
      height.writeBigInt64LE(id.height);
      expect(height.toString("hex")).toBe(vector.height_le_hex);
      const sequence = Buffer.alloc(8);
      sequence.writeBigUInt64LE(id.sequence);
      expect(sequence.toString("hex")).toBe(vector.sequence_le_hex);
    });
  }
});

type SettlementVector = {
  name: string;
  owner_hex: string;
  height: string;
  sequence: string;
  status: number;
  settlement: unknown;
  settlement_key_hex: string;
  settlement_value_hex: string;
  proof_nonce_hex: string;
};
const settlements: SettlementVector[] = JSON.parse(readFileSync(new URL("../../../../go/node/sink/v1/testdata/settlement_vectors.json", import.meta.url), "utf8"));

describe("native canonical settlement vectors", () => {
  for (const vector of settlements) {
    it(vector.name, () => {
      const value = SettlementData.fromJSON(vector.settlement);
      const encoded = SettlementData.encode(value).finish();
      expect(Buffer.from(encoded).toString("hex")).toBe(vector.settlement_value_hex);
      expect(SettlementData.decode(encoded)).toEqual(value);
      const key = Buffer.alloc(40);
      key.set([3, 1, vector.status, 20]);
      Buffer.from(vector.owner_hex, "hex").copy(key, 4);
      key.writeBigInt64BE(BigInt(vector.height), 24);
      key.writeBigUInt64BE(BigInt(vector.sequence), 32);
      expect(key.toString("hex")).toBe(vector.settlement_key_hex);
      expect(createHash("sha256").update(key).digest("hex")).toBe(vector.proof_nonce_hex);
      if (vector.name === "maximum_evidence") expect(encoded.length).toBe(6500);
    });
  }
});
