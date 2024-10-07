/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.4.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ModelRegisterLicenseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelRegisterLicenseRequest{}

// ModelRegisterLicenseRequest struct for ModelRegisterLicenseRequest
type ModelRegisterLicenseRequest struct {
	Email *string `json:"email,omitempty"`
	LicenseKey string `json:"license_key"`
}

type _ModelRegisterLicenseRequest ModelRegisterLicenseRequest

// NewModelRegisterLicenseRequest instantiates a new ModelRegisterLicenseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelRegisterLicenseRequest(licenseKey string) *ModelRegisterLicenseRequest {
	this := ModelRegisterLicenseRequest{}
	this.LicenseKey = licenseKey
	return &this
}

// NewModelRegisterLicenseRequestWithDefaults instantiates a new ModelRegisterLicenseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelRegisterLicenseRequestWithDefaults() *ModelRegisterLicenseRequest {
	this := ModelRegisterLicenseRequest{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ModelRegisterLicenseRequest) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRegisterLicenseRequest) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ModelRegisterLicenseRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ModelRegisterLicenseRequest) SetEmail(v string) {
	o.Email = &v
}

// GetLicenseKey returns the LicenseKey field value
func (o *ModelRegisterLicenseRequest) GetLicenseKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LicenseKey
}

// GetLicenseKeyOk returns a tuple with the LicenseKey field value
// and a boolean to check if the value has been set.
func (o *ModelRegisterLicenseRequest) GetLicenseKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LicenseKey, true
}

// SetLicenseKey sets field value
func (o *ModelRegisterLicenseRequest) SetLicenseKey(v string) {
	o.LicenseKey = v
}

func (o ModelRegisterLicenseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelRegisterLicenseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	toSerialize["license_key"] = o.LicenseKey
	return toSerialize, nil
}

func (o *ModelRegisterLicenseRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"license_key",
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

	varModelRegisterLicenseRequest := _ModelRegisterLicenseRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelRegisterLicenseRequest)

	if err != nil {
		return err
	}

	*o = ModelRegisterLicenseRequest(varModelRegisterLicenseRequest)

	return err
}

type NullableModelRegisterLicenseRequest struct {
	value *ModelRegisterLicenseRequest
	isSet bool
}

func (v NullableModelRegisterLicenseRequest) Get() *ModelRegisterLicenseRequest {
	return v.value
}

func (v *NullableModelRegisterLicenseRequest) Set(val *ModelRegisterLicenseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelRegisterLicenseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelRegisterLicenseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelRegisterLicenseRequest(val *ModelRegisterLicenseRequest) *NullableModelRegisterLicenseRequest {
	return &NullableModelRegisterLicenseRequest{value: val, isSet: true}
}

func (v NullableModelRegisterLicenseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelRegisterLicenseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


