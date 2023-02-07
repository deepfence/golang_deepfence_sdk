# ModelComplianceScanTriggerReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**ModelScanFilter**](ModelScanFilter.md) |  | 
**NodeIds** | [**[]ModelNodeIdentifier**](ModelNodeIdentifier.md) |  | 

## Methods

### NewModelComplianceScanTriggerReq

`func NewModelComplianceScanTriggerReq(filters ModelScanFilter, nodeIds []ModelNodeIdentifier, ) *ModelComplianceScanTriggerReq`

NewModelComplianceScanTriggerReq instantiates a new ModelComplianceScanTriggerReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelComplianceScanTriggerReqWithDefaults

`func NewModelComplianceScanTriggerReqWithDefaults() *ModelComplianceScanTriggerReq`

NewModelComplianceScanTriggerReqWithDefaults instantiates a new ModelComplianceScanTriggerReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ModelComplianceScanTriggerReq) GetFilters() ModelScanFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ModelComplianceScanTriggerReq) GetFiltersOk() (*ModelScanFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ModelComplianceScanTriggerReq) SetFilters(v ModelScanFilter)`

SetFilters sets Filters field to given value.


### GetNodeIds

`func (o *ModelComplianceScanTriggerReq) GetNodeIds() []ModelNodeIdentifier`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *ModelComplianceScanTriggerReq) GetNodeIdsOk() (*[]ModelNodeIdentifier, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *ModelComplianceScanTriggerReq) SetNodeIds(v []ModelNodeIdentifier)`

SetNodeIds sets NodeIds field to given value.


### SetNodeIdsNil

`func (o *ModelComplianceScanTriggerReq) SetNodeIdsNil(b bool)`

 SetNodeIdsNil sets the value for NodeIds to be an explicit nil

### UnsetNodeIds
`func (o *ModelComplianceScanTriggerReq) UnsetNodeIds()`

UnsetNodeIds ensures that no value is present for NodeIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


