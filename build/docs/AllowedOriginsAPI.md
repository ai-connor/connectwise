# \AllowedOriginsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemAllowedoriginsById**](AllowedOriginsAPI.md#DeleteSystemAllowedoriginsById) | **Delete** /system/allowedorigins/{id} | Delete AllowedOrigin
[**GetSystemAllowedorigins**](AllowedOriginsAPI.md#GetSystemAllowedorigins) | **Get** /system/allowedorigins | Get List of AllowedOrigin
[**GetSystemAllowedoriginsById**](AllowedOriginsAPI.md#GetSystemAllowedoriginsById) | **Get** /system/allowedorigins/{id} | Get AllowedOrigin
[**GetSystemAllowedoriginsCount**](AllowedOriginsAPI.md#GetSystemAllowedoriginsCount) | **Get** /system/allowedorigins/count | Get Count of AllowedOrigin
[**PatchSystemAllowedoriginsById**](AllowedOriginsAPI.md#PatchSystemAllowedoriginsById) | **Patch** /system/allowedorigins/{id} | Patch AllowedOrigin
[**PostSystemAllowedorigins**](AllowedOriginsAPI.md#PostSystemAllowedorigins) | **Post** /system/allowedorigins | Post AllowedOrigin
[**PutSystemAllowedoriginsById**](AllowedOriginsAPI.md#PutSystemAllowedoriginsById) | **Put** /system/allowedorigins/{id} | Put AllowedOrigin



## DeleteSystemAllowedoriginsById

> DeleteSystemAllowedoriginsById(ctx, id).ClientId(clientId).Execute()

Delete AllowedOrigin

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
	id := int32(56) // int32 | allowedoriginId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AllowedOriginsAPI.DeleteSystemAllowedoriginsById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedOriginsAPI.DeleteSystemAllowedoriginsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | allowedoriginId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemAllowedoriginsByIdRequest struct via the builder pattern


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


## GetSystemAllowedorigins

> []AllowedOrigin GetSystemAllowedorigins(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of AllowedOrigin

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
	resp, r, err := apiClient.AllowedOriginsAPI.GetSystemAllowedorigins(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedOriginsAPI.GetSystemAllowedorigins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemAllowedorigins`: []AllowedOrigin
	fmt.Fprintf(os.Stdout, "Response from `AllowedOriginsAPI.GetSystemAllowedorigins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemAllowedoriginsRequest struct via the builder pattern


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

[**[]AllowedOrigin**](AllowedOrigin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemAllowedoriginsById

> AllowedOrigin GetSystemAllowedoriginsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get AllowedOrigin

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
	id := int32(56) // int32 | allowedoriginId
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
	resp, r, err := apiClient.AllowedOriginsAPI.GetSystemAllowedoriginsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedOriginsAPI.GetSystemAllowedoriginsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemAllowedoriginsById`: AllowedOrigin
	fmt.Fprintf(os.Stdout, "Response from `AllowedOriginsAPI.GetSystemAllowedoriginsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | allowedoriginId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemAllowedoriginsByIdRequest struct via the builder pattern


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

[**AllowedOrigin**](AllowedOrigin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemAllowedoriginsCount

> Count GetSystemAllowedoriginsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of AllowedOrigin

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
	resp, r, err := apiClient.AllowedOriginsAPI.GetSystemAllowedoriginsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedOriginsAPI.GetSystemAllowedoriginsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemAllowedoriginsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `AllowedOriginsAPI.GetSystemAllowedoriginsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemAllowedoriginsCountRequest struct via the builder pattern


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


## PatchSystemAllowedoriginsById

> AllowedOrigin PatchSystemAllowedoriginsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch AllowedOrigin

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
	id := int32(56) // int32 | allowedoriginId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllowedOriginsAPI.PatchSystemAllowedoriginsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedOriginsAPI.PatchSystemAllowedoriginsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemAllowedoriginsById`: AllowedOrigin
	fmt.Fprintf(os.Stdout, "Response from `AllowedOriginsAPI.PatchSystemAllowedoriginsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | allowedoriginId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemAllowedoriginsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**AllowedOrigin**](AllowedOrigin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemAllowedorigins

> AllowedOrigin PostSystemAllowedorigins(ctx).ClientId(clientId).AllowedOrigin(allowedOrigin).Execute()

Post AllowedOrigin

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
	allowedOrigin := *openapiclient.NewAllowedOrigin("Origin_example", "Description_example") // AllowedOrigin | origin

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllowedOriginsAPI.PostSystemAllowedorigins(context.Background()).ClientId(clientId).AllowedOrigin(allowedOrigin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedOriginsAPI.PostSystemAllowedorigins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemAllowedorigins`: AllowedOrigin
	fmt.Fprintf(os.Stdout, "Response from `AllowedOriginsAPI.PostSystemAllowedorigins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemAllowedoriginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **allowedOrigin** | [**AllowedOrigin**](AllowedOrigin.md) | origin | 

### Return type

[**AllowedOrigin**](AllowedOrigin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemAllowedoriginsById

> AllowedOrigin PutSystemAllowedoriginsById(ctx, id).ClientId(clientId).AllowedOrigin(allowedOrigin).Execute()

Put AllowedOrigin

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
	id := int32(56) // int32 | allowedoriginId
	clientId := "clientId_example" // string | 
	allowedOrigin := *openapiclient.NewAllowedOrigin("Origin_example", "Description_example") // AllowedOrigin | origin

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllowedOriginsAPI.PutSystemAllowedoriginsById(context.Background(), id).ClientId(clientId).AllowedOrigin(allowedOrigin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedOriginsAPI.PutSystemAllowedoriginsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemAllowedoriginsById`: AllowedOrigin
	fmt.Fprintf(os.Stdout, "Response from `AllowedOriginsAPI.PutSystemAllowedoriginsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | allowedoriginId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemAllowedoriginsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **allowedOrigin** | [**AllowedOrigin**](AllowedOrigin.md) | origin | 

### Return type

[**AllowedOrigin**](AllowedOrigin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

