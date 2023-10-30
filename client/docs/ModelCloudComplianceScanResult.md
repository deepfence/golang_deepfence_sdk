# ModelCloudComplianceScanResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BenchmarkType** | **[]string** |  | 
**CloudAccountId** | **string** |  | 
**CompliancePercentage** | **float32** |  | 
**Compliances** | [**[]ModelCloudCompliance**](ModelCloudCompliance.md) |  | 
**CreatedAt** | **int64** |  | 
**DockerContainerName** | **string** |  | 
**DockerImageName** | **string** |  | 
**HostName** | **string** |  | 
**KubernetesClusterName** | **string** |  | 
**NodeId** | **string** |  | 
**NodeName** | **string** |  | 
**NodeType** | **string** |  | 
**ScanId** | **string** |  | 
**StatusCounts** | **map[string]int32** |  | 
**UpdatedAt** | **int64** |  | 

## Methods

### NewModelCloudComplianceScanResult

`func NewModelCloudComplianceScanResult(benchmarkType []string, cloudAccountId string, compliancePercentage float32, compliances []ModelCloudCompliance, createdAt int64, dockerContainerName string, dockerImageName string, hostName string, kubernetesClusterName string, nodeId string, nodeName string, nodeType string, scanId string, statusCounts map[string]int32, updatedAt int64, ) *ModelCloudComplianceScanResult`

NewModelCloudComplianceScanResult instantiates a new ModelCloudComplianceScanResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCloudComplianceScanResultWithDefaults

`func NewModelCloudComplianceScanResultWithDefaults() *ModelCloudComplianceScanResult`

NewModelCloudComplianceScanResultWithDefaults instantiates a new ModelCloudComplianceScanResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBenchmarkType

`func (o *ModelCloudComplianceScanResult) GetBenchmarkType() []string`

GetBenchmarkType returns the BenchmarkType field if non-nil, zero value otherwise.

### GetBenchmarkTypeOk

`func (o *ModelCloudComplianceScanResult) GetBenchmarkTypeOk() (*[]string, bool)`

GetBenchmarkTypeOk returns a tuple with the BenchmarkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmarkType

`func (o *ModelCloudComplianceScanResult) SetBenchmarkType(v []string)`

SetBenchmarkType sets BenchmarkType field to given value.


### SetBenchmarkTypeNil

`func (o *ModelCloudComplianceScanResult) SetBenchmarkTypeNil(b bool)`

 SetBenchmarkTypeNil sets the value for BenchmarkType to be an explicit nil

### UnsetBenchmarkType
`func (o *ModelCloudComplianceScanResult) UnsetBenchmarkType()`

UnsetBenchmarkType ensures that no value is present for BenchmarkType, not even an explicit nil
### GetCloudAccountId

`func (o *ModelCloudComplianceScanResult) GetCloudAccountId() string`

GetCloudAccountId returns the CloudAccountId field if non-nil, zero value otherwise.

### GetCloudAccountIdOk

`func (o *ModelCloudComplianceScanResult) GetCloudAccountIdOk() (*string, bool)`

GetCloudAccountIdOk returns a tuple with the CloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccountId

`func (o *ModelCloudComplianceScanResult) SetCloudAccountId(v string)`

SetCloudAccountId sets CloudAccountId field to given value.


### GetCompliancePercentage

`func (o *ModelCloudComplianceScanResult) GetCompliancePercentage() float32`

GetCompliancePercentage returns the CompliancePercentage field if non-nil, zero value otherwise.

### GetCompliancePercentageOk

`func (o *ModelCloudComplianceScanResult) GetCompliancePercentageOk() (*float32, bool)`

GetCompliancePercentageOk returns a tuple with the CompliancePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliancePercentage

`func (o *ModelCloudComplianceScanResult) SetCompliancePercentage(v float32)`

SetCompliancePercentage sets CompliancePercentage field to given value.


### GetCompliances

`func (o *ModelCloudComplianceScanResult) GetCompliances() []ModelCloudCompliance`

GetCompliances returns the Compliances field if non-nil, zero value otherwise.

### GetCompliancesOk

`func (o *ModelCloudComplianceScanResult) GetCompliancesOk() (*[]ModelCloudCompliance, bool)`

GetCompliancesOk returns a tuple with the Compliances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliances

`func (o *ModelCloudComplianceScanResult) SetCompliances(v []ModelCloudCompliance)`

SetCompliances sets Compliances field to given value.


### SetCompliancesNil

`func (o *ModelCloudComplianceScanResult) SetCompliancesNil(b bool)`

 SetCompliancesNil sets the value for Compliances to be an explicit nil

### UnsetCompliances
`func (o *ModelCloudComplianceScanResult) UnsetCompliances()`

UnsetCompliances ensures that no value is present for Compliances, not even an explicit nil
### GetCreatedAt

`func (o *ModelCloudComplianceScanResult) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModelCloudComplianceScanResult) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModelCloudComplianceScanResult) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetDockerContainerName

`func (o *ModelCloudComplianceScanResult) GetDockerContainerName() string`

GetDockerContainerName returns the DockerContainerName field if non-nil, zero value otherwise.

### GetDockerContainerNameOk

`func (o *ModelCloudComplianceScanResult) GetDockerContainerNameOk() (*string, bool)`

GetDockerContainerNameOk returns a tuple with the DockerContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerContainerName

`func (o *ModelCloudComplianceScanResult) SetDockerContainerName(v string)`

SetDockerContainerName sets DockerContainerName field to given value.


### GetDockerImageName

`func (o *ModelCloudComplianceScanResult) GetDockerImageName() string`

GetDockerImageName returns the DockerImageName field if non-nil, zero value otherwise.

### GetDockerImageNameOk

`func (o *ModelCloudComplianceScanResult) GetDockerImageNameOk() (*string, bool)`

GetDockerImageNameOk returns a tuple with the DockerImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageName

`func (o *ModelCloudComplianceScanResult) SetDockerImageName(v string)`

SetDockerImageName sets DockerImageName field to given value.


### GetHostName

`func (o *ModelCloudComplianceScanResult) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ModelCloudComplianceScanResult) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ModelCloudComplianceScanResult) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetKubernetesClusterName

`func (o *ModelCloudComplianceScanResult) GetKubernetesClusterName() string`

GetKubernetesClusterName returns the KubernetesClusterName field if non-nil, zero value otherwise.

### GetKubernetesClusterNameOk

`func (o *ModelCloudComplianceScanResult) GetKubernetesClusterNameOk() (*string, bool)`

GetKubernetesClusterNameOk returns a tuple with the KubernetesClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesClusterName

`func (o *ModelCloudComplianceScanResult) SetKubernetesClusterName(v string)`

SetKubernetesClusterName sets KubernetesClusterName field to given value.


### GetNodeId

`func (o *ModelCloudComplianceScanResult) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelCloudComplianceScanResult) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelCloudComplianceScanResult) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNodeName

`func (o *ModelCloudComplianceScanResult) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *ModelCloudComplianceScanResult) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *ModelCloudComplianceScanResult) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.


### GetNodeType

`func (o *ModelCloudComplianceScanResult) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *ModelCloudComplianceScanResult) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *ModelCloudComplianceScanResult) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetScanId

`func (o *ModelCloudComplianceScanResult) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *ModelCloudComplianceScanResult) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *ModelCloudComplianceScanResult) SetScanId(v string)`

SetScanId sets ScanId field to given value.


### GetStatusCounts

`func (o *ModelCloudComplianceScanResult) GetStatusCounts() map[string]int32`

GetStatusCounts returns the StatusCounts field if non-nil, zero value otherwise.

### GetStatusCountsOk

`func (o *ModelCloudComplianceScanResult) GetStatusCountsOk() (*map[string]int32, bool)`

GetStatusCountsOk returns a tuple with the StatusCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCounts

`func (o *ModelCloudComplianceScanResult) SetStatusCounts(v map[string]int32)`

SetStatusCounts sets StatusCounts field to given value.


### SetStatusCountsNil

`func (o *ModelCloudComplianceScanResult) SetStatusCountsNil(b bool)`

 SetStatusCountsNil sets the value for StatusCounts to be an explicit nil

### UnsetStatusCounts
`func (o *ModelCloudComplianceScanResult) UnsetStatusCounts()`

UnsetStatusCounts ensures that no value is present for StatusCounts, not even an explicit nil
### GetUpdatedAt

`func (o *ModelCloudComplianceScanResult) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModelCloudComplianceScanResult) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModelCloudComplianceScanResult) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


