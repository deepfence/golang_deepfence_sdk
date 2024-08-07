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
)

// checks if the DiagnosisGetDiagnosticLogsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiagnosisGetDiagnosticLogsResponse{}

// DiagnosisGetDiagnosticLogsResponse struct for DiagnosisGetDiagnosticLogsResponse
type DiagnosisGetDiagnosticLogsResponse struct {
	AgentLogs []DiagnosisDiagnosticLogsLink `json:"agent_logs,omitempty"`
	ConsoleLogs []DiagnosisDiagnosticLogsLink `json:"console_logs,omitempty"`
}

// NewDiagnosisGetDiagnosticLogsResponse instantiates a new DiagnosisGetDiagnosticLogsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiagnosisGetDiagnosticLogsResponse() *DiagnosisGetDiagnosticLogsResponse {
	this := DiagnosisGetDiagnosticLogsResponse{}
	return &this
}

// NewDiagnosisGetDiagnosticLogsResponseWithDefaults instantiates a new DiagnosisGetDiagnosticLogsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiagnosisGetDiagnosticLogsResponseWithDefaults() *DiagnosisGetDiagnosticLogsResponse {
	this := DiagnosisGetDiagnosticLogsResponse{}
	return &this
}

// GetAgentLogs returns the AgentLogs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiagnosisGetDiagnosticLogsResponse) GetAgentLogs() []DiagnosisDiagnosticLogsLink {
	if o == nil {
		var ret []DiagnosisDiagnosticLogsLink
		return ret
	}
	return o.AgentLogs
}

// GetAgentLogsOk returns a tuple with the AgentLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiagnosisGetDiagnosticLogsResponse) GetAgentLogsOk() ([]DiagnosisDiagnosticLogsLink, bool) {
	if o == nil || IsNil(o.AgentLogs) {
		return nil, false
	}
	return o.AgentLogs, true
}

// HasAgentLogs returns a boolean if a field has been set.
func (o *DiagnosisGetDiagnosticLogsResponse) HasAgentLogs() bool {
	if o != nil && !IsNil(o.AgentLogs) {
		return true
	}

	return false
}

// SetAgentLogs gets a reference to the given []DiagnosisDiagnosticLogsLink and assigns it to the AgentLogs field.
func (o *DiagnosisGetDiagnosticLogsResponse) SetAgentLogs(v []DiagnosisDiagnosticLogsLink) {
	o.AgentLogs = v
}

// GetConsoleLogs returns the ConsoleLogs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiagnosisGetDiagnosticLogsResponse) GetConsoleLogs() []DiagnosisDiagnosticLogsLink {
	if o == nil {
		var ret []DiagnosisDiagnosticLogsLink
		return ret
	}
	return o.ConsoleLogs
}

// GetConsoleLogsOk returns a tuple with the ConsoleLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiagnosisGetDiagnosticLogsResponse) GetConsoleLogsOk() ([]DiagnosisDiagnosticLogsLink, bool) {
	if o == nil || IsNil(o.ConsoleLogs) {
		return nil, false
	}
	return o.ConsoleLogs, true
}

// HasConsoleLogs returns a boolean if a field has been set.
func (o *DiagnosisGetDiagnosticLogsResponse) HasConsoleLogs() bool {
	if o != nil && !IsNil(o.ConsoleLogs) {
		return true
	}

	return false
}

// SetConsoleLogs gets a reference to the given []DiagnosisDiagnosticLogsLink and assigns it to the ConsoleLogs field.
func (o *DiagnosisGetDiagnosticLogsResponse) SetConsoleLogs(v []DiagnosisDiagnosticLogsLink) {
	o.ConsoleLogs = v
}

func (o DiagnosisGetDiagnosticLogsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiagnosisGetDiagnosticLogsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AgentLogs != nil {
		toSerialize["agent_logs"] = o.AgentLogs
	}
	if o.ConsoleLogs != nil {
		toSerialize["console_logs"] = o.ConsoleLogs
	}
	return toSerialize, nil
}

type NullableDiagnosisGetDiagnosticLogsResponse struct {
	value *DiagnosisGetDiagnosticLogsResponse
	isSet bool
}

func (v NullableDiagnosisGetDiagnosticLogsResponse) Get() *DiagnosisGetDiagnosticLogsResponse {
	return v.value
}

func (v *NullableDiagnosisGetDiagnosticLogsResponse) Set(val *DiagnosisGetDiagnosticLogsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDiagnosisGetDiagnosticLogsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDiagnosisGetDiagnosticLogsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiagnosisGetDiagnosticLogsResponse(val *DiagnosisGetDiagnosticLogsResponse) *NullableDiagnosisGetDiagnosticLogsResponse {
	return &NullableDiagnosisGetDiagnosticLogsResponse{value: val, isSet: true}
}

func (v NullableDiagnosisGetDiagnosticLogsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiagnosisGetDiagnosticLogsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


