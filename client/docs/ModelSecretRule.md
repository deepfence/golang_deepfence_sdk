# ModelSecretRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Level** | **string** |  | 
**Masked** | **bool** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Part** | Pointer to **string** |  | [optional] 
**SignatureToMatch** | Pointer to **string** |  | [optional] 
**UpdatedAt** | **int32** |  | 

## Methods

### NewModelSecretRule

`func NewModelSecretRule(level string, masked bool, updatedAt int32, ) *ModelSecretRule`

NewModelSecretRule instantiates a new ModelSecretRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSecretRuleWithDefaults

`func NewModelSecretRuleWithDefaults() *ModelSecretRule`

NewModelSecretRuleWithDefaults instantiates a new ModelSecretRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelSecretRule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelSecretRule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelSecretRule) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelSecretRule) HasId() bool`

HasId returns a boolean if a field has been set.

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


