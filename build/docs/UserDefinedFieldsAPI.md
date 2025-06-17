# \UserDefinedFieldsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemUserDefinedFieldsById**](UserDefinedFieldsAPI.md#DeleteSystemUserDefinedFieldsById) | **Delete** /system/userDefinedFields/{id} | Delete UserDefinedField
[**GetSystemUserDefinedFields**](UserDefinedFieldsAPI.md#GetSystemUserDefinedFields) | **Get** /system/userDefinedFields | Get List of UserDefinedField
[**GetSystemUserDefinedFieldsById**](UserDefinedFieldsAPI.md#GetSystemUserDefinedFieldsById) | **Get** /system/userDefinedFields/{id} | Get UserDefinedField
[**GetSystemUserDefinedFieldsCount**](UserDefinedFieldsAPI.md#GetSystemUserDefinedFieldsCount) | **Get** /system/userDefinedFields/count | Get Count of UserDefinedField
[**PatchSystemUserDefinedFieldsById**](UserDefinedFieldsAPI.md#PatchSystemUserDefinedFieldsById) | **Patch** /system/userDefinedFields/{id} | Patch UserDefinedField
[**PostSystemUserDefinedFields**](UserDefinedFieldsAPI.md#PostSystemUserDefinedFields) | **Post** /system/userDefinedFields | Post UserDefinedField
[**PutSystemUserDefinedFieldsById**](UserDefinedFieldsAPI.md#PutSystemUserDefinedFieldsById) | **Put** /system/userDefinedFields/{id} | Put UserDefinedField



## DeleteSystemUserDefinedFieldsById

> DeleteSystemUserDefinedFieldsById(ctx, id).ClientId(clientId).Execute()

Delete UserDefinedField

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
	id := int32(56) // int32 | userDefinedFieldId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserDefinedFieldsAPI.DeleteSystemUserDefinedFieldsById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDefinedFieldsAPI.DeleteSystemUserDefinedFieldsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | userDefinedFieldId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemUserDefinedFieldsByIdRequest struct via the builder pattern


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


## GetSystemUserDefinedFields

> []UserDefinedField GetSystemUserDefinedFields(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of UserDefinedField

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
	resp, r, err := apiClient.UserDefinedFieldsAPI.GetSystemUserDefinedFields(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDefinedFieldsAPI.GetSystemUserDefinedFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemUserDefinedFields`: []UserDefinedField
	fmt.Fprintf(os.Stdout, "Response from `UserDefinedFieldsAPI.GetSystemUserDefinedFields`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemUserDefinedFieldsRequest struct via the builder pattern


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

[**[]UserDefinedField**](UserDefinedField.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemUserDefinedFieldsById

> UserDefinedField GetSystemUserDefinedFieldsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get UserDefinedField

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
	id := int32(56) // int32 | userDefinedFieldId
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
	resp, r, err := apiClient.UserDefinedFieldsAPI.GetSystemUserDefinedFieldsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDefinedFieldsAPI.GetSystemUserDefinedFieldsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemUserDefinedFieldsById`: UserDefinedField
	fmt.Fprintf(os.Stdout, "Response from `UserDefinedFieldsAPI.GetSystemUserDefinedFieldsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | userDefinedFieldId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemUserDefinedFieldsByIdRequest struct via the builder pattern


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

[**UserDefinedField**](UserDefinedField.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemUserDefinedFieldsCount

> Count GetSystemUserDefinedFieldsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of UserDefinedField

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
	resp, r, err := apiClient.UserDefinedFieldsAPI.GetSystemUserDefinedFieldsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDefinedFieldsAPI.GetSystemUserDefinedFieldsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemUserDefinedFieldsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `UserDefinedFieldsAPI.GetSystemUserDefinedFieldsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemUserDefinedFieldsCountRequest struct via the builder pattern


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


## PatchSystemUserDefinedFieldsById

> UserDefinedField PatchSystemUserDefinedFieldsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch UserDefinedField

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
	id := int32(56) // int32 | userDefinedFieldId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserDefinedFieldsAPI.PatchSystemUserDefinedFieldsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDefinedFieldsAPI.PatchSystemUserDefinedFieldsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemUserDefinedFieldsById`: UserDefinedField
	fmt.Fprintf(os.Stdout, "Response from `UserDefinedFieldsAPI.PatchSystemUserDefinedFieldsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | userDefinedFieldId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemUserDefinedFieldsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**UserDefinedField**](UserDefinedField.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemUserDefinedFields

> UserDefinedField PostSystemUserDefinedFields(ctx).ClientId(clientId).UserDefinedField(userDefinedField).Execute()

Post UserDefinedField

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
	userDefinedField := *openapiclient.NewUserDefinedField(NullableInt32(123), "Caption_example", NullableInt32(123), "FieldTypeIdentifier_example") // UserDefinedField | userDefinedField

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserDefinedFieldsAPI.PostSystemUserDefinedFields(context.Background()).ClientId(clientId).UserDefinedField(userDefinedField).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDefinedFieldsAPI.PostSystemUserDefinedFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemUserDefinedFields`: UserDefinedField
	fmt.Fprintf(os.Stdout, "Response from `UserDefinedFieldsAPI.PostSystemUserDefinedFields`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemUserDefinedFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **userDefinedField** | [**UserDefinedField**](UserDefinedField.md) | userDefinedField | 

### Return type

[**UserDefinedField**](UserDefinedField.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemUserDefinedFieldsById

> UserDefinedField PutSystemUserDefinedFieldsById(ctx, id).ClientId(clientId).UserDefinedField(userDefinedField).Execute()

Put UserDefinedField

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
	id := int32(56) // int32 | userDefinedFieldId
	clientId := "clientId_example" // string | 
	userDefinedField := *openapiclient.NewUserDefinedField(NullableInt32(123), "Caption_example", NullableInt32(123), "FieldTypeIdentifier_example") // UserDefinedField | userDefinedField

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserDefinedFieldsAPI.PutSystemUserDefinedFieldsById(context.Background(), id).ClientId(clientId).UserDefinedField(userDefinedField).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDefinedFieldsAPI.PutSystemUserDefinedFieldsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemUserDefinedFieldsById`: UserDefinedField
	fmt.Fprintf(os.Stdout, "Response from `UserDefinedFieldsAPI.PutSystemUserDefinedFieldsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | userDefinedFieldId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemUserDefinedFieldsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **userDefinedField** | [**UserDefinedField**](UserDefinedField.md) | userDefinedField | 

### Return type

[**UserDefinedField**](UserDefinedField.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

