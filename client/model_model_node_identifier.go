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

// checks if the ModelNodeIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelNodeIdentifier{}

// ModelNodeIdentifier struct for ModelNodeIdentifier
type ModelNodeIdentifier struct {
	NodeId string `json:"node_id"`
	NodeType string `json:"node_type"`
}

type _ModelNodeIdentifier ModelNodeIdentifier

// NewModelNodeIdentifier instantiates a new ModelNodeIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelNodeIdentifier(nodeId string, nodeType string) *ModelNodeIdentifier {
	this := ModelNodeIdentifier{}
	this.NodeId = nodeId
	this.NodeType = nodeType
	return &this
}

// NewModelNodeIdentifierWithDefaults instantiates a new ModelNodeIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelNodeIdentifierWithDefaults() *ModelNodeIdentifier {
	this := ModelNodeIdentifier{}
	return &this
}

// GetNodeId returns the NodeId field value
func (o *ModelNodeIdentifier) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelNodeIdentifier) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelNodeIdentifier) SetNodeId(v string) {
	o.NodeId = v
}

// GetNodeType returns the NodeType field value
func (o *ModelNodeIdentifier) GetNodeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *ModelNodeIdentifier) GetNodeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value
func (o *ModelNodeIdentifier) SetNodeType(v string) {
	o.NodeType = v
}

func (o ModelNodeIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelNodeIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_type"] = o.NodeType
	return toSerialize, nil
}

func (o *ModelNodeIdentifier) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"node_id",
		"node_type",
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

	varModelNodeIdentifier := _ModelNodeIdentifier{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelNodeIdentifier)

	if err != nil {
		return err
	}

	*o = ModelNodeIdentifier(varModelNodeIdentifier)

	return err
}

type NullableModelNodeIdentifier struct {
	value *ModelNodeIdentifier
	isSet bool
}

func (v NullableModelNodeIdentifier) Get() *ModelNodeIdentifier {
	return v.value
}

func (v *NullableModelNodeIdentifier) Set(val *ModelNodeIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableModelNodeIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableModelNodeIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelNodeIdentifier(val *ModelNodeIdentifier) *NullableModelNodeIdentifier {
	return &NullableModelNodeIdentifier{value: val, isSet: true}
}

func (v NullableModelNodeIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelNodeIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


