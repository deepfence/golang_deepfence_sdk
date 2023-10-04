# \CommonAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Eula**](CommonAPI.md#Eula) | **Get** /deepfence/end-user-license-agreement | Get End User License Agreement
[**GetScanReportFields**](CommonAPI.md#GetScanReportFields) | **Get** /deepfence/scan/results/fields | Get Scan Report Fields



## Eula

> ModelMessageResponse Eula(ctx).Execute()

Get End User License Agreement



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommonAPI.Eula(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommonAPI.Eula``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Eula`: ModelMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `CommonAPI.Eula`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEulaRequest struct via the builder pattern


### Return type

[**ModelMessageResponse**](ModelMessageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScanReportFields

> ModelScanReportFieldsResponse GetScanReportFields(ctx).Execute()

Get Scan Report Fields



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommonAPI.GetScanReportFields(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommonAPI.GetScanReportFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScanReportFields`: ModelScanReportFieldsResponse
    fmt.Fprintf(os.Stdout, "Response from `CommonAPI.GetScanReportFields`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetScanReportFieldsRequest struct via the builder pattern


### Return type

[**ModelScanReportFieldsResponse**](ModelScanReportFieldsResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

