# \PortalConfigurationProjectSetupsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompanyPortalConfigurationsByParentIdProjectSetups**](PortalConfigurationProjectSetupsAPI.md#GetCompanyPortalConfigurationsByParentIdProjectSetups) | **Get** /company/portalConfigurations/{parentId}/projectSetups | Get List of PortalConfigurationProjectSetup
[**GetCompanyPortalConfigurationsByParentIdProjectSetupsById**](PortalConfigurationProjectSetupsAPI.md#GetCompanyPortalConfigurationsByParentIdProjectSetupsById) | **Get** /company/portalConfigurations/{parentId}/projectSetups/{id} | Get PortalConfigurationProjectSetup
[**GetCompanyPortalConfigurationsByParentIdProjectSetupsCount**](PortalConfigurationProjectSetupsAPI.md#GetCompanyPortalConfigurationsByParentIdProjectSetupsCount) | **Get** /company/portalConfigurations/{parentId}/projectSetups/count | Get Count of PortalConfigurationProjectSetup
[**PatchCompanyPortalConfigurationsByParentIdProjectSetupsById**](PortalConfigurationProjectSetupsAPI.md#PatchCompanyPortalConfigurationsByParentIdProjectSetupsById) | **Patch** /company/portalConfigurations/{parentId}/projectSetups/{id} | Patch PortalConfigurationProjectSetup
[**PutCompanyPortalConfigurationsByParentIdProjectSetupsById**](PortalConfigurationProjectSetupsAPI.md#PutCompanyPortalConfigurationsByParentIdProjectSetupsById) | **Put** /company/portalConfigurations/{parentId}/projectSetups/{id} | Put PortalConfigurationProjectSetup



## GetCompanyPortalConfigurationsByParentIdProjectSetups

> []PortalConfigurationProjectSetup GetCompanyPortalConfigurationsByParentIdProjectSetups(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of PortalConfigurationProjectSetup

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
	resp, r, err := apiClient.PortalConfigurationProjectSetupsAPI.GetCompanyPortalConfigurationsByParentIdProjectSetups(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationProjectSetupsAPI.GetCompanyPortalConfigurationsByParentIdProjectSetups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyPortalConfigurationsByParentIdProjectSetups`: []PortalConfigurationProjectSetup
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationProjectSetupsAPI.GetCompanyPortalConfigurationsByParentIdProjectSetups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPortalConfigurationsByParentIdProjectSetupsRequest struct via the builder pattern


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

[**[]PortalConfigurationProjectSetup**](PortalConfigurationProjectSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyPortalConfigurationsByParentIdProjectSetupsById

> PortalConfigurationProjectSetup GetCompanyPortalConfigurationsByParentIdProjectSetupsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get PortalConfigurationProjectSetup

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
	id := int32(56) // int32 | projectSetupId
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
	resp, r, err := apiClient.PortalConfigurationProjectSetupsAPI.GetCompanyPortalConfigurationsByParentIdProjectSetupsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationProjectSetupsAPI.GetCompanyPortalConfigurationsByParentIdProjectSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyPortalConfigurationsByParentIdProjectSetupsById`: PortalConfigurationProjectSetup
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationProjectSetupsAPI.GetCompanyPortalConfigurationsByParentIdProjectSetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | projectSetupId | 
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPortalConfigurationsByParentIdProjectSetupsByIdRequest struct via the builder pattern


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

[**PortalConfigurationProjectSetup**](PortalConfigurationProjectSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyPortalConfigurationsByParentIdProjectSetupsCount

> Count GetCompanyPortalConfigurationsByParentIdProjectSetupsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of PortalConfigurationProjectSetup

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
	resp, r, err := apiClient.PortalConfigurationProjectSetupsAPI.GetCompanyPortalConfigurationsByParentIdProjectSetupsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationProjectSetupsAPI.GetCompanyPortalConfigurationsByParentIdProjectSetupsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyPortalConfigurationsByParentIdProjectSetupsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationProjectSetupsAPI.GetCompanyPortalConfigurationsByParentIdProjectSetupsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPortalConfigurationsByParentIdProjectSetupsCountRequest struct via the builder pattern


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


## PatchCompanyPortalConfigurationsByParentIdProjectSetupsById

> PortalConfigurationProjectSetup PatchCompanyPortalConfigurationsByParentIdProjectSetupsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch PortalConfigurationProjectSetup

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
	id := int32(56) // int32 | projectSetupId
	parentId := int32(56) // int32 | portalConfigurationId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalConfigurationProjectSetupsAPI.PatchCompanyPortalConfigurationsByParentIdProjectSetupsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationProjectSetupsAPI.PatchCompanyPortalConfigurationsByParentIdProjectSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyPortalConfigurationsByParentIdProjectSetupsById`: PortalConfigurationProjectSetup
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationProjectSetupsAPI.PatchCompanyPortalConfigurationsByParentIdProjectSetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | projectSetupId | 
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyPortalConfigurationsByParentIdProjectSetupsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**PortalConfigurationProjectSetup**](PortalConfigurationProjectSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyPortalConfigurationsByParentIdProjectSetupsById

> PortalConfigurationProjectSetup PutCompanyPortalConfigurationsByParentIdProjectSetupsById(ctx, id, parentId).ClientId(clientId).PortalConfigurationProjectSetup(portalConfigurationProjectSetup).Execute()

Put PortalConfigurationProjectSetup

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
	id := int32(56) // int32 | projectSetupId
	parentId := int32(56) // int32 | portalConfigurationId
	clientId := "clientId_example" // string | 
	portalConfigurationProjectSetup := *openapiclient.NewPortalConfigurationProjectSetup("OnlyDisplay_example") // PortalConfigurationProjectSetup | portalConfigurationProjectSetup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalConfigurationProjectSetupsAPI.PutCompanyPortalConfigurationsByParentIdProjectSetupsById(context.Background(), id, parentId).ClientId(clientId).PortalConfigurationProjectSetup(portalConfigurationProjectSetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationProjectSetupsAPI.PutCompanyPortalConfigurationsByParentIdProjectSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyPortalConfigurationsByParentIdProjectSetupsById`: PortalConfigurationProjectSetup
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationProjectSetupsAPI.PutCompanyPortalConfigurationsByParentIdProjectSetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | projectSetupId | 
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyPortalConfigurationsByParentIdProjectSetupsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **portalConfigurationProjectSetup** | [**PortalConfigurationProjectSetup**](PortalConfigurationProjectSetup.md) | portalConfigurationProjectSetup | 

### Return type

[**PortalConfigurationProjectSetup**](PortalConfigurationProjectSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

