# \InvoicePaymentsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceInvoicesByParentIdPaymentsById**](InvoicePaymentsAPI.md#DeleteFinanceInvoicesByParentIdPaymentsById) | **Delete** /finance/invoices/{parentId}/payments/{id} | Delete Payment
[**GetFinanceInvoicesByParentIdPayments**](InvoicePaymentsAPI.md#GetFinanceInvoicesByParentIdPayments) | **Get** /finance/invoices/{parentId}/payments | Get List of Payment
[**GetFinanceInvoicesByParentIdPaymentsById**](InvoicePaymentsAPI.md#GetFinanceInvoicesByParentIdPaymentsById) | **Get** /finance/invoices/{parentId}/payments/{id} | Get Payment
[**PatchFinanceInvoicesByParentIdPaymentsById**](InvoicePaymentsAPI.md#PatchFinanceInvoicesByParentIdPaymentsById) | **Patch** /finance/invoices/{parentId}/payments/{id} | Patch Payment
[**PostFinanceInvoicesByParentIdPayments**](InvoicePaymentsAPI.md#PostFinanceInvoicesByParentIdPayments) | **Post** /finance/invoices/{parentId}/payments | Post Payment
[**PutFinanceInvoicesByParentIdPaymentsById**](InvoicePaymentsAPI.md#PutFinanceInvoicesByParentIdPaymentsById) | **Put** /finance/invoices/{parentId}/payments/{id} | Put Payment



## DeleteFinanceInvoicesByParentIdPaymentsById

> DeleteFinanceInvoicesByParentIdPaymentsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete Payment

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InvoicePaymentsAPI.DeleteFinanceInvoicesByParentIdPaymentsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicePaymentsAPI.DeleteFinanceInvoicesByParentIdPaymentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | paymentId | 
**parentId** | **int32** | invoiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceInvoicesByParentIdPaymentsByIdRequest struct via the builder pattern


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


## GetFinanceInvoicesByParentIdPayments

> []InvoicePayment GetFinanceInvoicesByParentIdPayments(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of Payment

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
	resp, r, err := apiClient.InvoicePaymentsAPI.GetFinanceInvoicesByParentIdPayments(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicePaymentsAPI.GetFinanceInvoicesByParentIdPayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceInvoicesByParentIdPayments`: []InvoicePayment
	fmt.Fprintf(os.Stdout, "Response from `InvoicePaymentsAPI.GetFinanceInvoicesByParentIdPayments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | invoiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceInvoicesByParentIdPaymentsRequest struct via the builder pattern


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

[**[]InvoicePayment**](InvoicePayment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceInvoicesByParentIdPaymentsById

> InvoicePayment GetFinanceInvoicesByParentIdPaymentsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Payment

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
	resp, r, err := apiClient.InvoicePaymentsAPI.GetFinanceInvoicesByParentIdPaymentsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicePaymentsAPI.GetFinanceInvoicesByParentIdPaymentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceInvoicesByParentIdPaymentsById`: InvoicePayment
	fmt.Fprintf(os.Stdout, "Response from `InvoicePaymentsAPI.GetFinanceInvoicesByParentIdPaymentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | paymentId | 
**parentId** | **int32** | invoiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceInvoicesByParentIdPaymentsByIdRequest struct via the builder pattern


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

[**InvoicePayment**](InvoicePayment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFinanceInvoicesByParentIdPaymentsById

> InvoicePayment PatchFinanceInvoicesByParentIdPaymentsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch Payment

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
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicePaymentsAPI.PatchFinanceInvoicesByParentIdPaymentsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicePaymentsAPI.PatchFinanceInvoicesByParentIdPaymentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceInvoicesByParentIdPaymentsById`: InvoicePayment
	fmt.Fprintf(os.Stdout, "Response from `InvoicePaymentsAPI.PatchFinanceInvoicesByParentIdPaymentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | paymentId | 
**parentId** | **int32** | invoiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceInvoicesByParentIdPaymentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**InvoicePayment**](InvoicePayment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceInvoicesByParentIdPayments

> InvoicePayment PostFinanceInvoicesByParentIdPayments(ctx, parentId).ClientId(clientId).InvoicePayment(invoicePayment).Execute()

Post Payment

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
	invoicePayment := *openapiclient.NewInvoicePayment() // InvoicePayment | payment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicePaymentsAPI.PostFinanceInvoicesByParentIdPayments(context.Background(), parentId).ClientId(clientId).InvoicePayment(invoicePayment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicePaymentsAPI.PostFinanceInvoicesByParentIdPayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceInvoicesByParentIdPayments`: InvoicePayment
	fmt.Fprintf(os.Stdout, "Response from `InvoicePaymentsAPI.PostFinanceInvoicesByParentIdPayments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | invoiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceInvoicesByParentIdPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **invoicePayment** | [**InvoicePayment**](InvoicePayment.md) | payment | 

### Return type

[**InvoicePayment**](InvoicePayment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceInvoicesByParentIdPaymentsById

> InvoicePayment PutFinanceInvoicesByParentIdPaymentsById(ctx, id, parentId).ClientId(clientId).InvoicePayment(invoicePayment).Execute()

Put Payment

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
	invoicePayment := *openapiclient.NewInvoicePayment() // InvoicePayment | payment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicePaymentsAPI.PutFinanceInvoicesByParentIdPaymentsById(context.Background(), id, parentId).ClientId(clientId).InvoicePayment(invoicePayment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicePaymentsAPI.PutFinanceInvoicesByParentIdPaymentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceInvoicesByParentIdPaymentsById`: InvoicePayment
	fmt.Fprintf(os.Stdout, "Response from `InvoicePaymentsAPI.PutFinanceInvoicesByParentIdPaymentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | paymentId | 
**parentId** | **int32** | invoiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceInvoicesByParentIdPaymentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **invoicePayment** | [**InvoicePayment**](InvoicePayment.md) | payment | 

### Return type

[**InvoicePayment**](InvoicePayment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

