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
)

// checks if the ModelResponseAccessToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelResponseAccessToken{}

// ModelResponseAccessToken struct for ModelResponseAccessToken
type ModelResponseAccessToken struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// NewModelResponseAccessToken instantiates a new ModelResponseAccessToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelResponseAccessToken(accessToken string, refreshToken string) *ModelResponseAccessToken {
	this := ModelResponseAccessToken{}
	this.AccessToken = accessToken
	this.RefreshToken = refreshToken
	return &this
}

// NewModelResponseAccessTokenWithDefaults instantiates a new ModelResponseAccessToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelResponseAccessTokenWithDefaults() *ModelResponseAccessToken {
	this := ModelResponseAccessToken{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *ModelResponseAccessToken) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *ModelResponseAccessToken) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *ModelResponseAccessToken) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetRefreshToken returns the RefreshToken field value
func (o *ModelResponseAccessToken) GetRefreshToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value
// and a boolean to check if the value has been set.
func (o *ModelResponseAccessToken) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefreshToken, true
}

// SetRefreshToken sets field value
func (o *ModelResponseAccessToken) SetRefreshToken(v string) {
	o.RefreshToken = v
}

func (o ModelResponseAccessToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelResponseAccessToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_token"] = o.AccessToken
	toSerialize["refresh_token"] = o.RefreshToken
	return toSerialize, nil
}

type NullableModelResponseAccessToken struct {
	value *ModelResponseAccessToken
	isSet bool
}

func (v NullableModelResponseAccessToken) Get() *ModelResponseAccessToken {
	return v.value
}

func (v *NullableModelResponseAccessToken) Set(val *ModelResponseAccessToken) {
	v.value = val
	v.isSet = true
}

func (v NullableModelResponseAccessToken) IsSet() bool {
	return v.isSet
}

func (v *NullableModelResponseAccessToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelResponseAccessToken(val *ModelResponseAccessToken) *NullableModelResponseAccessToken {
	return &NullableModelResponseAccessToken{value: val, isSet: true}
}

func (v NullableModelResponseAccessToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelResponseAccessToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


