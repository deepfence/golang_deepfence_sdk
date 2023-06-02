# IngestersCloudComplianceScanStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**ComplianceCheckTypes** | Pointer to **[]string** |  | [optional] 
**Result** | Pointer to [**IngestersComplianceStats**](IngestersComplianceStats.md) |  | [optional] 
**ScanId** | Pointer to **string** |  | [optional] 
**ScanMessage** | Pointer to **string** |  | [optional] 
**ScanStatus** | Pointer to **string** |  | [optional] 
**TotalChecks** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewIngestersCloudComplianceScanStatus

`func NewIngestersCloudComplianceScanStatus() *IngestersCloudComplianceScanStatus`

NewIngestersCloudComplianceScanStatus instantiates a new IngestersCloudComplianceScanStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestersCloudComplianceScanStatusWithDefaults

`func NewIngestersCloudComplianceScanStatusWithDefaults() *IngestersCloudComplianceScanStatus`

NewIngestersCloudComplianceScanStatusWithDefaults instantiates a new IngestersCloudComplianceScanStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *IngestersCloudComplianceScanStatus) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *IngestersCloudComplianceScanStatus) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *IngestersCloudComplianceScanStatus) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *IngestersCloudComplianceScanStatus) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetComplianceCheckTypes

`func (o *IngestersCloudComplianceScanStatus) GetComplianceCheckTypes() []string`

GetComplianceCheckTypes returns the ComplianceCheckTypes field if non-nil, zero value otherwise.

### GetComplianceCheckTypesOk

`func (o *IngestersCloudComplianceScanStatus) GetComplianceCheckTypesOk() (*[]string, bool)`

GetComplianceCheckTypesOk returns a tuple with the ComplianceCheckTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceCheckTypes

`func (o *IngestersCloudComplianceScanStatus) SetComplianceCheckTypes(v []string)`

SetComplianceCheckTypes sets ComplianceCheckTypes field to given value.

### HasComplianceCheckTypes

`func (o *IngestersCloudComplianceScanStatus) HasComplianceCheckTypes() bool`

HasComplianceCheckTypes returns a boolean if a field has been set.

### SetComplianceCheckTypesNil

`func (o *IngestersCloudComplianceScanStatus) SetComplianceCheckTypesNil(b bool)`

 SetComplianceCheckTypesNil sets the value for ComplianceCheckTypes to be an explicit nil

### UnsetComplianceCheckTypes
`func (o *IngestersCloudComplianceScanStatus) UnsetComplianceCheckTypes()`

UnsetComplianceCheckTypes ensures that no value is present for ComplianceCheckTypes, not even an explicit nil
### GetResult

`func (o *IngestersCloudComplianceScanStatus) GetResult() IngestersComplianceStats`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *IngestersCloudComplianceScanStatus) GetResultOk() (*IngestersComplianceStats, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *IngestersCloudComplianceScanStatus) SetResult(v IngestersComplianceStats)`

SetResult sets Result field to given value.

### HasResult

`func (o *IngestersCloudComplianceScanStatus) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetScanId

`func (o *IngestersCloudComplianceScanStatus) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *IngestersCloudComplianceScanStatus) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *IngestersCloudComplianceScanStatus) SetScanId(v string)`

SetScanId sets ScanId field to given value.

### HasScanId

`func (o *IngestersCloudComplianceScanStatus) HasScanId() bool`

HasScanId returns a boolean if a field has been set.

### GetScanMessage

`func (o *IngestersCloudComplianceScanStatus) GetScanMessage() string`

GetScanMessage returns the ScanMessage field if non-nil, zero value otherwise.

### GetScanMessageOk

`func (o *IngestersCloudComplianceScanStatus) GetScanMessageOk() (*string, bool)`

GetScanMessageOk returns a tuple with the ScanMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanMessage

`func (o *IngestersCloudComplianceScanStatus) SetScanMessage(v string)`

SetScanMessage sets ScanMessage field to given value.

### HasScanMessage

`func (o *IngestersCloudComplianceScanStatus) HasScanMessage() bool`

HasScanMessage returns a boolean if a field has been set.

### GetScanStatus

`func (o *IngestersCloudComplianceScanStatus) GetScanStatus() string`

GetScanStatus returns the ScanStatus field if non-nil, zero value otherwise.

### GetScanStatusOk

`func (o *IngestersCloudComplianceScanStatus) GetScanStatusOk() (*string, bool)`

GetScanStatusOk returns a tuple with the ScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStatus

`func (o *IngestersCloudComplianceScanStatus) SetScanStatus(v string)`

SetScanStatus sets ScanStatus field to given value.

### HasScanStatus

`func (o *IngestersCloudComplianceScanStatus) HasScanStatus() bool`

HasScanStatus returns a boolean if a field has been set.

### GetTotalChecks

`func (o *IngestersCloudComplianceScanStatus) GetTotalChecks() int32`

GetTotalChecks returns the TotalChecks field if non-nil, zero value otherwise.

### GetTotalChecksOk

`func (o *IngestersCloudComplianceScanStatus) GetTotalChecksOk() (*int32, bool)`

GetTotalChecksOk returns a tuple with the TotalChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalChecks

`func (o *IngestersCloudComplianceScanStatus) SetTotalChecks(v int32)`

SetTotalChecks sets TotalChecks field to given value.

### HasTotalChecks

`func (o *IngestersCloudComplianceScanStatus) HasTotalChecks() bool`

HasTotalChecks returns a boolean if a field has been set.

### GetType

`func (o *IngestersCloudComplianceScanStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IngestersCloudComplianceScanStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IngestersCloudComplianceScanStatus) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IngestersCloudComplianceScanStatus) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


