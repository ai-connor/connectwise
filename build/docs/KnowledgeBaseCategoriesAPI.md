# \KnowledgeBaseCategoriesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServiceKnowledgeBaseCategoriesById**](KnowledgeBaseCategoriesAPI.md#DeleteServiceKnowledgeBaseCategoriesById) | **Delete** /service/knowledgeBaseCategories/{id} | Delete KnowledgeBaseCategory
[**GetServiceKnowledgeBaseCategories**](KnowledgeBaseCategoriesAPI.md#GetServiceKnowledgeBaseCategories) | **Get** /service/knowledgeBaseCategories | Get List of KnowledgeBaseCategory
[**GetServiceKnowledgeBaseCategoriesById**](KnowledgeBaseCategoriesAPI.md#GetServiceKnowledgeBaseCategoriesById) | **Get** /service/knowledgeBaseCategories/{id} | Get KnowledgeBaseCategory
[**GetServiceKnowledgeBaseCategoriesCount**](KnowledgeBaseCategoriesAPI.md#GetServiceKnowledgeBaseCategoriesCount) | **Get** /service/knowledgeBaseCategories/count | Get Count of KnowledgeBaseCategory
[**PatchServiceKnowledgeBaseCategoriesById**](KnowledgeBaseCategoriesAPI.md#PatchServiceKnowledgeBaseCategoriesById) | **Patch** /service/knowledgeBaseCategories/{id} | Patch KnowledgeBaseCategory
[**PostServiceKnowledgeBaseCategories**](KnowledgeBaseCategoriesAPI.md#PostServiceKnowledgeBaseCategories) | **Post** /service/knowledgeBaseCategories | Post KnowledgeBaseCategory
[**PutServiceKnowledgeBaseCategoriesById**](KnowledgeBaseCategoriesAPI.md#PutServiceKnowledgeBaseCategoriesById) | **Put** /service/knowledgeBaseCategories/{id} | Put KnowledgeBaseCategory



## DeleteServiceKnowledgeBaseCategoriesById

> DeleteServiceKnowledgeBaseCategoriesById(ctx, id).ClientId(clientId).Execute()

Delete KnowledgeBaseCategory

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
	id := int32(56) // int32 | knowledgeBaseCategoryId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KnowledgeBaseCategoriesAPI.DeleteServiceKnowledgeBaseCategoriesById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseCategoriesAPI.DeleteServiceKnowledgeBaseCategoriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | knowledgeBaseCategoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceKnowledgeBaseCategoriesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceKnowledgeBaseCategories

> []KnowledgeBaseCategory GetServiceKnowledgeBaseCategories(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of KnowledgeBaseCategory

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
	resp, r, err := apiClient.KnowledgeBaseCategoriesAPI.GetServiceKnowledgeBaseCategories(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseCategoriesAPI.GetServiceKnowledgeBaseCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceKnowledgeBaseCategories`: []KnowledgeBaseCategory
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseCategoriesAPI.GetServiceKnowledgeBaseCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceKnowledgeBaseCategoriesRequest struct via the builder pattern


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

[**[]KnowledgeBaseCategory**](KnowledgeBaseCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceKnowledgeBaseCategoriesById

> KnowledgeBaseCategory GetServiceKnowledgeBaseCategoriesById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get KnowledgeBaseCategory

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
	id := int32(56) // int32 | knowledgeBaseCategoryId
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
	resp, r, err := apiClient.KnowledgeBaseCategoriesAPI.GetServiceKnowledgeBaseCategoriesById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseCategoriesAPI.GetServiceKnowledgeBaseCategoriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceKnowledgeBaseCategoriesById`: KnowledgeBaseCategory
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseCategoriesAPI.GetServiceKnowledgeBaseCategoriesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | knowledgeBaseCategoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceKnowledgeBaseCategoriesByIdRequest struct via the builder pattern


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

[**KnowledgeBaseCategory**](KnowledgeBaseCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceKnowledgeBaseCategoriesCount

> Count GetServiceKnowledgeBaseCategoriesCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of KnowledgeBaseCategory

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
	resp, r, err := apiClient.KnowledgeBaseCategoriesAPI.GetServiceKnowledgeBaseCategoriesCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseCategoriesAPI.GetServiceKnowledgeBaseCategoriesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceKnowledgeBaseCategoriesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseCategoriesAPI.GetServiceKnowledgeBaseCategoriesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceKnowledgeBaseCategoriesCountRequest struct via the builder pattern


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

[**Count**](Count.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchServiceKnowledgeBaseCategoriesById

> KnowledgeBaseCategory PatchServiceKnowledgeBaseCategoriesById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch KnowledgeBaseCategory

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
	id := int32(56) // int32 | knowledgeBaseCategoryId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeBaseCategoriesAPI.PatchServiceKnowledgeBaseCategoriesById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseCategoriesAPI.PatchServiceKnowledgeBaseCategoriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchServiceKnowledgeBaseCategoriesById`: KnowledgeBaseCategory
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseCategoriesAPI.PatchServiceKnowledgeBaseCategoriesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | knowledgeBaseCategoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchServiceKnowledgeBaseCategoriesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**KnowledgeBaseCategory**](KnowledgeBaseCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServiceKnowledgeBaseCategories

> KnowledgeBaseCategory PostServiceKnowledgeBaseCategories(ctx).ClientId(clientId).KnowledgeBaseCategory(knowledgeBaseCategory).Execute()

Post KnowledgeBaseCategory

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
	knowledgeBaseCategory := *openapiclient.NewKnowledgeBaseCategory("Name_example") // KnowledgeBaseCategory | knowledgeBaseCategory

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeBaseCategoriesAPI.PostServiceKnowledgeBaseCategories(context.Background()).ClientId(clientId).KnowledgeBaseCategory(knowledgeBaseCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseCategoriesAPI.PostServiceKnowledgeBaseCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServiceKnowledgeBaseCategories`: KnowledgeBaseCategory
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseCategoriesAPI.PostServiceKnowledgeBaseCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostServiceKnowledgeBaseCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **knowledgeBaseCategory** | [**KnowledgeBaseCategory**](KnowledgeBaseCategory.md) | knowledgeBaseCategory | 

### Return type

[**KnowledgeBaseCategory**](KnowledgeBaseCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServiceKnowledgeBaseCategoriesById

> KnowledgeBaseCategory PutServiceKnowledgeBaseCategoriesById(ctx, id).ClientId(clientId).KnowledgeBaseCategory(knowledgeBaseCategory).Execute()

Put KnowledgeBaseCategory

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
	id := int32(56) // int32 | knowledgeBaseCategoryId
	clientId := "clientId_example" // string | 
	knowledgeBaseCategory := *openapiclient.NewKnowledgeBaseCategory("Name_example") // KnowledgeBaseCategory | knowledgeBaseCategory

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeBaseCategoriesAPI.PutServiceKnowledgeBaseCategoriesById(context.Background(), id).ClientId(clientId).KnowledgeBaseCategory(knowledgeBaseCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseCategoriesAPI.PutServiceKnowledgeBaseCategoriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServiceKnowledgeBaseCategoriesById`: KnowledgeBaseCategory
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseCategoriesAPI.PutServiceKnowledgeBaseCategoriesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | knowledgeBaseCategoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServiceKnowledgeBaseCategoriesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **knowledgeBaseCategory** | [**KnowledgeBaseCategory**](KnowledgeBaseCategory.md) | knowledgeBaseCategory | 

### Return type

[**KnowledgeBaseCategory**](KnowledgeBaseCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

