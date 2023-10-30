# ModelCompliance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComplianceCheckType** | **string** |  | 
**Description** | **string** |  | 
**Masked** | **bool** |  | 
**NodeId** | **string** |  | 
**NodeType** | **string** |  | 
**RemediationAnsible** | **string** |  | 
**RemediationPuppet** | **string** |  | 
**RemediationScript** | **string** |  | 
**Resource** | **string** |  | 
**Resources** | Pointer to [**[]ModelBasicNode**](ModelBasicNode.md) |  | [optional] 
**RuleId** | **string** |  | 
**Status** | **string** |  | 
**TestCategory** | **string** |  | 
**TestDesc** | **string** |  | 
**TestNumber** | **string** |  | 
**TestRationale** | **string** |  | 
**TestSeverity** | **string** |  | 
**UpdatedAt** | **int32** |  | 

## Methods

### NewModelCompliance

`func NewModelCompliance(complianceCheckType string, description string, masked bool, nodeId string, nodeType string, remediationAnsible string, remediationPuppet string, remediationScript string, resource string, ruleId string, status string, testCategory string, testDesc string, testNumber string, testRationale string, testSeverity string, updatedAt int32, ) *ModelCompliance`

NewModelCompliance instantiates a new ModelCompliance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelComplianceWithDefaults

`func NewModelComplianceWithDefaults() *ModelCompliance`

NewModelComplianceWithDefaults instantiates a new ModelCompliance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComplianceCheckType

`func (o *ModelCompliance) GetComplianceCheckType() string`

GetComplianceCheckType returns the ComplianceCheckType field if non-nil, zero value otherwise.

### GetComplianceCheckTypeOk

`func (o *ModelCompliance) GetComplianceCheckTypeOk() (*string, bool)`

GetComplianceCheckTypeOk returns a tuple with the ComplianceCheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceCheckType

`func (o *ModelCompliance) SetComplianceCheckType(v string)`

SetComplianceCheckType sets ComplianceCheckType field to given value.


### GetDescription

`func (o *ModelCompliance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelCompliance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelCompliance) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMasked

`func (o *ModelCompliance) GetMasked() bool`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *ModelCompliance) GetMaskedOk() (*bool, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *ModelCompliance) SetMasked(v bool)`

SetMasked sets Masked field to given value.


### GetNodeId

`func (o *ModelCompliance) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelCompliance) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelCompliance) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNodeType

`func (o *ModelCompliance) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *ModelCompliance) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *ModelCompliance) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetRemediationAnsible

`func (o *ModelCompliance) GetRemediationAnsible() string`

GetRemediationAnsible returns the RemediationAnsible field if non-nil, zero value otherwise.

### GetRemediationAnsibleOk

`func (o *ModelCompliance) GetRemediationAnsibleOk() (*string, bool)`

GetRemediationAnsibleOk returns a tuple with the RemediationAnsible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationAnsible

`func (o *ModelCompliance) SetRemediationAnsible(v string)`

SetRemediationAnsible sets RemediationAnsible field to given value.


### GetRemediationPuppet

`func (o *ModelCompliance) GetRemediationPuppet() string`

GetRemediationPuppet returns the RemediationPuppet field if non-nil, zero value otherwise.

### GetRemediationPuppetOk

`func (o *ModelCompliance) GetRemediationPuppetOk() (*string, bool)`

GetRemediationPuppetOk returns a tuple with the RemediationPuppet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationPuppet

`func (o *ModelCompliance) SetRemediationPuppet(v string)`

SetRemediationPuppet sets RemediationPuppet field to given value.


### GetRemediationScript

`func (o *ModelCompliance) GetRemediationScript() string`

GetRemediationScript returns the RemediationScript field if non-nil, zero value otherwise.

### GetRemediationScriptOk

`func (o *ModelCompliance) GetRemediationScriptOk() (*string, bool)`

GetRemediationScriptOk returns a tuple with the RemediationScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationScript

`func (o *ModelCompliance) SetRemediationScript(v string)`

SetRemediationScript sets RemediationScript field to given value.


### GetResource

`func (o *ModelCompliance) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ModelCompliance) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ModelCompliance) SetResource(v string)`

SetResource sets Resource field to given value.


### GetResources

`func (o *ModelCompliance) GetResources() []ModelBasicNode`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ModelCompliance) GetResourcesOk() (*[]ModelBasicNode, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ModelCompliance) SetResources(v []ModelBasicNode)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ModelCompliance) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResourcesNil

`func (o *ModelCompliance) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *ModelCompliance) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil
### GetRuleId

`func (o *ModelCompliance) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *ModelCompliance) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *ModelCompliance) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetStatus

`func (o *ModelCompliance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelCompliance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelCompliance) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTestCategory

`func (o *ModelCompliance) GetTestCategory() string`

GetTestCategory returns the TestCategory field if non-nil, zero value otherwise.

### GetTestCategoryOk

`func (o *ModelCompliance) GetTestCategoryOk() (*string, bool)`

GetTestCategoryOk returns a tuple with the TestCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCategory

`func (o *ModelCompliance) SetTestCategory(v string)`

SetTestCategory sets TestCategory field to given value.


### GetTestDesc

`func (o *ModelCompliance) GetTestDesc() string`

GetTestDesc returns the TestDesc field if non-nil, zero value otherwise.

### GetTestDescOk

`func (o *ModelCompliance) GetTestDescOk() (*string, bool)`

GetTestDescOk returns a tuple with the TestDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestDesc

`func (o *ModelCompliance) SetTestDesc(v string)`

SetTestDesc sets TestDesc field to given value.


### GetTestNumber

`func (o *ModelCompliance) GetTestNumber() string`

GetTestNumber returns the TestNumber field if non-nil, zero value otherwise.

### GetTestNumberOk

`func (o *ModelCompliance) GetTestNumberOk() (*string, bool)`

GetTestNumberOk returns a tuple with the TestNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestNumber

`func (o *ModelCompliance) SetTestNumber(v string)`

SetTestNumber sets TestNumber field to given value.


### GetTestRationale

`func (o *ModelCompliance) GetTestRationale() string`

GetTestRationale returns the TestRationale field if non-nil, zero value otherwise.

### GetTestRationaleOk

`func (o *ModelCompliance) GetTestRationaleOk() (*string, bool)`

GetTestRationaleOk returns a tuple with the TestRationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRationale

`func (o *ModelCompliance) SetTestRationale(v string)`

SetTestRationale sets TestRationale field to given value.


### GetTestSeverity

`func (o *ModelCompliance) GetTestSeverity() string`

GetTestSeverity returns the TestSeverity field if non-nil, zero value otherwise.

### GetTestSeverityOk

`func (o *ModelCompliance) GetTestSeverityOk() (*string, bool)`

GetTestSeverityOk returns a tuple with the TestSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestSeverity

`func (o *ModelCompliance) SetTestSeverity(v string)`

SetTestSeverity sets TestSeverity field to given value.


### GetUpdatedAt

`func (o *ModelCompliance) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelCompliance) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelCompliance) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


