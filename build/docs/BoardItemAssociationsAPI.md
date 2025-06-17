# \BoardItemAssociationsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceBoardsByGrandparentIdItemsByParentIdAssociations**](BoardItemAssociationsAPI.md#GetServiceBoardsByGrandparentIdItemsByParentIdAssociations) | **Get** /service/boards/{grandparentId}/items/{parentId}/associations | Get List of BoardItemAssociation
[**GetServiceBoardsByGrandparentIdItemsByParentIdAssociationsById**](BoardItemAssociationsAPI.md#GetServiceBoardsByGrandparentIdItemsByParentIdAssociationsById) | **Get** /service/boards/{grandparentId}/items/{parentId}/associations/{id} | Get BoardItemAssociation
[**GetServiceBoardsByGrandparentIdItemsByParentIdAssociationsCount**](BoardItemAssociationsAPI.md#GetServiceBoardsByGrandparentIdItemsByParentIdAssociationsCount) | **Get** /service/boards/{grandparentId}/items/{parentId}/associations/count | Get Count of BoardItemAssociation
[**PatchServiceBoardsByGrandparentIdItemsByParentIdAssociationsById**](BoardItemAssociationsAPI.md#PatchServiceBoardsByGrandparentIdItemsByParentIdAssociationsById) | **Patch** /service/boards/{grandparentId}/items/{parentId}/associations/{id} | Patch BoardItemAssociation
[**PutServiceBoardsByGrandparentIdItemsByParentIdAssociationsById**](BoardItemAssociationsAPI.md#PutServiceBoardsByGrandparentIdItemsByParentIdAssociationsById) | **Put** /service/boards/{grandparentId}/items/{parentId}/associations/{id} | Put BoardItemAssociation



## GetServiceBoardsByGrandparentIdItemsByParentIdAssociations

> []BoardItemAssociation GetServiceBoardsByGrandparentIdItemsByParentIdAssociations(ctx, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of BoardItemAssociation

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
	parentId := int32(56) // int32 | itemId
	grandparentId := int32(56) // int32 | boardId
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
	resp, r, err := apiClient.BoardItemAssociationsAPI.GetServiceBoardsByGrandparentIdItemsByParentIdAssociations(context.Background(), parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardItemAssociationsAPI.GetServiceBoardsByGrandparentIdItemsByParentIdAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceBoardsByGrandparentIdItemsByParentIdAssociations`: []BoardItemAssociation
	fmt.Fprintf(os.Stdout, "Response from `BoardItemAssociationsAPI.GetServiceBoardsByGrandparentIdItemsByParentIdAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | itemId | 
**grandparentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceBoardsByGrandparentIdItemsByParentIdAssociationsRequest struct via the builder pattern


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

[**[]BoardItemAssociation**](BoardItemAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceBoardsByGrandparentIdItemsByParentIdAssociationsById

> BoardItemAssociation GetServiceBoardsByGrandparentIdItemsByParentIdAssociationsById(ctx, id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get BoardItemAssociation

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
	id := int32(56) // int32 | associationId
	parentId := int32(56) // int32 | itemId
	grandparentId := int32(56) // int32 | boardId
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
	resp, r, err := apiClient.BoardItemAssociationsAPI.GetServiceBoardsByGrandparentIdItemsByParentIdAssociationsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardItemAssociationsAPI.GetServiceBoardsByGrandparentIdItemsByParentIdAssociationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceBoardsByGrandparentIdItemsByParentIdAssociationsById`: BoardItemAssociation
	fmt.Fprintf(os.Stdout, "Response from `BoardItemAssociationsAPI.GetServiceBoardsByGrandparentIdItemsByParentIdAssociationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | associationId | 
**parentId** | **int32** | itemId | 
**grandparentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceBoardsByGrandparentIdItemsByParentIdAssociationsByIdRequest struct via the builder pattern


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

[**BoardItemAssociation**](BoardItemAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceBoardsByGrandparentIdItemsByParentIdAssociationsCount

> Count GetServiceBoardsByGrandparentIdItemsByParentIdAssociationsCount(ctx, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of BoardItemAssociation

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
	parentId := int32(56) // int32 | itemId
	grandparentId := int32(56) // int32 | boardId
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
	resp, r, err := apiClient.BoardItemAssociationsAPI.GetServiceBoardsByGrandparentIdItemsByParentIdAssociationsCount(context.Background(), parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardItemAssociationsAPI.GetServiceBoardsByGrandparentIdItemsByParentIdAssociationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceBoardsByGrandparentIdItemsByParentIdAssociationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `BoardItemAssociationsAPI.GetServiceBoardsByGrandparentIdItemsByParentIdAssociationsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | itemId | 
**grandparentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceBoardsByGrandparentIdItemsByParentIdAssociationsCountRequest struct via the builder pattern


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


## PatchServiceBoardsByGrandparentIdItemsByParentIdAssociationsById

> BoardItemAssociation PatchServiceBoardsByGrandparentIdItemsByParentIdAssociationsById(ctx, id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch BoardItemAssociation

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
	id := int32(56) // int32 | associationId
	parentId := int32(56) // int32 | itemId
	grandparentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardItemAssociationsAPI.PatchServiceBoardsByGrandparentIdItemsByParentIdAssociationsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardItemAssociationsAPI.PatchServiceBoardsByGrandparentIdItemsByParentIdAssociationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchServiceBoardsByGrandparentIdItemsByParentIdAssociationsById`: BoardItemAssociation
	fmt.Fprintf(os.Stdout, "Response from `BoardItemAssociationsAPI.PatchServiceBoardsByGrandparentIdItemsByParentIdAssociationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | associationId | 
**parentId** | **int32** | itemId | 
**grandparentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchServiceBoardsByGrandparentIdItemsByParentIdAssociationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**BoardItemAssociation**](BoardItemAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServiceBoardsByGrandparentIdItemsByParentIdAssociationsById

> BoardItemAssociation PutServiceBoardsByGrandparentIdItemsByParentIdAssociationsById(ctx, id, parentId, grandparentId).ClientId(clientId).BoardItemAssociation(boardItemAssociation).Execute()

Put BoardItemAssociation

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
	id := int32(56) // int32 | associationId
	parentId := int32(56) // int32 | itemId
	grandparentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 
	boardItemAssociation := *openapiclient.NewBoardItemAssociation(int32(123)) // BoardItemAssociation | itemAssociation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardItemAssociationsAPI.PutServiceBoardsByGrandparentIdItemsByParentIdAssociationsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).BoardItemAssociation(boardItemAssociation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardItemAssociationsAPI.PutServiceBoardsByGrandparentIdItemsByParentIdAssociationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServiceBoardsByGrandparentIdItemsByParentIdAssociationsById`: BoardItemAssociation
	fmt.Fprintf(os.Stdout, "Response from `BoardItemAssociationsAPI.PutServiceBoardsByGrandparentIdItemsByParentIdAssociationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | associationId | 
**parentId** | **int32** | itemId | 
**grandparentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServiceBoardsByGrandparentIdItemsByParentIdAssociationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **boardItemAssociation** | [**BoardItemAssociation**](BoardItemAssociation.md) | itemAssociation | 

### Return type

[**BoardItemAssociation**](BoardItemAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

