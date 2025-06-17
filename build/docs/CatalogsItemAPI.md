# \CatalogsItemAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProcurementCatalogById**](CatalogsItemAPI.md#DeleteProcurementCatalogById) | **Delete** /procurement/catalog/{id} | Delete CatalogItem
[**DeleteProcurementCatalogByParentIdVendorsById**](CatalogsItemAPI.md#DeleteProcurementCatalogByParentIdVendorsById) | **Delete** /procurement/catalog/{parentId}/vendors/{id} | Delete CatalogItem
[**GetProcurementCatalog**](CatalogsItemAPI.md#GetProcurementCatalog) | **Get** /procurement/catalog | Get List of CatalogItem
[**GetProcurementCatalogByCatalogItemIdentifierQuantityOnHand**](CatalogsItemAPI.md#GetProcurementCatalogByCatalogItemIdentifierQuantityOnHand) | **Get** /procurement/catalog/{catalogItemIdentifier}/quantityOnHand | Get Count of CatalogItem
[**GetProcurementCatalogById**](CatalogsItemAPI.md#GetProcurementCatalogById) | **Get** /procurement/catalog/{id} | Get CatalogItem
[**GetProcurementCatalogByIdInfo**](CatalogsItemAPI.md#GetProcurementCatalogByIdInfo) | **Get** /procurement/catalog/{id}/info | Get CatalogItemInfo
[**GetProcurementCatalogCount**](CatalogsItemAPI.md#GetProcurementCatalogCount) | **Get** /procurement/catalog/count | Get Count of CatalogItem
[**GetProcurementCatalogInfo**](CatalogsItemAPI.md#GetProcurementCatalogInfo) | **Get** /procurement/catalog/info | Get List of CatalogItemInfo
[**GetProcurementCatalogInfoCount**](CatalogsItemAPI.md#GetProcurementCatalogInfoCount) | **Get** /procurement/catalog/info/count | Get Count of CatalogItemInfo
[**GetProcurementCatalogVendorsByParentId**](CatalogsItemAPI.md#GetProcurementCatalogVendorsByParentId) | **Get** /procurement/catalog/vendors/{parentId} | Get List of CatalogItem
[**PatchProcurementCatalogById**](CatalogsItemAPI.md#PatchProcurementCatalogById) | **Patch** /procurement/catalog/{id} | Patch CatalogItem
[**PostProcurementCatalog**](CatalogsItemAPI.md#PostProcurementCatalog) | **Post** /procurement/catalog | Post CatalogItem
[**PostProcurementCatalogByIdCopy**](CatalogsItemAPI.md#PostProcurementCatalogByIdCopy) | **Post** /procurement/catalog/{id}/copy | Post Copy CatalogItem
[**PostProcurementCatalogByIdPricing**](CatalogsItemAPI.md#PostProcurementCatalogByIdPricing) | **Post** /procurement/catalog/{id}/pricing | Post CatalogPricing
[**PostProcurementCatalogVendors**](CatalogsItemAPI.md#PostProcurementCatalogVendors) | **Post** /procurement/catalog/vendors | Post CatalogItem
[**PutProcurementCatalogById**](CatalogsItemAPI.md#PutProcurementCatalogById) | **Put** /procurement/catalog/{id} | Put CatalogItem
[**PutProcurementCatalogByParentIdVendorsById**](CatalogsItemAPI.md#PutProcurementCatalogByParentIdVendorsById) | **Put** /procurement/catalog/{parentId}/vendors/{id} | Put CatalogItem



## DeleteProcurementCatalogById

> DeleteProcurementCatalogById(ctx, id).ClientId(clientId).Execute()

Delete CatalogItem

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
	id := int32(56) // int32 | catalogId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogsItemAPI.DeleteProcurementCatalogById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsItemAPI.DeleteProcurementCatalogById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | catalogId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcurementCatalogByIdRequest struct via the builder pattern


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


## DeleteProcurementCatalogByParentIdVendorsById

> DeleteProcurementCatalogByParentIdVendorsById(ctx, parentId, id).ClientId(clientId).Execute()

Delete CatalogItem

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
	id := int32(56) // int32 | vendorId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogsItemAPI.DeleteProcurementCatalogByParentIdVendorsById(context.Background(), parentId, id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsItemAPI.DeleteProcurementCatalogByParentIdVendorsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | catalogId | 
**id** | **int32** | vendorId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcurementCatalogByParentIdVendorsByIdRequest struct via the builder pattern


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


## GetProcurementCatalog

> []CatalogItem GetProcurementCatalog(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of CatalogItem

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
	resp, r, err := apiClient.CatalogsItemAPI.GetProcurementCatalog(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsItemAPI.GetProcurementCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementCatalog`: []CatalogItem
	fmt.Fprintf(os.Stdout, "Response from `CatalogsItemAPI.GetProcurementCatalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementCatalogRequest struct via the builder pattern


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

[**[]CatalogItem**](CatalogItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementCatalogByCatalogItemIdentifierQuantityOnHand

> Count GetProcurementCatalogByCatalogItemIdentifierQuantityOnHand(ctx, catalogItemIdentifier, warehouseBinId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of CatalogItem

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
	catalogItemIdentifier := "catalogItemIdentifier_example" // string | catalogItemIdentifier
	clientId := "clientId_example" // string | 
	warehouseBinId := int32(56) // int32 | warehouseBinId
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
	resp, r, err := apiClient.CatalogsItemAPI.GetProcurementCatalogByCatalogItemIdentifierQuantityOnHand(context.Background(), catalogItemIdentifier, warehouseBinId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsItemAPI.GetProcurementCatalogByCatalogItemIdentifierQuantityOnHand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementCatalogByCatalogItemIdentifierQuantityOnHand`: Count
	fmt.Fprintf(os.Stdout, "Response from `CatalogsItemAPI.GetProcurementCatalogByCatalogItemIdentifierQuantityOnHand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogItemIdentifier** | **string** | catalogItemIdentifier | 
**warehouseBinId** | **int32** | warehouseBinId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementCatalogByCatalogItemIdentifierQuantityOnHandRequest struct via the builder pattern


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


## GetProcurementCatalogById

> CatalogItem GetProcurementCatalogById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get CatalogItem

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
	id := int32(56) // int32 | catalogId
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
	resp, r, err := apiClient.CatalogsItemAPI.GetProcurementCatalogById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsItemAPI.GetProcurementCatalogById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementCatalogById`: CatalogItem
	fmt.Fprintf(os.Stdout, "Response from `CatalogsItemAPI.GetProcurementCatalogById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | catalogId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementCatalogByIdRequest struct via the builder pattern


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

[**CatalogItem**](CatalogItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementCatalogByIdInfo

> CatalogItemInfo GetProcurementCatalogByIdInfo(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get CatalogItemInfo

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
	id := int32(56) // int32 | catalogId
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
	resp, r, err := apiClient.CatalogsItemAPI.GetProcurementCatalogByIdInfo(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsItemAPI.GetProcurementCatalogByIdInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementCatalogByIdInfo`: CatalogItemInfo
	fmt.Fprintf(os.Stdout, "Response from `CatalogsItemAPI.GetProcurementCatalogByIdInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | catalogId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementCatalogByIdInfoRequest struct via the builder pattern


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

[**CatalogItemInfo**](CatalogItemInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementCatalogCount

> Count GetProcurementCatalogCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of CatalogItem

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
	resp, r, err := apiClient.CatalogsItemAPI.GetProcurementCatalogCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsItemAPI.GetProcurementCatalogCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementCatalogCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `CatalogsItemAPI.GetProcurementCatalogCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementCatalogCountRequest struct via the builder pattern


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


## GetProcurementCatalogInfo

> []CatalogItemInfo GetProcurementCatalogInfo(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of CatalogItemInfo

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
	resp, r, err := apiClient.CatalogsItemAPI.GetProcurementCatalogInfo(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsItemAPI.GetProcurementCatalogInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementCatalogInfo`: []CatalogItemInfo
	fmt.Fprintf(os.Stdout, "Response from `CatalogsItemAPI.GetProcurementCatalogInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementCatalogInfoRequest struct via the builder pattern


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

[**[]CatalogItemInfo**](CatalogItemInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementCatalogInfoCount

> Count GetProcurementCatalogInfoCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of CatalogItemInfo

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
	resp, r, err := apiClient.CatalogsItemAPI.GetProcurementCatalogInfoCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsItemAPI.GetProcurementCatalogInfoCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementCatalogInfoCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `CatalogsItemAPI.GetProcurementCatalogInfoCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementCatalogInfoCountRequest struct via the builder pattern


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


## GetProcurementCatalogVendorsByParentId

> []CatalogVendors GetProcurementCatalogVendorsByParentId(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of CatalogItem

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
	parentId := int32(56) // int32 | vendorId
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
	resp, r, err := apiClient.CatalogsItemAPI.GetProcurementCatalogVendorsByParentId(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsItemAPI.GetProcurementCatalogVendorsByParentId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementCatalogVendorsByParentId`: []CatalogVendors
	fmt.Fprintf(os.Stdout, "Response from `CatalogsItemAPI.GetProcurementCatalogVendorsByParentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | vendorId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementCatalogVendorsByParentIdRequest struct via the builder pattern


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

[**[]CatalogVendors**](CatalogVendors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchProcurementCatalogById

> CatalogItem PatchProcurementCatalogById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch CatalogItem

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
	id := int32(56) // int32 | catalogId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsItemAPI.PatchProcurementCatalogById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsItemAPI.PatchProcurementCatalogById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProcurementCatalogById`: CatalogItem
	fmt.Fprintf(os.Stdout, "Response from `CatalogsItemAPI.PatchProcurementCatalogById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | catalogId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProcurementCatalogByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**CatalogItem**](CatalogItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProcurementCatalog

> CatalogItem PostProcurementCatalog(ctx).ClientId(clientId).CatalogItem(catalogItem).Execute()

Post CatalogItem

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
	catalogItem := *openapiclient.NewCatalogItem("Identifier_example", "Description_example", *openapiclient.NewProductSubCategoryReference(), *openapiclient.NewProductTypeReference(), "CustomerDescription_example") // CatalogItem | catalogItem

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsItemAPI.PostProcurementCatalog(context.Background()).ClientId(clientId).CatalogItem(catalogItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsItemAPI.PostProcurementCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementCatalog`: CatalogItem
	fmt.Fprintf(os.Stdout, "Response from `CatalogsItemAPI.PostProcurementCatalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **catalogItem** | [**CatalogItem**](CatalogItem.md) | catalogItem | 

### Return type

[**CatalogItem**](CatalogItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProcurementCatalogByIdCopy

> CatalogItem PostProcurementCatalogByIdCopy(ctx, id).ClientId(clientId).Execute()

Post Copy CatalogItem

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
	id := int32(56) // int32 | catalogId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsItemAPI.PostProcurementCatalogByIdCopy(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsItemAPI.PostProcurementCatalogByIdCopy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementCatalogByIdCopy`: CatalogItem
	fmt.Fprintf(os.Stdout, "Response from `CatalogsItemAPI.PostProcurementCatalogByIdCopy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | catalogId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementCatalogByIdCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 

### Return type

[**CatalogItem**](CatalogItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProcurementCatalogByIdPricing

> CatalogPricing PostProcurementCatalogByIdPricing(ctx, id).ClientId(clientId).CatalogPricing(catalogPricing).Execute()

Post CatalogPricing

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
	id := int32(56) // int32 | catalogId
	clientId := "clientId_example" // string | 
	catalogPricing := *openapiclient.NewCatalogPricing() // CatalogPricing | catalogPricing

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsItemAPI.PostProcurementCatalogByIdPricing(context.Background(), id).ClientId(clientId).CatalogPricing(catalogPricing).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsItemAPI.PostProcurementCatalogByIdPricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementCatalogByIdPricing`: CatalogPricing
	fmt.Fprintf(os.Stdout, "Response from `CatalogsItemAPI.PostProcurementCatalogByIdPricing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | catalogId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementCatalogByIdPricingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **catalogPricing** | [**CatalogPricing**](CatalogPricing.md) | catalogPricing | 

### Return type

[**CatalogPricing**](CatalogPricing.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProcurementCatalogVendors

> CatalogVendors PostProcurementCatalogVendors(ctx).ClientId(clientId).CatalogVendors(catalogVendors).Execute()

Post CatalogItem

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
	catalogVendors := *openapiclient.NewCatalogVendors() // CatalogVendors | catalogVendors

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsItemAPI.PostProcurementCatalogVendors(context.Background()).ClientId(clientId).CatalogVendors(catalogVendors).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsItemAPI.PostProcurementCatalogVendors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementCatalogVendors`: CatalogVendors
	fmt.Fprintf(os.Stdout, "Response from `CatalogsItemAPI.PostProcurementCatalogVendors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementCatalogVendorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **catalogVendors** | [**CatalogVendors**](CatalogVendors.md) | catalogVendors | 

### Return type

[**CatalogVendors**](CatalogVendors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProcurementCatalogById

> CatalogItem PutProcurementCatalogById(ctx, id).ClientId(clientId).CatalogItem(catalogItem).Execute()

Put CatalogItem

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
	id := int32(56) // int32 | catalogId
	clientId := "clientId_example" // string | 
	catalogItem := *openapiclient.NewCatalogItem("Identifier_example", "Description_example", *openapiclient.NewProductSubCategoryReference(), *openapiclient.NewProductTypeReference(), "CustomerDescription_example") // CatalogItem | catalogItem

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsItemAPI.PutProcurementCatalogById(context.Background(), id).ClientId(clientId).CatalogItem(catalogItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsItemAPI.PutProcurementCatalogById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProcurementCatalogById`: CatalogItem
	fmt.Fprintf(os.Stdout, "Response from `CatalogsItemAPI.PutProcurementCatalogById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | catalogId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProcurementCatalogByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **catalogItem** | [**CatalogItem**](CatalogItem.md) | catalogItem | 

### Return type

[**CatalogItem**](CatalogItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProcurementCatalogByParentIdVendorsById

> CatalogVendors PutProcurementCatalogByParentIdVendorsById(ctx, id, parentId).ClientId(clientId).CatalogItem(catalogItem).Execute()

Put CatalogItem

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
	id := int32(56) // int32 | vendorId
	parentId := int32(56) // int32 | catalogId
	clientId := "clientId_example" // string | 
	catalogItem := *openapiclient.NewCatalogItem("Identifier_example", "Description_example", *openapiclient.NewProductSubCategoryReference(), *openapiclient.NewProductTypeReference(), "CustomerDescription_example") // CatalogItem | catalogItem

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsItemAPI.PutProcurementCatalogByParentIdVendorsById(context.Background(), id, parentId).ClientId(clientId).CatalogItem(catalogItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsItemAPI.PutProcurementCatalogByParentIdVendorsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProcurementCatalogByParentIdVendorsById`: CatalogVendors
	fmt.Fprintf(os.Stdout, "Response from `CatalogsItemAPI.PutProcurementCatalogByParentIdVendorsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | vendorId | 
**parentId** | **int32** | catalogId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProcurementCatalogByParentIdVendorsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **catalogItem** | [**CatalogItem**](CatalogItem.md) | catalogItem | 

### Return type

[**CatalogVendors**](CatalogVendors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

