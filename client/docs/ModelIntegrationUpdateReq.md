# ModelIntegrationUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Filters** | Pointer to [**ModelIntegrationFilters**](ModelIntegrationFilters.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IntegrationType** | Pointer to **string** |  | [optional] 
**NotificationType** | Pointer to **string** |  | [optional] 

## Methods

### NewModelIntegrationUpdateReq

`func NewModelIntegrationUpdateReq() *ModelIntegrationUpdateReq`

NewModelIntegrationUpdateReq instantiates a new ModelIntegrationUpdateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelIntegrationUpdateReqWithDefaults

`func NewModelIntegrationUpdateReqWithDefaults() *ModelIntegrationUpdateReq`

NewModelIntegrationUpdateReqWithDefaults instantiates a new ModelIntegrationUpdateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *ModelIntegrationUpdateReq) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ModelIntegrationUpdateReq) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ModelIntegrationUpdateReq) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ModelIntegrationUpdateReq) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *ModelIntegrationUpdateReq) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *ModelIntegrationUpdateReq) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetFilters

`func (o *ModelIntegrationUpdateReq) GetFilters() ModelIntegrationFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ModelIntegrationUpdateReq) GetFiltersOk() (*ModelIntegrationFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ModelIntegrationUpdateReq) SetFilters(v ModelIntegrationFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ModelIntegrationUpdateReq) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetId

`func (o *ModelIntegrationUpdateReq) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelIntegrationUpdateReq) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelIntegrationUpdateReq) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelIntegrationUpdateReq) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIntegrationType

`func (o *ModelIntegrationUpdateReq) GetIntegrationType() string`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *ModelIntegrationUpdateReq) GetIntegrationTypeOk() (*string, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *ModelIntegrationUpdateReq) SetIntegrationType(v string)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *ModelIntegrationUpdateReq) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### GetNotificationType

`func (o *ModelIntegrationUpdateReq) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *ModelIntegrationUpdateReq) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *ModelIntegrationUpdateReq) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.

### HasNotificationType

`func (o *ModelIntegrationUpdateReq) HasNotificationType() bool`

HasNotificationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


