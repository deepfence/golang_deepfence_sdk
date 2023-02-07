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

// checks if the ModelFieldsFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelFieldsFilter{}

// ModelFieldsFilter struct for ModelFieldsFilter
type ModelFieldsFilter struct {
	FieldsValues []ModelKeyValue `json:"fields_values"`
}

// NewModelFieldsFilter instantiates a new ModelFieldsFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelFieldsFilter(fieldsValues []ModelKeyValue) *ModelFieldsFilter {
	this := ModelFieldsFilter{}
	this.FieldsValues = fieldsValues
	return &this
}

// NewModelFieldsFilterWithDefaults instantiates a new ModelFieldsFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelFieldsFilterWithDefaults() *ModelFieldsFilter {
	this := ModelFieldsFilter{}
	return &this
}

// GetFieldsValues returns the FieldsValues field value
// If the value is explicit nil, the zero value for []ModelKeyValue will be returned
func (o *ModelFieldsFilter) GetFieldsValues() []ModelKeyValue {
	if o == nil {
		var ret []ModelKeyValue
		return ret
	}

	return o.FieldsValues
}

// GetFieldsValuesOk returns a tuple with the FieldsValues field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelFieldsFilter) GetFieldsValuesOk() ([]ModelKeyValue, bool) {
	if o == nil || isNil(o.FieldsValues) {
		return nil, false
	}
	return o.FieldsValues, true
}

// SetFieldsValues sets field value
func (o *ModelFieldsFilter) SetFieldsValues(v []ModelKeyValue) {
	o.FieldsValues = v
}

func (o ModelFieldsFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelFieldsFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FieldsValues != nil {
		toSerialize["fields_values"] = o.FieldsValues
	}
	return toSerialize, nil
}

type NullableModelFieldsFilter struct {
	value *ModelFieldsFilter
	isSet bool
}

func (v NullableModelFieldsFilter) Get() *ModelFieldsFilter {
	return v.value
}

func (v *NullableModelFieldsFilter) Set(val *ModelFieldsFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableModelFieldsFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableModelFieldsFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelFieldsFilter(val *ModelFieldsFilter) *NullableModelFieldsFilter {
	return &NullableModelFieldsFilter{value: val, isSet: true}
}

func (v NullableModelFieldsFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelFieldsFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


