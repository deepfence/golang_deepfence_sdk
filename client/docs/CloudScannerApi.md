# \CloudScannerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IngestCloudCompliances**](CloudScannerApi.md#IngestCloudCompliances) | **Post** /deepfence/ingest/cloud-compliance | Ingest Cloud Compliances
[**StartCloudComplianceScans**](CloudScannerApi.md#StartCloudComplianceScans) | **Post** /deepfence/scan/start/cloud-compliance | Start Cloud Compliance Scans



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


## StartCloudComplianceScans

> ModelScanTriggerResp StartCloudComplianceScans(ctx).ModelCloudComplianceScanTriggerReq(modelCloudComplianceScanTriggerReq).Execute()

Start Cloud Compliance Scans



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
    modelCloudComplianceScanTriggerReq := *openapiclient.NewModelCloudComplianceScanTriggerReq([]openapiclient.ModelCloudComplianceScanTrigger{*openapiclient.NewModelCloudComplianceScanTrigger([]string{"BenchmarkTypes_example"}, "NodeId_example")}) // ModelCloudComplianceScanTriggerReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudScannerApi.StartCloudComplianceScans(context.Background()).ModelCloudComplianceScanTriggerReq(modelCloudComplianceScanTriggerReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudScannerApi.StartCloudComplianceScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartCloudComplianceScans`: ModelScanTriggerResp
    fmt.Fprintf(os.Stdout, "Response from `CloudScannerApi.StartCloudComplianceScans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartCloudComplianceScansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelCloudComplianceScanTriggerReq** | [**ModelCloudComplianceScanTriggerReq**](ModelCloudComplianceScanTriggerReq.md) |  | 

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

