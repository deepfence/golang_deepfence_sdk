# \DiagnosisAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiagnosticNotification**](DiagnosisAPI.md#DiagnosticNotification) | **Get** /deepfence/diagnosis/notification | Get Diagnostic Notification
[**GenerateAgentDiagnosticLogs**](DiagnosisAPI.md#GenerateAgentDiagnosticLogs) | **Post** /deepfence/diagnosis/agent-logs | Generate Agent Diagnostic Logs
[**GenerateCloudScannerDiagnosticLogs**](DiagnosisAPI.md#GenerateCloudScannerDiagnosticLogs) | **Post** /deepfence/diagnosis/cloud-scanner-logs | Generate Cloud Scanner Diagnostic Logs
[**GenerateConsoleDiagnosticLogs**](DiagnosisAPI.md#GenerateConsoleDiagnosticLogs) | **Post** /deepfence/diagnosis/console-logs | Generate Console Diagnostic Logs
[**GetDiagnosticLogs**](DiagnosisAPI.md#GetDiagnosticLogs) | **Get** /deepfence/diagnosis/diagnostic-logs | Get Diagnostic Logs
[**UpdateAgentDiagnosticLogsStatus**](DiagnosisAPI.md#UpdateAgentDiagnosticLogsStatus) | **Put** /deepfence/diagnosis/agent-logs/status/{node_id} | Update Agent Diagnostic Logs Status
[**UpdateCloudScannerDiagnosticLogsStatus**](DiagnosisAPI.md#UpdateCloudScannerDiagnosticLogsStatus) | **Put** /deepfence/diagnosis/cloud-scanner-logs/status/{node_id} | Update Cloud Scanner Diagnostic Logs Status



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
	resp, r, err := apiClient.DiagnosisAPI.DiagnosticNotification(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiagnosisAPI.DiagnosticNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiagnosticNotification`: []DiagnosisDiagnosticNotification
	fmt.Fprintf(os.Stdout, "Response from `DiagnosisAPI.DiagnosticNotification`: %v\n", resp)
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
	r, err := apiClient.DiagnosisAPI.GenerateAgentDiagnosticLogs(context.Background()).DiagnosisGenerateAgentDiagnosticLogsRequest(diagnosisGenerateAgentDiagnosticLogsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiagnosisAPI.GenerateAgentDiagnosticLogs``: %v\n", err)
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


## GenerateCloudScannerDiagnosticLogs

> GenerateCloudScannerDiagnosticLogs(ctx).DiagnosisGenerateCloudScannerDiagnosticLogsRequest(diagnosisGenerateCloudScannerDiagnosticLogsRequest).Execute()

Generate Cloud Scanner Diagnostic Logs



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
	diagnosisGenerateCloudScannerDiagnosticLogsRequest := *openapiclient.NewDiagnosisGenerateCloudScannerDiagnosticLogsRequest([]openapiclient.DiagnosisNodeIdentifier{*openapiclient.NewDiagnosisNodeIdentifier("NodeId_example", "NodeType_example")}, int32(123)) // DiagnosisGenerateCloudScannerDiagnosticLogsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DiagnosisAPI.GenerateCloudScannerDiagnosticLogs(context.Background()).DiagnosisGenerateCloudScannerDiagnosticLogsRequest(diagnosisGenerateCloudScannerDiagnosticLogsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiagnosisAPI.GenerateCloudScannerDiagnosticLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateCloudScannerDiagnosticLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **diagnosisGenerateCloudScannerDiagnosticLogsRequest** | [**DiagnosisGenerateCloudScannerDiagnosticLogsRequest**](DiagnosisGenerateCloudScannerDiagnosticLogsRequest.md) |  | 

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
	r, err := apiClient.DiagnosisAPI.GenerateConsoleDiagnosticLogs(context.Background()).DiagnosisGenerateConsoleDiagnosticLogsRequest(diagnosisGenerateConsoleDiagnosticLogsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiagnosisAPI.GenerateConsoleDiagnosticLogs``: %v\n", err)
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
	resp, r, err := apiClient.DiagnosisAPI.GetDiagnosticLogs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiagnosisAPI.GetDiagnosticLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDiagnosticLogs`: DiagnosisGetDiagnosticLogsResponse
	fmt.Fprintf(os.Stdout, "Response from `DiagnosisAPI.GetDiagnosticLogs`: %v\n", resp)
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
	diagnosisDiagnosticLogsStatus := *openapiclient.NewDiagnosisDiagnosticLogsStatus("Status_example") // DiagnosisDiagnosticLogsStatus |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DiagnosisAPI.UpdateAgentDiagnosticLogsStatus(context.Background(), nodeId).DiagnosisDiagnosticLogsStatus(diagnosisDiagnosticLogsStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiagnosisAPI.UpdateAgentDiagnosticLogsStatus``: %v\n", err)
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


## UpdateCloudScannerDiagnosticLogsStatus

> UpdateCloudScannerDiagnosticLogsStatus(ctx, nodeId).DiagnosisDiagnosticLogsStatus(diagnosisDiagnosticLogsStatus).Execute()

Update Cloud Scanner Diagnostic Logs Status



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
	diagnosisDiagnosticLogsStatus := *openapiclient.NewDiagnosisDiagnosticLogsStatus("Status_example") // DiagnosisDiagnosticLogsStatus |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DiagnosisAPI.UpdateCloudScannerDiagnosticLogsStatus(context.Background(), nodeId).DiagnosisDiagnosticLogsStatus(diagnosisDiagnosticLogsStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiagnosisAPI.UpdateCloudScannerDiagnosticLogsStatus``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateCloudScannerDiagnosticLogsStatusRequest struct via the builder pattern


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

