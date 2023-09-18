# ModelCloudComplianceScanDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**Benchmarks** | Pointer to [**[]ModelCloudComplianceBenchmark**](ModelCloudComplianceBenchmark.md) |  | [optional] 
**ScanId** | Pointer to **string** |  | [optional] 
**ScanTypes** | Pointer to **[]string** |  | [optional] 
**StopRequested** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelCloudComplianceScanDetails

`func NewModelCloudComplianceScanDetails() *ModelCloudComplianceScanDetails`

NewModelCloudComplianceScanDetails instantiates a new ModelCloudComplianceScanDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCloudComplianceScanDetailsWithDefaults

`func NewModelCloudComplianceScanDetailsWithDefaults() *ModelCloudComplianceScanDetails`

NewModelCloudComplianceScanDetailsWithDefaults instantiates a new ModelCloudComplianceScanDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ModelCloudComplianceScanDetails) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ModelCloudComplianceScanDetails) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ModelCloudComplianceScanDetails) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ModelCloudComplianceScanDetails) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetBenchmarks

`func (o *ModelCloudComplianceScanDetails) GetBenchmarks() []ModelCloudComplianceBenchmark`

GetBenchmarks returns the Benchmarks field if non-nil, zero value otherwise.

### GetBenchmarksOk

`func (o *ModelCloudComplianceScanDetails) GetBenchmarksOk() (*[]ModelCloudComplianceBenchmark, bool)`

GetBenchmarksOk returns a tuple with the Benchmarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmarks

`func (o *ModelCloudComplianceScanDetails) SetBenchmarks(v []ModelCloudComplianceBenchmark)`

SetBenchmarks sets Benchmarks field to given value.

### HasBenchmarks

`func (o *ModelCloudComplianceScanDetails) HasBenchmarks() bool`

HasBenchmarks returns a boolean if a field has been set.

### SetBenchmarksNil

`func (o *ModelCloudComplianceScanDetails) SetBenchmarksNil(b bool)`

 SetBenchmarksNil sets the value for Benchmarks to be an explicit nil

### UnsetBenchmarks
`func (o *ModelCloudComplianceScanDetails) UnsetBenchmarks()`

UnsetBenchmarks ensures that no value is present for Benchmarks, not even an explicit nil
### GetScanId

`func (o *ModelCloudComplianceScanDetails) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *ModelCloudComplianceScanDetails) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *ModelCloudComplianceScanDetails) SetScanId(v string)`

SetScanId sets ScanId field to given value.

### HasScanId

`func (o *ModelCloudComplianceScanDetails) HasScanId() bool`

HasScanId returns a boolean if a field has been set.

### GetScanTypes

`func (o *ModelCloudComplianceScanDetails) GetScanTypes() []string`

GetScanTypes returns the ScanTypes field if non-nil, zero value otherwise.

### GetScanTypesOk

`func (o *ModelCloudComplianceScanDetails) GetScanTypesOk() (*[]string, bool)`

GetScanTypesOk returns a tuple with the ScanTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanTypes

`func (o *ModelCloudComplianceScanDetails) SetScanTypes(v []string)`

SetScanTypes sets ScanTypes field to given value.

### HasScanTypes

`func (o *ModelCloudComplianceScanDetails) HasScanTypes() bool`

HasScanTypes returns a boolean if a field has been set.

### SetScanTypesNil

`func (o *ModelCloudComplianceScanDetails) SetScanTypesNil(b bool)`

 SetScanTypesNil sets the value for ScanTypes to be an explicit nil

### UnsetScanTypes
`func (o *ModelCloudComplianceScanDetails) UnsetScanTypes()`

UnsetScanTypes ensures that no value is present for ScanTypes, not even an explicit nil
### GetStopRequested

`func (o *ModelCloudComplianceScanDetails) GetStopRequested() bool`

GetStopRequested returns the StopRequested field if non-nil, zero value otherwise.

### GetStopRequestedOk

`func (o *ModelCloudComplianceScanDetails) GetStopRequestedOk() (*bool, bool)`

GetStopRequestedOk returns a tuple with the StopRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopRequested

`func (o *ModelCloudComplianceScanDetails) SetStopRequested(v bool)`

SetStopRequested sets StopRequested field to given value.

### HasStopRequested

`func (o *ModelCloudComplianceScanDetails) HasStopRequested() bool`

HasStopRequested returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


