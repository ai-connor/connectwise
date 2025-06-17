# \ServiceSurveyQuestionsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServiceSurveysByParentIdQuestionsById**](ServiceSurveyQuestionsAPI.md#DeleteServiceSurveysByParentIdQuestionsById) | **Delete** /service/surveys/{parentId}/questions/{id} | Delete ServiceSurveyQuestion
[**GetServiceSurveysByParentIdQuestions**](ServiceSurveyQuestionsAPI.md#GetServiceSurveysByParentIdQuestions) | **Get** /service/surveys/{parentId}/questions | Get List of ServiceSurveyQuestion
[**GetServiceSurveysByParentIdQuestionsById**](ServiceSurveyQuestionsAPI.md#GetServiceSurveysByParentIdQuestionsById) | **Get** /service/surveys/{parentId}/questions/{id} | Get ServiceSurveyQuestion
[**GetServiceSurveysByParentIdQuestionsCount**](ServiceSurveyQuestionsAPI.md#GetServiceSurveysByParentIdQuestionsCount) | **Get** /service/surveys/{parentId}/questions/count | Get Count of ServiceSurveyQuestion
[**PatchServiceSurveysByParentIdQuestionsById**](ServiceSurveyQuestionsAPI.md#PatchServiceSurveysByParentIdQuestionsById) | **Patch** /service/surveys/{parentId}/questions/{id} | Patch ServiceSurveyQuestion
[**PostServiceSurveysByParentIdQuestions**](ServiceSurveyQuestionsAPI.md#PostServiceSurveysByParentIdQuestions) | **Post** /service/surveys/{parentId}/questions | Post ServiceSurveyQuestion
[**PostServiceSurveysByParentIdQuestionsByIdCopy**](ServiceSurveyQuestionsAPI.md#PostServiceSurveysByParentIdQuestionsByIdCopy) | **Post** /service/surveys/{parentId}/questions/{id}/copy | Post ServiceSurveyQuestion
[**PutServiceSurveysByParentIdQuestionsById**](ServiceSurveyQuestionsAPI.md#PutServiceSurveysByParentIdQuestionsById) | **Put** /service/surveys/{parentId}/questions/{id} | Put ServiceSurveyQuestion



## DeleteServiceSurveysByParentIdQuestionsById

> DeleteServiceSurveysByParentIdQuestionsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ServiceSurveyQuestion

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
	r, err := apiClient.ServiceSurveyQuestionsAPI.DeleteServiceSurveysByParentIdQuestionsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSurveyQuestionsAPI.DeleteServiceSurveysByParentIdQuestionsById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteServiceSurveysByParentIdQuestionsByIdRequest struct via the builder pattern


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


## GetServiceSurveysByParentIdQuestions

> []ServiceSurveyQuestion GetServiceSurveysByParentIdQuestions(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ServiceSurveyQuestion

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
	resp, r, err := apiClient.ServiceSurveyQuestionsAPI.GetServiceSurveysByParentIdQuestions(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSurveyQuestionsAPI.GetServiceSurveysByParentIdQuestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceSurveysByParentIdQuestions`: []ServiceSurveyQuestion
	fmt.Fprintf(os.Stdout, "Response from `ServiceSurveyQuestionsAPI.GetServiceSurveysByParentIdQuestions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceSurveysByParentIdQuestionsRequest struct via the builder pattern


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

[**[]ServiceSurveyQuestion**](ServiceSurveyQuestion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceSurveysByParentIdQuestionsById

> ServiceSurveyQuestion GetServiceSurveysByParentIdQuestionsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ServiceSurveyQuestion

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
	resp, r, err := apiClient.ServiceSurveyQuestionsAPI.GetServiceSurveysByParentIdQuestionsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSurveyQuestionsAPI.GetServiceSurveysByParentIdQuestionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceSurveysByParentIdQuestionsById`: ServiceSurveyQuestion
	fmt.Fprintf(os.Stdout, "Response from `ServiceSurveyQuestionsAPI.GetServiceSurveysByParentIdQuestionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | questionId | 
**parentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceSurveysByParentIdQuestionsByIdRequest struct via the builder pattern


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

[**ServiceSurveyQuestion**](ServiceSurveyQuestion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceSurveysByParentIdQuestionsCount

> Count GetServiceSurveysByParentIdQuestionsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ServiceSurveyQuestion

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
	resp, r, err := apiClient.ServiceSurveyQuestionsAPI.GetServiceSurveysByParentIdQuestionsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSurveyQuestionsAPI.GetServiceSurveysByParentIdQuestionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceSurveysByParentIdQuestionsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ServiceSurveyQuestionsAPI.GetServiceSurveysByParentIdQuestionsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceSurveysByParentIdQuestionsCountRequest struct via the builder pattern


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


## PatchServiceSurveysByParentIdQuestionsById

> ServiceSurveyQuestion PatchServiceSurveysByParentIdQuestionsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ServiceSurveyQuestion

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
	resp, r, err := apiClient.ServiceSurveyQuestionsAPI.PatchServiceSurveysByParentIdQuestionsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSurveyQuestionsAPI.PatchServiceSurveysByParentIdQuestionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchServiceSurveysByParentIdQuestionsById`: ServiceSurveyQuestion
	fmt.Fprintf(os.Stdout, "Response from `ServiceSurveyQuestionsAPI.PatchServiceSurveysByParentIdQuestionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | questionId | 
**parentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchServiceSurveysByParentIdQuestionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ServiceSurveyQuestion**](ServiceSurveyQuestion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServiceSurveysByParentIdQuestions

> ServiceSurveyQuestion PostServiceSurveysByParentIdQuestions(ctx, parentId).ClientId(clientId).ServiceSurveyQuestion(serviceSurveyQuestion).Execute()

Post ServiceSurveyQuestion

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
	serviceSurveyQuestion := *openapiclient.NewServiceSurveyQuestion("Type_example", "Question_example") // ServiceSurveyQuestion | serviceSurveyQuestion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSurveyQuestionsAPI.PostServiceSurveysByParentIdQuestions(context.Background(), parentId).ClientId(clientId).ServiceSurveyQuestion(serviceSurveyQuestion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSurveyQuestionsAPI.PostServiceSurveysByParentIdQuestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServiceSurveysByParentIdQuestions`: ServiceSurveyQuestion
	fmt.Fprintf(os.Stdout, "Response from `ServiceSurveyQuestionsAPI.PostServiceSurveysByParentIdQuestions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServiceSurveysByParentIdQuestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **serviceSurveyQuestion** | [**ServiceSurveyQuestion**](ServiceSurveyQuestion.md) | serviceSurveyQuestion | 

### Return type

[**ServiceSurveyQuestion**](ServiceSurveyQuestion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServiceSurveysByParentIdQuestionsByIdCopy

> ServiceSurveyQuestion PostServiceSurveysByParentIdQuestionsByIdCopy(ctx, id, parentId).ClientId(clientId).Execute()

Post ServiceSurveyQuestion

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
	resp, r, err := apiClient.ServiceSurveyQuestionsAPI.PostServiceSurveysByParentIdQuestionsByIdCopy(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSurveyQuestionsAPI.PostServiceSurveysByParentIdQuestionsByIdCopy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServiceSurveysByParentIdQuestionsByIdCopy`: ServiceSurveyQuestion
	fmt.Fprintf(os.Stdout, "Response from `ServiceSurveyQuestionsAPI.PostServiceSurveysByParentIdQuestionsByIdCopy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | questionId | 
**parentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServiceSurveysByParentIdQuestionsByIdCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 

### Return type

[**ServiceSurveyQuestion**](ServiceSurveyQuestion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServiceSurveysByParentIdQuestionsById

> ServiceSurveyQuestion PutServiceSurveysByParentIdQuestionsById(ctx, id, parentId).ClientId(clientId).ServiceSurveyQuestion(serviceSurveyQuestion).Execute()

Put ServiceSurveyQuestion

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
	serviceSurveyQuestion := *openapiclient.NewServiceSurveyQuestion("Type_example", "Question_example") // ServiceSurveyQuestion | serviceSurveyQuestion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSurveyQuestionsAPI.PutServiceSurveysByParentIdQuestionsById(context.Background(), id, parentId).ClientId(clientId).ServiceSurveyQuestion(serviceSurveyQuestion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSurveyQuestionsAPI.PutServiceSurveysByParentIdQuestionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServiceSurveysByParentIdQuestionsById`: ServiceSurveyQuestion
	fmt.Fprintf(os.Stdout, "Response from `ServiceSurveyQuestionsAPI.PutServiceSurveysByParentIdQuestionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | questionId | 
**parentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServiceSurveysByParentIdQuestionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **serviceSurveyQuestion** | [**ServiceSurveyQuestion**](ServiceSurveyQuestion.md) | serviceSurveyQuestion | 

### Return type

[**ServiceSurveyQuestion**](ServiceSurveyQuestion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

