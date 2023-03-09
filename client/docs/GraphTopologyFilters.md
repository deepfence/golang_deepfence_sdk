# GraphTopologyFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudFilter** | **[]string** |  | 
**FieldFilters** | [**ReportersFieldsFilters**](ReportersFieldsFilters.md) |  | 
**HostFilter** | **[]string** |  | 
**KubernetesFilter** | **[]string** |  | 
**PodFilter** | **[]string** |  | 
**RegionFilter** | **[]string** |  | 
**ServiceFilter** | **[]string** |  | 

## Methods

### NewGraphTopologyFilters

`func NewGraphTopologyFilters(cloudFilter []string, fieldFilters ReportersFieldsFilters, hostFilter []string, kubernetesFilter []string, podFilter []string, regionFilter []string, serviceFilter []string, ) *GraphTopologyFilters`

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
### GetServiceFilter

`func (o *GraphTopologyFilters) GetServiceFilter() []string`

GetServiceFilter returns the ServiceFilter field if non-nil, zero value otherwise.

### GetServiceFilterOk

`func (o *GraphTopologyFilters) GetServiceFilterOk() (*[]string, bool)`

GetServiceFilterOk returns a tuple with the ServiceFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFilter

`func (o *GraphTopologyFilters) SetServiceFilter(v []string)`

SetServiceFilter sets ServiceFilter field to given value.


### SetServiceFilterNil

`func (o *GraphTopologyFilters) SetServiceFilterNil(b bool)`

 SetServiceFilterNil sets the value for ServiceFilter to be an explicit nil

### UnsetServiceFilter
`func (o *GraphTopologyFilters) UnsetServiceFilter()`

UnsetServiceFilter ensures that no value is present for ServiceFilter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


