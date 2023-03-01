# ModelNodesInScanResultRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultIds** | **[]string** |  | 
**ScanType** | **string** |  | 

## Methods

### NewModelNodesInScanResultRequest

`func NewModelNodesInScanResultRequest(resultIds []string, scanType string, ) *ModelNodesInScanResultRequest`

NewModelNodesInScanResultRequest instantiates a new ModelNodesInScanResultRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelNodesInScanResultRequestWithDefaults

`func NewModelNodesInScanResultRequestWithDefaults() *ModelNodesInScanResultRequest`

NewModelNodesInScanResultRequestWithDefaults instantiates a new ModelNodesInScanResultRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultIds

`func (o *ModelNodesInScanResultRequest) GetResultIds() []string`

GetResultIds returns the ResultIds field if non-nil, zero value otherwise.

### GetResultIdsOk

`func (o *ModelNodesInScanResultRequest) GetResultIdsOk() (*[]string, bool)`

GetResultIdsOk returns a tuple with the ResultIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultIds

`func (o *ModelNodesInScanResultRequest) SetResultIds(v []string)`

SetResultIds sets ResultIds field to given value.


### SetResultIdsNil

`func (o *ModelNodesInScanResultRequest) SetResultIdsNil(b bool)`

 SetResultIdsNil sets the value for ResultIds to be an explicit nil

### UnsetResultIds
`func (o *ModelNodesInScanResultRequest) UnsetResultIds()`

UnsetResultIds ensures that no value is present for ResultIds, not even an explicit nil
### GetScanType

`func (o *ModelNodesInScanResultRequest) GetScanType() string`

GetScanType returns the ScanType field if non-nil, zero value otherwise.

### GetScanTypeOk

`func (o *ModelNodesInScanResultRequest) GetScanTypeOk() (*string, bool)`

GetScanTypeOk returns a tuple with the ScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanType

`func (o *ModelNodesInScanResultRequest) SetScanType(v string)`

SetScanType sets ScanType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


