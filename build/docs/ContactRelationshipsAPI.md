# \ContactRelationshipsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyContactsRelationshipsById**](ContactRelationshipsAPI.md#DeleteCompanyContactsRelationshipsById) | **Delete** /company/contacts/relationships/{id} | Delete ContactRelationship
[**GetCompanyContactsRelationships**](ContactRelationshipsAPI.md#GetCompanyContactsRelationships) | **Get** /company/contacts/relationships | Get List of ContactRelationship
[**GetCompanyContactsRelationshipsById**](ContactRelationshipsAPI.md#GetCompanyContactsRelationshipsById) | **Get** /company/contacts/relationships/{id} | Get ContactRelationship
[**GetCompanyContactsRelationshipsCount**](ContactRelationshipsAPI.md#GetCompanyContactsRelationshipsCount) | **Get** /company/contacts/relationships/count | Get Count of ContactRelationship
[**PatchCompanyContactsRelationshipsById**](ContactRelationshipsAPI.md#PatchCompanyContactsRelationshipsById) | **Patch** /company/contacts/relationships/{id} | Patch ContactRelationship
[**PostCompanyContactsRelationships**](ContactRelationshipsAPI.md#PostCompanyContactsRelationships) | **Post** /company/contacts/relationships | Post ContactRelationship
[**PutCompanyContactsRelationshipsById**](ContactRelationshipsAPI.md#PutCompanyContactsRelationshipsById) | **Put** /company/contacts/relationships/{id} | Put ContactRelationship



## DeleteCompanyContactsRelationshipsById

> DeleteCompanyContactsRelationshipsById(ctx, id).ClientId(clientId).Execute()

Delete ContactRelationship

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
	id := int32(56) // int32 | relationshipId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContactRelationshipsAPI.DeleteCompanyContactsRelationshipsById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactRelationshipsAPI.DeleteCompanyContactsRelationshipsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | relationshipId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyContactsRelationshipsByIdRequest struct via the builder pattern


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


## GetCompanyContactsRelationships

> []ContactRelationship GetCompanyContactsRelationships(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ContactRelationship

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
	resp, r, err := apiClient.ContactRelationshipsAPI.GetCompanyContactsRelationships(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactRelationshipsAPI.GetCompanyContactsRelationships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyContactsRelationships`: []ContactRelationship
	fmt.Fprintf(os.Stdout, "Response from `ContactRelationshipsAPI.GetCompanyContactsRelationships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyContactsRelationshipsRequest struct via the builder pattern


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

[**[]ContactRelationship**](ContactRelationship.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyContactsRelationshipsById

> ContactRelationship GetCompanyContactsRelationshipsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ContactRelationship

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
	id := int32(56) // int32 | relationshipId
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
	resp, r, err := apiClient.ContactRelationshipsAPI.GetCompanyContactsRelationshipsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactRelationshipsAPI.GetCompanyContactsRelationshipsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyContactsRelationshipsById`: ContactRelationship
	fmt.Fprintf(os.Stdout, "Response from `ContactRelationshipsAPI.GetCompanyContactsRelationshipsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | relationshipId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyContactsRelationshipsByIdRequest struct via the builder pattern


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

[**ContactRelationship**](ContactRelationship.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyContactsRelationshipsCount

> Count GetCompanyContactsRelationshipsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ContactRelationship

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
	resp, r, err := apiClient.ContactRelationshipsAPI.GetCompanyContactsRelationshipsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactRelationshipsAPI.GetCompanyContactsRelationshipsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyContactsRelationshipsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ContactRelationshipsAPI.GetCompanyContactsRelationshipsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyContactsRelationshipsCountRequest struct via the builder pattern


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


## PatchCompanyContactsRelationshipsById

> ContactRelationship PatchCompanyContactsRelationshipsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ContactRelationship

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
	id := int32(56) // int32 | relationshipId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactRelationshipsAPI.PatchCompanyContactsRelationshipsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactRelationshipsAPI.PatchCompanyContactsRelationshipsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyContactsRelationshipsById`: ContactRelationship
	fmt.Fprintf(os.Stdout, "Response from `ContactRelationshipsAPI.PatchCompanyContactsRelationshipsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | relationshipId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyContactsRelationshipsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ContactRelationship**](ContactRelationship.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyContactsRelationships

> ContactRelationship PostCompanyContactsRelationships(ctx).ClientId(clientId).ContactRelationship(contactRelationship).Execute()

Post ContactRelationship

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
	contactRelationship := *openapiclient.NewContactRelationship("Name_example") // ContactRelationship | contactRelationship

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactRelationshipsAPI.PostCompanyContactsRelationships(context.Background()).ClientId(clientId).ContactRelationship(contactRelationship).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactRelationshipsAPI.PostCompanyContactsRelationships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyContactsRelationships`: ContactRelationship
	fmt.Fprintf(os.Stdout, "Response from `ContactRelationshipsAPI.PostCompanyContactsRelationships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyContactsRelationshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **contactRelationship** | [**ContactRelationship**](ContactRelationship.md) | contactRelationship | 

### Return type

[**ContactRelationship**](ContactRelationship.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyContactsRelationshipsById

> ContactRelationship PutCompanyContactsRelationshipsById(ctx, id).ClientId(clientId).ContactRelationship(contactRelationship).Execute()

Put ContactRelationship

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
	id := int32(56) // int32 | relationshipId
	clientId := "clientId_example" // string | 
	contactRelationship := *openapiclient.NewContactRelationship("Name_example") // ContactRelationship | contactRelationship

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactRelationshipsAPI.PutCompanyContactsRelationshipsById(context.Background(), id).ClientId(clientId).ContactRelationship(contactRelationship).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactRelationshipsAPI.PutCompanyContactsRelationshipsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyContactsRelationshipsById`: ContactRelationship
	fmt.Fprintf(os.Stdout, "Response from `ContactRelationshipsAPI.PutCompanyContactsRelationshipsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | relationshipId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyContactsRelationshipsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **contactRelationship** | [**ContactRelationship**](ContactRelationship.md) | contactRelationship | 

### Return type

[**ContactRelationship**](ContactRelationship.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

