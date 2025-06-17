# \OrderStatusNotificationsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSalesOrdersStatusesByParentIdNotificationsById**](OrderStatusNotificationsAPI.md#DeleteSalesOrdersStatusesByParentIdNotificationsById) | **Delete** /sales/orders/statuses/{parentId}/notifications/{id} | Delete OrderStatusNotification
[**GetSalesOrdersStatusesByParentIdNotifications**](OrderStatusNotificationsAPI.md#GetSalesOrdersStatusesByParentIdNotifications) | **Get** /sales/orders/statuses/{parentId}/notifications | Get List of OrderStatusNotification
[**GetSalesOrdersStatusesByParentIdNotificationsById**](OrderStatusNotificationsAPI.md#GetSalesOrdersStatusesByParentIdNotificationsById) | **Get** /sales/orders/statuses/{parentId}/notifications/{id} | Get OrderStatusNotification
[**GetSalesOrdersStatusesByParentIdNotificationsCount**](OrderStatusNotificationsAPI.md#GetSalesOrdersStatusesByParentIdNotificationsCount) | **Get** /sales/orders/statuses/{parentId}/notifications/count | Get Count of OrderStatusNotification
[**PatchSalesOrdersStatusesByParentIdNotificationsById**](OrderStatusNotificationsAPI.md#PatchSalesOrdersStatusesByParentIdNotificationsById) | **Patch** /sales/orders/statuses/{parentId}/notifications/{id} | Patch OrderStatusNotification
[**PostSalesOrdersStatusesByParentIdNotifications**](OrderStatusNotificationsAPI.md#PostSalesOrdersStatusesByParentIdNotifications) | **Post** /sales/orders/statuses/{parentId}/notifications | Post OrderStatusNotification
[**PutSalesOrdersStatusesByParentIdNotificationsById**](OrderStatusNotificationsAPI.md#PutSalesOrdersStatusesByParentIdNotificationsById) | **Put** /sales/orders/statuses/{parentId}/notifications/{id} | Put OrderStatusNotification



## DeleteSalesOrdersStatusesByParentIdNotificationsById

> DeleteSalesOrdersStatusesByParentIdNotificationsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete OrderStatusNotification

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
	parentId := int32(56) // int32 | statusId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrderStatusNotificationsAPI.DeleteSalesOrdersStatusesByParentIdNotificationsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderStatusNotificationsAPI.DeleteSalesOrdersStatusesByParentIdNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | notificationId | 
**parentId** | **int32** | statusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSalesOrdersStatusesByParentIdNotificationsByIdRequest struct via the builder pattern


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


## GetSalesOrdersStatusesByParentIdNotifications

> []OrderStatusNotification GetSalesOrdersStatusesByParentIdNotifications(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of OrderStatusNotification

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
	parentId := int32(56) // int32 | statusId
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
	resp, r, err := apiClient.OrderStatusNotificationsAPI.GetSalesOrdersStatusesByParentIdNotifications(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderStatusNotificationsAPI.GetSalesOrdersStatusesByParentIdNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesOrdersStatusesByParentIdNotifications`: []OrderStatusNotification
	fmt.Fprintf(os.Stdout, "Response from `OrderStatusNotificationsAPI.GetSalesOrdersStatusesByParentIdNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | statusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesOrdersStatusesByParentIdNotificationsRequest struct via the builder pattern


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

[**[]OrderStatusNotification**](OrderStatusNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalesOrdersStatusesByParentIdNotificationsById

> OrderStatusNotification GetSalesOrdersStatusesByParentIdNotificationsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get OrderStatusNotification

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
	parentId := int32(56) // int32 | statusId
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
	resp, r, err := apiClient.OrderStatusNotificationsAPI.GetSalesOrdersStatusesByParentIdNotificationsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderStatusNotificationsAPI.GetSalesOrdersStatusesByParentIdNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesOrdersStatusesByParentIdNotificationsById`: OrderStatusNotification
	fmt.Fprintf(os.Stdout, "Response from `OrderStatusNotificationsAPI.GetSalesOrdersStatusesByParentIdNotificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | notificationId | 
**parentId** | **int32** | statusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesOrdersStatusesByParentIdNotificationsByIdRequest struct via the builder pattern


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

[**OrderStatusNotification**](OrderStatusNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalesOrdersStatusesByParentIdNotificationsCount

> Count GetSalesOrdersStatusesByParentIdNotificationsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of OrderStatusNotification

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
	parentId := int32(56) // int32 | statusId
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
	resp, r, err := apiClient.OrderStatusNotificationsAPI.GetSalesOrdersStatusesByParentIdNotificationsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderStatusNotificationsAPI.GetSalesOrdersStatusesByParentIdNotificationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesOrdersStatusesByParentIdNotificationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `OrderStatusNotificationsAPI.GetSalesOrdersStatusesByParentIdNotificationsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | statusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesOrdersStatusesByParentIdNotificationsCountRequest struct via the builder pattern


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


## PatchSalesOrdersStatusesByParentIdNotificationsById

> OrderStatusNotification PatchSalesOrdersStatusesByParentIdNotificationsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch OrderStatusNotification

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
	parentId := int32(56) // int32 | statusId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderStatusNotificationsAPI.PatchSalesOrdersStatusesByParentIdNotificationsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderStatusNotificationsAPI.PatchSalesOrdersStatusesByParentIdNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSalesOrdersStatusesByParentIdNotificationsById`: OrderStatusNotification
	fmt.Fprintf(os.Stdout, "Response from `OrderStatusNotificationsAPI.PatchSalesOrdersStatusesByParentIdNotificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | notificationId | 
**parentId** | **int32** | statusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSalesOrdersStatusesByParentIdNotificationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**OrderStatusNotification**](OrderStatusNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSalesOrdersStatusesByParentIdNotifications

> OrderStatusNotification PostSalesOrdersStatusesByParentIdNotifications(ctx, parentId).ClientId(clientId).OrderStatusNotification(orderStatusNotification).Execute()

Post OrderStatusNotification

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
	parentId := int32(56) // int32 | statusId
	clientId := "clientId_example" // string | 
	orderStatusNotification := *openapiclient.NewOrderStatusNotification(*openapiclient.NewNotificationRecipientReference()) // OrderStatusNotification | orderStatusNotification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderStatusNotificationsAPI.PostSalesOrdersStatusesByParentIdNotifications(context.Background(), parentId).ClientId(clientId).OrderStatusNotification(orderStatusNotification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderStatusNotificationsAPI.PostSalesOrdersStatusesByParentIdNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSalesOrdersStatusesByParentIdNotifications`: OrderStatusNotification
	fmt.Fprintf(os.Stdout, "Response from `OrderStatusNotificationsAPI.PostSalesOrdersStatusesByParentIdNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | statusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSalesOrdersStatusesByParentIdNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **orderStatusNotification** | [**OrderStatusNotification**](OrderStatusNotification.md) | orderStatusNotification | 

### Return type

[**OrderStatusNotification**](OrderStatusNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSalesOrdersStatusesByParentIdNotificationsById

> OrderStatusNotification PutSalesOrdersStatusesByParentIdNotificationsById(ctx, id, parentId).ClientId(clientId).OrderStatusNotification(orderStatusNotification).Execute()

Put OrderStatusNotification

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
	parentId := int32(56) // int32 | statusId
	clientId := "clientId_example" // string | 
	orderStatusNotification := *openapiclient.NewOrderStatusNotification(*openapiclient.NewNotificationRecipientReference()) // OrderStatusNotification | orderStatusNotification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderStatusNotificationsAPI.PutSalesOrdersStatusesByParentIdNotificationsById(context.Background(), id, parentId).ClientId(clientId).OrderStatusNotification(orderStatusNotification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderStatusNotificationsAPI.PutSalesOrdersStatusesByParentIdNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSalesOrdersStatusesByParentIdNotificationsById`: OrderStatusNotification
	fmt.Fprintf(os.Stdout, "Response from `OrderStatusNotificationsAPI.PutSalesOrdersStatusesByParentIdNotificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | notificationId | 
**parentId** | **int32** | statusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSalesOrdersStatusesByParentIdNotificationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **orderStatusNotification** | [**OrderStatusNotification**](OrderStatusNotification.md) | orderStatusNotification | 

### Return type

[**OrderStatusNotification**](OrderStatusNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

