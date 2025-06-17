# \ProjectTemplatesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProjectProjectTemplatesById**](ProjectTemplatesAPI.md#DeleteProjectProjectTemplatesById) | **Delete** /project/projectTemplates/{id} | Delete ProjectTemplates
[**GetProjectProjectTemplates**](ProjectTemplatesAPI.md#GetProjectProjectTemplates) | **Get** /project/projectTemplates/ | Get List of ProjectTemplates
[**GetProjectProjectTemplatesById**](ProjectTemplatesAPI.md#GetProjectProjectTemplatesById) | **Get** /project/projectTemplates/{id} | Get ProjectTemplates
[**GetProjectProjectTemplatesByIdWorkplan**](ProjectTemplatesAPI.md#GetProjectProjectTemplatesByIdWorkplan) | **Get** /project/projectTemplates/{id}/workplan | Get ProjectTemplatesWorkplan
[**GetProjectProjectTemplatesCount**](ProjectTemplatesAPI.md#GetProjectProjectTemplatesCount) | **Get** /project/projectTemplates/count | Get Count of ProjectTemplates
[**PatchProjectProjectTemplatesById**](ProjectTemplatesAPI.md#PatchProjectProjectTemplatesById) | **Patch** /project/projectTemplates/{id} | Patch ProjectTemplates
[**PostProjectProjectTemplates**](ProjectTemplatesAPI.md#PostProjectProjectTemplates) | **Post** /project/projectTemplates/ | Post ProjectTemplates
[**PostProjectProjectTemplatesCreateFromProjectById**](ProjectTemplatesAPI.md#PostProjectProjectTemplatesCreateFromProjectById) | **Post** /project/projectTemplates/createFromProject/{id} | Post CreateFromProject
[**PutProjectProjectTemplatesById**](ProjectTemplatesAPI.md#PutProjectProjectTemplatesById) | **Put** /project/projectTemplates/{id} | Put ProjectTemplates



## DeleteProjectProjectTemplatesById

> DeleteProjectProjectTemplatesById(ctx, id).ClientId(clientId).Execute()

Delete ProjectTemplates

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
	id := int32(56) // int32 | ProjectTemplateId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectTemplatesAPI.DeleteProjectProjectTemplatesById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatesAPI.DeleteProjectProjectTemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ProjectTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectProjectTemplatesByIdRequest struct via the builder pattern


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


## GetProjectProjectTemplates

> []ProjectTemplate GetProjectProjectTemplates(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProjectTemplates

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
	resp, r, err := apiClient.ProjectTemplatesAPI.GetProjectProjectTemplates(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatesAPI.GetProjectProjectTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectTemplates`: []ProjectTemplate
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatesAPI.GetProjectProjectTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectTemplatesRequest struct via the builder pattern


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

[**[]ProjectTemplate**](ProjectTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectProjectTemplatesById

> ProjectTemplate GetProjectProjectTemplatesById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ProjectTemplates

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
	id := int32(56) // int32 | ProjectTemplateId
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
	resp, r, err := apiClient.ProjectTemplatesAPI.GetProjectProjectTemplatesById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatesAPI.GetProjectProjectTemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectTemplatesById`: ProjectTemplate
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatesAPI.GetProjectProjectTemplatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ProjectTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectTemplatesByIdRequest struct via the builder pattern


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

[**ProjectTemplate**](ProjectTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectProjectTemplatesByIdWorkplan

> []ProjectTemplateWorkPlan GetProjectProjectTemplatesByIdWorkplan(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ProjectTemplatesWorkplan

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
	id := int32(56) // int32 | ProjectTemplateId
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
	resp, r, err := apiClient.ProjectTemplatesAPI.GetProjectProjectTemplatesByIdWorkplan(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatesAPI.GetProjectProjectTemplatesByIdWorkplan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectTemplatesByIdWorkplan`: []ProjectTemplateWorkPlan
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatesAPI.GetProjectProjectTemplatesByIdWorkplan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ProjectTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectTemplatesByIdWorkplanRequest struct via the builder pattern


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

[**[]ProjectTemplateWorkPlan**](ProjectTemplateWorkPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectProjectTemplatesCount

> Count GetProjectProjectTemplatesCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ProjectTemplates

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
	resp, r, err := apiClient.ProjectTemplatesAPI.GetProjectProjectTemplatesCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatesAPI.GetProjectProjectTemplatesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectTemplatesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatesAPI.GetProjectProjectTemplatesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectTemplatesCountRequest struct via the builder pattern


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


## PatchProjectProjectTemplatesById

> ProjectTemplate PatchProjectProjectTemplatesById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ProjectTemplates

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
	id := int32(56) // int32 | ProjectTemplateId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTemplatesAPI.PatchProjectProjectTemplatesById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatesAPI.PatchProjectProjectTemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProjectProjectTemplatesById`: ProjectTemplate
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatesAPI.PatchProjectProjectTemplatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ProjectTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProjectProjectTemplatesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ProjectTemplate**](ProjectTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectProjectTemplates

> ProjectTemplate PostProjectProjectTemplates(ctx).ClientId(clientId).ProjectTemplate(projectTemplate).Execute()

Post ProjectTemplates

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
	projectTemplate := *openapiclient.NewProjectTemplate("Name_example") // ProjectTemplate | ProjectTemplate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTemplatesAPI.PostProjectProjectTemplates(context.Background()).ClientId(clientId).ProjectTemplate(projectTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatesAPI.PostProjectProjectTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectProjectTemplates`: ProjectTemplate
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatesAPI.PostProjectProjectTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectProjectTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **projectTemplate** | [**ProjectTemplate**](ProjectTemplate.md) | ProjectTemplate | 

### Return type

[**ProjectTemplate**](ProjectTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectProjectTemplatesCreateFromProjectById

> ProjectTemplate PostProjectProjectTemplatesCreateFromProjectById(ctx, id).ClientId(clientId).ProjectTemplate(projectTemplate).Execute()

Post CreateFromProject

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
	id := int32(56) // int32 | projectId
	clientId := "clientId_example" // string | 
	projectTemplate := *openapiclient.NewProjectTemplate("Name_example") // ProjectTemplate | ProjectTemplate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTemplatesAPI.PostProjectProjectTemplatesCreateFromProjectById(context.Background(), id).ClientId(clientId).ProjectTemplate(projectTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatesAPI.PostProjectProjectTemplatesCreateFromProjectById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectProjectTemplatesCreateFromProjectById`: ProjectTemplate
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatesAPI.PostProjectProjectTemplatesCreateFromProjectById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectProjectTemplatesCreateFromProjectByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **projectTemplate** | [**ProjectTemplate**](ProjectTemplate.md) | ProjectTemplate | 

### Return type

[**ProjectTemplate**](ProjectTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProjectProjectTemplatesById

> ProjectTemplate PutProjectProjectTemplatesById(ctx, id).ClientId(clientId).ProjectTemplate(projectTemplate).Execute()

Put ProjectTemplates

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
	id := int32(56) // int32 | ProjectTemplateId
	clientId := "clientId_example" // string | 
	projectTemplate := *openapiclient.NewProjectTemplate("Name_example") // ProjectTemplate | companyTypeAssociation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTemplatesAPI.PutProjectProjectTemplatesById(context.Background(), id).ClientId(clientId).ProjectTemplate(projectTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTemplatesAPI.PutProjectProjectTemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProjectProjectTemplatesById`: ProjectTemplate
	fmt.Fprintf(os.Stdout, "Response from `ProjectTemplatesAPI.PutProjectProjectTemplatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ProjectTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProjectProjectTemplatesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **projectTemplate** | [**ProjectTemplate**](ProjectTemplate.md) | companyTypeAssociation | 

### Return type

[**ProjectTemplate**](ProjectTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

