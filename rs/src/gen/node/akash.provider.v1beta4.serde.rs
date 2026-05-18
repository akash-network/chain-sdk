// @generated
impl serde::Serialize for EventProviderCreated {
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
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.EventProviderCreated", len)?;
        if !self.owner.is_empty() {
            struct_ser.serialize_field("owner", &self.owner)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventProviderCreated {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "owner",
        ];

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
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventProviderCreated;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.EventProviderCreated")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventProviderCreated, V::Error>
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
                Ok(EventProviderCreated {
                    owner: owner__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.EventProviderCreated", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventProviderDeleted {
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
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.EventProviderDeleted", len)?;
        if !self.owner.is_empty() {
            struct_ser.serialize_field("owner", &self.owner)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventProviderDeleted {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "owner",
        ];

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
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventProviderDeleted;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.EventProviderDeleted")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventProviderDeleted, V::Error>
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
                Ok(EventProviderDeleted {
                    owner: owner__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.EventProviderDeleted", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventProviderMaintenanceClosed {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.maintenance_id != 0 {
            len += 1;
        }
        if !self.provider.is_empty() {
            len += 1;
        }
        if self.closed_at.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.EventProviderMaintenanceClosed", len)?;
        if self.maintenance_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("maintenanceId", ToString::to_string(&self.maintenance_id).as_str())?;
        }
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if let Some(v) = self.closed_at.as_ref() {
            struct_ser.serialize_field("closedAt", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventProviderMaintenanceClosed {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "maintenance_id",
            "maintenanceId",
            "provider",
            "closed_at",
            "closedAt",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            MaintenanceId,
            Provider,
            ClosedAt,
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
                            "maintenanceId" | "maintenance_id" => Ok(GeneratedField::MaintenanceId),
                            "provider" => Ok(GeneratedField::Provider),
                            "closedAt" | "closed_at" => Ok(GeneratedField::ClosedAt),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventProviderMaintenanceClosed;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.EventProviderMaintenanceClosed")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventProviderMaintenanceClosed, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut maintenance_id__ = None;
                let mut provider__ = None;
                let mut closed_at__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::MaintenanceId => {
                            if maintenance_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maintenanceId"));
                            }
                            maintenance_id__ =
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::ClosedAt => {
                            if closed_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("closedAt"));
                            }
                            closed_at__ = map_.next_value()?;
                        }
                    }
                }
                Ok(EventProviderMaintenanceClosed {
                    maintenance_id: maintenance_id__.unwrap_or_default(),
                    provider: provider__.unwrap_or_default(),
                    closed_at: closed_at__,
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.EventProviderMaintenanceClosed", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventProviderMaintenanceOpened {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.maintenance_id != 0 {
            len += 1;
        }
        if !self.provider.is_empty() {
            len += 1;
        }
        if self.maintenance_type != 0 {
            len += 1;
        }
        if self.starts_at.is_some() {
            len += 1;
        }
        if self.expected_ends_at.is_some() {
            len += 1;
        }
        if !self.metadata_hash.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.EventProviderMaintenanceOpened", len)?;
        if self.maintenance_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("maintenanceId", ToString::to_string(&self.maintenance_id).as_str())?;
        }
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if self.maintenance_type != 0 {
            let v = ProviderMaintenanceType::try_from(self.maintenance_type)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.maintenance_type)))?;
            struct_ser.serialize_field("maintenanceType", &v)?;
        }
        if let Some(v) = self.starts_at.as_ref() {
            struct_ser.serialize_field("startsAt", v)?;
        }
        if let Some(v) = self.expected_ends_at.as_ref() {
            struct_ser.serialize_field("expectedEndsAt", v)?;
        }
        if !self.metadata_hash.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("metadataHash", pbjson::private::base64::encode(&self.metadata_hash).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventProviderMaintenanceOpened {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "maintenance_id",
            "maintenanceId",
            "provider",
            "maintenance_type",
            "maintenanceType",
            "starts_at",
            "startsAt",
            "expected_ends_at",
            "expectedEndsAt",
            "metadata_hash",
            "metadataHash",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            MaintenanceId,
            Provider,
            MaintenanceType,
            StartsAt,
            ExpectedEndsAt,
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
                            "maintenanceId" | "maintenance_id" => Ok(GeneratedField::MaintenanceId),
                            "provider" => Ok(GeneratedField::Provider),
                            "maintenanceType" | "maintenance_type" => Ok(GeneratedField::MaintenanceType),
                            "startsAt" | "starts_at" => Ok(GeneratedField::StartsAt),
                            "expectedEndsAt" | "expected_ends_at" => Ok(GeneratedField::ExpectedEndsAt),
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
            type Value = EventProviderMaintenanceOpened;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.EventProviderMaintenanceOpened")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventProviderMaintenanceOpened, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut maintenance_id__ = None;
                let mut provider__ = None;
                let mut maintenance_type__ = None;
                let mut starts_at__ = None;
                let mut expected_ends_at__ = None;
                let mut metadata_hash__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::MaintenanceId => {
                            if maintenance_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maintenanceId"));
                            }
                            maintenance_id__ =
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::MaintenanceType => {
                            if maintenance_type__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maintenanceType"));
                            }
                            maintenance_type__ = Some(map_.next_value::<ProviderMaintenanceType>()? as i32);
                        }
                        GeneratedField::StartsAt => {
                            if starts_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("startsAt"));
                            }
                            starts_at__ = map_.next_value()?;
                        }
                        GeneratedField::ExpectedEndsAt => {
                            if expected_ends_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("expectedEndsAt"));
                            }
                            expected_ends_at__ = map_.next_value()?;
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
                Ok(EventProviderMaintenanceOpened {
                    maintenance_id: maintenance_id__.unwrap_or_default(),
                    provider: provider__.unwrap_or_default(),
                    maintenance_type: maintenance_type__.unwrap_or_default(),
                    starts_at: starts_at__,
                    expected_ends_at: expected_ends_at__,
                    metadata_hash: metadata_hash__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.EventProviderMaintenanceOpened", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for EventProviderUpdated {
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
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.EventProviderUpdated", len)?;
        if !self.owner.is_empty() {
            struct_ser.serialize_field("owner", &self.owner)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for EventProviderUpdated {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "owner",
        ];

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
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = EventProviderUpdated;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.EventProviderUpdated")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<EventProviderUpdated, V::Error>
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
                Ok(EventProviderUpdated {
                    owner: owner__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.EventProviderUpdated", FIELDS, GeneratedVisitor)
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
        if !self.providers.is_empty() {
            len += 1;
        }
        if self.params.is_some() {
            len += 1;
        }
        if !self.maintenances.is_empty() {
            len += 1;
        }
        if self.next_maintenance_id != 0 {
            len += 1;
        }
        if !self.registrations.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.GenesisState", len)?;
        if !self.providers.is_empty() {
            struct_ser.serialize_field("providers", &self.providers)?;
        }
        if let Some(v) = self.params.as_ref() {
            struct_ser.serialize_field("params", v)?;
        }
        if !self.maintenances.is_empty() {
            struct_ser.serialize_field("maintenances", &self.maintenances)?;
        }
        if self.next_maintenance_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("nextMaintenanceId", ToString::to_string(&self.next_maintenance_id).as_str())?;
        }
        if !self.registrations.is_empty() {
            struct_ser.serialize_field("registrations", &self.registrations)?;
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
            "providers",
            "params",
            "maintenances",
            "next_maintenance_id",
            "nextMaintenanceId",
            "registrations",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Providers,
            Params,
            Maintenances,
            NextMaintenanceId,
            Registrations,
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
                            "providers" => Ok(GeneratedField::Providers),
                            "params" => Ok(GeneratedField::Params),
                            "maintenances" => Ok(GeneratedField::Maintenances),
                            "nextMaintenanceId" | "next_maintenance_id" => Ok(GeneratedField::NextMaintenanceId),
                            "registrations" => Ok(GeneratedField::Registrations),
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
                formatter.write_str("struct akash.provider.v1beta4.GenesisState")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<GenesisState, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut providers__ = None;
                let mut params__ = None;
                let mut maintenances__ = None;
                let mut next_maintenance_id__ = None;
                let mut registrations__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Providers => {
                            if providers__.is_some() {
                                return Err(serde::de::Error::duplicate_field("providers"));
                            }
                            providers__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Params => {
                            if params__.is_some() {
                                return Err(serde::de::Error::duplicate_field("params"));
                            }
                            params__ = map_.next_value()?;
                        }
                        GeneratedField::Maintenances => {
                            if maintenances__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maintenances"));
                            }
                            maintenances__ = Some(map_.next_value()?);
                        }
                        GeneratedField::NextMaintenanceId => {
                            if next_maintenance_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("nextMaintenanceId"));
                            }
                            next_maintenance_id__ =
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Registrations => {
                            if registrations__.is_some() {
                                return Err(serde::de::Error::duplicate_field("registrations"));
                            }
                            registrations__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(GenesisState {
                    providers: providers__.unwrap_or_default(),
                    params: params__,
                    maintenances: maintenances__.unwrap_or_default(),
                    next_maintenance_id: next_maintenance_id__.unwrap_or_default(),
                    registrations: registrations__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.GenesisState", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for Info {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.email.is_empty() {
            len += 1;
        }
        if !self.website.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.Info", len)?;
        if !self.email.is_empty() {
            struct_ser.serialize_field("email", &self.email)?;
        }
        if !self.website.is_empty() {
            struct_ser.serialize_field("website", &self.website)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for Info {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "email",
            "website",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Email,
            Website,
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
                            "email" => Ok(GeneratedField::Email),
                            "website" => Ok(GeneratedField::Website),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = Info;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.Info")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<Info, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut email__ = None;
                let mut website__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Email => {
                            if email__.is_some() {
                                return Err(serde::de::Error::duplicate_field("email"));
                            }
                            email__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Website => {
                            if website__.is_some() {
                                return Err(serde::de::Error::duplicate_field("website"));
                            }
                            website__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(Info {
                    email: email__.unwrap_or_default(),
                    website: website__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.Info", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgCloseProviderMaintenance {
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
        if self.maintenance_id != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.MsgCloseProviderMaintenance", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if self.maintenance_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("maintenanceId", ToString::to_string(&self.maintenance_id).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgCloseProviderMaintenance {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "maintenance_id",
            "maintenanceId",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            MaintenanceId,
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
                            "maintenanceId" | "maintenance_id" => Ok(GeneratedField::MaintenanceId),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgCloseProviderMaintenance;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.MsgCloseProviderMaintenance")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgCloseProviderMaintenance, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut maintenance_id__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::MaintenanceId => {
                            if maintenance_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maintenanceId"));
                            }
                            maintenance_id__ =
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(MsgCloseProviderMaintenance {
                    provider: provider__.unwrap_or_default(),
                    maintenance_id: maintenance_id__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.MsgCloseProviderMaintenance", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgCloseProviderMaintenanceResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.provider.v1beta4.MsgCloseProviderMaintenanceResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgCloseProviderMaintenanceResponse {
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
            type Value = MsgCloseProviderMaintenanceResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.MsgCloseProviderMaintenanceResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgCloseProviderMaintenanceResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgCloseProviderMaintenanceResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.MsgCloseProviderMaintenanceResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgCreateProvider {
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
        if !self.host_uri.is_empty() {
            len += 1;
        }
        if !self.attributes.is_empty() {
            len += 1;
        }
        if self.info.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.MsgCreateProvider", len)?;
        if !self.owner.is_empty() {
            struct_ser.serialize_field("owner", &self.owner)?;
        }
        if !self.host_uri.is_empty() {
            struct_ser.serialize_field("hostUri", &self.host_uri)?;
        }
        if !self.attributes.is_empty() {
            struct_ser.serialize_field("attributes", &self.attributes)?;
        }
        if let Some(v) = self.info.as_ref() {
            struct_ser.serialize_field("info", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgCreateProvider {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "owner",
            "host_uri",
            "hostUri",
            "attributes",
            "info",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Owner,
            HostUri,
            Attributes,
            Info,
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
                            "hostUri" | "host_uri" => Ok(GeneratedField::HostUri),
                            "attributes" => Ok(GeneratedField::Attributes),
                            "info" => Ok(GeneratedField::Info),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgCreateProvider;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.MsgCreateProvider")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgCreateProvider, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut owner__ = None;
                let mut host_uri__ = None;
                let mut attributes__ = None;
                let mut info__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Owner => {
                            if owner__.is_some() {
                                return Err(serde::de::Error::duplicate_field("owner"));
                            }
                            owner__ = Some(map_.next_value()?);
                        }
                        GeneratedField::HostUri => {
                            if host_uri__.is_some() {
                                return Err(serde::de::Error::duplicate_field("hostUri"));
                            }
                            host_uri__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Attributes => {
                            if attributes__.is_some() {
                                return Err(serde::de::Error::duplicate_field("attributes"));
                            }
                            attributes__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Info => {
                            if info__.is_some() {
                                return Err(serde::de::Error::duplicate_field("info"));
                            }
                            info__ = map_.next_value()?;
                        }
                    }
                }
                Ok(MsgCreateProvider {
                    owner: owner__.unwrap_or_default(),
                    host_uri: host_uri__.unwrap_or_default(),
                    attributes: attributes__.unwrap_or_default(),
                    info: info__,
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.MsgCreateProvider", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgCreateProviderResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.provider.v1beta4.MsgCreateProviderResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgCreateProviderResponse {
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
            type Value = MsgCreateProviderResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.MsgCreateProviderResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgCreateProviderResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgCreateProviderResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.MsgCreateProviderResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgDeleteProvider {
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
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.MsgDeleteProvider", len)?;
        if !self.owner.is_empty() {
            struct_ser.serialize_field("owner", &self.owner)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgDeleteProvider {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "owner",
        ];

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
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgDeleteProvider;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.MsgDeleteProvider")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgDeleteProvider, V::Error>
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
                Ok(MsgDeleteProvider {
                    owner: owner__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.MsgDeleteProvider", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgDeleteProviderResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.provider.v1beta4.MsgDeleteProviderResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgDeleteProviderResponse {
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
            type Value = MsgDeleteProviderResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.MsgDeleteProviderResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgDeleteProviderResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgDeleteProviderResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.MsgDeleteProviderResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgOpenProviderMaintenance {
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
        if self.maintenance_type != 0 {
            len += 1;
        }
        if self.starts_at.is_some() {
            len += 1;
        }
        if self.expected_ends_at.is_some() {
            len += 1;
        }
        if !self.metadata_hash.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.MsgOpenProviderMaintenance", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if self.maintenance_type != 0 {
            let v = ProviderMaintenanceType::try_from(self.maintenance_type)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.maintenance_type)))?;
            struct_ser.serialize_field("maintenanceType", &v)?;
        }
        if let Some(v) = self.starts_at.as_ref() {
            struct_ser.serialize_field("startsAt", v)?;
        }
        if let Some(v) = self.expected_ends_at.as_ref() {
            struct_ser.serialize_field("expectedEndsAt", v)?;
        }
        if !self.metadata_hash.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("metadataHash", pbjson::private::base64::encode(&self.metadata_hash).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgOpenProviderMaintenance {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "maintenance_type",
            "maintenanceType",
            "starts_at",
            "startsAt",
            "expected_ends_at",
            "expectedEndsAt",
            "metadata_hash",
            "metadataHash",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            MaintenanceType,
            StartsAt,
            ExpectedEndsAt,
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
                            "maintenanceType" | "maintenance_type" => Ok(GeneratedField::MaintenanceType),
                            "startsAt" | "starts_at" => Ok(GeneratedField::StartsAt),
                            "expectedEndsAt" | "expected_ends_at" => Ok(GeneratedField::ExpectedEndsAt),
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
            type Value = MsgOpenProviderMaintenance;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.MsgOpenProviderMaintenance")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgOpenProviderMaintenance, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut maintenance_type__ = None;
                let mut starts_at__ = None;
                let mut expected_ends_at__ = None;
                let mut metadata_hash__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::MaintenanceType => {
                            if maintenance_type__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maintenanceType"));
                            }
                            maintenance_type__ = Some(map_.next_value::<ProviderMaintenanceType>()? as i32);
                        }
                        GeneratedField::StartsAt => {
                            if starts_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("startsAt"));
                            }
                            starts_at__ = map_.next_value()?;
                        }
                        GeneratedField::ExpectedEndsAt => {
                            if expected_ends_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("expectedEndsAt"));
                            }
                            expected_ends_at__ = map_.next_value()?;
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
                Ok(MsgOpenProviderMaintenance {
                    provider: provider__.unwrap_or_default(),
                    maintenance_type: maintenance_type__.unwrap_or_default(),
                    starts_at: starts_at__,
                    expected_ends_at: expected_ends_at__,
                    metadata_hash: metadata_hash__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.MsgOpenProviderMaintenance", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgOpenProviderMaintenanceResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.maintenance_id != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.MsgOpenProviderMaintenanceResponse", len)?;
        if self.maintenance_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("maintenanceId", ToString::to_string(&self.maintenance_id).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgOpenProviderMaintenanceResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "maintenance_id",
            "maintenanceId",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            MaintenanceId,
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
                            "maintenanceId" | "maintenance_id" => Ok(GeneratedField::MaintenanceId),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgOpenProviderMaintenanceResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.MsgOpenProviderMaintenanceResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgOpenProviderMaintenanceResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut maintenance_id__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::MaintenanceId => {
                            if maintenance_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maintenanceId"));
                            }
                            maintenance_id__ =
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(MsgOpenProviderMaintenanceResponse {
                    maintenance_id: maintenance_id__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.MsgOpenProviderMaintenanceResponse", FIELDS, GeneratedVisitor)
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
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.MsgUpdateParams", len)?;
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
                formatter.write_str("struct akash.provider.v1beta4.MsgUpdateParams")
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
        deserializer.deserialize_struct("akash.provider.v1beta4.MsgUpdateParams", FIELDS, GeneratedVisitor)
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
        let struct_ser = serializer.serialize_struct("akash.provider.v1beta4.MsgUpdateParamsResponse", len)?;
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
                formatter.write_str("struct akash.provider.v1beta4.MsgUpdateParamsResponse")
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
        deserializer.deserialize_struct("akash.provider.v1beta4.MsgUpdateParamsResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgUpdateProvider {
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
        if !self.host_uri.is_empty() {
            len += 1;
        }
        if !self.attributes.is_empty() {
            len += 1;
        }
        if self.info.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.MsgUpdateProvider", len)?;
        if !self.owner.is_empty() {
            struct_ser.serialize_field("owner", &self.owner)?;
        }
        if !self.host_uri.is_empty() {
            struct_ser.serialize_field("hostUri", &self.host_uri)?;
        }
        if !self.attributes.is_empty() {
            struct_ser.serialize_field("attributes", &self.attributes)?;
        }
        if let Some(v) = self.info.as_ref() {
            struct_ser.serialize_field("info", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgUpdateProvider {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "owner",
            "host_uri",
            "hostUri",
            "attributes",
            "info",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Owner,
            HostUri,
            Attributes,
            Info,
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
                            "hostUri" | "host_uri" => Ok(GeneratedField::HostUri),
                            "attributes" => Ok(GeneratedField::Attributes),
                            "info" => Ok(GeneratedField::Info),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = MsgUpdateProvider;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.MsgUpdateProvider")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgUpdateProvider, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut owner__ = None;
                let mut host_uri__ = None;
                let mut attributes__ = None;
                let mut info__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Owner => {
                            if owner__.is_some() {
                                return Err(serde::de::Error::duplicate_field("owner"));
                            }
                            owner__ = Some(map_.next_value()?);
                        }
                        GeneratedField::HostUri => {
                            if host_uri__.is_some() {
                                return Err(serde::de::Error::duplicate_field("hostUri"));
                            }
                            host_uri__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Attributes => {
                            if attributes__.is_some() {
                                return Err(serde::de::Error::duplicate_field("attributes"));
                            }
                            attributes__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Info => {
                            if info__.is_some() {
                                return Err(serde::de::Error::duplicate_field("info"));
                            }
                            info__ = map_.next_value()?;
                        }
                    }
                }
                Ok(MsgUpdateProvider {
                    owner: owner__.unwrap_or_default(),
                    host_uri: host_uri__.unwrap_or_default(),
                    attributes: attributes__.unwrap_or_default(),
                    info: info__,
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.MsgUpdateProvider", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for MsgUpdateProviderResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.provider.v1beta4.MsgUpdateProviderResponse", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for MsgUpdateProviderResponse {
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
            type Value = MsgUpdateProviderResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.MsgUpdateProviderResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<MsgUpdateProviderResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(MsgUpdateProviderResponse {
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.MsgUpdateProviderResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for Provider {
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
        if !self.host_uri.is_empty() {
            len += 1;
        }
        if !self.attributes.is_empty() {
            len += 1;
        }
        if self.info.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.Provider", len)?;
        if !self.owner.is_empty() {
            struct_ser.serialize_field("owner", &self.owner)?;
        }
        if !self.host_uri.is_empty() {
            struct_ser.serialize_field("hostUri", &self.host_uri)?;
        }
        if !self.attributes.is_empty() {
            struct_ser.serialize_field("attributes", &self.attributes)?;
        }
        if let Some(v) = self.info.as_ref() {
            struct_ser.serialize_field("info", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for Provider {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "owner",
            "host_uri",
            "hostUri",
            "attributes",
            "info",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Owner,
            HostUri,
            Attributes,
            Info,
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
                            "hostUri" | "host_uri" => Ok(GeneratedField::HostUri),
                            "attributes" => Ok(GeneratedField::Attributes),
                            "info" => Ok(GeneratedField::Info),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = Provider;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.Provider")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<Provider, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut owner__ = None;
                let mut host_uri__ = None;
                let mut attributes__ = None;
                let mut info__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Owner => {
                            if owner__.is_some() {
                                return Err(serde::de::Error::duplicate_field("owner"));
                            }
                            owner__ = Some(map_.next_value()?);
                        }
                        GeneratedField::HostUri => {
                            if host_uri__.is_some() {
                                return Err(serde::de::Error::duplicate_field("hostUri"));
                            }
                            host_uri__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Attributes => {
                            if attributes__.is_some() {
                                return Err(serde::de::Error::duplicate_field("attributes"));
                            }
                            attributes__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Info => {
                            if info__.is_some() {
                                return Err(serde::de::Error::duplicate_field("info"));
                            }
                            info__ = map_.next_value()?;
                        }
                    }
                }
                Ok(Provider {
                    owner: owner__.unwrap_or_default(),
                    host_uri: host_uri__.unwrap_or_default(),
                    attributes: attributes__.unwrap_or_default(),
                    info: info__,
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.Provider", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for ProviderMaintenanceParams {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.maintenance_max_duration.is_some() {
            len += 1;
        }
        if self.maintenance_max_lookahead.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.ProviderMaintenanceParams", len)?;
        if let Some(v) = self.maintenance_max_duration.as_ref() {
            struct_ser.serialize_field("maintenanceMaxDuration", v)?;
        }
        if let Some(v) = self.maintenance_max_lookahead.as_ref() {
            struct_ser.serialize_field("maintenanceMaxLookahead", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for ProviderMaintenanceParams {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "maintenance_max_duration",
            "maintenanceMaxDuration",
            "maintenance_max_lookahead",
            "maintenanceMaxLookahead",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            MaintenanceMaxDuration,
            MaintenanceMaxLookahead,
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
                            "maintenanceMaxDuration" | "maintenance_max_duration" => Ok(GeneratedField::MaintenanceMaxDuration),
                            "maintenanceMaxLookahead" | "maintenance_max_lookahead" => Ok(GeneratedField::MaintenanceMaxLookahead),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = ProviderMaintenanceParams;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.ProviderMaintenanceParams")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<ProviderMaintenanceParams, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut maintenance_max_duration__ = None;
                let mut maintenance_max_lookahead__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::MaintenanceMaxDuration => {
                            if maintenance_max_duration__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maintenanceMaxDuration"));
                            }
                            maintenance_max_duration__ = map_.next_value()?;
                        }
                        GeneratedField::MaintenanceMaxLookahead => {
                            if maintenance_max_lookahead__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maintenanceMaxLookahead"));
                            }
                            maintenance_max_lookahead__ = map_.next_value()?;
                        }
                    }
                }
                Ok(ProviderMaintenanceParams {
                    maintenance_max_duration: maintenance_max_duration__,
                    maintenance_max_lookahead: maintenance_max_lookahead__,
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.ProviderMaintenanceParams", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for ProviderMaintenanceRecord {
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
        if self.maintenance_type != 0 {
            len += 1;
        }
        if self.starts_at.is_some() {
            len += 1;
        }
        if self.expected_ends_at.is_some() {
            len += 1;
        }
        if self.opened_at.is_some() {
            len += 1;
        }
        if self.closed_at.is_some() {
            len += 1;
        }
        if !self.metadata_hash.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.ProviderMaintenanceRecord", len)?;
        if self.id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("id", ToString::to_string(&self.id).as_str())?;
        }
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if self.maintenance_type != 0 {
            let v = ProviderMaintenanceType::try_from(self.maintenance_type)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.maintenance_type)))?;
            struct_ser.serialize_field("maintenanceType", &v)?;
        }
        if let Some(v) = self.starts_at.as_ref() {
            struct_ser.serialize_field("startsAt", v)?;
        }
        if let Some(v) = self.expected_ends_at.as_ref() {
            struct_ser.serialize_field("expectedEndsAt", v)?;
        }
        if let Some(v) = self.opened_at.as_ref() {
            struct_ser.serialize_field("openedAt", v)?;
        }
        if let Some(v) = self.closed_at.as_ref() {
            struct_ser.serialize_field("closedAt", v)?;
        }
        if !self.metadata_hash.is_empty() {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("metadataHash", pbjson::private::base64::encode(&self.metadata_hash).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for ProviderMaintenanceRecord {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "id",
            "provider",
            "maintenance_type",
            "maintenanceType",
            "starts_at",
            "startsAt",
            "expected_ends_at",
            "expectedEndsAt",
            "opened_at",
            "openedAt",
            "closed_at",
            "closedAt",
            "metadata_hash",
            "metadataHash",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Id,
            Provider,
            MaintenanceType,
            StartsAt,
            ExpectedEndsAt,
            OpenedAt,
            ClosedAt,
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
                            "id" => Ok(GeneratedField::Id),
                            "provider" => Ok(GeneratedField::Provider),
                            "maintenanceType" | "maintenance_type" => Ok(GeneratedField::MaintenanceType),
                            "startsAt" | "starts_at" => Ok(GeneratedField::StartsAt),
                            "expectedEndsAt" | "expected_ends_at" => Ok(GeneratedField::ExpectedEndsAt),
                            "openedAt" | "opened_at" => Ok(GeneratedField::OpenedAt),
                            "closedAt" | "closed_at" => Ok(GeneratedField::ClosedAt),
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
            type Value = ProviderMaintenanceRecord;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.ProviderMaintenanceRecord")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<ProviderMaintenanceRecord, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut id__ = None;
                let mut provider__ = None;
                let mut maintenance_type__ = None;
                let mut starts_at__ = None;
                let mut expected_ends_at__ = None;
                let mut opened_at__ = None;
                let mut closed_at__ = None;
                let mut metadata_hash__ = None;
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
                        GeneratedField::MaintenanceType => {
                            if maintenance_type__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maintenanceType"));
                            }
                            maintenance_type__ = Some(map_.next_value::<ProviderMaintenanceType>()? as i32);
                        }
                        GeneratedField::StartsAt => {
                            if starts_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("startsAt"));
                            }
                            starts_at__ = map_.next_value()?;
                        }
                        GeneratedField::ExpectedEndsAt => {
                            if expected_ends_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("expectedEndsAt"));
                            }
                            expected_ends_at__ = map_.next_value()?;
                        }
                        GeneratedField::OpenedAt => {
                            if opened_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("openedAt"));
                            }
                            opened_at__ = map_.next_value()?;
                        }
                        GeneratedField::ClosedAt => {
                            if closed_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("closedAt"));
                            }
                            closed_at__ = map_.next_value()?;
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
                Ok(ProviderMaintenanceRecord {
                    id: id__.unwrap_or_default(),
                    provider: provider__.unwrap_or_default(),
                    maintenance_type: maintenance_type__.unwrap_or_default(),
                    starts_at: starts_at__,
                    expected_ends_at: expected_ends_at__,
                    opened_at: opened_at__,
                    closed_at: closed_at__,
                    metadata_hash: metadata_hash__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.ProviderMaintenanceRecord", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for ProviderMaintenanceStatus {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "provider_maintenance_status_unspecified",
            Self::Scheduled => "provider_maintenance_status_scheduled",
            Self::Active => "provider_maintenance_status_active",
            Self::Elapsed => "provider_maintenance_status_elapsed",
            Self::Closed => "provider_maintenance_status_closed",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for ProviderMaintenanceStatus {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider_maintenance_status_unspecified",
            "provider_maintenance_status_scheduled",
            "provider_maintenance_status_active",
            "provider_maintenance_status_elapsed",
            "provider_maintenance_status_closed",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = ProviderMaintenanceStatus;

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
                    "provider_maintenance_status_unspecified" => Ok(ProviderMaintenanceStatus::Unspecified),
                    "provider_maintenance_status_scheduled" => Ok(ProviderMaintenanceStatus::Scheduled),
                    "provider_maintenance_status_active" => Ok(ProviderMaintenanceStatus::Active),
                    "provider_maintenance_status_elapsed" => Ok(ProviderMaintenanceStatus::Elapsed),
                    "provider_maintenance_status_closed" => Ok(ProviderMaintenanceStatus::Closed),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for ProviderMaintenanceType {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "provider_maintenance_type_unspecified",
            Self::Planned => "provider_maintenance_type_planned",
            Self::Emergency => "provider_maintenance_type_emergency",
            Self::Security => "provider_maintenance_type_security",
            Self::Network => "provider_maintenance_type_network",
            Self::Capacity => "provider_maintenance_type_capacity",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for ProviderMaintenanceType {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider_maintenance_type_unspecified",
            "provider_maintenance_type_planned",
            "provider_maintenance_type_emergency",
            "provider_maintenance_type_security",
            "provider_maintenance_type_network",
            "provider_maintenance_type_capacity",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = ProviderMaintenanceType;

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
                    "provider_maintenance_type_unspecified" => Ok(ProviderMaintenanceType::Unspecified),
                    "provider_maintenance_type_planned" => Ok(ProviderMaintenanceType::Planned),
                    "provider_maintenance_type_emergency" => Ok(ProviderMaintenanceType::Emergency),
                    "provider_maintenance_type_security" => Ok(ProviderMaintenanceType::Security),
                    "provider_maintenance_type_network" => Ok(ProviderMaintenanceType::Network),
                    "provider_maintenance_type_capacity" => Ok(ProviderMaintenanceType::Capacity),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for ProviderMaintenanceWithStatus {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.record.is_some() {
            len += 1;
        }
        if self.status != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.ProviderMaintenanceWithStatus", len)?;
        if let Some(v) = self.record.as_ref() {
            struct_ser.serialize_field("record", v)?;
        }
        if self.status != 0 {
            let v = ProviderMaintenanceStatus::try_from(self.status)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.status)))?;
            struct_ser.serialize_field("status", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for ProviderMaintenanceWithStatus {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "record",
            "status",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Record,
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
                            "record" => Ok(GeneratedField::Record),
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
            type Value = ProviderMaintenanceWithStatus;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.ProviderMaintenanceWithStatus")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<ProviderMaintenanceWithStatus, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut record__ = None;
                let mut status__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Record => {
                            if record__.is_some() {
                                return Err(serde::de::Error::duplicate_field("record"));
                            }
                            record__ = map_.next_value()?;
                        }
                        GeneratedField::Status => {
                            if status__.is_some() {
                                return Err(serde::de::Error::duplicate_field("status"));
                            }
                            status__ = Some(map_.next_value::<ProviderMaintenanceStatus>()? as i32);
                        }
                    }
                }
                Ok(ProviderMaintenanceWithStatus {
                    record: record__,
                    status: status__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.ProviderMaintenanceWithStatus", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for ProviderRegistration {
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
        if self.registered_at.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.ProviderRegistration", len)?;
        if !self.owner.is_empty() {
            struct_ser.serialize_field("owner", &self.owner)?;
        }
        if let Some(v) = self.registered_at.as_ref() {
            struct_ser.serialize_field("registeredAt", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for ProviderRegistration {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "owner",
            "registered_at",
            "registeredAt",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Owner,
            RegisteredAt,
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
                            "registeredAt" | "registered_at" => Ok(GeneratedField::RegisteredAt),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = ProviderRegistration;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.ProviderRegistration")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<ProviderRegistration, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut owner__ = None;
                let mut registered_at__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Owner => {
                            if owner__.is_some() {
                                return Err(serde::de::Error::duplicate_field("owner"));
                            }
                            owner__ = Some(map_.next_value()?);
                        }
                        GeneratedField::RegisteredAt => {
                            if registered_at__.is_some() {
                                return Err(serde::de::Error::duplicate_field("registeredAt"));
                            }
                            registered_at__ = map_.next_value()?;
                        }
                    }
                }
                Ok(ProviderRegistration {
                    owner: owner__.unwrap_or_default(),
                    registered_at: registered_at__,
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.ProviderRegistration", FIELDS, GeneratedVisitor)
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
        let struct_ser = serializer.serialize_struct("akash.provider.v1beta4.QueryParamsRequest", len)?;
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
                formatter.write_str("struct akash.provider.v1beta4.QueryParamsRequest")
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
        deserializer.deserialize_struct("akash.provider.v1beta4.QueryParamsRequest", FIELDS, GeneratedVisitor)
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
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.QueryParamsResponse", len)?;
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
                formatter.write_str("struct akash.provider.v1beta4.QueryParamsResponse")
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
        deserializer.deserialize_struct("akash.provider.v1beta4.QueryParamsResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryProviderMaintenanceRequest {
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
        if self.maintenance_id != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.QueryProviderMaintenanceRequest", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if self.maintenance_id != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("maintenanceId", ToString::to_string(&self.maintenance_id).as_str())?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryProviderMaintenanceRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "provider",
            "maintenance_id",
            "maintenanceId",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Provider,
            MaintenanceId,
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
                            "maintenanceId" | "maintenance_id" => Ok(GeneratedField::MaintenanceId),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryProviderMaintenanceRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.QueryProviderMaintenanceRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryProviderMaintenanceRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut provider__ = None;
                let mut maintenance_id__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value()?);
                        }
                        GeneratedField::MaintenanceId => {
                            if maintenance_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maintenanceId"));
                            }
                            maintenance_id__ =
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(QueryProviderMaintenanceRequest {
                    provider: provider__.unwrap_or_default(),
                    maintenance_id: maintenance_id__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.QueryProviderMaintenanceRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryProviderMaintenanceResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.maintenance.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.QueryProviderMaintenanceResponse", len)?;
        if let Some(v) = self.maintenance.as_ref() {
            struct_ser.serialize_field("maintenance", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryProviderMaintenanceResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "maintenance",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Maintenance,
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
                            "maintenance" => Ok(GeneratedField::Maintenance),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryProviderMaintenanceResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.QueryProviderMaintenanceResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryProviderMaintenanceResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut maintenance__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Maintenance => {
                            if maintenance__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maintenance"));
                            }
                            maintenance__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryProviderMaintenanceResponse {
                    maintenance: maintenance__,
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.QueryProviderMaintenanceResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryProviderMaintenancesRequest {
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
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.QueryProviderMaintenancesRequest", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        if self.status_filter != 0 {
            let v = ProviderMaintenanceStatus::try_from(self.status_filter)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.status_filter)))?;
            struct_ser.serialize_field("statusFilter", &v)?;
        }
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryProviderMaintenancesRequest {
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
            type Value = QueryProviderMaintenancesRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.QueryProviderMaintenancesRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryProviderMaintenancesRequest, V::Error>
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
                            status_filter__ = Some(map_.next_value::<ProviderMaintenanceStatus>()? as i32);
                        }
                        GeneratedField::Pagination => {
                            if pagination__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pagination"));
                            }
                            pagination__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryProviderMaintenancesRequest {
                    provider: provider__.unwrap_or_default(),
                    status_filter: status_filter__.unwrap_or_default(),
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.QueryProviderMaintenancesRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryProviderMaintenancesResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.maintenance.is_empty() {
            len += 1;
        }
        if self.pagination.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.QueryProviderMaintenancesResponse", len)?;
        if !self.maintenance.is_empty() {
            struct_ser.serialize_field("maintenance", &self.maintenance)?;
        }
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryProviderMaintenancesResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "maintenance",
            "pagination",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Maintenance,
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
                            "maintenance" => Ok(GeneratedField::Maintenance),
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
            type Value = QueryProviderMaintenancesResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.QueryProviderMaintenancesResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryProviderMaintenancesResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut maintenance__ = None;
                let mut pagination__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Maintenance => {
                            if maintenance__.is_some() {
                                return Err(serde::de::Error::duplicate_field("maintenance"));
                            }
                            maintenance__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Pagination => {
                            if pagination__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pagination"));
                            }
                            pagination__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryProviderMaintenancesResponse {
                    maintenance: maintenance__.unwrap_or_default(),
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.QueryProviderMaintenancesResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryProviderRequest {
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
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.QueryProviderRequest", len)?;
        if !self.owner.is_empty() {
            struct_ser.serialize_field("owner", &self.owner)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryProviderRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "owner",
        ];

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
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryProviderRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.QueryProviderRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryProviderRequest, V::Error>
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
                Ok(QueryProviderRequest {
                    owner: owner__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.QueryProviderRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryProviderResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.provider.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.QueryProviderResponse", len)?;
        if let Some(v) = self.provider.as_ref() {
            struct_ser.serialize_field("provider", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryProviderResponse {
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
            type Value = QueryProviderResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.QueryProviderResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryProviderResponse, V::Error>
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
                            provider__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryProviderResponse {
                    provider: provider__,
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.QueryProviderResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryProvidersRequest {
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
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.QueryProvidersRequest", len)?;
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryProvidersRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "pagination",
        ];

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

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
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
            type Value = QueryProvidersRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.QueryProvidersRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryProvidersRequest, V::Error>
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
                Ok(QueryProvidersRequest {
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.QueryProvidersRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryProvidersResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.providers.is_empty() {
            len += 1;
        }
        if self.pagination.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.QueryProvidersResponse", len)?;
        if !self.providers.is_empty() {
            struct_ser.serialize_field("providers", &self.providers)?;
        }
        if let Some(v) = self.pagination.as_ref() {
            struct_ser.serialize_field("pagination", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryProvidersResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "providers",
            "pagination",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Providers,
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
                            "providers" => Ok(GeneratedField::Providers),
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
            type Value = QueryProvidersResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.QueryProvidersResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryProvidersResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut providers__ = None;
                let mut pagination__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Providers => {
                            if providers__.is_some() {
                                return Err(serde::de::Error::duplicate_field("providers"));
                            }
                            providers__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Pagination => {
                            if pagination__.is_some() {
                                return Err(serde::de::Error::duplicate_field("pagination"));
                            }
                            pagination__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryProvidersResponse {
                    providers: providers__.unwrap_or_default(),
                    pagination: pagination__,
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.QueryProvidersResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryRegistrationRequest {
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
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.QueryRegistrationRequest", len)?;
        if !self.provider.is_empty() {
            struct_ser.serialize_field("provider", &self.provider)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryRegistrationRequest {
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
            type Value = QueryRegistrationRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.QueryRegistrationRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryRegistrationRequest, V::Error>
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
                Ok(QueryRegistrationRequest {
                    provider: provider__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.QueryRegistrationRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for QueryRegistrationResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.registration.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.provider.v1beta4.QueryRegistrationResponse", len)?;
        if let Some(v) = self.registration.as_ref() {
            struct_ser.serialize_field("registration", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for QueryRegistrationResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "registration",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Registration,
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
                            "registration" => Ok(GeneratedField::Registration),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = QueryRegistrationResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.provider.v1beta4.QueryRegistrationResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<QueryRegistrationResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut registration__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Registration => {
                            if registration__.is_some() {
                                return Err(serde::de::Error::duplicate_field("registration"));
                            }
                            registration__ = map_.next_value()?;
                        }
                    }
                }
                Ok(QueryRegistrationResponse {
                    registration: registration__,
                })
            }
        }
        deserializer.deserialize_struct("akash.provider.v1beta4.QueryRegistrationResponse", FIELDS, GeneratedVisitor)
    }
}
