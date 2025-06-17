# \PurchaseOrderStatusNotificationsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProcurementPurchaseorderstatusesByParentIdNotificationsById**](PurchaseOrderStatusNotificationsAPI.md#DeleteProcurementPurchaseorderstatusesByParentIdNotificationsById) | **Delete** /procurement/purchaseorderstatuses/{parentId}/notifications/{id} | Delete PurchaseOrderStatusNotification
[**GetProcurementPurchaseorderstatusesByParentIdNotifications**](PurchaseOrderStatusNotificationsAPI.md#GetProcurementPurchaseorderstatusesByParentIdNotifications) | **Get** /procurement/purchaseorderstatuses/{parentId}/notifications | Get List of PurchaseOrderStatusNotification
[**GetProcurementPurchaseorderstatusesByParentIdNotificationsById**](PurchaseOrderStatusNotificationsAPI.md#GetProcurementPurchaseorderstatusesByParentIdNotificationsById) | **Get** /procurement/purchaseorderstatuses/{parentId}/notifications/{id} | Get PurchaseOrderStatusNotification
[**GetProcurementPurchaseorderstatusesByParentIdNotificationsCount**](PurchaseOrderStatusNotificationsAPI.md#GetProcurementPurchaseorderstatusesByParentIdNotificationsCount) | **Get** /procurement/purchaseorderstatuses/{parentId}/notifications/count | Get Count of PurchaseOrderStatusNotification
[**PatchProcurementPurchaseorderstatusesByParentIdNotificationsById**](PurchaseOrderStatusNotificationsAPI.md#PatchProcurementPurchaseorderstatusesByParentIdNotificationsById) | **Patch** /procurement/purchaseorderstatuses/{parentId}/notifications/{id} | Patch PurchaseOrderStatusNotification
[**PostProcurementPurchaseorderstatusesByParentIdNotifications**](PurchaseOrderStatusNotificationsAPI.md#PostProcurementPurchaseorderstatusesByParentIdNotifications) | **Post** /procurement/purchaseorderstatuses/{parentId}/notifications | Post PurchaseOrderStatusNotification
[**PutProcurementPurchaseorderstatusesByParentIdNotificationsById**](PurchaseOrderStatusNotificationsAPI.md#PutProcurementPurchaseorderstatusesByParentIdNotificationsById) | **Put** /procurement/purchaseorderstatuses/{parentId}/notifications/{id} | Put PurchaseOrderStatusNotification



## DeleteProcurementPurchaseorderstatusesByParentIdNotificationsById

> DeleteProcurementPurchaseorderstatusesByParentIdNotificationsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete PurchaseOrderStatusNotification

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
	id := int32(56) // int32 | notificationId
	parentId := int32(56) // int32 | purchaseorderstatusId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PurchaseOrderStatusNotificationsAPI.DeleteProcurementPurchaseorderstatusesByParentIdNotificationsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderStatusNotificationsAPI.DeleteProcurementPurchaseorderstatusesByParentIdNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | notificationId | 
**parentId** | **int32** | purchaseorderstatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcurementPurchaseorderstatusesByParentIdNotificationsByIdRequest struct via the builder pattern


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


## GetProcurementPurchaseorderstatusesByParentIdNotifications

> []PurchaseOrderStatusNotification GetProcurementPurchaseorderstatusesByParentIdNotifications(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of PurchaseOrderStatusNotification

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
	parentId := int32(56) // int32 | purchaseorderstatusId
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
	resp, r, err := apiClient.PurchaseOrderStatusNotificationsAPI.GetProcurementPurchaseorderstatusesByParentIdNotifications(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderStatusNotificationsAPI.GetProcurementPurchaseorderstatusesByParentIdNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementPurchaseorderstatusesByParentIdNotifications`: []PurchaseOrderStatusNotification
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderStatusNotificationsAPI.GetProcurementPurchaseorderstatusesByParentIdNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | purchaseorderstatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementPurchaseorderstatusesByParentIdNotificationsRequest struct via the builder pattern


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

[**[]PurchaseOrderStatusNotification**](PurchaseOrderStatusNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementPurchaseorderstatusesByParentIdNotificationsById

> PurchaseOrderStatusNotification GetProcurementPurchaseorderstatusesByParentIdNotificationsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get PurchaseOrderStatusNotification

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
	id := int32(56) // int32 | notificationId
	parentId := int32(56) // int32 | purchaseorderstatusId
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
	resp, r, err := apiClient.PurchaseOrderStatusNotificationsAPI.GetProcurementPurchaseorderstatusesByParentIdNotificationsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderStatusNotificationsAPI.GetProcurementPurchaseorderstatusesByParentIdNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementPurchaseorderstatusesByParentIdNotificationsById`: PurchaseOrderStatusNotification
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderStatusNotificationsAPI.GetProcurementPurchaseorderstatusesByParentIdNotificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | notificationId | 
**parentId** | **int32** | purchaseorderstatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementPurchaseorderstatusesByParentIdNotificationsByIdRequest struct via the builder pattern


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

[**PurchaseOrderStatusNotification**](PurchaseOrderStatusNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementPurchaseorderstatusesByParentIdNotificationsCount

> Count GetProcurementPurchaseorderstatusesByParentIdNotificationsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of PurchaseOrderStatusNotification

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
	parentId := int32(56) // int32 | purchaseorderstatusId
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
	resp, r, err := apiClient.PurchaseOrderStatusNotificationsAPI.GetProcurementPurchaseorderstatusesByParentIdNotificationsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderStatusNotificationsAPI.GetProcurementPurchaseorderstatusesByParentIdNotificationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementPurchaseorderstatusesByParentIdNotificationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderStatusNotificationsAPI.GetProcurementPurchaseorderstatusesByParentIdNotificationsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | purchaseorderstatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementPurchaseorderstatusesByParentIdNotificationsCountRequest struct via the builder pattern


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


## PatchProcurementPurchaseorderstatusesByParentIdNotificationsById

> PurchaseOrderStatusNotification PatchProcurementPurchaseorderstatusesByParentIdNotificationsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch PurchaseOrderStatusNotification

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
	id := int32(56) // int32 | notificationId
	parentId := int32(56) // int32 | purchaseorderstatusId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrderStatusNotificationsAPI.PatchProcurementPurchaseorderstatusesByParentIdNotificationsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderStatusNotificationsAPI.PatchProcurementPurchaseorderstatusesByParentIdNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProcurementPurchaseorderstatusesByParentIdNotificationsById`: PurchaseOrderStatusNotification
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderStatusNotificationsAPI.PatchProcurementPurchaseorderstatusesByParentIdNotificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | notificationId | 
**parentId** | **int32** | purchaseorderstatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProcurementPurchaseorderstatusesByParentIdNotificationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**PurchaseOrderStatusNotification**](PurchaseOrderStatusNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProcurementPurchaseorderstatusesByParentIdNotifications

> PurchaseOrderStatusNotification PostProcurementPurchaseorderstatusesByParentIdNotifications(ctx, parentId).ClientId(clientId).PurchaseOrderStatusNotification(purchaseOrderStatusNotification).Execute()

Post PurchaseOrderStatusNotification

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
	parentId := int32(56) // int32 | purchaseorderstatusId
	clientId := "clientId_example" // string | 
	purchaseOrderStatusNotification := *openapiclient.NewPurchaseOrderStatusNotification(*openapiclient.NewNotificationRecipientReference()) // PurchaseOrderStatusNotification | purchaseOrderStatusNotification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrderStatusNotificationsAPI.PostProcurementPurchaseorderstatusesByParentIdNotifications(context.Background(), parentId).ClientId(clientId).PurchaseOrderStatusNotification(purchaseOrderStatusNotification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderStatusNotificationsAPI.PostProcurementPurchaseorderstatusesByParentIdNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementPurchaseorderstatusesByParentIdNotifications`: PurchaseOrderStatusNotification
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderStatusNotificationsAPI.PostProcurementPurchaseorderstatusesByParentIdNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | purchaseorderstatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementPurchaseorderstatusesByParentIdNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **purchaseOrderStatusNotification** | [**PurchaseOrderStatusNotification**](PurchaseOrderStatusNotification.md) | purchaseOrderStatusNotification | 

### Return type

[**PurchaseOrderStatusNotification**](PurchaseOrderStatusNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProcurementPurchaseorderstatusesByParentIdNotificationsById

> PurchaseOrderStatusNotification PutProcurementPurchaseorderstatusesByParentIdNotificationsById(ctx, id, parentId).ClientId(clientId).PurchaseOrderStatusNotification(purchaseOrderStatusNotification).Execute()

Put PurchaseOrderStatusNotification

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
	id := int32(56) // int32 | notificationId
	parentId := int32(56) // int32 | purchaseorderstatusId
	clientId := "clientId_example" // string | 
	purchaseOrderStatusNotification := *openapiclient.NewPurchaseOrderStatusNotification(*openapiclient.NewNotificationRecipientReference()) // PurchaseOrderStatusNotification | purchaseOrderStatusNotification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrderStatusNotificationsAPI.PutProcurementPurchaseorderstatusesByParentIdNotificationsById(context.Background(), id, parentId).ClientId(clientId).PurchaseOrderStatusNotification(purchaseOrderStatusNotification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderStatusNotificationsAPI.PutProcurementPurchaseorderstatusesByParentIdNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProcurementPurchaseorderstatusesByParentIdNotificationsById`: PurchaseOrderStatusNotification
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderStatusNotificationsAPI.PutProcurementPurchaseorderstatusesByParentIdNotificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | notificationId | 
**parentId** | **int32** | purchaseorderstatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProcurementPurchaseorderstatusesByParentIdNotificationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **purchaseOrderStatusNotification** | [**PurchaseOrderStatusNotification**](PurchaseOrderStatusNotification.md) | purchaseOrderStatusNotification | 

### Return type

[**PurchaseOrderStatusNotification**](PurchaseOrderStatusNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

