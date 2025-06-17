# \CompanyCompanyTypeAssociationsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyCompaniesByParentIdTypeAssociationsById**](CompanyCompanyTypeAssociationsAPI.md#DeleteCompanyCompaniesByParentIdTypeAssociationsById) | **Delete** /company/companies/{parentId}/typeAssociations/{id} | Delete CompanyTypeAssociation
[**GetCompanyCompaniesByParentIdTypeAssociations**](CompanyCompanyTypeAssociationsAPI.md#GetCompanyCompaniesByParentIdTypeAssociations) | **Get** /company/companies/{parentId}/typeAssociations | Get List of CompanyTypeAssociation
[**GetCompanyCompaniesByParentIdTypeAssociationsById**](CompanyCompanyTypeAssociationsAPI.md#GetCompanyCompaniesByParentIdTypeAssociationsById) | **Get** /company/companies/{parentId}/typeAssociations/{id} | Get CompanyTypeAssociation
[**GetCompanyCompaniesByParentIdTypeAssociationsCount**](CompanyCompanyTypeAssociationsAPI.md#GetCompanyCompaniesByParentIdTypeAssociationsCount) | **Get** /company/companies/{parentId}/typeAssociations/count | Get Count of CompanyTypeAssociation
[**PatchCompanyCompaniesByParentIdTypeAssociationsById**](CompanyCompanyTypeAssociationsAPI.md#PatchCompanyCompaniesByParentIdTypeAssociationsById) | **Patch** /company/companies/{parentId}/typeAssociations/{id} | Patch CompanyTypeAssociation
[**PostCompanyCompaniesByParentIdTypeAssociations**](CompanyCompanyTypeAssociationsAPI.md#PostCompanyCompaniesByParentIdTypeAssociations) | **Post** /company/companies/{parentId}/typeAssociations | Post CompanyTypeAssociation
[**PutCompanyCompaniesByParentIdTypeAssociationsById**](CompanyCompanyTypeAssociationsAPI.md#PutCompanyCompaniesByParentIdTypeAssociationsById) | **Put** /company/companies/{parentId}/typeAssociations/{id} | Put CompanyTypeAssociation



## DeleteCompanyCompaniesByParentIdTypeAssociationsById

> DeleteCompanyCompaniesByParentIdTypeAssociationsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete CompanyTypeAssociation

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
	id := int32(56) // int32 | typeAssociationId
	parentId := int32(56) // int32 | companyId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CompanyCompanyTypeAssociationsAPI.DeleteCompanyCompaniesByParentIdTypeAssociationsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyCompanyTypeAssociationsAPI.DeleteCompanyCompaniesByParentIdTypeAssociationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | typeAssociationId | 
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyCompaniesByParentIdTypeAssociationsByIdRequest struct via the builder pattern


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


## GetCompanyCompaniesByParentIdTypeAssociations

> []CompanyCompanyTypeAssociationCompanyTypeAssociation GetCompanyCompaniesByParentIdTypeAssociations(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of CompanyTypeAssociation

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
	resp, r, err := apiClient.CompanyCompanyTypeAssociationsAPI.GetCompanyCompaniesByParentIdTypeAssociations(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyCompanyTypeAssociationsAPI.GetCompanyCompaniesByParentIdTypeAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCompaniesByParentIdTypeAssociations`: []CompanyCompanyTypeAssociationCompanyTypeAssociation
	fmt.Fprintf(os.Stdout, "Response from `CompanyCompanyTypeAssociationsAPI.GetCompanyCompaniesByParentIdTypeAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCompaniesByParentIdTypeAssociationsRequest struct via the builder pattern


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

[**[]CompanyCompanyTypeAssociationCompanyTypeAssociation**](CompanyCompanyTypeAssociationCompanyTypeAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyCompaniesByParentIdTypeAssociationsById

> CompanyCompanyTypeAssociationCompanyTypeAssociation GetCompanyCompaniesByParentIdTypeAssociationsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get CompanyTypeAssociation

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
	id := int32(56) // int32 | typeAssociationId
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
	resp, r, err := apiClient.CompanyCompanyTypeAssociationsAPI.GetCompanyCompaniesByParentIdTypeAssociationsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyCompanyTypeAssociationsAPI.GetCompanyCompaniesByParentIdTypeAssociationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCompaniesByParentIdTypeAssociationsById`: CompanyCompanyTypeAssociationCompanyTypeAssociation
	fmt.Fprintf(os.Stdout, "Response from `CompanyCompanyTypeAssociationsAPI.GetCompanyCompaniesByParentIdTypeAssociationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | typeAssociationId | 
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCompaniesByParentIdTypeAssociationsByIdRequest struct via the builder pattern


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

[**CompanyCompanyTypeAssociationCompanyTypeAssociation**](CompanyCompanyTypeAssociationCompanyTypeAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyCompaniesByParentIdTypeAssociationsCount

> Count GetCompanyCompaniesByParentIdTypeAssociationsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of CompanyTypeAssociation

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
	resp, r, err := apiClient.CompanyCompanyTypeAssociationsAPI.GetCompanyCompaniesByParentIdTypeAssociationsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyCompanyTypeAssociationsAPI.GetCompanyCompaniesByParentIdTypeAssociationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCompaniesByParentIdTypeAssociationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `CompanyCompanyTypeAssociationsAPI.GetCompanyCompaniesByParentIdTypeAssociationsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCompaniesByParentIdTypeAssociationsCountRequest struct via the builder pattern


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


## PatchCompanyCompaniesByParentIdTypeAssociationsById

> CompanyCompanyTypeAssociationCompanyTypeAssociation PatchCompanyCompaniesByParentIdTypeAssociationsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch CompanyTypeAssociation

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
	id := int32(56) // int32 | typeAssociationId
	parentId := int32(56) // int32 | companyId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyCompanyTypeAssociationsAPI.PatchCompanyCompaniesByParentIdTypeAssociationsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyCompanyTypeAssociationsAPI.PatchCompanyCompaniesByParentIdTypeAssociationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyCompaniesByParentIdTypeAssociationsById`: CompanyCompanyTypeAssociationCompanyTypeAssociation
	fmt.Fprintf(os.Stdout, "Response from `CompanyCompanyTypeAssociationsAPI.PatchCompanyCompaniesByParentIdTypeAssociationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | typeAssociationId | 
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyCompaniesByParentIdTypeAssociationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**CompanyCompanyTypeAssociationCompanyTypeAssociation**](CompanyCompanyTypeAssociationCompanyTypeAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyCompaniesByParentIdTypeAssociations

> CompanyCompanyTypeAssociationCompanyTypeAssociation PostCompanyCompaniesByParentIdTypeAssociations(ctx, parentId).ClientId(clientId).CompanyCompanyTypeAssociationCompanyTypeAssociation(companyCompanyTypeAssociationCompanyTypeAssociation).Execute()

Post CompanyTypeAssociation

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
	companyCompanyTypeAssociationCompanyTypeAssociation := *openapiclient.NewCompanyCompanyTypeAssociationCompanyTypeAssociation(*openapiclient.NewCompanyTypeReference(), *openapiclient.NewCompanyReference()) // CompanyCompanyTypeAssociationCompanyTypeAssociation | companyTypeAssociation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyCompanyTypeAssociationsAPI.PostCompanyCompaniesByParentIdTypeAssociations(context.Background(), parentId).ClientId(clientId).CompanyCompanyTypeAssociationCompanyTypeAssociation(companyCompanyTypeAssociationCompanyTypeAssociation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyCompanyTypeAssociationsAPI.PostCompanyCompaniesByParentIdTypeAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyCompaniesByParentIdTypeAssociations`: CompanyCompanyTypeAssociationCompanyTypeAssociation
	fmt.Fprintf(os.Stdout, "Response from `CompanyCompanyTypeAssociationsAPI.PostCompanyCompaniesByParentIdTypeAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyCompaniesByParentIdTypeAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **companyCompanyTypeAssociationCompanyTypeAssociation** | [**CompanyCompanyTypeAssociationCompanyTypeAssociation**](CompanyCompanyTypeAssociationCompanyTypeAssociation.md) | companyTypeAssociation | 

### Return type

[**CompanyCompanyTypeAssociationCompanyTypeAssociation**](CompanyCompanyTypeAssociationCompanyTypeAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyCompaniesByParentIdTypeAssociationsById

> CompanyCompanyTypeAssociationCompanyTypeAssociation PutCompanyCompaniesByParentIdTypeAssociationsById(ctx, id, parentId).ClientId(clientId).CompanyCompanyTypeAssociationCompanyTypeAssociation(companyCompanyTypeAssociationCompanyTypeAssociation).Execute()

Put CompanyTypeAssociation

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
	id := int32(56) // int32 | typeAssociationId
	parentId := int32(56) // int32 | companyId
	clientId := "clientId_example" // string | 
	companyCompanyTypeAssociationCompanyTypeAssociation := *openapiclient.NewCompanyCompanyTypeAssociationCompanyTypeAssociation(*openapiclient.NewCompanyTypeReference(), *openapiclient.NewCompanyReference()) // CompanyCompanyTypeAssociationCompanyTypeAssociation | companyTypeAssociation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyCompanyTypeAssociationsAPI.PutCompanyCompaniesByParentIdTypeAssociationsById(context.Background(), id, parentId).ClientId(clientId).CompanyCompanyTypeAssociationCompanyTypeAssociation(companyCompanyTypeAssociationCompanyTypeAssociation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyCompanyTypeAssociationsAPI.PutCompanyCompaniesByParentIdTypeAssociationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyCompaniesByParentIdTypeAssociationsById`: CompanyCompanyTypeAssociationCompanyTypeAssociation
	fmt.Fprintf(os.Stdout, "Response from `CompanyCompanyTypeAssociationsAPI.PutCompanyCompaniesByParentIdTypeAssociationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | typeAssociationId | 
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyCompaniesByParentIdTypeAssociationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **companyCompanyTypeAssociationCompanyTypeAssociation** | [**CompanyCompanyTypeAssociationCompanyTypeAssociation**](CompanyCompanyTypeAssociationCompanyTypeAssociation.md) | companyTypeAssociation | 

### Return type

[**CompanyCompanyTypeAssociationCompanyTypeAssociation**](CompanyCompanyTypeAssociationCompanyTypeAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

