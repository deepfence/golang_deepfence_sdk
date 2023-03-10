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

// checks if the ReportMetadataRow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportMetadataRow{}

// ReportMetadataRow struct for ReportMetadataRow
type ReportMetadataRow struct {
	DataType *string `json:"dataType,omitempty"`
	Id *string `json:"id,omitempty"`
	Label *string `json:"label,omitempty"`
	Priority *float32 `json:"priority,omitempty"`
	Truncate *int32 `json:"truncate,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewReportMetadataRow instantiates a new ReportMetadataRow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportMetadataRow() *ReportMetadataRow {
	this := ReportMetadataRow{}
	return &this
}

// NewReportMetadataRowWithDefaults instantiates a new ReportMetadataRow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportMetadataRowWithDefaults() *ReportMetadataRow {
	this := ReportMetadataRow{}
	return &this
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *ReportMetadataRow) GetDataType() string {
	if o == nil || IsNil(o.DataType) {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportMetadataRow) GetDataTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DataType) {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *ReportMetadataRow) HasDataType() bool {
	if o != nil && !IsNil(o.DataType) {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *ReportMetadataRow) SetDataType(v string) {
	o.DataType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReportMetadataRow) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportMetadataRow) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReportMetadataRow) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReportMetadataRow) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ReportMetadataRow) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportMetadataRow) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ReportMetadataRow) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ReportMetadataRow) SetLabel(v string) {
	o.Label = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ReportMetadataRow) GetPriority() float32 {
	if o == nil || IsNil(o.Priority) {
		var ret float32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportMetadataRow) GetPriorityOk() (*float32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ReportMetadataRow) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given float32 and assigns it to the Priority field.
func (o *ReportMetadataRow) SetPriority(v float32) {
	o.Priority = &v
}

// GetTruncate returns the Truncate field value if set, zero value otherwise.
func (o *ReportMetadataRow) GetTruncate() int32 {
	if o == nil || IsNil(o.Truncate) {
		var ret int32
		return ret
	}
	return *o.Truncate
}

// GetTruncateOk returns a tuple with the Truncate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportMetadataRow) GetTruncateOk() (*int32, bool) {
	if o == nil || IsNil(o.Truncate) {
		return nil, false
	}
	return o.Truncate, true
}

// HasTruncate returns a boolean if a field has been set.
func (o *ReportMetadataRow) HasTruncate() bool {
	if o != nil && !IsNil(o.Truncate) {
		return true
	}

	return false
}

// SetTruncate gets a reference to the given int32 and assigns it to the Truncate field.
func (o *ReportMetadataRow) SetTruncate(v int32) {
	o.Truncate = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ReportMetadataRow) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportMetadataRow) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ReportMetadataRow) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ReportMetadataRow) SetValue(v string) {
	o.Value = &v
}

func (o ReportMetadataRow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportMetadataRow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataType) {
		toSerialize["dataType"] = o.DataType
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Truncate) {
		toSerialize["truncate"] = o.Truncate
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableReportMetadataRow struct {
	value *ReportMetadataRow
	isSet bool
}

func (v NullableReportMetadataRow) Get() *ReportMetadataRow {
	return v.value
}

func (v *NullableReportMetadataRow) Set(val *ReportMetadataRow) {
	v.value = val
	v.isSet = true
}

func (v NullableReportMetadataRow) IsSet() bool {
	return v.isSet
}

func (v *NullableReportMetadataRow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportMetadataRow(val *ReportMetadataRow) *NullableReportMetadataRow {
	return &NullableReportMetadataRow{value: val, isSet: true}
}

func (v NullableReportMetadataRow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportMetadataRow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


