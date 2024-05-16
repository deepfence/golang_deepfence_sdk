/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.2.1
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UtilsReportOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UtilsReportOptions{}

// UtilsReportOptions struct for UtilsReportOptions
type UtilsReportOptions struct {
	SbomFormat *string `json:"sbom_format,omitempty"`
}

// NewUtilsReportOptions instantiates a new UtilsReportOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUtilsReportOptions() *UtilsReportOptions {
	this := UtilsReportOptions{}
	return &this
}

// NewUtilsReportOptionsWithDefaults instantiates a new UtilsReportOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUtilsReportOptionsWithDefaults() *UtilsReportOptions {
	this := UtilsReportOptions{}
	return &this
}

// GetSbomFormat returns the SbomFormat field value if set, zero value otherwise.
func (o *UtilsReportOptions) GetSbomFormat() string {
	if o == nil || IsNil(o.SbomFormat) {
		var ret string
		return ret
	}
	return *o.SbomFormat
}

// GetSbomFormatOk returns a tuple with the SbomFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsReportOptions) GetSbomFormatOk() (*string, bool) {
	if o == nil || IsNil(o.SbomFormat) {
		return nil, false
	}
	return o.SbomFormat, true
}

// HasSbomFormat returns a boolean if a field has been set.
func (o *UtilsReportOptions) HasSbomFormat() bool {
	if o != nil && !IsNil(o.SbomFormat) {
		return true
	}

	return false
}

// SetSbomFormat gets a reference to the given string and assigns it to the SbomFormat field.
func (o *UtilsReportOptions) SetSbomFormat(v string) {
	o.SbomFormat = &v
}

func (o UtilsReportOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UtilsReportOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SbomFormat) {
		toSerialize["sbom_format"] = o.SbomFormat
	}
	return toSerialize, nil
}

type NullableUtilsReportOptions struct {
	value *UtilsReportOptions
	isSet bool
}

func (v NullableUtilsReportOptions) Get() *UtilsReportOptions {
	return v.value
}

func (v *NullableUtilsReportOptions) Set(val *UtilsReportOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableUtilsReportOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableUtilsReportOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtilsReportOptions(val *UtilsReportOptions) *NullableUtilsReportOptions {
	return &NullableUtilsReportOptions{value: val, isSet: true}
}

func (v NullableUtilsReportOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtilsReportOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


