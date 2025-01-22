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
)

// checks if the ModelDownloadReportResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelDownloadReportResponse{}

// ModelDownloadReportResponse struct for ModelDownloadReportResponse
type ModelDownloadReportResponse struct {
	UrlLink *string `json:"url_link,omitempty"`
}

// NewModelDownloadReportResponse instantiates a new ModelDownloadReportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelDownloadReportResponse() *ModelDownloadReportResponse {
	this := ModelDownloadReportResponse{}
	return &this
}

// NewModelDownloadReportResponseWithDefaults instantiates a new ModelDownloadReportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelDownloadReportResponseWithDefaults() *ModelDownloadReportResponse {
	this := ModelDownloadReportResponse{}
	return &this
}

// GetUrlLink returns the UrlLink field value if set, zero value otherwise.
func (o *ModelDownloadReportResponse) GetUrlLink() string {
	if o == nil || IsNil(o.UrlLink) {
		var ret string
		return ret
	}
	return *o.UrlLink
}

// GetUrlLinkOk returns a tuple with the UrlLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelDownloadReportResponse) GetUrlLinkOk() (*string, bool) {
	if o == nil || IsNil(o.UrlLink) {
		return nil, false
	}
	return o.UrlLink, true
}

// HasUrlLink returns a boolean if a field has been set.
func (o *ModelDownloadReportResponse) HasUrlLink() bool {
	if o != nil && !IsNil(o.UrlLink) {
		return true
	}

	return false
}

// SetUrlLink gets a reference to the given string and assigns it to the UrlLink field.
func (o *ModelDownloadReportResponse) SetUrlLink(v string) {
	o.UrlLink = &v
}

func (o ModelDownloadReportResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelDownloadReportResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UrlLink) {
		toSerialize["url_link"] = o.UrlLink
	}
	return toSerialize, nil
}

type NullableModelDownloadReportResponse struct {
	value *ModelDownloadReportResponse
	isSet bool
}

func (v NullableModelDownloadReportResponse) Get() *ModelDownloadReportResponse {
	return v.value
}

func (v *NullableModelDownloadReportResponse) Set(val *ModelDownloadReportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelDownloadReportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelDownloadReportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelDownloadReportResponse(val *ModelDownloadReportResponse) *NullableModelDownloadReportResponse {
	return &NullableModelDownloadReportResponse{value: val, isSet: true}
}

func (v NullableModelDownloadReportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelDownloadReportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


