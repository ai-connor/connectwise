# \WarehouseBinsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProcurementWarehouseBinsById**](WarehouseBinsAPI.md#DeleteProcurementWarehouseBinsById) | **Delete** /procurement/warehouseBins/{id} | Delete WarehouseBin
[**GetProcurementWarehouseBins**](WarehouseBinsAPI.md#GetProcurementWarehouseBins) | **Get** /procurement/warehouseBins | Get List of WarehouseBin
[**GetProcurementWarehouseBinsById**](WarehouseBinsAPI.md#GetProcurementWarehouseBinsById) | **Get** /procurement/warehouseBins/{id} | Get WarehouseBin
[**GetProcurementWarehouseBinsCount**](WarehouseBinsAPI.md#GetProcurementWarehouseBinsCount) | **Get** /procurement/warehouseBins/count | Get Count of WarehouseBin
[**PatchProcurementWarehouseBinsById**](WarehouseBinsAPI.md#PatchProcurementWarehouseBinsById) | **Patch** /procurement/warehouseBins/{id} | Patch WarehouseBin
[**PostProcurementWarehouseBins**](WarehouseBinsAPI.md#PostProcurementWarehouseBins) | **Post** /procurement/warehouseBins | Post WarehouseBin
[**PutProcurementWarehouseBinsById**](WarehouseBinsAPI.md#PutProcurementWarehouseBinsById) | **Put** /procurement/warehouseBins/{id} | Put WarehouseBin



## DeleteProcurementWarehouseBinsById

> DeleteProcurementWarehouseBinsById(ctx, id).ClientId(clientId).Execute()

Delete WarehouseBin

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
	id := int32(56) // int32 | warehouseBinId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WarehouseBinsAPI.DeleteProcurementWarehouseBinsById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseBinsAPI.DeleteProcurementWarehouseBinsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | warehouseBinId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcurementWarehouseBinsByIdRequest struct via the builder pattern


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


## GetProcurementWarehouseBins

> []WarehouseBin GetProcurementWarehouseBins(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of WarehouseBin

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
	resp, r, err := apiClient.WarehouseBinsAPI.GetProcurementWarehouseBins(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseBinsAPI.GetProcurementWarehouseBins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementWarehouseBins`: []WarehouseBin
	fmt.Fprintf(os.Stdout, "Response from `WarehouseBinsAPI.GetProcurementWarehouseBins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementWarehouseBinsRequest struct via the builder pattern


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

[**[]WarehouseBin**](WarehouseBin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementWarehouseBinsById

> WarehouseBin GetProcurementWarehouseBinsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get WarehouseBin

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
	id := int32(56) // int32 | warehouseBinId
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
	resp, r, err := apiClient.WarehouseBinsAPI.GetProcurementWarehouseBinsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseBinsAPI.GetProcurementWarehouseBinsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementWarehouseBinsById`: WarehouseBin
	fmt.Fprintf(os.Stdout, "Response from `WarehouseBinsAPI.GetProcurementWarehouseBinsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | warehouseBinId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementWarehouseBinsByIdRequest struct via the builder pattern


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

[**WarehouseBin**](WarehouseBin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementWarehouseBinsCount

> Count GetProcurementWarehouseBinsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of WarehouseBin

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
	resp, r, err := apiClient.WarehouseBinsAPI.GetProcurementWarehouseBinsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseBinsAPI.GetProcurementWarehouseBinsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementWarehouseBinsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `WarehouseBinsAPI.GetProcurementWarehouseBinsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementWarehouseBinsCountRequest struct via the builder pattern


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


## PatchProcurementWarehouseBinsById

> WarehouseBin PatchProcurementWarehouseBinsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch WarehouseBin

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
	id := int32(56) // int32 | warehouseBinId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehouseBinsAPI.PatchProcurementWarehouseBinsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseBinsAPI.PatchProcurementWarehouseBinsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProcurementWarehouseBinsById`: WarehouseBin
	fmt.Fprintf(os.Stdout, "Response from `WarehouseBinsAPI.PatchProcurementWarehouseBinsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | warehouseBinId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProcurementWarehouseBinsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**WarehouseBin**](WarehouseBin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProcurementWarehouseBins

> WarehouseBin PostProcurementWarehouseBins(ctx).ClientId(clientId).WarehouseBin(warehouseBin).Execute()

Post WarehouseBin

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
	warehouseBin := *openapiclient.NewWarehouseBin("Name_example", *openapiclient.NewWarehouseReference()) // WarehouseBin | warehouseBin

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehouseBinsAPI.PostProcurementWarehouseBins(context.Background()).ClientId(clientId).WarehouseBin(warehouseBin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseBinsAPI.PostProcurementWarehouseBins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementWarehouseBins`: WarehouseBin
	fmt.Fprintf(os.Stdout, "Response from `WarehouseBinsAPI.PostProcurementWarehouseBins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementWarehouseBinsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **warehouseBin** | [**WarehouseBin**](WarehouseBin.md) | warehouseBin | 

### Return type

[**WarehouseBin**](WarehouseBin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProcurementWarehouseBinsById

> WarehouseBin PutProcurementWarehouseBinsById(ctx, id).ClientId(clientId).WarehouseBin(warehouseBin).Execute()

Put WarehouseBin

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
	id := int32(56) // int32 | warehouseBinId
	clientId := "clientId_example" // string | 
	warehouseBin := *openapiclient.NewWarehouseBin("Name_example", *openapiclient.NewWarehouseReference()) // WarehouseBin | warehouseBin

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehouseBinsAPI.PutProcurementWarehouseBinsById(context.Background(), id).ClientId(clientId).WarehouseBin(warehouseBin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseBinsAPI.PutProcurementWarehouseBinsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProcurementWarehouseBinsById`: WarehouseBin
	fmt.Fprintf(os.Stdout, "Response from `WarehouseBinsAPI.PutProcurementWarehouseBinsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | warehouseBinId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProcurementWarehouseBinsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **warehouseBin** | [**WarehouseBin**](WarehouseBin.md) | warehouseBin | 

### Return type

[**WarehouseBin**](WarehouseBin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

