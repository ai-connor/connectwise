# \TaxCodeWorkRoleExemptionsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceTaxCodesByParentIdWorkRoleExemptionsById**](TaxCodeWorkRoleExemptionsAPI.md#DeleteFinanceTaxCodesByParentIdWorkRoleExemptionsById) | **Delete** /finance/taxCodes/{parentId}/workRoleExemptions/{id} | Delete WorkRoleExemption
[**GetFinanceTaxCodesByParentIdWorkRoleExemptions**](TaxCodeWorkRoleExemptionsAPI.md#GetFinanceTaxCodesByParentIdWorkRoleExemptions) | **Get** /finance/taxCodes/{parentId}/workRoleExemptions | Get List of WorkRoleExemption
[**GetFinanceTaxCodesByParentIdWorkRoleExemptionsById**](TaxCodeWorkRoleExemptionsAPI.md#GetFinanceTaxCodesByParentIdWorkRoleExemptionsById) | **Get** /finance/taxCodes/{parentId}/workRoleExemptions/{id} | Get WorkRoleExemption
[**GetFinanceTaxCodesByParentIdWorkRoleExemptionsCount**](TaxCodeWorkRoleExemptionsAPI.md#GetFinanceTaxCodesByParentIdWorkRoleExemptionsCount) | **Get** /finance/taxCodes/{parentId}/workRoleExemptions/count | Get Count of WorkRoleExemption
[**PatchFinanceTaxCodesByParentIdWorkRoleExemptionsById**](TaxCodeWorkRoleExemptionsAPI.md#PatchFinanceTaxCodesByParentIdWorkRoleExemptionsById) | **Patch** /finance/taxCodes/{parentId}/workRoleExemptions/{id} | Patch WorkRoleExemption
[**PostFinanceTaxCodesByParentIdWorkRoleExemptions**](TaxCodeWorkRoleExemptionsAPI.md#PostFinanceTaxCodesByParentIdWorkRoleExemptions) | **Post** /finance/taxCodes/{parentId}/workRoleExemptions | Post WorkRoleExemption
[**PutFinanceTaxCodesByParentIdWorkRoleExemptionsById**](TaxCodeWorkRoleExemptionsAPI.md#PutFinanceTaxCodesByParentIdWorkRoleExemptionsById) | **Put** /finance/taxCodes/{parentId}/workRoleExemptions/{id} | Put WorkRoleExemption



## DeleteFinanceTaxCodesByParentIdWorkRoleExemptionsById

> DeleteFinanceTaxCodesByParentIdWorkRoleExemptionsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete WorkRoleExemption

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
	id := int32(56) // int32 | workRoleExemptionId
	parentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaxCodeWorkRoleExemptionsAPI.DeleteFinanceTaxCodesByParentIdWorkRoleExemptionsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeWorkRoleExemptionsAPI.DeleteFinanceTaxCodesByParentIdWorkRoleExemptionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | workRoleExemptionId | 
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceTaxCodesByParentIdWorkRoleExemptionsByIdRequest struct via the builder pattern


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


## GetFinanceTaxCodesByParentIdWorkRoleExemptions

> []WorkRoleExemption GetFinanceTaxCodesByParentIdWorkRoleExemptions(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of WorkRoleExemption

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
	parentId := int32(56) // int32 | taxCodeId
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
	resp, r, err := apiClient.TaxCodeWorkRoleExemptionsAPI.GetFinanceTaxCodesByParentIdWorkRoleExemptions(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeWorkRoleExemptionsAPI.GetFinanceTaxCodesByParentIdWorkRoleExemptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByParentIdWorkRoleExemptions`: []WorkRoleExemption
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeWorkRoleExemptionsAPI.GetFinanceTaxCodesByParentIdWorkRoleExemptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByParentIdWorkRoleExemptionsRequest struct via the builder pattern


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

[**[]WorkRoleExemption**](WorkRoleExemption.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceTaxCodesByParentIdWorkRoleExemptionsById

> WorkRoleExemption GetFinanceTaxCodesByParentIdWorkRoleExemptionsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get WorkRoleExemption

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
	id := int32(56) // int32 | workRoleExemptionId
	parentId := int32(56) // int32 | taxCodeId
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
	resp, r, err := apiClient.TaxCodeWorkRoleExemptionsAPI.GetFinanceTaxCodesByParentIdWorkRoleExemptionsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeWorkRoleExemptionsAPI.GetFinanceTaxCodesByParentIdWorkRoleExemptionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByParentIdWorkRoleExemptionsById`: WorkRoleExemption
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeWorkRoleExemptionsAPI.GetFinanceTaxCodesByParentIdWorkRoleExemptionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | workRoleExemptionId | 
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByParentIdWorkRoleExemptionsByIdRequest struct via the builder pattern


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

[**WorkRoleExemption**](WorkRoleExemption.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceTaxCodesByParentIdWorkRoleExemptionsCount

> Count GetFinanceTaxCodesByParentIdWorkRoleExemptionsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of WorkRoleExemption

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
	parentId := int32(56) // int32 | taxCodeId
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
	resp, r, err := apiClient.TaxCodeWorkRoleExemptionsAPI.GetFinanceTaxCodesByParentIdWorkRoleExemptionsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeWorkRoleExemptionsAPI.GetFinanceTaxCodesByParentIdWorkRoleExemptionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByParentIdWorkRoleExemptionsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeWorkRoleExemptionsAPI.GetFinanceTaxCodesByParentIdWorkRoleExemptionsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByParentIdWorkRoleExemptionsCountRequest struct via the builder pattern


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


## PatchFinanceTaxCodesByParentIdWorkRoleExemptionsById

> WorkRoleExemption PatchFinanceTaxCodesByParentIdWorkRoleExemptionsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch WorkRoleExemption

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
	id := int32(56) // int32 | workRoleExemptionId
	parentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxCodeWorkRoleExemptionsAPI.PatchFinanceTaxCodesByParentIdWorkRoleExemptionsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeWorkRoleExemptionsAPI.PatchFinanceTaxCodesByParentIdWorkRoleExemptionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceTaxCodesByParentIdWorkRoleExemptionsById`: WorkRoleExemption
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeWorkRoleExemptionsAPI.PatchFinanceTaxCodesByParentIdWorkRoleExemptionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | workRoleExemptionId | 
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceTaxCodesByParentIdWorkRoleExemptionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**WorkRoleExemption**](WorkRoleExemption.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceTaxCodesByParentIdWorkRoleExemptions

> WorkRoleExemption PostFinanceTaxCodesByParentIdWorkRoleExemptions(ctx, parentId).ClientId(clientId).WorkRoleExemption(workRoleExemption).Execute()

Post WorkRoleExemption

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
	parentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 
	workRoleExemption := *openapiclient.NewWorkRoleExemption(*openapiclient.NewWorkRoleReference()) // WorkRoleExemption | workRoleExemption

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxCodeWorkRoleExemptionsAPI.PostFinanceTaxCodesByParentIdWorkRoleExemptions(context.Background(), parentId).ClientId(clientId).WorkRoleExemption(workRoleExemption).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeWorkRoleExemptionsAPI.PostFinanceTaxCodesByParentIdWorkRoleExemptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceTaxCodesByParentIdWorkRoleExemptions`: WorkRoleExemption
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeWorkRoleExemptionsAPI.PostFinanceTaxCodesByParentIdWorkRoleExemptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceTaxCodesByParentIdWorkRoleExemptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **workRoleExemption** | [**WorkRoleExemption**](WorkRoleExemption.md) | workRoleExemption | 

### Return type

[**WorkRoleExemption**](WorkRoleExemption.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceTaxCodesByParentIdWorkRoleExemptionsById

> WorkRoleExemption PutFinanceTaxCodesByParentIdWorkRoleExemptionsById(ctx, id, parentId).ClientId(clientId).WorkRoleExemption(workRoleExemption).Execute()

Put WorkRoleExemption

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
	id := int32(56) // int32 | workRoleExemptionId
	parentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 
	workRoleExemption := *openapiclient.NewWorkRoleExemption(*openapiclient.NewWorkRoleReference()) // WorkRoleExemption | workRoleExemption

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxCodeWorkRoleExemptionsAPI.PutFinanceTaxCodesByParentIdWorkRoleExemptionsById(context.Background(), id, parentId).ClientId(clientId).WorkRoleExemption(workRoleExemption).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeWorkRoleExemptionsAPI.PutFinanceTaxCodesByParentIdWorkRoleExemptionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceTaxCodesByParentIdWorkRoleExemptionsById`: WorkRoleExemption
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeWorkRoleExemptionsAPI.PutFinanceTaxCodesByParentIdWorkRoleExemptionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | workRoleExemptionId | 
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceTaxCodesByParentIdWorkRoleExemptionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **workRoleExemption** | [**WorkRoleExemption**](WorkRoleExemption.md) | workRoleExemption | 

### Return type

[**WorkRoleExemption**](WorkRoleExemption.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

