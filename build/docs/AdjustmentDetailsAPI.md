# \AdjustmentDetailsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProcurementAdjustmentsByParentIdDetailsById**](AdjustmentDetailsAPI.md#DeleteProcurementAdjustmentsByParentIdDetailsById) | **Delete** /procurement/adjustments/{parentId}/details/{id} | Delete AdjustmentDetail
[**GetProcurementAdjustmentsByParentIdDetails**](AdjustmentDetailsAPI.md#GetProcurementAdjustmentsByParentIdDetails) | **Get** /procurement/adjustments/{parentId}/details | Get List of AdjustmentDetail
[**GetProcurementAdjustmentsByParentIdDetailsById**](AdjustmentDetailsAPI.md#GetProcurementAdjustmentsByParentIdDetailsById) | **Get** /procurement/adjustments/{parentId}/details/{id} | Get AdjustmentDetail
[**GetProcurementAdjustmentsByParentIdDetailsCount**](AdjustmentDetailsAPI.md#GetProcurementAdjustmentsByParentIdDetailsCount) | **Get** /procurement/adjustments/{parentId}/details/count | Get Count of AdjustmentDetail
[**PostProcurementAdjustmentsByParentIdDetails**](AdjustmentDetailsAPI.md#PostProcurementAdjustmentsByParentIdDetails) | **Post** /procurement/adjustments/{parentId}/details | Post AdjustmentDetail



## DeleteProcurementAdjustmentsByParentIdDetailsById

> DeleteProcurementAdjustmentsByParentIdDetailsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete AdjustmentDetail

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
	parentId := int32(56) // int32 | adjustmentId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdjustmentDetailsAPI.DeleteProcurementAdjustmentsByParentIdDetailsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdjustmentDetailsAPI.DeleteProcurementAdjustmentsByParentIdDetailsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | detailId | 
**parentId** | **int32** | adjustmentId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcurementAdjustmentsByParentIdDetailsByIdRequest struct via the builder pattern


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


## GetProcurementAdjustmentsByParentIdDetails

> []AdjustmentDetail GetProcurementAdjustmentsByParentIdDetails(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of AdjustmentDetail

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
	parentId := int32(56) // int32 | adjustmentId
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
	resp, r, err := apiClient.AdjustmentDetailsAPI.GetProcurementAdjustmentsByParentIdDetails(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdjustmentDetailsAPI.GetProcurementAdjustmentsByParentIdDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementAdjustmentsByParentIdDetails`: []AdjustmentDetail
	fmt.Fprintf(os.Stdout, "Response from `AdjustmentDetailsAPI.GetProcurementAdjustmentsByParentIdDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | adjustmentId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementAdjustmentsByParentIdDetailsRequest struct via the builder pattern


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

[**[]AdjustmentDetail**](AdjustmentDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementAdjustmentsByParentIdDetailsById

> AdjustmentDetail GetProcurementAdjustmentsByParentIdDetailsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get AdjustmentDetail

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
	parentId := int32(56) // int32 | adjustmentId
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
	resp, r, err := apiClient.AdjustmentDetailsAPI.GetProcurementAdjustmentsByParentIdDetailsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdjustmentDetailsAPI.GetProcurementAdjustmentsByParentIdDetailsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementAdjustmentsByParentIdDetailsById`: AdjustmentDetail
	fmt.Fprintf(os.Stdout, "Response from `AdjustmentDetailsAPI.GetProcurementAdjustmentsByParentIdDetailsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | detailId | 
**parentId** | **int32** | adjustmentId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementAdjustmentsByParentIdDetailsByIdRequest struct via the builder pattern


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

[**AdjustmentDetail**](AdjustmentDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementAdjustmentsByParentIdDetailsCount

> Count GetProcurementAdjustmentsByParentIdDetailsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of AdjustmentDetail

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
	parentId := int32(56) // int32 | adjustmentId
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
	resp, r, err := apiClient.AdjustmentDetailsAPI.GetProcurementAdjustmentsByParentIdDetailsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdjustmentDetailsAPI.GetProcurementAdjustmentsByParentIdDetailsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementAdjustmentsByParentIdDetailsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `AdjustmentDetailsAPI.GetProcurementAdjustmentsByParentIdDetailsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | adjustmentId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementAdjustmentsByParentIdDetailsCountRequest struct via the builder pattern


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


## PostProcurementAdjustmentsByParentIdDetails

> AdjustmentDetail PostProcurementAdjustmentsByParentIdDetails(ctx, parentId).ClientId(clientId).AdjustmentDetail(adjustmentDetail).Execute()

Post AdjustmentDetail

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
	parentId := int32(56) // int32 | adjustmentId
	clientId := "clientId_example" // string | 
	adjustmentDetail := *openapiclient.NewAdjustmentDetail(*openapiclient.NewCatalogItemReference(), *openapiclient.NewWarehouseReference(), *openapiclient.NewWarehouseBinReference(), NullableInt32(123)) // AdjustmentDetail | adjustmentDetail

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdjustmentDetailsAPI.PostProcurementAdjustmentsByParentIdDetails(context.Background(), parentId).ClientId(clientId).AdjustmentDetail(adjustmentDetail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdjustmentDetailsAPI.PostProcurementAdjustmentsByParentIdDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementAdjustmentsByParentIdDetails`: AdjustmentDetail
	fmt.Fprintf(os.Stdout, "Response from `AdjustmentDetailsAPI.PostProcurementAdjustmentsByParentIdDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | adjustmentId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementAdjustmentsByParentIdDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **adjustmentDetail** | [**AdjustmentDetail**](AdjustmentDetail.md) | adjustmentDetail | 

### Return type

[**AdjustmentDetail**](AdjustmentDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

