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
)

// checks if the IngestersSecretMatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestersSecretMatch{}

// IngestersSecretMatch struct for IngestersSecretMatch
type IngestersSecretMatch struct {
	FullFilename *string `json:"full_filename,omitempty"`
	MatchedContent *string `json:"matched_content,omitempty"`
	RelativeEndingIndex *int32 `json:"relative_ending_index,omitempty"`
	RelativeStartingIndex *int32 `json:"relative_starting_index,omitempty"`
	StartingIndex *int32 `json:"starting_index,omitempty"`
}

// NewIngestersSecretMatch instantiates a new IngestersSecretMatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestersSecretMatch() *IngestersSecretMatch {
	this := IngestersSecretMatch{}
	return &this
}

// NewIngestersSecretMatchWithDefaults instantiates a new IngestersSecretMatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestersSecretMatchWithDefaults() *IngestersSecretMatch {
	this := IngestersSecretMatch{}
	return &this
}

// GetFullFilename returns the FullFilename field value if set, zero value otherwise.
func (o *IngestersSecretMatch) GetFullFilename() string {
	if o == nil || IsNil(o.FullFilename) {
		var ret string
		return ret
	}
	return *o.FullFilename
}

// GetFullFilenameOk returns a tuple with the FullFilename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecretMatch) GetFullFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.FullFilename) {
		return nil, false
	}
	return o.FullFilename, true
}

// HasFullFilename returns a boolean if a field has been set.
func (o *IngestersSecretMatch) HasFullFilename() bool {
	if o != nil && !IsNil(o.FullFilename) {
		return true
	}

	return false
}

// SetFullFilename gets a reference to the given string and assigns it to the FullFilename field.
func (o *IngestersSecretMatch) SetFullFilename(v string) {
	o.FullFilename = &v
}

// GetMatchedContent returns the MatchedContent field value if set, zero value otherwise.
func (o *IngestersSecretMatch) GetMatchedContent() string {
	if o == nil || IsNil(o.MatchedContent) {
		var ret string
		return ret
	}
	return *o.MatchedContent
}

// GetMatchedContentOk returns a tuple with the MatchedContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecretMatch) GetMatchedContentOk() (*string, bool) {
	if o == nil || IsNil(o.MatchedContent) {
		return nil, false
	}
	return o.MatchedContent, true
}

// HasMatchedContent returns a boolean if a field has been set.
func (o *IngestersSecretMatch) HasMatchedContent() bool {
	if o != nil && !IsNil(o.MatchedContent) {
		return true
	}

	return false
}

// SetMatchedContent gets a reference to the given string and assigns it to the MatchedContent field.
func (o *IngestersSecretMatch) SetMatchedContent(v string) {
	o.MatchedContent = &v
}

// GetRelativeEndingIndex returns the RelativeEndingIndex field value if set, zero value otherwise.
func (o *IngestersSecretMatch) GetRelativeEndingIndex() int32 {
	if o == nil || IsNil(o.RelativeEndingIndex) {
		var ret int32
		return ret
	}
	return *o.RelativeEndingIndex
}

// GetRelativeEndingIndexOk returns a tuple with the RelativeEndingIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecretMatch) GetRelativeEndingIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.RelativeEndingIndex) {
		return nil, false
	}
	return o.RelativeEndingIndex, true
}

// HasRelativeEndingIndex returns a boolean if a field has been set.
func (o *IngestersSecretMatch) HasRelativeEndingIndex() bool {
	if o != nil && !IsNil(o.RelativeEndingIndex) {
		return true
	}

	return false
}

// SetRelativeEndingIndex gets a reference to the given int32 and assigns it to the RelativeEndingIndex field.
func (o *IngestersSecretMatch) SetRelativeEndingIndex(v int32) {
	o.RelativeEndingIndex = &v
}

// GetRelativeStartingIndex returns the RelativeStartingIndex field value if set, zero value otherwise.
func (o *IngestersSecretMatch) GetRelativeStartingIndex() int32 {
	if o == nil || IsNil(o.RelativeStartingIndex) {
		var ret int32
		return ret
	}
	return *o.RelativeStartingIndex
}

// GetRelativeStartingIndexOk returns a tuple with the RelativeStartingIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecretMatch) GetRelativeStartingIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.RelativeStartingIndex) {
		return nil, false
	}
	return o.RelativeStartingIndex, true
}

// HasRelativeStartingIndex returns a boolean if a field has been set.
func (o *IngestersSecretMatch) HasRelativeStartingIndex() bool {
	if o != nil && !IsNil(o.RelativeStartingIndex) {
		return true
	}

	return false
}

// SetRelativeStartingIndex gets a reference to the given int32 and assigns it to the RelativeStartingIndex field.
func (o *IngestersSecretMatch) SetRelativeStartingIndex(v int32) {
	o.RelativeStartingIndex = &v
}

// GetStartingIndex returns the StartingIndex field value if set, zero value otherwise.
func (o *IngestersSecretMatch) GetStartingIndex() int32 {
	if o == nil || IsNil(o.StartingIndex) {
		var ret int32
		return ret
	}
	return *o.StartingIndex
}

// GetStartingIndexOk returns a tuple with the StartingIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecretMatch) GetStartingIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.StartingIndex) {
		return nil, false
	}
	return o.StartingIndex, true
}

// HasStartingIndex returns a boolean if a field has been set.
func (o *IngestersSecretMatch) HasStartingIndex() bool {
	if o != nil && !IsNil(o.StartingIndex) {
		return true
	}

	return false
}

// SetStartingIndex gets a reference to the given int32 and assigns it to the StartingIndex field.
func (o *IngestersSecretMatch) SetStartingIndex(v int32) {
	o.StartingIndex = &v
}

func (o IngestersSecretMatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestersSecretMatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FullFilename) {
		toSerialize["full_filename"] = o.FullFilename
	}
	if !IsNil(o.MatchedContent) {
		toSerialize["matched_content"] = o.MatchedContent
	}
	if !IsNil(o.RelativeEndingIndex) {
		toSerialize["relative_ending_index"] = o.RelativeEndingIndex
	}
	if !IsNil(o.RelativeStartingIndex) {
		toSerialize["relative_starting_index"] = o.RelativeStartingIndex
	}
	if !IsNil(o.StartingIndex) {
		toSerialize["starting_index"] = o.StartingIndex
	}
	return toSerialize, nil
}

type NullableIngestersSecretMatch struct {
	value *IngestersSecretMatch
	isSet bool
}

func (v NullableIngestersSecretMatch) Get() *IngestersSecretMatch {
	return v.value
}

func (v *NullableIngestersSecretMatch) Set(val *IngestersSecretMatch) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestersSecretMatch) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestersSecretMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestersSecretMatch(val *IngestersSecretMatch) *NullableIngestersSecretMatch {
	return &NullableIngestersSecretMatch{value: val, isSet: true}
}

func (v NullableIngestersSecretMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestersSecretMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


