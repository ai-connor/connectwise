# \CustomReportParametersAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemCustomReportsByParentIdParametersById**](CustomReportParametersAPI.md#DeleteSystemCustomReportsByParentIdParametersById) | **Delete** /system/customReports/{parentId}/parameters/{id} | Delete CustomReportParameter
[**GetSystemCustomReportsByParentIdParameters**](CustomReportParametersAPI.md#GetSystemCustomReportsByParentIdParameters) | **Get** /system/customReports/{parentId}/parameters | Get List of CustomReportParameter
[**GetSystemCustomReportsByParentIdParametersById**](CustomReportParametersAPI.md#GetSystemCustomReportsByParentIdParametersById) | **Get** /system/customReports/{parentId}/parameters/{id} | Get CustomReportParameter
[**GetSystemCustomReportsByParentIdParametersCount**](CustomReportParametersAPI.md#GetSystemCustomReportsByParentIdParametersCount) | **Get** /system/customReports/{parentId}/parameters/count | Get Count of CustomReportParameter
[**PatchSystemCustomReportsByParentIdParametersById**](CustomReportParametersAPI.md#PatchSystemCustomReportsByParentIdParametersById) | **Patch** /system/customReports/{parentId}/parameters/{id} | Patch CustomReportParameter
[**PostSystemCustomReportsByParentIdParameters**](CustomReportParametersAPI.md#PostSystemCustomReportsByParentIdParameters) | **Post** /system/customReports/{parentId}/parameters | Post CustomReportParameter
[**PutSystemCustomReportsByParentIdParametersById**](CustomReportParametersAPI.md#PutSystemCustomReportsByParentIdParametersById) | **Put** /system/customReports/{parentId}/parameters/{id} | Put CustomReportParameter



## DeleteSystemCustomReportsByParentIdParametersById

> DeleteSystemCustomReportsByParentIdParametersById(ctx, id, parentId).ClientId(clientId).Execute()

Delete CustomReportParameter

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
	id := int32(56) // int32 | parameterId
	parentId := int32(56) // int32 | customReportId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomReportParametersAPI.DeleteSystemCustomReportsByParentIdParametersById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomReportParametersAPI.DeleteSystemCustomReportsByParentIdParametersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | parameterId | 
**parentId** | **int32** | customReportId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemCustomReportsByParentIdParametersByIdRequest struct via the builder pattern


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


## GetSystemCustomReportsByParentIdParameters

> []CustomReportParameter GetSystemCustomReportsByParentIdParameters(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of CustomReportParameter

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
	parentId := int32(56) // int32 | customReportId
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
	resp, r, err := apiClient.CustomReportParametersAPI.GetSystemCustomReportsByParentIdParameters(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomReportParametersAPI.GetSystemCustomReportsByParentIdParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemCustomReportsByParentIdParameters`: []CustomReportParameter
	fmt.Fprintf(os.Stdout, "Response from `CustomReportParametersAPI.GetSystemCustomReportsByParentIdParameters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | customReportId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemCustomReportsByParentIdParametersRequest struct via the builder pattern


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

[**[]CustomReportParameter**](CustomReportParameter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemCustomReportsByParentIdParametersById

> CustomReportParameter GetSystemCustomReportsByParentIdParametersById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get CustomReportParameter

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
	id := int32(56) // int32 | parameterId
	parentId := int32(56) // int32 | customReportId
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
	resp, r, err := apiClient.CustomReportParametersAPI.GetSystemCustomReportsByParentIdParametersById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomReportParametersAPI.GetSystemCustomReportsByParentIdParametersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemCustomReportsByParentIdParametersById`: CustomReportParameter
	fmt.Fprintf(os.Stdout, "Response from `CustomReportParametersAPI.GetSystemCustomReportsByParentIdParametersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | parameterId | 
**parentId** | **int32** | customReportId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemCustomReportsByParentIdParametersByIdRequest struct via the builder pattern


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

[**CustomReportParameter**](CustomReportParameter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemCustomReportsByParentIdParametersCount

> Count GetSystemCustomReportsByParentIdParametersCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of CustomReportParameter

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
	parentId := int32(56) // int32 | customReportId
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
	resp, r, err := apiClient.CustomReportParametersAPI.GetSystemCustomReportsByParentIdParametersCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomReportParametersAPI.GetSystemCustomReportsByParentIdParametersCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemCustomReportsByParentIdParametersCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `CustomReportParametersAPI.GetSystemCustomReportsByParentIdParametersCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | customReportId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemCustomReportsByParentIdParametersCountRequest struct via the builder pattern


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


## PatchSystemCustomReportsByParentIdParametersById

> CustomReportParameter PatchSystemCustomReportsByParentIdParametersById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch CustomReportParameter

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
	id := int32(56) // int32 | parameterId
	parentId := int32(56) // int32 | customReportId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomReportParametersAPI.PatchSystemCustomReportsByParentIdParametersById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomReportParametersAPI.PatchSystemCustomReportsByParentIdParametersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemCustomReportsByParentIdParametersById`: CustomReportParameter
	fmt.Fprintf(os.Stdout, "Response from `CustomReportParametersAPI.PatchSystemCustomReportsByParentIdParametersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | parameterId | 
**parentId** | **int32** | customReportId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemCustomReportsByParentIdParametersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**CustomReportParameter**](CustomReportParameter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemCustomReportsByParentIdParameters

> CustomReportParameter PostSystemCustomReportsByParentIdParameters(ctx, parentId).ClientId(clientId).CustomReportParameter(customReportParameter).Execute()

Post CustomReportParameter

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
	parentId := int32(56) // int32 | customReportId
	clientId := "clientId_example" // string | 
	customReportParameter := *openapiclient.NewCustomReportParameter() // CustomReportParameter | customReportParameter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomReportParametersAPI.PostSystemCustomReportsByParentIdParameters(context.Background(), parentId).ClientId(clientId).CustomReportParameter(customReportParameter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomReportParametersAPI.PostSystemCustomReportsByParentIdParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemCustomReportsByParentIdParameters`: CustomReportParameter
	fmt.Fprintf(os.Stdout, "Response from `CustomReportParametersAPI.PostSystemCustomReportsByParentIdParameters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | customReportId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemCustomReportsByParentIdParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **customReportParameter** | [**CustomReportParameter**](CustomReportParameter.md) | customReportParameter | 

### Return type

[**CustomReportParameter**](CustomReportParameter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemCustomReportsByParentIdParametersById

> CustomReportParameter PutSystemCustomReportsByParentIdParametersById(ctx, id, parentId).ClientId(clientId).CustomReportParameter(customReportParameter).Execute()

Put CustomReportParameter

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
	id := int32(56) // int32 | parameterId
	parentId := int32(56) // int32 | customReportId
	clientId := "clientId_example" // string | 
	customReportParameter := *openapiclient.NewCustomReportParameter() // CustomReportParameter | customReportParameter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomReportParametersAPI.PutSystemCustomReportsByParentIdParametersById(context.Background(), id, parentId).ClientId(clientId).CustomReportParameter(customReportParameter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomReportParametersAPI.PutSystemCustomReportsByParentIdParametersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemCustomReportsByParentIdParametersById`: CustomReportParameter
	fmt.Fprintf(os.Stdout, "Response from `CustomReportParametersAPI.PutSystemCustomReportsByParentIdParametersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | parameterId | 
**parentId** | **int32** | customReportId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemCustomReportsByParentIdParametersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **customReportParameter** | [**CustomReportParameter**](CustomReportParameter.md) | customReportParameter | 

### Return type

[**CustomReportParameter**](CustomReportParameter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

