// @generated
impl serde::Serialize for Bid {
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
        if self.state != 0 {
            len += 1;
        }
        if !self.prices.is_empty() {
            len += 1;
        }
        if self.created_at != 0 {
            len += 1;
        }
        if !self.resources_offer.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.Bid", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if self.state != 0 {
            let v = bid::State::try_from(self.state)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.state)))?;
            struct_ser.serialize_field("state", &v)?;
        }
        if !self.prices.is_empty() {
            struct_ser.serialize_field("prices", &self.prices)?;
        }
        if self.created_at != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("createdAt", ToString::to_string(&self.created_at).as_str())?;
        }
        if !self.resources_offer.is_empty() {
            struct_ser.serialize_field("resourcesOffer", &self.resources_offer)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for Bid {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "id",
            "state",
            "prices",
            "created_at",
            "createdAt",
            "resources_offer",
            "resourcesOffer",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Id,
            State,
            Prices,
            CreatedAt,
            ResourcesOffer,
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
                            "state" => Ok(GeneratedField::State),
                            "prices" => Ok(GeneratedField::Prices),
                            "createdAt" | "created_at" => Ok(GeneratedField::CreatedAt),
                            "resourcesOffer" | "resources_offer" => Ok(GeneratedField::ResourcesOffer),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = Bid;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.Bid")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<Bid, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                let mut state__ = None;
                let mut prices__ = None;
                let mut created_at__ = None;
                let mut resources_offer__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = map_.next_value()?;
                        }
                        GeneratedField::State => {
                            if state__.is_some() {
                                return Err(serde::de::Error::duplicate_field("state"));
                            }
                            state__ = Some(map_.next_value::<bid::State>()? as i32);
                        }
                        GeneratedField::Prices => {
                            if prices__.is_some() {
                                return Err(serde::de::Error::duplicate_field("prices"));
                            }
                            prices__ = Some(map_.next_value()?);
                        }
                        GeneratedField::CreatedAt => {
                            if created_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("createdAt"));
                            }
                            created_at__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::ResourcesOffer => {
                            if resources_offer__.is_some() {
                                return Err(serde::de::Error::duplicate_field("resourcesOffer"));
                            }
                            resources_offer__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(Bid {
                    id: id__,
                    state: state__.unwrap_or_default(),
                    prices: prices__.unwrap_or_default(),
                    created_at: created_at__.unwrap_or_default(),
                    resources_offer: resources_offer__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.Bid", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for bid::State {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Invalid => "invalid",
            Self::Open => "open",
            Self::Active => "active",
            Self::Lost => "lost",
            Self::Closed => "closed",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for bid::State {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "invalid",
            "open",
            "active",
            "lost",
            "closed",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = bid::State;

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
                    "invalid" => Ok(bid::State::Invalid),
                    "open" => Ok(bid::State::Open),
                    "active" => Ok(bid::State::Active),
                    "lost" => Ok(bid::State::Lost),
                    "closed" => Ok(bid::State::Closed),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for BidFilters {
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
        if self.dseq != 0 {
            len += 1;
        }
        if self.gseq != 0 {
            len += 1;
        }
        if self.oseq != 0 {
            len += 1;
        }
        if !self.provider.is_empty() {
            len += 1;
        }
        if !self.state.is_empty() {
            len += 1;
        }
        if self.bseq != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.BidFilters", len)?;
        if !self.owner.is_empty() {
            struct_ser.serialize_field("owner", &self.owner)?;
        }
        if self.dseq != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("dseq", ToString::to_string(&self.dseq).as_str())?;
        }
        if self.gseq != 0 {
            struct_ser.serialize_field("gseq", &self.gseq)?;
        }
        if self.oseq != 0 {
            struct_ser.serialize_field("oseq", &self.oseq)?;
        }
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if !self.state.is_empty() {
            struct_ser.serialize_field("state", &self.state)?;
        }
        if self.bseq != 0 {
            struct_ser.serialize_field("bseq", &self.bseq)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for BidFilters {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "owner",
            "dseq",
            "gseq",
            "oseq",
            "provider",
            "state",
            "bseq",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Owner,
            Dseq,
            Gseq,
            Oseq,
            Provider,
            State,
            Bseq,
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
                            "dseq" => Ok(GeneratedField::Dseq),
                            "gseq" => Ok(GeneratedField::Gseq),
                            "oseq" => Ok(GeneratedField::Oseq),
                            "provider" => Ok(GeneratedField::Provider),
                            "state" => Ok(GeneratedField::State),
                            "bseq" => Ok(GeneratedField::Bseq),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = BidFilters;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.BidFilters")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<BidFilters, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut owner__ = None;
                let mut dseq__ = None;
                let mut gseq__ = None;
                let mut oseq__ = None;
                let mut provider__ = None;
                let mut state__ = None;
                let mut bseq__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Owner => {
                            if owner__.is_some() {
                                return Err(serde::de::Error::duplicate_field("owner"));
                            }
                            owner__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Dseq => {
                            if dseq__.is_some() {
                                return Err(serde::de::Error::duplicate_field("dseq"));
                            }
                            dseq__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Gseq => {
                            if gseq__.is_some() {
                                return Err(serde::de::Error::duplicate_field("gseq"));
                            }
                            gseq__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Oseq => {
                            if oseq__.is_some() {
                                return Err(serde::de::Error::duplicate_field("oseq"));
                            }
                            oseq__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::State => {
                            if state__.is_some() {
                                return Err(serde::de::Error::duplicate_field("state"));
                            }
                            state__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Bseq => {
                            if bseq__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bseq"));
                            }
                            bseq__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(BidFilters {
                    owner: owner__.unwrap_or_default(),
                    dseq: dseq__.unwrap_or_default(),
                    gseq: gseq__.unwrap_or_default(),
                    oseq: oseq__.unwrap_or_default(),
                    provider: provider__.unwrap_or_default(),
                    state: state__.unwrap_or_default(),
                    bseq: bseq__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.BidFilters", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for BidId {
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
        if self.dseq != 0 {
            len += 1;
        }
        if self.gseq != 0 {
            len += 1;
        }
        if self.oseq != 0 {
            len += 1;
        }
        if !self.provider.is_empty() {
            len += 1;
        }
        if self.bseq != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.BidID", len)?;
        if !self.owner.is_empty() {
            struct_ser.serialize_field("owner", &self.owner)?;
        }
        if self.dseq != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("dseq", ToString::to_string(&self.dseq).as_str())?;
        }
        if self.gseq != 0 {
            struct_ser.serialize_field("gseq", &self.gseq)?;
        }
        if self.oseq != 0 {
            struct_ser.serialize_field("oseq", &self.oseq)?;
        }
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if self.bseq != 0 {
            struct_ser.serialize_field("bseq", &self.bseq)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for BidId {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "owner",
            "dseq",
            "gseq",
            "oseq",
            "provider",
            "bseq",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Owner,
            Dseq,
            Gseq,
            Oseq,
            Provider,
            Bseq,
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
                            "dseq" => Ok(GeneratedField::Dseq),
                            "gseq" => Ok(GeneratedField::Gseq),
                            "oseq" => Ok(GeneratedField::Oseq),
                            "provider" => Ok(GeneratedField::Provider),
                            "bseq" => Ok(GeneratedField::Bseq),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = BidId;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.BidID")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<BidId, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut owner__ = None;
                let mut dseq__ = None;
                let mut gseq__ = None;
                let mut oseq__ = None;
                let mut provider__ = None;
                let mut bseq__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Owner => {
                            if owner__.is_some() {
                                return Err(serde::de::Error::duplicate_field("owner"));
                            }
                            owner__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Dseq => {
                            if dseq__.is_some() {
                                return Err(serde::de::Error::duplicate_field("dseq"));
                            }
                            dseq__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Gseq => {
                            if gseq__.is_some() {
                                return Err(serde::de::Error::duplicate_field("gseq"));
                            }
                            gseq__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Oseq => {
                            if oseq__.is_some() {
                                return Err(serde::de::Error::duplicate_field("oseq"));
                            }
                            oseq__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Bseq => {
                            if bseq__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bseq"));
                            }
                            bseq__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(BidId {
                    owner: owner__.unwrap_or_default(),
                    dseq: dseq__.unwrap_or_default(),
                    gseq: gseq__.unwrap_or_default(),
                    oseq: oseq__.unwrap_or_default(),
                    provider: provider__.unwrap_or_default(),
                    bseq: bseq__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.BidID", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventBidClosed {
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
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.EventBidClosed", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventBidClosed {
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
            type Value = EventBidClosed;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.EventBidClosed")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventBidClosed, V::Error>
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
                Ok(EventBidClosed {
                    id: id__,
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.EventBidClosed", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventBidCreated {
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
        if !self.prices.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.EventBidCreated", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if !self.prices.is_empty() {
            struct_ser.serialize_field("prices", &self.prices)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventBidCreated {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "id",
            "prices",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Id,
            Prices,
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
                            "prices" => Ok(GeneratedField::Prices),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventBidCreated;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.EventBidCreated")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventBidCreated, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                let mut prices__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = map_.next_value()?;
                        }
                        GeneratedField::Prices => {
                            if prices__.is_some() {
                                return Err(serde::de::Error::duplicate_field("prices"));
                            }
                            prices__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(EventBidCreated {
                    id: id__,
                    prices: prices__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.EventBidCreated", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventLeaseClosed {
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
        if self.reason != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.EventLeaseClosed", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if self.reason != 0 {
            let v = LeaseClosedReason::try_from(self.reason)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.reason)))?;
            struct_ser.serialize_field("reason", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventLeaseClosed {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "id",
            "reason",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
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
            type Value = EventLeaseClosed;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.EventLeaseClosed")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventLeaseClosed, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                let mut reason__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
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
                            reason__ = Some(map_.next_value::<LeaseClosedReason>()? as i32);
                        }
                    }
                }
                Ok(EventLeaseClosed {
                    id: id__,
                    reason: reason__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.EventLeaseClosed", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventLeaseCreated {
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
        if !self.prices.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.EventLeaseCreated", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if !self.prices.is_empty() {
            struct_ser.serialize_field("prices", &self.prices)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventLeaseCreated {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "id",
            "prices",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Id,
            Prices,
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
                            "prices" => Ok(GeneratedField::Prices),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventLeaseCreated;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.EventLeaseCreated")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventLeaseCreated, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                let mut prices__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = map_.next_value()?;
                        }
                        GeneratedField::Prices => {
                            if prices__.is_some() {
                                return Err(serde::de::Error::duplicate_field("prices"));
                            }
                            prices__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(EventLeaseCreated {
                    id: id__,
                    prices: prices__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.EventLeaseCreated", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventOrderClosed {
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
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.EventOrderClosed", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventOrderClosed {
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
            type Value = EventOrderClosed;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.EventOrderClosed")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventOrderClosed, V::Error>
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
                Ok(EventOrderClosed {
                    id: id__,
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.EventOrderClosed", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventOrderCreated {
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
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.EventOrderCreated", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventOrderCreated {
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
            type Value = EventOrderCreated;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.EventOrderCreated")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventOrderCreated, V::Error>
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
                Ok(EventOrderCreated {
                    id: id__,
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.EventOrderCreated", FIELDS, GeneratedVisitor)
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
        if !self.orders.is_empty() {
            len += 1;
        }
        if !self.leases.is_empty() {
            len += 1;
        }
        if !self.bids.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.GenesisState", len)?;
        if let Some(v) = self.params.as_ref() {
            struct_ser.serialize_field("params", v)?;
        }
        if !self.orders.is_empty() {
            struct_ser.serialize_field("orders", &self.orders)?;
        }
        if !self.leases.is_empty() {
            struct_ser.serialize_field("leases", &self.leases)?;
        }
        if !self.bids.is_empty() {
            struct_ser.serialize_field("bids", &self.bids)?;
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
            "orders",
            "leases",
            "bids",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Params,
            Orders,
            Leases,
            Bids,
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
                            "orders" => Ok(GeneratedField::Orders),
                            "leases" => Ok(GeneratedField::Leases),
                            "bids" => Ok(GeneratedField::Bids),
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
                formatter.write_str("struct akash.market.v2beta1.GenesisState")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<GenesisState, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut params__ = None;
                let mut orders__ = None;
                let mut leases__ = None;
                let mut bids__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Params => {
                            if params__.is_some() {
                                return Err(serde::de::Error::duplicate_field("params"));
                            }
                            params__ = map_.next_value()?;
                        }
                        GeneratedField::Orders => {
                            if orders__.is_some() {
                                return Err(serde::de::Error::duplicate_field("orders"));
                            }
                            orders__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Leases => {
                            if leases__.is_some() {
                                return Err(serde::de::Error::duplicate_field("leases"));
                            }
                            leases__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Bids => {
                            if bids__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bids"));
                            }
                            bids__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(GenesisState {
                    params: params__,
                    orders: orders__.unwrap_or_default(),
                    leases: leases__.unwrap_or_default(),
                    bids: bids__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.GenesisState", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for Lease {
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
        if self.state != 0 {
            len += 1;
        }
        if !self.prices.is_empty() {
            len += 1;
        }
        if self.created_at != 0 {
            len += 1;
        }
        if self.closed_on != 0 {
            len += 1;
        }
        if self.reason != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.Lease", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if self.state != 0 {
            let v = lease::State::try_from(self.state)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.state)))?;
            struct_ser.serialize_field("state", &v)?;
        }
        if !self.prices.is_empty() {
            struct_ser.serialize_field("prices", &self.prices)?;
        }
        if self.created_at != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("createdAt", ToString::to_string(&self.created_at).as_str())?;
        }
        if self.closed_on != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("closedOn", ToString::to_string(&self.closed_on).as_str())?;
        }
        if self.reason != 0 {
            let v = LeaseClosedReason::try_from(self.reason)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.reason)))?;
            struct_ser.serialize_field("reason", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for Lease {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "id",
            "state",
            "prices",
            "created_at",
            "createdAt",
            "closed_on",
            "closedOn",
            "reason",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Id,
            State,
            Prices,
            CreatedAt,
            ClosedOn,
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
                            "id" => Ok(GeneratedField::Id),
                            "state" => Ok(GeneratedField::State),
                            "prices" => Ok(GeneratedField::Prices),
                            "createdAt" | "created_at" => Ok(GeneratedField::CreatedAt),
                            "closedOn" | "closed_on" => Ok(GeneratedField::ClosedOn),
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
            type Value = Lease;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.Lease")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<Lease, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                let mut state__ = None;
                let mut prices__ = None;
                let mut created_at__ = None;
                let mut closed_on__ = None;
                let mut reason__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = map_.next_value()?;
                        }
                        GeneratedField::State => {
                            if state__.is_some() {
                                return Err(serde::de::Error::duplicate_field("state"));
                            }
                            state__ = Some(map_.next_value::<lease::State>()? as i32);
                        }
                        GeneratedField::Prices => {
                            if prices__.is_some() {
                                return Err(serde::de::Error::duplicate_field("prices"));
                            }
                            prices__ = Some(map_.next_value()?);
                        }
                        GeneratedField::CreatedAt => {
                            if created_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("createdAt"));
                            }
                            created_at__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::ClosedOn => {
                            if closed_on__.is_some() {
                                return Err(serde::de::Error::duplicate_field("closedOn"));
                            }
                            closed_on__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Reason => {
                            if reason__.is_some() {
                                return Err(serde::de::Error::duplicate_field("reason"));
                            }
                            reason__ = Some(map_.next_value::<LeaseClosedReason>()? as i32);
                        }
                    }
                }
                Ok(Lease {
                    id: id__,
                    state: state__.unwrap_or_default(),
                    prices: prices__.unwrap_or_default(),
                    created_at: created_at__.unwrap_or_default(),
                    closed_on: closed_on__.unwrap_or_default(),
                    reason: reason__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.Lease", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for lease::State {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Invalid => "invalid",
            Self::Active => "active",
            Self::InsufficientFunds => "insufficient_funds",
            Self::Closed => "closed",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for lease::State {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "invalid",
            "active",
            "insufficient_funds",
            "closed",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = lease::State;

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
                    "invalid" => Ok(lease::State::Invalid),
                    "active" => Ok(lease::State::Active),
                    "insufficient_funds" => Ok(lease::State::InsufficientFunds),
                    "closed" => Ok(lease::State::Closed),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for LeaseClosedReason {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::LeaseClosedInvalid => "lease_closed_invalid",
            Self::LeaseClosedOwner => "lease_closed_owner",
            Self::Unstable => "lease_closed_reason_unstable",
            Self::Decommission => "lease_closed_reason_decommission",
            Self::Unspecified => "lease_closed_reason_unspecified",
            Self::ManifestTimeout => "lease_closed_reason_manifest_timeout",
            Self::InsufficientFunds => "lease_closed_reason_insufficient_funds",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for LeaseClosedReason {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "lease_closed_invalid",
            "lease_closed_owner",
            "lease_closed_reason_unstable",
            "lease_closed_reason_decommission",
            "lease_closed_reason_unspecified",
            "lease_closed_reason_manifest_timeout",
            "lease_closed_reason_insufficient_funds",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = LeaseClosedReason;

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
                    "lease_closed_invalid" => Ok(LeaseClosedReason::LeaseClosedInvalid),
                    "lease_closed_owner" => Ok(LeaseClosedReason::LeaseClosedOwner),
                    "lease_closed_reason_unstable" => Ok(LeaseClosedReason::Unstable),
                    "lease_closed_reason_decommission" => Ok(LeaseClosedReason::Decommission),
                    "lease_closed_reason_unspecified" => Ok(LeaseClosedReason::Unspecified),
                    "lease_closed_reason_manifest_timeout" => Ok(LeaseClosedReason::ManifestTimeout),
                    "lease_closed_reason_insufficient_funds" => Ok(LeaseClosedReason::InsufficientFunds),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for LeaseFilters {
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
        if self.dseq != 0 {
            len += 1;
        }
        if self.gseq != 0 {
            len += 1;
        }
        if self.oseq != 0 {
            len += 1;
        }
        if !self.provider.is_empty() {
            len += 1;
        }
        if !self.state.is_empty() {
            len += 1;
        }
        if self.bseq != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.LeaseFilters", len)?;
        if !self.owner.is_empty() {
            struct_ser.serialize_field("owner", &self.owner)?;
        }
        if self.dseq != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("dseq", ToString::to_string(&self.dseq).as_str())?;
        }
        if self.gseq != 0 {
            struct_ser.serialize_field("gseq", &self.gseq)?;
        }
        if self.oseq != 0 {
            struct_ser.serialize_field("oseq", &self.oseq)?;
        }
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if !self.state.is_empty() {
            struct_ser.serialize_field("state", &self.state)?;
        }
        if self.bseq != 0 {
            struct_ser.serialize_field("bseq", &self.bseq)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for LeaseFilters {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "owner",
            "dseq",
            "gseq",
            "oseq",
            "provider",
            "state",
            "bseq",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Owner,
            Dseq,
            Gseq,
            Oseq,
            Provider,
            State,
            Bseq,
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
                            "dseq" => Ok(GeneratedField::Dseq),
                            "gseq" => Ok(GeneratedField::Gseq),
                            "oseq" => Ok(GeneratedField::Oseq),
                            "provider" => Ok(GeneratedField::Provider),
                            "state" => Ok(GeneratedField::State),
                            "bseq" => Ok(GeneratedField::Bseq),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = LeaseFilters;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.LeaseFilters")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<LeaseFilters, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut owner__ = None;
                let mut dseq__ = None;
                let mut gseq__ = None;
                let mut oseq__ = None;
                let mut provider__ = None;
                let mut state__ = None;
                let mut bseq__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Owner => {
                            if owner__.is_some() {
                                return Err(serde::de::Error::duplicate_field("owner"));
                            }
                            owner__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Dseq => {
                            if dseq__.is_some() {
                                return Err(serde::de::Error::duplicate_field("dseq"));
                            }
                            dseq__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Gseq => {
                            if gseq__.is_some() {
                                return Err(serde::de::Error::duplicate_field("gseq"));
                            }
                            gseq__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Oseq => {
                            if oseq__.is_some() {
                                return Err(serde::de::Error::duplicate_field("oseq"));
                            }
                            oseq__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::State => {
                            if state__.is_some() {
                                return Err(serde::de::Error::duplicate_field("state"));
                            }
                            state__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Bseq => {
                            if bseq__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bseq"));
                            }
                            bseq__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(LeaseFilters {
                    owner: owner__.unwrap_or_default(),
                    dseq: dseq__.unwrap_or_default(),
                    gseq: gseq__.unwrap_or_default(),
                    oseq: oseq__.unwrap_or_default(),
                    provider: provider__.unwrap_or_default(),
                    state: state__.unwrap_or_default(),
                    bseq: bseq__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.LeaseFilters", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for LeaseId {
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
        if self.dseq != 0 {
            len += 1;
        }
        if self.gseq != 0 {
            len += 1;
        }
        if self.oseq != 0 {
            len += 1;
        }
        if !self.provider.is_empty() {
            len += 1;
        }
        if self.bseq != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.LeaseID", len)?;
        if !self.owner.is_empty() {
            struct_ser.serialize_field("owner", &self.owner)?;
        }
        if self.dseq != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("dseq", ToString::to_string(&self.dseq).as_str())?;
        }
        if self.gseq != 0 {
            struct_ser.serialize_field("gseq", &self.gseq)?;
        }
        if self.oseq != 0 {
            struct_ser.serialize_field("oseq", &self.oseq)?;
        }
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if self.bseq != 0 {
            struct_ser.serialize_field("bseq", &self.bseq)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for LeaseId {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "owner",
            "dseq",
            "gseq",
            "oseq",
            "provider",
            "bseq",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Owner,
            Dseq,
            Gseq,
            Oseq,
            Provider,
            Bseq,
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
                            "dseq" => Ok(GeneratedField::Dseq),
                            "gseq" => Ok(GeneratedField::Gseq),
                            "oseq" => Ok(GeneratedField::Oseq),
                            "provider" => Ok(GeneratedField::Provider),
                            "bseq" => Ok(GeneratedField::Bseq),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = LeaseId;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.LeaseID")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<LeaseId, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut owner__ = None;
                let mut dseq__ = None;
                let mut gseq__ = None;
                let mut oseq__ = None;
                let mut provider__ = None;
                let mut bseq__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Owner => {
                            if owner__.is_some() {
                                return Err(serde::de::Error::duplicate_field("owner"));
                            }
                            owner__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Dseq => {
                            if dseq__.is_some() {
                                return Err(serde::de::Error::duplicate_field("dseq"));
                            }
                            dseq__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Gseq => {
                            if gseq__.is_some() {
                                return Err(serde::de::Error::duplicate_field("gseq"));
                            }
                            gseq__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Oseq => {
                            if oseq__.is_some() {
                                return Err(serde::de::Error::duplicate_field("oseq"));
                            }
                            oseq__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Bseq => {
                            if bseq__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bseq"));
                            }
                            bseq__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(LeaseId {
                    owner: owner__.unwrap_or_default(),
                    dseq: dseq__.unwrap_or_default(),
                    gseq: gseq__.unwrap_or_default(),
                    oseq: oseq__.unwrap_or_default(),
                    provider: provider__.unwrap_or_default(),
                    bseq: bseq__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.LeaseID", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgCloseBid {
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
        if self.reason != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.MsgCloseBid", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if self.reason != 0 {
            let v = LeaseClosedReason::try_from(self.reason)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.reason)))?;
            struct_ser.serialize_field("reason", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgCloseBid {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "id",
            "reason",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
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
            type Value = MsgCloseBid;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.MsgCloseBid")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgCloseBid, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                let mut reason__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
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
                            reason__ = Some(map_.next_value::<LeaseClosedReason>()? as i32);
                        }
                    }
                }
                Ok(MsgCloseBid {
                    id: id__,
                    reason: reason__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.MsgCloseBid", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgCloseBidResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.market.v2beta1.MsgCloseBidResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgCloseBidResponse {
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
            type Value = MsgCloseBidResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.MsgCloseBidResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgCloseBidResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgCloseBidResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.MsgCloseBidResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgCloseLease {
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
        if self.reason != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.MsgCloseLease", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if self.reason != 0 {
            let v = LeaseClosedReason::try_from(self.reason)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.reason)))?;
            struct_ser.serialize_field("reason", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgCloseLease {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "id",
            "reason",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
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
            type Value = MsgCloseLease;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.MsgCloseLease")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgCloseLease, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                let mut reason__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
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
                            reason__ = Some(map_.next_value::<LeaseClosedReason>()? as i32);
                        }
                    }
                }
                Ok(MsgCloseLease {
                    id: id__,
                    reason: reason__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.MsgCloseLease", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgCloseLeaseResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.market.v2beta1.MsgCloseLeaseResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgCloseLeaseResponse {
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
            type Value = MsgCloseLeaseResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.MsgCloseLeaseResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgCloseLeaseResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgCloseLeaseResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.MsgCloseLeaseResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgCreateBid {
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
        if !self.prices.is_empty() {
            len += 1;
        }
        if self.deposit.is_some() {
            len += 1;
        }
        if !self.resources_offer.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.MsgCreateBid", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if !self.prices.is_empty() {
            struct_ser.serialize_field("prices", &self.prices)?;
        }
        if let Some(v) = self.deposit.as_ref() {
            struct_ser.serialize_field("deposit", v)?;
        }
        if !self.resources_offer.is_empty() {
            struct_ser.serialize_field("resourcesOffer", &self.resources_offer)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgCreateBid {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "id",
            "prices",
            "deposit",
            "resources_offer",
            "resourcesOffer",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Id,
            Prices,
            Deposit,
            ResourcesOffer,
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
                            "prices" => Ok(GeneratedField::Prices),
                            "deposit" => Ok(GeneratedField::Deposit),
                            "resourcesOffer" | "resources_offer" => Ok(GeneratedField::ResourcesOffer),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgCreateBid;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.MsgCreateBid")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgCreateBid, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                let mut prices__ = None;
                let mut deposit__ = None;
                let mut resources_offer__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = map_.next_value()?;
                        }
                        GeneratedField::Prices => {
                            if prices__.is_some() {
                                return Err(serde::de::Error::duplicate_field("prices"));
                            }
                            prices__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Deposit => {
                            if deposit__.is_some() {
                                return Err(serde::de::Error::duplicate_field("deposit"));
                            }
                            deposit__ = map_.next_value()?;
                        }
                        GeneratedField::ResourcesOffer => {
                            if resources_offer__.is_some() {
                                return Err(serde::de::Error::duplicate_field("resourcesOffer"));
                            }
                            resources_offer__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(MsgCreateBid {
                    id: id__,
                    prices: prices__.unwrap_or_default(),
                    deposit: deposit__,
                    resources_offer: resources_offer__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.MsgCreateBid", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgCreateBidResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.market.v2beta1.MsgCreateBidResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgCreateBidResponse {
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
            type Value = MsgCreateBidResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.MsgCreateBidResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgCreateBidResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgCreateBidResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.MsgCreateBidResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgCreateLease {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.bid_id.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.MsgCreateLease", len)?;
        if let Some(v) = self.bid_id.as_ref() {
            struct_ser.serialize_field("bidId", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgCreateLease {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "bid_id",
            "bidId",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            BidId,
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
                            "bidId" | "bid_id" => Ok(GeneratedField::BidId),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgCreateLease;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.MsgCreateLease")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgCreateLease, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut bid_id__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::BidId => {
                            if bid_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bidId"));
                            }
                            bid_id__ = map_.next_value()?;
                        }
                    }
                }
                Ok(MsgCreateLease {
                    bid_id: bid_id__,
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.MsgCreateLease", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgCreateLeaseResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.market.v2beta1.MsgCreateLeaseResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgCreateLeaseResponse {
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
            type Value = MsgCreateLeaseResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.MsgCreateLeaseResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgCreateLeaseResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgCreateLeaseResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.MsgCreateLeaseResponse", FIELDS, GeneratedVisitor)
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
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.MsgUpdateParams", len)?;
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
                formatter.write_str("struct akash.market.v2beta1.MsgUpdateParams")
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
        deserializer.deserialize_struct("akash.market.v2beta1.MsgUpdateParams", FIELDS, GeneratedVisitor)
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
        let struct_ser = serializer.serialize_struct("akash.market.v2beta1.MsgUpdateParamsResponse", len)?;
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
                formatter.write_str("struct akash.market.v2beta1.MsgUpdateParamsResponse")
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
        deserializer.deserialize_struct("akash.market.v2beta1.MsgUpdateParamsResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgWithdrawLease {
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
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.MsgWithdrawLease", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgWithdrawLease {
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
            type Value = MsgWithdrawLease;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.MsgWithdrawLease")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgWithdrawLease, V::Error>
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
                Ok(MsgWithdrawLease {
                    id: id__,
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.MsgWithdrawLease", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgWithdrawLeaseResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.market.v2beta1.MsgWithdrawLeaseResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgWithdrawLeaseResponse {
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
            type Value = MsgWithdrawLeaseResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.MsgWithdrawLeaseResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgWithdrawLeaseResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgWithdrawLeaseResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.MsgWithdrawLeaseResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for Order {
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
        if self.state != 0 {
            len += 1;
        }
        if self.spec.is_some() {
            len += 1;
        }
        if self.created_at != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.Order", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        if self.state != 0 {
            let v = order::State::try_from(self.state)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.state)))?;
            struct_ser.serialize_field("state", &v)?;
        }
        if let Some(v) = self.spec.as_ref() {
            struct_ser.serialize_field("spec", v)?;
        }
        if self.created_at != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("createdAt", ToString::to_string(&self.created_at).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for Order {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "id",
            "state",
            "spec",
            "created_at",
            "createdAt",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Id,
            State,
            Spec,
            CreatedAt,
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
                            "state" => Ok(GeneratedField::State),
                            "spec" => Ok(GeneratedField::Spec),
                            "createdAt" | "created_at" => Ok(GeneratedField::CreatedAt),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = Order;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.Order")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<Order, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                let mut state__ = None;
                let mut spec__ = None;
                let mut created_at__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Id => {
                            if id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("id"));
                            }
                            id__ = map_.next_value()?;
                        }
                        GeneratedField::State => {
                            if state__.is_some() {
                                return Err(serde::de::Error::duplicate_field("state"));
                            }
                            state__ = Some(map_.next_value::<order::State>()? as i32);
                        }
                        GeneratedField::Spec => {
                            if spec__.is_some() {
                                return Err(serde::de::Error::duplicate_field("spec"));
                            }
                            spec__ = map_.next_value()?;
                        }
                        GeneratedField::CreatedAt => {
                            if created_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("createdAt"));
                            }
                            created_at__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(Order {
                    id: id__,
                    state: state__.unwrap_or_default(),
                    spec: spec__,
                    created_at: created_at__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.Order", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for order::State {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Invalid => "invalid",
            Self::Open => "open",
            Self::Active => "active",
            Self::Closed => "closed",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for order::State {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "invalid",
            "open",
            "active",
            "closed",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = order::State;

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
                    "invalid" => Ok(order::State::Invalid),
                    "open" => Ok(order::State::Open),
                    "active" => Ok(order::State::Active),
                    "closed" => Ok(order::State::Closed),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for OrderFilters {
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
        if self.dseq != 0 {
            len += 1;
        }
        if self.gseq != 0 {
            len += 1;
        }
        if self.oseq != 0 {
            len += 1;
        }
        if !self.state.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.OrderFilters", len)?;
        if !self.owner.is_empty() {
            struct_ser.serialize_field("owner", &self.owner)?;
        }
        if self.dseq != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("dseq", ToString::to_string(&self.dseq).as_str())?;
        }
        if self.gseq != 0 {
            struct_ser.serialize_field("gseq", &self.gseq)?;
        }
        if self.oseq != 0 {
            struct_ser.serialize_field("oseq", &self.oseq)?;
        }
        if !self.state.is_empty() {
            struct_ser.serialize_field("state", &self.state)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for OrderFilters {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "owner",
            "dseq",
            "gseq",
            "oseq",
            "state",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Owner,
            Dseq,
            Gseq,
            Oseq,
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
                            "owner" => Ok(GeneratedField::Owner),
                            "dseq" => Ok(GeneratedField::Dseq),
                            "gseq" => Ok(GeneratedField::Gseq),
                            "oseq" => Ok(GeneratedField::Oseq),
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
            type Value = OrderFilters;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.OrderFilters")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<OrderFilters, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut owner__ = None;
                let mut dseq__ = None;
                let mut gseq__ = None;
                let mut oseq__ = None;
                let mut state__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Owner => {
                            if owner__.is_some() {
                                return Err(serde::de::Error::duplicate_field("owner"));
                            }
                            owner__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Dseq => {
                            if dseq__.is_some() {
                                return Err(serde::de::Error::duplicate_field("dseq"));
                            }
                            dseq__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Gseq => {
                            if gseq__.is_some() {
                                return Err(serde::de::Error::duplicate_field("gseq"));
                            }
                            gseq__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Oseq => {
                            if oseq__.is_some() {
                                return Err(serde::de::Error::duplicate_field("oseq"));
                            }
                            oseq__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::State => {
                            if state__.is_some() {
                                return Err(serde::de::Error::duplicate_field("state"));
                            }
                            state__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(OrderFilters {
                    owner: owner__.unwrap_or_default(),
                    dseq: dseq__.unwrap_or_default(),
                    gseq: gseq__.unwrap_or_default(),
                    oseq: oseq__.unwrap_or_default(),
                    state: state__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.OrderFilters", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for OrderId {
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
        if self.dseq != 0 {
            len += 1;
        }
        if self.gseq != 0 {
            len += 1;
        }
        if self.oseq != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.OrderID", len)?;
        if !self.owner.is_empty() {
            struct_ser.serialize_field("owner", &self.owner)?;
        }
        if self.dseq != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("dseq", ToString::to_string(&self.dseq).as_str())?;
        }
        if self.gseq != 0 {
            struct_ser.serialize_field("gseq", &self.gseq)?;
        }
        if self.oseq != 0 {
            struct_ser.serialize_field("oseq", &self.oseq)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for OrderId {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "owner",
            "dseq",
            "gseq",
            "oseq",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Owner,
            Dseq,
            Gseq,
            Oseq,
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
                            "dseq" => Ok(GeneratedField::Dseq),
                            "gseq" => Ok(GeneratedField::Gseq),
                            "oseq" => Ok(GeneratedField::Oseq),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = OrderId;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.OrderID")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<OrderId, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut owner__ = None;
                let mut dseq__ = None;
                let mut gseq__ = None;
                let mut oseq__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Owner => {
                            if owner__.is_some() {
                                return Err(serde::de::Error::duplicate_field("owner"));
                            }
                            owner__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Dseq => {
                            if dseq__.is_some() {
                                return Err(serde::de::Error::duplicate_field("dseq"));
                            }
                            dseq__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Gseq => {
                            if gseq__.is_some() {
                                return Err(serde::de::Error::duplicate_field("gseq"));
                            }
                            gseq__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Oseq => {
                            if oseq__.is_some() {
                                return Err(serde::de::Error::duplicate_field("oseq"));
                            }
                            oseq__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(OrderId {
                    owner: owner__.unwrap_or_default(),
                    dseq: dseq__.unwrap_or_default(),
                    gseq: gseq__.unwrap_or_default(),
                    oseq: oseq__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.OrderID", FIELDS, GeneratedVisitor)
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
        if self.bid_min_deposit.is_some() {
            len += 1;
        }
        if self.order_max_bids != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.Params", len)?;
        if let Some(v) = self.bid_min_deposit.as_ref() {
            struct_ser.serialize_field("bidMinDeposit", v)?;
        }
        if self.order_max_bids != 0 {
            struct_ser.serialize_field("orderMaxBids", &self.order_max_bids)?;
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
            "bid_min_deposit",
            "bidMinDeposit",
            "order_max_bids",
            "orderMaxBids",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            BidMinDeposit,
            OrderMaxBids,
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
                            "bidMinDeposit" | "bid_min_deposit" => Ok(GeneratedField::BidMinDeposit),
                            "orderMaxBids" | "order_max_bids" => Ok(GeneratedField::OrderMaxBids),
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
                formatter.write_str("struct akash.market.v2beta1.Params")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<Params, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut bid_min_deposit__ = None;
                let mut order_max_bids__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::BidMinDeposit => {
                            if bid_min_deposit__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bidMinDeposit"));
                            }
                            bid_min_deposit__ = map_.next_value()?;
                        }
                        GeneratedField::OrderMaxBids => {
                            if order_max_bids__.is_some() {
                                return Err(serde::de::Error::duplicate_field("orderMaxBids"));
                            }
                            order_max_bids__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(Params {
                    bid_min_deposit: bid_min_deposit__,
                    order_max_bids: order_max_bids__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.Params", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryBidRequest {
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
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.QueryBidRequest", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryBidRequest {
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
            type Value = QueryBidRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.QueryBidRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryBidRequest, V::Error>
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
                Ok(QueryBidRequest {
                    id: id__,
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.QueryBidRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryBidResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.bid.is_some() {
            len += 1;
        }
        if self.escrow_account.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.QueryBidResponse", len)?;
        if let Some(v) = self.bid.as_ref() {
            struct_ser.serialize_field("bid", v)?;
        }
        if let Some(v) = self.escrow_account.as_ref() {
            struct_ser.serialize_field("escrowAccount", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryBidResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "bid",
            "escrow_account",
            "escrowAccount",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Bid,
            EscrowAccount,
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
                            "bid" => Ok(GeneratedField::Bid),
                            "escrowAccount" | "escrow_account" => Ok(GeneratedField::EscrowAccount),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryBidResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.QueryBidResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryBidResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut bid__ = None;
                let mut escrow_account__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Bid => {
                            if bid__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bid"));
                            }
                            bid__ = map_.next_value()?;
                        }
                        GeneratedField::EscrowAccount => {
                            if escrow_account__.is_some() {
                                return Err(serde::de::Error::duplicate_field("escrowAccount"));
                            }
                            escrow_account__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryBidResponse {
                    bid: bid__,
                    escrow_account: escrow_account__,
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.QueryBidResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryBidsRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.filters.is_some() {
            len += 1;
        }
        if self.pagination.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.QueryBidsRequest", len)?;
        if let Some(v) = self.filters.as_ref() {
            struct_ser.serialize_field("filters", v)?;
        }
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryBidsRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "filters",
            "pagination",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Filters,
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
                            "filters" => Ok(GeneratedField::Filters),
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
            type Value = QueryBidsRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.QueryBidsRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryBidsRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut filters__ = None;
                let mut pagination__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Filters => {
                            if filters__.is_some() {
                                return Err(serde::de::Error::duplicate_field("filters"));
                            }
                            filters__ = map_.next_value()?;
                        }
                        GeneratedField::Pagination => {
                            if pagination__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pagination"));
                            }
                            pagination__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryBidsRequest {
                    filters: filters__,
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.QueryBidsRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryBidsResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.bids.is_empty() {
            len += 1;
        }
        if self.pagination.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.QueryBidsResponse", len)?;
        if !self.bids.is_empty() {
            struct_ser.serialize_field("bids", &self.bids)?;
        }
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryBidsResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "bids",
            "pagination",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Bids,
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
                            "bids" => Ok(GeneratedField::Bids),
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
            type Value = QueryBidsResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.QueryBidsResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryBidsResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut bids__ = None;
                let mut pagination__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Bids => {
                            if bids__.is_some() {
                                return Err(serde::de::Error::duplicate_field("bids"));
                            }
                            bids__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Pagination => {
                            if pagination__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pagination"));
                            }
                            pagination__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryBidsResponse {
                    bids: bids__.unwrap_or_default(),
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.QueryBidsResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryLeaseRequest {
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
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.QueryLeaseRequest", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryLeaseRequest {
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
            type Value = QueryLeaseRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.QueryLeaseRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryLeaseRequest, V::Error>
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
                Ok(QueryLeaseRequest {
                    id: id__,
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.QueryLeaseRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryLeaseResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.lease.is_some() {
            len += 1;
        }
        if self.escrow_payment.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.QueryLeaseResponse", len)?;
        if let Some(v) = self.lease.as_ref() {
            struct_ser.serialize_field("lease", v)?;
        }
        if let Some(v) = self.escrow_payment.as_ref() {
            struct_ser.serialize_field("escrowPayment", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryLeaseResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "lease",
            "escrow_payment",
            "escrowPayment",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Lease,
            EscrowPayment,
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
                            "lease" => Ok(GeneratedField::Lease),
                            "escrowPayment" | "escrow_payment" => Ok(GeneratedField::EscrowPayment),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryLeaseResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.QueryLeaseResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryLeaseResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut lease__ = None;
                let mut escrow_payment__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Lease => {
                            if lease__.is_some() {
                                return Err(serde::de::Error::duplicate_field("lease"));
                            }
                            lease__ = map_.next_value()?;
                        }
                        GeneratedField::EscrowPayment => {
                            if escrow_payment__.is_some() {
                                return Err(serde::de::Error::duplicate_field("escrowPayment"));
                            }
                            escrow_payment__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryLeaseResponse {
                    lease: lease__,
                    escrow_payment: escrow_payment__,
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.QueryLeaseResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryLeasesRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.filters.is_some() {
            len += 1;
        }
        if self.pagination.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.QueryLeasesRequest", len)?;
        if let Some(v) = self.filters.as_ref() {
            struct_ser.serialize_field("filters", v)?;
        }
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryLeasesRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "filters",
            "pagination",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Filters,
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
                            "filters" => Ok(GeneratedField::Filters),
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
            type Value = QueryLeasesRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.QueryLeasesRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryLeasesRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut filters__ = None;
                let mut pagination__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Filters => {
                            if filters__.is_some() {
                                return Err(serde::de::Error::duplicate_field("filters"));
                            }
                            filters__ = map_.next_value()?;
                        }
                        GeneratedField::Pagination => {
                            if pagination__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pagination"));
                            }
                            pagination__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryLeasesRequest {
                    filters: filters__,
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.QueryLeasesRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryLeasesResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.leases.is_empty() {
            len += 1;
        }
        if self.pagination.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.QueryLeasesResponse", len)?;
        if !self.leases.is_empty() {
            struct_ser.serialize_field("leases", &self.leases)?;
        }
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryLeasesResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "leases",
            "pagination",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Leases,
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
                            "leases" => Ok(GeneratedField::Leases),
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
            type Value = QueryLeasesResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.QueryLeasesResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryLeasesResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut leases__ = None;
                let mut pagination__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Leases => {
                            if leases__.is_some() {
                                return Err(serde::de::Error::duplicate_field("leases"));
                            }
                            leases__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Pagination => {
                            if pagination__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pagination"));
                            }
                            pagination__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryLeasesResponse {
                    leases: leases__.unwrap_or_default(),
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.QueryLeasesResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryOrderRequest {
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
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.QueryOrderRequest", len)?;
        if let Some(v) = self.id.as_ref() {
            struct_ser.serialize_field("id", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryOrderRequest {
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
            type Value = QueryOrderRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.QueryOrderRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryOrderRequest, V::Error>
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
                Ok(QueryOrderRequest {
                    id: id__,
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.QueryOrderRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryOrderResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.order.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.QueryOrderResponse", len)?;
        if let Some(v) = self.order.as_ref() {
            struct_ser.serialize_field("order", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryOrderResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "order",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Order,
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
                            "order" => Ok(GeneratedField::Order),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryOrderResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.QueryOrderResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryOrderResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut order__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Order => {
                            if order__.is_some() {
                                return Err(serde::de::Error::duplicate_field("order"));
                            }
                            order__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryOrderResponse {
                    order: order__,
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.QueryOrderResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryOrdersRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.filters.is_some() {
            len += 1;
        }
        if self.pagination.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.QueryOrdersRequest", len)?;
        if let Some(v) = self.filters.as_ref() {
            struct_ser.serialize_field("filters", v)?;
        }
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryOrdersRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "filters",
            "pagination",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Filters,
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
                            "filters" => Ok(GeneratedField::Filters),
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
            type Value = QueryOrdersRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.QueryOrdersRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryOrdersRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut filters__ = None;
                let mut pagination__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Filters => {
                            if filters__.is_some() {
                                return Err(serde::de::Error::duplicate_field("filters"));
                            }
                            filters__ = map_.next_value()?;
                        }
                        GeneratedField::Pagination => {
                            if pagination__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pagination"));
                            }
                            pagination__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryOrdersRequest {
                    filters: filters__,
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.QueryOrdersRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryOrdersResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.orders.is_empty() {
            len += 1;
        }
        if self.pagination.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.QueryOrdersResponse", len)?;
        if !self.orders.is_empty() {
            struct_ser.serialize_field("orders", &self.orders)?;
        }
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryOrdersResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "orders",
            "pagination",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Orders,
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
                            "orders" => Ok(GeneratedField::Orders),
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
            type Value = QueryOrdersResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.QueryOrdersResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryOrdersResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut orders__ = None;
                let mut pagination__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Orders => {
                            if orders__.is_some() {
                                return Err(serde::de::Error::duplicate_field("orders"));
                            }
                            orders__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Pagination => {
                            if pagination__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pagination"));
                            }
                            pagination__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryOrdersResponse {
                    orders: orders__.unwrap_or_default(),
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.QueryOrdersResponse", FIELDS, GeneratedVisitor)
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
        let struct_ser = serializer.serialize_struct("akash.market.v2beta1.QueryParamsRequest", len)?;
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
                formatter.write_str("struct akash.market.v2beta1.QueryParamsRequest")
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
        deserializer.deserialize_struct("akash.market.v2beta1.QueryParamsRequest", FIELDS, GeneratedVisitor)
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
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.QueryParamsResponse", len)?;
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
                formatter.write_str("struct akash.market.v2beta1.QueryParamsResponse")
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
        deserializer.deserialize_struct("akash.market.v2beta1.QueryParamsResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for ResourceOffer {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.resources.is_some() {
            len += 1;
        }
        if self.count != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.market.v2beta1.ResourceOffer", len)?;
        if let Some(v) = self.resources.as_ref() {
            struct_ser.serialize_field("resources", v)?;
        }
        if self.count != 0 {
            struct_ser.serialize_field("count", &self.count)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for ResourceOffer {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "resources",
            "count",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Resources,
            Count,
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
                            "resources" => Ok(GeneratedField::Resources),
                            "count" => Ok(GeneratedField::Count),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = ResourceOffer;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.market.v2beta1.ResourceOffer")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<ResourceOffer, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut resources__ = None;
                let mut count__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Resources => {
                            if resources__.is_some() {
                                return Err(serde::de::Error::duplicate_field("resources"));
                            }
                            resources__ = map_.next_value()?;
                        }
                        GeneratedField::Count => {
                            if count__.is_some() {
                                return Err(serde::de::Error::duplicate_field("count"));
                            }
                            count__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(ResourceOffer {
                    resources: resources__,
                    count: count__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.market.v2beta1.ResourceOffer", FIELDS, GeneratedVisitor)
    }
}
