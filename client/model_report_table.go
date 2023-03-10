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

// checks if the ReportTable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportTable{}

// ReportTable struct for ReportTable
type ReportTable struct {
	Columns []ReportColumn `json:"columns,omitempty"`
	Id *string `json:"id,omitempty"`
	Label *string `json:"label,omitempty"`
	Rows []ReportRow `json:"rows,omitempty"`
	TruncationCount *int32 `json:"truncationCount,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewReportTable instantiates a new ReportTable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportTable() *ReportTable {
	this := ReportTable{}
	return &this
}

// NewReportTableWithDefaults instantiates a new ReportTable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportTableWithDefaults() *ReportTable {
	this := ReportTable{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportTable) GetColumns() []ReportColumn {
	if o == nil {
		var ret []ReportColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportTable) GetColumnsOk() ([]ReportColumn, bool) {
	if o == nil || IsNil(o.Columns) {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *ReportTable) HasColumns() bool {
	if o != nil && IsNil(o.Columns) {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []ReportColumn and assigns it to the Columns field.
func (o *ReportTable) SetColumns(v []ReportColumn) {
	o.Columns = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReportTable) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportTable) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReportTable) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReportTable) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ReportTable) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportTable) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ReportTable) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ReportTable) SetLabel(v string) {
	o.Label = &v
}

// GetRows returns the Rows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportTable) GetRows() []ReportRow {
	if o == nil {
		var ret []ReportRow
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportTable) GetRowsOk() ([]ReportRow, bool) {
	if o == nil || IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *ReportTable) HasRows() bool {
	if o != nil && IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []ReportRow and assigns it to the Rows field.
func (o *ReportTable) SetRows(v []ReportRow) {
	o.Rows = v
}

// GetTruncationCount returns the TruncationCount field value if set, zero value otherwise.
func (o *ReportTable) GetTruncationCount() int32 {
	if o == nil || IsNil(o.TruncationCount) {
		var ret int32
		return ret
	}
	return *o.TruncationCount
}

// GetTruncationCountOk returns a tuple with the TruncationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportTable) GetTruncationCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TruncationCount) {
		return nil, false
	}
	return o.TruncationCount, true
}

// HasTruncationCount returns a boolean if a field has been set.
func (o *ReportTable) HasTruncationCount() bool {
	if o != nil && !IsNil(o.TruncationCount) {
		return true
	}

	return false
}

// SetTruncationCount gets a reference to the given int32 and assigns it to the TruncationCount field.
func (o *ReportTable) SetTruncationCount(v int32) {
	o.TruncationCount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ReportTable) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportTable) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ReportTable) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ReportTable) SetType(v string) {
	o.Type = &v
}

func (o ReportTable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportTable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if o.Rows != nil {
		toSerialize["rows"] = o.Rows
	}
	if !IsNil(o.TruncationCount) {
		toSerialize["truncationCount"] = o.TruncationCount
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableReportTable struct {
	value *ReportTable
	isSet bool
}

func (v NullableReportTable) Get() *ReportTable {
	return v.value
}

func (v *NullableReportTable) Set(val *ReportTable) {
	v.value = val
	v.isSet = true
}

func (v NullableReportTable) IsSet() bool {
	return v.isSet
}

func (v *NullableReportTable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportTable(val *ReportTable) *NullableReportTable {
	return &NullableReportTable{value: val, isSet: true}
}

func (v NullableReportTable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportTable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


