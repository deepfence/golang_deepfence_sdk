# SearchChainedSearchFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextFilter** | Pointer to **interface{}** |  | [optional] 
**NodeFilter** | [**SearchSearchFilter**](SearchSearchFilter.md) |  | 
**RelationShip** | **string** |  | 

## Methods

### NewSearchChainedSearchFilter

`func NewSearchChainedSearchFilter(nodeFilter SearchSearchFilter, relationShip string, ) *SearchChainedSearchFilter`

NewSearchChainedSearchFilter instantiates a new SearchChainedSearchFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchChainedSearchFilterWithDefaults

`func NewSearchChainedSearchFilterWithDefaults() *SearchChainedSearchFilter`

NewSearchChainedSearchFilterWithDefaults instantiates a new SearchChainedSearchFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextFilter

`func (o *SearchChainedSearchFilter) GetNextFilter() interface{}`

GetNextFilter returns the NextFilter field if non-nil, zero value otherwise.

### GetNextFilterOk

`func (o *SearchChainedSearchFilter) GetNextFilterOk() (*interface{}, bool)`

GetNextFilterOk returns a tuple with the NextFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFilter

`func (o *SearchChainedSearchFilter) SetNextFilter(v interface{})`

SetNextFilter sets NextFilter field to given value.

### HasNextFilter

`func (o *SearchChainedSearchFilter) HasNextFilter() bool`

HasNextFilter returns a boolean if a field has been set.

### SetNextFilterNil

`func (o *SearchChainedSearchFilter) SetNextFilterNil(b bool)`

 SetNextFilterNil sets the value for NextFilter to be an explicit nil

### UnsetNextFilter
`func (o *SearchChainedSearchFilter) UnsetNextFilter()`

UnsetNextFilter ensures that no value is present for NextFilter, not even an explicit nil
### GetNodeFilter

`func (o *SearchChainedSearchFilter) GetNodeFilter() SearchSearchFilter`

GetNodeFilter returns the NodeFilter field if non-nil, zero value otherwise.

### GetNodeFilterOk

`func (o *SearchChainedSearchFilter) GetNodeFilterOk() (*SearchSearchFilter, bool)`

GetNodeFilterOk returns a tuple with the NodeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeFilter

`func (o *SearchChainedSearchFilter) SetNodeFilter(v SearchSearchFilter)`

SetNodeFilter sets NodeFilter field to given value.


### GetRelationShip

`func (o *SearchChainedSearchFilter) GetRelationShip() string`

GetRelationShip returns the RelationShip field if non-nil, zero value otherwise.

### GetRelationShipOk

`func (o *SearchChainedSearchFilter) GetRelationShipOk() (*string, bool)`

GetRelationShipOk returns a tuple with the RelationShip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationShip

`func (o *SearchChainedSearchFilter) SetRelationShip(v string)`

SetRelationShip sets RelationShip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


