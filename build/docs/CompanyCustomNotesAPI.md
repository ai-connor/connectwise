# \CompanyCustomNotesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyCompaniesByParentIdCustomStatusNotesById**](CompanyCustomNotesAPI.md#DeleteCompanyCompaniesByParentIdCustomStatusNotesById) | **Delete** /company/companies/{parentId}/customStatusNotes/{id} | Delete CompanyCustomNote
[**GetCompanyCompaniesByParentIdCustomStatusNotes**](CompanyCustomNotesAPI.md#GetCompanyCompaniesByParentIdCustomStatusNotes) | **Get** /company/companies/{parentId}/customStatusNotes | Get List of CompanyCustomNote
[**GetCompanyCompaniesByParentIdCustomStatusNotesById**](CompanyCustomNotesAPI.md#GetCompanyCompaniesByParentIdCustomStatusNotesById) | **Get** /company/companies/{parentId}/customStatusNotes/{id} | Get CompanyCustomNote
[**GetCompanyCompaniesByParentIdCustomStatusNotesCount**](CompanyCustomNotesAPI.md#GetCompanyCompaniesByParentIdCustomStatusNotesCount) | **Get** /company/companies/{parentId}/customStatusNotes/count | Get Count of CompanyCustomNote
[**PatchCompanyCompaniesByParentIdCustomStatusNotesById**](CompanyCustomNotesAPI.md#PatchCompanyCompaniesByParentIdCustomStatusNotesById) | **Patch** /company/companies/{parentId}/customStatusNotes/{id} | Patch CompanyCustomNote
[**PostCompanyCompaniesByParentIdCustomStatusNotes**](CompanyCustomNotesAPI.md#PostCompanyCompaniesByParentIdCustomStatusNotes) | **Post** /company/companies/{parentId}/customStatusNotes | Post CompanyCustomNote
[**PutCompanyCompaniesByParentIdCustomStatusNotesById**](CompanyCustomNotesAPI.md#PutCompanyCompaniesByParentIdCustomStatusNotesById) | **Put** /company/companies/{parentId}/customStatusNotes/{id} | Put CompanyCustomNote



## DeleteCompanyCompaniesByParentIdCustomStatusNotesById

> DeleteCompanyCompaniesByParentIdCustomStatusNotesById(ctx, id, parentId).ClientId(clientId).Execute()

Delete CompanyCustomNote

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
	id := int32(56) // int32 | customStatusNoteId
	parentId := int32(56) // int32 | companyId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CompanyCustomNotesAPI.DeleteCompanyCompaniesByParentIdCustomStatusNotesById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyCustomNotesAPI.DeleteCompanyCompaniesByParentIdCustomStatusNotesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | customStatusNoteId | 
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyCompaniesByParentIdCustomStatusNotesByIdRequest struct via the builder pattern


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


## GetCompanyCompaniesByParentIdCustomStatusNotes

> []CompanyCustomNote GetCompanyCompaniesByParentIdCustomStatusNotes(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of CompanyCustomNote

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
	parentId := int32(56) // int32 | companyId
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
	resp, r, err := apiClient.CompanyCustomNotesAPI.GetCompanyCompaniesByParentIdCustomStatusNotes(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyCustomNotesAPI.GetCompanyCompaniesByParentIdCustomStatusNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCompaniesByParentIdCustomStatusNotes`: []CompanyCustomNote
	fmt.Fprintf(os.Stdout, "Response from `CompanyCustomNotesAPI.GetCompanyCompaniesByParentIdCustomStatusNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCompaniesByParentIdCustomStatusNotesRequest struct via the builder pattern


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

[**[]CompanyCustomNote**](CompanyCustomNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyCompaniesByParentIdCustomStatusNotesById

> CompanyCustomNote GetCompanyCompaniesByParentIdCustomStatusNotesById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get CompanyCustomNote

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
	id := int32(56) // int32 | customStatusNoteId
	parentId := int32(56) // int32 | companyId
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
	resp, r, err := apiClient.CompanyCustomNotesAPI.GetCompanyCompaniesByParentIdCustomStatusNotesById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyCustomNotesAPI.GetCompanyCompaniesByParentIdCustomStatusNotesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCompaniesByParentIdCustomStatusNotesById`: CompanyCustomNote
	fmt.Fprintf(os.Stdout, "Response from `CompanyCustomNotesAPI.GetCompanyCompaniesByParentIdCustomStatusNotesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | customStatusNoteId | 
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCompaniesByParentIdCustomStatusNotesByIdRequest struct via the builder pattern


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

[**CompanyCustomNote**](CompanyCustomNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyCompaniesByParentIdCustomStatusNotesCount

> Count GetCompanyCompaniesByParentIdCustomStatusNotesCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of CompanyCustomNote

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
	parentId := int32(56) // int32 | companyId
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
	resp, r, err := apiClient.CompanyCustomNotesAPI.GetCompanyCompaniesByParentIdCustomStatusNotesCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyCustomNotesAPI.GetCompanyCompaniesByParentIdCustomStatusNotesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCompaniesByParentIdCustomStatusNotesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `CompanyCustomNotesAPI.GetCompanyCompaniesByParentIdCustomStatusNotesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCompaniesByParentIdCustomStatusNotesCountRequest struct via the builder pattern


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


## PatchCompanyCompaniesByParentIdCustomStatusNotesById

> CompanyCustomNote PatchCompanyCompaniesByParentIdCustomStatusNotesById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch CompanyCustomNote

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
	id := int32(56) // int32 | customStatusNoteId
	parentId := int32(56) // int32 | companyId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyCustomNotesAPI.PatchCompanyCompaniesByParentIdCustomStatusNotesById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyCustomNotesAPI.PatchCompanyCompaniesByParentIdCustomStatusNotesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyCompaniesByParentIdCustomStatusNotesById`: CompanyCustomNote
	fmt.Fprintf(os.Stdout, "Response from `CompanyCustomNotesAPI.PatchCompanyCompaniesByParentIdCustomStatusNotesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | customStatusNoteId | 
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyCompaniesByParentIdCustomStatusNotesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**CompanyCustomNote**](CompanyCustomNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyCompaniesByParentIdCustomStatusNotes

> CompanyCustomNote PostCompanyCompaniesByParentIdCustomStatusNotes(ctx, parentId).ClientId(clientId).CompanyCustomNote(companyCustomNote).Execute()

Post CompanyCustomNote

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
	parentId := int32(56) // int32 | companyId
	clientId := "clientId_example" // string | 
	companyCustomNote := *openapiclient.NewCompanyCustomNote("CustomNote_example", *openapiclient.NewCompanyStatusReference()) // CompanyCustomNote | customNote

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyCustomNotesAPI.PostCompanyCompaniesByParentIdCustomStatusNotes(context.Background(), parentId).ClientId(clientId).CompanyCustomNote(companyCustomNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyCustomNotesAPI.PostCompanyCompaniesByParentIdCustomStatusNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyCompaniesByParentIdCustomStatusNotes`: CompanyCustomNote
	fmt.Fprintf(os.Stdout, "Response from `CompanyCustomNotesAPI.PostCompanyCompaniesByParentIdCustomStatusNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyCompaniesByParentIdCustomStatusNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **companyCustomNote** | [**CompanyCustomNote**](CompanyCustomNote.md) | customNote | 

### Return type

[**CompanyCustomNote**](CompanyCustomNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyCompaniesByParentIdCustomStatusNotesById

> CompanyCustomNote PutCompanyCompaniesByParentIdCustomStatusNotesById(ctx, id, parentId).ClientId(clientId).CompanyCustomNote(companyCustomNote).Execute()

Put CompanyCustomNote

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
	id := int32(56) // int32 | customStatusNoteId
	parentId := int32(56) // int32 | companyId
	clientId := "clientId_example" // string | 
	companyCustomNote := *openapiclient.NewCompanyCustomNote("CustomNote_example", *openapiclient.NewCompanyStatusReference()) // CompanyCustomNote | customNote

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyCustomNotesAPI.PutCompanyCompaniesByParentIdCustomStatusNotesById(context.Background(), id, parentId).ClientId(clientId).CompanyCustomNote(companyCustomNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyCustomNotesAPI.PutCompanyCompaniesByParentIdCustomStatusNotesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyCompaniesByParentIdCustomStatusNotesById`: CompanyCustomNote
	fmt.Fprintf(os.Stdout, "Response from `CompanyCustomNotesAPI.PutCompanyCompaniesByParentIdCustomStatusNotesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | customStatusNoteId | 
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyCompaniesByParentIdCustomStatusNotesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **companyCustomNote** | [**CompanyCustomNote**](CompanyCustomNote.md) | customNote | 

### Return type

[**CompanyCustomNote**](CompanyCustomNote.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

