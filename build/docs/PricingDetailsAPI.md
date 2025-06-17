# \PricingDetailsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProcurementPricingschedulesByParentIdDetailsById**](PricingDetailsAPI.md#DeleteProcurementPricingschedulesByParentIdDetailsById) | **Delete** /procurement/pricingschedules/{parentId}/details/{id} | Delete PricingDetail
[**GetProcurementPricingschedulesByParentIdDetails**](PricingDetailsAPI.md#GetProcurementPricingschedulesByParentIdDetails) | **Get** /procurement/pricingschedules/{parentId}/details | Get List of PricingDetail
[**GetProcurementPricingschedulesByParentIdDetailsById**](PricingDetailsAPI.md#GetProcurementPricingschedulesByParentIdDetailsById) | **Get** /procurement/pricingschedules/{parentId}/details/{id} | Get PricingDetail
[**GetProcurementPricingschedulesByParentIdDetailsCount**](PricingDetailsAPI.md#GetProcurementPricingschedulesByParentIdDetailsCount) | **Get** /procurement/pricingschedules/{parentId}/details/count | Get Count of PricingDetail
[**PatchProcurementPricingschedulesByParentIdDetailsById**](PricingDetailsAPI.md#PatchProcurementPricingschedulesByParentIdDetailsById) | **Patch** /procurement/pricingschedules/{parentId}/details/{id} | Patch PricingDetail
[**PostProcurementPricingschedulesByParentIdDetails**](PricingDetailsAPI.md#PostProcurementPricingschedulesByParentIdDetails) | **Post** /procurement/pricingschedules/{parentId}/details | Post PricingDetail
[**PutProcurementPricingschedulesByParentIdDetailsById**](PricingDetailsAPI.md#PutProcurementPricingschedulesByParentIdDetailsById) | **Put** /procurement/pricingschedules/{parentId}/details/{id} | Put PricingDetail



## DeleteProcurementPricingschedulesByParentIdDetailsById

> DeleteProcurementPricingschedulesByParentIdDetailsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete PricingDetail

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
	id := int32(56) // int32 | detailId
	parentId := int32(56) // int32 | pricingscheduleId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PricingDetailsAPI.DeleteProcurementPricingschedulesByParentIdDetailsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingDetailsAPI.DeleteProcurementPricingschedulesByParentIdDetailsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | detailId | 
**parentId** | **int32** | pricingscheduleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcurementPricingschedulesByParentIdDetailsByIdRequest struct via the builder pattern


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


## GetProcurementPricingschedulesByParentIdDetails

> []PricingDetail GetProcurementPricingschedulesByParentIdDetails(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of PricingDetail

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
	parentId := int32(56) // int32 | pricingscheduleId
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
	resp, r, err := apiClient.PricingDetailsAPI.GetProcurementPricingschedulesByParentIdDetails(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingDetailsAPI.GetProcurementPricingschedulesByParentIdDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementPricingschedulesByParentIdDetails`: []PricingDetail
	fmt.Fprintf(os.Stdout, "Response from `PricingDetailsAPI.GetProcurementPricingschedulesByParentIdDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | pricingscheduleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementPricingschedulesByParentIdDetailsRequest struct via the builder pattern


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

[**[]PricingDetail**](PricingDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementPricingschedulesByParentIdDetailsById

> PricingDetail GetProcurementPricingschedulesByParentIdDetailsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get PricingDetail

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
	id := int32(56) // int32 | detailId
	parentId := int32(56) // int32 | pricingscheduleId
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
	resp, r, err := apiClient.PricingDetailsAPI.GetProcurementPricingschedulesByParentIdDetailsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingDetailsAPI.GetProcurementPricingschedulesByParentIdDetailsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementPricingschedulesByParentIdDetailsById`: PricingDetail
	fmt.Fprintf(os.Stdout, "Response from `PricingDetailsAPI.GetProcurementPricingschedulesByParentIdDetailsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | detailId | 
**parentId** | **int32** | pricingscheduleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementPricingschedulesByParentIdDetailsByIdRequest struct via the builder pattern


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

[**PricingDetail**](PricingDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementPricingschedulesByParentIdDetailsCount

> Count GetProcurementPricingschedulesByParentIdDetailsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of PricingDetail

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
	parentId := int32(56) // int32 | pricingscheduleId
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
	resp, r, err := apiClient.PricingDetailsAPI.GetProcurementPricingschedulesByParentIdDetailsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingDetailsAPI.GetProcurementPricingschedulesByParentIdDetailsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementPricingschedulesByParentIdDetailsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `PricingDetailsAPI.GetProcurementPricingschedulesByParentIdDetailsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | pricingscheduleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementPricingschedulesByParentIdDetailsCountRequest struct via the builder pattern


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


## PatchProcurementPricingschedulesByParentIdDetailsById

> PricingDetail PatchProcurementPricingschedulesByParentIdDetailsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch PricingDetail

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
	id := int32(56) // int32 | detailId
	parentId := int32(56) // int32 | pricingscheduleId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingDetailsAPI.PatchProcurementPricingschedulesByParentIdDetailsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingDetailsAPI.PatchProcurementPricingschedulesByParentIdDetailsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProcurementPricingschedulesByParentIdDetailsById`: PricingDetail
	fmt.Fprintf(os.Stdout, "Response from `PricingDetailsAPI.PatchProcurementPricingschedulesByParentIdDetailsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | detailId | 
**parentId** | **int32** | pricingscheduleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProcurementPricingschedulesByParentIdDetailsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**PricingDetail**](PricingDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProcurementPricingschedulesByParentIdDetails

> PricingDetail PostProcurementPricingschedulesByParentIdDetails(ctx, parentId).ClientId(clientId).PricingDetail(pricingDetail).Execute()

Post PricingDetail

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ai-connor/connectwise"
)

func main() {
	parentId := int32(56) // int32 | pricingscheduleId
	clientId := "clientId_example" // string | 
	pricingDetail := *openapiclient.NewPricingDetail(time.Now()) // PricingDetail | pricingDetail

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingDetailsAPI.PostProcurementPricingschedulesByParentIdDetails(context.Background(), parentId).ClientId(clientId).PricingDetail(pricingDetail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingDetailsAPI.PostProcurementPricingschedulesByParentIdDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementPricingschedulesByParentIdDetails`: PricingDetail
	fmt.Fprintf(os.Stdout, "Response from `PricingDetailsAPI.PostProcurementPricingschedulesByParentIdDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | pricingscheduleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementPricingschedulesByParentIdDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **pricingDetail** | [**PricingDetail**](PricingDetail.md) | pricingDetail | 

### Return type

[**PricingDetail**](PricingDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProcurementPricingschedulesByParentIdDetailsById

> PricingDetail PutProcurementPricingschedulesByParentIdDetailsById(ctx, id, parentId).ClientId(clientId).PricingDetail(pricingDetail).Execute()

Put PricingDetail

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ai-connor/connectwise"
)

func main() {
	id := int32(56) // int32 | detailId
	parentId := int32(56) // int32 | pricingscheduleId
	clientId := "clientId_example" // string | 
	pricingDetail := *openapiclient.NewPricingDetail(time.Now()) // PricingDetail | pricingDetail

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingDetailsAPI.PutProcurementPricingschedulesByParentIdDetailsById(context.Background(), id, parentId).ClientId(clientId).PricingDetail(pricingDetail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingDetailsAPI.PutProcurementPricingschedulesByParentIdDetailsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProcurementPricingschedulesByParentIdDetailsById`: PricingDetail
	fmt.Fprintf(os.Stdout, "Response from `PricingDetailsAPI.PutProcurementPricingschedulesByParentIdDetailsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | detailId | 
**parentId** | **int32** | pricingscheduleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProcurementPricingschedulesByParentIdDetailsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **pricingDetail** | [**PricingDetail**](PricingDetail.md) | pricingDetail | 

### Return type

[**PricingDetail**](PricingDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

