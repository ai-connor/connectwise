# \ConnectWiseHostedSetupsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemConnectwisehostedsetupsById**](ConnectWiseHostedSetupsAPI.md#DeleteSystemConnectwisehostedsetupsById) | **Delete** /system/connectwisehostedsetups/{id} | Delete ConnectWiseHostedSetup
[**GetSystemConnectwisehostedsetups**](ConnectWiseHostedSetupsAPI.md#GetSystemConnectwisehostedsetups) | **Get** /system/connectwisehostedsetups | Get List of ConnectWiseHostedSetup
[**GetSystemConnectwisehostedsetupsById**](ConnectWiseHostedSetupsAPI.md#GetSystemConnectwisehostedsetupsById) | **Get** /system/connectwisehostedsetups/{id} | Get ConnectWiseHostedSetup
[**GetSystemConnectwisehostedsetupsCount**](ConnectWiseHostedSetupsAPI.md#GetSystemConnectwisehostedsetupsCount) | **Get** /system/connectwisehostedsetups/count | Get Count of ConnectWiseHostedSetup
[**PatchSystemConnectwisehostedsetupsById**](ConnectWiseHostedSetupsAPI.md#PatchSystemConnectwisehostedsetupsById) | **Patch** /system/connectwisehostedsetups/{id} | Patch ConnectWiseHostedSetup
[**PostSystemConnectwisehostedsetups**](ConnectWiseHostedSetupsAPI.md#PostSystemConnectwisehostedsetups) | **Post** /system/connectwisehostedsetups | Post ConnectWiseHostedSetup
[**PutSystemConnectwisehostedsetupsById**](ConnectWiseHostedSetupsAPI.md#PutSystemConnectwisehostedsetupsById) | **Put** /system/connectwisehostedsetups/{id} | Put ConnectWiseHostedSetup



## DeleteSystemConnectwisehostedsetupsById

> DeleteSystemConnectwisehostedsetupsById(ctx, id).ClientId(clientId).Execute()

Delete ConnectWiseHostedSetup

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
	id := int32(56) // int32 | connectwisehostedsetupId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectWiseHostedSetupsAPI.DeleteSystemConnectwisehostedsetupsById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectWiseHostedSetupsAPI.DeleteSystemConnectwisehostedsetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | connectwisehostedsetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemConnectwisehostedsetupsByIdRequest struct via the builder pattern


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


## GetSystemConnectwisehostedsetups

> []ConnectWiseHostedSetup GetSystemConnectwisehostedsetups(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ConnectWiseHostedSetup

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
	resp, r, err := apiClient.ConnectWiseHostedSetupsAPI.GetSystemConnectwisehostedsetups(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectWiseHostedSetupsAPI.GetSystemConnectwisehostedsetups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemConnectwisehostedsetups`: []ConnectWiseHostedSetup
	fmt.Fprintf(os.Stdout, "Response from `ConnectWiseHostedSetupsAPI.GetSystemConnectwisehostedsetups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemConnectwisehostedsetupsRequest struct via the builder pattern


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

[**[]ConnectWiseHostedSetup**](ConnectWiseHostedSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemConnectwisehostedsetupsById

> ConnectWiseHostedSetup GetSystemConnectwisehostedsetupsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ConnectWiseHostedSetup

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
	id := int32(56) // int32 | connectwisehostedsetupId
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
	resp, r, err := apiClient.ConnectWiseHostedSetupsAPI.GetSystemConnectwisehostedsetupsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectWiseHostedSetupsAPI.GetSystemConnectwisehostedsetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemConnectwisehostedsetupsById`: ConnectWiseHostedSetup
	fmt.Fprintf(os.Stdout, "Response from `ConnectWiseHostedSetupsAPI.GetSystemConnectwisehostedsetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | connectwisehostedsetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemConnectwisehostedsetupsByIdRequest struct via the builder pattern


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

[**ConnectWiseHostedSetup**](ConnectWiseHostedSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemConnectwisehostedsetupsCount

> Count GetSystemConnectwisehostedsetupsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ConnectWiseHostedSetup

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
	resp, r, err := apiClient.ConnectWiseHostedSetupsAPI.GetSystemConnectwisehostedsetupsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectWiseHostedSetupsAPI.GetSystemConnectwisehostedsetupsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemConnectwisehostedsetupsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ConnectWiseHostedSetupsAPI.GetSystemConnectwisehostedsetupsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemConnectwisehostedsetupsCountRequest struct via the builder pattern


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


## PatchSystemConnectwisehostedsetupsById

> ConnectWiseHostedSetup PatchSystemConnectwisehostedsetupsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ConnectWiseHostedSetup

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
	id := int32(56) // int32 | connectwisehostedsetupId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectWiseHostedSetupsAPI.PatchSystemConnectwisehostedsetupsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectWiseHostedSetupsAPI.PatchSystemConnectwisehostedsetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemConnectwisehostedsetupsById`: ConnectWiseHostedSetup
	fmt.Fprintf(os.Stdout, "Response from `ConnectWiseHostedSetupsAPI.PatchSystemConnectwisehostedsetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | connectwisehostedsetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemConnectwisehostedsetupsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ConnectWiseHostedSetup**](ConnectWiseHostedSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemConnectwisehostedsetups

> ConnectWiseHostedSetup PostSystemConnectwisehostedsetups(ctx).ClientId(clientId).ConnectWiseHostedSetup(connectWiseHostedSetup).Execute()

Post ConnectWiseHostedSetup

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
	connectWiseHostedSetup := *openapiclient.NewConnectWiseHostedSetup(NullableInt32(123), "Description_example", "Url_example", "Type_example") // ConnectWiseHostedSetup | connectWiseHostedSetup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectWiseHostedSetupsAPI.PostSystemConnectwisehostedsetups(context.Background()).ClientId(clientId).ConnectWiseHostedSetup(connectWiseHostedSetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectWiseHostedSetupsAPI.PostSystemConnectwisehostedsetups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemConnectwisehostedsetups`: ConnectWiseHostedSetup
	fmt.Fprintf(os.Stdout, "Response from `ConnectWiseHostedSetupsAPI.PostSystemConnectwisehostedsetups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemConnectwisehostedsetupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **connectWiseHostedSetup** | [**ConnectWiseHostedSetup**](ConnectWiseHostedSetup.md) | connectWiseHostedSetup | 

### Return type

[**ConnectWiseHostedSetup**](ConnectWiseHostedSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemConnectwisehostedsetupsById

> ConnectWiseHostedSetup PutSystemConnectwisehostedsetupsById(ctx, id).ClientId(clientId).ConnectWiseHostedSetup(connectWiseHostedSetup).Execute()

Put ConnectWiseHostedSetup

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
	id := int32(56) // int32 | connectwisehostedsetupId
	clientId := "clientId_example" // string | 
	connectWiseHostedSetup := *openapiclient.NewConnectWiseHostedSetup(NullableInt32(123), "Description_example", "Url_example", "Type_example") // ConnectWiseHostedSetup | connectWiseHostedSetup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectWiseHostedSetupsAPI.PutSystemConnectwisehostedsetupsById(context.Background(), id).ClientId(clientId).ConnectWiseHostedSetup(connectWiseHostedSetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectWiseHostedSetupsAPI.PutSystemConnectwisehostedsetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemConnectwisehostedsetupsById`: ConnectWiseHostedSetup
	fmt.Fprintf(os.Stdout, "Response from `ConnectWiseHostedSetupsAPI.PutSystemConnectwisehostedsetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | connectwisehostedsetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemConnectwisehostedsetupsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **connectWiseHostedSetup** | [**ConnectWiseHostedSetup**](ConnectWiseHostedSetup.md) | connectWiseHostedSetup | 

### Return type

[**ConnectWiseHostedSetup**](ConnectWiseHostedSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

