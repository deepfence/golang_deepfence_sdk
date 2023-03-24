# \ThreatApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetThreatGraph**](ThreatApi.md#GetThreatGraph) | **Post** /deepfence/graph/threat | Get Threat Graph



## GetThreatGraph

> map[string]GraphProviderThreatGraph GetThreatGraph(ctx).GraphThreatFilters(graphThreatFilters).Execute()

Get Threat Graph



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
    graphThreatFilters := *openapiclient.NewGraphThreatFilters(*openapiclient.NewGraphCloudProviderFilter([]string{"AccountIds_example"}), *openapiclient.NewGraphCloudProviderFilter([]string{"AccountIds_example"}), false, *openapiclient.NewGraphCloudProviderFilter([]string{"AccountIds_example"}), "Type_example") // GraphThreatFilters |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThreatApi.GetThreatGraph(context.Background()).GraphThreatFilters(graphThreatFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThreatApi.GetThreatGraph``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetThreatGraph`: map[string]GraphProviderThreatGraph
    fmt.Fprintf(os.Stdout, "Response from `ThreatApi.GetThreatGraph`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetThreatGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphThreatFilters** | [**GraphThreatFilters**](GraphThreatFilters.md) |  | 

### Return type

[**map[string]GraphProviderThreatGraph**](GraphProviderThreatGraph.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

