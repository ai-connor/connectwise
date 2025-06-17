# \InvoiceRoutingsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceInvoicesByParentIdRoutingsById**](InvoiceRoutingsAPI.md#DeleteFinanceInvoicesByParentIdRoutingsById) | **Delete** /finance/invoices/{parentId}/routings/{id} | Delete Invoice Routings
[**GetFinanceInvoicesByParentIdRoutings**](InvoiceRoutingsAPI.md#GetFinanceInvoicesByParentIdRoutings) | **Get** /finance/invoices/{parentId}/routings | Get List of Invoice Routings
[**GetFinanceInvoicesByParentIdRoutingsById**](InvoiceRoutingsAPI.md#GetFinanceInvoicesByParentIdRoutingsById) | **Get** /finance/invoices/{parentId}/routings/{id} | Get Invoice Routings
[**GetFinanceInvoicesByParentIdRoutingsCount**](InvoiceRoutingsAPI.md#GetFinanceInvoicesByParentIdRoutingsCount) | **Get** /finance/invoices/{parentId}/routings/count | Get Count of Invoice Routings
[**PatchFinanceInvoicesByParentIdRoutingsById**](InvoiceRoutingsAPI.md#PatchFinanceInvoicesByParentIdRoutingsById) | **Patch** /finance/invoices/{parentId}/routings/{id} | Patch Invoice Routings
[**PostFinanceInvoicesByParentIdRoutings**](InvoiceRoutingsAPI.md#PostFinanceInvoicesByParentIdRoutings) | **Post** /finance/invoices/{parentId}/routings | Post Invoice Routings
[**PutFinanceInvoicesByParentIdRoutingsById**](InvoiceRoutingsAPI.md#PutFinanceInvoicesByParentIdRoutingsById) | **Put** /finance/invoices/{parentId}/routings/{id} | Put Invoice Routings



## DeleteFinanceInvoicesByParentIdRoutingsById

> DeleteFinanceInvoicesByParentIdRoutingsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete Invoice Routings

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
	id := int32(56) // int32 | InvoiceRoutingsId
	parentId := int32(56) // int32 | Invoice
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InvoiceRoutingsAPI.DeleteFinanceInvoicesByParentIdRoutingsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceRoutingsAPI.DeleteFinanceInvoicesByParentIdRoutingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | InvoiceRoutingsId | 
**parentId** | **int32** | Invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceInvoicesByParentIdRoutingsByIdRequest struct via the builder pattern


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


## GetFinanceInvoicesByParentIdRoutings

> []InvoiceRouting GetFinanceInvoicesByParentIdRoutings(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of Invoice Routings

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
	parentId := int32(56) // int32 | Invoice
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
	resp, r, err := apiClient.InvoiceRoutingsAPI.GetFinanceInvoicesByParentIdRoutings(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceRoutingsAPI.GetFinanceInvoicesByParentIdRoutings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceInvoicesByParentIdRoutings`: []InvoiceRouting
	fmt.Fprintf(os.Stdout, "Response from `InvoiceRoutingsAPI.GetFinanceInvoicesByParentIdRoutings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | Invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceInvoicesByParentIdRoutingsRequest struct via the builder pattern


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

[**[]InvoiceRouting**](InvoiceRouting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceInvoicesByParentIdRoutingsById

> InvoiceRouting GetFinanceInvoicesByParentIdRoutingsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Invoice Routings

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
	id := int32(56) // int32 | InvoiceRoutingsId
	parentId := int32(56) // int32 | Invoice
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
	resp, r, err := apiClient.InvoiceRoutingsAPI.GetFinanceInvoicesByParentIdRoutingsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceRoutingsAPI.GetFinanceInvoicesByParentIdRoutingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceInvoicesByParentIdRoutingsById`: InvoiceRouting
	fmt.Fprintf(os.Stdout, "Response from `InvoiceRoutingsAPI.GetFinanceInvoicesByParentIdRoutingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | InvoiceRoutingsId | 
**parentId** | **int32** | Invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceInvoicesByParentIdRoutingsByIdRequest struct via the builder pattern


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

[**InvoiceRouting**](InvoiceRouting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceInvoicesByParentIdRoutingsCount

> Count GetFinanceInvoicesByParentIdRoutingsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of Invoice Routings

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
	parentId := int32(56) // int32 | Invoice
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
	resp, r, err := apiClient.InvoiceRoutingsAPI.GetFinanceInvoicesByParentIdRoutingsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceRoutingsAPI.GetFinanceInvoicesByParentIdRoutingsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceInvoicesByParentIdRoutingsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `InvoiceRoutingsAPI.GetFinanceInvoicesByParentIdRoutingsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | Invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceInvoicesByParentIdRoutingsCountRequest struct via the builder pattern


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


## PatchFinanceInvoicesByParentIdRoutingsById

> InvoiceRouting PatchFinanceInvoicesByParentIdRoutingsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch Invoice Routings

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
	id := int32(56) // int32 | InvoiceRoutingsId
	parentId := int32(56) // int32 | Invoice
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceRoutingsAPI.PatchFinanceInvoicesByParentIdRoutingsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceRoutingsAPI.PatchFinanceInvoicesByParentIdRoutingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceInvoicesByParentIdRoutingsById`: InvoiceRouting
	fmt.Fprintf(os.Stdout, "Response from `InvoiceRoutingsAPI.PatchFinanceInvoicesByParentIdRoutingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | InvoiceRoutingsId | 
**parentId** | **int32** | Invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceInvoicesByParentIdRoutingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**InvoiceRouting**](InvoiceRouting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceInvoicesByParentIdRoutings

> InvoiceRouting PostFinanceInvoicesByParentIdRoutings(ctx, parentId).ClientId(clientId).InvoiceRouting(invoiceRouting).Execute()

Post Invoice Routings

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
	parentId := int32(56) // int32 | Invoice
	clientId := "clientId_example" // string | 
	invoiceRouting := *openapiclient.NewInvoiceRouting() // InvoiceRouting | InvoiceRouting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceRoutingsAPI.PostFinanceInvoicesByParentIdRoutings(context.Background(), parentId).ClientId(clientId).InvoiceRouting(invoiceRouting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceRoutingsAPI.PostFinanceInvoicesByParentIdRoutings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceInvoicesByParentIdRoutings`: InvoiceRouting
	fmt.Fprintf(os.Stdout, "Response from `InvoiceRoutingsAPI.PostFinanceInvoicesByParentIdRoutings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | Invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceInvoicesByParentIdRoutingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **invoiceRouting** | [**InvoiceRouting**](InvoiceRouting.md) | InvoiceRouting | 

### Return type

[**InvoiceRouting**](InvoiceRouting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceInvoicesByParentIdRoutingsById

> InvoiceRouting PutFinanceInvoicesByParentIdRoutingsById(ctx, id, parentId).ClientId(clientId).InvoiceRouting(invoiceRouting).Execute()

Put Invoice Routings

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
	id := int32(56) // int32 | InvoiceRoutingId
	parentId := int32(56) // int32 | Invoice
	clientId := "clientId_example" // string | 
	invoiceRouting := *openapiclient.NewInvoiceRouting() // InvoiceRouting | companyTypeAssociation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceRoutingsAPI.PutFinanceInvoicesByParentIdRoutingsById(context.Background(), id, parentId).ClientId(clientId).InvoiceRouting(invoiceRouting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceRoutingsAPI.PutFinanceInvoicesByParentIdRoutingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceInvoicesByParentIdRoutingsById`: InvoiceRouting
	fmt.Fprintf(os.Stdout, "Response from `InvoiceRoutingsAPI.PutFinanceInvoicesByParentIdRoutingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | InvoiceRoutingId | 
**parentId** | **int32** | Invoice | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceInvoicesByParentIdRoutingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **invoiceRouting** | [**InvoiceRouting**](InvoiceRouting.md) | companyTypeAssociation | 

### Return type

[**InvoiceRouting**](InvoiceRouting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

