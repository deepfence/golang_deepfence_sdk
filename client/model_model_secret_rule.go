/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.5.3
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ModelSecretRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSecretRule{}

// ModelSecretRule struct for ModelSecretRule
type ModelSecretRule struct {
	Level string `json:"level"`
	Masked bool `json:"masked"`
	Part *string `json:"part,omitempty"`
	Payload string `json:"payload"`
	RuleId *string `json:"rule_id,omitempty"`
	Severity string `json:"severity"`
	SignatureToMatch *string `json:"signature_to_match,omitempty"`
	Summary string `json:"summary"`
	UpdatedAt int32 `json:"updated_at"`
}

type _ModelSecretRule ModelSecretRule

// NewModelSecretRule instantiates a new ModelSecretRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSecretRule(level string, masked bool, payload string, severity string, summary string, updatedAt int32) *ModelSecretRule {
	this := ModelSecretRule{}
	this.Level = level
	this.Masked = masked
	this.Payload = payload
	this.Severity = severity
	this.Summary = summary
	this.UpdatedAt = updatedAt
	return &this
}

// NewModelSecretRuleWithDefaults instantiates a new ModelSecretRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSecretRuleWithDefaults() *ModelSecretRule {
	this := ModelSecretRule{}
	return &this
}

// GetLevel returns the Level field value
func (o *ModelSecretRule) GetLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *ModelSecretRule) GetLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *ModelSecretRule) SetLevel(v string) {
	o.Level = v
}

// GetMasked returns the Masked field value
func (o *ModelSecretRule) GetMasked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Masked
}

// GetMaskedOk returns a tuple with the Masked field value
// and a boolean to check if the value has been set.
func (o *ModelSecretRule) GetMaskedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Masked, true
}

// SetMasked sets field value
func (o *ModelSecretRule) SetMasked(v bool) {
	o.Masked = v
}

// GetPart returns the Part field value if set, zero value otherwise.
func (o *ModelSecretRule) GetPart() string {
	if o == nil || IsNil(o.Part) {
		var ret string
		return ret
	}
	return *o.Part
}

// GetPartOk returns a tuple with the Part field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSecretRule) GetPartOk() (*string, bool) {
	if o == nil || IsNil(o.Part) {
		return nil, false
	}
	return o.Part, true
}

// HasPart returns a boolean if a field has been set.
func (o *ModelSecretRule) HasPart() bool {
	if o != nil && !IsNil(o.Part) {
		return true
	}

	return false
}

// SetPart gets a reference to the given string and assigns it to the Part field.
func (o *ModelSecretRule) SetPart(v string) {
	o.Part = &v
}

// GetPayload returns the Payload field value
func (o *ModelSecretRule) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *ModelSecretRule) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *ModelSecretRule) SetPayload(v string) {
	o.Payload = v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *ModelSecretRule) GetRuleId() string {
	if o == nil || IsNil(o.RuleId) {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSecretRule) GetRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RuleId) {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *ModelSecretRule) HasRuleId() bool {
	if o != nil && !IsNil(o.RuleId) {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *ModelSecretRule) SetRuleId(v string) {
	o.RuleId = &v
}

// GetSeverity returns the Severity field value
func (o *ModelSecretRule) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *ModelSecretRule) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *ModelSecretRule) SetSeverity(v string) {
	o.Severity = v
}

// GetSignatureToMatch returns the SignatureToMatch field value if set, zero value otherwise.
func (o *ModelSecretRule) GetSignatureToMatch() string {
	if o == nil || IsNil(o.SignatureToMatch) {
		var ret string
		return ret
	}
	return *o.SignatureToMatch
}

// GetSignatureToMatchOk returns a tuple with the SignatureToMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSecretRule) GetSignatureToMatchOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureToMatch) {
		return nil, false
	}
	return o.SignatureToMatch, true
}

// HasSignatureToMatch returns a boolean if a field has been set.
func (o *ModelSecretRule) HasSignatureToMatch() bool {
	if o != nil && !IsNil(o.SignatureToMatch) {
		return true
	}

	return false
}

// SetSignatureToMatch gets a reference to the given string and assigns it to the SignatureToMatch field.
func (o *ModelSecretRule) SetSignatureToMatch(v string) {
	o.SignatureToMatch = &v
}

// GetSummary returns the Summary field value
func (o *ModelSecretRule) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *ModelSecretRule) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *ModelSecretRule) SetSummary(v string) {
	o.Summary = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ModelSecretRule) GetUpdatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelSecretRule) GetUpdatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ModelSecretRule) SetUpdatedAt(v int32) {
	o.UpdatedAt = v
}

func (o ModelSecretRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSecretRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["level"] = o.Level
	toSerialize["masked"] = o.Masked
	if !IsNil(o.Part) {
		toSerialize["part"] = o.Part
	}
	toSerialize["payload"] = o.Payload
	if !IsNil(o.RuleId) {
		toSerialize["rule_id"] = o.RuleId
	}
	toSerialize["severity"] = o.Severity
	if !IsNil(o.SignatureToMatch) {
		toSerialize["signature_to_match"] = o.SignatureToMatch
	}
	toSerialize["summary"] = o.Summary
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *ModelSecretRule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"level",
		"masked",
		"payload",
		"severity",
		"summary",
		"updated_at",
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

	varModelSecretRule := _ModelSecretRule{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelSecretRule)

	if err != nil {
		return err
	}

	*o = ModelSecretRule(varModelSecretRule)

	return err
}

type NullableModelSecretRule struct {
	value *ModelSecretRule
	isSet bool
}

func (v NullableModelSecretRule) Get() *ModelSecretRule {
	return v.value
}

func (v *NullableModelSecretRule) Set(val *ModelSecretRule) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSecretRule) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSecretRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSecretRule(val *ModelSecretRule) *NullableModelSecretRule {
	return &NullableModelSecretRule{value: val, isSet: true}
}

func (v NullableModelSecretRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSecretRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


