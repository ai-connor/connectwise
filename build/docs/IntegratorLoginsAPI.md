# \IntegratorLoginsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemIntegratorloginsById**](IntegratorLoginsAPI.md#DeleteSystemIntegratorloginsById) | **Delete** /system/integratorlogins/{id} | Delete IntegratorLogin
[**GetSystemIntegratorlogins**](IntegratorLoginsAPI.md#GetSystemIntegratorlogins) | **Get** /system/integratorlogins | Get List of IntegratorLogin
[**GetSystemIntegratorloginsById**](IntegratorLoginsAPI.md#GetSystemIntegratorloginsById) | **Get** /system/integratorlogins/{id} | Get IntegratorLogin
[**GetSystemIntegratorloginsCount**](IntegratorLoginsAPI.md#GetSystemIntegratorloginsCount) | **Get** /system/integratorlogins/count | Get Count of IntegratorLogin
[**PatchSystemIntegratorloginsById**](IntegratorLoginsAPI.md#PatchSystemIntegratorloginsById) | **Patch** /system/integratorlogins/{id} | Patch IntegratorLogin
[**PostSystemIntegratorlogins**](IntegratorLoginsAPI.md#PostSystemIntegratorlogins) | **Post** /system/integratorlogins | Post IntegratorLogin
[**PutSystemIntegratorloginsById**](IntegratorLoginsAPI.md#PutSystemIntegratorloginsById) | **Put** /system/integratorlogins/{id} | Put IntegratorLogin



## DeleteSystemIntegratorloginsById

> DeleteSystemIntegratorloginsById(ctx, id).ClientId(clientId).Execute()

Delete IntegratorLogin

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
	id := int32(56) // int32 | integratorloginId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegratorLoginsAPI.DeleteSystemIntegratorloginsById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegratorLoginsAPI.DeleteSystemIntegratorloginsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | integratorloginId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemIntegratorloginsByIdRequest struct via the builder pattern


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


## GetSystemIntegratorlogins

> []IntegratorLogin GetSystemIntegratorlogins(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of IntegratorLogin

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
	resp, r, err := apiClient.IntegratorLoginsAPI.GetSystemIntegratorlogins(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegratorLoginsAPI.GetSystemIntegratorlogins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemIntegratorlogins`: []IntegratorLogin
	fmt.Fprintf(os.Stdout, "Response from `IntegratorLoginsAPI.GetSystemIntegratorlogins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemIntegratorloginsRequest struct via the builder pattern


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

[**[]IntegratorLogin**](IntegratorLogin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemIntegratorloginsById

> IntegratorLogin GetSystemIntegratorloginsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get IntegratorLogin

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
	id := int32(56) // int32 | integratorloginId
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
	resp, r, err := apiClient.IntegratorLoginsAPI.GetSystemIntegratorloginsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegratorLoginsAPI.GetSystemIntegratorloginsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemIntegratorloginsById`: IntegratorLogin
	fmt.Fprintf(os.Stdout, "Response from `IntegratorLoginsAPI.GetSystemIntegratorloginsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | integratorloginId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemIntegratorloginsByIdRequest struct via the builder pattern


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

[**IntegratorLogin**](IntegratorLogin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemIntegratorloginsCount

> Count GetSystemIntegratorloginsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of IntegratorLogin

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
	resp, r, err := apiClient.IntegratorLoginsAPI.GetSystemIntegratorloginsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegratorLoginsAPI.GetSystemIntegratorloginsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemIntegratorloginsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `IntegratorLoginsAPI.GetSystemIntegratorloginsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemIntegratorloginsCountRequest struct via the builder pattern


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


## PatchSystemIntegratorloginsById

> IntegratorLogin PatchSystemIntegratorloginsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch IntegratorLogin

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
	id := int32(56) // int32 | integratorloginId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegratorLoginsAPI.PatchSystemIntegratorloginsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegratorLoginsAPI.PatchSystemIntegratorloginsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemIntegratorloginsById`: IntegratorLogin
	fmt.Fprintf(os.Stdout, "Response from `IntegratorLoginsAPI.PatchSystemIntegratorloginsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | integratorloginId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemIntegratorloginsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**IntegratorLogin**](IntegratorLogin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemIntegratorlogins

> IntegratorLogin PostSystemIntegratorlogins(ctx).ClientId(clientId).IntegratorLogin(integratorLogin).Execute()

Post IntegratorLogin

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
	integratorLogin := *openapiclient.NewIntegratorLogin("Username_example") // IntegratorLogin | integratorLogin

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegratorLoginsAPI.PostSystemIntegratorlogins(context.Background()).ClientId(clientId).IntegratorLogin(integratorLogin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegratorLoginsAPI.PostSystemIntegratorlogins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemIntegratorlogins`: IntegratorLogin
	fmt.Fprintf(os.Stdout, "Response from `IntegratorLoginsAPI.PostSystemIntegratorlogins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemIntegratorloginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **integratorLogin** | [**IntegratorLogin**](IntegratorLogin.md) | integratorLogin | 

### Return type

[**IntegratorLogin**](IntegratorLogin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemIntegratorloginsById

> IntegratorLogin PutSystemIntegratorloginsById(ctx, id).ClientId(clientId).IntegratorLogin(integratorLogin).Execute()

Put IntegratorLogin

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
	id := int32(56) // int32 | integratorloginId
	clientId := "clientId_example" // string | 
	integratorLogin := *openapiclient.NewIntegratorLogin("Username_example") // IntegratorLogin | integratorLogin

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegratorLoginsAPI.PutSystemIntegratorloginsById(context.Background(), id).ClientId(clientId).IntegratorLogin(integratorLogin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegratorLoginsAPI.PutSystemIntegratorloginsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemIntegratorloginsById`: IntegratorLogin
	fmt.Fprintf(os.Stdout, "Response from `IntegratorLoginsAPI.PutSystemIntegratorloginsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | integratorloginId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemIntegratorloginsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **integratorLogin** | [**IntegratorLogin**](IntegratorLogin.md) | integratorLogin | 

### Return type

[**IntegratorLogin**](IntegratorLogin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

