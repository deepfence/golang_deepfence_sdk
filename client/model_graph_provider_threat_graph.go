/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.5.5
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GraphProviderThreatGraph type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphProviderThreatGraph{}

// GraphProviderThreatGraph struct for GraphProviderThreatGraph
type GraphProviderThreatGraph struct {
	CloudComplianceCount int32 `json:"cloud_compliance_count"`
	CloudWarnAlarmCount int32 `json:"cloud_warn_alarm_count"`
	ComplianceCount int32 `json:"compliance_count"`
	ExploitableSecretsCount int32 `json:"exploitable_secrets_count"`
	ExploitableVulnerabilitiesCount int32 `json:"exploitable_vulnerabilities_count"`
	Resources []GraphThreatNodeInfo `json:"resources"`
	SecretsCount int32 `json:"secrets_count"`
	VulnerabilityCount int32 `json:"vulnerability_count"`
	WarnAlarmCount int32 `json:"warn_alarm_count"`
}

type _GraphProviderThreatGraph GraphProviderThreatGraph

// NewGraphProviderThreatGraph instantiates a new GraphProviderThreatGraph object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphProviderThreatGraph(cloudComplianceCount int32, cloudWarnAlarmCount int32, complianceCount int32, exploitableSecretsCount int32, exploitableVulnerabilitiesCount int32, resources []GraphThreatNodeInfo, secretsCount int32, vulnerabilityCount int32, warnAlarmCount int32) *GraphProviderThreatGraph {
	this := GraphProviderThreatGraph{}
	this.CloudComplianceCount = cloudComplianceCount
	this.CloudWarnAlarmCount = cloudWarnAlarmCount
	this.ComplianceCount = complianceCount
	this.ExploitableSecretsCount = exploitableSecretsCount
	this.ExploitableVulnerabilitiesCount = exploitableVulnerabilitiesCount
	this.Resources = resources
	this.SecretsCount = secretsCount
	this.VulnerabilityCount = vulnerabilityCount
	this.WarnAlarmCount = warnAlarmCount
	return &this
}

// NewGraphProviderThreatGraphWithDefaults instantiates a new GraphProviderThreatGraph object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphProviderThreatGraphWithDefaults() *GraphProviderThreatGraph {
	this := GraphProviderThreatGraph{}
	return &this
}

// GetCloudComplianceCount returns the CloudComplianceCount field value
func (o *GraphProviderThreatGraph) GetCloudComplianceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CloudComplianceCount
}

// GetCloudComplianceCountOk returns a tuple with the CloudComplianceCount field value
// and a boolean to check if the value has been set.
func (o *GraphProviderThreatGraph) GetCloudComplianceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudComplianceCount, true
}

// SetCloudComplianceCount sets field value
func (o *GraphProviderThreatGraph) SetCloudComplianceCount(v int32) {
	o.CloudComplianceCount = v
}

// GetCloudWarnAlarmCount returns the CloudWarnAlarmCount field value
func (o *GraphProviderThreatGraph) GetCloudWarnAlarmCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CloudWarnAlarmCount
}

// GetCloudWarnAlarmCountOk returns a tuple with the CloudWarnAlarmCount field value
// and a boolean to check if the value has been set.
func (o *GraphProviderThreatGraph) GetCloudWarnAlarmCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudWarnAlarmCount, true
}

// SetCloudWarnAlarmCount sets field value
func (o *GraphProviderThreatGraph) SetCloudWarnAlarmCount(v int32) {
	o.CloudWarnAlarmCount = v
}

// GetComplianceCount returns the ComplianceCount field value
func (o *GraphProviderThreatGraph) GetComplianceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ComplianceCount
}

// GetComplianceCountOk returns a tuple with the ComplianceCount field value
// and a boolean to check if the value has been set.
func (o *GraphProviderThreatGraph) GetComplianceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComplianceCount, true
}

// SetComplianceCount sets field value
func (o *GraphProviderThreatGraph) SetComplianceCount(v int32) {
	o.ComplianceCount = v
}

// GetExploitableSecretsCount returns the ExploitableSecretsCount field value
func (o *GraphProviderThreatGraph) GetExploitableSecretsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExploitableSecretsCount
}

// GetExploitableSecretsCountOk returns a tuple with the ExploitableSecretsCount field value
// and a boolean to check if the value has been set.
func (o *GraphProviderThreatGraph) GetExploitableSecretsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExploitableSecretsCount, true
}

// SetExploitableSecretsCount sets field value
func (o *GraphProviderThreatGraph) SetExploitableSecretsCount(v int32) {
	o.ExploitableSecretsCount = v
}

// GetExploitableVulnerabilitiesCount returns the ExploitableVulnerabilitiesCount field value
func (o *GraphProviderThreatGraph) GetExploitableVulnerabilitiesCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExploitableVulnerabilitiesCount
}

// GetExploitableVulnerabilitiesCountOk returns a tuple with the ExploitableVulnerabilitiesCount field value
// and a boolean to check if the value has been set.
func (o *GraphProviderThreatGraph) GetExploitableVulnerabilitiesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExploitableVulnerabilitiesCount, true
}

// SetExploitableVulnerabilitiesCount sets field value
func (o *GraphProviderThreatGraph) SetExploitableVulnerabilitiesCount(v int32) {
	o.ExploitableVulnerabilitiesCount = v
}

// GetResources returns the Resources field value
// If the value is explicit nil, the zero value for []GraphThreatNodeInfo will be returned
func (o *GraphProviderThreatGraph) GetResources() []GraphThreatNodeInfo {
	if o == nil {
		var ret []GraphThreatNodeInfo
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphProviderThreatGraph) GetResourcesOk() ([]GraphThreatNodeInfo, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// SetResources sets field value
func (o *GraphProviderThreatGraph) SetResources(v []GraphThreatNodeInfo) {
	o.Resources = v
}

// GetSecretsCount returns the SecretsCount field value
func (o *GraphProviderThreatGraph) GetSecretsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SecretsCount
}

// GetSecretsCountOk returns a tuple with the SecretsCount field value
// and a boolean to check if the value has been set.
func (o *GraphProviderThreatGraph) GetSecretsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretsCount, true
}

// SetSecretsCount sets field value
func (o *GraphProviderThreatGraph) SetSecretsCount(v int32) {
	o.SecretsCount = v
}

// GetVulnerabilityCount returns the VulnerabilityCount field value
func (o *GraphProviderThreatGraph) GetVulnerabilityCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VulnerabilityCount
}

// GetVulnerabilityCountOk returns a tuple with the VulnerabilityCount field value
// and a boolean to check if the value has been set.
func (o *GraphProviderThreatGraph) GetVulnerabilityCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VulnerabilityCount, true
}

// SetVulnerabilityCount sets field value
func (o *GraphProviderThreatGraph) SetVulnerabilityCount(v int32) {
	o.VulnerabilityCount = v
}

// GetWarnAlarmCount returns the WarnAlarmCount field value
func (o *GraphProviderThreatGraph) GetWarnAlarmCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WarnAlarmCount
}

// GetWarnAlarmCountOk returns a tuple with the WarnAlarmCount field value
// and a boolean to check if the value has been set.
func (o *GraphProviderThreatGraph) GetWarnAlarmCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WarnAlarmCount, true
}

// SetWarnAlarmCount sets field value
func (o *GraphProviderThreatGraph) SetWarnAlarmCount(v int32) {
	o.WarnAlarmCount = v
}

func (o GraphProviderThreatGraph) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphProviderThreatGraph) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloud_compliance_count"] = o.CloudComplianceCount
	toSerialize["cloud_warn_alarm_count"] = o.CloudWarnAlarmCount
	toSerialize["compliance_count"] = o.ComplianceCount
	toSerialize["exploitable_secrets_count"] = o.ExploitableSecretsCount
	toSerialize["exploitable_vulnerabilities_count"] = o.ExploitableVulnerabilitiesCount
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	toSerialize["secrets_count"] = o.SecretsCount
	toSerialize["vulnerability_count"] = o.VulnerabilityCount
	toSerialize["warn_alarm_count"] = o.WarnAlarmCount
	return toSerialize, nil
}

func (o *GraphProviderThreatGraph) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloud_compliance_count",
		"cloud_warn_alarm_count",
		"compliance_count",
		"exploitable_secrets_count",
		"exploitable_vulnerabilities_count",
		"resources",
		"secrets_count",
		"vulnerability_count",
		"warn_alarm_count",
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

	varGraphProviderThreatGraph := _GraphProviderThreatGraph{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGraphProviderThreatGraph)

	if err != nil {
		return err
	}

	*o = GraphProviderThreatGraph(varGraphProviderThreatGraph)

	return err
}

type NullableGraphProviderThreatGraph struct {
	value *GraphProviderThreatGraph
	isSet bool
}

func (v NullableGraphProviderThreatGraph) Get() *GraphProviderThreatGraph {
	return v.value
}

func (v *NullableGraphProviderThreatGraph) Set(val *GraphProviderThreatGraph) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphProviderThreatGraph) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphProviderThreatGraph) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphProviderThreatGraph(val *GraphProviderThreatGraph) *NullableGraphProviderThreatGraph {
	return &NullableGraphProviderThreatGraph{value: val, isSet: true}
}

func (v NullableGraphProviderThreatGraph) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphProviderThreatGraph) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


