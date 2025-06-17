# \BoardStatusNotificationsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById**](BoardStatusNotificationsAPI.md#DeleteServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById) | **Delete** /service/boards/{grandparentId}/statuses/{parentId}/notifications/{id} | Delete BoardStatusNotification
[**GetServiceBoardsByGrandparentIdStatusesByParentIdNotifications**](BoardStatusNotificationsAPI.md#GetServiceBoardsByGrandparentIdStatusesByParentIdNotifications) | **Get** /service/boards/{grandparentId}/statuses/{parentId}/notifications | Get List of BoardStatusNotification
[**GetServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById**](BoardStatusNotificationsAPI.md#GetServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById) | **Get** /service/boards/{grandparentId}/statuses/{parentId}/notifications/{id} | Get BoardStatusNotification
[**GetServiceBoardsByGrandparentIdStatusesByParentIdNotificationsCount**](BoardStatusNotificationsAPI.md#GetServiceBoardsByGrandparentIdStatusesByParentIdNotificationsCount) | **Get** /service/boards/{grandparentId}/statuses/{parentId}/notifications/count | Get Count of BoardStatusNotification
[**PatchServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById**](BoardStatusNotificationsAPI.md#PatchServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById) | **Patch** /service/boards/{grandparentId}/statuses/{parentId}/notifications/{id} | Patch BoardStatusNotification
[**PostServiceBoardsByGrandparentIdStatusesByParentIdNotifications**](BoardStatusNotificationsAPI.md#PostServiceBoardsByGrandparentIdStatusesByParentIdNotifications) | **Post** /service/boards/{grandparentId}/statuses/{parentId}/notifications | Post BoardStatusNotification
[**PutServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById**](BoardStatusNotificationsAPI.md#PutServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById) | **Put** /service/boards/{grandparentId}/statuses/{parentId}/notifications/{id} | Put BoardStatusNotification



## DeleteServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById

> DeleteServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById(ctx, id, parentId, grandparentId).ClientId(clientId).Execute()

Delete BoardStatusNotification

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
	grandparentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BoardStatusNotificationsAPI.DeleteServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardStatusNotificationsAPI.DeleteServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById``: %v\n", err)
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
**grandparentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceBoardsByGrandparentIdStatusesByParentIdNotificationsByIdRequest struct via the builder pattern


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


## GetServiceBoardsByGrandparentIdStatusesByParentIdNotifications

> []BoardStatusNotification GetServiceBoardsByGrandparentIdStatusesByParentIdNotifications(ctx, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of BoardStatusNotification

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
	grandparentId := int32(56) // int32 | boardId
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
	resp, r, err := apiClient.BoardStatusNotificationsAPI.GetServiceBoardsByGrandparentIdStatusesByParentIdNotifications(context.Background(), parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardStatusNotificationsAPI.GetServiceBoardsByGrandparentIdStatusesByParentIdNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceBoardsByGrandparentIdStatusesByParentIdNotifications`: []BoardStatusNotification
	fmt.Fprintf(os.Stdout, "Response from `BoardStatusNotificationsAPI.GetServiceBoardsByGrandparentIdStatusesByParentIdNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | statusId | 
**grandparentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceBoardsByGrandparentIdStatusesByParentIdNotificationsRequest struct via the builder pattern


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

[**[]BoardStatusNotification**](BoardStatusNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById

> BoardStatusNotification GetServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById(ctx, id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get BoardStatusNotification

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
	grandparentId := int32(56) // int32 | boardId
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
	resp, r, err := apiClient.BoardStatusNotificationsAPI.GetServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardStatusNotificationsAPI.GetServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById`: BoardStatusNotification
	fmt.Fprintf(os.Stdout, "Response from `BoardStatusNotificationsAPI.GetServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | notificationId | 
**parentId** | **int32** | statusId | 
**grandparentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceBoardsByGrandparentIdStatusesByParentIdNotificationsByIdRequest struct via the builder pattern


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

[**BoardStatusNotification**](BoardStatusNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceBoardsByGrandparentIdStatusesByParentIdNotificationsCount

> Count GetServiceBoardsByGrandparentIdStatusesByParentIdNotificationsCount(ctx, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of BoardStatusNotification

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
	grandparentId := int32(56) // int32 | boardId
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
	resp, r, err := apiClient.BoardStatusNotificationsAPI.GetServiceBoardsByGrandparentIdStatusesByParentIdNotificationsCount(context.Background(), parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardStatusNotificationsAPI.GetServiceBoardsByGrandparentIdStatusesByParentIdNotificationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceBoardsByGrandparentIdStatusesByParentIdNotificationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `BoardStatusNotificationsAPI.GetServiceBoardsByGrandparentIdStatusesByParentIdNotificationsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | statusId | 
**grandparentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceBoardsByGrandparentIdStatusesByParentIdNotificationsCountRequest struct via the builder pattern


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


## PatchServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById

> BoardStatusNotification PatchServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById(ctx, id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch BoardStatusNotification

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
	grandparentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardStatusNotificationsAPI.PatchServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardStatusNotificationsAPI.PatchServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById`: BoardStatusNotification
	fmt.Fprintf(os.Stdout, "Response from `BoardStatusNotificationsAPI.PatchServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | notificationId | 
**parentId** | **int32** | statusId | 
**grandparentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchServiceBoardsByGrandparentIdStatusesByParentIdNotificationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**BoardStatusNotification**](BoardStatusNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServiceBoardsByGrandparentIdStatusesByParentIdNotifications

> BoardStatusNotification PostServiceBoardsByGrandparentIdStatusesByParentIdNotifications(ctx, parentId, grandparentId).ClientId(clientId).BoardStatusNotification(boardStatusNotification).Execute()

Post BoardStatusNotification

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
	grandparentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 
	boardStatusNotification := *openapiclient.NewBoardStatusNotification(*openapiclient.NewNotificationRecipientReference()) // BoardStatusNotification | boardStatusNotification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardStatusNotificationsAPI.PostServiceBoardsByGrandparentIdStatusesByParentIdNotifications(context.Background(), parentId, grandparentId).ClientId(clientId).BoardStatusNotification(boardStatusNotification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardStatusNotificationsAPI.PostServiceBoardsByGrandparentIdStatusesByParentIdNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServiceBoardsByGrandparentIdStatusesByParentIdNotifications`: BoardStatusNotification
	fmt.Fprintf(os.Stdout, "Response from `BoardStatusNotificationsAPI.PostServiceBoardsByGrandparentIdStatusesByParentIdNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | statusId | 
**grandparentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServiceBoardsByGrandparentIdStatusesByParentIdNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **boardStatusNotification** | [**BoardStatusNotification**](BoardStatusNotification.md) | boardStatusNotification | 

### Return type

[**BoardStatusNotification**](BoardStatusNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById

> BoardStatusNotification PutServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById(ctx, id, parentId, grandparentId).ClientId(clientId).BoardStatusNotification(boardStatusNotification).Execute()

Put BoardStatusNotification

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
	grandparentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 
	boardStatusNotification := *openapiclient.NewBoardStatusNotification(*openapiclient.NewNotificationRecipientReference()) // BoardStatusNotification | boardStatusNotification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardStatusNotificationsAPI.PutServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).BoardStatusNotification(boardStatusNotification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardStatusNotificationsAPI.PutServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById`: BoardStatusNotification
	fmt.Fprintf(os.Stdout, "Response from `BoardStatusNotificationsAPI.PutServiceBoardsByGrandparentIdStatusesByParentIdNotificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | notificationId | 
**parentId** | **int32** | statusId | 
**grandparentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServiceBoardsByGrandparentIdStatusesByParentIdNotificationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **boardStatusNotification** | [**BoardStatusNotification**](BoardStatusNotification.md) | boardStatusNotification | 

### Return type

[**BoardStatusNotification**](BoardStatusNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

