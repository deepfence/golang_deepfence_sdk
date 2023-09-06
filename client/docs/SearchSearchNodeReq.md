# SearchSearchNodeReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtendedNodeFilter** | Pointer to [**SearchSearchFilter**](SearchSearchFilter.md) |  | [optional] 
**NodeFilter** | [**SearchSearchFilter**](SearchSearchFilter.md) |  | 
**RelatedNodeFilter** | Pointer to [**SearchChainedSearchFilter**](SearchChainedSearchFilter.md) |  | [optional] 
**Window** | [**ModelFetchWindow**](ModelFetchWindow.md) |  | 

## Methods

### NewSearchSearchNodeReq

`func NewSearchSearchNodeReq(nodeFilter SearchSearchFilter, window ModelFetchWindow, ) *SearchSearchNodeReq`

NewSearchSearchNodeReq instantiates a new SearchSearchNodeReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSearchNodeReqWithDefaults

`func NewSearchSearchNodeReqWithDefaults() *SearchSearchNodeReq`

NewSearchSearchNodeReqWithDefaults instantiates a new SearchSearchNodeReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtendedNodeFilter

`func (o *SearchSearchNodeReq) GetExtendedNodeFilter() SearchSearchFilter`

GetExtendedNodeFilter returns the ExtendedNodeFilter field if non-nil, zero value otherwise.

### GetExtendedNodeFilterOk

`func (o *SearchSearchNodeReq) GetExtendedNodeFilterOk() (*SearchSearchFilter, bool)`

GetExtendedNodeFilterOk returns a tuple with the ExtendedNodeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedNodeFilter

`func (o *SearchSearchNodeReq) SetExtendedNodeFilter(v SearchSearchFilter)`

SetExtendedNodeFilter sets ExtendedNodeFilter field to given value.

### HasExtendedNodeFilter

`func (o *SearchSearchNodeReq) HasExtendedNodeFilter() bool`

HasExtendedNodeFilter returns a boolean if a field has been set.

### GetNodeFilter

`func (o *SearchSearchNodeReq) GetNodeFilter() SearchSearchFilter`

GetNodeFilter returns the NodeFilter field if non-nil, zero value otherwise.

### GetNodeFilterOk

`func (o *SearchSearchNodeReq) GetNodeFilterOk() (*SearchSearchFilter, bool)`

GetNodeFilterOk returns a tuple with the NodeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeFilter

`func (o *SearchSearchNodeReq) SetNodeFilter(v SearchSearchFilter)`

SetNodeFilter sets NodeFilter field to given value.


### GetRelatedNodeFilter

`func (o *SearchSearchNodeReq) GetRelatedNodeFilter() SearchChainedSearchFilter`

GetRelatedNodeFilter returns the RelatedNodeFilter field if non-nil, zero value otherwise.

### GetRelatedNodeFilterOk

`func (o *SearchSearchNodeReq) GetRelatedNodeFilterOk() (*SearchChainedSearchFilter, bool)`

GetRelatedNodeFilterOk returns a tuple with the RelatedNodeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedNodeFilter

`func (o *SearchSearchNodeReq) SetRelatedNodeFilter(v SearchChainedSearchFilter)`

SetRelatedNodeFilter sets RelatedNodeFilter field to given value.

### HasRelatedNodeFilter

`func (o *SearchSearchNodeReq) HasRelatedNodeFilter() bool`

HasRelatedNodeFilter returns a boolean if a field has been set.

### GetWindow

`func (o *SearchSearchNodeReq) GetWindow() ModelFetchWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *SearchSearchNodeReq) GetWindowOk() (*ModelFetchWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *SearchSearchNodeReq) SetWindow(v ModelFetchWindow)`

SetWindow sets Window field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


