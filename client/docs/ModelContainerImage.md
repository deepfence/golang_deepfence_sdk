# ModelContainerImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | [**[]ModelContainer**](ModelContainer.md) |  | 
**DockerImageCreatedAt** | **string** |  | 
**DockerImageId** | **string** |  | 
**DockerImageName** | **string** |  | 
**DockerImageSize** | **string** |  | 
**DockerImageTag** | **string** |  | 
**DockerImageTagList** | **[]string** |  | 
**DockerImageVirtualSize** | **string** |  | 
**ImageNodeId** | **string** |  | 
**MalwareLatestScanId** | **string** |  | 
**MalwareScanStatus** | **string** |  | 
**MalwaresCount** | **int32** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**NodeId** | **string** |  | 
**NodeName** | **string** |  | 
**SecretLatestScanId** | **string** |  | 
**SecretScanStatus** | **string** |  | 
**SecretsCount** | **int32** |  | 
**VulnerabilitiesCount** | **int32** |  | 
**VulnerabilityLatestScanId** | **string** |  | 
**VulnerabilityScanStatus** | **string** |  | 

## Methods

### NewModelContainerImage

`func NewModelContainerImage(containers []ModelContainer, dockerImageCreatedAt string, dockerImageId string, dockerImageName string, dockerImageSize string, dockerImageTag string, dockerImageTagList []string, dockerImageVirtualSize string, imageNodeId string, malwareLatestScanId string, malwareScanStatus string, malwaresCount int32, nodeId string, nodeName string, secretLatestScanId string, secretScanStatus string, secretsCount int32, vulnerabilitiesCount int32, vulnerabilityLatestScanId string, vulnerabilityScanStatus string, ) *ModelContainerImage`

NewModelContainerImage instantiates a new ModelContainerImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelContainerImageWithDefaults

`func NewModelContainerImageWithDefaults() *ModelContainerImage`

NewModelContainerImageWithDefaults instantiates a new ModelContainerImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainers

`func (o *ModelContainerImage) GetContainers() []ModelContainer`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *ModelContainerImage) GetContainersOk() (*[]ModelContainer, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *ModelContainerImage) SetContainers(v []ModelContainer)`

SetContainers sets Containers field to given value.


### SetContainersNil

`func (o *ModelContainerImage) SetContainersNil(b bool)`

 SetContainersNil sets the value for Containers to be an explicit nil

### UnsetContainers
`func (o *ModelContainerImage) UnsetContainers()`

UnsetContainers ensures that no value is present for Containers, not even an explicit nil
### GetDockerImageCreatedAt

`func (o *ModelContainerImage) GetDockerImageCreatedAt() string`

GetDockerImageCreatedAt returns the DockerImageCreatedAt field if non-nil, zero value otherwise.

### GetDockerImageCreatedAtOk

`func (o *ModelContainerImage) GetDockerImageCreatedAtOk() (*string, bool)`

GetDockerImageCreatedAtOk returns a tuple with the DockerImageCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageCreatedAt

`func (o *ModelContainerImage) SetDockerImageCreatedAt(v string)`

SetDockerImageCreatedAt sets DockerImageCreatedAt field to given value.


### GetDockerImageId

`func (o *ModelContainerImage) GetDockerImageId() string`

GetDockerImageId returns the DockerImageId field if non-nil, zero value otherwise.

### GetDockerImageIdOk

`func (o *ModelContainerImage) GetDockerImageIdOk() (*string, bool)`

GetDockerImageIdOk returns a tuple with the DockerImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageId

`func (o *ModelContainerImage) SetDockerImageId(v string)`

SetDockerImageId sets DockerImageId field to given value.


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


### GetDockerImageTagList

`func (o *ModelContainerImage) GetDockerImageTagList() []string`

GetDockerImageTagList returns the DockerImageTagList field if non-nil, zero value otherwise.

### GetDockerImageTagListOk

`func (o *ModelContainerImage) GetDockerImageTagListOk() (*[]string, bool)`

GetDockerImageTagListOk returns a tuple with the DockerImageTagList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageTagList

`func (o *ModelContainerImage) SetDockerImageTagList(v []string)`

SetDockerImageTagList sets DockerImageTagList field to given value.


### SetDockerImageTagListNil

`func (o *ModelContainerImage) SetDockerImageTagListNil(b bool)`

 SetDockerImageTagListNil sets the value for DockerImageTagList to be an explicit nil

### UnsetDockerImageTagList
`func (o *ModelContainerImage) UnsetDockerImageTagList()`

UnsetDockerImageTagList ensures that no value is present for DockerImageTagList, not even an explicit nil
### GetDockerImageVirtualSize

`func (o *ModelContainerImage) GetDockerImageVirtualSize() string`

GetDockerImageVirtualSize returns the DockerImageVirtualSize field if non-nil, zero value otherwise.

### GetDockerImageVirtualSizeOk

`func (o *ModelContainerImage) GetDockerImageVirtualSizeOk() (*string, bool)`

GetDockerImageVirtualSizeOk returns a tuple with the DockerImageVirtualSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageVirtualSize

`func (o *ModelContainerImage) SetDockerImageVirtualSize(v string)`

SetDockerImageVirtualSize sets DockerImageVirtualSize field to given value.


### GetImageNodeId

`func (o *ModelContainerImage) GetImageNodeId() string`

GetImageNodeId returns the ImageNodeId field if non-nil, zero value otherwise.

### GetImageNodeIdOk

`func (o *ModelContainerImage) GetImageNodeIdOk() (*string, bool)`

GetImageNodeIdOk returns a tuple with the ImageNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageNodeId

`func (o *ModelContainerImage) SetImageNodeId(v string)`

SetImageNodeId sets ImageNodeId field to given value.


### GetMalwareLatestScanId

`func (o *ModelContainerImage) GetMalwareLatestScanId() string`

GetMalwareLatestScanId returns the MalwareLatestScanId field if non-nil, zero value otherwise.

### GetMalwareLatestScanIdOk

`func (o *ModelContainerImage) GetMalwareLatestScanIdOk() (*string, bool)`

GetMalwareLatestScanIdOk returns a tuple with the MalwareLatestScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalwareLatestScanId

`func (o *ModelContainerImage) SetMalwareLatestScanId(v string)`

SetMalwareLatestScanId sets MalwareLatestScanId field to given value.


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

### HasMetadata

`func (o *ModelContainerImage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ModelContainerImage) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ModelContainerImage) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
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


### GetSecretLatestScanId

`func (o *ModelContainerImage) GetSecretLatestScanId() string`

GetSecretLatestScanId returns the SecretLatestScanId field if non-nil, zero value otherwise.

### GetSecretLatestScanIdOk

`func (o *ModelContainerImage) GetSecretLatestScanIdOk() (*string, bool)`

GetSecretLatestScanIdOk returns a tuple with the SecretLatestScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretLatestScanId

`func (o *ModelContainerImage) SetSecretLatestScanId(v string)`

SetSecretLatestScanId sets SecretLatestScanId field to given value.


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


### GetVulnerabilityLatestScanId

`func (o *ModelContainerImage) GetVulnerabilityLatestScanId() string`

GetVulnerabilityLatestScanId returns the VulnerabilityLatestScanId field if non-nil, zero value otherwise.

### GetVulnerabilityLatestScanIdOk

`func (o *ModelContainerImage) GetVulnerabilityLatestScanIdOk() (*string, bool)`

GetVulnerabilityLatestScanIdOk returns a tuple with the VulnerabilityLatestScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityLatestScanId

`func (o *ModelContainerImage) SetVulnerabilityLatestScanId(v string)`

SetVulnerabilityLatestScanId sets VulnerabilityLatestScanId field to given value.


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


