# \ControlsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAgentControls**](ControlsApi.md#GetAgentControls) | **Post** /deepfence/controls/agent | Fetch Agent Actions
[**GetAgentInitControls**](ControlsApi.md#GetAgentInitControls) | **Post** /deepfence/controls/agent-init | Fetch Agent Init Actions
[**GetKubernetesScannerControls**](ControlsApi.md#GetKubernetesScannerControls) | **Post** /deepfence/controls/kubernetes-scanner | Fetch Kubernetes Scanner Actions
[**UpgradeAgentVersion**](ControlsApi.md#UpgradeAgentVersion) | **Post** /deepfence/controls/agent-upgrade | Schedule new agent version upgrade



## GetAgentControls

> ControlsAgentControls GetAgentControls(ctx).ModelAgentId(modelAgentId).Execute()

Fetch Agent Actions



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
    modelAgentId := *openapiclient.NewModelAgentId("NodeId_example") // ModelAgentId |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ControlsApi.GetAgentControls(context.Background()).ModelAgentId(modelAgentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControlsApi.GetAgentControls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentControls`: ControlsAgentControls
    fmt.Fprintf(os.Stdout, "Response from `ControlsApi.GetAgentControls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentControlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelAgentId** | [**ModelAgentId**](ModelAgentId.md) |  | 

### Return type

[**ControlsAgentControls**](ControlsAgentControls.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentInitControls

> ControlsAgentControls GetAgentInitControls(ctx).ModelAgentId(modelAgentId).Execute()

Fetch Agent Init Actions



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
    modelAgentId := *openapiclient.NewModelAgentId("NodeId_example") // ModelAgentId |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ControlsApi.GetAgentInitControls(context.Background()).ModelAgentId(modelAgentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControlsApi.GetAgentInitControls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentInitControls`: ControlsAgentControls
    fmt.Fprintf(os.Stdout, "Response from `ControlsApi.GetAgentInitControls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentInitControlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelAgentId** | [**ModelAgentId**](ModelAgentId.md) |  | 

### Return type

[**ControlsAgentControls**](ControlsAgentControls.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesScannerControls

> ControlsKubernetesScannerControlResponse GetKubernetesScannerControls(ctx).ModelAgentId(modelAgentId).Execute()

Fetch Kubernetes Scanner Actions



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
    modelAgentId := *openapiclient.NewModelAgentId("NodeId_example") // ModelAgentId |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ControlsApi.GetKubernetesScannerControls(context.Background()).ModelAgentId(modelAgentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControlsApi.GetKubernetesScannerControls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesScannerControls`: ControlsKubernetesScannerControlResponse
    fmt.Fprintf(os.Stdout, "Response from `ControlsApi.GetKubernetesScannerControls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesScannerControlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelAgentId** | [**ModelAgentId**](ModelAgentId.md) |  | 

### Return type

[**ControlsKubernetesScannerControlResponse**](ControlsKubernetesScannerControlResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeAgentVersion

> UpgradeAgentVersion(ctx).ModelAgentUpgrade(modelAgentUpgrade).Execute()

Schedule new agent version upgrade



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
    modelAgentUpgrade := *openapiclient.NewModelAgentUpgrade("NodeId_example", "Version_example") // ModelAgentUpgrade |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ControlsApi.UpgradeAgentVersion(context.Background()).ModelAgentUpgrade(modelAgentUpgrade).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControlsApi.UpgradeAgentVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeAgentVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelAgentUpgrade** | [**ModelAgentUpgrade**](ModelAgentUpgrade.md) |  | 

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

