# ModelSecretScanResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DockerContainerName** | **string** |  | 
**DockerImageName** | **string** |  | 
**HostName** | **string** |  | 
**KubernetesClusterName** | **string** |  | 
**NodeId** | **string** |  | 
**NodeName** | **string** |  | 
**NodeType** | **string** |  | 
**Rule2Secrets** | **map[string][]int32** |  | 
**Rules** | [**[]ModelRule**](ModelRule.md) |  | 
**ScanId** | **string** |  | 
**Secrets** | [**[]ModelSecret**](ModelSecret.md) |  | 
**SeverityCounts** | **map[string]int32** |  | 

## Methods

### NewModelSecretScanResult

`func NewModelSecretScanResult(dockerContainerName string, dockerImageName string, hostName string, kubernetesClusterName string, nodeId string, nodeName string, nodeType string, rule2Secrets map[string][]int32, rules []ModelRule, scanId string, secrets []ModelSecret, severityCounts map[string]int32, ) *ModelSecretScanResult`

NewModelSecretScanResult instantiates a new ModelSecretScanResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSecretScanResultWithDefaults

`func NewModelSecretScanResultWithDefaults() *ModelSecretScanResult`

NewModelSecretScanResultWithDefaults instantiates a new ModelSecretScanResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetRule2Secrets

`func (o *ModelSecretScanResult) GetRule2Secrets() map[string][]int32`

GetRule2Secrets returns the Rule2Secrets field if non-nil, zero value otherwise.

### GetRule2SecretsOk

`func (o *ModelSecretScanResult) GetRule2SecretsOk() (*map[string][]int32, bool)`

GetRule2SecretsOk returns a tuple with the Rule2Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule2Secrets

`func (o *ModelSecretScanResult) SetRule2Secrets(v map[string][]int32)`

SetRule2Secrets sets Rule2Secrets field to given value.


### SetRule2SecretsNil

`func (o *ModelSecretScanResult) SetRule2SecretsNil(b bool)`

 SetRule2SecretsNil sets the value for Rule2Secrets to be an explicit nil

### UnsetRule2Secrets
`func (o *ModelSecretScanResult) UnsetRule2Secrets()`

UnsetRule2Secrets ensures that no value is present for Rule2Secrets, not even an explicit nil
### GetRules

`func (o *ModelSecretScanResult) GetRules() []ModelRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ModelSecretScanResult) GetRulesOk() (*[]ModelRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ModelSecretScanResult) SetRules(v []ModelRule)`

SetRules sets Rules field to given value.


### SetRulesNil

`func (o *ModelSecretScanResult) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *ModelSecretScanResult) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil
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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


