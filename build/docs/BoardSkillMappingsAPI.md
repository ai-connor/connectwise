# \BoardSkillMappingsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServiceBoardsByParentIdSkillMappingsById**](BoardSkillMappingsAPI.md#DeleteServiceBoardsByParentIdSkillMappingsById) | **Delete** /service/boards/{parentId}/skillMappings/{id} | Delete BoardSkillMappings
[**GetServiceBoardsByParentIdSkillMappings**](BoardSkillMappingsAPI.md#GetServiceBoardsByParentIdSkillMappings) | **Get** /service/boards/{parentId}/skillMappings/ | Get List of BoardSkillMappings
[**GetServiceBoardsByParentIdSkillMappingsById**](BoardSkillMappingsAPI.md#GetServiceBoardsByParentIdSkillMappingsById) | **Get** /service/boards/{parentId}/skillMappings/{id} | Get BoardSkillMappings
[**GetServiceBoardsByParentIdSkillMappingsCount**](BoardSkillMappingsAPI.md#GetServiceBoardsByParentIdSkillMappingsCount) | **Get** /service/boards/{parentId}/skillMappings/count | Get Count of BoardSkillMappings
[**PatchServiceBoardsByParentIdSkillMappingsById**](BoardSkillMappingsAPI.md#PatchServiceBoardsByParentIdSkillMappingsById) | **Patch** /service/boards/{parentId}/skillMappings/{id} | Patch BoardSkillMappings
[**PostServiceBoardsByParentIdSkillMappings**](BoardSkillMappingsAPI.md#PostServiceBoardsByParentIdSkillMappings) | **Post** /service/boards/{parentId}/skillMappings/ | Post BoardSkillMappings
[**PutServiceBoardsByParentIdSkillMappingsById**](BoardSkillMappingsAPI.md#PutServiceBoardsByParentIdSkillMappingsById) | **Put** /service/boards/{parentId}/skillMappings/{id} | Put BoardSkillMappings



## DeleteServiceBoardsByParentIdSkillMappingsById

> DeleteServiceBoardsByParentIdSkillMappingsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete BoardSkillMappings

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
	id := int32(56) // int32 | BoardSkillMappingId
	parentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BoardSkillMappingsAPI.DeleteServiceBoardsByParentIdSkillMappingsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardSkillMappingsAPI.DeleteServiceBoardsByParentIdSkillMappingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | BoardSkillMappingId | 
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceBoardsByParentIdSkillMappingsByIdRequest struct via the builder pattern


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


## GetServiceBoardsByParentIdSkillMappings

> []BoardSkillMapping GetServiceBoardsByParentIdSkillMappings(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of BoardSkillMappings

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
	resp, r, err := apiClient.BoardSkillMappingsAPI.GetServiceBoardsByParentIdSkillMappings(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardSkillMappingsAPI.GetServiceBoardsByParentIdSkillMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceBoardsByParentIdSkillMappings`: []BoardSkillMapping
	fmt.Fprintf(os.Stdout, "Response from `BoardSkillMappingsAPI.GetServiceBoardsByParentIdSkillMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceBoardsByParentIdSkillMappingsRequest struct via the builder pattern


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

[**[]BoardSkillMapping**](BoardSkillMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceBoardsByParentIdSkillMappingsById

> BoardSkillMapping GetServiceBoardsByParentIdSkillMappingsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get BoardSkillMappings

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
	id := int32(56) // int32 | BoardSkillMappingId
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
	resp, r, err := apiClient.BoardSkillMappingsAPI.GetServiceBoardsByParentIdSkillMappingsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardSkillMappingsAPI.GetServiceBoardsByParentIdSkillMappingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceBoardsByParentIdSkillMappingsById`: BoardSkillMapping
	fmt.Fprintf(os.Stdout, "Response from `BoardSkillMappingsAPI.GetServiceBoardsByParentIdSkillMappingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | BoardSkillMappingId | 
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceBoardsByParentIdSkillMappingsByIdRequest struct via the builder pattern


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

[**BoardSkillMapping**](BoardSkillMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceBoardsByParentIdSkillMappingsCount

> Count GetServiceBoardsByParentIdSkillMappingsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of BoardSkillMappings

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
	resp, r, err := apiClient.BoardSkillMappingsAPI.GetServiceBoardsByParentIdSkillMappingsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardSkillMappingsAPI.GetServiceBoardsByParentIdSkillMappingsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceBoardsByParentIdSkillMappingsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `BoardSkillMappingsAPI.GetServiceBoardsByParentIdSkillMappingsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceBoardsByParentIdSkillMappingsCountRequest struct via the builder pattern


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


## PatchServiceBoardsByParentIdSkillMappingsById

> BoardSkillMapping PatchServiceBoardsByParentIdSkillMappingsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch BoardSkillMappings

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
	id := int32(56) // int32 | BoardSkillMappingId
	parentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardSkillMappingsAPI.PatchServiceBoardsByParentIdSkillMappingsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardSkillMappingsAPI.PatchServiceBoardsByParentIdSkillMappingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchServiceBoardsByParentIdSkillMappingsById`: BoardSkillMapping
	fmt.Fprintf(os.Stdout, "Response from `BoardSkillMappingsAPI.PatchServiceBoardsByParentIdSkillMappingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | BoardSkillMappingId | 
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchServiceBoardsByParentIdSkillMappingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**BoardSkillMapping**](BoardSkillMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServiceBoardsByParentIdSkillMappings

> BoardSkillMapping PostServiceBoardsByParentIdSkillMappings(ctx, parentId).ClientId(clientId).BoardSkillMapping(boardSkillMapping).Execute()

Post BoardSkillMappings

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
	boardSkillMapping := *openapiclient.NewBoardSkillMapping(*openapiclient.NewServiceTypeReference(), *openapiclient.NewSkillCategoryReference(), *openapiclient.NewSkillReference()) // BoardSkillMapping | BoardSkillMapping

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardSkillMappingsAPI.PostServiceBoardsByParentIdSkillMappings(context.Background(), parentId).ClientId(clientId).BoardSkillMapping(boardSkillMapping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardSkillMappingsAPI.PostServiceBoardsByParentIdSkillMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServiceBoardsByParentIdSkillMappings`: BoardSkillMapping
	fmt.Fprintf(os.Stdout, "Response from `BoardSkillMappingsAPI.PostServiceBoardsByParentIdSkillMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServiceBoardsByParentIdSkillMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **boardSkillMapping** | [**BoardSkillMapping**](BoardSkillMapping.md) | BoardSkillMapping | 

### Return type

[**BoardSkillMapping**](BoardSkillMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServiceBoardsByParentIdSkillMappingsById

> BoardSkillMapping PutServiceBoardsByParentIdSkillMappingsById(ctx, id, parentId).ClientId(clientId).BoardSkillMapping(boardSkillMapping).Execute()

Put BoardSkillMappings

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
	id := int32(56) // int32 | BoardSkillMappingId
	parentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 
	boardSkillMapping := *openapiclient.NewBoardSkillMapping(*openapiclient.NewServiceTypeReference(), *openapiclient.NewSkillCategoryReference(), *openapiclient.NewSkillReference()) // BoardSkillMapping | boardSkillMapping

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardSkillMappingsAPI.PutServiceBoardsByParentIdSkillMappingsById(context.Background(), id, parentId).ClientId(clientId).BoardSkillMapping(boardSkillMapping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardSkillMappingsAPI.PutServiceBoardsByParentIdSkillMappingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServiceBoardsByParentIdSkillMappingsById`: BoardSkillMapping
	fmt.Fprintf(os.Stdout, "Response from `BoardSkillMappingsAPI.PutServiceBoardsByParentIdSkillMappingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | BoardSkillMappingId | 
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServiceBoardsByParentIdSkillMappingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **boardSkillMapping** | [**BoardSkillMapping**](BoardSkillMapping.md) | boardSkillMapping | 

### Return type

[**BoardSkillMapping**](BoardSkillMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

