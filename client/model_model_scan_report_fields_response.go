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
)

// checks if the ModelScanReportFieldsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelScanReportFieldsResponse{}

// ModelScanReportFieldsResponse struct for ModelScanReportFieldsResponse
type ModelScanReportFieldsResponse struct {
	Compliance []string `json:"compliance,omitempty"`
	Malware []string `json:"malware,omitempty"`
	Secret []string `json:"secret,omitempty"`
	Vulnerability []string `json:"vulnerability,omitempty"`
}

// NewModelScanReportFieldsResponse instantiates a new ModelScanReportFieldsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelScanReportFieldsResponse() *ModelScanReportFieldsResponse {
	this := ModelScanReportFieldsResponse{}
	return &this
}

// NewModelScanReportFieldsResponseWithDefaults instantiates a new ModelScanReportFieldsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelScanReportFieldsResponseWithDefaults() *ModelScanReportFieldsResponse {
	this := ModelScanReportFieldsResponse{}
	return &this
}

// GetCompliance returns the Compliance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelScanReportFieldsResponse) GetCompliance() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Compliance
}

// GetComplianceOk returns a tuple with the Compliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelScanReportFieldsResponse) GetComplianceOk() ([]string, bool) {
	if o == nil || IsNil(o.Compliance) {
		return nil, false
	}
	return o.Compliance, true
}

// HasCompliance returns a boolean if a field has been set.
func (o *ModelScanReportFieldsResponse) HasCompliance() bool {
	if o != nil && !IsNil(o.Compliance) {
		return true
	}

	return false
}

// SetCompliance gets a reference to the given []string and assigns it to the Compliance field.
func (o *ModelScanReportFieldsResponse) SetCompliance(v []string) {
	o.Compliance = v
}

// GetMalware returns the Malware field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelScanReportFieldsResponse) GetMalware() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Malware
}

// GetMalwareOk returns a tuple with the Malware field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelScanReportFieldsResponse) GetMalwareOk() ([]string, bool) {
	if o == nil || IsNil(o.Malware) {
		return nil, false
	}
	return o.Malware, true
}

// HasMalware returns a boolean if a field has been set.
func (o *ModelScanReportFieldsResponse) HasMalware() bool {
	if o != nil && !IsNil(o.Malware) {
		return true
	}

	return false
}

// SetMalware gets a reference to the given []string and assigns it to the Malware field.
func (o *ModelScanReportFieldsResponse) SetMalware(v []string) {
	o.Malware = v
}

// GetSecret returns the Secret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelScanReportFieldsResponse) GetSecret() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelScanReportFieldsResponse) GetSecretOk() ([]string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ModelScanReportFieldsResponse) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given []string and assigns it to the Secret field.
func (o *ModelScanReportFieldsResponse) SetSecret(v []string) {
	o.Secret = v
}

// GetVulnerability returns the Vulnerability field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelScanReportFieldsResponse) GetVulnerability() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Vulnerability
}

// GetVulnerabilityOk returns a tuple with the Vulnerability field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelScanReportFieldsResponse) GetVulnerabilityOk() ([]string, bool) {
	if o == nil || IsNil(o.Vulnerability) {
		return nil, false
	}
	return o.Vulnerability, true
}

// HasVulnerability returns a boolean if a field has been set.
func (o *ModelScanReportFieldsResponse) HasVulnerability() bool {
	if o != nil && !IsNil(o.Vulnerability) {
		return true
	}

	return false
}

// SetVulnerability gets a reference to the given []string and assigns it to the Vulnerability field.
func (o *ModelScanReportFieldsResponse) SetVulnerability(v []string) {
	o.Vulnerability = v
}

func (o ModelScanReportFieldsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelScanReportFieldsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Compliance != nil {
		toSerialize["compliance"] = o.Compliance
	}
	if o.Malware != nil {
		toSerialize["malware"] = o.Malware
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.Vulnerability != nil {
		toSerialize["vulnerability"] = o.Vulnerability
	}
	return toSerialize, nil
}

type NullableModelScanReportFieldsResponse struct {
	value *ModelScanReportFieldsResponse
	isSet bool
}

func (v NullableModelScanReportFieldsResponse) Get() *ModelScanReportFieldsResponse {
	return v.value
}

func (v *NullableModelScanReportFieldsResponse) Set(val *ModelScanReportFieldsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelScanReportFieldsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelScanReportFieldsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelScanReportFieldsResponse(val *ModelScanReportFieldsResponse) *NullableModelScanReportFieldsResponse {
	return &NullableModelScanReportFieldsResponse{value: val, isSet: true}
}

func (v NullableModelScanReportFieldsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelScanReportFieldsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


