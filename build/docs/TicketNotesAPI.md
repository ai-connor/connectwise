# \TicketNotesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProjectTicketsByParentIdNotesById**](TicketNotesAPI.md#DeleteProjectTicketsByParentIdNotesById) | **Delete** /project/tickets/{parentId}/notes/{id} | Delete TicketNote
[**DeleteServiceTicketsByParentIdNotesById**](TicketNotesAPI.md#DeleteServiceTicketsByParentIdNotesById) | **Delete** /service/tickets/{parentId}/notes/{id} | Delete ServiceNote
[**GetProjectTicketsByParentIdNotes**](TicketNotesAPI.md#GetProjectTicketsByParentIdNotes) | **Get** /project/tickets/{parentId}/notes | Get List of TicketNote
[**GetProjectTicketsByParentIdNotesById**](TicketNotesAPI.md#GetProjectTicketsByParentIdNotesById) | **Get** /project/tickets/{parentId}/notes/{id} | Get TicketNote
[**GetProjectTicketsByParentIdNotesCount**](TicketNotesAPI.md#GetProjectTicketsByParentIdNotesCount) | **Get** /project/tickets/{parentId}/notes/count | Get Count of TicketNote
[**GetServiceTicketsByParentIdNotes**](TicketNotesAPI.md#GetServiceTicketsByParentIdNotes) | **Get** /service/tickets/{parentId}/notes | Get List of ServiceNote
[**GetServiceTicketsByParentIdNotesById**](TicketNotesAPI.md#GetServiceTicketsByParentIdNotesById) | **Get** /service/tickets/{parentId}/notes/{id} | Get ServiceNote
[**GetServiceTicketsByParentIdNotesCount**](TicketNotesAPI.md#GetServiceTicketsByParentIdNotesCount) | **Get** /service/tickets/{parentId}/notes/count | Get Count of ServiceNote
[**PatchProjectTicketsByParentIdNotesById**](TicketNotesAPI.md#PatchProjectTicketsByParentIdNotesById) | **Patch** /project/tickets/{parentId}/notes/{id} | Patch TicketNote
[**PatchServiceTicketsByParentIdNotesById**](TicketNotesAPI.md#PatchServiceTicketsByParentIdNotesById) | **Patch** /service/tickets/{parentId}/notes/{id} | Patch ServiceNote
[**PostProjectTicketsByParentIdNotes**](TicketNotesAPI.md#PostProjectTicketsByParentIdNotes) | **Post** /project/tickets/{parentId}/notes | Post TicketNote
[**PostServiceTicketsByParentIdNotes**](TicketNotesAPI.md#PostServiceTicketsByParentIdNotes) | **Post** /service/tickets/{parentId}/notes | Post ServiceNote
[**PutProjectTicketsByParentIdNotesById**](TicketNotesAPI.md#PutProjectTicketsByParentIdNotesById) | **Put** /project/tickets/{parentId}/notes/{id} | Put TicketNote
[**PutServiceTicketsByParentIdNotesById**](TicketNotesAPI.md#PutServiceTicketsByParentIdNotesById) | **Put** /service/tickets/{parentId}/notes/{id} | Put ServiceNote



## DeleteProjectTicketsByParentIdNotesById

> DeleteProjectTicketsByParentIdNotesById(ctx, id, parentId).ClientId(clientId).Execute()

Delete TicketNote

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
	parentId := int32(56) // int32 | ticketId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TicketNotesAPI.DeleteProjectTicketsByParentIdNotesById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketNotesAPI.DeleteProjectTicketsByParentIdNotesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | noteId | 
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectTicketsByParentIdNotesByIdRequest struct via the builder pattern


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


## DeleteServiceTicketsByParentIdNotesById

> DeleteServiceTicketsByParentIdNotesById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ServiceNote

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
	parentId := int32(56) // int32 | ticketId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TicketNotesAPI.DeleteServiceTicketsByParentIdNotesById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketNotesAPI.DeleteServiceTicketsByParentIdNotesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | noteId | 
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceTicketsByParentIdNotesByIdRequest struct via the builder pattern


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


## GetProjectTicketsByParentIdNotes

> []TicketNote GetProjectTicketsByParentIdNotes(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of TicketNote

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
	resp, r, err := apiClient.TicketNotesAPI.GetProjectTicketsByParentIdNotes(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketNotesAPI.GetProjectTicketsByParentIdNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectTicketsByParentIdNotes`: []TicketNote
	fmt.Fprintf(os.Stdout, "Response from `TicketNotesAPI.GetProjectTicketsByParentIdNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectTicketsByParentIdNotesRequest struct via the builder pattern


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

[**[]TicketNote**](TicketNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectTicketsByParentIdNotesById

> TicketNote GetProjectTicketsByParentIdNotesById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get TicketNote

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
	resp, r, err := apiClient.TicketNotesAPI.GetProjectTicketsByParentIdNotesById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketNotesAPI.GetProjectTicketsByParentIdNotesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectTicketsByParentIdNotesById`: TicketNote
	fmt.Fprintf(os.Stdout, "Response from `TicketNotesAPI.GetProjectTicketsByParentIdNotesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | noteId | 
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectTicketsByParentIdNotesByIdRequest struct via the builder pattern


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

[**TicketNote**](TicketNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectTicketsByParentIdNotesCount

> Count GetProjectTicketsByParentIdNotesCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of TicketNote

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
	resp, r, err := apiClient.TicketNotesAPI.GetProjectTicketsByParentIdNotesCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketNotesAPI.GetProjectTicketsByParentIdNotesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectTicketsByParentIdNotesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `TicketNotesAPI.GetProjectTicketsByParentIdNotesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectTicketsByParentIdNotesCountRequest struct via the builder pattern


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


## GetServiceTicketsByParentIdNotes

> []ServiceNote GetServiceTicketsByParentIdNotes(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ServiceNote

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
	resp, r, err := apiClient.TicketNotesAPI.GetServiceTicketsByParentIdNotes(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketNotesAPI.GetServiceTicketsByParentIdNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceTicketsByParentIdNotes`: []ServiceNote
	fmt.Fprintf(os.Stdout, "Response from `TicketNotesAPI.GetServiceTicketsByParentIdNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceTicketsByParentIdNotesRequest struct via the builder pattern


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

[**[]ServiceNote**](ServiceNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceTicketsByParentIdNotesById

> ServiceNote GetServiceTicketsByParentIdNotesById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ServiceNote

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
	resp, r, err := apiClient.TicketNotesAPI.GetServiceTicketsByParentIdNotesById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketNotesAPI.GetServiceTicketsByParentIdNotesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceTicketsByParentIdNotesById`: ServiceNote
	fmt.Fprintf(os.Stdout, "Response from `TicketNotesAPI.GetServiceTicketsByParentIdNotesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | noteId | 
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceTicketsByParentIdNotesByIdRequest struct via the builder pattern


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

[**ServiceNote**](ServiceNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceTicketsByParentIdNotesCount

> Count GetServiceTicketsByParentIdNotesCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ServiceNote

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
	resp, r, err := apiClient.TicketNotesAPI.GetServiceTicketsByParentIdNotesCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketNotesAPI.GetServiceTicketsByParentIdNotesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceTicketsByParentIdNotesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `TicketNotesAPI.GetServiceTicketsByParentIdNotesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceTicketsByParentIdNotesCountRequest struct via the builder pattern


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


## PatchProjectTicketsByParentIdNotesById

> TicketNote PatchProjectTicketsByParentIdNotesById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch TicketNote

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
	parentId := int32(56) // int32 | ticketId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketNotesAPI.PatchProjectTicketsByParentIdNotesById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketNotesAPI.PatchProjectTicketsByParentIdNotesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProjectTicketsByParentIdNotesById`: TicketNote
	fmt.Fprintf(os.Stdout, "Response from `TicketNotesAPI.PatchProjectTicketsByParentIdNotesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | noteId | 
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProjectTicketsByParentIdNotesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**TicketNote**](TicketNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchServiceTicketsByParentIdNotesById

> ServiceNote PatchServiceTicketsByParentIdNotesById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ServiceNote

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
	parentId := int32(56) // int32 | ticketId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketNotesAPI.PatchServiceTicketsByParentIdNotesById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketNotesAPI.PatchServiceTicketsByParentIdNotesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchServiceTicketsByParentIdNotesById`: ServiceNote
	fmt.Fprintf(os.Stdout, "Response from `TicketNotesAPI.PatchServiceTicketsByParentIdNotesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | noteId | 
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchServiceTicketsByParentIdNotesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ServiceNote**](ServiceNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectTicketsByParentIdNotes

> TicketNote PostProjectTicketsByParentIdNotes(ctx, parentId).ClientId(clientId).TicketNote(ticketNote).Execute()

Post TicketNote

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
	ticketNote := *openapiclient.NewTicketNote() // TicketNote | ticketNote

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketNotesAPI.PostProjectTicketsByParentIdNotes(context.Background(), parentId).ClientId(clientId).TicketNote(ticketNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketNotesAPI.PostProjectTicketsByParentIdNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectTicketsByParentIdNotes`: TicketNote
	fmt.Fprintf(os.Stdout, "Response from `TicketNotesAPI.PostProjectTicketsByParentIdNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectTicketsByParentIdNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **ticketNote** | [**TicketNote**](TicketNote.md) | ticketNote | 

### Return type

[**TicketNote**](TicketNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServiceTicketsByParentIdNotes

> ServiceNote PostServiceTicketsByParentIdNotes(ctx, parentId).ClientId(clientId).ServiceNote(serviceNote).Execute()

Post ServiceNote

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
	serviceNote := *openapiclient.NewServiceNote() // ServiceNote | serviceNote

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketNotesAPI.PostServiceTicketsByParentIdNotes(context.Background(), parentId).ClientId(clientId).ServiceNote(serviceNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketNotesAPI.PostServiceTicketsByParentIdNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServiceTicketsByParentIdNotes`: ServiceNote
	fmt.Fprintf(os.Stdout, "Response from `TicketNotesAPI.PostServiceTicketsByParentIdNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServiceTicketsByParentIdNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **serviceNote** | [**ServiceNote**](ServiceNote.md) | serviceNote | 

### Return type

[**ServiceNote**](ServiceNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProjectTicketsByParentIdNotesById

> TicketNote PutProjectTicketsByParentIdNotesById(ctx, id, parentId).ClientId(clientId).TicketNote(ticketNote).Execute()

Put TicketNote

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
	parentId := int32(56) // int32 | ticketId
	clientId := "clientId_example" // string | 
	ticketNote := *openapiclient.NewTicketNote() // TicketNote | ticketNote

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketNotesAPI.PutProjectTicketsByParentIdNotesById(context.Background(), id, parentId).ClientId(clientId).TicketNote(ticketNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketNotesAPI.PutProjectTicketsByParentIdNotesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProjectTicketsByParentIdNotesById`: TicketNote
	fmt.Fprintf(os.Stdout, "Response from `TicketNotesAPI.PutProjectTicketsByParentIdNotesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | noteId | 
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProjectTicketsByParentIdNotesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **ticketNote** | [**TicketNote**](TicketNote.md) | ticketNote | 

### Return type

[**TicketNote**](TicketNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServiceTicketsByParentIdNotesById

> ServiceNote PutServiceTicketsByParentIdNotesById(ctx, id, parentId).ClientId(clientId).ServiceNote(serviceNote).Execute()

Put ServiceNote

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
	parentId := int32(56) // int32 | ticketId
	clientId := "clientId_example" // string | 
	serviceNote := *openapiclient.NewServiceNote() // ServiceNote | serviceNote

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketNotesAPI.PutServiceTicketsByParentIdNotesById(context.Background(), id, parentId).ClientId(clientId).ServiceNote(serviceNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketNotesAPI.PutServiceTicketsByParentIdNotesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServiceTicketsByParentIdNotesById`: ServiceNote
	fmt.Fprintf(os.Stdout, "Response from `TicketNotesAPI.PutServiceTicketsByParentIdNotesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | noteId | 
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServiceTicketsByParentIdNotesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **serviceNote** | [**ServiceNote**](ServiceNote.md) | serviceNote | 

### Return type

[**ServiceNote**](ServiceNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

