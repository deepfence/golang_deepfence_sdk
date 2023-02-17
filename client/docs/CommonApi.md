# \CommonApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteScanResult**](CommonApi.md#DeleteScanResult) | **Post** /deepfence/scan/results/action/delete | Delete Scans Results
[**MaskScanResult**](CommonApi.md#MaskScanResult) | **Post** /deepfence/scan/results/action/mask | Mask Scans Results
[**NotifyScanResult**](CommonApi.md#NotifyScanResult) | **Post** /deepfence/scan/results/action/notify | Notify Scans Results
[**UnmaskScanResult**](CommonApi.md#UnmaskScanResult) | **Post** /deepfence/scan/results/action/unmask | Unmask Scans Results



## DeleteScanResult

> DeleteScanResult(ctx).ModelScanResultsActionRequest(modelScanResultsActionRequest).Execute()

Delete Scans Results



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
    modelScanResultsActionRequest := *openapiclient.NewModelScanResultsActionRequest([]string{"NodeIds_example"}, "ScanType_example") // ModelScanResultsActionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CommonApi.DeleteScanResult(context.Background()).ModelScanResultsActionRequest(modelScanResultsActionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommonApi.DeleteScanResult``: %v\n", err)
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


## MaskScanResult

> MaskScanResult(ctx).ModelScanResultsActionRequest(modelScanResultsActionRequest).Execute()

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
    modelScanResultsActionRequest := *openapiclient.NewModelScanResultsActionRequest([]string{"NodeIds_example"}, "ScanType_example") // ModelScanResultsActionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CommonApi.MaskScanResult(context.Background()).ModelScanResultsActionRequest(modelScanResultsActionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommonApi.MaskScanResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMaskScanResultRequest struct via the builder pattern


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
    modelScanResultsActionRequest := *openapiclient.NewModelScanResultsActionRequest([]string{"NodeIds_example"}, "ScanType_example") // ModelScanResultsActionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CommonApi.NotifyScanResult(context.Background()).ModelScanResultsActionRequest(modelScanResultsActionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommonApi.NotifyScanResult``: %v\n", err)
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

> UnmaskScanResult(ctx).ModelScanResultsActionRequest(modelScanResultsActionRequest).Execute()

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
    modelScanResultsActionRequest := *openapiclient.NewModelScanResultsActionRequest([]string{"NodeIds_example"}, "ScanType_example") // ModelScanResultsActionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CommonApi.UnmaskScanResult(context.Background()).ModelScanResultsActionRequest(modelScanResultsActionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommonApi.UnmaskScanResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnmaskScanResultRequest struct via the builder pattern


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

