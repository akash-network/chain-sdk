import { createHash } from "node:crypto";
import { readFileSync } from "node:fs";

import { describe, expect, it } from "vitest";

import { ID, RecordData } from "../../generated/protos/akash/sink/v1/types.ts";

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
  void_order_key_hex: string;
  void_order_value_hex: string;
  mint_proof_nonce_hex: string;
  void_proof_nonce_hex: string;
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
      const transfer = Buffer.concat([Buffer.from([0x12, 0]), suffix]);
      const order = Buffer.concat([Buffer.from([0x12, 1]), suffix]);
      expect(transfer.toString("hex")).toBe(vector.transfer_key_hex);
      expect(order.toString("hex")).toBe(vector.void_order_key_hex);
      const record = Buffer.alloc(41);
      record[0] = 2;
      Buffer.from(vector.recipient_hex, "hex").copy(record, 1);
      record.writeBigUInt64LE(BigInt(vector.amount), 33);
      expect(record.toString("hex")).toBe(vector.record_value_hex);
      expect(vector.void_order_value_hex).toBe("02");
      expect(createHash("sha256").update(transfer).digest("hex")).toBe(vector.mint_proof_nonce_hex);
      expect(createHash("sha256").update(order).digest("hex")).toBe(vector.void_proof_nonce_hex);
      const height = Buffer.alloc(8);
      height.writeBigInt64LE(id.height);
      expect(height.toString("hex")).toBe(vector.height_le_hex);
      const sequence = Buffer.alloc(8);
      sequence.writeBigUInt64LE(id.sequence);
      expect(sequence.toString("hex")).toBe(vector.sequence_le_hex);
    });
  }
});
