# \ControlsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableAgentPlugin**](ControlsAPI.md#DisableAgentPlugin) | **Post** /deepfence/controls/agent-plugins/disable | Schedule new agent plugin version disabling
[**DisableCloudNodeControls**](ControlsAPI.md#DisableCloudNodeControls) | **Post** /deepfence/controls/cloud-node/disable | Disable Cloud Node Controls
[**EnableAgentPlugin**](ControlsAPI.md#EnableAgentPlugin) | **Post** /deepfence/controls/agent-plugins/enable | Schedule new agent plugin version enabling
[**EnableCloudNodeControls**](ControlsAPI.md#EnableCloudNodeControls) | **Post** /deepfence/controls/cloud-node/enable | Enable Cloud Node Controls
[**GetAgentControls**](ControlsAPI.md#GetAgentControls) | **Post** /deepfence/controls/agent | Fetch Agent Actions
[**GetAgentInitControls**](ControlsAPI.md#GetAgentInitControls) | **Post** /deepfence/controls/agent-init | Fetch Agent Init Actions
[**GetCloudNodeControls**](ControlsAPI.md#GetCloudNodeControls) | **Post** /deepfence/controls/cloud-node | Fetch Cloud Node Controls
[**GetKubernetesClusterControls**](ControlsAPI.md#GetKubernetesClusterControls) | **Post** /deepfence/controls/kubernetes-cluster | Fetch Kubernetes Cluster Actions
[**UpgradeAgentVersion**](ControlsAPI.md#UpgradeAgentVersion) | **Post** /deepfence/controls/agent-upgrade | Schedule new agent version upgrade



## DisableAgentPlugin

> DisableAgentPlugin(ctx).ModelAgentPluginDisable(modelAgentPluginDisable).Execute()

Schedule new agent plugin version disabling



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
    modelAgentPluginDisable := *openapiclient.NewModelAgentPluginDisable("NodeId_example", "PluginName_example") // ModelAgentPluginDisable |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ControlsAPI.DisableAgentPlugin(context.Background()).ModelAgentPluginDisable(modelAgentPluginDisable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControlsAPI.DisableAgentPlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableAgentPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelAgentPluginDisable** | [**ModelAgentPluginDisable**](ModelAgentPluginDisable.md) |  | 

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
    r, err := apiClient.ControlsAPI.DisableCloudNodeControls(context.Background()).ModelCloudNodeEnableDisableReq(modelCloudNodeEnableDisableReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControlsAPI.DisableCloudNodeControls``: %v\n", err)
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


## EnableAgentPlugin

> EnableAgentPlugin(ctx).ModelAgentPluginEnable(modelAgentPluginEnable).Execute()

Schedule new agent plugin version enabling



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
    modelAgentPluginEnable := *openapiclient.NewModelAgentPluginEnable("NodeId_example", "PluginName_example", "Version_example") // ModelAgentPluginEnable |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ControlsAPI.EnableAgentPlugin(context.Background()).ModelAgentPluginEnable(modelAgentPluginEnable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControlsAPI.EnableAgentPlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableAgentPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelAgentPluginEnable** | [**ModelAgentPluginEnable**](ModelAgentPluginEnable.md) |  | 

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
    r, err := apiClient.ControlsAPI.EnableCloudNodeControls(context.Background()).ModelCloudNodeEnableDisableReq(modelCloudNodeEnableDisableReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControlsAPI.EnableCloudNodeControls``: %v\n", err)
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

> ControlsAgentControls GetAgentControls(ctx).ModelAgentID(modelAgentID).Execute()

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
    modelAgentID := *openapiclient.NewModelAgentID(int32(123), "NodeId_example") // ModelAgentID |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ControlsAPI.GetAgentControls(context.Background()).ModelAgentID(modelAgentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControlsAPI.GetAgentControls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentControls`: ControlsAgentControls
    fmt.Fprintf(os.Stdout, "Response from `ControlsAPI.GetAgentControls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentControlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelAgentID** | [**ModelAgentID**](ModelAgentID.md) |  | 

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
    resp, r, err := apiClient.ControlsAPI.GetAgentInitControls(context.Background()).ModelInitAgentReq(modelInitAgentReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControlsAPI.GetAgentInitControls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentInitControls`: ControlsAgentControls
    fmt.Fprintf(os.Stdout, "Response from `ControlsAPI.GetAgentInitControls`: %v\n", resp)
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
    resp, r, err := apiClient.ControlsAPI.GetCloudNodeControls(context.Background()).ModelCloudNodeControlReq(modelCloudNodeControlReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControlsAPI.GetCloudNodeControls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudNodeControls`: ModelCloudNodeControlResp
    fmt.Fprintf(os.Stdout, "Response from `ControlsAPI.GetCloudNodeControls`: %v\n", resp)
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

> ControlsAgentControls GetKubernetesClusterControls(ctx).ModelAgentID(modelAgentID).Execute()

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
    modelAgentID := *openapiclient.NewModelAgentID(int32(123), "NodeId_example") // ModelAgentID |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ControlsAPI.GetKubernetesClusterControls(context.Background()).ModelAgentID(modelAgentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControlsAPI.GetKubernetesClusterControls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesClusterControls`: ControlsAgentControls
    fmt.Fprintf(os.Stdout, "Response from `ControlsAPI.GetKubernetesClusterControls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesClusterControlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelAgentID** | [**ModelAgentID**](ModelAgentID.md) |  | 

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
    r, err := apiClient.ControlsAPI.UpgradeAgentVersion(context.Background()).ModelAgentUpgrade(modelAgentUpgrade).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ControlsAPI.UpgradeAgentVersion``: %v\n", err)
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

