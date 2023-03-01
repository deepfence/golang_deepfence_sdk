# \SecretScanApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountResultsSecretScan**](SecretScanApi.md#CountResultsSecretScan) | **Post** /deepfence/scan/results/count/secret | Get Secret Scans Results
[**IngestSecretScanStatus**](SecretScanApi.md#IngestSecretScanStatus) | **Post** /deepfence/ingest/secret-scan-logs | Ingest Secrets Scan Status
[**IngestSecrets**](SecretScanApi.md#IngestSecrets) | **Post** /deepfence/ingest/secrets | Ingest Secrets
[**ListSecretScan**](SecretScanApi.md#ListSecretScan) | **Post** /deepfence/scan/list/secret | Get Secret Scans List
[**ResultsSecretScan**](SecretScanApi.md#ResultsSecretScan) | **Post** /deepfence/scan/results/secret | Get Secret Scans Results
[**StartSecretScan**](SecretScanApi.md#StartSecretScan) | **Post** /deepfence/scan/start/secret | Start Secret Scan
[**StatusSecretScan**](SecretScanApi.md#StatusSecretScan) | **Post** /deepfence/scan/status/secret | Get Secret Scan Status
[**StopSecretScan**](SecretScanApi.md#StopSecretScan) | **Post** /deepfence/scan/stop/secret | Stop Secret Scan



## CountResultsSecretScan

> SearchSearchCountResp CountResultsSecretScan(ctx).ModelScanResultsReq(modelScanResultsReq).Execute()

Get Secret Scans Results



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
    modelScanResultsReq := *openapiclient.NewModelScanResultsReq(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]string{"OrderFields_example"})), "ScanId_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ModelScanResultsReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretScanApi.CountResultsSecretScan(context.Background()).ModelScanResultsReq(modelScanResultsReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretScanApi.CountResultsSecretScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountResultsSecretScan`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SecretScanApi.CountResultsSecretScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountResultsSecretScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelScanResultsReq** | [**ModelScanResultsReq**](ModelScanResultsReq.md) |  | 

### Return type

[**SearchSearchCountResp**](SearchSearchCountResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IngestSecretScanStatus

> IngestSecretScanStatus(ctx).IngestersSecretScanStatus(ingestersSecretScanStatus).Execute()

Ingest Secrets Scan Status



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
    ingestersSecretScanStatus := []openapiclient.IngestersSecretScanStatus{*openapiclient.NewIngestersSecretScanStatus()} // []IngestersSecretScanStatus |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SecretScanApi.IngestSecretScanStatus(context.Background()).IngestersSecretScanStatus(ingestersSecretScanStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretScanApi.IngestSecretScanStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestSecretScanStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestersSecretScanStatus** | [**[]IngestersSecretScanStatus**](IngestersSecretScanStatus.md) |  | 

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


## IngestSecrets

> IngestSecrets(ctx).IngestersSecret(ingestersSecret).Execute()

Ingest Secrets



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
    ingestersSecret := []openapiclient.IngestersSecret{*openapiclient.NewIngestersSecret()} // []IngestersSecret |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SecretScanApi.IngestSecrets(context.Background()).IngestersSecret(ingestersSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretScanApi.IngestSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestersSecret** | [**[]IngestersSecret**](IngestersSecret.md) |  | 

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


## ListSecretScan

> ModelScanListResp ListSecretScan(ctx).ModelScanListReq(modelScanListReq).Execute()

Get Secret Scans List



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
    modelScanListReq := *openapiclient.NewModelScanListReq([]openapiclient.ModelNodeIdentifier{*openapiclient.NewModelNodeIdentifier("NodeId_example", "NodeType_example")}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ModelScanListReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretScanApi.ListSecretScan(context.Background()).ModelScanListReq(modelScanListReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretScanApi.ListSecretScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSecretScan`: ModelScanListResp
    fmt.Fprintf(os.Stdout, "Response from `SecretScanApi.ListSecretScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSecretScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelScanListReq** | [**ModelScanListReq**](ModelScanListReq.md) |  | 

### Return type

[**ModelScanListResp**](ModelScanListResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResultsSecretScan

> ModelSecretScanResult ResultsSecretScan(ctx).ModelScanResultsReq(modelScanResultsReq).Execute()

Get Secret Scans Results



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
    modelScanResultsReq := *openapiclient.NewModelScanResultsReq(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]string{"OrderFields_example"})), "ScanId_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ModelScanResultsReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretScanApi.ResultsSecretScan(context.Background()).ModelScanResultsReq(modelScanResultsReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretScanApi.ResultsSecretScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResultsSecretScan`: ModelSecretScanResult
    fmt.Fprintf(os.Stdout, "Response from `SecretScanApi.ResultsSecretScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResultsSecretScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelScanResultsReq** | [**ModelScanResultsReq**](ModelScanResultsReq.md) |  | 

### Return type

[**ModelSecretScanResult**](ModelSecretScanResult.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartSecretScan

> ModelScanTriggerResp StartSecretScan(ctx).ModelSecretScanTriggerReq(modelSecretScanTriggerReq).Execute()

Start Secret Scan



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
    modelSecretScanTriggerReq := *openapiclient.NewModelSecretScanTriggerReq(*openapiclient.NewModelScanFilter(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}})), []openapiclient.ModelNodeIdentifier{*openapiclient.NewModelNodeIdentifier("NodeId_example", "NodeType_example")}) // ModelSecretScanTriggerReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretScanApi.StartSecretScan(context.Background()).ModelSecretScanTriggerReq(modelSecretScanTriggerReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretScanApi.StartSecretScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartSecretScan`: ModelScanTriggerResp
    fmt.Fprintf(os.Stdout, "Response from `SecretScanApi.StartSecretScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartSecretScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelSecretScanTriggerReq** | [**ModelSecretScanTriggerReq**](ModelSecretScanTriggerReq.md) |  | 

### Return type

[**ModelScanTriggerResp**](ModelScanTriggerResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatusSecretScan

> ModelScanStatusResp StatusSecretScan(ctx).ModelScanStatusReq(modelScanStatusReq).Execute()

Get Secret Scan Status



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
    modelScanStatusReq := *openapiclient.NewModelScanStatusReq("BulkScanId_example", []string{"ScanIds_example"}) // ModelScanStatusReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecretScanApi.StatusSecretScan(context.Background()).ModelScanStatusReq(modelScanStatusReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretScanApi.StatusSecretScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StatusSecretScan`: ModelScanStatusResp
    fmt.Fprintf(os.Stdout, "Response from `SecretScanApi.StatusSecretScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStatusSecretScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelScanStatusReq** | [**ModelScanStatusReq**](ModelScanStatusReq.md) |  | 

### Return type

[**ModelScanStatusResp**](ModelScanStatusResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopSecretScan

> StopSecretScan(ctx).ModelSecretScanTriggerReq(modelSecretScanTriggerReq).Execute()

Stop Secret Scan



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
    modelSecretScanTriggerReq := *openapiclient.NewModelSecretScanTriggerReq(*openapiclient.NewModelScanFilter(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}})), []openapiclient.ModelNodeIdentifier{*openapiclient.NewModelNodeIdentifier("NodeId_example", "NodeType_example")}) // ModelSecretScanTriggerReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SecretScanApi.StopSecretScan(context.Background()).ModelSecretScanTriggerReq(modelSecretScanTriggerReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecretScanApi.StopSecretScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStopSecretScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelSecretScanTriggerReq** | [**ModelSecretScanTriggerReq**](ModelSecretScanTriggerReq.md) |  | 

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

