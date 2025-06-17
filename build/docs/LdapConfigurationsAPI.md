# \LdapConfigurationsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemLdapConfigurationsById**](LdapConfigurationsAPI.md#DeleteSystemLdapConfigurationsById) | **Delete** /system/ldapConfigurations/{id} | Delete LdapConfiguration
[**GetSystemLdapConfigurations**](LdapConfigurationsAPI.md#GetSystemLdapConfigurations) | **Get** /system/ldapConfigurations | Get List of LdapConfiguration
[**GetSystemLdapConfigurationsById**](LdapConfigurationsAPI.md#GetSystemLdapConfigurationsById) | **Get** /system/ldapConfigurations/{id} | Get LdapConfiguration
[**GetSystemLdapConfigurationsCount**](LdapConfigurationsAPI.md#GetSystemLdapConfigurationsCount) | **Get** /system/ldapConfigurations/count | Get Count of LdapConfiguration
[**PatchSystemLdapConfigurationsById**](LdapConfigurationsAPI.md#PatchSystemLdapConfigurationsById) | **Patch** /system/ldapConfigurations/{id} | Patch LdapConfiguration
[**PostSystemLdapConfigurations**](LdapConfigurationsAPI.md#PostSystemLdapConfigurations) | **Post** /system/ldapConfigurations | Post LdapConfiguration
[**PostSystemLdapConfigurationsTestLink**](LdapConfigurationsAPI.md#PostSystemLdapConfigurationsTestLink) | **Post** /system/ldapConfigurations/testLink | Post SuccessResponse
[**PutSystemLdapConfigurationsById**](LdapConfigurationsAPI.md#PutSystemLdapConfigurationsById) | **Put** /system/ldapConfigurations/{id} | Put LdapConfiguration



## DeleteSystemLdapConfigurationsById

> DeleteSystemLdapConfigurationsById(ctx, id).ClientId(clientId).Execute()

Delete LdapConfiguration

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
	id := int32(56) // int32 | ldapConfigurationId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LdapConfigurationsAPI.DeleteSystemLdapConfigurationsById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LdapConfigurationsAPI.DeleteSystemLdapConfigurationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ldapConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemLdapConfigurationsByIdRequest struct via the builder pattern


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


## GetSystemLdapConfigurations

> []LdapConfiguration GetSystemLdapConfigurations(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of LdapConfiguration

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
	resp, r, err := apiClient.LdapConfigurationsAPI.GetSystemLdapConfigurations(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LdapConfigurationsAPI.GetSystemLdapConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemLdapConfigurations`: []LdapConfiguration
	fmt.Fprintf(os.Stdout, "Response from `LdapConfigurationsAPI.GetSystemLdapConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemLdapConfigurationsRequest struct via the builder pattern


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

[**[]LdapConfiguration**](LdapConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemLdapConfigurationsById

> LdapConfiguration GetSystemLdapConfigurationsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get LdapConfiguration

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
	id := int32(56) // int32 | ldapConfigurationId
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
	resp, r, err := apiClient.LdapConfigurationsAPI.GetSystemLdapConfigurationsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LdapConfigurationsAPI.GetSystemLdapConfigurationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemLdapConfigurationsById`: LdapConfiguration
	fmt.Fprintf(os.Stdout, "Response from `LdapConfigurationsAPI.GetSystemLdapConfigurationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ldapConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemLdapConfigurationsByIdRequest struct via the builder pattern


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

[**LdapConfiguration**](LdapConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemLdapConfigurationsCount

> Count GetSystemLdapConfigurationsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of LdapConfiguration

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
	resp, r, err := apiClient.LdapConfigurationsAPI.GetSystemLdapConfigurationsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LdapConfigurationsAPI.GetSystemLdapConfigurationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemLdapConfigurationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `LdapConfigurationsAPI.GetSystemLdapConfigurationsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemLdapConfigurationsCountRequest struct via the builder pattern


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


## PatchSystemLdapConfigurationsById

> LdapConfiguration PatchSystemLdapConfigurationsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch LdapConfiguration

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
	id := int32(56) // int32 | ldapConfigurationId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LdapConfigurationsAPI.PatchSystemLdapConfigurationsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LdapConfigurationsAPI.PatchSystemLdapConfigurationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemLdapConfigurationsById`: LdapConfiguration
	fmt.Fprintf(os.Stdout, "Response from `LdapConfigurationsAPI.PatchSystemLdapConfigurationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ldapConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemLdapConfigurationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**LdapConfiguration**](LdapConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemLdapConfigurations

> LdapConfiguration PostSystemLdapConfigurations(ctx).ClientId(clientId).LdapConfiguration(ldapConfiguration).Execute()

Post LdapConfiguration

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
	ldapConfiguration := *openapiclient.NewLdapConfiguration("Name_example", "Server_example", "Domain_example") // LdapConfiguration | ldapConfiguration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LdapConfigurationsAPI.PostSystemLdapConfigurations(context.Background()).ClientId(clientId).LdapConfiguration(ldapConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LdapConfigurationsAPI.PostSystemLdapConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemLdapConfigurations`: LdapConfiguration
	fmt.Fprintf(os.Stdout, "Response from `LdapConfigurationsAPI.PostSystemLdapConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemLdapConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **ldapConfiguration** | [**LdapConfiguration**](LdapConfiguration.md) | ldapConfiguration | 

### Return type

[**LdapConfiguration**](LdapConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemLdapConfigurationsTestLink

> SuccessResponse PostSystemLdapConfigurationsTestLink(ctx).ClientId(clientId).LdapConfigurationTestLink(ldapConfigurationTestLink).Execute()

Post SuccessResponse

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
	ldapConfigurationTestLink := *openapiclient.NewLdapConfigurationTestLink() // LdapConfigurationTestLink | server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LdapConfigurationsAPI.PostSystemLdapConfigurationsTestLink(context.Background()).ClientId(clientId).LdapConfigurationTestLink(ldapConfigurationTestLink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LdapConfigurationsAPI.PostSystemLdapConfigurationsTestLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemLdapConfigurationsTestLink`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `LdapConfigurationsAPI.PostSystemLdapConfigurationsTestLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemLdapConfigurationsTestLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **ldapConfigurationTestLink** | [**LdapConfigurationTestLink**](LdapConfigurationTestLink.md) | server | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemLdapConfigurationsById

> LdapConfiguration PutSystemLdapConfigurationsById(ctx, id).ClientId(clientId).LdapConfiguration(ldapConfiguration).Execute()

Put LdapConfiguration

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
	id := int32(56) // int32 | ldapConfigurationId
	clientId := "clientId_example" // string | 
	ldapConfiguration := *openapiclient.NewLdapConfiguration("Name_example", "Server_example", "Domain_example") // LdapConfiguration | ldapConfiguration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LdapConfigurationsAPI.PutSystemLdapConfigurationsById(context.Background(), id).ClientId(clientId).LdapConfiguration(ldapConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LdapConfigurationsAPI.PutSystemLdapConfigurationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemLdapConfigurationsById`: LdapConfiguration
	fmt.Fprintf(os.Stdout, "Response from `LdapConfigurationsAPI.PutSystemLdapConfigurationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ldapConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemLdapConfigurationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **ldapConfiguration** | [**LdapConfiguration**](LdapConfiguration.md) | ldapConfiguration | 

### Return type

[**LdapConfiguration**](LdapConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

