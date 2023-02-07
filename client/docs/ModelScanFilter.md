# ModelScanFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerScanFilter** | [**ModelFieldsFilter**](ModelFieldsFilter.md) |  | 
**HostScanFilter** | [**ModelFieldsFilter**](ModelFieldsFilter.md) |  | 
**ImageScanFilter** | [**ModelFieldsFilter**](ModelFieldsFilter.md) |  | 

## Methods

### NewModelScanFilter

`func NewModelScanFilter(containerScanFilter ModelFieldsFilter, hostScanFilter ModelFieldsFilter, imageScanFilter ModelFieldsFilter, ) *ModelScanFilter`

NewModelScanFilter instantiates a new ModelScanFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelScanFilterWithDefaults

`func NewModelScanFilterWithDefaults() *ModelScanFilter`

NewModelScanFilterWithDefaults instantiates a new ModelScanFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerScanFilter

`func (o *ModelScanFilter) GetContainerScanFilter() ModelFieldsFilter`

GetContainerScanFilter returns the ContainerScanFilter field if non-nil, zero value otherwise.

### GetContainerScanFilterOk

`func (o *ModelScanFilter) GetContainerScanFilterOk() (*ModelFieldsFilter, bool)`

GetContainerScanFilterOk returns a tuple with the ContainerScanFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScanFilter

`func (o *ModelScanFilter) SetContainerScanFilter(v ModelFieldsFilter)`

SetContainerScanFilter sets ContainerScanFilter field to given value.


### GetHostScanFilter

`func (o *ModelScanFilter) GetHostScanFilter() ModelFieldsFilter`

GetHostScanFilter returns the HostScanFilter field if non-nil, zero value otherwise.

### GetHostScanFilterOk

`func (o *ModelScanFilter) GetHostScanFilterOk() (*ModelFieldsFilter, bool)`

GetHostScanFilterOk returns a tuple with the HostScanFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostScanFilter

`func (o *ModelScanFilter) SetHostScanFilter(v ModelFieldsFilter)`

SetHostScanFilter sets HostScanFilter field to given value.


### GetImageScanFilter

`func (o *ModelScanFilter) GetImageScanFilter() ModelFieldsFilter`

GetImageScanFilter returns the ImageScanFilter field if non-nil, zero value otherwise.

### GetImageScanFilterOk

`func (o *ModelScanFilter) GetImageScanFilterOk() (*ModelFieldsFilter, bool)`

GetImageScanFilterOk returns a tuple with the ImageScanFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageScanFilter

`func (o *ModelScanFilter) SetImageScanFilter(v ModelFieldsFilter)`

SetImageScanFilter sets ImageScanFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


