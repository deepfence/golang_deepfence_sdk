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

// checks if the ModelRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelRule{}

// ModelRule struct for ModelRule
type ModelRule struct {
	Id int32 `json:"id"`
	Level string `json:"level"`
	Name string `json:"name"`
	Part string `json:"part"`
	Score float32 `json:"score"`
	SignatureToMatch string `json:"signature_to_match"`
}

// NewModelRule instantiates a new ModelRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelRule(id int32, level string, name string, part string, score float32, signatureToMatch string) *ModelRule {
	this := ModelRule{}
	this.Id = id
	this.Level = level
	this.Name = name
	this.Part = part
	this.Score = score
	this.SignatureToMatch = signatureToMatch
	return &this
}

// NewModelRuleWithDefaults instantiates a new ModelRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelRuleWithDefaults() *ModelRule {
	this := ModelRule{}
	return &this
}

// GetId returns the Id field value
func (o *ModelRule) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModelRule) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModelRule) SetId(v int32) {
	o.Id = v
}

// GetLevel returns the Level field value
func (o *ModelRule) GetLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *ModelRule) GetLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *ModelRule) SetLevel(v string) {
	o.Level = v
}

// GetName returns the Name field value
func (o *ModelRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModelRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ModelRule) SetName(v string) {
	o.Name = v
}

// GetPart returns the Part field value
func (o *ModelRule) GetPart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Part
}

// GetPartOk returns a tuple with the Part field value
// and a boolean to check if the value has been set.
func (o *ModelRule) GetPartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Part, true
}

// SetPart sets field value
func (o *ModelRule) SetPart(v string) {
	o.Part = v
}

// GetScore returns the Score field value
func (o *ModelRule) GetScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *ModelRule) GetScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *ModelRule) SetScore(v float32) {
	o.Score = v
}

// GetSignatureToMatch returns the SignatureToMatch field value
func (o *ModelRule) GetSignatureToMatch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignatureToMatch
}

// GetSignatureToMatchOk returns a tuple with the SignatureToMatch field value
// and a boolean to check if the value has been set.
func (o *ModelRule) GetSignatureToMatchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignatureToMatch, true
}

// SetSignatureToMatch sets field value
func (o *ModelRule) SetSignatureToMatch(v string) {
	o.SignatureToMatch = v
}

func (o ModelRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["level"] = o.Level
	toSerialize["name"] = o.Name
	toSerialize["part"] = o.Part
	toSerialize["score"] = o.Score
	toSerialize["signature_to_match"] = o.SignatureToMatch
	return toSerialize, nil
}

type NullableModelRule struct {
	value *ModelRule
	isSet bool
}

func (v NullableModelRule) Get() *ModelRule {
	return v.value
}

func (v *NullableModelRule) Set(val *ModelRule) {
	v.value = val
	v.isSet = true
}

func (v NullableModelRule) IsSet() bool {
	return v.isSet
}

func (v *NullableModelRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelRule(val *ModelRule) *NullableModelRule {
	return &NullableModelRule{value: val, isSet: true}
}

func (v NullableModelRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


