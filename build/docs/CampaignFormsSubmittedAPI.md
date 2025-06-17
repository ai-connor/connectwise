# \CampaignFormsSubmittedAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMarketingCampaignsByParentIdFormsSubmittedById**](CampaignFormsSubmittedAPI.md#DeleteMarketingCampaignsByParentIdFormsSubmittedById) | **Delete** /marketing/campaigns/{parentId}/formsSubmitted/{id} | Delete FormSubmitted
[**GetMarketingCampaignsByParentIdFormsSubmitted**](CampaignFormsSubmittedAPI.md#GetMarketingCampaignsByParentIdFormsSubmitted) | **Get** /marketing/campaigns/{parentId}/formsSubmitted | Get List of FormSubmitted
[**GetMarketingCampaignsByParentIdFormsSubmittedById**](CampaignFormsSubmittedAPI.md#GetMarketingCampaignsByParentIdFormsSubmittedById) | **Get** /marketing/campaigns/{parentId}/formsSubmitted/{id} | Get FormSubmitted
[**GetMarketingCampaignsByParentIdFormsSubmittedCount**](CampaignFormsSubmittedAPI.md#GetMarketingCampaignsByParentIdFormsSubmittedCount) | **Get** /marketing/campaigns/{parentId}/formsSubmitted/count | Get Count of FormSubmitted
[**PatchMarketingCampaignsByParentIdFormsSubmittedById**](CampaignFormsSubmittedAPI.md#PatchMarketingCampaignsByParentIdFormsSubmittedById) | **Patch** /marketing/campaigns/{parentId}/formsSubmitted/{id} | Patch FormSubmitted
[**PostMarketingCampaignsByParentIdFormsSubmitted**](CampaignFormsSubmittedAPI.md#PostMarketingCampaignsByParentIdFormsSubmitted) | **Post** /marketing/campaigns/{parentId}/formsSubmitted | Post FormSubmitted
[**PutMarketingCampaignsByParentIdFormsSubmittedById**](CampaignFormsSubmittedAPI.md#PutMarketingCampaignsByParentIdFormsSubmittedById) | **Put** /marketing/campaigns/{parentId}/formsSubmitted/{id} | Put FormSubmitted



## DeleteMarketingCampaignsByParentIdFormsSubmittedById

> DeleteMarketingCampaignsByParentIdFormsSubmittedById(ctx, id, parentId).ClientId(clientId).Execute()

Delete FormSubmitted

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
	id := int32(56) // int32 | formsSubmittedId
	parentId := int32(56) // int32 | campaignId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CampaignFormsSubmittedAPI.DeleteMarketingCampaignsByParentIdFormsSubmittedById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignFormsSubmittedAPI.DeleteMarketingCampaignsByParentIdFormsSubmittedById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | formsSubmittedId | 
**parentId** | **int32** | campaignId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketingCampaignsByParentIdFormsSubmittedByIdRequest struct via the builder pattern


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


## GetMarketingCampaignsByParentIdFormsSubmitted

> []FormSubmitted GetMarketingCampaignsByParentIdFormsSubmitted(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of FormSubmitted

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
	parentId := int32(56) // int32 | campaignId
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
	resp, r, err := apiClient.CampaignFormsSubmittedAPI.GetMarketingCampaignsByParentIdFormsSubmitted(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignFormsSubmittedAPI.GetMarketingCampaignsByParentIdFormsSubmitted``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingCampaignsByParentIdFormsSubmitted`: []FormSubmitted
	fmt.Fprintf(os.Stdout, "Response from `CampaignFormsSubmittedAPI.GetMarketingCampaignsByParentIdFormsSubmitted`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | campaignId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingCampaignsByParentIdFormsSubmittedRequest struct via the builder pattern


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

[**[]FormSubmitted**](FormSubmitted.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingCampaignsByParentIdFormsSubmittedById

> FormSubmitted GetMarketingCampaignsByParentIdFormsSubmittedById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get FormSubmitted

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
	id := int32(56) // int32 | formsSubmittedId
	parentId := int32(56) // int32 | campaignId
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
	resp, r, err := apiClient.CampaignFormsSubmittedAPI.GetMarketingCampaignsByParentIdFormsSubmittedById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignFormsSubmittedAPI.GetMarketingCampaignsByParentIdFormsSubmittedById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingCampaignsByParentIdFormsSubmittedById`: FormSubmitted
	fmt.Fprintf(os.Stdout, "Response from `CampaignFormsSubmittedAPI.GetMarketingCampaignsByParentIdFormsSubmittedById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | formsSubmittedId | 
**parentId** | **int32** | campaignId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingCampaignsByParentIdFormsSubmittedByIdRequest struct via the builder pattern


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

[**FormSubmitted**](FormSubmitted.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingCampaignsByParentIdFormsSubmittedCount

> Count GetMarketingCampaignsByParentIdFormsSubmittedCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of FormSubmitted

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
	parentId := int32(56) // int32 | campaignId
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
	resp, r, err := apiClient.CampaignFormsSubmittedAPI.GetMarketingCampaignsByParentIdFormsSubmittedCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignFormsSubmittedAPI.GetMarketingCampaignsByParentIdFormsSubmittedCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingCampaignsByParentIdFormsSubmittedCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `CampaignFormsSubmittedAPI.GetMarketingCampaignsByParentIdFormsSubmittedCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | campaignId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingCampaignsByParentIdFormsSubmittedCountRequest struct via the builder pattern


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


## PatchMarketingCampaignsByParentIdFormsSubmittedById

> FormSubmitted PatchMarketingCampaignsByParentIdFormsSubmittedById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch FormSubmitted

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
	id := int32(56) // int32 | formsSubmittedId
	parentId := int32(56) // int32 | campaignId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignFormsSubmittedAPI.PatchMarketingCampaignsByParentIdFormsSubmittedById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignFormsSubmittedAPI.PatchMarketingCampaignsByParentIdFormsSubmittedById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMarketingCampaignsByParentIdFormsSubmittedById`: FormSubmitted
	fmt.Fprintf(os.Stdout, "Response from `CampaignFormsSubmittedAPI.PatchMarketingCampaignsByParentIdFormsSubmittedById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | formsSubmittedId | 
**parentId** | **int32** | campaignId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMarketingCampaignsByParentIdFormsSubmittedByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**FormSubmitted**](FormSubmitted.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingCampaignsByParentIdFormsSubmitted

> FormSubmitted PostMarketingCampaignsByParentIdFormsSubmitted(ctx, parentId).ClientId(clientId).FormSubmitted(formSubmitted).Execute()

Post FormSubmitted

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
	parentId := int32(56) // int32 | campaignId
	clientId := "clientId_example" // string | 
	formSubmitted := *openapiclient.NewFormSubmitted(NullableInt32(123), "Url_example") // FormSubmitted | formSubmitted

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignFormsSubmittedAPI.PostMarketingCampaignsByParentIdFormsSubmitted(context.Background(), parentId).ClientId(clientId).FormSubmitted(formSubmitted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignFormsSubmittedAPI.PostMarketingCampaignsByParentIdFormsSubmitted``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingCampaignsByParentIdFormsSubmitted`: FormSubmitted
	fmt.Fprintf(os.Stdout, "Response from `CampaignFormsSubmittedAPI.PostMarketingCampaignsByParentIdFormsSubmitted`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | campaignId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingCampaignsByParentIdFormsSubmittedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **formSubmitted** | [**FormSubmitted**](FormSubmitted.md) | formSubmitted | 

### Return type

[**FormSubmitted**](FormSubmitted.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMarketingCampaignsByParentIdFormsSubmittedById

> FormSubmitted PutMarketingCampaignsByParentIdFormsSubmittedById(ctx, id, parentId).ClientId(clientId).FormSubmitted(formSubmitted).Execute()

Put FormSubmitted

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
	id := int32(56) // int32 | formsSubmittedId
	parentId := int32(56) // int32 | campaignId
	clientId := "clientId_example" // string | 
	formSubmitted := *openapiclient.NewFormSubmitted(NullableInt32(123), "Url_example") // FormSubmitted | formSubmitted

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignFormsSubmittedAPI.PutMarketingCampaignsByParentIdFormsSubmittedById(context.Background(), id, parentId).ClientId(clientId).FormSubmitted(formSubmitted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignFormsSubmittedAPI.PutMarketingCampaignsByParentIdFormsSubmittedById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutMarketingCampaignsByParentIdFormsSubmittedById`: FormSubmitted
	fmt.Fprintf(os.Stdout, "Response from `CampaignFormsSubmittedAPI.PutMarketingCampaignsByParentIdFormsSubmittedById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | formsSubmittedId | 
**parentId** | **int32** | campaignId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMarketingCampaignsByParentIdFormsSubmittedByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **formSubmitted** | [**FormSubmitted**](FormSubmitted.md) | formSubmitted | 

### Return type

[**FormSubmitted**](FormSubmitted.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

