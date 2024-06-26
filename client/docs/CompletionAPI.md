# \CompletionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteCloudAccount**](CompletionAPI.md#CompleteCloudAccount) | **Post** /deepfence/complete/cloud-account | Get Completion for cloud account fields
[**CompleteCloudCompliance**](CompletionAPI.md#CompleteCloudCompliance) | **Post** /deepfence/complete/cloud-compliance | Get Completion for cloud compliance fields
[**CompleteCloudResources**](CompletionAPI.md#CompleteCloudResources) | **Post** /deepfence/complete/cloud-resources | Get Completion for cloud resources fields
[**CompleteComplianceInfo**](CompletionAPI.md#CompleteComplianceInfo) | **Post** /deepfence/complete/compliance | Get Completion for compliance fields
[**CompleteContainerInfo**](CompletionAPI.md#CompleteContainerInfo) | **Post** /deepfence/complete/container | Get Completion for Container fields
[**CompleteHostInfo**](CompletionAPI.md#CompleteHostInfo) | **Post** /deepfence/complete/host | Get Completion for host fields
[**CompletePodInfo**](CompletionAPI.md#CompletePodInfo) | **Post** /deepfence/complete/pod | Get Completion for Pod fields
[**CompleteProcessInfo**](CompletionAPI.md#CompleteProcessInfo) | **Post** /deepfence/complete/process | Get Completion for process fields
[**CompleteVulnerabilityInfo**](CompletionAPI.md#CompleteVulnerabilityInfo) | **Post** /deepfence/complete/vulnerability | Get Completion for vulnerability fields



## CompleteCloudAccount

> CompletionCompletionNodeFieldRes CompleteCloudAccount(ctx).CompletionCompletionNodeFieldReq(completionCompletionNodeFieldReq).Execute()

Get Completion for cloud account fields



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
	completionCompletionNodeFieldReq := *openapiclient.NewCompletionCompletionNodeFieldReq("Completion_example", "FieldName_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // CompletionCompletionNodeFieldReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompletionAPI.CompleteCloudAccount(context.Background()).CompletionCompletionNodeFieldReq(completionCompletionNodeFieldReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompletionAPI.CompleteCloudAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteCloudAccount`: CompletionCompletionNodeFieldRes
	fmt.Fprintf(os.Stdout, "Response from `CompletionAPI.CompleteCloudAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteCloudAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **completionCompletionNodeFieldReq** | [**CompletionCompletionNodeFieldReq**](CompletionCompletionNodeFieldReq.md) |  | 

### Return type

[**CompletionCompletionNodeFieldRes**](CompletionCompletionNodeFieldRes.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteCloudCompliance

> CompletionCompletionNodeFieldRes CompleteCloudCompliance(ctx).CompletionCompletionNodeFieldReq(completionCompletionNodeFieldReq).Execute()

Get Completion for cloud compliance fields



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
	completionCompletionNodeFieldReq := *openapiclient.NewCompletionCompletionNodeFieldReq("Completion_example", "FieldName_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // CompletionCompletionNodeFieldReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompletionAPI.CompleteCloudCompliance(context.Background()).CompletionCompletionNodeFieldReq(completionCompletionNodeFieldReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompletionAPI.CompleteCloudCompliance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteCloudCompliance`: CompletionCompletionNodeFieldRes
	fmt.Fprintf(os.Stdout, "Response from `CompletionAPI.CompleteCloudCompliance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteCloudComplianceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **completionCompletionNodeFieldReq** | [**CompletionCompletionNodeFieldReq**](CompletionCompletionNodeFieldReq.md) |  | 

### Return type

[**CompletionCompletionNodeFieldRes**](CompletionCompletionNodeFieldRes.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteCloudResources

> CompletionCompletionNodeFieldRes CompleteCloudResources(ctx).CompletionCompletionNodeFieldReq(completionCompletionNodeFieldReq).Execute()

Get Completion for cloud resources fields



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
	completionCompletionNodeFieldReq := *openapiclient.NewCompletionCompletionNodeFieldReq("Completion_example", "FieldName_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // CompletionCompletionNodeFieldReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompletionAPI.CompleteCloudResources(context.Background()).CompletionCompletionNodeFieldReq(completionCompletionNodeFieldReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompletionAPI.CompleteCloudResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteCloudResources`: CompletionCompletionNodeFieldRes
	fmt.Fprintf(os.Stdout, "Response from `CompletionAPI.CompleteCloudResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteCloudResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **completionCompletionNodeFieldReq** | [**CompletionCompletionNodeFieldReq**](CompletionCompletionNodeFieldReq.md) |  | 

### Return type

[**CompletionCompletionNodeFieldRes**](CompletionCompletionNodeFieldRes.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteComplianceInfo

> CompletionCompletionNodeFieldRes CompleteComplianceInfo(ctx).CompletionCompletionNodeFieldReq(completionCompletionNodeFieldReq).Execute()

Get Completion for compliance fields



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
	completionCompletionNodeFieldReq := *openapiclient.NewCompletionCompletionNodeFieldReq("Completion_example", "FieldName_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // CompletionCompletionNodeFieldReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompletionAPI.CompleteComplianceInfo(context.Background()).CompletionCompletionNodeFieldReq(completionCompletionNodeFieldReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompletionAPI.CompleteComplianceInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteComplianceInfo`: CompletionCompletionNodeFieldRes
	fmt.Fprintf(os.Stdout, "Response from `CompletionAPI.CompleteComplianceInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteComplianceInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **completionCompletionNodeFieldReq** | [**CompletionCompletionNodeFieldReq**](CompletionCompletionNodeFieldReq.md) |  | 

### Return type

[**CompletionCompletionNodeFieldRes**](CompletionCompletionNodeFieldRes.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteContainerInfo

> CompletionCompletionNodeFieldRes CompleteContainerInfo(ctx).CompletionCompletionNodeFieldReq(completionCompletionNodeFieldReq).Execute()

Get Completion for Container fields



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
	completionCompletionNodeFieldReq := *openapiclient.NewCompletionCompletionNodeFieldReq("Completion_example", "FieldName_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // CompletionCompletionNodeFieldReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompletionAPI.CompleteContainerInfo(context.Background()).CompletionCompletionNodeFieldReq(completionCompletionNodeFieldReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompletionAPI.CompleteContainerInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteContainerInfo`: CompletionCompletionNodeFieldRes
	fmt.Fprintf(os.Stdout, "Response from `CompletionAPI.CompleteContainerInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteContainerInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **completionCompletionNodeFieldReq** | [**CompletionCompletionNodeFieldReq**](CompletionCompletionNodeFieldReq.md) |  | 

### Return type

[**CompletionCompletionNodeFieldRes**](CompletionCompletionNodeFieldRes.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteHostInfo

> CompletionCompletionNodeFieldRes CompleteHostInfo(ctx).CompletionCompletionNodeFieldReq(completionCompletionNodeFieldReq).Execute()

Get Completion for host fields



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
	completionCompletionNodeFieldReq := *openapiclient.NewCompletionCompletionNodeFieldReq("Completion_example", "FieldName_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // CompletionCompletionNodeFieldReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompletionAPI.CompleteHostInfo(context.Background()).CompletionCompletionNodeFieldReq(completionCompletionNodeFieldReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompletionAPI.CompleteHostInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteHostInfo`: CompletionCompletionNodeFieldRes
	fmt.Fprintf(os.Stdout, "Response from `CompletionAPI.CompleteHostInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteHostInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **completionCompletionNodeFieldReq** | [**CompletionCompletionNodeFieldReq**](CompletionCompletionNodeFieldReq.md) |  | 

### Return type

[**CompletionCompletionNodeFieldRes**](CompletionCompletionNodeFieldRes.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompletePodInfo

> CompletionCompletionNodeFieldRes CompletePodInfo(ctx).CompletionCompletionNodeFieldReq(completionCompletionNodeFieldReq).Execute()

Get Completion for Pod fields



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
	completionCompletionNodeFieldReq := *openapiclient.NewCompletionCompletionNodeFieldReq("Completion_example", "FieldName_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // CompletionCompletionNodeFieldReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompletionAPI.CompletePodInfo(context.Background()).CompletionCompletionNodeFieldReq(completionCompletionNodeFieldReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompletionAPI.CompletePodInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompletePodInfo`: CompletionCompletionNodeFieldRes
	fmt.Fprintf(os.Stdout, "Response from `CompletionAPI.CompletePodInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompletePodInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **completionCompletionNodeFieldReq** | [**CompletionCompletionNodeFieldReq**](CompletionCompletionNodeFieldReq.md) |  | 

### Return type

[**CompletionCompletionNodeFieldRes**](CompletionCompletionNodeFieldRes.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteProcessInfo

> CompletionCompletionNodeFieldRes CompleteProcessInfo(ctx).CompletionCompletionNodeFieldReq(completionCompletionNodeFieldReq).Execute()

Get Completion for process fields



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
	completionCompletionNodeFieldReq := *openapiclient.NewCompletionCompletionNodeFieldReq("Completion_example", "FieldName_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // CompletionCompletionNodeFieldReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompletionAPI.CompleteProcessInfo(context.Background()).CompletionCompletionNodeFieldReq(completionCompletionNodeFieldReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompletionAPI.CompleteProcessInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteProcessInfo`: CompletionCompletionNodeFieldRes
	fmt.Fprintf(os.Stdout, "Response from `CompletionAPI.CompleteProcessInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteProcessInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **completionCompletionNodeFieldReq** | [**CompletionCompletionNodeFieldReq**](CompletionCompletionNodeFieldReq.md) |  | 

### Return type

[**CompletionCompletionNodeFieldRes**](CompletionCompletionNodeFieldRes.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteVulnerabilityInfo

> CompletionCompletionNodeFieldRes CompleteVulnerabilityInfo(ctx).CompletionCompletionNodeFieldReq(completionCompletionNodeFieldReq).Execute()

Get Completion for vulnerability fields



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
	completionCompletionNodeFieldReq := *openapiclient.NewCompletionCompletionNodeFieldReq("Completion_example", "FieldName_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // CompletionCompletionNodeFieldReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompletionAPI.CompleteVulnerabilityInfo(context.Background()).CompletionCompletionNodeFieldReq(completionCompletionNodeFieldReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompletionAPI.CompleteVulnerabilityInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteVulnerabilityInfo`: CompletionCompletionNodeFieldRes
	fmt.Fprintf(os.Stdout, "Response from `CompletionAPI.CompleteVulnerabilityInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteVulnerabilityInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **completionCompletionNodeFieldReq** | [**CompletionCompletionNodeFieldReq**](CompletionCompletionNodeFieldReq.md) |  | 

### Return type

[**CompletionCompletionNodeFieldRes**](CompletionCompletionNodeFieldRes.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

