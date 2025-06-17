# \TaxCodeXRefsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceTaxCodesByParentIdTaxCodeXRefsById**](TaxCodeXRefsAPI.md#DeleteFinanceTaxCodesByParentIdTaxCodeXRefsById) | **Delete** /finance/taxCodes/{parentId}/taxCodeXRefs/{id} | Delete TaxCodeXRef
[**GetFinanceTaxCodesByParentIdTaxCodeXRefs**](TaxCodeXRefsAPI.md#GetFinanceTaxCodesByParentIdTaxCodeXRefs) | **Get** /finance/taxCodes/{parentId}/taxCodeXRefs | Get List of TaxCodeXRef
[**GetFinanceTaxCodesByParentIdTaxCodeXRefsById**](TaxCodeXRefsAPI.md#GetFinanceTaxCodesByParentIdTaxCodeXRefsById) | **Get** /finance/taxCodes/{parentId}/taxCodeXRefs/{id} | Get TaxCodeXRef
[**GetFinanceTaxCodesByParentIdTaxCodeXRefsCount**](TaxCodeXRefsAPI.md#GetFinanceTaxCodesByParentIdTaxCodeXRefsCount) | **Get** /finance/taxCodes/{parentId}/taxCodeXRefs/count | Get Count of TaxCodeXRef
[**PatchFinanceTaxCodesByParentIdTaxCodeXRefsById**](TaxCodeXRefsAPI.md#PatchFinanceTaxCodesByParentIdTaxCodeXRefsById) | **Patch** /finance/taxCodes/{parentId}/taxCodeXRefs/{id} | Patch TaxCodeXRef
[**PostFinanceTaxCodesByParentIdTaxCodeXRefs**](TaxCodeXRefsAPI.md#PostFinanceTaxCodesByParentIdTaxCodeXRefs) | **Post** /finance/taxCodes/{parentId}/taxCodeXRefs | Post TaxCodeXRef
[**PutFinanceTaxCodesByParentIdTaxCodeXRefsById**](TaxCodeXRefsAPI.md#PutFinanceTaxCodesByParentIdTaxCodeXRefsById) | **Put** /finance/taxCodes/{parentId}/taxCodeXRefs/{id} | Put TaxCodeXRef



## DeleteFinanceTaxCodesByParentIdTaxCodeXRefsById

> DeleteFinanceTaxCodesByParentIdTaxCodeXRefsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete TaxCodeXRef

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
	id := int32(56) // int32 | taxCodeXRefId
	parentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaxCodeXRefsAPI.DeleteFinanceTaxCodesByParentIdTaxCodeXRefsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeXRefsAPI.DeleteFinanceTaxCodesByParentIdTaxCodeXRefsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxCodeXRefId | 
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceTaxCodesByParentIdTaxCodeXRefsByIdRequest struct via the builder pattern


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


## GetFinanceTaxCodesByParentIdTaxCodeXRefs

> []TaxCodeXRef GetFinanceTaxCodesByParentIdTaxCodeXRefs(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of TaxCodeXRef

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
	resp, r, err := apiClient.TaxCodeXRefsAPI.GetFinanceTaxCodesByParentIdTaxCodeXRefs(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeXRefsAPI.GetFinanceTaxCodesByParentIdTaxCodeXRefs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByParentIdTaxCodeXRefs`: []TaxCodeXRef
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeXRefsAPI.GetFinanceTaxCodesByParentIdTaxCodeXRefs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByParentIdTaxCodeXRefsRequest struct via the builder pattern


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

[**[]TaxCodeXRef**](TaxCodeXRef.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceTaxCodesByParentIdTaxCodeXRefsById

> TaxCodeXRef GetFinanceTaxCodesByParentIdTaxCodeXRefsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get TaxCodeXRef

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
	id := int32(56) // int32 | taxCodeXRefId
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
	resp, r, err := apiClient.TaxCodeXRefsAPI.GetFinanceTaxCodesByParentIdTaxCodeXRefsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeXRefsAPI.GetFinanceTaxCodesByParentIdTaxCodeXRefsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByParentIdTaxCodeXRefsById`: TaxCodeXRef
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeXRefsAPI.GetFinanceTaxCodesByParentIdTaxCodeXRefsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxCodeXRefId | 
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByParentIdTaxCodeXRefsByIdRequest struct via the builder pattern


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

[**TaxCodeXRef**](TaxCodeXRef.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceTaxCodesByParentIdTaxCodeXRefsCount

> Count GetFinanceTaxCodesByParentIdTaxCodeXRefsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of TaxCodeXRef

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
	resp, r, err := apiClient.TaxCodeXRefsAPI.GetFinanceTaxCodesByParentIdTaxCodeXRefsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeXRefsAPI.GetFinanceTaxCodesByParentIdTaxCodeXRefsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByParentIdTaxCodeXRefsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeXRefsAPI.GetFinanceTaxCodesByParentIdTaxCodeXRefsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByParentIdTaxCodeXRefsCountRequest struct via the builder pattern


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


## PatchFinanceTaxCodesByParentIdTaxCodeXRefsById

> TaxCodeXRef PatchFinanceTaxCodesByParentIdTaxCodeXRefsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch TaxCodeXRef

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
	id := int32(56) // int32 | taxCodeXRefId
	parentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxCodeXRefsAPI.PatchFinanceTaxCodesByParentIdTaxCodeXRefsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeXRefsAPI.PatchFinanceTaxCodesByParentIdTaxCodeXRefsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceTaxCodesByParentIdTaxCodeXRefsById`: TaxCodeXRef
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeXRefsAPI.PatchFinanceTaxCodesByParentIdTaxCodeXRefsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxCodeXRefId | 
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceTaxCodesByParentIdTaxCodeXRefsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**TaxCodeXRef**](TaxCodeXRef.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceTaxCodesByParentIdTaxCodeXRefs

> TaxCodeXRef PostFinanceTaxCodesByParentIdTaxCodeXRefs(ctx, parentId).ClientId(clientId).TaxCodeXRef(taxCodeXRef).Execute()

Post TaxCodeXRef

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
	taxCodeXRef := *openapiclient.NewTaxCodeXRef("Description_example") // TaxCodeXRef | taxCodeXRef

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxCodeXRefsAPI.PostFinanceTaxCodesByParentIdTaxCodeXRefs(context.Background(), parentId).ClientId(clientId).TaxCodeXRef(taxCodeXRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeXRefsAPI.PostFinanceTaxCodesByParentIdTaxCodeXRefs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceTaxCodesByParentIdTaxCodeXRefs`: TaxCodeXRef
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeXRefsAPI.PostFinanceTaxCodesByParentIdTaxCodeXRefs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceTaxCodesByParentIdTaxCodeXRefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **taxCodeXRef** | [**TaxCodeXRef**](TaxCodeXRef.md) | taxCodeXRef | 

### Return type

[**TaxCodeXRef**](TaxCodeXRef.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceTaxCodesByParentIdTaxCodeXRefsById

> TaxCodeXRef PutFinanceTaxCodesByParentIdTaxCodeXRefsById(ctx, id, parentId).ClientId(clientId).TaxCodeXRef(taxCodeXRef).Execute()

Put TaxCodeXRef

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
	id := int32(56) // int32 | taxCodeXRefId
	parentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 
	taxCodeXRef := *openapiclient.NewTaxCodeXRef("Description_example") // TaxCodeXRef | taxCodeXRef

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxCodeXRefsAPI.PutFinanceTaxCodesByParentIdTaxCodeXRefsById(context.Background(), id, parentId).ClientId(clientId).TaxCodeXRef(taxCodeXRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeXRefsAPI.PutFinanceTaxCodesByParentIdTaxCodeXRefsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceTaxCodesByParentIdTaxCodeXRefsById`: TaxCodeXRef
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeXRefsAPI.PutFinanceTaxCodesByParentIdTaxCodeXRefsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | taxCodeXRefId | 
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceTaxCodesByParentIdTaxCodeXRefsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **taxCodeXRef** | [**TaxCodeXRef**](TaxCodeXRef.md) | taxCodeXRef | 

### Return type

[**TaxCodeXRef**](TaxCodeXRef.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

