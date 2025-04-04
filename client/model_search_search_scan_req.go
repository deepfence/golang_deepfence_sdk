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
	"bytes"
	"fmt"
)

// checks if the SearchSearchScanReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchSearchScanReq{}

// SearchSearchScanReq struct for SearchSearchScanReq
type SearchSearchScanReq struct {
	NodeFilters SearchSearchFilter `json:"node_filters"`
	RelatedNodeFilter *SearchChainedSearchFilter `json:"related_node_filter,omitempty"`
	ScanFilters SearchSearchFilter `json:"scan_filters"`
	Window ModelFetchWindow `json:"window"`
}

type _SearchSearchScanReq SearchSearchScanReq

// NewSearchSearchScanReq instantiates a new SearchSearchScanReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchSearchScanReq(nodeFilters SearchSearchFilter, scanFilters SearchSearchFilter, window ModelFetchWindow) *SearchSearchScanReq {
	this := SearchSearchScanReq{}
	this.NodeFilters = nodeFilters
	this.ScanFilters = scanFilters
	this.Window = window
	return &this
}

// NewSearchSearchScanReqWithDefaults instantiates a new SearchSearchScanReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchSearchScanReqWithDefaults() *SearchSearchScanReq {
	this := SearchSearchScanReq{}
	return &this
}

// GetNodeFilters returns the NodeFilters field value
func (o *SearchSearchScanReq) GetNodeFilters() SearchSearchFilter {
	if o == nil {
		var ret SearchSearchFilter
		return ret
	}

	return o.NodeFilters
}

// GetNodeFiltersOk returns a tuple with the NodeFilters field value
// and a boolean to check if the value has been set.
func (o *SearchSearchScanReq) GetNodeFiltersOk() (*SearchSearchFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeFilters, true
}

// SetNodeFilters sets field value
func (o *SearchSearchScanReq) SetNodeFilters(v SearchSearchFilter) {
	o.NodeFilters = v
}

// GetRelatedNodeFilter returns the RelatedNodeFilter field value if set, zero value otherwise.
func (o *SearchSearchScanReq) GetRelatedNodeFilter() SearchChainedSearchFilter {
	if o == nil || IsNil(o.RelatedNodeFilter) {
		var ret SearchChainedSearchFilter
		return ret
	}
	return *o.RelatedNodeFilter
}

// GetRelatedNodeFilterOk returns a tuple with the RelatedNodeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSearchScanReq) GetRelatedNodeFilterOk() (*SearchChainedSearchFilter, bool) {
	if o == nil || IsNil(o.RelatedNodeFilter) {
		return nil, false
	}
	return o.RelatedNodeFilter, true
}

// HasRelatedNodeFilter returns a boolean if a field has been set.
func (o *SearchSearchScanReq) HasRelatedNodeFilter() bool {
	if o != nil && !IsNil(o.RelatedNodeFilter) {
		return true
	}

	return false
}

// SetRelatedNodeFilter gets a reference to the given SearchChainedSearchFilter and assigns it to the RelatedNodeFilter field.
func (o *SearchSearchScanReq) SetRelatedNodeFilter(v SearchChainedSearchFilter) {
	o.RelatedNodeFilter = &v
}

// GetScanFilters returns the ScanFilters field value
func (o *SearchSearchScanReq) GetScanFilters() SearchSearchFilter {
	if o == nil {
		var ret SearchSearchFilter
		return ret
	}

	return o.ScanFilters
}

// GetScanFiltersOk returns a tuple with the ScanFilters field value
// and a boolean to check if the value has been set.
func (o *SearchSearchScanReq) GetScanFiltersOk() (*SearchSearchFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScanFilters, true
}

// SetScanFilters sets field value
func (o *SearchSearchScanReq) SetScanFilters(v SearchSearchFilter) {
	o.ScanFilters = v
}

// GetWindow returns the Window field value
func (o *SearchSearchScanReq) GetWindow() ModelFetchWindow {
	if o == nil {
		var ret ModelFetchWindow
		return ret
	}

	return o.Window
}

// GetWindowOk returns a tuple with the Window field value
// and a boolean to check if the value has been set.
func (o *SearchSearchScanReq) GetWindowOk() (*ModelFetchWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Window, true
}

// SetWindow sets field value
func (o *SearchSearchScanReq) SetWindow(v ModelFetchWindow) {
	o.Window = v
}

func (o SearchSearchScanReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchSearchScanReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["node_filters"] = o.NodeFilters
	if !IsNil(o.RelatedNodeFilter) {
		toSerialize["related_node_filter"] = o.RelatedNodeFilter
	}
	toSerialize["scan_filters"] = o.ScanFilters
	toSerialize["window"] = o.Window
	return toSerialize, nil
}

func (o *SearchSearchScanReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"node_filters",
		"scan_filters",
		"window",
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

	varSearchSearchScanReq := _SearchSearchScanReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchSearchScanReq)

	if err != nil {
		return err
	}

	*o = SearchSearchScanReq(varSearchSearchScanReq)

	return err
}

type NullableSearchSearchScanReq struct {
	value *SearchSearchScanReq
	isSet bool
}

func (v NullableSearchSearchScanReq) Get() *SearchSearchScanReq {
	return v.value
}

func (v *NullableSearchSearchScanReq) Set(val *SearchSearchScanReq) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSearchScanReq) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSearchScanReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSearchScanReq(val *SearchSearchScanReq) *NullableSearchSearchScanReq {
	return &NullableSearchSearchScanReq{value: val, isSet: true}
}

func (v NullableSearchSearchScanReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSearchScanReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


