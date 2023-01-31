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

// checks if the ModelScanTriggerResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelScanTriggerResp{}

// ModelScanTriggerResp struct for ModelScanTriggerResp
type ModelScanTriggerResp struct {
	BulkScanId string `json:"bulk_scan_id"`
	ScanIds []string `json:"scan_ids"`
}

// NewModelScanTriggerResp instantiates a new ModelScanTriggerResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelScanTriggerResp(bulkScanId string, scanIds []string) *ModelScanTriggerResp {
	this := ModelScanTriggerResp{}
	this.BulkScanId = bulkScanId
	this.ScanIds = scanIds
	return &this
}

// NewModelScanTriggerRespWithDefaults instantiates a new ModelScanTriggerResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelScanTriggerRespWithDefaults() *ModelScanTriggerResp {
	this := ModelScanTriggerResp{}
	return &this
}

// GetBulkScanId returns the BulkScanId field value
func (o *ModelScanTriggerResp) GetBulkScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BulkScanId
}

// GetBulkScanIdOk returns a tuple with the BulkScanId field value
// and a boolean to check if the value has been set.
func (o *ModelScanTriggerResp) GetBulkScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BulkScanId, true
}

// SetBulkScanId sets field value
func (o *ModelScanTriggerResp) SetBulkScanId(v string) {
	o.BulkScanId = v
}

// GetScanIds returns the ScanIds field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ModelScanTriggerResp) GetScanIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ScanIds
}

// GetScanIdsOk returns a tuple with the ScanIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelScanTriggerResp) GetScanIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ScanIds) {
		return nil, false
	}
	return o.ScanIds, true
}

// SetScanIds sets field value
func (o *ModelScanTriggerResp) SetScanIds(v []string) {
	o.ScanIds = v
}

func (o ModelScanTriggerResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelScanTriggerResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bulk_scan_id"] = o.BulkScanId
	if o.ScanIds != nil {
		toSerialize["scan_ids"] = o.ScanIds
	}
	return toSerialize, nil
}

type NullableModelScanTriggerResp struct {
	value *ModelScanTriggerResp
	isSet bool
}

func (v NullableModelScanTriggerResp) Get() *ModelScanTriggerResp {
	return v.value
}

func (v *NullableModelScanTriggerResp) Set(val *ModelScanTriggerResp) {
	v.value = val
	v.isSet = true
}

func (v NullableModelScanTriggerResp) IsSet() bool {
	return v.isSet
}

func (v *NullableModelScanTriggerResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelScanTriggerResp(val *ModelScanTriggerResp) *NullableModelScanTriggerResp {
	return &NullableModelScanTriggerResp{value: val, isSet: true}
}

func (v NullableModelScanTriggerResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelScanTriggerResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


