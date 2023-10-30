# ModelSecretScanResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudAccountId** | **string** |  | 
**CreatedAt** | **int64** |  | 
**DockerContainerName** | **string** |  | 
**DockerImageName** | **string** |  | 
**HostName** | **string** |  | 
**KubernetesClusterName** | **string** |  | 
**NodeId** | **string** |  | 
**NodeName** | **string** |  | 
**NodeType** | **string** |  | 
**ScanId** | **string** |  | 
**Secrets** | [**[]ModelSecret**](ModelSecret.md) |  | 
**SeverityCounts** | **map[string]int32** |  | 
**UpdatedAt** | **int64** |  | 

## Methods

### NewModelSecretScanResult

`func NewModelSecretScanResult(cloudAccountId string, createdAt int64, dockerContainerName string, dockerImageName string, hostName string, kubernetesClusterName string, nodeId string, nodeName string, nodeType string, scanId string, secrets []ModelSecret, severityCounts map[string]int32, updatedAt int64, ) *ModelSecretScanResult`

NewModelSecretScanResult instantiates a new ModelSecretScanResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSecretScanResultWithDefaults

`func NewModelSecretScanResultWithDefaults() *ModelSecretScanResult`

NewModelSecretScanResultWithDefaults instantiates a new ModelSecretScanResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudAccountId

`func (o *ModelSecretScanResult) GetCloudAccountId() string`

GetCloudAccountId returns the CloudAccountId field if non-nil, zero value otherwise.

### GetCloudAccountIdOk

`func (o *ModelSecretScanResult) GetCloudAccountIdOk() (*string, bool)`

GetCloudAccountIdOk returns a tuple with the CloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccountId

`func (o *ModelSecretScanResult) SetCloudAccountId(v string)`

SetCloudAccountId sets CloudAccountId field to given value.


### GetCreatedAt

`func (o *ModelSecretScanResult) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelSecretScanResult) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelSecretScanResult) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetDockerContainerName

`func (o *ModelSecretScanResult) GetDockerContainerName() string`

GetDockerContainerName returns the DockerContainerName field if non-nil, zero value otherwise.

### GetDockerContainerNameOk

`func (o *ModelSecretScanResult) GetDockerContainerNameOk() (*string, bool)`

GetDockerContainerNameOk returns a tuple with the DockerContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContainerName

`func (o *ModelSecretScanResult) SetDockerContainerName(v string)`

SetDockerContainerName sets DockerContainerName field to given value.


### GetDockerImageName

`func (o *ModelSecretScanResult) GetDockerImageName() string`

GetDockerImageName returns the DockerImageName field if non-nil, zero value otherwise.

### GetDockerImageNameOk

`func (o *ModelSecretScanResult) GetDockerImageNameOk() (*string, bool)`

GetDockerImageNameOk returns a tuple with the DockerImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageName

`func (o *ModelSecretScanResult) SetDockerImageName(v string)`

SetDockerImageName sets DockerImageName field to given value.


### GetHostName

`func (o *ModelSecretScanResult) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ModelSecretScanResult) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ModelSecretScanResult) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetKubernetesClusterName

`func (o *ModelSecretScanResult) GetKubernetesClusterName() string`

GetKubernetesClusterName returns the KubernetesClusterName field if non-nil, zero value otherwise.

### GetKubernetesClusterNameOk

`func (o *ModelSecretScanResult) GetKubernetesClusterNameOk() (*string, bool)`

GetKubernetesClusterNameOk returns a tuple with the KubernetesClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesClusterName

`func (o *ModelSecretScanResult) SetKubernetesClusterName(v string)`

SetKubernetesClusterName sets KubernetesClusterName field to given value.


### GetNodeId

`func (o *ModelSecretScanResult) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelSecretScanResult) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelSecretScanResult) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNodeName

`func (o *ModelSecretScanResult) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *ModelSecretScanResult) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *ModelSecretScanResult) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.


### GetNodeType

`func (o *ModelSecretScanResult) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *ModelSecretScanResult) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *ModelSecretScanResult) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetScanId

`func (o *ModelSecretScanResult) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *ModelSecretScanResult) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *ModelSecretScanResult) SetScanId(v string)`

SetScanId sets ScanId field to given value.


### GetSecrets

`func (o *ModelSecretScanResult) GetSecrets() []ModelSecret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *ModelSecretScanResult) GetSecretsOk() (*[]ModelSecret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *ModelSecretScanResult) SetSecrets(v []ModelSecret)`

SetSecrets sets Secrets field to given value.


### SetSecretsNil

`func (o *ModelSecretScanResult) SetSecretsNil(b bool)`

 SetSecretsNil sets the value for Secrets to be an explicit nil

### UnsetSecrets
`func (o *ModelSecretScanResult) UnsetSecrets()`

UnsetSecrets ensures that no value is present for Secrets, not even an explicit nil
### GetSeverityCounts

`func (o *ModelSecretScanResult) GetSeverityCounts() map[string]int32`

GetSeverityCounts returns the SeverityCounts field if non-nil, zero value otherwise.

### GetSeverityCountsOk

`func (o *ModelSecretScanResult) GetSeverityCountsOk() (*map[string]int32, bool)`

GetSeverityCountsOk returns a tuple with the SeverityCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityCounts

`func (o *ModelSecretScanResult) SetSeverityCounts(v map[string]int32)`

SetSeverityCounts sets SeverityCounts field to given value.


### SetSeverityCountsNil

`func (o *ModelSecretScanResult) SetSeverityCountsNil(b bool)`

 SetSeverityCountsNil sets the value for SeverityCounts to be an explicit nil

### UnsetSeverityCounts
`func (o *ModelSecretScanResult) UnsetSeverityCounts()`

UnsetSeverityCounts ensures that no value is present for SeverityCounts, not even an explicit nil
### GetUpdatedAt

`func (o *ModelSecretScanResult) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelSecretScanResult) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelSecretScanResult) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


