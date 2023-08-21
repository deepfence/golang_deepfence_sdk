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

// checks if the ModelCloudInstanceDeployment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCloudInstanceDeployment{}

// ModelCloudInstanceDeployment struct for ModelCloudInstanceDeployment
type ModelCloudInstanceDeployment struct {
	AccountId *string `json:"account_id,omitempty"`
	InstanceId *string `json:"instance_id,omitempty"`
	Region *string `json:"region,omitempty"`
}

// NewModelCloudInstanceDeployment instantiates a new ModelCloudInstanceDeployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCloudInstanceDeployment() *ModelCloudInstanceDeployment {
	this := ModelCloudInstanceDeployment{}
	return &this
}

// NewModelCloudInstanceDeploymentWithDefaults instantiates a new ModelCloudInstanceDeployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCloudInstanceDeploymentWithDefaults() *ModelCloudInstanceDeployment {
	this := ModelCloudInstanceDeployment{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ModelCloudInstanceDeployment) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudInstanceDeployment) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ModelCloudInstanceDeployment) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ModelCloudInstanceDeployment) SetAccountId(v string) {
	o.AccountId = &v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *ModelCloudInstanceDeployment) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudInstanceDeployment) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *ModelCloudInstanceDeployment) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *ModelCloudInstanceDeployment) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ModelCloudInstanceDeployment) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudInstanceDeployment) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ModelCloudInstanceDeployment) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ModelCloudInstanceDeployment) SetRegion(v string) {
	o.Region = &v
}

func (o ModelCloudInstanceDeployment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCloudInstanceDeployment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.InstanceId) {
		toSerialize["instance_id"] = o.InstanceId
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	return toSerialize, nil
}

type NullableModelCloudInstanceDeployment struct {
	value *ModelCloudInstanceDeployment
	isSet bool
}

func (v NullableModelCloudInstanceDeployment) Get() *ModelCloudInstanceDeployment {
	return v.value
}

func (v *NullableModelCloudInstanceDeployment) Set(val *ModelCloudInstanceDeployment) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCloudInstanceDeployment) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCloudInstanceDeployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCloudInstanceDeployment(val *ModelCloudInstanceDeployment) *NullableModelCloudInstanceDeployment {
	return &NullableModelCloudInstanceDeployment{value: val, isSet: true}
}

func (v NullableModelCloudInstanceDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCloudInstanceDeployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


