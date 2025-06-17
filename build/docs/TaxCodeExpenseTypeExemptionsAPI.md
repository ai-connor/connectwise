# \TaxCodeExpenseTypeExemptionsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceTaxCodesByParentIdExpenseTypeExemptionsById**](TaxCodeExpenseTypeExemptionsAPI.md#DeleteFinanceTaxCodesByParentIdExpenseTypeExemptionsById) | **Delete** /finance/taxCodes/{parentId}/expenseTypeExemptions/{id} | Delete ExpenseTypeExemption
[**GetFinanceTaxCodesByParentIdExpenseTypeExemptions**](TaxCodeExpenseTypeExemptionsAPI.md#GetFinanceTaxCodesByParentIdExpenseTypeExemptions) | **Get** /finance/taxCodes/{parentId}/expenseTypeExemptions | Get List of ExpenseTypeExemption
[**GetFinanceTaxCodesByParentIdExpenseTypeExemptionsById**](TaxCodeExpenseTypeExemptionsAPI.md#GetFinanceTaxCodesByParentIdExpenseTypeExemptionsById) | **Get** /finance/taxCodes/{parentId}/expenseTypeExemptions/{id} | Get ExpenseTypeExemption
[**GetFinanceTaxCodesByParentIdExpenseTypeExemptionsCount**](TaxCodeExpenseTypeExemptionsAPI.md#GetFinanceTaxCodesByParentIdExpenseTypeExemptionsCount) | **Get** /finance/taxCodes/{parentId}/expenseTypeExemptions/count | Get Count of ExpenseTypeExemption
[**PatchFinanceTaxCodesByParentIdExpenseTypeExemptionsById**](TaxCodeExpenseTypeExemptionsAPI.md#PatchFinanceTaxCodesByParentIdExpenseTypeExemptionsById) | **Patch** /finance/taxCodes/{parentId}/expenseTypeExemptions/{id} | Patch ExpenseTypeExemption
[**PostFinanceTaxCodesByParentIdExpenseTypeExemptions**](TaxCodeExpenseTypeExemptionsAPI.md#PostFinanceTaxCodesByParentIdExpenseTypeExemptions) | **Post** /finance/taxCodes/{parentId}/expenseTypeExemptions | Post ExpenseTypeExemption
[**PutFinanceTaxCodesByParentIdExpenseTypeExemptionsById**](TaxCodeExpenseTypeExemptionsAPI.md#PutFinanceTaxCodesByParentIdExpenseTypeExemptionsById) | **Put** /finance/taxCodes/{parentId}/expenseTypeExemptions/{id} | Put ExpenseTypeExemption



## DeleteFinanceTaxCodesByParentIdExpenseTypeExemptionsById

> DeleteFinanceTaxCodesByParentIdExpenseTypeExemptionsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ExpenseTypeExemption

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
	id := int32(56) // int32 | expenseTypeExemptionId
	parentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaxCodeExpenseTypeExemptionsAPI.DeleteFinanceTaxCodesByParentIdExpenseTypeExemptionsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeExpenseTypeExemptionsAPI.DeleteFinanceTaxCodesByParentIdExpenseTypeExemptionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | expenseTypeExemptionId | 
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceTaxCodesByParentIdExpenseTypeExemptionsByIdRequest struct via the builder pattern


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


## GetFinanceTaxCodesByParentIdExpenseTypeExemptions

> []ExpenseTypeExemption GetFinanceTaxCodesByParentIdExpenseTypeExemptions(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ExpenseTypeExemption

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
	resp, r, err := apiClient.TaxCodeExpenseTypeExemptionsAPI.GetFinanceTaxCodesByParentIdExpenseTypeExemptions(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeExpenseTypeExemptionsAPI.GetFinanceTaxCodesByParentIdExpenseTypeExemptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByParentIdExpenseTypeExemptions`: []ExpenseTypeExemption
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeExpenseTypeExemptionsAPI.GetFinanceTaxCodesByParentIdExpenseTypeExemptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByParentIdExpenseTypeExemptionsRequest struct via the builder pattern


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

[**[]ExpenseTypeExemption**](ExpenseTypeExemption.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceTaxCodesByParentIdExpenseTypeExemptionsById

> ExpenseTypeExemption GetFinanceTaxCodesByParentIdExpenseTypeExemptionsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ExpenseTypeExemption

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
	id := int32(56) // int32 | expenseTypeExemptionId
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
	resp, r, err := apiClient.TaxCodeExpenseTypeExemptionsAPI.GetFinanceTaxCodesByParentIdExpenseTypeExemptionsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeExpenseTypeExemptionsAPI.GetFinanceTaxCodesByParentIdExpenseTypeExemptionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByParentIdExpenseTypeExemptionsById`: ExpenseTypeExemption
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeExpenseTypeExemptionsAPI.GetFinanceTaxCodesByParentIdExpenseTypeExemptionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | expenseTypeExemptionId | 
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByParentIdExpenseTypeExemptionsByIdRequest struct via the builder pattern


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

[**ExpenseTypeExemption**](ExpenseTypeExemption.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceTaxCodesByParentIdExpenseTypeExemptionsCount

> Count GetFinanceTaxCodesByParentIdExpenseTypeExemptionsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ExpenseTypeExemption

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
	resp, r, err := apiClient.TaxCodeExpenseTypeExemptionsAPI.GetFinanceTaxCodesByParentIdExpenseTypeExemptionsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeExpenseTypeExemptionsAPI.GetFinanceTaxCodesByParentIdExpenseTypeExemptionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByParentIdExpenseTypeExemptionsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeExpenseTypeExemptionsAPI.GetFinanceTaxCodesByParentIdExpenseTypeExemptionsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByParentIdExpenseTypeExemptionsCountRequest struct via the builder pattern


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


## PatchFinanceTaxCodesByParentIdExpenseTypeExemptionsById

> ExpenseTypeExemption PatchFinanceTaxCodesByParentIdExpenseTypeExemptionsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ExpenseTypeExemption

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
	id := int32(56) // int32 | expenseTypeExemptionId
	parentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxCodeExpenseTypeExemptionsAPI.PatchFinanceTaxCodesByParentIdExpenseTypeExemptionsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeExpenseTypeExemptionsAPI.PatchFinanceTaxCodesByParentIdExpenseTypeExemptionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceTaxCodesByParentIdExpenseTypeExemptionsById`: ExpenseTypeExemption
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeExpenseTypeExemptionsAPI.PatchFinanceTaxCodesByParentIdExpenseTypeExemptionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | expenseTypeExemptionId | 
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceTaxCodesByParentIdExpenseTypeExemptionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ExpenseTypeExemption**](ExpenseTypeExemption.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceTaxCodesByParentIdExpenseTypeExemptions

> ExpenseTypeExemption PostFinanceTaxCodesByParentIdExpenseTypeExemptions(ctx, parentId).ClientId(clientId).ExpenseTypeExemption(expenseTypeExemption).Execute()

Post ExpenseTypeExemption

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
	expenseTypeExemption := *openapiclient.NewExpenseTypeExemption(*openapiclient.NewExpenseTypeReference()) // ExpenseTypeExemption | expenseTypeExemption

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxCodeExpenseTypeExemptionsAPI.PostFinanceTaxCodesByParentIdExpenseTypeExemptions(context.Background(), parentId).ClientId(clientId).ExpenseTypeExemption(expenseTypeExemption).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeExpenseTypeExemptionsAPI.PostFinanceTaxCodesByParentIdExpenseTypeExemptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceTaxCodesByParentIdExpenseTypeExemptions`: ExpenseTypeExemption
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeExpenseTypeExemptionsAPI.PostFinanceTaxCodesByParentIdExpenseTypeExemptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceTaxCodesByParentIdExpenseTypeExemptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **expenseTypeExemption** | [**ExpenseTypeExemption**](ExpenseTypeExemption.md) | expenseTypeExemption | 

### Return type

[**ExpenseTypeExemption**](ExpenseTypeExemption.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceTaxCodesByParentIdExpenseTypeExemptionsById

> ExpenseTypeExemption PutFinanceTaxCodesByParentIdExpenseTypeExemptionsById(ctx, id, parentId).ClientId(clientId).ExpenseTypeExemption(expenseTypeExemption).Execute()

Put ExpenseTypeExemption

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
	id := int32(56) // int32 | expenseTypeExemptionId
	parentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 
	expenseTypeExemption := *openapiclient.NewExpenseTypeExemption(*openapiclient.NewExpenseTypeReference()) // ExpenseTypeExemption | expenseTypeExemption

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxCodeExpenseTypeExemptionsAPI.PutFinanceTaxCodesByParentIdExpenseTypeExemptionsById(context.Background(), id, parentId).ClientId(clientId).ExpenseTypeExemption(expenseTypeExemption).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeExpenseTypeExemptionsAPI.PutFinanceTaxCodesByParentIdExpenseTypeExemptionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceTaxCodesByParentIdExpenseTypeExemptionsById`: ExpenseTypeExemption
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeExpenseTypeExemptionsAPI.PutFinanceTaxCodesByParentIdExpenseTypeExemptionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | expenseTypeExemptionId | 
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceTaxCodesByParentIdExpenseTypeExemptionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **expenseTypeExemption** | [**ExpenseTypeExemption**](ExpenseTypeExemption.md) | expenseTypeExemption | 

### Return type

[**ExpenseTypeExemption**](ExpenseTypeExemption.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

