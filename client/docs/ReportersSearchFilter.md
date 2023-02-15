# ReportersSearchFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**ReportersFieldsFilters**](ReportersFieldsFilters.md) |  | 
**InFieldFilter** | **[]string** |  | 

## Methods

### NewReportersSearchFilter

`func NewReportersSearchFilter(filters ReportersFieldsFilters, inFieldFilter []string, ) *ReportersSearchFilter`

NewReportersSearchFilter instantiates a new ReportersSearchFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportersSearchFilterWithDefaults

`func NewReportersSearchFilterWithDefaults() *ReportersSearchFilter`

NewReportersSearchFilterWithDefaults instantiates a new ReportersSearchFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ReportersSearchFilter) GetFilters() ReportersFieldsFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ReportersSearchFilter) GetFiltersOk() (*ReportersFieldsFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ReportersSearchFilter) SetFilters(v ReportersFieldsFilters)`

SetFilters sets Filters field to given value.


### GetInFieldFilter

`func (o *ReportersSearchFilter) GetInFieldFilter() []string`

GetInFieldFilter returns the InFieldFilter field if non-nil, zero value otherwise.

### GetInFieldFilterOk

`func (o *ReportersSearchFilter) GetInFieldFilterOk() (*[]string, bool)`

GetInFieldFilterOk returns a tuple with the InFieldFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInFieldFilter

`func (o *ReportersSearchFilter) SetInFieldFilter(v []string)`

SetInFieldFilter sets InFieldFilter field to given value.


### SetInFieldFilterNil

`func (o *ReportersSearchFilter) SetInFieldFilterNil(b bool)`

 SetInFieldFilterNil sets the value for InFieldFilter to be an explicit nil

### UnsetInFieldFilter
`func (o *ReportersSearchFilter) UnsetInFieldFilter()`

UnsetInFieldFilter ensures that no value is present for InFieldFilter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


