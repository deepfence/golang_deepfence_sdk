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

// checks if the ModelEmailConfigurationAdd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelEmailConfigurationAdd{}

// ModelEmailConfigurationAdd struct for ModelEmailConfigurationAdd
type ModelEmailConfigurationAdd struct {
	AmazonAccessKey *string `json:"amazon_access_key,omitempty"`
	AmazonSecretKey *string `json:"amazon_secret_key,omitempty"`
	Apikey *string `json:"apikey,omitempty"`
	CreatedByUserId *int32 `json:"created_by_user_id,omitempty"`
	EmailId *string `json:"email_id,omitempty"`
	EmailProvider *string `json:"email_provider,omitempty"`
	Password *string `json:"password,omitempty"`
	Port *string `json:"port,omitempty"`
	SesRegion *string `json:"ses_region,omitempty"`
	Smtp *string `json:"smtp,omitempty"`
}

// NewModelEmailConfigurationAdd instantiates a new ModelEmailConfigurationAdd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelEmailConfigurationAdd() *ModelEmailConfigurationAdd {
	this := ModelEmailConfigurationAdd{}
	return &this
}

// NewModelEmailConfigurationAddWithDefaults instantiates a new ModelEmailConfigurationAdd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelEmailConfigurationAddWithDefaults() *ModelEmailConfigurationAdd {
	this := ModelEmailConfigurationAdd{}
	return &this
}

// GetAmazonAccessKey returns the AmazonAccessKey field value if set, zero value otherwise.
func (o *ModelEmailConfigurationAdd) GetAmazonAccessKey() string {
	if o == nil || IsNil(o.AmazonAccessKey) {
		var ret string
		return ret
	}
	return *o.AmazonAccessKey
}

// GetAmazonAccessKeyOk returns a tuple with the AmazonAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEmailConfigurationAdd) GetAmazonAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AmazonAccessKey) {
		return nil, false
	}
	return o.AmazonAccessKey, true
}

// HasAmazonAccessKey returns a boolean if a field has been set.
func (o *ModelEmailConfigurationAdd) HasAmazonAccessKey() bool {
	if o != nil && !IsNil(o.AmazonAccessKey) {
		return true
	}

	return false
}

// SetAmazonAccessKey gets a reference to the given string and assigns it to the AmazonAccessKey field.
func (o *ModelEmailConfigurationAdd) SetAmazonAccessKey(v string) {
	o.AmazonAccessKey = &v
}

// GetAmazonSecretKey returns the AmazonSecretKey field value if set, zero value otherwise.
func (o *ModelEmailConfigurationAdd) GetAmazonSecretKey() string {
	if o == nil || IsNil(o.AmazonSecretKey) {
		var ret string
		return ret
	}
	return *o.AmazonSecretKey
}

// GetAmazonSecretKeyOk returns a tuple with the AmazonSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEmailConfigurationAdd) GetAmazonSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AmazonSecretKey) {
		return nil, false
	}
	return o.AmazonSecretKey, true
}

// HasAmazonSecretKey returns a boolean if a field has been set.
func (o *ModelEmailConfigurationAdd) HasAmazonSecretKey() bool {
	if o != nil && !IsNil(o.AmazonSecretKey) {
		return true
	}

	return false
}

// SetAmazonSecretKey gets a reference to the given string and assigns it to the AmazonSecretKey field.
func (o *ModelEmailConfigurationAdd) SetAmazonSecretKey(v string) {
	o.AmazonSecretKey = &v
}

// GetApikey returns the Apikey field value if set, zero value otherwise.
func (o *ModelEmailConfigurationAdd) GetApikey() string {
	if o == nil || IsNil(o.Apikey) {
		var ret string
		return ret
	}
	return *o.Apikey
}

// GetApikeyOk returns a tuple with the Apikey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEmailConfigurationAdd) GetApikeyOk() (*string, bool) {
	if o == nil || IsNil(o.Apikey) {
		return nil, false
	}
	return o.Apikey, true
}

// HasApikey returns a boolean if a field has been set.
func (o *ModelEmailConfigurationAdd) HasApikey() bool {
	if o != nil && !IsNil(o.Apikey) {
		return true
	}

	return false
}

// SetApikey gets a reference to the given string and assigns it to the Apikey field.
func (o *ModelEmailConfigurationAdd) SetApikey(v string) {
	o.Apikey = &v
}

// GetCreatedByUserId returns the CreatedByUserId field value if set, zero value otherwise.
func (o *ModelEmailConfigurationAdd) GetCreatedByUserId() int32 {
	if o == nil || IsNil(o.CreatedByUserId) {
		var ret int32
		return ret
	}
	return *o.CreatedByUserId
}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEmailConfigurationAdd) GetCreatedByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedByUserId) {
		return nil, false
	}
	return o.CreatedByUserId, true
}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *ModelEmailConfigurationAdd) HasCreatedByUserId() bool {
	if o != nil && !IsNil(o.CreatedByUserId) {
		return true
	}

	return false
}

// SetCreatedByUserId gets a reference to the given int32 and assigns it to the CreatedByUserId field.
func (o *ModelEmailConfigurationAdd) SetCreatedByUserId(v int32) {
	o.CreatedByUserId = &v
}

// GetEmailId returns the EmailId field value if set, zero value otherwise.
func (o *ModelEmailConfigurationAdd) GetEmailId() string {
	if o == nil || IsNil(o.EmailId) {
		var ret string
		return ret
	}
	return *o.EmailId
}

// GetEmailIdOk returns a tuple with the EmailId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEmailConfigurationAdd) GetEmailIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmailId) {
		return nil, false
	}
	return o.EmailId, true
}

// HasEmailId returns a boolean if a field has been set.
func (o *ModelEmailConfigurationAdd) HasEmailId() bool {
	if o != nil && !IsNil(o.EmailId) {
		return true
	}

	return false
}

// SetEmailId gets a reference to the given string and assigns it to the EmailId field.
func (o *ModelEmailConfigurationAdd) SetEmailId(v string) {
	o.EmailId = &v
}

// GetEmailProvider returns the EmailProvider field value if set, zero value otherwise.
func (o *ModelEmailConfigurationAdd) GetEmailProvider() string {
	if o == nil || IsNil(o.EmailProvider) {
		var ret string
		return ret
	}
	return *o.EmailProvider
}

// GetEmailProviderOk returns a tuple with the EmailProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEmailConfigurationAdd) GetEmailProviderOk() (*string, bool) {
	if o == nil || IsNil(o.EmailProvider) {
		return nil, false
	}
	return o.EmailProvider, true
}

// HasEmailProvider returns a boolean if a field has been set.
func (o *ModelEmailConfigurationAdd) HasEmailProvider() bool {
	if o != nil && !IsNil(o.EmailProvider) {
		return true
	}

	return false
}

// SetEmailProvider gets a reference to the given string and assigns it to the EmailProvider field.
func (o *ModelEmailConfigurationAdd) SetEmailProvider(v string) {
	o.EmailProvider = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ModelEmailConfigurationAdd) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEmailConfigurationAdd) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ModelEmailConfigurationAdd) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ModelEmailConfigurationAdd) SetPassword(v string) {
	o.Password = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ModelEmailConfigurationAdd) GetPort() string {
	if o == nil || IsNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEmailConfigurationAdd) GetPortOk() (*string, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ModelEmailConfigurationAdd) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *ModelEmailConfigurationAdd) SetPort(v string) {
	o.Port = &v
}

// GetSesRegion returns the SesRegion field value if set, zero value otherwise.
func (o *ModelEmailConfigurationAdd) GetSesRegion() string {
	if o == nil || IsNil(o.SesRegion) {
		var ret string
		return ret
	}
	return *o.SesRegion
}

// GetSesRegionOk returns a tuple with the SesRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEmailConfigurationAdd) GetSesRegionOk() (*string, bool) {
	if o == nil || IsNil(o.SesRegion) {
		return nil, false
	}
	return o.SesRegion, true
}

// HasSesRegion returns a boolean if a field has been set.
func (o *ModelEmailConfigurationAdd) HasSesRegion() bool {
	if o != nil && !IsNil(o.SesRegion) {
		return true
	}

	return false
}

// SetSesRegion gets a reference to the given string and assigns it to the SesRegion field.
func (o *ModelEmailConfigurationAdd) SetSesRegion(v string) {
	o.SesRegion = &v
}

// GetSmtp returns the Smtp field value if set, zero value otherwise.
func (o *ModelEmailConfigurationAdd) GetSmtp() string {
	if o == nil || IsNil(o.Smtp) {
		var ret string
		return ret
	}
	return *o.Smtp
}

// GetSmtpOk returns a tuple with the Smtp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEmailConfigurationAdd) GetSmtpOk() (*string, bool) {
	if o == nil || IsNil(o.Smtp) {
		return nil, false
	}
	return o.Smtp, true
}

// HasSmtp returns a boolean if a field has been set.
func (o *ModelEmailConfigurationAdd) HasSmtp() bool {
	if o != nil && !IsNil(o.Smtp) {
		return true
	}

	return false
}

// SetSmtp gets a reference to the given string and assigns it to the Smtp field.
func (o *ModelEmailConfigurationAdd) SetSmtp(v string) {
	o.Smtp = &v
}

func (o ModelEmailConfigurationAdd) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelEmailConfigurationAdd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AmazonAccessKey) {
		toSerialize["amazon_access_key"] = o.AmazonAccessKey
	}
	if !IsNil(o.AmazonSecretKey) {
		toSerialize["amazon_secret_key"] = o.AmazonSecretKey
	}
	if !IsNil(o.Apikey) {
		toSerialize["apikey"] = o.Apikey
	}
	if !IsNil(o.CreatedByUserId) {
		toSerialize["created_by_user_id"] = o.CreatedByUserId
	}
	if !IsNil(o.EmailId) {
		toSerialize["email_id"] = o.EmailId
	}
	if !IsNil(o.EmailProvider) {
		toSerialize["email_provider"] = o.EmailProvider
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.SesRegion) {
		toSerialize["ses_region"] = o.SesRegion
	}
	if !IsNil(o.Smtp) {
		toSerialize["smtp"] = o.Smtp
	}
	return toSerialize, nil
}

type NullableModelEmailConfigurationAdd struct {
	value *ModelEmailConfigurationAdd
	isSet bool
}

func (v NullableModelEmailConfigurationAdd) Get() *ModelEmailConfigurationAdd {
	return v.value
}

func (v *NullableModelEmailConfigurationAdd) Set(val *ModelEmailConfigurationAdd) {
	v.value = val
	v.isSet = true
}

func (v NullableModelEmailConfigurationAdd) IsSet() bool {
	return v.isSet
}

func (v *NullableModelEmailConfigurationAdd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelEmailConfigurationAdd(val *ModelEmailConfigurationAdd) *NullableModelEmailConfigurationAdd {
	return &NullableModelEmailConfigurationAdd{value: val, isSet: true}
}

func (v NullableModelEmailConfigurationAdd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelEmailConfigurationAdd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


