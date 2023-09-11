# IngestersCloudResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLevel** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**AllowBlobPublicAccess** | Pointer to **bool** |  | [optional] 
**Arn** | Pointer to **string** |  | [optional] 
**AttachedPolicyArns** | Pointer to **interface{}** |  | [optional] 
**BlockPublicAcls** | Pointer to **bool** |  | [optional] 
**BlockPublicPolicy** | Pointer to **bool** |  | [optional] 
**BucketPolicyIsPublic** | Pointer to **bool** |  | [optional] 
**CidrIpv4** | Pointer to **string** |  | [optional] 
**CloudProvider** | Pointer to **string** |  | [optional] 
**ClusterArn** | Pointer to **string** |  | [optional] 
**ClusterName** | Pointer to **string** |  | [optional] 
**Connectivity** | Pointer to **string** |  | [optional] 
**ContainerDefinitions** | Pointer to **interface{}** |  | [optional] 
**Containers** | Pointer to **interface{}** |  | [optional] 
**CreateDate** | Pointer to **string** |  | [optional] 
**DbClusterIdentifier** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EventNotificationConfiguration** | Pointer to **interface{}** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to **interface{}** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**IamInstanceProfileArn** | Pointer to **string** |  | [optional] 
**IamInstanceProfileId** | Pointer to **string** |  | [optional] 
**IamPolicy** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IgnorePublicAcls** | Pointer to **bool** |  | [optional] 
**IngressSettings** | Pointer to **string** |  | [optional] 
**InlinePolicies** | Pointer to **interface{}** |  | [optional] 
**InstanceId** | Pointer to **string** |  | [optional] 
**InstanceProfileArns** | Pointer to **interface{}** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**Instances** | Pointer to **interface{}** |  | [optional] 
**IpConfiguration** | Pointer to **interface{}** |  | [optional] 
**IsEgress** | Pointer to **bool** |  | [optional] 
**LastStatus** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NetworkConfiguration** | Pointer to **interface{}** |  | [optional] 
**NetworkInterfaces** | Pointer to **interface{}** |  | [optional] 
**NetworkMode** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**OrganizationMasterAccountArn** | Pointer to **string** |  | [optional] 
**OrganizationMasterAccountEmail** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Policy** | Pointer to **interface{}** |  | [optional] 
**PolicyStd** | Pointer to **interface{}** |  | [optional] 
**PrivateDnsName** | Pointer to **string** |  | [optional] 
**PrivateIpAddress** | Pointer to **string** |  | [optional] 
**Privilege** | Pointer to **string** |  | [optional] 
**PublicAccess** | Pointer to **string** |  | [optional] 
**PublicIpAddress** | Pointer to **string** |  | [optional] 
**PublicIps** | Pointer to **interface{}** |  | [optional] 
**PublicNetworkAccess** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**ResourceId** | Pointer to **string** |  | [optional] 
**ResourceVpcConfig** | Pointer to **interface{}** |  | [optional] 
**ResourcesVpcConfig** | Pointer to **interface{}** |  | [optional] 
**RestrictPublicBuckets** | Pointer to **bool** |  | [optional] 
**Scheme** | Pointer to **string** |  | [optional] 
**SecurityGroups** | Pointer to **interface{}** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**StorageAccountName** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **interface{}** |  | [optional] 
**TargetGroupArn** | Pointer to **string** |  | [optional] 
**TargetHealthDescriptions** | Pointer to **interface{}** |  | [optional] 
**TaskArn** | Pointer to **string** |  | [optional] 
**TaskDefinition** | Pointer to **interface{}** |  | [optional] 
**TaskDefinitionArn** | Pointer to **string** |  | [optional] 
**UserGroups** | Pointer to **interface{}** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Users** | Pointer to **interface{}** |  | [optional] 
**VpcId** | Pointer to **string** |  | [optional] 
**VpcOptions** | Pointer to **interface{}** |  | [optional] 
**VpcSecurityGroupIds** | Pointer to **interface{}** |  | [optional] 
**VpcSecurityGroups** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewIngestersCloudResource

`func NewIngestersCloudResource() *IngestersCloudResource`

NewIngestersCloudResource instantiates a new IngestersCloudResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestersCloudResourceWithDefaults

`func NewIngestersCloudResourceWithDefaults() *IngestersCloudResource`

NewIngestersCloudResourceWithDefaults instantiates a new IngestersCloudResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessLevel

`func (o *IngestersCloudResource) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *IngestersCloudResource) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *IngestersCloudResource) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *IngestersCloudResource) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetAccountId

`func (o *IngestersCloudResource) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *IngestersCloudResource) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *IngestersCloudResource) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *IngestersCloudResource) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAction

`func (o *IngestersCloudResource) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IngestersCloudResource) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IngestersCloudResource) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *IngestersCloudResource) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAllowBlobPublicAccess

`func (o *IngestersCloudResource) GetAllowBlobPublicAccess() bool`

GetAllowBlobPublicAccess returns the AllowBlobPublicAccess field if non-nil, zero value otherwise.

### GetAllowBlobPublicAccessOk

`func (o *IngestersCloudResource) GetAllowBlobPublicAccessOk() (*bool, bool)`

GetAllowBlobPublicAccessOk returns a tuple with the AllowBlobPublicAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBlobPublicAccess

`func (o *IngestersCloudResource) SetAllowBlobPublicAccess(v bool)`

SetAllowBlobPublicAccess sets AllowBlobPublicAccess field to given value.

### HasAllowBlobPublicAccess

`func (o *IngestersCloudResource) HasAllowBlobPublicAccess() bool`

HasAllowBlobPublicAccess returns a boolean if a field has been set.

### GetArn

`func (o *IngestersCloudResource) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *IngestersCloudResource) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *IngestersCloudResource) SetArn(v string)`

SetArn sets Arn field to given value.

### HasArn

`func (o *IngestersCloudResource) HasArn() bool`

HasArn returns a boolean if a field has been set.

### GetAttachedPolicyArns

`func (o *IngestersCloudResource) GetAttachedPolicyArns() interface{}`

GetAttachedPolicyArns returns the AttachedPolicyArns field if non-nil, zero value otherwise.

### GetAttachedPolicyArnsOk

`func (o *IngestersCloudResource) GetAttachedPolicyArnsOk() (*interface{}, bool)`

GetAttachedPolicyArnsOk returns a tuple with the AttachedPolicyArns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedPolicyArns

`func (o *IngestersCloudResource) SetAttachedPolicyArns(v interface{})`

SetAttachedPolicyArns sets AttachedPolicyArns field to given value.

### HasAttachedPolicyArns

`func (o *IngestersCloudResource) HasAttachedPolicyArns() bool`

HasAttachedPolicyArns returns a boolean if a field has been set.

### SetAttachedPolicyArnsNil

`func (o *IngestersCloudResource) SetAttachedPolicyArnsNil(b bool)`

 SetAttachedPolicyArnsNil sets the value for AttachedPolicyArns to be an explicit nil

### UnsetAttachedPolicyArns
`func (o *IngestersCloudResource) UnsetAttachedPolicyArns()`

UnsetAttachedPolicyArns ensures that no value is present for AttachedPolicyArns, not even an explicit nil
### GetBlockPublicAcls

`func (o *IngestersCloudResource) GetBlockPublicAcls() bool`

GetBlockPublicAcls returns the BlockPublicAcls field if non-nil, zero value otherwise.

### GetBlockPublicAclsOk

`func (o *IngestersCloudResource) GetBlockPublicAclsOk() (*bool, bool)`

GetBlockPublicAclsOk returns a tuple with the BlockPublicAcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockPublicAcls

`func (o *IngestersCloudResource) SetBlockPublicAcls(v bool)`

SetBlockPublicAcls sets BlockPublicAcls field to given value.

### HasBlockPublicAcls

`func (o *IngestersCloudResource) HasBlockPublicAcls() bool`

HasBlockPublicAcls returns a boolean if a field has been set.

### GetBlockPublicPolicy

`func (o *IngestersCloudResource) GetBlockPublicPolicy() bool`

GetBlockPublicPolicy returns the BlockPublicPolicy field if non-nil, zero value otherwise.

### GetBlockPublicPolicyOk

`func (o *IngestersCloudResource) GetBlockPublicPolicyOk() (*bool, bool)`

GetBlockPublicPolicyOk returns a tuple with the BlockPublicPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockPublicPolicy

`func (o *IngestersCloudResource) SetBlockPublicPolicy(v bool)`

SetBlockPublicPolicy sets BlockPublicPolicy field to given value.

### HasBlockPublicPolicy

`func (o *IngestersCloudResource) HasBlockPublicPolicy() bool`

HasBlockPublicPolicy returns a boolean if a field has been set.

### GetBucketPolicyIsPublic

`func (o *IngestersCloudResource) GetBucketPolicyIsPublic() bool`

GetBucketPolicyIsPublic returns the BucketPolicyIsPublic field if non-nil, zero value otherwise.

### GetBucketPolicyIsPublicOk

`func (o *IngestersCloudResource) GetBucketPolicyIsPublicOk() (*bool, bool)`

GetBucketPolicyIsPublicOk returns a tuple with the BucketPolicyIsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketPolicyIsPublic

`func (o *IngestersCloudResource) SetBucketPolicyIsPublic(v bool)`

SetBucketPolicyIsPublic sets BucketPolicyIsPublic field to given value.

### HasBucketPolicyIsPublic

`func (o *IngestersCloudResource) HasBucketPolicyIsPublic() bool`

HasBucketPolicyIsPublic returns a boolean if a field has been set.

### GetCidrIpv4

`func (o *IngestersCloudResource) GetCidrIpv4() string`

GetCidrIpv4 returns the CidrIpv4 field if non-nil, zero value otherwise.

### GetCidrIpv4Ok

`func (o *IngestersCloudResource) GetCidrIpv4Ok() (*string, bool)`

GetCidrIpv4Ok returns a tuple with the CidrIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrIpv4

`func (o *IngestersCloudResource) SetCidrIpv4(v string)`

SetCidrIpv4 sets CidrIpv4 field to given value.

### HasCidrIpv4

`func (o *IngestersCloudResource) HasCidrIpv4() bool`

HasCidrIpv4 returns a boolean if a field has been set.

### GetCloudProvider

`func (o *IngestersCloudResource) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *IngestersCloudResource) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *IngestersCloudResource) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *IngestersCloudResource) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetClusterArn

`func (o *IngestersCloudResource) GetClusterArn() string`

GetClusterArn returns the ClusterArn field if non-nil, zero value otherwise.

### GetClusterArnOk

`func (o *IngestersCloudResource) GetClusterArnOk() (*string, bool)`

GetClusterArnOk returns a tuple with the ClusterArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterArn

`func (o *IngestersCloudResource) SetClusterArn(v string)`

SetClusterArn sets ClusterArn field to given value.

### HasClusterArn

`func (o *IngestersCloudResource) HasClusterArn() bool`

HasClusterArn returns a boolean if a field has been set.

### GetClusterName

`func (o *IngestersCloudResource) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *IngestersCloudResource) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *IngestersCloudResource) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *IngestersCloudResource) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetConnectivity

`func (o *IngestersCloudResource) GetConnectivity() string`

GetConnectivity returns the Connectivity field if non-nil, zero value otherwise.

### GetConnectivityOk

`func (o *IngestersCloudResource) GetConnectivityOk() (*string, bool)`

GetConnectivityOk returns a tuple with the Connectivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivity

`func (o *IngestersCloudResource) SetConnectivity(v string)`

SetConnectivity sets Connectivity field to given value.

### HasConnectivity

`func (o *IngestersCloudResource) HasConnectivity() bool`

HasConnectivity returns a boolean if a field has been set.

### GetContainerDefinitions

`func (o *IngestersCloudResource) GetContainerDefinitions() interface{}`

GetContainerDefinitions returns the ContainerDefinitions field if non-nil, zero value otherwise.

### GetContainerDefinitionsOk

`func (o *IngestersCloudResource) GetContainerDefinitionsOk() (*interface{}, bool)`

GetContainerDefinitionsOk returns a tuple with the ContainerDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerDefinitions

`func (o *IngestersCloudResource) SetContainerDefinitions(v interface{})`

SetContainerDefinitions sets ContainerDefinitions field to given value.

### HasContainerDefinitions

`func (o *IngestersCloudResource) HasContainerDefinitions() bool`

HasContainerDefinitions returns a boolean if a field has been set.

### SetContainerDefinitionsNil

`func (o *IngestersCloudResource) SetContainerDefinitionsNil(b bool)`

 SetContainerDefinitionsNil sets the value for ContainerDefinitions to be an explicit nil

### UnsetContainerDefinitions
`func (o *IngestersCloudResource) UnsetContainerDefinitions()`

UnsetContainerDefinitions ensures that no value is present for ContainerDefinitions, not even an explicit nil
### GetContainers

`func (o *IngestersCloudResource) GetContainers() interface{}`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *IngestersCloudResource) GetContainersOk() (*interface{}, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *IngestersCloudResource) SetContainers(v interface{})`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *IngestersCloudResource) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### SetContainersNil

`func (o *IngestersCloudResource) SetContainersNil(b bool)`

 SetContainersNil sets the value for Containers to be an explicit nil

### UnsetContainers
`func (o *IngestersCloudResource) UnsetContainers()`

UnsetContainers ensures that no value is present for Containers, not even an explicit nil
### GetCreateDate

`func (o *IngestersCloudResource) GetCreateDate() string`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *IngestersCloudResource) GetCreateDateOk() (*string, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *IngestersCloudResource) SetCreateDate(v string)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *IngestersCloudResource) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetDbClusterIdentifier

`func (o *IngestersCloudResource) GetDbClusterIdentifier() string`

GetDbClusterIdentifier returns the DbClusterIdentifier field if non-nil, zero value otherwise.

### GetDbClusterIdentifierOk

`func (o *IngestersCloudResource) GetDbClusterIdentifierOk() (*string, bool)`

GetDbClusterIdentifierOk returns a tuple with the DbClusterIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbClusterIdentifier

`func (o *IngestersCloudResource) SetDbClusterIdentifier(v string)`

SetDbClusterIdentifier sets DbClusterIdentifier field to given value.

### HasDbClusterIdentifier

`func (o *IngestersCloudResource) HasDbClusterIdentifier() bool`

HasDbClusterIdentifier returns a boolean if a field has been set.

### GetDescription

`func (o *IngestersCloudResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IngestersCloudResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IngestersCloudResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IngestersCloudResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEventNotificationConfiguration

`func (o *IngestersCloudResource) GetEventNotificationConfiguration() interface{}`

GetEventNotificationConfiguration returns the EventNotificationConfiguration field if non-nil, zero value otherwise.

### GetEventNotificationConfigurationOk

`func (o *IngestersCloudResource) GetEventNotificationConfigurationOk() (*interface{}, bool)`

GetEventNotificationConfigurationOk returns a tuple with the EventNotificationConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotificationConfiguration

`func (o *IngestersCloudResource) SetEventNotificationConfiguration(v interface{})`

SetEventNotificationConfiguration sets EventNotificationConfiguration field to given value.

### HasEventNotificationConfiguration

`func (o *IngestersCloudResource) HasEventNotificationConfiguration() bool`

HasEventNotificationConfiguration returns a boolean if a field has been set.

### SetEventNotificationConfigurationNil

`func (o *IngestersCloudResource) SetEventNotificationConfigurationNil(b bool)`

 SetEventNotificationConfigurationNil sets the value for EventNotificationConfiguration to be an explicit nil

### UnsetEventNotificationConfiguration
`func (o *IngestersCloudResource) UnsetEventNotificationConfiguration()`

UnsetEventNotificationConfiguration ensures that no value is present for EventNotificationConfiguration, not even an explicit nil
### GetGroup

`func (o *IngestersCloudResource) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *IngestersCloudResource) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *IngestersCloudResource) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *IngestersCloudResource) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetGroupId

`func (o *IngestersCloudResource) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *IngestersCloudResource) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *IngestersCloudResource) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *IngestersCloudResource) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroups

`func (o *IngestersCloudResource) GetGroups() interface{}`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *IngestersCloudResource) GetGroupsOk() (*interface{}, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *IngestersCloudResource) SetGroups(v interface{})`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *IngestersCloudResource) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *IngestersCloudResource) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *IngestersCloudResource) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetHostName

`func (o *IngestersCloudResource) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *IngestersCloudResource) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *IngestersCloudResource) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *IngestersCloudResource) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetIamInstanceProfileArn

`func (o *IngestersCloudResource) GetIamInstanceProfileArn() string`

GetIamInstanceProfileArn returns the IamInstanceProfileArn field if non-nil, zero value otherwise.

### GetIamInstanceProfileArnOk

`func (o *IngestersCloudResource) GetIamInstanceProfileArnOk() (*string, bool)`

GetIamInstanceProfileArnOk returns a tuple with the IamInstanceProfileArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamInstanceProfileArn

`func (o *IngestersCloudResource) SetIamInstanceProfileArn(v string)`

SetIamInstanceProfileArn sets IamInstanceProfileArn field to given value.

### HasIamInstanceProfileArn

`func (o *IngestersCloudResource) HasIamInstanceProfileArn() bool`

HasIamInstanceProfileArn returns a boolean if a field has been set.

### GetIamInstanceProfileId

`func (o *IngestersCloudResource) GetIamInstanceProfileId() string`

GetIamInstanceProfileId returns the IamInstanceProfileId field if non-nil, zero value otherwise.

### GetIamInstanceProfileIdOk

`func (o *IngestersCloudResource) GetIamInstanceProfileIdOk() (*string, bool)`

GetIamInstanceProfileIdOk returns a tuple with the IamInstanceProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamInstanceProfileId

`func (o *IngestersCloudResource) SetIamInstanceProfileId(v string)`

SetIamInstanceProfileId sets IamInstanceProfileId field to given value.

### HasIamInstanceProfileId

`func (o *IngestersCloudResource) HasIamInstanceProfileId() bool`

HasIamInstanceProfileId returns a boolean if a field has been set.

### GetIamPolicy

`func (o *IngestersCloudResource) GetIamPolicy() interface{}`

GetIamPolicy returns the IamPolicy field if non-nil, zero value otherwise.

### GetIamPolicyOk

`func (o *IngestersCloudResource) GetIamPolicyOk() (*interface{}, bool)`

GetIamPolicyOk returns a tuple with the IamPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamPolicy

`func (o *IngestersCloudResource) SetIamPolicy(v interface{})`

SetIamPolicy sets IamPolicy field to given value.

### HasIamPolicy

`func (o *IngestersCloudResource) HasIamPolicy() bool`

HasIamPolicy returns a boolean if a field has been set.

### SetIamPolicyNil

`func (o *IngestersCloudResource) SetIamPolicyNil(b bool)`

 SetIamPolicyNil sets the value for IamPolicy to be an explicit nil

### UnsetIamPolicy
`func (o *IngestersCloudResource) UnsetIamPolicy()`

UnsetIamPolicy ensures that no value is present for IamPolicy, not even an explicit nil
### GetId

`func (o *IngestersCloudResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IngestersCloudResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IngestersCloudResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IngestersCloudResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgnorePublicAcls

`func (o *IngestersCloudResource) GetIgnorePublicAcls() bool`

GetIgnorePublicAcls returns the IgnorePublicAcls field if non-nil, zero value otherwise.

### GetIgnorePublicAclsOk

`func (o *IngestersCloudResource) GetIgnorePublicAclsOk() (*bool, bool)`

GetIgnorePublicAclsOk returns a tuple with the IgnorePublicAcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnorePublicAcls

`func (o *IngestersCloudResource) SetIgnorePublicAcls(v bool)`

SetIgnorePublicAcls sets IgnorePublicAcls field to given value.

### HasIgnorePublicAcls

`func (o *IngestersCloudResource) HasIgnorePublicAcls() bool`

HasIgnorePublicAcls returns a boolean if a field has been set.

### GetIngressSettings

`func (o *IngestersCloudResource) GetIngressSettings() string`

GetIngressSettings returns the IngressSettings field if non-nil, zero value otherwise.

### GetIngressSettingsOk

`func (o *IngestersCloudResource) GetIngressSettingsOk() (*string, bool)`

GetIngressSettingsOk returns a tuple with the IngressSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressSettings

`func (o *IngestersCloudResource) SetIngressSettings(v string)`

SetIngressSettings sets IngressSettings field to given value.

### HasIngressSettings

`func (o *IngestersCloudResource) HasIngressSettings() bool`

HasIngressSettings returns a boolean if a field has been set.

### GetInlinePolicies

`func (o *IngestersCloudResource) GetInlinePolicies() interface{}`

GetInlinePolicies returns the InlinePolicies field if non-nil, zero value otherwise.

### GetInlinePoliciesOk

`func (o *IngestersCloudResource) GetInlinePoliciesOk() (*interface{}, bool)`

GetInlinePoliciesOk returns a tuple with the InlinePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlinePolicies

`func (o *IngestersCloudResource) SetInlinePolicies(v interface{})`

SetInlinePolicies sets InlinePolicies field to given value.

### HasInlinePolicies

`func (o *IngestersCloudResource) HasInlinePolicies() bool`

HasInlinePolicies returns a boolean if a field has been set.

### SetInlinePoliciesNil

`func (o *IngestersCloudResource) SetInlinePoliciesNil(b bool)`

 SetInlinePoliciesNil sets the value for InlinePolicies to be an explicit nil

### UnsetInlinePolicies
`func (o *IngestersCloudResource) UnsetInlinePolicies()`

UnsetInlinePolicies ensures that no value is present for InlinePolicies, not even an explicit nil
### GetInstanceId

`func (o *IngestersCloudResource) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *IngestersCloudResource) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *IngestersCloudResource) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *IngestersCloudResource) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInstanceProfileArns

`func (o *IngestersCloudResource) GetInstanceProfileArns() interface{}`

GetInstanceProfileArns returns the InstanceProfileArns field if non-nil, zero value otherwise.

### GetInstanceProfileArnsOk

`func (o *IngestersCloudResource) GetInstanceProfileArnsOk() (*interface{}, bool)`

GetInstanceProfileArnsOk returns a tuple with the InstanceProfileArns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceProfileArns

`func (o *IngestersCloudResource) SetInstanceProfileArns(v interface{})`

SetInstanceProfileArns sets InstanceProfileArns field to given value.

### HasInstanceProfileArns

`func (o *IngestersCloudResource) HasInstanceProfileArns() bool`

HasInstanceProfileArns returns a boolean if a field has been set.

### SetInstanceProfileArnsNil

`func (o *IngestersCloudResource) SetInstanceProfileArnsNil(b bool)`

 SetInstanceProfileArnsNil sets the value for InstanceProfileArns to be an explicit nil

### UnsetInstanceProfileArns
`func (o *IngestersCloudResource) UnsetInstanceProfileArns()`

UnsetInstanceProfileArns ensures that no value is present for InstanceProfileArns, not even an explicit nil
### GetInstanceType

`func (o *IngestersCloudResource) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *IngestersCloudResource) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *IngestersCloudResource) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *IngestersCloudResource) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetInstances

`func (o *IngestersCloudResource) GetInstances() interface{}`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *IngestersCloudResource) GetInstancesOk() (*interface{}, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *IngestersCloudResource) SetInstances(v interface{})`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *IngestersCloudResource) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### SetInstancesNil

`func (o *IngestersCloudResource) SetInstancesNil(b bool)`

 SetInstancesNil sets the value for Instances to be an explicit nil

### UnsetInstances
`func (o *IngestersCloudResource) UnsetInstances()`

UnsetInstances ensures that no value is present for Instances, not even an explicit nil
### GetIpConfiguration

`func (o *IngestersCloudResource) GetIpConfiguration() interface{}`

GetIpConfiguration returns the IpConfiguration field if non-nil, zero value otherwise.

### GetIpConfigurationOk

`func (o *IngestersCloudResource) GetIpConfigurationOk() (*interface{}, bool)`

GetIpConfigurationOk returns a tuple with the IpConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfiguration

`func (o *IngestersCloudResource) SetIpConfiguration(v interface{})`

SetIpConfiguration sets IpConfiguration field to given value.

### HasIpConfiguration

`func (o *IngestersCloudResource) HasIpConfiguration() bool`

HasIpConfiguration returns a boolean if a field has been set.

### SetIpConfigurationNil

`func (o *IngestersCloudResource) SetIpConfigurationNil(b bool)`

 SetIpConfigurationNil sets the value for IpConfiguration to be an explicit nil

### UnsetIpConfiguration
`func (o *IngestersCloudResource) UnsetIpConfiguration()`

UnsetIpConfiguration ensures that no value is present for IpConfiguration, not even an explicit nil
### GetIsEgress

`func (o *IngestersCloudResource) GetIsEgress() bool`

GetIsEgress returns the IsEgress field if non-nil, zero value otherwise.

### GetIsEgressOk

`func (o *IngestersCloudResource) GetIsEgressOk() (*bool, bool)`

GetIsEgressOk returns a tuple with the IsEgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEgress

`func (o *IngestersCloudResource) SetIsEgress(v bool)`

SetIsEgress sets IsEgress field to given value.

### HasIsEgress

`func (o *IngestersCloudResource) HasIsEgress() bool`

HasIsEgress returns a boolean if a field has been set.

### GetLastStatus

`func (o *IngestersCloudResource) GetLastStatus() string`

GetLastStatus returns the LastStatus field if non-nil, zero value otherwise.

### GetLastStatusOk

`func (o *IngestersCloudResource) GetLastStatusOk() (*string, bool)`

GetLastStatusOk returns a tuple with the LastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatus

`func (o *IngestersCloudResource) SetLastStatus(v string)`

SetLastStatus sets LastStatus field to given value.

### HasLastStatus

`func (o *IngestersCloudResource) HasLastStatus() bool`

HasLastStatus returns a boolean if a field has been set.

### GetName

`func (o *IngestersCloudResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IngestersCloudResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IngestersCloudResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IngestersCloudResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkConfiguration

`func (o *IngestersCloudResource) GetNetworkConfiguration() interface{}`

GetNetworkConfiguration returns the NetworkConfiguration field if non-nil, zero value otherwise.

### GetNetworkConfigurationOk

`func (o *IngestersCloudResource) GetNetworkConfigurationOk() (*interface{}, bool)`

GetNetworkConfigurationOk returns a tuple with the NetworkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfiguration

`func (o *IngestersCloudResource) SetNetworkConfiguration(v interface{})`

SetNetworkConfiguration sets NetworkConfiguration field to given value.

### HasNetworkConfiguration

`func (o *IngestersCloudResource) HasNetworkConfiguration() bool`

HasNetworkConfiguration returns a boolean if a field has been set.

### SetNetworkConfigurationNil

`func (o *IngestersCloudResource) SetNetworkConfigurationNil(b bool)`

 SetNetworkConfigurationNil sets the value for NetworkConfiguration to be an explicit nil

### UnsetNetworkConfiguration
`func (o *IngestersCloudResource) UnsetNetworkConfiguration()`

UnsetNetworkConfiguration ensures that no value is present for NetworkConfiguration, not even an explicit nil
### GetNetworkInterfaces

`func (o *IngestersCloudResource) GetNetworkInterfaces() interface{}`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *IngestersCloudResource) GetNetworkInterfacesOk() (*interface{}, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *IngestersCloudResource) SetNetworkInterfaces(v interface{})`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *IngestersCloudResource) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### SetNetworkInterfacesNil

`func (o *IngestersCloudResource) SetNetworkInterfacesNil(b bool)`

 SetNetworkInterfacesNil sets the value for NetworkInterfaces to be an explicit nil

### UnsetNetworkInterfaces
`func (o *IngestersCloudResource) UnsetNetworkInterfaces()`

UnsetNetworkInterfaces ensures that no value is present for NetworkInterfaces, not even an explicit nil
### GetNetworkMode

`func (o *IngestersCloudResource) GetNetworkMode() string`

GetNetworkMode returns the NetworkMode field if non-nil, zero value otherwise.

### GetNetworkModeOk

`func (o *IngestersCloudResource) GetNetworkModeOk() (*string, bool)`

GetNetworkModeOk returns a tuple with the NetworkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMode

`func (o *IngestersCloudResource) SetNetworkMode(v string)`

SetNetworkMode sets NetworkMode field to given value.

### HasNetworkMode

`func (o *IngestersCloudResource) HasNetworkMode() bool`

HasNetworkMode returns a boolean if a field has been set.

### GetOrganizationId

`func (o *IngestersCloudResource) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *IngestersCloudResource) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *IngestersCloudResource) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *IngestersCloudResource) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOrganizationMasterAccountArn

`func (o *IngestersCloudResource) GetOrganizationMasterAccountArn() string`

GetOrganizationMasterAccountArn returns the OrganizationMasterAccountArn field if non-nil, zero value otherwise.

### GetOrganizationMasterAccountArnOk

`func (o *IngestersCloudResource) GetOrganizationMasterAccountArnOk() (*string, bool)`

GetOrganizationMasterAccountArnOk returns a tuple with the OrganizationMasterAccountArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationMasterAccountArn

`func (o *IngestersCloudResource) SetOrganizationMasterAccountArn(v string)`

SetOrganizationMasterAccountArn sets OrganizationMasterAccountArn field to given value.

### HasOrganizationMasterAccountArn

`func (o *IngestersCloudResource) HasOrganizationMasterAccountArn() bool`

HasOrganizationMasterAccountArn returns a boolean if a field has been set.

### GetOrganizationMasterAccountEmail

`func (o *IngestersCloudResource) GetOrganizationMasterAccountEmail() string`

GetOrganizationMasterAccountEmail returns the OrganizationMasterAccountEmail field if non-nil, zero value otherwise.

### GetOrganizationMasterAccountEmailOk

`func (o *IngestersCloudResource) GetOrganizationMasterAccountEmailOk() (*string, bool)`

GetOrganizationMasterAccountEmailOk returns a tuple with the OrganizationMasterAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationMasterAccountEmail

`func (o *IngestersCloudResource) SetOrganizationMasterAccountEmail(v string)`

SetOrganizationMasterAccountEmail sets OrganizationMasterAccountEmail field to given value.

### HasOrganizationMasterAccountEmail

`func (o *IngestersCloudResource) HasOrganizationMasterAccountEmail() bool`

HasOrganizationMasterAccountEmail returns a boolean if a field has been set.

### GetPath

`func (o *IngestersCloudResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *IngestersCloudResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *IngestersCloudResource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *IngestersCloudResource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPolicy

`func (o *IngestersCloudResource) GetPolicy() interface{}`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *IngestersCloudResource) GetPolicyOk() (*interface{}, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *IngestersCloudResource) SetPolicy(v interface{})`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *IngestersCloudResource) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *IngestersCloudResource) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *IngestersCloudResource) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
### GetPolicyStd

`func (o *IngestersCloudResource) GetPolicyStd() interface{}`

GetPolicyStd returns the PolicyStd field if non-nil, zero value otherwise.

### GetPolicyStdOk

`func (o *IngestersCloudResource) GetPolicyStdOk() (*interface{}, bool)`

GetPolicyStdOk returns a tuple with the PolicyStd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyStd

`func (o *IngestersCloudResource) SetPolicyStd(v interface{})`

SetPolicyStd sets PolicyStd field to given value.

### HasPolicyStd

`func (o *IngestersCloudResource) HasPolicyStd() bool`

HasPolicyStd returns a boolean if a field has been set.

### SetPolicyStdNil

`func (o *IngestersCloudResource) SetPolicyStdNil(b bool)`

 SetPolicyStdNil sets the value for PolicyStd to be an explicit nil

### UnsetPolicyStd
`func (o *IngestersCloudResource) UnsetPolicyStd()`

UnsetPolicyStd ensures that no value is present for PolicyStd, not even an explicit nil
### GetPrivateDnsName

`func (o *IngestersCloudResource) GetPrivateDnsName() string`

GetPrivateDnsName returns the PrivateDnsName field if non-nil, zero value otherwise.

### GetPrivateDnsNameOk

`func (o *IngestersCloudResource) GetPrivateDnsNameOk() (*string, bool)`

GetPrivateDnsNameOk returns a tuple with the PrivateDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDnsName

`func (o *IngestersCloudResource) SetPrivateDnsName(v string)`

SetPrivateDnsName sets PrivateDnsName field to given value.

### HasPrivateDnsName

`func (o *IngestersCloudResource) HasPrivateDnsName() bool`

HasPrivateDnsName returns a boolean if a field has been set.

### GetPrivateIpAddress

`func (o *IngestersCloudResource) GetPrivateIpAddress() string`

GetPrivateIpAddress returns the PrivateIpAddress field if non-nil, zero value otherwise.

### GetPrivateIpAddressOk

`func (o *IngestersCloudResource) GetPrivateIpAddressOk() (*string, bool)`

GetPrivateIpAddressOk returns a tuple with the PrivateIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpAddress

`func (o *IngestersCloudResource) SetPrivateIpAddress(v string)`

SetPrivateIpAddress sets PrivateIpAddress field to given value.

### HasPrivateIpAddress

`func (o *IngestersCloudResource) HasPrivateIpAddress() bool`

HasPrivateIpAddress returns a boolean if a field has been set.

### GetPrivilege

`func (o *IngestersCloudResource) GetPrivilege() string`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *IngestersCloudResource) GetPrivilegeOk() (*string, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *IngestersCloudResource) SetPrivilege(v string)`

SetPrivilege sets Privilege field to given value.

### HasPrivilege

`func (o *IngestersCloudResource) HasPrivilege() bool`

HasPrivilege returns a boolean if a field has been set.

### GetPublicAccess

`func (o *IngestersCloudResource) GetPublicAccess() string`

GetPublicAccess returns the PublicAccess field if non-nil, zero value otherwise.

### GetPublicAccessOk

`func (o *IngestersCloudResource) GetPublicAccessOk() (*string, bool)`

GetPublicAccessOk returns a tuple with the PublicAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAccess

`func (o *IngestersCloudResource) SetPublicAccess(v string)`

SetPublicAccess sets PublicAccess field to given value.

### HasPublicAccess

`func (o *IngestersCloudResource) HasPublicAccess() bool`

HasPublicAccess returns a boolean if a field has been set.

### GetPublicIpAddress

`func (o *IngestersCloudResource) GetPublicIpAddress() string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *IngestersCloudResource) GetPublicIpAddressOk() (*string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddress

`func (o *IngestersCloudResource) SetPublicIpAddress(v string)`

SetPublicIpAddress sets PublicIpAddress field to given value.

### HasPublicIpAddress

`func (o *IngestersCloudResource) HasPublicIpAddress() bool`

HasPublicIpAddress returns a boolean if a field has been set.

### GetPublicIps

`func (o *IngestersCloudResource) GetPublicIps() interface{}`

GetPublicIps returns the PublicIps field if non-nil, zero value otherwise.

### GetPublicIpsOk

`func (o *IngestersCloudResource) GetPublicIpsOk() (*interface{}, bool)`

GetPublicIpsOk returns a tuple with the PublicIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIps

`func (o *IngestersCloudResource) SetPublicIps(v interface{})`

SetPublicIps sets PublicIps field to given value.

### HasPublicIps

`func (o *IngestersCloudResource) HasPublicIps() bool`

HasPublicIps returns a boolean if a field has been set.

### SetPublicIpsNil

`func (o *IngestersCloudResource) SetPublicIpsNil(b bool)`

 SetPublicIpsNil sets the value for PublicIps to be an explicit nil

### UnsetPublicIps
`func (o *IngestersCloudResource) UnsetPublicIps()`

UnsetPublicIps ensures that no value is present for PublicIps, not even an explicit nil
### GetPublicNetworkAccess

`func (o *IngestersCloudResource) GetPublicNetworkAccess() string`

GetPublicNetworkAccess returns the PublicNetworkAccess field if non-nil, zero value otherwise.

### GetPublicNetworkAccessOk

`func (o *IngestersCloudResource) GetPublicNetworkAccessOk() (*string, bool)`

GetPublicNetworkAccessOk returns a tuple with the PublicNetworkAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetworkAccess

`func (o *IngestersCloudResource) SetPublicNetworkAccess(v string)`

SetPublicNetworkAccess sets PublicNetworkAccess field to given value.

### HasPublicNetworkAccess

`func (o *IngestersCloudResource) HasPublicNetworkAccess() bool`

HasPublicNetworkAccess returns a boolean if a field has been set.

### GetRegion

`func (o *IngestersCloudResource) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *IngestersCloudResource) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *IngestersCloudResource) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *IngestersCloudResource) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetResourceId

`func (o *IngestersCloudResource) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *IngestersCloudResource) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *IngestersCloudResource) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *IngestersCloudResource) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceVpcConfig

`func (o *IngestersCloudResource) GetResourceVpcConfig() interface{}`

GetResourceVpcConfig returns the ResourceVpcConfig field if non-nil, zero value otherwise.

### GetResourceVpcConfigOk

`func (o *IngestersCloudResource) GetResourceVpcConfigOk() (*interface{}, bool)`

GetResourceVpcConfigOk returns a tuple with the ResourceVpcConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVpcConfig

`func (o *IngestersCloudResource) SetResourceVpcConfig(v interface{})`

SetResourceVpcConfig sets ResourceVpcConfig field to given value.

### HasResourceVpcConfig

`func (o *IngestersCloudResource) HasResourceVpcConfig() bool`

HasResourceVpcConfig returns a boolean if a field has been set.

### SetResourceVpcConfigNil

`func (o *IngestersCloudResource) SetResourceVpcConfigNil(b bool)`

 SetResourceVpcConfigNil sets the value for ResourceVpcConfig to be an explicit nil

### UnsetResourceVpcConfig
`func (o *IngestersCloudResource) UnsetResourceVpcConfig()`

UnsetResourceVpcConfig ensures that no value is present for ResourceVpcConfig, not even an explicit nil
### GetResourcesVpcConfig

`func (o *IngestersCloudResource) GetResourcesVpcConfig() interface{}`

GetResourcesVpcConfig returns the ResourcesVpcConfig field if non-nil, zero value otherwise.

### GetResourcesVpcConfigOk

`func (o *IngestersCloudResource) GetResourcesVpcConfigOk() (*interface{}, bool)`

GetResourcesVpcConfigOk returns a tuple with the ResourcesVpcConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesVpcConfig

`func (o *IngestersCloudResource) SetResourcesVpcConfig(v interface{})`

SetResourcesVpcConfig sets ResourcesVpcConfig field to given value.

### HasResourcesVpcConfig

`func (o *IngestersCloudResource) HasResourcesVpcConfig() bool`

HasResourcesVpcConfig returns a boolean if a field has been set.

### SetResourcesVpcConfigNil

`func (o *IngestersCloudResource) SetResourcesVpcConfigNil(b bool)`

 SetResourcesVpcConfigNil sets the value for ResourcesVpcConfig to be an explicit nil

### UnsetResourcesVpcConfig
`func (o *IngestersCloudResource) UnsetResourcesVpcConfig()`

UnsetResourcesVpcConfig ensures that no value is present for ResourcesVpcConfig, not even an explicit nil
### GetRestrictPublicBuckets

`func (o *IngestersCloudResource) GetRestrictPublicBuckets() bool`

GetRestrictPublicBuckets returns the RestrictPublicBuckets field if non-nil, zero value otherwise.

### GetRestrictPublicBucketsOk

`func (o *IngestersCloudResource) GetRestrictPublicBucketsOk() (*bool, bool)`

GetRestrictPublicBucketsOk returns a tuple with the RestrictPublicBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictPublicBuckets

`func (o *IngestersCloudResource) SetRestrictPublicBuckets(v bool)`

SetRestrictPublicBuckets sets RestrictPublicBuckets field to given value.

### HasRestrictPublicBuckets

`func (o *IngestersCloudResource) HasRestrictPublicBuckets() bool`

HasRestrictPublicBuckets returns a boolean if a field has been set.

### GetScheme

`func (o *IngestersCloudResource) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *IngestersCloudResource) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *IngestersCloudResource) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *IngestersCloudResource) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *IngestersCloudResource) GetSecurityGroups() interface{}`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *IngestersCloudResource) GetSecurityGroupsOk() (*interface{}, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *IngestersCloudResource) SetSecurityGroups(v interface{})`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *IngestersCloudResource) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *IngestersCloudResource) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *IngestersCloudResource) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetServiceName

`func (o *IngestersCloudResource) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *IngestersCloudResource) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *IngestersCloudResource) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *IngestersCloudResource) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetStorageAccountName

`func (o *IngestersCloudResource) GetStorageAccountName() string`

GetStorageAccountName returns the StorageAccountName field if non-nil, zero value otherwise.

### GetStorageAccountNameOk

`func (o *IngestersCloudResource) GetStorageAccountNameOk() (*string, bool)`

GetStorageAccountNameOk returns a tuple with the StorageAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccountName

`func (o *IngestersCloudResource) SetStorageAccountName(v string)`

SetStorageAccountName sets StorageAccountName field to given value.

### HasStorageAccountName

`func (o *IngestersCloudResource) HasStorageAccountName() bool`

HasStorageAccountName returns a boolean if a field has been set.

### GetTags

`func (o *IngestersCloudResource) GetTags() interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IngestersCloudResource) GetTagsOk() (*interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IngestersCloudResource) SetTags(v interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *IngestersCloudResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *IngestersCloudResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *IngestersCloudResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTargetGroupArn

`func (o *IngestersCloudResource) GetTargetGroupArn() string`

GetTargetGroupArn returns the TargetGroupArn field if non-nil, zero value otherwise.

### GetTargetGroupArnOk

`func (o *IngestersCloudResource) GetTargetGroupArnOk() (*string, bool)`

GetTargetGroupArnOk returns a tuple with the TargetGroupArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupArn

`func (o *IngestersCloudResource) SetTargetGroupArn(v string)`

SetTargetGroupArn sets TargetGroupArn field to given value.

### HasTargetGroupArn

`func (o *IngestersCloudResource) HasTargetGroupArn() bool`

HasTargetGroupArn returns a boolean if a field has been set.

### GetTargetHealthDescriptions

`func (o *IngestersCloudResource) GetTargetHealthDescriptions() interface{}`

GetTargetHealthDescriptions returns the TargetHealthDescriptions field if non-nil, zero value otherwise.

### GetTargetHealthDescriptionsOk

`func (o *IngestersCloudResource) GetTargetHealthDescriptionsOk() (*interface{}, bool)`

GetTargetHealthDescriptionsOk returns a tuple with the TargetHealthDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetHealthDescriptions

`func (o *IngestersCloudResource) SetTargetHealthDescriptions(v interface{})`

SetTargetHealthDescriptions sets TargetHealthDescriptions field to given value.

### HasTargetHealthDescriptions

`func (o *IngestersCloudResource) HasTargetHealthDescriptions() bool`

HasTargetHealthDescriptions returns a boolean if a field has been set.

### SetTargetHealthDescriptionsNil

`func (o *IngestersCloudResource) SetTargetHealthDescriptionsNil(b bool)`

 SetTargetHealthDescriptionsNil sets the value for TargetHealthDescriptions to be an explicit nil

### UnsetTargetHealthDescriptions
`func (o *IngestersCloudResource) UnsetTargetHealthDescriptions()`

UnsetTargetHealthDescriptions ensures that no value is present for TargetHealthDescriptions, not even an explicit nil
### GetTaskArn

`func (o *IngestersCloudResource) GetTaskArn() string`

GetTaskArn returns the TaskArn field if non-nil, zero value otherwise.

### GetTaskArnOk

`func (o *IngestersCloudResource) GetTaskArnOk() (*string, bool)`

GetTaskArnOk returns a tuple with the TaskArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskArn

`func (o *IngestersCloudResource) SetTaskArn(v string)`

SetTaskArn sets TaskArn field to given value.

### HasTaskArn

`func (o *IngestersCloudResource) HasTaskArn() bool`

HasTaskArn returns a boolean if a field has been set.

### GetTaskDefinition

`func (o *IngestersCloudResource) GetTaskDefinition() interface{}`

GetTaskDefinition returns the TaskDefinition field if non-nil, zero value otherwise.

### GetTaskDefinitionOk

`func (o *IngestersCloudResource) GetTaskDefinitionOk() (*interface{}, bool)`

GetTaskDefinitionOk returns a tuple with the TaskDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinition

`func (o *IngestersCloudResource) SetTaskDefinition(v interface{})`

SetTaskDefinition sets TaskDefinition field to given value.

### HasTaskDefinition

`func (o *IngestersCloudResource) HasTaskDefinition() bool`

HasTaskDefinition returns a boolean if a field has been set.

### SetTaskDefinitionNil

`func (o *IngestersCloudResource) SetTaskDefinitionNil(b bool)`

 SetTaskDefinitionNil sets the value for TaskDefinition to be an explicit nil

### UnsetTaskDefinition
`func (o *IngestersCloudResource) UnsetTaskDefinition()`

UnsetTaskDefinition ensures that no value is present for TaskDefinition, not even an explicit nil
### GetTaskDefinitionArn

`func (o *IngestersCloudResource) GetTaskDefinitionArn() string`

GetTaskDefinitionArn returns the TaskDefinitionArn field if non-nil, zero value otherwise.

### GetTaskDefinitionArnOk

`func (o *IngestersCloudResource) GetTaskDefinitionArnOk() (*string, bool)`

GetTaskDefinitionArnOk returns a tuple with the TaskDefinitionArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefinitionArn

`func (o *IngestersCloudResource) SetTaskDefinitionArn(v string)`

SetTaskDefinitionArn sets TaskDefinitionArn field to given value.

### HasTaskDefinitionArn

`func (o *IngestersCloudResource) HasTaskDefinitionArn() bool`

HasTaskDefinitionArn returns a boolean if a field has been set.

### GetUserGroups

`func (o *IngestersCloudResource) GetUserGroups() interface{}`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *IngestersCloudResource) GetUserGroupsOk() (*interface{}, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *IngestersCloudResource) SetUserGroups(v interface{})`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *IngestersCloudResource) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.

### SetUserGroupsNil

`func (o *IngestersCloudResource) SetUserGroupsNil(b bool)`

 SetUserGroupsNil sets the value for UserGroups to be an explicit nil

### UnsetUserGroups
`func (o *IngestersCloudResource) UnsetUserGroups()`

UnsetUserGroups ensures that no value is present for UserGroups, not even an explicit nil
### GetUserId

`func (o *IngestersCloudResource) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *IngestersCloudResource) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *IngestersCloudResource) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *IngestersCloudResource) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsers

`func (o *IngestersCloudResource) GetUsers() interface{}`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *IngestersCloudResource) GetUsersOk() (*interface{}, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *IngestersCloudResource) SetUsers(v interface{})`

SetUsers sets Users field to given value.

### HasUsers

`func (o *IngestersCloudResource) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *IngestersCloudResource) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *IngestersCloudResource) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil
### GetVpcId

`func (o *IngestersCloudResource) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *IngestersCloudResource) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *IngestersCloudResource) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *IngestersCloudResource) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetVpcOptions

`func (o *IngestersCloudResource) GetVpcOptions() interface{}`

GetVpcOptions returns the VpcOptions field if non-nil, zero value otherwise.

### GetVpcOptionsOk

`func (o *IngestersCloudResource) GetVpcOptionsOk() (*interface{}, bool)`

GetVpcOptionsOk returns a tuple with the VpcOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcOptions

`func (o *IngestersCloudResource) SetVpcOptions(v interface{})`

SetVpcOptions sets VpcOptions field to given value.

### HasVpcOptions

`func (o *IngestersCloudResource) HasVpcOptions() bool`

HasVpcOptions returns a boolean if a field has been set.

### SetVpcOptionsNil

`func (o *IngestersCloudResource) SetVpcOptionsNil(b bool)`

 SetVpcOptionsNil sets the value for VpcOptions to be an explicit nil

### UnsetVpcOptions
`func (o *IngestersCloudResource) UnsetVpcOptions()`

UnsetVpcOptions ensures that no value is present for VpcOptions, not even an explicit nil
### GetVpcSecurityGroupIds

`func (o *IngestersCloudResource) GetVpcSecurityGroupIds() interface{}`

GetVpcSecurityGroupIds returns the VpcSecurityGroupIds field if non-nil, zero value otherwise.

### GetVpcSecurityGroupIdsOk

`func (o *IngestersCloudResource) GetVpcSecurityGroupIdsOk() (*interface{}, bool)`

GetVpcSecurityGroupIdsOk returns a tuple with the VpcSecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcSecurityGroupIds

`func (o *IngestersCloudResource) SetVpcSecurityGroupIds(v interface{})`

SetVpcSecurityGroupIds sets VpcSecurityGroupIds field to given value.

### HasVpcSecurityGroupIds

`func (o *IngestersCloudResource) HasVpcSecurityGroupIds() bool`

HasVpcSecurityGroupIds returns a boolean if a field has been set.

### SetVpcSecurityGroupIdsNil

`func (o *IngestersCloudResource) SetVpcSecurityGroupIdsNil(b bool)`

 SetVpcSecurityGroupIdsNil sets the value for VpcSecurityGroupIds to be an explicit nil

### UnsetVpcSecurityGroupIds
`func (o *IngestersCloudResource) UnsetVpcSecurityGroupIds()`

UnsetVpcSecurityGroupIds ensures that no value is present for VpcSecurityGroupIds, not even an explicit nil
### GetVpcSecurityGroups

`func (o *IngestersCloudResource) GetVpcSecurityGroups() interface{}`

GetVpcSecurityGroups returns the VpcSecurityGroups field if non-nil, zero value otherwise.

### GetVpcSecurityGroupsOk

`func (o *IngestersCloudResource) GetVpcSecurityGroupsOk() (*interface{}, bool)`

GetVpcSecurityGroupsOk returns a tuple with the VpcSecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcSecurityGroups

`func (o *IngestersCloudResource) SetVpcSecurityGroups(v interface{})`

SetVpcSecurityGroups sets VpcSecurityGroups field to given value.

### HasVpcSecurityGroups

`func (o *IngestersCloudResource) HasVpcSecurityGroups() bool`

HasVpcSecurityGroups returns a boolean if a field has been set.

### SetVpcSecurityGroupsNil

`func (o *IngestersCloudResource) SetVpcSecurityGroupsNil(b bool)`

 SetVpcSecurityGroupsNil sets the value for VpcSecurityGroups to be an explicit nil

### UnsetVpcSecurityGroups
`func (o *IngestersCloudResource) UnsetVpcSecurityGroups()`

UnsetVpcSecurityGroups ensures that no value is present for VpcSecurityGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


