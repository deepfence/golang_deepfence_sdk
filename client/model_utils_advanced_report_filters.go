/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.5.5
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UtilsAdvancedReportFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UtilsAdvancedReportFilters{}

// UtilsAdvancedReportFilters struct for UtilsAdvancedReportFilters
type UtilsAdvancedReportFilters struct {
	ContainerName []string `json:"container_name,omitempty"`
	HostName []string `json:"host_name,omitempty"`
	ImageName []string `json:"image_name,omitempty"`
	KubernetesClusterName []string `json:"kubernetes_cluster_name,omitempty"`
	Masked []bool `json:"masked,omitempty"`
	NodeId []string `json:"node_id,omitempty"`
	PodName []string `json:"pod_name,omitempty"`
	ScanStatus []string `json:"scan_status,omitempty"`
}

// NewUtilsAdvancedReportFilters instantiates a new UtilsAdvancedReportFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUtilsAdvancedReportFilters() *UtilsAdvancedReportFilters {
	this := UtilsAdvancedReportFilters{}
	return &this
}

// NewUtilsAdvancedReportFiltersWithDefaults instantiates a new UtilsAdvancedReportFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUtilsAdvancedReportFiltersWithDefaults() *UtilsAdvancedReportFilters {
	this := UtilsAdvancedReportFilters{}
	return &this
}

// GetContainerName returns the ContainerName field value if set, zero value otherwise.
func (o *UtilsAdvancedReportFilters) GetContainerName() []string {
	if o == nil || IsNil(o.ContainerName) {
		var ret []string
		return ret
	}
	return o.ContainerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsAdvancedReportFilters) GetContainerNameOk() ([]string, bool) {
	if o == nil || IsNil(o.ContainerName) {
		return nil, false
	}
	return o.ContainerName, true
}

// HasContainerName returns a boolean if a field has been set.
func (o *UtilsAdvancedReportFilters) HasContainerName() bool {
	if o != nil && !IsNil(o.ContainerName) {
		return true
	}

	return false
}

// SetContainerName gets a reference to the given []string and assigns it to the ContainerName field.
func (o *UtilsAdvancedReportFilters) SetContainerName(v []string) {
	o.ContainerName = v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *UtilsAdvancedReportFilters) GetHostName() []string {
	if o == nil || IsNil(o.HostName) {
		var ret []string
		return ret
	}
	return o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsAdvancedReportFilters) GetHostNameOk() ([]string, bool) {
	if o == nil || IsNil(o.HostName) {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *UtilsAdvancedReportFilters) HasHostName() bool {
	if o != nil && !IsNil(o.HostName) {
		return true
	}

	return false
}

// SetHostName gets a reference to the given []string and assigns it to the HostName field.
func (o *UtilsAdvancedReportFilters) SetHostName(v []string) {
	o.HostName = v
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *UtilsAdvancedReportFilters) GetImageName() []string {
	if o == nil || IsNil(o.ImageName) {
		var ret []string
		return ret
	}
	return o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsAdvancedReportFilters) GetImageNameOk() ([]string, bool) {
	if o == nil || IsNil(o.ImageName) {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *UtilsAdvancedReportFilters) HasImageName() bool {
	if o != nil && !IsNil(o.ImageName) {
		return true
	}

	return false
}

// SetImageName gets a reference to the given []string and assigns it to the ImageName field.
func (o *UtilsAdvancedReportFilters) SetImageName(v []string) {
	o.ImageName = v
}

// GetKubernetesClusterName returns the KubernetesClusterName field value if set, zero value otherwise.
func (o *UtilsAdvancedReportFilters) GetKubernetesClusterName() []string {
	if o == nil || IsNil(o.KubernetesClusterName) {
		var ret []string
		return ret
	}
	return o.KubernetesClusterName
}

// GetKubernetesClusterNameOk returns a tuple with the KubernetesClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsAdvancedReportFilters) GetKubernetesClusterNameOk() ([]string, bool) {
	if o == nil || IsNil(o.KubernetesClusterName) {
		return nil, false
	}
	return o.KubernetesClusterName, true
}

// HasKubernetesClusterName returns a boolean if a field has been set.
func (o *UtilsAdvancedReportFilters) HasKubernetesClusterName() bool {
	if o != nil && !IsNil(o.KubernetesClusterName) {
		return true
	}

	return false
}

// SetKubernetesClusterName gets a reference to the given []string and assigns it to the KubernetesClusterName field.
func (o *UtilsAdvancedReportFilters) SetKubernetesClusterName(v []string) {
	o.KubernetesClusterName = v
}

// GetMasked returns the Masked field value if set, zero value otherwise.
func (o *UtilsAdvancedReportFilters) GetMasked() []bool {
	if o == nil || IsNil(o.Masked) {
		var ret []bool
		return ret
	}
	return o.Masked
}

// GetMaskedOk returns a tuple with the Masked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsAdvancedReportFilters) GetMaskedOk() ([]bool, bool) {
	if o == nil || IsNil(o.Masked) {
		return nil, false
	}
	return o.Masked, true
}

// HasMasked returns a boolean if a field has been set.
func (o *UtilsAdvancedReportFilters) HasMasked() bool {
	if o != nil && !IsNil(o.Masked) {
		return true
	}

	return false
}

// SetMasked gets a reference to the given []bool and assigns it to the Masked field.
func (o *UtilsAdvancedReportFilters) SetMasked(v []bool) {
	o.Masked = v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *UtilsAdvancedReportFilters) GetNodeId() []string {
	if o == nil || IsNil(o.NodeId) {
		var ret []string
		return ret
	}
	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsAdvancedReportFilters) GetNodeIdOk() ([]string, bool) {
	if o == nil || IsNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *UtilsAdvancedReportFilters) HasNodeId() bool {
	if o != nil && !IsNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given []string and assigns it to the NodeId field.
func (o *UtilsAdvancedReportFilters) SetNodeId(v []string) {
	o.NodeId = v
}

// GetPodName returns the PodName field value if set, zero value otherwise.
func (o *UtilsAdvancedReportFilters) GetPodName() []string {
	if o == nil || IsNil(o.PodName) {
		var ret []string
		return ret
	}
	return o.PodName
}

// GetPodNameOk returns a tuple with the PodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsAdvancedReportFilters) GetPodNameOk() ([]string, bool) {
	if o == nil || IsNil(o.PodName) {
		return nil, false
	}
	return o.PodName, true
}

// HasPodName returns a boolean if a field has been set.
func (o *UtilsAdvancedReportFilters) HasPodName() bool {
	if o != nil && !IsNil(o.PodName) {
		return true
	}

	return false
}

// SetPodName gets a reference to the given []string and assigns it to the PodName field.
func (o *UtilsAdvancedReportFilters) SetPodName(v []string) {
	o.PodName = v
}

// GetScanStatus returns the ScanStatus field value if set, zero value otherwise.
func (o *UtilsAdvancedReportFilters) GetScanStatus() []string {
	if o == nil || IsNil(o.ScanStatus) {
		var ret []string
		return ret
	}
	return o.ScanStatus
}

// GetScanStatusOk returns a tuple with the ScanStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsAdvancedReportFilters) GetScanStatusOk() ([]string, bool) {
	if o == nil || IsNil(o.ScanStatus) {
		return nil, false
	}
	return o.ScanStatus, true
}

// HasScanStatus returns a boolean if a field has been set.
func (o *UtilsAdvancedReportFilters) HasScanStatus() bool {
	if o != nil && !IsNil(o.ScanStatus) {
		return true
	}

	return false
}

// SetScanStatus gets a reference to the given []string and assigns it to the ScanStatus field.
func (o *UtilsAdvancedReportFilters) SetScanStatus(v []string) {
	o.ScanStatus = v
}

func (o UtilsAdvancedReportFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UtilsAdvancedReportFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContainerName) {
		toSerialize["container_name"] = o.ContainerName
	}
	if !IsNil(o.HostName) {
		toSerialize["host_name"] = o.HostName
	}
	if !IsNil(o.ImageName) {
		toSerialize["image_name"] = o.ImageName
	}
	if !IsNil(o.KubernetesClusterName) {
		toSerialize["kubernetes_cluster_name"] = o.KubernetesClusterName
	}
	if !IsNil(o.Masked) {
		toSerialize["masked"] = o.Masked
	}
	if !IsNil(o.NodeId) {
		toSerialize["node_id"] = o.NodeId
	}
	if !IsNil(o.PodName) {
		toSerialize["pod_name"] = o.PodName
	}
	if !IsNil(o.ScanStatus) {
		toSerialize["scan_status"] = o.ScanStatus
	}
	return toSerialize, nil
}

type NullableUtilsAdvancedReportFilters struct {
	value *UtilsAdvancedReportFilters
	isSet bool
}

func (v NullableUtilsAdvancedReportFilters) Get() *UtilsAdvancedReportFilters {
	return v.value
}

func (v *NullableUtilsAdvancedReportFilters) Set(val *UtilsAdvancedReportFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableUtilsAdvancedReportFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableUtilsAdvancedReportFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtilsAdvancedReportFilters(val *UtilsAdvancedReportFilters) *NullableUtilsAdvancedReportFilters {
	return &NullableUtilsAdvancedReportFilters{value: val, isSet: true}
}

func (v NullableUtilsAdvancedReportFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtilsAdvancedReportFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


