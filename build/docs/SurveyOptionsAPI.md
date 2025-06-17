# \SurveyOptionsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById**](SurveyOptionsAPI.md#DeleteServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById) | **Delete** /service/surveys/{grandparentId}/questions/{parentId}/options/{id} | Delete SurveyOption
[**GetServiceSurveysByGrandparentIdQuestionsByParentIdOptions**](SurveyOptionsAPI.md#GetServiceSurveysByGrandparentIdQuestionsByParentIdOptions) | **Get** /service/surveys/{grandparentId}/questions/{parentId}/options | Get List of SurveyOption
[**GetServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById**](SurveyOptionsAPI.md#GetServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById) | **Get** /service/surveys/{grandparentId}/questions/{parentId}/options/{id} | Get SurveyOption
[**GetServiceSurveysByGrandparentIdQuestionsByParentIdOptionsCount**](SurveyOptionsAPI.md#GetServiceSurveysByGrandparentIdQuestionsByParentIdOptionsCount) | **Get** /service/surveys/{grandparentId}/questions/{parentId}/options/count | Get Count of SurveyOption
[**PatchServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById**](SurveyOptionsAPI.md#PatchServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById) | **Patch** /service/surveys/{grandparentId}/questions/{parentId}/options/{id} | Patch SurveyOption
[**PostServiceSurveysByGrandparentIdQuestionsByParentIdOptions**](SurveyOptionsAPI.md#PostServiceSurveysByGrandparentIdQuestionsByParentIdOptions) | **Post** /service/surveys/{grandparentId}/questions/{parentId}/options | Post SurveyOption
[**PutServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById**](SurveyOptionsAPI.md#PutServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById) | **Put** /service/surveys/{grandparentId}/questions/{parentId}/options/{id} | Put SurveyOption



## DeleteServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById

> DeleteServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById(ctx, id, parentId, grandparentId).ClientId(clientId).Execute()

Delete SurveyOption

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
	id := int32(56) // int32 | optionId
	parentId := int32(56) // int32 | questionId
	grandparentId := int32(56) // int32 | surveyId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SurveyOptionsAPI.DeleteServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyOptionsAPI.DeleteServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | optionId | 
**parentId** | **int32** | questionId | 
**grandparentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceSurveysByGrandparentIdQuestionsByParentIdOptionsByIdRequest struct via the builder pattern


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


## GetServiceSurveysByGrandparentIdQuestionsByParentIdOptions

> []SurveyOption GetServiceSurveysByGrandparentIdQuestionsByParentIdOptions(ctx, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of SurveyOption

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
	parentId := int32(56) // int32 | questionId
	grandparentId := int32(56) // int32 | surveyId
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
	resp, r, err := apiClient.SurveyOptionsAPI.GetServiceSurveysByGrandparentIdQuestionsByParentIdOptions(context.Background(), parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyOptionsAPI.GetServiceSurveysByGrandparentIdQuestionsByParentIdOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceSurveysByGrandparentIdQuestionsByParentIdOptions`: []SurveyOption
	fmt.Fprintf(os.Stdout, "Response from `SurveyOptionsAPI.GetServiceSurveysByGrandparentIdQuestionsByParentIdOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | questionId | 
**grandparentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceSurveysByGrandparentIdQuestionsByParentIdOptionsRequest struct via the builder pattern


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

[**[]SurveyOption**](SurveyOption.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById

> SurveyOption GetServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById(ctx, id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get SurveyOption

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
	id := int32(56) // int32 | optionId
	parentId := int32(56) // int32 | questionId
	grandparentId := int32(56) // int32 | surveyId
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
	resp, r, err := apiClient.SurveyOptionsAPI.GetServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyOptionsAPI.GetServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById`: SurveyOption
	fmt.Fprintf(os.Stdout, "Response from `SurveyOptionsAPI.GetServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | optionId | 
**parentId** | **int32** | questionId | 
**grandparentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceSurveysByGrandparentIdQuestionsByParentIdOptionsByIdRequest struct via the builder pattern


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

[**SurveyOption**](SurveyOption.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceSurveysByGrandparentIdQuestionsByParentIdOptionsCount

> Count GetServiceSurveysByGrandparentIdQuestionsByParentIdOptionsCount(ctx, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of SurveyOption

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
	parentId := int32(56) // int32 | questionId
	grandparentId := int32(56) // int32 | surveyId
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
	resp, r, err := apiClient.SurveyOptionsAPI.GetServiceSurveysByGrandparentIdQuestionsByParentIdOptionsCount(context.Background(), parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyOptionsAPI.GetServiceSurveysByGrandparentIdQuestionsByParentIdOptionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceSurveysByGrandparentIdQuestionsByParentIdOptionsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `SurveyOptionsAPI.GetServiceSurveysByGrandparentIdQuestionsByParentIdOptionsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | questionId | 
**grandparentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceSurveysByGrandparentIdQuestionsByParentIdOptionsCountRequest struct via the builder pattern


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


## PatchServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById

> SurveyOption PatchServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById(ctx, id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch SurveyOption

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
	id := int32(56) // int32 | optionId
	parentId := int32(56) // int32 | questionId
	grandparentId := int32(56) // int32 | surveyId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveyOptionsAPI.PatchServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyOptionsAPI.PatchServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById`: SurveyOption
	fmt.Fprintf(os.Stdout, "Response from `SurveyOptionsAPI.PatchServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | optionId | 
**parentId** | **int32** | questionId | 
**grandparentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchServiceSurveysByGrandparentIdQuestionsByParentIdOptionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**SurveyOption**](SurveyOption.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServiceSurveysByGrandparentIdQuestionsByParentIdOptions

> SurveyOption PostServiceSurveysByGrandparentIdQuestionsByParentIdOptions(ctx, parentId, grandparentId).ClientId(clientId).SurveyOption(surveyOption).Execute()

Post SurveyOption

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
	parentId := int32(56) // int32 | questionId
	grandparentId := int32(56) // int32 | surveyId
	clientId := "clientId_example" // string | 
	surveyOption := *openapiclient.NewSurveyOption("Caption_example", NullableInt32(123)) // SurveyOption | surveyOption

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveyOptionsAPI.PostServiceSurveysByGrandparentIdQuestionsByParentIdOptions(context.Background(), parentId, grandparentId).ClientId(clientId).SurveyOption(surveyOption).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyOptionsAPI.PostServiceSurveysByGrandparentIdQuestionsByParentIdOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServiceSurveysByGrandparentIdQuestionsByParentIdOptions`: SurveyOption
	fmt.Fprintf(os.Stdout, "Response from `SurveyOptionsAPI.PostServiceSurveysByGrandparentIdQuestionsByParentIdOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | questionId | 
**grandparentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServiceSurveysByGrandparentIdQuestionsByParentIdOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **surveyOption** | [**SurveyOption**](SurveyOption.md) | surveyOption | 

### Return type

[**SurveyOption**](SurveyOption.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById

> SurveyOption PutServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById(ctx, id, parentId, grandparentId).ClientId(clientId).SurveyOption(surveyOption).Execute()

Put SurveyOption

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
	id := int32(56) // int32 | optionId
	parentId := int32(56) // int32 | questionId
	grandparentId := int32(56) // int32 | surveyId
	clientId := "clientId_example" // string | 
	surveyOption := *openapiclient.NewSurveyOption("Caption_example", NullableInt32(123)) // SurveyOption | surveyOption

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveyOptionsAPI.PutServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).SurveyOption(surveyOption).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyOptionsAPI.PutServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById`: SurveyOption
	fmt.Fprintf(os.Stdout, "Response from `SurveyOptionsAPI.PutServiceSurveysByGrandparentIdQuestionsByParentIdOptionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | optionId | 
**parentId** | **int32** | questionId | 
**grandparentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServiceSurveysByGrandparentIdQuestionsByParentIdOptionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **surveyOption** | [**SurveyOption**](SurveyOption.md) | surveyOption | 

### Return type

[**SurveyOption**](SurveyOption.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

