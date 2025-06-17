# \ProjectTemplateTicketsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProjectProjectTemplatesByParentIdProjectTemplateTicketsById**](ProjectTemplateTicketsAPI.md#DeleteProjectProjectTemplatesByParentIdProjectTemplateTicketsById) | **Delete** /project/projectTemplates/{parentId}/projectTemplateTickets/{id} | Delete ProjectTemplateTickets
[**GetProjectProjectTemplatesByParentIdProjectTemplateTickets**](ProjectTemplateTicketsAPI.md#GetProjectProjectTemplatesByParentIdProjectTemplateTickets) | **Get** /project/projectTemplates/{parentId}/projectTemplateTickets | Get List of ProjectTemplateTickets
[**GetProjectProjectTemplatesByParentIdProjectTemplateTicketsById**](ProjectTemplateTicketsAPI.md#GetProjectProjectTemplatesByParentIdProjectTemplateTicketsById) | **Get** /project/projectTemplates/{parentId}/projectTemplateTickets/{id} | Get ProjectTemplateTickets
[**GetProjectProjectTemplatesByParentIdProjectTemplateTicketsCount**](ProjectTemplateTicketsAPI.md#GetProjectProjectTemplatesByParentIdProjectTemplateTicketsCount) | **Get** /project/projectTemplates/{parentId}/projectTemplateTickets/count | Get Count of ProjectTemplateTickets
[**GetProjectProjectTemplatesProjectTemplateTickets**](ProjectTemplateTicketsAPI.md#GetProjectProjectTemplatesProjectTemplateTickets) | **Get** /project/projectTemplates/projectTemplateTickets | Get List of ProjectTemplateTickets
[**PatchProjectProjectTemplatesByParentIdProjectTemplateTicketsById**](ProjectTemplateTicketsAPI.md#PatchProjectProjectTemplatesByParentIdProjectTemplateTicketsById) | **Patch** /project/projectTemplates/{parentId}/projectTemplateTickets/{id} | Patch ProjectTemplateTickets
[**PostProjectProjectTemplatesByParentIdProjectTemplateTickets**](ProjectTemplateTicketsAPI.md#PostProjectProjectTemplatesByParentIdProjectTemplateTickets) | **Post** /project/projectTemplates/{parentId}/projectTemplateTickets | Post ProjectTemplateTickets
[**PutProjectProjectTemplatesByParentIdProjectTemplateTicketsById**](ProjectTemplateTicketsAPI.md#PutProjectProjectTemplatesByParentIdProjectTemplateTicketsById) | **Put** /project/projectTemplates/{parentId}/projectTemplateTickets/{id} | Put ProjectTemplateTickets



## DeleteProjectProjectTemplatesByParentIdProjectTemplateTicketsById

> DeleteProjectProjectTemplatesByParentIdProjectTemplateTicketsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ProjectTemplateTickets

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
	id := int32(56) // int32 | ProjectTemplateTicketId
	parentId := int32(56) // int32 | projectTemplateId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectTemplateTicketsAPI.DeleteProjectProjectTemplatesByParentIdProjectTemplateTicketsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplateTicketsAPI.DeleteProjectProjectTemplatesByParentIdProjectTemplateTicketsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ProjectTemplateTicketId | 
**parentId** | **int32** | projectTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectProjectTemplatesByParentIdProjectTemplateTicketsByIdRequest struct via the builder pattern


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


## GetProjectProjectTemplatesByParentIdProjectTemplateTickets

> []ProjectTemplateTicket GetProjectProjectTemplatesByParentIdProjectTemplateTickets(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProjectTemplateTickets

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
	parentId := int32(56) // int32 | projectTemplateId
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
	resp, r, err := apiClient.ProjectTemplateTicketsAPI.GetProjectProjectTemplatesByParentIdProjectTemplateTickets(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplateTicketsAPI.GetProjectProjectTemplatesByParentIdProjectTemplateTickets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectTemplatesByParentIdProjectTemplateTickets`: []ProjectTemplateTicket
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplateTicketsAPI.GetProjectProjectTemplatesByParentIdProjectTemplateTickets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | projectTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectTemplatesByParentIdProjectTemplateTicketsRequest struct via the builder pattern


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

[**[]ProjectTemplateTicket**](ProjectTemplateTicket.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectProjectTemplatesByParentIdProjectTemplateTicketsById

> ProjectTemplateTicket GetProjectProjectTemplatesByParentIdProjectTemplateTicketsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ProjectTemplateTickets

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
	id := int32(56) // int32 | ProjectTemplateTicketId
	parentId := int32(56) // int32 | projectTemplateId
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
	resp, r, err := apiClient.ProjectTemplateTicketsAPI.GetProjectProjectTemplatesByParentIdProjectTemplateTicketsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplateTicketsAPI.GetProjectProjectTemplatesByParentIdProjectTemplateTicketsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectTemplatesByParentIdProjectTemplateTicketsById`: ProjectTemplateTicket
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplateTicketsAPI.GetProjectProjectTemplatesByParentIdProjectTemplateTicketsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ProjectTemplateTicketId | 
**parentId** | **int32** | projectTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectTemplatesByParentIdProjectTemplateTicketsByIdRequest struct via the builder pattern


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

[**ProjectTemplateTicket**](ProjectTemplateTicket.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectProjectTemplatesByParentIdProjectTemplateTicketsCount

> Count GetProjectProjectTemplatesByParentIdProjectTemplateTicketsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ProjectTemplateTickets

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
	parentId := int32(56) // int32 | projectTemplateId
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
	resp, r, err := apiClient.ProjectTemplateTicketsAPI.GetProjectProjectTemplatesByParentIdProjectTemplateTicketsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplateTicketsAPI.GetProjectProjectTemplatesByParentIdProjectTemplateTicketsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectTemplatesByParentIdProjectTemplateTicketsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplateTicketsAPI.GetProjectProjectTemplatesByParentIdProjectTemplateTicketsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | projectTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectTemplatesByParentIdProjectTemplateTicketsCountRequest struct via the builder pattern


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


## GetProjectProjectTemplatesProjectTemplateTickets

> []ProjectTemplateTicket GetProjectProjectTemplatesProjectTemplateTickets(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProjectTemplateTickets

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
	parentId := int32(56) // int32 | projectTemplateId
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
	resp, r, err := apiClient.ProjectTemplateTicketsAPI.GetProjectProjectTemplatesProjectTemplateTickets(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplateTicketsAPI.GetProjectProjectTemplatesProjectTemplateTickets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectTemplatesProjectTemplateTickets`: []ProjectTemplateTicket
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplateTicketsAPI.GetProjectProjectTemplatesProjectTemplateTickets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | projectTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectTemplatesProjectTemplateTicketsRequest struct via the builder pattern


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

[**[]ProjectTemplateTicket**](ProjectTemplateTicket.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchProjectProjectTemplatesByParentIdProjectTemplateTicketsById

> ProjectTemplateTicket PatchProjectProjectTemplatesByParentIdProjectTemplateTicketsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ProjectTemplateTickets

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
	id := int32(56) // int32 | ProjectTemplateTicketId
	parentId := int32(56) // int32 | projectTemplateId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTemplateTicketsAPI.PatchProjectProjectTemplatesByParentIdProjectTemplateTicketsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplateTicketsAPI.PatchProjectProjectTemplatesByParentIdProjectTemplateTicketsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProjectProjectTemplatesByParentIdProjectTemplateTicketsById`: ProjectTemplateTicket
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplateTicketsAPI.PatchProjectProjectTemplatesByParentIdProjectTemplateTicketsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ProjectTemplateTicketId | 
**parentId** | **int32** | projectTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProjectProjectTemplatesByParentIdProjectTemplateTicketsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ProjectTemplateTicket**](ProjectTemplateTicket.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectProjectTemplatesByParentIdProjectTemplateTickets

> ProjectTemplateTicket PostProjectProjectTemplatesByParentIdProjectTemplateTickets(ctx, parentId).ClientId(clientId).ProjectTemplateTicket(projectTemplateTicket).Execute()

Post ProjectTemplateTickets

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
	parentId := int32(56) // int32 | projectTemplateId
	clientId := "clientId_example" // string | 
	projectTemplateTicket := *openapiclient.NewProjectTemplateTicket("Description_example") // ProjectTemplateTicket | ProjectTemplateTicket

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTemplateTicketsAPI.PostProjectProjectTemplatesByParentIdProjectTemplateTickets(context.Background(), parentId).ClientId(clientId).ProjectTemplateTicket(projectTemplateTicket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplateTicketsAPI.PostProjectProjectTemplatesByParentIdProjectTemplateTickets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectProjectTemplatesByParentIdProjectTemplateTickets`: ProjectTemplateTicket
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplateTicketsAPI.PostProjectProjectTemplatesByParentIdProjectTemplateTickets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | projectTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectProjectTemplatesByParentIdProjectTemplateTicketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **projectTemplateTicket** | [**ProjectTemplateTicket**](ProjectTemplateTicket.md) | ProjectTemplateTicket | 

### Return type

[**ProjectTemplateTicket**](ProjectTemplateTicket.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProjectProjectTemplatesByParentIdProjectTemplateTicketsById

> ProjectTemplateTicket PutProjectProjectTemplatesByParentIdProjectTemplateTicketsById(ctx, id, parentId).ClientId(clientId).ProjectTemplateTicket(projectTemplateTicket).Execute()

Put ProjectTemplateTickets

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
	id := int32(56) // int32 | ProjectTemplateTicketId
	parentId := int32(56) // int32 | projectTemplateId
	clientId := "clientId_example" // string | 
	projectTemplateTicket := *openapiclient.NewProjectTemplateTicket("Description_example") // ProjectTemplateTicket | companyTypeAssociation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTemplateTicketsAPI.PutProjectProjectTemplatesByParentIdProjectTemplateTicketsById(context.Background(), id, parentId).ClientId(clientId).ProjectTemplateTicket(projectTemplateTicket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplateTicketsAPI.PutProjectProjectTemplatesByParentIdProjectTemplateTicketsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProjectProjectTemplatesByParentIdProjectTemplateTicketsById`: ProjectTemplateTicket
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplateTicketsAPI.PutProjectProjectTemplatesByParentIdProjectTemplateTicketsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ProjectTemplateTicketId | 
**parentId** | **int32** | projectTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProjectProjectTemplatesByParentIdProjectTemplateTicketsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **projectTemplateTicket** | [**ProjectTemplateTicket**](ProjectTemplateTicket.md) | companyTypeAssociation | 

### Return type

[**ProjectTemplateTicket**](ProjectTemplateTicket.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

