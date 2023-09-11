# ModelCloudNodeAccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**CloudProvider** | Pointer to **string** |  | [optional] 
**CompliancePercentage** | Pointer to **float32** |  | [optional] 
**LastScanId** | Pointer to **string** |  | [optional] 
**LastScanStatus** | Pointer to **string** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**NodeName** | Pointer to **string** |  | [optional] 
**ScanStatusMap** | Pointer to **map[string]int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewModelCloudNodeAccountInfo

`func NewModelCloudNodeAccountInfo() *ModelCloudNodeAccountInfo`

NewModelCloudNodeAccountInfo instantiates a new ModelCloudNodeAccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCloudNodeAccountInfoWithDefaults

`func NewModelCloudNodeAccountInfoWithDefaults() *ModelCloudNodeAccountInfo`

NewModelCloudNodeAccountInfoWithDefaults instantiates a new ModelCloudNodeAccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ModelCloudNodeAccountInfo) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ModelCloudNodeAccountInfo) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ModelCloudNodeAccountInfo) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ModelCloudNodeAccountInfo) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCloudProvider

`func (o *ModelCloudNodeAccountInfo) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ModelCloudNodeAccountInfo) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ModelCloudNodeAccountInfo) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *ModelCloudNodeAccountInfo) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCompliancePercentage

`func (o *ModelCloudNodeAccountInfo) GetCompliancePercentage() float32`

GetCompliancePercentage returns the CompliancePercentage field if non-nil, zero value otherwise.

### GetCompliancePercentageOk

`func (o *ModelCloudNodeAccountInfo) GetCompliancePercentageOk() (*float32, bool)`

GetCompliancePercentageOk returns a tuple with the CompliancePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliancePercentage

`func (o *ModelCloudNodeAccountInfo) SetCompliancePercentage(v float32)`

SetCompliancePercentage sets CompliancePercentage field to given value.

### HasCompliancePercentage

`func (o *ModelCloudNodeAccountInfo) HasCompliancePercentage() bool`

HasCompliancePercentage returns a boolean if a field has been set.

### GetLastScanId

`func (o *ModelCloudNodeAccountInfo) GetLastScanId() string`

GetLastScanId returns the LastScanId field if non-nil, zero value otherwise.

### GetLastScanIdOk

`func (o *ModelCloudNodeAccountInfo) GetLastScanIdOk() (*string, bool)`

GetLastScanIdOk returns a tuple with the LastScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScanId

`func (o *ModelCloudNodeAccountInfo) SetLastScanId(v string)`

SetLastScanId sets LastScanId field to given value.

### HasLastScanId

`func (o *ModelCloudNodeAccountInfo) HasLastScanId() bool`

HasLastScanId returns a boolean if a field has been set.

### GetLastScanStatus

`func (o *ModelCloudNodeAccountInfo) GetLastScanStatus() string`

GetLastScanStatus returns the LastScanStatus field if non-nil, zero value otherwise.

### GetLastScanStatusOk

`func (o *ModelCloudNodeAccountInfo) GetLastScanStatusOk() (*string, bool)`

GetLastScanStatusOk returns a tuple with the LastScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScanStatus

`func (o *ModelCloudNodeAccountInfo) SetLastScanStatus(v string)`

SetLastScanStatus sets LastScanStatus field to given value.

### HasLastScanStatus

`func (o *ModelCloudNodeAccountInfo) HasLastScanStatus() bool`

HasLastScanStatus returns a boolean if a field has been set.

### GetNodeId

`func (o *ModelCloudNodeAccountInfo) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelCloudNodeAccountInfo) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelCloudNodeAccountInfo) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ModelCloudNodeAccountInfo) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNodeName

`func (o *ModelCloudNodeAccountInfo) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *ModelCloudNodeAccountInfo) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *ModelCloudNodeAccountInfo) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *ModelCloudNodeAccountInfo) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetScanStatusMap

`func (o *ModelCloudNodeAccountInfo) GetScanStatusMap() map[string]int32`

GetScanStatusMap returns the ScanStatusMap field if non-nil, zero value otherwise.

### GetScanStatusMapOk

`func (o *ModelCloudNodeAccountInfo) GetScanStatusMapOk() (*map[string]int32, bool)`

GetScanStatusMapOk returns a tuple with the ScanStatusMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStatusMap

`func (o *ModelCloudNodeAccountInfo) SetScanStatusMap(v map[string]int32)`

SetScanStatusMap sets ScanStatusMap field to given value.

### HasScanStatusMap

`func (o *ModelCloudNodeAccountInfo) HasScanStatusMap() bool`

HasScanStatusMap returns a boolean if a field has been set.

### SetScanStatusMapNil

`func (o *ModelCloudNodeAccountInfo) SetScanStatusMapNil(b bool)`

 SetScanStatusMapNil sets the value for ScanStatusMap to be an explicit nil

### UnsetScanStatusMap
`func (o *ModelCloudNodeAccountInfo) UnsetScanStatusMap()`

UnsetScanStatusMap ensures that no value is present for ScanStatusMap, not even an explicit nil
### GetVersion

`func (o *ModelCloudNodeAccountInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelCloudNodeAccountInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelCloudNodeAccountInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ModelCloudNodeAccountInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


