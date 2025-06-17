# \UnitOfMeasuresAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProcurementUnitOfMeasuresById**](UnitOfMeasuresAPI.md#DeleteProcurementUnitOfMeasuresById) | **Delete** /procurement/unitOfMeasures/{id} | Delete UnitOfMeasure
[**GetProcurementUnitOfMeasures**](UnitOfMeasuresAPI.md#GetProcurementUnitOfMeasures) | **Get** /procurement/unitOfMeasures | Get List of UnitOfMeasure
[**GetProcurementUnitOfMeasuresById**](UnitOfMeasuresAPI.md#GetProcurementUnitOfMeasuresById) | **Get** /procurement/unitOfMeasures/{id} | Get UnitOfMeasure
[**GetProcurementUnitOfMeasuresCount**](UnitOfMeasuresAPI.md#GetProcurementUnitOfMeasuresCount) | **Get** /procurement/unitOfMeasures/count | Get Count of UnitOfMeasure
[**PatchProcurementUnitOfMeasuresById**](UnitOfMeasuresAPI.md#PatchProcurementUnitOfMeasuresById) | **Patch** /procurement/unitOfMeasures/{id} | Patch UnitOfMeasure
[**PostProcurementUnitOfMeasures**](UnitOfMeasuresAPI.md#PostProcurementUnitOfMeasures) | **Post** /procurement/unitOfMeasures | Post UnitOfMeasure
[**PutProcurementUnitOfMeasuresById**](UnitOfMeasuresAPI.md#PutProcurementUnitOfMeasuresById) | **Put** /procurement/unitOfMeasures/{id} | Put UnitOfMeasure



## DeleteProcurementUnitOfMeasuresById

> DeleteProcurementUnitOfMeasuresById(ctx, id).ClientId(clientId).Execute()

Delete UnitOfMeasure

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
	id := int32(56) // int32 | unitOfMeasureId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UnitOfMeasuresAPI.DeleteProcurementUnitOfMeasuresById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnitOfMeasuresAPI.DeleteProcurementUnitOfMeasuresById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | unitOfMeasureId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcurementUnitOfMeasuresByIdRequest struct via the builder pattern


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


## GetProcurementUnitOfMeasures

> []UnitOfMeasure GetProcurementUnitOfMeasures(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of UnitOfMeasure

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
	resp, r, err := apiClient.UnitOfMeasuresAPI.GetProcurementUnitOfMeasures(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnitOfMeasuresAPI.GetProcurementUnitOfMeasures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementUnitOfMeasures`: []UnitOfMeasure
	fmt.Fprintf(os.Stdout, "Response from `UnitOfMeasuresAPI.GetProcurementUnitOfMeasures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementUnitOfMeasuresRequest struct via the builder pattern


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

[**[]UnitOfMeasure**](UnitOfMeasure.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementUnitOfMeasuresById

> UnitOfMeasure GetProcurementUnitOfMeasuresById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get UnitOfMeasure

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
	id := int32(56) // int32 | unitOfMeasureId
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
	resp, r, err := apiClient.UnitOfMeasuresAPI.GetProcurementUnitOfMeasuresById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnitOfMeasuresAPI.GetProcurementUnitOfMeasuresById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementUnitOfMeasuresById`: UnitOfMeasure
	fmt.Fprintf(os.Stdout, "Response from `UnitOfMeasuresAPI.GetProcurementUnitOfMeasuresById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | unitOfMeasureId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementUnitOfMeasuresByIdRequest struct via the builder pattern


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

[**UnitOfMeasure**](UnitOfMeasure.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementUnitOfMeasuresCount

> Count GetProcurementUnitOfMeasuresCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of UnitOfMeasure

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
	resp, r, err := apiClient.UnitOfMeasuresAPI.GetProcurementUnitOfMeasuresCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnitOfMeasuresAPI.GetProcurementUnitOfMeasuresCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementUnitOfMeasuresCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `UnitOfMeasuresAPI.GetProcurementUnitOfMeasuresCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementUnitOfMeasuresCountRequest struct via the builder pattern


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


## PatchProcurementUnitOfMeasuresById

> UnitOfMeasure PatchProcurementUnitOfMeasuresById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch UnitOfMeasure

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
	id := int32(56) // int32 | unitOfMeasureId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UnitOfMeasuresAPI.PatchProcurementUnitOfMeasuresById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnitOfMeasuresAPI.PatchProcurementUnitOfMeasuresById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProcurementUnitOfMeasuresById`: UnitOfMeasure
	fmt.Fprintf(os.Stdout, "Response from `UnitOfMeasuresAPI.PatchProcurementUnitOfMeasuresById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | unitOfMeasureId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProcurementUnitOfMeasuresByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**UnitOfMeasure**](UnitOfMeasure.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProcurementUnitOfMeasures

> UnitOfMeasure PostProcurementUnitOfMeasures(ctx).ClientId(clientId).UnitOfMeasure(unitOfMeasure).Execute()

Post UnitOfMeasure

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
	unitOfMeasure := *openapiclient.NewUnitOfMeasure("Name_example") // UnitOfMeasure | unitOfMeasure

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UnitOfMeasuresAPI.PostProcurementUnitOfMeasures(context.Background()).ClientId(clientId).UnitOfMeasure(unitOfMeasure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnitOfMeasuresAPI.PostProcurementUnitOfMeasures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementUnitOfMeasures`: UnitOfMeasure
	fmt.Fprintf(os.Stdout, "Response from `UnitOfMeasuresAPI.PostProcurementUnitOfMeasures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementUnitOfMeasuresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **unitOfMeasure** | [**UnitOfMeasure**](UnitOfMeasure.md) | unitOfMeasure | 

### Return type

[**UnitOfMeasure**](UnitOfMeasure.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProcurementUnitOfMeasuresById

> UnitOfMeasure PutProcurementUnitOfMeasuresById(ctx, id).ClientId(clientId).UnitOfMeasure(unitOfMeasure).Execute()

Put UnitOfMeasure

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
	id := int32(56) // int32 | unitOfMeasureId
	clientId := "clientId_example" // string | 
	unitOfMeasure := *openapiclient.NewUnitOfMeasure("Name_example") // UnitOfMeasure | unitOfMeasure

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UnitOfMeasuresAPI.PutProcurementUnitOfMeasuresById(context.Background(), id).ClientId(clientId).UnitOfMeasure(unitOfMeasure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnitOfMeasuresAPI.PutProcurementUnitOfMeasuresById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProcurementUnitOfMeasuresById`: UnitOfMeasure
	fmt.Fprintf(os.Stdout, "Response from `UnitOfMeasuresAPI.PutProcurementUnitOfMeasuresById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | unitOfMeasureId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProcurementUnitOfMeasuresByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **unitOfMeasure** | [**UnitOfMeasure**](UnitOfMeasure.md) | unitOfMeasure | 

### Return type

[**UnitOfMeasure**](UnitOfMeasure.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

