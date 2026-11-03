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
    use super::akash::sink::v1::{
        GenesisAccount, GenesisRecord, GenesisState, Id, Params, QueryReportersResponse,
        RecordData, ReporterPolicy, SettlementData,
    };
    use super::cosmos::base::v1beta1::Coin;
    use prost::Message;
    use serde::Deserialize;
    use sha2::{Digest, Sha256};

    #[derive(Deserialize)]
    struct ParamsVector {
        name: String,
        threshold: String,
        duration_nanos: String,
        json: serde_json::Value,
        proto_hex: String,
    }

    #[test]
    fn native_params_vectors() {
        let vectors: Vec<ParamsVector> = serde_json::from_str(include_str!(
            "../../go/node/sink/v1/testdata/params_vectors.json"
        ))
        .unwrap();
        assert!(!vectors.is_empty());

        for vector in vectors {
            let wire = hex::decode(&vector.proto_hex).unwrap();
            let from_wire = Params::decode(wire.as_slice()).unwrap();
            let from_json: Params = serde_json::from_value(vector.json).unwrap();
            assert_eq!(from_json, from_wire, "{}", vector.name);
            assert_eq!(from_json.encode_to_vec(), wire, "{}", vector.name);
            assert_eq!(from_wire.encode_to_vec(), wire, "{}", vector.name);

            let json = serde_json::to_value(&from_wire).unwrap();
            assert_eq!(
                json["reportersThreshold"], vector.threshold,
                "{}",
                vector.name
            );
            let roundtrip: Params = serde_json::from_value(json).unwrap();
            assert_eq!(roundtrip.encode_to_vec(), wire, "{}", vector.name);

            let duration = from_wire.reporter_challenge_period.unwrap();
            let nanos = i128::from(duration.seconds) * 1_000_000_000 + i128::from(duration.nanos);
            assert_eq!(nanos.to_string(), vector.duration_nanos, "{}", vector.name);
        }
    }

    #[test]
    fn threshold_json_preserves_exact_decimal_precision() {
        for (decimal, atoms, canonical) in [
            ("0.75", "750000000000000000", "0.750000000000000000"),
            ("0", "0", "0.000000000000000000"),
            ("1", "1000000000000000000", "1.000000000000000000"),
            ("0.000000000000000001", "1", "0.000000000000000001"),
            (
                "0.999999999999999999",
                "999999999999999999",
                "0.999999999999999999",
            ),
            (
                "0.666666666666666667",
                "666666666666666667",
                "0.666666666666666667",
            ),
            ("+000.75", "750000000000000000", "0.750000000000000000"),
            ("-0.0", "0", "0.000000000000000000"),
            ("-0.75", "-750000000000000000", "-0.750000000000000000"),
            (
                "123456789012345678901234567890.123456789012345678",
                "123456789012345678901234567890123456789012345678",
                "123456789012345678901234567890.123456789012345678",
            ),
        ] {
            for field in ["reporters_threshold", "reportersThreshold"] {
                let params: Params =
                    serde_json::from_value(serde_json::json!({ field: decimal })).unwrap();
                assert_eq!(params.reporters_threshold, atoms, "{decimal}");
                assert_eq!(
                    serde_json::to_value(&params).unwrap()["reportersThreshold"],
                    canonical,
                    "{decimal}"
                );
            }
            let policy: ReporterPolicy =
                serde_json::from_value(serde_json::json!({ "threshold": decimal })).unwrap();
            assert_eq!(policy.threshold, atoms, "{decimal}");
            assert_eq!(
                serde_json::to_value(&policy).unwrap()["threshold"],
                canonical,
                "{decimal}"
            );
            let mut wire = vec![26, u8::try_from(atoms.len()).unwrap()];
            wire.extend(atoms.as_bytes());
            assert_eq!(policy.encode_to_vec(), wire, "{decimal}");
            assert_eq!(ReporterPolicy::decode(wire.as_slice()).unwrap(), policy);
        }
    }

    #[test]
    fn threshold_json_rejects_invalid_decimals() {
        for value in [
            serde_json::json!(0.75),
            serde_json::Value::Null,
            serde_json::json!(true),
            serde_json::json!([]),
            serde_json::json!({}),
        ] {
            assert!(serde_json::from_value::<Params>(serde_json::json!({
                "reporters_threshold": value,
            }))
            .is_err());
            assert!(serde_json::from_value::<ReporterPolicy>(serde_json::json!({
                "threshold": value,
            }))
            .is_err());
        }

        for value in [
            "",
            " ",
            ".",
            "0.",
            ".75",
            "-.75",
            "1.0.0",
            "1e-18",
            "NaN",
            "Infinity",
            "-",
            "++1",
            "--1",
            "0.0000000000000000001",
            "0.7500000000000000000",
        ] {
            assert!(
                serde_json::from_value::<Params>(serde_json::json!({
                    "reporters_threshold": value,
                }))
                .is_err(),
                "{value}"
            );
            assert!(
                serde_json::from_value::<ReporterPolicy>(serde_json::json!({
                    "threshold": value,
                }))
                .is_err(),
                "{value}"
            );
        }

        assert!(serde_json::from_str::<Params>(
            r#"{"reporters_threshold":"0.75","reportersThreshold":"1"}"#
        )
        .is_err());
        assert_eq!(
            serde_json::from_str::<Params>("{}").unwrap(),
            Params::default()
        );
        assert_eq!(
            serde_json::to_value(Params::default()).unwrap(),
            serde_json::json!({})
        );
        assert!(
            serde_json::from_str::<ReporterPolicy>(r#"{"threshold":"0.75","threshold":"1"}"#)
                .is_err()
        );
        assert_eq!(
            serde_json::from_str::<ReporterPolicy>("{}").unwrap(),
            ReporterPolicy::default()
        );
        assert_eq!(
            serde_json::to_value(ReporterPolicy::default()).unwrap(),
            serde_json::json!({})
        );
    }

    #[test]
    fn threshold_serialization_rejects_non_atomic_values() {
        for value in ["0.75", "1e18", "NaN", "-", " ", "--1"] {
            let params = Params {
                reporters_threshold: value.into(),
                ..Params::default()
            };
            assert!(serde_json::to_value(params).is_err(), "{value}");
            let policy = ReporterPolicy {
                threshold: value.into(),
                ..ReporterPolicy::default()
            };
            assert!(serde_json::to_value(policy).is_err(), "{value}");
        }
    }

    #[test]
    fn threshold_serialization_rejects_ambiguous_leading_zeros() {
        for value in ["000750000000000000000", "010", "+010", "-010", "00", "08"] {
            let params = Params {
                reporters_threshold: value.into(),
                ..Params::default()
            };
            assert!(serde_json::to_value(params).is_err(), "{value}");
            let policy = ReporterPolicy {
                threshold: value.into(),
                ..ReporterPolicy::default()
            };
            assert!(serde_json::to_value(policy).is_err(), "{value}");
        }
    }

    #[test]
    fn threshold_json_enforces_legacy_dec_range() {
        let max_whole =
            "115792089237316195423570985008687907853269984665640564039457584007913129639935";
        let overflow_whole =
            "115792089237316195423570985008687907853269984665640564039457584007913129639936";
        for sign in ["", "-"] {
            let decimal = format!("{sign}{max_whole}.999999999999999999");
            let params: Params = serde_json::from_value(serde_json::json!({
                "reporters_threshold": decimal,
            }))
            .unwrap();
            assert_eq!(
                params.reporters_threshold,
                format!("{sign}{max_whole}999999999999999999")
            );
            assert_eq!(
                serde_json::to_value(params).unwrap()["reportersThreshold"],
                decimal
            );

            assert!(serde_json::from_value::<Params>(serde_json::json!({
                "reporters_threshold": format!("{sign}{overflow_whole}"),
            }))
            .is_err());
            let overflow = Params {
                reporters_threshold: format!("{sign}{overflow_whole}000000000000000000"),
                ..Params::default()
            };
            assert!(serde_json::to_value(overflow).is_err());
            let policy: ReporterPolicy = serde_json::from_value(serde_json::json!({
                "threshold": decimal,
            }))
            .unwrap();
            assert_eq!(
                policy.threshold,
                format!("{sign}{max_whole}999999999999999999")
            );
            assert_eq!(serde_json::to_value(policy).unwrap()["threshold"], decimal);
            assert!(serde_json::from_value::<ReporterPolicy>(serde_json::json!({
                "threshold": format!("{sign}{overflow_whole}"),
            }))
            .is_err());
            let overflow = ReporterPolicy {
                threshold: format!("{sign}{overflow_whole}000000000000000000"),
                ..ReporterPolicy::default()
            };
            assert!(serde_json::to_value(overflow).is_err());
        }
    }

    #[test]
    fn nested_policy_json_and_wire_preserve_decimal_precision() {
        let policy_json = serde_json::json!({
            "version": u64::MAX.to_string(),
            "members": ["akash1reporter"],
            "threshold": "0.666666666666666667",
        });
        let policy: ReporterPolicy = serde_json::from_value(policy_json.clone()).unwrap();
        let settlement: SettlementData = serde_json::from_value(serde_json::json!({
            "policy": policy_json,
        }))
        .unwrap();
        let reporters: QueryReportersResponse = serde_json::from_value(serde_json::json!({
            "reporters": [policy_json],
        }))
        .unwrap();
        let genesis: GenesisState = serde_json::from_value(serde_json::json!({
            "policy_sequence": u64::MAX.to_string(),
            "policies": [policy_json],
            "records": [{ "settlement": { "policy": policy_json } }],
        }))
        .unwrap();

        assert_eq!(policy.threshold, "666666666666666667");
        assert_eq!(settlement.policy.as_ref(), Some(&policy));
        assert_eq!(reporters.reporters, vec![policy.clone()]);
        assert_eq!(genesis.policies, vec![policy]);
        assert_eq!(genesis.records[0].settlement.as_ref(), Some(&settlement));
        assert_eq!(
            SettlementData::decode(settlement.encode_to_vec().as_slice()).unwrap(),
            settlement
        );
        assert_eq!(
            QueryReportersResponse::decode(reporters.encode_to_vec().as_slice()).unwrap(),
            reporters
        );
        assert_eq!(
            GenesisState::decode(genesis.encode_to_vec().as_slice()).unwrap(),
            genesis
        );
        let json = serde_json::to_value(&genesis).unwrap();
        assert_eq!(json["policySequence"], u64::MAX.to_string());
        assert_eq!(json["policies"][0], policy_json);
        assert_eq!(json["records"][0]["settlement"]["policy"], policy_json);
        assert_eq!(
            serde_json::from_value::<GenesisState>(json).unwrap(),
            genesis
        );
    }

    #[test]
    fn genesis_uses_current_field_numbers_and_policy_sequence() {
        let genesis = GenesisState {
            params: Some(Params::default()),
            totals: Some(Default::default()),
            policy_sequence: 10,
            policies: vec![ReporterPolicy::default()],
            send_usage: Some(Default::default()),
            refund_usage: Some(Default::default()),
            records: vec![GenesisRecord::default()],
            accounts: vec![GenesisAccount::default()],
        };
        let wire = [10, 0, 18, 0, 24, 10, 34, 0, 42, 0, 50, 0, 58, 0, 66, 0];

        assert_eq!(genesis.encode_to_vec(), wire);
        assert_eq!(GenesisState::decode(wire.as_slice()).unwrap(), genesis);
        let json = serde_json::to_value(&genesis).unwrap();
        assert_eq!(json["policySequence"], "10");
        for removed in [
            "schemaVersion",
            "schema_version",
            "currentPolicyVersion",
            "current_policy_version",
        ] {
            assert!(json.get(removed).is_none(), "{removed}");
        }
    }

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
                assert_eq!(value.len(), 6519);
            }
        }
    }
}
