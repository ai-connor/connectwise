# \PurchaseOrderLineItemsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProcurementPurchaseordersByParentIdLineitems**](PurchaseOrderLineItemsAPI.md#DeleteProcurementPurchaseordersByParentIdLineitems) | **Delete** /procurement/purchaseorders/{parentId}/lineitems | Delete PurchaseOrderLineItem
[**DeleteProcurementPurchaseordersByParentIdLineitemsBulk**](PurchaseOrderLineItemsAPI.md#DeleteProcurementPurchaseordersByParentIdLineitemsBulk) | **Delete** /procurement/purchaseorders/{parentId}/lineitems/bulk | Delete BulkResult
[**GetProcurementPurchaseordersByParentIdLineitems**](PurchaseOrderLineItemsAPI.md#GetProcurementPurchaseordersByParentIdLineitems) | **Get** /procurement/purchaseorders/{parentId}/lineitems | Get List of PurchaseOrderLineItem
[**GetProcurementPurchaseordersByParentIdLineitemsById**](PurchaseOrderLineItemsAPI.md#GetProcurementPurchaseordersByParentIdLineitemsById) | **Get** /procurement/purchaseorders/{parentId}/lineitems/{id} | Get PurchaseOrderLineItem
[**GetProcurementPurchaseordersByParentIdLineitemsCount**](PurchaseOrderLineItemsAPI.md#GetProcurementPurchaseordersByParentIdLineitemsCount) | **Get** /procurement/purchaseorders/{parentId}/lineitems/count | Get Count of PurchaseOrderLineItem
[**PatchProcurementPurchaseordersByParentIdLineitemsById**](PurchaseOrderLineItemsAPI.md#PatchProcurementPurchaseordersByParentIdLineitemsById) | **Patch** /procurement/purchaseorders/{parentId}/lineitems/{id} | Patch PurchaseOrderLineItem
[**PostProcurementPurchaseordersByParentIdLineitems**](PurchaseOrderLineItemsAPI.md#PostProcurementPurchaseordersByParentIdLineitems) | **Post** /procurement/purchaseorders/{parentId}/lineitems | Post PurchaseOrderLineItem
[**PostProcurementPurchaseordersByParentIdLineitemsBulk**](PurchaseOrderLineItemsAPI.md#PostProcurementPurchaseordersByParentIdLineitemsBulk) | **Post** /procurement/purchaseorders/{parentId}/lineitems/bulk | Post BulkResult
[**PutProcurementPurchaseordersByParentIdLineitemsBulk**](PurchaseOrderLineItemsAPI.md#PutProcurementPurchaseordersByParentIdLineitemsBulk) | **Put** /procurement/purchaseorders/{parentId}/lineitems/bulk | Put BulkResult
[**PutProcurementPurchaseordersByParentIdLineitemsById**](PurchaseOrderLineItemsAPI.md#PutProcurementPurchaseordersByParentIdLineitemsById) | **Put** /procurement/purchaseorders/{parentId}/lineitems/{id} | Put PurchaseOrderLineItem



## DeleteProcurementPurchaseordersByParentIdLineitems

> DeleteProcurementPurchaseordersByParentIdLineitems(ctx, parentId, id).ClientId(clientId).Execute()

Delete PurchaseOrderLineItem

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
	parentId := int32(56) // int32 | purchaseorderId
	clientId := "clientId_example" // string | 
	id := int32(56) // int32 | lineitemId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PurchaseOrderLineItemsAPI.DeleteProcurementPurchaseordersByParentIdLineitems(context.Background(), parentId, id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderLineItemsAPI.DeleteProcurementPurchaseordersByParentIdLineitems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | purchaseorderId | 
**id** | **int32** | lineitemId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcurementPurchaseordersByParentIdLineitemsRequest struct via the builder pattern


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


## DeleteProcurementPurchaseordersByParentIdLineitemsBulk

> BulkResult DeleteProcurementPurchaseordersByParentIdLineitemsBulk(ctx, parentId).ClientId(clientId).IdCollection(idCollection).Execute()

Delete BulkResult

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
	parentId := int32(56) // int32 | purchaseorderId
	clientId := "clientId_example" // string | 
	idCollection := *openapiclient.NewIdCollection() // IdCollection | purchaseOrderLineItems

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrderLineItemsAPI.DeleteProcurementPurchaseordersByParentIdLineitemsBulk(context.Background(), parentId).ClientId(clientId).IdCollection(idCollection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderLineItemsAPI.DeleteProcurementPurchaseordersByParentIdLineitemsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProcurementPurchaseordersByParentIdLineitemsBulk`: BulkResult
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderLineItemsAPI.DeleteProcurementPurchaseordersByParentIdLineitemsBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | purchaseorderId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcurementPurchaseordersByParentIdLineitemsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **idCollection** | [**IdCollection**](IdCollection.md) | purchaseOrderLineItems | 

### Return type

[**BulkResult**](BulkResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementPurchaseordersByParentIdLineitems

> []PurchaseOrderLineItem GetProcurementPurchaseordersByParentIdLineitems(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of PurchaseOrderLineItem

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
	parentId := int32(56) // int32 | purchaseorderId
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
	resp, r, err := apiClient.PurchaseOrderLineItemsAPI.GetProcurementPurchaseordersByParentIdLineitems(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderLineItemsAPI.GetProcurementPurchaseordersByParentIdLineitems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementPurchaseordersByParentIdLineitems`: []PurchaseOrderLineItem
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderLineItemsAPI.GetProcurementPurchaseordersByParentIdLineitems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | purchaseorderId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementPurchaseordersByParentIdLineitemsRequest struct via the builder pattern


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

[**[]PurchaseOrderLineItem**](PurchaseOrderLineItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementPurchaseordersByParentIdLineitemsById

> PurchaseOrderLineItem GetProcurementPurchaseordersByParentIdLineitemsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get PurchaseOrderLineItem

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
	id := int32(56) // int32 | lineitemId
	parentId := int32(56) // int32 | purchaseorderId
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
	resp, r, err := apiClient.PurchaseOrderLineItemsAPI.GetProcurementPurchaseordersByParentIdLineitemsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderLineItemsAPI.GetProcurementPurchaseordersByParentIdLineitemsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementPurchaseordersByParentIdLineitemsById`: PurchaseOrderLineItem
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderLineItemsAPI.GetProcurementPurchaseordersByParentIdLineitemsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | lineitemId | 
**parentId** | **int32** | purchaseorderId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementPurchaseordersByParentIdLineitemsByIdRequest struct via the builder pattern


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

[**PurchaseOrderLineItem**](PurchaseOrderLineItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementPurchaseordersByParentIdLineitemsCount

> Count GetProcurementPurchaseordersByParentIdLineitemsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of PurchaseOrderLineItem

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
	parentId := int32(56) // int32 | purchaseorderId
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
	resp, r, err := apiClient.PurchaseOrderLineItemsAPI.GetProcurementPurchaseordersByParentIdLineitemsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderLineItemsAPI.GetProcurementPurchaseordersByParentIdLineitemsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementPurchaseordersByParentIdLineitemsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderLineItemsAPI.GetProcurementPurchaseordersByParentIdLineitemsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | purchaseorderId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementPurchaseordersByParentIdLineitemsCountRequest struct via the builder pattern


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


## PatchProcurementPurchaseordersByParentIdLineitemsById

> PurchaseOrderLineItem PatchProcurementPurchaseordersByParentIdLineitemsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch PurchaseOrderLineItem

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
	id := int32(56) // int32 | lineitemId
	parentId := int32(56) // int32 | purchaseorderId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrderLineItemsAPI.PatchProcurementPurchaseordersByParentIdLineitemsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderLineItemsAPI.PatchProcurementPurchaseordersByParentIdLineitemsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProcurementPurchaseordersByParentIdLineitemsById`: PurchaseOrderLineItem
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderLineItemsAPI.PatchProcurementPurchaseordersByParentIdLineitemsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | lineitemId | 
**parentId** | **int32** | purchaseorderId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProcurementPurchaseordersByParentIdLineitemsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**PurchaseOrderLineItem**](PurchaseOrderLineItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProcurementPurchaseordersByParentIdLineitems

> PurchaseOrderLineItem PostProcurementPurchaseordersByParentIdLineitems(ctx, parentId).ClientId(clientId).PurchaseOrderLineItem(purchaseOrderLineItem).Execute()

Post PurchaseOrderLineItem

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
	parentId := int32(56) // int32 | purchaseorderId
	clientId := "clientId_example" // string | 
	purchaseOrderLineItem := *openapiclient.NewPurchaseOrderLineItem("Description_example", NullableInt32(123), *openapiclient.NewIvItemReference(), NullableFloat64(123), *openapiclient.NewUnitOfMeasureReference()) // PurchaseOrderLineItem | purchaseOrderLineItem

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrderLineItemsAPI.PostProcurementPurchaseordersByParentIdLineitems(context.Background(), parentId).ClientId(clientId).PurchaseOrderLineItem(purchaseOrderLineItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderLineItemsAPI.PostProcurementPurchaseordersByParentIdLineitems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementPurchaseordersByParentIdLineitems`: PurchaseOrderLineItem
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderLineItemsAPI.PostProcurementPurchaseordersByParentIdLineitems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | purchaseorderId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementPurchaseordersByParentIdLineitemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **purchaseOrderLineItem** | [**PurchaseOrderLineItem**](PurchaseOrderLineItem.md) | purchaseOrderLineItem | 

### Return type

[**PurchaseOrderLineItem**](PurchaseOrderLineItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProcurementPurchaseordersByParentIdLineitemsBulk

> BulkResult PostProcurementPurchaseordersByParentIdLineitemsBulk(ctx, parentId).ClientId(clientId).PurchaseOrderLineItem(purchaseOrderLineItem).Execute()

Post BulkResult

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
	parentId := int32(56) // int32 | purchaseorderId
	clientId := "clientId_example" // string | 
	purchaseOrderLineItem := []openapiclient.PurchaseOrderLineItem{*openapiclient.NewPurchaseOrderLineItem("Description_example", NullableInt32(123), *openapiclient.NewIvItemReference(), NullableFloat64(123), *openapiclient.NewUnitOfMeasureReference())} // []PurchaseOrderLineItem | List of PurchaseOrderLineItem

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrderLineItemsAPI.PostProcurementPurchaseordersByParentIdLineitemsBulk(context.Background(), parentId).ClientId(clientId).PurchaseOrderLineItem(purchaseOrderLineItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderLineItemsAPI.PostProcurementPurchaseordersByParentIdLineitemsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementPurchaseordersByParentIdLineitemsBulk`: BulkResult
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderLineItemsAPI.PostProcurementPurchaseordersByParentIdLineitemsBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | purchaseorderId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementPurchaseordersByParentIdLineitemsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **purchaseOrderLineItem** | [**[]PurchaseOrderLineItem**](PurchaseOrderLineItem.md) | List of PurchaseOrderLineItem | 

### Return type

[**BulkResult**](BulkResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProcurementPurchaseordersByParentIdLineitemsBulk

> BulkResult PutProcurementPurchaseordersByParentIdLineitemsBulk(ctx, parentId).ClientId(clientId).PurchaseOrderLineItem(purchaseOrderLineItem).Execute()

Put BulkResult

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
	parentId := int32(56) // int32 | purchaseorderId
	clientId := "clientId_example" // string | 
	purchaseOrderLineItem := []openapiclient.PurchaseOrderLineItem{*openapiclient.NewPurchaseOrderLineItem("Description_example", NullableInt32(123), *openapiclient.NewIvItemReference(), NullableFloat64(123), *openapiclient.NewUnitOfMeasureReference())} // []PurchaseOrderLineItem | List of PurchaseOrderLineItem

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrderLineItemsAPI.PutProcurementPurchaseordersByParentIdLineitemsBulk(context.Background(), parentId).ClientId(clientId).PurchaseOrderLineItem(purchaseOrderLineItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderLineItemsAPI.PutProcurementPurchaseordersByParentIdLineitemsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProcurementPurchaseordersByParentIdLineitemsBulk`: BulkResult
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderLineItemsAPI.PutProcurementPurchaseordersByParentIdLineitemsBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | purchaseorderId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProcurementPurchaseordersByParentIdLineitemsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **purchaseOrderLineItem** | [**[]PurchaseOrderLineItem**](PurchaseOrderLineItem.md) | List of PurchaseOrderLineItem | 

### Return type

[**BulkResult**](BulkResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProcurementPurchaseordersByParentIdLineitemsById

> PurchaseOrderLineItem PutProcurementPurchaseordersByParentIdLineitemsById(ctx, id, parentId).ClientId(clientId).PurchaseOrderLineItem(purchaseOrderLineItem).Execute()

Put PurchaseOrderLineItem

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
	id := int32(56) // int32 | lineitemId
	parentId := int32(56) // int32 | purchaseorderId
	clientId := "clientId_example" // string | 
	purchaseOrderLineItem := *openapiclient.NewPurchaseOrderLineItem("Description_example", NullableInt32(123), *openapiclient.NewIvItemReference(), NullableFloat64(123), *openapiclient.NewUnitOfMeasureReference()) // PurchaseOrderLineItem | purchaseOrderLineItem

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrderLineItemsAPI.PutProcurementPurchaseordersByParentIdLineitemsById(context.Background(), id, parentId).ClientId(clientId).PurchaseOrderLineItem(purchaseOrderLineItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderLineItemsAPI.PutProcurementPurchaseordersByParentIdLineitemsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProcurementPurchaseordersByParentIdLineitemsById`: PurchaseOrderLineItem
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderLineItemsAPI.PutProcurementPurchaseordersByParentIdLineitemsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | lineitemId | 
**parentId** | **int32** | purchaseorderId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProcurementPurchaseordersByParentIdLineitemsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **purchaseOrderLineItem** | [**PurchaseOrderLineItem**](PurchaseOrderLineItem.md) | purchaseOrderLineItem | 

### Return type

[**PurchaseOrderLineItem**](PurchaseOrderLineItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

