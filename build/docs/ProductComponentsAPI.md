# \ProductComponentsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProcurementProductsByParentIdComponentsById**](ProductComponentsAPI.md#DeleteProcurementProductsByParentIdComponentsById) | **Delete** /procurement/products/{parentId}/components/{id} | Delete ProductComponent
[**GetProcurementProductsByParentIdComponents**](ProductComponentsAPI.md#GetProcurementProductsByParentIdComponents) | **Get** /procurement/products/{parentId}/components | Get List of ProductComponent
[**GetProcurementProductsByParentIdComponentsById**](ProductComponentsAPI.md#GetProcurementProductsByParentIdComponentsById) | **Get** /procurement/products/{parentId}/components/{id} | Get List of ProductComponent
[**GetProcurementProductsByParentIdComponentsCount**](ProductComponentsAPI.md#GetProcurementProductsByParentIdComponentsCount) | **Get** /procurement/products/{parentId}/components/count | Get Count of ProductComponent
[**PatchProcurementProductsByParentIdComponentsById**](ProductComponentsAPI.md#PatchProcurementProductsByParentIdComponentsById) | **Patch** /procurement/products/{parentId}/components/{id} | Patch List of ProductComponent
[**PostProcurementProductsByParentIdComponents**](ProductComponentsAPI.md#PostProcurementProductsByParentIdComponents) | **Post** /procurement/products/{parentId}/components | Post List of ProductComponent
[**PutProcurementProductsByParentIdComponentsById**](ProductComponentsAPI.md#PutProcurementProductsByParentIdComponentsById) | **Put** /procurement/products/{parentId}/components/{id} | Put List of ProductComponent



## DeleteProcurementProductsByParentIdComponentsById

> DeleteProcurementProductsByParentIdComponentsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ProductComponent

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
	id := int32(56) // int32 | componentId
	parentId := int32(56) // int32 | productId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProductComponentsAPI.DeleteProcurementProductsByParentIdComponentsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductComponentsAPI.DeleteProcurementProductsByParentIdComponentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | componentId | 
**parentId** | **int32** | productId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcurementProductsByParentIdComponentsByIdRequest struct via the builder pattern


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


## GetProcurementProductsByParentIdComponents

> []ProductComponent GetProcurementProductsByParentIdComponents(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProductComponent

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
	parentId := int32(56) // int32 | productId
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
	resp, r, err := apiClient.ProductComponentsAPI.GetProcurementProductsByParentIdComponents(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductComponentsAPI.GetProcurementProductsByParentIdComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementProductsByParentIdComponents`: []ProductComponent
	fmt.Fprintf(os.Stdout, "Response from `ProductComponentsAPI.GetProcurementProductsByParentIdComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | productId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementProductsByParentIdComponentsRequest struct via the builder pattern


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

[**[]ProductComponent**](ProductComponent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementProductsByParentIdComponentsById

> []ProductComponent GetProcurementProductsByParentIdComponentsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProductComponent

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
	id := int32(56) // int32 | componentId
	parentId := int32(56) // int32 | productId
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
	resp, r, err := apiClient.ProductComponentsAPI.GetProcurementProductsByParentIdComponentsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductComponentsAPI.GetProcurementProductsByParentIdComponentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementProductsByParentIdComponentsById`: []ProductComponent
	fmt.Fprintf(os.Stdout, "Response from `ProductComponentsAPI.GetProcurementProductsByParentIdComponentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | componentId | 
**parentId** | **int32** | productId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementProductsByParentIdComponentsByIdRequest struct via the builder pattern


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

[**[]ProductComponent**](ProductComponent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementProductsByParentIdComponentsCount

> Count GetProcurementProductsByParentIdComponentsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ProductComponent

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
	parentId := int32(56) // int32 | productId
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
	resp, r, err := apiClient.ProductComponentsAPI.GetProcurementProductsByParentIdComponentsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductComponentsAPI.GetProcurementProductsByParentIdComponentsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementProductsByParentIdComponentsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ProductComponentsAPI.GetProcurementProductsByParentIdComponentsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | productId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementProductsByParentIdComponentsCountRequest struct via the builder pattern


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


## PatchProcurementProductsByParentIdComponentsById

> []ProductComponent PatchProcurementProductsByParentIdComponentsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch List of ProductComponent

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
	id := int32(56) // int32 | componentId
	parentId := int32(56) // int32 | productId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductComponentsAPI.PatchProcurementProductsByParentIdComponentsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductComponentsAPI.PatchProcurementProductsByParentIdComponentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProcurementProductsByParentIdComponentsById`: []ProductComponent
	fmt.Fprintf(os.Stdout, "Response from `ProductComponentsAPI.PatchProcurementProductsByParentIdComponentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | componentId | 
**parentId** | **int32** | productId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProcurementProductsByParentIdComponentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**[]ProductComponent**](ProductComponent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProcurementProductsByParentIdComponents

> []ProductComponent PostProcurementProductsByParentIdComponents(ctx, parentId).ClientId(clientId).ProductComponent(productComponent).Execute()

Post List of ProductComponent

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
	parentId := int32(56) // int32 | productId
	clientId := "clientId_example" // string | 
	productComponent := *openapiclient.NewProductComponent(NullableFloat64(123), *openapiclient.NewCatalogItemReference()) // ProductComponent | productComponent

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductComponentsAPI.PostProcurementProductsByParentIdComponents(context.Background(), parentId).ClientId(clientId).ProductComponent(productComponent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductComponentsAPI.PostProcurementProductsByParentIdComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementProductsByParentIdComponents`: []ProductComponent
	fmt.Fprintf(os.Stdout, "Response from `ProductComponentsAPI.PostProcurementProductsByParentIdComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | productId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementProductsByParentIdComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **productComponent** | [**ProductComponent**](ProductComponent.md) | productComponent | 

### Return type

[**[]ProductComponent**](ProductComponent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProcurementProductsByParentIdComponentsById

> []ProductComponent PutProcurementProductsByParentIdComponentsById(ctx, id, parentId).ClientId(clientId).ProductComponent(productComponent).Execute()

Put List of ProductComponent

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
	id := int32(56) // int32 | componentId
	parentId := int32(56) // int32 | productId
	clientId := "clientId_example" // string | 
	productComponent := *openapiclient.NewProductComponent(NullableFloat64(123), *openapiclient.NewCatalogItemReference()) // ProductComponent | productComponent

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductComponentsAPI.PutProcurementProductsByParentIdComponentsById(context.Background(), id, parentId).ClientId(clientId).ProductComponent(productComponent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductComponentsAPI.PutProcurementProductsByParentIdComponentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProcurementProductsByParentIdComponentsById`: []ProductComponent
	fmt.Fprintf(os.Stdout, "Response from `ProductComponentsAPI.PutProcurementProductsByParentIdComponentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | componentId | 
**parentId** | **int32** | productId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProcurementProductsByParentIdComponentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **productComponent** | [**ProductComponent**](ProductComponent.md) | productComponent | 

### Return type

[**[]ProductComponent**](ProductComponent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

