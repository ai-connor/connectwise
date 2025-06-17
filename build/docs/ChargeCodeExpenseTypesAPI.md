# \ChargeCodeExpenseTypesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTimeChargeCodesByParentIdExpenseTypesById**](ChargeCodeExpenseTypesAPI.md#DeleteTimeChargeCodesByParentIdExpenseTypesById) | **Delete** /time/chargeCodes/{parentId}/expenseTypes/{id} | Delete ChargeCodeExpenseType
[**GetTimeChargeCodesByParentIdExpenseTypes**](ChargeCodeExpenseTypesAPI.md#GetTimeChargeCodesByParentIdExpenseTypes) | **Get** /time/chargeCodes/{parentId}/expenseTypes | Get List of ChargeCodeExpenseType
[**GetTimeChargeCodesByParentIdExpenseTypesById**](ChargeCodeExpenseTypesAPI.md#GetTimeChargeCodesByParentIdExpenseTypesById) | **Get** /time/chargeCodes/{parentId}/expenseTypes/{id} | Get ChargeCodeExpenseType
[**GetTimeChargeCodesByParentIdExpenseTypesCount**](ChargeCodeExpenseTypesAPI.md#GetTimeChargeCodesByParentIdExpenseTypesCount) | **Get** /time/chargeCodes/{parentId}/expenseTypes/count | Get Count of ChargeCodeExpenseType
[**PatchTimeChargeCodesByParentIdExpenseTypesById**](ChargeCodeExpenseTypesAPI.md#PatchTimeChargeCodesByParentIdExpenseTypesById) | **Patch** /time/chargeCodes/{parentId}/expenseTypes/{id} | Patch ChargeCodeExpenseType
[**PostTimeChargeCodesByParentIdExpenseTypes**](ChargeCodeExpenseTypesAPI.md#PostTimeChargeCodesByParentIdExpenseTypes) | **Post** /time/chargeCodes/{parentId}/expenseTypes | Post ChargeCodeExpenseType
[**PutTimeChargeCodesByParentIdExpenseTypesById**](ChargeCodeExpenseTypesAPI.md#PutTimeChargeCodesByParentIdExpenseTypesById) | **Put** /time/chargeCodes/{parentId}/expenseTypes/{id} | Put ChargeCodeExpenseType



## DeleteTimeChargeCodesByParentIdExpenseTypesById

> DeleteTimeChargeCodesByParentIdExpenseTypesById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ChargeCodeExpenseType

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
	id := int32(56) // int32 | expenseTypeId
	parentId := int32(56) // int32 | chargeCodeId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChargeCodeExpenseTypesAPI.DeleteTimeChargeCodesByParentIdExpenseTypesById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChargeCodeExpenseTypesAPI.DeleteTimeChargeCodesByParentIdExpenseTypesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | expenseTypeId | 
**parentId** | **int32** | chargeCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTimeChargeCodesByParentIdExpenseTypesByIdRequest struct via the builder pattern


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


## GetTimeChargeCodesByParentIdExpenseTypes

> []ChargeCodeExpenseType GetTimeChargeCodesByParentIdExpenseTypes(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ChargeCodeExpenseType

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
	parentId := int32(56) // int32 | chargeCodeId
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
	resp, r, err := apiClient.ChargeCodeExpenseTypesAPI.GetTimeChargeCodesByParentIdExpenseTypes(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChargeCodeExpenseTypesAPI.GetTimeChargeCodesByParentIdExpenseTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeChargeCodesByParentIdExpenseTypes`: []ChargeCodeExpenseType
	fmt.Fprintf(os.Stdout, "Response from `ChargeCodeExpenseTypesAPI.GetTimeChargeCodesByParentIdExpenseTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | chargeCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeChargeCodesByParentIdExpenseTypesRequest struct via the builder pattern


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

[**[]ChargeCodeExpenseType**](ChargeCodeExpenseType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeChargeCodesByParentIdExpenseTypesById

> ChargeCodeExpenseType GetTimeChargeCodesByParentIdExpenseTypesById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ChargeCodeExpenseType

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
	id := int32(56) // int32 | expenseTypeId
	parentId := int32(56) // int32 | chargeCodeId
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
	resp, r, err := apiClient.ChargeCodeExpenseTypesAPI.GetTimeChargeCodesByParentIdExpenseTypesById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChargeCodeExpenseTypesAPI.GetTimeChargeCodesByParentIdExpenseTypesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeChargeCodesByParentIdExpenseTypesById`: ChargeCodeExpenseType
	fmt.Fprintf(os.Stdout, "Response from `ChargeCodeExpenseTypesAPI.GetTimeChargeCodesByParentIdExpenseTypesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | expenseTypeId | 
**parentId** | **int32** | chargeCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeChargeCodesByParentIdExpenseTypesByIdRequest struct via the builder pattern


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

[**ChargeCodeExpenseType**](ChargeCodeExpenseType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeChargeCodesByParentIdExpenseTypesCount

> Count GetTimeChargeCodesByParentIdExpenseTypesCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ChargeCodeExpenseType

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
	parentId := int32(56) // int32 | chargeCodeId
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
	resp, r, err := apiClient.ChargeCodeExpenseTypesAPI.GetTimeChargeCodesByParentIdExpenseTypesCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChargeCodeExpenseTypesAPI.GetTimeChargeCodesByParentIdExpenseTypesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeChargeCodesByParentIdExpenseTypesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ChargeCodeExpenseTypesAPI.GetTimeChargeCodesByParentIdExpenseTypesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | chargeCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeChargeCodesByParentIdExpenseTypesCountRequest struct via the builder pattern


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


## PatchTimeChargeCodesByParentIdExpenseTypesById

> ChargeCodeExpenseType PatchTimeChargeCodesByParentIdExpenseTypesById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ChargeCodeExpenseType

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
	id := int32(56) // int32 | expenseTypeId
	parentId := int32(56) // int32 | chargeCodeId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChargeCodeExpenseTypesAPI.PatchTimeChargeCodesByParentIdExpenseTypesById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChargeCodeExpenseTypesAPI.PatchTimeChargeCodesByParentIdExpenseTypesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTimeChargeCodesByParentIdExpenseTypesById`: ChargeCodeExpenseType
	fmt.Fprintf(os.Stdout, "Response from `ChargeCodeExpenseTypesAPI.PatchTimeChargeCodesByParentIdExpenseTypesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | expenseTypeId | 
**parentId** | **int32** | chargeCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTimeChargeCodesByParentIdExpenseTypesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ChargeCodeExpenseType**](ChargeCodeExpenseType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTimeChargeCodesByParentIdExpenseTypes

> ChargeCodeExpenseType PostTimeChargeCodesByParentIdExpenseTypes(ctx, parentId).ClientId(clientId).ChargeCodeExpenseType(chargeCodeExpenseType).Execute()

Post ChargeCodeExpenseType

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
	parentId := int32(56) // int32 | chargeCodeId
	clientId := "clientId_example" // string | 
	chargeCodeExpenseType := *openapiclient.NewChargeCodeExpenseType(*openapiclient.NewExpenseTypeReference()) // ChargeCodeExpenseType | chargeCodeExpenseType

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChargeCodeExpenseTypesAPI.PostTimeChargeCodesByParentIdExpenseTypes(context.Background(), parentId).ClientId(clientId).ChargeCodeExpenseType(chargeCodeExpenseType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChargeCodeExpenseTypesAPI.PostTimeChargeCodesByParentIdExpenseTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTimeChargeCodesByParentIdExpenseTypes`: ChargeCodeExpenseType
	fmt.Fprintf(os.Stdout, "Response from `ChargeCodeExpenseTypesAPI.PostTimeChargeCodesByParentIdExpenseTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | chargeCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTimeChargeCodesByParentIdExpenseTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **chargeCodeExpenseType** | [**ChargeCodeExpenseType**](ChargeCodeExpenseType.md) | chargeCodeExpenseType | 

### Return type

[**ChargeCodeExpenseType**](ChargeCodeExpenseType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTimeChargeCodesByParentIdExpenseTypesById

> ChargeCodeExpenseType PutTimeChargeCodesByParentIdExpenseTypesById(ctx, id, parentId).ClientId(clientId).ChargeCodeExpenseType(chargeCodeExpenseType).Execute()

Put ChargeCodeExpenseType

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
	id := int32(56) // int32 | expenseTypeId
	parentId := int32(56) // int32 | chargeCodeId
	clientId := "clientId_example" // string | 
	chargeCodeExpenseType := *openapiclient.NewChargeCodeExpenseType(*openapiclient.NewExpenseTypeReference()) // ChargeCodeExpenseType | chargeCodeExpenseType

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChargeCodeExpenseTypesAPI.PutTimeChargeCodesByParentIdExpenseTypesById(context.Background(), id, parentId).ClientId(clientId).ChargeCodeExpenseType(chargeCodeExpenseType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChargeCodeExpenseTypesAPI.PutTimeChargeCodesByParentIdExpenseTypesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTimeChargeCodesByParentIdExpenseTypesById`: ChargeCodeExpenseType
	fmt.Fprintf(os.Stdout, "Response from `ChargeCodeExpenseTypesAPI.PutTimeChargeCodesByParentIdExpenseTypesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | expenseTypeId | 
**parentId** | **int32** | chargeCodeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTimeChargeCodesByParentIdExpenseTypesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **chargeCodeExpenseType** | [**ChargeCodeExpenseType**](ChargeCodeExpenseType.md) | chargeCodeExpenseType | 

### Return type

[**ChargeCodeExpenseType**](ChargeCodeExpenseType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

