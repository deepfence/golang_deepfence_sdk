# ModelScanCompareReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseScanId** | **string** |  | 
**FieldsFilter** | [**ReportersFieldsFilters**](ReportersFieldsFilters.md) |  | 
**ToScanId** | **string** |  | 
**Window** | [**ModelFetchWindow**](ModelFetchWindow.md) |  | 

## Methods

### NewModelScanCompareReq

`func NewModelScanCompareReq(baseScanId string, fieldsFilter ReportersFieldsFilters, toScanId string, window ModelFetchWindow, ) *ModelScanCompareReq`

NewModelScanCompareReq instantiates a new ModelScanCompareReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelScanCompareReqWithDefaults

`func NewModelScanCompareReqWithDefaults() *ModelScanCompareReq`

NewModelScanCompareReqWithDefaults instantiates a new ModelScanCompareReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseScanId

`func (o *ModelScanCompareReq) GetBaseScanId() string`

GetBaseScanId returns the BaseScanId field if non-nil, zero value otherwise.

### GetBaseScanIdOk

`func (o *ModelScanCompareReq) GetBaseScanIdOk() (*string, bool)`

GetBaseScanIdOk returns a tuple with the BaseScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScanId

`func (o *ModelScanCompareReq) SetBaseScanId(v string)`

SetBaseScanId sets BaseScanId field to given value.


### GetFieldsFilter

`func (o *ModelScanCompareReq) GetFieldsFilter() ReportersFieldsFilters`

GetFieldsFilter returns the FieldsFilter field if non-nil, zero value otherwise.

### GetFieldsFilterOk

`func (o *ModelScanCompareReq) GetFieldsFilterOk() (*ReportersFieldsFilters, bool)`

GetFieldsFilterOk returns a tuple with the FieldsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsFilter

`func (o *ModelScanCompareReq) SetFieldsFilter(v ReportersFieldsFilters)`

SetFieldsFilter sets FieldsFilter field to given value.


### GetToScanId

`func (o *ModelScanCompareReq) GetToScanId() string`

GetToScanId returns the ToScanId field if non-nil, zero value otherwise.

### GetToScanIdOk

`func (o *ModelScanCompareReq) GetToScanIdOk() (*string, bool)`

GetToScanIdOk returns a tuple with the ToScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToScanId

`func (o *ModelScanCompareReq) SetToScanId(v string)`

SetToScanId sets ToScanId field to given value.


### GetWindow

`func (o *ModelScanCompareReq) GetWindow() ModelFetchWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *ModelScanCompareReq) GetWindowOk() (*ModelFetchWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *ModelScanCompareReq) SetWindow(v ModelFetchWindow)`

SetWindow sets Window field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


