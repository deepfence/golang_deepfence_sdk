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

// checks if the ControlsAgentBeat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControlsAgentBeat{}

// ControlsAgentBeat struct for ControlsAgentBeat
type ControlsAgentBeat struct {
	Beatrate int32 `json:"beatrate"`
}

// NewControlsAgentBeat instantiates a new ControlsAgentBeat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControlsAgentBeat(beatrate int32) *ControlsAgentBeat {
	this := ControlsAgentBeat{}
	this.Beatrate = beatrate
	return &this
}

// NewControlsAgentBeatWithDefaults instantiates a new ControlsAgentBeat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControlsAgentBeatWithDefaults() *ControlsAgentBeat {
	this := ControlsAgentBeat{}
	return &this
}

// GetBeatrate returns the Beatrate field value
func (o *ControlsAgentBeat) GetBeatrate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Beatrate
}

// GetBeatrateOk returns a tuple with the Beatrate field value
// and a boolean to check if the value has been set.
func (o *ControlsAgentBeat) GetBeatrateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Beatrate, true
}

// SetBeatrate sets field value
func (o *ControlsAgentBeat) SetBeatrate(v int32) {
	o.Beatrate = v
}

func (o ControlsAgentBeat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControlsAgentBeat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["beatrate"] = o.Beatrate
	return toSerialize, nil
}

type NullableControlsAgentBeat struct {
	value *ControlsAgentBeat
	isSet bool
}

func (v NullableControlsAgentBeat) Get() *ControlsAgentBeat {
	return v.value
}

func (v *NullableControlsAgentBeat) Set(val *ControlsAgentBeat) {
	v.value = val
	v.isSet = true
}

func (v NullableControlsAgentBeat) IsSet() bool {
	return v.isSet
}

func (v *NullableControlsAgentBeat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControlsAgentBeat(val *ControlsAgentBeat) *NullableControlsAgentBeat {
	return &NullableControlsAgentBeat{value: val, isSet: true}
}

func (v NullableControlsAgentBeat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControlsAgentBeat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


