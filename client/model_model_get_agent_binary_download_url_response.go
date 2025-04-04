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

// checks if the ModelGetAgentBinaryDownloadURLResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelGetAgentBinaryDownloadURLResponse{}

// ModelGetAgentBinaryDownloadURLResponse struct for ModelGetAgentBinaryDownloadURLResponse
type ModelGetAgentBinaryDownloadURLResponse struct {
	AgentBinaryAmd64DownloadUrl *string `json:"agent_binary_amd64_download_url,omitempty"`
	AgentBinaryArm64DownloadUrl *string `json:"agent_binary_arm64_download_url,omitempty"`
	StartAgentScriptDownloadUrl *string `json:"start_agent_script_download_url,omitempty"`
	UninstallAgentScriptDownloadUrl *string `json:"uninstall_agent_script_download_url,omitempty"`
}

// NewModelGetAgentBinaryDownloadURLResponse instantiates a new ModelGetAgentBinaryDownloadURLResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelGetAgentBinaryDownloadURLResponse() *ModelGetAgentBinaryDownloadURLResponse {
	this := ModelGetAgentBinaryDownloadURLResponse{}
	return &this
}

// NewModelGetAgentBinaryDownloadURLResponseWithDefaults instantiates a new ModelGetAgentBinaryDownloadURLResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelGetAgentBinaryDownloadURLResponseWithDefaults() *ModelGetAgentBinaryDownloadURLResponse {
	this := ModelGetAgentBinaryDownloadURLResponse{}
	return &this
}

// GetAgentBinaryAmd64DownloadUrl returns the AgentBinaryAmd64DownloadUrl field value if set, zero value otherwise.
func (o *ModelGetAgentBinaryDownloadURLResponse) GetAgentBinaryAmd64DownloadUrl() string {
	if o == nil || IsNil(o.AgentBinaryAmd64DownloadUrl) {
		var ret string
		return ret
	}
	return *o.AgentBinaryAmd64DownloadUrl
}

// GetAgentBinaryAmd64DownloadUrlOk returns a tuple with the AgentBinaryAmd64DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelGetAgentBinaryDownloadURLResponse) GetAgentBinaryAmd64DownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AgentBinaryAmd64DownloadUrl) {
		return nil, false
	}
	return o.AgentBinaryAmd64DownloadUrl, true
}

// HasAgentBinaryAmd64DownloadUrl returns a boolean if a field has been set.
func (o *ModelGetAgentBinaryDownloadURLResponse) HasAgentBinaryAmd64DownloadUrl() bool {
	if o != nil && !IsNil(o.AgentBinaryAmd64DownloadUrl) {
		return true
	}

	return false
}

// SetAgentBinaryAmd64DownloadUrl gets a reference to the given string and assigns it to the AgentBinaryAmd64DownloadUrl field.
func (o *ModelGetAgentBinaryDownloadURLResponse) SetAgentBinaryAmd64DownloadUrl(v string) {
	o.AgentBinaryAmd64DownloadUrl = &v
}

// GetAgentBinaryArm64DownloadUrl returns the AgentBinaryArm64DownloadUrl field value if set, zero value otherwise.
func (o *ModelGetAgentBinaryDownloadURLResponse) GetAgentBinaryArm64DownloadUrl() string {
	if o == nil || IsNil(o.AgentBinaryArm64DownloadUrl) {
		var ret string
		return ret
	}
	return *o.AgentBinaryArm64DownloadUrl
}

// GetAgentBinaryArm64DownloadUrlOk returns a tuple with the AgentBinaryArm64DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelGetAgentBinaryDownloadURLResponse) GetAgentBinaryArm64DownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AgentBinaryArm64DownloadUrl) {
		return nil, false
	}
	return o.AgentBinaryArm64DownloadUrl, true
}

// HasAgentBinaryArm64DownloadUrl returns a boolean if a field has been set.
func (o *ModelGetAgentBinaryDownloadURLResponse) HasAgentBinaryArm64DownloadUrl() bool {
	if o != nil && !IsNil(o.AgentBinaryArm64DownloadUrl) {
		return true
	}

	return false
}

// SetAgentBinaryArm64DownloadUrl gets a reference to the given string and assigns it to the AgentBinaryArm64DownloadUrl field.
func (o *ModelGetAgentBinaryDownloadURLResponse) SetAgentBinaryArm64DownloadUrl(v string) {
	o.AgentBinaryArm64DownloadUrl = &v
}

// GetStartAgentScriptDownloadUrl returns the StartAgentScriptDownloadUrl field value if set, zero value otherwise.
func (o *ModelGetAgentBinaryDownloadURLResponse) GetStartAgentScriptDownloadUrl() string {
	if o == nil || IsNil(o.StartAgentScriptDownloadUrl) {
		var ret string
		return ret
	}
	return *o.StartAgentScriptDownloadUrl
}

// GetStartAgentScriptDownloadUrlOk returns a tuple with the StartAgentScriptDownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelGetAgentBinaryDownloadURLResponse) GetStartAgentScriptDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.StartAgentScriptDownloadUrl) {
		return nil, false
	}
	return o.StartAgentScriptDownloadUrl, true
}

// HasStartAgentScriptDownloadUrl returns a boolean if a field has been set.
func (o *ModelGetAgentBinaryDownloadURLResponse) HasStartAgentScriptDownloadUrl() bool {
	if o != nil && !IsNil(o.StartAgentScriptDownloadUrl) {
		return true
	}

	return false
}

// SetStartAgentScriptDownloadUrl gets a reference to the given string and assigns it to the StartAgentScriptDownloadUrl field.
func (o *ModelGetAgentBinaryDownloadURLResponse) SetStartAgentScriptDownloadUrl(v string) {
	o.StartAgentScriptDownloadUrl = &v
}

// GetUninstallAgentScriptDownloadUrl returns the UninstallAgentScriptDownloadUrl field value if set, zero value otherwise.
func (o *ModelGetAgentBinaryDownloadURLResponse) GetUninstallAgentScriptDownloadUrl() string {
	if o == nil || IsNil(o.UninstallAgentScriptDownloadUrl) {
		var ret string
		return ret
	}
	return *o.UninstallAgentScriptDownloadUrl
}

// GetUninstallAgentScriptDownloadUrlOk returns a tuple with the UninstallAgentScriptDownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelGetAgentBinaryDownloadURLResponse) GetUninstallAgentScriptDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.UninstallAgentScriptDownloadUrl) {
		return nil, false
	}
	return o.UninstallAgentScriptDownloadUrl, true
}

// HasUninstallAgentScriptDownloadUrl returns a boolean if a field has been set.
func (o *ModelGetAgentBinaryDownloadURLResponse) HasUninstallAgentScriptDownloadUrl() bool {
	if o != nil && !IsNil(o.UninstallAgentScriptDownloadUrl) {
		return true
	}

	return false
}

// SetUninstallAgentScriptDownloadUrl gets a reference to the given string and assigns it to the UninstallAgentScriptDownloadUrl field.
func (o *ModelGetAgentBinaryDownloadURLResponse) SetUninstallAgentScriptDownloadUrl(v string) {
	o.UninstallAgentScriptDownloadUrl = &v
}

func (o ModelGetAgentBinaryDownloadURLResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelGetAgentBinaryDownloadURLResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgentBinaryAmd64DownloadUrl) {
		toSerialize["agent_binary_amd64_download_url"] = o.AgentBinaryAmd64DownloadUrl
	}
	if !IsNil(o.AgentBinaryArm64DownloadUrl) {
		toSerialize["agent_binary_arm64_download_url"] = o.AgentBinaryArm64DownloadUrl
	}
	if !IsNil(o.StartAgentScriptDownloadUrl) {
		toSerialize["start_agent_script_download_url"] = o.StartAgentScriptDownloadUrl
	}
	if !IsNil(o.UninstallAgentScriptDownloadUrl) {
		toSerialize["uninstall_agent_script_download_url"] = o.UninstallAgentScriptDownloadUrl
	}
	return toSerialize, nil
}

type NullableModelGetAgentBinaryDownloadURLResponse struct {
	value *ModelGetAgentBinaryDownloadURLResponse
	isSet bool
}

func (v NullableModelGetAgentBinaryDownloadURLResponse) Get() *ModelGetAgentBinaryDownloadURLResponse {
	return v.value
}

func (v *NullableModelGetAgentBinaryDownloadURLResponse) Set(val *ModelGetAgentBinaryDownloadURLResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelGetAgentBinaryDownloadURLResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelGetAgentBinaryDownloadURLResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelGetAgentBinaryDownloadURLResponse(val *ModelGetAgentBinaryDownloadURLResponse) *NullableModelGetAgentBinaryDownloadURLResponse {
	return &NullableModelGetAgentBinaryDownloadURLResponse{value: val, isSet: true}
}

func (v NullableModelGetAgentBinaryDownloadURLResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelGetAgentBinaryDownloadURLResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


