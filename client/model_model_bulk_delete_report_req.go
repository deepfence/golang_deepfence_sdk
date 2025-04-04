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

// checks if the ModelBulkDeleteReportReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelBulkDeleteReportReq{}

// ModelBulkDeleteReportReq struct for ModelBulkDeleteReportReq
type ModelBulkDeleteReportReq struct {
	ReportIds []string `json:"report_ids"`
}

type _ModelBulkDeleteReportReq ModelBulkDeleteReportReq

// NewModelBulkDeleteReportReq instantiates a new ModelBulkDeleteReportReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelBulkDeleteReportReq(reportIds []string) *ModelBulkDeleteReportReq {
	this := ModelBulkDeleteReportReq{}
	this.ReportIds = reportIds
	return &this
}

// NewModelBulkDeleteReportReqWithDefaults instantiates a new ModelBulkDeleteReportReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelBulkDeleteReportReqWithDefaults() *ModelBulkDeleteReportReq {
	this := ModelBulkDeleteReportReq{}
	return &this
}

// GetReportIds returns the ReportIds field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ModelBulkDeleteReportReq) GetReportIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ReportIds
}

// GetReportIdsOk returns a tuple with the ReportIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelBulkDeleteReportReq) GetReportIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ReportIds) {
		return nil, false
	}
	return o.ReportIds, true
}

// SetReportIds sets field value
func (o *ModelBulkDeleteReportReq) SetReportIds(v []string) {
	o.ReportIds = v
}

func (o ModelBulkDeleteReportReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelBulkDeleteReportReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ReportIds != nil {
		toSerialize["report_ids"] = o.ReportIds
	}
	return toSerialize, nil
}

func (o *ModelBulkDeleteReportReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"report_ids",
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

	varModelBulkDeleteReportReq := _ModelBulkDeleteReportReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelBulkDeleteReportReq)

	if err != nil {
		return err
	}

	*o = ModelBulkDeleteReportReq(varModelBulkDeleteReportReq)

	return err
}

type NullableModelBulkDeleteReportReq struct {
	value *ModelBulkDeleteReportReq
	isSet bool
}

func (v NullableModelBulkDeleteReportReq) Get() *ModelBulkDeleteReportReq {
	return v.value
}

func (v *NullableModelBulkDeleteReportReq) Set(val *ModelBulkDeleteReportReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelBulkDeleteReportReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelBulkDeleteReportReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelBulkDeleteReportReq(val *ModelBulkDeleteReportReq) *NullableModelBulkDeleteReportReq {
	return &NullableModelBulkDeleteReportReq{value: val, isSet: true}
}

func (v NullableModelBulkDeleteReportReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelBulkDeleteReportReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


