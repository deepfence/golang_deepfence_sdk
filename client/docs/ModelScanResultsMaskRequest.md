# ModelScanResultsMaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaskAcrossHostsAndImages** | Pointer to **bool** |  | [optional] 
**NodeIds** | **[]string** |  | 
**ScanId** | **string** |  | 
**ScanType** | **string** |  | 

## Methods

### NewModelScanResultsMaskRequest

`func NewModelScanResultsMaskRequest(nodeIds []string, scanId string, scanType string, ) *ModelScanResultsMaskRequest`

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

### GetNodeIds

`func (o *ModelScanResultsMaskRequest) GetNodeIds() []string`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *ModelScanResultsMaskRequest) GetNodeIdsOk() (*[]string, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *ModelScanResultsMaskRequest) SetNodeIds(v []string)`

SetNodeIds sets NodeIds field to given value.


### SetNodeIdsNil

`func (o *ModelScanResultsMaskRequest) SetNodeIdsNil(b bool)`

 SetNodeIdsNil sets the value for NodeIds to be an explicit nil

### UnsetNodeIds
`func (o *ModelScanResultsMaskRequest) UnsetNodeIds()`

UnsetNodeIds ensures that no value is present for NodeIds, not even an explicit nil
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


