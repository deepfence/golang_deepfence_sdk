/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.2.1
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ModelCloudNodeAccountRegisterResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCloudNodeAccountRegisterResp{}

// ModelCloudNodeAccountRegisterResp struct for ModelCloudNodeAccountRegisterResp
type ModelCloudNodeAccountRegisterResp struct {
	Data *ModelCloudNodeAccountRegisterRespData `json:"data,omitempty"`
}

// NewModelCloudNodeAccountRegisterResp instantiates a new ModelCloudNodeAccountRegisterResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCloudNodeAccountRegisterResp() *ModelCloudNodeAccountRegisterResp {
	this := ModelCloudNodeAccountRegisterResp{}
	return &this
}

// NewModelCloudNodeAccountRegisterRespWithDefaults instantiates a new ModelCloudNodeAccountRegisterResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCloudNodeAccountRegisterRespWithDefaults() *ModelCloudNodeAccountRegisterResp {
	this := ModelCloudNodeAccountRegisterResp{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModelCloudNodeAccountRegisterResp) GetData() ModelCloudNodeAccountRegisterRespData {
	if o == nil || IsNil(o.Data) {
		var ret ModelCloudNodeAccountRegisterRespData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeAccountRegisterResp) GetDataOk() (*ModelCloudNodeAccountRegisterRespData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModelCloudNodeAccountRegisterResp) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ModelCloudNodeAccountRegisterRespData and assigns it to the Data field.
func (o *ModelCloudNodeAccountRegisterResp) SetData(v ModelCloudNodeAccountRegisterRespData) {
	o.Data = &v
}

func (o ModelCloudNodeAccountRegisterResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCloudNodeAccountRegisterResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableModelCloudNodeAccountRegisterResp struct {
	value *ModelCloudNodeAccountRegisterResp
	isSet bool
}

func (v NullableModelCloudNodeAccountRegisterResp) Get() *ModelCloudNodeAccountRegisterResp {
	return v.value
}

func (v *NullableModelCloudNodeAccountRegisterResp) Set(val *ModelCloudNodeAccountRegisterResp) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCloudNodeAccountRegisterResp) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCloudNodeAccountRegisterResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCloudNodeAccountRegisterResp(val *ModelCloudNodeAccountRegisterResp) *NullableModelCloudNodeAccountRegisterResp {
	return &NullableModelCloudNodeAccountRegisterResp{value: val, isSet: true}
}

func (v NullableModelCloudNodeAccountRegisterResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCloudNodeAccountRegisterResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


