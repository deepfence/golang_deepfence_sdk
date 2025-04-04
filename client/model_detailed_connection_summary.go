/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.5.5
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the DetailedConnectionSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetailedConnectionSummary{}

// DetailedConnectionSummary struct for DetailedConnectionSummary
type DetailedConnectionSummary struct {
	Source *string `json:"source,omitempty"`
	Target *string `json:"target,omitempty"`
}

// NewDetailedConnectionSummary instantiates a new DetailedConnectionSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailedConnectionSummary() *DetailedConnectionSummary {
	this := DetailedConnectionSummary{}
	return &this
}

// NewDetailedConnectionSummaryWithDefaults instantiates a new DetailedConnectionSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailedConnectionSummaryWithDefaults() *DetailedConnectionSummary {
	this := DetailedConnectionSummary{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *DetailedConnectionSummary) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedConnectionSummary) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *DetailedConnectionSummary) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *DetailedConnectionSummary) SetSource(v string) {
	o.Source = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *DetailedConnectionSummary) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedConnectionSummary) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *DetailedConnectionSummary) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *DetailedConnectionSummary) SetTarget(v string) {
	o.Target = &v
}

func (o DetailedConnectionSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetailedConnectionSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	return toSerialize, nil
}

type NullableDetailedConnectionSummary struct {
	value *DetailedConnectionSummary
	isSet bool
}

func (v NullableDetailedConnectionSummary) Get() *DetailedConnectionSummary {
	return v.value
}

func (v *NullableDetailedConnectionSummary) Set(val *DetailedConnectionSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailedConnectionSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailedConnectionSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailedConnectionSummary(val *DetailedConnectionSummary) *NullableDetailedConnectionSummary {
	return &NullableDetailedConnectionSummary{value: val, isSet: true}
}

func (v NullableDetailedConnectionSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailedConnectionSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


