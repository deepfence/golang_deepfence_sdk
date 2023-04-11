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

// checks if the ModelSettingUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSettingUpdateRequest{}

// ModelSettingUpdateRequest struct for ModelSettingUpdateRequest
type ModelSettingUpdateRequest struct {
	Description string `json:"description"`
	Id int32 `json:"id"`
	Key string `json:"key"`
	Label string `json:"label"`
	Value interface{} `json:"value"`
}

// NewModelSettingUpdateRequest instantiates a new ModelSettingUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSettingUpdateRequest(description string, id int32, key string, label string, value interface{}) *ModelSettingUpdateRequest {
	this := ModelSettingUpdateRequest{}
	this.Description = description
	this.Id = id
	this.Key = key
	this.Label = label
	this.Value = value
	return &this
}

// NewModelSettingUpdateRequestWithDefaults instantiates a new ModelSettingUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSettingUpdateRequestWithDefaults() *ModelSettingUpdateRequest {
	this := ModelSettingUpdateRequest{}
	return &this
}

// GetDescription returns the Description field value
func (o *ModelSettingUpdateRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ModelSettingUpdateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ModelSettingUpdateRequest) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *ModelSettingUpdateRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModelSettingUpdateRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModelSettingUpdateRequest) SetId(v int32) {
	o.Id = v
}

// GetKey returns the Key field value
func (o *ModelSettingUpdateRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ModelSettingUpdateRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ModelSettingUpdateRequest) SetKey(v string) {
	o.Key = v
}

// GetLabel returns the Label field value
func (o *ModelSettingUpdateRequest) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ModelSettingUpdateRequest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ModelSettingUpdateRequest) SetLabel(v string) {
	o.Label = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ModelSettingUpdateRequest) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSettingUpdateRequest) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ModelSettingUpdateRequest) SetValue(v interface{}) {
	o.Value = v
}

func (o ModelSettingUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSettingUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	toSerialize["key"] = o.Key
	toSerialize["label"] = o.Label
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableModelSettingUpdateRequest struct {
	value *ModelSettingUpdateRequest
	isSet bool
}

func (v NullableModelSettingUpdateRequest) Get() *ModelSettingUpdateRequest {
	return v.value
}

func (v *NullableModelSettingUpdateRequest) Set(val *ModelSettingUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSettingUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSettingUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSettingUpdateRequest(val *ModelSettingUpdateRequest) *NullableModelSettingUpdateRequest {
	return &NullableModelSettingUpdateRequest{value: val, isSet: true}
}

func (v NullableModelSettingUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSettingUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


