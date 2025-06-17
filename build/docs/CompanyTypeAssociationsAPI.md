# \CompanyTypeAssociationsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyCompanyTypeAssociationsById**](CompanyTypeAssociationsAPI.md#DeleteCompanyCompanyTypeAssociationsById) | **Delete** /company/companyTypeAssociations/{id} | Delete CompanyTypeAssociation
[**GetCompanyCompanyTypeAssociations**](CompanyTypeAssociationsAPI.md#GetCompanyCompanyTypeAssociations) | **Get** /company/companyTypeAssociations | Get List of CompanyTypeAssociation
[**GetCompanyCompanyTypeAssociationsById**](CompanyTypeAssociationsAPI.md#GetCompanyCompanyTypeAssociationsById) | **Get** /company/companyTypeAssociations/{id} | Get CompanyTypeAssociation
[**GetCompanyCompanyTypeAssociationsCount**](CompanyTypeAssociationsAPI.md#GetCompanyCompanyTypeAssociationsCount) | **Get** /company/companyTypeAssociations/count | Get Count of CompanyTypeAssociation
[**PatchCompanyCompanyTypeAssociationsById**](CompanyTypeAssociationsAPI.md#PatchCompanyCompanyTypeAssociationsById) | **Patch** /company/companyTypeAssociations/{id} | Patch CompanyTypeAssociation
[**PostCompanyCompanyTypeAssociations**](CompanyTypeAssociationsAPI.md#PostCompanyCompanyTypeAssociations) | **Post** /company/companyTypeAssociations | Post CompanyTypeAssociation
[**PutCompanyCompanyTypeAssociationsById**](CompanyTypeAssociationsAPI.md#PutCompanyCompanyTypeAssociationsById) | **Put** /company/companyTypeAssociations/{id} | Put CompanyTypeAssociation



## DeleteCompanyCompanyTypeAssociationsById

> DeleteCompanyCompanyTypeAssociationsById(ctx, id).ClientId(clientId).Execute()

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
	id := int32(56) // int32 | companyTypeAssociationId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CompanyTypeAssociationsAPI.DeleteCompanyCompanyTypeAssociationsById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyTypeAssociationsAPI.DeleteCompanyCompanyTypeAssociationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | companyTypeAssociationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyCompanyTypeAssociationsByIdRequest struct via the builder pattern


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


## GetCompanyCompanyTypeAssociations

> []CompanyCompanyTypeAssociation GetCompanyCompanyTypeAssociations(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

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
	resp, r, err := apiClient.CompanyTypeAssociationsAPI.GetCompanyCompanyTypeAssociations(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyTypeAssociationsAPI.GetCompanyCompanyTypeAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCompanyTypeAssociations`: []CompanyCompanyTypeAssociation
	fmt.Fprintf(os.Stdout, "Response from `CompanyTypeAssociationsAPI.GetCompanyCompanyTypeAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCompanyTypeAssociationsRequest struct via the builder pattern


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

[**[]CompanyCompanyTypeAssociation**](CompanyCompanyTypeAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyCompanyTypeAssociationsById

> CompanyCompanyTypeAssociation GetCompanyCompanyTypeAssociationsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

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
	id := int32(56) // int32 | companyTypeAssociationId
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
	resp, r, err := apiClient.CompanyTypeAssociationsAPI.GetCompanyCompanyTypeAssociationsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyTypeAssociationsAPI.GetCompanyCompanyTypeAssociationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCompanyTypeAssociationsById`: CompanyCompanyTypeAssociation
	fmt.Fprintf(os.Stdout, "Response from `CompanyTypeAssociationsAPI.GetCompanyCompanyTypeAssociationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | companyTypeAssociationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCompanyTypeAssociationsByIdRequest struct via the builder pattern


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

[**CompanyCompanyTypeAssociation**](CompanyCompanyTypeAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyCompanyTypeAssociationsCount

> Count GetCompanyCompanyTypeAssociationsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

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
	resp, r, err := apiClient.CompanyTypeAssociationsAPI.GetCompanyCompanyTypeAssociationsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyTypeAssociationsAPI.GetCompanyCompanyTypeAssociationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCompanyTypeAssociationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `CompanyTypeAssociationsAPI.GetCompanyCompanyTypeAssociationsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCompanyTypeAssociationsCountRequest struct via the builder pattern


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


## PatchCompanyCompanyTypeAssociationsById

> CompanyCompanyTypeAssociation PatchCompanyCompanyTypeAssociationsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

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
	id := int32(56) // int32 | companyTypeAssociationId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyTypeAssociationsAPI.PatchCompanyCompanyTypeAssociationsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyTypeAssociationsAPI.PatchCompanyCompanyTypeAssociationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyCompanyTypeAssociationsById`: CompanyCompanyTypeAssociation
	fmt.Fprintf(os.Stdout, "Response from `CompanyTypeAssociationsAPI.PatchCompanyCompanyTypeAssociationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | companyTypeAssociationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyCompanyTypeAssociationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**CompanyCompanyTypeAssociation**](CompanyCompanyTypeAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyCompanyTypeAssociations

> CompanyCompanyTypeAssociation PostCompanyCompanyTypeAssociations(ctx).ClientId(clientId).CompanyCompanyTypeAssociation(companyCompanyTypeAssociation).Execute()

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
	clientId := "clientId_example" // string | 
	companyCompanyTypeAssociation := *openapiclient.NewCompanyCompanyTypeAssociation(*openapiclient.NewCompanyTypeReference(), *openapiclient.NewCompanyReference()) // CompanyCompanyTypeAssociation | companyTypeAssociation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyTypeAssociationsAPI.PostCompanyCompanyTypeAssociations(context.Background()).ClientId(clientId).CompanyCompanyTypeAssociation(companyCompanyTypeAssociation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyTypeAssociationsAPI.PostCompanyCompanyTypeAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyCompanyTypeAssociations`: CompanyCompanyTypeAssociation
	fmt.Fprintf(os.Stdout, "Response from `CompanyTypeAssociationsAPI.PostCompanyCompanyTypeAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyCompanyTypeAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **companyCompanyTypeAssociation** | [**CompanyCompanyTypeAssociation**](CompanyCompanyTypeAssociation.md) | companyTypeAssociation | 

### Return type

[**CompanyCompanyTypeAssociation**](CompanyCompanyTypeAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyCompanyTypeAssociationsById

> CompanyCompanyTypeAssociation PutCompanyCompanyTypeAssociationsById(ctx, id).ClientId(clientId).CompanyCompanyTypeAssociation(companyCompanyTypeAssociation).Execute()

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
	id := int32(56) // int32 | companyTypeAssociationId
	clientId := "clientId_example" // string | 
	companyCompanyTypeAssociation := *openapiclient.NewCompanyCompanyTypeAssociation(*openapiclient.NewCompanyTypeReference(), *openapiclient.NewCompanyReference()) // CompanyCompanyTypeAssociation | companyTypeAssociation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyTypeAssociationsAPI.PutCompanyCompanyTypeAssociationsById(context.Background(), id).ClientId(clientId).CompanyCompanyTypeAssociation(companyCompanyTypeAssociation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyTypeAssociationsAPI.PutCompanyCompanyTypeAssociationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyCompanyTypeAssociationsById`: CompanyCompanyTypeAssociation
	fmt.Fprintf(os.Stdout, "Response from `CompanyTypeAssociationsAPI.PutCompanyCompanyTypeAssociationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | companyTypeAssociationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyCompanyTypeAssociationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **companyCompanyTypeAssociation** | [**CompanyCompanyTypeAssociation**](CompanyCompanyTypeAssociation.md) | companyTypeAssociation | 

### Return type

[**CompanyCompanyTypeAssociation**](CompanyCompanyTypeAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

