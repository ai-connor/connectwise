# \PurchasingDemandsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostProcurementPurchasingDemands**](PurchasingDemandsAPI.md#PostProcurementPurchasingDemands) | **Post** /procurement/purchasingDemands | Post PurchasingDemand



## PostProcurementPurchasingDemands

> PurchasingDemand PostProcurementPurchasingDemands(ctx).ClientId(clientId).PurchasingDemand(purchasingDemand).Execute()

Post PurchasingDemand

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
	clientId := "clientId_example" // string | 
	purchasingDemand := *openapiclient.NewPurchasingDemand() // PurchasingDemand | purchasingDemand

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchasingDemandsAPI.PostProcurementPurchasingDemands(context.Background()).ClientId(clientId).PurchasingDemand(purchasingDemand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchasingDemandsAPI.PostProcurementPurchasingDemands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementPurchasingDemands`: PurchasingDemand
	fmt.Fprintf(os.Stdout, "Response from `PurchasingDemandsAPI.PostProcurementPurchasingDemands`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementPurchasingDemandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **purchasingDemand** | [**PurchasingDemand**](PurchasingDemand.md) | purchasingDemand | 

### Return type

[**PurchasingDemand**](PurchasingDemand.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

