# \CloudNodesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeployCloudResourceAgent**](CloudNodesAPI.md#DeployCloudResourceAgent) | **Post** /deepfence/cloud-resource/deploy-agent | Deploy Agent on Cloud Resource
[**ListCloudNodeAccount**](CloudNodesAPI.md#ListCloudNodeAccount) | **Post** /deepfence/cloud-node/list/accounts | List Cloud Node Accounts
[**ListCloudProviders**](CloudNodesAPI.md#ListCloudProviders) | **Get** /deepfence/cloud-node/list/providers | List Cloud Node Providers
[**RegisterCloudNodeAccount**](CloudNodesAPI.md#RegisterCloudNodeAccount) | **Post** /deepfence/cloud-node/account | Register Cloud Node Account



## DeployCloudResourceAgent

> ModelMessageResponse DeployCloudResourceAgent(ctx).ModelCloudResourceDeployAgentReq(modelCloudResourceDeployAgentReq).Execute()

Deploy Agent on Cloud Resource



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
    modelCloudResourceDeployAgentReq := *openapiclient.NewModelCloudResourceDeployAgentReq([]string{"NodeIds_example"}) // ModelCloudResourceDeployAgentReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudNodesAPI.DeployCloudResourceAgent(context.Background()).ModelCloudResourceDeployAgentReq(modelCloudResourceDeployAgentReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudNodesAPI.DeployCloudResourceAgent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeployCloudResourceAgent`: ModelMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudNodesAPI.DeployCloudResourceAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeployCloudResourceAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelCloudResourceDeployAgentReq** | [**ModelCloudResourceDeployAgentReq**](ModelCloudResourceDeployAgentReq.md) |  | 

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
    modelCloudNodeAccountsListReq := *openapiclient.NewModelCloudNodeAccountsListReq(*openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ModelCloudNodeAccountsListReq |  (optional)

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


## RegisterCloudNodeAccount

> ModelCloudNodeAccountRegisterResp RegisterCloudNodeAccount(ctx).ModelCloudNodeAccountRegisterReq(modelCloudNodeAccountRegisterReq).Execute()

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
    modelCloudNodeAccountRegisterReq := *openapiclient.NewModelCloudNodeAccountRegisterReq("CloudAccount_example", "CloudProvider_example", "NodeId_example") // ModelCloudNodeAccountRegisterReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudNodesAPI.RegisterCloudNodeAccount(context.Background()).ModelCloudNodeAccountRegisterReq(modelCloudNodeAccountRegisterReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudNodesAPI.RegisterCloudNodeAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterCloudNodeAccount`: ModelCloudNodeAccountRegisterResp
    fmt.Fprintf(os.Stdout, "Response from `CloudNodesAPI.RegisterCloudNodeAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterCloudNodeAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelCloudNodeAccountRegisterReq** | [**ModelCloudNodeAccountRegisterReq**](ModelCloudNodeAccountRegisterReq.md) |  | 

### Return type

[**ModelCloudNodeAccountRegisterResp**](ModelCloudNodeAccountRegisterResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

