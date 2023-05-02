# ModelScanListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldsFilter** | [**ReportersFieldsFilters**](ReportersFieldsFilters.md) |  | 
**NodeIds** | [**[]ModelNodeIdentifier**](ModelNodeIdentifier.md) |  | 
**Window** | [**ModelFetchWindow**](ModelFetchWindow.md) |  | 

## Methods

### NewModelScanListReq

`func NewModelScanListReq(fieldsFilter ReportersFieldsFilters, nodeIds []ModelNodeIdentifier, window ModelFetchWindow, ) *ModelScanListReq`

NewModelScanListReq instantiates a new ModelScanListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelScanListReqWithDefaults

`func NewModelScanListReqWithDefaults() *ModelScanListReq`

NewModelScanListReqWithDefaults instantiates a new ModelScanListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldsFilter

`func (o *ModelScanListReq) GetFieldsFilter() ReportersFieldsFilters`

GetFieldsFilter returns the FieldsFilter field if non-nil, zero value otherwise.

### GetFieldsFilterOk

`func (o *ModelScanListReq) GetFieldsFilterOk() (*ReportersFieldsFilters, bool)`

GetFieldsFilterOk returns a tuple with the FieldsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsFilter

`func (o *ModelScanListReq) SetFieldsFilter(v ReportersFieldsFilters)`

SetFieldsFilter sets FieldsFilter field to given value.


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


