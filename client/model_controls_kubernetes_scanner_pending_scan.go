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

// checks if the ControlsKubernetesScannerPendingScan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControlsKubernetesScannerPendingScan{}

// ControlsKubernetesScannerPendingScan struct for ControlsKubernetesScannerPendingScan
type ControlsKubernetesScannerPendingScan struct {
	AccountId *string `json:"account_id,omitempty"`
	Controls []string `json:"controls,omitempty"`
	ScanId *string `json:"scan_id,omitempty"`
	ScanType *string `json:"scan_type,omitempty"`
}

// NewControlsKubernetesScannerPendingScan instantiates a new ControlsKubernetesScannerPendingScan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControlsKubernetesScannerPendingScan() *ControlsKubernetesScannerPendingScan {
	this := ControlsKubernetesScannerPendingScan{}
	return &this
}

// NewControlsKubernetesScannerPendingScanWithDefaults instantiates a new ControlsKubernetesScannerPendingScan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControlsKubernetesScannerPendingScanWithDefaults() *ControlsKubernetesScannerPendingScan {
	this := ControlsKubernetesScannerPendingScan{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ControlsKubernetesScannerPendingScan) GetAccountId() string {
	if o == nil || isNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsKubernetesScannerPendingScan) GetAccountIdOk() (*string, bool) {
	if o == nil || isNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ControlsKubernetesScannerPendingScan) HasAccountId() bool {
	if o != nil && !isNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ControlsKubernetesScannerPendingScan) SetAccountId(v string) {
	o.AccountId = &v
}

// GetControls returns the Controls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ControlsKubernetesScannerPendingScan) GetControls() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Controls
}

// GetControlsOk returns a tuple with the Controls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ControlsKubernetesScannerPendingScan) GetControlsOk() ([]string, bool) {
	if o == nil || isNil(o.Controls) {
		return nil, false
	}
	return o.Controls, true
}

// HasControls returns a boolean if a field has been set.
func (o *ControlsKubernetesScannerPendingScan) HasControls() bool {
	if o != nil && isNil(o.Controls) {
		return true
	}

	return false
}

// SetControls gets a reference to the given []string and assigns it to the Controls field.
func (o *ControlsKubernetesScannerPendingScan) SetControls(v []string) {
	o.Controls = v
}

// GetScanId returns the ScanId field value if set, zero value otherwise.
func (o *ControlsKubernetesScannerPendingScan) GetScanId() string {
	if o == nil || isNil(o.ScanId) {
		var ret string
		return ret
	}
	return *o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsKubernetesScannerPendingScan) GetScanIdOk() (*string, bool) {
	if o == nil || isNil(o.ScanId) {
		return nil, false
	}
	return o.ScanId, true
}

// HasScanId returns a boolean if a field has been set.
func (o *ControlsKubernetesScannerPendingScan) HasScanId() bool {
	if o != nil && !isNil(o.ScanId) {
		return true
	}

	return false
}

// SetScanId gets a reference to the given string and assigns it to the ScanId field.
func (o *ControlsKubernetesScannerPendingScan) SetScanId(v string) {
	o.ScanId = &v
}

// GetScanType returns the ScanType field value if set, zero value otherwise.
func (o *ControlsKubernetesScannerPendingScan) GetScanType() string {
	if o == nil || isNil(o.ScanType) {
		var ret string
		return ret
	}
	return *o.ScanType
}

// GetScanTypeOk returns a tuple with the ScanType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsKubernetesScannerPendingScan) GetScanTypeOk() (*string, bool) {
	if o == nil || isNil(o.ScanType) {
		return nil, false
	}
	return o.ScanType, true
}

// HasScanType returns a boolean if a field has been set.
func (o *ControlsKubernetesScannerPendingScan) HasScanType() bool {
	if o != nil && !isNil(o.ScanType) {
		return true
	}

	return false
}

// SetScanType gets a reference to the given string and assigns it to the ScanType field.
func (o *ControlsKubernetesScannerPendingScan) SetScanType(v string) {
	o.ScanType = &v
}

func (o ControlsKubernetesScannerPendingScan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControlsKubernetesScannerPendingScan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if o.Controls != nil {
		toSerialize["controls"] = o.Controls
	}
	if !isNil(o.ScanId) {
		toSerialize["scan_id"] = o.ScanId
	}
	if !isNil(o.ScanType) {
		toSerialize["scan_type"] = o.ScanType
	}
	return toSerialize, nil
}

type NullableControlsKubernetesScannerPendingScan struct {
	value *ControlsKubernetesScannerPendingScan
	isSet bool
}

func (v NullableControlsKubernetesScannerPendingScan) Get() *ControlsKubernetesScannerPendingScan {
	return v.value
}

func (v *NullableControlsKubernetesScannerPendingScan) Set(val *ControlsKubernetesScannerPendingScan) {
	v.value = val
	v.isSet = true
}

func (v NullableControlsKubernetesScannerPendingScan) IsSet() bool {
	return v.isSet
}

func (v *NullableControlsKubernetesScannerPendingScan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControlsKubernetesScannerPendingScan(val *ControlsKubernetesScannerPendingScan) *NullableControlsKubernetesScannerPendingScan {
	return &NullableControlsKubernetesScannerPendingScan{value: val, isSet: true}
}

func (v NullableControlsKubernetesScannerPendingScan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControlsKubernetesScannerPendingScan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

