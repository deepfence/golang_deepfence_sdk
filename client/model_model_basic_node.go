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

// checks if the ModelBasicNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelBasicNode{}

// ModelBasicNode struct for ModelBasicNode
type ModelBasicNode struct {
	Active bool `json:"active"`
	HostName string `json:"host_name"`
	Name string `json:"name"`
	NodeId string `json:"node_id"`
	NodeType string `json:"node_type"`
}

type _ModelBasicNode ModelBasicNode

// NewModelBasicNode instantiates a new ModelBasicNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelBasicNode(active bool, hostName string, name string, nodeId string, nodeType string) *ModelBasicNode {
	this := ModelBasicNode{}
	this.Active = active
	this.HostName = hostName
	this.Name = name
	this.NodeId = nodeId
	this.NodeType = nodeType
	return &this
}

// NewModelBasicNodeWithDefaults instantiates a new ModelBasicNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelBasicNodeWithDefaults() *ModelBasicNode {
	this := ModelBasicNode{}
	return &this
}

// GetActive returns the Active field value
func (o *ModelBasicNode) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *ModelBasicNode) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *ModelBasicNode) SetActive(v bool) {
	o.Active = v
}

// GetHostName returns the HostName field value
func (o *ModelBasicNode) GetHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value
// and a boolean to check if the value has been set.
func (o *ModelBasicNode) GetHostNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostName, true
}

// SetHostName sets field value
func (o *ModelBasicNode) SetHostName(v string) {
	o.HostName = v
}

// GetName returns the Name field value
func (o *ModelBasicNode) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModelBasicNode) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ModelBasicNode) SetName(v string) {
	o.Name = v
}

// GetNodeId returns the NodeId field value
func (o *ModelBasicNode) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelBasicNode) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelBasicNode) SetNodeId(v string) {
	o.NodeId = v
}

// GetNodeType returns the NodeType field value
func (o *ModelBasicNode) GetNodeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *ModelBasicNode) GetNodeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value
func (o *ModelBasicNode) SetNodeType(v string) {
	o.NodeType = v
}

func (o ModelBasicNode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelBasicNode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["host_name"] = o.HostName
	toSerialize["name"] = o.Name
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_type"] = o.NodeType
	return toSerialize, nil
}

func (o *ModelBasicNode) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"active",
		"host_name",
		"name",
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

	varModelBasicNode := _ModelBasicNode{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelBasicNode)

	if err != nil {
		return err
	}

	*o = ModelBasicNode(varModelBasicNode)

	return err
}

type NullableModelBasicNode struct {
	value *ModelBasicNode
	isSet bool
}

func (v NullableModelBasicNode) Get() *ModelBasicNode {
	return v.value
}

func (v *NullableModelBasicNode) Set(val *ModelBasicNode) {
	v.value = val
	v.isSet = true
}

func (v NullableModelBasicNode) IsSet() bool {
	return v.isSet
}

func (v *NullableModelBasicNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelBasicNode(val *ModelBasicNode) *NullableModelBasicNode {
	return &NullableModelBasicNode{value: val, isSet: true}
}

func (v NullableModelBasicNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelBasicNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


