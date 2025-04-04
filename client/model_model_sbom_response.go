/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.5.6
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ModelSbomResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSbomResponse{}

// ModelSbomResponse struct for ModelSbomResponse
type ModelSbomResponse struct {
	CveId *string `json:"cve_id,omitempty"`
	CveNodeId *string `json:"cve_node_id,omitempty"`
	Licenses []string `json:"licenses,omitempty"`
	Locations []string `json:"locations,omitempty"`
	PackageName *string `json:"package_name,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewModelSbomResponse instantiates a new ModelSbomResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSbomResponse() *ModelSbomResponse {
	this := ModelSbomResponse{}
	return &this
}

// NewModelSbomResponseWithDefaults instantiates a new ModelSbomResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSbomResponseWithDefaults() *ModelSbomResponse {
	this := ModelSbomResponse{}
	return &this
}

// GetCveId returns the CveId field value if set, zero value otherwise.
func (o *ModelSbomResponse) GetCveId() string {
	if o == nil || IsNil(o.CveId) {
		var ret string
		return ret
	}
	return *o.CveId
}

// GetCveIdOk returns a tuple with the CveId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSbomResponse) GetCveIdOk() (*string, bool) {
	if o == nil || IsNil(o.CveId) {
		return nil, false
	}
	return o.CveId, true
}

// HasCveId returns a boolean if a field has been set.
func (o *ModelSbomResponse) HasCveId() bool {
	if o != nil && !IsNil(o.CveId) {
		return true
	}

	return false
}

// SetCveId gets a reference to the given string and assigns it to the CveId field.
func (o *ModelSbomResponse) SetCveId(v string) {
	o.CveId = &v
}

// GetCveNodeId returns the CveNodeId field value if set, zero value otherwise.
func (o *ModelSbomResponse) GetCveNodeId() string {
	if o == nil || IsNil(o.CveNodeId) {
		var ret string
		return ret
	}
	return *o.CveNodeId
}

// GetCveNodeIdOk returns a tuple with the CveNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSbomResponse) GetCveNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.CveNodeId) {
		return nil, false
	}
	return o.CveNodeId, true
}

// HasCveNodeId returns a boolean if a field has been set.
func (o *ModelSbomResponse) HasCveNodeId() bool {
	if o != nil && !IsNil(o.CveNodeId) {
		return true
	}

	return false
}

// SetCveNodeId gets a reference to the given string and assigns it to the CveNodeId field.
func (o *ModelSbomResponse) SetCveNodeId(v string) {
	o.CveNodeId = &v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *ModelSbomResponse) GetLicenses() []string {
	if o == nil || IsNil(o.Licenses) {
		var ret []string
		return ret
	}
	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSbomResponse) GetLicensesOk() ([]string, bool) {
	if o == nil || IsNil(o.Licenses) {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *ModelSbomResponse) HasLicenses() bool {
	if o != nil && !IsNil(o.Licenses) {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given []string and assigns it to the Licenses field.
func (o *ModelSbomResponse) SetLicenses(v []string) {
	o.Licenses = v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *ModelSbomResponse) GetLocations() []string {
	if o == nil || IsNil(o.Locations) {
		var ret []string
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSbomResponse) GetLocationsOk() ([]string, bool) {
	if o == nil || IsNil(o.Locations) {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *ModelSbomResponse) HasLocations() bool {
	if o != nil && !IsNil(o.Locations) {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []string and assigns it to the Locations field.
func (o *ModelSbomResponse) SetLocations(v []string) {
	o.Locations = v
}

// GetPackageName returns the PackageName field value if set, zero value otherwise.
func (o *ModelSbomResponse) GetPackageName() string {
	if o == nil || IsNil(o.PackageName) {
		var ret string
		return ret
	}
	return *o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSbomResponse) GetPackageNameOk() (*string, bool) {
	if o == nil || IsNil(o.PackageName) {
		return nil, false
	}
	return o.PackageName, true
}

// HasPackageName returns a boolean if a field has been set.
func (o *ModelSbomResponse) HasPackageName() bool {
	if o != nil && !IsNil(o.PackageName) {
		return true
	}

	return false
}

// SetPackageName gets a reference to the given string and assigns it to the PackageName field.
func (o *ModelSbomResponse) SetPackageName(v string) {
	o.PackageName = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ModelSbomResponse) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSbomResponse) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ModelSbomResponse) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *ModelSbomResponse) SetSeverity(v string) {
	o.Severity = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ModelSbomResponse) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSbomResponse) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ModelSbomResponse) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ModelSbomResponse) SetVersion(v string) {
	o.Version = &v
}

func (o ModelSbomResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSbomResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CveId) {
		toSerialize["cve_id"] = o.CveId
	}
	if !IsNil(o.CveNodeId) {
		toSerialize["cve_node_id"] = o.CveNodeId
	}
	if !IsNil(o.Licenses) {
		toSerialize["licenses"] = o.Licenses
	}
	if !IsNil(o.Locations) {
		toSerialize["locations"] = o.Locations
	}
	if !IsNil(o.PackageName) {
		toSerialize["package_name"] = o.PackageName
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableModelSbomResponse struct {
	value *ModelSbomResponse
	isSet bool
}

func (v NullableModelSbomResponse) Get() *ModelSbomResponse {
	return v.value
}

func (v *NullableModelSbomResponse) Set(val *ModelSbomResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSbomResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSbomResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSbomResponse(val *ModelSbomResponse) *NullableModelSbomResponse {
	return &NullableModelSbomResponse{value: val, isSet: true}
}

func (v NullableModelSbomResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSbomResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


