# \MarketplaceImportsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemMarketplaceimportGetdefinitionById**](MarketplaceImportsAPI.md#GetSystemMarketplaceimportGetdefinitionById) | **Get** /system/marketplaceimport/getdefinition/{id} | Get MarketplaceImport
[**PostSystemMarketplaceimportImport**](MarketplaceImportsAPI.md#PostSystemMarketplaceimportImport) | **Post** /system/marketplaceimport/import | Post MarketplaceImport



## GetSystemMarketplaceimportGetdefinitionById

> MarketplaceImport GetSystemMarketplaceimportGetdefinitionById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get MarketplaceImport

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
	id := int32(56) // int32 | getdefinitionId
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
	resp, r, err := apiClient.MarketplaceImportsAPI.GetSystemMarketplaceimportGetdefinitionById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceImportsAPI.GetSystemMarketplaceimportGetdefinitionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMarketplaceimportGetdefinitionById`: MarketplaceImport
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceImportsAPI.GetSystemMarketplaceimportGetdefinitionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | getdefinitionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMarketplaceimportGetdefinitionByIdRequest struct via the builder pattern


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

[**MarketplaceImport**](MarketplaceImport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemMarketplaceimportImport

> MarketplaceImport PostSystemMarketplaceimportImport(ctx).ClientId(clientId).MarketplaceImport(marketplaceImport).Execute()

Post MarketplaceImport

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
	marketplaceImport := *openapiclient.NewMarketplaceImport() // MarketplaceImport | marketplaceImport

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceImportsAPI.PostSystemMarketplaceimportImport(context.Background()).ClientId(clientId).MarketplaceImport(marketplaceImport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceImportsAPI.PostSystemMarketplaceimportImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemMarketplaceimportImport`: MarketplaceImport
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceImportsAPI.PostSystemMarketplaceimportImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemMarketplaceimportImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **marketplaceImport** | [**MarketplaceImport**](MarketplaceImport.md) | marketplaceImport | 

### Return type

[**MarketplaceImport**](MarketplaceImport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

