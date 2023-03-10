# ModelContainerImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComplianceScanStatus** | **string** |  | 
**CompliancesCount** | **int32** |  | 
**DockerImageName** | **string** |  | 
**DockerImageSize** | **string** |  | 
**DockerImageTag** | **string** |  | 
**MalwareScanStatus** | **string** |  | 
**MalwaresCount** | **int32** |  | 
**Metadata** | **map[string]interface{}** |  | 
**Metrics** | [**ModelComputeMetrics**](ModelComputeMetrics.md) |  | 
**NodeId** | **string** |  | 
**NodeName** | **string** |  | 
**SecretScanStatus** | **string** |  | 
**SecretsCount** | **int32** |  | 
**VulnerabilitiesCount** | **int32** |  | 
**VulnerabilityScanStatus** | **string** |  | 

## Methods

### NewModelContainerImage

`func NewModelContainerImage(complianceScanStatus string, compliancesCount int32, dockerImageName string, dockerImageSize string, dockerImageTag string, malwareScanStatus string, malwaresCount int32, metadata map[string]interface{}, metrics ModelComputeMetrics, nodeId string, nodeName string, secretScanStatus string, secretsCount int32, vulnerabilitiesCount int32, vulnerabilityScanStatus string, ) *ModelContainerImage`

NewModelContainerImage instantiates a new ModelContainerImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelContainerImageWithDefaults

`func NewModelContainerImageWithDefaults() *ModelContainerImage`

NewModelContainerImageWithDefaults instantiates a new ModelContainerImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComplianceScanStatus

`func (o *ModelContainerImage) GetComplianceScanStatus() string`

GetComplianceScanStatus returns the ComplianceScanStatus field if non-nil, zero value otherwise.

### GetComplianceScanStatusOk

`func (o *ModelContainerImage) GetComplianceScanStatusOk() (*string, bool)`

GetComplianceScanStatusOk returns a tuple with the ComplianceScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceScanStatus

`func (o *ModelContainerImage) SetComplianceScanStatus(v string)`

SetComplianceScanStatus sets ComplianceScanStatus field to given value.


### GetCompliancesCount

`func (o *ModelContainerImage) GetCompliancesCount() int32`

GetCompliancesCount returns the CompliancesCount field if non-nil, zero value otherwise.

### GetCompliancesCountOk

`func (o *ModelContainerImage) GetCompliancesCountOk() (*int32, bool)`

GetCompliancesCountOk returns a tuple with the CompliancesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliancesCount

`func (o *ModelContainerImage) SetCompliancesCount(v int32)`

SetCompliancesCount sets CompliancesCount field to given value.


### GetDockerImageName

`func (o *ModelContainerImage) GetDockerImageName() string`

GetDockerImageName returns the DockerImageName field if non-nil, zero value otherwise.

### GetDockerImageNameOk

`func (o *ModelContainerImage) GetDockerImageNameOk() (*string, bool)`

GetDockerImageNameOk returns a tuple with the DockerImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageName

`func (o *ModelContainerImage) SetDockerImageName(v string)`

SetDockerImageName sets DockerImageName field to given value.


### GetDockerImageSize

`func (o *ModelContainerImage) GetDockerImageSize() string`

GetDockerImageSize returns the DockerImageSize field if non-nil, zero value otherwise.

### GetDockerImageSizeOk

`func (o *ModelContainerImage) GetDockerImageSizeOk() (*string, bool)`

GetDockerImageSizeOk returns a tuple with the DockerImageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageSize

`func (o *ModelContainerImage) SetDockerImageSize(v string)`

SetDockerImageSize sets DockerImageSize field to given value.


### GetDockerImageTag

`func (o *ModelContainerImage) GetDockerImageTag() string`

GetDockerImageTag returns the DockerImageTag field if non-nil, zero value otherwise.

### GetDockerImageTagOk

`func (o *ModelContainerImage) GetDockerImageTagOk() (*string, bool)`

GetDockerImageTagOk returns a tuple with the DockerImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageTag

`func (o *ModelContainerImage) SetDockerImageTag(v string)`

SetDockerImageTag sets DockerImageTag field to given value.


### GetMalwareScanStatus

`func (o *ModelContainerImage) GetMalwareScanStatus() string`

GetMalwareScanStatus returns the MalwareScanStatus field if non-nil, zero value otherwise.

### GetMalwareScanStatusOk

`func (o *ModelContainerImage) GetMalwareScanStatusOk() (*string, bool)`

GetMalwareScanStatusOk returns a tuple with the MalwareScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalwareScanStatus

`func (o *ModelContainerImage) SetMalwareScanStatus(v string)`

SetMalwareScanStatus sets MalwareScanStatus field to given value.


### GetMalwaresCount

`func (o *ModelContainerImage) GetMalwaresCount() int32`

GetMalwaresCount returns the MalwaresCount field if non-nil, zero value otherwise.

### GetMalwaresCountOk

`func (o *ModelContainerImage) GetMalwaresCountOk() (*int32, bool)`

GetMalwaresCountOk returns a tuple with the MalwaresCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalwaresCount

`func (o *ModelContainerImage) SetMalwaresCount(v int32)`

SetMalwaresCount sets MalwaresCount field to given value.


### GetMetadata

`func (o *ModelContainerImage) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ModelContainerImage) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ModelContainerImage) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetMetrics

`func (o *ModelContainerImage) GetMetrics() ModelComputeMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ModelContainerImage) GetMetricsOk() (*ModelComputeMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ModelContainerImage) SetMetrics(v ModelComputeMetrics)`

SetMetrics sets Metrics field to given value.


### GetNodeId

`func (o *ModelContainerImage) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelContainerImage) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelContainerImage) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNodeName

`func (o *ModelContainerImage) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *ModelContainerImage) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *ModelContainerImage) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.


### GetSecretScanStatus

`func (o *ModelContainerImage) GetSecretScanStatus() string`

GetSecretScanStatus returns the SecretScanStatus field if non-nil, zero value otherwise.

### GetSecretScanStatusOk

`func (o *ModelContainerImage) GetSecretScanStatusOk() (*string, bool)`

GetSecretScanStatusOk returns a tuple with the SecretScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretScanStatus

`func (o *ModelContainerImage) SetSecretScanStatus(v string)`

SetSecretScanStatus sets SecretScanStatus field to given value.


### GetSecretsCount

`func (o *ModelContainerImage) GetSecretsCount() int32`

GetSecretsCount returns the SecretsCount field if non-nil, zero value otherwise.

### GetSecretsCountOk

`func (o *ModelContainerImage) GetSecretsCountOk() (*int32, bool)`

GetSecretsCountOk returns a tuple with the SecretsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsCount

`func (o *ModelContainerImage) SetSecretsCount(v int32)`

SetSecretsCount sets SecretsCount field to given value.


### GetVulnerabilitiesCount

`func (o *ModelContainerImage) GetVulnerabilitiesCount() int32`

GetVulnerabilitiesCount returns the VulnerabilitiesCount field if non-nil, zero value otherwise.

### GetVulnerabilitiesCountOk

`func (o *ModelContainerImage) GetVulnerabilitiesCountOk() (*int32, bool)`

GetVulnerabilitiesCountOk returns a tuple with the VulnerabilitiesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilitiesCount

`func (o *ModelContainerImage) SetVulnerabilitiesCount(v int32)`

SetVulnerabilitiesCount sets VulnerabilitiesCount field to given value.


### GetVulnerabilityScanStatus

`func (o *ModelContainerImage) GetVulnerabilityScanStatus() string`

GetVulnerabilityScanStatus returns the VulnerabilityScanStatus field if non-nil, zero value otherwise.

### GetVulnerabilityScanStatusOk

`func (o *ModelContainerImage) GetVulnerabilityScanStatusOk() (*string, bool)`

GetVulnerabilityScanStatusOk returns a tuple with the VulnerabilityScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityScanStatus

`func (o *ModelContainerImage) SetVulnerabilityScanStatus(v string)`

SetVulnerabilityScanStatus sets VulnerabilityScanStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


