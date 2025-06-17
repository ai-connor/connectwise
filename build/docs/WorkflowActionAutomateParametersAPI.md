# \WorkflowActionAutomateParametersAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemWorkflowActionsByParentIdAutomateParametersById**](WorkflowActionAutomateParametersAPI.md#DeleteSystemWorkflowActionsByParentIdAutomateParametersById) | **Delete** /system/workflowActions/{parentId}/automateParameters/{id} | Delete WorkflowActionAutomateParameter
[**GetSystemWorkflowActionsAutomateParameters**](WorkflowActionAutomateParametersAPI.md#GetSystemWorkflowActionsAutomateParameters) | **Get** /system/workflowActions/automateParameters | Get List of WorkflowActionAutomateParameter
[**GetSystemWorkflowActionsAutomateParametersById**](WorkflowActionAutomateParametersAPI.md#GetSystemWorkflowActionsAutomateParametersById) | **Get** /system/workflowActions/automateParameters/{id} | Get WorkflowActionAutomateParameter
[**GetSystemWorkflowActionsByParentIdAutomateParameters**](WorkflowActionAutomateParametersAPI.md#GetSystemWorkflowActionsByParentIdAutomateParameters) | **Get** /system/workflowActions/{parentId}/automateParameters | Get List of WorkflowActionAutomateParameter
[**GetSystemWorkflowActionsByParentIdAutomateParametersById**](WorkflowActionAutomateParametersAPI.md#GetSystemWorkflowActionsByParentIdAutomateParametersById) | **Get** /system/workflowActions/{parentId}/automateParameters/{id} | Get WorkflowActionAutomateParameter
[**GetSystemWorkflowActionsByParentIdAutomateParametersCount**](WorkflowActionAutomateParametersAPI.md#GetSystemWorkflowActionsByParentIdAutomateParametersCount) | **Get** /system/workflowActions/{parentId}/automateParameters/count | Get Count of WorkflowActionAutomateParameter
[**PatchSystemWorkflowActionsByParentIdAutomateParametersById**](WorkflowActionAutomateParametersAPI.md#PatchSystemWorkflowActionsByParentIdAutomateParametersById) | **Patch** /system/workflowActions/{parentId}/automateParameters/{id} | Patch WorkflowActionAutomateParameter
[**PostSystemWorkflowActionsByParentIdAutomateParameters**](WorkflowActionAutomateParametersAPI.md#PostSystemWorkflowActionsByParentIdAutomateParameters) | **Post** /system/workflowActions/{parentId}/automateParameters | Post WorkflowActionAutomateParameter
[**PutSystemWorkflowActionsByParentIdAutomateParametersById**](WorkflowActionAutomateParametersAPI.md#PutSystemWorkflowActionsByParentIdAutomateParametersById) | **Put** /system/workflowActions/{parentId}/automateParameters/{id} | Put WorkflowActionAutomateParameter



## DeleteSystemWorkflowActionsByParentIdAutomateParametersById

> DeleteSystemWorkflowActionsByParentIdAutomateParametersById(ctx, id, parentId).ClientId(clientId).Execute()

Delete WorkflowActionAutomateParameter

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
	id := int32(56) // int32 | automateParameterId
	parentId := int32(56) // int32 | workflowActionId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowActionAutomateParametersAPI.DeleteSystemWorkflowActionsByParentIdAutomateParametersById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionAutomateParametersAPI.DeleteSystemWorkflowActionsByParentIdAutomateParametersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | automateParameterId | 
**parentId** | **int32** | workflowActionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemWorkflowActionsByParentIdAutomateParametersByIdRequest struct via the builder pattern


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


## GetSystemWorkflowActionsAutomateParameters

> []WorkflowActionAutomateParameter GetSystemWorkflowActionsAutomateParameters(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of WorkflowActionAutomateParameter

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
	resp, r, err := apiClient.WorkflowActionAutomateParametersAPI.GetSystemWorkflowActionsAutomateParameters(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionAutomateParametersAPI.GetSystemWorkflowActionsAutomateParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemWorkflowActionsAutomateParameters`: []WorkflowActionAutomateParameter
	fmt.Fprintf(os.Stdout, "Response from `WorkflowActionAutomateParametersAPI.GetSystemWorkflowActionsAutomateParameters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemWorkflowActionsAutomateParametersRequest struct via the builder pattern


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

[**[]WorkflowActionAutomateParameter**](WorkflowActionAutomateParameter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemWorkflowActionsAutomateParametersById

> WorkflowActionAutomateParameter GetSystemWorkflowActionsAutomateParametersById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get WorkflowActionAutomateParameter

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
	id := int32(56) // int32 | automateParameterId
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
	resp, r, err := apiClient.WorkflowActionAutomateParametersAPI.GetSystemWorkflowActionsAutomateParametersById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionAutomateParametersAPI.GetSystemWorkflowActionsAutomateParametersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemWorkflowActionsAutomateParametersById`: WorkflowActionAutomateParameter
	fmt.Fprintf(os.Stdout, "Response from `WorkflowActionAutomateParametersAPI.GetSystemWorkflowActionsAutomateParametersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | automateParameterId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemWorkflowActionsAutomateParametersByIdRequest struct via the builder pattern


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

[**WorkflowActionAutomateParameter**](WorkflowActionAutomateParameter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemWorkflowActionsByParentIdAutomateParameters

> []WorkflowActionAutomateParameter GetSystemWorkflowActionsByParentIdAutomateParameters(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of WorkflowActionAutomateParameter

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
	parentId := int32(56) // int32 | workflowActionId
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
	resp, r, err := apiClient.WorkflowActionAutomateParametersAPI.GetSystemWorkflowActionsByParentIdAutomateParameters(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionAutomateParametersAPI.GetSystemWorkflowActionsByParentIdAutomateParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemWorkflowActionsByParentIdAutomateParameters`: []WorkflowActionAutomateParameter
	fmt.Fprintf(os.Stdout, "Response from `WorkflowActionAutomateParametersAPI.GetSystemWorkflowActionsByParentIdAutomateParameters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | workflowActionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemWorkflowActionsByParentIdAutomateParametersRequest struct via the builder pattern


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

[**[]WorkflowActionAutomateParameter**](WorkflowActionAutomateParameter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemWorkflowActionsByParentIdAutomateParametersById

> WorkflowActionAutomateParameter GetSystemWorkflowActionsByParentIdAutomateParametersById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get WorkflowActionAutomateParameter

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
	id := int32(56) // int32 | automateParameterId
	parentId := int32(56) // int32 | workflowActionId
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
	resp, r, err := apiClient.WorkflowActionAutomateParametersAPI.GetSystemWorkflowActionsByParentIdAutomateParametersById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionAutomateParametersAPI.GetSystemWorkflowActionsByParentIdAutomateParametersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemWorkflowActionsByParentIdAutomateParametersById`: WorkflowActionAutomateParameter
	fmt.Fprintf(os.Stdout, "Response from `WorkflowActionAutomateParametersAPI.GetSystemWorkflowActionsByParentIdAutomateParametersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | automateParameterId | 
**parentId** | **int32** | workflowActionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemWorkflowActionsByParentIdAutomateParametersByIdRequest struct via the builder pattern


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

[**WorkflowActionAutomateParameter**](WorkflowActionAutomateParameter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemWorkflowActionsByParentIdAutomateParametersCount

> Count GetSystemWorkflowActionsByParentIdAutomateParametersCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of WorkflowActionAutomateParameter

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
	parentId := int32(56) // int32 | workflowActionId
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
	resp, r, err := apiClient.WorkflowActionAutomateParametersAPI.GetSystemWorkflowActionsByParentIdAutomateParametersCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionAutomateParametersAPI.GetSystemWorkflowActionsByParentIdAutomateParametersCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemWorkflowActionsByParentIdAutomateParametersCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `WorkflowActionAutomateParametersAPI.GetSystemWorkflowActionsByParentIdAutomateParametersCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | workflowActionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemWorkflowActionsByParentIdAutomateParametersCountRequest struct via the builder pattern


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


## PatchSystemWorkflowActionsByParentIdAutomateParametersById

> WorkflowActionAutomateParameter PatchSystemWorkflowActionsByParentIdAutomateParametersById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch WorkflowActionAutomateParameter

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
	id := int32(56) // int32 | automateParameterId
	parentId := int32(56) // int32 | workflowActionId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowActionAutomateParametersAPI.PatchSystemWorkflowActionsByParentIdAutomateParametersById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionAutomateParametersAPI.PatchSystemWorkflowActionsByParentIdAutomateParametersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemWorkflowActionsByParentIdAutomateParametersById`: WorkflowActionAutomateParameter
	fmt.Fprintf(os.Stdout, "Response from `WorkflowActionAutomateParametersAPI.PatchSystemWorkflowActionsByParentIdAutomateParametersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | automateParameterId | 
**parentId** | **int32** | workflowActionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemWorkflowActionsByParentIdAutomateParametersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**WorkflowActionAutomateParameter**](WorkflowActionAutomateParameter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemWorkflowActionsByParentIdAutomateParameters

> WorkflowActionAutomateParameter PostSystemWorkflowActionsByParentIdAutomateParameters(ctx, parentId).ClientId(clientId).WorkflowActionAutomateParameter(workflowActionAutomateParameter).Execute()

Post WorkflowActionAutomateParameter

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
	parentId := int32(56) // int32 | workflowActionId
	clientId := "clientId_example" // string | 
	workflowActionAutomateParameter := *openapiclient.NewWorkflowActionAutomateParameter("Name_example") // WorkflowActionAutomateParameter | workflowActionAutomateParameter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowActionAutomateParametersAPI.PostSystemWorkflowActionsByParentIdAutomateParameters(context.Background(), parentId).ClientId(clientId).WorkflowActionAutomateParameter(workflowActionAutomateParameter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionAutomateParametersAPI.PostSystemWorkflowActionsByParentIdAutomateParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemWorkflowActionsByParentIdAutomateParameters`: WorkflowActionAutomateParameter
	fmt.Fprintf(os.Stdout, "Response from `WorkflowActionAutomateParametersAPI.PostSystemWorkflowActionsByParentIdAutomateParameters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | workflowActionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemWorkflowActionsByParentIdAutomateParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **workflowActionAutomateParameter** | [**WorkflowActionAutomateParameter**](WorkflowActionAutomateParameter.md) | workflowActionAutomateParameter | 

### Return type

[**WorkflowActionAutomateParameter**](WorkflowActionAutomateParameter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemWorkflowActionsByParentIdAutomateParametersById

> WorkflowActionAutomateParameter PutSystemWorkflowActionsByParentIdAutomateParametersById(ctx, id, parentId).ClientId(clientId).WorkflowActionAutomateParameter(workflowActionAutomateParameter).Execute()

Put WorkflowActionAutomateParameter

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
	id := int32(56) // int32 | automateParameterId
	parentId := int32(56) // int32 | workflowActionId
	clientId := "clientId_example" // string | 
	workflowActionAutomateParameter := *openapiclient.NewWorkflowActionAutomateParameter("Name_example") // WorkflowActionAutomateParameter | workflowActionAutomateParameter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowActionAutomateParametersAPI.PutSystemWorkflowActionsByParentIdAutomateParametersById(context.Background(), id, parentId).ClientId(clientId).WorkflowActionAutomateParameter(workflowActionAutomateParameter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionAutomateParametersAPI.PutSystemWorkflowActionsByParentIdAutomateParametersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemWorkflowActionsByParentIdAutomateParametersById`: WorkflowActionAutomateParameter
	fmt.Fprintf(os.Stdout, "Response from `WorkflowActionAutomateParametersAPI.PutSystemWorkflowActionsByParentIdAutomateParametersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | automateParameterId | 
**parentId** | **int32** | workflowActionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemWorkflowActionsByParentIdAutomateParametersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **workflowActionAutomateParameter** | [**WorkflowActionAutomateParameter**](WorkflowActionAutomateParameter.md) | workflowActionAutomateParameter | 

### Return type

[**WorkflowActionAutomateParameter**](WorkflowActionAutomateParameter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

