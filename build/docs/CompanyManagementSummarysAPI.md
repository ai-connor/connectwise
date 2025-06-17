# \CompanyManagementSummarysAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyCompaniesByParentIdManagementSummaryReportsById**](CompanyManagementSummarysAPI.md#DeleteCompanyCompaniesByParentIdManagementSummaryReportsById) | **Delete** /company/companies/{parentId}/managementSummaryReports/{id} | Delete CompanyManagementSummary
[**GetCompanyCompaniesByParentIdManagementSummaryReports**](CompanyManagementSummarysAPI.md#GetCompanyCompaniesByParentIdManagementSummaryReports) | **Get** /company/companies/{parentId}/managementSummaryReports | Get List of CompanyManagementSummary
[**GetCompanyCompaniesByParentIdManagementSummaryReportsById**](CompanyManagementSummarysAPI.md#GetCompanyCompaniesByParentIdManagementSummaryReportsById) | **Get** /company/companies/{parentId}/managementSummaryReports/{id} | Get CompanyManagementSummary
[**GetCompanyCompaniesByParentIdManagementSummaryReportsCount**](CompanyManagementSummarysAPI.md#GetCompanyCompaniesByParentIdManagementSummaryReportsCount) | **Get** /company/companies/{parentId}/managementSummaryReports/count | Get Count of CompanyManagementSummary
[**PatchCompanyCompaniesByParentIdManagementSummaryReportsById**](CompanyManagementSummarysAPI.md#PatchCompanyCompaniesByParentIdManagementSummaryReportsById) | **Patch** /company/companies/{parentId}/managementSummaryReports/{id} | Patch CompanyManagementSummary
[**PostCompanyCompaniesByParentIdManagementSummaryReports**](CompanyManagementSummarysAPI.md#PostCompanyCompaniesByParentIdManagementSummaryReports) | **Post** /company/companies/{parentId}/managementSummaryReports | Post CompanyManagementSummary
[**PutCompanyCompaniesByParentIdManagementSummaryReportsById**](CompanyManagementSummarysAPI.md#PutCompanyCompaniesByParentIdManagementSummaryReportsById) | **Put** /company/companies/{parentId}/managementSummaryReports/{id} | Put CompanyManagementSummary



## DeleteCompanyCompaniesByParentIdManagementSummaryReportsById

> DeleteCompanyCompaniesByParentIdManagementSummaryReportsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete CompanyManagementSummary

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
	id := int32(56) // int32 | managementSummaryReportId
	parentId := int32(56) // int32 | companyId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CompanyManagementSummarysAPI.DeleteCompanyCompaniesByParentIdManagementSummaryReportsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyManagementSummarysAPI.DeleteCompanyCompaniesByParentIdManagementSummaryReportsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementSummaryReportId | 
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyCompaniesByParentIdManagementSummaryReportsByIdRequest struct via the builder pattern


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


## GetCompanyCompaniesByParentIdManagementSummaryReports

> []CompanyManagementSummary GetCompanyCompaniesByParentIdManagementSummaryReports(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of CompanyManagementSummary

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
	resp, r, err := apiClient.CompanyManagementSummarysAPI.GetCompanyCompaniesByParentIdManagementSummaryReports(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyManagementSummarysAPI.GetCompanyCompaniesByParentIdManagementSummaryReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCompaniesByParentIdManagementSummaryReports`: []CompanyManagementSummary
	fmt.Fprintf(os.Stdout, "Response from `CompanyManagementSummarysAPI.GetCompanyCompaniesByParentIdManagementSummaryReports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCompaniesByParentIdManagementSummaryReportsRequest struct via the builder pattern


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

[**[]CompanyManagementSummary**](CompanyManagementSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyCompaniesByParentIdManagementSummaryReportsById

> CompanyManagementSummary GetCompanyCompaniesByParentIdManagementSummaryReportsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get CompanyManagementSummary

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
	id := int32(56) // int32 | managementSummaryReportId
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
	resp, r, err := apiClient.CompanyManagementSummarysAPI.GetCompanyCompaniesByParentIdManagementSummaryReportsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyManagementSummarysAPI.GetCompanyCompaniesByParentIdManagementSummaryReportsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCompaniesByParentIdManagementSummaryReportsById`: CompanyManagementSummary
	fmt.Fprintf(os.Stdout, "Response from `CompanyManagementSummarysAPI.GetCompanyCompaniesByParentIdManagementSummaryReportsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementSummaryReportId | 
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCompaniesByParentIdManagementSummaryReportsByIdRequest struct via the builder pattern


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

[**CompanyManagementSummary**](CompanyManagementSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyCompaniesByParentIdManagementSummaryReportsCount

> Count GetCompanyCompaniesByParentIdManagementSummaryReportsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of CompanyManagementSummary

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
	resp, r, err := apiClient.CompanyManagementSummarysAPI.GetCompanyCompaniesByParentIdManagementSummaryReportsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyManagementSummarysAPI.GetCompanyCompaniesByParentIdManagementSummaryReportsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCompaniesByParentIdManagementSummaryReportsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `CompanyManagementSummarysAPI.GetCompanyCompaniesByParentIdManagementSummaryReportsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCompaniesByParentIdManagementSummaryReportsCountRequest struct via the builder pattern


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


## PatchCompanyCompaniesByParentIdManagementSummaryReportsById

> CompanyManagementSummary PatchCompanyCompaniesByParentIdManagementSummaryReportsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch CompanyManagementSummary

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
	id := int32(56) // int32 | managementSummaryReportId
	parentId := int32(56) // int32 | companyId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyManagementSummarysAPI.PatchCompanyCompaniesByParentIdManagementSummaryReportsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyManagementSummarysAPI.PatchCompanyCompaniesByParentIdManagementSummaryReportsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyCompaniesByParentIdManagementSummaryReportsById`: CompanyManagementSummary
	fmt.Fprintf(os.Stdout, "Response from `CompanyManagementSummarysAPI.PatchCompanyCompaniesByParentIdManagementSummaryReportsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementSummaryReportId | 
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyCompaniesByParentIdManagementSummaryReportsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**CompanyManagementSummary**](CompanyManagementSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyCompaniesByParentIdManagementSummaryReports

> CompanyManagementSummary PostCompanyCompaniesByParentIdManagementSummaryReports(ctx, parentId).ClientId(clientId).CompanyManagementSummary(companyManagementSummary).Execute()

Post CompanyManagementSummary

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
	companyManagementSummary := *openapiclient.NewCompanyManagementSummary("GroupIdentifier_example") // CompanyManagementSummary | managementSummary

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyManagementSummarysAPI.PostCompanyCompaniesByParentIdManagementSummaryReports(context.Background(), parentId).ClientId(clientId).CompanyManagementSummary(companyManagementSummary).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyManagementSummarysAPI.PostCompanyCompaniesByParentIdManagementSummaryReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyCompaniesByParentIdManagementSummaryReports`: CompanyManagementSummary
	fmt.Fprintf(os.Stdout, "Response from `CompanyManagementSummarysAPI.PostCompanyCompaniesByParentIdManagementSummaryReports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyCompaniesByParentIdManagementSummaryReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **companyManagementSummary** | [**CompanyManagementSummary**](CompanyManagementSummary.md) | managementSummary | 

### Return type

[**CompanyManagementSummary**](CompanyManagementSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyCompaniesByParentIdManagementSummaryReportsById

> CompanyManagementSummary PutCompanyCompaniesByParentIdManagementSummaryReportsById(ctx, id, parentId).ClientId(clientId).CompanyManagementSummary(companyManagementSummary).Execute()

Put CompanyManagementSummary

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
	id := int32(56) // int32 | managementSummaryReportId
	parentId := int32(56) // int32 | companyId
	clientId := "clientId_example" // string | 
	companyManagementSummary := *openapiclient.NewCompanyManagementSummary("GroupIdentifier_example") // CompanyManagementSummary | managementSummary

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyManagementSummarysAPI.PutCompanyCompaniesByParentIdManagementSummaryReportsById(context.Background(), id, parentId).ClientId(clientId).CompanyManagementSummary(companyManagementSummary).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyManagementSummarysAPI.PutCompanyCompaniesByParentIdManagementSummaryReportsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyCompaniesByParentIdManagementSummaryReportsById`: CompanyManagementSummary
	fmt.Fprintf(os.Stdout, "Response from `CompanyManagementSummarysAPI.PutCompanyCompaniesByParentIdManagementSummaryReportsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementSummaryReportId | 
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyCompaniesByParentIdManagementSummaryReportsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **companyManagementSummary** | [**CompanyManagementSummary**](CompanyManagementSummary.md) | managementSummary | 

### Return type

[**CompanyManagementSummary**](CompanyManagementSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

