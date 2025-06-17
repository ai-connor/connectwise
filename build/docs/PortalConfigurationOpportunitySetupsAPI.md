# \PortalConfigurationOpportunitySetupsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompanyPortalConfigurationsByParentIdOpportunitySetups**](PortalConfigurationOpportunitySetupsAPI.md#GetCompanyPortalConfigurationsByParentIdOpportunitySetups) | **Get** /company/portalConfigurations/{parentId}/opportunitySetups | Get List of PortalConfigurationOpportunitySetup
[**GetCompanyPortalConfigurationsByParentIdOpportunitySetupsById**](PortalConfigurationOpportunitySetupsAPI.md#GetCompanyPortalConfigurationsByParentIdOpportunitySetupsById) | **Get** /company/portalConfigurations/{parentId}/opportunitySetups/{id} | Get PortalConfigurationOpportunitySetup
[**PatchCompanyPortalConfigurationsByParentIdOpportunitySetups**](PortalConfigurationOpportunitySetupsAPI.md#PatchCompanyPortalConfigurationsByParentIdOpportunitySetups) | **Patch** /company/portalConfigurations/{parentId}/opportunitySetups | Patch PortalConfigurationOpportunitySetup
[**PatchCompanyPortalConfigurationsByParentIdOpportunitySetupsById**](PortalConfigurationOpportunitySetupsAPI.md#PatchCompanyPortalConfigurationsByParentIdOpportunitySetupsById) | **Patch** /company/portalConfigurations/{parentId}/opportunitySetups/{id} | Patch PortalConfigurationOpportunitySetup
[**PutCompanyPortalConfigurationsByParentIdOpportunitySetups**](PortalConfigurationOpportunitySetupsAPI.md#PutCompanyPortalConfigurationsByParentIdOpportunitySetups) | **Put** /company/portalConfigurations/{parentId}/opportunitySetups | Put PortalConfigurationOpportunitySetup
[**PutCompanyPortalConfigurationsByParentIdOpportunitySetupsById**](PortalConfigurationOpportunitySetupsAPI.md#PutCompanyPortalConfigurationsByParentIdOpportunitySetupsById) | **Put** /company/portalConfigurations/{parentId}/opportunitySetups/{id} | Put PortalConfigurationOpportunitySetup



## GetCompanyPortalConfigurationsByParentIdOpportunitySetups

> []PortalConfigurationOpportunitySetup GetCompanyPortalConfigurationsByParentIdOpportunitySetups(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of PortalConfigurationOpportunitySetup

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
	resp, r, err := apiClient.PortalConfigurationOpportunitySetupsAPI.GetCompanyPortalConfigurationsByParentIdOpportunitySetups(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationOpportunitySetupsAPI.GetCompanyPortalConfigurationsByParentIdOpportunitySetups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyPortalConfigurationsByParentIdOpportunitySetups`: []PortalConfigurationOpportunitySetup
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationOpportunitySetupsAPI.GetCompanyPortalConfigurationsByParentIdOpportunitySetups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPortalConfigurationsByParentIdOpportunitySetupsRequest struct via the builder pattern


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

[**[]PortalConfigurationOpportunitySetup**](PortalConfigurationOpportunitySetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyPortalConfigurationsByParentIdOpportunitySetupsById

> PortalConfigurationOpportunitySetup GetCompanyPortalConfigurationsByParentIdOpportunitySetupsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get PortalConfigurationOpportunitySetup

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
	id := int32(56) // int32 | opportunitySetupId
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
	resp, r, err := apiClient.PortalConfigurationOpportunitySetupsAPI.GetCompanyPortalConfigurationsByParentIdOpportunitySetupsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationOpportunitySetupsAPI.GetCompanyPortalConfigurationsByParentIdOpportunitySetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyPortalConfigurationsByParentIdOpportunitySetupsById`: PortalConfigurationOpportunitySetup
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationOpportunitySetupsAPI.GetCompanyPortalConfigurationsByParentIdOpportunitySetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | opportunitySetupId | 
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPortalConfigurationsByParentIdOpportunitySetupsByIdRequest struct via the builder pattern


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

[**PortalConfigurationOpportunitySetup**](PortalConfigurationOpportunitySetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCompanyPortalConfigurationsByParentIdOpportunitySetups

> PortalConfigurationOpportunitySetup PatchCompanyPortalConfigurationsByParentIdOpportunitySetups(ctx, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch PortalConfigurationOpportunitySetup

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
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalConfigurationOpportunitySetupsAPI.PatchCompanyPortalConfigurationsByParentIdOpportunitySetups(context.Background(), parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationOpportunitySetupsAPI.PatchCompanyPortalConfigurationsByParentIdOpportunitySetups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyPortalConfigurationsByParentIdOpportunitySetups`: PortalConfigurationOpportunitySetup
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationOpportunitySetupsAPI.PatchCompanyPortalConfigurationsByParentIdOpportunitySetups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyPortalConfigurationsByParentIdOpportunitySetupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**PortalConfigurationOpportunitySetup**](PortalConfigurationOpportunitySetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCompanyPortalConfigurationsByParentIdOpportunitySetupsById

> PortalConfigurationOpportunitySetup PatchCompanyPortalConfigurationsByParentIdOpportunitySetupsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch PortalConfigurationOpportunitySetup

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
	id := int32(56) // int32 | opportunitySetupId
	parentId := int32(56) // int32 | portalConfigurationId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalConfigurationOpportunitySetupsAPI.PatchCompanyPortalConfigurationsByParentIdOpportunitySetupsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationOpportunitySetupsAPI.PatchCompanyPortalConfigurationsByParentIdOpportunitySetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyPortalConfigurationsByParentIdOpportunitySetupsById`: PortalConfigurationOpportunitySetup
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationOpportunitySetupsAPI.PatchCompanyPortalConfigurationsByParentIdOpportunitySetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | opportunitySetupId | 
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyPortalConfigurationsByParentIdOpportunitySetupsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**PortalConfigurationOpportunitySetup**](PortalConfigurationOpportunitySetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyPortalConfigurationsByParentIdOpportunitySetups

> PortalConfigurationOpportunitySetup PutCompanyPortalConfigurationsByParentIdOpportunitySetups(ctx, parentId).ClientId(clientId).PortalConfigurationOpportunitySetup(portalConfigurationOpportunitySetup).Execute()

Put PortalConfigurationOpportunitySetup

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
	portalConfigurationOpportunitySetup := *openapiclient.NewPortalConfigurationOpportunitySetup() // PortalConfigurationOpportunitySetup | opportunitySetup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalConfigurationOpportunitySetupsAPI.PutCompanyPortalConfigurationsByParentIdOpportunitySetups(context.Background(), parentId).ClientId(clientId).PortalConfigurationOpportunitySetup(portalConfigurationOpportunitySetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationOpportunitySetupsAPI.PutCompanyPortalConfigurationsByParentIdOpportunitySetups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyPortalConfigurationsByParentIdOpportunitySetups`: PortalConfigurationOpportunitySetup
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationOpportunitySetupsAPI.PutCompanyPortalConfigurationsByParentIdOpportunitySetups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyPortalConfigurationsByParentIdOpportunitySetupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **portalConfigurationOpportunitySetup** | [**PortalConfigurationOpportunitySetup**](PortalConfigurationOpportunitySetup.md) | opportunitySetup | 

### Return type

[**PortalConfigurationOpportunitySetup**](PortalConfigurationOpportunitySetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyPortalConfigurationsByParentIdOpportunitySetupsById

> PortalConfigurationOpportunitySetup PutCompanyPortalConfigurationsByParentIdOpportunitySetupsById(ctx, id, parentId).ClientId(clientId).PortalConfigurationOpportunitySetup(portalConfigurationOpportunitySetup).Execute()

Put PortalConfigurationOpportunitySetup

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
	id := int32(56) // int32 | opportunitySetupId
	parentId := int32(56) // int32 | portalConfigurationId
	clientId := "clientId_example" // string | 
	portalConfigurationOpportunitySetup := *openapiclient.NewPortalConfigurationOpportunitySetup() // PortalConfigurationOpportunitySetup | opportunitySetup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalConfigurationOpportunitySetupsAPI.PutCompanyPortalConfigurationsByParentIdOpportunitySetupsById(context.Background(), id, parentId).ClientId(clientId).PortalConfigurationOpportunitySetup(portalConfigurationOpportunitySetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalConfigurationOpportunitySetupsAPI.PutCompanyPortalConfigurationsByParentIdOpportunitySetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyPortalConfigurationsByParentIdOpportunitySetupsById`: PortalConfigurationOpportunitySetup
	fmt.Fprintf(os.Stdout, "Response from `PortalConfigurationOpportunitySetupsAPI.PutCompanyPortalConfigurationsByParentIdOpportunitySetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | opportunitySetupId | 
**parentId** | **int32** | portalConfigurationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyPortalConfigurationsByParentIdOpportunitySetupsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **portalConfigurationOpportunitySetup** | [**PortalConfigurationOpportunitySetup**](PortalConfigurationOpportunitySetup.md) | opportunitySetup | 

### Return type

[**PortalConfigurationOpportunitySetup**](PortalConfigurationOpportunitySetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

