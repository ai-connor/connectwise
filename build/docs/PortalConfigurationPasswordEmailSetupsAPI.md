# \PortalConfigurationPasswordEmailSetupsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompanyPortalConfigurationsByParentIdPasswordEmailSetups**](PortalConfigurationPasswordEmailSetupsAPI.md#GetCompanyPortalConfigurationsByParentIdPasswordEmailSetups) | **Get** /company/portalConfigurations/{parentId}/passwordEmailSetups | Get List of PortalConfigurationPasswordEmailSetup
[**GetCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById**](PortalConfigurationPasswordEmailSetupsAPI.md#GetCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById) | **Get** /company/portalConfigurations/{parentId}/passwordEmailSetups/{id} | Get PortalConfigurationPasswordEmailSetup
[**PatchCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById**](PortalConfigurationPasswordEmailSetupsAPI.md#PatchCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById) | **Patch** /company/portalConfigurations/{parentId}/passwordEmailSetups/{id} | Patch PortalConfigurationPasswordEmailSetup
[**PutCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById**](PortalConfigurationPasswordEmailSetupsAPI.md#PutCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById) | **Put** /company/portalConfigurations/{parentId}/passwordEmailSetups/{id} | Put PortalConfigurationPasswordEmailSetup



## GetCompanyPortalConfigurationsByParentIdPasswordEmailSetups

> []PortalConfigurationPasswordEmailSetup GetCompanyPortalConfigurationsByParentIdPasswordEmailSetups(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of PortalConfigurationPasswordEmailSetup

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
	resp, r, err := apiClient.PortalConfigurationPasswordEmailSetupsAPI.GetCompanyPortalConfigurationsByParentIdPasswordEmailSetups(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationPasswordEmailSetupsAPI.GetCompanyPortalConfigurationsByParentIdPasswordEmailSetups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyPortalConfigurationsByParentIdPasswordEmailSetups`: []PortalConfigurationPasswordEmailSetup
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationPasswordEmailSetupsAPI.GetCompanyPortalConfigurationsByParentIdPasswordEmailSetups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPortalConfigurationsByParentIdPasswordEmailSetupsRequest struct via the builder pattern


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

[**[]PortalConfigurationPasswordEmailSetup**](PortalConfigurationPasswordEmailSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById

> PortalConfigurationPasswordEmailSetup GetCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get PortalConfigurationPasswordEmailSetup

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
	id := int32(56) // int32 | passwordEmailSetupId
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
	resp, r, err := apiClient.PortalConfigurationPasswordEmailSetupsAPI.GetCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationPasswordEmailSetupsAPI.GetCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById`: PortalConfigurationPasswordEmailSetup
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationPasswordEmailSetupsAPI.GetCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | passwordEmailSetupId | 
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPortalConfigurationsByParentIdPasswordEmailSetupsByIdRequest struct via the builder pattern


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

[**PortalConfigurationPasswordEmailSetup**](PortalConfigurationPasswordEmailSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById

> PortalConfigurationPasswordEmailSetup PatchCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch PortalConfigurationPasswordEmailSetup

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
	id := int32(56) // int32 | passwordEmailSetupId
	parentId := int32(56) // int32 | portalConfigurationId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalConfigurationPasswordEmailSetupsAPI.PatchCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationPasswordEmailSetupsAPI.PatchCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById`: PortalConfigurationPasswordEmailSetup
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationPasswordEmailSetupsAPI.PatchCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | passwordEmailSetupId | 
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyPortalConfigurationsByParentIdPasswordEmailSetupsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**PortalConfigurationPasswordEmailSetup**](PortalConfigurationPasswordEmailSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById

> PortalConfigurationPasswordEmailSetup PutCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById(ctx, id, parentId).ClientId(clientId).PortalConfigurationPasswordEmailSetup(portalConfigurationPasswordEmailSetup).Execute()

Put PortalConfigurationPasswordEmailSetup

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
	id := int32(56) // int32 | passwordEmailSetupId
	parentId := int32(56) // int32 | portalConfigurationId
	clientId := "clientId_example" // string | 
	portalConfigurationPasswordEmailSetup := *openapiclient.NewPortalConfigurationPasswordEmailSetup() // PortalConfigurationPasswordEmailSetup | passwordEmailSetup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalConfigurationPasswordEmailSetupsAPI.PutCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById(context.Background(), id, parentId).ClientId(clientId).PortalConfigurationPasswordEmailSetup(portalConfigurationPasswordEmailSetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationPasswordEmailSetupsAPI.PutCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById`: PortalConfigurationPasswordEmailSetup
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationPasswordEmailSetupsAPI.PutCompanyPortalConfigurationsByParentIdPasswordEmailSetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | passwordEmailSetupId | 
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyPortalConfigurationsByParentIdPasswordEmailSetupsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **portalConfigurationPasswordEmailSetup** | [**PortalConfigurationPasswordEmailSetup**](PortalConfigurationPasswordEmailSetup.md) | passwordEmailSetup | 

### Return type

[**PortalConfigurationPasswordEmailSetup**](PortalConfigurationPasswordEmailSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

