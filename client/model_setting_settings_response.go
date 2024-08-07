/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v3.0.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SettingSettingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingSettingsResponse{}

// SettingSettingsResponse struct for SettingSettingsResponse
type SettingSettingsResponse struct {
	Description string `json:"description"`
	Id int32 `json:"id"`
	Key string `json:"key"`
	Label string `json:"label"`
	Value interface{} `json:"value"`
}

type _SettingSettingsResponse SettingSettingsResponse

// NewSettingSettingsResponse instantiates a new SettingSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingSettingsResponse(description string, id int32, key string, label string, value interface{}) *SettingSettingsResponse {
	this := SettingSettingsResponse{}
	this.Description = description
	this.Id = id
	this.Key = key
	this.Label = label
	this.Value = value
	return &this
}

// NewSettingSettingsResponseWithDefaults instantiates a new SettingSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingSettingsResponseWithDefaults() *SettingSettingsResponse {
	this := SettingSettingsResponse{}
	return &this
}

// GetDescription returns the Description field value
func (o *SettingSettingsResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SettingSettingsResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SettingSettingsResponse) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *SettingSettingsResponse) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SettingSettingsResponse) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SettingSettingsResponse) SetId(v int32) {
	o.Id = v
}

// GetKey returns the Key field value
func (o *SettingSettingsResponse) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SettingSettingsResponse) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *SettingSettingsResponse) SetKey(v string) {
	o.Key = v
}

// GetLabel returns the Label field value
func (o *SettingSettingsResponse) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *SettingSettingsResponse) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *SettingSettingsResponse) SetLabel(v string) {
	o.Label = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SettingSettingsResponse) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingSettingsResponse) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SettingSettingsResponse) SetValue(v interface{}) {
	o.Value = v
}

func (o SettingSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingSettingsResponse) ToMap() (map[string]interface{}, error) {
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

func (o *SettingSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"id",
		"key",
		"label",
		"value",
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

	varSettingSettingsResponse := _SettingSettingsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSettingSettingsResponse)

	if err != nil {
		return err
	}

	*o = SettingSettingsResponse(varSettingSettingsResponse)

	return err
}

type NullableSettingSettingsResponse struct {
	value *SettingSettingsResponse
	isSet bool
}

func (v NullableSettingSettingsResponse) Get() *SettingSettingsResponse {
	return v.value
}

func (v *NullableSettingSettingsResponse) Set(val *SettingSettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingSettingsResponse(val *SettingSettingsResponse) *NullableSettingSettingsResponse {
	return &NullableSettingSettingsResponse{value: val, isSet: true}
}

func (v NullableSettingSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


