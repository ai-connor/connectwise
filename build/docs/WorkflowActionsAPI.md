# \WorkflowActionsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemWorkflowsByGrandparentIdEventsByParentIdActionsById**](WorkflowActionsAPI.md#DeleteSystemWorkflowsByGrandparentIdEventsByParentIdActionsById) | **Delete** /system/workflows/{grandparentId}/events/{parentId}/actions/{id} | Delete WorkflowAction
[**GetSystemWorkflowsByGrandparentIdEventsByParentIdActions**](WorkflowActionsAPI.md#GetSystemWorkflowsByGrandparentIdEventsByParentIdActions) | **Get** /system/workflows/{grandparentId}/events/{parentId}/actions | Get List of WorkflowAction
[**GetSystemWorkflowsByGrandparentIdEventsByParentIdActionsById**](WorkflowActionsAPI.md#GetSystemWorkflowsByGrandparentIdEventsByParentIdActionsById) | **Get** /system/workflows/{grandparentId}/events/{parentId}/actions/{id} | Get WorkflowAction
[**GetSystemWorkflowsByGrandparentIdEventsByParentIdActionsCount**](WorkflowActionsAPI.md#GetSystemWorkflowsByGrandparentIdEventsByParentIdActionsCount) | **Get** /system/workflows/{grandparentId}/events/{parentId}/actions/count | Get Count of WorkflowAction
[**GetSystemWorkflowsEventsActions**](WorkflowActionsAPI.md#GetSystemWorkflowsEventsActions) | **Get** /system/workflows/events/actions | Get List of WorkflowAction
[**GetSystemWorkflowsEventsActionsById**](WorkflowActionsAPI.md#GetSystemWorkflowsEventsActionsById) | **Get** /system/workflows/events/actions/{id} | Get WorkflowAction
[**PatchSystemWorkflowsByGrandparentIdEventsByParentIdActionsById**](WorkflowActionsAPI.md#PatchSystemWorkflowsByGrandparentIdEventsByParentIdActionsById) | **Patch** /system/workflows/{grandparentId}/events/{parentId}/actions/{id} | Patch WorkflowAction
[**PostSystemWorkflowsByGrandparentIdEventsByParentIdActions**](WorkflowActionsAPI.md#PostSystemWorkflowsByGrandparentIdEventsByParentIdActions) | **Post** /system/workflows/{grandparentId}/events/{parentId}/actions | Post WorkflowAction
[**PutSystemWorkflowsByGrandparentIdEventsByParentIdActionsById**](WorkflowActionsAPI.md#PutSystemWorkflowsByGrandparentIdEventsByParentIdActionsById) | **Put** /system/workflows/{grandparentId}/events/{parentId}/actions/{id} | Put WorkflowAction



## DeleteSystemWorkflowsByGrandparentIdEventsByParentIdActionsById

> DeleteSystemWorkflowsByGrandparentIdEventsByParentIdActionsById(ctx, id, parentId, grandparentId).ClientId(clientId).Execute()

Delete WorkflowAction

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
	id := int32(56) // int32 | actionId
	parentId := int32(56) // int32 | eventId
	grandparentId := int32(56) // int32 | workflowId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowActionsAPI.DeleteSystemWorkflowsByGrandparentIdEventsByParentIdActionsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.DeleteSystemWorkflowsByGrandparentIdEventsByParentIdActionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | actionId | 
**parentId** | **int32** | eventId | 
**grandparentId** | **int32** | workflowId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemWorkflowsByGrandparentIdEventsByParentIdActionsByIdRequest struct via the builder pattern


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


## GetSystemWorkflowsByGrandparentIdEventsByParentIdActions

> []WorkflowAction GetSystemWorkflowsByGrandparentIdEventsByParentIdActions(ctx, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of WorkflowAction

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
	parentId := int32(56) // int32 | eventId
	grandparentId := int32(56) // int32 | workflowId
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
	resp, r, err := apiClient.WorkflowActionsAPI.GetSystemWorkflowsByGrandparentIdEventsByParentIdActions(context.Background(), parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.GetSystemWorkflowsByGrandparentIdEventsByParentIdActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemWorkflowsByGrandparentIdEventsByParentIdActions`: []WorkflowAction
	fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.GetSystemWorkflowsByGrandparentIdEventsByParentIdActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | eventId | 
**grandparentId** | **int32** | workflowId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemWorkflowsByGrandparentIdEventsByParentIdActionsRequest struct via the builder pattern


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

[**[]WorkflowAction**](WorkflowAction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemWorkflowsByGrandparentIdEventsByParentIdActionsById

> WorkflowAction GetSystemWorkflowsByGrandparentIdEventsByParentIdActionsById(ctx, id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get WorkflowAction

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
	id := int32(56) // int32 | actionId
	parentId := int32(56) // int32 | eventId
	grandparentId := int32(56) // int32 | workflowId
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
	resp, r, err := apiClient.WorkflowActionsAPI.GetSystemWorkflowsByGrandparentIdEventsByParentIdActionsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.GetSystemWorkflowsByGrandparentIdEventsByParentIdActionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemWorkflowsByGrandparentIdEventsByParentIdActionsById`: WorkflowAction
	fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.GetSystemWorkflowsByGrandparentIdEventsByParentIdActionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | actionId | 
**parentId** | **int32** | eventId | 
**grandparentId** | **int32** | workflowId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemWorkflowsByGrandparentIdEventsByParentIdActionsByIdRequest struct via the builder pattern


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

[**WorkflowAction**](WorkflowAction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemWorkflowsByGrandparentIdEventsByParentIdActionsCount

> Count GetSystemWorkflowsByGrandparentIdEventsByParentIdActionsCount(ctx, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of WorkflowAction

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
	parentId := int32(56) // int32 | eventId
	grandparentId := int32(56) // int32 | workflowId
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
	resp, r, err := apiClient.WorkflowActionsAPI.GetSystemWorkflowsByGrandparentIdEventsByParentIdActionsCount(context.Background(), parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.GetSystemWorkflowsByGrandparentIdEventsByParentIdActionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemWorkflowsByGrandparentIdEventsByParentIdActionsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.GetSystemWorkflowsByGrandparentIdEventsByParentIdActionsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | eventId | 
**grandparentId** | **int32** | workflowId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemWorkflowsByGrandparentIdEventsByParentIdActionsCountRequest struct via the builder pattern


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


## GetSystemWorkflowsEventsActions

> []WorkflowAction GetSystemWorkflowsEventsActions(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of WorkflowAction

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
	resp, r, err := apiClient.WorkflowActionsAPI.GetSystemWorkflowsEventsActions(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.GetSystemWorkflowsEventsActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemWorkflowsEventsActions`: []WorkflowAction
	fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.GetSystemWorkflowsEventsActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemWorkflowsEventsActionsRequest struct via the builder pattern


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

[**[]WorkflowAction**](WorkflowAction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemWorkflowsEventsActionsById

> WorkflowAction GetSystemWorkflowsEventsActionsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get WorkflowAction

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
	id := int32(56) // int32 | actionId
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
	resp, r, err := apiClient.WorkflowActionsAPI.GetSystemWorkflowsEventsActionsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.GetSystemWorkflowsEventsActionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemWorkflowsEventsActionsById`: WorkflowAction
	fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.GetSystemWorkflowsEventsActionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | actionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemWorkflowsEventsActionsByIdRequest struct via the builder pattern


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

[**WorkflowAction**](WorkflowAction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSystemWorkflowsByGrandparentIdEventsByParentIdActionsById

> WorkflowAction PatchSystemWorkflowsByGrandparentIdEventsByParentIdActionsById(ctx, id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch WorkflowAction

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
	id := int32(56) // int32 | actionId
	parentId := int32(56) // int32 | eventId
	grandparentId := int32(56) // int32 | workflowId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowActionsAPI.PatchSystemWorkflowsByGrandparentIdEventsByParentIdActionsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.PatchSystemWorkflowsByGrandparentIdEventsByParentIdActionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemWorkflowsByGrandparentIdEventsByParentIdActionsById`: WorkflowAction
	fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.PatchSystemWorkflowsByGrandparentIdEventsByParentIdActionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | actionId | 
**parentId** | **int32** | eventId | 
**grandparentId** | **int32** | workflowId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemWorkflowsByGrandparentIdEventsByParentIdActionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**WorkflowAction**](WorkflowAction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemWorkflowsByGrandparentIdEventsByParentIdActions

> WorkflowAction PostSystemWorkflowsByGrandparentIdEventsByParentIdActions(ctx, parentId, grandparentId).ClientId(clientId).WorkflowAction(workflowAction).Execute()

Post WorkflowAction

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
	parentId := int32(56) // int32 | eventId
	grandparentId := int32(56) // int32 | workflowId
	clientId := "clientId_example" // string | 
	workflowAction := *openapiclient.NewWorkflowAction(*openapiclient.NewNotifyTypeReference()) // WorkflowAction | workflowAction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowActionsAPI.PostSystemWorkflowsByGrandparentIdEventsByParentIdActions(context.Background(), parentId, grandparentId).ClientId(clientId).WorkflowAction(workflowAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.PostSystemWorkflowsByGrandparentIdEventsByParentIdActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemWorkflowsByGrandparentIdEventsByParentIdActions`: WorkflowAction
	fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.PostSystemWorkflowsByGrandparentIdEventsByParentIdActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | eventId | 
**grandparentId** | **int32** | workflowId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemWorkflowsByGrandparentIdEventsByParentIdActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **workflowAction** | [**WorkflowAction**](WorkflowAction.md) | workflowAction | 

### Return type

[**WorkflowAction**](WorkflowAction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemWorkflowsByGrandparentIdEventsByParentIdActionsById

> WorkflowAction PutSystemWorkflowsByGrandparentIdEventsByParentIdActionsById(ctx, id, parentId, grandparentId).ClientId(clientId).WorkflowAction(workflowAction).Execute()

Put WorkflowAction

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
	id := int32(56) // int32 | actionId
	parentId := int32(56) // int32 | eventId
	grandparentId := int32(56) // int32 | workflowId
	clientId := "clientId_example" // string | 
	workflowAction := *openapiclient.NewWorkflowAction(*openapiclient.NewNotifyTypeReference()) // WorkflowAction | workflowAction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowActionsAPI.PutSystemWorkflowsByGrandparentIdEventsByParentIdActionsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).WorkflowAction(workflowAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionsAPI.PutSystemWorkflowsByGrandparentIdEventsByParentIdActionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemWorkflowsByGrandparentIdEventsByParentIdActionsById`: WorkflowAction
	fmt.Fprintf(os.Stdout, "Response from `WorkflowActionsAPI.PutSystemWorkflowsByGrandparentIdEventsByParentIdActionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | actionId | 
**parentId** | **int32** | eventId | 
**grandparentId** | **int32** | workflowId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemWorkflowsByGrandparentIdEventsByParentIdActionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **workflowAction** | [**WorkflowAction**](WorkflowAction.md) | workflowAction | 

### Return type

[**WorkflowAction**](WorkflowAction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

