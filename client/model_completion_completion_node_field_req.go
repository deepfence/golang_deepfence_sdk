/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.3.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CompletionCompletionNodeFieldReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompletionCompletionNodeFieldReq{}

// CompletionCompletionNodeFieldReq struct for CompletionCompletionNodeFieldReq
type CompletionCompletionNodeFieldReq struct {
	Completion string `json:"completion"`
	FieldName string `json:"field_name"`
	Filters *ReportersFieldsFilters `json:"filters,omitempty"`
	ScanId *string `json:"scan_id,omitempty"`
	Window ModelFetchWindow `json:"window"`
}

type _CompletionCompletionNodeFieldReq CompletionCompletionNodeFieldReq

// NewCompletionCompletionNodeFieldReq instantiates a new CompletionCompletionNodeFieldReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompletionCompletionNodeFieldReq(completion string, fieldName string, window ModelFetchWindow) *CompletionCompletionNodeFieldReq {
	this := CompletionCompletionNodeFieldReq{}
	this.Completion = completion
	this.FieldName = fieldName
	this.Window = window
	return &this
}

// NewCompletionCompletionNodeFieldReqWithDefaults instantiates a new CompletionCompletionNodeFieldReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompletionCompletionNodeFieldReqWithDefaults() *CompletionCompletionNodeFieldReq {
	this := CompletionCompletionNodeFieldReq{}
	return &this
}

// GetCompletion returns the Completion field value
func (o *CompletionCompletionNodeFieldReq) GetCompletion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Completion
}

// GetCompletionOk returns a tuple with the Completion field value
// and a boolean to check if the value has been set.
func (o *CompletionCompletionNodeFieldReq) GetCompletionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Completion, true
}

// SetCompletion sets field value
func (o *CompletionCompletionNodeFieldReq) SetCompletion(v string) {
	o.Completion = v
}

// GetFieldName returns the FieldName field value
func (o *CompletionCompletionNodeFieldReq) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *CompletionCompletionNodeFieldReq) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *CompletionCompletionNodeFieldReq) SetFieldName(v string) {
	o.FieldName = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *CompletionCompletionNodeFieldReq) GetFilters() ReportersFieldsFilters {
	if o == nil || IsNil(o.Filters) {
		var ret ReportersFieldsFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletionCompletionNodeFieldReq) GetFiltersOk() (*ReportersFieldsFilters, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *CompletionCompletionNodeFieldReq) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given ReportersFieldsFilters and assigns it to the Filters field.
func (o *CompletionCompletionNodeFieldReq) SetFilters(v ReportersFieldsFilters) {
	o.Filters = &v
}

// GetScanId returns the ScanId field value if set, zero value otherwise.
func (o *CompletionCompletionNodeFieldReq) GetScanId() string {
	if o == nil || IsNil(o.ScanId) {
		var ret string
		return ret
	}
	return *o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletionCompletionNodeFieldReq) GetScanIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScanId) {
		return nil, false
	}
	return o.ScanId, true
}

// HasScanId returns a boolean if a field has been set.
func (o *CompletionCompletionNodeFieldReq) HasScanId() bool {
	if o != nil && !IsNil(o.ScanId) {
		return true
	}

	return false
}

// SetScanId gets a reference to the given string and assigns it to the ScanId field.
func (o *CompletionCompletionNodeFieldReq) SetScanId(v string) {
	o.ScanId = &v
}

// GetWindow returns the Window field value
func (o *CompletionCompletionNodeFieldReq) GetWindow() ModelFetchWindow {
	if o == nil {
		var ret ModelFetchWindow
		return ret
	}

	return o.Window
}

// GetWindowOk returns a tuple with the Window field value
// and a boolean to check if the value has been set.
func (o *CompletionCompletionNodeFieldReq) GetWindowOk() (*ModelFetchWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Window, true
}

// SetWindow sets field value
func (o *CompletionCompletionNodeFieldReq) SetWindow(v ModelFetchWindow) {
	o.Window = v
}

func (o CompletionCompletionNodeFieldReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompletionCompletionNodeFieldReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["completion"] = o.Completion
	toSerialize["field_name"] = o.FieldName
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.ScanId) {
		toSerialize["scan_id"] = o.ScanId
	}
	toSerialize["window"] = o.Window
	return toSerialize, nil
}

func (o *CompletionCompletionNodeFieldReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"completion",
		"field_name",
		"window",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCompletionCompletionNodeFieldReq := _CompletionCompletionNodeFieldReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCompletionCompletionNodeFieldReq)

	if err != nil {
		return err
	}

	*o = CompletionCompletionNodeFieldReq(varCompletionCompletionNodeFieldReq)

	return err
}

type NullableCompletionCompletionNodeFieldReq struct {
	value *CompletionCompletionNodeFieldReq
	isSet bool
}

func (v NullableCompletionCompletionNodeFieldReq) Get() *CompletionCompletionNodeFieldReq {
	return v.value
}

func (v *NullableCompletionCompletionNodeFieldReq) Set(val *CompletionCompletionNodeFieldReq) {
	v.value = val
	v.isSet = true
}

func (v NullableCompletionCompletionNodeFieldReq) IsSet() bool {
	return v.isSet
}

func (v *NullableCompletionCompletionNodeFieldReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompletionCompletionNodeFieldReq(val *CompletionCompletionNodeFieldReq) *NullableCompletionCompletionNodeFieldReq {
	return &NullableCompletionCompletionNodeFieldReq{value: val, isSet: true}
}

func (v NullableCompletionCompletionNodeFieldReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompletionCompletionNodeFieldReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


