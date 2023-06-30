# \LookupAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCloudCompliances**](LookupAPI.md#GetCloudCompliances) | **Post** /deepfence/lookup/cloud-compliances | Retrieve Cloud Compliances data
[**GetCloudResources**](LookupAPI.md#GetCloudResources) | **Post** /deepfence/lookup/cloud-resources | Get Cloud Resources
[**GetCompliances**](LookupAPI.md#GetCompliances) | **Post** /deepfence/lookup/compliances | Retrieve Compliances data
[**GetContainerImages**](LookupAPI.md#GetContainerImages) | **Post** /deepfence/lookup/containerimages | Retrieve Container Images data
[**GetContainers**](LookupAPI.md#GetContainers) | **Post** /deepfence/lookup/containers | Retrieve Containers data
[**GetHosts**](LookupAPI.md#GetHosts) | **Post** /deepfence/lookup/hosts | Retrieve Hosts data
[**GetKubernetesClusters**](LookupAPI.md#GetKubernetesClusters) | **Post** /deepfence/lookup/kubernetesclusters | Retrieve K8s data
[**GetMalwares**](LookupAPI.md#GetMalwares) | **Post** /deepfence/lookup/malwares | Retrieve Malwares data
[**GetPods**](LookupAPI.md#GetPods) | **Post** /deepfence/lookup/pods | Retrieve Pods data
[**GetProcesses**](LookupAPI.md#GetProcesses) | **Post** /deepfence/lookup/processes | Retrieve Processes data
[**GetRegistryAccount**](LookupAPI.md#GetRegistryAccount) | **Post** /deepfence/lookup/registryaccount | Get Images in Registry
[**GetSecrets**](LookupAPI.md#GetSecrets) | **Post** /deepfence/lookup/secrets | Retrieve Secrets data
[**GetVulnerabilities**](LookupAPI.md#GetVulnerabilities) | **Post** /deepfence/lookup/vulnerabilities | Retrieve Vulnerabilities data



## GetCloudCompliances

> []ModelCloudCompliance GetCloudCompliances(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve Cloud Compliances data



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
    resp, r, err := apiClient.LookupAPI.GetCloudCompliances(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetCloudCompliances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudCompliances`: []ModelCloudCompliance
    fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetCloudCompliances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudCompliancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelCloudCompliance**](ModelCloudCompliance.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    resp, r, err := apiClient.LookupAPI.GetCloudResources(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetCloudResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudResources`: []ModelCloudResource
    fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetCloudResources`: %v\n", resp)
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


## GetCompliances

> []ModelCompliance GetCompliances(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve Compliances data



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
    resp, r, err := apiClient.LookupAPI.GetCompliances(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetCompliances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompliances`: []ModelCompliance
    fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetCompliances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompliancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelCompliance**](ModelCompliance.md)

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
    resp, r, err := apiClient.LookupAPI.GetContainerImages(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetContainerImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerImages`: []ModelContainerImage
    fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetContainerImages`: %v\n", resp)
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
    resp, r, err := apiClient.LookupAPI.GetContainers(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainers`: []ModelContainer
    fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetContainers`: %v\n", resp)
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
    resp, r, err := apiClient.LookupAPI.GetHosts(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHosts`: []ModelHost
    fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetHosts`: %v\n", resp)
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
    resp, r, err := apiClient.LookupAPI.GetKubernetesClusters(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetKubernetesClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesClusters`: []ModelKubernetesCluster
    fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetKubernetesClusters`: %v\n", resp)
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


## GetMalwares

> []ModelMalware GetMalwares(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve Malwares data



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
    resp, r, err := apiClient.LookupAPI.GetMalwares(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetMalwares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMalwares`: []ModelMalware
    fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetMalwares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMalwaresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelMalware**](ModelMalware.md)

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
    resp, r, err := apiClient.LookupAPI.GetPods(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetPods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPods`: []ModelPod
    fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetPods`: %v\n", resp)
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
    resp, r, err := apiClient.LookupAPI.GetProcesses(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetProcesses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProcesses`: []ModelProcess
    fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetProcesses`: %v\n", resp)
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
    resp, r, err := apiClient.LookupAPI.GetRegistryAccount(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetRegistryAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegistryAccount`: []ModelRegistryAccount
    fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetRegistryAccount`: %v\n", resp)
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


## GetSecrets

> []ModelSecret GetSecrets(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve Secrets data



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
    resp, r, err := apiClient.LookupAPI.GetSecrets(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecrets`: []ModelSecret
    fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetSecrets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelSecret**](ModelSecret.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVulnerabilities

> []ModelVulnerability GetVulnerabilities(ctx).LookupLookupFilter(lookupLookupFilter).Execute()

Retrieve Vulnerabilities data



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
    resp, r, err := apiClient.LookupAPI.GetVulnerabilities(context.Background()).LookupLookupFilter(lookupLookupFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupAPI.GetVulnerabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVulnerabilities`: []ModelVulnerability
    fmt.Fprintf(os.Stdout, "Response from `LookupAPI.GetVulnerabilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVulnerabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupLookupFilter** | [**LookupLookupFilter**](LookupLookupFilter.md) |  | 

### Return type

[**[]ModelVulnerability**](ModelVulnerability.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

