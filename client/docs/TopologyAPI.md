# \TopologyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContainersTopologyGraph**](TopologyAPI.md#GetContainersTopologyGraph) | **Post** /deepfence/graph/topology/containers | Get Containers Topology Graph
[**GetHostsTopologyGraph**](TopologyAPI.md#GetHostsTopologyGraph) | **Post** /deepfence/graph/topology/hosts | Get Hosts Topology Graph
[**GetKubernetesTopologyGraph**](TopologyAPI.md#GetKubernetesTopologyGraph) | **Post** /deepfence/graph/topology/kubernetes | Get Kubernetes Topology Graph
[**GetPodsTopologyGraph**](TopologyAPI.md#GetPodsTopologyGraph) | **Post** /deepfence/graph/topology/pods | Get Pods Topology Graph
[**GetTopologyDelta**](TopologyAPI.md#GetTopologyDelta) | **Post** /deepfence/graph/topology/delta | Get Topology Delta
[**GetTopologyGraph**](TopologyAPI.md#GetTopologyGraph) | **Post** /deepfence/graph/topology/ | Get Topology Graph
[**IngestAgentReport**](TopologyAPI.md#IngestAgentReport) | **Post** /deepfence/ingest/report | Ingest Topology Data
[**IngestSyncAgentReport**](TopologyAPI.md#IngestSyncAgentReport) | **Post** /deepfence/ingest/sync-report | Ingest Topology Data



## GetContainersTopologyGraph

> ModelGraphResult GetContainersTopologyGraph(ctx).GraphTopologyFilters(graphTopologyFilters).Execute()

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
    graphTopologyFilters := *openapiclient.NewGraphTopologyFilters([]string{"CloudFilter_example"}, []string{"ContainerFilter_example"}, *openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"HostFilter_example"}, []string{"KubernetesFilter_example"}, []string{"PodFilter_example"}, []string{"RegionFilter_example"}, false) // GraphTopologyFilters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopologyAPI.GetContainersTopologyGraph(context.Background()).GraphTopologyFilters(graphTopologyFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.GetContainersTopologyGraph``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainersTopologyGraph`: ModelGraphResult
    fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.GetContainersTopologyGraph`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContainersTopologyGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphTopologyFilters** | [**GraphTopologyFilters**](GraphTopologyFilters.md) |  | 

### Return type

[**ModelGraphResult**](ModelGraphResult.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostsTopologyGraph

> ModelGraphResult GetHostsTopologyGraph(ctx).GraphTopologyFilters(graphTopologyFilters).Execute()

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
    graphTopologyFilters := *openapiclient.NewGraphTopologyFilters([]string{"CloudFilter_example"}, []string{"ContainerFilter_example"}, *openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"HostFilter_example"}, []string{"KubernetesFilter_example"}, []string{"PodFilter_example"}, []string{"RegionFilter_example"}, false) // GraphTopologyFilters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopologyAPI.GetHostsTopologyGraph(context.Background()).GraphTopologyFilters(graphTopologyFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.GetHostsTopologyGraph``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostsTopologyGraph`: ModelGraphResult
    fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.GetHostsTopologyGraph`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHostsTopologyGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphTopologyFilters** | [**GraphTopologyFilters**](GraphTopologyFilters.md) |  | 

### Return type

[**ModelGraphResult**](ModelGraphResult.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesTopologyGraph

> ModelGraphResult GetKubernetesTopologyGraph(ctx).GraphTopologyFilters(graphTopologyFilters).Execute()

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
    graphTopologyFilters := *openapiclient.NewGraphTopologyFilters([]string{"CloudFilter_example"}, []string{"ContainerFilter_example"}, *openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"HostFilter_example"}, []string{"KubernetesFilter_example"}, []string{"PodFilter_example"}, []string{"RegionFilter_example"}, false) // GraphTopologyFilters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopologyAPI.GetKubernetesTopologyGraph(context.Background()).GraphTopologyFilters(graphTopologyFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.GetKubernetesTopologyGraph``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesTopologyGraph`: ModelGraphResult
    fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.GetKubernetesTopologyGraph`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesTopologyGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphTopologyFilters** | [**GraphTopologyFilters**](GraphTopologyFilters.md) |  | 

### Return type

[**ModelGraphResult**](ModelGraphResult.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPodsTopologyGraph

> ModelGraphResult GetPodsTopologyGraph(ctx).GraphTopologyFilters(graphTopologyFilters).Execute()

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
    graphTopologyFilters := *openapiclient.NewGraphTopologyFilters([]string{"CloudFilter_example"}, []string{"ContainerFilter_example"}, *openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"HostFilter_example"}, []string{"KubernetesFilter_example"}, []string{"PodFilter_example"}, []string{"RegionFilter_example"}, false) // GraphTopologyFilters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopologyAPI.GetPodsTopologyGraph(context.Background()).GraphTopologyFilters(graphTopologyFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.GetPodsTopologyGraph``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPodsTopologyGraph`: ModelGraphResult
    fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.GetPodsTopologyGraph`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPodsTopologyGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphTopologyFilters** | [**GraphTopologyFilters**](GraphTopologyFilters.md) |  | 

### Return type

[**ModelGraphResult**](ModelGraphResult.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopologyDelta

> ModelTopologyDeltaResponse GetTopologyDelta(ctx).ModelTopologyDeltaReq(modelTopologyDeltaReq).Execute()

Get Topology Delta



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
    modelTopologyDeltaReq := *openapiclient.NewModelTopologyDeltaReq(false, int64(123), false, int64(123), []string{"EntityTypes_example"}) // ModelTopologyDeltaReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopologyAPI.GetTopologyDelta(context.Background()).ModelTopologyDeltaReq(modelTopologyDeltaReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.GetTopologyDelta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopologyDelta`: ModelTopologyDeltaResponse
    fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.GetTopologyDelta`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTopologyDeltaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelTopologyDeltaReq** | [**ModelTopologyDeltaReq**](ModelTopologyDeltaReq.md) |  | 

### Return type

[**ModelTopologyDeltaResponse**](ModelTopologyDeltaResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopologyGraph

> ModelGraphResult GetTopologyGraph(ctx).GraphTopologyFilters(graphTopologyFilters).Execute()

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
    graphTopologyFilters := *openapiclient.NewGraphTopologyFilters([]string{"CloudFilter_example"}, []string{"ContainerFilter_example"}, *openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"HostFilter_example"}, []string{"KubernetesFilter_example"}, []string{"PodFilter_example"}, []string{"RegionFilter_example"}, false) // GraphTopologyFilters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopologyAPI.GetTopologyGraph(context.Background()).GraphTopologyFilters(graphTopologyFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.GetTopologyGraph``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopologyGraph`: ModelGraphResult
    fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.GetTopologyGraph`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTopologyGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphTopologyFilters** | [**GraphTopologyFilters**](GraphTopologyFilters.md) |  | 

### Return type

[**ModelGraphResult**](ModelGraphResult.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IngestAgentReport

> ControlsAgentBeat IngestAgentReport(ctx).ReportRawReport(reportRawReport).Execute()

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
    resp, r, err := apiClient.TopologyAPI.IngestAgentReport(context.Background()).ReportRawReport(reportRawReport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.IngestAgentReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IngestAgentReport`: ControlsAgentBeat
    fmt.Fprintf(os.Stdout, "Response from `TopologyAPI.IngestAgentReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestAgentReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportRawReport** | [**ReportRawReport**](ReportRawReport.md) |  | 

### Return type

[**ControlsAgentBeat**](ControlsAgentBeat.md)

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
    ingestersReportIngestionData := *openapiclient.NewIngestersReportIngestionData([]map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, int32(123), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // IngestersReportIngestionData |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TopologyAPI.IngestSyncAgentReport(context.Background()).IngestersReportIngestionData(ingestersReportIngestionData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopologyAPI.IngestSyncAgentReport``: %v\n", err)
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

