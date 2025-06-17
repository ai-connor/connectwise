# \EPayConfigurationsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemEPayConfigurationsById**](EPayConfigurationsAPI.md#DeleteSystemEPayConfigurationsById) | **Delete** /system/ePayConfigurations/{id} | Delete EPayConfiguration
[**GetSystemEPayConfigurations**](EPayConfigurationsAPI.md#GetSystemEPayConfigurations) | **Get** /system/ePayConfigurations | Get List of EPayConfiguration
[**GetSystemEPayConfigurationsById**](EPayConfigurationsAPI.md#GetSystemEPayConfigurationsById) | **Get** /system/ePayConfigurations/{id} | Get EPayConfiguration
[**GetSystemEPayConfigurationsCount**](EPayConfigurationsAPI.md#GetSystemEPayConfigurationsCount) | **Get** /system/ePayConfigurations/count | Get Count of EPayConfiguration
[**PatchSystemEPayConfigurationsById**](EPayConfigurationsAPI.md#PatchSystemEPayConfigurationsById) | **Patch** /system/ePayConfigurations/{id} | Patch EPayConfiguration
[**PostSystemEPayConfigurations**](EPayConfigurationsAPI.md#PostSystemEPayConfigurations) | **Post** /system/ePayConfigurations | Post EPayConfiguration
[**PutSystemEPayConfigurationsById**](EPayConfigurationsAPI.md#PutSystemEPayConfigurationsById) | **Put** /system/ePayConfigurations/{id} | Put EPayConfiguration



## DeleteSystemEPayConfigurationsById

> DeleteSystemEPayConfigurationsById(ctx, id).ClientId(clientId).Execute()

Delete EPayConfiguration

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
	id := int32(56) // int32 | ePayConfigurationId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EPayConfigurationsAPI.DeleteSystemEPayConfigurationsById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EPayConfigurationsAPI.DeleteSystemEPayConfigurationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ePayConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemEPayConfigurationsByIdRequest struct via the builder pattern


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


## GetSystemEPayConfigurations

> []EPayConfiguration GetSystemEPayConfigurations(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of EPayConfiguration

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
	resp, r, err := apiClient.EPayConfigurationsAPI.GetSystemEPayConfigurations(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EPayConfigurationsAPI.GetSystemEPayConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemEPayConfigurations`: []EPayConfiguration
	fmt.Fprintf(os.Stdout, "Response from `EPayConfigurationsAPI.GetSystemEPayConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemEPayConfigurationsRequest struct via the builder pattern


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

[**[]EPayConfiguration**](EPayConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemEPayConfigurationsById

> EPayConfiguration GetSystemEPayConfigurationsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get EPayConfiguration

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
	id := int32(56) // int32 | ePayConfigurationId
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
	resp, r, err := apiClient.EPayConfigurationsAPI.GetSystemEPayConfigurationsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EPayConfigurationsAPI.GetSystemEPayConfigurationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemEPayConfigurationsById`: EPayConfiguration
	fmt.Fprintf(os.Stdout, "Response from `EPayConfigurationsAPI.GetSystemEPayConfigurationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ePayConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemEPayConfigurationsByIdRequest struct via the builder pattern


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

[**EPayConfiguration**](EPayConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemEPayConfigurationsCount

> Count GetSystemEPayConfigurationsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of EPayConfiguration

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
	resp, r, err := apiClient.EPayConfigurationsAPI.GetSystemEPayConfigurationsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EPayConfigurationsAPI.GetSystemEPayConfigurationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemEPayConfigurationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `EPayConfigurationsAPI.GetSystemEPayConfigurationsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemEPayConfigurationsCountRequest struct via the builder pattern


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


## PatchSystemEPayConfigurationsById

> EPayConfiguration PatchSystemEPayConfigurationsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch EPayConfiguration

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
	id := int32(56) // int32 | ePayConfigurationId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EPayConfigurationsAPI.PatchSystemEPayConfigurationsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EPayConfigurationsAPI.PatchSystemEPayConfigurationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemEPayConfigurationsById`: EPayConfiguration
	fmt.Fprintf(os.Stdout, "Response from `EPayConfigurationsAPI.PatchSystemEPayConfigurationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ePayConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemEPayConfigurationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**EPayConfiguration**](EPayConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemEPayConfigurations

> EPayConfiguration PostSystemEPayConfigurations(ctx).ClientId(clientId).EPayConfiguration(ePayConfiguration).Execute()

Post EPayConfiguration

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
	ePayConfiguration := *openapiclient.NewEPayConfiguration(*openapiclient.NewSystemLocationReference(), *openapiclient.NewCurrencyReference(), "Url_example", "StoreIdentifier_example") // EPayConfiguration | ePayConfiguration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EPayConfigurationsAPI.PostSystemEPayConfigurations(context.Background()).ClientId(clientId).EPayConfiguration(ePayConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EPayConfigurationsAPI.PostSystemEPayConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemEPayConfigurations`: EPayConfiguration
	fmt.Fprintf(os.Stdout, "Response from `EPayConfigurationsAPI.PostSystemEPayConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemEPayConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **ePayConfiguration** | [**EPayConfiguration**](EPayConfiguration.md) | ePayConfiguration | 

### Return type

[**EPayConfiguration**](EPayConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemEPayConfigurationsById

> EPayConfiguration PutSystemEPayConfigurationsById(ctx, id).ClientId(clientId).EPayConfiguration(ePayConfiguration).Execute()

Put EPayConfiguration

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
	id := int32(56) // int32 | ePayConfigurationId
	clientId := "clientId_example" // string | 
	ePayConfiguration := *openapiclient.NewEPayConfiguration(*openapiclient.NewSystemLocationReference(), *openapiclient.NewCurrencyReference(), "Url_example", "StoreIdentifier_example") // EPayConfiguration | ePayConfiguration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EPayConfigurationsAPI.PutSystemEPayConfigurationsById(context.Background(), id).ClientId(clientId).EPayConfiguration(ePayConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EPayConfigurationsAPI.PutSystemEPayConfigurationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemEPayConfigurationsById`: EPayConfiguration
	fmt.Fprintf(os.Stdout, "Response from `EPayConfigurationsAPI.PutSystemEPayConfigurationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ePayConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemEPayConfigurationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **ePayConfiguration** | [**EPayConfiguration**](EPayConfiguration.md) | ePayConfiguration | 

### Return type

[**EPayConfiguration**](EPayConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

