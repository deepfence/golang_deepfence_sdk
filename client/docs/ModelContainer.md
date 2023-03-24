# ModelContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudComplianceLatestScanId** | **string** |  | 
**CloudComplianceScanStatus** | **string** |  | 
**CloudCompliancesCount** | **int32** |  | 
**ComplianceLatestScanId** | **string** |  | 
**ComplianceScanStatus** | **string** |  | 
**CompliancesCount** | **int32** |  | 
**DockerContainerName** | **string** |  | 
**DockerLabels** | **map[string]interface{}** |  | 
**HostName** | **string** |  | 
**Image** | [**ModelContainerImage**](ModelContainerImage.md) |  | 
**MalwareLatestScanId** | **string** |  | 
**MalwareScanStatus** | **string** |  | 
**MalwaresCount** | **int32** |  | 
**Metadata** | **map[string]interface{}** |  | 
**Metrics** | [**ModelComputeMetrics**](ModelComputeMetrics.md) |  | 
**NodeId** | **string** |  | 
**NodeName** | **string** |  | 
**Processes** | [**[]ModelProcess**](ModelProcess.md) |  | 
**SecretLatestScan** | **string** |  | 
**SecretScanStatus** | **string** |  | 
**SecretsCount** | **int32** |  | 
**VulnerabilitiesCount** | **int32** |  | 
**VulnerabilityLatestScanId** | **string** |  | 
**VulnerabilityScanStatus** | **string** |  | 

## Methods

### NewModelContainer

`func NewModelContainer(cloudComplianceLatestScanId string, cloudComplianceScanStatus string, cloudCompliancesCount int32, complianceLatestScanId string, complianceScanStatus string, compliancesCount int32, dockerContainerName string, dockerLabels map[string]interface{}, hostName string, image ModelContainerImage, malwareLatestScanId string, malwareScanStatus string, malwaresCount int32, metadata map[string]interface{}, metrics ModelComputeMetrics, nodeId string, nodeName string, processes []ModelProcess, secretLatestScan string, secretScanStatus string, secretsCount int32, vulnerabilitiesCount int32, vulnerabilityLatestScanId string, vulnerabilityScanStatus string, ) *ModelContainer`

NewModelContainer instantiates a new ModelContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelContainerWithDefaults

`func NewModelContainerWithDefaults() *ModelContainer`

NewModelContainerWithDefaults instantiates a new ModelContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudComplianceLatestScanId

`func (o *ModelContainer) GetCloudComplianceLatestScanId() string`

GetCloudComplianceLatestScanId returns the CloudComplianceLatestScanId field if non-nil, zero value otherwise.

### GetCloudComplianceLatestScanIdOk

`func (o *ModelContainer) GetCloudComplianceLatestScanIdOk() (*string, bool)`

GetCloudComplianceLatestScanIdOk returns a tuple with the CloudComplianceLatestScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudComplianceLatestScanId

`func (o *ModelContainer) SetCloudComplianceLatestScanId(v string)`

SetCloudComplianceLatestScanId sets CloudComplianceLatestScanId field to given value.


### GetCloudComplianceScanStatus

`func (o *ModelContainer) GetCloudComplianceScanStatus() string`

GetCloudComplianceScanStatus returns the CloudComplianceScanStatus field if non-nil, zero value otherwise.

### GetCloudComplianceScanStatusOk

`func (o *ModelContainer) GetCloudComplianceScanStatusOk() (*string, bool)`

GetCloudComplianceScanStatusOk returns a tuple with the CloudComplianceScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudComplianceScanStatus

`func (o *ModelContainer) SetCloudComplianceScanStatus(v string)`

SetCloudComplianceScanStatus sets CloudComplianceScanStatus field to given value.


### GetCloudCompliancesCount

`func (o *ModelContainer) GetCloudCompliancesCount() int32`

GetCloudCompliancesCount returns the CloudCompliancesCount field if non-nil, zero value otherwise.

### GetCloudCompliancesCountOk

`func (o *ModelContainer) GetCloudCompliancesCountOk() (*int32, bool)`

GetCloudCompliancesCountOk returns a tuple with the CloudCompliancesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCompliancesCount

`func (o *ModelContainer) SetCloudCompliancesCount(v int32)`

SetCloudCompliancesCount sets CloudCompliancesCount field to given value.


### GetComplianceLatestScanId

`func (o *ModelContainer) GetComplianceLatestScanId() string`

GetComplianceLatestScanId returns the ComplianceLatestScanId field if non-nil, zero value otherwise.

### GetComplianceLatestScanIdOk

`func (o *ModelContainer) GetComplianceLatestScanIdOk() (*string, bool)`

GetComplianceLatestScanIdOk returns a tuple with the ComplianceLatestScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceLatestScanId

`func (o *ModelContainer) SetComplianceLatestScanId(v string)`

SetComplianceLatestScanId sets ComplianceLatestScanId field to given value.


### GetComplianceScanStatus

`func (o *ModelContainer) GetComplianceScanStatus() string`

GetComplianceScanStatus returns the ComplianceScanStatus field if non-nil, zero value otherwise.

### GetComplianceScanStatusOk

`func (o *ModelContainer) GetComplianceScanStatusOk() (*string, bool)`

GetComplianceScanStatusOk returns a tuple with the ComplianceScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceScanStatus

`func (o *ModelContainer) SetComplianceScanStatus(v string)`

SetComplianceScanStatus sets ComplianceScanStatus field to given value.


### GetCompliancesCount

`func (o *ModelContainer) GetCompliancesCount() int32`

GetCompliancesCount returns the CompliancesCount field if non-nil, zero value otherwise.

### GetCompliancesCountOk

`func (o *ModelContainer) GetCompliancesCountOk() (*int32, bool)`

GetCompliancesCountOk returns a tuple with the CompliancesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliancesCount

`func (o *ModelContainer) SetCompliancesCount(v int32)`

SetCompliancesCount sets CompliancesCount field to given value.


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


### GetMetadata

`func (o *ModelContainer) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ModelContainer) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ModelContainer) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetMetrics

`func (o *ModelContainer) GetMetrics() ModelComputeMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ModelContainer) GetMetricsOk() (*ModelComputeMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ModelContainer) SetMetrics(v ModelComputeMetrics)`

SetMetrics sets Metrics field to given value.


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
### GetSecretLatestScan

`func (o *ModelContainer) GetSecretLatestScan() string`

GetSecretLatestScan returns the SecretLatestScan field if non-nil, zero value otherwise.

### GetSecretLatestScanOk

`func (o *ModelContainer) GetSecretLatestScanOk() (*string, bool)`

GetSecretLatestScanOk returns a tuple with the SecretLatestScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretLatestScan

`func (o *ModelContainer) SetSecretLatestScan(v string)`

SetSecretLatestScan sets SecretLatestScan field to given value.


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


