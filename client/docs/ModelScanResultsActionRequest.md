# ModelScanResultsActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeIds** | **[]string** |  | 
**ScanId** | **string** |  | 
**ScanType** | **string** |  | 

## Methods

### NewModelScanResultsActionRequest

`func NewModelScanResultsActionRequest(nodeIds []string, scanId string, scanType string, ) *ModelScanResultsActionRequest`

NewModelScanResultsActionRequest instantiates a new ModelScanResultsActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelScanResultsActionRequestWithDefaults

`func NewModelScanResultsActionRequestWithDefaults() *ModelScanResultsActionRequest`

NewModelScanResultsActionRequestWithDefaults instantiates a new ModelScanResultsActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeIds

`func (o *ModelScanResultsActionRequest) GetNodeIds() []string`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *ModelScanResultsActionRequest) GetNodeIdsOk() (*[]string, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *ModelScanResultsActionRequest) SetNodeIds(v []string)`

SetNodeIds sets NodeIds field to given value.


### SetNodeIdsNil

`func (o *ModelScanResultsActionRequest) SetNodeIdsNil(b bool)`

 SetNodeIdsNil sets the value for NodeIds to be an explicit nil

### UnsetNodeIds
`func (o *ModelScanResultsActionRequest) UnsetNodeIds()`

UnsetNodeIds ensures that no value is present for NodeIds, not even an explicit nil
### GetScanId

`func (o *ModelScanResultsActionRequest) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *ModelScanResultsActionRequest) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *ModelScanResultsActionRequest) SetScanId(v string)`

SetScanId sets ScanId field to given value.


### GetScanType

`func (o *ModelScanResultsActionRequest) GetScanType() string`

GetScanType returns the ScanType field if non-nil, zero value otherwise.

### GetScanTypeOk

`func (o *ModelScanResultsActionRequest) GetScanTypeOk() (*string, bool)`

GetScanTypeOk returns a tuple with the ScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanType

`func (o *ModelScanResultsActionRequest) SetScanType(v string)`

SetScanType sets ScanType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


