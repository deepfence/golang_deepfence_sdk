# \SearchAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountCloudComplianceScans**](SearchAPI.md#CountCloudComplianceScans) | **Post** /deepfence/search/count/cloud-compliance/scans | Count Cloud Compliance Scan results
[**CountCloudCompliances**](SearchAPI.md#CountCloudCompliances) | **Post** /deepfence/search/count/cloud-compliances | Count Cloud compliances
[**CountCloudResources**](SearchAPI.md#CountCloudResources) | **Post** /deepfence/search/count/cloud-resources | Count Cloud resources
[**CountComplianceScans**](SearchAPI.md#CountComplianceScans) | **Post** /deepfence/search/count/compliance/scans | Count Compliance Scan results
[**CountCompliances**](SearchAPI.md#CountCompliances) | **Post** /deepfence/search/count/compliances | Count Compliances
[**CountContainerImages**](SearchAPI.md#CountContainerImages) | **Post** /deepfence/search/count/images | Count Container images
[**CountContainers**](SearchAPI.md#CountContainers) | **Post** /deepfence/search/count/containers | Count Containers data
[**CountHosts**](SearchAPI.md#CountHosts) | **Post** /deepfence/search/count/hosts | Count hosts
[**CountKubernetesClusters**](SearchAPI.md#CountKubernetesClusters) | **Post** /deepfence/search/count/kubernetes-clusters | Count Kubernetes clusters
[**CountMalwareScans**](SearchAPI.md#CountMalwareScans) | **Post** /deepfence/search/count/malware/scans | Count Malware Scan results
[**CountMalwares**](SearchAPI.md#CountMalwares) | **Post** /deepfence/search/count/malwares | Count Malwares
[**CountNodes**](SearchAPI.md#CountNodes) | **Get** /deepfence/search/count/nodes | Count nodes
[**CountPods**](SearchAPI.md#CountPods) | **Post** /deepfence/search/count/pods | Count Pods
[**CountSecrets**](SearchAPI.md#CountSecrets) | **Post** /deepfence/search/count/secrets | Count Secrets
[**CountSecretsScans**](SearchAPI.md#CountSecretsScans) | **Post** /deepfence/search/count/secret/scans | Count Secret Scan results
[**CountVulnerabilities**](SearchAPI.md#CountVulnerabilities) | **Post** /deepfence/search/count/vulnerabilities | Count Vulnerabilities
[**CountVulnerabilityScans**](SearchAPI.md#CountVulnerabilityScans) | **Post** /deepfence/search/count/vulnerability/scans | Count Vulnerability Scan results
[**GetCloudComplianceFilters**](SearchAPI.md#GetCloudComplianceFilters) | **Post** /deepfence/filters/cloud-compliance | Get Cloud Compliance Filters
[**GetComplianceFilters**](SearchAPI.md#GetComplianceFilters) | **Post** /deepfence/filters/compliance | Get Compliance Filters
[**SearchCloudComplianceScans**](SearchAPI.md#SearchCloudComplianceScans) | **Post** /deepfence/search/cloud-compliance/scans | Search Cloud Compliance Scan results
[**SearchCloudCompliances**](SearchAPI.md#SearchCloudCompliances) | **Post** /deepfence/search/cloud-compliances | Search Cloud compliances
[**SearchCloudNodes**](SearchAPI.md#SearchCloudNodes) | **Post** /deepfence/search/cloud-nodes | Search Cloud Nodes
[**SearchCloudResources**](SearchAPI.md#SearchCloudResources) | **Post** /deepfence/search/cloud-resources | Search Cloud Resources
[**SearchComplianceScans**](SearchAPI.md#SearchComplianceScans) | **Post** /deepfence/search/compliance/scans | Search Compliance Scan results
[**SearchCompliances**](SearchAPI.md#SearchCompliances) | **Post** /deepfence/search/compliances | Search Compliances
[**SearchContainerImages**](SearchAPI.md#SearchContainerImages) | **Post** /deepfence/search/images | Search Container images
[**SearchContainers**](SearchAPI.md#SearchContainers) | **Post** /deepfence/search/containers | Search Containers data
[**SearchHosts**](SearchAPI.md#SearchHosts) | **Post** /deepfence/search/hosts | Search hosts
[**SearchKubernetesClusters**](SearchAPI.md#SearchKubernetesClusters) | **Post** /deepfence/search/kubernetes-clusters | Search Kuberenetes Clusters
[**SearchMalwareScans**](SearchAPI.md#SearchMalwareScans) | **Post** /deepfence/search/malware/scans | Search Malware Scan results
[**SearchMalwares**](SearchAPI.md#SearchMalwares) | **Post** /deepfence/search/malwares | Search Malwares
[**SearchPods**](SearchAPI.md#SearchPods) | **Post** /deepfence/search/pods | Search Pods
[**SearchSecrets**](SearchAPI.md#SearchSecrets) | **Post** /deepfence/search/secrets | Search Secrets
[**SearchSecretsScans**](SearchAPI.md#SearchSecretsScans) | **Post** /deepfence/search/secret/scans | Search Secrets Scan results
[**SearchVulnerabilities**](SearchAPI.md#SearchVulnerabilities) | **Post** /deepfence/search/vulnerabilities | Search Vulnerabilities
[**SearchVulnerabilityScans**](SearchAPI.md#SearchVulnerabilityScans) | **Post** /deepfence/search/vulnerability/scans | Search Vulnerability Scan results



## CountCloudComplianceScans

> SearchSearchCountResp CountCloudComplianceScans(ctx).SearchSearchScanReq(searchSearchScanReq).Execute()

Count Cloud Compliance Scan results



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
    searchSearchScanReq := *openapiclient.NewSearchSearchScanReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchScanReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.CountCloudComplianceScans(context.Background()).SearchSearchScanReq(searchSearchScanReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CountCloudComplianceScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountCloudComplianceScans`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CountCloudComplianceScans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountCloudComplianceScansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchScanReq** | [**SearchSearchScanReq**](SearchSearchScanReq.md) |  | 

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


## CountCloudCompliances

> SearchSearchCountResp CountCloudCompliances(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

Count Cloud compliances



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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.CountCloudCompliances(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CountCloudCompliances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountCloudCompliances`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CountCloudCompliances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountCloudCompliancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchNodeReq** | [**SearchSearchNodeReq**](SearchSearchNodeReq.md) |  | 

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


## CountCloudResources

> SearchSearchCountResp CountCloudResources(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

Count Cloud resources



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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.CountCloudResources(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CountCloudResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountCloudResources`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CountCloudResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountCloudResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchNodeReq** | [**SearchSearchNodeReq**](SearchSearchNodeReq.md) |  | 

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


## CountComplianceScans

> SearchSearchCountResp CountComplianceScans(ctx).SearchSearchScanReq(searchSearchScanReq).Execute()

Count Compliance Scan results



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
    searchSearchScanReq := *openapiclient.NewSearchSearchScanReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchScanReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.CountComplianceScans(context.Background()).SearchSearchScanReq(searchSearchScanReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CountComplianceScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountComplianceScans`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CountComplianceScans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountComplianceScansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchScanReq** | [**SearchSearchScanReq**](SearchSearchScanReq.md) |  | 

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


## CountCompliances

> SearchSearchCountResp CountCompliances(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

Count Compliances



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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.CountCompliances(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CountCompliances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountCompliances`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CountCompliances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountCompliancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchNodeReq** | [**SearchSearchNodeReq**](SearchSearchNodeReq.md) |  | 

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


## CountContainerImages

> SearchSearchCountResp CountContainerImages(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

Count Container images



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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.CountContainerImages(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CountContainerImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountContainerImages`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CountContainerImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountContainerImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchNodeReq** | [**SearchSearchNodeReq**](SearchSearchNodeReq.md) |  | 

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


## CountContainers

> SearchSearchCountResp CountContainers(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

Count Containers data



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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.CountContainers(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CountContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountContainers`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CountContainers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchNodeReq** | [**SearchSearchNodeReq**](SearchSearchNodeReq.md) |  | 

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


## CountHosts

> SearchSearchCountResp CountHosts(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

Count hosts



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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.CountHosts(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CountHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountHosts`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CountHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchNodeReq** | [**SearchSearchNodeReq**](SearchSearchNodeReq.md) |  | 

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


## CountKubernetesClusters

> SearchSearchCountResp CountKubernetesClusters(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

Count Kubernetes clusters



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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.CountKubernetesClusters(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CountKubernetesClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountKubernetesClusters`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CountKubernetesClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountKubernetesClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchNodeReq** | [**SearchSearchNodeReq**](SearchSearchNodeReq.md) |  | 

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


## CountMalwareScans

> SearchSearchCountResp CountMalwareScans(ctx).SearchSearchScanReq(searchSearchScanReq).Execute()

Count Malware Scan results



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
    searchSearchScanReq := *openapiclient.NewSearchSearchScanReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchScanReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.CountMalwareScans(context.Background()).SearchSearchScanReq(searchSearchScanReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CountMalwareScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountMalwareScans`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CountMalwareScans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountMalwareScansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchScanReq** | [**SearchSearchScanReq**](SearchSearchScanReq.md) |  | 

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


## CountMalwares

> SearchSearchCountResp CountMalwares(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

Count Malwares



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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.CountMalwares(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CountMalwares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountMalwares`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CountMalwares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountMalwaresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchNodeReq** | [**SearchSearchNodeReq**](SearchSearchNodeReq.md) |  | 

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


## CountNodes

> SearchNodeCountResp CountNodes(ctx).Execute()

Count nodes



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
    resp, r, err := apiClient.SearchAPI.CountNodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CountNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountNodes`: SearchNodeCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CountNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCountNodesRequest struct via the builder pattern


### Return type

[**SearchNodeCountResp**](SearchNodeCountResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountPods

> SearchSearchCountResp CountPods(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

Count Pods



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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.CountPods(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CountPods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountPods`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CountPods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountPodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchNodeReq** | [**SearchSearchNodeReq**](SearchSearchNodeReq.md) |  | 

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


## CountSecrets

> SearchSearchCountResp CountSecrets(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

Count Secrets



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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.CountSecrets(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CountSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountSecrets`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CountSecrets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchNodeReq** | [**SearchSearchNodeReq**](SearchSearchNodeReq.md) |  | 

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


## CountSecretsScans

> SearchSearchCountResp CountSecretsScans(ctx).SearchSearchScanReq(searchSearchScanReq).Execute()

Count Secret Scan results



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
    searchSearchScanReq := *openapiclient.NewSearchSearchScanReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchScanReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.CountSecretsScans(context.Background()).SearchSearchScanReq(searchSearchScanReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CountSecretsScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountSecretsScans`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CountSecretsScans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountSecretsScansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchScanReq** | [**SearchSearchScanReq**](SearchSearchScanReq.md) |  | 

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


## CountVulnerabilities

> SearchSearchCountResp CountVulnerabilities(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

Count Vulnerabilities



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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.CountVulnerabilities(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CountVulnerabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountVulnerabilities`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CountVulnerabilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountVulnerabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchNodeReq** | [**SearchSearchNodeReq**](SearchSearchNodeReq.md) |  | 

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


## CountVulnerabilityScans

> SearchSearchCountResp CountVulnerabilityScans(ctx).SearchSearchScanReq(searchSearchScanReq).Execute()

Count Vulnerability Scan results



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
    searchSearchScanReq := *openapiclient.NewSearchSearchScanReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchScanReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.CountVulnerabilityScans(context.Background()).SearchSearchScanReq(searchSearchScanReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CountVulnerabilityScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountVulnerabilityScans`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.CountVulnerabilityScans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountVulnerabilityScansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchScanReq** | [**SearchSearchScanReq**](SearchSearchScanReq.md) |  | 

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


## GetCloudComplianceFilters

> ModelFiltersResult GetCloudComplianceFilters(ctx).ModelFiltersReq(modelFiltersReq).Execute()

Get Cloud Compliance Filters



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
    modelFiltersReq := *openapiclient.NewModelFiltersReq([]string{"Filters_example"}) // ModelFiltersReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.GetCloudComplianceFilters(context.Background()).ModelFiltersReq(modelFiltersReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.GetCloudComplianceFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudComplianceFilters`: ModelFiltersResult
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.GetCloudComplianceFilters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudComplianceFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelFiltersReq** | [**ModelFiltersReq**](ModelFiltersReq.md) |  | 

### Return type

[**ModelFiltersResult**](ModelFiltersResult.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComplianceFilters

> ModelFiltersResult GetComplianceFilters(ctx).ModelFiltersReq(modelFiltersReq).Execute()

Get Compliance Filters



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
    modelFiltersReq := *openapiclient.NewModelFiltersReq([]string{"Filters_example"}) // ModelFiltersReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.GetComplianceFilters(context.Background()).ModelFiltersReq(modelFiltersReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.GetComplianceFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComplianceFilters`: ModelFiltersResult
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.GetComplianceFilters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetComplianceFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelFiltersReq** | [**ModelFiltersReq**](ModelFiltersReq.md) |  | 

### Return type

[**ModelFiltersResult**](ModelFiltersResult.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCloudComplianceScans

> []ModelScanInfo SearchCloudComplianceScans(ctx).SearchSearchScanReq(searchSearchScanReq).Execute()

Search Cloud Compliance Scan results



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
    searchSearchScanReq := *openapiclient.NewSearchSearchScanReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchScanReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchCloudComplianceScans(context.Background()).SearchSearchScanReq(searchSearchScanReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchCloudComplianceScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCloudComplianceScans`: []ModelScanInfo
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchCloudComplianceScans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCloudComplianceScansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchScanReq** | [**SearchSearchScanReq**](SearchSearchScanReq.md) |  | 

### Return type

[**[]ModelScanInfo**](ModelScanInfo.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCloudCompliances

> []ModelCloudCompliance SearchCloudCompliances(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchCloudCompliances(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchCloudCompliances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCloudCompliances`: []ModelCloudCompliance
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchCloudCompliances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCloudCompliancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchNodeReq** | [**SearchSearchNodeReq**](SearchSearchNodeReq.md) |  | 

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


## SearchCloudNodes

> []ModelCloudNodeAccountInfo SearchCloudNodes(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

Search Cloud Nodes



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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchCloudNodes(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchCloudNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCloudNodes`: []ModelCloudNodeAccountInfo
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchCloudNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCloudNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchNodeReq** | [**SearchSearchNodeReq**](SearchSearchNodeReq.md) |  | 

### Return type

[**[]ModelCloudNodeAccountInfo**](ModelCloudNodeAccountInfo.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCloudResources

> []ModelCloudResource SearchCloudResources(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

Search Cloud Resources



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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchCloudResources(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchCloudResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCloudResources`: []ModelCloudResource
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchCloudResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCloudResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchNodeReq** | [**SearchSearchNodeReq**](SearchSearchNodeReq.md) |  | 

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


## SearchComplianceScans

> []ModelScanInfo SearchComplianceScans(ctx).SearchSearchScanReq(searchSearchScanReq).Execute()

Search Compliance Scan results



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
    searchSearchScanReq := *openapiclient.NewSearchSearchScanReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchScanReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchComplianceScans(context.Background()).SearchSearchScanReq(searchSearchScanReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchComplianceScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchComplianceScans`: []ModelScanInfo
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchComplianceScans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchComplianceScansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchScanReq** | [**SearchSearchScanReq**](SearchSearchScanReq.md) |  | 

### Return type

[**[]ModelScanInfo**](ModelScanInfo.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCompliances

> []ModelCompliance SearchCompliances(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchCompliances(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchCompliances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCompliances`: []ModelCompliance
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchCompliances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCompliancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchNodeReq** | [**SearchSearchNodeReq**](SearchSearchNodeReq.md) |  | 

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

> []ModelContainerImage SearchContainerImages(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchContainerImages(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchContainerImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchContainerImages`: []ModelContainerImage
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchContainerImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchContainerImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchNodeReq** | [**SearchSearchNodeReq**](SearchSearchNodeReq.md) |  | 

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

> []ModelContainer SearchContainers(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchContainers(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchContainers`: []ModelContainer
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchContainers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchNodeReq** | [**SearchSearchNodeReq**](SearchSearchNodeReq.md) |  | 

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

> []ModelHost SearchHosts(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchHosts(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchHosts`: []ModelHost
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchHosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchHostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchNodeReq** | [**SearchSearchNodeReq**](SearchSearchNodeReq.md) |  | 

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


## SearchKubernetesClusters

> []ModelKubernetesCluster SearchKubernetesClusters(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

Search Kuberenetes Clusters



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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchKubernetesClusters(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchKubernetesClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchKubernetesClusters`: []ModelKubernetesCluster
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchKubernetesClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchKubernetesClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchNodeReq** | [**SearchSearchNodeReq**](SearchSearchNodeReq.md) |  | 

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


## SearchMalwareScans

> []ModelScanInfo SearchMalwareScans(ctx).SearchSearchScanReq(searchSearchScanReq).Execute()

Search Malware Scan results



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
    searchSearchScanReq := *openapiclient.NewSearchSearchScanReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchScanReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchMalwareScans(context.Background()).SearchSearchScanReq(searchSearchScanReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchMalwareScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchMalwareScans`: []ModelScanInfo
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchMalwareScans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchMalwareScansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchScanReq** | [**SearchSearchScanReq**](SearchSearchScanReq.md) |  | 

### Return type

[**[]ModelScanInfo**](ModelScanInfo.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchMalwares

> []ModelMalware SearchMalwares(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchMalwares(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchMalwares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchMalwares`: []ModelMalware
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchMalwares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchMalwaresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchNodeReq** | [**SearchSearchNodeReq**](SearchSearchNodeReq.md) |  | 

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


## SearchPods

> []ModelPod SearchPods(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

Search Pods



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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchPods(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchPods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchPods`: []ModelPod
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchPods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchNodeReq** | [**SearchSearchNodeReq**](SearchSearchNodeReq.md) |  | 

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


## SearchSecrets

> []ModelSecret SearchSecrets(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchSecrets(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSecrets`: []ModelSecret
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchSecrets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchNodeReq** | [**SearchSearchNodeReq**](SearchSearchNodeReq.md) |  | 

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


## SearchSecretsScans

> []ModelScanInfo SearchSecretsScans(ctx).SearchSearchScanReq(searchSearchScanReq).Execute()

Search Secrets Scan results



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
    searchSearchScanReq := *openapiclient.NewSearchSearchScanReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchScanReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchSecretsScans(context.Background()).SearchSearchScanReq(searchSearchScanReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchSecretsScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSecretsScans`: []ModelScanInfo
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchSecretsScans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSecretsScansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchScanReq** | [**SearchSearchScanReq**](SearchSearchScanReq.md) |  | 

### Return type

[**[]ModelScanInfo**](ModelScanInfo.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchVulnerabilities

> []ModelVulnerability SearchVulnerabilities(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchVulnerabilities(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchVulnerabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchVulnerabilities`: []ModelVulnerability
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchVulnerabilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchVulnerabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchNodeReq** | [**SearchSearchNodeReq**](SearchSearchNodeReq.md) |  | 

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


## SearchVulnerabilityScans

> []ModelScanInfo SearchVulnerabilityScans(ctx).SearchSearchScanReq(searchSearchScanReq).Execute()

Search Vulnerability Scan results



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
    searchSearchScanReq := *openapiclient.NewSearchSearchScanReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters([]openapiclient.ReportersCompareFilter{*openapiclient.NewReportersCompareFilter("FieldName_example", interface{}(123), false)}, *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter([]openapiclient.ReportersOrderSpec{*openapiclient.NewReportersOrderSpec(false, "FieldName_example")})), []string{"InFieldFilter_example"}, *openapiclient.NewModelFetchWindow(int32(123), int32(123))), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchScanReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.SearchVulnerabilityScans(context.Background()).SearchSearchScanReq(searchSearchScanReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchVulnerabilityScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchVulnerabilityScans`: []ModelScanInfo
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchVulnerabilityScans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchVulnerabilityScansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchSearchScanReq** | [**SearchSearchScanReq**](SearchSearchScanReq.md) |  | 

### Return type

[**[]ModelScanInfo**](ModelScanInfo.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

