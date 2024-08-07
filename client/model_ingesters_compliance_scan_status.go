/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v3.0.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the IngestersComplianceScanStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestersComplianceScanStatus{}

// IngestersComplianceScanStatus struct for IngestersComplianceScanStatus
type IngestersComplianceScanStatus struct {
	ScanId *string `json:"scan_id,omitempty"`
	ScanMessage *string `json:"scan_message,omitempty"`
	ScanStatus *string `json:"scan_status,omitempty"`
}

// NewIngestersComplianceScanStatus instantiates a new IngestersComplianceScanStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestersComplianceScanStatus() *IngestersComplianceScanStatus {
	this := IngestersComplianceScanStatus{}
	return &this
}

// NewIngestersComplianceScanStatusWithDefaults instantiates a new IngestersComplianceScanStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestersComplianceScanStatusWithDefaults() *IngestersComplianceScanStatus {
	this := IngestersComplianceScanStatus{}
	return &this
}

// GetScanId returns the ScanId field value if set, zero value otherwise.
func (o *IngestersComplianceScanStatus) GetScanId() string {
	if o == nil || IsNil(o.ScanId) {
		var ret string
		return ret
	}
	return *o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceScanStatus) GetScanIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScanId) {
		return nil, false
	}
	return o.ScanId, true
}

// HasScanId returns a boolean if a field has been set.
func (o *IngestersComplianceScanStatus) HasScanId() bool {
	if o != nil && !IsNil(o.ScanId) {
		return true
	}

	return false
}

// SetScanId gets a reference to the given string and assigns it to the ScanId field.
func (o *IngestersComplianceScanStatus) SetScanId(v string) {
	o.ScanId = &v
}

// GetScanMessage returns the ScanMessage field value if set, zero value otherwise.
func (o *IngestersComplianceScanStatus) GetScanMessage() string {
	if o == nil || IsNil(o.ScanMessage) {
		var ret string
		return ret
	}
	return *o.ScanMessage
}

// GetScanMessageOk returns a tuple with the ScanMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceScanStatus) GetScanMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ScanMessage) {
		return nil, false
	}
	return o.ScanMessage, true
}

// HasScanMessage returns a boolean if a field has been set.
func (o *IngestersComplianceScanStatus) HasScanMessage() bool {
	if o != nil && !IsNil(o.ScanMessage) {
		return true
	}

	return false
}

// SetScanMessage gets a reference to the given string and assigns it to the ScanMessage field.
func (o *IngestersComplianceScanStatus) SetScanMessage(v string) {
	o.ScanMessage = &v
}

// GetScanStatus returns the ScanStatus field value if set, zero value otherwise.
func (o *IngestersComplianceScanStatus) GetScanStatus() string {
	if o == nil || IsNil(o.ScanStatus) {
		var ret string
		return ret
	}
	return *o.ScanStatus
}

// GetScanStatusOk returns a tuple with the ScanStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceScanStatus) GetScanStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ScanStatus) {
		return nil, false
	}
	return o.ScanStatus, true
}

// HasScanStatus returns a boolean if a field has been set.
func (o *IngestersComplianceScanStatus) HasScanStatus() bool {
	if o != nil && !IsNil(o.ScanStatus) {
		return true
	}

	return false
}

// SetScanStatus gets a reference to the given string and assigns it to the ScanStatus field.
func (o *IngestersComplianceScanStatus) SetScanStatus(v string) {
	o.ScanStatus = &v
}

func (o IngestersComplianceScanStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestersComplianceScanStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScanId) {
		toSerialize["scan_id"] = o.ScanId
	}
	if !IsNil(o.ScanMessage) {
		toSerialize["scan_message"] = o.ScanMessage
	}
	if !IsNil(o.ScanStatus) {
		toSerialize["scan_status"] = o.ScanStatus
	}
	return toSerialize, nil
}

type NullableIngestersComplianceScanStatus struct {
	value *IngestersComplianceScanStatus
	isSet bool
}

func (v NullableIngestersComplianceScanStatus) Get() *IngestersComplianceScanStatus {
	return v.value
}

func (v *NullableIngestersComplianceScanStatus) Set(val *IngestersComplianceScanStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestersComplianceScanStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestersComplianceScanStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestersComplianceScanStatus(val *IngestersComplianceScanStatus) *NullableIngestersComplianceScanStatus {
	return &NullableIngestersComplianceScanStatus{value: val, isSet: true}
}

func (v NullableIngestersComplianceScanStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestersComplianceScanStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


