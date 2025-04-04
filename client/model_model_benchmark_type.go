/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.5.6
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// ModelBenchmarkType the model 'ModelBenchmarkType'
type ModelBenchmarkType string

// List of ModelBenchmarkType
const (
	HIPAA ModelBenchmarkType = "hipaa"
	GDPR ModelBenchmarkType = "gdpr"
	PCI ModelBenchmarkType = "pci"
	NIST ModelBenchmarkType = "nist"
	CIS ModelBenchmarkType = "cis"
	SOC_2 ModelBenchmarkType = "soc_2"
	NSA_CISA ModelBenchmarkType = "nsa-cisa"
	AWS_FOUNDATIONAL_SECURITY ModelBenchmarkType = "aws_foundational_security"
)

// All allowed values of ModelBenchmarkType enum
var AllowedModelBenchmarkTypeEnumValues = []ModelBenchmarkType{
	"hipaa",
	"gdpr",
	"pci",
	"nist",
	"cis",
	"soc_2",
	"nsa-cisa",
	"aws_foundational_security",
}

func (v *ModelBenchmarkType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModelBenchmarkType(value)
	for _, existing := range AllowedModelBenchmarkTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModelBenchmarkType", value)
}

// NewModelBenchmarkTypeFromValue returns a pointer to a valid ModelBenchmarkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModelBenchmarkTypeFromValue(v string) (*ModelBenchmarkType, error) {
	ev := ModelBenchmarkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModelBenchmarkType: valid values are %v", v, AllowedModelBenchmarkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModelBenchmarkType) IsValid() bool {
	for _, existing := range AllowedModelBenchmarkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ModelBenchmarkType value
func (v ModelBenchmarkType) Ptr() *ModelBenchmarkType {
	return &v
}

type NullableModelBenchmarkType struct {
	value *ModelBenchmarkType
	isSet bool
}

func (v NullableModelBenchmarkType) Get() *ModelBenchmarkType {
	return v.value
}

func (v *NullableModelBenchmarkType) Set(val *ModelBenchmarkType) {
	v.value = val
	v.isSet = true
}

func (v NullableModelBenchmarkType) IsSet() bool {
	return v.isSet
}

func (v *NullableModelBenchmarkType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelBenchmarkType(val *ModelBenchmarkType) *NullableModelBenchmarkType {
	return &NullableModelBenchmarkType{value: val, isSet: true}
}

func (v NullableModelBenchmarkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelBenchmarkType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

