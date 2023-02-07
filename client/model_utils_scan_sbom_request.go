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

// checks if the UtilsScanSbomRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UtilsScanSbomRequest{}

// UtilsScanSbomRequest struct for UtilsScanSbomRequest
type UtilsScanSbomRequest struct {
	ContainerName *string `json:"container_name,omitempty"`
	HostName *string `json:"host_name,omitempty"`
	ImageId *string `json:"image_id,omitempty"`
	ImageName *string `json:"image_name,omitempty"`
	KubernetesClusterName *string `json:"kubernetes_cluster_name,omitempty"`
	Mode *string `json:"mode,omitempty"`
	NodeId *string `json:"node_id,omitempty"`
	NodeType *string `json:"node_type,omitempty"`
	RegistryId *string `json:"registry_id,omitempty"`
	Sbom string `json:"sbom"`
	SbomFilePath *string `json:"sbom_file_path,omitempty"`
	ScanId string `json:"scan_id"`
	ScanType *string `json:"scan_type,omitempty"`
}

// NewUtilsScanSbomRequest instantiates a new UtilsScanSbomRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUtilsScanSbomRequest(sbom string, scanId string) *UtilsScanSbomRequest {
	this := UtilsScanSbomRequest{}
	this.Sbom = sbom
	this.ScanId = scanId
	return &this
}

// NewUtilsScanSbomRequestWithDefaults instantiates a new UtilsScanSbomRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUtilsScanSbomRequestWithDefaults() *UtilsScanSbomRequest {
	this := UtilsScanSbomRequest{}
	return &this
}

// GetContainerName returns the ContainerName field value if set, zero value otherwise.
func (o *UtilsScanSbomRequest) GetContainerName() string {
	if o == nil || isNil(o.ContainerName) {
		var ret string
		return ret
	}
	return *o.ContainerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsScanSbomRequest) GetContainerNameOk() (*string, bool) {
	if o == nil || isNil(o.ContainerName) {
		return nil, false
	}
	return o.ContainerName, true
}

// HasContainerName returns a boolean if a field has been set.
func (o *UtilsScanSbomRequest) HasContainerName() bool {
	if o != nil && !isNil(o.ContainerName) {
		return true
	}

	return false
}

// SetContainerName gets a reference to the given string and assigns it to the ContainerName field.
func (o *UtilsScanSbomRequest) SetContainerName(v string) {
	o.ContainerName = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *UtilsScanSbomRequest) GetHostName() string {
	if o == nil || isNil(o.HostName) {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsScanSbomRequest) GetHostNameOk() (*string, bool) {
	if o == nil || isNil(o.HostName) {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *UtilsScanSbomRequest) HasHostName() bool {
	if o != nil && !isNil(o.HostName) {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *UtilsScanSbomRequest) SetHostName(v string) {
	o.HostName = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *UtilsScanSbomRequest) GetImageId() string {
	if o == nil || isNil(o.ImageId) {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsScanSbomRequest) GetImageIdOk() (*string, bool) {
	if o == nil || isNil(o.ImageId) {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *UtilsScanSbomRequest) HasImageId() bool {
	if o != nil && !isNil(o.ImageId) {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *UtilsScanSbomRequest) SetImageId(v string) {
	o.ImageId = &v
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *UtilsScanSbomRequest) GetImageName() string {
	if o == nil || isNil(o.ImageName) {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsScanSbomRequest) GetImageNameOk() (*string, bool) {
	if o == nil || isNil(o.ImageName) {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *UtilsScanSbomRequest) HasImageName() bool {
	if o != nil && !isNil(o.ImageName) {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *UtilsScanSbomRequest) SetImageName(v string) {
	o.ImageName = &v
}

// GetKubernetesClusterName returns the KubernetesClusterName field value if set, zero value otherwise.
func (o *UtilsScanSbomRequest) GetKubernetesClusterName() string {
	if o == nil || isNil(o.KubernetesClusterName) {
		var ret string
		return ret
	}
	return *o.KubernetesClusterName
}

// GetKubernetesClusterNameOk returns a tuple with the KubernetesClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsScanSbomRequest) GetKubernetesClusterNameOk() (*string, bool) {
	if o == nil || isNil(o.KubernetesClusterName) {
		return nil, false
	}
	return o.KubernetesClusterName, true
}

// HasKubernetesClusterName returns a boolean if a field has been set.
func (o *UtilsScanSbomRequest) HasKubernetesClusterName() bool {
	if o != nil && !isNil(o.KubernetesClusterName) {
		return true
	}

	return false
}

// SetKubernetesClusterName gets a reference to the given string and assigns it to the KubernetesClusterName field.
func (o *UtilsScanSbomRequest) SetKubernetesClusterName(v string) {
	o.KubernetesClusterName = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *UtilsScanSbomRequest) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsScanSbomRequest) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *UtilsScanSbomRequest) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *UtilsScanSbomRequest) SetMode(v string) {
	o.Mode = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *UtilsScanSbomRequest) GetNodeId() string {
	if o == nil || isNil(o.NodeId) {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsScanSbomRequest) GetNodeIdOk() (*string, bool) {
	if o == nil || isNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *UtilsScanSbomRequest) HasNodeId() bool {
	if o != nil && !isNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *UtilsScanSbomRequest) SetNodeId(v string) {
	o.NodeId = &v
}

// GetNodeType returns the NodeType field value if set, zero value otherwise.
func (o *UtilsScanSbomRequest) GetNodeType() string {
	if o == nil || isNil(o.NodeType) {
		var ret string
		return ret
	}
	return *o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsScanSbomRequest) GetNodeTypeOk() (*string, bool) {
	if o == nil || isNil(o.NodeType) {
		return nil, false
	}
	return o.NodeType, true
}

// HasNodeType returns a boolean if a field has been set.
func (o *UtilsScanSbomRequest) HasNodeType() bool {
	if o != nil && !isNil(o.NodeType) {
		return true
	}

	return false
}

// SetNodeType gets a reference to the given string and assigns it to the NodeType field.
func (o *UtilsScanSbomRequest) SetNodeType(v string) {
	o.NodeType = &v
}

// GetRegistryId returns the RegistryId field value if set, zero value otherwise.
func (o *UtilsScanSbomRequest) GetRegistryId() string {
	if o == nil || isNil(o.RegistryId) {
		var ret string
		return ret
	}
	return *o.RegistryId
}

// GetRegistryIdOk returns a tuple with the RegistryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsScanSbomRequest) GetRegistryIdOk() (*string, bool) {
	if o == nil || isNil(o.RegistryId) {
		return nil, false
	}
	return o.RegistryId, true
}

// HasRegistryId returns a boolean if a field has been set.
func (o *UtilsScanSbomRequest) HasRegistryId() bool {
	if o != nil && !isNil(o.RegistryId) {
		return true
	}

	return false
}

// SetRegistryId gets a reference to the given string and assigns it to the RegistryId field.
func (o *UtilsScanSbomRequest) SetRegistryId(v string) {
	o.RegistryId = &v
}

// GetSbom returns the Sbom field value
func (o *UtilsScanSbomRequest) GetSbom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sbom
}

// GetSbomOk returns a tuple with the Sbom field value
// and a boolean to check if the value has been set.
func (o *UtilsScanSbomRequest) GetSbomOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sbom, true
}

// SetSbom sets field value
func (o *UtilsScanSbomRequest) SetSbom(v string) {
	o.Sbom = v
}

// GetSbomFilePath returns the SbomFilePath field value if set, zero value otherwise.
func (o *UtilsScanSbomRequest) GetSbomFilePath() string {
	if o == nil || isNil(o.SbomFilePath) {
		var ret string
		return ret
	}
	return *o.SbomFilePath
}

// GetSbomFilePathOk returns a tuple with the SbomFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsScanSbomRequest) GetSbomFilePathOk() (*string, bool) {
	if o == nil || isNil(o.SbomFilePath) {
		return nil, false
	}
	return o.SbomFilePath, true
}

// HasSbomFilePath returns a boolean if a field has been set.
func (o *UtilsScanSbomRequest) HasSbomFilePath() bool {
	if o != nil && !isNil(o.SbomFilePath) {
		return true
	}

	return false
}

// SetSbomFilePath gets a reference to the given string and assigns it to the SbomFilePath field.
func (o *UtilsScanSbomRequest) SetSbomFilePath(v string) {
	o.SbomFilePath = &v
}

// GetScanId returns the ScanId field value
func (o *UtilsScanSbomRequest) GetScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value
// and a boolean to check if the value has been set.
func (o *UtilsScanSbomRequest) GetScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScanId, true
}

// SetScanId sets field value
func (o *UtilsScanSbomRequest) SetScanId(v string) {
	o.ScanId = v
}

// GetScanType returns the ScanType field value if set, zero value otherwise.
func (o *UtilsScanSbomRequest) GetScanType() string {
	if o == nil || isNil(o.ScanType) {
		var ret string
		return ret
	}
	return *o.ScanType
}

// GetScanTypeOk returns a tuple with the ScanType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsScanSbomRequest) GetScanTypeOk() (*string, bool) {
	if o == nil || isNil(o.ScanType) {
		return nil, false
	}
	return o.ScanType, true
}

// HasScanType returns a boolean if a field has been set.
func (o *UtilsScanSbomRequest) HasScanType() bool {
	if o != nil && !isNil(o.ScanType) {
		return true
	}

	return false
}

// SetScanType gets a reference to the given string and assigns it to the ScanType field.
func (o *UtilsScanSbomRequest) SetScanType(v string) {
	o.ScanType = &v
}

func (o UtilsScanSbomRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UtilsScanSbomRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ContainerName) {
		toSerialize["container_name"] = o.ContainerName
	}
	if !isNil(o.HostName) {
		toSerialize["host_name"] = o.HostName
	}
	if !isNil(o.ImageId) {
		toSerialize["image_id"] = o.ImageId
	}
	if !isNil(o.ImageName) {
		toSerialize["image_name"] = o.ImageName
	}
	if !isNil(o.KubernetesClusterName) {
		toSerialize["kubernetes_cluster_name"] = o.KubernetesClusterName
	}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !isNil(o.NodeId) {
		toSerialize["node_id"] = o.NodeId
	}
	if !isNil(o.NodeType) {
		toSerialize["node_type"] = o.NodeType
	}
	if !isNil(o.RegistryId) {
		toSerialize["registry_id"] = o.RegistryId
	}
	toSerialize["sbom"] = o.Sbom
	if !isNil(o.SbomFilePath) {
		toSerialize["sbom_file_path"] = o.SbomFilePath
	}
	toSerialize["scan_id"] = o.ScanId
	if !isNil(o.ScanType) {
		toSerialize["scan_type"] = o.ScanType
	}
	return toSerialize, nil
}

type NullableUtilsScanSbomRequest struct {
	value *UtilsScanSbomRequest
	isSet bool
}

func (v NullableUtilsScanSbomRequest) Get() *UtilsScanSbomRequest {
	return v.value
}

func (v *NullableUtilsScanSbomRequest) Set(val *UtilsScanSbomRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUtilsScanSbomRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUtilsScanSbomRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtilsScanSbomRequest(val *UtilsScanSbomRequest) *NullableUtilsScanSbomRequest {
	return &NullableUtilsScanSbomRequest{value: val, isSet: true}
}

func (v NullableUtilsScanSbomRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtilsScanSbomRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


