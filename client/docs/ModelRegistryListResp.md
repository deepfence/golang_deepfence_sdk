# ModelRegistryListResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IsSyncing** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**NonSecret** | Pointer to **interface{}** |  | [optional] 
**RegistryType** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **int32** |  | [optional] 

## Methods

### NewModelRegistryListResp

`func NewModelRegistryListResp() *ModelRegistryListResp`

NewModelRegistryListResp instantiates a new ModelRegistryListResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelRegistryListRespWithDefaults

`func NewModelRegistryListRespWithDefaults() *ModelRegistryListResp`

NewModelRegistryListRespWithDefaults instantiates a new ModelRegistryListResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ModelRegistryListResp) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelRegistryListResp) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelRegistryListResp) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ModelRegistryListResp) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *ModelRegistryListResp) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelRegistryListResp) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelRegistryListResp) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelRegistryListResp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsSyncing

`func (o *ModelRegistryListResp) GetIsSyncing() bool`

GetIsSyncing returns the IsSyncing field if non-nil, zero value otherwise.

### GetIsSyncingOk

`func (o *ModelRegistryListResp) GetIsSyncingOk() (*bool, bool)`

GetIsSyncingOk returns a tuple with the IsSyncing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncing

`func (o *ModelRegistryListResp) SetIsSyncing(v bool)`

SetIsSyncing sets IsSyncing field to given value.

### HasIsSyncing

`func (o *ModelRegistryListResp) HasIsSyncing() bool`

HasIsSyncing returns a boolean if a field has been set.

### GetName

`func (o *ModelRegistryListResp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelRegistryListResp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelRegistryListResp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelRegistryListResp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeId

`func (o *ModelRegistryListResp) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelRegistryListResp) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelRegistryListResp) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ModelRegistryListResp) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNonSecret

`func (o *ModelRegistryListResp) GetNonSecret() interface{}`

GetNonSecret returns the NonSecret field if non-nil, zero value otherwise.

### GetNonSecretOk

`func (o *ModelRegistryListResp) GetNonSecretOk() (*interface{}, bool)`

GetNonSecretOk returns a tuple with the NonSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSecret

`func (o *ModelRegistryListResp) SetNonSecret(v interface{})`

SetNonSecret sets NonSecret field to given value.

### HasNonSecret

`func (o *ModelRegistryListResp) HasNonSecret() bool`

HasNonSecret returns a boolean if a field has been set.

### SetNonSecretNil

`func (o *ModelRegistryListResp) SetNonSecretNil(b bool)`

 SetNonSecretNil sets the value for NonSecret to be an explicit nil

### UnsetNonSecret
`func (o *ModelRegistryListResp) UnsetNonSecret()`

UnsetNonSecret ensures that no value is present for NonSecret, not even an explicit nil
### GetRegistryType

`func (o *ModelRegistryListResp) GetRegistryType() string`

GetRegistryType returns the RegistryType field if non-nil, zero value otherwise.

### GetRegistryTypeOk

`func (o *ModelRegistryListResp) GetRegistryTypeOk() (*string, bool)`

GetRegistryTypeOk returns a tuple with the RegistryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryType

`func (o *ModelRegistryListResp) SetRegistryType(v string)`

SetRegistryType sets RegistryType field to given value.

### HasRegistryType

`func (o *ModelRegistryListResp) HasRegistryType() bool`

HasRegistryType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModelRegistryListResp) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelRegistryListResp) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelRegistryListResp) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModelRegistryListResp) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


