# ModelPod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | [**[]ModelContainer**](ModelContainer.md) |  | 
**HostNodeId** | **string** |  | 
**KubernetesName** | **string** |  | 
**KubernetesNamespace** | **string** |  | 
**Metadata** | **map[string]interface{}** |  | 
**Metrics** | [**ModelComputeMetrics**](ModelComputeMetrics.md) |  | 
**NodeId** | **string** |  | 
**Processes** | [**[]ModelProcess**](ModelProcess.md) |  | 

## Methods

### NewModelPod

`func NewModelPod(containers []ModelContainer, hostNodeId string, kubernetesName string, kubernetesNamespace string, metadata map[string]interface{}, metrics ModelComputeMetrics, nodeId string, processes []ModelProcess, ) *ModelPod`

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
### GetHostNodeId

`func (o *ModelPod) GetHostNodeId() string`

GetHostNodeId returns the HostNodeId field if non-nil, zero value otherwise.

### GetHostNodeIdOk

`func (o *ModelPod) GetHostNodeIdOk() (*string, bool)`

GetHostNodeIdOk returns a tuple with the HostNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNodeId

`func (o *ModelPod) SetHostNodeId(v string)`

SetHostNodeId sets HostNodeId field to given value.


### GetKubernetesName

`func (o *ModelPod) GetKubernetesName() string`

GetKubernetesName returns the KubernetesName field if non-nil, zero value otherwise.

### GetKubernetesNameOk

`func (o *ModelPod) GetKubernetesNameOk() (*string, bool)`

GetKubernetesNameOk returns a tuple with the KubernetesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesName

`func (o *ModelPod) SetKubernetesName(v string)`

SetKubernetesName sets KubernetesName field to given value.


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


### GetMetadata

`func (o *ModelPod) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ModelPod) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ModelPod) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetMetrics

`func (o *ModelPod) GetMetrics() ModelComputeMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ModelPod) GetMetricsOk() (*ModelComputeMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ModelPod) SetMetrics(v ModelComputeMetrics)`

SetMetrics sets Metrics field to given value.


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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


