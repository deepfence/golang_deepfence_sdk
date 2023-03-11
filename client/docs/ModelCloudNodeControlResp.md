# ModelCloudNodeControlResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Controls** | Pointer to [**[]ModelCloudNodeComplianceControl**](ModelCloudNodeComplianceControl.md) |  | [optional] 

## Methods

### NewModelCloudNodeControlResp

`func NewModelCloudNodeControlResp() *ModelCloudNodeControlResp`

NewModelCloudNodeControlResp instantiates a new ModelCloudNodeControlResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCloudNodeControlRespWithDefaults

`func NewModelCloudNodeControlRespWithDefaults() *ModelCloudNodeControlResp`

NewModelCloudNodeControlRespWithDefaults instantiates a new ModelCloudNodeControlResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControls

`func (o *ModelCloudNodeControlResp) GetControls() []ModelCloudNodeComplianceControl`

GetControls returns the Controls field if non-nil, zero value otherwise.

### GetControlsOk

`func (o *ModelCloudNodeControlResp) GetControlsOk() (*[]ModelCloudNodeComplianceControl, bool)`

GetControlsOk returns a tuple with the Controls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControls

`func (o *ModelCloudNodeControlResp) SetControls(v []ModelCloudNodeComplianceControl)`

SetControls sets Controls field to given value.

### HasControls

`func (o *ModelCloudNodeControlResp) HasControls() bool`

HasControls returns a boolean if a field has been set.

### SetControlsNil

`func (o *ModelCloudNodeControlResp) SetControlsNil(b bool)`

 SetControlsNil sets the value for Controls to be an explicit nil

### UnsetControls
`func (o *ModelCloudNodeControlResp) UnsetControls()`

UnsetControls ensures that no value is present for Controls, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


