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
	"bytes"
	"fmt"
)

// checks if the ModelUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelUser{}

// ModelUser struct for ModelUser
type ModelUser struct {
	Company string `json:"company"`
	CompanyId *int32 `json:"company_id,omitempty"`
	CurrentUser NullableBool `json:"current_user,omitempty"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	Groups map[string]string `json:"groups,omitempty"`
	Id *int32 `json:"id,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	LastName string `json:"last_name"`
	PasswordInvalidated *bool `json:"password_invalidated,omitempty"`
	Role *string `json:"role,omitempty"`
	RoleId *int32 `json:"role_id,omitempty"`
}

type _ModelUser ModelUser

// NewModelUser instantiates a new ModelUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelUser(company string, email string, firstName string, lastName string) *ModelUser {
	this := ModelUser{}
	this.Company = company
	this.Email = email
	this.FirstName = firstName
	this.LastName = lastName
	return &this
}

// NewModelUserWithDefaults instantiates a new ModelUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelUserWithDefaults() *ModelUser {
	this := ModelUser{}
	return &this
}

// GetCompany returns the Company field value
func (o *ModelUser) GetCompany() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Company
}

// GetCompanyOk returns a tuple with the Company field value
// and a boolean to check if the value has been set.
func (o *ModelUser) GetCompanyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Company, true
}

// SetCompany sets field value
func (o *ModelUser) SetCompany(v string) {
	o.Company = v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *ModelUser) GetCompanyId() int32 {
	if o == nil || IsNil(o.CompanyId) {
		var ret int32
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUser) GetCompanyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CompanyId) {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *ModelUser) HasCompanyId() bool {
	if o != nil && !IsNil(o.CompanyId) {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given int32 and assigns it to the CompanyId field.
func (o *ModelUser) SetCompanyId(v int32) {
	o.CompanyId = &v
}

// GetCurrentUser returns the CurrentUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelUser) GetCurrentUser() bool {
	if o == nil || IsNil(o.CurrentUser.Get()) {
		var ret bool
		return ret
	}
	return *o.CurrentUser.Get()
}

// GetCurrentUserOk returns a tuple with the CurrentUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelUser) GetCurrentUserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentUser.Get(), o.CurrentUser.IsSet()
}

// HasCurrentUser returns a boolean if a field has been set.
func (o *ModelUser) HasCurrentUser() bool {
	if o != nil && o.CurrentUser.IsSet() {
		return true
	}

	return false
}

// SetCurrentUser gets a reference to the given NullableBool and assigns it to the CurrentUser field.
func (o *ModelUser) SetCurrentUser(v bool) {
	o.CurrentUser.Set(&v)
}
// SetCurrentUserNil sets the value for CurrentUser to be an explicit nil
func (o *ModelUser) SetCurrentUserNil() {
	o.CurrentUser.Set(nil)
}

// UnsetCurrentUser ensures that no value is present for CurrentUser, not even an explicit nil
func (o *ModelUser) UnsetCurrentUser() {
	o.CurrentUser.Unset()
}

// GetEmail returns the Email field value
func (o *ModelUser) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ModelUser) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ModelUser) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value
func (o *ModelUser) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *ModelUser) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *ModelUser) SetFirstName(v string) {
	o.FirstName = v
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelUser) GetGroups() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelUser) GetGroupsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return &o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ModelUser) HasGroups() bool {
	if o != nil && IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given map[string]string and assigns it to the Groups field.
func (o *ModelUser) SetGroups(v map[string]string) {
	o.Groups = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelUser) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUser) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelUser) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelUser) SetId(v int32) {
	o.Id = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *ModelUser) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUser) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *ModelUser) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *ModelUser) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetLastName returns the LastName field value
func (o *ModelUser) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *ModelUser) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *ModelUser) SetLastName(v string) {
	o.LastName = v
}

// GetPasswordInvalidated returns the PasswordInvalidated field value if set, zero value otherwise.
func (o *ModelUser) GetPasswordInvalidated() bool {
	if o == nil || IsNil(o.PasswordInvalidated) {
		var ret bool
		return ret
	}
	return *o.PasswordInvalidated
}

// GetPasswordInvalidatedOk returns a tuple with the PasswordInvalidated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUser) GetPasswordInvalidatedOk() (*bool, bool) {
	if o == nil || IsNil(o.PasswordInvalidated) {
		return nil, false
	}
	return o.PasswordInvalidated, true
}

// HasPasswordInvalidated returns a boolean if a field has been set.
func (o *ModelUser) HasPasswordInvalidated() bool {
	if o != nil && !IsNil(o.PasswordInvalidated) {
		return true
	}

	return false
}

// SetPasswordInvalidated gets a reference to the given bool and assigns it to the PasswordInvalidated field.
func (o *ModelUser) SetPasswordInvalidated(v bool) {
	o.PasswordInvalidated = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ModelUser) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUser) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ModelUser) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *ModelUser) SetRole(v string) {
	o.Role = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *ModelUser) GetRoleId() int32 {
	if o == nil || IsNil(o.RoleId) {
		var ret int32
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUser) GetRoleIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RoleId) {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *ModelUser) HasRoleId() bool {
	if o != nil && !IsNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given int32 and assigns it to the RoleId field.
func (o *ModelUser) SetRoleId(v int32) {
	o.RoleId = &v
}

func (o ModelUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["company"] = o.Company
	if !IsNil(o.CompanyId) {
		toSerialize["company_id"] = o.CompanyId
	}
	if o.CurrentUser.IsSet() {
		toSerialize["current_user"] = o.CurrentUser.Get()
	}
	toSerialize["email"] = o.Email
	toSerialize["first_name"] = o.FirstName
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	toSerialize["last_name"] = o.LastName
	if !IsNil(o.PasswordInvalidated) {
		toSerialize["password_invalidated"] = o.PasswordInvalidated
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.RoleId) {
		toSerialize["role_id"] = o.RoleId
	}
	return toSerialize, nil
}

func (o *ModelUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"company",
		"email",
		"first_name",
		"last_name",
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

	varModelUser := _ModelUser{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelUser)

	if err != nil {
		return err
	}

	*o = ModelUser(varModelUser)

	return err
}

type NullableModelUser struct {
	value *ModelUser
	isSet bool
}

func (v NullableModelUser) Get() *ModelUser {
	return v.value
}

func (v *NullableModelUser) Set(val *ModelUser) {
	v.value = val
	v.isSet = true
}

func (v NullableModelUser) IsSet() bool {
	return v.isSet
}

func (v *NullableModelUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelUser(val *ModelUser) *NullableModelUser {
	return &NullableModelUser{value: val, isSet: true}
}

func (v NullableModelUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


