# ModelCloudComplianceControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**CategoryHierarchy** | Pointer to **[]string** |  | [optional] 
**CategoryHierarchyShort** | Pointer to **string** |  | [optional] 
**CloudProvider** | Pointer to **string** |  | [optional] 
**ComplianceType** | Pointer to **string** |  | [optional] 
**ControlId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Documentation** | Pointer to **string** |  | [optional] 
**Executable** | Pointer to **bool** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**ParentControlHierarchy** | Pointer to **[]string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewModelCloudComplianceControl

`func NewModelCloudComplianceControl() *ModelCloudComplianceControl`

NewModelCloudComplianceControl instantiates a new ModelCloudComplianceControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCloudComplianceControlWithDefaults

`func NewModelCloudComplianceControlWithDefaults() *ModelCloudComplianceControl`

NewModelCloudComplianceControlWithDefaults instantiates a new ModelCloudComplianceControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ModelCloudComplianceControl) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ModelCloudComplianceControl) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ModelCloudComplianceControl) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ModelCloudComplianceControl) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCategory

`func (o *ModelCloudComplianceControl) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ModelCloudComplianceControl) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ModelCloudComplianceControl) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ModelCloudComplianceControl) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCategoryHierarchy

`func (o *ModelCloudComplianceControl) GetCategoryHierarchy() []string`

GetCategoryHierarchy returns the CategoryHierarchy field if non-nil, zero value otherwise.

### GetCategoryHierarchyOk

`func (o *ModelCloudComplianceControl) GetCategoryHierarchyOk() (*[]string, bool)`

GetCategoryHierarchyOk returns a tuple with the CategoryHierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryHierarchy

`func (o *ModelCloudComplianceControl) SetCategoryHierarchy(v []string)`

SetCategoryHierarchy sets CategoryHierarchy field to given value.

### HasCategoryHierarchy

`func (o *ModelCloudComplianceControl) HasCategoryHierarchy() bool`

HasCategoryHierarchy returns a boolean if a field has been set.

### SetCategoryHierarchyNil

`func (o *ModelCloudComplianceControl) SetCategoryHierarchyNil(b bool)`

 SetCategoryHierarchyNil sets the value for CategoryHierarchy to be an explicit nil

### UnsetCategoryHierarchy
`func (o *ModelCloudComplianceControl) UnsetCategoryHierarchy()`

UnsetCategoryHierarchy ensures that no value is present for CategoryHierarchy, not even an explicit nil
### GetCategoryHierarchyShort

`func (o *ModelCloudComplianceControl) GetCategoryHierarchyShort() string`

GetCategoryHierarchyShort returns the CategoryHierarchyShort field if non-nil, zero value otherwise.

### GetCategoryHierarchyShortOk

`func (o *ModelCloudComplianceControl) GetCategoryHierarchyShortOk() (*string, bool)`

GetCategoryHierarchyShortOk returns a tuple with the CategoryHierarchyShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryHierarchyShort

`func (o *ModelCloudComplianceControl) SetCategoryHierarchyShort(v string)`

SetCategoryHierarchyShort sets CategoryHierarchyShort field to given value.

### HasCategoryHierarchyShort

`func (o *ModelCloudComplianceControl) HasCategoryHierarchyShort() bool`

HasCategoryHierarchyShort returns a boolean if a field has been set.

### GetCloudProvider

`func (o *ModelCloudComplianceControl) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ModelCloudComplianceControl) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ModelCloudComplianceControl) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *ModelCloudComplianceControl) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetComplianceType

`func (o *ModelCloudComplianceControl) GetComplianceType() string`

GetComplianceType returns the ComplianceType field if non-nil, zero value otherwise.

### GetComplianceTypeOk

`func (o *ModelCloudComplianceControl) GetComplianceTypeOk() (*string, bool)`

GetComplianceTypeOk returns a tuple with the ComplianceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceType

`func (o *ModelCloudComplianceControl) SetComplianceType(v string)`

SetComplianceType sets ComplianceType field to given value.

### HasComplianceType

`func (o *ModelCloudComplianceControl) HasComplianceType() bool`

HasComplianceType returns a boolean if a field has been set.

### GetControlId

`func (o *ModelCloudComplianceControl) GetControlId() string`

GetControlId returns the ControlId field if non-nil, zero value otherwise.

### GetControlIdOk

`func (o *ModelCloudComplianceControl) GetControlIdOk() (*string, bool)`

GetControlIdOk returns a tuple with the ControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlId

`func (o *ModelCloudComplianceControl) SetControlId(v string)`

SetControlId sets ControlId field to given value.

### HasControlId

`func (o *ModelCloudComplianceControl) HasControlId() bool`

HasControlId returns a boolean if a field has been set.

### GetDescription

`func (o *ModelCloudComplianceControl) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelCloudComplianceControl) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelCloudComplianceControl) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelCloudComplianceControl) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisabled

`func (o *ModelCloudComplianceControl) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ModelCloudComplianceControl) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ModelCloudComplianceControl) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ModelCloudComplianceControl) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDocumentation

`func (o *ModelCloudComplianceControl) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *ModelCloudComplianceControl) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *ModelCloudComplianceControl) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *ModelCloudComplianceControl) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetExecutable

`func (o *ModelCloudComplianceControl) GetExecutable() bool`

GetExecutable returns the Executable field if non-nil, zero value otherwise.

### GetExecutableOk

`func (o *ModelCloudComplianceControl) GetExecutableOk() (*bool, bool)`

GetExecutableOk returns a tuple with the Executable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutable

`func (o *ModelCloudComplianceControl) SetExecutable(v bool)`

SetExecutable sets Executable field to given value.

### HasExecutable

`func (o *ModelCloudComplianceControl) HasExecutable() bool`

HasExecutable returns a boolean if a field has been set.

### GetNodeId

`func (o *ModelCloudComplianceControl) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelCloudComplianceControl) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelCloudComplianceControl) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ModelCloudComplianceControl) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetParentControlHierarchy

`func (o *ModelCloudComplianceControl) GetParentControlHierarchy() []string`

GetParentControlHierarchy returns the ParentControlHierarchy field if non-nil, zero value otherwise.

### GetParentControlHierarchyOk

`func (o *ModelCloudComplianceControl) GetParentControlHierarchyOk() (*[]string, bool)`

GetParentControlHierarchyOk returns a tuple with the ParentControlHierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentControlHierarchy

`func (o *ModelCloudComplianceControl) SetParentControlHierarchy(v []string)`

SetParentControlHierarchy sets ParentControlHierarchy field to given value.

### HasParentControlHierarchy

`func (o *ModelCloudComplianceControl) HasParentControlHierarchy() bool`

HasParentControlHierarchy returns a boolean if a field has been set.

### SetParentControlHierarchyNil

`func (o *ModelCloudComplianceControl) SetParentControlHierarchyNil(b bool)`

 SetParentControlHierarchyNil sets the value for ParentControlHierarchy to be an explicit nil

### UnsetParentControlHierarchy
`func (o *ModelCloudComplianceControl) UnsetParentControlHierarchy()`

UnsetParentControlHierarchy ensures that no value is present for ParentControlHierarchy, not even an explicit nil
### GetService

`func (o *ModelCloudComplianceControl) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ModelCloudComplianceControl) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ModelCloudComplianceControl) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ModelCloudComplianceControl) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTitle

`func (o *ModelCloudComplianceControl) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModelCloudComplianceControl) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModelCloudComplianceControl) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ModelCloudComplianceControl) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


