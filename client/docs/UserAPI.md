# \UserAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCurrentUser**](UserAPI.md#DeleteCurrentUser) | **Delete** /deepfence/user | Delete Current User
[**DeleteUser**](UserAPI.md#DeleteUser) | **Delete** /deepfence/users/{id} | Delete User by User ID
[**GetApiTokens**](UserAPI.md#GetApiTokens) | **Get** /deepfence/api-token | Get User&#39;s API Tokens
[**GetCurrentUser**](UserAPI.md#GetCurrentUser) | **Get** /deepfence/user | Get Current User
[**GetUser**](UserAPI.md#GetUser) | **Get** /deepfence/users/{id} | Get User by User ID
[**GetUsers**](UserAPI.md#GetUsers) | **Get** /deepfence/users | Get all users
[**InviteUser**](UserAPI.md#InviteUser) | **Post** /deepfence/user/invite | Invite User
[**RegisterInvitedUser**](UserAPI.md#RegisterInvitedUser) | **Post** /deepfence/user/invite/register | Register Invited User
[**RegisterUser**](UserAPI.md#RegisterUser) | **Post** /deepfence/user/register | Register User
[**ResetApiTokens**](UserAPI.md#ResetApiTokens) | **Post** /deepfence/api-token/reset | Reset User&#39;s API Tokens
[**ResetPasswordRequest**](UserAPI.md#ResetPasswordRequest) | **Post** /deepfence/user/reset-password/request | Reset Password Request
[**UpdateCurrentUser**](UserAPI.md#UpdateCurrentUser) | **Put** /deepfence/user | Update Current User
[**UpdatePassword**](UserAPI.md#UpdatePassword) | **Put** /deepfence/user/password | Update Password
[**UpdateUser**](UserAPI.md#UpdateUser) | **Put** /deepfence/users/{id} | Update User by User ID
[**VerifyResetPasswordRequest**](UserAPI.md#VerifyResetPasswordRequest) | **Post** /deepfence/user/reset-password/verify | Verify and Reset Password



## DeleteCurrentUser

> DeleteCurrentUser(ctx).Execute()

Delete Current User



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
    r, err := apiClient.UserAPI.DeleteCurrentUser(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DeleteCurrentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCurrentUserRequest struct via the builder pattern


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


## DeleteUser

> DeleteUser(ctx, id).Execute()

Delete User by User ID



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
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserAPI.DeleteUser(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


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


## GetApiTokens

> []ModelAPITokenResponse GetApiTokens(ctx).Execute()

Get User's API Tokens



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
    resp, r, err := apiClient.UserAPI.GetApiTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetApiTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiTokens`: []ModelAPITokenResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetApiTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiTokensRequest struct via the builder pattern


### Return type

[**[]ModelAPITokenResponse**](ModelAPITokenResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUser

> ModelUser GetCurrentUser(ctx).Execute()

Get Current User



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
    resp, r, err := apiClient.UserAPI.GetCurrentUser(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetCurrentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentUser`: ModelUser
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetCurrentUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserRequest struct via the builder pattern


### Return type

[**ModelUser**](ModelUser.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> ModelUser GetUser(ctx, id).Execute()

Get User by User ID



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
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.GetUser(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: ModelUser
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelUser**](ModelUser.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> []ModelUser GetUsers(ctx).Execute()

Get all users



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
    resp, r, err := apiClient.UserAPI.GetUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsers`: []ModelUser
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


### Return type

[**[]ModelUser**](ModelUser.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteUser

> ModelInviteUserResponse InviteUser(ctx).ModelInviteUserRequest(modelInviteUserRequest).Execute()

Invite User



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
    modelInviteUserRequest := *openapiclient.NewModelInviteUserRequest("Action_example", "Email_example", "Role_example") // ModelInviteUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.InviteUser(context.Background()).ModelInviteUserRequest(modelInviteUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.InviteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InviteUser`: ModelInviteUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.InviteUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInviteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelInviteUserRequest** | [**ModelInviteUserRequest**](ModelInviteUserRequest.md) |  | 

### Return type

[**ModelInviteUserResponse**](ModelInviteUserResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterInvitedUser

> ModelLoginResponse RegisterInvitedUser(ctx).ModelRegisterInvitedUserRequest(modelRegisterInvitedUserRequest).Execute()

Register Invited User



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
    modelRegisterInvitedUserRequest := *openapiclient.NewModelRegisterInvitedUserRequest("Code_example", "FirstName_example", "LastName_example", "Namespace_example", "Password_example") // ModelRegisterInvitedUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.RegisterInvitedUser(context.Background()).ModelRegisterInvitedUserRequest(modelRegisterInvitedUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.RegisterInvitedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterInvitedUser`: ModelLoginResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.RegisterInvitedUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterInvitedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelRegisterInvitedUserRequest** | [**ModelRegisterInvitedUserRequest**](ModelRegisterInvitedUserRequest.md) |  | 

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


## RegisterUser

> ModelLoginResponse RegisterUser(ctx).ModelUserRegisterRequest(modelUserRegisterRequest).Execute()

Register User



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
    modelUserRegisterRequest := *openapiclient.NewModelUserRegisterRequest("Company_example", "ConsoleUrl_example", "Email_example", "FirstName_example", "LastName_example", "Password_example") // ModelUserRegisterRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.RegisterUser(context.Background()).ModelUserRegisterRequest(modelUserRegisterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.RegisterUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterUser`: ModelLoginResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.RegisterUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelUserRegisterRequest** | [**ModelUserRegisterRequest**](ModelUserRegisterRequest.md) |  | 

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


## ResetApiTokens

> []ModelAPITokenResponse ResetApiTokens(ctx).Execute()

Reset User's API Tokens



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
    resp, r, err := apiClient.UserAPI.ResetApiTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ResetApiTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetApiTokens`: []ModelAPITokenResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.ResetApiTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiResetApiTokensRequest struct via the builder pattern


### Return type

[**[]ModelAPITokenResponse**](ModelAPITokenResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPasswordRequest

> ModelMessageResponse ResetPasswordRequest(ctx).ModelPasswordResetRequest(modelPasswordResetRequest).Execute()

Reset Password Request



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
    modelPasswordResetRequest := *openapiclient.NewModelPasswordResetRequest("Email_example") // ModelPasswordResetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.ResetPasswordRequest(context.Background()).ModelPasswordResetRequest(modelPasswordResetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ResetPasswordRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetPasswordRequest`: ModelMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.ResetPasswordRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelPasswordResetRequest** | [**ModelPasswordResetRequest**](ModelPasswordResetRequest.md) |  | 

### Return type

[**ModelMessageResponse**](ModelMessageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCurrentUser

> ModelUser UpdateCurrentUser(ctx).ModelUpdateUserRequest(modelUpdateUserRequest).Execute()

Update Current User



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
    modelUpdateUserRequest := *openapiclient.NewModelUpdateUserRequest() // ModelUpdateUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.UpdateCurrentUser(context.Background()).ModelUpdateUserRequest(modelUpdateUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateCurrentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCurrentUser`: ModelUser
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.UpdateCurrentUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCurrentUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelUpdateUserRequest** | [**ModelUpdateUserRequest**](ModelUpdateUserRequest.md) |  | 

### Return type

[**ModelUser**](ModelUser.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePassword

> UpdatePassword(ctx).ModelUpdateUserPasswordRequest(modelUpdateUserPasswordRequest).Execute()

Update Password



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
    modelUpdateUserPasswordRequest := *openapiclient.NewModelUpdateUserPasswordRequest("NewPassword_example", "OldPassword_example") // ModelUpdateUserPasswordRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserAPI.UpdatePassword(context.Background()).ModelUpdateUserPasswordRequest(modelUpdateUserPasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdatePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelUpdateUserPasswordRequest** | [**ModelUpdateUserPasswordRequest**](ModelUpdateUserPasswordRequest.md) |  | 

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


## UpdateUser

> ModelUser UpdateUser(ctx, id).ModelUpdateUserIDRequest(modelUpdateUserIDRequest).Execute()

Update User by User ID



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
    id := int32(56) // int32 | 
    modelUpdateUserIDRequest := *openapiclient.NewModelUpdateUserIDRequest() // ModelUpdateUserIDRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAPI.UpdateUser(context.Background(), id).ModelUpdateUserIDRequest(modelUpdateUserIDRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: ModelUser
    fmt.Fprintf(os.Stdout, "Response from `UserAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelUpdateUserIDRequest** | [**ModelUpdateUserIDRequest**](ModelUpdateUserIDRequest.md) |  | 

### Return type

[**ModelUser**](ModelUser.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyResetPasswordRequest

> VerifyResetPasswordRequest(ctx).ModelPasswordResetVerifyRequest(modelPasswordResetVerifyRequest).Execute()

Verify and Reset Password



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
    modelPasswordResetVerifyRequest := *openapiclient.NewModelPasswordResetVerifyRequest("Code_example", "Namespace_example", "Password_example") // ModelPasswordResetVerifyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserAPI.VerifyResetPasswordRequest(context.Background()).ModelPasswordResetVerifyRequest(modelPasswordResetVerifyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.VerifyResetPasswordRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyResetPasswordRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelPasswordResetVerifyRequest** | [**ModelPasswordResetVerifyRequest**](ModelPasswordResetVerifyRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

