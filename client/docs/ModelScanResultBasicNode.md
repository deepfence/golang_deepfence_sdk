# ModelScanResultBasicNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicNodes** | [**[]ModelBasicNode**](ModelBasicNode.md) |  | 
**ResultId** | **string** |  | 

## Methods

### NewModelScanResultBasicNode

`func NewModelScanResultBasicNode(basicNodes []ModelBasicNode, resultId string, ) *ModelScanResultBasicNode`

NewModelScanResultBasicNode instantiates a new ModelScanResultBasicNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelScanResultBasicNodeWithDefaults

`func NewModelScanResultBasicNodeWithDefaults() *ModelScanResultBasicNode`

NewModelScanResultBasicNodeWithDefaults instantiates a new ModelScanResultBasicNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasicNodes

`func (o *ModelScanResultBasicNode) GetBasicNodes() []ModelBasicNode`

GetBasicNodes returns the BasicNodes field if non-nil, zero value otherwise.

### GetBasicNodesOk

`func (o *ModelScanResultBasicNode) GetBasicNodesOk() (*[]ModelBasicNode, bool)`

GetBasicNodesOk returns a tuple with the BasicNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicNodes

`func (o *ModelScanResultBasicNode) SetBasicNodes(v []ModelBasicNode)`

SetBasicNodes sets BasicNodes field to given value.


### SetBasicNodesNil

`func (o *ModelScanResultBasicNode) SetBasicNodesNil(b bool)`

 SetBasicNodesNil sets the value for BasicNodes to be an explicit nil

### UnsetBasicNodes
`func (o *ModelScanResultBasicNode) UnsetBasicNodes()`

UnsetBasicNodes ensures that no value is present for BasicNodes, not even an explicit nil
### GetResultId

`func (o *ModelScanResultBasicNode) GetResultId() string`

GetResultId returns the ResultId field if non-nil, zero value otherwise.

### GetResultIdOk

`func (o *ModelScanResultBasicNode) GetResultIdOk() (*string, bool)`

GetResultIdOk returns a tuple with the ResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultId

`func (o *ModelScanResultBasicNode) SetResultId(v string)`

SetResultId sets ResultId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


