// @generated
/// Generated client implementations.
pub mod query_client {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    use tonic::codegen::http::Uri;
    #[derive(Debug, Clone)]
    pub struct QueryClient<T> {
        inner: tonic::client::Grpc<T>,
    }
    impl QueryClient<tonic::transport::Channel> {
        /// Attempt to create a new client by connecting to a given endpoint.
        pub async fn connect<D>(dst: D) -> Result<Self, tonic::transport::Error>
        where
            D: TryInto<tonic::transport::Endpoint>,
            D::Error: Into<StdError>,
        {
            let conn = tonic::transport::Endpoint::new(dst)?.connect().await?;
            Ok(Self::new(conn))
        }
    }
    impl<T> QueryClient<T>
    where
        T: tonic::client::GrpcService<tonic::body::BoxBody>,
        T::Error: Into<StdError>,
        T::ResponseBody: Body<Data = Bytes> + Send + 'static,
        <T::ResponseBody as Body>::Error: Into<StdError> + Send,
    {
        pub fn new(inner: T) -> Self {
            let inner = tonic::client::Grpc::new(inner);
            Self { inner }
        }
        pub fn with_origin(inner: T, origin: Uri) -> Self {
            let inner = tonic::client::Grpc::with_origin(inner, origin);
            Self { inner }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> QueryClient<InterceptedService<T, F>>
        where
            F: tonic::service::Interceptor,
            T::ResponseBody: Default,
            T: tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
                Response = http::Response<
                    <T as tonic::client::GrpcService<tonic::body::BoxBody>>::ResponseBody,
                >,
            >,
            <T as tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
            >>::Error: Into<StdError> + Send + Sync,
        {
            QueryClient::new(InterceptedService::new(inner, interceptor))
        }
        /// Compress requests with the given encoding.
        ///
        /// This requires the server to support it otherwise it might respond with an
        /// error.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.send_compressed(encoding);
            self
        }
        /// Enable decompressing responses.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.accept_compressed(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_decoding_message_size(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_encoding_message_size(limit);
            self
        }
        pub async fn auditor(
            &mut self,
            request: impl tonic::IntoRequest<super::QueryAuditorRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryAuditorResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Query/Auditor",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("akash.verification.v1.Query", "Auditor"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn auditors(
            &mut self,
            request: impl tonic::IntoRequest<super::QueryAuditorsRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryAuditorsResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Query/Auditors",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("akash.verification.v1.Query", "Auditors"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn attestation(
            &mut self,
            request: impl tonic::IntoRequest<super::QueryAttestationRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryAttestationResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Query/Attestation",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("akash.verification.v1.Query", "Attestation"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn provider_attestations(
            &mut self,
            request: impl tonic::IntoRequest<super::QueryProviderAttestationsRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryProviderAttestationsResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Query/ProviderAttestations",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new(
                        "akash.verification.v1.Query",
                        "ProviderAttestations",
                    ),
                );
            self.inner.unary(req, path, codec).await
        }
        pub async fn auditor_attestations(
            &mut self,
            request: impl tonic::IntoRequest<super::QueryAuditorAttestationsRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryAuditorAttestationsResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Query/AuditorAttestations",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("akash.verification.v1.Query", "AuditorAttestations"),
                );
            self.inner.unary(req, path, codec).await
        }
        pub async fn discrepancy(
            &mut self,
            request: impl tonic::IntoRequest<super::QueryDiscrepancyRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryDiscrepancyResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Query/Discrepancy",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("akash.verification.v1.Query", "Discrepancy"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn discrepancies(
            &mut self,
            request: impl tonic::IntoRequest<super::QueryDiscrepanciesRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryDiscrepanciesResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Query/Discrepancies",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("akash.verification.v1.Query", "Discrepancies"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn audit_escrow(
            &mut self,
            request: impl tonic::IntoRequest<super::QueryAuditEscrowRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryAuditEscrowResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Query/AuditEscrow",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("akash.verification.v1.Query", "AuditEscrow"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn provider_audit_escrows(
            &mut self,
            request: impl tonic::IntoRequest<super::QueryProviderAuditEscrowsRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryProviderAuditEscrowsResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Query/ProviderAuditEscrows",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new(
                        "akash.verification.v1.Query",
                        "ProviderAuditEscrows",
                    ),
                );
            self.inner.unary(req, path, codec).await
        }
        pub async fn provider_verification_grace(
            &mut self,
            request: impl tonic::IntoRequest<
                super::QueryProviderVerificationGraceRequest,
            >,
        ) -> std::result::Result<
            tonic::Response<super::QueryProviderVerificationGraceResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Query/ProviderVerificationGrace",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new(
                        "akash.verification.v1.Query",
                        "ProviderVerificationGrace",
                    ),
                );
            self.inner.unary(req, path, codec).await
        }
        pub async fn provider_bond(
            &mut self,
            request: impl tonic::IntoRequest<super::QueryProviderBondRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryProviderBondResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Query/ProviderBond",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("akash.verification.v1.Query", "ProviderBond"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn provider_snapshot(
            &mut self,
            request: impl tonic::IntoRequest<super::QueryProviderSnapshotRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryProviderSnapshotResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Query/ProviderSnapshot",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("akash.verification.v1.Query", "ProviderSnapshot"),
                );
            self.inner.unary(req, path, codec).await
        }
        pub async fn params(
            &mut self,
            request: impl tonic::IntoRequest<super::QueryParamsRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryParamsResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Query/Params",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("akash.verification.v1.Query", "Params"));
            self.inner.unary(req, path, codec).await
        }
    }
}
/// Generated server implementations.
pub mod query_server {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    /// Generated trait containing gRPC methods that should be implemented for use with QueryServer.
    #[async_trait]
    pub trait Query: Send + Sync + 'static {
        async fn auditor(
            &self,
            request: tonic::Request<super::QueryAuditorRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryAuditorResponse>,
            tonic::Status,
        >;
        async fn auditors(
            &self,
            request: tonic::Request<super::QueryAuditorsRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryAuditorsResponse>,
            tonic::Status,
        >;
        async fn attestation(
            &self,
            request: tonic::Request<super::QueryAttestationRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryAttestationResponse>,
            tonic::Status,
        >;
        async fn provider_attestations(
            &self,
            request: tonic::Request<super::QueryProviderAttestationsRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryProviderAttestationsResponse>,
            tonic::Status,
        >;
        async fn auditor_attestations(
            &self,
            request: tonic::Request<super::QueryAuditorAttestationsRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryAuditorAttestationsResponse>,
            tonic::Status,
        >;
        async fn discrepancy(
            &self,
            request: tonic::Request<super::QueryDiscrepancyRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryDiscrepancyResponse>,
            tonic::Status,
        >;
        async fn discrepancies(
            &self,
            request: tonic::Request<super::QueryDiscrepanciesRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryDiscrepanciesResponse>,
            tonic::Status,
        >;
        async fn audit_escrow(
            &self,
            request: tonic::Request<super::QueryAuditEscrowRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryAuditEscrowResponse>,
            tonic::Status,
        >;
        async fn provider_audit_escrows(
            &self,
            request: tonic::Request<super::QueryProviderAuditEscrowsRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryProviderAuditEscrowsResponse>,
            tonic::Status,
        >;
        async fn provider_verification_grace(
            &self,
            request: tonic::Request<super::QueryProviderVerificationGraceRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryProviderVerificationGraceResponse>,
            tonic::Status,
        >;
        async fn provider_bond(
            &self,
            request: tonic::Request<super::QueryProviderBondRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryProviderBondResponse>,
            tonic::Status,
        >;
        async fn provider_snapshot(
            &self,
            request: tonic::Request<super::QueryProviderSnapshotRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryProviderSnapshotResponse>,
            tonic::Status,
        >;
        async fn params(
            &self,
            request: tonic::Request<super::QueryParamsRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryParamsResponse>,
            tonic::Status,
        >;
    }
    #[derive(Debug)]
    pub struct QueryServer<T: Query> {
        inner: Arc<T>,
        accept_compression_encodings: EnabledCompressionEncodings,
        send_compression_encodings: EnabledCompressionEncodings,
        max_decoding_message_size: Option<usize>,
        max_encoding_message_size: Option<usize>,
    }
    impl<T: Query> QueryServer<T> {
        pub fn new(inner: T) -> Self {
            Self::from_arc(Arc::new(inner))
        }
        pub fn from_arc(inner: Arc<T>) -> Self {
            Self {
                inner,
                accept_compression_encodings: Default::default(),
                send_compression_encodings: Default::default(),
                max_decoding_message_size: None,
                max_encoding_message_size: None,
            }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> InterceptedService<Self, F>
        where
            F: tonic::service::Interceptor,
        {
            InterceptedService::new(Self::new(inner), interceptor)
        }
        /// Enable decompressing requests with the given encoding.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.accept_compression_encodings.enable(encoding);
            self
        }
        /// Compress responses with the given encoding, if the client supports it.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.send_compression_encodings.enable(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.max_decoding_message_size = Some(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.max_encoding_message_size = Some(limit);
            self
        }
    }
    impl<T, B> tonic::codegen::Service<http::Request<B>> for QueryServer<T>
    where
        T: Query,
        B: Body + Send + 'static,
        B::Error: Into<StdError> + Send + 'static,
    {
        type Response = http::Response<tonic::body::BoxBody>;
        type Error = std::convert::Infallible;
        type Future = BoxFuture<Self::Response, Self::Error>;
        fn poll_ready(
            &mut self,
            _cx: &mut Context<'_>,
        ) -> Poll<std::result::Result<(), Self::Error>> {
            Poll::Ready(Ok(()))
        }
        fn call(&mut self, req: http::Request<B>) -> Self::Future {
            match req.uri().path() {
                "/akash.verification.v1.Query/Auditor" => {
                    #[allow(non_camel_case_types)]
                    struct AuditorSvc<T: Query>(pub Arc<T>);
                    impl<
                        T: Query,
                    > tonic::server::UnaryService<super::QueryAuditorRequest>
                    for AuditorSvc<T> {
                        type Response = super::QueryAuditorResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::QueryAuditorRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Query>::auditor(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = AuditorSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Query/Auditors" => {
                    #[allow(non_camel_case_types)]
                    struct AuditorsSvc<T: Query>(pub Arc<T>);
                    impl<
                        T: Query,
                    > tonic::server::UnaryService<super::QueryAuditorsRequest>
                    for AuditorsSvc<T> {
                        type Response = super::QueryAuditorsResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::QueryAuditorsRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Query>::auditors(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = AuditorsSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Query/Attestation" => {
                    #[allow(non_camel_case_types)]
                    struct AttestationSvc<T: Query>(pub Arc<T>);
                    impl<
                        T: Query,
                    > tonic::server::UnaryService<super::QueryAttestationRequest>
                    for AttestationSvc<T> {
                        type Response = super::QueryAttestationResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::QueryAttestationRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Query>::attestation(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = AttestationSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Query/ProviderAttestations" => {
                    #[allow(non_camel_case_types)]
                    struct ProviderAttestationsSvc<T: Query>(pub Arc<T>);
                    impl<
                        T: Query,
                    > tonic::server::UnaryService<
                        super::QueryProviderAttestationsRequest,
                    > for ProviderAttestationsSvc<T> {
                        type Response = super::QueryProviderAttestationsResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<
                                super::QueryProviderAttestationsRequest,
                            >,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Query>::provider_attestations(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = ProviderAttestationsSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Query/AuditorAttestations" => {
                    #[allow(non_camel_case_types)]
                    struct AuditorAttestationsSvc<T: Query>(pub Arc<T>);
                    impl<
                        T: Query,
                    > tonic::server::UnaryService<super::QueryAuditorAttestationsRequest>
                    for AuditorAttestationsSvc<T> {
                        type Response = super::QueryAuditorAttestationsResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<
                                super::QueryAuditorAttestationsRequest,
                            >,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Query>::auditor_attestations(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = AuditorAttestationsSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Query/Discrepancy" => {
                    #[allow(non_camel_case_types)]
                    struct DiscrepancySvc<T: Query>(pub Arc<T>);
                    impl<
                        T: Query,
                    > tonic::server::UnaryService<super::QueryDiscrepancyRequest>
                    for DiscrepancySvc<T> {
                        type Response = super::QueryDiscrepancyResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::QueryDiscrepancyRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Query>::discrepancy(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = DiscrepancySvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Query/Discrepancies" => {
                    #[allow(non_camel_case_types)]
                    struct DiscrepanciesSvc<T: Query>(pub Arc<T>);
                    impl<
                        T: Query,
                    > tonic::server::UnaryService<super::QueryDiscrepanciesRequest>
                    for DiscrepanciesSvc<T> {
                        type Response = super::QueryDiscrepanciesResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::QueryDiscrepanciesRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Query>::discrepancies(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = DiscrepanciesSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Query/AuditEscrow" => {
                    #[allow(non_camel_case_types)]
                    struct AuditEscrowSvc<T: Query>(pub Arc<T>);
                    impl<
                        T: Query,
                    > tonic::server::UnaryService<super::QueryAuditEscrowRequest>
                    for AuditEscrowSvc<T> {
                        type Response = super::QueryAuditEscrowResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::QueryAuditEscrowRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Query>::audit_escrow(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = AuditEscrowSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Query/ProviderAuditEscrows" => {
                    #[allow(non_camel_case_types)]
                    struct ProviderAuditEscrowsSvc<T: Query>(pub Arc<T>);
                    impl<
                        T: Query,
                    > tonic::server::UnaryService<
                        super::QueryProviderAuditEscrowsRequest,
                    > for ProviderAuditEscrowsSvc<T> {
                        type Response = super::QueryProviderAuditEscrowsResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<
                                super::QueryProviderAuditEscrowsRequest,
                            >,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Query>::provider_audit_escrows(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = ProviderAuditEscrowsSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Query/ProviderVerificationGrace" => {
                    #[allow(non_camel_case_types)]
                    struct ProviderVerificationGraceSvc<T: Query>(pub Arc<T>);
                    impl<
                        T: Query,
                    > tonic::server::UnaryService<
                        super::QueryProviderVerificationGraceRequest,
                    > for ProviderVerificationGraceSvc<T> {
                        type Response = super::QueryProviderVerificationGraceResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<
                                super::QueryProviderVerificationGraceRequest,
                            >,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Query>::provider_verification_grace(&inner, request)
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = ProviderVerificationGraceSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Query/ProviderBond" => {
                    #[allow(non_camel_case_types)]
                    struct ProviderBondSvc<T: Query>(pub Arc<T>);
                    impl<
                        T: Query,
                    > tonic::server::UnaryService<super::QueryProviderBondRequest>
                    for ProviderBondSvc<T> {
                        type Response = super::QueryProviderBondResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::QueryProviderBondRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Query>::provider_bond(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = ProviderBondSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Query/ProviderSnapshot" => {
                    #[allow(non_camel_case_types)]
                    struct ProviderSnapshotSvc<T: Query>(pub Arc<T>);
                    impl<
                        T: Query,
                    > tonic::server::UnaryService<super::QueryProviderSnapshotRequest>
                    for ProviderSnapshotSvc<T> {
                        type Response = super::QueryProviderSnapshotResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::QueryProviderSnapshotRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Query>::provider_snapshot(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = ProviderSnapshotSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Query/Params" => {
                    #[allow(non_camel_case_types)]
                    struct ParamsSvc<T: Query>(pub Arc<T>);
                    impl<T: Query> tonic::server::UnaryService<super::QueryParamsRequest>
                    for ParamsSvc<T> {
                        type Response = super::QueryParamsResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::QueryParamsRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Query>::params(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = ParamsSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                _ => {
                    Box::pin(async move {
                        Ok(
                            http::Response::builder()
                                .status(200)
                                .header("grpc-status", tonic::Code::Unimplemented as i32)
                                .header(
                                    http::header::CONTENT_TYPE,
                                    tonic::metadata::GRPC_CONTENT_TYPE,
                                )
                                .body(empty_body())
                                .unwrap(),
                        )
                    })
                }
            }
        }
    }
    impl<T: Query> Clone for QueryServer<T> {
        fn clone(&self) -> Self {
            let inner = self.inner.clone();
            Self {
                inner,
                accept_compression_encodings: self.accept_compression_encodings,
                send_compression_encodings: self.send_compression_encodings,
                max_decoding_message_size: self.max_decoding_message_size,
                max_encoding_message_size: self.max_encoding_message_size,
            }
        }
    }
    impl<T: Query> tonic::server::NamedService for QueryServer<T> {
        const NAME: &'static str = "akash.verification.v1.Query";
    }
}
/// Generated client implementations.
pub mod msg_client {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    use tonic::codegen::http::Uri;
    #[derive(Debug, Clone)]
    pub struct MsgClient<T> {
        inner: tonic::client::Grpc<T>,
    }
    impl MsgClient<tonic::transport::Channel> {
        /// Attempt to create a new client by connecting to a given endpoint.
        pub async fn connect<D>(dst: D) -> Result<Self, tonic::transport::Error>
        where
            D: TryInto<tonic::transport::Endpoint>,
            D::Error: Into<StdError>,
        {
            let conn = tonic::transport::Endpoint::new(dst)?.connect().await?;
            Ok(Self::new(conn))
        }
    }
    impl<T> MsgClient<T>
    where
        T: tonic::client::GrpcService<tonic::body::BoxBody>,
        T::Error: Into<StdError>,
        T::ResponseBody: Body<Data = Bytes> + Send + 'static,
        <T::ResponseBody as Body>::Error: Into<StdError> + Send,
    {
        pub fn new(inner: T) -> Self {
            let inner = tonic::client::Grpc::new(inner);
            Self { inner }
        }
        pub fn with_origin(inner: T, origin: Uri) -> Self {
            let inner = tonic::client::Grpc::with_origin(inner, origin);
            Self { inner }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> MsgClient<InterceptedService<T, F>>
        where
            F: tonic::service::Interceptor,
            T::ResponseBody: Default,
            T: tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
                Response = http::Response<
                    <T as tonic::client::GrpcService<tonic::body::BoxBody>>::ResponseBody,
                >,
            >,
            <T as tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
            >>::Error: Into<StdError> + Send + Sync,
        {
            MsgClient::new(InterceptedService::new(inner, interceptor))
        }
        /// Compress requests with the given encoding.
        ///
        /// This requires the server to support it otherwise it might respond with an
        /// error.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.send_compressed(encoding);
            self
        }
        /// Enable decompressing responses.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.accept_compressed(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_decoding_message_size(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_encoding_message_size(limit);
            self
        }
        /** PostAuditorBond posts (or tops up) an auditor's verification bond.
*/
        pub async fn post_auditor_bond(
            &mut self,
            request: impl tonic::IntoRequest<super::MsgPostAuditorBond>,
        ) -> std::result::Result<
            tonic::Response<super::MsgPostAuditorBondResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Msg/PostAuditorBond",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("akash.verification.v1.Msg", "PostAuditorBond"));
            self.inner.unary(req, path, codec).await
        }
        /** OpenAuditEscrow opens a new audit escrow funding a pending attestation.
*/
        pub async fn open_audit_escrow(
            &mut self,
            request: impl tonic::IntoRequest<super::MsgOpenAuditEscrow>,
        ) -> std::result::Result<
            tonic::Response<super::MsgOpenAuditEscrowResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Msg/OpenAuditEscrow",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("akash.verification.v1.Msg", "OpenAuditEscrow"));
            self.inner.unary(req, path, codec).await
        }
        /** CancelAuditEscrow cancels an open, unconsumed audit escrow before expiry
 and returns the fee and provider deposit to the provider.
*/
        pub async fn cancel_audit_escrow(
            &mut self,
            request: impl tonic::IntoRequest<super::MsgCancelAuditEscrow>,
        ) -> std::result::Result<
            tonic::Response<super::MsgCancelAuditEscrowResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Msg/CancelAuditEscrow",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("akash.verification.v1.Msg", "CancelAuditEscrow"),
                );
            self.inner.unary(req, path, codec).await
        }
        /** SettleAuditEscrow settles an unconsumed audit escrow with an explicit
 reason, fault attribution, and evidence reference.
*/
        pub async fn settle_audit_escrow(
            &mut self,
            request: impl tonic::IntoRequest<super::MsgSettleAuditEscrow>,
        ) -> std::result::Result<
            tonic::Response<super::MsgSettleAuditEscrowResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Msg/SettleAuditEscrow",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("akash.verification.v1.Msg", "SettleAuditEscrow"),
                );
            self.inner.unary(req, path, codec).await
        }
        /** SubmitAttestation submits an attestation about a provider; the first
 valid submission against a matching open escrow consumes it.
*/
        pub async fn submit_attestation(
            &mut self,
            request: impl tonic::IntoRequest<super::MsgSubmitAttestation>,
        ) -> std::result::Result<
            tonic::Response<super::MsgSubmitAttestationResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Msg/SubmitAttestation",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("akash.verification.v1.Msg", "SubmitAttestation"),
                );
            self.inner.unary(req, path, codec).await
        }
        /** RevokeAttestation revokes a previously submitted attestation with a
 typed reason and evidence reference.
*/
        pub async fn revoke_attestation(
            &mut self,
            request: impl tonic::IntoRequest<super::MsgRevokeAttestation>,
        ) -> std::result::Result<
            tonic::Response<super::MsgRevokeAttestationResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Msg/RevokeAttestation",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("akash.verification.v1.Msg", "RevokeAttestation"),
                );
            self.inner.unary(req, path, codec).await
        }
        /** RemoveAttestation voluntarily removes an attestation associated with the
 signing provider.
*/
        pub async fn remove_attestation(
            &mut self,
            request: impl tonic::IntoRequest<super::MsgRemoveAttestation>,
        ) -> std::result::Result<
            tonic::Response<super::MsgRemoveAttestationResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Msg/RemoveAttestation",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("akash.verification.v1.Msg", "RemoveAttestation"),
                );
            self.inner.unary(req, path, codec).await
        }
        /** ResignAuditor voluntarily exits the auditor role and begins unbonding
 of any posted auditor bond.
*/
        pub async fn resign_auditor(
            &mut self,
            request: impl tonic::IntoRequest<super::MsgResignAuditor>,
        ) -> std::result::Result<
            tonic::Response<super::MsgResignAuditorResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Msg/ResignAuditor",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("akash.verification.v1.Msg", "ResignAuditor"));
            self.inner.unary(req, path, codec).await
        }
        /** PostProviderBond posts (or tops up) a provider's resource-scaled
 verification bond.
*/
        pub async fn post_provider_bond(
            &mut self,
            request: impl tonic::IntoRequest<super::MsgPostProviderBond>,
        ) -> std::result::Result<
            tonic::Response<super::MsgPostProviderBondResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Msg/PostProviderBond",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("akash.verification.v1.Msg", "PostProviderBond"),
                );
            self.inner.unary(req, path, codec).await
        }
        /** WithdrawProviderBond initiates withdrawal of part or all of a provider's
 verification bond.
*/
        pub async fn withdraw_provider_bond(
            &mut self,
            request: impl tonic::IntoRequest<super::MsgWithdrawProviderBond>,
        ) -> std::result::Result<
            tonic::Response<super::MsgWithdrawProviderBondResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Msg/WithdrawProviderBond",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("akash.verification.v1.Msg", "WithdrawProviderBond"),
                );
            self.inner.unary(req, path, codec).await
        }
        /** PostSnapshotHash posts the provider's most recent resource snapshot hash
 and inline resource summary.
*/
        pub async fn post_snapshot_hash(
            &mut self,
            request: impl tonic::IntoRequest<super::MsgPostSnapshotHash>,
        ) -> std::result::Result<
            tonic::Response<super::MsgPostSnapshotHashResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Msg/PostSnapshotHash",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("akash.verification.v1.Msg", "PostSnapshotHash"),
                );
            self.inner.unary(req, path, codec).await
        }
        /** RegisterAuditor registers a new auditor with a maximum attestation tier;
 governance only.
*/
        pub async fn register_auditor(
            &mut self,
            request: impl tonic::IntoRequest<super::MsgRegisterAuditor>,
        ) -> std::result::Result<
            tonic::Response<super::MsgRegisterAuditorResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Msg/RegisterAuditor",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("akash.verification.v1.Msg", "RegisterAuditor"));
            self.inner.unary(req, path, codec).await
        }
        /** RenewAuditor renews an auditor's registration and resets the renewal
 deadline; governance only.
*/
        pub async fn renew_auditor(
            &mut self,
            request: impl tonic::IntoRequest<super::MsgRenewAuditor>,
        ) -> std::result::Result<
            tonic::Response<super::MsgRenewAuditorResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Msg/RenewAuditor",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("akash.verification.v1.Msg", "RenewAuditor"));
            self.inner.unary(req, path, codec).await
        }
        /** RemoveAuditor removes an auditor from the active set; governance only.
*/
        pub async fn remove_auditor(
            &mut self,
            request: impl tonic::IntoRequest<super::MsgRemoveAuditor>,
        ) -> std::result::Result<
            tonic::Response<super::MsgRemoveAuditorResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Msg/RemoveAuditor",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("akash.verification.v1.Msg", "RemoveAuditor"));
            self.inner.unary(req, path, codec).await
        }
        /** RevokeProviderAttestation revokes a single attestation for a specific
 provider/auditor pair; governance only.
*/
        pub async fn revoke_provider_attestation(
            &mut self,
            request: impl tonic::IntoRequest<super::MsgRevokeProviderAttestation>,
        ) -> std::result::Result<
            tonic::Response<super::MsgRevokeProviderAttestationResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Msg/RevokeProviderAttestation",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new(
                        "akash.verification.v1.Msg",
                        "RevokeProviderAttestation",
                    ),
                );
            self.inner.unary(req, path, codec).await
        }
        /** RevokeAllProviderAttestations revokes every active attestation for a
 single provider; governance only.
*/
        pub async fn revoke_all_provider_attestations(
            &mut self,
            request: impl tonic::IntoRequest<super::MsgRevokeAllProviderAttestations>,
        ) -> std::result::Result<
            tonic::Response<super::MsgRevokeAllProviderAttestationsResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Msg/RevokeAllProviderAttestations",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new(
                        "akash.verification.v1.Msg",
                        "RevokeAllProviderAttestations",
                    ),
                );
            self.inner.unary(req, path, codec).await
        }
        /** RevokeAuditorAttestations revokes every active attestation issued by a
 single auditor; governance only.
*/
        pub async fn revoke_auditor_attestations(
            &mut self,
            request: impl tonic::IntoRequest<super::MsgRevokeAuditorAttestations>,
        ) -> std::result::Result<
            tonic::Response<super::MsgRevokeAuditorAttestationsResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Msg/RevokeAuditorAttestations",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new(
                        "akash.verification.v1.Msg",
                        "RevokeAuditorAttestations",
                    ),
                );
            self.inner.unary(req, path, codec).await
        }
        /** ResolveDiscrepancy resolves a pending discrepancy between two auditors
 over the same provider; governance only.
*/
        pub async fn resolve_discrepancy(
            &mut self,
            request: impl tonic::IntoRequest<super::MsgResolveDiscrepancy>,
        ) -> std::result::Result<
            tonic::Response<super::MsgResolveDiscrepancyResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Msg/ResolveDiscrepancy",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("akash.verification.v1.Msg", "ResolveDiscrepancy"),
                );
            self.inner.unary(req, path, codec).await
        }
        /** SlashProviderBond slashes a fraction of a provider's verification bond;
 governance only.
*/
        pub async fn slash_provider_bond(
            &mut self,
            request: impl tonic::IntoRequest<super::MsgSlashProviderBond>,
        ) -> std::result::Result<
            tonic::Response<super::MsgSlashProviderBondResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Msg/SlashProviderBond",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("akash.verification.v1.Msg", "SlashProviderBond"),
                );
            self.inner.unary(req, path, codec).await
        }
        /** UpdateParams updates the x/verification module parameters; governance
 only.
*/
        pub async fn update_params(
            &mut self,
            request: impl tonic::IntoRequest<super::MsgUpdateParams>,
        ) -> std::result::Result<
            tonic::Response<super::MsgUpdateParamsResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/akash.verification.v1.Msg/UpdateParams",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("akash.verification.v1.Msg", "UpdateParams"));
            self.inner.unary(req, path, codec).await
        }
    }
}
/// Generated server implementations.
pub mod msg_server {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    /// Generated trait containing gRPC methods that should be implemented for use with MsgServer.
    #[async_trait]
    pub trait Msg: Send + Sync + 'static {
        /** PostAuditorBond posts (or tops up) an auditor's verification bond.
*/
        async fn post_auditor_bond(
            &self,
            request: tonic::Request<super::MsgPostAuditorBond>,
        ) -> std::result::Result<
            tonic::Response<super::MsgPostAuditorBondResponse>,
            tonic::Status,
        >;
        /** OpenAuditEscrow opens a new audit escrow funding a pending attestation.
*/
        async fn open_audit_escrow(
            &self,
            request: tonic::Request<super::MsgOpenAuditEscrow>,
        ) -> std::result::Result<
            tonic::Response<super::MsgOpenAuditEscrowResponse>,
            tonic::Status,
        >;
        /** CancelAuditEscrow cancels an open, unconsumed audit escrow before expiry
 and returns the fee and provider deposit to the provider.
*/
        async fn cancel_audit_escrow(
            &self,
            request: tonic::Request<super::MsgCancelAuditEscrow>,
        ) -> std::result::Result<
            tonic::Response<super::MsgCancelAuditEscrowResponse>,
            tonic::Status,
        >;
        /** SettleAuditEscrow settles an unconsumed audit escrow with an explicit
 reason, fault attribution, and evidence reference.
*/
        async fn settle_audit_escrow(
            &self,
            request: tonic::Request<super::MsgSettleAuditEscrow>,
        ) -> std::result::Result<
            tonic::Response<super::MsgSettleAuditEscrowResponse>,
            tonic::Status,
        >;
        /** SubmitAttestation submits an attestation about a provider; the first
 valid submission against a matching open escrow consumes it.
*/
        async fn submit_attestation(
            &self,
            request: tonic::Request<super::MsgSubmitAttestation>,
        ) -> std::result::Result<
            tonic::Response<super::MsgSubmitAttestationResponse>,
            tonic::Status,
        >;
        /** RevokeAttestation revokes a previously submitted attestation with a
 typed reason and evidence reference.
*/
        async fn revoke_attestation(
            &self,
            request: tonic::Request<super::MsgRevokeAttestation>,
        ) -> std::result::Result<
            tonic::Response<super::MsgRevokeAttestationResponse>,
            tonic::Status,
        >;
        /** RemoveAttestation voluntarily removes an attestation associated with the
 signing provider.
*/
        async fn remove_attestation(
            &self,
            request: tonic::Request<super::MsgRemoveAttestation>,
        ) -> std::result::Result<
            tonic::Response<super::MsgRemoveAttestationResponse>,
            tonic::Status,
        >;
        /** ResignAuditor voluntarily exits the auditor role and begins unbonding
 of any posted auditor bond.
*/
        async fn resign_auditor(
            &self,
            request: tonic::Request<super::MsgResignAuditor>,
        ) -> std::result::Result<
            tonic::Response<super::MsgResignAuditorResponse>,
            tonic::Status,
        >;
        /** PostProviderBond posts (or tops up) a provider's resource-scaled
 verification bond.
*/
        async fn post_provider_bond(
            &self,
            request: tonic::Request<super::MsgPostProviderBond>,
        ) -> std::result::Result<
            tonic::Response<super::MsgPostProviderBondResponse>,
            tonic::Status,
        >;
        /** WithdrawProviderBond initiates withdrawal of part or all of a provider's
 verification bond.
*/
        async fn withdraw_provider_bond(
            &self,
            request: tonic::Request<super::MsgWithdrawProviderBond>,
        ) -> std::result::Result<
            tonic::Response<super::MsgWithdrawProviderBondResponse>,
            tonic::Status,
        >;
        /** PostSnapshotHash posts the provider's most recent resource snapshot hash
 and inline resource summary.
*/
        async fn post_snapshot_hash(
            &self,
            request: tonic::Request<super::MsgPostSnapshotHash>,
        ) -> std::result::Result<
            tonic::Response<super::MsgPostSnapshotHashResponse>,
            tonic::Status,
        >;
        /** RegisterAuditor registers a new auditor with a maximum attestation tier;
 governance only.
*/
        async fn register_auditor(
            &self,
            request: tonic::Request<super::MsgRegisterAuditor>,
        ) -> std::result::Result<
            tonic::Response<super::MsgRegisterAuditorResponse>,
            tonic::Status,
        >;
        /** RenewAuditor renews an auditor's registration and resets the renewal
 deadline; governance only.
*/
        async fn renew_auditor(
            &self,
            request: tonic::Request<super::MsgRenewAuditor>,
        ) -> std::result::Result<
            tonic::Response<super::MsgRenewAuditorResponse>,
            tonic::Status,
        >;
        /** RemoveAuditor removes an auditor from the active set; governance only.
*/
        async fn remove_auditor(
            &self,
            request: tonic::Request<super::MsgRemoveAuditor>,
        ) -> std::result::Result<
            tonic::Response<super::MsgRemoveAuditorResponse>,
            tonic::Status,
        >;
        /** RevokeProviderAttestation revokes a single attestation for a specific
 provider/auditor pair; governance only.
*/
        async fn revoke_provider_attestation(
            &self,
            request: tonic::Request<super::MsgRevokeProviderAttestation>,
        ) -> std::result::Result<
            tonic::Response<super::MsgRevokeProviderAttestationResponse>,
            tonic::Status,
        >;
        /** RevokeAllProviderAttestations revokes every active attestation for a
 single provider; governance only.
*/
        async fn revoke_all_provider_attestations(
            &self,
            request: tonic::Request<super::MsgRevokeAllProviderAttestations>,
        ) -> std::result::Result<
            tonic::Response<super::MsgRevokeAllProviderAttestationsResponse>,
            tonic::Status,
        >;
        /** RevokeAuditorAttestations revokes every active attestation issued by a
 single auditor; governance only.
*/
        async fn revoke_auditor_attestations(
            &self,
            request: tonic::Request<super::MsgRevokeAuditorAttestations>,
        ) -> std::result::Result<
            tonic::Response<super::MsgRevokeAuditorAttestationsResponse>,
            tonic::Status,
        >;
        /** ResolveDiscrepancy resolves a pending discrepancy between two auditors
 over the same provider; governance only.
*/
        async fn resolve_discrepancy(
            &self,
            request: tonic::Request<super::MsgResolveDiscrepancy>,
        ) -> std::result::Result<
            tonic::Response<super::MsgResolveDiscrepancyResponse>,
            tonic::Status,
        >;
        /** SlashProviderBond slashes a fraction of a provider's verification bond;
 governance only.
*/
        async fn slash_provider_bond(
            &self,
            request: tonic::Request<super::MsgSlashProviderBond>,
        ) -> std::result::Result<
            tonic::Response<super::MsgSlashProviderBondResponse>,
            tonic::Status,
        >;
        /** UpdateParams updates the x/verification module parameters; governance
 only.
*/
        async fn update_params(
            &self,
            request: tonic::Request<super::MsgUpdateParams>,
        ) -> std::result::Result<
            tonic::Response<super::MsgUpdateParamsResponse>,
            tonic::Status,
        >;
    }
    #[derive(Debug)]
    pub struct MsgServer<T: Msg> {
        inner: Arc<T>,
        accept_compression_encodings: EnabledCompressionEncodings,
        send_compression_encodings: EnabledCompressionEncodings,
        max_decoding_message_size: Option<usize>,
        max_encoding_message_size: Option<usize>,
    }
    impl<T: Msg> MsgServer<T> {
        pub fn new(inner: T) -> Self {
            Self::from_arc(Arc::new(inner))
        }
        pub fn from_arc(inner: Arc<T>) -> Self {
            Self {
                inner,
                accept_compression_encodings: Default::default(),
                send_compression_encodings: Default::default(),
                max_decoding_message_size: None,
                max_encoding_message_size: None,
            }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> InterceptedService<Self, F>
        where
            F: tonic::service::Interceptor,
        {
            InterceptedService::new(Self::new(inner), interceptor)
        }
        /// Enable decompressing requests with the given encoding.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.accept_compression_encodings.enable(encoding);
            self
        }
        /// Compress responses with the given encoding, if the client supports it.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.send_compression_encodings.enable(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.max_decoding_message_size = Some(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.max_encoding_message_size = Some(limit);
            self
        }
    }
    impl<T, B> tonic::codegen::Service<http::Request<B>> for MsgServer<T>
    where
        T: Msg,
        B: Body + Send + 'static,
        B::Error: Into<StdError> + Send + 'static,
    {
        type Response = http::Response<tonic::body::BoxBody>;
        type Error = std::convert::Infallible;
        type Future = BoxFuture<Self::Response, Self::Error>;
        fn poll_ready(
            &mut self,
            _cx: &mut Context<'_>,
        ) -> Poll<std::result::Result<(), Self::Error>> {
            Poll::Ready(Ok(()))
        }
        fn call(&mut self, req: http::Request<B>) -> Self::Future {
            match req.uri().path() {
                "/akash.verification.v1.Msg/PostAuditorBond" => {
                    #[allow(non_camel_case_types)]
                    struct PostAuditorBondSvc<T: Msg>(pub Arc<T>);
                    impl<T: Msg> tonic::server::UnaryService<super::MsgPostAuditorBond>
                    for PostAuditorBondSvc<T> {
                        type Response = super::MsgPostAuditorBondResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::MsgPostAuditorBond>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Msg>::post_auditor_bond(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = PostAuditorBondSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Msg/OpenAuditEscrow" => {
                    #[allow(non_camel_case_types)]
                    struct OpenAuditEscrowSvc<T: Msg>(pub Arc<T>);
                    impl<T: Msg> tonic::server::UnaryService<super::MsgOpenAuditEscrow>
                    for OpenAuditEscrowSvc<T> {
                        type Response = super::MsgOpenAuditEscrowResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::MsgOpenAuditEscrow>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Msg>::open_audit_escrow(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = OpenAuditEscrowSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Msg/CancelAuditEscrow" => {
                    #[allow(non_camel_case_types)]
                    struct CancelAuditEscrowSvc<T: Msg>(pub Arc<T>);
                    impl<T: Msg> tonic::server::UnaryService<super::MsgCancelAuditEscrow>
                    for CancelAuditEscrowSvc<T> {
                        type Response = super::MsgCancelAuditEscrowResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::MsgCancelAuditEscrow>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Msg>::cancel_audit_escrow(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = CancelAuditEscrowSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Msg/SettleAuditEscrow" => {
                    #[allow(non_camel_case_types)]
                    struct SettleAuditEscrowSvc<T: Msg>(pub Arc<T>);
                    impl<T: Msg> tonic::server::UnaryService<super::MsgSettleAuditEscrow>
                    for SettleAuditEscrowSvc<T> {
                        type Response = super::MsgSettleAuditEscrowResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::MsgSettleAuditEscrow>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Msg>::settle_audit_escrow(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = SettleAuditEscrowSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Msg/SubmitAttestation" => {
                    #[allow(non_camel_case_types)]
                    struct SubmitAttestationSvc<T: Msg>(pub Arc<T>);
                    impl<T: Msg> tonic::server::UnaryService<super::MsgSubmitAttestation>
                    for SubmitAttestationSvc<T> {
                        type Response = super::MsgSubmitAttestationResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::MsgSubmitAttestation>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Msg>::submit_attestation(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = SubmitAttestationSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Msg/RevokeAttestation" => {
                    #[allow(non_camel_case_types)]
                    struct RevokeAttestationSvc<T: Msg>(pub Arc<T>);
                    impl<T: Msg> tonic::server::UnaryService<super::MsgRevokeAttestation>
                    for RevokeAttestationSvc<T> {
                        type Response = super::MsgRevokeAttestationResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::MsgRevokeAttestation>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Msg>::revoke_attestation(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = RevokeAttestationSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Msg/RemoveAttestation" => {
                    #[allow(non_camel_case_types)]
                    struct RemoveAttestationSvc<T: Msg>(pub Arc<T>);
                    impl<T: Msg> tonic::server::UnaryService<super::MsgRemoveAttestation>
                    for RemoveAttestationSvc<T> {
                        type Response = super::MsgRemoveAttestationResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::MsgRemoveAttestation>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Msg>::remove_attestation(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = RemoveAttestationSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Msg/ResignAuditor" => {
                    #[allow(non_camel_case_types)]
                    struct ResignAuditorSvc<T: Msg>(pub Arc<T>);
                    impl<T: Msg> tonic::server::UnaryService<super::MsgResignAuditor>
                    for ResignAuditorSvc<T> {
                        type Response = super::MsgResignAuditorResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::MsgResignAuditor>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Msg>::resign_auditor(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = ResignAuditorSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Msg/PostProviderBond" => {
                    #[allow(non_camel_case_types)]
                    struct PostProviderBondSvc<T: Msg>(pub Arc<T>);
                    impl<T: Msg> tonic::server::UnaryService<super::MsgPostProviderBond>
                    for PostProviderBondSvc<T> {
                        type Response = super::MsgPostProviderBondResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::MsgPostProviderBond>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Msg>::post_provider_bond(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = PostProviderBondSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Msg/WithdrawProviderBond" => {
                    #[allow(non_camel_case_types)]
                    struct WithdrawProviderBondSvc<T: Msg>(pub Arc<T>);
                    impl<
                        T: Msg,
                    > tonic::server::UnaryService<super::MsgWithdrawProviderBond>
                    for WithdrawProviderBondSvc<T> {
                        type Response = super::MsgWithdrawProviderBondResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::MsgWithdrawProviderBond>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Msg>::withdraw_provider_bond(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = WithdrawProviderBondSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Msg/PostSnapshotHash" => {
                    #[allow(non_camel_case_types)]
                    struct PostSnapshotHashSvc<T: Msg>(pub Arc<T>);
                    impl<T: Msg> tonic::server::UnaryService<super::MsgPostSnapshotHash>
                    for PostSnapshotHashSvc<T> {
                        type Response = super::MsgPostSnapshotHashResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::MsgPostSnapshotHash>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Msg>::post_snapshot_hash(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = PostSnapshotHashSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Msg/RegisterAuditor" => {
                    #[allow(non_camel_case_types)]
                    struct RegisterAuditorSvc<T: Msg>(pub Arc<T>);
                    impl<T: Msg> tonic::server::UnaryService<super::MsgRegisterAuditor>
                    for RegisterAuditorSvc<T> {
                        type Response = super::MsgRegisterAuditorResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::MsgRegisterAuditor>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Msg>::register_auditor(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = RegisterAuditorSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Msg/RenewAuditor" => {
                    #[allow(non_camel_case_types)]
                    struct RenewAuditorSvc<T: Msg>(pub Arc<T>);
                    impl<T: Msg> tonic::server::UnaryService<super::MsgRenewAuditor>
                    for RenewAuditorSvc<T> {
                        type Response = super::MsgRenewAuditorResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::MsgRenewAuditor>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Msg>::renew_auditor(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = RenewAuditorSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Msg/RemoveAuditor" => {
                    #[allow(non_camel_case_types)]
                    struct RemoveAuditorSvc<T: Msg>(pub Arc<T>);
                    impl<T: Msg> tonic::server::UnaryService<super::MsgRemoveAuditor>
                    for RemoveAuditorSvc<T> {
                        type Response = super::MsgRemoveAuditorResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::MsgRemoveAuditor>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Msg>::remove_auditor(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = RemoveAuditorSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Msg/RevokeProviderAttestation" => {
                    #[allow(non_camel_case_types)]
                    struct RevokeProviderAttestationSvc<T: Msg>(pub Arc<T>);
                    impl<
                        T: Msg,
                    > tonic::server::UnaryService<super::MsgRevokeProviderAttestation>
                    for RevokeProviderAttestationSvc<T> {
                        type Response = super::MsgRevokeProviderAttestationResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::MsgRevokeProviderAttestation>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Msg>::revoke_provider_attestation(&inner, request)
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = RevokeProviderAttestationSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Msg/RevokeAllProviderAttestations" => {
                    #[allow(non_camel_case_types)]
                    struct RevokeAllProviderAttestationsSvc<T: Msg>(pub Arc<T>);
                    impl<
                        T: Msg,
                    > tonic::server::UnaryService<
                        super::MsgRevokeAllProviderAttestations,
                    > for RevokeAllProviderAttestationsSvc<T> {
                        type Response = super::MsgRevokeAllProviderAttestationsResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<
                                super::MsgRevokeAllProviderAttestations,
                            >,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Msg>::revoke_all_provider_attestations(
                                        &inner,
                                        request,
                                    )
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = RevokeAllProviderAttestationsSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Msg/RevokeAuditorAttestations" => {
                    #[allow(non_camel_case_types)]
                    struct RevokeAuditorAttestationsSvc<T: Msg>(pub Arc<T>);
                    impl<
                        T: Msg,
                    > tonic::server::UnaryService<super::MsgRevokeAuditorAttestations>
                    for RevokeAuditorAttestationsSvc<T> {
                        type Response = super::MsgRevokeAuditorAttestationsResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::MsgRevokeAuditorAttestations>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Msg>::revoke_auditor_attestations(&inner, request)
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = RevokeAuditorAttestationsSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Msg/ResolveDiscrepancy" => {
                    #[allow(non_camel_case_types)]
                    struct ResolveDiscrepancySvc<T: Msg>(pub Arc<T>);
                    impl<
                        T: Msg,
                    > tonic::server::UnaryService<super::MsgResolveDiscrepancy>
                    for ResolveDiscrepancySvc<T> {
                        type Response = super::MsgResolveDiscrepancyResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::MsgResolveDiscrepancy>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Msg>::resolve_discrepancy(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = ResolveDiscrepancySvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Msg/SlashProviderBond" => {
                    #[allow(non_camel_case_types)]
                    struct SlashProviderBondSvc<T: Msg>(pub Arc<T>);
                    impl<T: Msg> tonic::server::UnaryService<super::MsgSlashProviderBond>
                    for SlashProviderBondSvc<T> {
                        type Response = super::MsgSlashProviderBondResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::MsgSlashProviderBond>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Msg>::slash_provider_bond(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = SlashProviderBondSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/akash.verification.v1.Msg/UpdateParams" => {
                    #[allow(non_camel_case_types)]
                    struct UpdateParamsSvc<T: Msg>(pub Arc<T>);
                    impl<T: Msg> tonic::server::UnaryService<super::MsgUpdateParams>
                    for UpdateParamsSvc<T> {
                        type Response = super::MsgUpdateParamsResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::MsgUpdateParams>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as Msg>::update_params(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = UpdateParamsSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                _ => {
                    Box::pin(async move {
                        Ok(
                            http::Response::builder()
                                .status(200)
                                .header("grpc-status", tonic::Code::Unimplemented as i32)
                                .header(
                                    http::header::CONTENT_TYPE,
                                    tonic::metadata::GRPC_CONTENT_TYPE,
                                )
                                .body(empty_body())
                                .unwrap(),
                        )
                    })
                }
            }
        }
    }
    impl<T: Msg> Clone for MsgServer<T> {
        fn clone(&self) -> Self {
            let inner = self.inner.clone();
            Self {
                inner,
                accept_compression_encodings: self.accept_compression_encodings,
                send_compression_encodings: self.send_compression_encodings,
                max_decoding_message_size: self.max_decoding_message_size,
                max_encoding_message_size: self.max_encoding_message_size,
            }
        }
    }
    impl<T: Msg> tonic::server::NamedService for MsgServer<T> {
        const NAME: &'static str = "akash.verification.v1.Msg";
    }
}
