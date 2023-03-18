# DiagnosisGenerateAgentDiagnosticLogsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeIds** | [**[]DiagnosisNodeIdentifier**](DiagnosisNodeIdentifier.md) |  | 
**Tail** | **int32** |  | 

## Methods

### NewDiagnosisGenerateAgentDiagnosticLogsRequest

`func NewDiagnosisGenerateAgentDiagnosticLogsRequest(nodeIds []DiagnosisNodeIdentifier, tail int32, ) *DiagnosisGenerateAgentDiagnosticLogsRequest`

NewDiagnosisGenerateAgentDiagnosticLogsRequest instantiates a new DiagnosisGenerateAgentDiagnosticLogsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosisGenerateAgentDiagnosticLogsRequestWithDefaults

`func NewDiagnosisGenerateAgentDiagnosticLogsRequestWithDefaults() *DiagnosisGenerateAgentDiagnosticLogsRequest`

NewDiagnosisGenerateAgentDiagnosticLogsRequestWithDefaults instantiates a new DiagnosisGenerateAgentDiagnosticLogsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeIds

`func (o *DiagnosisGenerateAgentDiagnosticLogsRequest) GetNodeIds() []DiagnosisNodeIdentifier`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *DiagnosisGenerateAgentDiagnosticLogsRequest) GetNodeIdsOk() (*[]DiagnosisNodeIdentifier, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *DiagnosisGenerateAgentDiagnosticLogsRequest) SetNodeIds(v []DiagnosisNodeIdentifier)`

SetNodeIds sets NodeIds field to given value.


### SetNodeIdsNil

`func (o *DiagnosisGenerateAgentDiagnosticLogsRequest) SetNodeIdsNil(b bool)`

 SetNodeIdsNil sets the value for NodeIds to be an explicit nil

### UnsetNodeIds
`func (o *DiagnosisGenerateAgentDiagnosticLogsRequest) UnsetNodeIds()`

UnsetNodeIds ensures that no value is present for NodeIds, not even an explicit nil
### GetTail

`func (o *DiagnosisGenerateAgentDiagnosticLogsRequest) GetTail() int32`

GetTail returns the Tail field if non-nil, zero value otherwise.

### GetTailOk

`func (o *DiagnosisGenerateAgentDiagnosticLogsRequest) GetTailOk() (*int32, bool)`

GetTailOk returns a tuple with the Tail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTail

`func (o *DiagnosisGenerateAgentDiagnosticLogsRequest) SetTail(v int32)`

SetTail sets Tail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


