# ReportMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentRunning** | Pointer to **bool** |  | [optional] 
**AvailabilityZone** | Pointer to **string** |  | [optional] 
**CloudAccountId** | Pointer to **string** |  | [optional] 
**CloudProvider** | Pointer to **string** |  | [optional] 
**CloudRegion** | Pointer to **string** |  | [optional] 
**Cmdline** | Pointer to **string** |  | [optional] 
**ConnectionCount** | Pointer to **int32** |  | [optional] 
**CopyOf** | Pointer to **string** |  | [optional] 
**CpuMax** | Pointer to **float32** |  | [optional] 
**CpuUsage** | Pointer to **float32** |  | [optional] 
**DockerContainerCommand** | Pointer to **string** |  | [optional] 
**DockerContainerCreated** | Pointer to **string** |  | [optional] 
**DockerContainerIps** | Pointer to **[]string** |  | [optional] 
**DockerContainerName** | Pointer to **string** |  | [optional] 
**DockerContainerNetworkMode** | Pointer to **string** |  | [optional] 
**DockerContainerNetworks** | Pointer to **string** |  | [optional] 
**DockerContainerPorts** | Pointer to **string** |  | [optional] 
**DockerContainerState** | Pointer to **string** |  | [optional] 
**DockerContainerStateHuman** | Pointer to **string** |  | [optional] 
**DockerEnv** | Pointer to **string** |  | [optional] 
**DockerImageCreatedAt** | Pointer to **string** |  | [optional] 
**DockerImageId** | Pointer to **string** |  | [optional] 
**DockerImageName** | Pointer to **string** |  | [optional] 
**DockerImageNameWithTag** | Pointer to **string** |  | [optional] 
**DockerImageSize** | Pointer to **string** |  | [optional] 
**DockerImageTag** | Pointer to **string** |  | [optional] 
**DockerImageVirtualSize** | Pointer to **string** |  | [optional] 
**DockerLabel** | Pointer to **string** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**InstanceId** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**InterfaceIpMap** | Pointer to **string** |  | [optional] 
**InterfaceIps** | Pointer to **[]string** |  | [optional] 
**InterfaceNames** | Pointer to **[]string** |  | [optional] 
**IsConsoleVm** | Pointer to **bool** |  | [optional] 
**KernelId** | Pointer to **string** |  | [optional] 
**KernelVersion** | Pointer to **string** |  | [optional] 
**KubernetesClusterId** | Pointer to **string** |  | [optional] 
**KubernetesClusterName** | Pointer to **string** |  | [optional] 
**KubernetesCreated** | Pointer to **string** |  | [optional] 
**KubernetesIngressIp** | Pointer to **[]string** |  | [optional] 
**KubernetesIp** | Pointer to **string** |  | [optional] 
**KubernetesIsInHostNetwork** | Pointer to **bool** |  | [optional] 
**KubernetesLabels** | Pointer to **string** |  | [optional] 
**KubernetesNamespace** | Pointer to **string** |  | [optional] 
**KubernetesPorts** | Pointer to **[]string** |  | [optional] 
**KubernetesPublicIp** | Pointer to **string** |  | [optional] 
**KubernetesState** | Pointer to **string** |  | [optional] 
**KubernetesType** | Pointer to **string** |  | [optional] 
**LocalCidr** | Pointer to **[]string** |  | [optional] 
**LocalNetworks** | Pointer to **[]string** |  | [optional] 
**MemoryMax** | Pointer to **int32** |  | [optional] 
**MemoryUsage** | Pointer to **int32** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**NodeName** | Pointer to **string** |  | [optional] 
**NodeType** | Pointer to **string** |  | [optional] 
**OpenFiles** | Pointer to **[]string** |  | [optional] 
**OpenFilesCount** | Pointer to **int32** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**Pid** | Pointer to **int32** |  | [optional] 
**PodId** | Pointer to **string** |  | [optional] 
**PodName** | Pointer to **string** |  | [optional] 
**Ppid** | Pointer to **int32** |  | [optional] 
**PrivateIp** | Pointer to **[]string** |  | [optional] 
**Pseudo** | Pointer to **bool** |  | [optional] 
**PublicIp** | Pointer to **[]string** |  | [optional] 
**ResourceGroup** | Pointer to **string** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Threads** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**Uptime** | Pointer to **int32** |  | [optional] 
**UserDefinedTags** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewReportMetadata

`func NewReportMetadata() *ReportMetadata`

NewReportMetadata instantiates a new ReportMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportMetadataWithDefaults

`func NewReportMetadataWithDefaults() *ReportMetadata`

NewReportMetadataWithDefaults instantiates a new ReportMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentRunning

`func (o *ReportMetadata) GetAgentRunning() bool`

GetAgentRunning returns the AgentRunning field if non-nil, zero value otherwise.

### GetAgentRunningOk

`func (o *ReportMetadata) GetAgentRunningOk() (*bool, bool)`

GetAgentRunningOk returns a tuple with the AgentRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentRunning

`func (o *ReportMetadata) SetAgentRunning(v bool)`

SetAgentRunning sets AgentRunning field to given value.

### HasAgentRunning

`func (o *ReportMetadata) HasAgentRunning() bool`

HasAgentRunning returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *ReportMetadata) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *ReportMetadata) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *ReportMetadata) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *ReportMetadata) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetCloudAccountId

`func (o *ReportMetadata) GetCloudAccountId() string`

GetCloudAccountId returns the CloudAccountId field if non-nil, zero value otherwise.

### GetCloudAccountIdOk

`func (o *ReportMetadata) GetCloudAccountIdOk() (*string, bool)`

GetCloudAccountIdOk returns a tuple with the CloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccountId

`func (o *ReportMetadata) SetCloudAccountId(v string)`

SetCloudAccountId sets CloudAccountId field to given value.

### HasCloudAccountId

`func (o *ReportMetadata) HasCloudAccountId() bool`

HasCloudAccountId returns a boolean if a field has been set.

### GetCloudProvider

`func (o *ReportMetadata) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ReportMetadata) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ReportMetadata) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *ReportMetadata) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCloudRegion

`func (o *ReportMetadata) GetCloudRegion() string`

GetCloudRegion returns the CloudRegion field if non-nil, zero value otherwise.

### GetCloudRegionOk

`func (o *ReportMetadata) GetCloudRegionOk() (*string, bool)`

GetCloudRegionOk returns a tuple with the CloudRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudRegion

`func (o *ReportMetadata) SetCloudRegion(v string)`

SetCloudRegion sets CloudRegion field to given value.

### HasCloudRegion

`func (o *ReportMetadata) HasCloudRegion() bool`

HasCloudRegion returns a boolean if a field has been set.

### GetCmdline

`func (o *ReportMetadata) GetCmdline() string`

GetCmdline returns the Cmdline field if non-nil, zero value otherwise.

### GetCmdlineOk

`func (o *ReportMetadata) GetCmdlineOk() (*string, bool)`

GetCmdlineOk returns a tuple with the Cmdline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmdline

`func (o *ReportMetadata) SetCmdline(v string)`

SetCmdline sets Cmdline field to given value.

### HasCmdline

`func (o *ReportMetadata) HasCmdline() bool`

HasCmdline returns a boolean if a field has been set.

### GetConnectionCount

`func (o *ReportMetadata) GetConnectionCount() int32`

GetConnectionCount returns the ConnectionCount field if non-nil, zero value otherwise.

### GetConnectionCountOk

`func (o *ReportMetadata) GetConnectionCountOk() (*int32, bool)`

GetConnectionCountOk returns a tuple with the ConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCount

`func (o *ReportMetadata) SetConnectionCount(v int32)`

SetConnectionCount sets ConnectionCount field to given value.

### HasConnectionCount

`func (o *ReportMetadata) HasConnectionCount() bool`

HasConnectionCount returns a boolean if a field has been set.

### GetCopyOf

`func (o *ReportMetadata) GetCopyOf() string`

GetCopyOf returns the CopyOf field if non-nil, zero value otherwise.

### GetCopyOfOk

`func (o *ReportMetadata) GetCopyOfOk() (*string, bool)`

GetCopyOfOk returns a tuple with the CopyOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyOf

`func (o *ReportMetadata) SetCopyOf(v string)`

SetCopyOf sets CopyOf field to given value.

### HasCopyOf

`func (o *ReportMetadata) HasCopyOf() bool`

HasCopyOf returns a boolean if a field has been set.

### GetCpuMax

`func (o *ReportMetadata) GetCpuMax() float32`

GetCpuMax returns the CpuMax field if non-nil, zero value otherwise.

### GetCpuMaxOk

`func (o *ReportMetadata) GetCpuMaxOk() (*float32, bool)`

GetCpuMaxOk returns a tuple with the CpuMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMax

`func (o *ReportMetadata) SetCpuMax(v float32)`

SetCpuMax sets CpuMax field to given value.

### HasCpuMax

`func (o *ReportMetadata) HasCpuMax() bool`

HasCpuMax returns a boolean if a field has been set.

### GetCpuUsage

`func (o *ReportMetadata) GetCpuUsage() float32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *ReportMetadata) GetCpuUsageOk() (*float32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *ReportMetadata) SetCpuUsage(v float32)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *ReportMetadata) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetDockerContainerCommand

`func (o *ReportMetadata) GetDockerContainerCommand() string`

GetDockerContainerCommand returns the DockerContainerCommand field if non-nil, zero value otherwise.

### GetDockerContainerCommandOk

`func (o *ReportMetadata) GetDockerContainerCommandOk() (*string, bool)`

GetDockerContainerCommandOk returns a tuple with the DockerContainerCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContainerCommand

`func (o *ReportMetadata) SetDockerContainerCommand(v string)`

SetDockerContainerCommand sets DockerContainerCommand field to given value.

### HasDockerContainerCommand

`func (o *ReportMetadata) HasDockerContainerCommand() bool`

HasDockerContainerCommand returns a boolean if a field has been set.

### GetDockerContainerCreated

`func (o *ReportMetadata) GetDockerContainerCreated() string`

GetDockerContainerCreated returns the DockerContainerCreated field if non-nil, zero value otherwise.

### GetDockerContainerCreatedOk

`func (o *ReportMetadata) GetDockerContainerCreatedOk() (*string, bool)`

GetDockerContainerCreatedOk returns a tuple with the DockerContainerCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContainerCreated

`func (o *ReportMetadata) SetDockerContainerCreated(v string)`

SetDockerContainerCreated sets DockerContainerCreated field to given value.

### HasDockerContainerCreated

`func (o *ReportMetadata) HasDockerContainerCreated() bool`

HasDockerContainerCreated returns a boolean if a field has been set.

### GetDockerContainerIps

`func (o *ReportMetadata) GetDockerContainerIps() []string`

GetDockerContainerIps returns the DockerContainerIps field if non-nil, zero value otherwise.

### GetDockerContainerIpsOk

`func (o *ReportMetadata) GetDockerContainerIpsOk() (*[]string, bool)`

GetDockerContainerIpsOk returns a tuple with the DockerContainerIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContainerIps

`func (o *ReportMetadata) SetDockerContainerIps(v []string)`

SetDockerContainerIps sets DockerContainerIps field to given value.

### HasDockerContainerIps

`func (o *ReportMetadata) HasDockerContainerIps() bool`

HasDockerContainerIps returns a boolean if a field has been set.

### GetDockerContainerName

`func (o *ReportMetadata) GetDockerContainerName() string`

GetDockerContainerName returns the DockerContainerName field if non-nil, zero value otherwise.

### GetDockerContainerNameOk

`func (o *ReportMetadata) GetDockerContainerNameOk() (*string, bool)`

GetDockerContainerNameOk returns a tuple with the DockerContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContainerName

`func (o *ReportMetadata) SetDockerContainerName(v string)`

SetDockerContainerName sets DockerContainerName field to given value.

### HasDockerContainerName

`func (o *ReportMetadata) HasDockerContainerName() bool`

HasDockerContainerName returns a boolean if a field has been set.

### GetDockerContainerNetworkMode

`func (o *ReportMetadata) GetDockerContainerNetworkMode() string`

GetDockerContainerNetworkMode returns the DockerContainerNetworkMode field if non-nil, zero value otherwise.

### GetDockerContainerNetworkModeOk

`func (o *ReportMetadata) GetDockerContainerNetworkModeOk() (*string, bool)`

GetDockerContainerNetworkModeOk returns a tuple with the DockerContainerNetworkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContainerNetworkMode

`func (o *ReportMetadata) SetDockerContainerNetworkMode(v string)`

SetDockerContainerNetworkMode sets DockerContainerNetworkMode field to given value.

### HasDockerContainerNetworkMode

`func (o *ReportMetadata) HasDockerContainerNetworkMode() bool`

HasDockerContainerNetworkMode returns a boolean if a field has been set.

### GetDockerContainerNetworks

`func (o *ReportMetadata) GetDockerContainerNetworks() string`

GetDockerContainerNetworks returns the DockerContainerNetworks field if non-nil, zero value otherwise.

### GetDockerContainerNetworksOk

`func (o *ReportMetadata) GetDockerContainerNetworksOk() (*string, bool)`

GetDockerContainerNetworksOk returns a tuple with the DockerContainerNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContainerNetworks

`func (o *ReportMetadata) SetDockerContainerNetworks(v string)`

SetDockerContainerNetworks sets DockerContainerNetworks field to given value.

### HasDockerContainerNetworks

`func (o *ReportMetadata) HasDockerContainerNetworks() bool`

HasDockerContainerNetworks returns a boolean if a field has been set.

### GetDockerContainerPorts

`func (o *ReportMetadata) GetDockerContainerPorts() string`

GetDockerContainerPorts returns the DockerContainerPorts field if non-nil, zero value otherwise.

### GetDockerContainerPortsOk

`func (o *ReportMetadata) GetDockerContainerPortsOk() (*string, bool)`

GetDockerContainerPortsOk returns a tuple with the DockerContainerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContainerPorts

`func (o *ReportMetadata) SetDockerContainerPorts(v string)`

SetDockerContainerPorts sets DockerContainerPorts field to given value.

### HasDockerContainerPorts

`func (o *ReportMetadata) HasDockerContainerPorts() bool`

HasDockerContainerPorts returns a boolean if a field has been set.

### GetDockerContainerState

`func (o *ReportMetadata) GetDockerContainerState() string`

GetDockerContainerState returns the DockerContainerState field if non-nil, zero value otherwise.

### GetDockerContainerStateOk

`func (o *ReportMetadata) GetDockerContainerStateOk() (*string, bool)`

GetDockerContainerStateOk returns a tuple with the DockerContainerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContainerState

`func (o *ReportMetadata) SetDockerContainerState(v string)`

SetDockerContainerState sets DockerContainerState field to given value.

### HasDockerContainerState

`func (o *ReportMetadata) HasDockerContainerState() bool`

HasDockerContainerState returns a boolean if a field has been set.

### GetDockerContainerStateHuman

`func (o *ReportMetadata) GetDockerContainerStateHuman() string`

GetDockerContainerStateHuman returns the DockerContainerStateHuman field if non-nil, zero value otherwise.

### GetDockerContainerStateHumanOk

`func (o *ReportMetadata) GetDockerContainerStateHumanOk() (*string, bool)`

GetDockerContainerStateHumanOk returns a tuple with the DockerContainerStateHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContainerStateHuman

`func (o *ReportMetadata) SetDockerContainerStateHuman(v string)`

SetDockerContainerStateHuman sets DockerContainerStateHuman field to given value.

### HasDockerContainerStateHuman

`func (o *ReportMetadata) HasDockerContainerStateHuman() bool`

HasDockerContainerStateHuman returns a boolean if a field has been set.

### GetDockerEnv

`func (o *ReportMetadata) GetDockerEnv() string`

GetDockerEnv returns the DockerEnv field if non-nil, zero value otherwise.

### GetDockerEnvOk

`func (o *ReportMetadata) GetDockerEnvOk() (*string, bool)`

GetDockerEnvOk returns a tuple with the DockerEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerEnv

`func (o *ReportMetadata) SetDockerEnv(v string)`

SetDockerEnv sets DockerEnv field to given value.

### HasDockerEnv

`func (o *ReportMetadata) HasDockerEnv() bool`

HasDockerEnv returns a boolean if a field has been set.

### GetDockerImageCreatedAt

`func (o *ReportMetadata) GetDockerImageCreatedAt() string`

GetDockerImageCreatedAt returns the DockerImageCreatedAt field if non-nil, zero value otherwise.

### GetDockerImageCreatedAtOk

`func (o *ReportMetadata) GetDockerImageCreatedAtOk() (*string, bool)`

GetDockerImageCreatedAtOk returns a tuple with the DockerImageCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageCreatedAt

`func (o *ReportMetadata) SetDockerImageCreatedAt(v string)`

SetDockerImageCreatedAt sets DockerImageCreatedAt field to given value.

### HasDockerImageCreatedAt

`func (o *ReportMetadata) HasDockerImageCreatedAt() bool`

HasDockerImageCreatedAt returns a boolean if a field has been set.

### GetDockerImageId

`func (o *ReportMetadata) GetDockerImageId() string`

GetDockerImageId returns the DockerImageId field if non-nil, zero value otherwise.

### GetDockerImageIdOk

`func (o *ReportMetadata) GetDockerImageIdOk() (*string, bool)`

GetDockerImageIdOk returns a tuple with the DockerImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageId

`func (o *ReportMetadata) SetDockerImageId(v string)`

SetDockerImageId sets DockerImageId field to given value.

### HasDockerImageId

`func (o *ReportMetadata) HasDockerImageId() bool`

HasDockerImageId returns a boolean if a field has been set.

### GetDockerImageName

`func (o *ReportMetadata) GetDockerImageName() string`

GetDockerImageName returns the DockerImageName field if non-nil, zero value otherwise.

### GetDockerImageNameOk

`func (o *ReportMetadata) GetDockerImageNameOk() (*string, bool)`

GetDockerImageNameOk returns a tuple with the DockerImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageName

`func (o *ReportMetadata) SetDockerImageName(v string)`

SetDockerImageName sets DockerImageName field to given value.

### HasDockerImageName

`func (o *ReportMetadata) HasDockerImageName() bool`

HasDockerImageName returns a boolean if a field has been set.

### GetDockerImageNameWithTag

`func (o *ReportMetadata) GetDockerImageNameWithTag() string`

GetDockerImageNameWithTag returns the DockerImageNameWithTag field if non-nil, zero value otherwise.

### GetDockerImageNameWithTagOk

`func (o *ReportMetadata) GetDockerImageNameWithTagOk() (*string, bool)`

GetDockerImageNameWithTagOk returns a tuple with the DockerImageNameWithTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageNameWithTag

`func (o *ReportMetadata) SetDockerImageNameWithTag(v string)`

SetDockerImageNameWithTag sets DockerImageNameWithTag field to given value.

### HasDockerImageNameWithTag

`func (o *ReportMetadata) HasDockerImageNameWithTag() bool`

HasDockerImageNameWithTag returns a boolean if a field has been set.

### GetDockerImageSize

`func (o *ReportMetadata) GetDockerImageSize() string`

GetDockerImageSize returns the DockerImageSize field if non-nil, zero value otherwise.

### GetDockerImageSizeOk

`func (o *ReportMetadata) GetDockerImageSizeOk() (*string, bool)`

GetDockerImageSizeOk returns a tuple with the DockerImageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageSize

`func (o *ReportMetadata) SetDockerImageSize(v string)`

SetDockerImageSize sets DockerImageSize field to given value.

### HasDockerImageSize

`func (o *ReportMetadata) HasDockerImageSize() bool`

HasDockerImageSize returns a boolean if a field has been set.

### GetDockerImageTag

`func (o *ReportMetadata) GetDockerImageTag() string`

GetDockerImageTag returns the DockerImageTag field if non-nil, zero value otherwise.

### GetDockerImageTagOk

`func (o *ReportMetadata) GetDockerImageTagOk() (*string, bool)`

GetDockerImageTagOk returns a tuple with the DockerImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageTag

`func (o *ReportMetadata) SetDockerImageTag(v string)`

SetDockerImageTag sets DockerImageTag field to given value.

### HasDockerImageTag

`func (o *ReportMetadata) HasDockerImageTag() bool`

HasDockerImageTag returns a boolean if a field has been set.

### GetDockerImageVirtualSize

`func (o *ReportMetadata) GetDockerImageVirtualSize() string`

GetDockerImageVirtualSize returns the DockerImageVirtualSize field if non-nil, zero value otherwise.

### GetDockerImageVirtualSizeOk

`func (o *ReportMetadata) GetDockerImageVirtualSizeOk() (*string, bool)`

GetDockerImageVirtualSizeOk returns a tuple with the DockerImageVirtualSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageVirtualSize

`func (o *ReportMetadata) SetDockerImageVirtualSize(v string)`

SetDockerImageVirtualSize sets DockerImageVirtualSize field to given value.

### HasDockerImageVirtualSize

`func (o *ReportMetadata) HasDockerImageVirtualSize() bool`

HasDockerImageVirtualSize returns a boolean if a field has been set.

### GetDockerLabel

`func (o *ReportMetadata) GetDockerLabel() string`

GetDockerLabel returns the DockerLabel field if non-nil, zero value otherwise.

### GetDockerLabelOk

`func (o *ReportMetadata) GetDockerLabelOk() (*string, bool)`

GetDockerLabelOk returns a tuple with the DockerLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerLabel

`func (o *ReportMetadata) SetDockerLabel(v string)`

SetDockerLabel sets DockerLabel field to given value.

### HasDockerLabel

`func (o *ReportMetadata) HasDockerLabel() bool`

HasDockerLabel returns a boolean if a field has been set.

### GetHostName

`func (o *ReportMetadata) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ReportMetadata) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ReportMetadata) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *ReportMetadata) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetInstanceId

`func (o *ReportMetadata) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ReportMetadata) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ReportMetadata) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ReportMetadata) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInstanceType

`func (o *ReportMetadata) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ReportMetadata) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ReportMetadata) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *ReportMetadata) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetInterfaceIpMap

`func (o *ReportMetadata) GetInterfaceIpMap() string`

GetInterfaceIpMap returns the InterfaceIpMap field if non-nil, zero value otherwise.

### GetInterfaceIpMapOk

`func (o *ReportMetadata) GetInterfaceIpMapOk() (*string, bool)`

GetInterfaceIpMapOk returns a tuple with the InterfaceIpMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIpMap

`func (o *ReportMetadata) SetInterfaceIpMap(v string)`

SetInterfaceIpMap sets InterfaceIpMap field to given value.

### HasInterfaceIpMap

`func (o *ReportMetadata) HasInterfaceIpMap() bool`

HasInterfaceIpMap returns a boolean if a field has been set.

### GetInterfaceIps

`func (o *ReportMetadata) GetInterfaceIps() []string`

GetInterfaceIps returns the InterfaceIps field if non-nil, zero value otherwise.

### GetInterfaceIpsOk

`func (o *ReportMetadata) GetInterfaceIpsOk() (*[]string, bool)`

GetInterfaceIpsOk returns a tuple with the InterfaceIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIps

`func (o *ReportMetadata) SetInterfaceIps(v []string)`

SetInterfaceIps sets InterfaceIps field to given value.

### HasInterfaceIps

`func (o *ReportMetadata) HasInterfaceIps() bool`

HasInterfaceIps returns a boolean if a field has been set.

### GetInterfaceNames

`func (o *ReportMetadata) GetInterfaceNames() []string`

GetInterfaceNames returns the InterfaceNames field if non-nil, zero value otherwise.

### GetInterfaceNamesOk

`func (o *ReportMetadata) GetInterfaceNamesOk() (*[]string, bool)`

GetInterfaceNamesOk returns a tuple with the InterfaceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceNames

`func (o *ReportMetadata) SetInterfaceNames(v []string)`

SetInterfaceNames sets InterfaceNames field to given value.

### HasInterfaceNames

`func (o *ReportMetadata) HasInterfaceNames() bool`

HasInterfaceNames returns a boolean if a field has been set.

### GetIsConsoleVm

`func (o *ReportMetadata) GetIsConsoleVm() bool`

GetIsConsoleVm returns the IsConsoleVm field if non-nil, zero value otherwise.

### GetIsConsoleVmOk

`func (o *ReportMetadata) GetIsConsoleVmOk() (*bool, bool)`

GetIsConsoleVmOk returns a tuple with the IsConsoleVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConsoleVm

`func (o *ReportMetadata) SetIsConsoleVm(v bool)`

SetIsConsoleVm sets IsConsoleVm field to given value.

### HasIsConsoleVm

`func (o *ReportMetadata) HasIsConsoleVm() bool`

HasIsConsoleVm returns a boolean if a field has been set.

### GetKernelId

`func (o *ReportMetadata) GetKernelId() string`

GetKernelId returns the KernelId field if non-nil, zero value otherwise.

### GetKernelIdOk

`func (o *ReportMetadata) GetKernelIdOk() (*string, bool)`

GetKernelIdOk returns a tuple with the KernelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelId

`func (o *ReportMetadata) SetKernelId(v string)`

SetKernelId sets KernelId field to given value.

### HasKernelId

`func (o *ReportMetadata) HasKernelId() bool`

HasKernelId returns a boolean if a field has been set.

### GetKernelVersion

`func (o *ReportMetadata) GetKernelVersion() string`

GetKernelVersion returns the KernelVersion field if non-nil, zero value otherwise.

### GetKernelVersionOk

`func (o *ReportMetadata) GetKernelVersionOk() (*string, bool)`

GetKernelVersionOk returns a tuple with the KernelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelVersion

`func (o *ReportMetadata) SetKernelVersion(v string)`

SetKernelVersion sets KernelVersion field to given value.

### HasKernelVersion

`func (o *ReportMetadata) HasKernelVersion() bool`

HasKernelVersion returns a boolean if a field has been set.

### GetKubernetesClusterId

`func (o *ReportMetadata) GetKubernetesClusterId() string`

GetKubernetesClusterId returns the KubernetesClusterId field if non-nil, zero value otherwise.

### GetKubernetesClusterIdOk

`func (o *ReportMetadata) GetKubernetesClusterIdOk() (*string, bool)`

GetKubernetesClusterIdOk returns a tuple with the KubernetesClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesClusterId

`func (o *ReportMetadata) SetKubernetesClusterId(v string)`

SetKubernetesClusterId sets KubernetesClusterId field to given value.

### HasKubernetesClusterId

`func (o *ReportMetadata) HasKubernetesClusterId() bool`

HasKubernetesClusterId returns a boolean if a field has been set.

### GetKubernetesClusterName

`func (o *ReportMetadata) GetKubernetesClusterName() string`

GetKubernetesClusterName returns the KubernetesClusterName field if non-nil, zero value otherwise.

### GetKubernetesClusterNameOk

`func (o *ReportMetadata) GetKubernetesClusterNameOk() (*string, bool)`

GetKubernetesClusterNameOk returns a tuple with the KubernetesClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesClusterName

`func (o *ReportMetadata) SetKubernetesClusterName(v string)`

SetKubernetesClusterName sets KubernetesClusterName field to given value.

### HasKubernetesClusterName

`func (o *ReportMetadata) HasKubernetesClusterName() bool`

HasKubernetesClusterName returns a boolean if a field has been set.

### GetKubernetesCreated

`func (o *ReportMetadata) GetKubernetesCreated() string`

GetKubernetesCreated returns the KubernetesCreated field if non-nil, zero value otherwise.

### GetKubernetesCreatedOk

`func (o *ReportMetadata) GetKubernetesCreatedOk() (*string, bool)`

GetKubernetesCreatedOk returns a tuple with the KubernetesCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesCreated

`func (o *ReportMetadata) SetKubernetesCreated(v string)`

SetKubernetesCreated sets KubernetesCreated field to given value.

### HasKubernetesCreated

`func (o *ReportMetadata) HasKubernetesCreated() bool`

HasKubernetesCreated returns a boolean if a field has been set.

### GetKubernetesIngressIp

`func (o *ReportMetadata) GetKubernetesIngressIp() []string`

GetKubernetesIngressIp returns the KubernetesIngressIp field if non-nil, zero value otherwise.

### GetKubernetesIngressIpOk

`func (o *ReportMetadata) GetKubernetesIngressIpOk() (*[]string, bool)`

GetKubernetesIngressIpOk returns a tuple with the KubernetesIngressIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesIngressIp

`func (o *ReportMetadata) SetKubernetesIngressIp(v []string)`

SetKubernetesIngressIp sets KubernetesIngressIp field to given value.

### HasKubernetesIngressIp

`func (o *ReportMetadata) HasKubernetesIngressIp() bool`

HasKubernetesIngressIp returns a boolean if a field has been set.

### GetKubernetesIp

`func (o *ReportMetadata) GetKubernetesIp() string`

GetKubernetesIp returns the KubernetesIp field if non-nil, zero value otherwise.

### GetKubernetesIpOk

`func (o *ReportMetadata) GetKubernetesIpOk() (*string, bool)`

GetKubernetesIpOk returns a tuple with the KubernetesIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesIp

`func (o *ReportMetadata) SetKubernetesIp(v string)`

SetKubernetesIp sets KubernetesIp field to given value.

### HasKubernetesIp

`func (o *ReportMetadata) HasKubernetesIp() bool`

HasKubernetesIp returns a boolean if a field has been set.

### GetKubernetesIsInHostNetwork

`func (o *ReportMetadata) GetKubernetesIsInHostNetwork() bool`

GetKubernetesIsInHostNetwork returns the KubernetesIsInHostNetwork field if non-nil, zero value otherwise.

### GetKubernetesIsInHostNetworkOk

`func (o *ReportMetadata) GetKubernetesIsInHostNetworkOk() (*bool, bool)`

GetKubernetesIsInHostNetworkOk returns a tuple with the KubernetesIsInHostNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesIsInHostNetwork

`func (o *ReportMetadata) SetKubernetesIsInHostNetwork(v bool)`

SetKubernetesIsInHostNetwork sets KubernetesIsInHostNetwork field to given value.

### HasKubernetesIsInHostNetwork

`func (o *ReportMetadata) HasKubernetesIsInHostNetwork() bool`

HasKubernetesIsInHostNetwork returns a boolean if a field has been set.

### GetKubernetesLabels

`func (o *ReportMetadata) GetKubernetesLabels() string`

GetKubernetesLabels returns the KubernetesLabels field if non-nil, zero value otherwise.

### GetKubernetesLabelsOk

`func (o *ReportMetadata) GetKubernetesLabelsOk() (*string, bool)`

GetKubernetesLabelsOk returns a tuple with the KubernetesLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesLabels

`func (o *ReportMetadata) SetKubernetesLabels(v string)`

SetKubernetesLabels sets KubernetesLabels field to given value.

### HasKubernetesLabels

`func (o *ReportMetadata) HasKubernetesLabels() bool`

HasKubernetesLabels returns a boolean if a field has been set.

### GetKubernetesNamespace

`func (o *ReportMetadata) GetKubernetesNamespace() string`

GetKubernetesNamespace returns the KubernetesNamespace field if non-nil, zero value otherwise.

### GetKubernetesNamespaceOk

`func (o *ReportMetadata) GetKubernetesNamespaceOk() (*string, bool)`

GetKubernetesNamespaceOk returns a tuple with the KubernetesNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesNamespace

`func (o *ReportMetadata) SetKubernetesNamespace(v string)`

SetKubernetesNamespace sets KubernetesNamespace field to given value.

### HasKubernetesNamespace

`func (o *ReportMetadata) HasKubernetesNamespace() bool`

HasKubernetesNamespace returns a boolean if a field has been set.

### GetKubernetesPorts

`func (o *ReportMetadata) GetKubernetesPorts() []string`

GetKubernetesPorts returns the KubernetesPorts field if non-nil, zero value otherwise.

### GetKubernetesPortsOk

`func (o *ReportMetadata) GetKubernetesPortsOk() (*[]string, bool)`

GetKubernetesPortsOk returns a tuple with the KubernetesPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesPorts

`func (o *ReportMetadata) SetKubernetesPorts(v []string)`

SetKubernetesPorts sets KubernetesPorts field to given value.

### HasKubernetesPorts

`func (o *ReportMetadata) HasKubernetesPorts() bool`

HasKubernetesPorts returns a boolean if a field has been set.

### GetKubernetesPublicIp

`func (o *ReportMetadata) GetKubernetesPublicIp() string`

GetKubernetesPublicIp returns the KubernetesPublicIp field if non-nil, zero value otherwise.

### GetKubernetesPublicIpOk

`func (o *ReportMetadata) GetKubernetesPublicIpOk() (*string, bool)`

GetKubernetesPublicIpOk returns a tuple with the KubernetesPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesPublicIp

`func (o *ReportMetadata) SetKubernetesPublicIp(v string)`

SetKubernetesPublicIp sets KubernetesPublicIp field to given value.

### HasKubernetesPublicIp

`func (o *ReportMetadata) HasKubernetesPublicIp() bool`

HasKubernetesPublicIp returns a boolean if a field has been set.

### GetKubernetesState

`func (o *ReportMetadata) GetKubernetesState() string`

GetKubernetesState returns the KubernetesState field if non-nil, zero value otherwise.

### GetKubernetesStateOk

`func (o *ReportMetadata) GetKubernetesStateOk() (*string, bool)`

GetKubernetesStateOk returns a tuple with the KubernetesState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesState

`func (o *ReportMetadata) SetKubernetesState(v string)`

SetKubernetesState sets KubernetesState field to given value.

### HasKubernetesState

`func (o *ReportMetadata) HasKubernetesState() bool`

HasKubernetesState returns a boolean if a field has been set.

### GetKubernetesType

`func (o *ReportMetadata) GetKubernetesType() string`

GetKubernetesType returns the KubernetesType field if non-nil, zero value otherwise.

### GetKubernetesTypeOk

`func (o *ReportMetadata) GetKubernetesTypeOk() (*string, bool)`

GetKubernetesTypeOk returns a tuple with the KubernetesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesType

`func (o *ReportMetadata) SetKubernetesType(v string)`

SetKubernetesType sets KubernetesType field to given value.

### HasKubernetesType

`func (o *ReportMetadata) HasKubernetesType() bool`

HasKubernetesType returns a boolean if a field has been set.

### GetLocalCidr

`func (o *ReportMetadata) GetLocalCidr() []string`

GetLocalCidr returns the LocalCidr field if non-nil, zero value otherwise.

### GetLocalCidrOk

`func (o *ReportMetadata) GetLocalCidrOk() (*[]string, bool)`

GetLocalCidrOk returns a tuple with the LocalCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCidr

`func (o *ReportMetadata) SetLocalCidr(v []string)`

SetLocalCidr sets LocalCidr field to given value.

### HasLocalCidr

`func (o *ReportMetadata) HasLocalCidr() bool`

HasLocalCidr returns a boolean if a field has been set.

### GetLocalNetworks

`func (o *ReportMetadata) GetLocalNetworks() []string`

GetLocalNetworks returns the LocalNetworks field if non-nil, zero value otherwise.

### GetLocalNetworksOk

`func (o *ReportMetadata) GetLocalNetworksOk() (*[]string, bool)`

GetLocalNetworksOk returns a tuple with the LocalNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNetworks

`func (o *ReportMetadata) SetLocalNetworks(v []string)`

SetLocalNetworks sets LocalNetworks field to given value.

### HasLocalNetworks

`func (o *ReportMetadata) HasLocalNetworks() bool`

HasLocalNetworks returns a boolean if a field has been set.

### GetMemoryMax

`func (o *ReportMetadata) GetMemoryMax() int32`

GetMemoryMax returns the MemoryMax field if non-nil, zero value otherwise.

### GetMemoryMaxOk

`func (o *ReportMetadata) GetMemoryMaxOk() (*int32, bool)`

GetMemoryMaxOk returns a tuple with the MemoryMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMax

`func (o *ReportMetadata) SetMemoryMax(v int32)`

SetMemoryMax sets MemoryMax field to given value.

### HasMemoryMax

`func (o *ReportMetadata) HasMemoryMax() bool`

HasMemoryMax returns a boolean if a field has been set.

### GetMemoryUsage

`func (o *ReportMetadata) GetMemoryUsage() int32`

GetMemoryUsage returns the MemoryUsage field if non-nil, zero value otherwise.

### GetMemoryUsageOk

`func (o *ReportMetadata) GetMemoryUsageOk() (*int32, bool)`

GetMemoryUsageOk returns a tuple with the MemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsage

`func (o *ReportMetadata) SetMemoryUsage(v int32)`

SetMemoryUsage sets MemoryUsage field to given value.

### HasMemoryUsage

`func (o *ReportMetadata) HasMemoryUsage() bool`

HasMemoryUsage returns a boolean if a field has been set.

### GetNodeId

`func (o *ReportMetadata) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ReportMetadata) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ReportMetadata) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ReportMetadata) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNodeName

`func (o *ReportMetadata) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *ReportMetadata) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *ReportMetadata) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *ReportMetadata) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetNodeType

`func (o *ReportMetadata) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *ReportMetadata) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *ReportMetadata) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *ReportMetadata) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### GetOpenFiles

`func (o *ReportMetadata) GetOpenFiles() []string`

GetOpenFiles returns the OpenFiles field if non-nil, zero value otherwise.

### GetOpenFilesOk

`func (o *ReportMetadata) GetOpenFilesOk() (*[]string, bool)`

GetOpenFilesOk returns a tuple with the OpenFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenFiles

`func (o *ReportMetadata) SetOpenFiles(v []string)`

SetOpenFiles sets OpenFiles field to given value.

### HasOpenFiles

`func (o *ReportMetadata) HasOpenFiles() bool`

HasOpenFiles returns a boolean if a field has been set.

### GetOpenFilesCount

`func (o *ReportMetadata) GetOpenFilesCount() int32`

GetOpenFilesCount returns the OpenFilesCount field if non-nil, zero value otherwise.

### GetOpenFilesCountOk

`func (o *ReportMetadata) GetOpenFilesCountOk() (*int32, bool)`

GetOpenFilesCountOk returns a tuple with the OpenFilesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenFilesCount

`func (o *ReportMetadata) SetOpenFilesCount(v int32)`

SetOpenFilesCount sets OpenFilesCount field to given value.

### HasOpenFilesCount

`func (o *ReportMetadata) HasOpenFilesCount() bool`

HasOpenFilesCount returns a boolean if a field has been set.

### GetOs

`func (o *ReportMetadata) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *ReportMetadata) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *ReportMetadata) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *ReportMetadata) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetPid

`func (o *ReportMetadata) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *ReportMetadata) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *ReportMetadata) SetPid(v int32)`

SetPid sets Pid field to given value.

### HasPid

`func (o *ReportMetadata) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPodId

`func (o *ReportMetadata) GetPodId() string`

GetPodId returns the PodId field if non-nil, zero value otherwise.

### GetPodIdOk

`func (o *ReportMetadata) GetPodIdOk() (*string, bool)`

GetPodIdOk returns a tuple with the PodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodId

`func (o *ReportMetadata) SetPodId(v string)`

SetPodId sets PodId field to given value.

### HasPodId

`func (o *ReportMetadata) HasPodId() bool`

HasPodId returns a boolean if a field has been set.

### GetPodName

`func (o *ReportMetadata) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *ReportMetadata) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *ReportMetadata) SetPodName(v string)`

SetPodName sets PodName field to given value.

### HasPodName

`func (o *ReportMetadata) HasPodName() bool`

HasPodName returns a boolean if a field has been set.

### GetPpid

`func (o *ReportMetadata) GetPpid() int32`

GetPpid returns the Ppid field if non-nil, zero value otherwise.

### GetPpidOk

`func (o *ReportMetadata) GetPpidOk() (*int32, bool)`

GetPpidOk returns a tuple with the Ppid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpid

`func (o *ReportMetadata) SetPpid(v int32)`

SetPpid sets Ppid field to given value.

### HasPpid

`func (o *ReportMetadata) HasPpid() bool`

HasPpid returns a boolean if a field has been set.

### GetPrivateIp

`func (o *ReportMetadata) GetPrivateIp() []string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *ReportMetadata) GetPrivateIpOk() (*[]string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *ReportMetadata) SetPrivateIp(v []string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *ReportMetadata) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetPseudo

`func (o *ReportMetadata) GetPseudo() bool`

GetPseudo returns the Pseudo field if non-nil, zero value otherwise.

### GetPseudoOk

`func (o *ReportMetadata) GetPseudoOk() (*bool, bool)`

GetPseudoOk returns a tuple with the Pseudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPseudo

`func (o *ReportMetadata) SetPseudo(v bool)`

SetPseudo sets Pseudo field to given value.

### HasPseudo

`func (o *ReportMetadata) HasPseudo() bool`

HasPseudo returns a boolean if a field has been set.

### GetPublicIp

`func (o *ReportMetadata) GetPublicIp() []string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *ReportMetadata) GetPublicIpOk() (*[]string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *ReportMetadata) SetPublicIp(v []string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *ReportMetadata) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetResourceGroup

`func (o *ReportMetadata) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *ReportMetadata) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *ReportMetadata) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *ReportMetadata) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetShortName

`func (o *ReportMetadata) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *ReportMetadata) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *ReportMetadata) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *ReportMetadata) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetTags

`func (o *ReportMetadata) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ReportMetadata) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ReportMetadata) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ReportMetadata) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThreads

`func (o *ReportMetadata) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *ReportMetadata) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *ReportMetadata) SetThreads(v int32)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *ReportMetadata) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### GetTimestamp

`func (o *ReportMetadata) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ReportMetadata) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ReportMetadata) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ReportMetadata) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUptime

`func (o *ReportMetadata) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *ReportMetadata) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *ReportMetadata) SetUptime(v int32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *ReportMetadata) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetUserDefinedTags

`func (o *ReportMetadata) GetUserDefinedTags() []string`

GetUserDefinedTags returns the UserDefinedTags field if non-nil, zero value otherwise.

### GetUserDefinedTagsOk

`func (o *ReportMetadata) GetUserDefinedTagsOk() (*[]string, bool)`

GetUserDefinedTagsOk returns a tuple with the UserDefinedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedTags

`func (o *ReportMetadata) SetUserDefinedTags(v []string)`

SetUserDefinedTags sets UserDefinedTags field to given value.

### HasUserDefinedTags

`func (o *ReportMetadata) HasUserDefinedTags() bool`

HasUserDefinedTags returns a boolean if a field has been set.

### GetVersion

`func (o *ReportMetadata) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ReportMetadata) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ReportMetadata) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ReportMetadata) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


