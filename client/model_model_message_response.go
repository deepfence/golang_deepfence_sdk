/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.4.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ModelMessageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelMessageResponse{}

// ModelMessageResponse struct for ModelMessageResponse
type ModelMessageResponse struct {
	Message string `json:"message"`
}

type _ModelMessageResponse ModelMessageResponse

// NewModelMessageResponse instantiates a new ModelMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelMessageResponse(message string) *ModelMessageResponse {
	this := ModelMessageResponse{}
	this.Message = message
	return &this
}

// NewModelMessageResponseWithDefaults instantiates a new ModelMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelMessageResponseWithDefaults() *ModelMessageResponse {
	this := ModelMessageResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *ModelMessageResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ModelMessageResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ModelMessageResponse) SetMessage(v string) {
	o.Message = v
}

func (o ModelMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelMessageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *ModelMessageResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
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

	varModelMessageResponse := _ModelMessageResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelMessageResponse)

	if err != nil {
		return err
	}

	*o = ModelMessageResponse(varModelMessageResponse)

	return err
}

type NullableModelMessageResponse struct {
	value *ModelMessageResponse
	isSet bool
}

func (v NullableModelMessageResponse) Get() *ModelMessageResponse {
	return v.value
}

func (v *NullableModelMessageResponse) Set(val *ModelMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelMessageResponse(val *ModelMessageResponse) *NullableModelMessageResponse {
	return &NullableModelMessageResponse{value: val, isSet: true}
}

func (v NullableModelMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


