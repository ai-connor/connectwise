# \ManagementNetworksSecurityAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemManagementNetworkSecuritiesById**](ManagementNetworksSecurityAPI.md#DeleteSystemManagementNetworkSecuritiesById) | **Delete** /system/managementNetworkSecurities/{id} | Delete ManagementNetworkSecurity
[**GetSystemManagementNetworkSecurities**](ManagementNetworksSecurityAPI.md#GetSystemManagementNetworkSecurities) | **Get** /system/managementNetworkSecurities | Get List of ManagementNetworkSecurity
[**GetSystemManagementNetworkSecuritiesById**](ManagementNetworksSecurityAPI.md#GetSystemManagementNetworkSecuritiesById) | **Get** /system/managementNetworkSecurities/{id} | Get ManagementNetworkSecurity
[**GetSystemManagementNetworkSecuritiesByIdTestCredentials**](ManagementNetworksSecurityAPI.md#GetSystemManagementNetworkSecuritiesByIdTestCredentials) | **Get** /system/managementNetworkSecurities/{id}/testCredentials | Get SuccessResponse
[**GetSystemManagementNetworkSecuritiesCount**](ManagementNetworksSecurityAPI.md#GetSystemManagementNetworkSecuritiesCount) | **Get** /system/managementNetworkSecurities/count | Get Count of ManagementNetworkSecurity
[**PatchSystemManagementNetworkSecuritiesById**](ManagementNetworksSecurityAPI.md#PatchSystemManagementNetworkSecuritiesById) | **Patch** /system/managementNetworkSecurities/{id} | Patch ManagementNetworkSecurity
[**PostSystemManagementNetworkSecurities**](ManagementNetworksSecurityAPI.md#PostSystemManagementNetworkSecurities) | **Post** /system/managementNetworkSecurities | Post ManagementNetworkSecurity
[**PutSystemManagementNetworkSecuritiesById**](ManagementNetworksSecurityAPI.md#PutSystemManagementNetworkSecuritiesById) | **Put** /system/managementNetworkSecurities/{id} | Put ManagementNetworkSecurity



## DeleteSystemManagementNetworkSecuritiesById

> DeleteSystemManagementNetworkSecuritiesById(ctx, id).ClientId(clientId).Execute()

Delete ManagementNetworkSecurity

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
	id := int32(56) // int32 | managementNetworkSecurityId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementNetworksSecurityAPI.DeleteSystemManagementNetworkSecuritiesById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementNetworksSecurityAPI.DeleteSystemManagementNetworkSecuritiesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementNetworkSecurityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemManagementNetworkSecuritiesByIdRequest struct via the builder pattern


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


## GetSystemManagementNetworkSecurities

> []ManagementNetworkSecurity GetSystemManagementNetworkSecurities(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ManagementNetworkSecurity

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
	resp, r, err := apiClient.ManagementNetworksSecurityAPI.GetSystemManagementNetworkSecurities(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementNetworksSecurityAPI.GetSystemManagementNetworkSecurities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemManagementNetworkSecurities`: []ManagementNetworkSecurity
	fmt.Fprintf(os.Stdout, "Response from `ManagementNetworksSecurityAPI.GetSystemManagementNetworkSecurities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemManagementNetworkSecuritiesRequest struct via the builder pattern


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

[**[]ManagementNetworkSecurity**](ManagementNetworkSecurity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemManagementNetworkSecuritiesById

> ManagementNetworkSecurity GetSystemManagementNetworkSecuritiesById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ManagementNetworkSecurity

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
	id := int32(56) // int32 | managementNetworkSecurityId
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
	resp, r, err := apiClient.ManagementNetworksSecurityAPI.GetSystemManagementNetworkSecuritiesById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementNetworksSecurityAPI.GetSystemManagementNetworkSecuritiesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemManagementNetworkSecuritiesById`: ManagementNetworkSecurity
	fmt.Fprintf(os.Stdout, "Response from `ManagementNetworksSecurityAPI.GetSystemManagementNetworkSecuritiesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementNetworkSecurityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemManagementNetworkSecuritiesByIdRequest struct via the builder pattern


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

[**ManagementNetworkSecurity**](ManagementNetworkSecurity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemManagementNetworkSecuritiesByIdTestCredentials

> SuccessResponse GetSystemManagementNetworkSecuritiesByIdTestCredentials(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get SuccessResponse

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
	id := int32(56) // int32 | managementNetworkSecurityId
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
	resp, r, err := apiClient.ManagementNetworksSecurityAPI.GetSystemManagementNetworkSecuritiesByIdTestCredentials(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementNetworksSecurityAPI.GetSystemManagementNetworkSecuritiesByIdTestCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemManagementNetworkSecuritiesByIdTestCredentials`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagementNetworksSecurityAPI.GetSystemManagementNetworkSecuritiesByIdTestCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementNetworkSecurityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemManagementNetworkSecuritiesByIdTestCredentialsRequest struct via the builder pattern


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

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemManagementNetworkSecuritiesCount

> Count GetSystemManagementNetworkSecuritiesCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ManagementNetworkSecurity

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
	resp, r, err := apiClient.ManagementNetworksSecurityAPI.GetSystemManagementNetworkSecuritiesCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementNetworksSecurityAPI.GetSystemManagementNetworkSecuritiesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemManagementNetworkSecuritiesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ManagementNetworksSecurityAPI.GetSystemManagementNetworkSecuritiesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemManagementNetworkSecuritiesCountRequest struct via the builder pattern


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


## PatchSystemManagementNetworkSecuritiesById

> ManagementNetworkSecurity PatchSystemManagementNetworkSecuritiesById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ManagementNetworkSecurity

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
	id := int32(56) // int32 | managementNetworkSecurityId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementNetworksSecurityAPI.PatchSystemManagementNetworkSecuritiesById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementNetworksSecurityAPI.PatchSystemManagementNetworkSecuritiesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemManagementNetworkSecuritiesById`: ManagementNetworkSecurity
	fmt.Fprintf(os.Stdout, "Response from `ManagementNetworksSecurityAPI.PatchSystemManagementNetworkSecuritiesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementNetworkSecurityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemManagementNetworkSecuritiesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ManagementNetworkSecurity**](ManagementNetworkSecurity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemManagementNetworkSecurities

> ManagementNetworkSecurity PostSystemManagementNetworkSecurities(ctx).ClientId(clientId).ManagementNetworkSecurity(managementNetworkSecurity).Execute()

Post ManagementNetworkSecurity

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
	managementNetworkSecurity := *openapiclient.NewManagementNetworkSecurity("Name_example", "Site_example") // ManagementNetworkSecurity | managementNetworkSecurity

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementNetworksSecurityAPI.PostSystemManagementNetworkSecurities(context.Background()).ClientId(clientId).ManagementNetworkSecurity(managementNetworkSecurity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementNetworksSecurityAPI.PostSystemManagementNetworkSecurities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemManagementNetworkSecurities`: ManagementNetworkSecurity
	fmt.Fprintf(os.Stdout, "Response from `ManagementNetworksSecurityAPI.PostSystemManagementNetworkSecurities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemManagementNetworkSecuritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **managementNetworkSecurity** | [**ManagementNetworkSecurity**](ManagementNetworkSecurity.md) | managementNetworkSecurity | 

### Return type

[**ManagementNetworkSecurity**](ManagementNetworkSecurity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemManagementNetworkSecuritiesById

> ManagementNetworkSecurity PutSystemManagementNetworkSecuritiesById(ctx, id).ClientId(clientId).ManagementNetworkSecurity(managementNetworkSecurity).Execute()

Put ManagementNetworkSecurity

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
	id := int32(56) // int32 | managementNetworkSecurityId
	clientId := "clientId_example" // string | 
	managementNetworkSecurity := *openapiclient.NewManagementNetworkSecurity("Name_example", "Site_example") // ManagementNetworkSecurity | managementNetworkSecurity

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementNetworksSecurityAPI.PutSystemManagementNetworkSecuritiesById(context.Background(), id).ClientId(clientId).ManagementNetworkSecurity(managementNetworkSecurity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementNetworksSecurityAPI.PutSystemManagementNetworkSecuritiesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemManagementNetworkSecuritiesById`: ManagementNetworkSecurity
	fmt.Fprintf(os.Stdout, "Response from `ManagementNetworksSecurityAPI.PutSystemManagementNetworkSecuritiesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementNetworkSecurityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemManagementNetworkSecuritiesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **managementNetworkSecurity** | [**ManagementNetworkSecurity**](ManagementNetworkSecurity.md) | managementNetworkSecurity | 

### Return type

[**ManagementNetworkSecurity**](ManagementNetworkSecurity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

