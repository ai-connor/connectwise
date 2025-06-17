# \OpportunityForecastItemsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSalesOpportunitiesByParentIdForecastById**](OpportunityForecastItemsAPI.md#DeleteSalesOpportunitiesByParentIdForecastById) | **Delete** /sales/opportunities/{parentId}/forecast/{id} | Delete ForecastItem
[**GetSalesOpportunitiesByParentIdForecastById**](OpportunityForecastItemsAPI.md#GetSalesOpportunitiesByParentIdForecastById) | **Get** /sales/opportunities/{parentId}/forecast/{id} | Get ForecastItem
[**PatchSalesOpportunitiesByParentIdForecastById**](OpportunityForecastItemsAPI.md#PatchSalesOpportunitiesByParentIdForecastById) | **Patch** /sales/opportunities/{parentId}/forecast/{id} | Patch ForecastItem
[**PostSalesOpportunitiesByParentIdForecastById**](OpportunityForecastItemsAPI.md#PostSalesOpportunitiesByParentIdForecastById) | **Post** /sales/opportunities/{parentId}/forecast/{id} | Post ForecastItem
[**PutSalesOpportunitiesByParentIdForecastById**](OpportunityForecastItemsAPI.md#PutSalesOpportunitiesByParentIdForecastById) | **Put** /sales/opportunities/{parentId}/forecast/{id} | Put ForecastItem



## DeleteSalesOpportunitiesByParentIdForecastById

> DeleteSalesOpportunitiesByParentIdForecastById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ForecastItem

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
	id := int32(56) // int32 | forecastId
	parentId := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OpportunityForecastItemsAPI.DeleteSalesOpportunitiesByParentIdForecastById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityForecastItemsAPI.DeleteSalesOpportunitiesByParentIdForecastById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | forecastId | 
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSalesOpportunitiesByParentIdForecastByIdRequest struct via the builder pattern


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


## GetSalesOpportunitiesByParentIdForecastById

> ForecastItem GetSalesOpportunitiesByParentIdForecastById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ForecastItem

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
	id := int32(56) // int32 | forecastId
	parentId := int32(56) // int32 | opportunityId
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
	resp, r, err := apiClient.OpportunityForecastItemsAPI.GetSalesOpportunitiesByParentIdForecastById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityForecastItemsAPI.GetSalesOpportunitiesByParentIdForecastById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesOpportunitiesByParentIdForecastById`: ForecastItem
	fmt.Fprintf(os.Stdout, "Response from `OpportunityForecastItemsAPI.GetSalesOpportunitiesByParentIdForecastById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | forecastId | 
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesOpportunitiesByParentIdForecastByIdRequest struct via the builder pattern


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

[**ForecastItem**](ForecastItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSalesOpportunitiesByParentIdForecastById

> ForecastItem PatchSalesOpportunitiesByParentIdForecastById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ForecastItem

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
	id := int32(56) // int32 | forecastId
	parentId := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpportunityForecastItemsAPI.PatchSalesOpportunitiesByParentIdForecastById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityForecastItemsAPI.PatchSalesOpportunitiesByParentIdForecastById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSalesOpportunitiesByParentIdForecastById`: ForecastItem
	fmt.Fprintf(os.Stdout, "Response from `OpportunityForecastItemsAPI.PatchSalesOpportunitiesByParentIdForecastById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | forecastId | 
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSalesOpportunitiesByParentIdForecastByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ForecastItem**](ForecastItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSalesOpportunitiesByParentIdForecastById

> ForecastItem PostSalesOpportunitiesByParentIdForecastById(ctx, id, parentId).ClientId(clientId).ForecastItem(forecastItem).Execute()

Post ForecastItem

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
	id := int32(56) // int32 | forecastId
	parentId := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 
	forecastItem := *openapiclient.NewForecastItem(*openapiclient.NewOpportunityReference(), *openapiclient.NewOpportunityStatusReference(), "ForecastType_example") // ForecastItem | forecast

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpportunityForecastItemsAPI.PostSalesOpportunitiesByParentIdForecastById(context.Background(), id, parentId).ClientId(clientId).ForecastItem(forecastItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityForecastItemsAPI.PostSalesOpportunitiesByParentIdForecastById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSalesOpportunitiesByParentIdForecastById`: ForecastItem
	fmt.Fprintf(os.Stdout, "Response from `OpportunityForecastItemsAPI.PostSalesOpportunitiesByParentIdForecastById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | forecastId | 
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSalesOpportunitiesByParentIdForecastByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **forecastItem** | [**ForecastItem**](ForecastItem.md) | forecast | 

### Return type

[**ForecastItem**](ForecastItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSalesOpportunitiesByParentIdForecastById

> ForecastItem PutSalesOpportunitiesByParentIdForecastById(ctx, id, parentId).ClientId(clientId).ForecastItem(forecastItem).Execute()

Put ForecastItem

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
	id := int32(56) // int32 | forecastId
	parentId := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 
	forecastItem := *openapiclient.NewForecastItem(*openapiclient.NewOpportunityReference(), *openapiclient.NewOpportunityStatusReference(), "ForecastType_example") // ForecastItem | forecast

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpportunityForecastItemsAPI.PutSalesOpportunitiesByParentIdForecastById(context.Background(), id, parentId).ClientId(clientId).ForecastItem(forecastItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityForecastItemsAPI.PutSalesOpportunitiesByParentIdForecastById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSalesOpportunitiesByParentIdForecastById`: ForecastItem
	fmt.Fprintf(os.Stdout, "Response from `OpportunityForecastItemsAPI.PutSalesOpportunitiesByParentIdForecastById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | forecastId | 
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSalesOpportunitiesByParentIdForecastByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **forecastItem** | [**ForecastItem**](ForecastItem.md) | forecast | 

### Return type

[**ForecastItem**](ForecastItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

