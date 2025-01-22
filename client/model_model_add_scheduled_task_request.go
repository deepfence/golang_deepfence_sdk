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

// checks if the ModelAddScheduledTaskRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelAddScheduledTaskRequest{}

// ModelAddScheduledTaskRequest struct for ModelAddScheduledTaskRequest
type ModelAddScheduledTaskRequest struct {
	Action string `json:"action"`
	BenchmarkTypes []ModelBenchmarkType `json:"benchmark_types"`
	CronExpr *string `json:"cron_expr,omitempty"`
	DeepfenceSystemScan *bool `json:"deepfence_system_scan,omitempty"`
	Description *string `json:"description,omitempty"`
	Filters ModelScanFilter `json:"filters"`
	IsPriority *bool `json:"is_priority,omitempty"`
	NodeIds []ModelNodeIdentifier `json:"node_ids"`
	ScanConfig []ModelVulnerabilityScanConfigLanguage `json:"scan_config"`
}

type _ModelAddScheduledTaskRequest ModelAddScheduledTaskRequest

// NewModelAddScheduledTaskRequest instantiates a new ModelAddScheduledTaskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelAddScheduledTaskRequest(action string, benchmarkTypes []ModelBenchmarkType, filters ModelScanFilter, nodeIds []ModelNodeIdentifier, scanConfig []ModelVulnerabilityScanConfigLanguage) *ModelAddScheduledTaskRequest {
	this := ModelAddScheduledTaskRequest{}
	this.Action = action
	this.BenchmarkTypes = benchmarkTypes
	this.Filters = filters
	this.NodeIds = nodeIds
	this.ScanConfig = scanConfig
	return &this
}

// NewModelAddScheduledTaskRequestWithDefaults instantiates a new ModelAddScheduledTaskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelAddScheduledTaskRequestWithDefaults() *ModelAddScheduledTaskRequest {
	this := ModelAddScheduledTaskRequest{}
	return &this
}

// GetAction returns the Action field value
func (o *ModelAddScheduledTaskRequest) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ModelAddScheduledTaskRequest) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ModelAddScheduledTaskRequest) SetAction(v string) {
	o.Action = v
}

// GetBenchmarkTypes returns the BenchmarkTypes field value
// If the value is explicit nil, the zero value for []ModelBenchmarkType will be returned
func (o *ModelAddScheduledTaskRequest) GetBenchmarkTypes() []ModelBenchmarkType {
	if o == nil {
		var ret []ModelBenchmarkType
		return ret
	}

	return o.BenchmarkTypes
}

// GetBenchmarkTypesOk returns a tuple with the BenchmarkTypes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelAddScheduledTaskRequest) GetBenchmarkTypesOk() ([]ModelBenchmarkType, bool) {
	if o == nil || IsNil(o.BenchmarkTypes) {
		return nil, false
	}
	return o.BenchmarkTypes, true
}

// SetBenchmarkTypes sets field value
func (o *ModelAddScheduledTaskRequest) SetBenchmarkTypes(v []ModelBenchmarkType) {
	o.BenchmarkTypes = v
}

// GetCronExpr returns the CronExpr field value if set, zero value otherwise.
func (o *ModelAddScheduledTaskRequest) GetCronExpr() string {
	if o == nil || IsNil(o.CronExpr) {
		var ret string
		return ret
	}
	return *o.CronExpr
}

// GetCronExprOk returns a tuple with the CronExpr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAddScheduledTaskRequest) GetCronExprOk() (*string, bool) {
	if o == nil || IsNil(o.CronExpr) {
		return nil, false
	}
	return o.CronExpr, true
}

// HasCronExpr returns a boolean if a field has been set.
func (o *ModelAddScheduledTaskRequest) HasCronExpr() bool {
	if o != nil && !IsNil(o.CronExpr) {
		return true
	}

	return false
}

// SetCronExpr gets a reference to the given string and assigns it to the CronExpr field.
func (o *ModelAddScheduledTaskRequest) SetCronExpr(v string) {
	o.CronExpr = &v
}

// GetDeepfenceSystemScan returns the DeepfenceSystemScan field value if set, zero value otherwise.
func (o *ModelAddScheduledTaskRequest) GetDeepfenceSystemScan() bool {
	if o == nil || IsNil(o.DeepfenceSystemScan) {
		var ret bool
		return ret
	}
	return *o.DeepfenceSystemScan
}

// GetDeepfenceSystemScanOk returns a tuple with the DeepfenceSystemScan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAddScheduledTaskRequest) GetDeepfenceSystemScanOk() (*bool, bool) {
	if o == nil || IsNil(o.DeepfenceSystemScan) {
		return nil, false
	}
	return o.DeepfenceSystemScan, true
}

// HasDeepfenceSystemScan returns a boolean if a field has been set.
func (o *ModelAddScheduledTaskRequest) HasDeepfenceSystemScan() bool {
	if o != nil && !IsNil(o.DeepfenceSystemScan) {
		return true
	}

	return false
}

// SetDeepfenceSystemScan gets a reference to the given bool and assigns it to the DeepfenceSystemScan field.
func (o *ModelAddScheduledTaskRequest) SetDeepfenceSystemScan(v bool) {
	o.DeepfenceSystemScan = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModelAddScheduledTaskRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAddScheduledTaskRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelAddScheduledTaskRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModelAddScheduledTaskRequest) SetDescription(v string) {
	o.Description = &v
}

// GetFilters returns the Filters field value
func (o *ModelAddScheduledTaskRequest) GetFilters() ModelScanFilter {
	if o == nil {
		var ret ModelScanFilter
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *ModelAddScheduledTaskRequest) GetFiltersOk() (*ModelScanFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filters, true
}

// SetFilters sets field value
func (o *ModelAddScheduledTaskRequest) SetFilters(v ModelScanFilter) {
	o.Filters = v
}

// GetIsPriority returns the IsPriority field value if set, zero value otherwise.
func (o *ModelAddScheduledTaskRequest) GetIsPriority() bool {
	if o == nil || IsNil(o.IsPriority) {
		var ret bool
		return ret
	}
	return *o.IsPriority
}

// GetIsPriorityOk returns a tuple with the IsPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAddScheduledTaskRequest) GetIsPriorityOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPriority) {
		return nil, false
	}
	return o.IsPriority, true
}

// HasIsPriority returns a boolean if a field has been set.
func (o *ModelAddScheduledTaskRequest) HasIsPriority() bool {
	if o != nil && !IsNil(o.IsPriority) {
		return true
	}

	return false
}

// SetIsPriority gets a reference to the given bool and assigns it to the IsPriority field.
func (o *ModelAddScheduledTaskRequest) SetIsPriority(v bool) {
	o.IsPriority = &v
}

// GetNodeIds returns the NodeIds field value
// If the value is explicit nil, the zero value for []ModelNodeIdentifier will be returned
func (o *ModelAddScheduledTaskRequest) GetNodeIds() []ModelNodeIdentifier {
	if o == nil {
		var ret []ModelNodeIdentifier
		return ret
	}

	return o.NodeIds
}

// GetNodeIdsOk returns a tuple with the NodeIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelAddScheduledTaskRequest) GetNodeIdsOk() ([]ModelNodeIdentifier, bool) {
	if o == nil || IsNil(o.NodeIds) {
		return nil, false
	}
	return o.NodeIds, true
}

// SetNodeIds sets field value
func (o *ModelAddScheduledTaskRequest) SetNodeIds(v []ModelNodeIdentifier) {
	o.NodeIds = v
}

// GetScanConfig returns the ScanConfig field value
// If the value is explicit nil, the zero value for []ModelVulnerabilityScanConfigLanguage will be returned
func (o *ModelAddScheduledTaskRequest) GetScanConfig() []ModelVulnerabilityScanConfigLanguage {
	if o == nil {
		var ret []ModelVulnerabilityScanConfigLanguage
		return ret
	}

	return o.ScanConfig
}

// GetScanConfigOk returns a tuple with the ScanConfig field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelAddScheduledTaskRequest) GetScanConfigOk() ([]ModelVulnerabilityScanConfigLanguage, bool) {
	if o == nil || IsNil(o.ScanConfig) {
		return nil, false
	}
	return o.ScanConfig, true
}

// SetScanConfig sets field value
func (o *ModelAddScheduledTaskRequest) SetScanConfig(v []ModelVulnerabilityScanConfigLanguage) {
	o.ScanConfig = v
}

func (o ModelAddScheduledTaskRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelAddScheduledTaskRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	if o.BenchmarkTypes != nil {
		toSerialize["benchmark_types"] = o.BenchmarkTypes
	}
	if !IsNil(o.CronExpr) {
		toSerialize["cron_expr"] = o.CronExpr
	}
	if !IsNil(o.DeepfenceSystemScan) {
		toSerialize["deepfence_system_scan"] = o.DeepfenceSystemScan
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["filters"] = o.Filters
	if !IsNil(o.IsPriority) {
		toSerialize["is_priority"] = o.IsPriority
	}
	if o.NodeIds != nil {
		toSerialize["node_ids"] = o.NodeIds
	}
	if o.ScanConfig != nil {
		toSerialize["scan_config"] = o.ScanConfig
	}
	return toSerialize, nil
}

func (o *ModelAddScheduledTaskRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"benchmark_types",
		"filters",
		"node_ids",
		"scan_config",
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

	varModelAddScheduledTaskRequest := _ModelAddScheduledTaskRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelAddScheduledTaskRequest)

	if err != nil {
		return err
	}

	*o = ModelAddScheduledTaskRequest(varModelAddScheduledTaskRequest)

	return err
}

type NullableModelAddScheduledTaskRequest struct {
	value *ModelAddScheduledTaskRequest
	isSet bool
}

func (v NullableModelAddScheduledTaskRequest) Get() *ModelAddScheduledTaskRequest {
	return v.value
}

func (v *NullableModelAddScheduledTaskRequest) Set(val *ModelAddScheduledTaskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelAddScheduledTaskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelAddScheduledTaskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelAddScheduledTaskRequest(val *ModelAddScheduledTaskRequest) *NullableModelAddScheduledTaskRequest {
	return &NullableModelAddScheduledTaskRequest{value: val, isSet: true}
}

func (v NullableModelAddScheduledTaskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelAddScheduledTaskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


