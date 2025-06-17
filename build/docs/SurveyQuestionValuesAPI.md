# \SurveyQuestionValuesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemSurveysByGrandparentIdQuestionsByParentIdValuesById**](SurveyQuestionValuesAPI.md#DeleteSystemSurveysByGrandparentIdQuestionsByParentIdValuesById) | **Delete** /system/surveys/{grandparentId}/questions/{parentId}/values/{id} | Delete SurveyQuestionValue
[**GetSystemSurveysByGrandparentIdQuestionsByParentIdValues**](SurveyQuestionValuesAPI.md#GetSystemSurveysByGrandparentIdQuestionsByParentIdValues) | **Get** /system/surveys/{grandparentId}/questions/{parentId}/values | Get List of SurveyQuestionValue
[**GetSystemSurveysByGrandparentIdQuestionsByParentIdValuesById**](SurveyQuestionValuesAPI.md#GetSystemSurveysByGrandparentIdQuestionsByParentIdValuesById) | **Get** /system/surveys/{grandparentId}/questions/{parentId}/values/{id} | Get SurveyQuestionValue
[**PatchSystemSurveysByGrandparentIdQuestionsByParentIdValuesById**](SurveyQuestionValuesAPI.md#PatchSystemSurveysByGrandparentIdQuestionsByParentIdValuesById) | **Patch** /system/surveys/{grandparentId}/questions/{parentId}/values/{id} | Patch SurveyQuestionValue
[**PostSystemSurveysByGrandparentIdQuestionsByParentIdValues**](SurveyQuestionValuesAPI.md#PostSystemSurveysByGrandparentIdQuestionsByParentIdValues) | **Post** /system/surveys/{grandparentId}/questions/{parentId}/values | Post SurveyQuestionValue
[**PutSystemSurveysByGrandparentIdQuestionsByParentIdValuesById**](SurveyQuestionValuesAPI.md#PutSystemSurveysByGrandparentIdQuestionsByParentIdValuesById) | **Put** /system/surveys/{grandparentId}/questions/{parentId}/values/{id} | Put SurveyQuestionValue



## DeleteSystemSurveysByGrandparentIdQuestionsByParentIdValuesById

> DeleteSystemSurveysByGrandparentIdQuestionsByParentIdValuesById(ctx, id, parentId, grandparentId).ClientId(clientId).Execute()

Delete SurveyQuestionValue

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
	id := int32(56) // int32 | valueId
	parentId := int32(56) // int32 | questionId
	grandparentId := int32(56) // int32 | surveyId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SurveyQuestionValuesAPI.DeleteSystemSurveysByGrandparentIdQuestionsByParentIdValuesById(context.Background(), id, parentId, grandparentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyQuestionValuesAPI.DeleteSystemSurveysByGrandparentIdQuestionsByParentIdValuesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | valueId | 
**parentId** | **int32** | questionId | 
**grandparentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemSurveysByGrandparentIdQuestionsByParentIdValuesByIdRequest struct via the builder pattern


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


## GetSystemSurveysByGrandparentIdQuestionsByParentIdValues

> []SurveyQuestionValue GetSystemSurveysByGrandparentIdQuestionsByParentIdValues(ctx, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of SurveyQuestionValue

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
	resp, r, err := apiClient.SurveyQuestionValuesAPI.GetSystemSurveysByGrandparentIdQuestionsByParentIdValues(context.Background(), parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyQuestionValuesAPI.GetSystemSurveysByGrandparentIdQuestionsByParentIdValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemSurveysByGrandparentIdQuestionsByParentIdValues`: []SurveyQuestionValue
	fmt.Fprintf(os.Stdout, "Response from `SurveyQuestionValuesAPI.GetSystemSurveysByGrandparentIdQuestionsByParentIdValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | questionId | 
**grandparentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemSurveysByGrandparentIdQuestionsByParentIdValuesRequest struct via the builder pattern


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

[**[]SurveyQuestionValue**](SurveyQuestionValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemSurveysByGrandparentIdQuestionsByParentIdValuesById

> SurveyQuestionValue GetSystemSurveysByGrandparentIdQuestionsByParentIdValuesById(ctx, id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get SurveyQuestionValue

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
	id := int32(56) // int32 | valueId
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
	resp, r, err := apiClient.SurveyQuestionValuesAPI.GetSystemSurveysByGrandparentIdQuestionsByParentIdValuesById(context.Background(), id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyQuestionValuesAPI.GetSystemSurveysByGrandparentIdQuestionsByParentIdValuesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemSurveysByGrandparentIdQuestionsByParentIdValuesById`: SurveyQuestionValue
	fmt.Fprintf(os.Stdout, "Response from `SurveyQuestionValuesAPI.GetSystemSurveysByGrandparentIdQuestionsByParentIdValuesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | valueId | 
**parentId** | **int32** | questionId | 
**grandparentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemSurveysByGrandparentIdQuestionsByParentIdValuesByIdRequest struct via the builder pattern


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

[**SurveyQuestionValue**](SurveyQuestionValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSystemSurveysByGrandparentIdQuestionsByParentIdValuesById

> SurveyQuestionValue PatchSystemSurveysByGrandparentIdQuestionsByParentIdValuesById(ctx, id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch SurveyQuestionValue

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
	id := int32(56) // int32 | valueId
	parentId := int32(56) // int32 | questionId
	grandparentId := int32(56) // int32 | surveyId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveyQuestionValuesAPI.PatchSystemSurveysByGrandparentIdQuestionsByParentIdValuesById(context.Background(), id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyQuestionValuesAPI.PatchSystemSurveysByGrandparentIdQuestionsByParentIdValuesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemSurveysByGrandparentIdQuestionsByParentIdValuesById`: SurveyQuestionValue
	fmt.Fprintf(os.Stdout, "Response from `SurveyQuestionValuesAPI.PatchSystemSurveysByGrandparentIdQuestionsByParentIdValuesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | valueId | 
**parentId** | **int32** | questionId | 
**grandparentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemSurveysByGrandparentIdQuestionsByParentIdValuesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**SurveyQuestionValue**](SurveyQuestionValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemSurveysByGrandparentIdQuestionsByParentIdValues

> SurveyQuestionValue PostSystemSurveysByGrandparentIdQuestionsByParentIdValues(ctx, parentId, grandparentId).ClientId(clientId).SurveyQuestionValue(surveyQuestionValue).Execute()

Post SurveyQuestionValue

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
	surveyQuestionValue := *openapiclient.NewSurveyQuestionValue("Value_example") // SurveyQuestionValue | surveyQuestionValue

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveyQuestionValuesAPI.PostSystemSurveysByGrandparentIdQuestionsByParentIdValues(context.Background(), parentId, grandparentId).ClientId(clientId).SurveyQuestionValue(surveyQuestionValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyQuestionValuesAPI.PostSystemSurveysByGrandparentIdQuestionsByParentIdValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemSurveysByGrandparentIdQuestionsByParentIdValues`: SurveyQuestionValue
	fmt.Fprintf(os.Stdout, "Response from `SurveyQuestionValuesAPI.PostSystemSurveysByGrandparentIdQuestionsByParentIdValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | questionId | 
**grandparentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemSurveysByGrandparentIdQuestionsByParentIdValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **surveyQuestionValue** | [**SurveyQuestionValue**](SurveyQuestionValue.md) | surveyQuestionValue | 

### Return type

[**SurveyQuestionValue**](SurveyQuestionValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemSurveysByGrandparentIdQuestionsByParentIdValuesById

> SurveyQuestionValue PutSystemSurveysByGrandparentIdQuestionsByParentIdValuesById(ctx, id, parentId, grandparentId).ClientId(clientId).SurveyQuestionValue(surveyQuestionValue).Execute()

Put SurveyQuestionValue

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
	id := int32(56) // int32 | valueId
	parentId := int32(56) // int32 | questionId
	grandparentId := int32(56) // int32 | surveyId
	clientId := "clientId_example" // string | 
	surveyQuestionValue := *openapiclient.NewSurveyQuestionValue("Value_example") // SurveyQuestionValue | surveyQuestionValue

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveyQuestionValuesAPI.PutSystemSurveysByGrandparentIdQuestionsByParentIdValuesById(context.Background(), id, parentId, grandparentId).ClientId(clientId).SurveyQuestionValue(surveyQuestionValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveyQuestionValuesAPI.PutSystemSurveysByGrandparentIdQuestionsByParentIdValuesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemSurveysByGrandparentIdQuestionsByParentIdValuesById`: SurveyQuestionValue
	fmt.Fprintf(os.Stdout, "Response from `SurveyQuestionValuesAPI.PutSystemSurveysByGrandparentIdQuestionsByParentIdValuesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | valueId | 
**parentId** | **int32** | questionId | 
**grandparentId** | **int32** | surveyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemSurveysByGrandparentIdQuestionsByParentIdValuesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **surveyQuestionValue** | [**SurveyQuestionValue**](SurveyQuestionValue.md) | surveyQuestionValue | 

### Return type

[**SurveyQuestionValue**](SurveyQuestionValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

