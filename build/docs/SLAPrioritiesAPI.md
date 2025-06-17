# \SLAPrioritiesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServiceSLAsByParentIdPrioritiesById**](SLAPrioritiesAPI.md#DeleteServiceSLAsByParentIdPrioritiesById) | **Delete** /service/SLAs/{parentId}/priorities/{id} | Delete SLAPriority
[**GetServiceSLAsByParentIdPriorities**](SLAPrioritiesAPI.md#GetServiceSLAsByParentIdPriorities) | **Get** /service/SLAs/{parentId}/priorities | Get List of SLAPriority
[**GetServiceSLAsByParentIdPrioritiesById**](SLAPrioritiesAPI.md#GetServiceSLAsByParentIdPrioritiesById) | **Get** /service/SLAs/{parentId}/priorities/{id} | Get SLAPriority
[**GetServiceSLAsByParentIdPrioritiesCount**](SLAPrioritiesAPI.md#GetServiceSLAsByParentIdPrioritiesCount) | **Get** /service/SLAs/{parentId}/priorities/count | Get Count of SLAPriority
[**PatchServiceSLAsByParentIdPrioritiesById**](SLAPrioritiesAPI.md#PatchServiceSLAsByParentIdPrioritiesById) | **Patch** /service/SLAs/{parentId}/priorities/{id} | Patch SLAPriority
[**PostServiceSLAsByParentIdPriorities**](SLAPrioritiesAPI.md#PostServiceSLAsByParentIdPriorities) | **Post** /service/SLAs/{parentId}/priorities | Post SLAPriority
[**PutServiceSLAsByParentIdPrioritiesById**](SLAPrioritiesAPI.md#PutServiceSLAsByParentIdPrioritiesById) | **Put** /service/SLAs/{parentId}/priorities/{id} | Put SLAPriority



## DeleteServiceSLAsByParentIdPrioritiesById

> DeleteServiceSLAsByParentIdPrioritiesById(ctx, id, parentId).ClientId(clientId).Execute()

Delete SLAPriority

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
	id := int32(56) // int32 | priorityId
	parentId := int32(56) // int32 | SLAId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SLAPrioritiesAPI.DeleteServiceSLAsByParentIdPrioritiesById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLAPrioritiesAPI.DeleteServiceSLAsByParentIdPrioritiesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | priorityId | 
**parentId** | **int32** | SLAId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceSLAsByParentIdPrioritiesByIdRequest struct via the builder pattern


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


## GetServiceSLAsByParentIdPriorities

> []SLAPriority GetServiceSLAsByParentIdPriorities(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of SLAPriority

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
	parentId := int32(56) // int32 | SLAId
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
	resp, r, err := apiClient.SLAPrioritiesAPI.GetServiceSLAsByParentIdPriorities(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLAPrioritiesAPI.GetServiceSLAsByParentIdPriorities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceSLAsByParentIdPriorities`: []SLAPriority
	fmt.Fprintf(os.Stdout, "Response from `SLAPrioritiesAPI.GetServiceSLAsByParentIdPriorities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | SLAId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceSLAsByParentIdPrioritiesRequest struct via the builder pattern


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

[**[]SLAPriority**](SLAPriority.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceSLAsByParentIdPrioritiesById

> SLAPriority GetServiceSLAsByParentIdPrioritiesById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get SLAPriority

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
	id := int32(56) // int32 | priorityId
	parentId := int32(56) // int32 | SLAId
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
	resp, r, err := apiClient.SLAPrioritiesAPI.GetServiceSLAsByParentIdPrioritiesById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLAPrioritiesAPI.GetServiceSLAsByParentIdPrioritiesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceSLAsByParentIdPrioritiesById`: SLAPriority
	fmt.Fprintf(os.Stdout, "Response from `SLAPrioritiesAPI.GetServiceSLAsByParentIdPrioritiesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | priorityId | 
**parentId** | **int32** | SLAId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceSLAsByParentIdPrioritiesByIdRequest struct via the builder pattern


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

[**SLAPriority**](SLAPriority.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceSLAsByParentIdPrioritiesCount

> Count GetServiceSLAsByParentIdPrioritiesCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of SLAPriority

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
	parentId := int32(56) // int32 | SLAId
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
	resp, r, err := apiClient.SLAPrioritiesAPI.GetServiceSLAsByParentIdPrioritiesCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLAPrioritiesAPI.GetServiceSLAsByParentIdPrioritiesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceSLAsByParentIdPrioritiesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `SLAPrioritiesAPI.GetServiceSLAsByParentIdPrioritiesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | SLAId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceSLAsByParentIdPrioritiesCountRequest struct via the builder pattern


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


## PatchServiceSLAsByParentIdPrioritiesById

> SLAPriority PatchServiceSLAsByParentIdPrioritiesById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch SLAPriority

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
	id := int32(56) // int32 | priorityId
	parentId := int32(56) // int32 | SLAId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLAPrioritiesAPI.PatchServiceSLAsByParentIdPrioritiesById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLAPrioritiesAPI.PatchServiceSLAsByParentIdPrioritiesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchServiceSLAsByParentIdPrioritiesById`: SLAPriority
	fmt.Fprintf(os.Stdout, "Response from `SLAPrioritiesAPI.PatchServiceSLAsByParentIdPrioritiesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | priorityId | 
**parentId** | **int32** | SLAId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchServiceSLAsByParentIdPrioritiesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**SLAPriority**](SLAPriority.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServiceSLAsByParentIdPriorities

> SLAPriority PostServiceSLAsByParentIdPriorities(ctx, parentId).ClientId(clientId).SLAPriority(sLAPriority).Execute()

Post SLAPriority

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
	parentId := int32(56) // int32 | SLAId
	clientId := "clientId_example" // string | 
	sLAPriority := *openapiclient.NewSLAPriority(*openapiclient.NewPriorityReference()) // SLAPriority | sLAPriority

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLAPrioritiesAPI.PostServiceSLAsByParentIdPriorities(context.Background(), parentId).ClientId(clientId).SLAPriority(sLAPriority).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLAPrioritiesAPI.PostServiceSLAsByParentIdPriorities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServiceSLAsByParentIdPriorities`: SLAPriority
	fmt.Fprintf(os.Stdout, "Response from `SLAPrioritiesAPI.PostServiceSLAsByParentIdPriorities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | SLAId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServiceSLAsByParentIdPrioritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **sLAPriority** | [**SLAPriority**](SLAPriority.md) | sLAPriority | 

### Return type

[**SLAPriority**](SLAPriority.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServiceSLAsByParentIdPrioritiesById

> SLAPriority PutServiceSLAsByParentIdPrioritiesById(ctx, id, parentId).ClientId(clientId).SLAPriority(sLAPriority).Execute()

Put SLAPriority

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
	id := int32(56) // int32 | priorityId
	parentId := int32(56) // int32 | SLAId
	clientId := "clientId_example" // string | 
	sLAPriority := *openapiclient.NewSLAPriority(*openapiclient.NewPriorityReference()) // SLAPriority | sLAPriority

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SLAPrioritiesAPI.PutServiceSLAsByParentIdPrioritiesById(context.Background(), id, parentId).ClientId(clientId).SLAPriority(sLAPriority).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SLAPrioritiesAPI.PutServiceSLAsByParentIdPrioritiesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServiceSLAsByParentIdPrioritiesById`: SLAPriority
	fmt.Fprintf(os.Stdout, "Response from `SLAPrioritiesAPI.PutServiceSLAsByParentIdPrioritiesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | priorityId | 
**parentId** | **int32** | SLAId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServiceSLAsByParentIdPrioritiesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **sLAPriority** | [**SLAPriority**](SLAPriority.md) | sLAPriority | 

### Return type

[**SLAPriority**](SLAPriority.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

