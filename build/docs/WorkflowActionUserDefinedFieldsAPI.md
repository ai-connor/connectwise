# \WorkflowActionUserDefinedFieldsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemWorkflowsUserdefinedfieldsActionsByParentId**](WorkflowActionUserDefinedFieldsAPI.md#DeleteSystemWorkflowsUserdefinedfieldsActionsByParentId) | **Delete** /system/workflows/userdefinedfields/actions/{parentId} | Delete WorkflowActionUserDefinedField
[**GetSystemWorkflowsUserdefinedfieldsEventsActions**](WorkflowActionUserDefinedFieldsAPI.md#GetSystemWorkflowsUserdefinedfieldsEventsActions) | **Get** /system/workflows/userdefinedfields/events/actions | Get List of WorkflowActionUserDefinedField
[**GetSystemWorkflowsUserdefinedfieldsEventsByGrandparentIdActionsByParentId**](WorkflowActionUserDefinedFieldsAPI.md#GetSystemWorkflowsUserdefinedfieldsEventsByGrandparentIdActionsByParentId) | **Get** /system/workflows/userdefinedfields/events/{grandparentId}/actions/{parentId} | Get List of WorkflowActionUserDefinedField
[**PatchSystemWorkflowsUserdefinedfieldsById**](WorkflowActionUserDefinedFieldsAPI.md#PatchSystemWorkflowsUserdefinedfieldsById) | **Patch** /system/workflows/userdefinedfields/{id} | Patch WorkflowActionUserDefinedField
[**PostSystemWorkflowsUserdefinedfieldsEventsByGrandparentId**](WorkflowActionUserDefinedFieldsAPI.md#PostSystemWorkflowsUserdefinedfieldsEventsByGrandparentId) | **Post** /system/workflows/userdefinedfields/events/{grandparentId} | Post WorkflowActionUserDefinedField
[**PutSystemWorkflowsUserdefinedfieldsById**](WorkflowActionUserDefinedFieldsAPI.md#PutSystemWorkflowsUserdefinedfieldsById) | **Put** /system/workflows/userdefinedfields/{id} | Put WorkflowActionUserDefinedField



## DeleteSystemWorkflowsUserdefinedfieldsActionsByParentId

> DeleteSystemWorkflowsUserdefinedfieldsActionsByParentId(ctx, parentId).ClientId(clientId).Execute()

Delete WorkflowActionUserDefinedField

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
	parentId := int32(56) // int32 | actionId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowActionUserDefinedFieldsAPI.DeleteSystemWorkflowsUserdefinedfieldsActionsByParentId(context.Background(), parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionUserDefinedFieldsAPI.DeleteSystemWorkflowsUserdefinedfieldsActionsByParentId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | actionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemWorkflowsUserdefinedfieldsActionsByParentIdRequest struct via the builder pattern


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


## GetSystemWorkflowsUserdefinedfieldsEventsActions

> []WorkflowActionUserDefinedField GetSystemWorkflowsUserdefinedfieldsEventsActions(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of WorkflowActionUserDefinedField

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
	resp, r, err := apiClient.WorkflowActionUserDefinedFieldsAPI.GetSystemWorkflowsUserdefinedfieldsEventsActions(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionUserDefinedFieldsAPI.GetSystemWorkflowsUserdefinedfieldsEventsActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemWorkflowsUserdefinedfieldsEventsActions`: []WorkflowActionUserDefinedField
	fmt.Fprintf(os.Stdout, "Response from `WorkflowActionUserDefinedFieldsAPI.GetSystemWorkflowsUserdefinedfieldsEventsActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemWorkflowsUserdefinedfieldsEventsActionsRequest struct via the builder pattern


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

[**[]WorkflowActionUserDefinedField**](WorkflowActionUserDefinedField.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemWorkflowsUserdefinedfieldsEventsByGrandparentIdActionsByParentId

> []WorkflowActionUserDefinedField GetSystemWorkflowsUserdefinedfieldsEventsByGrandparentIdActionsByParentId(ctx, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of WorkflowActionUserDefinedField

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
	parentId := int32(56) // int32 | actionId
	grandparentId := int32(56) // int32 | evnetId
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
	resp, r, err := apiClient.WorkflowActionUserDefinedFieldsAPI.GetSystemWorkflowsUserdefinedfieldsEventsByGrandparentIdActionsByParentId(context.Background(), parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionUserDefinedFieldsAPI.GetSystemWorkflowsUserdefinedfieldsEventsByGrandparentIdActionsByParentId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemWorkflowsUserdefinedfieldsEventsByGrandparentIdActionsByParentId`: []WorkflowActionUserDefinedField
	fmt.Fprintf(os.Stdout, "Response from `WorkflowActionUserDefinedFieldsAPI.GetSystemWorkflowsUserdefinedfieldsEventsByGrandparentIdActionsByParentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | actionId | 
**grandparentId** | **int32** | evnetId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemWorkflowsUserdefinedfieldsEventsByGrandparentIdActionsByParentIdRequest struct via the builder pattern


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

[**[]WorkflowActionUserDefinedField**](WorkflowActionUserDefinedField.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSystemWorkflowsUserdefinedfieldsById

> WorkflowActionUserDefinedField PatchSystemWorkflowsUserdefinedfieldsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch WorkflowActionUserDefinedField

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
	id := int32(56) // int32 | workflowActionUserDefinedFieldId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowActionUserDefinedFieldsAPI.PatchSystemWorkflowsUserdefinedfieldsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionUserDefinedFieldsAPI.PatchSystemWorkflowsUserdefinedfieldsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemWorkflowsUserdefinedfieldsById`: WorkflowActionUserDefinedField
	fmt.Fprintf(os.Stdout, "Response from `WorkflowActionUserDefinedFieldsAPI.PatchSystemWorkflowsUserdefinedfieldsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | workflowActionUserDefinedFieldId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemWorkflowsUserdefinedfieldsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**WorkflowActionUserDefinedField**](WorkflowActionUserDefinedField.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemWorkflowsUserdefinedfieldsEventsByGrandparentId

> WorkflowActionUserDefinedField PostSystemWorkflowsUserdefinedfieldsEventsByGrandparentId(ctx, grandparentId).ClientId(clientId).WorkflowActionUserDefinedField(workflowActionUserDefinedField).Execute()

Post WorkflowActionUserDefinedField

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
	grandparentId := int32(56) // int32 | eventId
	clientId := "clientId_example" // string | 
	workflowActionUserDefinedField := *openapiclient.NewWorkflowActionUserDefinedField() // WorkflowActionUserDefinedField | workflowActionUserDefinedField

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowActionUserDefinedFieldsAPI.PostSystemWorkflowsUserdefinedfieldsEventsByGrandparentId(context.Background(), grandparentId).ClientId(clientId).WorkflowActionUserDefinedField(workflowActionUserDefinedField).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionUserDefinedFieldsAPI.PostSystemWorkflowsUserdefinedfieldsEventsByGrandparentId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemWorkflowsUserdefinedfieldsEventsByGrandparentId`: WorkflowActionUserDefinedField
	fmt.Fprintf(os.Stdout, "Response from `WorkflowActionUserDefinedFieldsAPI.PostSystemWorkflowsUserdefinedfieldsEventsByGrandparentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**grandparentId** | **int32** | eventId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemWorkflowsUserdefinedfieldsEventsByGrandparentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **workflowActionUserDefinedField** | [**WorkflowActionUserDefinedField**](WorkflowActionUserDefinedField.md) | workflowActionUserDefinedField | 

### Return type

[**WorkflowActionUserDefinedField**](WorkflowActionUserDefinedField.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemWorkflowsUserdefinedfieldsById

> WorkflowActionUserDefinedField PutSystemWorkflowsUserdefinedfieldsById(ctx, id).ClientId(clientId).WorkflowActionUserDefinedField(workflowActionUserDefinedField).Execute()

Put WorkflowActionUserDefinedField

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
	id := int32(56) // int32 | workflowActionUserDefinedFieldId
	clientId := "clientId_example" // string | 
	workflowActionUserDefinedField := *openapiclient.NewWorkflowActionUserDefinedField() // WorkflowActionUserDefinedField | workflowActionUserDefinedField

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowActionUserDefinedFieldsAPI.PutSystemWorkflowsUserdefinedfieldsById(context.Background(), id).ClientId(clientId).WorkflowActionUserDefinedField(workflowActionUserDefinedField).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionUserDefinedFieldsAPI.PutSystemWorkflowsUserdefinedfieldsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemWorkflowsUserdefinedfieldsById`: WorkflowActionUserDefinedField
	fmt.Fprintf(os.Stdout, "Response from `WorkflowActionUserDefinedFieldsAPI.PutSystemWorkflowsUserdefinedfieldsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | workflowActionUserDefinedFieldId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemWorkflowsUserdefinedfieldsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **workflowActionUserDefinedField** | [**WorkflowActionUserDefinedField**](WorkflowActionUserDefinedField.md) | workflowActionUserDefinedField | 

### Return type

[**WorkflowActionUserDefinedField**](WorkflowActionUserDefinedField.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

