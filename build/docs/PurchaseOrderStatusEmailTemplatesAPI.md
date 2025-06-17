# \PurchaseOrderStatusEmailTemplatesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProcurementPurchaseorderstatusesByParentIdEmailtemplatesById**](PurchaseOrderStatusEmailTemplatesAPI.md#DeleteProcurementPurchaseorderstatusesByParentIdEmailtemplatesById) | **Delete** /procurement/purchaseorderstatuses/{parentId}/emailtemplates/{id} | Delete PurchaseOrderStatusEmailTemplate
[**GetProcurementPurchaseorderstatusesByParentIdEmailtemplates**](PurchaseOrderStatusEmailTemplatesAPI.md#GetProcurementPurchaseorderstatusesByParentIdEmailtemplates) | **Get** /procurement/purchaseorderstatuses/{parentId}/emailtemplates/ | Get List of PurchaseOrderStatusEmailTemplate
[**GetProcurementPurchaseorderstatusesByParentIdEmailtemplatesById**](PurchaseOrderStatusEmailTemplatesAPI.md#GetProcurementPurchaseorderstatusesByParentIdEmailtemplatesById) | **Get** /procurement/purchaseorderstatuses/{parentId}/emailtemplates/{id} | Get PurchaseOrderStatusEmailTemplate
[**GetProcurementPurchaseorderstatusesByParentIdEmailtemplatesCount**](PurchaseOrderStatusEmailTemplatesAPI.md#GetProcurementPurchaseorderstatusesByParentIdEmailtemplatesCount) | **Get** /procurement/purchaseorderstatuses/{parentId}/emailtemplates/count | Get Count of PurchaseOrderStatusEmailTemplate
[**PatchProcurementPurchaseorderstatusesByParentIdEmailtemplatesById**](PurchaseOrderStatusEmailTemplatesAPI.md#PatchProcurementPurchaseorderstatusesByParentIdEmailtemplatesById) | **Patch** /procurement/purchaseorderstatuses/{parentId}/emailtemplates/{id} | Patch PurchaseOrderStatusEmailTemplate
[**PostProcurementPurchaseorderstatusesByParentIdEmailtemplates**](PurchaseOrderStatusEmailTemplatesAPI.md#PostProcurementPurchaseorderstatusesByParentIdEmailtemplates) | **Post** /procurement/purchaseorderstatuses/{parentId}/emailtemplates/ | Post PurchaseOrderStatusEmailTemplate
[**PutProcurementPurchaseorderstatusesByParentIdEmailtemplatesById**](PurchaseOrderStatusEmailTemplatesAPI.md#PutProcurementPurchaseorderstatusesByParentIdEmailtemplatesById) | **Put** /procurement/purchaseorderstatuses/{parentId}/emailtemplates/{id} | Put PurchaseOrderStatusEmailTemplate



## DeleteProcurementPurchaseorderstatusesByParentIdEmailtemplatesById

> DeleteProcurementPurchaseorderstatusesByParentIdEmailtemplatesById(ctx, id, parentId).ClientId(clientId).Execute()

Delete PurchaseOrderStatusEmailTemplate

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
	parentId := int32(56) // int32 | purchaseorderstatusId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PurchaseOrderStatusEmailTemplatesAPI.DeleteProcurementPurchaseorderstatusesByParentIdEmailtemplatesById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderStatusEmailTemplatesAPI.DeleteProcurementPurchaseorderstatusesByParentIdEmailtemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailtemplateId | 
**parentId** | **int32** | purchaseorderstatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcurementPurchaseorderstatusesByParentIdEmailtemplatesByIdRequest struct via the builder pattern


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


## GetProcurementPurchaseorderstatusesByParentIdEmailtemplates

> []PurchaseOrderStatusEmailTemplate GetProcurementPurchaseorderstatusesByParentIdEmailtemplates(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of PurchaseOrderStatusEmailTemplate

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
	parentId := int32(56) // int32 | purchaseorderstatusId
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
	resp, r, err := apiClient.PurchaseOrderStatusEmailTemplatesAPI.GetProcurementPurchaseorderstatusesByParentIdEmailtemplates(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderStatusEmailTemplatesAPI.GetProcurementPurchaseorderstatusesByParentIdEmailtemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementPurchaseorderstatusesByParentIdEmailtemplates`: []PurchaseOrderStatusEmailTemplate
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderStatusEmailTemplatesAPI.GetProcurementPurchaseorderstatusesByParentIdEmailtemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | purchaseorderstatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementPurchaseorderstatusesByParentIdEmailtemplatesRequest struct via the builder pattern


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

[**[]PurchaseOrderStatusEmailTemplate**](PurchaseOrderStatusEmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementPurchaseorderstatusesByParentIdEmailtemplatesById

> PurchaseOrderStatusEmailTemplate GetProcurementPurchaseorderstatusesByParentIdEmailtemplatesById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get PurchaseOrderStatusEmailTemplate

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
	parentId := int32(56) // int32 | purchaseorderstatusId
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
	resp, r, err := apiClient.PurchaseOrderStatusEmailTemplatesAPI.GetProcurementPurchaseorderstatusesByParentIdEmailtemplatesById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderStatusEmailTemplatesAPI.GetProcurementPurchaseorderstatusesByParentIdEmailtemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementPurchaseorderstatusesByParentIdEmailtemplatesById`: PurchaseOrderStatusEmailTemplate
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderStatusEmailTemplatesAPI.GetProcurementPurchaseorderstatusesByParentIdEmailtemplatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailtemplateId | 
**parentId** | **int32** | purchaseorderstatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementPurchaseorderstatusesByParentIdEmailtemplatesByIdRequest struct via the builder pattern


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

[**PurchaseOrderStatusEmailTemplate**](PurchaseOrderStatusEmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementPurchaseorderstatusesByParentIdEmailtemplatesCount

> Count GetProcurementPurchaseorderstatusesByParentIdEmailtemplatesCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of PurchaseOrderStatusEmailTemplate

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
	parentId := int32(56) // int32 | purchaseorderstatusId
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
	resp, r, err := apiClient.PurchaseOrderStatusEmailTemplatesAPI.GetProcurementPurchaseorderstatusesByParentIdEmailtemplatesCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderStatusEmailTemplatesAPI.GetProcurementPurchaseorderstatusesByParentIdEmailtemplatesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementPurchaseorderstatusesByParentIdEmailtemplatesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderStatusEmailTemplatesAPI.GetProcurementPurchaseorderstatusesByParentIdEmailtemplatesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | purchaseorderstatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementPurchaseorderstatusesByParentIdEmailtemplatesCountRequest struct via the builder pattern


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


## PatchProcurementPurchaseorderstatusesByParentIdEmailtemplatesById

> PurchaseOrderStatusEmailTemplate PatchProcurementPurchaseorderstatusesByParentIdEmailtemplatesById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch PurchaseOrderStatusEmailTemplate

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
	parentId := int32(56) // int32 | purchaseorderstatusId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrderStatusEmailTemplatesAPI.PatchProcurementPurchaseorderstatusesByParentIdEmailtemplatesById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderStatusEmailTemplatesAPI.PatchProcurementPurchaseorderstatusesByParentIdEmailtemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProcurementPurchaseorderstatusesByParentIdEmailtemplatesById`: PurchaseOrderStatusEmailTemplate
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderStatusEmailTemplatesAPI.PatchProcurementPurchaseorderstatusesByParentIdEmailtemplatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailtemplateId | 
**parentId** | **int32** | purchaseorderstatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProcurementPurchaseorderstatusesByParentIdEmailtemplatesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**PurchaseOrderStatusEmailTemplate**](PurchaseOrderStatusEmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProcurementPurchaseorderstatusesByParentIdEmailtemplates

> PurchaseOrderStatusEmailTemplate PostProcurementPurchaseorderstatusesByParentIdEmailtemplates(ctx, parentId).ClientId(clientId).PurchaseOrderStatusEmailTemplate(purchaseOrderStatusEmailTemplate).Execute()

Post PurchaseOrderStatusEmailTemplate

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
	parentId := int32(56) // int32 | purchaseorderstatusId
	clientId := "clientId_example" // string | 
	purchaseOrderStatusEmailTemplate := *openapiclient.NewPurchaseOrderStatusEmailTemplate("Subject_example") // PurchaseOrderStatusEmailTemplate | purchaseOrderStatusEmailTemplate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrderStatusEmailTemplatesAPI.PostProcurementPurchaseorderstatusesByParentIdEmailtemplates(context.Background(), parentId).ClientId(clientId).PurchaseOrderStatusEmailTemplate(purchaseOrderStatusEmailTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderStatusEmailTemplatesAPI.PostProcurementPurchaseorderstatusesByParentIdEmailtemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementPurchaseorderstatusesByParentIdEmailtemplates`: PurchaseOrderStatusEmailTemplate
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderStatusEmailTemplatesAPI.PostProcurementPurchaseorderstatusesByParentIdEmailtemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | purchaseorderstatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementPurchaseorderstatusesByParentIdEmailtemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **purchaseOrderStatusEmailTemplate** | [**PurchaseOrderStatusEmailTemplate**](PurchaseOrderStatusEmailTemplate.md) | purchaseOrderStatusEmailTemplate | 

### Return type

[**PurchaseOrderStatusEmailTemplate**](PurchaseOrderStatusEmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProcurementPurchaseorderstatusesByParentIdEmailtemplatesById

> PurchaseOrderStatusEmailTemplate PutProcurementPurchaseorderstatusesByParentIdEmailtemplatesById(ctx, id, parentId).ClientId(clientId).PurchaseOrderStatusEmailTemplate(purchaseOrderStatusEmailTemplate).Execute()

Put PurchaseOrderStatusEmailTemplate

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
	parentId := int32(56) // int32 | purchaseorderstatusId
	clientId := "clientId_example" // string | 
	purchaseOrderStatusEmailTemplate := *openapiclient.NewPurchaseOrderStatusEmailTemplate("Subject_example") // PurchaseOrderStatusEmailTemplate | purchaseOrderStatusEmailTemplate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PurchaseOrderStatusEmailTemplatesAPI.PutProcurementPurchaseorderstatusesByParentIdEmailtemplatesById(context.Background(), id, parentId).ClientId(clientId).PurchaseOrderStatusEmailTemplate(purchaseOrderStatusEmailTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PurchaseOrderStatusEmailTemplatesAPI.PutProcurementPurchaseorderstatusesByParentIdEmailtemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProcurementPurchaseorderstatusesByParentIdEmailtemplatesById`: PurchaseOrderStatusEmailTemplate
	fmt.Fprintf(os.Stdout, "Response from `PurchaseOrderStatusEmailTemplatesAPI.PutProcurementPurchaseorderstatusesByParentIdEmailtemplatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailtemplateId | 
**parentId** | **int32** | purchaseorderstatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProcurementPurchaseorderstatusesByParentIdEmailtemplatesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **purchaseOrderStatusEmailTemplate** | [**PurchaseOrderStatusEmailTemplate**](PurchaseOrderStatusEmailTemplate.md) | purchaseOrderStatusEmailTemplate | 

### Return type

[**PurchaseOrderStatusEmailTemplate**](PurchaseOrderStatusEmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

