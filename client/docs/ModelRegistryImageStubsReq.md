# ModelRegistryImageStubsReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageFilter** | [**ReportersFieldsFilters**](ReportersFieldsFilters.md) |  | 
**RegistryId** | **string** |  | 
**Window** | [**ModelFetchWindow**](ModelFetchWindow.md) |  | 

## Methods

### NewModelRegistryImageStubsReq

`func NewModelRegistryImageStubsReq(imageFilter ReportersFieldsFilters, registryId string, window ModelFetchWindow, ) *ModelRegistryImageStubsReq`

NewModelRegistryImageStubsReq instantiates a new ModelRegistryImageStubsReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelRegistryImageStubsReqWithDefaults

`func NewModelRegistryImageStubsReqWithDefaults() *ModelRegistryImageStubsReq`

NewModelRegistryImageStubsReqWithDefaults instantiates a new ModelRegistryImageStubsReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageFilter

`func (o *ModelRegistryImageStubsReq) GetImageFilter() ReportersFieldsFilters`

GetImageFilter returns the ImageFilter field if non-nil, zero value otherwise.

### GetImageFilterOk

`func (o *ModelRegistryImageStubsReq) GetImageFilterOk() (*ReportersFieldsFilters, bool)`

GetImageFilterOk returns a tuple with the ImageFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFilter

`func (o *ModelRegistryImageStubsReq) SetImageFilter(v ReportersFieldsFilters)`

SetImageFilter sets ImageFilter field to given value.


### GetRegistryId

`func (o *ModelRegistryImageStubsReq) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *ModelRegistryImageStubsReq) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *ModelRegistryImageStubsReq) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.


### GetWindow

`func (o *ModelRegistryImageStubsReq) GetWindow() ModelFetchWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *ModelRegistryImageStubsReq) GetWindowOk() (*ModelFetchWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *ModelRegistryImageStubsReq) SetWindow(v ModelFetchWindow)`

SetWindow sets Window field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


