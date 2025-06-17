# \TaxableProductTypeLevelsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById**](TaxableProductTypeLevelsAPI.md#DeleteFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById) | **Delete** /finance/taxCodes/{grandparentId}/productTypeExemptions/{parentId}/taxableProductTypeLevels/{id} | Delete TaxableProductTypeLevel
[**GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevels**](TaxableProductTypeLevelsAPI.md#GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevels) | **Get** /finance/taxCodes/{grandparentId}/productTypeExemptions/{parentId}/taxableProductTypeLevels | Get List of TaxableProductTypeLevel
[**GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById**](TaxableProductTypeLevelsAPI.md#GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById) | **Get** /finance/taxCodes/{grandparentId}/productTypeExemptions/{parentId}/taxableProductTypeLevels/{id} | Get TaxableProductTypeLevel
[**GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsCount**](TaxableProductTypeLevelsAPI.md#GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsCount) | **Get** /finance/taxCodes/{grandparentId}/productTypeExemptions/{parentId}/taxableProductTypeLevels/count | Get Count of TaxableProductTypeLevel
[**PatchFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById**](TaxableProductTypeLevelsAPI.md#PatchFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById) | **Patch** /finance/taxCodes/{grandparentId}/productTypeExemptions/{parentId}/taxableProductTypeLevels/{id} | Patch TaxableProductTypeLevel
[**PostFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevels**](TaxableProductTypeLevelsAPI.md#PostFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevels) | **Post** /finance/taxCodes/{grandparentId}/productTypeExemptions/{parentId}/taxableProductTypeLevels | Post TaxableProductTypeLevel
[**PutFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById**](TaxableProductTypeLevelsAPI.md#PutFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById) | **Put** /finance/taxCodes/{grandparentId}/productTypeExemptions/{parentId}/taxableProductTypeLevels/{id} | Put TaxableProductTypeLevel



## DeleteFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById

> DeleteFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById(ctx, id, parentId, grandparentId).ClientId(clientId).Execute()

Delete TaxableProductTypeLevel

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
	id := int32(56) // int32 | taxableProductTypeLevelId
	parentId := int32(56) // int32 | productTypeExemptionId
	grandparentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaxableProductTypeLevelsAPI.DeleteFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableProductTypeLevelsAPI.DeleteFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxableProductTypeLevelId | 
**parentId** | **int32** | productTypeExemptionId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsByIdRequest struct via the builder pattern


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


## GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevels

> []TaxableProductTypeLevel GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevels(ctx, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of TaxableProductTypeLevel

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
	parentId := int32(56) // int32 | productTypeExemptionId
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
	resp, r, err := apiClient.TaxableProductTypeLevelsAPI.GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevels(context.Background(), parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableProductTypeLevelsAPI.GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevels`: []TaxableProductTypeLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxableProductTypeLevelsAPI.GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | productTypeExemptionId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsRequest struct via the builder pattern


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

[**[]TaxableProductTypeLevel**](TaxableProductTypeLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById

> TaxableProductTypeLevel GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById(ctx, id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get TaxableProductTypeLevel

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
	id := int32(56) // int32 | taxableProductTypeLevelId
	parentId := int32(56) // int32 | productTypeExemptionId
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
	resp, r, err := apiClient.TaxableProductTypeLevelsAPI.GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableProductTypeLevelsAPI.GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById`: TaxableProductTypeLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxableProductTypeLevelsAPI.GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxableProductTypeLevelId | 
**parentId** | **int32** | productTypeExemptionId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsByIdRequest struct via the builder pattern


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

[**TaxableProductTypeLevel**](TaxableProductTypeLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsCount

> Count GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsCount(ctx, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of TaxableProductTypeLevel

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
	parentId := int32(56) // int32 | productTypeExemptionId
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
	resp, r, err := apiClient.TaxableProductTypeLevelsAPI.GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsCount(context.Background(), parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableProductTypeLevelsAPI.GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `TaxableProductTypeLevelsAPI.GetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | productTypeExemptionId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsCountRequest struct via the builder pattern


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


## PatchFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById

> TaxableProductTypeLevel PatchFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById(ctx, id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch TaxableProductTypeLevel

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
	id := int32(56) // int32 | taxableProductTypeLevelId
	parentId := int32(56) // int32 | productTypeExemptionId
	grandparentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxableProductTypeLevelsAPI.PatchFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableProductTypeLevelsAPI.PatchFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById`: TaxableProductTypeLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxableProductTypeLevelsAPI.PatchFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxableProductTypeLevelId | 
**parentId** | **int32** | productTypeExemptionId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**TaxableProductTypeLevel**](TaxableProductTypeLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevels

> TaxableProductTypeLevel PostFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevels(ctx, parentId, grandparentId).ClientId(clientId).TaxableProductTypeLevel(taxableProductTypeLevel).Execute()

Post TaxableProductTypeLevel

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
	parentId := int32(56) // int32 | productTypeExemptionId
	grandparentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 
	taxableProductTypeLevel := *openapiclient.NewTaxableProductTypeLevel(*openapiclient.NewTaxCodeLevelReference()) // TaxableProductTypeLevel | taxableProductTypeLevel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxableProductTypeLevelsAPI.PostFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevels(context.Background(), parentId, grandparentId).ClientId(clientId).TaxableProductTypeLevel(taxableProductTypeLevel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableProductTypeLevelsAPI.PostFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevels`: TaxableProductTypeLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxableProductTypeLevelsAPI.PostFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | productTypeExemptionId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **taxableProductTypeLevel** | [**TaxableProductTypeLevel**](TaxableProductTypeLevel.md) | taxableProductTypeLevel | 

### Return type

[**TaxableProductTypeLevel**](TaxableProductTypeLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById

> TaxableProductTypeLevel PutFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById(ctx, id, parentId, grandparentId).ClientId(clientId).TaxableProductTypeLevel(taxableProductTypeLevel).Execute()

Put TaxableProductTypeLevel

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
	id := int32(56) // int32 | taxableProductTypeLevelId
	parentId := int32(56) // int32 | productTypeExemptionId
	grandparentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 
	taxableProductTypeLevel := *openapiclient.NewTaxableProductTypeLevel(*openapiclient.NewTaxCodeLevelReference()) // TaxableProductTypeLevel | taxableProductTypeLevel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxableProductTypeLevelsAPI.PutFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).TaxableProductTypeLevel(taxableProductTypeLevel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableProductTypeLevelsAPI.PutFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById`: TaxableProductTypeLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxableProductTypeLevelsAPI.PutFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxableProductTypeLevelId | 
**parentId** | **int32** | productTypeExemptionId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceTaxCodesByGrandparentIdProductTypeExemptionsByParentIdTaxableProductTypeLevelsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **taxableProductTypeLevel** | [**TaxableProductTypeLevel**](TaxableProductTypeLevel.md) | taxableProductTypeLevel | 

### Return type

[**TaxableProductTypeLevel**](TaxableProductTypeLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

