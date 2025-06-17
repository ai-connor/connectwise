# \AllowedFileTypeAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemAllowedfiletypesById**](AllowedFileTypeAPI.md#DeleteSystemAllowedfiletypesById) | **Delete** /system/allowedfiletypes/{id} | Delete AllowedFileType
[**GetSystemAllowedfiletypes**](AllowedFileTypeAPI.md#GetSystemAllowedfiletypes) | **Get** /system/allowedfiletypes/ | Get List of AllowedFileType
[**GetSystemAllowedfiletypesById**](AllowedFileTypeAPI.md#GetSystemAllowedfiletypesById) | **Get** /system/allowedfiletypes/{id} | Get AllowedFileType
[**GetSystemAllowedfiletypesCount**](AllowedFileTypeAPI.md#GetSystemAllowedfiletypesCount) | **Get** /system/allowedfiletypes/count | Get Count of AllowedFileType
[**PatchSystemAllowedfiletypesById**](AllowedFileTypeAPI.md#PatchSystemAllowedfiletypesById) | **Patch** /system/allowedfiletypes/{id} | Patch AllowedFileType
[**PostSystemAllowedFileTypes**](AllowedFileTypeAPI.md#PostSystemAllowedFileTypes) | **Post** /system/AllowedFileTypes/ | Post AllowedFileType
[**PutSystemAllowedfiletypesById**](AllowedFileTypeAPI.md#PutSystemAllowedfiletypesById) | **Put** /system/allowedfiletypes/{id} | Put AllowedFileType



## DeleteSystemAllowedfiletypesById

> DeleteSystemAllowedfiletypesById(ctx, id).ClientId(clientId).Execute()

Delete AllowedFileType

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
	id := int32(56) // int32 | allowedfiletypeId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AllowedFileTypeAPI.DeleteSystemAllowedfiletypesById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedFileTypeAPI.DeleteSystemAllowedfiletypesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | allowedfiletypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemAllowedfiletypesByIdRequest struct via the builder pattern


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


## GetSystemAllowedfiletypes

> []AllowedFileType GetSystemAllowedfiletypes(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of AllowedFileType

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
	resp, r, err := apiClient.AllowedFileTypeAPI.GetSystemAllowedfiletypes(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedFileTypeAPI.GetSystemAllowedfiletypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemAllowedfiletypes`: []AllowedFileType
	fmt.Fprintf(os.Stdout, "Response from `AllowedFileTypeAPI.GetSystemAllowedfiletypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemAllowedfiletypesRequest struct via the builder pattern


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

[**[]AllowedFileType**](AllowedFileType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemAllowedfiletypesById

> AllowedFileType GetSystemAllowedfiletypesById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get AllowedFileType

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
	id := int32(56) // int32 | allowedfiletypeId
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
	resp, r, err := apiClient.AllowedFileTypeAPI.GetSystemAllowedfiletypesById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedFileTypeAPI.GetSystemAllowedfiletypesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemAllowedfiletypesById`: AllowedFileType
	fmt.Fprintf(os.Stdout, "Response from `AllowedFileTypeAPI.GetSystemAllowedfiletypesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | allowedfiletypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemAllowedfiletypesByIdRequest struct via the builder pattern


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

[**AllowedFileType**](AllowedFileType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemAllowedfiletypesCount

> Count GetSystemAllowedfiletypesCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of AllowedFileType

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
	resp, r, err := apiClient.AllowedFileTypeAPI.GetSystemAllowedfiletypesCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedFileTypeAPI.GetSystemAllowedfiletypesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemAllowedfiletypesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `AllowedFileTypeAPI.GetSystemAllowedfiletypesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemAllowedfiletypesCountRequest struct via the builder pattern


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


## PatchSystemAllowedfiletypesById

> AllowedFileType PatchSystemAllowedfiletypesById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch AllowedFileType

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
	id := int32(56) // int32 | allowedFileTypesId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllowedFileTypeAPI.PatchSystemAllowedfiletypesById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedFileTypeAPI.PatchSystemAllowedfiletypesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemAllowedfiletypesById`: AllowedFileType
	fmt.Fprintf(os.Stdout, "Response from `AllowedFileTypeAPI.PatchSystemAllowedfiletypesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | allowedFileTypesId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemAllowedfiletypesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**AllowedFileType**](AllowedFileType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemAllowedFileTypes

> AllowedFileType PostSystemAllowedFileTypes(ctx).ClientId(clientId).AllowedFileType(allowedFileType).Execute()

Post AllowedFileType

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
	allowedFileType := *openapiclient.NewAllowedFileType("FileType_example") // AllowedFileType | allowedFileType

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllowedFileTypeAPI.PostSystemAllowedFileTypes(context.Background()).ClientId(clientId).AllowedFileType(allowedFileType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedFileTypeAPI.PostSystemAllowedFileTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemAllowedFileTypes`: AllowedFileType
	fmt.Fprintf(os.Stdout, "Response from `AllowedFileTypeAPI.PostSystemAllowedFileTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemAllowedFileTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **allowedFileType** | [**AllowedFileType**](AllowedFileType.md) | allowedFileType | 

### Return type

[**AllowedFileType**](AllowedFileType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemAllowedfiletypesById

> AllowedFileType PutSystemAllowedfiletypesById(ctx, id).ClientId(clientId).AllowedFileType(allowedFileType).Execute()

Put AllowedFileType

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
	id := int32(56) // int32 | allowedFileTypeId
	clientId := "clientId_example" // string | 
	allowedFileType := *openapiclient.NewAllowedFileType("FileType_example") // AllowedFileType | AllowedFileType

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AllowedFileTypeAPI.PutSystemAllowedfiletypesById(context.Background(), id).ClientId(clientId).AllowedFileType(allowedFileType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AllowedFileTypeAPI.PutSystemAllowedfiletypesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemAllowedfiletypesById`: AllowedFileType
	fmt.Fprintf(os.Stdout, "Response from `AllowedFileTypeAPI.PutSystemAllowedfiletypesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | allowedFileTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemAllowedfiletypesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **allowedFileType** | [**AllowedFileType**](AllowedFileType.md) | AllowedFileType | 

### Return type

[**AllowedFileType**](AllowedFileType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

