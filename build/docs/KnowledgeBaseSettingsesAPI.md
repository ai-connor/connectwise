# \KnowledgeBaseSettingsesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceKnowledgebasesettings**](KnowledgeBaseSettingsesAPI.md#GetServiceKnowledgebasesettings) | **Get** /service/knowledgebasesettings | Get KnowledgeBaseSettings
[**GetServiceKnowledgebasesettingsById**](KnowledgeBaseSettingsesAPI.md#GetServiceKnowledgebasesettingsById) | **Get** /service/knowledgebasesettings/{id} | Get KnowledgeBaseSettings
[**PatchServiceKnowledgebasesettingsById**](KnowledgeBaseSettingsesAPI.md#PatchServiceKnowledgebasesettingsById) | **Patch** /service/knowledgebasesettings/{id} | Patch KnowledgeBaseSettings
[**PostServiceKnowledgebasesettings**](KnowledgeBaseSettingsesAPI.md#PostServiceKnowledgebasesettings) | **Post** /service/knowledgebasesettings | Post KnowledgeBaseSettings
[**PutServiceKnowledgebasesettingsById**](KnowledgeBaseSettingsesAPI.md#PutServiceKnowledgebasesettingsById) | **Put** /service/knowledgebasesettings/{id} | Put KnowledgeBaseSettings



## GetServiceKnowledgebasesettings

> KnowledgeBaseSettings GetServiceKnowledgebasesettings(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get KnowledgeBaseSettings

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ai-connor/connectwise"
)

func main() {
	clientId := "clientId_example" // string | 
	conditions := "conditions_example" // string |  (optional)
	childConditions := "childConditions_example" // string |  (optional)
	customFieldConditions := "customFieldConditions_example" // string |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	fields := "fields_example" // string |  (optional)
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	pageId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeBaseSettingsesAPI.GetServiceKnowledgebasesettings(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseSettingsesAPI.GetServiceKnowledgebasesettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceKnowledgebasesettings`: KnowledgeBaseSettings
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseSettingsesAPI.GetServiceKnowledgebasesettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceKnowledgebasesettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **conditions** | **string** |  | 
 **childConditions** | **string** |  | 
 **customFieldConditions** | **string** |  | 
 **orderBy** | **string** |  | 
 **fields** | **string** |  | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **pageId** | **int32** |  | 

### Return type

[**KnowledgeBaseSettings**](KnowledgeBaseSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceKnowledgebasesettingsById

> KnowledgeBaseSettings GetServiceKnowledgebasesettingsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get KnowledgeBaseSettings

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ai-connor/connectwise"
)

func main() {
	id := int32(56) // int32 | knowledgebasesettingId
	clientId := "clientId_example" // string | 
	conditions := "conditions_example" // string |  (optional)
	childConditions := "childConditions_example" // string |  (optional)
	customFieldConditions := "customFieldConditions_example" // string |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	fields := "fields_example" // string |  (optional)
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	pageId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeBaseSettingsesAPI.GetServiceKnowledgebasesettingsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseSettingsesAPI.GetServiceKnowledgebasesettingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceKnowledgebasesettingsById`: KnowledgeBaseSettings
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseSettingsesAPI.GetServiceKnowledgebasesettingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | knowledgebasesettingId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceKnowledgebasesettingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **conditions** | **string** |  | 
 **childConditions** | **string** |  | 
 **customFieldConditions** | **string** |  | 
 **orderBy** | **string** |  | 
 **fields** | **string** |  | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **pageId** | **int32** |  | 

### Return type

[**KnowledgeBaseSettings**](KnowledgeBaseSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchServiceKnowledgebasesettingsById

> KnowledgeBaseSettings PatchServiceKnowledgebasesettingsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch KnowledgeBaseSettings

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ai-connor/connectwise"
)

func main() {
	id := int32(56) // int32 | knowledgebasesettingId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeBaseSettingsesAPI.PatchServiceKnowledgebasesettingsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseSettingsesAPI.PatchServiceKnowledgebasesettingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchServiceKnowledgebasesettingsById`: KnowledgeBaseSettings
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseSettingsesAPI.PatchServiceKnowledgebasesettingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | knowledgebasesettingId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchServiceKnowledgebasesettingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**KnowledgeBaseSettings**](KnowledgeBaseSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServiceKnowledgebasesettings

> KnowledgeBaseSettings PostServiceKnowledgebasesettings(ctx).ClientId(clientId).KnowledgeBaseSettings(knowledgeBaseSettings).Execute()

Post KnowledgeBaseSettings

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ai-connor/connectwise"
)

func main() {
	clientId := "clientId_example" // string | 
	knowledgeBaseSettings := *openapiclient.NewKnowledgeBaseSettings(false) // KnowledgeBaseSettings | knowledgeBaseSettings

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeBaseSettingsesAPI.PostServiceKnowledgebasesettings(context.Background()).ClientId(clientId).KnowledgeBaseSettings(knowledgeBaseSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseSettingsesAPI.PostServiceKnowledgebasesettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServiceKnowledgebasesettings`: KnowledgeBaseSettings
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseSettingsesAPI.PostServiceKnowledgebasesettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostServiceKnowledgebasesettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **knowledgeBaseSettings** | [**KnowledgeBaseSettings**](KnowledgeBaseSettings.md) | knowledgeBaseSettings | 

### Return type

[**KnowledgeBaseSettings**](KnowledgeBaseSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServiceKnowledgebasesettingsById

> KnowledgeBaseSettings PutServiceKnowledgebasesettingsById(ctx, id).ClientId(clientId).KnowledgeBaseSettings(knowledgeBaseSettings).Execute()

Put KnowledgeBaseSettings

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ai-connor/connectwise"
)

func main() {
	id := int32(56) // int32 | knowledgebasesettingId
	clientId := "clientId_example" // string | 
	knowledgeBaseSettings := *openapiclient.NewKnowledgeBaseSettings(false) // KnowledgeBaseSettings | knowledgeBaseSettings

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeBaseSettingsesAPI.PutServiceKnowledgebasesettingsById(context.Background(), id).ClientId(clientId).KnowledgeBaseSettings(knowledgeBaseSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseSettingsesAPI.PutServiceKnowledgebasesettingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServiceKnowledgebasesettingsById`: KnowledgeBaseSettings
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseSettingsesAPI.PutServiceKnowledgebasesettingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | knowledgebasesettingId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServiceKnowledgebasesettingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **knowledgeBaseSettings** | [**KnowledgeBaseSettings**](KnowledgeBaseSettings.md) | knowledgeBaseSettings | 

### Return type

[**KnowledgeBaseSettings**](KnowledgeBaseSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

