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

// checks if the ReportersOrderSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportersOrderSpec{}

// ReportersOrderSpec struct for ReportersOrderSpec
type ReportersOrderSpec struct {
	Descending bool `json:"descending"`
	FieldName string `json:"field_name"`
}

// NewReportersOrderSpec instantiates a new ReportersOrderSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportersOrderSpec(descending bool, fieldName string) *ReportersOrderSpec {
	this := ReportersOrderSpec{}
	this.Descending = descending
	this.FieldName = fieldName
	return &this
}

// NewReportersOrderSpecWithDefaults instantiates a new ReportersOrderSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportersOrderSpecWithDefaults() *ReportersOrderSpec {
	this := ReportersOrderSpec{}
	return &this
}

// GetDescending returns the Descending field value
func (o *ReportersOrderSpec) GetDescending() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Descending
}

// GetDescendingOk returns a tuple with the Descending field value
// and a boolean to check if the value has been set.
func (o *ReportersOrderSpec) GetDescendingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Descending, true
}

// SetDescending sets field value
func (o *ReportersOrderSpec) SetDescending(v bool) {
	o.Descending = v
}

// GetFieldName returns the FieldName field value
func (o *ReportersOrderSpec) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *ReportersOrderSpec) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *ReportersOrderSpec) SetFieldName(v string) {
	o.FieldName = v
}

func (o ReportersOrderSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportersOrderSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["descending"] = o.Descending
	toSerialize["field_name"] = o.FieldName
	return toSerialize, nil
}

type NullableReportersOrderSpec struct {
	value *ReportersOrderSpec
	isSet bool
}

func (v NullableReportersOrderSpec) Get() *ReportersOrderSpec {
	return v.value
}

func (v *NullableReportersOrderSpec) Set(val *ReportersOrderSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableReportersOrderSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableReportersOrderSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportersOrderSpec(val *ReportersOrderSpec) *NullableReportersOrderSpec {
	return &NullableReportersOrderSpec{value: val, isSet: true}
}

func (v NullableReportersOrderSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportersOrderSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

