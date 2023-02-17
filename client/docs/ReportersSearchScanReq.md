# ReportersSearchScanReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeFilters** | [**ReportersSearchFilter**](ReportersSearchFilter.md) |  | 
**ScanFilters** | [**ReportersSearchFilter**](ReportersSearchFilter.md) |  | 
**Window** | [**ModelFetchWindow**](ModelFetchWindow.md) |  | 

## Methods

### NewReportersSearchScanReq

`func NewReportersSearchScanReq(nodeFilters ReportersSearchFilter, scanFilters ReportersSearchFilter, window ModelFetchWindow, ) *ReportersSearchScanReq`

NewReportersSearchScanReq instantiates a new ReportersSearchScanReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportersSearchScanReqWithDefaults

`func NewReportersSearchScanReqWithDefaults() *ReportersSearchScanReq`

NewReportersSearchScanReqWithDefaults instantiates a new ReportersSearchScanReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeFilters

`func (o *ReportersSearchScanReq) GetNodeFilters() ReportersSearchFilter`

GetNodeFilters returns the NodeFilters field if non-nil, zero value otherwise.

### GetNodeFiltersOk

`func (o *ReportersSearchScanReq) GetNodeFiltersOk() (*ReportersSearchFilter, bool)`

GetNodeFiltersOk returns a tuple with the NodeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeFilters

`func (o *ReportersSearchScanReq) SetNodeFilters(v ReportersSearchFilter)`

SetNodeFilters sets NodeFilters field to given value.


### GetScanFilters

`func (o *ReportersSearchScanReq) GetScanFilters() ReportersSearchFilter`

GetScanFilters returns the ScanFilters field if non-nil, zero value otherwise.

### GetScanFiltersOk

`func (o *ReportersSearchScanReq) GetScanFiltersOk() (*ReportersSearchFilter, bool)`

GetScanFiltersOk returns a tuple with the ScanFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanFilters

`func (o *ReportersSearchScanReq) SetScanFilters(v ReportersSearchFilter)`

SetScanFilters sets ScanFilters field to given value.


### GetWindow

`func (o *ReportersSearchScanReq) GetWindow() ModelFetchWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *ReportersSearchScanReq) GetWindowOk() (*ModelFetchWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *ReportersSearchScanReq) SetWindow(v ModelFetchWindow)`

SetWindow sets Window field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


