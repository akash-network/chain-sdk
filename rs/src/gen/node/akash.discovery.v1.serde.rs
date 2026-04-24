// @generated
impl serde::Serialize for Akash {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.client_info.is_some() {
            len += 1;
        }
        if !self.supported_versions.is_empty() {
            len += 1;
        }
        if !self.chain_id.is_empty() {
            len += 1;
        }
        if !self.node_version.is_empty() {
            len += 1;
        }
        if !self.min_client_version.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.discovery.v1.Akash", len)?;
        if let Some(v) = self.client_info.as_ref() {
            struct_ser.serialize_field("clientInfo", v)?;
        }
        if !self.supported_versions.is_empty() {
            struct_ser.serialize_field("supportedVersions", &self.supported_versions)?;
        }
        if !self.chain_id.is_empty() {
            struct_ser.serialize_field("chainId", &self.chain_id)?;
        }
        if !self.node_version.is_empty() {
            struct_ser.serialize_field("nodeVersion", &self.node_version)?;
        }
        if !self.min_client_version.is_empty() {
            struct_ser.serialize_field("minClientVersion", &self.min_client_version)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for Akash {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "client_info",
            "clientInfo",
            "supported_versions",
            "supportedVersions",
            "chain_id",
            "chainId",
            "node_version",
            "nodeVersion",
            "min_client_version",
            "minClientVersion",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            ClientInfo,
            SupportedVersions,
            ChainId,
            NodeVersion,
            MinClientVersion,
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
                            "clientInfo" | "client_info" => Ok(GeneratedField::ClientInfo),
                            "supportedVersions" | "supported_versions" => Ok(GeneratedField::SupportedVersions),
                            "chainId" | "chain_id" => Ok(GeneratedField::ChainId),
                            "nodeVersion" | "node_version" => Ok(GeneratedField::NodeVersion),
                            "minClientVersion" | "min_client_version" => Ok(GeneratedField::MinClientVersion),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = Akash;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.discovery.v1.Akash")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<Akash, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut client_info__ = None;
                let mut supported_versions__ = None;
                let mut chain_id__ = None;
                let mut node_version__ = None;
                let mut min_client_version__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::ClientInfo => {
                            if client_info__.is_some() {
                                return Err(serde::de::Error::duplicate_field("clientInfo"));
                            }
                            client_info__ = map_.next_value()?;
                        }
                        GeneratedField::SupportedVersions => {
                            if supported_versions__.is_some() {
                                return Err(serde::de::Error::duplicate_field("supportedVersions"));
                            }
                            supported_versions__ = Some(map_.next_value()?);
                        }
                        GeneratedField::ChainId => {
                            if chain_id__.is_some() {
                                return Err(serde::de::Error::duplicate_field("chainId"));
                            }
                            chain_id__ = Some(map_.next_value()?);
                        }
                        GeneratedField::NodeVersion => {
                            if node_version__.is_some() {
                                return Err(serde::de::Error::duplicate_field("nodeVersion"));
                            }
                            node_version__ = Some(map_.next_value()?);
                        }
                        GeneratedField::MinClientVersion => {
                            if min_client_version__.is_some() {
                                return Err(serde::de::Error::duplicate_field("minClientVersion"));
                            }
                            min_client_version__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(Akash {
                    client_info: client_info__,
                    supported_versions: supported_versions__.unwrap_or_default(),
                    chain_id: chain_id__.unwrap_or_default(),
                    node_version: node_version__.unwrap_or_default(),
                    min_client_version: min_client_version__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.discovery.v1.Akash", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for ClientInfo {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.api_version.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.discovery.v1.ClientInfo", len)?;
        if !self.api_version.is_empty() {
            struct_ser.serialize_field("apiVersion", &self.api_version)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for ClientInfo {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "api_version",
            "apiVersion",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            ApiVersion,
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
                            "apiVersion" | "api_version" => Ok(GeneratedField::ApiVersion),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = ClientInfo;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.discovery.v1.ClientInfo")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<ClientInfo, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut api_version__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::ApiVersion => {
                            if api_version__.is_some() {
                                return Err(serde::de::Error::duplicate_field("apiVersion"));
                            }
                            api_version__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(ClientInfo {
                    api_version: api_version__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.discovery.v1.ClientInfo", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for GetInfoRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let len = 0;
        let struct_ser = serializer.serialize_struct("akash.discovery.v1.GetInfoRequest", len)?;
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for GetInfoRequest {
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
            type Value = GetInfoRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.discovery.v1.GetInfoRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<GetInfoRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                while map_.next_key::<GeneratedField>()?.is_some() {
                    let _ = map_.next_value::<serde::de::IgnoredAny>()?;
                }
                Ok(GetInfoRequest {
                })
            }
        }
        deserializer.deserialize_struct("akash.discovery.v1.GetInfoRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for GetInfoResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.info.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.discovery.v1.GetInfoResponse", len)?;
        if let Some(v) = self.info.as_ref() {
            struct_ser.serialize_field("info", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for GetInfoResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "info",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
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
            type Value = GetInfoResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.discovery.v1.GetInfoResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<GetInfoResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut info__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Info => {
                            if info__.is_some() {
                                return Err(serde::de::Error::duplicate_field("info"));
                            }
                            info__ = map_.next_value()?;
                        }
                    }
                }
                Ok(GetInfoResponse {
                    info: info__,
                })
            }
        }
        deserializer.deserialize_struct("akash.discovery.v1.GetInfoResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for ModuleVersion {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.module.is_empty() {
            len += 1;
        }
        if !self.version.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.discovery.v1.ModuleVersion", len)?;
        if !self.module.is_empty() {
            struct_ser.serialize_field("module", &self.module)?;
        }
        if !self.version.is_empty() {
            struct_ser.serialize_field("version", &self.version)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for ModuleVersion {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "module",
            "version",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Module,
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

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "module" => Ok(GeneratedField::Module),
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
            type Value = ModuleVersion;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.discovery.v1.ModuleVersion")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<ModuleVersion, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut module__ = None;
                let mut version__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Module => {
                            if module__.is_some() {
                                return Err(serde::de::Error::duplicate_field("module"));
                            }
                            module__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Version => {
                            if version__.is_some() {
                                return Err(serde::de::Error::duplicate_field("version"));
                            }
                            version__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(ModuleVersion {
                    module: module__.unwrap_or_default(),
                    version: version__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.discovery.v1.ModuleVersion", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for VersionInfo {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.api_version.is_empty() {
            len += 1;
        }
        if !self.modules.is_empty() {
            len += 1;
        }
        if !self.features.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("akash.discovery.v1.VersionInfo", len)?;
        if !self.api_version.is_empty() {
            struct_ser.serialize_field("apiVersion", &self.api_version)?;
        }
        if !self.modules.is_empty() {
            struct_ser.serialize_field("modules", &self.modules)?;
        }
        if !self.features.is_empty() {
            struct_ser.serialize_field("features", &self.features)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for VersionInfo {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "api_version",
            "apiVersion",
            "modules",
            "features",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            ApiVersion,
            Modules,
            Features,
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
                            "apiVersion" | "api_version" => Ok(GeneratedField::ApiVersion),
                            "modules" => Ok(GeneratedField::Modules),
                            "features" => Ok(GeneratedField::Features),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = VersionInfo;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct akash.discovery.v1.VersionInfo")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<VersionInfo, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut api_version__ = None;
                let mut modules__ = None;
                let mut features__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::ApiVersion => {
                            if api_version__.is_some() {
                                return Err(serde::de::Error::duplicate_field("apiVersion"));
                            }
                            api_version__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Modules => {
                            if modules__.is_some() {
                                return Err(serde::de::Error::duplicate_field("modules"));
                            }
                            modules__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Features => {
                            if features__.is_some() {
                                return Err(serde::de::Error::duplicate_field("features"));
                            }
                            features__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(VersionInfo {
                    api_version: api_version__.unwrap_or_default(),
                    modules: modules__.unwrap_or_default(),
                    features: features__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("akash.discovery.v1.VersionInfo", FIELDS, GeneratedVisitor)
    }
}
