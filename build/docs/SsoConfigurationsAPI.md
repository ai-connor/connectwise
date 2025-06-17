# \SsoConfigurationsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemSsoConfigurationsById**](SsoConfigurationsAPI.md#DeleteSystemSsoConfigurationsById) | **Delete** /system/ssoConfigurations/{id} | Delete SsoConfiguration
[**GetSystemSsoConfigurations**](SsoConfigurationsAPI.md#GetSystemSsoConfigurations) | **Get** /system/ssoConfigurations | Get List of SsoConfiguration
[**GetSystemSsoConfigurationsById**](SsoConfigurationsAPI.md#GetSystemSsoConfigurationsById) | **Get** /system/ssoConfigurations/{id} | Get SsoConfiguration
[**GetSystemSsoConfigurationsCount**](SsoConfigurationsAPI.md#GetSystemSsoConfigurationsCount) | **Get** /system/ssoConfigurations/count | Get Count of SsoConfiguration
[**PatchSystemSsoConfigurationsById**](SsoConfigurationsAPI.md#PatchSystemSsoConfigurationsById) | **Patch** /system/ssoConfigurations/{id} | Patch SsoConfiguration
[**PostSystemSsoConfigurations**](SsoConfigurationsAPI.md#PostSystemSsoConfigurations) | **Post** /system/ssoConfigurations | Post SsoConfiguration
[**PostSystemSsoConfigurationsByIdRegistertoken**](SsoConfigurationsAPI.md#PostSystemSsoConfigurationsByIdRegistertoken) | **Post** /system/ssoConfigurations/{id}/registertoken | Post SsoConfiguration
[**PostSystemSsoConfigurationsByIdSubmitmembers**](SsoConfigurationsAPI.md#PostSystemSsoConfigurationsByIdSubmitmembers) | **Post** /system/ssoConfigurations/{id}/submitmembers | Post SsoConfiguration
[**PutSystemSsoConfigurationsById**](SsoConfigurationsAPI.md#PutSystemSsoConfigurationsById) | **Put** /system/ssoConfigurations/{id} | Put SsoConfiguration



## DeleteSystemSsoConfigurationsById

> DeleteSystemSsoConfigurationsById(ctx, id).ClientId(clientId).Execute()

Delete SsoConfiguration

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
	id := int32(56) // int32 | ssoConfigurationId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SsoConfigurationsAPI.DeleteSystemSsoConfigurationsById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoConfigurationsAPI.DeleteSystemSsoConfigurationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ssoConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemSsoConfigurationsByIdRequest struct via the builder pattern


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


## GetSystemSsoConfigurations

> []SsoConfiguration GetSystemSsoConfigurations(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of SsoConfiguration

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
	resp, r, err := apiClient.SsoConfigurationsAPI.GetSystemSsoConfigurations(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoConfigurationsAPI.GetSystemSsoConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemSsoConfigurations`: []SsoConfiguration
	fmt.Fprintf(os.Stdout, "Response from `SsoConfigurationsAPI.GetSystemSsoConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemSsoConfigurationsRequest struct via the builder pattern


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

[**[]SsoConfiguration**](SsoConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemSsoConfigurationsById

> SsoConfiguration GetSystemSsoConfigurationsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get SsoConfiguration

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
	id := int32(56) // int32 | ssoConfigurationId
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
	resp, r, err := apiClient.SsoConfigurationsAPI.GetSystemSsoConfigurationsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoConfigurationsAPI.GetSystemSsoConfigurationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemSsoConfigurationsById`: SsoConfiguration
	fmt.Fprintf(os.Stdout, "Response from `SsoConfigurationsAPI.GetSystemSsoConfigurationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ssoConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemSsoConfigurationsByIdRequest struct via the builder pattern


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

[**SsoConfiguration**](SsoConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemSsoConfigurationsCount

> Count GetSystemSsoConfigurationsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of SsoConfiguration

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
	resp, r, err := apiClient.SsoConfigurationsAPI.GetSystemSsoConfigurationsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoConfigurationsAPI.GetSystemSsoConfigurationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemSsoConfigurationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `SsoConfigurationsAPI.GetSystemSsoConfigurationsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemSsoConfigurationsCountRequest struct via the builder pattern


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


## PatchSystemSsoConfigurationsById

> SsoConfiguration PatchSystemSsoConfigurationsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch SsoConfiguration

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
	id := int32(56) // int32 | ssoConfigurationId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SsoConfigurationsAPI.PatchSystemSsoConfigurationsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoConfigurationsAPI.PatchSystemSsoConfigurationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemSsoConfigurationsById`: SsoConfiguration
	fmt.Fprintf(os.Stdout, "Response from `SsoConfigurationsAPI.PatchSystemSsoConfigurationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ssoConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemSsoConfigurationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**SsoConfiguration**](SsoConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemSsoConfigurations

> SsoConfiguration PostSystemSsoConfigurations(ctx).ClientId(clientId).SsoConfiguration(ssoConfiguration).Execute()

Post SsoConfiguration

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
	ssoConfiguration := *openapiclient.NewSsoConfiguration("Name_example", "SsoType_example", []int32{int32(123)}) // SsoConfiguration | ssoConfiguration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SsoConfigurationsAPI.PostSystemSsoConfigurations(context.Background()).ClientId(clientId).SsoConfiguration(ssoConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoConfigurationsAPI.PostSystemSsoConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemSsoConfigurations`: SsoConfiguration
	fmt.Fprintf(os.Stdout, "Response from `SsoConfigurationsAPI.PostSystemSsoConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemSsoConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **ssoConfiguration** | [**SsoConfiguration**](SsoConfiguration.md) | ssoConfiguration | 

### Return type

[**SsoConfiguration**](SsoConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemSsoConfigurationsByIdRegistertoken

> SsoConfiguration PostSystemSsoConfigurationsByIdRegistertoken(ctx, id).ClientId(clientId).SsoConfiguration(ssoConfiguration).Execute()

Post SsoConfiguration

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
	id := int32(56) // int32 | ssoConfigurationId
	clientId := "clientId_example" // string | 
	ssoConfiguration := *openapiclient.NewSsoConfiguration("Name_example", "SsoType_example", []int32{int32(123)}) // SsoConfiguration | ssoConfiguration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SsoConfigurationsAPI.PostSystemSsoConfigurationsByIdRegistertoken(context.Background(), id).ClientId(clientId).SsoConfiguration(ssoConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoConfigurationsAPI.PostSystemSsoConfigurationsByIdRegistertoken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemSsoConfigurationsByIdRegistertoken`: SsoConfiguration
	fmt.Fprintf(os.Stdout, "Response from `SsoConfigurationsAPI.PostSystemSsoConfigurationsByIdRegistertoken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ssoConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemSsoConfigurationsByIdRegistertokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **ssoConfiguration** | [**SsoConfiguration**](SsoConfiguration.md) | ssoConfiguration | 

### Return type

[**SsoConfiguration**](SsoConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemSsoConfigurationsByIdSubmitmembers

> SsoConfiguration PostSystemSsoConfigurationsByIdSubmitmembers(ctx, id).ClientId(clientId).SsoConfiguration(ssoConfiguration).Execute()

Post SsoConfiguration

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
	id := int32(56) // int32 | ssoConfigurationId
	clientId := "clientId_example" // string | 
	ssoConfiguration := *openapiclient.NewSsoConfiguration("Name_example", "SsoType_example", []int32{int32(123)}) // SsoConfiguration | ssoConfiguration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SsoConfigurationsAPI.PostSystemSsoConfigurationsByIdSubmitmembers(context.Background(), id).ClientId(clientId).SsoConfiguration(ssoConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoConfigurationsAPI.PostSystemSsoConfigurationsByIdSubmitmembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemSsoConfigurationsByIdSubmitmembers`: SsoConfiguration
	fmt.Fprintf(os.Stdout, "Response from `SsoConfigurationsAPI.PostSystemSsoConfigurationsByIdSubmitmembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ssoConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemSsoConfigurationsByIdSubmitmembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **ssoConfiguration** | [**SsoConfiguration**](SsoConfiguration.md) | ssoConfiguration | 

### Return type

[**SsoConfiguration**](SsoConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemSsoConfigurationsById

> SsoConfiguration PutSystemSsoConfigurationsById(ctx, id).ClientId(clientId).SsoConfiguration(ssoConfiguration).Execute()

Put SsoConfiguration

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
	id := int32(56) // int32 | ssoConfigurationId
	clientId := "clientId_example" // string | 
	ssoConfiguration := *openapiclient.NewSsoConfiguration("Name_example", "SsoType_example", []int32{int32(123)}) // SsoConfiguration | ssoConfiguration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SsoConfigurationsAPI.PutSystemSsoConfigurationsById(context.Background(), id).ClientId(clientId).SsoConfiguration(ssoConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SsoConfigurationsAPI.PutSystemSsoConfigurationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemSsoConfigurationsById`: SsoConfiguration
	fmt.Fprintf(os.Stdout, "Response from `SsoConfigurationsAPI.PutSystemSsoConfigurationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ssoConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemSsoConfigurationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **ssoConfiguration** | [**SsoConfiguration**](SsoConfiguration.md) | ssoConfiguration | 

### Return type

[**SsoConfiguration**](SsoConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

