# \TodayPageCategoriesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemTodayPageCategoriesById**](TodayPageCategoriesAPI.md#DeleteSystemTodayPageCategoriesById) | **Delete** /system/todayPageCategories/{id} | Delete TodayPageCategory
[**GetSystemTodayPageCategories**](TodayPageCategoriesAPI.md#GetSystemTodayPageCategories) | **Get** /system/todayPageCategories | Get List of TodayPageCategory
[**GetSystemTodayPageCategoriesById**](TodayPageCategoriesAPI.md#GetSystemTodayPageCategoriesById) | **Get** /system/todayPageCategories/{id} | Get TodayPageCategory
[**GetSystemTodayPageCategoriesCount**](TodayPageCategoriesAPI.md#GetSystemTodayPageCategoriesCount) | **Get** /system/todayPageCategories/count | Get Count of TodayPageCategory
[**PatchSystemTodayPageCategoriesById**](TodayPageCategoriesAPI.md#PatchSystemTodayPageCategoriesById) | **Patch** /system/todayPageCategories/{id} | Patch TodayPageCategory
[**PostSystemTodayPageCategories**](TodayPageCategoriesAPI.md#PostSystemTodayPageCategories) | **Post** /system/todayPageCategories | Post TodayPageCategory
[**PutSystemTodayPageCategoriesById**](TodayPageCategoriesAPI.md#PutSystemTodayPageCategoriesById) | **Put** /system/todayPageCategories/{id} | Put TodayPageCategory



## DeleteSystemTodayPageCategoriesById

> DeleteSystemTodayPageCategoriesById(ctx, id).ClientId(clientId).Execute()

Delete TodayPageCategory

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
	id := int32(56) // int32 | todayPageCategoryId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TodayPageCategoriesAPI.DeleteSystemTodayPageCategoriesById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodayPageCategoriesAPI.DeleteSystemTodayPageCategoriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | todayPageCategoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemTodayPageCategoriesByIdRequest struct via the builder pattern


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


## GetSystemTodayPageCategories

> []TodayPageCategory GetSystemTodayPageCategories(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of TodayPageCategory

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
	resp, r, err := apiClient.TodayPageCategoriesAPI.GetSystemTodayPageCategories(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodayPageCategoriesAPI.GetSystemTodayPageCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemTodayPageCategories`: []TodayPageCategory
	fmt.Fprintf(os.Stdout, "Response from `TodayPageCategoriesAPI.GetSystemTodayPageCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemTodayPageCategoriesRequest struct via the builder pattern


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

[**[]TodayPageCategory**](TodayPageCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemTodayPageCategoriesById

> TodayPageCategory GetSystemTodayPageCategoriesById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get TodayPageCategory

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
	id := int32(56) // int32 | todayPageCategoryId
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
	resp, r, err := apiClient.TodayPageCategoriesAPI.GetSystemTodayPageCategoriesById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodayPageCategoriesAPI.GetSystemTodayPageCategoriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemTodayPageCategoriesById`: TodayPageCategory
	fmt.Fprintf(os.Stdout, "Response from `TodayPageCategoriesAPI.GetSystemTodayPageCategoriesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | todayPageCategoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemTodayPageCategoriesByIdRequest struct via the builder pattern


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

[**TodayPageCategory**](TodayPageCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemTodayPageCategoriesCount

> Count GetSystemTodayPageCategoriesCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of TodayPageCategory

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
	resp, r, err := apiClient.TodayPageCategoriesAPI.GetSystemTodayPageCategoriesCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodayPageCategoriesAPI.GetSystemTodayPageCategoriesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemTodayPageCategoriesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `TodayPageCategoriesAPI.GetSystemTodayPageCategoriesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemTodayPageCategoriesCountRequest struct via the builder pattern


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


## PatchSystemTodayPageCategoriesById

> TodayPageCategory PatchSystemTodayPageCategoriesById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch TodayPageCategory

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
	id := int32(56) // int32 | todayPageCategoryId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TodayPageCategoriesAPI.PatchSystemTodayPageCategoriesById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodayPageCategoriesAPI.PatchSystemTodayPageCategoriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemTodayPageCategoriesById`: TodayPageCategory
	fmt.Fprintf(os.Stdout, "Response from `TodayPageCategoriesAPI.PatchSystemTodayPageCategoriesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | todayPageCategoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemTodayPageCategoriesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**TodayPageCategory**](TodayPageCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemTodayPageCategories

> TodayPageCategory PostSystemTodayPageCategories(ctx).ClientId(clientId).TodayPageCategory(todayPageCategory).Execute()

Post TodayPageCategory

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
	todayPageCategory := *openapiclient.NewTodayPageCategory("Name_example", NullableInt32(123)) // TodayPageCategory | todayPageCategory

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TodayPageCategoriesAPI.PostSystemTodayPageCategories(context.Background()).ClientId(clientId).TodayPageCategory(todayPageCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodayPageCategoriesAPI.PostSystemTodayPageCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemTodayPageCategories`: TodayPageCategory
	fmt.Fprintf(os.Stdout, "Response from `TodayPageCategoriesAPI.PostSystemTodayPageCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemTodayPageCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **todayPageCategory** | [**TodayPageCategory**](TodayPageCategory.md) | todayPageCategory | 

### Return type

[**TodayPageCategory**](TodayPageCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemTodayPageCategoriesById

> TodayPageCategory PutSystemTodayPageCategoriesById(ctx, id).ClientId(clientId).TodayPageCategory(todayPageCategory).Execute()

Put TodayPageCategory

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
	id := int32(56) // int32 | todayPageCategoryId
	clientId := "clientId_example" // string | 
	todayPageCategory := *openapiclient.NewTodayPageCategory("Name_example", NullableInt32(123)) // TodayPageCategory | todayPageCategory

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TodayPageCategoriesAPI.PutSystemTodayPageCategoriesById(context.Background(), id).ClientId(clientId).TodayPageCategory(todayPageCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodayPageCategoriesAPI.PutSystemTodayPageCategoriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemTodayPageCategoriesById`: TodayPageCategory
	fmt.Fprintf(os.Stdout, "Response from `TodayPageCategoriesAPI.PutSystemTodayPageCategoriesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | todayPageCategoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemTodayPageCategoriesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **todayPageCategory** | [**TodayPageCategory**](TodayPageCategory.md) | todayPageCategory | 

### Return type

[**TodayPageCategory**](TodayPageCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

