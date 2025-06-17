# \ProjectBoardTeamMembersAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProjectBoardsByGrandparentIdTeamsByParentIdMembersById**](ProjectBoardTeamMembersAPI.md#DeleteProjectBoardsByGrandparentIdTeamsByParentIdMembersById) | **Delete** /project/boards/{grandparentId}/teams/{parentId}/members/{id} | Delete ProjectBoardTeamMember
[**GetProjectBoardsByGrandparentIdTeamsByParentIdMembers**](ProjectBoardTeamMembersAPI.md#GetProjectBoardsByGrandparentIdTeamsByParentIdMembers) | **Get** /project/boards/{grandparentId}/teams/{parentId}/members | Get List of ProjectBoardTeamMember
[**GetProjectBoardsByGrandparentIdTeamsByParentIdMembersById**](ProjectBoardTeamMembersAPI.md#GetProjectBoardsByGrandparentIdTeamsByParentIdMembersById) | **Get** /project/boards/{grandparentId}/teams/{parentId}/members/{id} | Get ProjectBoardTeamMember
[**PatchProjectBoardsByGrandparentIdTeamsByParentIdMembersById**](ProjectBoardTeamMembersAPI.md#PatchProjectBoardsByGrandparentIdTeamsByParentIdMembersById) | **Patch** /project/boards/{grandparentId}/teams/{parentId}/members/{id} | Patch ProjectBoardTeamMember
[**PostProjectBoardsByGrandparentIdTeamsByParentIdMembers**](ProjectBoardTeamMembersAPI.md#PostProjectBoardsByGrandparentIdTeamsByParentIdMembers) | **Post** /project/boards/{grandparentId}/teams/{parentId}/members | Post ProjectBoardTeamMember
[**PutProjectBoardsByGrandparentIdTeamsByParentIdMembersById**](ProjectBoardTeamMembersAPI.md#PutProjectBoardsByGrandparentIdTeamsByParentIdMembersById) | **Put** /project/boards/{grandparentId}/teams/{parentId}/members/{id} | Put ProjectBoardTeamMember



## DeleteProjectBoardsByGrandparentIdTeamsByParentIdMembersById

> DeleteProjectBoardsByGrandparentIdTeamsByParentIdMembersById(ctx, id, parentId, grandparentId).ClientId(clientId).Execute()

Delete ProjectBoardTeamMember

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
	id := int32(56) // int32 | memberId
	parentId := int32(56) // int32 | teamId
	grandparentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectBoardTeamMembersAPI.DeleteProjectBoardsByGrandparentIdTeamsByParentIdMembersById(context.Background(), id, parentId, grandparentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBoardTeamMembersAPI.DeleteProjectBoardsByGrandparentIdTeamsByParentIdMembersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | memberId | 
**parentId** | **int32** | teamId | 
**grandparentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectBoardsByGrandparentIdTeamsByParentIdMembersByIdRequest struct via the builder pattern


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


## GetProjectBoardsByGrandparentIdTeamsByParentIdMembers

> []ProjectBoardTeamMember GetProjectBoardsByGrandparentIdTeamsByParentIdMembers(ctx, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProjectBoardTeamMember

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
	parentId := int32(56) // int32 | teamId
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
	resp, r, err := apiClient.ProjectBoardTeamMembersAPI.GetProjectBoardsByGrandparentIdTeamsByParentIdMembers(context.Background(), parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBoardTeamMembersAPI.GetProjectBoardsByGrandparentIdTeamsByParentIdMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectBoardsByGrandparentIdTeamsByParentIdMembers`: []ProjectBoardTeamMember
	fmt.Fprintf(os.Stdout, "Response from `ProjectBoardTeamMembersAPI.GetProjectBoardsByGrandparentIdTeamsByParentIdMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | teamId | 
**grandparentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectBoardsByGrandparentIdTeamsByParentIdMembersRequest struct via the builder pattern


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

[**[]ProjectBoardTeamMember**](ProjectBoardTeamMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectBoardsByGrandparentIdTeamsByParentIdMembersById

> ProjectBoardTeamMember GetProjectBoardsByGrandparentIdTeamsByParentIdMembersById(ctx, id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ProjectBoardTeamMember

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
	id := int32(56) // int32 | memberId
	parentId := int32(56) // int32 | teamId
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
	resp, r, err := apiClient.ProjectBoardTeamMembersAPI.GetProjectBoardsByGrandparentIdTeamsByParentIdMembersById(context.Background(), id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBoardTeamMembersAPI.GetProjectBoardsByGrandparentIdTeamsByParentIdMembersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectBoardsByGrandparentIdTeamsByParentIdMembersById`: ProjectBoardTeamMember
	fmt.Fprintf(os.Stdout, "Response from `ProjectBoardTeamMembersAPI.GetProjectBoardsByGrandparentIdTeamsByParentIdMembersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | memberId | 
**parentId** | **int32** | teamId | 
**grandparentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectBoardsByGrandparentIdTeamsByParentIdMembersByIdRequest struct via the builder pattern


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

[**ProjectBoardTeamMember**](ProjectBoardTeamMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchProjectBoardsByGrandparentIdTeamsByParentIdMembersById

> ProjectBoardTeamMember PatchProjectBoardsByGrandparentIdTeamsByParentIdMembersById(ctx, id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ProjectBoardTeamMember

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
	id := int32(56) // int32 | memberId
	parentId := int32(56) // int32 | teamId
	grandparentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectBoardTeamMembersAPI.PatchProjectBoardsByGrandparentIdTeamsByParentIdMembersById(context.Background(), id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBoardTeamMembersAPI.PatchProjectBoardsByGrandparentIdTeamsByParentIdMembersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProjectBoardsByGrandparentIdTeamsByParentIdMembersById`: ProjectBoardTeamMember
	fmt.Fprintf(os.Stdout, "Response from `ProjectBoardTeamMembersAPI.PatchProjectBoardsByGrandparentIdTeamsByParentIdMembersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | memberId | 
**parentId** | **int32** | teamId | 
**grandparentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProjectBoardsByGrandparentIdTeamsByParentIdMembersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ProjectBoardTeamMember**](ProjectBoardTeamMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectBoardsByGrandparentIdTeamsByParentIdMembers

> ProjectBoardTeamMember PostProjectBoardsByGrandparentIdTeamsByParentIdMembers(ctx, parentId, grandparentId).ClientId(clientId).ProjectBoardTeamMember(projectBoardTeamMember).Execute()

Post ProjectBoardTeamMember

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
	parentId := int32(56) // int32 | teamId
	grandparentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 
	projectBoardTeamMember := *openapiclient.NewProjectBoardTeamMember(*openapiclient.NewMemberReference(), *openapiclient.NewProjectRoleReference()) // ProjectBoardTeamMember | teamMember

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectBoardTeamMembersAPI.PostProjectBoardsByGrandparentIdTeamsByParentIdMembers(context.Background(), parentId, grandparentId).ClientId(clientId).ProjectBoardTeamMember(projectBoardTeamMember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBoardTeamMembersAPI.PostProjectBoardsByGrandparentIdTeamsByParentIdMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectBoardsByGrandparentIdTeamsByParentIdMembers`: ProjectBoardTeamMember
	fmt.Fprintf(os.Stdout, "Response from `ProjectBoardTeamMembersAPI.PostProjectBoardsByGrandparentIdTeamsByParentIdMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | teamId | 
**grandparentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectBoardsByGrandparentIdTeamsByParentIdMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **projectBoardTeamMember** | [**ProjectBoardTeamMember**](ProjectBoardTeamMember.md) | teamMember | 

### Return type

[**ProjectBoardTeamMember**](ProjectBoardTeamMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProjectBoardsByGrandparentIdTeamsByParentIdMembersById

> ProjectBoardTeamMember PutProjectBoardsByGrandparentIdTeamsByParentIdMembersById(ctx, id, parentId, grandparentId).ClientId(clientId).ProjectBoardTeamMember(projectBoardTeamMember).Execute()

Put ProjectBoardTeamMember

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
	id := int32(56) // int32 | memberId
	parentId := int32(56) // int32 | teamId
	grandparentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 
	projectBoardTeamMember := *openapiclient.NewProjectBoardTeamMember(*openapiclient.NewMemberReference(), *openapiclient.NewProjectRoleReference()) // ProjectBoardTeamMember | teamMember

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectBoardTeamMembersAPI.PutProjectBoardsByGrandparentIdTeamsByParentIdMembersById(context.Background(), id, parentId, grandparentId).ClientId(clientId).ProjectBoardTeamMember(projectBoardTeamMember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBoardTeamMembersAPI.PutProjectBoardsByGrandparentIdTeamsByParentIdMembersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProjectBoardsByGrandparentIdTeamsByParentIdMembersById`: ProjectBoardTeamMember
	fmt.Fprintf(os.Stdout, "Response from `ProjectBoardTeamMembersAPI.PutProjectBoardsByGrandparentIdTeamsByParentIdMembersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | memberId | 
**parentId** | **int32** | teamId | 
**grandparentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProjectBoardsByGrandparentIdTeamsByParentIdMembersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **projectBoardTeamMember** | [**ProjectBoardTeamMember**](ProjectBoardTeamMember.md) | teamMember | 

### Return type

[**ProjectBoardTeamMember**](ProjectBoardTeamMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

