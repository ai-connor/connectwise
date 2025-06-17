# \ProjectTemplateTasksAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById**](ProjectTemplateTasksAPI.md#DeleteProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById) | **Delete** /project/projectTemplates/{grandParentId}/projectTemplateTickets/{parentId}/tasks/{id} | Delete ProjectTemplateTasks
[**GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasks**](ProjectTemplateTasksAPI.md#GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasks) | **Get** /project/projectTemplates/{grandParentId}/projectTemplateTickets/{parentId}/tasks | Get List of ProjectTemplateTasks
[**GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById**](ProjectTemplateTasksAPI.md#GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById) | **Get** /project/projectTemplates/{grandParentId}/projectTemplateTickets/{parentId}/tasks/{id} | Get ProjectTemplateTasks
[**GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksCount**](ProjectTemplateTasksAPI.md#GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksCount) | **Get** /project/projectTemplates/{grandParentId}/projectTemplateTickets/{parentId}/tasks/count | Get Count of ProjectTemplateTasks
[**GetProjectProjectTemplatesProjectTemplateTicketsTasks**](ProjectTemplateTasksAPI.md#GetProjectProjectTemplatesProjectTemplateTicketsTasks) | **Get** /project/projectTemplates/projectTemplateTickets/tasks | Get List of ProjectTemplateTasks
[**PatchProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById**](ProjectTemplateTasksAPI.md#PatchProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById) | **Patch** /project/projectTemplates/{grandParentId}/projectTemplateTickets/{parentId}/tasks/{id} | Patch ProjectTemplateTasks
[**PostProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasks**](ProjectTemplateTasksAPI.md#PostProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasks) | **Post** /project/projectTemplates/{grandParentId}/projectTemplateTickets/{parentId}/tasks | Post ProjectTemplateTasks
[**PutProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById**](ProjectTemplateTasksAPI.md#PutProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById) | **Put** /project/projectTemplates/{grandParentId}/projectTemplateTickets/{parentId}/tasks/{id} | Put ProjectTemplateTasks



## DeleteProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById

> DeleteProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById(ctx, id, grandParentId, parentId).ClientId(clientId).Execute()

Delete ProjectTemplateTasks

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
	id := int32(56) // int32 | ProjectTemplateTaskId
	grandParentId := int32(56) // int32 | ProjectTemplateId
	parentId := int32(56) // int32 | ProjectTemplateTicketId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectTemplateTasksAPI.DeleteProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById(context.Background(), id, grandParentId, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplateTasksAPI.DeleteProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ProjectTemplateTaskId | 
**grandParentId** | **int32** | ProjectTemplateId | 
**parentId** | **int32** | ProjectTemplateTicketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksByIdRequest struct via the builder pattern


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


## GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasks

> []ProjectTemplateTask GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasks(ctx, grandParentId, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProjectTemplateTasks

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
	grandParentId := int32(56) // int32 | ProjectTemplateId
	parentId := int32(56) // int32 | ProjectTemplateTicketId
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
	resp, r, err := apiClient.ProjectTemplateTasksAPI.GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasks(context.Background(), grandParentId, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplateTasksAPI.GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasks`: []ProjectTemplateTask
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplateTasksAPI.GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**grandParentId** | **int32** | ProjectTemplateId | 
**parentId** | **int32** | ProjectTemplateTicketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksRequest struct via the builder pattern


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

[**[]ProjectTemplateTask**](ProjectTemplateTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById

> ProjectTemplateTask GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById(ctx, id, grandParentId, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ProjectTemplateTasks

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
	id := int32(56) // int32 | ProjectTemplateTaskId
	grandParentId := int32(56) // int32 | ProjectTemplateId
	parentId := int32(56) // int32 | ProjectTemplateTicketId
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
	resp, r, err := apiClient.ProjectTemplateTasksAPI.GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById(context.Background(), id, grandParentId, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplateTasksAPI.GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById`: ProjectTemplateTask
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplateTasksAPI.GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ProjectTemplateTaskId | 
**grandParentId** | **int32** | ProjectTemplateId | 
**parentId** | **int32** | ProjectTemplateTicketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksByIdRequest struct via the builder pattern


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

[**ProjectTemplateTask**](ProjectTemplateTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksCount

> Count GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksCount(ctx, grandParentId, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ProjectTemplateTasks

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
	grandParentId := int32(56) // int32 | ProjectTemplateId
	parentId := int32(56) // int32 | ProjectTemplateTicketId
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
	resp, r, err := apiClient.ProjectTemplateTasksAPI.GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksCount(context.Background(), grandParentId, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplateTasksAPI.GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplateTasksAPI.GetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**grandParentId** | **int32** | ProjectTemplateId | 
**parentId** | **int32** | ProjectTemplateTicketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksCountRequest struct via the builder pattern


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


## GetProjectProjectTemplatesProjectTemplateTicketsTasks

> []ProjectTemplateTask GetProjectProjectTemplatesProjectTemplateTicketsTasks(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProjectTemplateTasks

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
	resp, r, err := apiClient.ProjectTemplateTasksAPI.GetProjectProjectTemplatesProjectTemplateTicketsTasks(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplateTasksAPI.GetProjectProjectTemplatesProjectTemplateTicketsTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectTemplatesProjectTemplateTicketsTasks`: []ProjectTemplateTask
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplateTasksAPI.GetProjectProjectTemplatesProjectTemplateTicketsTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectTemplatesProjectTemplateTicketsTasksRequest struct via the builder pattern


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

[**[]ProjectTemplateTask**](ProjectTemplateTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById

> ProjectTemplateTask PatchProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById(ctx, id, grandParentId, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ProjectTemplateTasks

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
	id := int32(56) // int32 | ProjectTemplateTaskId
	grandParentId := int32(56) // int32 | ProjectTemplateId
	parentId := int32(56) // int32 | ProjectTemplateTicketId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTemplateTasksAPI.PatchProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById(context.Background(), id, grandParentId, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplateTasksAPI.PatchProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById`: ProjectTemplateTask
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplateTasksAPI.PatchProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ProjectTemplateTaskId | 
**grandParentId** | **int32** | ProjectTemplateId | 
**parentId** | **int32** | ProjectTemplateTicketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ProjectTemplateTask**](ProjectTemplateTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasks

> ProjectTemplateTask PostProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasks(ctx, grandParentId, parentId).ClientId(clientId).ProjectTemplateTask(projectTemplateTask).Execute()

Post ProjectTemplateTasks

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
	grandParentId := int32(56) // int32 | ProjectTemplateId
	parentId := int32(56) // int32 | ProjectTemplateTicketId
	clientId := "clientId_example" // string | 
	projectTemplateTask := *openapiclient.NewProjectTemplateTask() // ProjectTemplateTask | ProjectTemplateTask

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTemplateTasksAPI.PostProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasks(context.Background(), grandParentId, parentId).ClientId(clientId).ProjectTemplateTask(projectTemplateTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplateTasksAPI.PostProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasks`: ProjectTemplateTask
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplateTasksAPI.PostProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**grandParentId** | **int32** | ProjectTemplateId | 
**parentId** | **int32** | ProjectTemplateTicketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **projectTemplateTask** | [**ProjectTemplateTask**](ProjectTemplateTask.md) | ProjectTemplateTask | 

### Return type

[**ProjectTemplateTask**](ProjectTemplateTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById

> ProjectTemplateTask PutProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById(ctx, id, grandParentId, parentId).ClientId(clientId).ProjectTemplateTask(projectTemplateTask).Execute()

Put ProjectTemplateTasks

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
	id := int32(56) // int32 | ProjectTemplateTaskId
	grandParentId := int32(56) // int32 | ProjectTemplateId
	parentId := int32(56) // int32 | ProjectTemplateTicketId
	clientId := "clientId_example" // string | 
	projectTemplateTask := *openapiclient.NewProjectTemplateTask() // ProjectTemplateTask | companyTypeAssociation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTemplateTasksAPI.PutProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById(context.Background(), id, grandParentId, parentId).ClientId(clientId).ProjectTemplateTask(projectTemplateTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplateTasksAPI.PutProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById`: ProjectTemplateTask
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplateTasksAPI.PutProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ProjectTemplateTaskId | 
**grandParentId** | **int32** | ProjectTemplateId | 
**parentId** | **int32** | ProjectTemplateTicketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProjectProjectTemplatesByGrandParentIdProjectTemplateTicketsByParentIdTasksByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **projectTemplateTask** | [**ProjectTemplateTask**](ProjectTemplateTask.md) | companyTypeAssociation | 

### Return type

[**ProjectTemplateTask**](ProjectTemplateTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

