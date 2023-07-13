# ModelRegistryImagesReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageFilter** | [**ReportersFieldsFilters**](ReportersFieldsFilters.md) |  | 
**RegistryId** | **string** |  | 
**Window** | [**ModelFetchWindow**](ModelFetchWindow.md) |  | 

## Methods

### NewModelRegistryImagesReq

`func NewModelRegistryImagesReq(imageFilter ReportersFieldsFilters, registryId string, window ModelFetchWindow, ) *ModelRegistryImagesReq`

NewModelRegistryImagesReq instantiates a new ModelRegistryImagesReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelRegistryImagesReqWithDefaults

`func NewModelRegistryImagesReqWithDefaults() *ModelRegistryImagesReq`

NewModelRegistryImagesReqWithDefaults instantiates a new ModelRegistryImagesReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageFilter

`func (o *ModelRegistryImagesReq) GetImageFilter() ReportersFieldsFilters`

GetImageFilter returns the ImageFilter field if non-nil, zero value otherwise.

### GetImageFilterOk

`func (o *ModelRegistryImagesReq) GetImageFilterOk() (*ReportersFieldsFilters, bool)`

GetImageFilterOk returns a tuple with the ImageFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFilter

`func (o *ModelRegistryImagesReq) SetImageFilter(v ReportersFieldsFilters)`

SetImageFilter sets ImageFilter field to given value.


### GetRegistryId

`func (o *ModelRegistryImagesReq) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *ModelRegistryImagesReq) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *ModelRegistryImagesReq) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.


### GetWindow

`func (o *ModelRegistryImagesReq) GetWindow() ModelFetchWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *ModelRegistryImagesReq) GetWindowOk() (*ModelFetchWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *ModelRegistryImagesReq) SetWindow(v ModelFetchWindow)`

SetWindow sets Window field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


