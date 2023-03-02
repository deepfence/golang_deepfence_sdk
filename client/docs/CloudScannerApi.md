# \CloudScannerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountResultsCloudComplianceScan**](CloudScannerApi.md#CountResultsCloudComplianceScan) | **Post** /deepfence/scan/results/count/cloud-compliance | Get Cloud Compliance Scan Results
[**IngestCloudComplianceScanStatus**](CloudScannerApi.md#IngestCloudComplianceScanStatus) | **Post** /deepfence/ingest/cloud-compliance-scan-status | Ingest Cloud Compliances
[**IngestCloudCompliances**](CloudScannerApi.md#IngestCloudCompliances) | **Post** /deepfence/ingest/cloud-compliance | Ingest Cloud Compliances
[**ResultsCloudComplianceScan**](CloudScannerApi.md#ResultsCloudComplianceScan) | **Post** /deepfence/scan/results/cloud-compliance | Get Cloud Compliance Scan Results
[**StatusCloudComplianceScan**](CloudScannerApi.md#StatusCloudComplianceScan) | **Post** /deepfence/scan/status/cloud-compliance | Get Cloud Compliance Scan Status



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
    modelScanResultsReq := *openapiclient.NewModelScanResultsReq(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), "ScanId_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ModelScanResultsReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudScannerApi.CountResultsCloudComplianceScan(context.Background()).ModelScanResultsReq(modelScanResultsReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudScannerApi.CountResultsCloudComplianceScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountResultsCloudComplianceScan`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `CloudScannerApi.CountResultsCloudComplianceScan`: %v\n", resp)
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

> IngestCloudComplianceScanStatus(ctx).IngestersCloudCompliance(ingestersCloudCompliance).Execute()

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
    r, err := apiClient.CloudScannerApi.IngestCloudComplianceScanStatus(context.Background()).IngestersCloudCompliance(ingestersCloudCompliance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudScannerApi.IngestCloudComplianceScanStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestCloudComplianceScanStatusRequest struct via the builder pattern


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
    r, err := apiClient.CloudScannerApi.IngestCloudCompliances(context.Background()).IngestersCloudCompliance(ingestersCloudCompliance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudScannerApi.IngestCloudCompliances``: %v\n", err)
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
    modelScanResultsReq := *openapiclient.NewModelScanResultsReq(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), "ScanId_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ModelScanResultsReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudScannerApi.ResultsCloudComplianceScan(context.Background()).ModelScanResultsReq(modelScanResultsReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudScannerApi.ResultsCloudComplianceScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResultsCloudComplianceScan`: ModelCloudComplianceScanResult
    fmt.Fprintf(os.Stdout, "Response from `CloudScannerApi.ResultsCloudComplianceScan`: %v\n", resp)
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
    resp, r, err := apiClient.CloudScannerApi.StatusCloudComplianceScan(context.Background()).ModelScanStatusReq(modelScanStatusReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudScannerApi.StatusCloudComplianceScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StatusCloudComplianceScan`: ModelComplianceScanStatusResp
    fmt.Fprintf(os.Stdout, "Response from `CloudScannerApi.StatusCloudComplianceScan`: %v\n", resp)
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

