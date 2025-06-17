# \EmailConnectorParsingStylesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemEmailConnectorsByParentIdParsingStylesById**](EmailConnectorParsingStylesAPI.md#DeleteSystemEmailConnectorsByParentIdParsingStylesById) | **Delete** /system/emailConnectors/{parentId}/parsingStyles/{id} | Delete EmailConnectorParsingStyle
[**GetSystemEmailConnectorsByParentIdParsingStyles**](EmailConnectorParsingStylesAPI.md#GetSystemEmailConnectorsByParentIdParsingStyles) | **Get** /system/emailConnectors/{parentId}/parsingStyles | Get List of EmailConnectorParsingStyle
[**GetSystemEmailConnectorsByParentIdParsingStylesById**](EmailConnectorParsingStylesAPI.md#GetSystemEmailConnectorsByParentIdParsingStylesById) | **Get** /system/emailConnectors/{parentId}/parsingStyles/{id} | Get EmailConnectorParsingStyle
[**GetSystemEmailConnectorsByParentIdParsingStylesCount**](EmailConnectorParsingStylesAPI.md#GetSystemEmailConnectorsByParentIdParsingStylesCount) | **Get** /system/emailConnectors/{parentId}/parsingStyles/count | Get Count of EmailConnectorParsingStyle
[**PatchSystemEmailConnectorsByParentIdParsingStylesById**](EmailConnectorParsingStylesAPI.md#PatchSystemEmailConnectorsByParentIdParsingStylesById) | **Patch** /system/emailConnectors/{parentId}/parsingStyles/{id} | Patch EmailConnectorParsingStyle
[**PostSystemEmailConnectorsByParentIdParsingStyles**](EmailConnectorParsingStylesAPI.md#PostSystemEmailConnectorsByParentIdParsingStyles) | **Post** /system/emailConnectors/{parentId}/parsingStyles | Post EmailConnectorParsingStyle
[**PutSystemEmailConnectorsByParentIdParsingStylesById**](EmailConnectorParsingStylesAPI.md#PutSystemEmailConnectorsByParentIdParsingStylesById) | **Put** /system/emailConnectors/{parentId}/parsingStyles/{id} | Put EmailConnectorParsingStyle



## DeleteSystemEmailConnectorsByParentIdParsingStylesById

> DeleteSystemEmailConnectorsByParentIdParsingStylesById(ctx, id, parentId).ClientId(clientId).Execute()

Delete EmailConnectorParsingStyle

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
	id := int32(56) // int32 | parsingStyleId
	parentId := int32(56) // int32 | emailConnectorId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmailConnectorParsingStylesAPI.DeleteSystemEmailConnectorsByParentIdParsingStylesById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailConnectorParsingStylesAPI.DeleteSystemEmailConnectorsByParentIdParsingStylesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | parsingStyleId | 
**parentId** | **int32** | emailConnectorId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemEmailConnectorsByParentIdParsingStylesByIdRequest struct via the builder pattern


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


## GetSystemEmailConnectorsByParentIdParsingStyles

> []EmailConnectorParsingStyle GetSystemEmailConnectorsByParentIdParsingStyles(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of EmailConnectorParsingStyle

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
	parentId := int32(56) // int32 | emailConnectorId
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
	resp, r, err := apiClient.EmailConnectorParsingStylesAPI.GetSystemEmailConnectorsByParentIdParsingStyles(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailConnectorParsingStylesAPI.GetSystemEmailConnectorsByParentIdParsingStyles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemEmailConnectorsByParentIdParsingStyles`: []EmailConnectorParsingStyle
	fmt.Fprintf(os.Stdout, "Response from `EmailConnectorParsingStylesAPI.GetSystemEmailConnectorsByParentIdParsingStyles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | emailConnectorId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemEmailConnectorsByParentIdParsingStylesRequest struct via the builder pattern


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

[**[]EmailConnectorParsingStyle**](EmailConnectorParsingStyle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemEmailConnectorsByParentIdParsingStylesById

> EmailConnectorParsingStyle GetSystemEmailConnectorsByParentIdParsingStylesById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get EmailConnectorParsingStyle

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
	id := int32(56) // int32 | parsingStyleId
	parentId := int32(56) // int32 | emailConnectorId
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
	resp, r, err := apiClient.EmailConnectorParsingStylesAPI.GetSystemEmailConnectorsByParentIdParsingStylesById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailConnectorParsingStylesAPI.GetSystemEmailConnectorsByParentIdParsingStylesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemEmailConnectorsByParentIdParsingStylesById`: EmailConnectorParsingStyle
	fmt.Fprintf(os.Stdout, "Response from `EmailConnectorParsingStylesAPI.GetSystemEmailConnectorsByParentIdParsingStylesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | parsingStyleId | 
**parentId** | **int32** | emailConnectorId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemEmailConnectorsByParentIdParsingStylesByIdRequest struct via the builder pattern


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

[**EmailConnectorParsingStyle**](EmailConnectorParsingStyle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemEmailConnectorsByParentIdParsingStylesCount

> Count GetSystemEmailConnectorsByParentIdParsingStylesCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of EmailConnectorParsingStyle

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
	parentId := int32(56) // int32 | emailConnectorId
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
	resp, r, err := apiClient.EmailConnectorParsingStylesAPI.GetSystemEmailConnectorsByParentIdParsingStylesCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailConnectorParsingStylesAPI.GetSystemEmailConnectorsByParentIdParsingStylesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemEmailConnectorsByParentIdParsingStylesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `EmailConnectorParsingStylesAPI.GetSystemEmailConnectorsByParentIdParsingStylesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | emailConnectorId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemEmailConnectorsByParentIdParsingStylesCountRequest struct via the builder pattern


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


## PatchSystemEmailConnectorsByParentIdParsingStylesById

> EmailConnectorParsingStyle PatchSystemEmailConnectorsByParentIdParsingStylesById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch EmailConnectorParsingStyle

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
	id := int32(56) // int32 | parsingStyleId
	parentId := int32(56) // int32 | emailConnectorId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailConnectorParsingStylesAPI.PatchSystemEmailConnectorsByParentIdParsingStylesById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailConnectorParsingStylesAPI.PatchSystemEmailConnectorsByParentIdParsingStylesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemEmailConnectorsByParentIdParsingStylesById`: EmailConnectorParsingStyle
	fmt.Fprintf(os.Stdout, "Response from `EmailConnectorParsingStylesAPI.PatchSystemEmailConnectorsByParentIdParsingStylesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | parsingStyleId | 
**parentId** | **int32** | emailConnectorId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemEmailConnectorsByParentIdParsingStylesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**EmailConnectorParsingStyle**](EmailConnectorParsingStyle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemEmailConnectorsByParentIdParsingStyles

> EmailConnectorParsingStyle PostSystemEmailConnectorsByParentIdParsingStyles(ctx, parentId).ClientId(clientId).EmailConnectorParsingStyle(emailConnectorParsingStyle).Execute()

Post EmailConnectorParsingStyle

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
	parentId := int32(56) // int32 | emailConnectorId
	clientId := "clientId_example" // string | 
	emailConnectorParsingStyle := *openapiclient.NewEmailConnectorParsingStyle(*openapiclient.NewEmailConnectorParsingTypeReference(), "ParseRule_example", NullableInt32(123)) // EmailConnectorParsingStyle | emailConnectorParsingStyle

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailConnectorParsingStylesAPI.PostSystemEmailConnectorsByParentIdParsingStyles(context.Background(), parentId).ClientId(clientId).EmailConnectorParsingStyle(emailConnectorParsingStyle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailConnectorParsingStylesAPI.PostSystemEmailConnectorsByParentIdParsingStyles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemEmailConnectorsByParentIdParsingStyles`: EmailConnectorParsingStyle
	fmt.Fprintf(os.Stdout, "Response from `EmailConnectorParsingStylesAPI.PostSystemEmailConnectorsByParentIdParsingStyles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | emailConnectorId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemEmailConnectorsByParentIdParsingStylesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **emailConnectorParsingStyle** | [**EmailConnectorParsingStyle**](EmailConnectorParsingStyle.md) | emailConnectorParsingStyle | 

### Return type

[**EmailConnectorParsingStyle**](EmailConnectorParsingStyle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemEmailConnectorsByParentIdParsingStylesById

> EmailConnectorParsingStyle PutSystemEmailConnectorsByParentIdParsingStylesById(ctx, id, parentId).ClientId(clientId).EmailConnectorParsingStyle(emailConnectorParsingStyle).Execute()

Put EmailConnectorParsingStyle

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
	id := int32(56) // int32 | parsingStyleId
	parentId := int32(56) // int32 | emailConnectorId
	clientId := "clientId_example" // string | 
	emailConnectorParsingStyle := *openapiclient.NewEmailConnectorParsingStyle(*openapiclient.NewEmailConnectorParsingTypeReference(), "ParseRule_example", NullableInt32(123)) // EmailConnectorParsingStyle | emailConnectorParsingStyle

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailConnectorParsingStylesAPI.PutSystemEmailConnectorsByParentIdParsingStylesById(context.Background(), id, parentId).ClientId(clientId).EmailConnectorParsingStyle(emailConnectorParsingStyle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailConnectorParsingStylesAPI.PutSystemEmailConnectorsByParentIdParsingStylesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemEmailConnectorsByParentIdParsingStylesById`: EmailConnectorParsingStyle
	fmt.Fprintf(os.Stdout, "Response from `EmailConnectorParsingStylesAPI.PutSystemEmailConnectorsByParentIdParsingStylesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | parsingStyleId | 
**parentId** | **int32** | emailConnectorId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemEmailConnectorsByParentIdParsingStylesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **emailConnectorParsingStyle** | [**EmailConnectorParsingStyle**](EmailConnectorParsingStyle.md) | emailConnectorParsingStyle | 

### Return type

[**EmailConnectorParsingStyle**](EmailConnectorParsingStyle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

