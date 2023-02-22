# ModelScanStatusReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BulkScanId** | **string** |  | 
**ScanIds** | **[]string** |  | 

## Methods

### NewModelScanStatusReq

`func NewModelScanStatusReq(bulkScanId string, scanIds []string, ) *ModelScanStatusReq`

NewModelScanStatusReq instantiates a new ModelScanStatusReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelScanStatusReqWithDefaults

`func NewModelScanStatusReqWithDefaults() *ModelScanStatusReq`

NewModelScanStatusReqWithDefaults instantiates a new ModelScanStatusReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBulkScanId

`func (o *ModelScanStatusReq) GetBulkScanId() string`

GetBulkScanId returns the BulkScanId field if non-nil, zero value otherwise.

### GetBulkScanIdOk

`func (o *ModelScanStatusReq) GetBulkScanIdOk() (*string, bool)`

GetBulkScanIdOk returns a tuple with the BulkScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkScanId

`func (o *ModelScanStatusReq) SetBulkScanId(v string)`

SetBulkScanId sets BulkScanId field to given value.


### GetScanIds

`func (o *ModelScanStatusReq) GetScanIds() []string`

GetScanIds returns the ScanIds field if non-nil, zero value otherwise.

### GetScanIdsOk

`func (o *ModelScanStatusReq) GetScanIdsOk() (*[]string, bool)`

GetScanIdsOk returns a tuple with the ScanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanIds

`func (o *ModelScanStatusReq) SetScanIds(v []string)`

SetScanIds sets ScanIds field to given value.


### SetScanIdsNil

`func (o *ModelScanStatusReq) SetScanIdsNil(b bool)`

 SetScanIdsNil sets the value for ScanIds to be an explicit nil

### UnsetScanIds
`func (o *ModelScanStatusReq) UnsetScanIds()`

UnsetScanIds ensures that no value is present for ScanIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


