# \SearchApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchCloudComplianceScans**](SearchApi.md#SearchCloudComplianceScans) | **Post** /deepfence/search/cloud-compliance/scans | Search Vulnerability Scan results
[**SearchCloudCompliances**](SearchApi.md#SearchCloudCompliances) | **Post** /deepfence/search/cloud-compliances | Search Cloud compliances
[**SearchComplianceScans**](SearchApi.md#SearchComplianceScans) | **Post** /deepfence/search/compliance/scans | Search Vulnerability Scan results
[**SearchCompliances**](SearchApi.md#SearchCompliances) | **Post** /deepfence/search/compliances | Search Compliances
[**SearchContainerImages**](SearchApi.md#SearchContainerImages) | **Post** /deepfence/search/images | Search Container images
[**SearchContainers**](SearchApi.md#SearchContainers) | **Post** /deepfence/search/containers | Search Containers data
[**SearchHosts**](SearchApi.md#SearchHosts) | **Post** /deepfence/search/hosts | Search hosts
[**SearchMalwareScans**](SearchApi.md#SearchMalwareScans) | **Post** /deepfence/search/malware/scans | Search Vulnerability Scan results
[**SearchMalwares**](SearchApi.md#SearchMalwares) | **Post** /deepfence/search/malwares | Search Malwares
[**SearchSecrets**](SearchApi.md#SearchSecrets) | **Post** /deepfence/search/secrets | Search Secrets
[**SearchSecretsScans**](SearchApi.md#SearchSecretsScans) | **Post** /deepfence/search/secret/scans | Search Vulnerability Scan results
[**SearchVulnerabilities**](SearchApi.md#SearchVulnerabilities) | **Post** /deepfence/search/vulnerabilities | Search Vulnerabilities
[**SearchVulnerabilityScans**](SearchApi.md#SearchVulnerabilityScans) | **Post** /deepfence/search/vulnerability/scans | Search Vulnerability Scan results



## SearchCloudComplianceScans

> []ModelScanInfo SearchCloudComplianceScans(ctx).ReportersSearchScanReq(reportersSearchScanReq).Execute()

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
    reportersSearchScanReq := *openapiclient.NewReportersSearchScanReq(*openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}), *openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ReportersSearchScanReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchCloudComplianceScans(context.Background()).ReportersSearchScanReq(reportersSearchScanReq).Execute()
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
 **reportersSearchScanReq** | [**ReportersSearchScanReq**](ReportersSearchScanReq.md) |  | 

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

> []ModelCloudCompliance SearchCloudCompliances(ctx).ReportersSearchFilter(reportersSearchFilter).Execute()

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
    reportersSearchFilter := *openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}) // ReportersSearchFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchCloudCompliances(context.Background()).ReportersSearchFilter(reportersSearchFilter).Execute()
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
 **reportersSearchFilter** | [**ReportersSearchFilter**](ReportersSearchFilter.md) |  | 

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

> []ModelScanInfo SearchComplianceScans(ctx).ReportersSearchScanReq(reportersSearchScanReq).Execute()

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
    reportersSearchScanReq := *openapiclient.NewReportersSearchScanReq(*openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}), *openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ReportersSearchScanReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchComplianceScans(context.Background()).ReportersSearchScanReq(reportersSearchScanReq).Execute()
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
 **reportersSearchScanReq** | [**ReportersSearchScanReq**](ReportersSearchScanReq.md) |  | 

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

> []ModelCompliance SearchCompliances(ctx).ReportersSearchFilter(reportersSearchFilter).Execute()

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
    reportersSearchFilter := *openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}) // ReportersSearchFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchCompliances(context.Background()).ReportersSearchFilter(reportersSearchFilter).Execute()
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
 **reportersSearchFilter** | [**ReportersSearchFilter**](ReportersSearchFilter.md) |  | 

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

> []ModelContainerImage SearchContainerImages(ctx).ReportersSearchFilter(reportersSearchFilter).Execute()

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
    reportersSearchFilter := *openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}) // ReportersSearchFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchContainerImages(context.Background()).ReportersSearchFilter(reportersSearchFilter).Execute()
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
 **reportersSearchFilter** | [**ReportersSearchFilter**](ReportersSearchFilter.md) |  | 

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

> []ModelContainer SearchContainers(ctx).ReportersSearchFilter(reportersSearchFilter).Execute()

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
    reportersSearchFilter := *openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}) // ReportersSearchFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchContainers(context.Background()).ReportersSearchFilter(reportersSearchFilter).Execute()
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
 **reportersSearchFilter** | [**ReportersSearchFilter**](ReportersSearchFilter.md) |  | 

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

> []ModelHost SearchHosts(ctx).ReportersSearchFilter(reportersSearchFilter).Execute()

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
    reportersSearchFilter := *openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}) // ReportersSearchFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchHosts(context.Background()).ReportersSearchFilter(reportersSearchFilter).Execute()
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
 **reportersSearchFilter** | [**ReportersSearchFilter**](ReportersSearchFilter.md) |  | 

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

> []ModelScanInfo SearchMalwareScans(ctx).ReportersSearchScanReq(reportersSearchScanReq).Execute()

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
    reportersSearchScanReq := *openapiclient.NewReportersSearchScanReq(*openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}), *openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ReportersSearchScanReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchMalwareScans(context.Background()).ReportersSearchScanReq(reportersSearchScanReq).Execute()
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
 **reportersSearchScanReq** | [**ReportersSearchScanReq**](ReportersSearchScanReq.md) |  | 

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

> []ModelMalware SearchMalwares(ctx).ReportersSearchFilter(reportersSearchFilter).Execute()

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
    reportersSearchFilter := *openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}) // ReportersSearchFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchMalwares(context.Background()).ReportersSearchFilter(reportersSearchFilter).Execute()
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
 **reportersSearchFilter** | [**ReportersSearchFilter**](ReportersSearchFilter.md) |  | 

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

> []ModelSecret SearchSecrets(ctx).ReportersSearchFilter(reportersSearchFilter).Execute()

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
    reportersSearchFilter := *openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}) // ReportersSearchFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchSecrets(context.Background()).ReportersSearchFilter(reportersSearchFilter).Execute()
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
 **reportersSearchFilter** | [**ReportersSearchFilter**](ReportersSearchFilter.md) |  | 

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

> []ModelScanInfo SearchSecretsScans(ctx).ReportersSearchScanReq(reportersSearchScanReq).Execute()

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
    reportersSearchScanReq := *openapiclient.NewReportersSearchScanReq(*openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}), *openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ReportersSearchScanReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchSecretsScans(context.Background()).ReportersSearchScanReq(reportersSearchScanReq).Execute()
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
 **reportersSearchScanReq** | [**ReportersSearchScanReq**](ReportersSearchScanReq.md) |  | 

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

> []ModelVulnerability SearchVulnerabilities(ctx).ReportersSearchFilter(reportersSearchFilter).Execute()

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
    reportersSearchFilter := *openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}) // ReportersSearchFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchVulnerabilities(context.Background()).ReportersSearchFilter(reportersSearchFilter).Execute()
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
 **reportersSearchFilter** | [**ReportersSearchFilter**](ReportersSearchFilter.md) |  | 

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

> []ModelScanInfo SearchVulnerabilityScans(ctx).ReportersSearchScanReq(reportersSearchScanReq).Execute()

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
    reportersSearchScanReq := *openapiclient.NewReportersSearchScanReq(*openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}), *openapiclient.NewReportersSearchFilter(*openapiclient.NewReportersFieldsFilters(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersOrderFilter(int32(123), "OrderField_example")), []string{"InFieldFilter_example"}), *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ReportersSearchScanReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.SearchVulnerabilityScans(context.Background()).ReportersSearchScanReq(reportersSearchScanReq).Execute()
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
 **reportersSearchScanReq** | [**ReportersSearchScanReq**](ReportersSearchScanReq.md) |  | 

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

