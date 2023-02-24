# ModelUpdateUserIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Role** | **string** |  | 

## Methods

### NewModelUpdateUserIdRequest

`func NewModelUpdateUserIdRequest(role string, ) *ModelUpdateUserIdRequest`

NewModelUpdateUserIdRequest instantiates a new ModelUpdateUserIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelUpdateUserIdRequestWithDefaults

`func NewModelUpdateUserIdRequestWithDefaults() *ModelUpdateUserIdRequest`

NewModelUpdateUserIdRequestWithDefaults instantiates a new ModelUpdateUserIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *ModelUpdateUserIdRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ModelUpdateUserIdRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ModelUpdateUserIdRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ModelUpdateUserIdRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetIsActive

`func (o *ModelUpdateUserIdRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ModelUpdateUserIdRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ModelUpdateUserIdRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ModelUpdateUserIdRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLastName

`func (o *ModelUpdateUserIdRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ModelUpdateUserIdRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ModelUpdateUserIdRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ModelUpdateUserIdRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetRole

`func (o *ModelUpdateUserIdRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ModelUpdateUserIdRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ModelUpdateUserIdRequest) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


