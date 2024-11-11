# ModelGenerateReportReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**UtilsReportFilters**](UtilsReportFilters.md) |  | [optional] 
**FromTimestamp** | Pointer to **int32** |  | [optional] 
**Options** | Pointer to [**UtilsReportOptions**](UtilsReportOptions.md) |  | [optional] 
**ReportType** | **string** |  | 
**ToTimestamp** | Pointer to **int32** |  | [optional] 
**ZippedReport** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelGenerateReportReq

`func NewModelGenerateReportReq(reportType string, ) *ModelGenerateReportReq`

NewModelGenerateReportReq instantiates a new ModelGenerateReportReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelGenerateReportReqWithDefaults

`func NewModelGenerateReportReqWithDefaults() *ModelGenerateReportReq`

NewModelGenerateReportReqWithDefaults instantiates a new ModelGenerateReportReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ModelGenerateReportReq) GetFilters() UtilsReportFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ModelGenerateReportReq) GetFiltersOk() (*UtilsReportFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ModelGenerateReportReq) SetFilters(v UtilsReportFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ModelGenerateReportReq) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetFromTimestamp

`func (o *ModelGenerateReportReq) GetFromTimestamp() int32`

GetFromTimestamp returns the FromTimestamp field if non-nil, zero value otherwise.

### GetFromTimestampOk

`func (o *ModelGenerateReportReq) GetFromTimestampOk() (*int32, bool)`

GetFromTimestampOk returns a tuple with the FromTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTimestamp

`func (o *ModelGenerateReportReq) SetFromTimestamp(v int32)`

SetFromTimestamp sets FromTimestamp field to given value.

### HasFromTimestamp

`func (o *ModelGenerateReportReq) HasFromTimestamp() bool`

HasFromTimestamp returns a boolean if a field has been set.

### GetOptions

`func (o *ModelGenerateReportReq) GetOptions() UtilsReportOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ModelGenerateReportReq) GetOptionsOk() (*UtilsReportOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ModelGenerateReportReq) SetOptions(v UtilsReportOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ModelGenerateReportReq) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetReportType

`func (o *ModelGenerateReportReq) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *ModelGenerateReportReq) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *ModelGenerateReportReq) SetReportType(v string)`

SetReportType sets ReportType field to given value.


### GetToTimestamp

`func (o *ModelGenerateReportReq) GetToTimestamp() int32`

GetToTimestamp returns the ToTimestamp field if non-nil, zero value otherwise.

### GetToTimestampOk

`func (o *ModelGenerateReportReq) GetToTimestampOk() (*int32, bool)`

GetToTimestampOk returns a tuple with the ToTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTimestamp

`func (o *ModelGenerateReportReq) SetToTimestamp(v int32)`

SetToTimestamp sets ToTimestamp field to given value.

### HasToTimestamp

`func (o *ModelGenerateReportReq) HasToTimestamp() bool`

HasToTimestamp returns a boolean if a field has been set.

### GetZippedReport

`func (o *ModelGenerateReportReq) GetZippedReport() bool`

GetZippedReport returns the ZippedReport field if non-nil, zero value otherwise.

### GetZippedReportOk

`func (o *ModelGenerateReportReq) GetZippedReportOk() (*bool, bool)`

GetZippedReportOk returns a tuple with the ZippedReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZippedReport

`func (o *ModelGenerateReportReq) SetZippedReport(v bool)`

SetZippedReport sets ZippedReport field to given value.

### HasZippedReport

`func (o *ModelGenerateReportReq) HasZippedReport() bool`

HasZippedReport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


