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

// checks if the IngestersReportIngestionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestersReportIngestionData{}

// IngestersReportIngestionData struct for IngestersReportIngestionData
type IngestersReportIngestionData struct {
	ContainerBatch []map[string]string `json:"container_batch"`
	ContainerEdgesBatch []map[string]interface{} `json:"container_edges_batch"`
	ContainerImageBatch []map[string]string `json:"container_image_batch"`
	ContainerImageEdgeBatch []map[string]interface{} `json:"container_image_edge_batch"`
	EndpointEdgesBatch []map[string]interface{} `json:"endpoint_edges_batch"`
	HostBatch []map[string]string `json:"host_batch"`
	Hosts []map[string]string `json:"hosts"`
	KubernetesClusterBatch []map[string]string `json:"kubernetes_cluster_batch"`
	KubernetesClusterEdgeBatch []map[string]interface{} `json:"kubernetes_cluster_edge_batch"`
	PodBatch []map[string]string `json:"pod_batch"`
	PodEdgesBatch []map[string]interface{} `json:"pod_edges_batch"`
	ProcessBatch []map[string]string `json:"process_batch"`
	ProcessEdgesBatch []map[string]interface{} `json:"process_edges_batch"`
}

// NewIngestersReportIngestionData instantiates a new IngestersReportIngestionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestersReportIngestionData(containerBatch []map[string]string, containerEdgesBatch []map[string]interface{}, containerImageBatch []map[string]string, containerImageEdgeBatch []map[string]interface{}, endpointEdgesBatch []map[string]interface{}, hostBatch []map[string]string, hosts []map[string]string, kubernetesClusterBatch []map[string]string, kubernetesClusterEdgeBatch []map[string]interface{}, podBatch []map[string]string, podEdgesBatch []map[string]interface{}, processBatch []map[string]string, processEdgesBatch []map[string]interface{}) *IngestersReportIngestionData {
	this := IngestersReportIngestionData{}
	this.ContainerBatch = containerBatch
	this.ContainerEdgesBatch = containerEdgesBatch
	this.ContainerImageBatch = containerImageBatch
	this.ContainerImageEdgeBatch = containerImageEdgeBatch
	this.EndpointEdgesBatch = endpointEdgesBatch
	this.HostBatch = hostBatch
	this.Hosts = hosts
	this.KubernetesClusterBatch = kubernetesClusterBatch
	this.KubernetesClusterEdgeBatch = kubernetesClusterEdgeBatch
	this.PodBatch = podBatch
	this.PodEdgesBatch = podEdgesBatch
	this.ProcessBatch = processBatch
	this.ProcessEdgesBatch = processEdgesBatch
	return &this
}

// NewIngestersReportIngestionDataWithDefaults instantiates a new IngestersReportIngestionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestersReportIngestionDataWithDefaults() *IngestersReportIngestionData {
	this := IngestersReportIngestionData{}
	return &this
}

// GetContainerBatch returns the ContainerBatch field value
// If the value is explicit nil, the zero value for []map[string]string will be returned
func (o *IngestersReportIngestionData) GetContainerBatch() []map[string]string {
	if o == nil {
		var ret []map[string]string
		return ret
	}

	return o.ContainerBatch
}

// GetContainerBatchOk returns a tuple with the ContainerBatch field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngestersReportIngestionData) GetContainerBatchOk() ([]map[string]string, bool) {
	if o == nil || IsNil(o.ContainerBatch) {
		return nil, false
	}
	return o.ContainerBatch, true
}

// SetContainerBatch sets field value
func (o *IngestersReportIngestionData) SetContainerBatch(v []map[string]string) {
	o.ContainerBatch = v
}

// GetContainerEdgesBatch returns the ContainerEdgesBatch field value
// If the value is explicit nil, the zero value for []map[string]interface{} will be returned
func (o *IngestersReportIngestionData) GetContainerEdgesBatch() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.ContainerEdgesBatch
}

// GetContainerEdgesBatchOk returns a tuple with the ContainerEdgesBatch field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngestersReportIngestionData) GetContainerEdgesBatchOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.ContainerEdgesBatch) {
		return nil, false
	}
	return o.ContainerEdgesBatch, true
}

// SetContainerEdgesBatch sets field value
func (o *IngestersReportIngestionData) SetContainerEdgesBatch(v []map[string]interface{}) {
	o.ContainerEdgesBatch = v
}

// GetContainerImageBatch returns the ContainerImageBatch field value
// If the value is explicit nil, the zero value for []map[string]string will be returned
func (o *IngestersReportIngestionData) GetContainerImageBatch() []map[string]string {
	if o == nil {
		var ret []map[string]string
		return ret
	}

	return o.ContainerImageBatch
}

// GetContainerImageBatchOk returns a tuple with the ContainerImageBatch field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngestersReportIngestionData) GetContainerImageBatchOk() ([]map[string]string, bool) {
	if o == nil || IsNil(o.ContainerImageBatch) {
		return nil, false
	}
	return o.ContainerImageBatch, true
}

// SetContainerImageBatch sets field value
func (o *IngestersReportIngestionData) SetContainerImageBatch(v []map[string]string) {
	o.ContainerImageBatch = v
}

// GetContainerImageEdgeBatch returns the ContainerImageEdgeBatch field value
// If the value is explicit nil, the zero value for []map[string]interface{} will be returned
func (o *IngestersReportIngestionData) GetContainerImageEdgeBatch() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.ContainerImageEdgeBatch
}

// GetContainerImageEdgeBatchOk returns a tuple with the ContainerImageEdgeBatch field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngestersReportIngestionData) GetContainerImageEdgeBatchOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.ContainerImageEdgeBatch) {
		return nil, false
	}
	return o.ContainerImageEdgeBatch, true
}

// SetContainerImageEdgeBatch sets field value
func (o *IngestersReportIngestionData) SetContainerImageEdgeBatch(v []map[string]interface{}) {
	o.ContainerImageEdgeBatch = v
}

// GetEndpointEdgesBatch returns the EndpointEdgesBatch field value
// If the value is explicit nil, the zero value for []map[string]interface{} will be returned
func (o *IngestersReportIngestionData) GetEndpointEdgesBatch() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.EndpointEdgesBatch
}

// GetEndpointEdgesBatchOk returns a tuple with the EndpointEdgesBatch field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngestersReportIngestionData) GetEndpointEdgesBatchOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.EndpointEdgesBatch) {
		return nil, false
	}
	return o.EndpointEdgesBatch, true
}

// SetEndpointEdgesBatch sets field value
func (o *IngestersReportIngestionData) SetEndpointEdgesBatch(v []map[string]interface{}) {
	o.EndpointEdgesBatch = v
}

// GetHostBatch returns the HostBatch field value
// If the value is explicit nil, the zero value for []map[string]string will be returned
func (o *IngestersReportIngestionData) GetHostBatch() []map[string]string {
	if o == nil {
		var ret []map[string]string
		return ret
	}

	return o.HostBatch
}

// GetHostBatchOk returns a tuple with the HostBatch field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngestersReportIngestionData) GetHostBatchOk() ([]map[string]string, bool) {
	if o == nil || IsNil(o.HostBatch) {
		return nil, false
	}
	return o.HostBatch, true
}

// SetHostBatch sets field value
func (o *IngestersReportIngestionData) SetHostBatch(v []map[string]string) {
	o.HostBatch = v
}

// GetHosts returns the Hosts field value
// If the value is explicit nil, the zero value for []map[string]string will be returned
func (o *IngestersReportIngestionData) GetHosts() []map[string]string {
	if o == nil {
		var ret []map[string]string
		return ret
	}

	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngestersReportIngestionData) GetHostsOk() ([]map[string]string, bool) {
	if o == nil || IsNil(o.Hosts) {
		return nil, false
	}
	return o.Hosts, true
}

// SetHosts sets field value
func (o *IngestersReportIngestionData) SetHosts(v []map[string]string) {
	o.Hosts = v
}

// GetKubernetesClusterBatch returns the KubernetesClusterBatch field value
// If the value is explicit nil, the zero value for []map[string]string will be returned
func (o *IngestersReportIngestionData) GetKubernetesClusterBatch() []map[string]string {
	if o == nil {
		var ret []map[string]string
		return ret
	}

	return o.KubernetesClusterBatch
}

// GetKubernetesClusterBatchOk returns a tuple with the KubernetesClusterBatch field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngestersReportIngestionData) GetKubernetesClusterBatchOk() ([]map[string]string, bool) {
	if o == nil || IsNil(o.KubernetesClusterBatch) {
		return nil, false
	}
	return o.KubernetesClusterBatch, true
}

// SetKubernetesClusterBatch sets field value
func (o *IngestersReportIngestionData) SetKubernetesClusterBatch(v []map[string]string) {
	o.KubernetesClusterBatch = v
}

// GetKubernetesClusterEdgeBatch returns the KubernetesClusterEdgeBatch field value
// If the value is explicit nil, the zero value for []map[string]interface{} will be returned
func (o *IngestersReportIngestionData) GetKubernetesClusterEdgeBatch() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.KubernetesClusterEdgeBatch
}

// GetKubernetesClusterEdgeBatchOk returns a tuple with the KubernetesClusterEdgeBatch field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngestersReportIngestionData) GetKubernetesClusterEdgeBatchOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.KubernetesClusterEdgeBatch) {
		return nil, false
	}
	return o.KubernetesClusterEdgeBatch, true
}

// SetKubernetesClusterEdgeBatch sets field value
func (o *IngestersReportIngestionData) SetKubernetesClusterEdgeBatch(v []map[string]interface{}) {
	o.KubernetesClusterEdgeBatch = v
}

// GetPodBatch returns the PodBatch field value
// If the value is explicit nil, the zero value for []map[string]string will be returned
func (o *IngestersReportIngestionData) GetPodBatch() []map[string]string {
	if o == nil {
		var ret []map[string]string
		return ret
	}

	return o.PodBatch
}

// GetPodBatchOk returns a tuple with the PodBatch field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngestersReportIngestionData) GetPodBatchOk() ([]map[string]string, bool) {
	if o == nil || IsNil(o.PodBatch) {
		return nil, false
	}
	return o.PodBatch, true
}

// SetPodBatch sets field value
func (o *IngestersReportIngestionData) SetPodBatch(v []map[string]string) {
	o.PodBatch = v
}

// GetPodEdgesBatch returns the PodEdgesBatch field value
// If the value is explicit nil, the zero value for []map[string]interface{} will be returned
func (o *IngestersReportIngestionData) GetPodEdgesBatch() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.PodEdgesBatch
}

// GetPodEdgesBatchOk returns a tuple with the PodEdgesBatch field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngestersReportIngestionData) GetPodEdgesBatchOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.PodEdgesBatch) {
		return nil, false
	}
	return o.PodEdgesBatch, true
}

// SetPodEdgesBatch sets field value
func (o *IngestersReportIngestionData) SetPodEdgesBatch(v []map[string]interface{}) {
	o.PodEdgesBatch = v
}

// GetProcessBatch returns the ProcessBatch field value
// If the value is explicit nil, the zero value for []map[string]string will be returned
func (o *IngestersReportIngestionData) GetProcessBatch() []map[string]string {
	if o == nil {
		var ret []map[string]string
		return ret
	}

	return o.ProcessBatch
}

// GetProcessBatchOk returns a tuple with the ProcessBatch field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngestersReportIngestionData) GetProcessBatchOk() ([]map[string]string, bool) {
	if o == nil || IsNil(o.ProcessBatch) {
		return nil, false
	}
	return o.ProcessBatch, true
}

// SetProcessBatch sets field value
func (o *IngestersReportIngestionData) SetProcessBatch(v []map[string]string) {
	o.ProcessBatch = v
}

// GetProcessEdgesBatch returns the ProcessEdgesBatch field value
// If the value is explicit nil, the zero value for []map[string]interface{} will be returned
func (o *IngestersReportIngestionData) GetProcessEdgesBatch() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.ProcessEdgesBatch
}

// GetProcessEdgesBatchOk returns a tuple with the ProcessEdgesBatch field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngestersReportIngestionData) GetProcessEdgesBatchOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.ProcessEdgesBatch) {
		return nil, false
	}
	return o.ProcessEdgesBatch, true
}

// SetProcessEdgesBatch sets field value
func (o *IngestersReportIngestionData) SetProcessEdgesBatch(v []map[string]interface{}) {
	o.ProcessEdgesBatch = v
}

func (o IngestersReportIngestionData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestersReportIngestionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ContainerBatch != nil {
		toSerialize["container_batch"] = o.ContainerBatch
	}
	if o.ContainerEdgesBatch != nil {
		toSerialize["container_edges_batch"] = o.ContainerEdgesBatch
	}
	if o.ContainerImageBatch != nil {
		toSerialize["container_image_batch"] = o.ContainerImageBatch
	}
	if o.ContainerImageEdgeBatch != nil {
		toSerialize["container_image_edge_batch"] = o.ContainerImageEdgeBatch
	}
	if o.EndpointEdgesBatch != nil {
		toSerialize["endpoint_edges_batch"] = o.EndpointEdgesBatch
	}
	if o.HostBatch != nil {
		toSerialize["host_batch"] = o.HostBatch
	}
	if o.Hosts != nil {
		toSerialize["hosts"] = o.Hosts
	}
	if o.KubernetesClusterBatch != nil {
		toSerialize["kubernetes_cluster_batch"] = o.KubernetesClusterBatch
	}
	if o.KubernetesClusterEdgeBatch != nil {
		toSerialize["kubernetes_cluster_edge_batch"] = o.KubernetesClusterEdgeBatch
	}
	if o.PodBatch != nil {
		toSerialize["pod_batch"] = o.PodBatch
	}
	if o.PodEdgesBatch != nil {
		toSerialize["pod_edges_batch"] = o.PodEdgesBatch
	}
	if o.ProcessBatch != nil {
		toSerialize["process_batch"] = o.ProcessBatch
	}
	if o.ProcessEdgesBatch != nil {
		toSerialize["process_edges_batch"] = o.ProcessEdgesBatch
	}
	return toSerialize, nil
}

type NullableIngestersReportIngestionData struct {
	value *IngestersReportIngestionData
	isSet bool
}

func (v NullableIngestersReportIngestionData) Get() *IngestersReportIngestionData {
	return v.value
}

func (v *NullableIngestersReportIngestionData) Set(val *IngestersReportIngestionData) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestersReportIngestionData) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestersReportIngestionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestersReportIngestionData(val *IngestersReportIngestionData) *NullableIngestersReportIngestionData {
	return &NullableIngestersReportIngestionData{value: val, isSet: true}
}

func (v NullableIngestersReportIngestionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestersReportIngestionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


