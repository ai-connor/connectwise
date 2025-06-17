# \SalesTeamMembersAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSalesSalesTeamsByParentIdMembersById**](SalesTeamMembersAPI.md#DeleteSalesSalesTeamsByParentIdMembersById) | **Delete** /sales/salesTeams/{parentId}/members/{id} | Delete SalesTeamMember
[**GetSalesSalesTeamsByParentIdMembers**](SalesTeamMembersAPI.md#GetSalesSalesTeamsByParentIdMembers) | **Get** /sales/salesTeams/{parentId}/members | Get List of SalesTeamMember
[**GetSalesSalesTeamsByParentIdMembersById**](SalesTeamMembersAPI.md#GetSalesSalesTeamsByParentIdMembersById) | **Get** /sales/salesTeams/{parentId}/members/{id} | Get SalesTeamMember
[**GetSalesSalesTeamsByParentIdMembersCount**](SalesTeamMembersAPI.md#GetSalesSalesTeamsByParentIdMembersCount) | **Get** /sales/salesTeams/{parentId}/members/count | Get Count of SalesTeamMember
[**PatchSalesSalesTeamsByParentIdMembersById**](SalesTeamMembersAPI.md#PatchSalesSalesTeamsByParentIdMembersById) | **Patch** /sales/salesTeams/{parentId}/members/{id} | Patch SalesTeamMember
[**PostSalesSalesTeamsByParentIdMembers**](SalesTeamMembersAPI.md#PostSalesSalesTeamsByParentIdMembers) | **Post** /sales/salesTeams/{parentId}/members | Post SalesTeamMember
[**PutSalesSalesTeamsByParentIdMembersById**](SalesTeamMembersAPI.md#PutSalesSalesTeamsByParentIdMembersById) | **Put** /sales/salesTeams/{parentId}/members/{id} | Put SalesTeamMember



## DeleteSalesSalesTeamsByParentIdMembersById

> DeleteSalesSalesTeamsByParentIdMembersById(ctx, id, parentId).ClientId(clientId).Execute()

Delete SalesTeamMember

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
	id := int32(56) // int32 | memberId
	parentId := int32(56) // int32 | salesTeamId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SalesTeamMembersAPI.DeleteSalesSalesTeamsByParentIdMembersById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesTeamMembersAPI.DeleteSalesSalesTeamsByParentIdMembersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | memberId | 
**parentId** | **int32** | salesTeamId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSalesSalesTeamsByParentIdMembersByIdRequest struct via the builder pattern


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


## GetSalesSalesTeamsByParentIdMembers

> []SalesTeamMember GetSalesSalesTeamsByParentIdMembers(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of SalesTeamMember

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
	parentId := int32(56) // int32 | salesTeamId
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
	resp, r, err := apiClient.SalesTeamMembersAPI.GetSalesSalesTeamsByParentIdMembers(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesTeamMembersAPI.GetSalesSalesTeamsByParentIdMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesSalesTeamsByParentIdMembers`: []SalesTeamMember
	fmt.Fprintf(os.Stdout, "Response from `SalesTeamMembersAPI.GetSalesSalesTeamsByParentIdMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | salesTeamId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesSalesTeamsByParentIdMembersRequest struct via the builder pattern


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

[**[]SalesTeamMember**](SalesTeamMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalesSalesTeamsByParentIdMembersById

> SalesTeamMember GetSalesSalesTeamsByParentIdMembersById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get SalesTeamMember

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
	id := int32(56) // int32 | memberId
	parentId := int32(56) // int32 | salesTeamId
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
	resp, r, err := apiClient.SalesTeamMembersAPI.GetSalesSalesTeamsByParentIdMembersById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesTeamMembersAPI.GetSalesSalesTeamsByParentIdMembersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesSalesTeamsByParentIdMembersById`: SalesTeamMember
	fmt.Fprintf(os.Stdout, "Response from `SalesTeamMembersAPI.GetSalesSalesTeamsByParentIdMembersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | memberId | 
**parentId** | **int32** | salesTeamId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesSalesTeamsByParentIdMembersByIdRequest struct via the builder pattern


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

[**SalesTeamMember**](SalesTeamMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalesSalesTeamsByParentIdMembersCount

> Count GetSalesSalesTeamsByParentIdMembersCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of SalesTeamMember

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
	parentId := int32(56) // int32 | salesTeamId
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
	resp, r, err := apiClient.SalesTeamMembersAPI.GetSalesSalesTeamsByParentIdMembersCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesTeamMembersAPI.GetSalesSalesTeamsByParentIdMembersCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesSalesTeamsByParentIdMembersCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `SalesTeamMembersAPI.GetSalesSalesTeamsByParentIdMembersCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | salesTeamId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesSalesTeamsByParentIdMembersCountRequest struct via the builder pattern


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


## PatchSalesSalesTeamsByParentIdMembersById

> SalesTeamMember PatchSalesSalesTeamsByParentIdMembersById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch SalesTeamMember

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
	id := int32(56) // int32 | memberId
	parentId := int32(56) // int32 | salesTeamId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalesTeamMembersAPI.PatchSalesSalesTeamsByParentIdMembersById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesTeamMembersAPI.PatchSalesSalesTeamsByParentIdMembersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSalesSalesTeamsByParentIdMembersById`: SalesTeamMember
	fmt.Fprintf(os.Stdout, "Response from `SalesTeamMembersAPI.PatchSalesSalesTeamsByParentIdMembersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | memberId | 
**parentId** | **int32** | salesTeamId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSalesSalesTeamsByParentIdMembersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**SalesTeamMember**](SalesTeamMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSalesSalesTeamsByParentIdMembers

> SalesTeamMember PostSalesSalesTeamsByParentIdMembers(ctx, parentId).ClientId(clientId).SalesTeamMember(salesTeamMember).Execute()

Post SalesTeamMember

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
	parentId := int32(56) // int32 | salesTeamId
	clientId := "clientId_example" // string | 
	salesTeamMember := *openapiclient.NewSalesTeamMember(*openapiclient.NewMemberReference()) // SalesTeamMember | salesTeamMember

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalesTeamMembersAPI.PostSalesSalesTeamsByParentIdMembers(context.Background(), parentId).ClientId(clientId).SalesTeamMember(salesTeamMember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesTeamMembersAPI.PostSalesSalesTeamsByParentIdMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSalesSalesTeamsByParentIdMembers`: SalesTeamMember
	fmt.Fprintf(os.Stdout, "Response from `SalesTeamMembersAPI.PostSalesSalesTeamsByParentIdMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | salesTeamId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSalesSalesTeamsByParentIdMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **salesTeamMember** | [**SalesTeamMember**](SalesTeamMember.md) | salesTeamMember | 

### Return type

[**SalesTeamMember**](SalesTeamMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSalesSalesTeamsByParentIdMembersById

> SalesTeamMember PutSalesSalesTeamsByParentIdMembersById(ctx, id, parentId).ClientId(clientId).SalesTeamMember(salesTeamMember).Execute()

Put SalesTeamMember

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
	id := int32(56) // int32 | memberId
	parentId := int32(56) // int32 | salesTeamId
	clientId := "clientId_example" // string | 
	salesTeamMember := *openapiclient.NewSalesTeamMember(*openapiclient.NewMemberReference()) // SalesTeamMember | salesTeamMember

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalesTeamMembersAPI.PutSalesSalesTeamsByParentIdMembersById(context.Background(), id, parentId).ClientId(clientId).SalesTeamMember(salesTeamMember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesTeamMembersAPI.PutSalesSalesTeamsByParentIdMembersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSalesSalesTeamsByParentIdMembersById`: SalesTeamMember
	fmt.Fprintf(os.Stdout, "Response from `SalesTeamMembersAPI.PutSalesSalesTeamsByParentIdMembersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | memberId | 
**parentId** | **int32** | salesTeamId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSalesSalesTeamsByParentIdMembersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **salesTeamMember** | [**SalesTeamMember**](SalesTeamMember.md) | salesTeamMember | 

### Return type

[**SalesTeamMember**](SalesTeamMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

