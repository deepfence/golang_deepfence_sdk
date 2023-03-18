# \RegistryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRegistry**](RegistryApi.md#AddRegistry) | **Post** /deepfence/registryaccount | Add Registry
[**CountImageStubs**](RegistryApi.md#CountImageStubs) | **Post** /deepfence/registryaccount/count/stubs | Count Image Stubs
[**CountImages**](RegistryApi.md#CountImages) | **Post** /deepfence/registryaccount/count/images | Count Registry Images
[**DeleteRegistry**](RegistryApi.md#DeleteRegistry) | **Delete** /deepfence/registryaccount/{registry_id} | Delete Registry
[**GetRegistrySummary**](RegistryApi.md#GetRegistrySummary) | **Get** /deepfence/registryaccount/{registry_id}/summary | Get Registry Summary
[**GetSummaryAll**](RegistryApi.md#GetSummaryAll) | **Get** /deepfence/registryaccount/summary | Get All Registries Summary By Type
[**GetSummaryByType**](RegistryApi.md#GetSummaryByType) | **Get** /deepfence/registryaccount/{registry_type}/summary-by-type | Get Registry Summary By Type
[**ListImageStubs**](RegistryApi.md#ListImageStubs) | **Post** /deepfence/registryaccount/stubs | List Image Stubs
[**ListImages**](RegistryApi.md#ListImages) | **Post** /deepfence/registryaccount/images | List Registry Images
[**ListRegistry**](RegistryApi.md#ListRegistry) | **Get** /deepfence/registryaccount | List Registries
[**UpdateRegistry**](RegistryApi.md#UpdateRegistry) | **Put** /deepfence/registryaccount/{registry_id} | Update Registry



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
    modelRegistryAddReq := *openapiclient.NewModelRegistryAddReq("Name_example", "RegistryType_example") // ModelRegistryAddReq |  (optional)

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


## CountImageStubs

> ModelRegistryCountResp CountImageStubs(ctx).ModelRegistryImageStubsReq(modelRegistryImageStubsReq).Execute()

Count Image Stubs



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
    modelRegistryImageStubsReq := *openapiclient.NewModelRegistryImageStubsReq(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), "RegistryId_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ModelRegistryImageStubsReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.CountImageStubs(context.Background()).ModelRegistryImageStubsReq(modelRegistryImageStubsReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.CountImageStubs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountImageStubs`: ModelRegistryCountResp
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.CountImageStubs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountImageStubsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelRegistryImageStubsReq** | [**ModelRegistryImageStubsReq**](ModelRegistryImageStubsReq.md) |  | 

### Return type

[**ModelRegistryCountResp**](ModelRegistryCountResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountImages

> ModelRegistryCountResp CountImages(ctx).ModelRegistryImagesReq(modelRegistryImagesReq).Execute()

Count Registry Images



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
    modelRegistryImagesReq := *openapiclient.NewModelRegistryImagesReq(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), "RegistryId_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ModelRegistryImagesReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.CountImages(context.Background()).ModelRegistryImagesReq(modelRegistryImagesReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.CountImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountImages`: ModelRegistryCountResp
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.CountImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelRegistryImagesReq** | [**ModelRegistryImagesReq**](ModelRegistryImagesReq.md) |  | 

### Return type

[**ModelRegistryCountResp**](ModelRegistryCountResp.md)

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

Delete Registry



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
**registryId** | **string** |  | 

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

> ModelSummary GetRegistrySummary(ctx, registryId).Execute()

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
    // response from `GetRegistrySummary`: ModelSummary
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

[**ModelSummary**](ModelSummary.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSummaryAll

> map[string]ModelSummary GetSummaryAll(ctx).Execute()

Get All Registries Summary By Type



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
    resp, r, err := apiClient.RegistryApi.GetSummaryAll(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.GetSummaryAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSummaryAll`: map[string]ModelSummary
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.GetSummaryAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSummaryAllRequest struct via the builder pattern


### Return type

[**map[string]ModelSummary**](ModelSummary.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSummaryByType

> ModelSummary GetSummaryByType(ctx, registryType).Execute()

Get Registry Summary By Type



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
    registryType := "registryType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.GetSummaryByType(context.Background(), registryType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.GetSummaryByType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSummaryByType`: ModelSummary
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.GetSummaryByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSummaryByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelSummary**](ModelSummary.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImageStubs

> []ModelImageStub ListImageStubs(ctx).ModelRegistryImageStubsReq(modelRegistryImageStubsReq).Execute()

List Image Stubs



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
    modelRegistryImageStubsReq := *openapiclient.NewModelRegistryImageStubsReq(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), "RegistryId_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ModelRegistryImageStubsReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.ListImageStubs(context.Background()).ModelRegistryImageStubsReq(modelRegistryImageStubsReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.ListImageStubs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImageStubs`: []ModelImageStub
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.ListImageStubs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListImageStubsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelRegistryImageStubsReq** | [**ModelRegistryImageStubsReq**](ModelRegistryImageStubsReq.md) |  | 

### Return type

[**[]ModelImageStub**](ModelImageStub.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImages

> []ModelContainerImage ListImages(ctx).ModelRegistryImagesReq(modelRegistryImagesReq).Execute()

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
    modelRegistryImagesReq := *openapiclient.NewModelRegistryImagesReq(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), "RegistryId_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ModelRegistryImagesReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegistryApi.ListImages(context.Background()).ModelRegistryImagesReq(modelRegistryImagesReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.ListImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImages`: []ModelContainerImage
    fmt.Fprintf(os.Stdout, "Response from `RegistryApi.ListImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelRegistryImagesReq** | [**ModelRegistryImagesReq**](ModelRegistryImagesReq.md) |  | 

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


## UpdateRegistry

> UpdateRegistry(ctx, registryId).ModelRegistryUpdateReq(modelRegistryUpdateReq).Execute()

Update Registry



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
    modelRegistryUpdateReq := *openapiclient.NewModelRegistryUpdateReq() // ModelRegistryUpdateReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RegistryApi.UpdateRegistry(context.Background(), registryId).ModelRegistryUpdateReq(modelRegistryUpdateReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistryApi.UpdateRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelRegistryUpdateReq** | [**ModelRegistryUpdateReq**](ModelRegistryUpdateReq.md) |  | 

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

