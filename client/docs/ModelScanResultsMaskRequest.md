# ModelScanResultsMaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaskAcrossHostsAndImages** | Pointer to **bool** |  | [optional] 
**MaskInThisHostOrImageTags** | Pointer to **bool** |  | [optional] 
**MaskInThisImageTag** | Pointer to **bool** |  | [optional] 
**ResultIds** | **[]string** |  | 
**ScanId** | **string** |  | 
**ScanType** | **string** |  | 

## Methods

### NewModelScanResultsMaskRequest

`func NewModelScanResultsMaskRequest(resultIds []string, scanId string, scanType string, ) *ModelScanResultsMaskRequest`

NewModelScanResultsMaskRequest instantiates a new ModelScanResultsMaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelScanResultsMaskRequestWithDefaults

`func NewModelScanResultsMaskRequestWithDefaults() *ModelScanResultsMaskRequest`

NewModelScanResultsMaskRequestWithDefaults instantiates a new ModelScanResultsMaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaskAcrossHostsAndImages

`func (o *ModelScanResultsMaskRequest) GetMaskAcrossHostsAndImages() bool`

GetMaskAcrossHostsAndImages returns the MaskAcrossHostsAndImages field if non-nil, zero value otherwise.

### GetMaskAcrossHostsAndImagesOk

`func (o *ModelScanResultsMaskRequest) GetMaskAcrossHostsAndImagesOk() (*bool, bool)`

GetMaskAcrossHostsAndImagesOk returns a tuple with the MaskAcrossHostsAndImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskAcrossHostsAndImages

`func (o *ModelScanResultsMaskRequest) SetMaskAcrossHostsAndImages(v bool)`

SetMaskAcrossHostsAndImages sets MaskAcrossHostsAndImages field to given value.

### HasMaskAcrossHostsAndImages

`func (o *ModelScanResultsMaskRequest) HasMaskAcrossHostsAndImages() bool`

HasMaskAcrossHostsAndImages returns a boolean if a field has been set.

### GetMaskInThisHostOrImageTags

`func (o *ModelScanResultsMaskRequest) GetMaskInThisHostOrImageTags() bool`

GetMaskInThisHostOrImageTags returns the MaskInThisHostOrImageTags field if non-nil, zero value otherwise.

### GetMaskInThisHostOrImageTagsOk

`func (o *ModelScanResultsMaskRequest) GetMaskInThisHostOrImageTagsOk() (*bool, bool)`

GetMaskInThisHostOrImageTagsOk returns a tuple with the MaskInThisHostOrImageTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskInThisHostOrImageTags

`func (o *ModelScanResultsMaskRequest) SetMaskInThisHostOrImageTags(v bool)`

SetMaskInThisHostOrImageTags sets MaskInThisHostOrImageTags field to given value.

### HasMaskInThisHostOrImageTags

`func (o *ModelScanResultsMaskRequest) HasMaskInThisHostOrImageTags() bool`

HasMaskInThisHostOrImageTags returns a boolean if a field has been set.

### GetMaskInThisImageTag

`func (o *ModelScanResultsMaskRequest) GetMaskInThisImageTag() bool`

GetMaskInThisImageTag returns the MaskInThisImageTag field if non-nil, zero value otherwise.

### GetMaskInThisImageTagOk

`func (o *ModelScanResultsMaskRequest) GetMaskInThisImageTagOk() (*bool, bool)`

GetMaskInThisImageTagOk returns a tuple with the MaskInThisImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskInThisImageTag

`func (o *ModelScanResultsMaskRequest) SetMaskInThisImageTag(v bool)`

SetMaskInThisImageTag sets MaskInThisImageTag field to given value.

### HasMaskInThisImageTag

`func (o *ModelScanResultsMaskRequest) HasMaskInThisImageTag() bool`

HasMaskInThisImageTag returns a boolean if a field has been set.

### GetResultIds

`func (o *ModelScanResultsMaskRequest) GetResultIds() []string`

GetResultIds returns the ResultIds field if non-nil, zero value otherwise.

### GetResultIdsOk

`func (o *ModelScanResultsMaskRequest) GetResultIdsOk() (*[]string, bool)`

GetResultIdsOk returns a tuple with the ResultIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultIds

`func (o *ModelScanResultsMaskRequest) SetResultIds(v []string)`

SetResultIds sets ResultIds field to given value.


### SetResultIdsNil

`func (o *ModelScanResultsMaskRequest) SetResultIdsNil(b bool)`

 SetResultIdsNil sets the value for ResultIds to be an explicit nil

### UnsetResultIds
`func (o *ModelScanResultsMaskRequest) UnsetResultIds()`

UnsetResultIds ensures that no value is present for ResultIds, not even an explicit nil
### GetScanId

`func (o *ModelScanResultsMaskRequest) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *ModelScanResultsMaskRequest) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *ModelScanResultsMaskRequest) SetScanId(v string)`

SetScanId sets ScanId field to given value.


### GetScanType

`func (o *ModelScanResultsMaskRequest) GetScanType() string`

GetScanType returns the ScanType field if non-nil, zero value otherwise.

### GetScanTypeOk

`func (o *ModelScanResultsMaskRequest) GetScanTypeOk() (*string, bool)`

GetScanTypeOk returns a tuple with the ScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanType

`func (o *ModelScanResultsMaskRequest) SetScanType(v string)`

SetScanType sets ScanType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


