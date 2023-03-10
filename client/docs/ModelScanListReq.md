# ModelScanListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeIds** | [**[]ModelNodeIdentifier**](ModelNodeIdentifier.md) |  | 
**ScanStatus** | Pointer to **[]string** |  | [optional] 
**Window** | [**ModelFetchWindow**](ModelFetchWindow.md) |  | 

## Methods

### NewModelScanListReq

`func NewModelScanListReq(nodeIds []ModelNodeIdentifier, window ModelFetchWindow, ) *ModelScanListReq`

NewModelScanListReq instantiates a new ModelScanListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelScanListReqWithDefaults

`func NewModelScanListReqWithDefaults() *ModelScanListReq`

NewModelScanListReqWithDefaults instantiates a new ModelScanListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeIds

`func (o *ModelScanListReq) GetNodeIds() []ModelNodeIdentifier`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *ModelScanListReq) GetNodeIdsOk() (*[]ModelNodeIdentifier, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *ModelScanListReq) SetNodeIds(v []ModelNodeIdentifier)`

SetNodeIds sets NodeIds field to given value.


### SetNodeIdsNil

`func (o *ModelScanListReq) SetNodeIdsNil(b bool)`

 SetNodeIdsNil sets the value for NodeIds to be an explicit nil

### UnsetNodeIds
`func (o *ModelScanListReq) UnsetNodeIds()`

UnsetNodeIds ensures that no value is present for NodeIds, not even an explicit nil
### GetScanStatus

`func (o *ModelScanListReq) GetScanStatus() []string`

GetScanStatus returns the ScanStatus field if non-nil, zero value otherwise.

### GetScanStatusOk

`func (o *ModelScanListReq) GetScanStatusOk() (*[]string, bool)`

GetScanStatusOk returns a tuple with the ScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStatus

`func (o *ModelScanListReq) SetScanStatus(v []string)`

SetScanStatus sets ScanStatus field to given value.

### HasScanStatus

`func (o *ModelScanListReq) HasScanStatus() bool`

HasScanStatus returns a boolean if a field has been set.

### SetScanStatusNil

`func (o *ModelScanListReq) SetScanStatusNil(b bool)`

 SetScanStatusNil sets the value for ScanStatus to be an explicit nil

### UnsetScanStatus
`func (o *ModelScanListReq) UnsetScanStatus()`

UnsetScanStatus ensures that no value is present for ScanStatus, not even an explicit nil
### GetWindow

`func (o *ModelScanListReq) GetWindow() ModelFetchWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *ModelScanListReq) GetWindowOk() (*ModelFetchWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *ModelScanListReq) SetWindow(v ModelFetchWindow)`

SetWindow sets Window field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


