// @generated
impl serde::Serialize for AccountUsage {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.total_sent.is_some() {
            len += 1;
        }
        if self.epoch_sent.is_some() {
            len += 1;
        }
        if self.last_epoch != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.AccountUsage", len)?;
        if let Some(v) = self.total_sent.as_ref() {
            struct_ser.serialize_field("totalSent", v)?;
        }
        if let Some(v) = self.epoch_sent.as_ref() {
            struct_ser.serialize_field("epochSent", v)?;
        }
        if self.last_epoch != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser
                .serialize_field("lastEpoch", ToString::to_string(&self.last_epoch).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for AccountUsage {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "total_sent",
            "totalSent",
            "epoch_sent",
            "epochSent",
            "last_epoch",
            "lastEpoch",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            TotalSent,
            EpochSent,
            LastEpoch,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "totalSent" | "total_sent" => Ok(GeneratedField::TotalSent),
                            "epochSent" | "epoch_sent" => Ok(GeneratedField::EpochSent),
                            "lastEpoch" | "last_epoch" => Ok(GeneratedField::LastEpoch),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = AccountUsage;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.AccountUsage")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<AccountUsage, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut total_sent__ = None;
                let mut epoch_sent__ = None;
                let mut last_epoch__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::TotalSent => {
                            if total_sent__.is_some() {
                                return Err(serde::de::Error::duplicate_field("totalSent"));
                            }
                            total_sent__ = map_.next_value()?;
                        }
                        GeneratedField::EpochSent => {
                            if epoch_sent__.is_some() {
                                return Err(serde::de::Error::duplicate_field("epochSent"));
                            }
                            epoch_sent__ = map_.next_value()?;
                        }
                        GeneratedField::LastEpoch => {
                            if last_epoch__.is_some() {
                                return Err(serde::de::Error::duplicate_field("lastEpoch"));
                            }
                            last_epoch__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                    }
                }
                Ok(AccountUsage {
                    total_sent: total_sent__,
                    epoch_sent: epoch_sent__,
                    last_epoch: last_epoch__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.AccountUsage", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for Decision {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "DECISION_UNSPECIFIED",
            Self::Restart => "DECISION_RESTART",
            Self::ConfirmMint => "DECISION_CONFIRM_MINT",
            Self::ConfirmVoid => "DECISION_CONFIRM_VOID",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for Decision {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "DECISION_UNSPECIFIED",
            "DECISION_RESTART",
            "DECISION_CONFIRM_MINT",
            "DECISION_CONFIRM_VOID",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = Decision;

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
                    "DECISION_UNSPECIFIED" => Ok(Decision::Unspecified),
                    "DECISION_RESTART" => Ok(Decision::Restart),
                    "DECISION_CONFIRM_MINT" => Ok(Decision::ConfirmMint),
                    "DECISION_CONFIRM_VOID" => Ok(Decision::ConfirmVoid),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for DestType {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Invalid => "DEST_TYPE_INVALID",
            Self::Solana => "DEST_TYPE_SOLANA",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for DestType {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["DEST_TYPE_INVALID", "DEST_TYPE_SOLANA"];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = DestType;

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
                    "DEST_TYPE_INVALID" => Ok(DestType::Invalid),
                    "DEST_TYPE_SOLANA" => Ok(DestType::Solana),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for EpochUsage {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.epoch != 0 {
            len += 1;
        }
        if self.amount.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.EpochUsage", len)?;
        if self.epoch != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("epoch", ToString::to_string(&self.epoch).as_str())?;
        }
        if let Some(v) = self.amount.as_ref() {
            struct_ser.serialize_field("amount", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EpochUsage {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["epoch", "amount"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Epoch,
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

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "epoch" => Ok(GeneratedField::Epoch),
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
            type Value = EpochUsage;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.EpochUsage")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EpochUsage, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut epoch__ = None;
                let mut amount__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Epoch => {
                            if epoch__.is_some() {
                                return Err(serde::de::Error::duplicate_field("epoch"));
                            }
                            epoch__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                        GeneratedField::Amount => {
                            if amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("amount"));
                            }
                            amount__ = map_.next_value()?;
                        }
                    }
                }
                Ok(EpochUsage {
                    epoch: epoch__.unwrap_or_default(),
                    amount: amount__,
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.EpochUsage", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventGuardianPaused {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.guardian.is_empty() {
            len += 1;
        }
        if self.deposits {
            len += 1;
        }
        if self.refunds {
            len += 1;
        }
        if self.enabled {
            len += 1;
        }
        if self.refunds_enabled {
            len += 1;
        }
        let mut struct_ser =
            serializer.serialize_struct("akash.sink.v1.EventGuardianPaused", len)?;
        if !self.guardian.is_empty() {
            struct_ser.serialize_field("guardian", &self.guardian)?;
        }
        if self.deposits {
            struct_ser.serialize_field("deposits", &self.deposits)?;
        }
        if self.refunds {
            struct_ser.serialize_field("refunds", &self.refunds)?;
        }
        if self.enabled {
            struct_ser.serialize_field("enabled", &self.enabled)?;
        }
        if self.refunds_enabled {
            struct_ser.serialize_field("refundsEnabled", &self.refunds_enabled)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventGuardianPaused {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "guardian",
            "deposits",
            "refunds",
            "enabled",
            "refunds_enabled",
            "refundsEnabled",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Guardian,
            Deposits,
            Refunds,
            Enabled,
            RefundsEnabled,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "guardian" => Ok(GeneratedField::Guardian),
                            "deposits" => Ok(GeneratedField::Deposits),
                            "refunds" => Ok(GeneratedField::Refunds),
                            "enabled" => Ok(GeneratedField::Enabled),
                            "refundsEnabled" | "refunds_enabled" => {
                                Ok(GeneratedField::RefundsEnabled)
                            }
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventGuardianPaused;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.EventGuardianPaused")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventGuardianPaused, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut guardian__ = None;
                let mut deposits__ = None;
                let mut refunds__ = None;
                let mut enabled__ = None;
                let mut refunds_enabled__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Guardian => {
                            if guardian__.is_some() {
                                return Err(serde::de::Error::duplicate_field("guardian"));
                            }
                            guardian__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Deposits => {
                            if deposits__.is_some() {
                                return Err(serde::de::Error::duplicate_field("deposits"));
                            }
                            deposits__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Refunds => {
                            if refunds__.is_some() {
                                return Err(serde::de::Error::duplicate_field("refunds"));
                            }
                            refunds__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Enabled => {
                            if enabled__.is_some() {
                                return Err(serde::de::Error::duplicate_field("enabled"));
                            }
                            enabled__ = Some(map_.next_value()?);
                        }
                        GeneratedField::RefundsEnabled => {
                            if refunds_enabled__.is_some() {
                                return Err(serde::de::Error::duplicate_field("refundsEnabled"));
                            }
                            refunds_enabled__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(EventGuardianPaused {
                    guardian: guardian__.unwrap_or_default(),
                    deposits: deposits__.unwrap_or_default(),
                    refunds: refunds__.unwrap_or_default(),
                    enabled: enabled__.unwrap_or_default(),
                    refunds_enabled: refunds_enabled__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.EventGuardianPaused",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for EventMintConfirmed {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.id.is_some() {
            len += 1;
        }
        let mut struct_ser =
            serializer.serialize_struct("akash.sink.v1.EventMintConfirmed", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventMintConfirmed {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["id"];

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

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
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
            type Value = EventMintConfirmed;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.EventMintConfirmed")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventMintConfirmed, V::Error>
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
                            id__ = map_.next_value()?;
                        }
                    }
                }
                Ok(EventMintConfirmed { id: id__ })
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.EventMintConfirmed",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for EventOutcomeReported {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.id.is_some() {
            len += 1;
        }
        if !self.reporter.is_empty() {
            len += 1;
        }
        if self.outcome != 0 {
            len += 1;
        }
        let mut struct_ser =
            serializer.serialize_struct("akash.sink.v1.EventOutcomeReported", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if !self.reporter.is_empty() {
            struct_ser.serialize_field("reporter", &self.reporter)?;
        }
        if self.outcome != 0 {
            let v = Outcome::try_from(self.outcome).map_err(|_| {
                serde::ser::Error::custom(format!("Invalid variant {}", self.outcome))
            })?;
            struct_ser.serialize_field("outcome", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventOutcomeReported {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["id", "reporter", "outcome"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Id,
            Reporter,
            Outcome,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "id" => Ok(GeneratedField::Id),
                            "reporter" => Ok(GeneratedField::Reporter),
                            "outcome" => Ok(GeneratedField::Outcome),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventOutcomeReported;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.EventOutcomeReported")
            }

            fn visit_map<V>(
                self,
                mut map_: V,
            ) -> std::result::Result<EventOutcomeReported, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                let mut reporter__ = None;
                let mut outcome__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = map_.next_value()?;
                        }
                        GeneratedField::Reporter => {
                            if reporter__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reporter"));
                            }
                            reporter__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Outcome => {
                            if outcome__.is_some() {
                                return Err(serde::de::Error::duplicate_field("outcome"));
                            }
                            outcome__ = Some(map_.next_value::<Outcome>()? as i32);
                        }
                    }
                }
                Ok(EventOutcomeReported {
                    id: id__,
                    reporter: reporter__.unwrap_or_default(),
                    outcome: outcome__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.EventOutcomeReported",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for EventPreexistingBalanceMoved {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.amount.is_empty() {
            len += 1;
        }
        if !self.recipient.is_empty() {
            len += 1;
        }
        let mut struct_ser =
            serializer.serialize_struct("akash.sink.v1.EventPreexistingBalanceMoved", len)?;
        if !self.amount.is_empty() {
            struct_ser.serialize_field("amount", &self.amount)?;
        }
        if !self.recipient.is_empty() {
            struct_ser.serialize_field("recipient", &self.recipient)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventPreexistingBalanceMoved {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["amount", "recipient"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Amount,
            Recipient,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "amount" => Ok(GeneratedField::Amount),
                            "recipient" => Ok(GeneratedField::Recipient),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventPreexistingBalanceMoved;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.EventPreexistingBalanceMoved")
            }

            fn visit_map<V>(
                self,
                mut map_: V,
            ) -> std::result::Result<EventPreexistingBalanceMoved, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut amount__ = None;
                let mut recipient__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Amount => {
                            if amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("amount"));
                            }
                            amount__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Recipient => {
                            if recipient__.is_some() {
                                return Err(serde::de::Error::duplicate_field("recipient"));
                            }
                            recipient__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(EventPreexistingBalanceMoved {
                    amount: amount__.unwrap_or_default(),
                    recipient: recipient__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.EventPreexistingBalanceMoved",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for EventReportConflict {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.id.is_some() {
            len += 1;
        }
        if !self.reporter.is_empty() {
            len += 1;
        }
        let mut struct_ser =
            serializer.serialize_struct("akash.sink.v1.EventReportConflict", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if !self.reporter.is_empty() {
            struct_ser.serialize_field("reporter", &self.reporter)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventReportConflict {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["id", "reporter"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Id,
            Reporter,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "id" => Ok(GeneratedField::Id),
                            "reporter" => Ok(GeneratedField::Reporter),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventReportConflict;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.EventReportConflict")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventReportConflict, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                let mut reporter__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = map_.next_value()?;
                        }
                        GeneratedField::Reporter => {
                            if reporter__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reporter"));
                            }
                            reporter__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(EventReportConflict {
                    id: id__,
                    reporter: reporter__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.EventReportConflict",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for EventTotalsUpdated {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.pending_amount.is_some() {
            len += 1;
        }
        if self.pending_count != 0 {
            len += 1;
        }
        if self.minted_amount.is_some() {
            len += 1;
        }
        if self.minted_count != 0 {
            len += 1;
        }
        if self.voided_amount.is_some() {
            len += 1;
        }
        if self.voided_count != 0 {
            len += 1;
        }
        if self.refund_pending_amount.is_some() {
            len += 1;
        }
        if self.refund_pending_count != 0 {
            len += 1;
        }
        if self.refunded_amount.is_some() {
            len += 1;
        }
        if self.refunded_count != 0 {
            len += 1;
        }
        if self.refunded_in_epoch.is_some() {
            len += 1;
        }
        let mut struct_ser =
            serializer.serialize_struct("akash.sink.v1.EventTotalsUpdated", len)?;
        if let Some(v) = self.pending_amount.as_ref() {
            struct_ser.serialize_field("pendingAmount", v)?;
        }
        if self.pending_count != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field(
                "pendingCount",
                ToString::to_string(&self.pending_count).as_str(),
            )?;
        }
        if let Some(v) = self.minted_amount.as_ref() {
            struct_ser.serialize_field("mintedAmount", v)?;
        }
        if self.minted_count != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field(
                "mintedCount",
                ToString::to_string(&self.minted_count).as_str(),
            )?;
        }
        if let Some(v) = self.voided_amount.as_ref() {
            struct_ser.serialize_field("voidedAmount", v)?;
        }
        if self.voided_count != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field(
                "voidedCount",
                ToString::to_string(&self.voided_count).as_str(),
            )?;
        }
        if let Some(v) = self.refund_pending_amount.as_ref() {
            struct_ser.serialize_field("refundPendingAmount", v)?;
        }
        if self.refund_pending_count != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field(
                "refundPendingCount",
                ToString::to_string(&self.refund_pending_count).as_str(),
            )?;
        }
        if let Some(v) = self.refunded_amount.as_ref() {
            struct_ser.serialize_field("refundedAmount", v)?;
        }
        if self.refunded_count != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field(
                "refundedCount",
                ToString::to_string(&self.refunded_count).as_str(),
            )?;
        }
        if let Some(v) = self.refunded_in_epoch.as_ref() {
            struct_ser.serialize_field("refundedInEpoch", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventTotalsUpdated {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "pending_amount",
            "pendingAmount",
            "pending_count",
            "pendingCount",
            "minted_amount",
            "mintedAmount",
            "minted_count",
            "mintedCount",
            "voided_amount",
            "voidedAmount",
            "voided_count",
            "voidedCount",
            "refund_pending_amount",
            "refundPendingAmount",
            "refund_pending_count",
            "refundPendingCount",
            "refunded_amount",
            "refundedAmount",
            "refunded_count",
            "refundedCount",
            "refunded_in_epoch",
            "refundedInEpoch",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            PendingAmount,
            PendingCount,
            MintedAmount,
            MintedCount,
            VoidedAmount,
            VoidedCount,
            RefundPendingAmount,
            RefundPendingCount,
            RefundedAmount,
            RefundedCount,
            RefundedInEpoch,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "pendingAmount" | "pending_amount" => Ok(GeneratedField::PendingAmount),
                            "pendingCount" | "pending_count" => Ok(GeneratedField::PendingCount),
                            "mintedAmount" | "minted_amount" => Ok(GeneratedField::MintedAmount),
                            "mintedCount" | "minted_count" => Ok(GeneratedField::MintedCount),
                            "voidedAmount" | "voided_amount" => Ok(GeneratedField::VoidedAmount),
                            "voidedCount" | "voided_count" => Ok(GeneratedField::VoidedCount),
                            "refundPendingAmount" | "refund_pending_amount" => {
                                Ok(GeneratedField::RefundPendingAmount)
                            }
                            "refundPendingCount" | "refund_pending_count" => {
                                Ok(GeneratedField::RefundPendingCount)
                            }
                            "refundedAmount" | "refunded_amount" => {
                                Ok(GeneratedField::RefundedAmount)
                            }
                            "refundedCount" | "refunded_count" => Ok(GeneratedField::RefundedCount),
                            "refundedInEpoch" | "refunded_in_epoch" => {
                                Ok(GeneratedField::RefundedInEpoch)
                            }
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventTotalsUpdated;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.EventTotalsUpdated")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventTotalsUpdated, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut pending_amount__ = None;
                let mut pending_count__ = None;
                let mut minted_amount__ = None;
                let mut minted_count__ = None;
                let mut voided_amount__ = None;
                let mut voided_count__ = None;
                let mut refund_pending_amount__ = None;
                let mut refund_pending_count__ = None;
                let mut refunded_amount__ = None;
                let mut refunded_count__ = None;
                let mut refunded_in_epoch__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::PendingAmount => {
                            if pending_amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pendingAmount"));
                            }
                            pending_amount__ = map_.next_value()?;
                        }
                        GeneratedField::PendingCount => {
                            if pending_count__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pendingCount"));
                            }
                            pending_count__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                        GeneratedField::MintedAmount => {
                            if minted_amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("mintedAmount"));
                            }
                            minted_amount__ = map_.next_value()?;
                        }
                        GeneratedField::MintedCount => {
                            if minted_count__.is_some() {
                                return Err(serde::de::Error::duplicate_field("mintedCount"));
                            }
                            minted_count__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                        GeneratedField::VoidedAmount => {
                            if voided_amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("voidedAmount"));
                            }
                            voided_amount__ = map_.next_value()?;
                        }
                        GeneratedField::VoidedCount => {
                            if voided_count__.is_some() {
                                return Err(serde::de::Error::duplicate_field("voidedCount"));
                            }
                            voided_count__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                        GeneratedField::RefundPendingAmount => {
                            if refund_pending_amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field(
                                    "refundPendingAmount",
                                ));
                            }
                            refund_pending_amount__ = map_.next_value()?;
                        }
                        GeneratedField::RefundPendingCount => {
                            if refund_pending_count__.is_some() {
                                return Err(serde::de::Error::duplicate_field(
                                    "refundPendingCount",
                                ));
                            }
                            refund_pending_count__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                        GeneratedField::RefundedAmount => {
                            if refunded_amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("refundedAmount"));
                            }
                            refunded_amount__ = map_.next_value()?;
                        }
                        GeneratedField::RefundedCount => {
                            if refunded_count__.is_some() {
                                return Err(serde::de::Error::duplicate_field("refundedCount"));
                            }
                            refunded_count__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                        GeneratedField::RefundedInEpoch => {
                            if refunded_in_epoch__.is_some() {
                                return Err(serde::de::Error::duplicate_field("refundedInEpoch"));
                            }
                            refunded_in_epoch__ = map_.next_value()?;
                        }
                    }
                }
                Ok(EventTotalsUpdated {
                    pending_amount: pending_amount__,
                    pending_count: pending_count__.unwrap_or_default(),
                    minted_amount: minted_amount__,
                    minted_count: minted_count__.unwrap_or_default(),
                    voided_amount: voided_amount__,
                    voided_count: voided_count__.unwrap_or_default(),
                    refund_pending_amount: refund_pending_amount__,
                    refund_pending_count: refund_pending_count__.unwrap_or_default(),
                    refunded_amount: refunded_amount__,
                    refunded_count: refunded_count__.unwrap_or_default(),
                    refunded_in_epoch: refunded_in_epoch__,
                })
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.EventTotalsUpdated",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for EventTransferCreated {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.id.is_some() {
            len += 1;
        }
        if !self.sender.is_empty() {
            len += 1;
        }
        if !self.recipient.is_empty() {
            len += 1;
        }
        if self.amount.is_some() {
            len += 1;
        }
        if self.height != 0 {
            len += 1;
        }
        let mut struct_ser =
            serializer.serialize_struct("akash.sink.v1.EventTransferCreated", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if !self.sender.is_empty() {
            struct_ser.serialize_field("sender", &self.sender)?;
        }
        if !self.recipient.is_empty() {
            struct_ser.serialize_field("recipient", &self.recipient)?;
        }
        if let Some(v) = self.amount.as_ref() {
            struct_ser.serialize_field("amount", v)?;
        }
        if self.height != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("height", ToString::to_string(&self.height).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventTransferCreated {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["id", "sender", "recipient", "amount", "height"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Id,
            Sender,
            Recipient,
            Amount,
            Height,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "id" => Ok(GeneratedField::Id),
                            "sender" => Ok(GeneratedField::Sender),
                            "recipient" => Ok(GeneratedField::Recipient),
                            "amount" => Ok(GeneratedField::Amount),
                            "height" => Ok(GeneratedField::Height),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventTransferCreated;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.EventTransferCreated")
            }

            fn visit_map<V>(
                self,
                mut map_: V,
            ) -> std::result::Result<EventTransferCreated, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                let mut sender__ = None;
                let mut recipient__ = None;
                let mut amount__ = None;
                let mut height__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = map_.next_value()?;
                        }
                        GeneratedField::Sender => {
                            if sender__.is_some() {
                                return Err(serde::de::Error::duplicate_field("sender"));
                            }
                            sender__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Recipient => {
                            if recipient__.is_some() {
                                return Err(serde::de::Error::duplicate_field("recipient"));
                            }
                            recipient__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Amount => {
                            if amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("amount"));
                            }
                            amount__ = map_.next_value()?;
                        }
                        GeneratedField::Height => {
                            if height__.is_some() {
                                return Err(serde::de::Error::duplicate_field("height"));
                            }
                            height__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                    }
                }
                Ok(EventTransferCreated {
                    id: id__,
                    sender: sender__.unwrap_or_default(),
                    recipient: recipient__.unwrap_or_default(),
                    amount: amount__,
                    height: height__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.EventTransferCreated",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for EventTransferRefunded {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.id.is_some() {
            len += 1;
        }
        if !self.sender.is_empty() {
            len += 1;
        }
        if self.amount.is_some() {
            len += 1;
        }
        let mut struct_ser =
            serializer.serialize_struct("akash.sink.v1.EventTransferRefunded", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if !self.sender.is_empty() {
            struct_ser.serialize_field("sender", &self.sender)?;
        }
        if let Some(v) = self.amount.as_ref() {
            struct_ser.serialize_field("amount", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventTransferRefunded {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["id", "sender", "amount"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Id,
            Sender,
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

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "id" => Ok(GeneratedField::Id),
                            "sender" => Ok(GeneratedField::Sender),
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
            type Value = EventTransferRefunded;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.EventTransferRefunded")
            }

            fn visit_map<V>(
                self,
                mut map_: V,
            ) -> std::result::Result<EventTransferRefunded, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                let mut sender__ = None;
                let mut amount__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = map_.next_value()?;
                        }
                        GeneratedField::Sender => {
                            if sender__.is_some() {
                                return Err(serde::de::Error::duplicate_field("sender"));
                            }
                            sender__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Amount => {
                            if amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("amount"));
                            }
                            amount__ = map_.next_value()?;
                        }
                    }
                }
                Ok(EventTransferRefunded {
                    id: id__,
                    sender: sender__.unwrap_or_default(),
                    amount: amount__,
                })
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.EventTransferRefunded",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for EventTransferResolved {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.id.is_some() {
            len += 1;
        }
        if self.decision != 0 {
            len += 1;
        }
        if !self.reason.is_empty() {
            len += 1;
        }
        let mut struct_ser =
            serializer.serialize_struct("akash.sink.v1.EventTransferResolved", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if self.decision != 0 {
            let v = Decision::try_from(self.decision).map_err(|_| {
                serde::ser::Error::custom(format!("Invalid variant {}", self.decision))
            })?;
            struct_ser.serialize_field("decision", &v)?;
        }
        if !self.reason.is_empty() {
            struct_ser.serialize_field("reason", &self.reason)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventTransferResolved {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["id", "decision", "reason"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Id,
            Decision,
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

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "id" => Ok(GeneratedField::Id),
                            "decision" => Ok(GeneratedField::Decision),
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
            type Value = EventTransferResolved;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.EventTransferResolved")
            }

            fn visit_map<V>(
                self,
                mut map_: V,
            ) -> std::result::Result<EventTransferResolved, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                let mut decision__ = None;
                let mut reason__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = map_.next_value()?;
                        }
                        GeneratedField::Decision => {
                            if decision__.is_some() {
                                return Err(serde::de::Error::duplicate_field("decision"));
                            }
                            decision__ = Some(map_.next_value::<Decision>()? as i32);
                        }
                        GeneratedField::Reason => {
                            if reason__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reason"));
                            }
                            reason__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(EventTransferResolved {
                    id: id__,
                    decision: decision__.unwrap_or_default(),
                    reason: reason__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.EventTransferResolved",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for EventVoidConfirmed {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.id.is_some() {
            len += 1;
        }
        if self.challenge_end.is_some() {
            len += 1;
        }
        let mut struct_ser =
            serializer.serialize_struct("akash.sink.v1.EventVoidConfirmed", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if let Some(v) = self.challenge_end.as_ref() {
            struct_ser.serialize_field("challengeEnd", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventVoidConfirmed {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["id", "challenge_end", "challengeEnd"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Id,
            ChallengeEnd,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "id" => Ok(GeneratedField::Id),
                            "challengeEnd" | "challenge_end" => Ok(GeneratedField::ChallengeEnd),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventVoidConfirmed;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.EventVoidConfirmed")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventVoidConfirmed, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                let mut challenge_end__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = map_.next_value()?;
                        }
                        GeneratedField::ChallengeEnd => {
                            if challenge_end__.is_some() {
                                return Err(serde::de::Error::duplicate_field("challengeEnd"));
                            }
                            challenge_end__ = map_.next_value()?;
                        }
                    }
                }
                Ok(EventVoidConfirmed {
                    id: id__,
                    challenge_end: challenge_end__,
                })
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.EventVoidConfirmed",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for EventVoidOrdered {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.id.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.EventVoidOrdered", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventVoidOrdered {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["id"];

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

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
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
            type Value = EventVoidOrdered;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.EventVoidOrdered")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventVoidOrdered, V::Error>
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
                            id__ = map_.next_value()?;
                        }
                    }
                }
                Ok(EventVoidOrdered { id: id__ })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.EventVoidOrdered", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventVoidVetoed {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.id.is_some() {
            len += 1;
        }
        if !self.signer.is_empty() {
            len += 1;
        }
        if !self.reason.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.EventVoidVetoed", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if !self.signer.is_empty() {
            struct_ser.serialize_field("signer", &self.signer)?;
        }
        if !self.reason.is_empty() {
            struct_ser.serialize_field("reason", &self.reason)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventVoidVetoed {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["id", "signer", "reason"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Id,
            Signer,
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

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "id" => Ok(GeneratedField::Id),
                            "signer" => Ok(GeneratedField::Signer),
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
            type Value = EventVoidVetoed;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.EventVoidVetoed")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventVoidVetoed, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                let mut signer__ = None;
                let mut reason__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = map_.next_value()?;
                        }
                        GeneratedField::Signer => {
                            if signer__.is_some() {
                                return Err(serde::de::Error::duplicate_field("signer"));
                            }
                            signer__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Reason => {
                            if reason__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reason"));
                            }
                            reason__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(EventVoidVetoed {
                    id: id__,
                    signer: signer__.unwrap_or_default(),
                    reason: reason__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.EventVoidVetoed", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for GenesisAccount {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.owner.is_empty() {
            len += 1;
        }
        if self.usage.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.GenesisAccount", len)?;
        if !self.owner.is_empty() {
            struct_ser.serialize_field("owner", &self.owner)?;
        }
        if let Some(v) = self.usage.as_ref() {
            struct_ser.serialize_field("usage", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for GenesisAccount {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["owner", "usage"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Owner,
            Usage,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "owner" => Ok(GeneratedField::Owner),
                            "usage" => Ok(GeneratedField::Usage),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = GenesisAccount;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.GenesisAccount")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<GenesisAccount, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut owner__ = None;
                let mut usage__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Owner => {
                            if owner__.is_some() {
                                return Err(serde::de::Error::duplicate_field("owner"));
                            }
                            owner__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Usage => {
                            if usage__.is_some() {
                                return Err(serde::de::Error::duplicate_field("usage"));
                            }
                            usage__ = map_.next_value()?;
                        }
                    }
                }
                Ok(GenesisAccount {
                    owner: owner__.unwrap_or_default(),
                    usage: usage__,
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.GenesisAccount", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for GenesisRecord {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.id.is_some() {
            len += 1;
        }
        if self.data.is_some() {
            len += 1;
        }
        if self.settlement.is_some() {
            len += 1;
        }
        if self.status != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.GenesisRecord", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if let Some(v) = self.data.as_ref() {
            struct_ser.serialize_field("data", v)?;
        }
        if let Some(v) = self.settlement.as_ref() {
            struct_ser.serialize_field("settlement", v)?;
        }
        if self.status != 0 {
            let v = Status::try_from(self.status).map_err(|_| {
                serde::ser::Error::custom(format!("Invalid variant {}", self.status))
            })?;
            struct_ser.serialize_field("status", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for GenesisRecord {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["id", "data", "settlement", "status"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Id,
            Data,
            Settlement,
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

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "id" => Ok(GeneratedField::Id),
                            "data" => Ok(GeneratedField::Data),
                            "settlement" => Ok(GeneratedField::Settlement),
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
            type Value = GenesisRecord;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.GenesisRecord")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<GenesisRecord, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                let mut data__ = None;
                let mut settlement__ = None;
                let mut status__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = map_.next_value()?;
                        }
                        GeneratedField::Data => {
                            if data__.is_some() {
                                return Err(serde::de::Error::duplicate_field("data"));
                            }
                            data__ = map_.next_value()?;
                        }
                        GeneratedField::Settlement => {
                            if settlement__.is_some() {
                                return Err(serde::de::Error::duplicate_field("settlement"));
                            }
                            settlement__ = map_.next_value()?;
                        }
                        GeneratedField::Status => {
                            if status__.is_some() {
                                return Err(serde::de::Error::duplicate_field("status"));
                            }
                            status__ = Some(map_.next_value::<Status>()? as i32);
                        }
                    }
                }
                Ok(GenesisRecord {
                    id: id__,
                    data: data__,
                    settlement: settlement__,
                    status: status__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.GenesisRecord", FIELDS, GeneratedVisitor)
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
        if self.totals.is_some() {
            len += 1;
        }
        if !self.policies.is_empty() {
            len += 1;
        }
        if self.send_usage.is_some() {
            len += 1;
        }
        if self.refund_usage.is_some() {
            len += 1;
        }
        if !self.records.is_empty() {
            len += 1;
        }
        if !self.accounts.is_empty() {
            len += 1;
        }
        if self.schema_version != 0 {
            len += 1;
        }
        if self.current_policy_version != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.GenesisState", len)?;
        if let Some(v) = self.params.as_ref() {
            struct_ser.serialize_field("params", v)?;
        }
        if let Some(v) = self.totals.as_ref() {
            struct_ser.serialize_field("totals", v)?;
        }
        if !self.policies.is_empty() {
            struct_ser.serialize_field("policies", &self.policies)?;
        }
        if let Some(v) = self.send_usage.as_ref() {
            struct_ser.serialize_field("sendUsage", v)?;
        }
        if let Some(v) = self.refund_usage.as_ref() {
            struct_ser.serialize_field("refundUsage", v)?;
        }
        if !self.records.is_empty() {
            struct_ser.serialize_field("records", &self.records)?;
        }
        if !self.accounts.is_empty() {
            struct_ser.serialize_field("accounts", &self.accounts)?;
        }
        if self.schema_version != 0 {
            struct_ser.serialize_field("schemaVersion", &self.schema_version)?;
        }
        if self.current_policy_version != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field(
                "currentPolicyVersion",
                ToString::to_string(&self.current_policy_version).as_str(),
            )?;
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
            "totals",
            "policies",
            "send_usage",
            "sendUsage",
            "refund_usage",
            "refundUsage",
            "records",
            "accounts",
            "schema_version",
            "schemaVersion",
            "current_policy_version",
            "currentPolicyVersion",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Params,
            Totals,
            Policies,
            SendUsage,
            RefundUsage,
            Records,
            Accounts,
            SchemaVersion,
            CurrentPolicyVersion,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "params" => Ok(GeneratedField::Params),
                            "totals" => Ok(GeneratedField::Totals),
                            "policies" => Ok(GeneratedField::Policies),
                            "sendUsage" | "send_usage" => Ok(GeneratedField::SendUsage),
                            "refundUsage" | "refund_usage" => Ok(GeneratedField::RefundUsage),
                            "records" => Ok(GeneratedField::Records),
                            "accounts" => Ok(GeneratedField::Accounts),
                            "schemaVersion" | "schema_version" => Ok(GeneratedField::SchemaVersion),
                            "currentPolicyVersion" | "current_policy_version" => {
                                Ok(GeneratedField::CurrentPolicyVersion)
                            }
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
                formatter.write_str("struct akash.sink.v1.GenesisState")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<GenesisState, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut params__ = None;
                let mut totals__ = None;
                let mut policies__ = None;
                let mut send_usage__ = None;
                let mut refund_usage__ = None;
                let mut records__ = None;
                let mut accounts__ = None;
                let mut schema_version__ = None;
                let mut current_policy_version__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Params => {
                            if params__.is_some() {
                                return Err(serde::de::Error::duplicate_field("params"));
                            }
                            params__ = map_.next_value()?;
                        }
                        GeneratedField::Totals => {
                            if totals__.is_some() {
                                return Err(serde::de::Error::duplicate_field("totals"));
                            }
                            totals__ = map_.next_value()?;
                        }
                        GeneratedField::Policies => {
                            if policies__.is_some() {
                                return Err(serde::de::Error::duplicate_field("policies"));
                            }
                            policies__ = Some(map_.next_value()?);
                        }
                        GeneratedField::SendUsage => {
                            if send_usage__.is_some() {
                                return Err(serde::de::Error::duplicate_field("sendUsage"));
                            }
                            send_usage__ = map_.next_value()?;
                        }
                        GeneratedField::RefundUsage => {
                            if refund_usage__.is_some() {
                                return Err(serde::de::Error::duplicate_field("refundUsage"));
                            }
                            refund_usage__ = map_.next_value()?;
                        }
                        GeneratedField::Records => {
                            if records__.is_some() {
                                return Err(serde::de::Error::duplicate_field("records"));
                            }
                            records__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Accounts => {
                            if accounts__.is_some() {
                                return Err(serde::de::Error::duplicate_field("accounts"));
                            }
                            accounts__ = Some(map_.next_value()?);
                        }
                        GeneratedField::SchemaVersion => {
                            if schema_version__.is_some() {
                                return Err(serde::de::Error::duplicate_field("schemaVersion"));
                            }
                            schema_version__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                        GeneratedField::CurrentPolicyVersion => {
                            if current_policy_version__.is_some() {
                                return Err(serde::de::Error::duplicate_field(
                                    "currentPolicyVersion",
                                ));
                            }
                            current_policy_version__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                    }
                }
                Ok(GenesisState {
                    params: params__,
                    totals: totals__,
                    policies: policies__.unwrap_or_default(),
                    send_usage: send_usage__,
                    refund_usage: refund_usage__,
                    records: records__.unwrap_or_default(),
                    accounts: accounts__.unwrap_or_default(),
                    schema_version: schema_version__.unwrap_or_default(),
                    current_policy_version: current_policy_version__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.GenesisState", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for Hold {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::None => "HOLD_NONE",
            Self::Veto => "HOLD_VETO",
            Self::Conflict => "HOLD_CONFLICT",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for Hold {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["HOLD_NONE", "HOLD_VETO", "HOLD_CONFLICT"];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = Hold;

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
                    "HOLD_NONE" => Ok(Hold::None),
                    "HOLD_VETO" => Ok(Hold::Veto),
                    "HOLD_CONFLICT" => Ok(Hold::Conflict),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for Id {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.owner.is_empty() {
            len += 1;
        }
        if self.height != 0 {
            len += 1;
        }
        if self.sequence != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.ID", len)?;
        if !self.owner.is_empty() {
            struct_ser.serialize_field("owner", &self.owner)?;
        }
        if self.height != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("height", ToString::to_string(&self.height).as_str())?;
        }
        if self.sequence != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("sequence", ToString::to_string(&self.sequence).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for Id {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["owner", "height", "sequence"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Owner,
            Height,
            Sequence,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "owner" => Ok(GeneratedField::Owner),
                            "height" => Ok(GeneratedField::Height),
                            "sequence" => Ok(GeneratedField::Sequence),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = Id;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.ID")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<Id, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut owner__ = None;
                let mut height__ = None;
                let mut sequence__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Owner => {
                            if owner__.is_some() {
                                return Err(serde::de::Error::duplicate_field("owner"));
                            }
                            owner__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Height => {
                            if height__.is_some() {
                                return Err(serde::de::Error::duplicate_field("height"));
                            }
                            height__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                        GeneratedField::Sequence => {
                            if sequence__.is_some() {
                                return Err(serde::de::Error::duplicate_field("sequence"));
                            }
                            sequence__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                    }
                }
                Ok(Id {
                    owner: owner__.unwrap_or_default(),
                    height: height__.unwrap_or_default(),
                    sequence: sequence__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.ID", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgConfirmMint {
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
        if self.id.is_some() {
            len += 1;
        }
        if self.evidence.is_some() {
            len += 1;
        }
        if !self.reason.is_empty() {
            len += 1;
        }
        if self.expected_round != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.MsgConfirmMint", len)?;
        if !self.authority.is_empty() {
            struct_ser.serialize_field("authority", &self.authority)?;
        }
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if let Some(v) = self.evidence.as_ref() {
            struct_ser.serialize_field("evidence", v)?;
        }
        if !self.reason.is_empty() {
            struct_ser.serialize_field("reason", &self.reason)?;
        }
        if self.expected_round != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field(
                "expectedRound",
                ToString::to_string(&self.expected_round).as_str(),
            )?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgConfirmMint {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "authority",
            "id",
            "evidence",
            "reason",
            "expected_round",
            "expectedRound",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Authority,
            Id,
            Evidence,
            Reason,
            ExpectedRound,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "authority" => Ok(GeneratedField::Authority),
                            "id" => Ok(GeneratedField::Id),
                            "evidence" => Ok(GeneratedField::Evidence),
                            "reason" => Ok(GeneratedField::Reason),
                            "expectedRound" | "expected_round" => Ok(GeneratedField::ExpectedRound),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgConfirmMint;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.MsgConfirmMint")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgConfirmMint, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut authority__ = None;
                let mut id__ = None;
                let mut evidence__ = None;
                let mut reason__ = None;
                let mut expected_round__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Authority => {
                            if authority__.is_some() {
                                return Err(serde::de::Error::duplicate_field("authority"));
                            }
                            authority__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = map_.next_value()?;
                        }
                        GeneratedField::Evidence => {
                            if evidence__.is_some() {
                                return Err(serde::de::Error::duplicate_field("evidence"));
                            }
                            evidence__ = map_.next_value()?;
                        }
                        GeneratedField::Reason => {
                            if reason__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reason"));
                            }
                            reason__ = Some(map_.next_value()?);
                        }
                        GeneratedField::ExpectedRound => {
                            if expected_round__.is_some() {
                                return Err(serde::de::Error::duplicate_field("expectedRound"));
                            }
                            expected_round__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                    }
                }
                Ok(MsgConfirmMint {
                    authority: authority__.unwrap_or_default(),
                    id: id__,
                    evidence: evidence__,
                    reason: reason__.unwrap_or_default(),
                    expected_round: expected_round__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.MsgConfirmMint", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgConfirmMintResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser =
            serializer.serialize_struct("akash.sink.v1.MsgConfirmMintResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgConfirmMintResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {}
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
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
            type Value = MsgConfirmMintResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.MsgConfirmMintResponse")
            }

            fn visit_map<V>(
                self,
                mut map_: V,
            ) -> std::result::Result<MsgConfirmMintResponse, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgConfirmMintResponse {})
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.MsgConfirmMintResponse",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for MsgConfirmVoid {
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
        if self.id.is_some() {
            len += 1;
        }
        if self.evidence.is_some() {
            len += 1;
        }
        if !self.reason.is_empty() {
            len += 1;
        }
        if self.expected_round != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.MsgConfirmVoid", len)?;
        if !self.authority.is_empty() {
            struct_ser.serialize_field("authority", &self.authority)?;
        }
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if let Some(v) = self.evidence.as_ref() {
            struct_ser.serialize_field("evidence", v)?;
        }
        if !self.reason.is_empty() {
            struct_ser.serialize_field("reason", &self.reason)?;
        }
        if self.expected_round != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field(
                "expectedRound",
                ToString::to_string(&self.expected_round).as_str(),
            )?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgConfirmVoid {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "authority",
            "id",
            "evidence",
            "reason",
            "expected_round",
            "expectedRound",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Authority,
            Id,
            Evidence,
            Reason,
            ExpectedRound,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "authority" => Ok(GeneratedField::Authority),
                            "id" => Ok(GeneratedField::Id),
                            "evidence" => Ok(GeneratedField::Evidence),
                            "reason" => Ok(GeneratedField::Reason),
                            "expectedRound" | "expected_round" => Ok(GeneratedField::ExpectedRound),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgConfirmVoid;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.MsgConfirmVoid")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgConfirmVoid, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut authority__ = None;
                let mut id__ = None;
                let mut evidence__ = None;
                let mut reason__ = None;
                let mut expected_round__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Authority => {
                            if authority__.is_some() {
                                return Err(serde::de::Error::duplicate_field("authority"));
                            }
                            authority__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = map_.next_value()?;
                        }
                        GeneratedField::Evidence => {
                            if evidence__.is_some() {
                                return Err(serde::de::Error::duplicate_field("evidence"));
                            }
                            evidence__ = map_.next_value()?;
                        }
                        GeneratedField::Reason => {
                            if reason__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reason"));
                            }
                            reason__ = Some(map_.next_value()?);
                        }
                        GeneratedField::ExpectedRound => {
                            if expected_round__.is_some() {
                                return Err(serde::de::Error::duplicate_field("expectedRound"));
                            }
                            expected_round__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                    }
                }
                Ok(MsgConfirmVoid {
                    authority: authority__.unwrap_or_default(),
                    id: id__,
                    evidence: evidence__,
                    reason: reason__.unwrap_or_default(),
                    expected_round: expected_round__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.MsgConfirmVoid", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgConfirmVoidResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser =
            serializer.serialize_struct("akash.sink.v1.MsgConfirmVoidResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgConfirmVoidResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {}
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
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
            type Value = MsgConfirmVoidResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.MsgConfirmVoidResponse")
            }

            fn visit_map<V>(
                self,
                mut map_: V,
            ) -> std::result::Result<MsgConfirmVoidResponse, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgConfirmVoidResponse {})
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.MsgConfirmVoidResponse",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for MsgFinalizeRefund {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.executor.is_empty() {
            len += 1;
        }
        if self.id.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.MsgFinalizeRefund", len)?;
        if !self.executor.is_empty() {
            struct_ser.serialize_field("executor", &self.executor)?;
        }
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgFinalizeRefund {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["executor", "id"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Executor,
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

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "executor" => Ok(GeneratedField::Executor),
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
            type Value = MsgFinalizeRefund;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.MsgFinalizeRefund")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgFinalizeRefund, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut executor__ = None;
                let mut id__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Executor => {
                            if executor__.is_some() {
                                return Err(serde::de::Error::duplicate_field("executor"));
                            }
                            executor__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = map_.next_value()?;
                        }
                    }
                }
                Ok(MsgFinalizeRefund {
                    executor: executor__.unwrap_or_default(),
                    id: id__,
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.MsgFinalizeRefund", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgFinalizeRefundResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser =
            serializer.serialize_struct("akash.sink.v1.MsgFinalizeRefundResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgFinalizeRefundResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {}
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
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
            type Value = MsgFinalizeRefundResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.MsgFinalizeRefundResponse")
            }

            fn visit_map<V>(
                self,
                mut map_: V,
            ) -> std::result::Result<MsgFinalizeRefundResponse, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgFinalizeRefundResponse {})
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.MsgFinalizeRefundResponse",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for MsgGuardianPause {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.guardian.is_empty() {
            len += 1;
        }
        if self.deposits {
            len += 1;
        }
        if self.refunds {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.MsgGuardianPause", len)?;
        if !self.guardian.is_empty() {
            struct_ser.serialize_field("guardian", &self.guardian)?;
        }
        if self.deposits {
            struct_ser.serialize_field("deposits", &self.deposits)?;
        }
        if self.refunds {
            struct_ser.serialize_field("refunds", &self.refunds)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgGuardianPause {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["guardian", "deposits", "refunds"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Guardian,
            Deposits,
            Refunds,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "guardian" => Ok(GeneratedField::Guardian),
                            "deposits" => Ok(GeneratedField::Deposits),
                            "refunds" => Ok(GeneratedField::Refunds),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgGuardianPause;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.MsgGuardianPause")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgGuardianPause, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut guardian__ = None;
                let mut deposits__ = None;
                let mut refunds__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Guardian => {
                            if guardian__.is_some() {
                                return Err(serde::de::Error::duplicate_field("guardian"));
                            }
                            guardian__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Deposits => {
                            if deposits__.is_some() {
                                return Err(serde::de::Error::duplicate_field("deposits"));
                            }
                            deposits__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Refunds => {
                            if refunds__.is_some() {
                                return Err(serde::de::Error::duplicate_field("refunds"));
                            }
                            refunds__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(MsgGuardianPause {
                    guardian: guardian__.unwrap_or_default(),
                    deposits: deposits__.unwrap_or_default(),
                    refunds: refunds__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.MsgGuardianPause", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgGuardianPauseResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser =
            serializer.serialize_struct("akash.sink.v1.MsgGuardianPauseResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgGuardianPauseResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {}
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
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
            type Value = MsgGuardianPauseResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.MsgGuardianPauseResponse")
            }

            fn visit_map<V>(
                self,
                mut map_: V,
            ) -> std::result::Result<MsgGuardianPauseResponse, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgGuardianPauseResponse {})
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.MsgGuardianPauseResponse",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for MsgReportOutcome {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.reporter.is_empty() {
            len += 1;
        }
        if self.id.is_some() {
            len += 1;
        }
        if self.policy_version != 0 {
            len += 1;
        }
        if self.round != 0 {
            len += 1;
        }
        if self.observation.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.MsgReportOutcome", len)?;
        if !self.reporter.is_empty() {
            struct_ser.serialize_field("reporter", &self.reporter)?;
        }
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if self.policy_version != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field(
                "policyVersion",
                ToString::to_string(&self.policy_version).as_str(),
            )?;
        }
        if self.round != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("round", ToString::to_string(&self.round).as_str())?;
        }
        if let Some(v) = self.observation.as_ref() {
            struct_ser.serialize_field("observation", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgReportOutcome {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "reporter",
            "id",
            "policy_version",
            "policyVersion",
            "round",
            "observation",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Reporter,
            Id,
            PolicyVersion,
            Round,
            Observation,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "reporter" => Ok(GeneratedField::Reporter),
                            "id" => Ok(GeneratedField::Id),
                            "policyVersion" | "policy_version" => Ok(GeneratedField::PolicyVersion),
                            "round" => Ok(GeneratedField::Round),
                            "observation" => Ok(GeneratedField::Observation),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgReportOutcome;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.MsgReportOutcome")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgReportOutcome, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut reporter__ = None;
                let mut id__ = None;
                let mut policy_version__ = None;
                let mut round__ = None;
                let mut observation__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Reporter => {
                            if reporter__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reporter"));
                            }
                            reporter__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = map_.next_value()?;
                        }
                        GeneratedField::PolicyVersion => {
                            if policy_version__.is_some() {
                                return Err(serde::de::Error::duplicate_field("policyVersion"));
                            }
                            policy_version__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                        GeneratedField::Round => {
                            if round__.is_some() {
                                return Err(serde::de::Error::duplicate_field("round"));
                            }
                            round__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                        GeneratedField::Observation => {
                            if observation__.is_some() {
                                return Err(serde::de::Error::duplicate_field("observation"));
                            }
                            observation__ = map_.next_value()?;
                        }
                    }
                }
                Ok(MsgReportOutcome {
                    reporter: reporter__.unwrap_or_default(),
                    id: id__,
                    policy_version: policy_version__.unwrap_or_default(),
                    round: round__.unwrap_or_default(),
                    observation: observation__,
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.MsgReportOutcome", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgReportOutcomeResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser =
            serializer.serialize_struct("akash.sink.v1.MsgReportOutcomeResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgReportOutcomeResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {}
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
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
            type Value = MsgReportOutcomeResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.MsgReportOutcomeResponse")
            }

            fn visit_map<V>(
                self,
                mut map_: V,
            ) -> std::result::Result<MsgReportOutcomeResponse, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgReportOutcomeResponse {})
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.MsgReportOutcomeResponse",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for MsgRequestVoid {
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
        if self.id.is_some() {
            len += 1;
        }
        if self.expected_round != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.MsgRequestVoid", len)?;
        if !self.authority.is_empty() {
            struct_ser.serialize_field("authority", &self.authority)?;
        }
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if self.expected_round != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field(
                "expectedRound",
                ToString::to_string(&self.expected_round).as_str(),
            )?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgRequestVoid {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["authority", "id", "expected_round", "expectedRound"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Authority,
            Id,
            ExpectedRound,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "authority" => Ok(GeneratedField::Authority),
                            "id" => Ok(GeneratedField::Id),
                            "expectedRound" | "expected_round" => Ok(GeneratedField::ExpectedRound),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgRequestVoid;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.MsgRequestVoid")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgRequestVoid, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut authority__ = None;
                let mut id__ = None;
                let mut expected_round__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Authority => {
                            if authority__.is_some() {
                                return Err(serde::de::Error::duplicate_field("authority"));
                            }
                            authority__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = map_.next_value()?;
                        }
                        GeneratedField::ExpectedRound => {
                            if expected_round__.is_some() {
                                return Err(serde::de::Error::duplicate_field("expectedRound"));
                            }
                            expected_round__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                    }
                }
                Ok(MsgRequestVoid {
                    authority: authority__.unwrap_or_default(),
                    id: id__,
                    expected_round: expected_round__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.MsgRequestVoid", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgRequestVoidResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser =
            serializer.serialize_struct("akash.sink.v1.MsgRequestVoidResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgRequestVoidResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {}
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
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
            type Value = MsgRequestVoidResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.MsgRequestVoidResponse")
            }

            fn visit_map<V>(
                self,
                mut map_: V,
            ) -> std::result::Result<MsgRequestVoidResponse, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgRequestVoidResponse {})
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.MsgRequestVoidResponse",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for MsgRestartTransfer {
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
        if self.id.is_some() {
            len += 1;
        }
        if !self.reason.is_empty() {
            len += 1;
        }
        if self.expected_round != 0 {
            len += 1;
        }
        let mut struct_ser =
            serializer.serialize_struct("akash.sink.v1.MsgRestartTransfer", len)?;
        if !self.authority.is_empty() {
            struct_ser.serialize_field("authority", &self.authority)?;
        }
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if !self.reason.is_empty() {
            struct_ser.serialize_field("reason", &self.reason)?;
        }
        if self.expected_round != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field(
                "expectedRound",
                ToString::to_string(&self.expected_round).as_str(),
            )?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgRestartTransfer {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "authority",
            "id",
            "reason",
            "expected_round",
            "expectedRound",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Authority,
            Id,
            Reason,
            ExpectedRound,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "authority" => Ok(GeneratedField::Authority),
                            "id" => Ok(GeneratedField::Id),
                            "reason" => Ok(GeneratedField::Reason),
                            "expectedRound" | "expected_round" => Ok(GeneratedField::ExpectedRound),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgRestartTransfer;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.MsgRestartTransfer")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgRestartTransfer, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut authority__ = None;
                let mut id__ = None;
                let mut reason__ = None;
                let mut expected_round__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Authority => {
                            if authority__.is_some() {
                                return Err(serde::de::Error::duplicate_field("authority"));
                            }
                            authority__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = map_.next_value()?;
                        }
                        GeneratedField::Reason => {
                            if reason__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reason"));
                            }
                            reason__ = Some(map_.next_value()?);
                        }
                        GeneratedField::ExpectedRound => {
                            if expected_round__.is_some() {
                                return Err(serde::de::Error::duplicate_field("expectedRound"));
                            }
                            expected_round__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                    }
                }
                Ok(MsgRestartTransfer {
                    authority: authority__.unwrap_or_default(),
                    id: id__,
                    reason: reason__.unwrap_or_default(),
                    expected_round: expected_round__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.MsgRestartTransfer",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for MsgRestartTransferResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser =
            serializer.serialize_struct("akash.sink.v1.MsgRestartTransferResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgRestartTransferResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {}
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
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
            type Value = MsgRestartTransferResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.MsgRestartTransferResponse")
            }

            fn visit_map<V>(
                self,
                mut map_: V,
            ) -> std::result::Result<MsgRestartTransferResponse, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgRestartTransferResponse {})
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.MsgRestartTransferResponse",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for MsgSink {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.dest_type != 0 {
            len += 1;
        }
        if !self.dest.is_empty() {
            len += 1;
        }
        if self.amount.is_some() {
            len += 1;
        }
        if !self.sender.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.MsgSink", len)?;
        if self.dest_type != 0 {
            let v = DestType::try_from(self.dest_type).map_err(|_| {
                serde::ser::Error::custom(format!("Invalid variant {}", self.dest_type))
            })?;
            struct_ser.serialize_field("destType", &v)?;
        }
        if !self.dest.is_empty() {
            struct_ser.serialize_field("dest", &self.dest)?;
        }
        if let Some(v) = self.amount.as_ref() {
            struct_ser.serialize_field("amount", v)?;
        }
        if !self.sender.is_empty() {
            struct_ser.serialize_field("sender", &self.sender)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgSink {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["dest_type", "destType", "dest", "amount", "sender"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            DestType,
            Dest,
            Amount,
            Sender,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "destType" | "dest_type" => Ok(GeneratedField::DestType),
                            "dest" => Ok(GeneratedField::Dest),
                            "amount" => Ok(GeneratedField::Amount),
                            "sender" => Ok(GeneratedField::Sender),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgSink;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.MsgSink")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgSink, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut dest_type__ = None;
                let mut dest__ = None;
                let mut amount__ = None;
                let mut sender__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::DestType => {
                            if dest_type__.is_some() {
                                return Err(serde::de::Error::duplicate_field("destType"));
                            }
                            dest_type__ = Some(map_.next_value::<DestType>()? as i32);
                        }
                        GeneratedField::Dest => {
                            if dest__.is_some() {
                                return Err(serde::de::Error::duplicate_field("dest"));
                            }
                            dest__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Amount => {
                            if amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("amount"));
                            }
                            amount__ = map_.next_value()?;
                        }
                        GeneratedField::Sender => {
                            if sender__.is_some() {
                                return Err(serde::de::Error::duplicate_field("sender"));
                            }
                            sender__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(MsgSink {
                    dest_type: dest_type__.unwrap_or_default(),
                    dest: dest__.unwrap_or_default(),
                    amount: amount__,
                    sender: sender__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.MsgSink", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgSinkResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.id.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.MsgSinkResponse", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgSinkResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["id"];

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

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
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
            type Value = MsgSinkResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.MsgSinkResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgSinkResponse, V::Error>
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
                            id__ = map_.next_value()?;
                        }
                    }
                }
                Ok(MsgSinkResponse { id: id__ })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.MsgSinkResponse", FIELDS, GeneratedVisitor)
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
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.MsgUpdateParams", len)?;
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
        const FIELDS: &[&str] = &["authority", "params"];

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

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
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
                formatter.write_str("struct akash.sink.v1.MsgUpdateParams")
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
        deserializer.deserialize_struct("akash.sink.v1.MsgUpdateParams", FIELDS, GeneratedVisitor)
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
        let struct_ser =
            serializer.serialize_struct("akash.sink.v1.MsgUpdateParamsResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgUpdateParamsResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {}
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
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
                formatter.write_str("struct akash.sink.v1.MsgUpdateParamsResponse")
            }

            fn visit_map<V>(
                self,
                mut map_: V,
            ) -> std::result::Result<MsgUpdateParamsResponse, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgUpdateParamsResponse {})
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.MsgUpdateParamsResponse",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for MsgVetoVoid {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.signer.is_empty() {
            len += 1;
        }
        if self.id.is_some() {
            len += 1;
        }
        if !self.reason.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.MsgVetoVoid", len)?;
        if !self.signer.is_empty() {
            struct_ser.serialize_field("signer", &self.signer)?;
        }
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if !self.reason.is_empty() {
            struct_ser.serialize_field("reason", &self.reason)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgVetoVoid {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["signer", "id", "reason"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Signer,
            Id,
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

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "signer" => Ok(GeneratedField::Signer),
                            "id" => Ok(GeneratedField::Id),
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
            type Value = MsgVetoVoid;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.MsgVetoVoid")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgVetoVoid, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut signer__ = None;
                let mut id__ = None;
                let mut reason__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Signer => {
                            if signer__.is_some() {
                                return Err(serde::de::Error::duplicate_field("signer"));
                            }
                            signer__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = map_.next_value()?;
                        }
                        GeneratedField::Reason => {
                            if reason__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reason"));
                            }
                            reason__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(MsgVetoVoid {
                    signer: signer__.unwrap_or_default(),
                    id: id__,
                    reason: reason__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.MsgVetoVoid", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgVetoVoidResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.sink.v1.MsgVetoVoidResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgVetoVoidResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {}
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
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
            type Value = MsgVetoVoidResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.MsgVetoVoidResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgVetoVoidResponse, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgVetoVoidResponse {})
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.MsgVetoVoidResponse",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for Observation {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.outcome != 0 {
            len += 1;
        }
        if self.slot != 0 {
            len += 1;
        }
        if !self.transaction_signature.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.Observation", len)?;
        if self.outcome != 0 {
            let v = Outcome::try_from(self.outcome).map_err(|_| {
                serde::ser::Error::custom(format!("Invalid variant {}", self.outcome))
            })?;
            struct_ser.serialize_field("outcome", &v)?;
        }
        if self.slot != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("slot", ToString::to_string(&self.slot).as_str())?;
        }
        if !self.transaction_signature.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field(
                "transactionSignature",
                pbjson::private::base64::encode(&self.transaction_signature).as_str(),
            )?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for Observation {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "outcome",
            "slot",
            "transaction_signature",
            "transactionSignature",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Outcome,
            Slot,
            TransactionSignature,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "outcome" => Ok(GeneratedField::Outcome),
                            "slot" => Ok(GeneratedField::Slot),
                            "transactionSignature" | "transaction_signature" => {
                                Ok(GeneratedField::TransactionSignature)
                            }
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = Observation;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.Observation")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<Observation, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut outcome__ = None;
                let mut slot__ = None;
                let mut transaction_signature__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Outcome => {
                            if outcome__.is_some() {
                                return Err(serde::de::Error::duplicate_field("outcome"));
                            }
                            outcome__ = Some(map_.next_value::<Outcome>()? as i32);
                        }
                        GeneratedField::Slot => {
                            if slot__.is_some() {
                                return Err(serde::de::Error::duplicate_field("slot"));
                            }
                            slot__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                        GeneratedField::TransactionSignature => {
                            if transaction_signature__.is_some() {
                                return Err(serde::de::Error::duplicate_field(
                                    "transactionSignature",
                                ));
                            }
                            transaction_signature__ = Some(
                                map_.next_value::<::pbjson::private::BytesDeserialize<_>>()?
                                    .0,
                            );
                        }
                    }
                }
                Ok(Observation {
                    outcome: outcome__.unwrap_or_default(),
                    slot: slot__.unwrap_or_default(),
                    transaction_signature: transaction_signature__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.Observation", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for Outcome {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "OUTCOME_UNSPECIFIED",
            Self::Minted => "OUTCOME_MINTED",
            Self::Voided => "OUTCOME_VOIDED",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for Outcome {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["OUTCOME_UNSPECIFIED", "OUTCOME_MINTED", "OUTCOME_VOIDED"];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = Outcome;

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
                    "OUTCOME_UNSPECIFIED" => Ok(Outcome::Unspecified),
                    "OUTCOME_MINTED" => Ok(Outcome::Minted),
                    "OUTCOME_VOIDED" => Ok(Outcome::Voided),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
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
        if self.enabled {
            len += 1;
        }
        if self.min_amount_per_transfer.is_some() {
            len += 1;
        }
        if self.max_amount_per_transfer.is_some() {
            len += 1;
        }
        if self.max_transfer_amount_per_epoch.is_some() {
            len += 1;
        }
        if self.max_epoch_amount_per_account.is_some() {
            len += 1;
        }
        if self.max_epoch_refund.is_some() {
            len += 1;
        }
        if !self.reporters.is_empty() {
            len += 1;
        }
        if !self.reporters_threshold.is_empty() {
            len += 1;
        }
        if self.reporter_challenge_period.is_some() {
            len += 1;
        }
        if !self.allowed_destination_vaults.is_empty() {
            len += 1;
        }
        if self.refunds_enabled {
            len += 1;
        }
        if !self.pause_guardian.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.Params", len)?;
        if self.enabled {
            struct_ser.serialize_field("enabled", &self.enabled)?;
        }
        if let Some(v) = self.min_amount_per_transfer.as_ref() {
            struct_ser.serialize_field("minAmountPerTransfer", v)?;
        }
        if let Some(v) = self.max_amount_per_transfer.as_ref() {
            struct_ser.serialize_field("maxAmountPerTransfer", v)?;
        }
        if let Some(v) = self.max_transfer_amount_per_epoch.as_ref() {
            struct_ser.serialize_field("maxTransferAmountPerEpoch", v)?;
        }
        if let Some(v) = self.max_epoch_amount_per_account.as_ref() {
            struct_ser.serialize_field("maxEpochAmountPerAccount", v)?;
        }
        if let Some(v) = self.max_epoch_refund.as_ref() {
            struct_ser.serialize_field("maxEpochRefund", v)?;
        }
        if !self.reporters.is_empty() {
            struct_ser.serialize_field("reporters", &self.reporters)?;
        }
        if !self.reporters_threshold.is_empty() {
            struct_ser.serialize_field(
                "reportersThreshold",
                &sink_legacy_dec::to_json(&self.reporters_threshold)
                    .map_err(serde::ser::Error::custom)?,
            )?;
        }
        if let Some(v) = self.reporter_challenge_period.as_ref() {
            struct_ser.serialize_field("reporterChallengePeriod", v)?;
        }
        if !self.allowed_destination_vaults.is_empty() {
            struct_ser
                .serialize_field("allowedDestinationVaults", &self.allowed_destination_vaults)?;
        }
        if self.refunds_enabled {
            struct_ser.serialize_field("refundsEnabled", &self.refunds_enabled)?;
        }
        if !self.pause_guardian.is_empty() {
            struct_ser.serialize_field("pauseGuardian", &self.pause_guardian)?;
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
            "enabled",
            "min_amount_per_transfer",
            "minAmountPerTransfer",
            "max_amount_per_transfer",
            "maxAmountPerTransfer",
            "max_transfer_amount_per_epoch",
            "maxTransferAmountPerEpoch",
            "max_epoch_amount_per_account",
            "maxEpochAmountPerAccount",
            "max_epoch_refund",
            "maxEpochRefund",
            "reporters",
            "reporters_threshold",
            "reportersThreshold",
            "reporter_challenge_period",
            "reporterChallengePeriod",
            "allowed_destination_vaults",
            "allowedDestinationVaults",
            "refunds_enabled",
            "refundsEnabled",
            "pause_guardian",
            "pauseGuardian",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Enabled,
            MinAmountPerTransfer,
            MaxAmountPerTransfer,
            MaxTransferAmountPerEpoch,
            MaxEpochAmountPerAccount,
            MaxEpochRefund,
            Reporters,
            ReportersThreshold,
            ReporterChallengePeriod,
            AllowedDestinationVaults,
            RefundsEnabled,
            PauseGuardian,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "enabled" => Ok(GeneratedField::Enabled),
                            "minAmountPerTransfer" | "min_amount_per_transfer" => {
                                Ok(GeneratedField::MinAmountPerTransfer)
                            }
                            "maxAmountPerTransfer" | "max_amount_per_transfer" => {
                                Ok(GeneratedField::MaxAmountPerTransfer)
                            }
                            "maxTransferAmountPerEpoch" | "max_transfer_amount_per_epoch" => {
                                Ok(GeneratedField::MaxTransferAmountPerEpoch)
                            }
                            "maxEpochAmountPerAccount" | "max_epoch_amount_per_account" => {
                                Ok(GeneratedField::MaxEpochAmountPerAccount)
                            }
                            "maxEpochRefund" | "max_epoch_refund" => {
                                Ok(GeneratedField::MaxEpochRefund)
                            }
                            "reporters" => Ok(GeneratedField::Reporters),
                            "reportersThreshold" | "reporters_threshold" => {
                                Ok(GeneratedField::ReportersThreshold)
                            }
                            "reporterChallengePeriod" | "reporter_challenge_period" => {
                                Ok(GeneratedField::ReporterChallengePeriod)
                            }
                            "allowedDestinationVaults" | "allowed_destination_vaults" => {
                                Ok(GeneratedField::AllowedDestinationVaults)
                            }
                            "refundsEnabled" | "refunds_enabled" => {
                                Ok(GeneratedField::RefundsEnabled)
                            }
                            "pauseGuardian" | "pause_guardian" => Ok(GeneratedField::PauseGuardian),
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
                formatter.write_str("struct akash.sink.v1.Params")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<Params, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut enabled__ = None;
                let mut min_amount_per_transfer__ = None;
                let mut max_amount_per_transfer__ = None;
                let mut max_transfer_amount_per_epoch__ = None;
                let mut max_epoch_amount_per_account__ = None;
                let mut max_epoch_refund__ = None;
                let mut reporters__ = None;
                let mut reporters_threshold__ = None;
                let mut reporter_challenge_period__ = None;
                let mut allowed_destination_vaults__ = None;
                let mut refunds_enabled__ = None;
                let mut pause_guardian__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Enabled => {
                            if enabled__.is_some() {
                                return Err(serde::de::Error::duplicate_field("enabled"));
                            }
                            enabled__ = Some(map_.next_value()?);
                        }
                        GeneratedField::MinAmountPerTransfer => {
                            if min_amount_per_transfer__.is_some() {
                                return Err(serde::de::Error::duplicate_field(
                                    "minAmountPerTransfer",
                                ));
                            }
                            min_amount_per_transfer__ = map_.next_value()?;
                        }
                        GeneratedField::MaxAmountPerTransfer => {
                            if max_amount_per_transfer__.is_some() {
                                return Err(serde::de::Error::duplicate_field(
                                    "maxAmountPerTransfer",
                                ));
                            }
                            max_amount_per_transfer__ = map_.next_value()?;
                        }
                        GeneratedField::MaxTransferAmountPerEpoch => {
                            if max_transfer_amount_per_epoch__.is_some() {
                                return Err(serde::de::Error::duplicate_field(
                                    "maxTransferAmountPerEpoch",
                                ));
                            }
                            max_transfer_amount_per_epoch__ = map_.next_value()?;
                        }
                        GeneratedField::MaxEpochAmountPerAccount => {
                            if max_epoch_amount_per_account__.is_some() {
                                return Err(serde::de::Error::duplicate_field(
                                    "maxEpochAmountPerAccount",
                                ));
                            }
                            max_epoch_amount_per_account__ = map_.next_value()?;
                        }
                        GeneratedField::MaxEpochRefund => {
                            if max_epoch_refund__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maxEpochRefund"));
                            }
                            max_epoch_refund__ = map_.next_value()?;
                        }
                        GeneratedField::Reporters => {
                            if reporters__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reporters"));
                            }
                            reporters__ = Some(map_.next_value()?);
                        }
                        GeneratedField::ReportersThreshold => {
                            if reporters_threshold__.is_some() {
                                return Err(serde::de::Error::duplicate_field(
                                    "reportersThreshold",
                                ));
                            }
                            reporters_threshold__ = Some(
                                sink_legacy_dec::from_json(&map_.next_value::<String>()?)
                                    .map_err(serde::de::Error::custom)?,
                            );
                        }
                        GeneratedField::ReporterChallengePeriod => {
                            if reporter_challenge_period__.is_some() {
                                return Err(serde::de::Error::duplicate_field(
                                    "reporterChallengePeriod",
                                ));
                            }
                            reporter_challenge_period__ = map_.next_value()?;
                        }
                        GeneratedField::AllowedDestinationVaults => {
                            if allowed_destination_vaults__.is_some() {
                                return Err(serde::de::Error::duplicate_field(
                                    "allowedDestinationVaults",
                                ));
                            }
                            allowed_destination_vaults__ = Some(map_.next_value()?);
                        }
                        GeneratedField::RefundsEnabled => {
                            if refunds_enabled__.is_some() {
                                return Err(serde::de::Error::duplicate_field("refundsEnabled"));
                            }
                            refunds_enabled__ = Some(map_.next_value()?);
                        }
                        GeneratedField::PauseGuardian => {
                            if pause_guardian__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pauseGuardian"));
                            }
                            pause_guardian__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(Params {
                    enabled: enabled__.unwrap_or_default(),
                    min_amount_per_transfer: min_amount_per_transfer__,
                    max_amount_per_transfer: max_amount_per_transfer__,
                    max_transfer_amount_per_epoch: max_transfer_amount_per_epoch__,
                    max_epoch_amount_per_account: max_epoch_amount_per_account__,
                    max_epoch_refund: max_epoch_refund__,
                    reporters: reporters__.unwrap_or_default(),
                    reporters_threshold: reporters_threshold__.unwrap_or_default(),
                    reporter_challenge_period: reporter_challenge_period__,
                    allowed_destination_vaults: allowed_destination_vaults__.unwrap_or_default(),
                    refunds_enabled: refunds_enabled__.unwrap_or_default(),
                    pause_guardian: pause_guardian__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.Params", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryAccountRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.owner.is_empty() {
            len += 1;
        }
        let mut struct_ser =
            serializer.serialize_struct("akash.sink.v1.QueryAccountRequest", len)?;
        if !self.owner.is_empty() {
            struct_ser.serialize_field("owner", &self.owner)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryAccountRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["owner"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Owner,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "owner" => Ok(GeneratedField::Owner),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryAccountRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.QueryAccountRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryAccountRequest, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut owner__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Owner => {
                            if owner__.is_some() {
                                return Err(serde::de::Error::duplicate_field("owner"));
                            }
                            owner__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(QueryAccountRequest {
                    owner: owner__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.QueryAccountRequest",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for QueryAccountResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.usage.is_some() {
            len += 1;
        }
        if self.epoch != 0 {
            len += 1;
        }
        let mut struct_ser =
            serializer.serialize_struct("akash.sink.v1.QueryAccountResponse", len)?;
        if let Some(v) = self.usage.as_ref() {
            struct_ser.serialize_field("usage", v)?;
        }
        if self.epoch != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("epoch", ToString::to_string(&self.epoch).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryAccountResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["usage", "epoch"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Usage,
            Epoch,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "usage" => Ok(GeneratedField::Usage),
                            "epoch" => Ok(GeneratedField::Epoch),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryAccountResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.QueryAccountResponse")
            }

            fn visit_map<V>(
                self,
                mut map_: V,
            ) -> std::result::Result<QueryAccountResponse, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut usage__ = None;
                let mut epoch__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Usage => {
                            if usage__.is_some() {
                                return Err(serde::de::Error::duplicate_field("usage"));
                            }
                            usage__ = map_.next_value()?;
                        }
                        GeneratedField::Epoch => {
                            if epoch__.is_some() {
                                return Err(serde::de::Error::duplicate_field("epoch"));
                            }
                            epoch__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                    }
                }
                Ok(QueryAccountResponse {
                    usage: usage__,
                    epoch: epoch__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.QueryAccountResponse",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for QueryDisputesRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.pagination.is_some() {
            len += 1;
        }
        let mut struct_ser =
            serializer.serialize_struct("akash.sink.v1.QueryDisputesRequest", len)?;
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryDisputesRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["pagination"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
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

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
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
            type Value = QueryDisputesRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.QueryDisputesRequest")
            }

            fn visit_map<V>(
                self,
                mut map_: V,
            ) -> std::result::Result<QueryDisputesRequest, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut pagination__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Pagination => {
                            if pagination__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pagination"));
                            }
                            pagination__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryDisputesRequest {
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.QueryDisputesRequest",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for QueryDisputesResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.transfers.is_empty() {
            len += 1;
        }
        if self.pagination.is_some() {
            len += 1;
        }
        let mut struct_ser =
            serializer.serialize_struct("akash.sink.v1.QueryDisputesResponse", len)?;
        if !self.transfers.is_empty() {
            struct_ser.serialize_field("transfers", &self.transfers)?;
        }
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryDisputesResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["transfers", "pagination"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Transfers,
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

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "transfers" => Ok(GeneratedField::Transfers),
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
            type Value = QueryDisputesResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.QueryDisputesResponse")
            }

            fn visit_map<V>(
                self,
                mut map_: V,
            ) -> std::result::Result<QueryDisputesResponse, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut transfers__ = None;
                let mut pagination__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Transfers => {
                            if transfers__.is_some() {
                                return Err(serde::de::Error::duplicate_field("transfers"));
                            }
                            transfers__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Pagination => {
                            if pagination__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pagination"));
                            }
                            pagination__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryDisputesResponse {
                    transfers: transfers__.unwrap_or_default(),
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.QueryDisputesResponse",
            FIELDS,
            GeneratedVisitor,
        )
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
        let struct_ser = serializer.serialize_struct("akash.sink.v1.QueryParamsRequest", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryParamsRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {}
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
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
                formatter.write_str("struct akash.sink.v1.QueryParamsRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryParamsRequest, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(QueryParamsRequest {})
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.QueryParamsRequest",
            FIELDS,
            GeneratedVisitor,
        )
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
        let mut struct_ser =
            serializer.serialize_struct("akash.sink.v1.QueryParamsResponse", len)?;
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
        const FIELDS: &[&str] = &["params"];

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

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
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
                formatter.write_str("struct akash.sink.v1.QueryParamsResponse")
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
                Ok(QueryParamsResponse { params: params__ })
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.QueryParamsResponse",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for QueryReportersRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.version != 0 {
            len += 1;
        }
        let mut struct_ser =
            serializer.serialize_struct("akash.sink.v1.QueryReportersRequest", len)?;
        if self.version != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("version", ToString::to_string(&self.version).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryReportersRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["version"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Version,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "version" => Ok(GeneratedField::Version),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryReportersRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.QueryReportersRequest")
            }

            fn visit_map<V>(
                self,
                mut map_: V,
            ) -> std::result::Result<QueryReportersRequest, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut version__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Version => {
                            if version__.is_some() {
                                return Err(serde::de::Error::duplicate_field("version"));
                            }
                            version__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                    }
                }
                Ok(QueryReportersRequest {
                    version: version__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.QueryReportersRequest",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for QueryReportersResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.reporters.is_empty() {
            len += 1;
        }
        let mut struct_ser =
            serializer.serialize_struct("akash.sink.v1.QueryReportersResponse", len)?;
        if !self.reporters.is_empty() {
            struct_ser.serialize_field("reporters", &self.reporters)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryReportersResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["reporters"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Reporters,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "reporters" => Ok(GeneratedField::Reporters),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryReportersResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.QueryReportersResponse")
            }

            fn visit_map<V>(
                self,
                mut map_: V,
            ) -> std::result::Result<QueryReportersResponse, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut reporters__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Reporters => {
                            if reporters__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reporters"));
                            }
                            reporters__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(QueryReportersResponse {
                    reporters: reporters__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.QueryReportersResponse",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for QueryStateRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.sink.v1.QueryStateRequest", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryStateRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {}
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
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
            type Value = QueryStateRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.QueryStateRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryStateRequest, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(QueryStateRequest {})
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.QueryStateRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryStateResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.totals.is_some() {
            len += 1;
        }
        if self.epoch != 0 {
            len += 1;
        }
        if self.send_usage.is_some() {
            len += 1;
        }
        if self.refund_usage.is_some() {
            len += 1;
        }
        if self.refund_head.is_some() {
            len += 1;
        }
        let mut struct_ser =
            serializer.serialize_struct("akash.sink.v1.QueryStateResponse", len)?;
        if let Some(v) = self.totals.as_ref() {
            struct_ser.serialize_field("totals", v)?;
        }
        if self.epoch != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("epoch", ToString::to_string(&self.epoch).as_str())?;
        }
        if let Some(v) = self.send_usage.as_ref() {
            struct_ser.serialize_field("sendUsage", v)?;
        }
        if let Some(v) = self.refund_usage.as_ref() {
            struct_ser.serialize_field("refundUsage", v)?;
        }
        if let Some(v) = self.refund_head.as_ref() {
            struct_ser.serialize_field("refundHead", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryStateResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "totals",
            "epoch",
            "send_usage",
            "sendUsage",
            "refund_usage",
            "refundUsage",
            "refund_head",
            "refundHead",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Totals,
            Epoch,
            SendUsage,
            RefundUsage,
            RefundHead,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "totals" => Ok(GeneratedField::Totals),
                            "epoch" => Ok(GeneratedField::Epoch),
                            "sendUsage" | "send_usage" => Ok(GeneratedField::SendUsage),
                            "refundUsage" | "refund_usage" => Ok(GeneratedField::RefundUsage),
                            "refundHead" | "refund_head" => Ok(GeneratedField::RefundHead),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryStateResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.QueryStateResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryStateResponse, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut totals__ = None;
                let mut epoch__ = None;
                let mut send_usage__ = None;
                let mut refund_usage__ = None;
                let mut refund_head__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Totals => {
                            if totals__.is_some() {
                                return Err(serde::de::Error::duplicate_field("totals"));
                            }
                            totals__ = map_.next_value()?;
                        }
                        GeneratedField::Epoch => {
                            if epoch__.is_some() {
                                return Err(serde::de::Error::duplicate_field("epoch"));
                            }
                            epoch__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                        GeneratedField::SendUsage => {
                            if send_usage__.is_some() {
                                return Err(serde::de::Error::duplicate_field("sendUsage"));
                            }
                            send_usage__ = map_.next_value()?;
                        }
                        GeneratedField::RefundUsage => {
                            if refund_usage__.is_some() {
                                return Err(serde::de::Error::duplicate_field("refundUsage"));
                            }
                            refund_usage__ = map_.next_value()?;
                        }
                        GeneratedField::RefundHead => {
                            if refund_head__.is_some() {
                                return Err(serde::de::Error::duplicate_field("refundHead"));
                            }
                            refund_head__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryStateResponse {
                    totals: totals__,
                    epoch: epoch__.unwrap_or_default(),
                    send_usage: send_usage__,
                    refund_usage: refund_usage__,
                    refund_head: refund_head__,
                })
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.QueryStateResponse",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for QueryTransferRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.id.is_some() {
            len += 1;
        }
        let mut struct_ser =
            serializer.serialize_struct("akash.sink.v1.QueryTransferRequest", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryTransferRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["id"];

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

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
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
            type Value = QueryTransferRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.QueryTransferRequest")
            }

            fn visit_map<V>(
                self,
                mut map_: V,
            ) -> std::result::Result<QueryTransferRequest, V::Error>
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
                            id__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryTransferRequest { id: id__ })
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.QueryTransferRequest",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for QueryTransferResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.transfer.is_some() {
            len += 1;
        }
        let mut struct_ser =
            serializer.serialize_struct("akash.sink.v1.QueryTransferResponse", len)?;
        if let Some(v) = self.transfer.as_ref() {
            struct_ser.serialize_field("transfer", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryTransferResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["transfer"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Transfer,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "transfer" => Ok(GeneratedField::Transfer),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryTransferResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.QueryTransferResponse")
            }

            fn visit_map<V>(
                self,
                mut map_: V,
            ) -> std::result::Result<QueryTransferResponse, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut transfer__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Transfer => {
                            if transfer__.is_some() {
                                return Err(serde::de::Error::duplicate_field("transfer"));
                            }
                            transfer__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryTransferResponse {
                    transfer: transfer__,
                })
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.QueryTransferResponse",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for QueryTransfersRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.owner.is_empty() {
            len += 1;
        }
        if self.status != 0 {
            len += 1;
        }
        if self.pagination.is_some() {
            len += 1;
        }
        let mut struct_ser =
            serializer.serialize_struct("akash.sink.v1.QueryTransfersRequest", len)?;
        if !self.owner.is_empty() {
            struct_ser.serialize_field("owner", &self.owner)?;
        }
        if self.status != 0 {
            let v = Status::try_from(self.status).map_err(|_| {
                serde::ser::Error::custom(format!("Invalid variant {}", self.status))
            })?;
            struct_ser.serialize_field("status", &v)?;
        }
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryTransfersRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["owner", "status", "pagination"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Owner,
            Status,
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

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "owner" => Ok(GeneratedField::Owner),
                            "status" => Ok(GeneratedField::Status),
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
            type Value = QueryTransfersRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.QueryTransfersRequest")
            }

            fn visit_map<V>(
                self,
                mut map_: V,
            ) -> std::result::Result<QueryTransfersRequest, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut owner__ = None;
                let mut status__ = None;
                let mut pagination__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Owner => {
                            if owner__.is_some() {
                                return Err(serde::de::Error::duplicate_field("owner"));
                            }
                            owner__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Status => {
                            if status__.is_some() {
                                return Err(serde::de::Error::duplicate_field("status"));
                            }
                            status__ = Some(map_.next_value::<Status>()? as i32);
                        }
                        GeneratedField::Pagination => {
                            if pagination__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pagination"));
                            }
                            pagination__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryTransfersRequest {
                    owner: owner__.unwrap_or_default(),
                    status: status__.unwrap_or_default(),
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.QueryTransfersRequest",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for QueryTransfersResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.transfers.is_empty() {
            len += 1;
        }
        if self.pagination.is_some() {
            len += 1;
        }
        let mut struct_ser =
            serializer.serialize_struct("akash.sink.v1.QueryTransfersResponse", len)?;
        if !self.transfers.is_empty() {
            struct_ser.serialize_field("transfers", &self.transfers)?;
        }
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryTransfersResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["transfers", "pagination"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Transfers,
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

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "transfers" => Ok(GeneratedField::Transfers),
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
            type Value = QueryTransfersResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.QueryTransfersResponse")
            }

            fn visit_map<V>(
                self,
                mut map_: V,
            ) -> std::result::Result<QueryTransfersResponse, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut transfers__ = None;
                let mut pagination__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Transfers => {
                            if transfers__.is_some() {
                                return Err(serde::de::Error::duplicate_field("transfers"));
                            }
                            transfers__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Pagination => {
                            if pagination__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pagination"));
                            }
                            pagination__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryTransfersResponse {
                    transfers: transfers__.unwrap_or_default(),
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct(
            "akash.sink.v1.QueryTransfersResponse",
            FIELDS,
            GeneratedVisitor,
        )
    }
}
impl serde::Serialize for RecordData {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.dest.is_empty() {
            len += 1;
        }
        if self.amount.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.RecordData", len)?;
        if !self.dest.is_empty() {
            struct_ser.serialize_field("dest", &self.dest)?;
        }
        if let Some(v) = self.amount.as_ref() {
            struct_ser.serialize_field("amount", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for RecordData {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["dest", "amount"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Dest,
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

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "dest" => Ok(GeneratedField::Dest),
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
            type Value = RecordData;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.RecordData")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<RecordData, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut dest__ = None;
                let mut amount__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Dest => {
                            if dest__.is_some() {
                                return Err(serde::de::Error::duplicate_field("dest"));
                            }
                            dest__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Amount => {
                            if amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("amount"));
                            }
                            amount__ = map_.next_value()?;
                        }
                    }
                }
                Ok(RecordData {
                    dest: dest__.unwrap_or_default(),
                    amount: amount__,
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.RecordData", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for ReporterPolicy {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.version != 0 {
            len += 1;
        }
        if !self.members.is_empty() {
            len += 1;
        }
        if self.threshold != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.ReporterPolicy", len)?;
        if self.version != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("version", ToString::to_string(&self.version).as_str())?;
        }
        if !self.members.is_empty() {
            struct_ser.serialize_field("members", &self.members)?;
        }
        if self.threshold != 0 {
            struct_ser.serialize_field("threshold", &self.threshold)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for ReporterPolicy {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["version", "members", "threshold"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Version,
            Members,
            Threshold,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "version" => Ok(GeneratedField::Version),
                            "members" => Ok(GeneratedField::Members),
                            "threshold" => Ok(GeneratedField::Threshold),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = ReporterPolicy;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.ReporterPolicy")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<ReporterPolicy, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut version__ = None;
                let mut members__ = None;
                let mut threshold__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Version => {
                            if version__.is_some() {
                                return Err(serde::de::Error::duplicate_field("version"));
                            }
                            version__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                        GeneratedField::Members => {
                            if members__.is_some() {
                                return Err(serde::de::Error::duplicate_field("members"));
                            }
                            members__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Threshold => {
                            if threshold__.is_some() {
                                return Err(serde::de::Error::duplicate_field("threshold"));
                            }
                            threshold__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                    }
                }
                Ok(ReporterPolicy {
                    version: version__.unwrap_or_default(),
                    members: members__.unwrap_or_default(),
                    threshold: threshold__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.ReporterPolicy", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for SettlementData {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.hold != 0 {
            len += 1;
        }
        if self.round != 0 {
            len += 1;
        }
        if self.policy.is_some() {
            len += 1;
        }
        if !self.votes.is_empty() {
            len += 1;
        }
        if self.challenge_end != 0 {
            len += 1;
        }
        if self.governance_observation.is_some() {
            len += 1;
        }
        if !self.veto_by.is_empty() {
            len += 1;
        }
        if !self.veto_reason.is_empty() {
            len += 1;
        }
        if self.void_authorization.is_some() {
            len += 1;
        }
        if self.veto_policy_version != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.SettlementData", len)?;
        if self.hold != 0 {
            let v = Hold::try_from(self.hold)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.hold)))?;
            struct_ser.serialize_field("hold", &v)?;
        }
        if self.round != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("round", ToString::to_string(&self.round).as_str())?;
        }
        if let Some(v) = self.policy.as_ref() {
            struct_ser.serialize_field("policy", v)?;
        }
        if !self.votes.is_empty() {
            struct_ser.serialize_field("votes", &self.votes)?;
        }
        if self.challenge_end != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field(
                "challengeEnd",
                ToString::to_string(&self.challenge_end).as_str(),
            )?;
        }
        if let Some(v) = self.governance_observation.as_ref() {
            struct_ser.serialize_field("governanceObservation", v)?;
        }
        if !self.veto_by.is_empty() {
            struct_ser.serialize_field("vetoBy", &self.veto_by)?;
        }
        if !self.veto_reason.is_empty() {
            struct_ser.serialize_field("vetoReason", &self.veto_reason)?;
        }
        if let Some(v) = self.void_authorization.as_ref() {
            struct_ser.serialize_field("voidAuthorization", v)?;
        }
        if self.veto_policy_version != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field(
                "vetoPolicyVersion",
                ToString::to_string(&self.veto_policy_version).as_str(),
            )?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for SettlementData {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "hold",
            "round",
            "policy",
            "votes",
            "challenge_end",
            "challengeEnd",
            "governance_observation",
            "governanceObservation",
            "veto_by",
            "vetoBy",
            "veto_reason",
            "vetoReason",
            "void_authorization",
            "voidAuthorization",
            "veto_policy_version",
            "vetoPolicyVersion",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Hold,
            Round,
            Policy,
            Votes,
            ChallengeEnd,
            GovernanceObservation,
            VetoBy,
            VetoReason,
            VoidAuthorization,
            VetoPolicyVersion,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "hold" => Ok(GeneratedField::Hold),
                            "round" => Ok(GeneratedField::Round),
                            "policy" => Ok(GeneratedField::Policy),
                            "votes" => Ok(GeneratedField::Votes),
                            "challengeEnd" | "challenge_end" => Ok(GeneratedField::ChallengeEnd),
                            "governanceObservation" | "governance_observation" => {
                                Ok(GeneratedField::GovernanceObservation)
                            }
                            "vetoBy" | "veto_by" => Ok(GeneratedField::VetoBy),
                            "vetoReason" | "veto_reason" => Ok(GeneratedField::VetoReason),
                            "voidAuthorization" | "void_authorization" => {
                                Ok(GeneratedField::VoidAuthorization)
                            }
                            "vetoPolicyVersion" | "veto_policy_version" => {
                                Ok(GeneratedField::VetoPolicyVersion)
                            }
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = SettlementData;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.SettlementData")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<SettlementData, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut hold__ = None;
                let mut round__ = None;
                let mut policy__ = None;
                let mut votes__ = None;
                let mut challenge_end__ = None;
                let mut governance_observation__ = None;
                let mut veto_by__ = None;
                let mut veto_reason__ = None;
                let mut void_authorization__ = None;
                let mut veto_policy_version__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Hold => {
                            if hold__.is_some() {
                                return Err(serde::de::Error::duplicate_field("hold"));
                            }
                            hold__ = Some(map_.next_value::<Hold>()? as i32);
                        }
                        GeneratedField::Round => {
                            if round__.is_some() {
                                return Err(serde::de::Error::duplicate_field("round"));
                            }
                            round__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                        GeneratedField::Policy => {
                            if policy__.is_some() {
                                return Err(serde::de::Error::duplicate_field("policy"));
                            }
                            policy__ = map_.next_value()?;
                        }
                        GeneratedField::Votes => {
                            if votes__.is_some() {
                                return Err(serde::de::Error::duplicate_field("votes"));
                            }
                            votes__ = Some(map_.next_value()?);
                        }
                        GeneratedField::ChallengeEnd => {
                            if challenge_end__.is_some() {
                                return Err(serde::de::Error::duplicate_field("challengeEnd"));
                            }
                            challenge_end__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                        GeneratedField::GovernanceObservation => {
                            if governance_observation__.is_some() {
                                return Err(serde::de::Error::duplicate_field(
                                    "governanceObservation",
                                ));
                            }
                            governance_observation__ = map_.next_value()?;
                        }
                        GeneratedField::VetoBy => {
                            if veto_by__.is_some() {
                                return Err(serde::de::Error::duplicate_field("vetoBy"));
                            }
                            veto_by__ = Some(map_.next_value()?);
                        }
                        GeneratedField::VetoReason => {
                            if veto_reason__.is_some() {
                                return Err(serde::de::Error::duplicate_field("vetoReason"));
                            }
                            veto_reason__ = Some(map_.next_value()?);
                        }
                        GeneratedField::VoidAuthorization => {
                            if void_authorization__.is_some() {
                                return Err(serde::de::Error::duplicate_field("voidAuthorization"));
                            }
                            void_authorization__ = map_.next_value()?;
                        }
                        GeneratedField::VetoPolicyVersion => {
                            if veto_policy_version__.is_some() {
                                return Err(serde::de::Error::duplicate_field("vetoPolicyVersion"));
                            }
                            veto_policy_version__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                    }
                }
                Ok(SettlementData {
                    hold: hold__.unwrap_or_default(),
                    round: round__.unwrap_or_default(),
                    policy: policy__,
                    votes: votes__.unwrap_or_default(),
                    challenge_end: challenge_end__.unwrap_or_default(),
                    governance_observation: governance_observation__,
                    veto_by: veto_by__.unwrap_or_default(),
                    veto_reason: veto_reason__.unwrap_or_default(),
                    void_authorization: void_authorization__,
                    veto_policy_version: veto_policy_version__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.SettlementData", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for SettlementId {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.status != 0 {
            len += 1;
        }
        if !self.owner.is_empty() {
            len += 1;
        }
        if self.height != 0 {
            len += 1;
        }
        if self.sequence != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.SettlementID", len)?;
        if self.status != 0 {
            let v = Status::try_from(self.status).map_err(|_| {
                serde::ser::Error::custom(format!("Invalid variant {}", self.status))
            })?;
            struct_ser.serialize_field("status", &v)?;
        }
        if !self.owner.is_empty() {
            struct_ser.serialize_field("owner", &self.owner)?;
        }
        if self.height != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("height", ToString::to_string(&self.height).as_str())?;
        }
        if self.sequence != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("sequence", ToString::to_string(&self.sequence).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for SettlementId {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["status", "owner", "height", "sequence"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Status,
            Owner,
            Height,
            Sequence,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "status" => Ok(GeneratedField::Status),
                            "owner" => Ok(GeneratedField::Owner),
                            "height" => Ok(GeneratedField::Height),
                            "sequence" => Ok(GeneratedField::Sequence),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = SettlementId;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.SettlementID")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<SettlementId, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut status__ = None;
                let mut owner__ = None;
                let mut height__ = None;
                let mut sequence__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Status => {
                            if status__.is_some() {
                                return Err(serde::de::Error::duplicate_field("status"));
                            }
                            status__ = Some(map_.next_value::<Status>()? as i32);
                        }
                        GeneratedField::Owner => {
                            if owner__.is_some() {
                                return Err(serde::de::Error::duplicate_field("owner"));
                            }
                            owner__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Height => {
                            if height__.is_some() {
                                return Err(serde::de::Error::duplicate_field("height"));
                            }
                            height__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                        GeneratedField::Sequence => {
                            if sequence__.is_some() {
                                return Err(serde::de::Error::duplicate_field("sequence"));
                            }
                            sequence__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                    }
                }
                Ok(SettlementId {
                    status: status__.unwrap_or_default(),
                    owner: owner__.unwrap_or_default(),
                    height: height__.unwrap_or_default(),
                    sequence: sequence__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.SettlementID", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for SolanaEvidence {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.slot != 0 {
            len += 1;
        }
        if !self.transaction_signature.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.SolanaEvidence", len)?;
        if self.slot != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("slot", ToString::to_string(&self.slot).as_str())?;
        }
        if !self.transaction_signature.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field(
                "transactionSignature",
                pbjson::private::base64::encode(&self.transaction_signature).as_str(),
            )?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for SolanaEvidence {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["slot", "transaction_signature", "transactionSignature"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Slot,
            TransactionSignature,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "slot" => Ok(GeneratedField::Slot),
                            "transactionSignature" | "transaction_signature" => {
                                Ok(GeneratedField::TransactionSignature)
                            }
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = SolanaEvidence;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.SolanaEvidence")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<SolanaEvidence, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut slot__ = None;
                let mut transaction_signature__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Slot => {
                            if slot__.is_some() {
                                return Err(serde::de::Error::duplicate_field("slot"));
                            }
                            slot__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                        GeneratedField::TransactionSignature => {
                            if transaction_signature__.is_some() {
                                return Err(serde::de::Error::duplicate_field(
                                    "transactionSignature",
                                ));
                            }
                            transaction_signature__ = Some(
                                map_.next_value::<::pbjson::private::BytesDeserialize<_>>()?
                                    .0,
                            );
                        }
                    }
                }
                Ok(SolanaEvidence {
                    slot: slot__.unwrap_or_default(),
                    transaction_signature: transaction_signature__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.SolanaEvidence", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for Status {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Invalid => "STATUS_INVALID",
            Self::TransferPending => "STATUS_TRANSFER_PENDING",
            Self::Minted => "STATUS_MINTED",
            Self::Voided => "STATUS_VOIDED",
            Self::RefundPending => "STATUS_REFUND_PENDING",
            Self::Refunded => "STATUS_REFUNDED",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for Status {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "STATUS_INVALID",
            "STATUS_TRANSFER_PENDING",
            "STATUS_MINTED",
            "STATUS_VOIDED",
            "STATUS_REFUND_PENDING",
            "STATUS_REFUNDED",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = Status;

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
                    "STATUS_INVALID" => Ok(Status::Invalid),
                    "STATUS_TRANSFER_PENDING" => Ok(Status::TransferPending),
                    "STATUS_MINTED" => Ok(Status::Minted),
                    "STATUS_VOIDED" => Ok(Status::Voided),
                    "STATUS_REFUND_PENDING" => Ok(Status::RefundPending),
                    "STATUS_REFUNDED" => Ok(Status::Refunded),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for Totals {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.pending_amount.is_some() {
            len += 1;
        }
        if self.pending_count != 0 {
            len += 1;
        }
        if self.minted_amount.is_some() {
            len += 1;
        }
        if self.minted_count != 0 {
            len += 1;
        }
        if self.voided_amount.is_some() {
            len += 1;
        }
        if self.voided_count != 0 {
            len += 1;
        }
        if self.refund_pending_amount.is_some() {
            len += 1;
        }
        if self.refund_pending_count != 0 {
            len += 1;
        }
        if self.refunded_amount.is_some() {
            len += 1;
        }
        if self.refunded_count != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.Totals", len)?;
        if let Some(v) = self.pending_amount.as_ref() {
            struct_ser.serialize_field("pendingAmount", v)?;
        }
        if self.pending_count != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field(
                "pendingCount",
                ToString::to_string(&self.pending_count).as_str(),
            )?;
        }
        if let Some(v) = self.minted_amount.as_ref() {
            struct_ser.serialize_field("mintedAmount", v)?;
        }
        if self.minted_count != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field(
                "mintedCount",
                ToString::to_string(&self.minted_count).as_str(),
            )?;
        }
        if let Some(v) = self.voided_amount.as_ref() {
            struct_ser.serialize_field("voidedAmount", v)?;
        }
        if self.voided_count != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field(
                "voidedCount",
                ToString::to_string(&self.voided_count).as_str(),
            )?;
        }
        if let Some(v) = self.refund_pending_amount.as_ref() {
            struct_ser.serialize_field("refundPendingAmount", v)?;
        }
        if self.refund_pending_count != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field(
                "refundPendingCount",
                ToString::to_string(&self.refund_pending_count).as_str(),
            )?;
        }
        if let Some(v) = self.refunded_amount.as_ref() {
            struct_ser.serialize_field("refundedAmount", v)?;
        }
        if self.refunded_count != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field(
                "refundedCount",
                ToString::to_string(&self.refunded_count).as_str(),
            )?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for Totals {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "pending_amount",
            "pendingAmount",
            "pending_count",
            "pendingCount",
            "minted_amount",
            "mintedAmount",
            "minted_count",
            "mintedCount",
            "voided_amount",
            "voidedAmount",
            "voided_count",
            "voidedCount",
            "refund_pending_amount",
            "refundPendingAmount",
            "refund_pending_count",
            "refundPendingCount",
            "refunded_amount",
            "refundedAmount",
            "refunded_count",
            "refundedCount",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            PendingAmount,
            PendingCount,
            MintedAmount,
            MintedCount,
            VoidedAmount,
            VoidedCount,
            RefundPendingAmount,
            RefundPendingCount,
            RefundedAmount,
            RefundedCount,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "pendingAmount" | "pending_amount" => Ok(GeneratedField::PendingAmount),
                            "pendingCount" | "pending_count" => Ok(GeneratedField::PendingCount),
                            "mintedAmount" | "minted_amount" => Ok(GeneratedField::MintedAmount),
                            "mintedCount" | "minted_count" => Ok(GeneratedField::MintedCount),
                            "voidedAmount" | "voided_amount" => Ok(GeneratedField::VoidedAmount),
                            "voidedCount" | "voided_count" => Ok(GeneratedField::VoidedCount),
                            "refundPendingAmount" | "refund_pending_amount" => {
                                Ok(GeneratedField::RefundPendingAmount)
                            }
                            "refundPendingCount" | "refund_pending_count" => {
                                Ok(GeneratedField::RefundPendingCount)
                            }
                            "refundedAmount" | "refunded_amount" => {
                                Ok(GeneratedField::RefundedAmount)
                            }
                            "refundedCount" | "refunded_count" => Ok(GeneratedField::RefundedCount),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = Totals;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.Totals")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<Totals, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut pending_amount__ = None;
                let mut pending_count__ = None;
                let mut minted_amount__ = None;
                let mut minted_count__ = None;
                let mut voided_amount__ = None;
                let mut voided_count__ = None;
                let mut refund_pending_amount__ = None;
                let mut refund_pending_count__ = None;
                let mut refunded_amount__ = None;
                let mut refunded_count__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::PendingAmount => {
                            if pending_amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pendingAmount"));
                            }
                            pending_amount__ = map_.next_value()?;
                        }
                        GeneratedField::PendingCount => {
                            if pending_count__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pendingCount"));
                            }
                            pending_count__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                        GeneratedField::MintedAmount => {
                            if minted_amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("mintedAmount"));
                            }
                            minted_amount__ = map_.next_value()?;
                        }
                        GeneratedField::MintedCount => {
                            if minted_count__.is_some() {
                                return Err(serde::de::Error::duplicate_field("mintedCount"));
                            }
                            minted_count__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                        GeneratedField::VoidedAmount => {
                            if voided_amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("voidedAmount"));
                            }
                            voided_amount__ = map_.next_value()?;
                        }
                        GeneratedField::VoidedCount => {
                            if voided_count__.is_some() {
                                return Err(serde::de::Error::duplicate_field("voidedCount"));
                            }
                            voided_count__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                        GeneratedField::RefundPendingAmount => {
                            if refund_pending_amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field(
                                    "refundPendingAmount",
                                ));
                            }
                            refund_pending_amount__ = map_.next_value()?;
                        }
                        GeneratedField::RefundPendingCount => {
                            if refund_pending_count__.is_some() {
                                return Err(serde::de::Error::duplicate_field(
                                    "refundPendingCount",
                                ));
                            }
                            refund_pending_count__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                        GeneratedField::RefundedAmount => {
                            if refunded_amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("refundedAmount"));
                            }
                            refunded_amount__ = map_.next_value()?;
                        }
                        GeneratedField::RefundedCount => {
                            if refunded_count__.is_some() {
                                return Err(serde::de::Error::duplicate_field("refundedCount"));
                            }
                            refunded_count__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                    }
                }
                Ok(Totals {
                    pending_amount: pending_amount__,
                    pending_count: pending_count__.unwrap_or_default(),
                    minted_amount: minted_amount__,
                    minted_count: minted_count__.unwrap_or_default(),
                    voided_amount: voided_amount__,
                    voided_count: voided_count__.unwrap_or_default(),
                    refund_pending_amount: refund_pending_amount__,
                    refund_pending_count: refund_pending_count__.unwrap_or_default(),
                    refunded_amount: refunded_amount__,
                    refunded_count: refunded_count__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.Totals", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for TransferState {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.id.is_some() {
            len += 1;
        }
        if self.data.is_some() {
            len += 1;
        }
        if self.settlement.is_some() {
            len += 1;
        }
        if self.status != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.TransferState", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if let Some(v) = self.data.as_ref() {
            struct_ser.serialize_field("data", v)?;
        }
        if let Some(v) = self.settlement.as_ref() {
            struct_ser.serialize_field("settlement", v)?;
        }
        if self.status != 0 {
            let v = Status::try_from(self.status).map_err(|_| {
                serde::ser::Error::custom(format!("Invalid variant {}", self.status))
            })?;
            struct_ser.serialize_field("status", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for TransferState {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["id", "data", "settlement", "status"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Id,
            Data,
            Settlement,
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

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "id" => Ok(GeneratedField::Id),
                            "data" => Ok(GeneratedField::Data),
                            "settlement" => Ok(GeneratedField::Settlement),
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
            type Value = TransferState;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.TransferState")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<TransferState, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                let mut data__ = None;
                let mut settlement__ = None;
                let mut status__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = map_.next_value()?;
                        }
                        GeneratedField::Data => {
                            if data__.is_some() {
                                return Err(serde::de::Error::duplicate_field("data"));
                            }
                            data__ = map_.next_value()?;
                        }
                        GeneratedField::Settlement => {
                            if settlement__.is_some() {
                                return Err(serde::de::Error::duplicate_field("settlement"));
                            }
                            settlement__ = map_.next_value()?;
                        }
                        GeneratedField::Status => {
                            if status__.is_some() {
                                return Err(serde::de::Error::duplicate_field("status"));
                            }
                            status__ = Some(map_.next_value::<Status>()? as i32);
                        }
                    }
                }
                Ok(TransferState {
                    id: id__,
                    data: data__,
                    settlement: settlement__,
                    status: status__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.TransferState", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for VoidAuthorization {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.requested_height != 0 {
            len += 1;
        }
        if self.requested_round != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.VoidAuthorization", len)?;
        if self.requested_height != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field(
                "requestedHeight",
                ToString::to_string(&self.requested_height).as_str(),
            )?;
        }
        if self.requested_round != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field(
                "requestedRound",
                ToString::to_string(&self.requested_round).as_str(),
            )?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for VoidAuthorization {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "requested_height",
            "requestedHeight",
            "requested_round",
            "requestedRound",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            RequestedHeight,
            RequestedRound,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "requestedHeight" | "requested_height" => {
                                Ok(GeneratedField::RequestedHeight)
                            }
                            "requestedRound" | "requested_round" => {
                                Ok(GeneratedField::RequestedRound)
                            }
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = VoidAuthorization;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.VoidAuthorization")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<VoidAuthorization, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut requested_height__ = None;
                let mut requested_round__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::RequestedHeight => {
                            if requested_height__.is_some() {
                                return Err(serde::de::Error::duplicate_field("requestedHeight"));
                            }
                            requested_height__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                        GeneratedField::RequestedRound => {
                            if requested_round__.is_some() {
                                return Err(serde::de::Error::duplicate_field("requestedRound"));
                            }
                            requested_round__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                    }
                }
                Ok(VoidAuthorization {
                    requested_height: requested_height__.unwrap_or_default(),
                    requested_round: requested_round__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.VoidAuthorization", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for Vote {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.reporter.is_empty() {
            len += 1;
        }
        if self.observation.is_some() {
            len += 1;
        }
        if self.policy_version != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.sink.v1.Vote", len)?;
        if !self.reporter.is_empty() {
            struct_ser.serialize_field("reporter", &self.reporter)?;
        }
        if let Some(v) = self.observation.as_ref() {
            struct_ser.serialize_field("observation", v)?;
        }
        if self.policy_version != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field(
                "policyVersion",
                ToString::to_string(&self.policy_version).as_str(),
            )?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for Vote {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &["reporter", "observation", "policy_version", "policyVersion"];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Reporter,
            Observation,
            PolicyVersion,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(
                        &self,
                        formatter: &mut std::fmt::Formatter<'_>,
                    ) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "reporter" => Ok(GeneratedField::Reporter),
                            "observation" => Ok(GeneratedField::Observation),
                            "policyVersion" | "policy_version" => Ok(GeneratedField::PolicyVersion),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = Vote;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.sink.v1.Vote")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<Vote, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut reporter__ = None;
                let mut observation__ = None;
                let mut policy_version__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Reporter => {
                            if reporter__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reporter"));
                            }
                            reporter__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Observation => {
                            if observation__.is_some() {
                                return Err(serde::de::Error::duplicate_field("observation"));
                            }
                            observation__ = map_.next_value()?;
                        }
                        GeneratedField::PolicyVersion => {
                            if policy_version__.is_some() {
                                return Err(serde::de::Error::duplicate_field("policyVersion"));
                            }
                            policy_version__ = Some(
                                map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?
                                    .0,
                            );
                        }
                    }
                }
                Ok(Vote {
                    reporter: reporter__.unwrap_or_default(),
                    observation: observation__,
                    policy_version: policy_version__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.sink.v1.Vote", FIELDS, GeneratedVisitor)
    }
}

// Generated from rs/serde/sink_legacy_dec.rs by rs/script/fix-prost-serde.py.
mod sink_legacy_dec {
    const PRECISION: usize = 18;
    // LegacyDec bounds the atomic magnitude below 2^256 * 10^18.
    const MAX_ATOMS: &str = concat!(
        "115792089237316195423570985008687907853269984665640564039457584007913129639935",
        "999999999999999999"
    );

    fn sign(value: &str) -> (bool, &str) {
        if let Some(digits) = value.strip_prefix('-') {
            (true, digits)
        } else {
            (false, value.strip_prefix('+').unwrap_or(value))
        }
    }

    fn digits(value: &str) -> bool {
        !value.is_empty() && value.bytes().all(|byte| byte.is_ascii_digit())
    }

    fn magnitude(value: &str) -> Result<&str, &'static str> {
        if !digits(value) {
            return Err("reporters_threshold requires decimal digits");
        }
        let value = value.trim_start_matches('0');
        if value.len() > MAX_ATOMS.len() || (value.len() == MAX_ATOMS.len() && value > MAX_ATOMS) {
            return Err("reporters_threshold exceeds the LegacyDec range");
        }
        Ok(if value.is_empty() { "0" } else { value })
    }

    pub(super) fn from_json(value: &str) -> Result<String, &'static str> {
        let (negative, value) = sign(value);
        let (whole, fraction) = match value.split_once('.') {
            Some((whole, fraction)) if digits(fraction) => (whole, fraction),
            Some(_) => return Err("reporters_threshold requires a decimal fraction"),
            None => (value, ""),
        };
        if !digits(whole) {
            return Err("reporters_threshold requires decimal digits");
        }
        if fraction.len() > PRECISION {
            return Err("reporters_threshold exceeds 18 decimal places");
        }

        let atoms = format!("{whole}{fraction:0<PRECISION$}");
        let atoms = magnitude(&atoms)?;
        let sign = if negative && atoms != "0" { "-" } else { "" };
        Ok(format!("{sign}{atoms}"))
    }

    pub(super) fn to_json(value: &str) -> Result<String, &'static str> {
        let (negative, value) = sign(value);
        if value.len() > 1 && value.starts_with('0') {
            return Err("reporters_threshold requires an unambiguous atomic integer");
        }
        let atoms = magnitude(value)?;
        let sign = if negative && atoms != "0" { "-" } else { "" };
        if atoms.len() <= PRECISION {
            return Ok(format!("{sign}0.{atoms:0>PRECISION$}"));
        }

        let (whole, fraction) = atoms.split_at(atoms.len() - PRECISION);
        Ok(format!("{sign}{whole}.{fraction}"))
    }
}
