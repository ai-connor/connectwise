# \TaxableWorkRoleLevelsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById**](TaxableWorkRoleLevelsAPI.md#DeleteFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById) | **Delete** /finance/taxCodes/{grandparentId}/workRoleExemptions/{parentId}/taxableWorkRoleLevels/{id} | Delete TaxableWorkRoleLevel
[**GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevels**](TaxableWorkRoleLevelsAPI.md#GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevels) | **Get** /finance/taxCodes/{grandparentId}/workRoleExemptions/{parentId}/taxableWorkRoleLevels | Get List of TaxableWorkRoleLevel
[**GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById**](TaxableWorkRoleLevelsAPI.md#GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById) | **Get** /finance/taxCodes/{grandparentId}/workRoleExemptions/{parentId}/taxableWorkRoleLevels/{id} | Get TaxableWorkRoleLevel
[**GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsCount**](TaxableWorkRoleLevelsAPI.md#GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsCount) | **Get** /finance/taxCodes/{grandparentId}/workRoleExemptions/{parentId}/taxableWorkRoleLevels/count | Get Count of TaxableWorkRoleLevel
[**PatchFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById**](TaxableWorkRoleLevelsAPI.md#PatchFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById) | **Patch** /finance/taxCodes/{grandparentId}/workRoleExemptions/{parentId}/taxableWorkRoleLevels/{id} | Patch TaxableWorkRoleLevel
[**PostFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevels**](TaxableWorkRoleLevelsAPI.md#PostFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevels) | **Post** /finance/taxCodes/{grandparentId}/workRoleExemptions/{parentId}/taxableWorkRoleLevels | Post TaxableWorkRoleLevel
[**PutFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById**](TaxableWorkRoleLevelsAPI.md#PutFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById) | **Put** /finance/taxCodes/{grandparentId}/workRoleExemptions/{parentId}/taxableWorkRoleLevels/{id} | Put TaxableWorkRoleLevel



## DeleteFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById

> DeleteFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById(ctx, id, parentId, grandparentId).ClientId(clientId).Execute()

Delete TaxableWorkRoleLevel

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
	id := int32(56) // int32 | taxableWorkRoleLevelId
	parentId := int32(56) // int32 | workRoleExemptionId
	grandparentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaxableWorkRoleLevelsAPI.DeleteFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableWorkRoleLevelsAPI.DeleteFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxableWorkRoleLevelId | 
**parentId** | **int32** | workRoleExemptionId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsByIdRequest struct via the builder pattern


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


## GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevels

> []TaxableWorkRoleLevel GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevels(ctx, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of TaxableWorkRoleLevel

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
	parentId := int32(56) // int32 | workRoleExemptionId
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
	resp, r, err := apiClient.TaxableWorkRoleLevelsAPI.GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevels(context.Background(), parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableWorkRoleLevelsAPI.GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevels`: []TaxableWorkRoleLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxableWorkRoleLevelsAPI.GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | workRoleExemptionId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsRequest struct via the builder pattern


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

[**[]TaxableWorkRoleLevel**](TaxableWorkRoleLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById

> TaxableWorkRoleLevel GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById(ctx, id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get TaxableWorkRoleLevel

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
	id := int32(56) // int32 | taxableWorkRoleLevelId
	parentId := int32(56) // int32 | workRoleExemptionId
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
	resp, r, err := apiClient.TaxableWorkRoleLevelsAPI.GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableWorkRoleLevelsAPI.GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById`: TaxableWorkRoleLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxableWorkRoleLevelsAPI.GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxableWorkRoleLevelId | 
**parentId** | **int32** | workRoleExemptionId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsByIdRequest struct via the builder pattern


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

[**TaxableWorkRoleLevel**](TaxableWorkRoleLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsCount

> Count GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsCount(ctx, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of TaxableWorkRoleLevel

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
	parentId := int32(56) // int32 | workRoleExemptionId
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
	resp, r, err := apiClient.TaxableWorkRoleLevelsAPI.GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsCount(context.Background(), parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableWorkRoleLevelsAPI.GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `TaxableWorkRoleLevelsAPI.GetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | workRoleExemptionId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsCountRequest struct via the builder pattern


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


## PatchFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById

> TaxableWorkRoleLevel PatchFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById(ctx, id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch TaxableWorkRoleLevel

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
	id := int32(56) // int32 | taxableWorkRoleLevelId
	parentId := int32(56) // int32 | workRoleExemptionId
	grandparentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxableWorkRoleLevelsAPI.PatchFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableWorkRoleLevelsAPI.PatchFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById`: TaxableWorkRoleLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxableWorkRoleLevelsAPI.PatchFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxableWorkRoleLevelId | 
**parentId** | **int32** | workRoleExemptionId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**TaxableWorkRoleLevel**](TaxableWorkRoleLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevels

> TaxableWorkRoleLevel PostFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevels(ctx, parentId, grandparentId).ClientId(clientId).TaxableWorkRoleLevel(taxableWorkRoleLevel).Execute()

Post TaxableWorkRoleLevel

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
	parentId := int32(56) // int32 | workRoleExemptionId
	grandparentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 
	taxableWorkRoleLevel := *openapiclient.NewTaxableWorkRoleLevel(*openapiclient.NewTaxCodeLevelReference()) // TaxableWorkRoleLevel | taxableWorkRoleLevel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxableWorkRoleLevelsAPI.PostFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevels(context.Background(), parentId, grandparentId).ClientId(clientId).TaxableWorkRoleLevel(taxableWorkRoleLevel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableWorkRoleLevelsAPI.PostFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevels`: TaxableWorkRoleLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxableWorkRoleLevelsAPI.PostFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | workRoleExemptionId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **taxableWorkRoleLevel** | [**TaxableWorkRoleLevel**](TaxableWorkRoleLevel.md) | taxableWorkRoleLevel | 

### Return type

[**TaxableWorkRoleLevel**](TaxableWorkRoleLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById

> TaxableWorkRoleLevel PutFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById(ctx, id, parentId, grandparentId).ClientId(clientId).TaxableWorkRoleLevel(taxableWorkRoleLevel).Execute()

Put TaxableWorkRoleLevel

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
	id := int32(56) // int32 | taxableWorkRoleLevelId
	parentId := int32(56) // int32 | workRoleExemptionId
	grandparentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 
	taxableWorkRoleLevel := *openapiclient.NewTaxableWorkRoleLevel(*openapiclient.NewTaxCodeLevelReference()) // TaxableWorkRoleLevel | taxableWorkRoleLevel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxableWorkRoleLevelsAPI.PutFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).TaxableWorkRoleLevel(taxableWorkRoleLevel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableWorkRoleLevelsAPI.PutFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById`: TaxableWorkRoleLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxableWorkRoleLevelsAPI.PutFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxableWorkRoleLevelId | 
**parentId** | **int32** | workRoleExemptionId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceTaxCodesByGrandparentIdWorkRoleExemptionsByParentIdTaxableWorkRoleLevelsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **taxableWorkRoleLevel** | [**TaxableWorkRoleLevel**](TaxableWorkRoleLevel.md) | taxableWorkRoleLevel | 

### Return type

[**TaxableWorkRoleLevel**](TaxableWorkRoleLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

