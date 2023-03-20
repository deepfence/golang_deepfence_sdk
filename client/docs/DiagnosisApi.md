# \DiagnosisApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiagnosticNotification**](DiagnosisApi.md#DiagnosticNotification) | **Get** /deepfence/diagnosis/notification | Get Diagnostic Notification
[**GenerateAgentDiagnosticLogs**](DiagnosisApi.md#GenerateAgentDiagnosticLogs) | **Post** /deepfence/diagnosis/agent-logs | Generate Agent Diagnostic Logs
[**GenerateConsoleDiagnosticLogs**](DiagnosisApi.md#GenerateConsoleDiagnosticLogs) | **Post** /deepfence/diagnosis/console-logs | Generate Console Diagnostic Logs
[**GetDiagnosticLogs**](DiagnosisApi.md#GetDiagnosticLogs) | **Get** /deepfence/diagnosis/diagnostic-logs | Get Diagnostic Logs
[**UpdateAgentDiagnosticLogsStatus**](DiagnosisApi.md#UpdateAgentDiagnosticLogsStatus) | **Post** /deepfence/diagnosis/agent-logs/status/{node_id} | Update Agent Diagnostic Logs Status



## DiagnosticNotification

> []DiagnosisDiagnosticNotification DiagnosticNotification(ctx).Execute()

Get Diagnostic Notification



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
    resp, r, err := apiClient.DiagnosisApi.DiagnosticNotification(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiagnosisApi.DiagnosticNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DiagnosticNotification`: []DiagnosisDiagnosticNotification
    fmt.Fprintf(os.Stdout, "Response from `DiagnosisApi.DiagnosticNotification`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDiagnosticNotificationRequest struct via the builder pattern


### Return type

[**[]DiagnosisDiagnosticNotification**](DiagnosisDiagnosticNotification.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateAgentDiagnosticLogs

> GenerateAgentDiagnosticLogs(ctx).DiagnosisGenerateAgentDiagnosticLogsRequest(diagnosisGenerateAgentDiagnosticLogsRequest).Execute()

Generate Agent Diagnostic Logs



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
    diagnosisGenerateAgentDiagnosticLogsRequest := *openapiclient.NewDiagnosisGenerateAgentDiagnosticLogsRequest([]openapiclient.DiagnosisNodeIdentifier{*openapiclient.NewDiagnosisNodeIdentifier("NodeId_example", "NodeType_example")}, int32(123)) // DiagnosisGenerateAgentDiagnosticLogsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DiagnosisApi.GenerateAgentDiagnosticLogs(context.Background()).DiagnosisGenerateAgentDiagnosticLogsRequest(diagnosisGenerateAgentDiagnosticLogsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiagnosisApi.GenerateAgentDiagnosticLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateAgentDiagnosticLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **diagnosisGenerateAgentDiagnosticLogsRequest** | [**DiagnosisGenerateAgentDiagnosticLogsRequest**](DiagnosisGenerateAgentDiagnosticLogsRequest.md) |  | 

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


## GenerateConsoleDiagnosticLogs

> GenerateConsoleDiagnosticLogs(ctx).DiagnosisGenerateConsoleDiagnosticLogsRequest(diagnosisGenerateConsoleDiagnosticLogsRequest).Execute()

Generate Console Diagnostic Logs



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
    diagnosisGenerateConsoleDiagnosticLogsRequest := *openapiclient.NewDiagnosisGenerateConsoleDiagnosticLogsRequest(int32(123)) // DiagnosisGenerateConsoleDiagnosticLogsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DiagnosisApi.GenerateConsoleDiagnosticLogs(context.Background()).DiagnosisGenerateConsoleDiagnosticLogsRequest(diagnosisGenerateConsoleDiagnosticLogsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiagnosisApi.GenerateConsoleDiagnosticLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateConsoleDiagnosticLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **diagnosisGenerateConsoleDiagnosticLogsRequest** | [**DiagnosisGenerateConsoleDiagnosticLogsRequest**](DiagnosisGenerateConsoleDiagnosticLogsRequest.md) |  | 

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


## GetDiagnosticLogs

> DiagnosisGetDiagnosticLogsResponse GetDiagnosticLogs(ctx).Execute()

Get Diagnostic Logs



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
    resp, r, err := apiClient.DiagnosisApi.GetDiagnosticLogs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiagnosisApi.GetDiagnosticLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiagnosticLogs`: DiagnosisGetDiagnosticLogsResponse
    fmt.Fprintf(os.Stdout, "Response from `DiagnosisApi.GetDiagnosticLogs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiagnosticLogsRequest struct via the builder pattern


### Return type

[**DiagnosisGetDiagnosticLogsResponse**](DiagnosisGetDiagnosticLogsResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAgentDiagnosticLogsStatus

> UpdateAgentDiagnosticLogsStatus(ctx, nodeId).DiagnosisDiagnosticLogsStatus(diagnosisDiagnosticLogsStatus).Execute()

Update Agent Diagnostic Logs Status



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
    nodeId := "nodeId_example" // string | 
    diagnosisDiagnosticLogsStatus := *openapiclient.NewDiagnosisDiagnosticLogsStatus("Message_example", "Status_example") // DiagnosisDiagnosticLogsStatus |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DiagnosisApi.UpdateAgentDiagnosticLogsStatus(context.Background(), nodeId).DiagnosisDiagnosticLogsStatus(diagnosisDiagnosticLogsStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiagnosisApi.UpdateAgentDiagnosticLogsStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgentDiagnosticLogsStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **diagnosisDiagnosticLogsStatus** | [**DiagnosisDiagnosticLogsStatus**](DiagnosisDiagnosticLogsStatus.md) |  | 

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

