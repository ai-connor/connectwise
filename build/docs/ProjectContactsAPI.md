# \ProjectContactsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProjectProjectsByParentIdContactsById**](ProjectContactsAPI.md#DeleteProjectProjectsByParentIdContactsById) | **Delete** /project/projects/{parentId}/contacts/{id} | Delete ProjectContact
[**GetProjectProjectsByParentIdContacts**](ProjectContactsAPI.md#GetProjectProjectsByParentIdContacts) | **Get** /project/projects/{parentId}/contacts | Get List of ProjectContact
[**GetProjectProjectsByParentIdContactsById**](ProjectContactsAPI.md#GetProjectProjectsByParentIdContactsById) | **Get** /project/projects/{parentId}/contacts/{id} | Get ProjectContact
[**PostProjectProjectsByParentIdContacts**](ProjectContactsAPI.md#PostProjectProjectsByParentIdContacts) | **Post** /project/projects/{parentId}/contacts | Post ProjectContact



## DeleteProjectProjectsByParentIdContactsById

> DeleteProjectProjectsByParentIdContactsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ProjectContact

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
	id := int32(56) // int32 | contactId
	parentId := int32(56) // int32 | projectId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectContactsAPI.DeleteProjectProjectsByParentIdContactsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectContactsAPI.DeleteProjectProjectsByParentIdContactsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | contactId | 
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectProjectsByParentIdContactsByIdRequest struct via the builder pattern


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


## GetProjectProjectsByParentIdContacts

> []ProjectContact GetProjectProjectsByParentIdContacts(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ProjectContact

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
	parentId := int32(56) // int32 | projectId
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
	resp, r, err := apiClient.ProjectContactsAPI.GetProjectProjectsByParentIdContacts(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectContactsAPI.GetProjectProjectsByParentIdContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectsByParentIdContacts`: []ProjectContact
	fmt.Fprintf(os.Stdout, "Response from `ProjectContactsAPI.GetProjectProjectsByParentIdContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectsByParentIdContactsRequest struct via the builder pattern


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

[**[]ProjectContact**](ProjectContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectProjectsByParentIdContactsById

> ProjectContact GetProjectProjectsByParentIdContactsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ProjectContact

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
	id := int32(56) // int32 | contactId
	parentId := int32(56) // int32 | projectId
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
	resp, r, err := apiClient.ProjectContactsAPI.GetProjectProjectsByParentIdContactsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectContactsAPI.GetProjectProjectsByParentIdContactsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProjectsByParentIdContactsById`: ProjectContact
	fmt.Fprintf(os.Stdout, "Response from `ProjectContactsAPI.GetProjectProjectsByParentIdContactsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | contactId | 
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectProjectsByParentIdContactsByIdRequest struct via the builder pattern


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

[**ProjectContact**](ProjectContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectProjectsByParentIdContacts

> ProjectContact PostProjectProjectsByParentIdContacts(ctx, parentId).ClientId(clientId).ProjectContact(projectContact).Execute()

Post ProjectContact

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
	parentId := int32(56) // int32 | projectId
	clientId := "clientId_example" // string | 
	projectContact := *openapiclient.NewProjectContact(*openapiclient.NewContactReference()) // ProjectContact | contact

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectContactsAPI.PostProjectProjectsByParentIdContacts(context.Background(), parentId).ClientId(clientId).ProjectContact(projectContact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectContactsAPI.PostProjectProjectsByParentIdContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectProjectsByParentIdContacts`: ProjectContact
	fmt.Fprintf(os.Stdout, "Response from `ProjectContactsAPI.PostProjectProjectsByParentIdContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectProjectsByParentIdContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **projectContact** | [**ProjectContact**](ProjectContact.md) | contact | 

### Return type

[**ProjectContact**](ProjectContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

