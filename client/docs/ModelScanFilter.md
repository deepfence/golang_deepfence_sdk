# ModelScanFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudAccountScanFilter** | [**ReportersContainsFilter**](ReportersContainsFilter.md) |  | 
**ContainerScanFilter** | [**ReportersContainsFilter**](ReportersContainsFilter.md) |  | 
**HostScanFilter** | [**ReportersContainsFilter**](ReportersContainsFilter.md) |  | 
**ImageScanFilter** | [**ReportersContainsFilter**](ReportersContainsFilter.md) |  | 
**KubernetesClusterScanFilter** | [**ReportersContainsFilter**](ReportersContainsFilter.md) |  | 

## Methods

### NewModelScanFilter

`func NewModelScanFilter(cloudAccountScanFilter ReportersContainsFilter, containerScanFilter ReportersContainsFilter, hostScanFilter ReportersContainsFilter, imageScanFilter ReportersContainsFilter, kubernetesClusterScanFilter ReportersContainsFilter, ) *ModelScanFilter`

NewModelScanFilter instantiates a new ModelScanFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelScanFilterWithDefaults

`func NewModelScanFilterWithDefaults() *ModelScanFilter`

NewModelScanFilterWithDefaults instantiates a new ModelScanFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudAccountScanFilter

`func (o *ModelScanFilter) GetCloudAccountScanFilter() ReportersContainsFilter`

GetCloudAccountScanFilter returns the CloudAccountScanFilter field if non-nil, zero value otherwise.

### GetCloudAccountScanFilterOk

`func (o *ModelScanFilter) GetCloudAccountScanFilterOk() (*ReportersContainsFilter, bool)`

GetCloudAccountScanFilterOk returns a tuple with the CloudAccountScanFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccountScanFilter

`func (o *ModelScanFilter) SetCloudAccountScanFilter(v ReportersContainsFilter)`

SetCloudAccountScanFilter sets CloudAccountScanFilter field to given value.


### GetContainerScanFilter

`func (o *ModelScanFilter) GetContainerScanFilter() ReportersContainsFilter`

GetContainerScanFilter returns the ContainerScanFilter field if non-nil, zero value otherwise.

### GetContainerScanFilterOk

`func (o *ModelScanFilter) GetContainerScanFilterOk() (*ReportersContainsFilter, bool)`

GetContainerScanFilterOk returns a tuple with the ContainerScanFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScanFilter

`func (o *ModelScanFilter) SetContainerScanFilter(v ReportersContainsFilter)`

SetContainerScanFilter sets ContainerScanFilter field to given value.


### GetHostScanFilter

`func (o *ModelScanFilter) GetHostScanFilter() ReportersContainsFilter`

GetHostScanFilter returns the HostScanFilter field if non-nil, zero value otherwise.

### GetHostScanFilterOk

`func (o *ModelScanFilter) GetHostScanFilterOk() (*ReportersContainsFilter, bool)`

GetHostScanFilterOk returns a tuple with the HostScanFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostScanFilter

`func (o *ModelScanFilter) SetHostScanFilter(v ReportersContainsFilter)`

SetHostScanFilter sets HostScanFilter field to given value.


### GetImageScanFilter

`func (o *ModelScanFilter) GetImageScanFilter() ReportersContainsFilter`

GetImageScanFilter returns the ImageScanFilter field if non-nil, zero value otherwise.

### GetImageScanFilterOk

`func (o *ModelScanFilter) GetImageScanFilterOk() (*ReportersContainsFilter, bool)`

GetImageScanFilterOk returns a tuple with the ImageScanFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageScanFilter

`func (o *ModelScanFilter) SetImageScanFilter(v ReportersContainsFilter)`

SetImageScanFilter sets ImageScanFilter field to given value.


### GetKubernetesClusterScanFilter

`func (o *ModelScanFilter) GetKubernetesClusterScanFilter() ReportersContainsFilter`

GetKubernetesClusterScanFilter returns the KubernetesClusterScanFilter field if non-nil, zero value otherwise.

### GetKubernetesClusterScanFilterOk

`func (o *ModelScanFilter) GetKubernetesClusterScanFilterOk() (*ReportersContainsFilter, bool)`

GetKubernetesClusterScanFilterOk returns a tuple with the KubernetesClusterScanFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesClusterScanFilter

`func (o *ModelScanFilter) SetKubernetesClusterScanFilter(v ReportersContainsFilter)`

SetKubernetesClusterScanFilter sets KubernetesClusterScanFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


