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

// checks if the ModelCompliance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCompliance{}

// ModelCompliance struct for ModelCompliance
type ModelCompliance struct {
	ComplianceCheckType string `json:"compliance_check_type"`
	Description string `json:"description"`
	Masked bool `json:"masked"`
	NodeId string `json:"node_id"`
	NodeType string `json:"node_type"`
	RemediationAnsible string `json:"remediation_ansible"`
	RemediationPuppet string `json:"remediation_puppet"`
	RemediationScript string `json:"remediation_script"`
	Resource string `json:"resource"`
	Resources []ModelBasicNode `json:"resources,omitempty"`
	RuleId string `json:"rule_id"`
	Status string `json:"status"`
	TestCategory string `json:"test_category"`
	TestDesc string `json:"test_desc"`
	TestNumber string `json:"test_number"`
	TestRationale string `json:"test_rationale"`
	TestSeverity string `json:"test_severity"`
	UpdatedAt int32 `json:"updated_at"`
}

// NewModelCompliance instantiates a new ModelCompliance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCompliance(complianceCheckType string, description string, masked bool, nodeId string, nodeType string, remediationAnsible string, remediationPuppet string, remediationScript string, resource string, ruleId string, status string, testCategory string, testDesc string, testNumber string, testRationale string, testSeverity string, updatedAt int32) *ModelCompliance {
	this := ModelCompliance{}
	this.ComplianceCheckType = complianceCheckType
	this.Description = description
	this.Masked = masked
	this.NodeId = nodeId
	this.NodeType = nodeType
	this.RemediationAnsible = remediationAnsible
	this.RemediationPuppet = remediationPuppet
	this.RemediationScript = remediationScript
	this.Resource = resource
	this.RuleId = ruleId
	this.Status = status
	this.TestCategory = testCategory
	this.TestDesc = testDesc
	this.TestNumber = testNumber
	this.TestRationale = testRationale
	this.TestSeverity = testSeverity
	this.UpdatedAt = updatedAt
	return &this
}

// NewModelComplianceWithDefaults instantiates a new ModelCompliance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelComplianceWithDefaults() *ModelCompliance {
	this := ModelCompliance{}
	return &this
}

// GetComplianceCheckType returns the ComplianceCheckType field value
func (o *ModelCompliance) GetComplianceCheckType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComplianceCheckType
}

// GetComplianceCheckTypeOk returns a tuple with the ComplianceCheckType field value
// and a boolean to check if the value has been set.
func (o *ModelCompliance) GetComplianceCheckTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComplianceCheckType, true
}

// SetComplianceCheckType sets field value
func (o *ModelCompliance) SetComplianceCheckType(v string) {
	o.ComplianceCheckType = v
}

// GetDescription returns the Description field value
func (o *ModelCompliance) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ModelCompliance) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ModelCompliance) SetDescription(v string) {
	o.Description = v
}

// GetMasked returns the Masked field value
func (o *ModelCompliance) GetMasked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Masked
}

// GetMaskedOk returns a tuple with the Masked field value
// and a boolean to check if the value has been set.
func (o *ModelCompliance) GetMaskedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Masked, true
}

// SetMasked sets field value
func (o *ModelCompliance) SetMasked(v bool) {
	o.Masked = v
}

// GetNodeId returns the NodeId field value
func (o *ModelCompliance) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelCompliance) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelCompliance) SetNodeId(v string) {
	o.NodeId = v
}

// GetNodeType returns the NodeType field value
func (o *ModelCompliance) GetNodeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *ModelCompliance) GetNodeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value
func (o *ModelCompliance) SetNodeType(v string) {
	o.NodeType = v
}

// GetRemediationAnsible returns the RemediationAnsible field value
func (o *ModelCompliance) GetRemediationAnsible() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemediationAnsible
}

// GetRemediationAnsibleOk returns a tuple with the RemediationAnsible field value
// and a boolean to check if the value has been set.
func (o *ModelCompliance) GetRemediationAnsibleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemediationAnsible, true
}

// SetRemediationAnsible sets field value
func (o *ModelCompliance) SetRemediationAnsible(v string) {
	o.RemediationAnsible = v
}

// GetRemediationPuppet returns the RemediationPuppet field value
func (o *ModelCompliance) GetRemediationPuppet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemediationPuppet
}

// GetRemediationPuppetOk returns a tuple with the RemediationPuppet field value
// and a boolean to check if the value has been set.
func (o *ModelCompliance) GetRemediationPuppetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemediationPuppet, true
}

// SetRemediationPuppet sets field value
func (o *ModelCompliance) SetRemediationPuppet(v string) {
	o.RemediationPuppet = v
}

// GetRemediationScript returns the RemediationScript field value
func (o *ModelCompliance) GetRemediationScript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemediationScript
}

// GetRemediationScriptOk returns a tuple with the RemediationScript field value
// and a boolean to check if the value has been set.
func (o *ModelCompliance) GetRemediationScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemediationScript, true
}

// SetRemediationScript sets field value
func (o *ModelCompliance) SetRemediationScript(v string) {
	o.RemediationScript = v
}

// GetResource returns the Resource field value
func (o *ModelCompliance) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *ModelCompliance) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *ModelCompliance) SetResource(v string) {
	o.Resource = v
}

// GetResources returns the Resources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCompliance) GetResources() []ModelBasicNode {
	if o == nil {
		var ret []ModelBasicNode
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCompliance) GetResourcesOk() ([]ModelBasicNode, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *ModelCompliance) HasResources() bool {
	if o != nil && IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []ModelBasicNode and assigns it to the Resources field.
func (o *ModelCompliance) SetResources(v []ModelBasicNode) {
	o.Resources = v
}

// GetRuleId returns the RuleId field value
func (o *ModelCompliance) GetRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *ModelCompliance) GetRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value
func (o *ModelCompliance) SetRuleId(v string) {
	o.RuleId = v
}

// GetStatus returns the Status field value
func (o *ModelCompliance) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ModelCompliance) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ModelCompliance) SetStatus(v string) {
	o.Status = v
}

// GetTestCategory returns the TestCategory field value
func (o *ModelCompliance) GetTestCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestCategory
}

// GetTestCategoryOk returns a tuple with the TestCategory field value
// and a boolean to check if the value has been set.
func (o *ModelCompliance) GetTestCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestCategory, true
}

// SetTestCategory sets field value
func (o *ModelCompliance) SetTestCategory(v string) {
	o.TestCategory = v
}

// GetTestDesc returns the TestDesc field value
func (o *ModelCompliance) GetTestDesc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestDesc
}

// GetTestDescOk returns a tuple with the TestDesc field value
// and a boolean to check if the value has been set.
func (o *ModelCompliance) GetTestDescOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestDesc, true
}

// SetTestDesc sets field value
func (o *ModelCompliance) SetTestDesc(v string) {
	o.TestDesc = v
}

// GetTestNumber returns the TestNumber field value
func (o *ModelCompliance) GetTestNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestNumber
}

// GetTestNumberOk returns a tuple with the TestNumber field value
// and a boolean to check if the value has been set.
func (o *ModelCompliance) GetTestNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestNumber, true
}

// SetTestNumber sets field value
func (o *ModelCompliance) SetTestNumber(v string) {
	o.TestNumber = v
}

// GetTestRationale returns the TestRationale field value
func (o *ModelCompliance) GetTestRationale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestRationale
}

// GetTestRationaleOk returns a tuple with the TestRationale field value
// and a boolean to check if the value has been set.
func (o *ModelCompliance) GetTestRationaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestRationale, true
}

// SetTestRationale sets field value
func (o *ModelCompliance) SetTestRationale(v string) {
	o.TestRationale = v
}

// GetTestSeverity returns the TestSeverity field value
func (o *ModelCompliance) GetTestSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestSeverity
}

// GetTestSeverityOk returns a tuple with the TestSeverity field value
// and a boolean to check if the value has been set.
func (o *ModelCompliance) GetTestSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestSeverity, true
}

// SetTestSeverity sets field value
func (o *ModelCompliance) SetTestSeverity(v string) {
	o.TestSeverity = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ModelCompliance) GetUpdatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelCompliance) GetUpdatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ModelCompliance) SetUpdatedAt(v int32) {
	o.UpdatedAt = v
}

func (o ModelCompliance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCompliance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["compliance_check_type"] = o.ComplianceCheckType
	toSerialize["description"] = o.Description
	toSerialize["masked"] = o.Masked
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_type"] = o.NodeType
	toSerialize["remediation_ansible"] = o.RemediationAnsible
	toSerialize["remediation_puppet"] = o.RemediationPuppet
	toSerialize["remediation_script"] = o.RemediationScript
	toSerialize["resource"] = o.Resource
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	toSerialize["rule_id"] = o.RuleId
	toSerialize["status"] = o.Status
	toSerialize["test_category"] = o.TestCategory
	toSerialize["test_desc"] = o.TestDesc
	toSerialize["test_number"] = o.TestNumber
	toSerialize["test_rationale"] = o.TestRationale
	toSerialize["test_severity"] = o.TestSeverity
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableModelCompliance struct {
	value *ModelCompliance
	isSet bool
}

func (v NullableModelCompliance) Get() *ModelCompliance {
	return v.value
}

func (v *NullableModelCompliance) Set(val *ModelCompliance) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCompliance) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCompliance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCompliance(val *ModelCompliance) *NullableModelCompliance {
	return &NullableModelCompliance{value: val, isSet: true}
}

func (v NullableModelCompliance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCompliance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


