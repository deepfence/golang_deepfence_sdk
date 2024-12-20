# ModelKubernetesCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentRunning** | **bool** |  | 
**Hosts** | [**[]ModelHost**](ModelHost.md) |  | 
**NodeId** | **string** |  | 
**NodeName** | **string** |  | 
**Tags** | **[]string** |  | 

## Methods

### NewModelKubernetesCluster

`func NewModelKubernetesCluster(agentRunning bool, hosts []ModelHost, nodeId string, nodeName string, tags []string, ) *ModelKubernetesCluster`

NewModelKubernetesCluster instantiates a new ModelKubernetesCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelKubernetesClusterWithDefaults

`func NewModelKubernetesClusterWithDefaults() *ModelKubernetesCluster`

NewModelKubernetesClusterWithDefaults instantiates a new ModelKubernetesCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentRunning

`func (o *ModelKubernetesCluster) GetAgentRunning() bool`

GetAgentRunning returns the AgentRunning field if non-nil, zero value otherwise.

### GetAgentRunningOk

`func (o *ModelKubernetesCluster) GetAgentRunningOk() (*bool, bool)`

GetAgentRunningOk returns a tuple with the AgentRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentRunning

`func (o *ModelKubernetesCluster) SetAgentRunning(v bool)`

SetAgentRunning sets AgentRunning field to given value.


### GetHosts

`func (o *ModelKubernetesCluster) GetHosts() []ModelHost`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *ModelKubernetesCluster) GetHostsOk() (*[]ModelHost, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *ModelKubernetesCluster) SetHosts(v []ModelHost)`

SetHosts sets Hosts field to given value.


### SetHostsNil

`func (o *ModelKubernetesCluster) SetHostsNil(b bool)`

 SetHostsNil sets the value for Hosts to be an explicit nil

### UnsetHosts
`func (o *ModelKubernetesCluster) UnsetHosts()`

UnsetHosts ensures that no value is present for Hosts, not even an explicit nil
### GetNodeId

`func (o *ModelKubernetesCluster) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelKubernetesCluster) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelKubernetesCluster) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNodeName

`func (o *ModelKubernetesCluster) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *ModelKubernetesCluster) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *ModelKubernetesCluster) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.


### GetTags

`func (o *ModelKubernetesCluster) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModelKubernetesCluster) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModelKubernetesCluster) SetTags(v []string)`

SetTags sets Tags field to given value.


### SetTagsNil

`func (o *ModelKubernetesCluster) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ModelKubernetesCluster) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


