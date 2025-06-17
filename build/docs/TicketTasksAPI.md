# \TicketTasksAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProjectTicketsByParentIdTasksById**](TicketTasksAPI.md#DeleteProjectTicketsByParentIdTasksById) | **Delete** /project/tickets/{parentId}/tasks/{id} | Delete TicketTask
[**DeleteServiceTicketsByParentIdTasksById**](TicketTasksAPI.md#DeleteServiceTicketsByParentIdTasksById) | **Delete** /service/tickets/{parentId}/tasks/{id} | Delete Task
[**GetProjectTicketsByParentIdTasks**](TicketTasksAPI.md#GetProjectTicketsByParentIdTasks) | **Get** /project/tickets/{parentId}/tasks | Get List of TicketTask
[**GetProjectTicketsByParentIdTasksById**](TicketTasksAPI.md#GetProjectTicketsByParentIdTasksById) | **Get** /project/tickets/{parentId}/tasks/{id} | Get TicketTask
[**GetProjectTicketsByParentIdTasksCount**](TicketTasksAPI.md#GetProjectTicketsByParentIdTasksCount) | **Get** /project/tickets/{parentId}/tasks/count | Get Count of TicketTask
[**GetServiceTicketsByParentIdTasks**](TicketTasksAPI.md#GetServiceTicketsByParentIdTasks) | **Get** /service/tickets/{parentId}/tasks | Get List of Task
[**GetServiceTicketsByParentIdTasksById**](TicketTasksAPI.md#GetServiceTicketsByParentIdTasksById) | **Get** /service/tickets/{parentId}/tasks/{id} | Get Task
[**GetServiceTicketsByParentIdTasksCount**](TicketTasksAPI.md#GetServiceTicketsByParentIdTasksCount) | **Get** /service/tickets/{parentId}/tasks/count | Get Count of Task
[**PatchProjectTicketsByParentIdTasksById**](TicketTasksAPI.md#PatchProjectTicketsByParentIdTasksById) | **Patch** /project/tickets/{parentId}/tasks/{id} | Patch TicketTask
[**PatchServiceTicketsByParentIdTasksById**](TicketTasksAPI.md#PatchServiceTicketsByParentIdTasksById) | **Patch** /service/tickets/{parentId}/tasks/{id} | Patch Task
[**PostProjectTicketsByParentIdTasks**](TicketTasksAPI.md#PostProjectTicketsByParentIdTasks) | **Post** /project/tickets/{parentId}/tasks | Post TicketTask
[**PostServiceTicketsByParentIdTasks**](TicketTasksAPI.md#PostServiceTicketsByParentIdTasks) | **Post** /service/tickets/{parentId}/tasks | Post Task
[**PutProjectTicketsByParentIdTasksById**](TicketTasksAPI.md#PutProjectTicketsByParentIdTasksById) | **Put** /project/tickets/{parentId}/tasks/{id} | Put TicketTask
[**PutServiceTicketsByParentIdTasksById**](TicketTasksAPI.md#PutServiceTicketsByParentIdTasksById) | **Put** /service/tickets/{parentId}/tasks/{id} | Put Task



## DeleteProjectTicketsByParentIdTasksById

> DeleteProjectTicketsByParentIdTasksById(ctx, id, parentId).ClientId(clientId).Execute()

Delete TicketTask

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
	id := int32(56) // int32 | taskId
	parentId := int32(56) // int32 | ticketId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TicketTasksAPI.DeleteProjectTicketsByParentIdTasksById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketTasksAPI.DeleteProjectTicketsByParentIdTasksById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taskId | 
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectTicketsByParentIdTasksByIdRequest struct via the builder pattern


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


## DeleteServiceTicketsByParentIdTasksById

> DeleteServiceTicketsByParentIdTasksById(ctx, id, parentId).ClientId(clientId).Execute()

Delete Task

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
	id := int32(56) // int32 | taskId
	parentId := int32(56) // int32 | ticketId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TicketTasksAPI.DeleteServiceTicketsByParentIdTasksById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketTasksAPI.DeleteServiceTicketsByParentIdTasksById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taskId | 
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceTicketsByParentIdTasksByIdRequest struct via the builder pattern


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


## GetProjectTicketsByParentIdTasks

> []TicketTask GetProjectTicketsByParentIdTasks(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of TicketTask

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
	parentId := int32(56) // int32 | ticketId
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
	resp, r, err := apiClient.TicketTasksAPI.GetProjectTicketsByParentIdTasks(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketTasksAPI.GetProjectTicketsByParentIdTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectTicketsByParentIdTasks`: []TicketTask
	fmt.Fprintf(os.Stdout, "Response from `TicketTasksAPI.GetProjectTicketsByParentIdTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectTicketsByParentIdTasksRequest struct via the builder pattern


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

[**[]TicketTask**](TicketTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectTicketsByParentIdTasksById

> TicketTask GetProjectTicketsByParentIdTasksById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get TicketTask

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
	id := int32(56) // int32 | taskId
	parentId := int32(56) // int32 | ticketId
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
	resp, r, err := apiClient.TicketTasksAPI.GetProjectTicketsByParentIdTasksById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketTasksAPI.GetProjectTicketsByParentIdTasksById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectTicketsByParentIdTasksById`: TicketTask
	fmt.Fprintf(os.Stdout, "Response from `TicketTasksAPI.GetProjectTicketsByParentIdTasksById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taskId | 
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectTicketsByParentIdTasksByIdRequest struct via the builder pattern


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

[**TicketTask**](TicketTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectTicketsByParentIdTasksCount

> Count GetProjectTicketsByParentIdTasksCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of TicketTask

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
	parentId := int32(56) // int32 | ticketId
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
	resp, r, err := apiClient.TicketTasksAPI.GetProjectTicketsByParentIdTasksCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketTasksAPI.GetProjectTicketsByParentIdTasksCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectTicketsByParentIdTasksCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `TicketTasksAPI.GetProjectTicketsByParentIdTasksCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectTicketsByParentIdTasksCountRequest struct via the builder pattern


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


## GetServiceTicketsByParentIdTasks

> []ServiceTask GetServiceTicketsByParentIdTasks(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of Task

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
	parentId := int32(56) // int32 | ticketId
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
	resp, r, err := apiClient.TicketTasksAPI.GetServiceTicketsByParentIdTasks(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketTasksAPI.GetServiceTicketsByParentIdTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceTicketsByParentIdTasks`: []ServiceTask
	fmt.Fprintf(os.Stdout, "Response from `TicketTasksAPI.GetServiceTicketsByParentIdTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceTicketsByParentIdTasksRequest struct via the builder pattern


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

[**[]ServiceTask**](ServiceTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceTicketsByParentIdTasksById

> ServiceTask GetServiceTicketsByParentIdTasksById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Task

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
	id := int32(56) // int32 | taskId
	parentId := int32(56) // int32 | ticketId
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
	resp, r, err := apiClient.TicketTasksAPI.GetServiceTicketsByParentIdTasksById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketTasksAPI.GetServiceTicketsByParentIdTasksById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceTicketsByParentIdTasksById`: ServiceTask
	fmt.Fprintf(os.Stdout, "Response from `TicketTasksAPI.GetServiceTicketsByParentIdTasksById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taskId | 
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceTicketsByParentIdTasksByIdRequest struct via the builder pattern


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

[**ServiceTask**](ServiceTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceTicketsByParentIdTasksCount

> Count GetServiceTicketsByParentIdTasksCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of Task

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
	parentId := int32(56) // int32 | ticketId
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
	resp, r, err := apiClient.TicketTasksAPI.GetServiceTicketsByParentIdTasksCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketTasksAPI.GetServiceTicketsByParentIdTasksCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceTicketsByParentIdTasksCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `TicketTasksAPI.GetServiceTicketsByParentIdTasksCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceTicketsByParentIdTasksCountRequest struct via the builder pattern


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


## PatchProjectTicketsByParentIdTasksById

> TicketTask PatchProjectTicketsByParentIdTasksById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch TicketTask

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
	id := int32(56) // int32 | taskId
	parentId := int32(56) // int32 | ticketId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketTasksAPI.PatchProjectTicketsByParentIdTasksById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketTasksAPI.PatchProjectTicketsByParentIdTasksById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProjectTicketsByParentIdTasksById`: TicketTask
	fmt.Fprintf(os.Stdout, "Response from `TicketTasksAPI.PatchProjectTicketsByParentIdTasksById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taskId | 
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProjectTicketsByParentIdTasksByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**TicketTask**](TicketTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchServiceTicketsByParentIdTasksById

> ServiceTask PatchServiceTicketsByParentIdTasksById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch Task

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
	id := int32(56) // int32 | taskId
	parentId := int32(56) // int32 | ticketId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketTasksAPI.PatchServiceTicketsByParentIdTasksById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketTasksAPI.PatchServiceTicketsByParentIdTasksById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchServiceTicketsByParentIdTasksById`: ServiceTask
	fmt.Fprintf(os.Stdout, "Response from `TicketTasksAPI.PatchServiceTicketsByParentIdTasksById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taskId | 
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchServiceTicketsByParentIdTasksByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ServiceTask**](ServiceTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectTicketsByParentIdTasks

> TicketTask PostProjectTicketsByParentIdTasks(ctx, parentId).ClientId(clientId).TicketTask(ticketTask).Execute()

Post TicketTask

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
	parentId := int32(56) // int32 | ticketId
	clientId := "clientId_example" // string | 
	ticketTask := *openapiclient.NewTicketTask() // TicketTask | ticketTask

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketTasksAPI.PostProjectTicketsByParentIdTasks(context.Background(), parentId).ClientId(clientId).TicketTask(ticketTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketTasksAPI.PostProjectTicketsByParentIdTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectTicketsByParentIdTasks`: TicketTask
	fmt.Fprintf(os.Stdout, "Response from `TicketTasksAPI.PostProjectTicketsByParentIdTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectTicketsByParentIdTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **ticketTask** | [**TicketTask**](TicketTask.md) | ticketTask | 

### Return type

[**TicketTask**](TicketTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServiceTicketsByParentIdTasks

> ServiceTask PostServiceTicketsByParentIdTasks(ctx, parentId).ClientId(clientId).ServiceTask(serviceTask).Execute()

Post Task

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
	parentId := int32(56) // int32 | ticketId
	clientId := "clientId_example" // string | 
	serviceTask := *openapiclient.NewServiceTask() // ServiceTask | task

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketTasksAPI.PostServiceTicketsByParentIdTasks(context.Background(), parentId).ClientId(clientId).ServiceTask(serviceTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketTasksAPI.PostServiceTicketsByParentIdTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServiceTicketsByParentIdTasks`: ServiceTask
	fmt.Fprintf(os.Stdout, "Response from `TicketTasksAPI.PostServiceTicketsByParentIdTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServiceTicketsByParentIdTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **serviceTask** | [**ServiceTask**](ServiceTask.md) | task | 

### Return type

[**ServiceTask**](ServiceTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProjectTicketsByParentIdTasksById

> TicketTask PutProjectTicketsByParentIdTasksById(ctx, id, parentId).ClientId(clientId).TicketTask(ticketTask).Execute()

Put TicketTask

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
	id := int32(56) // int32 | taskId
	parentId := int32(56) // int32 | ticketId
	clientId := "clientId_example" // string | 
	ticketTask := *openapiclient.NewTicketTask() // TicketTask | ticketTask

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketTasksAPI.PutProjectTicketsByParentIdTasksById(context.Background(), id, parentId).ClientId(clientId).TicketTask(ticketTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketTasksAPI.PutProjectTicketsByParentIdTasksById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProjectTicketsByParentIdTasksById`: TicketTask
	fmt.Fprintf(os.Stdout, "Response from `TicketTasksAPI.PutProjectTicketsByParentIdTasksById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taskId | 
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProjectTicketsByParentIdTasksByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **ticketTask** | [**TicketTask**](TicketTask.md) | ticketTask | 

### Return type

[**TicketTask**](TicketTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServiceTicketsByParentIdTasksById

> ServiceTask PutServiceTicketsByParentIdTasksById(ctx, id, parentId).ClientId(clientId).ServiceTask(serviceTask).Execute()

Put Task

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
	id := int32(56) // int32 | taskId
	parentId := int32(56) // int32 | ticketId
	clientId := "clientId_example" // string | 
	serviceTask := *openapiclient.NewServiceTask() // ServiceTask | task

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketTasksAPI.PutServiceTicketsByParentIdTasksById(context.Background(), id, parentId).ClientId(clientId).ServiceTask(serviceTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketTasksAPI.PutServiceTicketsByParentIdTasksById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServiceTicketsByParentIdTasksById`: ServiceTask
	fmt.Fprintf(os.Stdout, "Response from `TicketTasksAPI.PutServiceTicketsByParentIdTasksById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taskId | 
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServiceTicketsByParentIdTasksByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **serviceTask** | [**ServiceTask**](ServiceTask.md) | task | 

### Return type

[**ServiceTask**](ServiceTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

