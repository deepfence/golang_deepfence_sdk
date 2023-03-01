# \SearchApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountCloudComplianceScans**](SearchApi.md#CountCloudComplianceScans) | **Post** /deepfence/search/count/cloud-compliance/scans | Count Cloud Compliance Scan results
[**CountCloudCompliances**](SearchApi.md#CountCloudCompliances) | **Post** /deepfence/search/count/cloud-compliances | Count Cloud compliances
[**CountComplianceScans**](SearchApi.md#CountComplianceScans) | **Post** /deepfence/search/count/compliance/scans | Count Compliance Scan results
[**CountCompliances**](SearchApi.md#CountCompliances) | **Post** /deepfence/search/count/compliances | Count Compliances
[**CountContainerImages**](SearchApi.md#CountContainerImages) | **Post** /deepfence/search/count/images | Count Container images
[**CountContainers**](SearchApi.md#CountContainers) | **Post** /deepfence/search/count/containers | Count Containers data
[**CountHosts**](SearchApi.md#CountHosts) | **Post** /deepfence/search/count/hosts | Count hosts
[**CountMalwareScans**](SearchApi.md#CountMalwareScans) | **Post** /deepfence/search/count/malware/scans | Count Malware Scan results
[**CountMalwares**](SearchApi.md#CountMalwares) | **Post** /deepfence/search/count/malwares | Count Malwares
[**CountSecrets**](SearchApi.md#CountSecrets) | **Post** /deepfence/search/count/secrets | Count Secrets
[**CountSecretsScans**](SearchApi.md#CountSecretsScans) | **Post** /deepfence/search/count/secret/scans | Count Secret Scan results
[**CountVulnerabilities**](SearchApi.md#CountVulnerabilities) | **Post** /deepfence/search/count/vulnerabilities | Count Vulnerabilities
[**CountVulnerabilityScans**](SearchApi.md#CountVulnerabilityScans) | **Post** /deepfence/search/count/vulnerability/scans | Count Vulnerability Scan results
[**SearchCloudComplianceScans**](SearchApi.md#SearchCloudComplianceScans) | **Post** /deepfence/search/cloud-compliance/scans | Search Cloud Compliance Scan results
[**SearchCloudCompliances**](SearchApi.md#SearchCloudCompliances) | **Post** /deepfence/search/cloud-compliances | Search Cloud compliances
[**SearchComplianceScans**](SearchApi.md#SearchComplianceScans) | **Post** /deepfence/search/compliance/scans | Search Compliance Scan results
[**SearchCompliances**](SearchApi.md#SearchCompliances) | **Post** /deepfence/search/compliances | Search Compliances
[**SearchContainerImages**](SearchApi.md#SearchContainerImages) | **Post** /deepfence/search/images | Search Container images
[**SearchContainers**](SearchApi.md#SearchContainers) | **Post** /deepfence/search/containers | Search Containers data
[**SearchHosts**](SearchApi.md#SearchHosts) | **Post** /deepfence/search/hosts | Search hosts
[**SearchMalwareScans**](SearchApi.md#SearchMalwareScans) | **Post** /deepfence/search/malware/scans | Search Malware Scan results
[**SearchMalwares**](SearchApi.md#SearchMalwares) | **Post** /deepfence/search/malwares | Search Malwares
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
    searchSearchScanReq := *openapiclient.NewSearchSearchScanReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchScanReq |  (optional)

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

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
    searchSearchScanReq := *openapiclient.NewSearchSearchScanReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchScanReq |  (optional)

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

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
    searchSearchScanReq := *openapiclient.NewSearchSearchScanReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchScanReq |  (optional)

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

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
    searchSearchScanReq := *openapiclient.NewSearchSearchScanReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchScanReq |  (optional)

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

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
    searchSearchScanReq := *openapiclient.NewSearchSearchScanReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchScanReq |  (optional)

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
    searchSearchScanReq := *openapiclient.NewSearchSearchScanReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchScanReq |  (optional)

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

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
    searchSearchScanReq := *openapiclient.NewSearchSearchScanReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchScanReq |  (optional)

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

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
    searchSearchScanReq := *openapiclient.NewSearchSearchScanReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchScanReq |  (optional)

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

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


## SearchSecrets

> []ModelSecretRule SearchSecrets(ctx).SearchSearchNodeReq(searchSearchNodeReq).Execute()

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchSecrets(context.Background()).SearchSearchNodeReq(searchSearchNodeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchSecrets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSecrets`: []ModelSecretRule
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

[**[]ModelSecretRule**](ModelSecretRule.md)

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
    searchSearchScanReq := *openapiclient.NewSearchSearchScanReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchScanReq |  (optional)

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
    searchSearchNodeReq := *openapiclient.NewSearchSearchNodeReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchNodeReq |  (optional)

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
    searchSearchScanReq := *openapiclient.NewSearchSearchScanReq(*openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewSearchSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersMatchFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter("OrderBy_example", []string{"OrderFields_example"})), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // SearchSearchScanReq |  (optional)

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

