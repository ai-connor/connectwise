# \AgreementTypeWorkTypeExclusionsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceAgreementTypesByParentIdWorkTypeExclusionsById**](AgreementTypeWorkTypeExclusionsAPI.md#DeleteFinanceAgreementTypesByParentIdWorkTypeExclusionsById) | **Delete** /finance/agreementTypes/{parentId}/workTypeExclusions/{id} | Delete AgreementTypeWorkTypeExclusion
[**GetFinanceAgreementTypesByParentIdWorkTypeExclusions**](AgreementTypeWorkTypeExclusionsAPI.md#GetFinanceAgreementTypesByParentIdWorkTypeExclusions) | **Get** /finance/agreementTypes/{parentId}/workTypeExclusions | Get List of AgreementTypeWorkTypeExclusion
[**GetFinanceAgreementTypesByParentIdWorkTypeExclusionsById**](AgreementTypeWorkTypeExclusionsAPI.md#GetFinanceAgreementTypesByParentIdWorkTypeExclusionsById) | **Get** /finance/agreementTypes/{parentId}/workTypeExclusions/{id} | Get AgreementTypeWorkTypeExclusion
[**GetFinanceAgreementTypesByParentIdWorkTypeExclusionsCount**](AgreementTypeWorkTypeExclusionsAPI.md#GetFinanceAgreementTypesByParentIdWorkTypeExclusionsCount) | **Get** /finance/agreementTypes/{parentId}/workTypeExclusions/count | Get Count of AgreementTypeWorkTypeExclusion
[**PostFinanceAgreementTypesByParentIdWorkTypeExclusions**](AgreementTypeWorkTypeExclusionsAPI.md#PostFinanceAgreementTypesByParentIdWorkTypeExclusions) | **Post** /finance/agreementTypes/{parentId}/workTypeExclusions | Post AgreementTypeWorkTypeExclusion



## DeleteFinanceAgreementTypesByParentIdWorkTypeExclusionsById

> DeleteFinanceAgreementTypesByParentIdWorkTypeExclusionsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete AgreementTypeWorkTypeExclusion

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
	id := int32(56) // int32 | workTypeExclusionId
	parentId := int32(56) // int32 | agreementTypeId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgreementTypeWorkTypeExclusionsAPI.DeleteFinanceAgreementTypesByParentIdWorkTypeExclusionsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkTypeExclusionsAPI.DeleteFinanceAgreementTypesByParentIdWorkTypeExclusionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | workTypeExclusionId | 
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceAgreementTypesByParentIdWorkTypeExclusionsByIdRequest struct via the builder pattern


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


## GetFinanceAgreementTypesByParentIdWorkTypeExclusions

> []AgreementTypeWorkTypeExclusion GetFinanceAgreementTypesByParentIdWorkTypeExclusions(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of AgreementTypeWorkTypeExclusion

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
	parentId := int32(56) // int32 | agreementTypeId
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
	resp, r, err := apiClient.AgreementTypeWorkTypeExclusionsAPI.GetFinanceAgreementTypesByParentIdWorkTypeExclusions(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkTypeExclusionsAPI.GetFinanceAgreementTypesByParentIdWorkTypeExclusions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementTypesByParentIdWorkTypeExclusions`: []AgreementTypeWorkTypeExclusion
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeWorkTypeExclusionsAPI.GetFinanceAgreementTypesByParentIdWorkTypeExclusions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementTypesByParentIdWorkTypeExclusionsRequest struct via the builder pattern


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

[**[]AgreementTypeWorkTypeExclusion**](AgreementTypeWorkTypeExclusion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAgreementTypesByParentIdWorkTypeExclusionsById

> AgreementTypeWorkTypeExclusion GetFinanceAgreementTypesByParentIdWorkTypeExclusionsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get AgreementTypeWorkTypeExclusion

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
	id := int32(56) // int32 | workTypeExclusionId
	parentId := int32(56) // int32 | agreementTypeId
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
	resp, r, err := apiClient.AgreementTypeWorkTypeExclusionsAPI.GetFinanceAgreementTypesByParentIdWorkTypeExclusionsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkTypeExclusionsAPI.GetFinanceAgreementTypesByParentIdWorkTypeExclusionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementTypesByParentIdWorkTypeExclusionsById`: AgreementTypeWorkTypeExclusion
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeWorkTypeExclusionsAPI.GetFinanceAgreementTypesByParentIdWorkTypeExclusionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | workTypeExclusionId | 
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementTypesByParentIdWorkTypeExclusionsByIdRequest struct via the builder pattern


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

[**AgreementTypeWorkTypeExclusion**](AgreementTypeWorkTypeExclusion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAgreementTypesByParentIdWorkTypeExclusionsCount

> Count GetFinanceAgreementTypesByParentIdWorkTypeExclusionsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of AgreementTypeWorkTypeExclusion

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
	parentId := int32(56) // int32 | agreementTypeId
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
	resp, r, err := apiClient.AgreementTypeWorkTypeExclusionsAPI.GetFinanceAgreementTypesByParentIdWorkTypeExclusionsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkTypeExclusionsAPI.GetFinanceAgreementTypesByParentIdWorkTypeExclusionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementTypesByParentIdWorkTypeExclusionsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeWorkTypeExclusionsAPI.GetFinanceAgreementTypesByParentIdWorkTypeExclusionsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementTypesByParentIdWorkTypeExclusionsCountRequest struct via the builder pattern


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


## PostFinanceAgreementTypesByParentIdWorkTypeExclusions

> AgreementTypeWorkTypeExclusion PostFinanceAgreementTypesByParentIdWorkTypeExclusions(ctx, parentId).ClientId(clientId).AgreementTypeWorkTypeExclusion(agreementTypeWorkTypeExclusion).Execute()

Post AgreementTypeWorkTypeExclusion

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
	parentId := int32(56) // int32 | agreementTypeId
	clientId := "clientId_example" // string | 
	agreementTypeWorkTypeExclusion := *openapiclient.NewAgreementTypeWorkTypeExclusion(*openapiclient.NewWorkTypeReference()) // AgreementTypeWorkTypeExclusion | workTypeExclusion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgreementTypeWorkTypeExclusionsAPI.PostFinanceAgreementTypesByParentIdWorkTypeExclusions(context.Background(), parentId).ClientId(clientId).AgreementTypeWorkTypeExclusion(agreementTypeWorkTypeExclusion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkTypeExclusionsAPI.PostFinanceAgreementTypesByParentIdWorkTypeExclusions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceAgreementTypesByParentIdWorkTypeExclusions`: AgreementTypeWorkTypeExclusion
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeWorkTypeExclusionsAPI.PostFinanceAgreementTypesByParentIdWorkTypeExclusions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceAgreementTypesByParentIdWorkTypeExclusionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **agreementTypeWorkTypeExclusion** | [**AgreementTypeWorkTypeExclusion**](AgreementTypeWorkTypeExclusion.md) | workTypeExclusion | 

### Return type

[**AgreementTypeWorkTypeExclusion**](AgreementTypeWorkTypeExclusion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

