# \LookupApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCloudResources**](LookupApi.md#GetCloudResources) | **Post** /deepfence/lookup/cloud-resources | Get Cloud Resources
[**GetContainerImages**](LookupApi.md#GetContainerImages) | **Post** /deepfence/lookup/containerimages | Retrieve Container Images data
[**GetContainers**](LookupApi.md#GetContainers) | **Post** /deepfence/lookup/containers | Retrieve Containers data
[**GetHosts**](LookupApi.md#GetHosts) | **Post** /deepfence/lookup/hosts | Retrieve Hosts data
[**GetKubernetesClusters**](LookupApi.md#GetKubernetesClusters) | **Post** /deepfence/lookup/kubernetesclusters | Retrieve K8s data
[**GetPods**](LookupApi.md#GetPods) | **Post** /deepfence/lookup/pods | Retrieve Pods data
[**GetProcesses**](LookupApi.md#GetProcesses) | **Post** /deepfence/lookup/processes | Retrieve Processes data
[**GetRegistryAccount**](LookupApi.md#GetRegistryAccount) | **Post** /deepfence/lookup/registryaccount | Get Images in Registry



## GetCloudResources

> []ModelCloudResource GetCloudResources(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Get Cloud Resources



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
    lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LookupApi.GetCloudResources(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupApi.GetCloudResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudResources`: []ModelCloudResource
    fmt.Fprintf(os.Stdout, "Response from `LookupApi.GetCloudResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelCloudResource**](ModelCloudResource.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerImages

> []ModelContainerImage GetContainerImages(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve Container Images data



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
    lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LookupApi.GetContainerImages(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupApi.GetContainerImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerImages`: []ModelContainerImage
    fmt.Fprintf(os.Stdout, "Response from `LookupApi.GetContainerImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelContainerImage**](ModelContainerImage.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainers

> []ModelContainer GetContainers(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve Containers data



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
    lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LookupApi.GetContainers(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupApi.GetContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainers`: []ModelContainer
    fmt.Fprintf(os.Stdout, "Response from `LookupApi.GetContainers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelContainer**](ModelContainer.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHosts

> []ModelHost GetHosts(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve Hosts data



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
    lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LookupApi.GetHosts(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupApi.GetHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHosts`: []ModelHost
    fmt.Fprintf(os.Stdout, "Response from `LookupApi.GetHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelHost**](ModelHost.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesClusters

> []ModelKubernetesCluster GetKubernetesClusters(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve K8s data



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
    lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LookupApi.GetKubernetesClusters(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupApi.GetKubernetesClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesClusters`: []ModelKubernetesCluster
    fmt.Fprintf(os.Stdout, "Response from `LookupApi.GetKubernetesClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelKubernetesCluster**](ModelKubernetesCluster.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPods

> []ModelPod GetPods(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve Pods data



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
    lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LookupApi.GetPods(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupApi.GetPods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPods`: []ModelPod
    fmt.Fprintf(os.Stdout, "Response from `LookupApi.GetPods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelPod**](ModelPod.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcesses

> []ModelProcess GetProcesses(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve Processes data



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
    lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LookupApi.GetProcesses(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupApi.GetProcesses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProcesses`: []ModelProcess
    fmt.Fprintf(os.Stdout, "Response from `LookupApi.GetProcesses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelProcess**](ModelProcess.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistryAccount

> []ModelRegistryAccount GetRegistryAccount(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Get Images in Registry



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
    lookupLookupFilter := *openapiclient.NewLookupLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // LookupLookupFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LookupApi.GetRegistryAccount(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupApi.GetRegistryAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegistryAccount`: []ModelRegistryAccount
    fmt.Fprintf(os.Stdout, "Response from `LookupApi.GetRegistryAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistryAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelRegistryAccount**](ModelRegistryAccount.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

