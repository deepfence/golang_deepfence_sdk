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

// checks if the ModelSecretScanResultRules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSecretScanResultRules{}

// ModelSecretScanResultRules struct for ModelSecretScanResultRules
type ModelSecretScanResultRules struct {
	Rules []string `json:"rules"`
}

type _ModelSecretScanResultRules ModelSecretScanResultRules

// NewModelSecretScanResultRules instantiates a new ModelSecretScanResultRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSecretScanResultRules(rules []string) *ModelSecretScanResultRules {
	this := ModelSecretScanResultRules{}
	this.Rules = rules
	return &this
}

// NewModelSecretScanResultRulesWithDefaults instantiates a new ModelSecretScanResultRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSecretScanResultRulesWithDefaults() *ModelSecretScanResultRules {
	this := ModelSecretScanResultRules{}
	return &this
}

// GetRules returns the Rules field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ModelSecretScanResultRules) GetRules() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSecretScanResultRules) GetRulesOk() ([]string, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *ModelSecretScanResultRules) SetRules(v []string) {
	o.Rules = v
}

func (o ModelSecretScanResultRules) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSecretScanResultRules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

func (o *ModelSecretScanResultRules) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rules",
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

	varModelSecretScanResultRules := _ModelSecretScanResultRules{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelSecretScanResultRules)

	if err != nil {
		return err
	}

	*o = ModelSecretScanResultRules(varModelSecretScanResultRules)

	return err
}

type NullableModelSecretScanResultRules struct {
	value *ModelSecretScanResultRules
	isSet bool
}

func (v NullableModelSecretScanResultRules) Get() *ModelSecretScanResultRules {
	return v.value
}

func (v *NullableModelSecretScanResultRules) Set(val *ModelSecretScanResultRules) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSecretScanResultRules) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSecretScanResultRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSecretScanResultRules(val *ModelSecretScanResultRules) *NullableModelSecretScanResultRules {
	return &NullableModelSecretScanResultRules{value: val, isSet: true}
}

func (v NullableModelSecretScanResultRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSecretScanResultRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


