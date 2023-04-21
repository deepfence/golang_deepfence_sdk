# \SearchApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountCloudComplianceScans**](SearchApi.md#CountCloudComplianceScans) | **Post** /deepfence/search/count/cloud-compliance/scans | Count Cloud Compliance Scan results
[**CountCloudCompliances**](SearchApi.md#CountCloudCompliances) | **Post** /deepfence/search/count/cloud-compliances | Count Cloud compliances
[**CountCloudResources**](SearchApi.md#CountCloudResources) | **Post** /deepfence/search/count/cloud-resources | Count Cloud resources
[**CountComplianceScans**](SearchApi.md#CountComplianceScans) | **Post** /deepfence/search/count/compliance/scans | Count Compliance Scan results
[**CountCompliances**](SearchApi.md#CountCompliances) | **Post** /deepfence/search/count/compliances | Count Compliances
[**CountContainerImages**](SearchApi.md#CountContainerImages) | **Post** /deepfence/search/count/images | Count Container images
[**CountContainers**](SearchApi.md#CountContainers) | **Post** /deepfence/search/count/containers | Count Containers data
[**CountHosts**](SearchApi.md#CountHosts) | **Post** /deepfence/search/count/hosts | Count hosts
[**CountKubernetesClusters**](SearchApi.md#CountKubernetesClusters) | **Post** /deepfence/search/count/kubernetes-clusters | Count Kubernetes clusters
[**CountMalwareScans**](SearchApi.md#CountMalwareScans) | **Post** /deepfence/search/count/malware/scans | Count Malware Scan results
[**CountMalwares**](SearchApi.md#CountMalwares) | **Post** /deepfence/search/count/malwares | Count Malwares
[**CountNodes**](SearchApi.md#CountNodes) | **Get** /deepfence/search/count/nodes | Count nodes
[**CountPods**](SearchApi.md#CountPods) | **Post** /deepfence/search/count/pods | Count Pods
[**CountSecrets**](SearchApi.md#CountSecrets) | **Post** /deepfence/search/count/secrets | Count Secrets
[**CountSecretsScans**](SearchApi.md#CountSecretsScans) | **Post** /deepfence/search/count/secret/scans | Count Secret Scan results
[**CountVulnerabilities**](SearchApi.md#CountVulnerabilities) | **Post** /deepfence/search/count/vulnerabilities | Count Vulnerabilities
[**CountVulnerabilityScans**](SearchApi.md#CountVulnerabilityScans) | **Post** /deepfence/search/count/vulnerability/scans | Count Vulnerability Scan results
[**GetCloudComplianceFilters**](SearchApi.md#GetCloudComplianceFilters) | **Post** /deepfence/filters/cloud-compliance | Get Cloud Compliance Filters
[**GetComplianceFilters**](SearchApi.md#GetComplianceFilters) | **Post** /deepfence/filters/compliance | Get Compliance Filters
[**SearchCloudComplianceScans**](SearchApi.md#SearchCloudComplianceScans) | **Post** /deepfence/search/cloud-compliance/scans | Search Cloud Compliance Scan results
[**SearchCloudCompliances**](SearchApi.md#SearchCloudCompliances) | **Post** /deepfence/search/cloud-compliances | Search Cloud compliances
[**SearchCloudResources**](SearchApi.md#SearchCloudResources) | **Post** /deepfence/search/cloud-resources | Search Cloud Resources
[**SearchComplianceScans**](SearchApi.md#SearchComplianceScans) | **Post** /deepfence/search/compliance/scans | Search Compliance Scan results
[**SearchCompliances**](SearchApi.md#SearchCompliances) | **Post** /deepfence/search/compliances | Search Compliances
[**SearchContainerImages**](SearchApi.md#SearchContainerImages) | **Post** /deepfence/search/images | Search Container images
[**SearchContainers**](SearchApi.md#SearchContainers) | **Post** /deepfence/search/containers | Search Containers data
[**SearchHosts**](SearchApi.md#SearchHosts) | **Post** /deepfence/search/hosts | Search hosts
[**SearchKubernetesClusters**](SearchApi.md#SearchKubernetesClusters) | **Post** /deepfence/search/kubernetes-clusters | Search Kuberenetes Clusters
[**SearchMalwareScans**](SearchApi.md#SearchMalwareScans) | **Post** /deepfence/search/malware/scans | Search Malware Scan results
[**SearchMalwares**](SearchApi.md#SearchMalwares) | **Post** /deepfence/search/malwares | Search Malwares
[**SearchPods**](SearchApi.md#SearchPods) | **Post** /deepfence/search/pods | Search Pods
[**SearchSecrets**](SearchApi.md#SearchSecrets) | **Post** /deepfence/search/secrets | Search Secrets
[**SearchSecretsScans**](SearchApi.md#SearchSecretsScans) | **Post** /deepfence/search/secret/scans | Search Secrets Scan results
[**SearchVulnerabilities**](SearchApi.md#SearchVulnerabilities) | **Post** /deepfence/search/vulnerabilities | Search Vulnerabilities
[**SearchVulnerabilityScans**](SearchApi.md#SearchVulnerabilityScans) | **Post** /deepfence/search/vulnerability/scans | Search Vulnerability Scan results



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
    resp, r, err := apiClient.SearchApi.CountCloudComplianceScans(context.Background()).SearchSearchScanReq(searchSearchScanReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.CountCloudComplianceScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountCloudComplianceScans`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.CountCloudComplianceScans`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.CountCloudCompliances(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.CountCloudCompliances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountCloudCompliances`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.CountCloudCompliances`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.CountCloudResources(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.CountCloudResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountCloudResources`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.CountCloudResources`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.CountComplianceScans(context.Background()).SearchSearchScanReq(searchSearchScanReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.CountComplianceScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountComplianceScans`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.CountComplianceScans`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.CountCompliances(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.CountCompliances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountCompliances`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.CountCompliances`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.CountContainerImages(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.CountContainerImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountContainerImages`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.CountContainerImages`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.CountContainers(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.CountContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountContainers`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.CountContainers`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.CountHosts(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.CountHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountHosts`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.CountHosts`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.CountKubernetesClusters(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.CountKubernetesClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountKubernetesClusters`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.CountKubernetesClusters`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.CountMalwareScans(context.Background()).SearchSearchScanReq(searchSearchScanReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.CountMalwareScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountMalwareScans`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.CountMalwareScans`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.CountMalwares(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.CountMalwares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountMalwares`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.CountMalwares`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.CountNodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.CountNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountNodes`: SearchNodeCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.CountNodes`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.CountPods(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.CountPods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountPods`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.CountPods`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.CountSecrets(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.CountSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountSecrets`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.CountSecrets`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.CountSecretsScans(context.Background()).SearchSearchScanReq(searchSearchScanReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.CountSecretsScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountSecretsScans`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.CountSecretsScans`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.CountVulnerabilities(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.CountVulnerabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountVulnerabilities`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.CountVulnerabilities`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.CountVulnerabilityScans(context.Background()).SearchSearchScanReq(searchSearchScanReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.CountVulnerabilityScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountVulnerabilityScans`: SearchSearchCountResp
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.CountVulnerabilityScans`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.GetCloudComplianceFilters(context.Background()).ModelFiltersReq(modelFiltersReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetCloudComplianceFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudComplianceFilters`: ModelFiltersResult
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetCloudComplianceFilters`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.GetComplianceFilters(context.Background()).ModelFiltersReq(modelFiltersReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetComplianceFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComplianceFilters`: ModelFiltersResult
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetComplianceFilters`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.SearchCloudComplianceScans(context.Background()).SearchSearchScanReq(searchSearchScanReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchCloudComplianceScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCloudComplianceScans`: []ModelScanInfo
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchCloudComplianceScans`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.SearchCloudCompliances(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchCloudCompliances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCloudCompliances`: []ModelCloudCompliance
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchCloudCompliances`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.SearchCloudResources(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchCloudResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCloudResources`: []ModelCloudResource
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchCloudResources`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.SearchComplianceScans(context.Background()).SearchSearchScanReq(searchSearchScanReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchComplianceScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchComplianceScans`: []ModelScanInfo
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchComplianceScans`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.SearchCompliances(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchCompliances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCompliances`: []ModelCompliance
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchCompliances`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.SearchContainerImages(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchContainerImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchContainerImages`: []ModelContainerImage
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchContainerImages`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.SearchContainers(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchContainers`: []ModelContainer
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchContainers`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.SearchHosts(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchHosts`: []ModelHost
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchHosts`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.SearchKubernetesClusters(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchKubernetesClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchKubernetesClusters`: []ModelKubernetesCluster
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchKubernetesClusters`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.SearchMalwareScans(context.Background()).SearchSearchScanReq(searchSearchScanReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchMalwareScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchMalwareScans`: []ModelScanInfo
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchMalwareScans`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.SearchMalwares(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchMalwares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchMalwares`: []ModelMalware
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchMalwares`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.SearchPods(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchPods``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchPods`: []ModelPod
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchPods`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.SearchSecrets(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSecrets`: []ModelSecret
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchSecrets`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.SearchSecretsScans(context.Background()).SearchSearchScanReq(searchSearchScanReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchSecretsScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSecretsScans`: []ModelScanInfo
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchSecretsScans`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.SearchVulnerabilities(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchVulnerabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchVulnerabilities`: []ModelVulnerability
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchVulnerabilities`: %v\n", resp)
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
    resp, r, err := apiClient.SearchApi.SearchVulnerabilityScans(context.Background()).SearchSearchScanReq(searchSearchScanReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchVulnerabilityScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchVulnerabilityScans`: []ModelScanInfo
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchVulnerabilityScans`: %v\n", resp)
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

