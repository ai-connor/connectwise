# \KnowledgeBaseArticlesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServiceKnowledgeBaseArticlesById**](KnowledgeBaseArticlesAPI.md#DeleteServiceKnowledgeBaseArticlesById) | **Delete** /service/knowledgeBaseArticles/{id} | Delete KnowledgeBaseArticle
[**GetServiceKnowledgeBaseArticles**](KnowledgeBaseArticlesAPI.md#GetServiceKnowledgeBaseArticles) | **Get** /service/knowledgeBaseArticles | Get List of KnowledgeBaseArticle
[**GetServiceKnowledgeBaseArticlesById**](KnowledgeBaseArticlesAPI.md#GetServiceKnowledgeBaseArticlesById) | **Get** /service/knowledgeBaseArticles/{id} | Get KnowledgeBaseArticle
[**GetServiceKnowledgeBaseArticlesCount**](KnowledgeBaseArticlesAPI.md#GetServiceKnowledgeBaseArticlesCount) | **Get** /service/knowledgeBaseArticles/count | Get Count of KnowledgeBaseArticle
[**PatchServiceKnowledgeBaseArticlesById**](KnowledgeBaseArticlesAPI.md#PatchServiceKnowledgeBaseArticlesById) | **Patch** /service/knowledgeBaseArticles/{id} | Patch KnowledgeBaseArticle
[**PostServiceKnowledgeBaseArticles**](KnowledgeBaseArticlesAPI.md#PostServiceKnowledgeBaseArticles) | **Post** /service/knowledgeBaseArticles | Post KnowledgeBaseArticle
[**PutServiceKnowledgeBaseArticlesById**](KnowledgeBaseArticlesAPI.md#PutServiceKnowledgeBaseArticlesById) | **Put** /service/knowledgeBaseArticles/{id} | Put KnowledgeBaseArticle



## DeleteServiceKnowledgeBaseArticlesById

> DeleteServiceKnowledgeBaseArticlesById(ctx, id).ClientId(clientId).Execute()

Delete KnowledgeBaseArticle

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
	id := int32(56) // int32 | knowledgeBaseArticleId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KnowledgeBaseArticlesAPI.DeleteServiceKnowledgeBaseArticlesById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseArticlesAPI.DeleteServiceKnowledgeBaseArticlesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | knowledgeBaseArticleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceKnowledgeBaseArticlesByIdRequest struct via the builder pattern


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


## GetServiceKnowledgeBaseArticles

> []KnowledgeBaseArticle GetServiceKnowledgeBaseArticles(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of KnowledgeBaseArticle

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
	resp, r, err := apiClient.KnowledgeBaseArticlesAPI.GetServiceKnowledgeBaseArticles(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseArticlesAPI.GetServiceKnowledgeBaseArticles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceKnowledgeBaseArticles`: []KnowledgeBaseArticle
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseArticlesAPI.GetServiceKnowledgeBaseArticles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceKnowledgeBaseArticlesRequest struct via the builder pattern


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

[**[]KnowledgeBaseArticle**](KnowledgeBaseArticle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceKnowledgeBaseArticlesById

> KnowledgeBaseArticle GetServiceKnowledgeBaseArticlesById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get KnowledgeBaseArticle

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
	id := int32(56) // int32 | knowledgeBaseArticleId
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
	resp, r, err := apiClient.KnowledgeBaseArticlesAPI.GetServiceKnowledgeBaseArticlesById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseArticlesAPI.GetServiceKnowledgeBaseArticlesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceKnowledgeBaseArticlesById`: KnowledgeBaseArticle
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseArticlesAPI.GetServiceKnowledgeBaseArticlesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | knowledgeBaseArticleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceKnowledgeBaseArticlesByIdRequest struct via the builder pattern


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

[**KnowledgeBaseArticle**](KnowledgeBaseArticle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceKnowledgeBaseArticlesCount

> Count GetServiceKnowledgeBaseArticlesCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of KnowledgeBaseArticle

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
	resp, r, err := apiClient.KnowledgeBaseArticlesAPI.GetServiceKnowledgeBaseArticlesCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseArticlesAPI.GetServiceKnowledgeBaseArticlesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceKnowledgeBaseArticlesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseArticlesAPI.GetServiceKnowledgeBaseArticlesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceKnowledgeBaseArticlesCountRequest struct via the builder pattern


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


## PatchServiceKnowledgeBaseArticlesById

> KnowledgeBaseArticle PatchServiceKnowledgeBaseArticlesById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch KnowledgeBaseArticle

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
	id := int32(56) // int32 | knowledgeBaseArticleId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeBaseArticlesAPI.PatchServiceKnowledgeBaseArticlesById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseArticlesAPI.PatchServiceKnowledgeBaseArticlesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchServiceKnowledgeBaseArticlesById`: KnowledgeBaseArticle
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseArticlesAPI.PatchServiceKnowledgeBaseArticlesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | knowledgeBaseArticleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchServiceKnowledgeBaseArticlesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**KnowledgeBaseArticle**](KnowledgeBaseArticle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostServiceKnowledgeBaseArticles

> KnowledgeBaseArticle PostServiceKnowledgeBaseArticles(ctx).ClientId(clientId).KnowledgeBaseArticle(knowledgeBaseArticle).Execute()

Post KnowledgeBaseArticle

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
	knowledgeBaseArticle := *openapiclient.NewKnowledgeBaseArticle("Title_example", "Issue_example", "Resolution_example") // KnowledgeBaseArticle | knowledgeBaseArticle

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeBaseArticlesAPI.PostServiceKnowledgeBaseArticles(context.Background()).ClientId(clientId).KnowledgeBaseArticle(knowledgeBaseArticle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseArticlesAPI.PostServiceKnowledgeBaseArticles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServiceKnowledgeBaseArticles`: KnowledgeBaseArticle
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseArticlesAPI.PostServiceKnowledgeBaseArticles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostServiceKnowledgeBaseArticlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **knowledgeBaseArticle** | [**KnowledgeBaseArticle**](KnowledgeBaseArticle.md) | knowledgeBaseArticle | 

### Return type

[**KnowledgeBaseArticle**](KnowledgeBaseArticle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServiceKnowledgeBaseArticlesById

> KnowledgeBaseArticle PutServiceKnowledgeBaseArticlesById(ctx, id).ClientId(clientId).KnowledgeBaseArticle(knowledgeBaseArticle).Execute()

Put KnowledgeBaseArticle

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
	id := int32(56) // int32 | knowledgeBaseArticleId
	clientId := "clientId_example" // string | 
	knowledgeBaseArticle := *openapiclient.NewKnowledgeBaseArticle("Title_example", "Issue_example", "Resolution_example") // KnowledgeBaseArticle | knowledgeBaseArticle

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeBaseArticlesAPI.PutServiceKnowledgeBaseArticlesById(context.Background(), id).ClientId(clientId).KnowledgeBaseArticle(knowledgeBaseArticle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeBaseArticlesAPI.PutServiceKnowledgeBaseArticlesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServiceKnowledgeBaseArticlesById`: KnowledgeBaseArticle
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeBaseArticlesAPI.PutServiceKnowledgeBaseArticlesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | knowledgeBaseArticleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServiceKnowledgeBaseArticlesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **knowledgeBaseArticle** | [**KnowledgeBaseArticle**](KnowledgeBaseArticle.md) | knowledgeBaseArticle | 

### Return type

[**KnowledgeBaseArticle**](KnowledgeBaseArticle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

