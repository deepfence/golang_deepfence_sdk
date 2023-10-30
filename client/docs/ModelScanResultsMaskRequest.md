# ModelScanResultsMaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaskAction** | **string** |  | 
**ResultIds** | **[]string** |  | 
**ScanId** | **string** |  | 
**ScanType** | **string** |  | 

## Methods

### NewModelScanResultsMaskRequest

`func NewModelScanResultsMaskRequest(maskAction string, resultIds []string, scanId string, scanType string, ) *ModelScanResultsMaskRequest`

NewModelScanResultsMaskRequest instantiates a new ModelScanResultsMaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelScanResultsMaskRequestWithDefaults

`func NewModelScanResultsMaskRequestWithDefaults() *ModelScanResultsMaskRequest`

NewModelScanResultsMaskRequestWithDefaults instantiates a new ModelScanResultsMaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaskAction

`func (o *ModelScanResultsMaskRequest) GetMaskAction() string`

GetMaskAction returns the MaskAction field if non-nil, zero value otherwise.

### GetMaskActionOk

`func (o *ModelScanResultsMaskRequest) GetMaskActionOk() (*string, bool)`

GetMaskActionOk returns a tuple with the MaskAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskAction

`func (o *ModelScanResultsMaskRequest) SetMaskAction(v string)`

SetMaskAction sets MaskAction field to given value.


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


