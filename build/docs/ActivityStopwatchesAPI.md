# \ActivityStopwatchesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTimeActivitystopwatchesById**](ActivityStopwatchesAPI.md#DeleteTimeActivitystopwatchesById) | **Delete** /time/activitystopwatches/{id} | Delete ActivityStopwatch
[**GetTimeActivitystopwatches**](ActivityStopwatchesAPI.md#GetTimeActivitystopwatches) | **Get** /time/activitystopwatches | Get List of ActivityStopwatch
[**GetTimeActivitystopwatchesById**](ActivityStopwatchesAPI.md#GetTimeActivitystopwatchesById) | **Get** /time/activitystopwatches/{id} | Get ActivityStopwatch
[**GetTimeActivitystopwatchesCount**](ActivityStopwatchesAPI.md#GetTimeActivitystopwatchesCount) | **Get** /time/activitystopwatches/count | Get Count of ActivityStopwatch
[**PatchTimeActivitystopwatchesById**](ActivityStopwatchesAPI.md#PatchTimeActivitystopwatchesById) | **Patch** /time/activitystopwatches/{id} | Patch ActivityStopwatch
[**PostTimeActivitystopwatches**](ActivityStopwatchesAPI.md#PostTimeActivitystopwatches) | **Post** /time/activitystopwatches | Post ActivityStopwatch
[**PutTimeActivitystopwatchesById**](ActivityStopwatchesAPI.md#PutTimeActivitystopwatchesById) | **Put** /time/activitystopwatches/{id} | Put ActivityStopwatch



## DeleteTimeActivitystopwatchesById

> DeleteTimeActivitystopwatchesById(ctx, id).ClientId(clientId).Execute()

Delete ActivityStopwatch

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
	id := int32(56) // int32 | activitystopwatcheId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ActivityStopwatchesAPI.DeleteTimeActivitystopwatchesById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityStopwatchesAPI.DeleteTimeActivitystopwatchesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | activitystopwatcheId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTimeActivitystopwatchesByIdRequest struct via the builder pattern


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


## GetTimeActivitystopwatches

> []ActivityStopwatch GetTimeActivitystopwatches(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ActivityStopwatch

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
	resp, r, err := apiClient.ActivityStopwatchesAPI.GetTimeActivitystopwatches(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityStopwatchesAPI.GetTimeActivitystopwatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeActivitystopwatches`: []ActivityStopwatch
	fmt.Fprintf(os.Stdout, "Response from `ActivityStopwatchesAPI.GetTimeActivitystopwatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeActivitystopwatchesRequest struct via the builder pattern


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

[**[]ActivityStopwatch**](ActivityStopwatch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeActivitystopwatchesById

> ActivityStopwatch GetTimeActivitystopwatchesById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ActivityStopwatch

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
	id := int32(56) // int32 | activitystopwatcheId
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
	resp, r, err := apiClient.ActivityStopwatchesAPI.GetTimeActivitystopwatchesById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityStopwatchesAPI.GetTimeActivitystopwatchesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeActivitystopwatchesById`: ActivityStopwatch
	fmt.Fprintf(os.Stdout, "Response from `ActivityStopwatchesAPI.GetTimeActivitystopwatchesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | activitystopwatcheId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeActivitystopwatchesByIdRequest struct via the builder pattern


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

[**ActivityStopwatch**](ActivityStopwatch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeActivitystopwatchesCount

> Count GetTimeActivitystopwatchesCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ActivityStopwatch

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
	resp, r, err := apiClient.ActivityStopwatchesAPI.GetTimeActivitystopwatchesCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityStopwatchesAPI.GetTimeActivitystopwatchesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeActivitystopwatchesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ActivityStopwatchesAPI.GetTimeActivitystopwatchesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeActivitystopwatchesCountRequest struct via the builder pattern


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


## PatchTimeActivitystopwatchesById

> ActivityStopwatch PatchTimeActivitystopwatchesById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ActivityStopwatch

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
	id := int32(56) // int32 | activitystopwatcheId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityStopwatchesAPI.PatchTimeActivitystopwatchesById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityStopwatchesAPI.PatchTimeActivitystopwatchesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTimeActivitystopwatchesById`: ActivityStopwatch
	fmt.Fprintf(os.Stdout, "Response from `ActivityStopwatchesAPI.PatchTimeActivitystopwatchesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | activitystopwatcheId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTimeActivitystopwatchesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ActivityStopwatch**](ActivityStopwatch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTimeActivitystopwatches

> ActivityStopwatch PostTimeActivitystopwatches(ctx).ClientId(clientId).ActivityStopwatch(activityStopwatch).Execute()

Post ActivityStopwatch

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
	activityStopwatch := *openapiclient.NewActivityStopwatch(NullableInt32(123), *openapiclient.NewMemberReference(), "Status_example") // ActivityStopwatch | activityStopwatch

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityStopwatchesAPI.PostTimeActivitystopwatches(context.Background()).ClientId(clientId).ActivityStopwatch(activityStopwatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityStopwatchesAPI.PostTimeActivitystopwatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTimeActivitystopwatches`: ActivityStopwatch
	fmt.Fprintf(os.Stdout, "Response from `ActivityStopwatchesAPI.PostTimeActivitystopwatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTimeActivitystopwatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **activityStopwatch** | [**ActivityStopwatch**](ActivityStopwatch.md) | activityStopwatch | 

### Return type

[**ActivityStopwatch**](ActivityStopwatch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTimeActivitystopwatchesById

> ActivityStopwatch PutTimeActivitystopwatchesById(ctx, id).ClientId(clientId).ActivityStopwatch(activityStopwatch).Execute()

Put ActivityStopwatch

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
	id := int32(56) // int32 | activitystopwatcheId
	clientId := "clientId_example" // string | 
	activityStopwatch := *openapiclient.NewActivityStopwatch(NullableInt32(123), *openapiclient.NewMemberReference(), "Status_example") // ActivityStopwatch | activityStopwatch

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityStopwatchesAPI.PutTimeActivitystopwatchesById(context.Background(), id).ClientId(clientId).ActivityStopwatch(activityStopwatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityStopwatchesAPI.PutTimeActivitystopwatchesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTimeActivitystopwatchesById`: ActivityStopwatch
	fmt.Fprintf(os.Stdout, "Response from `ActivityStopwatchesAPI.PutTimeActivitystopwatchesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | activitystopwatcheId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTimeActivitystopwatchesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **activityStopwatch** | [**ActivityStopwatch**](ActivityStopwatch.md) | activityStopwatch | 

### Return type

[**ActivityStopwatch**](ActivityStopwatch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

