# \MemberDelegationsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemMembersByParentIdDelegationsById**](MemberDelegationsAPI.md#DeleteSystemMembersByParentIdDelegationsById) | **Delete** /system/members/{parentId}/delegations/{id} | Delete MemberDelegation
[**DeleteSystemMyAccountByParentIdDelegationsById**](MemberDelegationsAPI.md#DeleteSystemMyAccountByParentIdDelegationsById) | **Delete** /system/myAccount/{parentId}/delegations/{id} | Delete MemberDelegation
[**GetSystemMembersByParentIdDelegations**](MemberDelegationsAPI.md#GetSystemMembersByParentIdDelegations) | **Get** /system/members/{parentId}/delegations | Get List of MemberDelegation
[**GetSystemMembersByParentIdDelegationsById**](MemberDelegationsAPI.md#GetSystemMembersByParentIdDelegationsById) | **Get** /system/members/{parentId}/delegations/{id} | Get MemberDelegation
[**GetSystemMembersByParentIdDelegationsCount**](MemberDelegationsAPI.md#GetSystemMembersByParentIdDelegationsCount) | **Get** /system/members/{parentId}/delegations/count | Get Count of MemberDelegation
[**GetSystemMyAccountByParentIdDelegations**](MemberDelegationsAPI.md#GetSystemMyAccountByParentIdDelegations) | **Get** /system/myAccount/{parentId}/delegations | Get List of MemberDelegation
[**GetSystemMyAccountByParentIdDelegationsById**](MemberDelegationsAPI.md#GetSystemMyAccountByParentIdDelegationsById) | **Get** /system/myAccount/{parentId}/delegations/{id} | Get MemberDelegation
[**GetSystemMyAccountByParentIdDelegationsCount**](MemberDelegationsAPI.md#GetSystemMyAccountByParentIdDelegationsCount) | **Get** /system/myAccount/{parentId}/delegations/count | Get Count of MemberDelegation
[**PatchSystemMembersByParentIdDelegationsById**](MemberDelegationsAPI.md#PatchSystemMembersByParentIdDelegationsById) | **Patch** /system/members/{parentId}/delegations/{id} | Patch MemberDelegation
[**PatchSystemMyAccountByParentIdDelegationsById**](MemberDelegationsAPI.md#PatchSystemMyAccountByParentIdDelegationsById) | **Patch** /system/myAccount/{parentId}/delegations/{id} | Patch MemberDelegation
[**PostSystemMembersByParentIdDelegations**](MemberDelegationsAPI.md#PostSystemMembersByParentIdDelegations) | **Post** /system/members/{parentId}/delegations | Post MemberDelegation
[**PostSystemMyAccountByParentIdDelegations**](MemberDelegationsAPI.md#PostSystemMyAccountByParentIdDelegations) | **Post** /system/myAccount/{parentId}/delegations | Post MemberDelegation
[**PutSystemMembersByParentIdDelegationsById**](MemberDelegationsAPI.md#PutSystemMembersByParentIdDelegationsById) | **Put** /system/members/{parentId}/delegations/{id} | Put MemberDelegation
[**PutSystemMyAccountByParentIdDelegationsById**](MemberDelegationsAPI.md#PutSystemMyAccountByParentIdDelegationsById) | **Put** /system/myAccount/{parentId}/delegations/{id} | Put MemberDelegation



## DeleteSystemMembersByParentIdDelegationsById

> DeleteSystemMembersByParentIdDelegationsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete MemberDelegation

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
	id := int32(56) // int32 | delegationId
	parentId := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MemberDelegationsAPI.DeleteSystemMembersByParentIdDelegationsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberDelegationsAPI.DeleteSystemMembersByParentIdDelegationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | delegationId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemMembersByParentIdDelegationsByIdRequest struct via the builder pattern


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


## DeleteSystemMyAccountByParentIdDelegationsById

> DeleteSystemMyAccountByParentIdDelegationsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete MemberDelegation

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
	id := int32(56) // int32 | delegationId
	parentId := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MemberDelegationsAPI.DeleteSystemMyAccountByParentIdDelegationsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberDelegationsAPI.DeleteSystemMyAccountByParentIdDelegationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | delegationId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemMyAccountByParentIdDelegationsByIdRequest struct via the builder pattern


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


## GetSystemMembersByParentIdDelegations

> []MemberDelegation GetSystemMembersByParentIdDelegations(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of MemberDelegation

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
	parentId := int32(56) // int32 | memberId
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
	resp, r, err := apiClient.MemberDelegationsAPI.GetSystemMembersByParentIdDelegations(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberDelegationsAPI.GetSystemMembersByParentIdDelegations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMembersByParentIdDelegations`: []MemberDelegation
	fmt.Fprintf(os.Stdout, "Response from `MemberDelegationsAPI.GetSystemMembersByParentIdDelegations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMembersByParentIdDelegationsRequest struct via the builder pattern


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

[**[]MemberDelegation**](MemberDelegation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemMembersByParentIdDelegationsById

> MemberDelegation GetSystemMembersByParentIdDelegationsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get MemberDelegation

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
	id := int32(56) // int32 | delegationId
	parentId := int32(56) // int32 | memberId
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
	resp, r, err := apiClient.MemberDelegationsAPI.GetSystemMembersByParentIdDelegationsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberDelegationsAPI.GetSystemMembersByParentIdDelegationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMembersByParentIdDelegationsById`: MemberDelegation
	fmt.Fprintf(os.Stdout, "Response from `MemberDelegationsAPI.GetSystemMembersByParentIdDelegationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | delegationId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMembersByParentIdDelegationsByIdRequest struct via the builder pattern


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

[**MemberDelegation**](MemberDelegation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemMembersByParentIdDelegationsCount

> Count GetSystemMembersByParentIdDelegationsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of MemberDelegation

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
	parentId := int32(56) // int32 | memberId
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
	resp, r, err := apiClient.MemberDelegationsAPI.GetSystemMembersByParentIdDelegationsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberDelegationsAPI.GetSystemMembersByParentIdDelegationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMembersByParentIdDelegationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `MemberDelegationsAPI.GetSystemMembersByParentIdDelegationsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMembersByParentIdDelegationsCountRequest struct via the builder pattern


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


## GetSystemMyAccountByParentIdDelegations

> []MemberDelegation GetSystemMyAccountByParentIdDelegations(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of MemberDelegation

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
	parentId := int32(56) // int32 | memberId
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
	resp, r, err := apiClient.MemberDelegationsAPI.GetSystemMyAccountByParentIdDelegations(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberDelegationsAPI.GetSystemMyAccountByParentIdDelegations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMyAccountByParentIdDelegations`: []MemberDelegation
	fmt.Fprintf(os.Stdout, "Response from `MemberDelegationsAPI.GetSystemMyAccountByParentIdDelegations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMyAccountByParentIdDelegationsRequest struct via the builder pattern


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

[**[]MemberDelegation**](MemberDelegation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemMyAccountByParentIdDelegationsById

> MemberDelegation GetSystemMyAccountByParentIdDelegationsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get MemberDelegation

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
	id := int32(56) // int32 | delegationId
	parentId := int32(56) // int32 | memberId
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
	resp, r, err := apiClient.MemberDelegationsAPI.GetSystemMyAccountByParentIdDelegationsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberDelegationsAPI.GetSystemMyAccountByParentIdDelegationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMyAccountByParentIdDelegationsById`: MemberDelegation
	fmt.Fprintf(os.Stdout, "Response from `MemberDelegationsAPI.GetSystemMyAccountByParentIdDelegationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | delegationId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMyAccountByParentIdDelegationsByIdRequest struct via the builder pattern


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

[**MemberDelegation**](MemberDelegation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemMyAccountByParentIdDelegationsCount

> Count GetSystemMyAccountByParentIdDelegationsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of MemberDelegation

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
	parentId := int32(56) // int32 | memberId
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
	resp, r, err := apiClient.MemberDelegationsAPI.GetSystemMyAccountByParentIdDelegationsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberDelegationsAPI.GetSystemMyAccountByParentIdDelegationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMyAccountByParentIdDelegationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `MemberDelegationsAPI.GetSystemMyAccountByParentIdDelegationsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMyAccountByParentIdDelegationsCountRequest struct via the builder pattern


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


## PatchSystemMembersByParentIdDelegationsById

> MemberDelegation PatchSystemMembersByParentIdDelegationsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch MemberDelegation

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
	id := int32(56) // int32 | delegationId
	parentId := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberDelegationsAPI.PatchSystemMembersByParentIdDelegationsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberDelegationsAPI.PatchSystemMembersByParentIdDelegationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemMembersByParentIdDelegationsById`: MemberDelegation
	fmt.Fprintf(os.Stdout, "Response from `MemberDelegationsAPI.PatchSystemMembersByParentIdDelegationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | delegationId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemMembersByParentIdDelegationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**MemberDelegation**](MemberDelegation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSystemMyAccountByParentIdDelegationsById

> MemberDelegation PatchSystemMyAccountByParentIdDelegationsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch MemberDelegation

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
	id := int32(56) // int32 | delegationId
	parentId := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberDelegationsAPI.PatchSystemMyAccountByParentIdDelegationsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberDelegationsAPI.PatchSystemMyAccountByParentIdDelegationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemMyAccountByParentIdDelegationsById`: MemberDelegation
	fmt.Fprintf(os.Stdout, "Response from `MemberDelegationsAPI.PatchSystemMyAccountByParentIdDelegationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | delegationId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemMyAccountByParentIdDelegationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**MemberDelegation**](MemberDelegation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemMembersByParentIdDelegations

> MemberDelegation PostSystemMembersByParentIdDelegations(ctx, parentId).ClientId(clientId).MemberDelegation(memberDelegation).Execute()

Post MemberDelegation

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ai-connor/connectwise"
)

func main() {
	parentId := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 
	memberDelegation := *openapiclient.NewMemberDelegation("DelegationType_example", *openapiclient.NewMemberReference(), time.Now(), time.Now()) // MemberDelegation | memberDelegation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberDelegationsAPI.PostSystemMembersByParentIdDelegations(context.Background(), parentId).ClientId(clientId).MemberDelegation(memberDelegation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberDelegationsAPI.PostSystemMembersByParentIdDelegations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemMembersByParentIdDelegations`: MemberDelegation
	fmt.Fprintf(os.Stdout, "Response from `MemberDelegationsAPI.PostSystemMembersByParentIdDelegations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemMembersByParentIdDelegationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **memberDelegation** | [**MemberDelegation**](MemberDelegation.md) | memberDelegation | 

### Return type

[**MemberDelegation**](MemberDelegation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemMyAccountByParentIdDelegations

> MemberDelegation PostSystemMyAccountByParentIdDelegations(ctx, parentId).ClientId(clientId).MemberDelegation(memberDelegation).Execute()

Post MemberDelegation

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ai-connor/connectwise"
)

func main() {
	parentId := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 
	memberDelegation := *openapiclient.NewMemberDelegation("DelegationType_example", *openapiclient.NewMemberReference(), time.Now(), time.Now()) // MemberDelegation | memberDelegation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberDelegationsAPI.PostSystemMyAccountByParentIdDelegations(context.Background(), parentId).ClientId(clientId).MemberDelegation(memberDelegation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberDelegationsAPI.PostSystemMyAccountByParentIdDelegations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemMyAccountByParentIdDelegations`: MemberDelegation
	fmt.Fprintf(os.Stdout, "Response from `MemberDelegationsAPI.PostSystemMyAccountByParentIdDelegations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemMyAccountByParentIdDelegationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **memberDelegation** | [**MemberDelegation**](MemberDelegation.md) | memberDelegation | 

### Return type

[**MemberDelegation**](MemberDelegation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemMembersByParentIdDelegationsById

> MemberDelegation PutSystemMembersByParentIdDelegationsById(ctx, id, parentId).ClientId(clientId).MemberDelegation(memberDelegation).Execute()

Put MemberDelegation

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ai-connor/connectwise"
)

func main() {
	id := int32(56) // int32 | delegationId
	parentId := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 
	memberDelegation := *openapiclient.NewMemberDelegation("DelegationType_example", *openapiclient.NewMemberReference(), time.Now(), time.Now()) // MemberDelegation | memberDelegation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberDelegationsAPI.PutSystemMembersByParentIdDelegationsById(context.Background(), id, parentId).ClientId(clientId).MemberDelegation(memberDelegation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberDelegationsAPI.PutSystemMembersByParentIdDelegationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemMembersByParentIdDelegationsById`: MemberDelegation
	fmt.Fprintf(os.Stdout, "Response from `MemberDelegationsAPI.PutSystemMembersByParentIdDelegationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | delegationId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemMembersByParentIdDelegationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **memberDelegation** | [**MemberDelegation**](MemberDelegation.md) | memberDelegation | 

### Return type

[**MemberDelegation**](MemberDelegation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemMyAccountByParentIdDelegationsById

> MemberDelegation PutSystemMyAccountByParentIdDelegationsById(ctx, id, parentId).ClientId(clientId).MemberDelegation(memberDelegation).Execute()

Put MemberDelegation

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ai-connor/connectwise"
)

func main() {
	id := int32(56) // int32 | delegationId
	parentId := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 
	memberDelegation := *openapiclient.NewMemberDelegation("DelegationType_example", *openapiclient.NewMemberReference(), time.Now(), time.Now()) // MemberDelegation | memberDelegation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberDelegationsAPI.PutSystemMyAccountByParentIdDelegationsById(context.Background(), id, parentId).ClientId(clientId).MemberDelegation(memberDelegation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberDelegationsAPI.PutSystemMyAccountByParentIdDelegationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemMyAccountByParentIdDelegationsById`: MemberDelegation
	fmt.Fprintf(os.Stdout, "Response from `MemberDelegationsAPI.PutSystemMyAccountByParentIdDelegationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | delegationId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemMyAccountByParentIdDelegationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **memberDelegation** | [**MemberDelegation**](MemberDelegation.md) | memberDelegation | 

### Return type

[**MemberDelegation**](MemberDelegation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

