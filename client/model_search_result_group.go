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
)

// checks if the SearchResultGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchResultGroup{}

// SearchResultGroup struct for SearchResultGroup
type SearchResultGroup struct {
	Count *int32 `json:"count,omitempty"`
	Name *string `json:"name,omitempty"`
	Severity *string `json:"severity,omitempty"`
}

// NewSearchResultGroup instantiates a new SearchResultGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchResultGroup() *SearchResultGroup {
	this := SearchResultGroup{}
	return &this
}

// NewSearchResultGroupWithDefaults instantiates a new SearchResultGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchResultGroupWithDefaults() *SearchResultGroup {
	this := SearchResultGroup{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SearchResultGroup) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultGroup) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SearchResultGroup) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *SearchResultGroup) SetCount(v int32) {
	o.Count = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SearchResultGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SearchResultGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SearchResultGroup) SetName(v string) {
	o.Name = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *SearchResultGroup) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultGroup) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *SearchResultGroup) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *SearchResultGroup) SetSeverity(v string) {
	o.Severity = &v
}

func (o SearchResultGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchResultGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	return toSerialize, nil
}

type NullableSearchResultGroup struct {
	value *SearchResultGroup
	isSet bool
}

func (v NullableSearchResultGroup) Get() *SearchResultGroup {
	return v.value
}

func (v *NullableSearchResultGroup) Set(val *SearchResultGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResultGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResultGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResultGroup(val *SearchResultGroup) *NullableSearchResultGroup {
	return &NullableSearchResultGroup{value: val, isSet: true}
}

func (v NullableSearchResultGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchResultGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


