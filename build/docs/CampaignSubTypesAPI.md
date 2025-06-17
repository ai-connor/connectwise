# \CampaignSubTypesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMarketingCampaignsSubTypesById**](CampaignSubTypesAPI.md#DeleteMarketingCampaignsSubTypesById) | **Delete** /marketing/campaigns/subTypes/{id} | Delete CampaignSubType
[**GetMarketingCampaignsSubTypes**](CampaignSubTypesAPI.md#GetMarketingCampaignsSubTypes) | **Get** /marketing/campaigns/subTypes | Get List of CampaignSubType
[**GetMarketingCampaignsSubTypesById**](CampaignSubTypesAPI.md#GetMarketingCampaignsSubTypesById) | **Get** /marketing/campaigns/subTypes/{id} | Get CampaignSubType
[**GetMarketingCampaignsSubTypesCount**](CampaignSubTypesAPI.md#GetMarketingCampaignsSubTypesCount) | **Get** /marketing/campaigns/subTypes/count | Get Count of CampaignSubType
[**PatchMarketingCampaignsSubTypesById**](CampaignSubTypesAPI.md#PatchMarketingCampaignsSubTypesById) | **Patch** /marketing/campaigns/subTypes/{id} | Patch CampaignSubType
[**PostMarketingCampaignsSubTypes**](CampaignSubTypesAPI.md#PostMarketingCampaignsSubTypes) | **Post** /marketing/campaigns/subTypes | Post CampaignSubType
[**PutMarketingCampaignsSubTypesById**](CampaignSubTypesAPI.md#PutMarketingCampaignsSubTypesById) | **Put** /marketing/campaigns/subTypes/{id} | Put CampaignSubType



## DeleteMarketingCampaignsSubTypesById

> DeleteMarketingCampaignsSubTypesById(ctx, id).ClientId(clientId).Execute()

Delete CampaignSubType

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
	id := int32(56) // int32 | subTypeId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CampaignSubTypesAPI.DeleteMarketingCampaignsSubTypesById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignSubTypesAPI.DeleteMarketingCampaignsSubTypesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | subTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketingCampaignsSubTypesByIdRequest struct via the builder pattern


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


## GetMarketingCampaignsSubTypes

> []CampaignSubTypeCampaignSubType GetMarketingCampaignsSubTypes(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of CampaignSubType

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
	resp, r, err := apiClient.CampaignSubTypesAPI.GetMarketingCampaignsSubTypes(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignSubTypesAPI.GetMarketingCampaignsSubTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingCampaignsSubTypes`: []CampaignSubTypeCampaignSubType
	fmt.Fprintf(os.Stdout, "Response from `CampaignSubTypesAPI.GetMarketingCampaignsSubTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingCampaignsSubTypesRequest struct via the builder pattern


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

[**[]CampaignSubTypeCampaignSubType**](CampaignSubTypeCampaignSubType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingCampaignsSubTypesById

> CampaignSubTypeCampaignSubType GetMarketingCampaignsSubTypesById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get CampaignSubType

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
	id := int32(56) // int32 | subTypeId
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
	resp, r, err := apiClient.CampaignSubTypesAPI.GetMarketingCampaignsSubTypesById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignSubTypesAPI.GetMarketingCampaignsSubTypesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingCampaignsSubTypesById`: CampaignSubTypeCampaignSubType
	fmt.Fprintf(os.Stdout, "Response from `CampaignSubTypesAPI.GetMarketingCampaignsSubTypesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | subTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingCampaignsSubTypesByIdRequest struct via the builder pattern


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

[**CampaignSubTypeCampaignSubType**](CampaignSubTypeCampaignSubType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingCampaignsSubTypesCount

> Count GetMarketingCampaignsSubTypesCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of CampaignSubType

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
	resp, r, err := apiClient.CampaignSubTypesAPI.GetMarketingCampaignsSubTypesCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignSubTypesAPI.GetMarketingCampaignsSubTypesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingCampaignsSubTypesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `CampaignSubTypesAPI.GetMarketingCampaignsSubTypesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingCampaignsSubTypesCountRequest struct via the builder pattern


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


## PatchMarketingCampaignsSubTypesById

> CampaignSubTypeCampaignSubType PatchMarketingCampaignsSubTypesById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch CampaignSubType

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
	id := int32(56) // int32 | subTypeId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignSubTypesAPI.PatchMarketingCampaignsSubTypesById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignSubTypesAPI.PatchMarketingCampaignsSubTypesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMarketingCampaignsSubTypesById`: CampaignSubTypeCampaignSubType
	fmt.Fprintf(os.Stdout, "Response from `CampaignSubTypesAPI.PatchMarketingCampaignsSubTypesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | subTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMarketingCampaignsSubTypesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**CampaignSubTypeCampaignSubType**](CampaignSubTypeCampaignSubType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingCampaignsSubTypes

> CampaignSubTypeCampaignSubType PostMarketingCampaignsSubTypes(ctx).ClientId(clientId).CampaignSubTypeCampaignSubType(campaignSubTypeCampaignSubType).Execute()

Post CampaignSubType

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
	campaignSubTypeCampaignSubType := *openapiclient.NewCampaignSubTypeCampaignSubType(*openapiclient.NewCampaignTypeReference(), "Name_example") // CampaignSubTypeCampaignSubType | campaignSubType

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignSubTypesAPI.PostMarketingCampaignsSubTypes(context.Background()).ClientId(clientId).CampaignSubTypeCampaignSubType(campaignSubTypeCampaignSubType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignSubTypesAPI.PostMarketingCampaignsSubTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingCampaignsSubTypes`: CampaignSubTypeCampaignSubType
	fmt.Fprintf(os.Stdout, "Response from `CampaignSubTypesAPI.PostMarketingCampaignsSubTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingCampaignsSubTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **campaignSubTypeCampaignSubType** | [**CampaignSubTypeCampaignSubType**](CampaignSubTypeCampaignSubType.md) | campaignSubType | 

### Return type

[**CampaignSubTypeCampaignSubType**](CampaignSubTypeCampaignSubType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMarketingCampaignsSubTypesById

> CampaignSubTypeCampaignSubType PutMarketingCampaignsSubTypesById(ctx, id).ClientId(clientId).CampaignSubTypeCampaignSubType(campaignSubTypeCampaignSubType).Execute()

Put CampaignSubType

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
	id := int32(56) // int32 | subTypeId
	clientId := "clientId_example" // string | 
	campaignSubTypeCampaignSubType := *openapiclient.NewCampaignSubTypeCampaignSubType(*openapiclient.NewCampaignTypeReference(), "Name_example") // CampaignSubTypeCampaignSubType | campaignSubType

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignSubTypesAPI.PutMarketingCampaignsSubTypesById(context.Background(), id).ClientId(clientId).CampaignSubTypeCampaignSubType(campaignSubTypeCampaignSubType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignSubTypesAPI.PutMarketingCampaignsSubTypesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutMarketingCampaignsSubTypesById`: CampaignSubTypeCampaignSubType
	fmt.Fprintf(os.Stdout, "Response from `CampaignSubTypesAPI.PutMarketingCampaignsSubTypesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | subTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMarketingCampaignsSubTypesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **campaignSubTypeCampaignSubType** | [**CampaignSubTypeCampaignSubType**](CampaignSubTypeCampaignSubType.md) | campaignSubType | 

### Return type

[**CampaignSubTypeCampaignSubType**](CampaignSubTypeCampaignSubType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

