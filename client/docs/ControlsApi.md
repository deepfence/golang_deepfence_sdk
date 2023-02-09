# \ControlsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAgentControls**](ControlsApi.md#GetAgentControls) | **Post** /deepfence/controls/agent | Fetch Agent Actions
[**GetAgentInitControls**](ControlsApi.md#GetAgentInitControls) | **Post** /deepfence/controls/agent-init | Fetch Agent Init Actions
[**GetKubernetesClusterControls**](ControlsApi.md#GetKubernetesClusterControls) | **Post** /deepfence/controls/kubernetes-cluster | Fetch Kubernetes Cluster Actions
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
    modelAgentId := *openapiclient.NewModelAgentId(int32(123), "NodeId_example") // ModelAgentId |  (optional)

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

> ControlsAgentControls GetAgentInitControls(ctx).ModelInitAgentReq(modelInitAgentReq).Execute()

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
    modelInitAgentReq := *openapiclient.NewModelInitAgentReq(int32(123), "NodeId_example", "Version_example") // ModelInitAgentReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ControlsApi.GetAgentInitControls(context.Background()).ModelInitAgentReq(modelInitAgentReq).Execute()
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
 **modelInitAgentReq** | [**ModelInitAgentReq**](ModelInitAgentReq.md) |  | 

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


## GetKubernetesClusterControls

> ControlsAgentControls GetKubernetesClusterControls(ctx).ModelAgentId(modelAgentId).Execute()

Fetch Kubernetes Cluster Actions



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
    modelAgentId := *openapiclient.NewModelAgentId(int32(123), "NodeId_example") // ModelAgentId |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ControlsApi.GetKubernetesClusterControls(context.Background()).ModelAgentId(modelAgentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControlsApi.GetKubernetesClusterControls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesClusterControls`: ControlsAgentControls
    fmt.Fprintf(os.Stdout, "Response from `ControlsApi.GetKubernetesClusterControls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesClusterControlsRequest struct via the builder pattern


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

