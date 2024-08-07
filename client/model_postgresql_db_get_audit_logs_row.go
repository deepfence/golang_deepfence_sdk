/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v3.0.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the PostgresqlDbGetAuditLogsRow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostgresqlDbGetAuditLogsRow{}

// PostgresqlDbGetAuditLogsRow struct for PostgresqlDbGetAuditLogsRow
type PostgresqlDbGetAuditLogsRow struct {
	Action *string `json:"action,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Email *string `json:"email,omitempty"`
	Event *string `json:"event,omitempty"`
	Resources *string `json:"resources,omitempty"`
	Role *string `json:"role,omitempty"`
	Success *bool `json:"success,omitempty"`
}

// NewPostgresqlDbGetAuditLogsRow instantiates a new PostgresqlDbGetAuditLogsRow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgresqlDbGetAuditLogsRow() *PostgresqlDbGetAuditLogsRow {
	this := PostgresqlDbGetAuditLogsRow{}
	return &this
}

// NewPostgresqlDbGetAuditLogsRowWithDefaults instantiates a new PostgresqlDbGetAuditLogsRow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresqlDbGetAuditLogsRowWithDefaults() *PostgresqlDbGetAuditLogsRow {
	this := PostgresqlDbGetAuditLogsRow{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PostgresqlDbGetAuditLogsRow) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlDbGetAuditLogsRow) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PostgresqlDbGetAuditLogsRow) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *PostgresqlDbGetAuditLogsRow) SetAction(v string) {
	o.Action = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PostgresqlDbGetAuditLogsRow) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlDbGetAuditLogsRow) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PostgresqlDbGetAuditLogsRow) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PostgresqlDbGetAuditLogsRow) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PostgresqlDbGetAuditLogsRow) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlDbGetAuditLogsRow) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PostgresqlDbGetAuditLogsRow) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PostgresqlDbGetAuditLogsRow) SetEmail(v string) {
	o.Email = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *PostgresqlDbGetAuditLogsRow) GetEvent() string {
	if o == nil || IsNil(o.Event) {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlDbGetAuditLogsRow) GetEventOk() (*string, bool) {
	if o == nil || IsNil(o.Event) {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *PostgresqlDbGetAuditLogsRow) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *PostgresqlDbGetAuditLogsRow) SetEvent(v string) {
	o.Event = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *PostgresqlDbGetAuditLogsRow) GetResources() string {
	if o == nil || IsNil(o.Resources) {
		var ret string
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlDbGetAuditLogsRow) GetResourcesOk() (*string, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *PostgresqlDbGetAuditLogsRow) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given string and assigns it to the Resources field.
func (o *PostgresqlDbGetAuditLogsRow) SetResources(v string) {
	o.Resources = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *PostgresqlDbGetAuditLogsRow) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlDbGetAuditLogsRow) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *PostgresqlDbGetAuditLogsRow) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *PostgresqlDbGetAuditLogsRow) SetRole(v string) {
	o.Role = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *PostgresqlDbGetAuditLogsRow) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlDbGetAuditLogsRow) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *PostgresqlDbGetAuditLogsRow) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *PostgresqlDbGetAuditLogsRow) SetSuccess(v bool) {
	o.Success = &v
}

func (o PostgresqlDbGetAuditLogsRow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostgresqlDbGetAuditLogsRow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Event) {
		toSerialize["event"] = o.Event
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullablePostgresqlDbGetAuditLogsRow struct {
	value *PostgresqlDbGetAuditLogsRow
	isSet bool
}

func (v NullablePostgresqlDbGetAuditLogsRow) Get() *PostgresqlDbGetAuditLogsRow {
	return v.value
}

func (v *NullablePostgresqlDbGetAuditLogsRow) Set(val *PostgresqlDbGetAuditLogsRow) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgresqlDbGetAuditLogsRow) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgresqlDbGetAuditLogsRow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgresqlDbGetAuditLogsRow(val *PostgresqlDbGetAuditLogsRow) *NullablePostgresqlDbGetAuditLogsRow {
	return &NullablePostgresqlDbGetAuditLogsRow{value: val, isSet: true}
}

func (v NullablePostgresqlDbGetAuditLogsRow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgresqlDbGetAuditLogsRow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


