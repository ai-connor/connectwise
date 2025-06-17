# \TaxableXRefLevelsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById**](TaxableXRefLevelsAPI.md#DeleteFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById) | **Delete** /finance/taxCodes/{grandparentId}/taxCodeXRefs/{parentId}/taxableXRefLevels/{id} | Delete TaxableXRefLevel
[**GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevels**](TaxableXRefLevelsAPI.md#GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevels) | **Get** /finance/taxCodes/{grandparentId}/taxCodeXRefs/{parentId}/taxableXRefLevels | Get List of TaxableXRefLevel
[**GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById**](TaxableXRefLevelsAPI.md#GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById) | **Get** /finance/taxCodes/{grandparentId}/taxCodeXRefs/{parentId}/taxableXRefLevels/{id} | Get TaxableXRefLevel
[**GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsCount**](TaxableXRefLevelsAPI.md#GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsCount) | **Get** /finance/taxCodes/{grandparentId}/taxCodeXRefs/{parentId}/taxableXRefLevels/count | Get Count of TaxableXRefLevel
[**PatchFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById**](TaxableXRefLevelsAPI.md#PatchFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById) | **Patch** /finance/taxCodes/{grandparentId}/taxCodeXRefs/{parentId}/taxableXRefLevels/{id} | Patch TaxableXRefLevel
[**PostFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevels**](TaxableXRefLevelsAPI.md#PostFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevels) | **Post** /finance/taxCodes/{grandparentId}/taxCodeXRefs/{parentId}/taxableXRefLevels | Post TaxableXRefLevel
[**PutFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById**](TaxableXRefLevelsAPI.md#PutFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById) | **Put** /finance/taxCodes/{grandparentId}/taxCodeXRefs/{parentId}/taxableXRefLevels/{id} | Put TaxableXRefLevel



## DeleteFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById

> DeleteFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById(ctx, id, parentId, grandparentId).ClientId(clientId).Execute()

Delete TaxableXRefLevel

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
	id := int32(56) // int32 | taxableXRefLevelId
	parentId := int32(56) // int32 | taxCodeXRefId
	grandparentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaxableXRefLevelsAPI.DeleteFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableXRefLevelsAPI.DeleteFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxableXRefLevelId | 
**parentId** | **int32** | taxCodeXRefId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsByIdRequest struct via the builder pattern


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


## GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevels

> []TaxableXRefLevel GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevels(ctx, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of TaxableXRefLevel

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
	parentId := int32(56) // int32 | taxCodeXRefId
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
	resp, r, err := apiClient.TaxableXRefLevelsAPI.GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevels(context.Background(), parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableXRefLevelsAPI.GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevels`: []TaxableXRefLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxableXRefLevelsAPI.GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | taxCodeXRefId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsRequest struct via the builder pattern


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

[**[]TaxableXRefLevel**](TaxableXRefLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById

> TaxableXRefLevel GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById(ctx, id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get TaxableXRefLevel

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
	id := int32(56) // int32 | taxableXRefLevelId
	parentId := int32(56) // int32 | taxCodeXRefId
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
	resp, r, err := apiClient.TaxableXRefLevelsAPI.GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableXRefLevelsAPI.GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById`: TaxableXRefLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxableXRefLevelsAPI.GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxableXRefLevelId | 
**parentId** | **int32** | taxCodeXRefId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsByIdRequest struct via the builder pattern


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

[**TaxableXRefLevel**](TaxableXRefLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsCount

> Count GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsCount(ctx, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of TaxableXRefLevel

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
	parentId := int32(56) // int32 | taxCodeXRefId
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
	resp, r, err := apiClient.TaxableXRefLevelsAPI.GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsCount(context.Background(), parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableXRefLevelsAPI.GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `TaxableXRefLevelsAPI.GetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | taxCodeXRefId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsCountRequest struct via the builder pattern


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


## PatchFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById

> TaxableXRefLevel PatchFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById(ctx, id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch TaxableXRefLevel

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
	id := int32(56) // int32 | taxableXRefLevelId
	parentId := int32(56) // int32 | taxCodeXRefId
	grandparentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxableXRefLevelsAPI.PatchFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableXRefLevelsAPI.PatchFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById`: TaxableXRefLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxableXRefLevelsAPI.PatchFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxableXRefLevelId | 
**parentId** | **int32** | taxCodeXRefId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**TaxableXRefLevel**](TaxableXRefLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevels

> TaxableXRefLevel PostFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevels(ctx, parentId, grandparentId).ClientId(clientId).TaxableXRefLevel(taxableXRefLevel).Execute()

Post TaxableXRefLevel

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
	parentId := int32(56) // int32 | taxCodeXRefId
	grandparentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 
	taxableXRefLevel := *openapiclient.NewTaxableXRefLevel(*openapiclient.NewTaxCodeLevelReference()) // TaxableXRefLevel | taxableXRefLevel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxableXRefLevelsAPI.PostFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevels(context.Background(), parentId, grandparentId).ClientId(clientId).TaxableXRefLevel(taxableXRefLevel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableXRefLevelsAPI.PostFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevels`: TaxableXRefLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxableXRefLevelsAPI.PostFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | taxCodeXRefId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **taxableXRefLevel** | [**TaxableXRefLevel**](TaxableXRefLevel.md) | taxableXRefLevel | 

### Return type

[**TaxableXRefLevel**](TaxableXRefLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById

> TaxableXRefLevel PutFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById(ctx, id, parentId, grandparentId).ClientId(clientId).TaxableXRefLevel(taxableXRefLevel).Execute()

Put TaxableXRefLevel

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
	id := int32(56) // int32 | taxableXRefLevelId
	parentId := int32(56) // int32 | taxCodeXRefId
	grandparentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 
	taxableXRefLevel := *openapiclient.NewTaxableXRefLevel(*openapiclient.NewTaxCodeLevelReference()) // TaxableXRefLevel | taxableXRefLevel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxableXRefLevelsAPI.PutFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById(context.Background(), id, parentId, grandparentId).ClientId(clientId).TaxableXRefLevel(taxableXRefLevel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxableXRefLevelsAPI.PutFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById`: TaxableXRefLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxableXRefLevelsAPI.PutFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxableXRefLevelId | 
**parentId** | **int32** | taxCodeXRefId | 
**grandparentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceTaxCodesByGrandparentIdTaxCodeXRefsByParentIdTaxableXRefLevelsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **taxableXRefLevel** | [**TaxableXRefLevel**](TaxableXRefLevel.md) | taxableXRefLevel | 

### Return type

[**TaxableXRefLevel**](TaxableXRefLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

