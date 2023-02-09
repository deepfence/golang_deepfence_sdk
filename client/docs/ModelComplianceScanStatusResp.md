# ModelComplianceScanStatusResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statuses** | [**[]ModelComplianceScanInfo**](ModelComplianceScanInfo.md) |  | 

## Methods

### NewModelComplianceScanStatusResp

`func NewModelComplianceScanStatusResp(statuses []ModelComplianceScanInfo, ) *ModelComplianceScanStatusResp`

NewModelComplianceScanStatusResp instantiates a new ModelComplianceScanStatusResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelComplianceScanStatusRespWithDefaults

`func NewModelComplianceScanStatusRespWithDefaults() *ModelComplianceScanStatusResp`

NewModelComplianceScanStatusRespWithDefaults instantiates a new ModelComplianceScanStatusResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatuses

`func (o *ModelComplianceScanStatusResp) GetStatuses() []ModelComplianceScanInfo`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *ModelComplianceScanStatusResp) GetStatusesOk() (*[]ModelComplianceScanInfo, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *ModelComplianceScanStatusResp) SetStatuses(v []ModelComplianceScanInfo)`

SetStatuses sets Statuses field to given value.


### SetStatusesNil

`func (o *ModelComplianceScanStatusResp) SetStatusesNil(b bool)`

 SetStatusesNil sets the value for Statuses to be an explicit nil

### UnsetStatuses
`func (o *ModelComplianceScanStatusResp) UnsetStatuses()`

UnsetStatuses ensures that no value is present for Statuses, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


