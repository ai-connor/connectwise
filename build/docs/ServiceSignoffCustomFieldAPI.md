# \ServiceSignoffCustomFieldAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServiceServiceSignoffByParentIdSignoffcustomfieldsById**](ServiceSignoffCustomFieldAPI.md#DeleteServiceServiceSignoffByParentIdSignoffcustomfieldsById) | **Delete** /service/serviceSignoff/{parentId}/signoffcustomfields/{id} | Delete ServiceSignoffCustomField
[**GetServiceServiceSignoffByParentIdSignoffcustomfields**](ServiceSignoffCustomFieldAPI.md#GetServiceServiceSignoffByParentIdSignoffcustomfields) | **Get** /service/serviceSignoff/{parentId}/signoffcustomfields | Get List of ServiceSignoffCustomField
[**GetServiceServiceSignoffByParentIdSignoffcustomfieldsById**](ServiceSignoffCustomFieldAPI.md#GetServiceServiceSignoffByParentIdSignoffcustomfieldsById) | **Get** /service/serviceSignoff/{parentId}/signoffcustomfields/{id} | Get ServiceSignoffCustomField
[**GetServiceServiceSignoffByParentIdSignoffcustomfieldsCount**](ServiceSignoffCustomFieldAPI.md#GetServiceServiceSignoffByParentIdSignoffcustomfieldsCount) | **Get** /service/serviceSignoff/{parentId}/signoffcustomfields/count | Get Count of ServiceSignoffCustomField
[**PatchServiceServiceSignoffByParentIdSignoffcustomfieldsById**](ServiceSignoffCustomFieldAPI.md#PatchServiceServiceSignoffByParentIdSignoffcustomfieldsById) | **Patch** /service/serviceSignoff/{parentId}/signoffcustomfields/{id} | Patch ServiceSignoffCustomField
[**PostServiceServiceSignoffByParentIdSignoffcustomfields**](ServiceSignoffCustomFieldAPI.md#PostServiceServiceSignoffByParentIdSignoffcustomfields) | **Post** /service/serviceSignoff/{parentId}/signoffcustomfields | Post ServiceSignoffCustomField
[**PutServiceServiceSignoffByParentIdSignoffcustomfieldsById**](ServiceSignoffCustomFieldAPI.md#PutServiceServiceSignoffByParentIdSignoffcustomfieldsById) | **Put** /service/serviceSignoff/{parentId}/signoffcustomfields/{id} | Put ServiceSignoffCustomField



## DeleteServiceServiceSignoffByParentIdSignoffcustomfieldsById

> DeleteServiceServiceSignoffByParentIdSignoffcustomfieldsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ServiceSignoffCustomField

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
	id := int32(56) // int32 | ServiceSignoffCustomFieldId
	parentId := int32(56) // int32 | serviceSignoffId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceSignoffCustomFieldAPI.DeleteServiceServiceSignoffByParentIdSignoffcustomfieldsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSignoffCustomFieldAPI.DeleteServiceServiceSignoffByParentIdSignoffcustomfieldsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ServiceSignoffCustomFieldId | 
**parentId** | **int32** | serviceSignoffId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceServiceSignoffByParentIdSignoffcustomfieldsByIdRequest struct via the builder pattern


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


## GetServiceServiceSignoffByParentIdSignoffcustomfields

> []ServiceSignoffCustomField GetServiceServiceSignoffByParentIdSignoffcustomfields(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ServiceSignoffCustomField

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
	parentId := int32(56) // int32 | serviceSignoffId
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
	resp, r, err := apiClient.ServiceSignoffCustomFieldAPI.GetServiceServiceSignoffByParentIdSignoffcustomfields(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSignoffCustomFieldAPI.GetServiceServiceSignoffByParentIdSignoffcustomfields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceServiceSignoffByParentIdSignoffcustomfields`: []ServiceSignoffCustomField
	fmt.Fprintf(os.Stdout, "Response from `ServiceSignoffCustomFieldAPI.GetServiceServiceSignoffByParentIdSignoffcustomfields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | serviceSignoffId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceServiceSignoffByParentIdSignoffcustomfieldsRequest struct via the builder pattern


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

[**[]ServiceSignoffCustomField**](ServiceSignoffCustomField.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceServiceSignoffByParentIdSignoffcustomfieldsById

> ServiceSignoffCustomField GetServiceServiceSignoffByParentIdSignoffcustomfieldsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ServiceSignoffCustomField

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
	id := int32(56) // int32 | ServiceSignoffCustomFieldId
	parentId := int32(56) // int32 | serviceSignoffId
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
	resp, r, err := apiClient.ServiceSignoffCustomFieldAPI.GetServiceServiceSignoffByParentIdSignoffcustomfieldsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSignoffCustomFieldAPI.GetServiceServiceSignoffByParentIdSignoffcustomfieldsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceServiceSignoffByParentIdSignoffcustomfieldsById`: ServiceSignoffCustomField
	fmt.Fprintf(os.Stdout, "Response from `ServiceSignoffCustomFieldAPI.GetServiceServiceSignoffByParentIdSignoffcustomfieldsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ServiceSignoffCustomFieldId | 
**parentId** | **int32** | serviceSignoffId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceServiceSignoffByParentIdSignoffcustomfieldsByIdRequest struct via the builder pattern


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

[**ServiceSignoffCustomField**](ServiceSignoffCustomField.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceServiceSignoffByParentIdSignoffcustomfieldsCount

> Count GetServiceServiceSignoffByParentIdSignoffcustomfieldsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ServiceSignoffCustomField

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
	parentId := int32(56) // int32 | serviceSignoffId
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
	resp, r, err := apiClient.ServiceSignoffCustomFieldAPI.GetServiceServiceSignoffByParentIdSignoffcustomfieldsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSignoffCustomFieldAPI.GetServiceServiceSignoffByParentIdSignoffcustomfieldsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceServiceSignoffByParentIdSignoffcustomfieldsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ServiceSignoffCustomFieldAPI.GetServiceServiceSignoffByParentIdSignoffcustomfieldsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | serviceSignoffId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceServiceSignoffByParentIdSignoffcustomfieldsCountRequest struct via the builder pattern


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


## PatchServiceServiceSignoffByParentIdSignoffcustomfieldsById

> ServiceSignoffCustomField PatchServiceServiceSignoffByParentIdSignoffcustomfieldsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ServiceSignoffCustomField

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
	id := int32(56) // int32 | ServiceSignoffCustomFieldId
	parentId := int32(56) // int32 | serviceSignoffId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSignoffCustomFieldAPI.PatchServiceServiceSignoffByParentIdSignoffcustomfieldsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSignoffCustomFieldAPI.PatchServiceServiceSignoffByParentIdSignoffcustomfieldsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchServiceServiceSignoffByParentIdSignoffcustomfieldsById`: ServiceSignoffCustomField
	fmt.Fprintf(os.Stdout, "Response from `ServiceSignoffCustomFieldAPI.PatchServiceServiceSignoffByParentIdSignoffcustomfieldsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ServiceSignoffCustomFieldId | 
**parentId** | **int32** | serviceSignoffId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchServiceServiceSignoffByParentIdSignoffcustomfieldsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ServiceSignoffCustomField**](ServiceSignoffCustomField.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServiceServiceSignoffByParentIdSignoffcustomfields

> ServiceSignoffCustomField PostServiceServiceSignoffByParentIdSignoffcustomfields(ctx, parentId).ClientId(clientId).ServiceSignoffCustomField(serviceSignoffCustomField).Execute()

Post ServiceSignoffCustomField

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
	parentId := int32(56) // int32 | serviceSignoffId
	clientId := "clientId_example" // string | 
	serviceSignoffCustomField := *openapiclient.NewServiceSignoffCustomField(NullableFloat64(123), "DisplaySection_example", *openapiclient.NewUserDefinedFieldReference()) // ServiceSignoffCustomField | serviceSignoff

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSignoffCustomFieldAPI.PostServiceServiceSignoffByParentIdSignoffcustomfields(context.Background(), parentId).ClientId(clientId).ServiceSignoffCustomField(serviceSignoffCustomField).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSignoffCustomFieldAPI.PostServiceServiceSignoffByParentIdSignoffcustomfields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServiceServiceSignoffByParentIdSignoffcustomfields`: ServiceSignoffCustomField
	fmt.Fprintf(os.Stdout, "Response from `ServiceSignoffCustomFieldAPI.PostServiceServiceSignoffByParentIdSignoffcustomfields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | serviceSignoffId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServiceServiceSignoffByParentIdSignoffcustomfieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **serviceSignoffCustomField** | [**ServiceSignoffCustomField**](ServiceSignoffCustomField.md) | serviceSignoff | 

### Return type

[**ServiceSignoffCustomField**](ServiceSignoffCustomField.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServiceServiceSignoffByParentIdSignoffcustomfieldsById

> ServiceSignoffCustomField PutServiceServiceSignoffByParentIdSignoffcustomfieldsById(ctx, id, parentId).ClientId(clientId).ServiceSignoffCustomField(serviceSignoffCustomField).Execute()

Put ServiceSignoffCustomField

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
	id := int32(56) // int32 | ServiceSignoffCustomFieldId
	parentId := int32(56) // int32 | serviceSignoffId
	clientId := "clientId_example" // string | 
	serviceSignoffCustomField := *openapiclient.NewServiceSignoffCustomField(NullableFloat64(123), "DisplaySection_example", *openapiclient.NewUserDefinedFieldReference()) // ServiceSignoffCustomField | ServiceSignoffCustomField

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSignoffCustomFieldAPI.PutServiceServiceSignoffByParentIdSignoffcustomfieldsById(context.Background(), id, parentId).ClientId(clientId).ServiceSignoffCustomField(serviceSignoffCustomField).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSignoffCustomFieldAPI.PutServiceServiceSignoffByParentIdSignoffcustomfieldsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServiceServiceSignoffByParentIdSignoffcustomfieldsById`: ServiceSignoffCustomField
	fmt.Fprintf(os.Stdout, "Response from `ServiceSignoffCustomFieldAPI.PutServiceServiceSignoffByParentIdSignoffcustomfieldsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ServiceSignoffCustomFieldId | 
**parentId** | **int32** | serviceSignoffId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServiceServiceSignoffByParentIdSignoffcustomfieldsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **serviceSignoffCustomField** | [**ServiceSignoffCustomField**](ServiceSignoffCustomField.md) | ServiceSignoffCustomField | 

### Return type

[**ServiceSignoffCustomField**](ServiceSignoffCustomField.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

