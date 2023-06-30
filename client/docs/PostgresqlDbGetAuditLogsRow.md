# PostgresqlDbGetAuditLogsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewPostgresqlDbGetAuditLogsRow

`func NewPostgresqlDbGetAuditLogsRow() *PostgresqlDbGetAuditLogsRow`

NewPostgresqlDbGetAuditLogsRow instantiates a new PostgresqlDbGetAuditLogsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlDbGetAuditLogsRowWithDefaults

`func NewPostgresqlDbGetAuditLogsRowWithDefaults() *PostgresqlDbGetAuditLogsRow`

NewPostgresqlDbGetAuditLogsRowWithDefaults instantiates a new PostgresqlDbGetAuditLogsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PostgresqlDbGetAuditLogsRow) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PostgresqlDbGetAuditLogsRow) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PostgresqlDbGetAuditLogsRow) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PostgresqlDbGetAuditLogsRow) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PostgresqlDbGetAuditLogsRow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PostgresqlDbGetAuditLogsRow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PostgresqlDbGetAuditLogsRow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PostgresqlDbGetAuditLogsRow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *PostgresqlDbGetAuditLogsRow) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PostgresqlDbGetAuditLogsRow) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PostgresqlDbGetAuditLogsRow) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PostgresqlDbGetAuditLogsRow) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEvent

`func (o *PostgresqlDbGetAuditLogsRow) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *PostgresqlDbGetAuditLogsRow) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *PostgresqlDbGetAuditLogsRow) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *PostgresqlDbGetAuditLogsRow) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetResources

`func (o *PostgresqlDbGetAuditLogsRow) GetResources() string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *PostgresqlDbGetAuditLogsRow) GetResourcesOk() (*string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *PostgresqlDbGetAuditLogsRow) SetResources(v string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *PostgresqlDbGetAuditLogsRow) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRole

`func (o *PostgresqlDbGetAuditLogsRow) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PostgresqlDbGetAuditLogsRow) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PostgresqlDbGetAuditLogsRow) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *PostgresqlDbGetAuditLogsRow) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSuccess

`func (o *PostgresqlDbGetAuditLogsRow) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *PostgresqlDbGetAuditLogsRow) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *PostgresqlDbGetAuditLogsRow) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *PostgresqlDbGetAuditLogsRow) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


