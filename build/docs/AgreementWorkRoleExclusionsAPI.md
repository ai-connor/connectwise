# \AgreementWorkRoleExclusionsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceAgreementsByParentIdWorkRoleExclusionsById**](AgreementWorkRoleExclusionsAPI.md#DeleteFinanceAgreementsByParentIdWorkRoleExclusionsById) | **Delete** /finance/agreements/{parentId}/workRoleExclusions/{id} | Delete AgreementWorkRoleExclusion
[**GetFinanceAgreementsByParentIdWorkRoleExclusions**](AgreementWorkRoleExclusionsAPI.md#GetFinanceAgreementsByParentIdWorkRoleExclusions) | **Get** /finance/agreements/{parentId}/workRoleExclusions | Get List of AgreementWorkRoleExclusion
[**GetFinanceAgreementsByParentIdWorkRoleExclusionsCount**](AgreementWorkRoleExclusionsAPI.md#GetFinanceAgreementsByParentIdWorkRoleExclusionsCount) | **Get** /finance/agreements/{parentId}/workRoleExclusions/count | Get Count of AgreementWorkRoleExclusion
[**PostFinanceAgreementsByParentIdWorkRoleExclusions**](AgreementWorkRoleExclusionsAPI.md#PostFinanceAgreementsByParentIdWorkRoleExclusions) | **Post** /finance/agreements/{parentId}/workRoleExclusions | Post AgreementWorkRoleExclusion



## DeleteFinanceAgreementsByParentIdWorkRoleExclusionsById

> DeleteFinanceAgreementsByParentIdWorkRoleExclusionsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete AgreementWorkRoleExclusion

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
	id := int32(56) // int32 | workRoleExclusionId
	parentId := int32(56) // int32 | agreementId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgreementWorkRoleExclusionsAPI.DeleteFinanceAgreementsByParentIdWorkRoleExclusionsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementWorkRoleExclusionsAPI.DeleteFinanceAgreementsByParentIdWorkRoleExclusionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | workRoleExclusionId | 
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceAgreementsByParentIdWorkRoleExclusionsByIdRequest struct via the builder pattern


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


## GetFinanceAgreementsByParentIdWorkRoleExclusions

> []AgreementWorkRoleExclusion GetFinanceAgreementsByParentIdWorkRoleExclusions(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of AgreementWorkRoleExclusion

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
	resp, r, err := apiClient.AgreementWorkRoleExclusionsAPI.GetFinanceAgreementsByParentIdWorkRoleExclusions(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementWorkRoleExclusionsAPI.GetFinanceAgreementsByParentIdWorkRoleExclusions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementsByParentIdWorkRoleExclusions`: []AgreementWorkRoleExclusion
	fmt.Fprintf(os.Stdout, "Response from `AgreementWorkRoleExclusionsAPI.GetFinanceAgreementsByParentIdWorkRoleExclusions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementsByParentIdWorkRoleExclusionsRequest struct via the builder pattern


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

[**[]AgreementWorkRoleExclusion**](AgreementWorkRoleExclusion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAgreementsByParentIdWorkRoleExclusionsCount

> Count GetFinanceAgreementsByParentIdWorkRoleExclusionsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of AgreementWorkRoleExclusion

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
	resp, r, err := apiClient.AgreementWorkRoleExclusionsAPI.GetFinanceAgreementsByParentIdWorkRoleExclusionsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementWorkRoleExclusionsAPI.GetFinanceAgreementsByParentIdWorkRoleExclusionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementsByParentIdWorkRoleExclusionsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `AgreementWorkRoleExclusionsAPI.GetFinanceAgreementsByParentIdWorkRoleExclusionsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementsByParentIdWorkRoleExclusionsCountRequest struct via the builder pattern


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


## PostFinanceAgreementsByParentIdWorkRoleExclusions

> AgreementWorkRoleExclusion PostFinanceAgreementsByParentIdWorkRoleExclusions(ctx, parentId).ClientId(clientId).AgreementWorkRoleExclusion(agreementWorkRoleExclusion).Execute()

Post AgreementWorkRoleExclusion

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
	agreementWorkRoleExclusion := *openapiclient.NewAgreementWorkRoleExclusion(*openapiclient.NewWorkRoleReference()) // AgreementWorkRoleExclusion | workRoleExclusion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgreementWorkRoleExclusionsAPI.PostFinanceAgreementsByParentIdWorkRoleExclusions(context.Background(), parentId).ClientId(clientId).AgreementWorkRoleExclusion(agreementWorkRoleExclusion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementWorkRoleExclusionsAPI.PostFinanceAgreementsByParentIdWorkRoleExclusions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceAgreementsByParentIdWorkRoleExclusions`: AgreementWorkRoleExclusion
	fmt.Fprintf(os.Stdout, "Response from `AgreementWorkRoleExclusionsAPI.PostFinanceAgreementsByParentIdWorkRoleExclusions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceAgreementsByParentIdWorkRoleExclusionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **agreementWorkRoleExclusion** | [**AgreementWorkRoleExclusion**](AgreementWorkRoleExclusion.md) | workRoleExclusion | 

### Return type

[**AgreementWorkRoleExclusion**](AgreementWorkRoleExclusion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

