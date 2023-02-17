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

// checks if the SearchSearchNodeReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchSearchNodeReq{}

// SearchSearchNodeReq struct for SearchSearchNodeReq
type SearchSearchNodeReq struct {
	NodeFilter SearchSearchFilter `json:"node_filter"`
	Window ModelFetchWindow `json:"window"`
}

// NewSearchSearchNodeReq instantiates a new SearchSearchNodeReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchSearchNodeReq(nodeFilter SearchSearchFilter, window ModelFetchWindow) *SearchSearchNodeReq {
	this := SearchSearchNodeReq{}
	this.NodeFilter = nodeFilter
	this.Window = window
	return &this
}

// NewSearchSearchNodeReqWithDefaults instantiates a new SearchSearchNodeReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchSearchNodeReqWithDefaults() *SearchSearchNodeReq {
	this := SearchSearchNodeReq{}
	return &this
}

// GetNodeFilter returns the NodeFilter field value
func (o *SearchSearchNodeReq) GetNodeFilter() SearchSearchFilter {
	if o == nil {
		var ret SearchSearchFilter
		return ret
	}

	return o.NodeFilter
}

// GetNodeFilterOk returns a tuple with the NodeFilter field value
// and a boolean to check if the value has been set.
func (o *SearchSearchNodeReq) GetNodeFilterOk() (*SearchSearchFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeFilter, true
}

// SetNodeFilter sets field value
func (o *SearchSearchNodeReq) SetNodeFilter(v SearchSearchFilter) {
	o.NodeFilter = v
}

// GetWindow returns the Window field value
func (o *SearchSearchNodeReq) GetWindow() ModelFetchWindow {
	if o == nil {
		var ret ModelFetchWindow
		return ret
	}

	return o.Window
}

// GetWindowOk returns a tuple with the Window field value
// and a boolean to check if the value has been set.
func (o *SearchSearchNodeReq) GetWindowOk() (*ModelFetchWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Window, true
}

// SetWindow sets field value
func (o *SearchSearchNodeReq) SetWindow(v ModelFetchWindow) {
	o.Window = v
}

func (o SearchSearchNodeReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchSearchNodeReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["node_filter"] = o.NodeFilter
	toSerialize["window"] = o.Window
	return toSerialize, nil
}

type NullableSearchSearchNodeReq struct {
	value *SearchSearchNodeReq
	isSet bool
}

func (v NullableSearchSearchNodeReq) Get() *SearchSearchNodeReq {
	return v.value
}

func (v *NullableSearchSearchNodeReq) Set(val *SearchSearchNodeReq) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSearchNodeReq) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSearchNodeReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSearchNodeReq(val *SearchSearchNodeReq) *NullableSearchSearchNodeReq {
	return &NullableSearchSearchNodeReq{value: val, isSet: true}
}

func (v NullableSearchSearchNodeReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSearchNodeReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


