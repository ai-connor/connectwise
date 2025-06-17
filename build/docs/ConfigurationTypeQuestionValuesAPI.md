# \ConfigurationTypeQuestionValuesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById**](ConfigurationTypeQuestionValuesAPI.md#DeleteCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById) | **Delete** /company/configurations/types/{grandparentId}/questions/{parentId}/values/{id} | Delete ConfigurationTypeQuestionValue
[**GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValues**](ConfigurationTypeQuestionValuesAPI.md#GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValues) | **Get** /company/configurations/types/{grandparentId}/questions/{parentId}/values | Get List of ConfigurationTypeQuestionValue
[**GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById**](ConfigurationTypeQuestionValuesAPI.md#GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById) | **Get** /company/configurations/types/{grandparentId}/questions/{parentId}/values/{id} | Get ConfigurationTypeQuestionValue
[**GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdUsages**](ConfigurationTypeQuestionValuesAPI.md#GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdUsages) | **Get** /company/configurations/types/{grandparentId}/questions/{parentId}/values/{id}/usages | Get List of Usage
[**GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdUsagesList**](ConfigurationTypeQuestionValuesAPI.md#GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdUsagesList) | **Get** /company/configurations/types/{grandparentId}/questions/{parentId}/values/{id}/usages/list | Get List of Usage
[**GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesCount**](ConfigurationTypeQuestionValuesAPI.md#GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesCount) | **Get** /company/configurations/types/{grandparentId}/questions/{parentId}/values/count | Get Count of ConfigurationTypeQuestionValue
[**PatchCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById**](ConfigurationTypeQuestionValuesAPI.md#PatchCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById) | **Patch** /company/configurations/types/{grandparentId}/questions/{parentId}/values/{id} | Patch ConfigurationTypeQuestionValue
[**PostCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValues**](ConfigurationTypeQuestionValuesAPI.md#PostCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValues) | **Post** /company/configurations/types/{grandparentId}/questions/{parentId}/values | Post ConfigurationTypeQuestionValue
[**PutCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById**](ConfigurationTypeQuestionValuesAPI.md#PutCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById) | **Put** /company/configurations/types/{grandparentId}/questions/{parentId}/values/{id} | Put ConfigurationTypeQuestionValue



## DeleteCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById

> DeleteCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById(ctx, id, parentId, grandparentId).ClientId(clientId).Execute()

Delete ConfigurationTypeQuestionValue

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
	grandparentId := int32(56) // int32 | typeId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationTypeQuestionValuesAPI.DeleteCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById(context.Background(), id, parentId, grandparentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTypeQuestionValuesAPI.DeleteCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById``: %v\n", err)
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
**grandparentId** | **int32** | typeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdRequest struct via the builder pattern


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


## GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValues

> []ConfigurationTypeQuestionValue GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValues(ctx, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ConfigurationTypeQuestionValue

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
	grandparentId := int32(56) // int32 | typeId
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
	resp, r, err := apiClient.ConfigurationTypeQuestionValuesAPI.GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValues(context.Background(), parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTypeQuestionValuesAPI.GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValues`: []ConfigurationTypeQuestionValue
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationTypeQuestionValuesAPI.GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | questionId | 
**grandparentId** | **int32** | typeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesRequest struct via the builder pattern


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

[**[]ConfigurationTypeQuestionValue**](ConfigurationTypeQuestionValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById

> ConfigurationTypeQuestionValue GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById(ctx, id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ConfigurationTypeQuestionValue

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
	grandparentId := int32(56) // int32 | typeId
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
	resp, r, err := apiClient.ConfigurationTypeQuestionValuesAPI.GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById(context.Background(), id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTypeQuestionValuesAPI.GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById`: ConfigurationTypeQuestionValue
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationTypeQuestionValuesAPI.GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | valueId | 
**parentId** | **int32** | questionId | 
**grandparentId** | **int32** | typeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdRequest struct via the builder pattern


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

[**ConfigurationTypeQuestionValue**](ConfigurationTypeQuestionValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdUsages

> []Usage GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdUsages(ctx, id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of Usage

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
	grandparentId := int32(56) // int32 | typeId
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
	resp, r, err := apiClient.ConfigurationTypeQuestionValuesAPI.GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdUsages(context.Background(), id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTypeQuestionValuesAPI.GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdUsages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdUsages`: []Usage
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationTypeQuestionValuesAPI.GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdUsages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | valueId | 
**parentId** | **int32** | questionId | 
**grandparentId** | **int32** | typeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdUsagesRequest struct via the builder pattern


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

[**[]Usage**](Usage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdUsagesList

> []Usage GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdUsagesList(ctx, id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of Usage

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
	grandparentId := int32(56) // int32 | typeId
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
	resp, r, err := apiClient.ConfigurationTypeQuestionValuesAPI.GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdUsagesList(context.Background(), id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTypeQuestionValuesAPI.GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdUsagesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdUsagesList`: []Usage
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationTypeQuestionValuesAPI.GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdUsagesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | valueId | 
**parentId** | **int32** | questionId | 
**grandparentId** | **int32** | typeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdUsagesListRequest struct via the builder pattern


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

[**[]Usage**](Usage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesCount

> Count GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesCount(ctx, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ConfigurationTypeQuestionValue

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
	grandparentId := int32(56) // int32 | typeId
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
	resp, r, err := apiClient.ConfigurationTypeQuestionValuesAPI.GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesCount(context.Background(), parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTypeQuestionValuesAPI.GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationTypeQuestionValuesAPI.GetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | questionId | 
**grandparentId** | **int32** | typeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesCountRequest struct via the builder pattern


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


## PatchCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById

> ConfigurationTypeQuestionValue PatchCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById(ctx, id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ConfigurationTypeQuestionValue

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
	grandparentId := int32(56) // int32 | typeId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationTypeQuestionValuesAPI.PatchCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById(context.Background(), id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTypeQuestionValuesAPI.PatchCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById`: ConfigurationTypeQuestionValue
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationTypeQuestionValuesAPI.PatchCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | valueId | 
**parentId** | **int32** | questionId | 
**grandparentId** | **int32** | typeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ConfigurationTypeQuestionValue**](ConfigurationTypeQuestionValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValues

> ConfigurationTypeQuestionValue PostCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValues(ctx, parentId, grandparentId).ClientId(clientId).ConfigurationTypeQuestionValue(configurationTypeQuestionValue).Execute()

Post ConfigurationTypeQuestionValue

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
	grandparentId := int32(56) // int32 | typeId
	clientId := "clientId_example" // string | 
	configurationTypeQuestionValue := *openapiclient.NewConfigurationTypeQuestionValue("Value_example") // ConfigurationTypeQuestionValue | configurationTypeQuestionValue

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationTypeQuestionValuesAPI.PostCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValues(context.Background(), parentId, grandparentId).ClientId(clientId).ConfigurationTypeQuestionValue(configurationTypeQuestionValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTypeQuestionValuesAPI.PostCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValues`: ConfigurationTypeQuestionValue
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationTypeQuestionValuesAPI.PostCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | questionId | 
**grandparentId** | **int32** | typeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **configurationTypeQuestionValue** | [**ConfigurationTypeQuestionValue**](ConfigurationTypeQuestionValue.md) | configurationTypeQuestionValue | 

### Return type

[**ConfigurationTypeQuestionValue**](ConfigurationTypeQuestionValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById

> ConfigurationTypeQuestionValue PutCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById(ctx, id, parentId, grandparentId).ClientId(clientId).ConfigurationTypeQuestionValue(configurationTypeQuestionValue).Execute()

Put ConfigurationTypeQuestionValue

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
	grandparentId := int32(56) // int32 | typeId
	clientId := "clientId_example" // string | 
	configurationTypeQuestionValue := *openapiclient.NewConfigurationTypeQuestionValue("Value_example") // ConfigurationTypeQuestionValue | configurationTypeQuestionValue

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationTypeQuestionValuesAPI.PutCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById(context.Background(), id, parentId, grandparentId).ClientId(clientId).ConfigurationTypeQuestionValue(configurationTypeQuestionValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationTypeQuestionValuesAPI.PutCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById`: ConfigurationTypeQuestionValue
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationTypeQuestionValuesAPI.PutCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | valueId | 
**parentId** | **int32** | questionId | 
**grandparentId** | **int32** | typeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyConfigurationsTypesByGrandparentIdQuestionsByParentIdValuesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **configurationTypeQuestionValue** | [**ConfigurationTypeQuestionValue**](ConfigurationTypeQuestionValue.md) | configurationTypeQuestionValue | 

### Return type

[**ConfigurationTypeQuestionValue**](ConfigurationTypeQuestionValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

