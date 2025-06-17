# \LegacySubCategoriesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProcurementCategoriesByParentIdSubcategoriesById**](LegacySubCategoriesAPI.md#DeleteProcurementCategoriesByParentIdSubcategoriesById) | **Delete** /procurement/categories/{parentId}/subcategories/{id} | Delete LegacySubCategory
[**GetProcurementCategoriesByParentIdSubcategories**](LegacySubCategoriesAPI.md#GetProcurementCategoriesByParentIdSubcategories) | **Get** /procurement/categories/{parentId}/subcategories/ | Get List of LegacySubCategory
[**GetProcurementCategoriesByParentIdSubcategoriesById**](LegacySubCategoriesAPI.md#GetProcurementCategoriesByParentIdSubcategoriesById) | **Get** /procurement/categories/{parentId}/subcategories/{id} | Get LegacySubCategory
[**GetProcurementCategoriesByParentIdSubcategoriesCount**](LegacySubCategoriesAPI.md#GetProcurementCategoriesByParentIdSubcategoriesCount) | **Get** /procurement/categories/{parentId}/subcategories/count | Get Count of LegacySubCategory
[**PatchProcurementCategoriesByParentIdSubcategoriesById**](LegacySubCategoriesAPI.md#PatchProcurementCategoriesByParentIdSubcategoriesById) | **Patch** /procurement/categories/{parentId}/subcategories/{id} | Patch LegacySubCategory
[**PostProcurementCategoriesByParentIdSubcategories**](LegacySubCategoriesAPI.md#PostProcurementCategoriesByParentIdSubcategories) | **Post** /procurement/categories/{parentId}/subcategories/ | Post LegacySubCategory
[**PutProcurementCategoriesByParentIdSubcategoriesById**](LegacySubCategoriesAPI.md#PutProcurementCategoriesByParentIdSubcategoriesById) | **Put** /procurement/categories/{parentId}/subcategories/{id} | Put LegacySubCategory



## DeleteProcurementCategoriesByParentIdSubcategoriesById

> DeleteProcurementCategoriesByParentIdSubcategoriesById(ctx, id, parentId).ClientId(clientId).Execute()

Delete LegacySubCategory

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
	id := int32(56) // int32 | subcategoryId
	parentId := int32(56) // int32 | categoryId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LegacySubCategoriesAPI.DeleteProcurementCategoriesByParentIdSubcategoriesById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegacySubCategoriesAPI.DeleteProcurementCategoriesByParentIdSubcategoriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | subcategoryId | 
**parentId** | **int32** | categoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcurementCategoriesByParentIdSubcategoriesByIdRequest struct via the builder pattern


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


## GetProcurementCategoriesByParentIdSubcategories

> []LegacySubCategory GetProcurementCategoriesByParentIdSubcategories(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of LegacySubCategory

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
	parentId := int32(56) // int32 | categoryId
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
	resp, r, err := apiClient.LegacySubCategoriesAPI.GetProcurementCategoriesByParentIdSubcategories(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegacySubCategoriesAPI.GetProcurementCategoriesByParentIdSubcategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementCategoriesByParentIdSubcategories`: []LegacySubCategory
	fmt.Fprintf(os.Stdout, "Response from `LegacySubCategoriesAPI.GetProcurementCategoriesByParentIdSubcategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | categoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementCategoriesByParentIdSubcategoriesRequest struct via the builder pattern


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

[**[]LegacySubCategory**](LegacySubCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementCategoriesByParentIdSubcategoriesById

> LegacySubCategory GetProcurementCategoriesByParentIdSubcategoriesById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get LegacySubCategory

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
	id := int32(56) // int32 | subcategoryId
	parentId := int32(56) // int32 | categoryId
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
	resp, r, err := apiClient.LegacySubCategoriesAPI.GetProcurementCategoriesByParentIdSubcategoriesById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegacySubCategoriesAPI.GetProcurementCategoriesByParentIdSubcategoriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementCategoriesByParentIdSubcategoriesById`: LegacySubCategory
	fmt.Fprintf(os.Stdout, "Response from `LegacySubCategoriesAPI.GetProcurementCategoriesByParentIdSubcategoriesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | subcategoryId | 
**parentId** | **int32** | categoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementCategoriesByParentIdSubcategoriesByIdRequest struct via the builder pattern


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

[**LegacySubCategory**](LegacySubCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementCategoriesByParentIdSubcategoriesCount

> Count GetProcurementCategoriesByParentIdSubcategoriesCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of LegacySubCategory

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
	parentId := int32(56) // int32 | categoryId
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
	resp, r, err := apiClient.LegacySubCategoriesAPI.GetProcurementCategoriesByParentIdSubcategoriesCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegacySubCategoriesAPI.GetProcurementCategoriesByParentIdSubcategoriesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementCategoriesByParentIdSubcategoriesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `LegacySubCategoriesAPI.GetProcurementCategoriesByParentIdSubcategoriesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | categoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementCategoriesByParentIdSubcategoriesCountRequest struct via the builder pattern


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


## PatchProcurementCategoriesByParentIdSubcategoriesById

> LegacySubCategory PatchProcurementCategoriesByParentIdSubcategoriesById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch LegacySubCategory

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
	id := int32(56) // int32 | subcategoryId
	parentId := int32(56) // int32 | categoryId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegacySubCategoriesAPI.PatchProcurementCategoriesByParentIdSubcategoriesById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegacySubCategoriesAPI.PatchProcurementCategoriesByParentIdSubcategoriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProcurementCategoriesByParentIdSubcategoriesById`: LegacySubCategory
	fmt.Fprintf(os.Stdout, "Response from `LegacySubCategoriesAPI.PatchProcurementCategoriesByParentIdSubcategoriesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | subcategoryId | 
**parentId** | **int32** | categoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProcurementCategoriesByParentIdSubcategoriesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**LegacySubCategory**](LegacySubCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProcurementCategoriesByParentIdSubcategories

> LegacySubCategory PostProcurementCategoriesByParentIdSubcategories(ctx, parentId).ClientId(clientId).LegacySubCategory(legacySubCategory).Execute()

Post LegacySubCategory

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
	parentId := int32(56) // int32 | categoryId
	clientId := "clientId_example" // string | 
	legacySubCategory := *openapiclient.NewLegacySubCategory("Name_example") // LegacySubCategory | subCategory

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegacySubCategoriesAPI.PostProcurementCategoriesByParentIdSubcategories(context.Background(), parentId).ClientId(clientId).LegacySubCategory(legacySubCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegacySubCategoriesAPI.PostProcurementCategoriesByParentIdSubcategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementCategoriesByParentIdSubcategories`: LegacySubCategory
	fmt.Fprintf(os.Stdout, "Response from `LegacySubCategoriesAPI.PostProcurementCategoriesByParentIdSubcategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | categoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementCategoriesByParentIdSubcategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **legacySubCategory** | [**LegacySubCategory**](LegacySubCategory.md) | subCategory | 

### Return type

[**LegacySubCategory**](LegacySubCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProcurementCategoriesByParentIdSubcategoriesById

> LegacySubCategory PutProcurementCategoriesByParentIdSubcategoriesById(ctx, id, parentId).ClientId(clientId).LegacySubCategory(legacySubCategory).Execute()

Put LegacySubCategory

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
	id := int32(56) // int32 | subcategoryId
	parentId := int32(56) // int32 | categoryId
	clientId := "clientId_example" // string | 
	legacySubCategory := *openapiclient.NewLegacySubCategory("Name_example") // LegacySubCategory | subCategory

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LegacySubCategoriesAPI.PutProcurementCategoriesByParentIdSubcategoriesById(context.Background(), id, parentId).ClientId(clientId).LegacySubCategory(legacySubCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegacySubCategoriesAPI.PutProcurementCategoriesByParentIdSubcategoriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProcurementCategoriesByParentIdSubcategoriesById`: LegacySubCategory
	fmt.Fprintf(os.Stdout, "Response from `LegacySubCategoriesAPI.PutProcurementCategoriesByParentIdSubcategoriesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | subcategoryId | 
**parentId** | **int32** | categoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProcurementCategoriesByParentIdSubcategoriesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **legacySubCategory** | [**LegacySubCategory**](LegacySubCategory.md) | subCategory | 

### Return type

[**LegacySubCategory**](LegacySubCategory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

