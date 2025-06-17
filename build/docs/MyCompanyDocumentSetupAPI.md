# \MyCompanyDocumentSetupAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemMycompanyDocuments**](MyCompanyDocumentSetupAPI.md#GetSystemMycompanyDocuments) | **Get** /system/mycompany/documents | Get List of DocumentSetup
[**GetSystemMycompanyDocumentsById**](MyCompanyDocumentSetupAPI.md#GetSystemMycompanyDocumentsById) | **Get** /system/mycompany/documents/{id} | Get DocumentSetup
[**PatchSystemMycompanyDocumentsById**](MyCompanyDocumentSetupAPI.md#PatchSystemMycompanyDocumentsById) | **Patch** /system/mycompany/documents/{id} | Patch DocumentSetup
[**PutSystemMycompanyDocumentsById**](MyCompanyDocumentSetupAPI.md#PutSystemMycompanyDocumentsById) | **Put** /system/mycompany/documents/{id} | Put DocumentSetup



## GetSystemMycompanyDocuments

> []DocumentSetup GetSystemMycompanyDocuments(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of DocumentSetup

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
	resp, r, err := apiClient.MyCompanyDocumentSetupAPI.GetSystemMycompanyDocuments(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyCompanyDocumentSetupAPI.GetSystemMycompanyDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMycompanyDocuments`: []DocumentSetup
	fmt.Fprintf(os.Stdout, "Response from `MyCompanyDocumentSetupAPI.GetSystemMycompanyDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMycompanyDocumentsRequest struct via the builder pattern


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

[**[]DocumentSetup**](DocumentSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemMycompanyDocumentsById

> DocumentSetup GetSystemMycompanyDocumentsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get DocumentSetup

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
	id := int32(56) // int32 | documentId
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
	resp, r, err := apiClient.MyCompanyDocumentSetupAPI.GetSystemMycompanyDocumentsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyCompanyDocumentSetupAPI.GetSystemMycompanyDocumentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMycompanyDocumentsById`: DocumentSetup
	fmt.Fprintf(os.Stdout, "Response from `MyCompanyDocumentSetupAPI.GetSystemMycompanyDocumentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | documentId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMycompanyDocumentsByIdRequest struct via the builder pattern


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

[**DocumentSetup**](DocumentSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSystemMycompanyDocumentsById

> DocumentSetup PatchSystemMycompanyDocumentsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch DocumentSetup

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
	id := int32(56) // int32 | documentId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyCompanyDocumentSetupAPI.PatchSystemMycompanyDocumentsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyCompanyDocumentSetupAPI.PatchSystemMycompanyDocumentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemMycompanyDocumentsById`: DocumentSetup
	fmt.Fprintf(os.Stdout, "Response from `MyCompanyDocumentSetupAPI.PatchSystemMycompanyDocumentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | documentId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemMycompanyDocumentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**DocumentSetup**](DocumentSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemMycompanyDocumentsById

> DocumentSetup PutSystemMycompanyDocumentsById(ctx, id).ClientId(clientId).DocumentSetup(documentSetup).Execute()

Put DocumentSetup

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
	id := int32(56) // int32 | documentId
	clientId := "clientId_example" // string | 
	documentSetup := *openapiclient.NewDocumentSetup() // DocumentSetup | document

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyCompanyDocumentSetupAPI.PutSystemMycompanyDocumentsById(context.Background(), id).ClientId(clientId).DocumentSetup(documentSetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyCompanyDocumentSetupAPI.PutSystemMycompanyDocumentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemMycompanyDocumentsById`: DocumentSetup
	fmt.Fprintf(os.Stdout, "Response from `MyCompanyDocumentSetupAPI.PutSystemMycompanyDocumentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | documentId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemMycompanyDocumentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **documentSetup** | [**DocumentSetup**](DocumentSetup.md) | document | 

### Return type

[**DocumentSetup**](DocumentSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

