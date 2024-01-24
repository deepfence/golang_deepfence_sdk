# \DiffAddAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiffAddCloudCompliance**](DiffAddAPI.md#DiffAddCloudCompliance) | **Post** /deepfence/diff-add/cloud-compliance | Get Cloud Compliance Diff
[**DiffAddCompliance**](DiffAddAPI.md#DiffAddCompliance) | **Post** /deepfence/diff-add/compliance | Get Compliance Diff
[**DiffAddMalware**](DiffAddAPI.md#DiffAddMalware) | **Post** /deepfence/diff-add/malware | Get Malware Diff
[**DiffAddSecret**](DiffAddAPI.md#DiffAddSecret) | **Post** /deepfence/diff-add/secret | Get Secret Diff
[**DiffAddVulnerability**](DiffAddAPI.md#DiffAddVulnerability) | **Post** /deepfence/diff-add/vulnerability | Get Vulnerability Diff



## DiffAddCloudCompliance

> ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCloudCompliance DiffAddCloudCompliance(ctx).ModelScanCompareReq(modelScanCompareReq).Execute()

Get Cloud Compliance Diff



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/deepfence/golang_deepfence_sdk/client"
)

func main() {
	modelScanCompareReq := *openapiclient.NewModelScanCompareReq("BaseScanId_example", *openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), "ToScanId_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ModelScanCompareReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiffAddAPI.DiffAddCloudCompliance(context.Background()).ModelScanCompareReq(modelScanCompareReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiffAddAPI.DiffAddCloudCompliance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiffAddCloudCompliance`: ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCloudCompliance
	fmt.Fprintf(os.Stdout, "Response from `DiffAddAPI.DiffAddCloudCompliance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiffAddCloudComplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelScanCompareReq** | [**ModelScanCompareReq**](ModelScanCompareReq.md) |  | 

### Return type

[**ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCloudCompliance**](ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCloudCompliance.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiffAddCompliance

> ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance DiffAddCompliance(ctx).ModelScanCompareReq(modelScanCompareReq).Execute()

Get Compliance Diff



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/deepfence/golang_deepfence_sdk/client"
)

func main() {
	modelScanCompareReq := *openapiclient.NewModelScanCompareReq("BaseScanId_example", *openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), "ToScanId_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ModelScanCompareReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiffAddAPI.DiffAddCompliance(context.Background()).ModelScanCompareReq(modelScanCompareReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiffAddAPI.DiffAddCompliance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiffAddCompliance`: ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance
	fmt.Fprintf(os.Stdout, "Response from `DiffAddAPI.DiffAddCompliance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiffAddComplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelScanCompareReq** | [**ModelScanCompareReq**](ModelScanCompareReq.md) |  | 

### Return type

[**ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance**](ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelCompliance.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiffAddMalware

> ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelMalware DiffAddMalware(ctx).ModelScanCompareReq(modelScanCompareReq).Execute()

Get Malware Diff



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/deepfence/golang_deepfence_sdk/client"
)

func main() {
	modelScanCompareReq := *openapiclient.NewModelScanCompareReq("BaseScanId_example", *openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), "ToScanId_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ModelScanCompareReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiffAddAPI.DiffAddMalware(context.Background()).ModelScanCompareReq(modelScanCompareReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiffAddAPI.DiffAddMalware``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiffAddMalware`: ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelMalware
	fmt.Fprintf(os.Stdout, "Response from `DiffAddAPI.DiffAddMalware`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiffAddMalwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelScanCompareReq** | [**ModelScanCompareReq**](ModelScanCompareReq.md) |  | 

### Return type

[**ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelMalware**](ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelMalware.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiffAddSecret

> ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret DiffAddSecret(ctx).ModelScanCompareReq(modelScanCompareReq).Execute()

Get Secret Diff



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/deepfence/golang_deepfence_sdk/client"
)

func main() {
	modelScanCompareReq := *openapiclient.NewModelScanCompareReq("BaseScanId_example", *openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), "ToScanId_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ModelScanCompareReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiffAddAPI.DiffAddSecret(context.Background()).ModelScanCompareReq(modelScanCompareReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiffAddAPI.DiffAddSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiffAddSecret`: ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret
	fmt.Fprintf(os.Stdout, "Response from `DiffAddAPI.DiffAddSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiffAddSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelScanCompareReq** | [**ModelScanCompareReq**](ModelScanCompareReq.md) |  | 

### Return type

[**ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret**](ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiffAddVulnerability

> ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelVulnerability DiffAddVulnerability(ctx).ModelScanCompareReq(modelScanCompareReq).Execute()

Get Vulnerability Diff



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/deepfence/golang_deepfence_sdk/client"
)

func main() {
	modelScanCompareReq := *openapiclient.NewModelScanCompareReq("BaseScanId_example", *openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), "ToScanId_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ModelScanCompareReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiffAddAPI.DiffAddVulnerability(context.Background()).ModelScanCompareReq(modelScanCompareReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiffAddAPI.DiffAddVulnerability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiffAddVulnerability`: ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelVulnerability
	fmt.Fprintf(os.Stdout, "Response from `DiffAddAPI.DiffAddVulnerability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiffAddVulnerabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelScanCompareReq** | [**ModelScanCompareReq**](ModelScanCompareReq.md) |  | 

### Return type

[**ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelVulnerability**](ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelVulnerability.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

