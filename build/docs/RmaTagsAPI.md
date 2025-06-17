# \RmaTagsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProcurementRmaTagsById**](RmaTagsAPI.md#DeleteProcurementRmaTagsById) | **Delete** /procurement/rmaTags/{id} | Delete RmaTag
[**GetProcurementRmaTags**](RmaTagsAPI.md#GetProcurementRmaTags) | **Get** /procurement/rmaTags | Get List of RmaTag
[**GetProcurementRmaTagsById**](RmaTagsAPI.md#GetProcurementRmaTagsById) | **Get** /procurement/rmaTags/{id} | Get RmaTag
[**GetProcurementRmaTagsCount**](RmaTagsAPI.md#GetProcurementRmaTagsCount) | **Get** /procurement/rmaTags/count | Get Count of RmaTag
[**GetProcurementRmaTagsDefault**](RmaTagsAPI.md#GetProcurementRmaTagsDefault) | **Get** /procurement/rmaTags/default | Get RmaTag
[**PatchProcurementRmaTagsById**](RmaTagsAPI.md#PatchProcurementRmaTagsById) | **Patch** /procurement/rmaTags/{id} | Patch RmaTag
[**PostProcurementRmaTags**](RmaTagsAPI.md#PostProcurementRmaTags) | **Post** /procurement/rmaTags | Post RmaTag
[**PutProcurementRmaTagsById**](RmaTagsAPI.md#PutProcurementRmaTagsById) | **Put** /procurement/rmaTags/{id} | Put RmaTag



## DeleteProcurementRmaTagsById

> DeleteProcurementRmaTagsById(ctx, id).ClientId(clientId).Execute()

Delete RmaTag

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
	id := int32(56) // int32 | rmaTagId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RmaTagsAPI.DeleteProcurementRmaTagsById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaTagsAPI.DeleteProcurementRmaTagsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | rmaTagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcurementRmaTagsByIdRequest struct via the builder pattern


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


## GetProcurementRmaTags

> []RmaTag GetProcurementRmaTags(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of RmaTag

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
	resp, r, err := apiClient.RmaTagsAPI.GetProcurementRmaTags(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaTagsAPI.GetProcurementRmaTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementRmaTags`: []RmaTag
	fmt.Fprintf(os.Stdout, "Response from `RmaTagsAPI.GetProcurementRmaTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementRmaTagsRequest struct via the builder pattern


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

[**[]RmaTag**](RmaTag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementRmaTagsById

> RmaTag GetProcurementRmaTagsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get RmaTag

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
	id := int32(56) // int32 | rmaTagId
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
	resp, r, err := apiClient.RmaTagsAPI.GetProcurementRmaTagsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaTagsAPI.GetProcurementRmaTagsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementRmaTagsById`: RmaTag
	fmt.Fprintf(os.Stdout, "Response from `RmaTagsAPI.GetProcurementRmaTagsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | rmaTagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementRmaTagsByIdRequest struct via the builder pattern


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

[**RmaTag**](RmaTag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementRmaTagsCount

> Count GetProcurementRmaTagsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of RmaTag

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
	resp, r, err := apiClient.RmaTagsAPI.GetProcurementRmaTagsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaTagsAPI.GetProcurementRmaTagsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementRmaTagsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `RmaTagsAPI.GetProcurementRmaTagsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementRmaTagsCountRequest struct via the builder pattern


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


## GetProcurementRmaTagsDefault

> RmaTag GetProcurementRmaTagsDefault(ctx, productId, billingLogId, ticketId, projectId, salesOrderId, companyId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get RmaTag

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
	productId := int32(56) // int32 | productId
	billingLogId := int32(56) // int32 | billingLogId
	ticketId := int32(56) // int32 | ticketId
	projectId := int32(56) // int32 | projectId
	salesOrderId := int32(56) // int32 | salesOrderId
	companyId := int32(56) // int32 | companyId
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
	resp, r, err := apiClient.RmaTagsAPI.GetProcurementRmaTagsDefault(context.Background(), productId, billingLogId, ticketId, projectId, salesOrderId, companyId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaTagsAPI.GetProcurementRmaTagsDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementRmaTagsDefault`: RmaTag
	fmt.Fprintf(os.Stdout, "Response from `RmaTagsAPI.GetProcurementRmaTagsDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **int32** | productId | 
**billingLogId** | **int32** | billingLogId | 
**ticketId** | **int32** | ticketId | 
**projectId** | **int32** | projectId | 
**salesOrderId** | **int32** | salesOrderId | 
**companyId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementRmaTagsDefaultRequest struct via the builder pattern


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

[**RmaTag**](RmaTag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchProcurementRmaTagsById

> RmaTag PatchProcurementRmaTagsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch RmaTag

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
	id := int32(56) // int32 | rmaTagId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RmaTagsAPI.PatchProcurementRmaTagsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaTagsAPI.PatchProcurementRmaTagsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProcurementRmaTagsById`: RmaTag
	fmt.Fprintf(os.Stdout, "Response from `RmaTagsAPI.PatchProcurementRmaTagsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | rmaTagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProcurementRmaTagsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**RmaTag**](RmaTag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProcurementRmaTags

> RmaTag PostProcurementRmaTags(ctx).ClientId(clientId).RmaTag(rmaTag).Execute()

Post RmaTag

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
	rmaTag := *openapiclient.NewRmaTag(*openapiclient.NewIvItemReference(), "ProductDescription_example", *openapiclient.NewRmaStatusReference(), *openapiclient.NewSystemLocationReference(), *openapiclient.NewSystemDepartmentReference(), *openapiclient.NewCompanyReference(), *openapiclient.NewRmaDispositionReference()) // RmaTag | rmaTag

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RmaTagsAPI.PostProcurementRmaTags(context.Background()).ClientId(clientId).RmaTag(rmaTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaTagsAPI.PostProcurementRmaTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementRmaTags`: RmaTag
	fmt.Fprintf(os.Stdout, "Response from `RmaTagsAPI.PostProcurementRmaTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementRmaTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **rmaTag** | [**RmaTag**](RmaTag.md) | rmaTag | 

### Return type

[**RmaTag**](RmaTag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProcurementRmaTagsById

> RmaTag PutProcurementRmaTagsById(ctx, id).ClientId(clientId).RmaTag(rmaTag).Execute()

Put RmaTag

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
	id := int32(56) // int32 | rmaTagId
	clientId := "clientId_example" // string | 
	rmaTag := *openapiclient.NewRmaTag(*openapiclient.NewIvItemReference(), "ProductDescription_example", *openapiclient.NewRmaStatusReference(), *openapiclient.NewSystemLocationReference(), *openapiclient.NewSystemDepartmentReference(), *openapiclient.NewCompanyReference(), *openapiclient.NewRmaDispositionReference()) // RmaTag | rmaTag

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RmaTagsAPI.PutProcurementRmaTagsById(context.Background(), id).ClientId(clientId).RmaTag(rmaTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaTagsAPI.PutProcurementRmaTagsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProcurementRmaTagsById`: RmaTag
	fmt.Fprintf(os.Stdout, "Response from `RmaTagsAPI.PutProcurementRmaTagsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | rmaTagId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProcurementRmaTagsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **rmaTag** | [**RmaTag**](RmaTag.md) | rmaTag | 

### Return type

[**RmaTag**](RmaTag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

