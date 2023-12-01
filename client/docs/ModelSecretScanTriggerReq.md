# ModelSecretScanTriggerReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**ModelScanFilter**](ModelScanFilter.md) |  | 
**IsPriority** | Pointer to **bool** |  | [optional] 
**NodeIds** | [**[]ModelNodeIdentifier**](ModelNodeIdentifier.md) |  | 

## Methods

### NewModelSecretScanTriggerReq

`func NewModelSecretScanTriggerReq(filters ModelScanFilter, nodeIds []ModelNodeIdentifier, ) *ModelSecretScanTriggerReq`

NewModelSecretScanTriggerReq instantiates a new ModelSecretScanTriggerReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSecretScanTriggerReqWithDefaults

`func NewModelSecretScanTriggerReqWithDefaults() *ModelSecretScanTriggerReq`

NewModelSecretScanTriggerReqWithDefaults instantiates a new ModelSecretScanTriggerReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ModelSecretScanTriggerReq) GetFilters() ModelScanFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ModelSecretScanTriggerReq) GetFiltersOk() (*ModelScanFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ModelSecretScanTriggerReq) SetFilters(v ModelScanFilter)`

SetFilters sets Filters field to given value.


### GetIsPriority

`func (o *ModelSecretScanTriggerReq) GetIsPriority() bool`

GetIsPriority returns the IsPriority field if non-nil, zero value otherwise.

### GetIsPriorityOk

`func (o *ModelSecretScanTriggerReq) GetIsPriorityOk() (*bool, bool)`

GetIsPriorityOk returns a tuple with the IsPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPriority

`func (o *ModelSecretScanTriggerReq) SetIsPriority(v bool)`

SetIsPriority sets IsPriority field to given value.

### HasIsPriority

`func (o *ModelSecretScanTriggerReq) HasIsPriority() bool`

HasIsPriority returns a boolean if a field has been set.

### GetNodeIds

`func (o *ModelSecretScanTriggerReq) GetNodeIds() []ModelNodeIdentifier`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *ModelSecretScanTriggerReq) GetNodeIdsOk() (*[]ModelNodeIdentifier, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *ModelSecretScanTriggerReq) SetNodeIds(v []ModelNodeIdentifier)`

SetNodeIds sets NodeIds field to given value.


### SetNodeIdsNil

`func (o *ModelSecretScanTriggerReq) SetNodeIdsNil(b bool)`

 SetNodeIdsNil sets the value for NodeIds to be an explicit nil

### UnsetNodeIds
`func (o *ModelSecretScanTriggerReq) UnsetNodeIds()`

UnsetNodeIds ensures that no value is present for NodeIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


