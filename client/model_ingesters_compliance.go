/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.4.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the IngestersCompliance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestersCompliance{}

// IngestersCompliance struct for IngestersCompliance
type IngestersCompliance struct {
	ComplianceCheckType *string `json:"compliance_check_type,omitempty"`
	Description *string `json:"description,omitempty"`
	NodeId *string `json:"node_id,omitempty"`
	NodeType *string `json:"node_type,omitempty"`
	RemediationAnsible *string `json:"remediation_ansible,omitempty"`
	RemediationPuppet *string `json:"remediation_puppet,omitempty"`
	RemediationScript *string `json:"remediation_script,omitempty"`
	Resource *string `json:"resource,omitempty"`
	ScanId *string `json:"scan_id,omitempty"`
	Status *string `json:"status,omitempty"`
	TestCategory *string `json:"test_category,omitempty"`
	TestDesc *string `json:"test_desc,omitempty"`
	TestNumber *string `json:"test_number,omitempty"`
	TestRationale *string `json:"test_rationale,omitempty"`
	TestSeverity *string `json:"test_severity,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewIngestersCompliance instantiates a new IngestersCompliance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestersCompliance() *IngestersCompliance {
	this := IngestersCompliance{}
	return &this
}

// NewIngestersComplianceWithDefaults instantiates a new IngestersCompliance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestersComplianceWithDefaults() *IngestersCompliance {
	this := IngestersCompliance{}
	return &this
}

// GetComplianceCheckType returns the ComplianceCheckType field value if set, zero value otherwise.
func (o *IngestersCompliance) GetComplianceCheckType() string {
	if o == nil || IsNil(o.ComplianceCheckType) {
		var ret string
		return ret
	}
	return *o.ComplianceCheckType
}

// GetComplianceCheckTypeOk returns a tuple with the ComplianceCheckType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCompliance) GetComplianceCheckTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ComplianceCheckType) {
		return nil, false
	}
	return o.ComplianceCheckType, true
}

// HasComplianceCheckType returns a boolean if a field has been set.
func (o *IngestersCompliance) HasComplianceCheckType() bool {
	if o != nil && !IsNil(o.ComplianceCheckType) {
		return true
	}

	return false
}

// SetComplianceCheckType gets a reference to the given string and assigns it to the ComplianceCheckType field.
func (o *IngestersCompliance) SetComplianceCheckType(v string) {
	o.ComplianceCheckType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IngestersCompliance) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCompliance) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IngestersCompliance) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IngestersCompliance) SetDescription(v string) {
	o.Description = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *IngestersCompliance) GetNodeId() string {
	if o == nil || IsNil(o.NodeId) {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCompliance) GetNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *IngestersCompliance) HasNodeId() bool {
	if o != nil && !IsNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *IngestersCompliance) SetNodeId(v string) {
	o.NodeId = &v
}

// GetNodeType returns the NodeType field value if set, zero value otherwise.
func (o *IngestersCompliance) GetNodeType() string {
	if o == nil || IsNil(o.NodeType) {
		var ret string
		return ret
	}
	return *o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCompliance) GetNodeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NodeType) {
		return nil, false
	}
	return o.NodeType, true
}

// HasNodeType returns a boolean if a field has been set.
func (o *IngestersCompliance) HasNodeType() bool {
	if o != nil && !IsNil(o.NodeType) {
		return true
	}

	return false
}

// SetNodeType gets a reference to the given string and assigns it to the NodeType field.
func (o *IngestersCompliance) SetNodeType(v string) {
	o.NodeType = &v
}

// GetRemediationAnsible returns the RemediationAnsible field value if set, zero value otherwise.
func (o *IngestersCompliance) GetRemediationAnsible() string {
	if o == nil || IsNil(o.RemediationAnsible) {
		var ret string
		return ret
	}
	return *o.RemediationAnsible
}

// GetRemediationAnsibleOk returns a tuple with the RemediationAnsible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCompliance) GetRemediationAnsibleOk() (*string, bool) {
	if o == nil || IsNil(o.RemediationAnsible) {
		return nil, false
	}
	return o.RemediationAnsible, true
}

// HasRemediationAnsible returns a boolean if a field has been set.
func (o *IngestersCompliance) HasRemediationAnsible() bool {
	if o != nil && !IsNil(o.RemediationAnsible) {
		return true
	}

	return false
}

// SetRemediationAnsible gets a reference to the given string and assigns it to the RemediationAnsible field.
func (o *IngestersCompliance) SetRemediationAnsible(v string) {
	o.RemediationAnsible = &v
}

// GetRemediationPuppet returns the RemediationPuppet field value if set, zero value otherwise.
func (o *IngestersCompliance) GetRemediationPuppet() string {
	if o == nil || IsNil(o.RemediationPuppet) {
		var ret string
		return ret
	}
	return *o.RemediationPuppet
}

// GetRemediationPuppetOk returns a tuple with the RemediationPuppet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCompliance) GetRemediationPuppetOk() (*string, bool) {
	if o == nil || IsNil(o.RemediationPuppet) {
		return nil, false
	}
	return o.RemediationPuppet, true
}

// HasRemediationPuppet returns a boolean if a field has been set.
func (o *IngestersCompliance) HasRemediationPuppet() bool {
	if o != nil && !IsNil(o.RemediationPuppet) {
		return true
	}

	return false
}

// SetRemediationPuppet gets a reference to the given string and assigns it to the RemediationPuppet field.
func (o *IngestersCompliance) SetRemediationPuppet(v string) {
	o.RemediationPuppet = &v
}

// GetRemediationScript returns the RemediationScript field value if set, zero value otherwise.
func (o *IngestersCompliance) GetRemediationScript() string {
	if o == nil || IsNil(o.RemediationScript) {
		var ret string
		return ret
	}
	return *o.RemediationScript
}

// GetRemediationScriptOk returns a tuple with the RemediationScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCompliance) GetRemediationScriptOk() (*string, bool) {
	if o == nil || IsNil(o.RemediationScript) {
		return nil, false
	}
	return o.RemediationScript, true
}

// HasRemediationScript returns a boolean if a field has been set.
func (o *IngestersCompliance) HasRemediationScript() bool {
	if o != nil && !IsNil(o.RemediationScript) {
		return true
	}

	return false
}

// SetRemediationScript gets a reference to the given string and assigns it to the RemediationScript field.
func (o *IngestersCompliance) SetRemediationScript(v string) {
	o.RemediationScript = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *IngestersCompliance) GetResource() string {
	if o == nil || IsNil(o.Resource) {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCompliance) GetResourceOk() (*string, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *IngestersCompliance) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *IngestersCompliance) SetResource(v string) {
	o.Resource = &v
}

// GetScanId returns the ScanId field value if set, zero value otherwise.
func (o *IngestersCompliance) GetScanId() string {
	if o == nil || IsNil(o.ScanId) {
		var ret string
		return ret
	}
	return *o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCompliance) GetScanIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScanId) {
		return nil, false
	}
	return o.ScanId, true
}

// HasScanId returns a boolean if a field has been set.
func (o *IngestersCompliance) HasScanId() bool {
	if o != nil && !IsNil(o.ScanId) {
		return true
	}

	return false
}

// SetScanId gets a reference to the given string and assigns it to the ScanId field.
func (o *IngestersCompliance) SetScanId(v string) {
	o.ScanId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IngestersCompliance) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCompliance) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IngestersCompliance) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IngestersCompliance) SetStatus(v string) {
	o.Status = &v
}

// GetTestCategory returns the TestCategory field value if set, zero value otherwise.
func (o *IngestersCompliance) GetTestCategory() string {
	if o == nil || IsNil(o.TestCategory) {
		var ret string
		return ret
	}
	return *o.TestCategory
}

// GetTestCategoryOk returns a tuple with the TestCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCompliance) GetTestCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.TestCategory) {
		return nil, false
	}
	return o.TestCategory, true
}

// HasTestCategory returns a boolean if a field has been set.
func (o *IngestersCompliance) HasTestCategory() bool {
	if o != nil && !IsNil(o.TestCategory) {
		return true
	}

	return false
}

// SetTestCategory gets a reference to the given string and assigns it to the TestCategory field.
func (o *IngestersCompliance) SetTestCategory(v string) {
	o.TestCategory = &v
}

// GetTestDesc returns the TestDesc field value if set, zero value otherwise.
func (o *IngestersCompliance) GetTestDesc() string {
	if o == nil || IsNil(o.TestDesc) {
		var ret string
		return ret
	}
	return *o.TestDesc
}

// GetTestDescOk returns a tuple with the TestDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCompliance) GetTestDescOk() (*string, bool) {
	if o == nil || IsNil(o.TestDesc) {
		return nil, false
	}
	return o.TestDesc, true
}

// HasTestDesc returns a boolean if a field has been set.
func (o *IngestersCompliance) HasTestDesc() bool {
	if o != nil && !IsNil(o.TestDesc) {
		return true
	}

	return false
}

// SetTestDesc gets a reference to the given string and assigns it to the TestDesc field.
func (o *IngestersCompliance) SetTestDesc(v string) {
	o.TestDesc = &v
}

// GetTestNumber returns the TestNumber field value if set, zero value otherwise.
func (o *IngestersCompliance) GetTestNumber() string {
	if o == nil || IsNil(o.TestNumber) {
		var ret string
		return ret
	}
	return *o.TestNumber
}

// GetTestNumberOk returns a tuple with the TestNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCompliance) GetTestNumberOk() (*string, bool) {
	if o == nil || IsNil(o.TestNumber) {
		return nil, false
	}
	return o.TestNumber, true
}

// HasTestNumber returns a boolean if a field has been set.
func (o *IngestersCompliance) HasTestNumber() bool {
	if o != nil && !IsNil(o.TestNumber) {
		return true
	}

	return false
}

// SetTestNumber gets a reference to the given string and assigns it to the TestNumber field.
func (o *IngestersCompliance) SetTestNumber(v string) {
	o.TestNumber = &v
}

// GetTestRationale returns the TestRationale field value if set, zero value otherwise.
func (o *IngestersCompliance) GetTestRationale() string {
	if o == nil || IsNil(o.TestRationale) {
		var ret string
		return ret
	}
	return *o.TestRationale
}

// GetTestRationaleOk returns a tuple with the TestRationale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCompliance) GetTestRationaleOk() (*string, bool) {
	if o == nil || IsNil(o.TestRationale) {
		return nil, false
	}
	return o.TestRationale, true
}

// HasTestRationale returns a boolean if a field has been set.
func (o *IngestersCompliance) HasTestRationale() bool {
	if o != nil && !IsNil(o.TestRationale) {
		return true
	}

	return false
}

// SetTestRationale gets a reference to the given string and assigns it to the TestRationale field.
func (o *IngestersCompliance) SetTestRationale(v string) {
	o.TestRationale = &v
}

// GetTestSeverity returns the TestSeverity field value if set, zero value otherwise.
func (o *IngestersCompliance) GetTestSeverity() string {
	if o == nil || IsNil(o.TestSeverity) {
		var ret string
		return ret
	}
	return *o.TestSeverity
}

// GetTestSeverityOk returns a tuple with the TestSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCompliance) GetTestSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.TestSeverity) {
		return nil, false
	}
	return o.TestSeverity, true
}

// HasTestSeverity returns a boolean if a field has been set.
func (o *IngestersCompliance) HasTestSeverity() bool {
	if o != nil && !IsNil(o.TestSeverity) {
		return true
	}

	return false
}

// SetTestSeverity gets a reference to the given string and assigns it to the TestSeverity field.
func (o *IngestersCompliance) SetTestSeverity(v string) {
	o.TestSeverity = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IngestersCompliance) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCompliance) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IngestersCompliance) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IngestersCompliance) SetType(v string) {
	o.Type = &v
}

func (o IngestersCompliance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestersCompliance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComplianceCheckType) {
		toSerialize["compliance_check_type"] = o.ComplianceCheckType
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.NodeId) {
		toSerialize["node_id"] = o.NodeId
	}
	if !IsNil(o.NodeType) {
		toSerialize["node_type"] = o.NodeType
	}
	if !IsNil(o.RemediationAnsible) {
		toSerialize["remediation_ansible"] = o.RemediationAnsible
	}
	if !IsNil(o.RemediationPuppet) {
		toSerialize["remediation_puppet"] = o.RemediationPuppet
	}
	if !IsNil(o.RemediationScript) {
		toSerialize["remediation_script"] = o.RemediationScript
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if !IsNil(o.ScanId) {
		toSerialize["scan_id"] = o.ScanId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TestCategory) {
		toSerialize["test_category"] = o.TestCategory
	}
	if !IsNil(o.TestDesc) {
		toSerialize["test_desc"] = o.TestDesc
	}
	if !IsNil(o.TestNumber) {
		toSerialize["test_number"] = o.TestNumber
	}
	if !IsNil(o.TestRationale) {
		toSerialize["test_rationale"] = o.TestRationale
	}
	if !IsNil(o.TestSeverity) {
		toSerialize["test_severity"] = o.TestSeverity
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableIngestersCompliance struct {
	value *IngestersCompliance
	isSet bool
}

func (v NullableIngestersCompliance) Get() *IngestersCompliance {
	return v.value
}

func (v *NullableIngestersCompliance) Set(val *IngestersCompliance) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestersCompliance) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestersCompliance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestersCompliance(val *IngestersCompliance) *NullableIngestersCompliance {
	return &NullableIngestersCompliance{value: val, isSet: true}
}

func (v NullableIngestersCompliance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestersCompliance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


