/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.5.6
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ModelFetchWindow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelFetchWindow{}

// ModelFetchWindow struct for ModelFetchWindow
type ModelFetchWindow struct {
	Offset int32 `json:"offset"`
	Size int32 `json:"size"`
}

type _ModelFetchWindow ModelFetchWindow

// NewModelFetchWindow instantiates a new ModelFetchWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelFetchWindow(offset int32, size int32) *ModelFetchWindow {
	this := ModelFetchWindow{}
	this.Offset = offset
	this.Size = size
	return &this
}

// NewModelFetchWindowWithDefaults instantiates a new ModelFetchWindow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelFetchWindowWithDefaults() *ModelFetchWindow {
	this := ModelFetchWindow{}
	return &this
}

// GetOffset returns the Offset field value
func (o *ModelFetchWindow) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *ModelFetchWindow) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *ModelFetchWindow) SetOffset(v int32) {
	o.Offset = v
}

// GetSize returns the Size field value
func (o *ModelFetchWindow) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ModelFetchWindow) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ModelFetchWindow) SetSize(v int32) {
	o.Size = v
}

func (o ModelFetchWindow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelFetchWindow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["offset"] = o.Offset
	toSerialize["size"] = o.Size
	return toSerialize, nil
}

func (o *ModelFetchWindow) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"offset",
		"size",
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

	varModelFetchWindow := _ModelFetchWindow{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelFetchWindow)

	if err != nil {
		return err
	}

	*o = ModelFetchWindow(varModelFetchWindow)

	return err
}

type NullableModelFetchWindow struct {
	value *ModelFetchWindow
	isSet bool
}

func (v NullableModelFetchWindow) Get() *ModelFetchWindow {
	return v.value
}

func (v *NullableModelFetchWindow) Set(val *ModelFetchWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableModelFetchWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableModelFetchWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelFetchWindow(val *ModelFetchWindow) *NullableModelFetchWindow {
	return &NullableModelFetchWindow{value: val, isSet: true}
}

func (v NullableModelFetchWindow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelFetchWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


