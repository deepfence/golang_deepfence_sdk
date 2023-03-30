# ModelHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentRunning** | **bool** |  | 
**AvailabilityZone** | Pointer to **string** |  | [optional] 
**CloudComplianceLatestScanId** | **string** |  | 
**CloudComplianceScanStatus** | **string** |  | 
**CloudCompliancesCount** | **int32** |  | 
**CloudProvider** | **string** |  | 
**CloudRegion** | **string** |  | 
**ComplianceLatestScanId** | **string** |  | 
**ComplianceScanStatus** | **string** |  | 
**CompliancesCount** | **int32** |  | 
**ContainerImages** | [**[]ModelContainerImage**](ModelContainerImage.md) |  | 
**Containers** | [**[]ModelContainer**](ModelContainer.md) |  | 
**CpuMax** | Pointer to **float32** |  | [optional] 
**CpuUsage** | Pointer to **float32** |  | [optional] 
**HostName** | **string** |  | 
**InstanceId** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**InterfaceIps** | **[]string** |  | 
**InterfaceNames** | **[]string** |  | 
**IsConsoleVm** | **bool** |  | 
**KernelId** | Pointer to **string** |  | [optional] 
**KernelVersion** | **string** |  | 
**LocalCidr** | **[]string** |  | 
**LocalNetworks** | Pointer to **[]string** |  | [optional] 
**MalwareLatestScanId** | **string** |  | 
**MalwareScanStatus** | **string** |  | 
**MalwaresCount** | **int32** |  | 
**MemoryMax** | Pointer to **int32** |  | [optional] 
**MemoryUsage** | Pointer to **int32** |  | [optional] 
**NodeId** | **string** |  | 
**NodeName** | **string** |  | 
**Os** | **string** |  | 
**Pods** | [**[]ModelPod**](ModelPod.md) |  | 
**PrivateIp** | Pointer to **[]string** |  | [optional] 
**Processes** | [**[]ModelProcess**](ModelProcess.md) |  | 
**PublicIp** | Pointer to **[]string** |  | [optional] 
**ResourceGroup** | Pointer to **string** |  | [optional] 
**SecretLatestScan** | **string** |  | 
**SecretScanStatus** | **string** |  | 
**SecretsCount** | **int32** |  | 
**Uptime** | **int32** |  | 
**Version** | **string** |  | 
**VulnerabilitiesCount** | **int32** |  | 
**VulnerabilityLatestScanId** | **string** |  | 
**VulnerabilityScanStatus** | **string** |  | 

## Methods

### NewModelHost

`func NewModelHost(agentRunning bool, cloudComplianceLatestScanId string, cloudComplianceScanStatus string, cloudCompliancesCount int32, cloudProvider string, cloudRegion string, complianceLatestScanId string, complianceScanStatus string, compliancesCount int32, containerImages []ModelContainerImage, containers []ModelContainer, hostName string, interfaceIps []string, interfaceNames []string, isConsoleVm bool, kernelVersion string, localCidr []string, malwareLatestScanId string, malwareScanStatus string, malwaresCount int32, nodeId string, nodeName string, os string, pods []ModelPod, processes []ModelProcess, secretLatestScan string, secretScanStatus string, secretsCount int32, uptime int32, version string, vulnerabilitiesCount int32, vulnerabilityLatestScanId string, vulnerabilityScanStatus string, ) *ModelHost`

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

### HasAvailabilityZone

`func (o *ModelHost) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetCloudComplianceLatestScanId

`func (o *ModelHost) GetCloudComplianceLatestScanId() string`

GetCloudComplianceLatestScanId returns the CloudComplianceLatestScanId field if non-nil, zero value otherwise.

### GetCloudComplianceLatestScanIdOk

`func (o *ModelHost) GetCloudComplianceLatestScanIdOk() (*string, bool)`

GetCloudComplianceLatestScanIdOk returns a tuple with the CloudComplianceLatestScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudComplianceLatestScanId

`func (o *ModelHost) SetCloudComplianceLatestScanId(v string)`

SetCloudComplianceLatestScanId sets CloudComplianceLatestScanId field to given value.


### GetCloudComplianceScanStatus

`func (o *ModelHost) GetCloudComplianceScanStatus() string`

GetCloudComplianceScanStatus returns the CloudComplianceScanStatus field if non-nil, zero value otherwise.

### GetCloudComplianceScanStatusOk

`func (o *ModelHost) GetCloudComplianceScanStatusOk() (*string, bool)`

GetCloudComplianceScanStatusOk returns a tuple with the CloudComplianceScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudComplianceScanStatus

`func (o *ModelHost) SetCloudComplianceScanStatus(v string)`

SetCloudComplianceScanStatus sets CloudComplianceScanStatus field to given value.


### GetCloudCompliancesCount

`func (o *ModelHost) GetCloudCompliancesCount() int32`

GetCloudCompliancesCount returns the CloudCompliancesCount field if non-nil, zero value otherwise.

### GetCloudCompliancesCountOk

`func (o *ModelHost) GetCloudCompliancesCountOk() (*int32, bool)`

GetCloudCompliancesCountOk returns a tuple with the CloudCompliancesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCompliancesCount

`func (o *ModelHost) SetCloudCompliancesCount(v int32)`

SetCloudCompliancesCount sets CloudCompliancesCount field to given value.


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

### HasCpuMax

`func (o *ModelHost) HasCpuMax() bool`

HasCpuMax returns a boolean if a field has been set.

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

### HasCpuUsage

`func (o *ModelHost) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

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

### HasInstanceId

`func (o *ModelHost) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

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

### HasInstanceType

`func (o *ModelHost) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetInterfaceIps

`func (o *ModelHost) GetInterfaceIps() []string`

GetInterfaceIps returns the InterfaceIps field if non-nil, zero value otherwise.

### GetInterfaceIpsOk

`func (o *ModelHost) GetInterfaceIpsOk() (*[]string, bool)`

GetInterfaceIpsOk returns a tuple with the InterfaceIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIps

`func (o *ModelHost) SetInterfaceIps(v []string)`

SetInterfaceIps sets InterfaceIps field to given value.


### SetInterfaceIpsNil

`func (o *ModelHost) SetInterfaceIpsNil(b bool)`

 SetInterfaceIpsNil sets the value for InterfaceIps to be an explicit nil

### UnsetInterfaceIps
`func (o *ModelHost) UnsetInterfaceIps()`

UnsetInterfaceIps ensures that no value is present for InterfaceIps, not even an explicit nil
### GetInterfaceNames

`func (o *ModelHost) GetInterfaceNames() []string`

GetInterfaceNames returns the InterfaceNames field if non-nil, zero value otherwise.

### GetInterfaceNamesOk

`func (o *ModelHost) GetInterfaceNamesOk() (*[]string, bool)`

GetInterfaceNamesOk returns a tuple with the InterfaceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceNames

`func (o *ModelHost) SetInterfaceNames(v []string)`

SetInterfaceNames sets InterfaceNames field to given value.


### SetInterfaceNamesNil

`func (o *ModelHost) SetInterfaceNamesNil(b bool)`

 SetInterfaceNamesNil sets the value for InterfaceNames to be an explicit nil

### UnsetInterfaceNames
`func (o *ModelHost) UnsetInterfaceNames()`

UnsetInterfaceNames ensures that no value is present for InterfaceNames, not even an explicit nil
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

### HasKernelId

`func (o *ModelHost) HasKernelId() bool`

HasKernelId returns a boolean if a field has been set.

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

`func (o *ModelHost) GetLocalCidr() []string`

GetLocalCidr returns the LocalCidr field if non-nil, zero value otherwise.

### GetLocalCidrOk

`func (o *ModelHost) GetLocalCidrOk() (*[]string, bool)`

GetLocalCidrOk returns a tuple with the LocalCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCidr

`func (o *ModelHost) SetLocalCidr(v []string)`

SetLocalCidr sets LocalCidr field to given value.


### SetLocalCidrNil

`func (o *ModelHost) SetLocalCidrNil(b bool)`

 SetLocalCidrNil sets the value for LocalCidr to be an explicit nil

### UnsetLocalCidr
`func (o *ModelHost) UnsetLocalCidr()`

UnsetLocalCidr ensures that no value is present for LocalCidr, not even an explicit nil
### GetLocalNetworks

`func (o *ModelHost) GetLocalNetworks() []string`

GetLocalNetworks returns the LocalNetworks field if non-nil, zero value otherwise.

### GetLocalNetworksOk

`func (o *ModelHost) GetLocalNetworksOk() (*[]string, bool)`

GetLocalNetworksOk returns a tuple with the LocalNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNetworks

`func (o *ModelHost) SetLocalNetworks(v []string)`

SetLocalNetworks sets LocalNetworks field to given value.

### HasLocalNetworks

`func (o *ModelHost) HasLocalNetworks() bool`

HasLocalNetworks returns a boolean if a field has been set.

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

### HasMemoryMax

`func (o *ModelHost) HasMemoryMax() bool`

HasMemoryMax returns a boolean if a field has been set.

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

### HasMemoryUsage

`func (o *ModelHost) HasMemoryUsage() bool`

HasMemoryUsage returns a boolean if a field has been set.

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

`func (o *ModelHost) GetPrivateIp() []string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *ModelHost) GetPrivateIpOk() (*[]string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *ModelHost) SetPrivateIp(v []string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *ModelHost) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

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

`func (o *ModelHost) GetPublicIp() []string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *ModelHost) GetPublicIpOk() (*[]string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *ModelHost) SetPublicIp(v []string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *ModelHost) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

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

### HasResourceGroup

`func (o *ModelHost) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetSecretLatestScan

`func (o *ModelHost) GetSecretLatestScan() string`

GetSecretLatestScan returns the SecretLatestScan field if non-nil, zero value otherwise.

### GetSecretLatestScanOk

`func (o *ModelHost) GetSecretLatestScanOk() (*string, bool)`

GetSecretLatestScanOk returns a tuple with the SecretLatestScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretLatestScan

`func (o *ModelHost) SetSecretLatestScan(v string)`

SetSecretLatestScan sets SecretLatestScan field to given value.


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


