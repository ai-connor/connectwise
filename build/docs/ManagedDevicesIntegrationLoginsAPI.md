# \ManagedDevicesIntegrationLoginsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyManagedDevicesIntegrationsByParentIdLoginsById**](ManagedDevicesIntegrationLoginsAPI.md#DeleteCompanyManagedDevicesIntegrationsByParentIdLoginsById) | **Delete** /company/managedDevicesIntegrations/{parentId}/logins/{id} | Delete ManagedDevicesIntegrationLogin
[**GetCompanyManagedDevicesIntegrationsByParentIdLogins**](ManagedDevicesIntegrationLoginsAPI.md#GetCompanyManagedDevicesIntegrationsByParentIdLogins) | **Get** /company/managedDevicesIntegrations/{parentId}/logins | Get List of ManagedDevicesIntegrationLogin
[**GetCompanyManagedDevicesIntegrationsByParentIdLoginsById**](ManagedDevicesIntegrationLoginsAPI.md#GetCompanyManagedDevicesIntegrationsByParentIdLoginsById) | **Get** /company/managedDevicesIntegrations/{parentId}/logins/{id} | Get ManagedDevicesIntegrationLogin
[**GetCompanyManagedDevicesIntegrationsByParentIdLoginsCount**](ManagedDevicesIntegrationLoginsAPI.md#GetCompanyManagedDevicesIntegrationsByParentIdLoginsCount) | **Get** /company/managedDevicesIntegrations/{parentId}/logins/count | Get Count of ManagedDevicesIntegrationLogin
[**PatchCompanyManagedDevicesIntegrationsByParentIdLoginsById**](ManagedDevicesIntegrationLoginsAPI.md#PatchCompanyManagedDevicesIntegrationsByParentIdLoginsById) | **Patch** /company/managedDevicesIntegrations/{parentId}/logins/{id} | Patch ManagedDevicesIntegrationLogin
[**PostCompanyManagedDevicesIntegrationsByParentIdLogins**](ManagedDevicesIntegrationLoginsAPI.md#PostCompanyManagedDevicesIntegrationsByParentIdLogins) | **Post** /company/managedDevicesIntegrations/{parentId}/logins | Post ManagedDevicesIntegrationLogin
[**PutCompanyManagedDevicesIntegrationsByParentIdLoginsById**](ManagedDevicesIntegrationLoginsAPI.md#PutCompanyManagedDevicesIntegrationsByParentIdLoginsById) | **Put** /company/managedDevicesIntegrations/{parentId}/logins/{id} | Put ManagedDevicesIntegrationLogin



## DeleteCompanyManagedDevicesIntegrationsByParentIdLoginsById

> ManagedDevicesIntegrationLogin DeleteCompanyManagedDevicesIntegrationsByParentIdLoginsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ManagedDevicesIntegrationLogin

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
	id := int32(56) // int32 | loginId
	parentId := int32(56) // int32 | managedDevicesIntegrationId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedDevicesIntegrationLoginsAPI.DeleteCompanyManagedDevicesIntegrationsByParentIdLoginsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationLoginsAPI.DeleteCompanyManagedDevicesIntegrationsByParentIdLoginsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCompanyManagedDevicesIntegrationsByParentIdLoginsById`: ManagedDevicesIntegrationLogin
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationLoginsAPI.DeleteCompanyManagedDevicesIntegrationsByParentIdLoginsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | loginId | 
**parentId** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyManagedDevicesIntegrationsByParentIdLoginsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 

### Return type

[**ManagedDevicesIntegrationLogin**](ManagedDevicesIntegrationLogin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyManagedDevicesIntegrationsByParentIdLogins

> []ManagedDevicesIntegrationLogin GetCompanyManagedDevicesIntegrationsByParentIdLogins(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ManagedDevicesIntegrationLogin

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
	resp, r, err := apiClient.ManagedDevicesIntegrationLoginsAPI.GetCompanyManagedDevicesIntegrationsByParentIdLogins(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationLoginsAPI.GetCompanyManagedDevicesIntegrationsByParentIdLogins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyManagedDevicesIntegrationsByParentIdLogins`: []ManagedDevicesIntegrationLogin
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationLoginsAPI.GetCompanyManagedDevicesIntegrationsByParentIdLogins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyManagedDevicesIntegrationsByParentIdLoginsRequest struct via the builder pattern


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

[**[]ManagedDevicesIntegrationLogin**](ManagedDevicesIntegrationLogin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyManagedDevicesIntegrationsByParentIdLoginsById

> ManagedDevicesIntegrationLogin GetCompanyManagedDevicesIntegrationsByParentIdLoginsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ManagedDevicesIntegrationLogin

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
	id := int32(56) // int32 | loginId
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
	resp, r, err := apiClient.ManagedDevicesIntegrationLoginsAPI.GetCompanyManagedDevicesIntegrationsByParentIdLoginsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationLoginsAPI.GetCompanyManagedDevicesIntegrationsByParentIdLoginsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyManagedDevicesIntegrationsByParentIdLoginsById`: ManagedDevicesIntegrationLogin
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationLoginsAPI.GetCompanyManagedDevicesIntegrationsByParentIdLoginsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | loginId | 
**parentId** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyManagedDevicesIntegrationsByParentIdLoginsByIdRequest struct via the builder pattern


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

[**ManagedDevicesIntegrationLogin**](ManagedDevicesIntegrationLogin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyManagedDevicesIntegrationsByParentIdLoginsCount

> Count GetCompanyManagedDevicesIntegrationsByParentIdLoginsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ManagedDevicesIntegrationLogin

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
	resp, r, err := apiClient.ManagedDevicesIntegrationLoginsAPI.GetCompanyManagedDevicesIntegrationsByParentIdLoginsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationLoginsAPI.GetCompanyManagedDevicesIntegrationsByParentIdLoginsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyManagedDevicesIntegrationsByParentIdLoginsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationLoginsAPI.GetCompanyManagedDevicesIntegrationsByParentIdLoginsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyManagedDevicesIntegrationsByParentIdLoginsCountRequest struct via the builder pattern


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


## PatchCompanyManagedDevicesIntegrationsByParentIdLoginsById

> ManagedDevicesIntegrationLogin PatchCompanyManagedDevicesIntegrationsByParentIdLoginsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ManagedDevicesIntegrationLogin

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
	id := int32(56) // int32 | loginId
	parentId := int32(56) // int32 | managedDevicesIntegrationId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedDevicesIntegrationLoginsAPI.PatchCompanyManagedDevicesIntegrationsByParentIdLoginsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationLoginsAPI.PatchCompanyManagedDevicesIntegrationsByParentIdLoginsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyManagedDevicesIntegrationsByParentIdLoginsById`: ManagedDevicesIntegrationLogin
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationLoginsAPI.PatchCompanyManagedDevicesIntegrationsByParentIdLoginsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | loginId | 
**parentId** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyManagedDevicesIntegrationsByParentIdLoginsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ManagedDevicesIntegrationLogin**](ManagedDevicesIntegrationLogin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyManagedDevicesIntegrationsByParentIdLogins

> ManagedDevicesIntegrationLogin PostCompanyManagedDevicesIntegrationsByParentIdLogins(ctx, parentId).ClientId(clientId).ManagedDevicesIntegrationLogin(managedDevicesIntegrationLogin).Execute()

Post ManagedDevicesIntegrationLogin

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
	managedDevicesIntegrationLogin := *openapiclient.NewManagedDevicesIntegrationLogin("Username_example", *openapiclient.NewMemberReference()) // ManagedDevicesIntegrationLogin | login

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedDevicesIntegrationLoginsAPI.PostCompanyManagedDevicesIntegrationsByParentIdLogins(context.Background(), parentId).ClientId(clientId).ManagedDevicesIntegrationLogin(managedDevicesIntegrationLogin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationLoginsAPI.PostCompanyManagedDevicesIntegrationsByParentIdLogins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyManagedDevicesIntegrationsByParentIdLogins`: ManagedDevicesIntegrationLogin
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationLoginsAPI.PostCompanyManagedDevicesIntegrationsByParentIdLogins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyManagedDevicesIntegrationsByParentIdLoginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **managedDevicesIntegrationLogin** | [**ManagedDevicesIntegrationLogin**](ManagedDevicesIntegrationLogin.md) | login | 

### Return type

[**ManagedDevicesIntegrationLogin**](ManagedDevicesIntegrationLogin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyManagedDevicesIntegrationsByParentIdLoginsById

> ManagedDevicesIntegrationLogin PutCompanyManagedDevicesIntegrationsByParentIdLoginsById(ctx, id, parentId).ClientId(clientId).ManagedDevicesIntegrationLogin(managedDevicesIntegrationLogin).Execute()

Put ManagedDevicesIntegrationLogin

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
	id := int32(56) // int32 | loginId
	parentId := int32(56) // int32 | managedDevicesIntegrationId
	clientId := "clientId_example" // string | 
	managedDevicesIntegrationLogin := *openapiclient.NewManagedDevicesIntegrationLogin("Username_example", *openapiclient.NewMemberReference()) // ManagedDevicesIntegrationLogin | login

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedDevicesIntegrationLoginsAPI.PutCompanyManagedDevicesIntegrationsByParentIdLoginsById(context.Background(), id, parentId).ClientId(clientId).ManagedDevicesIntegrationLogin(managedDevicesIntegrationLogin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDevicesIntegrationLoginsAPI.PutCompanyManagedDevicesIntegrationsByParentIdLoginsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyManagedDevicesIntegrationsByParentIdLoginsById`: ManagedDevicesIntegrationLogin
	fmt.Fprintf(os.Stdout, "Response from `ManagedDevicesIntegrationLoginsAPI.PutCompanyManagedDevicesIntegrationsByParentIdLoginsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | loginId | 
**parentId** | **int32** | managedDevicesIntegrationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyManagedDevicesIntegrationsByParentIdLoginsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **managedDevicesIntegrationLogin** | [**ManagedDevicesIntegrationLogin**](ManagedDevicesIntegrationLogin.md) | login | 

### Return type

[**ManagedDevicesIntegrationLogin**](ManagedDevicesIntegrationLogin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

