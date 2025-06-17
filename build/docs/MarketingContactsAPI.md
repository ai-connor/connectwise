# \MarketingContactsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMarketingGroupsByParentIdContactsById**](MarketingContactsAPI.md#DeleteMarketingGroupsByParentIdContactsById) | **Delete** /marketing/groups/{parentId}/contacts/{id} | Delete MarketingContact
[**GetMarketingGroupsByParentIdContacts**](MarketingContactsAPI.md#GetMarketingGroupsByParentIdContacts) | **Get** /marketing/groups/{parentId}/contacts | Get List of MarketingContact
[**GetMarketingGroupsByParentIdContactsById**](MarketingContactsAPI.md#GetMarketingGroupsByParentIdContactsById) | **Get** /marketing/groups/{parentId}/contacts/{id} | Get MarketingContact
[**GetMarketingGroupsByParentIdContactsCount**](MarketingContactsAPI.md#GetMarketingGroupsByParentIdContactsCount) | **Get** /marketing/groups/{parentId}/contacts/count | Get Count of MarketingContact
[**PatchMarketingGroupsByParentIdContactsById**](MarketingContactsAPI.md#PatchMarketingGroupsByParentIdContactsById) | **Patch** /marketing/groups/{parentId}/contacts/{id} | Patch MarketingContact
[**PostMarketingGroupsByParentIdContacts**](MarketingContactsAPI.md#PostMarketingGroupsByParentIdContacts) | **Post** /marketing/groups/{parentId}/contacts | Post MarketingContact
[**PutMarketingGroupsByParentIdContactsById**](MarketingContactsAPI.md#PutMarketingGroupsByParentIdContactsById) | **Put** /marketing/groups/{parentId}/contacts/{id} | Put MarketingContact



## DeleteMarketingGroupsByParentIdContactsById

> DeleteMarketingGroupsByParentIdContactsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete MarketingContact

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
	parentId := int32(56) // int32 | groupId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MarketingContactsAPI.DeleteMarketingGroupsByParentIdContactsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingContactsAPI.DeleteMarketingGroupsByParentIdContactsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | contactId | 
**parentId** | **int32** | groupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketingGroupsByParentIdContactsByIdRequest struct via the builder pattern


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


## GetMarketingGroupsByParentIdContacts

> []MarketingContact GetMarketingGroupsByParentIdContacts(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of MarketingContact

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
	resp, r, err := apiClient.MarketingContactsAPI.GetMarketingGroupsByParentIdContacts(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingContactsAPI.GetMarketingGroupsByParentIdContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingGroupsByParentIdContacts`: []MarketingContact
	fmt.Fprintf(os.Stdout, "Response from `MarketingContactsAPI.GetMarketingGroupsByParentIdContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | groupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingGroupsByParentIdContactsRequest struct via the builder pattern


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

[**[]MarketingContact**](MarketingContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingGroupsByParentIdContactsById

> MarketingContact GetMarketingGroupsByParentIdContactsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get MarketingContact

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
	resp, r, err := apiClient.MarketingContactsAPI.GetMarketingGroupsByParentIdContactsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingContactsAPI.GetMarketingGroupsByParentIdContactsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingGroupsByParentIdContactsById`: MarketingContact
	fmt.Fprintf(os.Stdout, "Response from `MarketingContactsAPI.GetMarketingGroupsByParentIdContactsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | contactId | 
**parentId** | **int32** | groupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingGroupsByParentIdContactsByIdRequest struct via the builder pattern


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

[**MarketingContact**](MarketingContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingGroupsByParentIdContactsCount

> Count GetMarketingGroupsByParentIdContactsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of MarketingContact

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
	resp, r, err := apiClient.MarketingContactsAPI.GetMarketingGroupsByParentIdContactsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingContactsAPI.GetMarketingGroupsByParentIdContactsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingGroupsByParentIdContactsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `MarketingContactsAPI.GetMarketingGroupsByParentIdContactsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | groupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingGroupsByParentIdContactsCountRequest struct via the builder pattern


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


## PatchMarketingGroupsByParentIdContactsById

> MarketingContact PatchMarketingGroupsByParentIdContactsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch MarketingContact

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
	parentId := int32(56) // int32 | groupId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingContactsAPI.PatchMarketingGroupsByParentIdContactsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingContactsAPI.PatchMarketingGroupsByParentIdContactsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMarketingGroupsByParentIdContactsById`: MarketingContact
	fmt.Fprintf(os.Stdout, "Response from `MarketingContactsAPI.PatchMarketingGroupsByParentIdContactsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | contactId | 
**parentId** | **int32** | groupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMarketingGroupsByParentIdContactsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**MarketingContact**](MarketingContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingGroupsByParentIdContacts

> MarketingContact PostMarketingGroupsByParentIdContacts(ctx, parentId).ClientId(clientId).MarketingContact(marketingContact).Execute()

Post MarketingContact

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
	marketingContact := *openapiclient.NewMarketingContact() // MarketingContact | marketingContact

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingContactsAPI.PostMarketingGroupsByParentIdContacts(context.Background(), parentId).ClientId(clientId).MarketingContact(marketingContact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingContactsAPI.PostMarketingGroupsByParentIdContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingGroupsByParentIdContacts`: MarketingContact
	fmt.Fprintf(os.Stdout, "Response from `MarketingContactsAPI.PostMarketingGroupsByParentIdContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | groupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingGroupsByParentIdContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **marketingContact** | [**MarketingContact**](MarketingContact.md) | marketingContact | 

### Return type

[**MarketingContact**](MarketingContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMarketingGroupsByParentIdContactsById

> MarketingContact PutMarketingGroupsByParentIdContactsById(ctx, id, parentId).ClientId(clientId).MarketingContact(marketingContact).Execute()

Put MarketingContact

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
	parentId := int32(56) // int32 | groupId
	clientId := "clientId_example" // string | 
	marketingContact := *openapiclient.NewMarketingContact() // MarketingContact | marketingContact

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingContactsAPI.PutMarketingGroupsByParentIdContactsById(context.Background(), id, parentId).ClientId(clientId).MarketingContact(marketingContact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingContactsAPI.PutMarketingGroupsByParentIdContactsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutMarketingGroupsByParentIdContactsById`: MarketingContact
	fmt.Fprintf(os.Stdout, "Response from `MarketingContactsAPI.PutMarketingGroupsByParentIdContactsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | contactId | 
**parentId** | **int32** | groupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMarketingGroupsByParentIdContactsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **marketingContact** | [**MarketingContact**](MarketingContact.md) | marketingContact | 

### Return type

[**MarketingContact**](MarketingContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

