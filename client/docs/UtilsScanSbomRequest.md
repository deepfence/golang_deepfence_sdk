# UtilsScanSbomRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerName** | Pointer to **string** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**ImageName** | Pointer to **string** |  | [optional] 
**KubernetesClusterName** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**NodeType** | Pointer to **string** |  | [optional] 
**RegistryId** | Pointer to **string** |  | [optional] 
**Sbom** | **string** |  | 
**SbomFilePath** | Pointer to **string** |  | [optional] 
**ScanId** | **string** |  | 
**ScanType** | Pointer to **string** |  | [optional] 
**SkipScan** | Pointer to **bool** |  | [optional] 

## Methods

### NewUtilsScanSbomRequest

`func NewUtilsScanSbomRequest(sbom string, scanId string, ) *UtilsScanSbomRequest`

NewUtilsScanSbomRequest instantiates a new UtilsScanSbomRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilsScanSbomRequestWithDefaults

`func NewUtilsScanSbomRequestWithDefaults() *UtilsScanSbomRequest`

NewUtilsScanSbomRequestWithDefaults instantiates a new UtilsScanSbomRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerName

`func (o *UtilsScanSbomRequest) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *UtilsScanSbomRequest) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *UtilsScanSbomRequest) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *UtilsScanSbomRequest) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetHostName

`func (o *UtilsScanSbomRequest) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *UtilsScanSbomRequest) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *UtilsScanSbomRequest) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *UtilsScanSbomRequest) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetImageId

`func (o *UtilsScanSbomRequest) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *UtilsScanSbomRequest) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *UtilsScanSbomRequest) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *UtilsScanSbomRequest) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetImageName

`func (o *UtilsScanSbomRequest) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *UtilsScanSbomRequest) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *UtilsScanSbomRequest) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *UtilsScanSbomRequest) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetKubernetesClusterName

`func (o *UtilsScanSbomRequest) GetKubernetesClusterName() string`

GetKubernetesClusterName returns the KubernetesClusterName field if non-nil, zero value otherwise.

### GetKubernetesClusterNameOk

`func (o *UtilsScanSbomRequest) GetKubernetesClusterNameOk() (*string, bool)`

GetKubernetesClusterNameOk returns a tuple with the KubernetesClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesClusterName

`func (o *UtilsScanSbomRequest) SetKubernetesClusterName(v string)`

SetKubernetesClusterName sets KubernetesClusterName field to given value.

### HasKubernetesClusterName

`func (o *UtilsScanSbomRequest) HasKubernetesClusterName() bool`

HasKubernetesClusterName returns a boolean if a field has been set.

### GetMode

`func (o *UtilsScanSbomRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *UtilsScanSbomRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *UtilsScanSbomRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *UtilsScanSbomRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNodeId

`func (o *UtilsScanSbomRequest) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *UtilsScanSbomRequest) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *UtilsScanSbomRequest) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *UtilsScanSbomRequest) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNodeType

`func (o *UtilsScanSbomRequest) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *UtilsScanSbomRequest) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *UtilsScanSbomRequest) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *UtilsScanSbomRequest) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### GetRegistryId

`func (o *UtilsScanSbomRequest) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *UtilsScanSbomRequest) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *UtilsScanSbomRequest) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *UtilsScanSbomRequest) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.

### GetSbom

`func (o *UtilsScanSbomRequest) GetSbom() string`

GetSbom returns the Sbom field if non-nil, zero value otherwise.

### GetSbomOk

`func (o *UtilsScanSbomRequest) GetSbomOk() (*string, bool)`

GetSbomOk returns a tuple with the Sbom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbom

`func (o *UtilsScanSbomRequest) SetSbom(v string)`

SetSbom sets Sbom field to given value.


### GetSbomFilePath

`func (o *UtilsScanSbomRequest) GetSbomFilePath() string`

GetSbomFilePath returns the SbomFilePath field if non-nil, zero value otherwise.

### GetSbomFilePathOk

`func (o *UtilsScanSbomRequest) GetSbomFilePathOk() (*string, bool)`

GetSbomFilePathOk returns a tuple with the SbomFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbomFilePath

`func (o *UtilsScanSbomRequest) SetSbomFilePath(v string)`

SetSbomFilePath sets SbomFilePath field to given value.

### HasSbomFilePath

`func (o *UtilsScanSbomRequest) HasSbomFilePath() bool`

HasSbomFilePath returns a boolean if a field has been set.

### GetScanId

`func (o *UtilsScanSbomRequest) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *UtilsScanSbomRequest) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *UtilsScanSbomRequest) SetScanId(v string)`

SetScanId sets ScanId field to given value.


### GetScanType

`func (o *UtilsScanSbomRequest) GetScanType() string`

GetScanType returns the ScanType field if non-nil, zero value otherwise.

### GetScanTypeOk

`func (o *UtilsScanSbomRequest) GetScanTypeOk() (*string, bool)`

GetScanTypeOk returns a tuple with the ScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanType

`func (o *UtilsScanSbomRequest) SetScanType(v string)`

SetScanType sets ScanType field to given value.

### HasScanType

`func (o *UtilsScanSbomRequest) HasScanType() bool`

HasScanType returns a boolean if a field has been set.

### GetSkipScan

`func (o *UtilsScanSbomRequest) GetSkipScan() bool`

GetSkipScan returns the SkipScan field if non-nil, zero value otherwise.

### GetSkipScanOk

`func (o *UtilsScanSbomRequest) GetSkipScanOk() (*bool, bool)`

GetSkipScanOk returns a tuple with the SkipScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipScan

`func (o *UtilsScanSbomRequest) SetSkipScan(v bool)`

SetSkipScan sets SkipScan field to given value.

### HasSkipScan

`func (o *UtilsScanSbomRequest) HasSkipScan() bool`

HasSkipScan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


