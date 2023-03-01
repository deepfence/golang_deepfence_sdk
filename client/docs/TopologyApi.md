# \TopologyApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContainersTopologyGraph**](TopologyApi.md#GetContainersTopologyGraph) | **Post** /deepfence/graph/topology/containers | Get Containers Topology Graph
[**GetHostsTopologyGraph**](TopologyApi.md#GetHostsTopologyGraph) | **Post** /deepfence/graph/topology/hosts | Get Hosts Topology Graph
[**GetKubernetesTopologyGraph**](TopologyApi.md#GetKubernetesTopologyGraph) | **Post** /deepfence/graph/topology/kubernetes | Get Kubernetes Topology Graph
[**GetPodsTopologyGraph**](TopologyApi.md#GetPodsTopologyGraph) | **Post** /deepfence/graph/topology/pods | Get Pods Topology Graph
[**GetTopologyGraph**](TopologyApi.md#GetTopologyGraph) | **Post** /deepfence/graph/topology/ | Get Topology Graph
[**IngestAgentReport**](TopologyApi.md#IngestAgentReport) | **Post** /deepfence/ingest/report | Ingest Topology Data
[**IngestSyncAgentReport**](TopologyApi.md#IngestSyncAgentReport) | **Post** /deepfence/ingest/sync-report | Ingest Topology Data



## GetContainersTopologyGraph

> ApiDocsGraphResult GetContainersTopologyGraph(ctx).GraphTopologyFilters(graphTopologyFilters).Execute()

Get Containers Topology Graph



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
    graphTopologyFilters := *openapiclient.NewGraphTopologyFilters([]string{"CloudFilter_example"}, *openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"HostFilter_example"}, []string{"KubernetesFilter_example"}, []string{"PodFilter_example"}, []string{"RegionFilter_example"}) // GraphTopologyFilters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopologyApi.GetContainersTopologyGraph(context.Background()).GraphTopologyFilters(graphTopologyFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologyApi.GetContainersTopologyGraph``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainersTopologyGraph`: ApiDocsGraphResult
    fmt.Fprintf(os.Stdout, "Response from `TopologyApi.GetContainersTopologyGraph`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContainersTopologyGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphTopologyFilters** | [**GraphTopologyFilters**](GraphTopologyFilters.md) |  | 

### Return type

[**ApiDocsGraphResult**](ApiDocsGraphResult.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostsTopologyGraph

> ApiDocsGraphResult GetHostsTopologyGraph(ctx).GraphTopologyFilters(graphTopologyFilters).Execute()

Get Hosts Topology Graph



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
    graphTopologyFilters := *openapiclient.NewGraphTopologyFilters([]string{"CloudFilter_example"}, *openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"HostFilter_example"}, []string{"KubernetesFilter_example"}, []string{"PodFilter_example"}, []string{"RegionFilter_example"}) // GraphTopologyFilters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopologyApi.GetHostsTopologyGraph(context.Background()).GraphTopologyFilters(graphTopologyFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologyApi.GetHostsTopologyGraph``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostsTopologyGraph`: ApiDocsGraphResult
    fmt.Fprintf(os.Stdout, "Response from `TopologyApi.GetHostsTopologyGraph`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHostsTopologyGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphTopologyFilters** | [**GraphTopologyFilters**](GraphTopologyFilters.md) |  | 

### Return type

[**ApiDocsGraphResult**](ApiDocsGraphResult.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesTopologyGraph

> ApiDocsGraphResult GetKubernetesTopologyGraph(ctx).GraphTopologyFilters(graphTopologyFilters).Execute()

Get Kubernetes Topology Graph



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
    graphTopologyFilters := *openapiclient.NewGraphTopologyFilters([]string{"CloudFilter_example"}, *openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"HostFilter_example"}, []string{"KubernetesFilter_example"}, []string{"PodFilter_example"}, []string{"RegionFilter_example"}) // GraphTopologyFilters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopologyApi.GetKubernetesTopologyGraph(context.Background()).GraphTopologyFilters(graphTopologyFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologyApi.GetKubernetesTopologyGraph``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesTopologyGraph`: ApiDocsGraphResult
    fmt.Fprintf(os.Stdout, "Response from `TopologyApi.GetKubernetesTopologyGraph`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesTopologyGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphTopologyFilters** | [**GraphTopologyFilters**](GraphTopologyFilters.md) |  | 

### Return type

[**ApiDocsGraphResult**](ApiDocsGraphResult.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPodsTopologyGraph

> ApiDocsGraphResult GetPodsTopologyGraph(ctx).GraphTopologyFilters(graphTopologyFilters).Execute()

Get Pods Topology Graph



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
    graphTopologyFilters := *openapiclient.NewGraphTopologyFilters([]string{"CloudFilter_example"}, *openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"HostFilter_example"}, []string{"KubernetesFilter_example"}, []string{"PodFilter_example"}, []string{"RegionFilter_example"}) // GraphTopologyFilters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopologyApi.GetPodsTopologyGraph(context.Background()).GraphTopologyFilters(graphTopologyFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologyApi.GetPodsTopologyGraph``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPodsTopologyGraph`: ApiDocsGraphResult
    fmt.Fprintf(os.Stdout, "Response from `TopologyApi.GetPodsTopologyGraph`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPodsTopologyGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphTopologyFilters** | [**GraphTopologyFilters**](GraphTopologyFilters.md) |  | 

### Return type

[**ApiDocsGraphResult**](ApiDocsGraphResult.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopologyGraph

> ApiDocsGraphResult GetTopologyGraph(ctx).GraphTopologyFilters(graphTopologyFilters).Execute()

Get Topology Graph



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
    graphTopologyFilters := *openapiclient.NewGraphTopologyFilters([]string{"CloudFilter_example"}, *openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"HostFilter_example"}, []string{"KubernetesFilter_example"}, []string{"PodFilter_example"}, []string{"RegionFilter_example"}) // GraphTopologyFilters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopologyApi.GetTopologyGraph(context.Background()).GraphTopologyFilters(graphTopologyFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologyApi.GetTopologyGraph``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopologyGraph`: ApiDocsGraphResult
    fmt.Fprintf(os.Stdout, "Response from `TopologyApi.GetTopologyGraph`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTopologyGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphTopologyFilters** | [**GraphTopologyFilters**](GraphTopologyFilters.md) |  | 

### Return type

[**ApiDocsGraphResult**](ApiDocsGraphResult.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IngestAgentReport

> IngestAgentReport(ctx).ReportRawReport(reportRawReport).Execute()

Ingest Topology Data



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
    reportRawReport := *openapiclient.NewReportRawReport("Payload_example") // ReportRawReport |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TopologyApi.IngestAgentReport(context.Background()).ReportRawReport(reportRawReport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologyApi.IngestAgentReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestAgentReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportRawReport** | [**ReportRawReport**](ReportRawReport.md) |  | 

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


## IngestSyncAgentReport

> IngestSyncAgentReport(ctx).IngestersReportIngestionData(ingestersReportIngestionData).Execute()

Ingest Topology Data



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
    ingestersReportIngestionData := *openapiclient.NewIngestersReportIngestionData([]map[string]string{map[string]string{"key": "Inner_example"}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]string{map[string]string{"key": "Inner_example"}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]string{map[string]string{"key": "Inner_example"}}, []map[string]string{map[string]string{"key": "Inner_example"}}, []map[string]string{map[string]string{"key": "Inner_example"}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]string{map[string]string{"key": "Inner_example"}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]string{map[string]string{"key": "Inner_example"}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // IngestersReportIngestionData |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TopologyApi.IngestSyncAgentReport(context.Background()).IngestersReportIngestionData(ingestersReportIngestionData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologyApi.IngestSyncAgentReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestSyncAgentReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestersReportIngestionData** | [**IngestersReportIngestionData**](IngestersReportIngestionData.md) |  | 

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

