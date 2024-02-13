# ModelRegistryAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerImages** | [**[]ModelContainerImage**](ModelContainerImage.md) |  | 
**Name** | **string** |  | 
**NodeId** | **string** |  | 
**RegistryType** | **string** |  | 
**Syncing** | **bool** |  | 

## Methods

### NewModelRegistryAccount

`func NewModelRegistryAccount(containerImages []ModelContainerImage, name string, nodeId string, registryType string, syncing bool, ) *ModelRegistryAccount`

NewModelRegistryAccount instantiates a new ModelRegistryAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelRegistryAccountWithDefaults

`func NewModelRegistryAccountWithDefaults() *ModelRegistryAccount`

NewModelRegistryAccountWithDefaults instantiates a new ModelRegistryAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerImages

`func (o *ModelRegistryAccount) GetContainerImages() []ModelContainerImage`

GetContainerImages returns the ContainerImages field if non-nil, zero value otherwise.

### GetContainerImagesOk

`func (o *ModelRegistryAccount) GetContainerImagesOk() (*[]ModelContainerImage, bool)`

GetContainerImagesOk returns a tuple with the ContainerImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerImages

`func (o *ModelRegistryAccount) SetContainerImages(v []ModelContainerImage)`

SetContainerImages sets ContainerImages field to given value.


### SetContainerImagesNil

`func (o *ModelRegistryAccount) SetContainerImagesNil(b bool)`

 SetContainerImagesNil sets the value for ContainerImages to be an explicit nil

### UnsetContainerImages
`func (o *ModelRegistryAccount) UnsetContainerImages()`

UnsetContainerImages ensures that no value is present for ContainerImages, not even an explicit nil
### GetName

`func (o *ModelRegistryAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelRegistryAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelRegistryAccount) SetName(v string)`

SetName sets Name field to given value.


### GetNodeId

`func (o *ModelRegistryAccount) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ModelRegistryAccount) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ModelRegistryAccount) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetRegistryType

`func (o *ModelRegistryAccount) GetRegistryType() string`

GetRegistryType returns the RegistryType field if non-nil, zero value otherwise.

### GetRegistryTypeOk

`func (o *ModelRegistryAccount) GetRegistryTypeOk() (*string, bool)`

GetRegistryTypeOk returns a tuple with the RegistryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryType

`func (o *ModelRegistryAccount) SetRegistryType(v string)`

SetRegistryType sets RegistryType field to given value.


### GetSyncing

`func (o *ModelRegistryAccount) GetSyncing() bool`

GetSyncing returns the Syncing field if non-nil, zero value otherwise.

### GetSyncingOk

`func (o *ModelRegistryAccount) GetSyncingOk() (*bool, bool)`

GetSyncingOk returns a tuple with the Syncing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncing

`func (o *ModelRegistryAccount) SetSyncing(v bool)`

SetSyncing sets Syncing field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


