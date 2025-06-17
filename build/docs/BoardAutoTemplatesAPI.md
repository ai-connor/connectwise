# \BoardAutoTemplatesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServiceBoardsByParentIdAutoTemplatesById**](BoardAutoTemplatesAPI.md#DeleteServiceBoardsByParentIdAutoTemplatesById) | **Delete** /service/boards/{parentId}/autoTemplates/{id} | Delete BoardAutoTemplate
[**GetServiceBoardsByParentIdAutoTemplates**](BoardAutoTemplatesAPI.md#GetServiceBoardsByParentIdAutoTemplates) | **Get** /service/boards/{parentId}/autoTemplates | Get List of BoardAutoTemplate
[**GetServiceBoardsByParentIdAutoTemplatesById**](BoardAutoTemplatesAPI.md#GetServiceBoardsByParentIdAutoTemplatesById) | **Get** /service/boards/{parentId}/autoTemplates/{id} | Get BoardAutoTemplate
[**GetServiceBoardsByParentIdAutoTemplatesCount**](BoardAutoTemplatesAPI.md#GetServiceBoardsByParentIdAutoTemplatesCount) | **Get** /service/boards/{parentId}/autoTemplates/count | Get Count of BoardAutoTemplate
[**PatchServiceBoardsByParentIdAutoTemplatesById**](BoardAutoTemplatesAPI.md#PatchServiceBoardsByParentIdAutoTemplatesById) | **Patch** /service/boards/{parentId}/autoTemplates/{id} | Patch BoardAutoTemplate
[**PostServiceBoardsByParentIdAutoTemplates**](BoardAutoTemplatesAPI.md#PostServiceBoardsByParentIdAutoTemplates) | **Post** /service/boards/{parentId}/autoTemplates | Post BoardAutoTemplate
[**PutServiceBoardsByParentIdAutoTemplatesById**](BoardAutoTemplatesAPI.md#PutServiceBoardsByParentIdAutoTemplatesById) | **Put** /service/boards/{parentId}/autoTemplates/{id} | Put BoardAutoTemplate



## DeleteServiceBoardsByParentIdAutoTemplatesById

> DeleteServiceBoardsByParentIdAutoTemplatesById(ctx, id, parentId).ClientId(clientId).Execute()

Delete BoardAutoTemplate

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
	id := int32(56) // int32 | autoTemplateId
	parentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BoardAutoTemplatesAPI.DeleteServiceBoardsByParentIdAutoTemplatesById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAutoTemplatesAPI.DeleteServiceBoardsByParentIdAutoTemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | autoTemplateId | 
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceBoardsByParentIdAutoTemplatesByIdRequest struct via the builder pattern


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


## GetServiceBoardsByParentIdAutoTemplates

> []BoardAutoTemplate GetServiceBoardsByParentIdAutoTemplates(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of BoardAutoTemplate

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
	resp, r, err := apiClient.BoardAutoTemplatesAPI.GetServiceBoardsByParentIdAutoTemplates(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAutoTemplatesAPI.GetServiceBoardsByParentIdAutoTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceBoardsByParentIdAutoTemplates`: []BoardAutoTemplate
	fmt.Fprintf(os.Stdout, "Response from `BoardAutoTemplatesAPI.GetServiceBoardsByParentIdAutoTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceBoardsByParentIdAutoTemplatesRequest struct via the builder pattern


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

[**[]BoardAutoTemplate**](BoardAutoTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceBoardsByParentIdAutoTemplatesById

> BoardAutoTemplate GetServiceBoardsByParentIdAutoTemplatesById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get BoardAutoTemplate

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
	id := int32(56) // int32 | autoTemplateId
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
	resp, r, err := apiClient.BoardAutoTemplatesAPI.GetServiceBoardsByParentIdAutoTemplatesById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAutoTemplatesAPI.GetServiceBoardsByParentIdAutoTemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceBoardsByParentIdAutoTemplatesById`: BoardAutoTemplate
	fmt.Fprintf(os.Stdout, "Response from `BoardAutoTemplatesAPI.GetServiceBoardsByParentIdAutoTemplatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | autoTemplateId | 
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceBoardsByParentIdAutoTemplatesByIdRequest struct via the builder pattern


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

[**BoardAutoTemplate**](BoardAutoTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceBoardsByParentIdAutoTemplatesCount

> Count GetServiceBoardsByParentIdAutoTemplatesCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of BoardAutoTemplate

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
	resp, r, err := apiClient.BoardAutoTemplatesAPI.GetServiceBoardsByParentIdAutoTemplatesCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAutoTemplatesAPI.GetServiceBoardsByParentIdAutoTemplatesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceBoardsByParentIdAutoTemplatesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `BoardAutoTemplatesAPI.GetServiceBoardsByParentIdAutoTemplatesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceBoardsByParentIdAutoTemplatesCountRequest struct via the builder pattern


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


## PatchServiceBoardsByParentIdAutoTemplatesById

> BoardAutoTemplate PatchServiceBoardsByParentIdAutoTemplatesById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch BoardAutoTemplate

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
	id := int32(56) // int32 | autoTemplateId
	parentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAutoTemplatesAPI.PatchServiceBoardsByParentIdAutoTemplatesById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAutoTemplatesAPI.PatchServiceBoardsByParentIdAutoTemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchServiceBoardsByParentIdAutoTemplatesById`: BoardAutoTemplate
	fmt.Fprintf(os.Stdout, "Response from `BoardAutoTemplatesAPI.PatchServiceBoardsByParentIdAutoTemplatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | autoTemplateId | 
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchServiceBoardsByParentIdAutoTemplatesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**BoardAutoTemplate**](BoardAutoTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServiceBoardsByParentIdAutoTemplates

> BoardAutoTemplate PostServiceBoardsByParentIdAutoTemplates(ctx, parentId).ClientId(clientId).BoardAutoTemplate(boardAutoTemplate).Execute()

Post BoardAutoTemplate

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
	boardAutoTemplate := *openapiclient.NewBoardAutoTemplate(*openapiclient.NewServiceTypeReference(), *openapiclient.NewServiceSubTypeReference(), *openapiclient.NewServiceItemReference(), *openapiclient.NewServiceTemplateReference()) // BoardAutoTemplate | boardAutoTemplate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAutoTemplatesAPI.PostServiceBoardsByParentIdAutoTemplates(context.Background(), parentId).ClientId(clientId).BoardAutoTemplate(boardAutoTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAutoTemplatesAPI.PostServiceBoardsByParentIdAutoTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServiceBoardsByParentIdAutoTemplates`: BoardAutoTemplate
	fmt.Fprintf(os.Stdout, "Response from `BoardAutoTemplatesAPI.PostServiceBoardsByParentIdAutoTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServiceBoardsByParentIdAutoTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **boardAutoTemplate** | [**BoardAutoTemplate**](BoardAutoTemplate.md) | boardAutoTemplate | 

### Return type

[**BoardAutoTemplate**](BoardAutoTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServiceBoardsByParentIdAutoTemplatesById

> BoardAutoTemplate PutServiceBoardsByParentIdAutoTemplatesById(ctx, id, parentId).ClientId(clientId).BoardAutoTemplate(boardAutoTemplate).Execute()

Put BoardAutoTemplate

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
	id := int32(56) // int32 | autoTemplateId
	parentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 
	boardAutoTemplate := *openapiclient.NewBoardAutoTemplate(*openapiclient.NewServiceTypeReference(), *openapiclient.NewServiceSubTypeReference(), *openapiclient.NewServiceItemReference(), *openapiclient.NewServiceTemplateReference()) // BoardAutoTemplate | boardAutoTemplate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAutoTemplatesAPI.PutServiceBoardsByParentIdAutoTemplatesById(context.Background(), id, parentId).ClientId(clientId).BoardAutoTemplate(boardAutoTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAutoTemplatesAPI.PutServiceBoardsByParentIdAutoTemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServiceBoardsByParentIdAutoTemplatesById`: BoardAutoTemplate
	fmt.Fprintf(os.Stdout, "Response from `BoardAutoTemplatesAPI.PutServiceBoardsByParentIdAutoTemplatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | autoTemplateId | 
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServiceBoardsByParentIdAutoTemplatesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **boardAutoTemplate** | [**BoardAutoTemplate**](BoardAutoTemplate.md) | boardAutoTemplate | 

### Return type

[**BoardAutoTemplate**](BoardAutoTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

