# ModelComplianceScanTriggerReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GenerateBulkScanId** | **bool** |  | 
**ScanTriggers** | [**[]ModelScanTrigger**](ModelScanTrigger.md) |  | 

## Methods

### NewModelComplianceScanTriggerReq

`func NewModelComplianceScanTriggerReq(generateBulkScanId bool, scanTriggers []ModelScanTrigger, ) *ModelComplianceScanTriggerReq`

NewModelComplianceScanTriggerReq instantiates a new ModelComplianceScanTriggerReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelComplianceScanTriggerReqWithDefaults

`func NewModelComplianceScanTriggerReqWithDefaults() *ModelComplianceScanTriggerReq`

NewModelComplianceScanTriggerReqWithDefaults instantiates a new ModelComplianceScanTriggerReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenerateBulkScanId

`func (o *ModelComplianceScanTriggerReq) GetGenerateBulkScanId() bool`

GetGenerateBulkScanId returns the GenerateBulkScanId field if non-nil, zero value otherwise.

### GetGenerateBulkScanIdOk

`func (o *ModelComplianceScanTriggerReq) GetGenerateBulkScanIdOk() (*bool, bool)`

GetGenerateBulkScanIdOk returns a tuple with the GenerateBulkScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateBulkScanId

`func (o *ModelComplianceScanTriggerReq) SetGenerateBulkScanId(v bool)`

SetGenerateBulkScanId sets GenerateBulkScanId field to given value.


### GetScanTriggers

`func (o *ModelComplianceScanTriggerReq) GetScanTriggers() []ModelScanTrigger`

GetScanTriggers returns the ScanTriggers field if non-nil, zero value otherwise.

### GetScanTriggersOk

`func (o *ModelComplianceScanTriggerReq) GetScanTriggersOk() (*[]ModelScanTrigger, bool)`

GetScanTriggersOk returns a tuple with the ScanTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanTriggers

`func (o *ModelComplianceScanTriggerReq) SetScanTriggers(v []ModelScanTrigger)`

SetScanTriggers sets ScanTriggers field to given value.


### SetScanTriggersNil

`func (o *ModelComplianceScanTriggerReq) SetScanTriggersNil(b bool)`

 SetScanTriggersNil sets the value for ScanTriggers to be an explicit nil

### UnsetScanTriggers
`func (o *ModelComplianceScanTriggerReq) UnsetScanTriggers()`

UnsetScanTriggers ensures that no value is present for ScanTriggers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


