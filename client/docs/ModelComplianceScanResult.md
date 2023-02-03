# ModelComplianceScanResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compliances** | [**[]ModelCompliance**](ModelCompliance.md) |  | 
**DockerContainerName** | **string** |  | 
**DockerImageName** | **string** |  | 
**HostName** | **string** |  | 
**KubernetesClusterName** | **string** |  | 
**NodeId** | **string** |  | 
**NodeName** | **string** |  | 
**NodeType** | **string** |  | 
**ScanId** | **string** |  | 

## Methods

### NewModelComplianceScanResult

`func NewModelComplianceScanResult(compliances []ModelCompliance, dockerContainerName string, dockerImageName string, hostName string, kubernetesClusterName string, nodeId string, nodeName string, nodeType string, scanId string, ) *ModelComplianceScanResult`

NewModelComplianceScanResult instantiates a new ModelComplianceScanResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelComplianceScanResultWithDefaults

`func NewModelComplianceScanResultWithDefaults() *ModelComplianceScanResult`

NewModelComplianceScanResultWithDefaults instantiates a new ModelComplianceScanResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompliances

`func (o *ModelComplianceScanResult) GetCompliances() []ModelCompliance`

GetCompliances returns the Compliances field if non-nil, zero value otherwise.

### GetCompliancesOk

`func (o *ModelComplianceScanResult) GetCompliancesOk() (*[]ModelCompliance, bool)`

GetCompliancesOk returns a tuple with the Compliances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliances

`func (o *ModelComplianceScanResult) SetCompliances(v []ModelCompliance)`

SetCompliances sets Compliances field to given value.


### SetCompliancesNil

`func (o *ModelComplianceScanResult) SetCompliancesNil(b bool)`

 SetCompliancesNil sets the value for Compliances to be an explicit nil

### UnsetCompliances
`func (o *ModelComplianceScanResult) UnsetCompliances()`

UnsetCompliances ensures that no value is present for Compliances, not even an explicit nil
### GetDockerContainerName

`func (o *ModelComplianceScanResult) GetDockerContainerName() string`

GetDockerContainerName returns the DockerContainerName field if non-nil, zero value otherwise.

### GetDockerContainerNameOk

`func (o *ModelComplianceScanResult) GetDockerContainerNameOk() (*string, bool)`

GetDockerContainerNameOk returns a tuple with the DockerContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContainerName

`func (o *ModelComplianceScanResult) SetDockerContainerName(v string)`

SetDockerContainerName sets DockerContainerName field to given value.


### GetDockerImageName

`func (o *ModelComplianceScanResult) GetDockerImageName() string`

GetDockerImageName returns the DockerImageName field if non-nil, zero value otherwise.

### GetDockerImageNameOk

`func (o *ModelComplianceScanResult) GetDockerImageNameOk() (*string, bool)`

GetDockerImageNameOk returns a tuple with the DockerImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageName

`func (o *ModelComplianceScanResult) SetDockerImageName(v string)`

SetDockerImageName sets DockerImageName field to given value.


### GetHostName

`func (o *ModelComplianceScanResult) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ModelComplianceScanResult) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ModelComplianceScanResult) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetKubernetesClusterName

`func (o *ModelComplianceScanResult) GetKubernetesClusterName() string`

GetKubernetesClusterName returns the KubernetesClusterName field if non-nil, zero value otherwise.

### GetKubernetesClusterNameOk

`func (o *ModelComplianceScanResult) GetKubernetesClusterNameOk() (*string, bool)`

GetKubernetesClusterNameOk returns a tuple with the KubernetesClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesClusterName

`func (o *ModelComplianceScanResult) SetKubernetesClusterName(v string)`

SetKubernetesClusterName sets KubernetesClusterName field to given value.


### GetNodeId

`func (o *ModelComplianceScanResult) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelComplianceScanResult) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelComplianceScanResult) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNodeName

`func (o *ModelComplianceScanResult) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *ModelComplianceScanResult) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *ModelComplianceScanResult) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.


### GetNodeType

`func (o *ModelComplianceScanResult) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *ModelComplianceScanResult) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *ModelComplianceScanResult) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetScanId

`func (o *ModelComplianceScanResult) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *ModelComplianceScanResult) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *ModelComplianceScanResult) SetScanId(v string)`

SetScanId sets ScanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


