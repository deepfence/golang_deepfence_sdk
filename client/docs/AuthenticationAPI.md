# \AuthenticationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthToken**](AuthenticationAPI.md#AuthToken) | **Post** /deepfence/auth/token | Get Access Token for API Token
[**AuthTokenRefresh**](AuthenticationAPI.md#AuthTokenRefresh) | **Post** /deepfence/auth/token/refresh | Refresh access token
[**Login**](AuthenticationAPI.md#Login) | **Post** /deepfence/user/login | Login API
[**Logout**](AuthenticationAPI.md#Logout) | **Post** /deepfence/user/logout | Logout API



## AuthToken

> ModelResponseAccessToken AuthToken(ctx).ModelApiAuthRequest(modelApiAuthRequest).Execute()

Get Access Token for API Token



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
    modelApiAuthRequest := *openapiclient.NewModelApiAuthRequest("ApiToken_example") // ModelApiAuthRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationAPI.AuthToken(context.Background()).ModelApiAuthRequest(modelApiAuthRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AuthToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthToken`: ModelResponseAccessToken
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AuthToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelApiAuthRequest** | [**ModelApiAuthRequest**](ModelApiAuthRequest.md) |  | 

### Return type

[**ModelResponseAccessToken**](ModelResponseAccessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenRefresh

> ModelResponseAccessToken AuthTokenRefresh(ctx).Execute()

Refresh access token



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
    resp, r, err := apiClient.AuthenticationAPI.AuthTokenRefresh(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AuthTokenRefresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokenRefresh`: ModelResponseAccessToken
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AuthTokenRefresh`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenRefreshRequest struct via the builder pattern


### Return type

[**ModelResponseAccessToken**](ModelResponseAccessToken.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> ModelLoginResponse Login(ctx).ModelLoginRequest(modelLoginRequest).Execute()

Login API



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
    modelLoginRequest := *openapiclient.NewModelLoginRequest("Email_example", "Password_example") // ModelLoginRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationAPI.Login(context.Background()).ModelLoginRequest(modelLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Login``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Login`: ModelLoginResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.Login`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelLoginRequest** | [**ModelLoginRequest**](ModelLoginRequest.md) |  | 

### Return type

[**ModelLoginResponse**](ModelLoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logout

> Logout(ctx).Execute()

Logout API



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
    r, err := apiClient.AuthenticationAPI.Logout(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Logout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


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

