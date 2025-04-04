/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.5.5
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ModelRegistryCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelRegistryCredentials{}

// ModelRegistryCredentials struct for ModelRegistryCredentials
type ModelRegistryCredentials struct {
	Password *string `json:"password,omitempty"`
	RegistryUrl *string `json:"registry_url,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewModelRegistryCredentials instantiates a new ModelRegistryCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelRegistryCredentials() *ModelRegistryCredentials {
	this := ModelRegistryCredentials{}
	return &this
}

// NewModelRegistryCredentialsWithDefaults instantiates a new ModelRegistryCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelRegistryCredentialsWithDefaults() *ModelRegistryCredentials {
	this := ModelRegistryCredentials{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ModelRegistryCredentials) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRegistryCredentials) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ModelRegistryCredentials) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ModelRegistryCredentials) SetPassword(v string) {
	o.Password = &v
}

// GetRegistryUrl returns the RegistryUrl field value if set, zero value otherwise.
func (o *ModelRegistryCredentials) GetRegistryUrl() string {
	if o == nil || IsNil(o.RegistryUrl) {
		var ret string
		return ret
	}
	return *o.RegistryUrl
}

// GetRegistryUrlOk returns a tuple with the RegistryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRegistryCredentials) GetRegistryUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RegistryUrl) {
		return nil, false
	}
	return o.RegistryUrl, true
}

// HasRegistryUrl returns a boolean if a field has been set.
func (o *ModelRegistryCredentials) HasRegistryUrl() bool {
	if o != nil && !IsNil(o.RegistryUrl) {
		return true
	}

	return false
}

// SetRegistryUrl gets a reference to the given string and assigns it to the RegistryUrl field.
func (o *ModelRegistryCredentials) SetRegistryUrl(v string) {
	o.RegistryUrl = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ModelRegistryCredentials) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRegistryCredentials) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ModelRegistryCredentials) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ModelRegistryCredentials) SetUsername(v string) {
	o.Username = &v
}

func (o ModelRegistryCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelRegistryCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.RegistryUrl) {
		toSerialize["registry_url"] = o.RegistryUrl
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableModelRegistryCredentials struct {
	value *ModelRegistryCredentials
	isSet bool
}

func (v NullableModelRegistryCredentials) Get() *ModelRegistryCredentials {
	return v.value
}

func (v *NullableModelRegistryCredentials) Set(val *ModelRegistryCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableModelRegistryCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableModelRegistryCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelRegistryCredentials(val *ModelRegistryCredentials) *NullableModelRegistryCredentials {
	return &NullableModelRegistryCredentials{value: val, isSet: true}
}

func (v NullableModelRegistryCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelRegistryCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


