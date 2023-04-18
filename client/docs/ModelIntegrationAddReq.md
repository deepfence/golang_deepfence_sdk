# ModelIntegrationAddReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Filters** | Pointer to [**ModelIntegrationFilters**](ModelIntegrationFilters.md) |  | [optional] 
**IntegrationType** | Pointer to **string** |  | [optional] 
**NotificationType** | Pointer to **string** |  | [optional] 

## Methods

### NewModelIntegrationAddReq

`func NewModelIntegrationAddReq() *ModelIntegrationAddReq`

NewModelIntegrationAddReq instantiates a new ModelIntegrationAddReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelIntegrationAddReqWithDefaults

`func NewModelIntegrationAddReqWithDefaults() *ModelIntegrationAddReq`

NewModelIntegrationAddReqWithDefaults instantiates a new ModelIntegrationAddReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *ModelIntegrationAddReq) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ModelIntegrationAddReq) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ModelIntegrationAddReq) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ModelIntegrationAddReq) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *ModelIntegrationAddReq) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *ModelIntegrationAddReq) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetFilters

`func (o *ModelIntegrationAddReq) GetFilters() ModelIntegrationFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ModelIntegrationAddReq) GetFiltersOk() (*ModelIntegrationFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ModelIntegrationAddReq) SetFilters(v ModelIntegrationFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ModelIntegrationAddReq) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetIntegrationType

`func (o *ModelIntegrationAddReq) GetIntegrationType() string`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *ModelIntegrationAddReq) GetIntegrationTypeOk() (*string, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *ModelIntegrationAddReq) SetIntegrationType(v string)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *ModelIntegrationAddReq) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### GetNotificationType

`func (o *ModelIntegrationAddReq) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *ModelIntegrationAddReq) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *ModelIntegrationAddReq) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.

### HasNotificationType

`func (o *ModelIntegrationAddReq) HasNotificationType() bool`

HasNotificationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


