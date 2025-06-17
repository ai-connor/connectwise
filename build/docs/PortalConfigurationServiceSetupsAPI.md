# \PortalConfigurationServiceSetupsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompanyPortalConfigurationsByParentIdServiceSetups**](PortalConfigurationServiceSetupsAPI.md#GetCompanyPortalConfigurationsByParentIdServiceSetups) | **Get** /company/portalConfigurations/{parentId}/serviceSetups | Get List of PortalConfigurationServiceSetup
[**GetCompanyPortalConfigurationsByParentIdServiceSetupsById**](PortalConfigurationServiceSetupsAPI.md#GetCompanyPortalConfigurationsByParentIdServiceSetupsById) | **Get** /company/portalConfigurations/{parentId}/serviceSetups/{id} | Get PortalConfigurationServiceSetup
[**GetCompanyPortalConfigurationsByParentIdServiceSetupsCount**](PortalConfigurationServiceSetupsAPI.md#GetCompanyPortalConfigurationsByParentIdServiceSetupsCount) | **Get** /company/portalConfigurations/{parentId}/serviceSetups/count | Get Count of PortalConfigurationServiceSetup
[**PatchCompanyPortalConfigurationsByParentIdServiceSetupsById**](PortalConfigurationServiceSetupsAPI.md#PatchCompanyPortalConfigurationsByParentIdServiceSetupsById) | **Patch** /company/portalConfigurations/{parentId}/serviceSetups/{id} | Patch PortalConfigurationServiceSetup
[**PutCompanyPortalConfigurationsByParentIdServiceSetupsById**](PortalConfigurationServiceSetupsAPI.md#PutCompanyPortalConfigurationsByParentIdServiceSetupsById) | **Put** /company/portalConfigurations/{parentId}/serviceSetups/{id} | Put PortalConfigurationServiceSetup



## GetCompanyPortalConfigurationsByParentIdServiceSetups

> []PortalConfigurationServiceSetup GetCompanyPortalConfigurationsByParentIdServiceSetups(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of PortalConfigurationServiceSetup

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
	parentId := int32(56) // int32 | portalConfigurationId
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
	resp, r, err := apiClient.PortalConfigurationServiceSetupsAPI.GetCompanyPortalConfigurationsByParentIdServiceSetups(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationServiceSetupsAPI.GetCompanyPortalConfigurationsByParentIdServiceSetups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyPortalConfigurationsByParentIdServiceSetups`: []PortalConfigurationServiceSetup
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationServiceSetupsAPI.GetCompanyPortalConfigurationsByParentIdServiceSetups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest struct via the builder pattern


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

[**[]PortalConfigurationServiceSetup**](PortalConfigurationServiceSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyPortalConfigurationsByParentIdServiceSetupsById

> PortalConfigurationServiceSetup GetCompanyPortalConfigurationsByParentIdServiceSetupsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get PortalConfigurationServiceSetup

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
	id := int32(56) // int32 | serviceSetupId
	parentId := int32(56) // int32 | portalConfigurationId
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
	resp, r, err := apiClient.PortalConfigurationServiceSetupsAPI.GetCompanyPortalConfigurationsByParentIdServiceSetupsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationServiceSetupsAPI.GetCompanyPortalConfigurationsByParentIdServiceSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyPortalConfigurationsByParentIdServiceSetupsById`: PortalConfigurationServiceSetup
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationServiceSetupsAPI.GetCompanyPortalConfigurationsByParentIdServiceSetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | serviceSetupId | 
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest struct via the builder pattern


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

[**PortalConfigurationServiceSetup**](PortalConfigurationServiceSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyPortalConfigurationsByParentIdServiceSetupsCount

> Count GetCompanyPortalConfigurationsByParentIdServiceSetupsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of PortalConfigurationServiceSetup

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
	parentId := int32(56) // int32 | portalConfigurationId
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
	resp, r, err := apiClient.PortalConfigurationServiceSetupsAPI.GetCompanyPortalConfigurationsByParentIdServiceSetupsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationServiceSetupsAPI.GetCompanyPortalConfigurationsByParentIdServiceSetupsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyPortalConfigurationsByParentIdServiceSetupsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationServiceSetupsAPI.GetCompanyPortalConfigurationsByParentIdServiceSetupsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest struct via the builder pattern


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


## PatchCompanyPortalConfigurationsByParentIdServiceSetupsById

> PortalConfigurationServiceSetup PatchCompanyPortalConfigurationsByParentIdServiceSetupsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch PortalConfigurationServiceSetup

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
	id := int32(56) // int32 | serviceSetupId
	parentId := int32(56) // int32 | portalConfigurationId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalConfigurationServiceSetupsAPI.PatchCompanyPortalConfigurationsByParentIdServiceSetupsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationServiceSetupsAPI.PatchCompanyPortalConfigurationsByParentIdServiceSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyPortalConfigurationsByParentIdServiceSetupsById`: PortalConfigurationServiceSetup
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationServiceSetupsAPI.PatchCompanyPortalConfigurationsByParentIdServiceSetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | serviceSetupId | 
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**PortalConfigurationServiceSetup**](PortalConfigurationServiceSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyPortalConfigurationsByParentIdServiceSetupsById

> PortalConfigurationServiceSetup PutCompanyPortalConfigurationsByParentIdServiceSetupsById(ctx, id, parentId).ClientId(clientId).PortalConfigurationServiceSetup(portalConfigurationServiceSetup).Execute()

Put PortalConfigurationServiceSetup

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
	id := int32(56) // int32 | serviceSetupId
	parentId := int32(56) // int32 | portalConfigurationId
	clientId := "clientId_example" // string | 
	portalConfigurationServiceSetup := *openapiclient.NewPortalConfigurationServiceSetup("DisplayClosedTicketsOption_example", *openapiclient.NewServiceSignoffReference(), *openapiclient.NewServiceSignoffReference()) // PortalConfigurationServiceSetup | portalConfigurationServiceSetup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalConfigurationServiceSetupsAPI.PutCompanyPortalConfigurationsByParentIdServiceSetupsById(context.Background(), id, parentId).ClientId(clientId).PortalConfigurationServiceSetup(portalConfigurationServiceSetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationServiceSetupsAPI.PutCompanyPortalConfigurationsByParentIdServiceSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyPortalConfigurationsByParentIdServiceSetupsById`: PortalConfigurationServiceSetup
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationServiceSetupsAPI.PutCompanyPortalConfigurationsByParentIdServiceSetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | serviceSetupId | 
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **portalConfigurationServiceSetup** | [**PortalConfigurationServiceSetup**](PortalConfigurationServiceSetup.md) | portalConfigurationServiceSetup | 

### Return type

[**PortalConfigurationServiceSetup**](PortalConfigurationServiceSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

