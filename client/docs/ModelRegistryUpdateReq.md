# ModelRegistryUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extras** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** |  | 
**NonSecret** | Pointer to **map[string]interface{}** |  | [optional] 
**RegistryType** | **string** |  | 
**Secret** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewModelRegistryUpdateReq

`func NewModelRegistryUpdateReq(name string, registryType string, ) *ModelRegistryUpdateReq`

NewModelRegistryUpdateReq instantiates a new ModelRegistryUpdateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelRegistryUpdateReqWithDefaults

`func NewModelRegistryUpdateReqWithDefaults() *ModelRegistryUpdateReq`

NewModelRegistryUpdateReqWithDefaults instantiates a new ModelRegistryUpdateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtras

`func (o *ModelRegistryUpdateReq) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *ModelRegistryUpdateReq) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *ModelRegistryUpdateReq) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *ModelRegistryUpdateReq) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *ModelRegistryUpdateReq) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *ModelRegistryUpdateReq) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetName

`func (o *ModelRegistryUpdateReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelRegistryUpdateReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelRegistryUpdateReq) SetName(v string)`

SetName sets Name field to given value.


### GetNonSecret

`func (o *ModelRegistryUpdateReq) GetNonSecret() map[string]interface{}`

GetNonSecret returns the NonSecret field if non-nil, zero value otherwise.

### GetNonSecretOk

`func (o *ModelRegistryUpdateReq) GetNonSecretOk() (*map[string]interface{}, bool)`

GetNonSecretOk returns a tuple with the NonSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSecret

`func (o *ModelRegistryUpdateReq) SetNonSecret(v map[string]interface{})`

SetNonSecret sets NonSecret field to given value.

### HasNonSecret

`func (o *ModelRegistryUpdateReq) HasNonSecret() bool`

HasNonSecret returns a boolean if a field has been set.

### SetNonSecretNil

`func (o *ModelRegistryUpdateReq) SetNonSecretNil(b bool)`

 SetNonSecretNil sets the value for NonSecret to be an explicit nil

### UnsetNonSecret
`func (o *ModelRegistryUpdateReq) UnsetNonSecret()`

UnsetNonSecret ensures that no value is present for NonSecret, not even an explicit nil
### GetRegistryType

`func (o *ModelRegistryUpdateReq) GetRegistryType() string`

GetRegistryType returns the RegistryType field if non-nil, zero value otherwise.

### GetRegistryTypeOk

`func (o *ModelRegistryUpdateReq) GetRegistryTypeOk() (*string, bool)`

GetRegistryTypeOk returns a tuple with the RegistryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryType

`func (o *ModelRegistryUpdateReq) SetRegistryType(v string)`

SetRegistryType sets RegistryType field to given value.


### GetSecret

`func (o *ModelRegistryUpdateReq) GetSecret() map[string]interface{}`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ModelRegistryUpdateReq) GetSecretOk() (*map[string]interface{}, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ModelRegistryUpdateReq) SetSecret(v map[string]interface{})`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ModelRegistryUpdateReq) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### SetSecretNil

`func (o *ModelRegistryUpdateReq) SetSecretNil(b bool)`

 SetSecretNil sets the value for Secret to be an explicit nil

### UnsetSecret
`func (o *ModelRegistryUpdateReq) UnsetSecret()`

UnsetSecret ensures that no value is present for Secret, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


