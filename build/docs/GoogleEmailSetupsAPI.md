# \GoogleEmailSetupsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemGoogleemailsetupById**](GoogleEmailSetupsAPI.md#DeleteSystemGoogleemailsetupById) | **Delete** /system/googleemailsetup/{id} | Delete GoogleEmailSetups
[**GetSystemGoogleemailsetup**](GoogleEmailSetupsAPI.md#GetSystemGoogleemailsetup) | **Get** /system/googleemailsetup/ | Get List of GoogleEmailSetups
[**GetSystemGoogleemailsetupById**](GoogleEmailSetupsAPI.md#GetSystemGoogleemailsetupById) | **Get** /system/googleemailsetup/{id} | Get GoogleEmailSetups
[**GetSystemGoogleemailsetupCount**](GoogleEmailSetupsAPI.md#GetSystemGoogleemailsetupCount) | **Get** /system/googleemailsetup/count | Get Count of GoogleEmailSetups
[**PatchSystemGoogleemailsetupById**](GoogleEmailSetupsAPI.md#PatchSystemGoogleemailsetupById) | **Patch** /system/googleemailsetup/{id} | Patch GoogleEmailSetups
[**PostSystemGoogleemailsetup**](GoogleEmailSetupsAPI.md#PostSystemGoogleemailsetup) | **Post** /system/googleemailsetup/ | Post GoogleEmailSetups
[**PostSystemGoogleemailsetupByIdTestConnection**](GoogleEmailSetupsAPI.md#PostSystemGoogleemailsetupByIdTestConnection) | **Post** /system/googleemailsetup/{id}/testConnection | Post SuccessResponse
[**PutSystemGoogleemailsetupById**](GoogleEmailSetupsAPI.md#PutSystemGoogleemailsetupById) | **Put** /system/googleemailsetup/{id} | Put GoogleEmailSetups



## DeleteSystemGoogleemailsetupById

> DeleteSystemGoogleemailsetupById(ctx, id).ClientId(clientId).Execute()

Delete GoogleEmailSetups

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
	id := int32(56) // int32 | GoogleEmailSetupId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GoogleEmailSetupsAPI.DeleteSystemGoogleemailsetupById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoogleEmailSetupsAPI.DeleteSystemGoogleemailsetupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | GoogleEmailSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemGoogleemailsetupByIdRequest struct via the builder pattern


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


## GetSystemGoogleemailsetup

> []GoogleEmailSetup GetSystemGoogleemailsetup(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of GoogleEmailSetups

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
	resp, r, err := apiClient.GoogleEmailSetupsAPI.GetSystemGoogleemailsetup(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoogleEmailSetupsAPI.GetSystemGoogleemailsetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemGoogleemailsetup`: []GoogleEmailSetup
	fmt.Fprintf(os.Stdout, "Response from `GoogleEmailSetupsAPI.GetSystemGoogleemailsetup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemGoogleemailsetupRequest struct via the builder pattern


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

[**[]GoogleEmailSetup**](GoogleEmailSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemGoogleemailsetupById

> GoogleEmailSetup GetSystemGoogleemailsetupById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get GoogleEmailSetups

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
	id := int32(56) // int32 | GoogleEmailSetupId
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
	resp, r, err := apiClient.GoogleEmailSetupsAPI.GetSystemGoogleemailsetupById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoogleEmailSetupsAPI.GetSystemGoogleemailsetupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemGoogleemailsetupById`: GoogleEmailSetup
	fmt.Fprintf(os.Stdout, "Response from `GoogleEmailSetupsAPI.GetSystemGoogleemailsetupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | GoogleEmailSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemGoogleemailsetupByIdRequest struct via the builder pattern


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

[**GoogleEmailSetup**](GoogleEmailSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemGoogleemailsetupCount

> Count GetSystemGoogleemailsetupCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of GoogleEmailSetups

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
	resp, r, err := apiClient.GoogleEmailSetupsAPI.GetSystemGoogleemailsetupCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoogleEmailSetupsAPI.GetSystemGoogleemailsetupCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemGoogleemailsetupCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `GoogleEmailSetupsAPI.GetSystemGoogleemailsetupCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemGoogleemailsetupCountRequest struct via the builder pattern


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


## PatchSystemGoogleemailsetupById

> GoogleEmailSetup PatchSystemGoogleemailsetupById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch GoogleEmailSetups

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
	id := int32(56) // int32 | GoogleEmailSetupId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoogleEmailSetupsAPI.PatchSystemGoogleemailsetupById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoogleEmailSetupsAPI.PatchSystemGoogleemailsetupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemGoogleemailsetupById`: GoogleEmailSetup
	fmt.Fprintf(os.Stdout, "Response from `GoogleEmailSetupsAPI.PatchSystemGoogleemailsetupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | GoogleEmailSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemGoogleemailsetupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**GoogleEmailSetup**](GoogleEmailSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemGoogleemailsetup

> GoogleEmailSetup PostSystemGoogleemailsetup(ctx).ClientId(clientId).GoogleEmailSetup(googleEmailSetup).Execute()

Post GoogleEmailSetups

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
	googleEmailSetup := *openapiclient.NewGoogleEmailSetup("Name_example", "Username_example", "InboxFolder_example", "ProcessedFolder_example", "FailedFolder_example") // GoogleEmailSetup | GoogleEmailSetup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoogleEmailSetupsAPI.PostSystemGoogleemailsetup(context.Background()).ClientId(clientId).GoogleEmailSetup(googleEmailSetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoogleEmailSetupsAPI.PostSystemGoogleemailsetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemGoogleemailsetup`: GoogleEmailSetup
	fmt.Fprintf(os.Stdout, "Response from `GoogleEmailSetupsAPI.PostSystemGoogleemailsetup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemGoogleemailsetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **googleEmailSetup** | [**GoogleEmailSetup**](GoogleEmailSetup.md) | GoogleEmailSetup | 

### Return type

[**GoogleEmailSetup**](GoogleEmailSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemGoogleemailsetupByIdTestConnection

> SuccessResponse PostSystemGoogleemailsetupByIdTestConnection(ctx, id).ClientId(clientId).Execute()

Post SuccessResponse

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
	id := int32(56) // int32 | emailSetupId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoogleEmailSetupsAPI.PostSystemGoogleemailsetupByIdTestConnection(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoogleEmailSetupsAPI.PostSystemGoogleemailsetupByIdTestConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemGoogleemailsetupByIdTestConnection`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `GoogleEmailSetupsAPI.PostSystemGoogleemailsetupByIdTestConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemGoogleemailsetupByIdTestConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemGoogleemailsetupById

> GoogleEmailSetup PutSystemGoogleemailsetupById(ctx, id).ClientId(clientId).GoogleEmailSetup(googleEmailSetup).Execute()

Put GoogleEmailSetups

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
	id := int32(56) // int32 | GoogleEmailSetupId
	clientId := "clientId_example" // string | 
	googleEmailSetup := *openapiclient.NewGoogleEmailSetup("Name_example", "Username_example", "InboxFolder_example", "ProcessedFolder_example", "FailedFolder_example") // GoogleEmailSetup | companyTypeAssociation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GoogleEmailSetupsAPI.PutSystemGoogleemailsetupById(context.Background(), id).ClientId(clientId).GoogleEmailSetup(googleEmailSetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GoogleEmailSetupsAPI.PutSystemGoogleemailsetupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemGoogleemailsetupById`: GoogleEmailSetup
	fmt.Fprintf(os.Stdout, "Response from `GoogleEmailSetupsAPI.PutSystemGoogleemailsetupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | GoogleEmailSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemGoogleemailsetupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **googleEmailSetup** | [**GoogleEmailSetup**](GoogleEmailSetup.md) | companyTypeAssociation | 

### Return type

[**GoogleEmailSetup**](GoogleEmailSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

