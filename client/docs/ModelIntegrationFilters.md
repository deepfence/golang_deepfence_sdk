# ModelIntegrationFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldsFilters** | Pointer to [**ReportersFieldsFilters**](ReportersFieldsFilters.md) |  | [optional] 
**NodeIds** | [**[]ModelNodeIdentifier**](ModelNodeIdentifier.md) |  | 

## Methods

### NewModelIntegrationFilters

`func NewModelIntegrationFilters(nodeIds []ModelNodeIdentifier, ) *ModelIntegrationFilters`

NewModelIntegrationFilters instantiates a new ModelIntegrationFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelIntegrationFiltersWithDefaults

`func NewModelIntegrationFiltersWithDefaults() *ModelIntegrationFilters`

NewModelIntegrationFiltersWithDefaults instantiates a new ModelIntegrationFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldsFilters

`func (o *ModelIntegrationFilters) GetFieldsFilters() ReportersFieldsFilters`

GetFieldsFilters returns the FieldsFilters field if non-nil, zero value otherwise.

### GetFieldsFiltersOk

`func (o *ModelIntegrationFilters) GetFieldsFiltersOk() (*ReportersFieldsFilters, bool)`

GetFieldsFiltersOk returns a tuple with the FieldsFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsFilters

`func (o *ModelIntegrationFilters) SetFieldsFilters(v ReportersFieldsFilters)`

SetFieldsFilters sets FieldsFilters field to given value.

### HasFieldsFilters

`func (o *ModelIntegrationFilters) HasFieldsFilters() bool`

HasFieldsFilters returns a boolean if a field has been set.

### GetNodeIds

`func (o *ModelIntegrationFilters) GetNodeIds() []ModelNodeIdentifier`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *ModelIntegrationFilters) GetNodeIdsOk() (*[]ModelNodeIdentifier, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *ModelIntegrationFilters) SetNodeIds(v []ModelNodeIdentifier)`

SetNodeIds sets NodeIds field to given value.


### SetNodeIdsNil

`func (o *ModelIntegrationFilters) SetNodeIdsNil(b bool)`

 SetNodeIdsNil sets the value for NodeIds to be an explicit nil

### UnsetNodeIds
`func (o *ModelIntegrationFilters) UnsetNodeIds()`

UnsetNodeIds ensures that no value is present for NodeIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


