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

// checks if the ControlsKubernetesScannerControlResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControlsKubernetesScannerControlResponse{}

// ControlsKubernetesScannerControlResponse struct for ControlsKubernetesScannerControlResponse
type ControlsKubernetesScannerControlResponse struct {
	Data *ControlsKubernetesScannerPendingScans `json:"data,omitempty"`
}

// NewControlsKubernetesScannerControlResponse instantiates a new ControlsKubernetesScannerControlResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControlsKubernetesScannerControlResponse() *ControlsKubernetesScannerControlResponse {
	this := ControlsKubernetesScannerControlResponse{}
	return &this
}

// NewControlsKubernetesScannerControlResponseWithDefaults instantiates a new ControlsKubernetesScannerControlResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControlsKubernetesScannerControlResponseWithDefaults() *ControlsKubernetesScannerControlResponse {
	this := ControlsKubernetesScannerControlResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ControlsKubernetesScannerControlResponse) GetData() ControlsKubernetesScannerPendingScans {
	if o == nil || isNil(o.Data) {
		var ret ControlsKubernetesScannerPendingScans
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsKubernetesScannerControlResponse) GetDataOk() (*ControlsKubernetesScannerPendingScans, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ControlsKubernetesScannerControlResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ControlsKubernetesScannerPendingScans and assigns it to the Data field.
func (o *ControlsKubernetesScannerControlResponse) SetData(v ControlsKubernetesScannerPendingScans) {
	o.Data = &v
}

func (o ControlsKubernetesScannerControlResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControlsKubernetesScannerControlResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableControlsKubernetesScannerControlResponse struct {
	value *ControlsKubernetesScannerControlResponse
	isSet bool
}

func (v NullableControlsKubernetesScannerControlResponse) Get() *ControlsKubernetesScannerControlResponse {
	return v.value
}

func (v *NullableControlsKubernetesScannerControlResponse) Set(val *ControlsKubernetesScannerControlResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableControlsKubernetesScannerControlResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableControlsKubernetesScannerControlResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControlsKubernetesScannerControlResponse(val *ControlsKubernetesScannerControlResponse) *NullableControlsKubernetesScannerControlResponse {
	return &NullableControlsKubernetesScannerControlResponse{value: val, isSet: true}
}

func (v NullableControlsKubernetesScannerControlResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControlsKubernetesScannerControlResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


