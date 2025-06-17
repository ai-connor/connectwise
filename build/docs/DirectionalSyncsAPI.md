# \DirectionalSyncsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProcurementDirectionalSyncsById**](DirectionalSyncsAPI.md#DeleteProcurementDirectionalSyncsById) | **Delete** /procurement/directionalSyncs/{id} | Delete DirectionalSync
[**GetProcurementDirectionalSyncs**](DirectionalSyncsAPI.md#GetProcurementDirectionalSyncs) | **Get** /procurement/directionalSyncs | Get List of DirectionalSync
[**GetProcurementDirectionalSyncsById**](DirectionalSyncsAPI.md#GetProcurementDirectionalSyncsById) | **Get** /procurement/directionalSyncs/{id} | Get DirectionalSync
[**GetProcurementDirectionalSyncsCount**](DirectionalSyncsAPI.md#GetProcurementDirectionalSyncsCount) | **Get** /procurement/directionalSyncs/count | Get Count of DirectionalSync
[**PatchProcurementDirectionalSyncsById**](DirectionalSyncsAPI.md#PatchProcurementDirectionalSyncsById) | **Patch** /procurement/directionalSyncs/{id} | Patch DirectionalSync
[**PostProcurementDirectionalSyncs**](DirectionalSyncsAPI.md#PostProcurementDirectionalSyncs) | **Post** /procurement/directionalSyncs | Post DirectionalSync
[**PutProcurementDirectionalSyncsById**](DirectionalSyncsAPI.md#PutProcurementDirectionalSyncsById) | **Put** /procurement/directionalSyncs/{id} | Put DirectionalSync



## DeleteProcurementDirectionalSyncsById

> DeleteProcurementDirectionalSyncsById(ctx, id).ClientId(clientId).Execute()

Delete DirectionalSync

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
	id := int32(56) // int32 | warehousId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DirectionalSyncsAPI.DeleteProcurementDirectionalSyncsById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectionalSyncsAPI.DeleteProcurementDirectionalSyncsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | warehousId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcurementDirectionalSyncsByIdRequest struct via the builder pattern


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


## GetProcurementDirectionalSyncs

> []DirectionalSync GetProcurementDirectionalSyncs(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of DirectionalSync

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
	resp, r, err := apiClient.DirectionalSyncsAPI.GetProcurementDirectionalSyncs(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectionalSyncsAPI.GetProcurementDirectionalSyncs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementDirectionalSyncs`: []DirectionalSync
	fmt.Fprintf(os.Stdout, "Response from `DirectionalSyncsAPI.GetProcurementDirectionalSyncs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementDirectionalSyncsRequest struct via the builder pattern


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

[**[]DirectionalSync**](DirectionalSync.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementDirectionalSyncsById

> DirectionalSync GetProcurementDirectionalSyncsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get DirectionalSync

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
	id := int32(56) // int32 | warehousId
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
	resp, r, err := apiClient.DirectionalSyncsAPI.GetProcurementDirectionalSyncsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectionalSyncsAPI.GetProcurementDirectionalSyncsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementDirectionalSyncsById`: DirectionalSync
	fmt.Fprintf(os.Stdout, "Response from `DirectionalSyncsAPI.GetProcurementDirectionalSyncsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | warehousId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementDirectionalSyncsByIdRequest struct via the builder pattern


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

[**DirectionalSync**](DirectionalSync.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementDirectionalSyncsCount

> Count GetProcurementDirectionalSyncsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of DirectionalSync

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
	resp, r, err := apiClient.DirectionalSyncsAPI.GetProcurementDirectionalSyncsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectionalSyncsAPI.GetProcurementDirectionalSyncsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementDirectionalSyncsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `DirectionalSyncsAPI.GetProcurementDirectionalSyncsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementDirectionalSyncsCountRequest struct via the builder pattern


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


## PatchProcurementDirectionalSyncsById

> DirectionalSync PatchProcurementDirectionalSyncsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch DirectionalSync

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
	id := int32(56) // int32 | warehousId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectionalSyncsAPI.PatchProcurementDirectionalSyncsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectionalSyncsAPI.PatchProcurementDirectionalSyncsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProcurementDirectionalSyncsById`: DirectionalSync
	fmt.Fprintf(os.Stdout, "Response from `DirectionalSyncsAPI.PatchProcurementDirectionalSyncsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | warehousId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProcurementDirectionalSyncsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**DirectionalSync**](DirectionalSync.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProcurementDirectionalSyncs

> DirectionalSync PostProcurementDirectionalSyncs(ctx).ClientId(clientId).DirectionalSync(directionalSync).Execute()

Post DirectionalSync

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
	directionalSync := *openapiclient.NewDirectionalSync("Name_example") // DirectionalSync | directionalSync

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectionalSyncsAPI.PostProcurementDirectionalSyncs(context.Background()).ClientId(clientId).DirectionalSync(directionalSync).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectionalSyncsAPI.PostProcurementDirectionalSyncs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementDirectionalSyncs`: DirectionalSync
	fmt.Fprintf(os.Stdout, "Response from `DirectionalSyncsAPI.PostProcurementDirectionalSyncs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementDirectionalSyncsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **directionalSync** | [**DirectionalSync**](DirectionalSync.md) | directionalSync | 

### Return type

[**DirectionalSync**](DirectionalSync.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProcurementDirectionalSyncsById

> DirectionalSync PutProcurementDirectionalSyncsById(ctx, id).ClientId(clientId).DirectionalSync(directionalSync).Execute()

Put DirectionalSync

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
	id := int32(56) // int32 | warehousId
	clientId := "clientId_example" // string | 
	directionalSync := *openapiclient.NewDirectionalSync("Name_example") // DirectionalSync | directionalSync

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectionalSyncsAPI.PutProcurementDirectionalSyncsById(context.Background(), id).ClientId(clientId).DirectionalSync(directionalSync).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectionalSyncsAPI.PutProcurementDirectionalSyncsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProcurementDirectionalSyncsById`: DirectionalSync
	fmt.Fprintf(os.Stdout, "Response from `DirectionalSyncsAPI.PutProcurementDirectionalSyncsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | warehousId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProcurementDirectionalSyncsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **directionalSync** | [**DirectionalSync**](DirectionalSync.md) | directionalSync | 

### Return type

[**DirectionalSync**](DirectionalSync.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

