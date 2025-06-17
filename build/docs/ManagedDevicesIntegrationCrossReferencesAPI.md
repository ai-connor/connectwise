# \ManagedDevicesIntegrationCrossReferencesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById**](ManagedDevicesIntegrationCrossReferencesAPI.md#DeleteCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById) | **Delete** /company/managedDevicesIntegrations/{parentId}/crossReferences/{id} | Delete ManagedDevicesIntegrationCrossReference
[**GetCompanyManagedDevicesIntegrationsByParentIdCrossReferences**](ManagedDevicesIntegrationCrossReferencesAPI.md#GetCompanyManagedDevicesIntegrationsByParentIdCrossReferences) | **Get** /company/managedDevicesIntegrations/{parentId}/crossReferences | Get List of ManagedDevicesIntegrationCrossReference
[**GetCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById**](ManagedDevicesIntegrationCrossReferencesAPI.md#GetCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById) | **Get** /company/managedDevicesIntegrations/{parentId}/crossReferences/{id} | Get ManagedDevicesIntegrationCrossReference
[**GetCompanyManagedDevicesIntegrationsByParentIdCrossReferencesCount**](ManagedDevicesIntegrationCrossReferencesAPI.md#GetCompanyManagedDevicesIntegrationsByParentIdCrossReferencesCount) | **Get** /company/managedDevicesIntegrations/{parentId}/crossReferences/count | Get Count of ManagedDevicesIntegrationCrossReference
[**PatchCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById**](ManagedDevicesIntegrationCrossReferencesAPI.md#PatchCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById) | **Patch** /company/managedDevicesIntegrations/{parentId}/crossReferences/{id} | Patch ManagedDevicesIntegrationCrossReference
[**PostCompanyManagedDevicesIntegrationsByParentIdCrossReferences**](ManagedDevicesIntegrationCrossReferencesAPI.md#PostCompanyManagedDevicesIntegrationsByParentIdCrossReferences) | **Post** /company/managedDevicesIntegrations/{parentId}/crossReferences | Post ManagedDevicesIntegrationCrossReference
[**PutCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById**](ManagedDevicesIntegrationCrossReferencesAPI.md#PutCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById) | **Put** /company/managedDevicesIntegrations/{parentId}/crossReferences/{id} | Put ManagedDevicesIntegrationCrossReference



## DeleteCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById

> ManagedDevicesIntegrationCrossReference DeleteCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ManagedDevicesIntegrationCrossReference

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
	id := int32(56) // int32 | crossReferenceId
	parentId := int32(56) // int32 | managedDevicesIntegrationId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedDevicesIntegrationCrossReferencesAPI.DeleteCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationCrossReferencesAPI.DeleteCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById`: ManagedDevicesIntegrationCrossReference
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationCrossReferencesAPI.DeleteCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | crossReferenceId | 
**parentId** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyManagedDevicesIntegrationsByParentIdCrossReferencesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 

### Return type

[**ManagedDevicesIntegrationCrossReference**](ManagedDevicesIntegrationCrossReference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyManagedDevicesIntegrationsByParentIdCrossReferences

> []ManagedDevicesIntegrationCrossReference GetCompanyManagedDevicesIntegrationsByParentIdCrossReferences(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ManagedDevicesIntegrationCrossReference

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
	parentId := int32(56) // int32 | managedDevicesIntegrationId
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
	resp, r, err := apiClient.ManagedDevicesIntegrationCrossReferencesAPI.GetCompanyManagedDevicesIntegrationsByParentIdCrossReferences(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationCrossReferencesAPI.GetCompanyManagedDevicesIntegrationsByParentIdCrossReferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyManagedDevicesIntegrationsByParentIdCrossReferences`: []ManagedDevicesIntegrationCrossReference
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationCrossReferencesAPI.GetCompanyManagedDevicesIntegrationsByParentIdCrossReferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyManagedDevicesIntegrationsByParentIdCrossReferencesRequest struct via the builder pattern


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

[**[]ManagedDevicesIntegrationCrossReference**](ManagedDevicesIntegrationCrossReference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById

> ManagedDevicesIntegrationCrossReference GetCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ManagedDevicesIntegrationCrossReference

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
	id := int32(56) // int32 | crossReferenceId
	parentId := int32(56) // int32 | managedDevicesIntegrationId
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
	resp, r, err := apiClient.ManagedDevicesIntegrationCrossReferencesAPI.GetCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationCrossReferencesAPI.GetCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById`: ManagedDevicesIntegrationCrossReference
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationCrossReferencesAPI.GetCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | crossReferenceId | 
**parentId** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyManagedDevicesIntegrationsByParentIdCrossReferencesByIdRequest struct via the builder pattern


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

[**ManagedDevicesIntegrationCrossReference**](ManagedDevicesIntegrationCrossReference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyManagedDevicesIntegrationsByParentIdCrossReferencesCount

> Count GetCompanyManagedDevicesIntegrationsByParentIdCrossReferencesCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ManagedDevicesIntegrationCrossReference

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
	parentId := int32(56) // int32 | managedDevicesIntegrationId
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
	resp, r, err := apiClient.ManagedDevicesIntegrationCrossReferencesAPI.GetCompanyManagedDevicesIntegrationsByParentIdCrossReferencesCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationCrossReferencesAPI.GetCompanyManagedDevicesIntegrationsByParentIdCrossReferencesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyManagedDevicesIntegrationsByParentIdCrossReferencesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationCrossReferencesAPI.GetCompanyManagedDevicesIntegrationsByParentIdCrossReferencesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyManagedDevicesIntegrationsByParentIdCrossReferencesCountRequest struct via the builder pattern


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


## PatchCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById

> ManagedDevicesIntegrationCrossReference PatchCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ManagedDevicesIntegrationCrossReference

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
	id := int32(56) // int32 | crossReferenceId
	parentId := int32(56) // int32 | managedDevicesIntegrationId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedDevicesIntegrationCrossReferencesAPI.PatchCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationCrossReferencesAPI.PatchCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById`: ManagedDevicesIntegrationCrossReference
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationCrossReferencesAPI.PatchCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | crossReferenceId | 
**parentId** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyManagedDevicesIntegrationsByParentIdCrossReferencesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ManagedDevicesIntegrationCrossReference**](ManagedDevicesIntegrationCrossReference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyManagedDevicesIntegrationsByParentIdCrossReferences

> ManagedDevicesIntegrationCrossReference PostCompanyManagedDevicesIntegrationsByParentIdCrossReferences(ctx, parentId).ClientId(clientId).ManagedDevicesIntegrationCrossReference(managedDevicesIntegrationCrossReference).Execute()

Post ManagedDevicesIntegrationCrossReference

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
	parentId := int32(56) // int32 | managedDevicesIntegrationId
	clientId := "clientId_example" // string | 
	managedDevicesIntegrationCrossReference := *openapiclient.NewManagedDevicesIntegrationCrossReference() // ManagedDevicesIntegrationCrossReference | crossReference

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedDevicesIntegrationCrossReferencesAPI.PostCompanyManagedDevicesIntegrationsByParentIdCrossReferences(context.Background(), parentId).ClientId(clientId).ManagedDevicesIntegrationCrossReference(managedDevicesIntegrationCrossReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationCrossReferencesAPI.PostCompanyManagedDevicesIntegrationsByParentIdCrossReferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyManagedDevicesIntegrationsByParentIdCrossReferences`: ManagedDevicesIntegrationCrossReference
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationCrossReferencesAPI.PostCompanyManagedDevicesIntegrationsByParentIdCrossReferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyManagedDevicesIntegrationsByParentIdCrossReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **managedDevicesIntegrationCrossReference** | [**ManagedDevicesIntegrationCrossReference**](ManagedDevicesIntegrationCrossReference.md) | crossReference | 

### Return type

[**ManagedDevicesIntegrationCrossReference**](ManagedDevicesIntegrationCrossReference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById

> ManagedDevicesIntegrationCrossReference PutCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById(ctx, id, parentId).ClientId(clientId).ManagedDevicesIntegrationCrossReference(managedDevicesIntegrationCrossReference).Execute()

Put ManagedDevicesIntegrationCrossReference

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
	id := int32(56) // int32 | crossReferenceId
	parentId := int32(56) // int32 | managedDevicesIntegrationId
	clientId := "clientId_example" // string | 
	managedDevicesIntegrationCrossReference := *openapiclient.NewManagedDevicesIntegrationCrossReference() // ManagedDevicesIntegrationCrossReference | crossReference

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedDevicesIntegrationCrossReferencesAPI.PutCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById(context.Background(), id, parentId).ClientId(clientId).ManagedDevicesIntegrationCrossReference(managedDevicesIntegrationCrossReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationCrossReferencesAPI.PutCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById`: ManagedDevicesIntegrationCrossReference
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationCrossReferencesAPI.PutCompanyManagedDevicesIntegrationsByParentIdCrossReferencesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | crossReferenceId | 
**parentId** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyManagedDevicesIntegrationsByParentIdCrossReferencesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **managedDevicesIntegrationCrossReference** | [**ManagedDevicesIntegrationCrossReference**](ManagedDevicesIntegrationCrossReference.md) | crossReference | 

### Return type

[**ManagedDevicesIntegrationCrossReference**](ManagedDevicesIntegrationCrossReference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

