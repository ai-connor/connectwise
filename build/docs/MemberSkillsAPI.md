# \MemberSkillsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemMembersByParentIdSkillsById**](MemberSkillsAPI.md#DeleteSystemMembersByParentIdSkillsById) | **Delete** /system/members/{parentId}/skills/{id} | Delete MemberSkill
[**DeleteSystemMyAccountByParentIdSkillsById**](MemberSkillsAPI.md#DeleteSystemMyAccountByParentIdSkillsById) | **Delete** /system/myAccount/{parentId}/skills/{id} | Delete MemberSkill
[**GetSystemMembersByParentIdSkills**](MemberSkillsAPI.md#GetSystemMembersByParentIdSkills) | **Get** /system/members/{parentId}/skills | Get List of MemberSkill
[**GetSystemMembersByParentIdSkillsById**](MemberSkillsAPI.md#GetSystemMembersByParentIdSkillsById) | **Get** /system/members/{parentId}/skills/{id} | Get MemberSkill
[**GetSystemMembersByParentIdSkillsCount**](MemberSkillsAPI.md#GetSystemMembersByParentIdSkillsCount) | **Get** /system/members/{parentId}/skills/count | Get Count of MemberSkill
[**GetSystemMyAccountByParentIdSkills**](MemberSkillsAPI.md#GetSystemMyAccountByParentIdSkills) | **Get** /system/myAccount/{parentId}/skills | Get List of MemberSkill
[**GetSystemMyAccountByParentIdSkillsById**](MemberSkillsAPI.md#GetSystemMyAccountByParentIdSkillsById) | **Get** /system/myAccount/{parentId}/skills/{id} | Get MemberSkill
[**GetSystemMyAccountByParentIdSkillsCount**](MemberSkillsAPI.md#GetSystemMyAccountByParentIdSkillsCount) | **Get** /system/myAccount/{parentId}/skills/count | Get Count of MemberSkill
[**PatchSystemMembersByParentIdSkillsById**](MemberSkillsAPI.md#PatchSystemMembersByParentIdSkillsById) | **Patch** /system/members/{parentId}/skills/{id} | Patch MemberSkill
[**PatchSystemMyAccountByParentIdSkillsById**](MemberSkillsAPI.md#PatchSystemMyAccountByParentIdSkillsById) | **Patch** /system/myAccount/{parentId}/skills/{id} | Patch MemberSkill
[**PostSystemMembersByParentIdSkills**](MemberSkillsAPI.md#PostSystemMembersByParentIdSkills) | **Post** /system/members/{parentId}/skills | Post MemberSkill
[**PostSystemMyAccountByParentIdSkills**](MemberSkillsAPI.md#PostSystemMyAccountByParentIdSkills) | **Post** /system/myAccount/{parentId}/skills | Post MemberSkill
[**PutSystemMembersByParentIdSkillsById**](MemberSkillsAPI.md#PutSystemMembersByParentIdSkillsById) | **Put** /system/members/{parentId}/skills/{id} | Put MemberSkill
[**PutSystemMyAccountByParentIdSkillsById**](MemberSkillsAPI.md#PutSystemMyAccountByParentIdSkillsById) | **Put** /system/myAccount/{parentId}/skills/{id} | Put MemberSkill



## DeleteSystemMembersByParentIdSkillsById

> DeleteSystemMembersByParentIdSkillsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete MemberSkill

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
	id := int32(56) // int32 | skillId
	parentId := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MemberSkillsAPI.DeleteSystemMembersByParentIdSkillsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberSkillsAPI.DeleteSystemMembersByParentIdSkillsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | skillId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemMembersByParentIdSkillsByIdRequest struct via the builder pattern


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


## DeleteSystemMyAccountByParentIdSkillsById

> DeleteSystemMyAccountByParentIdSkillsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete MemberSkill

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
	id := int32(56) // int32 | skillId
	parentId := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MemberSkillsAPI.DeleteSystemMyAccountByParentIdSkillsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberSkillsAPI.DeleteSystemMyAccountByParentIdSkillsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | skillId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemMyAccountByParentIdSkillsByIdRequest struct via the builder pattern


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


## GetSystemMembersByParentIdSkills

> []MemberSkill GetSystemMembersByParentIdSkills(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of MemberSkill

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
	resp, r, err := apiClient.MemberSkillsAPI.GetSystemMembersByParentIdSkills(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberSkillsAPI.GetSystemMembersByParentIdSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMembersByParentIdSkills`: []MemberSkill
	fmt.Fprintf(os.Stdout, "Response from `MemberSkillsAPI.GetSystemMembersByParentIdSkills`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMembersByParentIdSkillsRequest struct via the builder pattern


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

[**[]MemberSkill**](MemberSkill.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemMembersByParentIdSkillsById

> MemberSkill GetSystemMembersByParentIdSkillsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get MemberSkill

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
	id := int32(56) // int32 | skillId
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
	resp, r, err := apiClient.MemberSkillsAPI.GetSystemMembersByParentIdSkillsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberSkillsAPI.GetSystemMembersByParentIdSkillsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMembersByParentIdSkillsById`: MemberSkill
	fmt.Fprintf(os.Stdout, "Response from `MemberSkillsAPI.GetSystemMembersByParentIdSkillsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | skillId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMembersByParentIdSkillsByIdRequest struct via the builder pattern


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

[**MemberSkill**](MemberSkill.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemMembersByParentIdSkillsCount

> Count GetSystemMembersByParentIdSkillsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of MemberSkill

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
	resp, r, err := apiClient.MemberSkillsAPI.GetSystemMembersByParentIdSkillsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberSkillsAPI.GetSystemMembersByParentIdSkillsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMembersByParentIdSkillsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `MemberSkillsAPI.GetSystemMembersByParentIdSkillsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMembersByParentIdSkillsCountRequest struct via the builder pattern


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


## GetSystemMyAccountByParentIdSkills

> []MemberSkill GetSystemMyAccountByParentIdSkills(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of MemberSkill

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
	resp, r, err := apiClient.MemberSkillsAPI.GetSystemMyAccountByParentIdSkills(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberSkillsAPI.GetSystemMyAccountByParentIdSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMyAccountByParentIdSkills`: []MemberSkill
	fmt.Fprintf(os.Stdout, "Response from `MemberSkillsAPI.GetSystemMyAccountByParentIdSkills`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMyAccountByParentIdSkillsRequest struct via the builder pattern


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

[**[]MemberSkill**](MemberSkill.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemMyAccountByParentIdSkillsById

> MemberSkill GetSystemMyAccountByParentIdSkillsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get MemberSkill

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
	id := int32(56) // int32 | skillId
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
	resp, r, err := apiClient.MemberSkillsAPI.GetSystemMyAccountByParentIdSkillsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberSkillsAPI.GetSystemMyAccountByParentIdSkillsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMyAccountByParentIdSkillsById`: MemberSkill
	fmt.Fprintf(os.Stdout, "Response from `MemberSkillsAPI.GetSystemMyAccountByParentIdSkillsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | skillId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMyAccountByParentIdSkillsByIdRequest struct via the builder pattern


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

[**MemberSkill**](MemberSkill.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemMyAccountByParentIdSkillsCount

> Count GetSystemMyAccountByParentIdSkillsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of MemberSkill

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
	resp, r, err := apiClient.MemberSkillsAPI.GetSystemMyAccountByParentIdSkillsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberSkillsAPI.GetSystemMyAccountByParentIdSkillsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMyAccountByParentIdSkillsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `MemberSkillsAPI.GetSystemMyAccountByParentIdSkillsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMyAccountByParentIdSkillsCountRequest struct via the builder pattern


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


## PatchSystemMembersByParentIdSkillsById

> MemberSkill PatchSystemMembersByParentIdSkillsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch MemberSkill

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
	id := int32(56) // int32 | skillId
	parentId := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberSkillsAPI.PatchSystemMembersByParentIdSkillsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberSkillsAPI.PatchSystemMembersByParentIdSkillsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemMembersByParentIdSkillsById`: MemberSkill
	fmt.Fprintf(os.Stdout, "Response from `MemberSkillsAPI.PatchSystemMembersByParentIdSkillsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | skillId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemMembersByParentIdSkillsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**MemberSkill**](MemberSkill.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSystemMyAccountByParentIdSkillsById

> MemberSkill PatchSystemMyAccountByParentIdSkillsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch MemberSkill

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
	id := int32(56) // int32 | skillId
	parentId := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberSkillsAPI.PatchSystemMyAccountByParentIdSkillsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberSkillsAPI.PatchSystemMyAccountByParentIdSkillsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemMyAccountByParentIdSkillsById`: MemberSkill
	fmt.Fprintf(os.Stdout, "Response from `MemberSkillsAPI.PatchSystemMyAccountByParentIdSkillsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | skillId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemMyAccountByParentIdSkillsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**MemberSkill**](MemberSkill.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemMembersByParentIdSkills

> MemberSkill PostSystemMembersByParentIdSkills(ctx, parentId).ClientId(clientId).MemberSkill(memberSkill).Execute()

Post MemberSkill

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
	memberSkill := *openapiclient.NewMemberSkill(*openapiclient.NewSkillReference(), "SkillLevel_example") // MemberSkill | memberSkill

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberSkillsAPI.PostSystemMembersByParentIdSkills(context.Background(), parentId).ClientId(clientId).MemberSkill(memberSkill).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberSkillsAPI.PostSystemMembersByParentIdSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemMembersByParentIdSkills`: MemberSkill
	fmt.Fprintf(os.Stdout, "Response from `MemberSkillsAPI.PostSystemMembersByParentIdSkills`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemMembersByParentIdSkillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **memberSkill** | [**MemberSkill**](MemberSkill.md) | memberSkill | 

### Return type

[**MemberSkill**](MemberSkill.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemMyAccountByParentIdSkills

> MemberSkill PostSystemMyAccountByParentIdSkills(ctx, parentId).ClientId(clientId).MemberSkill(memberSkill).Execute()

Post MemberSkill

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
	memberSkill := *openapiclient.NewMemberSkill(*openapiclient.NewSkillReference(), "SkillLevel_example") // MemberSkill | memberSkill

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberSkillsAPI.PostSystemMyAccountByParentIdSkills(context.Background(), parentId).ClientId(clientId).MemberSkill(memberSkill).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberSkillsAPI.PostSystemMyAccountByParentIdSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemMyAccountByParentIdSkills`: MemberSkill
	fmt.Fprintf(os.Stdout, "Response from `MemberSkillsAPI.PostSystemMyAccountByParentIdSkills`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemMyAccountByParentIdSkillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **memberSkill** | [**MemberSkill**](MemberSkill.md) | memberSkill | 

### Return type

[**MemberSkill**](MemberSkill.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemMembersByParentIdSkillsById

> MemberSkill PutSystemMembersByParentIdSkillsById(ctx, id, parentId).ClientId(clientId).MemberSkill(memberSkill).Execute()

Put MemberSkill

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
	id := int32(56) // int32 | skillId
	parentId := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 
	memberSkill := *openapiclient.NewMemberSkill(*openapiclient.NewSkillReference(), "SkillLevel_example") // MemberSkill | memberSkill

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberSkillsAPI.PutSystemMembersByParentIdSkillsById(context.Background(), id, parentId).ClientId(clientId).MemberSkill(memberSkill).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberSkillsAPI.PutSystemMembersByParentIdSkillsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemMembersByParentIdSkillsById`: MemberSkill
	fmt.Fprintf(os.Stdout, "Response from `MemberSkillsAPI.PutSystemMembersByParentIdSkillsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | skillId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemMembersByParentIdSkillsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **memberSkill** | [**MemberSkill**](MemberSkill.md) | memberSkill | 

### Return type

[**MemberSkill**](MemberSkill.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemMyAccountByParentIdSkillsById

> MemberSkill PutSystemMyAccountByParentIdSkillsById(ctx, id, parentId).ClientId(clientId).MemberSkill(memberSkill).Execute()

Put MemberSkill

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
	id := int32(56) // int32 | skillId
	parentId := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 
	memberSkill := *openapiclient.NewMemberSkill(*openapiclient.NewSkillReference(), "SkillLevel_example") // MemberSkill | memberSkill

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberSkillsAPI.PutSystemMyAccountByParentIdSkillsById(context.Background(), id, parentId).ClientId(clientId).MemberSkill(memberSkill).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberSkillsAPI.PutSystemMyAccountByParentIdSkillsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemMyAccountByParentIdSkillsById`: MemberSkill
	fmt.Fprintf(os.Stdout, "Response from `MemberSkillsAPI.PutSystemMyAccountByParentIdSkillsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | skillId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemMyAccountByParentIdSkillsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **memberSkill** | [**MemberSkill**](MemberSkill.md) | memberSkill | 

### Return type

[**MemberSkill**](MemberSkill.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

