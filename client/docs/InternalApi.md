# \InternalApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConsoleApiToken**](InternalApi.md#GetConsoleApiToken) | **Get** /deepfence/internal/console-api-token | Get api-token for console agent



## GetConsoleApiToken

> ModelApiAuthRequest GetConsoleApiToken(ctx).Execute()

Get api-token for console agent



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.GetConsoleApiToken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.GetConsoleApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConsoleApiToken`: ModelApiAuthRequest
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.GetConsoleApiToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsoleApiTokenRequest struct via the builder pattern


### Return type

[**ModelApiAuthRequest**](ModelApiAuthRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

