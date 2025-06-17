# \AgreementWorkTypeExclusionsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceAgreementsByParentIdWorkTypeExclusionsById**](AgreementWorkTypeExclusionsAPI.md#DeleteFinanceAgreementsByParentIdWorkTypeExclusionsById) | **Delete** /finance/agreements/{parentId}/workTypeExclusions/{id} | Delete AgreementWorkTypeExclusion
[**GetFinanceAgreementsByParentIdWorkTypeExclusions**](AgreementWorkTypeExclusionsAPI.md#GetFinanceAgreementsByParentIdWorkTypeExclusions) | **Get** /finance/agreements/{parentId}/workTypeExclusions | Get List of AgreementWorkTypeExclusion
[**GetFinanceAgreementsByParentIdWorkTypeExclusionsCount**](AgreementWorkTypeExclusionsAPI.md#GetFinanceAgreementsByParentIdWorkTypeExclusionsCount) | **Get** /finance/agreements/{parentId}/workTypeExclusions/count | Get Count of AgreementWorkTypeExclusion
[**PostFinanceAgreementsByParentIdWorkTypeExclusions**](AgreementWorkTypeExclusionsAPI.md#PostFinanceAgreementsByParentIdWorkTypeExclusions) | **Post** /finance/agreements/{parentId}/workTypeExclusions | Post AgreementWorkTypeExclusion



## DeleteFinanceAgreementsByParentIdWorkTypeExclusionsById

> DeleteFinanceAgreementsByParentIdWorkTypeExclusionsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete AgreementWorkTypeExclusion

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
	parentId := int32(56) // int32 | agreementId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgreementWorkTypeExclusionsAPI.DeleteFinanceAgreementsByParentIdWorkTypeExclusionsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementWorkTypeExclusionsAPI.DeleteFinanceAgreementsByParentIdWorkTypeExclusionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | workTypeExclusionId | 
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceAgreementsByParentIdWorkTypeExclusionsByIdRequest struct via the builder pattern


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


## GetFinanceAgreementsByParentIdWorkTypeExclusions

> []AgreementWorkTypeExclusion GetFinanceAgreementsByParentIdWorkTypeExclusions(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of AgreementWorkTypeExclusion

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
	parentId := int32(56) // int32 | agreementId
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
	resp, r, err := apiClient.AgreementWorkTypeExclusionsAPI.GetFinanceAgreementsByParentIdWorkTypeExclusions(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementWorkTypeExclusionsAPI.GetFinanceAgreementsByParentIdWorkTypeExclusions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementsByParentIdWorkTypeExclusions`: []AgreementWorkTypeExclusion
	fmt.Fprintf(os.Stdout, "Response from `AgreementWorkTypeExclusionsAPI.GetFinanceAgreementsByParentIdWorkTypeExclusions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementsByParentIdWorkTypeExclusionsRequest struct via the builder pattern


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

[**[]AgreementWorkTypeExclusion**](AgreementWorkTypeExclusion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAgreementsByParentIdWorkTypeExclusionsCount

> Count GetFinanceAgreementsByParentIdWorkTypeExclusionsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of AgreementWorkTypeExclusion

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
	parentId := int32(56) // int32 | agreementId
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
	resp, r, err := apiClient.AgreementWorkTypeExclusionsAPI.GetFinanceAgreementsByParentIdWorkTypeExclusionsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementWorkTypeExclusionsAPI.GetFinanceAgreementsByParentIdWorkTypeExclusionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementsByParentIdWorkTypeExclusionsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `AgreementWorkTypeExclusionsAPI.GetFinanceAgreementsByParentIdWorkTypeExclusionsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementsByParentIdWorkTypeExclusionsCountRequest struct via the builder pattern


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


## PostFinanceAgreementsByParentIdWorkTypeExclusions

> AgreementWorkTypeExclusion PostFinanceAgreementsByParentIdWorkTypeExclusions(ctx, parentId).ClientId(clientId).AgreementWorkTypeExclusion(agreementWorkTypeExclusion).Execute()

Post AgreementWorkTypeExclusion

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
	parentId := int32(56) // int32 | agreementId
	clientId := "clientId_example" // string | 
	agreementWorkTypeExclusion := *openapiclient.NewAgreementWorkTypeExclusion(*openapiclient.NewWorkTypeReference()) // AgreementWorkTypeExclusion | workTypeExclusion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgreementWorkTypeExclusionsAPI.PostFinanceAgreementsByParentIdWorkTypeExclusions(context.Background(), parentId).ClientId(clientId).AgreementWorkTypeExclusion(agreementWorkTypeExclusion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementWorkTypeExclusionsAPI.PostFinanceAgreementsByParentIdWorkTypeExclusions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceAgreementsByParentIdWorkTypeExclusions`: AgreementWorkTypeExclusion
	fmt.Fprintf(os.Stdout, "Response from `AgreementWorkTypeExclusionsAPI.PostFinanceAgreementsByParentIdWorkTypeExclusions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceAgreementsByParentIdWorkTypeExclusionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **agreementWorkTypeExclusion** | [**AgreementWorkTypeExclusion**](AgreementWorkTypeExclusion.md) | workTypeExclusion | 

### Return type

[**AgreementWorkTypeExclusion**](AgreementWorkTypeExclusion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

