# ModelComplianceScanTriggerReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BenchmarkTypes** | **[]string** |  | 
**Filters** | [**ModelScanFilter**](ModelScanFilter.md) |  | 
**IsPriority** | Pointer to **bool** |  | [optional] 
**NodeIds** | [**[]ModelNodeIdentifier**](ModelNodeIdentifier.md) |  | 

## Methods

### NewModelComplianceScanTriggerReq

`func NewModelComplianceScanTriggerReq(benchmarkTypes []string, filters ModelScanFilter, nodeIds []ModelNodeIdentifier, ) *ModelComplianceScanTriggerReq`

NewModelComplianceScanTriggerReq instantiates a new ModelComplianceScanTriggerReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelComplianceScanTriggerReqWithDefaults

`func NewModelComplianceScanTriggerReqWithDefaults() *ModelComplianceScanTriggerReq`

NewModelComplianceScanTriggerReqWithDefaults instantiates a new ModelComplianceScanTriggerReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBenchmarkTypes

`func (o *ModelComplianceScanTriggerReq) GetBenchmarkTypes() []string`

GetBenchmarkTypes returns the BenchmarkTypes field if non-nil, zero value otherwise.

### GetBenchmarkTypesOk

`func (o *ModelComplianceScanTriggerReq) GetBenchmarkTypesOk() (*[]string, bool)`

GetBenchmarkTypesOk returns a tuple with the BenchmarkTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmarkTypes

`func (o *ModelComplianceScanTriggerReq) SetBenchmarkTypes(v []string)`

SetBenchmarkTypes sets BenchmarkTypes field to given value.


### SetBenchmarkTypesNil

`func (o *ModelComplianceScanTriggerReq) SetBenchmarkTypesNil(b bool)`

 SetBenchmarkTypesNil sets the value for BenchmarkTypes to be an explicit nil

### UnsetBenchmarkTypes
`func (o *ModelComplianceScanTriggerReq) UnsetBenchmarkTypes()`

UnsetBenchmarkTypes ensures that no value is present for BenchmarkTypes, not even an explicit nil
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


### GetIsPriority

`func (o *ModelComplianceScanTriggerReq) GetIsPriority() bool`

GetIsPriority returns the IsPriority field if non-nil, zero value otherwise.

### GetIsPriorityOk

`func (o *ModelComplianceScanTriggerReq) GetIsPriorityOk() (*bool, bool)`

GetIsPriorityOk returns a tuple with the IsPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPriority

`func (o *ModelComplianceScanTriggerReq) SetIsPriority(v bool)`

SetIsPriority sets IsPriority field to given value.

### HasIsPriority

`func (o *ModelComplianceScanTriggerReq) HasIsPriority() bool`

HasIsPriority returns a boolean if a field has been set.

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


