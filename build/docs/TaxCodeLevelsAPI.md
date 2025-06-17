# \TaxCodeLevelsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceTaxCodesByParentIdTaxCodeLevelsById**](TaxCodeLevelsAPI.md#DeleteFinanceTaxCodesByParentIdTaxCodeLevelsById) | **Delete** /finance/taxCodes/{parentId}/taxCodeLevels/{id} | Delete TaxCodeLevel
[**GetFinanceTaxCodesByParentIdTaxCodeLevels**](TaxCodeLevelsAPI.md#GetFinanceTaxCodesByParentIdTaxCodeLevels) | **Get** /finance/taxCodes/{parentId}/taxCodeLevels | Get List of TaxCodeLevel
[**GetFinanceTaxCodesByParentIdTaxCodeLevelsById**](TaxCodeLevelsAPI.md#GetFinanceTaxCodesByParentIdTaxCodeLevelsById) | **Get** /finance/taxCodes/{parentId}/taxCodeLevels/{id} | Get TaxCodeLevel
[**GetFinanceTaxCodesByParentIdTaxCodeLevelsCount**](TaxCodeLevelsAPI.md#GetFinanceTaxCodesByParentIdTaxCodeLevelsCount) | **Get** /finance/taxCodes/{parentId}/taxCodeLevels/count | Get Count of TaxCodeLevel
[**PatchFinanceTaxCodesByParentIdTaxCodeLevelsById**](TaxCodeLevelsAPI.md#PatchFinanceTaxCodesByParentIdTaxCodeLevelsById) | **Patch** /finance/taxCodes/{parentId}/taxCodeLevels/{id} | Patch TaxCodeLevel
[**PostFinanceTaxCodesByParentIdTaxCodeLevels**](TaxCodeLevelsAPI.md#PostFinanceTaxCodesByParentIdTaxCodeLevels) | **Post** /finance/taxCodes/{parentId}/taxCodeLevels | Post TaxCodeLevel
[**PutFinanceTaxCodesByParentIdTaxCodeLevelsById**](TaxCodeLevelsAPI.md#PutFinanceTaxCodesByParentIdTaxCodeLevelsById) | **Put** /finance/taxCodes/{parentId}/taxCodeLevels/{id} | Put TaxCodeLevel



## DeleteFinanceTaxCodesByParentIdTaxCodeLevelsById

> DeleteFinanceTaxCodesByParentIdTaxCodeLevelsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete TaxCodeLevel

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
	id := int32(56) // int32 | taxCodeLevelId
	parentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaxCodeLevelsAPI.DeleteFinanceTaxCodesByParentIdTaxCodeLevelsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeLevelsAPI.DeleteFinanceTaxCodesByParentIdTaxCodeLevelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxCodeLevelId | 
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceTaxCodesByParentIdTaxCodeLevelsByIdRequest struct via the builder pattern


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


## GetFinanceTaxCodesByParentIdTaxCodeLevels

> []TaxCodeLevel GetFinanceTaxCodesByParentIdTaxCodeLevels(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of TaxCodeLevel

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
	resp, r, err := apiClient.TaxCodeLevelsAPI.GetFinanceTaxCodesByParentIdTaxCodeLevels(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeLevelsAPI.GetFinanceTaxCodesByParentIdTaxCodeLevels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByParentIdTaxCodeLevels`: []TaxCodeLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeLevelsAPI.GetFinanceTaxCodesByParentIdTaxCodeLevels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByParentIdTaxCodeLevelsRequest struct via the builder pattern


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

[**[]TaxCodeLevel**](TaxCodeLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceTaxCodesByParentIdTaxCodeLevelsById

> TaxCodeLevel GetFinanceTaxCodesByParentIdTaxCodeLevelsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get TaxCodeLevel

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
	id := int32(56) // int32 | taxCodeLevelId
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
	resp, r, err := apiClient.TaxCodeLevelsAPI.GetFinanceTaxCodesByParentIdTaxCodeLevelsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeLevelsAPI.GetFinanceTaxCodesByParentIdTaxCodeLevelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByParentIdTaxCodeLevelsById`: TaxCodeLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeLevelsAPI.GetFinanceTaxCodesByParentIdTaxCodeLevelsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxCodeLevelId | 
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByParentIdTaxCodeLevelsByIdRequest struct via the builder pattern


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

[**TaxCodeLevel**](TaxCodeLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceTaxCodesByParentIdTaxCodeLevelsCount

> Count GetFinanceTaxCodesByParentIdTaxCodeLevelsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of TaxCodeLevel

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
	resp, r, err := apiClient.TaxCodeLevelsAPI.GetFinanceTaxCodesByParentIdTaxCodeLevelsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeLevelsAPI.GetFinanceTaxCodesByParentIdTaxCodeLevelsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByParentIdTaxCodeLevelsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeLevelsAPI.GetFinanceTaxCodesByParentIdTaxCodeLevelsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByParentIdTaxCodeLevelsCountRequest struct via the builder pattern


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


## PatchFinanceTaxCodesByParentIdTaxCodeLevelsById

> TaxCodeLevel PatchFinanceTaxCodesByParentIdTaxCodeLevelsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch TaxCodeLevel

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
	id := int32(56) // int32 | taxCodeLevelId
	parentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxCodeLevelsAPI.PatchFinanceTaxCodesByParentIdTaxCodeLevelsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeLevelsAPI.PatchFinanceTaxCodesByParentIdTaxCodeLevelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceTaxCodesByParentIdTaxCodeLevelsById`: TaxCodeLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeLevelsAPI.PatchFinanceTaxCodesByParentIdTaxCodeLevelsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxCodeLevelId | 
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceTaxCodesByParentIdTaxCodeLevelsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**TaxCodeLevel**](TaxCodeLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceTaxCodesByParentIdTaxCodeLevels

> TaxCodeLevel PostFinanceTaxCodesByParentIdTaxCodeLevels(ctx, parentId).ClientId(clientId).TaxCodeLevel(taxCodeLevel).Execute()

Post TaxCodeLevel

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
	taxCodeLevel := *openapiclient.NewTaxCodeLevel(NullableFloat64(123), "RateType_example") // TaxCodeLevel | taxCodeLevel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxCodeLevelsAPI.PostFinanceTaxCodesByParentIdTaxCodeLevels(context.Background(), parentId).ClientId(clientId).TaxCodeLevel(taxCodeLevel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeLevelsAPI.PostFinanceTaxCodesByParentIdTaxCodeLevels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceTaxCodesByParentIdTaxCodeLevels`: TaxCodeLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeLevelsAPI.PostFinanceTaxCodesByParentIdTaxCodeLevels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceTaxCodesByParentIdTaxCodeLevelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **taxCodeLevel** | [**TaxCodeLevel**](TaxCodeLevel.md) | taxCodeLevel | 

### Return type

[**TaxCodeLevel**](TaxCodeLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceTaxCodesByParentIdTaxCodeLevelsById

> TaxCodeLevel PutFinanceTaxCodesByParentIdTaxCodeLevelsById(ctx, id, parentId).ClientId(clientId).TaxCodeLevel(taxCodeLevel).Execute()

Put TaxCodeLevel

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
	id := int32(56) // int32 | taxCodeLevelId
	parentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 
	taxCodeLevel := *openapiclient.NewTaxCodeLevel(NullableFloat64(123), "RateType_example") // TaxCodeLevel | taxCodeLevel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxCodeLevelsAPI.PutFinanceTaxCodesByParentIdTaxCodeLevelsById(context.Background(), id, parentId).ClientId(clientId).TaxCodeLevel(taxCodeLevel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeLevelsAPI.PutFinanceTaxCodesByParentIdTaxCodeLevelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceTaxCodesByParentIdTaxCodeLevelsById`: TaxCodeLevel
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeLevelsAPI.PutFinanceTaxCodesByParentIdTaxCodeLevelsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxCodeLevelId | 
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceTaxCodesByParentIdTaxCodeLevelsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **taxCodeLevel** | [**TaxCodeLevel**](TaxCodeLevel.md) | taxCodeLevel | 

### Return type

[**TaxCodeLevel**](TaxCodeLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

