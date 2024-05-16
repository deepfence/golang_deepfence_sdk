/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.2.1
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SearchNodeCountResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchNodeCountResp{}

// SearchNodeCountResp struct for SearchNodeCountResp
type SearchNodeCountResp struct {
	CloudProvider int32 `json:"cloud_provider"`
	Container int32 `json:"container"`
	ContainerImage int32 `json:"container_image"`
	Host int32 `json:"host"`
	KubernetesCluster int32 `json:"kubernetes_cluster"`
	Namespace int32 `json:"namespace"`
	Pod int32 `json:"pod"`
}

type _SearchNodeCountResp SearchNodeCountResp

// NewSearchNodeCountResp instantiates a new SearchNodeCountResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchNodeCountResp(cloudProvider int32, container int32, containerImage int32, host int32, kubernetesCluster int32, namespace int32, pod int32) *SearchNodeCountResp {
	this := SearchNodeCountResp{}
	this.CloudProvider = cloudProvider
	this.Container = container
	this.ContainerImage = containerImage
	this.Host = host
	this.KubernetesCluster = kubernetesCluster
	this.Namespace = namespace
	this.Pod = pod
	return &this
}

// NewSearchNodeCountRespWithDefaults instantiates a new SearchNodeCountResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchNodeCountRespWithDefaults() *SearchNodeCountResp {
	this := SearchNodeCountResp{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value
func (o *SearchNodeCountResp) GetCloudProvider() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *SearchNodeCountResp) GetCloudProviderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *SearchNodeCountResp) SetCloudProvider(v int32) {
	o.CloudProvider = v
}

// GetContainer returns the Container field value
func (o *SearchNodeCountResp) GetContainer() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Container
}

// GetContainerOk returns a tuple with the Container field value
// and a boolean to check if the value has been set.
func (o *SearchNodeCountResp) GetContainerOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Container, true
}

// SetContainer sets field value
func (o *SearchNodeCountResp) SetContainer(v int32) {
	o.Container = v
}

// GetContainerImage returns the ContainerImage field value
func (o *SearchNodeCountResp) GetContainerImage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ContainerImage
}

// GetContainerImageOk returns a tuple with the ContainerImage field value
// and a boolean to check if the value has been set.
func (o *SearchNodeCountResp) GetContainerImageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerImage, true
}

// SetContainerImage sets field value
func (o *SearchNodeCountResp) SetContainerImage(v int32) {
	o.ContainerImage = v
}

// GetHost returns the Host field value
func (o *SearchNodeCountResp) GetHost() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *SearchNodeCountResp) GetHostOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *SearchNodeCountResp) SetHost(v int32) {
	o.Host = v
}

// GetKubernetesCluster returns the KubernetesCluster field value
func (o *SearchNodeCountResp) GetKubernetesCluster() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.KubernetesCluster
}

// GetKubernetesClusterOk returns a tuple with the KubernetesCluster field value
// and a boolean to check if the value has been set.
func (o *SearchNodeCountResp) GetKubernetesClusterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubernetesCluster, true
}

// SetKubernetesCluster sets field value
func (o *SearchNodeCountResp) SetKubernetesCluster(v int32) {
	o.KubernetesCluster = v
}

// GetNamespace returns the Namespace field value
func (o *SearchNodeCountResp) GetNamespace() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *SearchNodeCountResp) GetNamespaceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *SearchNodeCountResp) SetNamespace(v int32) {
	o.Namespace = v
}

// GetPod returns the Pod field value
func (o *SearchNodeCountResp) GetPod() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pod
}

// GetPodOk returns a tuple with the Pod field value
// and a boolean to check if the value has been set.
func (o *SearchNodeCountResp) GetPodOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pod, true
}

// SetPod sets field value
func (o *SearchNodeCountResp) SetPod(v int32) {
	o.Pod = v
}

func (o SearchNodeCountResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchNodeCountResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloud_provider"] = o.CloudProvider
	toSerialize["container"] = o.Container
	toSerialize["container_image"] = o.ContainerImage
	toSerialize["host"] = o.Host
	toSerialize["kubernetes_cluster"] = o.KubernetesCluster
	toSerialize["namespace"] = o.Namespace
	toSerialize["pod"] = o.Pod
	return toSerialize, nil
}

func (o *SearchNodeCountResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloud_provider",
		"container",
		"container_image",
		"host",
		"kubernetes_cluster",
		"namespace",
		"pod",
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

	varSearchNodeCountResp := _SearchNodeCountResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchNodeCountResp)

	if err != nil {
		return err
	}

	*o = SearchNodeCountResp(varSearchNodeCountResp)

	return err
}

type NullableSearchNodeCountResp struct {
	value *SearchNodeCountResp
	isSet bool
}

func (v NullableSearchNodeCountResp) Get() *SearchNodeCountResp {
	return v.value
}

func (v *NullableSearchNodeCountResp) Set(val *SearchNodeCountResp) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchNodeCountResp) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchNodeCountResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchNodeCountResp(val *SearchNodeCountResp) *NullableSearchNodeCountResp {
	return &NullableSearchNodeCountResp{value: val, isSet: true}
}

func (v NullableSearchNodeCountResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchNodeCountResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


