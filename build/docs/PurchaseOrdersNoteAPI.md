# \PurchaseOrdersNoteAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProcurementPurchaseordersByParentIdNotesById**](PurchaseOrdersNoteAPI.md#DeleteProcurementPurchaseordersByParentIdNotesById) | **Delete** /procurement/purchaseorders/{parentId}/notes/{id} | Delete PurchaseOrderNote
[**GetProcurementPurchaseordersByParentIdNotesById**](PurchaseOrdersNoteAPI.md#GetProcurementPurchaseordersByParentIdNotesById) | **Get** /procurement/purchaseorders/{parentId}/notes/{id} | Get PurchaseOrderNote
[**GetProcurementPurchaseordersByParentIdNotesCount**](PurchaseOrdersNoteAPI.md#GetProcurementPurchaseordersByParentIdNotesCount) | **Get** /procurement/purchaseorders/{parentId}/notes/count | Get Count of PurchaseOrdersNote
[**PatchProcurementPurchaseordersByParentIdNotesById**](PurchaseOrdersNoteAPI.md#PatchProcurementPurchaseordersByParentIdNotesById) | **Patch** /procurement/purchaseorders/{parentId}/notes/{id} | Patch PurchaseOrderNote
[**PostProcurementPurchaseordersByParentIdNotes**](PurchaseOrdersNoteAPI.md#PostProcurementPurchaseordersByParentIdNotes) | **Post** /procurement/purchaseorders/{parentId}/notes | Post PurchaseOrderNote
[**PutProcurementPurchaseordersByParentIdNotesById**](PurchaseOrdersNoteAPI.md#PutProcurementPurchaseordersByParentIdNotesById) | **Put** /procurement/purchaseorders/{parentId}/notes/{id} | Put PurchaseOrderNote



## DeleteProcurementPurchaseordersByParentIdNotesById

> DeleteProcurementPurchaseordersByParentIdNotesById(ctx, id, parentId).ClientId(clientId).Execute()

Delete PurchaseOrderNote

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
	parentId := int32(56) // int32 | PurchaseHeaderRecID
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PurchaseOrdersNoteAPI.DeleteProcurementPurchaseordersByParentIdNotesById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrdersNoteAPI.DeleteProcurementPurchaseordersByParentIdNotesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | noteId | 
**parentId** | **int32** | PurchaseHeaderRecID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcurementPurchaseordersByParentIdNotesByIdRequest struct via the builder pattern


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


## GetProcurementPurchaseordersByParentIdNotesById

> PurchaseOrderNote GetProcurementPurchaseordersByParentIdNotesById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get PurchaseOrderNote

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
	parentId := int32(56) // int32 | PurchaseHeaderRecID
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
	resp, r, err := apiClient.PurchaseOrdersNoteAPI.GetProcurementPurchaseordersByParentIdNotesById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrdersNoteAPI.GetProcurementPurchaseordersByParentIdNotesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementPurchaseordersByParentIdNotesById`: PurchaseOrderNote
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrdersNoteAPI.GetProcurementPurchaseordersByParentIdNotesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | noteId | 
**parentId** | **int32** | PurchaseHeaderRecID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementPurchaseordersByParentIdNotesByIdRequest struct via the builder pattern


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

[**PurchaseOrderNote**](PurchaseOrderNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementPurchaseordersByParentIdNotesCount

> Count GetProcurementPurchaseordersByParentIdNotesCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of PurchaseOrdersNote

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
	parentId := int32(56) // int32 | PurchaseHeaderRecID
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
	resp, r, err := apiClient.PurchaseOrdersNoteAPI.GetProcurementPurchaseordersByParentIdNotesCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrdersNoteAPI.GetProcurementPurchaseordersByParentIdNotesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementPurchaseordersByParentIdNotesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrdersNoteAPI.GetProcurementPurchaseordersByParentIdNotesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | PurchaseHeaderRecID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementPurchaseordersByParentIdNotesCountRequest struct via the builder pattern


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


## PatchProcurementPurchaseordersByParentIdNotesById

> PurchaseOrderNote PatchProcurementPurchaseordersByParentIdNotesById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch PurchaseOrderNote

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
	parentId := int32(56) // int32 | PurchaseHeaderRecID
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrdersNoteAPI.PatchProcurementPurchaseordersByParentIdNotesById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrdersNoteAPI.PatchProcurementPurchaseordersByParentIdNotesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProcurementPurchaseordersByParentIdNotesById`: PurchaseOrderNote
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrdersNoteAPI.PatchProcurementPurchaseordersByParentIdNotesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | noteId | 
**parentId** | **int32** | PurchaseHeaderRecID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProcurementPurchaseordersByParentIdNotesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**PurchaseOrderNote**](PurchaseOrderNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProcurementPurchaseordersByParentIdNotes

> PurchaseOrderNote PostProcurementPurchaseordersByParentIdNotes(ctx, parentId).ClientId(clientId).PurchaseOrderNote(purchaseOrderNote).Execute()

Post PurchaseOrderNote

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
	parentId := int32(56) // int32 | PurchaseHeaderRecID
	clientId := "clientId_example" // string | 
	purchaseOrderNote := *openapiclient.NewPurchaseOrderNote() // PurchaseOrderNote | PurchaseOrderNote

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrdersNoteAPI.PostProcurementPurchaseordersByParentIdNotes(context.Background(), parentId).ClientId(clientId).PurchaseOrderNote(purchaseOrderNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrdersNoteAPI.PostProcurementPurchaseordersByParentIdNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementPurchaseordersByParentIdNotes`: PurchaseOrderNote
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrdersNoteAPI.PostProcurementPurchaseordersByParentIdNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | PurchaseHeaderRecID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementPurchaseordersByParentIdNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **purchaseOrderNote** | [**PurchaseOrderNote**](PurchaseOrderNote.md) | PurchaseOrderNote | 

### Return type

[**PurchaseOrderNote**](PurchaseOrderNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProcurementPurchaseordersByParentIdNotesById

> PurchaseOrderNote PutProcurementPurchaseordersByParentIdNotesById(ctx, id, parentId).ClientId(clientId).PurchaseOrderNote(purchaseOrderNote).Execute()

Put PurchaseOrderNote

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
	parentId := int32(56) // int32 | PurchaseHeaderRecID
	clientId := "clientId_example" // string | 
	purchaseOrderNote := *openapiclient.NewPurchaseOrderNote() // PurchaseOrderNote | PurchaseOrderNote

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrdersNoteAPI.PutProcurementPurchaseordersByParentIdNotesById(context.Background(), id, parentId).ClientId(clientId).PurchaseOrderNote(purchaseOrderNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrdersNoteAPI.PutProcurementPurchaseordersByParentIdNotesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProcurementPurchaseordersByParentIdNotesById`: PurchaseOrderNote
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrdersNoteAPI.PutProcurementPurchaseordersByParentIdNotesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | noteId | 
**parentId** | **int32** | PurchaseHeaderRecID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProcurementPurchaseordersByParentIdNotesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **purchaseOrderNote** | [**PurchaseOrderNote**](PurchaseOrderNote.md) | PurchaseOrderNote | 

### Return type

[**PurchaseOrderNote**](PurchaseOrderNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

