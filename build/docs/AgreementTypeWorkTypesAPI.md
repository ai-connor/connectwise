# \AgreementTypeWorkTypesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceAgreementTypesByParentIdWorktypesById**](AgreementTypeWorkTypesAPI.md#DeleteFinanceAgreementTypesByParentIdWorktypesById) | **Delete** /finance/agreementTypes/{parentId}/worktypes/{id} | Delete AgreementTypeWorkType
[**GetFinanceAgreementTypesByParentIdWorktypes**](AgreementTypeWorkTypesAPI.md#GetFinanceAgreementTypesByParentIdWorktypes) | **Get** /finance/agreementTypes/{parentId}/worktypes | Get List of AgreementTypeWorkType
[**GetFinanceAgreementTypesByParentIdWorktypesById**](AgreementTypeWorkTypesAPI.md#GetFinanceAgreementTypesByParentIdWorktypesById) | **Get** /finance/agreementTypes/{parentId}/worktypes/{id} | Get AgreementTypeWorkType
[**GetFinanceAgreementTypesByParentIdWorktypesCount**](AgreementTypeWorkTypesAPI.md#GetFinanceAgreementTypesByParentIdWorktypesCount) | **Get** /finance/agreementTypes/{parentId}/worktypes/count | Get Count of AgreementTypeWorkType
[**PatchFinanceAgreementTypesByParentIdWorktypesById**](AgreementTypeWorkTypesAPI.md#PatchFinanceAgreementTypesByParentIdWorktypesById) | **Patch** /finance/agreementTypes/{parentId}/worktypes/{id} | Patch AgreementTypeWorkType
[**PostFinanceAgreementTypesByParentIdWorktypes**](AgreementTypeWorkTypesAPI.md#PostFinanceAgreementTypesByParentIdWorktypes) | **Post** /finance/agreementTypes/{parentId}/worktypes | Post AgreementTypeWorkType
[**PutFinanceAgreementTypesByParentIdWorktypesById**](AgreementTypeWorkTypesAPI.md#PutFinanceAgreementTypesByParentIdWorktypesById) | **Put** /finance/agreementTypes/{parentId}/worktypes/{id} | Put AgreementTypeWorkType



## DeleteFinanceAgreementTypesByParentIdWorktypesById

> DeleteFinanceAgreementTypesByParentIdWorktypesById(ctx, id, parentId).ClientId(clientId).Execute()

Delete AgreementTypeWorkType

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
	id := int32(56) // int32 | worktypeId
	parentId := int32(56) // int32 | agreementTypeId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgreementTypeWorkTypesAPI.DeleteFinanceAgreementTypesByParentIdWorktypesById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkTypesAPI.DeleteFinanceAgreementTypesByParentIdWorktypesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | worktypeId | 
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceAgreementTypesByParentIdWorktypesByIdRequest struct via the builder pattern


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


## GetFinanceAgreementTypesByParentIdWorktypes

> []AgreementTypeWorkType GetFinanceAgreementTypesByParentIdWorktypes(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of AgreementTypeWorkType

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
	resp, r, err := apiClient.AgreementTypeWorkTypesAPI.GetFinanceAgreementTypesByParentIdWorktypes(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkTypesAPI.GetFinanceAgreementTypesByParentIdWorktypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementTypesByParentIdWorktypes`: []AgreementTypeWorkType
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeWorkTypesAPI.GetFinanceAgreementTypesByParentIdWorktypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementTypesByParentIdWorktypesRequest struct via the builder pattern


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

[**[]AgreementTypeWorkType**](AgreementTypeWorkType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAgreementTypesByParentIdWorktypesById

> AgreementTypeWorkType GetFinanceAgreementTypesByParentIdWorktypesById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get AgreementTypeWorkType

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
	id := int32(56) // int32 | worktypeId
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
	resp, r, err := apiClient.AgreementTypeWorkTypesAPI.GetFinanceAgreementTypesByParentIdWorktypesById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkTypesAPI.GetFinanceAgreementTypesByParentIdWorktypesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementTypesByParentIdWorktypesById`: AgreementTypeWorkType
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeWorkTypesAPI.GetFinanceAgreementTypesByParentIdWorktypesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | worktypeId | 
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementTypesByParentIdWorktypesByIdRequest struct via the builder pattern


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

[**AgreementTypeWorkType**](AgreementTypeWorkType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAgreementTypesByParentIdWorktypesCount

> Count GetFinanceAgreementTypesByParentIdWorktypesCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of AgreementTypeWorkType

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
	resp, r, err := apiClient.AgreementTypeWorkTypesAPI.GetFinanceAgreementTypesByParentIdWorktypesCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkTypesAPI.GetFinanceAgreementTypesByParentIdWorktypesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementTypesByParentIdWorktypesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeWorkTypesAPI.GetFinanceAgreementTypesByParentIdWorktypesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementTypesByParentIdWorktypesCountRequest struct via the builder pattern


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


## PatchFinanceAgreementTypesByParentIdWorktypesById

> AgreementTypeWorkType PatchFinanceAgreementTypesByParentIdWorktypesById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch AgreementTypeWorkType

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
	id := int32(56) // int32 | worktypeId
	parentId := int32(56) // int32 | agreementTypeId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgreementTypeWorkTypesAPI.PatchFinanceAgreementTypesByParentIdWorktypesById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkTypesAPI.PatchFinanceAgreementTypesByParentIdWorktypesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceAgreementTypesByParentIdWorktypesById`: AgreementTypeWorkType
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeWorkTypesAPI.PatchFinanceAgreementTypesByParentIdWorktypesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | worktypeId | 
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceAgreementTypesByParentIdWorktypesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**AgreementTypeWorkType**](AgreementTypeWorkType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceAgreementTypesByParentIdWorktypes

> AgreementTypeWorkType PostFinanceAgreementTypesByParentIdWorktypes(ctx, parentId).ClientId(clientId).AgreementTypeWorkType(agreementTypeWorkType).Execute()

Post AgreementTypeWorkType

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
	agreementTypeWorkType := *openapiclient.NewAgreementTypeWorkType("RateType_example", "BillTime_example", "OverageRateType_example") // AgreementTypeWorkType | workType

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgreementTypeWorkTypesAPI.PostFinanceAgreementTypesByParentIdWorktypes(context.Background(), parentId).ClientId(clientId).AgreementTypeWorkType(agreementTypeWorkType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkTypesAPI.PostFinanceAgreementTypesByParentIdWorktypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceAgreementTypesByParentIdWorktypes`: AgreementTypeWorkType
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeWorkTypesAPI.PostFinanceAgreementTypesByParentIdWorktypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceAgreementTypesByParentIdWorktypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **agreementTypeWorkType** | [**AgreementTypeWorkType**](AgreementTypeWorkType.md) | workType | 

### Return type

[**AgreementTypeWorkType**](AgreementTypeWorkType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceAgreementTypesByParentIdWorktypesById

> AgreementTypeWorkType PutFinanceAgreementTypesByParentIdWorktypesById(ctx, id, parentId).ClientId(clientId).AgreementTypeWorkType(agreementTypeWorkType).Execute()

Put AgreementTypeWorkType

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
	id := int32(56) // int32 | worktypeId
	parentId := int32(56) // int32 | agreementTypeId
	clientId := "clientId_example" // string | 
	agreementTypeWorkType := *openapiclient.NewAgreementTypeWorkType("RateType_example", "BillTime_example", "OverageRateType_example") // AgreementTypeWorkType | workType

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgreementTypeWorkTypesAPI.PutFinanceAgreementTypesByParentIdWorktypesById(context.Background(), id, parentId).ClientId(clientId).AgreementTypeWorkType(agreementTypeWorkType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkTypesAPI.PutFinanceAgreementTypesByParentIdWorktypesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceAgreementTypesByParentIdWorktypesById`: AgreementTypeWorkType
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeWorkTypesAPI.PutFinanceAgreementTypesByParentIdWorktypesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | worktypeId | 
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceAgreementTypesByParentIdWorktypesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **agreementTypeWorkType** | [**AgreementTypeWorkType**](AgreementTypeWorkType.md) | workType | 

### Return type

[**AgreementTypeWorkType**](AgreementTypeWorkType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

