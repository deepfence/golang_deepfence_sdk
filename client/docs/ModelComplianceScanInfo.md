# ModelComplianceScanInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BenchmarkType** | **string** |  | 
**NodeId** | **string** |  | 
**NodeType** | **string** |  | 
**ScanId** | **string** |  | 
**SeverityCounts** | **map[string]int32** |  | 
**Status** | **string** |  | 
**UpdatedAt** | **int64** |  | 

## Methods

### NewModelComplianceScanInfo

`func NewModelComplianceScanInfo(benchmarkType string, nodeId string, nodeType string, scanId string, severityCounts map[string]int32, status string, updatedAt int64, ) *ModelComplianceScanInfo`

NewModelComplianceScanInfo instantiates a new ModelComplianceScanInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelComplianceScanInfoWithDefaults

`func NewModelComplianceScanInfoWithDefaults() *ModelComplianceScanInfo`

NewModelComplianceScanInfoWithDefaults instantiates a new ModelComplianceScanInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBenchmarkType

`func (o *ModelComplianceScanInfo) GetBenchmarkType() string`

GetBenchmarkType returns the BenchmarkType field if non-nil, zero value otherwise.

### GetBenchmarkTypeOk

`func (o *ModelComplianceScanInfo) GetBenchmarkTypeOk() (*string, bool)`

GetBenchmarkTypeOk returns a tuple with the BenchmarkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmarkType

`func (o *ModelComplianceScanInfo) SetBenchmarkType(v string)`

SetBenchmarkType sets BenchmarkType field to given value.


### GetNodeId

`func (o *ModelComplianceScanInfo) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelComplianceScanInfo) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelComplianceScanInfo) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNodeType

`func (o *ModelComplianceScanInfo) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *ModelComplianceScanInfo) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *ModelComplianceScanInfo) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetScanId

`func (o *ModelComplianceScanInfo) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *ModelComplianceScanInfo) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *ModelComplianceScanInfo) SetScanId(v string)`

SetScanId sets ScanId field to given value.


### GetSeverityCounts

`func (o *ModelComplianceScanInfo) GetSeverityCounts() map[string]int32`

GetSeverityCounts returns the SeverityCounts field if non-nil, zero value otherwise.

### GetSeverityCountsOk

`func (o *ModelComplianceScanInfo) GetSeverityCountsOk() (*map[string]int32, bool)`

GetSeverityCountsOk returns a tuple with the SeverityCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityCounts

`func (o *ModelComplianceScanInfo) SetSeverityCounts(v map[string]int32)`

SetSeverityCounts sets SeverityCounts field to given value.


### SetSeverityCountsNil

`func (o *ModelComplianceScanInfo) SetSeverityCountsNil(b bool)`

 SetSeverityCountsNil sets the value for SeverityCounts to be an explicit nil

### UnsetSeverityCounts
`func (o *ModelComplianceScanInfo) UnsetSeverityCounts()`

UnsetSeverityCounts ensures that no value is present for SeverityCounts, not even an explicit nil
### GetStatus

`func (o *ModelComplianceScanInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelComplianceScanInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelComplianceScanInfo) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *ModelComplianceScanInfo) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelComplianceScanInfo) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelComplianceScanInfo) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


