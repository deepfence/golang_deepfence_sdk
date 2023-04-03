# ModelGenerateReportReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** |  | [optional] 
**Filters** | Pointer to [**UtilsReportFilters**](UtilsReportFilters.md) |  | [optional] 
**ReportType** | **string** |  | 

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

### GetDuration

`func (o *ModelGenerateReportReq) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ModelGenerateReportReq) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ModelGenerateReportReq) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ModelGenerateReportReq) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


