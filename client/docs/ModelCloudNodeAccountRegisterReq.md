# ModelCloudNodeAccountRegisterReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**CloudProvider** | **string** |  | 
**HostNodeId** | **string** |  | 
**IsOrganizationDeployment** | Pointer to **bool** |  | [optional] 
**MonitoredAccountIds** | Pointer to **map[string]string** |  | [optional] 
**NodeId** | **string** |  | 
**OrganizationAccountId** | Pointer to **string** |  | [optional] 
**Version** | **string** |  | 

## Methods

### NewModelCloudNodeAccountRegisterReq

`func NewModelCloudNodeAccountRegisterReq(accountId string, cloudProvider string, hostNodeId string, nodeId string, version string, ) *ModelCloudNodeAccountRegisterReq`

NewModelCloudNodeAccountRegisterReq instantiates a new ModelCloudNodeAccountRegisterReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCloudNodeAccountRegisterReqWithDefaults

`func NewModelCloudNodeAccountRegisterReqWithDefaults() *ModelCloudNodeAccountRegisterReq`

NewModelCloudNodeAccountRegisterReqWithDefaults instantiates a new ModelCloudNodeAccountRegisterReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ModelCloudNodeAccountRegisterReq) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ModelCloudNodeAccountRegisterReq) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ModelCloudNodeAccountRegisterReq) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCloudProvider

`func (o *ModelCloudNodeAccountRegisterReq) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ModelCloudNodeAccountRegisterReq) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ModelCloudNodeAccountRegisterReq) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetHostNodeId

`func (o *ModelCloudNodeAccountRegisterReq) GetHostNodeId() string`

GetHostNodeId returns the HostNodeId field if non-nil, zero value otherwise.

### GetHostNodeIdOk

`func (o *ModelCloudNodeAccountRegisterReq) GetHostNodeIdOk() (*string, bool)`

GetHostNodeIdOk returns a tuple with the HostNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNodeId

`func (o *ModelCloudNodeAccountRegisterReq) SetHostNodeId(v string)`

SetHostNodeId sets HostNodeId field to given value.


### GetIsOrganizationDeployment

`func (o *ModelCloudNodeAccountRegisterReq) GetIsOrganizationDeployment() bool`

GetIsOrganizationDeployment returns the IsOrganizationDeployment field if non-nil, zero value otherwise.

### GetIsOrganizationDeploymentOk

`func (o *ModelCloudNodeAccountRegisterReq) GetIsOrganizationDeploymentOk() (*bool, bool)`

GetIsOrganizationDeploymentOk returns a tuple with the IsOrganizationDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrganizationDeployment

`func (o *ModelCloudNodeAccountRegisterReq) SetIsOrganizationDeployment(v bool)`

SetIsOrganizationDeployment sets IsOrganizationDeployment field to given value.

### HasIsOrganizationDeployment

`func (o *ModelCloudNodeAccountRegisterReq) HasIsOrganizationDeployment() bool`

HasIsOrganizationDeployment returns a boolean if a field has been set.

### GetMonitoredAccountIds

`func (o *ModelCloudNodeAccountRegisterReq) GetMonitoredAccountIds() map[string]string`

GetMonitoredAccountIds returns the MonitoredAccountIds field if non-nil, zero value otherwise.

### GetMonitoredAccountIdsOk

`func (o *ModelCloudNodeAccountRegisterReq) GetMonitoredAccountIdsOk() (*map[string]string, bool)`

GetMonitoredAccountIdsOk returns a tuple with the MonitoredAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredAccountIds

`func (o *ModelCloudNodeAccountRegisterReq) SetMonitoredAccountIds(v map[string]string)`

SetMonitoredAccountIds sets MonitoredAccountIds field to given value.

### HasMonitoredAccountIds

`func (o *ModelCloudNodeAccountRegisterReq) HasMonitoredAccountIds() bool`

HasMonitoredAccountIds returns a boolean if a field has been set.

### SetMonitoredAccountIdsNil

`func (o *ModelCloudNodeAccountRegisterReq) SetMonitoredAccountIdsNil(b bool)`

 SetMonitoredAccountIdsNil sets the value for MonitoredAccountIds to be an explicit nil

### UnsetMonitoredAccountIds
`func (o *ModelCloudNodeAccountRegisterReq) UnsetMonitoredAccountIds()`

UnsetMonitoredAccountIds ensures that no value is present for MonitoredAccountIds, not even an explicit nil
### GetNodeId

`func (o *ModelCloudNodeAccountRegisterReq) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelCloudNodeAccountRegisterReq) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelCloudNodeAccountRegisterReq) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetOrganizationAccountId

`func (o *ModelCloudNodeAccountRegisterReq) GetOrganizationAccountId() string`

GetOrganizationAccountId returns the OrganizationAccountId field if non-nil, zero value otherwise.

### GetOrganizationAccountIdOk

`func (o *ModelCloudNodeAccountRegisterReq) GetOrganizationAccountIdOk() (*string, bool)`

GetOrganizationAccountIdOk returns a tuple with the OrganizationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationAccountId

`func (o *ModelCloudNodeAccountRegisterReq) SetOrganizationAccountId(v string)`

SetOrganizationAccountId sets OrganizationAccountId field to given value.

### HasOrganizationAccountId

`func (o *ModelCloudNodeAccountRegisterReq) HasOrganizationAccountId() bool`

HasOrganizationAccountId returns a boolean if a field has been set.

### GetVersion

`func (o *ModelCloudNodeAccountRegisterReq) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelCloudNodeAccountRegisterReq) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelCloudNodeAccountRegisterReq) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


