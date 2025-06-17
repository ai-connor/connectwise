# \ProductPickingShippingDetailsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProcurementProductsByParentIdPickingShippingDetailsById**](ProductPickingShippingDetailsAPI.md#DeleteProcurementProductsByParentIdPickingShippingDetailsById) | **Delete** /procurement/products/{parentId}/pickingShippingDetails/{id} | Delete ProductPickingShippingDetail
[**GetProcurementProductsByParentIdPickingShippingDetails**](ProductPickingShippingDetailsAPI.md#GetProcurementProductsByParentIdPickingShippingDetails) | **Get** /procurement/products/{parentId}/pickingShippingDetails | Get List of ProductPickingShippingDetail
[**GetProcurementProductsByParentIdPickingShippingDetailsById**](ProductPickingShippingDetailsAPI.md#GetProcurementProductsByParentIdPickingShippingDetailsById) | **Get** /procurement/products/{parentId}/pickingShippingDetails/{id} | Get List of ProductPickingShippingDetail
[**GetProcurementProductsByParentIdPickingShippingDetailsCount**](ProductPickingShippingDetailsAPI.md#GetProcurementProductsByParentIdPickingShippingDetailsCount) | **Get** /procurement/products/{parentId}/pickingShippingDetails/count | Get Count of ProductPickingShippingDetail
[**PatchProcurementProductsByParentIdPickingShippingDetailsById**](ProductPickingShippingDetailsAPI.md#PatchProcurementProductsByParentIdPickingShippingDetailsById) | **Patch** /procurement/products/{parentId}/pickingShippingDetails/{id} | Patch List of ProductPickingShippingDetail
[**PostProcurementProductsByParentIdPickingShippingDetails**](ProductPickingShippingDetailsAPI.md#PostProcurementProductsByParentIdPickingShippingDetails) | **Post** /procurement/products/{parentId}/pickingShippingDetails | Post List of ProductPickingShippingDetail
[**PutProcurementProductsByParentIdPickingShippingDetailsById**](ProductPickingShippingDetailsAPI.md#PutProcurementProductsByParentIdPickingShippingDetailsById) | **Put** /procurement/products/{parentId}/pickingShippingDetails/{id} | Put List of ProductPickingShippingDetail



## DeleteProcurementProductsByParentIdPickingShippingDetailsById

> DeleteProcurementProductsByParentIdPickingShippingDetailsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ProductPickingShippingDetail

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
	id := int32(56) // int32 | pickingShippingDetailId
	parentId := int32(56) // int32 | productId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProductPickingShippingDetailsAPI.DeleteProcurementProductsByParentIdPickingShippingDetailsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductPickingShippingDetailsAPI.DeleteProcurementProductsByParentIdPickingShippingDetailsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | pickingShippingDetailId | 
**parentId** | **int32** | productId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcurementProductsByParentIdPickingShippingDetailsByIdRequest struct via the builder pattern


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


## GetProcurementProductsByParentIdPickingShippingDetails

> []ProductPickingShippingDetail GetProcurementProductsByParentIdPickingShippingDetails(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProductPickingShippingDetail

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
	resp, r, err := apiClient.ProductPickingShippingDetailsAPI.GetProcurementProductsByParentIdPickingShippingDetails(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductPickingShippingDetailsAPI.GetProcurementProductsByParentIdPickingShippingDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementProductsByParentIdPickingShippingDetails`: []ProductPickingShippingDetail
	fmt.Fprintf(os.Stdout, "Response from `ProductPickingShippingDetailsAPI.GetProcurementProductsByParentIdPickingShippingDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | productId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementProductsByParentIdPickingShippingDetailsRequest struct via the builder pattern


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

[**[]ProductPickingShippingDetail**](ProductPickingShippingDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementProductsByParentIdPickingShippingDetailsById

> []ProductPickingShippingDetail GetProcurementProductsByParentIdPickingShippingDetailsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProductPickingShippingDetail

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
	id := int32(56) // int32 | pickingShippingDetailId
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
	resp, r, err := apiClient.ProductPickingShippingDetailsAPI.GetProcurementProductsByParentIdPickingShippingDetailsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductPickingShippingDetailsAPI.GetProcurementProductsByParentIdPickingShippingDetailsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementProductsByParentIdPickingShippingDetailsById`: []ProductPickingShippingDetail
	fmt.Fprintf(os.Stdout, "Response from `ProductPickingShippingDetailsAPI.GetProcurementProductsByParentIdPickingShippingDetailsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | pickingShippingDetailId | 
**parentId** | **int32** | productId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementProductsByParentIdPickingShippingDetailsByIdRequest struct via the builder pattern


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

[**[]ProductPickingShippingDetail**](ProductPickingShippingDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementProductsByParentIdPickingShippingDetailsCount

> Count GetProcurementProductsByParentIdPickingShippingDetailsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ProductPickingShippingDetail

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
	resp, r, err := apiClient.ProductPickingShippingDetailsAPI.GetProcurementProductsByParentIdPickingShippingDetailsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductPickingShippingDetailsAPI.GetProcurementProductsByParentIdPickingShippingDetailsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementProductsByParentIdPickingShippingDetailsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ProductPickingShippingDetailsAPI.GetProcurementProductsByParentIdPickingShippingDetailsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | productId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementProductsByParentIdPickingShippingDetailsCountRequest struct via the builder pattern


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


## PatchProcurementProductsByParentIdPickingShippingDetailsById

> []ProductPickingShippingDetail PatchProcurementProductsByParentIdPickingShippingDetailsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch List of ProductPickingShippingDetail

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
	id := int32(56) // int32 | pickingShippingDetailId
	parentId := int32(56) // int32 | productId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductPickingShippingDetailsAPI.PatchProcurementProductsByParentIdPickingShippingDetailsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductPickingShippingDetailsAPI.PatchProcurementProductsByParentIdPickingShippingDetailsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProcurementProductsByParentIdPickingShippingDetailsById`: []ProductPickingShippingDetail
	fmt.Fprintf(os.Stdout, "Response from `ProductPickingShippingDetailsAPI.PatchProcurementProductsByParentIdPickingShippingDetailsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | pickingShippingDetailId | 
**parentId** | **int32** | productId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProcurementProductsByParentIdPickingShippingDetailsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**[]ProductPickingShippingDetail**](ProductPickingShippingDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProcurementProductsByParentIdPickingShippingDetails

> []ProductPickingShippingDetail PostProcurementProductsByParentIdPickingShippingDetails(ctx, parentId).ClientId(clientId).ProductPickingShippingDetail(productPickingShippingDetail).Execute()

Post List of ProductPickingShippingDetail

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
	productPickingShippingDetail := *openapiclient.NewProductPickingShippingDetail(*openapiclient.NewWarehouseReference(), *openapiclient.NewWarehouseBinReference()) // ProductPickingShippingDetail | productPickingShippingDetails

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductPickingShippingDetailsAPI.PostProcurementProductsByParentIdPickingShippingDetails(context.Background(), parentId).ClientId(clientId).ProductPickingShippingDetail(productPickingShippingDetail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductPickingShippingDetailsAPI.PostProcurementProductsByParentIdPickingShippingDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementProductsByParentIdPickingShippingDetails`: []ProductPickingShippingDetail
	fmt.Fprintf(os.Stdout, "Response from `ProductPickingShippingDetailsAPI.PostProcurementProductsByParentIdPickingShippingDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | productId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementProductsByParentIdPickingShippingDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **productPickingShippingDetail** | [**ProductPickingShippingDetail**](ProductPickingShippingDetail.md) | productPickingShippingDetails | 

### Return type

[**[]ProductPickingShippingDetail**](ProductPickingShippingDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProcurementProductsByParentIdPickingShippingDetailsById

> []ProductPickingShippingDetail PutProcurementProductsByParentIdPickingShippingDetailsById(ctx, id, parentId).ClientId(clientId).ProductPickingShippingDetail(productPickingShippingDetail).Execute()

Put List of ProductPickingShippingDetail

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
	id := int32(56) // int32 | pickingShippingDetailId
	parentId := int32(56) // int32 | productId
	clientId := "clientId_example" // string | 
	productPickingShippingDetail := *openapiclient.NewProductPickingShippingDetail(*openapiclient.NewWarehouseReference(), *openapiclient.NewWarehouseBinReference()) // ProductPickingShippingDetail | productPickingShippingDetails

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductPickingShippingDetailsAPI.PutProcurementProductsByParentIdPickingShippingDetailsById(context.Background(), id, parentId).ClientId(clientId).ProductPickingShippingDetail(productPickingShippingDetail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductPickingShippingDetailsAPI.PutProcurementProductsByParentIdPickingShippingDetailsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProcurementProductsByParentIdPickingShippingDetailsById`: []ProductPickingShippingDetail
	fmt.Fprintf(os.Stdout, "Response from `ProductPickingShippingDetailsAPI.PutProcurementProductsByParentIdPickingShippingDetailsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | pickingShippingDetailId | 
**parentId** | **int32** | productId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProcurementProductsByParentIdPickingShippingDetailsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **productPickingShippingDetail** | [**ProductPickingShippingDetail**](ProductPickingShippingDetail.md) | productPickingShippingDetails | 

### Return type

[**[]ProductPickingShippingDetail**](ProductPickingShippingDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

