# ModelAddScheduledTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**BenchmarkTypes** | **[]string** |  | 
**CronExpr** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Filters** | [**ModelScanFilter**](ModelScanFilter.md) |  | 
**IsPriority** | Pointer to **bool** |  | [optional] 
**NodeIds** | [**[]ModelNodeIdentifier**](ModelNodeIdentifier.md) |  | 
**ScanConfig** | [**[]ModelVulnerabilityScanConfigLanguage**](ModelVulnerabilityScanConfigLanguage.md) |  | 

## Methods

### NewModelAddScheduledTaskRequest

`func NewModelAddScheduledTaskRequest(action string, benchmarkTypes []string, filters ModelScanFilter, nodeIds []ModelNodeIdentifier, scanConfig []ModelVulnerabilityScanConfigLanguage, ) *ModelAddScheduledTaskRequest`

NewModelAddScheduledTaskRequest instantiates a new ModelAddScheduledTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAddScheduledTaskRequestWithDefaults

`func NewModelAddScheduledTaskRequestWithDefaults() *ModelAddScheduledTaskRequest`

NewModelAddScheduledTaskRequestWithDefaults instantiates a new ModelAddScheduledTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ModelAddScheduledTaskRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ModelAddScheduledTaskRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ModelAddScheduledTaskRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetBenchmarkTypes

`func (o *ModelAddScheduledTaskRequest) GetBenchmarkTypes() []string`

GetBenchmarkTypes returns the BenchmarkTypes field if non-nil, zero value otherwise.

### GetBenchmarkTypesOk

`func (o *ModelAddScheduledTaskRequest) GetBenchmarkTypesOk() (*[]string, bool)`

GetBenchmarkTypesOk returns a tuple with the BenchmarkTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmarkTypes

`func (o *ModelAddScheduledTaskRequest) SetBenchmarkTypes(v []string)`

SetBenchmarkTypes sets BenchmarkTypes field to given value.


### SetBenchmarkTypesNil

`func (o *ModelAddScheduledTaskRequest) SetBenchmarkTypesNil(b bool)`

 SetBenchmarkTypesNil sets the value for BenchmarkTypes to be an explicit nil

### UnsetBenchmarkTypes
`func (o *ModelAddScheduledTaskRequest) UnsetBenchmarkTypes()`

UnsetBenchmarkTypes ensures that no value is present for BenchmarkTypes, not even an explicit nil
### GetCronExpr

`func (o *ModelAddScheduledTaskRequest) GetCronExpr() string`

GetCronExpr returns the CronExpr field if non-nil, zero value otherwise.

### GetCronExprOk

`func (o *ModelAddScheduledTaskRequest) GetCronExprOk() (*string, bool)`

GetCronExprOk returns a tuple with the CronExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpr

`func (o *ModelAddScheduledTaskRequest) SetCronExpr(v string)`

SetCronExpr sets CronExpr field to given value.

### HasCronExpr

`func (o *ModelAddScheduledTaskRequest) HasCronExpr() bool`

HasCronExpr returns a boolean if a field has been set.

### GetDescription

`func (o *ModelAddScheduledTaskRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelAddScheduledTaskRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelAddScheduledTaskRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelAddScheduledTaskRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilters

`func (o *ModelAddScheduledTaskRequest) GetFilters() ModelScanFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ModelAddScheduledTaskRequest) GetFiltersOk() (*ModelScanFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ModelAddScheduledTaskRequest) SetFilters(v ModelScanFilter)`

SetFilters sets Filters field to given value.


### GetIsPriority

`func (o *ModelAddScheduledTaskRequest) GetIsPriority() bool`

GetIsPriority returns the IsPriority field if non-nil, zero value otherwise.

### GetIsPriorityOk

`func (o *ModelAddScheduledTaskRequest) GetIsPriorityOk() (*bool, bool)`

GetIsPriorityOk returns a tuple with the IsPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPriority

`func (o *ModelAddScheduledTaskRequest) SetIsPriority(v bool)`

SetIsPriority sets IsPriority field to given value.

### HasIsPriority

`func (o *ModelAddScheduledTaskRequest) HasIsPriority() bool`

HasIsPriority returns a boolean if a field has been set.

### GetNodeIds

`func (o *ModelAddScheduledTaskRequest) GetNodeIds() []ModelNodeIdentifier`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *ModelAddScheduledTaskRequest) GetNodeIdsOk() (*[]ModelNodeIdentifier, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *ModelAddScheduledTaskRequest) SetNodeIds(v []ModelNodeIdentifier)`

SetNodeIds sets NodeIds field to given value.


### SetNodeIdsNil

`func (o *ModelAddScheduledTaskRequest) SetNodeIdsNil(b bool)`

 SetNodeIdsNil sets the value for NodeIds to be an explicit nil

### UnsetNodeIds
`func (o *ModelAddScheduledTaskRequest) UnsetNodeIds()`

UnsetNodeIds ensures that no value is present for NodeIds, not even an explicit nil
### GetScanConfig

`func (o *ModelAddScheduledTaskRequest) GetScanConfig() []ModelVulnerabilityScanConfigLanguage`

GetScanConfig returns the ScanConfig field if non-nil, zero value otherwise.

### GetScanConfigOk

`func (o *ModelAddScheduledTaskRequest) GetScanConfigOk() (*[]ModelVulnerabilityScanConfigLanguage, bool)`

GetScanConfigOk returns a tuple with the ScanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanConfig

`func (o *ModelAddScheduledTaskRequest) SetScanConfig(v []ModelVulnerabilityScanConfigLanguage)`

SetScanConfig sets ScanConfig field to given value.


### SetScanConfigNil

`func (o *ModelAddScheduledTaskRequest) SetScanConfigNil(b bool)`

 SetScanConfigNil sets the value for ScanConfig to be an explicit nil

### UnsetScanConfig
`func (o *ModelAddScheduledTaskRequest) UnsetScanConfig()`

UnsetScanConfig ensures that no value is present for ScanConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


