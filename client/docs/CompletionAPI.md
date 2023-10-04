# \CompletionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteProcessInfo**](CompletionAPI.md#CompleteProcessInfo) | **Post** /deepfence/complete/process | Get Completion for process fields



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

