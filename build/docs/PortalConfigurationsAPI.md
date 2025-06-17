# \PortalConfigurationsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyPortalConfigurationsById**](PortalConfigurationsAPI.md#DeleteCompanyPortalConfigurationsById) | **Delete** /company/portalConfigurations/{id} | Delete PortalConfiguration
[**GetCompanyPortalConfigurations**](PortalConfigurationsAPI.md#GetCompanyPortalConfigurations) | **Get** /company/portalConfigurations | Get List of PortalConfiguration
[**GetCompanyPortalConfigurationsById**](PortalConfigurationsAPI.md#GetCompanyPortalConfigurationsById) | **Get** /company/portalConfigurations/{id} | Get PortalConfiguration
[**GetCompanyPortalConfigurationsCount**](PortalConfigurationsAPI.md#GetCompanyPortalConfigurationsCount) | **Get** /company/portalConfigurations/count | Get Count of PortalConfiguration
[**PatchCompanyPortalConfigurationsById**](PortalConfigurationsAPI.md#PatchCompanyPortalConfigurationsById) | **Patch** /company/portalConfigurations/{id} | Patch PortalConfiguration
[**PostCompanyPortalConfigurations**](PortalConfigurationsAPI.md#PostCompanyPortalConfigurations) | **Post** /company/portalConfigurations | Post PortalConfiguration
[**PostCompanyPortalConfigurationsCopy**](PortalConfigurationsAPI.md#PostCompanyPortalConfigurationsCopy) | **Post** /company/portalConfigurations/copy | Post PortalConfiguration
[**PutCompanyPortalConfigurationsById**](PortalConfigurationsAPI.md#PutCompanyPortalConfigurationsById) | **Put** /company/portalConfigurations/{id} | Put PortalConfiguration



## DeleteCompanyPortalConfigurationsById

> DeleteCompanyPortalConfigurationsById(ctx, id).ClientId(clientId).Execute()

Delete PortalConfiguration

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
	id := int32(56) // int32 | portalConfigurationId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PortalConfigurationsAPI.DeleteCompanyPortalConfigurationsById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationsAPI.DeleteCompanyPortalConfigurationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyPortalConfigurationsByIdRequest struct via the builder pattern


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


## GetCompanyPortalConfigurations

> []PortalConfiguration GetCompanyPortalConfigurations(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of PortalConfiguration

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
	resp, r, err := apiClient.PortalConfigurationsAPI.GetCompanyPortalConfigurations(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationsAPI.GetCompanyPortalConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyPortalConfigurations`: []PortalConfiguration
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationsAPI.GetCompanyPortalConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPortalConfigurationsRequest struct via the builder pattern


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

[**[]PortalConfiguration**](PortalConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyPortalConfigurationsById

> PortalConfiguration GetCompanyPortalConfigurationsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get PortalConfiguration

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
	id := int32(56) // int32 | portalConfigurationId
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
	resp, r, err := apiClient.PortalConfigurationsAPI.GetCompanyPortalConfigurationsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationsAPI.GetCompanyPortalConfigurationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyPortalConfigurationsById`: PortalConfiguration
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationsAPI.GetCompanyPortalConfigurationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPortalConfigurationsByIdRequest struct via the builder pattern


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

[**PortalConfiguration**](PortalConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyPortalConfigurationsCount

> Count GetCompanyPortalConfigurationsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of PortalConfiguration

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
	resp, r, err := apiClient.PortalConfigurationsAPI.GetCompanyPortalConfigurationsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationsAPI.GetCompanyPortalConfigurationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyPortalConfigurationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationsAPI.GetCompanyPortalConfigurationsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPortalConfigurationsCountRequest struct via the builder pattern


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


## PatchCompanyPortalConfigurationsById

> PortalConfiguration PatchCompanyPortalConfigurationsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch PortalConfiguration

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
	id := int32(56) // int32 | portalConfigurationId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalConfigurationsAPI.PatchCompanyPortalConfigurationsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationsAPI.PatchCompanyPortalConfigurationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyPortalConfigurationsById`: PortalConfiguration
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationsAPI.PatchCompanyPortalConfigurationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyPortalConfigurationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**PortalConfiguration**](PortalConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyPortalConfigurations

> PortalConfiguration PostCompanyPortalConfigurations(ctx).ClientId(clientId).PortalConfiguration(portalConfiguration).Execute()

Post PortalConfiguration

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
	portalConfiguration := *openapiclient.NewPortalConfiguration("Name_example") // PortalConfiguration | portalConfiguration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalConfigurationsAPI.PostCompanyPortalConfigurations(context.Background()).ClientId(clientId).PortalConfiguration(portalConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationsAPI.PostCompanyPortalConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyPortalConfigurations`: PortalConfiguration
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationsAPI.PostCompanyPortalConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyPortalConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **portalConfiguration** | [**PortalConfiguration**](PortalConfiguration.md) | portalConfiguration | 

### Return type

[**PortalConfiguration**](PortalConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyPortalConfigurationsCopy

> PortalConfiguration PostCompanyPortalConfigurationsCopy(ctx).ClientId(clientId).PortalConfiguration(portalConfiguration).Execute()

Post PortalConfiguration

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
	portalConfiguration := *openapiclient.NewPortalConfiguration("Name_example") // PortalConfiguration | copy

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalConfigurationsAPI.PostCompanyPortalConfigurationsCopy(context.Background()).ClientId(clientId).PortalConfiguration(portalConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationsAPI.PostCompanyPortalConfigurationsCopy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyPortalConfigurationsCopy`: PortalConfiguration
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationsAPI.PostCompanyPortalConfigurationsCopy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyPortalConfigurationsCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **portalConfiguration** | [**PortalConfiguration**](PortalConfiguration.md) | copy | 

### Return type

[**PortalConfiguration**](PortalConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyPortalConfigurationsById

> PortalConfiguration PutCompanyPortalConfigurationsById(ctx, id).ClientId(clientId).PortalConfiguration(portalConfiguration).Execute()

Put PortalConfiguration

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
	id := int32(56) // int32 | portalConfigurationId
	clientId := "clientId_example" // string | 
	portalConfiguration := *openapiclient.NewPortalConfiguration("Name_example") // PortalConfiguration | portalConfiguration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalConfigurationsAPI.PutCompanyPortalConfigurationsById(context.Background(), id).ClientId(clientId).PortalConfiguration(portalConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationsAPI.PutCompanyPortalConfigurationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyPortalConfigurationsById`: PortalConfiguration
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationsAPI.PutCompanyPortalConfigurationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyPortalConfigurationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **portalConfiguration** | [**PortalConfiguration**](PortalConfiguration.md) | portalConfiguration | 

### Return type

[**PortalConfiguration**](PortalConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

