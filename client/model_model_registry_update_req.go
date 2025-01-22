/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.5.3
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ModelRegistryUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelRegistryUpdateReq{}

// ModelRegistryUpdateReq struct for ModelRegistryUpdateReq
type ModelRegistryUpdateReq struct {
	Extras map[string]interface{} `json:"extras,omitempty"`
	Name string `json:"name"`
	NonSecret map[string]interface{} `json:"non_secret,omitempty"`
	RegistryType string `json:"registry_type"`
	Secret map[string]interface{} `json:"secret,omitempty"`
}

type _ModelRegistryUpdateReq ModelRegistryUpdateReq

// NewModelRegistryUpdateReq instantiates a new ModelRegistryUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelRegistryUpdateReq(name string, registryType string) *ModelRegistryUpdateReq {
	this := ModelRegistryUpdateReq{}
	this.Name = name
	this.RegistryType = registryType
	return &this
}

// NewModelRegistryUpdateReqWithDefaults instantiates a new ModelRegistryUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelRegistryUpdateReqWithDefaults() *ModelRegistryUpdateReq {
	this := ModelRegistryUpdateReq{}
	return &this
}

// GetExtras returns the Extras field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelRegistryUpdateReq) GetExtras() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelRegistryUpdateReq) GetExtrasOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Extras) {
		return map[string]interface{}{}, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *ModelRegistryUpdateReq) HasExtras() bool {
	if o != nil && !IsNil(o.Extras) {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *ModelRegistryUpdateReq) SetExtras(v map[string]interface{}) {
	o.Extras = v
}

// GetName returns the Name field value
func (o *ModelRegistryUpdateReq) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModelRegistryUpdateReq) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ModelRegistryUpdateReq) SetName(v string) {
	o.Name = v
}

// GetNonSecret returns the NonSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelRegistryUpdateReq) GetNonSecret() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.NonSecret
}

// GetNonSecretOk returns a tuple with the NonSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelRegistryUpdateReq) GetNonSecretOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.NonSecret) {
		return map[string]interface{}{}, false
	}
	return o.NonSecret, true
}

// HasNonSecret returns a boolean if a field has been set.
func (o *ModelRegistryUpdateReq) HasNonSecret() bool {
	if o != nil && !IsNil(o.NonSecret) {
		return true
	}

	return false
}

// SetNonSecret gets a reference to the given map[string]interface{} and assigns it to the NonSecret field.
func (o *ModelRegistryUpdateReq) SetNonSecret(v map[string]interface{}) {
	o.NonSecret = v
}

// GetRegistryType returns the RegistryType field value
func (o *ModelRegistryUpdateReq) GetRegistryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegistryType
}

// GetRegistryTypeOk returns a tuple with the RegistryType field value
// and a boolean to check if the value has been set.
func (o *ModelRegistryUpdateReq) GetRegistryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegistryType, true
}

// SetRegistryType sets field value
func (o *ModelRegistryUpdateReq) SetRegistryType(v string) {
	o.RegistryType = v
}

// GetSecret returns the Secret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelRegistryUpdateReq) GetSecret() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelRegistryUpdateReq) GetSecretOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Secret) {
		return map[string]interface{}{}, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ModelRegistryUpdateReq) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given map[string]interface{} and assigns it to the Secret field.
func (o *ModelRegistryUpdateReq) SetSecret(v map[string]interface{}) {
	o.Secret = v
}

func (o ModelRegistryUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelRegistryUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	toSerialize["name"] = o.Name
	if o.NonSecret != nil {
		toSerialize["non_secret"] = o.NonSecret
	}
	toSerialize["registry_type"] = o.RegistryType
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	return toSerialize, nil
}

func (o *ModelRegistryUpdateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"registry_type",
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

	varModelRegistryUpdateReq := _ModelRegistryUpdateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelRegistryUpdateReq)

	if err != nil {
		return err
	}

	*o = ModelRegistryUpdateReq(varModelRegistryUpdateReq)

	return err
}

type NullableModelRegistryUpdateReq struct {
	value *ModelRegistryUpdateReq
	isSet bool
}

func (v NullableModelRegistryUpdateReq) Get() *ModelRegistryUpdateReq {
	return v.value
}

func (v *NullableModelRegistryUpdateReq) Set(val *ModelRegistryUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelRegistryUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelRegistryUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelRegistryUpdateReq(val *ModelRegistryUpdateReq) *NullableModelRegistryUpdateReq {
	return &NullableModelRegistryUpdateReq{value: val, isSet: true}
}

func (v NullableModelRegistryUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelRegistryUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


