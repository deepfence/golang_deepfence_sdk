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
	"bytes"
	"fmt"
)

// checks if the ModelPasswordResetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelPasswordResetRequest{}

// ModelPasswordResetRequest struct for ModelPasswordResetRequest
type ModelPasswordResetRequest struct {
	Email string `json:"email"`
}

type _ModelPasswordResetRequest ModelPasswordResetRequest

// NewModelPasswordResetRequest instantiates a new ModelPasswordResetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelPasswordResetRequest(email string) *ModelPasswordResetRequest {
	this := ModelPasswordResetRequest{}
	this.Email = email
	return &this
}

// NewModelPasswordResetRequestWithDefaults instantiates a new ModelPasswordResetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelPasswordResetRequestWithDefaults() *ModelPasswordResetRequest {
	this := ModelPasswordResetRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *ModelPasswordResetRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ModelPasswordResetRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ModelPasswordResetRequest) SetEmail(v string) {
	o.Email = v
}

func (o ModelPasswordResetRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelPasswordResetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

func (o *ModelPasswordResetRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varModelPasswordResetRequest := _ModelPasswordResetRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelPasswordResetRequest)

	if err != nil {
		return err
	}

	*o = ModelPasswordResetRequest(varModelPasswordResetRequest)

	return err
}

type NullableModelPasswordResetRequest struct {
	value *ModelPasswordResetRequest
	isSet bool
}

func (v NullableModelPasswordResetRequest) Get() *ModelPasswordResetRequest {
	return v.value
}

func (v *NullableModelPasswordResetRequest) Set(val *ModelPasswordResetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelPasswordResetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelPasswordResetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelPasswordResetRequest(val *ModelPasswordResetRequest) *NullableModelPasswordResetRequest {
	return &NullableModelPasswordResetRequest{value: val, isSet: true}
}

func (v NullableModelPasswordResetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelPasswordResetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


