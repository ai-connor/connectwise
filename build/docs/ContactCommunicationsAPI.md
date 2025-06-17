# \ContactCommunicationsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyContactsByParentIdCommunicationsById**](ContactCommunicationsAPI.md#DeleteCompanyContactsByParentIdCommunicationsById) | **Delete** /company/contacts/{parentId}/communications/{id} | Delete ContactCommunication
[**GetCompanyContactsByParentIdCommunications**](ContactCommunicationsAPI.md#GetCompanyContactsByParentIdCommunications) | **Get** /company/contacts/{parentId}/communications | Get List of ContactCommunication
[**GetCompanyContactsByParentIdCommunicationsById**](ContactCommunicationsAPI.md#GetCompanyContactsByParentIdCommunicationsById) | **Get** /company/contacts/{parentId}/communications/{id} | Get ContactCommunication
[**GetCompanyContactsByParentIdCommunicationsCount**](ContactCommunicationsAPI.md#GetCompanyContactsByParentIdCommunicationsCount) | **Get** /company/contacts/{parentId}/communications/count | Get Count of ContactCommunication
[**PatchCompanyContactsByParentIdCommunicationsById**](ContactCommunicationsAPI.md#PatchCompanyContactsByParentIdCommunicationsById) | **Patch** /company/contacts/{parentId}/communications/{id} | Patch ContactCommunication
[**PostCompanyContactsByParentIdCommunications**](ContactCommunicationsAPI.md#PostCompanyContactsByParentIdCommunications) | **Post** /company/contacts/{parentId}/communications | Post ContactCommunication
[**PutCompanyContactsByParentIdCommunicationsById**](ContactCommunicationsAPI.md#PutCompanyContactsByParentIdCommunicationsById) | **Put** /company/contacts/{parentId}/communications/{id} | Put ContactCommunication



## DeleteCompanyContactsByParentIdCommunicationsById

> DeleteCompanyContactsByParentIdCommunicationsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ContactCommunication

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
	id := int32(56) // int32 | communicationId
	parentId := int32(56) // int32 | contactId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContactCommunicationsAPI.DeleteCompanyContactsByParentIdCommunicationsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactCommunicationsAPI.DeleteCompanyContactsByParentIdCommunicationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | communicationId | 
**parentId** | **int32** | contactId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyContactsByParentIdCommunicationsByIdRequest struct via the builder pattern


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


## GetCompanyContactsByParentIdCommunications

> []ContactCommunication GetCompanyContactsByParentIdCommunications(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ContactCommunication

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
	resp, r, err := apiClient.ContactCommunicationsAPI.GetCompanyContactsByParentIdCommunications(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactCommunicationsAPI.GetCompanyContactsByParentIdCommunications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyContactsByParentIdCommunications`: []ContactCommunication
	fmt.Fprintf(os.Stdout, "Response from `ContactCommunicationsAPI.GetCompanyContactsByParentIdCommunications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | contactId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyContactsByParentIdCommunicationsRequest struct via the builder pattern


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

[**[]ContactCommunication**](ContactCommunication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyContactsByParentIdCommunicationsById

> ContactCommunication GetCompanyContactsByParentIdCommunicationsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ContactCommunication

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
	id := int32(56) // int32 | communicationId
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
	resp, r, err := apiClient.ContactCommunicationsAPI.GetCompanyContactsByParentIdCommunicationsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactCommunicationsAPI.GetCompanyContactsByParentIdCommunicationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyContactsByParentIdCommunicationsById`: ContactCommunication
	fmt.Fprintf(os.Stdout, "Response from `ContactCommunicationsAPI.GetCompanyContactsByParentIdCommunicationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | communicationId | 
**parentId** | **int32** | contactId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyContactsByParentIdCommunicationsByIdRequest struct via the builder pattern


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

[**ContactCommunication**](ContactCommunication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyContactsByParentIdCommunicationsCount

> Count GetCompanyContactsByParentIdCommunicationsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ContactCommunication

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
	resp, r, err := apiClient.ContactCommunicationsAPI.GetCompanyContactsByParentIdCommunicationsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactCommunicationsAPI.GetCompanyContactsByParentIdCommunicationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyContactsByParentIdCommunicationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ContactCommunicationsAPI.GetCompanyContactsByParentIdCommunicationsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | contactId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyContactsByParentIdCommunicationsCountRequest struct via the builder pattern


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


## PatchCompanyContactsByParentIdCommunicationsById

> ContactCommunication PatchCompanyContactsByParentIdCommunicationsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ContactCommunication

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
	id := int32(56) // int32 | communicationId
	parentId := int32(56) // int32 | contactId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactCommunicationsAPI.PatchCompanyContactsByParentIdCommunicationsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactCommunicationsAPI.PatchCompanyContactsByParentIdCommunicationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyContactsByParentIdCommunicationsById`: ContactCommunication
	fmt.Fprintf(os.Stdout, "Response from `ContactCommunicationsAPI.PatchCompanyContactsByParentIdCommunicationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | communicationId | 
**parentId** | **int32** | contactId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyContactsByParentIdCommunicationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ContactCommunication**](ContactCommunication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyContactsByParentIdCommunications

> ContactCommunication PostCompanyContactsByParentIdCommunications(ctx, parentId).ClientId(clientId).ContactCommunication(contactCommunication).Execute()

Post ContactCommunication

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
	contactCommunication := *openapiclient.NewContactCommunication(*openapiclient.NewCommunicationTypeReference(), "Value_example") // ContactCommunication | contactCommunication

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactCommunicationsAPI.PostCompanyContactsByParentIdCommunications(context.Background(), parentId).ClientId(clientId).ContactCommunication(contactCommunication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactCommunicationsAPI.PostCompanyContactsByParentIdCommunications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyContactsByParentIdCommunications`: ContactCommunication
	fmt.Fprintf(os.Stdout, "Response from `ContactCommunicationsAPI.PostCompanyContactsByParentIdCommunications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | contactId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyContactsByParentIdCommunicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **contactCommunication** | [**ContactCommunication**](ContactCommunication.md) | contactCommunication | 

### Return type

[**ContactCommunication**](ContactCommunication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyContactsByParentIdCommunicationsById

> ContactCommunication PutCompanyContactsByParentIdCommunicationsById(ctx, id, parentId).ClientId(clientId).ContactCommunication(contactCommunication).Execute()

Put ContactCommunication

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
	id := int32(56) // int32 | communicationId
	parentId := int32(56) // int32 | contactId
	clientId := "clientId_example" // string | 
	contactCommunication := *openapiclient.NewContactCommunication(*openapiclient.NewCommunicationTypeReference(), "Value_example") // ContactCommunication | contactCommunication

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactCommunicationsAPI.PutCompanyContactsByParentIdCommunicationsById(context.Background(), id, parentId).ClientId(clientId).ContactCommunication(contactCommunication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactCommunicationsAPI.PutCompanyContactsByParentIdCommunicationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyContactsByParentIdCommunicationsById`: ContactCommunication
	fmt.Fprintf(os.Stdout, "Response from `ContactCommunicationsAPI.PutCompanyContactsByParentIdCommunicationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | communicationId | 
**parentId** | **int32** | contactId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyContactsByParentIdCommunicationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **contactCommunication** | [**ContactCommunication**](ContactCommunication.md) | contactCommunication | 

### Return type

[**ContactCommunication**](ContactCommunication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

