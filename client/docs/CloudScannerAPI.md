# \CloudScannerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountResultsCloudComplianceScan**](CloudScannerAPI.md#CountResultsCloudComplianceScan) | **Post** /deepfence/scan/results/count/cloud-compliance | Get Cloud Compliance Scan Results
[**IngestCloudComplianceScanStatus**](CloudScannerAPI.md#IngestCloudComplianceScanStatus) | **Post** /deepfence/ingest/cloud-compliance-status | Ingest Cloud Compliances scan status
[**IngestCloudCompliances**](CloudScannerAPI.md#IngestCloudCompliances) | **Post** /deepfence/ingest/cloud-compliance | Ingest Cloud Compliances
[**ListCloudComplianceScan**](CloudScannerAPI.md#ListCloudComplianceScan) | **Post** /deepfence/scan/list/cloud-compliance | Get Cloud Compliance Scans List
[**ResultsCloudComplianceScan**](CloudScannerAPI.md#ResultsCloudComplianceScan) | **Post** /deepfence/scan/results/cloud-compliance | Get Cloud Compliance Scan Results
[**StatusCloudComplianceScan**](CloudScannerAPI.md#StatusCloudComplianceScan) | **Post** /deepfence/scan/status/cloud-compliance | Get Cloud Compliance Scan Status



## CountResultsCloudComplianceScan

> SearchSearchCountResp CountResultsCloudComplianceScan(ctx).ModelScanResultsReq(modelScanResultsReq).Execute()

Get Cloud Compliance Scan Results



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
	modelScanResultsReq := *openapiclient.NewModelScanResultsReq(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), "ScanId_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ModelScanResultsReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudScannerAPI.CountResultsCloudComplianceScan(context.Background()).ModelScanResultsReq(modelScanResultsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudScannerAPI.CountResultsCloudComplianceScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountResultsCloudComplianceScan`: SearchSearchCountResp
	fmt.Fprintf(os.Stdout, "Response from `CloudScannerAPI.CountResultsCloudComplianceScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountResultsCloudComplianceScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelScanResultsReq** | [**ModelScanResultsReq**](ModelScanResultsReq.md) |  | 

### Return type

[**SearchSearchCountResp**](SearchSearchCountResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IngestCloudComplianceScanStatus

> IngestCloudComplianceScanStatus(ctx).IngestersCloudComplianceScanStatus(ingestersCloudComplianceScanStatus).Execute()

Ingest Cloud Compliances scan status



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
	ingestersCloudComplianceScanStatus := []openapiclient.IngestersCloudComplianceScanStatus{*openapiclient.NewIngestersCloudComplianceScanStatus()} // []IngestersCloudComplianceScanStatus |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudScannerAPI.IngestCloudComplianceScanStatus(context.Background()).IngestersCloudComplianceScanStatus(ingestersCloudComplianceScanStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudScannerAPI.IngestCloudComplianceScanStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestCloudComplianceScanStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestersCloudComplianceScanStatus** | [**[]IngestersCloudComplianceScanStatus**](IngestersCloudComplianceScanStatus.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IngestCloudCompliances

> IngestCloudCompliances(ctx).IngestersCloudCompliance(ingestersCloudCompliance).Execute()

Ingest Cloud Compliances



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
	ingestersCloudCompliance := []openapiclient.IngestersCloudCompliance{*openapiclient.NewIngestersCloudCompliance()} // []IngestersCloudCompliance |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudScannerAPI.IngestCloudCompliances(context.Background()).IngestersCloudCompliance(ingestersCloudCompliance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudScannerAPI.IngestCloudCompliances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestCloudCompliancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestersCloudCompliance** | [**[]IngestersCloudCompliance**](IngestersCloudCompliance.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudComplianceScan

> ModelScanListResp ListCloudComplianceScan(ctx).ModelScanListReq(modelScanListReq).Execute()

Get Cloud Compliance Scans List



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
	modelScanListReq := *openapiclient.NewModelScanListReq(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []openapiclient.ModelNodeIdentifier{*openapiclient.NewModelNodeIdentifier("NodeId_example", "NodeType_example")}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ModelScanListReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudScannerAPI.ListCloudComplianceScan(context.Background()).ModelScanListReq(modelScanListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudScannerAPI.ListCloudComplianceScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCloudComplianceScan`: ModelScanListResp
	fmt.Fprintf(os.Stdout, "Response from `CloudScannerAPI.ListCloudComplianceScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCloudComplianceScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelScanListReq** | [**ModelScanListReq**](ModelScanListReq.md) |  | 

### Return type

[**ModelScanListResp**](ModelScanListResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResultsCloudComplianceScan

> ModelCloudComplianceScanResult ResultsCloudComplianceScan(ctx).ModelScanResultsReq(modelScanResultsReq).Execute()

Get Cloud Compliance Scan Results



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
	modelScanResultsReq := *openapiclient.NewModelScanResultsReq(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), "ScanId_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ModelScanResultsReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudScannerAPI.ResultsCloudComplianceScan(context.Background()).ModelScanResultsReq(modelScanResultsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudScannerAPI.ResultsCloudComplianceScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResultsCloudComplianceScan`: ModelCloudComplianceScanResult
	fmt.Fprintf(os.Stdout, "Response from `CloudScannerAPI.ResultsCloudComplianceScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResultsCloudComplianceScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelScanResultsReq** | [**ModelScanResultsReq**](ModelScanResultsReq.md) |  | 

### Return type

[**ModelCloudComplianceScanResult**](ModelCloudComplianceScanResult.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatusCloudComplianceScan

> ModelComplianceScanStatusResp StatusCloudComplianceScan(ctx).ModelScanStatusReq(modelScanStatusReq).Execute()

Get Cloud Compliance Scan Status



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
	modelScanStatusReq := *openapiclient.NewModelScanStatusReq("BulkScanId_example", []string{"ScanIds_example"}) // ModelScanStatusReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudScannerAPI.StatusCloudComplianceScan(context.Background()).ModelScanStatusReq(modelScanStatusReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudScannerAPI.StatusCloudComplianceScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatusCloudComplianceScan`: ModelComplianceScanStatusResp
	fmt.Fprintf(os.Stdout, "Response from `CloudScannerAPI.StatusCloudComplianceScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStatusCloudComplianceScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelScanStatusReq** | [**ModelScanStatusReq**](ModelScanStatusReq.md) |  | 

### Return type

[**ModelComplianceScanStatusResp**](ModelComplianceScanStatusResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

