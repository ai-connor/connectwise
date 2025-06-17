# \OpportunityForecastsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSalesOpportunitiesByParentIdForecast**](OpportunityForecastsAPI.md#DeleteSalesOpportunitiesByParentIdForecast) | **Delete** /sales/opportunities/{parentId}/forecast/ | Delete Forecast
[**GetSalesOpportunitiesByParentIdForecast**](OpportunityForecastsAPI.md#GetSalesOpportunitiesByParentIdForecast) | **Get** /sales/opportunities/{parentId}/forecast | Get List of Forecast
[**GetSalesOpportunitiesByParentIdForecastCount**](OpportunityForecastsAPI.md#GetSalesOpportunitiesByParentIdForecastCount) | **Get** /sales/opportunities/{parentId}/forecast/count | Get Count of Forecast
[**PatchSalesOpportunitiesByParentIdForecast**](OpportunityForecastsAPI.md#PatchSalesOpportunitiesByParentIdForecast) | **Patch** /sales/opportunities/{parentId}/forecast/ | Patch Forecast
[**PostSalesOpportunitiesByParentIdForecast**](OpportunityForecastsAPI.md#PostSalesOpportunitiesByParentIdForecast) | **Post** /sales/opportunities/{parentId}/forecast | Post Forecast
[**PostSalesOpportunitiesByParentIdForecastCopyById**](OpportunityForecastsAPI.md#PostSalesOpportunitiesByParentIdForecastCopyById) | **Post** /sales/opportunities/{parentId}/forecast/copy/{id} | Post SuccessResponse
[**PutSalesOpportunitiesByParentIdForecast**](OpportunityForecastsAPI.md#PutSalesOpportunitiesByParentIdForecast) | **Put** /sales/opportunities/{parentId}/forecast/ | Put Forecast



## DeleteSalesOpportunitiesByParentIdForecast

> DeleteSalesOpportunitiesByParentIdForecast(ctx, parentId).ClientId(clientId).Execute()

Delete Forecast

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
	parentId := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OpportunityForecastsAPI.DeleteSalesOpportunitiesByParentIdForecast(context.Background(), parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityForecastsAPI.DeleteSalesOpportunitiesByParentIdForecast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSalesOpportunitiesByParentIdForecastRequest struct via the builder pattern


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


## GetSalesOpportunitiesByParentIdForecast

> []Forecast GetSalesOpportunitiesByParentIdForecast(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of Forecast

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
	resp, r, err := apiClient.OpportunityForecastsAPI.GetSalesOpportunitiesByParentIdForecast(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityForecastsAPI.GetSalesOpportunitiesByParentIdForecast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesOpportunitiesByParentIdForecast`: []Forecast
	fmt.Fprintf(os.Stdout, "Response from `OpportunityForecastsAPI.GetSalesOpportunitiesByParentIdForecast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesOpportunitiesByParentIdForecastRequest struct via the builder pattern


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

[**[]Forecast**](Forecast.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalesOpportunitiesByParentIdForecastCount

> Count GetSalesOpportunitiesByParentIdForecastCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of Forecast

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
	resp, r, err := apiClient.OpportunityForecastsAPI.GetSalesOpportunitiesByParentIdForecastCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityForecastsAPI.GetSalesOpportunitiesByParentIdForecastCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesOpportunitiesByParentIdForecastCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `OpportunityForecastsAPI.GetSalesOpportunitiesByParentIdForecastCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesOpportunitiesByParentIdForecastCountRequest struct via the builder pattern


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


## PatchSalesOpportunitiesByParentIdForecast

> Forecast PatchSalesOpportunitiesByParentIdForecast(ctx, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch Forecast

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
	parentId := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpportunityForecastsAPI.PatchSalesOpportunitiesByParentIdForecast(context.Background(), parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityForecastsAPI.PatchSalesOpportunitiesByParentIdForecast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSalesOpportunitiesByParentIdForecast`: Forecast
	fmt.Fprintf(os.Stdout, "Response from `OpportunityForecastsAPI.PatchSalesOpportunitiesByParentIdForecast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSalesOpportunitiesByParentIdForecastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**Forecast**](Forecast.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSalesOpportunitiesByParentIdForecast

> Forecast PostSalesOpportunitiesByParentIdForecast(ctx, parentId).ClientId(clientId).Forecast(forecast).Execute()

Post Forecast

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
	parentId := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 
	forecast := *openapiclient.NewForecast() // Forecast | forecast

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpportunityForecastsAPI.PostSalesOpportunitiesByParentIdForecast(context.Background(), parentId).ClientId(clientId).Forecast(forecast).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityForecastsAPI.PostSalesOpportunitiesByParentIdForecast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSalesOpportunitiesByParentIdForecast`: Forecast
	fmt.Fprintf(os.Stdout, "Response from `OpportunityForecastsAPI.PostSalesOpportunitiesByParentIdForecast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSalesOpportunitiesByParentIdForecastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **forecast** | [**Forecast**](Forecast.md) | forecast | 

### Return type

[**Forecast**](Forecast.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSalesOpportunitiesByParentIdForecastCopyById

> SuccessResponse PostSalesOpportunitiesByParentIdForecastCopyById(ctx, id, parentId).ClientId(clientId).Execute()

Post SuccessResponse

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
	id := int32(56) // int32 | copyId
	parentId := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpportunityForecastsAPI.PostSalesOpportunitiesByParentIdForecastCopyById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityForecastsAPI.PostSalesOpportunitiesByParentIdForecastCopyById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSalesOpportunitiesByParentIdForecastCopyById`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `OpportunityForecastsAPI.PostSalesOpportunitiesByParentIdForecastCopyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | copyId | 
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSalesOpportunitiesByParentIdForecastCopyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSalesOpportunitiesByParentIdForecast

> Forecast PutSalesOpportunitiesByParentIdForecast(ctx, parentId).ClientId(clientId).Forecast(forecast).Execute()

Put Forecast

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
	parentId := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 
	forecast := *openapiclient.NewForecast() // Forecast | forecast

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpportunityForecastsAPI.PutSalesOpportunitiesByParentIdForecast(context.Background(), parentId).ClientId(clientId).Forecast(forecast).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityForecastsAPI.PutSalesOpportunitiesByParentIdForecast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSalesOpportunitiesByParentIdForecast`: Forecast
	fmt.Fprintf(os.Stdout, "Response from `OpportunityForecastsAPI.PutSalesOpportunitiesByParentIdForecast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSalesOpportunitiesByParentIdForecastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **forecast** | [**Forecast**](Forecast.md) | forecast | 

### Return type

[**Forecast**](Forecast.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

