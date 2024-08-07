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

// checks if the ModelFiltersReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelFiltersReq{}

// ModelFiltersReq struct for ModelFiltersReq
type ModelFiltersReq struct {
	Filters []string `json:"filters"`
	Having map[string]interface{} `json:"having,omitempty"`
}

type _ModelFiltersReq ModelFiltersReq

// NewModelFiltersReq instantiates a new ModelFiltersReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelFiltersReq(filters []string) *ModelFiltersReq {
	this := ModelFiltersReq{}
	this.Filters = filters
	return &this
}

// NewModelFiltersReqWithDefaults instantiates a new ModelFiltersReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelFiltersReqWithDefaults() *ModelFiltersReq {
	this := ModelFiltersReq{}
	return &this
}

// GetFilters returns the Filters field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ModelFiltersReq) GetFilters() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelFiltersReq) GetFiltersOk() ([]string, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// SetFilters sets field value
func (o *ModelFiltersReq) SetFilters(v []string) {
	o.Filters = v
}

// GetHaving returns the Having field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelFiltersReq) GetHaving() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Having
}

// GetHavingOk returns a tuple with the Having field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelFiltersReq) GetHavingOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Having) {
		return map[string]interface{}{}, false
	}
	return o.Having, true
}

// HasHaving returns a boolean if a field has been set.
func (o *ModelFiltersReq) HasHaving() bool {
	if o != nil && !IsNil(o.Having) {
		return true
	}

	return false
}

// SetHaving gets a reference to the given map[string]interface{} and assigns it to the Having field.
func (o *ModelFiltersReq) SetHaving(v map[string]interface{}) {
	o.Having = v
}

func (o ModelFiltersReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelFiltersReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.Having != nil {
		toSerialize["having"] = o.Having
	}
	return toSerialize, nil
}

func (o *ModelFiltersReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filters",
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

	varModelFiltersReq := _ModelFiltersReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelFiltersReq)

	if err != nil {
		return err
	}

	*o = ModelFiltersReq(varModelFiltersReq)

	return err
}

type NullableModelFiltersReq struct {
	value *ModelFiltersReq
	isSet bool
}

func (v NullableModelFiltersReq) Get() *ModelFiltersReq {
	return v.value
}

func (v *NullableModelFiltersReq) Set(val *ModelFiltersReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelFiltersReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelFiltersReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelFiltersReq(val *ModelFiltersReq) *NullableModelFiltersReq {
	return &NullableModelFiltersReq{value: val, isSet: true}
}

func (v NullableModelFiltersReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelFiltersReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


