# \ProjectTemplatePhasesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProjectProjectTemplatesByParentIdProjectTemplatePhasesById**](ProjectTemplatePhasesAPI.md#DeleteProjectProjectTemplatesByParentIdProjectTemplatePhasesById) | **Delete** /project/projectTemplates/{parentId}/projectTemplatePhases/{id} | Delete ProjectTemplatePhases
[**GetProjectProjectTemplatesByParentIdProjectTemplatePhases**](ProjectTemplatePhasesAPI.md#GetProjectProjectTemplatesByParentIdProjectTemplatePhases) | **Get** /project/projectTemplates/{parentId}/projectTemplatePhases | Get List of ProjectTemplatePhases
[**GetProjectProjectTemplatesByParentIdProjectTemplatePhasesById**](ProjectTemplatePhasesAPI.md#GetProjectProjectTemplatesByParentIdProjectTemplatePhasesById) | **Get** /project/projectTemplates/{parentId}/projectTemplatePhases/{id} | Get ProjectTemplatePhases
[**GetProjectProjectTemplatesProjectTemplatePhases**](ProjectTemplatePhasesAPI.md#GetProjectProjectTemplatesProjectTemplatePhases) | **Get** /project/projectTemplates/projectTemplatePhases | Get List of ProjectTemplatePhases
[**PatchProjectProjectTemplatesByParentIdProjectTemplatePhasesById**](ProjectTemplatePhasesAPI.md#PatchProjectProjectTemplatesByParentIdProjectTemplatePhasesById) | **Patch** /project/projectTemplates/{parentId}/projectTemplatePhases/{id} | Patch ProjectTemplatePhases
[**PostProjectProjectTemplatesByParentIdProjectTemplatePhases**](ProjectTemplatePhasesAPI.md#PostProjectProjectTemplatesByParentIdProjectTemplatePhases) | **Post** /project/projectTemplates/{parentId}/projectTemplatePhases | Post ProjectTemplatePhases
[**PutProjectProjectTemplatesByParentIdProjectTemplatePhasesById**](ProjectTemplatePhasesAPI.md#PutProjectProjectTemplatesByParentIdProjectTemplatePhasesById) | **Put** /project/projectTemplates/{parentId}/projectTemplatePhases/{id} | Put ProjectTemplatePhases



## DeleteProjectProjectTemplatesByParentIdProjectTemplatePhasesById

> DeleteProjectProjectTemplatesByParentIdProjectTemplatePhasesById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ProjectTemplatePhases

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
	id := int32(56) // int32 | ProjectTemplatePhaseId
	parentId := int32(56) // int32 | templateProjectId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectTemplatePhasesAPI.DeleteProjectProjectTemplatesByParentIdProjectTemplatePhasesById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatePhasesAPI.DeleteProjectProjectTemplatesByParentIdProjectTemplatePhasesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ProjectTemplatePhaseId | 
**parentId** | **int32** | templateProjectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectProjectTemplatesByParentIdProjectTemplatePhasesByIdRequest struct via the builder pattern


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


## GetProjectProjectTemplatesByParentIdProjectTemplatePhases

> []ProjectTemplatePhase GetProjectProjectTemplatesByParentIdProjectTemplatePhases(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProjectTemplatePhases

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
	parentId := int32(56) // int32 | templateProjectId
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
	resp, r, err := apiClient.ProjectTemplatePhasesAPI.GetProjectProjectTemplatesByParentIdProjectTemplatePhases(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatePhasesAPI.GetProjectProjectTemplatesByParentIdProjectTemplatePhases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectTemplatesByParentIdProjectTemplatePhases`: []ProjectTemplatePhase
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatePhasesAPI.GetProjectProjectTemplatesByParentIdProjectTemplatePhases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | templateProjectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectTemplatesByParentIdProjectTemplatePhasesRequest struct via the builder pattern


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

[**[]ProjectTemplatePhase**](ProjectTemplatePhase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectProjectTemplatesByParentIdProjectTemplatePhasesById

> ProjectTemplatePhase GetProjectProjectTemplatesByParentIdProjectTemplatePhasesById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ProjectTemplatePhases

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
	id := int32(56) // int32 | ProjectTemplatePhaseId
	parentId := int32(56) // int32 | templateProjectId
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
	resp, r, err := apiClient.ProjectTemplatePhasesAPI.GetProjectProjectTemplatesByParentIdProjectTemplatePhasesById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatePhasesAPI.GetProjectProjectTemplatesByParentIdProjectTemplatePhasesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectTemplatesByParentIdProjectTemplatePhasesById`: ProjectTemplatePhase
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatePhasesAPI.GetProjectProjectTemplatesByParentIdProjectTemplatePhasesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ProjectTemplatePhaseId | 
**parentId** | **int32** | templateProjectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectTemplatesByParentIdProjectTemplatePhasesByIdRequest struct via the builder pattern


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

[**ProjectTemplatePhase**](ProjectTemplatePhase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectProjectTemplatesProjectTemplatePhases

> []ProjectTemplatePhase GetProjectProjectTemplatesProjectTemplatePhases(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProjectTemplatePhases

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
	parentId := int32(56) // int32 | templateProjectId
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
	resp, r, err := apiClient.ProjectTemplatePhasesAPI.GetProjectProjectTemplatesProjectTemplatePhases(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatePhasesAPI.GetProjectProjectTemplatesProjectTemplatePhases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectTemplatesProjectTemplatePhases`: []ProjectTemplatePhase
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatePhasesAPI.GetProjectProjectTemplatesProjectTemplatePhases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | templateProjectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectTemplatesProjectTemplatePhasesRequest struct via the builder pattern


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

[**[]ProjectTemplatePhase**](ProjectTemplatePhase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchProjectProjectTemplatesByParentIdProjectTemplatePhasesById

> ProjectTemplatePhase PatchProjectProjectTemplatesByParentIdProjectTemplatePhasesById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ProjectTemplatePhases

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
	id := int32(56) // int32 | ProjectTemplatePhaseId
	parentId := int32(56) // int32 | templateProjectId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTemplatePhasesAPI.PatchProjectProjectTemplatesByParentIdProjectTemplatePhasesById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatePhasesAPI.PatchProjectProjectTemplatesByParentIdProjectTemplatePhasesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProjectProjectTemplatesByParentIdProjectTemplatePhasesById`: ProjectTemplatePhase
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatePhasesAPI.PatchProjectProjectTemplatesByParentIdProjectTemplatePhasesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ProjectTemplatePhaseId | 
**parentId** | **int32** | templateProjectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProjectProjectTemplatesByParentIdProjectTemplatePhasesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ProjectTemplatePhase**](ProjectTemplatePhase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectProjectTemplatesByParentIdProjectTemplatePhases

> ProjectTemplatePhase PostProjectProjectTemplatesByParentIdProjectTemplatePhases(ctx, parentId).ClientId(clientId).ProjectTemplatePhase(projectTemplatePhase).Execute()

Post ProjectTemplatePhases

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
	parentId := int32(56) // int32 | templateProjectId
	clientId := "clientId_example" // string | 
	projectTemplatePhase := *openapiclient.NewProjectTemplatePhase() // ProjectTemplatePhase | ProjectTemplatePhase

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTemplatePhasesAPI.PostProjectProjectTemplatesByParentIdProjectTemplatePhases(context.Background(), parentId).ClientId(clientId).ProjectTemplatePhase(projectTemplatePhase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatePhasesAPI.PostProjectProjectTemplatesByParentIdProjectTemplatePhases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectProjectTemplatesByParentIdProjectTemplatePhases`: ProjectTemplatePhase
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatePhasesAPI.PostProjectProjectTemplatesByParentIdProjectTemplatePhases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | templateProjectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectProjectTemplatesByParentIdProjectTemplatePhasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **projectTemplatePhase** | [**ProjectTemplatePhase**](ProjectTemplatePhase.md) | ProjectTemplatePhase | 

### Return type

[**ProjectTemplatePhase**](ProjectTemplatePhase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProjectProjectTemplatesByParentIdProjectTemplatePhasesById

> ProjectTemplatePhase PutProjectProjectTemplatesByParentIdProjectTemplatePhasesById(ctx, id, parentId).ClientId(clientId).ProjectTemplatePhase(projectTemplatePhase).Execute()

Put ProjectTemplatePhases

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
	id := int32(56) // int32 | ProjectTemplatePhaseId
	parentId := int32(56) // int32 | templateProjectId
	clientId := "clientId_example" // string | 
	projectTemplatePhase := *openapiclient.NewProjectTemplatePhase() // ProjectTemplatePhase | projectTemplatePhase

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTemplatePhasesAPI.PutProjectProjectTemplatesByParentIdProjectTemplatePhasesById(context.Background(), id, parentId).ClientId(clientId).ProjectTemplatePhase(projectTemplatePhase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatePhasesAPI.PutProjectProjectTemplatesByParentIdProjectTemplatePhasesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProjectProjectTemplatesByParentIdProjectTemplatePhasesById`: ProjectTemplatePhase
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatePhasesAPI.PutProjectProjectTemplatesByParentIdProjectTemplatePhasesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ProjectTemplatePhaseId | 
**parentId** | **int32** | templateProjectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProjectProjectTemplatesByParentIdProjectTemplatePhasesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **projectTemplatePhase** | [**ProjectTemplatePhase**](ProjectTemplatePhase.md) | projectTemplatePhase | 

### Return type

[**ProjectTemplatePhase**](ProjectTemplatePhase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

