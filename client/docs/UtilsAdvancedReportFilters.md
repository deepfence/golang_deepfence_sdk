# UtilsAdvancedReportFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerName** | Pointer to **[]string** |  | [optional] 
**HostName** | Pointer to **[]string** |  | [optional] 
**ImageName** | Pointer to **[]string** |  | [optional] 
**KubernetesClusterName** | Pointer to **[]string** |  | [optional] 
**Masked** | Pointer to **[]bool** |  | [optional] 
**NodeId** | Pointer to **[]string** |  | [optional] 
**PodName** | Pointer to **[]string** |  | [optional] 
**ScanStatus** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUtilsAdvancedReportFilters

`func NewUtilsAdvancedReportFilters() *UtilsAdvancedReportFilters`

NewUtilsAdvancedReportFilters instantiates a new UtilsAdvancedReportFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilsAdvancedReportFiltersWithDefaults

`func NewUtilsAdvancedReportFiltersWithDefaults() *UtilsAdvancedReportFilters`

NewUtilsAdvancedReportFiltersWithDefaults instantiates a new UtilsAdvancedReportFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerName

`func (o *UtilsAdvancedReportFilters) GetContainerName() []string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *UtilsAdvancedReportFilters) GetContainerNameOk() (*[]string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *UtilsAdvancedReportFilters) SetContainerName(v []string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *UtilsAdvancedReportFilters) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetHostName

`func (o *UtilsAdvancedReportFilters) GetHostName() []string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *UtilsAdvancedReportFilters) GetHostNameOk() (*[]string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *UtilsAdvancedReportFilters) SetHostName(v []string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *UtilsAdvancedReportFilters) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetImageName

`func (o *UtilsAdvancedReportFilters) GetImageName() []string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *UtilsAdvancedReportFilters) GetImageNameOk() (*[]string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *UtilsAdvancedReportFilters) SetImageName(v []string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *UtilsAdvancedReportFilters) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetKubernetesClusterName

`func (o *UtilsAdvancedReportFilters) GetKubernetesClusterName() []string`

GetKubernetesClusterName returns the KubernetesClusterName field if non-nil, zero value otherwise.

### GetKubernetesClusterNameOk

`func (o *UtilsAdvancedReportFilters) GetKubernetesClusterNameOk() (*[]string, bool)`

GetKubernetesClusterNameOk returns a tuple with the KubernetesClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesClusterName

`func (o *UtilsAdvancedReportFilters) SetKubernetesClusterName(v []string)`

SetKubernetesClusterName sets KubernetesClusterName field to given value.

### HasKubernetesClusterName

`func (o *UtilsAdvancedReportFilters) HasKubernetesClusterName() bool`

HasKubernetesClusterName returns a boolean if a field has been set.

### GetMasked

`func (o *UtilsAdvancedReportFilters) GetMasked() []bool`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *UtilsAdvancedReportFilters) GetMaskedOk() (*[]bool, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *UtilsAdvancedReportFilters) SetMasked(v []bool)`

SetMasked sets Masked field to given value.

### HasMasked

`func (o *UtilsAdvancedReportFilters) HasMasked() bool`

HasMasked returns a boolean if a field has been set.

### GetNodeId

`func (o *UtilsAdvancedReportFilters) GetNodeId() []string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *UtilsAdvancedReportFilters) GetNodeIdOk() (*[]string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *UtilsAdvancedReportFilters) SetNodeId(v []string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *UtilsAdvancedReportFilters) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetPodName

`func (o *UtilsAdvancedReportFilters) GetPodName() []string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *UtilsAdvancedReportFilters) GetPodNameOk() (*[]string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *UtilsAdvancedReportFilters) SetPodName(v []string)`

SetPodName sets PodName field to given value.

### HasPodName

`func (o *UtilsAdvancedReportFilters) HasPodName() bool`

HasPodName returns a boolean if a field has been set.

### GetScanStatus

`func (o *UtilsAdvancedReportFilters) GetScanStatus() []string`

GetScanStatus returns the ScanStatus field if non-nil, zero value otherwise.

### GetScanStatusOk

`func (o *UtilsAdvancedReportFilters) GetScanStatusOk() (*[]string, bool)`

GetScanStatusOk returns a tuple with the ScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStatus

`func (o *UtilsAdvancedReportFilters) SetScanStatus(v []string)`

SetScanStatus sets ScanStatus field to given value.

### HasScanStatus

`func (o *UtilsAdvancedReportFilters) HasScanStatus() bool`

HasScanStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


