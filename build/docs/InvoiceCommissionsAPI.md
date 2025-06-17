# \InvoiceCommissionsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFinanceInvoicesByParentIdCommissions**](InvoiceCommissionsAPI.md#GetFinanceInvoicesByParentIdCommissions) | **Get** /finance/invoices/{parentId}/commissions | Get List of InvoiceCommissions
[**GetFinanceInvoicesByParentIdCommissionsById**](InvoiceCommissionsAPI.md#GetFinanceInvoicesByParentIdCommissionsById) | **Get** /finance/invoices/{parentId}/commissions/{id} | Get InvoiceCommissions
[**PatchFinanceInvoicesByParentIdCommissionsById**](InvoiceCommissionsAPI.md#PatchFinanceInvoicesByParentIdCommissionsById) | **Patch** /finance/invoices/{parentId}/commissions/{id} | Patch InvoiceCommissions
[**PostFinanceInvoicesByParentIdCommissionsRecalculate**](InvoiceCommissionsAPI.md#PostFinanceInvoicesByParentIdCommissionsRecalculate) | **Post** /finance/invoices/{parentId}/commissions/recalculate | Recalculate InvoiceCommissions



## GetFinanceInvoicesByParentIdCommissions

> []InvoiceCommission GetFinanceInvoicesByParentIdCommissions(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of InvoiceCommissions

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
	parentId := int32(56) // int32 | invoiceId
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
	resp, r, err := apiClient.InvoiceCommissionsAPI.GetFinanceInvoicesByParentIdCommissions(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceCommissionsAPI.GetFinanceInvoicesByParentIdCommissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceInvoicesByParentIdCommissions`: []InvoiceCommission
	fmt.Fprintf(os.Stdout, "Response from `InvoiceCommissionsAPI.GetFinanceInvoicesByParentIdCommissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | invoiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceInvoicesByParentIdCommissionsRequest struct via the builder pattern


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

[**[]InvoiceCommission**](InvoiceCommission.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceInvoicesByParentIdCommissionsById

> InvoiceCommission GetFinanceInvoicesByParentIdCommissionsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get InvoiceCommissions

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
	id := int32(56) // int32 | paymentId
	parentId := int32(56) // int32 | invoiceId
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
	resp, r, err := apiClient.InvoiceCommissionsAPI.GetFinanceInvoicesByParentIdCommissionsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceCommissionsAPI.GetFinanceInvoicesByParentIdCommissionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceInvoicesByParentIdCommissionsById`: InvoiceCommission
	fmt.Fprintf(os.Stdout, "Response from `InvoiceCommissionsAPI.GetFinanceInvoicesByParentIdCommissionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | paymentId | 
**parentId** | **int32** | invoiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceInvoicesByParentIdCommissionsByIdRequest struct via the builder pattern


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

[**InvoiceCommission**](InvoiceCommission.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFinanceInvoicesByParentIdCommissionsById

> InvoiceCommission PatchFinanceInvoicesByParentIdCommissionsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch InvoiceCommissions

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
	id := int32(56) // int32 | InvoiceCommissionId
	parentId := int32(56) // int32 | invoiceId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceCommissionsAPI.PatchFinanceInvoicesByParentIdCommissionsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceCommissionsAPI.PatchFinanceInvoicesByParentIdCommissionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceInvoicesByParentIdCommissionsById`: InvoiceCommission
	fmt.Fprintf(os.Stdout, "Response from `InvoiceCommissionsAPI.PatchFinanceInvoicesByParentIdCommissionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | InvoiceCommissionId | 
**parentId** | **int32** | invoiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceInvoicesByParentIdCommissionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**InvoiceCommission**](InvoiceCommission.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceInvoicesByParentIdCommissionsRecalculate

> PostFinanceInvoicesByParentIdCommissionsRecalculate(ctx, parentId).ClientId(clientId).Execute()

Recalculate InvoiceCommissions

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
	parentId := int32(56) // int32 | invoiceId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InvoiceCommissionsAPI.PostFinanceInvoicesByParentIdCommissionsRecalculate(context.Background(), parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceCommissionsAPI.PostFinanceInvoicesByParentIdCommissionsRecalculate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | invoiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceInvoicesByParentIdCommissionsRecalculateRequest struct via the builder pattern


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

