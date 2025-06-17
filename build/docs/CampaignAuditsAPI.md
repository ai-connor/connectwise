# \CampaignAuditsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMarketingCampaignsByParentIdAuditsById**](CampaignAuditsAPI.md#DeleteMarketingCampaignsByParentIdAuditsById) | **Delete** /marketing/campaigns/{parentId}/audits/{id} | Delete CampaignAudit
[**GetMarketingCampaignsByParentIdAudits**](CampaignAuditsAPI.md#GetMarketingCampaignsByParentIdAudits) | **Get** /marketing/campaigns/{parentId}/audits | Get List of CampaignAudit
[**GetMarketingCampaignsByParentIdAuditsById**](CampaignAuditsAPI.md#GetMarketingCampaignsByParentIdAuditsById) | **Get** /marketing/campaigns/{parentId}/audits/{id} | Get CampaignAudit
[**GetMarketingCampaignsByParentIdAuditsCount**](CampaignAuditsAPI.md#GetMarketingCampaignsByParentIdAuditsCount) | **Get** /marketing/campaigns/{parentId}/audits/count | Get Count of CampaignAudit
[**PatchMarketingCampaignsByParentIdAuditsById**](CampaignAuditsAPI.md#PatchMarketingCampaignsByParentIdAuditsById) | **Patch** /marketing/campaigns/{parentId}/audits/{id} | Patch CampaignAudit
[**PostMarketingCampaignsByParentIdAudits**](CampaignAuditsAPI.md#PostMarketingCampaignsByParentIdAudits) | **Post** /marketing/campaigns/{parentId}/audits | Post CampaignAudit
[**PutMarketingCampaignsByParentIdAuditsById**](CampaignAuditsAPI.md#PutMarketingCampaignsByParentIdAuditsById) | **Put** /marketing/campaigns/{parentId}/audits/{id} | Put CampaignAudit



## DeleteMarketingCampaignsByParentIdAuditsById

> DeleteMarketingCampaignsByParentIdAuditsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete CampaignAudit

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
	id := int32(56) // int32 | auditId
	parentId := int32(56) // int32 | campaignId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CampaignAuditsAPI.DeleteMarketingCampaignsByParentIdAuditsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAuditsAPI.DeleteMarketingCampaignsByParentIdAuditsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | auditId | 
**parentId** | **int32** | campaignId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketingCampaignsByParentIdAuditsByIdRequest struct via the builder pattern


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


## GetMarketingCampaignsByParentIdAudits

> []CampaignAudit GetMarketingCampaignsByParentIdAudits(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of CampaignAudit

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
	resp, r, err := apiClient.CampaignAuditsAPI.GetMarketingCampaignsByParentIdAudits(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAuditsAPI.GetMarketingCampaignsByParentIdAudits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingCampaignsByParentIdAudits`: []CampaignAudit
	fmt.Fprintf(os.Stdout, "Response from `CampaignAuditsAPI.GetMarketingCampaignsByParentIdAudits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | campaignId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingCampaignsByParentIdAuditsRequest struct via the builder pattern


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

[**[]CampaignAudit**](CampaignAudit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingCampaignsByParentIdAuditsById

> CampaignAudit GetMarketingCampaignsByParentIdAuditsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get CampaignAudit

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
	id := int32(56) // int32 | auditId
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
	resp, r, err := apiClient.CampaignAuditsAPI.GetMarketingCampaignsByParentIdAuditsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAuditsAPI.GetMarketingCampaignsByParentIdAuditsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingCampaignsByParentIdAuditsById`: CampaignAudit
	fmt.Fprintf(os.Stdout, "Response from `CampaignAuditsAPI.GetMarketingCampaignsByParentIdAuditsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | auditId | 
**parentId** | **int32** | campaignId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingCampaignsByParentIdAuditsByIdRequest struct via the builder pattern


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

[**CampaignAudit**](CampaignAudit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingCampaignsByParentIdAuditsCount

> Count GetMarketingCampaignsByParentIdAuditsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of CampaignAudit

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
	resp, r, err := apiClient.CampaignAuditsAPI.GetMarketingCampaignsByParentIdAuditsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAuditsAPI.GetMarketingCampaignsByParentIdAuditsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingCampaignsByParentIdAuditsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `CampaignAuditsAPI.GetMarketingCampaignsByParentIdAuditsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | campaignId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingCampaignsByParentIdAuditsCountRequest struct via the builder pattern


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


## PatchMarketingCampaignsByParentIdAuditsById

> CampaignAudit PatchMarketingCampaignsByParentIdAuditsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch CampaignAudit

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
	id := int32(56) // int32 | auditId
	parentId := int32(56) // int32 | campaignId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAuditsAPI.PatchMarketingCampaignsByParentIdAuditsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAuditsAPI.PatchMarketingCampaignsByParentIdAuditsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMarketingCampaignsByParentIdAuditsById`: CampaignAudit
	fmt.Fprintf(os.Stdout, "Response from `CampaignAuditsAPI.PatchMarketingCampaignsByParentIdAuditsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | auditId | 
**parentId** | **int32** | campaignId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMarketingCampaignsByParentIdAuditsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**CampaignAudit**](CampaignAudit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingCampaignsByParentIdAudits

> CampaignAudit PostMarketingCampaignsByParentIdAudits(ctx, parentId).ClientId(clientId).CampaignAudit(campaignAudit).Execute()

Post CampaignAudit

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
	campaignAudit := *openapiclient.NewCampaignAudit(NullableInt32(123)) // CampaignAudit | campaignAudit

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAuditsAPI.PostMarketingCampaignsByParentIdAudits(context.Background(), parentId).ClientId(clientId).CampaignAudit(campaignAudit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAuditsAPI.PostMarketingCampaignsByParentIdAudits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingCampaignsByParentIdAudits`: CampaignAudit
	fmt.Fprintf(os.Stdout, "Response from `CampaignAuditsAPI.PostMarketingCampaignsByParentIdAudits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | campaignId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingCampaignsByParentIdAuditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **campaignAudit** | [**CampaignAudit**](CampaignAudit.md) | campaignAudit | 

### Return type

[**CampaignAudit**](CampaignAudit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMarketingCampaignsByParentIdAuditsById

> CampaignAudit PutMarketingCampaignsByParentIdAuditsById(ctx, id, parentId).ClientId(clientId).CampaignAudit(campaignAudit).Execute()

Put CampaignAudit

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
	id := int32(56) // int32 | auditId
	parentId := int32(56) // int32 | campaignId
	clientId := "clientId_example" // string | 
	campaignAudit := *openapiclient.NewCampaignAudit(NullableInt32(123)) // CampaignAudit | campaignAudit

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAuditsAPI.PutMarketingCampaignsByParentIdAuditsById(context.Background(), id, parentId).ClientId(clientId).CampaignAudit(campaignAudit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAuditsAPI.PutMarketingCampaignsByParentIdAuditsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutMarketingCampaignsByParentIdAuditsById`: CampaignAudit
	fmt.Fprintf(os.Stdout, "Response from `CampaignAuditsAPI.PutMarketingCampaignsByParentIdAuditsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | auditId | 
**parentId** | **int32** | campaignId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMarketingCampaignsByParentIdAuditsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **campaignAudit** | [**CampaignAudit**](CampaignAudit.md) | campaignAudit | 

### Return type

[**CampaignAudit**](CampaignAudit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

