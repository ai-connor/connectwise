# \SalesOrdersLineItemsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSalesOrdersByParentIdLineitemsById**](SalesOrdersLineItemsAPI.md#DeleteSalesOrdersByParentIdLineitemsById) | **Delete** /sales/orders/{parentId}/lineitems/{id} | Delete SalesOrdersLineItems
[**GetSalesOrdersByParentIdLineitems**](SalesOrdersLineItemsAPI.md#GetSalesOrdersByParentIdLineitems) | **Get** /sales/orders/{parentId}/lineitems/ | Get List of SalesOrdersLineItems
[**GetSalesOrdersByParentIdLineitemsById**](SalesOrdersLineItemsAPI.md#GetSalesOrdersByParentIdLineitemsById) | **Get** /sales/orders/{parentId}/lineitems/{id} | Get SalesOrdersLineItems
[**GetSalesOrdersByParentIdLineitemsCount**](SalesOrdersLineItemsAPI.md#GetSalesOrdersByParentIdLineitemsCount) | **Get** /sales/orders/{parentId}/lineitems/count | Get Count of SalesOrdersLineItems
[**PatchSalesOrdersByParentIdLineitemsById**](SalesOrdersLineItemsAPI.md#PatchSalesOrdersByParentIdLineitemsById) | **Patch** /sales/orders/{parentId}/lineitems/{id} | Patch SalesOrdersLineItems
[**PostSalesOrdersByParentIdLineitems**](SalesOrdersLineItemsAPI.md#PostSalesOrdersByParentIdLineitems) | **Post** /sales/orders/{parentId}/lineitems/ | Post SalesOrdersLineItems
[**PutSalesOrdersByParentIdLineitemsById**](SalesOrdersLineItemsAPI.md#PutSalesOrdersByParentIdLineitemsById) | **Put** /sales/orders/{parentId}/lineitems/{id} | Put SalesOrdersLineItems



## DeleteSalesOrdersByParentIdLineitemsById

> DeleteSalesOrdersByParentIdLineitemsById(ctx, parentId, id).ClientId(clientId).Execute()

Delete SalesOrdersLineItems

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
	parentId := int32(56) // int32 | salesOrderId
	id := int32(56) // int32 | salesOrdersLineItemId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SalesOrdersLineItemsAPI.DeleteSalesOrdersByParentIdLineitemsById(context.Background(), parentId, id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesOrdersLineItemsAPI.DeleteSalesOrdersByParentIdLineitemsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | salesOrderId | 
**id** | **int32** | salesOrdersLineItemId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSalesOrdersByParentIdLineitemsByIdRequest struct via the builder pattern


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


## GetSalesOrdersByParentIdLineitems

> []SalesOrdersLineItem GetSalesOrdersByParentIdLineitems(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of SalesOrdersLineItems

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
	parentId := int32(56) // int32 | salesOrderId
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
	resp, r, err := apiClient.SalesOrdersLineItemsAPI.GetSalesOrdersByParentIdLineitems(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesOrdersLineItemsAPI.GetSalesOrdersByParentIdLineitems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesOrdersByParentIdLineitems`: []SalesOrdersLineItem
	fmt.Fprintf(os.Stdout, "Response from `SalesOrdersLineItemsAPI.GetSalesOrdersByParentIdLineitems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | salesOrderId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesOrdersByParentIdLineitemsRequest struct via the builder pattern


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

[**[]SalesOrdersLineItem**](SalesOrdersLineItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalesOrdersByParentIdLineitemsById

> SalesOrdersLineItem GetSalesOrdersByParentIdLineitemsById(ctx, parentId, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get SalesOrdersLineItems

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
	parentId := int32(56) // int32 | salesOrderId
	id := int32(56) // int32 | salesOrdersLineItemId
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
	resp, r, err := apiClient.SalesOrdersLineItemsAPI.GetSalesOrdersByParentIdLineitemsById(context.Background(), parentId, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesOrdersLineItemsAPI.GetSalesOrdersByParentIdLineitemsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesOrdersByParentIdLineitemsById`: SalesOrdersLineItem
	fmt.Fprintf(os.Stdout, "Response from `SalesOrdersLineItemsAPI.GetSalesOrdersByParentIdLineitemsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | salesOrderId | 
**id** | **int32** | salesOrdersLineItemId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesOrdersByParentIdLineitemsByIdRequest struct via the builder pattern


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

[**SalesOrdersLineItem**](SalesOrdersLineItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalesOrdersByParentIdLineitemsCount

> Count GetSalesOrdersByParentIdLineitemsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of SalesOrdersLineItems

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
	parentId := int32(56) // int32 | salesOrderId
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
	resp, r, err := apiClient.SalesOrdersLineItemsAPI.GetSalesOrdersByParentIdLineitemsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesOrdersLineItemsAPI.GetSalesOrdersByParentIdLineitemsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesOrdersByParentIdLineitemsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `SalesOrdersLineItemsAPI.GetSalesOrdersByParentIdLineitemsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | salesOrderId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesOrdersByParentIdLineitemsCountRequest struct via the builder pattern


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


## PatchSalesOrdersByParentIdLineitemsById

> SalesOrdersLineItem PatchSalesOrdersByParentIdLineitemsById(ctx, parentId, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch SalesOrdersLineItems

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
	parentId := int32(56) // int32 | salesOrderId
	id := int32(56) // int32 | salesOrdersLineItemId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalesOrdersLineItemsAPI.PatchSalesOrdersByParentIdLineitemsById(context.Background(), parentId, id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesOrdersLineItemsAPI.PatchSalesOrdersByParentIdLineitemsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSalesOrdersByParentIdLineitemsById`: SalesOrdersLineItem
	fmt.Fprintf(os.Stdout, "Response from `SalesOrdersLineItemsAPI.PatchSalesOrdersByParentIdLineitemsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | salesOrderId | 
**id** | **int32** | salesOrdersLineItemId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSalesOrdersByParentIdLineitemsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**SalesOrdersLineItem**](SalesOrdersLineItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSalesOrdersByParentIdLineitems

> SalesOrdersLineItem PostSalesOrdersByParentIdLineitems(ctx, parentId).ClientId(clientId).SalesOrdersLineItem(salesOrdersLineItem).Execute()

Post SalesOrdersLineItems

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
	parentId := int32(56) // int32 | salesOrderId
	clientId := "clientId_example" // string | 
	salesOrdersLineItem := *openapiclient.NewSalesOrdersLineItem(*openapiclient.NewSalesOrderReference()) // SalesOrdersLineItem | salesOrdersLineItem

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalesOrdersLineItemsAPI.PostSalesOrdersByParentIdLineitems(context.Background(), parentId).ClientId(clientId).SalesOrdersLineItem(salesOrdersLineItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesOrdersLineItemsAPI.PostSalesOrdersByParentIdLineitems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSalesOrdersByParentIdLineitems`: SalesOrdersLineItem
	fmt.Fprintf(os.Stdout, "Response from `SalesOrdersLineItemsAPI.PostSalesOrdersByParentIdLineitems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | salesOrderId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSalesOrdersByParentIdLineitemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **salesOrdersLineItem** | [**SalesOrdersLineItem**](SalesOrdersLineItem.md) | salesOrdersLineItem | 

### Return type

[**SalesOrdersLineItem**](SalesOrdersLineItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSalesOrdersByParentIdLineitemsById

> SalesOrdersLineItem PutSalesOrdersByParentIdLineitemsById(ctx, parentId, id).ClientId(clientId).SalesOrdersLineItem(salesOrdersLineItem).Execute()

Put SalesOrdersLineItems

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
	parentId := int32(56) // int32 | salesOrderId
	id := int32(56) // int32 | salesOrdersLineItemId
	clientId := "clientId_example" // string | 
	salesOrdersLineItem := *openapiclient.NewSalesOrdersLineItem(*openapiclient.NewSalesOrderReference()) // SalesOrdersLineItem | salesOrdersLineItem

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalesOrdersLineItemsAPI.PutSalesOrdersByParentIdLineitemsById(context.Background(), parentId, id).ClientId(clientId).SalesOrdersLineItem(salesOrdersLineItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesOrdersLineItemsAPI.PutSalesOrdersByParentIdLineitemsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSalesOrdersByParentIdLineitemsById`: SalesOrdersLineItem
	fmt.Fprintf(os.Stdout, "Response from `SalesOrdersLineItemsAPI.PutSalesOrdersByParentIdLineitemsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | salesOrderId | 
**id** | **int32** | salesOrdersLineItemId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSalesOrdersByParentIdLineitemsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **salesOrdersLineItem** | [**SalesOrdersLineItem**](SalesOrdersLineItem.md) | salesOrdersLineItem | 

### Return type

[**SalesOrdersLineItem**](SalesOrdersLineItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

