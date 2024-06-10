# \SettingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEmailConfiguration**](SettingsAPI.md#AddEmailConfiguration) | **Post** /deepfence/settings/email | Add Email Configuration
[**AddScheduledTask**](SettingsAPI.md#AddScheduledTask) | **Post** /deepfence/scheduled-task | Add scheduled task
[**DeleteCustomScheduledTask**](SettingsAPI.md#DeleteCustomScheduledTask) | **Delete** /deepfence/scheduled-task/{id} | Delete Custom Schedule task
[**DeleteEmailConfiguration**](SettingsAPI.md#DeleteEmailConfiguration) | **Delete** /deepfence/settings/email/{config_id} | Delete Email Configurations
[**DeleteLicense**](SettingsAPI.md#DeleteLicense) | **Delete** /deepfence/license | Delete License
[**GenerateLicense**](SettingsAPI.md#GenerateLicense) | **Post** /deepfence/license/generate | Generate License Key
[**GetAgentBinaryDownloadURL**](SettingsAPI.md#GetAgentBinaryDownloadURL) | **Get** /deepfence/agent-deployment/binary/download-url | Get agent binary download url
[**GetAgentVersions**](SettingsAPI.md#GetAgentVersions) | **Get** /deepfence/settings/agent/versions | Get available agent versions
[**GetEmailConfiguration**](SettingsAPI.md#GetEmailConfiguration) | **Get** /deepfence/settings/email | Get Email Configurations
[**GetLicense**](SettingsAPI.md#GetLicense) | **Get** /deepfence/license | Get License Details
[**GetScheduledTasks**](SettingsAPI.md#GetScheduledTasks) | **Get** /deepfence/scheduled-task | Get scheduled tasks
[**GetSettings**](SettingsAPI.md#GetSettings) | **Get** /deepfence/settings/global-settings | Get settings
[**GetUserAuditLogs**](SettingsAPI.md#GetUserAuditLogs) | **Post** /deepfence/settings/user-audit-log | Get user audit logs
[**GetUserAuditLogsCount**](SettingsAPI.md#GetUserAuditLogsCount) | **Get** /deepfence/settings/user-audit-log/count | Get user audit logs count
[**RegisterLicense**](SettingsAPI.md#RegisterLicense) | **Post** /deepfence/license | Register License
[**TestConfiguredEmail**](SettingsAPI.md#TestConfiguredEmail) | **Post** /deepfence/settings/email/test | Test Configured Email
[**TestUnconfiguredEmail**](SettingsAPI.md#TestUnconfiguredEmail) | **Post** /deepfence/settings/email/test-unconfigured | Test Unconfigured Email
[**UpdateScheduledTask**](SettingsAPI.md#UpdateScheduledTask) | **Patch** /deepfence/scheduled-task/{id} | Update scheduled task
[**UpdateSetting**](SettingsAPI.md#UpdateSetting) | **Patch** /deepfence/settings/global-settings/{id} | Update setting
[**UploadAgentVersion**](SettingsAPI.md#UploadAgentVersion) | **Put** /deepfence/settings/agent/version | Upload New agent version
[**UploadMalwareRules**](SettingsAPI.md#UploadMalwareRules) | **Put** /deepfence/database/malware | Upload Malware Rules
[**UploadPostureControls**](SettingsAPI.md#UploadPostureControls) | **Put** /deepfence/database/posture | Upload Posture Controls
[**UploadSecretsRules**](SettingsAPI.md#UploadSecretsRules) | **Put** /deepfence/database/secret | Upload Secrets Rules
[**UploadVulnerabilityDatabase**](SettingsAPI.md#UploadVulnerabilityDatabase) | **Put** /deepfence/database/vulnerability | Upload Vulnerability Database



## AddEmailConfiguration

> ModelMessageResponse AddEmailConfiguration(ctx).ModelEmailConfigurationAdd(modelEmailConfigurationAdd).Execute()

Add Email Configuration



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
	modelEmailConfigurationAdd := *openapiclient.NewModelEmailConfigurationAdd() // ModelEmailConfigurationAdd |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.AddEmailConfiguration(context.Background()).ModelEmailConfigurationAdd(modelEmailConfigurationAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.AddEmailConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddEmailConfiguration`: ModelMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.AddEmailConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddEmailConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelEmailConfigurationAdd** | [**ModelEmailConfigurationAdd**](ModelEmailConfigurationAdd.md) |  | 

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


## AddScheduledTask

> AddScheduledTask(ctx).ModelAddScheduledTaskRequest(modelAddScheduledTaskRequest).Execute()

Add scheduled task



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
	modelAddScheduledTaskRequest := *openapiclient.NewModelAddScheduledTaskRequest("Action_example", []openapiclient.ModelBenchmarkType{openapiclient.ModelBenchmarkType("hipaa")}, *openapiclient.NewModelScanFilter(*openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}}), *openapiclient.NewReportersContainsFilter(map[string][]interface{}{"key": []interface{}{nil}})), []openapiclient.ModelNodeIdentifier{*openapiclient.NewModelNodeIdentifier("NodeId_example", "NodeType_example")}, []openapiclient.ModelVulnerabilityScanConfigLanguage{*openapiclient.NewModelVulnerabilityScanConfigLanguage("Language_example")}) // ModelAddScheduledTaskRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.AddScheduledTask(context.Background()).ModelAddScheduledTaskRequest(modelAddScheduledTaskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.AddScheduledTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddScheduledTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelAddScheduledTaskRequest** | [**ModelAddScheduledTaskRequest**](ModelAddScheduledTaskRequest.md) |  | 

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


## DeleteCustomScheduledTask

> DeleteCustomScheduledTask(ctx, id).Execute()

Delete Custom Schedule task



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
	r, err := apiClient.SettingsAPI.DeleteCustomScheduledTask(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.DeleteCustomScheduledTask``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCustomScheduledTaskRequest struct via the builder pattern


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


## DeleteEmailConfiguration

> DeleteEmailConfiguration(ctx, configId).Execute()

Delete Email Configurations



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
	configId := "configId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.DeleteEmailConfiguration(context.Background(), configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.DeleteEmailConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmailConfigurationRequest struct via the builder pattern


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


## DeleteLicense

> DeleteLicense(ctx).Execute()

Delete License



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
	r, err := apiClient.SettingsAPI.DeleteLicense(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.DeleteLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLicenseRequest struct via the builder pattern


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


## GenerateLicense

> ModelGenerateLicenseResponse GenerateLicense(ctx).ModelGenerateLicenseRequest(modelGenerateLicenseRequest).Execute()

Generate License Key



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
	modelGenerateLicenseRequest := *openapiclient.NewModelGenerateLicenseRequest("Company_example", "Email_example", "FirstName_example", "LastName_example", false) // ModelGenerateLicenseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.GenerateLicense(context.Background()).ModelGenerateLicenseRequest(modelGenerateLicenseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GenerateLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateLicense`: ModelGenerateLicenseResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GenerateLicense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelGenerateLicenseRequest** | [**ModelGenerateLicenseRequest**](ModelGenerateLicenseRequest.md) |  | 

### Return type

[**ModelGenerateLicenseResponse**](ModelGenerateLicenseResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentBinaryDownloadURL

> ModelGetAgentBinaryDownloadURLResponse GetAgentBinaryDownloadURL(ctx).Execute()

Get agent binary download url



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
	resp, r, err := apiClient.SettingsAPI.GetAgentBinaryDownloadURL(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetAgentBinaryDownloadURL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentBinaryDownloadURL`: ModelGetAgentBinaryDownloadURLResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetAgentBinaryDownloadURL`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentBinaryDownloadURLRequest struct via the builder pattern


### Return type

[**ModelGetAgentBinaryDownloadURLResponse**](ModelGetAgentBinaryDownloadURLResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentVersions

> ModelListAgentVersionResp GetAgentVersions(ctx).Execute()

Get available agent versions



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
	resp, r, err := apiClient.SettingsAPI.GetAgentVersions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetAgentVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentVersions`: ModelListAgentVersionResp
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetAgentVersions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentVersionsRequest struct via the builder pattern


### Return type

[**ModelListAgentVersionResp**](ModelListAgentVersionResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailConfiguration

> []ModelEmailConfigurationResp GetEmailConfiguration(ctx).Execute()

Get Email Configurations



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
	resp, r, err := apiClient.SettingsAPI.GetEmailConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetEmailConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmailConfiguration`: []ModelEmailConfigurationResp
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetEmailConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailConfigurationRequest struct via the builder pattern


### Return type

[**[]ModelEmailConfigurationResp**](ModelEmailConfigurationResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicense

> ModelLicense GetLicense(ctx).Execute()

Get License Details



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
	resp, r, err := apiClient.SettingsAPI.GetLicense(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLicense`: ModelLicense
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetLicense`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseRequest struct via the builder pattern


### Return type

[**ModelLicense**](ModelLicense.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScheduledTasks

> []PostgresqlDbScheduler GetScheduledTasks(ctx).Execute()

Get scheduled tasks



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
	resp, r, err := apiClient.SettingsAPI.GetScheduledTasks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetScheduledTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScheduledTasks`: []PostgresqlDbScheduler
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetScheduledTasks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduledTasksRequest struct via the builder pattern


### Return type

[**[]PostgresqlDbScheduler**](PostgresqlDbScheduler.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettings

> []ModelSettingsResponse GetSettings(ctx).Execute()

Get settings



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
	resp, r, err := apiClient.SettingsAPI.GetSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettings`: []ModelSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsRequest struct via the builder pattern


### Return type

[**[]ModelSettingsResponse**](ModelSettingsResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserAuditLogs

> []PostgresqlDbGetAuditLogsRow GetUserAuditLogs(ctx).ModelGetAuditLogsRequest(modelGetAuditLogsRequest).Execute()

Get user audit logs



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
	modelGetAuditLogsRequest := *openapiclient.NewModelGetAuditLogsRequest(*openapiclient.NewModelFetchWindow(int32(123), int32(123))) // ModelGetAuditLogsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.GetUserAuditLogs(context.Background()).ModelGetAuditLogsRequest(modelGetAuditLogsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetUserAuditLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserAuditLogs`: []PostgresqlDbGetAuditLogsRow
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetUserAuditLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelGetAuditLogsRequest** | [**ModelGetAuditLogsRequest**](ModelGetAuditLogsRequest.md) |  | 

### Return type

[**[]PostgresqlDbGetAuditLogsRow**](PostgresqlDbGetAuditLogsRow.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserAuditLogsCount

> SearchSearchCountResp GetUserAuditLogsCount(ctx).Execute()

Get user audit logs count



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
	resp, r, err := apiClient.SettingsAPI.GetUserAuditLogsCount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetUserAuditLogsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserAuditLogsCount`: SearchSearchCountResp
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetUserAuditLogsCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAuditLogsCountRequest struct via the builder pattern


### Return type

[**SearchSearchCountResp**](SearchSearchCountResp.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterLicense

> ModelRegisterLicenseResponse RegisterLicense(ctx).ModelRegisterLicenseRequest(modelRegisterLicenseRequest).Execute()

Register License



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
	modelRegisterLicenseRequest := *openapiclient.NewModelRegisterLicenseRequest("LicenseKey_example") // ModelRegisterLicenseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.RegisterLicense(context.Background()).ModelRegisterLicenseRequest(modelRegisterLicenseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.RegisterLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterLicense`: ModelRegisterLicenseResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.RegisterLicense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelRegisterLicenseRequest** | [**ModelRegisterLicenseRequest**](ModelRegisterLicenseRequest.md) |  | 

### Return type

[**ModelRegisterLicenseResponse**](ModelRegisterLicenseResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestConfiguredEmail

> ModelMessageResponse TestConfiguredEmail(ctx).Execute()

Test Configured Email



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
	resp, r, err := apiClient.SettingsAPI.TestConfiguredEmail(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.TestConfiguredEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestConfiguredEmail`: ModelMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.TestConfiguredEmail`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestConfiguredEmailRequest struct via the builder pattern


### Return type

[**ModelMessageResponse**](ModelMessageResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestUnconfiguredEmail

> ModelMessageResponse TestUnconfiguredEmail(ctx).ModelEmailConfigurationAdd(modelEmailConfigurationAdd).Execute()

Test Unconfigured Email



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
	modelEmailConfigurationAdd := *openapiclient.NewModelEmailConfigurationAdd() // ModelEmailConfigurationAdd |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.TestUnconfiguredEmail(context.Background()).ModelEmailConfigurationAdd(modelEmailConfigurationAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.TestUnconfiguredEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestUnconfiguredEmail`: ModelMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.TestUnconfiguredEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestUnconfiguredEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelEmailConfigurationAdd** | [**ModelEmailConfigurationAdd**](ModelEmailConfigurationAdd.md) |  | 

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


## UpdateScheduledTask

> UpdateScheduledTask(ctx, id).ModelUpdateScheduledTaskRequest(modelUpdateScheduledTaskRequest).Execute()

Update scheduled task



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
	modelUpdateScheduledTaskRequest := *openapiclient.NewModelUpdateScheduledTaskRequest(false) // ModelUpdateScheduledTaskRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.UpdateScheduledTask(context.Background(), id).ModelUpdateScheduledTaskRequest(modelUpdateScheduledTaskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.UpdateScheduledTask``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateScheduledTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelUpdateScheduledTaskRequest** | [**ModelUpdateScheduledTaskRequest**](ModelUpdateScheduledTaskRequest.md) |  | 

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


## UpdateSetting

> UpdateSetting(ctx, id).ModelSettingUpdateRequest(modelSettingUpdateRequest).Execute()

Update setting



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
	modelSettingUpdateRequest := *openapiclient.NewModelSettingUpdateRequest("Key_example", "Value_example") // ModelSettingUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.UpdateSetting(context.Background(), id).ModelSettingUpdateRequest(modelSettingUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.UpdateSetting``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelSettingUpdateRequest** | [**ModelSettingUpdateRequest**](ModelSettingUpdateRequest.md) |  | 

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


## UploadAgentVersion

> UploadAgentVersion(ctx).Tarball(tarball).Execute()

Upload New agent version



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
	tarball := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.UploadAgentVersion(context.Background()).Tarball(tarball).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.UploadAgentVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadAgentVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tarball** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadMalwareRules

> ModelMessageResponse UploadMalwareRules(ctx).Database(database).Execute()

Upload Malware Rules



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
	database := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.UploadMalwareRules(context.Background()).Database(database).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.UploadMalwareRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadMalwareRules`: ModelMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.UploadMalwareRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadMalwareRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **database** | ***os.File** |  | 

### Return type

[**ModelMessageResponse**](ModelMessageResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadPostureControls

> ModelMessageResponse UploadPostureControls(ctx).Database(database).Execute()

Upload Posture Controls



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
	database := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.UploadPostureControls(context.Background()).Database(database).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.UploadPostureControls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadPostureControls`: ModelMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.UploadPostureControls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadPostureControlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **database** | ***os.File** |  | 

### Return type

[**ModelMessageResponse**](ModelMessageResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadSecretsRules

> ModelMessageResponse UploadSecretsRules(ctx).Database(database).Execute()

Upload Secrets Rules



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
	database := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.UploadSecretsRules(context.Background()).Database(database).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.UploadSecretsRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadSecretsRules`: ModelMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.UploadSecretsRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadSecretsRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **database** | ***os.File** |  | 

### Return type

[**ModelMessageResponse**](ModelMessageResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadVulnerabilityDatabase

> ModelMessageResponse UploadVulnerabilityDatabase(ctx).Database(database).Execute()

Upload Vulnerability Database



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
	database := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.UploadVulnerabilityDatabase(context.Background()).Database(database).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.UploadVulnerabilityDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadVulnerabilityDatabase`: ModelMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.UploadVulnerabilityDatabase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadVulnerabilityDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **database** | ***os.File** |  | 

### Return type

[**ModelMessageResponse**](ModelMessageResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

