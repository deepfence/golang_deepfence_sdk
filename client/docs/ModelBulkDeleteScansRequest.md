# ModelBulkDeleteScansRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**ReportersFieldsFilters**](ReportersFieldsFilters.md) |  | 
**ScanType** | **string** |  | 

## Methods

### NewModelBulkDeleteScansRequest

`func NewModelBulkDeleteScansRequest(filters ReportersFieldsFilters, scanType string, ) *ModelBulkDeleteScansRequest`

NewModelBulkDeleteScansRequest instantiates a new ModelBulkDeleteScansRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelBulkDeleteScansRequestWithDefaults

`func NewModelBulkDeleteScansRequestWithDefaults() *ModelBulkDeleteScansRequest`

NewModelBulkDeleteScansRequestWithDefaults instantiates a new ModelBulkDeleteScansRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ModelBulkDeleteScansRequest) GetFilters() ReportersFieldsFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ModelBulkDeleteScansRequest) GetFiltersOk() (*ReportersFieldsFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ModelBulkDeleteScansRequest) SetFilters(v ReportersFieldsFilters)`

SetFilters sets Filters field to given value.


### GetScanType

`func (o *ModelBulkDeleteScansRequest) GetScanType() string`

GetScanType returns the ScanType field if non-nil, zero value otherwise.

### GetScanTypeOk

`func (o *ModelBulkDeleteScansRequest) GetScanTypeOk() (*string, bool)`

GetScanTypeOk returns a tuple with the ScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanType

`func (o *ModelBulkDeleteScansRequest) SetScanType(v string)`

SetScanType sets ScanType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


