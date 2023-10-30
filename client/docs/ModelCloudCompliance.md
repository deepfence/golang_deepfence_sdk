# ModelCloudCompliance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**CloudProvider** | **string** |  | 
**ComplianceCheckType** | **string** |  | 
**ControlId** | **string** |  | 
**Count** | **int32** |  | 
**Description** | **string** |  | 
**Group** | **string** |  | 
**Masked** | **bool** |  | 
**NodeId** | **string** |  | 
**NodeName** | **string** |  | 
**Reason** | **string** |  | 
**Region** | **string** |  | 
**Resource** | **string** |  | 
**Resources** | Pointer to [**[]ModelBasicNode**](ModelBasicNode.md) |  | [optional] 
**Service** | **string** |  | 
**Severity** | **string** |  | 
**Status** | **string** |  | 
**Title** | **string** |  | 
**Type** | **string** |  | 
**UpdatedAt** | **int32** |  | 

## Methods

### NewModelCloudCompliance

`func NewModelCloudCompliance(accountId string, cloudProvider string, complianceCheckType string, controlId string, count int32, description string, group string, masked bool, nodeId string, nodeName string, reason string, region string, resource string, service string, severity string, status string, title string, type_ string, updatedAt int32, ) *ModelCloudCompliance`

NewModelCloudCompliance instantiates a new ModelCloudCompliance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCloudComplianceWithDefaults

`func NewModelCloudComplianceWithDefaults() *ModelCloudCompliance`

NewModelCloudComplianceWithDefaults instantiates a new ModelCloudCompliance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ModelCloudCompliance) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ModelCloudCompliance) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ModelCloudCompliance) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCloudProvider

`func (o *ModelCloudCompliance) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ModelCloudCompliance) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ModelCloudCompliance) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetComplianceCheckType

`func (o *ModelCloudCompliance) GetComplianceCheckType() string`

GetComplianceCheckType returns the ComplianceCheckType field if non-nil, zero value otherwise.

### GetComplianceCheckTypeOk

`func (o *ModelCloudCompliance) GetComplianceCheckTypeOk() (*string, bool)`

GetComplianceCheckTypeOk returns a tuple with the ComplianceCheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceCheckType

`func (o *ModelCloudCompliance) SetComplianceCheckType(v string)`

SetComplianceCheckType sets ComplianceCheckType field to given value.


### GetControlId

`func (o *ModelCloudCompliance) GetControlId() string`

GetControlId returns the ControlId field if non-nil, zero value otherwise.

### GetControlIdOk

`func (o *ModelCloudCompliance) GetControlIdOk() (*string, bool)`

GetControlIdOk returns a tuple with the ControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlId

`func (o *ModelCloudCompliance) SetControlId(v string)`

SetControlId sets ControlId field to given value.


### GetCount

`func (o *ModelCloudCompliance) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ModelCloudCompliance) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ModelCloudCompliance) SetCount(v int32)`

SetCount sets Count field to given value.


### GetDescription

`func (o *ModelCloudCompliance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelCloudCompliance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelCloudCompliance) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetGroup

`func (o *ModelCloudCompliance) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ModelCloudCompliance) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ModelCloudCompliance) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetMasked

`func (o *ModelCloudCompliance) GetMasked() bool`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *ModelCloudCompliance) GetMaskedOk() (*bool, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *ModelCloudCompliance) SetMasked(v bool)`

SetMasked sets Masked field to given value.


### GetNodeId

`func (o *ModelCloudCompliance) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelCloudCompliance) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelCloudCompliance) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNodeName

`func (o *ModelCloudCompliance) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *ModelCloudCompliance) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *ModelCloudCompliance) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.


### GetReason

`func (o *ModelCloudCompliance) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ModelCloudCompliance) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ModelCloudCompliance) SetReason(v string)`

SetReason sets Reason field to given value.


### GetRegion

`func (o *ModelCloudCompliance) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ModelCloudCompliance) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ModelCloudCompliance) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetResource

`func (o *ModelCloudCompliance) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ModelCloudCompliance) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ModelCloudCompliance) SetResource(v string)`

SetResource sets Resource field to given value.


### GetResources

`func (o *ModelCloudCompliance) GetResources() []ModelBasicNode`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ModelCloudCompliance) GetResourcesOk() (*[]ModelBasicNode, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ModelCloudCompliance) SetResources(v []ModelBasicNode)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ModelCloudCompliance) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResourcesNil

`func (o *ModelCloudCompliance) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *ModelCloudCompliance) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil
### GetService

`func (o *ModelCloudCompliance) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ModelCloudCompliance) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ModelCloudCompliance) SetService(v string)`

SetService sets Service field to given value.


### GetSeverity

`func (o *ModelCloudCompliance) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ModelCloudCompliance) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ModelCloudCompliance) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetStatus

`func (o *ModelCloudCompliance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelCloudCompliance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelCloudCompliance) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *ModelCloudCompliance) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModelCloudCompliance) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModelCloudCompliance) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *ModelCloudCompliance) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelCloudCompliance) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelCloudCompliance) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *ModelCloudCompliance) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelCloudCompliance) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelCloudCompliance) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


