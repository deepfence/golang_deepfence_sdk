# \ThreatAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIndividualThreatGraph**](ThreatAPI.md#GetIndividualThreatGraph) | **Post** /deepfence/graph/threat/individual | Get Vulnerability Threat Graph
[**GetThreatGraph**](ThreatAPI.md#GetThreatGraph) | **Post** /deepfence/graph/threat | Get Threat Graph



## GetIndividualThreatGraph

> []GraphIndividualThreatGraph GetIndividualThreatGraph(ctx).GraphIndividualThreatGraphRequest(graphIndividualThreatGraphRequest).Execute()

Get Vulnerability Threat Graph



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
	graphIndividualThreatGraphRequest := *openapiclient.NewGraphIndividualThreatGraphRequest("GraphType_example", "IssueType_example") // GraphIndividualThreatGraphRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreatAPI.GetIndividualThreatGraph(context.Background()).GraphIndividualThreatGraphRequest(graphIndividualThreatGraphRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatAPI.GetIndividualThreatGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIndividualThreatGraph`: []GraphIndividualThreatGraph
	fmt.Fprintf(os.Stdout, "Response from `ThreatAPI.GetIndividualThreatGraph`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIndividualThreatGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphIndividualThreatGraphRequest** | [**GraphIndividualThreatGraphRequest**](GraphIndividualThreatGraphRequest.md) |  | 

### Return type

[**[]GraphIndividualThreatGraph**](GraphIndividualThreatGraph.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.ThreatAPI.GetThreatGraph(context.Background()).GraphThreatFilters(graphThreatFilters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatAPI.GetThreatGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThreatGraph`: map[string]GraphProviderThreatGraph
	fmt.Fprintf(os.Stdout, "Response from `ThreatAPI.GetThreatGraph`: %v\n", resp)
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

