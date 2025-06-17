# \CampaignEmailsOpenedAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMarketingCampaignsByParentIdEmailsOpenedById**](CampaignEmailsOpenedAPI.md#DeleteMarketingCampaignsByParentIdEmailsOpenedById) | **Delete** /marketing/campaigns/{parentId}/emailsOpened/{id} | Delete EmailOpened
[**GetMarketingCampaignsByParentIdEmailsOpened**](CampaignEmailsOpenedAPI.md#GetMarketingCampaignsByParentIdEmailsOpened) | **Get** /marketing/campaigns/{parentId}/emailsOpened | Get List of EmailOpened
[**GetMarketingCampaignsByParentIdEmailsOpenedById**](CampaignEmailsOpenedAPI.md#GetMarketingCampaignsByParentIdEmailsOpenedById) | **Get** /marketing/campaigns/{parentId}/emailsOpened/{id} | Get EmailOpened
[**GetMarketingCampaignsByParentIdEmailsOpenedCount**](CampaignEmailsOpenedAPI.md#GetMarketingCampaignsByParentIdEmailsOpenedCount) | **Get** /marketing/campaigns/{parentId}/emailsOpened/count | Get Count of EmailOpened
[**PatchMarketingCampaignsByParentIdEmailsOpenedById**](CampaignEmailsOpenedAPI.md#PatchMarketingCampaignsByParentIdEmailsOpenedById) | **Patch** /marketing/campaigns/{parentId}/emailsOpened/{id} | Patch EmailOpened
[**PostMarketingCampaignsByParentIdEmailsOpened**](CampaignEmailsOpenedAPI.md#PostMarketingCampaignsByParentIdEmailsOpened) | **Post** /marketing/campaigns/{parentId}/emailsOpened | Post EmailOpened
[**PutMarketingCampaignsByParentIdEmailsOpenedById**](CampaignEmailsOpenedAPI.md#PutMarketingCampaignsByParentIdEmailsOpenedById) | **Put** /marketing/campaigns/{parentId}/emailsOpened/{id} | Put EmailOpened



## DeleteMarketingCampaignsByParentIdEmailsOpenedById

> DeleteMarketingCampaignsByParentIdEmailsOpenedById(ctx, id, parentId).ClientId(clientId).Execute()

Delete EmailOpened

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
	id := int32(56) // int32 | emailsOpenedId
	parentId := int32(56) // int32 | campaignId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CampaignEmailsOpenedAPI.DeleteMarketingCampaignsByParentIdEmailsOpenedById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignEmailsOpenedAPI.DeleteMarketingCampaignsByParentIdEmailsOpenedById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailsOpenedId | 
**parentId** | **int32** | campaignId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketingCampaignsByParentIdEmailsOpenedByIdRequest struct via the builder pattern


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


## GetMarketingCampaignsByParentIdEmailsOpened

> []EmailOpened GetMarketingCampaignsByParentIdEmailsOpened(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of EmailOpened

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
	resp, r, err := apiClient.CampaignEmailsOpenedAPI.GetMarketingCampaignsByParentIdEmailsOpened(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignEmailsOpenedAPI.GetMarketingCampaignsByParentIdEmailsOpened``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingCampaignsByParentIdEmailsOpened`: []EmailOpened
	fmt.Fprintf(os.Stdout, "Response from `CampaignEmailsOpenedAPI.GetMarketingCampaignsByParentIdEmailsOpened`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | campaignId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingCampaignsByParentIdEmailsOpenedRequest struct via the builder pattern


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

[**[]EmailOpened**](EmailOpened.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingCampaignsByParentIdEmailsOpenedById

> EmailOpened GetMarketingCampaignsByParentIdEmailsOpenedById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get EmailOpened

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
	id := int32(56) // int32 | emailsOpenedId
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
	resp, r, err := apiClient.CampaignEmailsOpenedAPI.GetMarketingCampaignsByParentIdEmailsOpenedById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignEmailsOpenedAPI.GetMarketingCampaignsByParentIdEmailsOpenedById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingCampaignsByParentIdEmailsOpenedById`: EmailOpened
	fmt.Fprintf(os.Stdout, "Response from `CampaignEmailsOpenedAPI.GetMarketingCampaignsByParentIdEmailsOpenedById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailsOpenedId | 
**parentId** | **int32** | campaignId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingCampaignsByParentIdEmailsOpenedByIdRequest struct via the builder pattern


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

[**EmailOpened**](EmailOpened.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingCampaignsByParentIdEmailsOpenedCount

> Count GetMarketingCampaignsByParentIdEmailsOpenedCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of EmailOpened

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
	resp, r, err := apiClient.CampaignEmailsOpenedAPI.GetMarketingCampaignsByParentIdEmailsOpenedCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignEmailsOpenedAPI.GetMarketingCampaignsByParentIdEmailsOpenedCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingCampaignsByParentIdEmailsOpenedCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `CampaignEmailsOpenedAPI.GetMarketingCampaignsByParentIdEmailsOpenedCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | campaignId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingCampaignsByParentIdEmailsOpenedCountRequest struct via the builder pattern


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


## PatchMarketingCampaignsByParentIdEmailsOpenedById

> EmailOpened PatchMarketingCampaignsByParentIdEmailsOpenedById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch EmailOpened

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
	id := int32(56) // int32 | emailsOpenedId
	parentId := int32(56) // int32 | campaignId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignEmailsOpenedAPI.PatchMarketingCampaignsByParentIdEmailsOpenedById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignEmailsOpenedAPI.PatchMarketingCampaignsByParentIdEmailsOpenedById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMarketingCampaignsByParentIdEmailsOpenedById`: EmailOpened
	fmt.Fprintf(os.Stdout, "Response from `CampaignEmailsOpenedAPI.PatchMarketingCampaignsByParentIdEmailsOpenedById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailsOpenedId | 
**parentId** | **int32** | campaignId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMarketingCampaignsByParentIdEmailsOpenedByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**EmailOpened**](EmailOpened.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingCampaignsByParentIdEmailsOpened

> EmailOpened PostMarketingCampaignsByParentIdEmailsOpened(ctx, parentId).ClientId(clientId).EmailOpened(emailOpened).Execute()

Post EmailOpened

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
	emailOpened := *openapiclient.NewEmailOpened(NullableInt32(123)) // EmailOpened | emailOpened

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignEmailsOpenedAPI.PostMarketingCampaignsByParentIdEmailsOpened(context.Background(), parentId).ClientId(clientId).EmailOpened(emailOpened).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignEmailsOpenedAPI.PostMarketingCampaignsByParentIdEmailsOpened``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingCampaignsByParentIdEmailsOpened`: EmailOpened
	fmt.Fprintf(os.Stdout, "Response from `CampaignEmailsOpenedAPI.PostMarketingCampaignsByParentIdEmailsOpened`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | campaignId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingCampaignsByParentIdEmailsOpenedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **emailOpened** | [**EmailOpened**](EmailOpened.md) | emailOpened | 

### Return type

[**EmailOpened**](EmailOpened.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMarketingCampaignsByParentIdEmailsOpenedById

> EmailOpened PutMarketingCampaignsByParentIdEmailsOpenedById(ctx, id, parentId).ClientId(clientId).EmailOpened(emailOpened).Execute()

Put EmailOpened

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
	id := int32(56) // int32 | emailsOpenedId
	parentId := int32(56) // int32 | campaignId
	clientId := "clientId_example" // string | 
	emailOpened := *openapiclient.NewEmailOpened(NullableInt32(123)) // EmailOpened | emailOpened

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignEmailsOpenedAPI.PutMarketingCampaignsByParentIdEmailsOpenedById(context.Background(), id, parentId).ClientId(clientId).EmailOpened(emailOpened).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignEmailsOpenedAPI.PutMarketingCampaignsByParentIdEmailsOpenedById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutMarketingCampaignsByParentIdEmailsOpenedById`: EmailOpened
	fmt.Fprintf(os.Stdout, "Response from `CampaignEmailsOpenedAPI.PutMarketingCampaignsByParentIdEmailsOpenedById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailsOpenedId | 
**parentId** | **int32** | campaignId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMarketingCampaignsByParentIdEmailsOpenedByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **emailOpened** | [**EmailOpened**](EmailOpened.md) | emailOpened | 

### Return type

[**EmailOpened**](EmailOpened.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

