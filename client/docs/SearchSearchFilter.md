# SearchSearchFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**ReportersFieldsFilters**](ReportersFieldsFilters.md) |  | 
**InFieldFilter** | **[]string** |  | 
**Window** | [**ModelFetchWindow**](ModelFetchWindow.md) |  | 

## Methods

### NewSearchSearchFilter

`func NewSearchSearchFilter(filters ReportersFieldsFilters, inFieldFilter []string, window ModelFetchWindow, ) *SearchSearchFilter`

NewSearchSearchFilter instantiates a new SearchSearchFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSearchFilterWithDefaults

`func NewSearchSearchFilterWithDefaults() *SearchSearchFilter`

NewSearchSearchFilterWithDefaults instantiates a new SearchSearchFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *SearchSearchFilter) GetFilters() ReportersFieldsFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SearchSearchFilter) GetFiltersOk() (*ReportersFieldsFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SearchSearchFilter) SetFilters(v ReportersFieldsFilters)`

SetFilters sets Filters field to given value.


### GetInFieldFilter

`func (o *SearchSearchFilter) GetInFieldFilter() []string`

GetInFieldFilter returns the InFieldFilter field if non-nil, zero value otherwise.

### GetInFieldFilterOk

`func (o *SearchSearchFilter) GetInFieldFilterOk() (*[]string, bool)`

GetInFieldFilterOk returns a tuple with the InFieldFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInFieldFilter

`func (o *SearchSearchFilter) SetInFieldFilter(v []string)`

SetInFieldFilter sets InFieldFilter field to given value.


### SetInFieldFilterNil

`func (o *SearchSearchFilter) SetInFieldFilterNil(b bool)`

 SetInFieldFilterNil sets the value for InFieldFilter to be an explicit nil

### UnsetInFieldFilter
`func (o *SearchSearchFilter) UnsetInFieldFilter()`

UnsetInFieldFilter ensures that no value is present for InFieldFilter, not even an explicit nil
### GetWindow

`func (o *SearchSearchFilter) GetWindow() ModelFetchWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *SearchSearchFilter) GetWindowOk() (*ModelFetchWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *SearchSearchFilter) SetWindow(v ModelFetchWindow)`

SetWindow sets Window field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


