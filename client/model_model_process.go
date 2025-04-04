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
	"bytes"
	"fmt"
)

// checks if the ModelProcess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelProcess{}

// ModelProcess struct for ModelProcess
type ModelProcess struct {
	Cmdline string `json:"cmdline"`
	CpuMax float32 `json:"cpu_max"`
	CpuUsage float32 `json:"cpu_usage"`
	MemoryMax int32 `json:"memory_max"`
	MemoryUsage int32 `json:"memory_usage"`
	NodeId string `json:"node_id"`
	NodeName string `json:"node_name"`
	OpenFilesCount int32 `json:"open_files_count"`
	Pid int32 `json:"pid"`
	Ppid int32 `json:"ppid"`
	ShortName string `json:"short_name"`
	Threads int32 `json:"threads"`
}

type _ModelProcess ModelProcess

// NewModelProcess instantiates a new ModelProcess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelProcess(cmdline string, cpuMax float32, cpuUsage float32, memoryMax int32, memoryUsage int32, nodeId string, nodeName string, openFilesCount int32, pid int32, ppid int32, shortName string, threads int32) *ModelProcess {
	this := ModelProcess{}
	this.Cmdline = cmdline
	this.CpuMax = cpuMax
	this.CpuUsage = cpuUsage
	this.MemoryMax = memoryMax
	this.MemoryUsage = memoryUsage
	this.NodeId = nodeId
	this.NodeName = nodeName
	this.OpenFilesCount = openFilesCount
	this.Pid = pid
	this.Ppid = ppid
	this.ShortName = shortName
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

// GetCpuMax returns the CpuMax field value
func (o *ModelProcess) GetCpuMax() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CpuMax
}

// GetCpuMaxOk returns a tuple with the CpuMax field value
// and a boolean to check if the value has been set.
func (o *ModelProcess) GetCpuMaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuMax, true
}

// SetCpuMax sets field value
func (o *ModelProcess) SetCpuMax(v float32) {
	o.CpuMax = v
}

// GetCpuUsage returns the CpuUsage field value
func (o *ModelProcess) GetCpuUsage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CpuUsage
}

// GetCpuUsageOk returns a tuple with the CpuUsage field value
// and a boolean to check if the value has been set.
func (o *ModelProcess) GetCpuUsageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuUsage, true
}

// SetCpuUsage sets field value
func (o *ModelProcess) SetCpuUsage(v float32) {
	o.CpuUsage = v
}

// GetMemoryMax returns the MemoryMax field value
func (o *ModelProcess) GetMemoryMax() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MemoryMax
}

// GetMemoryMaxOk returns a tuple with the MemoryMax field value
// and a boolean to check if the value has been set.
func (o *ModelProcess) GetMemoryMaxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemoryMax, true
}

// SetMemoryMax sets field value
func (o *ModelProcess) SetMemoryMax(v int32) {
	o.MemoryMax = v
}

// GetMemoryUsage returns the MemoryUsage field value
func (o *ModelProcess) GetMemoryUsage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MemoryUsage
}

// GetMemoryUsageOk returns a tuple with the MemoryUsage field value
// and a boolean to check if the value has been set.
func (o *ModelProcess) GetMemoryUsageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemoryUsage, true
}

// SetMemoryUsage sets field value
func (o *ModelProcess) SetMemoryUsage(v int32) {
	o.MemoryUsage = v
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

// GetNodeName returns the NodeName field value
func (o *ModelProcess) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *ModelProcess) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value
func (o *ModelProcess) SetNodeName(v string) {
	o.NodeName = v
}

// GetOpenFilesCount returns the OpenFilesCount field value
func (o *ModelProcess) GetOpenFilesCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OpenFilesCount
}

// GetOpenFilesCountOk returns a tuple with the OpenFilesCount field value
// and a boolean to check if the value has been set.
func (o *ModelProcess) GetOpenFilesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpenFilesCount, true
}

// SetOpenFilesCount sets field value
func (o *ModelProcess) SetOpenFilesCount(v int32) {
	o.OpenFilesCount = v
}

// GetPid returns the Pid field value
func (o *ModelProcess) GetPid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pid
}

// GetPidOk returns a tuple with the Pid field value
// and a boolean to check if the value has been set.
func (o *ModelProcess) GetPidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pid, true
}

// SetPid sets field value
func (o *ModelProcess) SetPid(v int32) {
	o.Pid = v
}

// GetPpid returns the Ppid field value
func (o *ModelProcess) GetPpid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Ppid
}

// GetPpidOk returns a tuple with the Ppid field value
// and a boolean to check if the value has been set.
func (o *ModelProcess) GetPpidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ppid, true
}

// SetPpid sets field value
func (o *ModelProcess) SetPpid(v int32) {
	o.Ppid = v
}

// GetShortName returns the ShortName field value
func (o *ModelProcess) GetShortName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value
// and a boolean to check if the value has been set.
func (o *ModelProcess) GetShortNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShortName, true
}

// SetShortName sets field value
func (o *ModelProcess) SetShortName(v string) {
	o.ShortName = v
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
	toSerialize["cpu_max"] = o.CpuMax
	toSerialize["cpu_usage"] = o.CpuUsage
	toSerialize["memory_max"] = o.MemoryMax
	toSerialize["memory_usage"] = o.MemoryUsage
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_name"] = o.NodeName
	toSerialize["open_files_count"] = o.OpenFilesCount
	toSerialize["pid"] = o.Pid
	toSerialize["ppid"] = o.Ppid
	toSerialize["short_name"] = o.ShortName
	toSerialize["threads"] = o.Threads
	return toSerialize, nil
}

func (o *ModelProcess) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cmdline",
		"cpu_max",
		"cpu_usage",
		"memory_max",
		"memory_usage",
		"node_id",
		"node_name",
		"open_files_count",
		"pid",
		"ppid",
		"short_name",
		"threads",
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

	varModelProcess := _ModelProcess{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelProcess)

	if err != nil {
		return err
	}

	*o = ModelProcess(varModelProcess)

	return err
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


