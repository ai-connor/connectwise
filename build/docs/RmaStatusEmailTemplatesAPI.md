# \RmaStatusEmailTemplatesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProcurementRmaStatusesByParentIdEmailtemplatesById**](RmaStatusEmailTemplatesAPI.md#DeleteProcurementRmaStatusesByParentIdEmailtemplatesById) | **Delete** /procurement/rmaStatuses/{parentId}/emailtemplates/{id} | Delete RmaStatusEmailTemplate
[**GetProcurementRmaStatusesByParentIdEmailTemplates**](RmaStatusEmailTemplatesAPI.md#GetProcurementRmaStatusesByParentIdEmailTemplates) | **Get** /procurement/rmaStatuses/{parentId}/emailTemplates/ | Get List of RmaStatusEmailTemplate
[**GetProcurementRmaStatusesByParentIdEmailtemplatesById**](RmaStatusEmailTemplatesAPI.md#GetProcurementRmaStatusesByParentIdEmailtemplatesById) | **Get** /procurement/rmaStatuses/{parentId}/emailtemplates/{id} | Get RmaStatusEmailTemplate
[**GetProcurementRmaStatusesByParentIdEmailtemplatesCount**](RmaStatusEmailTemplatesAPI.md#GetProcurementRmaStatusesByParentIdEmailtemplatesCount) | **Get** /procurement/rmaStatuses/{parentId}/emailtemplates/count | Get Count of RmaStatusEmailTemplate
[**PatchProcurementRmaStatusesByParentIdEmailtemplatesById**](RmaStatusEmailTemplatesAPI.md#PatchProcurementRmaStatusesByParentIdEmailtemplatesById) | **Patch** /procurement/rmaStatuses/{parentId}/emailtemplates/{id} | Patch RmaStatusEmailTemplate
[**PostProcurementRmaStatusesByParentIdEmailtemplates**](RmaStatusEmailTemplatesAPI.md#PostProcurementRmaStatusesByParentIdEmailtemplates) | **Post** /procurement/rmaStatuses/{parentId}/emailtemplates/ | Post RmaStatusEmailTemplate
[**PutProcurementRmaStatusesByParentIdEmailtemplatesById**](RmaStatusEmailTemplatesAPI.md#PutProcurementRmaStatusesByParentIdEmailtemplatesById) | **Put** /procurement/rmaStatuses/{parentId}/emailtemplates/{id} | Put RmaStatusEmailTemplate



## DeleteProcurementRmaStatusesByParentIdEmailtemplatesById

> DeleteProcurementRmaStatusesByParentIdEmailtemplatesById(ctx, id, parentId).ClientId(clientId).Execute()

Delete RmaStatusEmailTemplate

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
	parentId := int32(56) // int32 | rmaStatusId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RmaStatusEmailTemplatesAPI.DeleteProcurementRmaStatusesByParentIdEmailtemplatesById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaStatusEmailTemplatesAPI.DeleteProcurementRmaStatusesByParentIdEmailtemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailtemplateId | 
**parentId** | **int32** | rmaStatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcurementRmaStatusesByParentIdEmailtemplatesByIdRequest struct via the builder pattern


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


## GetProcurementRmaStatusesByParentIdEmailTemplates

> []RmaStatusEmailTemplate GetProcurementRmaStatusesByParentIdEmailTemplates(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of RmaStatusEmailTemplate

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
	parentId := int32(56) // int32 | rmaStatusId
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
	resp, r, err := apiClient.RmaStatusEmailTemplatesAPI.GetProcurementRmaStatusesByParentIdEmailTemplates(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaStatusEmailTemplatesAPI.GetProcurementRmaStatusesByParentIdEmailTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementRmaStatusesByParentIdEmailTemplates`: []RmaStatusEmailTemplate
	fmt.Fprintf(os.Stdout, "Response from `RmaStatusEmailTemplatesAPI.GetProcurementRmaStatusesByParentIdEmailTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | rmaStatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementRmaStatusesByParentIdEmailTemplatesRequest struct via the builder pattern


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

[**[]RmaStatusEmailTemplate**](RmaStatusEmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementRmaStatusesByParentIdEmailtemplatesById

> RmaStatusEmailTemplate GetProcurementRmaStatusesByParentIdEmailtemplatesById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get RmaStatusEmailTemplate

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
	parentId := int32(56) // int32 | rmaStatusId
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
	resp, r, err := apiClient.RmaStatusEmailTemplatesAPI.GetProcurementRmaStatusesByParentIdEmailtemplatesById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaStatusEmailTemplatesAPI.GetProcurementRmaStatusesByParentIdEmailtemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementRmaStatusesByParentIdEmailtemplatesById`: RmaStatusEmailTemplate
	fmt.Fprintf(os.Stdout, "Response from `RmaStatusEmailTemplatesAPI.GetProcurementRmaStatusesByParentIdEmailtemplatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailtemplateId | 
**parentId** | **int32** | rmaStatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementRmaStatusesByParentIdEmailtemplatesByIdRequest struct via the builder pattern


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

[**RmaStatusEmailTemplate**](RmaStatusEmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementRmaStatusesByParentIdEmailtemplatesCount

> Count GetProcurementRmaStatusesByParentIdEmailtemplatesCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of RmaStatusEmailTemplate

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
	parentId := int32(56) // int32 | rmaStatusId
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
	resp, r, err := apiClient.RmaStatusEmailTemplatesAPI.GetProcurementRmaStatusesByParentIdEmailtemplatesCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaStatusEmailTemplatesAPI.GetProcurementRmaStatusesByParentIdEmailtemplatesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementRmaStatusesByParentIdEmailtemplatesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `RmaStatusEmailTemplatesAPI.GetProcurementRmaStatusesByParentIdEmailtemplatesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | rmaStatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementRmaStatusesByParentIdEmailtemplatesCountRequest struct via the builder pattern


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


## PatchProcurementRmaStatusesByParentIdEmailtemplatesById

> RmaStatusEmailTemplate PatchProcurementRmaStatusesByParentIdEmailtemplatesById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch RmaStatusEmailTemplate

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
	parentId := int32(56) // int32 | rmaStatusId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RmaStatusEmailTemplatesAPI.PatchProcurementRmaStatusesByParentIdEmailtemplatesById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaStatusEmailTemplatesAPI.PatchProcurementRmaStatusesByParentIdEmailtemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProcurementRmaStatusesByParentIdEmailtemplatesById`: RmaStatusEmailTemplate
	fmt.Fprintf(os.Stdout, "Response from `RmaStatusEmailTemplatesAPI.PatchProcurementRmaStatusesByParentIdEmailtemplatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailtemplateId | 
**parentId** | **int32** | rmaStatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProcurementRmaStatusesByParentIdEmailtemplatesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**RmaStatusEmailTemplate**](RmaStatusEmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProcurementRmaStatusesByParentIdEmailtemplates

> RmaStatusEmailTemplate PostProcurementRmaStatusesByParentIdEmailtemplates(ctx, parentId).ClientId(clientId).RmaStatusEmailTemplate(rmaStatusEmailTemplate).Execute()

Post RmaStatusEmailTemplate

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
	parentId := int32(56) // int32 | rmaStatusId
	clientId := "clientId_example" // string | 
	rmaStatusEmailTemplate := *openapiclient.NewRmaStatusEmailTemplate("Subject_example", "Body_example") // RmaStatusEmailTemplate | rmaStatusEmailTemplate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RmaStatusEmailTemplatesAPI.PostProcurementRmaStatusesByParentIdEmailtemplates(context.Background(), parentId).ClientId(clientId).RmaStatusEmailTemplate(rmaStatusEmailTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaStatusEmailTemplatesAPI.PostProcurementRmaStatusesByParentIdEmailtemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementRmaStatusesByParentIdEmailtemplates`: RmaStatusEmailTemplate
	fmt.Fprintf(os.Stdout, "Response from `RmaStatusEmailTemplatesAPI.PostProcurementRmaStatusesByParentIdEmailtemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | rmaStatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementRmaStatusesByParentIdEmailtemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **rmaStatusEmailTemplate** | [**RmaStatusEmailTemplate**](RmaStatusEmailTemplate.md) | rmaStatusEmailTemplate | 

### Return type

[**RmaStatusEmailTemplate**](RmaStatusEmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProcurementRmaStatusesByParentIdEmailtemplatesById

> RmaStatusEmailTemplate PutProcurementRmaStatusesByParentIdEmailtemplatesById(ctx, id, parentId).ClientId(clientId).RmaStatusEmailTemplate(rmaStatusEmailTemplate).Execute()

Put RmaStatusEmailTemplate

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
	parentId := int32(56) // int32 | rmaStatusId
	clientId := "clientId_example" // string | 
	rmaStatusEmailTemplate := *openapiclient.NewRmaStatusEmailTemplate("Subject_example", "Body_example") // RmaStatusEmailTemplate | rmaStatusEmailTemplate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RmaStatusEmailTemplatesAPI.PutProcurementRmaStatusesByParentIdEmailtemplatesById(context.Background(), id, parentId).ClientId(clientId).RmaStatusEmailTemplate(rmaStatusEmailTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RmaStatusEmailTemplatesAPI.PutProcurementRmaStatusesByParentIdEmailtemplatesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProcurementRmaStatusesByParentIdEmailtemplatesById`: RmaStatusEmailTemplate
	fmt.Fprintf(os.Stdout, "Response from `RmaStatusEmailTemplatesAPI.PutProcurementRmaStatusesByParentIdEmailtemplatesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailtemplateId | 
**parentId** | **int32** | rmaStatusId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProcurementRmaStatusesByParentIdEmailtemplatesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **rmaStatusEmailTemplate** | [**RmaStatusEmailTemplate**](RmaStatusEmailTemplate.md) | rmaStatusEmailTemplate | 

### Return type

[**RmaStatusEmailTemplate**](RmaStatusEmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

