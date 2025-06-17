# \InvoiceEmailTemplatesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceInvoiceEmailTemplatesById**](InvoiceEmailTemplatesAPI.md#DeleteFinanceInvoiceEmailTemplatesById) | **Delete** /finance/invoiceEmailTemplates/{id} | Delete InvoiceEmailTemplate
[**GetFinanceInvoiceEmailTemplates**](InvoiceEmailTemplatesAPI.md#GetFinanceInvoiceEmailTemplates) | **Get** /finance/invoiceEmailTemplates | Get List of InvoiceEmailTemplate
[**GetFinanceInvoiceEmailTemplatesById**](InvoiceEmailTemplatesAPI.md#GetFinanceInvoiceEmailTemplatesById) | **Get** /finance/invoiceEmailTemplates/{id} | Get InvoiceEmailTemplate
[**GetFinanceInvoiceEmailTemplatesByIdUsages**](InvoiceEmailTemplatesAPI.md#GetFinanceInvoiceEmailTemplatesByIdUsages) | **Get** /finance/invoiceEmailTemplates/{id}/usages | Get List of Usage Count
[**GetFinanceInvoiceEmailTemplatesByIdUsagesList**](InvoiceEmailTemplatesAPI.md#GetFinanceInvoiceEmailTemplatesByIdUsagesList) | **Get** /finance/invoiceEmailTemplates/{id}/usages/list | Get List of Usage
[**GetFinanceInvoiceEmailTemplatesCount**](InvoiceEmailTemplatesAPI.md#GetFinanceInvoiceEmailTemplatesCount) | **Get** /finance/invoiceEmailTemplates/count | Get Count of InvoiceEmailTemplate
[**PatchFinanceInvoiceEmailTemplatesById**](InvoiceEmailTemplatesAPI.md#PatchFinanceInvoiceEmailTemplatesById) | **Patch** /finance/invoiceEmailTemplates/{id} | Patch InvoiceEmailTemplate
[**PostFinanceInvoiceEmailTemplates**](InvoiceEmailTemplatesAPI.md#PostFinanceInvoiceEmailTemplates) | **Post** /finance/invoiceEmailTemplates | Post InvoiceEmailTemplate
[**PutFinanceInvoiceEmailTemplatesById**](InvoiceEmailTemplatesAPI.md#PutFinanceInvoiceEmailTemplatesById) | **Put** /finance/invoiceEmailTemplates/{id} | Put InvoiceEmailTemplate



## DeleteFinanceInvoiceEmailTemplatesById

> DeleteFinanceInvoiceEmailTemplatesById(ctx, id).ClientId(clientId).Execute()

Delete InvoiceEmailTemplate

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
	id := int32(56) // int32 | invoiceEmailTemplateId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InvoiceEmailTemplatesAPI.DeleteFinanceInvoiceEmailTemplatesById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceEmailTemplatesAPI.DeleteFinanceInvoiceEmailTemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | invoiceEmailTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceInvoiceEmailTemplatesByIdRequest struct via the builder pattern


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


## GetFinanceInvoiceEmailTemplates

> []InvoiceEmailTemplate GetFinanceInvoiceEmailTemplates(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of InvoiceEmailTemplate

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
	resp, r, err := apiClient.InvoiceEmailTemplatesAPI.GetFinanceInvoiceEmailTemplates(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceEmailTemplatesAPI.GetFinanceInvoiceEmailTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceInvoiceEmailTemplates`: []InvoiceEmailTemplate
	fmt.Fprintf(os.Stdout, "Response from `InvoiceEmailTemplatesAPI.GetFinanceInvoiceEmailTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceInvoiceEmailTemplatesRequest struct via the builder pattern


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

[**[]InvoiceEmailTemplate**](InvoiceEmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceInvoiceEmailTemplatesById

> InvoiceEmailTemplate GetFinanceInvoiceEmailTemplatesById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get InvoiceEmailTemplate

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
	id := int32(56) // int32 | invoiceEmailTemplateId
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
	resp, r, err := apiClient.InvoiceEmailTemplatesAPI.GetFinanceInvoiceEmailTemplatesById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceEmailTemplatesAPI.GetFinanceInvoiceEmailTemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceInvoiceEmailTemplatesById`: InvoiceEmailTemplate
	fmt.Fprintf(os.Stdout, "Response from `InvoiceEmailTemplatesAPI.GetFinanceInvoiceEmailTemplatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | invoiceEmailTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceInvoiceEmailTemplatesByIdRequest struct via the builder pattern


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

[**InvoiceEmailTemplate**](InvoiceEmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceInvoiceEmailTemplatesByIdUsages

> []Usage GetFinanceInvoiceEmailTemplatesByIdUsages(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

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
	id := int32(56) // int32 | invoiceEmailTemplateId
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
	resp, r, err := apiClient.InvoiceEmailTemplatesAPI.GetFinanceInvoiceEmailTemplatesByIdUsages(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceEmailTemplatesAPI.GetFinanceInvoiceEmailTemplatesByIdUsages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceInvoiceEmailTemplatesByIdUsages`: []Usage
	fmt.Fprintf(os.Stdout, "Response from `InvoiceEmailTemplatesAPI.GetFinanceInvoiceEmailTemplatesByIdUsages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | invoiceEmailTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceInvoiceEmailTemplatesByIdUsagesRequest struct via the builder pattern


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


## GetFinanceInvoiceEmailTemplatesByIdUsagesList

> []Usage GetFinanceInvoiceEmailTemplatesByIdUsagesList(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

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
	id := int32(56) // int32 | invoiceEmailTemplateId
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
	resp, r, err := apiClient.InvoiceEmailTemplatesAPI.GetFinanceInvoiceEmailTemplatesByIdUsagesList(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceEmailTemplatesAPI.GetFinanceInvoiceEmailTemplatesByIdUsagesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceInvoiceEmailTemplatesByIdUsagesList`: []Usage
	fmt.Fprintf(os.Stdout, "Response from `InvoiceEmailTemplatesAPI.GetFinanceInvoiceEmailTemplatesByIdUsagesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | invoiceEmailTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceInvoiceEmailTemplatesByIdUsagesListRequest struct via the builder pattern


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


## GetFinanceInvoiceEmailTemplatesCount

> Count GetFinanceInvoiceEmailTemplatesCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of InvoiceEmailTemplate

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
	resp, r, err := apiClient.InvoiceEmailTemplatesAPI.GetFinanceInvoiceEmailTemplatesCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceEmailTemplatesAPI.GetFinanceInvoiceEmailTemplatesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceInvoiceEmailTemplatesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `InvoiceEmailTemplatesAPI.GetFinanceInvoiceEmailTemplatesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceInvoiceEmailTemplatesCountRequest struct via the builder pattern


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


## PatchFinanceInvoiceEmailTemplatesById

> InvoiceEmailTemplate PatchFinanceInvoiceEmailTemplatesById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch InvoiceEmailTemplate

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
	id := int32(56) // int32 | invoiceEmailTemplateId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceEmailTemplatesAPI.PatchFinanceInvoiceEmailTemplatesById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceEmailTemplatesAPI.PatchFinanceInvoiceEmailTemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceInvoiceEmailTemplatesById`: InvoiceEmailTemplate
	fmt.Fprintf(os.Stdout, "Response from `InvoiceEmailTemplatesAPI.PatchFinanceInvoiceEmailTemplatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | invoiceEmailTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceInvoiceEmailTemplatesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**InvoiceEmailTemplate**](InvoiceEmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceInvoiceEmailTemplates

> InvoiceEmailTemplate PostFinanceInvoiceEmailTemplates(ctx).ClientId(clientId).InvoiceEmailTemplate(invoiceEmailTemplate).Execute()

Post InvoiceEmailTemplate

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
	invoiceEmailTemplate := *openapiclient.NewInvoiceEmailTemplate("Name_example", "Subject_example") // InvoiceEmailTemplate | invoiceEmailTemplate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceEmailTemplatesAPI.PostFinanceInvoiceEmailTemplates(context.Background()).ClientId(clientId).InvoiceEmailTemplate(invoiceEmailTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceEmailTemplatesAPI.PostFinanceInvoiceEmailTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceInvoiceEmailTemplates`: InvoiceEmailTemplate
	fmt.Fprintf(os.Stdout, "Response from `InvoiceEmailTemplatesAPI.PostFinanceInvoiceEmailTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceInvoiceEmailTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **invoiceEmailTemplate** | [**InvoiceEmailTemplate**](InvoiceEmailTemplate.md) | invoiceEmailTemplate | 

### Return type

[**InvoiceEmailTemplate**](InvoiceEmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceInvoiceEmailTemplatesById

> InvoiceEmailTemplate PutFinanceInvoiceEmailTemplatesById(ctx, id).ClientId(clientId).InvoiceEmailTemplate(invoiceEmailTemplate).Execute()

Put InvoiceEmailTemplate

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
	id := int32(56) // int32 | invoiceEmailTemplateId
	clientId := "clientId_example" // string | 
	invoiceEmailTemplate := *openapiclient.NewInvoiceEmailTemplate("Name_example", "Subject_example") // InvoiceEmailTemplate | invoiceEmailTemplate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceEmailTemplatesAPI.PutFinanceInvoiceEmailTemplatesById(context.Background(), id).ClientId(clientId).InvoiceEmailTemplate(invoiceEmailTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceEmailTemplatesAPI.PutFinanceInvoiceEmailTemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceInvoiceEmailTemplatesById`: InvoiceEmailTemplate
	fmt.Fprintf(os.Stdout, "Response from `InvoiceEmailTemplatesAPI.PutFinanceInvoiceEmailTemplatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | invoiceEmailTemplateId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceInvoiceEmailTemplatesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **invoiceEmailTemplate** | [**InvoiceEmailTemplate**](InvoiceEmailTemplate.md) | invoiceEmailTemplate | 

### Return type

[**InvoiceEmailTemplate**](InvoiceEmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

