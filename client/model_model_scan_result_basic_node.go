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

// checks if the ModelScanResultBasicNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelScanResultBasicNode{}

// ModelScanResultBasicNode struct for ModelScanResultBasicNode
type ModelScanResultBasicNode struct {
	BasicNodes []ModelBasicNode `json:"basic_nodes"`
	ResultId string `json:"result_id"`
}

type _ModelScanResultBasicNode ModelScanResultBasicNode

// NewModelScanResultBasicNode instantiates a new ModelScanResultBasicNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelScanResultBasicNode(basicNodes []ModelBasicNode, resultId string) *ModelScanResultBasicNode {
	this := ModelScanResultBasicNode{}
	this.BasicNodes = basicNodes
	this.ResultId = resultId
	return &this
}

// NewModelScanResultBasicNodeWithDefaults instantiates a new ModelScanResultBasicNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelScanResultBasicNodeWithDefaults() *ModelScanResultBasicNode {
	this := ModelScanResultBasicNode{}
	return &this
}

// GetBasicNodes returns the BasicNodes field value
// If the value is explicit nil, the zero value for []ModelBasicNode will be returned
func (o *ModelScanResultBasicNode) GetBasicNodes() []ModelBasicNode {
	if o == nil {
		var ret []ModelBasicNode
		return ret
	}

	return o.BasicNodes
}

// GetBasicNodesOk returns a tuple with the BasicNodes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelScanResultBasicNode) GetBasicNodesOk() ([]ModelBasicNode, bool) {
	if o == nil || IsNil(o.BasicNodes) {
		return nil, false
	}
	return o.BasicNodes, true
}

// SetBasicNodes sets field value
func (o *ModelScanResultBasicNode) SetBasicNodes(v []ModelBasicNode) {
	o.BasicNodes = v
}

// GetResultId returns the ResultId field value
func (o *ModelScanResultBasicNode) GetResultId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResultId
}

// GetResultIdOk returns a tuple with the ResultId field value
// and a boolean to check if the value has been set.
func (o *ModelScanResultBasicNode) GetResultIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultId, true
}

// SetResultId sets field value
func (o *ModelScanResultBasicNode) SetResultId(v string) {
	o.ResultId = v
}

func (o ModelScanResultBasicNode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelScanResultBasicNode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BasicNodes != nil {
		toSerialize["basic_nodes"] = o.BasicNodes
	}
	toSerialize["result_id"] = o.ResultId
	return toSerialize, nil
}

func (o *ModelScanResultBasicNode) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"basic_nodes",
		"result_id",
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

	varModelScanResultBasicNode := _ModelScanResultBasicNode{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelScanResultBasicNode)

	if err != nil {
		return err
	}

	*o = ModelScanResultBasicNode(varModelScanResultBasicNode)

	return err
}

type NullableModelScanResultBasicNode struct {
	value *ModelScanResultBasicNode
	isSet bool
}

func (v NullableModelScanResultBasicNode) Get() *ModelScanResultBasicNode {
	return v.value
}

func (v *NullableModelScanResultBasicNode) Set(val *ModelScanResultBasicNode) {
	v.value = val
	v.isSet = true
}

func (v NullableModelScanResultBasicNode) IsSet() bool {
	return v.isSet
}

func (v *NullableModelScanResultBasicNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelScanResultBasicNode(val *ModelScanResultBasicNode) *NullableModelScanResultBasicNode {
	return &NullableModelScanResultBasicNode{value: val, isSet: true}
}

func (v NullableModelScanResultBasicNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelScanResultBasicNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


