# ModelRegisterLicenseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**LicenseKey** | **string** |  | 

## Methods

### NewModelRegisterLicenseRequest

`func NewModelRegisterLicenseRequest(licenseKey string, ) *ModelRegisterLicenseRequest`

NewModelRegisterLicenseRequest instantiates a new ModelRegisterLicenseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelRegisterLicenseRequestWithDefaults

`func NewModelRegisterLicenseRequestWithDefaults() *ModelRegisterLicenseRequest`

NewModelRegisterLicenseRequestWithDefaults instantiates a new ModelRegisterLicenseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ModelRegisterLicenseRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ModelRegisterLicenseRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ModelRegisterLicenseRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ModelRegisterLicenseRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLicenseKey

`func (o *ModelRegisterLicenseRequest) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *ModelRegisterLicenseRequest) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *ModelRegisterLicenseRequest) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


