# \BillingTermsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceBillingTermsById**](BillingTermsAPI.md#DeleteFinanceBillingTermsById) | **Delete** /finance/billingTerms/{id} | Delete BillingTerm
[**GetFinanceBillingTerms**](BillingTermsAPI.md#GetFinanceBillingTerms) | **Get** /finance/billingTerms | Get List of BillingTerm
[**GetFinanceBillingTermsById**](BillingTermsAPI.md#GetFinanceBillingTermsById) | **Get** /finance/billingTerms/{id} | Get BillingTerm
[**GetFinanceBillingTermsByIdUsages**](BillingTermsAPI.md#GetFinanceBillingTermsByIdUsages) | **Get** /finance/billingTerms/{id}/usages | Get List of Usage Count
[**GetFinanceBillingTermsByIdUsagesList**](BillingTermsAPI.md#GetFinanceBillingTermsByIdUsagesList) | **Get** /finance/billingTerms/{id}/usages/list | Get List of Usage
[**GetFinanceBillingTermsCount**](BillingTermsAPI.md#GetFinanceBillingTermsCount) | **Get** /finance/billingTerms/count | Get Count of BillingTerm
[**PatchFinanceBillingTermsById**](BillingTermsAPI.md#PatchFinanceBillingTermsById) | **Patch** /finance/billingTerms/{id} | Patch BillingTerm
[**PostFinanceBillingTerms**](BillingTermsAPI.md#PostFinanceBillingTerms) | **Post** /finance/billingTerms | Post BillingTerm
[**PutFinanceBillingTermsById**](BillingTermsAPI.md#PutFinanceBillingTermsById) | **Put** /finance/billingTerms/{id} | Put BillingTerm



## DeleteFinanceBillingTermsById

> DeleteFinanceBillingTermsById(ctx, id).ClientId(clientId).Execute()

Delete BillingTerm

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
	id := int32(56) // int32 | billingTermId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BillingTermsAPI.DeleteFinanceBillingTermsById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingTermsAPI.DeleteFinanceBillingTermsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | billingTermId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceBillingTermsByIdRequest struct via the builder pattern


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


## GetFinanceBillingTerms

> []BillingTerm GetFinanceBillingTerms(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of BillingTerm

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
	resp, r, err := apiClient.BillingTermsAPI.GetFinanceBillingTerms(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingTermsAPI.GetFinanceBillingTerms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceBillingTerms`: []BillingTerm
	fmt.Fprintf(os.Stdout, "Response from `BillingTermsAPI.GetFinanceBillingTerms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceBillingTermsRequest struct via the builder pattern


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

[**[]BillingTerm**](BillingTerm.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceBillingTermsById

> BillingTerm GetFinanceBillingTermsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get BillingTerm

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
	id := int32(56) // int32 | billingTermId
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
	resp, r, err := apiClient.BillingTermsAPI.GetFinanceBillingTermsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingTermsAPI.GetFinanceBillingTermsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceBillingTermsById`: BillingTerm
	fmt.Fprintf(os.Stdout, "Response from `BillingTermsAPI.GetFinanceBillingTermsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | billingTermId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceBillingTermsByIdRequest struct via the builder pattern


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

[**BillingTerm**](BillingTerm.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceBillingTermsByIdUsages

> []Usage GetFinanceBillingTermsByIdUsages(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of Usage Count

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
	id := int32(56) // int32 | billingTermId
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
	resp, r, err := apiClient.BillingTermsAPI.GetFinanceBillingTermsByIdUsages(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingTermsAPI.GetFinanceBillingTermsByIdUsages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceBillingTermsByIdUsages`: []Usage
	fmt.Fprintf(os.Stdout, "Response from `BillingTermsAPI.GetFinanceBillingTermsByIdUsages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | billingTermId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceBillingTermsByIdUsagesRequest struct via the builder pattern


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

[**[]Usage**](Usage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceBillingTermsByIdUsagesList

> []Usage GetFinanceBillingTermsByIdUsagesList(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of Usage

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
	id := int32(56) // int32 | billingTermId
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
	resp, r, err := apiClient.BillingTermsAPI.GetFinanceBillingTermsByIdUsagesList(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingTermsAPI.GetFinanceBillingTermsByIdUsagesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceBillingTermsByIdUsagesList`: []Usage
	fmt.Fprintf(os.Stdout, "Response from `BillingTermsAPI.GetFinanceBillingTermsByIdUsagesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | billingTermId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceBillingTermsByIdUsagesListRequest struct via the builder pattern


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

[**[]Usage**](Usage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceBillingTermsCount

> Count GetFinanceBillingTermsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of BillingTerm

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
	resp, r, err := apiClient.BillingTermsAPI.GetFinanceBillingTermsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingTermsAPI.GetFinanceBillingTermsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceBillingTermsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `BillingTermsAPI.GetFinanceBillingTermsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceBillingTermsCountRequest struct via the builder pattern


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


## PatchFinanceBillingTermsById

> BillingTerm PatchFinanceBillingTermsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch BillingTerm

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
	id := int32(56) // int32 | billingTermId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingTermsAPI.PatchFinanceBillingTermsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingTermsAPI.PatchFinanceBillingTermsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceBillingTermsById`: BillingTerm
	fmt.Fprintf(os.Stdout, "Response from `BillingTermsAPI.PatchFinanceBillingTermsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | billingTermId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceBillingTermsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**BillingTerm**](BillingTerm.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceBillingTerms

> BillingTerm PostFinanceBillingTerms(ctx).ClientId(clientId).BillingTerm(billingTerm).Execute()

Post BillingTerm

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
	clientId := "clientId_example" // string | 
	billingTerm := *openapiclient.NewBillingTerm("Name_example", NullableInt32(123)) // BillingTerm | billingTerms

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingTermsAPI.PostFinanceBillingTerms(context.Background()).ClientId(clientId).BillingTerm(billingTerm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingTermsAPI.PostFinanceBillingTerms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceBillingTerms`: BillingTerm
	fmt.Fprintf(os.Stdout, "Response from `BillingTermsAPI.PostFinanceBillingTerms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceBillingTermsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **billingTerm** | [**BillingTerm**](BillingTerm.md) | billingTerms | 

### Return type

[**BillingTerm**](BillingTerm.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceBillingTermsById

> BillingTerm PutFinanceBillingTermsById(ctx, id).ClientId(clientId).BillingTerm(billingTerm).Execute()

Put BillingTerm

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
	id := int32(56) // int32 | billingTermId
	clientId := "clientId_example" // string | 
	billingTerm := *openapiclient.NewBillingTerm("Name_example", NullableInt32(123)) // BillingTerm | billingTerms

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingTermsAPI.PutFinanceBillingTermsById(context.Background(), id).ClientId(clientId).BillingTerm(billingTerm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingTermsAPI.PutFinanceBillingTermsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceBillingTermsById`: BillingTerm
	fmt.Fprintf(os.Stdout, "Response from `BillingTermsAPI.PutFinanceBillingTermsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | billingTermId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceBillingTermsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **billingTerm** | [**BillingTerm**](BillingTerm.md) | billingTerms | 

### Return type

[**BillingTerm**](BillingTerm.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

