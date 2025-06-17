# \SurveyQuestionsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemSurveysByParentIdQuestionsById**](SurveyQuestionsAPI.md#DeleteSystemSurveysByParentIdQuestionsById) | **Delete** /system/surveys/{parentId}/questions/{id} | Delete SurveyQuestion
[**GetSystemSurveysByParentIdQuestions**](SurveyQuestionsAPI.md#GetSystemSurveysByParentIdQuestions) | **Get** /system/surveys/{parentId}/questions | Get List of SurveyQuestion
[**GetSystemSurveysByParentIdQuestionsById**](SurveyQuestionsAPI.md#GetSystemSurveysByParentIdQuestionsById) | **Get** /system/surveys/{parentId}/questions/{id} | Get SurveyQuestion
[**GetSystemSurveysByParentIdQuestionsCount**](SurveyQuestionsAPI.md#GetSystemSurveysByParentIdQuestionsCount) | **Get** /system/surveys/{parentId}/questions/count | Get Count of SurveyQuestion
[**PatchSystemSurveysByParentIdQuestionsById**](SurveyQuestionsAPI.md#PatchSystemSurveysByParentIdQuestionsById) | **Patch** /system/surveys/{parentId}/questions/{id} | Patch SurveyQuestion
[**PostSystemSurveysByParentIdQuestions**](SurveyQuestionsAPI.md#PostSystemSurveysByParentIdQuestions) | **Post** /system/surveys/{parentId}/questions | Post SurveyQuestion
[**PutSystemSurveysByParentIdQuestionsById**](SurveyQuestionsAPI.md#PutSystemSurveysByParentIdQuestionsById) | **Put** /system/surveys/{parentId}/questions/{id} | Put SurveyQuestion



## DeleteSystemSurveysByParentIdQuestionsById

> DeleteSystemSurveysByParentIdQuestionsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete SurveyQuestion

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
	id := int32(56) // int32 | questionId
	parentId := int32(56) // int32 | surveyId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SurveyQuestionsAPI.DeleteSystemSurveysByParentIdQuestionsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyQuestionsAPI.DeleteSystemSurveysByParentIdQuestionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | questionId | 
**parentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemSurveysByParentIdQuestionsByIdRequest struct via the builder pattern


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


## GetSystemSurveysByParentIdQuestions

> []SurveyQuestion GetSystemSurveysByParentIdQuestions(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of SurveyQuestion

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
	resp, r, err := apiClient.SurveyQuestionsAPI.GetSystemSurveysByParentIdQuestions(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyQuestionsAPI.GetSystemSurveysByParentIdQuestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemSurveysByParentIdQuestions`: []SurveyQuestion
	fmt.Fprintf(os.Stdout, "Response from `SurveyQuestionsAPI.GetSystemSurveysByParentIdQuestions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemSurveysByParentIdQuestionsRequest struct via the builder pattern


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

[**[]SurveyQuestion**](SurveyQuestion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemSurveysByParentIdQuestionsById

> SurveyQuestion GetSystemSurveysByParentIdQuestionsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get SurveyQuestion

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
	id := int32(56) // int32 | questionId
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
	resp, r, err := apiClient.SurveyQuestionsAPI.GetSystemSurveysByParentIdQuestionsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyQuestionsAPI.GetSystemSurveysByParentIdQuestionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemSurveysByParentIdQuestionsById`: SurveyQuestion
	fmt.Fprintf(os.Stdout, "Response from `SurveyQuestionsAPI.GetSystemSurveysByParentIdQuestionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | questionId | 
**parentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemSurveysByParentIdQuestionsByIdRequest struct via the builder pattern


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

[**SurveyQuestion**](SurveyQuestion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemSurveysByParentIdQuestionsCount

> Count GetSystemSurveysByParentIdQuestionsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of SurveyQuestion

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
	resp, r, err := apiClient.SurveyQuestionsAPI.GetSystemSurveysByParentIdQuestionsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyQuestionsAPI.GetSystemSurveysByParentIdQuestionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemSurveysByParentIdQuestionsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `SurveyQuestionsAPI.GetSystemSurveysByParentIdQuestionsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemSurveysByParentIdQuestionsCountRequest struct via the builder pattern


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


## PatchSystemSurveysByParentIdQuestionsById

> SurveyQuestion PatchSystemSurveysByParentIdQuestionsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch SurveyQuestion

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
	id := int32(56) // int32 | questionId
	parentId := int32(56) // int32 | surveyId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveyQuestionsAPI.PatchSystemSurveysByParentIdQuestionsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyQuestionsAPI.PatchSystemSurveysByParentIdQuestionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemSurveysByParentIdQuestionsById`: SurveyQuestion
	fmt.Fprintf(os.Stdout, "Response from `SurveyQuestionsAPI.PatchSystemSurveysByParentIdQuestionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | questionId | 
**parentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemSurveysByParentIdQuestionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**SurveyQuestion**](SurveyQuestion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemSurveysByParentIdQuestions

> SurveyQuestion PostSystemSurveysByParentIdQuestions(ctx, parentId).ClientId(clientId).SurveyQuestion(surveyQuestion).Execute()

Post SurveyQuestion

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
	surveyQuestion := *openapiclient.NewSurveyQuestion("FieldType_example", "EntryType_example", NullableFloat64(123), "Question_example") // SurveyQuestion | surveyQuestion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveyQuestionsAPI.PostSystemSurveysByParentIdQuestions(context.Background(), parentId).ClientId(clientId).SurveyQuestion(surveyQuestion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyQuestionsAPI.PostSystemSurveysByParentIdQuestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemSurveysByParentIdQuestions`: SurveyQuestion
	fmt.Fprintf(os.Stdout, "Response from `SurveyQuestionsAPI.PostSystemSurveysByParentIdQuestions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemSurveysByParentIdQuestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **surveyQuestion** | [**SurveyQuestion**](SurveyQuestion.md) | surveyQuestion | 

### Return type

[**SurveyQuestion**](SurveyQuestion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemSurveysByParentIdQuestionsById

> SurveyQuestion PutSystemSurveysByParentIdQuestionsById(ctx, id, parentId).ClientId(clientId).SurveyQuestion(surveyQuestion).Execute()

Put SurveyQuestion

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
	id := int32(56) // int32 | questionId
	parentId := int32(56) // int32 | surveyId
	clientId := "clientId_example" // string | 
	surveyQuestion := *openapiclient.NewSurveyQuestion("FieldType_example", "EntryType_example", NullableFloat64(123), "Question_example") // SurveyQuestion | surveyQuestion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveyQuestionsAPI.PutSystemSurveysByParentIdQuestionsById(context.Background(), id, parentId).ClientId(clientId).SurveyQuestion(surveyQuestion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyQuestionsAPI.PutSystemSurveysByParentIdQuestionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemSurveysByParentIdQuestionsById`: SurveyQuestion
	fmt.Fprintf(os.Stdout, "Response from `SurveyQuestionsAPI.PutSystemSurveysByParentIdQuestionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | questionId | 
**parentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemSurveysByParentIdQuestionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **surveyQuestion** | [**SurveyQuestion**](SurveyQuestion.md) | surveyQuestion | 

### Return type

[**SurveyQuestion**](SurveyQuestion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

