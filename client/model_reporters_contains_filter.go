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

// checks if the ReportersContainsFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportersContainsFilter{}

// ReportersContainsFilter struct for ReportersContainsFilter
type ReportersContainsFilter struct {
	FilterIn map[string][]interface{} `json:"filter_in"`
}

type _ReportersContainsFilter ReportersContainsFilter

// NewReportersContainsFilter instantiates a new ReportersContainsFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportersContainsFilter(filterIn map[string][]interface{}) *ReportersContainsFilter {
	this := ReportersContainsFilter{}
	this.FilterIn = filterIn
	return &this
}

// NewReportersContainsFilterWithDefaults instantiates a new ReportersContainsFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportersContainsFilterWithDefaults() *ReportersContainsFilter {
	this := ReportersContainsFilter{}
	return &this
}

// GetFilterIn returns the FilterIn field value
// If the value is explicit nil, the zero value for map[string][]interface{} will be returned
func (o *ReportersContainsFilter) GetFilterIn() map[string][]interface{} {
	if o == nil {
		var ret map[string][]interface{}
		return ret
	}

	return o.FilterIn
}

// GetFilterInOk returns a tuple with the FilterIn field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportersContainsFilter) GetFilterInOk() (*map[string][]interface{}, bool) {
	if o == nil || IsNil(o.FilterIn) {
		return nil, false
	}
	return &o.FilterIn, true
}

// SetFilterIn sets field value
func (o *ReportersContainsFilter) SetFilterIn(v map[string][]interface{}) {
	o.FilterIn = v
}

func (o ReportersContainsFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportersContainsFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FilterIn != nil {
		toSerialize["filter_in"] = o.FilterIn
	}
	return toSerialize, nil
}

func (o *ReportersContainsFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filter_in",
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

	varReportersContainsFilter := _ReportersContainsFilter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReportersContainsFilter)

	if err != nil {
		return err
	}

	*o = ReportersContainsFilter(varReportersContainsFilter)

	return err
}

type NullableReportersContainsFilter struct {
	value *ReportersContainsFilter
	isSet bool
}

func (v NullableReportersContainsFilter) Get() *ReportersContainsFilter {
	return v.value
}

func (v *NullableReportersContainsFilter) Set(val *ReportersContainsFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableReportersContainsFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableReportersContainsFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportersContainsFilter(val *ReportersContainsFilter) *NullableReportersContainsFilter {
	return &NullableReportersContainsFilter{value: val, isSet: true}
}

func (v NullableReportersContainsFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportersContainsFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


