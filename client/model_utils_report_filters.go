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
	"bytes"
	"fmt"
)

// checks if the UtilsReportFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UtilsReportFilters{}

// UtilsReportFilters struct for UtilsReportFilters
type UtilsReportFilters struct {
	AdvancedReportFilters *UtilsAdvancedReportFilters `json:"advanced_report_filters,omitempty"`
	IncludeDeadNodes *bool `json:"include_dead_nodes,omitempty"`
	MostExploitableReport *bool `json:"most_exploitable_report,omitempty"`
	NodeType string `json:"node_type"`
	ScanId *string `json:"scan_id,omitempty"`
	ScanType string `json:"scan_type"`
	SeverityOrCheckType []string `json:"severity_or_check_type,omitempty"`
}

type _UtilsReportFilters UtilsReportFilters

// NewUtilsReportFilters instantiates a new UtilsReportFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUtilsReportFilters(nodeType string, scanType string) *UtilsReportFilters {
	this := UtilsReportFilters{}
	this.NodeType = nodeType
	this.ScanType = scanType
	return &this
}

// NewUtilsReportFiltersWithDefaults instantiates a new UtilsReportFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUtilsReportFiltersWithDefaults() *UtilsReportFilters {
	this := UtilsReportFilters{}
	return &this
}

// GetAdvancedReportFilters returns the AdvancedReportFilters field value if set, zero value otherwise.
func (o *UtilsReportFilters) GetAdvancedReportFilters() UtilsAdvancedReportFilters {
	if o == nil || IsNil(o.AdvancedReportFilters) {
		var ret UtilsAdvancedReportFilters
		return ret
	}
	return *o.AdvancedReportFilters
}

// GetAdvancedReportFiltersOk returns a tuple with the AdvancedReportFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsReportFilters) GetAdvancedReportFiltersOk() (*UtilsAdvancedReportFilters, bool) {
	if o == nil || IsNil(o.AdvancedReportFilters) {
		return nil, false
	}
	return o.AdvancedReportFilters, true
}

// HasAdvancedReportFilters returns a boolean if a field has been set.
func (o *UtilsReportFilters) HasAdvancedReportFilters() bool {
	if o != nil && !IsNil(o.AdvancedReportFilters) {
		return true
	}

	return false
}

// SetAdvancedReportFilters gets a reference to the given UtilsAdvancedReportFilters and assigns it to the AdvancedReportFilters field.
func (o *UtilsReportFilters) SetAdvancedReportFilters(v UtilsAdvancedReportFilters) {
	o.AdvancedReportFilters = &v
}

// GetIncludeDeadNodes returns the IncludeDeadNodes field value if set, zero value otherwise.
func (o *UtilsReportFilters) GetIncludeDeadNodes() bool {
	if o == nil || IsNil(o.IncludeDeadNodes) {
		var ret bool
		return ret
	}
	return *o.IncludeDeadNodes
}

// GetIncludeDeadNodesOk returns a tuple with the IncludeDeadNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsReportFilters) GetIncludeDeadNodesOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeDeadNodes) {
		return nil, false
	}
	return o.IncludeDeadNodes, true
}

// HasIncludeDeadNodes returns a boolean if a field has been set.
func (o *UtilsReportFilters) HasIncludeDeadNodes() bool {
	if o != nil && !IsNil(o.IncludeDeadNodes) {
		return true
	}

	return false
}

// SetIncludeDeadNodes gets a reference to the given bool and assigns it to the IncludeDeadNodes field.
func (o *UtilsReportFilters) SetIncludeDeadNodes(v bool) {
	o.IncludeDeadNodes = &v
}

// GetMostExploitableReport returns the MostExploitableReport field value if set, zero value otherwise.
func (o *UtilsReportFilters) GetMostExploitableReport() bool {
	if o == nil || IsNil(o.MostExploitableReport) {
		var ret bool
		return ret
	}
	return *o.MostExploitableReport
}

// GetMostExploitableReportOk returns a tuple with the MostExploitableReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsReportFilters) GetMostExploitableReportOk() (*bool, bool) {
	if o == nil || IsNil(o.MostExploitableReport) {
		return nil, false
	}
	return o.MostExploitableReport, true
}

// HasMostExploitableReport returns a boolean if a field has been set.
func (o *UtilsReportFilters) HasMostExploitableReport() bool {
	if o != nil && !IsNil(o.MostExploitableReport) {
		return true
	}

	return false
}

// SetMostExploitableReport gets a reference to the given bool and assigns it to the MostExploitableReport field.
func (o *UtilsReportFilters) SetMostExploitableReport(v bool) {
	o.MostExploitableReport = &v
}

// GetNodeType returns the NodeType field value
func (o *UtilsReportFilters) GetNodeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *UtilsReportFilters) GetNodeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value
func (o *UtilsReportFilters) SetNodeType(v string) {
	o.NodeType = v
}

// GetScanId returns the ScanId field value if set, zero value otherwise.
func (o *UtilsReportFilters) GetScanId() string {
	if o == nil || IsNil(o.ScanId) {
		var ret string
		return ret
	}
	return *o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsReportFilters) GetScanIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScanId) {
		return nil, false
	}
	return o.ScanId, true
}

// HasScanId returns a boolean if a field has been set.
func (o *UtilsReportFilters) HasScanId() bool {
	if o != nil && !IsNil(o.ScanId) {
		return true
	}

	return false
}

// SetScanId gets a reference to the given string and assigns it to the ScanId field.
func (o *UtilsReportFilters) SetScanId(v string) {
	o.ScanId = &v
}

// GetScanType returns the ScanType field value
func (o *UtilsReportFilters) GetScanType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScanType
}

// GetScanTypeOk returns a tuple with the ScanType field value
// and a boolean to check if the value has been set.
func (o *UtilsReportFilters) GetScanTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScanType, true
}

// SetScanType sets field value
func (o *UtilsReportFilters) SetScanType(v string) {
	o.ScanType = v
}

// GetSeverityOrCheckType returns the SeverityOrCheckType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UtilsReportFilters) GetSeverityOrCheckType() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SeverityOrCheckType
}

// GetSeverityOrCheckTypeOk returns a tuple with the SeverityOrCheckType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UtilsReportFilters) GetSeverityOrCheckTypeOk() ([]string, bool) {
	if o == nil || IsNil(o.SeverityOrCheckType) {
		return nil, false
	}
	return o.SeverityOrCheckType, true
}

// HasSeverityOrCheckType returns a boolean if a field has been set.
func (o *UtilsReportFilters) HasSeverityOrCheckType() bool {
	if o != nil && !IsNil(o.SeverityOrCheckType) {
		return true
	}

	return false
}

// SetSeverityOrCheckType gets a reference to the given []string and assigns it to the SeverityOrCheckType field.
func (o *UtilsReportFilters) SetSeverityOrCheckType(v []string) {
	o.SeverityOrCheckType = v
}

func (o UtilsReportFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UtilsReportFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdvancedReportFilters) {
		toSerialize["advanced_report_filters"] = o.AdvancedReportFilters
	}
	if !IsNil(o.IncludeDeadNodes) {
		toSerialize["include_dead_nodes"] = o.IncludeDeadNodes
	}
	if !IsNil(o.MostExploitableReport) {
		toSerialize["most_exploitable_report"] = o.MostExploitableReport
	}
	toSerialize["node_type"] = o.NodeType
	if !IsNil(o.ScanId) {
		toSerialize["scan_id"] = o.ScanId
	}
	toSerialize["scan_type"] = o.ScanType
	if o.SeverityOrCheckType != nil {
		toSerialize["severity_or_check_type"] = o.SeverityOrCheckType
	}
	return toSerialize, nil
}

func (o *UtilsReportFilters) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"node_type",
		"scan_type",
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

	varUtilsReportFilters := _UtilsReportFilters{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUtilsReportFilters)

	if err != nil {
		return err
	}

	*o = UtilsReportFilters(varUtilsReportFilters)

	return err
}

type NullableUtilsReportFilters struct {
	value *UtilsReportFilters
	isSet bool
}

func (v NullableUtilsReportFilters) Get() *UtilsReportFilters {
	return v.value
}

func (v *NullableUtilsReportFilters) Set(val *UtilsReportFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableUtilsReportFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableUtilsReportFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtilsReportFilters(val *UtilsReportFilters) *NullableUtilsReportFilters {
	return &NullableUtilsReportFilters{value: val, isSet: true}
}

func (v NullableUtilsReportFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtilsReportFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


