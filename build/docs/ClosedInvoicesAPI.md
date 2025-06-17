# \ClosedInvoicesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PatchFinanceClosedInvoicesById**](ClosedInvoicesAPI.md#PatchFinanceClosedInvoicesById) | **Patch** /finance/closedInvoices/{id} | Patch ClosedInvoice
[**PutFinanceClosedInvoicesById**](ClosedInvoicesAPI.md#PutFinanceClosedInvoicesById) | **Put** /finance/closedInvoices/{id} | Put ClosedInvoice



## PatchFinanceClosedInvoicesById

> ClosedInvoice PatchFinanceClosedInvoicesById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ClosedInvoice

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
	id := int32(56) // int32 | closedInvoiceId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClosedInvoicesAPI.PatchFinanceClosedInvoicesById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClosedInvoicesAPI.PatchFinanceClosedInvoicesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceClosedInvoicesById`: ClosedInvoice
	fmt.Fprintf(os.Stdout, "Response from `ClosedInvoicesAPI.PatchFinanceClosedInvoicesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | closedInvoiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceClosedInvoicesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ClosedInvoice**](ClosedInvoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceClosedInvoicesById

> ClosedInvoice PutFinanceClosedInvoicesById(ctx, id).ClientId(clientId).ClosedInvoice(closedInvoice).Execute()

Put ClosedInvoice

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
	id := int32(56) // int32 | closedInvoiceId
	clientId := "clientId_example" // string | 
	closedInvoice := *openapiclient.NewClosedInvoice() // ClosedInvoice | closedInvoice

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClosedInvoicesAPI.PutFinanceClosedInvoicesById(context.Background(), id).ClientId(clientId).ClosedInvoice(closedInvoice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClosedInvoicesAPI.PutFinanceClosedInvoicesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceClosedInvoicesById`: ClosedInvoice
	fmt.Fprintf(os.Stdout, "Response from `ClosedInvoicesAPI.PutFinanceClosedInvoicesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | closedInvoiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceClosedInvoicesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **closedInvoice** | [**ClosedInvoice**](ClosedInvoice.md) | closedInvoice | 

### Return type

[**ClosedInvoice**](ClosedInvoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

