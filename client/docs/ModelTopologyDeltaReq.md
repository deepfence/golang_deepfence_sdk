# ModelTopologyDeltaReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addition** | **bool** |  | 
**AdditionTimestamp** | **int64** |  | 
**Deletion** | **bool** |  | 
**DeletionTimestamp** | **int64** |  | 
**EntityTypes** | **[]string** |  | 

## Methods

### NewModelTopologyDeltaReq

`func NewModelTopologyDeltaReq(addition bool, additionTimestamp int64, deletion bool, deletionTimestamp int64, entityTypes []string, ) *ModelTopologyDeltaReq`

NewModelTopologyDeltaReq instantiates a new ModelTopologyDeltaReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelTopologyDeltaReqWithDefaults

`func NewModelTopologyDeltaReqWithDefaults() *ModelTopologyDeltaReq`

NewModelTopologyDeltaReqWithDefaults instantiates a new ModelTopologyDeltaReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddition

`func (o *ModelTopologyDeltaReq) GetAddition() bool`

GetAddition returns the Addition field if non-nil, zero value otherwise.

### GetAdditionOk

`func (o *ModelTopologyDeltaReq) GetAdditionOk() (*bool, bool)`

GetAdditionOk returns a tuple with the Addition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddition

`func (o *ModelTopologyDeltaReq) SetAddition(v bool)`

SetAddition sets Addition field to given value.


### GetAdditionTimestamp

`func (o *ModelTopologyDeltaReq) GetAdditionTimestamp() int64`

GetAdditionTimestamp returns the AdditionTimestamp field if non-nil, zero value otherwise.

### GetAdditionTimestampOk

`func (o *ModelTopologyDeltaReq) GetAdditionTimestampOk() (*int64, bool)`

GetAdditionTimestampOk returns a tuple with the AdditionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionTimestamp

`func (o *ModelTopologyDeltaReq) SetAdditionTimestamp(v int64)`

SetAdditionTimestamp sets AdditionTimestamp field to given value.


### GetDeletion

`func (o *ModelTopologyDeltaReq) GetDeletion() bool`

GetDeletion returns the Deletion field if non-nil, zero value otherwise.

### GetDeletionOk

`func (o *ModelTopologyDeltaReq) GetDeletionOk() (*bool, bool)`

GetDeletionOk returns a tuple with the Deletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletion

`func (o *ModelTopologyDeltaReq) SetDeletion(v bool)`

SetDeletion sets Deletion field to given value.


### GetDeletionTimestamp

`func (o *ModelTopologyDeltaReq) GetDeletionTimestamp() int64`

GetDeletionTimestamp returns the DeletionTimestamp field if non-nil, zero value otherwise.

### GetDeletionTimestampOk

`func (o *ModelTopologyDeltaReq) GetDeletionTimestampOk() (*int64, bool)`

GetDeletionTimestampOk returns a tuple with the DeletionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTimestamp

`func (o *ModelTopologyDeltaReq) SetDeletionTimestamp(v int64)`

SetDeletionTimestamp sets DeletionTimestamp field to given value.


### GetEntityTypes

`func (o *ModelTopologyDeltaReq) GetEntityTypes() []string`

GetEntityTypes returns the EntityTypes field if non-nil, zero value otherwise.

### GetEntityTypesOk

`func (o *ModelTopologyDeltaReq) GetEntityTypesOk() (*[]string, bool)`

GetEntityTypesOk returns a tuple with the EntityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypes

`func (o *ModelTopologyDeltaReq) SetEntityTypes(v []string)`

SetEntityTypes sets EntityTypes field to given value.


### SetEntityTypesNil

`func (o *ModelTopologyDeltaReq) SetEntityTypesNil(b bool)`

 SetEntityTypesNil sets the value for EntityTypes to be an explicit nil

### UnsetEntityTypes
`func (o *ModelTopologyDeltaReq) UnsetEntityTypes()`

UnsetEntityTypes ensures that no value is present for EntityTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


