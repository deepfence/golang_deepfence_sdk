# ModelHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudComplianceLatestScanId** | **string** |  | 
**CloudComplianceScanStatus** | **string** |  | 
**CloudCompliancesCount** | **int32** |  | 
**CloudMetadata** | **map[string]interface{}** |  | 
**ComplianceLatestScanId** | **string** |  | 
**ComplianceScanStatus** | **string** |  | 
**CompliancesCount** | **int32** |  | 
**ContainerImages** | [**[]ModelContainerImage**](ModelContainerImage.md) |  | 
**Containers** | [**[]ModelContainer**](ModelContainer.md) |  | 
**HostName** | **string** |  | 
**InterfaceIps** | **string** |  | 
**InterfaceNames** | **string** |  | 
**KernelVersion** | **string** |  | 
**MalwareLatestScanId** | **string** |  | 
**MalwareScanStatus** | **string** |  | 
**MalwaresCount** | **int32** |  | 
**Metrics** | [**ModelComputeMetrics**](ModelComputeMetrics.md) |  | 
**NodeId** | **string** |  | 
**NodeName** | **string** |  | 
**Pods** | [**[]ModelPod**](ModelPod.md) |  | 
**Processes** | [**[]ModelProcess**](ModelProcess.md) |  | 
**SecretLatestScan** | **string** |  | 
**SecretScanStatus** | **string** |  | 
**SecretsCount** | **int32** |  | 
**Uptime** | **string** |  | 
**VulnerabilitiesCount** | **int32** |  | 
**VulnerabilityLatestScanId** | **string** |  | 
**VulnerabilityScanStatus** | **string** |  | 

## Methods

### NewModelHost

`func NewModelHost(cloudComplianceLatestScanId string, cloudComplianceScanStatus string, cloudCompliancesCount int32, cloudMetadata map[string]interface{}, complianceLatestScanId string, complianceScanStatus string, compliancesCount int32, containerImages []ModelContainerImage, containers []ModelContainer, hostName string, interfaceIps string, interfaceNames string, kernelVersion string, malwareLatestScanId string, malwareScanStatus string, malwaresCount int32, metrics ModelComputeMetrics, nodeId string, nodeName string, pods []ModelPod, processes []ModelProcess, secretLatestScan string, secretScanStatus string, secretsCount int32, uptime string, vulnerabilitiesCount int32, vulnerabilityLatestScanId string, vulnerabilityScanStatus string, ) *ModelHost`

NewModelHost instantiates a new ModelHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelHostWithDefaults

`func NewModelHostWithDefaults() *ModelHost`

NewModelHostWithDefaults instantiates a new ModelHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetCloudMetadata

`func (o *ModelHost) GetCloudMetadata() map[string]interface{}`

GetCloudMetadata returns the CloudMetadata field if non-nil, zero value otherwise.

### GetCloudMetadataOk

`func (o *ModelHost) GetCloudMetadataOk() (*map[string]interface{}, bool)`

GetCloudMetadataOk returns a tuple with the CloudMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudMetadata

`func (o *ModelHost) SetCloudMetadata(v map[string]interface{})`

SetCloudMetadata sets CloudMetadata field to given value.


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


### GetInterfaceIps

`func (o *ModelHost) GetInterfaceIps() string`

GetInterfaceIps returns the InterfaceIps field if non-nil, zero value otherwise.

### GetInterfaceIpsOk

`func (o *ModelHost) GetInterfaceIpsOk() (*string, bool)`

GetInterfaceIpsOk returns a tuple with the InterfaceIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIps

`func (o *ModelHost) SetInterfaceIps(v string)`

SetInterfaceIps sets InterfaceIps field to given value.


### GetInterfaceNames

`func (o *ModelHost) GetInterfaceNames() string`

GetInterfaceNames returns the InterfaceNames field if non-nil, zero value otherwise.

### GetInterfaceNamesOk

`func (o *ModelHost) GetInterfaceNamesOk() (*string, bool)`

GetInterfaceNamesOk returns a tuple with the InterfaceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceNames

`func (o *ModelHost) SetInterfaceNames(v string)`

SetInterfaceNames sets InterfaceNames field to given value.


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


### GetMetrics

`func (o *ModelHost) GetMetrics() ModelComputeMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ModelHost) GetMetricsOk() (*ModelComputeMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ModelHost) SetMetrics(v ModelComputeMetrics)`

SetMetrics sets Metrics field to given value.


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

`func (o *ModelHost) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *ModelHost) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *ModelHost) SetUptime(v string)`

SetUptime sets Uptime field to given value.


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


