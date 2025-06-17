# \ConfigurationTypeQuestionValueInfosAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfo**](ConfigurationTypeQuestionValueInfosAPI.md#GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfo) | **Get** /configurations/types/{grandparentId}/questions/{parentId}/values/{id}/info | Get ConfigurationTypeQuestionValueInfo
[**GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfo**](ConfigurationTypeQuestionValueInfosAPI.md#GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfo) | **Get** /configurations/types/{grandparentId}/questions/{parentId}/values/info | Get ConfigurationTypeQuestionValueInfo
[**GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCount**](ConfigurationTypeQuestionValueInfosAPI.md#GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCount) | **Get** /configurations/types/{grandparentId}/questions/{parentId}/values/info/count | Get Count of ConfigurationTypeQuestionValueInfos



## GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfo

> ConfigurationTypeQuestionInfo GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfo(ctx, grandparentId, parentId, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ConfigurationTypeQuestionValueInfo

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
	grandparentId := int32(56) // int32 | ConfigurationTypeQuestionInfo
	parentId := int32(56) // int32 | ConfigurationTypeQuestionInfo
	id := int32(56) // int32 | ConfigurationTypeQuestionInfo
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
	resp, r, err := apiClient.ConfigurationTypeQuestionValueInfosAPI.GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfo(context.Background(), grandparentId, parentId, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTypeQuestionValueInfosAPI.GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfo`: ConfigurationTypeQuestionInfo
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationTypeQuestionValueInfosAPI.GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**grandparentId** | **int32** | ConfigurationTypeQuestionInfo | 
**parentId** | **int32** | ConfigurationTypeQuestionInfo | 
**id** | **int32** | ConfigurationTypeQuestionInfo | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdInfoRequest struct via the builder pattern


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

[**ConfigurationTypeQuestionInfo**](ConfigurationTypeQuestionInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfo

> []ConfigurationTypeQuestionValueInfo GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfo(ctx, grandparentId, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ConfigurationTypeQuestionValueInfo

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
	grandparentId := int32(56) // int32 | ConfigurationTypeQuestionValueInfo
	parentId := int32(56) // int32 | ConfigurationTypeQuestionValueInfo
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
	resp, r, err := apiClient.ConfigurationTypeQuestionValueInfosAPI.GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfo(context.Background(), grandparentId, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTypeQuestionValueInfosAPI.GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfo`: []ConfigurationTypeQuestionValueInfo
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationTypeQuestionValueInfosAPI.GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**grandparentId** | **int32** | ConfigurationTypeQuestionValueInfo | 
**parentId** | **int32** | ConfigurationTypeQuestionValueInfo | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoRequest struct via the builder pattern


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

[**[]ConfigurationTypeQuestionValueInfo**](ConfigurationTypeQuestionValueInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCount

> Count GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCount(ctx, grandparentId, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ConfigurationTypeQuestionValueInfos

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
	grandparentId := int32(56) // int32 | ConfigurationTypeQuestionValueInfo
	parentId := int32(56) // int32 | ConfigurationTypeQuestionValueInfo
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
	resp, r, err := apiClient.ConfigurationTypeQuestionValueInfosAPI.GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCount(context.Background(), grandparentId, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTypeQuestionValueInfosAPI.GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationTypeQuestionValueInfosAPI.GetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**grandparentId** | **int32** | ConfigurationTypeQuestionValueInfo | 
**parentId** | **int32** | ConfigurationTypeQuestionValueInfo | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesInfoCountRequest struct via the builder pattern


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

