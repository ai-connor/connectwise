# \BillingSetupRoutingsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceBillingSetupsByParentIdRoutingsById**](BillingSetupRoutingsAPI.md#DeleteFinanceBillingSetupsByParentIdRoutingsById) | **Delete** /finance/billingSetups/{parentId}/routings/{id} | Delete BillingSetupRouting
[**GetFinanceBillingSetupsByParentIdRoutings**](BillingSetupRoutingsAPI.md#GetFinanceBillingSetupsByParentIdRoutings) | **Get** /finance/billingSetups/{parentId}/routings | Get List of BillingSetupRouting
[**GetFinanceBillingSetupsByParentIdRoutingsById**](BillingSetupRoutingsAPI.md#GetFinanceBillingSetupsByParentIdRoutingsById) | **Get** /finance/billingSetups/{parentId}/routings/{id} | Get BillingSetupRouting
[**GetFinanceBillingSetupsByParentIdRoutingsCount**](BillingSetupRoutingsAPI.md#GetFinanceBillingSetupsByParentIdRoutingsCount) | **Get** /finance/billingSetups/{parentId}/routings/count | Get Count of BillingSetupRouting
[**PatchFinanceBillingSetupsByParentIdRoutingsById**](BillingSetupRoutingsAPI.md#PatchFinanceBillingSetupsByParentIdRoutingsById) | **Patch** /finance/billingSetups/{parentId}/routings/{id} | Patch BillingSetupRouting
[**PostFinanceBillingSetupsByParentIdRoutings**](BillingSetupRoutingsAPI.md#PostFinanceBillingSetupsByParentIdRoutings) | **Post** /finance/billingSetups/{parentId}/routings | Post BillingSetupRouting
[**PutFinanceBillingSetupsByParentIdRoutingsById**](BillingSetupRoutingsAPI.md#PutFinanceBillingSetupsByParentIdRoutingsById) | **Put** /finance/billingSetups/{parentId}/routings/{id} | Put BillingSetupRouting



## DeleteFinanceBillingSetupsByParentIdRoutingsById

> DeleteFinanceBillingSetupsByParentIdRoutingsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete BillingSetupRouting

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
	id := int32(56) // int32 | routingId
	parentId := int32(56) // int32 | billingSetupId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BillingSetupRoutingsAPI.DeleteFinanceBillingSetupsByParentIdRoutingsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingSetupRoutingsAPI.DeleteFinanceBillingSetupsByParentIdRoutingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | routingId | 
**parentId** | **int32** | billingSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceBillingSetupsByParentIdRoutingsByIdRequest struct via the builder pattern


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


## GetFinanceBillingSetupsByParentIdRoutings

> []BillingSetupRouting GetFinanceBillingSetupsByParentIdRoutings(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of BillingSetupRouting

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
	parentId := int32(56) // int32 | billingSetupId
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
	resp, r, err := apiClient.BillingSetupRoutingsAPI.GetFinanceBillingSetupsByParentIdRoutings(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingSetupRoutingsAPI.GetFinanceBillingSetupsByParentIdRoutings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceBillingSetupsByParentIdRoutings`: []BillingSetupRouting
	fmt.Fprintf(os.Stdout, "Response from `BillingSetupRoutingsAPI.GetFinanceBillingSetupsByParentIdRoutings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | billingSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceBillingSetupsByParentIdRoutingsRequest struct via the builder pattern


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

[**[]BillingSetupRouting**](BillingSetupRouting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceBillingSetupsByParentIdRoutingsById

> BillingSetupRouting GetFinanceBillingSetupsByParentIdRoutingsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get BillingSetupRouting

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
	id := int32(56) // int32 | routingId
	parentId := int32(56) // int32 | billingSetupId
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
	resp, r, err := apiClient.BillingSetupRoutingsAPI.GetFinanceBillingSetupsByParentIdRoutingsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingSetupRoutingsAPI.GetFinanceBillingSetupsByParentIdRoutingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceBillingSetupsByParentIdRoutingsById`: BillingSetupRouting
	fmt.Fprintf(os.Stdout, "Response from `BillingSetupRoutingsAPI.GetFinanceBillingSetupsByParentIdRoutingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | routingId | 
**parentId** | **int32** | billingSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceBillingSetupsByParentIdRoutingsByIdRequest struct via the builder pattern


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

[**BillingSetupRouting**](BillingSetupRouting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceBillingSetupsByParentIdRoutingsCount

> Count GetFinanceBillingSetupsByParentIdRoutingsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of BillingSetupRouting

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
	parentId := int32(56) // int32 | billingSetupId
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
	resp, r, err := apiClient.BillingSetupRoutingsAPI.GetFinanceBillingSetupsByParentIdRoutingsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingSetupRoutingsAPI.GetFinanceBillingSetupsByParentIdRoutingsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceBillingSetupsByParentIdRoutingsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `BillingSetupRoutingsAPI.GetFinanceBillingSetupsByParentIdRoutingsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | billingSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceBillingSetupsByParentIdRoutingsCountRequest struct via the builder pattern


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


## PatchFinanceBillingSetupsByParentIdRoutingsById

> BillingSetupRouting PatchFinanceBillingSetupsByParentIdRoutingsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch BillingSetupRouting

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
	id := int32(56) // int32 | routingId
	parentId := int32(56) // int32 | billingSetupId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingSetupRoutingsAPI.PatchFinanceBillingSetupsByParentIdRoutingsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingSetupRoutingsAPI.PatchFinanceBillingSetupsByParentIdRoutingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceBillingSetupsByParentIdRoutingsById`: BillingSetupRouting
	fmt.Fprintf(os.Stdout, "Response from `BillingSetupRoutingsAPI.PatchFinanceBillingSetupsByParentIdRoutingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | routingId | 
**parentId** | **int32** | billingSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceBillingSetupsByParentIdRoutingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**BillingSetupRouting**](BillingSetupRouting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceBillingSetupsByParentIdRoutings

> BillingSetupRouting PostFinanceBillingSetupsByParentIdRoutings(ctx, parentId).ClientId(clientId).BillingSetupRouting(billingSetupRouting).Execute()

Post BillingSetupRouting

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
	parentId := int32(56) // int32 | billingSetupId
	clientId := "clientId_example" // string | 
	billingSetupRouting := *openapiclient.NewBillingSetupRouting(NullableInt32(123), "InvoiceRule_example", "RoutingRule_example") // BillingSetupRouting | billingSetupRouting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingSetupRoutingsAPI.PostFinanceBillingSetupsByParentIdRoutings(context.Background(), parentId).ClientId(clientId).BillingSetupRouting(billingSetupRouting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingSetupRoutingsAPI.PostFinanceBillingSetupsByParentIdRoutings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceBillingSetupsByParentIdRoutings`: BillingSetupRouting
	fmt.Fprintf(os.Stdout, "Response from `BillingSetupRoutingsAPI.PostFinanceBillingSetupsByParentIdRoutings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | billingSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceBillingSetupsByParentIdRoutingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **billingSetupRouting** | [**BillingSetupRouting**](BillingSetupRouting.md) | billingSetupRouting | 

### Return type

[**BillingSetupRouting**](BillingSetupRouting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceBillingSetupsByParentIdRoutingsById

> BillingSetupRouting PutFinanceBillingSetupsByParentIdRoutingsById(ctx, id, parentId).ClientId(clientId).BillingSetupRouting(billingSetupRouting).Execute()

Put BillingSetupRouting

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
	id := int32(56) // int32 | routingId
	parentId := int32(56) // int32 | billingSetupId
	clientId := "clientId_example" // string | 
	billingSetupRouting := *openapiclient.NewBillingSetupRouting(NullableInt32(123), "InvoiceRule_example", "RoutingRule_example") // BillingSetupRouting | billingSetupRouting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingSetupRoutingsAPI.PutFinanceBillingSetupsByParentIdRoutingsById(context.Background(), id, parentId).ClientId(clientId).BillingSetupRouting(billingSetupRouting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingSetupRoutingsAPI.PutFinanceBillingSetupsByParentIdRoutingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceBillingSetupsByParentIdRoutingsById`: BillingSetupRouting
	fmt.Fprintf(os.Stdout, "Response from `BillingSetupRoutingsAPI.PutFinanceBillingSetupsByParentIdRoutingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | routingId | 
**parentId** | **int32** | billingSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceBillingSetupsByParentIdRoutingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **billingSetupRouting** | [**BillingSetupRouting**](BillingSetupRouting.md) | billingSetupRouting | 

### Return type

[**BillingSetupRouting**](BillingSetupRouting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

