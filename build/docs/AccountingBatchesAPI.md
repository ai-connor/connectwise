# \AccountingBatchesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceAccountingBatchesById**](AccountingBatchesAPI.md#DeleteFinanceAccountingBatchesById) | **Delete** /finance/accounting/batches/{id} | Delete GLExport
[**GetFinanceAccountingBatches**](AccountingBatchesAPI.md#GetFinanceAccountingBatches) | **Get** /finance/accounting/batches | Get List of AccountingBatch
[**GetFinanceAccountingBatchesById**](AccountingBatchesAPI.md#GetFinanceAccountingBatchesById) | **Get** /finance/accounting/batches/{id} | Get AccountingBatch
[**GetFinanceAccountingBatchesCount**](AccountingBatchesAPI.md#GetFinanceAccountingBatchesCount) | **Get** /finance/accounting/batches/count | Get Count of AccountingBatch
[**PostFinanceAccountingBatches**](AccountingBatchesAPI.md#PostFinanceAccountingBatches) | **Post** /finance/accounting/batches | Post GLExport
[**PostFinanceAccountingBatchesByIdExport**](AccountingBatchesAPI.md#PostFinanceAccountingBatchesByIdExport) | **Post** /finance/accounting/batches/{id}/export | Post GLExport
[**PostFinanceAccountingExport**](AccountingBatchesAPI.md#PostFinanceAccountingExport) | **Post** /finance/accounting/export | Post GLExport



## DeleteFinanceAccountingBatchesById

> DeleteFinanceAccountingBatchesById(ctx, id).ClientId(clientId).Execute()

Delete GLExport

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
	id := int32(56) // int32 | batcheId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountingBatchesAPI.DeleteFinanceAccountingBatchesById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingBatchesAPI.DeleteFinanceAccountingBatchesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | batcheId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceAccountingBatchesByIdRequest struct via the builder pattern


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


## GetFinanceAccountingBatches

> []AccountingBatch GetFinanceAccountingBatches(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of AccountingBatch

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
	resp, r, err := apiClient.AccountingBatchesAPI.GetFinanceAccountingBatches(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingBatchesAPI.GetFinanceAccountingBatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAccountingBatches`: []AccountingBatch
	fmt.Fprintf(os.Stdout, "Response from `AccountingBatchesAPI.GetFinanceAccountingBatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAccountingBatchesRequest struct via the builder pattern


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

[**[]AccountingBatch**](AccountingBatch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAccountingBatchesById

> AccountingBatch GetFinanceAccountingBatchesById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get AccountingBatch

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
	id := int32(56) // int32 | batcheId
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
	resp, r, err := apiClient.AccountingBatchesAPI.GetFinanceAccountingBatchesById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingBatchesAPI.GetFinanceAccountingBatchesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAccountingBatchesById`: AccountingBatch
	fmt.Fprintf(os.Stdout, "Response from `AccountingBatchesAPI.GetFinanceAccountingBatchesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | batcheId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAccountingBatchesByIdRequest struct via the builder pattern


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

[**AccountingBatch**](AccountingBatch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAccountingBatchesCount

> Count GetFinanceAccountingBatchesCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of AccountingBatch

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
	resp, r, err := apiClient.AccountingBatchesAPI.GetFinanceAccountingBatchesCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingBatchesAPI.GetFinanceAccountingBatchesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAccountingBatchesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `AccountingBatchesAPI.GetFinanceAccountingBatchesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAccountingBatchesCountRequest struct via the builder pattern


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


## PostFinanceAccountingBatches

> GLExport PostFinanceAccountingBatches(ctx).ClientId(clientId).CreateAccountingBatchRequest(createAccountingBatchRequest).Execute()

Post GLExport

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
	createAccountingBatchRequest := *openapiclient.NewCreateAccountingBatchRequest([]*int32{nil}) // CreateAccountingBatchRequest | accountingBatchParameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingBatchesAPI.PostFinanceAccountingBatches(context.Background()).ClientId(clientId).CreateAccountingBatchRequest(createAccountingBatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingBatchesAPI.PostFinanceAccountingBatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceAccountingBatches`: GLExport
	fmt.Fprintf(os.Stdout, "Response from `AccountingBatchesAPI.PostFinanceAccountingBatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceAccountingBatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **createAccountingBatchRequest** | [**CreateAccountingBatchRequest**](CreateAccountingBatchRequest.md) | accountingBatchParameters | 

### Return type

[**GLExport**](GLExport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceAccountingBatchesByIdExport

> GLExport PostFinanceAccountingBatchesByIdExport(ctx, id).ClientId(clientId).ExportAccountingBatchRequest(exportAccountingBatchRequest).Execute()

Post GLExport

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
	id := int32(56) // int32 | batcheId
	clientId := "clientId_example" // string | 
	exportAccountingBatchRequest := *openapiclient.NewExportAccountingBatchRequest() // ExportAccountingBatchRequest | batchExportParameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingBatchesAPI.PostFinanceAccountingBatchesByIdExport(context.Background(), id).ClientId(clientId).ExportAccountingBatchRequest(exportAccountingBatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingBatchesAPI.PostFinanceAccountingBatchesByIdExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceAccountingBatchesByIdExport`: GLExport
	fmt.Fprintf(os.Stdout, "Response from `AccountingBatchesAPI.PostFinanceAccountingBatchesByIdExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | batcheId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceAccountingBatchesByIdExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **exportAccountingBatchRequest** | [**ExportAccountingBatchRequest**](ExportAccountingBatchRequest.md) | batchExportParameters | 

### Return type

[**GLExport**](GLExport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceAccountingExport

> GLExport PostFinanceAccountingExport(ctx).ClientId(clientId).ExportAccountingBatchRequest(exportAccountingBatchRequest).Execute()

Post GLExport

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
	exportAccountingBatchRequest := *openapiclient.NewExportAccountingBatchRequest() // ExportAccountingBatchRequest | batchExportParameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingBatchesAPI.PostFinanceAccountingExport(context.Background()).ClientId(clientId).ExportAccountingBatchRequest(exportAccountingBatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingBatchesAPI.PostFinanceAccountingExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceAccountingExport`: GLExport
	fmt.Fprintf(os.Stdout, "Response from `AccountingBatchesAPI.PostFinanceAccountingExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceAccountingExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **exportAccountingBatchRequest** | [**ExportAccountingBatchRequest**](ExportAccountingBatchRequest.md) | batchExportParameters | 

### Return type

[**GLExport**](GLExport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

