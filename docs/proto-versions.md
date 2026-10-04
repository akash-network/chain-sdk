# Protobuf Msg Types by Module and Version

This document provides a comprehensive mapping of all protobuf types and their available versions in the chain-sdk project.

## Node Modules - Msg Types

| Module | Msg Type | v1beta3 | v1beta4 | v1beta5 | v1 |
|:-------|:---------|:-------:|:-------:|:-------:|:--:|
| **audit** | MsgSignProviderAttributes | ✓ | | | ✓ |
| | MsgDeleteProviderAttributes | ✓ | | | ✓ |
| **cert** | MsgCreateCertificate | ✓ | | | ✓ |
| | MsgRevokeCertificate | ✓ | | | ✓ |
| **deployment** | MsgCreateDeployment | ✓ | ✓ | | |
| | MsgUpdateDeployment | ✓ | ✓ | | |
| | MsgCloseDeployment | ✓ | ✓ | | |
| | MsgDepositDeployment | ✓ | | | |
| | MsgCloseGroup | ✓ | ✓ | | |
| | MsgPauseGroup | ✓ | ✓ | | |
| | MsgStartGroup | ✓ | ✓ | | |
| | MsgUpdateParams | | ✓ | | |
| **escrow** | MsgAccountDeposit | | | | ✓ |
| **market** | MsgCreateBid | | ✓ | ✓ | |
| | MsgCloseBid | | ✓ | ✓ | |
| | MsgCreateLease | | ✓ | ✓ | |
| | MsgWithdrawLease | | ✓ | ✓ | |
| | MsgCloseLease | | ✓ | ✓ | |
| | MsgUpdateParams | | | ✓ | |
| **provider** | MsgCreateProvider | ✓ | ✓ | | |
| | MsgUpdateProvider | ✓ | ✓ | | |
| | MsgDeleteProvider | ✓ | ✓ | | |
| **take** | MsgUpdateParams | | | | ✓ |

## Provider Modules - Proto Types

| Module | Proto Type | v1 | v2beta3 |
|:-------|:-----------|:--:|:-------:| 
| **inventory** | Cluster | ✓ | |
| | Node | ✓ | |
| | NodeCapabilities | ✓ | |
| | NodeResources | ✓ | |
| | ResourcePair | ✓ | |
| | CPU | ✓ | |
| | CPUInfo | ✓ | |
| | GPU | ✓ | |
| | GPUInfo | ✓ | |
| | Storage | ✓ | |
| | StorageInfo | ✓ | |
| **manifest** | Group | | ✓ |
| | Service | | ✓ |
| | ServiceExpose | | ✓ |
| | ServiceExposeHTTPOptions | | ✓ |
| | ServiceParams | | ✓ |
| | StorageParams | | ✓ |
| | ImageCredentials | | ✓ |
| **provider** | Status | ✓ | |
| | ClusterStatus | ✓ | |
| | BidEngineStatus | ✓ | |
| | Inventory | ✓ | |
| | ResourcesMetric | ✓ | |
| | Leases | ✓ | |
| | ReservationsMetric | ✓ | |
| | Reservations | ✓ | |
| | LeaseServiceStatus | ✓ | |
| | LeaseIPStatus | ✓ | |
| | ForwarderPortStatus | ✓ | |
| | ServiceStatus | ✓ | |
| | SendManifestRequest | ✓ | |
| | SendManifestResponse | ✓ | |
| | ServiceLogsRequest | ✓ | |
| | ServiceLogs | ✓ | |
| | ServiceLogsResponse | ✓ | |
| | ShellRequest | ✓ | |
| | ServiceStatusRequest | ✓ | |
| | ServiceStatusResponse | ✓ | |
