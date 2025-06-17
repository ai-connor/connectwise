# \AccountingUnpostedInvoiceTaxableLevelsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevels**](AccountingUnpostedInvoiceTaxableLevelsAPI.md#GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevels) | **Get** /finance/accounting/unpostedinvoices/{parentId}/taxableLevels | Get List of UnpostedInvoiceTaxableLevel
[**GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevelsById**](AccountingUnpostedInvoiceTaxableLevelsAPI.md#GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevelsById) | **Get** /finance/accounting/unpostedinvoices/{parentId}/taxableLevels/{id} | Get UnpostedInvoiceTaxableLevel
[**GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevelsCount**](AccountingUnpostedInvoiceTaxableLevelsAPI.md#GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevelsCount) | **Get** /finance/accounting/unpostedinvoices/{parentId}/taxableLevels/count | Get Count of UnpostedInvoiceTaxableLevel



## GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevels

> []UnpostedInvoiceTaxableLevel GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevels(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of UnpostedInvoiceTaxableLevel

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
	parentId := int32(56) // int32 | unpostedinvoiceId
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
	resp, r, err := apiClient.AccountingUnpostedInvoiceTaxableLevelsAPI.GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevels(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingUnpostedInvoiceTaxableLevelsAPI.GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevels`: []UnpostedInvoiceTaxableLevel
	fmt.Fprintf(os.Stdout, "Response from `AccountingUnpostedInvoiceTaxableLevelsAPI.GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | unpostedinvoiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevelsRequest struct via the builder pattern


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

[**[]UnpostedInvoiceTaxableLevel**](UnpostedInvoiceTaxableLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevelsById

> UnpostedInvoiceTaxableLevel GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevelsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get UnpostedInvoiceTaxableLevel

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
	parentId := int32(56) // int32 | unpostedinvoiceId
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
	resp, r, err := apiClient.AccountingUnpostedInvoiceTaxableLevelsAPI.GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevelsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingUnpostedInvoiceTaxableLevelsAPI.GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevelsById`: UnpostedInvoiceTaxableLevel
	fmt.Fprintf(os.Stdout, "Response from `AccountingUnpostedInvoiceTaxableLevelsAPI.GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevelsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxableLevelId | 
**parentId** | **int32** | unpostedinvoiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevelsByIdRequest struct via the builder pattern


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

[**UnpostedInvoiceTaxableLevel**](UnpostedInvoiceTaxableLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevelsCount

> Count GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevelsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of UnpostedInvoiceTaxableLevel

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
	parentId := int32(56) // int32 | unpostedinvoiceId
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
	resp, r, err := apiClient.AccountingUnpostedInvoiceTaxableLevelsAPI.GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevelsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingUnpostedInvoiceTaxableLevelsAPI.GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevelsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevelsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `AccountingUnpostedInvoiceTaxableLevelsAPI.GetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevelsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | unpostedinvoiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAccountingUnpostedinvoicesByParentIdTaxableLevelsCountRequest struct via the builder pattern


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

