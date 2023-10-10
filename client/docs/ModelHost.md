# ModelHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentRunning** | **bool** |  | 
**AvailabilityZone** | **string** |  | 
**CloudAccountId** | **string** |  | 
**CloudProvider** | **string** |  | 
**CloudRegion** | **string** |  | 
**ComplianceLatestScanId** | **string** |  | 
**ComplianceScanStatus** | **string** |  | 
**CompliancesCount** | **int32** |  | 
**ContainerImages** | [**[]ModelContainerImage**](ModelContainerImage.md) |  | 
**Containers** | [**[]ModelContainer**](ModelContainer.md) |  | 
**CpuMax** | **float32** |  | 
**CpuUsage** | **float32** |  | 
**HostName** | **string** |  | 
**InboundConnections** | [**[]ModelConnection**](ModelConnection.md) |  | 
**InstanceId** | **string** |  | 
**InstanceType** | **string** |  | 
**IsConsoleVm** | **bool** |  | 
**KernelId** | **string** |  | 
**KernelVersion** | **string** |  | 
**LocalCidr** | **[]interface{}** |  | 
**LocalNetworks** | **[]interface{}** |  | 
**MalwareLatestScanId** | **string** |  | 
**MalwareScanStatus** | **string** |  | 
**MalwaresCount** | **int32** |  | 
**MemoryMax** | **int32** |  | 
**MemoryUsage** | **int32** |  | 
**NodeId** | **string** |  | 
**NodeName** | **string** |  | 
**Os** | **string** |  | 
**OutboundConnections** | [**[]ModelConnection**](ModelConnection.md) |  | 
**Pods** | [**[]ModelPod**](ModelPod.md) |  | 
**PrivateIp** | **[]interface{}** |  | 
**Processes** | [**[]ModelProcess**](ModelProcess.md) |  | 
**PublicIp** | **[]interface{}** |  | 
**ResourceGroup** | **string** |  | 
**SecretLatestScanId** | **string** |  | 
**SecretScanStatus** | **string** |  | 
**SecretsCount** | **int32** |  | 
**Uptime** | **int32** |  | 
**Version** | **string** |  | 
**VulnerabilitiesCount** | **int32** |  | 
**VulnerabilityLatestScanId** | **string** |  | 
**VulnerabilityScanStatus** | **string** |  | 

## Methods

### NewModelHost

`func NewModelHost(agentRunning bool, availabilityZone string, cloudAccountId string, cloudProvider string, cloudRegion string, complianceLatestScanId string, complianceScanStatus string, compliancesCount int32, containerImages []ModelContainerImage, containers []ModelContainer, cpuMax float32, cpuUsage float32, hostName string, inboundConnections []ModelConnection, instanceId string, instanceType string, isConsoleVm bool, kernelId string, kernelVersion string, localCidr []interface{}, localNetworks []interface{}, malwareLatestScanId string, malwareScanStatus string, malwaresCount int32, memoryMax int32, memoryUsage int32, nodeId string, nodeName string, os string, outboundConnections []ModelConnection, pods []ModelPod, privateIp []interface{}, processes []ModelProcess, publicIp []interface{}, resourceGroup string, secretLatestScanId string, secretScanStatus string, secretsCount int32, uptime int32, version string, vulnerabilitiesCount int32, vulnerabilityLatestScanId string, vulnerabilityScanStatus string, ) *ModelHost`

NewModelHost instantiates a new ModelHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelHostWithDefaults

`func NewModelHostWithDefaults() *ModelHost`

NewModelHostWithDefaults instantiates a new ModelHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentRunning

`func (o *ModelHost) GetAgentRunning() bool`

GetAgentRunning returns the AgentRunning field if non-nil, zero value otherwise.

### GetAgentRunningOk

`func (o *ModelHost) GetAgentRunningOk() (*bool, bool)`

GetAgentRunningOk returns a tuple with the AgentRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentRunning

`func (o *ModelHost) SetAgentRunning(v bool)`

SetAgentRunning sets AgentRunning field to given value.


### GetAvailabilityZone

`func (o *ModelHost) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *ModelHost) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *ModelHost) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetCloudAccountId

`func (o *ModelHost) GetCloudAccountId() string`

GetCloudAccountId returns the CloudAccountId field if non-nil, zero value otherwise.

### GetCloudAccountIdOk

`func (o *ModelHost) GetCloudAccountIdOk() (*string, bool)`

GetCloudAccountIdOk returns a tuple with the CloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccountId

`func (o *ModelHost) SetCloudAccountId(v string)`

SetCloudAccountId sets CloudAccountId field to given value.


### GetCloudProvider

`func (o *ModelHost) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ModelHost) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ModelHost) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetCloudRegion

`func (o *ModelHost) GetCloudRegion() string`

GetCloudRegion returns the CloudRegion field if non-nil, zero value otherwise.

### GetCloudRegionOk

`func (o *ModelHost) GetCloudRegionOk() (*string, bool)`

GetCloudRegionOk returns a tuple with the CloudRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudRegion

`func (o *ModelHost) SetCloudRegion(v string)`

SetCloudRegion sets CloudRegion field to given value.


### GetComplianceLatestScanId

`func (o *ModelHost) GetComplianceLatestScanId() string`

GetComplianceLatestScanId returns the ComplianceLatestScanId field if non-nil, zero value otherwise.

### GetComplianceLatestScanIdOk

`func (o *ModelHost) GetComplianceLatestScanIdOk() (*string, bool)`

GetComplianceLatestScanIdOk returns a tuple with the ComplianceLatestScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceLatestScanId

`func (o *ModelHost) SetComplianceLatestScanId(v string)`

SetComplianceLatestScanId sets ComplianceLatestScanId field to given value.


### GetComplianceScanStatus

`func (o *ModelHost) GetComplianceScanStatus() string`

GetComplianceScanStatus returns the ComplianceScanStatus field if non-nil, zero value otherwise.

### GetComplianceScanStatusOk

`func (o *ModelHost) GetComplianceScanStatusOk() (*string, bool)`

GetComplianceScanStatusOk returns a tuple with the ComplianceScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceScanStatus

`func (o *ModelHost) SetComplianceScanStatus(v string)`

SetComplianceScanStatus sets ComplianceScanStatus field to given value.


### GetCompliancesCount

`func (o *ModelHost) GetCompliancesCount() int32`

GetCompliancesCount returns the CompliancesCount field if non-nil, zero value otherwise.

### GetCompliancesCountOk

`func (o *ModelHost) GetCompliancesCountOk() (*int32, bool)`

GetCompliancesCountOk returns a tuple with the CompliancesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliancesCount

`func (o *ModelHost) SetCompliancesCount(v int32)`

SetCompliancesCount sets CompliancesCount field to given value.


### GetContainerImages

`func (o *ModelHost) GetContainerImages() []ModelContainerImage`

GetContainerImages returns the ContainerImages field if non-nil, zero value otherwise.

### GetContainerImagesOk

`func (o *ModelHost) GetContainerImagesOk() (*[]ModelContainerImage, bool)`

GetContainerImagesOk returns a tuple with the ContainerImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerImages

`func (o *ModelHost) SetContainerImages(v []ModelContainerImage)`

SetContainerImages sets ContainerImages field to given value.


### SetContainerImagesNil

`func (o *ModelHost) SetContainerImagesNil(b bool)`

 SetContainerImagesNil sets the value for ContainerImages to be an explicit nil

### UnsetContainerImages
`func (o *ModelHost) UnsetContainerImages()`

UnsetContainerImages ensures that no value is present for ContainerImages, not even an explicit nil
### GetContainers

`func (o *ModelHost) GetContainers() []ModelContainer`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *ModelHost) GetContainersOk() (*[]ModelContainer, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *ModelHost) SetContainers(v []ModelContainer)`

SetContainers sets Containers field to given value.


### SetContainersNil

`func (o *ModelHost) SetContainersNil(b bool)`

 SetContainersNil sets the value for Containers to be an explicit nil

### UnsetContainers
`func (o *ModelHost) UnsetContainers()`

UnsetContainers ensures that no value is present for Containers, not even an explicit nil
### GetCpuMax

`func (o *ModelHost) GetCpuMax() float32`

GetCpuMax returns the CpuMax field if non-nil, zero value otherwise.

### GetCpuMaxOk

`func (o *ModelHost) GetCpuMaxOk() (*float32, bool)`

GetCpuMaxOk returns a tuple with the CpuMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMax

`func (o *ModelHost) SetCpuMax(v float32)`

SetCpuMax sets CpuMax field to given value.


### GetCpuUsage

`func (o *ModelHost) GetCpuUsage() float32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *ModelHost) GetCpuUsageOk() (*float32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *ModelHost) SetCpuUsage(v float32)`

SetCpuUsage sets CpuUsage field to given value.


### GetHostName

`func (o *ModelHost) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ModelHost) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ModelHost) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetInboundConnections

`func (o *ModelHost) GetInboundConnections() []ModelConnection`

GetInboundConnections returns the InboundConnections field if non-nil, zero value otherwise.

### GetInboundConnectionsOk

`func (o *ModelHost) GetInboundConnectionsOk() (*[]ModelConnection, bool)`

GetInboundConnectionsOk returns a tuple with the InboundConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundConnections

`func (o *ModelHost) SetInboundConnections(v []ModelConnection)`

SetInboundConnections sets InboundConnections field to given value.


### SetInboundConnectionsNil

`func (o *ModelHost) SetInboundConnectionsNil(b bool)`

 SetInboundConnectionsNil sets the value for InboundConnections to be an explicit nil

### UnsetInboundConnections
`func (o *ModelHost) UnsetInboundConnections()`

UnsetInboundConnections ensures that no value is present for InboundConnections, not even an explicit nil
### GetInstanceId

`func (o *ModelHost) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ModelHost) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ModelHost) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetInstanceType

`func (o *ModelHost) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ModelHost) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ModelHost) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.


### GetIsConsoleVm

`func (o *ModelHost) GetIsConsoleVm() bool`

GetIsConsoleVm returns the IsConsoleVm field if non-nil, zero value otherwise.

### GetIsConsoleVmOk

`func (o *ModelHost) GetIsConsoleVmOk() (*bool, bool)`

GetIsConsoleVmOk returns a tuple with the IsConsoleVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConsoleVm

`func (o *ModelHost) SetIsConsoleVm(v bool)`

SetIsConsoleVm sets IsConsoleVm field to given value.


### GetKernelId

`func (o *ModelHost) GetKernelId() string`

GetKernelId returns the KernelId field if non-nil, zero value otherwise.

### GetKernelIdOk

`func (o *ModelHost) GetKernelIdOk() (*string, bool)`

GetKernelIdOk returns a tuple with the KernelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelId

`func (o *ModelHost) SetKernelId(v string)`

SetKernelId sets KernelId field to given value.


### GetKernelVersion

`func (o *ModelHost) GetKernelVersion() string`

GetKernelVersion returns the KernelVersion field if non-nil, zero value otherwise.

### GetKernelVersionOk

`func (o *ModelHost) GetKernelVersionOk() (*string, bool)`

GetKernelVersionOk returns a tuple with the KernelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelVersion

`func (o *ModelHost) SetKernelVersion(v string)`

SetKernelVersion sets KernelVersion field to given value.


### GetLocalCidr

`func (o *ModelHost) GetLocalCidr() []interface{}`

GetLocalCidr returns the LocalCidr field if non-nil, zero value otherwise.

### GetLocalCidrOk

`func (o *ModelHost) GetLocalCidrOk() (*[]interface{}, bool)`

GetLocalCidrOk returns a tuple with the LocalCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCidr

`func (o *ModelHost) SetLocalCidr(v []interface{})`

SetLocalCidr sets LocalCidr field to given value.


### SetLocalCidrNil

`func (o *ModelHost) SetLocalCidrNil(b bool)`

 SetLocalCidrNil sets the value for LocalCidr to be an explicit nil

### UnsetLocalCidr
`func (o *ModelHost) UnsetLocalCidr()`

UnsetLocalCidr ensures that no value is present for LocalCidr, not even an explicit nil
### GetLocalNetworks

`func (o *ModelHost) GetLocalNetworks() []interface{}`

GetLocalNetworks returns the LocalNetworks field if non-nil, zero value otherwise.

### GetLocalNetworksOk

`func (o *ModelHost) GetLocalNetworksOk() (*[]interface{}, bool)`

GetLocalNetworksOk returns a tuple with the LocalNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNetworks

`func (o *ModelHost) SetLocalNetworks(v []interface{})`

SetLocalNetworks sets LocalNetworks field to given value.


### SetLocalNetworksNil

`func (o *ModelHost) SetLocalNetworksNil(b bool)`

 SetLocalNetworksNil sets the value for LocalNetworks to be an explicit nil

### UnsetLocalNetworks
`func (o *ModelHost) UnsetLocalNetworks()`

UnsetLocalNetworks ensures that no value is present for LocalNetworks, not even an explicit nil
### GetMalwareLatestScanId

`func (o *ModelHost) GetMalwareLatestScanId() string`

GetMalwareLatestScanId returns the MalwareLatestScanId field if non-nil, zero value otherwise.

### GetMalwareLatestScanIdOk

`func (o *ModelHost) GetMalwareLatestScanIdOk() (*string, bool)`

GetMalwareLatestScanIdOk returns a tuple with the MalwareLatestScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalwareLatestScanId

`func (o *ModelHost) SetMalwareLatestScanId(v string)`

SetMalwareLatestScanId sets MalwareLatestScanId field to given value.


### GetMalwareScanStatus

`func (o *ModelHost) GetMalwareScanStatus() string`

GetMalwareScanStatus returns the MalwareScanStatus field if non-nil, zero value otherwise.

### GetMalwareScanStatusOk

`func (o *ModelHost) GetMalwareScanStatusOk() (*string, bool)`

GetMalwareScanStatusOk returns a tuple with the MalwareScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalwareScanStatus

`func (o *ModelHost) SetMalwareScanStatus(v string)`

SetMalwareScanStatus sets MalwareScanStatus field to given value.


### GetMalwaresCount

`func (o *ModelHost) GetMalwaresCount() int32`

GetMalwaresCount returns the MalwaresCount field if non-nil, zero value otherwise.

### GetMalwaresCountOk

`func (o *ModelHost) GetMalwaresCountOk() (*int32, bool)`

GetMalwaresCountOk returns a tuple with the MalwaresCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalwaresCount

`func (o *ModelHost) SetMalwaresCount(v int32)`

SetMalwaresCount sets MalwaresCount field to given value.


### GetMemoryMax

`func (o *ModelHost) GetMemoryMax() int32`

GetMemoryMax returns the MemoryMax field if non-nil, zero value otherwise.

### GetMemoryMaxOk

`func (o *ModelHost) GetMemoryMaxOk() (*int32, bool)`

GetMemoryMaxOk returns a tuple with the MemoryMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMax

`func (o *ModelHost) SetMemoryMax(v int32)`

SetMemoryMax sets MemoryMax field to given value.


### GetMemoryUsage

`func (o *ModelHost) GetMemoryUsage() int32`

GetMemoryUsage returns the MemoryUsage field if non-nil, zero value otherwise.

### GetMemoryUsageOk

`func (o *ModelHost) GetMemoryUsageOk() (*int32, bool)`

GetMemoryUsageOk returns a tuple with the MemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsage

`func (o *ModelHost) SetMemoryUsage(v int32)`

SetMemoryUsage sets MemoryUsage field to given value.


### GetNodeId

`func (o *ModelHost) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelHost) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelHost) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNodeName

`func (o *ModelHost) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *ModelHost) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *ModelHost) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.


### GetOs

`func (o *ModelHost) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *ModelHost) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *ModelHost) SetOs(v string)`

SetOs sets Os field to given value.


### GetOutboundConnections

`func (o *ModelHost) GetOutboundConnections() []ModelConnection`

GetOutboundConnections returns the OutboundConnections field if non-nil, zero value otherwise.

### GetOutboundConnectionsOk

`func (o *ModelHost) GetOutboundConnectionsOk() (*[]ModelConnection, bool)`

GetOutboundConnectionsOk returns a tuple with the OutboundConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundConnections

`func (o *ModelHost) SetOutboundConnections(v []ModelConnection)`

SetOutboundConnections sets OutboundConnections field to given value.


### SetOutboundConnectionsNil

`func (o *ModelHost) SetOutboundConnectionsNil(b bool)`

 SetOutboundConnectionsNil sets the value for OutboundConnections to be an explicit nil

### UnsetOutboundConnections
`func (o *ModelHost) UnsetOutboundConnections()`

UnsetOutboundConnections ensures that no value is present for OutboundConnections, not even an explicit nil
### GetPods

`func (o *ModelHost) GetPods() []ModelPod`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *ModelHost) GetPodsOk() (*[]ModelPod, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *ModelHost) SetPods(v []ModelPod)`

SetPods sets Pods field to given value.


### SetPodsNil

`func (o *ModelHost) SetPodsNil(b bool)`

 SetPodsNil sets the value for Pods to be an explicit nil

### UnsetPods
`func (o *ModelHost) UnsetPods()`

UnsetPods ensures that no value is present for Pods, not even an explicit nil
### GetPrivateIp

`func (o *ModelHost) GetPrivateIp() []interface{}`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *ModelHost) GetPrivateIpOk() (*[]interface{}, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *ModelHost) SetPrivateIp(v []interface{})`

SetPrivateIp sets PrivateIp field to given value.


### SetPrivateIpNil

`func (o *ModelHost) SetPrivateIpNil(b bool)`

 SetPrivateIpNil sets the value for PrivateIp to be an explicit nil

### UnsetPrivateIp
`func (o *ModelHost) UnsetPrivateIp()`

UnsetPrivateIp ensures that no value is present for PrivateIp, not even an explicit nil
### GetProcesses

`func (o *ModelHost) GetProcesses() []ModelProcess`

GetProcesses returns the Processes field if non-nil, zero value otherwise.

### GetProcessesOk

`func (o *ModelHost) GetProcessesOk() (*[]ModelProcess, bool)`

GetProcessesOk returns a tuple with the Processes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcesses

`func (o *ModelHost) SetProcesses(v []ModelProcess)`

SetProcesses sets Processes field to given value.


### SetProcessesNil

`func (o *ModelHost) SetProcessesNil(b bool)`

 SetProcessesNil sets the value for Processes to be an explicit nil

### UnsetProcesses
`func (o *ModelHost) UnsetProcesses()`

UnsetProcesses ensures that no value is present for Processes, not even an explicit nil
### GetPublicIp

`func (o *ModelHost) GetPublicIp() []interface{}`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *ModelHost) GetPublicIpOk() (*[]interface{}, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *ModelHost) SetPublicIp(v []interface{})`

SetPublicIp sets PublicIp field to given value.


### SetPublicIpNil

`func (o *ModelHost) SetPublicIpNil(b bool)`

 SetPublicIpNil sets the value for PublicIp to be an explicit nil

### UnsetPublicIp
`func (o *ModelHost) UnsetPublicIp()`

UnsetPublicIp ensures that no value is present for PublicIp, not even an explicit nil
### GetResourceGroup

`func (o *ModelHost) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *ModelHost) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *ModelHost) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.


### GetSecretLatestScanId

`func (o *ModelHost) GetSecretLatestScanId() string`

GetSecretLatestScanId returns the SecretLatestScanId field if non-nil, zero value otherwise.

### GetSecretLatestScanIdOk

`func (o *ModelHost) GetSecretLatestScanIdOk() (*string, bool)`

GetSecretLatestScanIdOk returns a tuple with the SecretLatestScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretLatestScanId

`func (o *ModelHost) SetSecretLatestScanId(v string)`

SetSecretLatestScanId sets SecretLatestScanId field to given value.


### GetSecretScanStatus

`func (o *ModelHost) GetSecretScanStatus() string`

GetSecretScanStatus returns the SecretScanStatus field if non-nil, zero value otherwise.

### GetSecretScanStatusOk

`func (o *ModelHost) GetSecretScanStatusOk() (*string, bool)`

GetSecretScanStatusOk returns a tuple with the SecretScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretScanStatus

`func (o *ModelHost) SetSecretScanStatus(v string)`

SetSecretScanStatus sets SecretScanStatus field to given value.


### GetSecretsCount

`func (o *ModelHost) GetSecretsCount() int32`

GetSecretsCount returns the SecretsCount field if non-nil, zero value otherwise.

### GetSecretsCountOk

`func (o *ModelHost) GetSecretsCountOk() (*int32, bool)`

GetSecretsCountOk returns a tuple with the SecretsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsCount

`func (o *ModelHost) SetSecretsCount(v int32)`

SetSecretsCount sets SecretsCount field to given value.


### GetUptime

`func (o *ModelHost) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *ModelHost) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *ModelHost) SetUptime(v int32)`

SetUptime sets Uptime field to given value.


### GetVersion

`func (o *ModelHost) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelHost) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelHost) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVulnerabilitiesCount

`func (o *ModelHost) GetVulnerabilitiesCount() int32`

GetVulnerabilitiesCount returns the VulnerabilitiesCount field if non-nil, zero value otherwise.

### GetVulnerabilitiesCountOk

`func (o *ModelHost) GetVulnerabilitiesCountOk() (*int32, bool)`

GetVulnerabilitiesCountOk returns a tuple with the VulnerabilitiesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilitiesCount

`func (o *ModelHost) SetVulnerabilitiesCount(v int32)`

SetVulnerabilitiesCount sets VulnerabilitiesCount field to given value.


### GetVulnerabilityLatestScanId

`func (o *ModelHost) GetVulnerabilityLatestScanId() string`

GetVulnerabilityLatestScanId returns the VulnerabilityLatestScanId field if non-nil, zero value otherwise.

### GetVulnerabilityLatestScanIdOk

`func (o *ModelHost) GetVulnerabilityLatestScanIdOk() (*string, bool)`

GetVulnerabilityLatestScanIdOk returns a tuple with the VulnerabilityLatestScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityLatestScanId

`func (o *ModelHost) SetVulnerabilityLatestScanId(v string)`

SetVulnerabilityLatestScanId sets VulnerabilityLatestScanId field to given value.


### GetVulnerabilityScanStatus

`func (o *ModelHost) GetVulnerabilityScanStatus() string`

GetVulnerabilityScanStatus returns the VulnerabilityScanStatus field if non-nil, zero value otherwise.

### GetVulnerabilityScanStatusOk

`func (o *ModelHost) GetVulnerabilityScanStatusOk() (*string, bool)`

GetVulnerabilityScanStatusOk returns a tuple with the VulnerabilityScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityScanStatus

`func (o *ModelHost) SetVulnerabilityScanStatus(v string)`

SetVulnerabilityScanStatus sets VulnerabilityScanStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


