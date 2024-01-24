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
	"bytes"
	"fmt"
)

// checks if the DiagnosisGenerateAgentDiagnosticLogsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiagnosisGenerateAgentDiagnosticLogsRequest{}

// DiagnosisGenerateAgentDiagnosticLogsRequest struct for DiagnosisGenerateAgentDiagnosticLogsRequest
type DiagnosisGenerateAgentDiagnosticLogsRequest struct {
	NodeIds []DiagnosisNodeIdentifier `json:"node_ids"`
	Tail int32 `json:"tail"`
}

type _DiagnosisGenerateAgentDiagnosticLogsRequest DiagnosisGenerateAgentDiagnosticLogsRequest

// NewDiagnosisGenerateAgentDiagnosticLogsRequest instantiates a new DiagnosisGenerateAgentDiagnosticLogsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiagnosisGenerateAgentDiagnosticLogsRequest(nodeIds []DiagnosisNodeIdentifier, tail int32) *DiagnosisGenerateAgentDiagnosticLogsRequest {
	this := DiagnosisGenerateAgentDiagnosticLogsRequest{}
	this.NodeIds = nodeIds
	this.Tail = tail
	return &this
}

// NewDiagnosisGenerateAgentDiagnosticLogsRequestWithDefaults instantiates a new DiagnosisGenerateAgentDiagnosticLogsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiagnosisGenerateAgentDiagnosticLogsRequestWithDefaults() *DiagnosisGenerateAgentDiagnosticLogsRequest {
	this := DiagnosisGenerateAgentDiagnosticLogsRequest{}
	return &this
}

// GetNodeIds returns the NodeIds field value
// If the value is explicit nil, the zero value for []DiagnosisNodeIdentifier will be returned
func (o *DiagnosisGenerateAgentDiagnosticLogsRequest) GetNodeIds() []DiagnosisNodeIdentifier {
	if o == nil {
		var ret []DiagnosisNodeIdentifier
		return ret
	}

	return o.NodeIds
}

// GetNodeIdsOk returns a tuple with the NodeIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiagnosisGenerateAgentDiagnosticLogsRequest) GetNodeIdsOk() ([]DiagnosisNodeIdentifier, bool) {
	if o == nil || IsNil(o.NodeIds) {
		return nil, false
	}
	return o.NodeIds, true
}

// SetNodeIds sets field value
func (o *DiagnosisGenerateAgentDiagnosticLogsRequest) SetNodeIds(v []DiagnosisNodeIdentifier) {
	o.NodeIds = v
}

// GetTail returns the Tail field value
func (o *DiagnosisGenerateAgentDiagnosticLogsRequest) GetTail() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Tail
}

// GetTailOk returns a tuple with the Tail field value
// and a boolean to check if the value has been set.
func (o *DiagnosisGenerateAgentDiagnosticLogsRequest) GetTailOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tail, true
}

// SetTail sets field value
func (o *DiagnosisGenerateAgentDiagnosticLogsRequest) SetTail(v int32) {
	o.Tail = v
}

func (o DiagnosisGenerateAgentDiagnosticLogsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiagnosisGenerateAgentDiagnosticLogsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NodeIds != nil {
		toSerialize["node_ids"] = o.NodeIds
	}
	toSerialize["tail"] = o.Tail
	return toSerialize, nil
}

func (o *DiagnosisGenerateAgentDiagnosticLogsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"node_ids",
		"tail",
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

	varDiagnosisGenerateAgentDiagnosticLogsRequest := _DiagnosisGenerateAgentDiagnosticLogsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDiagnosisGenerateAgentDiagnosticLogsRequest)

	if err != nil {
		return err
	}

	*o = DiagnosisGenerateAgentDiagnosticLogsRequest(varDiagnosisGenerateAgentDiagnosticLogsRequest)

	return err
}

type NullableDiagnosisGenerateAgentDiagnosticLogsRequest struct {
	value *DiagnosisGenerateAgentDiagnosticLogsRequest
	isSet bool
}

func (v NullableDiagnosisGenerateAgentDiagnosticLogsRequest) Get() *DiagnosisGenerateAgentDiagnosticLogsRequest {
	return v.value
}

func (v *NullableDiagnosisGenerateAgentDiagnosticLogsRequest) Set(val *DiagnosisGenerateAgentDiagnosticLogsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDiagnosisGenerateAgentDiagnosticLogsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDiagnosisGenerateAgentDiagnosticLogsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiagnosisGenerateAgentDiagnosticLogsRequest(val *DiagnosisGenerateAgentDiagnosticLogsRequest) *NullableDiagnosisGenerateAgentDiagnosticLogsRequest {
	return &NullableDiagnosisGenerateAgentDiagnosticLogsRequest{value: val, isSet: true}
}

func (v NullableDiagnosisGenerateAgentDiagnosticLogsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiagnosisGenerateAgentDiagnosticLogsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


