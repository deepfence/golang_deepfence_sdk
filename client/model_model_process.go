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

// checks if the ModelProcess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelProcess{}

// ModelProcess struct for ModelProcess
type ModelProcess struct {
	Cmdline string `json:"cmdline"`
	Metadata map[string]interface{} `json:"metadata"`
	Metrics ModelComputeMetrics `json:"metrics"`
	Name string `json:"name"`
	NodeId string `json:"node_id"`
	Pid string `json:"pid"`
	Ppid string `json:"ppid"`
	Threads int32 `json:"threads"`
}

// NewModelProcess instantiates a new ModelProcess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelProcess(cmdline string, metadata map[string]interface{}, metrics ModelComputeMetrics, name string, nodeId string, pid string, ppid string, threads int32) *ModelProcess {
	this := ModelProcess{}
	this.Cmdline = cmdline
	this.Metadata = metadata
	this.Metrics = metrics
	this.Name = name
	this.NodeId = nodeId
	this.Pid = pid
	this.Ppid = ppid
	this.Threads = threads
	return &this
}

// NewModelProcessWithDefaults instantiates a new ModelProcess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelProcessWithDefaults() *ModelProcess {
	this := ModelProcess{}
	return &this
}

// GetCmdline returns the Cmdline field value
func (o *ModelProcess) GetCmdline() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cmdline
}

// GetCmdlineOk returns a tuple with the Cmdline field value
// and a boolean to check if the value has been set.
func (o *ModelProcess) GetCmdlineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cmdline, true
}

// SetCmdline sets field value
func (o *ModelProcess) SetCmdline(v string) {
	o.Cmdline = v
}

// GetMetadata returns the Metadata field value
func (o *ModelProcess) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ModelProcess) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *ModelProcess) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMetrics returns the Metrics field value
func (o *ModelProcess) GetMetrics() ModelComputeMetrics {
	if o == nil {
		var ret ModelComputeMetrics
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *ModelProcess) GetMetricsOk() (*ModelComputeMetrics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// SetMetrics sets field value
func (o *ModelProcess) SetMetrics(v ModelComputeMetrics) {
	o.Metrics = v
}

// GetName returns the Name field value
func (o *ModelProcess) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModelProcess) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ModelProcess) SetName(v string) {
	o.Name = v
}

// GetNodeId returns the NodeId field value
func (o *ModelProcess) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelProcess) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelProcess) SetNodeId(v string) {
	o.NodeId = v
}

// GetPid returns the Pid field value
func (o *ModelProcess) GetPid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pid
}

// GetPidOk returns a tuple with the Pid field value
// and a boolean to check if the value has been set.
func (o *ModelProcess) GetPidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pid, true
}

// SetPid sets field value
func (o *ModelProcess) SetPid(v string) {
	o.Pid = v
}

// GetPpid returns the Ppid field value
func (o *ModelProcess) GetPpid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ppid
}

// GetPpidOk returns a tuple with the Ppid field value
// and a boolean to check if the value has been set.
func (o *ModelProcess) GetPpidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ppid, true
}

// SetPpid sets field value
func (o *ModelProcess) SetPpid(v string) {
	o.Ppid = v
}

// GetThreads returns the Threads field value
func (o *ModelProcess) GetThreads() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Threads
}

// GetThreadsOk returns a tuple with the Threads field value
// and a boolean to check if the value has been set.
func (o *ModelProcess) GetThreadsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threads, true
}

// SetThreads sets field value
func (o *ModelProcess) SetThreads(v int32) {
	o.Threads = v
}

func (o ModelProcess) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelProcess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cmdline"] = o.Cmdline
	toSerialize["metadata"] = o.Metadata
	toSerialize["metrics"] = o.Metrics
	toSerialize["name"] = o.Name
	toSerialize["node_id"] = o.NodeId
	toSerialize["pid"] = o.Pid
	toSerialize["ppid"] = o.Ppid
	toSerialize["threads"] = o.Threads
	return toSerialize, nil
}

type NullableModelProcess struct {
	value *ModelProcess
	isSet bool
}

func (v NullableModelProcess) Get() *ModelProcess {
	return v.value
}

func (v *NullableModelProcess) Set(val *ModelProcess) {
	v.value = val
	v.isSet = true
}

func (v NullableModelProcess) IsSet() bool {
	return v.isSet
}

func (v *NullableModelProcess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelProcess(val *ModelProcess) *NullableModelProcess {
	return &NullableModelProcess{value: val, isSet: true}
}

func (v NullableModelProcess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelProcess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


