# \ProjectBoardKanbanSettingsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProjectBoardsByParentIdKanbanSettingsById**](ProjectBoardKanbanSettingsAPI.md#DeleteProjectBoardsByParentIdKanbanSettingsById) | **Delete** /project/boards/{parentId}/kanbanSettings/{id} | Delete ProjectBoardKanbanSetting
[**GetProjectBoardsByParentIdKanbanSettings**](ProjectBoardKanbanSettingsAPI.md#GetProjectBoardsByParentIdKanbanSettings) | **Get** /project/boards/{parentId}/kanbanSettings | Get List of ProjectBoardKanbanSettings
[**GetProjectBoardsByParentIdKanbanSettingsById**](ProjectBoardKanbanSettingsAPI.md#GetProjectBoardsByParentIdKanbanSettingsById) | **Get** /project/boards/{parentId}/kanbanSettings/{id} | Get ProjectBoardKanbanSetting
[**PatchProjectBoardsByParentIdKanbanSettingsById**](ProjectBoardKanbanSettingsAPI.md#PatchProjectBoardsByParentIdKanbanSettingsById) | **Patch** /project/boards/{parentId}/kanbanSettings/{id} | Patch ProjectBoardKanbanSetting
[**PostProjectBoardsByParentIdKanbanSettings**](ProjectBoardKanbanSettingsAPI.md#PostProjectBoardsByParentIdKanbanSettings) | **Post** /project/boards/{parentId}/kanbanSettings | Post ProjectBoardKanbanSetting
[**PutProjectBoardsByParentIdKanbanSettingsById**](ProjectBoardKanbanSettingsAPI.md#PutProjectBoardsByParentIdKanbanSettingsById) | **Put** /project/boards/{parentId}/kanbanSettings/{id} | Put ProjectBoardKanbanSetting



## DeleteProjectBoardsByParentIdKanbanSettingsById

> DeleteProjectBoardsByParentIdKanbanSettingsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ProjectBoardKanbanSetting

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
	id := int32(56) // int32 | KanbanId
	parentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectBoardKanbanSettingsAPI.DeleteProjectBoardsByParentIdKanbanSettingsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBoardKanbanSettingsAPI.DeleteProjectBoardsByParentIdKanbanSettingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | KanbanId | 
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectBoardsByParentIdKanbanSettingsByIdRequest struct via the builder pattern


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


## GetProjectBoardsByParentIdKanbanSettings

> []ProjectBoardKanbanSetting GetProjectBoardsByParentIdKanbanSettings(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProjectBoardKanbanSettings

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
	parentId := int32(56) // int32 | boardId
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
	resp, r, err := apiClient.ProjectBoardKanbanSettingsAPI.GetProjectBoardsByParentIdKanbanSettings(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBoardKanbanSettingsAPI.GetProjectBoardsByParentIdKanbanSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectBoardsByParentIdKanbanSettings`: []ProjectBoardKanbanSetting
	fmt.Fprintf(os.Stdout, "Response from `ProjectBoardKanbanSettingsAPI.GetProjectBoardsByParentIdKanbanSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectBoardsByParentIdKanbanSettingsRequest struct via the builder pattern


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

[**[]ProjectBoardKanbanSetting**](ProjectBoardKanbanSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectBoardsByParentIdKanbanSettingsById

> ProjectBoardKanbanSetting GetProjectBoardsByParentIdKanbanSettingsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ProjectBoardKanbanSetting

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
	id := int32(56) // int32 | KanbanId
	parentId := int32(56) // int32 | boardId
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
	resp, r, err := apiClient.ProjectBoardKanbanSettingsAPI.GetProjectBoardsByParentIdKanbanSettingsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBoardKanbanSettingsAPI.GetProjectBoardsByParentIdKanbanSettingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectBoardsByParentIdKanbanSettingsById`: ProjectBoardKanbanSetting
	fmt.Fprintf(os.Stdout, "Response from `ProjectBoardKanbanSettingsAPI.GetProjectBoardsByParentIdKanbanSettingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | KanbanId | 
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectBoardsByParentIdKanbanSettingsByIdRequest struct via the builder pattern


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

[**ProjectBoardKanbanSetting**](ProjectBoardKanbanSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchProjectBoardsByParentIdKanbanSettingsById

> ProjectBoardKanbanSetting PatchProjectBoardsByParentIdKanbanSettingsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ProjectBoardKanbanSetting

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
	id := int32(56) // int32 | KanbanId
	parentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectBoardKanbanSettingsAPI.PatchProjectBoardsByParentIdKanbanSettingsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBoardKanbanSettingsAPI.PatchProjectBoardsByParentIdKanbanSettingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProjectBoardsByParentIdKanbanSettingsById`: ProjectBoardKanbanSetting
	fmt.Fprintf(os.Stdout, "Response from `ProjectBoardKanbanSettingsAPI.PatchProjectBoardsByParentIdKanbanSettingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | KanbanId | 
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProjectBoardsByParentIdKanbanSettingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ProjectBoardKanbanSetting**](ProjectBoardKanbanSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectBoardsByParentIdKanbanSettings

> ProjectBoardKanbanSetting PostProjectBoardsByParentIdKanbanSettings(ctx, parentId).ClientId(clientId).ProjectBoardKanbanSetting(projectBoardKanbanSetting).Execute()

Post ProjectBoardKanbanSetting

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
	parentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 
	projectBoardKanbanSetting := *openapiclient.NewProjectBoardKanbanSetting("Name_example") // ProjectBoardKanbanSetting | kanbanSettings

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectBoardKanbanSettingsAPI.PostProjectBoardsByParentIdKanbanSettings(context.Background(), parentId).ClientId(clientId).ProjectBoardKanbanSetting(projectBoardKanbanSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBoardKanbanSettingsAPI.PostProjectBoardsByParentIdKanbanSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectBoardsByParentIdKanbanSettings`: ProjectBoardKanbanSetting
	fmt.Fprintf(os.Stdout, "Response from `ProjectBoardKanbanSettingsAPI.PostProjectBoardsByParentIdKanbanSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectBoardsByParentIdKanbanSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **projectBoardKanbanSetting** | [**ProjectBoardKanbanSetting**](ProjectBoardKanbanSetting.md) | kanbanSettings | 

### Return type

[**ProjectBoardKanbanSetting**](ProjectBoardKanbanSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProjectBoardsByParentIdKanbanSettingsById

> ProjectBoardKanbanSetting PutProjectBoardsByParentIdKanbanSettingsById(ctx, id, parentId).ClientId(clientId).ProjectBoardKanbanSetting(projectBoardKanbanSetting).Execute()

Put ProjectBoardKanbanSetting

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
	id := int32(56) // int32 | KanbanId
	parentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 
	projectBoardKanbanSetting := *openapiclient.NewProjectBoardKanbanSetting("Name_example") // ProjectBoardKanbanSetting | Kanban

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectBoardKanbanSettingsAPI.PutProjectBoardsByParentIdKanbanSettingsById(context.Background(), id, parentId).ClientId(clientId).ProjectBoardKanbanSetting(projectBoardKanbanSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBoardKanbanSettingsAPI.PutProjectBoardsByParentIdKanbanSettingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProjectBoardsByParentIdKanbanSettingsById`: ProjectBoardKanbanSetting
	fmt.Fprintf(os.Stdout, "Response from `ProjectBoardKanbanSettingsAPI.PutProjectBoardsByParentIdKanbanSettingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | KanbanId | 
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProjectBoardsByParentIdKanbanSettingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **projectBoardKanbanSetting** | [**ProjectBoardKanbanSetting**](ProjectBoardKanbanSetting.md) | Kanban | 

### Return type

[**ProjectBoardKanbanSetting**](ProjectBoardKanbanSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

