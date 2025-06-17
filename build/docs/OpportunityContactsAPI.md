# \OpportunityContactsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSalesOpportunitiesByParentIdContactsById**](OpportunityContactsAPI.md#DeleteSalesOpportunitiesByParentIdContactsById) | **Delete** /sales/opportunities/{parentId}/contacts/{id} | Delete OpportunityContact
[**GetSalesOpportunitiesByParentIdContacts**](OpportunityContactsAPI.md#GetSalesOpportunitiesByParentIdContacts) | **Get** /sales/opportunities/{parentId}/contacts | Get List of OpportunityContact
[**GetSalesOpportunitiesByParentIdContactsById**](OpportunityContactsAPI.md#GetSalesOpportunitiesByParentIdContactsById) | **Get** /sales/opportunities/{parentId}/contacts/{id} | Get OpportunityContact
[**GetSalesOpportunitiesByParentIdContactsCount**](OpportunityContactsAPI.md#GetSalesOpportunitiesByParentIdContactsCount) | **Get** /sales/opportunities/{parentId}/contacts/count | Get Count of OpportunityContact
[**PatchSalesOpportunitiesByParentIdContactsById**](OpportunityContactsAPI.md#PatchSalesOpportunitiesByParentIdContactsById) | **Patch** /sales/opportunities/{parentId}/contacts/{id} | Patch OpportunityContact
[**PostSalesOpportunitiesByParentIdContacts**](OpportunityContactsAPI.md#PostSalesOpportunitiesByParentIdContacts) | **Post** /sales/opportunities/{parentId}/contacts | Post OpportunityContact
[**PutSalesOpportunitiesByParentIdContactsById**](OpportunityContactsAPI.md#PutSalesOpportunitiesByParentIdContactsById) | **Put** /sales/opportunities/{parentId}/contacts/{id} | Put OpportunityContact



## DeleteSalesOpportunitiesByParentIdContactsById

> DeleteSalesOpportunitiesByParentIdContactsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete OpportunityContact

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
	parentId := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OpportunityContactsAPI.DeleteSalesOpportunitiesByParentIdContactsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityContactsAPI.DeleteSalesOpportunitiesByParentIdContactsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | contactId | 
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSalesOpportunitiesByParentIdContactsByIdRequest struct via the builder pattern


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


## GetSalesOpportunitiesByParentIdContacts

> []OpportunityContact GetSalesOpportunitiesByParentIdContacts(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of OpportunityContact

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
	parentId := int32(56) // int32 | opportunityId
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
	resp, r, err := apiClient.OpportunityContactsAPI.GetSalesOpportunitiesByParentIdContacts(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityContactsAPI.GetSalesOpportunitiesByParentIdContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesOpportunitiesByParentIdContacts`: []OpportunityContact
	fmt.Fprintf(os.Stdout, "Response from `OpportunityContactsAPI.GetSalesOpportunitiesByParentIdContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesOpportunitiesByParentIdContactsRequest struct via the builder pattern


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

[**[]OpportunityContact**](OpportunityContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalesOpportunitiesByParentIdContactsById

> OpportunityContact GetSalesOpportunitiesByParentIdContactsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get OpportunityContact

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
	parentId := int32(56) // int32 | opportunityId
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
	resp, r, err := apiClient.OpportunityContactsAPI.GetSalesOpportunitiesByParentIdContactsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityContactsAPI.GetSalesOpportunitiesByParentIdContactsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesOpportunitiesByParentIdContactsById`: OpportunityContact
	fmt.Fprintf(os.Stdout, "Response from `OpportunityContactsAPI.GetSalesOpportunitiesByParentIdContactsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | contactId | 
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesOpportunitiesByParentIdContactsByIdRequest struct via the builder pattern


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

[**OpportunityContact**](OpportunityContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalesOpportunitiesByParentIdContactsCount

> Count GetSalesOpportunitiesByParentIdContactsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of OpportunityContact

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
	parentId := int32(56) // int32 | opportunityId
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
	resp, r, err := apiClient.OpportunityContactsAPI.GetSalesOpportunitiesByParentIdContactsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityContactsAPI.GetSalesOpportunitiesByParentIdContactsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesOpportunitiesByParentIdContactsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `OpportunityContactsAPI.GetSalesOpportunitiesByParentIdContactsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesOpportunitiesByParentIdContactsCountRequest struct via the builder pattern


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


## PatchSalesOpportunitiesByParentIdContactsById

> OpportunityContact PatchSalesOpportunitiesByParentIdContactsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch OpportunityContact

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
	parentId := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpportunityContactsAPI.PatchSalesOpportunitiesByParentIdContactsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityContactsAPI.PatchSalesOpportunitiesByParentIdContactsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSalesOpportunitiesByParentIdContactsById`: OpportunityContact
	fmt.Fprintf(os.Stdout, "Response from `OpportunityContactsAPI.PatchSalesOpportunitiesByParentIdContactsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | contactId | 
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSalesOpportunitiesByParentIdContactsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**OpportunityContact**](OpportunityContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSalesOpportunitiesByParentIdContacts

> OpportunityContact PostSalesOpportunitiesByParentIdContacts(ctx, parentId).ClientId(clientId).OpportunityContact(opportunityContact).Execute()

Post OpportunityContact

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
	parentId := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 
	opportunityContact := *openapiclient.NewOpportunityContact(*openapiclient.NewContactReference()) // OpportunityContact | opportunityContact

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpportunityContactsAPI.PostSalesOpportunitiesByParentIdContacts(context.Background(), parentId).ClientId(clientId).OpportunityContact(opportunityContact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityContactsAPI.PostSalesOpportunitiesByParentIdContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSalesOpportunitiesByParentIdContacts`: OpportunityContact
	fmt.Fprintf(os.Stdout, "Response from `OpportunityContactsAPI.PostSalesOpportunitiesByParentIdContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSalesOpportunitiesByParentIdContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **opportunityContact** | [**OpportunityContact**](OpportunityContact.md) | opportunityContact | 

### Return type

[**OpportunityContact**](OpportunityContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSalesOpportunitiesByParentIdContactsById

> OpportunityContact PutSalesOpportunitiesByParentIdContactsById(ctx, id, parentId).ClientId(clientId).OpportunityContact(opportunityContact).Execute()

Put OpportunityContact

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
	parentId := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 
	opportunityContact := *openapiclient.NewOpportunityContact(*openapiclient.NewContactReference()) // OpportunityContact | opportunityContact

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpportunityContactsAPI.PutSalesOpportunitiesByParentIdContactsById(context.Background(), id, parentId).ClientId(clientId).OpportunityContact(opportunityContact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityContactsAPI.PutSalesOpportunitiesByParentIdContactsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSalesOpportunitiesByParentIdContactsById`: OpportunityContact
	fmt.Fprintf(os.Stdout, "Response from `OpportunityContactsAPI.PutSalesOpportunitiesByParentIdContactsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | contactId | 
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSalesOpportunitiesByParentIdContactsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **opportunityContact** | [**OpportunityContact**](OpportunityContact.md) | opportunityContact | 

### Return type

[**OpportunityContact**](OpportunityContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

