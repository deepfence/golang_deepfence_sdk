# ModelSecretScanResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageLayerId** | **string** |  | 
**ContainerName** | **string** |  | 
**HostName** | **string** |  | 
**KubernetesClusterName** | **string** |  | 
**Masked** | **string** |  | 
**NodeId** | **string** |  | 
**NodeName** | **string** |  | 
**NodeType** | **string** |  | 
**Rule2Secrets** | **map[string][]int32** |  | 
**Rules** | [**[]ModelRule**](ModelRule.md) |  | 
**ScanId** | **string** |  | 
**Secrets** | [**[]ModelSecret**](ModelSecret.md) |  | 

## Methods

### NewModelSecretScanResult

`func NewModelSecretScanResult(imageLayerId string, containerName string, hostName string, kubernetesClusterName string, masked string, nodeId string, nodeName string, nodeType string, rule2Secrets map[string][]int32, rules []ModelRule, scanId string, secrets []ModelSecret, ) *ModelSecretScanResult`

NewModelSecretScanResult instantiates a new ModelSecretScanResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSecretScanResultWithDefaults

`func NewModelSecretScanResultWithDefaults() *ModelSecretScanResult`

NewModelSecretScanResultWithDefaults instantiates a new ModelSecretScanResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageLayerId

`func (o *ModelSecretScanResult) GetImageLayerId() string`

GetImageLayerId returns the ImageLayerId field if non-nil, zero value otherwise.

### GetImageLayerIdOk

`func (o *ModelSecretScanResult) GetImageLayerIdOk() (*string, bool)`

GetImageLayerIdOk returns a tuple with the ImageLayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageLayerId

`func (o *ModelSecretScanResult) SetImageLayerId(v string)`

SetImageLayerId sets ImageLayerId field to given value.


### GetContainerName

`func (o *ModelSecretScanResult) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *ModelSecretScanResult) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *ModelSecretScanResult) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


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


### GetMasked

`func (o *ModelSecretScanResult) GetMasked() string`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *ModelSecretScanResult) GetMaskedOk() (*string, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *ModelSecretScanResult) SetMasked(v string)`

SetMasked sets Masked field to given value.


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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


