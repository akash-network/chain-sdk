// @generated
impl serde::Serialize for BmRecord {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.burned_from.is_empty() {
            len += 1;
        }
        if !self.minted_to.is_empty() {
            len += 1;
        }
        if !self.burner.is_empty() {
            len += 1;
        }
        if !self.minter.is_empty() {
            len += 1;
        }
        if self.burned.is_some() {
            len += 1;
        }
        if self.minted.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.bme.v1.BMRecord", len)?;
        if !self.burned_from.is_empty() {
            struct_ser.serialize_field("burnedFrom", &self.burned_from)?;
        }
        if !self.minted_to.is_empty() {
            struct_ser.serialize_field("mintedTo", &self.minted_to)?;
        }
        if !self.burner.is_empty() {
            struct_ser.serialize_field("burner", &self.burner)?;
        }
        if !self.minter.is_empty() {
            struct_ser.serialize_field("minter", &self.minter)?;
        }
        if let Some(v) = self.burned.as_ref() {
            struct_ser.serialize_field("burned", v)?;
        }
        if let Some(v) = self.minted.as_ref() {
            struct_ser.serialize_field("minted", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for BmRecord {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "burned_from",
            "burnedFrom",
            "minted_to",
            "mintedTo",
            "burner",
            "minter",
            "burned",
            "minted",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            BurnedFrom,
            MintedTo,
            Burner,
            Minter,
            Burned,
            Minted,
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
                            "burnedFrom" | "burned_from" => Ok(GeneratedField::BurnedFrom),
                            "mintedTo" | "minted_to" => Ok(GeneratedField::MintedTo),
                            "burner" => Ok(GeneratedField::Burner),
                            "minter" => Ok(GeneratedField::Minter),
                            "burned" => Ok(GeneratedField::Burned),
                            "minted" => Ok(GeneratedField::Minted),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = BmRecord;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.bme.v1.BMRecord")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<BmRecord, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut burned_from__ = None;
                let mut minted_to__ = None;
                let mut burner__ = None;
                let mut minter__ = None;
                let mut burned__ = None;
                let mut minted__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::BurnedFrom => {
                            if burned_from__.is_some() {
                                return Err(serde::de::Error::duplicate_field("burnedFrom"));
                            }
                            burned_from__ = Some(map_.next_value()?);
                        }
                        GeneratedField::MintedTo => {
                            if minted_to__.is_some() {
                                return Err(serde::de::Error::duplicate_field("mintedTo"));
                            }
                            minted_to__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Burner => {
                            if burner__.is_some() {
                                return Err(serde::de::Error::duplicate_field("burner"));
                            }
                            burner__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Minter => {
                            if minter__.is_some() {
                                return Err(serde::de::Error::duplicate_field("minter"));
                            }
                            minter__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Burned => {
                            if burned__.is_some() {
                                return Err(serde::de::Error::duplicate_field("burned"));
                            }
                            burned__ = map_.next_value()?;
                        }
                        GeneratedField::Minted => {
                            if minted__.is_some() {
                                return Err(serde::de::Error::duplicate_field("minted"));
                            }
                            minted__ = map_.next_value()?;
                        }
                    }
                }
                Ok(BmRecord {
                    burned_from: burned_from__.unwrap_or_default(),
                    minted_to: minted_to__.unwrap_or_default(),
                    burner: burner__.unwrap_or_default(),
                    minter: minter__.unwrap_or_default(),
                    burned: burned__,
                    minted: minted__,
                })
            }
        }
        deserializer.deserialize_struct("akash.bme.v1.BMRecord", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for BurnMintPair {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.burned.is_some() {
            len += 1;
        }
        if self.minted.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.bme.v1.BurnMintPair", len)?;
        if let Some(v) = self.burned.as_ref() {
            struct_ser.serialize_field("burned", v)?;
        }
        if let Some(v) = self.minted.as_ref() {
            struct_ser.serialize_field("minted", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for BurnMintPair {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "burned",
            "minted",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Burned,
            Minted,
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
                            "burned" => Ok(GeneratedField::Burned),
                            "minted" => Ok(GeneratedField::Minted),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = BurnMintPair;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.bme.v1.BurnMintPair")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<BurnMintPair, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut burned__ = None;
                let mut minted__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Burned => {
                            if burned__.is_some() {
                                return Err(serde::de::Error::duplicate_field("burned"));
                            }
                            burned__ = map_.next_value()?;
                        }
                        GeneratedField::Minted => {
                            if minted__.is_some() {
                                return Err(serde::de::Error::duplicate_field("minted"));
                            }
                            minted__ = map_.next_value()?;
                        }
                    }
                }
                Ok(BurnMintPair {
                    burned: burned__,
                    minted: minted__,
                })
            }
        }
        deserializer.deserialize_struct("akash.bme.v1.BurnMintPair", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for CircuitBreakerStatus {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "CIRCUIT_BREAKER_STATUS_UNSPECIFIED",
            Self::Healthy => "CIRCUIT_BREAKER_STATUS_HEALTHY",
            Self::Warning => "CIRCUIT_BREAKER_STATUS_WARNING",
            Self::Halt => "CIRCUIT_BREAKER_STATUS_HALT",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for CircuitBreakerStatus {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "CIRCUIT_BREAKER_STATUS_UNSPECIFIED",
            "CIRCUIT_BREAKER_STATUS_HEALTHY",
            "CIRCUIT_BREAKER_STATUS_WARNING",
            "CIRCUIT_BREAKER_STATUS_HALT",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = CircuitBreakerStatus;

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
                    "CIRCUIT_BREAKER_STATUS_UNSPECIFIED" => Ok(CircuitBreakerStatus::Unspecified),
                    "CIRCUIT_BREAKER_STATUS_HEALTHY" => Ok(CircuitBreakerStatus::Healthy),
                    "CIRCUIT_BREAKER_STATUS_WARNING" => Ok(CircuitBreakerStatus::Warning),
                    "CIRCUIT_BREAKER_STATUS_HALT" => Ok(CircuitBreakerStatus::Halt),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for CoinPrice {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.coin.is_some() {
            len += 1;
        }
        if !self.price.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.bme.v1.CoinPrice", len)?;
        if let Some(v) = self.coin.as_ref() {
            struct_ser.serialize_field("coin", v)?;
        }
        if !self.price.is_empty() {
            struct_ser.serialize_field("price", &self.price)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for CoinPrice {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "coin",
            "price",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Coin,
            Price,
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
                            "coin" => Ok(GeneratedField::Coin),
                            "price" => Ok(GeneratedField::Price),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = CoinPrice;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.bme.v1.CoinPrice")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<CoinPrice, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut coin__ = None;
                let mut price__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Coin => {
                            if coin__.is_some() {
                                return Err(serde::de::Error::duplicate_field("coin"));
                            }
                            coin__ = map_.next_value()?;
                        }
                        GeneratedField::Price => {
                            if price__.is_some() {
                                return Err(serde::de::Error::duplicate_field("price"));
                            }
                            price__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(CoinPrice {
                    coin: coin__,
                    price: price__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.bme.v1.CoinPrice", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for CollateralRatio {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.ratio.is_empty() {
            len += 1;
        }
        if self.status != 0 {
            len += 1;
        }
        if !self.reference_price.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.bme.v1.CollateralRatio", len)?;
        if !self.ratio.is_empty() {
            struct_ser.serialize_field("ratio", &self.ratio)?;
        }
        if self.status != 0 {
            let v = CircuitBreakerStatus::try_from(self.status)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.status)))?;
            struct_ser.serialize_field("status", &v)?;
        }
        if !self.reference_price.is_empty() {
            struct_ser.serialize_field("referencePrice", &self.reference_price)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for CollateralRatio {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "ratio",
            "status",
            "reference_price",
            "referencePrice",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Ratio,
            Status,
            ReferencePrice,
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
                            "ratio" => Ok(GeneratedField::Ratio),
                            "status" => Ok(GeneratedField::Status),
                            "referencePrice" | "reference_price" => Ok(GeneratedField::ReferencePrice),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = CollateralRatio;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.bme.v1.CollateralRatio")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<CollateralRatio, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut ratio__ = None;
                let mut status__ = None;
                let mut reference_price__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Ratio => {
                            if ratio__.is_some() {
                                return Err(serde::de::Error::duplicate_field("ratio"));
                            }
                            ratio__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Status => {
                            if status__.is_some() {
                                return Err(serde::de::Error::duplicate_field("status"));
                            }
                            status__ = Some(map_.next_value::<CircuitBreakerStatus>()? as i32);
                        }
                        GeneratedField::ReferencePrice => {
                            if reference_price__.is_some() {
                                return Err(serde::de::Error::duplicate_field("referencePrice"));
                            }
                            reference_price__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(CollateralRatio {
                    ratio: ratio__.unwrap_or_default(),
                    status: status__.unwrap_or_default(),
                    reference_price: reference_price__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.bme.v1.CollateralRatio", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventBmRecord {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.burned_from.is_empty() {
            len += 1;
        }
        if !self.minted_to.is_empty() {
            len += 1;
        }
        if !self.burner.is_empty() {
            len += 1;
        }
        if !self.minter.is_empty() {
            len += 1;
        }
        if self.burned.is_some() {
            len += 1;
        }
        if self.minted.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.bme.v1.EventBMRecord", len)?;
        if !self.burned_from.is_empty() {
            struct_ser.serialize_field("burnedFrom", &self.burned_from)?;
        }
        if !self.minted_to.is_empty() {
            struct_ser.serialize_field("mintedTo", &self.minted_to)?;
        }
        if !self.burner.is_empty() {
            struct_ser.serialize_field("burner", &self.burner)?;
        }
        if !self.minter.is_empty() {
            struct_ser.serialize_field("minter", &self.minter)?;
        }
        if let Some(v) = self.burned.as_ref() {
            struct_ser.serialize_field("burned", v)?;
        }
        if let Some(v) = self.minted.as_ref() {
            struct_ser.serialize_field("minted", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventBmRecord {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "burned_from",
            "burnedFrom",
            "minted_to",
            "mintedTo",
            "burner",
            "minter",
            "burned",
            "minted",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            BurnedFrom,
            MintedTo,
            Burner,
            Minter,
            Burned,
            Minted,
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
                            "burnedFrom" | "burned_from" => Ok(GeneratedField::BurnedFrom),
                            "mintedTo" | "minted_to" => Ok(GeneratedField::MintedTo),
                            "burner" => Ok(GeneratedField::Burner),
                            "minter" => Ok(GeneratedField::Minter),
                            "burned" => Ok(GeneratedField::Burned),
                            "minted" => Ok(GeneratedField::Minted),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventBmRecord;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.bme.v1.EventBMRecord")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventBmRecord, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut burned_from__ = None;
                let mut minted_to__ = None;
                let mut burner__ = None;
                let mut minter__ = None;
                let mut burned__ = None;
                let mut minted__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::BurnedFrom => {
                            if burned_from__.is_some() {
                                return Err(serde::de::Error::duplicate_field("burnedFrom"));
                            }
                            burned_from__ = Some(map_.next_value()?);
                        }
                        GeneratedField::MintedTo => {
                            if minted_to__.is_some() {
                                return Err(serde::de::Error::duplicate_field("mintedTo"));
                            }
                            minted_to__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Burner => {
                            if burner__.is_some() {
                                return Err(serde::de::Error::duplicate_field("burner"));
                            }
                            burner__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Minter => {
                            if minter__.is_some() {
                                return Err(serde::de::Error::duplicate_field("minter"));
                            }
                            minter__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Burned => {
                            if burned__.is_some() {
                                return Err(serde::de::Error::duplicate_field("burned"));
                            }
                            burned__ = map_.next_value()?;
                        }
                        GeneratedField::Minted => {
                            if minted__.is_some() {
                                return Err(serde::de::Error::duplicate_field("minted"));
                            }
                            minted__ = map_.next_value()?;
                        }
                    }
                }
                Ok(EventBmRecord {
                    burned_from: burned_from__.unwrap_or_default(),
                    minted_to: minted_to__.unwrap_or_default(),
                    burner: burner__.unwrap_or_default(),
                    minter: minter__.unwrap_or_default(),
                    burned: burned__,
                    minted: minted__,
                })
            }
        }
        deserializer.deserialize_struct("akash.bme.v1.EventBMRecord", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventCircuitBreakerStatusChange {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.previous_status != 0 {
            len += 1;
        }
        if self.new_status != 0 {
            len += 1;
        }
        if !self.collateral_ratio.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.bme.v1.EventCircuitBreakerStatusChange", len)?;
        if self.previous_status != 0 {
            let v = CircuitBreakerStatus::try_from(self.previous_status)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.previous_status)))?;
            struct_ser.serialize_field("previousStatus", &v)?;
        }
        if self.new_status != 0 {
            let v = CircuitBreakerStatus::try_from(self.new_status)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.new_status)))?;
            struct_ser.serialize_field("newStatus", &v)?;
        }
        if !self.collateral_ratio.is_empty() {
            struct_ser.serialize_field("collateralRatio", &self.collateral_ratio)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventCircuitBreakerStatusChange {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "previous_status",
            "previousStatus",
            "new_status",
            "newStatus",
            "collateral_ratio",
            "collateralRatio",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            PreviousStatus,
            NewStatus,
            CollateralRatio,
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
                            "previousStatus" | "previous_status" => Ok(GeneratedField::PreviousStatus),
                            "newStatus" | "new_status" => Ok(GeneratedField::NewStatus),
                            "collateralRatio" | "collateral_ratio" => Ok(GeneratedField::CollateralRatio),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventCircuitBreakerStatusChange;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.bme.v1.EventCircuitBreakerStatusChange")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventCircuitBreakerStatusChange, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut previous_status__ = None;
                let mut new_status__ = None;
                let mut collateral_ratio__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::PreviousStatus => {
                            if previous_status__.is_some() {
                                return Err(serde::de::Error::duplicate_field("previousStatus"));
                            }
                            previous_status__ = Some(map_.next_value::<CircuitBreakerStatus>()? as i32);
                        }
                        GeneratedField::NewStatus => {
                            if new_status__.is_some() {
                                return Err(serde::de::Error::duplicate_field("newStatus"));
                            }
                            new_status__ = Some(map_.next_value::<CircuitBreakerStatus>()? as i32);
                        }
                        GeneratedField::CollateralRatio => {
                            if collateral_ratio__.is_some() {
                                return Err(serde::de::Error::duplicate_field("collateralRatio"));
                            }
                            collateral_ratio__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(EventCircuitBreakerStatusChange {
                    previous_status: previous_status__.unwrap_or_default(),
                    new_status: new_status__.unwrap_or_default(),
                    collateral_ratio: collateral_ratio__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.bme.v1.EventCircuitBreakerStatusChange", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventVaultSeeded {
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
        if !self.source.is_empty() {
            len += 1;
        }
        if self.new_vault_balance.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.bme.v1.EventVaultSeeded", len)?;
        if let Some(v) = self.amount.as_ref() {
            struct_ser.serialize_field("amount", v)?;
        }
        if !self.source.is_empty() {
            struct_ser.serialize_field("source", &self.source)?;
        }
        if let Some(v) = self.new_vault_balance.as_ref() {
            struct_ser.serialize_field("newVaultBalance", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventVaultSeeded {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "amount",
            "source",
            "new_vault_balance",
            "newVaultBalance",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Amount,
            Source,
            NewVaultBalance,
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
                            "source" => Ok(GeneratedField::Source),
                            "newVaultBalance" | "new_vault_balance" => Ok(GeneratedField::NewVaultBalance),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventVaultSeeded;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.bme.v1.EventVaultSeeded")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventVaultSeeded, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut amount__ = None;
                let mut source__ = None;
                let mut new_vault_balance__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Amount => {
                            if amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("amount"));
                            }
                            amount__ = map_.next_value()?;
                        }
                        GeneratedField::Source => {
                            if source__.is_some() {
                                return Err(serde::de::Error::duplicate_field("source"));
                            }
                            source__ = Some(map_.next_value()?);
                        }
                        GeneratedField::NewVaultBalance => {
                            if new_vault_balance__.is_some() {
                                return Err(serde::de::Error::duplicate_field("newVaultBalance"));
                            }
                            new_vault_balance__ = map_.next_value()?;
                        }
                    }
                }
                Ok(EventVaultSeeded {
                    amount: amount__,
                    source: source__.unwrap_or_default(),
                    new_vault_balance: new_vault_balance__,
                })
            }
        }
        deserializer.deserialize_struct("akash.bme.v1.EventVaultSeeded", FIELDS, GeneratedVisitor)
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
        if self.state.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.bme.v1.GenesisState", len)?;
        if let Some(v) = self.params.as_ref() {
            struct_ser.serialize_field("params", v)?;
        }
        if let Some(v) = self.state.as_ref() {
            struct_ser.serialize_field("state", v)?;
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
            "state",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Params,
            State,
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
                            "state" => Ok(GeneratedField::State),
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
                formatter.write_str("struct akash.bme.v1.GenesisState")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<GenesisState, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut params__ = None;
                let mut state__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Params => {
                            if params__.is_some() {
                                return Err(serde::de::Error::duplicate_field("params"));
                            }
                            params__ = map_.next_value()?;
                        }
                        GeneratedField::State => {
                            if state__.is_some() {
                                return Err(serde::de::Error::duplicate_field("state"));
                            }
                            state__ = map_.next_value()?;
                        }
                    }
                }
                Ok(GenesisState {
                    params: params__,
                    state: state__,
                })
            }
        }
        deserializer.deserialize_struct("akash.bme.v1.GenesisState", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for LedgerId {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.height != 0 {
            len += 1;
        }
        if self.sequence != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.bme.v1.LedgerID", len)?;
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
impl<'de> serde::Deserialize<'de> for LedgerId {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "height",
            "sequence",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
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

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
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
            type Value = LedgerId;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.bme.v1.LedgerID")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<LedgerId, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut height__ = None;
                let mut sequence__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Height => {
                            if height__.is_some() {
                                return Err(serde::de::Error::duplicate_field("height"));
                            }
                            height__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Sequence => {
                            if sequence__.is_some() {
                                return Err(serde::de::Error::duplicate_field("sequence"));
                            }
                            sequence__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(LedgerId {
                    height: height__.unwrap_or_default(),
                    sequence: sequence__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.bme.v1.LedgerID", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgBurnMint {
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
        if !self.to.is_empty() {
            len += 1;
        }
        if self.coins_to_burn.is_some() {
            len += 1;
        }
        if !self.denom_to_mint.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.bme.v1.MsgBurnMint", len)?;
        if !self.owner.is_empty() {
            struct_ser.serialize_field("owner", &self.owner)?;
        }
        if !self.to.is_empty() {
            struct_ser.serialize_field("to", &self.to)?;
        }
        if let Some(v) = self.coins_to_burn.as_ref() {
            struct_ser.serialize_field("coinsToBurn", v)?;
        }
        if !self.denom_to_mint.is_empty() {
            struct_ser.serialize_field("denomToMint", &self.denom_to_mint)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgBurnMint {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "owner",
            "to",
            "coins_to_burn",
            "coinsToBurn",
            "denom_to_mint",
            "denomToMint",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Owner,
            To,
            CoinsToBurn,
            DenomToMint,
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
                            "owner" => Ok(GeneratedField::Owner),
                            "to" => Ok(GeneratedField::To),
                            "coinsToBurn" | "coins_to_burn" => Ok(GeneratedField::CoinsToBurn),
                            "denomToMint" | "denom_to_mint" => Ok(GeneratedField::DenomToMint),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgBurnMint;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.bme.v1.MsgBurnMint")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgBurnMint, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut owner__ = None;
                let mut to__ = None;
                let mut coins_to_burn__ = None;
                let mut denom_to_mint__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Owner => {
                            if owner__.is_some() {
                                return Err(serde::de::Error::duplicate_field("owner"));
                            }
                            owner__ = Some(map_.next_value()?);
                        }
                        GeneratedField::To => {
                            if to__.is_some() {
                                return Err(serde::de::Error::duplicate_field("to"));
                            }
                            to__ = Some(map_.next_value()?);
                        }
                        GeneratedField::CoinsToBurn => {
                            if coins_to_burn__.is_some() {
                                return Err(serde::de::Error::duplicate_field("coinsToBurn"));
                            }
                            coins_to_burn__ = map_.next_value()?;
                        }
                        GeneratedField::DenomToMint => {
                            if denom_to_mint__.is_some() {
                                return Err(serde::de::Error::duplicate_field("denomToMint"));
                            }
                            denom_to_mint__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(MsgBurnMint {
                    owner: owner__.unwrap_or_default(),
                    to: to__.unwrap_or_default(),
                    coins_to_burn: coins_to_burn__,
                    denom_to_mint: denom_to_mint__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.bme.v1.MsgBurnMint", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgBurnMintResponse {
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
        if !self.to.is_empty() {
            len += 1;
        }
        if self.burned.is_some() {
            len += 1;
        }
        if self.minted.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.bme.v1.MsgBurnMintResponse", len)?;
        if !self.owner.is_empty() {
            struct_ser.serialize_field("owner", &self.owner)?;
        }
        if !self.to.is_empty() {
            struct_ser.serialize_field("to", &self.to)?;
        }
        if let Some(v) = self.burned.as_ref() {
            struct_ser.serialize_field("burned", v)?;
        }
        if let Some(v) = self.minted.as_ref() {
            struct_ser.serialize_field("minted", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgBurnMintResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "owner",
            "to",
            "burned",
            "minted",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Owner,
            To,
            Burned,
            Minted,
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
                            "owner" => Ok(GeneratedField::Owner),
                            "to" => Ok(GeneratedField::To),
                            "burned" => Ok(GeneratedField::Burned),
                            "minted" => Ok(GeneratedField::Minted),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgBurnMintResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.bme.v1.MsgBurnMintResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgBurnMintResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut owner__ = None;
                let mut to__ = None;
                let mut burned__ = None;
                let mut minted__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Owner => {
                            if owner__.is_some() {
                                return Err(serde::de::Error::duplicate_field("owner"));
                            }
                            owner__ = Some(map_.next_value()?);
                        }
                        GeneratedField::To => {
                            if to__.is_some() {
                                return Err(serde::de::Error::duplicate_field("to"));
                            }
                            to__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Burned => {
                            if burned__.is_some() {
                                return Err(serde::de::Error::duplicate_field("burned"));
                            }
                            burned__ = map_.next_value()?;
                        }
                        GeneratedField::Minted => {
                            if minted__.is_some() {
                                return Err(serde::de::Error::duplicate_field("minted"));
                            }
                            minted__ = map_.next_value()?;
                        }
                    }
                }
                Ok(MsgBurnMintResponse {
                    owner: owner__.unwrap_or_default(),
                    to: to__.unwrap_or_default(),
                    burned: burned__,
                    minted: minted__,
                })
            }
        }
        deserializer.deserialize_struct("akash.bme.v1.MsgBurnMintResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgSeedVault {
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
        if self.amount.is_some() {
            len += 1;
        }
        if !self.source.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.bme.v1.MsgSeedVault", len)?;
        if !self.authority.is_empty() {
            struct_ser.serialize_field("authority", &self.authority)?;
        }
        if let Some(v) = self.amount.as_ref() {
            struct_ser.serialize_field("amount", v)?;
        }
        if !self.source.is_empty() {
            struct_ser.serialize_field("source", &self.source)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgSeedVault {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "authority",
            "amount",
            "source",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Authority,
            Amount,
            Source,
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
                            "amount" => Ok(GeneratedField::Amount),
                            "source" => Ok(GeneratedField::Source),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgSeedVault;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.bme.v1.MsgSeedVault")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgSeedVault, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut authority__ = None;
                let mut amount__ = None;
                let mut source__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Authority => {
                            if authority__.is_some() {
                                return Err(serde::de::Error::duplicate_field("authority"));
                            }
                            authority__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Amount => {
                            if amount__.is_some() {
                                return Err(serde::de::Error::duplicate_field("amount"));
                            }
                            amount__ = map_.next_value()?;
                        }
                        GeneratedField::Source => {
                            if source__.is_some() {
                                return Err(serde::de::Error::duplicate_field("source"));
                            }
                            source__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(MsgSeedVault {
                    authority: authority__.unwrap_or_default(),
                    amount: amount__,
                    source: source__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.bme.v1.MsgSeedVault", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgSeedVaultResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.vault_akt.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.bme.v1.MsgSeedVaultResponse", len)?;
        if !self.vault_akt.is_empty() {
            struct_ser.serialize_field("vaultAkt", &self.vault_akt)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgSeedVaultResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "vault_akt",
            "vaultAkt",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            VaultAkt,
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
                            "vaultAkt" | "vault_akt" => Ok(GeneratedField::VaultAkt),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgSeedVaultResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.bme.v1.MsgSeedVaultResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgSeedVaultResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut vault_akt__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::VaultAkt => {
                            if vault_akt__.is_some() {
                                return Err(serde::de::Error::duplicate_field("vaultAkt"));
                            }
                            vault_akt__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(MsgSeedVaultResponse {
                    vault_akt: vault_akt__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.bme.v1.MsgSeedVaultResponse", FIELDS, GeneratedVisitor)
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
        let mut struct_ser = serializer.serialize_struct("akash.bme.v1.MsgUpdateParams", len)?;
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
                formatter.write_str("struct akash.bme.v1.MsgUpdateParams")
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
        deserializer.deserialize_struct("akash.bme.v1.MsgUpdateParams", FIELDS, GeneratedVisitor)
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
        let struct_ser = serializer.serialize_struct("akash.bme.v1.MsgUpdateParamsResponse", len)?;
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
                formatter.write_str("struct akash.bme.v1.MsgUpdateParamsResponse")
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
        deserializer.deserialize_struct("akash.bme.v1.MsgUpdateParamsResponse", FIELDS, GeneratedVisitor)
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
        if self.oracle_twap_window.is_some() {
            len += 1;
        }
        if self.oracle_outlier_threshold_bps != 0 {
            len += 1;
        }
        if self.circuit_breaker_warn_threshold != 0 {
            len += 1;
        }
        if self.circuit_breaker_halt_threshold != 0 {
            len += 1;
        }
        if self.mint_spread_bps != 0 {
            len += 1;
        }
        if self.settle_spread_bps != 0 {
            len += 1;
        }
        if self.settlement_epoch_blocks != 0 {
            len += 1;
        }
        if self.min_runway_blocks != 0 {
            len += 1;
        }
        if self.snapshot_interval_blocks != 0 {
            len += 1;
        }
        if self.snapshot_retention_seconds != 0 {
            len += 1;
        }
        if self.snapshot_bucket_duration != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.bme.v1.Params", len)?;
        if self.enabled {
            struct_ser.serialize_field("enabled", &self.enabled)?;
        }
        if let Some(v) = self.oracle_twap_window.as_ref() {
            struct_ser.serialize_field("oracleTwapWindow", v)?;
        }
        if self.oracle_outlier_threshold_bps != 0 {
            struct_ser.serialize_field("oracleOutlierThresholdBps", &self.oracle_outlier_threshold_bps)?;
        }
        if self.circuit_breaker_warn_threshold != 0 {
            struct_ser.serialize_field("circuitBreakerWarnThreshold", &self.circuit_breaker_warn_threshold)?;
        }
        if self.circuit_breaker_halt_threshold != 0 {
            struct_ser.serialize_field("circuitBreakerHaltThreshold", &self.circuit_breaker_halt_threshold)?;
        }
        if self.mint_spread_bps != 0 {
            struct_ser.serialize_field("mintSpreadBps", &self.mint_spread_bps)?;
        }
        if self.settle_spread_bps != 0 {
            struct_ser.serialize_field("settleSpreadBps", &self.settle_spread_bps)?;
        }
        if self.settlement_epoch_blocks != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("settlementEpochBlocks", ToString::to_string(&self.settlement_epoch_blocks).as_str())?;
        }
        if self.min_runway_blocks != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("minRunwayBlocks", ToString::to_string(&self.min_runway_blocks).as_str())?;
        }
        if self.snapshot_interval_blocks != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("snapshotIntervalBlocks", ToString::to_string(&self.snapshot_interval_blocks).as_str())?;
        }
        if self.snapshot_retention_seconds != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("snapshotRetentionSeconds", ToString::to_string(&self.snapshot_retention_seconds).as_str())?;
        }
        if self.snapshot_bucket_duration != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("snapshotBucketDuration", ToString::to_string(&self.snapshot_bucket_duration).as_str())?;
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
            "oracle_twap_window",
            "oracleTwapWindow",
            "oracle_outlier_threshold_bps",
            "oracleOutlierThresholdBps",
            "circuit_breaker_warn_threshold",
            "circuitBreakerWarnThreshold",
            "circuit_breaker_halt_threshold",
            "circuitBreakerHaltThreshold",
            "mint_spread_bps",
            "mintSpreadBps",
            "settle_spread_bps",
            "settleSpreadBps",
            "settlement_epoch_blocks",
            "settlementEpochBlocks",
            "min_runway_blocks",
            "minRunwayBlocks",
            "snapshot_interval_blocks",
            "snapshotIntervalBlocks",
            "snapshot_retention_seconds",
            "snapshotRetentionSeconds",
            "snapshot_bucket_duration",
            "snapshotBucketDuration",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Enabled,
            OracleTwapWindow,
            OracleOutlierThresholdBps,
            CircuitBreakerWarnThreshold,
            CircuitBreakerHaltThreshold,
            MintSpreadBps,
            SettleSpreadBps,
            SettlementEpochBlocks,
            MinRunwayBlocks,
            SnapshotIntervalBlocks,
            SnapshotRetentionSeconds,
            SnapshotBucketDuration,
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
                            "enabled" => Ok(GeneratedField::Enabled),
                            "oracleTwapWindow" | "oracle_twap_window" => Ok(GeneratedField::OracleTwapWindow),
                            "oracleOutlierThresholdBps" | "oracle_outlier_threshold_bps" => Ok(GeneratedField::OracleOutlierThresholdBps),
                            "circuitBreakerWarnThreshold" | "circuit_breaker_warn_threshold" => Ok(GeneratedField::CircuitBreakerWarnThreshold),
                            "circuitBreakerHaltThreshold" | "circuit_breaker_halt_threshold" => Ok(GeneratedField::CircuitBreakerHaltThreshold),
                            "mintSpreadBps" | "mint_spread_bps" => Ok(GeneratedField::MintSpreadBps),
                            "settleSpreadBps" | "settle_spread_bps" => Ok(GeneratedField::SettleSpreadBps),
                            "settlementEpochBlocks" | "settlement_epoch_blocks" => Ok(GeneratedField::SettlementEpochBlocks),
                            "minRunwayBlocks" | "min_runway_blocks" => Ok(GeneratedField::MinRunwayBlocks),
                            "snapshotIntervalBlocks" | "snapshot_interval_blocks" => Ok(GeneratedField::SnapshotIntervalBlocks),
                            "snapshotRetentionSeconds" | "snapshot_retention_seconds" => Ok(GeneratedField::SnapshotRetentionSeconds),
                            "snapshotBucketDuration" | "snapshot_bucket_duration" => Ok(GeneratedField::SnapshotBucketDuration),
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
                formatter.write_str("struct akash.bme.v1.Params")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<Params, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut enabled__ = None;
                let mut oracle_twap_window__ = None;
                let mut oracle_outlier_threshold_bps__ = None;
                let mut circuit_breaker_warn_threshold__ = None;
                let mut circuit_breaker_halt_threshold__ = None;
                let mut mint_spread_bps__ = None;
                let mut settle_spread_bps__ = None;
                let mut settlement_epoch_blocks__ = None;
                let mut min_runway_blocks__ = None;
                let mut snapshot_interval_blocks__ = None;
                let mut snapshot_retention_seconds__ = None;
                let mut snapshot_bucket_duration__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Enabled => {
                            if enabled__.is_some() {
                                return Err(serde::de::Error::duplicate_field("enabled"));
                            }
                            enabled__ = Some(map_.next_value()?);
                        }
                        GeneratedField::OracleTwapWindow => {
                            if oracle_twap_window__.is_some() {
                                return Err(serde::de::Error::duplicate_field("oracleTwapWindow"));
                            }
                            oracle_twap_window__ = map_.next_value()?;
                        }
                        GeneratedField::OracleOutlierThresholdBps => {
                            if oracle_outlier_threshold_bps__.is_some() {
                                return Err(serde::de::Error::duplicate_field("oracleOutlierThresholdBps"));
                            }
                            oracle_outlier_threshold_bps__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::CircuitBreakerWarnThreshold => {
                            if circuit_breaker_warn_threshold__.is_some() {
                                return Err(serde::de::Error::duplicate_field("circuitBreakerWarnThreshold"));
                            }
                            circuit_breaker_warn_threshold__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::CircuitBreakerHaltThreshold => {
                            if circuit_breaker_halt_threshold__.is_some() {
                                return Err(serde::de::Error::duplicate_field("circuitBreakerHaltThreshold"));
                            }
                            circuit_breaker_halt_threshold__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::MintSpreadBps => {
                            if mint_spread_bps__.is_some() {
                                return Err(serde::de::Error::duplicate_field("mintSpreadBps"));
                            }
                            mint_spread_bps__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::SettleSpreadBps => {
                            if settle_spread_bps__.is_some() {
                                return Err(serde::de::Error::duplicate_field("settleSpreadBps"));
                            }
                            settle_spread_bps__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::SettlementEpochBlocks => {
                            if settlement_epoch_blocks__.is_some() {
                                return Err(serde::de::Error::duplicate_field("settlementEpochBlocks"));
                            }
                            settlement_epoch_blocks__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::MinRunwayBlocks => {
                            if min_runway_blocks__.is_some() {
                                return Err(serde::de::Error::duplicate_field("minRunwayBlocks"));
                            }
                            min_runway_blocks__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::SnapshotIntervalBlocks => {
                            if snapshot_interval_blocks__.is_some() {
                                return Err(serde::de::Error::duplicate_field("snapshotIntervalBlocks"));
                            }
                            snapshot_interval_blocks__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::SnapshotRetentionSeconds => {
                            if snapshot_retention_seconds__.is_some() {
                                return Err(serde::de::Error::duplicate_field("snapshotRetentionSeconds"));
                            }
                            snapshot_retention_seconds__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::SnapshotBucketDuration => {
                            if snapshot_bucket_duration__.is_some() {
                                return Err(serde::de::Error::duplicate_field("snapshotBucketDuration"));
                            }
                            snapshot_bucket_duration__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(Params {
                    enabled: enabled__.unwrap_or_default(),
                    oracle_twap_window: oracle_twap_window__,
                    oracle_outlier_threshold_bps: oracle_outlier_threshold_bps__.unwrap_or_default(),
                    circuit_breaker_warn_threshold: circuit_breaker_warn_threshold__.unwrap_or_default(),
                    circuit_breaker_halt_threshold: circuit_breaker_halt_threshold__.unwrap_or_default(),
                    mint_spread_bps: mint_spread_bps__.unwrap_or_default(),
                    settle_spread_bps: settle_spread_bps__.unwrap_or_default(),
                    settlement_epoch_blocks: settlement_epoch_blocks__.unwrap_or_default(),
                    min_runway_blocks: min_runway_blocks__.unwrap_or_default(),
                    snapshot_interval_blocks: snapshot_interval_blocks__.unwrap_or_default(),
                    snapshot_retention_seconds: snapshot_retention_seconds__.unwrap_or_default(),
                    snapshot_bucket_duration: snapshot_bucket_duration__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.bme.v1.Params", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryCircuitBreakerStatusRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.bme.v1.QueryCircuitBreakerStatusRequest", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryCircuitBreakerStatusRequest {
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
            type Value = QueryCircuitBreakerStatusRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.bme.v1.QueryCircuitBreakerStatusRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryCircuitBreakerStatusRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(QueryCircuitBreakerStatusRequest {
                })
            }
        }
        deserializer.deserialize_struct("akash.bme.v1.QueryCircuitBreakerStatusRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryCircuitBreakerStatusResponse {
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
        if !self.collateral_ratio.is_empty() {
            len += 1;
        }
        if !self.warn_threshold.is_empty() {
            len += 1;
        }
        if !self.halt_threshold.is_empty() {
            len += 1;
        }
        if self.mints_allowed {
            len += 1;
        }
        if self.settlements_allowed {
            len += 1;
        }
        if self.refunds_allowed {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.bme.v1.QueryCircuitBreakerStatusResponse", len)?;
        if self.status != 0 {
            let v = CircuitBreakerStatus::try_from(self.status)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.status)))?;
            struct_ser.serialize_field("status", &v)?;
        }
        if !self.collateral_ratio.is_empty() {
            struct_ser.serialize_field("collateralRatio", &self.collateral_ratio)?;
        }
        if !self.warn_threshold.is_empty() {
            struct_ser.serialize_field("warnThreshold", &self.warn_threshold)?;
        }
        if !self.halt_threshold.is_empty() {
            struct_ser.serialize_field("haltThreshold", &self.halt_threshold)?;
        }
        if self.mints_allowed {
            struct_ser.serialize_field("mintsAllowed", &self.mints_allowed)?;
        }
        if self.settlements_allowed {
            struct_ser.serialize_field("settlementsAllowed", &self.settlements_allowed)?;
        }
        if self.refunds_allowed {
            struct_ser.serialize_field("refundsAllowed", &self.refunds_allowed)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryCircuitBreakerStatusResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "status",
            "collateral_ratio",
            "collateralRatio",
            "warn_threshold",
            "warnThreshold",
            "halt_threshold",
            "haltThreshold",
            "mints_allowed",
            "mintsAllowed",
            "settlements_allowed",
            "settlementsAllowed",
            "refunds_allowed",
            "refundsAllowed",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Status,
            CollateralRatio,
            WarnThreshold,
            HaltThreshold,
            MintsAllowed,
            SettlementsAllowed,
            RefundsAllowed,
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
                            "status" => Ok(GeneratedField::Status),
                            "collateralRatio" | "collateral_ratio" => Ok(GeneratedField::CollateralRatio),
                            "warnThreshold" | "warn_threshold" => Ok(GeneratedField::WarnThreshold),
                            "haltThreshold" | "halt_threshold" => Ok(GeneratedField::HaltThreshold),
                            "mintsAllowed" | "mints_allowed" => Ok(GeneratedField::MintsAllowed),
                            "settlementsAllowed" | "settlements_allowed" => Ok(GeneratedField::SettlementsAllowed),
                            "refundsAllowed" | "refunds_allowed" => Ok(GeneratedField::RefundsAllowed),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryCircuitBreakerStatusResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.bme.v1.QueryCircuitBreakerStatusResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryCircuitBreakerStatusResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut status__ = None;
                let mut collateral_ratio__ = None;
                let mut warn_threshold__ = None;
                let mut halt_threshold__ = None;
                let mut mints_allowed__ = None;
                let mut settlements_allowed__ = None;
                let mut refunds_allowed__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Status => {
                            if status__.is_some() {
                                return Err(serde::de::Error::duplicate_field("status"));
                            }
                            status__ = Some(map_.next_value::<CircuitBreakerStatus>()? as i32);
                        }
                        GeneratedField::CollateralRatio => {
                            if collateral_ratio__.is_some() {
                                return Err(serde::de::Error::duplicate_field("collateralRatio"));
                            }
                            collateral_ratio__ = Some(map_.next_value()?);
                        }
                        GeneratedField::WarnThreshold => {
                            if warn_threshold__.is_some() {
                                return Err(serde::de::Error::duplicate_field("warnThreshold"));
                            }
                            warn_threshold__ = Some(map_.next_value()?);
                        }
                        GeneratedField::HaltThreshold => {
                            if halt_threshold__.is_some() {
                                return Err(serde::de::Error::duplicate_field("haltThreshold"));
                            }
                            halt_threshold__ = Some(map_.next_value()?);
                        }
                        GeneratedField::MintsAllowed => {
                            if mints_allowed__.is_some() {
                                return Err(serde::de::Error::duplicate_field("mintsAllowed"));
                            }
                            mints_allowed__ = Some(map_.next_value()?);
                        }
                        GeneratedField::SettlementsAllowed => {
                            if settlements_allowed__.is_some() {
                                return Err(serde::de::Error::duplicate_field("settlementsAllowed"));
                            }
                            settlements_allowed__ = Some(map_.next_value()?);
                        }
                        GeneratedField::RefundsAllowed => {
                            if refunds_allowed__.is_some() {
                                return Err(serde::de::Error::duplicate_field("refundsAllowed"));
                            }
                            refunds_allowed__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(QueryCircuitBreakerStatusResponse {
                    status: status__.unwrap_or_default(),
                    collateral_ratio: collateral_ratio__.unwrap_or_default(),
                    warn_threshold: warn_threshold__.unwrap_or_default(),
                    halt_threshold: halt_threshold__.unwrap_or_default(),
                    mints_allowed: mints_allowed__.unwrap_or_default(),
                    settlements_allowed: settlements_allowed__.unwrap_or_default(),
                    refunds_allowed: refunds_allowed__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.bme.v1.QueryCircuitBreakerStatusResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryCollateralRatioRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.bme.v1.QueryCollateralRatioRequest", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryCollateralRatioRequest {
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
            type Value = QueryCollateralRatioRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.bme.v1.QueryCollateralRatioRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryCollateralRatioRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(QueryCollateralRatioRequest {
                })
            }
        }
        deserializer.deserialize_struct("akash.bme.v1.QueryCollateralRatioRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryCollateralRatioResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.collateral_ratio.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.bme.v1.QueryCollateralRatioResponse", len)?;
        if let Some(v) = self.collateral_ratio.as_ref() {
            struct_ser.serialize_field("collateralRatio", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryCollateralRatioResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "collateral_ratio",
            "collateralRatio",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            CollateralRatio,
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
                            "collateralRatio" | "collateral_ratio" => Ok(GeneratedField::CollateralRatio),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryCollateralRatioResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.bme.v1.QueryCollateralRatioResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryCollateralRatioResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut collateral_ratio__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::CollateralRatio => {
                            if collateral_ratio__.is_some() {
                                return Err(serde::de::Error::duplicate_field("collateralRatio"));
                            }
                            collateral_ratio__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryCollateralRatioResponse {
                    collateral_ratio: collateral_ratio__,
                })
            }
        }
        deserializer.deserialize_struct("akash.bme.v1.QueryCollateralRatioResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryNetBurnRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.period.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.bme.v1.QueryNetBurnRequest", len)?;
        if !self.period.is_empty() {
            struct_ser.serialize_field("period", &self.period)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryNetBurnRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "period",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Period,
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
                            "period" => Ok(GeneratedField::Period),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryNetBurnRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.bme.v1.QueryNetBurnRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryNetBurnRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut period__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Period => {
                            if period__.is_some() {
                                return Err(serde::de::Error::duplicate_field("period"));
                            }
                            period__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(QueryNetBurnRequest {
                    period: period__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.bme.v1.QueryNetBurnRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryNetBurnResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.net_burn.is_some() {
            len += 1;
        }
        if self.total_burned.is_some() {
            len += 1;
        }
        if self.total_minted.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.bme.v1.QueryNetBurnResponse", len)?;
        if let Some(v) = self.net_burn.as_ref() {
            struct_ser.serialize_field("netBurn", v)?;
        }
        if let Some(v) = self.total_burned.as_ref() {
            struct_ser.serialize_field("totalBurned", v)?;
        }
        if let Some(v) = self.total_minted.as_ref() {
            struct_ser.serialize_field("totalMinted", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryNetBurnResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "net_burn",
            "netBurn",
            "total_burned",
            "totalBurned",
            "total_minted",
            "totalMinted",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            NetBurn,
            TotalBurned,
            TotalMinted,
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
                            "netBurn" | "net_burn" => Ok(GeneratedField::NetBurn),
                            "totalBurned" | "total_burned" => Ok(GeneratedField::TotalBurned),
                            "totalMinted" | "total_minted" => Ok(GeneratedField::TotalMinted),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryNetBurnResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.bme.v1.QueryNetBurnResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryNetBurnResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut net_burn__ = None;
                let mut total_burned__ = None;
                let mut total_minted__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::NetBurn => {
                            if net_burn__.is_some() {
                                return Err(serde::de::Error::duplicate_field("netBurn"));
                            }
                            net_burn__ = map_.next_value()?;
                        }
                        GeneratedField::TotalBurned => {
                            if total_burned__.is_some() {
                                return Err(serde::de::Error::duplicate_field("totalBurned"));
                            }
                            total_burned__ = map_.next_value()?;
                        }
                        GeneratedField::TotalMinted => {
                            if total_minted__.is_some() {
                                return Err(serde::de::Error::duplicate_field("totalMinted"));
                            }
                            total_minted__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryNetBurnResponse {
                    net_burn: net_burn__,
                    total_burned: total_burned__,
                    total_minted: total_minted__,
                })
            }
        }
        deserializer.deserialize_struct("akash.bme.v1.QueryNetBurnResponse", FIELDS, GeneratedVisitor)
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
        let struct_ser = serializer.serialize_struct("akash.bme.v1.QueryParamsRequest", len)?;
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
                formatter.write_str("struct akash.bme.v1.QueryParamsRequest")
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
        deserializer.deserialize_struct("akash.bme.v1.QueryParamsRequest", FIELDS, GeneratedVisitor)
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
        let mut struct_ser = serializer.serialize_struct("akash.bme.v1.QueryParamsResponse", len)?;
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
                formatter.write_str("struct akash.bme.v1.QueryParamsResponse")
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
        deserializer.deserialize_struct("akash.bme.v1.QueryParamsResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryVaultStateRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.bme.v1.QueryVaultStateRequest", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryVaultStateRequest {
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
            type Value = QueryVaultStateRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.bme.v1.QueryVaultStateRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryVaultStateRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(QueryVaultStateRequest {
                })
            }
        }
        deserializer.deserialize_struct("akash.bme.v1.QueryVaultStateRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryVaultStateResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.vault_state.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.bme.v1.QueryVaultStateResponse", len)?;
        if let Some(v) = self.vault_state.as_ref() {
            struct_ser.serialize_field("vaultState", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryVaultStateResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "vault_state",
            "vaultState",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            VaultState,
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
                            "vaultState" | "vault_state" => Ok(GeneratedField::VaultState),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryVaultStateResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.bme.v1.QueryVaultStateResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryVaultStateResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut vault_state__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::VaultState => {
                            if vault_state__.is_some() {
                                return Err(serde::de::Error::duplicate_field("vaultState"));
                            }
                            vault_state__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryVaultStateResponse {
                    vault_state: vault_state__,
                })
            }
        }
        deserializer.deserialize_struct("akash.bme.v1.QueryVaultStateResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for State {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.last_updated != 0 {
            len += 1;
        }
        if !self.burned.is_empty() {
            len += 1;
        }
        if !self.minted.is_empty() {
            len += 1;
        }
        if !self.balances.is_empty() {
            len += 1;
        }
        if !self.remint_credits.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.bme.v1.State", len)?;
        if self.last_updated != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("lastUpdated", ToString::to_string(&self.last_updated).as_str())?;
        }
        if !self.burned.is_empty() {
            struct_ser.serialize_field("burned", &self.burned)?;
        }
        if !self.minted.is_empty() {
            struct_ser.serialize_field("minted", &self.minted)?;
        }
        if !self.balances.is_empty() {
            struct_ser.serialize_field("balances", &self.balances)?;
        }
        if !self.remint_credits.is_empty() {
            struct_ser.serialize_field("remintCredits", &self.remint_credits)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for State {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "last_updated",
            "lastUpdated",
            "burned",
            "minted",
            "balances",
            "remint_credits",
            "remintCredits",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            LastUpdated,
            Burned,
            Minted,
            Balances,
            RemintCredits,
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
                            "lastUpdated" | "last_updated" => Ok(GeneratedField::LastUpdated),
                            "burned" => Ok(GeneratedField::Burned),
                            "minted" => Ok(GeneratedField::Minted),
                            "balances" => Ok(GeneratedField::Balances),
                            "remintCredits" | "remint_credits" => Ok(GeneratedField::RemintCredits),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = State;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.bme.v1.State")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<State, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut last_updated__ = None;
                let mut burned__ = None;
                let mut minted__ = None;
                let mut balances__ = None;
                let mut remint_credits__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::LastUpdated => {
                            if last_updated__.is_some() {
                                return Err(serde::de::Error::duplicate_field("lastUpdated"));
                            }
                            last_updated__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Burned => {
                            if burned__.is_some() {
                                return Err(serde::de::Error::duplicate_field("burned"));
                            }
                            burned__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Minted => {
                            if minted__.is_some() {
                                return Err(serde::de::Error::duplicate_field("minted"));
                            }
                            minted__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Balances => {
                            if balances__.is_some() {
                                return Err(serde::de::Error::duplicate_field("balances"));
                            }
                            balances__ = Some(map_.next_value()?);
                        }
                        GeneratedField::RemintCredits => {
                            if remint_credits__.is_some() {
                                return Err(serde::de::Error::duplicate_field("remintCredits"));
                            }
                            remint_credits__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(State {
                    last_updated: last_updated__.unwrap_or_default(),
                    burned: burned__.unwrap_or_default(),
                    minted: minted__.unwrap_or_default(),
                    balances: balances__.unwrap_or_default(),
                    remint_credits: remint_credits__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.bme.v1.State", FIELDS, GeneratedVisitor)
    }
}
