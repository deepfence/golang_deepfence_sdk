# ModelScanResultsReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldsFilter** | [**ReportersFieldsFilters**](ReportersFieldsFilters.md) |  | 
**ScanId** | **string** |  | 
**Window** | [**ModelFetchWindow**](ModelFetchWindow.md) |  | 

## Methods

### NewModelScanResultsReq

`func NewModelScanResultsReq(fieldsFilter ReportersFieldsFilters, scanId string, window ModelFetchWindow, ) *ModelScanResultsReq`

NewModelScanResultsReq instantiates a new ModelScanResultsReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelScanResultsReqWithDefaults

`func NewModelScanResultsReqWithDefaults() *ModelScanResultsReq`

NewModelScanResultsReqWithDefaults instantiates a new ModelScanResultsReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldsFilter

`func (o *ModelScanResultsReq) GetFieldsFilter() ReportersFieldsFilters`

GetFieldsFilter returns the FieldsFilter field if non-nil, zero value otherwise.

### GetFieldsFilterOk

`func (o *ModelScanResultsReq) GetFieldsFilterOk() (*ReportersFieldsFilters, bool)`

GetFieldsFilterOk returns a tuple with the FieldsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsFilter

`func (o *ModelScanResultsReq) SetFieldsFilter(v ReportersFieldsFilters)`

SetFieldsFilter sets FieldsFilter field to given value.


### GetScanId

`func (o *ModelScanResultsReq) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *ModelScanResultsReq) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *ModelScanResultsReq) SetScanId(v string)`

SetScanId sets ScanId field to given value.


### GetWindow

`func (o *ModelScanResultsReq) GetWindow() ModelFetchWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *ModelScanResultsReq) GetWindowOk() (*ModelFetchWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *ModelScanResultsReq) SetWindow(v ModelFetchWindow)`

SetWindow sets Window field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


