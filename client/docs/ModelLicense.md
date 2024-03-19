# ModelLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentHosts** | Pointer to **int32** |  | [optional] 
**DeepfenceSupportEmail** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**LicenseEmail** | Pointer to **string** |  | [optional] 
**LicenseEmailDomain** | Pointer to **string** |  | [optional] 
**LicenseType** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**NoOfCloudAccounts** | Pointer to **int32** |  | [optional] 
**NoOfHosts** | Pointer to **int32** |  | [optional] 
**NoOfImagesInRegistry** | Pointer to **int32** |  | [optional] 
**NoOfRegistries** | Pointer to **int32** |  | [optional] 
**NotificationThresholdPercentage** | Pointer to **int32** |  | [optional] 
**NotificationThresholdUpdatedAt** | Pointer to **int32** |  | [optional] 
**RegistryCredentials** | Pointer to [**ModelRegistryCredentials**](ModelRegistryCredentials.md) |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 

## Methods

### NewModelLicense

`func NewModelLicense() *ModelLicense`

NewModelLicense instantiates a new ModelLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelLicenseWithDefaults

`func NewModelLicenseWithDefaults() *ModelLicense`

NewModelLicenseWithDefaults instantiates a new ModelLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentHosts

`func (o *ModelLicense) GetCurrentHosts() int32`

GetCurrentHosts returns the CurrentHosts field if non-nil, zero value otherwise.

### GetCurrentHostsOk

`func (o *ModelLicense) GetCurrentHostsOk() (*int32, bool)`

GetCurrentHostsOk returns a tuple with the CurrentHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentHosts

`func (o *ModelLicense) SetCurrentHosts(v int32)`

SetCurrentHosts sets CurrentHosts field to given value.

### HasCurrentHosts

`func (o *ModelLicense) HasCurrentHosts() bool`

HasCurrentHosts returns a boolean if a field has been set.

### GetDeepfenceSupportEmail

`func (o *ModelLicense) GetDeepfenceSupportEmail() string`

GetDeepfenceSupportEmail returns the DeepfenceSupportEmail field if non-nil, zero value otherwise.

### GetDeepfenceSupportEmailOk

`func (o *ModelLicense) GetDeepfenceSupportEmailOk() (*string, bool)`

GetDeepfenceSupportEmailOk returns a tuple with the DeepfenceSupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepfenceSupportEmail

`func (o *ModelLicense) SetDeepfenceSupportEmail(v string)`

SetDeepfenceSupportEmail sets DeepfenceSupportEmail field to given value.

### HasDeepfenceSupportEmail

`func (o *ModelLicense) HasDeepfenceSupportEmail() bool`

HasDeepfenceSupportEmail returns a boolean if a field has been set.

### GetDescription

`func (o *ModelLicense) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelLicense) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelLicense) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelLicense) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndDate

`func (o *ModelLicense) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ModelLicense) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ModelLicense) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ModelLicense) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetIsActive

`func (o *ModelLicense) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ModelLicense) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ModelLicense) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ModelLicense) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetKey

`func (o *ModelLicense) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ModelLicense) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ModelLicense) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ModelLicense) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLicenseEmail

`func (o *ModelLicense) GetLicenseEmail() string`

GetLicenseEmail returns the LicenseEmail field if non-nil, zero value otherwise.

### GetLicenseEmailOk

`func (o *ModelLicense) GetLicenseEmailOk() (*string, bool)`

GetLicenseEmailOk returns a tuple with the LicenseEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseEmail

`func (o *ModelLicense) SetLicenseEmail(v string)`

SetLicenseEmail sets LicenseEmail field to given value.

### HasLicenseEmail

`func (o *ModelLicense) HasLicenseEmail() bool`

HasLicenseEmail returns a boolean if a field has been set.

### GetLicenseEmailDomain

`func (o *ModelLicense) GetLicenseEmailDomain() string`

GetLicenseEmailDomain returns the LicenseEmailDomain field if non-nil, zero value otherwise.

### GetLicenseEmailDomainOk

`func (o *ModelLicense) GetLicenseEmailDomainOk() (*string, bool)`

GetLicenseEmailDomainOk returns a tuple with the LicenseEmailDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseEmailDomain

`func (o *ModelLicense) SetLicenseEmailDomain(v string)`

SetLicenseEmailDomain sets LicenseEmailDomain field to given value.

### HasLicenseEmailDomain

`func (o *ModelLicense) HasLicenseEmailDomain() bool`

HasLicenseEmailDomain returns a boolean if a field has been set.

### GetLicenseType

`func (o *ModelLicense) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *ModelLicense) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *ModelLicense) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *ModelLicense) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetMessage

`func (o *ModelLicense) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ModelLicense) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ModelLicense) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ModelLicense) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNoOfCloudAccounts

`func (o *ModelLicense) GetNoOfCloudAccounts() int32`

GetNoOfCloudAccounts returns the NoOfCloudAccounts field if non-nil, zero value otherwise.

### GetNoOfCloudAccountsOk

`func (o *ModelLicense) GetNoOfCloudAccountsOk() (*int32, bool)`

GetNoOfCloudAccountsOk returns a tuple with the NoOfCloudAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfCloudAccounts

`func (o *ModelLicense) SetNoOfCloudAccounts(v int32)`

SetNoOfCloudAccounts sets NoOfCloudAccounts field to given value.

### HasNoOfCloudAccounts

`func (o *ModelLicense) HasNoOfCloudAccounts() bool`

HasNoOfCloudAccounts returns a boolean if a field has been set.

### GetNoOfHosts

`func (o *ModelLicense) GetNoOfHosts() int32`

GetNoOfHosts returns the NoOfHosts field if non-nil, zero value otherwise.

### GetNoOfHostsOk

`func (o *ModelLicense) GetNoOfHostsOk() (*int32, bool)`

GetNoOfHostsOk returns a tuple with the NoOfHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfHosts

`func (o *ModelLicense) SetNoOfHosts(v int32)`

SetNoOfHosts sets NoOfHosts field to given value.

### HasNoOfHosts

`func (o *ModelLicense) HasNoOfHosts() bool`

HasNoOfHosts returns a boolean if a field has been set.

### GetNoOfImagesInRegistry

`func (o *ModelLicense) GetNoOfImagesInRegistry() int32`

GetNoOfImagesInRegistry returns the NoOfImagesInRegistry field if non-nil, zero value otherwise.

### GetNoOfImagesInRegistryOk

`func (o *ModelLicense) GetNoOfImagesInRegistryOk() (*int32, bool)`

GetNoOfImagesInRegistryOk returns a tuple with the NoOfImagesInRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfImagesInRegistry

`func (o *ModelLicense) SetNoOfImagesInRegistry(v int32)`

SetNoOfImagesInRegistry sets NoOfImagesInRegistry field to given value.

### HasNoOfImagesInRegistry

`func (o *ModelLicense) HasNoOfImagesInRegistry() bool`

HasNoOfImagesInRegistry returns a boolean if a field has been set.

### GetNoOfRegistries

`func (o *ModelLicense) GetNoOfRegistries() int32`

GetNoOfRegistries returns the NoOfRegistries field if non-nil, zero value otherwise.

### GetNoOfRegistriesOk

`func (o *ModelLicense) GetNoOfRegistriesOk() (*int32, bool)`

GetNoOfRegistriesOk returns a tuple with the NoOfRegistries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfRegistries

`func (o *ModelLicense) SetNoOfRegistries(v int32)`

SetNoOfRegistries sets NoOfRegistries field to given value.

### HasNoOfRegistries

`func (o *ModelLicense) HasNoOfRegistries() bool`

HasNoOfRegistries returns a boolean if a field has been set.

### GetNotificationThresholdPercentage

`func (o *ModelLicense) GetNotificationThresholdPercentage() int32`

GetNotificationThresholdPercentage returns the NotificationThresholdPercentage field if non-nil, zero value otherwise.

### GetNotificationThresholdPercentageOk

`func (o *ModelLicense) GetNotificationThresholdPercentageOk() (*int32, bool)`

GetNotificationThresholdPercentageOk returns a tuple with the NotificationThresholdPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationThresholdPercentage

`func (o *ModelLicense) SetNotificationThresholdPercentage(v int32)`

SetNotificationThresholdPercentage sets NotificationThresholdPercentage field to given value.

### HasNotificationThresholdPercentage

`func (o *ModelLicense) HasNotificationThresholdPercentage() bool`

HasNotificationThresholdPercentage returns a boolean if a field has been set.

### GetNotificationThresholdUpdatedAt

`func (o *ModelLicense) GetNotificationThresholdUpdatedAt() int32`

GetNotificationThresholdUpdatedAt returns the NotificationThresholdUpdatedAt field if non-nil, zero value otherwise.

### GetNotificationThresholdUpdatedAtOk

`func (o *ModelLicense) GetNotificationThresholdUpdatedAtOk() (*int32, bool)`

GetNotificationThresholdUpdatedAtOk returns a tuple with the NotificationThresholdUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationThresholdUpdatedAt

`func (o *ModelLicense) SetNotificationThresholdUpdatedAt(v int32)`

SetNotificationThresholdUpdatedAt sets NotificationThresholdUpdatedAt field to given value.

### HasNotificationThresholdUpdatedAt

`func (o *ModelLicense) HasNotificationThresholdUpdatedAt() bool`

HasNotificationThresholdUpdatedAt returns a boolean if a field has been set.

### GetRegistryCredentials

`func (o *ModelLicense) GetRegistryCredentials() ModelRegistryCredentials`

GetRegistryCredentials returns the RegistryCredentials field if non-nil, zero value otherwise.

### GetRegistryCredentialsOk

`func (o *ModelLicense) GetRegistryCredentialsOk() (*ModelRegistryCredentials, bool)`

GetRegistryCredentialsOk returns a tuple with the RegistryCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryCredentials

`func (o *ModelLicense) SetRegistryCredentials(v ModelRegistryCredentials)`

SetRegistryCredentials sets RegistryCredentials field to given value.

### HasRegistryCredentials

`func (o *ModelLicense) HasRegistryCredentials() bool`

HasRegistryCredentials returns a boolean if a field has been set.

### GetStartDate

`func (o *ModelLicense) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ModelLicense) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ModelLicense) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ModelLicense) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


