# \IntegrationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIntegration**](IntegrationAPI.md#AddIntegration) | **Post** /deepfence/integration | Add Integration
[**DeleteIntegration**](IntegrationAPI.md#DeleteIntegration) | **Delete** /deepfence/integration/{integration_id} | Delete Integration
[**ListIntegration**](IntegrationAPI.md#ListIntegration) | **Get** /deepfence/integration | List Integrations
[**UpdateIntegration**](IntegrationAPI.md#UpdateIntegration) | **Put** /deepfence/integration/{integration_id} | Update Integration



## AddIntegration

> ModelMessageResponse AddIntegration(ctx).ModelIntegrationAddReq(modelIntegrationAddReq).Execute()

Add Integration



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
	modelIntegrationAddReq := *openapiclient.NewModelIntegrationAddReq() // ModelIntegrationAddReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.AddIntegration(context.Background()).ModelIntegrationAddReq(modelIntegrationAddReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.AddIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddIntegration`: ModelMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.AddIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelIntegrationAddReq** | [**ModelIntegrationAddReq**](ModelIntegrationAddReq.md) |  | 

### Return type

[**ModelMessageResponse**](ModelMessageResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIntegration

> DeleteIntegration(ctx, integrationId).Execute()

Delete Integration



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
	integrationId := "integrationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationAPI.DeleteIntegration(context.Background(), integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.DeleteIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegration

> []ModelIntegrationListResp ListIntegration(ctx).Execute()

List Integrations



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
	resp, r, err := apiClient.IntegrationAPI.ListIntegration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.ListIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntegration`: []ModelIntegrationListResp
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.ListIntegration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationRequest struct via the builder pattern


### Return type

[**[]ModelIntegrationListResp**](ModelIntegrationListResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIntegration

> ModelMessageResponse UpdateIntegration(ctx, integrationId).ModelIntegrationUpdateReq(modelIntegrationUpdateReq).Execute()

Update Integration



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
	integrationId := "integrationId_example" // string | 
	modelIntegrationUpdateReq := *openapiclient.NewModelIntegrationUpdateReq() // ModelIntegrationUpdateReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.UpdateIntegration(context.Background(), integrationId).ModelIntegrationUpdateReq(modelIntegrationUpdateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.UpdateIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIntegration`: ModelMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.UpdateIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelIntegrationUpdateReq** | [**ModelIntegrationUpdateReq**](ModelIntegrationUpdateReq.md) |  | 

### Return type

[**ModelMessageResponse**](ModelMessageResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

