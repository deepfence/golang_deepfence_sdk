# ModelGenerateLicenseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GenerateLicenseLink** | Pointer to **string** |  | [optional] 
**Message** | **string** |  | 
**Success** | **bool** |  | 

## Methods

### NewModelGenerateLicenseResponse

`func NewModelGenerateLicenseResponse(message string, success bool, ) *ModelGenerateLicenseResponse`

NewModelGenerateLicenseResponse instantiates a new ModelGenerateLicenseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelGenerateLicenseResponseWithDefaults

`func NewModelGenerateLicenseResponseWithDefaults() *ModelGenerateLicenseResponse`

NewModelGenerateLicenseResponseWithDefaults instantiates a new ModelGenerateLicenseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenerateLicenseLink

`func (o *ModelGenerateLicenseResponse) GetGenerateLicenseLink() string`

GetGenerateLicenseLink returns the GenerateLicenseLink field if non-nil, zero value otherwise.

### GetGenerateLicenseLinkOk

`func (o *ModelGenerateLicenseResponse) GetGenerateLicenseLinkOk() (*string, bool)`

GetGenerateLicenseLinkOk returns a tuple with the GenerateLicenseLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateLicenseLink

`func (o *ModelGenerateLicenseResponse) SetGenerateLicenseLink(v string)`

SetGenerateLicenseLink sets GenerateLicenseLink field to given value.

### HasGenerateLicenseLink

`func (o *ModelGenerateLicenseResponse) HasGenerateLicenseLink() bool`

HasGenerateLicenseLink returns a boolean if a field has been set.

### GetMessage

`func (o *ModelGenerateLicenseResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ModelGenerateLicenseResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ModelGenerateLicenseResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSuccess

`func (o *ModelGenerateLicenseResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ModelGenerateLicenseResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ModelGenerateLicenseResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


