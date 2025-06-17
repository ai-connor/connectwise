# \ContactTypeAssociationsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyContactTypeAssociationsById**](ContactTypeAssociationsAPI.md#DeleteCompanyContactTypeAssociationsById) | **Delete** /company/contactTypeAssociations/{id} | Delete ContactTypeAssociation
[**GetCompanyContactTypeAssociations**](ContactTypeAssociationsAPI.md#GetCompanyContactTypeAssociations) | **Get** /company/contactTypeAssociations | Get List of ContactTypeAssociation
[**GetCompanyContactTypeAssociationsById**](ContactTypeAssociationsAPI.md#GetCompanyContactTypeAssociationsById) | **Get** /company/contactTypeAssociations/{id} | Get ContactTypeAssociation
[**GetCompanyContactTypeAssociationsCount**](ContactTypeAssociationsAPI.md#GetCompanyContactTypeAssociationsCount) | **Get** /company/contactTypeAssociations/count | Get Count of ContactTypeAssociation
[**PatchCompanyContactTypeAssociationsById**](ContactTypeAssociationsAPI.md#PatchCompanyContactTypeAssociationsById) | **Patch** /company/contactTypeAssociations/{id} | Patch ContactTypeAssociation
[**PostCompanyContactTypeAssociations**](ContactTypeAssociationsAPI.md#PostCompanyContactTypeAssociations) | **Post** /company/contactTypeAssociations | Post ContactTypeAssociation
[**PutCompanyContactTypeAssociationsById**](ContactTypeAssociationsAPI.md#PutCompanyContactTypeAssociationsById) | **Put** /company/contactTypeAssociations/{id} | Put ContactTypeAssociation



## DeleteCompanyContactTypeAssociationsById

> DeleteCompanyContactTypeAssociationsById(ctx, id).ClientId(clientId).Execute()

Delete ContactTypeAssociation

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
	id := int32(56) // int32 | contactTypeAssociationId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContactTypeAssociationsAPI.DeleteCompanyContactTypeAssociationsById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactTypeAssociationsAPI.DeleteCompanyContactTypeAssociationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | contactTypeAssociationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyContactTypeAssociationsByIdRequest struct via the builder pattern


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


## GetCompanyContactTypeAssociations

> []CompanyContactTypeAssociation GetCompanyContactTypeAssociations(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ContactTypeAssociation

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
	resp, r, err := apiClient.ContactTypeAssociationsAPI.GetCompanyContactTypeAssociations(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactTypeAssociationsAPI.GetCompanyContactTypeAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyContactTypeAssociations`: []CompanyContactTypeAssociation
	fmt.Fprintf(os.Stdout, "Response from `ContactTypeAssociationsAPI.GetCompanyContactTypeAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyContactTypeAssociationsRequest struct via the builder pattern


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

[**[]CompanyContactTypeAssociation**](CompanyContactTypeAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyContactTypeAssociationsById

> CompanyContactTypeAssociation GetCompanyContactTypeAssociationsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ContactTypeAssociation

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
	id := int32(56) // int32 | contactTypeAssociationId
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
	resp, r, err := apiClient.ContactTypeAssociationsAPI.GetCompanyContactTypeAssociationsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactTypeAssociationsAPI.GetCompanyContactTypeAssociationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyContactTypeAssociationsById`: CompanyContactTypeAssociation
	fmt.Fprintf(os.Stdout, "Response from `ContactTypeAssociationsAPI.GetCompanyContactTypeAssociationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | contactTypeAssociationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyContactTypeAssociationsByIdRequest struct via the builder pattern


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

[**CompanyContactTypeAssociation**](CompanyContactTypeAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyContactTypeAssociationsCount

> Count GetCompanyContactTypeAssociationsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ContactTypeAssociation

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
	resp, r, err := apiClient.ContactTypeAssociationsAPI.GetCompanyContactTypeAssociationsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactTypeAssociationsAPI.GetCompanyContactTypeAssociationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyContactTypeAssociationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ContactTypeAssociationsAPI.GetCompanyContactTypeAssociationsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyContactTypeAssociationsCountRequest struct via the builder pattern


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


## PatchCompanyContactTypeAssociationsById

> CompanyContactTypeAssociation PatchCompanyContactTypeAssociationsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ContactTypeAssociation

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
	id := int32(56) // int32 | contactTypeAssociationId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactTypeAssociationsAPI.PatchCompanyContactTypeAssociationsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactTypeAssociationsAPI.PatchCompanyContactTypeAssociationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyContactTypeAssociationsById`: CompanyContactTypeAssociation
	fmt.Fprintf(os.Stdout, "Response from `ContactTypeAssociationsAPI.PatchCompanyContactTypeAssociationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | contactTypeAssociationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyContactTypeAssociationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**CompanyContactTypeAssociation**](CompanyContactTypeAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyContactTypeAssociations

> CompanyContactTypeAssociation PostCompanyContactTypeAssociations(ctx).ClientId(clientId).CompanyContactTypeAssociation(companyContactTypeAssociation).Execute()

Post ContactTypeAssociation

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
	companyContactTypeAssociation := *openapiclient.NewCompanyContactTypeAssociation(*openapiclient.NewContactTypeReference(), *openapiclient.NewContactReference()) // CompanyContactTypeAssociation | contactTypeAssociation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactTypeAssociationsAPI.PostCompanyContactTypeAssociations(context.Background()).ClientId(clientId).CompanyContactTypeAssociation(companyContactTypeAssociation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactTypeAssociationsAPI.PostCompanyContactTypeAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyContactTypeAssociations`: CompanyContactTypeAssociation
	fmt.Fprintf(os.Stdout, "Response from `ContactTypeAssociationsAPI.PostCompanyContactTypeAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyContactTypeAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **companyContactTypeAssociation** | [**CompanyContactTypeAssociation**](CompanyContactTypeAssociation.md) | contactTypeAssociation | 

### Return type

[**CompanyContactTypeAssociation**](CompanyContactTypeAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyContactTypeAssociationsById

> CompanyContactTypeAssociation PutCompanyContactTypeAssociationsById(ctx, id).ClientId(clientId).CompanyContactTypeAssociation(companyContactTypeAssociation).Execute()

Put ContactTypeAssociation

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
	id := int32(56) // int32 | contactTypeAssociationId
	clientId := "clientId_example" // string | 
	companyContactTypeAssociation := *openapiclient.NewCompanyContactTypeAssociation(*openapiclient.NewContactTypeReference(), *openapiclient.NewContactReference()) // CompanyContactTypeAssociation | contactTypeAssociation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactTypeAssociationsAPI.PutCompanyContactTypeAssociationsById(context.Background(), id).ClientId(clientId).CompanyContactTypeAssociation(companyContactTypeAssociation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactTypeAssociationsAPI.PutCompanyContactTypeAssociationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyContactTypeAssociationsById`: CompanyContactTypeAssociation
	fmt.Fprintf(os.Stdout, "Response from `ContactTypeAssociationsAPI.PutCompanyContactTypeAssociationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | contactTypeAssociationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyContactTypeAssociationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **companyContactTypeAssociation** | [**CompanyContactTypeAssociation**](CompanyContactTypeAssociation.md) | contactTypeAssociation | 

### Return type

[**CompanyContactTypeAssociation**](CompanyContactTypeAssociation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

