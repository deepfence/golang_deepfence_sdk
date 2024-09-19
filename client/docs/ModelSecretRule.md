# ModelSecretRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | **string** |  | 
**Masked** | **bool** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Part** | Pointer to **string** |  | [optional] 
**Payload** | **string** |  | 
**RuleId** | Pointer to **string** |  | [optional] 
**Severity** | **string** |  | 
**SignatureToMatch** | Pointer to **string** |  | [optional] 
**Summary** | **string** |  | 
**UpdatedAt** | **int32** |  | 

## Methods

### NewModelSecretRule

`func NewModelSecretRule(level string, masked bool, payload string, severity string, summary string, updatedAt int32, ) *ModelSecretRule`

NewModelSecretRule instantiates a new ModelSecretRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSecretRuleWithDefaults

`func NewModelSecretRuleWithDefaults() *ModelSecretRule`

NewModelSecretRuleWithDefaults instantiates a new ModelSecretRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *ModelSecretRule) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ModelSecretRule) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ModelSecretRule) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetMasked

`func (o *ModelSecretRule) GetMasked() bool`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *ModelSecretRule) GetMaskedOk() (*bool, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *ModelSecretRule) SetMasked(v bool)`

SetMasked sets Masked field to given value.


### GetName

`func (o *ModelSecretRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelSecretRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelSecretRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelSecretRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPart

`func (o *ModelSecretRule) GetPart() string`

GetPart returns the Part field if non-nil, zero value otherwise.

### GetPartOk

`func (o *ModelSecretRule) GetPartOk() (*string, bool)`

GetPartOk returns a tuple with the Part field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPart

`func (o *ModelSecretRule) SetPart(v string)`

SetPart sets Part field to given value.

### HasPart

`func (o *ModelSecretRule) HasPart() bool`

HasPart returns a boolean if a field has been set.

### GetPayload

`func (o *ModelSecretRule) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ModelSecretRule) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ModelSecretRule) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetRuleId

`func (o *ModelSecretRule) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *ModelSecretRule) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *ModelSecretRule) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *ModelSecretRule) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetSeverity

`func (o *ModelSecretRule) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ModelSecretRule) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ModelSecretRule) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetSignatureToMatch

`func (o *ModelSecretRule) GetSignatureToMatch() string`

GetSignatureToMatch returns the SignatureToMatch field if non-nil, zero value otherwise.

### GetSignatureToMatchOk

`func (o *ModelSecretRule) GetSignatureToMatchOk() (*string, bool)`

GetSignatureToMatchOk returns a tuple with the SignatureToMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureToMatch

`func (o *ModelSecretRule) SetSignatureToMatch(v string)`

SetSignatureToMatch sets SignatureToMatch field to given value.

### HasSignatureToMatch

`func (o *ModelSecretRule) HasSignatureToMatch() bool`

HasSignatureToMatch returns a boolean if a field has been set.

### GetSummary

`func (o *ModelSecretRule) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ModelSecretRule) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ModelSecretRule) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetUpdatedAt

`func (o *ModelSecretRule) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelSecretRule) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelSecretRule) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


