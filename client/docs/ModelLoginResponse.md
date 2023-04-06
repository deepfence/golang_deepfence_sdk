# ModelLoginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** |  | 
**OnboardingRequired** | **bool** |  | 
**PasswordInvalidated** | **bool** |  | 
**RefreshToken** | **string** |  | 

## Methods

### NewModelLoginResponse

`func NewModelLoginResponse(accessToken string, onboardingRequired bool, passwordInvalidated bool, refreshToken string, ) *ModelLoginResponse`

NewModelLoginResponse instantiates a new ModelLoginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelLoginResponseWithDefaults

`func NewModelLoginResponseWithDefaults() *ModelLoginResponse`

NewModelLoginResponseWithDefaults instantiates a new ModelLoginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *ModelLoginResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ModelLoginResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ModelLoginResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetOnboardingRequired

`func (o *ModelLoginResponse) GetOnboardingRequired() bool`

GetOnboardingRequired returns the OnboardingRequired field if non-nil, zero value otherwise.

### GetOnboardingRequiredOk

`func (o *ModelLoginResponse) GetOnboardingRequiredOk() (*bool, bool)`

GetOnboardingRequiredOk returns a tuple with the OnboardingRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingRequired

`func (o *ModelLoginResponse) SetOnboardingRequired(v bool)`

SetOnboardingRequired sets OnboardingRequired field to given value.


### GetPasswordInvalidated

`func (o *ModelLoginResponse) GetPasswordInvalidated() bool`

GetPasswordInvalidated returns the PasswordInvalidated field if non-nil, zero value otherwise.

### GetPasswordInvalidatedOk

`func (o *ModelLoginResponse) GetPasswordInvalidatedOk() (*bool, bool)`

GetPasswordInvalidatedOk returns a tuple with the PasswordInvalidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordInvalidated

`func (o *ModelLoginResponse) SetPasswordInvalidated(v bool)`

SetPasswordInvalidated sets PasswordInvalidated field to given value.


### GetRefreshToken

`func (o *ModelLoginResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *ModelLoginResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *ModelLoginResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


