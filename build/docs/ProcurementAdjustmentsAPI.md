# \ProcurementAdjustmentsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProcurementAdjustmentsById**](ProcurementAdjustmentsAPI.md#DeleteProcurementAdjustmentsById) | **Delete** /procurement/adjustments/{id} | Delete ProcurementAdjustment
[**GetProcurementAdjustments**](ProcurementAdjustmentsAPI.md#GetProcurementAdjustments) | **Get** /procurement/adjustments | Get List of ProcurementAdjustment
[**GetProcurementAdjustmentsById**](ProcurementAdjustmentsAPI.md#GetProcurementAdjustmentsById) | **Get** /procurement/adjustments/{id} | Get ProcurementAdjustment
[**GetProcurementAdjustmentsCount**](ProcurementAdjustmentsAPI.md#GetProcurementAdjustmentsCount) | **Get** /procurement/adjustments/count | Get Count of ProcurementAdjustment
[**PatchProcurementAdjustmentsById**](ProcurementAdjustmentsAPI.md#PatchProcurementAdjustmentsById) | **Patch** /procurement/adjustments/{id} | Patch ProcurementAdjustment
[**PostProcurementAdjustments**](ProcurementAdjustmentsAPI.md#PostProcurementAdjustments) | **Post** /procurement/adjustments | Post ProcurementAdjustment
[**PutProcurementAdjustmentsById**](ProcurementAdjustmentsAPI.md#PutProcurementAdjustmentsById) | **Put** /procurement/adjustments/{id} | Put ProcurementAdjustment



## DeleteProcurementAdjustmentsById

> DeleteProcurementAdjustmentsById(ctx, id).ClientId(clientId).Execute()

Delete ProcurementAdjustment

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
	id := int32(56) // int32 | adjustmentId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProcurementAdjustmentsAPI.DeleteProcurementAdjustmentsById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcurementAdjustmentsAPI.DeleteProcurementAdjustmentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | adjustmentId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcurementAdjustmentsByIdRequest struct via the builder pattern


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


## GetProcurementAdjustments

> []ProcurementAdjustment GetProcurementAdjustments(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProcurementAdjustment

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
	resp, r, err := apiClient.ProcurementAdjustmentsAPI.GetProcurementAdjustments(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcurementAdjustmentsAPI.GetProcurementAdjustments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementAdjustments`: []ProcurementAdjustment
	fmt.Fprintf(os.Stdout, "Response from `ProcurementAdjustmentsAPI.GetProcurementAdjustments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementAdjustmentsRequest struct via the builder pattern


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

[**[]ProcurementAdjustment**](ProcurementAdjustment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementAdjustmentsById

> ProcurementAdjustment GetProcurementAdjustmentsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ProcurementAdjustment

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
	id := int32(56) // int32 | adjustmentId
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
	resp, r, err := apiClient.ProcurementAdjustmentsAPI.GetProcurementAdjustmentsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcurementAdjustmentsAPI.GetProcurementAdjustmentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementAdjustmentsById`: ProcurementAdjustment
	fmt.Fprintf(os.Stdout, "Response from `ProcurementAdjustmentsAPI.GetProcurementAdjustmentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | adjustmentId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementAdjustmentsByIdRequest struct via the builder pattern


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

[**ProcurementAdjustment**](ProcurementAdjustment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementAdjustmentsCount

> Count GetProcurementAdjustmentsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ProcurementAdjustment

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
	resp, r, err := apiClient.ProcurementAdjustmentsAPI.GetProcurementAdjustmentsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcurementAdjustmentsAPI.GetProcurementAdjustmentsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementAdjustmentsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ProcurementAdjustmentsAPI.GetProcurementAdjustmentsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementAdjustmentsCountRequest struct via the builder pattern


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


## PatchProcurementAdjustmentsById

> ProcurementAdjustment PatchProcurementAdjustmentsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ProcurementAdjustment

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
	id := int32(56) // int32 | adjustmentId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcurementAdjustmentsAPI.PatchProcurementAdjustmentsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcurementAdjustmentsAPI.PatchProcurementAdjustmentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProcurementAdjustmentsById`: ProcurementAdjustment
	fmt.Fprintf(os.Stdout, "Response from `ProcurementAdjustmentsAPI.PatchProcurementAdjustmentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | adjustmentId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProcurementAdjustmentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ProcurementAdjustment**](ProcurementAdjustment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProcurementAdjustments

> ProcurementAdjustment PostProcurementAdjustments(ctx).ClientId(clientId).ProcurementAdjustment(procurementAdjustment).Execute()

Post ProcurementAdjustment

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
	procurementAdjustment := *openapiclient.NewProcurementAdjustment("Identifier_example", *openapiclient.NewAdjustmentTypeReference()) // ProcurementAdjustment | adjustment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcurementAdjustmentsAPI.PostProcurementAdjustments(context.Background()).ClientId(clientId).ProcurementAdjustment(procurementAdjustment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcurementAdjustmentsAPI.PostProcurementAdjustments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementAdjustments`: ProcurementAdjustment
	fmt.Fprintf(os.Stdout, "Response from `ProcurementAdjustmentsAPI.PostProcurementAdjustments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementAdjustmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **procurementAdjustment** | [**ProcurementAdjustment**](ProcurementAdjustment.md) | adjustment | 

### Return type

[**ProcurementAdjustment**](ProcurementAdjustment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProcurementAdjustmentsById

> ProcurementAdjustment PutProcurementAdjustmentsById(ctx, id).ClientId(clientId).ProcurementAdjustment(procurementAdjustment).Execute()

Put ProcurementAdjustment

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
	id := int32(56) // int32 | adjustmentId
	clientId := "clientId_example" // string | 
	procurementAdjustment := *openapiclient.NewProcurementAdjustment("Identifier_example", *openapiclient.NewAdjustmentTypeReference()) // ProcurementAdjustment | adjustment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProcurementAdjustmentsAPI.PutProcurementAdjustmentsById(context.Background(), id).ClientId(clientId).ProcurementAdjustment(procurementAdjustment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProcurementAdjustmentsAPI.PutProcurementAdjustmentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProcurementAdjustmentsById`: ProcurementAdjustment
	fmt.Fprintf(os.Stdout, "Response from `ProcurementAdjustmentsAPI.PutProcurementAdjustmentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | adjustmentId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProcurementAdjustmentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **procurementAdjustment** | [**ProcurementAdjustment**](ProcurementAdjustment.md) | adjustment | 

### Return type

[**ProcurementAdjustment**](ProcurementAdjustment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

