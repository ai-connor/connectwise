# \ProjectTicketNotesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProjectTicketsByParentIdAllNotes**](ProjectTicketNotesAPI.md#GetProjectTicketsByParentIdAllNotes) | **Get** /project/tickets/{parentId}/allNotes | Get List of ProjectTicketNote
[**PostProjectTicketNoteByIdMarkAs**](ProjectTicketNotesAPI.md#PostProjectTicketNoteByIdMarkAs) | **Post** /project/ticketNote/{id}/markAs/ | Post ProjectTicketNote



## GetProjectTicketsByParentIdAllNotes

> []ProjectTicketNote GetProjectTicketsByParentIdAllNotes(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProjectTicketNote

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
	resp, r, err := apiClient.ProjectTicketNotesAPI.GetProjectTicketsByParentIdAllNotes(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTicketNotesAPI.GetProjectTicketsByParentIdAllNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectTicketsByParentIdAllNotes`: []ProjectTicketNote
	fmt.Fprintf(os.Stdout, "Response from `ProjectTicketNotesAPI.GetProjectTicketsByParentIdAllNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | ticketId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectTicketsByParentIdAllNotesRequest struct via the builder pattern


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

[**[]ProjectTicketNote**](ProjectTicketNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectTicketNoteByIdMarkAs

> PostProjectTicketNoteByIdMarkAs(ctx, id).ClientId(clientId).ProjectTicketNote(projectTicketNote).Execute()

Post ProjectTicketNote

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
	id := int32(56) // int32 | ticketNoteId
	clientId := "clientId_example" // string | 
	projectTicketNote := *openapiclient.NewProjectTicketNote() // ProjectTicketNote | item

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectTicketNotesAPI.PostProjectTicketNoteByIdMarkAs(context.Background(), id).ClientId(clientId).ProjectTicketNote(projectTicketNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTicketNotesAPI.PostProjectTicketNoteByIdMarkAs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ticketNoteId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectTicketNoteByIdMarkAsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **projectTicketNote** | [**ProjectTicketNote**](ProjectTicketNote.md) | item | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

