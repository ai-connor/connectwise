# \ManagedDevicesIntegrationNotificationsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyManagedDevicesIntegrationsByParentIdNotificationsById**](ManagedDevicesIntegrationNotificationsAPI.md#DeleteCompanyManagedDevicesIntegrationsByParentIdNotificationsById) | **Delete** /company/managedDevicesIntegrations/{parentId}/notifications/{id} | Delete ManagedDevicesIntegrationNotification
[**GetCompanyManagedDevicesIntegrationsByParentIdNotifications**](ManagedDevicesIntegrationNotificationsAPI.md#GetCompanyManagedDevicesIntegrationsByParentIdNotifications) | **Get** /company/managedDevicesIntegrations/{parentId}/notifications | Get List of ManagedDevicesIntegrationNotification
[**GetCompanyManagedDevicesIntegrationsByParentIdNotificationsById**](ManagedDevicesIntegrationNotificationsAPI.md#GetCompanyManagedDevicesIntegrationsByParentIdNotificationsById) | **Get** /company/managedDevicesIntegrations/{parentId}/notifications/{id} | Get ManagedDevicesIntegrationNotification
[**GetCompanyManagedDevicesIntegrationsByParentIdNotificationsCount**](ManagedDevicesIntegrationNotificationsAPI.md#GetCompanyManagedDevicesIntegrationsByParentIdNotificationsCount) | **Get** /company/managedDevicesIntegrations/{parentId}/notifications/count | Get Count of ManagedDevicesIntegrationNotification
[**PatchCompanyManagedDevicesIntegrationsByParentIdNotificationsById**](ManagedDevicesIntegrationNotificationsAPI.md#PatchCompanyManagedDevicesIntegrationsByParentIdNotificationsById) | **Patch** /company/managedDevicesIntegrations/{parentId}/notifications/{id} | Patch ManagedDevicesIntegrationNotification
[**PostCompanyManagedDevicesIntegrationsByParentIdNotifications**](ManagedDevicesIntegrationNotificationsAPI.md#PostCompanyManagedDevicesIntegrationsByParentIdNotifications) | **Post** /company/managedDevicesIntegrations/{parentId}/notifications | Post ManagedDevicesIntegrationNotification
[**PutCompanyManagedDevicesIntegrationsByParentIdNotificationsById**](ManagedDevicesIntegrationNotificationsAPI.md#PutCompanyManagedDevicesIntegrationsByParentIdNotificationsById) | **Put** /company/managedDevicesIntegrations/{parentId}/notifications/{id} | Put ManagedDevicesIntegrationNotification



## DeleteCompanyManagedDevicesIntegrationsByParentIdNotificationsById

> ManagedDevicesIntegrationNotification DeleteCompanyManagedDevicesIntegrationsByParentIdNotificationsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ManagedDevicesIntegrationNotification

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
	parentId := int32(56) // int32 | managedDevicesIntegrationId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedDevicesIntegrationNotificationsAPI.DeleteCompanyManagedDevicesIntegrationsByParentIdNotificationsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationNotificationsAPI.DeleteCompanyManagedDevicesIntegrationsByParentIdNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCompanyManagedDevicesIntegrationsByParentIdNotificationsById`: ManagedDevicesIntegrationNotification
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationNotificationsAPI.DeleteCompanyManagedDevicesIntegrationsByParentIdNotificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | notificationId | 
**parentId** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyManagedDevicesIntegrationsByParentIdNotificationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 

### Return type

[**ManagedDevicesIntegrationNotification**](ManagedDevicesIntegrationNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyManagedDevicesIntegrationsByParentIdNotifications

> []ManagedDevicesIntegrationNotification GetCompanyManagedDevicesIntegrationsByParentIdNotifications(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ManagedDevicesIntegrationNotification

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
	parentId := int32(56) // int32 | managedDevicesIntegrationId
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
	resp, r, err := apiClient.ManagedDevicesIntegrationNotificationsAPI.GetCompanyManagedDevicesIntegrationsByParentIdNotifications(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationNotificationsAPI.GetCompanyManagedDevicesIntegrationsByParentIdNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyManagedDevicesIntegrationsByParentIdNotifications`: []ManagedDevicesIntegrationNotification
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationNotificationsAPI.GetCompanyManagedDevicesIntegrationsByParentIdNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyManagedDevicesIntegrationsByParentIdNotificationsRequest struct via the builder pattern


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

[**[]ManagedDevicesIntegrationNotification**](ManagedDevicesIntegrationNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyManagedDevicesIntegrationsByParentIdNotificationsById

> ManagedDevicesIntegrationNotification GetCompanyManagedDevicesIntegrationsByParentIdNotificationsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ManagedDevicesIntegrationNotification

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
	parentId := int32(56) // int32 | managedDevicesIntegrationId
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
	resp, r, err := apiClient.ManagedDevicesIntegrationNotificationsAPI.GetCompanyManagedDevicesIntegrationsByParentIdNotificationsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationNotificationsAPI.GetCompanyManagedDevicesIntegrationsByParentIdNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyManagedDevicesIntegrationsByParentIdNotificationsById`: ManagedDevicesIntegrationNotification
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationNotificationsAPI.GetCompanyManagedDevicesIntegrationsByParentIdNotificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | notificationId | 
**parentId** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyManagedDevicesIntegrationsByParentIdNotificationsByIdRequest struct via the builder pattern


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

[**ManagedDevicesIntegrationNotification**](ManagedDevicesIntegrationNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyManagedDevicesIntegrationsByParentIdNotificationsCount

> Count GetCompanyManagedDevicesIntegrationsByParentIdNotificationsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ManagedDevicesIntegrationNotification

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
	parentId := int32(56) // int32 | managedDevicesIntegrationId
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
	resp, r, err := apiClient.ManagedDevicesIntegrationNotificationsAPI.GetCompanyManagedDevicesIntegrationsByParentIdNotificationsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationNotificationsAPI.GetCompanyManagedDevicesIntegrationsByParentIdNotificationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyManagedDevicesIntegrationsByParentIdNotificationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationNotificationsAPI.GetCompanyManagedDevicesIntegrationsByParentIdNotificationsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyManagedDevicesIntegrationsByParentIdNotificationsCountRequest struct via the builder pattern


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


## PatchCompanyManagedDevicesIntegrationsByParentIdNotificationsById

> ManagedDevicesIntegrationNotification PatchCompanyManagedDevicesIntegrationsByParentIdNotificationsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ManagedDevicesIntegrationNotification

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
	parentId := int32(56) // int32 | managedDevicesIntegrationId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedDevicesIntegrationNotificationsAPI.PatchCompanyManagedDevicesIntegrationsByParentIdNotificationsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationNotificationsAPI.PatchCompanyManagedDevicesIntegrationsByParentIdNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyManagedDevicesIntegrationsByParentIdNotificationsById`: ManagedDevicesIntegrationNotification
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationNotificationsAPI.PatchCompanyManagedDevicesIntegrationsByParentIdNotificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | notificationId | 
**parentId** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyManagedDevicesIntegrationsByParentIdNotificationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ManagedDevicesIntegrationNotification**](ManagedDevicesIntegrationNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyManagedDevicesIntegrationsByParentIdNotifications

> ManagedDevicesIntegrationNotification PostCompanyManagedDevicesIntegrationsByParentIdNotifications(ctx, parentId).ClientId(clientId).ManagedDevicesIntegrationNotification(managedDevicesIntegrationNotification).Execute()

Post ManagedDevicesIntegrationNotification

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
	parentId := int32(56) // int32 | managedDevicesIntegrationId
	clientId := "clientId_example" // string | 
	managedDevicesIntegrationNotification := *openapiclient.NewManagedDevicesIntegrationNotification(*openapiclient.NewNotificationRecipientReference(), "LogType_example") // ManagedDevicesIntegrationNotification | notification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedDevicesIntegrationNotificationsAPI.PostCompanyManagedDevicesIntegrationsByParentIdNotifications(context.Background(), parentId).ClientId(clientId).ManagedDevicesIntegrationNotification(managedDevicesIntegrationNotification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationNotificationsAPI.PostCompanyManagedDevicesIntegrationsByParentIdNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyManagedDevicesIntegrationsByParentIdNotifications`: ManagedDevicesIntegrationNotification
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationNotificationsAPI.PostCompanyManagedDevicesIntegrationsByParentIdNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyManagedDevicesIntegrationsByParentIdNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **managedDevicesIntegrationNotification** | [**ManagedDevicesIntegrationNotification**](ManagedDevicesIntegrationNotification.md) | notification | 

### Return type

[**ManagedDevicesIntegrationNotification**](ManagedDevicesIntegrationNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyManagedDevicesIntegrationsByParentIdNotificationsById

> ManagedDevicesIntegrationNotification PutCompanyManagedDevicesIntegrationsByParentIdNotificationsById(ctx, id, parentId).ClientId(clientId).ManagedDevicesIntegrationNotification(managedDevicesIntegrationNotification).Execute()

Put ManagedDevicesIntegrationNotification

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
	parentId := int32(56) // int32 | managedDevicesIntegrationId
	clientId := "clientId_example" // string | 
	managedDevicesIntegrationNotification := *openapiclient.NewManagedDevicesIntegrationNotification(*openapiclient.NewNotificationRecipientReference(), "LogType_example") // ManagedDevicesIntegrationNotification | notification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedDevicesIntegrationNotificationsAPI.PutCompanyManagedDevicesIntegrationsByParentIdNotificationsById(context.Background(), id, parentId).ClientId(clientId).ManagedDevicesIntegrationNotification(managedDevicesIntegrationNotification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationNotificationsAPI.PutCompanyManagedDevicesIntegrationsByParentIdNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyManagedDevicesIntegrationsByParentIdNotificationsById`: ManagedDevicesIntegrationNotification
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationNotificationsAPI.PutCompanyManagedDevicesIntegrationsByParentIdNotificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | notificationId | 
**parentId** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyManagedDevicesIntegrationsByParentIdNotificationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **managedDevicesIntegrationNotification** | [**ManagedDevicesIntegrationNotification**](ManagedDevicesIntegrationNotification.md) | notification | 

### Return type

[**ManagedDevicesIntegrationNotification**](ManagedDevicesIntegrationNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

