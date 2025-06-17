# \AgreementAdditionsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceAgreementsByParentIdAdditionsById**](AgreementAdditionsAPI.md#DeleteFinanceAgreementsByParentIdAdditionsById) | **Delete** /finance/agreements/{parentId}/additions/{id} | Delete Addition
[**GetFinanceAgreementsByParentIdAdditions**](AgreementAdditionsAPI.md#GetFinanceAgreementsByParentIdAdditions) | **Get** /finance/agreements/{parentId}/additions | Get List of Addition
[**GetFinanceAgreementsByParentIdAdditionsById**](AgreementAdditionsAPI.md#GetFinanceAgreementsByParentIdAdditionsById) | **Get** /finance/agreements/{parentId}/additions/{id} | Get Addition
[**GetFinanceAgreementsByParentIdAdditionsCount**](AgreementAdditionsAPI.md#GetFinanceAgreementsByParentIdAdditionsCount) | **Get** /finance/agreements/{parentId}/additions/count | Get Count of Addition
[**PatchFinanceAgreementsByParentIdAdditionsById**](AgreementAdditionsAPI.md#PatchFinanceAgreementsByParentIdAdditionsById) | **Patch** /finance/agreements/{parentId}/additions/{id} | Patch Addition
[**PostFinanceAgreementsByParentIdAdditions**](AgreementAdditionsAPI.md#PostFinanceAgreementsByParentIdAdditions) | **Post** /finance/agreements/{parentId}/additions | Post Addition
[**PutFinanceAgreementsByParentIdAdditionsById**](AgreementAdditionsAPI.md#PutFinanceAgreementsByParentIdAdditionsById) | **Put** /finance/agreements/{parentId}/additions/{id} | Put Addition



## DeleteFinanceAgreementsByParentIdAdditionsById

> DeleteFinanceAgreementsByParentIdAdditionsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete Addition

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
	id := int32(56) // int32 | additionId
	parentId := int32(56) // int32 | agreementId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgreementAdditionsAPI.DeleteFinanceAgreementsByParentIdAdditionsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementAdditionsAPI.DeleteFinanceAgreementsByParentIdAdditionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | additionId | 
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceAgreementsByParentIdAdditionsByIdRequest struct via the builder pattern


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


## GetFinanceAgreementsByParentIdAdditions

> []Addition GetFinanceAgreementsByParentIdAdditions(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of Addition

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
	resp, r, err := apiClient.AgreementAdditionsAPI.GetFinanceAgreementsByParentIdAdditions(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementAdditionsAPI.GetFinanceAgreementsByParentIdAdditions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementsByParentIdAdditions`: []Addition
	fmt.Fprintf(os.Stdout, "Response from `AgreementAdditionsAPI.GetFinanceAgreementsByParentIdAdditions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementsByParentIdAdditionsRequest struct via the builder pattern


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

[**[]Addition**](Addition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAgreementsByParentIdAdditionsById

> Addition GetFinanceAgreementsByParentIdAdditionsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Addition

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
	id := int32(56) // int32 | additionId
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
	resp, r, err := apiClient.AgreementAdditionsAPI.GetFinanceAgreementsByParentIdAdditionsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementAdditionsAPI.GetFinanceAgreementsByParentIdAdditionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementsByParentIdAdditionsById`: Addition
	fmt.Fprintf(os.Stdout, "Response from `AgreementAdditionsAPI.GetFinanceAgreementsByParentIdAdditionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | additionId | 
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementsByParentIdAdditionsByIdRequest struct via the builder pattern


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

[**Addition**](Addition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAgreementsByParentIdAdditionsCount

> Count GetFinanceAgreementsByParentIdAdditionsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of Addition

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
	resp, r, err := apiClient.AgreementAdditionsAPI.GetFinanceAgreementsByParentIdAdditionsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementAdditionsAPI.GetFinanceAgreementsByParentIdAdditionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementsByParentIdAdditionsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `AgreementAdditionsAPI.GetFinanceAgreementsByParentIdAdditionsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementsByParentIdAdditionsCountRequest struct via the builder pattern


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


## PatchFinanceAgreementsByParentIdAdditionsById

> Addition PatchFinanceAgreementsByParentIdAdditionsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch Addition

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
	id := int32(56) // int32 | additionId
	parentId := int32(56) // int32 | agreementId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgreementAdditionsAPI.PatchFinanceAgreementsByParentIdAdditionsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementAdditionsAPI.PatchFinanceAgreementsByParentIdAdditionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceAgreementsByParentIdAdditionsById`: Addition
	fmt.Fprintf(os.Stdout, "Response from `AgreementAdditionsAPI.PatchFinanceAgreementsByParentIdAdditionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | additionId | 
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceAgreementsByParentIdAdditionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**Addition**](Addition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceAgreementsByParentIdAdditions

> Addition PostFinanceAgreementsByParentIdAdditions(ctx, parentId).ClientId(clientId).Addition(addition).Execute()

Post Addition

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
	addition := *openapiclient.NewAddition(*openapiclient.NewIvItemReference(), "BillCustomer_example") // Addition | addition

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgreementAdditionsAPI.PostFinanceAgreementsByParentIdAdditions(context.Background(), parentId).ClientId(clientId).Addition(addition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementAdditionsAPI.PostFinanceAgreementsByParentIdAdditions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceAgreementsByParentIdAdditions`: Addition
	fmt.Fprintf(os.Stdout, "Response from `AgreementAdditionsAPI.PostFinanceAgreementsByParentIdAdditions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceAgreementsByParentIdAdditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **addition** | [**Addition**](Addition.md) | addition | 

### Return type

[**Addition**](Addition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceAgreementsByParentIdAdditionsById

> Addition PutFinanceAgreementsByParentIdAdditionsById(ctx, id, parentId).ClientId(clientId).Addition(addition).Execute()

Put Addition

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
	id := int32(56) // int32 | additionId
	parentId := int32(56) // int32 | agreementId
	clientId := "clientId_example" // string | 
	addition := *openapiclient.NewAddition(*openapiclient.NewIvItemReference(), "BillCustomer_example") // Addition | addition

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgreementAdditionsAPI.PutFinanceAgreementsByParentIdAdditionsById(context.Background(), id, parentId).ClientId(clientId).Addition(addition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementAdditionsAPI.PutFinanceAgreementsByParentIdAdditionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceAgreementsByParentIdAdditionsById`: Addition
	fmt.Fprintf(os.Stdout, "Response from `AgreementAdditionsAPI.PutFinanceAgreementsByParentIdAdditionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | additionId | 
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceAgreementsByParentIdAdditionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **addition** | [**Addition**](Addition.md) | addition | 

### Return type

[**Addition**](Addition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

