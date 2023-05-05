# \ScanResultsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkDeleteScans**](ScanResultsApi.md#BulkDeleteScans) | **Post** /deepfence/scans/bulk/delete | Bulk Delete Scans
[**DeleteScanResult**](ScanResultsApi.md#DeleteScanResult) | **Patch** /deepfence/scan/results/action/delete | Delete selected scan results
[**DeleteScanResultsForScanID**](ScanResultsApi.md#DeleteScanResultsForScanID) | **Delete** /deepfence/scan/{scan_type}/{scan_id} | Delete all scan results for a scan id
[**DownloadScanResults**](ScanResultsApi.md#DownloadScanResults) | **Get** /deepfence/scan/{scan_type}/{scan_id}/download | Download Scans Results
[**GetAllNodesInScanResults**](ScanResultsApi.md#GetAllNodesInScanResults) | **Post** /deepfence/scan/nodes-in-result | Get all nodes in given scan result ids
[**MaskScanResult**](ScanResultsApi.md#MaskScanResult) | **Post** /deepfence/scan/results/action/mask | Mask Scans Results
[**NotifyScanResult**](ScanResultsApi.md#NotifyScanResult) | **Post** /deepfence/scan/results/action/notify | Notify Scans Results
[**UnmaskScanResult**](ScanResultsApi.md#UnmaskScanResult) | **Post** /deepfence/scan/results/action/unmask | Unmask Scans Results



## BulkDeleteScans

> BulkDeleteScans(ctx).ModelBulkDeleteScansRequest(modelBulkDeleteScansRequest).Execute()

Bulk Delete Scans



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
    modelBulkDeleteScansRequest := *openapiclient.NewModelBulkDeleteScansRequest(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), "ScanType_example") // ModelBulkDeleteScansRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ScanResultsApi.BulkDeleteScans(context.Background()).ModelBulkDeleteScansRequest(modelBulkDeleteScansRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScanResultsApi.BulkDeleteScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkDeleteScansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelBulkDeleteScansRequest** | [**ModelBulkDeleteScansRequest**](ModelBulkDeleteScansRequest.md) |  | 

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


## DeleteScanResult

> DeleteScanResult(ctx).ModelScanResultsActionRequest(modelScanResultsActionRequest).Execute()

Delete selected scan results



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
    modelScanResultsActionRequest := *openapiclient.NewModelScanResultsActionRequest([]string{"ResultIds_example"}, "ScanId_example", "ScanType_example") // ModelScanResultsActionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ScanResultsApi.DeleteScanResult(context.Background()).ModelScanResultsActionRequest(modelScanResultsActionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScanResultsApi.DeleteScanResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScanResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelScanResultsActionRequest** | [**ModelScanResultsActionRequest**](ModelScanResultsActionRequest.md) |  | 

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


## DeleteScanResultsForScanID

> DeleteScanResultsForScanID(ctx, scanId, scanType).Execute()

Delete all scan results for a scan id



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
    scanId := "scanId_example" // string | 
    scanType := "scanType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ScanResultsApi.DeleteScanResultsForScanID(context.Background(), scanId, scanType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScanResultsApi.DeleteScanResultsForScanID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scanId** | **string** |  | 
**scanType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScanResultsForScanIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadScanResults

> ModelDownloadScanResultsResponse DownloadScanResults(ctx, scanId, scanType).Execute()

Download Scans Results



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
    scanId := "scanId_example" // string | 
    scanType := "scanType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScanResultsApi.DownloadScanResults(context.Background(), scanId, scanType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScanResultsApi.DownloadScanResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadScanResults`: ModelDownloadScanResultsResponse
    fmt.Fprintf(os.Stdout, "Response from `ScanResultsApi.DownloadScanResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scanId** | **string** |  | 
**scanType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadScanResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelDownloadScanResultsResponse**](ModelDownloadScanResultsResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllNodesInScanResults

> []ModelScanResultBasicNode GetAllNodesInScanResults(ctx).ModelNodesInScanResultRequest(modelNodesInScanResultRequest).Execute()

Get all nodes in given scan result ids



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
    modelNodesInScanResultRequest := *openapiclient.NewModelNodesInScanResultRequest([]string{"ResultIds_example"}, "ScanType_example") // ModelNodesInScanResultRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScanResultsApi.GetAllNodesInScanResults(context.Background()).ModelNodesInScanResultRequest(modelNodesInScanResultRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScanResultsApi.GetAllNodesInScanResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllNodesInScanResults`: []ModelScanResultBasicNode
    fmt.Fprintf(os.Stdout, "Response from `ScanResultsApi.GetAllNodesInScanResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllNodesInScanResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelNodesInScanResultRequest** | [**ModelNodesInScanResultRequest**](ModelNodesInScanResultRequest.md) |  | 

### Return type

[**[]ModelScanResultBasicNode**](ModelScanResultBasicNode.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaskScanResult

> MaskScanResult(ctx).ModelScanResultsMaskRequest(modelScanResultsMaskRequest).Execute()

Mask Scans Results



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
    modelScanResultsMaskRequest := *openapiclient.NewModelScanResultsMaskRequest([]string{"ResultIds_example"}, "ScanId_example", "ScanType_example") // ModelScanResultsMaskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ScanResultsApi.MaskScanResult(context.Background()).ModelScanResultsMaskRequest(modelScanResultsMaskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScanResultsApi.MaskScanResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMaskScanResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelScanResultsMaskRequest** | [**ModelScanResultsMaskRequest**](ModelScanResultsMaskRequest.md) |  | 

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


## NotifyScanResult

> NotifyScanResult(ctx).ModelScanResultsActionRequest(modelScanResultsActionRequest).Execute()

Notify Scans Results



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
    modelScanResultsActionRequest := *openapiclient.NewModelScanResultsActionRequest([]string{"ResultIds_example"}, "ScanId_example", "ScanType_example") // ModelScanResultsActionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ScanResultsApi.NotifyScanResult(context.Background()).ModelScanResultsActionRequest(modelScanResultsActionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScanResultsApi.NotifyScanResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotifyScanResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelScanResultsActionRequest** | [**ModelScanResultsActionRequest**](ModelScanResultsActionRequest.md) |  | 

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


## UnmaskScanResult

> UnmaskScanResult(ctx).ModelScanResultsMaskRequest(modelScanResultsMaskRequest).Execute()

Unmask Scans Results



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
    modelScanResultsMaskRequest := *openapiclient.NewModelScanResultsMaskRequest([]string{"ResultIds_example"}, "ScanId_example", "ScanType_example") // ModelScanResultsMaskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ScanResultsApi.UnmaskScanResult(context.Background()).ModelScanResultsMaskRequest(modelScanResultsMaskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScanResultsApi.UnmaskScanResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnmaskScanResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelScanResultsMaskRequest** | [**ModelScanResultsMaskRequest**](ModelScanResultsMaskRequest.md) |  | 

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

