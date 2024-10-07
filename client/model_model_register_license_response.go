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

// checks if the ModelRegisterLicenseResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelRegisterLicenseResponse{}

// ModelRegisterLicenseResponse struct for ModelRegisterLicenseResponse
type ModelRegisterLicenseResponse struct {
	EmailDomain string `json:"email_domain"`
	LicenseKey string `json:"license_key"`
}

type _ModelRegisterLicenseResponse ModelRegisterLicenseResponse

// NewModelRegisterLicenseResponse instantiates a new ModelRegisterLicenseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelRegisterLicenseResponse(emailDomain string, licenseKey string) *ModelRegisterLicenseResponse {
	this := ModelRegisterLicenseResponse{}
	this.EmailDomain = emailDomain
	this.LicenseKey = licenseKey
	return &this
}

// NewModelRegisterLicenseResponseWithDefaults instantiates a new ModelRegisterLicenseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelRegisterLicenseResponseWithDefaults() *ModelRegisterLicenseResponse {
	this := ModelRegisterLicenseResponse{}
	return &this
}

// GetEmailDomain returns the EmailDomain field value
func (o *ModelRegisterLicenseResponse) GetEmailDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailDomain
}

// GetEmailDomainOk returns a tuple with the EmailDomain field value
// and a boolean to check if the value has been set.
func (o *ModelRegisterLicenseResponse) GetEmailDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailDomain, true
}

// SetEmailDomain sets field value
func (o *ModelRegisterLicenseResponse) SetEmailDomain(v string) {
	o.EmailDomain = v
}

// GetLicenseKey returns the LicenseKey field value
func (o *ModelRegisterLicenseResponse) GetLicenseKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LicenseKey
}

// GetLicenseKeyOk returns a tuple with the LicenseKey field value
// and a boolean to check if the value has been set.
func (o *ModelRegisterLicenseResponse) GetLicenseKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LicenseKey, true
}

// SetLicenseKey sets field value
func (o *ModelRegisterLicenseResponse) SetLicenseKey(v string) {
	o.LicenseKey = v
}

func (o ModelRegisterLicenseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelRegisterLicenseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email_domain"] = o.EmailDomain
	toSerialize["license_key"] = o.LicenseKey
	return toSerialize, nil
}

func (o *ModelRegisterLicenseResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email_domain",
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

	varModelRegisterLicenseResponse := _ModelRegisterLicenseResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelRegisterLicenseResponse)

	if err != nil {
		return err
	}

	*o = ModelRegisterLicenseResponse(varModelRegisterLicenseResponse)

	return err
}

type NullableModelRegisterLicenseResponse struct {
	value *ModelRegisterLicenseResponse
	isSet bool
}

func (v NullableModelRegisterLicenseResponse) Get() *ModelRegisterLicenseResponse {
	return v.value
}

func (v *NullableModelRegisterLicenseResponse) Set(val *ModelRegisterLicenseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelRegisterLicenseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelRegisterLicenseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelRegisterLicenseResponse(val *ModelRegisterLicenseResponse) *NullableModelRegisterLicenseResponse {
	return &NullableModelRegisterLicenseResponse{value: val, isSet: true}
}

func (v NullableModelRegisterLicenseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelRegisterLicenseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


