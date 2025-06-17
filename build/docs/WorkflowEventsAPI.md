# \WorkflowEventsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemWorkflowsByParentIdEventsById**](WorkflowEventsAPI.md#DeleteSystemWorkflowsByParentIdEventsById) | **Delete** /system/workflows/{parentId}/events/{id} | Delete WorkflowEvent
[**GetSystemWorkflowsByParentIdEvents**](WorkflowEventsAPI.md#GetSystemWorkflowsByParentIdEvents) | **Get** /system/workflows/{parentId}/events | Get List of WorkflowEvent
[**GetSystemWorkflowsByParentIdEventsById**](WorkflowEventsAPI.md#GetSystemWorkflowsByParentIdEventsById) | **Get** /system/workflows/{parentId}/events/{id} | Get WorkflowEvent
[**GetSystemWorkflowsByParentIdEventsByIdTest**](WorkflowEventsAPI.md#GetSystemWorkflowsByParentIdEventsByIdTest) | **Get** /system/workflows/{parentId}/events/{id}/test | Get workflow test results
[**GetSystemWorkflowsByParentIdEventsCount**](WorkflowEventsAPI.md#GetSystemWorkflowsByParentIdEventsCount) | **Get** /system/workflows/{parentId}/events/count | Get Count of WorkflowEvent
[**GetSystemWorkflowsEvents**](WorkflowEventsAPI.md#GetSystemWorkflowsEvents) | **Get** /system/workflows/events | Get List of WorkflowEvent
[**GetSystemWorkflowsEventsById**](WorkflowEventsAPI.md#GetSystemWorkflowsEventsById) | **Get** /system/workflows/events/{id} | Get WorkflowEvent
[**PatchSystemWorkflowsByParentIdEventsById**](WorkflowEventsAPI.md#PatchSystemWorkflowsByParentIdEventsById) | **Patch** /system/workflows/{parentId}/events/{id} | Patch WorkflowEvent
[**PostSystemWorkflowsByParentIdEvents**](WorkflowEventsAPI.md#PostSystemWorkflowsByParentIdEvents) | **Post** /system/workflows/{parentId}/events | Post WorkflowEvent
[**PostSystemWorkflowsByParentIdEventsByIdCopy**](WorkflowEventsAPI.md#PostSystemWorkflowsByParentIdEventsByIdCopy) | **Post** /system/workflows/{parentId}/events/{id}/copy | Post WorkflowEvent
[**PutSystemWorkflowsByParentIdEventsById**](WorkflowEventsAPI.md#PutSystemWorkflowsByParentIdEventsById) | **Put** /system/workflows/{parentId}/events/{id} | Put WorkflowEvent



## DeleteSystemWorkflowsByParentIdEventsById

> DeleteSystemWorkflowsByParentIdEventsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete WorkflowEvent

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
	id := int32(56) // int32 | eventId
	parentId := int32(56) // int32 | workflowId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowEventsAPI.DeleteSystemWorkflowsByParentIdEventsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowEventsAPI.DeleteSystemWorkflowsByParentIdEventsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | eventId | 
**parentId** | **int32** | workflowId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemWorkflowsByParentIdEventsByIdRequest struct via the builder pattern


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


## GetSystemWorkflowsByParentIdEvents

> []WorkflowEvent GetSystemWorkflowsByParentIdEvents(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of WorkflowEvent

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
	parentId := int32(56) // int32 | workflowId
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
	resp, r, err := apiClient.WorkflowEventsAPI.GetSystemWorkflowsByParentIdEvents(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowEventsAPI.GetSystemWorkflowsByParentIdEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemWorkflowsByParentIdEvents`: []WorkflowEvent
	fmt.Fprintf(os.Stdout, "Response from `WorkflowEventsAPI.GetSystemWorkflowsByParentIdEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | workflowId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemWorkflowsByParentIdEventsRequest struct via the builder pattern


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

[**[]WorkflowEvent**](WorkflowEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemWorkflowsByParentIdEventsById

> WorkflowEvent GetSystemWorkflowsByParentIdEventsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get WorkflowEvent

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
	id := int32(56) // int32 | eventId
	parentId := int32(56) // int32 | workflowId
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
	resp, r, err := apiClient.WorkflowEventsAPI.GetSystemWorkflowsByParentIdEventsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowEventsAPI.GetSystemWorkflowsByParentIdEventsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemWorkflowsByParentIdEventsById`: WorkflowEvent
	fmt.Fprintf(os.Stdout, "Response from `WorkflowEventsAPI.GetSystemWorkflowsByParentIdEventsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | eventId | 
**parentId** | **int32** | workflowId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemWorkflowsByParentIdEventsByIdRequest struct via the builder pattern


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

[**WorkflowEvent**](WorkflowEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemWorkflowsByParentIdEventsByIdTest

> []map[string]map[string]interface{} GetSystemWorkflowsByParentIdEventsByIdTest(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get workflow test results

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
	id := int32(56) // int32 | eventId
	parentId := int32(56) // int32 | workflowId
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
	resp, r, err := apiClient.WorkflowEventsAPI.GetSystemWorkflowsByParentIdEventsByIdTest(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowEventsAPI.GetSystemWorkflowsByParentIdEventsByIdTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemWorkflowsByParentIdEventsByIdTest`: []map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowEventsAPI.GetSystemWorkflowsByParentIdEventsByIdTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | eventId | 
**parentId** | **int32** | workflowId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemWorkflowsByParentIdEventsByIdTestRequest struct via the builder pattern


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

[**[]map[string]map[string]interface{}**](map.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemWorkflowsByParentIdEventsCount

> Count GetSystemWorkflowsByParentIdEventsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of WorkflowEvent

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
	parentId := int32(56) // int32 | workflowId
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
	resp, r, err := apiClient.WorkflowEventsAPI.GetSystemWorkflowsByParentIdEventsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowEventsAPI.GetSystemWorkflowsByParentIdEventsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemWorkflowsByParentIdEventsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `WorkflowEventsAPI.GetSystemWorkflowsByParentIdEventsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | workflowId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemWorkflowsByParentIdEventsCountRequest struct via the builder pattern


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


## GetSystemWorkflowsEvents

> []WorkflowEvent GetSystemWorkflowsEvents(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of WorkflowEvent

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
	resp, r, err := apiClient.WorkflowEventsAPI.GetSystemWorkflowsEvents(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowEventsAPI.GetSystemWorkflowsEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemWorkflowsEvents`: []WorkflowEvent
	fmt.Fprintf(os.Stdout, "Response from `WorkflowEventsAPI.GetSystemWorkflowsEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemWorkflowsEventsRequest struct via the builder pattern


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

[**[]WorkflowEvent**](WorkflowEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemWorkflowsEventsById

> WorkflowEvent GetSystemWorkflowsEventsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get WorkflowEvent

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
	id := int32(56) // int32 | eventId
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
	resp, r, err := apiClient.WorkflowEventsAPI.GetSystemWorkflowsEventsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowEventsAPI.GetSystemWorkflowsEventsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemWorkflowsEventsById`: WorkflowEvent
	fmt.Fprintf(os.Stdout, "Response from `WorkflowEventsAPI.GetSystemWorkflowsEventsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemWorkflowsEventsByIdRequest struct via the builder pattern


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

[**WorkflowEvent**](WorkflowEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSystemWorkflowsByParentIdEventsById

> WorkflowEvent PatchSystemWorkflowsByParentIdEventsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch WorkflowEvent

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
	id := int32(56) // int32 | eventId
	parentId := int32(56) // int32 | workflowId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowEventsAPI.PatchSystemWorkflowsByParentIdEventsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowEventsAPI.PatchSystemWorkflowsByParentIdEventsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemWorkflowsByParentIdEventsById`: WorkflowEvent
	fmt.Fprintf(os.Stdout, "Response from `WorkflowEventsAPI.PatchSystemWorkflowsByParentIdEventsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | eventId | 
**parentId** | **int32** | workflowId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemWorkflowsByParentIdEventsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**WorkflowEvent**](WorkflowEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemWorkflowsByParentIdEvents

> WorkflowEvent PostSystemWorkflowsByParentIdEvents(ctx, parentId).ClientId(clientId).WorkflowEvent(workflowEvent).Execute()

Post WorkflowEvent

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
	parentId := int32(56) // int32 | workflowId
	clientId := "clientId_example" // string | 
	workflowEvent := *openapiclient.NewWorkflowEvent("EventCondition_example") // WorkflowEvent | workflowEvent

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowEventsAPI.PostSystemWorkflowsByParentIdEvents(context.Background(), parentId).ClientId(clientId).WorkflowEvent(workflowEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowEventsAPI.PostSystemWorkflowsByParentIdEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemWorkflowsByParentIdEvents`: WorkflowEvent
	fmt.Fprintf(os.Stdout, "Response from `WorkflowEventsAPI.PostSystemWorkflowsByParentIdEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | workflowId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemWorkflowsByParentIdEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **workflowEvent** | [**WorkflowEvent**](WorkflowEvent.md) | workflowEvent | 

### Return type

[**WorkflowEvent**](WorkflowEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemWorkflowsByParentIdEventsByIdCopy

> WorkflowEvent PostSystemWorkflowsByParentIdEventsByIdCopy(ctx, id, parentId).ClientId(clientId).Execute()

Post WorkflowEvent

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
	id := int32(56) // int32 | eventId
	parentId := int32(56) // int32 | workflowId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowEventsAPI.PostSystemWorkflowsByParentIdEventsByIdCopy(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowEventsAPI.PostSystemWorkflowsByParentIdEventsByIdCopy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemWorkflowsByParentIdEventsByIdCopy`: WorkflowEvent
	fmt.Fprintf(os.Stdout, "Response from `WorkflowEventsAPI.PostSystemWorkflowsByParentIdEventsByIdCopy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | eventId | 
**parentId** | **int32** | workflowId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemWorkflowsByParentIdEventsByIdCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 

### Return type

[**WorkflowEvent**](WorkflowEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemWorkflowsByParentIdEventsById

> WorkflowEvent PutSystemWorkflowsByParentIdEventsById(ctx, id, parentId).ClientId(clientId).WorkflowEvent(workflowEvent).Execute()

Put WorkflowEvent

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
	id := int32(56) // int32 | eventId
	parentId := int32(56) // int32 | workflowId
	clientId := "clientId_example" // string | 
	workflowEvent := *openapiclient.NewWorkflowEvent("EventCondition_example") // WorkflowEvent | workflowEvent

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowEventsAPI.PutSystemWorkflowsByParentIdEventsById(context.Background(), id, parentId).ClientId(clientId).WorkflowEvent(workflowEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowEventsAPI.PutSystemWorkflowsByParentIdEventsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemWorkflowsByParentIdEventsById`: WorkflowEvent
	fmt.Fprintf(os.Stdout, "Response from `WorkflowEventsAPI.PutSystemWorkflowsByParentIdEventsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | eventId | 
**parentId** | **int32** | workflowId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemWorkflowsByParentIdEventsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **workflowEvent** | [**WorkflowEvent**](WorkflowEvent.md) | workflowEvent | 

### Return type

[**WorkflowEvent**](WorkflowEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

