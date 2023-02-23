# ModelScanResultsActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocIds** | **[]string** |  | 
**ScanId** | **string** |  | 
**ScanType** | **string** |  | 

## Methods

### NewModelScanResultsActionRequest

`func NewModelScanResultsActionRequest(docIds []string, scanId string, scanType string, ) *ModelScanResultsActionRequest`

NewModelScanResultsActionRequest instantiates a new ModelScanResultsActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelScanResultsActionRequestWithDefaults

`func NewModelScanResultsActionRequestWithDefaults() *ModelScanResultsActionRequest`

NewModelScanResultsActionRequestWithDefaults instantiates a new ModelScanResultsActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocIds

`func (o *ModelScanResultsActionRequest) GetDocIds() []string`

GetDocIds returns the DocIds field if non-nil, zero value otherwise.

### GetDocIdsOk

`func (o *ModelScanResultsActionRequest) GetDocIdsOk() (*[]string, bool)`

GetDocIdsOk returns a tuple with the DocIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocIds

`func (o *ModelScanResultsActionRequest) SetDocIds(v []string)`

SetDocIds sets DocIds field to given value.


### SetDocIdsNil

`func (o *ModelScanResultsActionRequest) SetDocIdsNil(b bool)`

 SetDocIdsNil sets the value for DocIds to be an explicit nil

### UnsetDocIds
`func (o *ModelScanResultsActionRequest) UnsetDocIds()`

UnsetDocIds ensures that no value is present for DocIds, not even an explicit nil
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


