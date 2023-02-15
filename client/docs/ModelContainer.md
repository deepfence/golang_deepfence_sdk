# ModelContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComplianceScanStatus** | Pointer to **string** |  | [optional] 
**CompliancesCount** | Pointer to **int32** |  | [optional] 
**DockerLabels** | **map[string]interface{}** |  | 
**HostName** | **string** |  | 
**Image** | [**ModelContainerImage**](ModelContainerImage.md) |  | 
**MalwareScanStatus** | Pointer to **string** |  | [optional] 
**MalwaresCount** | Pointer to **int32** |  | [optional] 
**Metadata** | **map[string]interface{}** |  | 
**Metrics** | [**ModelComputeMetrics**](ModelComputeMetrics.md) |  | 
**Name** | **string** |  | 
**NodeId** | **string** |  | 
**Processes** | [**[]ModelProcess**](ModelProcess.md) |  | 
**SecretScanStatus** | Pointer to **string** |  | [optional] 
**SecretsCount** | Pointer to **int32** |  | [optional] 
**VulnerabilitiesCount** | Pointer to **int32** |  | [optional] 
**VulnerabilityScanStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewModelContainer

`func NewModelContainer(dockerLabels map[string]interface{}, hostName string, image ModelContainerImage, metadata map[string]interface{}, metrics ModelComputeMetrics, name string, nodeId string, processes []ModelProcess, ) *ModelContainer`

NewModelContainer instantiates a new ModelContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelContainerWithDefaults

`func NewModelContainerWithDefaults() *ModelContainer`

NewModelContainerWithDefaults instantiates a new ModelContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### HasComplianceScanStatus

`func (o *ModelContainer) HasComplianceScanStatus() bool`

HasComplianceScanStatus returns a boolean if a field has been set.

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

### HasCompliancesCount

`func (o *ModelContainer) HasCompliancesCount() bool`

HasCompliancesCount returns a boolean if a field has been set.

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

### HasMalwareScanStatus

`func (o *ModelContainer) HasMalwareScanStatus() bool`

HasMalwareScanStatus returns a boolean if a field has been set.

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

### HasMalwaresCount

`func (o *ModelContainer) HasMalwaresCount() bool`

HasMalwaresCount returns a boolean if a field has been set.

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


### GetName

`func (o *ModelContainer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelContainer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelContainer) SetName(v string)`

SetName sets Name field to given value.


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

### HasSecretScanStatus

`func (o *ModelContainer) HasSecretScanStatus() bool`

HasSecretScanStatus returns a boolean if a field has been set.

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

### HasSecretsCount

`func (o *ModelContainer) HasSecretsCount() bool`

HasSecretsCount returns a boolean if a field has been set.

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

### HasVulnerabilitiesCount

`func (o *ModelContainer) HasVulnerabilitiesCount() bool`

HasVulnerabilitiesCount returns a boolean if a field has been set.

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

### HasVulnerabilityScanStatus

`func (o *ModelContainer) HasVulnerabilityScanStatus() bool`

HasVulnerabilityScanStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


