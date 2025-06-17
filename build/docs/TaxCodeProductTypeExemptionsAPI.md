# \TaxCodeProductTypeExemptionsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceTaxCodesByParentIdProductTypeExemptionsById**](TaxCodeProductTypeExemptionsAPI.md#DeleteFinanceTaxCodesByParentIdProductTypeExemptionsById) | **Delete** /finance/taxCodes/{parentId}/productTypeExemptions/{id} | Delete ProductTypeExemption
[**GetFinanceTaxCodesByParentIdProductTypeExemptions**](TaxCodeProductTypeExemptionsAPI.md#GetFinanceTaxCodesByParentIdProductTypeExemptions) | **Get** /finance/taxCodes/{parentId}/productTypeExemptions | Get List of ProductTypeExemption
[**GetFinanceTaxCodesByParentIdProductTypeExemptionsById**](TaxCodeProductTypeExemptionsAPI.md#GetFinanceTaxCodesByParentIdProductTypeExemptionsById) | **Get** /finance/taxCodes/{parentId}/productTypeExemptions/{id} | Get ProductTypeExemption
[**GetFinanceTaxCodesByParentIdProductTypeExemptionsCount**](TaxCodeProductTypeExemptionsAPI.md#GetFinanceTaxCodesByParentIdProductTypeExemptionsCount) | **Get** /finance/taxCodes/{parentId}/productTypeExemptions/count | Get Count of ProductTypeExemption
[**PatchFinanceTaxCodesByParentIdProductTypeExemptionsById**](TaxCodeProductTypeExemptionsAPI.md#PatchFinanceTaxCodesByParentIdProductTypeExemptionsById) | **Patch** /finance/taxCodes/{parentId}/productTypeExemptions/{id} | Patch ProductTypeExemption
[**PostFinanceTaxCodesByParentIdProductTypeExemptions**](TaxCodeProductTypeExemptionsAPI.md#PostFinanceTaxCodesByParentIdProductTypeExemptions) | **Post** /finance/taxCodes/{parentId}/productTypeExemptions | Post ProductTypeExemption
[**PutFinanceTaxCodesByParentIdProductTypeExemptionsById**](TaxCodeProductTypeExemptionsAPI.md#PutFinanceTaxCodesByParentIdProductTypeExemptionsById) | **Put** /finance/taxCodes/{parentId}/productTypeExemptions/{id} | Put ProductTypeExemption



## DeleteFinanceTaxCodesByParentIdProductTypeExemptionsById

> DeleteFinanceTaxCodesByParentIdProductTypeExemptionsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ProductTypeExemption

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
	id := int32(56) // int32 | productTypeExemptionId
	parentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaxCodeProductTypeExemptionsAPI.DeleteFinanceTaxCodesByParentIdProductTypeExemptionsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeProductTypeExemptionsAPI.DeleteFinanceTaxCodesByParentIdProductTypeExemptionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | productTypeExemptionId | 
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceTaxCodesByParentIdProductTypeExemptionsByIdRequest struct via the builder pattern


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


## GetFinanceTaxCodesByParentIdProductTypeExemptions

> []ProductTypeExemption GetFinanceTaxCodesByParentIdProductTypeExemptions(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProductTypeExemption

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
	resp, r, err := apiClient.TaxCodeProductTypeExemptionsAPI.GetFinanceTaxCodesByParentIdProductTypeExemptions(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeProductTypeExemptionsAPI.GetFinanceTaxCodesByParentIdProductTypeExemptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByParentIdProductTypeExemptions`: []ProductTypeExemption
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeProductTypeExemptionsAPI.GetFinanceTaxCodesByParentIdProductTypeExemptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByParentIdProductTypeExemptionsRequest struct via the builder pattern


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

[**[]ProductTypeExemption**](ProductTypeExemption.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceTaxCodesByParentIdProductTypeExemptionsById

> ProductTypeExemption GetFinanceTaxCodesByParentIdProductTypeExemptionsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ProductTypeExemption

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
	id := int32(56) // int32 | productTypeExemptionId
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
	resp, r, err := apiClient.TaxCodeProductTypeExemptionsAPI.GetFinanceTaxCodesByParentIdProductTypeExemptionsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeProductTypeExemptionsAPI.GetFinanceTaxCodesByParentIdProductTypeExemptionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByParentIdProductTypeExemptionsById`: ProductTypeExemption
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeProductTypeExemptionsAPI.GetFinanceTaxCodesByParentIdProductTypeExemptionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | productTypeExemptionId | 
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByParentIdProductTypeExemptionsByIdRequest struct via the builder pattern


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

[**ProductTypeExemption**](ProductTypeExemption.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceTaxCodesByParentIdProductTypeExemptionsCount

> Count GetFinanceTaxCodesByParentIdProductTypeExemptionsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ProductTypeExemption

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
	resp, r, err := apiClient.TaxCodeProductTypeExemptionsAPI.GetFinanceTaxCodesByParentIdProductTypeExemptionsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeProductTypeExemptionsAPI.GetFinanceTaxCodesByParentIdProductTypeExemptionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceTaxCodesByParentIdProductTypeExemptionsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeProductTypeExemptionsAPI.GetFinanceTaxCodesByParentIdProductTypeExemptionsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceTaxCodesByParentIdProductTypeExemptionsCountRequest struct via the builder pattern


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


## PatchFinanceTaxCodesByParentIdProductTypeExemptionsById

> ProductTypeExemption PatchFinanceTaxCodesByParentIdProductTypeExemptionsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ProductTypeExemption

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
	id := int32(56) // int32 | productTypeExemptionId
	parentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxCodeProductTypeExemptionsAPI.PatchFinanceTaxCodesByParentIdProductTypeExemptionsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeProductTypeExemptionsAPI.PatchFinanceTaxCodesByParentIdProductTypeExemptionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceTaxCodesByParentIdProductTypeExemptionsById`: ProductTypeExemption
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeProductTypeExemptionsAPI.PatchFinanceTaxCodesByParentIdProductTypeExemptionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | productTypeExemptionId | 
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceTaxCodesByParentIdProductTypeExemptionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ProductTypeExemption**](ProductTypeExemption.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceTaxCodesByParentIdProductTypeExemptions

> ProductTypeExemption PostFinanceTaxCodesByParentIdProductTypeExemptions(ctx, parentId).ClientId(clientId).ProductTypeExemption(productTypeExemption).Execute()

Post ProductTypeExemption

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
	productTypeExemption := *openapiclient.NewProductTypeExemption(*openapiclient.NewProductTypeReference()) // ProductTypeExemption | productTypeExemption

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxCodeProductTypeExemptionsAPI.PostFinanceTaxCodesByParentIdProductTypeExemptions(context.Background(), parentId).ClientId(clientId).ProductTypeExemption(productTypeExemption).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeProductTypeExemptionsAPI.PostFinanceTaxCodesByParentIdProductTypeExemptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceTaxCodesByParentIdProductTypeExemptions`: ProductTypeExemption
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeProductTypeExemptionsAPI.PostFinanceTaxCodesByParentIdProductTypeExemptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceTaxCodesByParentIdProductTypeExemptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **productTypeExemption** | [**ProductTypeExemption**](ProductTypeExemption.md) | productTypeExemption | 

### Return type

[**ProductTypeExemption**](ProductTypeExemption.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceTaxCodesByParentIdProductTypeExemptionsById

> ProductTypeExemption PutFinanceTaxCodesByParentIdProductTypeExemptionsById(ctx, id, parentId).ClientId(clientId).ProductTypeExemption(productTypeExemption).Execute()

Put ProductTypeExemption

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
	id := int32(56) // int32 | productTypeExemptionId
	parentId := int32(56) // int32 | taxCodeId
	clientId := "clientId_example" // string | 
	productTypeExemption := *openapiclient.NewProductTypeExemption(*openapiclient.NewProductTypeReference()) // ProductTypeExemption | productTypeExemption

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxCodeProductTypeExemptionsAPI.PutFinanceTaxCodesByParentIdProductTypeExemptionsById(context.Background(), id, parentId).ClientId(clientId).ProductTypeExemption(productTypeExemption).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxCodeProductTypeExemptionsAPI.PutFinanceTaxCodesByParentIdProductTypeExemptionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceTaxCodesByParentIdProductTypeExemptionsById`: ProductTypeExemption
	fmt.Fprintf(os.Stdout, "Response from `TaxCodeProductTypeExemptionsAPI.PutFinanceTaxCodesByParentIdProductTypeExemptionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | productTypeExemptionId | 
**parentId** | **int32** | taxCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceTaxCodesByParentIdProductTypeExemptionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **productTypeExemption** | [**ProductTypeExemption**](ProductTypeExemption.md) | productTypeExemption | 

### Return type

[**ProductTypeExemption**](ProductTypeExemption.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

