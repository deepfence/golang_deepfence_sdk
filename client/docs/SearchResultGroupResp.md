# SearchResultGroupResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]SearchResultGroup**](SearchResultGroup.md) |  | [optional] 

## Methods

### NewSearchResultGroupResp

`func NewSearchResultGroupResp() *SearchResultGroupResp`

NewSearchResultGroupResp instantiates a new SearchResultGroupResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultGroupRespWithDefaults

`func NewSearchResultGroupRespWithDefaults() *SearchResultGroupResp`

NewSearchResultGroupRespWithDefaults instantiates a new SearchResultGroupResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *SearchResultGroupResp) GetGroups() []SearchResultGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *SearchResultGroupResp) GetGroupsOk() (*[]SearchResultGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *SearchResultGroupResp) SetGroups(v []SearchResultGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *SearchResultGroupResp) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *SearchResultGroupResp) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *SearchResultGroupResp) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


