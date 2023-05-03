# ModelDownloadScanResultsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScanInfo** | Pointer to [**ModelScanResultsCommon**](ModelScanResultsCommon.md) |  | [optional] 
**ScanResults** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewModelDownloadScanResultsResponse

`func NewModelDownloadScanResultsResponse() *ModelDownloadScanResultsResponse`

NewModelDownloadScanResultsResponse instantiates a new ModelDownloadScanResultsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelDownloadScanResultsResponseWithDefaults

`func NewModelDownloadScanResultsResponseWithDefaults() *ModelDownloadScanResultsResponse`

NewModelDownloadScanResultsResponseWithDefaults instantiates a new ModelDownloadScanResultsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScanInfo

`func (o *ModelDownloadScanResultsResponse) GetScanInfo() ModelScanResultsCommon`

GetScanInfo returns the ScanInfo field if non-nil, zero value otherwise.

### GetScanInfoOk

`func (o *ModelDownloadScanResultsResponse) GetScanInfoOk() (*ModelScanResultsCommon, bool)`

GetScanInfoOk returns a tuple with the ScanInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanInfo

`func (o *ModelDownloadScanResultsResponse) SetScanInfo(v ModelScanResultsCommon)`

SetScanInfo sets ScanInfo field to given value.

### HasScanInfo

`func (o *ModelDownloadScanResultsResponse) HasScanInfo() bool`

HasScanInfo returns a boolean if a field has been set.

### GetScanResults

`func (o *ModelDownloadScanResultsResponse) GetScanResults() []interface{}`

GetScanResults returns the ScanResults field if non-nil, zero value otherwise.

### GetScanResultsOk

`func (o *ModelDownloadScanResultsResponse) GetScanResultsOk() (*[]interface{}, bool)`

GetScanResultsOk returns a tuple with the ScanResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanResults

`func (o *ModelDownloadScanResultsResponse) SetScanResults(v []interface{})`

SetScanResults sets ScanResults field to given value.

### HasScanResults

`func (o *ModelDownloadScanResultsResponse) HasScanResults() bool`

HasScanResults returns a boolean if a field has been set.

### SetScanResultsNil

`func (o *ModelDownloadScanResultsResponse) SetScanResultsNil(b bool)`

 SetScanResultsNil sets the value for ScanResults to be an explicit nil

### UnsetScanResults
`func (o *ModelDownloadScanResultsResponse) UnsetScanResults()`

UnsetScanResults ensures that no value is present for ScanResults, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


