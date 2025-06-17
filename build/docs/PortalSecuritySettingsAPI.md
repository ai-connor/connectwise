# \PortalSecuritySettingsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompanyPortalSecuritySettings**](PortalSecuritySettingsAPI.md#GetCompanyPortalSecuritySettings) | **Get** /company/portalSecuritySettings | Get List of PortalSecuritySetting
[**GetCompanyPortalSecuritySettingsById**](PortalSecuritySettingsAPI.md#GetCompanyPortalSecuritySettingsById) | **Get** /company/portalSecuritySettings/{id} | Get PortalSecuritySetting
[**GetCompanyPortalSecuritySettingsCount**](PortalSecuritySettingsAPI.md#GetCompanyPortalSecuritySettingsCount) | **Get** /company/portalSecuritySettings/count | Get Count of PortalSecuritySetting
[**PatchCompanyPortalSecuritySettingsById**](PortalSecuritySettingsAPI.md#PatchCompanyPortalSecuritySettingsById) | **Patch** /company/portalSecuritySettings/{id} | Patch PortalSecuritySetting
[**PutCompanyPortalSecuritySettingsById**](PortalSecuritySettingsAPI.md#PutCompanyPortalSecuritySettingsById) | **Put** /company/portalSecuritySettings/{id} | Put PortalSecuritySetting



## GetCompanyPortalSecuritySettings

> []PortalSecuritySetting GetCompanyPortalSecuritySettings(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of PortalSecuritySetting

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
	resp, r, err := apiClient.PortalSecuritySettingsAPI.GetCompanyPortalSecuritySettings(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalSecuritySettingsAPI.GetCompanyPortalSecuritySettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyPortalSecuritySettings`: []PortalSecuritySetting
	fmt.Fprintf(os.Stdout, "Response from `PortalSecuritySettingsAPI.GetCompanyPortalSecuritySettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPortalSecuritySettingsRequest struct via the builder pattern


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

[**[]PortalSecuritySetting**](PortalSecuritySetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyPortalSecuritySettingsById

> PortalSecuritySetting GetCompanyPortalSecuritySettingsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get PortalSecuritySetting

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
	id := int32(56) // int32 | portalSecuritySettingId
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
	resp, r, err := apiClient.PortalSecuritySettingsAPI.GetCompanyPortalSecuritySettingsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalSecuritySettingsAPI.GetCompanyPortalSecuritySettingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyPortalSecuritySettingsById`: PortalSecuritySetting
	fmt.Fprintf(os.Stdout, "Response from `PortalSecuritySettingsAPI.GetCompanyPortalSecuritySettingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | portalSecuritySettingId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPortalSecuritySettingsByIdRequest struct via the builder pattern


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

[**PortalSecuritySetting**](PortalSecuritySetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyPortalSecuritySettingsCount

> Count GetCompanyPortalSecuritySettingsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of PortalSecuritySetting

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
	resp, r, err := apiClient.PortalSecuritySettingsAPI.GetCompanyPortalSecuritySettingsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalSecuritySettingsAPI.GetCompanyPortalSecuritySettingsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyPortalSecuritySettingsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `PortalSecuritySettingsAPI.GetCompanyPortalSecuritySettingsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyPortalSecuritySettingsCountRequest struct via the builder pattern


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


## PatchCompanyPortalSecuritySettingsById

> PortalSecuritySetting PatchCompanyPortalSecuritySettingsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch PortalSecuritySetting

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
	id := int32(56) // int32 | portalSecuritySettingId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalSecuritySettingsAPI.PatchCompanyPortalSecuritySettingsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalSecuritySettingsAPI.PatchCompanyPortalSecuritySettingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyPortalSecuritySettingsById`: PortalSecuritySetting
	fmt.Fprintf(os.Stdout, "Response from `PortalSecuritySettingsAPI.PatchCompanyPortalSecuritySettingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | portalSecuritySettingId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyPortalSecuritySettingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**PortalSecuritySetting**](PortalSecuritySetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyPortalSecuritySettingsById

> PortalSecuritySetting PutCompanyPortalSecuritySettingsById(ctx, id).ClientId(clientId).PortalSecuritySetting(portalSecuritySetting).Execute()

Put PortalSecuritySetting

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
	id := int32(56) // int32 | portalSecuritySettingId
	clientId := "clientId_example" // string | 
	portalSecuritySetting := *openapiclient.NewPortalSecuritySetting() // PortalSecuritySetting | portalSecurity

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalSecuritySettingsAPI.PutCompanyPortalSecuritySettingsById(context.Background(), id).ClientId(clientId).PortalSecuritySetting(portalSecuritySetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalSecuritySettingsAPI.PutCompanyPortalSecuritySettingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyPortalSecuritySettingsById`: PortalSecuritySetting
	fmt.Fprintf(os.Stdout, "Response from `PortalSecuritySettingsAPI.PutCompanyPortalSecuritySettingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | portalSecuritySettingId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyPortalSecuritySettingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **portalSecuritySetting** | [**PortalSecuritySetting**](PortalSecuritySetting.md) | portalSecurity | 

### Return type

[**PortalSecuritySetting**](PortalSecuritySetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

