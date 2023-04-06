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

// checks if the ModelPod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelPod{}

// ModelPod struct for ModelPod
type ModelPod struct {
	Containers []ModelContainer `json:"containers"`
	HostName string `json:"host_name"`
	KubernetesClusterId string `json:"kubernetes_cluster_id"`
	KubernetesClusterName string `json:"kubernetes_cluster_name"`
	KubernetesIp string `json:"kubernetes_ip"`
	KubernetesIsInHostNetwork string `json:"kubernetes_is_in_host_network"`
	KubernetesNamespace string `json:"kubernetes_namespace"`
	KubernetesState string `json:"kubernetes_state"`
	NodeId string `json:"node_id"`
	NodeName string `json:"node_name"`
	Processes []ModelProcess `json:"processes"`
}

// NewModelPod instantiates a new ModelPod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelPod(containers []ModelContainer, hostName string, kubernetesClusterId string, kubernetesClusterName string, kubernetesIp string, kubernetesIsInHostNetwork string, kubernetesNamespace string, kubernetesState string, nodeId string, nodeName string, processes []ModelProcess) *ModelPod {
	this := ModelPod{}
	this.Containers = containers
	this.HostName = hostName
	this.KubernetesClusterId = kubernetesClusterId
	this.KubernetesClusterName = kubernetesClusterName
	this.KubernetesIp = kubernetesIp
	this.KubernetesIsInHostNetwork = kubernetesIsInHostNetwork
	this.KubernetesNamespace = kubernetesNamespace
	this.KubernetesState = kubernetesState
	this.NodeId = nodeId
	this.NodeName = nodeName
	this.Processes = processes
	return &this
}

// NewModelPodWithDefaults instantiates a new ModelPod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelPodWithDefaults() *ModelPod {
	this := ModelPod{}
	return &this
}

// GetContainers returns the Containers field value
// If the value is explicit nil, the zero value for []ModelContainer will be returned
func (o *ModelPod) GetContainers() []ModelContainer {
	if o == nil {
		var ret []ModelContainer
		return ret
	}

	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPod) GetContainersOk() ([]ModelContainer, bool) {
	if o == nil || IsNil(o.Containers) {
		return nil, false
	}
	return o.Containers, true
}

// SetContainers sets field value
func (o *ModelPod) SetContainers(v []ModelContainer) {
	o.Containers = v
}

// GetHostName returns the HostName field value
func (o *ModelPod) GetHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value
// and a boolean to check if the value has been set.
func (o *ModelPod) GetHostNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostName, true
}

// SetHostName sets field value
func (o *ModelPod) SetHostName(v string) {
	o.HostName = v
}

// GetKubernetesClusterId returns the KubernetesClusterId field value
func (o *ModelPod) GetKubernetesClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubernetesClusterId
}

// GetKubernetesClusterIdOk returns a tuple with the KubernetesClusterId field value
// and a boolean to check if the value has been set.
func (o *ModelPod) GetKubernetesClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubernetesClusterId, true
}

// SetKubernetesClusterId sets field value
func (o *ModelPod) SetKubernetesClusterId(v string) {
	o.KubernetesClusterId = v
}

// GetKubernetesClusterName returns the KubernetesClusterName field value
func (o *ModelPod) GetKubernetesClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubernetesClusterName
}

// GetKubernetesClusterNameOk returns a tuple with the KubernetesClusterName field value
// and a boolean to check if the value has been set.
func (o *ModelPod) GetKubernetesClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubernetesClusterName, true
}

// SetKubernetesClusterName sets field value
func (o *ModelPod) SetKubernetesClusterName(v string) {
	o.KubernetesClusterName = v
}

// GetKubernetesIp returns the KubernetesIp field value
func (o *ModelPod) GetKubernetesIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubernetesIp
}

// GetKubernetesIpOk returns a tuple with the KubernetesIp field value
// and a boolean to check if the value has been set.
func (o *ModelPod) GetKubernetesIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubernetesIp, true
}

// SetKubernetesIp sets field value
func (o *ModelPod) SetKubernetesIp(v string) {
	o.KubernetesIp = v
}

// GetKubernetesIsInHostNetwork returns the KubernetesIsInHostNetwork field value
func (o *ModelPod) GetKubernetesIsInHostNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubernetesIsInHostNetwork
}

// GetKubernetesIsInHostNetworkOk returns a tuple with the KubernetesIsInHostNetwork field value
// and a boolean to check if the value has been set.
func (o *ModelPod) GetKubernetesIsInHostNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubernetesIsInHostNetwork, true
}

// SetKubernetesIsInHostNetwork sets field value
func (o *ModelPod) SetKubernetesIsInHostNetwork(v string) {
	o.KubernetesIsInHostNetwork = v
}

// GetKubernetesNamespace returns the KubernetesNamespace field value
func (o *ModelPod) GetKubernetesNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubernetesNamespace
}

// GetKubernetesNamespaceOk returns a tuple with the KubernetesNamespace field value
// and a boolean to check if the value has been set.
func (o *ModelPod) GetKubernetesNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubernetesNamespace, true
}

// SetKubernetesNamespace sets field value
func (o *ModelPod) SetKubernetesNamespace(v string) {
	o.KubernetesNamespace = v
}

// GetKubernetesState returns the KubernetesState field value
func (o *ModelPod) GetKubernetesState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubernetesState
}

// GetKubernetesStateOk returns a tuple with the KubernetesState field value
// and a boolean to check if the value has been set.
func (o *ModelPod) GetKubernetesStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubernetesState, true
}

// SetKubernetesState sets field value
func (o *ModelPod) SetKubernetesState(v string) {
	o.KubernetesState = v
}

// GetNodeId returns the NodeId field value
func (o *ModelPod) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelPod) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelPod) SetNodeId(v string) {
	o.NodeId = v
}

// GetNodeName returns the NodeName field value
func (o *ModelPod) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *ModelPod) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value
func (o *ModelPod) SetNodeName(v string) {
	o.NodeName = v
}

// GetProcesses returns the Processes field value
// If the value is explicit nil, the zero value for []ModelProcess will be returned
func (o *ModelPod) GetProcesses() []ModelProcess {
	if o == nil {
		var ret []ModelProcess
		return ret
	}

	return o.Processes
}

// GetProcessesOk returns a tuple with the Processes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPod) GetProcessesOk() ([]ModelProcess, bool) {
	if o == nil || IsNil(o.Processes) {
		return nil, false
	}
	return o.Processes, true
}

// SetProcesses sets field value
func (o *ModelPod) SetProcesses(v []ModelProcess) {
	o.Processes = v
}

func (o ModelPod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelPod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Containers != nil {
		toSerialize["containers"] = o.Containers
	}
	toSerialize["host_name"] = o.HostName
	toSerialize["kubernetes_cluster_id"] = o.KubernetesClusterId
	toSerialize["kubernetes_cluster_name"] = o.KubernetesClusterName
	toSerialize["kubernetes_ip"] = o.KubernetesIp
	toSerialize["kubernetes_is_in_host_network"] = o.KubernetesIsInHostNetwork
	toSerialize["kubernetes_namespace"] = o.KubernetesNamespace
	toSerialize["kubernetes_state"] = o.KubernetesState
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_name"] = o.NodeName
	if o.Processes != nil {
		toSerialize["processes"] = o.Processes
	}
	return toSerialize, nil
}

type NullableModelPod struct {
	value *ModelPod
	isSet bool
}

func (v NullableModelPod) Get() *ModelPod {
	return v.value
}

func (v *NullableModelPod) Set(val *ModelPod) {
	v.value = val
	v.isSet = true
}

func (v NullableModelPod) IsSet() bool {
	return v.isSet
}

func (v *NullableModelPod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelPod(val *ModelPod) *NullableModelPod {
	return &NullableModelPod{value: val, isSet: true}
}

func (v NullableModelPod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelPod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


