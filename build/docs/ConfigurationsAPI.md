# \ConfigurationsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyConfigurationsBulk**](ConfigurationsAPI.md#DeleteCompanyConfigurationsBulk) | **Delete** /company/configurations/bulk | Delete BulkResult
[**DeleteCompanyConfigurationsById**](ConfigurationsAPI.md#DeleteCompanyConfigurationsById) | **Delete** /company/configurations/{id} | Delete Configuration
[**GetCompanyConfigurations**](ConfigurationsAPI.md#GetCompanyConfigurations) | **Get** /company/configurations | Get List of Configuration
[**GetCompanyConfigurationsById**](ConfigurationsAPI.md#GetCompanyConfigurationsById) | **Get** /company/configurations/{id} | Get Configuration
[**GetCompanyConfigurationsByIdQuickAccessCount**](ConfigurationsAPI.md#GetCompanyConfigurationsByIdQuickAccessCount) | **Get** /company/configurations/{id}/quickAccess/count | Get Configuration Tab Count
[**GetCompanyConfigurationsCount**](ConfigurationsAPI.md#GetCompanyConfigurationsCount) | **Get** /company/configurations/count | Get Count of Configuration
[**PatchCompanyConfigurationsById**](ConfigurationsAPI.md#PatchCompanyConfigurationsById) | **Patch** /company/configurations/{id} | Patch Configuration
[**PatchCompanyConfigurationsByIdChangeType**](ConfigurationsAPI.md#PatchCompanyConfigurationsByIdChangeType) | **Patch** /company/configurations/{id}/changeType | Patch Configuration
[**PostCompanyConfigurations**](ConfigurationsAPI.md#PostCompanyConfigurations) | **Post** /company/configurations | Post Configuration
[**PostCompanyConfigurationsBulk**](ConfigurationsAPI.md#PostCompanyConfigurationsBulk) | **Post** /company/configurations/bulk | Post Configuration
[**PutCompanyConfigurationsBulk**](ConfigurationsAPI.md#PutCompanyConfigurationsBulk) | **Put** /company/configurations/bulk | Put Configuration
[**PutCompanyConfigurationsById**](ConfigurationsAPI.md#PutCompanyConfigurationsById) | **Put** /company/configurations/{id} | Put Configuration



## DeleteCompanyConfigurationsBulk

> BulkResult DeleteCompanyConfigurationsBulk(ctx).ClientId(clientId).Execute()

Delete BulkResult

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationsAPI.DeleteCompanyConfigurationsBulk(context.Background()).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.DeleteCompanyConfigurationsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCompanyConfigurationsBulk`: BulkResult
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.DeleteCompanyConfigurationsBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyConfigurationsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 

### Return type

[**BulkResult**](BulkResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCompanyConfigurationsById

> DeleteCompanyConfigurationsById(ctx, id).ClientId(clientId).Execute()

Delete Configuration

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
	id := int32(56) // int32 | configurationId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationsAPI.DeleteCompanyConfigurationsById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.DeleteCompanyConfigurationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | configurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyConfigurationsByIdRequest struct via the builder pattern


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


## GetCompanyConfigurations

> []CompanyConfiguration GetCompanyConfigurations(ctx).ClientId(clientId).ManagedIdentifier(managedIdentifier).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of Configuration

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
	managedIdentifier := "managedIdentifier_example" // string | managedIdentifier (optional)
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
	resp, r, err := apiClient.ConfigurationsAPI.GetCompanyConfigurations(context.Background()).ClientId(clientId).ManagedIdentifier(managedIdentifier).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.GetCompanyConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyConfigurations`: []CompanyConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.GetCompanyConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **managedIdentifier** | **string** | managedIdentifier | 
 **conditions** | **string** |  | 
 **childConditions** | **string** |  | 
 **customFieldConditions** | **string** |  | 
 **orderBy** | **string** |  | 
 **fields** | **string** |  | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **pageId** | **int32** |  | 

### Return type

[**[]CompanyConfiguration**](CompanyConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyConfigurationsById

> CompanyConfiguration GetCompanyConfigurationsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Configuration

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
	id := int32(56) // int32 | configurationId
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
	resp, r, err := apiClient.ConfigurationsAPI.GetCompanyConfigurationsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.GetCompanyConfigurationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyConfigurationsById`: CompanyConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.GetCompanyConfigurationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | configurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyConfigurationsByIdRequest struct via the builder pattern


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

[**CompanyConfiguration**](CompanyConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyConfigurationsByIdQuickAccessCount

> map[string]interface{} GetCompanyConfigurationsByIdQuickAccessCount(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Configuration Tab Count

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
	id := int32(56) // int32 | configurationId
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
	resp, r, err := apiClient.ConfigurationsAPI.GetCompanyConfigurationsByIdQuickAccessCount(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.GetCompanyConfigurationsByIdQuickAccessCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyConfigurationsByIdQuickAccessCount`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.GetCompanyConfigurationsByIdQuickAccessCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | configurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyConfigurationsByIdQuickAccessCountRequest struct via the builder pattern


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

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyConfigurationsCount

> Count GetCompanyConfigurationsCount(ctx).ClientId(clientId).ManagedIdentifier(managedIdentifier).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of Configuration

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
	managedIdentifier := "managedIdentifier_example" // string | managedIdentifier (optional)
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
	resp, r, err := apiClient.ConfigurationsAPI.GetCompanyConfigurationsCount(context.Background()).ClientId(clientId).ManagedIdentifier(managedIdentifier).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.GetCompanyConfigurationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyConfigurationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.GetCompanyConfigurationsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyConfigurationsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **managedIdentifier** | **string** | managedIdentifier | 
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


## PatchCompanyConfigurationsById

> CompanyConfiguration PatchCompanyConfigurationsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).ManagedInformation(managedInformation).Execute()

Patch Configuration

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
	id := int32(56) // int32 | configurationId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation
	managedInformation := *openapiclient.NewManagedInformation() // ManagedInformation | managedInformation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationsAPI.PatchCompanyConfigurationsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).ManagedInformation(managedInformation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.PatchCompanyConfigurationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyConfigurationsById`: CompanyConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.PatchCompanyConfigurationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | configurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyConfigurationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 
 **managedInformation** | [**ManagedInformation**](ManagedInformation.md) | managedInformation | 

### Return type

[**CompanyConfiguration**](CompanyConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCompanyConfigurationsByIdChangeType

> CompanyConfiguration PatchCompanyConfigurationsByIdChangeType(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch Configuration

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
	id := int32(56) // int32 | configurationId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationsAPI.PatchCompanyConfigurationsByIdChangeType(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.PatchCompanyConfigurationsByIdChangeType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyConfigurationsByIdChangeType`: CompanyConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.PatchCompanyConfigurationsByIdChangeType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | configurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyConfigurationsByIdChangeTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**CompanyConfiguration**](CompanyConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyConfigurations

> CompanyConfiguration PostCompanyConfigurations(ctx).ClientId(clientId).CompanyConfiguration(companyConfiguration).ManagedInformation(managedInformation).Execute()

Post Configuration

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
	companyConfiguration := *openapiclient.NewCompanyConfiguration("Name_example", *openapiclient.NewConfigurationTypeReference(), *openapiclient.NewCompanyReference()) // CompanyConfiguration | configuration
	managedInformation := *openapiclient.NewManagedInformation() // ManagedInformation | managedInformation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationsAPI.PostCompanyConfigurations(context.Background()).ClientId(clientId).CompanyConfiguration(companyConfiguration).ManagedInformation(managedInformation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.PostCompanyConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyConfigurations`: CompanyConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.PostCompanyConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **companyConfiguration** | [**CompanyConfiguration**](CompanyConfiguration.md) | configuration | 
 **managedInformation** | [**ManagedInformation**](ManagedInformation.md) | managedInformation | 

### Return type

[**CompanyConfiguration**](CompanyConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyConfigurationsBulk

> CompanyConfiguration PostCompanyConfigurationsBulk(ctx).ClientId(clientId).CompanyConfiguration(companyConfiguration).ManagedInformation(managedInformation).Execute()

Post Configuration

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
	companyConfiguration := []openapiclient.CompanyConfiguration{*openapiclient.NewCompanyConfiguration("Name_example", *openapiclient.NewConfigurationTypeReference(), *openapiclient.NewCompanyReference())} // []CompanyConfiguration | List of Configuration
	managedInformation := *openapiclient.NewManagedInformation() // ManagedInformation | managedInformation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationsAPI.PostCompanyConfigurationsBulk(context.Background()).ClientId(clientId).CompanyConfiguration(companyConfiguration).ManagedInformation(managedInformation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.PostCompanyConfigurationsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyConfigurationsBulk`: CompanyConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.PostCompanyConfigurationsBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyConfigurationsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **companyConfiguration** | [**[]CompanyConfiguration**](CompanyConfiguration.md) | List of Configuration | 
 **managedInformation** | [**ManagedInformation**](ManagedInformation.md) | managedInformation | 

### Return type

[**CompanyConfiguration**](CompanyConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyConfigurationsBulk

> CompanyConfiguration PutCompanyConfigurationsBulk(ctx).ClientId(clientId).CompanyConfiguration(companyConfiguration).ManagedInformation(managedInformation).Execute()

Put Configuration

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
	companyConfiguration := []openapiclient.CompanyConfiguration{*openapiclient.NewCompanyConfiguration("Name_example", *openapiclient.NewConfigurationTypeReference(), *openapiclient.NewCompanyReference())} // []CompanyConfiguration | List of Configuration
	managedInformation := *openapiclient.NewManagedInformation() // ManagedInformation | managedInformation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationsAPI.PutCompanyConfigurationsBulk(context.Background()).ClientId(clientId).CompanyConfiguration(companyConfiguration).ManagedInformation(managedInformation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.PutCompanyConfigurationsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyConfigurationsBulk`: CompanyConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.PutCompanyConfigurationsBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyConfigurationsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **companyConfiguration** | [**[]CompanyConfiguration**](CompanyConfiguration.md) | List of Configuration | 
 **managedInformation** | [**ManagedInformation**](ManagedInformation.md) | managedInformation | 

### Return type

[**CompanyConfiguration**](CompanyConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyConfigurationsById

> CompanyConfiguration PutCompanyConfigurationsById(ctx, id).ClientId(clientId).CompanyConfiguration(companyConfiguration).ManagedInformation(managedInformation).Execute()

Put Configuration

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
	id := int32(56) // int32 | configurationId
	clientId := "clientId_example" // string | 
	companyConfiguration := *openapiclient.NewCompanyConfiguration("Name_example", *openapiclient.NewConfigurationTypeReference(), *openapiclient.NewCompanyReference()) // CompanyConfiguration | configuration
	managedInformation := *openapiclient.NewManagedInformation() // ManagedInformation | managedInformation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationsAPI.PutCompanyConfigurationsById(context.Background(), id).ClientId(clientId).CompanyConfiguration(companyConfiguration).ManagedInformation(managedInformation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationsAPI.PutCompanyConfigurationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyConfigurationsById`: CompanyConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationsAPI.PutCompanyConfigurationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | configurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyConfigurationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **companyConfiguration** | [**CompanyConfiguration**](CompanyConfiguration.md) | configuration | 
 **managedInformation** | [**ManagedInformation**](ManagedInformation.md) | managedInformation | 

### Return type

[**CompanyConfiguration**](CompanyConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

