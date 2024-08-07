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

// checks if the ModelKubernetesCluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelKubernetesCluster{}

// ModelKubernetesCluster struct for ModelKubernetesCluster
type ModelKubernetesCluster struct {
	AgentRunning bool `json:"agent_running"`
	Hosts []ModelHost `json:"hosts"`
	NodeId string `json:"node_id"`
	NodeName string `json:"node_name"`
}

type _ModelKubernetesCluster ModelKubernetesCluster

// NewModelKubernetesCluster instantiates a new ModelKubernetesCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelKubernetesCluster(agentRunning bool, hosts []ModelHost, nodeId string, nodeName string) *ModelKubernetesCluster {
	this := ModelKubernetesCluster{}
	this.AgentRunning = agentRunning
	this.Hosts = hosts
	this.NodeId = nodeId
	this.NodeName = nodeName
	return &this
}

// NewModelKubernetesClusterWithDefaults instantiates a new ModelKubernetesCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelKubernetesClusterWithDefaults() *ModelKubernetesCluster {
	this := ModelKubernetesCluster{}
	return &this
}

// GetAgentRunning returns the AgentRunning field value
func (o *ModelKubernetesCluster) GetAgentRunning() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AgentRunning
}

// GetAgentRunningOk returns a tuple with the AgentRunning field value
// and a boolean to check if the value has been set.
func (o *ModelKubernetesCluster) GetAgentRunningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentRunning, true
}

// SetAgentRunning sets field value
func (o *ModelKubernetesCluster) SetAgentRunning(v bool) {
	o.AgentRunning = v
}

// GetHosts returns the Hosts field value
// If the value is explicit nil, the zero value for []ModelHost will be returned
func (o *ModelKubernetesCluster) GetHosts() []ModelHost {
	if o == nil {
		var ret []ModelHost
		return ret
	}

	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelKubernetesCluster) GetHostsOk() ([]ModelHost, bool) {
	if o == nil || IsNil(o.Hosts) {
		return nil, false
	}
	return o.Hosts, true
}

// SetHosts sets field value
func (o *ModelKubernetesCluster) SetHosts(v []ModelHost) {
	o.Hosts = v
}

// GetNodeId returns the NodeId field value
func (o *ModelKubernetesCluster) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelKubernetesCluster) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelKubernetesCluster) SetNodeId(v string) {
	o.NodeId = v
}

// GetNodeName returns the NodeName field value
func (o *ModelKubernetesCluster) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *ModelKubernetesCluster) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value
func (o *ModelKubernetesCluster) SetNodeName(v string) {
	o.NodeName = v
}

func (o ModelKubernetesCluster) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelKubernetesCluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agent_running"] = o.AgentRunning
	if o.Hosts != nil {
		toSerialize["hosts"] = o.Hosts
	}
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_name"] = o.NodeName
	return toSerialize, nil
}

func (o *ModelKubernetesCluster) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"agent_running",
		"hosts",
		"node_id",
		"node_name",
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

	varModelKubernetesCluster := _ModelKubernetesCluster{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelKubernetesCluster)

	if err != nil {
		return err
	}

	*o = ModelKubernetesCluster(varModelKubernetesCluster)

	return err
}

type NullableModelKubernetesCluster struct {
	value *ModelKubernetesCluster
	isSet bool
}

func (v NullableModelKubernetesCluster) Get() *ModelKubernetesCluster {
	return v.value
}

func (v *NullableModelKubernetesCluster) Set(val *ModelKubernetesCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableModelKubernetesCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableModelKubernetesCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelKubernetesCluster(val *ModelKubernetesCluster) *NullableModelKubernetesCluster {
	return &NullableModelKubernetesCluster{value: val, isSet: true}
}

func (v NullableModelKubernetesCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelKubernetesCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


