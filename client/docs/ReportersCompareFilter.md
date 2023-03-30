# ReportersCompareFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | **string** |  | 
**FieldValue** | **interface{}** |  | 
**GreaterThan** | **bool** |  | 

## Methods

### NewReportersCompareFilter

`func NewReportersCompareFilter(fieldName string, fieldValue interface{}, greaterThan bool, ) *ReportersCompareFilter`

NewReportersCompareFilter instantiates a new ReportersCompareFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportersCompareFilterWithDefaults

`func NewReportersCompareFilterWithDefaults() *ReportersCompareFilter`

NewReportersCompareFilterWithDefaults instantiates a new ReportersCompareFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *ReportersCompareFilter) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *ReportersCompareFilter) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *ReportersCompareFilter) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetFieldValue

`func (o *ReportersCompareFilter) GetFieldValue() interface{}`

GetFieldValue returns the FieldValue field if non-nil, zero value otherwise.

### GetFieldValueOk

`func (o *ReportersCompareFilter) GetFieldValueOk() (*interface{}, bool)`

GetFieldValueOk returns a tuple with the FieldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValue

`func (o *ReportersCompareFilter) SetFieldValue(v interface{})`

SetFieldValue sets FieldValue field to given value.


### SetFieldValueNil

`func (o *ReportersCompareFilter) SetFieldValueNil(b bool)`

 SetFieldValueNil sets the value for FieldValue to be an explicit nil

### UnsetFieldValue
`func (o *ReportersCompareFilter) UnsetFieldValue()`

UnsetFieldValue ensures that no value is present for FieldValue, not even an explicit nil
### GetGreaterThan

`func (o *ReportersCompareFilter) GetGreaterThan() bool`

GetGreaterThan returns the GreaterThan field if non-nil, zero value otherwise.

### GetGreaterThanOk

`func (o *ReportersCompareFilter) GetGreaterThanOk() (*bool, bool)`

GetGreaterThanOk returns a tuple with the GreaterThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreaterThan

`func (o *ReportersCompareFilter) SetGreaterThan(v bool)`

SetGreaterThan sets GreaterThan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


