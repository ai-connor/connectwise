# \AccountingPackageSetupsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFinanceAccountingPackageSetup**](AccountingPackageSetupsAPI.md#GetFinanceAccountingPackageSetup) | **Get** /finance/accountingPackageSetup | Get List of AccountingPackageSetup
[**GetFinanceAccountingPackageSetupById**](AccountingPackageSetupsAPI.md#GetFinanceAccountingPackageSetupById) | **Get** /finance/accountingPackageSetup/{id} | Get AccountingPackageSetup
[**GetFinanceAccountingPackageSetupCount**](AccountingPackageSetupsAPI.md#GetFinanceAccountingPackageSetupCount) | **Get** /finance/accountingPackageSetup/count | Get Count of AccountingPackageSetup
[**PatchFinanceAccountingPackageSetupById**](AccountingPackageSetupsAPI.md#PatchFinanceAccountingPackageSetupById) | **Patch** /finance/accountingPackageSetup/{id} | Patch AccountingPackageSetup
[**PutFinanceAccountingPackageSetupById**](AccountingPackageSetupsAPI.md#PutFinanceAccountingPackageSetupById) | **Put** /finance/accountingPackageSetup/{id} | Put AccountingPackageSetup



## GetFinanceAccountingPackageSetup

> []AccountingPackageSetup GetFinanceAccountingPackageSetup(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of AccountingPackageSetup

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
	resp, r, err := apiClient.AccountingPackageSetupsAPI.GetFinanceAccountingPackageSetup(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingPackageSetupsAPI.GetFinanceAccountingPackageSetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAccountingPackageSetup`: []AccountingPackageSetup
	fmt.Fprintf(os.Stdout, "Response from `AccountingPackageSetupsAPI.GetFinanceAccountingPackageSetup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAccountingPackageSetupRequest struct via the builder pattern


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

[**[]AccountingPackageSetup**](AccountingPackageSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAccountingPackageSetupById

> AccountingPackageSetup GetFinanceAccountingPackageSetupById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get AccountingPackageSetup

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
	id := int32(56) // int32 | accountingPackageSetupId
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
	resp, r, err := apiClient.AccountingPackageSetupsAPI.GetFinanceAccountingPackageSetupById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingPackageSetupsAPI.GetFinanceAccountingPackageSetupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAccountingPackageSetupById`: AccountingPackageSetup
	fmt.Fprintf(os.Stdout, "Response from `AccountingPackageSetupsAPI.GetFinanceAccountingPackageSetupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | accountingPackageSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAccountingPackageSetupByIdRequest struct via the builder pattern


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

[**AccountingPackageSetup**](AccountingPackageSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAccountingPackageSetupCount

> Count GetFinanceAccountingPackageSetupCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of AccountingPackageSetup

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
	resp, r, err := apiClient.AccountingPackageSetupsAPI.GetFinanceAccountingPackageSetupCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingPackageSetupsAPI.GetFinanceAccountingPackageSetupCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAccountingPackageSetupCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `AccountingPackageSetupsAPI.GetFinanceAccountingPackageSetupCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAccountingPackageSetupCountRequest struct via the builder pattern


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


## PatchFinanceAccountingPackageSetupById

> AccountingPackageSetup PatchFinanceAccountingPackageSetupById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch AccountingPackageSetup

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
	id := int32(56) // int32 | accountingPackageSetupId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingPackageSetupsAPI.PatchFinanceAccountingPackageSetupById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingPackageSetupsAPI.PatchFinanceAccountingPackageSetupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceAccountingPackageSetupById`: AccountingPackageSetup
	fmt.Fprintf(os.Stdout, "Response from `AccountingPackageSetupsAPI.PatchFinanceAccountingPackageSetupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | accountingPackageSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceAccountingPackageSetupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**AccountingPackageSetup**](AccountingPackageSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceAccountingPackageSetupById

> AccountingPackageSetup PutFinanceAccountingPackageSetupById(ctx, id).ClientId(clientId).AccountingPackageSetup(accountingPackageSetup).Execute()

Put AccountingPackageSetup

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
	id := int32(56) // int32 | accountingPackageSetupId
	clientId := "clientId_example" // string | 
	accountingPackageSetup := *openapiclient.NewAccountingPackageSetup(*openapiclient.NewAccountingPackageReference()) // AccountingPackageSetup | accountingPackageSetup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingPackageSetupsAPI.PutFinanceAccountingPackageSetupById(context.Background(), id).ClientId(clientId).AccountingPackageSetup(accountingPackageSetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingPackageSetupsAPI.PutFinanceAccountingPackageSetupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceAccountingPackageSetupById`: AccountingPackageSetup
	fmt.Fprintf(os.Stdout, "Response from `AccountingPackageSetupsAPI.PutFinanceAccountingPackageSetupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | accountingPackageSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceAccountingPackageSetupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **accountingPackageSetup** | [**AccountingPackageSetup**](AccountingPackageSetup.md) | accountingPackageSetup | 

### Return type

[**AccountingPackageSetup**](AccountingPackageSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

