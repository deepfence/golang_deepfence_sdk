# PostgresqlDbScheduler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CronExpr** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**IsSystem** | Pointer to **bool** |  | [optional] 
**LastRanAt** | Pointer to **map[string]interface{}** |  | [optional] 
**Payload** | Pointer to **interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPostgresqlDbScheduler

`func NewPostgresqlDbScheduler() *PostgresqlDbScheduler`

NewPostgresqlDbScheduler instantiates a new PostgresqlDbScheduler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlDbSchedulerWithDefaults

`func NewPostgresqlDbSchedulerWithDefaults() *PostgresqlDbScheduler`

NewPostgresqlDbSchedulerWithDefaults instantiates a new PostgresqlDbScheduler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PostgresqlDbScheduler) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PostgresqlDbScheduler) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PostgresqlDbScheduler) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PostgresqlDbScheduler) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PostgresqlDbScheduler) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PostgresqlDbScheduler) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PostgresqlDbScheduler) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PostgresqlDbScheduler) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCronExpr

`func (o *PostgresqlDbScheduler) GetCronExpr() string`

GetCronExpr returns the CronExpr field if non-nil, zero value otherwise.

### GetCronExprOk

`func (o *PostgresqlDbScheduler) GetCronExprOk() (*string, bool)`

GetCronExprOk returns a tuple with the CronExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpr

`func (o *PostgresqlDbScheduler) SetCronExpr(v string)`

SetCronExpr sets CronExpr field to given value.

### HasCronExpr

`func (o *PostgresqlDbScheduler) HasCronExpr() bool`

HasCronExpr returns a boolean if a field has been set.

### GetDescription

`func (o *PostgresqlDbScheduler) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostgresqlDbScheduler) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostgresqlDbScheduler) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostgresqlDbScheduler) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *PostgresqlDbScheduler) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostgresqlDbScheduler) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostgresqlDbScheduler) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PostgresqlDbScheduler) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsEnabled

`func (o *PostgresqlDbScheduler) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *PostgresqlDbScheduler) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *PostgresqlDbScheduler) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *PostgresqlDbScheduler) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsSystem

`func (o *PostgresqlDbScheduler) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *PostgresqlDbScheduler) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *PostgresqlDbScheduler) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.

### HasIsSystem

`func (o *PostgresqlDbScheduler) HasIsSystem() bool`

HasIsSystem returns a boolean if a field has been set.

### GetLastRanAt

`func (o *PostgresqlDbScheduler) GetLastRanAt() map[string]interface{}`

GetLastRanAt returns the LastRanAt field if non-nil, zero value otherwise.

### GetLastRanAtOk

`func (o *PostgresqlDbScheduler) GetLastRanAtOk() (*map[string]interface{}, bool)`

GetLastRanAtOk returns a tuple with the LastRanAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRanAt

`func (o *PostgresqlDbScheduler) SetLastRanAt(v map[string]interface{})`

SetLastRanAt sets LastRanAt field to given value.

### HasLastRanAt

`func (o *PostgresqlDbScheduler) HasLastRanAt() bool`

HasLastRanAt returns a boolean if a field has been set.

### GetPayload

`func (o *PostgresqlDbScheduler) GetPayload() interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *PostgresqlDbScheduler) GetPayloadOk() (*interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *PostgresqlDbScheduler) SetPayload(v interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *PostgresqlDbScheduler) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *PostgresqlDbScheduler) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *PostgresqlDbScheduler) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetStatus

`func (o *PostgresqlDbScheduler) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostgresqlDbScheduler) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostgresqlDbScheduler) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PostgresqlDbScheduler) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PostgresqlDbScheduler) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PostgresqlDbScheduler) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PostgresqlDbScheduler) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PostgresqlDbScheduler) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


