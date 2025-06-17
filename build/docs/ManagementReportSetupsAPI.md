# \ManagementReportSetupsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompanyCompaniesByParentIdManagementReportSetup**](ManagementReportSetupsAPI.md#GetCompanyCompaniesByParentIdManagementReportSetup) | **Get** /company/companies/{parentId}/managementReportSetup | Get List of ManagementReportSetup
[**PatchCompanyCompaniesByParentIdManagementReportSetupById**](ManagementReportSetupsAPI.md#PatchCompanyCompaniesByParentIdManagementReportSetupById) | **Patch** /company/companies/{parentId}/managementReportSetup/{id} | Patch ManagementReportSetup
[**PostCompanyCompaniesByParentIdManagementReportSetup**](ManagementReportSetupsAPI.md#PostCompanyCompaniesByParentIdManagementReportSetup) | **Post** /company/companies/{parentId}/managementReportSetup | Post ManagementReportSetup
[**PutCompanyCompaniesByParentIdManagementReportSetupById**](ManagementReportSetupsAPI.md#PutCompanyCompaniesByParentIdManagementReportSetupById) | **Put** /company/companies/{parentId}/managementReportSetup/{id} | Put ManagementReportSetup



## GetCompanyCompaniesByParentIdManagementReportSetup

> []ManagementReportSetup GetCompanyCompaniesByParentIdManagementReportSetup(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ManagementReportSetup

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
	parentId := int32(56) // int32 | companyId
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
	resp, r, err := apiClient.ManagementReportSetupsAPI.GetCompanyCompaniesByParentIdManagementReportSetup(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementReportSetupsAPI.GetCompanyCompaniesByParentIdManagementReportSetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCompaniesByParentIdManagementReportSetup`: []ManagementReportSetup
	fmt.Fprintf(os.Stdout, "Response from `ManagementReportSetupsAPI.GetCompanyCompaniesByParentIdManagementReportSetup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCompaniesByParentIdManagementReportSetupRequest struct via the builder pattern


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

[**[]ManagementReportSetup**](ManagementReportSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCompanyCompaniesByParentIdManagementReportSetupById

> ManagementReportSetup PatchCompanyCompaniesByParentIdManagementReportSetupById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ManagementReportSetup

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
	id := int32(56) // int32 | managementReportSetupId
	parentId := int32(56) // int32 | companyId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementReportSetupsAPI.PatchCompanyCompaniesByParentIdManagementReportSetupById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementReportSetupsAPI.PatchCompanyCompaniesByParentIdManagementReportSetupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyCompaniesByParentIdManagementReportSetupById`: ManagementReportSetup
	fmt.Fprintf(os.Stdout, "Response from `ManagementReportSetupsAPI.PatchCompanyCompaniesByParentIdManagementReportSetupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementReportSetupId | 
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyCompaniesByParentIdManagementReportSetupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ManagementReportSetup**](ManagementReportSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyCompaniesByParentIdManagementReportSetup

> ManagementReportSetup PostCompanyCompaniesByParentIdManagementReportSetup(ctx, parentId).ClientId(clientId).ManagementReportSetup(managementReportSetup).Execute()

Post ManagementReportSetup

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
	parentId := int32(56) // int32 | companyId
	clientId := "clientId_example" // string | 
	managementReportSetup := *openapiclient.NewManagementReportSetup(false) // ManagementReportSetup | managementReportSetup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementReportSetupsAPI.PostCompanyCompaniesByParentIdManagementReportSetup(context.Background(), parentId).ClientId(clientId).ManagementReportSetup(managementReportSetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementReportSetupsAPI.PostCompanyCompaniesByParentIdManagementReportSetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyCompaniesByParentIdManagementReportSetup`: ManagementReportSetup
	fmt.Fprintf(os.Stdout, "Response from `ManagementReportSetupsAPI.PostCompanyCompaniesByParentIdManagementReportSetup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyCompaniesByParentIdManagementReportSetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **managementReportSetup** | [**ManagementReportSetup**](ManagementReportSetup.md) | managementReportSetup | 

### Return type

[**ManagementReportSetup**](ManagementReportSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyCompaniesByParentIdManagementReportSetupById

> ManagementReportSetup PutCompanyCompaniesByParentIdManagementReportSetupById(ctx, id, parentId).ClientId(clientId).ManagementReportSetup(managementReportSetup).Execute()

Put ManagementReportSetup

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
	id := int32(56) // int32 | managementReportSetupId
	parentId := int32(56) // int32 | companyId
	clientId := "clientId_example" // string | 
	managementReportSetup := *openapiclient.NewManagementReportSetup(false) // ManagementReportSetup | managementReportSetup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementReportSetupsAPI.PutCompanyCompaniesByParentIdManagementReportSetupById(context.Background(), id, parentId).ClientId(clientId).ManagementReportSetup(managementReportSetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementReportSetupsAPI.PutCompanyCompaniesByParentIdManagementReportSetupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyCompaniesByParentIdManagementReportSetupById`: ManagementReportSetup
	fmt.Fprintf(os.Stdout, "Response from `ManagementReportSetupsAPI.PutCompanyCompaniesByParentIdManagementReportSetupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementReportSetupId | 
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyCompaniesByParentIdManagementReportSetupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **managementReportSetup** | [**ManagementReportSetup**](ManagementReportSetup.md) | managementReportSetup | 

### Return type

[**ManagementReportSetup**](ManagementReportSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

