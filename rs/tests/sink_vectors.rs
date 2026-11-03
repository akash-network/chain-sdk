pub mod cosmos {
    pub mod base {
        pub mod v1beta1 {
            include!("../../.cache/sink-rust-deps/cosmos.base.v1beta1.rs");
            include!("../../.cache/sink-rust-deps/cosmos.base.v1beta1.serde.rs");
        }
        pub mod query {
            pub mod v1beta1 {
                include!("../../.cache/sink-rust-deps/cosmos.base.query.v1beta1.rs");
                include!("../../.cache/sink-rust-deps/cosmos.base.query.v1beta1.serde.rs");
            }
        }
    }
}
pub mod akash {
    pub mod sink {
        pub mod v1 {
            include!("../src/gen/node/akash.sink.v1.rs");
            include!("../src/gen/node/akash.sink.v1.serde.rs");
        }
    }
}

#[cfg(test)]
mod tests {
    use super::akash::sink::v1::{Id, RecordData, SettlementData};
    use super::cosmos::base::v1beta1::Coin;
    use prost::Message;
    use serde::Deserialize;
    use sha2::{Digest, Sha256};

    #[derive(Deserialize)]
    struct Vector {
        name: String,
        owner: String,
        owner_hex: String,
        height: String,
        sequence: String,
        recipient: String,
        amount: String,
        id_text: String,
        transfer_key_hex: String,
        record_value_hex: String,
        mint_proof_nonce_hex: String,
        height_le_hex: String,
        sequence_le_hex: String,
    }

    #[test]
    fn proof_vectors() {
        let vectors: Vec<Vector> = serde_json::from_str(include_str!(
            "../../go/node/sink/v1/testdata/proof_vectors.json"
        ))
        .unwrap();
        for v in vectors {
            let id = Id {
                owner: v.owner,
                height: v.height.parse().unwrap(),
                sequence: v.sequence.parse().unwrap(),
            };
            assert_eq!(Id::decode(id.encode_to_vec().as_slice()).unwrap(), id);
            let json = serde_json::to_string(&id).unwrap();
            assert_eq!(serde_json::from_str::<Id>(&json).unwrap(), id);
            assert_eq!(
                format!("{}/{}/{}", id.owner, id.height, id.sequence),
                v.id_text
            );
            let mut suffix = vec![20];
            suffix.extend(hex::decode(v.owner_hex).unwrap());
            suffix.extend(id.height.to_be_bytes());
            suffix.extend(id.sequence.to_be_bytes());
            let mut transfer = vec![0x02, 0];
            transfer.extend(&suffix);
            assert_eq!(hex::encode(&transfer), v.transfer_key_hex, "{}", v.name);
            let record = RecordData {
                dest: v.recipient,
                amount: Some(Coin {
                    denom: "uakt".into(),
                    amount: v.amount,
                }),
            };
            assert_eq!(hex::encode(record.encode_to_vec()), v.record_value_hex);
            assert_eq!(
                hex::encode(Sha256::digest(transfer)),
                v.mint_proof_nonce_hex
            );
            assert_eq!(hex::encode(id.height.to_le_bytes()), v.height_le_hex);
            assert_eq!(hex::encode(id.sequence.to_le_bytes()), v.sequence_le_hex);
        }
    }
    #[derive(Deserialize)]
    struct SettlementVector {
        name: String,
        owner_hex: String,
        height: String,
        sequence: String,
        status: u8,
        settlement: SettlementData,
        settlement_key_hex: String,
        settlement_value_hex: String,
        proof_nonce_hex: String,
    }

    #[test]
    fn settlement_vectors() {
        let vectors: Vec<SettlementVector> = serde_json::from_str(include_str!(
            "../../go/node/sink/v1/testdata/settlement_vectors.json"
        ))
        .unwrap();
        for v in vectors {
            let mut key = vec![3, 1, v.status, 20];
            key.extend(hex::decode(v.owner_hex).unwrap());
            key.extend(v.height.parse::<i64>().unwrap().to_be_bytes());
            key.extend(v.sequence.parse::<u64>().unwrap().to_be_bytes());
            assert_eq!(hex::encode(&key), v.settlement_key_hex, "{}", v.name);
            assert_eq!(hex::encode(Sha256::digest(&key)), v.proof_nonce_hex);
            let value = v.settlement.encode_to_vec();
            assert_eq!(hex::encode(&value), v.settlement_value_hex, "{}", v.name);
            assert_eq!(
                SettlementData::decode(value.as_slice()).unwrap(),
                v.settlement
            );
            if v.name == "maximum_evidence" {
                assert_eq!(value.len(), 6500);
            }
        }
    }
}
