# \AgreementBoardDefaultsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceAgreementsByParentIdBoardDefaultsById**](AgreementBoardDefaultsAPI.md#DeleteFinanceAgreementsByParentIdBoardDefaultsById) | **Delete** /finance/agreements/{parentId}/boardDefaults/{id} | Delete BoardDefault
[**GetFinanceAgreementsByParentIdBoardDefaults**](AgreementBoardDefaultsAPI.md#GetFinanceAgreementsByParentIdBoardDefaults) | **Get** /finance/agreements/{parentId}/boardDefaults | Get List of BoardDefault
[**GetFinanceAgreementsByParentIdBoardDefaultsById**](AgreementBoardDefaultsAPI.md#GetFinanceAgreementsByParentIdBoardDefaultsById) | **Get** /finance/agreements/{parentId}/boardDefaults/{id} | Get BoardDefault
[**GetFinanceAgreementsByParentIdBoardDefaultsCount**](AgreementBoardDefaultsAPI.md#GetFinanceAgreementsByParentIdBoardDefaultsCount) | **Get** /finance/agreements/{parentId}/boardDefaults/count | Get Count of BoardDefault
[**PatchFinanceAgreementsByParentIdBoardDefaultsById**](AgreementBoardDefaultsAPI.md#PatchFinanceAgreementsByParentIdBoardDefaultsById) | **Patch** /finance/agreements/{parentId}/boardDefaults/{id} | Patch BoardDefault
[**PostFinanceAgreementsByParentIdBoardDefaults**](AgreementBoardDefaultsAPI.md#PostFinanceAgreementsByParentIdBoardDefaults) | **Post** /finance/agreements/{parentId}/boardDefaults | Post BoardDefault
[**PutFinanceAgreementsByParentIdBoardDefaultsById**](AgreementBoardDefaultsAPI.md#PutFinanceAgreementsByParentIdBoardDefaultsById) | **Put** /finance/agreements/{parentId}/boardDefaults/{id} | Put BoardDefault



## DeleteFinanceAgreementsByParentIdBoardDefaultsById

> DeleteFinanceAgreementsByParentIdBoardDefaultsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete BoardDefault

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
	id := int32(56) // int32 | boardDefaultId
	parentId := int32(56) // int32 | agreementId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgreementBoardDefaultsAPI.DeleteFinanceAgreementsByParentIdBoardDefaultsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementBoardDefaultsAPI.DeleteFinanceAgreementsByParentIdBoardDefaultsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | boardDefaultId | 
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceAgreementsByParentIdBoardDefaultsByIdRequest struct via the builder pattern


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


## GetFinanceAgreementsByParentIdBoardDefaults

> []BoardDefault GetFinanceAgreementsByParentIdBoardDefaults(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of BoardDefault

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
	parentId := int32(56) // int32 | agreementId
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
	resp, r, err := apiClient.AgreementBoardDefaultsAPI.GetFinanceAgreementsByParentIdBoardDefaults(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementBoardDefaultsAPI.GetFinanceAgreementsByParentIdBoardDefaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementsByParentIdBoardDefaults`: []BoardDefault
	fmt.Fprintf(os.Stdout, "Response from `AgreementBoardDefaultsAPI.GetFinanceAgreementsByParentIdBoardDefaults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementsByParentIdBoardDefaultsRequest struct via the builder pattern


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

[**[]BoardDefault**](BoardDefault.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAgreementsByParentIdBoardDefaultsById

> BoardDefault GetFinanceAgreementsByParentIdBoardDefaultsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get BoardDefault

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
	id := int32(56) // int32 | boardDefaultId
	parentId := int32(56) // int32 | agreementId
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
	resp, r, err := apiClient.AgreementBoardDefaultsAPI.GetFinanceAgreementsByParentIdBoardDefaultsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementBoardDefaultsAPI.GetFinanceAgreementsByParentIdBoardDefaultsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementsByParentIdBoardDefaultsById`: BoardDefault
	fmt.Fprintf(os.Stdout, "Response from `AgreementBoardDefaultsAPI.GetFinanceAgreementsByParentIdBoardDefaultsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | boardDefaultId | 
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementsByParentIdBoardDefaultsByIdRequest struct via the builder pattern


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

[**BoardDefault**](BoardDefault.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAgreementsByParentIdBoardDefaultsCount

> Count GetFinanceAgreementsByParentIdBoardDefaultsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of BoardDefault

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
	parentId := int32(56) // int32 | agreementId
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
	resp, r, err := apiClient.AgreementBoardDefaultsAPI.GetFinanceAgreementsByParentIdBoardDefaultsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementBoardDefaultsAPI.GetFinanceAgreementsByParentIdBoardDefaultsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementsByParentIdBoardDefaultsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `AgreementBoardDefaultsAPI.GetFinanceAgreementsByParentIdBoardDefaultsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementsByParentIdBoardDefaultsCountRequest struct via the builder pattern


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


## PatchFinanceAgreementsByParentIdBoardDefaultsById

> BoardDefault PatchFinanceAgreementsByParentIdBoardDefaultsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch BoardDefault

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
	id := int32(56) // int32 | boardDefaultId
	parentId := int32(56) // int32 | agreementId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgreementBoardDefaultsAPI.PatchFinanceAgreementsByParentIdBoardDefaultsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementBoardDefaultsAPI.PatchFinanceAgreementsByParentIdBoardDefaultsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceAgreementsByParentIdBoardDefaultsById`: BoardDefault
	fmt.Fprintf(os.Stdout, "Response from `AgreementBoardDefaultsAPI.PatchFinanceAgreementsByParentIdBoardDefaultsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | boardDefaultId | 
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceAgreementsByParentIdBoardDefaultsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**BoardDefault**](BoardDefault.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceAgreementsByParentIdBoardDefaults

> BoardDefault PostFinanceAgreementsByParentIdBoardDefaults(ctx, parentId).ClientId(clientId).BoardDefault(boardDefault).Execute()

Post BoardDefault

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
	parentId := int32(56) // int32 | agreementId
	clientId := "clientId_example" // string | 
	boardDefault := *openapiclient.NewBoardDefault(*openapiclient.NewBoardReference()) // BoardDefault | boardDefault

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgreementBoardDefaultsAPI.PostFinanceAgreementsByParentIdBoardDefaults(context.Background(), parentId).ClientId(clientId).BoardDefault(boardDefault).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementBoardDefaultsAPI.PostFinanceAgreementsByParentIdBoardDefaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceAgreementsByParentIdBoardDefaults`: BoardDefault
	fmt.Fprintf(os.Stdout, "Response from `AgreementBoardDefaultsAPI.PostFinanceAgreementsByParentIdBoardDefaults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceAgreementsByParentIdBoardDefaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **boardDefault** | [**BoardDefault**](BoardDefault.md) | boardDefault | 

### Return type

[**BoardDefault**](BoardDefault.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceAgreementsByParentIdBoardDefaultsById

> BoardDefault PutFinanceAgreementsByParentIdBoardDefaultsById(ctx, id, parentId).ClientId(clientId).BoardDefault(boardDefault).Execute()

Put BoardDefault

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
	id := int32(56) // int32 | boardDefaultId
	parentId := int32(56) // int32 | agreementId
	clientId := "clientId_example" // string | 
	boardDefault := *openapiclient.NewBoardDefault(*openapiclient.NewBoardReference()) // BoardDefault | boardDefault

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgreementBoardDefaultsAPI.PutFinanceAgreementsByParentIdBoardDefaultsById(context.Background(), id, parentId).ClientId(clientId).BoardDefault(boardDefault).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementBoardDefaultsAPI.PutFinanceAgreementsByParentIdBoardDefaultsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceAgreementsByParentIdBoardDefaultsById`: BoardDefault
	fmt.Fprintf(os.Stdout, "Response from `AgreementBoardDefaultsAPI.PutFinanceAgreementsByParentIdBoardDefaultsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | boardDefaultId | 
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceAgreementsByParentIdBoardDefaultsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **boardDefault** | [**BoardDefault**](BoardDefault.md) | boardDefault | 

### Return type

[**BoardDefault**](BoardDefault.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

