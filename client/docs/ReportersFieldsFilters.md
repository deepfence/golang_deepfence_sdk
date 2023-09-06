# ReportersFieldsFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompareFilter** | [**[]ReportersCompareFilter**](ReportersCompareFilter.md) |  | 
**ContainsFilter** | [**ReportersContainsFilter**](ReportersContainsFilter.md) |  | 
**ContainsInArrayFilter** | Pointer to [**ReportersContainsFilter**](ReportersContainsFilter.md) |  | [optional] 
**MatchFilter** | [**ReportersMatchFilter**](ReportersMatchFilter.md) |  | 
**MatchInArrayFilter** | Pointer to [**ReportersMatchFilter**](ReportersMatchFilter.md) |  | [optional] 
**NotContainsFilter** | Pointer to [**ReportersContainsFilter**](ReportersContainsFilter.md) |  | [optional] 
**OrderFilter** | [**ReportersOrderFilter**](ReportersOrderFilter.md) |  | 

## Methods

### NewReportersFieldsFilters

`func NewReportersFieldsFilters(compareFilter []ReportersCompareFilter, containsFilter ReportersContainsFilter, matchFilter ReportersMatchFilter, orderFilter ReportersOrderFilter, ) *ReportersFieldsFilters`

NewReportersFieldsFilters instantiates a new ReportersFieldsFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportersFieldsFiltersWithDefaults

`func NewReportersFieldsFiltersWithDefaults() *ReportersFieldsFilters`

NewReportersFieldsFiltersWithDefaults instantiates a new ReportersFieldsFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompareFilter

`func (o *ReportersFieldsFilters) GetCompareFilter() []ReportersCompareFilter`

GetCompareFilter returns the CompareFilter field if non-nil, zero value otherwise.

### GetCompareFilterOk

`func (o *ReportersFieldsFilters) GetCompareFilterOk() (*[]ReportersCompareFilter, bool)`

GetCompareFilterOk returns a tuple with the CompareFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareFilter

`func (o *ReportersFieldsFilters) SetCompareFilter(v []ReportersCompareFilter)`

SetCompareFilter sets CompareFilter field to given value.


### SetCompareFilterNil

`func (o *ReportersFieldsFilters) SetCompareFilterNil(b bool)`

 SetCompareFilterNil sets the value for CompareFilter to be an explicit nil

### UnsetCompareFilter
`func (o *ReportersFieldsFilters) UnsetCompareFilter()`

UnsetCompareFilter ensures that no value is present for CompareFilter, not even an explicit nil
### GetContainsFilter

`func (o *ReportersFieldsFilters) GetContainsFilter() ReportersContainsFilter`

GetContainsFilter returns the ContainsFilter field if non-nil, zero value otherwise.

### GetContainsFilterOk

`func (o *ReportersFieldsFilters) GetContainsFilterOk() (*ReportersContainsFilter, bool)`

GetContainsFilterOk returns a tuple with the ContainsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsFilter

`func (o *ReportersFieldsFilters) SetContainsFilter(v ReportersContainsFilter)`

SetContainsFilter sets ContainsFilter field to given value.


### GetContainsInArrayFilter

`func (o *ReportersFieldsFilters) GetContainsInArrayFilter() ReportersContainsFilter`

GetContainsInArrayFilter returns the ContainsInArrayFilter field if non-nil, zero value otherwise.

### GetContainsInArrayFilterOk

`func (o *ReportersFieldsFilters) GetContainsInArrayFilterOk() (*ReportersContainsFilter, bool)`

GetContainsInArrayFilterOk returns a tuple with the ContainsInArrayFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsInArrayFilter

`func (o *ReportersFieldsFilters) SetContainsInArrayFilter(v ReportersContainsFilter)`

SetContainsInArrayFilter sets ContainsInArrayFilter field to given value.

### HasContainsInArrayFilter

`func (o *ReportersFieldsFilters) HasContainsInArrayFilter() bool`

HasContainsInArrayFilter returns a boolean if a field has been set.

### GetMatchFilter

`func (o *ReportersFieldsFilters) GetMatchFilter() ReportersMatchFilter`

GetMatchFilter returns the MatchFilter field if non-nil, zero value otherwise.

### GetMatchFilterOk

`func (o *ReportersFieldsFilters) GetMatchFilterOk() (*ReportersMatchFilter, bool)`

GetMatchFilterOk returns a tuple with the MatchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchFilter

`func (o *ReportersFieldsFilters) SetMatchFilter(v ReportersMatchFilter)`

SetMatchFilter sets MatchFilter field to given value.


### GetMatchInArrayFilter

`func (o *ReportersFieldsFilters) GetMatchInArrayFilter() ReportersMatchFilter`

GetMatchInArrayFilter returns the MatchInArrayFilter field if non-nil, zero value otherwise.

### GetMatchInArrayFilterOk

`func (o *ReportersFieldsFilters) GetMatchInArrayFilterOk() (*ReportersMatchFilter, bool)`

GetMatchInArrayFilterOk returns a tuple with the MatchInArrayFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchInArrayFilter

`func (o *ReportersFieldsFilters) SetMatchInArrayFilter(v ReportersMatchFilter)`

SetMatchInArrayFilter sets MatchInArrayFilter field to given value.

### HasMatchInArrayFilter

`func (o *ReportersFieldsFilters) HasMatchInArrayFilter() bool`

HasMatchInArrayFilter returns a boolean if a field has been set.

### GetNotContainsFilter

`func (o *ReportersFieldsFilters) GetNotContainsFilter() ReportersContainsFilter`

GetNotContainsFilter returns the NotContainsFilter field if non-nil, zero value otherwise.

### GetNotContainsFilterOk

`func (o *ReportersFieldsFilters) GetNotContainsFilterOk() (*ReportersContainsFilter, bool)`

GetNotContainsFilterOk returns a tuple with the NotContainsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotContainsFilter

`func (o *ReportersFieldsFilters) SetNotContainsFilter(v ReportersContainsFilter)`

SetNotContainsFilter sets NotContainsFilter field to given value.

### HasNotContainsFilter

`func (o *ReportersFieldsFilters) HasNotContainsFilter() bool`

HasNotContainsFilter returns a boolean if a field has been set.

### GetOrderFilter

`func (o *ReportersFieldsFilters) GetOrderFilter() ReportersOrderFilter`

GetOrderFilter returns the OrderFilter field if non-nil, zero value otherwise.

### GetOrderFilterOk

`func (o *ReportersFieldsFilters) GetOrderFilterOk() (*ReportersOrderFilter, bool)`

GetOrderFilterOk returns a tuple with the OrderFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderFilter

`func (o *ReportersFieldsFilters) SetOrderFilter(v ReportersOrderFilter)`

SetOrderFilter sets OrderFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


