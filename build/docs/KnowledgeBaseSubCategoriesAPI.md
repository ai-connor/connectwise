# \KnowledgeBaseSubCategoriesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServiceKnowledgeBaseSubCategoriesById**](KnowledgeBaseSubCategoriesAPI.md#DeleteServiceKnowledgeBaseSubCategoriesById) | **Delete** /service/knowledgeBaseSubCategories/{id} | Delete KnowledgeBaseSubCategory
[**GetServiceKnowledgeBaseSubCategories**](KnowledgeBaseSubCategoriesAPI.md#GetServiceKnowledgeBaseSubCategories) | **Get** /service/knowledgeBaseSubCategories | Get List of KnowledgeBaseSubCategory
[**GetServiceKnowledgeBaseSubCategoriesById**](KnowledgeBaseSubCategoriesAPI.md#GetServiceKnowledgeBaseSubCategoriesById) | **Get** /service/knowledgeBaseSubCategories/{id} | Get KnowledgeBaseSubCategory
[**GetServiceKnowledgeBaseSubCategoriesByIdUsages**](KnowledgeBaseSubCategoriesAPI.md#GetServiceKnowledgeBaseSubCategoriesByIdUsages) | **Get** /service/knowledgeBaseSubCategories/{id}/usages | Get List of Usage Count
[**GetServiceKnowledgeBaseSubCategoriesByIdUsagesList**](KnowledgeBaseSubCategoriesAPI.md#GetServiceKnowledgeBaseSubCategoriesByIdUsagesList) | **Get** /service/knowledgeBaseSubCategories/{id}/usages/list | Get List of Usage
[**GetServiceKnowledgeBaseSubCategoriesCount**](KnowledgeBaseSubCategoriesAPI.md#GetServiceKnowledgeBaseSubCategoriesCount) | **Get** /service/knowledgeBaseSubCategories/count | Get Count of KnowledgeBaseSubCategory
[**PatchServiceKnowledgeBaseSubCategoriesById**](KnowledgeBaseSubCategoriesAPI.md#PatchServiceKnowledgeBaseSubCategoriesById) | **Patch** /service/knowledgeBaseSubCategories/{id} | Patch KnowledgeBaseSubCategory
[**PostServiceKnowledgeBaseSubCategories**](KnowledgeBaseSubCategoriesAPI.md#PostServiceKnowledgeBaseSubCategories) | **Post** /service/knowledgeBaseSubCategories | Post KnowledgeBaseSubCategory
[**PutServiceKnowledgeBaseSubCategoriesById**](KnowledgeBaseSubCategoriesAPI.md#PutServiceKnowledgeBaseSubCategoriesById) | **Put** /service/knowledgeBaseSubCategories/{id} | Put KnowledgeBaseSubCategory



## DeleteServiceKnowledgeBaseSubCategoriesById

> DeleteServiceKnowledgeBaseSubCategoriesById(ctx, id).ClientId(clientId).Execute()

Delete KnowledgeBaseSubCategory

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
	id := int32(56) // int32 | knowledgeBaseSubCategoryId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KnowledgeBaseSubCategoriesAPI.DeleteServiceKnowledgeBaseSubCategoriesById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseSubCategoriesAPI.DeleteServiceKnowledgeBaseSubCategoriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | knowledgeBaseSubCategoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceKnowledgeBaseSubCategoriesByIdRequest struct via the builder pattern


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


## GetServiceKnowledgeBaseSubCategories

> []KnowledgeBaseSubCategory GetServiceKnowledgeBaseSubCategories(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of KnowledgeBaseSubCategory

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
	resp, r, err := apiClient.KnowledgeBaseSubCategoriesAPI.GetServiceKnowledgeBaseSubCategories(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseSubCategoriesAPI.GetServiceKnowledgeBaseSubCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceKnowledgeBaseSubCategories`: []KnowledgeBaseSubCategory
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseSubCategoriesAPI.GetServiceKnowledgeBaseSubCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceKnowledgeBaseSubCategoriesRequest struct via the builder pattern


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

[**[]KnowledgeBaseSubCategory**](KnowledgeBaseSubCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceKnowledgeBaseSubCategoriesById

> KnowledgeBaseSubCategory GetServiceKnowledgeBaseSubCategoriesById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get KnowledgeBaseSubCategory

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
	id := int32(56) // int32 | knowledgeBaseSubCategoryId
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
	resp, r, err := apiClient.KnowledgeBaseSubCategoriesAPI.GetServiceKnowledgeBaseSubCategoriesById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseSubCategoriesAPI.GetServiceKnowledgeBaseSubCategoriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceKnowledgeBaseSubCategoriesById`: KnowledgeBaseSubCategory
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseSubCategoriesAPI.GetServiceKnowledgeBaseSubCategoriesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | knowledgeBaseSubCategoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceKnowledgeBaseSubCategoriesByIdRequest struct via the builder pattern


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

[**KnowledgeBaseSubCategory**](KnowledgeBaseSubCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceKnowledgeBaseSubCategoriesByIdUsages

> []Usage GetServiceKnowledgeBaseSubCategoriesByIdUsages(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of Usage Count

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
	id := int32(56) // int32 | knowledgeBaseSubCategoryId
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
	resp, r, err := apiClient.KnowledgeBaseSubCategoriesAPI.GetServiceKnowledgeBaseSubCategoriesByIdUsages(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseSubCategoriesAPI.GetServiceKnowledgeBaseSubCategoriesByIdUsages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceKnowledgeBaseSubCategoriesByIdUsages`: []Usage
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseSubCategoriesAPI.GetServiceKnowledgeBaseSubCategoriesByIdUsages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | knowledgeBaseSubCategoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceKnowledgeBaseSubCategoriesByIdUsagesRequest struct via the builder pattern


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

[**[]Usage**](Usage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceKnowledgeBaseSubCategoriesByIdUsagesList

> []Usage GetServiceKnowledgeBaseSubCategoriesByIdUsagesList(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of Usage

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
	id := int32(56) // int32 | knowledgeBaseSubCategoryId
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
	resp, r, err := apiClient.KnowledgeBaseSubCategoriesAPI.GetServiceKnowledgeBaseSubCategoriesByIdUsagesList(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseSubCategoriesAPI.GetServiceKnowledgeBaseSubCategoriesByIdUsagesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceKnowledgeBaseSubCategoriesByIdUsagesList`: []Usage
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseSubCategoriesAPI.GetServiceKnowledgeBaseSubCategoriesByIdUsagesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | knowledgeBaseSubCategoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceKnowledgeBaseSubCategoriesByIdUsagesListRequest struct via the builder pattern


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

[**[]Usage**](Usage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceKnowledgeBaseSubCategoriesCount

> Count GetServiceKnowledgeBaseSubCategoriesCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of KnowledgeBaseSubCategory

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
	resp, r, err := apiClient.KnowledgeBaseSubCategoriesAPI.GetServiceKnowledgeBaseSubCategoriesCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseSubCategoriesAPI.GetServiceKnowledgeBaseSubCategoriesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceKnowledgeBaseSubCategoriesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseSubCategoriesAPI.GetServiceKnowledgeBaseSubCategoriesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceKnowledgeBaseSubCategoriesCountRequest struct via the builder pattern


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


## PatchServiceKnowledgeBaseSubCategoriesById

> KnowledgeBaseSubCategory PatchServiceKnowledgeBaseSubCategoriesById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch KnowledgeBaseSubCategory

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
	id := int32(56) // int32 | knowledgeBaseSubCategoryId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeBaseSubCategoriesAPI.PatchServiceKnowledgeBaseSubCategoriesById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseSubCategoriesAPI.PatchServiceKnowledgeBaseSubCategoriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchServiceKnowledgeBaseSubCategoriesById`: KnowledgeBaseSubCategory
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseSubCategoriesAPI.PatchServiceKnowledgeBaseSubCategoriesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | knowledgeBaseSubCategoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchServiceKnowledgeBaseSubCategoriesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**KnowledgeBaseSubCategory**](KnowledgeBaseSubCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServiceKnowledgeBaseSubCategories

> KnowledgeBaseSubCategory PostServiceKnowledgeBaseSubCategories(ctx).ClientId(clientId).KnowledgeBaseSubCategory(knowledgeBaseSubCategory).Execute()

Post KnowledgeBaseSubCategory

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
	knowledgeBaseSubCategory := *openapiclient.NewKnowledgeBaseSubCategory("Name_example", *openapiclient.NewKBCategoryReference()) // KnowledgeBaseSubCategory | knowledgeBaseSubCategory

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeBaseSubCategoriesAPI.PostServiceKnowledgeBaseSubCategories(context.Background()).ClientId(clientId).KnowledgeBaseSubCategory(knowledgeBaseSubCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseSubCategoriesAPI.PostServiceKnowledgeBaseSubCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServiceKnowledgeBaseSubCategories`: KnowledgeBaseSubCategory
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseSubCategoriesAPI.PostServiceKnowledgeBaseSubCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostServiceKnowledgeBaseSubCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **knowledgeBaseSubCategory** | [**KnowledgeBaseSubCategory**](KnowledgeBaseSubCategory.md) | knowledgeBaseSubCategory | 

### Return type

[**KnowledgeBaseSubCategory**](KnowledgeBaseSubCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServiceKnowledgeBaseSubCategoriesById

> KnowledgeBaseSubCategory PutServiceKnowledgeBaseSubCategoriesById(ctx, id).ClientId(clientId).KnowledgeBaseSubCategory(knowledgeBaseSubCategory).Execute()

Put KnowledgeBaseSubCategory

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
	id := int32(56) // int32 | knowledgeBaseSubCategoryId
	clientId := "clientId_example" // string | 
	knowledgeBaseSubCategory := *openapiclient.NewKnowledgeBaseSubCategory("Name_example", *openapiclient.NewKBCategoryReference()) // KnowledgeBaseSubCategory | knowledgeBaseSubCategory

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeBaseSubCategoriesAPI.PutServiceKnowledgeBaseSubCategoriesById(context.Background(), id).ClientId(clientId).KnowledgeBaseSubCategory(knowledgeBaseSubCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseSubCategoriesAPI.PutServiceKnowledgeBaseSubCategoriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServiceKnowledgeBaseSubCategoriesById`: KnowledgeBaseSubCategory
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseSubCategoriesAPI.PutServiceKnowledgeBaseSubCategoriesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | knowledgeBaseSubCategoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServiceKnowledgeBaseSubCategoriesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **knowledgeBaseSubCategory** | [**KnowledgeBaseSubCategory**](KnowledgeBaseSubCategory.md) | knowledgeBaseSubCategory | 

### Return type

[**KnowledgeBaseSubCategory**](KnowledgeBaseSubCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

