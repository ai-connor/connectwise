# \ContactContactTypeAssociationsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyContactsByParentIdTypeAssociationsById**](ContactContactTypeAssociationsAPI.md#DeleteCompanyContactsByParentIdTypeAssociationsById) | **Delete** /company/contacts/{parentId}/typeAssociations/{id} | Delete ContactTypeAssociation
[**GetCompanyContactsByParentIdTypeAssociations**](ContactContactTypeAssociationsAPI.md#GetCompanyContactsByParentIdTypeAssociations) | **Get** /company/contacts/{parentId}/typeAssociations | Get List of ContactTypeAssociation
[**GetCompanyContactsByParentIdTypeAssociationsById**](ContactContactTypeAssociationsAPI.md#GetCompanyContactsByParentIdTypeAssociationsById) | **Get** /company/contacts/{parentId}/typeAssociations/{id} | Get ContactTypeAssociation
[**GetCompanyContactsByParentIdTypeAssociationsCount**](ContactContactTypeAssociationsAPI.md#GetCompanyContactsByParentIdTypeAssociationsCount) | **Get** /company/contacts/{parentId}/typeAssociations/count | Get Count of ContactTypeAssociation
[**PatchCompanyContactsByParentIdTypeAssociationsById**](ContactContactTypeAssociationsAPI.md#PatchCompanyContactsByParentIdTypeAssociationsById) | **Patch** /company/contacts/{parentId}/typeAssociations/{id} | Patch ContactTypeAssociation
[**PostCompanyContactsByParentIdTypeAssociations**](ContactContactTypeAssociationsAPI.md#PostCompanyContactsByParentIdTypeAssociations) | **Post** /company/contacts/{parentId}/typeAssociations | Post ContactTypeAssociation
[**PutCompanyContactsByParentIdTypeAssociationsById**](ContactContactTypeAssociationsAPI.md#PutCompanyContactsByParentIdTypeAssociationsById) | **Put** /company/contacts/{parentId}/typeAssociations/{id} | Put ContactTypeAssociation



## DeleteCompanyContactsByParentIdTypeAssociationsById

> DeleteCompanyContactsByParentIdTypeAssociationsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ContactTypeAssociation

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
	id := int32(56) // int32 | typeAssociationId
	parentId := int32(56) // int32 | contactId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContactContactTypeAssociationsAPI.DeleteCompanyContactsByParentIdTypeAssociationsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactContactTypeAssociationsAPI.DeleteCompanyContactsByParentIdTypeAssociationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | typeAssociationId | 
**parentId** | **int32** | contactId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyContactsByParentIdTypeAssociationsByIdRequest struct via the builder pattern


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


## GetCompanyContactsByParentIdTypeAssociations

> []ContactContactTypeAssociationContactTypeAssociation GetCompanyContactsByParentIdTypeAssociations(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ContactTypeAssociation

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
	parentId := int32(56) // int32 | contactId
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
	resp, r, err := apiClient.ContactContactTypeAssociationsAPI.GetCompanyContactsByParentIdTypeAssociations(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactContactTypeAssociationsAPI.GetCompanyContactsByParentIdTypeAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyContactsByParentIdTypeAssociations`: []ContactContactTypeAssociationContactTypeAssociation
	fmt.Fprintf(os.Stdout, "Response from `ContactContactTypeAssociationsAPI.GetCompanyContactsByParentIdTypeAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | contactId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyContactsByParentIdTypeAssociationsRequest struct via the builder pattern


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

[**[]ContactContactTypeAssociationContactTypeAssociation**](ContactContactTypeAssociationContactTypeAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyContactsByParentIdTypeAssociationsById

> ContactContactTypeAssociationContactTypeAssociation GetCompanyContactsByParentIdTypeAssociationsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ContactTypeAssociation

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
	id := int32(56) // int32 | typeAssociationId
	parentId := int32(56) // int32 | contactId
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
	resp, r, err := apiClient.ContactContactTypeAssociationsAPI.GetCompanyContactsByParentIdTypeAssociationsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactContactTypeAssociationsAPI.GetCompanyContactsByParentIdTypeAssociationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyContactsByParentIdTypeAssociationsById`: ContactContactTypeAssociationContactTypeAssociation
	fmt.Fprintf(os.Stdout, "Response from `ContactContactTypeAssociationsAPI.GetCompanyContactsByParentIdTypeAssociationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | typeAssociationId | 
**parentId** | **int32** | contactId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyContactsByParentIdTypeAssociationsByIdRequest struct via the builder pattern


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

[**ContactContactTypeAssociationContactTypeAssociation**](ContactContactTypeAssociationContactTypeAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyContactsByParentIdTypeAssociationsCount

> Count GetCompanyContactsByParentIdTypeAssociationsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ContactTypeAssociation

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
	parentId := int32(56) // int32 | contactId
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
	resp, r, err := apiClient.ContactContactTypeAssociationsAPI.GetCompanyContactsByParentIdTypeAssociationsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactContactTypeAssociationsAPI.GetCompanyContactsByParentIdTypeAssociationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyContactsByParentIdTypeAssociationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ContactContactTypeAssociationsAPI.GetCompanyContactsByParentIdTypeAssociationsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | contactId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyContactsByParentIdTypeAssociationsCountRequest struct via the builder pattern


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


## PatchCompanyContactsByParentIdTypeAssociationsById

> ContactContactTypeAssociationContactTypeAssociation PatchCompanyContactsByParentIdTypeAssociationsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ContactTypeAssociation

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
	id := int32(56) // int32 | typeAssociationId
	parentId := int32(56) // int32 | contactId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactContactTypeAssociationsAPI.PatchCompanyContactsByParentIdTypeAssociationsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactContactTypeAssociationsAPI.PatchCompanyContactsByParentIdTypeAssociationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyContactsByParentIdTypeAssociationsById`: ContactContactTypeAssociationContactTypeAssociation
	fmt.Fprintf(os.Stdout, "Response from `ContactContactTypeAssociationsAPI.PatchCompanyContactsByParentIdTypeAssociationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | typeAssociationId | 
**parentId** | **int32** | contactId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyContactsByParentIdTypeAssociationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ContactContactTypeAssociationContactTypeAssociation**](ContactContactTypeAssociationContactTypeAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyContactsByParentIdTypeAssociations

> ContactContactTypeAssociationContactTypeAssociation PostCompanyContactsByParentIdTypeAssociations(ctx, parentId).ClientId(clientId).ContactContactTypeAssociationContactTypeAssociation(contactContactTypeAssociationContactTypeAssociation).Execute()

Post ContactTypeAssociation

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
	parentId := int32(56) // int32 | contactId
	clientId := "clientId_example" // string | 
	contactContactTypeAssociationContactTypeAssociation := *openapiclient.NewContactContactTypeAssociationContactTypeAssociation(*openapiclient.NewContactTypeReference(), *openapiclient.NewContactReference()) // ContactContactTypeAssociationContactTypeAssociation | contactTypeAssociation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactContactTypeAssociationsAPI.PostCompanyContactsByParentIdTypeAssociations(context.Background(), parentId).ClientId(clientId).ContactContactTypeAssociationContactTypeAssociation(contactContactTypeAssociationContactTypeAssociation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactContactTypeAssociationsAPI.PostCompanyContactsByParentIdTypeAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyContactsByParentIdTypeAssociations`: ContactContactTypeAssociationContactTypeAssociation
	fmt.Fprintf(os.Stdout, "Response from `ContactContactTypeAssociationsAPI.PostCompanyContactsByParentIdTypeAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | contactId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyContactsByParentIdTypeAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **contactContactTypeAssociationContactTypeAssociation** | [**ContactContactTypeAssociationContactTypeAssociation**](ContactContactTypeAssociationContactTypeAssociation.md) | contactTypeAssociation | 

### Return type

[**ContactContactTypeAssociationContactTypeAssociation**](ContactContactTypeAssociationContactTypeAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyContactsByParentIdTypeAssociationsById

> ContactContactTypeAssociationContactTypeAssociation PutCompanyContactsByParentIdTypeAssociationsById(ctx, id, parentId).ClientId(clientId).ContactContactTypeAssociationContactTypeAssociation(contactContactTypeAssociationContactTypeAssociation).Execute()

Put ContactTypeAssociation

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
	id := int32(56) // int32 | typeAssociationId
	parentId := int32(56) // int32 | contactId
	clientId := "clientId_example" // string | 
	contactContactTypeAssociationContactTypeAssociation := *openapiclient.NewContactContactTypeAssociationContactTypeAssociation(*openapiclient.NewContactTypeReference(), *openapiclient.NewContactReference()) // ContactContactTypeAssociationContactTypeAssociation | contactTypeAssociation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactContactTypeAssociationsAPI.PutCompanyContactsByParentIdTypeAssociationsById(context.Background(), id, parentId).ClientId(clientId).ContactContactTypeAssociationContactTypeAssociation(contactContactTypeAssociationContactTypeAssociation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactContactTypeAssociationsAPI.PutCompanyContactsByParentIdTypeAssociationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyContactsByParentIdTypeAssociationsById`: ContactContactTypeAssociationContactTypeAssociation
	fmt.Fprintf(os.Stdout, "Response from `ContactContactTypeAssociationsAPI.PutCompanyContactsByParentIdTypeAssociationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | typeAssociationId | 
**parentId** | **int32** | contactId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyContactsByParentIdTypeAssociationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **contactContactTypeAssociationContactTypeAssociation** | [**ContactContactTypeAssociationContactTypeAssociation**](ContactContactTypeAssociationContactTypeAssociation.md) | contactTypeAssociation | 

### Return type

[**ContactContactTypeAssociationContactTypeAssociation**](ContactContactTypeAssociationContactTypeAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

