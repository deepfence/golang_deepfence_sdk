# ModelScanResultsActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifyIndividual** | Pointer to **bool** |  | [optional] 
**ResultIds** | **[]string** |  | 
**ScanId** | **string** |  | 
**ScanType** | **string** |  | 

## Methods

### NewModelScanResultsActionRequest

`func NewModelScanResultsActionRequest(resultIds []string, scanId string, scanType string, ) *ModelScanResultsActionRequest`

NewModelScanResultsActionRequest instantiates a new ModelScanResultsActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelScanResultsActionRequestWithDefaults

`func NewModelScanResultsActionRequestWithDefaults() *ModelScanResultsActionRequest`

NewModelScanResultsActionRequestWithDefaults instantiates a new ModelScanResultsActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifyIndividual

`func (o *ModelScanResultsActionRequest) GetNotifyIndividual() bool`

GetNotifyIndividual returns the NotifyIndividual field if non-nil, zero value otherwise.

### GetNotifyIndividualOk

`func (o *ModelScanResultsActionRequest) GetNotifyIndividualOk() (*bool, bool)`

GetNotifyIndividualOk returns a tuple with the NotifyIndividual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyIndividual

`func (o *ModelScanResultsActionRequest) SetNotifyIndividual(v bool)`

SetNotifyIndividual sets NotifyIndividual field to given value.

### HasNotifyIndividual

`func (o *ModelScanResultsActionRequest) HasNotifyIndividual() bool`

HasNotifyIndividual returns a boolean if a field has been set.

### GetResultIds

`func (o *ModelScanResultsActionRequest) GetResultIds() []string`

GetResultIds returns the ResultIds field if non-nil, zero value otherwise.

### GetResultIdsOk

`func (o *ModelScanResultsActionRequest) GetResultIdsOk() (*[]string, bool)`

GetResultIdsOk returns a tuple with the ResultIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultIds

`func (o *ModelScanResultsActionRequest) SetResultIds(v []string)`

SetResultIds sets ResultIds field to given value.


### SetResultIdsNil

`func (o *ModelScanResultsActionRequest) SetResultIdsNil(b bool)`

 SetResultIdsNil sets the value for ResultIds to be an explicit nil

### UnsetResultIds
`func (o *ModelScanResultsActionRequest) UnsetResultIds()`

UnsetResultIds ensures that no value is present for ResultIds, not even an explicit nil
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


