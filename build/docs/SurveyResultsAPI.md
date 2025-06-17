# \SurveyResultsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServiceSurveysByParentIdResultsById**](SurveyResultsAPI.md#DeleteServiceSurveysByParentIdResultsById) | **Delete** /service/surveys/{parentId}/results/{id} | Delete SurveyResult
[**GetServiceSurveysByParentIdResults**](SurveyResultsAPI.md#GetServiceSurveysByParentIdResults) | **Get** /service/surveys/{parentId}/results | Get List of SurveyResult
[**GetServiceSurveysByParentIdResultsById**](SurveyResultsAPI.md#GetServiceSurveysByParentIdResultsById) | **Get** /service/surveys/{parentId}/results/{id} | Get SurveyResult
[**GetServiceSurveysByParentIdResultsCount**](SurveyResultsAPI.md#GetServiceSurveysByParentIdResultsCount) | **Get** /service/surveys/{parentId}/results/count | Get Count of SurveyResult
[**PatchServiceSurveysByParentIdResultsById**](SurveyResultsAPI.md#PatchServiceSurveysByParentIdResultsById) | **Patch** /service/surveys/{parentId}/results/{id} | Patch SurveyResult
[**PostServiceSurveysByParentIdResults**](SurveyResultsAPI.md#PostServiceSurveysByParentIdResults) | **Post** /service/surveys/{parentId}/results | Post SurveyResult
[**PutServiceSurveysByParentIdResultsById**](SurveyResultsAPI.md#PutServiceSurveysByParentIdResultsById) | **Put** /service/surveys/{parentId}/results/{id} | Put SurveyResult



## DeleteServiceSurveysByParentIdResultsById

> DeleteServiceSurveysByParentIdResultsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete SurveyResult

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
	id := int32(56) // int32 | resultId
	parentId := int32(56) // int32 | surveyId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SurveyResultsAPI.DeleteServiceSurveysByParentIdResultsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyResultsAPI.DeleteServiceSurveysByParentIdResultsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | resultId | 
**parentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceSurveysByParentIdResultsByIdRequest struct via the builder pattern


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


## GetServiceSurveysByParentIdResults

> []SurveyResult GetServiceSurveysByParentIdResults(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of SurveyResult

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
	parentId := int32(56) // int32 | surveyId
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
	resp, r, err := apiClient.SurveyResultsAPI.GetServiceSurveysByParentIdResults(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyResultsAPI.GetServiceSurveysByParentIdResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceSurveysByParentIdResults`: []SurveyResult
	fmt.Fprintf(os.Stdout, "Response from `SurveyResultsAPI.GetServiceSurveysByParentIdResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceSurveysByParentIdResultsRequest struct via the builder pattern


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

[**[]SurveyResult**](SurveyResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceSurveysByParentIdResultsById

> SurveyResult GetServiceSurveysByParentIdResultsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get SurveyResult

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
	id := int32(56) // int32 | resultId
	parentId := int32(56) // int32 | surveyId
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
	resp, r, err := apiClient.SurveyResultsAPI.GetServiceSurveysByParentIdResultsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyResultsAPI.GetServiceSurveysByParentIdResultsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceSurveysByParentIdResultsById`: SurveyResult
	fmt.Fprintf(os.Stdout, "Response from `SurveyResultsAPI.GetServiceSurveysByParentIdResultsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | resultId | 
**parentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceSurveysByParentIdResultsByIdRequest struct via the builder pattern


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

[**SurveyResult**](SurveyResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceSurveysByParentIdResultsCount

> Count GetServiceSurveysByParentIdResultsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of SurveyResult

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
	parentId := int32(56) // int32 | surveyId
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
	resp, r, err := apiClient.SurveyResultsAPI.GetServiceSurveysByParentIdResultsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyResultsAPI.GetServiceSurveysByParentIdResultsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceSurveysByParentIdResultsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `SurveyResultsAPI.GetServiceSurveysByParentIdResultsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceSurveysByParentIdResultsCountRequest struct via the builder pattern


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


## PatchServiceSurveysByParentIdResultsById

> SurveyResult PatchServiceSurveysByParentIdResultsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch SurveyResult

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
	id := int32(56) // int32 | resultId
	parentId := int32(56) // int32 | surveyId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveyResultsAPI.PatchServiceSurveysByParentIdResultsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyResultsAPI.PatchServiceSurveysByParentIdResultsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchServiceSurveysByParentIdResultsById`: SurveyResult
	fmt.Fprintf(os.Stdout, "Response from `SurveyResultsAPI.PatchServiceSurveysByParentIdResultsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | resultId | 
**parentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchServiceSurveysByParentIdResultsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**SurveyResult**](SurveyResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServiceSurveysByParentIdResults

> SurveyResult PostServiceSurveysByParentIdResults(ctx, parentId).ClientId(clientId).SurveyResult(surveyResult).Execute()

Post SurveyResult

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
	parentId := int32(56) // int32 | surveyId
	clientId := "clientId_example" // string | 
	surveyResult := *openapiclient.NewSurveyResult(NullableInt32(123)) // SurveyResult | surveyResult

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveyResultsAPI.PostServiceSurveysByParentIdResults(context.Background(), parentId).ClientId(clientId).SurveyResult(surveyResult).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyResultsAPI.PostServiceSurveysByParentIdResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServiceSurveysByParentIdResults`: SurveyResult
	fmt.Fprintf(os.Stdout, "Response from `SurveyResultsAPI.PostServiceSurveysByParentIdResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServiceSurveysByParentIdResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **surveyResult** | [**SurveyResult**](SurveyResult.md) | surveyResult | 

### Return type

[**SurveyResult**](SurveyResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServiceSurveysByParentIdResultsById

> SurveyResult PutServiceSurveysByParentIdResultsById(ctx, id, parentId).ClientId(clientId).SurveyResult(surveyResult).Execute()

Put SurveyResult

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
	id := int32(56) // int32 | resultId
	parentId := int32(56) // int32 | surveyId
	clientId := "clientId_example" // string | 
	surveyResult := *openapiclient.NewSurveyResult(NullableInt32(123)) // SurveyResult | surveyResult

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveyResultsAPI.PutServiceSurveysByParentIdResultsById(context.Background(), id, parentId).ClientId(clientId).SurveyResult(surveyResult).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyResultsAPI.PutServiceSurveysByParentIdResultsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServiceSurveysByParentIdResultsById`: SurveyResult
	fmt.Fprintf(os.Stdout, "Response from `SurveyResultsAPI.PutServiceSurveysByParentIdResultsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | resultId | 
**parentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServiceSurveysByParentIdResultsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **surveyResult** | [**SurveyResult**](SurveyResult.md) | surveyResult | 

### Return type

[**SurveyResult**](SurveyResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

