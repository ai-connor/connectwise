# \ManagedDevicesIntegrationsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyManagedDevicesIntegrationsById**](ManagedDevicesIntegrationsAPI.md#DeleteCompanyManagedDevicesIntegrationsById) | **Delete** /company/managedDevicesIntegrations/{id} | Delete ManagedDevicesIntegration
[**GetCompanyManagedDevicesIntegrations**](ManagedDevicesIntegrationsAPI.md#GetCompanyManagedDevicesIntegrations) | **Get** /company/managedDevicesIntegrations | Get List of ManagedDevicesIntegration
[**GetCompanyManagedDevicesIntegrationsById**](ManagedDevicesIntegrationsAPI.md#GetCompanyManagedDevicesIntegrationsById) | **Get** /company/managedDevicesIntegrations/{id} | Get ManagedDevicesIntegration
[**GetCompanyManagedDevicesIntegrationsByIdUsages**](ManagedDevicesIntegrationsAPI.md#GetCompanyManagedDevicesIntegrationsByIdUsages) | **Get** /company/managedDevicesIntegrations/{id}/usages | Get List of Usage Count
[**GetCompanyManagedDevicesIntegrationsByIdUsagesList**](ManagedDevicesIntegrationsAPI.md#GetCompanyManagedDevicesIntegrationsByIdUsagesList) | **Get** /company/managedDevicesIntegrations/{id}/usages/list | Get List of Usage
[**GetCompanyManagedDevicesIntegrationsCount**](ManagedDevicesIntegrationsAPI.md#GetCompanyManagedDevicesIntegrationsCount) | **Get** /company/managedDevicesIntegrations/count | Get Count of Usage
[**PatchCompanyManagedDevicesIntegrationsById**](ManagedDevicesIntegrationsAPI.md#PatchCompanyManagedDevicesIntegrationsById) | **Patch** /company/managedDevicesIntegrations/{id} | Patch ManagedDevicesIntegration
[**PostCompanyManagedDevicesIntegrations**](ManagedDevicesIntegrationsAPI.md#PostCompanyManagedDevicesIntegrations) | **Post** /company/managedDevicesIntegrations | Post ManagedDevicesIntegration
[**PutCompanyManagedDevicesIntegrationsById**](ManagedDevicesIntegrationsAPI.md#PutCompanyManagedDevicesIntegrationsById) | **Put** /company/managedDevicesIntegrations/{id} | Put ManagedDevicesIntegration



## DeleteCompanyManagedDevicesIntegrationsById

> DeleteCompanyManagedDevicesIntegrationsById(ctx, id).ClientId(clientId).Execute()

Delete ManagedDevicesIntegration

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
	id := int32(56) // int32 | managedDevicesIntegrationId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagedDevicesIntegrationsAPI.DeleteCompanyManagedDevicesIntegrationsById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationsAPI.DeleteCompanyManagedDevicesIntegrationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyManagedDevicesIntegrationsByIdRequest struct via the builder pattern


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


## GetCompanyManagedDevicesIntegrations

> []ManagedDevicesIntegration GetCompanyManagedDevicesIntegrations(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ManagedDevicesIntegration

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
	resp, r, err := apiClient.ManagedDevicesIntegrationsAPI.GetCompanyManagedDevicesIntegrations(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationsAPI.GetCompanyManagedDevicesIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyManagedDevicesIntegrations`: []ManagedDevicesIntegration
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationsAPI.GetCompanyManagedDevicesIntegrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyManagedDevicesIntegrationsRequest struct via the builder pattern


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

[**[]ManagedDevicesIntegration**](ManagedDevicesIntegration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyManagedDevicesIntegrationsById

> ManagedDevicesIntegration GetCompanyManagedDevicesIntegrationsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ManagedDevicesIntegration

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
	id := int32(56) // int32 | managedDevicesIntegrationId
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
	resp, r, err := apiClient.ManagedDevicesIntegrationsAPI.GetCompanyManagedDevicesIntegrationsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationsAPI.GetCompanyManagedDevicesIntegrationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyManagedDevicesIntegrationsById`: ManagedDevicesIntegration
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationsAPI.GetCompanyManagedDevicesIntegrationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyManagedDevicesIntegrationsByIdRequest struct via the builder pattern


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

[**ManagedDevicesIntegration**](ManagedDevicesIntegration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyManagedDevicesIntegrationsByIdUsages

> []Usage GetCompanyManagedDevicesIntegrationsByIdUsages(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of Usage Count

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
	id := int32(56) // int32 | managedDevicesIntegrationId
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
	resp, r, err := apiClient.ManagedDevicesIntegrationsAPI.GetCompanyManagedDevicesIntegrationsByIdUsages(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationsAPI.GetCompanyManagedDevicesIntegrationsByIdUsages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyManagedDevicesIntegrationsByIdUsages`: []Usage
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationsAPI.GetCompanyManagedDevicesIntegrationsByIdUsages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyManagedDevicesIntegrationsByIdUsagesRequest struct via the builder pattern


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

[**[]Usage**](Usage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyManagedDevicesIntegrationsByIdUsagesList

> []Usage GetCompanyManagedDevicesIntegrationsByIdUsagesList(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of Usage

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
	id := int32(56) // int32 | managedDevicesIntegrationId
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
	resp, r, err := apiClient.ManagedDevicesIntegrationsAPI.GetCompanyManagedDevicesIntegrationsByIdUsagesList(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationsAPI.GetCompanyManagedDevicesIntegrationsByIdUsagesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyManagedDevicesIntegrationsByIdUsagesList`: []Usage
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationsAPI.GetCompanyManagedDevicesIntegrationsByIdUsagesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyManagedDevicesIntegrationsByIdUsagesListRequest struct via the builder pattern


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

[**[]Usage**](Usage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyManagedDevicesIntegrationsCount

> Count GetCompanyManagedDevicesIntegrationsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of Usage

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
	resp, r, err := apiClient.ManagedDevicesIntegrationsAPI.GetCompanyManagedDevicesIntegrationsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationsAPI.GetCompanyManagedDevicesIntegrationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyManagedDevicesIntegrationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationsAPI.GetCompanyManagedDevicesIntegrationsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyManagedDevicesIntegrationsCountRequest struct via the builder pattern


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


## PatchCompanyManagedDevicesIntegrationsById

> ManagedDevicesIntegration PatchCompanyManagedDevicesIntegrationsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ManagedDevicesIntegration

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
	id := int32(56) // int32 | managedDevicesIntegrationId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedDevicesIntegrationsAPI.PatchCompanyManagedDevicesIntegrationsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationsAPI.PatchCompanyManagedDevicesIntegrationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyManagedDevicesIntegrationsById`: ManagedDevicesIntegration
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationsAPI.PatchCompanyManagedDevicesIntegrationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyManagedDevicesIntegrationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ManagedDevicesIntegration**](ManagedDevicesIntegration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyManagedDevicesIntegrations

> ManagedDevicesIntegration PostCompanyManagedDevicesIntegrations(ctx).ClientId(clientId).ManagedDevicesIntegration(managedDevicesIntegration).Execute()

Post ManagedDevicesIntegration

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
	managedDevicesIntegration := *openapiclient.NewManagedDevicesIntegration("Name_example", "Solution_example", "LoginBy_example", "DefaultBillingLevel_example") // ManagedDevicesIntegration | managedDevicesIntegration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedDevicesIntegrationsAPI.PostCompanyManagedDevicesIntegrations(context.Background()).ClientId(clientId).ManagedDevicesIntegration(managedDevicesIntegration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationsAPI.PostCompanyManagedDevicesIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyManagedDevicesIntegrations`: ManagedDevicesIntegration
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationsAPI.PostCompanyManagedDevicesIntegrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyManagedDevicesIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **managedDevicesIntegration** | [**ManagedDevicesIntegration**](ManagedDevicesIntegration.md) | managedDevicesIntegration | 

### Return type

[**ManagedDevicesIntegration**](ManagedDevicesIntegration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyManagedDevicesIntegrationsById

> ManagedDevicesIntegration PutCompanyManagedDevicesIntegrationsById(ctx, id).ClientId(clientId).ManagedDevicesIntegration(managedDevicesIntegration).Execute()

Put ManagedDevicesIntegration

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
	id := int32(56) // int32 | managedDevicesIntegrationId
	clientId := "clientId_example" // string | 
	managedDevicesIntegration := *openapiclient.NewManagedDevicesIntegration("Name_example", "Solution_example", "LoginBy_example", "DefaultBillingLevel_example") // ManagedDevicesIntegration | managedDevicesIntegration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedDevicesIntegrationsAPI.PutCompanyManagedDevicesIntegrationsById(context.Background(), id).ClientId(clientId).ManagedDevicesIntegration(managedDevicesIntegration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationsAPI.PutCompanyManagedDevicesIntegrationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyManagedDevicesIntegrationsById`: ManagedDevicesIntegration
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationsAPI.PutCompanyManagedDevicesIntegrationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyManagedDevicesIntegrationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **managedDevicesIntegration** | [**ManagedDevicesIntegration**](ManagedDevicesIntegration.md) | managedDevicesIntegration | 

### Return type

[**ManagedDevicesIntegration**](ManagedDevicesIntegration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

