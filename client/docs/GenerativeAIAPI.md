# \GenerativeAIAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGenerativeAiIntegrationBedrock**](GenerativeAIAPI.md#AddGenerativeAiIntegrationBedrock) | **Post** /deepfence/generative-ai-integration/bedrock | Add AWS Bedrock Generative AI Integration
[**AddGenerativeAiIntegrationOpenAI**](GenerativeAIAPI.md#AddGenerativeAiIntegrationOpenAI) | **Post** /deepfence/generative-ai-integration/openai | Add OpenAI Generative AI Integration
[**DeleteGenerativeAiIntegration**](GenerativeAIAPI.md#DeleteGenerativeAiIntegration) | **Delete** /deepfence/generative-ai-integration/{integration_id} | Delete Generative AI Integration
[**GenerativeAiIntegrationCloudPostureQuery**](GenerativeAIAPI.md#GenerativeAiIntegrationCloudPostureQuery) | **Post** /deepfence/generative-ai-integration/query/cloud-posture | Send Cloud Posture query to Generative AI Integration
[**GenerativeAiIntegrationKubernetesPostureQuery**](GenerativeAIAPI.md#GenerativeAiIntegrationKubernetesPostureQuery) | **Post** /deepfence/generative-ai-integration/query/kubernetes-posture | Send Kubernetes Posture query to Generative AI Integration
[**GenerativeAiIntegrationLinuxPostureQuery**](GenerativeAIAPI.md#GenerativeAiIntegrationLinuxPostureQuery) | **Post** /deepfence/generative-ai-integration/query/linux-posture | Send Linux Posture query to Generative AI Integration
[**GenerativeAiIntegrationMalwareQuery**](GenerativeAIAPI.md#GenerativeAiIntegrationMalwareQuery) | **Post** /deepfence/generative-ai-integration/query/malware | Send Malware query to Generative AI Integration
[**GenerativeAiIntegrationSecretQuery**](GenerativeAIAPI.md#GenerativeAiIntegrationSecretQuery) | **Post** /deepfence/generative-ai-integration/query/secret | Send Secret query to Generative AI Integration
[**GenerativeAiIntegrationVulnerabilityQuery**](GenerativeAIAPI.md#GenerativeAiIntegrationVulnerabilityQuery) | **Post** /deepfence/generative-ai-integration/query/vulnerability | Send Vulnerability query to Generative AI Integration
[**ListGenerativeAiIntegration**](GenerativeAIAPI.md#ListGenerativeAiIntegration) | **Get** /deepfence/generative-ai-integration | List Generative AI Integrations
[**SetDefaultGenerativeAiIntegration**](GenerativeAIAPI.md#SetDefaultGenerativeAiIntegration) | **Put** /deepfence/generative-ai-integration/{integration_id}/default | Set Default Generative AI Integration



## AddGenerativeAiIntegrationBedrock

> ModelMessageResponse AddGenerativeAiIntegrationBedrock(ctx).ModelAddGenerativeAiBedrockIntegration(modelAddGenerativeAiBedrockIntegration).Execute()

Add AWS Bedrock Generative AI Integration



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
    modelAddGenerativeAiBedrockIntegration := *openapiclient.NewModelAddGenerativeAiBedrockIntegration("AwsRegion_example", "ModelId_example") // ModelAddGenerativeAiBedrockIntegration |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenerativeAIAPI.AddGenerativeAiIntegrationBedrock(context.Background()).ModelAddGenerativeAiBedrockIntegration(modelAddGenerativeAiBedrockIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenerativeAIAPI.AddGenerativeAiIntegrationBedrock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddGenerativeAiIntegrationBedrock`: ModelMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `GenerativeAIAPI.AddGenerativeAiIntegrationBedrock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddGenerativeAiIntegrationBedrockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelAddGenerativeAiBedrockIntegration** | [**ModelAddGenerativeAiBedrockIntegration**](ModelAddGenerativeAiBedrockIntegration.md) |  | 

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


## AddGenerativeAiIntegrationOpenAI

> ModelMessageResponse AddGenerativeAiIntegrationOpenAI(ctx).ModelAddGenerativeAiOpenAIIntegration(modelAddGenerativeAiOpenAIIntegration).Execute()

Add OpenAI Generative AI Integration



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
    modelAddGenerativeAiOpenAIIntegration := *openapiclient.NewModelAddGenerativeAiOpenAIIntegration("ApiKey_example", "ModelId_example") // ModelAddGenerativeAiOpenAIIntegration |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenerativeAIAPI.AddGenerativeAiIntegrationOpenAI(context.Background()).ModelAddGenerativeAiOpenAIIntegration(modelAddGenerativeAiOpenAIIntegration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenerativeAIAPI.AddGenerativeAiIntegrationOpenAI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddGenerativeAiIntegrationOpenAI`: ModelMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `GenerativeAIAPI.AddGenerativeAiIntegrationOpenAI`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddGenerativeAiIntegrationOpenAIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelAddGenerativeAiOpenAIIntegration** | [**ModelAddGenerativeAiOpenAIIntegration**](ModelAddGenerativeAiOpenAIIntegration.md) |  | 

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


## DeleteGenerativeAiIntegration

> DeleteGenerativeAiIntegration(ctx, integrationId).Execute()

Delete Generative AI Integration



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
    r, err := apiClient.GenerativeAIAPI.DeleteGenerativeAiIntegration(context.Background(), integrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenerativeAIAPI.DeleteGenerativeAiIntegration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteGenerativeAiIntegrationRequest struct via the builder pattern


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


## GenerativeAiIntegrationCloudPostureQuery

> string GenerativeAiIntegrationCloudPostureQuery(ctx).ModelGenerativeAiIntegrationCloudPostureRequest(modelGenerativeAiIntegrationCloudPostureRequest).Execute()

Send Cloud Posture query to Generative AI Integration



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
    modelGenerativeAiIntegrationCloudPostureRequest := *openapiclient.NewModelGenerativeAiIntegrationCloudPostureRequest("CloudProvider_example", "ComplianceCheckType_example", "QueryType_example", "RemediationFormat_example", "Title_example") // ModelGenerativeAiIntegrationCloudPostureRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenerativeAIAPI.GenerativeAiIntegrationCloudPostureQuery(context.Background()).ModelGenerativeAiIntegrationCloudPostureRequest(modelGenerativeAiIntegrationCloudPostureRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenerativeAIAPI.GenerativeAiIntegrationCloudPostureQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerativeAiIntegrationCloudPostureQuery`: string
    fmt.Fprintf(os.Stdout, "Response from `GenerativeAIAPI.GenerativeAiIntegrationCloudPostureQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerativeAiIntegrationCloudPostureQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelGenerativeAiIntegrationCloudPostureRequest** | [**ModelGenerativeAiIntegrationCloudPostureRequest**](ModelGenerativeAiIntegrationCloudPostureRequest.md) |  | 

### Return type

**string**

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerativeAiIntegrationKubernetesPostureQuery

> string GenerativeAiIntegrationKubernetesPostureQuery(ctx).ModelGenerativeAiIntegrationKubernetesPostureRequest(modelGenerativeAiIntegrationKubernetesPostureRequest).Execute()

Send Kubernetes Posture query to Generative AI Integration



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
    modelGenerativeAiIntegrationKubernetesPostureRequest := *openapiclient.NewModelGenerativeAiIntegrationKubernetesPostureRequest("ComplianceCheckType_example", "Description_example", "QueryType_example", "RemediationFormat_example") // ModelGenerativeAiIntegrationKubernetesPostureRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenerativeAIAPI.GenerativeAiIntegrationKubernetesPostureQuery(context.Background()).ModelGenerativeAiIntegrationKubernetesPostureRequest(modelGenerativeAiIntegrationKubernetesPostureRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenerativeAIAPI.GenerativeAiIntegrationKubernetesPostureQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerativeAiIntegrationKubernetesPostureQuery`: string
    fmt.Fprintf(os.Stdout, "Response from `GenerativeAIAPI.GenerativeAiIntegrationKubernetesPostureQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerativeAiIntegrationKubernetesPostureQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelGenerativeAiIntegrationKubernetesPostureRequest** | [**ModelGenerativeAiIntegrationKubernetesPostureRequest**](ModelGenerativeAiIntegrationKubernetesPostureRequest.md) |  | 

### Return type

**string**

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerativeAiIntegrationLinuxPostureQuery

> string GenerativeAiIntegrationLinuxPostureQuery(ctx).ModelGenerativeAiIntegrationLinuxPostureRequest(modelGenerativeAiIntegrationLinuxPostureRequest).Execute()

Send Linux Posture query to Generative AI Integration



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
    modelGenerativeAiIntegrationLinuxPostureRequest := *openapiclient.NewModelGenerativeAiIntegrationLinuxPostureRequest("ComplianceCheckType_example", "Description_example", "QueryType_example", "RemediationFormat_example", "TestNumber_example") // ModelGenerativeAiIntegrationLinuxPostureRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenerativeAIAPI.GenerativeAiIntegrationLinuxPostureQuery(context.Background()).ModelGenerativeAiIntegrationLinuxPostureRequest(modelGenerativeAiIntegrationLinuxPostureRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenerativeAIAPI.GenerativeAiIntegrationLinuxPostureQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerativeAiIntegrationLinuxPostureQuery`: string
    fmt.Fprintf(os.Stdout, "Response from `GenerativeAIAPI.GenerativeAiIntegrationLinuxPostureQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerativeAiIntegrationLinuxPostureQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelGenerativeAiIntegrationLinuxPostureRequest** | [**ModelGenerativeAiIntegrationLinuxPostureRequest**](ModelGenerativeAiIntegrationLinuxPostureRequest.md) |  | 

### Return type

**string**

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerativeAiIntegrationMalwareQuery

> string GenerativeAiIntegrationMalwareQuery(ctx).ModelGenerativeAiIntegrationMalwareRequest(modelGenerativeAiIntegrationMalwareRequest).Execute()

Send Malware query to Generative AI Integration



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
    modelGenerativeAiIntegrationMalwareRequest := *openapiclient.NewModelGenerativeAiIntegrationMalwareRequest("Info_example", "QueryType_example", "RuleName_example") // ModelGenerativeAiIntegrationMalwareRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenerativeAIAPI.GenerativeAiIntegrationMalwareQuery(context.Background()).ModelGenerativeAiIntegrationMalwareRequest(modelGenerativeAiIntegrationMalwareRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenerativeAIAPI.GenerativeAiIntegrationMalwareQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerativeAiIntegrationMalwareQuery`: string
    fmt.Fprintf(os.Stdout, "Response from `GenerativeAIAPI.GenerativeAiIntegrationMalwareQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerativeAiIntegrationMalwareQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelGenerativeAiIntegrationMalwareRequest** | [**ModelGenerativeAiIntegrationMalwareRequest**](ModelGenerativeAiIntegrationMalwareRequest.md) |  | 

### Return type

**string**

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerativeAiIntegrationSecretQuery

> string GenerativeAiIntegrationSecretQuery(ctx).ModelGenerativeAiIntegrationSecretRequest(modelGenerativeAiIntegrationSecretRequest).Execute()

Send Secret query to Generative AI Integration



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
    modelGenerativeAiIntegrationSecretRequest := *openapiclient.NewModelGenerativeAiIntegrationSecretRequest("Name_example", "QueryType_example") // ModelGenerativeAiIntegrationSecretRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenerativeAIAPI.GenerativeAiIntegrationSecretQuery(context.Background()).ModelGenerativeAiIntegrationSecretRequest(modelGenerativeAiIntegrationSecretRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenerativeAIAPI.GenerativeAiIntegrationSecretQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerativeAiIntegrationSecretQuery`: string
    fmt.Fprintf(os.Stdout, "Response from `GenerativeAIAPI.GenerativeAiIntegrationSecretQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerativeAiIntegrationSecretQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelGenerativeAiIntegrationSecretRequest** | [**ModelGenerativeAiIntegrationSecretRequest**](ModelGenerativeAiIntegrationSecretRequest.md) |  | 

### Return type

**string**

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerativeAiIntegrationVulnerabilityQuery

> string GenerativeAiIntegrationVulnerabilityQuery(ctx).ModelGenerativeAiIntegrationVulnerabilityRequest(modelGenerativeAiIntegrationVulnerabilityRequest).Execute()

Send Vulnerability query to Generative AI Integration



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
    modelGenerativeAiIntegrationVulnerabilityRequest := *openapiclient.NewModelGenerativeAiIntegrationVulnerabilityRequest("CveCausedByPackage_example", "CveId_example", "CveType_example", "QueryType_example", "RemediationFormat_example") // ModelGenerativeAiIntegrationVulnerabilityRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenerativeAIAPI.GenerativeAiIntegrationVulnerabilityQuery(context.Background()).ModelGenerativeAiIntegrationVulnerabilityRequest(modelGenerativeAiIntegrationVulnerabilityRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenerativeAIAPI.GenerativeAiIntegrationVulnerabilityQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerativeAiIntegrationVulnerabilityQuery`: string
    fmt.Fprintf(os.Stdout, "Response from `GenerativeAIAPI.GenerativeAiIntegrationVulnerabilityQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerativeAiIntegrationVulnerabilityQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelGenerativeAiIntegrationVulnerabilityRequest** | [**ModelGenerativeAiIntegrationVulnerabilityRequest**](ModelGenerativeAiIntegrationVulnerabilityRequest.md) |  | 

### Return type

**string**

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGenerativeAiIntegration

> []ModelGenerativeAiIntegrationListResponse ListGenerativeAiIntegration(ctx).IntegrationType(integrationType).Execute()

List Generative AI Integrations



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
    integrationType := "integrationType_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenerativeAIAPI.ListGenerativeAiIntegration(context.Background()).IntegrationType(integrationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenerativeAIAPI.ListGenerativeAiIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGenerativeAiIntegration`: []ModelGenerativeAiIntegrationListResponse
    fmt.Fprintf(os.Stdout, "Response from `GenerativeAIAPI.ListGenerativeAiIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGenerativeAiIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integrationType** | **string** |  | 

### Return type

[**[]ModelGenerativeAiIntegrationListResponse**](ModelGenerativeAiIntegrationListResponse.md)

### Authorization

[bearer_token](../README.md#bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDefaultGenerativeAiIntegration

> SetDefaultGenerativeAiIntegration(ctx, integrationId).Execute()

Set Default Generative AI Integration



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
    r, err := apiClient.GenerativeAIAPI.SetDefaultGenerativeAiIntegration(context.Background(), integrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenerativeAIAPI.SetDefaultGenerativeAiIntegration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetDefaultGenerativeAiIntegrationRequest struct via the builder pattern


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

