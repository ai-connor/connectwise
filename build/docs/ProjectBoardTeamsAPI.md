# \ProjectBoardTeamsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProjectBoardsByParentIdTeamsById**](ProjectBoardTeamsAPI.md#DeleteProjectBoardsByParentIdTeamsById) | **Delete** /project/boards/{parentId}/teams/{id} | Delete ProjectBoardTeam
[**GetProjectBoardsByParentIdTeams**](ProjectBoardTeamsAPI.md#GetProjectBoardsByParentIdTeams) | **Get** /project/boards/{parentId}/teams | Get List of ProjectBoardTeam
[**GetProjectBoardsByParentIdTeamsById**](ProjectBoardTeamsAPI.md#GetProjectBoardsByParentIdTeamsById) | **Get** /project/boards/{parentId}/teams/{id} | Get ProjectBoardTeam
[**GetProjectBoardsByParentIdTeamsCount**](ProjectBoardTeamsAPI.md#GetProjectBoardsByParentIdTeamsCount) | **Get** /project/boards/{parentId}/teams/count | Get Count of ProjectBoardTeam
[**PatchProjectBoardsByParentIdTeamsById**](ProjectBoardTeamsAPI.md#PatchProjectBoardsByParentIdTeamsById) | **Patch** /project/boards/{parentId}/teams/{id} | Patch ProjectBoardTeam
[**PostProjectBoardsByParentIdTeams**](ProjectBoardTeamsAPI.md#PostProjectBoardsByParentIdTeams) | **Post** /project/boards/{parentId}/teams | Post ProjectBoardTeam
[**PutProjectBoardsByParentIdTeamsById**](ProjectBoardTeamsAPI.md#PutProjectBoardsByParentIdTeamsById) | **Put** /project/boards/{parentId}/teams/{id} | Put ProjectBoardTeam



## DeleteProjectBoardsByParentIdTeamsById

> DeleteProjectBoardsByParentIdTeamsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ProjectBoardTeam

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
	id := int32(56) // int32 | teamId
	parentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectBoardTeamsAPI.DeleteProjectBoardsByParentIdTeamsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBoardTeamsAPI.DeleteProjectBoardsByParentIdTeamsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | teamId | 
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectBoardsByParentIdTeamsByIdRequest struct via the builder pattern


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


## GetProjectBoardsByParentIdTeams

> []ProjectBoardTeam GetProjectBoardsByParentIdTeams(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProjectBoardTeam

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
	resp, r, err := apiClient.ProjectBoardTeamsAPI.GetProjectBoardsByParentIdTeams(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBoardTeamsAPI.GetProjectBoardsByParentIdTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectBoardsByParentIdTeams`: []ProjectBoardTeam
	fmt.Fprintf(os.Stdout, "Response from `ProjectBoardTeamsAPI.GetProjectBoardsByParentIdTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectBoardsByParentIdTeamsRequest struct via the builder pattern


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

[**[]ProjectBoardTeam**](ProjectBoardTeam.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectBoardsByParentIdTeamsById

> ProjectBoardTeam GetProjectBoardsByParentIdTeamsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ProjectBoardTeam

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
	id := int32(56) // int32 | teamId
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
	resp, r, err := apiClient.ProjectBoardTeamsAPI.GetProjectBoardsByParentIdTeamsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBoardTeamsAPI.GetProjectBoardsByParentIdTeamsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectBoardsByParentIdTeamsById`: ProjectBoardTeam
	fmt.Fprintf(os.Stdout, "Response from `ProjectBoardTeamsAPI.GetProjectBoardsByParentIdTeamsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | teamId | 
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectBoardsByParentIdTeamsByIdRequest struct via the builder pattern


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

[**ProjectBoardTeam**](ProjectBoardTeam.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectBoardsByParentIdTeamsCount

> Count GetProjectBoardsByParentIdTeamsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ProjectBoardTeam

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
	resp, r, err := apiClient.ProjectBoardTeamsAPI.GetProjectBoardsByParentIdTeamsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBoardTeamsAPI.GetProjectBoardsByParentIdTeamsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectBoardsByParentIdTeamsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ProjectBoardTeamsAPI.GetProjectBoardsByParentIdTeamsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectBoardsByParentIdTeamsCountRequest struct via the builder pattern


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


## PatchProjectBoardsByParentIdTeamsById

> ProjectBoardTeam PatchProjectBoardsByParentIdTeamsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ProjectBoardTeam

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
	id := int32(56) // int32 | teamId
	parentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectBoardTeamsAPI.PatchProjectBoardsByParentIdTeamsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBoardTeamsAPI.PatchProjectBoardsByParentIdTeamsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProjectBoardsByParentIdTeamsById`: ProjectBoardTeam
	fmt.Fprintf(os.Stdout, "Response from `ProjectBoardTeamsAPI.PatchProjectBoardsByParentIdTeamsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | teamId | 
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProjectBoardsByParentIdTeamsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ProjectBoardTeam**](ProjectBoardTeam.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectBoardsByParentIdTeams

> ProjectBoardTeam PostProjectBoardsByParentIdTeams(ctx, parentId).ClientId(clientId).ProjectBoardTeam(projectBoardTeam).Execute()

Post ProjectBoardTeam

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
	projectBoardTeam := *openapiclient.NewProjectBoardTeam("Name_example") // ProjectBoardTeam | team

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectBoardTeamsAPI.PostProjectBoardsByParentIdTeams(context.Background(), parentId).ClientId(clientId).ProjectBoardTeam(projectBoardTeam).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBoardTeamsAPI.PostProjectBoardsByParentIdTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectBoardsByParentIdTeams`: ProjectBoardTeam
	fmt.Fprintf(os.Stdout, "Response from `ProjectBoardTeamsAPI.PostProjectBoardsByParentIdTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectBoardsByParentIdTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **projectBoardTeam** | [**ProjectBoardTeam**](ProjectBoardTeam.md) | team | 

### Return type

[**ProjectBoardTeam**](ProjectBoardTeam.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProjectBoardsByParentIdTeamsById

> ProjectBoardTeam PutProjectBoardsByParentIdTeamsById(ctx, id, parentId).ClientId(clientId).ProjectBoardTeam(projectBoardTeam).Execute()

Put ProjectBoardTeam

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
	id := int32(56) // int32 | teamId
	parentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 
	projectBoardTeam := *openapiclient.NewProjectBoardTeam("Name_example") // ProjectBoardTeam | team

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectBoardTeamsAPI.PutProjectBoardsByParentIdTeamsById(context.Background(), id, parentId).ClientId(clientId).ProjectBoardTeam(projectBoardTeam).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBoardTeamsAPI.PutProjectBoardsByParentIdTeamsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProjectBoardsByParentIdTeamsById`: ProjectBoardTeam
	fmt.Fprintf(os.Stdout, "Response from `ProjectBoardTeamsAPI.PutProjectBoardsByParentIdTeamsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | teamId | 
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProjectBoardsByParentIdTeamsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **projectBoardTeam** | [**ProjectBoardTeam**](ProjectBoardTeam.md) | team | 

### Return type

[**ProjectBoardTeam**](ProjectBoardTeam.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

