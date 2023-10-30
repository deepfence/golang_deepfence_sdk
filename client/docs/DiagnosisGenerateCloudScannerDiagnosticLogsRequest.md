# DiagnosisGenerateCloudScannerDiagnosticLogsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeIds** | [**[]DiagnosisNodeIdentifier**](DiagnosisNodeIdentifier.md) |  | 
**Tail** | **int32** |  | 

## Methods

### NewDiagnosisGenerateCloudScannerDiagnosticLogsRequest

`func NewDiagnosisGenerateCloudScannerDiagnosticLogsRequest(nodeIds []DiagnosisNodeIdentifier, tail int32, ) *DiagnosisGenerateCloudScannerDiagnosticLogsRequest`

NewDiagnosisGenerateCloudScannerDiagnosticLogsRequest instantiates a new DiagnosisGenerateCloudScannerDiagnosticLogsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosisGenerateCloudScannerDiagnosticLogsRequestWithDefaults

`func NewDiagnosisGenerateCloudScannerDiagnosticLogsRequestWithDefaults() *DiagnosisGenerateCloudScannerDiagnosticLogsRequest`

NewDiagnosisGenerateCloudScannerDiagnosticLogsRequestWithDefaults instantiates a new DiagnosisGenerateCloudScannerDiagnosticLogsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeIds

`func (o *DiagnosisGenerateCloudScannerDiagnosticLogsRequest) GetNodeIds() []DiagnosisNodeIdentifier`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *DiagnosisGenerateCloudScannerDiagnosticLogsRequest) GetNodeIdsOk() (*[]DiagnosisNodeIdentifier, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *DiagnosisGenerateCloudScannerDiagnosticLogsRequest) SetNodeIds(v []DiagnosisNodeIdentifier)`

SetNodeIds sets NodeIds field to given value.


### SetNodeIdsNil

`func (o *DiagnosisGenerateCloudScannerDiagnosticLogsRequest) SetNodeIdsNil(b bool)`

 SetNodeIdsNil sets the value for NodeIds to be an explicit nil

### UnsetNodeIds
`func (o *DiagnosisGenerateCloudScannerDiagnosticLogsRequest) UnsetNodeIds()`

UnsetNodeIds ensures that no value is present for NodeIds, not even an explicit nil
### GetTail

`func (o *DiagnosisGenerateCloudScannerDiagnosticLogsRequest) GetTail() int32`

GetTail returns the Tail field if non-nil, zero value otherwise.

### GetTailOk

`func (o *DiagnosisGenerateCloudScannerDiagnosticLogsRequest) GetTailOk() (*int32, bool)`

GetTailOk returns a tuple with the Tail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTail

`func (o *DiagnosisGenerateCloudScannerDiagnosticLogsRequest) SetTail(v int32)`

SetTail sets Tail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


