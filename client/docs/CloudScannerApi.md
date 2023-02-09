# \CloudScannerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IngestCloudComplianceScanStatus**](CloudScannerApi.md#IngestCloudComplianceScanStatus) | **Post** /deepfence/ingest/cloud-compliance-scan-status | Ingest Cloud Compliances
[**IngestCloudCompliances**](CloudScannerApi.md#IngestCloudCompliances) | **Post** /deepfence/ingest/cloud-compliance | Ingest Cloud Compliances
[**StatusCloudComplianceScan**](CloudScannerApi.md#StatusCloudComplianceScan) | **Get** /deepfence/scan/status/cloud-compliance | Get Cloud Compliance Scan Status



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
    openapiclient "./openapi"
)

func main() {
    ingestersCloudCompliance := []openapiclient.IngestersCloudCompliance{*openapiclient.NewIngestersCloudCompliance()} // []IngestersCloudCompliance |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudScannerApi.IngestCloudComplianceScanStatus(context.Background()).IngestersCloudCompliance(ingestersCloudCompliance).Execute()
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
    openapiclient "./openapi"
)

func main() {
    ingestersCloudCompliance := []openapiclient.IngestersCloudCompliance{*openapiclient.NewIngestersCloudCompliance()} // []IngestersCloudCompliance |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudScannerApi.IngestCloudCompliances(context.Background()).IngestersCloudCompliance(ingestersCloudCompliance).Execute()
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


## StatusCloudComplianceScan

> ModelComplianceScanStatusResp StatusCloudComplianceScan(ctx).ScanIds(scanIds).BulkScanId(bulkScanId).Execute()

Get Cloud Compliance Scan Status



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    scanIds := []string{"Inner_example"} // []string | 
    bulkScanId := "bulkScanId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudScannerApi.StatusCloudComplianceScan(context.Background()).ScanIds(scanIds).BulkScanId(bulkScanId).Execute()
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

