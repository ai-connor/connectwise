# \ContactGroupsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyContactsByParentIdGroupsById**](ContactGroupsAPI.md#DeleteCompanyContactsByParentIdGroupsById) | **Delete** /company/contacts/{parentId}/groups/{id} | Delete ContactGroup
[**GetCompanyContactsByParentIdGroups**](ContactGroupsAPI.md#GetCompanyContactsByParentIdGroups) | **Get** /company/contacts/{parentId}/groups | Get List of ContactGroup
[**GetCompanyContactsByParentIdGroupsById**](ContactGroupsAPI.md#GetCompanyContactsByParentIdGroupsById) | **Get** /company/contacts/{parentId}/groups/{id} | Get ContactGroup
[**GetCompanyContactsByParentIdGroupsCount**](ContactGroupsAPI.md#GetCompanyContactsByParentIdGroupsCount) | **Get** /company/contacts/{parentId}/groups/count | Get Count of ContactGroup
[**PatchCompanyContactsByParentIdGroupsById**](ContactGroupsAPI.md#PatchCompanyContactsByParentIdGroupsById) | **Patch** /company/contacts/{parentId}/groups/{id} | Patch ContactGroup
[**PostCompanyContactsByParentIdGroups**](ContactGroupsAPI.md#PostCompanyContactsByParentIdGroups) | **Post** /company/contacts/{parentId}/groups | Post ContactGroup
[**PutCompanyContactsByParentIdGroupsById**](ContactGroupsAPI.md#PutCompanyContactsByParentIdGroupsById) | **Put** /company/contacts/{parentId}/groups/{id} | Put ContactGroup



## DeleteCompanyContactsByParentIdGroupsById

> DeleteCompanyContactsByParentIdGroupsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ContactGroup

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
	id := int32(56) // int32 | groupId
	parentId := int32(56) // int32 | contactId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContactGroupsAPI.DeleteCompanyContactsByParentIdGroupsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactGroupsAPI.DeleteCompanyContactsByParentIdGroupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | groupId | 
**parentId** | **int32** | contactId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyContactsByParentIdGroupsByIdRequest struct via the builder pattern


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


## GetCompanyContactsByParentIdGroups

> []ContactGroup GetCompanyContactsByParentIdGroups(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ContactGroup

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
	parentId := int32(56) // int32 | contactId
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
	resp, r, err := apiClient.ContactGroupsAPI.GetCompanyContactsByParentIdGroups(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactGroupsAPI.GetCompanyContactsByParentIdGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyContactsByParentIdGroups`: []ContactGroup
	fmt.Fprintf(os.Stdout, "Response from `ContactGroupsAPI.GetCompanyContactsByParentIdGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | contactId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyContactsByParentIdGroupsRequest struct via the builder pattern


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

[**[]ContactGroup**](ContactGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyContactsByParentIdGroupsById

> ContactGroup GetCompanyContactsByParentIdGroupsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ContactGroup

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
	id := int32(56) // int32 | groupId
	parentId := int32(56) // int32 | contactId
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
	resp, r, err := apiClient.ContactGroupsAPI.GetCompanyContactsByParentIdGroupsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactGroupsAPI.GetCompanyContactsByParentIdGroupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyContactsByParentIdGroupsById`: ContactGroup
	fmt.Fprintf(os.Stdout, "Response from `ContactGroupsAPI.GetCompanyContactsByParentIdGroupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | groupId | 
**parentId** | **int32** | contactId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyContactsByParentIdGroupsByIdRequest struct via the builder pattern


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

[**ContactGroup**](ContactGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyContactsByParentIdGroupsCount

> Count GetCompanyContactsByParentIdGroupsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ContactGroup

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
	parentId := int32(56) // int32 | contactId
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
	resp, r, err := apiClient.ContactGroupsAPI.GetCompanyContactsByParentIdGroupsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactGroupsAPI.GetCompanyContactsByParentIdGroupsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyContactsByParentIdGroupsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ContactGroupsAPI.GetCompanyContactsByParentIdGroupsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | contactId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyContactsByParentIdGroupsCountRequest struct via the builder pattern


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


## PatchCompanyContactsByParentIdGroupsById

> ContactGroup PatchCompanyContactsByParentIdGroupsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ContactGroup

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
	id := int32(56) // int32 | groupId
	parentId := int32(56) // int32 | contactId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactGroupsAPI.PatchCompanyContactsByParentIdGroupsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactGroupsAPI.PatchCompanyContactsByParentIdGroupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyContactsByParentIdGroupsById`: ContactGroup
	fmt.Fprintf(os.Stdout, "Response from `ContactGroupsAPI.PatchCompanyContactsByParentIdGroupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | groupId | 
**parentId** | **int32** | contactId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyContactsByParentIdGroupsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ContactGroup**](ContactGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyContactsByParentIdGroups

> ContactGroup PostCompanyContactsByParentIdGroups(ctx, parentId).ClientId(clientId).ContactGroup(contactGroup).Execute()

Post ContactGroup

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
	parentId := int32(56) // int32 | contactId
	clientId := "clientId_example" // string | 
	contactGroup := *openapiclient.NewContactGroup(*openapiclient.NewGroupReference()) // ContactGroup | contactGroup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactGroupsAPI.PostCompanyContactsByParentIdGroups(context.Background(), parentId).ClientId(clientId).ContactGroup(contactGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactGroupsAPI.PostCompanyContactsByParentIdGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyContactsByParentIdGroups`: ContactGroup
	fmt.Fprintf(os.Stdout, "Response from `ContactGroupsAPI.PostCompanyContactsByParentIdGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | contactId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyContactsByParentIdGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **contactGroup** | [**ContactGroup**](ContactGroup.md) | contactGroup | 

### Return type

[**ContactGroup**](ContactGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyContactsByParentIdGroupsById

> ContactGroup PutCompanyContactsByParentIdGroupsById(ctx, id, parentId).ClientId(clientId).ContactGroup(contactGroup).Execute()

Put ContactGroup

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
	id := int32(56) // int32 | groupId
	parentId := int32(56) // int32 | contactId
	clientId := "clientId_example" // string | 
	contactGroup := *openapiclient.NewContactGroup(*openapiclient.NewGroupReference()) // ContactGroup | contactGroup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactGroupsAPI.PutCompanyContactsByParentIdGroupsById(context.Background(), id, parentId).ClientId(clientId).ContactGroup(contactGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactGroupsAPI.PutCompanyContactsByParentIdGroupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyContactsByParentIdGroupsById`: ContactGroup
	fmt.Fprintf(os.Stdout, "Response from `ContactGroupsAPI.PutCompanyContactsByParentIdGroupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | groupId | 
**parentId** | **int32** | contactId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyContactsByParentIdGroupsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **contactGroup** | [**ContactGroup**](ContactGroup.md) | contactGroup | 

### Return type

[**ContactGroup**](ContactGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

