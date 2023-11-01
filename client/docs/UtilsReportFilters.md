# UtilsReportFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedReportFilters** | Pointer to [**UtilsAdvancedReportFilters**](UtilsAdvancedReportFilters.md) |  | [optional] 
**IncludeDeadNodes** | Pointer to **bool** |  | [optional] 
**MostExploitableReport** | Pointer to **bool** |  | [optional] 
**NodeType** | **string** |  | 
**ScanId** | Pointer to **string** |  | [optional] 
**ScanType** | **string** |  | 
**SeverityOrCheckType** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUtilsReportFilters

`func NewUtilsReportFilters(nodeType string, scanType string, ) *UtilsReportFilters`

NewUtilsReportFilters instantiates a new UtilsReportFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilsReportFiltersWithDefaults

`func NewUtilsReportFiltersWithDefaults() *UtilsReportFilters`

NewUtilsReportFiltersWithDefaults instantiates a new UtilsReportFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedReportFilters

`func (o *UtilsReportFilters) GetAdvancedReportFilters() UtilsAdvancedReportFilters`

GetAdvancedReportFilters returns the AdvancedReportFilters field if non-nil, zero value otherwise.

### GetAdvancedReportFiltersOk

`func (o *UtilsReportFilters) GetAdvancedReportFiltersOk() (*UtilsAdvancedReportFilters, bool)`

GetAdvancedReportFiltersOk returns a tuple with the AdvancedReportFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedReportFilters

`func (o *UtilsReportFilters) SetAdvancedReportFilters(v UtilsAdvancedReportFilters)`

SetAdvancedReportFilters sets AdvancedReportFilters field to given value.

### HasAdvancedReportFilters

`func (o *UtilsReportFilters) HasAdvancedReportFilters() bool`

HasAdvancedReportFilters returns a boolean if a field has been set.

### GetIncludeDeadNodes

`func (o *UtilsReportFilters) GetIncludeDeadNodes() bool`

GetIncludeDeadNodes returns the IncludeDeadNodes field if non-nil, zero value otherwise.

### GetIncludeDeadNodesOk

`func (o *UtilsReportFilters) GetIncludeDeadNodesOk() (*bool, bool)`

GetIncludeDeadNodesOk returns a tuple with the IncludeDeadNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDeadNodes

`func (o *UtilsReportFilters) SetIncludeDeadNodes(v bool)`

SetIncludeDeadNodes sets IncludeDeadNodes field to given value.

### HasIncludeDeadNodes

`func (o *UtilsReportFilters) HasIncludeDeadNodes() bool`

HasIncludeDeadNodes returns a boolean if a field has been set.

### GetMostExploitableReport

`func (o *UtilsReportFilters) GetMostExploitableReport() bool`

GetMostExploitableReport returns the MostExploitableReport field if non-nil, zero value otherwise.

### GetMostExploitableReportOk

`func (o *UtilsReportFilters) GetMostExploitableReportOk() (*bool, bool)`

GetMostExploitableReportOk returns a tuple with the MostExploitableReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostExploitableReport

`func (o *UtilsReportFilters) SetMostExploitableReport(v bool)`

SetMostExploitableReport sets MostExploitableReport field to given value.

### HasMostExploitableReport

`func (o *UtilsReportFilters) HasMostExploitableReport() bool`

HasMostExploitableReport returns a boolean if a field has been set.

### GetNodeType

`func (o *UtilsReportFilters) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *UtilsReportFilters) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *UtilsReportFilters) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetScanId

`func (o *UtilsReportFilters) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *UtilsReportFilters) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *UtilsReportFilters) SetScanId(v string)`

SetScanId sets ScanId field to given value.

### HasScanId

`func (o *UtilsReportFilters) HasScanId() bool`

HasScanId returns a boolean if a field has been set.

### GetScanType

`func (o *UtilsReportFilters) GetScanType() string`

GetScanType returns the ScanType field if non-nil, zero value otherwise.

### GetScanTypeOk

`func (o *UtilsReportFilters) GetScanTypeOk() (*string, bool)`

GetScanTypeOk returns a tuple with the ScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanType

`func (o *UtilsReportFilters) SetScanType(v string)`

SetScanType sets ScanType field to given value.


### GetSeverityOrCheckType

`func (o *UtilsReportFilters) GetSeverityOrCheckType() []string`

GetSeverityOrCheckType returns the SeverityOrCheckType field if non-nil, zero value otherwise.

### GetSeverityOrCheckTypeOk

`func (o *UtilsReportFilters) GetSeverityOrCheckTypeOk() (*[]string, bool)`

GetSeverityOrCheckTypeOk returns a tuple with the SeverityOrCheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityOrCheckType

`func (o *UtilsReportFilters) SetSeverityOrCheckType(v []string)`

SetSeverityOrCheckType sets SeverityOrCheckType field to given value.

### HasSeverityOrCheckType

`func (o *UtilsReportFilters) HasSeverityOrCheckType() bool`

HasSeverityOrCheckType returns a boolean if a field has been set.

### SetSeverityOrCheckTypeNil

`func (o *UtilsReportFilters) SetSeverityOrCheckTypeNil(b bool)`

 SetSeverityOrCheckTypeNil sets the value for SeverityOrCheckType to be an explicit nil

### UnsetSeverityOrCheckType
`func (o *UtilsReportFilters) UnsetSeverityOrCheckType()`

UnsetSeverityOrCheckType ensures that no value is present for SeverityOrCheckType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


