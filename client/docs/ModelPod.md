# ModelPod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | [**[]ModelContainer**](ModelContainer.md) |  | 
**HostName** | **string** |  | 
**KubernetesClusterId** | **string** |  | 
**KubernetesClusterName** | **string** |  | 
**KubernetesCreated** | **string** |  | 
**KubernetesIp** | **string** |  | 
**KubernetesIsInHostNetwork** | **bool** |  | 
**KubernetesLabels** | **map[string]interface{}** |  | 
**KubernetesNamespace** | **string** |  | 
**KubernetesState** | **string** |  | 
**MalwareScanStatus** | **string** |  | 
**NodeId** | **string** |  | 
**NodeName** | **string** |  | 
**PodName** | **string** |  | 
**Processes** | [**[]ModelProcess**](ModelProcess.md) |  | 
**SecretScanStatus** | **string** |  | 
**VulnerabilityScanStatus** | **string** |  | 

## Methods

### NewModelPod

`func NewModelPod(containers []ModelContainer, hostName string, kubernetesClusterId string, kubernetesClusterName string, kubernetesCreated string, kubernetesIp string, kubernetesIsInHostNetwork bool, kubernetesLabels map[string]interface{}, kubernetesNamespace string, kubernetesState string, malwareScanStatus string, nodeId string, nodeName string, podName string, processes []ModelProcess, secretScanStatus string, vulnerabilityScanStatus string, ) *ModelPod`

NewModelPod instantiates a new ModelPod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPodWithDefaults

`func NewModelPodWithDefaults() *ModelPod`

NewModelPodWithDefaults instantiates a new ModelPod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainers

`func (o *ModelPod) GetContainers() []ModelContainer`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *ModelPod) GetContainersOk() (*[]ModelContainer, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *ModelPod) SetContainers(v []ModelContainer)`

SetContainers sets Containers field to given value.


### SetContainersNil

`func (o *ModelPod) SetContainersNil(b bool)`

 SetContainersNil sets the value for Containers to be an explicit nil

### UnsetContainers
`func (o *ModelPod) UnsetContainers()`

UnsetContainers ensures that no value is present for Containers, not even an explicit nil
### GetHostName

`func (o *ModelPod) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ModelPod) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ModelPod) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetKubernetesClusterId

`func (o *ModelPod) GetKubernetesClusterId() string`

GetKubernetesClusterId returns the KubernetesClusterId field if non-nil, zero value otherwise.

### GetKubernetesClusterIdOk

`func (o *ModelPod) GetKubernetesClusterIdOk() (*string, bool)`

GetKubernetesClusterIdOk returns a tuple with the KubernetesClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesClusterId

`func (o *ModelPod) SetKubernetesClusterId(v string)`

SetKubernetesClusterId sets KubernetesClusterId field to given value.


### GetKubernetesClusterName

`func (o *ModelPod) GetKubernetesClusterName() string`

GetKubernetesClusterName returns the KubernetesClusterName field if non-nil, zero value otherwise.

### GetKubernetesClusterNameOk

`func (o *ModelPod) GetKubernetesClusterNameOk() (*string, bool)`

GetKubernetesClusterNameOk returns a tuple with the KubernetesClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesClusterName

`func (o *ModelPod) SetKubernetesClusterName(v string)`

SetKubernetesClusterName sets KubernetesClusterName field to given value.


### GetKubernetesCreated

`func (o *ModelPod) GetKubernetesCreated() string`

GetKubernetesCreated returns the KubernetesCreated field if non-nil, zero value otherwise.

### GetKubernetesCreatedOk

`func (o *ModelPod) GetKubernetesCreatedOk() (*string, bool)`

GetKubernetesCreatedOk returns a tuple with the KubernetesCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesCreated

`func (o *ModelPod) SetKubernetesCreated(v string)`

SetKubernetesCreated sets KubernetesCreated field to given value.


### GetKubernetesIp

`func (o *ModelPod) GetKubernetesIp() string`

GetKubernetesIp returns the KubernetesIp field if non-nil, zero value otherwise.

### GetKubernetesIpOk

`func (o *ModelPod) GetKubernetesIpOk() (*string, bool)`

GetKubernetesIpOk returns a tuple with the KubernetesIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesIp

`func (o *ModelPod) SetKubernetesIp(v string)`

SetKubernetesIp sets KubernetesIp field to given value.


### GetKubernetesIsInHostNetwork

`func (o *ModelPod) GetKubernetesIsInHostNetwork() bool`

GetKubernetesIsInHostNetwork returns the KubernetesIsInHostNetwork field if non-nil, zero value otherwise.

### GetKubernetesIsInHostNetworkOk

`func (o *ModelPod) GetKubernetesIsInHostNetworkOk() (*bool, bool)`

GetKubernetesIsInHostNetworkOk returns a tuple with the KubernetesIsInHostNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesIsInHostNetwork

`func (o *ModelPod) SetKubernetesIsInHostNetwork(v bool)`

SetKubernetesIsInHostNetwork sets KubernetesIsInHostNetwork field to given value.


### GetKubernetesLabels

`func (o *ModelPod) GetKubernetesLabels() map[string]interface{}`

GetKubernetesLabels returns the KubernetesLabels field if non-nil, zero value otherwise.

### GetKubernetesLabelsOk

`func (o *ModelPod) GetKubernetesLabelsOk() (*map[string]interface{}, bool)`

GetKubernetesLabelsOk returns a tuple with the KubernetesLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesLabels

`func (o *ModelPod) SetKubernetesLabels(v map[string]interface{})`

SetKubernetesLabels sets KubernetesLabels field to given value.


### SetKubernetesLabelsNil

`func (o *ModelPod) SetKubernetesLabelsNil(b bool)`

 SetKubernetesLabelsNil sets the value for KubernetesLabels to be an explicit nil

### UnsetKubernetesLabels
`func (o *ModelPod) UnsetKubernetesLabels()`

UnsetKubernetesLabels ensures that no value is present for KubernetesLabels, not even an explicit nil
### GetKubernetesNamespace

`func (o *ModelPod) GetKubernetesNamespace() string`

GetKubernetesNamespace returns the KubernetesNamespace field if non-nil, zero value otherwise.

### GetKubernetesNamespaceOk

`func (o *ModelPod) GetKubernetesNamespaceOk() (*string, bool)`

GetKubernetesNamespaceOk returns a tuple with the KubernetesNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesNamespace

`func (o *ModelPod) SetKubernetesNamespace(v string)`

SetKubernetesNamespace sets KubernetesNamespace field to given value.


### GetKubernetesState

`func (o *ModelPod) GetKubernetesState() string`

GetKubernetesState returns the KubernetesState field if non-nil, zero value otherwise.

### GetKubernetesStateOk

`func (o *ModelPod) GetKubernetesStateOk() (*string, bool)`

GetKubernetesStateOk returns a tuple with the KubernetesState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesState

`func (o *ModelPod) SetKubernetesState(v string)`

SetKubernetesState sets KubernetesState field to given value.


### GetMalwareScanStatus

`func (o *ModelPod) GetMalwareScanStatus() string`

GetMalwareScanStatus returns the MalwareScanStatus field if non-nil, zero value otherwise.

### GetMalwareScanStatusOk

`func (o *ModelPod) GetMalwareScanStatusOk() (*string, bool)`

GetMalwareScanStatusOk returns a tuple with the MalwareScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalwareScanStatus

`func (o *ModelPod) SetMalwareScanStatus(v string)`

SetMalwareScanStatus sets MalwareScanStatus field to given value.


### GetNodeId

`func (o *ModelPod) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelPod) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelPod) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNodeName

`func (o *ModelPod) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *ModelPod) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *ModelPod) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.


### GetPodName

`func (o *ModelPod) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *ModelPod) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *ModelPod) SetPodName(v string)`

SetPodName sets PodName field to given value.


### GetProcesses

`func (o *ModelPod) GetProcesses() []ModelProcess`

GetProcesses returns the Processes field if non-nil, zero value otherwise.

### GetProcessesOk

`func (o *ModelPod) GetProcessesOk() (*[]ModelProcess, bool)`

GetProcessesOk returns a tuple with the Processes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcesses

`func (o *ModelPod) SetProcesses(v []ModelProcess)`

SetProcesses sets Processes field to given value.


### SetProcessesNil

`func (o *ModelPod) SetProcessesNil(b bool)`

 SetProcessesNil sets the value for Processes to be an explicit nil

### UnsetProcesses
`func (o *ModelPod) UnsetProcesses()`

UnsetProcesses ensures that no value is present for Processes, not even an explicit nil
### GetSecretScanStatus

`func (o *ModelPod) GetSecretScanStatus() string`

GetSecretScanStatus returns the SecretScanStatus field if non-nil, zero value otherwise.

### GetSecretScanStatusOk

`func (o *ModelPod) GetSecretScanStatusOk() (*string, bool)`

GetSecretScanStatusOk returns a tuple with the SecretScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretScanStatus

`func (o *ModelPod) SetSecretScanStatus(v string)`

SetSecretScanStatus sets SecretScanStatus field to given value.


### GetVulnerabilityScanStatus

`func (o *ModelPod) GetVulnerabilityScanStatus() string`

GetVulnerabilityScanStatus returns the VulnerabilityScanStatus field if non-nil, zero value otherwise.

### GetVulnerabilityScanStatusOk

`func (o *ModelPod) GetVulnerabilityScanStatusOk() (*string, bool)`

GetVulnerabilityScanStatusOk returns a tuple with the VulnerabilityScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityScanStatus

`func (o *ModelPod) SetVulnerabilityScanStatus(v string)`

SetVulnerabilityScanStatus sets VulnerabilityScanStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


