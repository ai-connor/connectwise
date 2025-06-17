# \PortalConfigurationsInvoiceSetupsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompanyPortalConfigurationsByParentIdInvoiceSetups**](PortalConfigurationsInvoiceSetupsAPI.md#GetCompanyPortalConfigurationsByParentIdInvoiceSetups) | **Get** /company/portalConfigurations/{parentId}/invoiceSetups | Get List of PortalConfigurationInvoiceSetup
[**GetCompanyPortalConfigurationsByParentIdInvoiceSetupsById**](PortalConfigurationsInvoiceSetupsAPI.md#GetCompanyPortalConfigurationsByParentIdInvoiceSetupsById) | **Get** /company/portalConfigurations/{parentId}/invoiceSetups/{id} | Get PortalConfigurationInvoiceSetup
[**GetCompanyPortalConfigurationsByParentIdInvoiceSetupsCount**](PortalConfigurationsInvoiceSetupsAPI.md#GetCompanyPortalConfigurationsByParentIdInvoiceSetupsCount) | **Get** /company/portalConfigurations/{parentId}/invoiceSetups/count | Get Count of PortalConfigurationInvoiceSetup
[**PatchCompanyPortalConfigurationsByParentIdInvoiceSetupsById**](PortalConfigurationsInvoiceSetupsAPI.md#PatchCompanyPortalConfigurationsByParentIdInvoiceSetupsById) | **Patch** /company/portalConfigurations/{parentId}/invoiceSetups/{id} | Patch PortalConfigurationInvoiceSetup
[**PostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransaction**](PortalConfigurationsInvoiceSetupsAPI.md#PostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransaction) | **Post** /company/portalConfigurations/{parentId}/invoiceSetups/{id}/testTransaction | Post SuccessResponse
[**PutCompanyPortalConfigurationsByParentIdInvoiceSetupsById**](PortalConfigurationsInvoiceSetupsAPI.md#PutCompanyPortalConfigurationsByParentIdInvoiceSetupsById) | **Put** /company/portalConfigurations/{parentId}/invoiceSetups/{id} | Put PortalConfigurationInvoiceSetup



## GetCompanyPortalConfigurationsByParentIdInvoiceSetups

> []PortalConfigurationInvoiceSetup GetCompanyPortalConfigurationsByParentIdInvoiceSetups(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of PortalConfigurationInvoiceSetup

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
	resp, r, err := apiClient.PortalConfigurationsInvoiceSetupsAPI.GetCompanyPortalConfigurationsByParentIdInvoiceSetups(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationsInvoiceSetupsAPI.GetCompanyPortalConfigurationsByParentIdInvoiceSetups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyPortalConfigurationsByParentIdInvoiceSetups`: []PortalConfigurationInvoiceSetup
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationsInvoiceSetupsAPI.GetCompanyPortalConfigurationsByParentIdInvoiceSetups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest struct via the builder pattern


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

[**[]PortalConfigurationInvoiceSetup**](PortalConfigurationInvoiceSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyPortalConfigurationsByParentIdInvoiceSetupsById

> PortalConfigurationInvoiceSetup GetCompanyPortalConfigurationsByParentIdInvoiceSetupsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get PortalConfigurationInvoiceSetup

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
	id := int32(56) // int32 | invoiceSetupId
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
	resp, r, err := apiClient.PortalConfigurationsInvoiceSetupsAPI.GetCompanyPortalConfigurationsByParentIdInvoiceSetupsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationsInvoiceSetupsAPI.GetCompanyPortalConfigurationsByParentIdInvoiceSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyPortalConfigurationsByParentIdInvoiceSetupsById`: PortalConfigurationInvoiceSetup
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationsInvoiceSetupsAPI.GetCompanyPortalConfigurationsByParentIdInvoiceSetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | invoiceSetupId | 
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest struct via the builder pattern


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

[**PortalConfigurationInvoiceSetup**](PortalConfigurationInvoiceSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyPortalConfigurationsByParentIdInvoiceSetupsCount

> Count GetCompanyPortalConfigurationsByParentIdInvoiceSetupsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of PortalConfigurationInvoiceSetup

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
	resp, r, err := apiClient.PortalConfigurationsInvoiceSetupsAPI.GetCompanyPortalConfigurationsByParentIdInvoiceSetupsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationsInvoiceSetupsAPI.GetCompanyPortalConfigurationsByParentIdInvoiceSetupsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyPortalConfigurationsByParentIdInvoiceSetupsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationsInvoiceSetupsAPI.GetCompanyPortalConfigurationsByParentIdInvoiceSetupsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest struct via the builder pattern


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


## PatchCompanyPortalConfigurationsByParentIdInvoiceSetupsById

> PortalConfigurationInvoiceSetup PatchCompanyPortalConfigurationsByParentIdInvoiceSetupsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch PortalConfigurationInvoiceSetup

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
	id := int32(56) // int32 | invoiceSetupId
	parentId := int32(56) // int32 | portalConfigurationId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalConfigurationsInvoiceSetupsAPI.PatchCompanyPortalConfigurationsByParentIdInvoiceSetupsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationsInvoiceSetupsAPI.PatchCompanyPortalConfigurationsByParentIdInvoiceSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyPortalConfigurationsByParentIdInvoiceSetupsById`: PortalConfigurationInvoiceSetup
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationsInvoiceSetupsAPI.PatchCompanyPortalConfigurationsByParentIdInvoiceSetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | invoiceSetupId | 
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**PortalConfigurationInvoiceSetup**](PortalConfigurationInvoiceSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransaction

> SuccessResponse PostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransaction(ctx, id, parentId).ClientId(clientId).PortalConfigurationInvoiceSetup(portalConfigurationInvoiceSetup).Execute()

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
	id := int32(56) // int32 | invoiceSetupId
	parentId := int32(56) // int32 | portalConfigurationId
	clientId := "clientId_example" // string | 
	portalConfigurationInvoiceSetup := *openapiclient.NewPortalConfigurationInvoiceSetup() // PortalConfigurationInvoiceSetup | portalConfigurationInvoiceSetup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalConfigurationsInvoiceSetupsAPI.PostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransaction(context.Background(), id, parentId).ClientId(clientId).PortalConfigurationInvoiceSetup(portalConfigurationInvoiceSetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationsInvoiceSetupsAPI.PostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransaction`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationsInvoiceSetupsAPI.PostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | invoiceSetupId | 
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **portalConfigurationInvoiceSetup** | [**PortalConfigurationInvoiceSetup**](PortalConfigurationInvoiceSetup.md) | portalConfigurationInvoiceSetup | 

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


## PutCompanyPortalConfigurationsByParentIdInvoiceSetupsById

> PortalConfigurationInvoiceSetup PutCompanyPortalConfigurationsByParentIdInvoiceSetupsById(ctx, id, parentId).ClientId(clientId).PortalConfigurationInvoiceSetup(portalConfigurationInvoiceSetup).Execute()

Put PortalConfigurationInvoiceSetup

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
	id := int32(56) // int32 | invoiceSetupId
	parentId := int32(56) // int32 | portalConfigurationId
	clientId := "clientId_example" // string | 
	portalConfigurationInvoiceSetup := *openapiclient.NewPortalConfigurationInvoiceSetup() // PortalConfigurationInvoiceSetup | portalConfigurationInvoiceSetup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalConfigurationsInvoiceSetupsAPI.PutCompanyPortalConfigurationsByParentIdInvoiceSetupsById(context.Background(), id, parentId).ClientId(clientId).PortalConfigurationInvoiceSetup(portalConfigurationInvoiceSetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationsInvoiceSetupsAPI.PutCompanyPortalConfigurationsByParentIdInvoiceSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyPortalConfigurationsByParentIdInvoiceSetupsById`: PortalConfigurationInvoiceSetup
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationsInvoiceSetupsAPI.PutCompanyPortalConfigurationsByParentIdInvoiceSetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | invoiceSetupId | 
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **portalConfigurationInvoiceSetup** | [**PortalConfigurationInvoiceSetup**](PortalConfigurationInvoiceSetup.md) | portalConfigurationInvoiceSetup | 

### Return type

[**PortalConfigurationInvoiceSetup**](PortalConfigurationInvoiceSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

