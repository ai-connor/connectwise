# \AccountingUnpostedProcurementTaxableLevelsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevels**](AccountingUnpostedProcurementTaxableLevelsAPI.md#GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevels) | **Get** /finance/accounting/unpostedprocurement/{parentId}/taxableLevels | Get List of UnpostedProcurementTaxableLevel
[**GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevelsById**](AccountingUnpostedProcurementTaxableLevelsAPI.md#GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevelsById) | **Get** /finance/accounting/unpostedprocurement/{parentId}/taxableLevels/{id} | Get UnpostedProcurementTaxableLevel
[**GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevelsCount**](AccountingUnpostedProcurementTaxableLevelsAPI.md#GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevelsCount) | **Get** /finance/accounting/unpostedprocurement/{parentId}/taxableLevels/count | Get Count of UnpostedProcurementTaxableLevel



## GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevels

> []UnpostedProcurementTaxableLevel GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevels(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of UnpostedProcurementTaxableLevel

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
	parentId := int32(56) // int32 | unpostedprocurementId
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
	resp, r, err := apiClient.AccountingUnpostedProcurementTaxableLevelsAPI.GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevels(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingUnpostedProcurementTaxableLevelsAPI.GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevels`: []UnpostedProcurementTaxableLevel
	fmt.Fprintf(os.Stdout, "Response from `AccountingUnpostedProcurementTaxableLevelsAPI.GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | unpostedprocurementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAccountingUnpostedprocurementByParentIdTaxableLevelsRequest struct via the builder pattern


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

[**[]UnpostedProcurementTaxableLevel**](UnpostedProcurementTaxableLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevelsById

> UnpostedProcurementTaxableLevel GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevelsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get UnpostedProcurementTaxableLevel

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
	id := int32(56) // int32 | taxableLevelId
	parentId := int32(56) // int32 | unpostedprocurementId
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
	resp, r, err := apiClient.AccountingUnpostedProcurementTaxableLevelsAPI.GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevelsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingUnpostedProcurementTaxableLevelsAPI.GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevelsById`: UnpostedProcurementTaxableLevel
	fmt.Fprintf(os.Stdout, "Response from `AccountingUnpostedProcurementTaxableLevelsAPI.GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevelsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxableLevelId | 
**parentId** | **int32** | unpostedprocurementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAccountingUnpostedprocurementByParentIdTaxableLevelsByIdRequest struct via the builder pattern


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

[**UnpostedProcurementTaxableLevel**](UnpostedProcurementTaxableLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevelsCount

> Count GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevelsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of UnpostedProcurementTaxableLevel

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
	parentId := int32(56) // int32 | unpostedprocurementId
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
	resp, r, err := apiClient.AccountingUnpostedProcurementTaxableLevelsAPI.GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevelsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingUnpostedProcurementTaxableLevelsAPI.GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevelsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevelsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `AccountingUnpostedProcurementTaxableLevelsAPI.GetFinanceAccountingUnpostedprocurementByParentIdTaxableLevelsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | unpostedprocurementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAccountingUnpostedprocurementByParentIdTaxableLevelsCountRequest struct via the builder pattern


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

