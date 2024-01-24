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
	"bytes"
	"fmt"
)

// checks if the ModelGenerativeAiIntegrationSecretRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelGenerativeAiIntegrationSecretRequest{}

// ModelGenerativeAiIntegrationSecretRequest struct for ModelGenerativeAiIntegrationSecretRequest
type ModelGenerativeAiIntegrationSecretRequest struct {
	IntegrationId *int32 `json:"integration_id,omitempty"`
	Name string `json:"name"`
	QueryType string `json:"query_type"`
}

type _ModelGenerativeAiIntegrationSecretRequest ModelGenerativeAiIntegrationSecretRequest

// NewModelGenerativeAiIntegrationSecretRequest instantiates a new ModelGenerativeAiIntegrationSecretRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelGenerativeAiIntegrationSecretRequest(name string, queryType string) *ModelGenerativeAiIntegrationSecretRequest {
	this := ModelGenerativeAiIntegrationSecretRequest{}
	this.Name = name
	this.QueryType = queryType
	return &this
}

// NewModelGenerativeAiIntegrationSecretRequestWithDefaults instantiates a new ModelGenerativeAiIntegrationSecretRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelGenerativeAiIntegrationSecretRequestWithDefaults() *ModelGenerativeAiIntegrationSecretRequest {
	this := ModelGenerativeAiIntegrationSecretRequest{}
	return &this
}

// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise.
func (o *ModelGenerativeAiIntegrationSecretRequest) GetIntegrationId() int32 {
	if o == nil || IsNil(o.IntegrationId) {
		var ret int32
		return ret
	}
	return *o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelGenerativeAiIntegrationSecretRequest) GetIntegrationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.IntegrationId) {
		return nil, false
	}
	return o.IntegrationId, true
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *ModelGenerativeAiIntegrationSecretRequest) HasIntegrationId() bool {
	if o != nil && !IsNil(o.IntegrationId) {
		return true
	}

	return false
}

// SetIntegrationId gets a reference to the given int32 and assigns it to the IntegrationId field.
func (o *ModelGenerativeAiIntegrationSecretRequest) SetIntegrationId(v int32) {
	o.IntegrationId = &v
}

// GetName returns the Name field value
func (o *ModelGenerativeAiIntegrationSecretRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModelGenerativeAiIntegrationSecretRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ModelGenerativeAiIntegrationSecretRequest) SetName(v string) {
	o.Name = v
}

// GetQueryType returns the QueryType field value
func (o *ModelGenerativeAiIntegrationSecretRequest) GetQueryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value
// and a boolean to check if the value has been set.
func (o *ModelGenerativeAiIntegrationSecretRequest) GetQueryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryType, true
}

// SetQueryType sets field value
func (o *ModelGenerativeAiIntegrationSecretRequest) SetQueryType(v string) {
	o.QueryType = v
}

func (o ModelGenerativeAiIntegrationSecretRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelGenerativeAiIntegrationSecretRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IntegrationId) {
		toSerialize["integration_id"] = o.IntegrationId
	}
	toSerialize["name"] = o.Name
	toSerialize["query_type"] = o.QueryType
	return toSerialize, nil
}

func (o *ModelGenerativeAiIntegrationSecretRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"query_type",
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

	varModelGenerativeAiIntegrationSecretRequest := _ModelGenerativeAiIntegrationSecretRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelGenerativeAiIntegrationSecretRequest)

	if err != nil {
		return err
	}

	*o = ModelGenerativeAiIntegrationSecretRequest(varModelGenerativeAiIntegrationSecretRequest)

	return err
}

type NullableModelGenerativeAiIntegrationSecretRequest struct {
	value *ModelGenerativeAiIntegrationSecretRequest
	isSet bool
}

func (v NullableModelGenerativeAiIntegrationSecretRequest) Get() *ModelGenerativeAiIntegrationSecretRequest {
	return v.value
}

func (v *NullableModelGenerativeAiIntegrationSecretRequest) Set(val *ModelGenerativeAiIntegrationSecretRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelGenerativeAiIntegrationSecretRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelGenerativeAiIntegrationSecretRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelGenerativeAiIntegrationSecretRequest(val *ModelGenerativeAiIntegrationSecretRequest) *NullableModelGenerativeAiIntegrationSecretRequest {
	return &NullableModelGenerativeAiIntegrationSecretRequest{value: val, isSet: true}
}

func (v NullableModelGenerativeAiIntegrationSecretRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelGenerativeAiIntegrationSecretRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


