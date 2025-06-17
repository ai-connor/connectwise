# \RmaStatusNotificationsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProcurementRmaStatusesByParentIdNotificationsById**](RmaStatusNotificationsAPI.md#DeleteProcurementRmaStatusesByParentIdNotificationsById) | **Delete** /procurement/rmaStatuses/{parentId}/notifications/{id} | Delete RmaStatusNotification
[**GetProcurementRmaStatusesByParentIdNotifications**](RmaStatusNotificationsAPI.md#GetProcurementRmaStatusesByParentIdNotifications) | **Get** /procurement/rmaStatuses/{parentId}/notifications | Get List of RmaStatusNotification
[**GetProcurementRmaStatusesByParentIdNotificationsById**](RmaStatusNotificationsAPI.md#GetProcurementRmaStatusesByParentIdNotificationsById) | **Get** /procurement/rmaStatuses/{parentId}/notifications/{id} | Get RmaStatusNotification
[**GetProcurementRmaStatusesByParentIdNotificationsCount**](RmaStatusNotificationsAPI.md#GetProcurementRmaStatusesByParentIdNotificationsCount) | **Get** /procurement/rmaStatuses/{parentId}/notifications/count | Get Count of RmaStatusNotification
[**PatchProcurementRmaStatusesByParentIdNotificationsById**](RmaStatusNotificationsAPI.md#PatchProcurementRmaStatusesByParentIdNotificationsById) | **Patch** /procurement/rmaStatuses/{parentId}/notifications/{id} | Patch RmaStatusNotification
[**PostProcurementRmaStatusesByParentIdNotifications**](RmaStatusNotificationsAPI.md#PostProcurementRmaStatusesByParentIdNotifications) | **Post** /procurement/rmaStatuses/{parentId}/notifications | Post RmaStatusNotification
[**PutProcurementRmaStatusesByParentIdNotificationsById**](RmaStatusNotificationsAPI.md#PutProcurementRmaStatusesByParentIdNotificationsById) | **Put** /procurement/rmaStatuses/{parentId}/notifications/{id} | Put RmaStatusNotification



## DeleteProcurementRmaStatusesByParentIdNotificationsById

> DeleteProcurementRmaStatusesByParentIdNotificationsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete RmaStatusNotification

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
	parentId := int32(56) // int32 | rmaStatusId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RmaStatusNotificationsAPI.DeleteProcurementRmaStatusesByParentIdNotificationsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaStatusNotificationsAPI.DeleteProcurementRmaStatusesByParentIdNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | notificationId | 
**parentId** | **int32** | rmaStatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcurementRmaStatusesByParentIdNotificationsByIdRequest struct via the builder pattern


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


## GetProcurementRmaStatusesByParentIdNotifications

> []RmaStatusNotification GetProcurementRmaStatusesByParentIdNotifications(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of RmaStatusNotification

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
	parentId := int32(56) // int32 | rmaStatusId
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
	resp, r, err := apiClient.RmaStatusNotificationsAPI.GetProcurementRmaStatusesByParentIdNotifications(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaStatusNotificationsAPI.GetProcurementRmaStatusesByParentIdNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementRmaStatusesByParentIdNotifications`: []RmaStatusNotification
	fmt.Fprintf(os.Stdout, "Response from `RmaStatusNotificationsAPI.GetProcurementRmaStatusesByParentIdNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | rmaStatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementRmaStatusesByParentIdNotificationsRequest struct via the builder pattern


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

[**[]RmaStatusNotification**](RmaStatusNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementRmaStatusesByParentIdNotificationsById

> RmaStatusNotification GetProcurementRmaStatusesByParentIdNotificationsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get RmaStatusNotification

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
	parentId := int32(56) // int32 | rmaStatusId
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
	resp, r, err := apiClient.RmaStatusNotificationsAPI.GetProcurementRmaStatusesByParentIdNotificationsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaStatusNotificationsAPI.GetProcurementRmaStatusesByParentIdNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementRmaStatusesByParentIdNotificationsById`: RmaStatusNotification
	fmt.Fprintf(os.Stdout, "Response from `RmaStatusNotificationsAPI.GetProcurementRmaStatusesByParentIdNotificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | notificationId | 
**parentId** | **int32** | rmaStatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementRmaStatusesByParentIdNotificationsByIdRequest struct via the builder pattern


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

[**RmaStatusNotification**](RmaStatusNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementRmaStatusesByParentIdNotificationsCount

> Count GetProcurementRmaStatusesByParentIdNotificationsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of RmaStatusNotification

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
	parentId := int32(56) // int32 | rmaStatusId
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
	resp, r, err := apiClient.RmaStatusNotificationsAPI.GetProcurementRmaStatusesByParentIdNotificationsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaStatusNotificationsAPI.GetProcurementRmaStatusesByParentIdNotificationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementRmaStatusesByParentIdNotificationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `RmaStatusNotificationsAPI.GetProcurementRmaStatusesByParentIdNotificationsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | rmaStatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementRmaStatusesByParentIdNotificationsCountRequest struct via the builder pattern


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


## PatchProcurementRmaStatusesByParentIdNotificationsById

> RmaStatusNotification PatchProcurementRmaStatusesByParentIdNotificationsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch RmaStatusNotification

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
	parentId := int32(56) // int32 | rmaStatusId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RmaStatusNotificationsAPI.PatchProcurementRmaStatusesByParentIdNotificationsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaStatusNotificationsAPI.PatchProcurementRmaStatusesByParentIdNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProcurementRmaStatusesByParentIdNotificationsById`: RmaStatusNotification
	fmt.Fprintf(os.Stdout, "Response from `RmaStatusNotificationsAPI.PatchProcurementRmaStatusesByParentIdNotificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | notificationId | 
**parentId** | **int32** | rmaStatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProcurementRmaStatusesByParentIdNotificationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**RmaStatusNotification**](RmaStatusNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProcurementRmaStatusesByParentIdNotifications

> RmaStatusNotification PostProcurementRmaStatusesByParentIdNotifications(ctx, parentId).ClientId(clientId).RmaStatusNotification(rmaStatusNotification).Execute()

Post RmaStatusNotification

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
	parentId := int32(56) // int32 | rmaStatusId
	clientId := "clientId_example" // string | 
	rmaStatusNotification := *openapiclient.NewRmaStatusNotification(*openapiclient.NewNotificationRecipientReference()) // RmaStatusNotification | rmaStatusNotification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RmaStatusNotificationsAPI.PostProcurementRmaStatusesByParentIdNotifications(context.Background(), parentId).ClientId(clientId).RmaStatusNotification(rmaStatusNotification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaStatusNotificationsAPI.PostProcurementRmaStatusesByParentIdNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementRmaStatusesByParentIdNotifications`: RmaStatusNotification
	fmt.Fprintf(os.Stdout, "Response from `RmaStatusNotificationsAPI.PostProcurementRmaStatusesByParentIdNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | rmaStatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementRmaStatusesByParentIdNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **rmaStatusNotification** | [**RmaStatusNotification**](RmaStatusNotification.md) | rmaStatusNotification | 

### Return type

[**RmaStatusNotification**](RmaStatusNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProcurementRmaStatusesByParentIdNotificationsById

> RmaStatusNotification PutProcurementRmaStatusesByParentIdNotificationsById(ctx, id, parentId).ClientId(clientId).RmaStatusNotification(rmaStatusNotification).Execute()

Put RmaStatusNotification

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
	parentId := int32(56) // int32 | rmaStatusId
	clientId := "clientId_example" // string | 
	rmaStatusNotification := *openapiclient.NewRmaStatusNotification(*openapiclient.NewNotificationRecipientReference()) // RmaStatusNotification | rmaStatusNotification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RmaStatusNotificationsAPI.PutProcurementRmaStatusesByParentIdNotificationsById(context.Background(), id, parentId).ClientId(clientId).RmaStatusNotification(rmaStatusNotification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaStatusNotificationsAPI.PutProcurementRmaStatusesByParentIdNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProcurementRmaStatusesByParentIdNotificationsById`: RmaStatusNotification
	fmt.Fprintf(os.Stdout, "Response from `RmaStatusNotificationsAPI.PutProcurementRmaStatusesByParentIdNotificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | notificationId | 
**parentId** | **int32** | rmaStatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProcurementRmaStatusesByParentIdNotificationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **rmaStatusNotification** | [**RmaStatusNotification**](RmaStatusNotification.md) | rmaStatusNotification | 

### Return type

[**RmaStatusNotification**](RmaStatusNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

