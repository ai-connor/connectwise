# \OrderStatusesEmailTemplateAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSalesOrdersStatusesByParentIdEmailtemplatesById**](OrderStatusesEmailTemplateAPI.md#DeleteSalesOrdersStatusesByParentIdEmailtemplatesById) | **Delete** /sales/orders/statuses/{parentId}/emailtemplates/{id} | Delete OrderStatusEmailTemplate
[**GetSalesOrdersStatusesByParentIdEmailtemplates**](OrderStatusesEmailTemplateAPI.md#GetSalesOrdersStatusesByParentIdEmailtemplates) | **Get** /sales/orders/statuses/{parentId}/emailtemplates/ | Get List of OrderStatusEmailTemplate
[**GetSalesOrdersStatusesByParentIdEmailtemplatesById**](OrderStatusesEmailTemplateAPI.md#GetSalesOrdersStatusesByParentIdEmailtemplatesById) | **Get** /sales/orders/statuses/{parentId}/emailtemplates/{id} | Get OrderStatusEmailTemplate
[**GetSalesOrdersStatusesByParentIdEmailtemplatesCount**](OrderStatusesEmailTemplateAPI.md#GetSalesOrdersStatusesByParentIdEmailtemplatesCount) | **Get** /sales/orders/statuses/{parentId}/emailtemplates/count | Get Count of OrderStatusEmailTemplate
[**PatchSalesOrdersStatusesByParentIdEmailtemplatesById**](OrderStatusesEmailTemplateAPI.md#PatchSalesOrdersStatusesByParentIdEmailtemplatesById) | **Patch** /sales/orders/statuses/{parentId}/emailtemplates/{id} | Patch OrderStatusEmailTemplate
[**PostSalesOrdersStatusesByParentIdEmailtemplates**](OrderStatusesEmailTemplateAPI.md#PostSalesOrdersStatusesByParentIdEmailtemplates) | **Post** /sales/orders/statuses/{parentId}/emailtemplates/ | Post OrderStatusEmailTemplate
[**PutSalesOrdersStatusesByParentIdEmailtemplatesById**](OrderStatusesEmailTemplateAPI.md#PutSalesOrdersStatusesByParentIdEmailtemplatesById) | **Put** /sales/orders/statuses/{parentId}/emailtemplates/{id} | Put OrderStatusEmailTemplate



## DeleteSalesOrdersStatusesByParentIdEmailtemplatesById

> DeleteSalesOrdersStatusesByParentIdEmailtemplatesById(ctx, id, parentId).ClientId(clientId).Execute()

Delete OrderStatusEmailTemplate

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
	id := int32(56) // int32 | emailtemplateId
	parentId := int32(56) // int32 | statusId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrderStatusesEmailTemplateAPI.DeleteSalesOrdersStatusesByParentIdEmailtemplatesById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderStatusesEmailTemplateAPI.DeleteSalesOrdersStatusesByParentIdEmailtemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailtemplateId | 
**parentId** | **int32** | statusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSalesOrdersStatusesByParentIdEmailtemplatesByIdRequest struct via the builder pattern


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


## GetSalesOrdersStatusesByParentIdEmailtemplates

> []OrderStatusEmailTemplate GetSalesOrdersStatusesByParentIdEmailtemplates(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of OrderStatusEmailTemplate

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
	parentId := int32(56) // int32 | statusId
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
	resp, r, err := apiClient.OrderStatusesEmailTemplateAPI.GetSalesOrdersStatusesByParentIdEmailtemplates(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderStatusesEmailTemplateAPI.GetSalesOrdersStatusesByParentIdEmailtemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesOrdersStatusesByParentIdEmailtemplates`: []OrderStatusEmailTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrderStatusesEmailTemplateAPI.GetSalesOrdersStatusesByParentIdEmailtemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | statusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesOrdersStatusesByParentIdEmailtemplatesRequest struct via the builder pattern


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

[**[]OrderStatusEmailTemplate**](OrderStatusEmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalesOrdersStatusesByParentIdEmailtemplatesById

> OrderStatusEmailTemplate GetSalesOrdersStatusesByParentIdEmailtemplatesById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get OrderStatusEmailTemplate

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
	id := int32(56) // int32 | emailtemplateId
	parentId := int32(56) // int32 | statusId
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
	resp, r, err := apiClient.OrderStatusesEmailTemplateAPI.GetSalesOrdersStatusesByParentIdEmailtemplatesById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderStatusesEmailTemplateAPI.GetSalesOrdersStatusesByParentIdEmailtemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesOrdersStatusesByParentIdEmailtemplatesById`: OrderStatusEmailTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrderStatusesEmailTemplateAPI.GetSalesOrdersStatusesByParentIdEmailtemplatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailtemplateId | 
**parentId** | **int32** | statusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesOrdersStatusesByParentIdEmailtemplatesByIdRequest struct via the builder pattern


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

[**OrderStatusEmailTemplate**](OrderStatusEmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalesOrdersStatusesByParentIdEmailtemplatesCount

> Count GetSalesOrdersStatusesByParentIdEmailtemplatesCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of OrderStatusEmailTemplate

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
	parentId := int32(56) // int32 | statusId
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
	resp, r, err := apiClient.OrderStatusesEmailTemplateAPI.GetSalesOrdersStatusesByParentIdEmailtemplatesCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderStatusesEmailTemplateAPI.GetSalesOrdersStatusesByParentIdEmailtemplatesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesOrdersStatusesByParentIdEmailtemplatesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `OrderStatusesEmailTemplateAPI.GetSalesOrdersStatusesByParentIdEmailtemplatesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | statusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesOrdersStatusesByParentIdEmailtemplatesCountRequest struct via the builder pattern


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


## PatchSalesOrdersStatusesByParentIdEmailtemplatesById

> OrderStatusEmailTemplate PatchSalesOrdersStatusesByParentIdEmailtemplatesById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch OrderStatusEmailTemplate

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
	id := int32(56) // int32 | emailtemplateId
	parentId := int32(56) // int32 | statusId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderStatusesEmailTemplateAPI.PatchSalesOrdersStatusesByParentIdEmailtemplatesById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderStatusesEmailTemplateAPI.PatchSalesOrdersStatusesByParentIdEmailtemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSalesOrdersStatusesByParentIdEmailtemplatesById`: OrderStatusEmailTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrderStatusesEmailTemplateAPI.PatchSalesOrdersStatusesByParentIdEmailtemplatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailtemplateId | 
**parentId** | **int32** | statusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSalesOrdersStatusesByParentIdEmailtemplatesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**OrderStatusEmailTemplate**](OrderStatusEmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSalesOrdersStatusesByParentIdEmailtemplates

> OrderStatusEmailTemplate PostSalesOrdersStatusesByParentIdEmailtemplates(ctx, parentId).ClientId(clientId).OrderStatusEmailTemplate(orderStatusEmailTemplate).Execute()

Post OrderStatusEmailTemplate

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
	parentId := int32(56) // int32 | statusId
	clientId := "clientId_example" // string | 
	orderStatusEmailTemplate := *openapiclient.NewOrderStatusEmailTemplate("Subject_example", "Body_example") // OrderStatusEmailTemplate | orderStatusEmailTemplate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderStatusesEmailTemplateAPI.PostSalesOrdersStatusesByParentIdEmailtemplates(context.Background(), parentId).ClientId(clientId).OrderStatusEmailTemplate(orderStatusEmailTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderStatusesEmailTemplateAPI.PostSalesOrdersStatusesByParentIdEmailtemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSalesOrdersStatusesByParentIdEmailtemplates`: OrderStatusEmailTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrderStatusesEmailTemplateAPI.PostSalesOrdersStatusesByParentIdEmailtemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | statusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSalesOrdersStatusesByParentIdEmailtemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **orderStatusEmailTemplate** | [**OrderStatusEmailTemplate**](OrderStatusEmailTemplate.md) | orderStatusEmailTemplate | 

### Return type

[**OrderStatusEmailTemplate**](OrderStatusEmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSalesOrdersStatusesByParentIdEmailtemplatesById

> OrderStatusEmailTemplate PutSalesOrdersStatusesByParentIdEmailtemplatesById(ctx, id, parentId).ClientId(clientId).OrderStatusEmailTemplate(orderStatusEmailTemplate).Execute()

Put OrderStatusEmailTemplate

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
	id := int32(56) // int32 | emailtemplateId
	parentId := int32(56) // int32 | statusId
	clientId := "clientId_example" // string | 
	orderStatusEmailTemplate := *openapiclient.NewOrderStatusEmailTemplate("Subject_example", "Body_example") // OrderStatusEmailTemplate | orderStatusEmailTemplate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderStatusesEmailTemplateAPI.PutSalesOrdersStatusesByParentIdEmailtemplatesById(context.Background(), id, parentId).ClientId(clientId).OrderStatusEmailTemplate(orderStatusEmailTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderStatusesEmailTemplateAPI.PutSalesOrdersStatusesByParentIdEmailtemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSalesOrdersStatusesByParentIdEmailtemplatesById`: OrderStatusEmailTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrderStatusesEmailTemplateAPI.PutSalesOrdersStatusesByParentIdEmailtemplatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailtemplateId | 
**parentId** | **int32** | statusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSalesOrdersStatusesByParentIdEmailtemplatesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **orderStatusEmailTemplate** | [**OrderStatusEmailTemplate**](OrderStatusEmailTemplate.md) | orderStatusEmailTemplate | 

### Return type

[**OrderStatusEmailTemplate**](OrderStatusEmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

