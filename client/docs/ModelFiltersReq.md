# ModelFiltersReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | **[]string** |  | 
**Having** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewModelFiltersReq

`func NewModelFiltersReq(filters []string, ) *ModelFiltersReq`

NewModelFiltersReq instantiates a new ModelFiltersReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelFiltersReqWithDefaults

`func NewModelFiltersReqWithDefaults() *ModelFiltersReq`

NewModelFiltersReqWithDefaults instantiates a new ModelFiltersReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ModelFiltersReq) GetFilters() []string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ModelFiltersReq) GetFiltersOk() (*[]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ModelFiltersReq) SetFilters(v []string)`

SetFilters sets Filters field to given value.


### SetFiltersNil

`func (o *ModelFiltersReq) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *ModelFiltersReq) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetHaving

`func (o *ModelFiltersReq) GetHaving() map[string]interface{}`

GetHaving returns the Having field if non-nil, zero value otherwise.

### GetHavingOk

`func (o *ModelFiltersReq) GetHavingOk() (*map[string]interface{}, bool)`

GetHavingOk returns a tuple with the Having field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaving

`func (o *ModelFiltersReq) SetHaving(v map[string]interface{})`

SetHaving sets Having field to given value.

### HasHaving

`func (o *ModelFiltersReq) HasHaving() bool`

HasHaving returns a boolean if a field has been set.

### SetHavingNil

`func (o *ModelFiltersReq) SetHavingNil(b bool)`

 SetHavingNil sets the value for Having to be an explicit nil

### UnsetHaving
`func (o *ModelFiltersReq) UnsetHaving()`

UnsetHaving ensures that no value is present for Having, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


