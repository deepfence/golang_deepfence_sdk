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

// checks if the ModelPod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelPod{}

// ModelPod struct for ModelPod
type ModelPod struct {
	Containers []ModelContainer `json:"containers"`
	HostName string `json:"host_name"`
	IsDeepfenceSystem bool `json:"is_deepfence_system"`
	KubernetesClusterId string `json:"kubernetes_cluster_id"`
	KubernetesClusterName string `json:"kubernetes_cluster_name"`
	KubernetesCreated string `json:"kubernetes_created"`
	KubernetesIp string `json:"kubernetes_ip"`
	KubernetesIsInHostNetwork bool `json:"kubernetes_is_in_host_network"`
	KubernetesLabels map[string]interface{} `json:"kubernetes_labels"`
	KubernetesNamespace string `json:"kubernetes_namespace"`
	KubernetesState string `json:"kubernetes_state"`
	MalwareScanStatus string `json:"malware_scan_status"`
	NodeId string `json:"node_id"`
	NodeName string `json:"node_name"`
	PodName string `json:"pod_name"`
	Processes []ModelProcess `json:"processes"`
	SecretScanStatus string `json:"secret_scan_status"`
	VulnerabilityScanStatus string `json:"vulnerability_scan_status"`
}

type _ModelPod ModelPod

// NewModelPod instantiates a new ModelPod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelPod(containers []ModelContainer, hostName string, isDeepfenceSystem bool, kubernetesClusterId string, kubernetesClusterName string, kubernetesCreated string, kubernetesIp string, kubernetesIsInHostNetwork bool, kubernetesLabels map[string]interface{}, kubernetesNamespace string, kubernetesState string, malwareScanStatus string, nodeId string, nodeName string, podName string, processes []ModelProcess, secretScanStatus string, vulnerabilityScanStatus string) *ModelPod {
	this := ModelPod{}
	this.Containers = containers
	this.HostName = hostName
	this.IsDeepfenceSystem = isDeepfenceSystem
	this.KubernetesClusterId = kubernetesClusterId
	this.KubernetesClusterName = kubernetesClusterName
	this.KubernetesCreated = kubernetesCreated
	this.KubernetesIp = kubernetesIp
	this.KubernetesIsInHostNetwork = kubernetesIsInHostNetwork
	this.KubernetesLabels = kubernetesLabels
	this.KubernetesNamespace = kubernetesNamespace
	this.KubernetesState = kubernetesState
	this.MalwareScanStatus = malwareScanStatus
	this.NodeId = nodeId
	this.NodeName = nodeName
	this.PodName = podName
	this.Processes = processes
	this.SecretScanStatus = secretScanStatus
	this.VulnerabilityScanStatus = vulnerabilityScanStatus
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

// GetIsDeepfenceSystem returns the IsDeepfenceSystem field value
func (o *ModelPod) GetIsDeepfenceSystem() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeepfenceSystem
}

// GetIsDeepfenceSystemOk returns a tuple with the IsDeepfenceSystem field value
// and a boolean to check if the value has been set.
func (o *ModelPod) GetIsDeepfenceSystemOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeepfenceSystem, true
}

// SetIsDeepfenceSystem sets field value
func (o *ModelPod) SetIsDeepfenceSystem(v bool) {
	o.IsDeepfenceSystem = v
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

// GetKubernetesCreated returns the KubernetesCreated field value
func (o *ModelPod) GetKubernetesCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubernetesCreated
}

// GetKubernetesCreatedOk returns a tuple with the KubernetesCreated field value
// and a boolean to check if the value has been set.
func (o *ModelPod) GetKubernetesCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubernetesCreated, true
}

// SetKubernetesCreated sets field value
func (o *ModelPod) SetKubernetesCreated(v string) {
	o.KubernetesCreated = v
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
func (o *ModelPod) GetKubernetesIsInHostNetwork() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.KubernetesIsInHostNetwork
}

// GetKubernetesIsInHostNetworkOk returns a tuple with the KubernetesIsInHostNetwork field value
// and a boolean to check if the value has been set.
func (o *ModelPod) GetKubernetesIsInHostNetworkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubernetesIsInHostNetwork, true
}

// SetKubernetesIsInHostNetwork sets field value
func (o *ModelPod) SetKubernetesIsInHostNetwork(v bool) {
	o.KubernetesIsInHostNetwork = v
}

// GetKubernetesLabels returns the KubernetesLabels field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *ModelPod) GetKubernetesLabels() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.KubernetesLabels
}

// GetKubernetesLabelsOk returns a tuple with the KubernetesLabels field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPod) GetKubernetesLabelsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.KubernetesLabels) {
		return map[string]interface{}{}, false
	}
	return o.KubernetesLabels, true
}

// SetKubernetesLabels sets field value
func (o *ModelPod) SetKubernetesLabels(v map[string]interface{}) {
	o.KubernetesLabels = v
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

// GetMalwareScanStatus returns the MalwareScanStatus field value
func (o *ModelPod) GetMalwareScanStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MalwareScanStatus
}

// GetMalwareScanStatusOk returns a tuple with the MalwareScanStatus field value
// and a boolean to check if the value has been set.
func (o *ModelPod) GetMalwareScanStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MalwareScanStatus, true
}

// SetMalwareScanStatus sets field value
func (o *ModelPod) SetMalwareScanStatus(v string) {
	o.MalwareScanStatus = v
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

// GetPodName returns the PodName field value
func (o *ModelPod) GetPodName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PodName
}

// GetPodNameOk returns a tuple with the PodName field value
// and a boolean to check if the value has been set.
func (o *ModelPod) GetPodNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PodName, true
}

// SetPodName sets field value
func (o *ModelPod) SetPodName(v string) {
	o.PodName = v
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

// GetSecretScanStatus returns the SecretScanStatus field value
func (o *ModelPod) GetSecretScanStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretScanStatus
}

// GetSecretScanStatusOk returns a tuple with the SecretScanStatus field value
// and a boolean to check if the value has been set.
func (o *ModelPod) GetSecretScanStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretScanStatus, true
}

// SetSecretScanStatus sets field value
func (o *ModelPod) SetSecretScanStatus(v string) {
	o.SecretScanStatus = v
}

// GetVulnerabilityScanStatus returns the VulnerabilityScanStatus field value
func (o *ModelPod) GetVulnerabilityScanStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VulnerabilityScanStatus
}

// GetVulnerabilityScanStatusOk returns a tuple with the VulnerabilityScanStatus field value
// and a boolean to check if the value has been set.
func (o *ModelPod) GetVulnerabilityScanStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VulnerabilityScanStatus, true
}

// SetVulnerabilityScanStatus sets field value
func (o *ModelPod) SetVulnerabilityScanStatus(v string) {
	o.VulnerabilityScanStatus = v
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
	toSerialize["is_deepfence_system"] = o.IsDeepfenceSystem
	toSerialize["kubernetes_cluster_id"] = o.KubernetesClusterId
	toSerialize["kubernetes_cluster_name"] = o.KubernetesClusterName
	toSerialize["kubernetes_created"] = o.KubernetesCreated
	toSerialize["kubernetes_ip"] = o.KubernetesIp
	toSerialize["kubernetes_is_in_host_network"] = o.KubernetesIsInHostNetwork
	if o.KubernetesLabels != nil {
		toSerialize["kubernetes_labels"] = o.KubernetesLabels
	}
	toSerialize["kubernetes_namespace"] = o.KubernetesNamespace
	toSerialize["kubernetes_state"] = o.KubernetesState
	toSerialize["malware_scan_status"] = o.MalwareScanStatus
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_name"] = o.NodeName
	toSerialize["pod_name"] = o.PodName
	if o.Processes != nil {
		toSerialize["processes"] = o.Processes
	}
	toSerialize["secret_scan_status"] = o.SecretScanStatus
	toSerialize["vulnerability_scan_status"] = o.VulnerabilityScanStatus
	return toSerialize, nil
}

func (o *ModelPod) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"containers",
		"host_name",
		"is_deepfence_system",
		"kubernetes_cluster_id",
		"kubernetes_cluster_name",
		"kubernetes_created",
		"kubernetes_ip",
		"kubernetes_is_in_host_network",
		"kubernetes_labels",
		"kubernetes_namespace",
		"kubernetes_state",
		"malware_scan_status",
		"node_id",
		"node_name",
		"pod_name",
		"processes",
		"secret_scan_status",
		"vulnerability_scan_status",
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

	varModelPod := _ModelPod{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelPod)

	if err != nil {
		return err
	}

	*o = ModelPod(varModelPod)

	return err
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


