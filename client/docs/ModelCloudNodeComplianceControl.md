# ModelCloudNodeComplianceControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryHierarchy** | Pointer to **[]string** |  | [optional] 
**CategoryHierarchyShort** | Pointer to **string** |  | [optional] 
**ComplianceType** | Pointer to **string** |  | [optional] 
**ControlId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**ProblemTitle** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewModelCloudNodeComplianceControl

`func NewModelCloudNodeComplianceControl() *ModelCloudNodeComplianceControl`

NewModelCloudNodeComplianceControl instantiates a new ModelCloudNodeComplianceControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCloudNodeComplianceControlWithDefaults

`func NewModelCloudNodeComplianceControlWithDefaults() *ModelCloudNodeComplianceControl`

NewModelCloudNodeComplianceControlWithDefaults instantiates a new ModelCloudNodeComplianceControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryHierarchy

`func (o *ModelCloudNodeComplianceControl) GetCategoryHierarchy() []string`

GetCategoryHierarchy returns the CategoryHierarchy field if non-nil, zero value otherwise.

### GetCategoryHierarchyOk

`func (o *ModelCloudNodeComplianceControl) GetCategoryHierarchyOk() (*[]string, bool)`

GetCategoryHierarchyOk returns a tuple with the CategoryHierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryHierarchy

`func (o *ModelCloudNodeComplianceControl) SetCategoryHierarchy(v []string)`

SetCategoryHierarchy sets CategoryHierarchy field to given value.

### HasCategoryHierarchy

`func (o *ModelCloudNodeComplianceControl) HasCategoryHierarchy() bool`

HasCategoryHierarchy returns a boolean if a field has been set.

### SetCategoryHierarchyNil

`func (o *ModelCloudNodeComplianceControl) SetCategoryHierarchyNil(b bool)`

 SetCategoryHierarchyNil sets the value for CategoryHierarchy to be an explicit nil

### UnsetCategoryHierarchy
`func (o *ModelCloudNodeComplianceControl) UnsetCategoryHierarchy()`

UnsetCategoryHierarchy ensures that no value is present for CategoryHierarchy, not even an explicit nil
### GetCategoryHierarchyShort

`func (o *ModelCloudNodeComplianceControl) GetCategoryHierarchyShort() string`

GetCategoryHierarchyShort returns the CategoryHierarchyShort field if non-nil, zero value otherwise.

### GetCategoryHierarchyShortOk

`func (o *ModelCloudNodeComplianceControl) GetCategoryHierarchyShortOk() (*string, bool)`

GetCategoryHierarchyShortOk returns a tuple with the CategoryHierarchyShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryHierarchyShort

`func (o *ModelCloudNodeComplianceControl) SetCategoryHierarchyShort(v string)`

SetCategoryHierarchyShort sets CategoryHierarchyShort field to given value.

### HasCategoryHierarchyShort

`func (o *ModelCloudNodeComplianceControl) HasCategoryHierarchyShort() bool`

HasCategoryHierarchyShort returns a boolean if a field has been set.

### GetComplianceType

`func (o *ModelCloudNodeComplianceControl) GetComplianceType() string`

GetComplianceType returns the ComplianceType field if non-nil, zero value otherwise.

### GetComplianceTypeOk

`func (o *ModelCloudNodeComplianceControl) GetComplianceTypeOk() (*string, bool)`

GetComplianceTypeOk returns a tuple with the ComplianceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceType

`func (o *ModelCloudNodeComplianceControl) SetComplianceType(v string)`

SetComplianceType sets ComplianceType field to given value.

### HasComplianceType

`func (o *ModelCloudNodeComplianceControl) HasComplianceType() bool`

HasComplianceType returns a boolean if a field has been set.

### GetControlId

`func (o *ModelCloudNodeComplianceControl) GetControlId() string`

GetControlId returns the ControlId field if non-nil, zero value otherwise.

### GetControlIdOk

`func (o *ModelCloudNodeComplianceControl) GetControlIdOk() (*string, bool)`

GetControlIdOk returns a tuple with the ControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlId

`func (o *ModelCloudNodeComplianceControl) SetControlId(v string)`

SetControlId sets ControlId field to given value.

### HasControlId

`func (o *ModelCloudNodeComplianceControl) HasControlId() bool`

HasControlId returns a boolean if a field has been set.

### GetDescription

`func (o *ModelCloudNodeComplianceControl) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelCloudNodeComplianceControl) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelCloudNodeComplianceControl) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelCloudNodeComplianceControl) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ModelCloudNodeComplianceControl) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ModelCloudNodeComplianceControl) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ModelCloudNodeComplianceControl) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ModelCloudNodeComplianceControl) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNodeId

`func (o *ModelCloudNodeComplianceControl) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelCloudNodeComplianceControl) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelCloudNodeComplianceControl) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ModelCloudNodeComplianceControl) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetProblemTitle

`func (o *ModelCloudNodeComplianceControl) GetProblemTitle() string`

GetProblemTitle returns the ProblemTitle field if non-nil, zero value otherwise.

### GetProblemTitleOk

`func (o *ModelCloudNodeComplianceControl) GetProblemTitleOk() (*string, bool)`

GetProblemTitleOk returns a tuple with the ProblemTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemTitle

`func (o *ModelCloudNodeComplianceControl) SetProblemTitle(v string)`

SetProblemTitle sets ProblemTitle field to given value.

### HasProblemTitle

`func (o *ModelCloudNodeComplianceControl) HasProblemTitle() bool`

HasProblemTitle returns a boolean if a field has been set.

### GetService

`func (o *ModelCloudNodeComplianceControl) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ModelCloudNodeComplianceControl) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ModelCloudNodeComplianceControl) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ModelCloudNodeComplianceControl) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTitle

`func (o *ModelCloudNodeComplianceControl) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModelCloudNodeComplianceControl) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModelCloudNodeComplianceControl) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ModelCloudNodeComplianceControl) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


