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
)

// checks if the ModelGenerateReportResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelGenerateReportResp{}

// ModelGenerateReportResp struct for ModelGenerateReportResp
type ModelGenerateReportResp struct {
	ReportId *string `json:"report_id,omitempty"`
}

// NewModelGenerateReportResp instantiates a new ModelGenerateReportResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelGenerateReportResp() *ModelGenerateReportResp {
	this := ModelGenerateReportResp{}
	return &this
}

// NewModelGenerateReportRespWithDefaults instantiates a new ModelGenerateReportResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelGenerateReportRespWithDefaults() *ModelGenerateReportResp {
	this := ModelGenerateReportResp{}
	return &this
}

// GetReportId returns the ReportId field value if set, zero value otherwise.
func (o *ModelGenerateReportResp) GetReportId() string {
	if o == nil || IsNil(o.ReportId) {
		var ret string
		return ret
	}
	return *o.ReportId
}

// GetReportIdOk returns a tuple with the ReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelGenerateReportResp) GetReportIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReportId) {
		return nil, false
	}
	return o.ReportId, true
}

// HasReportId returns a boolean if a field has been set.
func (o *ModelGenerateReportResp) HasReportId() bool {
	if o != nil && !IsNil(o.ReportId) {
		return true
	}

	return false
}

// SetReportId gets a reference to the given string and assigns it to the ReportId field.
func (o *ModelGenerateReportResp) SetReportId(v string) {
	o.ReportId = &v
}

func (o ModelGenerateReportResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelGenerateReportResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReportId) {
		toSerialize["report_id"] = o.ReportId
	}
	return toSerialize, nil
}

type NullableModelGenerateReportResp struct {
	value *ModelGenerateReportResp
	isSet bool
}

func (v NullableModelGenerateReportResp) Get() *ModelGenerateReportResp {
	return v.value
}

func (v *NullableModelGenerateReportResp) Set(val *ModelGenerateReportResp) {
	v.value = val
	v.isSet = true
}

func (v NullableModelGenerateReportResp) IsSet() bool {
	return v.isSet
}

func (v *NullableModelGenerateReportResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelGenerateReportResp(val *ModelGenerateReportResp) *NullableModelGenerateReportResp {
	return &NullableModelGenerateReportResp{value: val, isSet: true}
}

func (v NullableModelGenerateReportResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelGenerateReportResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


