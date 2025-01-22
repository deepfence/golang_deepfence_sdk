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

// checks if the ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance{}

// ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance struct for ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance
type ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance struct {
	New []ModelCompliance `json:"new"`
}

type _ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance

// NewModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance instantiates a new ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance(new []ModelCompliance) *ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance {
	this := ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance{}
	this.New = new
	return &this
}

// NewModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelComplianceWithDefaults instantiates a new ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelComplianceWithDefaults() *ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance {
	this := ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance{}
	return &this
}

// GetNew returns the New field value
// If the value is explicit nil, the zero value for []ModelCompliance will be returned
func (o *ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance) GetNew() []ModelCompliance {
	if o == nil {
		var ret []ModelCompliance
		return ret
	}

	return o.New
}

// GetNewOk returns a tuple with the New field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance) GetNewOk() ([]ModelCompliance, bool) {
	if o == nil || IsNil(o.New) {
		return nil, false
	}
	return o.New, true
}

// SetNew sets field value
func (o *ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance) SetNew(v []ModelCompliance) {
	o.New = v
}

func (o ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.New != nil {
		toSerialize["new"] = o.New
	}
	return toSerialize, nil
}

func (o *ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"new",
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

	varModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance := _ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance)

	if err != nil {
		return err
	}

	*o = ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance(varModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance)

	return err
}

type NullableModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance struct {
	value *ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance
	isSet bool
}

func (v NullableModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance) Get() *ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance {
	return v.value
}

func (v *NullableModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance) Set(val *ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance) {
	v.value = val
	v.isSet = true
}

func (v NullableModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance) IsSet() bool {
	return v.isSet
}

func (v *NullableModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance(val *ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance) *NullableModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance {
	return &NullableModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance{value: val, isSet: true}
}

func (v NullableModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


