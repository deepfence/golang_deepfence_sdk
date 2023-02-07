# ModelScanStatusResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statuses** | [**map[string]ModelScanInfo**](ModelScanInfo.md) |  | 

## Methods

### NewModelScanStatusResp

`func NewModelScanStatusResp(statuses map[string]ModelScanInfo, ) *ModelScanStatusResp`

NewModelScanStatusResp instantiates a new ModelScanStatusResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelScanStatusRespWithDefaults

`func NewModelScanStatusRespWithDefaults() *ModelScanStatusResp`

NewModelScanStatusRespWithDefaults instantiates a new ModelScanStatusResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatuses

`func (o *ModelScanStatusResp) GetStatuses() map[string]ModelScanInfo`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *ModelScanStatusResp) GetStatusesOk() (*map[string]ModelScanInfo, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *ModelScanStatusResp) SetStatuses(v map[string]ModelScanInfo)`

SetStatuses sets Statuses field to given value.


### SetStatusesNil

`func (o *ModelScanStatusResp) SetStatusesNil(b bool)`

 SetStatusesNil sets the value for Statuses to be an explicit nil

### UnsetStatuses
`func (o *ModelScanStatusResp) UnsetStatuses()`

UnsetStatuses ensures that no value is present for Statuses, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


