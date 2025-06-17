# \MarketingCompaniesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMarketingGroupsByParentIdCompaniesById**](MarketingCompaniesAPI.md#DeleteMarketingGroupsByParentIdCompaniesById) | **Delete** /marketing/groups/{parentId}/companies/{id} | Delete MarketingCompany
[**GetMarketingGroupsByParentIdCompanies**](MarketingCompaniesAPI.md#GetMarketingGroupsByParentIdCompanies) | **Get** /marketing/groups/{parentId}/companies | Get List of MarketingCompany
[**GetMarketingGroupsByParentIdCompaniesById**](MarketingCompaniesAPI.md#GetMarketingGroupsByParentIdCompaniesById) | **Get** /marketing/groups/{parentId}/companies/{id} | Get MarketingCompany
[**GetMarketingGroupsByParentIdCompaniesCount**](MarketingCompaniesAPI.md#GetMarketingGroupsByParentIdCompaniesCount) | **Get** /marketing/groups/{parentId}/companies/count | Get Count of MarketingCompany
[**PatchMarketingGroupsByParentIdCompaniesById**](MarketingCompaniesAPI.md#PatchMarketingGroupsByParentIdCompaniesById) | **Patch** /marketing/groups/{parentId}/companies/{id} | Patch MarketingCompany
[**PostMarketingGroupsByParentIdCompanies**](MarketingCompaniesAPI.md#PostMarketingGroupsByParentIdCompanies) | **Post** /marketing/groups/{parentId}/companies | Post MarketingCompany
[**PutMarketingGroupsByParentIdCompaniesById**](MarketingCompaniesAPI.md#PutMarketingGroupsByParentIdCompaniesById) | **Put** /marketing/groups/{parentId}/companies/{id} | Put MarketingCompany



## DeleteMarketingGroupsByParentIdCompaniesById

> DeleteMarketingGroupsByParentIdCompaniesById(ctx, id, parentId).ClientId(clientId).Execute()

Delete MarketingCompany

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
	id := int32(56) // int32 | companyId
	parentId := int32(56) // int32 | groupId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MarketingCompaniesAPI.DeleteMarketingGroupsByParentIdCompaniesById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingCompaniesAPI.DeleteMarketingGroupsByParentIdCompaniesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | companyId | 
**parentId** | **int32** | groupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketingGroupsByParentIdCompaniesByIdRequest struct via the builder pattern


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


## GetMarketingGroupsByParentIdCompanies

> []MarketingCompany GetMarketingGroupsByParentIdCompanies(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of MarketingCompany

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
	parentId := int32(56) // int32 | groupId
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
	resp, r, err := apiClient.MarketingCompaniesAPI.GetMarketingGroupsByParentIdCompanies(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingCompaniesAPI.GetMarketingGroupsByParentIdCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingGroupsByParentIdCompanies`: []MarketingCompany
	fmt.Fprintf(os.Stdout, "Response from `MarketingCompaniesAPI.GetMarketingGroupsByParentIdCompanies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | groupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingGroupsByParentIdCompaniesRequest struct via the builder pattern


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

[**[]MarketingCompany**](MarketingCompany.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingGroupsByParentIdCompaniesById

> MarketingCompany GetMarketingGroupsByParentIdCompaniesById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get MarketingCompany

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
	id := int32(56) // int32 | companyId
	parentId := int32(56) // int32 | groupId
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
	resp, r, err := apiClient.MarketingCompaniesAPI.GetMarketingGroupsByParentIdCompaniesById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingCompaniesAPI.GetMarketingGroupsByParentIdCompaniesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingGroupsByParentIdCompaniesById`: MarketingCompany
	fmt.Fprintf(os.Stdout, "Response from `MarketingCompaniesAPI.GetMarketingGroupsByParentIdCompaniesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | companyId | 
**parentId** | **int32** | groupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingGroupsByParentIdCompaniesByIdRequest struct via the builder pattern


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

[**MarketingCompany**](MarketingCompany.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingGroupsByParentIdCompaniesCount

> Count GetMarketingGroupsByParentIdCompaniesCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of MarketingCompany

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
	parentId := int32(56) // int32 | groupId
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
	resp, r, err := apiClient.MarketingCompaniesAPI.GetMarketingGroupsByParentIdCompaniesCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingCompaniesAPI.GetMarketingGroupsByParentIdCompaniesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingGroupsByParentIdCompaniesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `MarketingCompaniesAPI.GetMarketingGroupsByParentIdCompaniesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | groupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingGroupsByParentIdCompaniesCountRequest struct via the builder pattern


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


## PatchMarketingGroupsByParentIdCompaniesById

> MarketingCompany PatchMarketingGroupsByParentIdCompaniesById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch MarketingCompany

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
	id := int32(56) // int32 | companyId
	parentId := int32(56) // int32 | groupId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingCompaniesAPI.PatchMarketingGroupsByParentIdCompaniesById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingCompaniesAPI.PatchMarketingGroupsByParentIdCompaniesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMarketingGroupsByParentIdCompaniesById`: MarketingCompany
	fmt.Fprintf(os.Stdout, "Response from `MarketingCompaniesAPI.PatchMarketingGroupsByParentIdCompaniesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | companyId | 
**parentId** | **int32** | groupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMarketingGroupsByParentIdCompaniesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**MarketingCompany**](MarketingCompany.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingGroupsByParentIdCompanies

> MarketingCompany PostMarketingGroupsByParentIdCompanies(ctx, parentId).ClientId(clientId).MarketingCompany(marketingCompany).Execute()

Post MarketingCompany

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
	parentId := int32(56) // int32 | groupId
	clientId := "clientId_example" // string | 
	marketingCompany := *openapiclient.NewMarketingCompany() // MarketingCompany | marketingCompany

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingCompaniesAPI.PostMarketingGroupsByParentIdCompanies(context.Background(), parentId).ClientId(clientId).MarketingCompany(marketingCompany).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingCompaniesAPI.PostMarketingGroupsByParentIdCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingGroupsByParentIdCompanies`: MarketingCompany
	fmt.Fprintf(os.Stdout, "Response from `MarketingCompaniesAPI.PostMarketingGroupsByParentIdCompanies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | groupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingGroupsByParentIdCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **marketingCompany** | [**MarketingCompany**](MarketingCompany.md) | marketingCompany | 

### Return type

[**MarketingCompany**](MarketingCompany.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMarketingGroupsByParentIdCompaniesById

> MarketingCompany PutMarketingGroupsByParentIdCompaniesById(ctx, id, parentId).ClientId(clientId).MarketingCompany(marketingCompany).Execute()

Put MarketingCompany

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
	id := int32(56) // int32 | companyId
	parentId := int32(56) // int32 | groupId
	clientId := "clientId_example" // string | 
	marketingCompany := *openapiclient.NewMarketingCompany() // MarketingCompany | marketingCompany

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingCompaniesAPI.PutMarketingGroupsByParentIdCompaniesById(context.Background(), id, parentId).ClientId(clientId).MarketingCompany(marketingCompany).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingCompaniesAPI.PutMarketingGroupsByParentIdCompaniesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutMarketingGroupsByParentIdCompaniesById`: MarketingCompany
	fmt.Fprintf(os.Stdout, "Response from `MarketingCompaniesAPI.PutMarketingGroupsByParentIdCompaniesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | companyId | 
**parentId** | **int32** | groupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMarketingGroupsByParentIdCompaniesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **marketingCompany** | [**MarketingCompany**](MarketingCompany.md) | marketingCompany | 

### Return type

[**MarketingCompany**](MarketingCompany.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

