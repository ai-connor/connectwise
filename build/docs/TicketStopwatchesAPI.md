# \TicketStopwatchesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTimeTicketstopwatchesById**](TicketStopwatchesAPI.md#DeleteTimeTicketstopwatchesById) | **Delete** /time/ticketstopwatches/{id} | Delete TicketStopwatch
[**GetTimeTicketstopwatches**](TicketStopwatchesAPI.md#GetTimeTicketstopwatches) | **Get** /time/ticketstopwatches | Get List of TicketStopwatch
[**GetTimeTicketstopwatchesById**](TicketStopwatchesAPI.md#GetTimeTicketstopwatchesById) | **Get** /time/ticketstopwatches/{id} | Get TicketStopwatch
[**GetTimeTicketstopwatchesCount**](TicketStopwatchesAPI.md#GetTimeTicketstopwatchesCount) | **Get** /time/ticketstopwatches/count | Get Count of TicketStopwatch
[**PatchTimeTicketstopwatchesById**](TicketStopwatchesAPI.md#PatchTimeTicketstopwatchesById) | **Patch** /time/ticketstopwatches/{id} | Patch TicketStopwatch
[**PostTimeTicketstopwatches**](TicketStopwatchesAPI.md#PostTimeTicketstopwatches) | **Post** /time/ticketstopwatches | Post TicketStopwatch
[**PutTimeTicketstopwatchesById**](TicketStopwatchesAPI.md#PutTimeTicketstopwatchesById) | **Put** /time/ticketstopwatches/{id} | Put TicketStopwatch



## DeleteTimeTicketstopwatchesById

> DeleteTimeTicketstopwatchesById(ctx, id).ClientId(clientId).Execute()

Delete TicketStopwatch

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
	id := int32(56) // int32 | ticketstopwatcheId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TicketStopwatchesAPI.DeleteTimeTicketstopwatchesById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketStopwatchesAPI.DeleteTimeTicketstopwatchesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ticketstopwatcheId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTimeTicketstopwatchesByIdRequest struct via the builder pattern


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


## GetTimeTicketstopwatches

> []TicketStopwatch GetTimeTicketstopwatches(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of TicketStopwatch

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
	resp, r, err := apiClient.TicketStopwatchesAPI.GetTimeTicketstopwatches(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketStopwatchesAPI.GetTimeTicketstopwatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeTicketstopwatches`: []TicketStopwatch
	fmt.Fprintf(os.Stdout, "Response from `TicketStopwatchesAPI.GetTimeTicketstopwatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeTicketstopwatchesRequest struct via the builder pattern


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

[**[]TicketStopwatch**](TicketStopwatch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeTicketstopwatchesById

> TicketStopwatch GetTimeTicketstopwatchesById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get TicketStopwatch

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
	id := int32(56) // int32 | ticketstopwatcheId
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
	resp, r, err := apiClient.TicketStopwatchesAPI.GetTimeTicketstopwatchesById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketStopwatchesAPI.GetTimeTicketstopwatchesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeTicketstopwatchesById`: TicketStopwatch
	fmt.Fprintf(os.Stdout, "Response from `TicketStopwatchesAPI.GetTimeTicketstopwatchesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ticketstopwatcheId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeTicketstopwatchesByIdRequest struct via the builder pattern


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

[**TicketStopwatch**](TicketStopwatch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeTicketstopwatchesCount

> Count GetTimeTicketstopwatchesCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of TicketStopwatch

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
	resp, r, err := apiClient.TicketStopwatchesAPI.GetTimeTicketstopwatchesCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketStopwatchesAPI.GetTimeTicketstopwatchesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeTicketstopwatchesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `TicketStopwatchesAPI.GetTimeTicketstopwatchesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeTicketstopwatchesCountRequest struct via the builder pattern


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


## PatchTimeTicketstopwatchesById

> TicketStopwatch PatchTimeTicketstopwatchesById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch TicketStopwatch

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
	id := int32(56) // int32 | ticketstopwatcheId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketStopwatchesAPI.PatchTimeTicketstopwatchesById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketStopwatchesAPI.PatchTimeTicketstopwatchesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTimeTicketstopwatchesById`: TicketStopwatch
	fmt.Fprintf(os.Stdout, "Response from `TicketStopwatchesAPI.PatchTimeTicketstopwatchesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ticketstopwatcheId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTimeTicketstopwatchesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**TicketStopwatch**](TicketStopwatch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTimeTicketstopwatches

> TicketStopwatch PostTimeTicketstopwatches(ctx).ClientId(clientId).TicketStopwatch(ticketStopwatch).Execute()

Post TicketStopwatch

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
	ticketStopwatch := *openapiclient.NewTicketStopwatch(*openapiclient.NewMemberReference(), "Status_example", *openapiclient.NewTicketReference()) // TicketStopwatch | ticketStopwatch

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketStopwatchesAPI.PostTimeTicketstopwatches(context.Background()).ClientId(clientId).TicketStopwatch(ticketStopwatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketStopwatchesAPI.PostTimeTicketstopwatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTimeTicketstopwatches`: TicketStopwatch
	fmt.Fprintf(os.Stdout, "Response from `TicketStopwatchesAPI.PostTimeTicketstopwatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTimeTicketstopwatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **ticketStopwatch** | [**TicketStopwatch**](TicketStopwatch.md) | ticketStopwatch | 

### Return type

[**TicketStopwatch**](TicketStopwatch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTimeTicketstopwatchesById

> TicketStopwatch PutTimeTicketstopwatchesById(ctx, id).ClientId(clientId).TicketStopwatch(ticketStopwatch).Execute()

Put TicketStopwatch

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
	id := int32(56) // int32 | ticketstopwatcheId
	clientId := "clientId_example" // string | 
	ticketStopwatch := *openapiclient.NewTicketStopwatch(*openapiclient.NewMemberReference(), "Status_example", *openapiclient.NewTicketReference()) // TicketStopwatch | ticketStopwatch

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketStopwatchesAPI.PutTimeTicketstopwatchesById(context.Background(), id).ClientId(clientId).TicketStopwatch(ticketStopwatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketStopwatchesAPI.PutTimeTicketstopwatchesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTimeTicketstopwatchesById`: TicketStopwatch
	fmt.Fprintf(os.Stdout, "Response from `TicketStopwatchesAPI.PutTimeTicketstopwatchesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ticketstopwatcheId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTimeTicketstopwatchesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **ticketStopwatch** | [**TicketStopwatch**](TicketStopwatch.md) | ticketStopwatch | 

### Return type

[**TicketStopwatch**](TicketStopwatch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

