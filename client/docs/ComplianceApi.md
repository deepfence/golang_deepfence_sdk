# \ComplianceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IngestCompliances**](ComplianceApi.md#IngestCompliances) | **Post** /deepfence/ingest/compliance | Ingest Compliances
[**ListComplianceScan**](ComplianceApi.md#ListComplianceScan) | **Post** /deepfence/scan/list/compliance | Get Compliance Scans List
[**ResultsComplianceScan**](ComplianceApi.md#ResultsComplianceScan) | **Post** /deepfence/scan/results/compliance | Get Compliance Scans Results
[**StartComplianceScan**](ComplianceApi.md#StartComplianceScan) | **Post** /deepfence/scan/start/compliance | Start Compliance Scan
[**StatusComplianceScan**](ComplianceApi.md#StatusComplianceScan) | **Get** /deepfence/scan/status/compliance | Get Compliance Scan Status
[**StopComplianceScan**](ComplianceApi.md#StopComplianceScan) | **Post** /deepfence/scan/stop/compliance | Stop Compliance Scan



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
    r, err := apiClient.ComplianceApi.IngestCompliances(context.Background()).IngestersCompliance(ingestersCompliance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComplianceApi.IngestCompliances``: %v\n", err)
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
    modelScanListReq := *openapiclient.NewModelScanListReq([]openapiclient.ModelNodeIdentifier{*openapiclient.NewModelNodeIdentifier("NodeId_example", "NodeType_example")}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ModelScanListReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComplianceApi.ListComplianceScan(context.Background()).ModelScanListReq(modelScanListReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComplianceApi.ListComplianceScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListComplianceScan`: ModelScanListResp
    fmt.Fprintf(os.Stdout, "Response from `ComplianceApi.ListComplianceScan`: %v\n", resp)
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
    modelScanResultsReq := *openapiclient.NewModelScanResultsReq(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderField_example")), "ScanId_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ModelScanResultsReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComplianceApi.ResultsComplianceScan(context.Background()).ModelScanResultsReq(modelScanResultsReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComplianceApi.ResultsComplianceScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResultsComplianceScan`: ModelComplianceScanResult
    fmt.Fprintf(os.Stdout, "Response from `ComplianceApi.ResultsComplianceScan`: %v\n", resp)
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
    resp, r, err := apiClient.ComplianceApi.StartComplianceScan(context.Background()).ModelComplianceScanTriggerReq(modelComplianceScanTriggerReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComplianceApi.StartComplianceScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartComplianceScan`: ModelScanTriggerResp
    fmt.Fprintf(os.Stdout, "Response from `ComplianceApi.StartComplianceScan`: %v\n", resp)
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

> ModelComplianceScanStatusResp StatusComplianceScan(ctx).ScanIds(scanIds).BulkScanId(bulkScanId).Execute()

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
    scanIds := []string{"Inner_example"} // []string | 
    bulkScanId := "bulkScanId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComplianceApi.StatusComplianceScan(context.Background()).ScanIds(scanIds).BulkScanId(bulkScanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComplianceApi.StatusComplianceScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StatusComplianceScan`: ModelComplianceScanStatusResp
    fmt.Fprintf(os.Stdout, "Response from `ComplianceApi.StatusComplianceScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStatusComplianceScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scanIds** | **[]string** |  | 
 **bulkScanId** | **string** |  | 

### Return type

[**ModelComplianceScanStatusResp**](ModelComplianceScanStatusResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopComplianceScan

> StopComplianceScan(ctx).ModelComplianceScanTriggerReq(modelComplianceScanTriggerReq).Execute()

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
    modelComplianceScanTriggerReq := *openapiclient.NewModelComplianceScanTriggerReq([]string{"BenchmarkTypes_example"}, *openapiclient.NewModelScanFilter(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}})), []openapiclient.ModelNodeIdentifier{*openapiclient.NewModelNodeIdentifier("NodeId_example", "NodeType_example")}) // ModelComplianceScanTriggerReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ComplianceApi.StopComplianceScan(context.Background()).ModelComplianceScanTriggerReq(modelComplianceScanTriggerReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComplianceApi.StopComplianceScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStopComplianceScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelComplianceScanTriggerReq** | [**ModelComplianceScanTriggerReq**](ModelComplianceScanTriggerReq.md) |  | 

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

