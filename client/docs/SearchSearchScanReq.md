# SearchSearchScanReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeFilters** | [**SearchSearchFilter**](SearchSearchFilter.md) |  | 
**ScanFilters** | [**SearchSearchFilter**](SearchSearchFilter.md) |  | 
**Window** | [**ModelFetchWindow**](ModelFetchWindow.md) |  | 

## Methods

### NewSearchSearchScanReq

`func NewSearchSearchScanReq(nodeFilters SearchSearchFilter, scanFilters SearchSearchFilter, window ModelFetchWindow, ) *SearchSearchScanReq`

NewSearchSearchScanReq instantiates a new SearchSearchScanReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSearchScanReqWithDefaults

`func NewSearchSearchScanReqWithDefaults() *SearchSearchScanReq`

NewSearchSearchScanReqWithDefaults instantiates a new SearchSearchScanReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeFilters

`func (o *SearchSearchScanReq) GetNodeFilters() SearchSearchFilter`

GetNodeFilters returns the NodeFilters field if non-nil, zero value otherwise.

### GetNodeFiltersOk

`func (o *SearchSearchScanReq) GetNodeFiltersOk() (*SearchSearchFilter, bool)`

GetNodeFiltersOk returns a tuple with the NodeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeFilters

`func (o *SearchSearchScanReq) SetNodeFilters(v SearchSearchFilter)`

SetNodeFilters sets NodeFilters field to given value.


### GetScanFilters

`func (o *SearchSearchScanReq) GetScanFilters() SearchSearchFilter`

GetScanFilters returns the ScanFilters field if non-nil, zero value otherwise.

### GetScanFiltersOk

`func (o *SearchSearchScanReq) GetScanFiltersOk() (*SearchSearchFilter, bool)`

GetScanFiltersOk returns a tuple with the ScanFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanFilters

`func (o *SearchSearchScanReq) SetScanFilters(v SearchSearchFilter)`

SetScanFilters sets ScanFilters field to given value.


### GetWindow

`func (o *SearchSearchScanReq) GetWindow() ModelFetchWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *SearchSearchScanReq) GetWindowOk() (*ModelFetchWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *SearchSearchScanReq) SetWindow(v ModelFetchWindow)`

SetWindow sets Window field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


