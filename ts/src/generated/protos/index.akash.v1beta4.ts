import { patched } from "./nodePatchMessage.ts";

export { ResourceValue } from "./akash/base/resources/v1beta4/resourcevalue.ts";
export { CPU } from "./akash/base/resources/v1beta4/cpu.ts";
export { Endpoint, Endpoint_Kind } from "./akash/base/resources/v1beta4/endpoint.ts";
export { GPU } from "./akash/base/resources/v1beta4/gpu.ts";
export { Memory } from "./akash/base/resources/v1beta4/memory.ts";
export { Storage } from "./akash/base/resources/v1beta4/storage.ts";
export { Resources } from "./akash/base/resources/v1beta4/resources.ts";

import { ResourceUnit as _ResourceUnit } from "./akash/deployment/v1beta4/resourceunit.ts";
export const ResourceUnit = patched(_ResourceUnit);

import { GroupSpec as _GroupSpec } from "./akash/deployment/v1beta4/groupspec.ts";
export const GroupSpec = patched(_GroupSpec);
export { MsgCreateDeploymentResponse, MsgUpdateDeployment, MsgUpdateDeploymentResponse, MsgCloseDeployment, MsgCloseDeploymentResponse } from "./akash/deployment/v1beta4/deploymentmsg.ts";

import { MsgCreateDeployment as _MsgCreateDeployment } from "./akash/deployment/v1beta4/deploymentmsg.ts";
export const MsgCreateDeployment = patched(_MsgCreateDeployment);
export { DeploymentFilters, GroupFilters } from "./akash/deployment/v1beta4/filters.ts";
export { Group_State } from "./akash/deployment/v1beta4/group.ts";

import { Group as _Group } from "./akash/deployment/v1beta4/group.ts";
export const Group = patched(_Group);
export { Params } from "./akash/deployment/v1beta4/params.ts";

import { GenesisDeployment as _GenesisDeployment, GenesisState as _GenesisState } from "./akash/deployment/v1beta4/genesis.ts";
export const GenesisDeployment = patched(_GenesisDeployment);
export const GenesisState = patched(_GenesisState);
export { MsgCloseGroup, MsgCloseGroupResponse, MsgPauseGroup, MsgPauseGroupResponse, MsgStartGroup, MsgStartGroupResponse } from "./akash/deployment/v1beta4/groupmsg.ts";
export { MsgUpdateParams, MsgUpdateParamsResponse } from "./akash/deployment/v1beta4/paramsmsg.ts";
export { QueryDeploymentsRequest, QueryDeploymentRequest, QueryGroupRequest, QueryParamsRequest, QueryParamsResponse } from "./akash/deployment/v1beta4/query.ts";

import { QueryDeploymentsResponse as _QueryDeploymentsResponse, QueryDeploymentResponse as _QueryDeploymentResponse, QueryGroupResponse as _QueryGroupResponse } from "./akash/deployment/v1beta4/query.ts";
export const QueryDeploymentsResponse = patched(_QueryDeploymentsResponse);
export const QueryDeploymentResponse = patched(_QueryDeploymentResponse);
export const QueryGroupResponse = patched(_QueryGroupResponse);
export { EventProviderCreated, EventProviderUpdated, EventProviderDeleted } from "./akash/provider/v1beta4/event.ts";
export { Info, Provider } from "./akash/provider/v1beta4/provider.ts";
export { GenesisState as Provider_GenesisState } from "./akash/provider/v1beta4/genesis.ts";
export { MsgCreateProvider, MsgCreateProviderResponse, MsgUpdateProvider, MsgUpdateProviderResponse, MsgDeleteProvider, MsgDeleteProviderResponse } from "./akash/provider/v1beta4/msg.ts";
export { QueryProvidersRequest, QueryProvidersResponse, QueryProviderRequest, QueryProviderResponse } from "./akash/provider/v1beta4/query.ts";
