# \ProjectSecurityRolesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProjectSecurityRolesById**](ProjectSecurityRolesAPI.md#DeleteProjectSecurityRolesById) | **Delete** /project/securityRoles/{id} | Delete ProjectSecurityRole
[**GetProjectSecurityRoles**](ProjectSecurityRolesAPI.md#GetProjectSecurityRoles) | **Get** /project/securityRoles | Get List of ProjectSecurityRole
[**GetProjectSecurityRolesById**](ProjectSecurityRolesAPI.md#GetProjectSecurityRolesById) | **Get** /project/securityRoles/{id} | Get ProjectSecurityRole
[**GetProjectSecurityRolesCount**](ProjectSecurityRolesAPI.md#GetProjectSecurityRolesCount) | **Get** /project/securityRoles/count | Get Count of ProjectSecurityRole
[**PatchProjectSecurityRolesById**](ProjectSecurityRolesAPI.md#PatchProjectSecurityRolesById) | **Patch** /project/securityRoles/{id} | Patch ProjectSecurityRole
[**PostProjectSecurityRoles**](ProjectSecurityRolesAPI.md#PostProjectSecurityRoles) | **Post** /project/securityRoles | Post ProjectSecurityRole
[**PutProjectSecurityRolesById**](ProjectSecurityRolesAPI.md#PutProjectSecurityRolesById) | **Put** /project/securityRoles/{id} | Put ProjectSecurityRole



## DeleteProjectSecurityRolesById

> DeleteProjectSecurityRolesById(ctx, id).ClientId(clientId).Execute()

Delete ProjectSecurityRole

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
	id := int32(56) // int32 | securityRoleId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectSecurityRolesAPI.DeleteProjectSecurityRolesById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectSecurityRolesAPI.DeleteProjectSecurityRolesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | securityRoleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectSecurityRolesByIdRequest struct via the builder pattern


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


## GetProjectSecurityRoles

> []ProjectSecurityRole GetProjectSecurityRoles(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProjectSecurityRole

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
	resp, r, err := apiClient.ProjectSecurityRolesAPI.GetProjectSecurityRoles(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectSecurityRolesAPI.GetProjectSecurityRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectSecurityRoles`: []ProjectSecurityRole
	fmt.Fprintf(os.Stdout, "Response from `ProjectSecurityRolesAPI.GetProjectSecurityRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectSecurityRolesRequest struct via the builder pattern


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

[**[]ProjectSecurityRole**](ProjectSecurityRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectSecurityRolesById

> ProjectSecurityRole GetProjectSecurityRolesById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ProjectSecurityRole

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
	id := int32(56) // int32 | securityRoleId
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
	resp, r, err := apiClient.ProjectSecurityRolesAPI.GetProjectSecurityRolesById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectSecurityRolesAPI.GetProjectSecurityRolesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectSecurityRolesById`: ProjectSecurityRole
	fmt.Fprintf(os.Stdout, "Response from `ProjectSecurityRolesAPI.GetProjectSecurityRolesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | securityRoleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectSecurityRolesByIdRequest struct via the builder pattern


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

[**ProjectSecurityRole**](ProjectSecurityRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectSecurityRolesCount

> Count GetProjectSecurityRolesCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ProjectSecurityRole

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
	resp, r, err := apiClient.ProjectSecurityRolesAPI.GetProjectSecurityRolesCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectSecurityRolesAPI.GetProjectSecurityRolesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectSecurityRolesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ProjectSecurityRolesAPI.GetProjectSecurityRolesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectSecurityRolesCountRequest struct via the builder pattern


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


## PatchProjectSecurityRolesById

> ProjectSecurityRole PatchProjectSecurityRolesById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ProjectSecurityRole

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
	id := int32(56) // int32 | securityRoleId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectSecurityRolesAPI.PatchProjectSecurityRolesById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectSecurityRolesAPI.PatchProjectSecurityRolesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProjectSecurityRolesById`: ProjectSecurityRole
	fmt.Fprintf(os.Stdout, "Response from `ProjectSecurityRolesAPI.PatchProjectSecurityRolesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | securityRoleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProjectSecurityRolesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ProjectSecurityRole**](ProjectSecurityRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectSecurityRoles

> ProjectSecurityRole PostProjectSecurityRoles(ctx).ClientId(clientId).ProjectSecurityRole(projectSecurityRole).Execute()

Post ProjectSecurityRole

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
	projectSecurityRole := *openapiclient.NewProjectSecurityRole("Name_example") // ProjectSecurityRole | projectSecurityRole

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectSecurityRolesAPI.PostProjectSecurityRoles(context.Background()).ClientId(clientId).ProjectSecurityRole(projectSecurityRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectSecurityRolesAPI.PostProjectSecurityRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectSecurityRoles`: ProjectSecurityRole
	fmt.Fprintf(os.Stdout, "Response from `ProjectSecurityRolesAPI.PostProjectSecurityRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectSecurityRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **projectSecurityRole** | [**ProjectSecurityRole**](ProjectSecurityRole.md) | projectSecurityRole | 

### Return type

[**ProjectSecurityRole**](ProjectSecurityRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProjectSecurityRolesById

> ProjectSecurityRole PutProjectSecurityRolesById(ctx, id).ClientId(clientId).ProjectSecurityRole(projectSecurityRole).Execute()

Put ProjectSecurityRole

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
	id := int32(56) // int32 | securityRoleId
	clientId := "clientId_example" // string | 
	projectSecurityRole := *openapiclient.NewProjectSecurityRole("Name_example") // ProjectSecurityRole | projectSecurityRole

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectSecurityRolesAPI.PutProjectSecurityRolesById(context.Background(), id).ClientId(clientId).ProjectSecurityRole(projectSecurityRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectSecurityRolesAPI.PutProjectSecurityRolesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProjectSecurityRolesById`: ProjectSecurityRole
	fmt.Fprintf(os.Stdout, "Response from `ProjectSecurityRolesAPI.PutProjectSecurityRolesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | securityRoleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProjectSecurityRolesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **projectSecurityRole** | [**ProjectSecurityRole**](ProjectSecurityRole.md) | projectSecurityRole | 

### Return type

[**ProjectSecurityRole**](ProjectSecurityRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

