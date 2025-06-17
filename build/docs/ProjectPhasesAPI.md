# \ProjectPhasesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProjectProjectsByParentIdPhasesById**](ProjectPhasesAPI.md#DeleteProjectProjectsByParentIdPhasesById) | **Delete** /project/projects/{parentId}/phases/{id} | Delete ProjectPhase
[**GetProjectProjectsByParentIdPhases**](ProjectPhasesAPI.md#GetProjectProjectsByParentIdPhases) | **Get** /project/projects/{parentId}/phases | Get List of ProjectPhase
[**GetProjectProjectsByParentIdPhasesById**](ProjectPhasesAPI.md#GetProjectProjectsByParentIdPhasesById) | **Get** /project/projects/{parentId}/phases/{id} | Get ProjectPhase
[**GetProjectProjectsByParentIdPhasesCount**](ProjectPhasesAPI.md#GetProjectProjectsByParentIdPhasesCount) | **Get** /project/projects/{parentId}/phases/count | Get Count of ProjectPhase
[**PatchProjectProjectsByParentIdPhasesById**](ProjectPhasesAPI.md#PatchProjectProjectsByParentIdPhasesById) | **Patch** /project/projects/{parentId}/phases/{id} | Patch ProjectPhase
[**PostProjectProjectsByParentIdPhases**](ProjectPhasesAPI.md#PostProjectProjectsByParentIdPhases) | **Post** /project/projects/{parentId}/phases | Post ProjectPhase
[**PutProjectProjectsByParentIdPhasesById**](ProjectPhasesAPI.md#PutProjectProjectsByParentIdPhasesById) | **Put** /project/projects/{parentId}/phases/{id} | Put ProjectPhase



## DeleteProjectProjectsByParentIdPhasesById

> DeleteProjectProjectsByParentIdPhasesById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ProjectPhase

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
	id := int32(56) // int32 | phasId
	parentId := int32(56) // int32 | projectId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectPhasesAPI.DeleteProjectProjectsByParentIdPhasesById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectPhasesAPI.DeleteProjectProjectsByParentIdPhasesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | phasId | 
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectProjectsByParentIdPhasesByIdRequest struct via the builder pattern


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


## GetProjectProjectsByParentIdPhases

> []ProjectPhase GetProjectProjectsByParentIdPhases(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProjectPhase

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
	resp, r, err := apiClient.ProjectPhasesAPI.GetProjectProjectsByParentIdPhases(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectPhasesAPI.GetProjectProjectsByParentIdPhases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectsByParentIdPhases`: []ProjectPhase
	fmt.Fprintf(os.Stdout, "Response from `ProjectPhasesAPI.GetProjectProjectsByParentIdPhases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectsByParentIdPhasesRequest struct via the builder pattern


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

[**[]ProjectPhase**](ProjectPhase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectProjectsByParentIdPhasesById

> ProjectPhase GetProjectProjectsByParentIdPhasesById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ProjectPhase

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
	id := int32(56) // int32 | phasId
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
	resp, r, err := apiClient.ProjectPhasesAPI.GetProjectProjectsByParentIdPhasesById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectPhasesAPI.GetProjectProjectsByParentIdPhasesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectsByParentIdPhasesById`: ProjectPhase
	fmt.Fprintf(os.Stdout, "Response from `ProjectPhasesAPI.GetProjectProjectsByParentIdPhasesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | phasId | 
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectsByParentIdPhasesByIdRequest struct via the builder pattern


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

[**ProjectPhase**](ProjectPhase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectProjectsByParentIdPhasesCount

> Count GetProjectProjectsByParentIdPhasesCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ProjectPhase

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
	resp, r, err := apiClient.ProjectPhasesAPI.GetProjectProjectsByParentIdPhasesCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectPhasesAPI.GetProjectProjectsByParentIdPhasesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectsByParentIdPhasesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ProjectPhasesAPI.GetProjectProjectsByParentIdPhasesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectsByParentIdPhasesCountRequest struct via the builder pattern


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


## PatchProjectProjectsByParentIdPhasesById

> ProjectPhase PatchProjectProjectsByParentIdPhasesById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ProjectPhase

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
	id := int32(56) // int32 | phasId
	parentId := int32(56) // int32 | projectId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectPhasesAPI.PatchProjectProjectsByParentIdPhasesById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectPhasesAPI.PatchProjectProjectsByParentIdPhasesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProjectProjectsByParentIdPhasesById`: ProjectPhase
	fmt.Fprintf(os.Stdout, "Response from `ProjectPhasesAPI.PatchProjectProjectsByParentIdPhasesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | phasId | 
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProjectProjectsByParentIdPhasesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ProjectPhase**](ProjectPhase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectProjectsByParentIdPhases

> ProjectPhase PostProjectProjectsByParentIdPhases(ctx, parentId).ClientId(clientId).ProjectPhase(projectPhase).Execute()

Post ProjectPhase

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
	projectPhase := *openapiclient.NewProjectPhase("Description_example") // ProjectPhase | projectPhase

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectPhasesAPI.PostProjectProjectsByParentIdPhases(context.Background(), parentId).ClientId(clientId).ProjectPhase(projectPhase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectPhasesAPI.PostProjectProjectsByParentIdPhases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectProjectsByParentIdPhases`: ProjectPhase
	fmt.Fprintf(os.Stdout, "Response from `ProjectPhasesAPI.PostProjectProjectsByParentIdPhases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectProjectsByParentIdPhasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **projectPhase** | [**ProjectPhase**](ProjectPhase.md) | projectPhase | 

### Return type

[**ProjectPhase**](ProjectPhase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProjectProjectsByParentIdPhasesById

> ProjectPhase PutProjectProjectsByParentIdPhasesById(ctx, id, parentId).ClientId(clientId).ProjectPhase(projectPhase).Execute()

Put ProjectPhase

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
	id := int32(56) // int32 | phasId
	parentId := int32(56) // int32 | projectId
	clientId := "clientId_example" // string | 
	projectPhase := *openapiclient.NewProjectPhase("Description_example") // ProjectPhase | projectPhase

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectPhasesAPI.PutProjectProjectsByParentIdPhasesById(context.Background(), id, parentId).ClientId(clientId).ProjectPhase(projectPhase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectPhasesAPI.PutProjectProjectsByParentIdPhasesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProjectProjectsByParentIdPhasesById`: ProjectPhase
	fmt.Fprintf(os.Stdout, "Response from `ProjectPhasesAPI.PutProjectProjectsByParentIdPhasesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | phasId | 
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProjectProjectsByParentIdPhasesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **projectPhase** | [**ProjectPhase**](ProjectPhase.md) | projectPhase | 

### Return type

[**ProjectPhase**](ProjectPhase.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

