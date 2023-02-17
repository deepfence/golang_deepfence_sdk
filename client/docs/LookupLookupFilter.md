# LookupLookupFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InFieldFilter** | **[]string** |  | 
**NodeIds** | **[]string** |  | 
**Window** | [**ModelFetchWindow**](ModelFetchWindow.md) |  | 

## Methods

### NewLookupLookupFilter

`func NewLookupLookupFilter(inFieldFilter []string, nodeIds []string, window ModelFetchWindow, ) *LookupLookupFilter`

NewLookupLookupFilter instantiates a new LookupLookupFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupLookupFilterWithDefaults

`func NewLookupLookupFilterWithDefaults() *LookupLookupFilter`

NewLookupLookupFilterWithDefaults instantiates a new LookupLookupFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInFieldFilter

`func (o *LookupLookupFilter) GetInFieldFilter() []string`

GetInFieldFilter returns the InFieldFilter field if non-nil, zero value otherwise.

### GetInFieldFilterOk

`func (o *LookupLookupFilter) GetInFieldFilterOk() (*[]string, bool)`

GetInFieldFilterOk returns a tuple with the InFieldFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInFieldFilter

`func (o *LookupLookupFilter) SetInFieldFilter(v []string)`

SetInFieldFilter sets InFieldFilter field to given value.


### SetInFieldFilterNil

`func (o *LookupLookupFilter) SetInFieldFilterNil(b bool)`

 SetInFieldFilterNil sets the value for InFieldFilter to be an explicit nil

### UnsetInFieldFilter
`func (o *LookupLookupFilter) UnsetInFieldFilter()`

UnsetInFieldFilter ensures that no value is present for InFieldFilter, not even an explicit nil
### GetNodeIds

`func (o *LookupLookupFilter) GetNodeIds() []string`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *LookupLookupFilter) GetNodeIdsOk() (*[]string, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *LookupLookupFilter) SetNodeIds(v []string)`

SetNodeIds sets NodeIds field to given value.


### SetNodeIdsNil

`func (o *LookupLookupFilter) SetNodeIdsNil(b bool)`

 SetNodeIdsNil sets the value for NodeIds to be an explicit nil

### UnsetNodeIds
`func (o *LookupLookupFilter) UnsetNodeIds()`

UnsetNodeIds ensures that no value is present for NodeIds, not even an explicit nil
### GetWindow

`func (o *LookupLookupFilter) GetWindow() ModelFetchWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *LookupLookupFilter) GetWindowOk() (*ModelFetchWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *LookupLookupFilter) SetWindow(v ModelFetchWindow)`

SetWindow sets Window field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


