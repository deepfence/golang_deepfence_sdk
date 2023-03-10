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

// checks if the ModelContainerImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelContainerImage{}

// ModelContainerImage struct for ModelContainerImage
type ModelContainerImage struct {
	ComplianceScanStatus string `json:"compliance_scan_status"`
	CompliancesCount int32 `json:"compliances_count"`
	DockerImageName string `json:"docker_image_name"`
	DockerImageSize string `json:"docker_image_size"`
	DockerImageTag string `json:"docker_image_tag"`
	MalwareScanStatus string `json:"malware_scan_status"`
	MalwaresCount int32 `json:"malwares_count"`
	Metadata map[string]interface{} `json:"metadata"`
	Metrics ModelComputeMetrics `json:"metrics"`
	NodeId string `json:"node_id"`
	NodeName string `json:"node_name"`
	SecretScanStatus string `json:"secret_scan_status"`
	SecretsCount int32 `json:"secrets_count"`
	VulnerabilitiesCount int32 `json:"vulnerabilities_count"`
	VulnerabilityScanStatus string `json:"vulnerability_scan_status"`
}

// NewModelContainerImage instantiates a new ModelContainerImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelContainerImage(complianceScanStatus string, compliancesCount int32, dockerImageName string, dockerImageSize string, dockerImageTag string, malwareScanStatus string, malwaresCount int32, metadata map[string]interface{}, metrics ModelComputeMetrics, nodeId string, nodeName string, secretScanStatus string, secretsCount int32, vulnerabilitiesCount int32, vulnerabilityScanStatus string) *ModelContainerImage {
	this := ModelContainerImage{}
	this.ComplianceScanStatus = complianceScanStatus
	this.CompliancesCount = compliancesCount
	this.DockerImageName = dockerImageName
	this.DockerImageSize = dockerImageSize
	this.DockerImageTag = dockerImageTag
	this.MalwareScanStatus = malwareScanStatus
	this.MalwaresCount = malwaresCount
	this.Metadata = metadata
	this.Metrics = metrics
	this.NodeId = nodeId
	this.NodeName = nodeName
	this.SecretScanStatus = secretScanStatus
	this.SecretsCount = secretsCount
	this.VulnerabilitiesCount = vulnerabilitiesCount
	this.VulnerabilityScanStatus = vulnerabilityScanStatus
	return &this
}

// NewModelContainerImageWithDefaults instantiates a new ModelContainerImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelContainerImageWithDefaults() *ModelContainerImage {
	this := ModelContainerImage{}
	return &this
}

// GetComplianceScanStatus returns the ComplianceScanStatus field value
func (o *ModelContainerImage) GetComplianceScanStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComplianceScanStatus
}

// GetComplianceScanStatusOk returns a tuple with the ComplianceScanStatus field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetComplianceScanStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComplianceScanStatus, true
}

// SetComplianceScanStatus sets field value
func (o *ModelContainerImage) SetComplianceScanStatus(v string) {
	o.ComplianceScanStatus = v
}

// GetCompliancesCount returns the CompliancesCount field value
func (o *ModelContainerImage) GetCompliancesCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CompliancesCount
}

// GetCompliancesCountOk returns a tuple with the CompliancesCount field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetCompliancesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompliancesCount, true
}

// SetCompliancesCount sets field value
func (o *ModelContainerImage) SetCompliancesCount(v int32) {
	o.CompliancesCount = v
}

// GetDockerImageName returns the DockerImageName field value
func (o *ModelContainerImage) GetDockerImageName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerImageName
}

// GetDockerImageNameOk returns a tuple with the DockerImageName field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetDockerImageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerImageName, true
}

// SetDockerImageName sets field value
func (o *ModelContainerImage) SetDockerImageName(v string) {
	o.DockerImageName = v
}

// GetDockerImageSize returns the DockerImageSize field value
func (o *ModelContainerImage) GetDockerImageSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerImageSize
}

// GetDockerImageSizeOk returns a tuple with the DockerImageSize field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetDockerImageSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerImageSize, true
}

// SetDockerImageSize sets field value
func (o *ModelContainerImage) SetDockerImageSize(v string) {
	o.DockerImageSize = v
}

// GetDockerImageTag returns the DockerImageTag field value
func (o *ModelContainerImage) GetDockerImageTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerImageTag
}

// GetDockerImageTagOk returns a tuple with the DockerImageTag field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetDockerImageTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerImageTag, true
}

// SetDockerImageTag sets field value
func (o *ModelContainerImage) SetDockerImageTag(v string) {
	o.DockerImageTag = v
}

// GetMalwareScanStatus returns the MalwareScanStatus field value
func (o *ModelContainerImage) GetMalwareScanStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MalwareScanStatus
}

// GetMalwareScanStatusOk returns a tuple with the MalwareScanStatus field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetMalwareScanStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MalwareScanStatus, true
}

// SetMalwareScanStatus sets field value
func (o *ModelContainerImage) SetMalwareScanStatus(v string) {
	o.MalwareScanStatus = v
}

// GetMalwaresCount returns the MalwaresCount field value
func (o *ModelContainerImage) GetMalwaresCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MalwaresCount
}

// GetMalwaresCountOk returns a tuple with the MalwaresCount field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetMalwaresCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MalwaresCount, true
}

// SetMalwaresCount sets field value
func (o *ModelContainerImage) SetMalwaresCount(v int32) {
	o.MalwaresCount = v
}

// GetMetadata returns the Metadata field value
func (o *ModelContainerImage) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *ModelContainerImage) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMetrics returns the Metrics field value
func (o *ModelContainerImage) GetMetrics() ModelComputeMetrics {
	if o == nil {
		var ret ModelComputeMetrics
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetMetricsOk() (*ModelComputeMetrics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// SetMetrics sets field value
func (o *ModelContainerImage) SetMetrics(v ModelComputeMetrics) {
	o.Metrics = v
}

// GetNodeId returns the NodeId field value
func (o *ModelContainerImage) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelContainerImage) SetNodeId(v string) {
	o.NodeId = v
}

// GetNodeName returns the NodeName field value
func (o *ModelContainerImage) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value
func (o *ModelContainerImage) SetNodeName(v string) {
	o.NodeName = v
}

// GetSecretScanStatus returns the SecretScanStatus field value
func (o *ModelContainerImage) GetSecretScanStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretScanStatus
}

// GetSecretScanStatusOk returns a tuple with the SecretScanStatus field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetSecretScanStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretScanStatus, true
}

// SetSecretScanStatus sets field value
func (o *ModelContainerImage) SetSecretScanStatus(v string) {
	o.SecretScanStatus = v
}

// GetSecretsCount returns the SecretsCount field value
func (o *ModelContainerImage) GetSecretsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SecretsCount
}

// GetSecretsCountOk returns a tuple with the SecretsCount field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetSecretsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretsCount, true
}

// SetSecretsCount sets field value
func (o *ModelContainerImage) SetSecretsCount(v int32) {
	o.SecretsCount = v
}

// GetVulnerabilitiesCount returns the VulnerabilitiesCount field value
func (o *ModelContainerImage) GetVulnerabilitiesCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VulnerabilitiesCount
}

// GetVulnerabilitiesCountOk returns a tuple with the VulnerabilitiesCount field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetVulnerabilitiesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VulnerabilitiesCount, true
}

// SetVulnerabilitiesCount sets field value
func (o *ModelContainerImage) SetVulnerabilitiesCount(v int32) {
	o.VulnerabilitiesCount = v
}

// GetVulnerabilityScanStatus returns the VulnerabilityScanStatus field value
func (o *ModelContainerImage) GetVulnerabilityScanStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VulnerabilityScanStatus
}

// GetVulnerabilityScanStatusOk returns a tuple with the VulnerabilityScanStatus field value
// and a boolean to check if the value has been set.
func (o *ModelContainerImage) GetVulnerabilityScanStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VulnerabilityScanStatus, true
}

// SetVulnerabilityScanStatus sets field value
func (o *ModelContainerImage) SetVulnerabilityScanStatus(v string) {
	o.VulnerabilityScanStatus = v
}

func (o ModelContainerImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelContainerImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["compliance_scan_status"] = o.ComplianceScanStatus
	toSerialize["compliances_count"] = o.CompliancesCount
	toSerialize["docker_image_name"] = o.DockerImageName
	toSerialize["docker_image_size"] = o.DockerImageSize
	toSerialize["docker_image_tag"] = o.DockerImageTag
	toSerialize["malware_scan_status"] = o.MalwareScanStatus
	toSerialize["malwares_count"] = o.MalwaresCount
	toSerialize["metadata"] = o.Metadata
	toSerialize["metrics"] = o.Metrics
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_name"] = o.NodeName
	toSerialize["secret_scan_status"] = o.SecretScanStatus
	toSerialize["secrets_count"] = o.SecretsCount
	toSerialize["vulnerabilities_count"] = o.VulnerabilitiesCount
	toSerialize["vulnerability_scan_status"] = o.VulnerabilityScanStatus
	return toSerialize, nil
}

type NullableModelContainerImage struct {
	value *ModelContainerImage
	isSet bool
}

func (v NullableModelContainerImage) Get() *ModelContainerImage {
	return v.value
}

func (v *NullableModelContainerImage) Set(val *ModelContainerImage) {
	v.value = val
	v.isSet = true
}

func (v NullableModelContainerImage) IsSet() bool {
	return v.isSet
}

func (v *NullableModelContainerImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelContainerImage(val *ModelContainerImage) *NullableModelContainerImage {
	return &NullableModelContainerImage{value: val, isSet: true}
}

func (v NullableModelContainerImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelContainerImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


