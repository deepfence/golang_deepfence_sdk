# GraphThreatFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsFilter** | **map[string]interface{}** |  | 
**AzureFilter** | **map[string]interface{}** |  | 
**CloudResourceOnly** | **bool** |  | 
**GcpFilter** | **map[string]interface{}** |  | 
**Type** | **string** |  | 

## Methods

### NewGraphThreatFilters

`func NewGraphThreatFilters(awsFilter map[string]interface{}, azureFilter map[string]interface{}, cloudResourceOnly bool, gcpFilter map[string]interface{}, type_ string, ) *GraphThreatFilters`

NewGraphThreatFilters instantiates a new GraphThreatFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphThreatFiltersWithDefaults

`func NewGraphThreatFiltersWithDefaults() *GraphThreatFilters`

NewGraphThreatFiltersWithDefaults instantiates a new GraphThreatFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsFilter

`func (o *GraphThreatFilters) GetAwsFilter() map[string]interface{}`

GetAwsFilter returns the AwsFilter field if non-nil, zero value otherwise.

### GetAwsFilterOk

`func (o *GraphThreatFilters) GetAwsFilterOk() (*map[string]interface{}, bool)`

GetAwsFilterOk returns a tuple with the AwsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsFilter

`func (o *GraphThreatFilters) SetAwsFilter(v map[string]interface{})`

SetAwsFilter sets AwsFilter field to given value.


### GetAzureFilter

`func (o *GraphThreatFilters) GetAzureFilter() map[string]interface{}`

GetAzureFilter returns the AzureFilter field if non-nil, zero value otherwise.

### GetAzureFilterOk

`func (o *GraphThreatFilters) GetAzureFilterOk() (*map[string]interface{}, bool)`

GetAzureFilterOk returns a tuple with the AzureFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureFilter

`func (o *GraphThreatFilters) SetAzureFilter(v map[string]interface{})`

SetAzureFilter sets AzureFilter field to given value.


### GetCloudResourceOnly

`func (o *GraphThreatFilters) GetCloudResourceOnly() bool`

GetCloudResourceOnly returns the CloudResourceOnly field if non-nil, zero value otherwise.

### GetCloudResourceOnlyOk

`func (o *GraphThreatFilters) GetCloudResourceOnlyOk() (*bool, bool)`

GetCloudResourceOnlyOk returns a tuple with the CloudResourceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudResourceOnly

`func (o *GraphThreatFilters) SetCloudResourceOnly(v bool)`

SetCloudResourceOnly sets CloudResourceOnly field to given value.


### GetGcpFilter

`func (o *GraphThreatFilters) GetGcpFilter() map[string]interface{}`

GetGcpFilter returns the GcpFilter field if non-nil, zero value otherwise.

### GetGcpFilterOk

`func (o *GraphThreatFilters) GetGcpFilterOk() (*map[string]interface{}, bool)`

GetGcpFilterOk returns a tuple with the GcpFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpFilter

`func (o *GraphThreatFilters) SetGcpFilter(v map[string]interface{})`

SetGcpFilter sets GcpFilter field to given value.


### GetType

`func (o *GraphThreatFilters) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphThreatFilters) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphThreatFilters) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


