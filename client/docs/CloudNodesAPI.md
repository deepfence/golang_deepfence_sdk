# \CloudNodesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCloudNodeAccount**](CloudNodesAPI.md#DeleteCloudNodeAccount) | **Patch** /deepfence/cloud-node/account/delete | Delete Cloud Node Account
[**ListCloudNodeAccount**](CloudNodesAPI.md#ListCloudNodeAccount) | **Post** /deepfence/cloud-node/list/accounts | List Cloud Node Accounts
[**ListCloudProviders**](CloudNodesAPI.md#ListCloudProviders) | **Get** /deepfence/cloud-node/list/providers | List Cloud Node Providers
[**RefreshCloudNodeAccount**](CloudNodesAPI.md#RefreshCloudNodeAccount) | **Post** /deepfence/cloud-node/account/refresh | Refresh Cloud Account
[**RegisterCloudNodeAccount**](CloudNodesAPI.md#RegisterCloudNodeAccount) | **Post** /deepfence/cloud-node/account | Register Cloud Node Account



## DeleteCloudNodeAccount

> DeleteCloudNodeAccount(ctx).ModelCloudAccountDeleteReq(modelCloudAccountDeleteReq).Execute()

Delete Cloud Node Account



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
	modelCloudAccountDeleteReq := *openapiclient.NewModelCloudAccountDeleteReq([]string{"NodeIds_example"}) // ModelCloudAccountDeleteReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudNodesAPI.DeleteCloudNodeAccount(context.Background()).ModelCloudAccountDeleteReq(modelCloudAccountDeleteReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudNodesAPI.DeleteCloudNodeAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCloudNodeAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelCloudAccountDeleteReq** | [**ModelCloudAccountDeleteReq**](ModelCloudAccountDeleteReq.md) |  | 

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


## ListCloudNodeAccount

> ModelCloudNodeAccountsListResp ListCloudNodeAccount(ctx).ModelCloudNodeAccountsListReq(modelCloudNodeAccountsListReq).Execute()

List Cloud Node Accounts



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
	modelCloudNodeAccountsListReq := *openapiclient.NewModelCloudNodeAccountsListReq("CloudProvider_example", *openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ModelCloudNodeAccountsListReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudNodesAPI.ListCloudNodeAccount(context.Background()).ModelCloudNodeAccountsListReq(modelCloudNodeAccountsListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudNodesAPI.ListCloudNodeAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCloudNodeAccount`: ModelCloudNodeAccountsListResp
	fmt.Fprintf(os.Stdout, "Response from `CloudNodesAPI.ListCloudNodeAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCloudNodeAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelCloudNodeAccountsListReq** | [**ModelCloudNodeAccountsListReq**](ModelCloudNodeAccountsListReq.md) |  | 

### Return type

[**ModelCloudNodeAccountsListResp**](ModelCloudNodeAccountsListResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudProviders

> ModelCloudNodeProvidersListResp ListCloudProviders(ctx).Execute()

List Cloud Node Providers



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
	resp, r, err := apiClient.CloudNodesAPI.ListCloudProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudNodesAPI.ListCloudProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCloudProviders`: ModelCloudNodeProvidersListResp
	fmt.Fprintf(os.Stdout, "Response from `CloudNodesAPI.ListCloudProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCloudProvidersRequest struct via the builder pattern


### Return type

[**ModelCloudNodeProvidersListResp**](ModelCloudNodeProvidersListResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshCloudNodeAccount

> RefreshCloudNodeAccount(ctx).ModelCloudAccountRefreshReq(modelCloudAccountRefreshReq).Execute()

Refresh Cloud Account



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
	modelCloudAccountRefreshReq := *openapiclient.NewModelCloudAccountRefreshReq([]string{"NodeIds_example"}) // ModelCloudAccountRefreshReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudNodesAPI.RefreshCloudNodeAccount(context.Background()).ModelCloudAccountRefreshReq(modelCloudAccountRefreshReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudNodesAPI.RefreshCloudNodeAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshCloudNodeAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelCloudAccountRefreshReq** | [**ModelCloudAccountRefreshReq**](ModelCloudAccountRefreshReq.md) |  | 

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


## RegisterCloudNodeAccount

> RegisterCloudNodeAccount(ctx).ModelCloudNodeAccountRegisterReq(modelCloudNodeAccountRegisterReq).Execute()

Register Cloud Node Account



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
	modelCloudNodeAccountRegisterReq := *openapiclient.NewModelCloudNodeAccountRegisterReq("AccountId_example", "CloudProvider_example", "HostNodeId_example", "NodeId_example", "Version_example") // ModelCloudNodeAccountRegisterReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudNodesAPI.RegisterCloudNodeAccount(context.Background()).ModelCloudNodeAccountRegisterReq(modelCloudNodeAccountRegisterReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudNodesAPI.RegisterCloudNodeAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterCloudNodeAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelCloudNodeAccountRegisterReq** | [**ModelCloudNodeAccountRegisterReq**](ModelCloudNodeAccountRegisterReq.md) |  | 

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

