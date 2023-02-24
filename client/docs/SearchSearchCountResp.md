# SearchSearchCountResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | **map[string]int32** |  | 
**Count** | **int32** |  | 

## Methods

### NewSearchSearchCountResp

`func NewSearchSearchCountResp(categories map[string]int32, count int32, ) *SearchSearchCountResp`

NewSearchSearchCountResp instantiates a new SearchSearchCountResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSearchCountRespWithDefaults

`func NewSearchSearchCountRespWithDefaults() *SearchSearchCountResp`

NewSearchSearchCountRespWithDefaults instantiates a new SearchSearchCountResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *SearchSearchCountResp) GetCategories() map[string]int32`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *SearchSearchCountResp) GetCategoriesOk() (*map[string]int32, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *SearchSearchCountResp) SetCategories(v map[string]int32)`

SetCategories sets Categories field to given value.


### SetCategoriesNil

`func (o *SearchSearchCountResp) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *SearchSearchCountResp) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil
### GetCount

`func (o *SearchSearchCountResp) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SearchSearchCountResp) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SearchSearchCountResp) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


