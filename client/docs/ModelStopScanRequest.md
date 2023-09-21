# ModelStopScanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeIds** | [**[]ModelNodeIdentifier**](ModelNodeIdentifier.md) |  | 
**ScanIds** | **[]string** |  | 
**ScanType** | **string** |  | 

## Methods

### NewModelStopScanRequest

`func NewModelStopScanRequest(nodeIds []ModelNodeIdentifier, scanIds []string, scanType string, ) *ModelStopScanRequest`

NewModelStopScanRequest instantiates a new ModelStopScanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelStopScanRequestWithDefaults

`func NewModelStopScanRequestWithDefaults() *ModelStopScanRequest`

NewModelStopScanRequestWithDefaults instantiates a new ModelStopScanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeIds

`func (o *ModelStopScanRequest) GetNodeIds() []ModelNodeIdentifier`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *ModelStopScanRequest) GetNodeIdsOk() (*[]ModelNodeIdentifier, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *ModelStopScanRequest) SetNodeIds(v []ModelNodeIdentifier)`

SetNodeIds sets NodeIds field to given value.


### SetNodeIdsNil

`func (o *ModelStopScanRequest) SetNodeIdsNil(b bool)`

 SetNodeIdsNil sets the value for NodeIds to be an explicit nil

### UnsetNodeIds
`func (o *ModelStopScanRequest) UnsetNodeIds()`

UnsetNodeIds ensures that no value is present for NodeIds, not even an explicit nil
### GetScanIds

`func (o *ModelStopScanRequest) GetScanIds() []string`

GetScanIds returns the ScanIds field if non-nil, zero value otherwise.

### GetScanIdsOk

`func (o *ModelStopScanRequest) GetScanIdsOk() (*[]string, bool)`

GetScanIdsOk returns a tuple with the ScanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanIds

`func (o *ModelStopScanRequest) SetScanIds(v []string)`

SetScanIds sets ScanIds field to given value.


### SetScanIdsNil

`func (o *ModelStopScanRequest) SetScanIdsNil(b bool)`

 SetScanIdsNil sets the value for ScanIds to be an explicit nil

### UnsetScanIds
`func (o *ModelStopScanRequest) UnsetScanIds()`

UnsetScanIds ensures that no value is present for ScanIds, not even an explicit nil
### GetScanType

`func (o *ModelStopScanRequest) GetScanType() string`

GetScanType returns the ScanType field if non-nil, zero value otherwise.

### GetScanTypeOk

`func (o *ModelStopScanRequest) GetScanTypeOk() (*string, bool)`

GetScanTypeOk returns a tuple with the ScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanType

`func (o *ModelStopScanRequest) SetScanType(v string)`

SetScanType sets ScanType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


