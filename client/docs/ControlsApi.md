# \ControlsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableCloudNodeControls**](ControlsApi.md#DisableCloudNodeControls) | **Post** /deepfence/controls/cloud-node/disable | Disable Cloud Node Controls
[**EnableCloudNodeControls**](ControlsApi.md#EnableCloudNodeControls) | **Post** /deepfence/controls/cloud-node/enable | Enable Cloud Node Controls
[**GetAgentControls**](ControlsApi.md#GetAgentControls) | **Post** /deepfence/controls/agent | Fetch Agent Actions
[**GetAgentInitControls**](ControlsApi.md#GetAgentInitControls) | **Post** /deepfence/controls/agent-init | Fetch Agent Init Actions
[**GetCloudNodeControls**](ControlsApi.md#GetCloudNodeControls) | **Post** /deepfence/controls/cloud-node | Fetch Cloud Node Controls
[**GetKubernetesClusterControls**](ControlsApi.md#GetKubernetesClusterControls) | **Post** /deepfence/controls/kubernetes-cluster | Fetch Kubernetes Cluster Actions
[**UpgradeAgentVersion**](ControlsApi.md#UpgradeAgentVersion) | **Post** /deepfence/controls/agent-upgrade | Schedule new agent version upgrade



## DisableCloudNodeControls

> DisableCloudNodeControls(ctx).ModelCloudNodeEnableDisableReq(modelCloudNodeEnableDisableReq).Execute()

Disable Cloud Node Controls



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
    modelCloudNodeEnableDisableReq := *openapiclient.NewModelCloudNodeEnableDisableReq() // ModelCloudNodeEnableDisableReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ControlsApi.DisableCloudNodeControls(context.Background()).ModelCloudNodeEnableDisableReq(modelCloudNodeEnableDisableReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControlsApi.DisableCloudNodeControls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableCloudNodeControlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelCloudNodeEnableDisableReq** | [**ModelCloudNodeEnableDisableReq**](ModelCloudNodeEnableDisableReq.md) |  | 

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


## EnableCloudNodeControls

> EnableCloudNodeControls(ctx).ModelCloudNodeEnableDisableReq(modelCloudNodeEnableDisableReq).Execute()

Enable Cloud Node Controls



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
    modelCloudNodeEnableDisableReq := *openapiclient.NewModelCloudNodeEnableDisableReq() // ModelCloudNodeEnableDisableReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ControlsApi.EnableCloudNodeControls(context.Background()).ModelCloudNodeEnableDisableReq(modelCloudNodeEnableDisableReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControlsApi.EnableCloudNodeControls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableCloudNodeControlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelCloudNodeEnableDisableReq** | [**ModelCloudNodeEnableDisableReq**](ModelCloudNodeEnableDisableReq.md) |  | 

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
    openapiclient "github.com/deepfence/golang_deepfence_sdk/client"
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
    openapiclient "github.com/deepfence/golang_deepfence_sdk/client"
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


## GetCloudNodeControls

> ModelCloudNodeControlResp GetCloudNodeControls(ctx).ModelCloudNodeControlReq(modelCloudNodeControlReq).Execute()

Fetch Cloud Node Controls



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
    modelCloudNodeControlReq := *openapiclient.NewModelCloudNodeControlReq("CloudProvider_example", "ComplianceType_example") // ModelCloudNodeControlReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ControlsApi.GetCloudNodeControls(context.Background()).ModelCloudNodeControlReq(modelCloudNodeControlReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControlsApi.GetCloudNodeControls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudNodeControls`: ModelCloudNodeControlResp
    fmt.Fprintf(os.Stdout, "Response from `ControlsApi.GetCloudNodeControls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudNodeControlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelCloudNodeControlReq** | [**ModelCloudNodeControlReq**](ModelCloudNodeControlReq.md) |  | 

### Return type

[**ModelCloudNodeControlResp**](ModelCloudNodeControlResp.md)

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
    openapiclient "github.com/deepfence/golang_deepfence_sdk/client"
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
    openapiclient "github.com/deepfence/golang_deepfence_sdk/client"
)

func main() {
    modelAgentUpgrade := *openapiclient.NewModelAgentUpgrade("NodeId_example", "Version_example") // ModelAgentUpgrade |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ControlsApi.UpgradeAgentVersion(context.Background()).ModelAgentUpgrade(modelAgentUpgrade).Execute()
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

