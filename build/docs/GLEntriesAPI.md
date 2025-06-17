# \GLEntriesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFinanceInvoicesByParentIdGlEntries**](GLEntriesAPI.md#GetFinanceInvoicesByParentIdGlEntries) | **Get** /finance/invoices/{parentId}/glEntries/ | Get List of GLEntries
[**GetFinanceInvoicesByParentIdGlEntriesById**](GLEntriesAPI.md#GetFinanceInvoicesByParentIdGlEntriesById) | **Get** /finance/invoices/{parentId}/glEntries/{id} | Get GLEntries
[**PatchFinanceInvoicesByParentIdGlEntriesById**](GLEntriesAPI.md#PatchFinanceInvoicesByParentIdGlEntriesById) | **Patch** /finance/invoices/{parentId}/glEntries/{id} | Patch GLEntries
[**PutFinanceInvoicesByParentIdGlEntriesById**](GLEntriesAPI.md#PutFinanceInvoicesByParentIdGlEntriesById) | **Put** /finance/invoices/{parentId}/glEntries/{id} | Put GLEntries



## GetFinanceInvoicesByParentIdGlEntries

> []GLEntry GetFinanceInvoicesByParentIdGlEntries(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of GLEntries

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
	parentId := int32(56) // int32 | invoiceId
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
	resp, r, err := apiClient.GLEntriesAPI.GetFinanceInvoicesByParentIdGlEntries(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GLEntriesAPI.GetFinanceInvoicesByParentIdGlEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceInvoicesByParentIdGlEntries`: []GLEntry
	fmt.Fprintf(os.Stdout, "Response from `GLEntriesAPI.GetFinanceInvoicesByParentIdGlEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | invoiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceInvoicesByParentIdGlEntriesRequest struct via the builder pattern


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

[**[]GLEntry**](GLEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceInvoicesByParentIdGlEntriesById

> GLEntry GetFinanceInvoicesByParentIdGlEntriesById(ctx, parentId, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get GLEntries

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
	parentId := int32(56) // int32 | invoiceId
	id := int32(56) // int32 | gLEntryId
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
	resp, r, err := apiClient.GLEntriesAPI.GetFinanceInvoicesByParentIdGlEntriesById(context.Background(), parentId, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GLEntriesAPI.GetFinanceInvoicesByParentIdGlEntriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceInvoicesByParentIdGlEntriesById`: GLEntry
	fmt.Fprintf(os.Stdout, "Response from `GLEntriesAPI.GetFinanceInvoicesByParentIdGlEntriesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | invoiceId | 
**id** | **int32** | gLEntryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceInvoicesByParentIdGlEntriesByIdRequest struct via the builder pattern


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

[**GLEntry**](GLEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFinanceInvoicesByParentIdGlEntriesById

> GLEntry PatchFinanceInvoicesByParentIdGlEntriesById(ctx, parentId, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch GLEntries

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
	parentId := int32(56) // int32 | invoiceId
	id := int32(56) // int32 | GLEntryId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GLEntriesAPI.PatchFinanceInvoicesByParentIdGlEntriesById(context.Background(), parentId, id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GLEntriesAPI.PatchFinanceInvoicesByParentIdGlEntriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceInvoicesByParentIdGlEntriesById`: GLEntry
	fmt.Fprintf(os.Stdout, "Response from `GLEntriesAPI.PatchFinanceInvoicesByParentIdGlEntriesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | invoiceId | 
**id** | **int32** | GLEntryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceInvoicesByParentIdGlEntriesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**GLEntry**](GLEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceInvoicesByParentIdGlEntriesById

> GLEntry PutFinanceInvoicesByParentIdGlEntriesById(ctx, parentId, id).ClientId(clientId).GLEntry(gLEntry).Execute()

Put GLEntries

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
	parentId := int32(56) // int32 | invoiceId
	id := int32(56) // int32 | gLEntryId
	clientId := "clientId_example" // string | 
	gLEntry := *openapiclient.NewGLEntry() // GLEntry | gLEntry

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GLEntriesAPI.PutFinanceInvoicesByParentIdGlEntriesById(context.Background(), parentId, id).ClientId(clientId).GLEntry(gLEntry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GLEntriesAPI.PutFinanceInvoicesByParentIdGlEntriesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceInvoicesByParentIdGlEntriesById`: GLEntry
	fmt.Fprintf(os.Stdout, "Response from `GLEntriesAPI.PutFinanceInvoicesByParentIdGlEntriesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | invoiceId | 
**id** | **int32** | gLEntryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceInvoicesByParentIdGlEntriesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **gLEntry** | [**GLEntry**](GLEntry.md) | gLEntry | 

### Return type

[**GLEntry**](GLEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

