# \AutoSyncTimesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemAutoSyncTimeById**](AutoSyncTimesAPI.md#DeleteSystemAutoSyncTimeById) | **Delete** /system/autoSyncTime/{id} | Delete AutoSyncTime
[**GetSystemAutoSyncTime**](AutoSyncTimesAPI.md#GetSystemAutoSyncTime) | **Get** /system/autoSyncTime | Get List of AutoSyncTime
[**GetSystemAutoSyncTimeById**](AutoSyncTimesAPI.md#GetSystemAutoSyncTimeById) | **Get** /system/autoSyncTime/{id} | Get AutoSyncTime
[**GetSystemAutoSyncTimeCount**](AutoSyncTimesAPI.md#GetSystemAutoSyncTimeCount) | **Get** /system/autoSyncTime/count | Get Count of AutoSyncTime
[**PatchSystemAutoSyncTimeById**](AutoSyncTimesAPI.md#PatchSystemAutoSyncTimeById) | **Patch** /system/autoSyncTime/{id} | Patch AutoSyncTime
[**PostSystemAutoSyncTime**](AutoSyncTimesAPI.md#PostSystemAutoSyncTime) | **Post** /system/autoSyncTime | Post AutoSyncTime
[**PutSystemAutoSyncTimeById**](AutoSyncTimesAPI.md#PutSystemAutoSyncTimeById) | **Put** /system/autoSyncTime/{id} | Put AutoSyncTime



## DeleteSystemAutoSyncTimeById

> DeleteSystemAutoSyncTimeById(ctx, id).ClientId(clientId).Execute()

Delete AutoSyncTime

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
	id := int32(56) // int32 | autoSyncTimeId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutoSyncTimesAPI.DeleteSystemAutoSyncTimeById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoSyncTimesAPI.DeleteSystemAutoSyncTimeById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | autoSyncTimeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemAutoSyncTimeByIdRequest struct via the builder pattern


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


## GetSystemAutoSyncTime

> []AutoSyncTime GetSystemAutoSyncTime(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of AutoSyncTime

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
	resp, r, err := apiClient.AutoSyncTimesAPI.GetSystemAutoSyncTime(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoSyncTimesAPI.GetSystemAutoSyncTime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemAutoSyncTime`: []AutoSyncTime
	fmt.Fprintf(os.Stdout, "Response from `AutoSyncTimesAPI.GetSystemAutoSyncTime`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemAutoSyncTimeRequest struct via the builder pattern


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

[**[]AutoSyncTime**](AutoSyncTime.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemAutoSyncTimeById

> AutoSyncTime GetSystemAutoSyncTimeById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get AutoSyncTime

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
	id := int32(56) // int32 | autoSyncTimeId
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
	resp, r, err := apiClient.AutoSyncTimesAPI.GetSystemAutoSyncTimeById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoSyncTimesAPI.GetSystemAutoSyncTimeById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemAutoSyncTimeById`: AutoSyncTime
	fmt.Fprintf(os.Stdout, "Response from `AutoSyncTimesAPI.GetSystemAutoSyncTimeById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | autoSyncTimeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemAutoSyncTimeByIdRequest struct via the builder pattern


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

[**AutoSyncTime**](AutoSyncTime.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemAutoSyncTimeCount

> Count GetSystemAutoSyncTimeCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of AutoSyncTime

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
	resp, r, err := apiClient.AutoSyncTimesAPI.GetSystemAutoSyncTimeCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoSyncTimesAPI.GetSystemAutoSyncTimeCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemAutoSyncTimeCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `AutoSyncTimesAPI.GetSystemAutoSyncTimeCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemAutoSyncTimeCountRequest struct via the builder pattern


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


## PatchSystemAutoSyncTimeById

> AutoSyncTime PatchSystemAutoSyncTimeById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch AutoSyncTime

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
	id := int32(56) // int32 | autoSyncTimeId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoSyncTimesAPI.PatchSystemAutoSyncTimeById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoSyncTimesAPI.PatchSystemAutoSyncTimeById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemAutoSyncTimeById`: AutoSyncTime
	fmt.Fprintf(os.Stdout, "Response from `AutoSyncTimesAPI.PatchSystemAutoSyncTimeById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | autoSyncTimeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemAutoSyncTimeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**AutoSyncTime**](AutoSyncTime.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemAutoSyncTime

> AutoSyncTime PostSystemAutoSyncTime(ctx).ClientId(clientId).AutoSyncTime(autoSyncTime).Execute()

Post AutoSyncTime

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
	autoSyncTime := *openapiclient.NewAutoSyncTime("SyncTime_example", *openapiclient.NewTimeZoneSetupReference()) // AutoSyncTime | autoSyncTime

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoSyncTimesAPI.PostSystemAutoSyncTime(context.Background()).ClientId(clientId).AutoSyncTime(autoSyncTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoSyncTimesAPI.PostSystemAutoSyncTime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemAutoSyncTime`: AutoSyncTime
	fmt.Fprintf(os.Stdout, "Response from `AutoSyncTimesAPI.PostSystemAutoSyncTime`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemAutoSyncTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **autoSyncTime** | [**AutoSyncTime**](AutoSyncTime.md) | autoSyncTime | 

### Return type

[**AutoSyncTime**](AutoSyncTime.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemAutoSyncTimeById

> AutoSyncTime PutSystemAutoSyncTimeById(ctx, id).ClientId(clientId).AutoSyncTime(autoSyncTime).Execute()

Put AutoSyncTime

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
	id := int32(56) // int32 | autoSyncTimeId
	clientId := "clientId_example" // string | 
	autoSyncTime := *openapiclient.NewAutoSyncTime("SyncTime_example", *openapiclient.NewTimeZoneSetupReference()) // AutoSyncTime | autoSyncTime

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoSyncTimesAPI.PutSystemAutoSyncTimeById(context.Background(), id).ClientId(clientId).AutoSyncTime(autoSyncTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoSyncTimesAPI.PutSystemAutoSyncTimeById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemAutoSyncTimeById`: AutoSyncTime
	fmt.Fprintf(os.Stdout, "Response from `AutoSyncTimesAPI.PutSystemAutoSyncTimeById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | autoSyncTimeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemAutoSyncTimeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **autoSyncTime** | [**AutoSyncTime**](AutoSyncTime.md) | autoSyncTime | 

### Return type

[**AutoSyncTime**](AutoSyncTime.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

