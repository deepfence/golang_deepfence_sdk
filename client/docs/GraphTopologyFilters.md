# GraphTopologyFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudFilter** | **[]string** |  | 
**ContainerFilter** | **[]string** |  | 
**FieldFilters** | [**ReportersFieldsFilters**](ReportersFieldsFilters.md) |  | 
**HostFilter** | **[]string** |  | 
**KubernetesFilter** | **[]string** |  | 
**PodFilter** | **[]string** |  | 
**RegionFilter** | **[]string** |  | 
**SkipConnections** | **bool** |  | 

## Methods

### NewGraphTopologyFilters

`func NewGraphTopologyFilters(cloudFilter []string, containerFilter []string, fieldFilters ReportersFieldsFilters, hostFilter []string, kubernetesFilter []string, podFilter []string, regionFilter []string, skipConnections bool, ) *GraphTopologyFilters`

NewGraphTopologyFilters instantiates a new GraphTopologyFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphTopologyFiltersWithDefaults

`func NewGraphTopologyFiltersWithDefaults() *GraphTopologyFilters`

NewGraphTopologyFiltersWithDefaults instantiates a new GraphTopologyFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudFilter

`func (o *GraphTopologyFilters) GetCloudFilter() []string`

GetCloudFilter returns the CloudFilter field if non-nil, zero value otherwise.

### GetCloudFilterOk

`func (o *GraphTopologyFilters) GetCloudFilterOk() (*[]string, bool)`

GetCloudFilterOk returns a tuple with the CloudFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudFilter

`func (o *GraphTopologyFilters) SetCloudFilter(v []string)`

SetCloudFilter sets CloudFilter field to given value.


### SetCloudFilterNil

`func (o *GraphTopologyFilters) SetCloudFilterNil(b bool)`

 SetCloudFilterNil sets the value for CloudFilter to be an explicit nil

### UnsetCloudFilter
`func (o *GraphTopologyFilters) UnsetCloudFilter()`

UnsetCloudFilter ensures that no value is present for CloudFilter, not even an explicit nil
### GetContainerFilter

`func (o *GraphTopologyFilters) GetContainerFilter() []string`

GetContainerFilter returns the ContainerFilter field if non-nil, zero value otherwise.

### GetContainerFilterOk

`func (o *GraphTopologyFilters) GetContainerFilterOk() (*[]string, bool)`

GetContainerFilterOk returns a tuple with the ContainerFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerFilter

`func (o *GraphTopologyFilters) SetContainerFilter(v []string)`

SetContainerFilter sets ContainerFilter field to given value.


### SetContainerFilterNil

`func (o *GraphTopologyFilters) SetContainerFilterNil(b bool)`

 SetContainerFilterNil sets the value for ContainerFilter to be an explicit nil

### UnsetContainerFilter
`func (o *GraphTopologyFilters) UnsetContainerFilter()`

UnsetContainerFilter ensures that no value is present for ContainerFilter, not even an explicit nil
### GetFieldFilters

`func (o *GraphTopologyFilters) GetFieldFilters() ReportersFieldsFilters`

GetFieldFilters returns the FieldFilters field if non-nil, zero value otherwise.

### GetFieldFiltersOk

`func (o *GraphTopologyFilters) GetFieldFiltersOk() (*ReportersFieldsFilters, bool)`

GetFieldFiltersOk returns a tuple with the FieldFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldFilters

`func (o *GraphTopologyFilters) SetFieldFilters(v ReportersFieldsFilters)`

SetFieldFilters sets FieldFilters field to given value.


### GetHostFilter

`func (o *GraphTopologyFilters) GetHostFilter() []string`

GetHostFilter returns the HostFilter field if non-nil, zero value otherwise.

### GetHostFilterOk

`func (o *GraphTopologyFilters) GetHostFilterOk() (*[]string, bool)`

GetHostFilterOk returns a tuple with the HostFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostFilter

`func (o *GraphTopologyFilters) SetHostFilter(v []string)`

SetHostFilter sets HostFilter field to given value.


### SetHostFilterNil

`func (o *GraphTopologyFilters) SetHostFilterNil(b bool)`

 SetHostFilterNil sets the value for HostFilter to be an explicit nil

### UnsetHostFilter
`func (o *GraphTopologyFilters) UnsetHostFilter()`

UnsetHostFilter ensures that no value is present for HostFilter, not even an explicit nil
### GetKubernetesFilter

`func (o *GraphTopologyFilters) GetKubernetesFilter() []string`

GetKubernetesFilter returns the KubernetesFilter field if non-nil, zero value otherwise.

### GetKubernetesFilterOk

`func (o *GraphTopologyFilters) GetKubernetesFilterOk() (*[]string, bool)`

GetKubernetesFilterOk returns a tuple with the KubernetesFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesFilter

`func (o *GraphTopologyFilters) SetKubernetesFilter(v []string)`

SetKubernetesFilter sets KubernetesFilter field to given value.


### SetKubernetesFilterNil

`func (o *GraphTopologyFilters) SetKubernetesFilterNil(b bool)`

 SetKubernetesFilterNil sets the value for KubernetesFilter to be an explicit nil

### UnsetKubernetesFilter
`func (o *GraphTopologyFilters) UnsetKubernetesFilter()`

UnsetKubernetesFilter ensures that no value is present for KubernetesFilter, not even an explicit nil
### GetPodFilter

`func (o *GraphTopologyFilters) GetPodFilter() []string`

GetPodFilter returns the PodFilter field if non-nil, zero value otherwise.

### GetPodFilterOk

`func (o *GraphTopologyFilters) GetPodFilterOk() (*[]string, bool)`

GetPodFilterOk returns a tuple with the PodFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodFilter

`func (o *GraphTopologyFilters) SetPodFilter(v []string)`

SetPodFilter sets PodFilter field to given value.


### SetPodFilterNil

`func (o *GraphTopologyFilters) SetPodFilterNil(b bool)`

 SetPodFilterNil sets the value for PodFilter to be an explicit nil

### UnsetPodFilter
`func (o *GraphTopologyFilters) UnsetPodFilter()`

UnsetPodFilter ensures that no value is present for PodFilter, not even an explicit nil
### GetRegionFilter

`func (o *GraphTopologyFilters) GetRegionFilter() []string`

GetRegionFilter returns the RegionFilter field if non-nil, zero value otherwise.

### GetRegionFilterOk

`func (o *GraphTopologyFilters) GetRegionFilterOk() (*[]string, bool)`

GetRegionFilterOk returns a tuple with the RegionFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionFilter

`func (o *GraphTopologyFilters) SetRegionFilter(v []string)`

SetRegionFilter sets RegionFilter field to given value.


### SetRegionFilterNil

`func (o *GraphTopologyFilters) SetRegionFilterNil(b bool)`

 SetRegionFilterNil sets the value for RegionFilter to be an explicit nil

### UnsetRegionFilter
`func (o *GraphTopologyFilters) UnsetRegionFilter()`

UnsetRegionFilter ensures that no value is present for RegionFilter, not even an explicit nil
### GetSkipConnections

`func (o *GraphTopologyFilters) GetSkipConnections() bool`

GetSkipConnections returns the SkipConnections field if non-nil, zero value otherwise.

### GetSkipConnectionsOk

`func (o *GraphTopologyFilters) GetSkipConnectionsOk() (*bool, bool)`

GetSkipConnectionsOk returns a tuple with the SkipConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipConnections

`func (o *GraphTopologyFilters) SetSkipConnections(v bool)`

SetSkipConnections sets SkipConnections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


