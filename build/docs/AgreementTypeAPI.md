# \AgreementTypeAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostFinanceAgreementsByIdCopy**](AgreementTypeAPI.md#PostFinanceAgreementsByIdCopy) | **Post** /finance/agreements/types{id}/copy | Post AgreementType



## PostFinanceAgreementsByIdCopy

> AgreementType PostFinanceAgreementsByIdCopy(ctx, id).ClientId(clientId).Execute()

Post AgreementType

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
	id := int32(56) // int32 | agreementTypeId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgreementTypeAPI.PostFinanceAgreementsByIdCopy(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeAPI.PostFinanceAgreementsByIdCopy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceAgreementsByIdCopy`: AgreementType
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeAPI.PostFinanceAgreementsByIdCopy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceAgreementsByIdCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 

### Return type

[**AgreementType**](AgreementType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

