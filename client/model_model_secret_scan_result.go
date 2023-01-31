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

// checks if the ModelSecretScanResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSecretScanResult{}

// ModelSecretScanResult struct for ModelSecretScanResult
type ModelSecretScanResult struct {
	ImageLayerId string `json:"ImageLayerId"`
	ContainerName string `json:"container_name"`
	HostName string `json:"host_name"`
	KubernetesClusterName string `json:"kubernetes_cluster_name"`
	Masked string `json:"masked"`
	NodeId string `json:"node_id"`
	NodeName string `json:"node_name"`
	NodeType string `json:"node_type"`
	Rule2Secrets map[string][]int32 `json:"rule_2_secrets"`
	Rules []ModelRule `json:"rules"`
	ScanId string `json:"scan_id"`
	Secrets []ModelSecret `json:"secrets"`
}

// NewModelSecretScanResult instantiates a new ModelSecretScanResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSecretScanResult(imageLayerId string, containerName string, hostName string, kubernetesClusterName string, masked string, nodeId string, nodeName string, nodeType string, rule2Secrets map[string][]int32, rules []ModelRule, scanId string, secrets []ModelSecret) *ModelSecretScanResult {
	this := ModelSecretScanResult{}
	this.ImageLayerId = imageLayerId
	this.ContainerName = containerName
	this.HostName = hostName
	this.KubernetesClusterName = kubernetesClusterName
	this.Masked = masked
	this.NodeId = nodeId
	this.NodeName = nodeName
	this.NodeType = nodeType
	this.Rule2Secrets = rule2Secrets
	this.Rules = rules
	this.ScanId = scanId
	this.Secrets = secrets
	return &this
}

// NewModelSecretScanResultWithDefaults instantiates a new ModelSecretScanResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSecretScanResultWithDefaults() *ModelSecretScanResult {
	this := ModelSecretScanResult{}
	return &this
}

// GetImageLayerId returns the ImageLayerId field value
func (o *ModelSecretScanResult) GetImageLayerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageLayerId
}

// GetImageLayerIdOk returns a tuple with the ImageLayerId field value
// and a boolean to check if the value has been set.
func (o *ModelSecretScanResult) GetImageLayerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageLayerId, true
}

// SetImageLayerId sets field value
func (o *ModelSecretScanResult) SetImageLayerId(v string) {
	o.ImageLayerId = v
}

// GetContainerName returns the ContainerName field value
func (o *ModelSecretScanResult) GetContainerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value
// and a boolean to check if the value has been set.
func (o *ModelSecretScanResult) GetContainerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerName, true
}

// SetContainerName sets field value
func (o *ModelSecretScanResult) SetContainerName(v string) {
	o.ContainerName = v
}

// GetHostName returns the HostName field value
func (o *ModelSecretScanResult) GetHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value
// and a boolean to check if the value has been set.
func (o *ModelSecretScanResult) GetHostNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostName, true
}

// SetHostName sets field value
func (o *ModelSecretScanResult) SetHostName(v string) {
	o.HostName = v
}

// GetKubernetesClusterName returns the KubernetesClusterName field value
func (o *ModelSecretScanResult) GetKubernetesClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubernetesClusterName
}

// GetKubernetesClusterNameOk returns a tuple with the KubernetesClusterName field value
// and a boolean to check if the value has been set.
func (o *ModelSecretScanResult) GetKubernetesClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubernetesClusterName, true
}

// SetKubernetesClusterName sets field value
func (o *ModelSecretScanResult) SetKubernetesClusterName(v string) {
	o.KubernetesClusterName = v
}

// GetMasked returns the Masked field value
func (o *ModelSecretScanResult) GetMasked() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Masked
}

// GetMaskedOk returns a tuple with the Masked field value
// and a boolean to check if the value has been set.
func (o *ModelSecretScanResult) GetMaskedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Masked, true
}

// SetMasked sets field value
func (o *ModelSecretScanResult) SetMasked(v string) {
	o.Masked = v
}

// GetNodeId returns the NodeId field value
func (o *ModelSecretScanResult) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelSecretScanResult) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelSecretScanResult) SetNodeId(v string) {
	o.NodeId = v
}

// GetNodeName returns the NodeName field value
func (o *ModelSecretScanResult) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *ModelSecretScanResult) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value
func (o *ModelSecretScanResult) SetNodeName(v string) {
	o.NodeName = v
}

// GetNodeType returns the NodeType field value
func (o *ModelSecretScanResult) GetNodeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *ModelSecretScanResult) GetNodeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value
func (o *ModelSecretScanResult) SetNodeType(v string) {
	o.NodeType = v
}

// GetRule2Secrets returns the Rule2Secrets field value
// If the value is explicit nil, the zero value for map[string][]int32 will be returned
func (o *ModelSecretScanResult) GetRule2Secrets() map[string][]int32 {
	if o == nil {
		var ret map[string][]int32
		return ret
	}

	return o.Rule2Secrets
}

// GetRule2SecretsOk returns a tuple with the Rule2Secrets field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSecretScanResult) GetRule2SecretsOk() (*map[string][]int32, bool) {
	if o == nil || isNil(o.Rule2Secrets) {
		return nil, false
	}
	return &o.Rule2Secrets, true
}

// SetRule2Secrets sets field value
func (o *ModelSecretScanResult) SetRule2Secrets(v map[string][]int32) {
	o.Rule2Secrets = v
}

// GetRules returns the Rules field value
// If the value is explicit nil, the zero value for []ModelRule will be returned
func (o *ModelSecretScanResult) GetRules() []ModelRule {
	if o == nil {
		var ret []ModelRule
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSecretScanResult) GetRulesOk() ([]ModelRule, bool) {
	if o == nil || isNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *ModelSecretScanResult) SetRules(v []ModelRule) {
	o.Rules = v
}

// GetScanId returns the ScanId field value
func (o *ModelSecretScanResult) GetScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value
// and a boolean to check if the value has been set.
func (o *ModelSecretScanResult) GetScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScanId, true
}

// SetScanId sets field value
func (o *ModelSecretScanResult) SetScanId(v string) {
	o.ScanId = v
}

// GetSecrets returns the Secrets field value
// If the value is explicit nil, the zero value for []ModelSecret will be returned
func (o *ModelSecretScanResult) GetSecrets() []ModelSecret {
	if o == nil {
		var ret []ModelSecret
		return ret
	}

	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSecretScanResult) GetSecretsOk() ([]ModelSecret, bool) {
	if o == nil || isNil(o.Secrets) {
		return nil, false
	}
	return o.Secrets, true
}

// SetSecrets sets field value
func (o *ModelSecretScanResult) SetSecrets(v []ModelSecret) {
	o.Secrets = v
}

func (o ModelSecretScanResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSecretScanResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ImageLayerId"] = o.ImageLayerId
	toSerialize["container_name"] = o.ContainerName
	toSerialize["host_name"] = o.HostName
	toSerialize["kubernetes_cluster_name"] = o.KubernetesClusterName
	toSerialize["masked"] = o.Masked
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_name"] = o.NodeName
	toSerialize["node_type"] = o.NodeType
	if o.Rule2Secrets != nil {
		toSerialize["rule_2_secrets"] = o.Rule2Secrets
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	toSerialize["scan_id"] = o.ScanId
	if o.Secrets != nil {
		toSerialize["secrets"] = o.Secrets
	}
	return toSerialize, nil
}

type NullableModelSecretScanResult struct {
	value *ModelSecretScanResult
	isSet bool
}

func (v NullableModelSecretScanResult) Get() *ModelSecretScanResult {
	return v.value
}

func (v *NullableModelSecretScanResult) Set(val *ModelSecretScanResult) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSecretScanResult) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSecretScanResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSecretScanResult(val *ModelSecretScanResult) *NullableModelSecretScanResult {
	return &NullableModelSecretScanResult{value: val, isSet: true}
}

func (v NullableModelSecretScanResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSecretScanResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


