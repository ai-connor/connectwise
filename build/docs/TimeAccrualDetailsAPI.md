# \TimeAccrualDetailsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTimeAccrualsByParentIdDetailsById**](TimeAccrualDetailsAPI.md#DeleteTimeAccrualsByParentIdDetailsById) | **Delete** /time/accruals/{parentId}/details/{id} | Delete TimeAccrualDetail
[**GetTimeAccrualsByParentIdDetails**](TimeAccrualDetailsAPI.md#GetTimeAccrualsByParentIdDetails) | **Get** /time/accruals/{parentId}/details | Get List of TimeAccrualDetail
[**GetTimeAccrualsByParentIdDetailsById**](TimeAccrualDetailsAPI.md#GetTimeAccrualsByParentIdDetailsById) | **Get** /time/accruals/{parentId}/details/{id} | Get TimeAccrualDetail
[**GetTimeAccrualsByParentIdDetailsCount**](TimeAccrualDetailsAPI.md#GetTimeAccrualsByParentIdDetailsCount) | **Get** /time/accruals/{parentId}/details/count | Get Count of TimeAccrualDetail
[**PatchTimeAccrualsByParentIdDetailsById**](TimeAccrualDetailsAPI.md#PatchTimeAccrualsByParentIdDetailsById) | **Patch** /time/accruals/{parentId}/details/{id} | Patch TimeAccrualDetail
[**PostTimeAccrualsByParentIdDetails**](TimeAccrualDetailsAPI.md#PostTimeAccrualsByParentIdDetails) | **Post** /time/accruals/{parentId}/details/ | Post TimeAccrualDetail
[**PutTimeAccrualsByParentIdDetailsById**](TimeAccrualDetailsAPI.md#PutTimeAccrualsByParentIdDetailsById) | **Put** /time/accruals/{parentId}/details/{id} | Put TimeAccrualDetail



## DeleteTimeAccrualsByParentIdDetailsById

> DeleteTimeAccrualsByParentIdDetailsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete TimeAccrualDetail

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
	parentId := int32(56) // int32 | accrualId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TimeAccrualDetailsAPI.DeleteTimeAccrualsByParentIdDetailsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeAccrualDetailsAPI.DeleteTimeAccrualsByParentIdDetailsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | detailId | 
**parentId** | **int32** | accrualId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTimeAccrualsByParentIdDetailsByIdRequest struct via the builder pattern


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


## GetTimeAccrualsByParentIdDetails

> []TimeAccrualDetail GetTimeAccrualsByParentIdDetails(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of TimeAccrualDetail

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
	parentId := int32(56) // int32 | accrualId
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
	resp, r, err := apiClient.TimeAccrualDetailsAPI.GetTimeAccrualsByParentIdDetails(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeAccrualDetailsAPI.GetTimeAccrualsByParentIdDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeAccrualsByParentIdDetails`: []TimeAccrualDetail
	fmt.Fprintf(os.Stdout, "Response from `TimeAccrualDetailsAPI.GetTimeAccrualsByParentIdDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | accrualId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeAccrualsByParentIdDetailsRequest struct via the builder pattern


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

[**[]TimeAccrualDetail**](TimeAccrualDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeAccrualsByParentIdDetailsById

> TimeAccrualDetail GetTimeAccrualsByParentIdDetailsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get TimeAccrualDetail

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
	parentId := int32(56) // int32 | accrualId
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
	resp, r, err := apiClient.TimeAccrualDetailsAPI.GetTimeAccrualsByParentIdDetailsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeAccrualDetailsAPI.GetTimeAccrualsByParentIdDetailsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeAccrualsByParentIdDetailsById`: TimeAccrualDetail
	fmt.Fprintf(os.Stdout, "Response from `TimeAccrualDetailsAPI.GetTimeAccrualsByParentIdDetailsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | detailId | 
**parentId** | **int32** | accrualId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeAccrualsByParentIdDetailsByIdRequest struct via the builder pattern


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

[**TimeAccrualDetail**](TimeAccrualDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeAccrualsByParentIdDetailsCount

> Count GetTimeAccrualsByParentIdDetailsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of TimeAccrualDetail

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
	parentId := int32(56) // int32 | accrualId
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
	resp, r, err := apiClient.TimeAccrualDetailsAPI.GetTimeAccrualsByParentIdDetailsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeAccrualDetailsAPI.GetTimeAccrualsByParentIdDetailsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeAccrualsByParentIdDetailsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `TimeAccrualDetailsAPI.GetTimeAccrualsByParentIdDetailsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | accrualId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeAccrualsByParentIdDetailsCountRequest struct via the builder pattern


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


## PatchTimeAccrualsByParentIdDetailsById

> TimeAccrualDetail PatchTimeAccrualsByParentIdDetailsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch TimeAccrualDetail

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
	parentId := int32(56) // int32 | accrualId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimeAccrualDetailsAPI.PatchTimeAccrualsByParentIdDetailsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeAccrualDetailsAPI.PatchTimeAccrualsByParentIdDetailsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTimeAccrualsByParentIdDetailsById`: TimeAccrualDetail
	fmt.Fprintf(os.Stdout, "Response from `TimeAccrualDetailsAPI.PatchTimeAccrualsByParentIdDetailsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | detailId | 
**parentId** | **int32** | accrualId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTimeAccrualsByParentIdDetailsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**TimeAccrualDetail**](TimeAccrualDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTimeAccrualsByParentIdDetails

> TimeAccrualDetail PostTimeAccrualsByParentIdDetails(ctx, parentId).ClientId(clientId).TimeAccrualDetail(timeAccrualDetail).Execute()

Post TimeAccrualDetail

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
	parentId := int32(56) // int32 | accrualId
	clientId := "clientId_example" // string | 
	timeAccrualDetail := *openapiclient.NewTimeAccrualDetail("AccrualType_example", NullableInt32(123), NullableInt32(123), NullableFloat64(123)) // TimeAccrualDetail | timeAccrualDetail

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimeAccrualDetailsAPI.PostTimeAccrualsByParentIdDetails(context.Background(), parentId).ClientId(clientId).TimeAccrualDetail(timeAccrualDetail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeAccrualDetailsAPI.PostTimeAccrualsByParentIdDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTimeAccrualsByParentIdDetails`: TimeAccrualDetail
	fmt.Fprintf(os.Stdout, "Response from `TimeAccrualDetailsAPI.PostTimeAccrualsByParentIdDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | accrualId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTimeAccrualsByParentIdDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **timeAccrualDetail** | [**TimeAccrualDetail**](TimeAccrualDetail.md) | timeAccrualDetail | 

### Return type

[**TimeAccrualDetail**](TimeAccrualDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTimeAccrualsByParentIdDetailsById

> TimeAccrualDetail PutTimeAccrualsByParentIdDetailsById(ctx, id, parentId).ClientId(clientId).TimeAccrualDetail(timeAccrualDetail).Execute()

Put TimeAccrualDetail

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
	parentId := int32(56) // int32 | accrualId
	clientId := "clientId_example" // string | 
	timeAccrualDetail := *openapiclient.NewTimeAccrualDetail("AccrualType_example", NullableInt32(123), NullableInt32(123), NullableFloat64(123)) // TimeAccrualDetail | timeAccrualDetail

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimeAccrualDetailsAPI.PutTimeAccrualsByParentIdDetailsById(context.Background(), id, parentId).ClientId(clientId).TimeAccrualDetail(timeAccrualDetail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeAccrualDetailsAPI.PutTimeAccrualsByParentIdDetailsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTimeAccrualsByParentIdDetailsById`: TimeAccrualDetail
	fmt.Fprintf(os.Stdout, "Response from `TimeAccrualDetailsAPI.PutTimeAccrualsByParentIdDetailsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | detailId | 
**parentId** | **int32** | accrualId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTimeAccrualsByParentIdDetailsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **timeAccrualDetail** | [**TimeAccrualDetail**](TimeAccrualDetail.md) | timeAccrualDetail | 

### Return type

[**TimeAccrualDetail**](TimeAccrualDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

