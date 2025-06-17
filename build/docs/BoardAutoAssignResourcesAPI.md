# \BoardAutoAssignResourcesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServiceBoardsByParentIdAutoAssignResourcesById**](BoardAutoAssignResourcesAPI.md#DeleteServiceBoardsByParentIdAutoAssignResourcesById) | **Delete** /service/boards/{parentId}/autoAssignResources/{id} | Delete BoardAutoAssignResource
[**GetServiceBoardsByParentIdAutoAssignResources**](BoardAutoAssignResourcesAPI.md#GetServiceBoardsByParentIdAutoAssignResources) | **Get** /service/boards/{parentId}/autoAssignResources | Get List of BoardAutoAssignResource
[**GetServiceBoardsByParentIdAutoAssignResourcesById**](BoardAutoAssignResourcesAPI.md#GetServiceBoardsByParentIdAutoAssignResourcesById) | **Get** /service/boards/{parentId}/autoAssignResources/{id} | Get BoardAutoAssignResource
[**GetServiceBoardsByParentIdAutoAssignResourcesCount**](BoardAutoAssignResourcesAPI.md#GetServiceBoardsByParentIdAutoAssignResourcesCount) | **Get** /service/boards/{parentId}/autoAssignResources/count | Get Count of BoardAutoAssignResource
[**PatchServiceBoardsByParentIdAutoAssignResourcesById**](BoardAutoAssignResourcesAPI.md#PatchServiceBoardsByParentIdAutoAssignResourcesById) | **Patch** /service/boards/{parentId}/autoAssignResources/{id} | Patch BoardAutoAssignResource
[**PostServiceBoardsByParentIdAutoAssignResources**](BoardAutoAssignResourcesAPI.md#PostServiceBoardsByParentIdAutoAssignResources) | **Post** /service/boards/{parentId}/autoAssignResources | Post BoardAutoAssignResource
[**PutServiceBoardsByParentIdAutoAssignResourcesById**](BoardAutoAssignResourcesAPI.md#PutServiceBoardsByParentIdAutoAssignResourcesById) | **Put** /service/boards/{parentId}/autoAssignResources/{id} | Put BoardAutoAssignResource



## DeleteServiceBoardsByParentIdAutoAssignResourcesById

> DeleteServiceBoardsByParentIdAutoAssignResourcesById(ctx, id, parentId).ClientId(clientId).Execute()

Delete BoardAutoAssignResource

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
	id := int32(56) // int32 | autoAssignResourceId
	parentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BoardAutoAssignResourcesAPI.DeleteServiceBoardsByParentIdAutoAssignResourcesById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAutoAssignResourcesAPI.DeleteServiceBoardsByParentIdAutoAssignResourcesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | autoAssignResourceId | 
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceBoardsByParentIdAutoAssignResourcesByIdRequest struct via the builder pattern


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


## GetServiceBoardsByParentIdAutoAssignResources

> []BoardAutoAssignResource GetServiceBoardsByParentIdAutoAssignResources(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of BoardAutoAssignResource

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
	resp, r, err := apiClient.BoardAutoAssignResourcesAPI.GetServiceBoardsByParentIdAutoAssignResources(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAutoAssignResourcesAPI.GetServiceBoardsByParentIdAutoAssignResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceBoardsByParentIdAutoAssignResources`: []BoardAutoAssignResource
	fmt.Fprintf(os.Stdout, "Response from `BoardAutoAssignResourcesAPI.GetServiceBoardsByParentIdAutoAssignResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceBoardsByParentIdAutoAssignResourcesRequest struct via the builder pattern


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

[**[]BoardAutoAssignResource**](BoardAutoAssignResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceBoardsByParentIdAutoAssignResourcesById

> BoardAutoAssignResource GetServiceBoardsByParentIdAutoAssignResourcesById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get BoardAutoAssignResource

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
	id := int32(56) // int32 | autoAssignResourceId
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
	resp, r, err := apiClient.BoardAutoAssignResourcesAPI.GetServiceBoardsByParentIdAutoAssignResourcesById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAutoAssignResourcesAPI.GetServiceBoardsByParentIdAutoAssignResourcesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceBoardsByParentIdAutoAssignResourcesById`: BoardAutoAssignResource
	fmt.Fprintf(os.Stdout, "Response from `BoardAutoAssignResourcesAPI.GetServiceBoardsByParentIdAutoAssignResourcesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | autoAssignResourceId | 
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceBoardsByParentIdAutoAssignResourcesByIdRequest struct via the builder pattern


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

[**BoardAutoAssignResource**](BoardAutoAssignResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceBoardsByParentIdAutoAssignResourcesCount

> Count GetServiceBoardsByParentIdAutoAssignResourcesCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of BoardAutoAssignResource

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
	resp, r, err := apiClient.BoardAutoAssignResourcesAPI.GetServiceBoardsByParentIdAutoAssignResourcesCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAutoAssignResourcesAPI.GetServiceBoardsByParentIdAutoAssignResourcesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceBoardsByParentIdAutoAssignResourcesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `BoardAutoAssignResourcesAPI.GetServiceBoardsByParentIdAutoAssignResourcesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceBoardsByParentIdAutoAssignResourcesCountRequest struct via the builder pattern


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


## PatchServiceBoardsByParentIdAutoAssignResourcesById

> BoardAutoAssignResource PatchServiceBoardsByParentIdAutoAssignResourcesById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch BoardAutoAssignResource

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
	id := int32(56) // int32 | autoAssignResourceId
	parentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAutoAssignResourcesAPI.PatchServiceBoardsByParentIdAutoAssignResourcesById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAutoAssignResourcesAPI.PatchServiceBoardsByParentIdAutoAssignResourcesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchServiceBoardsByParentIdAutoAssignResourcesById`: BoardAutoAssignResource
	fmt.Fprintf(os.Stdout, "Response from `BoardAutoAssignResourcesAPI.PatchServiceBoardsByParentIdAutoAssignResourcesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | autoAssignResourceId | 
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchServiceBoardsByParentIdAutoAssignResourcesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**BoardAutoAssignResource**](BoardAutoAssignResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServiceBoardsByParentIdAutoAssignResources

> BoardAutoAssignResource PostServiceBoardsByParentIdAutoAssignResources(ctx, parentId).ClientId(clientId).BoardAutoAssignResource(boardAutoAssignResource).Execute()

Post BoardAutoAssignResource

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
	boardAutoAssignResource := *openapiclient.NewBoardAutoAssignResource(*openapiclient.NewNotificationRecipientReference()) // BoardAutoAssignResource | boardAutoAssignResource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAutoAssignResourcesAPI.PostServiceBoardsByParentIdAutoAssignResources(context.Background(), parentId).ClientId(clientId).BoardAutoAssignResource(boardAutoAssignResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAutoAssignResourcesAPI.PostServiceBoardsByParentIdAutoAssignResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServiceBoardsByParentIdAutoAssignResources`: BoardAutoAssignResource
	fmt.Fprintf(os.Stdout, "Response from `BoardAutoAssignResourcesAPI.PostServiceBoardsByParentIdAutoAssignResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServiceBoardsByParentIdAutoAssignResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **boardAutoAssignResource** | [**BoardAutoAssignResource**](BoardAutoAssignResource.md) | boardAutoAssignResource | 

### Return type

[**BoardAutoAssignResource**](BoardAutoAssignResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServiceBoardsByParentIdAutoAssignResourcesById

> BoardAutoAssignResource PutServiceBoardsByParentIdAutoAssignResourcesById(ctx, id, parentId).ClientId(clientId).BoardAutoAssignResource(boardAutoAssignResource).Execute()

Put BoardAutoAssignResource

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
	id := int32(56) // int32 | autoAssignResourceId
	parentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 
	boardAutoAssignResource := *openapiclient.NewBoardAutoAssignResource(*openapiclient.NewNotificationRecipientReference()) // BoardAutoAssignResource | boardAutoAssignResource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAutoAssignResourcesAPI.PutServiceBoardsByParentIdAutoAssignResourcesById(context.Background(), id, parentId).ClientId(clientId).BoardAutoAssignResource(boardAutoAssignResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAutoAssignResourcesAPI.PutServiceBoardsByParentIdAutoAssignResourcesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServiceBoardsByParentIdAutoAssignResourcesById`: BoardAutoAssignResource
	fmt.Fprintf(os.Stdout, "Response from `BoardAutoAssignResourcesAPI.PutServiceBoardsByParentIdAutoAssignResourcesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | autoAssignResourceId | 
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServiceBoardsByParentIdAutoAssignResourcesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **boardAutoAssignResource** | [**BoardAutoAssignResource**](BoardAutoAssignResource.md) | boardAutoAssignResource | 

### Return type

[**BoardAutoAssignResource**](BoardAutoAssignResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

