# \RegistryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRegistry**](RegistryApi.md#AddRegistry) | **Post** /deepfence/registryaccount | Add Registry
[**DeleteRegistry**](RegistryApi.md#DeleteRegistry) | **Delete** /deepfence/registryaccount/{registry_id} | Add Registry
[**GetRegistrySummary**](RegistryApi.md#GetRegistrySummary) | **Get** /deepfence/registryaccount/{registry_id}/summary | Get Registry Summary
[**GetSummary**](RegistryApi.md#GetSummary) | **Get** /deepfence/registryaccount/summary | Get All Registries Summary
[**ListImageTags**](RegistryApi.md#ListImageTags) | **Get** /deepfence/registryaccount/{registry_id}/images/{image_name}/tags | List Image Tags
[**ListImages**](RegistryApi.md#ListImages) | **Get** /deepfence/registryaccount/{registry_id}/images | List Registry Images
[**ListRegistry**](RegistryApi.md#ListRegistry) | **Get** /deepfence/registryaccount | List Registries



## AddRegistry

> AddRegistry(ctx).ModelRegistryAddReq(modelRegistryAddReq).Execute()

Add Registry



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
    modelRegistryAddReq := *openapiclient.NewModelRegistryAddReq() // ModelRegistryAddReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RegistryApi.AddRegistry(context.Background()).ModelRegistryAddReq(modelRegistryAddReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.AddRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelRegistryAddReq** | [**ModelRegistryAddReq**](ModelRegistryAddReq.md) |  | 

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


## DeleteRegistry

> DeleteRegistry(ctx, registryId).Execute()

Add Registry



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
    registryId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RegistryApi.DeleteRegistry(context.Background(), registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.DeleteRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistrySummary

> map[string]int32 GetRegistrySummary(ctx, registryId).Execute()

Get Registry Summary



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
    registryId := "registryId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.GetRegistrySummary(context.Background(), registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.GetRegistrySummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegistrySummary`: map[string]int32
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.GetRegistrySummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistrySummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]int32**

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSummary

> map[string]int32 GetSummary(ctx).Execute()

Get All Registries Summary



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
    resp, r, err := apiClient.RegistryApi.GetSummary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.GetSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSummary`: map[string]int32
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.GetSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSummaryRequest struct via the builder pattern


### Return type

**map[string]int32**

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImageTags

> []ModelContainerImage ListImageTags(ctx, registryId, imageName).Execute()

List Image Tags



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
    registryId := "registryId_example" // string | 
    imageName := "imageName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.ListImageTags(context.Background(), registryId, imageName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.ListImageTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImageTags`: []ModelContainerImage
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.ListImageTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** |  | 
**imageName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListImageTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ModelContainerImage**](ModelContainerImage.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImages

> []ModelContainerImageWithTags ListImages(ctx, registryId).Execute()

List Registry Images



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
    registryId := "registryId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.ListImages(context.Background(), registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.ListImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImages`: []ModelContainerImageWithTags
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.ListImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ModelContainerImageWithTags**](ModelContainerImageWithTags.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRegistry

> []ModelRegistryListResp ListRegistry(ctx).Execute()

List Registries



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
    resp, r, err := apiClient.RegistryApi.ListRegistry(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.ListRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRegistry`: []ModelRegistryListResp
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.ListRegistry`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRegistryRequest struct via the builder pattern


### Return type

[**[]ModelRegistryListResp**](ModelRegistryListResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

