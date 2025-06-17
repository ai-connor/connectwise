# \BoardExcludedMembersAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteServiceBoardsByParentIdExcludedMembersById**](BoardExcludedMembersAPI.md#DeleteServiceBoardsByParentIdExcludedMembersById) | **Delete** /service/boards/{parentId}/excludedMembers/{id} | Delete BoardExcludedMember
[**GetServiceBoardsByParentIdExcludedMembers**](BoardExcludedMembersAPI.md#GetServiceBoardsByParentIdExcludedMembers) | **Get** /service/boards/{parentId}/excludedMembers | Get List of BoardExcludedMember
[**GetServiceBoardsByParentIdExcludedMembersById**](BoardExcludedMembersAPI.md#GetServiceBoardsByParentIdExcludedMembersById) | **Get** /service/boards/{parentId}/excludedMembers/{id} | Get BoardExcludedMember
[**GetServiceBoardsByParentIdExcludedMembersCount**](BoardExcludedMembersAPI.md#GetServiceBoardsByParentIdExcludedMembersCount) | **Get** /service/boards/{parentId}/excludedMembers/count | Get Count of BoardExcludedMember
[**PostServiceBoardsByParentIdExcludedMembers**](BoardExcludedMembersAPI.md#PostServiceBoardsByParentIdExcludedMembers) | **Post** /service/boards/{parentId}/excludedMembers | Post BoardExcludedMember



## DeleteServiceBoardsByParentIdExcludedMembersById

> DeleteServiceBoardsByParentIdExcludedMembersById(ctx, id, parentId).ClientId(clientId).Execute()

Delete BoardExcludedMember

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
	id := int32(56) // int32 | excludedMemberId
	parentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BoardExcludedMembersAPI.DeleteServiceBoardsByParentIdExcludedMembersById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardExcludedMembersAPI.DeleteServiceBoardsByParentIdExcludedMembersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | excludedMemberId | 
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceBoardsByParentIdExcludedMembersByIdRequest struct via the builder pattern


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


## GetServiceBoardsByParentIdExcludedMembers

> []BoardExcludedMember GetServiceBoardsByParentIdExcludedMembers(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of BoardExcludedMember

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
	parentId := int32(56) // int32 | boardId
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
	resp, r, err := apiClient.BoardExcludedMembersAPI.GetServiceBoardsByParentIdExcludedMembers(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardExcludedMembersAPI.GetServiceBoardsByParentIdExcludedMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceBoardsByParentIdExcludedMembers`: []BoardExcludedMember
	fmt.Fprintf(os.Stdout, "Response from `BoardExcludedMembersAPI.GetServiceBoardsByParentIdExcludedMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceBoardsByParentIdExcludedMembersRequest struct via the builder pattern


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

[**[]BoardExcludedMember**](BoardExcludedMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceBoardsByParentIdExcludedMembersById

> BoardExcludedMember GetServiceBoardsByParentIdExcludedMembersById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get BoardExcludedMember

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
	id := int32(56) // int32 | excludedMemberId
	parentId := int32(56) // int32 | boardId
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
	resp, r, err := apiClient.BoardExcludedMembersAPI.GetServiceBoardsByParentIdExcludedMembersById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardExcludedMembersAPI.GetServiceBoardsByParentIdExcludedMembersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceBoardsByParentIdExcludedMembersById`: BoardExcludedMember
	fmt.Fprintf(os.Stdout, "Response from `BoardExcludedMembersAPI.GetServiceBoardsByParentIdExcludedMembersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | excludedMemberId | 
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceBoardsByParentIdExcludedMembersByIdRequest struct via the builder pattern


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

[**BoardExcludedMember**](BoardExcludedMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceBoardsByParentIdExcludedMembersCount

> Count GetServiceBoardsByParentIdExcludedMembersCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of BoardExcludedMember

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
	parentId := int32(56) // int32 | boardId
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
	resp, r, err := apiClient.BoardExcludedMembersAPI.GetServiceBoardsByParentIdExcludedMembersCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardExcludedMembersAPI.GetServiceBoardsByParentIdExcludedMembersCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceBoardsByParentIdExcludedMembersCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `BoardExcludedMembersAPI.GetServiceBoardsByParentIdExcludedMembersCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceBoardsByParentIdExcludedMembersCountRequest struct via the builder pattern


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


## PostServiceBoardsByParentIdExcludedMembers

> BoardExcludedMember PostServiceBoardsByParentIdExcludedMembers(ctx, parentId).ClientId(clientId).BoardExcludedMember(boardExcludedMember).Execute()

Post BoardExcludedMember

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
	parentId := int32(56) // int32 | boardId
	clientId := "clientId_example" // string | 
	boardExcludedMember := *openapiclient.NewBoardExcludedMember() // BoardExcludedMember | boardExcludedMember

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardExcludedMembersAPI.PostServiceBoardsByParentIdExcludedMembers(context.Background(), parentId).ClientId(clientId).BoardExcludedMember(boardExcludedMember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardExcludedMembersAPI.PostServiceBoardsByParentIdExcludedMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServiceBoardsByParentIdExcludedMembers`: BoardExcludedMember
	fmt.Fprintf(os.Stdout, "Response from `BoardExcludedMembersAPI.PostServiceBoardsByParentIdExcludedMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | boardId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostServiceBoardsByParentIdExcludedMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **boardExcludedMember** | [**BoardExcludedMember**](BoardExcludedMember.md) | boardExcludedMember | 

### Return type

[**BoardExcludedMember**](BoardExcludedMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

