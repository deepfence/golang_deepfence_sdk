# \KubernetesScannerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegisterKubernetesScanner**](KubernetesScannerApi.md#RegisterKubernetesScanner) | **Post** /deepfence/ingest/kubernetes-scanner | Register Kubernetes Scanner



## RegisterKubernetesScanner

> RegisterKubernetesScanner(ctx).IngestersRegisterKubernetesScannerRequest(ingestersRegisterKubernetesScannerRequest).Execute()

Register Kubernetes Scanner



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
    ingestersRegisterKubernetesScannerRequest := []openapiclient.IngestersRegisterKubernetesScannerRequest{*openapiclient.NewIngestersRegisterKubernetesScannerRequest()} // []IngestersRegisterKubernetesScannerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesScannerApi.RegisterKubernetesScanner(context.Background()).IngestersRegisterKubernetesScannerRequest(ingestersRegisterKubernetesScannerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesScannerApi.RegisterKubernetesScanner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterKubernetesScannerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestersRegisterKubernetesScannerRequest** | [**[]IngestersRegisterKubernetesScannerRequest**](IngestersRegisterKubernetesScannerRequest.md) |  | 

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

