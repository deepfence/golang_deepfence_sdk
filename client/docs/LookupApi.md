# \LookupApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContainerImages**](LookupApi.md#GetContainerImages) | **Post** /deepfence/lookup/containerimages | Retrieve Container Images data
[**GetContainers**](LookupApi.md#GetContainers) | **Post** /deepfence/lookup/containers | Retrieve Containers data
[**GetHosts**](LookupApi.md#GetHosts) | **Post** /deepfence/lookup/hosts | Retrieve Hosts data
[**GetKubernetesClusters**](LookupApi.md#GetKubernetesClusters) | **Post** /deepfence/lookup/kubernetesclusters | Retrieve K8s data
[**GetPods**](LookupApi.md#GetPods) | **Post** /deepfence/lookup/pods | Retrieve Pods data
[**GetProcesses**](LookupApi.md#GetProcesses) | **Post** /deepfence/lookup/processes | Retrieve Processes data
[**GetRegistryAccount**](LookupApi.md#GetRegistryAccount) | **Post** /deepfence/lookup/registryaccount | Get Images in Registry
[**SearchCloudCompliances**](LookupApi.md#SearchCloudCompliances) | **Post** /deepfence/search/cloud-compliances | Search Cloud compliances
[**SearchCompliances**](LookupApi.md#SearchCompliances) | **Post** /deepfence/search/compliances | Search Compliances
[**SearchContainerImages**](LookupApi.md#SearchContainerImages) | **Post** /deepfence/search/images | Search Container images
[**SearchContainers**](LookupApi.md#SearchContainers) | **Post** /deepfence/search/containers | Search Containers data
[**SearchHosts**](LookupApi.md#SearchHosts) | **Post** /deepfence/search/hosts | Search hosts
[**SearchMalwares**](LookupApi.md#SearchMalwares) | **Post** /deepfence/search/malwares | Search Malwares
[**SearchSecrets**](LookupApi.md#SearchSecrets) | **Post** /deepfence/search/secrets | Search Secrets
[**SearchVulnerabilities**](LookupApi.md#SearchVulnerabilities) | **Post** /deepfence/search/vulnerabilities | Search Vulnerabilities



## GetContainerImages

> []ModelContainerImage GetContainerImages(ctx).ReportersLookupFilter(reportersLookupFilter).Execute()

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
    reportersLookupFilter := *openapiclient.NewReportersLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}) // ReportersLookupFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LookupApi.GetContainerImages(context.Background()).ReportersLookupFilter(reportersLookupFilter).Execute()
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
 **reportersLookupFilter** | [**ReportersLookupFilter**](ReportersLookupFilter.md) |  | 

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

> []ModelContainer GetContainers(ctx).ReportersLookupFilter(reportersLookupFilter).Execute()

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
    reportersLookupFilter := *openapiclient.NewReportersLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}) // ReportersLookupFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LookupApi.GetContainers(context.Background()).ReportersLookupFilter(reportersLookupFilter).Execute()
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
 **reportersLookupFilter** | [**ReportersLookupFilter**](ReportersLookupFilter.md) |  | 

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

> []ModelHost GetHosts(ctx).ReportersLookupFilter(reportersLookupFilter).Execute()

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
    reportersLookupFilter := *openapiclient.NewReportersLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}) // ReportersLookupFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LookupApi.GetHosts(context.Background()).ReportersLookupFilter(reportersLookupFilter).Execute()
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
 **reportersLookupFilter** | [**ReportersLookupFilter**](ReportersLookupFilter.md) |  | 

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

> []ModelKubernetesCluster GetKubernetesClusters(ctx).ReportersLookupFilter(reportersLookupFilter).Execute()

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
    reportersLookupFilter := *openapiclient.NewReportersLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}) // ReportersLookupFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LookupApi.GetKubernetesClusters(context.Background()).ReportersLookupFilter(reportersLookupFilter).Execute()
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
 **reportersLookupFilter** | [**ReportersLookupFilter**](ReportersLookupFilter.md) |  | 

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

> []ModelPod GetPods(ctx).ReportersLookupFilter(reportersLookupFilter).Execute()

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
    reportersLookupFilter := *openapiclient.NewReportersLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}) // ReportersLookupFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LookupApi.GetPods(context.Background()).ReportersLookupFilter(reportersLookupFilter).Execute()
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
 **reportersLookupFilter** | [**ReportersLookupFilter**](ReportersLookupFilter.md) |  | 

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

> []ModelProcess GetProcesses(ctx).ReportersLookupFilter(reportersLookupFilter).Execute()

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
    reportersLookupFilter := *openapiclient.NewReportersLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}) // ReportersLookupFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LookupApi.GetProcesses(context.Background()).ReportersLookupFilter(reportersLookupFilter).Execute()
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
 **reportersLookupFilter** | [**ReportersLookupFilter**](ReportersLookupFilter.md) |  | 

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

> []ModelRegistryAccount GetRegistryAccount(ctx).ReportersLookupFilter(reportersLookupFilter).Execute()

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
    reportersLookupFilter := *openapiclient.NewReportersLookupFilter([]string{"InFieldFilter_example"}, []string{"NodeIds_example"}) // ReportersLookupFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LookupApi.GetRegistryAccount(context.Background()).ReportersLookupFilter(reportersLookupFilter).Execute()
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
 **reportersLookupFilter** | [**ReportersLookupFilter**](ReportersLookupFilter.md) |  | 

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


## SearchCloudCompliances

> []ModelCloudCompliance SearchCloudCompliances(ctx).ReportersSearchFilter(reportersSearchFilter).Execute()

Search Cloud compliances



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
    reportersSearchFilter := *openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}) // ReportersSearchFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LookupApi.SearchCloudCompliances(context.Background()).ReportersSearchFilter(reportersSearchFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupApi.SearchCloudCompliances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCloudCompliances`: []ModelCloudCompliance
    fmt.Fprintf(os.Stdout, "Response from `LookupApi.SearchCloudCompliances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCloudCompliancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportersSearchFilter** | [**ReportersSearchFilter**](ReportersSearchFilter.md) |  | 

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


## SearchCompliances

> []ModelCompliance SearchCompliances(ctx).ReportersSearchFilter(reportersSearchFilter).Execute()

Search Compliances



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
    reportersSearchFilter := *openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}) // ReportersSearchFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LookupApi.SearchCompliances(context.Background()).ReportersSearchFilter(reportersSearchFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupApi.SearchCompliances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCompliances`: []ModelCompliance
    fmt.Fprintf(os.Stdout, "Response from `LookupApi.SearchCompliances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCompliancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportersSearchFilter** | [**ReportersSearchFilter**](ReportersSearchFilter.md) |  | 

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


## SearchContainerImages

> []ModelContainerImage SearchContainerImages(ctx).ReportersSearchFilter(reportersSearchFilter).Execute()

Search Container images



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
    reportersSearchFilter := *openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}) // ReportersSearchFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LookupApi.SearchContainerImages(context.Background()).ReportersSearchFilter(reportersSearchFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupApi.SearchContainerImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchContainerImages`: []ModelContainerImage
    fmt.Fprintf(os.Stdout, "Response from `LookupApi.SearchContainerImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchContainerImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportersSearchFilter** | [**ReportersSearchFilter**](ReportersSearchFilter.md) |  | 

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


## SearchContainers

> []ModelContainer SearchContainers(ctx).ReportersSearchFilter(reportersSearchFilter).Execute()

Search Containers data



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
    reportersSearchFilter := *openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}) // ReportersSearchFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LookupApi.SearchContainers(context.Background()).ReportersSearchFilter(reportersSearchFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupApi.SearchContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchContainers`: []ModelContainer
    fmt.Fprintf(os.Stdout, "Response from `LookupApi.SearchContainers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportersSearchFilter** | [**ReportersSearchFilter**](ReportersSearchFilter.md) |  | 

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


## SearchHosts

> []ModelHost SearchHosts(ctx).ReportersSearchFilter(reportersSearchFilter).Execute()

Search hosts



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
    reportersSearchFilter := *openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}) // ReportersSearchFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LookupApi.SearchHosts(context.Background()).ReportersSearchFilter(reportersSearchFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupApi.SearchHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchHosts`: []ModelHost
    fmt.Fprintf(os.Stdout, "Response from `LookupApi.SearchHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportersSearchFilter** | [**ReportersSearchFilter**](ReportersSearchFilter.md) |  | 

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


## SearchMalwares

> []ModelMalware SearchMalwares(ctx).ReportersSearchFilter(reportersSearchFilter).Execute()

Search Malwares



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
    reportersSearchFilter := *openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}) // ReportersSearchFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LookupApi.SearchMalwares(context.Background()).ReportersSearchFilter(reportersSearchFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupApi.SearchMalwares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchMalwares`: []ModelMalware
    fmt.Fprintf(os.Stdout, "Response from `LookupApi.SearchMalwares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchMalwaresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportersSearchFilter** | [**ReportersSearchFilter**](ReportersSearchFilter.md) |  | 

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


## SearchSecrets

> []ModelSecret SearchSecrets(ctx).ReportersSearchFilter(reportersSearchFilter).Execute()

Search Secrets



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
    reportersSearchFilter := *openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}) // ReportersSearchFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LookupApi.SearchSecrets(context.Background()).ReportersSearchFilter(reportersSearchFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupApi.SearchSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSecrets`: []ModelSecret
    fmt.Fprintf(os.Stdout, "Response from `LookupApi.SearchSecrets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportersSearchFilter** | [**ReportersSearchFilter**](ReportersSearchFilter.md) |  | 

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


## SearchVulnerabilities

> []ModelVulnerability SearchVulnerabilities(ctx).ReportersSearchFilter(reportersSearchFilter).Execute()

Search Vulnerabilities



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
    reportersSearchFilter := *openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}) // ReportersSearchFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LookupApi.SearchVulnerabilities(context.Background()).ReportersSearchFilter(reportersSearchFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupApi.SearchVulnerabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchVulnerabilities`: []ModelVulnerability
    fmt.Fprintf(os.Stdout, "Response from `LookupApi.SearchVulnerabilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchVulnerabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportersSearchFilter** | [**ReportersSearchFilter**](ReportersSearchFilter.md) |  | 

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

