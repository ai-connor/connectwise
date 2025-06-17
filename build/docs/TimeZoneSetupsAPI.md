# \TimeZoneSetupsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemTimeZoneSetupsById**](TimeZoneSetupsAPI.md#DeleteSystemTimeZoneSetupsById) | **Delete** /system/timeZoneSetups/{id} | Delete TimeZoneSetup
[**GetSystemTimeZoneSetups**](TimeZoneSetupsAPI.md#GetSystemTimeZoneSetups) | **Get** /system/timeZoneSetups | Get List of TimeZoneSetup
[**GetSystemTimeZoneSetupsById**](TimeZoneSetupsAPI.md#GetSystemTimeZoneSetupsById) | **Get** /system/timeZoneSetups/{id} | Get TimeZoneSetup
[**GetSystemTimeZoneSetupsCount**](TimeZoneSetupsAPI.md#GetSystemTimeZoneSetupsCount) | **Get** /system/timeZoneSetups/count | Get Count of TimeZoneSetup
[**PatchSystemTimeZoneSetupsById**](TimeZoneSetupsAPI.md#PatchSystemTimeZoneSetupsById) | **Patch** /system/timeZoneSetups/{id} | Patch TimeZoneSetup
[**PostSystemTimeZoneSetups**](TimeZoneSetupsAPI.md#PostSystemTimeZoneSetups) | **Post** /system/timeZoneSetups | Post TimeZoneSetup
[**PutSystemTimeZoneSetupsById**](TimeZoneSetupsAPI.md#PutSystemTimeZoneSetupsById) | **Put** /system/timeZoneSetups/{id} | Put TimeZoneSetup



## DeleteSystemTimeZoneSetupsById

> DeleteSystemTimeZoneSetupsById(ctx, id).ClientId(clientId).Execute()

Delete TimeZoneSetup

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
	id := int32(56) // int32 | timeZoneSetupId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TimeZoneSetupsAPI.DeleteSystemTimeZoneSetupsById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeZoneSetupsAPI.DeleteSystemTimeZoneSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | timeZoneSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemTimeZoneSetupsByIdRequest struct via the builder pattern


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


## GetSystemTimeZoneSetups

> []TimeZoneSetup GetSystemTimeZoneSetups(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of TimeZoneSetup

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
	resp, r, err := apiClient.TimeZoneSetupsAPI.GetSystemTimeZoneSetups(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeZoneSetupsAPI.GetSystemTimeZoneSetups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemTimeZoneSetups`: []TimeZoneSetup
	fmt.Fprintf(os.Stdout, "Response from `TimeZoneSetupsAPI.GetSystemTimeZoneSetups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemTimeZoneSetupsRequest struct via the builder pattern


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

[**[]TimeZoneSetup**](TimeZoneSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemTimeZoneSetupsById

> TimeZoneSetup GetSystemTimeZoneSetupsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get TimeZoneSetup

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
	id := int32(56) // int32 | timeZoneSetupId
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
	resp, r, err := apiClient.TimeZoneSetupsAPI.GetSystemTimeZoneSetupsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeZoneSetupsAPI.GetSystemTimeZoneSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemTimeZoneSetupsById`: TimeZoneSetup
	fmt.Fprintf(os.Stdout, "Response from `TimeZoneSetupsAPI.GetSystemTimeZoneSetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | timeZoneSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemTimeZoneSetupsByIdRequest struct via the builder pattern


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

[**TimeZoneSetup**](TimeZoneSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemTimeZoneSetupsCount

> Count GetSystemTimeZoneSetupsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of TimeZoneSetup

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
	resp, r, err := apiClient.TimeZoneSetupsAPI.GetSystemTimeZoneSetupsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeZoneSetupsAPI.GetSystemTimeZoneSetupsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemTimeZoneSetupsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `TimeZoneSetupsAPI.GetSystemTimeZoneSetupsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemTimeZoneSetupsCountRequest struct via the builder pattern


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


## PatchSystemTimeZoneSetupsById

> TimeZoneSetup PatchSystemTimeZoneSetupsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch TimeZoneSetup

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
	id := int32(56) // int32 | timeZoneSetupId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimeZoneSetupsAPI.PatchSystemTimeZoneSetupsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeZoneSetupsAPI.PatchSystemTimeZoneSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemTimeZoneSetupsById`: TimeZoneSetup
	fmt.Fprintf(os.Stdout, "Response from `TimeZoneSetupsAPI.PatchSystemTimeZoneSetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | timeZoneSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemTimeZoneSetupsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**TimeZoneSetup**](TimeZoneSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemTimeZoneSetups

> TimeZoneSetup PostSystemTimeZoneSetups(ctx).ClientId(clientId).TimeZoneSetup(timeZoneSetup).Execute()

Post TimeZoneSetup

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
	timeZoneSetup := *openapiclient.NewTimeZoneSetup("Name_example", *openapiclient.NewTimeZoneReference()) // TimeZoneSetup | timeZoneSetup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimeZoneSetupsAPI.PostSystemTimeZoneSetups(context.Background()).ClientId(clientId).TimeZoneSetup(timeZoneSetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeZoneSetupsAPI.PostSystemTimeZoneSetups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemTimeZoneSetups`: TimeZoneSetup
	fmt.Fprintf(os.Stdout, "Response from `TimeZoneSetupsAPI.PostSystemTimeZoneSetups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemTimeZoneSetupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **timeZoneSetup** | [**TimeZoneSetup**](TimeZoneSetup.md) | timeZoneSetup | 

### Return type

[**TimeZoneSetup**](TimeZoneSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemTimeZoneSetupsById

> TimeZoneSetup PutSystemTimeZoneSetupsById(ctx, id).ClientId(clientId).TimeZoneSetup(timeZoneSetup).Execute()

Put TimeZoneSetup

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
	id := int32(56) // int32 | timeZoneSetupId
	clientId := "clientId_example" // string | 
	timeZoneSetup := *openapiclient.NewTimeZoneSetup("Name_example", *openapiclient.NewTimeZoneReference()) // TimeZoneSetup | timeZoneSetup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimeZoneSetupsAPI.PutSystemTimeZoneSetupsById(context.Background(), id).ClientId(clientId).TimeZoneSetup(timeZoneSetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeZoneSetupsAPI.PutSystemTimeZoneSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemTimeZoneSetupsById`: TimeZoneSetup
	fmt.Fprintf(os.Stdout, "Response from `TimeZoneSetupsAPI.PutSystemTimeZoneSetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | timeZoneSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemTimeZoneSetupsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **timeZoneSetup** | [**TimeZoneSetup**](TimeZoneSetup.md) | timeZoneSetup | 

### Return type

[**TimeZoneSetup**](TimeZoneSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

