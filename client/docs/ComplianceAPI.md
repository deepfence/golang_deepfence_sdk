# \ComplianceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountResultsComplianceScan**](ComplianceAPI.md#CountResultsComplianceScan) | **Post** /deepfence/scan/results/count/compliance | Get Compliance Scans Results
[**IngestComplianceScanStatus**](ComplianceAPI.md#IngestComplianceScanStatus) | **Post** /deepfence/ingest/compliance-scan-logs | Ingest Compliance Scan Status
[**IngestCompliances**](ComplianceAPI.md#IngestCompliances) | **Post** /deepfence/ingest/compliance | Ingest Compliances
[**ListComplianceScan**](ComplianceAPI.md#ListComplianceScan) | **Post** /deepfence/scan/list/compliance | Get Compliance Scans List
[**ResultsComplianceScan**](ComplianceAPI.md#ResultsComplianceScan) | **Post** /deepfence/scan/results/compliance | Get Compliance Scans Results
[**StartComplianceScan**](ComplianceAPI.md#StartComplianceScan) | **Post** /deepfence/scan/start/compliance | Start Compliance Scan
[**StatusComplianceScan**](ComplianceAPI.md#StatusComplianceScan) | **Post** /deepfence/scan/status/compliance | Get Compliance Scan Status
[**StopComplianceScan**](ComplianceAPI.md#StopComplianceScan) | **Post** /deepfence/scan/stop/compliance | Stop Compliance Scan



## CountResultsComplianceScan

> SearchSearchCountResp CountResultsComplianceScan(ctx).ModelScanResultsReq(modelScanResultsReq).Execute()

Get Compliance Scans Results



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
	resp, r, err := apiClient.ComplianceAPI.CountResultsComplianceScan(context.Background()).ModelScanResultsReq(modelScanResultsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.CountResultsComplianceScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountResultsComplianceScan`: SearchSearchCountResp
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.CountResultsComplianceScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountResultsComplianceScanRequest struct via the builder pattern


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


## IngestComplianceScanStatus

> IngestComplianceScanStatus(ctx).IngestersComplianceScanStatus(ingestersComplianceScanStatus).Execute()

Ingest Compliance Scan Status



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
	ingestersComplianceScanStatus := []openapiclient.IngestersComplianceScanStatus{*openapiclient.NewIngestersComplianceScanStatus()} // []IngestersComplianceScanStatus |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComplianceAPI.IngestComplianceScanStatus(context.Background()).IngestersComplianceScanStatus(ingestersComplianceScanStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.IngestComplianceScanStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestComplianceScanStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestersComplianceScanStatus** | [**[]IngestersComplianceScanStatus**](IngestersComplianceScanStatus.md) |  | 

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


## IngestCompliances

> IngestCompliances(ctx).IngestersCompliance(ingestersCompliance).Execute()

Ingest Compliances



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
	ingestersCompliance := []openapiclient.IngestersCompliance{*openapiclient.NewIngestersCompliance()} // []IngestersCompliance |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComplianceAPI.IngestCompliances(context.Background()).IngestersCompliance(ingestersCompliance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.IngestCompliances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestCompliancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestersCompliance** | [**[]IngestersCompliance**](IngestersCompliance.md) |  | 

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


## ListComplianceScan

> ModelScanListResp ListComplianceScan(ctx).ModelScanListReq(modelScanListReq).Execute()

Get Compliance Scans List



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
	resp, r, err := apiClient.ComplianceAPI.ListComplianceScan(context.Background()).ModelScanListReq(modelScanListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.ListComplianceScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListComplianceScan`: ModelScanListResp
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.ListComplianceScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListComplianceScanRequest struct via the builder pattern


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


## ResultsComplianceScan

> ModelComplianceScanResult ResultsComplianceScan(ctx).ModelScanResultsReq(modelScanResultsReq).Execute()

Get Compliance Scans Results



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
	resp, r, err := apiClient.ComplianceAPI.ResultsComplianceScan(context.Background()).ModelScanResultsReq(modelScanResultsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.ResultsComplianceScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResultsComplianceScan`: ModelComplianceScanResult
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.ResultsComplianceScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResultsComplianceScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelScanResultsReq** | [**ModelScanResultsReq**](ModelScanResultsReq.md) |  | 

### Return type

[**ModelComplianceScanResult**](ModelComplianceScanResult.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartComplianceScan

> ModelScanTriggerResp StartComplianceScan(ctx).ModelComplianceScanTriggerReq(modelComplianceScanTriggerReq).Execute()

Start Compliance Scan



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
	modelComplianceScanTriggerReq := *openapiclient.NewModelComplianceScanTriggerReq([]string{"BenchmarkTypes_example"}, *openapiclient.NewModelScanFilter(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}})), []openapiclient.ModelNodeIdentifier{*openapiclient.NewModelNodeIdentifier("NodeId_example", "NodeType_example")}) // ModelComplianceScanTriggerReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.StartComplianceScan(context.Background()).ModelComplianceScanTriggerReq(modelComplianceScanTriggerReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.StartComplianceScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartComplianceScan`: ModelScanTriggerResp
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.StartComplianceScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartComplianceScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelComplianceScanTriggerReq** | [**ModelComplianceScanTriggerReq**](ModelComplianceScanTriggerReq.md) |  | 

### Return type

[**ModelScanTriggerResp**](ModelScanTriggerResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatusComplianceScan

> ModelScanStatusResp StatusComplianceScan(ctx).ModelScanStatusReq(modelScanStatusReq).Execute()

Get Compliance Scan Status



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
	resp, r, err := apiClient.ComplianceAPI.StatusComplianceScan(context.Background()).ModelScanStatusReq(modelScanStatusReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.StatusComplianceScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatusComplianceScan`: ModelScanStatusResp
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.StatusComplianceScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStatusComplianceScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelScanStatusReq** | [**ModelScanStatusReq**](ModelScanStatusReq.md) |  | 

### Return type

[**ModelScanStatusResp**](ModelScanStatusResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopComplianceScan

> StopComplianceScan(ctx).ModelStopScanRequest(modelStopScanRequest).Execute()

Stop Compliance Scan



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
	modelStopScanRequest := *openapiclient.NewModelStopScanRequest([]string{"ScanIds_example"}, "ScanType_example") // ModelStopScanRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComplianceAPI.StopComplianceScan(context.Background()).ModelStopScanRequest(modelStopScanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.StopComplianceScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStopComplianceScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelStopScanRequest** | [**ModelStopScanRequest**](ModelStopScanRequest.md) |  | 

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

