# \AgreementTypeBoardDefaultsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceAgreementTypesByParentIdBoardDefaultsById**](AgreementTypeBoardDefaultsAPI.md#DeleteFinanceAgreementTypesByParentIdBoardDefaultsById) | **Delete** /finance/agreementTypes/{parentId}/boardDefaults/{id} | Delete AgreementTypeBoardDefault
[**GetFinanceAgreementTypesByParentIdBoardDefaults**](AgreementTypeBoardDefaultsAPI.md#GetFinanceAgreementTypesByParentIdBoardDefaults) | **Get** /finance/agreementTypes/{parentId}/boardDefaults | Get List of AgreementTypeBoardDefault
[**GetFinanceAgreementTypesByParentIdBoardDefaultsById**](AgreementTypeBoardDefaultsAPI.md#GetFinanceAgreementTypesByParentIdBoardDefaultsById) | **Get** /finance/agreementTypes/{parentId}/boardDefaults/{id} | Get AgreementTypeBoardDefault
[**GetFinanceAgreementTypesByParentIdBoardDefaultsCount**](AgreementTypeBoardDefaultsAPI.md#GetFinanceAgreementTypesByParentIdBoardDefaultsCount) | **Get** /finance/agreementTypes/{parentId}/boardDefaults/count | Get Count of AgreementTypeBoardDefault
[**PatchFinanceAgreementTypesByParentIdBoardDefaultsById**](AgreementTypeBoardDefaultsAPI.md#PatchFinanceAgreementTypesByParentIdBoardDefaultsById) | **Patch** /finance/agreementTypes/{parentId}/boardDefaults/{id} | Patch AgreementTypeBoardDefault
[**PostFinanceAgreementTypesByParentIdBoardDefaults**](AgreementTypeBoardDefaultsAPI.md#PostFinanceAgreementTypesByParentIdBoardDefaults) | **Post** /finance/agreementTypes/{parentId}/boardDefaults | Post AgreementTypeBoardDefault
[**PutFinanceAgreementTypesByParentIdBoardDefaultsById**](AgreementTypeBoardDefaultsAPI.md#PutFinanceAgreementTypesByParentIdBoardDefaultsById) | **Put** /finance/agreementTypes/{parentId}/boardDefaults/{id} | Put AgreementTypeBoardDefault



## DeleteFinanceAgreementTypesByParentIdBoardDefaultsById

> DeleteFinanceAgreementTypesByParentIdBoardDefaultsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete AgreementTypeBoardDefault

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
	parentId := int32(56) // int32 | agreementTypeId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgreementTypeBoardDefaultsAPI.DeleteFinanceAgreementTypesByParentIdBoardDefaultsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeBoardDefaultsAPI.DeleteFinanceAgreementTypesByParentIdBoardDefaultsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | boardDefaultId | 
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceAgreementTypesByParentIdBoardDefaultsByIdRequest struct via the builder pattern


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


## GetFinanceAgreementTypesByParentIdBoardDefaults

> []AgreementTypeBoardDefault GetFinanceAgreementTypesByParentIdBoardDefaults(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of AgreementTypeBoardDefault

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
	parentId := int32(56) // int32 | agreementTypeId
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
	resp, r, err := apiClient.AgreementTypeBoardDefaultsAPI.GetFinanceAgreementTypesByParentIdBoardDefaults(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeBoardDefaultsAPI.GetFinanceAgreementTypesByParentIdBoardDefaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementTypesByParentIdBoardDefaults`: []AgreementTypeBoardDefault
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeBoardDefaultsAPI.GetFinanceAgreementTypesByParentIdBoardDefaults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementTypesByParentIdBoardDefaultsRequest struct via the builder pattern


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

[**[]AgreementTypeBoardDefault**](AgreementTypeBoardDefault.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAgreementTypesByParentIdBoardDefaultsById

> AgreementTypeBoardDefault GetFinanceAgreementTypesByParentIdBoardDefaultsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get AgreementTypeBoardDefault

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
	parentId := int32(56) // int32 | agreementTypeId
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
	resp, r, err := apiClient.AgreementTypeBoardDefaultsAPI.GetFinanceAgreementTypesByParentIdBoardDefaultsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeBoardDefaultsAPI.GetFinanceAgreementTypesByParentIdBoardDefaultsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementTypesByParentIdBoardDefaultsById`: AgreementTypeBoardDefault
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeBoardDefaultsAPI.GetFinanceAgreementTypesByParentIdBoardDefaultsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | boardDefaultId | 
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementTypesByParentIdBoardDefaultsByIdRequest struct via the builder pattern


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

[**AgreementTypeBoardDefault**](AgreementTypeBoardDefault.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAgreementTypesByParentIdBoardDefaultsCount

> Count GetFinanceAgreementTypesByParentIdBoardDefaultsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of AgreementTypeBoardDefault

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
	parentId := int32(56) // int32 | agreementTypeId
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
	resp, r, err := apiClient.AgreementTypeBoardDefaultsAPI.GetFinanceAgreementTypesByParentIdBoardDefaultsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeBoardDefaultsAPI.GetFinanceAgreementTypesByParentIdBoardDefaultsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementTypesByParentIdBoardDefaultsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeBoardDefaultsAPI.GetFinanceAgreementTypesByParentIdBoardDefaultsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementTypesByParentIdBoardDefaultsCountRequest struct via the builder pattern


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


## PatchFinanceAgreementTypesByParentIdBoardDefaultsById

> AgreementTypeBoardDefault PatchFinanceAgreementTypesByParentIdBoardDefaultsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch AgreementTypeBoardDefault

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
	parentId := int32(56) // int32 | agreementTypeId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgreementTypeBoardDefaultsAPI.PatchFinanceAgreementTypesByParentIdBoardDefaultsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeBoardDefaultsAPI.PatchFinanceAgreementTypesByParentIdBoardDefaultsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceAgreementTypesByParentIdBoardDefaultsById`: AgreementTypeBoardDefault
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeBoardDefaultsAPI.PatchFinanceAgreementTypesByParentIdBoardDefaultsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | boardDefaultId | 
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceAgreementTypesByParentIdBoardDefaultsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**AgreementTypeBoardDefault**](AgreementTypeBoardDefault.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceAgreementTypesByParentIdBoardDefaults

> AgreementTypeBoardDefault PostFinanceAgreementTypesByParentIdBoardDefaults(ctx, parentId).ClientId(clientId).AgreementTypeBoardDefault(agreementTypeBoardDefault).Execute()

Post AgreementTypeBoardDefault

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
	parentId := int32(56) // int32 | agreementTypeId
	clientId := "clientId_example" // string | 
	agreementTypeBoardDefault := *openapiclient.NewAgreementTypeBoardDefault(*openapiclient.NewSystemLocationReference()) // AgreementTypeBoardDefault | boardDefault

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgreementTypeBoardDefaultsAPI.PostFinanceAgreementTypesByParentIdBoardDefaults(context.Background(), parentId).ClientId(clientId).AgreementTypeBoardDefault(agreementTypeBoardDefault).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeBoardDefaultsAPI.PostFinanceAgreementTypesByParentIdBoardDefaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceAgreementTypesByParentIdBoardDefaults`: AgreementTypeBoardDefault
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeBoardDefaultsAPI.PostFinanceAgreementTypesByParentIdBoardDefaults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceAgreementTypesByParentIdBoardDefaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **agreementTypeBoardDefault** | [**AgreementTypeBoardDefault**](AgreementTypeBoardDefault.md) | boardDefault | 

### Return type

[**AgreementTypeBoardDefault**](AgreementTypeBoardDefault.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceAgreementTypesByParentIdBoardDefaultsById

> AgreementTypeBoardDefault PutFinanceAgreementTypesByParentIdBoardDefaultsById(ctx, id, parentId).ClientId(clientId).AgreementTypeBoardDefault(agreementTypeBoardDefault).Execute()

Put AgreementTypeBoardDefault

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
	parentId := int32(56) // int32 | agreementTypeId
	clientId := "clientId_example" // string | 
	agreementTypeBoardDefault := *openapiclient.NewAgreementTypeBoardDefault(*openapiclient.NewSystemLocationReference()) // AgreementTypeBoardDefault | boardDefault

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgreementTypeBoardDefaultsAPI.PutFinanceAgreementTypesByParentIdBoardDefaultsById(context.Background(), id, parentId).ClientId(clientId).AgreementTypeBoardDefault(agreementTypeBoardDefault).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeBoardDefaultsAPI.PutFinanceAgreementTypesByParentIdBoardDefaultsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceAgreementTypesByParentIdBoardDefaultsById`: AgreementTypeBoardDefault
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeBoardDefaultsAPI.PutFinanceAgreementTypesByParentIdBoardDefaultsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | boardDefaultId | 
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceAgreementTypesByParentIdBoardDefaultsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **agreementTypeBoardDefault** | [**AgreementTypeBoardDefault**](AgreementTypeBoardDefault.md) | boardDefault | 

### Return type

[**AgreementTypeBoardDefault**](AgreementTypeBoardDefault.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

