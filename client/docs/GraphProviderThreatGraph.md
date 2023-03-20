# GraphProviderThreatGraph

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudComplianceCount** | **int32** |  | 
**ComplianceCount** | **int32** |  | 
**Resources** | [**[]GraphThreatNodeInfo**](GraphThreatNodeInfo.md) |  | 
**SecretsCount** | **int32** |  | 
**VulnerabilityCount** | **int32** |  | 

## Methods

### NewGraphProviderThreatGraph

`func NewGraphProviderThreatGraph(cloudComplianceCount int32, complianceCount int32, resources []GraphThreatNodeInfo, secretsCount int32, vulnerabilityCount int32, ) *GraphProviderThreatGraph`

NewGraphProviderThreatGraph instantiates a new GraphProviderThreatGraph object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphProviderThreatGraphWithDefaults

`func NewGraphProviderThreatGraphWithDefaults() *GraphProviderThreatGraph`

NewGraphProviderThreatGraphWithDefaults instantiates a new GraphProviderThreatGraph object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudComplianceCount

`func (o *GraphProviderThreatGraph) GetCloudComplianceCount() int32`

GetCloudComplianceCount returns the CloudComplianceCount field if non-nil, zero value otherwise.

### GetCloudComplianceCountOk

`func (o *GraphProviderThreatGraph) GetCloudComplianceCountOk() (*int32, bool)`

GetCloudComplianceCountOk returns a tuple with the CloudComplianceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudComplianceCount

`func (o *GraphProviderThreatGraph) SetCloudComplianceCount(v int32)`

SetCloudComplianceCount sets CloudComplianceCount field to given value.


### GetComplianceCount

`func (o *GraphProviderThreatGraph) GetComplianceCount() int32`

GetComplianceCount returns the ComplianceCount field if non-nil, zero value otherwise.

### GetComplianceCountOk

`func (o *GraphProviderThreatGraph) GetComplianceCountOk() (*int32, bool)`

GetComplianceCountOk returns a tuple with the ComplianceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceCount

`func (o *GraphProviderThreatGraph) SetComplianceCount(v int32)`

SetComplianceCount sets ComplianceCount field to given value.


### GetResources

`func (o *GraphProviderThreatGraph) GetResources() []GraphThreatNodeInfo`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *GraphProviderThreatGraph) GetResourcesOk() (*[]GraphThreatNodeInfo, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *GraphProviderThreatGraph) SetResources(v []GraphThreatNodeInfo)`

SetResources sets Resources field to given value.


### SetResourcesNil

`func (o *GraphProviderThreatGraph) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *GraphProviderThreatGraph) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil
### GetSecretsCount

`func (o *GraphProviderThreatGraph) GetSecretsCount() int32`

GetSecretsCount returns the SecretsCount field if non-nil, zero value otherwise.

### GetSecretsCountOk

`func (o *GraphProviderThreatGraph) GetSecretsCountOk() (*int32, bool)`

GetSecretsCountOk returns a tuple with the SecretsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsCount

`func (o *GraphProviderThreatGraph) SetSecretsCount(v int32)`

SetSecretsCount sets SecretsCount field to given value.


### GetVulnerabilityCount

`func (o *GraphProviderThreatGraph) GetVulnerabilityCount() int32`

GetVulnerabilityCount returns the VulnerabilityCount field if non-nil, zero value otherwise.

### GetVulnerabilityCountOk

`func (o *GraphProviderThreatGraph) GetVulnerabilityCountOk() (*int32, bool)`

GetVulnerabilityCountOk returns a tuple with the VulnerabilityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityCount

`func (o *GraphProviderThreatGraph) SetVulnerabilityCount(v int32)`

SetVulnerabilityCount sets VulnerabilityCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


