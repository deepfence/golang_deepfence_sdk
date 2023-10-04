# ModelTopologyDeltaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionTimestamp** | Pointer to **int64** |  | [optional] 
**Additons** | Pointer to [**[]ModelNodeIdentifier**](ModelNodeIdentifier.md) |  | [optional] 
**DeletionTimestamp** | Pointer to **int64** |  | [optional] 
**Deletions** | Pointer to [**[]ModelNodeIdentifier**](ModelNodeIdentifier.md) |  | [optional] 

## Methods

### NewModelTopologyDeltaResponse

`func NewModelTopologyDeltaResponse() *ModelTopologyDeltaResponse`

NewModelTopologyDeltaResponse instantiates a new ModelTopologyDeltaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelTopologyDeltaResponseWithDefaults

`func NewModelTopologyDeltaResponseWithDefaults() *ModelTopologyDeltaResponse`

NewModelTopologyDeltaResponseWithDefaults instantiates a new ModelTopologyDeltaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionTimestamp

`func (o *ModelTopologyDeltaResponse) GetAdditionTimestamp() int64`

GetAdditionTimestamp returns the AdditionTimestamp field if non-nil, zero value otherwise.

### GetAdditionTimestampOk

`func (o *ModelTopologyDeltaResponse) GetAdditionTimestampOk() (*int64, bool)`

GetAdditionTimestampOk returns a tuple with the AdditionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionTimestamp

`func (o *ModelTopologyDeltaResponse) SetAdditionTimestamp(v int64)`

SetAdditionTimestamp sets AdditionTimestamp field to given value.

### HasAdditionTimestamp

`func (o *ModelTopologyDeltaResponse) HasAdditionTimestamp() bool`

HasAdditionTimestamp returns a boolean if a field has been set.

### GetAdditons

`func (o *ModelTopologyDeltaResponse) GetAdditons() []ModelNodeIdentifier`

GetAdditons returns the Additons field if non-nil, zero value otherwise.

### GetAdditonsOk

`func (o *ModelTopologyDeltaResponse) GetAdditonsOk() (*[]ModelNodeIdentifier, bool)`

GetAdditonsOk returns a tuple with the Additons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditons

`func (o *ModelTopologyDeltaResponse) SetAdditons(v []ModelNodeIdentifier)`

SetAdditons sets Additons field to given value.

### HasAdditons

`func (o *ModelTopologyDeltaResponse) HasAdditons() bool`

HasAdditons returns a boolean if a field has been set.

### SetAdditonsNil

`func (o *ModelTopologyDeltaResponse) SetAdditonsNil(b bool)`

 SetAdditonsNil sets the value for Additons to be an explicit nil

### UnsetAdditons
`func (o *ModelTopologyDeltaResponse) UnsetAdditons()`

UnsetAdditons ensures that no value is present for Additons, not even an explicit nil
### GetDeletionTimestamp

`func (o *ModelTopologyDeltaResponse) GetDeletionTimestamp() int64`

GetDeletionTimestamp returns the DeletionTimestamp field if non-nil, zero value otherwise.

### GetDeletionTimestampOk

`func (o *ModelTopologyDeltaResponse) GetDeletionTimestampOk() (*int64, bool)`

GetDeletionTimestampOk returns a tuple with the DeletionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTimestamp

`func (o *ModelTopologyDeltaResponse) SetDeletionTimestamp(v int64)`

SetDeletionTimestamp sets DeletionTimestamp field to given value.

### HasDeletionTimestamp

`func (o *ModelTopologyDeltaResponse) HasDeletionTimestamp() bool`

HasDeletionTimestamp returns a boolean if a field has been set.

### GetDeletions

`func (o *ModelTopologyDeltaResponse) GetDeletions() []ModelNodeIdentifier`

GetDeletions returns the Deletions field if non-nil, zero value otherwise.

### GetDeletionsOk

`func (o *ModelTopologyDeltaResponse) GetDeletionsOk() (*[]ModelNodeIdentifier, bool)`

GetDeletionsOk returns a tuple with the Deletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletions

`func (o *ModelTopologyDeltaResponse) SetDeletions(v []ModelNodeIdentifier)`

SetDeletions sets Deletions field to given value.

### HasDeletions

`func (o *ModelTopologyDeltaResponse) HasDeletions() bool`

HasDeletions returns a boolean if a field has been set.

### SetDeletionsNil

`func (o *ModelTopologyDeltaResponse) SetDeletionsNil(b bool)`

 SetDeletionsNil sets the value for Deletions to be an explicit nil

### UnsetDeletions
`func (o *ModelTopologyDeltaResponse) UnsetDeletions()`

UnsetDeletions ensures that no value is present for Deletions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


