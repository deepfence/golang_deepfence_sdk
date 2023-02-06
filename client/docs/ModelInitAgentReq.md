# ModelInitAgentReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableWorkload** | **int32** |  | 
**NodeId** | **string** |  | 
**Version** | **string** |  | 

## Methods

### NewModelInitAgentReq

`func NewModelInitAgentReq(availableWorkload int32, nodeId string, version string, ) *ModelInitAgentReq`

NewModelInitAgentReq instantiates a new ModelInitAgentReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelInitAgentReqWithDefaults

`func NewModelInitAgentReqWithDefaults() *ModelInitAgentReq`

NewModelInitAgentReqWithDefaults instantiates a new ModelInitAgentReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableWorkload

`func (o *ModelInitAgentReq) GetAvailableWorkload() int32`

GetAvailableWorkload returns the AvailableWorkload field if non-nil, zero value otherwise.

### GetAvailableWorkloadOk

`func (o *ModelInitAgentReq) GetAvailableWorkloadOk() (*int32, bool)`

GetAvailableWorkloadOk returns a tuple with the AvailableWorkload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableWorkload

`func (o *ModelInitAgentReq) SetAvailableWorkload(v int32)`

SetAvailableWorkload sets AvailableWorkload field to given value.


### GetNodeId

`func (o *ModelInitAgentReq) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelInitAgentReq) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelInitAgentReq) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetVersion

`func (o *ModelInitAgentReq) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelInitAgentReq) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelInitAgentReq) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


