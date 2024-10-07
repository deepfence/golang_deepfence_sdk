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
)

// checks if the ModelRegistryCountResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelRegistryCountResp{}

// ModelRegistryCountResp struct for ModelRegistryCountResp
type ModelRegistryCountResp struct {
	Count *int32 `json:"count,omitempty"`
}

// NewModelRegistryCountResp instantiates a new ModelRegistryCountResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelRegistryCountResp() *ModelRegistryCountResp {
	this := ModelRegistryCountResp{}
	return &this
}

// NewModelRegistryCountRespWithDefaults instantiates a new ModelRegistryCountResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelRegistryCountRespWithDefaults() *ModelRegistryCountResp {
	this := ModelRegistryCountResp{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ModelRegistryCountResp) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRegistryCountResp) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ModelRegistryCountResp) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ModelRegistryCountResp) SetCount(v int32) {
	o.Count = &v
}

func (o ModelRegistryCountResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelRegistryCountResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableModelRegistryCountResp struct {
	value *ModelRegistryCountResp
	isSet bool
}

func (v NullableModelRegistryCountResp) Get() *ModelRegistryCountResp {
	return v.value
}

func (v *NullableModelRegistryCountResp) Set(val *ModelRegistryCountResp) {
	v.value = val
	v.isSet = true
}

func (v NullableModelRegistryCountResp) IsSet() bool {
	return v.isSet
}

func (v *NullableModelRegistryCountResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelRegistryCountResp(val *ModelRegistryCountResp) *NullableModelRegistryCountResp {
	return &NullableModelRegistryCountResp{value: val, isSet: true}
}

func (v NullableModelRegistryCountResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelRegistryCountResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


