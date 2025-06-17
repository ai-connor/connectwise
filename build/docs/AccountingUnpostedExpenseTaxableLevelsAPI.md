# \AccountingUnpostedExpenseTaxableLevelsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevels**](AccountingUnpostedExpenseTaxableLevelsAPI.md#GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevels) | **Get** /finance/accounting/unpostedexpenses/{parentId}/taxableLevels | Get List of UnpostedExpenseTaxableLevel
[**GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevelsById**](AccountingUnpostedExpenseTaxableLevelsAPI.md#GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevelsById) | **Get** /finance/accounting/unpostedexpenses/{parentId}/taxableLevels/{id} | Get UnpostedExpenseTaxableLevel
[**GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevelsCount**](AccountingUnpostedExpenseTaxableLevelsAPI.md#GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevelsCount) | **Get** /finance/accounting/unpostedexpenses/{parentId}/taxableLevels/count | Get Count of UnpostedExpenseTaxableLevel



## GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevels

> []UnpostedExpenseTaxableLevel GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevels(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of UnpostedExpenseTaxableLevel

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
	parentId := int32(56) // int32 | unpostedexpensId
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
	resp, r, err := apiClient.AccountingUnpostedExpenseTaxableLevelsAPI.GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevels(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingUnpostedExpenseTaxableLevelsAPI.GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevels`: []UnpostedExpenseTaxableLevel
	fmt.Fprintf(os.Stdout, "Response from `AccountingUnpostedExpenseTaxableLevelsAPI.GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | unpostedexpensId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAccountingUnpostedexpensesByParentIdTaxableLevelsRequest struct via the builder pattern


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

[**[]UnpostedExpenseTaxableLevel**](UnpostedExpenseTaxableLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevelsById

> UnpostedExpenseTaxableLevel GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevelsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get UnpostedExpenseTaxableLevel

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
	parentId := int32(56) // int32 | unpostedexpensId
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
	resp, r, err := apiClient.AccountingUnpostedExpenseTaxableLevelsAPI.GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevelsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingUnpostedExpenseTaxableLevelsAPI.GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevelsById`: UnpostedExpenseTaxableLevel
	fmt.Fprintf(os.Stdout, "Response from `AccountingUnpostedExpenseTaxableLevelsAPI.GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevelsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxableLevelId | 
**parentId** | **int32** | unpostedexpensId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAccountingUnpostedexpensesByParentIdTaxableLevelsByIdRequest struct via the builder pattern


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

[**UnpostedExpenseTaxableLevel**](UnpostedExpenseTaxableLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevelsCount

> Count GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevelsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of UnpostedExpenseTaxableLevel

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
	parentId := int32(56) // int32 | unpostedexpensId
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
	resp, r, err := apiClient.AccountingUnpostedExpenseTaxableLevelsAPI.GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevelsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingUnpostedExpenseTaxableLevelsAPI.GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevelsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevelsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `AccountingUnpostedExpenseTaxableLevelsAPI.GetFinanceAccountingUnpostedexpensesByParentIdTaxableLevelsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | unpostedexpensId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAccountingUnpostedexpensesByParentIdTaxableLevelsCountRequest struct via the builder pattern


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

