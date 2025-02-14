# CompletionCompletionNodeFieldReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completion** | **string** |  | 
**FieldName** | **string** |  | 
**Filters** | Pointer to [**ReportersFieldsFilters**](ReportersFieldsFilters.md) |  | [optional] 
**IsArrayField** | Pointer to **bool** |  | [optional] 
**ScanId** | Pointer to **string** |  | [optional] 
**Window** | [**ModelFetchWindow**](ModelFetchWindow.md) |  | 

## Methods

### NewCompletionCompletionNodeFieldReq

`func NewCompletionCompletionNodeFieldReq(completion string, fieldName string, window ModelFetchWindow, ) *CompletionCompletionNodeFieldReq`

NewCompletionCompletionNodeFieldReq instantiates a new CompletionCompletionNodeFieldReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompletionCompletionNodeFieldReqWithDefaults

`func NewCompletionCompletionNodeFieldReqWithDefaults() *CompletionCompletionNodeFieldReq`

NewCompletionCompletionNodeFieldReqWithDefaults instantiates a new CompletionCompletionNodeFieldReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletion

`func (o *CompletionCompletionNodeFieldReq) GetCompletion() string`

GetCompletion returns the Completion field if non-nil, zero value otherwise.

### GetCompletionOk

`func (o *CompletionCompletionNodeFieldReq) GetCompletionOk() (*string, bool)`

GetCompletionOk returns a tuple with the Completion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletion

`func (o *CompletionCompletionNodeFieldReq) SetCompletion(v string)`

SetCompletion sets Completion field to given value.


### GetFieldName

`func (o *CompletionCompletionNodeFieldReq) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *CompletionCompletionNodeFieldReq) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *CompletionCompletionNodeFieldReq) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetFilters

`func (o *CompletionCompletionNodeFieldReq) GetFilters() ReportersFieldsFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CompletionCompletionNodeFieldReq) GetFiltersOk() (*ReportersFieldsFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CompletionCompletionNodeFieldReq) SetFilters(v ReportersFieldsFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *CompletionCompletionNodeFieldReq) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetIsArrayField

`func (o *CompletionCompletionNodeFieldReq) GetIsArrayField() bool`

GetIsArrayField returns the IsArrayField field if non-nil, zero value otherwise.

### GetIsArrayFieldOk

`func (o *CompletionCompletionNodeFieldReq) GetIsArrayFieldOk() (*bool, bool)`

GetIsArrayFieldOk returns a tuple with the IsArrayField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArrayField

`func (o *CompletionCompletionNodeFieldReq) SetIsArrayField(v bool)`

SetIsArrayField sets IsArrayField field to given value.

### HasIsArrayField

`func (o *CompletionCompletionNodeFieldReq) HasIsArrayField() bool`

HasIsArrayField returns a boolean if a field has been set.

### GetScanId

`func (o *CompletionCompletionNodeFieldReq) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *CompletionCompletionNodeFieldReq) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *CompletionCompletionNodeFieldReq) SetScanId(v string)`

SetScanId sets ScanId field to given value.

### HasScanId

`func (o *CompletionCompletionNodeFieldReq) HasScanId() bool`

HasScanId returns a boolean if a field has been set.

### GetWindow

`func (o *CompletionCompletionNodeFieldReq) GetWindow() ModelFetchWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *CompletionCompletionNodeFieldReq) GetWindowOk() (*ModelFetchWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *CompletionCompletionNodeFieldReq) SetWindow(v ModelFetchWindow)`

SetWindow sets Window field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


