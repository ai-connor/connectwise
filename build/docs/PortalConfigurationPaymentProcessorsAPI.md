# \PortalConfigurationPaymentProcessorsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessors**](PortalConfigurationPaymentProcessorsAPI.md#GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessors) | **Get** /company/portalConfigurations/invoiceSetup/paymentProcessors | Get List of PortalConfigurationPaymentProcessor
[**GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessorsById**](PortalConfigurationPaymentProcessorsAPI.md#GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessorsById) | **Get** /company/portalConfigurations/invoiceSetup/paymentProcessors/{id} | Get PortalConfigurationPaymentProcessor
[**GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessorsCount**](PortalConfigurationPaymentProcessorsAPI.md#GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessorsCount) | **Get** /company/portalConfigurations/invoiceSetup/paymentProcessors/count | Get Count of PortalConfigurationPaymentProcessor



## GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessors

> []PortalConfigurationPaymentProcessor GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessors(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of PortalConfigurationPaymentProcessor

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
	resp, r, err := apiClient.PortalConfigurationPaymentProcessorsAPI.GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessors(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationPaymentProcessorsAPI.GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessors`: []PortalConfigurationPaymentProcessor
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationPaymentProcessorsAPI.GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPortalConfigurationsInvoiceSetupPaymentProcessorsRequest struct via the builder pattern


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

[**[]PortalConfigurationPaymentProcessor**](PortalConfigurationPaymentProcessor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessorsById

> PortalConfigurationPaymentProcessor GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessorsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get PortalConfigurationPaymentProcessor

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
	id := int32(56) // int32 | paymentProcessorId
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
	resp, r, err := apiClient.PortalConfigurationPaymentProcessorsAPI.GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessorsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationPaymentProcessorsAPI.GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessorsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessorsById`: PortalConfigurationPaymentProcessor
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationPaymentProcessorsAPI.GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessorsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | paymentProcessorId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPortalConfigurationsInvoiceSetupPaymentProcessorsByIdRequest struct via the builder pattern


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

[**PortalConfigurationPaymentProcessor**](PortalConfigurationPaymentProcessor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessorsCount

> Count GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessorsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of PortalConfigurationPaymentProcessor

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
	resp, r, err := apiClient.PortalConfigurationPaymentProcessorsAPI.GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessorsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationPaymentProcessorsAPI.GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessorsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessorsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationPaymentProcessorsAPI.GetCompanyPortalConfigurationsInvoiceSetupPaymentProcessorsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPortalConfigurationsInvoiceSetupPaymentProcessorsCountRequest struct via the builder pattern


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

