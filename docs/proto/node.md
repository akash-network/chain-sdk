<!-- This file is auto-generated. Please do not modify it yourself. -->
 # Protobuf Documentation
 <a name="top"></a>

 ## Table of Contents
 
 - [akash/verification/v1/types.proto](#akash/verification/v1/types.proto)
     - [AttestationRevocationReason](#akash.verification.v1.AttestationRevocationReason)
     - [AttestationStatus](#akash.verification.v1.AttestationStatus)
     - [AuditEscrowSettlementReason](#akash.verification.v1.AuditEscrowSettlementReason)
     - [AuditEscrowStatus](#akash.verification.v1.AuditEscrowStatus)
     - [AuditorSelectionMode](#akash.verification.v1.AuditorSelectionMode)
     - [AuditorStatus](#akash.verification.v1.AuditorStatus)
     - [BondStatus](#akash.verification.v1.BondStatus)
     - [CapabilityFlag](#akash.verification.v1.CapabilityFlag)
     - [DepositStatus](#akash.verification.v1.DepositStatus)
     - [DiscrepancyResolutionReason](#akash.verification.v1.DiscrepancyResolutionReason)
     - [DiscrepancyStatus](#akash.verification.v1.DiscrepancyStatus)
     - [FaultAttribution](#akash.verification.v1.FaultAttribution)
     - [FeeStatus](#akash.verification.v1.FeeStatus)
     - [GovernanceAttestationReason](#akash.verification.v1.GovernanceAttestationReason)
     - [ProviderBondSlashReason](#akash.verification.v1.ProviderBondSlashReason)
     - [ProviderDepositStatus](#akash.verification.v1.ProviderDepositStatus)
     - [VerificationGraceStatus](#akash.verification.v1.VerificationGraceStatus)
     - [VerificationTier](#akash.verification.v1.VerificationTier)
     - [VoidedReason](#akash.verification.v1.VoidedReason)
   
 - [akash/verification/v1/verificationrequirement.proto](#akash/verification/v1/verificationrequirement.proto)
     - [VerificationRequirement](#akash.verification.v1.VerificationRequirement)
   
 - [akash/base/attributes/v1/attribute.proto](#akash/base/attributes/v1/attribute.proto)
     - [Attribute](#akash.base.attributes.v1.Attribute)
     - [PlacementRequirements](#akash.base.attributes.v1.PlacementRequirements)
     - [SignedBy](#akash.base.attributes.v1.SignedBy)
   
 - [akash/audit/v1/audit.proto](#akash/audit/v1/audit.proto)
     - [AttributesFilters](#akash.audit.v1.AttributesFilters)
     - [AuditedAttributesStore](#akash.audit.v1.AuditedAttributesStore)
     - [AuditedProvider](#akash.audit.v1.AuditedProvider)
   
 - [akash/audit/v1/event.proto](#akash/audit/v1/event.proto)
     - [EventTrustedAuditorCreated](#akash.audit.v1.EventTrustedAuditorCreated)
     - [EventTrustedAuditorDeleted](#akash.audit.v1.EventTrustedAuditorDeleted)
   
 - [akash/audit/v1/genesis.proto](#akash/audit/v1/genesis.proto)
     - [GenesisState](#akash.audit.v1.GenesisState)
   
 - [akash/audit/v1/msg.proto](#akash/audit/v1/msg.proto)
     - [MsgDeleteProviderAttributes](#akash.audit.v1.MsgDeleteProviderAttributes)
     - [MsgDeleteProviderAttributesResponse](#akash.audit.v1.MsgDeleteProviderAttributesResponse)
     - [MsgSignProviderAttributes](#akash.audit.v1.MsgSignProviderAttributes)
     - [MsgSignProviderAttributesResponse](#akash.audit.v1.MsgSignProviderAttributesResponse)
   
 - [akash/audit/v1/query.proto](#akash/audit/v1/query.proto)
     - [QueryAllProvidersAttributesRequest](#akash.audit.v1.QueryAllProvidersAttributesRequest)
     - [QueryAuditorAttributesRequest](#akash.audit.v1.QueryAuditorAttributesRequest)
     - [QueryProviderAttributesRequest](#akash.audit.v1.QueryProviderAttributesRequest)
     - [QueryProviderAuditorRequest](#akash.audit.v1.QueryProviderAuditorRequest)
     - [QueryProviderRequest](#akash.audit.v1.QueryProviderRequest)
     - [QueryProvidersResponse](#akash.audit.v1.QueryProvidersResponse)
   
     - [Query](#akash.audit.v1.Query)
   
 - [akash/audit/v1/service.proto](#akash/audit/v1/service.proto)
     - [Msg](#akash.audit.v1.Msg)
   
 - [akash/base/deposit/v1/deposit.proto](#akash/base/deposit/v1/deposit.proto)
     - [Deposit](#akash.base.deposit.v1.Deposit)
   
     - [Source](#akash.base.deposit.v1.Source)
   
 - [akash/base/offchain/sign/v1/sign.proto](#akash/base/offchain/sign/v1/sign.proto)
     - [MsgSignData](#akash.base.offchain.sign.v1.MsgSignData)
   
 - [akash/base/resources/v1beta4/resourcevalue.proto](#akash/base/resources/v1beta4/resourcevalue.proto)
     - [ResourceValue](#akash.base.resources.v1beta4.ResourceValue)
   
 - [akash/base/resources/v1beta4/cpu.proto](#akash/base/resources/v1beta4/cpu.proto)
     - [CPU](#akash.base.resources.v1beta4.CPU)
   
 - [akash/base/resources/v1beta4/endpoint.proto](#akash/base/resources/v1beta4/endpoint.proto)
     - [Endpoint](#akash.base.resources.v1beta4.Endpoint)
   
     - [Endpoint.Kind](#akash.base.resources.v1beta4.Endpoint.Kind)
   
 - [akash/base/resources/v1beta4/gpu.proto](#akash/base/resources/v1beta4/gpu.proto)
     - [GPU](#akash.base.resources.v1beta4.GPU)
   
 - [akash/base/resources/v1beta4/memory.proto](#akash/base/resources/v1beta4/memory.proto)
     - [Memory](#akash.base.resources.v1beta4.Memory)
   
 - [akash/base/resources/v1beta4/storage.proto](#akash/base/resources/v1beta4/storage.proto)
     - [Storage](#akash.base.resources.v1beta4.Storage)
   
 - [akash/base/resources/v1beta4/resources.proto](#akash/base/resources/v1beta4/resources.proto)
     - [Resources](#akash.base.resources.v1beta4.Resources)
   
 - [akash/bme/v1/types.proto](#akash/bme/v1/types.proto)
     - [BurnMintPair](#akash.bme.v1.BurnMintPair)
     - [CoinPrice](#akash.bme.v1.CoinPrice)
     - [CollateralRatio](#akash.bme.v1.CollateralRatio)
     - [LedgerCanceledRecord](#akash.bme.v1.LedgerCanceledRecord)
     - [LedgerID](#akash.bme.v1.LedgerID)
     - [LedgerPendingRecord](#akash.bme.v1.LedgerPendingRecord)
     - [LedgerRecord](#akash.bme.v1.LedgerRecord)
     - [LedgerRecordID](#akash.bme.v1.LedgerRecordID)
     - [State](#akash.bme.v1.State)
     - [Status](#akash.bme.v1.Status)
   
     - [LedgerCanceledRecord.BMCancelReason](#akash.bme.v1.LedgerCanceledRecord.BMCancelReason)
     - [LedgerRecordStatus](#akash.bme.v1.LedgerRecordStatus)
     - [MintStatus](#akash.bme.v1.MintStatus)
   
 - [akash/bme/v1/events.proto](#akash/bme/v1/events.proto)
     - [EventLedgerRecordCanceled](#akash.bme.v1.EventLedgerRecordCanceled)
     - [EventLedgerRecordExecuted](#akash.bme.v1.EventLedgerRecordExecuted)
     - [EventMintStatusChange](#akash.bme.v1.EventMintStatusChange)
     - [EventVaultFunded](#akash.bme.v1.EventVaultFunded)
   
 - [akash/bme/v1/filters.proto](#akash/bme/v1/filters.proto)
     - [LedgerRecordFilters](#akash.bme.v1.LedgerRecordFilters)
   
 - [akash/bme/v1/params.proto](#akash/bme/v1/params.proto)
     - [Params](#akash.bme.v1.Params)
   
 - [akash/bme/v1/genesis.proto](#akash/bme/v1/genesis.proto)
     - [GenesisLedgerPendingRecord](#akash.bme.v1.GenesisLedgerPendingRecord)
     - [GenesisLedgerRecord](#akash.bme.v1.GenesisLedgerRecord)
     - [GenesisLedgerState](#akash.bme.v1.GenesisLedgerState)
     - [GenesisState](#akash.bme.v1.GenesisState)
     - [GenesisVaultState](#akash.bme.v1.GenesisVaultState)
   
 - [akash/bme/v1/msgs.proto](#akash/bme/v1/msgs.proto)
     - [MsgBurnACT](#akash.bme.v1.MsgBurnACT)
     - [MsgBurnACTResponse](#akash.bme.v1.MsgBurnACTResponse)
     - [MsgBurnMint](#akash.bme.v1.MsgBurnMint)
     - [MsgBurnMintResponse](#akash.bme.v1.MsgBurnMintResponse)
     - [MsgFundVault](#akash.bme.v1.MsgFundVault)
     - [MsgFundVaultResponse](#akash.bme.v1.MsgFundVaultResponse)
     - [MsgMintACT](#akash.bme.v1.MsgMintACT)
     - [MsgMintACTResponse](#akash.bme.v1.MsgMintACTResponse)
     - [MsgUpdateParams](#akash.bme.v1.MsgUpdateParams)
     - [MsgUpdateParamsResponse](#akash.bme.v1.MsgUpdateParamsResponse)
   
 - [akash/bme/v1/query.proto](#akash/bme/v1/query.proto)
     - [QueryLedgerRecordEntry](#akash.bme.v1.QueryLedgerRecordEntry)
     - [QueryLedgerRecordsRequest](#akash.bme.v1.QueryLedgerRecordsRequest)
     - [QueryLedgerRecordsResponse](#akash.bme.v1.QueryLedgerRecordsResponse)
     - [QueryParamsRequest](#akash.bme.v1.QueryParamsRequest)
     - [QueryParamsResponse](#akash.bme.v1.QueryParamsResponse)
     - [QueryStatusRequest](#akash.bme.v1.QueryStatusRequest)
     - [QueryStatusResponse](#akash.bme.v1.QueryStatusResponse)
     - [QueryVaultStateRequest](#akash.bme.v1.QueryVaultStateRequest)
     - [QueryVaultStateResponse](#akash.bme.v1.QueryVaultStateResponse)
   
     - [Query](#akash.bme.v1.Query)
   
 - [akash/bme/v1/service.proto](#akash/bme/v1/service.proto)
     - [Msg](#akash.bme.v1.Msg)
   
 - [akash/cert/v1/cert.proto](#akash/cert/v1/cert.proto)
     - [Certificate](#akash.cert.v1.Certificate)
     - [ID](#akash.cert.v1.ID)
   
     - [State](#akash.cert.v1.State)
   
 - [akash/cert/v1/filters.proto](#akash/cert/v1/filters.proto)
     - [CertificateFilter](#akash.cert.v1.CertificateFilter)
   
 - [akash/cert/v1/genesis.proto](#akash/cert/v1/genesis.proto)
     - [GenesisCertificate](#akash.cert.v1.GenesisCertificate)
     - [GenesisState](#akash.cert.v1.GenesisState)
   
 - [akash/cert/v1/msg.proto](#akash/cert/v1/msg.proto)
     - [MsgCreateCertificate](#akash.cert.v1.MsgCreateCertificate)
     - [MsgCreateCertificateResponse](#akash.cert.v1.MsgCreateCertificateResponse)
     - [MsgRevokeCertificate](#akash.cert.v1.MsgRevokeCertificate)
     - [MsgRevokeCertificateResponse](#akash.cert.v1.MsgRevokeCertificateResponse)
   
 - [akash/cert/v1/query.proto](#akash/cert/v1/query.proto)
     - [CertificateResponse](#akash.cert.v1.CertificateResponse)
     - [QueryCertificatesRequest](#akash.cert.v1.QueryCertificatesRequest)
     - [QueryCertificatesResponse](#akash.cert.v1.QueryCertificatesResponse)
   
     - [Query](#akash.cert.v1.Query)
   
 - [akash/cert/v1/service.proto](#akash/cert/v1/service.proto)
     - [Msg](#akash.cert.v1.Msg)
   
 - [akash/deployment/v1/deployment.proto](#akash/deployment/v1/deployment.proto)
     - [Deployment](#akash.deployment.v1.Deployment)
     - [DeploymentID](#akash.deployment.v1.DeploymentID)
     - [DeploymentReclamation](#akash.deployment.v1.DeploymentReclamation)
   
     - [Deployment.State](#akash.deployment.v1.Deployment.State)
   
 - [akash/deployment/v1/group.proto](#akash/deployment/v1/group.proto)
     - [GroupID](#akash.deployment.v1.GroupID)
   
 - [akash/deployment/v1/event.proto](#akash/deployment/v1/event.proto)
     - [EventDeploymentClosed](#akash.deployment.v1.EventDeploymentClosed)
     - [EventDeploymentCreated](#akash.deployment.v1.EventDeploymentCreated)
     - [EventDeploymentUpdated](#akash.deployment.v1.EventDeploymentUpdated)
     - [EventGroupClosed](#akash.deployment.v1.EventGroupClosed)
     - [EventGroupPaused](#akash.deployment.v1.EventGroupPaused)
     - [EventGroupStarted](#akash.deployment.v1.EventGroupStarted)
   
 - [akash/deployment/v1beta4/resourceunit.proto](#akash/deployment/v1beta4/resourceunit.proto)
     - [ResourceUnit](#akash.deployment.v1beta4.ResourceUnit)
   
 - [akash/deployment/v1beta4/groupspec.proto](#akash/deployment/v1beta4/groupspec.proto)
     - [GroupSpec](#akash.deployment.v1beta4.GroupSpec)
   
 - [akash/deployment/v1beta4/deploymentmsg.proto](#akash/deployment/v1beta4/deploymentmsg.proto)
     - [MsgCloseDeployment](#akash.deployment.v1beta4.MsgCloseDeployment)
     - [MsgCloseDeploymentResponse](#akash.deployment.v1beta4.MsgCloseDeploymentResponse)
     - [MsgCreateDeployment](#akash.deployment.v1beta4.MsgCreateDeployment)
     - [MsgCreateDeploymentResponse](#akash.deployment.v1beta4.MsgCreateDeploymentResponse)
     - [MsgUpdateDeployment](#akash.deployment.v1beta4.MsgUpdateDeployment)
     - [MsgUpdateDeploymentResponse](#akash.deployment.v1beta4.MsgUpdateDeploymentResponse)
   
 - [akash/deployment/v1beta4/filters.proto](#akash/deployment/v1beta4/filters.proto)
     - [DeploymentFilters](#akash.deployment.v1beta4.DeploymentFilters)
     - [GroupFilters](#akash.deployment.v1beta4.GroupFilters)
   
 - [akash/deployment/v1beta4/group.proto](#akash/deployment/v1beta4/group.proto)
     - [Group](#akash.deployment.v1beta4.Group)
   
     - [Group.State](#akash.deployment.v1beta4.Group.State)
   
 - [akash/deployment/v1beta4/params.proto](#akash/deployment/v1beta4/params.proto)
     - [Params](#akash.deployment.v1beta4.Params)
   
 - [akash/deployment/v1beta4/genesis.proto](#akash/deployment/v1beta4/genesis.proto)
     - [GenesisDeployment](#akash.deployment.v1beta4.GenesisDeployment)
     - [GenesisState](#akash.deployment.v1beta4.GenesisState)
   
 - [akash/deployment/v1beta4/groupmsg.proto](#akash/deployment/v1beta4/groupmsg.proto)
     - [MsgCloseGroup](#akash.deployment.v1beta4.MsgCloseGroup)
     - [MsgCloseGroupResponse](#akash.deployment.v1beta4.MsgCloseGroupResponse)
     - [MsgPauseGroup](#akash.deployment.v1beta4.MsgPauseGroup)
     - [MsgPauseGroupResponse](#akash.deployment.v1beta4.MsgPauseGroupResponse)
     - [MsgStartGroup](#akash.deployment.v1beta4.MsgStartGroup)
     - [MsgStartGroupResponse](#akash.deployment.v1beta4.MsgStartGroupResponse)
   
 - [akash/deployment/v1beta4/paramsmsg.proto](#akash/deployment/v1beta4/paramsmsg.proto)
     - [MsgUpdateParams](#akash.deployment.v1beta4.MsgUpdateParams)
     - [MsgUpdateParamsResponse](#akash.deployment.v1beta4.MsgUpdateParamsResponse)
   
 - [akash/escrow/id/v1/id.proto](#akash/escrow/id/v1/id.proto)
     - [Account](#akash.escrow.id.v1.Account)
     - [Payment](#akash.escrow.id.v1.Payment)
   
     - [Scope](#akash.escrow.id.v1.Scope)
   
 - [akash/escrow/types/v1/balance.proto](#akash/escrow/types/v1/balance.proto)
     - [Balance](#akash.escrow.types.v1.Balance)
   
 - [akash/escrow/types/v1/deposit.proto](#akash/escrow/types/v1/deposit.proto)
     - [Depositor](#akash.escrow.types.v1.Depositor)
   
 - [akash/escrow/types/v1/state.proto](#akash/escrow/types/v1/state.proto)
     - [State](#akash.escrow.types.v1.State)
   
 - [akash/escrow/types/v1/account.proto](#akash/escrow/types/v1/account.proto)
     - [Account](#akash.escrow.types.v1.Account)
     - [AccountState](#akash.escrow.types.v1.AccountState)
   
 - [akash/deployment/v1beta4/query.proto](#akash/deployment/v1beta4/query.proto)
     - [QueryDeploymentRequest](#akash.deployment.v1beta4.QueryDeploymentRequest)
     - [QueryDeploymentResponse](#akash.deployment.v1beta4.QueryDeploymentResponse)
     - [QueryDeploymentsRequest](#akash.deployment.v1beta4.QueryDeploymentsRequest)
     - [QueryDeploymentsResponse](#akash.deployment.v1beta4.QueryDeploymentsResponse)
     - [QueryGroupRequest](#akash.deployment.v1beta4.QueryGroupRequest)
     - [QueryGroupResponse](#akash.deployment.v1beta4.QueryGroupResponse)
     - [QueryParamsRequest](#akash.deployment.v1beta4.QueryParamsRequest)
     - [QueryParamsResponse](#akash.deployment.v1beta4.QueryParamsResponse)
   
     - [Query](#akash.deployment.v1beta4.Query)
   
 - [akash/deployment/v1beta4/service.proto](#akash/deployment/v1beta4/service.proto)
     - [Msg](#akash.deployment.v1beta4.Msg)
   
 - [akash/discovery/v1/client_info.proto](#akash/discovery/v1/client_info.proto)
     - [ClientInfo](#akash.discovery.v1.ClientInfo)
     - [ModuleVersion](#akash.discovery.v1.ModuleVersion)
     - [VersionInfo](#akash.discovery.v1.VersionInfo)
   
 - [akash/discovery/v1/akash.proto](#akash/discovery/v1/akash.proto)
     - [Akash](#akash.discovery.v1.Akash)
   
 - [akash/discovery/v1/service.proto](#akash/discovery/v1/service.proto)
     - [GetInfoRequest](#akash.discovery.v1.GetInfoRequest)
     - [GetInfoResponse](#akash.discovery.v1.GetInfoResponse)
   
     - [Discovery](#akash.discovery.v1.Discovery)
   
 - [akash/downtimedetector/v1beta1/downtime_duration.proto](#akash/downtimedetector/v1beta1/downtime_duration.proto)
     - [Downtime](#akash.downtimedetector.v1beta1.Downtime)
   
 - [akash/downtimedetector/v1beta1/genesis.proto](#akash/downtimedetector/v1beta1/genesis.proto)
     - [GenesisDowntimeEntry](#akash.downtimedetector.v1beta1.GenesisDowntimeEntry)
     - [GenesisState](#akash.downtimedetector.v1beta1.GenesisState)
   
 - [akash/downtimedetector/v1beta1/query.proto](#akash/downtimedetector/v1beta1/query.proto)
     - [RecoveredSinceDowntimeOfLengthRequest](#akash.downtimedetector.v1beta1.RecoveredSinceDowntimeOfLengthRequest)
     - [RecoveredSinceDowntimeOfLengthResponse](#akash.downtimedetector.v1beta1.RecoveredSinceDowntimeOfLengthResponse)
   
     - [Query](#akash.downtimedetector.v1beta1.Query)
   
 - [akash/epochs/v1beta1/events.proto](#akash/epochs/v1beta1/events.proto)
     - [EventEpochEnd](#akash.epochs.v1beta1.EventEpochEnd)
     - [EventEpochStart](#akash.epochs.v1beta1.EventEpochStart)
   
 - [akash/epochs/v1beta1/genesis.proto](#akash/epochs/v1beta1/genesis.proto)
     - [EpochInfo](#akash.epochs.v1beta1.EpochInfo)
     - [GenesisState](#akash.epochs.v1beta1.GenesisState)
   
 - [akash/epochs/v1beta1/query.proto](#akash/epochs/v1beta1/query.proto)
     - [QueryCurrentEpochRequest](#akash.epochs.v1beta1.QueryCurrentEpochRequest)
     - [QueryCurrentEpochResponse](#akash.epochs.v1beta1.QueryCurrentEpochResponse)
     - [QueryEpochInfosRequest](#akash.epochs.v1beta1.QueryEpochInfosRequest)
     - [QueryEpochInfosResponse](#akash.epochs.v1beta1.QueryEpochInfosResponse)
   
     - [Query](#akash.epochs.v1beta1.Query)
   
 - [akash/escrow/types/v1/payment.proto](#akash/escrow/types/v1/payment.proto)
     - [Payment](#akash.escrow.types.v1.Payment)
     - [PaymentState](#akash.escrow.types.v1.PaymentState)
   
 - [akash/escrow/v1/authz.proto](#akash/escrow/v1/authz.proto)
     - [DepositAuthorization](#akash.escrow.v1.DepositAuthorization)
   
     - [DepositAuthorization.Scope](#akash.escrow.v1.DepositAuthorization.Scope)
   
 - [akash/escrow/v1/genesis.proto](#akash/escrow/v1/genesis.proto)
     - [GenesisState](#akash.escrow.v1.GenesisState)
   
 - [akash/escrow/v1/msg.proto](#akash/escrow/v1/msg.proto)
     - [MsgAccountDeposit](#akash.escrow.v1.MsgAccountDeposit)
     - [MsgAccountDepositResponse](#akash.escrow.v1.MsgAccountDepositResponse)
   
 - [akash/escrow/v1/query.proto](#akash/escrow/v1/query.proto)
     - [QueryAccountsRequest](#akash.escrow.v1.QueryAccountsRequest)
     - [QueryAccountsResponse](#akash.escrow.v1.QueryAccountsResponse)
     - [QueryPaymentsRequest](#akash.escrow.v1.QueryPaymentsRequest)
     - [QueryPaymentsResponse](#akash.escrow.v1.QueryPaymentsResponse)
   
     - [Query](#akash.escrow.v1.Query)
   
 - [akash/escrow/v1/service.proto](#akash/escrow/v1/service.proto)
     - [Msg](#akash.escrow.v1.Msg)
   
 - [akash/market/v1/bid.proto](#akash/market/v1/bid.proto)
     - [BidID](#akash.market.v1.BidID)
   
 - [akash/market/v1/order.proto](#akash/market/v1/order.proto)
     - [OrderID](#akash.market.v1.OrderID)
   
 - [akash/market/v1/types.proto](#akash/market/v1/types.proto)
     - [LeaseClosedReason](#akash.market.v1.LeaseClosedReason)
   
 - [akash/market/v1/reclamation.proto](#akash/market/v1/reclamation.proto)
     - [Reclamation](#akash.market.v1.Reclamation)
   
 - [akash/market/v1/lease.proto](#akash/market/v1/lease.proto)
     - [Lease](#akash.market.v1.Lease)
     - [LeaseID](#akash.market.v1.LeaseID)
   
     - [Lease.State](#akash.market.v1.Lease.State)
   
 - [akash/market/v1/event.proto](#akash/market/v1/event.proto)
     - [EventBidClosed](#akash.market.v1.EventBidClosed)
     - [EventBidCreated](#akash.market.v1.EventBidCreated)
     - [EventLeaseClosed](#akash.market.v1.EventLeaseClosed)
     - [EventLeaseCreated](#akash.market.v1.EventLeaseCreated)
     - [EventLeaseReclaimStarted](#akash.market.v1.EventLeaseReclaimStarted)
     - [EventOrderClosed](#akash.market.v1.EventOrderClosed)
     - [EventOrderCreated](#akash.market.v1.EventOrderCreated)
   
 - [akash/market/v1/filters.proto](#akash/market/v1/filters.proto)
     - [LeaseFilters](#akash.market.v1.LeaseFilters)
   
 - [akash/market/v1/stats.proto](#akash/market/v1/stats.proto)
     - [ProviderLeaseStats](#akash.market.v1.ProviderLeaseStats)
     - [ProviderLeaseStatsByReason](#akash.market.v1.ProviderLeaseStatsByReason)
   
 - [akash/market/v1beta5/resourcesoffer.proto](#akash/market/v1beta5/resourcesoffer.proto)
     - [EndpointOfferPrice](#akash.market.v1beta5.EndpointOfferPrice)
     - [OfferPrices](#akash.market.v1beta5.OfferPrices)
     - [ResourceOffer](#akash.market.v1beta5.ResourceOffer)
     - [StorageOfferPrice](#akash.market.v1beta5.StorageOfferPrice)
   
 - [akash/market/v1beta5/bid.proto](#akash/market/v1beta5/bid.proto)
     - [Bid](#akash.market.v1beta5.Bid)
   
     - [Bid.State](#akash.market.v1beta5.Bid.State)
   
 - [akash/market/v1beta5/bidmsg.proto](#akash/market/v1beta5/bidmsg.proto)
     - [MsgCloseBid](#akash.market.v1beta5.MsgCloseBid)
     - [MsgCloseBidResponse](#akash.market.v1beta5.MsgCloseBidResponse)
     - [MsgCreateBid](#akash.market.v1beta5.MsgCreateBid)
     - [MsgCreateBidResponse](#akash.market.v1beta5.MsgCreateBidResponse)
   
 - [akash/market/v1beta5/filters.proto](#akash/market/v1beta5/filters.proto)
     - [BidFilters](#akash.market.v1beta5.BidFilters)
     - [OrderFilters](#akash.market.v1beta5.OrderFilters)
   
 - [akash/market/v1beta5/params.proto](#akash/market/v1beta5/params.proto)
     - [Params](#akash.market.v1beta5.Params)
   
 - [akash/market/v1beta5/order.proto](#akash/market/v1beta5/order.proto)
     - [Order](#akash.market.v1beta5.Order)
   
     - [Order.State](#akash.market.v1beta5.Order.State)
   
 - [akash/market/v1beta5/genesis.proto](#akash/market/v1beta5/genesis.proto)
     - [GenesisState](#akash.market.v1beta5.GenesisState)
   
 - [akash/market/v1beta5/leasemsg.proto](#akash/market/v1beta5/leasemsg.proto)
     - [MsgCloseLease](#akash.market.v1beta5.MsgCloseLease)
     - [MsgCloseLeaseResponse](#akash.market.v1beta5.MsgCloseLeaseResponse)
     - [MsgCreateLease](#akash.market.v1beta5.MsgCreateLease)
     - [MsgCreateLeaseResponse](#akash.market.v1beta5.MsgCreateLeaseResponse)
     - [MsgLeaseStartReclaim](#akash.market.v1beta5.MsgLeaseStartReclaim)
     - [MsgLeaseStartReclaimResponse](#akash.market.v1beta5.MsgLeaseStartReclaimResponse)
     - [MsgWithdrawLease](#akash.market.v1beta5.MsgWithdrawLease)
     - [MsgWithdrawLeaseResponse](#akash.market.v1beta5.MsgWithdrawLeaseResponse)
   
 - [akash/market/v1beta5/paramsmsg.proto](#akash/market/v1beta5/paramsmsg.proto)
     - [MsgUpdateParams](#akash.market.v1beta5.MsgUpdateParams)
     - [MsgUpdateParamsResponse](#akash.market.v1beta5.MsgUpdateParamsResponse)
   
 - [akash/market/v1beta5/query.proto](#akash/market/v1beta5/query.proto)
     - [QueryBidRequest](#akash.market.v1beta5.QueryBidRequest)
     - [QueryBidResponse](#akash.market.v1beta5.QueryBidResponse)
     - [QueryBidsRequest](#akash.market.v1beta5.QueryBidsRequest)
     - [QueryBidsResponse](#akash.market.v1beta5.QueryBidsResponse)
     - [QueryLeaseRequest](#akash.market.v1beta5.QueryLeaseRequest)
     - [QueryLeaseResponse](#akash.market.v1beta5.QueryLeaseResponse)
     - [QueryLeasesRequest](#akash.market.v1beta5.QueryLeasesRequest)
     - [QueryLeasesResponse](#akash.market.v1beta5.QueryLeasesResponse)
     - [QueryOrderRequest](#akash.market.v1beta5.QueryOrderRequest)
     - [QueryOrderResponse](#akash.market.v1beta5.QueryOrderResponse)
     - [QueryOrdersRequest](#akash.market.v1beta5.QueryOrdersRequest)
     - [QueryOrdersResponse](#akash.market.v1beta5.QueryOrdersResponse)
     - [QueryParamsRequest](#akash.market.v1beta5.QueryParamsRequest)
     - [QueryParamsResponse](#akash.market.v1beta5.QueryParamsResponse)
     - [QueryProviderLeaseStatsRequest](#akash.market.v1beta5.QueryProviderLeaseStatsRequest)
     - [QueryProviderLeaseStatsResponse](#akash.market.v1beta5.QueryProviderLeaseStatsResponse)
   
     - [Query](#akash.market.v1beta5.Query)
   
 - [akash/market/v1beta5/service.proto](#akash/market/v1beta5/service.proto)
     - [Msg](#akash.market.v1beta5.Msg)
   
 - [akash/oracle/v1/prices.proto](#akash/oracle/v1/prices.proto)
     - [AggregatedPrice](#akash.oracle.v1.AggregatedPrice)
     - [DataID](#akash.oracle.v1.DataID)
     - [PriceData](#akash.oracle.v1.PriceData)
     - [PriceDataID](#akash.oracle.v1.PriceDataID)
     - [PriceDataRecordID](#akash.oracle.v1.PriceDataRecordID)
     - [PriceDataState](#akash.oracle.v1.PriceDataState)
     - [PriceHealth](#akash.oracle.v1.PriceHealth)
     - [PricesFilter](#akash.oracle.v1.PricesFilter)
     - [QueryPricesRequest](#akash.oracle.v1.QueryPricesRequest)
     - [QueryPricesResponse](#akash.oracle.v1.QueryPricesResponse)
   
 - [akash/oracle/v1/events.proto](#akash/oracle/v1/events.proto)
     - [EventAggregatedPrice](#akash.oracle.v1.EventAggregatedPrice)
     - [EventPriceData](#akash.oracle.v1.EventPriceData)
     - [EventPriceRecovered](#akash.oracle.v1.EventPriceRecovered)
     - [EventPriceStaleWarning](#akash.oracle.v1.EventPriceStaleWarning)
     - [EventPriceStaled](#akash.oracle.v1.EventPriceStaled)
   
 - [akash/oracle/v1/params.proto](#akash/oracle/v1/params.proto)
     - [Params](#akash.oracle.v1.Params)
   
 - [akash/oracle/v1/genesis.proto](#akash/oracle/v1/genesis.proto)
     - [GenesisState](#akash.oracle.v1.GenesisState)
   
 - [akash/oracle/v1/msgs.proto](#akash/oracle/v1/msgs.proto)
     - [MsgAddPriceEntry](#akash.oracle.v1.MsgAddPriceEntry)
     - [MsgAddPriceEntryResponse](#akash.oracle.v1.MsgAddPriceEntryResponse)
     - [MsgUpdateParams](#akash.oracle.v1.MsgUpdateParams)
     - [MsgUpdateParamsResponse](#akash.oracle.v1.MsgUpdateParamsResponse)
   
 - [akash/oracle/v1/query.proto](#akash/oracle/v1/query.proto)
     - [QueryAggregatedPriceRequest](#akash.oracle.v1.QueryAggregatedPriceRequest)
     - [QueryAggregatedPriceResponse](#akash.oracle.v1.QueryAggregatedPriceResponse)
     - [QueryParamsRequest](#akash.oracle.v1.QueryParamsRequest)
     - [QueryParamsResponse](#akash.oracle.v1.QueryParamsResponse)
   
     - [Query](#akash.oracle.v1.Query)
   
 - [akash/oracle/v1/service.proto](#akash/oracle/v1/service.proto)
     - [Msg](#akash.oracle.v1.Msg)
   
 - [akash/oracle/v2/prices.proto](#akash/oracle/v2/prices.proto)
     - [AggregatedPrice](#akash.oracle.v2.AggregatedPrice)
     - [DataID](#akash.oracle.v2.DataID)
     - [PriceData](#akash.oracle.v2.PriceData)
     - [PriceDataID](#akash.oracle.v2.PriceDataID)
     - [PriceDataRecordID](#akash.oracle.v2.PriceDataRecordID)
     - [PriceDataState](#akash.oracle.v2.PriceDataState)
     - [PriceHealth](#akash.oracle.v2.PriceHealth)
     - [PriceLatestDataState](#akash.oracle.v2.PriceLatestDataState)
   
 - [akash/oracle/v2/events.proto](#akash/oracle/v2/events.proto)
     - [EventAggregatedPrice](#akash.oracle.v2.EventAggregatedPrice)
     - [EventPriceData](#akash.oracle.v2.EventPriceData)
     - [EventPriceRecovered](#akash.oracle.v2.EventPriceRecovered)
     - [EventPriceStaleWarning](#akash.oracle.v2.EventPriceStaleWarning)
     - [EventPriceStaled](#akash.oracle.v2.EventPriceStaled)
   
 - [akash/oracle/v2/params.proto](#akash/oracle/v2/params.proto)
     - [Params](#akash.oracle.v2.Params)
   
 - [akash/oracle/v2/genesis.proto](#akash/oracle/v2/genesis.proto)
     - [GenesisLatestPricesIDs](#akash.oracle.v2.GenesisLatestPricesIDs)
     - [GenesisSourceID](#akash.oracle.v2.GenesisSourceID)
     - [GenesisState](#akash.oracle.v2.GenesisState)
   
 - [akash/oracle/v2/msgs.proto](#akash/oracle/v2/msgs.proto)
     - [MsgAddPriceEntry](#akash.oracle.v2.MsgAddPriceEntry)
     - [MsgAddPriceEntryResponse](#akash.oracle.v2.MsgAddPriceEntryResponse)
     - [MsgUpdateParams](#akash.oracle.v2.MsgUpdateParams)
     - [MsgUpdateParamsResponse](#akash.oracle.v2.MsgUpdateParamsResponse)
   
 - [akash/oracle/v2/query.proto](#akash/oracle/v2/query.proto)
     - [PricesFilter](#akash.oracle.v2.PricesFilter)
     - [QueryAggregatedPriceRequest](#akash.oracle.v2.QueryAggregatedPriceRequest)
     - [QueryAggregatedPriceResponse](#akash.oracle.v2.QueryAggregatedPriceResponse)
     - [QueryParamsRequest](#akash.oracle.v2.QueryParamsRequest)
     - [QueryParamsResponse](#akash.oracle.v2.QueryParamsResponse)
     - [QueryPricesRequest](#akash.oracle.v2.QueryPricesRequest)
     - [QueryPricesResponse](#akash.oracle.v2.QueryPricesResponse)
   
     - [Query](#akash.oracle.v2.Query)
   
 - [akash/oracle/v2/service.proto](#akash/oracle/v2/service.proto)
     - [Msg](#akash.oracle.v2.Msg)
   
 - [akash/provider/v1beta4/maintenance.proto](#akash/provider/v1beta4/maintenance.proto)
     - [ProviderMaintenanceRecord](#akash.provider.v1beta4.ProviderMaintenanceRecord)
     - [ProviderMaintenanceWithStatus](#akash.provider.v1beta4.ProviderMaintenanceWithStatus)
   
     - [ProviderMaintenanceStatus](#akash.provider.v1beta4.ProviderMaintenanceStatus)
     - [ProviderMaintenanceType](#akash.provider.v1beta4.ProviderMaintenanceType)
   
 - [akash/provider/v1beta4/event.proto](#akash/provider/v1beta4/event.proto)
     - [EventProviderCreated](#akash.provider.v1beta4.EventProviderCreated)
     - [EventProviderDeleted](#akash.provider.v1beta4.EventProviderDeleted)
     - [EventProviderMaintenanceClosed](#akash.provider.v1beta4.EventProviderMaintenanceClosed)
     - [EventProviderMaintenanceOpened](#akash.provider.v1beta4.EventProviderMaintenanceOpened)
     - [EventProviderUpdated](#akash.provider.v1beta4.EventProviderUpdated)
   
 - [akash/provider/v1beta4/provider.proto](#akash/provider/v1beta4/provider.proto)
     - [Info](#akash.provider.v1beta4.Info)
     - [Provider](#akash.provider.v1beta4.Provider)
     - [ProviderRegistration](#akash.provider.v1beta4.ProviderRegistration)
   
 - [akash/provider/v1beta4/params.proto](#akash/provider/v1beta4/params.proto)
     - [ProviderMaintenanceParams](#akash.provider.v1beta4.ProviderMaintenanceParams)
   
 - [akash/provider/v1beta4/genesis.proto](#akash/provider/v1beta4/genesis.proto)
     - [GenesisState](#akash.provider.v1beta4.GenesisState)
   
 - [akash/provider/v1beta4/msg.proto](#akash/provider/v1beta4/msg.proto)
     - [MsgCloseProviderMaintenance](#akash.provider.v1beta4.MsgCloseProviderMaintenance)
     - [MsgCloseProviderMaintenanceResponse](#akash.provider.v1beta4.MsgCloseProviderMaintenanceResponse)
     - [MsgCreateProvider](#akash.provider.v1beta4.MsgCreateProvider)
     - [MsgCreateProviderResponse](#akash.provider.v1beta4.MsgCreateProviderResponse)
     - [MsgDeleteProvider](#akash.provider.v1beta4.MsgDeleteProvider)
     - [MsgDeleteProviderResponse](#akash.provider.v1beta4.MsgDeleteProviderResponse)
     - [MsgOpenProviderMaintenance](#akash.provider.v1beta4.MsgOpenProviderMaintenance)
     - [MsgOpenProviderMaintenanceResponse](#akash.provider.v1beta4.MsgOpenProviderMaintenanceResponse)
     - [MsgUpdateProvider](#akash.provider.v1beta4.MsgUpdateProvider)
     - [MsgUpdateProviderResponse](#akash.provider.v1beta4.MsgUpdateProviderResponse)
   
 - [akash/provider/v1beta4/paramsmsg.proto](#akash/provider/v1beta4/paramsmsg.proto)
     - [MsgUpdateParams](#akash.provider.v1beta4.MsgUpdateParams)
     - [MsgUpdateParamsResponse](#akash.provider.v1beta4.MsgUpdateParamsResponse)
   
 - [akash/provider/v1beta4/query.proto](#akash/provider/v1beta4/query.proto)
     - [QueryParamsRequest](#akash.provider.v1beta4.QueryParamsRequest)
     - [QueryParamsResponse](#akash.provider.v1beta4.QueryParamsResponse)
     - [QueryProviderMaintenanceRequest](#akash.provider.v1beta4.QueryProviderMaintenanceRequest)
     - [QueryProviderMaintenanceResponse](#akash.provider.v1beta4.QueryProviderMaintenanceResponse)
     - [QueryProviderMaintenancesRequest](#akash.provider.v1beta4.QueryProviderMaintenancesRequest)
     - [QueryProviderMaintenancesResponse](#akash.provider.v1beta4.QueryProviderMaintenancesResponse)
     - [QueryProviderRequest](#akash.provider.v1beta4.QueryProviderRequest)
     - [QueryProviderResponse](#akash.provider.v1beta4.QueryProviderResponse)
     - [QueryProvidersRequest](#akash.provider.v1beta4.QueryProvidersRequest)
     - [QueryProvidersResponse](#akash.provider.v1beta4.QueryProvidersResponse)
     - [QueryRegistrationRequest](#akash.provider.v1beta4.QueryRegistrationRequest)
     - [QueryRegistrationResponse](#akash.provider.v1beta4.QueryRegistrationResponse)
   
     - [Query](#akash.provider.v1beta4.Query)
   
 - [akash/provider/v1beta4/service.proto](#akash/provider/v1beta4/service.proto)
     - [Msg](#akash.provider.v1beta4.Msg)
   
 - [akash/take/v1/params.proto](#akash/take/v1/params.proto)
     - [DenomTakeRate](#akash.take.v1.DenomTakeRate)
     - [Params](#akash.take.v1.Params)
   
 - [akash/take/v1/genesis.proto](#akash/take/v1/genesis.proto)
     - [GenesisState](#akash.take.v1.GenesisState)
   
 - [akash/take/v1/paramsmsg.proto](#akash/take/v1/paramsmsg.proto)
     - [MsgUpdateParams](#akash.take.v1.MsgUpdateParams)
     - [MsgUpdateParamsResponse](#akash.take.v1.MsgUpdateParamsResponse)
   
 - [akash/take/v1/query.proto](#akash/take/v1/query.proto)
     - [QueryParamsRequest](#akash.take.v1.QueryParamsRequest)
     - [QueryParamsResponse](#akash.take.v1.QueryParamsResponse)
   
     - [Query](#akash.take.v1.Query)
   
 - [akash/take/v1/service.proto](#akash/take/v1/service.proto)
     - [Msg](#akash.take.v1.Msg)
   
 - [akash/verification/v1/event.proto](#akash/verification/v1/event.proto)
     - [EventAttestationExpired](#akash.verification.v1.EventAttestationExpired)
     - [EventAttestationReplaced](#akash.verification.v1.EventAttestationReplaced)
     - [EventAttestationRevoked](#akash.verification.v1.EventAttestationRevoked)
     - [EventAttestationSubmitted](#akash.verification.v1.EventAttestationSubmitted)
     - [EventAttestationVoided](#akash.verification.v1.EventAttestationVoided)
     - [EventAuditEscrowOpened](#akash.verification.v1.EventAuditEscrowOpened)
     - [EventAuditEscrowSettled](#akash.verification.v1.EventAuditEscrowSettled)
     - [EventAuditorBondPosted](#akash.verification.v1.EventAuditorBondPosted)
     - [EventAuditorFrozen](#akash.verification.v1.EventAuditorFrozen)
     - [EventAuditorLapsed](#akash.verification.v1.EventAuditorLapsed)
     - [EventAuditorRegistered](#akash.verification.v1.EventAuditorRegistered)
     - [EventAuditorRemoved](#akash.verification.v1.EventAuditorRemoved)
     - [EventAuditorRenewed](#akash.verification.v1.EventAuditorRenewed)
     - [EventAuditorResigned](#akash.verification.v1.EventAuditorResigned)
     - [EventDepositReturnedToAuditor](#akash.verification.v1.EventDepositReturnedToAuditor)
     - [EventDepositSlashed](#akash.verification.v1.EventDepositSlashed)
     - [EventDiscrepancyDetected](#akash.verification.v1.EventDiscrepancyDetected)
     - [EventDiscrepancyResolved](#akash.verification.v1.EventDiscrepancyResolved)
     - [EventDiscrepancyTimedOut](#akash.verification.v1.EventDiscrepancyTimedOut)
     - [EventFeeEscrowed](#akash.verification.v1.EventFeeEscrowed)
     - [EventFeeReleasedToAuditor](#akash.verification.v1.EventFeeReleasedToAuditor)
     - [EventFeeReturnedToProvider](#akash.verification.v1.EventFeeReturnedToProvider)
     - [EventProviderBondPosted](#akash.verification.v1.EventProviderBondPosted)
     - [EventProviderBondSlashed](#akash.verification.v1.EventProviderBondSlashed)
     - [EventProviderBondWithdrawalCompleted](#akash.verification.v1.EventProviderBondWithdrawalCompleted)
     - [EventProviderBondWithdrawalInitiated](#akash.verification.v1.EventProviderBondWithdrawalInitiated)
     - [EventSnapshotHashPosted](#akash.verification.v1.EventSnapshotHashPosted)
     - [EventSnapshotResumed](#akash.verification.v1.EventSnapshotResumed)
     - [EventSnapshotSuspended](#akash.verification.v1.EventSnapshotSuspended)
     - [EventVerificationGraceEnded](#akash.verification.v1.EventVerificationGraceEnded)
     - [EventVerificationGraceStarted](#akash.verification.v1.EventVerificationGraceStarted)
   
 - [akash/verification/v1/state.proto](#akash/verification/v1/state.proto)
     - [AttestationRecord](#akash.verification.v1.AttestationRecord)
     - [AuditEscrowRecord](#akash.verification.v1.AuditEscrowRecord)
     - [AuditorRecord](#akash.verification.v1.AuditorRecord)
     - [DiscrepancyEvent](#akash.verification.v1.DiscrepancyEvent)
     - [ProviderBondRecord](#akash.verification.v1.ProviderBondRecord)
     - [ProviderSnapshotRecord](#akash.verification.v1.ProviderSnapshotRecord)
     - [ProviderVerificationGraceRecord](#akash.verification.v1.ProviderVerificationGraceRecord)
     - [ResourceSummary](#akash.verification.v1.ResourceSummary)
     - [SoftwareIdentity](#akash.verification.v1.SoftwareIdentity)
     - [UnbondingEntry](#akash.verification.v1.UnbondingEntry)
     - [VerificationStoreRecord](#akash.verification.v1.VerificationStoreRecord)
   
 - [akash/verification/v1/params.proto](#akash/verification/v1/params.proto)
     - [Params](#akash.verification.v1.Params)
   
 - [akash/verification/v1/genesis.proto](#akash/verification/v1/genesis.proto)
     - [GenesisState](#akash.verification.v1.GenesisState)
   
 - [akash/verification/v1/msg.proto](#akash/verification/v1/msg.proto)
     - [MsgCancelAuditEscrow](#akash.verification.v1.MsgCancelAuditEscrow)
     - [MsgCancelAuditEscrowResponse](#akash.verification.v1.MsgCancelAuditEscrowResponse)
     - [MsgOpenAuditEscrow](#akash.verification.v1.MsgOpenAuditEscrow)
     - [MsgOpenAuditEscrowResponse](#akash.verification.v1.MsgOpenAuditEscrowResponse)
     - [MsgPostAuditorBond](#akash.verification.v1.MsgPostAuditorBond)
     - [MsgPostAuditorBondResponse](#akash.verification.v1.MsgPostAuditorBondResponse)
     - [MsgPostProviderBond](#akash.verification.v1.MsgPostProviderBond)
     - [MsgPostProviderBondResponse](#akash.verification.v1.MsgPostProviderBondResponse)
     - [MsgPostSnapshotHash](#akash.verification.v1.MsgPostSnapshotHash)
     - [MsgPostSnapshotHashResponse](#akash.verification.v1.MsgPostSnapshotHashResponse)
     - [MsgRegisterAuditor](#akash.verification.v1.MsgRegisterAuditor)
     - [MsgRegisterAuditorResponse](#akash.verification.v1.MsgRegisterAuditorResponse)
     - [MsgRemoveAttestation](#akash.verification.v1.MsgRemoveAttestation)
     - [MsgRemoveAttestationResponse](#akash.verification.v1.MsgRemoveAttestationResponse)
     - [MsgRemoveAuditor](#akash.verification.v1.MsgRemoveAuditor)
     - [MsgRemoveAuditorResponse](#akash.verification.v1.MsgRemoveAuditorResponse)
     - [MsgRenewAuditor](#akash.verification.v1.MsgRenewAuditor)
     - [MsgRenewAuditorResponse](#akash.verification.v1.MsgRenewAuditorResponse)
     - [MsgResignAuditor](#akash.verification.v1.MsgResignAuditor)
     - [MsgResignAuditorResponse](#akash.verification.v1.MsgResignAuditorResponse)
     - [MsgResolveDiscrepancy](#akash.verification.v1.MsgResolveDiscrepancy)
     - [MsgResolveDiscrepancyResponse](#akash.verification.v1.MsgResolveDiscrepancyResponse)
     - [MsgRevokeAllProviderAttestations](#akash.verification.v1.MsgRevokeAllProviderAttestations)
     - [MsgRevokeAllProviderAttestationsResponse](#akash.verification.v1.MsgRevokeAllProviderAttestationsResponse)
     - [MsgRevokeAttestation](#akash.verification.v1.MsgRevokeAttestation)
     - [MsgRevokeAttestationResponse](#akash.verification.v1.MsgRevokeAttestationResponse)
     - [MsgRevokeAuditorAttestations](#akash.verification.v1.MsgRevokeAuditorAttestations)
     - [MsgRevokeAuditorAttestationsResponse](#akash.verification.v1.MsgRevokeAuditorAttestationsResponse)
     - [MsgRevokeProviderAttestation](#akash.verification.v1.MsgRevokeProviderAttestation)
     - [MsgRevokeProviderAttestationResponse](#akash.verification.v1.MsgRevokeProviderAttestationResponse)
     - [MsgSettleAuditEscrow](#akash.verification.v1.MsgSettleAuditEscrow)
     - [MsgSettleAuditEscrowResponse](#akash.verification.v1.MsgSettleAuditEscrowResponse)
     - [MsgSlashProviderBond](#akash.verification.v1.MsgSlashProviderBond)
     - [MsgSlashProviderBondResponse](#akash.verification.v1.MsgSlashProviderBondResponse)
     - [MsgSubmitAttestation](#akash.verification.v1.MsgSubmitAttestation)
     - [MsgSubmitAttestationResponse](#akash.verification.v1.MsgSubmitAttestationResponse)
     - [MsgUpdateParams](#akash.verification.v1.MsgUpdateParams)
     - [MsgUpdateParamsResponse](#akash.verification.v1.MsgUpdateParamsResponse)
     - [MsgWithdrawProviderBond](#akash.verification.v1.MsgWithdrawProviderBond)
     - [MsgWithdrawProviderBondResponse](#akash.verification.v1.MsgWithdrawProviderBondResponse)
   
 - [akash/verification/v1/query.proto](#akash/verification/v1/query.proto)
     - [QueryAttestationRequest](#akash.verification.v1.QueryAttestationRequest)
     - [QueryAttestationResponse](#akash.verification.v1.QueryAttestationResponse)
     - [QueryAuditEscrowRequest](#akash.verification.v1.QueryAuditEscrowRequest)
     - [QueryAuditEscrowResponse](#akash.verification.v1.QueryAuditEscrowResponse)
     - [QueryAuditorAttestationsRequest](#akash.verification.v1.QueryAuditorAttestationsRequest)
     - [QueryAuditorAttestationsResponse](#akash.verification.v1.QueryAuditorAttestationsResponse)
     - [QueryAuditorRequest](#akash.verification.v1.QueryAuditorRequest)
     - [QueryAuditorResponse](#akash.verification.v1.QueryAuditorResponse)
     - [QueryAuditorsRequest](#akash.verification.v1.QueryAuditorsRequest)
     - [QueryAuditorsResponse](#akash.verification.v1.QueryAuditorsResponse)
     - [QueryDiscrepanciesRequest](#akash.verification.v1.QueryDiscrepanciesRequest)
     - [QueryDiscrepanciesResponse](#akash.verification.v1.QueryDiscrepanciesResponse)
     - [QueryDiscrepancyRequest](#akash.verification.v1.QueryDiscrepancyRequest)
     - [QueryDiscrepancyResponse](#akash.verification.v1.QueryDiscrepancyResponse)
     - [QueryParamsRequest](#akash.verification.v1.QueryParamsRequest)
     - [QueryParamsResponse](#akash.verification.v1.QueryParamsResponse)
     - [QueryProviderAttestationsRequest](#akash.verification.v1.QueryProviderAttestationsRequest)
     - [QueryProviderAttestationsResponse](#akash.verification.v1.QueryProviderAttestationsResponse)
     - [QueryProviderAuditEscrowsRequest](#akash.verification.v1.QueryProviderAuditEscrowsRequest)
     - [QueryProviderAuditEscrowsResponse](#akash.verification.v1.QueryProviderAuditEscrowsResponse)
     - [QueryProviderBondRequest](#akash.verification.v1.QueryProviderBondRequest)
     - [QueryProviderBondResponse](#akash.verification.v1.QueryProviderBondResponse)
     - [QueryProviderSnapshotRequest](#akash.verification.v1.QueryProviderSnapshotRequest)
     - [QueryProviderSnapshotResponse](#akash.verification.v1.QueryProviderSnapshotResponse)
     - [QueryProviderVerificationGraceRequest](#akash.verification.v1.QueryProviderVerificationGraceRequest)
     - [QueryProviderVerificationGraceResponse](#akash.verification.v1.QueryProviderVerificationGraceResponse)
   
     - [Query](#akash.verification.v1.Query)
   
 - [akash/verification/v1/service.proto](#akash/verification/v1/service.proto)
     - [Msg](#akash.verification.v1.Msg)
   
 - [akash/wasm/v1/event.proto](#akash/wasm/v1/event.proto)
     - [EventMsgBlocked](#akash.wasm.v1.EventMsgBlocked)
   
 - [akash/wasm/v1/params.proto](#akash/wasm/v1/params.proto)
     - [Params](#akash.wasm.v1.Params)
   
 - [akash/wasm/v1/genesis.proto](#akash/wasm/v1/genesis.proto)
     - [GenesisState](#akash.wasm.v1.GenesisState)
   
 - [akash/wasm/v1/paramsmsg.proto](#akash/wasm/v1/paramsmsg.proto)
     - [MsgUpdateParams](#akash.wasm.v1.MsgUpdateParams)
     - [MsgUpdateParamsResponse](#akash.wasm.v1.MsgUpdateParamsResponse)
   
 - [akash/wasm/v1/query.proto](#akash/wasm/v1/query.proto)
     - [QueryParamsRequest](#akash.wasm.v1.QueryParamsRequest)
     - [QueryParamsResponse](#akash.wasm.v1.QueryParamsResponse)
   
     - [Query](#akash.wasm.v1.Query)
   
 - [akash/wasm/v1/service.proto](#akash/wasm/v1/service.proto)
     - [Msg](#akash.wasm.v1.Msg)
   
 - [Scalar Value Types](#scalar-value-types)

 
 
 <a name="akash/verification/v1/types.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/verification/v1/types.proto
 

  <!-- end messages -->

 
 <a name="akash.verification.v1.AttestationRevocationReason"></a>

 ### AttestationRevocationReason
 AttestationRevocationReason enumerates the typed reasons an auditor may revoke
an attestation it previously submitted.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | attestation_revocation_reason_unspecified | 0 | AttestationRevocationReasonUnspecified is the default; no reason has been set. |
 | attestation_revocation_reason_provider_no_longer_qualifies | 1 | AttestationRevocationReasonProviderNoLongerQualifies indicates the provider no longer meets the tier criteria. |
 | attestation_revocation_reason_snapshot_mismatch | 2 | AttestationRevocationReasonSnapshotMismatch indicates the provider's snapshot does not match the attested claim. |
 | attestation_revocation_reason_software_identity_changed | 3 | AttestationRevocationReasonSoftwareIdentityChanged indicates the provider's software identity changed. |
 | attestation_revocation_reason_capability_misrepresented | 4 | AttestationRevocationReasonCapabilityMisrepresented indicates the provider misrepresented its capabilities. |
 | attestation_revocation_reason_provider_non_responsive | 5 | AttestationRevocationReasonProviderNonResponsive indicates the provider became non-responsive to the auditor. |
 | attestation_revocation_reason_auditor_evidence_error | 6 | AttestationRevocationReasonAuditorEvidenceError indicates an error in the auditor's evidence forced revocation. |
 | attestation_revocation_reason_auditor_operational_exit | 7 | AttestationRevocationReasonAuditorOperationalExit indicates the auditor is exiting operations. |
 

 
 <a name="akash.verification.v1.AttestationStatus"></a>

 ### AttestationStatus
 AttestationStatus represents the lifecycle status of an attestation record.
Disputed attestations are stored as status=Voided with voided_reason=Discrepancy
and a pending DiscrepancyEvent. Replacement is a transition event before the
(provider, auditor) attestation record is overwritten, not a persisted status.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | attestation_status_unspecified | 0 | AttestationStatusUnspecified is the default; no status has been set. |
 | attestation_status_valid | 1 | AttestationStatusValid indicates the attestation is currently valid. |
 | attestation_status_voided | 2 | AttestationStatusVoided indicates the attestation has been voided (see VoidedReason). |
 | attestation_status_expired | 3 | AttestationStatusExpired indicates the attestation passed its TTL without renewal. |
 | attestation_status_revoked | 4 | AttestationStatusRevoked indicates the attestation was revoked by the auditor. |
 | attestation_status_removed | 5 | AttestationStatusRemoved indicates the attestation was removed by governance. |
 

 
 <a name="akash.verification.v1.AuditEscrowSettlementReason"></a>

 ### AuditEscrowSettlementReason
 AuditEscrowSettlementReason enumerates the reasons an audit escrow may settle.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | audit_escrow_settlement_reason_unspecified | 0 | AuditEscrowSettlementReasonUnspecified is the default; no reason has been set. |
 | audit_escrow_settlement_reason_cancelled_unconsumed | 1 | AuditEscrowSettlementReasonCancelledUnconsumed indicates the escrow was cancelled without an audit being consumed. |
 | audit_escrow_settlement_reason_expired_unconsumed | 2 | AuditEscrowSettlementReasonExpiredUnconsumed indicates the escrow expired without an audit being consumed. |
 | audit_escrow_settlement_reason_provider_fault | 3 | AuditEscrowSettlementReasonProviderFault indicates settlement was for provider fault. |
 | audit_escrow_settlement_reason_no_fault | 4 | AuditEscrowSettlementReasonNoFault indicates settlement without fault attribution. |
 

 
 <a name="akash.verification.v1.AuditEscrowStatus"></a>

 ### AuditEscrowStatus
 AuditEscrowStatus represents the lifecycle status of an audit escrow record.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | audit_escrow_status_unspecified | 0 | AuditEscrowStatusUnspecified is the default; no status has been set. |
 | audit_escrow_status_open | 1 | AuditEscrowStatusOpen indicates the escrow is open and may be consumed by an attestation. |
 | audit_escrow_status_consumed | 2 | AuditEscrowStatusConsumed indicates the escrow has been consumed by an attestation. |
 | audit_escrow_status_cancelled | 3 | AuditEscrowStatusCancelled indicates the escrow was cancelled before consumption. |
 | audit_escrow_status_expired | 4 | AuditEscrowStatusExpired indicates the escrow expired before consumption. |
 | audit_escrow_status_settled | 5 | AuditEscrowStatusSettled indicates the escrow has been finally settled. |
 

 
 <a name="akash.verification.v1.AuditorSelectionMode"></a>

 ### AuditorSelectionMode
 AuditorSelectionMode controls how the `required_auditors` list on a
VerificationRequirement is evaluated against a provider's set of attestations.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | auditor_selection_mode_unspecified | 0 | AuditorSelectionModeUnspecified is the default; treated as Any. |
 | auditor_selection_mode_any | 1 | AuditorSelectionModeAny matches when any one of the required auditors has attested. |
 | auditor_selection_mode_all | 2 | AuditorSelectionModeAll matches only when all required auditors have attested. |
 

 
 <a name="akash.verification.v1.AuditorStatus"></a>

 ### AuditorStatus
 AuditorStatus represents the lifecycle status of an auditor record.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | auditor_status_unspecified | 0 | AuditorStatusUnspecified is the default; no status has been set. |
 | auditor_status_pending_bond | 6 | AuditorStatusPendingBond indicates the auditor has been approved by governance but has not posted enough bond to submit attestations. |
 | auditor_status_active | 1 | AuditorStatusActive indicates the auditor is bonded and may submit attestations. |
 | auditor_status_frozen | 2 | AuditorStatusFrozen indicates the auditor has been temporarily suspended (e.g. discrepancy threshold). |
 | auditor_status_lapsed | 3 | AuditorStatusLapsed indicates the auditor's renewal deadline has passed without renewal. |
 | auditor_status_resigned | 4 | AuditorStatusResigned indicates the auditor voluntarily resigned. |
 | auditor_status_removed | 5 | AuditorStatusRemoved indicates the auditor was removed by governance. |
 

 
 <a name="akash.verification.v1.BondStatus"></a>

 ### BondStatus
 BondStatus represents the status of an auditor's bond.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | bond_status_unspecified | 0 | BondStatusUnspecified is the default; no bond status has been set. |
 | bond_status_not_bonded | 4 | BondStatusNotBonded indicates no active auditor bond is currently posted. |
 | bond_status_bonded | 1 | BondStatusBonded indicates the bond is fully posted and active. |
 | bond_status_frozen | 2 | BondStatusFrozen indicates the bond is locked while the auditor is frozen. |
 | bond_status_unbonding | 3 | BondStatusUnbonding indicates the bond is currently in the unbonding period. |
 

 
 <a name="akash.verification.v1.CapabilityFlag"></a>

 ### CapabilityFlag
 CapabilityFlag enumerates the optional provider capabilities that may be asserted
by attestations and requested by deployments via the SDL.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | capability_unspecified | 0 | CapabilityUnspecified is the default; no capability has been asserted. |
 | capability_tee_hardware_attestation | 1 | CapabilityTEEHardwareAttestation indicates the provider asserts TEE hardware attestation support. |
 | capability_confidential_computing | 2 | CapabilityConfidentialComputing indicates the provider asserts confidential computing support. |
 | capability_persistent_storage | 3 | CapabilityPersistentStorage indicates the provider asserts persistent storage support. |
 | capability_bare_metal | 4 | CapabilityBareMetal indicates the provider asserts bare-metal hosting support. |
 

 
 <a name="akash.verification.v1.DepositStatus"></a>

 ### DepositStatus
 DepositStatus tracks the lifecycle of an auditor's anti-griefing deposit.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | deposit_status_unspecified | 0 | DepositStatusUnspecified is the default; no deposit status has been set. |
 | deposit_status_escrowed | 1 | DepositStatusEscrowed indicates the deposit is held in escrow. |
 | deposit_status_pending_discrepancy | 2 | DepositStatusPendingDiscrepancy indicates the deposit is locked pending discrepancy resolution. |
 | deposit_status_returned_to_auditor | 3 | DepositStatusReturnedToAuditor indicates the deposit was returned to the auditor. |
 | deposit_status_slashed | 4 | DepositStatusSlashed indicates the deposit was slashed. |
 

 
 <a name="akash.verification.v1.DiscrepancyResolutionReason"></a>

 ### DiscrepancyResolutionReason
 DiscrepancyResolutionReason enumerates the reasons a discrepancy may be resolved.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | discrepancy_resolution_reason_unspecified | 0 | DiscrepancyResolutionReasonUnspecified is the default; no reason has been set. |
 | discrepancy_resolution_reason_auditor_a_correct | 1 | DiscrepancyResolutionReasonAuditorACorrect indicates auditor A's attestation was found correct. |
 | discrepancy_resolution_reason_auditor_b_correct | 2 | DiscrepancyResolutionReasonAuditorBCorrect indicates auditor B's attestation was found correct. |
 | discrepancy_resolution_reason_both_auditors_wrong | 3 | DiscrepancyResolutionReasonBothAuditorsWrong indicates both auditors' attestations were wrong. |
 | discrepancy_resolution_reason_provider_fault | 4 | DiscrepancyResolutionReasonProviderFault indicates the provider was at fault. |
 | discrepancy_resolution_reason_shared_fault | 5 | DiscrepancyResolutionReasonSharedFault indicates fault is shared between the parties. |
 | discrepancy_resolution_reason_evidence_inconclusive | 6 | DiscrepancyResolutionReasonEvidenceInconclusive indicates evidence was inconclusive. |
 | discrepancy_resolution_reason_governance_timeout_review | 7 | DiscrepancyResolutionReasonGovernanceTimeoutReview indicates governance reviewed a timed-out discrepancy. |
 

 
 <a name="akash.verification.v1.DiscrepancyStatus"></a>

 ### DiscrepancyStatus
 DiscrepancyStatus represents the lifecycle status of a discrepancy event.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | discrepancy_status_unspecified | 0 | DiscrepancyStatusUnspecified is the default; no status has been set. |
 | discrepancy_status_pending | 1 | DiscrepancyStatusPending indicates the discrepancy is awaiting resolution. |
 | discrepancy_status_resolved | 2 | DiscrepancyStatusResolved indicates the discrepancy has been resolved. |
 | discrepancy_status_timed_out | 3 | DiscrepancyStatusTimedOut indicates the discrepancy resolution window timed out. |
 

 
 <a name="akash.verification.v1.FaultAttribution"></a>

 ### FaultAttribution
 FaultAttribution identifies the responsible party in a settlement or resolution decision.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | fault_attribution_unspecified | 0 | FaultAttributionUnspecified is the default; no attribution has been set. |
 | fault_attribution_provider_fault | 1 | FaultAttributionProviderFault assigns fault to the provider. |
 | fault_attribution_auditor_fault | 2 | FaultAttributionAuditorFault assigns fault to the auditor. |
 | fault_attribution_shared_fault | 3 | FaultAttributionSharedFault attributes fault to both parties. |
 | fault_attribution_no_fault | 4 | FaultAttributionNoFault indicates neither party is at fault. |
 | fault_attribution_inconclusive | 5 | FaultAttributionInconclusive indicates the evidence does not allow attribution. |
 

 
 <a name="akash.verification.v1.FeeStatus"></a>

 ### FeeStatus
 FeeStatus tracks the lifecycle of the fee paid by a provider into an audit escrow.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | fee_status_unspecified | 0 | FeeStatusUnspecified is the default; no fee status has been set. |
 | fee_status_escrowed | 1 | FeeStatusEscrowed indicates the fee is held in escrow pending settlement. |
 | fee_status_released_to_auditor | 2 | FeeStatusReleasedToAuditor indicates the fee was released to the auditor upon settlement. |
 | fee_status_returned_to_provider | 3 | FeeStatusReturnedToProvider indicates the fee was returned to the provider upon settlement. |
 

 
 <a name="akash.verification.v1.GovernanceAttestationReason"></a>

 ### GovernanceAttestationReason
 GovernanceAttestationReason enumerates the typed reasons governance may revoke
or void attestations independent of the auditor.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | governance_attestation_reason_unspecified | 0 | GovernanceAttestationReasonUnspecified is the default; no reason has been set. |
 | governance_attestation_reason_fraudulent_provider | 1 | GovernanceAttestationReasonFraudulentProvider indicates governance found the provider fraudulent. |
 | governance_attestation_reason_compromised_provider | 2 | GovernanceAttestationReasonCompromisedProvider indicates the provider was compromised. |
 | governance_attestation_reason_provider_non_cooperation | 3 | GovernanceAttestationReasonProviderNonCooperation indicates the provider refused to cooperate with audit. |
 | governance_attestation_reason_faulty_auditor | 4 | GovernanceAttestationReasonFaultyAuditor indicates the auditor was found faulty. |
 | governance_attestation_reason_negligent_auditor | 5 | GovernanceAttestationReasonNegligentAuditor indicates the auditor was found negligent. |
 | governance_attestation_reason_evidence_insufficient | 6 | GovernanceAttestationReasonEvidenceInsufficient indicates evidence supporting the attestation was insufficient. |
 | governance_attestation_reason_emergency_safety_action | 7 | GovernanceAttestationReasonEmergencySafetyAction indicates an emergency safety action. |
 

 
 <a name="akash.verification.v1.ProviderBondSlashReason"></a>

 ### ProviderBondSlashReason
 ProviderBondSlashReason enumerates the reasons a provider bond may be slashed.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | provider_bond_slash_reason_unspecified | 0 | ProviderBondSlashReasonUnspecified is the default; no reason has been set. |
 | provider_bond_slash_reason_resource_misrepresentation | 1 | ProviderBondSlashReasonResourceMisrepresentation indicates the provider misrepresented its resources. |
 | provider_bond_slash_reason_capacity_overstatement | 2 | ProviderBondSlashReasonCapacityOverstatement indicates the provider overstated its capacity. |
 | provider_bond_slash_reason_fraudulent_snapshot | 3 | ProviderBondSlashReasonFraudulentSnapshot indicates the provider submitted a fraudulent snapshot. |
 | provider_bond_slash_reason_provider_compromise | 4 | ProviderBondSlashReasonProviderCompromise indicates the provider was compromised. |
 | provider_bond_slash_reason_sla_breach | 5 | ProviderBondSlashReasonSLABreach indicates the provider breached an SLA. |
 | provider_bond_slash_reason_non_cooperation_during_audit | 6 | ProviderBondSlashReasonNonCooperationDuringAudit indicates the provider did not cooperate during an audit. |
 

 
 <a name="akash.verification.v1.ProviderDepositStatus"></a>

 ### ProviderDepositStatus
 ProviderDepositStatus tracks the lifecycle of a provider's audit-escrow deposit.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | provider_deposit_status_unspecified | 0 | ProviderDepositStatusUnspecified is the default; no deposit status has been set. |
 | provider_deposit_status_escrowed | 1 | ProviderDepositStatusEscrowed indicates the provider's deposit is held in escrow. |
 | provider_deposit_status_returned_to_provider | 2 | ProviderDepositStatusReturnedToProvider indicates the deposit was returned to the provider. |
 | provider_deposit_status_slashed | 3 | ProviderDepositStatusSlashed indicates the deposit was slashed. |
 

 
 <a name="akash.verification.v1.VerificationGraceStatus"></a>

 ### VerificationGraceStatus
 VerificationGraceStatus tracks the status of a provider's verification grace period.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | verification_grace_status_unspecified | 0 | VerificationGraceStatusUnspecified is the default; no grace status has been set. |
 | verification_grace_status_active | 1 | VerificationGraceStatusActive indicates the grace window is open. |
 | verification_grace_status_expired | 2 | VerificationGraceStatusExpired indicates the grace window has elapsed. |
 | verification_grace_status_terminated | 3 | VerificationGraceStatusTerminated indicates the grace window was terminated early. |
 

 
 <a name="akash.verification.v1.VerificationTier"></a>

 ### VerificationTier
 VerificationTier represents provider verification levels.
Higher numeric value = higher trust. TierUnspecified (L0, permissionless)
means no attestation has been recorded for the provider.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | verification_tier_unspecified | 0 | TierUnspecified (L0) is the default; no attestation has been recorded. |
 | verification_tier_identified | 1 | TierIdentified (L1) indicates the provider has been identified by an auditor. |
 | verification_tier_verified | 2 | TierVerified (L2) indicates the provider's identity and capabilities have been verified. |
 | verification_tier_established | 3 | TierEstablished (L3) indicates the provider has a track record meeting the L3 criteria. |
 | verification_tier_trusted | 4 | TierTrusted (L4) is the highest tier, reserved for providers with long-running clean history. |
 

 
 <a name="akash.verification.v1.VoidedReason"></a>

 ### VoidedReason
 VoidedReason explains why an attestation transitioned to the Voided status.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | voided_reason_unspecified | 0 | VoidedReasonUnspecified is the default; no reason has been set. |
 | voided_reason_discrepancy | 1 | VoidedReasonDiscrepancy indicates the attestation was voided due to an open discrepancy. |
 | voided_reason_governance | 2 | VoidedReasonGovernance indicates the attestation was voided by governance action. |
 | voided_reason_bond_withdrawn | 3 | VoidedReasonBondWithdrawn indicates the auditor's bond was withdrawn. |
 | voided_reason_bond_slashed | 4 | VoidedReasonBondSlashed indicates the auditor's bond was slashed. |
 

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/verification/v1/verificationrequirement.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/verification/v1/verificationrequirement.proto
 

 
 <a name="akash.verification.v1.VerificationRequirement"></a>

 ### VerificationRequirement
 VerificationRequirement holds the verification constraints that a deployment
places on the providers eligible to bid on a group. It is referenced from the
deployment-module `PlacementRequirements` message; this file is split out from
`state.proto` so the deployment-side import surface stays minimal (only the
shared enums in `types.proto` need to come along).
A `min_tier` of 0 (TierUnspecified) means "no requirement".

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `min_tier` | [VerificationTier](#akash.verification.v1.VerificationTier) |  | MinTier is the minimum verification tier required of bidding providers. A value of 0 (TierUnspecified) means no tier requirement. |
 | `required_capabilities` | [CapabilityFlag](#akash.verification.v1.CapabilityFlag) | repeated | RequiredCapabilities is the set of capability flags providers must assert. |
 | `required_auditors` | [string](#string) | repeated | RequiredAuditors is an optional list of specific auditor bech32 addresses whose attestations must be present on the provider, interpreted per `auditor_mode`. |
 | `auditor_mode` | [AuditorSelectionMode](#akash.verification.v1.AuditorSelectionMode) |  | AuditorMode controls how `required_auditors` is evaluated. AuditorSelectionModeUnspecified is treated as Any. |
 | `min_auditor_count` | [uint32](#uint32) |  | MinAuditorCount is the minimum number of independent auditors that must have attested the provider, regardless of identity. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/base/attributes/v1/attribute.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/base/attributes/v1/attribute.proto
 

 
 <a name="akash.base.attributes.v1.Attribute"></a>

 ### Attribute
 Attribute represents an arbitrary attribute key-value pair.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `key` | [string](#string) |  | Key of the attribute (e.g., "region", "cpu_architecture", etc.). |
 | `value` | [string](#string) |  | Value of the attribute (e.g., "us-west", "x86_64", etc.). |
 
 

 

 
 <a name="akash.base.attributes.v1.PlacementRequirements"></a>

 ### PlacementRequirements
 PlacementRequirements represents the requirements for a provider placement on the network.
It is used to specify the characteristics and constraints of a provider that can be used to satisfy a deployment request.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `signed_by` | [SignedBy](#akash.base.attributes.v1.SignedBy) |  | SignedBy holds the list of keys that tenants expect to have signatures from. |
 | `attributes` | [Attribute](#akash.base.attributes.v1.Attribute) | repeated | Attribute holds the list of attributes tenant expects from the provider. |
 | `verification` | [akash.verification.v1.VerificationRequirement](#akash.verification.v1.VerificationRequirement) |  | Verification holds the verification requirements for this placement. |
 
 

 

 
 <a name="akash.base.attributes.v1.SignedBy"></a>

 ### SignedBy
 SignedBy represents validation accounts that tenant expects signatures for provider attributes.
AllOf has precedence i.e. if there is at least one entry AnyOf is ignored regardless to how many
entries there.

TODO: this behaviour to be discussed

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `all_of` | [string](#string) | repeated | AllOf indicates all keys in this list must have signed attributes. |
 | `any_of` | [string](#string) | repeated | AnyOf means that at least one of the keys from the list must have signed attributes. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/audit/v1/audit.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/audit/v1/audit.proto
 

 
 <a name="akash.audit.v1.AttributesFilters"></a>

 ### AttributesFilters
 AttributesFilters defines attribute filters that can be used to filter deployments.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `auditors` | [string](#string) | repeated | Auditors contains a list of auditor account bech32 addresses. |
 | `owners` | [string](#string) | repeated | Owners contains a list of owner account bech32 addresses. |
 
 

 

 
 <a name="akash.audit.v1.AuditedAttributesStore"></a>

 ### AuditedAttributesStore
 AuditedAttributesStore stores the audited attributes of the provider.
Attributes that have been audited are those that have been verified by an auditor.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `attributes` | [akash.base.attributes.v1.Attribute](#akash.base.attributes.v1.Attribute) | repeated | Attributes holds a list of key-value pairs of provider attributes. Attributes are arbitrary values that a provider exposes. |
 
 

 

 
 <a name="akash.audit.v1.AuditedProvider"></a>

 ### AuditedProvider
 AuditedProvider stores owner, auditor and attributes details.
An AuditedProvider is a provider that has undergone a verification or auditing process to ensure that it meets certain standards or requirements by an auditor.
An auditor can be any valid account on-chain.
NOTE: There are certain teams providing auditing services, which should be accounted for when deploying.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the account bech32 address of the provider. It is a string representing a valid account address.

Example: "akash1..." |
 | `auditor` | [string](#string) |  | Auditor is the account bech32 address of the auditor. It is a string representing a valid account address.

Example: "akash1..." |
 | `attributes` | [akash.base.attributes.v1.Attribute](#akash.base.attributes.v1.Attribute) | repeated | Attributes holds a list of key-value pairs of provider attributes. Attributes are arbitrary values that a provider exposes. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/audit/v1/event.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/audit/v1/event.proto
 

 
 <a name="akash.audit.v1.EventTrustedAuditorCreated"></a>

 ### EventTrustedAuditorCreated
 EventTrustedAuditorCreated defines an SDK message for when a trusted auditor is created.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the account bech32 address of the provider. It is a string representing a valid account address.

Example: "akash1..." |
 | `auditor` | [string](#string) |  | Auditor is the account address of the auditor. It is a string representing a valid account address.

Example: "akash1..." |
 
 

 

 
 <a name="akash.audit.v1.EventTrustedAuditorDeleted"></a>

 ### EventTrustedAuditorDeleted
 EventTrustedAuditorDeleted defines an event for when a trusted auditor is deleted.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the account bech32 address of the provider. It is a string representing a valid account address.

Example: "akash1..." |
 | `auditor` | [string](#string) |  | Auditor is the account address of the auditor. It is a string representing a valid account address.

Example: "akash1..." |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/audit/v1/genesis.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/audit/v1/genesis.proto
 

 
 <a name="akash.audit.v1.GenesisState"></a>

 ### GenesisState
 GenesisState defines the basic genesis state used by audit module.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `providers` | [AuditedProvider](#akash.audit.v1.AuditedProvider) | repeated | Providers contains a list of audited providers account addresses. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/audit/v1/msg.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/audit/v1/msg.proto
 

 
 <a name="akash.audit.v1.MsgDeleteProviderAttributes"></a>

 ### MsgDeleteProviderAttributes
 MsgDeleteProviderAttributes defined the Msg/DeleteProviderAttributes

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the account bech32 address of the provider. It is a string representing a valid account address.

Example: "akash1..." |
 | `auditor` | [string](#string) |  | Auditor is the account address of the auditor. It is a string representing a valid account address.

Example: "akash1..." |
 | `keys` | [string](#string) | repeated | Keys holds a list of keys of audited provider attributes to delete from the audit. |
 
 

 

 
 <a name="akash.audit.v1.MsgDeleteProviderAttributesResponse"></a>

 ### MsgDeleteProviderAttributesResponse
 MsgDeleteProviderAttributesResponse defines the Msg/ProviderAttributes response type.

 

 

 
 <a name="akash.audit.v1.MsgSignProviderAttributes"></a>

 ### MsgSignProviderAttributes
 MsgSignProviderAttributes defines an SDK message for signing a provider attributes.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the account bech32 address of the provider. It is a string representing a valid account address.

Example: "akash1..." |
 | `auditor` | [string](#string) |  | Auditor is the account address of the auditor. It is a string representing a valid account address.

Example: "akash1..." |
 | `attributes` | [akash.base.attributes.v1.Attribute](#akash.base.attributes.v1.Attribute) | repeated | Attributes holds a list of key-value pairs of provider attributes to be audited. Attributes are arbitrary values that a provider exposes. |
 
 

 

 
 <a name="akash.audit.v1.MsgSignProviderAttributesResponse"></a>

 ### MsgSignProviderAttributesResponse
 MsgSignProviderAttributesResponse defines the Msg/CreateProvider response type.

 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/audit/v1/query.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/audit/v1/query.proto
 

 
 <a name="akash.audit.v1.QueryAllProvidersAttributesRequest"></a>

 ### QueryAllProvidersAttributesRequest
 QueryAllProvidersAttributesRequest is request type for the Query/All Providers RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | Pagination is used to paginate the request. |
 
 

 

 
 <a name="akash.audit.v1.QueryAuditorAttributesRequest"></a>

 ### QueryAuditorAttributesRequest
 QueryAuditorAttributesRequest is request type for the Query/Providers RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `auditor` | [string](#string) |  | Auditor is the account address of the auditor. It is a string representing a valid account address.

Example: "akash1..." |
 | `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | Pagination is used to paginate the request. |
 
 

 

 
 <a name="akash.audit.v1.QueryProviderAttributesRequest"></a>

 ### QueryProviderAttributesRequest
 QueryProviderAttributesRequest is request type for the Query/Provider RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the account bech32 address of the provider. It is a string representing a valid account address.

Example: "akash1..." |
 | `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | Pagination is used to paginate the request. |
 
 

 

 
 <a name="akash.audit.v1.QueryProviderAuditorRequest"></a>

 ### QueryProviderAuditorRequest
 QueryProviderAuditorRequest is request type for the Query/Providers RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `auditor` | [string](#string) |  | Auditor is the account address of the auditor. It is a string representing a valid account address.

Example: "akash1..." |
 | `owner` | [string](#string) |  | Owner is the account bech32 address of the provider. It is a string representing a valid account address.

Example: "akash1..." |
 
 

 

 
 <a name="akash.audit.v1.QueryProviderRequest"></a>

 ### QueryProviderRequest
 QueryProviderRequest is request type for the Query/Provider RPC method

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `auditor` | [string](#string) |  | Auditor is the account address of the auditor. It is a string representing a valid account address.

Example: "akash1..." |
 | `owner` | [string](#string) |  | Owner is the account bech32 address of the provider. It is a string representing a valid account address.

Example: "akash1..." |
 
 

 

 
 <a name="akash.audit.v1.QueryProvidersResponse"></a>

 ### QueryProvidersResponse
 QueryProvidersResponse is response type for the Query/Providers RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `providers` | [AuditedProvider](#akash.audit.v1.AuditedProvider) | repeated | Providers contains a list of audited providers account addresses. |
 | `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | Pagination is used to paginate results. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.audit.v1.Query"></a>

 ### Query
 Query defines the gRPC querier service for the audit package.

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `AllProvidersAttributes` | [QueryAllProvidersAttributesRequest](#akash.audit.v1.QueryAllProvidersAttributesRequest) | [QueryProvidersResponse](#akash.audit.v1.QueryProvidersResponse) | AllProvidersAttributes queries all providers. buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE buf:lint:ignore RPC_RESPONSE_STANDARD_NAME | GET|/akash/audit/v1/audit/attributes/list|
 | `ProviderAttributes` | [QueryProviderAttributesRequest](#akash.audit.v1.QueryProviderAttributesRequest) | [QueryProvidersResponse](#akash.audit.v1.QueryProvidersResponse) | ProviderAttributes queries all provider signed attributes. buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE buf:lint:ignore RPC_RESPONSE_STANDARD_NAME | GET|/akash/audit/v1/audit/attributes/{owner}/list|
 | `ProviderAuditorAttributes` | [QueryProviderAuditorRequest](#akash.audit.v1.QueryProviderAuditorRequest) | [QueryProvidersResponse](#akash.audit.v1.QueryProvidersResponse) | ProviderAuditorAttributes queries provider signed attributes by specific auditor. buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE buf:lint:ignore RPC_RESPONSE_STANDARD_NAME | GET|/akash/audit/v1/audit/attributes/{auditor}/{owner}|
 | `AuditorAttributes` | [QueryAuditorAttributesRequest](#akash.audit.v1.QueryAuditorAttributesRequest) | [QueryProvidersResponse](#akash.audit.v1.QueryProvidersResponse) | AuditorAttributes queries all providers signed by this auditor. buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE buf:lint:ignore RPC_RESPONSE_STANDARD_NAME | GET|/akash/provider/v1/auditor/{auditor}/list|
 
  <!-- end services -->

 
 
 <a name="akash/audit/v1/service.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/audit/v1/service.proto
 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.audit.v1.Msg"></a>

 ### Msg
 Msg defines the audit Msg service.

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `SignProviderAttributes` | [MsgSignProviderAttributes](#akash.audit.v1.MsgSignProviderAttributes) | [MsgSignProviderAttributesResponse](#akash.audit.v1.MsgSignProviderAttributesResponse) | SignProviderAttributes defines a method that signs provider attributes. | |
 | `DeleteProviderAttributes` | [MsgDeleteProviderAttributes](#akash.audit.v1.MsgDeleteProviderAttributes) | [MsgDeleteProviderAttributesResponse](#akash.audit.v1.MsgDeleteProviderAttributesResponse) | DeleteProviderAttributes defines a method that deletes provider attributes. | |
 
  <!-- end services -->

 
 
 <a name="akash/base/deposit/v1/deposit.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/base/deposit/v1/deposit.proto
 

 
 <a name="akash.base.deposit.v1.Deposit"></a>

 ### Deposit
 Deposit is a data type used by MsgCreateDeployment, MsgDepositDeployment and MsgCreateBid to indicate source of the deposit.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | amount specifies the amount of coins to include in the deployment's first deposit. |
 | `sources` | [Source](#akash.base.deposit.v1.Source) | repeated | Sources is the set of deposit sources, each entry must be unique. |
 
 

 

  <!-- end messages -->

 
 <a name="akash.base.deposit.v1.Source"></a>

 ### Source
 Source is an enum which lists source of funds for deployment deposit.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | invalid | 0 | Prefix should start with 0 in enum. So declaring dummy state. |
 | balance | 1 | SourceBalance denotes account balance as source of funds |
 | grant | 2 | SourceGrant denotes authz grants as source of funds |
 

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/base/offchain/sign/v1/sign.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/base/offchain/sign/v1/sign.proto
 

 
 <a name="akash.base.offchain.sign.v1.MsgSignData"></a>

 ### MsgSignData
 MsgSignData defines an arbitrary, general-purpose, off-chain message

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `signer` | [string](#string) |  | Signer is the sdk.AccAddress of the message signer |
 | `data` | [bytes](#bytes) |  | Data represents the raw bytes of the content that is signed (text, json, etc) |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/base/resources/v1beta4/resourcevalue.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/base/resources/v1beta4/resourcevalue.proto
 

 
 <a name="akash.base.resources.v1beta4.ResourceValue"></a>

 ### ResourceValue
 Unit stores cpu, memory and storage metrics.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `val` | [bytes](#bytes) |  |  |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/base/resources/v1beta4/cpu.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/base/resources/v1beta4/cpu.proto
 

 
 <a name="akash.base.resources.v1beta4.CPU"></a>

 ### CPU
 CPU stores resource units and cpu config attributes.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `units` | [ResourceValue](#akash.base.resources.v1beta4.ResourceValue) |  | Units of the CPU, which represents the number of CPUs available. This field is required and must be a non-negative integer. |
 | `attributes` | [akash.base.attributes.v1.Attribute](#akash.base.attributes.v1.Attribute) | repeated | Attributes holds a list of key-value attributes that describe the GPU, such as its model, memory and interface. This field is required and must be a list of `Attribute` messages. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/base/resources/v1beta4/endpoint.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/base/resources/v1beta4/endpoint.proto
 

 
 <a name="akash.base.resources.v1beta4.Endpoint"></a>

 ### Endpoint
 Endpoint describes a publicly accessible IP service.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `kind` | [Endpoint.Kind](#akash.base.resources.v1beta4.Endpoint.Kind) |  | Kind describes how the endpoint is implemented when the lease is deployed. |
 | `sequence_number` | [uint32](#uint32) |  | SequenceNumber represents a sequence number for the Endpoint. |
 
 

 

  <!-- end messages -->

 
 <a name="akash.base.resources.v1beta4.Endpoint.Kind"></a>

 ### Endpoint.Kind
 Kind describes how the endpoint is implemented when the lease is deployed.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | SHARED_HTTP | 0 | Describes an endpoint that becomes a Kubernetes Ingress. |
 | RANDOM_PORT | 1 | Describes an endpoint that becomes a Kubernetes NodePort. |
 | LEASED_IP | 2 | Describes an endpoint that becomes a leased IP. |
 

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/base/resources/v1beta4/gpu.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/base/resources/v1beta4/gpu.proto
 

 
 <a name="akash.base.resources.v1beta4.GPU"></a>

 ### GPU
 GPU stores resource units and gpu configuration attributes.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `units` | [ResourceValue](#akash.base.resources.v1beta4.ResourceValue) |  | The resource value of the GPU, which represents the number of GPUs available. This field is required and must be a non-negative integer. |
 | `attributes` | [akash.base.attributes.v1.Attribute](#akash.base.attributes.v1.Attribute) | repeated |  |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/base/resources/v1beta4/memory.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/base/resources/v1beta4/memory.proto
 

 
 <a name="akash.base.resources.v1beta4.Memory"></a>

 ### Memory
 Memory stores resource quantity and memory attributes.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `quantity` | [ResourceValue](#akash.base.resources.v1beta4.ResourceValue) |  | Quantity of memory available, which represents the amount of memory in bytes. This field is required and must be a non-negative integer. |
 | `attributes` | [akash.base.attributes.v1.Attribute](#akash.base.attributes.v1.Attribute) | repeated | Attributes that describe the memory, such as its type and speed. This field is required and must be a list of Attribute key-values. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/base/resources/v1beta4/storage.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/base/resources/v1beta4/storage.proto
 

 
 <a name="akash.base.resources.v1beta4.Storage"></a>

 ### Storage
 Storage stores resource quantity and storage attributes.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `name` | [string](#string) |  | Name holds an arbitrary name for the storage resource. |
 | `quantity` | [ResourceValue](#akash.base.resources.v1beta4.ResourceValue) |  | Quantity of storage available, which represents the amount of memory in bytes. This field is required and must be a non-negative integer. |
 | `attributes` | [akash.base.attributes.v1.Attribute](#akash.base.attributes.v1.Attribute) | repeated | Attributes that describe the storage. This field is required and must be a list of Attribute key-values. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/base/resources/v1beta4/resources.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/base/resources/v1beta4/resources.proto
 

 
 <a name="akash.base.resources.v1beta4.Resources"></a>

 ### Resources
 Resources describes all available resources types for deployment/node etc
if field is nil resource is not present in the given data-structure

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [uint32](#uint32) |  | Id is a unique identifier for the resources. |
 | `cpu` | [CPU](#akash.base.resources.v1beta4.CPU) |  | CPU resources available, including the architecture, number of cores and other details. This field is optional and can be empty if no CPU resources are available. |
 | `memory` | [Memory](#akash.base.resources.v1beta4.Memory) |  | Memory resources available, including the quantity and attributes. This field is optional and can be empty if no memory resources are available. |
 | `storage` | [Storage](#akash.base.resources.v1beta4.Storage) | repeated | Storage resources available, including the quantity and attributes. This field is optional and can be empty if no storage resources are available. |
 | `gpu` | [GPU](#akash.base.resources.v1beta4.GPU) |  | GPU resources available, including the type, architecture and other details. This field is optional and can be empty if no GPU resources are available. |
 | `endpoints` | [Endpoint](#akash.base.resources.v1beta4.Endpoint) | repeated | Endpoint resources available |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/bme/v1/types.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/bme/v1/types.proto
 

 
 <a name="akash.bme.v1.BurnMintPair"></a>

 ### BurnMintPair
 BurnMintPair represents a pair of burn and mint operations with their respective prices

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `burned` | [CoinPrice](#akash.bme.v1.CoinPrice) |  | burned is the coin burned |
 | `minted` | [CoinPrice](#akash.bme.v1.CoinPrice) |  | minted is coin minted |
 
 

 

 
 <a name="akash.bme.v1.CoinPrice"></a>

 ### CoinPrice
 CoinPrice represents a coin amount with its associated oracle price at a specific point in time

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `coin` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | coin is the token amount |
 | `price` | [string](#string) |  | price (at oracle) of the coin at burn/mint event |
 
 

 

 
 <a name="akash.bme.v1.CollateralRatio"></a>

 ### CollateralRatio
 CollateralRatio represents the current collateral ratio

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `ratio` | [string](#string) |  | ratio is CR = (VaultAKT * Price) / OutstandingACT |
 | `status` | [MintStatus](#akash.bme.v1.MintStatus) |  | status indicates the current circuit breaker status |
 | `reference_price` | [string](#string) |  | reference_price is the price used to calculate CR |
 
 

 

 
 <a name="akash.bme.v1.LedgerCanceledRecord"></a>

 ### LedgerCanceledRecord
 LedgerCanceledRecord

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | owner source of the coins to be burned |
 | `cancel_reason` | [LedgerCanceledRecord.BMCancelReason](#akash.bme.v1.LedgerCanceledRecord.BMCancelReason) |  | cancel_reason |
 | `to` | [string](#string) |  | to destination of the minted coins. if minted coin is ACT, "to" must be same as signer |
 | `coins_to_burn` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | coins_to_burn |
 | `denom_to_mint` | [string](#string) |  | denom_to_mint |
 
 

 

 
 <a name="akash.bme.v1.LedgerID"></a>

 ### LedgerID
 LedgerID uniquely identifies a ledger entry by block height and sequence number

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `height` | [int64](#int64) |  | height is the block height when the ledger entry was created |
 | `sequence` | [int64](#int64) |  | sequence is the sequence number within the block (for ordering) |
 
 

 

 
 <a name="akash.bme.v1.LedgerPendingRecord"></a>

 ### LedgerPendingRecord
 LedgerPendingRecord

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | owner source of the coins to be burned |
 | `to` | [string](#string) |  | to destination of the minted coins. if minted coin is ACT, "to" must be same as signer |
 | `coins_to_burn` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | coins_to_burn |
 | `denom_to_mint` | [string](#string) |  | denom_to_mint |
 | `attempts` | [uint32](#uint32) |  | attempts is the number of times this record has been processed and failed with a retriable error |
 
 

 

 
 <a name="akash.bme.v1.LedgerRecord"></a>

 ### LedgerRecord
 LedgerRecord stores information of burn/mint event of token A burn to mint token B

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `burned_from` | [string](#string) |  | burned_from source address of the tokens burned |
 | `minted_to` | [string](#string) |  | minted_to destination address of the tokens minted |
 | `burner` | [string](#string) |  | module is module account performing burn |
 | `minter` | [string](#string) |  | module is module account performing mint |
 | `burned` | [CoinPrice](#akash.bme.v1.CoinPrice) |  | burned is the coin burned at price |
 | `minted` | [CoinPrice](#akash.bme.v1.CoinPrice) |  | minted is coin minted at price |
 | `spread` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
 | `remint_credit_issued` | [CoinPrice](#akash.bme.v1.CoinPrice) |  |  |
 | `remint_credit_accrued` | [CoinPrice](#akash.bme.v1.CoinPrice) |  |  |
 
 

 

 
 <a name="akash.bme.v1.LedgerRecordID"></a>

 ### LedgerRecordID
 LedgerRecordID

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `denom` | [string](#string) |  | denom is the asset denomination |
 | `to_denom` | [string](#string) |  | to_denom is what denom swap to |
 | `source` | [string](#string) |  |  |
 | `height` | [int64](#int64) |  |  |
 | `sequence` | [int64](#int64) |  |  |
 
 

 

 
 <a name="akash.bme.v1.State"></a>

 ### State
 State tracks net burn metrics since BME start

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `balances` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | burned is the cumulative burn for tracked tokens |
 | `total_burned` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | burned is the cumulative burn for tracked tokens |
 | `total_minted` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | minted is the cumulative mint back for tracked tokens |
 | `remint_credits` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | remint_credits tracks available credits for reminting tokens (e.g., from previous burns that can be reminted without additional collateral) |
 
 

 

 
 <a name="akash.bme.v1.Status"></a>

 ### Status
 Status stores status of mint operations

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `status` | [MintStatus](#akash.bme.v1.MintStatus) |  |  |
 | `previous_status` | [MintStatus](#akash.bme.v1.MintStatus) |  |  |
 | `epoch_height_diff` | [int64](#int64) |  |  |
 
 

 

  <!-- end messages -->

 
 <a name="akash.bme.v1.LedgerCanceledRecord.BMCancelReason"></a>

 ### LedgerCanceledRecord.BMCancelReason
 BMCancelReason is an enum indicating reasons of failure for burn/mint request

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | unknown | 0 | Prefix should start with 0 in enum. So declaring dummy state. |
 | epsilon | 1 | BMCancelReasonEpsilon the result of conversion is below the smallest meaningful difference (10^-6) |
 | zero_price | 2 | BMCancelReasonZeroPrice oracle price is zero |
 | insufficient_funds | 3 | BMCancelReasonInsufficientFunds insufficient vault/supply funds |
 | invalid_denom | 4 | BMCancelReasonInvalidDenom denomination is not registered |
 | invalid_amount | 5 | BMCancelReasonInvalidAmount zero or invalid burn amount |
 | minimum_mint | 6 | BMCancelReasonMinimumMint mint output below minimum threshold |
 | mint_failed | 7 | BMCancelReasonMintFailed bank MintCoins operation failed |
 | burn_failed | 8 | BMCancelReasonBurnFailed bank BurnCoins operation failed |
 | max_attempts | 9 | BMCancelReasonMaxAttempts exceeded maximum pending processing attempts |
 

 
 <a name="akash.bme.v1.LedgerRecordStatus"></a>

 ### LedgerRecordStatus
 LedgerRecordStatus indicates the current state of a burn/mint ledger record

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | ledger_record_status_invalid | 0 | LEDGER_RECORD_STATUS_INVALID is the default/uninitialized value This status should never appear in a valid ledger record |
 | ledger_record_status_pending | 1 | LEDGER_RECORD_STATUS_PENDING indicates a burn/mint operation has been initiated but not yet executed (e.g., waiting for oracle price or circuit breaker clearance) |
 | ledger_record_status_executed | 2 | LEDGER_RECORD_STATUS_EXECUTED indicates the burn/mint operation has been successfully completed and tokens have been burned and minted |
 | ledger_record_status_canceled | 3 | LEDGER_RECORD_STATUS_CANCELED indicates the burn/mint operation has encountered error and funds have been returned to the owner successfully completed and tokens have been burned and minted |
 

 
 <a name="akash.bme.v1.MintStatus"></a>

 ### MintStatus
 MintStatus indicates the current state of mint

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | mint_status_unspecified | 0 | MINT_STATUS_UNSPECIFIED is the default value |
 | mint_status_healthy | 1 | MINT_STATUS_HEALTHY indicates normal operation (CR > warn threshold) |
 | mint_status_warning | 2 | MINT_STATUS_WARNING indicates CR is below warning threshold |
 | mint_status_halt_cr | 3 | MINT_STATUS_HALT_CR indicates CR is below halt threshold, mints paused |
 | mint_status_halt_oracle | 4 | MINT_STATUS_HALT_ORACLE indicates circuit breaker tripped due to unhealthy oracle price |
 

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/bme/v1/events.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/bme/v1/events.proto
 

 
 <a name="akash.bme.v1.EventLedgerRecordCanceled"></a>

 ### EventLedgerRecordCanceled
 EventLedgerRecordCanceled emitted information of unsuccessful burn/mint event

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [LedgerRecordID](#akash.bme.v1.LedgerRecordID) |  | burned_from source address of the tokens burned |
 | `cancel_reason` | [LedgerCanceledRecord.BMCancelReason](#akash.bme.v1.LedgerCanceledRecord.BMCancelReason) |  | fail_reason |
 | `owner` | [string](#string) |  | owner source of the coins to be burned |
 | `to` | [string](#string) |  | to destination of the minted coins. if minted coin is ACT, "to" must be same as signer |
 | `coins_to_burn` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | coins_to_burn |
 | `denom_to_mint` | [string](#string) |  | denom_to_mint |
 
 

 

 
 <a name="akash.bme.v1.EventLedgerRecordExecuted"></a>

 ### EventLedgerRecordExecuted
 EventLedgerRecordExecuted emitted information of burn/mint event of token A burn to mint token B

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [LedgerRecordID](#akash.bme.v1.LedgerRecordID) |  | burned_from source address of the tokens burned |
 | `burned_from` | [string](#string) |  | burned_from source address of the tokens burned |
 | `minted_to` | [string](#string) |  | minted_to destination address of the tokens minted |
 | `burner` | [string](#string) |  | module is module account performing burn |
 | `minter` | [string](#string) |  | module is module account performing mint |
 | `burned` | [CoinPrice](#akash.bme.v1.CoinPrice) |  | burned is the coin burned at price |
 | `minted` | [CoinPrice](#akash.bme.v1.CoinPrice) |  | minted is coin minted at price |
 | `spread` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
 | `remint_credit_issued` | [CoinPrice](#akash.bme.v1.CoinPrice) |  |  |
 | `remint_credit_accrued` | [CoinPrice](#akash.bme.v1.CoinPrice) |  |  |
 
 

 

 
 <a name="akash.bme.v1.EventMintStatusChange"></a>

 ### EventMintStatusChange
 EventCircuitBreakerStatusChange is emitted when circuit breaker status changes

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `previous_status` | [MintStatus](#akash.bme.v1.MintStatus) |  | previous_status is the previous status |
 | `new_status` | [MintStatus](#akash.bme.v1.MintStatus) |  | new_status is the new status |
 | `collateral_ratio` | [string](#string) |  | collateral_ratio is the CR that triggered the change |
 
 

 

 
 <a name="akash.bme.v1.EventVaultFunded"></a>

 ### EventVaultFunded
 EventVaultFunded is emitted when the vault is seeded with AKT

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | amount is the AKT amount added to vault |
 | `source` | [string](#string) |  | source is where the funds came from |
 | `new_vault_balance` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | new_vault_balance is the new vault balance |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/bme/v1/filters.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/bme/v1/filters.proto
 

 
 <a name="akash.bme.v1.LedgerRecordFilters"></a>

 ### LedgerRecordFilters
 LedgerRecordFilters defines filters used to filter ledger records

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `source` | [string](#string) |  | source is the account address of the user who initiated the burn/mint |
 | `denom` | [string](#string) |  | denom filters by the burn denomination |
 | `to_denom` | [string](#string) |  | to_denom filters by the mint denomination |
 | `status` | [string](#string) |  | status filters by record status (pending, executed or failed). Uses the string representation of LedgerRecordStatus enum values. If empty, returns both pending and executed records. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/bme/v1/params.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/bme/v1/params.proto
 

 
 <a name="akash.bme.v1.Params"></a>

 ### Params
 Params defines the parameters for the BME module

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `circuit_breaker_warn_threshold` | [uint32](#uint32) |  | circuit_breaker_warn_threshold is the CR below which warning is triggered Stored as basis points * 100 (e.g., 9500 = 0.95) |
 | `circuit_breaker_halt_threshold` | [uint32](#uint32) |  | circuit_breaker_halt_threshold is the CR below which mints are halted Stored as basis points * 100 (e.g., 9000 = 0.90) |
 | `min_epoch_blocks` | [int64](#int64) |  | min_epoch_blocks is the minimum amount of blocks required for ACT mints |
 | `epoch_blocks_backoff_percent` | [uint32](#uint32) |  | epoch_blocks_backoff increase of runway_blocks in % during warn threshold for drop in 1 basis point of circuit_breaker_warn_threshold Stored as basis points * 100 (e.g., 9500 = 0.95) e.g: runway_blocks = 100 min_runway_blocks_backoff = 1000 circuit_breaker_warn_threshold drops from 0.95 to 0.94 then runway_blocks = (100*0.1 + 100) = 110

 circuit_breaker_warn_threshold drops from 0.94 to 0.92 then runway_blocks = (110*(0.1*2) + 110) = 132 |
 | `mint_spread_bps` | [uint32](#uint32) |  | mint_spread_bps is the spread in basis points applied during ACT mint (default: 25 bps = 0.25%) |
 | `settle_spread_bps` | [uint32](#uint32) |  | settle_spread_bps is the spread in basis points applied during settlement (default: 0 for no provider tax) |
 | `max_endblocker_records` | [uint32](#uint32) |  | max_endblocker_records is the deterministic upper bound on pending ledger records processed in a single EndBlocker invocation. |
 | `min_mint` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | min_mint minimum amount of ACT required to be minted in the new transaction |
 | `max_pending_attempts` | [uint32](#uint32) |  | max_pending_attempts is the maximum number of EndBlocker processing attempts for a pending record before it is canceled. Applies to retriable errors only. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/bme/v1/genesis.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/bme/v1/genesis.proto
 

 
 <a name="akash.bme.v1.GenesisLedgerPendingRecord"></a>

 ### GenesisLedgerPendingRecord
 GenesisLedgerPendingRecord

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [LedgerRecordID](#akash.bme.v1.LedgerRecordID) |  |  |
 | `record` | [LedgerPendingRecord](#akash.bme.v1.LedgerPendingRecord) |  |  |
 
 

 

 
 <a name="akash.bme.v1.GenesisLedgerRecord"></a>

 ### GenesisLedgerRecord
 GenesisLedgerRecord

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [LedgerRecordID](#akash.bme.v1.LedgerRecordID) |  |  |
 | `record` | [LedgerRecord](#akash.bme.v1.LedgerRecord) |  |  |
 
 

 

 
 <a name="akash.bme.v1.GenesisLedgerState"></a>

 ### GenesisLedgerState
 GenesisLedgerState

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `records` | [GenesisLedgerRecord](#akash.bme.v1.GenesisLedgerRecord) | repeated |  |
 | `pending_records` | [GenesisLedgerPendingRecord](#akash.bme.v1.GenesisLedgerPendingRecord) | repeated |  |
 
 

 

 
 <a name="akash.bme.v1.GenesisState"></a>

 ### GenesisState
 GenesisState defines the BME module's genesis state

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `params` | [Params](#akash.bme.v1.Params) |  | params defines the module parameters |
 | `state` | [GenesisVaultState](#akash.bme.v1.GenesisVaultState) |  | state is the initial vault state |
 | `ledger` | [GenesisLedgerState](#akash.bme.v1.GenesisLedgerState) |  |  |
 
 

 

 
 <a name="akash.bme.v1.GenesisVaultState"></a>

 ### GenesisVaultState
 GenesisVaultState

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `total_burned` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | burned is the cumulative burn for tracked tokens |
 | `total_minted` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | minted is the cumulative mint back for tracked tokens |
 | `remint_credits` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | remint_credits tracks available credits for reminting tokens (e.g., from previous burns that can be reminted without additional collateral) |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/bme/v1/msgs.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/bme/v1/msgs.proto
 

 
 <a name="akash.bme.v1.MsgBurnACT"></a>

 ### MsgBurnACT
 MsgMintACT defines the message for burning one token to mint another
Allows burning AKT to mint ACT, or burning unused ACT back to AKT

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | owner source of the coins to be burned |
 | `to` | [string](#string) |  | to destination of the minted coins. if minted coin is ACT, "to" must be same as signer |
 | `coins_to_burn` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | coins_to_burn |
 
 

 

 
 <a name="akash.bme.v1.MsgBurnACTResponse"></a>

 ### MsgBurnACTResponse
 MsgBurnMintResponse is the response type for MsgBurnMint

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [LedgerRecordID](#akash.bme.v1.LedgerRecordID) |  |  |
 | `status` | [LedgerRecordStatus](#akash.bme.v1.LedgerRecordStatus) |  |  |
 
 

 

 
 <a name="akash.bme.v1.MsgBurnMint"></a>

 ### MsgBurnMint
 MsgBurnMint defines the message for burning one token to mint another
Allows burning AKT to mint ACT, or burning unused ACT back to AKT

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | owner source of the coins to be burned |
 | `to` | [string](#string) |  | to destination of the minted coins. if minted coin is ACT, "to" must be same as signer |
 | `coins_to_burn` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | coins_to_burn |
 | `denom_to_mint` | [string](#string) |  | denom_to_mint |
 
 

 

 
 <a name="akash.bme.v1.MsgBurnMintResponse"></a>

 ### MsgBurnMintResponse
 MsgBurnMintResponse is the response type for MsgBurnMint

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [LedgerRecordID](#akash.bme.v1.LedgerRecordID) |  |  |
 | `status` | [LedgerRecordStatus](#akash.bme.v1.LedgerRecordStatus) |  |  |
 
 

 

 
 <a name="akash.bme.v1.MsgFundVault"></a>

 ### MsgFundVault
 MsgFundVault defines the message for funding the BME vault with AKT
This is used to provide an initial volatility buffer

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `authority` | [string](#string) |  | authority is the address that controls the module (governance) |
 | `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | amount is the AKT amount to seed the vault with |
 | `source` | [string](#string) |  | source is the source of funds (e.g., community pool) |
 
 

 

 
 <a name="akash.bme.v1.MsgFundVaultResponse"></a>

 ### MsgFundVaultResponse
 MsgFundVaultResponse is the response type for MsgFundVault

 

 

 
 <a name="akash.bme.v1.MsgMintACT"></a>

 ### MsgMintACT
 MsgMintACT defines the message for burning one token to mint another
Allows burning AKT to mint ACT, or burning unused ACT back to AKT

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | owner source of the coins to be burned |
 | `to` | [string](#string) |  | to destination of the minted coins. if minted coin is ACT, "to" must be same as signer |
 | `coins_to_burn` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | coins_to_burn |
 
 

 

 
 <a name="akash.bme.v1.MsgMintACTResponse"></a>

 ### MsgMintACTResponse
 MsgBurnMintResponse is the response type for MsgBurnMint

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [LedgerRecordID](#akash.bme.v1.LedgerRecordID) |  |  |
 | `status` | [LedgerRecordStatus](#akash.bme.v1.LedgerRecordStatus) |  |  |
 
 

 

 
 <a name="akash.bme.v1.MsgUpdateParams"></a>

 ### MsgUpdateParams
 MsgUpdateParams defines the message for updating module parameters

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `authority` | [string](#string) |  | authority is the address that controls the module (governance) |
 | `params` | [Params](#akash.bme.v1.Params) |  | params defines the updated parameters |
 
 

 

 
 <a name="akash.bme.v1.MsgUpdateParamsResponse"></a>

 ### MsgUpdateParamsResponse
 MsgUpdateParamsResponse is the response type for MsgUpdateParams

 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/bme/v1/query.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/bme/v1/query.proto
 

 
 <a name="akash.bme.v1.QueryLedgerRecordEntry"></a>

 ### QueryLedgerRecordEntry
 QueryLedgerRecordEntry wraps a ledger record with its ID and status

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [LedgerRecordID](#akash.bme.v1.LedgerRecordID) |  | id is the unique identifier of the ledger record |
 | `status` | [LedgerRecordStatus](#akash.bme.v1.LedgerRecordStatus) |  | status indicates whether this record is pending or executed |
 | `pending_record` | [LedgerPendingRecord](#akash.bme.v1.LedgerPendingRecord) |  | pending_record is set when the record status is pending |
 | `executed_record` | [LedgerRecord](#akash.bme.v1.LedgerRecord) |  | executed_record is set when the record status is executed |
 | `canceled_record` | [LedgerCanceledRecord](#akash.bme.v1.LedgerCanceledRecord) |  | canceled_record is set when the record status is failed |
 
 

 

 
 <a name="akash.bme.v1.QueryLedgerRecordsRequest"></a>

 ### QueryLedgerRecordsRequest
 QueryLedgerRecordsRequest is the request type for the Query/LedgerRecords RPC method

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `filters` | [LedgerRecordFilters](#akash.bme.v1.LedgerRecordFilters) |  | filters holds the ledger record fields to filter the request |
 | `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines the pagination for the request |
 
 

 

 
 <a name="akash.bme.v1.QueryLedgerRecordsResponse"></a>

 ### QueryLedgerRecordsResponse
 QueryLedgerRecordsResponse is the response type for the Query/LedgerRecords RPC method

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `records` | [QueryLedgerRecordEntry](#akash.bme.v1.QueryLedgerRecordEntry) | repeated | records is a list of ledger records matching the filters |
 | `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination contains the information about response pagination |
 
 

 

 
 <a name="akash.bme.v1.QueryParamsRequest"></a>

 ### QueryParamsRequest
 QueryParamsRequest is the request type for the Query/Params RPC method

 

 

 
 <a name="akash.bme.v1.QueryParamsResponse"></a>

 ### QueryParamsResponse
 QueryParamsResponse is the response type for the Query/Params RPC method

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `params` | [Params](#akash.bme.v1.Params) |  |  |
 
 

 

 
 <a name="akash.bme.v1.QueryStatusRequest"></a>

 ### QueryStatusRequest
 QueryStatusRequest is the request type for the circuit breaker status

 

 

 
 <a name="akash.bme.v1.QueryStatusResponse"></a>

 ### QueryStatusResponse
 QueryMintStatusResponse is the response type for the circuit breaker status

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `status` | [MintStatus](#akash.bme.v1.MintStatus) |  | status is the current circuit breaker status |
 | `collateral_ratio` | [string](#string) |  | collateral_ratio is the current CR |
 | `warn_threshold` | [string](#string) |  | warn_threshold is the warning threshold |
 | `halt_threshold` | [string](#string) |  | halt_threshold is the halt threshold |
 | `mints_allowed` | [bool](#bool) |  | mints_allowed indicates if new ACT mints are allowed |
 | `refunds_allowed` | [bool](#bool) |  | refunds_allowed indicates if ACT refunds are allowed |
 
 

 

 
 <a name="akash.bme.v1.QueryVaultStateRequest"></a>

 ### QueryVaultStateRequest
 QueryVaultStateRequest is the request type for the Query/VaultState RPC method

 

 

 
 <a name="akash.bme.v1.QueryVaultStateResponse"></a>

 ### QueryVaultStateResponse
 QueryVaultStateResponse is the response type for the Query/VaultState RPC method

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `vault_state` | [State](#akash.bme.v1.State) |  |  |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.bme.v1.Query"></a>

 ### Query
 Query defines the gRPC querier service for the BME module

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `Params` | [QueryParamsRequest](#akash.bme.v1.QueryParamsRequest) | [QueryParamsResponse](#akash.bme.v1.QueryParamsResponse) | Params returns the module parameters | GET|/akash/bme/v1/params|
 | `VaultState` | [QueryVaultStateRequest](#akash.bme.v1.QueryVaultStateRequest) | [QueryVaultStateResponse](#akash.bme.v1.QueryVaultStateResponse) | VaultState returns the current vault state | GET|/akash/bme/v1/vault|
 | `Status` | [QueryStatusRequest](#akash.bme.v1.QueryStatusRequest) | [QueryStatusResponse](#akash.bme.v1.QueryStatusResponse) | Status returns the current circuit breaker status | GET|/akash/bme/v1/status|
 | `LedgerRecords` | [QueryLedgerRecordsRequest](#akash.bme.v1.QueryLedgerRecordsRequest) | [QueryLedgerRecordsResponse](#akash.bme.v1.QueryLedgerRecordsResponse) | LedgerRecords queries ledger records with optional filters for status, source, denom, to_denom | GET|/akash/bme/v1/ledger|
 
  <!-- end services -->

 
 
 <a name="akash/bme/v1/service.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/bme/v1/service.proto
 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.bme.v1.Msg"></a>

 ### Msg
 Msg defines the BME (Burn/Mint Engine) transaction service.
The BME module manages the burn and mint operations for ACT tokens,
maintaining collateral ratios and enforcing circuit breaker rules.

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `UpdateParams` | [MsgUpdateParams](#akash.bme.v1.MsgUpdateParams) | [MsgUpdateParamsResponse](#akash.bme.v1.MsgUpdateParamsResponse) | UpdateParams updates the module parameters. This operation can only be performed through governance proposals. | |
 | `BurnMint` | [MsgBurnMint](#akash.bme.v1.MsgBurnMint) | [MsgBurnMintResponse](#akash.bme.v1.MsgBurnMintResponse) | BurnMint allows users to burn one token and mint another at current oracle prices. Typically used to burn unused ACT tokens back to AKT. The operation may be delayed or rejected based on circuit breaker status. | |
 | `MintACT` | [MsgMintACT](#akash.bme.v1.MsgMintACT) | [MsgMintACTResponse](#akash.bme.v1.MsgMintACTResponse) | MintACT mints ACT tokens by burning the specified source token. The mint amount is calculated based on current oracle prices and the collateral ratio. May be halted if circuit breaker is triggered. | |
 | `BurnACT` | [MsgBurnACT](#akash.bme.v1.MsgBurnACT) | [MsgBurnACTResponse](#akash.bme.v1.MsgBurnACTResponse) | BurnACT burns ACT tokens and mints the specified destination token. The burn operation uses remint credits when available, otherwise requires adequate collateral backing based on oracle prices. | |
 | `FundVault` | [MsgFundVault](#akash.bme.v1.MsgFundVault) | [MsgFundVaultResponse](#akash.bme.v1.MsgFundVaultResponse) | FundVault seeds the BME vault with AKT from a designated source (e.g., community pool). This provides the initial volatility buffer required for burn/mint operations. Can only be executed through governance proposals. | |
 
  <!-- end services -->

 
 
 <a name="akash/cert/v1/cert.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/cert/v1/cert.proto
 

 
 <a name="akash.cert.v1.Certificate"></a>

 ### Certificate
 Certificate stores state, certificate and it's public key.
The certificate is required for several transactions including deployment of a workload to verify the identity of the tenant and secure the deployment.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `state` | [State](#akash.cert.v1.State) |  | State is the state of the certificate. CertificateValid denotes state for deployment active. CertificateRevoked denotes state for deployment closed. |
 | `cert` | [bytes](#bytes) |  | Cert holds the bytes of the certificate. |
 | `pubkey` | [bytes](#bytes) |  | PubKey holds the public key of the certificate. |
 
 

 

 
 <a name="akash.cert.v1.ID"></a>

 ### ID
 ID stores owner and sequence number.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the account address of the user who owns the certificate. It is a string representing a valid account address.

Example: "akash1..." |
 | `serial` | [string](#string) |  | Serial is a sequence number for the certificate. |
 
 

 

  <!-- end messages -->

 
 <a name="akash.cert.v1.State"></a>

 ### State
 State is an enum which refers to state of the certificate.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | invalid | 0 | Prefix should start with 0 in enum. So declaring dummy state. |
 | valid | 1 | CertificateValid denotes state for deployment active. |
 | revoked | 2 | CertificateRevoked denotes state for deployment closed. |
 

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/cert/v1/filters.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/cert/v1/filters.proto
 

 
 <a name="akash.cert.v1.CertificateFilter"></a>

 ### CertificateFilter
 CertificateFilter defines filters used to filter certificates.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the account address of the user who owns the certificate. It is a string representing a valid account address.

Example: "akash1..." |
 | `serial` | [string](#string) |  | Serial is a sequence number for the certificate. |
 | `state` | [string](#string) |  | State is the state of the certificate. CertificateValid denotes state for deployment active. CertificateRevoked denotes state for deployment closed. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/cert/v1/genesis.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/cert/v1/genesis.proto
 

 
 <a name="akash.cert.v1.GenesisCertificate"></a>

 ### GenesisCertificate
 GenesisCertificate defines certificate entry at genesis.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the account address of the user who owns the certificate. It is a string representing a valid account address.

Example: "akash1..." |
 | `certificate` | [Certificate](#akash.cert.v1.Certificate) |  | Certificate holds the certificate. |
 
 

 

 
 <a name="akash.cert.v1.GenesisState"></a>

 ### GenesisState
 GenesisState defines the basic genesis state used by cert module.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `certificates` | [GenesisCertificate](#akash.cert.v1.GenesisCertificate) | repeated | Certificates is a list of genesis certificates. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/cert/v1/msg.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/cert/v1/msg.proto
 

 
 <a name="akash.cert.v1.MsgCreateCertificate"></a>

 ### MsgCreateCertificate
 MsgCreateCertificate defines an SDK message for creating certificate.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the account address of the user who owns the certificate. It is a string representing a valid account address.

Example: "akash1..." |
 | `cert` | [bytes](#bytes) |  | Cert holds the bytes representing the certificate. |
 | `pubkey` | [bytes](#bytes) |  | PubKey holds the public key. |
 
 

 

 
 <a name="akash.cert.v1.MsgCreateCertificateResponse"></a>

 ### MsgCreateCertificateResponse
 MsgCreateCertificateResponse defines the Msg/CreateCertificate response type.

 

 

 
 <a name="akash.cert.v1.MsgRevokeCertificate"></a>

 ### MsgRevokeCertificate
 MsgRevokeCertificate defines an SDK message for revoking certificate.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [ID](#akash.cert.v1.ID) |  | Id corresponds to the certificate ID which includes owner and sequence number. |
 
 

 

 
 <a name="akash.cert.v1.MsgRevokeCertificateResponse"></a>

 ### MsgRevokeCertificateResponse
 MsgRevokeCertificateResponse defines the Msg/RevokeCertificate response type.

 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/cert/v1/query.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/cert/v1/query.proto
 

 
 <a name="akash.cert.v1.CertificateResponse"></a>

 ### CertificateResponse
 CertificateResponse contains a single X509 certificate and its serial number.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `certificate` | [Certificate](#akash.cert.v1.Certificate) |  | Certificate holds the certificate. |
 | `serial` | [string](#string) |  | Serial is a sequence number for the certificate. |
 
 

 

 
 <a name="akash.cert.v1.QueryCertificatesRequest"></a>

 ### QueryCertificatesRequest
 QueryDeploymentsRequest is request type for the Query/Deployments RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `filter` | [CertificateFilter](#akash.cert.v1.CertificateFilter) |  | Filter allows for filtering of results. |
 | `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | Pagination is used to paginate the request. |
 
 

 

 
 <a name="akash.cert.v1.QueryCertificatesResponse"></a>

 ### QueryCertificatesResponse
 QueryCertificatesResponse is response type for the Query/Certificates RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `certificates` | [CertificateResponse](#akash.cert.v1.CertificateResponse) | repeated | Certificates is a list of certificate. |
 | `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | Pagination contains the information about response pagination. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.cert.v1.Query"></a>

 ### Query
 Query defines the gRPC querier service for certificates.

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `Certificates` | [QueryCertificatesRequest](#akash.cert.v1.QueryCertificatesRequest) | [QueryCertificatesResponse](#akash.cert.v1.QueryCertificatesResponse) | Certificates queries certificates on-chain. | GET|/akash/cert/v1/certificates/list|
 
  <!-- end services -->

 
 
 <a name="akash/cert/v1/service.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/cert/v1/service.proto
 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.cert.v1.Msg"></a>

 ### Msg
 Msg defines the provider Msg service.

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `CreateCertificate` | [MsgCreateCertificate](#akash.cert.v1.MsgCreateCertificate) | [MsgCreateCertificateResponse](#akash.cert.v1.MsgCreateCertificateResponse) | CreateCertificate defines a method to create new certificate given proper inputs. | |
 | `RevokeCertificate` | [MsgRevokeCertificate](#akash.cert.v1.MsgRevokeCertificate) | [MsgRevokeCertificateResponse](#akash.cert.v1.MsgRevokeCertificateResponse) | RevokeCertificate defines a method to revoke the certificate. | |
 
  <!-- end services -->

 
 
 <a name="akash/deployment/v1/deployment.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/deployment/v1/deployment.proto
 

 
 <a name="akash.deployment.v1.Deployment"></a>

 ### Deployment
 Deployment stores deploymentID, state and checksum details.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [DeploymentID](#akash.deployment.v1.DeploymentID) |  | ID is the unique identifier of the deployment. |
 | `state` | [Deployment.State](#akash.deployment.v1.Deployment.State) |  | State defines the sate of the deployment. A deployment can be either active or inactive. |
 | `hash` | [bytes](#bytes) |  | Hash is an hashed representation of the deployment. |
 | `created_at` | [int64](#int64) |  | CreatedAt indicates when the deployment was created as a block height value. |
 | `reclamation` | [DeploymentReclamation](#akash.deployment.v1.DeploymentReclamation) |  | reclamation stores the deployment's reclamation requirements for persistence. Needed so that StartGroup can propagate reclamation to newly created orders. |
 
 

 

 
 <a name="akash.deployment.v1.DeploymentID"></a>

 ### DeploymentID
 DeploymentID represents a unique identifier for a specific deployment on the network.
It is composed of two fields: an owner address and a sequence number (dseq).

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the account bech32 address of the user who owns the deployment. It is a string representing a valid bech32 account address.

Example: "akash1..." |
 | `dseq` | [uint64](#uint64) |  | Dseq (deployment sequence number) is a unique numeric identifier for the deployment. It is used to differentiate deployments created by the same owner. |
 
 

 

 
 <a name="akash.deployment.v1.DeploymentReclamation"></a>

 ### DeploymentReclamation
 DeploymentReclamation defines the tenant's reclamation requirements.
Stored on the Deployment and propagated to Orders.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `min_window` | [google.protobuf.Duration](#google.protobuf.Duration) |  | min_window is the minimum reclamation window the tenant requires. |
 
 

 

  <!-- end messages -->

 
 <a name="akash.deployment.v1.Deployment.State"></a>

 ### Deployment.State
 State is an enum which refers to state of deployment.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | invalid | 0 | Prefix should start with 0 in enum. So declaring dummy state. |
 | active | 1 | DeploymentActive denotes state for deployment active. |
 | closed | 2 | DeploymentClosed denotes state for deployment closed. |
 

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/deployment/v1/group.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/deployment/v1/group.proto
 

 
 <a name="akash.deployment.v1.GroupID"></a>

 ### GroupID
 GroupID uniquely identifies a group within a deployment on the network.
A group represents a specific collection of resources or configurations
within a deployment.
It stores owner, deployment sequence number (dseq) and group sequence number (gseq).

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the account address of the user who owns the group. It is a string representing a valid account address.

Example: "akash1..." |
 | `dseq` | [uint64](#uint64) |  | Dseq (deployment sequence number) is a unique numeric identifier for the deployment. It is used to differentiate deployments created by the same owner. |
 | `gseq` | [uint32](#uint32) |  | Gseq (group sequence number) is a unique numeric identifier for the group. It is used to differentiate groups created by the same owner in a deployment. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/deployment/v1/event.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/deployment/v1/event.proto
 

 
 <a name="akash.deployment.v1.EventDeploymentClosed"></a>

 ### EventDeploymentClosed
 EventDeploymentClosed is triggered when deployment is closed on chain.
It contains all the information required to identify a deployment.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [DeploymentID](#akash.deployment.v1.DeploymentID) |  | ID is the unique identifier of the deployment. |
 
 

 

 
 <a name="akash.deployment.v1.EventDeploymentCreated"></a>

 ### EventDeploymentCreated
 EventDeploymentCreated event is triggered when deployment is created on chain.
It contains all the information required to identify a deployment.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [DeploymentID](#akash.deployment.v1.DeploymentID) |  | ID is the unique identifier of the deployment. |
 | `hash` | [bytes](#bytes) |  | Hash is an hashed representation of the deployment. |
 
 

 

 
 <a name="akash.deployment.v1.EventDeploymentUpdated"></a>

 ### EventDeploymentUpdated
 EventDeploymentUpdated is triggered when deployment is updated on chain.
It contains all the information required to identify a deployment.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [DeploymentID](#akash.deployment.v1.DeploymentID) |  | ID is the unique identifier of the deployment. |
 | `hash` | [bytes](#bytes) |  | Hash is an hashed representation of the deployment. |
 
 

 

 
 <a name="akash.deployment.v1.EventGroupClosed"></a>

 ### EventGroupClosed
 EventGroupClosed is triggered when deployment group is closed.
It contains all the information required to identify a group.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [GroupID](#akash.deployment.v1.GroupID) |  | ID is the unique identifier of the group. |
 
 

 

 
 <a name="akash.deployment.v1.EventGroupPaused"></a>

 ### EventGroupPaused
 EventGroupPaused is triggered when deployment group is paused.
It contains all the information required to identify a group.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [GroupID](#akash.deployment.v1.GroupID) |  | ID is the unique identifier of the group. |
 
 

 

 
 <a name="akash.deployment.v1.EventGroupStarted"></a>

 ### EventGroupStarted
 EventGroupStarted is triggered when deployment group is started.
It contains all the information required to identify a group.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [GroupID](#akash.deployment.v1.GroupID) |  | ID is the unique identifier of the group. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/deployment/v1beta4/resourceunit.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/deployment/v1beta4/resourceunit.proto
 

 
 <a name="akash.deployment.v1beta4.ResourceUnit"></a>

 ### ResourceUnit
 ResourceUnit extends Resources and adds Count along with the Price.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `resource` | [akash.base.resources.v1beta4.Resources](#akash.base.resources.v1beta4.Resources) |  | Resource holds the amount of resources. |
 | `count` | [uint32](#uint32) |  | Count corresponds to the amount of replicas to run of the resources. |
 | `price` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | Price holds the pricing for the resource units. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/deployment/v1beta4/groupspec.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/deployment/v1beta4/groupspec.proto
 

 
 <a name="akash.deployment.v1beta4.GroupSpec"></a>

 ### GroupSpec
 GroupSpec defines a specification for a group in a deployment on the network.
This includes attributes like the group name, placement requirements, and resource units.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `name` | [string](#string) |  | Name is the name of the group. |
 | `requirements` | [akash.base.attributes.v1.PlacementRequirements](#akash.base.attributes.v1.PlacementRequirements) |  | Requirements specifies the placement requirements for the group. This determines where the resources in the group can be deployed. |
 | `resources` | [ResourceUnit](#akash.deployment.v1beta4.ResourceUnit) | repeated | Resources is a list containing the resource units allocated to the group. Each ResourceUnit defines the specific resources (e.g., CPU, memory) assigned. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/deployment/v1beta4/deploymentmsg.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/deployment/v1beta4/deploymentmsg.proto
 

 
 <a name="akash.deployment.v1beta4.MsgCloseDeployment"></a>

 ### MsgCloseDeployment
 MsgCloseDeployment defines an SDK message for closing deployment

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [akash.deployment.v1.DeploymentID](#akash.deployment.v1.DeploymentID) |  | ID is the unique identifier of the deployment. |
 
 

 

 
 <a name="akash.deployment.v1beta4.MsgCloseDeploymentResponse"></a>

 ### MsgCloseDeploymentResponse
 MsgCloseDeploymentResponse defines the Msg/CloseDeployment response type.

 

 

 
 <a name="akash.deployment.v1beta4.MsgCreateDeployment"></a>

 ### MsgCreateDeployment
 MsgCreateDeployment defines an SDK message for creating deployment.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [akash.deployment.v1.DeploymentID](#akash.deployment.v1.DeploymentID) |  | ID is the unique identifier of the deployment. |
 | `groups` | [GroupSpec](#akash.deployment.v1beta4.GroupSpec) | repeated | GroupSpec is a list of group specifications for the deployment. This field is required and must be a list of GroupSpec. |
 | `hash` | [bytes](#bytes) |  | Hash of the deployment. |
 | `deposit` | [akash.base.deposit.v1.Deposit](#akash.base.deposit.v1.Deposit) |  | Deposit specifies the amount of coins to include in the deployment's first deposit. |
 | `reclamation` | [akash.deployment.v1.DeploymentReclamation](#akash.deployment.v1.DeploymentReclamation) |  | reclamation specifies the deployment-level reclamation requirements. Nil means the tenant does not require reclamation. |
 
 

 

 
 <a name="akash.deployment.v1beta4.MsgCreateDeploymentResponse"></a>

 ### MsgCreateDeploymentResponse
 MsgCreateDeploymentResponse defines the Msg/CreateDeployment response type.

 

 

 
 <a name="akash.deployment.v1beta4.MsgUpdateDeployment"></a>

 ### MsgUpdateDeployment
 MsgUpdateDeployment defines an SDK message for updating deployment.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [akash.deployment.v1.DeploymentID](#akash.deployment.v1.DeploymentID) |  | ID is the unique identifier of the deployment. |
 | `hash` | [bytes](#bytes) |  | Hash of the deployment. |
 
 

 

 
 <a name="akash.deployment.v1beta4.MsgUpdateDeploymentResponse"></a>

 ### MsgUpdateDeploymentResponse
 MsgUpdateDeploymentResponse defines the Msg/UpdateDeployment response type.

 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/deployment/v1beta4/filters.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/deployment/v1beta4/filters.proto
 

 
 <a name="akash.deployment.v1beta4.DeploymentFilters"></a>

 ### DeploymentFilters
 DeploymentFilters defines filters used to filter deployments.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the account bech32 address of the user who owns the deployment. It is a string representing a valid bech32 account address.

Example: "akash1..." |
 | `dseq` | [uint64](#uint64) |  | Dseq (deployment sequence number) is a unique numeric identifier for the deployment. It is used to differentiate deployments created by the same owner. |
 | `state` | [string](#string) |  | State defines the sate of the deployment. A deployment can be either active or inactive. |
 
 

 

 
 <a name="akash.deployment.v1beta4.GroupFilters"></a>

 ### GroupFilters
 GroupFilters defines filters used to filter groups

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the account address of the user who owns the group. It is a string representing a valid account address.

Example: "akash1..." |
 | `dseq` | [uint64](#uint64) |  | Dseq (deployment sequence number) is a unique numeric identifier for the deployment. It is used to differentiate deployments created by the same owner. |
 | `gseq` | [uint64](#uint64) |  | Gseq (group sequence number) is a unique numeric identifier for the group. It is used to differentiate groups created by the same owner in a deployment. |
 | `state` | [string](#string) |  | State defines the sate of the deployment. A deployment can be either active or inactive. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/deployment/v1beta4/group.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/deployment/v1beta4/group.proto
 

 
 <a name="akash.deployment.v1beta4.Group"></a>

 ### Group
 Group stores group id, state and specifications of a group.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [akash.deployment.v1.GroupID](#akash.deployment.v1.GroupID) |  | Id is the unique identifier for the group. |
 | `state` | [Group.State](#akash.deployment.v1beta4.Group.State) |  | State represents the state of the group. |
 | `group_spec` | [GroupSpec](#akash.deployment.v1beta4.GroupSpec) |  | GroupSpec holds the specification of a the Group. |
 | `created_at` | [int64](#int64) |  | CreatedAt is the block height at which the deployment was created. |
 
 

 

  <!-- end messages -->

 
 <a name="akash.deployment.v1beta4.Group.State"></a>

 ### Group.State
 State is an enum which refers to state of group.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | invalid | 0 | Prefix should start with 0 in enum. So declaring dummy state. |
 | open | 1 | GroupOpen denotes state for group open. |
 | paused | 2 | GroupOrdered denotes state for group ordered. |
 | insufficient_funds | 3 | GroupInsufficientFunds denotes state for group insufficient_funds. |
 | closed | 4 | GroupClosed denotes state for group closed. |
 

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/deployment/v1beta4/params.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/deployment/v1beta4/params.proto
 

 
 <a name="akash.deployment.v1beta4.Params"></a>

 ### Params
 Params defines the parameters for the x/deployment module.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `min_deposits` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | MinDeposits holds a list of the minimum amount of deposits for each a coin. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/deployment/v1beta4/genesis.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/deployment/v1beta4/genesis.proto
 

 
 <a name="akash.deployment.v1beta4.GenesisDeployment"></a>

 ### GenesisDeployment
 GenesisDeployment defines the basic genesis state used by deployment module.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `deployment` | [akash.deployment.v1.Deployment](#akash.deployment.v1.Deployment) |  | Deployments represents a deployment on the network. |
 | `groups` | [Group](#akash.deployment.v1beta4.Group) | repeated | Groups is a list of groups within a Deployment. |
 
 

 

 
 <a name="akash.deployment.v1beta4.GenesisState"></a>

 ### GenesisState
 GenesisState stores slice of genesis deployment instance.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `deployments` | [GenesisDeployment](#akash.deployment.v1beta4.GenesisDeployment) | repeated | Deployments is a list of deployments on the network. |
 | `params` | [Params](#akash.deployment.v1beta4.Params) |  | Params defines the parameters for the x/deployment module. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/deployment/v1beta4/groupmsg.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/deployment/v1beta4/groupmsg.proto
 

 
 <a name="akash.deployment.v1beta4.MsgCloseGroup"></a>

 ### MsgCloseGroup
 MsgCloseGroup defines SDK message to close a single Group within a Deployment.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [akash.deployment.v1.GroupID](#akash.deployment.v1.GroupID) |  | Id is the unique identifier of the Group. |
 
 

 

 
 <a name="akash.deployment.v1beta4.MsgCloseGroupResponse"></a>

 ### MsgCloseGroupResponse
 MsgCloseGroupResponse defines the Msg/CloseGroup response type.

 

 

 
 <a name="akash.deployment.v1beta4.MsgPauseGroup"></a>

 ### MsgPauseGroup
 MsgPauseGroup defines SDK message to pause a single Group within a Deployment.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [akash.deployment.v1.GroupID](#akash.deployment.v1.GroupID) |  | Id is the unique identifier of the Group. |
 
 

 

 
 <a name="akash.deployment.v1beta4.MsgPauseGroupResponse"></a>

 ### MsgPauseGroupResponse
 MsgPauseGroupResponse defines the Msg/PauseGroup response type.

 

 

 
 <a name="akash.deployment.v1beta4.MsgStartGroup"></a>

 ### MsgStartGroup
 MsgStartGroup defines SDK message to start a single Group within a Deployment.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [akash.deployment.v1.GroupID](#akash.deployment.v1.GroupID) |  | Id is the unique identifier of the Group. |
 
 

 

 
 <a name="akash.deployment.v1beta4.MsgStartGroupResponse"></a>

 ### MsgStartGroupResponse
 MsgStartGroupResponse defines the Msg/StartGroup response type.

 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/deployment/v1beta4/paramsmsg.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/deployment/v1beta4/paramsmsg.proto
 

 
 <a name="akash.deployment.v1beta4.MsgUpdateParams"></a>

 ### MsgUpdateParams
 MsgUpdateParams is the Msg/UpdateParams request type.

Since: akash v1.0.0

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `authority` | [string](#string) |  | Authority is the address of the governance account. |
 | `params` | [Params](#akash.deployment.v1beta4.Params) |  | Params defines the x/deployment parameters to update.

NOTE: All parameters must be supplied. |
 
 

 

 
 <a name="akash.deployment.v1beta4.MsgUpdateParamsResponse"></a>

 ### MsgUpdateParamsResponse
 MsgUpdateParamsResponse defines the response structure for executing a
MsgUpdateParams message.

Since: akash v1.0.0

 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/escrow/id/v1/id.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/escrow/id/v1/id.proto
 

 
 <a name="akash.escrow.id.v1.Account"></a>

 ### Account
 Account is the account identifier.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `scope` | [Scope](#akash.escrow.id.v1.Scope) |  |  |
 | `xid` | [string](#string) |  |  |
 
 

 

 
 <a name="akash.escrow.id.v1.Payment"></a>

 ### Payment
 Payment is the payment identifier.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `aid` | [Account](#akash.escrow.id.v1.Account) |  |  |
 | `xid` | [string](#string) |  |  |
 
 

 

  <!-- end messages -->

 
 <a name="akash.escrow.id.v1.Scope"></a>

 ### Scope
 Scope is an enum which refers to the account scope

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | invalid | 0 | Prefix should start with 0 in enum. So declaring dummy state. |
 | deployment | 1 | DeploymentActive denotes state for deployment active. |
 | bid | 2 | DeploymentClosed denotes state for deployment closed. |
 

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/escrow/types/v1/balance.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/escrow/types/v1/balance.proto
 

 
 <a name="akash.escrow.types.v1.Balance"></a>

 ### Balance
 Balance holds the unspent coin received from all deposits with same denom
DecCoin is not being used here as it does not support negative values,
and balance may go negative if account is overdrawn.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `denom` | [string](#string) |  |  |
 | `amount` | [string](#string) |  |  |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/escrow/types/v1/deposit.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/escrow/types/v1/deposit.proto
 

 
 <a name="akash.escrow.types.v1.Depositor"></a>

 ### Depositor
 Depositor stores state of a deposit.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the bech32 address of the depositor. It is a string representing a valid account address.

Example: "akash1..." If depositor is same as the owner, then any incoming coins are added to the Balance. If depositor isn't same as the owner, then any incoming coins are added to the Funds. |
 | `height` | [int64](#int64) |  | Height blockchain height at which deposit was created |
 | `source` | [akash.base.deposit.v1.Source](#akash.base.deposit.v1.Source) |  | Source indicated origination of the funds |
 | `balance` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | Balance amount of funds available to spend in this deposit. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/escrow/types/v1/state.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/escrow/types/v1/state.proto
 

  <!-- end messages -->

 
 <a name="akash.escrow.types.v1.State"></a>

 ### State
 State stores state for an escrow account.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | invalid | 0 | AccountStateInvalid is an invalid state. |
 | open | 1 | StateOpen is the state when an object is open. |
 | closed | 2 | StateClosed is the state when an object is closed. |
 | overdrawn | 3 | StateOverdrawn is the state when an object are overdrawn. |
 

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/escrow/types/v1/account.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/escrow/types/v1/account.proto
 

 
 <a name="akash.escrow.types.v1.Account"></a>

 ### Account
 Account

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [akash.escrow.id.v1.Account](#akash.escrow.id.v1.Account) |  |  |
 | `state` | [AccountState](#akash.escrow.types.v1.AccountState) |  |  |
 
 

 

 
 <a name="akash.escrow.types.v1.AccountState"></a>

 ### AccountState
 Account stores state for an escrow account.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the account bech32 address of the user who owns the deployment. It is a string representing a valid bech32 account address.

Example: "akash1..." |
 | `state` | [State](#akash.escrow.types.v1.State) |  | State represents the current state of an Account. |
 | `transferred` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated | Transferred total coins spent by this account. |
 | `settled_at` | [int64](#int64) |  | SettledAt represents the block height at which this account was last settled. |
 | `funds` | [Balance](#akash.escrow.types.v1.Balance) | repeated | Funds holds the unspent coins received from all deposits |
 | `deposits` | [Depositor](#akash.escrow.types.v1.Depositor) | repeated |  |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/deployment/v1beta4/query.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/deployment/v1beta4/query.proto
 

 
 <a name="akash.deployment.v1beta4.QueryDeploymentRequest"></a>

 ### QueryDeploymentRequest
 QueryDeploymentRequest is request type for the Query/Deployment RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [akash.deployment.v1.DeploymentID](#akash.deployment.v1.DeploymentID) |  | Id is the unique identifier of the deployment. |
 
 

 

 
 <a name="akash.deployment.v1beta4.QueryDeploymentResponse"></a>

 ### QueryDeploymentResponse
 QueryDeploymentResponse is response type for the Query/Deployment RPC method

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `deployment` | [akash.deployment.v1.Deployment](#akash.deployment.v1.Deployment) |  | Deployment represents a deployment on the network. |
 | `groups` | [Group](#akash.deployment.v1beta4.Group) | repeated | Groups is a list of deployment groups. |
 | `escrow_account` | [akash.escrow.types.v1.Account](#akash.escrow.types.v1.Account) |  | EscrowAccount represents an escrow mechanism where funds are held. This ensures that obligations of both tenants and providers involved in the transaction are met without direct access to each other's accounts. |
 
 

 

 
 <a name="akash.deployment.v1beta4.QueryDeploymentsRequest"></a>

 ### QueryDeploymentsRequest
 QueryDeploymentsRequest is request type for the Query/Deployments RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `filters` | [DeploymentFilters](#akash.deployment.v1beta4.DeploymentFilters) |  | Filters holds the deployment fields to filter the request. |
 | `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | Pagination is used to paginate the request. |
 
 

 

 
 <a name="akash.deployment.v1beta4.QueryDeploymentsResponse"></a>

 ### QueryDeploymentsResponse
 QueryDeploymentsResponse is response type for the Query/Deployments RPC method

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `deployments` | [QueryDeploymentResponse](#akash.deployment.v1beta4.QueryDeploymentResponse) | repeated | Deployments is a list of Deployments. |
 | `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | Pagination contains the information about response pagination. |
 
 

 

 
 <a name="akash.deployment.v1beta4.QueryGroupRequest"></a>

 ### QueryGroupRequest
 QueryGroupRequest is request type for the Query/Group RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [akash.deployment.v1.GroupID](#akash.deployment.v1.GroupID) |  | Id is the unique identifier of the Group. |
 
 

 

 
 <a name="akash.deployment.v1beta4.QueryGroupResponse"></a>

 ### QueryGroupResponse
 QueryGroupResponse is response type for the Query/Group RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `group` | [Group](#akash.deployment.v1beta4.Group) |  | Group holds a deployment Group. |
 
 

 

 
 <a name="akash.deployment.v1beta4.QueryParamsRequest"></a>

 ### QueryParamsRequest
 QueryParamsRequest is the request type for the Query/Params RPC method.

 

 

 
 <a name="akash.deployment.v1beta4.QueryParamsResponse"></a>

 ### QueryParamsResponse
 QueryParamsResponse is the response type for the Query/Params RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `params` | [Params](#akash.deployment.v1beta4.Params) |  | params defines the parameters of the module. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.deployment.v1beta4.Query"></a>

 ### Query
 Query defines the gRPC querier service for the Deployments package.

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `Deployments` | [QueryDeploymentsRequest](#akash.deployment.v1beta4.QueryDeploymentsRequest) | [QueryDeploymentsResponse](#akash.deployment.v1beta4.QueryDeploymentsResponse) | Deployments queries deployments. | GET|/akash/deployment/v1beta4/deployments/list|
 | `Deployment` | [QueryDeploymentRequest](#akash.deployment.v1beta4.QueryDeploymentRequest) | [QueryDeploymentResponse](#akash.deployment.v1beta4.QueryDeploymentResponse) | Deployment queries deployment details. | GET|/akash/deployment/v1beta4/deployments/info|
 | `Group` | [QueryGroupRequest](#akash.deployment.v1beta4.QueryGroupRequest) | [QueryGroupResponse](#akash.deployment.v1beta4.QueryGroupResponse) | Group queries group details. | GET|/akash/deployment/v1beta4/groups/info|
 | `Params` | [QueryParamsRequest](#akash.deployment.v1beta4.QueryParamsRequest) | [QueryParamsResponse](#akash.deployment.v1beta4.QueryParamsResponse) | Params returns the total set of deployment parameters. | GET|/akash/deployment/v1beta4/params|
 
  <!-- end services -->

 
 
 <a name="akash/deployment/v1beta4/service.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/deployment/v1beta4/service.proto
 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.deployment.v1beta4.Msg"></a>

 ### Msg
 Msg defines the x/deployment Msg service.

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `CreateDeployment` | [MsgCreateDeployment](#akash.deployment.v1beta4.MsgCreateDeployment) | [MsgCreateDeploymentResponse](#akash.deployment.v1beta4.MsgCreateDeploymentResponse) | CreateDeployment defines a method to create new deployment given proper inputs. | |
 | `UpdateDeployment` | [MsgUpdateDeployment](#akash.deployment.v1beta4.MsgUpdateDeployment) | [MsgUpdateDeploymentResponse](#akash.deployment.v1beta4.MsgUpdateDeploymentResponse) | UpdateDeployment defines a method to update a deployment given proper inputs. | |
 | `CloseDeployment` | [MsgCloseDeployment](#akash.deployment.v1beta4.MsgCloseDeployment) | [MsgCloseDeploymentResponse](#akash.deployment.v1beta4.MsgCloseDeploymentResponse) | CloseDeployment defines a method to close a deployment given proper inputs. | |
 | `CloseGroup` | [MsgCloseGroup](#akash.deployment.v1beta4.MsgCloseGroup) | [MsgCloseGroupResponse](#akash.deployment.v1beta4.MsgCloseGroupResponse) | CloseGroup defines a method to close a group of a deployment given proper inputs. | |
 | `PauseGroup` | [MsgPauseGroup](#akash.deployment.v1beta4.MsgPauseGroup) | [MsgPauseGroupResponse](#akash.deployment.v1beta4.MsgPauseGroupResponse) | PauseGroup defines a method to pause a group of a deployment given proper inputs. | |
 | `StartGroup` | [MsgStartGroup](#akash.deployment.v1beta4.MsgStartGroup) | [MsgStartGroupResponse](#akash.deployment.v1beta4.MsgStartGroupResponse) | StartGroup defines a method to start a group of a deployment given proper inputs. | |
 | `UpdateParams` | [MsgUpdateParams](#akash.deployment.v1beta4.MsgUpdateParams) | [MsgUpdateParamsResponse](#akash.deployment.v1beta4.MsgUpdateParamsResponse) | UpdateParams defines a governance operation for updating the x/deployment module parameters. The authority is hard-coded to the x/gov module account.

Since: akash v1.0.0 | |
 
  <!-- end services -->

 
 
 <a name="akash/discovery/v1/client_info.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/discovery/v1/client_info.proto
 

 
 <a name="akash.discovery.v1.ClientInfo"></a>

 ### ClientInfo
 ClientInfo is the akash specific client info.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `api_version` | [string](#string) |  | ApiVersion is the version of the API running on the client. |
 
 

 

 
 <a name="akash.discovery.v1.ModuleVersion"></a>

 ### ModuleVersion
 ModuleVersion describes a single module and its API version.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `module` | [string](#string) |  | Module is the name of the module (e.g., "deployment", "market", "oracle"). |
 | `version` | [string](#string) |  | Version is the API version of the module (e.g., "v1beta4", "v1beta5", "v2"). |
 
 

 

 
 <a name="akash.discovery.v1.VersionInfo"></a>

 ### VersionInfo
 VersionInfo describes a complete API version and its metadata.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `api_version` | [string](#string) |  | ApiVersion is the composite API version identifier (e.g., "v1beta4"). |
 | `modules` | [ModuleVersion](#akash.discovery.v1.ModuleVersion) | repeated | Modules lists the per-module versions included in this API version. |
 | `features` | [string](#string) | repeated | Features lists optional feature flags supported by this API version. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/discovery/v1/akash.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/discovery/v1/akash.proto
 

 
 <a name="akash.discovery.v1.Akash"></a>

 ### Akash
 Akash akash specific RPC parameters.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `client_info` | [ClientInfo](#akash.discovery.v1.ClientInfo) |  | ClientInfo holds information about the client. Kept for backward compatibility. New clients should use supported_versions. |
 | `supported_versions` | [VersionInfo](#akash.discovery.v1.VersionInfo) | repeated | SupportedVersions lists all API versions the node supports. Clients should pick the best match from this list. |
 | `chain_id` | [string](#string) |  | ChainID is the identifier of the blockchain network. |
 | `node_version` | [string](#string) |  | NodeVersion is the software version of the node. |
 | `min_client_version` | [string](#string) |  | MinClientVersion is the minimum client version the node accepts. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/discovery/v1/service.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/discovery/v1/service.proto
 

 
 <a name="akash.discovery.v1.GetInfoRequest"></a>

 ### GetInfoRequest
 GetInfoRequest is the request type for the Discovery/GetInfo RPC method.

 

 

 
 <a name="akash.discovery.v1.GetInfoResponse"></a>

 ### GetInfoResponse
 GetInfoResponse is the response type for the Discovery/GetInfo RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `info` | [Akash](#akash.discovery.v1.Akash) |  | Info contains the node's version and capability information. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.discovery.v1.Discovery"></a>

 ### Discovery
 Discovery provides version and capability information about the node.

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `GetInfo` | [GetInfoRequest](#akash.discovery.v1.GetInfoRequest) | [GetInfoResponse](#akash.discovery.v1.GetInfoResponse) | GetInfo returns the node's supported API versions and metadata. | GET|/akash/discovery/v1/info|
 
  <!-- end services -->

 
 
 <a name="akash/downtimedetector/v1beta1/downtime_duration.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/downtimedetector/v1beta1/downtime_duration.proto
 

  <!-- end messages -->

 
 <a name="akash.downtimedetector.v1beta1.Downtime"></a>

 ### Downtime
 Downtime defines the predefined downtime durations that can be tracked
by the downtime detector module to monitor chain availability

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | DURATION_30S | 0 | DURATION_30S represents a 30 second downtime period |
 | DURATION_1M | 1 | DURATION_1M represents a 1 minute downtime period |
 | DURATION_2M | 2 | DURATION_2M represents a 2 minute downtime period |
 | DURATION_3M | 3 | DURATION_3M represents a 3 minute downtime period |
 | DURATION_4M | 4 | DURATION_4M represents a 4 minute downtime period |
 | DURATION_5M | 5 | DURATION_5M represents a 5 minute downtime period |
 | DURATION_10M | 6 | DURATION_10M represents a 10 minute downtime period |
 | DURATION_20M | 7 | DURATION_20M represents a 20 minute downtime period |
 | DURATION_30M | 8 | DURATION_30M represents a 30 minute downtime period |
 | DURATION_40M | 9 | DURATION_40M represents a 40 minute downtime period |
 | DURATION_50M | 10 | DURATION_50M represents a 50 minute downtime period |
 | DURATION_1H | 11 | DURATION_1H represents a 1 hour downtime period |
 | DURATION_1_5H | 12 | DURATION_1_5H represents a 1.5 hour downtime period |
 | DURATION_2H | 13 | DURATION_2H represents a 2 hour downtime period |
 | DURATION_2_5H | 14 | DURATION_2_5H represents a 2.5 hour downtime period |
 | DURATION_3H | 15 | DURATION_3H represents a 3 hour downtime period |
 | DURATION_4H | 16 | DURATION_4H represents a 4 hour downtime period |
 | DURATION_5H | 17 | DURATION_5H represents a 5 hour downtime period |
 | DURATION_6H | 18 | DURATION_6H represents a 6 hour downtime period |
 | DURATION_9H | 19 | DURATION_9H represents a 9 hour downtime period |
 | DURATION_12H | 20 | DURATION_12H represents a 12 hour downtime period |
 | DURATION_18H | 21 | DURATION_18H represents a 18 hour downtime period |
 | DURATION_24H | 22 | DURATION_24H represents a 24 hour downtime period |
 | DURATION_36H | 23 | DURATION_36H represents a 36 hour downtime period |
 | DURATION_48H | 24 | DURATION_48H represents a 48 hour downtime period |
 

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/downtimedetector/v1beta1/genesis.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/downtimedetector/v1beta1/genesis.proto
 

 
 <a name="akash.downtimedetector.v1beta1.GenesisDowntimeEntry"></a>

 ### GenesisDowntimeEntry
 GenesisDowntimeEntry tracks the last occurrence of a specific downtime duration

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `duration` | [Downtime](#akash.downtimedetector.v1beta1.Downtime) |  | duration is the downtime period being tracked |
 | `last_downtime` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | last_downtime is the timestamp when this downtime duration was last observed |
 
 

 

 
 <a name="akash.downtimedetector.v1beta1.GenesisState"></a>

 ### GenesisState
 GenesisState defines the downtime detector module's genesis state

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `downtimes` | [GenesisDowntimeEntry](#akash.downtimedetector.v1beta1.GenesisDowntimeEntry) | repeated | downtimes is the list of tracked downtime entries |
 | `last_block_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | last_block_time is the timestamp of the last processed block |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/downtimedetector/v1beta1/query.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/downtimedetector/v1beta1/query.proto
 

 
 <a name="akash.downtimedetector.v1beta1.RecoveredSinceDowntimeOfLengthRequest"></a>

 ### RecoveredSinceDowntimeOfLengthRequest
 RecoveredSinceDowntimeOfLengthRequest is the request type for querying if the chain
has been operational for at least the specified recovery duration since experiencing
downtime of the specified length

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `downtime` | [Downtime](#akash.downtimedetector.v1beta1.Downtime) |  | downtime is the downtime duration to check against |
 | `recovery` | [google.protobuf.Duration](#google.protobuf.Duration) |  | recovery is the minimum recovery duration required since the downtime |
 
 

 

 
 <a name="akash.downtimedetector.v1beta1.RecoveredSinceDowntimeOfLengthResponse"></a>

 ### RecoveredSinceDowntimeOfLengthResponse
 RecoveredSinceDowntimeOfLengthResponse is the response type for the recovery query

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `succesfully_recovered` | [bool](#bool) |  | succesfully_recovered indicates if the chain has been up for at least the recovery duration since the last downtime of the specified length |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.downtimedetector.v1beta1.Query"></a>

 ### Query
 Query defines the gRPC querier service for the downtime detector module

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `RecoveredSinceDowntimeOfLength` | [RecoveredSinceDowntimeOfLengthRequest](#akash.downtimedetector.v1beta1.RecoveredSinceDowntimeOfLengthRequest) | [RecoveredSinceDowntimeOfLengthResponse](#akash.downtimedetector.v1beta1.RecoveredSinceDowntimeOfLengthResponse) | RecoveredSinceDowntimeOfLength queries if the chain has recovered for a specified duration since experiencing downtime of a given length | GET|/akash/downtime-detector/v1beta1/RecoveredSinceDowntimeOfLength|
 
  <!-- end services -->

 
 
 <a name="akash/epochs/v1beta1/events.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/epochs/v1beta1/events.proto
 

 
 <a name="akash.epochs.v1beta1.EventEpochEnd"></a>

 ### EventEpochEnd
 EventEpochEnd is an event emitted when an epoch end.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `epoch_number` | [int64](#int64) |  |  |
 
 

 

 
 <a name="akash.epochs.v1beta1.EventEpochStart"></a>

 ### EventEpochStart
 EventEpochStart is an event emitted when an epoch start.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `epoch_number` | [int64](#int64) |  |  |
 | `epoch_start_time` | [int64](#int64) |  |  |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/epochs/v1beta1/genesis.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/epochs/v1beta1/genesis.proto
 

 
 <a name="akash.epochs.v1beta1.EpochInfo"></a>

 ### EpochInfo
 EpochInfo is a struct that describes the data going into
a timer defined by the x/epochs module.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [string](#string) |  | id is a unique reference to this particular timer. |
 | `start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | start_time is the time at which the timer first ever ticks. If start_time is in the future, the epoch will not begin until the start time. |
 | `duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  | duration is the time in between epoch ticks. In order for intended behavior to be met, duration should be greater than the chains expected block time. Duration must be non-zero. |
 | `current_epoch` | [int64](#int64) |  | current_epoch is the current epoch number, or in other words, how many times has the timer 'ticked'. The first tick (current_epoch=1) is defined as the first block whose blocktime is greater than the EpochInfo start_time. |
 | `current_epoch_start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | current_epoch_start_time describes the start time of the current timer interval. The interval is (current_epoch_start_time, current_epoch_start_time + duration] When the timer ticks, this is set to current_epoch_start_time = last_epoch_start_time + duration only one timer tick for a given identifier can occur per block.

NOTE! The current_epoch_start_time may diverge significantly from the wall-clock time the epoch began at. Wall-clock time of epoch start may be >> current_epoch_start_time. Suppose current_epoch_start_time = 10, duration = 5. Suppose the chain goes offline at t=14, and comes back online at t=30, and produces blocks at every successive time. (t=31, 32, etc.) * The t=30 block will start the epoch for (10, 15] * The t=31 block will start the epoch for (15, 20] * The t=32 block will start the epoch for (20, 25] * The t=33 block will start the epoch for (25, 30] * The t=34 block will start the epoch for (30, 35] * The **t=36** block will start the epoch for (35, 40] |
 | `epoch_counting_started` | [bool](#bool) |  | epoch_counting_started is a boolean, that indicates whether this epoch timer has began yet. |
 | `current_epoch_start_height` | [int64](#int64) |  | current_epoch_start_height is the block height at which the current epoch started. (The block height at which the timer last ticked) |
 
 

 

 
 <a name="akash.epochs.v1beta1.GenesisState"></a>

 ### GenesisState
 GenesisState defines the epochs module's genesis state.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `epochs` | [EpochInfo](#akash.epochs.v1beta1.EpochInfo) | repeated |  |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/epochs/v1beta1/query.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/epochs/v1beta1/query.proto
 

 
 <a name="akash.epochs.v1beta1.QueryCurrentEpochRequest"></a>

 ### QueryCurrentEpochRequest
 QueryCurrentEpochRequest defines the gRPC request structure for
querying an epoch by its identifier.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `identifier` | [string](#string) |  |  |
 
 

 

 
 <a name="akash.epochs.v1beta1.QueryCurrentEpochResponse"></a>

 ### QueryCurrentEpochResponse
 QueryCurrentEpochResponse defines the gRPC response structure for
querying an epoch by its identifier.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `current_epoch` | [int64](#int64) |  |  |
 
 

 

 
 <a name="akash.epochs.v1beta1.QueryEpochInfosRequest"></a>

 ### QueryEpochInfosRequest
 QueryEpochInfosRequest defines the gRPC request structure for
querying all epoch info.

 

 

 
 <a name="akash.epochs.v1beta1.QueryEpochInfosResponse"></a>

 ### QueryEpochInfosResponse
 QueryEpochInfosRequest defines the gRPC response structure for
querying all epoch info.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `epochs` | [EpochInfo](#akash.epochs.v1beta1.EpochInfo) | repeated |  |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.epochs.v1beta1.Query"></a>

 ### Query
 Query defines the gRPC querier service.

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `EpochInfos` | [QueryEpochInfosRequest](#akash.epochs.v1beta1.QueryEpochInfosRequest) | [QueryEpochInfosResponse](#akash.epochs.v1beta1.QueryEpochInfosResponse) | EpochInfos provide running epochInfos | GET|/cosmos/epochs/v1beta1/epochs|
 | `CurrentEpoch` | [QueryCurrentEpochRequest](#akash.epochs.v1beta1.QueryCurrentEpochRequest) | [QueryCurrentEpochResponse](#akash.epochs.v1beta1.QueryCurrentEpochResponse) | CurrentEpoch provide current epoch of specified identifier | GET|/cosmos/epochs/v1beta1/current_epoch|
 
  <!-- end services -->

 
 
 <a name="akash/escrow/types/v1/payment.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/escrow/types/v1/payment.proto
 

 
 <a name="akash.escrow.types.v1.Payment"></a>

 ### Payment
 Payment

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [akash.escrow.id.v1.Payment](#akash.escrow.id.v1.Payment) |  |  |
 | `state` | [PaymentState](#akash.escrow.types.v1.PaymentState) |  |  |
 
 

 

 
 <a name="akash.escrow.types.v1.PaymentState"></a>

 ### PaymentState
 Payment stores state for a payment.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the account bech32 address of the user who owns the deployment. It is a string representing a valid bech32 account address.

Example: "akash1..." |
 | `state` | [State](#akash.escrow.types.v1.State) |  | State represents the state of the Payment. |
 | `rate` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | Rate holds the rate of the Payment. |
 | `balance` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | Balance is the current available coins. |
 | `unsettled` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | Unsettled is the amount needed to settle payment if account is overdrawn |
 | `withdrawn` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Withdrawn corresponds to the amount of coins withdrawn by the Payment. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/escrow/v1/authz.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/escrow/v1/authz.proto
 

 
 <a name="akash.escrow.v1.DepositAuthorization"></a>

 ### DepositAuthorization
 DepositAuthorization allows the grantee to deposit up to spend_limit coins from
the granter's account for Akash deployments and bids. This authorization is used
within the Cosmos SDK authz module to grant scoped permissions for deposit operations.
The authorization can be restricted to specific scopes (deployment or bid) to limit
what types of deposits the grantee is authorized to make on behalf of the granter.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `spend_limit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | SpendLimit is the maximum amount the grantee is authorized to spend from the granter's account. This limit applies cumulatively across all deposit operations within the authorized scopes. Once this limit is reached, the authorization becomes invalid and no further deposits can be made. Deprecated: use spend_limits instead |
 | `scopes` | [DepositAuthorization.Scope](#akash.escrow.v1.DepositAuthorization.Scope) | repeated | Scopes defines the specific types of deposit operations this authorization permits. This provides fine-grained control over what operations the grantee can perform using the granter's funds. |
 | `spend_limits` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | SpendLimits specifies the maximum amount per denomination the grantee is authorized to spend. Each entry represents the limit for a specific denomination, enforced independently. Once an individual denomination's limit is exhausted, no further deposits can be made in that denomination. |
 
 

 

  <!-- end messages -->

 
 <a name="akash.escrow.v1.DepositAuthorization.Scope"></a>

 ### DepositAuthorization.Scope
 Scope defines the types of deposit operations that can be authorized.
This enum is used to restrict the authorization to specific deposit contexts,
allowing fine-grained permission control within the authz system.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | invalid | 0 | Prefix should start with 0 in enum. So declaring dummy state. |
 | deployment | 1 | DepositScopeDeployment allows deposits for deployment-related operations. |
 | bid | 2 | DepositScopeBid allows deposits for bid-related operations. |
 

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/escrow/v1/genesis.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/escrow/v1/genesis.proto
 

 
 <a name="akash.escrow.v1.GenesisState"></a>

 ### GenesisState
 GenesisState defines the basic genesis state used by the escrow module.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `accounts` | [akash.escrow.types.v1.Account](#akash.escrow.types.v1.Account) | repeated | Accounts is a list of accounts on the genesis state. |
 | `payments` | [akash.escrow.types.v1.Payment](#akash.escrow.types.v1.Payment) | repeated | Payments is a list of fractional payments on the genesis state.. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/escrow/v1/msg.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/escrow/v1/msg.proto
 

 
 <a name="akash.escrow.v1.MsgAccountDeposit"></a>

 ### MsgAccountDeposit
 MsgAccountDeposit represents a message to deposit funds into an existing escrow account
on the blockchain. This is part of the interaction mechanism for managing
deployment-related resources.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `signer` | [string](#string) |  | Signer is the account bech32 address of the user who wants to deposit into an escrow account. Does not necessarily needs to be an owner of the deployment. It is a string representing a valid bech32 account address.

Example: "akash1..." |
 | `id` | [akash.escrow.id.v1.Account](#akash.escrow.id.v1.Account) |  | ID is the unique identifier of the account. |
 | `deposit` | [akash.base.deposit.v1.Deposit](#akash.base.deposit.v1.Deposit) |  | Deposit contains information about the deposit amount and the source of the deposit to the escrow account. |
 
 

 

 
 <a name="akash.escrow.v1.MsgAccountDepositResponse"></a>

 ### MsgAccountDepositResponse
 MsgAccountDepositResponse defines response type for the MsgDeposit.

 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/escrow/v1/query.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/escrow/v1/query.proto
 

 
 <a name="akash.escrow.v1.QueryAccountsRequest"></a>

 ### QueryAccountsRequest
 QueryAccountRequest is request type for the Query/Account RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `state` | [string](#string) |  | State represents the current state of an Account. |
 | `xid` | [string](#string) |  | Scope holds the scope of the account. |
 | `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | Pagination is used to paginate the request. |
 
 

 

 
 <a name="akash.escrow.v1.QueryAccountsResponse"></a>

 ### QueryAccountsResponse
 QueryProvidersResponse is response type for the Query/Providers RPC method

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `accounts` | [akash.escrow.types.v1.Account](#akash.escrow.types.v1.Account) | repeated | Accounts is a list of Account. |
 | `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | Pagination contains the information about response pagination. |
 
 

 

 
 <a name="akash.escrow.v1.QueryPaymentsRequest"></a>

 ### QueryPaymentsRequest
 QueryPaymentRequest is request type for the Query/Payment RPC method

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `state` | [string](#string) |  | State represents the current state of a Payment. |
 | `xid` | [string](#string) |  |  |
 | `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | Pagination is used to paginate the request. |
 
 

 

 
 <a name="akash.escrow.v1.QueryPaymentsResponse"></a>

 ### QueryPaymentsResponse
 QueryProvidersResponse is response type for the Query/Providers RPC method

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `payments` | [akash.escrow.types.v1.Payment](#akash.escrow.types.v1.Payment) | repeated | Payments is a list of payments. |
 | `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | Pagination contains the information about response pagination. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.escrow.v1.Query"></a>

 ### Query
 Query defines the gRPC querier service for the escrow package.

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `Accounts` | [QueryAccountsRequest](#akash.escrow.v1.QueryAccountsRequest) | [QueryAccountsResponse](#akash.escrow.v1.QueryAccountsResponse) | buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE buf:lint:ignore RPC_RESPONSE_STANDARD_NAME Accounts queries all accounts. | GET|/akash/escrow/v1/types/accounts|
 | `Payments` | [QueryPaymentsRequest](#akash.escrow.v1.QueryPaymentsRequest) | [QueryPaymentsResponse](#akash.escrow.v1.QueryPaymentsResponse) | buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE buf:lint:ignore RPC_RESPONSE_STANDARD_NAME Payments queries all payments. | GET|/akash/escrow/v1/types/payments|
 
  <!-- end services -->

 
 
 <a name="akash/escrow/v1/service.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/escrow/v1/service.proto
 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.escrow.v1.Msg"></a>

 ### Msg
 Msg defines the x/deployment Msg service.

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `AccountDeposit` | [MsgAccountDeposit](#akash.escrow.v1.MsgAccountDeposit) | [MsgAccountDepositResponse](#akash.escrow.v1.MsgAccountDepositResponse) | AccountDeposit deposits more funds into the escrow account. | |
 
  <!-- end services -->

 
 
 <a name="akash/market/v1/bid.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/market/v1/bid.proto
 

 
 <a name="akash.market.v1.BidID"></a>

 ### BidID
 BidID stores owner and all other seq numbers.
A successful bid becomes a Lease(ID).

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the account bech32 address of the user who owns the deployment. It is a string representing a valid bech32 account address.

Example: "akash1..." |
 | `dseq` | [uint64](#uint64) |  | Dseq (deployment sequence number) is a unique numeric identifier for the deployment. It is used to differentiate deployments created by the same owner. |
 | `gseq` | [uint32](#uint32) |  | Gseq (group sequence number) is a unique numeric identifier for the group. It is used to differentiate groups created by the same owner in a deployment. |
 | `oseq` | [uint32](#uint32) |  | Oseq (order sequence) distinguishes multiple orders associated with a single deployment. Oseq is incremented when a lease associated with an existing deployment is closed, and a new order is generated. |
 | `provider` | [string](#string) |  | Provider is the account bech32 address of the provider making the bid. It is a string representing a valid account bech32 address.

Example: "akash1..." |
 | `bseq` | [uint32](#uint32) |  | BSeq (bid sequence) distinguishes multiple bids associated with a single deployment from same provider. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/market/v1/order.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/market/v1/order.proto
 

 
 <a name="akash.market.v1.OrderID"></a>

 ### OrderID
 OrderId stores owner and all other seq numbers.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the account bech32 address of the user who owns the deployment. It is a string representing a valid bech32 account address.

Example: "akash1..." |
 | `dseq` | [uint64](#uint64) |  | Dseq (deployment sequence number) is a unique numeric identifier for the deployment. It is used to differentiate deployments created by the same owner. |
 | `gseq` | [uint32](#uint32) |  | Gseq (group sequence number) is a unique numeric identifier for the group. It is used to differentiate groups created by the same owner in a deployment. |
 | `oseq` | [uint32](#uint32) |  | Oseq (order sequence) distinguishes multiple orders associated with a single deployment. Oseq is incremented when a lease associated with an existing deployment is closed, and a new order is generated. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/market/v1/types.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/market/v1/types.proto
 

  <!-- end messages -->

 
 <a name="akash.market.v1.LeaseClosedReason"></a>

 ### LeaseClosedReason
 LeaseClosedReason indicates reason bid was closed

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | lease_closed_invalid | 0 | LeaseClosedReasonInvalid represents the default zero value for LeaseClosedReason. This value indicates an uninitialized or invalid lease closure reason and should not be used |
 | lease_closed_owner | 1 | values between 1..9999 indicate owner-initiated close. |
 | lease_closed_reason_unstable | 10000 | values between 10000..19999 are indicating provider initiated close. lease_closed_reason_unstable lease workloads have been unstable |
 | lease_closed_reason_decommission | 10001 | lease_closed_reason_decommission provider is being decommissioned |
 | lease_closed_reason_unspecified | 10002 | lease_closed_reason_unspecified provider did not specify reason |
 | lease_closed_reason_manifest_timeout | 10003 | lease_closed_reason_manifest_timeout provider closed leases due to manifest not received |
 | lease_closed_reason_provider | 10004 | lease_closed_reason_provider provider closed the lease |
 | lease_closed_reason_insufficient_funds | 20000 | values between 20000..29999 indicate network-initiated close. |
 

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/market/v1/reclamation.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/market/v1/reclamation.proto
 

 
 <a name="akash.market.v1.Reclamation"></a>

 ### Reclamation
 Reclamation defines the runtime reclamation state stored on a Lease.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `window` | [google.protobuf.Duration](#google.protobuf.Duration) |  | window is the negotiated reclamation window duration (from the winning bid). |
 | `started_at` | [int64](#int64) |  | started_at is the block height at which reclamation was initiated. Zero means reclamation has not been started yet. |
 | `deadline` | [int64](#int64) |  | deadline is the unix timestamp at which the reclamation window expires. Zero means reclamation has not been started yet. |
 | `reason` | [LeaseClosedReason](#akash.market.v1.LeaseClosedReason) |  | reason is the provider's stated reason for reclamation. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/market/v1/lease.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/market/v1/lease.proto
 

 
 <a name="akash.market.v1.Lease"></a>

 ### Lease
 Lease stores LeaseID, state of lease and price.
The Lease defines the terms under which the provider allocates resources to fulfill
the tenant's deployment requirements.
Leases are paid from the tenant to the provider through a deposit and withdraw mechanism and are priced in blocks.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [LeaseID](#akash.market.v1.LeaseID) |  | Id is the unique identifier of the Lease. |
 | `state` | [Lease.State](#akash.market.v1.Lease.State) |  | State represents the state of the Lease. |
 | `price` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | Price holds the settled price for the Lease. |
 | `created_at` | [int64](#int64) |  | CreatedAt is the block height at which the Lease was created. |
 | `closed_on` | [int64](#int64) |  | ClosedOn is the block height at which the Lease was closed. |
 | `reason` | [LeaseClosedReason](#akash.market.v1.LeaseClosedReason) |  |  |
 | `reclamation` | [Reclamation](#akash.market.v1.Reclamation) |  | Reclamation holds reclamation configuration and state, if applicable. Nil if reclamation is not configured for this lease. |
 
 

 

 
 <a name="akash.market.v1.LeaseID"></a>

 ### LeaseID
 LeaseID stores bid details of lease.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the account bech32 address of the user who owns the deployment. It is a string representing a valid bech32 account address.

Example: "akash1..." |
 | `dseq` | [uint64](#uint64) |  | Dseq (deployment sequence number) is a unique numeric identifier for the deployment. It is used to differentiate deployments created by the same owner. |
 | `gseq` | [uint32](#uint32) |  | Gseq (group sequence number) is a unique numeric identifier for the group. It is used to differentiate groups created by the same owner in a deployment. |
 | `oseq` | [uint32](#uint32) |  | Oseq (order sequence) distinguishes multiple orders associated with a single deployment. Oseq is incremented when a lease associated with an existing deployment is closed, and a new order is generated. |
 | `provider` | [string](#string) |  | Provider is the account bech32 address of the provider making the bid. It is a string representing a valid account bech32 address.

Example: "akash1..." |
 | `bseq` | [uint32](#uint32) |  | BSeq (bid sequence) distinguishes multiple bids associated with a single deployment from same provider. |
 
 

 

  <!-- end messages -->

 
 <a name="akash.market.v1.Lease.State"></a>

 ### Lease.State
 State is an enum which refers to state of lease.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | invalid | 0 | Prefix should start with 0 in enum. So declaring dummy state. |
 | active | 1 | LeaseActive denotes state for lease active. |
 | insufficient_funds | 2 | LeaseInsufficientFunds denotes state for lease insufficient_funds. |
 | closed | 3 | LeaseClosed denotes state for lease closed. |
 | reclaiming | 4 | LeaseReclaiming denotes a lease in reclamation (grace period before closure). |
 

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/market/v1/event.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/market/v1/event.proto
 

 
 <a name="akash.market.v1.EventBidClosed"></a>

 ### EventBidClosed
 EventBidClosed is triggered when a bid is closed.
It contains all the information required to identify a bid.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [BidID](#akash.market.v1.BidID) |  | Id is the unique identifier of the Bid. |
 
 

 

 
 <a name="akash.market.v1.EventBidCreated"></a>

 ### EventBidCreated
 EventBidCreated is triggered when a bid is created.
It contains all the information required to identify a bid.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [BidID](#akash.market.v1.BidID) |  | Id is the unique identifier of the Bid. |
 | `price` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | Price stated on the Bid. |
 
 

 

 
 <a name="akash.market.v1.EventLeaseClosed"></a>

 ### EventLeaseClosed
 EventLeaseClosed is triggered when a lease is closed.
It contains all the information required to identify a lease.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [LeaseID](#akash.market.v1.LeaseID) |  | Id is the unique identifier of the Lease. |
 | `reason` | [LeaseClosedReason](#akash.market.v1.LeaseClosedReason) |  |  |
 
 

 

 
 <a name="akash.market.v1.EventLeaseCreated"></a>

 ### EventLeaseCreated
 EventLeaseCreated is triggered when a lease is created.
It contains all the information required to identify a lease.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [LeaseID](#akash.market.v1.LeaseID) |  | Id is the unique identifier of the Lease. |
 | `price` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | Price settled for the lease. |
 
 

 

 
 <a name="akash.market.v1.EventLeaseReclaimStarted"></a>

 ### EventLeaseReclaimStarted
 EventLeaseReclaimStarted is triggered when a provider initiates reclamation on a lease.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [LeaseID](#akash.market.v1.LeaseID) |  | Id is the unique identifier of the Lease. |
 | `reason` | [LeaseClosedReason](#akash.market.v1.LeaseClosedReason) |  | reason is the provider's stated reason for reclamation. |
 | `deadline` | [int64](#int64) |  | deadline is the unix timestamp when the reclamation window expires. |
 
 

 

 
 <a name="akash.market.v1.EventOrderClosed"></a>

 ### EventOrderClosed
 EventOrderClosed is triggered when an order is closed.
It contains all the information required to identify an order.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [OrderID](#akash.market.v1.OrderID) |  | Id is the unique identifier of the Order. |
 
 

 

 
 <a name="akash.market.v1.EventOrderCreated"></a>

 ### EventOrderCreated
 EventOrderCreated is triggered when an order is created.
It contains all the information required to identify an order.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [OrderID](#akash.market.v1.OrderID) |  | Id is the unique identifier of the Order. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/market/v1/filters.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/market/v1/filters.proto
 

 
 <a name="akash.market.v1.LeaseFilters"></a>

 ### LeaseFilters
 LeaseFilters defines flags for lease list filtering.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the account bech32 address of the user who owns the deployment. It is a string representing a valid bech32 account address.

Example: "akash1..." |
 | `dseq` | [uint64](#uint64) |  | Dseq (deployment sequence number) is a unique numeric identifier for the deployment. It is used to differentiate deployments created by the same owner. |
 | `gseq` | [uint32](#uint32) |  | Gseq (group sequence number) is a unique numeric identifier for the group. It is used to differentiate groups created by the same owner in a deployment. |
 | `oseq` | [uint32](#uint32) |  | Oseq (order sequence) distinguishes multiple orders associated with a single deployment. Oseq is incremented when a lease associated with an existing deployment is closed, and a new order is generated. |
 | `provider` | [string](#string) |  | Provider is the account bech32 address of the provider making the bid. It is a string representing a valid account bech32 address.

Example: "akash1..." |
 | `state` | [string](#string) |  | State represents the state of the lease. |
 | `bseq` | [uint32](#uint32) |  | BSeq (bid sequence) distinguishes multiple bids associated with a single deployment from same provider. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/market/v1/stats.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/market/v1/stats.proto
 

 
 <a name="akash.market.v1.ProviderLeaseStats"></a>

 ### ProviderLeaseStats
 ProviderLeaseStats stores aggregate lease-completion stats for a provider.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `total_leases` | [uint64](#uint64) |  | TotalLeases is the total number of leases included in the aggregate. |
 | `completed_leases` | [uint64](#uint64) |  | CompletedLeases is the number of leases not closed for provider-fault reasons. |
 | `provider_faulted_leases` | [uint64](#uint64) |  | ProviderFaultedLeases is the number of provider-fault lease closures. |
 | `provider_faults` | [ProviderLeaseStatsByReason](#akash.market.v1.ProviderLeaseStatsByReason) | repeated | ProviderFaults is the provider-fault lease count split by close reason. |
 
 

 

 
 <a name="akash.market.v1.ProviderLeaseStatsByReason"></a>

 ### ProviderLeaseStatsByReason
 ProviderLeaseStatsByReason stores a provider-fault lease count for one close
reason.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `reason` | [LeaseClosedReason](#akash.market.v1.LeaseClosedReason) |  | Reason is the provider-initiated close reason being counted. |
 | `count` | [uint64](#uint64) |  | Count is the number of leases closed for this reason. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/market/v1beta5/resourcesoffer.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/market/v1beta5/resourcesoffer.proto
 

 
 <a name="akash.market.v1beta5.EndpointOfferPrice"></a>

 ### EndpointOfferPrice
 EndpointOfferPrice represents the price a provider is offering for a specific
kind of network endpoint. Providers may price each endpoint kind differently
(e.g., a leased IP may cost more than a shared HTTP ingress). This type is
used as a repeated field within OfferPrices to express per-kind endpoint
pricing in a bid.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `kind` | [akash.base.resources.v1beta4.Endpoint.Kind](#akash.base.resources.v1beta4.Endpoint.Kind) |  | Kind specifies the type of network endpoint being priced. Possible values: - SHARED_HTTP (0): A Kubernetes Ingress endpoint. - RANDOM_PORT (1): A Kubernetes NodePort endpoint. - LEASED_IP (2): A dedicated leased IP endpoint. |
 | `price` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | Price is the offered price per unit of this endpoint kind, expressed as a DecCoin (decimal coin) to allow fractional pricing. When nil, no explicit price is set for this endpoint kind. |
 
 

 

 
 <a name="akash.market.v1beta5.OfferPrices"></a>

 ### OfferPrices
 OfferPrices contains the complete pricing breakdown that a provider includes
in a bid for a deployment resource group. Each field represents the price for
a specific compute resource type. All price fields use DecCoin (decimal coin)
to support fractional pricing denominated in any supported token.

This message is embedded as a nullable field on ResourceOffer, which in turn
is carried by Bid and MsgCreateBid messages. A nil OfferPrices on a
ResourceOffer indicates that no per-resource pricing was specified.

Field 1 is reserved for backward compatibility with a previously removed field.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `cpu` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | Cpu is the offered price for CPU resources. When nil, no explicit CPU price is set. |
 | `memory` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | Memory is the offered price for memory resources. When nil, no explicit memory price is set. |
 | `storage` | [StorageOfferPrice](#akash.market.v1beta5.StorageOfferPrice) | repeated | Storage is a list of per-class storage prices. Each entry corresponds to a named storage class (e.g., "default", "ssd") and its associated price. Multiple entries allow providers to price different storage tiers independently. |
 | `gpu` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | Gpu is the offered price for GPU resources. When nil, no explicit GPU price is set. |
 | `endpoints` | [EndpointOfferPrice](#akash.market.v1beta5.EndpointOfferPrice) | repeated | Endpoints is a list of per-kind endpoint prices. Each entry corresponds to a network endpoint kind (SHARED_HTTP, RANDOM_PORT, or LEASED_IP) and its associated price. Multiple entries allow providers to price different endpoint types independently. |
 
 

 

 
 <a name="akash.market.v1beta5.ResourceOffer"></a>

 ### ResourceOffer
 ResourceOffer describes resources that provider is offering
for deployment.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `resources` | [akash.base.resources.v1beta4.Resources](#akash.base.resources.v1beta4.Resources) |  | Resources holds information about bid resources. |
 | `count` | [uint32](#uint32) |  | Count is the number of resources. |
 | `prices` | [OfferPrices](#akash.market.v1beta5.OfferPrices) |  | Prices contains per-resource pricing details (CPU, memory, storage, GPU, endpoints) for this offer. |
 
 

 

 
 <a name="akash.market.v1beta5.StorageOfferPrice"></a>

 ### StorageOfferPrice
 StorageOfferPrice represents the price a provider is offering for a specific
class of persistent storage. Providers may offer multiple storage classes
(e.g., SSD, HDD, NVMe), each identified by name and priced independently.
This type is used as a repeated field within OfferPrices to express
per-class storage pricing in a bid.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `name` | [string](#string) |  | Name holds an arbitrary name for the storage class (e.g., "default", "ssd", "hdd"). This must match a storage class name from the corresponding resource specification. |
 | `price` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | Price is the offered price per unit of this storage class, expressed as a DecCoin (decimal coin) to allow fractional pricing. When nil, no explicit price is set for this storage class. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/market/v1beta5/bid.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/market/v1beta5/bid.proto
 

 
 <a name="akash.market.v1beta5.Bid"></a>

 ### Bid
 Bid stores BidID, state of bid and price.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [akash.market.v1.BidID](#akash.market.v1.BidID) |  | BidID stores owner and all other seq numbers. A successful bid becomes a Lease(ID). |
 | `state` | [Bid.State](#akash.market.v1beta5.Bid.State) |  | State represents the state of the Bid. |
 | `price` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | Price holds the pricing stated on the Bid. |
 | `created_at` | [int64](#int64) |  | CreatedAt is the block height at which the Bid was created. |
 | `resources_offer` | [ResourceOffer](#akash.market.v1beta5.ResourceOffer) | repeated | ResourceOffer is a list of offers. |
 | `reclamation_window` | [google.protobuf.Duration](#google.protobuf.Duration) |  | reclamation_window is the reclamation window offered by this provider. |
 
 

 

  <!-- end messages -->

 
 <a name="akash.market.v1beta5.Bid.State"></a>

 ### Bid.State
 BidState is an enum which refers to state of bid.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | invalid | 0 | Prefix should start with 0 in enum. So declaring dummy state. |
 | open | 1 | BidOpen denotes state for bid open. |
 | active | 2 | BidMatched denotes state for bid open. |
 | lost | 3 | BidLost denotes state for bid lost. |
 | closed | 4 | BidClosed denotes state for bid closed. |
 

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/market/v1beta5/bidmsg.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/market/v1beta5/bidmsg.proto
 

 
 <a name="akash.market.v1beta5.MsgCloseBid"></a>

 ### MsgCloseBid
 MsgCloseBid defines an SDK message for closing bid.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [akash.market.v1.BidID](#akash.market.v1.BidID) |  | Id is the unique identifier of the Bid. |
 | `reason` | [akash.market.v1.LeaseClosedReason](#akash.market.v1.LeaseClosedReason) |  |  |
 
 

 

 
 <a name="akash.market.v1beta5.MsgCloseBidResponse"></a>

 ### MsgCloseBidResponse
 MsgCloseBidResponse defines the Msg/CloseBid response type.

 

 

 
 <a name="akash.market.v1beta5.MsgCreateBid"></a>

 ### MsgCreateBid
 MsgCreateBid defines an SDK message for creating Bid.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [akash.market.v1.BidID](#akash.market.v1.BidID) |  |  |
 | `price` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | Price holds the pricing stated on the Bid. |
 | `deposit` | [akash.base.deposit.v1.Deposit](#akash.base.deposit.v1.Deposit) |  | Deposit holds the amount of coins to deposit. |
 | `resources_offer` | [ResourceOffer](#akash.market.v1beta5.ResourceOffer) | repeated | ResourceOffer is a list of resource offers. |
 | `reclamation_window` | [google.protobuf.Duration](#google.protobuf.Duration) |  | reclamation_window is the reclamation window duration the provider offers. If the order requires reclamation, this must be >= the order's min_window. Nil means the provider does not offer reclamation on this bid. |
 
 

 

 
 <a name="akash.market.v1beta5.MsgCreateBidResponse"></a>

 ### MsgCreateBidResponse
 MsgCreateBidResponse defines the Msg/CreateBid response type.

 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/market/v1beta5/filters.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/market/v1beta5/filters.proto
 

 
 <a name="akash.market.v1beta5.BidFilters"></a>

 ### BidFilters
 BidFilters defines flags for bid list filter.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the account bech32 address of the user who owns the deployment. It is a string representing a valid bech32 account address.

Example: "akash1..." |
 | `dseq` | [uint64](#uint64) |  | Dseq (deployment sequence number) is a unique numeric identifier for the deployment. It is used to differentiate deployments created by the same owner. |
 | `gseq` | [uint32](#uint32) |  | Gseq (group sequence number) is a unique numeric identifier for the group. It is used to differentiate groups created by the same owner in a deployment. |
 | `oseq` | [uint32](#uint32) |  | Oseq (order sequence) distinguishes multiple orders associated with a single deployment. Oseq is incremented when a lease associated with an existing deployment is closed, and a new order is generated. |
 | `provider` | [string](#string) |  | Provider is the account bech32 address of the provider making the bid. It is a string representing a valid account bech32 address.

Example: "akash1..." |
 | `state` | [string](#string) |  | State represents the state of the lease. |
 | `bseq` | [uint32](#uint32) |  | BSeq (bid sequence) distinguishes multiple bids associated with a single deployment from same provider. |
 
 

 

 
 <a name="akash.market.v1beta5.OrderFilters"></a>

 ### OrderFilters
 OrderFilters defines flags for order list filter

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the account bech32 address of the user who owns the deployment. It is a string representing a valid bech32 account address.

Example: "akash1..." |
 | `dseq` | [uint64](#uint64) |  | Dseq (deployment sequence number) is a unique numeric identifier for the deployment. It is used to differentiate deployments created by the same owner. |
 | `gseq` | [uint32](#uint32) |  | Gseq (group sequence number) is a unique numeric identifier for the group. It is used to differentiate groups created by the same owner in a deployment. |
 | `oseq` | [uint32](#uint32) |  | Oseq (order sequence) distinguishes multiple orders associated with a single deployment. Oseq is incremented when a lease associated with an existing deployment is closed, and a new order is generated. |
 | `state` | [string](#string) |  | State represents the state of the lease. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/market/v1beta5/params.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/market/v1beta5/params.proto
 

 
 <a name="akash.market.v1beta5.Params"></a>

 ### Params
 Params is the params for the x/market module.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `bid_min_deposit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | BidMinDeposit is a parameter for the minimum deposit on a Bid. Deprecated: use BidMinDeposits |
 | `order_max_bids` | [uint32](#uint32) |  | OrderMaxBids is a parameter for the maximum number of bids in an order. |
 | `bid_min_deposits` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | BidMinDeposits is a parameter for the minimum deposits per denom on a Bid. |
 | `min_reclamation_window` | [google.protobuf.Duration](#google.protobuf.Duration) |  | min_reclamation_window is the minimum reclamation window duration allowed. |
 | `max_reclamation_window` | [google.protobuf.Duration](#google.protobuf.Duration) |  | max_reclamation_window is the maximum reclamation window duration allowed. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/market/v1beta5/order.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/market/v1beta5/order.proto
 

 
 <a name="akash.market.v1beta5.Order"></a>

 ### Order
 Order stores orderID, state of order and other details.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [akash.market.v1.OrderID](#akash.market.v1.OrderID) |  | Id is the unique identifier of the order. |
 | `state` | [Order.State](#akash.market.v1beta5.Order.State) |  |  |
 | `spec` | [akash.deployment.v1beta4.GroupSpec](#akash.deployment.v1beta4.GroupSpec) |  |  |
 | `created_at` | [int64](#int64) |  |  |
 | `reclamation` | [akash.deployment.v1.DeploymentReclamation](#akash.deployment.v1.DeploymentReclamation) |  | reclamation is the deployment-level reclamation requirement, propagated to the order. Nil means the deployment does not require reclamation. |
 
 

 

  <!-- end messages -->

 
 <a name="akash.market.v1beta5.Order.State"></a>

 ### Order.State
 State is an enum which refers to state of order.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | invalid | 0 | Prefix should start with 0 in enum. So declaring dummy state. |
 | open | 1 | OrderOpen denotes state for order open. |
 | active | 2 | OrderMatched denotes state for order matched. |
 | closed | 3 | OrderClosed denotes state for order lost. |
 

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/market/v1beta5/genesis.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/market/v1beta5/genesis.proto
 

 
 <a name="akash.market.v1beta5.GenesisState"></a>

 ### GenesisState
 GenesisState defines the basic genesis state used by market module.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `params` | [Params](#akash.market.v1beta5.Params) |  | Params holds parameters of the genesis of market. |
 | `orders` | [Order](#akash.market.v1beta5.Order) | repeated | Orders is a list of orders in the genesis state. |
 | `leases` | [akash.market.v1.Lease](#akash.market.v1.Lease) | repeated | Leases is a list of leases in the genesis state. |
 | `bids` | [Bid](#akash.market.v1beta5.Bid) | repeated | Bids is a list of bids in the genesis state. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/market/v1beta5/leasemsg.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/market/v1beta5/leasemsg.proto
 

 
 <a name="akash.market.v1beta5.MsgCloseLease"></a>

 ### MsgCloseLease
 MsgCloseLease defines an SDK message for closing order.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [akash.market.v1.LeaseID](#akash.market.v1.LeaseID) |  | LeaseID is the unique identifier of the Lease. |
 | `reason` | [akash.market.v1.LeaseClosedReason](#akash.market.v1.LeaseClosedReason) |  |  |
 
 

 

 
 <a name="akash.market.v1beta5.MsgCloseLeaseResponse"></a>

 ### MsgCloseLeaseResponse
 MsgCloseLeaseResponse defines the Msg/CloseLease response type.

 

 

 
 <a name="akash.market.v1beta5.MsgCreateLease"></a>

 ### MsgCreateLease
 MsgCreateLease is sent to create a lease.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `bid_id` | [akash.market.v1.BidID](#akash.market.v1.BidID) |  | BidId is the unique identifier of the Bid. |
 
 

 

 
 <a name="akash.market.v1beta5.MsgCreateLeaseResponse"></a>

 ### MsgCreateLeaseResponse
 MsgCreateLeaseResponse is the response from creating a lease.

 

 

 
 <a name="akash.market.v1beta5.MsgLeaseStartReclaim"></a>

 ### MsgLeaseStartReclaim
 MsgLeaseStartReclaim is sent by the provider to initiate reclamation on an active lease.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [akash.market.v1.LeaseID](#akash.market.v1.LeaseID) |  | Id is the unique identifier of the Lease. |
 | `reason` | [akash.market.v1.LeaseClosedReason](#akash.market.v1.LeaseClosedReason) |  | reason is the provider's stated reason for initiating reclamation. |
 
 

 

 
 <a name="akash.market.v1beta5.MsgLeaseStartReclaimResponse"></a>

 ### MsgLeaseStartReclaimResponse
 MsgLeaseStartReclaimResponse is the response from starting lease reclamation.

 

 

 
 <a name="akash.market.v1beta5.MsgWithdrawLease"></a>

 ### MsgWithdrawLease
 MsgWithdrawLease defines an SDK message for withdrawing lease funds.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [akash.market.v1.LeaseID](#akash.market.v1.LeaseID) |  | BidId is the unique identifier of the Bid. |
 
 

 

 
 <a name="akash.market.v1beta5.MsgWithdrawLeaseResponse"></a>

 ### MsgWithdrawLeaseResponse
 MsgWithdrawLeaseResponse defines the Msg/WithdrawLease response type.

 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/market/v1beta5/paramsmsg.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/market/v1beta5/paramsmsg.proto
 

 
 <a name="akash.market.v1beta5.MsgUpdateParams"></a>

 ### MsgUpdateParams
 MsgUpdateParams is the Msg/UpdateParams request type.

Since: akash v1.0.0

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `authority` | [string](#string) |  | authority is the address of the governance account. |
 | `params` | [Params](#akash.market.v1beta5.Params) |  | params defines the x/deployment parameters to update.

NOTE: All parameters must be supplied. |
 
 

 

 
 <a name="akash.market.v1beta5.MsgUpdateParamsResponse"></a>

 ### MsgUpdateParamsResponse
 MsgUpdateParamsResponse defines the response structure for executing a
MsgUpdateParams message.

Since: akash v1.0.0

 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/market/v1beta5/query.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/market/v1beta5/query.proto
 

 
 <a name="akash.market.v1beta5.QueryBidRequest"></a>

 ### QueryBidRequest
 QueryBidRequest is request type for the Query/Bid RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [akash.market.v1.BidID](#akash.market.v1.BidID) |  | Id is the unique identifier for the Bid. |
 
 

 

 
 <a name="akash.market.v1beta5.QueryBidResponse"></a>

 ### QueryBidResponse
 QueryBidResponse is response type for the Query/Bid RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `bid` | [Bid](#akash.market.v1beta5.Bid) |  | Bid represents a deployment bid. |
 | `escrow_account` | [akash.escrow.types.v1.Account](#akash.escrow.types.v1.Account) |  | EscrowAccount represents the escrow account created for the Bid. |
 
 

 

 
 <a name="akash.market.v1beta5.QueryBidsRequest"></a>

 ### QueryBidsRequest
 QueryBidsRequest is request type for the Query/Bids RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `filters` | [BidFilters](#akash.market.v1beta5.BidFilters) |  | Filters holds the fields to filter bids. |
 | `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | Pagination is used to paginate the request. |
 
 

 

 
 <a name="akash.market.v1beta5.QueryBidsResponse"></a>

 ### QueryBidsResponse
 QueryBidsResponse is response type for the Query/Bids RPC method

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `bids` | [QueryBidResponse](#akash.market.v1beta5.QueryBidResponse) | repeated | Bids is a list of deployment bids. |
 | `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | Pagination contains the information about response pagination. |
 
 

 

 
 <a name="akash.market.v1beta5.QueryLeaseRequest"></a>

 ### QueryLeaseRequest
 QueryLeaseRequest is request type for the Query/Lease RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [akash.market.v1.LeaseID](#akash.market.v1.LeaseID) |  | Id is the unique identifier of the Lease. |
 
 

 

 
 <a name="akash.market.v1beta5.QueryLeaseResponse"></a>

 ### QueryLeaseResponse
 QueryLeaseResponse is response type for the Query/Lease RPC method

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `lease` | [akash.market.v1.Lease](#akash.market.v1.Lease) |  | Lease holds the lease for a deployment. |
 | `escrow_payment` | [akash.escrow.types.v1.Payment](#akash.escrow.types.v1.Payment) |  | EscrowPayment holds information about the Lease's fractional payment. |
 
 

 

 
 <a name="akash.market.v1beta5.QueryLeasesRequest"></a>

 ### QueryLeasesRequest
 QueryLeasesRequest is request type for the Query/Leases RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `filters` | [akash.market.v1.LeaseFilters](#akash.market.v1.LeaseFilters) |  | Filters holds the fields to filter leases. |
 | `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | Pagination is used to paginate the request. |
 
 

 

 
 <a name="akash.market.v1beta5.QueryLeasesResponse"></a>

 ### QueryLeasesResponse
 QueryLeasesResponse is response type for the Query/Leases RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `leases` | [QueryLeaseResponse](#akash.market.v1beta5.QueryLeaseResponse) | repeated | Leases is a list of Lease. |
 | `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | Pagination contains the information about response pagination. |
 
 

 

 
 <a name="akash.market.v1beta5.QueryOrderRequest"></a>

 ### QueryOrderRequest
 QueryOrderRequest is request type for the Query/Order RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [akash.market.v1.OrderID](#akash.market.v1.OrderID) |  | Id is the unique identifier of the Order. |
 
 

 

 
 <a name="akash.market.v1beta5.QueryOrderResponse"></a>

 ### QueryOrderResponse
 QueryOrderResponse is response type for the Query/Order RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `order` | [Order](#akash.market.v1beta5.Order) |  | Order represents a market order. |
 
 

 

 
 <a name="akash.market.v1beta5.QueryOrdersRequest"></a>

 ### QueryOrdersRequest
 QueryOrdersRequest is request type for the Query/Orders RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `filters` | [OrderFilters](#akash.market.v1beta5.OrderFilters) |  | Filters holds the fields to filter orders. |
 | `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | Pagination is used to paginate the request. |
 
 

 

 
 <a name="akash.market.v1beta5.QueryOrdersResponse"></a>

 ### QueryOrdersResponse
 QueryOrdersResponse is response type for the Query/Orders RPC method

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `orders` | [Order](#akash.market.v1beta5.Order) | repeated | Orders is a list of market orders. |
 | `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | Pagination contains the information about response pagination. |
 
 

 

 
 <a name="akash.market.v1beta5.QueryParamsRequest"></a>

 ### QueryParamsRequest
 QueryParamsRequest is the request type for the Query/Params RPC method.

 

 

 
 <a name="akash.market.v1beta5.QueryParamsResponse"></a>

 ### QueryParamsResponse
 QueryParamsResponse is the response type for the Query/Params RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `params` | [Params](#akash.market.v1beta5.Params) |  | params defines the parameters of the module. |
 
 

 

 
 <a name="akash.market.v1beta5.QueryProviderLeaseStatsRequest"></a>

 ### QueryProviderLeaseStatsRequest
 QueryProviderLeaseStatsRequest is the request type for the
Query/ProviderLeaseStats RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the account bech32 address of the provider being queried.

Example: "akash1..." |
 
 

 

 
 <a name="akash.market.v1beta5.QueryProviderLeaseStatsResponse"></a>

 ### QueryProviderLeaseStatsResponse
 QueryProviderLeaseStatsResponse is the response type for the
Query/ProviderLeaseStats RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `stats` | [akash.market.v1.ProviderLeaseStats](#akash.market.v1.ProviderLeaseStats) |  | Stats holds the aggregate lease-completion stats for the provider. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.market.v1beta5.Query"></a>

 ### Query
 Query defines the gRPC querier service for the market package.

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `Orders` | [QueryOrdersRequest](#akash.market.v1beta5.QueryOrdersRequest) | [QueryOrdersResponse](#akash.market.v1beta5.QueryOrdersResponse) | Orders queries orders with filters. | GET|/akash/market/v1beta5/orders/list|
 | `Order` | [QueryOrderRequest](#akash.market.v1beta5.QueryOrderRequest) | [QueryOrderResponse](#akash.market.v1beta5.QueryOrderResponse) | Order queries order details. | GET|/akash/market/v1beta5/orders/info|
 | `Bids` | [QueryBidsRequest](#akash.market.v1beta5.QueryBidsRequest) | [QueryBidsResponse](#akash.market.v1beta5.QueryBidsResponse) | Bids queries bids with filters. | GET|/akash/market/v1beta5/bids/list|
 | `Bid` | [QueryBidRequest](#akash.market.v1beta5.QueryBidRequest) | [QueryBidResponse](#akash.market.v1beta5.QueryBidResponse) | Bid queries bid details. | GET|/akash/market/v1beta5/bids/info|
 | `Leases` | [QueryLeasesRequest](#akash.market.v1beta5.QueryLeasesRequest) | [QueryLeasesResponse](#akash.market.v1beta5.QueryLeasesResponse) | Leases queries leases with filters. | GET|/akash/market/v1beta5/leases/list|
 | `Lease` | [QueryLeaseRequest](#akash.market.v1beta5.QueryLeaseRequest) | [QueryLeaseResponse](#akash.market.v1beta5.QueryLeaseResponse) | Lease queries lease details. | GET|/akash/market/v1beta5/leases/info|
 | `ProviderLeaseStats` | [QueryProviderLeaseStatsRequest](#akash.market.v1beta5.QueryProviderLeaseStatsRequest) | [QueryProviderLeaseStatsResponse](#akash.market.v1beta5.QueryProviderLeaseStatsResponse) | ProviderLeaseStats queries aggregate lease-completion stats for a provider. | GET|/akash/market/v1beta5/providers/{provider}/lease-stats|
 | `Params` | [QueryParamsRequest](#akash.market.v1beta5.QueryParamsRequest) | [QueryParamsResponse](#akash.market.v1beta5.QueryParamsResponse) | Params returns the total set of market parameters. | GET|/akash/market/v1beta5/params|
 
  <!-- end services -->

 
 
 <a name="akash/market/v1beta5/service.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/market/v1beta5/service.proto
 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.market.v1beta5.Msg"></a>

 ### Msg
 Msg defines the market Msg service.

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `CreateBid` | [MsgCreateBid](#akash.market.v1beta5.MsgCreateBid) | [MsgCreateBidResponse](#akash.market.v1beta5.MsgCreateBidResponse) | CreateBid defines a method to create a bid given proper inputs. | |
 | `CloseBid` | [MsgCloseBid](#akash.market.v1beta5.MsgCloseBid) | [MsgCloseBidResponse](#akash.market.v1beta5.MsgCloseBidResponse) | CloseBid defines a method to close a bid given proper inputs. | |
 | `WithdrawLease` | [MsgWithdrawLease](#akash.market.v1beta5.MsgWithdrawLease) | [MsgWithdrawLeaseResponse](#akash.market.v1beta5.MsgWithdrawLeaseResponse) | WithdrawLease withdraws accrued funds from the lease payment | |
 | `CreateLease` | [MsgCreateLease](#akash.market.v1beta5.MsgCreateLease) | [MsgCreateLeaseResponse](#akash.market.v1beta5.MsgCreateLeaseResponse) | CreateLease creates a new lease | |
 | `CloseLease` | [MsgCloseLease](#akash.market.v1beta5.MsgCloseLease) | [MsgCloseLeaseResponse](#akash.market.v1beta5.MsgCloseLeaseResponse) | CloseLease defines a method to close an order given proper inputs. | |
 | `LeaseStartReclaim` | [MsgLeaseStartReclaim](#akash.market.v1beta5.MsgLeaseStartReclaim) | [MsgLeaseStartReclaimResponse](#akash.market.v1beta5.MsgLeaseStartReclaimResponse) | LeaseStartReclaim initiates the reclamation window on an active lease. | |
 | `UpdateParams` | [MsgUpdateParams](#akash.market.v1beta5.MsgUpdateParams) | [MsgUpdateParamsResponse](#akash.market.v1beta5.MsgUpdateParamsResponse) | UpdateParams defines a governance operation for updating the x/market module parameters. The authority is hard-coded to the x/gov module account.

Since: akash v1.0.0 | |
 
  <!-- end services -->

 
 
 <a name="akash/oracle/v1/prices.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/oracle/v1/prices.proto
 

 
 <a name="akash.oracle.v1.AggregatedPrice"></a>

 ### AggregatedPrice
 AggregatedPrice represents the final aggregated price from all sources

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `denom` | [string](#string) |  | denom is the asset denomination |
 | `twap` | [string](#string) |  | twap is the time-weighted average price over the configured window |
 | `median_price` | [string](#string) |  | median_price is the median of all source prices |
 | `min_price` | [string](#string) |  | min_price is the minimum price from all sources |
 | `max_price` | [string](#string) |  | max_price is the maximum price from all sources |
 | `timestamp` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | timestamp is when the aggregated price was computed |
 | `num_sources` | [uint32](#uint32) |  | num_sources is the number of price sources contributing to this aggregation |
 | `deviation_bps` | [uint64](#uint64) |  | deviation_bps is the price deviation in basis points between min and max prices |
 
 

 

 
 <a name="akash.oracle.v1.DataID"></a>

 ### DataID
 DataID uniquely identifies a price pair by asset and base denomination

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `denom` | [string](#string) |  | denom is the asset denomination (e.g., "uakt") |
 | `base_denom` | [string](#string) |  | base_denom is the base denomination for the price pair (e.g., "usd") |
 
 

 

 
 <a name="akash.oracle.v1.PriceData"></a>

 ### PriceData
 PriceData combines a price record identifier with its state

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [PriceDataRecordID](#akash.oracle.v1.PriceDataRecordID) |  | id uniquely identifies this price record |
 | `state` | [PriceDataState](#akash.oracle.v1.PriceDataState) |  | state contains the price value and timestamp |
 
 

 

 
 <a name="akash.oracle.v1.PriceDataID"></a>

 ### PriceDataID
 PriceDataID identifies price data from a specific source for a specific pair

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `source` | [uint32](#uint32) |  | source is the index of the price source (oracle provider) |
 | `denom` | [string](#string) |  | denom is the asset denomination |
 | `base_denom` | [string](#string) |  | base_denom is the base denomination for the price pair |
 
 

 

 
 <a name="akash.oracle.v1.PriceDataRecordID"></a>

 ### PriceDataRecordID
 PriceDataRecordID represents a price from a specific source at a specific time.
It also represents a single data point in TWAP history

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `source` | [uint32](#uint32) |  | source is the index of the price source (oracle provider) |
 | `denom` | [string](#string) |  | denom is the asset denomination |
 | `base_denom` | [string](#string) |  | base_denom is the base denomination for the price pair |
 | `height` | [int64](#int64) |  | height is the block height when this price was recorded |
 
 

 

 
 <a name="akash.oracle.v1.PriceDataState"></a>

 ### PriceDataState
 PriceDataState represents the price value and timestamp for a price entry

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `price` | [string](#string) |  | price is the decimal price value |
 | `timestamp` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | timestamp is when the price was recorded |
 
 

 

 
 <a name="akash.oracle.v1.PriceHealth"></a>

 ### PriceHealth
 PriceHealth represents the health status of a price feed

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `denom` | [string](#string) |  | denom is the asset denomination |
 | `is_healthy` | [bool](#bool) |  | is_healthy indicates if the price feed meets all health requirements |
 | `has_min_sources` | [bool](#bool) |  | has_min_sources indicates if minimum number of sources are reporting |
 | `deviation_ok` | [bool](#bool) |  | deviation_ok indicates if price deviation is within acceptable limits |
 | `total_sources` | [uint32](#uint32) |  | total_sources indicates total amount of sources registered for price calculations |
 | `total_healthy_sources` | [uint32](#uint32) |  | total_healthy_sources indicates total usable sources for price calculations |
 | `failure_reason` | [string](#string) | repeated | failure_reason lists reasons for unhealthy status, if any |
 
 

 

 
 <a name="akash.oracle.v1.PricesFilter"></a>

 ### PricesFilter
 PricesFilter defines filters used to query price data

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `asset_denom` | [string](#string) |  | asset_denom is the asset denomination to filter by |
 | `base_denom` | [string](#string) |  | base_denom is the base denomination to filter by |
 | `height` | [int64](#int64) |  | height is the block height to filter by |
 
 

 

 
 <a name="akash.oracle.v1.QueryPricesRequest"></a>

 ### QueryPricesRequest
 QueryPricesRequest is the request type for querying price history

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `filters` | [PricesFilter](#akash.oracle.v1.PricesFilter) |  | filters holds the price fields to filter the request |
 | `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination is used to paginate the request |
 
 

 

 
 <a name="akash.oracle.v1.QueryPricesResponse"></a>

 ### QueryPricesResponse
 QueryPricesResponse is the response type for querying price history

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `prices` | [PriceData](#akash.oracle.v1.PriceData) | repeated | prices is the list of historical price data matching the filters |
 | `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination contains the information about response pagination |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/oracle/v1/events.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/oracle/v1/events.proto
 

 
 <a name="akash.oracle.v1.EventAggregatedPrice"></a>

 ### EventAggregatedPrice
 EventAggregatedPrice is emitted when aggregated price has an update

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `price` | [AggregatedPrice](#akash.oracle.v1.AggregatedPrice) |  |  |
 
 

 

 
 <a name="akash.oracle.v1.EventPriceData"></a>

 ### EventPriceData
 EventPriceData is emitted when new price data is added to the oracle

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `source` | [string](#string) |  | source is the address of the price source (oracle provider) |
 | `id` | [DataID](#akash.oracle.v1.DataID) |  | id identifies the price pair (denom and base_denom) |
 | `data` | [PriceDataState](#akash.oracle.v1.PriceDataState) |  | data contains the price value and timestamp |
 
 

 

 
 <a name="akash.oracle.v1.EventPriceRecovered"></a>

 ### EventPriceRecovered
 EventPriceRecovered is emitted when a stale price has started receiving updates again

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [DataID](#akash.oracle.v1.DataID) |  | id identifies the price pair |
 | `height` | [int64](#int64) |  | height is the block height when the price recovery was detected |
 
 

 

 
 <a name="akash.oracle.v1.EventPriceStaleWarning"></a>

 ### EventPriceStaleWarning
 EventPriceStaleWarning is emitted when price has not been updated and is about to become stale

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `source` | [string](#string) |  | source is the address of the price source |
 | `id` | [DataID](#akash.oracle.v1.DataID) |  | id identifies the price pair |
 | `last_height` | [int64](#int64) |  | last_height is the block height when the price was last updated |
 | `blocks_to_stall` | [int64](#int64) |  | blocks_to_stall is the number of blocks until the price becomes stale |
 
 

 

 
 <a name="akash.oracle.v1.EventPriceStaled"></a>

 ### EventPriceStaled
 EventPriceStaled is emitted when a price has become stale

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [DataID](#akash.oracle.v1.DataID) |  | id identifies the price pair |
 | `last_height` | [int64](#int64) |  | last_height is the block height when the price was last updated before becoming stale |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/oracle/v1/params.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/oracle/v1/params.proto
 

 
 <a name="akash.oracle.v1.Params"></a>

 ### Params
 Params defines the parameters for the oracle module

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `sources` | [string](#string) | repeated | sources addresses allowed to write prices into oracle module those are to be smartcontract addresses |
 | `min_price_sources` | [uint32](#uint32) |  | Minimum number of price sources required (default: 2) |
 | `max_price_staleness_blocks` | [int64](#int64) |  | Maximum price staleness in blocks (default: 50 = ~ 5 minutes) |
 | `twap_window` | [int64](#int64) |  | TWAP window in blocks (default: 50 = ~ 5 minutes) |
 | `max_price_deviation_bps` | [uint64](#uint64) |  | Maximum price deviation in basis points (default: 150 = 1.5%) |
 | `feed_contracts_params` | [google.protobuf.Any](#google.protobuf.Any) | repeated | feed_contracts_params contains the configuration for the price feed contracts |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/oracle/v1/genesis.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/oracle/v1/genesis.proto
 

 
 <a name="akash.oracle.v1.GenesisState"></a>

 ### GenesisState
 GenesisState defines the oracle module's genesis state

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `params` | [Params](#akash.oracle.v1.Params) |  | params holds the oracle module parameters |
 | `prices` | [PriceData](#akash.oracle.v1.PriceData) | repeated | prices is the list of all historical price data entries |
 | `latest_height` | [PriceDataID](#akash.oracle.v1.PriceDataID) | repeated | latest_height tracks the most recent block height for each price feed source |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/oracle/v1/msgs.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/oracle/v1/msgs.proto
 

 
 <a name="akash.oracle.v1.MsgAddPriceEntry"></a>

 ### MsgAddPriceEntry
 MsgAddPriceEntry defines an SDK message to add oracle price entry.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `signer` | [string](#string) |  | Signer is the bech32 address of the account of the provider. It is a string representing a valid account address.

Example: "akash1..." |
 | `id` | [DataID](#akash.oracle.v1.DataID) |  | id uniquely identifies the price data by denomination and base denomination |
 | `price` | [PriceDataState](#akash.oracle.v1.PriceDataState) |  | price contains the price value and timestamp for this entry |
 
 

 

 
 <a name="akash.oracle.v1.MsgAddPriceEntryResponse"></a>

 ### MsgAddPriceEntryResponse
 MsgAddPriceEntryResponse defines the Msg/MsgAddDPriceEntry response type.

 

 

 
 <a name="akash.oracle.v1.MsgUpdateParams"></a>

 ### MsgUpdateParams
 MsgUpdateParams is the Msg/UpdateParams request type.

Since: akash v2.0.0

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `authority` | [string](#string) |  | authority is the address of the governance account. |
 | `params` | [Params](#akash.oracle.v1.Params) |  | params defines the x/oracle parameters to update.

NOTE: All parameters must be supplied. |
 
 

 

 
 <a name="akash.oracle.v1.MsgUpdateParamsResponse"></a>

 ### MsgUpdateParamsResponse
 MsgUpdateParamsResponse defines the response structure for executing a
MsgUpdateParams message.

Since: akash v2.0.0

 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/oracle/v1/query.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/oracle/v1/query.proto
 

 
 <a name="akash.oracle.v1.QueryAggregatedPriceRequest"></a>

 ### QueryAggregatedPriceRequest
 QueryAggregatedPriceRequest is the request type for aggregated price.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `denom` | [string](#string) |  | denom is the asset denomination |
 
 

 

 
 <a name="akash.oracle.v1.QueryAggregatedPriceResponse"></a>

 ### QueryAggregatedPriceResponse
 QueryAggregatedPriceResponse is the response type for aggregated price.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `aggregated_price` | [AggregatedPrice](#akash.oracle.v1.AggregatedPrice) |  | aggregated_price is the aggregated price data |
 | `price_health` | [PriceHealth](#akash.oracle.v1.PriceHealth) |  | price_health is the health status for the price feed |
 
 

 

 
 <a name="akash.oracle.v1.QueryParamsRequest"></a>

 ### QueryParamsRequest
 QueryParamsRequest is the request type for the Query/Params RPC method.

 

 

 
 <a name="akash.oracle.v1.QueryParamsResponse"></a>

 ### QueryParamsResponse
 QueryParamsResponse is the response type for the Query/Params RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `params` | [Params](#akash.oracle.v1.Params) |  | params defines the parameters of the module. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.oracle.v1.Query"></a>

 ### Query
 Query defines the gRPC querier service of the oracle package.

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `Prices` | [QueryPricesRequest](#akash.oracle.v1.QueryPricesRequest) | [QueryPricesResponse](#akash.oracle.v1.QueryPricesResponse) | Prices query prices for specific denom | GET|/akash/oracle/v1/prices|
 | `Params` | [QueryParamsRequest](#akash.oracle.v1.QueryParamsRequest) | [QueryParamsResponse](#akash.oracle.v1.QueryParamsResponse) | Params returns the total set of oracle parameters. | GET|/akash/oracle/v1/params|
 | `AggregatedPrice` | [QueryAggregatedPriceRequest](#akash.oracle.v1.QueryAggregatedPriceRequest) | [QueryAggregatedPriceResponse](#akash.oracle.v1.QueryAggregatedPriceResponse) | AggregatedPrice queries the aggregated price for a given denom. | GET|/akash/oracle/v1/aggregated_price/{denom}|
 
  <!-- end services -->

 
 
 <a name="akash/oracle/v1/service.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/oracle/v1/service.proto
 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.oracle.v1.Msg"></a>

 ### Msg
 Msg defines the oracle Msg service for managing price feeds

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `AddPriceEntry` | [MsgAddPriceEntry](#akash.oracle.v1.MsgAddPriceEntry) | [MsgAddPriceEntryResponse](#akash.oracle.v1.MsgAddPriceEntryResponse) | AddPriceEntry adds a new price entry for a denomination from an authorized source | |
 | `UpdateParams` | [MsgUpdateParams](#akash.oracle.v1.MsgUpdateParams) | [MsgUpdateParamsResponse](#akash.oracle.v1.MsgUpdateParamsResponse) | UpdateParams defines a governance operation for updating the x/wasm module parameters. The authority is hard-coded to the x/gov module account.

Since: akash v2.0.0 | |
 
  <!-- end services -->

 
 
 <a name="akash/oracle/v2/prices.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/oracle/v2/prices.proto
 

 
 <a name="akash.oracle.v2.AggregatedPrice"></a>

 ### AggregatedPrice
 AggregatedPrice represents the final aggregated price from all sources

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `denom` | [string](#string) |  | denom is the asset denomination |
 | `twap` | [string](#string) |  | twap is the time-weighted average price over the configured window |
 | `median_price` | [string](#string) |  | median_price is the median of all source prices |
 | `min_price` | [string](#string) |  | min_price is the minimum price from all sources |
 | `max_price` | [string](#string) |  | max_price is the maximum price from all sources |
 | `timestamp` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | timestamp is when the aggregated price was computed |
 | `num_sources` | [uint32](#uint32) |  | num_sources is the number of price sources contributing to this aggregation |
 | `deviation_bps` | [uint64](#uint64) |  | deviation_bps is the price deviation in basis points between min and max prices |
 
 

 

 
 <a name="akash.oracle.v2.DataID"></a>

 ### DataID
 DataID uniquely identifies a price pair by asset and base denomination

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `denom` | [string](#string) |  | denom is the asset denomination (e.g., "uakt") |
 | `base_denom` | [string](#string) |  | base_denom is the base denomination for the price pair (e.g., "usd") |
 
 

 

 
 <a name="akash.oracle.v2.PriceData"></a>

 ### PriceData
 PriceData combines a price record identifier with its state

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [PriceDataRecordID](#akash.oracle.v2.PriceDataRecordID) |  | id uniquely identifies this price record |
 | `state` | [PriceDataState](#akash.oracle.v2.PriceDataState) |  | state contains the price value and timestamp |
 
 

 

 
 <a name="akash.oracle.v2.PriceDataID"></a>

 ### PriceDataID
 PriceDataID identifies price data from a specific source for a specific pair

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `source` | [uint32](#uint32) |  | source is the index of the price source (oracle provider) |
 | `denom` | [string](#string) |  | denom is the asset denomination |
 | `base_denom` | [string](#string) |  | base_denom is the base denomination for the price pair |
 
 

 

 
 <a name="akash.oracle.v2.PriceDataRecordID"></a>

 ### PriceDataRecordID
 PriceDataRecordID represents a price from a specific source at a specific time.
It also represents a single data point in TWAP history

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `source` | [uint32](#uint32) |  | source is the index of the price source (oracle provider) |
 | `denom` | [string](#string) |  | denom is the asset denomination |
 | `base_denom` | [string](#string) |  | base_denom is the base denomination for the price pair |
 | `timestamp` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | timestamp is when this price was recorded |
 | `sequence` | [uint64](#uint64) |  | sequence disambiguates multiple price entries at the same timestamp |
 
 

 

 
 <a name="akash.oracle.v2.PriceDataState"></a>

 ### PriceDataState
 PriceDataState represents the price value

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `price` | [string](#string) |  | price is the decimal price value |
 
 

 

 
 <a name="akash.oracle.v2.PriceHealth"></a>

 ### PriceHealth
 PriceHealth represents the health status of a price feed

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `denom` | [string](#string) |  | denom is the asset denomination |
 | `is_healthy` | [bool](#bool) |  | is_healthy indicates if the price feed meets all health requirements |
 | `has_min_sources` | [bool](#bool) |  | has_min_sources indicates if minimum number of sources are reporting |
 | `deviation_ok` | [bool](#bool) |  | deviation_ok indicates if price deviation is within acceptable limits |
 | `total_sources` | [uint32](#uint32) |  | total_sources indicates total amount of sources registered for price calculations |
 | `total_healthy_sources` | [uint32](#uint32) |  | total_healthy_sources indicates total usable sources for price calculations |
 | `failure_reason` | [string](#string) | repeated | failure_reason lists reasons for unhealthy status, if any |
 
 

 

 
 <a name="akash.oracle.v2.PriceLatestDataState"></a>

 ### PriceLatestDataState
 PriceLatestDataState holds the timestamp of the most recent price record for a source/pair

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `timestamp` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | timestamp is when the price was recorded |
 | `sequence` | [uint64](#uint64) |  | sequence disambiguates multiple price entries at the same timestamp |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/oracle/v2/events.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/oracle/v2/events.proto
 

 
 <a name="akash.oracle.v2.EventAggregatedPrice"></a>

 ### EventAggregatedPrice
 EventAggregatedPrice is emitted when aggregated price has an update

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `price` | [AggregatedPrice](#akash.oracle.v2.AggregatedPrice) |  | price is the aggregated price data |
 
 

 

 
 <a name="akash.oracle.v2.EventPriceData"></a>

 ### EventPriceData
 EventPriceData is emitted when new price data is added to the oracle

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `source` | [string](#string) |  | source is the address of the price source (oracle provider) |
 | `id` | [DataID](#akash.oracle.v2.DataID) |  | id identifies the price pair (denom and base_denom) |
 | `price` | [string](#string) |  | price is the decimal price value |
 | `timestamp` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | timestamp is when this price was recorded |
 
 

 

 
 <a name="akash.oracle.v2.EventPriceRecovered"></a>

 ### EventPriceRecovered
 EventPriceRecovered is emitted when a stale price has started receiving updates again

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [DataID](#akash.oracle.v2.DataID) |  | id identifies the price pair |
 | `height` | [int64](#int64) |  | height is the block height when the price recovery was detected |
 
 

 

 
 <a name="akash.oracle.v2.EventPriceStaleWarning"></a>

 ### EventPriceStaleWarning
 EventPriceStaleWarning is emitted when price has not been updated and is about to become stale

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [DataID](#akash.oracle.v2.DataID) |  | id identifies the price pair |
 | `last_height` | [int64](#int64) |  | last_height is the block height when the price was last updated |
 | `blocks_to_stall` | [int64](#int64) |  | blocks_to_stall is the number of blocks until the price becomes stale |
 
 

 

 
 <a name="akash.oracle.v2.EventPriceStaled"></a>

 ### EventPriceStaled
 EventPriceStaled is emitted when a price has become stale

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [DataID](#akash.oracle.v2.DataID) |  | id identifies the price pair |
 | `last_height` | [int64](#int64) |  | last_height is the block height when the price was last updated before becoming stale |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/oracle/v2/params.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/oracle/v2/params.proto
 

 
 <a name="akash.oracle.v2.Params"></a>

 ### Params
 Params defines the parameters for the oracle module

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `sources` | [string](#string) | repeated | sources addresses allowed to write prices into oracle module those are to be smartcontract addresses |
 | `min_price_sources` | [uint32](#uint32) |  | Minimum number of price sources required (default: 1) |
 | `max_price_staleness_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  | Maximum price staleness in seconds (default: 60s) |
 | `twap_window` | [google.protobuf.Duration](#google.protobuf.Duration) |  | TWAP window as a duration (default: 5s) |
 | `max_price_deviation_bps` | [uint64](#uint64) |  | Maximum price deviation in basis points (default: 150 = 1.5%) |
 | `feed_contracts_params` | [google.protobuf.Any](#google.protobuf.Any) | repeated | feed_contracts_params contains the configuration for the price feed contracts |
 | `price_retention` | [google.protobuf.Duration](#google.protobuf.Duration) |  | price_retention is how long to keep price records (default: 24h) |
 | `prune_epoch` | [string](#string) |  | prune_epoch is the epoch identifier that triggers pruning (default: "hour") |
 | `max_prune_per_epoch` | [int64](#int64) |  | max_prune_per_epoch is the max records to delete per epoch pruning pass (default: 1000) |
 | `max_future_time_drift` | [google.protobuf.Duration](#google.protobuf.Duration) |  | max_future_time_drift is the maximum amount of time a price timestamp may exceed the current block time (default: 1m) |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/oracle/v2/genesis.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/oracle/v2/genesis.proto
 

 
 <a name="akash.oracle.v2.GenesisLatestPricesIDs"></a>

 ### GenesisLatestPricesIDs
 GenesisLatestPricesIDs stores the latest price state for a given price pair

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [PriceDataID](#akash.oracle.v2.PriceDataID) |  | id identifies the price pair (source, denom, base_denom) |
 | `state` | [PriceLatestDataState](#akash.oracle.v2.PriceLatestDataState) |  | state holds the timestamp of the latest price record |
 
 

 

 
 <a name="akash.oracle.v2.GenesisSourceID"></a>

 ### GenesisSourceID
 GenesisSourceID maps an oracle source address to its numeric identifier

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `address` | [string](#string) |  | address is the bech32 address of the oracle source |
 | `id` | [uint32](#uint32) |  | id is the numeric identifier assigned to the source |
 
 

 

 
 <a name="akash.oracle.v2.GenesisState"></a>

 ### GenesisState
 GenesisState defines the oracle module's genesis state

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `params` | [Params](#akash.oracle.v2.Params) |  | params holds the oracle module parameters |
 | `prices` | [PriceData](#akash.oracle.v2.PriceData) | repeated | prices is the list of all historical price data entries |
 | `latest_prices_ids` | [GenesisLatestPricesIDs](#akash.oracle.v2.GenesisLatestPricesIDs) | repeated | latest_height tracks the most recent block height for each price feed source |
 | `source_ids` | [GenesisSourceID](#akash.oracle.v2.GenesisSourceID) | repeated | source_ids is the list of oracle source address-to-ID mappings |
 | `source_seq` | [uint64](#uint64) |  | source_seq is the next available source identifier sequence number |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/oracle/v2/msgs.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/oracle/v2/msgs.proto
 

 
 <a name="akash.oracle.v2.MsgAddPriceEntry"></a>

 ### MsgAddPriceEntry
 MsgAddPriceEntry defines an SDK message to add oracle price entry.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `signer` | [string](#string) |  | Signer is the bech32 address of the account of the provider. It is a string representing a valid account address.

Example: "akash1..." |
 | `id` | [DataID](#akash.oracle.v2.DataID) |  | id uniquely identifies the price data by denomination and base denomination |
 | `price` | [string](#string) |  | price is the decimal price value |
 | `timestamp` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | timestamp is when this price was observed |
 
 

 

 
 <a name="akash.oracle.v2.MsgAddPriceEntryResponse"></a>

 ### MsgAddPriceEntryResponse
 MsgAddPriceEntryResponse defines the Msg/MsgAddPriceEntry response type.

 

 

 
 <a name="akash.oracle.v2.MsgUpdateParams"></a>

 ### MsgUpdateParams
 MsgUpdateParams is the Msg/UpdateParams request type.

Since: akash v2.0.0

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `authority` | [string](#string) |  | authority is the address of the governance account. |
 | `params` | [Params](#akash.oracle.v2.Params) |  | params defines the x/oracle parameters to update.

NOTE: All parameters must be supplied. |
 
 

 

 
 <a name="akash.oracle.v2.MsgUpdateParamsResponse"></a>

 ### MsgUpdateParamsResponse
 MsgUpdateParamsResponse defines the response structure for executing a
MsgUpdateParams message.

Since: akash v2.0.0

 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/oracle/v2/query.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/oracle/v2/query.proto
 

 
 <a name="akash.oracle.v2.PricesFilter"></a>

 ### PricesFilter
 PricesFilter defines filters used to query price data

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `asset_denom` | [string](#string) |  | asset_denom is the asset denomination to filter by |
 | `base_denom` | [string](#string) |  | base_denom is the base denomination to filter by |
 | `start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | start_time is the inclusive start of the time range to filter by |
 | `end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | end_time is the inclusive end of the time range to filter by |
 
 

 

 
 <a name="akash.oracle.v2.QueryAggregatedPriceRequest"></a>

 ### QueryAggregatedPriceRequest
 QueryAggregatedPriceRequest is the request type for aggregated price.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `denom` | [string](#string) |  | denom is the asset denomination |
 
 

 

 
 <a name="akash.oracle.v2.QueryAggregatedPriceResponse"></a>

 ### QueryAggregatedPriceResponse
 QueryAggregatedPriceResponse is the response type for aggregated price.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `aggregated_price` | [AggregatedPrice](#akash.oracle.v2.AggregatedPrice) |  | aggregated_price is the aggregated price data |
 | `price_health` | [PriceHealth](#akash.oracle.v2.PriceHealth) |  | price_health is the health status for the price feed |
 
 

 

 
 <a name="akash.oracle.v2.QueryParamsRequest"></a>

 ### QueryParamsRequest
 QueryParamsRequest is the request type for the Query/Params RPC method.

 

 

 
 <a name="akash.oracle.v2.QueryParamsResponse"></a>

 ### QueryParamsResponse
 QueryParamsResponse is the response type for the Query/Params RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `params` | [Params](#akash.oracle.v2.Params) |  | params defines the parameters of the module. |
 
 

 

 
 <a name="akash.oracle.v2.QueryPricesRequest"></a>

 ### QueryPricesRequest
 QueryPricesRequest is the request type for querying price history

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `filters` | [PricesFilter](#akash.oracle.v2.PricesFilter) |  | filters holds the price fields to filter the request |
 | `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination is used to paginate the request |
 
 

 

 
 <a name="akash.oracle.v2.QueryPricesResponse"></a>

 ### QueryPricesResponse
 QueryPricesResponse is the response type for querying price history

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `prices` | [PriceData](#akash.oracle.v2.PriceData) | repeated | prices is the list of historical price data matching the filters |
 | `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination contains the information about response pagination |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.oracle.v2.Query"></a>

 ### Query
 Query defines the gRPC querier service of the oracle package.

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `Prices` | [QueryPricesRequest](#akash.oracle.v2.QueryPricesRequest) | [QueryPricesResponse](#akash.oracle.v2.QueryPricesResponse) | Prices query prices for specific denom | GET|/akash/oracle/v2/prices|
 | `Params` | [QueryParamsRequest](#akash.oracle.v2.QueryParamsRequest) | [QueryParamsResponse](#akash.oracle.v2.QueryParamsResponse) | Params returns the total set of oracle parameters. | GET|/akash/oracle/v2/params|
 | `AggregatedPrice` | [QueryAggregatedPriceRequest](#akash.oracle.v2.QueryAggregatedPriceRequest) | [QueryAggregatedPriceResponse](#akash.oracle.v2.QueryAggregatedPriceResponse) | AggregatedPrice queries the aggregated price for a given denom. | GET|/akash/oracle/v2/aggregated_price/{denom=**}|
 
  <!-- end services -->

 
 
 <a name="akash/oracle/v2/service.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/oracle/v2/service.proto
 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.oracle.v2.Msg"></a>

 ### Msg
 Msg defines the oracle Msg service for managing price feeds

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `AddPriceEntry` | [MsgAddPriceEntry](#akash.oracle.v2.MsgAddPriceEntry) | [MsgAddPriceEntryResponse](#akash.oracle.v2.MsgAddPriceEntryResponse) | AddPriceEntry adds a new price entry for a denomination from an authorized source | |
 | `UpdateParams` | [MsgUpdateParams](#akash.oracle.v2.MsgUpdateParams) | [MsgUpdateParamsResponse](#akash.oracle.v2.MsgUpdateParamsResponse) | UpdateParams defines a governance operation for updating the x/oracle module parameters. The authority is hard-coded to the x/gov module account.

Since: akash v2.0.0 | |
 
  <!-- end services -->

 
 
 <a name="akash/provider/v1beta4/maintenance.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/provider/v1beta4/maintenance.proto
 

 
 <a name="akash.provider.v1beta4.ProviderMaintenanceRecord"></a>

 ### ProviderMaintenanceRecord
 ProviderMaintenanceRecord is an on-chain provider maintenance record.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [uint64](#uint64) |  | id is the maintenance identifier. |
 | `provider` | [string](#string) |  | provider is the bech32 address of the provider owning the maintenance window.

Example: "akash1..." |
 | `maintenance_type` | [ProviderMaintenanceType](#akash.provider.v1beta4.ProviderMaintenanceType) |  | maintenance_type is the declared category of the window. |
 | `starts_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | starts_at is the wall-clock time at which the maintenance window begins. |
 | `expected_ends_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | expected_ends_at is the wall-clock time at which the provider expects the window to end. |
 | `opened_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | opened_at is the block time at which the window was opened. |
 | `closed_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | closed_at is the block time at which the window was closed. |
 | `metadata_hash` | [bytes](#bytes) |  | metadata_hash is an optional, opaque hash of off-chain metadata. |
 
 

 

 
 <a name="akash.provider.v1beta4.ProviderMaintenanceWithStatus"></a>

 ### ProviderMaintenanceWithStatus
 ProviderMaintenanceWithStatus pairs a maintenance record with its status.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `record` | [ProviderMaintenanceRecord](#akash.provider.v1beta4.ProviderMaintenanceRecord) |  | record is the stored maintenance window. |
 | `status` | [ProviderMaintenanceStatus](#akash.provider.v1beta4.ProviderMaintenanceStatus) |  | status is the derived lifecycle state of the record at query time. |
 
 

 

  <!-- end messages -->

 
 <a name="akash.provider.v1beta4.ProviderMaintenanceStatus"></a>

 ### ProviderMaintenanceStatus
 ProviderMaintenanceStatus enumerates provider maintenance lifecycle states.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | provider_maintenance_status_unspecified | 0 | provider_maintenance_status_unspecified is the zero value. |
 | provider_maintenance_status_scheduled | 1 | provider_maintenance_status_scheduled means the window has not started. |
 | provider_maintenance_status_active | 2 | provider_maintenance_status_active means the window is active. |
 | provider_maintenance_status_elapsed | 3 | provider_maintenance_status_elapsed means the window reached expected_ends_at. |
 | provider_maintenance_status_closed | 4 | provider_maintenance_status_closed means the window was closed explicitly. |
 

 
 <a name="akash.provider.v1beta4.ProviderMaintenanceType"></a>

 ### ProviderMaintenanceType
 ProviderMaintenanceType enumerates provider maintenance window types.

 | Name | Number | Description |
 | ---- | ------ | ----------- |
 | provider_maintenance_type_unspecified | 0 | provider_maintenance_type_unspecified is the zero value. |
 | provider_maintenance_type_planned | 1 | provider_maintenance_type_planned represents a scheduled, non-urgent maintenance window communicated to tenants ahead of time. |
 | provider_maintenance_type_emergency | 2 | provider_maintenance_type_emergency represents an urgent, unplanned maintenance window driven by operational incidents. |
 | provider_maintenance_type_security | 3 | provider_maintenance_type_security represents a window opened to apply a security patch or to remediate a security event. |
 | provider_maintenance_type_network | 4 | provider_maintenance_type_network represents a window driven by network connectivity work (e.g., upstream provider, peering, or DNS changes). |
 | provider_maintenance_type_capacity | 5 | provider_maintenance_type_capacity represents a window opened to perform capacity changes such as adding, draining, or removing hardware. |
 

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/provider/v1beta4/event.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/provider/v1beta4/event.proto
 

 
 <a name="akash.provider.v1beta4.EventProviderCreated"></a>

 ### EventProviderCreated
 EventProviderCreated defines an SDK message for provider created event.
It contains all the required information to identify a provider on-chain.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the bech32 address of the account of the provider. It is a string representing a valid account address.

Example: "akash1..." |
 
 

 

 
 <a name="akash.provider.v1beta4.EventProviderDeleted"></a>

 ### EventProviderDeleted
 EventProviderDeleted defines an SDK message for provider deleted event.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the bech32 address of the account of the provider. It is a string representing a valid account address.

Example: "akash1..." |
 
 

 

 
 <a name="akash.provider.v1beta4.EventProviderMaintenanceClosed"></a>

 ### EventProviderMaintenanceClosed
 EventProviderMaintenanceClosed is emitted when provider maintenance closes.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `maintenance_id` | [uint64](#uint64) |  | maintenance_id is the identifier of the closed maintenance record. |
 | `provider` | [string](#string) |  | provider is the bech32 address of the provider that closed the window.

Example: "akash1..." |
 | `closed_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | closed_at is the block time at which the window closed. |
 
 

 

 
 <a name="akash.provider.v1beta4.EventProviderMaintenanceOpened"></a>

 ### EventProviderMaintenanceOpened
 EventProviderMaintenanceOpened is emitted when provider maintenance opens.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `maintenance_id` | [uint64](#uint64) |  | maintenance_id is the identifier assigned to the new maintenance record. |
 | `provider` | [string](#string) |  | provider is the bech32 address of the provider that opened the window.

Example: "akash1..." |
 | `maintenance_type` | [ProviderMaintenanceType](#akash.provider.v1beta4.ProviderMaintenanceType) |  | maintenance_type is the declared category of the window. |
 | `starts_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | starts_at is the wall-clock time at which the maintenance window begins. |
 | `expected_ends_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | expected_ends_at is the wall-clock time at which the provider expects the window to end. |
 | `metadata_hash` | [bytes](#bytes) |  | metadata_hash is the optional, opaque hash of off-chain metadata. |
 
 

 

 
 <a name="akash.provider.v1beta4.EventProviderUpdated"></a>

 ### EventProviderUpdated
 EventProviderUpdated defines an SDK message for provider updated event.
It contains all the required information to identify a provider on-chain.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the bech32 address of the account of the provider. It is a string representing a valid account address.

Example: "akash1..." |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/provider/v1beta4/provider.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/provider/v1beta4/provider.proto
 

 
 <a name="akash.provider.v1beta4.Info"></a>

 ### Info
 Info contains information on the provider.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `email` | [string](#string) |  | Email is the email address to contact the provider. |
 | `website` | [string](#string) |  | Website is the URL to the landing page or socials of the provider. |
 
 

 

 
 <a name="akash.provider.v1beta4.Provider"></a>

 ### Provider
 Provider stores owner and host details.
Akash providers are entities that contribute computing resources to the network.
They can be individuals or organizations with underutilized computing resources, such as data centers or personal servers.
Providers participate in the network by running the Akash node software and setting the price for their services.
Users can then choose a provider based on factors such as cost, performance, and location.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the bech32 address of the account of the provider. It is a string representing a valid account address.

Example: "akash1..." |
 | `host_uri` | [string](#string) |  | HostURI is the Uniform Resource Identifier for provider connection. This URI is used to directly connect to the provider to perform tasks such as sending the manifest. |
 | `attributes` | [akash.base.attributes.v1.Attribute](#akash.base.attributes.v1.Attribute) | repeated | Attributes is a list of arbitrary attribute key-value pairs. |
 | `info` | [Info](#akash.provider.v1beta4.Info) |  | Info contains additional provider information. |
 
 

 

 
 <a name="akash.provider.v1beta4.ProviderRegistration"></a>

 ### ProviderRegistration
 ProviderRegistration captures when a provider was registered.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | owner is the bech32 address of the provider account.

Example: "akash1..." |
 | `registered_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | registered_at is the block time at which the provider was registered. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/provider/v1beta4/params.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/provider/v1beta4/params.proto
 

 
 <a name="akash.provider.v1beta4.ProviderMaintenanceParams"></a>

 ### ProviderMaintenanceParams
 ProviderMaintenanceParams defines maintenance window parameters.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `maintenance_max_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  | maintenance_max_duration is the maximum allowed value of (expected_ends_at - starts_at) when opening a maintenance window. |
 | `maintenance_max_lookahead` | [google.protobuf.Duration](#google.protobuf.Duration) |  | maintenance_max_lookahead is the maximum allowed value of (starts_at - block_time) when opening a maintenance window. A value of zero means windows may only start at or before the current block time. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/provider/v1beta4/genesis.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/provider/v1beta4/genesis.proto
 

 
 <a name="akash.provider.v1beta4.GenesisState"></a>

 ### GenesisState
 GenesisState defines the basic genesis state used by provider module.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `providers` | [Provider](#akash.provider.v1beta4.Provider) | repeated | Providers is a list of genesis providers. |
 | `params` | [ProviderMaintenanceParams](#akash.provider.v1beta4.ProviderMaintenanceParams) |  | Params is the provider module parameter set. |
 | `maintenances` | [ProviderMaintenanceRecord](#akash.provider.v1beta4.ProviderMaintenanceRecord) | repeated | Maintenances is the list of provider maintenance records. |
 | `next_maintenance_id` | [uint64](#uint64) |  | NextMaintenanceID is the next provider maintenance identifier. |
 | `registrations` | [ProviderRegistration](#akash.provider.v1beta4.ProviderRegistration) | repeated | Registrations is the list of provider registration records to import at genesis. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/provider/v1beta4/msg.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/provider/v1beta4/msg.proto
 

 
 <a name="akash.provider.v1beta4.MsgCloseProviderMaintenance"></a>

 ### MsgCloseProviderMaintenance
 MsgCloseProviderMaintenance closes an open provider maintenance window.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | provider is the bech32 address of the provider closing the maintenance window.

Example: "akash1..." |
 | `maintenance_id` | [uint64](#uint64) |  | maintenance_id is the identifier of the record to close. |
 
 

 

 
 <a name="akash.provider.v1beta4.MsgCloseProviderMaintenanceResponse"></a>

 ### MsgCloseProviderMaintenanceResponse
 MsgCloseProviderMaintenanceResponse is the response type for
MsgCloseProviderMaintenance.

 

 

 
 <a name="akash.provider.v1beta4.MsgCreateProvider"></a>

 ### MsgCreateProvider
 MsgCreateProvider defines an SDK message for creating a provider.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the bech32 address of the account of the provider. It is a string representing a valid account address.

Example: "akash1..." |
 | `host_uri` | [string](#string) |  | HostURI is the Uniform Resource Identifier for provider connection. This URI is used to directly connect to the provider to perform tasks such as sending the manifest. |
 | `attributes` | [akash.base.attributes.v1.Attribute](#akash.base.attributes.v1.Attribute) | repeated | Attributes is a list of arbitrary attribute key-value pairs. |
 | `info` | [Info](#akash.provider.v1beta4.Info) |  | Info contains additional provider information. |
 
 

 

 
 <a name="akash.provider.v1beta4.MsgCreateProviderResponse"></a>

 ### MsgCreateProviderResponse
 MsgCreateProviderResponse defines the Msg/CreateProvider response type.

 

 

 
 <a name="akash.provider.v1beta4.MsgDeleteProvider"></a>

 ### MsgDeleteProvider
 MsgDeleteProvider defines an SDK message for deleting a provider

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  |  |
 
 

 

 
 <a name="akash.provider.v1beta4.MsgDeleteProviderResponse"></a>

 ### MsgDeleteProviderResponse
 MsgDeleteProviderResponse defines the Msg/DeleteProvider response type.

 

 

 
 <a name="akash.provider.v1beta4.MsgOpenProviderMaintenance"></a>

 ### MsgOpenProviderMaintenance
 MsgOpenProviderMaintenance opens a maintenance window for a provider.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | provider is the bech32 address of the provider opening the maintenance window.

Example: "akash1..." |
 | `maintenance_type` | [ProviderMaintenanceType](#akash.provider.v1beta4.ProviderMaintenanceType) |  | maintenance_type is the declared category of the window. |
 | `starts_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | starts_at is the wall-clock time at which the maintenance window begins. |
 | `expected_ends_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | expected_ends_at is the wall-clock time at which the provider expects the window to end. |
 | `metadata_hash` | [bytes](#bytes) |  | metadata_hash is an optional, opaque commitment to off-chain explanatory metadata. The chain does not interpret this value. |
 
 

 

 
 <a name="akash.provider.v1beta4.MsgOpenProviderMaintenanceResponse"></a>

 ### MsgOpenProviderMaintenanceResponse
 MsgOpenProviderMaintenanceResponse is the response type for
MsgOpenProviderMaintenance.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `maintenance_id` | [uint64](#uint64) |  | maintenance_id is the identifier assigned to the maintenance window. |
 
 

 

 
 <a name="akash.provider.v1beta4.MsgUpdateProvider"></a>

 ### MsgUpdateProvider
 MsgUpdateProvider defines an SDK message for updating a provider

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  |  |
 | `host_uri` | [string](#string) |  |  |
 | `attributes` | [akash.base.attributes.v1.Attribute](#akash.base.attributes.v1.Attribute) | repeated |  |
 | `info` | [Info](#akash.provider.v1beta4.Info) |  |  |
 
 

 

 
 <a name="akash.provider.v1beta4.MsgUpdateProviderResponse"></a>

 ### MsgUpdateProviderResponse
 MsgUpdateProviderResponse defines the Msg/UpdateProvider response type.

 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/provider/v1beta4/paramsmsg.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/provider/v1beta4/paramsmsg.proto
 

 
 <a name="akash.provider.v1beta4.MsgUpdateParams"></a>

 ### MsgUpdateParams
 MsgUpdateParams is the Msg/UpdateParams request type.

Since: akash v1.0.0

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `authority` | [string](#string) |  | authority is the address of the governance account. |
 | `params` | [ProviderMaintenanceParams](#akash.provider.v1beta4.ProviderMaintenanceParams) |  | params defines the x/provider parameters to update. |
 
 

 

 
 <a name="akash.provider.v1beta4.MsgUpdateParamsResponse"></a>

 ### MsgUpdateParamsResponse
 MsgUpdateParamsResponse is the response type for MsgUpdateParams.

Since: akash v1.0.0

 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/provider/v1beta4/query.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/provider/v1beta4/query.proto
 

 
 <a name="akash.provider.v1beta4.QueryParamsRequest"></a>

 ### QueryParamsRequest
 QueryParamsRequest is the request type for the Query/Params RPC method.

 

 

 
 <a name="akash.provider.v1beta4.QueryParamsResponse"></a>

 ### QueryParamsResponse
 QueryParamsResponse is the response type for the Query/Params RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `params` | [ProviderMaintenanceParams](#akash.provider.v1beta4.ProviderMaintenanceParams) |  | params defines the parameters of the x/provider module. |
 
 

 

 
 <a name="akash.provider.v1beta4.QueryProviderMaintenanceRequest"></a>

 ### QueryProviderMaintenanceRequest
 QueryProviderMaintenanceRequest is the request type for the
Query/ProviderMaintenance RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | provider is the bech32 address of the provider whose maintenance record is being looked up.

Example: "akash1..." |
 | `maintenance_id` | [uint64](#uint64) |  | maintenance_id is the identifier of the maintenance record. |
 
 

 

 
 <a name="akash.provider.v1beta4.QueryProviderMaintenanceResponse"></a>

 ### QueryProviderMaintenanceResponse
 QueryProviderMaintenanceResponse is the response type for the
Query/ProviderMaintenance RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `maintenance` | [ProviderMaintenanceWithStatus](#akash.provider.v1beta4.ProviderMaintenanceWithStatus) |  | maintenance is the requested maintenance record. |
 
 

 

 
 <a name="akash.provider.v1beta4.QueryProviderMaintenancesRequest"></a>

 ### QueryProviderMaintenancesRequest
 QueryProviderMaintenancesRequest is the request type for the
Query/ProviderMaintenances RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | provider is the bech32 address of the provider whose maintenance records are being listed.

Example: "akash1..." |
 | `status_filter` | [ProviderMaintenanceStatus](#akash.provider.v1beta4.ProviderMaintenanceStatus) |  | status_filter optionally restricts the results by status. |
 | `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination is used to paginate the request. |
 
 

 

 
 <a name="akash.provider.v1beta4.QueryProviderMaintenancesResponse"></a>

 ### QueryProviderMaintenancesResponse
 QueryProviderMaintenancesResponse is the response type for the
Query/ProviderMaintenances RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `maintenance` | [ProviderMaintenanceWithStatus](#akash.provider.v1beta4.ProviderMaintenanceWithStatus) | repeated | maintenance is the list of records matching the request. |
 | `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination contains the information about response pagination. |
 
 

 

 
 <a name="akash.provider.v1beta4.QueryProviderRequest"></a>

 ### QueryProviderRequest
 QueryProviderRequest is request type for the Query/Provider RPC method

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `owner` | [string](#string) |  | Owner is the bech32 address of the account of the provider. It is a string representing a valid account address.

Example: "akash1..." |
 
 

 

 
 <a name="akash.provider.v1beta4.QueryProviderResponse"></a>

 ### QueryProviderResponse
 QueryProviderResponse is response type for the Query/Provider RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [Provider](#akash.provider.v1beta4.Provider) |  | Provider holds the representation of a provider on the network. |
 
 

 

 
 <a name="akash.provider.v1beta4.QueryProvidersRequest"></a>

 ### QueryProvidersRequest
 QueryProvidersRequest is request type for the Query/Providers RPC method

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | Pagination is used to paginate the request. |
 
 

 

 
 <a name="akash.provider.v1beta4.QueryProvidersResponse"></a>

 ### QueryProvidersResponse
 QueryProvidersResponse is response type for the Query/Providers RPC method

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `providers` | [Provider](#akash.provider.v1beta4.Provider) | repeated | Providers is a list of providers on the network. |
 | `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | Pagination contains the information about response pagination. |
 
 

 

 
 <a name="akash.provider.v1beta4.QueryRegistrationRequest"></a>

 ### QueryRegistrationRequest
 QueryRegistrationRequest is the request type for the Query/Registration RPC
method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | provider is the bech32 address of the provider whose registration record is being looked up.

Example: "akash1..." |
 
 

 

 
 <a name="akash.provider.v1beta4.QueryRegistrationResponse"></a>

 ### QueryRegistrationResponse
 QueryRegistrationResponse is the response type for the Query/Registration
RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `registration` | [ProviderRegistration](#akash.provider.v1beta4.ProviderRegistration) |  | registration is the provider registration record. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.provider.v1beta4.Query"></a>

 ### Query
 Query defines the gRPC querier service for the provider package.

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `Providers` | [QueryProvidersRequest](#akash.provider.v1beta4.QueryProvidersRequest) | [QueryProvidersResponse](#akash.provider.v1beta4.QueryProvidersResponse) | Providers queries providers | GET|/akash/provider/v1beta4/providers|
 | `Provider` | [QueryProviderRequest](#akash.provider.v1beta4.QueryProviderRequest) | [QueryProviderResponse](#akash.provider.v1beta4.QueryProviderResponse) | Provider queries provider details | GET|/akash/provider/v1beta4/providers/{owner}|
 | `ProviderMaintenance` | [QueryProviderMaintenanceRequest](#akash.provider.v1beta4.QueryProviderMaintenanceRequest) | [QueryProviderMaintenanceResponse](#akash.provider.v1beta4.QueryProviderMaintenanceResponse) | ProviderMaintenance queries a provider maintenance record. | GET|/akash/provider/v1beta4/providers/{provider}/maintenance/{maintenance_id}|
 | `ProviderMaintenances` | [QueryProviderMaintenancesRequest](#akash.provider.v1beta4.QueryProviderMaintenancesRequest) | [QueryProviderMaintenancesResponse](#akash.provider.v1beta4.QueryProviderMaintenancesResponse) | ProviderMaintenances queries provider maintenance records. | GET|/akash/provider/v1beta4/providers/{provider}/maintenance|
 | `Params` | [QueryParamsRequest](#akash.provider.v1beta4.QueryParamsRequest) | [QueryParamsResponse](#akash.provider.v1beta4.QueryParamsResponse) | Params returns the x/provider ProviderMaintenanceParams. | GET|/akash/provider/v1beta4/params|
 | `Registration` | [QueryRegistrationRequest](#akash.provider.v1beta4.QueryRegistrationRequest) | [QueryRegistrationResponse](#akash.provider.v1beta4.QueryRegistrationResponse) | Registration queries provider registration details. | GET|/akash/provider/v1beta4/providers/{provider}/registration|
 
  <!-- end services -->

 
 
 <a name="akash/provider/v1beta4/service.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/provider/v1beta4/service.proto
 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.provider.v1beta4.Msg"></a>

 ### Msg
 Msg defines the provider Msg service.

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `CreateProvider` | [MsgCreateProvider](#akash.provider.v1beta4.MsgCreateProvider) | [MsgCreateProviderResponse](#akash.provider.v1beta4.MsgCreateProviderResponse) | CreateProvider defines a method that creates a provider given the proper inputs. | |
 | `UpdateProvider` | [MsgUpdateProvider](#akash.provider.v1beta4.MsgUpdateProvider) | [MsgUpdateProviderResponse](#akash.provider.v1beta4.MsgUpdateProviderResponse) | UpdateProvider defines a method that updates a provider given the proper inputs. | |
 | `DeleteProvider` | [MsgDeleteProvider](#akash.provider.v1beta4.MsgDeleteProvider) | [MsgDeleteProviderResponse](#akash.provider.v1beta4.MsgDeleteProviderResponse) | DeleteProvider defines a method that deletes a provider given the proper inputs. | |
 | `OpenProviderMaintenance` | [MsgOpenProviderMaintenance](#akash.provider.v1beta4.MsgOpenProviderMaintenance) | [MsgOpenProviderMaintenanceResponse](#akash.provider.v1beta4.MsgOpenProviderMaintenanceResponse) | OpenProviderMaintenance opens a provider maintenance window. | |
 | `CloseProviderMaintenance` | [MsgCloseProviderMaintenance](#akash.provider.v1beta4.MsgCloseProviderMaintenance) | [MsgCloseProviderMaintenanceResponse](#akash.provider.v1beta4.MsgCloseProviderMaintenanceResponse) | CloseProviderMaintenance closes an open maintenance window. | |
 | `UpdateParams` | [MsgUpdateParams](#akash.provider.v1beta4.MsgUpdateParams) | [MsgUpdateParamsResponse](#akash.provider.v1beta4.MsgUpdateParamsResponse) | UpdateParams is a governance operation for updating the x/provider parameters.

Since: akash v1.0.0 | |
 
  <!-- end services -->

 
 
 <a name="akash/take/v1/params.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/take/v1/params.proto
 

 
 <a name="akash.take.v1.DenomTakeRate"></a>

 ### DenomTakeRate
 DenomTakeRate describes take rate for specified denom.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `denom` | [string](#string) |  | Denom is the denomination of the take rate (uakt, usdc, etc.). |
 | `rate` | [uint32](#uint32) |  | Rate is the value of the take rate. |
 
 

 

 
 <a name="akash.take.v1.Params"></a>

 ### Params
 Params defines the parameters for the x/take package.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `denom_take_rates` | [DenomTakeRate](#akash.take.v1.DenomTakeRate) | repeated | DenomTakeRates is a list of configured take rates. |
 | `default_take_rate` | [uint32](#uint32) |  | DefaultTakeRate holds the default take rate. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/take/v1/genesis.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/take/v1/genesis.proto
 

 
 <a name="akash.take.v1.GenesisState"></a>

 ### GenesisState
 GenesisState stores slice of genesis staking parameters.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `params` | [Params](#akash.take.v1.Params) |  | Params holds parameters of the genesis of staking. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/take/v1/paramsmsg.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/take/v1/paramsmsg.proto
 

 
 <a name="akash.take.v1.MsgUpdateParams"></a>

 ### MsgUpdateParams
 MsgUpdateParams is the Msg/UpdateParams request type.

Since: akash v1.0.0

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `authority` | [string](#string) |  | authority is the address of the governance account. |
 | `params` | [Params](#akash.take.v1.Params) |  | params defines the x/deployment parameters to update.

NOTE: All parameters must be supplied. |
 
 

 

 
 <a name="akash.take.v1.MsgUpdateParamsResponse"></a>

 ### MsgUpdateParamsResponse
 MsgUpdateParamsResponse defines the response structure for executing a
MsgUpdateParams message.

Since: akash v1.0.0

 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/take/v1/query.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/take/v1/query.proto
 

 
 <a name="akash.take.v1.QueryParamsRequest"></a>

 ### QueryParamsRequest
 QueryParamsRequest is the request type for the Query/Params RPC method.

 

 

 
 <a name="akash.take.v1.QueryParamsResponse"></a>

 ### QueryParamsResponse
 QueryParamsResponse is the response type for the Query/Params RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `params` | [Params](#akash.take.v1.Params) |  | params defines the parameters of the module. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.take.v1.Query"></a>

 ### Query
 Query defines the gRPC querier service of the take package.

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `Params` | [QueryParamsRequest](#akash.take.v1.QueryParamsRequest) | [QueryParamsResponse](#akash.take.v1.QueryParamsResponse) | Params returns the total set of take parameters. | GET|/akash/take/v1/params|
 
  <!-- end services -->

 
 
 <a name="akash/take/v1/service.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/take/v1/service.proto
 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.take.v1.Msg"></a>

 ### Msg
 Msg defines the take Msg service

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `UpdateParams` | [MsgUpdateParams](#akash.take.v1.MsgUpdateParams) | [MsgUpdateParamsResponse](#akash.take.v1.MsgUpdateParamsResponse) | UpdateParams defines a governance operation for updating the x/market module parameters. The authority is hard-coded to the x/gov module account.

Since: akash v1.0.0 | |
 
  <!-- end services -->

 
 
 <a name="akash/verification/v1/event.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/verification/v1/event.proto
 

 
 <a name="akash.verification.v1.EventAttestationExpired"></a>

 ### EventAttestationExpired
 EventAttestationExpired is emitted when an attestation expires.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the provider whose attestation expired. |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the auditor whose attestation expired. |
 | `tier` | [VerificationTier](#akash.verification.v1.VerificationTier) |  | Tier is the tier of the expired attestation. |
 
 

 

 
 <a name="akash.verification.v1.EventAttestationReplaced"></a>

 ### EventAttestationReplaced
 EventAttestationReplaced is emitted when a new attestation replaces an existing
attestation from the same auditor for the same provider.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the provider involved. |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the auditor involved. |
 | `old_tier` | [VerificationTier](#akash.verification.v1.VerificationTier) |  | OldTier is the tier of the previous attestation being replaced. |
 | `new_tier` | [VerificationTier](#akash.verification.v1.VerificationTier) |  | NewTier is the tier of the new attestation. |
 | `old_audit_escrow_id` | [uint64](#uint64) |  | OldAuditEscrowID is the audit escrow that authorized the previous attestation. |
 | `new_audit_escrow_id` | [uint64](#uint64) |  | NewAuditEscrowID is the audit escrow that authorized the new attestation. |
 
 

 

 
 <a name="akash.verification.v1.EventAttestationRevoked"></a>

 ### EventAttestationRevoked
 EventAttestationRevoked is emitted when an attestation is revoked by the auditor.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the provider involved. |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the auditor involved. |
 | `initiator` | [string](#string) |  | Initiator is the textual identifier of the party that initiated the revocation (e.g. "auditor", "governance"). |
 | `reason` | [AttestationRevocationReason](#akash.verification.v1.AttestationRevocationReason) |  | Reason is the typed reason recorded for the revocation. |
 
 

 

 
 <a name="akash.verification.v1.EventAttestationSubmitted"></a>

 ### EventAttestationSubmitted
 EventAttestationSubmitted is emitted when an auditor submits an attestation.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the attested provider. |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the issuing auditor. |
 | `tier` | [VerificationTier](#akash.verification.v1.VerificationTier) |  | Tier is the verification tier asserted by the attestation. |
 | `capabilities` | [CapabilityFlag](#akash.verification.v1.CapabilityFlag) | repeated | Capabilities is the set of capability flags asserted by the attestation. |
 | `expires_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | ExpiresAt is the chain timestamp at which the attestation expires. |
 | `audit_escrow_id` | [uint64](#uint64) |  | AuditEscrowID is the audit escrow identifier that authorized the attestation. |
 
 

 

 
 <a name="akash.verification.v1.EventAttestationVoided"></a>

 ### EventAttestationVoided
 EventAttestationVoided is emitted when an attestation transitions to voided state.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the provider whose attestation was voided. |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the auditor whose attestation was voided. |
 | `reason` | [VoidedReason](#akash.verification.v1.VoidedReason) |  | Reason is the typed reason recorded for the void. |
 
 

 

 
 <a name="akash.verification.v1.EventAuditEscrowOpened"></a>

 ### EventAuditEscrowOpened
 EventAuditEscrowOpened is emitted when a provider opens an audit escrow.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `audit_escrow_id` | [uint64](#uint64) |  | AuditEscrowID is the identifier of the newly opened audit escrow. |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the provider opening the escrow. |
 | `fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Fee is the audit fee escrowed. |
 | `provider_deposit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | ProviderDeposit is the deposit a provider escrows alongside the fee. |
 
 

 

 
 <a name="akash.verification.v1.EventAuditEscrowSettled"></a>

 ### EventAuditEscrowSettled
 EventAuditEscrowSettled is emitted when an audit escrow reaches final settlement.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `audit_escrow_id` | [uint64](#uint64) |  | AuditEscrowID is the identifier of the settled audit escrow. |
 | `reason` | [AuditEscrowSettlementReason](#akash.verification.v1.AuditEscrowSettlementReason) |  | Reason is the typed settlement reason recorded. |
 | `fault_attribution` | [FaultAttribution](#akash.verification.v1.FaultAttribution) |  | FaultAttribution is the typed fault attribution recorded at settlement. |
 
 

 

 
 <a name="akash.verification.v1.EventAuditorBondPosted"></a>

 ### EventAuditorBondPosted
 EventAuditorBondPosted is emitted when an auditor posts (or tops up) bond.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the bonding auditor. |
 | `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Amount is the coin amount posted as bond. |
 
 

 

 
 <a name="akash.verification.v1.EventAuditorFrozen"></a>

 ### EventAuditorFrozen
 EventAuditorFrozen is emitted when an auditor crosses the discrepancy
threshold and is frozen pending governance review.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the frozen auditor. |
 | `discrepancy_id` | [uint64](#uint64) |  | DiscrepancyID is the identifier of the discrepancy that triggered the freeze. |
 
 

 

 
 <a name="akash.verification.v1.EventAuditorLapsed"></a>

 ### EventAuditorLapsed
 EventAuditorLapsed is emitted when an auditor's renewal deadline passes
without renewal.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the lapsed auditor. |
 
 

 

 
 <a name="akash.verification.v1.EventAuditorRegistered"></a>

 ### EventAuditorRegistered
 EventAuditorRegistered is emitted when a new auditor is registered.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the newly registered auditor. |
 | `max_attestation_tier` | [VerificationTier](#akash.verification.v1.VerificationTier) |  | MaxAttestationTier is the maximum tier the auditor may attest at. |
 
 

 

 
 <a name="akash.verification.v1.EventAuditorRemoved"></a>

 ### EventAuditorRemoved
 EventAuditorRemoved is emitted when governance removes an auditor.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the removed auditor. |
 
 

 

 
 <a name="akash.verification.v1.EventAuditorRenewed"></a>

 ### EventAuditorRenewed
 EventAuditorRenewed is emitted when an auditor's registration is renewed.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the renewed auditor. |
 | `new_deadline` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | NewDeadline is the new renewal deadline assigned to the auditor. |
 
 

 

 
 <a name="akash.verification.v1.EventAuditorResigned"></a>

 ### EventAuditorResigned
 EventAuditorResigned is emitted when an auditor voluntarily resigns.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the resigning auditor. |
 
 

 

 
 <a name="akash.verification.v1.EventDepositReturnedToAuditor"></a>

 ### EventDepositReturnedToAuditor
 EventDepositReturnedToAuditor is emitted when an auditor's anti-griefing
deposit is returned.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the auditor receiving the deposit. |
 | `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Amount is the coin amount returned to the auditor. |
 
 

 

 
 <a name="akash.verification.v1.EventDepositSlashed"></a>

 ### EventDepositSlashed
 EventDepositSlashed is emitted when an auditor's anti-griefing deposit is slashed.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the auditor whose deposit was slashed. |
 | `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Amount is the coin amount slashed from the deposit. |
 
 

 

 
 <a name="akash.verification.v1.EventDiscrepancyDetected"></a>

 ### EventDiscrepancyDetected
 EventDiscrepancyDetected is emitted when the chain detects two conflicting
attestations from different auditors for the same provider.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `discrepancy_id` | [uint64](#uint64) |  | DiscrepancyID is the identifier of the newly created discrepancy record. |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the disputed provider. |
 | `auditor_a` | [string](#string) |  | AuditorA is the bech32 account address of the first auditor in the dispute. |
 | `tier_a` | [VerificationTier](#akash.verification.v1.VerificationTier) |  | TierA is the tier asserted by auditor A. |
 | `auditor_b` | [string](#string) |  | AuditorB is the bech32 account address of the second auditor in the dispute. |
 | `tier_b` | [VerificationTier](#akash.verification.v1.VerificationTier) |  | TierB is the tier asserted by auditor B. |
 
 

 

 
 <a name="akash.verification.v1.EventDiscrepancyResolved"></a>

 ### EventDiscrepancyResolved
 EventDiscrepancyResolved is emitted when a discrepancy is resolved by governance.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `discrepancy_id` | [uint64](#uint64) |  | DiscrepancyID is the identifier of the resolved discrepancy record. |
 | `vindicated_auditor` | [string](#string) |  | VindicatedAuditor is the identifier of the auditor vindicated by the resolution; carries an empty string when neither party is vindicated. |
 | `reason` | [DiscrepancyResolutionReason](#akash.verification.v1.DiscrepancyResolutionReason) |  | Reason is the typed resolution reason recorded by governance. |
 | `fault_attribution` | [FaultAttribution](#akash.verification.v1.FaultAttribution) |  | FaultAttribution is the typed fault attribution recorded at resolution. |
 
 

 

 
 <a name="akash.verification.v1.EventDiscrepancyTimedOut"></a>

 ### EventDiscrepancyTimedOut
 EventDiscrepancyTimedOut is emitted when a discrepancy's resolution window
elapses without governance action.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `discrepancy_id` | [uint64](#uint64) |  | DiscrepancyID is the identifier of the timed-out discrepancy record. |
 | `auditor_a` | [string](#string) |  | AuditorA is the bech32 account address of the first auditor in the dispute. |
 | `auditor_b` | [string](#string) |  | AuditorB is the bech32 account address of the second auditor in the dispute. |
 
 

 

 
 <a name="akash.verification.v1.EventFeeEscrowed"></a>

 ### EventFeeEscrowed
 EventFeeEscrowed is emitted when an audit fee is escrowed by a provider.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the escrowing provider. |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the consuming auditor (may be empty). |
 | `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Amount is the coin amount escrowed. |
 
 

 

 
 <a name="akash.verification.v1.EventFeeReleasedToAuditor"></a>

 ### EventFeeReleasedToAuditor
 EventFeeReleasedToAuditor is emitted when an escrowed fee is released to the auditor.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the auditor receiving the fee. |
 | `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Amount is the coin amount released to the auditor. |
 
 

 

 
 <a name="akash.verification.v1.EventFeeReturnedToProvider"></a>

 ### EventFeeReturnedToProvider
 EventFeeReturnedToProvider is emitted when an escrowed fee is returned to the provider.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the provider receiving the fee. |
 | `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Amount is the coin amount returned to the provider. |
 
 

 

 
 <a name="akash.verification.v1.EventProviderBondPosted"></a>

 ### EventProviderBondPosted
 EventProviderBondPosted is emitted when a provider posts (or tops up) bond.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the bonding provider. |
 | `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Amount is the coin amount just posted in this transaction. |
 | `total_bonded` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | TotalBonded is the resulting total bonded amount after this posting. |
 
 

 

 
 <a name="akash.verification.v1.EventProviderBondSlashed"></a>

 ### EventProviderBondSlashed
 EventProviderBondSlashed is emitted when a provider's bond is slashed.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the slashed provider. |
 | `slashed_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | SlashedAmount is the coin amount slashed in this action. |
 | `reason` | [ProviderBondSlashReason](#akash.verification.v1.ProviderBondSlashReason) |  | Reason is the typed slash reason recorded. |
 
 

 

 
 <a name="akash.verification.v1.EventProviderBondWithdrawalCompleted"></a>

 ### EventProviderBondWithdrawalCompleted
 EventProviderBondWithdrawalCompleted is emitted when an unbonding entry
matures and the amount is returned to the provider.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the provider receiving the unbonded coins. |
 | `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Amount is the coin amount returned to the provider. |
 
 

 

 
 <a name="akash.verification.v1.EventProviderBondWithdrawalInitiated"></a>

 ### EventProviderBondWithdrawalInitiated
 EventProviderBondWithdrawalInitiated is emitted when a provider begins
unbonding part or all of its bond.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the unbonding provider. |
 | `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Amount is the coin amount the provider is unbonding. |
 | `completion_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | CompletionTime is the chain timestamp at which the unbonding completes. |
 
 

 

 
 <a name="akash.verification.v1.EventSnapshotHashPosted"></a>

 ### EventSnapshotHashPosted
 EventSnapshotHashPosted is emitted when a provider posts a snapshot hash.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the snapshotting provider. |
 | `snapshot_hash` | [bytes](#bytes) |  | SnapshotHash is the opaque hash bytes posted by the provider. |
 | `compliance_deadline` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | ComplianceDeadline is the chain timestamp by which the next snapshot must be posted. |
 
 

 

 
 <a name="akash.verification.v1.EventSnapshotResumed"></a>

 ### EventSnapshotResumed
 EventSnapshotResumed is emitted when a previously suspended provider snapshot
returns to compliance.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the resumed provider. |
 
 

 

 
 <a name="akash.verification.v1.EventSnapshotSuspended"></a>

 ### EventSnapshotSuspended
 EventSnapshotSuspended is emitted when a provider's snapshot is suspended
for missing a compliance deadline.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the suspended provider. |
 
 

 

 
 <a name="akash.verification.v1.EventVerificationGraceEnded"></a>

 ### EventVerificationGraceEnded
 EventVerificationGraceEnded is emitted when a verification grace window
concludes for any reason.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `grace_record_id` | [uint64](#uint64) |  | GraceRecordID is the identifier of the closed grace record. |
 | `status` | [VerificationGraceStatus](#akash.verification.v1.VerificationGraceStatus) |  | Status is the final lifecycle status of the closed grace record. |
 
 

 

 
 <a name="akash.verification.v1.EventVerificationGraceStarted"></a>

 ### EventVerificationGraceStarted
 EventVerificationGraceStarted is emitted when a provider enters a verification
grace window.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `grace_record_id` | [uint64](#uint64) |  | GraceRecordID is the identifier of the newly opened grace record. |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the provider entering grace. |
 | `preserved_tier` | [VerificationTier](#akash.verification.v1.VerificationTier) |  | PreservedTier is the verification tier preserved for the provider during grace. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/verification/v1/state.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/verification/v1/state.proto
 

 
 <a name="akash.verification.v1.AttestationRecord"></a>

 ### AttestationRecord
 AttestationRecord captures a single attestation issued by an auditor about
a provider, including the asserted tier and capabilities, the evidence
reference, the audit fee, lifecycle and fault state, anti-griefing
deposit, and the audit-escrow ID that authorized the attestation.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the attested provider. |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the issuing auditor. |
 | `tier` | [VerificationTier](#akash.verification.v1.VerificationTier) |  | Tier is the verification tier the auditor asserts for the provider. |
 | `capabilities` | [CapabilityFlag](#akash.verification.v1.CapabilityFlag) | repeated | Capabilities is the set of capability flags the auditor asserts. |
 | `evidence_hash` | [bytes](#bytes) |  | EvidenceHash is an opaque hash referencing the auditor's evidence (off-chain). |
 | `fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Fee is the audit fee paid by the provider for this attestation. |
 | `fee_status` | [FeeStatus](#akash.verification.v1.FeeStatus) |  | FeeStatus is the lifecycle status of the fee held against this attestation. |
 | `created_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | CreatedAt is the chain timestamp at which the attestation was created. |
 | `expires_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | ExpiresAt is the chain timestamp at which the attestation expires. |
 | `status` | [AttestationStatus](#akash.verification.v1.AttestationStatus) |  | Status is the lifecycle status of the attestation. |
 | `voided_reason` | [VoidedReason](#akash.verification.v1.VoidedReason) |  | VoidedReason carries the typed reason when status == Voided. |
 | `deposit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Deposit is the anti-griefing deposit posted by the auditor. |
 | `deposit_status` | [DepositStatus](#akash.verification.v1.DepositStatus) |  | DepositStatus is the lifecycle status of the auditor's deposit. |
 | `audit_escrow_id` | [uint64](#uint64) |  | AuditEscrowID is the identifier of the audit escrow that authorized this attestation. |
 | `fault_attribution` | [FaultAttribution](#akash.verification.v1.FaultAttribution) |  | FaultAttribution carries the fault attribution recorded after settlement or governance resolution involving this attestation. |
 
 

 

 
 <a name="akash.verification.v1.AuditEscrowRecord"></a>

 ### AuditEscrowRecord
 AuditEscrowRecord captures an audit escrow opened by a provider to fund a
pending attestation, including the requested tier and capabilities, the
audit fee and provider deposit, lifecycle status, lifecycle timestamps,
metadata reference, and final settlement attribution.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [uint64](#uint64) |  | ID is the unique identifier of the audit escrow. |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the provider that opened the escrow. |
 | `consumed_by_auditor` | [string](#string) |  | ConsumedByAuditor is the bech32 address of the auditor that consumed the escrow when filing an attestation; empty until consumption. |
 | `requested_tier` | [VerificationTier](#akash.verification.v1.VerificationTier) |  | RequestedTier is the verification tier the provider is requesting. |
 | `requested_capabilities` | [CapabilityFlag](#akash.verification.v1.CapabilityFlag) | repeated | RequestedCapabilities is the set of capability flags the provider is requesting. |
 | `fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Fee is the audit fee escrowed by the provider. |
 | `fee_status` | [FeeStatus](#akash.verification.v1.FeeStatus) |  | FeeStatus is the lifecycle status of the escrowed fee. |
 | `provider_deposit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | ProviderDeposit is the deposit a provider escrows alongside the fee. |
 | `provider_deposit_status` | [ProviderDepositStatus](#akash.verification.v1.ProviderDepositStatus) |  | ProviderDepositStatus is the lifecycle status of the provider deposit. |
 | `status` | [AuditEscrowStatus](#akash.verification.v1.AuditEscrowStatus) |  | Status is the lifecycle status of the audit escrow. |
 | `opened_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | OpenedAt is the chain timestamp at which the escrow was opened. |
 | `consumed_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | ConsumedAt is the chain timestamp at which the escrow was consumed; nil until consumption. |
 | `expires_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | ExpiresAt is the chain timestamp at which an unconsumed escrow expires. |
 | `metadata_hash` | [bytes](#bytes) |  | MetadataHash is an opaque hash referencing escrow metadata (off-chain). |
 | `settlement_reason` | [AuditEscrowSettlementReason](#akash.verification.v1.AuditEscrowSettlementReason) |  | SettlementReason is the typed reason recorded when the escrow is settled. |
 | `fault_attribution` | [FaultAttribution](#akash.verification.v1.FaultAttribution) |  | FaultAttribution is the fault attribution recorded at settlement. |
 
 

 

 
 <a name="akash.verification.v1.AuditorRecord"></a>

 ### AuditorRecord
 AuditorRecord captures the on-chain state of a registered auditor including
status, maximum attestation tier, bonded amount, registration and renewal
timestamps, discrepancy counter, and any pending bond unbonding completion.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `address` | [string](#string) |  | Address is the bech32 account address of the auditor. |
 | `status` | [AuditorStatus](#akash.verification.v1.AuditorStatus) |  | Status is the lifecycle status of the auditor record. |
 | `max_attestation_tier` | [VerificationTier](#akash.verification.v1.VerificationTier) |  | MaxAttestationTier is the highest verification tier this auditor is permitted to attest once the required bond has been posted. |
 | `bond_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | BondAmount is the coin amount currently bonded by the auditor. |
 | `bond_status` | [BondStatus](#akash.verification.v1.BondStatus) |  | BondStatus is the current status of the auditor bond. |
 | `metadata_hash` | [bytes](#bytes) |  | MetadataHash is an opaque hash referencing auditor metadata (off-chain). |
 | `registered_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | RegisteredAt is the chain timestamp at which this auditor was registered. |
 | `renewal_deadline` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | RenewalDeadline is the chain timestamp by which the auditor must renew. |
 | `discrepancy_count` | [uint64](#uint64) |  | DiscrepancyCount is the running count of discrepancies attributed to this auditor; used by the keeper to enforce the discrepancy threshold. |
 | `bond_unbonding_completion_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | BondUnbondingCompletionTime is the time at which a pending bond unbonding completes; nil when no unbonding is in progress. |
 
 

 

 
 <a name="akash.verification.v1.DiscrepancyEvent"></a>

 ### DiscrepancyEvent
 DiscrepancyEvent captures a disagreement between two auditors over a
provider, including both attestations' asserted tiers, the resolution
state, pointers to the governance proposal and grace record (if any),
and the final resolution attribution recorded at settlement.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [uint64](#uint64) |  | ID is the unique identifier of the discrepancy event. |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the provider in dispute. |
 | `auditor_a` | [string](#string) |  | AuditorA is the bech32 account address of the first auditor. |
 | `auditor_a_tier` | [VerificationTier](#akash.verification.v1.VerificationTier) |  | AuditorATier is the tier asserted by auditor A. |
 | `auditor_b` | [string](#string) |  | AuditorB is the bech32 account address of the second auditor. |
 | `auditor_b_tier` | [VerificationTier](#akash.verification.v1.VerificationTier) |  | AuditorBTier is the tier asserted by auditor B. |
 | `timestamp` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Timestamp is the chain time at which the discrepancy was detected. |
 | `resolution_status` | [DiscrepancyStatus](#akash.verification.v1.DiscrepancyStatus) |  | ResolutionStatus is the lifecycle status of the resolution flow. |
 | `resolution_proposal_id` | [uint64](#uint64) |  | ResolutionProposalID is the governance proposal that resolved the discrepancy; zero when no proposal has been filed. |
 | `grace_record_id` | [uint64](#uint64) |  | GraceRecordID is the provider verification grace record opened in response to this discrepancy; zero when no grace record exists. |
 | `resolution_reason` | [DiscrepancyResolutionReason](#akash.verification.v1.DiscrepancyResolutionReason) |  | ResolutionReason is the typed reason recorded at resolution. |
 | `fault_attribution` | [FaultAttribution](#akash.verification.v1.FaultAttribution) |  | FaultAttribution is the fault attribution recorded at resolution. |
 | `resolution_evidence_hash` | [bytes](#bytes) |  | ResolutionEvidenceHash is an opaque hash referencing the evidence used at resolution (off-chain). |
 
 

 

 
 <a name="akash.verification.v1.ProviderBondRecord"></a>

 ### ProviderBondRecord
 ProviderBondRecord holds a provider's resource-scaled verification bond,
any unbonding entries currently in progress, and slash history.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the bonded provider. |
 | `bonded_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | BondedAmount is the coin amount currently bonded by the provider. |
 | `unbonding_entries` | [UnbondingEntry](#akash.verification.v1.UnbondingEntry) | repeated | UnbondingEntries is the list of pending unbondings against this bond. |
 | `slashed` | [bool](#bool) |  | Slashed indicates whether this provider bond has ever been slashed. |
 | `last_slash_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | LastSlashTime is the chain timestamp of the most recent slash; nil if the bond has never been slashed. |
 
 

 

 
 <a name="akash.verification.v1.ProviderSnapshotRecord"></a>

 ### ProviderSnapshotRecord
 ProviderSnapshotRecord is the on-chain record of a provider's most recent
resource snapshot, including the snapshot hash, resource summary, posting
and snapshot timestamps, the compliance deadline, and suspension state.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the snapshotting provider. |
 | `snapshot_hash` | [bytes](#bytes) |  | SnapshotHash is the opaque hash of the snapshot payload (off-chain). |
 | `resource_summary` | [ResourceSummary](#akash.verification.v1.ResourceSummary) |  | ResourceSummary is the inline resource counts the provider claims at snapshot time. |
 | `posted_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | PostedAt is the chain timestamp at which the snapshot record was posted on-chain. |
 | `snapshot_timestamp` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | SnapshotTimestamp is the timestamp embedded by the provider in the snapshot itself. |
 | `compliance_deadline` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | ComplianceDeadline is the chain timestamp by which the next snapshot must be posted to remain compliant with the snapshot cadence. |
 | `suspended` | [bool](#bool) |  | Suspended indicates whether this provider snapshot has been suspended for missing the compliance deadline. |
 
 

 

 
 <a name="akash.verification.v1.ProviderVerificationGraceRecord"></a>

 ### ProviderVerificationGraceRecord
 ProviderVerificationGraceRecord preserves a provider's verification tier
across a discrepancy by tracking the source discrepancies, preserved tier,
and the grace window during which the provider may continue to operate.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [uint64](#uint64) |  | ID is the unique identifier of the grace record. |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the provider in grace. |
 | `preserved_tier` | [VerificationTier](#akash.verification.v1.VerificationTier) |  | PreservedTier is the verification tier preserved for the provider during grace. |
 | `source_discrepancy_ids` | [uint64](#uint64) | repeated | SourceDiscrepancyIDs is the list of discrepancy IDs that opened this grace. |
 | `started_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | StartedAt is the chain timestamp at which the grace window began. |
 | `expires_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | ExpiresAt is the chain timestamp at which the grace window expires. |
 | `status` | [VerificationGraceStatus](#akash.verification.v1.VerificationGraceStatus) |  | Status is the lifecycle status of the grace window. |
 
 

 

 
 <a name="akash.verification.v1.ResourceSummary"></a>

 ### ResourceSummary
 ResourceSummary captures the resource counts a provider claims at snapshot
time. Used by the chain to scale the provider bond requirement and as
evidence in snapshot records. Software identity fields are evidence-only and
do not drive on-chain enforcement.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `total_gpus` | [uint32](#uint32) |  | TotalGPUs is the total number of GPUs the provider claims. |
 | `total_vcpus` | [uint32](#uint32) |  | TotalVCPUs is the total number of virtual CPUs the provider claims. |
 | `total_memory_mb` | [uint64](#uint64) |  | TotalMemoryMB is the total memory in megabytes the provider claims. |
 | `total_storage_mb` | [uint64](#uint64) |  | TotalStorageMB is the total storage in megabytes the provider claims. |
 | `active_leases` | [uint32](#uint32) |  | ActiveLeases is the number of leases currently active on the provider. |
 | `software_version` | [string](#string) |  | SoftwareVersion is the provider software version string kept for compatibility. |
 | `software_signature` | [bytes](#bytes) |  | SoftwareSignature is the provider software signature kept for compatibility. |
 | `software_identity` | [SoftwareIdentity](#akash.verification.v1.SoftwareIdentity) |  | SoftwareIdentity carries structured release artifact metadata. |
 
 

 

 
 <a name="akash.verification.v1.SoftwareIdentity"></a>

 ### SoftwareIdentity
 SoftwareIdentity carries release artifact identity and signature metadata.
Providers post these fields as evidence; auditors verify them off-chain
against the published Akash release key.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `version` | [string](#string) |  | Version is the provider or inventory software version string. |
 | `artifact_ref` | [string](#string) |  | ArtifactRef identifies the release artifact whose digest/signature is reported. |
 | `digest_algorithm` | [string](#string) |  | DigestAlgorithm identifies the digest algorithm, e.g. sha3-256. |
 | `digest` | [bytes](#bytes) |  | Digest is the release artifact digest bytes. |
 | `signature_type` | [string](#string) |  | SignatureType identifies the signature format, e.g. cosign. |
 | `signature` | [bytes](#bytes) |  | Signature is the detached signature bytes when carried inline. |
 | `signature_ref` | [string](#string) |  | SignatureRef identifies an external signature or bundle. |
 | `public_key_ref` | [string](#string) |  | PublicKeyRef identifies the published release public key. |
 
 

 

 
 <a name="akash.verification.v1.UnbondingEntry"></a>

 ### UnbondingEntry
 UnbondingEntry represents a single bond-unbonding amount and completion
time queued against a provider or auditor bond.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Amount is the coin amount queued for unbonding. |
 | `completion_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | CompletionTime is the chain timestamp at which this entry completes unbonding and the amount becomes spendable. |
 
 

 

 
 <a name="akash.verification.v1.VerificationStoreRecord"></a>

 ### VerificationStoreRecord
 VerificationStoreRecord is a generic any-typed wrapper used by genesis
import/export to carry sub-record payloads addressed by their proto type URL.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `type_url` | [string](#string) |  | TypeURL is the fully-qualified proto type URL of the wrapped record. |
 | `value` | [bytes](#bytes) |  | Value is the serialized bytes of the wrapped record. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/verification/v1/params.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/verification/v1/params.proto
 

 
 <a name="akash.verification.v1.Params"></a>

 ### Params
 Params defines the on-chain parameters for the verification module.
Includes per-tier auditor bond and TTL settings, per-tier minimum audit fees,
resource-scaled provider bond amounts, history requirements for higher tiers,
EndBlocker budget caps, escrow timing, anti-griefing deposits, the feature
flag controlling module activation, and per-tier contact-response targets.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `bond_l1` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | BondL1 is the auditor bond required to attest at L1 (Identified). |
 | `bond_l2` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | BondL2 is the auditor bond required to attest at L2 (Verified). |
 | `bond_l3` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | BondL3 is the auditor bond required to attest at L3 (Established). |
 | `bond_l4` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | BondL4 is the auditor bond required to attest at L4 (Trusted). |
 | `ttl_l1` | [google.protobuf.Duration](#google.protobuf.Duration) |  | TTLL1 is the time-to-live of an L1 attestation before it expires. |
 | `ttl_l2` | [google.protobuf.Duration](#google.protobuf.Duration) |  | TTLL2 is the time-to-live of an L2 attestation before it expires. |
 | `ttl_l3` | [google.protobuf.Duration](#google.protobuf.Duration) |  | TTLL3 is the time-to-live of an L3 attestation before it expires. |
 | `ttl_l4` | [google.protobuf.Duration](#google.protobuf.Duration) |  | TTLL4 is the time-to-live of an L4 attestation before it expires. |
 | `min_fee_l1` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | MinFeeL1 is the minimum audit fee a provider must escrow for an L1 audit. |
 | `min_fee_l2` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | MinFeeL2 is the minimum audit fee a provider must escrow for an L2 audit. |
 | `min_fee_l3` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | MinFeeL3 is the minimum audit fee a provider must escrow for an L3 audit. |
 | `min_fee_l4` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | MinFeeL4 is the minimum audit fee a provider must escrow for an L4 audit. |
 | `discrepancy_threshold` | [uint32](#uint32) |  | DiscrepancyThreshold is the number of discrepancies that triggers auditor freeze. |
 | `auditor_unbonding_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  | AuditorUnbondingPeriod is the unbonding duration applied to an auditor bond upon withdrawal initiation. |
 | `renewal_period_l1` | [google.protobuf.Duration](#google.protobuf.Duration) |  | RenewalPeriodL1 is the auditor renewal period for L1 attestations. |
 | `renewal_period_l2` | [google.protobuf.Duration](#google.protobuf.Duration) |  | RenewalPeriodL2 is the auditor renewal period for L2 attestations. |
 | `renewal_period_l3` | [google.protobuf.Duration](#google.protobuf.Duration) |  | RenewalPeriodL3 is the auditor renewal period for L3 attestations. |
 | `renewal_period_l4` | [google.protobuf.Duration](#google.protobuf.Duration) |  | RenewalPeriodL4 is the auditor renewal period for L4 attestations. |
 | `snapshot_hash_interval` | [google.protobuf.Duration](#google.protobuf.Duration) |  | SnapshotHashInterval is the cadence at which providers must post a snapshot hash. |
 | `max_snapshot_age` | [google.protobuf.Duration](#google.protobuf.Duration) |  | MaxSnapshotAge is the maximum age a provider snapshot may have before being stale. |
 | `bond_gpu_l2` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | BondGpuL2 is the per-GPU resource-scaled provider bond at L2. |
 | `bond_gpu_l3` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | BondGpuL3 is the per-GPU resource-scaled provider bond at L3. |
 | `bond_gpu_l4` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | BondGpuL4 is the per-GPU resource-scaled provider bond at L4. |
 | `bond_vcpu_l2` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | BondVcpuL2 is the per-vCPU resource-scaled provider bond at L2. |
 | `bond_vcpu_l3` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | BondVcpuL3 is the per-vCPU resource-scaled provider bond at L3. |
 | `bond_vcpu_l4` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | BondVcpuL4 is the per-vCPU resource-scaled provider bond at L4. |
 | `bond_mem_gb_l2` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | BondMemGbL2 is the per-GB-memory resource-scaled provider bond at L2. |
 | `bond_mem_gb_l3` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | BondMemGbL3 is the per-GB-memory resource-scaled provider bond at L3. |
 | `bond_mem_gb_l4` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | BondMemGbL4 is the per-GB-memory resource-scaled provider bond at L4. |
 | `bond_storage_tb_l2` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | BondStorageTbL2 is the per-TB-storage resource-scaled provider bond at L2. |
 | `bond_storage_tb_l3` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | BondStorageTbL3 is the per-TB-storage resource-scaled provider bond at L3. |
 | `bond_storage_tb_l4` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | BondStorageTbL4 is the per-TB-storage resource-scaled provider bond at L4. |
 | `provider_bond_unbonding_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  | ProviderBondUnbondingPeriod is the unbonding duration applied to a provider bond upon withdrawal initiation. |
 | `min_age_l2` | [google.protobuf.Duration](#google.protobuf.Duration) |  | MinAgeL2 is the minimum provider age required to attain L2. |
 | `min_age_l3` | [google.protobuf.Duration](#google.protobuf.Duration) |  | MinAgeL3 is the minimum provider age required to attain L3. |
 | `min_age_l4` | [google.protobuf.Duration](#google.protobuf.Duration) |  | MinAgeL4 is the minimum provider age required to attain L4. |
 | `min_lease_completion_bps_l3` | [uint32](#uint32) |  | MinLeaseCompletionBpsL3 is the minimum lease-completion rate (basis points) for L3. |
 | `min_lease_completion_bps_l4` | [uint32](#uint32) |  | MinLeaseCompletionBpsL4 is the minimum lease-completion rate (basis points) for L4. |
 | `clean_history_window_l3` | [google.protobuf.Duration](#google.protobuf.Duration) |  | CleanHistoryWindowL3 is the clean-history window required to attain L3. |
 | `clean_history_window_l4` | [google.protobuf.Duration](#google.protobuf.Duration) |  | CleanHistoryWindowL4 is the clean-history window required to attain L4. |
 | `min_l3_duration_for_l4` | [google.protobuf.Duration](#google.protobuf.Duration) |  | MinL3DurationForL4 is the minimum continuous L3 duration required before promotion to L4. |
 | `min_leases_for_completion_rate` | [uint32](#uint32) |  | MinLeasesForCompletionRate is the minimum number of leases required before the lease-completion ratio is considered meaningful. |
 | `max_endblocker_attestation_expiries` | [uint32](#uint32) |  | MaxEndblockerAttestationExpiries caps the number of attestation expiries processed per EndBlocker. |
 | `max_endblocker_snapshot_suspensions` | [uint32](#uint32) |  | MaxEndblockerSnapshotSuspensions caps the number of snapshot suspensions processed per EndBlocker. |
 | `max_endblocker_unbonding_completions` | [uint32](#uint32) |  | MaxEndblockerUnbondingCompletions caps the number of unbonding completions processed per EndBlocker. |
 | `max_endblocker_discrepancy_timeouts` | [uint32](#uint32) |  | MaxEndblockerDiscrepancyTimeouts caps the number of discrepancy timeouts processed per EndBlocker. |
 | `max_endblocker_audit_escrow_expiries` | [uint32](#uint32) |  | MaxEndblockerAuditEscrowExpiries caps the number of audit escrow expiries processed per EndBlocker. |
 | `max_endblocker_grace_expiries` | [uint32](#uint32) |  | MaxEndblockerGraceExpiries caps the number of verification-grace expiries processed per EndBlocker. |
 | `discrepancy_resolution_timeout` | [google.protobuf.Duration](#google.protobuf.Duration) |  | DiscrepancyResolutionTimeout is the maximum time a discrepancy may remain pending before timing out. |
 | `attestation_deposit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | AttestationDeposit is the anti-griefing deposit required when submitting an attestation. |
 | `discrepancy_grace_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  | DiscrepancyGracePeriod is the grace window applied after a discrepancy is opened. |
 | `provider_audit_deposit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | ProviderAuditDeposit is the deposit a provider must escrow when opening an audit escrow. |
 | `verification_module_active` | [bool](#bool) |  | VerificationModuleActive is the feature flag controlling whether the verification module enforces filtering and accepts new transactions. |
 | `contact_response_critical_l1` | [google.protobuf.Duration](#google.protobuf.Duration) |  | ContactResponseCriticalL1 is the maximum critical-contact response time at L1. |
 | `contact_response_critical_l2` | [google.protobuf.Duration](#google.protobuf.Duration) |  | ContactResponseCriticalL2 is the maximum critical-contact response time at L2. |
 | `contact_response_critical_l3` | [google.protobuf.Duration](#google.protobuf.Duration) |  | ContactResponseCriticalL3 is the maximum critical-contact response time at L3. |
 | `contact_response_critical_l4` | [google.protobuf.Duration](#google.protobuf.Duration) |  | ContactResponseCriticalL4 is the maximum critical-contact response time at L4. |
 | `contact_response_standard_l1` | [google.protobuf.Duration](#google.protobuf.Duration) |  | ContactResponseStandardL1 is the maximum standard-contact response time at L1. |
 | `contact_response_standard_l2` | [google.protobuf.Duration](#google.protobuf.Duration) |  | ContactResponseStandardL2 is the maximum standard-contact response time at L2. |
 | `contact_response_standard_l3` | [google.protobuf.Duration](#google.protobuf.Duration) |  | ContactResponseStandardL3 is the maximum standard-contact response time at L3. |
 | `contact_response_standard_l4` | [google.protobuf.Duration](#google.protobuf.Duration) |  | ContactResponseStandardL4 is the maximum standard-contact response time at L4. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/verification/v1/genesis.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/verification/v1/genesis.proto
 

 
 <a name="akash.verification.v1.GenesisState"></a>

 ### GenesisState
 GenesisState defines the genesis state of the verification module. It
contains the on-chain parameters and the full set of records the module
is responsible for: auditors, attestations, discrepancies (and the next ID
counter), provider bonds, provider snapshot records, audit escrows (and
the next ID counter), and provider verification grace records (and the
next ID counter).

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `params` | [Params](#akash.verification.v1.Params) |  | Params holds the on-chain parameters for the verification module. |
 | `auditors` | [AuditorRecord](#akash.verification.v1.AuditorRecord) | repeated | Auditors is the list of registered auditor records. |
 | `attestations` | [AttestationRecord](#akash.verification.v1.AttestationRecord) | repeated | Attestations is the list of attestation records. |
 | `discrepancies` | [DiscrepancyEvent](#akash.verification.v1.DiscrepancyEvent) | repeated | Discrepancies is the list of discrepancy events. |
 | `provider_bonds` | [ProviderBondRecord](#akash.verification.v1.ProviderBondRecord) | repeated | ProviderBonds is the list of provider bond records. |
 | `provider_snapshots` | [ProviderSnapshotRecord](#akash.verification.v1.ProviderSnapshotRecord) | repeated | ProviderSnapshots is the list of provider snapshot records. |
 | `next_discrepancy_id` | [uint64](#uint64) |  | NextDiscrepancyID is the monotonically increasing ID assigned to the next discrepancy event. |
 | `audit_escrows` | [AuditEscrowRecord](#akash.verification.v1.AuditEscrowRecord) | repeated | AuditEscrows is the list of audit escrow records. |
 | `next_audit_escrow_id` | [uint64](#uint64) |  | NextAuditEscrowID is the monotonically increasing ID assigned to the next audit escrow record. |
 | `verification_graces` | [ProviderVerificationGraceRecord](#akash.verification.v1.ProviderVerificationGraceRecord) | repeated | VerificationGraces is the list of provider verification grace records. |
 | `next_grace_record_id` | [uint64](#uint64) |  | NextGraceRecordID is the monotonically increasing ID assigned to the next provider verification grace record. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/verification/v1/msg.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/verification/v1/msg.proto
 

 
 <a name="akash.verification.v1.MsgCancelAuditEscrow"></a>

 ### MsgCancelAuditEscrow
 MsgCancelAuditEscrow is the provider-signed message used to cancel an open,
unconsumed audit escrow before its expiry. Cancellation returns the audit
fee and the provider deposit to the provider.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the provider that opened the escrow and the signer of this message. |
 | `audit_escrow_id` | [uint64](#uint64) |  | AuditEscrowID is the identifier of the audit escrow being cancelled. |
 
 

 

 
 <a name="akash.verification.v1.MsgCancelAuditEscrowResponse"></a>

 ### MsgCancelAuditEscrowResponse
 MsgCancelAuditEscrowResponse is the response type for Msg/CancelAuditEscrow.

 

 

 
 <a name="akash.verification.v1.MsgOpenAuditEscrow"></a>

 ### MsgOpenAuditEscrow
 MsgOpenAuditEscrow is the provider-signed message used to open an audit
escrow that funds a pending attestation. The provider escrows both an audit
fee and a provider deposit, sets the requested tier and capabilities, and
declares an expiry after which the escrow may be settled if no auditor has
consumed it. The message does not name an auditor; the first valid
attestation against this escrow consumes it.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the provider opening the escrow and the signer of this message. |
 | `requested_tier` | [VerificationTier](#akash.verification.v1.VerificationTier) |  | RequestedTier is the verification tier the provider is requesting. |
 | `requested_capabilities` | [CapabilityFlag](#akash.verification.v1.CapabilityFlag) | repeated | RequestedCapabilities is the set of capability flags the provider is requesting that an auditor verify. |
 | `fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Fee is the audit fee the provider is escrowing for the eventual auditor. |
 | `provider_deposit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | ProviderDeposit is the deposit the provider escrows alongside the fee. It can only be returned to the provider or slashed to the community pool via governance-authorized settlement. |
 | `expires_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | ExpiresAt is the chain timestamp at which an unconsumed escrow expires and may be settled. |
 | `metadata_hash` | [bytes](#bytes) |  | MetadataHash is an opaque hash referencing escrow metadata (off-chain). |
 
 

 

 
 <a name="akash.verification.v1.MsgOpenAuditEscrowResponse"></a>

 ### MsgOpenAuditEscrowResponse
 MsgOpenAuditEscrowResponse is the response type for Msg/OpenAuditEscrow.
It carries the identifier assigned to the newly opened audit escrow.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `audit_escrow_id` | [uint64](#uint64) |  | AuditEscrowID is the unique identifier assigned to the opened audit escrow. |
 
 

 

 
 <a name="akash.verification.v1.MsgPostAuditorBond"></a>

 ### MsgPostAuditorBond
 MsgPostAuditorBond is the auditor-signed message used to post (or top up)
the auditor's verification bond. The bond is held by the verification module
account and determines the maximum attestation tier the auditor may use.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the bond-posting auditor and the signer of this message. |
 | `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Amount is the coin amount being posted as auditor bond. |
 
 

 

 
 <a name="akash.verification.v1.MsgPostAuditorBondResponse"></a>

 ### MsgPostAuditorBondResponse
 MsgPostAuditorBondResponse is the response type for Msg/PostAuditorBond.

 

 

 
 <a name="akash.verification.v1.MsgPostProviderBond"></a>

 ### MsgPostProviderBond
 MsgPostProviderBond is the provider-signed message used to post (or top up)
the provider's resource-scaled verification bond. The bond is held by the
verification module account.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the bond-posting provider and the signer of this message. |
 | `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Amount is the coin amount being posted as provider bond. |
 
 

 

 
 <a name="akash.verification.v1.MsgPostProviderBondResponse"></a>

 ### MsgPostProviderBondResponse
 MsgPostProviderBondResponse is the response type for Msg/PostProviderBond.

 

 

 
 <a name="akash.verification.v1.MsgPostSnapshotHash"></a>

 ### MsgPostSnapshotHash
 MsgPostSnapshotHash is the provider-signed message used to post the
provider's most recent resource snapshot hash, the inline resource summary
claimed by the provider, and the provider-side snapshot timestamp.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the snapshotting provider and the signer of this message. |
 | `snapshot_hash` | [bytes](#bytes) |  | SnapshotHash is the opaque hash of the snapshot payload (off-chain). |
 | `resource_summary` | [ResourceSummary](#akash.verification.v1.ResourceSummary) |  | ResourceSummary is the inline resource counts the provider claims at snapshot time; used to scale the provider bond requirement. |
 | `snapshot_timestamp` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | SnapshotTimestamp is the timestamp embedded by the provider in the snapshot itself. |
 
 

 

 
 <a name="akash.verification.v1.MsgPostSnapshotHashResponse"></a>

 ### MsgPostSnapshotHashResponse
 MsgPostSnapshotHashResponse is the response type for Msg/PostSnapshotHash.

 

 

 
 <a name="akash.verification.v1.MsgRegisterAuditor"></a>

 ### MsgRegisterAuditor
 MsgRegisterAuditor is the governance-authority-signed message used to
register a new auditor with a maximum attestation tier. There is no
auditor self-registration path in v1; only governance may register
auditors.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `authority` | [string](#string) |  | Authority is the bech32 address of the governance account authorized to register auditors and the signer of this message. |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the auditor being registered. |
 | `max_attestation_tier` | [VerificationTier](#akash.verification.v1.VerificationTier) |  | MaxAttestationTier is the highest verification tier this auditor is permitted to attest, set by governance at registration. |
 | `metadata_hash` | [bytes](#bytes) |  | MetadataHash is an opaque hash referencing auditor metadata (off-chain). |
 
 

 

 
 <a name="akash.verification.v1.MsgRegisterAuditorResponse"></a>

 ### MsgRegisterAuditorResponse
 MsgRegisterAuditorResponse is the response type for Msg/RegisterAuditor.

 

 

 
 <a name="akash.verification.v1.MsgRemoveAttestation"></a>

 ### MsgRemoveAttestation
 MsgRemoveAttestation is the provider-signed message used to voluntarily
remove an attestation associated with the provider. No reason or evidence
is required; the attestation transitions to the Removed status.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the provider requesting removal and the signer of this message. |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the issuing auditor of the attestation being removed. |
 
 

 

 
 <a name="akash.verification.v1.MsgRemoveAttestationResponse"></a>

 ### MsgRemoveAttestationResponse
 MsgRemoveAttestationResponse is the response type for Msg/RemoveAttestation.

 

 

 
 <a name="akash.verification.v1.MsgRemoveAuditor"></a>

 ### MsgRemoveAuditor
 MsgRemoveAuditor is the governance-authority-signed message used to remove
an auditor from the active set. Removal moves the auditor record to the
Removed status.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `authority` | [string](#string) |  | Authority is the bech32 address of the governance account authorized to remove auditors and the signer of this message. |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the auditor being removed. |
 
 

 

 
 <a name="akash.verification.v1.MsgRemoveAuditorResponse"></a>

 ### MsgRemoveAuditorResponse
 MsgRemoveAuditorResponse is the response type for Msg/RemoveAuditor.

 

 

 
 <a name="akash.verification.v1.MsgRenewAuditor"></a>

 ### MsgRenewAuditor
 MsgRenewAuditor is the governance-authority-signed message used to renew an
auditor's registration and reset the renewal deadline.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `authority` | [string](#string) |  | Authority is the bech32 address of the governance account authorized to renew auditors and the signer of this message. |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the auditor whose registration is being renewed. |
 
 

 

 
 <a name="akash.verification.v1.MsgRenewAuditorResponse"></a>

 ### MsgRenewAuditorResponse
 MsgRenewAuditorResponse is the response type for Msg/RenewAuditor.

 

 

 
 <a name="akash.verification.v1.MsgResignAuditor"></a>

 ### MsgResignAuditor
 MsgResignAuditor is the auditor-signed message used by an auditor to
voluntarily exit the auditor role. Resignation moves the auditor record to
the Resigned status and begins unbonding of any posted auditor bond.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the resigning auditor and the signer of this message. |
 
 

 

 
 <a name="akash.verification.v1.MsgResignAuditorResponse"></a>

 ### MsgResignAuditorResponse
 MsgResignAuditorResponse is the response type for Msg/ResignAuditor.

 

 

 
 <a name="akash.verification.v1.MsgResolveDiscrepancy"></a>

 ### MsgResolveDiscrepancy
 MsgResolveDiscrepancy is the governance-authority-signed message used to
resolve a pending discrepancy between two auditors over the same provider.
The authority names the vindicated auditor (if any), optionally slashes
either or both auditor bonds, and records the typed resolution reason,
fault attribution, and evidence reference.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `authority` | [string](#string) |  | Authority is the bech32 address of the governance account authorized to resolve discrepancies and the signer of this message. |
 | `discrepancy_id` | [uint64](#uint64) |  | DiscrepancyID is the identifier of the discrepancy being resolved. |
 | `vindicated_auditor` | [string](#string) |  | VindicatedAuditor is the bech32 account address of the auditor whose attestation is upheld by the resolution; may be empty when no auditor is vindicated. |
 | `slash_auditor_a` | [bool](#bool) |  | SlashAuditorA indicates whether auditor A's bond should be slashed as part of this resolution. |
 | `slash_auditor_b` | [bool](#bool) |  | SlashAuditorB indicates whether auditor B's bond should be slashed as part of this resolution. |
 | `reason` | [DiscrepancyResolutionReason](#akash.verification.v1.DiscrepancyResolutionReason) |  | Reason is the typed discrepancy resolution reason recorded at resolution. |
 | `fault_attribution` | [FaultAttribution](#akash.verification.v1.FaultAttribution) |  | FaultAttribution is the fault attribution recorded at resolution. |
 | `evidence_hash` | [bytes](#bytes) |  | EvidenceHash is an opaque hash referencing the evidence considered at resolution (off-chain). |
 
 

 

 
 <a name="akash.verification.v1.MsgResolveDiscrepancyResponse"></a>

 ### MsgResolveDiscrepancyResponse
 MsgResolveDiscrepancyResponse is the response type for
Msg/ResolveDiscrepancy.

 

 

 
 <a name="akash.verification.v1.MsgRevokeAllProviderAttestations"></a>

 ### MsgRevokeAllProviderAttestations
 MsgRevokeAllProviderAttestations is the governance-authority-signed message
used to revoke every active attestation for a single provider, with a
typed reason, fault attribution, and evidence reference.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `authority` | [string](#string) |  | Authority is the bech32 address of the governance account authorized to revoke attestations and the signer of this message. |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the provider whose attestations are being revoked en masse. |
 | `reason` | [GovernanceAttestationReason](#akash.verification.v1.GovernanceAttestationReason) |  | Reason is the typed governance attestation reason recorded against the revocations. |
 | `fault_attribution` | [FaultAttribution](#akash.verification.v1.FaultAttribution) |  | FaultAttribution is the fault attribution recorded against the revocations. |
 | `evidence_hash` | [bytes](#bytes) |  | EvidenceHash is an opaque hash referencing the evidence backing the revocations (off-chain). |
 
 

 

 
 <a name="akash.verification.v1.MsgRevokeAllProviderAttestationsResponse"></a>

 ### MsgRevokeAllProviderAttestationsResponse
 MsgRevokeAllProviderAttestationsResponse is the response type for
Msg/RevokeAllProviderAttestations.

 

 

 
 <a name="akash.verification.v1.MsgRevokeAttestation"></a>

 ### MsgRevokeAttestation
 MsgRevokeAttestation is the auditor-signed message used to revoke a
previously submitted attestation for a typed reason, with an evidence
reference. Revocation transitions the attestation to the Revoked status.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the attested provider. |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the revoking auditor and the signer of this message. |
 | `reason` | [AttestationRevocationReason](#akash.verification.v1.AttestationRevocationReason) |  | Reason is the typed revocation reason recorded against the attestation. |
 | `evidence_hash` | [bytes](#bytes) |  | EvidenceHash is an opaque hash referencing the evidence backing the revocation (off-chain). |
 
 

 

 
 <a name="akash.verification.v1.MsgRevokeAttestationResponse"></a>

 ### MsgRevokeAttestationResponse
 MsgRevokeAttestationResponse is the response type for Msg/RevokeAttestation.

 

 

 
 <a name="akash.verification.v1.MsgRevokeAuditorAttestations"></a>

 ### MsgRevokeAuditorAttestations
 MsgRevokeAuditorAttestations is the governance-authority-signed message
used to revoke every active attestation issued by a single auditor, with a
typed reason, fault attribution, and evidence reference.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `authority` | [string](#string) |  | Authority is the bech32 address of the governance account authorized to revoke attestations and the signer of this message. |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the auditor whose attestations are being revoked en masse. |
 | `reason` | [GovernanceAttestationReason](#akash.verification.v1.GovernanceAttestationReason) |  | Reason is the typed governance attestation reason recorded against the revocations. |
 | `fault_attribution` | [FaultAttribution](#akash.verification.v1.FaultAttribution) |  | FaultAttribution is the fault attribution recorded against the revocations. |
 | `evidence_hash` | [bytes](#bytes) |  | EvidenceHash is an opaque hash referencing the evidence backing the revocations (off-chain). |
 
 

 

 
 <a name="akash.verification.v1.MsgRevokeAuditorAttestationsResponse"></a>

 ### MsgRevokeAuditorAttestationsResponse
 MsgRevokeAuditorAttestationsResponse is the response type for
Msg/RevokeAuditorAttestations.

 

 

 
 <a name="akash.verification.v1.MsgRevokeProviderAttestation"></a>

 ### MsgRevokeProviderAttestation
 MsgRevokeProviderAttestation is the governance-authority-signed message
used to revoke a single attestation issued by a specific auditor for a
specific provider, with a typed reason, fault attribution, and evidence
reference.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `authority` | [string](#string) |  | Authority is the bech32 address of the governance account authorized to revoke attestations and the signer of this message. |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the provider whose attestation is being revoked. |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the auditor whose attestation is being revoked. |
 | `reason` | [GovernanceAttestationReason](#akash.verification.v1.GovernanceAttestationReason) |  | Reason is the typed governance attestation reason recorded against the revocation. |
 | `fault_attribution` | [FaultAttribution](#akash.verification.v1.FaultAttribution) |  | FaultAttribution is the fault attribution recorded at revocation. |
 | `evidence_hash` | [bytes](#bytes) |  | EvidenceHash is an opaque hash referencing the evidence backing the revocation (off-chain). |
 
 

 

 
 <a name="akash.verification.v1.MsgRevokeProviderAttestationResponse"></a>

 ### MsgRevokeProviderAttestationResponse
 MsgRevokeProviderAttestationResponse is the response type for
Msg/RevokeProviderAttestation.

 

 

 
 <a name="akash.verification.v1.MsgSettleAuditEscrow"></a>

 ### MsgSettleAuditEscrow
 MsgSettleAuditEscrow is the governance-authority-signed message used to
settle an unconsumed audit escrow with an explicit reason, fault
attribution, and evidence reference. NoFault returns provider-funded coins
to the provider; ProviderFault returns the fee to the provider and slashes
the provider deposit to the community pool.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `authority` | [string](#string) |  | Authority is the bech32 address of the governance account authorized to settle the escrow and the signer of this message. |
 | `audit_escrow_id` | [uint64](#uint64) |  | AuditEscrowID is the identifier of the audit escrow being settled. |
 | `reason` | [AuditEscrowSettlementReason](#akash.verification.v1.AuditEscrowSettlementReason) |  | Reason is the typed settlement reason recorded against the escrow. |
 | `fault_attribution` | [FaultAttribution](#akash.verification.v1.FaultAttribution) |  | FaultAttribution is the fault attribution recorded against the escrow at settlement; it must be consistent with the reason. |
 | `evidence_hash` | [bytes](#bytes) |  | EvidenceHash is an opaque hash referencing the evidence considered at settlement (off-chain). |
 
 

 

 
 <a name="akash.verification.v1.MsgSettleAuditEscrowResponse"></a>

 ### MsgSettleAuditEscrowResponse
 MsgSettleAuditEscrowResponse is the response type for Msg/SettleAuditEscrow.

 

 

 
 <a name="akash.verification.v1.MsgSlashProviderBond"></a>

 ### MsgSlashProviderBond
 MsgSlashProviderBond is the governance-authority-signed message used to
slash a fraction of a provider's verification bond for a typed reason,
backed by an evidence reference. The fraction is a LegacyDec in [0, 1].

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `authority` | [string](#string) |  | Authority is the bech32 address of the governance account authorized to slash provider bonds and the signer of this message. |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the provider whose bond is being slashed. |
 | `slash_fraction` | [string](#string) |  | SlashFraction is the fraction of the bond to slash, expressed as a LegacyDec in the inclusive range [0, 1]. |
 | `reason` | [ProviderBondSlashReason](#akash.verification.v1.ProviderBondSlashReason) |  | Reason is the typed provider bond slash reason recorded against the slash. |
 | `evidence_hash` | [bytes](#bytes) |  | EvidenceHash is an opaque hash referencing the evidence backing the slash (off-chain). |
 
 

 

 
 <a name="akash.verification.v1.MsgSlashProviderBondResponse"></a>

 ### MsgSlashProviderBondResponse
 MsgSlashProviderBondResponse is the response type for
Msg/SlashProviderBond.

 

 

 
 <a name="akash.verification.v1.MsgSubmitAttestation"></a>

 ### MsgSubmitAttestation
 MsgSubmitAttestation is the auditor-signed message used to submit an
attestation about a provider at a requested tier and capability set. The
auditor includes a typed evidence reference, the audit fee, an anti-griefing
deposit, and the identifier of the audit escrow that authorizes the
attestation. The first valid submission against an open, matching, unexpired
escrow consumes it and records the submitting auditor.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the attested provider. |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the issuing auditor and the signer of this message. |
 | `tier` | [VerificationTier](#akash.verification.v1.VerificationTier) |  | Tier is the verification tier the auditor asserts for the provider. |
 | `capabilities` | [CapabilityFlag](#akash.verification.v1.CapabilityFlag) | repeated | Capabilities is the set of capability flags the auditor asserts. The attestation must include at least the escrow's requested capabilities; additional capabilities the auditor verified may be included. |
 | `evidence_hash` | [bytes](#bytes) |  | EvidenceHash is an opaque hash referencing the auditor's evidence (off-chain). |
 | `fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Fee is the audit fee paid by the provider for this attestation, sourced from the consumed audit escrow. |
 | `deposit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Deposit is the anti-griefing deposit posted by the auditor alongside the attestation. |
 | `audit_escrow_id` | [uint64](#uint64) |  | AuditEscrowID is the identifier of the open audit escrow that authorizes this attestation. |
 
 

 

 
 <a name="akash.verification.v1.MsgSubmitAttestationResponse"></a>

 ### MsgSubmitAttestationResponse
 MsgSubmitAttestationResponse is the response type for Msg/SubmitAttestation.

 

 

 
 <a name="akash.verification.v1.MsgUpdateParams"></a>

 ### MsgUpdateParams
 MsgUpdateParams is the governance-authority-signed message used to update
the x/verification module parameters. All parameters must be supplied.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `authority` | [string](#string) |  | Authority is the bech32 address of the governance account authorized to update module parameters and the signer of this message. |
 | `params` | [Params](#akash.verification.v1.Params) |  | Params defines the x/verification parameters to update. All parameters must be supplied. |
 
 

 

 
 <a name="akash.verification.v1.MsgUpdateParamsResponse"></a>

 ### MsgUpdateParamsResponse
 MsgUpdateParamsResponse is the response type for Msg/UpdateParams.

 

 

 
 <a name="akash.verification.v1.MsgWithdrawProviderBond"></a>

 ### MsgWithdrawProviderBond
 MsgWithdrawProviderBond is the provider-signed message used to initiate
withdrawal of part or all of the provider bond. Withdrawal enqueues an
unbonding entry against the provider bond record.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the bond-withdrawing provider and the signer of this message. |
 | `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Amount is the coin amount being withdrawn from the provider bond. |
 
 

 

 
 <a name="akash.verification.v1.MsgWithdrawProviderBondResponse"></a>

 ### MsgWithdrawProviderBondResponse
 MsgWithdrawProviderBondResponse is the response type for
Msg/WithdrawProviderBond.

 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/verification/v1/query.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/verification/v1/query.proto
 

 
 <a name="akash.verification.v1.QueryAttestationRequest"></a>

 ### QueryAttestationRequest
 QueryAttestationRequest is the request type for the Query/Attestation RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the attested provider. |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the auditor that submitted the attestation being requested. |
 
 

 

 
 <a name="akash.verification.v1.QueryAttestationResponse"></a>

 ### QueryAttestationResponse
 QueryAttestationResponse is the response type for the Query/Attestation RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `attestation` | [AttestationRecord](#akash.verification.v1.AttestationRecord) |  | Attestation is the AttestationRecord for the requested (provider, auditor) pair. |
 
 

 

 
 <a name="akash.verification.v1.QueryAuditEscrowRequest"></a>

 ### QueryAuditEscrowRequest
 QueryAuditEscrowRequest is the request type for the Query/AuditEscrow RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [uint64](#uint64) |  | ID is the numeric identifier of the audit escrow being requested. |
 
 

 

 
 <a name="akash.verification.v1.QueryAuditEscrowResponse"></a>

 ### QueryAuditEscrowResponse
 QueryAuditEscrowResponse is the response type for the Query/AuditEscrow RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `escrow` | [AuditEscrowRecord](#akash.verification.v1.AuditEscrowRecord) |  | Escrow is the AuditEscrowRecord for the requested id. |
 
 

 

 
 <a name="akash.verification.v1.QueryAuditorAttestationsRequest"></a>

 ### QueryAuditorAttestationsRequest
 QueryAuditorAttestationsRequest is the request type for the
Query/AuditorAttestations RPC method.

Note: this request intentionally does not include a status filter — callers
receive every attestation the auditor has emitted, regardless of status.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the auditor whose attestations are being queried. |
 | `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | Pagination is the standard Cosmos page-request used to paginate results. |
 
 

 

 
 <a name="akash.verification.v1.QueryAuditorAttestationsResponse"></a>

 ### QueryAuditorAttestationsResponse
 QueryAuditorAttestationsResponse is the response type for the
Query/AuditorAttestations RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `attestations` | [AttestationRecord](#akash.verification.v1.AttestationRecord) | repeated | Attestations is the page of AttestationRecord values emitted by the requested auditor. |
 | `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | Pagination is the standard Cosmos page-response carrying the next-key and total counters for the matched set. |
 
 

 

 
 <a name="akash.verification.v1.QueryAuditorRequest"></a>

 ### QueryAuditorRequest
 QueryAuditorRequest is the request type for the Query/Auditor RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `auditor` | [string](#string) |  | Auditor is the bech32 account address of the auditor being queried. Example: "akash1...". |
 
 

 

 
 <a name="akash.verification.v1.QueryAuditorResponse"></a>

 ### QueryAuditorResponse
 QueryAuditorResponse is the response type for the Query/Auditor RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `auditor` | [AuditorRecord](#akash.verification.v1.AuditorRecord) |  | Auditor is the on-chain AuditorRecord for the requested address. |
 
 

 

 
 <a name="akash.verification.v1.QueryAuditorsRequest"></a>

 ### QueryAuditorsRequest
 QueryAuditorsRequest is the request type for the Query/Auditors RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `status_filter` | [AuditorStatus](#akash.verification.v1.AuditorStatus) |  | StatusFilter restricts results to auditors with the given AuditorStatus. Set to AuditorStatusUnspecified (0) to return auditors in any status. |
 | `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | Pagination is the standard Cosmos page-request used to paginate results. |
 
 

 

 
 <a name="akash.verification.v1.QueryAuditorsResponse"></a>

 ### QueryAuditorsResponse
 QueryAuditorsResponse is the response type for the Query/Auditors RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `auditors` | [AuditorRecord](#akash.verification.v1.AuditorRecord) | repeated | Auditors is the page of AuditorRecord values matching the request filter. |
 | `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | Pagination is the standard Cosmos page-response carrying the next-key and total counters for the matched set. |
 
 

 

 
 <a name="akash.verification.v1.QueryDiscrepanciesRequest"></a>

 ### QueryDiscrepanciesRequest
 QueryDiscrepanciesRequest is the request type for the Query/Discrepancies RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `status_filter` | [DiscrepancyStatus](#akash.verification.v1.DiscrepancyStatus) |  | StatusFilter restricts results to discrepancies in the given DiscrepancyStatus. Set to DiscrepancyStatusUnspecified (0) to return discrepancies in any status. |
 | `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | Pagination is the standard Cosmos page-request used to paginate results. |
 
 

 

 
 <a name="akash.verification.v1.QueryDiscrepanciesResponse"></a>

 ### QueryDiscrepanciesResponse
 QueryDiscrepanciesResponse is the response type for the Query/Discrepancies RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `discrepancies` | [DiscrepancyEvent](#akash.verification.v1.DiscrepancyEvent) | repeated | Discrepancies is the page of DiscrepancyEvent values matching the request. |
 | `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | Pagination is the standard Cosmos page-response carrying the next-key and total counters for the matched set. |
 
 

 

 
 <a name="akash.verification.v1.QueryDiscrepancyRequest"></a>

 ### QueryDiscrepancyRequest
 QueryDiscrepancyRequest is the request type for the Query/Discrepancy RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `id` | [uint64](#uint64) |  | ID is the numeric identifier of the discrepancy event being requested. |
 
 

 

 
 <a name="akash.verification.v1.QueryDiscrepancyResponse"></a>

 ### QueryDiscrepancyResponse
 QueryDiscrepancyResponse is the response type for the Query/Discrepancy RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `discrepancy` | [DiscrepancyEvent](#akash.verification.v1.DiscrepancyEvent) |  | Discrepancy is the DiscrepancyEvent for the requested id. |
 
 

 

 
 <a name="akash.verification.v1.QueryParamsRequest"></a>

 ### QueryParamsRequest
 QueryParamsRequest is the request type for the Query/Params RPC method.

 

 

 
 <a name="akash.verification.v1.QueryParamsResponse"></a>

 ### QueryParamsResponse
 QueryParamsResponse is the response type for the Query/Params RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `params` | [Params](#akash.verification.v1.Params) |  | Params is the current parameter set for the verification module. |
 
 

 

 
 <a name="akash.verification.v1.QueryProviderAttestationsRequest"></a>

 ### QueryProviderAttestationsRequest
 QueryProviderAttestationsRequest is the request type for the
Query/ProviderAttestations RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the provider whose attestations are being queried. |
 | `status_filter` | [AttestationStatus](#akash.verification.v1.AttestationStatus) |  | StatusFilter restricts results to attestations whose AttestationStatus matches the filter. Set to AttestationStatusUnspecified (0) to return attestations in any status. |
 | `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | Pagination is the standard Cosmos page-request used to paginate results. |
 
 

 

 
 <a name="akash.verification.v1.QueryProviderAttestationsResponse"></a>

 ### QueryProviderAttestationsResponse
 QueryProviderAttestationsResponse is the response type for the
Query/ProviderAttestations RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `attestations` | [AttestationRecord](#akash.verification.v1.AttestationRecord) | repeated | Attestations is the page of AttestationRecord values matching the request. |
 | `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | Pagination is the standard Cosmos page-response carrying the next-key and total counters for the matched set. |
 
 

 

 
 <a name="akash.verification.v1.QueryProviderAuditEscrowsRequest"></a>

 ### QueryProviderAuditEscrowsRequest
 QueryProviderAuditEscrowsRequest is the request type for the
Query/ProviderAuditEscrows RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the provider whose audit-escrow records are being queried. |
 | `status_filter` | [AuditEscrowStatus](#akash.verification.v1.AuditEscrowStatus) |  | StatusFilter restricts results to audit-escrow records whose AuditEscrowStatus matches the filter. Set to AuditEscrowStatusUnspecified (0) to return records in any status. |
 | `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | Pagination is the standard Cosmos page-request used to paginate results. |
 
 

 

 
 <a name="akash.verification.v1.QueryProviderAuditEscrowsResponse"></a>

 ### QueryProviderAuditEscrowsResponse
 QueryProviderAuditEscrowsResponse is the response type for the
Query/ProviderAuditEscrows RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `escrows` | [AuditEscrowRecord](#akash.verification.v1.AuditEscrowRecord) | repeated | Escrows is the page of AuditEscrowRecord values matching the request. |
 | `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | Pagination is the standard Cosmos page-response carrying the next-key and total counters for the matched set. |
 
 

 

 
 <a name="akash.verification.v1.QueryProviderBondRequest"></a>

 ### QueryProviderBondRequest
 QueryProviderBondRequest is the request type for the Query/ProviderBond RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the provider whose bond record is being queried. |
 
 

 

 
 <a name="akash.verification.v1.QueryProviderBondResponse"></a>

 ### QueryProviderBondResponse
 QueryProviderBondResponse is the response type for the Query/ProviderBond RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `bond` | [ProviderBondRecord](#akash.verification.v1.ProviderBondRecord) |  | Bond is the on-chain ProviderBondRecord for the requested provider. |
 | `required_for_current_tier` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | RequiredForCurrentTier is the required bond amount for the provider's current tier. |
 
 

 

 
 <a name="akash.verification.v1.QueryProviderSnapshotRequest"></a>

 ### QueryProviderSnapshotRequest
 QueryProviderSnapshotRequest is the request type for the
Query/ProviderSnapshot RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the provider whose latest snapshot record is being queried. |
 
 

 

 
 <a name="akash.verification.v1.QueryProviderSnapshotResponse"></a>

 ### QueryProviderSnapshotResponse
 QueryProviderSnapshotResponse is the response type for the
Query/ProviderSnapshot RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `snapshot` | [ProviderSnapshotRecord](#akash.verification.v1.ProviderSnapshotRecord) |  | Snapshot is the latest ProviderSnapshotRecord posted on-chain by the requested provider. |
 
 

 

 
 <a name="akash.verification.v1.QueryProviderVerificationGraceRequest"></a>

 ### QueryProviderVerificationGraceRequest
 QueryProviderVerificationGraceRequest is the request type for the
Query/ProviderVerificationGrace RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `provider` | [string](#string) |  | Provider is the bech32 account address of the provider whose verification-grace record is being queried. |
 
 

 

 
 <a name="akash.verification.v1.QueryProviderVerificationGraceResponse"></a>

 ### QueryProviderVerificationGraceResponse
 QueryProviderVerificationGraceResponse is the response type for the
Query/ProviderVerificationGrace RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `grace` | [ProviderVerificationGraceRecord](#akash.verification.v1.ProviderVerificationGraceRecord) |  | Grace is the ProviderVerificationGraceRecord currently tracked for the requested provider. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.verification.v1.Query"></a>

 ### Query
 Query defines the gRPC querier service for the verification package.

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `Auditor` | [QueryAuditorRequest](#akash.verification.v1.QueryAuditorRequest) | [QueryAuditorResponse](#akash.verification.v1.QueryAuditorResponse) | Auditor returns the on-chain record for a single auditor identified by its bech32 address. buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE buf:lint:ignore RPC_RESPONSE_STANDARD_NAME | GET|/akash/verification/v1/auditors/{auditor}|
 | `Auditors` | [QueryAuditorsRequest](#akash.verification.v1.QueryAuditorsRequest) | [QueryAuditorsResponse](#akash.verification.v1.QueryAuditorsResponse) | Auditors returns a paginated list of auditor records, optionally filtered by AuditorStatus. buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE buf:lint:ignore RPC_RESPONSE_STANDARD_NAME | GET|/akash/verification/v1/auditors|
 | `Attestation` | [QueryAttestationRequest](#akash.verification.v1.QueryAttestationRequest) | [QueryAttestationResponse](#akash.verification.v1.QueryAttestationResponse) | Attestation returns the attestation record for a specific (provider, auditor) pair. buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE buf:lint:ignore RPC_RESPONSE_STANDARD_NAME | GET|/akash/verification/v1/attestations/{provider}/{auditor}|
 | `ProviderAttestations` | [QueryProviderAttestationsRequest](#akash.verification.v1.QueryProviderAttestationsRequest) | [QueryProviderAttestationsResponse](#akash.verification.v1.QueryProviderAttestationsResponse) | ProviderAttestations returns a paginated list of all attestation records for the given provider, optionally filtered by AttestationStatus. buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE buf:lint:ignore RPC_RESPONSE_STANDARD_NAME | GET|/akash/verification/v1/providers/{provider}/attestations|
 | `AuditorAttestations` | [QueryAuditorAttestationsRequest](#akash.verification.v1.QueryAuditorAttestationsRequest) | [QueryAuditorAttestationsResponse](#akash.verification.v1.QueryAuditorAttestationsResponse) | AuditorAttestations returns a paginated list of all attestation records submitted by the given auditor. No status filter is applied; callers can page through every attestation the auditor has emitted. buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE buf:lint:ignore RPC_RESPONSE_STANDARD_NAME | GET|/akash/verification/v1/auditors/{auditor}/attestations|
 | `Discrepancy` | [QueryDiscrepancyRequest](#akash.verification.v1.QueryDiscrepancyRequest) | [QueryDiscrepancyResponse](#akash.verification.v1.QueryDiscrepancyResponse) | Discrepancy returns a single discrepancy event by its numeric id. buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE buf:lint:ignore RPC_RESPONSE_STANDARD_NAME | GET|/akash/verification/v1/discrepancies/{id}|
 | `Discrepancies` | [QueryDiscrepanciesRequest](#akash.verification.v1.QueryDiscrepanciesRequest) | [QueryDiscrepanciesResponse](#akash.verification.v1.QueryDiscrepanciesResponse) | Discrepancies returns a paginated list of discrepancy events, optionally filtered by DiscrepancyStatus. buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE buf:lint:ignore RPC_RESPONSE_STANDARD_NAME | GET|/akash/verification/v1/discrepancies|
 | `AuditEscrow` | [QueryAuditEscrowRequest](#akash.verification.v1.QueryAuditEscrowRequest) | [QueryAuditEscrowResponse](#akash.verification.v1.QueryAuditEscrowResponse) | AuditEscrow returns a single audit-escrow record by its numeric id. buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE buf:lint:ignore RPC_RESPONSE_STANDARD_NAME | GET|/akash/verification/v1/audit-escrows/{id}|
 | `ProviderAuditEscrows` | [QueryProviderAuditEscrowsRequest](#akash.verification.v1.QueryProviderAuditEscrowsRequest) | [QueryProviderAuditEscrowsResponse](#akash.verification.v1.QueryProviderAuditEscrowsResponse) | ProviderAuditEscrows returns a paginated list of audit-escrow records opened by the given provider, optionally filtered by AuditEscrowStatus. buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE buf:lint:ignore RPC_RESPONSE_STANDARD_NAME | GET|/akash/verification/v1/providers/{provider}/audit-escrows|
 | `ProviderVerificationGrace` | [QueryProviderVerificationGraceRequest](#akash.verification.v1.QueryProviderVerificationGraceRequest) | [QueryProviderVerificationGraceResponse](#akash.verification.v1.QueryProviderVerificationGraceResponse) | ProviderVerificationGrace returns the verification-grace record for the given provider, if one is currently tracked. buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE buf:lint:ignore RPC_RESPONSE_STANDARD_NAME | GET|/akash/verification/v1/providers/{provider}/grace|
 | `ProviderBond` | [QueryProviderBondRequest](#akash.verification.v1.QueryProviderBondRequest) | [QueryProviderBondResponse](#akash.verification.v1.QueryProviderBondResponse) | ProviderBond returns the provider's bond record. buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE buf:lint:ignore RPC_RESPONSE_STANDARD_NAME | GET|/akash/verification/v1/providers/{provider}/bond|
 | `ProviderSnapshot` | [QueryProviderSnapshotRequest](#akash.verification.v1.QueryProviderSnapshotRequest) | [QueryProviderSnapshotResponse](#akash.verification.v1.QueryProviderSnapshotResponse) | ProviderSnapshot returns the most recent provider snapshot record posted on-chain by the given provider. buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE buf:lint:ignore RPC_RESPONSE_STANDARD_NAME | GET|/akash/verification/v1/providers/{provider}/snapshot|
 | `Params` | [QueryParamsRequest](#akash.verification.v1.QueryParamsRequest) | [QueryParamsResponse](#akash.verification.v1.QueryParamsResponse) | Params returns the current parameter set for the verification module. buf:lint:ignore RPC_REQUEST_RESPONSE_UNIQUE buf:lint:ignore RPC_RESPONSE_STANDARD_NAME | GET|/akash/verification/v1/params|
 
  <!-- end services -->

 
 
 <a name="akash/verification/v1/service.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/verification/v1/service.proto
 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.verification.v1.Msg"></a>

 ### Msg
 Msg defines the x/verification Msg service. It carries all transaction RPCs
for the verification module, including auditor lifecycle, attestation
submission and revocation, audit escrow flow, provider bond management,
snapshot posting, discrepancy resolution, and parameter updates.

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `PostAuditorBond` | [MsgPostAuditorBond](#akash.verification.v1.MsgPostAuditorBond) | [MsgPostAuditorBondResponse](#akash.verification.v1.MsgPostAuditorBondResponse) | PostAuditorBond posts (or tops up) an auditor's verification bond. | |
 | `OpenAuditEscrow` | [MsgOpenAuditEscrow](#akash.verification.v1.MsgOpenAuditEscrow) | [MsgOpenAuditEscrowResponse](#akash.verification.v1.MsgOpenAuditEscrowResponse) | OpenAuditEscrow opens a new audit escrow funding a pending attestation. | |
 | `CancelAuditEscrow` | [MsgCancelAuditEscrow](#akash.verification.v1.MsgCancelAuditEscrow) | [MsgCancelAuditEscrowResponse](#akash.verification.v1.MsgCancelAuditEscrowResponse) | CancelAuditEscrow cancels an open, unconsumed audit escrow before expiry and returns the fee and provider deposit to the provider. | |
 | `SettleAuditEscrow` | [MsgSettleAuditEscrow](#akash.verification.v1.MsgSettleAuditEscrow) | [MsgSettleAuditEscrowResponse](#akash.verification.v1.MsgSettleAuditEscrowResponse) | SettleAuditEscrow settles an unconsumed audit escrow with an explicit reason, fault attribution, and evidence reference. | |
 | `SubmitAttestation` | [MsgSubmitAttestation](#akash.verification.v1.MsgSubmitAttestation) | [MsgSubmitAttestationResponse](#akash.verification.v1.MsgSubmitAttestationResponse) | SubmitAttestation submits an attestation about a provider; the first valid submission against a matching open escrow consumes it. | |
 | `RevokeAttestation` | [MsgRevokeAttestation](#akash.verification.v1.MsgRevokeAttestation) | [MsgRevokeAttestationResponse](#akash.verification.v1.MsgRevokeAttestationResponse) | RevokeAttestation revokes a previously submitted attestation with a typed reason and evidence reference. | |
 | `RemoveAttestation` | [MsgRemoveAttestation](#akash.verification.v1.MsgRemoveAttestation) | [MsgRemoveAttestationResponse](#akash.verification.v1.MsgRemoveAttestationResponse) | RemoveAttestation voluntarily removes an attestation associated with the signing provider. | |
 | `ResignAuditor` | [MsgResignAuditor](#akash.verification.v1.MsgResignAuditor) | [MsgResignAuditorResponse](#akash.verification.v1.MsgResignAuditorResponse) | ResignAuditor voluntarily exits the auditor role and begins unbonding of any posted auditor bond. | |
 | `PostProviderBond` | [MsgPostProviderBond](#akash.verification.v1.MsgPostProviderBond) | [MsgPostProviderBondResponse](#akash.verification.v1.MsgPostProviderBondResponse) | PostProviderBond posts (or tops up) a provider's resource-scaled verification bond. | |
 | `WithdrawProviderBond` | [MsgWithdrawProviderBond](#akash.verification.v1.MsgWithdrawProviderBond) | [MsgWithdrawProviderBondResponse](#akash.verification.v1.MsgWithdrawProviderBondResponse) | WithdrawProviderBond initiates withdrawal of part or all of a provider's verification bond. | |
 | `PostSnapshotHash` | [MsgPostSnapshotHash](#akash.verification.v1.MsgPostSnapshotHash) | [MsgPostSnapshotHashResponse](#akash.verification.v1.MsgPostSnapshotHashResponse) | PostSnapshotHash posts the provider's most recent resource snapshot hash and inline resource summary. | |
 | `RegisterAuditor` | [MsgRegisterAuditor](#akash.verification.v1.MsgRegisterAuditor) | [MsgRegisterAuditorResponse](#akash.verification.v1.MsgRegisterAuditorResponse) | RegisterAuditor registers a new auditor with a maximum attestation tier; governance only. | |
 | `RenewAuditor` | [MsgRenewAuditor](#akash.verification.v1.MsgRenewAuditor) | [MsgRenewAuditorResponse](#akash.verification.v1.MsgRenewAuditorResponse) | RenewAuditor renews an auditor's registration and resets the renewal deadline; governance only. | |
 | `RemoveAuditor` | [MsgRemoveAuditor](#akash.verification.v1.MsgRemoveAuditor) | [MsgRemoveAuditorResponse](#akash.verification.v1.MsgRemoveAuditorResponse) | RemoveAuditor removes an auditor from the active set; governance only. | |
 | `RevokeProviderAttestation` | [MsgRevokeProviderAttestation](#akash.verification.v1.MsgRevokeProviderAttestation) | [MsgRevokeProviderAttestationResponse](#akash.verification.v1.MsgRevokeProviderAttestationResponse) | RevokeProviderAttestation revokes a single attestation for a specific provider/auditor pair; governance only. | |
 | `RevokeAllProviderAttestations` | [MsgRevokeAllProviderAttestations](#akash.verification.v1.MsgRevokeAllProviderAttestations) | [MsgRevokeAllProviderAttestationsResponse](#akash.verification.v1.MsgRevokeAllProviderAttestationsResponse) | RevokeAllProviderAttestations revokes every active attestation for a single provider; governance only. | |
 | `RevokeAuditorAttestations` | [MsgRevokeAuditorAttestations](#akash.verification.v1.MsgRevokeAuditorAttestations) | [MsgRevokeAuditorAttestationsResponse](#akash.verification.v1.MsgRevokeAuditorAttestationsResponse) | RevokeAuditorAttestations revokes every active attestation issued by a single auditor; governance only. | |
 | `ResolveDiscrepancy` | [MsgResolveDiscrepancy](#akash.verification.v1.MsgResolveDiscrepancy) | [MsgResolveDiscrepancyResponse](#akash.verification.v1.MsgResolveDiscrepancyResponse) | ResolveDiscrepancy resolves a pending discrepancy between two auditors over the same provider; governance only. | |
 | `SlashProviderBond` | [MsgSlashProviderBond](#akash.verification.v1.MsgSlashProviderBond) | [MsgSlashProviderBondResponse](#akash.verification.v1.MsgSlashProviderBondResponse) | SlashProviderBond slashes a fraction of a provider's verification bond; governance only. | |
 | `UpdateParams` | [MsgUpdateParams](#akash.verification.v1.MsgUpdateParams) | [MsgUpdateParamsResponse](#akash.verification.v1.MsgUpdateParamsResponse) | UpdateParams updates the x/verification module parameters; governance only. | |
 
  <!-- end services -->

 
 
 <a name="akash/wasm/v1/event.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/wasm/v1/event.proto
 

 
 <a name="akash.wasm.v1.EventMsgBlocked"></a>

 ### EventMsgBlocked
 EventMsgBlocked is triggered when smart contract does not
pass message filter

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `contract_address` | [string](#string) |  |  |
 | `msg_type` | [string](#string) |  |  |
 | `reason` | [string](#string) |  |  |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/wasm/v1/params.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/wasm/v1/params.proto
 

 
 <a name="akash.wasm.v1.Params"></a>

 ### Params
 Params defines the parameters for the x/wasm package.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `blocked_addresses` | [string](#string) | repeated |  |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/wasm/v1/genesis.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/wasm/v1/genesis.proto
 

 
 <a name="akash.wasm.v1.GenesisState"></a>

 ### GenesisState
 GenesisState stores slice of genesis wasm parameters.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `params` | [Params](#akash.wasm.v1.Params) |  | Params holds parameters of the genesis of akash wasm. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/wasm/v1/paramsmsg.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/wasm/v1/paramsmsg.proto
 

 
 <a name="akash.wasm.v1.MsgUpdateParams"></a>

 ### MsgUpdateParams
 MsgUpdateParams is the Msg/UpdateParams request type.

Since: akash v1.0.0

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `authority` | [string](#string) |  | authority is the address of the governance account. |
 | `params` | [Params](#akash.wasm.v1.Params) |  | params defines the x/wasm parameters to update.

NOTE: All parameters must be supplied. |
 
 

 

 
 <a name="akash.wasm.v1.MsgUpdateParamsResponse"></a>

 ### MsgUpdateParamsResponse
 MsgUpdateParamsResponse defines the response structure for executing a
MsgUpdateParams message.

Since: akash v1.0.0

 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

  <!-- end services -->

 
 
 <a name="akash/wasm/v1/query.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/wasm/v1/query.proto
 

 
 <a name="akash.wasm.v1.QueryParamsRequest"></a>

 ### QueryParamsRequest
 QueryParamsRequest is the request type for the Query/Params RPC method.

 

 

 
 <a name="akash.wasm.v1.QueryParamsResponse"></a>

 ### QueryParamsResponse
 QueryParamsResponse is the response type for the Query/Params RPC method.

 
 | Field | Type | Label | Description |
 | ----- | ---- | ----- | ----------- |
 | `params` | [Params](#akash.wasm.v1.Params) |  | params defines the parameters of the module. |
 
 

 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.wasm.v1.Query"></a>

 ### Query
 Query defines the gRPC querier service of the wasm package.

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `Params` | [QueryParamsRequest](#akash.wasm.v1.QueryParamsRequest) | [QueryParamsResponse](#akash.wasm.v1.QueryParamsResponse) | Params returns the total set of wasm parameters. | GET|/akash/wasm/v1/params|
 
  <!-- end services -->

 
 
 <a name="akash/wasm/v1/service.proto"></a>
 <p align="right"><a href="#top">Top</a></p>

 ## akash/wasm/v1/service.proto
 

  <!-- end messages -->

  <!-- end enums -->

  <!-- end HasExtensions -->

 
 <a name="akash.wasm.v1.Msg"></a>

 ### Msg
 Msg defines the wasm Msg service

 | Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
 | ----------- | ------------ | ------------- | ------------| ------- | -------- |
 | `UpdateParams` | [MsgUpdateParams](#akash.wasm.v1.MsgUpdateParams) | [MsgUpdateParamsResponse](#akash.wasm.v1.MsgUpdateParamsResponse) | UpdateParams defines a governance operation for updating the x/wasm module parameters. The authority is hard-coded to the x/gov module account.

Since: akash v2.0.0 | |
 
  <!-- end services -->

 

 ## Scalar Value Types

 | .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
 | ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
 | <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
 | <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
 | <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
 | <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
 | <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
 | <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
 | <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
 | <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
 | <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
 | <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
 | <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
 | <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
 | <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
 | <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
 | <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |
 
