# \ManagementItSolutionAgreementInterfaceParametersAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyManagementItSolutionsByParentIdManagementProductsById**](ManagementItSolutionAgreementInterfaceParametersAPI.md#DeleteCompanyManagementItSolutionsByParentIdManagementProductsById) | **Delete** /company/managementItSolutions/{parentId}/managementProducts/{id} | Delete ManagementItSolutionAgreementInterfaceParameter
[**GetCompanyManagementItSolutionsByParentIdManagementProducts**](ManagementItSolutionAgreementInterfaceParametersAPI.md#GetCompanyManagementItSolutionsByParentIdManagementProducts) | **Get** /company/managementItSolutions/{parentId}/managementProducts | Get List of ManagementItSolutionAgreementInterfaceParameter
[**GetCompanyManagementItSolutionsByParentIdManagementProductsById**](ManagementItSolutionAgreementInterfaceParametersAPI.md#GetCompanyManagementItSolutionsByParentIdManagementProductsById) | **Get** /company/managementItSolutions/{parentId}/managementProducts/{id} | Get ManagementItSolutionAgreementInterfaceParameter
[**GetCompanyManagementItSolutionsByParentIdManagementProductsCount**](ManagementItSolutionAgreementInterfaceParametersAPI.md#GetCompanyManagementItSolutionsByParentIdManagementProductsCount) | **Get** /company/managementItSolutions/{parentId}/managementProducts/count | Get Count of ManagementItSolutionAgreementInterfaceParameter
[**PatchCompanyManagementItSolutionsByParentIdManagementProductsById**](ManagementItSolutionAgreementInterfaceParametersAPI.md#PatchCompanyManagementItSolutionsByParentIdManagementProductsById) | **Patch** /company/managementItSolutions/{parentId}/managementProducts/{id} | Patch ManagementItSolutionAgreementInterfaceParameter
[**PostCompanyManagementItSolutionsByParentIdManagementProducts**](ManagementItSolutionAgreementInterfaceParametersAPI.md#PostCompanyManagementItSolutionsByParentIdManagementProducts) | **Post** /company/managementItSolutions/{parentId}/managementProducts | Post ManagementItSolutionAgreementInterfaceParameter
[**PutCompanyManagementItSolutionsByParentIdManagementProductsById**](ManagementItSolutionAgreementInterfaceParametersAPI.md#PutCompanyManagementItSolutionsByParentIdManagementProductsById) | **Put** /company/managementItSolutions/{parentId}/managementProducts/{id} | Put ManagementItSolutionAgreementInterfaceParameter



## DeleteCompanyManagementItSolutionsByParentIdManagementProductsById

> ManagementItSolutionAgreementInterfaceParameter DeleteCompanyManagementItSolutionsByParentIdManagementProductsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ManagementItSolutionAgreementInterfaceParameter

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
	id := int32(56) // int32 | managementProductId
	parentId := int32(56) // int32 | managementItSolutionId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementItSolutionAgreementInterfaceParametersAPI.DeleteCompanyManagementItSolutionsByParentIdManagementProductsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementItSolutionAgreementInterfaceParametersAPI.DeleteCompanyManagementItSolutionsByParentIdManagementProductsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCompanyManagementItSolutionsByParentIdManagementProductsById`: ManagementItSolutionAgreementInterfaceParameter
	fmt.Fprintf(os.Stdout, "Response from `ManagementItSolutionAgreementInterfaceParametersAPI.DeleteCompanyManagementItSolutionsByParentIdManagementProductsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementProductId | 
**parentId** | **int32** | managementItSolutionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyManagementItSolutionsByParentIdManagementProductsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 

### Return type

[**ManagementItSolutionAgreementInterfaceParameter**](ManagementItSolutionAgreementInterfaceParameter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyManagementItSolutionsByParentIdManagementProducts

> []ManagementItSolutionAgreementInterfaceParameter GetCompanyManagementItSolutionsByParentIdManagementProducts(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ManagementItSolutionAgreementInterfaceParameter

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
	parentId := int32(56) // int32 | managementItSolutionId
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
	resp, r, err := apiClient.ManagementItSolutionAgreementInterfaceParametersAPI.GetCompanyManagementItSolutionsByParentIdManagementProducts(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementItSolutionAgreementInterfaceParametersAPI.GetCompanyManagementItSolutionsByParentIdManagementProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyManagementItSolutionsByParentIdManagementProducts`: []ManagementItSolutionAgreementInterfaceParameter
	fmt.Fprintf(os.Stdout, "Response from `ManagementItSolutionAgreementInterfaceParametersAPI.GetCompanyManagementItSolutionsByParentIdManagementProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | managementItSolutionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyManagementItSolutionsByParentIdManagementProductsRequest struct via the builder pattern


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

[**[]ManagementItSolutionAgreementInterfaceParameter**](ManagementItSolutionAgreementInterfaceParameter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyManagementItSolutionsByParentIdManagementProductsById

> ManagementItSolutionAgreementInterfaceParameter GetCompanyManagementItSolutionsByParentIdManagementProductsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ManagementItSolutionAgreementInterfaceParameter

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
	id := int32(56) // int32 | managementProductId
	parentId := int32(56) // int32 | managementItSolutionId
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
	resp, r, err := apiClient.ManagementItSolutionAgreementInterfaceParametersAPI.GetCompanyManagementItSolutionsByParentIdManagementProductsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementItSolutionAgreementInterfaceParametersAPI.GetCompanyManagementItSolutionsByParentIdManagementProductsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyManagementItSolutionsByParentIdManagementProductsById`: ManagementItSolutionAgreementInterfaceParameter
	fmt.Fprintf(os.Stdout, "Response from `ManagementItSolutionAgreementInterfaceParametersAPI.GetCompanyManagementItSolutionsByParentIdManagementProductsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementProductId | 
**parentId** | **int32** | managementItSolutionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyManagementItSolutionsByParentIdManagementProductsByIdRequest struct via the builder pattern


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

[**ManagementItSolutionAgreementInterfaceParameter**](ManagementItSolutionAgreementInterfaceParameter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyManagementItSolutionsByParentIdManagementProductsCount

> Count GetCompanyManagementItSolutionsByParentIdManagementProductsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ManagementItSolutionAgreementInterfaceParameter

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
	parentId := int32(56) // int32 | managementItSolutionId
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
	resp, r, err := apiClient.ManagementItSolutionAgreementInterfaceParametersAPI.GetCompanyManagementItSolutionsByParentIdManagementProductsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementItSolutionAgreementInterfaceParametersAPI.GetCompanyManagementItSolutionsByParentIdManagementProductsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyManagementItSolutionsByParentIdManagementProductsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ManagementItSolutionAgreementInterfaceParametersAPI.GetCompanyManagementItSolutionsByParentIdManagementProductsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | managementItSolutionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyManagementItSolutionsByParentIdManagementProductsCountRequest struct via the builder pattern


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


## PatchCompanyManagementItSolutionsByParentIdManagementProductsById

> ManagementItSolutionAgreementInterfaceParameter PatchCompanyManagementItSolutionsByParentIdManagementProductsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ManagementItSolutionAgreementInterfaceParameter

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
	id := int32(56) // int32 | managementProductId
	parentId := int32(56) // int32 | managementItSolutionId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementItSolutionAgreementInterfaceParametersAPI.PatchCompanyManagementItSolutionsByParentIdManagementProductsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementItSolutionAgreementInterfaceParametersAPI.PatchCompanyManagementItSolutionsByParentIdManagementProductsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyManagementItSolutionsByParentIdManagementProductsById`: ManagementItSolutionAgreementInterfaceParameter
	fmt.Fprintf(os.Stdout, "Response from `ManagementItSolutionAgreementInterfaceParametersAPI.PatchCompanyManagementItSolutionsByParentIdManagementProductsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementProductId | 
**parentId** | **int32** | managementItSolutionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyManagementItSolutionsByParentIdManagementProductsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ManagementItSolutionAgreementInterfaceParameter**](ManagementItSolutionAgreementInterfaceParameter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyManagementItSolutionsByParentIdManagementProducts

> ManagementItSolutionAgreementInterfaceParameter PostCompanyManagementItSolutionsByParentIdManagementProducts(ctx, parentId).ClientId(clientId).ManagementItSolutionAgreementInterfaceParameter(managementItSolutionAgreementInterfaceParameter).Execute()

Post ManagementItSolutionAgreementInterfaceParameter

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
	parentId := int32(56) // int32 | managementItSolutionId
	clientId := "clientId_example" // string | 
	managementItSolutionAgreementInterfaceParameter := *openapiclient.NewManagementItSolutionAgreementInterfaceParameter(*openapiclient.NewAgreementTypeReference()) // ManagementItSolutionAgreementInterfaceParameter | managementProduct

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementItSolutionAgreementInterfaceParametersAPI.PostCompanyManagementItSolutionsByParentIdManagementProducts(context.Background(), parentId).ClientId(clientId).ManagementItSolutionAgreementInterfaceParameter(managementItSolutionAgreementInterfaceParameter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementItSolutionAgreementInterfaceParametersAPI.PostCompanyManagementItSolutionsByParentIdManagementProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyManagementItSolutionsByParentIdManagementProducts`: ManagementItSolutionAgreementInterfaceParameter
	fmt.Fprintf(os.Stdout, "Response from `ManagementItSolutionAgreementInterfaceParametersAPI.PostCompanyManagementItSolutionsByParentIdManagementProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | managementItSolutionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyManagementItSolutionsByParentIdManagementProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **managementItSolutionAgreementInterfaceParameter** | [**ManagementItSolutionAgreementInterfaceParameter**](ManagementItSolutionAgreementInterfaceParameter.md) | managementProduct | 

### Return type

[**ManagementItSolutionAgreementInterfaceParameter**](ManagementItSolutionAgreementInterfaceParameter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyManagementItSolutionsByParentIdManagementProductsById

> ManagementItSolutionAgreementInterfaceParameter PutCompanyManagementItSolutionsByParentIdManagementProductsById(ctx, id, parentId).ClientId(clientId).ManagementItSolutionAgreementInterfaceParameter(managementItSolutionAgreementInterfaceParameter).Execute()

Put ManagementItSolutionAgreementInterfaceParameter

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
	id := int32(56) // int32 | managementProductId
	parentId := int32(56) // int32 | managementItSolutionId
	clientId := "clientId_example" // string | 
	managementItSolutionAgreementInterfaceParameter := *openapiclient.NewManagementItSolutionAgreementInterfaceParameter(*openapiclient.NewAgreementTypeReference()) // ManagementItSolutionAgreementInterfaceParameter | managementProduct

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementItSolutionAgreementInterfaceParametersAPI.PutCompanyManagementItSolutionsByParentIdManagementProductsById(context.Background(), id, parentId).ClientId(clientId).ManagementItSolutionAgreementInterfaceParameter(managementItSolutionAgreementInterfaceParameter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementItSolutionAgreementInterfaceParametersAPI.PutCompanyManagementItSolutionsByParentIdManagementProductsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyManagementItSolutionsByParentIdManagementProductsById`: ManagementItSolutionAgreementInterfaceParameter
	fmt.Fprintf(os.Stdout, "Response from `ManagementItSolutionAgreementInterfaceParametersAPI.PutCompanyManagementItSolutionsByParentIdManagementProductsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementProductId | 
**parentId** | **int32** | managementItSolutionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyManagementItSolutionsByParentIdManagementProductsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **managementItSolutionAgreementInterfaceParameter** | [**ManagementItSolutionAgreementInterfaceParameter**](ManagementItSolutionAgreementInterfaceParameter.md) | managementProduct | 

### Return type

[**ManagementItSolutionAgreementInterfaceParameter**](ManagementItSolutionAgreementInterfaceParameter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

