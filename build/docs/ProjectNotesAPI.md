# \ProjectNotesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProjectProjectsByParentIdNotesById**](ProjectNotesAPI.md#DeleteProjectProjectsByParentIdNotesById) | **Delete** /project/projects/{parentId}/notes/{id} | Delete ProjectNote
[**GetProjectProjectsByParentIdNotes**](ProjectNotesAPI.md#GetProjectProjectsByParentIdNotes) | **Get** /project/projects/{parentId}/notes | Get List of ProjectNote
[**GetProjectProjectsByParentIdNotesById**](ProjectNotesAPI.md#GetProjectProjectsByParentIdNotesById) | **Get** /project/projects/{parentId}/notes/{id} | Get ProjectNote
[**GetProjectProjectsByParentIdNotesCount**](ProjectNotesAPI.md#GetProjectProjectsByParentIdNotesCount) | **Get** /project/projects/{parentId}/notes/count | Get Count of ProjectNote
[**PatchProjectProjectsByParentIdNotesById**](ProjectNotesAPI.md#PatchProjectProjectsByParentIdNotesById) | **Patch** /project/projects/{parentId}/notes/{id} | Patch ProjectNote
[**PostProjectProjectsByParentIdNotes**](ProjectNotesAPI.md#PostProjectProjectsByParentIdNotes) | **Post** /project/projects/{parentId}/notes | Post ProjectNote
[**PutProjectProjectsByParentIdNotesById**](ProjectNotesAPI.md#PutProjectProjectsByParentIdNotesById) | **Put** /project/projects/{parentId}/notes/{id} | Put ProjectNote



## DeleteProjectProjectsByParentIdNotesById

> DeleteProjectProjectsByParentIdNotesById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ProjectNote

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
	id := int32(56) // int32 | noteId
	parentId := int32(56) // int32 | projectId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectNotesAPI.DeleteProjectProjectsByParentIdNotesById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectNotesAPI.DeleteProjectProjectsByParentIdNotesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | noteId | 
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectProjectsByParentIdNotesByIdRequest struct via the builder pattern


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


## GetProjectProjectsByParentIdNotes

> []ProjectNote GetProjectProjectsByParentIdNotes(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProjectNote

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
	parentId := int32(56) // int32 | projectId
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
	resp, r, err := apiClient.ProjectNotesAPI.GetProjectProjectsByParentIdNotes(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectNotesAPI.GetProjectProjectsByParentIdNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectsByParentIdNotes`: []ProjectNote
	fmt.Fprintf(os.Stdout, "Response from `ProjectNotesAPI.GetProjectProjectsByParentIdNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectsByParentIdNotesRequest struct via the builder pattern


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

[**[]ProjectNote**](ProjectNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectProjectsByParentIdNotesById

> ProjectNote GetProjectProjectsByParentIdNotesById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ProjectNote

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
	id := int32(56) // int32 | noteId
	parentId := int32(56) // int32 | projectId
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
	resp, r, err := apiClient.ProjectNotesAPI.GetProjectProjectsByParentIdNotesById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectNotesAPI.GetProjectProjectsByParentIdNotesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectsByParentIdNotesById`: ProjectNote
	fmt.Fprintf(os.Stdout, "Response from `ProjectNotesAPI.GetProjectProjectsByParentIdNotesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | noteId | 
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectsByParentIdNotesByIdRequest struct via the builder pattern


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

[**ProjectNote**](ProjectNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectProjectsByParentIdNotesCount

> Count GetProjectProjectsByParentIdNotesCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ProjectNote

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
	parentId := int32(56) // int32 | projectId
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
	resp, r, err := apiClient.ProjectNotesAPI.GetProjectProjectsByParentIdNotesCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectNotesAPI.GetProjectProjectsByParentIdNotesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectsByParentIdNotesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ProjectNotesAPI.GetProjectProjectsByParentIdNotesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectsByParentIdNotesCountRequest struct via the builder pattern


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


## PatchProjectProjectsByParentIdNotesById

> ProjectNote PatchProjectProjectsByParentIdNotesById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ProjectNote

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
	id := int32(56) // int32 | noteId
	parentId := int32(56) // int32 | projectId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectNotesAPI.PatchProjectProjectsByParentIdNotesById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectNotesAPI.PatchProjectProjectsByParentIdNotesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProjectProjectsByParentIdNotesById`: ProjectNote
	fmt.Fprintf(os.Stdout, "Response from `ProjectNotesAPI.PatchProjectProjectsByParentIdNotesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | noteId | 
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProjectProjectsByParentIdNotesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ProjectNote**](ProjectNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectProjectsByParentIdNotes

> ProjectNote PostProjectProjectsByParentIdNotes(ctx, parentId).ClientId(clientId).ProjectNote(projectNote).Execute()

Post ProjectNote

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
	parentId := int32(56) // int32 | projectId
	clientId := "clientId_example" // string | 
	projectNote := *openapiclient.NewProjectNote("Text_example") // ProjectNote | note

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectNotesAPI.PostProjectProjectsByParentIdNotes(context.Background(), parentId).ClientId(clientId).ProjectNote(projectNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectNotesAPI.PostProjectProjectsByParentIdNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectProjectsByParentIdNotes`: ProjectNote
	fmt.Fprintf(os.Stdout, "Response from `ProjectNotesAPI.PostProjectProjectsByParentIdNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectProjectsByParentIdNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **projectNote** | [**ProjectNote**](ProjectNote.md) | note | 

### Return type

[**ProjectNote**](ProjectNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProjectProjectsByParentIdNotesById

> ProjectNote PutProjectProjectsByParentIdNotesById(ctx, id, parentId).ClientId(clientId).ProjectNote(projectNote).Execute()

Put ProjectNote

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
	id := int32(56) // int32 | noteId
	parentId := int32(56) // int32 | projectId
	clientId := "clientId_example" // string | 
	projectNote := *openapiclient.NewProjectNote("Text_example") // ProjectNote | note

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectNotesAPI.PutProjectProjectsByParentIdNotesById(context.Background(), id, parentId).ClientId(clientId).ProjectNote(projectNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectNotesAPI.PutProjectProjectsByParentIdNotesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProjectProjectsByParentIdNotesById`: ProjectNote
	fmt.Fprintf(os.Stdout, "Response from `ProjectNotesAPI.PutProjectProjectsByParentIdNotesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | noteId | 
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProjectProjectsByParentIdNotesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **projectNote** | [**ProjectNote**](ProjectNote.md) | note | 

### Return type

[**ProjectNote**](ProjectNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

