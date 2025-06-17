# \ConfigurationTypeQuestionsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyConfigurationsTypesByParentIdQuestionsById**](ConfigurationTypeQuestionsAPI.md#DeleteCompanyConfigurationsTypesByParentIdQuestionsById) | **Delete** /company/configurations/types/{parentId}/questions/{id} | Delete ConfigurationTypeQuestion
[**GetCompanyConfigurationsTypesByParentIdQuestions**](ConfigurationTypeQuestionsAPI.md#GetCompanyConfigurationsTypesByParentIdQuestions) | **Get** /company/configurations/types/{parentId}/questions | Get List of ConfigurationTypeQuestion
[**GetCompanyConfigurationsTypesByParentIdQuestionsById**](ConfigurationTypeQuestionsAPI.md#GetCompanyConfigurationsTypesByParentIdQuestionsById) | **Get** /company/configurations/types/{parentId}/questions/{id} | Get ConfigurationTypeQuestion
[**GetCompanyConfigurationsTypesByParentIdQuestionsCount**](ConfigurationTypeQuestionsAPI.md#GetCompanyConfigurationsTypesByParentIdQuestionsCount) | **Get** /company/configurations/types/{parentId}/questions/count | Get Count of ConfigurationTypeQuestion
[**PatchCompanyConfigurationsTypesByParentIdQuestionsById**](ConfigurationTypeQuestionsAPI.md#PatchCompanyConfigurationsTypesByParentIdQuestionsById) | **Patch** /company/configurations/types/{parentId}/questions/{id} | Patch ConfigurationTypeQuestion
[**PostCompanyConfigurationsTypesByParentIdQuestions**](ConfigurationTypeQuestionsAPI.md#PostCompanyConfigurationsTypesByParentIdQuestions) | **Post** /company/configurations/types/{parentId}/questions | Post ConfigurationTypeQuestion
[**PutCompanyConfigurationsTypesByParentIdQuestionsById**](ConfigurationTypeQuestionsAPI.md#PutCompanyConfigurationsTypesByParentIdQuestionsById) | **Put** /company/configurations/types/{parentId}/questions/{id} | Put ConfigurationTypeQuestion



## DeleteCompanyConfigurationsTypesByParentIdQuestionsById

> DeleteCompanyConfigurationsTypesByParentIdQuestionsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ConfigurationTypeQuestion

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
	parentId := int32(56) // int32 | typeId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationTypeQuestionsAPI.DeleteCompanyConfigurationsTypesByParentIdQuestionsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTypeQuestionsAPI.DeleteCompanyConfigurationsTypesByParentIdQuestionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | questionId | 
**parentId** | **int32** | typeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyConfigurationsTypesByParentIdQuestionsByIdRequest struct via the builder pattern


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


## GetCompanyConfigurationsTypesByParentIdQuestions

> []ConfigurationTypeQuestion GetCompanyConfigurationsTypesByParentIdQuestions(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ConfigurationTypeQuestion

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
	parentId := int32(56) // int32 | typeId
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
	resp, r, err := apiClient.ConfigurationTypeQuestionsAPI.GetCompanyConfigurationsTypesByParentIdQuestions(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTypeQuestionsAPI.GetCompanyConfigurationsTypesByParentIdQuestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyConfigurationsTypesByParentIdQuestions`: []ConfigurationTypeQuestion
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationTypeQuestionsAPI.GetCompanyConfigurationsTypesByParentIdQuestions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | typeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyConfigurationsTypesByParentIdQuestionsRequest struct via the builder pattern


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

[**[]ConfigurationTypeQuestion**](ConfigurationTypeQuestion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyConfigurationsTypesByParentIdQuestionsById

> ConfigurationTypeQuestion GetCompanyConfigurationsTypesByParentIdQuestionsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ConfigurationTypeQuestion

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
	parentId := int32(56) // int32 | typeId
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
	resp, r, err := apiClient.ConfigurationTypeQuestionsAPI.GetCompanyConfigurationsTypesByParentIdQuestionsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTypeQuestionsAPI.GetCompanyConfigurationsTypesByParentIdQuestionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyConfigurationsTypesByParentIdQuestionsById`: ConfigurationTypeQuestion
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationTypeQuestionsAPI.GetCompanyConfigurationsTypesByParentIdQuestionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | questionId | 
**parentId** | **int32** | typeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyConfigurationsTypesByParentIdQuestionsByIdRequest struct via the builder pattern


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

[**ConfigurationTypeQuestion**](ConfigurationTypeQuestion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyConfigurationsTypesByParentIdQuestionsCount

> Count GetCompanyConfigurationsTypesByParentIdQuestionsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ConfigurationTypeQuestion

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
	parentId := int32(56) // int32 | typeId
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
	resp, r, err := apiClient.ConfigurationTypeQuestionsAPI.GetCompanyConfigurationsTypesByParentIdQuestionsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTypeQuestionsAPI.GetCompanyConfigurationsTypesByParentIdQuestionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyConfigurationsTypesByParentIdQuestionsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationTypeQuestionsAPI.GetCompanyConfigurationsTypesByParentIdQuestionsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | typeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyConfigurationsTypesByParentIdQuestionsCountRequest struct via the builder pattern


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


## PatchCompanyConfigurationsTypesByParentIdQuestionsById

> ConfigurationTypeQuestion PatchCompanyConfigurationsTypesByParentIdQuestionsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ConfigurationTypeQuestion

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
	parentId := int32(56) // int32 | typeId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationTypeQuestionsAPI.PatchCompanyConfigurationsTypesByParentIdQuestionsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTypeQuestionsAPI.PatchCompanyConfigurationsTypesByParentIdQuestionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyConfigurationsTypesByParentIdQuestionsById`: ConfigurationTypeQuestion
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationTypeQuestionsAPI.PatchCompanyConfigurationsTypesByParentIdQuestionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | questionId | 
**parentId** | **int32** | typeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyConfigurationsTypesByParentIdQuestionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ConfigurationTypeQuestion**](ConfigurationTypeQuestion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyConfigurationsTypesByParentIdQuestions

> ConfigurationTypeQuestion PostCompanyConfigurationsTypesByParentIdQuestions(ctx, parentId).ClientId(clientId).ConfigurationTypeQuestion(configurationTypeQuestion).Execute()

Post ConfigurationTypeQuestion

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
	parentId := int32(56) // int32 | typeId
	clientId := "clientId_example" // string | 
	configurationTypeQuestion := *openapiclient.NewConfigurationTypeQuestion("FieldType_example", "EntryType_example", NullableFloat64(123), "Question_example") // ConfigurationTypeQuestion | configurationTypeQuestion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationTypeQuestionsAPI.PostCompanyConfigurationsTypesByParentIdQuestions(context.Background(), parentId).ClientId(clientId).ConfigurationTypeQuestion(configurationTypeQuestion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTypeQuestionsAPI.PostCompanyConfigurationsTypesByParentIdQuestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyConfigurationsTypesByParentIdQuestions`: ConfigurationTypeQuestion
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationTypeQuestionsAPI.PostCompanyConfigurationsTypesByParentIdQuestions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | typeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyConfigurationsTypesByParentIdQuestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **configurationTypeQuestion** | [**ConfigurationTypeQuestion**](ConfigurationTypeQuestion.md) | configurationTypeQuestion | 

### Return type

[**ConfigurationTypeQuestion**](ConfigurationTypeQuestion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyConfigurationsTypesByParentIdQuestionsById

> ConfigurationTypeQuestion PutCompanyConfigurationsTypesByParentIdQuestionsById(ctx, id, parentId).ClientId(clientId).ConfigurationTypeQuestion(configurationTypeQuestion).Execute()

Put ConfigurationTypeQuestion

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
	parentId := int32(56) // int32 | typeId
	clientId := "clientId_example" // string | 
	configurationTypeQuestion := *openapiclient.NewConfigurationTypeQuestion("FieldType_example", "EntryType_example", NullableFloat64(123), "Question_example") // ConfigurationTypeQuestion | configurationTypeQuestion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationTypeQuestionsAPI.PutCompanyConfigurationsTypesByParentIdQuestionsById(context.Background(), id, parentId).ClientId(clientId).ConfigurationTypeQuestion(configurationTypeQuestion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTypeQuestionsAPI.PutCompanyConfigurationsTypesByParentIdQuestionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyConfigurationsTypesByParentIdQuestionsById`: ConfigurationTypeQuestion
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationTypeQuestionsAPI.PutCompanyConfigurationsTypesByParentIdQuestionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | questionId | 
**parentId** | **int32** | typeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyConfigurationsTypesByParentIdQuestionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **configurationTypeQuestion** | [**ConfigurationTypeQuestion**](ConfigurationTypeQuestion.md) | configurationTypeQuestion | 

### Return type

[**ConfigurationTypeQuestion**](ConfigurationTypeQuestion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

