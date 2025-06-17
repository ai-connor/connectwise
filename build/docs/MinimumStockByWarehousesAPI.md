# \MinimumStockByWarehousesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProcurementCatalogByParentIdMinimumStockByWarehouseById**](MinimumStockByWarehousesAPI.md#DeleteProcurementCatalogByParentIdMinimumStockByWarehouseById) | **Delete** /procurement/catalog/{parentId}/minimumStockByWarehouse/{id} | Delete MinimumStockByWarehouse
[**GetProcurementCatalogByParentIdMinimumStockByWarehouse**](MinimumStockByWarehousesAPI.md#GetProcurementCatalogByParentIdMinimumStockByWarehouse) | **Get** /procurement/catalog/{parentId}/minimumStockByWarehouse | Get List of MinimumStockByWarehouse
[**GetProcurementCatalogByParentIdMinimumStockByWarehouseById**](MinimumStockByWarehousesAPI.md#GetProcurementCatalogByParentIdMinimumStockByWarehouseById) | **Get** /procurement/catalog/{parentId}/minimumStockByWarehouse/{id} | Get MinimumStockByWarehouse
[**GetProcurementCatalogByParentIdMinimumStockByWarehouseCount**](MinimumStockByWarehousesAPI.md#GetProcurementCatalogByParentIdMinimumStockByWarehouseCount) | **Get** /procurement/catalog/{parentId}/minimumStockByWarehouse/count | Get Count of MinimumStockByWarehouse
[**PatchProcurementCatalogByParentIdMinimumStockByWarehouseById**](MinimumStockByWarehousesAPI.md#PatchProcurementCatalogByParentIdMinimumStockByWarehouseById) | **Patch** /procurement/catalog/{parentId}/minimumStockByWarehouse/{id} | Patch MinimumStockByWarehouse
[**PostProcurementCatalogByParentIdMinimumStockByWarehouse**](MinimumStockByWarehousesAPI.md#PostProcurementCatalogByParentIdMinimumStockByWarehouse) | **Post** /procurement/catalog/{parentId}/minimumStockByWarehouse | Post MinimumStockByWarehouse
[**PutProcurementCatalogByParentIdMinimumStockByWarehouseById**](MinimumStockByWarehousesAPI.md#PutProcurementCatalogByParentIdMinimumStockByWarehouseById) | **Put** /procurement/catalog/{parentId}/minimumStockByWarehouse/{id} | Put MinimumStockByWarehouse



## DeleteProcurementCatalogByParentIdMinimumStockByWarehouseById

> DeleteProcurementCatalogByParentIdMinimumStockByWarehouseById(ctx, id, parentId).ClientId(clientId).Execute()

Delete MinimumStockByWarehouse

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
	id := int32(56) // int32 | minimumStockByWarehouseId
	parentId := int32(56) // int32 | catalogId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MinimumStockByWarehousesAPI.DeleteProcurementCatalogByParentIdMinimumStockByWarehouseById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MinimumStockByWarehousesAPI.DeleteProcurementCatalogByParentIdMinimumStockByWarehouseById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | minimumStockByWarehouseId | 
**parentId** | **int32** | catalogId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcurementCatalogByParentIdMinimumStockByWarehouseByIdRequest struct via the builder pattern


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


## GetProcurementCatalogByParentIdMinimumStockByWarehouse

> []MinimumStockByWarehouse GetProcurementCatalogByParentIdMinimumStockByWarehouse(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of MinimumStockByWarehouse

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
	resp, r, err := apiClient.MinimumStockByWarehousesAPI.GetProcurementCatalogByParentIdMinimumStockByWarehouse(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MinimumStockByWarehousesAPI.GetProcurementCatalogByParentIdMinimumStockByWarehouse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementCatalogByParentIdMinimumStockByWarehouse`: []MinimumStockByWarehouse
	fmt.Fprintf(os.Stdout, "Response from `MinimumStockByWarehousesAPI.GetProcurementCatalogByParentIdMinimumStockByWarehouse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | catalogId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementCatalogByParentIdMinimumStockByWarehouseRequest struct via the builder pattern


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

[**[]MinimumStockByWarehouse**](MinimumStockByWarehouse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementCatalogByParentIdMinimumStockByWarehouseById

> MinimumStockByWarehouse GetProcurementCatalogByParentIdMinimumStockByWarehouseById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get MinimumStockByWarehouse

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
	id := int32(56) // int32 | minimumStockByWarehouseId
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
	resp, r, err := apiClient.MinimumStockByWarehousesAPI.GetProcurementCatalogByParentIdMinimumStockByWarehouseById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MinimumStockByWarehousesAPI.GetProcurementCatalogByParentIdMinimumStockByWarehouseById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementCatalogByParentIdMinimumStockByWarehouseById`: MinimumStockByWarehouse
	fmt.Fprintf(os.Stdout, "Response from `MinimumStockByWarehousesAPI.GetProcurementCatalogByParentIdMinimumStockByWarehouseById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | minimumStockByWarehouseId | 
**parentId** | **int32** | catalogId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementCatalogByParentIdMinimumStockByWarehouseByIdRequest struct via the builder pattern


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

[**MinimumStockByWarehouse**](MinimumStockByWarehouse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementCatalogByParentIdMinimumStockByWarehouseCount

> Count GetProcurementCatalogByParentIdMinimumStockByWarehouseCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of MinimumStockByWarehouse

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
	resp, r, err := apiClient.MinimumStockByWarehousesAPI.GetProcurementCatalogByParentIdMinimumStockByWarehouseCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MinimumStockByWarehousesAPI.GetProcurementCatalogByParentIdMinimumStockByWarehouseCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementCatalogByParentIdMinimumStockByWarehouseCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `MinimumStockByWarehousesAPI.GetProcurementCatalogByParentIdMinimumStockByWarehouseCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | catalogId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementCatalogByParentIdMinimumStockByWarehouseCountRequest struct via the builder pattern


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


## PatchProcurementCatalogByParentIdMinimumStockByWarehouseById

> MinimumStockByWarehouse PatchProcurementCatalogByParentIdMinimumStockByWarehouseById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch MinimumStockByWarehouse

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
	id := int32(56) // int32 | minimumStockByWarehouseId
	parentId := int32(56) // int32 | catalogId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MinimumStockByWarehousesAPI.PatchProcurementCatalogByParentIdMinimumStockByWarehouseById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MinimumStockByWarehousesAPI.PatchProcurementCatalogByParentIdMinimumStockByWarehouseById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProcurementCatalogByParentIdMinimumStockByWarehouseById`: MinimumStockByWarehouse
	fmt.Fprintf(os.Stdout, "Response from `MinimumStockByWarehousesAPI.PatchProcurementCatalogByParentIdMinimumStockByWarehouseById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | minimumStockByWarehouseId | 
**parentId** | **int32** | catalogId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProcurementCatalogByParentIdMinimumStockByWarehouseByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**MinimumStockByWarehouse**](MinimumStockByWarehouse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProcurementCatalogByParentIdMinimumStockByWarehouse

> MinimumStockByWarehouse PostProcurementCatalogByParentIdMinimumStockByWarehouse(ctx, parentId).ClientId(clientId).MinimumStockByWarehouse(minimumStockByWarehouse).Execute()

Post MinimumStockByWarehouse

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
	minimumStockByWarehouse := *openapiclient.NewMinimumStockByWarehouse(*openapiclient.NewWarehouseReference(), NullableInt32(123)) // MinimumStockByWarehouse | minimumStockByWarehouse

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MinimumStockByWarehousesAPI.PostProcurementCatalogByParentIdMinimumStockByWarehouse(context.Background(), parentId).ClientId(clientId).MinimumStockByWarehouse(minimumStockByWarehouse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MinimumStockByWarehousesAPI.PostProcurementCatalogByParentIdMinimumStockByWarehouse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementCatalogByParentIdMinimumStockByWarehouse`: MinimumStockByWarehouse
	fmt.Fprintf(os.Stdout, "Response from `MinimumStockByWarehousesAPI.PostProcurementCatalogByParentIdMinimumStockByWarehouse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | catalogId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementCatalogByParentIdMinimumStockByWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **minimumStockByWarehouse** | [**MinimumStockByWarehouse**](MinimumStockByWarehouse.md) | minimumStockByWarehouse | 

### Return type

[**MinimumStockByWarehouse**](MinimumStockByWarehouse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProcurementCatalogByParentIdMinimumStockByWarehouseById

> MinimumStockByWarehouse PutProcurementCatalogByParentIdMinimumStockByWarehouseById(ctx, id, parentId).ClientId(clientId).MinimumStockByWarehouse(minimumStockByWarehouse).Execute()

Put MinimumStockByWarehouse

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
	id := int32(56) // int32 | minimumStockByWarehouseId
	parentId := int32(56) // int32 | catalogId
	clientId := "clientId_example" // string | 
	minimumStockByWarehouse := *openapiclient.NewMinimumStockByWarehouse(*openapiclient.NewWarehouseReference(), NullableInt32(123)) // MinimumStockByWarehouse | minimumStockByWarehouse

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MinimumStockByWarehousesAPI.PutProcurementCatalogByParentIdMinimumStockByWarehouseById(context.Background(), id, parentId).ClientId(clientId).MinimumStockByWarehouse(minimumStockByWarehouse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MinimumStockByWarehousesAPI.PutProcurementCatalogByParentIdMinimumStockByWarehouseById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProcurementCatalogByParentIdMinimumStockByWarehouseById`: MinimumStockByWarehouse
	fmt.Fprintf(os.Stdout, "Response from `MinimumStockByWarehousesAPI.PutProcurementCatalogByParentIdMinimumStockByWarehouseById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | minimumStockByWarehouseId | 
**parentId** | **int32** | catalogId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProcurementCatalogByParentIdMinimumStockByWarehouseByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **minimumStockByWarehouse** | [**MinimumStockByWarehouse**](MinimumStockByWarehouse.md) | minimumStockByWarehouse | 

### Return type

[**MinimumStockByWarehouse**](MinimumStockByWarehouse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

