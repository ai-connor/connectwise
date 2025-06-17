# \ProjectSecurityRoleSettingsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProjectSecurityRolesByParentIdSettings**](ProjectSecurityRoleSettingsAPI.md#GetProjectSecurityRolesByParentIdSettings) | **Get** /project/securityRoles/{parentId}/settings | Get List of ProjectSecurityRoleSetting
[**GetProjectSecurityRolesByParentIdSettingsById**](ProjectSecurityRoleSettingsAPI.md#GetProjectSecurityRolesByParentIdSettingsById) | **Get** /project/securityRoles/{parentId}/settings/{id} | Get ProjectSecurityRoleSetting
[**GetProjectSecurityRolesByParentIdSettingsCount**](ProjectSecurityRoleSettingsAPI.md#GetProjectSecurityRolesByParentIdSettingsCount) | **Get** /project/securityRoles/{parentId}/settings/count | Get Count of ProjectSecurityRoleSetting
[**PatchProjectSecurityRolesByParentIdSettingsById**](ProjectSecurityRoleSettingsAPI.md#PatchProjectSecurityRolesByParentIdSettingsById) | **Patch** /project/securityRoles/{parentId}/settings/{id} | Patch ProjectSecurityRoleSetting
[**PutProjectSecurityRolesByParentIdSettingsById**](ProjectSecurityRoleSettingsAPI.md#PutProjectSecurityRolesByParentIdSettingsById) | **Put** /project/securityRoles/{parentId}/settings/{id} | Put ProjectSecurityRoleSetting



## GetProjectSecurityRolesByParentIdSettings

> []ProjectSecurityRoleSetting GetProjectSecurityRolesByParentIdSettings(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProjectSecurityRoleSetting

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
	parentId := int32(56) // int32 | securityRoleId
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
	resp, r, err := apiClient.ProjectSecurityRoleSettingsAPI.GetProjectSecurityRolesByParentIdSettings(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectSecurityRoleSettingsAPI.GetProjectSecurityRolesByParentIdSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectSecurityRolesByParentIdSettings`: []ProjectSecurityRoleSetting
	fmt.Fprintf(os.Stdout, "Response from `ProjectSecurityRoleSettingsAPI.GetProjectSecurityRolesByParentIdSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | securityRoleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectSecurityRolesByParentIdSettingsRequest struct via the builder pattern


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

[**[]ProjectSecurityRoleSetting**](ProjectSecurityRoleSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectSecurityRolesByParentIdSettingsById

> ProjectSecurityRoleSetting GetProjectSecurityRolesByParentIdSettingsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ProjectSecurityRoleSetting

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
	id := int32(56) // int32 | settingId
	parentId := int32(56) // int32 | securityRoleId
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
	resp, r, err := apiClient.ProjectSecurityRoleSettingsAPI.GetProjectSecurityRolesByParentIdSettingsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectSecurityRoleSettingsAPI.GetProjectSecurityRolesByParentIdSettingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectSecurityRolesByParentIdSettingsById`: ProjectSecurityRoleSetting
	fmt.Fprintf(os.Stdout, "Response from `ProjectSecurityRoleSettingsAPI.GetProjectSecurityRolesByParentIdSettingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | settingId | 
**parentId** | **int32** | securityRoleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectSecurityRolesByParentIdSettingsByIdRequest struct via the builder pattern


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

[**ProjectSecurityRoleSetting**](ProjectSecurityRoleSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectSecurityRolesByParentIdSettingsCount

> Count GetProjectSecurityRolesByParentIdSettingsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ProjectSecurityRoleSetting

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
	parentId := int32(56) // int32 | securityRoleId
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
	resp, r, err := apiClient.ProjectSecurityRoleSettingsAPI.GetProjectSecurityRolesByParentIdSettingsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectSecurityRoleSettingsAPI.GetProjectSecurityRolesByParentIdSettingsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectSecurityRolesByParentIdSettingsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ProjectSecurityRoleSettingsAPI.GetProjectSecurityRolesByParentIdSettingsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | securityRoleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectSecurityRolesByParentIdSettingsCountRequest struct via the builder pattern


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


## PatchProjectSecurityRolesByParentIdSettingsById

> ProjectSecurityRoleSetting PatchProjectSecurityRolesByParentIdSettingsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ProjectSecurityRoleSetting

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
	id := int32(56) // int32 | settingId
	parentId := int32(56) // int32 | securityRoleId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectSecurityRoleSettingsAPI.PatchProjectSecurityRolesByParentIdSettingsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectSecurityRoleSettingsAPI.PatchProjectSecurityRolesByParentIdSettingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProjectSecurityRolesByParentIdSettingsById`: ProjectSecurityRoleSetting
	fmt.Fprintf(os.Stdout, "Response from `ProjectSecurityRoleSettingsAPI.PatchProjectSecurityRolesByParentIdSettingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | settingId | 
**parentId** | **int32** | securityRoleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProjectSecurityRolesByParentIdSettingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ProjectSecurityRoleSetting**](ProjectSecurityRoleSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProjectSecurityRolesByParentIdSettingsById

> ProjectSecurityRoleSetting PutProjectSecurityRolesByParentIdSettingsById(ctx, id, parentId).ClientId(clientId).ProjectSecurityRoleSetting(projectSecurityRoleSetting).Execute()

Put ProjectSecurityRoleSetting

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
	id := int32(56) // int32 | settingId
	parentId := int32(56) // int32 | securityRoleId
	clientId := "clientId_example" // string | 
	projectSecurityRoleSetting := *openapiclient.NewProjectSecurityRoleSetting() // ProjectSecurityRoleSetting | projectSecurityRoleSetting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectSecurityRoleSettingsAPI.PutProjectSecurityRolesByParentIdSettingsById(context.Background(), id, parentId).ClientId(clientId).ProjectSecurityRoleSetting(projectSecurityRoleSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectSecurityRoleSettingsAPI.PutProjectSecurityRolesByParentIdSettingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProjectSecurityRolesByParentIdSettingsById`: ProjectSecurityRoleSetting
	fmt.Fprintf(os.Stdout, "Response from `ProjectSecurityRoleSettingsAPI.PutProjectSecurityRolesByParentIdSettingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | settingId | 
**parentId** | **int32** | securityRoleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProjectSecurityRolesByParentIdSettingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **projectSecurityRoleSetting** | [**ProjectSecurityRoleSetting**](ProjectSecurityRoleSetting.md) | projectSecurityRoleSetting | 

### Return type

[**ProjectSecurityRoleSetting**](ProjectSecurityRoleSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

