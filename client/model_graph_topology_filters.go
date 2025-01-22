/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.5.3
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GraphTopologyFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphTopologyFilters{}

// GraphTopologyFilters struct for GraphTopologyFilters
type GraphTopologyFilters struct {
	CloudFilter []string `json:"cloud_filter"`
	ContainerFilter []string `json:"container_filter"`
	FieldFilters ReportersFieldsFilters `json:"field_filters"`
	HostFilter []string `json:"host_filter"`
	KubernetesFilter []string `json:"kubernetes_filter"`
	PodFilter []string `json:"pod_filter"`
	RegionFilter []string `json:"region_filter"`
	SkipConnections bool `json:"skip_connections"`
}

type _GraphTopologyFilters GraphTopologyFilters

// NewGraphTopologyFilters instantiates a new GraphTopologyFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphTopologyFilters(cloudFilter []string, containerFilter []string, fieldFilters ReportersFieldsFilters, hostFilter []string, kubernetesFilter []string, podFilter []string, regionFilter []string, skipConnections bool) *GraphTopologyFilters {
	this := GraphTopologyFilters{}
	this.CloudFilter = cloudFilter
	this.ContainerFilter = containerFilter
	this.FieldFilters = fieldFilters
	this.HostFilter = hostFilter
	this.KubernetesFilter = kubernetesFilter
	this.PodFilter = podFilter
	this.RegionFilter = regionFilter
	this.SkipConnections = skipConnections
	return &this
}

// NewGraphTopologyFiltersWithDefaults instantiates a new GraphTopologyFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphTopologyFiltersWithDefaults() *GraphTopologyFilters {
	this := GraphTopologyFilters{}
	return &this
}

// GetCloudFilter returns the CloudFilter field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *GraphTopologyFilters) GetCloudFilter() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CloudFilter
}

// GetCloudFilterOk returns a tuple with the CloudFilter field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphTopologyFilters) GetCloudFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.CloudFilter) {
		return nil, false
	}
	return o.CloudFilter, true
}

// SetCloudFilter sets field value
func (o *GraphTopologyFilters) SetCloudFilter(v []string) {
	o.CloudFilter = v
}

// GetContainerFilter returns the ContainerFilter field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *GraphTopologyFilters) GetContainerFilter() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ContainerFilter
}

// GetContainerFilterOk returns a tuple with the ContainerFilter field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphTopologyFilters) GetContainerFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.ContainerFilter) {
		return nil, false
	}
	return o.ContainerFilter, true
}

// SetContainerFilter sets field value
func (o *GraphTopologyFilters) SetContainerFilter(v []string) {
	o.ContainerFilter = v
}

// GetFieldFilters returns the FieldFilters field value
func (o *GraphTopologyFilters) GetFieldFilters() ReportersFieldsFilters {
	if o == nil {
		var ret ReportersFieldsFilters
		return ret
	}

	return o.FieldFilters
}

// GetFieldFiltersOk returns a tuple with the FieldFilters field value
// and a boolean to check if the value has been set.
func (o *GraphTopologyFilters) GetFieldFiltersOk() (*ReportersFieldsFilters, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldFilters, true
}

// SetFieldFilters sets field value
func (o *GraphTopologyFilters) SetFieldFilters(v ReportersFieldsFilters) {
	o.FieldFilters = v
}

// GetHostFilter returns the HostFilter field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *GraphTopologyFilters) GetHostFilter() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.HostFilter
}

// GetHostFilterOk returns a tuple with the HostFilter field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphTopologyFilters) GetHostFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.HostFilter) {
		return nil, false
	}
	return o.HostFilter, true
}

// SetHostFilter sets field value
func (o *GraphTopologyFilters) SetHostFilter(v []string) {
	o.HostFilter = v
}

// GetKubernetesFilter returns the KubernetesFilter field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *GraphTopologyFilters) GetKubernetesFilter() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.KubernetesFilter
}

// GetKubernetesFilterOk returns a tuple with the KubernetesFilter field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphTopologyFilters) GetKubernetesFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.KubernetesFilter) {
		return nil, false
	}
	return o.KubernetesFilter, true
}

// SetKubernetesFilter sets field value
func (o *GraphTopologyFilters) SetKubernetesFilter(v []string) {
	o.KubernetesFilter = v
}

// GetPodFilter returns the PodFilter field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *GraphTopologyFilters) GetPodFilter() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PodFilter
}

// GetPodFilterOk returns a tuple with the PodFilter field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphTopologyFilters) GetPodFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.PodFilter) {
		return nil, false
	}
	return o.PodFilter, true
}

// SetPodFilter sets field value
func (o *GraphTopologyFilters) SetPodFilter(v []string) {
	o.PodFilter = v
}

// GetRegionFilter returns the RegionFilter field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *GraphTopologyFilters) GetRegionFilter() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RegionFilter
}

// GetRegionFilterOk returns a tuple with the RegionFilter field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphTopologyFilters) GetRegionFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.RegionFilter) {
		return nil, false
	}
	return o.RegionFilter, true
}

// SetRegionFilter sets field value
func (o *GraphTopologyFilters) SetRegionFilter(v []string) {
	o.RegionFilter = v
}

// GetSkipConnections returns the SkipConnections field value
func (o *GraphTopologyFilters) GetSkipConnections() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SkipConnections
}

// GetSkipConnectionsOk returns a tuple with the SkipConnections field value
// and a boolean to check if the value has been set.
func (o *GraphTopologyFilters) GetSkipConnectionsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkipConnections, true
}

// SetSkipConnections sets field value
func (o *GraphTopologyFilters) SetSkipConnections(v bool) {
	o.SkipConnections = v
}

func (o GraphTopologyFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphTopologyFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CloudFilter != nil {
		toSerialize["cloud_filter"] = o.CloudFilter
	}
	if o.ContainerFilter != nil {
		toSerialize["container_filter"] = o.ContainerFilter
	}
	toSerialize["field_filters"] = o.FieldFilters
	if o.HostFilter != nil {
		toSerialize["host_filter"] = o.HostFilter
	}
	if o.KubernetesFilter != nil {
		toSerialize["kubernetes_filter"] = o.KubernetesFilter
	}
	if o.PodFilter != nil {
		toSerialize["pod_filter"] = o.PodFilter
	}
	if o.RegionFilter != nil {
		toSerialize["region_filter"] = o.RegionFilter
	}
	toSerialize["skip_connections"] = o.SkipConnections
	return toSerialize, nil
}

func (o *GraphTopologyFilters) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloud_filter",
		"container_filter",
		"field_filters",
		"host_filter",
		"kubernetes_filter",
		"pod_filter",
		"region_filter",
		"skip_connections",
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

	varGraphTopologyFilters := _GraphTopologyFilters{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGraphTopologyFilters)

	if err != nil {
		return err
	}

	*o = GraphTopologyFilters(varGraphTopologyFilters)

	return err
}

type NullableGraphTopologyFilters struct {
	value *GraphTopologyFilters
	isSet bool
}

func (v NullableGraphTopologyFilters) Get() *GraphTopologyFilters {
	return v.value
}

func (v *NullableGraphTopologyFilters) Set(val *GraphTopologyFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphTopologyFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphTopologyFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphTopologyFilters(val *GraphTopologyFilters) *NullableGraphTopologyFilters {
	return &NullableGraphTopologyFilters{value: val, isSet: true}
}

func (v NullableGraphTopologyFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphTopologyFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


