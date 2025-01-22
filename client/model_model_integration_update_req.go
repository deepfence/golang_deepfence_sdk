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

// checks if the ModelIntegrationUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelIntegrationUpdateReq{}

// ModelIntegrationUpdateReq struct for ModelIntegrationUpdateReq
type ModelIntegrationUpdateReq struct {
	Config map[string]interface{} `json:"config,omitempty"`
	Filters *ModelIntegrationFilters `json:"filters,omitempty"`
	Id *int32 `json:"id,omitempty"`
	IntegrationType *string `json:"integration_type,omitempty"`
	NotificationType *string `json:"notification_type,omitempty"`
	SendSummary *bool `json:"send_summary,omitempty"`
}

// NewModelIntegrationUpdateReq instantiates a new ModelIntegrationUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelIntegrationUpdateReq() *ModelIntegrationUpdateReq {
	this := ModelIntegrationUpdateReq{}
	return &this
}

// NewModelIntegrationUpdateReqWithDefaults instantiates a new ModelIntegrationUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelIntegrationUpdateReqWithDefaults() *ModelIntegrationUpdateReq {
	this := ModelIntegrationUpdateReq{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelIntegrationUpdateReq) GetConfig() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelIntegrationUpdateReq) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ModelIntegrationUpdateReq) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *ModelIntegrationUpdateReq) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ModelIntegrationUpdateReq) GetFilters() ModelIntegrationFilters {
	if o == nil || IsNil(o.Filters) {
		var ret ModelIntegrationFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelIntegrationUpdateReq) GetFiltersOk() (*ModelIntegrationFilters, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ModelIntegrationUpdateReq) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given ModelIntegrationFilters and assigns it to the Filters field.
func (o *ModelIntegrationUpdateReq) SetFilters(v ModelIntegrationFilters) {
	o.Filters = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelIntegrationUpdateReq) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelIntegrationUpdateReq) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelIntegrationUpdateReq) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelIntegrationUpdateReq) SetId(v int32) {
	o.Id = &v
}

// GetIntegrationType returns the IntegrationType field value if set, zero value otherwise.
func (o *ModelIntegrationUpdateReq) GetIntegrationType() string {
	if o == nil || IsNil(o.IntegrationType) {
		var ret string
		return ret
	}
	return *o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelIntegrationUpdateReq) GetIntegrationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IntegrationType) {
		return nil, false
	}
	return o.IntegrationType, true
}

// HasIntegrationType returns a boolean if a field has been set.
func (o *ModelIntegrationUpdateReq) HasIntegrationType() bool {
	if o != nil && !IsNil(o.IntegrationType) {
		return true
	}

	return false
}

// SetIntegrationType gets a reference to the given string and assigns it to the IntegrationType field.
func (o *ModelIntegrationUpdateReq) SetIntegrationType(v string) {
	o.IntegrationType = &v
}

// GetNotificationType returns the NotificationType field value if set, zero value otherwise.
func (o *ModelIntegrationUpdateReq) GetNotificationType() string {
	if o == nil || IsNil(o.NotificationType) {
		var ret string
		return ret
	}
	return *o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelIntegrationUpdateReq) GetNotificationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationType) {
		return nil, false
	}
	return o.NotificationType, true
}

// HasNotificationType returns a boolean if a field has been set.
func (o *ModelIntegrationUpdateReq) HasNotificationType() bool {
	if o != nil && !IsNil(o.NotificationType) {
		return true
	}

	return false
}

// SetNotificationType gets a reference to the given string and assigns it to the NotificationType field.
func (o *ModelIntegrationUpdateReq) SetNotificationType(v string) {
	o.NotificationType = &v
}

// GetSendSummary returns the SendSummary field value if set, zero value otherwise.
func (o *ModelIntegrationUpdateReq) GetSendSummary() bool {
	if o == nil || IsNil(o.SendSummary) {
		var ret bool
		return ret
	}
	return *o.SendSummary
}

// GetSendSummaryOk returns a tuple with the SendSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelIntegrationUpdateReq) GetSendSummaryOk() (*bool, bool) {
	if o == nil || IsNil(o.SendSummary) {
		return nil, false
	}
	return o.SendSummary, true
}

// HasSendSummary returns a boolean if a field has been set.
func (o *ModelIntegrationUpdateReq) HasSendSummary() bool {
	if o != nil && !IsNil(o.SendSummary) {
		return true
	}

	return false
}

// SetSendSummary gets a reference to the given bool and assigns it to the SendSummary field.
func (o *ModelIntegrationUpdateReq) SetSendSummary(v bool) {
	o.SendSummary = &v
}

func (o ModelIntegrationUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelIntegrationUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IntegrationType) {
		toSerialize["integration_type"] = o.IntegrationType
	}
	if !IsNil(o.NotificationType) {
		toSerialize["notification_type"] = o.NotificationType
	}
	if !IsNil(o.SendSummary) {
		toSerialize["send_summary"] = o.SendSummary
	}
	return toSerialize, nil
}

type NullableModelIntegrationUpdateReq struct {
	value *ModelIntegrationUpdateReq
	isSet bool
}

func (v NullableModelIntegrationUpdateReq) Get() *ModelIntegrationUpdateReq {
	return v.value
}

func (v *NullableModelIntegrationUpdateReq) Set(val *ModelIntegrationUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelIntegrationUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelIntegrationUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelIntegrationUpdateReq(val *ModelIntegrationUpdateReq) *NullableModelIntegrationUpdateReq {
	return &NullableModelIntegrationUpdateReq{value: val, isSet: true}
}

func (v NullableModelIntegrationUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelIntegrationUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


