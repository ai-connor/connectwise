# \TaxableExpenseTypeLevelsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById**](TaxableExpenseTypeLevelsAPI.md#DeleteFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById) | **Delete** /finance/taxCodes/{grandparentId}/expenseTypeExemptions/{parentId}/taxableExpenseTypeLevels/{id} | Delete TaxableExpenseTypeLevel
[**GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevels**](TaxableExpenseTypeLevelsAPI.md#GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevels) | **Get** /finance/taxCodes/{grandparentId}/expenseTypeExemptions/{parentId}/taxableExpenseTypeLevels | Get List of TaxableExpenseTypeLevel
[**GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById**](TaxableExpenseTypeLevelsAPI.md#GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById) | **Get** /finance/taxCodes/{grandparentId}/expenseTypeExemptions/{parentId}/taxableExpenseTypeLevels/{id} | Get TaxableExpenseTypeLevel
[**GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsCount**](TaxableExpenseTypeLevelsAPI.md#GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsCount) | **Get** /finance/taxCodes/{grandparentId}/expenseTypeExemptions/{parentId}/taxableExpenseTypeLevels/count | Get Count of TaxableExpenseTypeLevel
[**PatchFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById**](TaxableExpenseTypeLevelsAPI.md#PatchFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById) | **Patch** /finance/taxCodes/{grandparentId}/expenseTypeExemptions/{parentId}/taxableExpenseTypeLevels/{id} | Patch TaxableExpenseTypeLevel
[**PostFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevels**](TaxableExpenseTypeLevelsAPI.md#PostFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevels) | **Post** /finance/taxCodes/{grandparentId}/expenseTypeExemptions/{parentId}/taxableExpenseTypeLevels | Post TaxableExpenseTypeLevel
[**PutFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById**](TaxableExpenseTypeLevelsAPI.md#PutFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById) | **Put** /finance/taxCodes/{grandparentId}/expenseTypeExemptions/{parentId}/taxableExpenseTypeLevels/{id} | Put TaxableExpenseTypeLevel



## DeleteFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById

> DeleteFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById(ctx, id, parentId, grandparentId).ClientId(clientId).Execute()

Delete TaxableExpenseTypeLevel

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
	id := int32(56) // int32 | taxableExpenseTypeLevelId
	parentId := int32(56) // int32 | expenseTypeExemptionId
	grandparentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaxableExpenseTypeLevelsAPI.DeleteFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableExpenseTypeLevelsAPI.DeleteFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxableExpenseTypeLevelId | 
**parentId** | **int32** | expenseTypeExemptionId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsByIdRequest struct via the builder pattern


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


## GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevels

> []TaxableExpenseTypeLevel GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevels(ctx, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of TaxableExpenseTypeLevel

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
	parentId := int32(56) // int32 | expenseTypeExemptionId
	grandparentId := int32(56) // int32 | taxCodeId
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
	resp, r, err := apiClient.TaxableExpenseTypeLevelsAPI.GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevels(context.Background(), parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableExpenseTypeLevelsAPI.GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevels`: []TaxableExpenseTypeLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxableExpenseTypeLevelsAPI.GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | expenseTypeExemptionId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsRequest struct via the builder pattern


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

[**[]TaxableExpenseTypeLevel**](TaxableExpenseTypeLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById

> TaxableExpenseTypeLevel GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById(ctx, id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get TaxableExpenseTypeLevel

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
	id := int32(56) // int32 | taxableExpenseTypeLevelId
	parentId := int32(56) // int32 | expenseTypeExemptionId
	grandparentId := int32(56) // int32 | taxCodeId
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
	resp, r, err := apiClient.TaxableExpenseTypeLevelsAPI.GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableExpenseTypeLevelsAPI.GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById`: TaxableExpenseTypeLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxableExpenseTypeLevelsAPI.GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxableExpenseTypeLevelId | 
**parentId** | **int32** | expenseTypeExemptionId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsByIdRequest struct via the builder pattern


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

[**TaxableExpenseTypeLevel**](TaxableExpenseTypeLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsCount

> Count GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsCount(ctx, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of TaxableExpenseTypeLevel

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
	parentId := int32(56) // int32 | expenseTypeExemptionId
	grandparentId := int32(56) // int32 | taxCodeId
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
	resp, r, err := apiClient.TaxableExpenseTypeLevelsAPI.GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsCount(context.Background(), parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableExpenseTypeLevelsAPI.GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `TaxableExpenseTypeLevelsAPI.GetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | expenseTypeExemptionId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsCountRequest struct via the builder pattern


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


## PatchFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById

> TaxableExpenseTypeLevel PatchFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById(ctx, id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch TaxableExpenseTypeLevel

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
	id := int32(56) // int32 | taxableExpenseTypeLevelId
	parentId := int32(56) // int32 | expenseTypeExemptionId
	grandparentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxableExpenseTypeLevelsAPI.PatchFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableExpenseTypeLevelsAPI.PatchFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById`: TaxableExpenseTypeLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxableExpenseTypeLevelsAPI.PatchFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxableExpenseTypeLevelId | 
**parentId** | **int32** | expenseTypeExemptionId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**TaxableExpenseTypeLevel**](TaxableExpenseTypeLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevels

> TaxableExpenseTypeLevel PostFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevels(ctx, parentId, grandparentId).ClientId(clientId).TaxableExpenseTypeLevel(taxableExpenseTypeLevel).Execute()

Post TaxableExpenseTypeLevel

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
	parentId := int32(56) // int32 | expenseTypeExemptionId
	grandparentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 
	taxableExpenseTypeLevel := *openapiclient.NewTaxableExpenseTypeLevel(*openapiclient.NewTaxCodeLevelReference()) // TaxableExpenseTypeLevel | taxableExpenseTypeLevel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxableExpenseTypeLevelsAPI.PostFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevels(context.Background(), parentId, grandparentId).ClientId(clientId).TaxableExpenseTypeLevel(taxableExpenseTypeLevel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableExpenseTypeLevelsAPI.PostFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevels`: TaxableExpenseTypeLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxableExpenseTypeLevelsAPI.PostFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | expenseTypeExemptionId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **taxableExpenseTypeLevel** | [**TaxableExpenseTypeLevel**](TaxableExpenseTypeLevel.md) | taxableExpenseTypeLevel | 

### Return type

[**TaxableExpenseTypeLevel**](TaxableExpenseTypeLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById

> TaxableExpenseTypeLevel PutFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById(ctx, id, parentId, grandparentId).ClientId(clientId).TaxableExpenseTypeLevel(taxableExpenseTypeLevel).Execute()

Put TaxableExpenseTypeLevel

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
	id := int32(56) // int32 | taxableExpenseTypeLevelId
	parentId := int32(56) // int32 | expenseTypeExemptionId
	grandparentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 
	taxableExpenseTypeLevel := *openapiclient.NewTaxableExpenseTypeLevel(*openapiclient.NewTaxCodeLevelReference()) // TaxableExpenseTypeLevel | taxableExpenseTypeLevel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxableExpenseTypeLevelsAPI.PutFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).TaxableExpenseTypeLevel(taxableExpenseTypeLevel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableExpenseTypeLevelsAPI.PutFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById`: TaxableExpenseTypeLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxableExpenseTypeLevelsAPI.PutFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxableExpenseTypeLevelId | 
**parentId** | **int32** | expenseTypeExemptionId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceTaxCodesByGrandparentIdExpenseTypeExemptionsByParentIdTaxableExpenseTypeLevelsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **taxableExpenseTypeLevel** | [**TaxableExpenseTypeLevel**](TaxableExpenseTypeLevel.md) | taxableExpenseTypeLevel | 

### Return type

[**TaxableExpenseTypeLevel**](TaxableExpenseTypeLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

