# \ProjectBillingRatesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProjectByParentIdBillingRatesById**](ProjectBillingRatesAPI.md#DeleteProjectByParentIdBillingRatesById) | **Delete** /project/{parentId}/billingRates/{id} | Delete ProjectBillingRate
[**GetProjectByParentIdBillingRates**](ProjectBillingRatesAPI.md#GetProjectByParentIdBillingRates) | **Get** /project/{parentId}/billingRates | Get List of ProjectBillingRate
[**GetProjectByParentIdBillingRatesById**](ProjectBillingRatesAPI.md#GetProjectByParentIdBillingRatesById) | **Get** /project/{parentId}/billingRates/{id} | Get ProjectBillingRate
[**GetProjectByParentIdBillingRatesCount**](ProjectBillingRatesAPI.md#GetProjectByParentIdBillingRatesCount) | **Get** /project/{parentId}/billingRates/count | Get Count of ProjectBillingRate
[**PatchProjectBillingRatesByParentIdBillingRatesById**](ProjectBillingRatesAPI.md#PatchProjectBillingRatesByParentIdBillingRatesById) | **Patch** /project/billingRates/{parentId}/billingRates/{id} | Patch ProjectBillingRate
[**PostProjectByParentIdBillingRates**](ProjectBillingRatesAPI.md#PostProjectByParentIdBillingRates) | **Post** /project/{parentId}/billingRates | Post ProjectBillingRate
[**PutProjectByParentIdBillingRatesById**](ProjectBillingRatesAPI.md#PutProjectByParentIdBillingRatesById) | **Put** /project/{parentId}/billingRates/{id} | Put ProjectBillingRate



## DeleteProjectByParentIdBillingRatesById

> DeleteProjectByParentIdBillingRatesById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ProjectBillingRate

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
	id := int32(56) // int32 | billingRate
	parentId := int32(56) // int32 | projectId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectBillingRatesAPI.DeleteProjectByParentIdBillingRatesById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBillingRatesAPI.DeleteProjectByParentIdBillingRatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | billingRate | 
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectByParentIdBillingRatesByIdRequest struct via the builder pattern


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


## GetProjectByParentIdBillingRates

> []ProjectBillingRate GetProjectByParentIdBillingRates(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProjectBillingRate

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
	resp, r, err := apiClient.ProjectBillingRatesAPI.GetProjectByParentIdBillingRates(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBillingRatesAPI.GetProjectByParentIdBillingRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectByParentIdBillingRates`: []ProjectBillingRate
	fmt.Fprintf(os.Stdout, "Response from `ProjectBillingRatesAPI.GetProjectByParentIdBillingRates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectByParentIdBillingRatesRequest struct via the builder pattern


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

[**[]ProjectBillingRate**](ProjectBillingRate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectByParentIdBillingRatesById

> ProjectBillingRate GetProjectByParentIdBillingRatesById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ProjectBillingRate

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
	id := int32(56) // int32 | billingRateId
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
	resp, r, err := apiClient.ProjectBillingRatesAPI.GetProjectByParentIdBillingRatesById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBillingRatesAPI.GetProjectByParentIdBillingRatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectByParentIdBillingRatesById`: ProjectBillingRate
	fmt.Fprintf(os.Stdout, "Response from `ProjectBillingRatesAPI.GetProjectByParentIdBillingRatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | billingRateId | 
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectByParentIdBillingRatesByIdRequest struct via the builder pattern


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

[**ProjectBillingRate**](ProjectBillingRate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectByParentIdBillingRatesCount

> Count GetProjectByParentIdBillingRatesCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ProjectBillingRate

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
	resp, r, err := apiClient.ProjectBillingRatesAPI.GetProjectByParentIdBillingRatesCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBillingRatesAPI.GetProjectByParentIdBillingRatesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectByParentIdBillingRatesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ProjectBillingRatesAPI.GetProjectByParentIdBillingRatesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectByParentIdBillingRatesCountRequest struct via the builder pattern


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


## PatchProjectBillingRatesByParentIdBillingRatesById

> ProjectBillingRate PatchProjectBillingRatesByParentIdBillingRatesById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ProjectBillingRate

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
	id := int32(56) // int32 | billingRateId
	parentId := int32(56) // int32 | projectId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectBillingRatesAPI.PatchProjectBillingRatesByParentIdBillingRatesById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBillingRatesAPI.PatchProjectBillingRatesByParentIdBillingRatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProjectBillingRatesByParentIdBillingRatesById`: ProjectBillingRate
	fmt.Fprintf(os.Stdout, "Response from `ProjectBillingRatesAPI.PatchProjectBillingRatesByParentIdBillingRatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | billingRateId | 
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProjectBillingRatesByParentIdBillingRatesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ProjectBillingRate**](ProjectBillingRate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectByParentIdBillingRates

> ProjectBillingRate PostProjectByParentIdBillingRates(ctx, parentId).ClientId(clientId).ProjectBillingRate(projectBillingRate).Execute()

Post ProjectBillingRate

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
	projectBillingRate := *openapiclient.NewProjectBillingRate() // ProjectBillingRate | billingRate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectBillingRatesAPI.PostProjectByParentIdBillingRates(context.Background(), parentId).ClientId(clientId).ProjectBillingRate(projectBillingRate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBillingRatesAPI.PostProjectByParentIdBillingRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectByParentIdBillingRates`: ProjectBillingRate
	fmt.Fprintf(os.Stdout, "Response from `ProjectBillingRatesAPI.PostProjectByParentIdBillingRates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectByParentIdBillingRatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **projectBillingRate** | [**ProjectBillingRate**](ProjectBillingRate.md) | billingRate | 

### Return type

[**ProjectBillingRate**](ProjectBillingRate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProjectByParentIdBillingRatesById

> ProjectBillingRate PutProjectByParentIdBillingRatesById(ctx, id, parentId).ClientId(clientId).ProjectBillingRate(projectBillingRate).Execute()

Put ProjectBillingRate

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
	id := int32(56) // int32 | billingRateId
	parentId := int32(56) // int32 | projectId
	clientId := "clientId_example" // string | 
	projectBillingRate := *openapiclient.NewProjectBillingRate() // ProjectBillingRate | billingRate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectBillingRatesAPI.PutProjectByParentIdBillingRatesById(context.Background(), id, parentId).ClientId(clientId).ProjectBillingRate(projectBillingRate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectBillingRatesAPI.PutProjectByParentIdBillingRatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProjectByParentIdBillingRatesById`: ProjectBillingRate
	fmt.Fprintf(os.Stdout, "Response from `ProjectBillingRatesAPI.PutProjectByParentIdBillingRatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | billingRateId | 
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProjectByParentIdBillingRatesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **projectBillingRate** | [**ProjectBillingRate**](ProjectBillingRate.md) | billingRate | 

### Return type

[**ProjectBillingRate**](ProjectBillingRate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

