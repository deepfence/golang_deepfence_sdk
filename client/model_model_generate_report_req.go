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

// checks if the ModelGenerateReportReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelGenerateReportReq{}

// ModelGenerateReportReq struct for ModelGenerateReportReq
type ModelGenerateReportReq struct {
	Duration *int32 `json:"duration,omitempty"`
	Filters *UtilsReportFilters `json:"filters,omitempty"`
	ReportType string `json:"report_type"`
}

type _ModelGenerateReportReq ModelGenerateReportReq

// NewModelGenerateReportReq instantiates a new ModelGenerateReportReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelGenerateReportReq(reportType string) *ModelGenerateReportReq {
	this := ModelGenerateReportReq{}
	this.ReportType = reportType
	return &this
}

// NewModelGenerateReportReqWithDefaults instantiates a new ModelGenerateReportReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelGenerateReportReqWithDefaults() *ModelGenerateReportReq {
	this := ModelGenerateReportReq{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *ModelGenerateReportReq) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelGenerateReportReq) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *ModelGenerateReportReq) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *ModelGenerateReportReq) SetDuration(v int32) {
	o.Duration = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ModelGenerateReportReq) GetFilters() UtilsReportFilters {
	if o == nil || IsNil(o.Filters) {
		var ret UtilsReportFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelGenerateReportReq) GetFiltersOk() (*UtilsReportFilters, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ModelGenerateReportReq) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given UtilsReportFilters and assigns it to the Filters field.
func (o *ModelGenerateReportReq) SetFilters(v UtilsReportFilters) {
	o.Filters = &v
}

// GetReportType returns the ReportType field value
func (o *ModelGenerateReportReq) GetReportType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value
// and a boolean to check if the value has been set.
func (o *ModelGenerateReportReq) GetReportTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportType, true
}

// SetReportType sets field value
func (o *ModelGenerateReportReq) SetReportType(v string) {
	o.ReportType = v
}

func (o ModelGenerateReportReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelGenerateReportReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	toSerialize["report_type"] = o.ReportType
	return toSerialize, nil
}

func (o *ModelGenerateReportReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"report_type",
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

	varModelGenerateReportReq := _ModelGenerateReportReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelGenerateReportReq)

	if err != nil {
		return err
	}

	*o = ModelGenerateReportReq(varModelGenerateReportReq)

	return err
}

type NullableModelGenerateReportReq struct {
	value *ModelGenerateReportReq
	isSet bool
}

func (v NullableModelGenerateReportReq) Get() *ModelGenerateReportReq {
	return v.value
}

func (v *NullableModelGenerateReportReq) Set(val *ModelGenerateReportReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelGenerateReportReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelGenerateReportReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelGenerateReportReq(val *ModelGenerateReportReq) *NullableModelGenerateReportReq {
	return &NullableModelGenerateReportReq{value: val, isSet: true}
}

func (v NullableModelGenerateReportReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelGenerateReportReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


