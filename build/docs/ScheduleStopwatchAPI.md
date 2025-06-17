# \ScheduleStopwatchAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTimeSchedulestopwatchesById**](ScheduleStopwatchAPI.md#DeleteTimeSchedulestopwatchesById) | **Delete** /time/schedulestopwatches/{id} | Delete ScheduleStopwatch
[**GetTimeSchedulestopwatches**](ScheduleStopwatchAPI.md#GetTimeSchedulestopwatches) | **Get** /time/schedulestopwatches | Get List of ScheduleStopwatch
[**GetTimeSchedulestopwatchesById**](ScheduleStopwatchAPI.md#GetTimeSchedulestopwatchesById) | **Get** /time/schedulestopwatches/{id} | Get ScheduleStopwatch
[**GetTimeSchedulestopwatchesCount**](ScheduleStopwatchAPI.md#GetTimeSchedulestopwatchesCount) | **Get** /time/schedulestopwatches/count | Get Count of ScheduleStopwatch
[**PatchTimeSchedulestopwatchesById**](ScheduleStopwatchAPI.md#PatchTimeSchedulestopwatchesById) | **Patch** /time/schedulestopwatches/{id} | Patch ScheduleStopwatch
[**PostTimeSchedulestopwatches**](ScheduleStopwatchAPI.md#PostTimeSchedulestopwatches) | **Post** /time/schedulestopwatches | Post ScheduleStopwatch
[**PutTimeSchedulestopwatchesById**](ScheduleStopwatchAPI.md#PutTimeSchedulestopwatchesById) | **Put** /time/schedulestopwatches/{id} | Put ScheduleStopwatch



## DeleteTimeSchedulestopwatchesById

> DeleteTimeSchedulestopwatchesById(ctx, id).ClientId(clientId).Execute()

Delete ScheduleStopwatch

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
	id := int32(56) // int32 | schedulestopwatcheId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScheduleStopwatchAPI.DeleteTimeSchedulestopwatchesById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleStopwatchAPI.DeleteTimeSchedulestopwatchesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | schedulestopwatcheId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTimeSchedulestopwatchesByIdRequest struct via the builder pattern


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


## GetTimeSchedulestopwatches

> []ScheduleStopwatch GetTimeSchedulestopwatches(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ScheduleStopwatch

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
	resp, r, err := apiClient.ScheduleStopwatchAPI.GetTimeSchedulestopwatches(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleStopwatchAPI.GetTimeSchedulestopwatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeSchedulestopwatches`: []ScheduleStopwatch
	fmt.Fprintf(os.Stdout, "Response from `ScheduleStopwatchAPI.GetTimeSchedulestopwatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeSchedulestopwatchesRequest struct via the builder pattern


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

[**[]ScheduleStopwatch**](ScheduleStopwatch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeSchedulestopwatchesById

> ScheduleStopwatch GetTimeSchedulestopwatchesById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ScheduleStopwatch

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
	id := int32(56) // int32 | schedulestopwatcheId
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
	resp, r, err := apiClient.ScheduleStopwatchAPI.GetTimeSchedulestopwatchesById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleStopwatchAPI.GetTimeSchedulestopwatchesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeSchedulestopwatchesById`: ScheduleStopwatch
	fmt.Fprintf(os.Stdout, "Response from `ScheduleStopwatchAPI.GetTimeSchedulestopwatchesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | schedulestopwatcheId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeSchedulestopwatchesByIdRequest struct via the builder pattern


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

[**ScheduleStopwatch**](ScheduleStopwatch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeSchedulestopwatchesCount

> Count GetTimeSchedulestopwatchesCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ScheduleStopwatch

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
	resp, r, err := apiClient.ScheduleStopwatchAPI.GetTimeSchedulestopwatchesCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleStopwatchAPI.GetTimeSchedulestopwatchesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeSchedulestopwatchesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ScheduleStopwatchAPI.GetTimeSchedulestopwatchesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeSchedulestopwatchesCountRequest struct via the builder pattern


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


## PatchTimeSchedulestopwatchesById

> ScheduleStopwatch PatchTimeSchedulestopwatchesById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ScheduleStopwatch

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
	id := int32(56) // int32 | schedulestopwatcheId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleStopwatchAPI.PatchTimeSchedulestopwatchesById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleStopwatchAPI.PatchTimeSchedulestopwatchesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTimeSchedulestopwatchesById`: ScheduleStopwatch
	fmt.Fprintf(os.Stdout, "Response from `ScheduleStopwatchAPI.PatchTimeSchedulestopwatchesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | schedulestopwatcheId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTimeSchedulestopwatchesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ScheduleStopwatch**](ScheduleStopwatch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTimeSchedulestopwatches

> ScheduleStopwatch PostTimeSchedulestopwatches(ctx).ClientId(clientId).ScheduleStopwatch(scheduleStopwatch).Execute()

Post ScheduleStopwatch

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
	scheduleStopwatch := *openapiclient.NewScheduleStopwatch(*openapiclient.NewMemberReference(), NullableInt32(123), "Status_example") // ScheduleStopwatch | scheduleStopwatch

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleStopwatchAPI.PostTimeSchedulestopwatches(context.Background()).ClientId(clientId).ScheduleStopwatch(scheduleStopwatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleStopwatchAPI.PostTimeSchedulestopwatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTimeSchedulestopwatches`: ScheduleStopwatch
	fmt.Fprintf(os.Stdout, "Response from `ScheduleStopwatchAPI.PostTimeSchedulestopwatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTimeSchedulestopwatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **scheduleStopwatch** | [**ScheduleStopwatch**](ScheduleStopwatch.md) | scheduleStopwatch | 

### Return type

[**ScheduleStopwatch**](ScheduleStopwatch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTimeSchedulestopwatchesById

> ScheduleStopwatch PutTimeSchedulestopwatchesById(ctx, id).ClientId(clientId).ScheduleStopwatch(scheduleStopwatch).Execute()

Put ScheduleStopwatch

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
	id := int32(56) // int32 | schedulestopwatcheId
	clientId := "clientId_example" // string | 
	scheduleStopwatch := *openapiclient.NewScheduleStopwatch(*openapiclient.NewMemberReference(), NullableInt32(123), "Status_example") // ScheduleStopwatch | scheduleStopwatch

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleStopwatchAPI.PutTimeSchedulestopwatchesById(context.Background(), id).ClientId(clientId).ScheduleStopwatch(scheduleStopwatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleStopwatchAPI.PutTimeSchedulestopwatchesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTimeSchedulestopwatchesById`: ScheduleStopwatch
	fmt.Fprintf(os.Stdout, "Response from `ScheduleStopwatchAPI.PutTimeSchedulestopwatchesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | schedulestopwatcheId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTimeSchedulestopwatchesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **scheduleStopwatch** | [**ScheduleStopwatch**](ScheduleStopwatch.md) | scheduleStopwatch | 

### Return type

[**ScheduleStopwatch**](ScheduleStopwatch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

