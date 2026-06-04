// @generated
impl serde::Serialize for AttestationRecord {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if !self.auditor.is_empty() {
            len += 1;
        }
        if self.tier != 0 {
            len += 1;
        }
        if !self.capabilities.is_empty() {
            len += 1;
        }
        if !self.evidence_hash.is_empty() {
            len += 1;
        }
        if self.fee.is_some() {
            len += 1;
        }
        if self.fee_status != 0 {
            len += 1;
        }
        if self.created_at.is_some() {
            len += 1;
        }
        if self.expires_at.is_some() {
            len += 1;
        }
        if self.status != 0 {
            len += 1;
        }
        if self.voided_reason != 0 {
            len += 1;
        }
        if self.deposit.is_some() {
            len += 1;
        }
        if self.deposit_status != 0 {
            len += 1;
        }
        if self.audit_escrow_id != 0 {
            len += 1;
        }
        if self.fault_attribution != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.AttestationRecord", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        if self.tier != 0 {
            let v = VerificationTier::try_from(self.tier)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.tier)))?;
            struct_ser.serialize_field("tier", &v)?;
        }
        if !self.capabilities.is_empty() {
            let v = self.capabilities.iter().cloned().map(|v| {
                CapabilityFlag::try_from(v)
                    .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", v)))
                }).collect::<std::result::Result<Vec<_>, _>>()?;
            struct_ser.serialize_field("capabilities", &v)?;
        }
        if !self.evidence_hash.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("evidenceHash", pbjson::private::base64::encode(&self.evidence_hash).as_str())?;
        }
        if let Some(v) = self.fee.as_ref() {
            struct_ser.serialize_field("fee", v)?;
        }
        if self.fee_status != 0 {
            let v = FeeStatus::try_from(self.fee_status)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.fee_status)))?;
            struct_ser.serialize_field("feeStatus", &v)?;
        }
        if let Some(v) = self.created_at.as_ref() {
            struct_ser.serialize_field("createdAt", v)?;
        }
        if let Some(v) = self.expires_at.as_ref() {
            struct_ser.serialize_field("expiresAt", v)?;
        }
        if self.status != 0 {
            let v = AttestationStatus::try_from(self.status)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.status)))?;
            struct_ser.serialize_field("status", &v)?;
        }
        if self.voided_reason != 0 {
            let v = VoidedReason::try_from(self.voided_reason)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.voided_reason)))?;
            struct_ser.serialize_field("voidedReason", &v)?;
        }
        if let Some(v) = self.deposit.as_ref() {
            struct_ser.serialize_field("deposit", v)?;
        }
        if self.deposit_status != 0 {
            let v = DepositStatus::try_from(self.deposit_status)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.deposit_status)))?;
            struct_ser.serialize_field("depositStatus", &v)?;
        }
        if self.audit_escrow_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("auditEscrowId", ToString::to_string(&self.audit_escrow_id).as_str())?;
        }
        if self.fault_attribution != 0 {
            let v = FaultAttribution::try_from(self.fault_attribution)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.fault_attribution)))?;
            struct_ser.serialize_field("faultAttribution", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for AttestationRecord {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "auditor",
            "tier",
            "capabilities",
            "evidence_hash",
            "evidenceHash",
            "fee",
            "fee_status",
            "feeStatus",
            "created_at",
            "createdAt",
            "expires_at",
            "expiresAt",
            "status",
            "voided_reason",
            "voidedReason",
            "deposit",
            "deposit_status",
            "depositStatus",
            "audit_escrow_id",
            "auditEscrowId",
            "fault_attribution",
            "faultAttribution",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            Auditor,
            Tier,
            Capabilities,
            EvidenceHash,
            Fee,
            FeeStatus,
            CreatedAt,
            ExpiresAt,
            Status,
            VoidedReason,
            Deposit,
            DepositStatus,
            AuditEscrowId,
            FaultAttribution,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "auditor" => Ok(GeneratedField::Auditor),
                            "tier" => Ok(GeneratedField::Tier),
                            "capabilities" => Ok(GeneratedField::Capabilities),
                            "evidenceHash" | "evidence_hash" => Ok(GeneratedField::EvidenceHash),
                            "fee" => Ok(GeneratedField::Fee),
                            "feeStatus" | "fee_status" => Ok(GeneratedField::FeeStatus),
                            "createdAt" | "created_at" => Ok(GeneratedField::CreatedAt),
                            "expiresAt" | "expires_at" => Ok(GeneratedField::ExpiresAt),
                            "status" => Ok(GeneratedField::Status),
                            "voidedReason" | "voided_reason" => Ok(GeneratedField::VoidedReason),
                            "deposit" => Ok(GeneratedField::Deposit),
                            "depositStatus" | "deposit_status" => Ok(GeneratedField::DepositStatus),
                            "auditEscrowId" | "audit_escrow_id" => Ok(GeneratedField::AuditEscrowId),
                            "faultAttribution" | "fault_attribution" => Ok(GeneratedField::FaultAttribution),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = AttestationRecord;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.AttestationRecord")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<AttestationRecord, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut auditor__ = None;
                let mut tier__ = None;
                let mut capabilities__ = None;
                let mut evidence_hash__ = None;
                let mut fee__ = None;
                let mut fee_status__ = None;
                let mut created_at__ = None;
                let mut expires_at__ = None;
                let mut status__ = None;
                let mut voided_reason__ = None;
                let mut deposit__ = None;
                let mut deposit_status__ = None;
                let mut audit_escrow_id__ = None;
                let mut fault_attribution__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Tier => {
                            if tier__.is_some() {
                                return Err(serde::de::Error::duplicate_field("tier"));
                            }
                            tier__ = Some(map_.next_value::<VerificationTier>()? as i32);
                        }
                        GeneratedField::Capabilities => {
                            if capabilities__.is_some() {
                                return Err(serde::de::Error::duplicate_field("capabilities"));
                            }
                            capabilities__ = Some(map_.next_value::<Vec<CapabilityFlag>>()?.into_iter().map(|x| x as i32).collect());
                        }
                        GeneratedField::EvidenceHash => {
                            if evidence_hash__.is_some() {
                                return Err(serde::de::Error::duplicate_field("evidenceHash"));
                            }
                            evidence_hash__ = 
                                Some(map_.next_value::<::pbjson::private::BytesDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Fee => {
                            if fee__.is_some() {
                                return Err(serde::de::Error::duplicate_field("fee"));
                            }
                            fee__ = map_.next_value()?;
                        }
                        GeneratedField::FeeStatus => {
                            if fee_status__.is_some() {
                                return Err(serde::de::Error::duplicate_field("feeStatus"));
                            }
                            fee_status__ = Some(map_.next_value::<FeeStatus>()? as i32);
                        }
                        GeneratedField::CreatedAt => {
                            if created_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("createdAt"));
                            }
                            created_at__ = map_.next_value()?;
                        }
                        GeneratedField::ExpiresAt => {
                            if expires_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("expiresAt"));
                            }
                            expires_at__ = map_.next_value()?;
                        }
                        GeneratedField::Status => {
                            if status__.is_some() {
                                return Err(serde::de::Error::duplicate_field("status"));
                            }
                            status__ = Some(map_.next_value::<AttestationStatus>()? as i32);
                        }
                        GeneratedField::VoidedReason => {
                            if voided_reason__.is_some() {
                                return Err(serde::de::Error::duplicate_field("voidedReason"));
                            }
                            voided_reason__ = Some(map_.next_value::<VoidedReason>()? as i32);
                        }
                        GeneratedField::Deposit => {
                            if deposit__.is_some() {
                                return Err(serde::de::Error::duplicate_field("deposit"));
                            }
                            deposit__ = map_.next_value()?;
                        }
                        GeneratedField::DepositStatus => {
                            if deposit_status__.is_some() {
                                return Err(serde::de::Error::duplicate_field("depositStatus"));
                            }
                            deposit_status__ = Some(map_.next_value::<DepositStatus>()? as i32);
                        }
                        GeneratedField::AuditEscrowId => {
                            if audit_escrow_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditEscrowId"));
                            }
                            audit_escrow_id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::FaultAttribution => {
                            if fault_attribution__.is_some() {
                                return Err(serde::de::Error::duplicate_field("faultAttribution"));
                            }
                            fault_attribution__ = Some(map_.next_value::<FaultAttribution>()? as i32);
                        }
                    }
                }
                Ok(AttestationRecord {
                    provider: provider__.unwrap_or_default(),
                    auditor: auditor__.unwrap_or_default(),
                    tier: tier__.unwrap_or_default(),
                    capabilities: capabilities__.unwrap_or_default(),
                    evidence_hash: evidence_hash__.unwrap_or_default(),
                    fee: fee__,
                    fee_status: fee_status__.unwrap_or_default(),
                    created_at: created_at__,
                    expires_at: expires_at__,
                    status: status__.unwrap_or_default(),
                    voided_reason: voided_reason__.unwrap_or_default(),
                    deposit: deposit__,
                    deposit_status: deposit_status__.unwrap_or_default(),
                    audit_escrow_id: audit_escrow_id__.unwrap_or_default(),
                    fault_attribution: fault_attribution__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.AttestationRecord", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for AttestationRevocationReason {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "attestation_revocation_reason_unspecified",
            Self::ProviderNoLongerQualifies => "attestation_revocation_reason_provider_no_longer_qualifies",
            Self::SnapshotMismatch => "attestation_revocation_reason_snapshot_mismatch",
            Self::SoftwareIdentityChanged => "attestation_revocation_reason_software_identity_changed",
            Self::CapabilityMisrepresented => "attestation_revocation_reason_capability_misrepresented",
            Self::ProviderNonResponsive => "attestation_revocation_reason_provider_non_responsive",
            Self::AuditorEvidenceError => "attestation_revocation_reason_auditor_evidence_error",
            Self::AuditorOperationalExit => "attestation_revocation_reason_auditor_operational_exit",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for AttestationRevocationReason {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "attestation_revocation_reason_unspecified",
            "attestation_revocation_reason_provider_no_longer_qualifies",
            "attestation_revocation_reason_snapshot_mismatch",
            "attestation_revocation_reason_software_identity_changed",
            "attestation_revocation_reason_capability_misrepresented",
            "attestation_revocation_reason_provider_non_responsive",
            "attestation_revocation_reason_auditor_evidence_error",
            "attestation_revocation_reason_auditor_operational_exit",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = AttestationRevocationReason;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                write!(formatter, "expected one of: {:?}", &FIELDS)
            }

            fn visit_i64<E>(self, v: i64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Signed(v), &self)
                    })
            }

            fn visit_u64<E>(self, v: u64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Unsigned(v), &self)
                    })
            }

            fn visit_str<E>(self, value: &str) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                match value {
                    "attestation_revocation_reason_unspecified" => Ok(AttestationRevocationReason::Unspecified),
                    "attestation_revocation_reason_provider_no_longer_qualifies" => Ok(AttestationRevocationReason::ProviderNoLongerQualifies),
                    "attestation_revocation_reason_snapshot_mismatch" => Ok(AttestationRevocationReason::SnapshotMismatch),
                    "attestation_revocation_reason_software_identity_changed" => Ok(AttestationRevocationReason::SoftwareIdentityChanged),
                    "attestation_revocation_reason_capability_misrepresented" => Ok(AttestationRevocationReason::CapabilityMisrepresented),
                    "attestation_revocation_reason_provider_non_responsive" => Ok(AttestationRevocationReason::ProviderNonResponsive),
                    "attestation_revocation_reason_auditor_evidence_error" => Ok(AttestationRevocationReason::AuditorEvidenceError),
                    "attestation_revocation_reason_auditor_operational_exit" => Ok(AttestationRevocationReason::AuditorOperationalExit),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for AttestationStatus {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "attestation_status_unspecified",
            Self::Valid => "attestation_status_valid",
            Self::Voided => "attestation_status_voided",
            Self::Expired => "attestation_status_expired",
            Self::Revoked => "attestation_status_revoked",
            Self::Removed => "attestation_status_removed",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for AttestationStatus {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "attestation_status_unspecified",
            "attestation_status_valid",
            "attestation_status_voided",
            "attestation_status_expired",
            "attestation_status_revoked",
            "attestation_status_removed",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = AttestationStatus;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                write!(formatter, "expected one of: {:?}", &FIELDS)
            }

            fn visit_i64<E>(self, v: i64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Signed(v), &self)
                    })
            }

            fn visit_u64<E>(self, v: u64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Unsigned(v), &self)
                    })
            }

            fn visit_str<E>(self, value: &str) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                match value {
                    "attestation_status_unspecified" => Ok(AttestationStatus::Unspecified),
                    "attestation_status_valid" => Ok(AttestationStatus::Valid),
                    "attestation_status_voided" => Ok(AttestationStatus::Voided),
                    "attestation_status_expired" => Ok(AttestationStatus::Expired),
                    "attestation_status_revoked" => Ok(AttestationStatus::Revoked),
                    "attestation_status_removed" => Ok(AttestationStatus::Removed),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for AuditEscrowRecord {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.id != 0 {
            len += 1;
        }
        if !self.provider.is_empty() {
            len += 1;
        }
        if !self.consumed_by_auditor.is_empty() {
            len += 1;
        }
        if self.requested_tier != 0 {
            len += 1;
        }
        if !self.requested_capabilities.is_empty() {
            len += 1;
        }
        if self.fee.is_some() {
            len += 1;
        }
        if self.fee_status != 0 {
            len += 1;
        }
        if self.provider_deposit.is_some() {
            len += 1;
        }
        if self.provider_deposit_status != 0 {
            len += 1;
        }
        if self.status != 0 {
            len += 1;
        }
        if self.opened_at.is_some() {
            len += 1;
        }
        if self.consumed_at.is_some() {
            len += 1;
        }
        if self.expires_at.is_some() {
            len += 1;
        }
        if !self.metadata_hash.is_empty() {
            len += 1;
        }
        if self.settlement_reason != 0 {
            len += 1;
        }
        if self.fault_attribution != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.AuditEscrowRecord", len)?;
        if self.id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("id", ToString::to_string(&self.id).as_str())?;
        }
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if !self.consumed_by_auditor.is_empty() {
            struct_ser.serialize_field("consumedByAuditor", &self.consumed_by_auditor)?;
        }
        if self.requested_tier != 0 {
            let v = VerificationTier::try_from(self.requested_tier)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.requested_tier)))?;
            struct_ser.serialize_field("requestedTier", &v)?;
        }
        if !self.requested_capabilities.is_empty() {
            let v = self.requested_capabilities.iter().cloned().map(|v| {
                CapabilityFlag::try_from(v)
                    .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", v)))
                }).collect::<std::result::Result<Vec<_>, _>>()?;
            struct_ser.serialize_field("requestedCapabilities", &v)?;
        }
        if let Some(v) = self.fee.as_ref() {
            struct_ser.serialize_field("fee", v)?;
        }
        if self.fee_status != 0 {
            let v = FeeStatus::try_from(self.fee_status)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.fee_status)))?;
            struct_ser.serialize_field("feeStatus", &v)?;
        }
        if let Some(v) = self.provider_deposit.as_ref() {
            struct_ser.serialize_field("providerDeposit", v)?;
        }
        if self.provider_deposit_status != 0 {
            let v = ProviderDepositStatus::try_from(self.provider_deposit_status)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.provider_deposit_status)))?;
            struct_ser.serialize_field("providerDepositStatus", &v)?;
        }
        if self.status != 0 {
            let v = AuditEscrowStatus::try_from(self.status)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.status)))?;
            struct_ser.serialize_field("status", &v)?;
        }
        if let Some(v) = self.opened_at.as_ref() {
            struct_ser.serialize_field("openedAt", v)?;
        }
        if let Some(v) = self.consumed_at.as_ref() {
            struct_ser.serialize_field("consumedAt", v)?;
        }
        if let Some(v) = self.expires_at.as_ref() {
            struct_ser.serialize_field("expiresAt", v)?;
        }
        if !self.metadata_hash.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("metadataHash", pbjson::private::base64::encode(&self.metadata_hash).as_str())?;
        }
        if self.settlement_reason != 0 {
            let v = AuditEscrowSettlementReason::try_from(self.settlement_reason)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.settlement_reason)))?;
            struct_ser.serialize_field("settlementReason", &v)?;
        }
        if self.fault_attribution != 0 {
            let v = FaultAttribution::try_from(self.fault_attribution)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.fault_attribution)))?;
            struct_ser.serialize_field("faultAttribution", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for AuditEscrowRecord {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "id",
            "provider",
            "consumed_by_auditor",
            "consumedByAuditor",
            "requested_tier",
            "requestedTier",
            "requested_capabilities",
            "requestedCapabilities",
            "fee",
            "fee_status",
            "feeStatus",
            "provider_deposit",
            "providerDeposit",
            "provider_deposit_status",
            "providerDepositStatus",
            "status",
            "opened_at",
            "openedAt",
            "consumed_at",
            "consumedAt",
            "expires_at",
            "expiresAt",
            "metadata_hash",
            "metadataHash",
            "settlement_reason",
            "settlementReason",
            "fault_attribution",
            "faultAttribution",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Id,
            Provider,
            ConsumedByAuditor,
            RequestedTier,
            RequestedCapabilities,
            Fee,
            FeeStatus,
            ProviderDeposit,
            ProviderDepositStatus,
            Status,
            OpenedAt,
            ConsumedAt,
            ExpiresAt,
            MetadataHash,
            SettlementReason,
            FaultAttribution,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "id" => Ok(GeneratedField::Id),
                            "provider" => Ok(GeneratedField::Provider),
                            "consumedByAuditor" | "consumed_by_auditor" => Ok(GeneratedField::ConsumedByAuditor),
                            "requestedTier" | "requested_tier" => Ok(GeneratedField::RequestedTier),
                            "requestedCapabilities" | "requested_capabilities" => Ok(GeneratedField::RequestedCapabilities),
                            "fee" => Ok(GeneratedField::Fee),
                            "feeStatus" | "fee_status" => Ok(GeneratedField::FeeStatus),
                            "providerDeposit" | "provider_deposit" => Ok(GeneratedField::ProviderDeposit),
                            "providerDepositStatus" | "provider_deposit_status" => Ok(GeneratedField::ProviderDepositStatus),
                            "status" => Ok(GeneratedField::Status),
                            "openedAt" | "opened_at" => Ok(GeneratedField::OpenedAt),
                            "consumedAt" | "consumed_at" => Ok(GeneratedField::ConsumedAt),
                            "expiresAt" | "expires_at" => Ok(GeneratedField::ExpiresAt),
                            "metadataHash" | "metadata_hash" => Ok(GeneratedField::MetadataHash),
                            "settlementReason" | "settlement_reason" => Ok(GeneratedField::SettlementReason),
                            "faultAttribution" | "fault_attribution" => Ok(GeneratedField::FaultAttribution),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = AuditEscrowRecord;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.AuditEscrowRecord")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<AuditEscrowRecord, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                let mut provider__ = None;
                let mut consumed_by_auditor__ = None;
                let mut requested_tier__ = None;
                let mut requested_capabilities__ = None;
                let mut fee__ = None;
                let mut fee_status__ = None;
                let mut provider_deposit__ = None;
                let mut provider_deposit_status__ = None;
                let mut status__ = None;
                let mut opened_at__ = None;
                let mut consumed_at__ = None;
                let mut expires_at__ = None;
                let mut metadata_hash__ = None;
                let mut settlement_reason__ = None;
                let mut fault_attribution__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::ConsumedByAuditor => {
                            if consumed_by_auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("consumedByAuditor"));
                            }
                            consumed_by_auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::RequestedTier => {
                            if requested_tier__.is_some() {
                                return Err(serde::de::Error::duplicate_field("requestedTier"));
                            }
                            requested_tier__ = Some(map_.next_value::<VerificationTier>()? as i32);
                        }
                        GeneratedField::RequestedCapabilities => {
                            if requested_capabilities__.is_some() {
                                return Err(serde::de::Error::duplicate_field("requestedCapabilities"));
                            }
                            requested_capabilities__ = Some(map_.next_value::<Vec<CapabilityFlag>>()?.into_iter().map(|x| x as i32).collect());
                        }
                        GeneratedField::Fee => {
                            if fee__.is_some() {
                                return Err(serde::de::Error::duplicate_field("fee"));
                            }
                            fee__ = map_.next_value()?;
                        }
                        GeneratedField::FeeStatus => {
                            if fee_status__.is_some() {
                                return Err(serde::de::Error::duplicate_field("feeStatus"));
                            }
                            fee_status__ = Some(map_.next_value::<FeeStatus>()? as i32);
                        }
                        GeneratedField::ProviderDeposit => {
                            if provider_deposit__.is_some() {
                                return Err(serde::de::Error::duplicate_field("providerDeposit"));
                            }
                            provider_deposit__ = map_.next_value()?;
                        }
                        GeneratedField::ProviderDepositStatus => {
                            if provider_deposit_status__.is_some() {
                                return Err(serde::de::Error::duplicate_field("providerDepositStatus"));
                            }
                            provider_deposit_status__ = Some(map_.next_value::<ProviderDepositStatus>()? as i32);
                        }
                        GeneratedField::Status => {
                            if status__.is_some() {
                                return Err(serde::de::Error::duplicate_field("status"));
                            }
                            status__ = Some(map_.next_value::<AuditEscrowStatus>()? as i32);
                        }
                        GeneratedField::OpenedAt => {
                            if opened_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("openedAt"));
                            }
                            opened_at__ = map_.next_value()?;
                        }
                        GeneratedField::ConsumedAt => {
                            if consumed_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("consumedAt"));
                            }
                            consumed_at__ = map_.next_value()?;
                        }
                        GeneratedField::ExpiresAt => {
                            if expires_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("expiresAt"));
                            }
                            expires_at__ = map_.next_value()?;
                        }
                        GeneratedField::MetadataHash => {
                            if metadata_hash__.is_some() {
                                return Err(serde::de::Error::duplicate_field("metadataHash"));
                            }
                            metadata_hash__ = 
                                Some(map_.next_value::<::pbjson::private::BytesDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::SettlementReason => {
                            if settlement_reason__.is_some() {
                                return Err(serde::de::Error::duplicate_field("settlementReason"));
                            }
                            settlement_reason__ = Some(map_.next_value::<AuditEscrowSettlementReason>()? as i32);
                        }
                        GeneratedField::FaultAttribution => {
                            if fault_attribution__.is_some() {
                                return Err(serde::de::Error::duplicate_field("faultAttribution"));
                            }
                            fault_attribution__ = Some(map_.next_value::<FaultAttribution>()? as i32);
                        }
                    }
                }
                Ok(AuditEscrowRecord {
                    id: id__.unwrap_or_default(),
                    provider: provider__.unwrap_or_default(),
                    consumed_by_auditor: consumed_by_auditor__.unwrap_or_default(),
                    requested_tier: requested_tier__.unwrap_or_default(),
                    requested_capabilities: requested_capabilities__.unwrap_or_default(),
                    fee: fee__,
                    fee_status: fee_status__.unwrap_or_default(),
                    provider_deposit: provider_deposit__,
                    provider_deposit_status: provider_deposit_status__.unwrap_or_default(),
                    status: status__.unwrap_or_default(),
                    opened_at: opened_at__,
                    consumed_at: consumed_at__,
                    expires_at: expires_at__,
                    metadata_hash: metadata_hash__.unwrap_or_default(),
                    settlement_reason: settlement_reason__.unwrap_or_default(),
                    fault_attribution: fault_attribution__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.AuditEscrowRecord", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for AuditEscrowSettlementReason {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "audit_escrow_settlement_reason_unspecified",
            Self::CancelledUnconsumed => "audit_escrow_settlement_reason_cancelled_unconsumed",
            Self::ExpiredUnconsumed => "audit_escrow_settlement_reason_expired_unconsumed",
            Self::ProviderFault => "audit_escrow_settlement_reason_provider_fault",
            Self::NoFault => "audit_escrow_settlement_reason_no_fault",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for AuditEscrowSettlementReason {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "audit_escrow_settlement_reason_unspecified",
            "audit_escrow_settlement_reason_cancelled_unconsumed",
            "audit_escrow_settlement_reason_expired_unconsumed",
            "audit_escrow_settlement_reason_provider_fault",
            "audit_escrow_settlement_reason_no_fault",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = AuditEscrowSettlementReason;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                write!(formatter, "expected one of: {:?}", &FIELDS)
            }

            fn visit_i64<E>(self, v: i64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Signed(v), &self)
                    })
            }

            fn visit_u64<E>(self, v: u64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Unsigned(v), &self)
                    })
            }

            fn visit_str<E>(self, value: &str) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                match value {
                    "audit_escrow_settlement_reason_unspecified" => Ok(AuditEscrowSettlementReason::Unspecified),
                    "audit_escrow_settlement_reason_cancelled_unconsumed" => Ok(AuditEscrowSettlementReason::CancelledUnconsumed),
                    "audit_escrow_settlement_reason_expired_unconsumed" => Ok(AuditEscrowSettlementReason::ExpiredUnconsumed),
                    "audit_escrow_settlement_reason_provider_fault" => Ok(AuditEscrowSettlementReason::ProviderFault),
                    "audit_escrow_settlement_reason_no_fault" => Ok(AuditEscrowSettlementReason::NoFault),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for AuditEscrowStatus {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "audit_escrow_status_unspecified",
            Self::Open => "audit_escrow_status_open",
            Self::Consumed => "audit_escrow_status_consumed",
            Self::Cancelled => "audit_escrow_status_cancelled",
            Self::Expired => "audit_escrow_status_expired",
            Self::Settled => "audit_escrow_status_settled",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for AuditEscrowStatus {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "audit_escrow_status_unspecified",
            "audit_escrow_status_open",
            "audit_escrow_status_consumed",
            "audit_escrow_status_cancelled",
            "audit_escrow_status_expired",
            "audit_escrow_status_settled",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = AuditEscrowStatus;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                write!(formatter, "expected one of: {:?}", &FIELDS)
            }

            fn visit_i64<E>(self, v: i64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Signed(v), &self)
                    })
            }

            fn visit_u64<E>(self, v: u64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Unsigned(v), &self)
                    })
            }

            fn visit_str<E>(self, value: &str) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                match value {
                    "audit_escrow_status_unspecified" => Ok(AuditEscrowStatus::Unspecified),
                    "audit_escrow_status_open" => Ok(AuditEscrowStatus::Open),
                    "audit_escrow_status_consumed" => Ok(AuditEscrowStatus::Consumed),
                    "audit_escrow_status_cancelled" => Ok(AuditEscrowStatus::Cancelled),
                    "audit_escrow_status_expired" => Ok(AuditEscrowStatus::Expired),
                    "audit_escrow_status_settled" => Ok(AuditEscrowStatus::Settled),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for AuditorRecord {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.address.is_empty() {
            len += 1;
        }
        if self.status != 0 {
            len += 1;
        }
        if self.max_attestation_tier != 0 {
            len += 1;
        }
        if self.bond_amount.is_some() {
            len += 1;
        }
        if self.bond_status != 0 {
            len += 1;
        }
        if !self.metadata_hash.is_empty() {
            len += 1;
        }
        if self.registered_at.is_some() {
            len += 1;
        }
        if self.renewal_deadline.is_some() {
            len += 1;
        }
        if self.discrepancy_count != 0 {
            len += 1;
        }
        if self.bond_unbonding_completion_time.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.AuditorRecord", len)?;
        if !self.address.is_empty() {
            struct_ser.serialize_field("address", &self.address)?;
        }
        if self.status != 0 {
            let v = AuditorStatus::try_from(self.status)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.status)))?;
            struct_ser.serialize_field("status", &v)?;
        }
        if self.max_attestation_tier != 0 {
            let v = VerificationTier::try_from(self.max_attestation_tier)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.max_attestation_tier)))?;
            struct_ser.serialize_field("maxAttestationTier", &v)?;
        }
        if let Some(v) = self.bond_amount.as_ref() {
            struct_ser.serialize_field("bondAmount", v)?;
        }
        if self.bond_status != 0 {
            let v = BondStatus::try_from(self.bond_status)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.bond_status)))?;
            struct_ser.serialize_field("bondStatus", &v)?;
        }
        if !self.metadata_hash.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("metadataHash", pbjson::private::base64::encode(&self.metadata_hash).as_str())?;
        }
        if let Some(v) = self.registered_at.as_ref() {
            struct_ser.serialize_field("registeredAt", v)?;
        }
        if let Some(v) = self.renewal_deadline.as_ref() {
            struct_ser.serialize_field("renewalDeadline", v)?;
        }
        if self.discrepancy_count != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("discrepancyCount", ToString::to_string(&self.discrepancy_count).as_str())?;
        }
        if let Some(v) = self.bond_unbonding_completion_time.as_ref() {
            struct_ser.serialize_field("bondUnbondingCompletionTime", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for AuditorRecord {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "address",
            "status",
            "max_attestation_tier",
            "maxAttestationTier",
            "bond_amount",
            "bondAmount",
            "bond_status",
            "bondStatus",
            "metadata_hash",
            "metadataHash",
            "registered_at",
            "registeredAt",
            "renewal_deadline",
            "renewalDeadline",
            "discrepancy_count",
            "discrepancyCount",
            "bond_unbonding_completion_time",
            "bondUnbondingCompletionTime",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Address,
            Status,
            MaxAttestationTier,
            BondAmount,
            BondStatus,
            MetadataHash,
            RegisteredAt,
            RenewalDeadline,
            DiscrepancyCount,
            BondUnbondingCompletionTime,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "address" => Ok(GeneratedField::Address),
                            "status" => Ok(GeneratedField::Status),
                            "maxAttestationTier" | "max_attestation_tier" => Ok(GeneratedField::MaxAttestationTier),
                            "bondAmount" | "bond_amount" => Ok(GeneratedField::BondAmount),
                            "bondStatus" | "bond_status" => Ok(GeneratedField::BondStatus),
                            "metadataHash" | "metadata_hash" => Ok(GeneratedField::MetadataHash),
                            "registeredAt" | "registered_at" => Ok(GeneratedField::RegisteredAt),
                            "renewalDeadline" | "renewal_deadline" => Ok(GeneratedField::RenewalDeadline),
                            "discrepancyCount" | "discrepancy_count" => Ok(GeneratedField::DiscrepancyCount),
                            "bondUnbondingCompletionTime" | "bond_unbonding_completion_time" => Ok(GeneratedField::BondUnbondingCompletionTime),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = AuditorRecord;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.AuditorRecord")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<AuditorRecord, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut address__ = None;
                let mut status__ = None;
                let mut max_attestation_tier__ = None;
                let mut bond_amount__ = None;
                let mut bond_status__ = None;
                let mut metadata_hash__ = None;
                let mut registered_at__ = None;
                let mut renewal_deadline__ = None;
                let mut discrepancy_count__ = None;
                let mut bond_unbonding_completion_time__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Address => {
                            if address__.is_some() {
                                return Err(serde::de::Error::duplicate_field("address"));
                            }
                            address__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Status => {
                            if status__.is_some() {
                                return Err(serde::de::Error::duplicate_field("status"));
                            }
                            status__ = Some(map_.next_value::<AuditorStatus>()? as i32);
                        }
                        GeneratedField::MaxAttestationTier => {
                            if max_attestation_tier__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maxAttestationTier"));
                            }
                            max_attestation_tier__ = Some(map_.next_value::<VerificationTier>()? as i32);
                        }
                        GeneratedField::BondAmount => {
                            if bond_amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bondAmount"));
                            }
                            bond_amount__ = map_.next_value()?;
                        }
                        GeneratedField::BondStatus => {
                            if bond_status__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bondStatus"));
                            }
                            bond_status__ = Some(map_.next_value::<BondStatus>()? as i32);
                        }
                        GeneratedField::MetadataHash => {
                            if metadata_hash__.is_some() {
                                return Err(serde::de::Error::duplicate_field("metadataHash"));
                            }
                            metadata_hash__ = 
                                Some(map_.next_value::<::pbjson::private::BytesDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::RegisteredAt => {
                            if registered_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("registeredAt"));
                            }
                            registered_at__ = map_.next_value()?;
                        }
                        GeneratedField::RenewalDeadline => {
                            if renewal_deadline__.is_some() {
                                return Err(serde::de::Error::duplicate_field("renewalDeadline"));
                            }
                            renewal_deadline__ = map_.next_value()?;
                        }
                        GeneratedField::DiscrepancyCount => {
                            if discrepancy_count__.is_some() {
                                return Err(serde::de::Error::duplicate_field("discrepancyCount"));
                            }
                            discrepancy_count__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::BondUnbondingCompletionTime => {
                            if bond_unbonding_completion_time__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bondUnbondingCompletionTime"));
                            }
                            bond_unbonding_completion_time__ = map_.next_value()?;
                        }
                    }
                }
                Ok(AuditorRecord {
                    address: address__.unwrap_or_default(),
                    status: status__.unwrap_or_default(),
                    max_attestation_tier: max_attestation_tier__.unwrap_or_default(),
                    bond_amount: bond_amount__,
                    bond_status: bond_status__.unwrap_or_default(),
                    metadata_hash: metadata_hash__.unwrap_or_default(),
                    registered_at: registered_at__,
                    renewal_deadline: renewal_deadline__,
                    discrepancy_count: discrepancy_count__.unwrap_or_default(),
                    bond_unbonding_completion_time: bond_unbonding_completion_time__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.AuditorRecord", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for AuditorSelectionMode {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "auditor_selection_mode_unspecified",
            Self::Any => "auditor_selection_mode_any",
            Self::All => "auditor_selection_mode_all",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for AuditorSelectionMode {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "auditor_selection_mode_unspecified",
            "auditor_selection_mode_any",
            "auditor_selection_mode_all",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = AuditorSelectionMode;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                write!(formatter, "expected one of: {:?}", &FIELDS)
            }

            fn visit_i64<E>(self, v: i64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Signed(v), &self)
                    })
            }

            fn visit_u64<E>(self, v: u64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Unsigned(v), &self)
                    })
            }

            fn visit_str<E>(self, value: &str) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                match value {
                    "auditor_selection_mode_unspecified" => Ok(AuditorSelectionMode::Unspecified),
                    "auditor_selection_mode_any" => Ok(AuditorSelectionMode::Any),
                    "auditor_selection_mode_all" => Ok(AuditorSelectionMode::All),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for AuditorStatus {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "auditor_status_unspecified",
            Self::PendingBond => "auditor_status_pending_bond",
            Self::Active => "auditor_status_active",
            Self::Frozen => "auditor_status_frozen",
            Self::Lapsed => "auditor_status_lapsed",
            Self::Resigned => "auditor_status_resigned",
            Self::Removed => "auditor_status_removed",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for AuditorStatus {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "auditor_status_unspecified",
            "auditor_status_pending_bond",
            "auditor_status_active",
            "auditor_status_frozen",
            "auditor_status_lapsed",
            "auditor_status_resigned",
            "auditor_status_removed",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = AuditorStatus;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                write!(formatter, "expected one of: {:?}", &FIELDS)
            }

            fn visit_i64<E>(self, v: i64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Signed(v), &self)
                    })
            }

            fn visit_u64<E>(self, v: u64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Unsigned(v), &self)
                    })
            }

            fn visit_str<E>(self, value: &str) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                match value {
                    "auditor_status_unspecified" => Ok(AuditorStatus::Unspecified),
                    "auditor_status_pending_bond" => Ok(AuditorStatus::PendingBond),
                    "auditor_status_active" => Ok(AuditorStatus::Active),
                    "auditor_status_frozen" => Ok(AuditorStatus::Frozen),
                    "auditor_status_lapsed" => Ok(AuditorStatus::Lapsed),
                    "auditor_status_resigned" => Ok(AuditorStatus::Resigned),
                    "auditor_status_removed" => Ok(AuditorStatus::Removed),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for BondStatus {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "bond_status_unspecified",
            Self::NotBonded => "bond_status_not_bonded",
            Self::Bonded => "bond_status_bonded",
            Self::Frozen => "bond_status_frozen",
            Self::Unbonding => "bond_status_unbonding",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for BondStatus {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "bond_status_unspecified",
            "bond_status_not_bonded",
            "bond_status_bonded",
            "bond_status_frozen",
            "bond_status_unbonding",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = BondStatus;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                write!(formatter, "expected one of: {:?}", &FIELDS)
            }

            fn visit_i64<E>(self, v: i64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Signed(v), &self)
                    })
            }

            fn visit_u64<E>(self, v: u64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Unsigned(v), &self)
                    })
            }

            fn visit_str<E>(self, value: &str) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                match value {
                    "bond_status_unspecified" => Ok(BondStatus::Unspecified),
                    "bond_status_not_bonded" => Ok(BondStatus::NotBonded),
                    "bond_status_bonded" => Ok(BondStatus::Bonded),
                    "bond_status_frozen" => Ok(BondStatus::Frozen),
                    "bond_status_unbonding" => Ok(BondStatus::Unbonding),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for CapabilityFlag {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::CapabilityUnspecified => "capability_unspecified",
            Self::CapabilityTeeHardwareAttestation => "capability_tee_hardware_attestation",
            Self::CapabilityConfidentialComputing => "capability_confidential_computing",
            Self::CapabilityPersistentStorage => "capability_persistent_storage",
            Self::CapabilityBareMetal => "capability_bare_metal",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for CapabilityFlag {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "capability_unspecified",
            "capability_tee_hardware_attestation",
            "capability_confidential_computing",
            "capability_persistent_storage",
            "capability_bare_metal",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = CapabilityFlag;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                write!(formatter, "expected one of: {:?}", &FIELDS)
            }

            fn visit_i64<E>(self, v: i64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Signed(v), &self)
                    })
            }

            fn visit_u64<E>(self, v: u64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Unsigned(v), &self)
                    })
            }

            fn visit_str<E>(self, value: &str) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                match value {
                    "capability_unspecified" => Ok(CapabilityFlag::CapabilityUnspecified),
                    "capability_tee_hardware_attestation" => Ok(CapabilityFlag::CapabilityTeeHardwareAttestation),
                    "capability_confidential_computing" => Ok(CapabilityFlag::CapabilityConfidentialComputing),
                    "capability_persistent_storage" => Ok(CapabilityFlag::CapabilityPersistentStorage),
                    "capability_bare_metal" => Ok(CapabilityFlag::CapabilityBareMetal),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for DepositStatus {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "deposit_status_unspecified",
            Self::Escrowed => "deposit_status_escrowed",
            Self::PendingDiscrepancy => "deposit_status_pending_discrepancy",
            Self::ReturnedToAuditor => "deposit_status_returned_to_auditor",
            Self::Slashed => "deposit_status_slashed",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for DepositStatus {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "deposit_status_unspecified",
            "deposit_status_escrowed",
            "deposit_status_pending_discrepancy",
            "deposit_status_returned_to_auditor",
            "deposit_status_slashed",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = DepositStatus;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                write!(formatter, "expected one of: {:?}", &FIELDS)
            }

            fn visit_i64<E>(self, v: i64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Signed(v), &self)
                    })
            }

            fn visit_u64<E>(self, v: u64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Unsigned(v), &self)
                    })
            }

            fn visit_str<E>(self, value: &str) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                match value {
                    "deposit_status_unspecified" => Ok(DepositStatus::Unspecified),
                    "deposit_status_escrowed" => Ok(DepositStatus::Escrowed),
                    "deposit_status_pending_discrepancy" => Ok(DepositStatus::PendingDiscrepancy),
                    "deposit_status_returned_to_auditor" => Ok(DepositStatus::ReturnedToAuditor),
                    "deposit_status_slashed" => Ok(DepositStatus::Slashed),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for DiscrepancyEvent {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.id != 0 {
            len += 1;
        }
        if !self.provider.is_empty() {
            len += 1;
        }
        if !self.auditor_a.is_empty() {
            len += 1;
        }
        if self.auditor_a_tier != 0 {
            len += 1;
        }
        if !self.auditor_b.is_empty() {
            len += 1;
        }
        if self.auditor_b_tier != 0 {
            len += 1;
        }
        if self.timestamp.is_some() {
            len += 1;
        }
        if self.resolution_status != 0 {
            len += 1;
        }
        if self.resolution_proposal_id != 0 {
            len += 1;
        }
        if self.grace_record_id != 0 {
            len += 1;
        }
        if self.resolution_reason != 0 {
            len += 1;
        }
        if self.fault_attribution != 0 {
            len += 1;
        }
        if !self.resolution_evidence_hash.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.DiscrepancyEvent", len)?;
        if self.id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("id", ToString::to_string(&self.id).as_str())?;
        }
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if !self.auditor_a.is_empty() {
            struct_ser.serialize_field("auditorA", &self.auditor_a)?;
        }
        if self.auditor_a_tier != 0 {
            let v = VerificationTier::try_from(self.auditor_a_tier)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.auditor_a_tier)))?;
            struct_ser.serialize_field("auditorATier", &v)?;
        }
        if !self.auditor_b.is_empty() {
            struct_ser.serialize_field("auditorB", &self.auditor_b)?;
        }
        if self.auditor_b_tier != 0 {
            let v = VerificationTier::try_from(self.auditor_b_tier)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.auditor_b_tier)))?;
            struct_ser.serialize_field("auditorBTier", &v)?;
        }
        if let Some(v) = self.timestamp.as_ref() {
            struct_ser.serialize_field("timestamp", v)?;
        }
        if self.resolution_status != 0 {
            let v = DiscrepancyStatus::try_from(self.resolution_status)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.resolution_status)))?;
            struct_ser.serialize_field("resolutionStatus", &v)?;
        }
        if self.resolution_proposal_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("resolutionProposalId", ToString::to_string(&self.resolution_proposal_id).as_str())?;
        }
        if self.grace_record_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("graceRecordId", ToString::to_string(&self.grace_record_id).as_str())?;
        }
        if self.resolution_reason != 0 {
            let v = DiscrepancyResolutionReason::try_from(self.resolution_reason)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.resolution_reason)))?;
            struct_ser.serialize_field("resolutionReason", &v)?;
        }
        if self.fault_attribution != 0 {
            let v = FaultAttribution::try_from(self.fault_attribution)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.fault_attribution)))?;
            struct_ser.serialize_field("faultAttribution", &v)?;
        }
        if !self.resolution_evidence_hash.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("resolutionEvidenceHash", pbjson::private::base64::encode(&self.resolution_evidence_hash).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for DiscrepancyEvent {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "id",
            "provider",
            "auditor_a",
            "auditorA",
            "auditor_a_tier",
            "auditorATier",
            "auditor_b",
            "auditorB",
            "auditor_b_tier",
            "auditorBTier",
            "timestamp",
            "resolution_status",
            "resolutionStatus",
            "resolution_proposal_id",
            "resolutionProposalId",
            "grace_record_id",
            "graceRecordId",
            "resolution_reason",
            "resolutionReason",
            "fault_attribution",
            "faultAttribution",
            "resolution_evidence_hash",
            "resolutionEvidenceHash",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Id,
            Provider,
            AuditorA,
            AuditorATier,
            AuditorB,
            AuditorBTier,
            Timestamp,
            ResolutionStatus,
            ResolutionProposalId,
            GraceRecordId,
            ResolutionReason,
            FaultAttribution,
            ResolutionEvidenceHash,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "id" => Ok(GeneratedField::Id),
                            "provider" => Ok(GeneratedField::Provider),
                            "auditorA" | "auditor_a" => Ok(GeneratedField::AuditorA),
                            "auditorATier" | "auditor_a_tier" => Ok(GeneratedField::AuditorATier),
                            "auditorB" | "auditor_b" => Ok(GeneratedField::AuditorB),
                            "auditorBTier" | "auditor_b_tier" => Ok(GeneratedField::AuditorBTier),
                            "timestamp" => Ok(GeneratedField::Timestamp),
                            "resolutionStatus" | "resolution_status" => Ok(GeneratedField::ResolutionStatus),
                            "resolutionProposalId" | "resolution_proposal_id" => Ok(GeneratedField::ResolutionProposalId),
                            "graceRecordId" | "grace_record_id" => Ok(GeneratedField::GraceRecordId),
                            "resolutionReason" | "resolution_reason" => Ok(GeneratedField::ResolutionReason),
                            "faultAttribution" | "fault_attribution" => Ok(GeneratedField::FaultAttribution),
                            "resolutionEvidenceHash" | "resolution_evidence_hash" => Ok(GeneratedField::ResolutionEvidenceHash),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = DiscrepancyEvent;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.DiscrepancyEvent")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<DiscrepancyEvent, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                let mut provider__ = None;
                let mut auditor_a__ = None;
                let mut auditor_a_tier__ = None;
                let mut auditor_b__ = None;
                let mut auditor_b_tier__ = None;
                let mut timestamp__ = None;
                let mut resolution_status__ = None;
                let mut resolution_proposal_id__ = None;
                let mut grace_record_id__ = None;
                let mut resolution_reason__ = None;
                let mut fault_attribution__ = None;
                let mut resolution_evidence_hash__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::AuditorA => {
                            if auditor_a__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditorA"));
                            }
                            auditor_a__ = Some(map_.next_value()?);
                        }
                        GeneratedField::AuditorATier => {
                            if auditor_a_tier__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditorATier"));
                            }
                            auditor_a_tier__ = Some(map_.next_value::<VerificationTier>()? as i32);
                        }
                        GeneratedField::AuditorB => {
                            if auditor_b__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditorB"));
                            }
                            auditor_b__ = Some(map_.next_value()?);
                        }
                        GeneratedField::AuditorBTier => {
                            if auditor_b_tier__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditorBTier"));
                            }
                            auditor_b_tier__ = Some(map_.next_value::<VerificationTier>()? as i32);
                        }
                        GeneratedField::Timestamp => {
                            if timestamp__.is_some() {
                                return Err(serde::de::Error::duplicate_field("timestamp"));
                            }
                            timestamp__ = map_.next_value()?;
                        }
                        GeneratedField::ResolutionStatus => {
                            if resolution_status__.is_some() {
                                return Err(serde::de::Error::duplicate_field("resolutionStatus"));
                            }
                            resolution_status__ = Some(map_.next_value::<DiscrepancyStatus>()? as i32);
                        }
                        GeneratedField::ResolutionProposalId => {
                            if resolution_proposal_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("resolutionProposalId"));
                            }
                            resolution_proposal_id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::GraceRecordId => {
                            if grace_record_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("graceRecordId"));
                            }
                            grace_record_id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::ResolutionReason => {
                            if resolution_reason__.is_some() {
                                return Err(serde::de::Error::duplicate_field("resolutionReason"));
                            }
                            resolution_reason__ = Some(map_.next_value::<DiscrepancyResolutionReason>()? as i32);
                        }
                        GeneratedField::FaultAttribution => {
                            if fault_attribution__.is_some() {
                                return Err(serde::de::Error::duplicate_field("faultAttribution"));
                            }
                            fault_attribution__ = Some(map_.next_value::<FaultAttribution>()? as i32);
                        }
                        GeneratedField::ResolutionEvidenceHash => {
                            if resolution_evidence_hash__.is_some() {
                                return Err(serde::de::Error::duplicate_field("resolutionEvidenceHash"));
                            }
                            resolution_evidence_hash__ = 
                                Some(map_.next_value::<::pbjson::private::BytesDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(DiscrepancyEvent {
                    id: id__.unwrap_or_default(),
                    provider: provider__.unwrap_or_default(),
                    auditor_a: auditor_a__.unwrap_or_default(),
                    auditor_a_tier: auditor_a_tier__.unwrap_or_default(),
                    auditor_b: auditor_b__.unwrap_or_default(),
                    auditor_b_tier: auditor_b_tier__.unwrap_or_default(),
                    timestamp: timestamp__,
                    resolution_status: resolution_status__.unwrap_or_default(),
                    resolution_proposal_id: resolution_proposal_id__.unwrap_or_default(),
                    grace_record_id: grace_record_id__.unwrap_or_default(),
                    resolution_reason: resolution_reason__.unwrap_or_default(),
                    fault_attribution: fault_attribution__.unwrap_or_default(),
                    resolution_evidence_hash: resolution_evidence_hash__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.DiscrepancyEvent", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for DiscrepancyResolutionReason {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "discrepancy_resolution_reason_unspecified",
            Self::AuditorACorrect => "discrepancy_resolution_reason_auditor_a_correct",
            Self::AuditorBCorrect => "discrepancy_resolution_reason_auditor_b_correct",
            Self::BothAuditorsWrong => "discrepancy_resolution_reason_both_auditors_wrong",
            Self::ProviderFault => "discrepancy_resolution_reason_provider_fault",
            Self::SharedFault => "discrepancy_resolution_reason_shared_fault",
            Self::EvidenceInconclusive => "discrepancy_resolution_reason_evidence_inconclusive",
            Self::GovernanceTimeoutReview => "discrepancy_resolution_reason_governance_timeout_review",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for DiscrepancyResolutionReason {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "discrepancy_resolution_reason_unspecified",
            "discrepancy_resolution_reason_auditor_a_correct",
            "discrepancy_resolution_reason_auditor_b_correct",
            "discrepancy_resolution_reason_both_auditors_wrong",
            "discrepancy_resolution_reason_provider_fault",
            "discrepancy_resolution_reason_shared_fault",
            "discrepancy_resolution_reason_evidence_inconclusive",
            "discrepancy_resolution_reason_governance_timeout_review",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = DiscrepancyResolutionReason;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                write!(formatter, "expected one of: {:?}", &FIELDS)
            }

            fn visit_i64<E>(self, v: i64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Signed(v), &self)
                    })
            }

            fn visit_u64<E>(self, v: u64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Unsigned(v), &self)
                    })
            }

            fn visit_str<E>(self, value: &str) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                match value {
                    "discrepancy_resolution_reason_unspecified" => Ok(DiscrepancyResolutionReason::Unspecified),
                    "discrepancy_resolution_reason_auditor_a_correct" => Ok(DiscrepancyResolutionReason::AuditorACorrect),
                    "discrepancy_resolution_reason_auditor_b_correct" => Ok(DiscrepancyResolutionReason::AuditorBCorrect),
                    "discrepancy_resolution_reason_both_auditors_wrong" => Ok(DiscrepancyResolutionReason::BothAuditorsWrong),
                    "discrepancy_resolution_reason_provider_fault" => Ok(DiscrepancyResolutionReason::ProviderFault),
                    "discrepancy_resolution_reason_shared_fault" => Ok(DiscrepancyResolutionReason::SharedFault),
                    "discrepancy_resolution_reason_evidence_inconclusive" => Ok(DiscrepancyResolutionReason::EvidenceInconclusive),
                    "discrepancy_resolution_reason_governance_timeout_review" => Ok(DiscrepancyResolutionReason::GovernanceTimeoutReview),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for DiscrepancyStatus {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "discrepancy_status_unspecified",
            Self::Pending => "discrepancy_status_pending",
            Self::Resolved => "discrepancy_status_resolved",
            Self::TimedOut => "discrepancy_status_timed_out",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for DiscrepancyStatus {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "discrepancy_status_unspecified",
            "discrepancy_status_pending",
            "discrepancy_status_resolved",
            "discrepancy_status_timed_out",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = DiscrepancyStatus;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                write!(formatter, "expected one of: {:?}", &FIELDS)
            }

            fn visit_i64<E>(self, v: i64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Signed(v), &self)
                    })
            }

            fn visit_u64<E>(self, v: u64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Unsigned(v), &self)
                    })
            }

            fn visit_str<E>(self, value: &str) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                match value {
                    "discrepancy_status_unspecified" => Ok(DiscrepancyStatus::Unspecified),
                    "discrepancy_status_pending" => Ok(DiscrepancyStatus::Pending),
                    "discrepancy_status_resolved" => Ok(DiscrepancyStatus::Resolved),
                    "discrepancy_status_timed_out" => Ok(DiscrepancyStatus::TimedOut),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for EventAttestationExpired {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if !self.auditor.is_empty() {
            len += 1;
        }
        if self.tier != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventAttestationExpired", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        if self.tier != 0 {
            let v = VerificationTier::try_from(self.tier)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.tier)))?;
            struct_ser.serialize_field("tier", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventAttestationExpired {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "auditor",
            "tier",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            Auditor,
            Tier,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "auditor" => Ok(GeneratedField::Auditor),
                            "tier" => Ok(GeneratedField::Tier),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventAttestationExpired;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventAttestationExpired")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventAttestationExpired, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut auditor__ = None;
                let mut tier__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Tier => {
                            if tier__.is_some() {
                                return Err(serde::de::Error::duplicate_field("tier"));
                            }
                            tier__ = Some(map_.next_value::<VerificationTier>()? as i32);
                        }
                    }
                }
                Ok(EventAttestationExpired {
                    provider: provider__.unwrap_or_default(),
                    auditor: auditor__.unwrap_or_default(),
                    tier: tier__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventAttestationExpired", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventAttestationReplaced {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if !self.auditor.is_empty() {
            len += 1;
        }
        if self.old_tier != 0 {
            len += 1;
        }
        if self.new_tier != 0 {
            len += 1;
        }
        if self.old_audit_escrow_id != 0 {
            len += 1;
        }
        if self.new_audit_escrow_id != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventAttestationReplaced", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        if self.old_tier != 0 {
            let v = VerificationTier::try_from(self.old_tier)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.old_tier)))?;
            struct_ser.serialize_field("oldTier", &v)?;
        }
        if self.new_tier != 0 {
            let v = VerificationTier::try_from(self.new_tier)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.new_tier)))?;
            struct_ser.serialize_field("newTier", &v)?;
        }
        if self.old_audit_escrow_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("oldAuditEscrowId", ToString::to_string(&self.old_audit_escrow_id).as_str())?;
        }
        if self.new_audit_escrow_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("newAuditEscrowId", ToString::to_string(&self.new_audit_escrow_id).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventAttestationReplaced {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "auditor",
            "old_tier",
            "oldTier",
            "new_tier",
            "newTier",
            "old_audit_escrow_id",
            "oldAuditEscrowId",
            "new_audit_escrow_id",
            "newAuditEscrowId",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            Auditor,
            OldTier,
            NewTier,
            OldAuditEscrowId,
            NewAuditEscrowId,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "auditor" => Ok(GeneratedField::Auditor),
                            "oldTier" | "old_tier" => Ok(GeneratedField::OldTier),
                            "newTier" | "new_tier" => Ok(GeneratedField::NewTier),
                            "oldAuditEscrowId" | "old_audit_escrow_id" => Ok(GeneratedField::OldAuditEscrowId),
                            "newAuditEscrowId" | "new_audit_escrow_id" => Ok(GeneratedField::NewAuditEscrowId),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventAttestationReplaced;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventAttestationReplaced")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventAttestationReplaced, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut auditor__ = None;
                let mut old_tier__ = None;
                let mut new_tier__ = None;
                let mut old_audit_escrow_id__ = None;
                let mut new_audit_escrow_id__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::OldTier => {
                            if old_tier__.is_some() {
                                return Err(serde::de::Error::duplicate_field("oldTier"));
                            }
                            old_tier__ = Some(map_.next_value::<VerificationTier>()? as i32);
                        }
                        GeneratedField::NewTier => {
                            if new_tier__.is_some() {
                                return Err(serde::de::Error::duplicate_field("newTier"));
                            }
                            new_tier__ = Some(map_.next_value::<VerificationTier>()? as i32);
                        }
                        GeneratedField::OldAuditEscrowId => {
                            if old_audit_escrow_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("oldAuditEscrowId"));
                            }
                            old_audit_escrow_id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::NewAuditEscrowId => {
                            if new_audit_escrow_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("newAuditEscrowId"));
                            }
                            new_audit_escrow_id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(EventAttestationReplaced {
                    provider: provider__.unwrap_or_default(),
                    auditor: auditor__.unwrap_or_default(),
                    old_tier: old_tier__.unwrap_or_default(),
                    new_tier: new_tier__.unwrap_or_default(),
                    old_audit_escrow_id: old_audit_escrow_id__.unwrap_or_default(),
                    new_audit_escrow_id: new_audit_escrow_id__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventAttestationReplaced", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventAttestationRevoked {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if !self.auditor.is_empty() {
            len += 1;
        }
        if !self.initiator.is_empty() {
            len += 1;
        }
        if self.reason != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventAttestationRevoked", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        if !self.initiator.is_empty() {
            struct_ser.serialize_field("initiator", &self.initiator)?;
        }
        if self.reason != 0 {
            let v = AttestationRevocationReason::try_from(self.reason)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.reason)))?;
            struct_ser.serialize_field("reason", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventAttestationRevoked {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "auditor",
            "initiator",
            "reason",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            Auditor,
            Initiator,
            Reason,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "auditor" => Ok(GeneratedField::Auditor),
                            "initiator" => Ok(GeneratedField::Initiator),
                            "reason" => Ok(GeneratedField::Reason),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventAttestationRevoked;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventAttestationRevoked")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventAttestationRevoked, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut auditor__ = None;
                let mut initiator__ = None;
                let mut reason__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Initiator => {
                            if initiator__.is_some() {
                                return Err(serde::de::Error::duplicate_field("initiator"));
                            }
                            initiator__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Reason => {
                            if reason__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reason"));
                            }
                            reason__ = Some(map_.next_value::<AttestationRevocationReason>()? as i32);
                        }
                    }
                }
                Ok(EventAttestationRevoked {
                    provider: provider__.unwrap_or_default(),
                    auditor: auditor__.unwrap_or_default(),
                    initiator: initiator__.unwrap_or_default(),
                    reason: reason__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventAttestationRevoked", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventAttestationSubmitted {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if !self.auditor.is_empty() {
            len += 1;
        }
        if self.tier != 0 {
            len += 1;
        }
        if !self.capabilities.is_empty() {
            len += 1;
        }
        if self.expires_at.is_some() {
            len += 1;
        }
        if self.audit_escrow_id != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventAttestationSubmitted", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        if self.tier != 0 {
            let v = VerificationTier::try_from(self.tier)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.tier)))?;
            struct_ser.serialize_field("tier", &v)?;
        }
        if !self.capabilities.is_empty() {
            let v = self.capabilities.iter().cloned().map(|v| {
                CapabilityFlag::try_from(v)
                    .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", v)))
                }).collect::<std::result::Result<Vec<_>, _>>()?;
            struct_ser.serialize_field("capabilities", &v)?;
        }
        if let Some(v) = self.expires_at.as_ref() {
            struct_ser.serialize_field("expiresAt", v)?;
        }
        if self.audit_escrow_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("auditEscrowId", ToString::to_string(&self.audit_escrow_id).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventAttestationSubmitted {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "auditor",
            "tier",
            "capabilities",
            "expires_at",
            "expiresAt",
            "audit_escrow_id",
            "auditEscrowId",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            Auditor,
            Tier,
            Capabilities,
            ExpiresAt,
            AuditEscrowId,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "auditor" => Ok(GeneratedField::Auditor),
                            "tier" => Ok(GeneratedField::Tier),
                            "capabilities" => Ok(GeneratedField::Capabilities),
                            "expiresAt" | "expires_at" => Ok(GeneratedField::ExpiresAt),
                            "auditEscrowId" | "audit_escrow_id" => Ok(GeneratedField::AuditEscrowId),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventAttestationSubmitted;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventAttestationSubmitted")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventAttestationSubmitted, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut auditor__ = None;
                let mut tier__ = None;
                let mut capabilities__ = None;
                let mut expires_at__ = None;
                let mut audit_escrow_id__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Tier => {
                            if tier__.is_some() {
                                return Err(serde::de::Error::duplicate_field("tier"));
                            }
                            tier__ = Some(map_.next_value::<VerificationTier>()? as i32);
                        }
                        GeneratedField::Capabilities => {
                            if capabilities__.is_some() {
                                return Err(serde::de::Error::duplicate_field("capabilities"));
                            }
                            capabilities__ = Some(map_.next_value::<Vec<CapabilityFlag>>()?.into_iter().map(|x| x as i32).collect());
                        }
                        GeneratedField::ExpiresAt => {
                            if expires_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("expiresAt"));
                            }
                            expires_at__ = map_.next_value()?;
                        }
                        GeneratedField::AuditEscrowId => {
                            if audit_escrow_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditEscrowId"));
                            }
                            audit_escrow_id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(EventAttestationSubmitted {
                    provider: provider__.unwrap_or_default(),
                    auditor: auditor__.unwrap_or_default(),
                    tier: tier__.unwrap_or_default(),
                    capabilities: capabilities__.unwrap_or_default(),
                    expires_at: expires_at__,
                    audit_escrow_id: audit_escrow_id__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventAttestationSubmitted", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventAttestationVoided {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if !self.auditor.is_empty() {
            len += 1;
        }
        if self.reason != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventAttestationVoided", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        if self.reason != 0 {
            let v = VoidedReason::try_from(self.reason)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.reason)))?;
            struct_ser.serialize_field("reason", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventAttestationVoided {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "auditor",
            "reason",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            Auditor,
            Reason,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "auditor" => Ok(GeneratedField::Auditor),
                            "reason" => Ok(GeneratedField::Reason),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventAttestationVoided;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventAttestationVoided")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventAttestationVoided, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut auditor__ = None;
                let mut reason__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Reason => {
                            if reason__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reason"));
                            }
                            reason__ = Some(map_.next_value::<VoidedReason>()? as i32);
                        }
                    }
                }
                Ok(EventAttestationVoided {
                    provider: provider__.unwrap_or_default(),
                    auditor: auditor__.unwrap_or_default(),
                    reason: reason__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventAttestationVoided", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventAuditEscrowOpened {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.audit_escrow_id != 0 {
            len += 1;
        }
        if !self.provider.is_empty() {
            len += 1;
        }
        if self.fee.is_some() {
            len += 1;
        }
        if self.provider_deposit.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventAuditEscrowOpened", len)?;
        if self.audit_escrow_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("auditEscrowId", ToString::to_string(&self.audit_escrow_id).as_str())?;
        }
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if let Some(v) = self.fee.as_ref() {
            struct_ser.serialize_field("fee", v)?;
        }
        if let Some(v) = self.provider_deposit.as_ref() {
            struct_ser.serialize_field("providerDeposit", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventAuditEscrowOpened {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "audit_escrow_id",
            "auditEscrowId",
            "provider",
            "fee",
            "provider_deposit",
            "providerDeposit",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            AuditEscrowId,
            Provider,
            Fee,
            ProviderDeposit,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "auditEscrowId" | "audit_escrow_id" => Ok(GeneratedField::AuditEscrowId),
                            "provider" => Ok(GeneratedField::Provider),
                            "fee" => Ok(GeneratedField::Fee),
                            "providerDeposit" | "provider_deposit" => Ok(GeneratedField::ProviderDeposit),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventAuditEscrowOpened;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventAuditEscrowOpened")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventAuditEscrowOpened, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut audit_escrow_id__ = None;
                let mut provider__ = None;
                let mut fee__ = None;
                let mut provider_deposit__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::AuditEscrowId => {
                            if audit_escrow_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditEscrowId"));
                            }
                            audit_escrow_id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Fee => {
                            if fee__.is_some() {
                                return Err(serde::de::Error::duplicate_field("fee"));
                            }
                            fee__ = map_.next_value()?;
                        }
                        GeneratedField::ProviderDeposit => {
                            if provider_deposit__.is_some() {
                                return Err(serde::de::Error::duplicate_field("providerDeposit"));
                            }
                            provider_deposit__ = map_.next_value()?;
                        }
                    }
                }
                Ok(EventAuditEscrowOpened {
                    audit_escrow_id: audit_escrow_id__.unwrap_or_default(),
                    provider: provider__.unwrap_or_default(),
                    fee: fee__,
                    provider_deposit: provider_deposit__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventAuditEscrowOpened", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventAuditEscrowSettled {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.audit_escrow_id != 0 {
            len += 1;
        }
        if self.reason != 0 {
            len += 1;
        }
        if self.fault_attribution != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventAuditEscrowSettled", len)?;
        if self.audit_escrow_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("auditEscrowId", ToString::to_string(&self.audit_escrow_id).as_str())?;
        }
        if self.reason != 0 {
            let v = AuditEscrowSettlementReason::try_from(self.reason)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.reason)))?;
            struct_ser.serialize_field("reason", &v)?;
        }
        if self.fault_attribution != 0 {
            let v = FaultAttribution::try_from(self.fault_attribution)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.fault_attribution)))?;
            struct_ser.serialize_field("faultAttribution", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventAuditEscrowSettled {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "audit_escrow_id",
            "auditEscrowId",
            "reason",
            "fault_attribution",
            "faultAttribution",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            AuditEscrowId,
            Reason,
            FaultAttribution,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "auditEscrowId" | "audit_escrow_id" => Ok(GeneratedField::AuditEscrowId),
                            "reason" => Ok(GeneratedField::Reason),
                            "faultAttribution" | "fault_attribution" => Ok(GeneratedField::FaultAttribution),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventAuditEscrowSettled;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventAuditEscrowSettled")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventAuditEscrowSettled, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut audit_escrow_id__ = None;
                let mut reason__ = None;
                let mut fault_attribution__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::AuditEscrowId => {
                            if audit_escrow_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditEscrowId"));
                            }
                            audit_escrow_id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Reason => {
                            if reason__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reason"));
                            }
                            reason__ = Some(map_.next_value::<AuditEscrowSettlementReason>()? as i32);
                        }
                        GeneratedField::FaultAttribution => {
                            if fault_attribution__.is_some() {
                                return Err(serde::de::Error::duplicate_field("faultAttribution"));
                            }
                            fault_attribution__ = Some(map_.next_value::<FaultAttribution>()? as i32);
                        }
                    }
                }
                Ok(EventAuditEscrowSettled {
                    audit_escrow_id: audit_escrow_id__.unwrap_or_default(),
                    reason: reason__.unwrap_or_default(),
                    fault_attribution: fault_attribution__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventAuditEscrowSettled", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventAuditorBondPosted {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.auditor.is_empty() {
            len += 1;
        }
        if self.amount.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventAuditorBondPosted", len)?;
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        if let Some(v) = self.amount.as_ref() {
            struct_ser.serialize_field("amount", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventAuditorBondPosted {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "auditor",
            "amount",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Auditor,
            Amount,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "auditor" => Ok(GeneratedField::Auditor),
                            "amount" => Ok(GeneratedField::Amount),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventAuditorBondPosted;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventAuditorBondPosted")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventAuditorBondPosted, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut auditor__ = None;
                let mut amount__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Amount => {
                            if amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("amount"));
                            }
                            amount__ = map_.next_value()?;
                        }
                    }
                }
                Ok(EventAuditorBondPosted {
                    auditor: auditor__.unwrap_or_default(),
                    amount: amount__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventAuditorBondPosted", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventAuditorFrozen {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.auditor.is_empty() {
            len += 1;
        }
        if self.discrepancy_id != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventAuditorFrozen", len)?;
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        if self.discrepancy_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("discrepancyId", ToString::to_string(&self.discrepancy_id).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventAuditorFrozen {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "auditor",
            "discrepancy_id",
            "discrepancyId",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Auditor,
            DiscrepancyId,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "auditor" => Ok(GeneratedField::Auditor),
                            "discrepancyId" | "discrepancy_id" => Ok(GeneratedField::DiscrepancyId),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventAuditorFrozen;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventAuditorFrozen")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventAuditorFrozen, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut auditor__ = None;
                let mut discrepancy_id__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::DiscrepancyId => {
                            if discrepancy_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("discrepancyId"));
                            }
                            discrepancy_id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(EventAuditorFrozen {
                    auditor: auditor__.unwrap_or_default(),
                    discrepancy_id: discrepancy_id__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventAuditorFrozen", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventAuditorLapsed {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.auditor.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventAuditorLapsed", len)?;
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventAuditorLapsed {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "auditor",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Auditor,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "auditor" => Ok(GeneratedField::Auditor),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventAuditorLapsed;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventAuditorLapsed")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventAuditorLapsed, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut auditor__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(EventAuditorLapsed {
                    auditor: auditor__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventAuditorLapsed", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventAuditorRegistered {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.auditor.is_empty() {
            len += 1;
        }
        if self.max_attestation_tier != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventAuditorRegistered", len)?;
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        if self.max_attestation_tier != 0 {
            let v = VerificationTier::try_from(self.max_attestation_tier)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.max_attestation_tier)))?;
            struct_ser.serialize_field("maxAttestationTier", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventAuditorRegistered {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "auditor",
            "max_attestation_tier",
            "maxAttestationTier",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Auditor,
            MaxAttestationTier,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "auditor" => Ok(GeneratedField::Auditor),
                            "maxAttestationTier" | "max_attestation_tier" => Ok(GeneratedField::MaxAttestationTier),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventAuditorRegistered;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventAuditorRegistered")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventAuditorRegistered, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut auditor__ = None;
                let mut max_attestation_tier__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::MaxAttestationTier => {
                            if max_attestation_tier__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maxAttestationTier"));
                            }
                            max_attestation_tier__ = Some(map_.next_value::<VerificationTier>()? as i32);
                        }
                    }
                }
                Ok(EventAuditorRegistered {
                    auditor: auditor__.unwrap_or_default(),
                    max_attestation_tier: max_attestation_tier__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventAuditorRegistered", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventAuditorRemoved {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.auditor.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventAuditorRemoved", len)?;
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventAuditorRemoved {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "auditor",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Auditor,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "auditor" => Ok(GeneratedField::Auditor),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventAuditorRemoved;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventAuditorRemoved")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventAuditorRemoved, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut auditor__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(EventAuditorRemoved {
                    auditor: auditor__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventAuditorRemoved", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventAuditorRenewed {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.auditor.is_empty() {
            len += 1;
        }
        if self.new_deadline.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventAuditorRenewed", len)?;
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        if let Some(v) = self.new_deadline.as_ref() {
            struct_ser.serialize_field("newDeadline", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventAuditorRenewed {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "auditor",
            "new_deadline",
            "newDeadline",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Auditor,
            NewDeadline,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "auditor" => Ok(GeneratedField::Auditor),
                            "newDeadline" | "new_deadline" => Ok(GeneratedField::NewDeadline),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventAuditorRenewed;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventAuditorRenewed")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventAuditorRenewed, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut auditor__ = None;
                let mut new_deadline__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::NewDeadline => {
                            if new_deadline__.is_some() {
                                return Err(serde::de::Error::duplicate_field("newDeadline"));
                            }
                            new_deadline__ = map_.next_value()?;
                        }
                    }
                }
                Ok(EventAuditorRenewed {
                    auditor: auditor__.unwrap_or_default(),
                    new_deadline: new_deadline__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventAuditorRenewed", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventAuditorResigned {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.auditor.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventAuditorResigned", len)?;
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventAuditorResigned {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "auditor",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Auditor,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "auditor" => Ok(GeneratedField::Auditor),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventAuditorResigned;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventAuditorResigned")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventAuditorResigned, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut auditor__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(EventAuditorResigned {
                    auditor: auditor__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventAuditorResigned", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventDepositReturnedToAuditor {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.auditor.is_empty() {
            len += 1;
        }
        if self.amount.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventDepositReturnedToAuditor", len)?;
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        if let Some(v) = self.amount.as_ref() {
            struct_ser.serialize_field("amount", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventDepositReturnedToAuditor {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "auditor",
            "amount",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Auditor,
            Amount,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "auditor" => Ok(GeneratedField::Auditor),
                            "amount" => Ok(GeneratedField::Amount),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventDepositReturnedToAuditor;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventDepositReturnedToAuditor")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventDepositReturnedToAuditor, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut auditor__ = None;
                let mut amount__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Amount => {
                            if amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("amount"));
                            }
                            amount__ = map_.next_value()?;
                        }
                    }
                }
                Ok(EventDepositReturnedToAuditor {
                    auditor: auditor__.unwrap_or_default(),
                    amount: amount__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventDepositReturnedToAuditor", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventDepositSlashed {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.auditor.is_empty() {
            len += 1;
        }
        if self.amount.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventDepositSlashed", len)?;
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        if let Some(v) = self.amount.as_ref() {
            struct_ser.serialize_field("amount", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventDepositSlashed {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "auditor",
            "amount",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Auditor,
            Amount,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "auditor" => Ok(GeneratedField::Auditor),
                            "amount" => Ok(GeneratedField::Amount),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventDepositSlashed;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventDepositSlashed")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventDepositSlashed, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut auditor__ = None;
                let mut amount__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Amount => {
                            if amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("amount"));
                            }
                            amount__ = map_.next_value()?;
                        }
                    }
                }
                Ok(EventDepositSlashed {
                    auditor: auditor__.unwrap_or_default(),
                    amount: amount__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventDepositSlashed", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventDiscrepancyDetected {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.discrepancy_id != 0 {
            len += 1;
        }
        if !self.provider.is_empty() {
            len += 1;
        }
        if !self.auditor_a.is_empty() {
            len += 1;
        }
        if self.tier_a != 0 {
            len += 1;
        }
        if !self.auditor_b.is_empty() {
            len += 1;
        }
        if self.tier_b != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventDiscrepancyDetected", len)?;
        if self.discrepancy_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("discrepancyId", ToString::to_string(&self.discrepancy_id).as_str())?;
        }
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if !self.auditor_a.is_empty() {
            struct_ser.serialize_field("auditorA", &self.auditor_a)?;
        }
        if self.tier_a != 0 {
            let v = VerificationTier::try_from(self.tier_a)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.tier_a)))?;
            struct_ser.serialize_field("tierA", &v)?;
        }
        if !self.auditor_b.is_empty() {
            struct_ser.serialize_field("auditorB", &self.auditor_b)?;
        }
        if self.tier_b != 0 {
            let v = VerificationTier::try_from(self.tier_b)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.tier_b)))?;
            struct_ser.serialize_field("tierB", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventDiscrepancyDetected {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "discrepancy_id",
            "discrepancyId",
            "provider",
            "auditor_a",
            "auditorA",
            "tier_a",
            "tierA",
            "auditor_b",
            "auditorB",
            "tier_b",
            "tierB",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            DiscrepancyId,
            Provider,
            AuditorA,
            TierA,
            AuditorB,
            TierB,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "discrepancyId" | "discrepancy_id" => Ok(GeneratedField::DiscrepancyId),
                            "provider" => Ok(GeneratedField::Provider),
                            "auditorA" | "auditor_a" => Ok(GeneratedField::AuditorA),
                            "tierA" | "tier_a" => Ok(GeneratedField::TierA),
                            "auditorB" | "auditor_b" => Ok(GeneratedField::AuditorB),
                            "tierB" | "tier_b" => Ok(GeneratedField::TierB),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventDiscrepancyDetected;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventDiscrepancyDetected")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventDiscrepancyDetected, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut discrepancy_id__ = None;
                let mut provider__ = None;
                let mut auditor_a__ = None;
                let mut tier_a__ = None;
                let mut auditor_b__ = None;
                let mut tier_b__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::DiscrepancyId => {
                            if discrepancy_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("discrepancyId"));
                            }
                            discrepancy_id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::AuditorA => {
                            if auditor_a__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditorA"));
                            }
                            auditor_a__ = Some(map_.next_value()?);
                        }
                        GeneratedField::TierA => {
                            if tier_a__.is_some() {
                                return Err(serde::de::Error::duplicate_field("tierA"));
                            }
                            tier_a__ = Some(map_.next_value::<VerificationTier>()? as i32);
                        }
                        GeneratedField::AuditorB => {
                            if auditor_b__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditorB"));
                            }
                            auditor_b__ = Some(map_.next_value()?);
                        }
                        GeneratedField::TierB => {
                            if tier_b__.is_some() {
                                return Err(serde::de::Error::duplicate_field("tierB"));
                            }
                            tier_b__ = Some(map_.next_value::<VerificationTier>()? as i32);
                        }
                    }
                }
                Ok(EventDiscrepancyDetected {
                    discrepancy_id: discrepancy_id__.unwrap_or_default(),
                    provider: provider__.unwrap_or_default(),
                    auditor_a: auditor_a__.unwrap_or_default(),
                    tier_a: tier_a__.unwrap_or_default(),
                    auditor_b: auditor_b__.unwrap_or_default(),
                    tier_b: tier_b__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventDiscrepancyDetected", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventDiscrepancyResolved {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.discrepancy_id != 0 {
            len += 1;
        }
        if !self.vindicated_auditor.is_empty() {
            len += 1;
        }
        if self.reason != 0 {
            len += 1;
        }
        if self.fault_attribution != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventDiscrepancyResolved", len)?;
        if self.discrepancy_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("discrepancyId", ToString::to_string(&self.discrepancy_id).as_str())?;
        }
        if !self.vindicated_auditor.is_empty() {
            struct_ser.serialize_field("vindicatedAuditor", &self.vindicated_auditor)?;
        }
        if self.reason != 0 {
            let v = DiscrepancyResolutionReason::try_from(self.reason)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.reason)))?;
            struct_ser.serialize_field("reason", &v)?;
        }
        if self.fault_attribution != 0 {
            let v = FaultAttribution::try_from(self.fault_attribution)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.fault_attribution)))?;
            struct_ser.serialize_field("faultAttribution", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventDiscrepancyResolved {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "discrepancy_id",
            "discrepancyId",
            "vindicated_auditor",
            "vindicatedAuditor",
            "reason",
            "fault_attribution",
            "faultAttribution",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            DiscrepancyId,
            VindicatedAuditor,
            Reason,
            FaultAttribution,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "discrepancyId" | "discrepancy_id" => Ok(GeneratedField::DiscrepancyId),
                            "vindicatedAuditor" | "vindicated_auditor" => Ok(GeneratedField::VindicatedAuditor),
                            "reason" => Ok(GeneratedField::Reason),
                            "faultAttribution" | "fault_attribution" => Ok(GeneratedField::FaultAttribution),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventDiscrepancyResolved;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventDiscrepancyResolved")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventDiscrepancyResolved, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut discrepancy_id__ = None;
                let mut vindicated_auditor__ = None;
                let mut reason__ = None;
                let mut fault_attribution__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::DiscrepancyId => {
                            if discrepancy_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("discrepancyId"));
                            }
                            discrepancy_id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::VindicatedAuditor => {
                            if vindicated_auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("vindicatedAuditor"));
                            }
                            vindicated_auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Reason => {
                            if reason__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reason"));
                            }
                            reason__ = Some(map_.next_value::<DiscrepancyResolutionReason>()? as i32);
                        }
                        GeneratedField::FaultAttribution => {
                            if fault_attribution__.is_some() {
                                return Err(serde::de::Error::duplicate_field("faultAttribution"));
                            }
                            fault_attribution__ = Some(map_.next_value::<FaultAttribution>()? as i32);
                        }
                    }
                }
                Ok(EventDiscrepancyResolved {
                    discrepancy_id: discrepancy_id__.unwrap_or_default(),
                    vindicated_auditor: vindicated_auditor__.unwrap_or_default(),
                    reason: reason__.unwrap_or_default(),
                    fault_attribution: fault_attribution__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventDiscrepancyResolved", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventDiscrepancyTimedOut {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.discrepancy_id != 0 {
            len += 1;
        }
        if !self.auditor_a.is_empty() {
            len += 1;
        }
        if !self.auditor_b.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventDiscrepancyTimedOut", len)?;
        if self.discrepancy_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("discrepancyId", ToString::to_string(&self.discrepancy_id).as_str())?;
        }
        if !self.auditor_a.is_empty() {
            struct_ser.serialize_field("auditorA", &self.auditor_a)?;
        }
        if !self.auditor_b.is_empty() {
            struct_ser.serialize_field("auditorB", &self.auditor_b)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventDiscrepancyTimedOut {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "discrepancy_id",
            "discrepancyId",
            "auditor_a",
            "auditorA",
            "auditor_b",
            "auditorB",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            DiscrepancyId,
            AuditorA,
            AuditorB,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "discrepancyId" | "discrepancy_id" => Ok(GeneratedField::DiscrepancyId),
                            "auditorA" | "auditor_a" => Ok(GeneratedField::AuditorA),
                            "auditorB" | "auditor_b" => Ok(GeneratedField::AuditorB),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventDiscrepancyTimedOut;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventDiscrepancyTimedOut")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventDiscrepancyTimedOut, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut discrepancy_id__ = None;
                let mut auditor_a__ = None;
                let mut auditor_b__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::DiscrepancyId => {
                            if discrepancy_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("discrepancyId"));
                            }
                            discrepancy_id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::AuditorA => {
                            if auditor_a__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditorA"));
                            }
                            auditor_a__ = Some(map_.next_value()?);
                        }
                        GeneratedField::AuditorB => {
                            if auditor_b__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditorB"));
                            }
                            auditor_b__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(EventDiscrepancyTimedOut {
                    discrepancy_id: discrepancy_id__.unwrap_or_default(),
                    auditor_a: auditor_a__.unwrap_or_default(),
                    auditor_b: auditor_b__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventDiscrepancyTimedOut", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventFeeEscrowed {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if !self.auditor.is_empty() {
            len += 1;
        }
        if self.amount.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventFeeEscrowed", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        if let Some(v) = self.amount.as_ref() {
            struct_ser.serialize_field("amount", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventFeeEscrowed {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "auditor",
            "amount",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            Auditor,
            Amount,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "auditor" => Ok(GeneratedField::Auditor),
                            "amount" => Ok(GeneratedField::Amount),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventFeeEscrowed;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventFeeEscrowed")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventFeeEscrowed, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut auditor__ = None;
                let mut amount__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Amount => {
                            if amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("amount"));
                            }
                            amount__ = map_.next_value()?;
                        }
                    }
                }
                Ok(EventFeeEscrowed {
                    provider: provider__.unwrap_or_default(),
                    auditor: auditor__.unwrap_or_default(),
                    amount: amount__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventFeeEscrowed", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventFeeReleasedToAuditor {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.auditor.is_empty() {
            len += 1;
        }
        if self.amount.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventFeeReleasedToAuditor", len)?;
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        if let Some(v) = self.amount.as_ref() {
            struct_ser.serialize_field("amount", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventFeeReleasedToAuditor {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "auditor",
            "amount",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Auditor,
            Amount,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "auditor" => Ok(GeneratedField::Auditor),
                            "amount" => Ok(GeneratedField::Amount),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventFeeReleasedToAuditor;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventFeeReleasedToAuditor")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventFeeReleasedToAuditor, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut auditor__ = None;
                let mut amount__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Amount => {
                            if amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("amount"));
                            }
                            amount__ = map_.next_value()?;
                        }
                    }
                }
                Ok(EventFeeReleasedToAuditor {
                    auditor: auditor__.unwrap_or_default(),
                    amount: amount__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventFeeReleasedToAuditor", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventFeeReturnedToProvider {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if self.amount.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventFeeReturnedToProvider", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if let Some(v) = self.amount.as_ref() {
            struct_ser.serialize_field("amount", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventFeeReturnedToProvider {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "amount",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            Amount,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "amount" => Ok(GeneratedField::Amount),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventFeeReturnedToProvider;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventFeeReturnedToProvider")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventFeeReturnedToProvider, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut amount__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Amount => {
                            if amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("amount"));
                            }
                            amount__ = map_.next_value()?;
                        }
                    }
                }
                Ok(EventFeeReturnedToProvider {
                    provider: provider__.unwrap_or_default(),
                    amount: amount__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventFeeReturnedToProvider", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventProviderBondPosted {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if self.amount.is_some() {
            len += 1;
        }
        if self.total_bonded.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventProviderBondPosted", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if let Some(v) = self.amount.as_ref() {
            struct_ser.serialize_field("amount", v)?;
        }
        if let Some(v) = self.total_bonded.as_ref() {
            struct_ser.serialize_field("totalBonded", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventProviderBondPosted {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "amount",
            "total_bonded",
            "totalBonded",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            Amount,
            TotalBonded,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "amount" => Ok(GeneratedField::Amount),
                            "totalBonded" | "total_bonded" => Ok(GeneratedField::TotalBonded),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventProviderBondPosted;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventProviderBondPosted")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventProviderBondPosted, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut amount__ = None;
                let mut total_bonded__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Amount => {
                            if amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("amount"));
                            }
                            amount__ = map_.next_value()?;
                        }
                        GeneratedField::TotalBonded => {
                            if total_bonded__.is_some() {
                                return Err(serde::de::Error::duplicate_field("totalBonded"));
                            }
                            total_bonded__ = map_.next_value()?;
                        }
                    }
                }
                Ok(EventProviderBondPosted {
                    provider: provider__.unwrap_or_default(),
                    amount: amount__,
                    total_bonded: total_bonded__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventProviderBondPosted", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventProviderBondSlashed {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if self.slashed_amount.is_some() {
            len += 1;
        }
        if self.reason != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventProviderBondSlashed", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if let Some(v) = self.slashed_amount.as_ref() {
            struct_ser.serialize_field("slashedAmount", v)?;
        }
        if self.reason != 0 {
            let v = ProviderBondSlashReason::try_from(self.reason)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.reason)))?;
            struct_ser.serialize_field("reason", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventProviderBondSlashed {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "slashed_amount",
            "slashedAmount",
            "reason",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            SlashedAmount,
            Reason,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "slashedAmount" | "slashed_amount" => Ok(GeneratedField::SlashedAmount),
                            "reason" => Ok(GeneratedField::Reason),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventProviderBondSlashed;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventProviderBondSlashed")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventProviderBondSlashed, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut slashed_amount__ = None;
                let mut reason__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::SlashedAmount => {
                            if slashed_amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("slashedAmount"));
                            }
                            slashed_amount__ = map_.next_value()?;
                        }
                        GeneratedField::Reason => {
                            if reason__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reason"));
                            }
                            reason__ = Some(map_.next_value::<ProviderBondSlashReason>()? as i32);
                        }
                    }
                }
                Ok(EventProviderBondSlashed {
                    provider: provider__.unwrap_or_default(),
                    slashed_amount: slashed_amount__,
                    reason: reason__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventProviderBondSlashed", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventProviderBondWithdrawalCompleted {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if self.amount.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventProviderBondWithdrawalCompleted", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if let Some(v) = self.amount.as_ref() {
            struct_ser.serialize_field("amount", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventProviderBondWithdrawalCompleted {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "amount",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            Amount,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "amount" => Ok(GeneratedField::Amount),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventProviderBondWithdrawalCompleted;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventProviderBondWithdrawalCompleted")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventProviderBondWithdrawalCompleted, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut amount__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Amount => {
                            if amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("amount"));
                            }
                            amount__ = map_.next_value()?;
                        }
                    }
                }
                Ok(EventProviderBondWithdrawalCompleted {
                    provider: provider__.unwrap_or_default(),
                    amount: amount__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventProviderBondWithdrawalCompleted", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventProviderBondWithdrawalInitiated {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if self.amount.is_some() {
            len += 1;
        }
        if self.completion_time.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventProviderBondWithdrawalInitiated", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if let Some(v) = self.amount.as_ref() {
            struct_ser.serialize_field("amount", v)?;
        }
        if let Some(v) = self.completion_time.as_ref() {
            struct_ser.serialize_field("completionTime", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventProviderBondWithdrawalInitiated {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "amount",
            "completion_time",
            "completionTime",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            Amount,
            CompletionTime,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "amount" => Ok(GeneratedField::Amount),
                            "completionTime" | "completion_time" => Ok(GeneratedField::CompletionTime),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventProviderBondWithdrawalInitiated;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventProviderBondWithdrawalInitiated")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventProviderBondWithdrawalInitiated, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut amount__ = None;
                let mut completion_time__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Amount => {
                            if amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("amount"));
                            }
                            amount__ = map_.next_value()?;
                        }
                        GeneratedField::CompletionTime => {
                            if completion_time__.is_some() {
                                return Err(serde::de::Error::duplicate_field("completionTime"));
                            }
                            completion_time__ = map_.next_value()?;
                        }
                    }
                }
                Ok(EventProviderBondWithdrawalInitiated {
                    provider: provider__.unwrap_or_default(),
                    amount: amount__,
                    completion_time: completion_time__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventProviderBondWithdrawalInitiated", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventSnapshotHashPosted {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if !self.snapshot_hash.is_empty() {
            len += 1;
        }
        if self.compliance_deadline.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventSnapshotHashPosted", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if !self.snapshot_hash.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("snapshotHash", pbjson::private::base64::encode(&self.snapshot_hash).as_str())?;
        }
        if let Some(v) = self.compliance_deadline.as_ref() {
            struct_ser.serialize_field("complianceDeadline", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventSnapshotHashPosted {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "snapshot_hash",
            "snapshotHash",
            "compliance_deadline",
            "complianceDeadline",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            SnapshotHash,
            ComplianceDeadline,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "snapshotHash" | "snapshot_hash" => Ok(GeneratedField::SnapshotHash),
                            "complianceDeadline" | "compliance_deadline" => Ok(GeneratedField::ComplianceDeadline),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventSnapshotHashPosted;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventSnapshotHashPosted")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventSnapshotHashPosted, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut snapshot_hash__ = None;
                let mut compliance_deadline__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::SnapshotHash => {
                            if snapshot_hash__.is_some() {
                                return Err(serde::de::Error::duplicate_field("snapshotHash"));
                            }
                            snapshot_hash__ = 
                                Some(map_.next_value::<::pbjson::private::BytesDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::ComplianceDeadline => {
                            if compliance_deadline__.is_some() {
                                return Err(serde::de::Error::duplicate_field("complianceDeadline"));
                            }
                            compliance_deadline__ = map_.next_value()?;
                        }
                    }
                }
                Ok(EventSnapshotHashPosted {
                    provider: provider__.unwrap_or_default(),
                    snapshot_hash: snapshot_hash__.unwrap_or_default(),
                    compliance_deadline: compliance_deadline__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventSnapshotHashPosted", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventSnapshotResumed {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventSnapshotResumed", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventSnapshotResumed {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventSnapshotResumed;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventSnapshotResumed")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventSnapshotResumed, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(EventSnapshotResumed {
                    provider: provider__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventSnapshotResumed", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventSnapshotSuspended {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventSnapshotSuspended", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventSnapshotSuspended {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventSnapshotSuspended;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventSnapshotSuspended")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventSnapshotSuspended, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(EventSnapshotSuspended {
                    provider: provider__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventSnapshotSuspended", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventVerificationGraceEnded {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.grace_record_id != 0 {
            len += 1;
        }
        if self.status != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventVerificationGraceEnded", len)?;
        if self.grace_record_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("graceRecordId", ToString::to_string(&self.grace_record_id).as_str())?;
        }
        if self.status != 0 {
            let v = VerificationGraceStatus::try_from(self.status)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.status)))?;
            struct_ser.serialize_field("status", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventVerificationGraceEnded {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "grace_record_id",
            "graceRecordId",
            "status",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            GraceRecordId,
            Status,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "graceRecordId" | "grace_record_id" => Ok(GeneratedField::GraceRecordId),
                            "status" => Ok(GeneratedField::Status),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventVerificationGraceEnded;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventVerificationGraceEnded")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventVerificationGraceEnded, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut grace_record_id__ = None;
                let mut status__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::GraceRecordId => {
                            if grace_record_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("graceRecordId"));
                            }
                            grace_record_id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Status => {
                            if status__.is_some() {
                                return Err(serde::de::Error::duplicate_field("status"));
                            }
                            status__ = Some(map_.next_value::<VerificationGraceStatus>()? as i32);
                        }
                    }
                }
                Ok(EventVerificationGraceEnded {
                    grace_record_id: grace_record_id__.unwrap_or_default(),
                    status: status__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventVerificationGraceEnded", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventVerificationGraceStarted {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.grace_record_id != 0 {
            len += 1;
        }
        if !self.provider.is_empty() {
            len += 1;
        }
        if self.preserved_tier != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.EventVerificationGraceStarted", len)?;
        if self.grace_record_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("graceRecordId", ToString::to_string(&self.grace_record_id).as_str())?;
        }
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if self.preserved_tier != 0 {
            let v = VerificationTier::try_from(self.preserved_tier)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.preserved_tier)))?;
            struct_ser.serialize_field("preservedTier", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventVerificationGraceStarted {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "grace_record_id",
            "graceRecordId",
            "provider",
            "preserved_tier",
            "preservedTier",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            GraceRecordId,
            Provider,
            PreservedTier,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "graceRecordId" | "grace_record_id" => Ok(GeneratedField::GraceRecordId),
                            "provider" => Ok(GeneratedField::Provider),
                            "preservedTier" | "preserved_tier" => Ok(GeneratedField::PreservedTier),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventVerificationGraceStarted;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.EventVerificationGraceStarted")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventVerificationGraceStarted, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut grace_record_id__ = None;
                let mut provider__ = None;
                let mut preserved_tier__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::GraceRecordId => {
                            if grace_record_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("graceRecordId"));
                            }
                            grace_record_id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::PreservedTier => {
                            if preserved_tier__.is_some() {
                                return Err(serde::de::Error::duplicate_field("preservedTier"));
                            }
                            preserved_tier__ = Some(map_.next_value::<VerificationTier>()? as i32);
                        }
                    }
                }
                Ok(EventVerificationGraceStarted {
                    grace_record_id: grace_record_id__.unwrap_or_default(),
                    provider: provider__.unwrap_or_default(),
                    preserved_tier: preserved_tier__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.EventVerificationGraceStarted", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for FaultAttribution {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "fault_attribution_unspecified",
            Self::ProviderFault => "fault_attribution_provider_fault",
            Self::AuditorFault => "fault_attribution_auditor_fault",
            Self::SharedFault => "fault_attribution_shared_fault",
            Self::NoFault => "fault_attribution_no_fault",
            Self::Inconclusive => "fault_attribution_inconclusive",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for FaultAttribution {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "fault_attribution_unspecified",
            "fault_attribution_provider_fault",
            "fault_attribution_auditor_fault",
            "fault_attribution_shared_fault",
            "fault_attribution_no_fault",
            "fault_attribution_inconclusive",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = FaultAttribution;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                write!(formatter, "expected one of: {:?}", &FIELDS)
            }

            fn visit_i64<E>(self, v: i64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Signed(v), &self)
                    })
            }

            fn visit_u64<E>(self, v: u64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Unsigned(v), &self)
                    })
            }

            fn visit_str<E>(self, value: &str) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                match value {
                    "fault_attribution_unspecified" => Ok(FaultAttribution::Unspecified),
                    "fault_attribution_provider_fault" => Ok(FaultAttribution::ProviderFault),
                    "fault_attribution_auditor_fault" => Ok(FaultAttribution::AuditorFault),
                    "fault_attribution_shared_fault" => Ok(FaultAttribution::SharedFault),
                    "fault_attribution_no_fault" => Ok(FaultAttribution::NoFault),
                    "fault_attribution_inconclusive" => Ok(FaultAttribution::Inconclusive),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for FeeStatus {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "fee_status_unspecified",
            Self::Escrowed => "fee_status_escrowed",
            Self::ReleasedToAuditor => "fee_status_released_to_auditor",
            Self::ReturnedToProvider => "fee_status_returned_to_provider",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for FeeStatus {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "fee_status_unspecified",
            "fee_status_escrowed",
            "fee_status_released_to_auditor",
            "fee_status_returned_to_provider",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = FeeStatus;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                write!(formatter, "expected one of: {:?}", &FIELDS)
            }

            fn visit_i64<E>(self, v: i64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Signed(v), &self)
                    })
            }

            fn visit_u64<E>(self, v: u64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Unsigned(v), &self)
                    })
            }

            fn visit_str<E>(self, value: &str) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                match value {
                    "fee_status_unspecified" => Ok(FeeStatus::Unspecified),
                    "fee_status_escrowed" => Ok(FeeStatus::Escrowed),
                    "fee_status_released_to_auditor" => Ok(FeeStatus::ReleasedToAuditor),
                    "fee_status_returned_to_provider" => Ok(FeeStatus::ReturnedToProvider),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for GenesisState {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.params.is_some() {
            len += 1;
        }
        if !self.auditors.is_empty() {
            len += 1;
        }
        if !self.attestations.is_empty() {
            len += 1;
        }
        if !self.discrepancies.is_empty() {
            len += 1;
        }
        if !self.provider_bonds.is_empty() {
            len += 1;
        }
        if !self.provider_snapshots.is_empty() {
            len += 1;
        }
        if self.next_discrepancy_id != 0 {
            len += 1;
        }
        if !self.audit_escrows.is_empty() {
            len += 1;
        }
        if self.next_audit_escrow_id != 0 {
            len += 1;
        }
        if !self.verification_graces.is_empty() {
            len += 1;
        }
        if self.next_grace_record_id != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.GenesisState", len)?;
        if let Some(v) = self.params.as_ref() {
            struct_ser.serialize_field("params", v)?;
        }
        if !self.auditors.is_empty() {
            struct_ser.serialize_field("auditors", &self.auditors)?;
        }
        if !self.attestations.is_empty() {
            struct_ser.serialize_field("attestations", &self.attestations)?;
        }
        if !self.discrepancies.is_empty() {
            struct_ser.serialize_field("discrepancies", &self.discrepancies)?;
        }
        if !self.provider_bonds.is_empty() {
            struct_ser.serialize_field("providerBonds", &self.provider_bonds)?;
        }
        if !self.provider_snapshots.is_empty() {
            struct_ser.serialize_field("providerSnapshots", &self.provider_snapshots)?;
        }
        if self.next_discrepancy_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("nextDiscrepancyId", ToString::to_string(&self.next_discrepancy_id).as_str())?;
        }
        if !self.audit_escrows.is_empty() {
            struct_ser.serialize_field("auditEscrows", &self.audit_escrows)?;
        }
        if self.next_audit_escrow_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("nextAuditEscrowId", ToString::to_string(&self.next_audit_escrow_id).as_str())?;
        }
        if !self.verification_graces.is_empty() {
            struct_ser.serialize_field("verificationGraces", &self.verification_graces)?;
        }
        if self.next_grace_record_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("nextGraceRecordId", ToString::to_string(&self.next_grace_record_id).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for GenesisState {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "params",
            "auditors",
            "attestations",
            "discrepancies",
            "provider_bonds",
            "providerBonds",
            "provider_snapshots",
            "providerSnapshots",
            "next_discrepancy_id",
            "nextDiscrepancyId",
            "audit_escrows",
            "auditEscrows",
            "next_audit_escrow_id",
            "nextAuditEscrowId",
            "verification_graces",
            "verificationGraces",
            "next_grace_record_id",
            "nextGraceRecordId",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Params,
            Auditors,
            Attestations,
            Discrepancies,
            ProviderBonds,
            ProviderSnapshots,
            NextDiscrepancyId,
            AuditEscrows,
            NextAuditEscrowId,
            VerificationGraces,
            NextGraceRecordId,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "params" => Ok(GeneratedField::Params),
                            "auditors" => Ok(GeneratedField::Auditors),
                            "attestations" => Ok(GeneratedField::Attestations),
                            "discrepancies" => Ok(GeneratedField::Discrepancies),
                            "providerBonds" | "provider_bonds" => Ok(GeneratedField::ProviderBonds),
                            "providerSnapshots" | "provider_snapshots" => Ok(GeneratedField::ProviderSnapshots),
                            "nextDiscrepancyId" | "next_discrepancy_id" => Ok(GeneratedField::NextDiscrepancyId),
                            "auditEscrows" | "audit_escrows" => Ok(GeneratedField::AuditEscrows),
                            "nextAuditEscrowId" | "next_audit_escrow_id" => Ok(GeneratedField::NextAuditEscrowId),
                            "verificationGraces" | "verification_graces" => Ok(GeneratedField::VerificationGraces),
                            "nextGraceRecordId" | "next_grace_record_id" => Ok(GeneratedField::NextGraceRecordId),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = GenesisState;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.GenesisState")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<GenesisState, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut params__ = None;
                let mut auditors__ = None;
                let mut attestations__ = None;
                let mut discrepancies__ = None;
                let mut provider_bonds__ = None;
                let mut provider_snapshots__ = None;
                let mut next_discrepancy_id__ = None;
                let mut audit_escrows__ = None;
                let mut next_audit_escrow_id__ = None;
                let mut verification_graces__ = None;
                let mut next_grace_record_id__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Params => {
                            if params__.is_some() {
                                return Err(serde::de::Error::duplicate_field("params"));
                            }
                            params__ = map_.next_value()?;
                        }
                        GeneratedField::Auditors => {
                            if auditors__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditors"));
                            }
                            auditors__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Attestations => {
                            if attestations__.is_some() {
                                return Err(serde::de::Error::duplicate_field("attestations"));
                            }
                            attestations__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Discrepancies => {
                            if discrepancies__.is_some() {
                                return Err(serde::de::Error::duplicate_field("discrepancies"));
                            }
                            discrepancies__ = Some(map_.next_value()?);
                        }
                        GeneratedField::ProviderBonds => {
                            if provider_bonds__.is_some() {
                                return Err(serde::de::Error::duplicate_field("providerBonds"));
                            }
                            provider_bonds__ = Some(map_.next_value()?);
                        }
                        GeneratedField::ProviderSnapshots => {
                            if provider_snapshots__.is_some() {
                                return Err(serde::de::Error::duplicate_field("providerSnapshots"));
                            }
                            provider_snapshots__ = Some(map_.next_value()?);
                        }
                        GeneratedField::NextDiscrepancyId => {
                            if next_discrepancy_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("nextDiscrepancyId"));
                            }
                            next_discrepancy_id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::AuditEscrows => {
                            if audit_escrows__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditEscrows"));
                            }
                            audit_escrows__ = Some(map_.next_value()?);
                        }
                        GeneratedField::NextAuditEscrowId => {
                            if next_audit_escrow_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("nextAuditEscrowId"));
                            }
                            next_audit_escrow_id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::VerificationGraces => {
                            if verification_graces__.is_some() {
                                return Err(serde::de::Error::duplicate_field("verificationGraces"));
                            }
                            verification_graces__ = Some(map_.next_value()?);
                        }
                        GeneratedField::NextGraceRecordId => {
                            if next_grace_record_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("nextGraceRecordId"));
                            }
                            next_grace_record_id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(GenesisState {
                    params: params__,
                    auditors: auditors__.unwrap_or_default(),
                    attestations: attestations__.unwrap_or_default(),
                    discrepancies: discrepancies__.unwrap_or_default(),
                    provider_bonds: provider_bonds__.unwrap_or_default(),
                    provider_snapshots: provider_snapshots__.unwrap_or_default(),
                    next_discrepancy_id: next_discrepancy_id__.unwrap_or_default(),
                    audit_escrows: audit_escrows__.unwrap_or_default(),
                    next_audit_escrow_id: next_audit_escrow_id__.unwrap_or_default(),
                    verification_graces: verification_graces__.unwrap_or_default(),
                    next_grace_record_id: next_grace_record_id__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.GenesisState", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for GovernanceAttestationReason {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "governance_attestation_reason_unspecified",
            Self::FraudulentProvider => "governance_attestation_reason_fraudulent_provider",
            Self::CompromisedProvider => "governance_attestation_reason_compromised_provider",
            Self::ProviderNonCooperation => "governance_attestation_reason_provider_non_cooperation",
            Self::FaultyAuditor => "governance_attestation_reason_faulty_auditor",
            Self::NegligentAuditor => "governance_attestation_reason_negligent_auditor",
            Self::EvidenceInsufficient => "governance_attestation_reason_evidence_insufficient",
            Self::EmergencySafetyAction => "governance_attestation_reason_emergency_safety_action",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for GovernanceAttestationReason {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "governance_attestation_reason_unspecified",
            "governance_attestation_reason_fraudulent_provider",
            "governance_attestation_reason_compromised_provider",
            "governance_attestation_reason_provider_non_cooperation",
            "governance_attestation_reason_faulty_auditor",
            "governance_attestation_reason_negligent_auditor",
            "governance_attestation_reason_evidence_insufficient",
            "governance_attestation_reason_emergency_safety_action",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = GovernanceAttestationReason;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                write!(formatter, "expected one of: {:?}", &FIELDS)
            }

            fn visit_i64<E>(self, v: i64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Signed(v), &self)
                    })
            }

            fn visit_u64<E>(self, v: u64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Unsigned(v), &self)
                    })
            }

            fn visit_str<E>(self, value: &str) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                match value {
                    "governance_attestation_reason_unspecified" => Ok(GovernanceAttestationReason::Unspecified),
                    "governance_attestation_reason_fraudulent_provider" => Ok(GovernanceAttestationReason::FraudulentProvider),
                    "governance_attestation_reason_compromised_provider" => Ok(GovernanceAttestationReason::CompromisedProvider),
                    "governance_attestation_reason_provider_non_cooperation" => Ok(GovernanceAttestationReason::ProviderNonCooperation),
                    "governance_attestation_reason_faulty_auditor" => Ok(GovernanceAttestationReason::FaultyAuditor),
                    "governance_attestation_reason_negligent_auditor" => Ok(GovernanceAttestationReason::NegligentAuditor),
                    "governance_attestation_reason_evidence_insufficient" => Ok(GovernanceAttestationReason::EvidenceInsufficient),
                    "governance_attestation_reason_emergency_safety_action" => Ok(GovernanceAttestationReason::EmergencySafetyAction),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for MsgCancelAuditEscrow {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if self.audit_escrow_id != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.MsgCancelAuditEscrow", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if self.audit_escrow_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("auditEscrowId", ToString::to_string(&self.audit_escrow_id).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgCancelAuditEscrow {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "audit_escrow_id",
            "auditEscrowId",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            AuditEscrowId,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "auditEscrowId" | "audit_escrow_id" => Ok(GeneratedField::AuditEscrowId),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgCancelAuditEscrow;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgCancelAuditEscrow")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgCancelAuditEscrow, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut audit_escrow_id__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::AuditEscrowId => {
                            if audit_escrow_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditEscrowId"));
                            }
                            audit_escrow_id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(MsgCancelAuditEscrow {
                    provider: provider__.unwrap_or_default(),
                    audit_escrow_id: audit_escrow_id__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgCancelAuditEscrow", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgCancelAuditEscrowResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.verification.v1.MsgCancelAuditEscrowResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgCancelAuditEscrowResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                            Err(serde::de::Error::unknown_field(value, FIELDS))
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgCancelAuditEscrowResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgCancelAuditEscrowResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgCancelAuditEscrowResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgCancelAuditEscrowResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgCancelAuditEscrowResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgOpenAuditEscrow {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if self.requested_tier != 0 {
            len += 1;
        }
        if !self.requested_capabilities.is_empty() {
            len += 1;
        }
        if self.fee.is_some() {
            len += 1;
        }
        if self.provider_deposit.is_some() {
            len += 1;
        }
        if self.expires_at.is_some() {
            len += 1;
        }
        if !self.metadata_hash.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.MsgOpenAuditEscrow", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if self.requested_tier != 0 {
            let v = VerificationTier::try_from(self.requested_tier)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.requested_tier)))?;
            struct_ser.serialize_field("requestedTier", &v)?;
        }
        if !self.requested_capabilities.is_empty() {
            let v = self.requested_capabilities.iter().cloned().map(|v| {
                CapabilityFlag::try_from(v)
                    .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", v)))
                }).collect::<std::result::Result<Vec<_>, _>>()?;
            struct_ser.serialize_field("requestedCapabilities", &v)?;
        }
        if let Some(v) = self.fee.as_ref() {
            struct_ser.serialize_field("fee", v)?;
        }
        if let Some(v) = self.provider_deposit.as_ref() {
            struct_ser.serialize_field("providerDeposit", v)?;
        }
        if let Some(v) = self.expires_at.as_ref() {
            struct_ser.serialize_field("expiresAt", v)?;
        }
        if !self.metadata_hash.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("metadataHash", pbjson::private::base64::encode(&self.metadata_hash).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgOpenAuditEscrow {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "requested_tier",
            "requestedTier",
            "requested_capabilities",
            "requestedCapabilities",
            "fee",
            "provider_deposit",
            "providerDeposit",
            "expires_at",
            "expiresAt",
            "metadata_hash",
            "metadataHash",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            RequestedTier,
            RequestedCapabilities,
            Fee,
            ProviderDeposit,
            ExpiresAt,
            MetadataHash,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "requestedTier" | "requested_tier" => Ok(GeneratedField::RequestedTier),
                            "requestedCapabilities" | "requested_capabilities" => Ok(GeneratedField::RequestedCapabilities),
                            "fee" => Ok(GeneratedField::Fee),
                            "providerDeposit" | "provider_deposit" => Ok(GeneratedField::ProviderDeposit),
                            "expiresAt" | "expires_at" => Ok(GeneratedField::ExpiresAt),
                            "metadataHash" | "metadata_hash" => Ok(GeneratedField::MetadataHash),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgOpenAuditEscrow;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgOpenAuditEscrow")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgOpenAuditEscrow, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut requested_tier__ = None;
                let mut requested_capabilities__ = None;
                let mut fee__ = None;
                let mut provider_deposit__ = None;
                let mut expires_at__ = None;
                let mut metadata_hash__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::RequestedTier => {
                            if requested_tier__.is_some() {
                                return Err(serde::de::Error::duplicate_field("requestedTier"));
                            }
                            requested_tier__ = Some(map_.next_value::<VerificationTier>()? as i32);
                        }
                        GeneratedField::RequestedCapabilities => {
                            if requested_capabilities__.is_some() {
                                return Err(serde::de::Error::duplicate_field("requestedCapabilities"));
                            }
                            requested_capabilities__ = Some(map_.next_value::<Vec<CapabilityFlag>>()?.into_iter().map(|x| x as i32).collect());
                        }
                        GeneratedField::Fee => {
                            if fee__.is_some() {
                                return Err(serde::de::Error::duplicate_field("fee"));
                            }
                            fee__ = map_.next_value()?;
                        }
                        GeneratedField::ProviderDeposit => {
                            if provider_deposit__.is_some() {
                                return Err(serde::de::Error::duplicate_field("providerDeposit"));
                            }
                            provider_deposit__ = map_.next_value()?;
                        }
                        GeneratedField::ExpiresAt => {
                            if expires_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("expiresAt"));
                            }
                            expires_at__ = map_.next_value()?;
                        }
                        GeneratedField::MetadataHash => {
                            if metadata_hash__.is_some() {
                                return Err(serde::de::Error::duplicate_field("metadataHash"));
                            }
                            metadata_hash__ = 
                                Some(map_.next_value::<::pbjson::private::BytesDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(MsgOpenAuditEscrow {
                    provider: provider__.unwrap_or_default(),
                    requested_tier: requested_tier__.unwrap_or_default(),
                    requested_capabilities: requested_capabilities__.unwrap_or_default(),
                    fee: fee__,
                    provider_deposit: provider_deposit__,
                    expires_at: expires_at__,
                    metadata_hash: metadata_hash__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgOpenAuditEscrow", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgOpenAuditEscrowResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.audit_escrow_id != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.MsgOpenAuditEscrowResponse", len)?;
        if self.audit_escrow_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("auditEscrowId", ToString::to_string(&self.audit_escrow_id).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgOpenAuditEscrowResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "audit_escrow_id",
            "auditEscrowId",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            AuditEscrowId,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "auditEscrowId" | "audit_escrow_id" => Ok(GeneratedField::AuditEscrowId),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgOpenAuditEscrowResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgOpenAuditEscrowResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgOpenAuditEscrowResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut audit_escrow_id__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::AuditEscrowId => {
                            if audit_escrow_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditEscrowId"));
                            }
                            audit_escrow_id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(MsgOpenAuditEscrowResponse {
                    audit_escrow_id: audit_escrow_id__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgOpenAuditEscrowResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgPostAuditorBond {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.auditor.is_empty() {
            len += 1;
        }
        if self.amount.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.MsgPostAuditorBond", len)?;
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        if let Some(v) = self.amount.as_ref() {
            struct_ser.serialize_field("amount", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgPostAuditorBond {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "auditor",
            "amount",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Auditor,
            Amount,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "auditor" => Ok(GeneratedField::Auditor),
                            "amount" => Ok(GeneratedField::Amount),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgPostAuditorBond;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgPostAuditorBond")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgPostAuditorBond, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut auditor__ = None;
                let mut amount__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Amount => {
                            if amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("amount"));
                            }
                            amount__ = map_.next_value()?;
                        }
                    }
                }
                Ok(MsgPostAuditorBond {
                    auditor: auditor__.unwrap_or_default(),
                    amount: amount__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgPostAuditorBond", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgPostAuditorBondResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.verification.v1.MsgPostAuditorBondResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgPostAuditorBondResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                            Err(serde::de::Error::unknown_field(value, FIELDS))
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgPostAuditorBondResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgPostAuditorBondResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgPostAuditorBondResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgPostAuditorBondResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgPostAuditorBondResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgPostProviderBond {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if self.amount.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.MsgPostProviderBond", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if let Some(v) = self.amount.as_ref() {
            struct_ser.serialize_field("amount", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgPostProviderBond {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "amount",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            Amount,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "amount" => Ok(GeneratedField::Amount),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgPostProviderBond;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgPostProviderBond")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgPostProviderBond, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut amount__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Amount => {
                            if amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("amount"));
                            }
                            amount__ = map_.next_value()?;
                        }
                    }
                }
                Ok(MsgPostProviderBond {
                    provider: provider__.unwrap_or_default(),
                    amount: amount__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgPostProviderBond", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgPostProviderBondResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.verification.v1.MsgPostProviderBondResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgPostProviderBondResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                            Err(serde::de::Error::unknown_field(value, FIELDS))
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgPostProviderBondResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgPostProviderBondResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgPostProviderBondResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgPostProviderBondResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgPostProviderBondResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgPostSnapshotHash {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if !self.snapshot_hash.is_empty() {
            len += 1;
        }
        if self.resource_summary.is_some() {
            len += 1;
        }
        if self.snapshot_timestamp.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.MsgPostSnapshotHash", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if !self.snapshot_hash.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("snapshotHash", pbjson::private::base64::encode(&self.snapshot_hash).as_str())?;
        }
        if let Some(v) = self.resource_summary.as_ref() {
            struct_ser.serialize_field("resourceSummary", v)?;
        }
        if let Some(v) = self.snapshot_timestamp.as_ref() {
            struct_ser.serialize_field("snapshotTimestamp", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgPostSnapshotHash {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "snapshot_hash",
            "snapshotHash",
            "resource_summary",
            "resourceSummary",
            "snapshot_timestamp",
            "snapshotTimestamp",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            SnapshotHash,
            ResourceSummary,
            SnapshotTimestamp,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "snapshotHash" | "snapshot_hash" => Ok(GeneratedField::SnapshotHash),
                            "resourceSummary" | "resource_summary" => Ok(GeneratedField::ResourceSummary),
                            "snapshotTimestamp" | "snapshot_timestamp" => Ok(GeneratedField::SnapshotTimestamp),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgPostSnapshotHash;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgPostSnapshotHash")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgPostSnapshotHash, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut snapshot_hash__ = None;
                let mut resource_summary__ = None;
                let mut snapshot_timestamp__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::SnapshotHash => {
                            if snapshot_hash__.is_some() {
                                return Err(serde::de::Error::duplicate_field("snapshotHash"));
                            }
                            snapshot_hash__ = 
                                Some(map_.next_value::<::pbjson::private::BytesDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::ResourceSummary => {
                            if resource_summary__.is_some() {
                                return Err(serde::de::Error::duplicate_field("resourceSummary"));
                            }
                            resource_summary__ = map_.next_value()?;
                        }
                        GeneratedField::SnapshotTimestamp => {
                            if snapshot_timestamp__.is_some() {
                                return Err(serde::de::Error::duplicate_field("snapshotTimestamp"));
                            }
                            snapshot_timestamp__ = map_.next_value()?;
                        }
                    }
                }
                Ok(MsgPostSnapshotHash {
                    provider: provider__.unwrap_or_default(),
                    snapshot_hash: snapshot_hash__.unwrap_or_default(),
                    resource_summary: resource_summary__,
                    snapshot_timestamp: snapshot_timestamp__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgPostSnapshotHash", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgPostSnapshotHashResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.verification.v1.MsgPostSnapshotHashResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgPostSnapshotHashResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                            Err(serde::de::Error::unknown_field(value, FIELDS))
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgPostSnapshotHashResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgPostSnapshotHashResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgPostSnapshotHashResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgPostSnapshotHashResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgPostSnapshotHashResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgRegisterAuditor {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.authority.is_empty() {
            len += 1;
        }
        if !self.auditor.is_empty() {
            len += 1;
        }
        if self.max_attestation_tier != 0 {
            len += 1;
        }
        if !self.metadata_hash.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.MsgRegisterAuditor", len)?;
        if !self.authority.is_empty() {
            struct_ser.serialize_field("authority", &self.authority)?;
        }
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        if self.max_attestation_tier != 0 {
            let v = VerificationTier::try_from(self.max_attestation_tier)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.max_attestation_tier)))?;
            struct_ser.serialize_field("maxAttestationTier", &v)?;
        }
        if !self.metadata_hash.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("metadataHash", pbjson::private::base64::encode(&self.metadata_hash).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgRegisterAuditor {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "authority",
            "auditor",
            "max_attestation_tier",
            "maxAttestationTier",
            "metadata_hash",
            "metadataHash",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Authority,
            Auditor,
            MaxAttestationTier,
            MetadataHash,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "authority" => Ok(GeneratedField::Authority),
                            "auditor" => Ok(GeneratedField::Auditor),
                            "maxAttestationTier" | "max_attestation_tier" => Ok(GeneratedField::MaxAttestationTier),
                            "metadataHash" | "metadata_hash" => Ok(GeneratedField::MetadataHash),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgRegisterAuditor;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgRegisterAuditor")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgRegisterAuditor, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut authority__ = None;
                let mut auditor__ = None;
                let mut max_attestation_tier__ = None;
                let mut metadata_hash__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Authority => {
                            if authority__.is_some() {
                                return Err(serde::de::Error::duplicate_field("authority"));
                            }
                            authority__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::MaxAttestationTier => {
                            if max_attestation_tier__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maxAttestationTier"));
                            }
                            max_attestation_tier__ = Some(map_.next_value::<VerificationTier>()? as i32);
                        }
                        GeneratedField::MetadataHash => {
                            if metadata_hash__.is_some() {
                                return Err(serde::de::Error::duplicate_field("metadataHash"));
                            }
                            metadata_hash__ = 
                                Some(map_.next_value::<::pbjson::private::BytesDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(MsgRegisterAuditor {
                    authority: authority__.unwrap_or_default(),
                    auditor: auditor__.unwrap_or_default(),
                    max_attestation_tier: max_attestation_tier__.unwrap_or_default(),
                    metadata_hash: metadata_hash__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgRegisterAuditor", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgRegisterAuditorResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.verification.v1.MsgRegisterAuditorResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgRegisterAuditorResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                            Err(serde::de::Error::unknown_field(value, FIELDS))
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgRegisterAuditorResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgRegisterAuditorResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgRegisterAuditorResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgRegisterAuditorResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgRegisterAuditorResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgRemoveAttestation {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if !self.auditor.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.MsgRemoveAttestation", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgRemoveAttestation {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "auditor",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            Auditor,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "auditor" => Ok(GeneratedField::Auditor),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgRemoveAttestation;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgRemoveAttestation")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgRemoveAttestation, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut auditor__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(MsgRemoveAttestation {
                    provider: provider__.unwrap_or_default(),
                    auditor: auditor__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgRemoveAttestation", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgRemoveAttestationResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.verification.v1.MsgRemoveAttestationResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgRemoveAttestationResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                            Err(serde::de::Error::unknown_field(value, FIELDS))
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgRemoveAttestationResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgRemoveAttestationResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgRemoveAttestationResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgRemoveAttestationResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgRemoveAttestationResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgRemoveAuditor {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.authority.is_empty() {
            len += 1;
        }
        if !self.auditor.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.MsgRemoveAuditor", len)?;
        if !self.authority.is_empty() {
            struct_ser.serialize_field("authority", &self.authority)?;
        }
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgRemoveAuditor {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "authority",
            "auditor",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Authority,
            Auditor,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "authority" => Ok(GeneratedField::Authority),
                            "auditor" => Ok(GeneratedField::Auditor),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgRemoveAuditor;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgRemoveAuditor")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgRemoveAuditor, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut authority__ = None;
                let mut auditor__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Authority => {
                            if authority__.is_some() {
                                return Err(serde::de::Error::duplicate_field("authority"));
                            }
                            authority__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(MsgRemoveAuditor {
                    authority: authority__.unwrap_or_default(),
                    auditor: auditor__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgRemoveAuditor", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgRemoveAuditorResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.verification.v1.MsgRemoveAuditorResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgRemoveAuditorResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                            Err(serde::de::Error::unknown_field(value, FIELDS))
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgRemoveAuditorResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgRemoveAuditorResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgRemoveAuditorResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgRemoveAuditorResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgRemoveAuditorResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgRenewAuditor {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.authority.is_empty() {
            len += 1;
        }
        if !self.auditor.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.MsgRenewAuditor", len)?;
        if !self.authority.is_empty() {
            struct_ser.serialize_field("authority", &self.authority)?;
        }
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgRenewAuditor {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "authority",
            "auditor",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Authority,
            Auditor,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "authority" => Ok(GeneratedField::Authority),
                            "auditor" => Ok(GeneratedField::Auditor),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgRenewAuditor;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgRenewAuditor")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgRenewAuditor, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut authority__ = None;
                let mut auditor__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Authority => {
                            if authority__.is_some() {
                                return Err(serde::de::Error::duplicate_field("authority"));
                            }
                            authority__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(MsgRenewAuditor {
                    authority: authority__.unwrap_or_default(),
                    auditor: auditor__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgRenewAuditor", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgRenewAuditorResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.verification.v1.MsgRenewAuditorResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgRenewAuditorResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                            Err(serde::de::Error::unknown_field(value, FIELDS))
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgRenewAuditorResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgRenewAuditorResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgRenewAuditorResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgRenewAuditorResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgRenewAuditorResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgResignAuditor {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.auditor.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.MsgResignAuditor", len)?;
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgResignAuditor {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "auditor",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Auditor,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "auditor" => Ok(GeneratedField::Auditor),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgResignAuditor;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgResignAuditor")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgResignAuditor, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut auditor__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(MsgResignAuditor {
                    auditor: auditor__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgResignAuditor", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgResignAuditorResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.verification.v1.MsgResignAuditorResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgResignAuditorResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                            Err(serde::de::Error::unknown_field(value, FIELDS))
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgResignAuditorResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgResignAuditorResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgResignAuditorResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgResignAuditorResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgResignAuditorResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgResolveDiscrepancy {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.authority.is_empty() {
            len += 1;
        }
        if self.discrepancy_id != 0 {
            len += 1;
        }
        if !self.vindicated_auditor.is_empty() {
            len += 1;
        }
        if self.slash_auditor_a {
            len += 1;
        }
        if self.slash_auditor_b {
            len += 1;
        }
        if self.reason != 0 {
            len += 1;
        }
        if self.fault_attribution != 0 {
            len += 1;
        }
        if !self.evidence_hash.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.MsgResolveDiscrepancy", len)?;
        if !self.authority.is_empty() {
            struct_ser.serialize_field("authority", &self.authority)?;
        }
        if self.discrepancy_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("discrepancyId", ToString::to_string(&self.discrepancy_id).as_str())?;
        }
        if !self.vindicated_auditor.is_empty() {
            struct_ser.serialize_field("vindicatedAuditor", &self.vindicated_auditor)?;
        }
        if self.slash_auditor_a {
            struct_ser.serialize_field("slashAuditorA", &self.slash_auditor_a)?;
        }
        if self.slash_auditor_b {
            struct_ser.serialize_field("slashAuditorB", &self.slash_auditor_b)?;
        }
        if self.reason != 0 {
            let v = DiscrepancyResolutionReason::try_from(self.reason)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.reason)))?;
            struct_ser.serialize_field("reason", &v)?;
        }
        if self.fault_attribution != 0 {
            let v = FaultAttribution::try_from(self.fault_attribution)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.fault_attribution)))?;
            struct_ser.serialize_field("faultAttribution", &v)?;
        }
        if !self.evidence_hash.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("evidenceHash", pbjson::private::base64::encode(&self.evidence_hash).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgResolveDiscrepancy {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "authority",
            "discrepancy_id",
            "discrepancyId",
            "vindicated_auditor",
            "vindicatedAuditor",
            "slash_auditor_a",
            "slashAuditorA",
            "slash_auditor_b",
            "slashAuditorB",
            "reason",
            "fault_attribution",
            "faultAttribution",
            "evidence_hash",
            "evidenceHash",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Authority,
            DiscrepancyId,
            VindicatedAuditor,
            SlashAuditorA,
            SlashAuditorB,
            Reason,
            FaultAttribution,
            EvidenceHash,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "authority" => Ok(GeneratedField::Authority),
                            "discrepancyId" | "discrepancy_id" => Ok(GeneratedField::DiscrepancyId),
                            "vindicatedAuditor" | "vindicated_auditor" => Ok(GeneratedField::VindicatedAuditor),
                            "slashAuditorA" | "slash_auditor_a" => Ok(GeneratedField::SlashAuditorA),
                            "slashAuditorB" | "slash_auditor_b" => Ok(GeneratedField::SlashAuditorB),
                            "reason" => Ok(GeneratedField::Reason),
                            "faultAttribution" | "fault_attribution" => Ok(GeneratedField::FaultAttribution),
                            "evidenceHash" | "evidence_hash" => Ok(GeneratedField::EvidenceHash),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgResolveDiscrepancy;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgResolveDiscrepancy")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgResolveDiscrepancy, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut authority__ = None;
                let mut discrepancy_id__ = None;
                let mut vindicated_auditor__ = None;
                let mut slash_auditor_a__ = None;
                let mut slash_auditor_b__ = None;
                let mut reason__ = None;
                let mut fault_attribution__ = None;
                let mut evidence_hash__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Authority => {
                            if authority__.is_some() {
                                return Err(serde::de::Error::duplicate_field("authority"));
                            }
                            authority__ = Some(map_.next_value()?);
                        }
                        GeneratedField::DiscrepancyId => {
                            if discrepancy_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("discrepancyId"));
                            }
                            discrepancy_id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::VindicatedAuditor => {
                            if vindicated_auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("vindicatedAuditor"));
                            }
                            vindicated_auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::SlashAuditorA => {
                            if slash_auditor_a__.is_some() {
                                return Err(serde::de::Error::duplicate_field("slashAuditorA"));
                            }
                            slash_auditor_a__ = Some(map_.next_value()?);
                        }
                        GeneratedField::SlashAuditorB => {
                            if slash_auditor_b__.is_some() {
                                return Err(serde::de::Error::duplicate_field("slashAuditorB"));
                            }
                            slash_auditor_b__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Reason => {
                            if reason__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reason"));
                            }
                            reason__ = Some(map_.next_value::<DiscrepancyResolutionReason>()? as i32);
                        }
                        GeneratedField::FaultAttribution => {
                            if fault_attribution__.is_some() {
                                return Err(serde::de::Error::duplicate_field("faultAttribution"));
                            }
                            fault_attribution__ = Some(map_.next_value::<FaultAttribution>()? as i32);
                        }
                        GeneratedField::EvidenceHash => {
                            if evidence_hash__.is_some() {
                                return Err(serde::de::Error::duplicate_field("evidenceHash"));
                            }
                            evidence_hash__ = 
                                Some(map_.next_value::<::pbjson::private::BytesDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(MsgResolveDiscrepancy {
                    authority: authority__.unwrap_or_default(),
                    discrepancy_id: discrepancy_id__.unwrap_or_default(),
                    vindicated_auditor: vindicated_auditor__.unwrap_or_default(),
                    slash_auditor_a: slash_auditor_a__.unwrap_or_default(),
                    slash_auditor_b: slash_auditor_b__.unwrap_or_default(),
                    reason: reason__.unwrap_or_default(),
                    fault_attribution: fault_attribution__.unwrap_or_default(),
                    evidence_hash: evidence_hash__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgResolveDiscrepancy", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgResolveDiscrepancyResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.verification.v1.MsgResolveDiscrepancyResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgResolveDiscrepancyResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                            Err(serde::de::Error::unknown_field(value, FIELDS))
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgResolveDiscrepancyResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgResolveDiscrepancyResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgResolveDiscrepancyResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgResolveDiscrepancyResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgResolveDiscrepancyResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgRevokeAllProviderAttestations {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.authority.is_empty() {
            len += 1;
        }
        if !self.provider.is_empty() {
            len += 1;
        }
        if self.reason != 0 {
            len += 1;
        }
        if self.fault_attribution != 0 {
            len += 1;
        }
        if !self.evidence_hash.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.MsgRevokeAllProviderAttestations", len)?;
        if !self.authority.is_empty() {
            struct_ser.serialize_field("authority", &self.authority)?;
        }
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if self.reason != 0 {
            let v = GovernanceAttestationReason::try_from(self.reason)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.reason)))?;
            struct_ser.serialize_field("reason", &v)?;
        }
        if self.fault_attribution != 0 {
            let v = FaultAttribution::try_from(self.fault_attribution)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.fault_attribution)))?;
            struct_ser.serialize_field("faultAttribution", &v)?;
        }
        if !self.evidence_hash.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("evidenceHash", pbjson::private::base64::encode(&self.evidence_hash).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgRevokeAllProviderAttestations {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "authority",
            "provider",
            "reason",
            "fault_attribution",
            "faultAttribution",
            "evidence_hash",
            "evidenceHash",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Authority,
            Provider,
            Reason,
            FaultAttribution,
            EvidenceHash,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "authority" => Ok(GeneratedField::Authority),
                            "provider" => Ok(GeneratedField::Provider),
                            "reason" => Ok(GeneratedField::Reason),
                            "faultAttribution" | "fault_attribution" => Ok(GeneratedField::FaultAttribution),
                            "evidenceHash" | "evidence_hash" => Ok(GeneratedField::EvidenceHash),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgRevokeAllProviderAttestations;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgRevokeAllProviderAttestations")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgRevokeAllProviderAttestations, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut authority__ = None;
                let mut provider__ = None;
                let mut reason__ = None;
                let mut fault_attribution__ = None;
                let mut evidence_hash__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Authority => {
                            if authority__.is_some() {
                                return Err(serde::de::Error::duplicate_field("authority"));
                            }
                            authority__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Reason => {
                            if reason__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reason"));
                            }
                            reason__ = Some(map_.next_value::<GovernanceAttestationReason>()? as i32);
                        }
                        GeneratedField::FaultAttribution => {
                            if fault_attribution__.is_some() {
                                return Err(serde::de::Error::duplicate_field("faultAttribution"));
                            }
                            fault_attribution__ = Some(map_.next_value::<FaultAttribution>()? as i32);
                        }
                        GeneratedField::EvidenceHash => {
                            if evidence_hash__.is_some() {
                                return Err(serde::de::Error::duplicate_field("evidenceHash"));
                            }
                            evidence_hash__ = 
                                Some(map_.next_value::<::pbjson::private::BytesDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(MsgRevokeAllProviderAttestations {
                    authority: authority__.unwrap_or_default(),
                    provider: provider__.unwrap_or_default(),
                    reason: reason__.unwrap_or_default(),
                    fault_attribution: fault_attribution__.unwrap_or_default(),
                    evidence_hash: evidence_hash__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgRevokeAllProviderAttestations", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgRevokeAllProviderAttestationsResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.verification.v1.MsgRevokeAllProviderAttestationsResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgRevokeAllProviderAttestationsResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                            Err(serde::de::Error::unknown_field(value, FIELDS))
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgRevokeAllProviderAttestationsResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgRevokeAllProviderAttestationsResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgRevokeAllProviderAttestationsResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgRevokeAllProviderAttestationsResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgRevokeAllProviderAttestationsResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgRevokeAttestation {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if !self.auditor.is_empty() {
            len += 1;
        }
        if self.reason != 0 {
            len += 1;
        }
        if !self.evidence_hash.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.MsgRevokeAttestation", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        if self.reason != 0 {
            let v = AttestationRevocationReason::try_from(self.reason)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.reason)))?;
            struct_ser.serialize_field("reason", &v)?;
        }
        if !self.evidence_hash.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("evidenceHash", pbjson::private::base64::encode(&self.evidence_hash).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgRevokeAttestation {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "auditor",
            "reason",
            "evidence_hash",
            "evidenceHash",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            Auditor,
            Reason,
            EvidenceHash,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "auditor" => Ok(GeneratedField::Auditor),
                            "reason" => Ok(GeneratedField::Reason),
                            "evidenceHash" | "evidence_hash" => Ok(GeneratedField::EvidenceHash),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgRevokeAttestation;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgRevokeAttestation")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgRevokeAttestation, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut auditor__ = None;
                let mut reason__ = None;
                let mut evidence_hash__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Reason => {
                            if reason__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reason"));
                            }
                            reason__ = Some(map_.next_value::<AttestationRevocationReason>()? as i32);
                        }
                        GeneratedField::EvidenceHash => {
                            if evidence_hash__.is_some() {
                                return Err(serde::de::Error::duplicate_field("evidenceHash"));
                            }
                            evidence_hash__ = 
                                Some(map_.next_value::<::pbjson::private::BytesDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(MsgRevokeAttestation {
                    provider: provider__.unwrap_or_default(),
                    auditor: auditor__.unwrap_or_default(),
                    reason: reason__.unwrap_or_default(),
                    evidence_hash: evidence_hash__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgRevokeAttestation", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgRevokeAttestationResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.verification.v1.MsgRevokeAttestationResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgRevokeAttestationResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                            Err(serde::de::Error::unknown_field(value, FIELDS))
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgRevokeAttestationResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgRevokeAttestationResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgRevokeAttestationResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgRevokeAttestationResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgRevokeAttestationResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgRevokeAuditorAttestations {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.authority.is_empty() {
            len += 1;
        }
        if !self.auditor.is_empty() {
            len += 1;
        }
        if self.reason != 0 {
            len += 1;
        }
        if self.fault_attribution != 0 {
            len += 1;
        }
        if !self.evidence_hash.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.MsgRevokeAuditorAttestations", len)?;
        if !self.authority.is_empty() {
            struct_ser.serialize_field("authority", &self.authority)?;
        }
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        if self.reason != 0 {
            let v = GovernanceAttestationReason::try_from(self.reason)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.reason)))?;
            struct_ser.serialize_field("reason", &v)?;
        }
        if self.fault_attribution != 0 {
            let v = FaultAttribution::try_from(self.fault_attribution)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.fault_attribution)))?;
            struct_ser.serialize_field("faultAttribution", &v)?;
        }
        if !self.evidence_hash.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("evidenceHash", pbjson::private::base64::encode(&self.evidence_hash).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgRevokeAuditorAttestations {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "authority",
            "auditor",
            "reason",
            "fault_attribution",
            "faultAttribution",
            "evidence_hash",
            "evidenceHash",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Authority,
            Auditor,
            Reason,
            FaultAttribution,
            EvidenceHash,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "authority" => Ok(GeneratedField::Authority),
                            "auditor" => Ok(GeneratedField::Auditor),
                            "reason" => Ok(GeneratedField::Reason),
                            "faultAttribution" | "fault_attribution" => Ok(GeneratedField::FaultAttribution),
                            "evidenceHash" | "evidence_hash" => Ok(GeneratedField::EvidenceHash),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgRevokeAuditorAttestations;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgRevokeAuditorAttestations")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgRevokeAuditorAttestations, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut authority__ = None;
                let mut auditor__ = None;
                let mut reason__ = None;
                let mut fault_attribution__ = None;
                let mut evidence_hash__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Authority => {
                            if authority__.is_some() {
                                return Err(serde::de::Error::duplicate_field("authority"));
                            }
                            authority__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Reason => {
                            if reason__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reason"));
                            }
                            reason__ = Some(map_.next_value::<GovernanceAttestationReason>()? as i32);
                        }
                        GeneratedField::FaultAttribution => {
                            if fault_attribution__.is_some() {
                                return Err(serde::de::Error::duplicate_field("faultAttribution"));
                            }
                            fault_attribution__ = Some(map_.next_value::<FaultAttribution>()? as i32);
                        }
                        GeneratedField::EvidenceHash => {
                            if evidence_hash__.is_some() {
                                return Err(serde::de::Error::duplicate_field("evidenceHash"));
                            }
                            evidence_hash__ = 
                                Some(map_.next_value::<::pbjson::private::BytesDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(MsgRevokeAuditorAttestations {
                    authority: authority__.unwrap_or_default(),
                    auditor: auditor__.unwrap_or_default(),
                    reason: reason__.unwrap_or_default(),
                    fault_attribution: fault_attribution__.unwrap_or_default(),
                    evidence_hash: evidence_hash__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgRevokeAuditorAttestations", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgRevokeAuditorAttestationsResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.verification.v1.MsgRevokeAuditorAttestationsResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgRevokeAuditorAttestationsResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                            Err(serde::de::Error::unknown_field(value, FIELDS))
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgRevokeAuditorAttestationsResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgRevokeAuditorAttestationsResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgRevokeAuditorAttestationsResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgRevokeAuditorAttestationsResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgRevokeAuditorAttestationsResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgRevokeProviderAttestation {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.authority.is_empty() {
            len += 1;
        }
        if !self.provider.is_empty() {
            len += 1;
        }
        if !self.auditor.is_empty() {
            len += 1;
        }
        if self.reason != 0 {
            len += 1;
        }
        if self.fault_attribution != 0 {
            len += 1;
        }
        if !self.evidence_hash.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.MsgRevokeProviderAttestation", len)?;
        if !self.authority.is_empty() {
            struct_ser.serialize_field("authority", &self.authority)?;
        }
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        if self.reason != 0 {
            let v = GovernanceAttestationReason::try_from(self.reason)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.reason)))?;
            struct_ser.serialize_field("reason", &v)?;
        }
        if self.fault_attribution != 0 {
            let v = FaultAttribution::try_from(self.fault_attribution)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.fault_attribution)))?;
            struct_ser.serialize_field("faultAttribution", &v)?;
        }
        if !self.evidence_hash.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("evidenceHash", pbjson::private::base64::encode(&self.evidence_hash).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgRevokeProviderAttestation {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "authority",
            "provider",
            "auditor",
            "reason",
            "fault_attribution",
            "faultAttribution",
            "evidence_hash",
            "evidenceHash",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Authority,
            Provider,
            Auditor,
            Reason,
            FaultAttribution,
            EvidenceHash,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "authority" => Ok(GeneratedField::Authority),
                            "provider" => Ok(GeneratedField::Provider),
                            "auditor" => Ok(GeneratedField::Auditor),
                            "reason" => Ok(GeneratedField::Reason),
                            "faultAttribution" | "fault_attribution" => Ok(GeneratedField::FaultAttribution),
                            "evidenceHash" | "evidence_hash" => Ok(GeneratedField::EvidenceHash),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgRevokeProviderAttestation;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgRevokeProviderAttestation")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgRevokeProviderAttestation, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut authority__ = None;
                let mut provider__ = None;
                let mut auditor__ = None;
                let mut reason__ = None;
                let mut fault_attribution__ = None;
                let mut evidence_hash__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Authority => {
                            if authority__.is_some() {
                                return Err(serde::de::Error::duplicate_field("authority"));
                            }
                            authority__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Reason => {
                            if reason__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reason"));
                            }
                            reason__ = Some(map_.next_value::<GovernanceAttestationReason>()? as i32);
                        }
                        GeneratedField::FaultAttribution => {
                            if fault_attribution__.is_some() {
                                return Err(serde::de::Error::duplicate_field("faultAttribution"));
                            }
                            fault_attribution__ = Some(map_.next_value::<FaultAttribution>()? as i32);
                        }
                        GeneratedField::EvidenceHash => {
                            if evidence_hash__.is_some() {
                                return Err(serde::de::Error::duplicate_field("evidenceHash"));
                            }
                            evidence_hash__ = 
                                Some(map_.next_value::<::pbjson::private::BytesDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(MsgRevokeProviderAttestation {
                    authority: authority__.unwrap_or_default(),
                    provider: provider__.unwrap_or_default(),
                    auditor: auditor__.unwrap_or_default(),
                    reason: reason__.unwrap_or_default(),
                    fault_attribution: fault_attribution__.unwrap_or_default(),
                    evidence_hash: evidence_hash__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgRevokeProviderAttestation", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgRevokeProviderAttestationResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.verification.v1.MsgRevokeProviderAttestationResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgRevokeProviderAttestationResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                            Err(serde::de::Error::unknown_field(value, FIELDS))
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgRevokeProviderAttestationResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgRevokeProviderAttestationResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgRevokeProviderAttestationResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgRevokeProviderAttestationResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgRevokeProviderAttestationResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgSettleAuditEscrow {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.authority.is_empty() {
            len += 1;
        }
        if self.audit_escrow_id != 0 {
            len += 1;
        }
        if self.reason != 0 {
            len += 1;
        }
        if self.fault_attribution != 0 {
            len += 1;
        }
        if !self.evidence_hash.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.MsgSettleAuditEscrow", len)?;
        if !self.authority.is_empty() {
            struct_ser.serialize_field("authority", &self.authority)?;
        }
        if self.audit_escrow_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("auditEscrowId", ToString::to_string(&self.audit_escrow_id).as_str())?;
        }
        if self.reason != 0 {
            let v = AuditEscrowSettlementReason::try_from(self.reason)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.reason)))?;
            struct_ser.serialize_field("reason", &v)?;
        }
        if self.fault_attribution != 0 {
            let v = FaultAttribution::try_from(self.fault_attribution)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.fault_attribution)))?;
            struct_ser.serialize_field("faultAttribution", &v)?;
        }
        if !self.evidence_hash.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("evidenceHash", pbjson::private::base64::encode(&self.evidence_hash).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgSettleAuditEscrow {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "authority",
            "audit_escrow_id",
            "auditEscrowId",
            "reason",
            "fault_attribution",
            "faultAttribution",
            "evidence_hash",
            "evidenceHash",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Authority,
            AuditEscrowId,
            Reason,
            FaultAttribution,
            EvidenceHash,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "authority" => Ok(GeneratedField::Authority),
                            "auditEscrowId" | "audit_escrow_id" => Ok(GeneratedField::AuditEscrowId),
                            "reason" => Ok(GeneratedField::Reason),
                            "faultAttribution" | "fault_attribution" => Ok(GeneratedField::FaultAttribution),
                            "evidenceHash" | "evidence_hash" => Ok(GeneratedField::EvidenceHash),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgSettleAuditEscrow;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgSettleAuditEscrow")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgSettleAuditEscrow, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut authority__ = None;
                let mut audit_escrow_id__ = None;
                let mut reason__ = None;
                let mut fault_attribution__ = None;
                let mut evidence_hash__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Authority => {
                            if authority__.is_some() {
                                return Err(serde::de::Error::duplicate_field("authority"));
                            }
                            authority__ = Some(map_.next_value()?);
                        }
                        GeneratedField::AuditEscrowId => {
                            if audit_escrow_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditEscrowId"));
                            }
                            audit_escrow_id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Reason => {
                            if reason__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reason"));
                            }
                            reason__ = Some(map_.next_value::<AuditEscrowSettlementReason>()? as i32);
                        }
                        GeneratedField::FaultAttribution => {
                            if fault_attribution__.is_some() {
                                return Err(serde::de::Error::duplicate_field("faultAttribution"));
                            }
                            fault_attribution__ = Some(map_.next_value::<FaultAttribution>()? as i32);
                        }
                        GeneratedField::EvidenceHash => {
                            if evidence_hash__.is_some() {
                                return Err(serde::de::Error::duplicate_field("evidenceHash"));
                            }
                            evidence_hash__ = 
                                Some(map_.next_value::<::pbjson::private::BytesDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(MsgSettleAuditEscrow {
                    authority: authority__.unwrap_or_default(),
                    audit_escrow_id: audit_escrow_id__.unwrap_or_default(),
                    reason: reason__.unwrap_or_default(),
                    fault_attribution: fault_attribution__.unwrap_or_default(),
                    evidence_hash: evidence_hash__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgSettleAuditEscrow", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgSettleAuditEscrowResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.verification.v1.MsgSettleAuditEscrowResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgSettleAuditEscrowResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                            Err(serde::de::Error::unknown_field(value, FIELDS))
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgSettleAuditEscrowResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgSettleAuditEscrowResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgSettleAuditEscrowResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgSettleAuditEscrowResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgSettleAuditEscrowResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgSlashProviderBond {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.authority.is_empty() {
            len += 1;
        }
        if !self.provider.is_empty() {
            len += 1;
        }
        if !self.slash_fraction.is_empty() {
            len += 1;
        }
        if self.reason != 0 {
            len += 1;
        }
        if !self.evidence_hash.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.MsgSlashProviderBond", len)?;
        if !self.authority.is_empty() {
            struct_ser.serialize_field("authority", &self.authority)?;
        }
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if !self.slash_fraction.is_empty() {
            struct_ser.serialize_field("slashFraction", &self.slash_fraction)?;
        }
        if self.reason != 0 {
            let v = ProviderBondSlashReason::try_from(self.reason)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.reason)))?;
            struct_ser.serialize_field("reason", &v)?;
        }
        if !self.evidence_hash.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("evidenceHash", pbjson::private::base64::encode(&self.evidence_hash).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgSlashProviderBond {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "authority",
            "provider",
            "slash_fraction",
            "slashFraction",
            "reason",
            "evidence_hash",
            "evidenceHash",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Authority,
            Provider,
            SlashFraction,
            Reason,
            EvidenceHash,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "authority" => Ok(GeneratedField::Authority),
                            "provider" => Ok(GeneratedField::Provider),
                            "slashFraction" | "slash_fraction" => Ok(GeneratedField::SlashFraction),
                            "reason" => Ok(GeneratedField::Reason),
                            "evidenceHash" | "evidence_hash" => Ok(GeneratedField::EvidenceHash),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgSlashProviderBond;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgSlashProviderBond")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgSlashProviderBond, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut authority__ = None;
                let mut provider__ = None;
                let mut slash_fraction__ = None;
                let mut reason__ = None;
                let mut evidence_hash__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Authority => {
                            if authority__.is_some() {
                                return Err(serde::de::Error::duplicate_field("authority"));
                            }
                            authority__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::SlashFraction => {
                            if slash_fraction__.is_some() {
                                return Err(serde::de::Error::duplicate_field("slashFraction"));
                            }
                            slash_fraction__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Reason => {
                            if reason__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reason"));
                            }
                            reason__ = Some(map_.next_value::<ProviderBondSlashReason>()? as i32);
                        }
                        GeneratedField::EvidenceHash => {
                            if evidence_hash__.is_some() {
                                return Err(serde::de::Error::duplicate_field("evidenceHash"));
                            }
                            evidence_hash__ = 
                                Some(map_.next_value::<::pbjson::private::BytesDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(MsgSlashProviderBond {
                    authority: authority__.unwrap_or_default(),
                    provider: provider__.unwrap_or_default(),
                    slash_fraction: slash_fraction__.unwrap_or_default(),
                    reason: reason__.unwrap_or_default(),
                    evidence_hash: evidence_hash__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgSlashProviderBond", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgSlashProviderBondResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.verification.v1.MsgSlashProviderBondResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgSlashProviderBondResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                            Err(serde::de::Error::unknown_field(value, FIELDS))
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgSlashProviderBondResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgSlashProviderBondResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgSlashProviderBondResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgSlashProviderBondResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgSlashProviderBondResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgSubmitAttestation {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if !self.auditor.is_empty() {
            len += 1;
        }
        if self.tier != 0 {
            len += 1;
        }
        if !self.capabilities.is_empty() {
            len += 1;
        }
        if !self.evidence_hash.is_empty() {
            len += 1;
        }
        if self.fee.is_some() {
            len += 1;
        }
        if self.deposit.is_some() {
            len += 1;
        }
        if self.audit_escrow_id != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.MsgSubmitAttestation", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        if self.tier != 0 {
            let v = VerificationTier::try_from(self.tier)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.tier)))?;
            struct_ser.serialize_field("tier", &v)?;
        }
        if !self.capabilities.is_empty() {
            let v = self.capabilities.iter().cloned().map(|v| {
                CapabilityFlag::try_from(v)
                    .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", v)))
                }).collect::<std::result::Result<Vec<_>, _>>()?;
            struct_ser.serialize_field("capabilities", &v)?;
        }
        if !self.evidence_hash.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("evidenceHash", pbjson::private::base64::encode(&self.evidence_hash).as_str())?;
        }
        if let Some(v) = self.fee.as_ref() {
            struct_ser.serialize_field("fee", v)?;
        }
        if let Some(v) = self.deposit.as_ref() {
            struct_ser.serialize_field("deposit", v)?;
        }
        if self.audit_escrow_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("auditEscrowId", ToString::to_string(&self.audit_escrow_id).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgSubmitAttestation {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "auditor",
            "tier",
            "capabilities",
            "evidence_hash",
            "evidenceHash",
            "fee",
            "deposit",
            "audit_escrow_id",
            "auditEscrowId",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            Auditor,
            Tier,
            Capabilities,
            EvidenceHash,
            Fee,
            Deposit,
            AuditEscrowId,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "auditor" => Ok(GeneratedField::Auditor),
                            "tier" => Ok(GeneratedField::Tier),
                            "capabilities" => Ok(GeneratedField::Capabilities),
                            "evidenceHash" | "evidence_hash" => Ok(GeneratedField::EvidenceHash),
                            "fee" => Ok(GeneratedField::Fee),
                            "deposit" => Ok(GeneratedField::Deposit),
                            "auditEscrowId" | "audit_escrow_id" => Ok(GeneratedField::AuditEscrowId),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgSubmitAttestation;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgSubmitAttestation")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgSubmitAttestation, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut auditor__ = None;
                let mut tier__ = None;
                let mut capabilities__ = None;
                let mut evidence_hash__ = None;
                let mut fee__ = None;
                let mut deposit__ = None;
                let mut audit_escrow_id__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Tier => {
                            if tier__.is_some() {
                                return Err(serde::de::Error::duplicate_field("tier"));
                            }
                            tier__ = Some(map_.next_value::<VerificationTier>()? as i32);
                        }
                        GeneratedField::Capabilities => {
                            if capabilities__.is_some() {
                                return Err(serde::de::Error::duplicate_field("capabilities"));
                            }
                            capabilities__ = Some(map_.next_value::<Vec<CapabilityFlag>>()?.into_iter().map(|x| x as i32).collect());
                        }
                        GeneratedField::EvidenceHash => {
                            if evidence_hash__.is_some() {
                                return Err(serde::de::Error::duplicate_field("evidenceHash"));
                            }
                            evidence_hash__ = 
                                Some(map_.next_value::<::pbjson::private::BytesDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Fee => {
                            if fee__.is_some() {
                                return Err(serde::de::Error::duplicate_field("fee"));
                            }
                            fee__ = map_.next_value()?;
                        }
                        GeneratedField::Deposit => {
                            if deposit__.is_some() {
                                return Err(serde::de::Error::duplicate_field("deposit"));
                            }
                            deposit__ = map_.next_value()?;
                        }
                        GeneratedField::AuditEscrowId => {
                            if audit_escrow_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditEscrowId"));
                            }
                            audit_escrow_id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(MsgSubmitAttestation {
                    provider: provider__.unwrap_or_default(),
                    auditor: auditor__.unwrap_or_default(),
                    tier: tier__.unwrap_or_default(),
                    capabilities: capabilities__.unwrap_or_default(),
                    evidence_hash: evidence_hash__.unwrap_or_default(),
                    fee: fee__,
                    deposit: deposit__,
                    audit_escrow_id: audit_escrow_id__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgSubmitAttestation", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgSubmitAttestationResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.verification.v1.MsgSubmitAttestationResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgSubmitAttestationResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                            Err(serde::de::Error::unknown_field(value, FIELDS))
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgSubmitAttestationResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgSubmitAttestationResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgSubmitAttestationResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgSubmitAttestationResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgSubmitAttestationResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgUpdateParams {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.authority.is_empty() {
            len += 1;
        }
        if self.params.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.MsgUpdateParams", len)?;
        if !self.authority.is_empty() {
            struct_ser.serialize_field("authority", &self.authority)?;
        }
        if let Some(v) = self.params.as_ref() {
            struct_ser.serialize_field("params", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgUpdateParams {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "authority",
            "params",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Authority,
            Params,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "authority" => Ok(GeneratedField::Authority),
                            "params" => Ok(GeneratedField::Params),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgUpdateParams;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgUpdateParams")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgUpdateParams, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut authority__ = None;
                let mut params__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Authority => {
                            if authority__.is_some() {
                                return Err(serde::de::Error::duplicate_field("authority"));
                            }
                            authority__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Params => {
                            if params__.is_some() {
                                return Err(serde::de::Error::duplicate_field("params"));
                            }
                            params__ = map_.next_value()?;
                        }
                    }
                }
                Ok(MsgUpdateParams {
                    authority: authority__.unwrap_or_default(),
                    params: params__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgUpdateParams", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgUpdateParamsResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.verification.v1.MsgUpdateParamsResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgUpdateParamsResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                            Err(serde::de::Error::unknown_field(value, FIELDS))
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgUpdateParamsResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgUpdateParamsResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgUpdateParamsResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgUpdateParamsResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgUpdateParamsResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgWithdrawProviderBond {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if self.amount.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.MsgWithdrawProviderBond", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if let Some(v) = self.amount.as_ref() {
            struct_ser.serialize_field("amount", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgWithdrawProviderBond {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "amount",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            Amount,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "amount" => Ok(GeneratedField::Amount),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgWithdrawProviderBond;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgWithdrawProviderBond")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgWithdrawProviderBond, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut amount__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Amount => {
                            if amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("amount"));
                            }
                            amount__ = map_.next_value()?;
                        }
                    }
                }
                Ok(MsgWithdrawProviderBond {
                    provider: provider__.unwrap_or_default(),
                    amount: amount__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgWithdrawProviderBond", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgWithdrawProviderBondResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.verification.v1.MsgWithdrawProviderBondResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgWithdrawProviderBondResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                            Err(serde::de::Error::unknown_field(value, FIELDS))
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgWithdrawProviderBondResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.MsgWithdrawProviderBondResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgWithdrawProviderBondResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgWithdrawProviderBondResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.MsgWithdrawProviderBondResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for Params {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.bond_l1.is_some() {
            len += 1;
        }
        if self.bond_l2.is_some() {
            len += 1;
        }
        if self.bond_l3.is_some() {
            len += 1;
        }
        if self.bond_l4.is_some() {
            len += 1;
        }
        if self.ttl_l1.is_some() {
            len += 1;
        }
        if self.ttl_l2.is_some() {
            len += 1;
        }
        if self.ttl_l3.is_some() {
            len += 1;
        }
        if self.ttl_l4.is_some() {
            len += 1;
        }
        if self.min_fee_l1.is_some() {
            len += 1;
        }
        if self.min_fee_l2.is_some() {
            len += 1;
        }
        if self.min_fee_l3.is_some() {
            len += 1;
        }
        if self.min_fee_l4.is_some() {
            len += 1;
        }
        if self.discrepancy_threshold != 0 {
            len += 1;
        }
        if self.auditor_unbonding_period.is_some() {
            len += 1;
        }
        if self.renewal_period_l1.is_some() {
            len += 1;
        }
        if self.renewal_period_l2.is_some() {
            len += 1;
        }
        if self.renewal_period_l3.is_some() {
            len += 1;
        }
        if self.renewal_period_l4.is_some() {
            len += 1;
        }
        if self.snapshot_hash_interval.is_some() {
            len += 1;
        }
        if self.max_snapshot_age.is_some() {
            len += 1;
        }
        if self.bond_gpu_l2.is_some() {
            len += 1;
        }
        if self.bond_gpu_l3.is_some() {
            len += 1;
        }
        if self.bond_gpu_l4.is_some() {
            len += 1;
        }
        if self.bond_vcpu_l2.is_some() {
            len += 1;
        }
        if self.bond_vcpu_l3.is_some() {
            len += 1;
        }
        if self.bond_vcpu_l4.is_some() {
            len += 1;
        }
        if self.bond_mem_gb_l2.is_some() {
            len += 1;
        }
        if self.bond_mem_gb_l3.is_some() {
            len += 1;
        }
        if self.bond_mem_gb_l4.is_some() {
            len += 1;
        }
        if self.bond_storage_tb_l2.is_some() {
            len += 1;
        }
        if self.bond_storage_tb_l3.is_some() {
            len += 1;
        }
        if self.bond_storage_tb_l4.is_some() {
            len += 1;
        }
        if self.provider_bond_unbonding_period.is_some() {
            len += 1;
        }
        if self.min_age_l2.is_some() {
            len += 1;
        }
        if self.min_age_l3.is_some() {
            len += 1;
        }
        if self.min_age_l4.is_some() {
            len += 1;
        }
        if self.min_lease_completion_bps_l3 != 0 {
            len += 1;
        }
        if self.min_lease_completion_bps_l4 != 0 {
            len += 1;
        }
        if self.clean_history_window_l3.is_some() {
            len += 1;
        }
        if self.clean_history_window_l4.is_some() {
            len += 1;
        }
        if self.min_l3_duration_for_l4.is_some() {
            len += 1;
        }
        if self.min_leases_for_completion_rate != 0 {
            len += 1;
        }
        if self.max_endblocker_attestation_expiries != 0 {
            len += 1;
        }
        if self.max_endblocker_snapshot_suspensions != 0 {
            len += 1;
        }
        if self.max_endblocker_unbonding_completions != 0 {
            len += 1;
        }
        if self.max_endblocker_discrepancy_timeouts != 0 {
            len += 1;
        }
        if self.max_endblocker_audit_escrow_expiries != 0 {
            len += 1;
        }
        if self.max_endblocker_grace_expiries != 0 {
            len += 1;
        }
        if self.discrepancy_resolution_timeout.is_some() {
            len += 1;
        }
        if self.attestation_deposit.is_some() {
            len += 1;
        }
        if self.discrepancy_grace_period.is_some() {
            len += 1;
        }
        if self.provider_audit_deposit.is_some() {
            len += 1;
        }
        if self.verification_module_active {
            len += 1;
        }
        if self.contact_response_critical_l1.is_some() {
            len += 1;
        }
        if self.contact_response_critical_l2.is_some() {
            len += 1;
        }
        if self.contact_response_critical_l3.is_some() {
            len += 1;
        }
        if self.contact_response_critical_l4.is_some() {
            len += 1;
        }
        if self.contact_response_standard_l1.is_some() {
            len += 1;
        }
        if self.contact_response_standard_l2.is_some() {
            len += 1;
        }
        if self.contact_response_standard_l3.is_some() {
            len += 1;
        }
        if self.contact_response_standard_l4.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.Params", len)?;
        if let Some(v) = self.bond_l1.as_ref() {
            struct_ser.serialize_field("bondL1", v)?;
        }
        if let Some(v) = self.bond_l2.as_ref() {
            struct_ser.serialize_field("bondL2", v)?;
        }
        if let Some(v) = self.bond_l3.as_ref() {
            struct_ser.serialize_field("bondL3", v)?;
        }
        if let Some(v) = self.bond_l4.as_ref() {
            struct_ser.serialize_field("bondL4", v)?;
        }
        if let Some(v) = self.ttl_l1.as_ref() {
            struct_ser.serialize_field("ttlL1", v)?;
        }
        if let Some(v) = self.ttl_l2.as_ref() {
            struct_ser.serialize_field("ttlL2", v)?;
        }
        if let Some(v) = self.ttl_l3.as_ref() {
            struct_ser.serialize_field("ttlL3", v)?;
        }
        if let Some(v) = self.ttl_l4.as_ref() {
            struct_ser.serialize_field("ttlL4", v)?;
        }
        if let Some(v) = self.min_fee_l1.as_ref() {
            struct_ser.serialize_field("minFeeL1", v)?;
        }
        if let Some(v) = self.min_fee_l2.as_ref() {
            struct_ser.serialize_field("minFeeL2", v)?;
        }
        if let Some(v) = self.min_fee_l3.as_ref() {
            struct_ser.serialize_field("minFeeL3", v)?;
        }
        if let Some(v) = self.min_fee_l4.as_ref() {
            struct_ser.serialize_field("minFeeL4", v)?;
        }
        if self.discrepancy_threshold != 0 {
            struct_ser.serialize_field("discrepancyThreshold", &self.discrepancy_threshold)?;
        }
        if let Some(v) = self.auditor_unbonding_period.as_ref() {
            struct_ser.serialize_field("auditorUnbondingPeriod", v)?;
        }
        if let Some(v) = self.renewal_period_l1.as_ref() {
            struct_ser.serialize_field("renewalPeriodL1", v)?;
        }
        if let Some(v) = self.renewal_period_l2.as_ref() {
            struct_ser.serialize_field("renewalPeriodL2", v)?;
        }
        if let Some(v) = self.renewal_period_l3.as_ref() {
            struct_ser.serialize_field("renewalPeriodL3", v)?;
        }
        if let Some(v) = self.renewal_period_l4.as_ref() {
            struct_ser.serialize_field("renewalPeriodL4", v)?;
        }
        if let Some(v) = self.snapshot_hash_interval.as_ref() {
            struct_ser.serialize_field("snapshotHashInterval", v)?;
        }
        if let Some(v) = self.max_snapshot_age.as_ref() {
            struct_ser.serialize_field("maxSnapshotAge", v)?;
        }
        if let Some(v) = self.bond_gpu_l2.as_ref() {
            struct_ser.serialize_field("bondGpuL2", v)?;
        }
        if let Some(v) = self.bond_gpu_l3.as_ref() {
            struct_ser.serialize_field("bondGpuL3", v)?;
        }
        if let Some(v) = self.bond_gpu_l4.as_ref() {
            struct_ser.serialize_field("bondGpuL4", v)?;
        }
        if let Some(v) = self.bond_vcpu_l2.as_ref() {
            struct_ser.serialize_field("bondVcpuL2", v)?;
        }
        if let Some(v) = self.bond_vcpu_l3.as_ref() {
            struct_ser.serialize_field("bondVcpuL3", v)?;
        }
        if let Some(v) = self.bond_vcpu_l4.as_ref() {
            struct_ser.serialize_field("bondVcpuL4", v)?;
        }
        if let Some(v) = self.bond_mem_gb_l2.as_ref() {
            struct_ser.serialize_field("bondMemGbL2", v)?;
        }
        if let Some(v) = self.bond_mem_gb_l3.as_ref() {
            struct_ser.serialize_field("bondMemGbL3", v)?;
        }
        if let Some(v) = self.bond_mem_gb_l4.as_ref() {
            struct_ser.serialize_field("bondMemGbL4", v)?;
        }
        if let Some(v) = self.bond_storage_tb_l2.as_ref() {
            struct_ser.serialize_field("bondStorageTbL2", v)?;
        }
        if let Some(v) = self.bond_storage_tb_l3.as_ref() {
            struct_ser.serialize_field("bondStorageTbL3", v)?;
        }
        if let Some(v) = self.bond_storage_tb_l4.as_ref() {
            struct_ser.serialize_field("bondStorageTbL4", v)?;
        }
        if let Some(v) = self.provider_bond_unbonding_period.as_ref() {
            struct_ser.serialize_field("providerBondUnbondingPeriod", v)?;
        }
        if let Some(v) = self.min_age_l2.as_ref() {
            struct_ser.serialize_field("minAgeL2", v)?;
        }
        if let Some(v) = self.min_age_l3.as_ref() {
            struct_ser.serialize_field("minAgeL3", v)?;
        }
        if let Some(v) = self.min_age_l4.as_ref() {
            struct_ser.serialize_field("minAgeL4", v)?;
        }
        if self.min_lease_completion_bps_l3 != 0 {
            struct_ser.serialize_field("minLeaseCompletionBpsL3", &self.min_lease_completion_bps_l3)?;
        }
        if self.min_lease_completion_bps_l4 != 0 {
            struct_ser.serialize_field("minLeaseCompletionBpsL4", &self.min_lease_completion_bps_l4)?;
        }
        if let Some(v) = self.clean_history_window_l3.as_ref() {
            struct_ser.serialize_field("cleanHistoryWindowL3", v)?;
        }
        if let Some(v) = self.clean_history_window_l4.as_ref() {
            struct_ser.serialize_field("cleanHistoryWindowL4", v)?;
        }
        if let Some(v) = self.min_l3_duration_for_l4.as_ref() {
            struct_ser.serialize_field("minL3DurationForL4", v)?;
        }
        if self.min_leases_for_completion_rate != 0 {
            struct_ser.serialize_field("minLeasesForCompletionRate", &self.min_leases_for_completion_rate)?;
        }
        if self.max_endblocker_attestation_expiries != 0 {
            struct_ser.serialize_field("maxEndblockerAttestationExpiries", &self.max_endblocker_attestation_expiries)?;
        }
        if self.max_endblocker_snapshot_suspensions != 0 {
            struct_ser.serialize_field("maxEndblockerSnapshotSuspensions", &self.max_endblocker_snapshot_suspensions)?;
        }
        if self.max_endblocker_unbonding_completions != 0 {
            struct_ser.serialize_field("maxEndblockerUnbondingCompletions", &self.max_endblocker_unbonding_completions)?;
        }
        if self.max_endblocker_discrepancy_timeouts != 0 {
            struct_ser.serialize_field("maxEndblockerDiscrepancyTimeouts", &self.max_endblocker_discrepancy_timeouts)?;
        }
        if self.max_endblocker_audit_escrow_expiries != 0 {
            struct_ser.serialize_field("maxEndblockerAuditEscrowExpiries", &self.max_endblocker_audit_escrow_expiries)?;
        }
        if self.max_endblocker_grace_expiries != 0 {
            struct_ser.serialize_field("maxEndblockerGraceExpiries", &self.max_endblocker_grace_expiries)?;
        }
        if let Some(v) = self.discrepancy_resolution_timeout.as_ref() {
            struct_ser.serialize_field("discrepancyResolutionTimeout", v)?;
        }
        if let Some(v) = self.attestation_deposit.as_ref() {
            struct_ser.serialize_field("attestationDeposit", v)?;
        }
        if let Some(v) = self.discrepancy_grace_period.as_ref() {
            struct_ser.serialize_field("discrepancyGracePeriod", v)?;
        }
        if let Some(v) = self.provider_audit_deposit.as_ref() {
            struct_ser.serialize_field("providerAuditDeposit", v)?;
        }
        if self.verification_module_active {
            struct_ser.serialize_field("verificationModuleActive", &self.verification_module_active)?;
        }
        if let Some(v) = self.contact_response_critical_l1.as_ref() {
            struct_ser.serialize_field("contactResponseCriticalL1", v)?;
        }
        if let Some(v) = self.contact_response_critical_l2.as_ref() {
            struct_ser.serialize_field("contactResponseCriticalL2", v)?;
        }
        if let Some(v) = self.contact_response_critical_l3.as_ref() {
            struct_ser.serialize_field("contactResponseCriticalL3", v)?;
        }
        if let Some(v) = self.contact_response_critical_l4.as_ref() {
            struct_ser.serialize_field("contactResponseCriticalL4", v)?;
        }
        if let Some(v) = self.contact_response_standard_l1.as_ref() {
            struct_ser.serialize_field("contactResponseStandardL1", v)?;
        }
        if let Some(v) = self.contact_response_standard_l2.as_ref() {
            struct_ser.serialize_field("contactResponseStandardL2", v)?;
        }
        if let Some(v) = self.contact_response_standard_l3.as_ref() {
            struct_ser.serialize_field("contactResponseStandardL3", v)?;
        }
        if let Some(v) = self.contact_response_standard_l4.as_ref() {
            struct_ser.serialize_field("contactResponseStandardL4", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for Params {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "bond_l1",
            "bondL1",
            "bond_l2",
            "bondL2",
            "bond_l3",
            "bondL3",
            "bond_l4",
            "bondL4",
            "ttl_l1",
            "ttlL1",
            "ttl_l2",
            "ttlL2",
            "ttl_l3",
            "ttlL3",
            "ttl_l4",
            "ttlL4",
            "min_fee_l1",
            "minFeeL1",
            "min_fee_l2",
            "minFeeL2",
            "min_fee_l3",
            "minFeeL3",
            "min_fee_l4",
            "minFeeL4",
            "discrepancy_threshold",
            "discrepancyThreshold",
            "auditor_unbonding_period",
            "auditorUnbondingPeriod",
            "renewal_period_l1",
            "renewalPeriodL1",
            "renewal_period_l2",
            "renewalPeriodL2",
            "renewal_period_l3",
            "renewalPeriodL3",
            "renewal_period_l4",
            "renewalPeriodL4",
            "snapshot_hash_interval",
            "snapshotHashInterval",
            "max_snapshot_age",
            "maxSnapshotAge",
            "bond_gpu_l2",
            "bondGpuL2",
            "bond_gpu_l3",
            "bondGpuL3",
            "bond_gpu_l4",
            "bondGpuL4",
            "bond_vcpu_l2",
            "bondVcpuL2",
            "bond_vcpu_l3",
            "bondVcpuL3",
            "bond_vcpu_l4",
            "bondVcpuL4",
            "bond_mem_gb_l2",
            "bondMemGbL2",
            "bond_mem_gb_l3",
            "bondMemGbL3",
            "bond_mem_gb_l4",
            "bondMemGbL4",
            "bond_storage_tb_l2",
            "bondStorageTbL2",
            "bond_storage_tb_l3",
            "bondStorageTbL3",
            "bond_storage_tb_l4",
            "bondStorageTbL4",
            "provider_bond_unbonding_period",
            "providerBondUnbondingPeriod",
            "min_age_l2",
            "minAgeL2",
            "min_age_l3",
            "minAgeL3",
            "min_age_l4",
            "minAgeL4",
            "min_lease_completion_bps_l3",
            "minLeaseCompletionBpsL3",
            "min_lease_completion_bps_l4",
            "minLeaseCompletionBpsL4",
            "clean_history_window_l3",
            "cleanHistoryWindowL3",
            "clean_history_window_l4",
            "cleanHistoryWindowL4",
            "min_l3_duration_for_l4",
            "minL3DurationForL4",
            "min_leases_for_completion_rate",
            "minLeasesForCompletionRate",
            "max_endblocker_attestation_expiries",
            "maxEndblockerAttestationExpiries",
            "max_endblocker_snapshot_suspensions",
            "maxEndblockerSnapshotSuspensions",
            "max_endblocker_unbonding_completions",
            "maxEndblockerUnbondingCompletions",
            "max_endblocker_discrepancy_timeouts",
            "maxEndblockerDiscrepancyTimeouts",
            "max_endblocker_audit_escrow_expiries",
            "maxEndblockerAuditEscrowExpiries",
            "max_endblocker_grace_expiries",
            "maxEndblockerGraceExpiries",
            "discrepancy_resolution_timeout",
            "discrepancyResolutionTimeout",
            "attestation_deposit",
            "attestationDeposit",
            "discrepancy_grace_period",
            "discrepancyGracePeriod",
            "provider_audit_deposit",
            "providerAuditDeposit",
            "verification_module_active",
            "verificationModuleActive",
            "contact_response_critical_l1",
            "contactResponseCriticalL1",
            "contact_response_critical_l2",
            "contactResponseCriticalL2",
            "contact_response_critical_l3",
            "contactResponseCriticalL3",
            "contact_response_critical_l4",
            "contactResponseCriticalL4",
            "contact_response_standard_l1",
            "contactResponseStandardL1",
            "contact_response_standard_l2",
            "contactResponseStandardL2",
            "contact_response_standard_l3",
            "contactResponseStandardL3",
            "contact_response_standard_l4",
            "contactResponseStandardL4",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            BondL1,
            BondL2,
            BondL3,
            BondL4,
            TtlL1,
            TtlL2,
            TtlL3,
            TtlL4,
            MinFeeL1,
            MinFeeL2,
            MinFeeL3,
            MinFeeL4,
            DiscrepancyThreshold,
            AuditorUnbondingPeriod,
            RenewalPeriodL1,
            RenewalPeriodL2,
            RenewalPeriodL3,
            RenewalPeriodL4,
            SnapshotHashInterval,
            MaxSnapshotAge,
            BondGpuL2,
            BondGpuL3,
            BondGpuL4,
            BondVcpuL2,
            BondVcpuL3,
            BondVcpuL4,
            BondMemGbL2,
            BondMemGbL3,
            BondMemGbL4,
            BondStorageTbL2,
            BondStorageTbL3,
            BondStorageTbL4,
            ProviderBondUnbondingPeriod,
            MinAgeL2,
            MinAgeL3,
            MinAgeL4,
            MinLeaseCompletionBpsL3,
            MinLeaseCompletionBpsL4,
            CleanHistoryWindowL3,
            CleanHistoryWindowL4,
            MinL3DurationForL4,
            MinLeasesForCompletionRate,
            MaxEndblockerAttestationExpiries,
            MaxEndblockerSnapshotSuspensions,
            MaxEndblockerUnbondingCompletions,
            MaxEndblockerDiscrepancyTimeouts,
            MaxEndblockerAuditEscrowExpiries,
            MaxEndblockerGraceExpiries,
            DiscrepancyResolutionTimeout,
            AttestationDeposit,
            DiscrepancyGracePeriod,
            ProviderAuditDeposit,
            VerificationModuleActive,
            ContactResponseCriticalL1,
            ContactResponseCriticalL2,
            ContactResponseCriticalL3,
            ContactResponseCriticalL4,
            ContactResponseStandardL1,
            ContactResponseStandardL2,
            ContactResponseStandardL3,
            ContactResponseStandardL4,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "bondL1" | "bond_l1" => Ok(GeneratedField::BondL1),
                            "bondL2" | "bond_l2" => Ok(GeneratedField::BondL2),
                            "bondL3" | "bond_l3" => Ok(GeneratedField::BondL3),
                            "bondL4" | "bond_l4" => Ok(GeneratedField::BondL4),
                            "ttlL1" | "ttl_l1" => Ok(GeneratedField::TtlL1),
                            "ttlL2" | "ttl_l2" => Ok(GeneratedField::TtlL2),
                            "ttlL3" | "ttl_l3" => Ok(GeneratedField::TtlL3),
                            "ttlL4" | "ttl_l4" => Ok(GeneratedField::TtlL4),
                            "minFeeL1" | "min_fee_l1" => Ok(GeneratedField::MinFeeL1),
                            "minFeeL2" | "min_fee_l2" => Ok(GeneratedField::MinFeeL2),
                            "minFeeL3" | "min_fee_l3" => Ok(GeneratedField::MinFeeL3),
                            "minFeeL4" | "min_fee_l4" => Ok(GeneratedField::MinFeeL4),
                            "discrepancyThreshold" | "discrepancy_threshold" => Ok(GeneratedField::DiscrepancyThreshold),
                            "auditorUnbondingPeriod" | "auditor_unbonding_period" => Ok(GeneratedField::AuditorUnbondingPeriod),
                            "renewalPeriodL1" | "renewal_period_l1" => Ok(GeneratedField::RenewalPeriodL1),
                            "renewalPeriodL2" | "renewal_period_l2" => Ok(GeneratedField::RenewalPeriodL2),
                            "renewalPeriodL3" | "renewal_period_l3" => Ok(GeneratedField::RenewalPeriodL3),
                            "renewalPeriodL4" | "renewal_period_l4" => Ok(GeneratedField::RenewalPeriodL4),
                            "snapshotHashInterval" | "snapshot_hash_interval" => Ok(GeneratedField::SnapshotHashInterval),
                            "maxSnapshotAge" | "max_snapshot_age" => Ok(GeneratedField::MaxSnapshotAge),
                            "bondGpuL2" | "bond_gpu_l2" => Ok(GeneratedField::BondGpuL2),
                            "bondGpuL3" | "bond_gpu_l3" => Ok(GeneratedField::BondGpuL3),
                            "bondGpuL4" | "bond_gpu_l4" => Ok(GeneratedField::BondGpuL4),
                            "bondVcpuL2" | "bond_vcpu_l2" => Ok(GeneratedField::BondVcpuL2),
                            "bondVcpuL3" | "bond_vcpu_l3" => Ok(GeneratedField::BondVcpuL3),
                            "bondVcpuL4" | "bond_vcpu_l4" => Ok(GeneratedField::BondVcpuL4),
                            "bondMemGbL2" | "bond_mem_gb_l2" => Ok(GeneratedField::BondMemGbL2),
                            "bondMemGbL3" | "bond_mem_gb_l3" => Ok(GeneratedField::BondMemGbL3),
                            "bondMemGbL4" | "bond_mem_gb_l4" => Ok(GeneratedField::BondMemGbL4),
                            "bondStorageTbL2" | "bond_storage_tb_l2" => Ok(GeneratedField::BondStorageTbL2),
                            "bondStorageTbL3" | "bond_storage_tb_l3" => Ok(GeneratedField::BondStorageTbL3),
                            "bondStorageTbL4" | "bond_storage_tb_l4" => Ok(GeneratedField::BondStorageTbL4),
                            "providerBondUnbondingPeriod" | "provider_bond_unbonding_period" => Ok(GeneratedField::ProviderBondUnbondingPeriod),
                            "minAgeL2" | "min_age_l2" => Ok(GeneratedField::MinAgeL2),
                            "minAgeL3" | "min_age_l3" => Ok(GeneratedField::MinAgeL3),
                            "minAgeL4" | "min_age_l4" => Ok(GeneratedField::MinAgeL4),
                            "minLeaseCompletionBpsL3" | "min_lease_completion_bps_l3" => Ok(GeneratedField::MinLeaseCompletionBpsL3),
                            "minLeaseCompletionBpsL4" | "min_lease_completion_bps_l4" => Ok(GeneratedField::MinLeaseCompletionBpsL4),
                            "cleanHistoryWindowL3" | "clean_history_window_l3" => Ok(GeneratedField::CleanHistoryWindowL3),
                            "cleanHistoryWindowL4" | "clean_history_window_l4" => Ok(GeneratedField::CleanHistoryWindowL4),
                            "minL3DurationForL4" | "min_l3_duration_for_l4" => Ok(GeneratedField::MinL3DurationForL4),
                            "minLeasesForCompletionRate" | "min_leases_for_completion_rate" => Ok(GeneratedField::MinLeasesForCompletionRate),
                            "maxEndblockerAttestationExpiries" | "max_endblocker_attestation_expiries" => Ok(GeneratedField::MaxEndblockerAttestationExpiries),
                            "maxEndblockerSnapshotSuspensions" | "max_endblocker_snapshot_suspensions" => Ok(GeneratedField::MaxEndblockerSnapshotSuspensions),
                            "maxEndblockerUnbondingCompletions" | "max_endblocker_unbonding_completions" => Ok(GeneratedField::MaxEndblockerUnbondingCompletions),
                            "maxEndblockerDiscrepancyTimeouts" | "max_endblocker_discrepancy_timeouts" => Ok(GeneratedField::MaxEndblockerDiscrepancyTimeouts),
                            "maxEndblockerAuditEscrowExpiries" | "max_endblocker_audit_escrow_expiries" => Ok(GeneratedField::MaxEndblockerAuditEscrowExpiries),
                            "maxEndblockerGraceExpiries" | "max_endblocker_grace_expiries" => Ok(GeneratedField::MaxEndblockerGraceExpiries),
                            "discrepancyResolutionTimeout" | "discrepancy_resolution_timeout" => Ok(GeneratedField::DiscrepancyResolutionTimeout),
                            "attestationDeposit" | "attestation_deposit" => Ok(GeneratedField::AttestationDeposit),
                            "discrepancyGracePeriod" | "discrepancy_grace_period" => Ok(GeneratedField::DiscrepancyGracePeriod),
                            "providerAuditDeposit" | "provider_audit_deposit" => Ok(GeneratedField::ProviderAuditDeposit),
                            "verificationModuleActive" | "verification_module_active" => Ok(GeneratedField::VerificationModuleActive),
                            "contactResponseCriticalL1" | "contact_response_critical_l1" => Ok(GeneratedField::ContactResponseCriticalL1),
                            "contactResponseCriticalL2" | "contact_response_critical_l2" => Ok(GeneratedField::ContactResponseCriticalL2),
                            "contactResponseCriticalL3" | "contact_response_critical_l3" => Ok(GeneratedField::ContactResponseCriticalL3),
                            "contactResponseCriticalL4" | "contact_response_critical_l4" => Ok(GeneratedField::ContactResponseCriticalL4),
                            "contactResponseStandardL1" | "contact_response_standard_l1" => Ok(GeneratedField::ContactResponseStandardL1),
                            "contactResponseStandardL2" | "contact_response_standard_l2" => Ok(GeneratedField::ContactResponseStandardL2),
                            "contactResponseStandardL3" | "contact_response_standard_l3" => Ok(GeneratedField::ContactResponseStandardL3),
                            "contactResponseStandardL4" | "contact_response_standard_l4" => Ok(GeneratedField::ContactResponseStandardL4),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = Params;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.Params")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<Params, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut bond_l1__ = None;
                let mut bond_l2__ = None;
                let mut bond_l3__ = None;
                let mut bond_l4__ = None;
                let mut ttl_l1__ = None;
                let mut ttl_l2__ = None;
                let mut ttl_l3__ = None;
                let mut ttl_l4__ = None;
                let mut min_fee_l1__ = None;
                let mut min_fee_l2__ = None;
                let mut min_fee_l3__ = None;
                let mut min_fee_l4__ = None;
                let mut discrepancy_threshold__ = None;
                let mut auditor_unbonding_period__ = None;
                let mut renewal_period_l1__ = None;
                let mut renewal_period_l2__ = None;
                let mut renewal_period_l3__ = None;
                let mut renewal_period_l4__ = None;
                let mut snapshot_hash_interval__ = None;
                let mut max_snapshot_age__ = None;
                let mut bond_gpu_l2__ = None;
                let mut bond_gpu_l3__ = None;
                let mut bond_gpu_l4__ = None;
                let mut bond_vcpu_l2__ = None;
                let mut bond_vcpu_l3__ = None;
                let mut bond_vcpu_l4__ = None;
                let mut bond_mem_gb_l2__ = None;
                let mut bond_mem_gb_l3__ = None;
                let mut bond_mem_gb_l4__ = None;
                let mut bond_storage_tb_l2__ = None;
                let mut bond_storage_tb_l3__ = None;
                let mut bond_storage_tb_l4__ = None;
                let mut provider_bond_unbonding_period__ = None;
                let mut min_age_l2__ = None;
                let mut min_age_l3__ = None;
                let mut min_age_l4__ = None;
                let mut min_lease_completion_bps_l3__ = None;
                let mut min_lease_completion_bps_l4__ = None;
                let mut clean_history_window_l3__ = None;
                let mut clean_history_window_l4__ = None;
                let mut min_l3_duration_for_l4__ = None;
                let mut min_leases_for_completion_rate__ = None;
                let mut max_endblocker_attestation_expiries__ = None;
                let mut max_endblocker_snapshot_suspensions__ = None;
                let mut max_endblocker_unbonding_completions__ = None;
                let mut max_endblocker_discrepancy_timeouts__ = None;
                let mut max_endblocker_audit_escrow_expiries__ = None;
                let mut max_endblocker_grace_expiries__ = None;
                let mut discrepancy_resolution_timeout__ = None;
                let mut attestation_deposit__ = None;
                let mut discrepancy_grace_period__ = None;
                let mut provider_audit_deposit__ = None;
                let mut verification_module_active__ = None;
                let mut contact_response_critical_l1__ = None;
                let mut contact_response_critical_l2__ = None;
                let mut contact_response_critical_l3__ = None;
                let mut contact_response_critical_l4__ = None;
                let mut contact_response_standard_l1__ = None;
                let mut contact_response_standard_l2__ = None;
                let mut contact_response_standard_l3__ = None;
                let mut contact_response_standard_l4__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::BondL1 => {
                            if bond_l1__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bondL1"));
                            }
                            bond_l1__ = map_.next_value()?;
                        }
                        GeneratedField::BondL2 => {
                            if bond_l2__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bondL2"));
                            }
                            bond_l2__ = map_.next_value()?;
                        }
                        GeneratedField::BondL3 => {
                            if bond_l3__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bondL3"));
                            }
                            bond_l3__ = map_.next_value()?;
                        }
                        GeneratedField::BondL4 => {
                            if bond_l4__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bondL4"));
                            }
                            bond_l4__ = map_.next_value()?;
                        }
                        GeneratedField::TtlL1 => {
                            if ttl_l1__.is_some() {
                                return Err(serde::de::Error::duplicate_field("ttlL1"));
                            }
                            ttl_l1__ = map_.next_value()?;
                        }
                        GeneratedField::TtlL2 => {
                            if ttl_l2__.is_some() {
                                return Err(serde::de::Error::duplicate_field("ttlL2"));
                            }
                            ttl_l2__ = map_.next_value()?;
                        }
                        GeneratedField::TtlL3 => {
                            if ttl_l3__.is_some() {
                                return Err(serde::de::Error::duplicate_field("ttlL3"));
                            }
                            ttl_l3__ = map_.next_value()?;
                        }
                        GeneratedField::TtlL4 => {
                            if ttl_l4__.is_some() {
                                return Err(serde::de::Error::duplicate_field("ttlL4"));
                            }
                            ttl_l4__ = map_.next_value()?;
                        }
                        GeneratedField::MinFeeL1 => {
                            if min_fee_l1__.is_some() {
                                return Err(serde::de::Error::duplicate_field("minFeeL1"));
                            }
                            min_fee_l1__ = map_.next_value()?;
                        }
                        GeneratedField::MinFeeL2 => {
                            if min_fee_l2__.is_some() {
                                return Err(serde::de::Error::duplicate_field("minFeeL2"));
                            }
                            min_fee_l2__ = map_.next_value()?;
                        }
                        GeneratedField::MinFeeL3 => {
                            if min_fee_l3__.is_some() {
                                return Err(serde::de::Error::duplicate_field("minFeeL3"));
                            }
                            min_fee_l3__ = map_.next_value()?;
                        }
                        GeneratedField::MinFeeL4 => {
                            if min_fee_l4__.is_some() {
                                return Err(serde::de::Error::duplicate_field("minFeeL4"));
                            }
                            min_fee_l4__ = map_.next_value()?;
                        }
                        GeneratedField::DiscrepancyThreshold => {
                            if discrepancy_threshold__.is_some() {
                                return Err(serde::de::Error::duplicate_field("discrepancyThreshold"));
                            }
                            discrepancy_threshold__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::AuditorUnbondingPeriod => {
                            if auditor_unbonding_period__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditorUnbondingPeriod"));
                            }
                            auditor_unbonding_period__ = map_.next_value()?;
                        }
                        GeneratedField::RenewalPeriodL1 => {
                            if renewal_period_l1__.is_some() {
                                return Err(serde::de::Error::duplicate_field("renewalPeriodL1"));
                            }
                            renewal_period_l1__ = map_.next_value()?;
                        }
                        GeneratedField::RenewalPeriodL2 => {
                            if renewal_period_l2__.is_some() {
                                return Err(serde::de::Error::duplicate_field("renewalPeriodL2"));
                            }
                            renewal_period_l2__ = map_.next_value()?;
                        }
                        GeneratedField::RenewalPeriodL3 => {
                            if renewal_period_l3__.is_some() {
                                return Err(serde::de::Error::duplicate_field("renewalPeriodL3"));
                            }
                            renewal_period_l3__ = map_.next_value()?;
                        }
                        GeneratedField::RenewalPeriodL4 => {
                            if renewal_period_l4__.is_some() {
                                return Err(serde::de::Error::duplicate_field("renewalPeriodL4"));
                            }
                            renewal_period_l4__ = map_.next_value()?;
                        }
                        GeneratedField::SnapshotHashInterval => {
                            if snapshot_hash_interval__.is_some() {
                                return Err(serde::de::Error::duplicate_field("snapshotHashInterval"));
                            }
                            snapshot_hash_interval__ = map_.next_value()?;
                        }
                        GeneratedField::MaxSnapshotAge => {
                            if max_snapshot_age__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maxSnapshotAge"));
                            }
                            max_snapshot_age__ = map_.next_value()?;
                        }
                        GeneratedField::BondGpuL2 => {
                            if bond_gpu_l2__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bondGpuL2"));
                            }
                            bond_gpu_l2__ = map_.next_value()?;
                        }
                        GeneratedField::BondGpuL3 => {
                            if bond_gpu_l3__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bondGpuL3"));
                            }
                            bond_gpu_l3__ = map_.next_value()?;
                        }
                        GeneratedField::BondGpuL4 => {
                            if bond_gpu_l4__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bondGpuL4"));
                            }
                            bond_gpu_l4__ = map_.next_value()?;
                        }
                        GeneratedField::BondVcpuL2 => {
                            if bond_vcpu_l2__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bondVcpuL2"));
                            }
                            bond_vcpu_l2__ = map_.next_value()?;
                        }
                        GeneratedField::BondVcpuL3 => {
                            if bond_vcpu_l3__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bondVcpuL3"));
                            }
                            bond_vcpu_l3__ = map_.next_value()?;
                        }
                        GeneratedField::BondVcpuL4 => {
                            if bond_vcpu_l4__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bondVcpuL4"));
                            }
                            bond_vcpu_l4__ = map_.next_value()?;
                        }
                        GeneratedField::BondMemGbL2 => {
                            if bond_mem_gb_l2__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bondMemGbL2"));
                            }
                            bond_mem_gb_l2__ = map_.next_value()?;
                        }
                        GeneratedField::BondMemGbL3 => {
                            if bond_mem_gb_l3__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bondMemGbL3"));
                            }
                            bond_mem_gb_l3__ = map_.next_value()?;
                        }
                        GeneratedField::BondMemGbL4 => {
                            if bond_mem_gb_l4__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bondMemGbL4"));
                            }
                            bond_mem_gb_l4__ = map_.next_value()?;
                        }
                        GeneratedField::BondStorageTbL2 => {
                            if bond_storage_tb_l2__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bondStorageTbL2"));
                            }
                            bond_storage_tb_l2__ = map_.next_value()?;
                        }
                        GeneratedField::BondStorageTbL3 => {
                            if bond_storage_tb_l3__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bondStorageTbL3"));
                            }
                            bond_storage_tb_l3__ = map_.next_value()?;
                        }
                        GeneratedField::BondStorageTbL4 => {
                            if bond_storage_tb_l4__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bondStorageTbL4"));
                            }
                            bond_storage_tb_l4__ = map_.next_value()?;
                        }
                        GeneratedField::ProviderBondUnbondingPeriod => {
                            if provider_bond_unbonding_period__.is_some() {
                                return Err(serde::de::Error::duplicate_field("providerBondUnbondingPeriod"));
                            }
                            provider_bond_unbonding_period__ = map_.next_value()?;
                        }
                        GeneratedField::MinAgeL2 => {
                            if min_age_l2__.is_some() {
                                return Err(serde::de::Error::duplicate_field("minAgeL2"));
                            }
                            min_age_l2__ = map_.next_value()?;
                        }
                        GeneratedField::MinAgeL3 => {
                            if min_age_l3__.is_some() {
                                return Err(serde::de::Error::duplicate_field("minAgeL3"));
                            }
                            min_age_l3__ = map_.next_value()?;
                        }
                        GeneratedField::MinAgeL4 => {
                            if min_age_l4__.is_some() {
                                return Err(serde::de::Error::duplicate_field("minAgeL4"));
                            }
                            min_age_l4__ = map_.next_value()?;
                        }
                        GeneratedField::MinLeaseCompletionBpsL3 => {
                            if min_lease_completion_bps_l3__.is_some() {
                                return Err(serde::de::Error::duplicate_field("minLeaseCompletionBpsL3"));
                            }
                            min_lease_completion_bps_l3__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::MinLeaseCompletionBpsL4 => {
                            if min_lease_completion_bps_l4__.is_some() {
                                return Err(serde::de::Error::duplicate_field("minLeaseCompletionBpsL4"));
                            }
                            min_lease_completion_bps_l4__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::CleanHistoryWindowL3 => {
                            if clean_history_window_l3__.is_some() {
                                return Err(serde::de::Error::duplicate_field("cleanHistoryWindowL3"));
                            }
                            clean_history_window_l3__ = map_.next_value()?;
                        }
                        GeneratedField::CleanHistoryWindowL4 => {
                            if clean_history_window_l4__.is_some() {
                                return Err(serde::de::Error::duplicate_field("cleanHistoryWindowL4"));
                            }
                            clean_history_window_l4__ = map_.next_value()?;
                        }
                        GeneratedField::MinL3DurationForL4 => {
                            if min_l3_duration_for_l4__.is_some() {
                                return Err(serde::de::Error::duplicate_field("minL3DurationForL4"));
                            }
                            min_l3_duration_for_l4__ = map_.next_value()?;
                        }
                        GeneratedField::MinLeasesForCompletionRate => {
                            if min_leases_for_completion_rate__.is_some() {
                                return Err(serde::de::Error::duplicate_field("minLeasesForCompletionRate"));
                            }
                            min_leases_for_completion_rate__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::MaxEndblockerAttestationExpiries => {
                            if max_endblocker_attestation_expiries__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maxEndblockerAttestationExpiries"));
                            }
                            max_endblocker_attestation_expiries__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::MaxEndblockerSnapshotSuspensions => {
                            if max_endblocker_snapshot_suspensions__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maxEndblockerSnapshotSuspensions"));
                            }
                            max_endblocker_snapshot_suspensions__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::MaxEndblockerUnbondingCompletions => {
                            if max_endblocker_unbonding_completions__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maxEndblockerUnbondingCompletions"));
                            }
                            max_endblocker_unbonding_completions__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::MaxEndblockerDiscrepancyTimeouts => {
                            if max_endblocker_discrepancy_timeouts__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maxEndblockerDiscrepancyTimeouts"));
                            }
                            max_endblocker_discrepancy_timeouts__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::MaxEndblockerAuditEscrowExpiries => {
                            if max_endblocker_audit_escrow_expiries__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maxEndblockerAuditEscrowExpiries"));
                            }
                            max_endblocker_audit_escrow_expiries__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::MaxEndblockerGraceExpiries => {
                            if max_endblocker_grace_expiries__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maxEndblockerGraceExpiries"));
                            }
                            max_endblocker_grace_expiries__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::DiscrepancyResolutionTimeout => {
                            if discrepancy_resolution_timeout__.is_some() {
                                return Err(serde::de::Error::duplicate_field("discrepancyResolutionTimeout"));
                            }
                            discrepancy_resolution_timeout__ = map_.next_value()?;
                        }
                        GeneratedField::AttestationDeposit => {
                            if attestation_deposit__.is_some() {
                                return Err(serde::de::Error::duplicate_field("attestationDeposit"));
                            }
                            attestation_deposit__ = map_.next_value()?;
                        }
                        GeneratedField::DiscrepancyGracePeriod => {
                            if discrepancy_grace_period__.is_some() {
                                return Err(serde::de::Error::duplicate_field("discrepancyGracePeriod"));
                            }
                            discrepancy_grace_period__ = map_.next_value()?;
                        }
                        GeneratedField::ProviderAuditDeposit => {
                            if provider_audit_deposit__.is_some() {
                                return Err(serde::de::Error::duplicate_field("providerAuditDeposit"));
                            }
                            provider_audit_deposit__ = map_.next_value()?;
                        }
                        GeneratedField::VerificationModuleActive => {
                            if verification_module_active__.is_some() {
                                return Err(serde::de::Error::duplicate_field("verificationModuleActive"));
                            }
                            verification_module_active__ = Some(map_.next_value()?);
                        }
                        GeneratedField::ContactResponseCriticalL1 => {
                            if contact_response_critical_l1__.is_some() {
                                return Err(serde::de::Error::duplicate_field("contactResponseCriticalL1"));
                            }
                            contact_response_critical_l1__ = map_.next_value()?;
                        }
                        GeneratedField::ContactResponseCriticalL2 => {
                            if contact_response_critical_l2__.is_some() {
                                return Err(serde::de::Error::duplicate_field("contactResponseCriticalL2"));
                            }
                            contact_response_critical_l2__ = map_.next_value()?;
                        }
                        GeneratedField::ContactResponseCriticalL3 => {
                            if contact_response_critical_l3__.is_some() {
                                return Err(serde::de::Error::duplicate_field("contactResponseCriticalL3"));
                            }
                            contact_response_critical_l3__ = map_.next_value()?;
                        }
                        GeneratedField::ContactResponseCriticalL4 => {
                            if contact_response_critical_l4__.is_some() {
                                return Err(serde::de::Error::duplicate_field("contactResponseCriticalL4"));
                            }
                            contact_response_critical_l4__ = map_.next_value()?;
                        }
                        GeneratedField::ContactResponseStandardL1 => {
                            if contact_response_standard_l1__.is_some() {
                                return Err(serde::de::Error::duplicate_field("contactResponseStandardL1"));
                            }
                            contact_response_standard_l1__ = map_.next_value()?;
                        }
                        GeneratedField::ContactResponseStandardL2 => {
                            if contact_response_standard_l2__.is_some() {
                                return Err(serde::de::Error::duplicate_field("contactResponseStandardL2"));
                            }
                            contact_response_standard_l2__ = map_.next_value()?;
                        }
                        GeneratedField::ContactResponseStandardL3 => {
                            if contact_response_standard_l3__.is_some() {
                                return Err(serde::de::Error::duplicate_field("contactResponseStandardL3"));
                            }
                            contact_response_standard_l3__ = map_.next_value()?;
                        }
                        GeneratedField::ContactResponseStandardL4 => {
                            if contact_response_standard_l4__.is_some() {
                                return Err(serde::de::Error::duplicate_field("contactResponseStandardL4"));
                            }
                            contact_response_standard_l4__ = map_.next_value()?;
                        }
                    }
                }
                Ok(Params {
                    bond_l1: bond_l1__,
                    bond_l2: bond_l2__,
                    bond_l3: bond_l3__,
                    bond_l4: bond_l4__,
                    ttl_l1: ttl_l1__,
                    ttl_l2: ttl_l2__,
                    ttl_l3: ttl_l3__,
                    ttl_l4: ttl_l4__,
                    min_fee_l1: min_fee_l1__,
                    min_fee_l2: min_fee_l2__,
                    min_fee_l3: min_fee_l3__,
                    min_fee_l4: min_fee_l4__,
                    discrepancy_threshold: discrepancy_threshold__.unwrap_or_default(),
                    auditor_unbonding_period: auditor_unbonding_period__,
                    renewal_period_l1: renewal_period_l1__,
                    renewal_period_l2: renewal_period_l2__,
                    renewal_period_l3: renewal_period_l3__,
                    renewal_period_l4: renewal_period_l4__,
                    snapshot_hash_interval: snapshot_hash_interval__,
                    max_snapshot_age: max_snapshot_age__,
                    bond_gpu_l2: bond_gpu_l2__,
                    bond_gpu_l3: bond_gpu_l3__,
                    bond_gpu_l4: bond_gpu_l4__,
                    bond_vcpu_l2: bond_vcpu_l2__,
                    bond_vcpu_l3: bond_vcpu_l3__,
                    bond_vcpu_l4: bond_vcpu_l4__,
                    bond_mem_gb_l2: bond_mem_gb_l2__,
                    bond_mem_gb_l3: bond_mem_gb_l3__,
                    bond_mem_gb_l4: bond_mem_gb_l4__,
                    bond_storage_tb_l2: bond_storage_tb_l2__,
                    bond_storage_tb_l3: bond_storage_tb_l3__,
                    bond_storage_tb_l4: bond_storage_tb_l4__,
                    provider_bond_unbonding_period: provider_bond_unbonding_period__,
                    min_age_l2: min_age_l2__,
                    min_age_l3: min_age_l3__,
                    min_age_l4: min_age_l4__,
                    min_lease_completion_bps_l3: min_lease_completion_bps_l3__.unwrap_or_default(),
                    min_lease_completion_bps_l4: min_lease_completion_bps_l4__.unwrap_or_default(),
                    clean_history_window_l3: clean_history_window_l3__,
                    clean_history_window_l4: clean_history_window_l4__,
                    min_l3_duration_for_l4: min_l3_duration_for_l4__,
                    min_leases_for_completion_rate: min_leases_for_completion_rate__.unwrap_or_default(),
                    max_endblocker_attestation_expiries: max_endblocker_attestation_expiries__.unwrap_or_default(),
                    max_endblocker_snapshot_suspensions: max_endblocker_snapshot_suspensions__.unwrap_or_default(),
                    max_endblocker_unbonding_completions: max_endblocker_unbonding_completions__.unwrap_or_default(),
                    max_endblocker_discrepancy_timeouts: max_endblocker_discrepancy_timeouts__.unwrap_or_default(),
                    max_endblocker_audit_escrow_expiries: max_endblocker_audit_escrow_expiries__.unwrap_or_default(),
                    max_endblocker_grace_expiries: max_endblocker_grace_expiries__.unwrap_or_default(),
                    discrepancy_resolution_timeout: discrepancy_resolution_timeout__,
                    attestation_deposit: attestation_deposit__,
                    discrepancy_grace_period: discrepancy_grace_period__,
                    provider_audit_deposit: provider_audit_deposit__,
                    verification_module_active: verification_module_active__.unwrap_or_default(),
                    contact_response_critical_l1: contact_response_critical_l1__,
                    contact_response_critical_l2: contact_response_critical_l2__,
                    contact_response_critical_l3: contact_response_critical_l3__,
                    contact_response_critical_l4: contact_response_critical_l4__,
                    contact_response_standard_l1: contact_response_standard_l1__,
                    contact_response_standard_l2: contact_response_standard_l2__,
                    contact_response_standard_l3: contact_response_standard_l3__,
                    contact_response_standard_l4: contact_response_standard_l4__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.Params", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for ProviderBondRecord {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if self.bonded_amount.is_some() {
            len += 1;
        }
        if !self.unbonding_entries.is_empty() {
            len += 1;
        }
        if self.slashed {
            len += 1;
        }
        if self.last_slash_time.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.ProviderBondRecord", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if let Some(v) = self.bonded_amount.as_ref() {
            struct_ser.serialize_field("bondedAmount", v)?;
        }
        if !self.unbonding_entries.is_empty() {
            struct_ser.serialize_field("unbondingEntries", &self.unbonding_entries)?;
        }
        if self.slashed {
            struct_ser.serialize_field("slashed", &self.slashed)?;
        }
        if let Some(v) = self.last_slash_time.as_ref() {
            struct_ser.serialize_field("lastSlashTime", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for ProviderBondRecord {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "bonded_amount",
            "bondedAmount",
            "unbonding_entries",
            "unbondingEntries",
            "slashed",
            "last_slash_time",
            "lastSlashTime",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            BondedAmount,
            UnbondingEntries,
            Slashed,
            LastSlashTime,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "bondedAmount" | "bonded_amount" => Ok(GeneratedField::BondedAmount),
                            "unbondingEntries" | "unbonding_entries" => Ok(GeneratedField::UnbondingEntries),
                            "slashed" => Ok(GeneratedField::Slashed),
                            "lastSlashTime" | "last_slash_time" => Ok(GeneratedField::LastSlashTime),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = ProviderBondRecord;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.ProviderBondRecord")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<ProviderBondRecord, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut bonded_amount__ = None;
                let mut unbonding_entries__ = None;
                let mut slashed__ = None;
                let mut last_slash_time__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::BondedAmount => {
                            if bonded_amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bondedAmount"));
                            }
                            bonded_amount__ = map_.next_value()?;
                        }
                        GeneratedField::UnbondingEntries => {
                            if unbonding_entries__.is_some() {
                                return Err(serde::de::Error::duplicate_field("unbondingEntries"));
                            }
                            unbonding_entries__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Slashed => {
                            if slashed__.is_some() {
                                return Err(serde::de::Error::duplicate_field("slashed"));
                            }
                            slashed__ = Some(map_.next_value()?);
                        }
                        GeneratedField::LastSlashTime => {
                            if last_slash_time__.is_some() {
                                return Err(serde::de::Error::duplicate_field("lastSlashTime"));
                            }
                            last_slash_time__ = map_.next_value()?;
                        }
                    }
                }
                Ok(ProviderBondRecord {
                    provider: provider__.unwrap_or_default(),
                    bonded_amount: bonded_amount__,
                    unbonding_entries: unbonding_entries__.unwrap_or_default(),
                    slashed: slashed__.unwrap_or_default(),
                    last_slash_time: last_slash_time__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.ProviderBondRecord", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for ProviderBondSlashReason {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "provider_bond_slash_reason_unspecified",
            Self::ResourceMisrepresentation => "provider_bond_slash_reason_resource_misrepresentation",
            Self::CapacityOverstatement => "provider_bond_slash_reason_capacity_overstatement",
            Self::FraudulentSnapshot => "provider_bond_slash_reason_fraudulent_snapshot",
            Self::ProviderCompromise => "provider_bond_slash_reason_provider_compromise",
            Self::SlaBreach => "provider_bond_slash_reason_sla_breach",
            Self::NonCooperationDuringAudit => "provider_bond_slash_reason_non_cooperation_during_audit",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for ProviderBondSlashReason {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider_bond_slash_reason_unspecified",
            "provider_bond_slash_reason_resource_misrepresentation",
            "provider_bond_slash_reason_capacity_overstatement",
            "provider_bond_slash_reason_fraudulent_snapshot",
            "provider_bond_slash_reason_provider_compromise",
            "provider_bond_slash_reason_sla_breach",
            "provider_bond_slash_reason_non_cooperation_during_audit",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = ProviderBondSlashReason;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                write!(formatter, "expected one of: {:?}", &FIELDS)
            }

            fn visit_i64<E>(self, v: i64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Signed(v), &self)
                    })
            }

            fn visit_u64<E>(self, v: u64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Unsigned(v), &self)
                    })
            }

            fn visit_str<E>(self, value: &str) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                match value {
                    "provider_bond_slash_reason_unspecified" => Ok(ProviderBondSlashReason::Unspecified),
                    "provider_bond_slash_reason_resource_misrepresentation" => Ok(ProviderBondSlashReason::ResourceMisrepresentation),
                    "provider_bond_slash_reason_capacity_overstatement" => Ok(ProviderBondSlashReason::CapacityOverstatement),
                    "provider_bond_slash_reason_fraudulent_snapshot" => Ok(ProviderBondSlashReason::FraudulentSnapshot),
                    "provider_bond_slash_reason_provider_compromise" => Ok(ProviderBondSlashReason::ProviderCompromise),
                    "provider_bond_slash_reason_sla_breach" => Ok(ProviderBondSlashReason::SlaBreach),
                    "provider_bond_slash_reason_non_cooperation_during_audit" => Ok(ProviderBondSlashReason::NonCooperationDuringAudit),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for ProviderDepositStatus {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "provider_deposit_status_unspecified",
            Self::Escrowed => "provider_deposit_status_escrowed",
            Self::ReturnedToProvider => "provider_deposit_status_returned_to_provider",
            Self::Slashed => "provider_deposit_status_slashed",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for ProviderDepositStatus {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider_deposit_status_unspecified",
            "provider_deposit_status_escrowed",
            "provider_deposit_status_returned_to_provider",
            "provider_deposit_status_slashed",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = ProviderDepositStatus;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                write!(formatter, "expected one of: {:?}", &FIELDS)
            }

            fn visit_i64<E>(self, v: i64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Signed(v), &self)
                    })
            }

            fn visit_u64<E>(self, v: u64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Unsigned(v), &self)
                    })
            }

            fn visit_str<E>(self, value: &str) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                match value {
                    "provider_deposit_status_unspecified" => Ok(ProviderDepositStatus::Unspecified),
                    "provider_deposit_status_escrowed" => Ok(ProviderDepositStatus::Escrowed),
                    "provider_deposit_status_returned_to_provider" => Ok(ProviderDepositStatus::ReturnedToProvider),
                    "provider_deposit_status_slashed" => Ok(ProviderDepositStatus::Slashed),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for ProviderSnapshotRecord {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if !self.snapshot_hash.is_empty() {
            len += 1;
        }
        if self.resource_summary.is_some() {
            len += 1;
        }
        if self.posted_at.is_some() {
            len += 1;
        }
        if self.snapshot_timestamp.is_some() {
            len += 1;
        }
        if self.compliance_deadline.is_some() {
            len += 1;
        }
        if self.suspended {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.ProviderSnapshotRecord", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if !self.snapshot_hash.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("snapshotHash", pbjson::private::base64::encode(&self.snapshot_hash).as_str())?;
        }
        if let Some(v) = self.resource_summary.as_ref() {
            struct_ser.serialize_field("resourceSummary", v)?;
        }
        if let Some(v) = self.posted_at.as_ref() {
            struct_ser.serialize_field("postedAt", v)?;
        }
        if let Some(v) = self.snapshot_timestamp.as_ref() {
            struct_ser.serialize_field("snapshotTimestamp", v)?;
        }
        if let Some(v) = self.compliance_deadline.as_ref() {
            struct_ser.serialize_field("complianceDeadline", v)?;
        }
        if self.suspended {
            struct_ser.serialize_field("suspended", &self.suspended)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for ProviderSnapshotRecord {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "snapshot_hash",
            "snapshotHash",
            "resource_summary",
            "resourceSummary",
            "posted_at",
            "postedAt",
            "snapshot_timestamp",
            "snapshotTimestamp",
            "compliance_deadline",
            "complianceDeadline",
            "suspended",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            SnapshotHash,
            ResourceSummary,
            PostedAt,
            SnapshotTimestamp,
            ComplianceDeadline,
            Suspended,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "snapshotHash" | "snapshot_hash" => Ok(GeneratedField::SnapshotHash),
                            "resourceSummary" | "resource_summary" => Ok(GeneratedField::ResourceSummary),
                            "postedAt" | "posted_at" => Ok(GeneratedField::PostedAt),
                            "snapshotTimestamp" | "snapshot_timestamp" => Ok(GeneratedField::SnapshotTimestamp),
                            "complianceDeadline" | "compliance_deadline" => Ok(GeneratedField::ComplianceDeadline),
                            "suspended" => Ok(GeneratedField::Suspended),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = ProviderSnapshotRecord;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.ProviderSnapshotRecord")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<ProviderSnapshotRecord, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut snapshot_hash__ = None;
                let mut resource_summary__ = None;
                let mut posted_at__ = None;
                let mut snapshot_timestamp__ = None;
                let mut compliance_deadline__ = None;
                let mut suspended__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::SnapshotHash => {
                            if snapshot_hash__.is_some() {
                                return Err(serde::de::Error::duplicate_field("snapshotHash"));
                            }
                            snapshot_hash__ = 
                                Some(map_.next_value::<::pbjson::private::BytesDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::ResourceSummary => {
                            if resource_summary__.is_some() {
                                return Err(serde::de::Error::duplicate_field("resourceSummary"));
                            }
                            resource_summary__ = map_.next_value()?;
                        }
                        GeneratedField::PostedAt => {
                            if posted_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("postedAt"));
                            }
                            posted_at__ = map_.next_value()?;
                        }
                        GeneratedField::SnapshotTimestamp => {
                            if snapshot_timestamp__.is_some() {
                                return Err(serde::de::Error::duplicate_field("snapshotTimestamp"));
                            }
                            snapshot_timestamp__ = map_.next_value()?;
                        }
                        GeneratedField::ComplianceDeadline => {
                            if compliance_deadline__.is_some() {
                                return Err(serde::de::Error::duplicate_field("complianceDeadline"));
                            }
                            compliance_deadline__ = map_.next_value()?;
                        }
                        GeneratedField::Suspended => {
                            if suspended__.is_some() {
                                return Err(serde::de::Error::duplicate_field("suspended"));
                            }
                            suspended__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(ProviderSnapshotRecord {
                    provider: provider__.unwrap_or_default(),
                    snapshot_hash: snapshot_hash__.unwrap_or_default(),
                    resource_summary: resource_summary__,
                    posted_at: posted_at__,
                    snapshot_timestamp: snapshot_timestamp__,
                    compliance_deadline: compliance_deadline__,
                    suspended: suspended__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.ProviderSnapshotRecord", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for ProviderVerificationGraceRecord {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.id != 0 {
            len += 1;
        }
        if !self.provider.is_empty() {
            len += 1;
        }
        if self.preserved_tier != 0 {
            len += 1;
        }
        if !self.source_discrepancy_ids.is_empty() {
            len += 1;
        }
        if self.started_at.is_some() {
            len += 1;
        }
        if self.expires_at.is_some() {
            len += 1;
        }
        if self.status != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.ProviderVerificationGraceRecord", len)?;
        if self.id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("id", ToString::to_string(&self.id).as_str())?;
        }
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if self.preserved_tier != 0 {
            let v = VerificationTier::try_from(self.preserved_tier)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.preserved_tier)))?;
            struct_ser.serialize_field("preservedTier", &v)?;
        }
        if !self.source_discrepancy_ids.is_empty() {
            struct_ser.serialize_field("sourceDiscrepancyIds", &self.source_discrepancy_ids.iter().map(ToString::to_string).collect::<Vec<_>>())?;
        }
        if let Some(v) = self.started_at.as_ref() {
            struct_ser.serialize_field("startedAt", v)?;
        }
        if let Some(v) = self.expires_at.as_ref() {
            struct_ser.serialize_field("expiresAt", v)?;
        }
        if self.status != 0 {
            let v = VerificationGraceStatus::try_from(self.status)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.status)))?;
            struct_ser.serialize_field("status", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for ProviderVerificationGraceRecord {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "id",
            "provider",
            "preserved_tier",
            "preservedTier",
            "source_discrepancy_ids",
            "sourceDiscrepancyIds",
            "started_at",
            "startedAt",
            "expires_at",
            "expiresAt",
            "status",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Id,
            Provider,
            PreservedTier,
            SourceDiscrepancyIds,
            StartedAt,
            ExpiresAt,
            Status,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "id" => Ok(GeneratedField::Id),
                            "provider" => Ok(GeneratedField::Provider),
                            "preservedTier" | "preserved_tier" => Ok(GeneratedField::PreservedTier),
                            "sourceDiscrepancyIds" | "source_discrepancy_ids" => Ok(GeneratedField::SourceDiscrepancyIds),
                            "startedAt" | "started_at" => Ok(GeneratedField::StartedAt),
                            "expiresAt" | "expires_at" => Ok(GeneratedField::ExpiresAt),
                            "status" => Ok(GeneratedField::Status),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = ProviderVerificationGraceRecord;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.ProviderVerificationGraceRecord")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<ProviderVerificationGraceRecord, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                let mut provider__ = None;
                let mut preserved_tier__ = None;
                let mut source_discrepancy_ids__ = None;
                let mut started_at__ = None;
                let mut expires_at__ = None;
                let mut status__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::PreservedTier => {
                            if preserved_tier__.is_some() {
                                return Err(serde::de::Error::duplicate_field("preservedTier"));
                            }
                            preserved_tier__ = Some(map_.next_value::<VerificationTier>()? as i32);
                        }
                        GeneratedField::SourceDiscrepancyIds => {
                            if source_discrepancy_ids__.is_some() {
                                return Err(serde::de::Error::duplicate_field("sourceDiscrepancyIds"));
                            }
                            source_discrepancy_ids__ = 
                                Some(map_.next_value::<Vec<::pbjson::private::NumberDeserialize<_>>>()?
                                    .into_iter().map(|x| x.0).collect())
                            ;
                        }
                        GeneratedField::StartedAt => {
                            if started_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("startedAt"));
                            }
                            started_at__ = map_.next_value()?;
                        }
                        GeneratedField::ExpiresAt => {
                            if expires_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("expiresAt"));
                            }
                            expires_at__ = map_.next_value()?;
                        }
                        GeneratedField::Status => {
                            if status__.is_some() {
                                return Err(serde::de::Error::duplicate_field("status"));
                            }
                            status__ = Some(map_.next_value::<VerificationGraceStatus>()? as i32);
                        }
                    }
                }
                Ok(ProviderVerificationGraceRecord {
                    id: id__.unwrap_or_default(),
                    provider: provider__.unwrap_or_default(),
                    preserved_tier: preserved_tier__.unwrap_or_default(),
                    source_discrepancy_ids: source_discrepancy_ids__.unwrap_or_default(),
                    started_at: started_at__,
                    expires_at: expires_at__,
                    status: status__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.ProviderVerificationGraceRecord", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryAttestationRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if !self.auditor.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryAttestationRequest", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryAttestationRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "auditor",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            Auditor,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "auditor" => Ok(GeneratedField::Auditor),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryAttestationRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryAttestationRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryAttestationRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut auditor__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(QueryAttestationRequest {
                    provider: provider__.unwrap_or_default(),
                    auditor: auditor__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryAttestationRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryAttestationResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.attestation.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryAttestationResponse", len)?;
        if let Some(v) = self.attestation.as_ref() {
            struct_ser.serialize_field("attestation", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryAttestationResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "attestation",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Attestation,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "attestation" => Ok(GeneratedField::Attestation),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryAttestationResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryAttestationResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryAttestationResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut attestation__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Attestation => {
                            if attestation__.is_some() {
                                return Err(serde::de::Error::duplicate_field("attestation"));
                            }
                            attestation__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryAttestationResponse {
                    attestation: attestation__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryAttestationResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryAuditEscrowRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.id != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryAuditEscrowRequest", len)?;
        if self.id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("id", ToString::to_string(&self.id).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryAuditEscrowRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "id",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Id,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "id" => Ok(GeneratedField::Id),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryAuditEscrowRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryAuditEscrowRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryAuditEscrowRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(QueryAuditEscrowRequest {
                    id: id__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryAuditEscrowRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryAuditEscrowResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.escrow.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryAuditEscrowResponse", len)?;
        if let Some(v) = self.escrow.as_ref() {
            struct_ser.serialize_field("escrow", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryAuditEscrowResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "escrow",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Escrow,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "escrow" => Ok(GeneratedField::Escrow),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryAuditEscrowResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryAuditEscrowResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryAuditEscrowResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut escrow__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Escrow => {
                            if escrow__.is_some() {
                                return Err(serde::de::Error::duplicate_field("escrow"));
                            }
                            escrow__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryAuditEscrowResponse {
                    escrow: escrow__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryAuditEscrowResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryAuditorAttestationsRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.auditor.is_empty() {
            len += 1;
        }
        if self.pagination.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryAuditorAttestationsRequest", len)?;
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryAuditorAttestationsRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "auditor",
            "pagination",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Auditor,
            Pagination,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "auditor" => Ok(GeneratedField::Auditor),
                            "pagination" => Ok(GeneratedField::Pagination),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryAuditorAttestationsRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryAuditorAttestationsRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryAuditorAttestationsRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut auditor__ = None;
                let mut pagination__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Pagination => {
                            if pagination__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pagination"));
                            }
                            pagination__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryAuditorAttestationsRequest {
                    auditor: auditor__.unwrap_or_default(),
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryAuditorAttestationsRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryAuditorAttestationsResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.attestations.is_empty() {
            len += 1;
        }
        if self.pagination.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryAuditorAttestationsResponse", len)?;
        if !self.attestations.is_empty() {
            struct_ser.serialize_field("attestations", &self.attestations)?;
        }
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryAuditorAttestationsResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "attestations",
            "pagination",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Attestations,
            Pagination,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "attestations" => Ok(GeneratedField::Attestations),
                            "pagination" => Ok(GeneratedField::Pagination),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryAuditorAttestationsResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryAuditorAttestationsResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryAuditorAttestationsResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut attestations__ = None;
                let mut pagination__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Attestations => {
                            if attestations__.is_some() {
                                return Err(serde::de::Error::duplicate_field("attestations"));
                            }
                            attestations__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Pagination => {
                            if pagination__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pagination"));
                            }
                            pagination__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryAuditorAttestationsResponse {
                    attestations: attestations__.unwrap_or_default(),
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryAuditorAttestationsResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryAuditorRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.auditor.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryAuditorRequest", len)?;
        if !self.auditor.is_empty() {
            struct_ser.serialize_field("auditor", &self.auditor)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryAuditorRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "auditor",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Auditor,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "auditor" => Ok(GeneratedField::Auditor),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryAuditorRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryAuditorRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryAuditorRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut auditor__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(QueryAuditorRequest {
                    auditor: auditor__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryAuditorRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryAuditorResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.auditor.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryAuditorResponse", len)?;
        if let Some(v) = self.auditor.as_ref() {
            struct_ser.serialize_field("auditor", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryAuditorResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "auditor",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Auditor,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "auditor" => Ok(GeneratedField::Auditor),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryAuditorResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryAuditorResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryAuditorResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut auditor__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Auditor => {
                            if auditor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditor"));
                            }
                            auditor__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryAuditorResponse {
                    auditor: auditor__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryAuditorResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryAuditorsRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.status_filter != 0 {
            len += 1;
        }
        if self.pagination.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryAuditorsRequest", len)?;
        if self.status_filter != 0 {
            let v = AuditorStatus::try_from(self.status_filter)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.status_filter)))?;
            struct_ser.serialize_field("statusFilter", &v)?;
        }
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryAuditorsRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "status_filter",
            "statusFilter",
            "pagination",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            StatusFilter,
            Pagination,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "statusFilter" | "status_filter" => Ok(GeneratedField::StatusFilter),
                            "pagination" => Ok(GeneratedField::Pagination),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryAuditorsRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryAuditorsRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryAuditorsRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut status_filter__ = None;
                let mut pagination__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::StatusFilter => {
                            if status_filter__.is_some() {
                                return Err(serde::de::Error::duplicate_field("statusFilter"));
                            }
                            status_filter__ = Some(map_.next_value::<AuditorStatus>()? as i32);
                        }
                        GeneratedField::Pagination => {
                            if pagination__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pagination"));
                            }
                            pagination__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryAuditorsRequest {
                    status_filter: status_filter__.unwrap_or_default(),
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryAuditorsRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryAuditorsResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.auditors.is_empty() {
            len += 1;
        }
        if self.pagination.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryAuditorsResponse", len)?;
        if !self.auditors.is_empty() {
            struct_ser.serialize_field("auditors", &self.auditors)?;
        }
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryAuditorsResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "auditors",
            "pagination",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Auditors,
            Pagination,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "auditors" => Ok(GeneratedField::Auditors),
                            "pagination" => Ok(GeneratedField::Pagination),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryAuditorsResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryAuditorsResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryAuditorsResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut auditors__ = None;
                let mut pagination__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Auditors => {
                            if auditors__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditors"));
                            }
                            auditors__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Pagination => {
                            if pagination__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pagination"));
                            }
                            pagination__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryAuditorsResponse {
                    auditors: auditors__.unwrap_or_default(),
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryAuditorsResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryDiscrepanciesRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.status_filter != 0 {
            len += 1;
        }
        if self.pagination.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryDiscrepanciesRequest", len)?;
        if self.status_filter != 0 {
            let v = DiscrepancyStatus::try_from(self.status_filter)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.status_filter)))?;
            struct_ser.serialize_field("statusFilter", &v)?;
        }
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryDiscrepanciesRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "status_filter",
            "statusFilter",
            "pagination",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            StatusFilter,
            Pagination,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "statusFilter" | "status_filter" => Ok(GeneratedField::StatusFilter),
                            "pagination" => Ok(GeneratedField::Pagination),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryDiscrepanciesRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryDiscrepanciesRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryDiscrepanciesRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut status_filter__ = None;
                let mut pagination__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::StatusFilter => {
                            if status_filter__.is_some() {
                                return Err(serde::de::Error::duplicate_field("statusFilter"));
                            }
                            status_filter__ = Some(map_.next_value::<DiscrepancyStatus>()? as i32);
                        }
                        GeneratedField::Pagination => {
                            if pagination__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pagination"));
                            }
                            pagination__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryDiscrepanciesRequest {
                    status_filter: status_filter__.unwrap_or_default(),
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryDiscrepanciesRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryDiscrepanciesResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.discrepancies.is_empty() {
            len += 1;
        }
        if self.pagination.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryDiscrepanciesResponse", len)?;
        if !self.discrepancies.is_empty() {
            struct_ser.serialize_field("discrepancies", &self.discrepancies)?;
        }
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryDiscrepanciesResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "discrepancies",
            "pagination",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Discrepancies,
            Pagination,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "discrepancies" => Ok(GeneratedField::Discrepancies),
                            "pagination" => Ok(GeneratedField::Pagination),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryDiscrepanciesResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryDiscrepanciesResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryDiscrepanciesResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut discrepancies__ = None;
                let mut pagination__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Discrepancies => {
                            if discrepancies__.is_some() {
                                return Err(serde::de::Error::duplicate_field("discrepancies"));
                            }
                            discrepancies__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Pagination => {
                            if pagination__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pagination"));
                            }
                            pagination__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryDiscrepanciesResponse {
                    discrepancies: discrepancies__.unwrap_or_default(),
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryDiscrepanciesResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryDiscrepancyRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.id != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryDiscrepancyRequest", len)?;
        if self.id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("id", ToString::to_string(&self.id).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryDiscrepancyRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "id",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Id,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "id" => Ok(GeneratedField::Id),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryDiscrepancyRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryDiscrepancyRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryDiscrepancyRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(QueryDiscrepancyRequest {
                    id: id__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryDiscrepancyRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryDiscrepancyResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.discrepancy.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryDiscrepancyResponse", len)?;
        if let Some(v) = self.discrepancy.as_ref() {
            struct_ser.serialize_field("discrepancy", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryDiscrepancyResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "discrepancy",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Discrepancy,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "discrepancy" => Ok(GeneratedField::Discrepancy),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryDiscrepancyResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryDiscrepancyResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryDiscrepancyResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut discrepancy__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Discrepancy => {
                            if discrepancy__.is_some() {
                                return Err(serde::de::Error::duplicate_field("discrepancy"));
                            }
                            discrepancy__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryDiscrepancyResponse {
                    discrepancy: discrepancy__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryDiscrepancyResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryParamsRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.verification.v1.QueryParamsRequest", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryParamsRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                            Err(serde::de::Error::unknown_field(value, FIELDS))
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryParamsRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryParamsRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryParamsRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(QueryParamsRequest {
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryParamsRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryParamsResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.params.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryParamsResponse", len)?;
        if let Some(v) = self.params.as_ref() {
            struct_ser.serialize_field("params", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryParamsResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "params",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Params,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "params" => Ok(GeneratedField::Params),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryParamsResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryParamsResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryParamsResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut params__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Params => {
                            if params__.is_some() {
                                return Err(serde::de::Error::duplicate_field("params"));
                            }
                            params__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryParamsResponse {
                    params: params__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryParamsResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryProviderAttestationsRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if self.status_filter != 0 {
            len += 1;
        }
        if self.pagination.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryProviderAttestationsRequest", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if self.status_filter != 0 {
            let v = AttestationStatus::try_from(self.status_filter)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.status_filter)))?;
            struct_ser.serialize_field("statusFilter", &v)?;
        }
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryProviderAttestationsRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "status_filter",
            "statusFilter",
            "pagination",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            StatusFilter,
            Pagination,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "statusFilter" | "status_filter" => Ok(GeneratedField::StatusFilter),
                            "pagination" => Ok(GeneratedField::Pagination),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryProviderAttestationsRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryProviderAttestationsRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryProviderAttestationsRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut status_filter__ = None;
                let mut pagination__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::StatusFilter => {
                            if status_filter__.is_some() {
                                return Err(serde::de::Error::duplicate_field("statusFilter"));
                            }
                            status_filter__ = Some(map_.next_value::<AttestationStatus>()? as i32);
                        }
                        GeneratedField::Pagination => {
                            if pagination__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pagination"));
                            }
                            pagination__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryProviderAttestationsRequest {
                    provider: provider__.unwrap_or_default(),
                    status_filter: status_filter__.unwrap_or_default(),
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryProviderAttestationsRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryProviderAttestationsResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.attestations.is_empty() {
            len += 1;
        }
        if self.pagination.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryProviderAttestationsResponse", len)?;
        if !self.attestations.is_empty() {
            struct_ser.serialize_field("attestations", &self.attestations)?;
        }
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryProviderAttestationsResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "attestations",
            "pagination",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Attestations,
            Pagination,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "attestations" => Ok(GeneratedField::Attestations),
                            "pagination" => Ok(GeneratedField::Pagination),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryProviderAttestationsResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryProviderAttestationsResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryProviderAttestationsResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut attestations__ = None;
                let mut pagination__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Attestations => {
                            if attestations__.is_some() {
                                return Err(serde::de::Error::duplicate_field("attestations"));
                            }
                            attestations__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Pagination => {
                            if pagination__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pagination"));
                            }
                            pagination__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryProviderAttestationsResponse {
                    attestations: attestations__.unwrap_or_default(),
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryProviderAttestationsResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryProviderAuditEscrowsRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        if self.status_filter != 0 {
            len += 1;
        }
        if self.pagination.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryProviderAuditEscrowsRequest", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if self.status_filter != 0 {
            let v = AuditEscrowStatus::try_from(self.status_filter)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.status_filter)))?;
            struct_ser.serialize_field("statusFilter", &v)?;
        }
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryProviderAuditEscrowsRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "status_filter",
            "statusFilter",
            "pagination",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            StatusFilter,
            Pagination,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            "statusFilter" | "status_filter" => Ok(GeneratedField::StatusFilter),
                            "pagination" => Ok(GeneratedField::Pagination),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryProviderAuditEscrowsRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryProviderAuditEscrowsRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryProviderAuditEscrowsRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut status_filter__ = None;
                let mut pagination__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::StatusFilter => {
                            if status_filter__.is_some() {
                                return Err(serde::de::Error::duplicate_field("statusFilter"));
                            }
                            status_filter__ = Some(map_.next_value::<AuditEscrowStatus>()? as i32);
                        }
                        GeneratedField::Pagination => {
                            if pagination__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pagination"));
                            }
                            pagination__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryProviderAuditEscrowsRequest {
                    provider: provider__.unwrap_or_default(),
                    status_filter: status_filter__.unwrap_or_default(),
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryProviderAuditEscrowsRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryProviderAuditEscrowsResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.escrows.is_empty() {
            len += 1;
        }
        if self.pagination.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryProviderAuditEscrowsResponse", len)?;
        if !self.escrows.is_empty() {
            struct_ser.serialize_field("escrows", &self.escrows)?;
        }
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryProviderAuditEscrowsResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "escrows",
            "pagination",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Escrows,
            Pagination,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "escrows" => Ok(GeneratedField::Escrows),
                            "pagination" => Ok(GeneratedField::Pagination),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryProviderAuditEscrowsResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryProviderAuditEscrowsResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryProviderAuditEscrowsResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut escrows__ = None;
                let mut pagination__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Escrows => {
                            if escrows__.is_some() {
                                return Err(serde::de::Error::duplicate_field("escrows"));
                            }
                            escrows__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Pagination => {
                            if pagination__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pagination"));
                            }
                            pagination__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryProviderAuditEscrowsResponse {
                    escrows: escrows__.unwrap_or_default(),
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryProviderAuditEscrowsResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryProviderBondRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryProviderBondRequest", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryProviderBondRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryProviderBondRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryProviderBondRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryProviderBondRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(QueryProviderBondRequest {
                    provider: provider__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryProviderBondRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryProviderBondResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.bond.is_some() {
            len += 1;
        }
        if self.required_for_current_tier.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryProviderBondResponse", len)?;
        if let Some(v) = self.bond.as_ref() {
            struct_ser.serialize_field("bond", v)?;
        }
        if let Some(v) = self.required_for_current_tier.as_ref() {
            struct_ser.serialize_field("requiredForCurrentTier", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryProviderBondResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "bond",
            "required_for_current_tier",
            "requiredForCurrentTier",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Bond,
            RequiredForCurrentTier,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "bond" => Ok(GeneratedField::Bond),
                            "requiredForCurrentTier" | "required_for_current_tier" => Ok(GeneratedField::RequiredForCurrentTier),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryProviderBondResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryProviderBondResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryProviderBondResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut bond__ = None;
                let mut required_for_current_tier__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Bond => {
                            if bond__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bond"));
                            }
                            bond__ = map_.next_value()?;
                        }
                        GeneratedField::RequiredForCurrentTier => {
                            if required_for_current_tier__.is_some() {
                                return Err(serde::de::Error::duplicate_field("requiredForCurrentTier"));
                            }
                            required_for_current_tier__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryProviderBondResponse {
                    bond: bond__,
                    required_for_current_tier: required_for_current_tier__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryProviderBondResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryProviderSnapshotRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryProviderSnapshotRequest", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryProviderSnapshotRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryProviderSnapshotRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryProviderSnapshotRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryProviderSnapshotRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(QueryProviderSnapshotRequest {
                    provider: provider__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryProviderSnapshotRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryProviderSnapshotResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.snapshot.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryProviderSnapshotResponse", len)?;
        if let Some(v) = self.snapshot.as_ref() {
            struct_ser.serialize_field("snapshot", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryProviderSnapshotResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "snapshot",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Snapshot,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "snapshot" => Ok(GeneratedField::Snapshot),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryProviderSnapshotResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryProviderSnapshotResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryProviderSnapshotResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut snapshot__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Snapshot => {
                            if snapshot__.is_some() {
                                return Err(serde::de::Error::duplicate_field("snapshot"));
                            }
                            snapshot__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryProviderSnapshotResponse {
                    snapshot: snapshot__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryProviderSnapshotResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryProviderVerificationGraceRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.provider.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryProviderVerificationGraceRequest", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryProviderVerificationGraceRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "provider" => Ok(GeneratedField::Provider),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryProviderVerificationGraceRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryProviderVerificationGraceRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryProviderVerificationGraceRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(QueryProviderVerificationGraceRequest {
                    provider: provider__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryProviderVerificationGraceRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryProviderVerificationGraceResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.grace.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.QueryProviderVerificationGraceResponse", len)?;
        if let Some(v) = self.grace.as_ref() {
            struct_ser.serialize_field("grace", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryProviderVerificationGraceResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "grace",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Grace,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "grace" => Ok(GeneratedField::Grace),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryProviderVerificationGraceResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.QueryProviderVerificationGraceResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryProviderVerificationGraceResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut grace__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Grace => {
                            if grace__.is_some() {
                                return Err(serde::de::Error::duplicate_field("grace"));
                            }
                            grace__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryProviderVerificationGraceResponse {
                    grace: grace__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.QueryProviderVerificationGraceResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for ResourceSummary {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.total_gpus != 0 {
            len += 1;
        }
        if self.total_vcpus != 0 {
            len += 1;
        }
        if self.total_memory_mb != 0 {
            len += 1;
        }
        if self.total_storage_mb != 0 {
            len += 1;
        }
        if self.active_leases != 0 {
            len += 1;
        }
        if !self.software_version.is_empty() {
            len += 1;
        }
        if !self.software_signature.is_empty() {
            len += 1;
        }
        if self.software_identity.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.ResourceSummary", len)?;
        if self.total_gpus != 0 {
            struct_ser.serialize_field("totalGpus", &self.total_gpus)?;
        }
        if self.total_vcpus != 0 {
            struct_ser.serialize_field("totalVcpus", &self.total_vcpus)?;
        }
        if self.total_memory_mb != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("totalMemoryMb", ToString::to_string(&self.total_memory_mb).as_str())?;
        }
        if self.total_storage_mb != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("totalStorageMb", ToString::to_string(&self.total_storage_mb).as_str())?;
        }
        if self.active_leases != 0 {
            struct_ser.serialize_field("activeLeases", &self.active_leases)?;
        }
        if !self.software_version.is_empty() {
            struct_ser.serialize_field("softwareVersion", &self.software_version)?;
        }
        if !self.software_signature.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("softwareSignature", pbjson::private::base64::encode(&self.software_signature).as_str())?;
        }
        if let Some(v) = self.software_identity.as_ref() {
            struct_ser.serialize_field("softwareIdentity", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for ResourceSummary {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "total_gpus",
            "totalGpus",
            "total_vcpus",
            "totalVcpus",
            "total_memory_mb",
            "totalMemoryMb",
            "total_storage_mb",
            "totalStorageMb",
            "active_leases",
            "activeLeases",
            "software_version",
            "softwareVersion",
            "software_signature",
            "softwareSignature",
            "software_identity",
            "softwareIdentity",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            TotalGpus,
            TotalVcpus,
            TotalMemoryMb,
            TotalStorageMb,
            ActiveLeases,
            SoftwareVersion,
            SoftwareSignature,
            SoftwareIdentity,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "totalGpus" | "total_gpus" => Ok(GeneratedField::TotalGpus),
                            "totalVcpus" | "total_vcpus" => Ok(GeneratedField::TotalVcpus),
                            "totalMemoryMb" | "total_memory_mb" => Ok(GeneratedField::TotalMemoryMb),
                            "totalStorageMb" | "total_storage_mb" => Ok(GeneratedField::TotalStorageMb),
                            "activeLeases" | "active_leases" => Ok(GeneratedField::ActiveLeases),
                            "softwareVersion" | "software_version" => Ok(GeneratedField::SoftwareVersion),
                            "softwareSignature" | "software_signature" => Ok(GeneratedField::SoftwareSignature),
                            "softwareIdentity" | "software_identity" => Ok(GeneratedField::SoftwareIdentity),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = ResourceSummary;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.ResourceSummary")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<ResourceSummary, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut total_gpus__ = None;
                let mut total_vcpus__ = None;
                let mut total_memory_mb__ = None;
                let mut total_storage_mb__ = None;
                let mut active_leases__ = None;
                let mut software_version__ = None;
                let mut software_signature__ = None;
                let mut software_identity__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::TotalGpus => {
                            if total_gpus__.is_some() {
                                return Err(serde::de::Error::duplicate_field("totalGpus"));
                            }
                            total_gpus__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::TotalVcpus => {
                            if total_vcpus__.is_some() {
                                return Err(serde::de::Error::duplicate_field("totalVcpus"));
                            }
                            total_vcpus__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::TotalMemoryMb => {
                            if total_memory_mb__.is_some() {
                                return Err(serde::de::Error::duplicate_field("totalMemoryMb"));
                            }
                            total_memory_mb__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::TotalStorageMb => {
                            if total_storage_mb__.is_some() {
                                return Err(serde::de::Error::duplicate_field("totalStorageMb"));
                            }
                            total_storage_mb__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::ActiveLeases => {
                            if active_leases__.is_some() {
                                return Err(serde::de::Error::duplicate_field("activeLeases"));
                            }
                            active_leases__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::SoftwareVersion => {
                            if software_version__.is_some() {
                                return Err(serde::de::Error::duplicate_field("softwareVersion"));
                            }
                            software_version__ = Some(map_.next_value()?);
                        }
                        GeneratedField::SoftwareSignature => {
                            if software_signature__.is_some() {
                                return Err(serde::de::Error::duplicate_field("softwareSignature"));
                            }
                            software_signature__ = 
                                Some(map_.next_value::<::pbjson::private::BytesDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::SoftwareIdentity => {
                            if software_identity__.is_some() {
                                return Err(serde::de::Error::duplicate_field("softwareIdentity"));
                            }
                            software_identity__ = map_.next_value()?;
                        }
                    }
                }
                Ok(ResourceSummary {
                    total_gpus: total_gpus__.unwrap_or_default(),
                    total_vcpus: total_vcpus__.unwrap_or_default(),
                    total_memory_mb: total_memory_mb__.unwrap_or_default(),
                    total_storage_mb: total_storage_mb__.unwrap_or_default(),
                    active_leases: active_leases__.unwrap_or_default(),
                    software_version: software_version__.unwrap_or_default(),
                    software_signature: software_signature__.unwrap_or_default(),
                    software_identity: software_identity__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.ResourceSummary", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for SoftwareIdentity {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.version.is_empty() {
            len += 1;
        }
        if !self.artifact_ref.is_empty() {
            len += 1;
        }
        if !self.digest_algorithm.is_empty() {
            len += 1;
        }
        if !self.digest.is_empty() {
            len += 1;
        }
        if !self.signature_type.is_empty() {
            len += 1;
        }
        if !self.signature.is_empty() {
            len += 1;
        }
        if !self.signature_ref.is_empty() {
            len += 1;
        }
        if !self.public_key_ref.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.SoftwareIdentity", len)?;
        if !self.version.is_empty() {
            struct_ser.serialize_field("version", &self.version)?;
        }
        if !self.artifact_ref.is_empty() {
            struct_ser.serialize_field("artifactRef", &self.artifact_ref)?;
        }
        if !self.digest_algorithm.is_empty() {
            struct_ser.serialize_field("digestAlgorithm", &self.digest_algorithm)?;
        }
        if !self.digest.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("digest", pbjson::private::base64::encode(&self.digest).as_str())?;
        }
        if !self.signature_type.is_empty() {
            struct_ser.serialize_field("signatureType", &self.signature_type)?;
        }
        if !self.signature.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("signature", pbjson::private::base64::encode(&self.signature).as_str())?;
        }
        if !self.signature_ref.is_empty() {
            struct_ser.serialize_field("signatureRef", &self.signature_ref)?;
        }
        if !self.public_key_ref.is_empty() {
            struct_ser.serialize_field("publicKeyRef", &self.public_key_ref)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for SoftwareIdentity {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "version",
            "artifact_ref",
            "artifactRef",
            "digest_algorithm",
            "digestAlgorithm",
            "digest",
            "signature_type",
            "signatureType",
            "signature",
            "signature_ref",
            "signatureRef",
            "public_key_ref",
            "publicKeyRef",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Version,
            ArtifactRef,
            DigestAlgorithm,
            Digest,
            SignatureType,
            Signature,
            SignatureRef,
            PublicKeyRef,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "version" => Ok(GeneratedField::Version),
                            "artifactRef" | "artifact_ref" => Ok(GeneratedField::ArtifactRef),
                            "digestAlgorithm" | "digest_algorithm" => Ok(GeneratedField::DigestAlgorithm),
                            "digest" => Ok(GeneratedField::Digest),
                            "signatureType" | "signature_type" => Ok(GeneratedField::SignatureType),
                            "signature" => Ok(GeneratedField::Signature),
                            "signatureRef" | "signature_ref" => Ok(GeneratedField::SignatureRef),
                            "publicKeyRef" | "public_key_ref" => Ok(GeneratedField::PublicKeyRef),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = SoftwareIdentity;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.SoftwareIdentity")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<SoftwareIdentity, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut version__ = None;
                let mut artifact_ref__ = None;
                let mut digest_algorithm__ = None;
                let mut digest__ = None;
                let mut signature_type__ = None;
                let mut signature__ = None;
                let mut signature_ref__ = None;
                let mut public_key_ref__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Version => {
                            if version__.is_some() {
                                return Err(serde::de::Error::duplicate_field("version"));
                            }
                            version__ = Some(map_.next_value()?);
                        }
                        GeneratedField::ArtifactRef => {
                            if artifact_ref__.is_some() {
                                return Err(serde::de::Error::duplicate_field("artifactRef"));
                            }
                            artifact_ref__ = Some(map_.next_value()?);
                        }
                        GeneratedField::DigestAlgorithm => {
                            if digest_algorithm__.is_some() {
                                return Err(serde::de::Error::duplicate_field("digestAlgorithm"));
                            }
                            digest_algorithm__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Digest => {
                            if digest__.is_some() {
                                return Err(serde::de::Error::duplicate_field("digest"));
                            }
                            digest__ = 
                                Some(map_.next_value::<::pbjson::private::BytesDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::SignatureType => {
                            if signature_type__.is_some() {
                                return Err(serde::de::Error::duplicate_field("signatureType"));
                            }
                            signature_type__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Signature => {
                            if signature__.is_some() {
                                return Err(serde::de::Error::duplicate_field("signature"));
                            }
                            signature__ = 
                                Some(map_.next_value::<::pbjson::private::BytesDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::SignatureRef => {
                            if signature_ref__.is_some() {
                                return Err(serde::de::Error::duplicate_field("signatureRef"));
                            }
                            signature_ref__ = Some(map_.next_value()?);
                        }
                        GeneratedField::PublicKeyRef => {
                            if public_key_ref__.is_some() {
                                return Err(serde::de::Error::duplicate_field("publicKeyRef"));
                            }
                            public_key_ref__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(SoftwareIdentity {
                    version: version__.unwrap_or_default(),
                    artifact_ref: artifact_ref__.unwrap_or_default(),
                    digest_algorithm: digest_algorithm__.unwrap_or_default(),
                    digest: digest__.unwrap_or_default(),
                    signature_type: signature_type__.unwrap_or_default(),
                    signature: signature__.unwrap_or_default(),
                    signature_ref: signature_ref__.unwrap_or_default(),
                    public_key_ref: public_key_ref__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.SoftwareIdentity", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for UnbondingEntry {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.amount.is_some() {
            len += 1;
        }
        if self.completion_time.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.UnbondingEntry", len)?;
        if let Some(v) = self.amount.as_ref() {
            struct_ser.serialize_field("amount", v)?;
        }
        if let Some(v) = self.completion_time.as_ref() {
            struct_ser.serialize_field("completionTime", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for UnbondingEntry {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "amount",
            "completion_time",
            "completionTime",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Amount,
            CompletionTime,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "amount" => Ok(GeneratedField::Amount),
                            "completionTime" | "completion_time" => Ok(GeneratedField::CompletionTime),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = UnbondingEntry;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.UnbondingEntry")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<UnbondingEntry, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut amount__ = None;
                let mut completion_time__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Amount => {
                            if amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("amount"));
                            }
                            amount__ = map_.next_value()?;
                        }
                        GeneratedField::CompletionTime => {
                            if completion_time__.is_some() {
                                return Err(serde::de::Error::duplicate_field("completionTime"));
                            }
                            completion_time__ = map_.next_value()?;
                        }
                    }
                }
                Ok(UnbondingEntry {
                    amount: amount__,
                    completion_time: completion_time__,
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.UnbondingEntry", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for VerificationGraceStatus {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "verification_grace_status_unspecified",
            Self::Active => "verification_grace_status_active",
            Self::Expired => "verification_grace_status_expired",
            Self::Terminated => "verification_grace_status_terminated",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for VerificationGraceStatus {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "verification_grace_status_unspecified",
            "verification_grace_status_active",
            "verification_grace_status_expired",
            "verification_grace_status_terminated",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = VerificationGraceStatus;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                write!(formatter, "expected one of: {:?}", &FIELDS)
            }

            fn visit_i64<E>(self, v: i64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Signed(v), &self)
                    })
            }

            fn visit_u64<E>(self, v: u64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Unsigned(v), &self)
                    })
            }

            fn visit_str<E>(self, value: &str) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                match value {
                    "verification_grace_status_unspecified" => Ok(VerificationGraceStatus::Unspecified),
                    "verification_grace_status_active" => Ok(VerificationGraceStatus::Active),
                    "verification_grace_status_expired" => Ok(VerificationGraceStatus::Expired),
                    "verification_grace_status_terminated" => Ok(VerificationGraceStatus::Terminated),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for VerificationRequirement {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.min_tier != 0 {
            len += 1;
        }
        if !self.required_capabilities.is_empty() {
            len += 1;
        }
        if !self.required_auditors.is_empty() {
            len += 1;
        }
        if self.auditor_mode != 0 {
            len += 1;
        }
        if self.min_auditor_count != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.VerificationRequirement", len)?;
        if self.min_tier != 0 {
            let v = VerificationTier::try_from(self.min_tier)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.min_tier)))?;
            struct_ser.serialize_field("minTier", &v)?;
        }
        if !self.required_capabilities.is_empty() {
            let v = self.required_capabilities.iter().cloned().map(|v| {
                CapabilityFlag::try_from(v)
                    .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", v)))
                }).collect::<std::result::Result<Vec<_>, _>>()?;
            struct_ser.serialize_field("requiredCapabilities", &v)?;
        }
        if !self.required_auditors.is_empty() {
            struct_ser.serialize_field("requiredAuditors", &self.required_auditors)?;
        }
        if self.auditor_mode != 0 {
            let v = AuditorSelectionMode::try_from(self.auditor_mode)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.auditor_mode)))?;
            struct_ser.serialize_field("auditorMode", &v)?;
        }
        if self.min_auditor_count != 0 {
            struct_ser.serialize_field("minAuditorCount", &self.min_auditor_count)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for VerificationRequirement {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "min_tier",
            "minTier",
            "required_capabilities",
            "requiredCapabilities",
            "required_auditors",
            "requiredAuditors",
            "auditor_mode",
            "auditorMode",
            "min_auditor_count",
            "minAuditorCount",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            MinTier,
            RequiredCapabilities,
            RequiredAuditors,
            AuditorMode,
            MinAuditorCount,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "minTier" | "min_tier" => Ok(GeneratedField::MinTier),
                            "requiredCapabilities" | "required_capabilities" => Ok(GeneratedField::RequiredCapabilities),
                            "requiredAuditors" | "required_auditors" => Ok(GeneratedField::RequiredAuditors),
                            "auditorMode" | "auditor_mode" => Ok(GeneratedField::AuditorMode),
                            "minAuditorCount" | "min_auditor_count" => Ok(GeneratedField::MinAuditorCount),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = VerificationRequirement;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.VerificationRequirement")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<VerificationRequirement, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut min_tier__ = None;
                let mut required_capabilities__ = None;
                let mut required_auditors__ = None;
                let mut auditor_mode__ = None;
                let mut min_auditor_count__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::MinTier => {
                            if min_tier__.is_some() {
                                return Err(serde::de::Error::duplicate_field("minTier"));
                            }
                            min_tier__ = Some(map_.next_value::<VerificationTier>()? as i32);
                        }
                        GeneratedField::RequiredCapabilities => {
                            if required_capabilities__.is_some() {
                                return Err(serde::de::Error::duplicate_field("requiredCapabilities"));
                            }
                            required_capabilities__ = Some(map_.next_value::<Vec<CapabilityFlag>>()?.into_iter().map(|x| x as i32).collect());
                        }
                        GeneratedField::RequiredAuditors => {
                            if required_auditors__.is_some() {
                                return Err(serde::de::Error::duplicate_field("requiredAuditors"));
                            }
                            required_auditors__ = Some(map_.next_value()?);
                        }
                        GeneratedField::AuditorMode => {
                            if auditor_mode__.is_some() {
                                return Err(serde::de::Error::duplicate_field("auditorMode"));
                            }
                            auditor_mode__ = Some(map_.next_value::<AuditorSelectionMode>()? as i32);
                        }
                        GeneratedField::MinAuditorCount => {
                            if min_auditor_count__.is_some() {
                                return Err(serde::de::Error::duplicate_field("minAuditorCount"));
                            }
                            min_auditor_count__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(VerificationRequirement {
                    min_tier: min_tier__.unwrap_or_default(),
                    required_capabilities: required_capabilities__.unwrap_or_default(),
                    required_auditors: required_auditors__.unwrap_or_default(),
                    auditor_mode: auditor_mode__.unwrap_or_default(),
                    min_auditor_count: min_auditor_count__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.VerificationRequirement", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for VerificationStoreRecord {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.type_url.is_empty() {
            len += 1;
        }
        if !self.value.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.verification.v1.VerificationStoreRecord", len)?;
        if !self.type_url.is_empty() {
            struct_ser.serialize_field("typeUrl", &self.type_url)?;
        }
        if !self.value.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("value", pbjson::private::base64::encode(&self.value).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for VerificationStoreRecord {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "type_url",
            "typeUrl",
            "value",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            TypeUrl,
            Value,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "typeUrl" | "type_url" => Ok(GeneratedField::TypeUrl),
                            "value" => Ok(GeneratedField::Value),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = VerificationStoreRecord;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.verification.v1.VerificationStoreRecord")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<VerificationStoreRecord, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut type_url__ = None;
                let mut value__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::TypeUrl => {
                            if type_url__.is_some() {
                                return Err(serde::de::Error::duplicate_field("typeUrl"));
                            }
                            type_url__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Value => {
                            if value__.is_some() {
                                return Err(serde::de::Error::duplicate_field("value"));
                            }
                            value__ = 
                                Some(map_.next_value::<::pbjson::private::BytesDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(VerificationStoreRecord {
                    type_url: type_url__.unwrap_or_default(),
                    value: value__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.verification.v1.VerificationStoreRecord", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for VerificationTier {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "verification_tier_unspecified",
            Self::Identified => "verification_tier_identified",
            Self::Verified => "verification_tier_verified",
            Self::Established => "verification_tier_established",
            Self::Trusted => "verification_tier_trusted",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for VerificationTier {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "verification_tier_unspecified",
            "verification_tier_identified",
            "verification_tier_verified",
            "verification_tier_established",
            "verification_tier_trusted",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = VerificationTier;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                write!(formatter, "expected one of: {:?}", &FIELDS)
            }

            fn visit_i64<E>(self, v: i64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Signed(v), &self)
                    })
            }

            fn visit_u64<E>(self, v: u64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Unsigned(v), &self)
                    })
            }

            fn visit_str<E>(self, value: &str) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                match value {
                    "verification_tier_unspecified" => Ok(VerificationTier::Unspecified),
                    "verification_tier_identified" => Ok(VerificationTier::Identified),
                    "verification_tier_verified" => Ok(VerificationTier::Verified),
                    "verification_tier_established" => Ok(VerificationTier::Established),
                    "verification_tier_trusted" => Ok(VerificationTier::Trusted),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for VoidedReason {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "voided_reason_unspecified",
            Self::Discrepancy => "voided_reason_discrepancy",
            Self::Governance => "voided_reason_governance",
            Self::BondWithdrawn => "voided_reason_bond_withdrawn",
            Self::BondSlashed => "voided_reason_bond_slashed",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for VoidedReason {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "voided_reason_unspecified",
            "voided_reason_discrepancy",
            "voided_reason_governance",
            "voided_reason_bond_withdrawn",
            "voided_reason_bond_slashed",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = VoidedReason;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                write!(formatter, "expected one of: {:?}", &FIELDS)
            }

            fn visit_i64<E>(self, v: i64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Signed(v), &self)
                    })
            }

            fn visit_u64<E>(self, v: u64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Unsigned(v), &self)
                    })
            }

            fn visit_str<E>(self, value: &str) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                match value {
                    "voided_reason_unspecified" => Ok(VoidedReason::Unspecified),
                    "voided_reason_discrepancy" => Ok(VoidedReason::Discrepancy),
                    "voided_reason_governance" => Ok(VoidedReason::Governance),
                    "voided_reason_bond_withdrawn" => Ok(VoidedReason::BondWithdrawn),
                    "voided_reason_bond_slashed" => Ok(VoidedReason::BondSlashed),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
