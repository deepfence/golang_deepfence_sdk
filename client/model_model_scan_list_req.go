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
	"bytes"
	"fmt"
)

// checks if the ModelScanListReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelScanListReq{}

// ModelScanListReq struct for ModelScanListReq
type ModelScanListReq struct {
	FieldsFilter ReportersFieldsFilters `json:"fields_filter"`
	NodeIds []ModelNodeIdentifier `json:"node_ids"`
	Window ModelFetchWindow `json:"window"`
}

type _ModelScanListReq ModelScanListReq

// NewModelScanListReq instantiates a new ModelScanListReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelScanListReq(fieldsFilter ReportersFieldsFilters, nodeIds []ModelNodeIdentifier, window ModelFetchWindow) *ModelScanListReq {
	this := ModelScanListReq{}
	this.FieldsFilter = fieldsFilter
	this.NodeIds = nodeIds
	this.Window = window
	return &this
}

// NewModelScanListReqWithDefaults instantiates a new ModelScanListReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelScanListReqWithDefaults() *ModelScanListReq {
	this := ModelScanListReq{}
	return &this
}

// GetFieldsFilter returns the FieldsFilter field value
func (o *ModelScanListReq) GetFieldsFilter() ReportersFieldsFilters {
	if o == nil {
		var ret ReportersFieldsFilters
		return ret
	}

	return o.FieldsFilter
}

// GetFieldsFilterOk returns a tuple with the FieldsFilter field value
// and a boolean to check if the value has been set.
func (o *ModelScanListReq) GetFieldsFilterOk() (*ReportersFieldsFilters, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldsFilter, true
}

// SetFieldsFilter sets field value
func (o *ModelScanListReq) SetFieldsFilter(v ReportersFieldsFilters) {
	o.FieldsFilter = v
}

// GetNodeIds returns the NodeIds field value
// If the value is explicit nil, the zero value for []ModelNodeIdentifier will be returned
func (o *ModelScanListReq) GetNodeIds() []ModelNodeIdentifier {
	if o == nil {
		var ret []ModelNodeIdentifier
		return ret
	}

	return o.NodeIds
}

// GetNodeIdsOk returns a tuple with the NodeIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelScanListReq) GetNodeIdsOk() ([]ModelNodeIdentifier, bool) {
	if o == nil || IsNil(o.NodeIds) {
		return nil, false
	}
	return o.NodeIds, true
}

// SetNodeIds sets field value
func (o *ModelScanListReq) SetNodeIds(v []ModelNodeIdentifier) {
	o.NodeIds = v
}

// GetWindow returns the Window field value
func (o *ModelScanListReq) GetWindow() ModelFetchWindow {
	if o == nil {
		var ret ModelFetchWindow
		return ret
	}

	return o.Window
}

// GetWindowOk returns a tuple with the Window field value
// and a boolean to check if the value has been set.
func (o *ModelScanListReq) GetWindowOk() (*ModelFetchWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Window, true
}

// SetWindow sets field value
func (o *ModelScanListReq) SetWindow(v ModelFetchWindow) {
	o.Window = v
}

func (o ModelScanListReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelScanListReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields_filter"] = o.FieldsFilter
	if o.NodeIds != nil {
		toSerialize["node_ids"] = o.NodeIds
	}
	toSerialize["window"] = o.Window
	return toSerialize, nil
}

func (o *ModelScanListReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fields_filter",
		"node_ids",
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

	varModelScanListReq := _ModelScanListReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelScanListReq)

	if err != nil {
		return err
	}

	*o = ModelScanListReq(varModelScanListReq)

	return err
}

type NullableModelScanListReq struct {
	value *ModelScanListReq
	isSet bool
}

func (v NullableModelScanListReq) Get() *ModelScanListReq {
	return v.value
}

func (v *NullableModelScanListReq) Set(val *ModelScanListReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelScanListReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelScanListReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelScanListReq(val *ModelScanListReq) *NullableModelScanListReq {
	return &NullableModelScanListReq{value: val, isSet: true}
}

func (v NullableModelScanListReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelScanListReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


