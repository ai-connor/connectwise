# \CatalogComponentsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProcurementCatalogByParentIdComponentsById**](CatalogComponentsAPI.md#DeleteProcurementCatalogByParentIdComponentsById) | **Delete** /procurement/catalog/{parentId}/components/{id} | Delete CatalogComponent
[**GetProcurementCatalogByParentIdComponents**](CatalogComponentsAPI.md#GetProcurementCatalogByParentIdComponents) | **Get** /procurement/catalog/{parentId}/components | Get List of CatalogComponent
[**GetProcurementCatalogByParentIdComponentsById**](CatalogComponentsAPI.md#GetProcurementCatalogByParentIdComponentsById) | **Get** /procurement/catalog/{parentId}/components/{id} | Get CatalogComponent
[**GetProcurementCatalogByParentIdComponentsCount**](CatalogComponentsAPI.md#GetProcurementCatalogByParentIdComponentsCount) | **Get** /procurement/catalog/{parentId}/components/count | Get Count of CatalogComponent
[**PatchProcurementCatalogByParentIdComponentsById**](CatalogComponentsAPI.md#PatchProcurementCatalogByParentIdComponentsById) | **Patch** /procurement/catalog/{parentId}/components/{id} | Patch CatalogComponent
[**PostProcurementCatalogByParentIdComponents**](CatalogComponentsAPI.md#PostProcurementCatalogByParentIdComponents) | **Post** /procurement/catalog/{parentId}/components | Post CatalogComponent
[**PutProcurementCatalogByParentIdComponentsById**](CatalogComponentsAPI.md#PutProcurementCatalogByParentIdComponentsById) | **Put** /procurement/catalog/{parentId}/components/{id} | Put CatalogComponent



## DeleteProcurementCatalogByParentIdComponentsById

> DeleteProcurementCatalogByParentIdComponentsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete CatalogComponent

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
	parentId := int32(56) // int32 | catalogId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogComponentsAPI.DeleteProcurementCatalogByParentIdComponentsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogComponentsAPI.DeleteProcurementCatalogByParentIdComponentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | componentId | 
**parentId** | **int32** | catalogId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcurementCatalogByParentIdComponentsByIdRequest struct via the builder pattern


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


## GetProcurementCatalogByParentIdComponents

> []CatalogComponent GetProcurementCatalogByParentIdComponents(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of CatalogComponent

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
	parentId := int32(56) // int32 | catalogId
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
	resp, r, err := apiClient.CatalogComponentsAPI.GetProcurementCatalogByParentIdComponents(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogComponentsAPI.GetProcurementCatalogByParentIdComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementCatalogByParentIdComponents`: []CatalogComponent
	fmt.Fprintf(os.Stdout, "Response from `CatalogComponentsAPI.GetProcurementCatalogByParentIdComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | catalogId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementCatalogByParentIdComponentsRequest struct via the builder pattern


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

[**[]CatalogComponent**](CatalogComponent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementCatalogByParentIdComponentsById

> CatalogComponent GetProcurementCatalogByParentIdComponentsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get CatalogComponent

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
	parentId := int32(56) // int32 | catalogId
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
	resp, r, err := apiClient.CatalogComponentsAPI.GetProcurementCatalogByParentIdComponentsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogComponentsAPI.GetProcurementCatalogByParentIdComponentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementCatalogByParentIdComponentsById`: CatalogComponent
	fmt.Fprintf(os.Stdout, "Response from `CatalogComponentsAPI.GetProcurementCatalogByParentIdComponentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | componentId | 
**parentId** | **int32** | catalogId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementCatalogByParentIdComponentsByIdRequest struct via the builder pattern


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

[**CatalogComponent**](CatalogComponent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementCatalogByParentIdComponentsCount

> Count GetProcurementCatalogByParentIdComponentsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of CatalogComponent

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
	parentId := int32(56) // int32 | catalogId
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
	resp, r, err := apiClient.CatalogComponentsAPI.GetProcurementCatalogByParentIdComponentsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogComponentsAPI.GetProcurementCatalogByParentIdComponentsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementCatalogByParentIdComponentsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `CatalogComponentsAPI.GetProcurementCatalogByParentIdComponentsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | catalogId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementCatalogByParentIdComponentsCountRequest struct via the builder pattern


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


## PatchProcurementCatalogByParentIdComponentsById

> CatalogComponent PatchProcurementCatalogByParentIdComponentsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch CatalogComponent

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
	parentId := int32(56) // int32 | catalogId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogComponentsAPI.PatchProcurementCatalogByParentIdComponentsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogComponentsAPI.PatchProcurementCatalogByParentIdComponentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProcurementCatalogByParentIdComponentsById`: CatalogComponent
	fmt.Fprintf(os.Stdout, "Response from `CatalogComponentsAPI.PatchProcurementCatalogByParentIdComponentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | componentId | 
**parentId** | **int32** | catalogId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProcurementCatalogByParentIdComponentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**CatalogComponent**](CatalogComponent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProcurementCatalogByParentIdComponents

> CatalogComponent PostProcurementCatalogByParentIdComponents(ctx, parentId).ClientId(clientId).CatalogComponent(catalogComponent).Execute()

Post CatalogComponent

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
	parentId := int32(56) // int32 | catalogId
	clientId := "clientId_example" // string | 
	catalogComponent := *openapiclient.NewCatalogComponent(NullableFloat64(123), *openapiclient.NewCatalogItemReference()) // CatalogComponent | catalogComponent

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogComponentsAPI.PostProcurementCatalogByParentIdComponents(context.Background(), parentId).ClientId(clientId).CatalogComponent(catalogComponent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogComponentsAPI.PostProcurementCatalogByParentIdComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementCatalogByParentIdComponents`: CatalogComponent
	fmt.Fprintf(os.Stdout, "Response from `CatalogComponentsAPI.PostProcurementCatalogByParentIdComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | catalogId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementCatalogByParentIdComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **catalogComponent** | [**CatalogComponent**](CatalogComponent.md) | catalogComponent | 

### Return type

[**CatalogComponent**](CatalogComponent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProcurementCatalogByParentIdComponentsById

> CatalogComponent PutProcurementCatalogByParentIdComponentsById(ctx, id, parentId).ClientId(clientId).CatalogComponent(catalogComponent).Execute()

Put CatalogComponent

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
	parentId := int32(56) // int32 | catalogId
	clientId := "clientId_example" // string | 
	catalogComponent := *openapiclient.NewCatalogComponent(NullableFloat64(123), *openapiclient.NewCatalogItemReference()) // CatalogComponent | catalogComponent

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogComponentsAPI.PutProcurementCatalogByParentIdComponentsById(context.Background(), id, parentId).ClientId(clientId).CatalogComponent(catalogComponent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogComponentsAPI.PutProcurementCatalogByParentIdComponentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProcurementCatalogByParentIdComponentsById`: CatalogComponent
	fmt.Fprintf(os.Stdout, "Response from `CatalogComponentsAPI.PutProcurementCatalogByParentIdComponentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | componentId | 
**parentId** | **int32** | catalogId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProcurementCatalogByParentIdComponentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **catalogComponent** | [**CatalogComponent**](CatalogComponent.md) | catalogComponent | 

### Return type

[**CatalogComponent**](CatalogComponent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

