# ModelContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuMax** | **float32** |  | 
**CpuUsage** | **float32** |  | 
**DockerContainerCommand** | **string** |  | 
**DockerContainerCreated** | **string** |  | 
**DockerContainerIps** | **[]interface{}** |  | 
**DockerContainerName** | **string** |  | 
**DockerContainerNetworkMode** | **string** |  | 
**DockerContainerNetworks** | **string** |  | 
**DockerContainerPorts** | **string** |  | 
**DockerContainerState** | **string** |  | 
**DockerContainerStateHuman** | **string** |  | 
**DockerLabels** | **map[string]interface{}** |  | 
**HostName** | **string** |  | 
**Image** | [**ModelContainerImage**](ModelContainerImage.md) |  | 
**MalwareLatestScanId** | **string** |  | 
**MalwareScanStatus** | **string** |  | 
**MalwaresCount** | **int32** |  | 
**MemoryMax** | **int32** |  | 
**MemoryUsage** | **int32** |  | 
**NodeId** | **string** |  | 
**NodeName** | **string** |  | 
**Processes** | [**[]ModelProcess**](ModelProcess.md) |  | 
**SecretLatestScanId** | **string** |  | 
**SecretScanStatus** | **string** |  | 
**SecretsCount** | **int32** |  | 
**Uptime** | **int32** |  | 
**VulnerabilitiesCount** | **int32** |  | 
**VulnerabilityLatestScanId** | **string** |  | 
**VulnerabilityScanStatus** | **string** |  | 

## Methods

### NewModelContainer

`func NewModelContainer(cpuMax float32, cpuUsage float32, dockerContainerCommand string, dockerContainerCreated string, dockerContainerIps []interface{}, dockerContainerName string, dockerContainerNetworkMode string, dockerContainerNetworks string, dockerContainerPorts string, dockerContainerState string, dockerContainerStateHuman string, dockerLabels map[string]interface{}, hostName string, image ModelContainerImage, malwareLatestScanId string, malwareScanStatus string, malwaresCount int32, memoryMax int32, memoryUsage int32, nodeId string, nodeName string, processes []ModelProcess, secretLatestScanId string, secretScanStatus string, secretsCount int32, uptime int32, vulnerabilitiesCount int32, vulnerabilityLatestScanId string, vulnerabilityScanStatus string, ) *ModelContainer`

NewModelContainer instantiates a new ModelContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelContainerWithDefaults

`func NewModelContainerWithDefaults() *ModelContainer`

NewModelContainerWithDefaults instantiates a new ModelContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuMax

`func (o *ModelContainer) GetCpuMax() float32`

GetCpuMax returns the CpuMax field if non-nil, zero value otherwise.

### GetCpuMaxOk

`func (o *ModelContainer) GetCpuMaxOk() (*float32, bool)`

GetCpuMaxOk returns a tuple with the CpuMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMax

`func (o *ModelContainer) SetCpuMax(v float32)`

SetCpuMax sets CpuMax field to given value.


### GetCpuUsage

`func (o *ModelContainer) GetCpuUsage() float32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *ModelContainer) GetCpuUsageOk() (*float32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *ModelContainer) SetCpuUsage(v float32)`

SetCpuUsage sets CpuUsage field to given value.


### GetDockerContainerCommand

`func (o *ModelContainer) GetDockerContainerCommand() string`

GetDockerContainerCommand returns the DockerContainerCommand field if non-nil, zero value otherwise.

### GetDockerContainerCommandOk

`func (o *ModelContainer) GetDockerContainerCommandOk() (*string, bool)`

GetDockerContainerCommandOk returns a tuple with the DockerContainerCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContainerCommand

`func (o *ModelContainer) SetDockerContainerCommand(v string)`

SetDockerContainerCommand sets DockerContainerCommand field to given value.


### GetDockerContainerCreated

`func (o *ModelContainer) GetDockerContainerCreated() string`

GetDockerContainerCreated returns the DockerContainerCreated field if non-nil, zero value otherwise.

### GetDockerContainerCreatedOk

`func (o *ModelContainer) GetDockerContainerCreatedOk() (*string, bool)`

GetDockerContainerCreatedOk returns a tuple with the DockerContainerCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContainerCreated

`func (o *ModelContainer) SetDockerContainerCreated(v string)`

SetDockerContainerCreated sets DockerContainerCreated field to given value.


### GetDockerContainerIps

`func (o *ModelContainer) GetDockerContainerIps() []interface{}`

GetDockerContainerIps returns the DockerContainerIps field if non-nil, zero value otherwise.

### GetDockerContainerIpsOk

`func (o *ModelContainer) GetDockerContainerIpsOk() (*[]interface{}, bool)`

GetDockerContainerIpsOk returns a tuple with the DockerContainerIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContainerIps

`func (o *ModelContainer) SetDockerContainerIps(v []interface{})`

SetDockerContainerIps sets DockerContainerIps field to given value.


### SetDockerContainerIpsNil

`func (o *ModelContainer) SetDockerContainerIpsNil(b bool)`

 SetDockerContainerIpsNil sets the value for DockerContainerIps to be an explicit nil

### UnsetDockerContainerIps
`func (o *ModelContainer) UnsetDockerContainerIps()`

UnsetDockerContainerIps ensures that no value is present for DockerContainerIps, not even an explicit nil
### GetDockerContainerName

`func (o *ModelContainer) GetDockerContainerName() string`

GetDockerContainerName returns the DockerContainerName field if non-nil, zero value otherwise.

### GetDockerContainerNameOk

`func (o *ModelContainer) GetDockerContainerNameOk() (*string, bool)`

GetDockerContainerNameOk returns a tuple with the DockerContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContainerName

`func (o *ModelContainer) SetDockerContainerName(v string)`

SetDockerContainerName sets DockerContainerName field to given value.


### GetDockerContainerNetworkMode

`func (o *ModelContainer) GetDockerContainerNetworkMode() string`

GetDockerContainerNetworkMode returns the DockerContainerNetworkMode field if non-nil, zero value otherwise.

### GetDockerContainerNetworkModeOk

`func (o *ModelContainer) GetDockerContainerNetworkModeOk() (*string, bool)`

GetDockerContainerNetworkModeOk returns a tuple with the DockerContainerNetworkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContainerNetworkMode

`func (o *ModelContainer) SetDockerContainerNetworkMode(v string)`

SetDockerContainerNetworkMode sets DockerContainerNetworkMode field to given value.


### GetDockerContainerNetworks

`func (o *ModelContainer) GetDockerContainerNetworks() string`

GetDockerContainerNetworks returns the DockerContainerNetworks field if non-nil, zero value otherwise.

### GetDockerContainerNetworksOk

`func (o *ModelContainer) GetDockerContainerNetworksOk() (*string, bool)`

GetDockerContainerNetworksOk returns a tuple with the DockerContainerNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContainerNetworks

`func (o *ModelContainer) SetDockerContainerNetworks(v string)`

SetDockerContainerNetworks sets DockerContainerNetworks field to given value.


### GetDockerContainerPorts

`func (o *ModelContainer) GetDockerContainerPorts() string`

GetDockerContainerPorts returns the DockerContainerPorts field if non-nil, zero value otherwise.

### GetDockerContainerPortsOk

`func (o *ModelContainer) GetDockerContainerPortsOk() (*string, bool)`

GetDockerContainerPortsOk returns a tuple with the DockerContainerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContainerPorts

`func (o *ModelContainer) SetDockerContainerPorts(v string)`

SetDockerContainerPorts sets DockerContainerPorts field to given value.


### GetDockerContainerState

`func (o *ModelContainer) GetDockerContainerState() string`

GetDockerContainerState returns the DockerContainerState field if non-nil, zero value otherwise.

### GetDockerContainerStateOk

`func (o *ModelContainer) GetDockerContainerStateOk() (*string, bool)`

GetDockerContainerStateOk returns a tuple with the DockerContainerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContainerState

`func (o *ModelContainer) SetDockerContainerState(v string)`

SetDockerContainerState sets DockerContainerState field to given value.


### GetDockerContainerStateHuman

`func (o *ModelContainer) GetDockerContainerStateHuman() string`

GetDockerContainerStateHuman returns the DockerContainerStateHuman field if non-nil, zero value otherwise.

### GetDockerContainerStateHumanOk

`func (o *ModelContainer) GetDockerContainerStateHumanOk() (*string, bool)`

GetDockerContainerStateHumanOk returns a tuple with the DockerContainerStateHuman field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContainerStateHuman

`func (o *ModelContainer) SetDockerContainerStateHuman(v string)`

SetDockerContainerStateHuman sets DockerContainerStateHuman field to given value.


### GetDockerLabels

`func (o *ModelContainer) GetDockerLabels() map[string]interface{}`

GetDockerLabels returns the DockerLabels field if non-nil, zero value otherwise.

### GetDockerLabelsOk

`func (o *ModelContainer) GetDockerLabelsOk() (*map[string]interface{}, bool)`

GetDockerLabelsOk returns a tuple with the DockerLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerLabels

`func (o *ModelContainer) SetDockerLabels(v map[string]interface{})`

SetDockerLabels sets DockerLabels field to given value.


### SetDockerLabelsNil

`func (o *ModelContainer) SetDockerLabelsNil(b bool)`

 SetDockerLabelsNil sets the value for DockerLabels to be an explicit nil

### UnsetDockerLabels
`func (o *ModelContainer) UnsetDockerLabels()`

UnsetDockerLabels ensures that no value is present for DockerLabels, not even an explicit nil
### GetHostName

`func (o *ModelContainer) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ModelContainer) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ModelContainer) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetImage

`func (o *ModelContainer) GetImage() ModelContainerImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ModelContainer) GetImageOk() (*ModelContainerImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ModelContainer) SetImage(v ModelContainerImage)`

SetImage sets Image field to given value.


### GetMalwareLatestScanId

`func (o *ModelContainer) GetMalwareLatestScanId() string`

GetMalwareLatestScanId returns the MalwareLatestScanId field if non-nil, zero value otherwise.

### GetMalwareLatestScanIdOk

`func (o *ModelContainer) GetMalwareLatestScanIdOk() (*string, bool)`

GetMalwareLatestScanIdOk returns a tuple with the MalwareLatestScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalwareLatestScanId

`func (o *ModelContainer) SetMalwareLatestScanId(v string)`

SetMalwareLatestScanId sets MalwareLatestScanId field to given value.


### GetMalwareScanStatus

`func (o *ModelContainer) GetMalwareScanStatus() string`

GetMalwareScanStatus returns the MalwareScanStatus field if non-nil, zero value otherwise.

### GetMalwareScanStatusOk

`func (o *ModelContainer) GetMalwareScanStatusOk() (*string, bool)`

GetMalwareScanStatusOk returns a tuple with the MalwareScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalwareScanStatus

`func (o *ModelContainer) SetMalwareScanStatus(v string)`

SetMalwareScanStatus sets MalwareScanStatus field to given value.


### GetMalwaresCount

`func (o *ModelContainer) GetMalwaresCount() int32`

GetMalwaresCount returns the MalwaresCount field if non-nil, zero value otherwise.

### GetMalwaresCountOk

`func (o *ModelContainer) GetMalwaresCountOk() (*int32, bool)`

GetMalwaresCountOk returns a tuple with the MalwaresCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalwaresCount

`func (o *ModelContainer) SetMalwaresCount(v int32)`

SetMalwaresCount sets MalwaresCount field to given value.


### GetMemoryMax

`func (o *ModelContainer) GetMemoryMax() int32`

GetMemoryMax returns the MemoryMax field if non-nil, zero value otherwise.

### GetMemoryMaxOk

`func (o *ModelContainer) GetMemoryMaxOk() (*int32, bool)`

GetMemoryMaxOk returns a tuple with the MemoryMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMax

`func (o *ModelContainer) SetMemoryMax(v int32)`

SetMemoryMax sets MemoryMax field to given value.


### GetMemoryUsage

`func (o *ModelContainer) GetMemoryUsage() int32`

GetMemoryUsage returns the MemoryUsage field if non-nil, zero value otherwise.

### GetMemoryUsageOk

`func (o *ModelContainer) GetMemoryUsageOk() (*int32, bool)`

GetMemoryUsageOk returns a tuple with the MemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsage

`func (o *ModelContainer) SetMemoryUsage(v int32)`

SetMemoryUsage sets MemoryUsage field to given value.


### GetNodeId

`func (o *ModelContainer) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelContainer) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelContainer) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNodeName

`func (o *ModelContainer) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *ModelContainer) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *ModelContainer) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.


### GetProcesses

`func (o *ModelContainer) GetProcesses() []ModelProcess`

GetProcesses returns the Processes field if non-nil, zero value otherwise.

### GetProcessesOk

`func (o *ModelContainer) GetProcessesOk() (*[]ModelProcess, bool)`

GetProcessesOk returns a tuple with the Processes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcesses

`func (o *ModelContainer) SetProcesses(v []ModelProcess)`

SetProcesses sets Processes field to given value.


### SetProcessesNil

`func (o *ModelContainer) SetProcessesNil(b bool)`

 SetProcessesNil sets the value for Processes to be an explicit nil

### UnsetProcesses
`func (o *ModelContainer) UnsetProcesses()`

UnsetProcesses ensures that no value is present for Processes, not even an explicit nil
### GetSecretLatestScanId

`func (o *ModelContainer) GetSecretLatestScanId() string`

GetSecretLatestScanId returns the SecretLatestScanId field if non-nil, zero value otherwise.

### GetSecretLatestScanIdOk

`func (o *ModelContainer) GetSecretLatestScanIdOk() (*string, bool)`

GetSecretLatestScanIdOk returns a tuple with the SecretLatestScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretLatestScanId

`func (o *ModelContainer) SetSecretLatestScanId(v string)`

SetSecretLatestScanId sets SecretLatestScanId field to given value.


### GetSecretScanStatus

`func (o *ModelContainer) GetSecretScanStatus() string`

GetSecretScanStatus returns the SecretScanStatus field if non-nil, zero value otherwise.

### GetSecretScanStatusOk

`func (o *ModelContainer) GetSecretScanStatusOk() (*string, bool)`

GetSecretScanStatusOk returns a tuple with the SecretScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretScanStatus

`func (o *ModelContainer) SetSecretScanStatus(v string)`

SetSecretScanStatus sets SecretScanStatus field to given value.


### GetSecretsCount

`func (o *ModelContainer) GetSecretsCount() int32`

GetSecretsCount returns the SecretsCount field if non-nil, zero value otherwise.

### GetSecretsCountOk

`func (o *ModelContainer) GetSecretsCountOk() (*int32, bool)`

GetSecretsCountOk returns a tuple with the SecretsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsCount

`func (o *ModelContainer) SetSecretsCount(v int32)`

SetSecretsCount sets SecretsCount field to given value.


### GetUptime

`func (o *ModelContainer) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *ModelContainer) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *ModelContainer) SetUptime(v int32)`

SetUptime sets Uptime field to given value.


### GetVulnerabilitiesCount

`func (o *ModelContainer) GetVulnerabilitiesCount() int32`

GetVulnerabilitiesCount returns the VulnerabilitiesCount field if non-nil, zero value otherwise.

### GetVulnerabilitiesCountOk

`func (o *ModelContainer) GetVulnerabilitiesCountOk() (*int32, bool)`

GetVulnerabilitiesCountOk returns a tuple with the VulnerabilitiesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilitiesCount

`func (o *ModelContainer) SetVulnerabilitiesCount(v int32)`

SetVulnerabilitiesCount sets VulnerabilitiesCount field to given value.


### GetVulnerabilityLatestScanId

`func (o *ModelContainer) GetVulnerabilityLatestScanId() string`

GetVulnerabilityLatestScanId returns the VulnerabilityLatestScanId field if non-nil, zero value otherwise.

### GetVulnerabilityLatestScanIdOk

`func (o *ModelContainer) GetVulnerabilityLatestScanIdOk() (*string, bool)`

GetVulnerabilityLatestScanIdOk returns a tuple with the VulnerabilityLatestScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityLatestScanId

`func (o *ModelContainer) SetVulnerabilityLatestScanId(v string)`

SetVulnerabilityLatestScanId sets VulnerabilityLatestScanId field to given value.


### GetVulnerabilityScanStatus

`func (o *ModelContainer) GetVulnerabilityScanStatus() string`

GetVulnerabilityScanStatus returns the VulnerabilityScanStatus field if non-nil, zero value otherwise.

### GetVulnerabilityScanStatusOk

`func (o *ModelContainer) GetVulnerabilityScanStatusOk() (*string, bool)`

GetVulnerabilityScanStatusOk returns a tuple with the VulnerabilityScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityScanStatus

`func (o *ModelContainer) SetVulnerabilityScanStatus(v string)`

SetVulnerabilityScanStatus sets VulnerabilityScanStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


