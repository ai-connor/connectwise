# \ReportCardDetailsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemReportCardsByParentIdDetailsById**](ReportCardDetailsAPI.md#DeleteSystemReportCardsByParentIdDetailsById) | **Delete** /system/reportCards/{parentId}/details/{id} | Delete ReportCardDetail
[**GetSystemReportCardsByParentIdDetails**](ReportCardDetailsAPI.md#GetSystemReportCardsByParentIdDetails) | **Get** /system/reportCards/{parentId}/details | Get List of ReportCardDetail
[**GetSystemReportCardsByParentIdDetailsById**](ReportCardDetailsAPI.md#GetSystemReportCardsByParentIdDetailsById) | **Get** /system/reportCards/{parentId}/details/{id} | Get ReportCardDetail
[**GetSystemReportCardsByParentIdDetailsCount**](ReportCardDetailsAPI.md#GetSystemReportCardsByParentIdDetailsCount) | **Get** /system/reportCards/{parentId}/details/count | Get Count of ReportCardDetail
[**PatchSystemReportCardsByParentIdDetailsById**](ReportCardDetailsAPI.md#PatchSystemReportCardsByParentIdDetailsById) | **Patch** /system/reportCards/{parentId}/details/{id} | Patch ReportCardDetail
[**PostSystemReportCardsByParentIdDetails**](ReportCardDetailsAPI.md#PostSystemReportCardsByParentIdDetails) | **Post** /system/reportCards/{parentId}/details | Post ReportCardDetail
[**PutSystemReportCardsByParentIdDetailsById**](ReportCardDetailsAPI.md#PutSystemReportCardsByParentIdDetailsById) | **Put** /system/reportCards/{parentId}/details/{id} | Put ReportCardDetail



## DeleteSystemReportCardsByParentIdDetailsById

> DeleteSystemReportCardsByParentIdDetailsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ReportCardDetail

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
	parentId := int32(56) // int32 | reportCardId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportCardDetailsAPI.DeleteSystemReportCardsByParentIdDetailsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportCardDetailsAPI.DeleteSystemReportCardsByParentIdDetailsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | detailId | 
**parentId** | **int32** | reportCardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemReportCardsByParentIdDetailsByIdRequest struct via the builder pattern


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


## GetSystemReportCardsByParentIdDetails

> []ReportCardDetail GetSystemReportCardsByParentIdDetails(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ReportCardDetail

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
	parentId := int32(56) // int32 | reportCardId
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
	resp, r, err := apiClient.ReportCardDetailsAPI.GetSystemReportCardsByParentIdDetails(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportCardDetailsAPI.GetSystemReportCardsByParentIdDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemReportCardsByParentIdDetails`: []ReportCardDetail
	fmt.Fprintf(os.Stdout, "Response from `ReportCardDetailsAPI.GetSystemReportCardsByParentIdDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | reportCardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemReportCardsByParentIdDetailsRequest struct via the builder pattern


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

[**[]ReportCardDetail**](ReportCardDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemReportCardsByParentIdDetailsById

> ReportCardDetail GetSystemReportCardsByParentIdDetailsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ReportCardDetail

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
	parentId := int32(56) // int32 | reportCardId
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
	resp, r, err := apiClient.ReportCardDetailsAPI.GetSystemReportCardsByParentIdDetailsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportCardDetailsAPI.GetSystemReportCardsByParentIdDetailsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemReportCardsByParentIdDetailsById`: ReportCardDetail
	fmt.Fprintf(os.Stdout, "Response from `ReportCardDetailsAPI.GetSystemReportCardsByParentIdDetailsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | detailId | 
**parentId** | **int32** | reportCardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemReportCardsByParentIdDetailsByIdRequest struct via the builder pattern


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

[**ReportCardDetail**](ReportCardDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemReportCardsByParentIdDetailsCount

> Count GetSystemReportCardsByParentIdDetailsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ReportCardDetail

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
	parentId := int32(56) // int32 | reportCardId
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
	resp, r, err := apiClient.ReportCardDetailsAPI.GetSystemReportCardsByParentIdDetailsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportCardDetailsAPI.GetSystemReportCardsByParentIdDetailsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemReportCardsByParentIdDetailsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ReportCardDetailsAPI.GetSystemReportCardsByParentIdDetailsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | reportCardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemReportCardsByParentIdDetailsCountRequest struct via the builder pattern


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


## PatchSystemReportCardsByParentIdDetailsById

> ReportCardDetail PatchSystemReportCardsByParentIdDetailsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ReportCardDetail

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
	parentId := int32(56) // int32 | reportCardId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportCardDetailsAPI.PatchSystemReportCardsByParentIdDetailsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportCardDetailsAPI.PatchSystemReportCardsByParentIdDetailsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemReportCardsByParentIdDetailsById`: ReportCardDetail
	fmt.Fprintf(os.Stdout, "Response from `ReportCardDetailsAPI.PatchSystemReportCardsByParentIdDetailsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | detailId | 
**parentId** | **int32** | reportCardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemReportCardsByParentIdDetailsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ReportCardDetail**](ReportCardDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemReportCardsByParentIdDetails

> ReportCardDetail PostSystemReportCardsByParentIdDetails(ctx, parentId).ClientId(clientId).ReportCardDetail(reportCardDetail).Execute()

Post ReportCardDetail

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
	parentId := int32(56) // int32 | reportCardId
	clientId := "clientId_example" // string | 
	reportCardDetail := *openapiclient.NewReportCardDetail(*openapiclient.NewKPIReference()) // ReportCardDetail | reportCardDetail

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportCardDetailsAPI.PostSystemReportCardsByParentIdDetails(context.Background(), parentId).ClientId(clientId).ReportCardDetail(reportCardDetail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportCardDetailsAPI.PostSystemReportCardsByParentIdDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemReportCardsByParentIdDetails`: ReportCardDetail
	fmt.Fprintf(os.Stdout, "Response from `ReportCardDetailsAPI.PostSystemReportCardsByParentIdDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | reportCardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemReportCardsByParentIdDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **reportCardDetail** | [**ReportCardDetail**](ReportCardDetail.md) | reportCardDetail | 

### Return type

[**ReportCardDetail**](ReportCardDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemReportCardsByParentIdDetailsById

> ReportCardDetail PutSystemReportCardsByParentIdDetailsById(ctx, id, parentId).ClientId(clientId).ReportCardDetail(reportCardDetail).Execute()

Put ReportCardDetail

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
	parentId := int32(56) // int32 | reportCardId
	clientId := "clientId_example" // string | 
	reportCardDetail := *openapiclient.NewReportCardDetail(*openapiclient.NewKPIReference()) // ReportCardDetail | reportCardDetail

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportCardDetailsAPI.PutSystemReportCardsByParentIdDetailsById(context.Background(), id, parentId).ClientId(clientId).ReportCardDetail(reportCardDetail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportCardDetailsAPI.PutSystemReportCardsByParentIdDetailsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemReportCardsByParentIdDetailsById`: ReportCardDetail
	fmt.Fprintf(os.Stdout, "Response from `ReportCardDetailsAPI.PutSystemReportCardsByParentIdDetailsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | detailId | 
**parentId** | **int32** | reportCardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemReportCardsByParentIdDetailsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **reportCardDetail** | [**ReportCardDetail**](ReportCardDetail.md) | reportCardDetail | 

### Return type

[**ReportCardDetail**](ReportCardDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

