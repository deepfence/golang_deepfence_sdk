# ModelCloudNodeControlReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** |  | 
**ComplianceType** | **string** |  | 
**NodeId** | Pointer to **string** |  | [optional] 

## Methods

### NewModelCloudNodeControlReq

`func NewModelCloudNodeControlReq(cloudProvider string, complianceType string, ) *ModelCloudNodeControlReq`

NewModelCloudNodeControlReq instantiates a new ModelCloudNodeControlReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCloudNodeControlReqWithDefaults

`func NewModelCloudNodeControlReqWithDefaults() *ModelCloudNodeControlReq`

NewModelCloudNodeControlReqWithDefaults instantiates a new ModelCloudNodeControlReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *ModelCloudNodeControlReq) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ModelCloudNodeControlReq) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ModelCloudNodeControlReq) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetComplianceType

`func (o *ModelCloudNodeControlReq) GetComplianceType() string`

GetComplianceType returns the ComplianceType field if non-nil, zero value otherwise.

### GetComplianceTypeOk

`func (o *ModelCloudNodeControlReq) GetComplianceTypeOk() (*string, bool)`

GetComplianceTypeOk returns a tuple with the ComplianceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceType

`func (o *ModelCloudNodeControlReq) SetComplianceType(v string)`

SetComplianceType sets ComplianceType field to given value.


### GetNodeId

`func (o *ModelCloudNodeControlReq) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelCloudNodeControlReq) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelCloudNodeControlReq) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ModelCloudNodeControlReq) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


