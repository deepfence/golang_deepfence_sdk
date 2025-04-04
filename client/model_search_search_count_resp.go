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

// checks if the SearchSearchCountResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchSearchCountResp{}

// SearchSearchCountResp struct for SearchSearchCountResp
type SearchSearchCountResp struct {
	Categories map[string]int32 `json:"categories"`
	Count int32 `json:"count"`
}

type _SearchSearchCountResp SearchSearchCountResp

// NewSearchSearchCountResp instantiates a new SearchSearchCountResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchSearchCountResp(categories map[string]int32, count int32) *SearchSearchCountResp {
	this := SearchSearchCountResp{}
	this.Categories = categories
	this.Count = count
	return &this
}

// NewSearchSearchCountRespWithDefaults instantiates a new SearchSearchCountResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchSearchCountRespWithDefaults() *SearchSearchCountResp {
	this := SearchSearchCountResp{}
	return &this
}

// GetCategories returns the Categories field value
// If the value is explicit nil, the zero value for map[string]int32 will be returned
func (o *SearchSearchCountResp) GetCategories() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchSearchCountResp) GetCategoriesOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return &o.Categories, true
}

// SetCategories sets field value
func (o *SearchSearchCountResp) SetCategories(v map[string]int32) {
	o.Categories = v
}

// GetCount returns the Count field value
func (o *SearchSearchCountResp) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *SearchSearchCountResp) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *SearchSearchCountResp) SetCount(v int32) {
	o.Count = v
}

func (o SearchSearchCountResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchSearchCountResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	toSerialize["count"] = o.Count
	return toSerialize, nil
}

func (o *SearchSearchCountResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"categories",
		"count",
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

	varSearchSearchCountResp := _SearchSearchCountResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchSearchCountResp)

	if err != nil {
		return err
	}

	*o = SearchSearchCountResp(varSearchSearchCountResp)

	return err
}

type NullableSearchSearchCountResp struct {
	value *SearchSearchCountResp
	isSet bool
}

func (v NullableSearchSearchCountResp) Get() *SearchSearchCountResp {
	return v.value
}

func (v *NullableSearchSearchCountResp) Set(val *SearchSearchCountResp) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSearchCountResp) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSearchCountResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSearchCountResp(val *SearchSearchCountResp) *NullableSearchSearchCountResp {
	return &NullableSearchSearchCountResp{value: val, isSet: true}
}

func (v NullableSearchSearchCountResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSearchCountResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


