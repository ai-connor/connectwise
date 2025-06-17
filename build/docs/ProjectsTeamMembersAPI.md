# \ProjectsTeamMembersAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProjectProjectsByParentIdTeamMembersById**](ProjectsTeamMembersAPI.md#DeleteProjectProjectsByParentIdTeamMembersById) | **Delete** /project/projects/{parentId}/teamMembers/{id} | Delete ProjectTeamMember
[**GetProjectProjectsByParentIdTeamMembers**](ProjectsTeamMembersAPI.md#GetProjectProjectsByParentIdTeamMembers) | **Get** /project/projects/{parentId}/teamMembers | Get List of ProjectTeamMember
[**GetProjectProjectsByParentIdTeamMembersById**](ProjectsTeamMembersAPI.md#GetProjectProjectsByParentIdTeamMembersById) | **Get** /project/projects/{parentId}/teamMembers/{id} | Get ProjectTeamMember
[**GetProjectProjectsByParentIdTeamMembersCount**](ProjectsTeamMembersAPI.md#GetProjectProjectsByParentIdTeamMembersCount) | **Get** /project/projects/{parentId}/teamMembers/count | Get Count of ProjectTeamMember
[**PatchProjectProjectsByParentIdTeamMembersById**](ProjectsTeamMembersAPI.md#PatchProjectProjectsByParentIdTeamMembersById) | **Patch** /project/projects/{parentId}/teamMembers/{id} | Patch ProjectTeamMember
[**PostProjectProjectsByParentIdTeamMembers**](ProjectsTeamMembersAPI.md#PostProjectProjectsByParentIdTeamMembers) | **Post** /project/projects/{parentId}/teamMembers | Post ProjectTeamMember
[**PutProjectProjectsByParentIdTeamMembersById**](ProjectsTeamMembersAPI.md#PutProjectProjectsByParentIdTeamMembersById) | **Put** /project/projects/{parentId}/teamMembers/{id} | Put ProjectTeamMember



## DeleteProjectProjectsByParentIdTeamMembersById

> DeleteProjectProjectsByParentIdTeamMembersById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ProjectTeamMember

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
	id := int32(56) // int32 | teamMemberId
	parentId := int32(56) // int32 | projectId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsTeamMembersAPI.DeleteProjectProjectsByParentIdTeamMembersById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsTeamMembersAPI.DeleteProjectProjectsByParentIdTeamMembersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | teamMemberId | 
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectProjectsByParentIdTeamMembersByIdRequest struct via the builder pattern


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


## GetProjectProjectsByParentIdTeamMembers

> []ProjectTeamMember GetProjectProjectsByParentIdTeamMembers(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProjectTeamMember

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
	parentId := int32(56) // int32 | projectId
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
	resp, r, err := apiClient.ProjectsTeamMembersAPI.GetProjectProjectsByParentIdTeamMembers(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsTeamMembersAPI.GetProjectProjectsByParentIdTeamMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectsByParentIdTeamMembers`: []ProjectTeamMember
	fmt.Fprintf(os.Stdout, "Response from `ProjectsTeamMembersAPI.GetProjectProjectsByParentIdTeamMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectsByParentIdTeamMembersRequest struct via the builder pattern


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

[**[]ProjectTeamMember**](ProjectTeamMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectProjectsByParentIdTeamMembersById

> ProjectTeamMember GetProjectProjectsByParentIdTeamMembersById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ProjectTeamMember

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
	id := int32(56) // int32 | teamMemberId
	parentId := int32(56) // int32 | projectId
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
	resp, r, err := apiClient.ProjectsTeamMembersAPI.GetProjectProjectsByParentIdTeamMembersById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsTeamMembersAPI.GetProjectProjectsByParentIdTeamMembersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectsByParentIdTeamMembersById`: ProjectTeamMember
	fmt.Fprintf(os.Stdout, "Response from `ProjectsTeamMembersAPI.GetProjectProjectsByParentIdTeamMembersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | teamMemberId | 
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectsByParentIdTeamMembersByIdRequest struct via the builder pattern


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

[**ProjectTeamMember**](ProjectTeamMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectProjectsByParentIdTeamMembersCount

> Count GetProjectProjectsByParentIdTeamMembersCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ProjectTeamMember

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
	parentId := int32(56) // int32 | projectId
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
	resp, r, err := apiClient.ProjectsTeamMembersAPI.GetProjectProjectsByParentIdTeamMembersCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsTeamMembersAPI.GetProjectProjectsByParentIdTeamMembersCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectsByParentIdTeamMembersCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ProjectsTeamMembersAPI.GetProjectProjectsByParentIdTeamMembersCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectsByParentIdTeamMembersCountRequest struct via the builder pattern


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


## PatchProjectProjectsByParentIdTeamMembersById

> ProjectTeamMember PatchProjectProjectsByParentIdTeamMembersById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ProjectTeamMember

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
	id := int32(56) // int32 | teamMemberId
	parentId := int32(56) // int32 | projectId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsTeamMembersAPI.PatchProjectProjectsByParentIdTeamMembersById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsTeamMembersAPI.PatchProjectProjectsByParentIdTeamMembersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProjectProjectsByParentIdTeamMembersById`: ProjectTeamMember
	fmt.Fprintf(os.Stdout, "Response from `ProjectsTeamMembersAPI.PatchProjectProjectsByParentIdTeamMembersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | teamMemberId | 
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProjectProjectsByParentIdTeamMembersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ProjectTeamMember**](ProjectTeamMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectProjectsByParentIdTeamMembers

> ProjectTeamMember PostProjectProjectsByParentIdTeamMembers(ctx, parentId).ClientId(clientId).ProjectTeamMember(projectTeamMember).Execute()

Post ProjectTeamMember

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
	parentId := int32(56) // int32 | projectId
	clientId := "clientId_example" // string | 
	projectTeamMember := *openapiclient.NewProjectTeamMember(*openapiclient.NewMemberReference(), *openapiclient.NewProjectRoleReference()) // ProjectTeamMember | teamMember

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsTeamMembersAPI.PostProjectProjectsByParentIdTeamMembers(context.Background(), parentId).ClientId(clientId).ProjectTeamMember(projectTeamMember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsTeamMembersAPI.PostProjectProjectsByParentIdTeamMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectProjectsByParentIdTeamMembers`: ProjectTeamMember
	fmt.Fprintf(os.Stdout, "Response from `ProjectsTeamMembersAPI.PostProjectProjectsByParentIdTeamMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectProjectsByParentIdTeamMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **projectTeamMember** | [**ProjectTeamMember**](ProjectTeamMember.md) | teamMember | 

### Return type

[**ProjectTeamMember**](ProjectTeamMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProjectProjectsByParentIdTeamMembersById

> ProjectTeamMember PutProjectProjectsByParentIdTeamMembersById(ctx, id, parentId).ClientId(clientId).ProjectTeamMember(projectTeamMember).Execute()

Put ProjectTeamMember

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
	id := int32(56) // int32 | teamMemberId
	parentId := int32(56) // int32 | projectId
	clientId := "clientId_example" // string | 
	projectTeamMember := *openapiclient.NewProjectTeamMember(*openapiclient.NewMemberReference(), *openapiclient.NewProjectRoleReference()) // ProjectTeamMember | teamMember

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsTeamMembersAPI.PutProjectProjectsByParentIdTeamMembersById(context.Background(), id, parentId).ClientId(clientId).ProjectTeamMember(projectTeamMember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsTeamMembersAPI.PutProjectProjectsByParentIdTeamMembersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProjectProjectsByParentIdTeamMembersById`: ProjectTeamMember
	fmt.Fprintf(os.Stdout, "Response from `ProjectsTeamMembersAPI.PutProjectProjectsByParentIdTeamMembersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | teamMemberId | 
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProjectProjectsByParentIdTeamMembersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **projectTeamMember** | [**ProjectTeamMember**](ProjectTeamMember.md) | teamMember | 

### Return type

[**ProjectTeamMember**](ProjectTeamMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

