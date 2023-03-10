/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: 2.0.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ModelUpdateUserIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelUpdateUserIdRequest{}

// ModelUpdateUserIdRequest struct for ModelUpdateUserIdRequest
type ModelUpdateUserIdRequest struct {
	FirstName *string `json:"first_name,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	Role string `json:"role"`
}

// NewModelUpdateUserIdRequest instantiates a new ModelUpdateUserIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelUpdateUserIdRequest(role string) *ModelUpdateUserIdRequest {
	this := ModelUpdateUserIdRequest{}
	this.Role = role
	return &this
}

// NewModelUpdateUserIdRequestWithDefaults instantiates a new ModelUpdateUserIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelUpdateUserIdRequestWithDefaults() *ModelUpdateUserIdRequest {
	this := ModelUpdateUserIdRequest{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ModelUpdateUserIdRequest) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUpdateUserIdRequest) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ModelUpdateUserIdRequest) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ModelUpdateUserIdRequest) SetFirstName(v string) {
	o.FirstName = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *ModelUpdateUserIdRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUpdateUserIdRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *ModelUpdateUserIdRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *ModelUpdateUserIdRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ModelUpdateUserIdRequest) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUpdateUserIdRequest) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ModelUpdateUserIdRequest) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ModelUpdateUserIdRequest) SetLastName(v string) {
	o.LastName = &v
}

// GetRole returns the Role field value
func (o *ModelUpdateUserIdRequest) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *ModelUpdateUserIdRequest) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *ModelUpdateUserIdRequest) SetRole(v string) {
	o.Role = v
}

func (o ModelUpdateUserIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelUpdateUserIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	toSerialize["role"] = o.Role
	return toSerialize, nil
}

type NullableModelUpdateUserIdRequest struct {
	value *ModelUpdateUserIdRequest
	isSet bool
}

func (v NullableModelUpdateUserIdRequest) Get() *ModelUpdateUserIdRequest {
	return v.value
}

func (v *NullableModelUpdateUserIdRequest) Set(val *ModelUpdateUserIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelUpdateUserIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelUpdateUserIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelUpdateUserIdRequest(val *ModelUpdateUserIdRequest) *NullableModelUpdateUserIdRequest {
	return &NullableModelUpdateUserIdRequest{value: val, isSet: true}
}

func (v NullableModelUpdateUserIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelUpdateUserIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


