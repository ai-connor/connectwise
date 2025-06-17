# \OnPremiseSearchSettingsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemOnPremiseSearchSetting**](OnPremiseSearchSettingsAPI.md#GetSystemOnPremiseSearchSetting) | **Get** /system/onPremiseSearchSetting/ | Get List of OnPremiseSearchSettings
[**GetSystemOnPremiseSearchSettingById**](OnPremiseSearchSettingsAPI.md#GetSystemOnPremiseSearchSettingById) | **Get** /system/onPremiseSearchSetting/{id} | Get OnPremiseSearchSettings
[**GetSystemOnPremiseSearchSettingCount**](OnPremiseSearchSettingsAPI.md#GetSystemOnPremiseSearchSettingCount) | **Get** /system/onPremiseSearchSetting/count | Get Count of OnPremiseSearchSettings
[**PatchSystemOnPremiseSearchSettingById**](OnPremiseSearchSettingsAPI.md#PatchSystemOnPremiseSearchSettingById) | **Patch** /system/onPremiseSearchSetting/{id} | Patch OnPremiseSearchSettings
[**PutSystemOnPremiseSearchSettingById**](OnPremiseSearchSettingsAPI.md#PutSystemOnPremiseSearchSettingById) | **Put** /system/onPremiseSearchSetting/{id} | Put OnPremiseSearchSettings              This does not update Solr. This allows you to set Manage to the Solr password.



## GetSystemOnPremiseSearchSetting

> []OnPremiseSearchSetting GetSystemOnPremiseSearchSetting(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of OnPremiseSearchSettings

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
	resp, r, err := apiClient.OnPremiseSearchSettingsAPI.GetSystemOnPremiseSearchSetting(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremiseSearchSettingsAPI.GetSystemOnPremiseSearchSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemOnPremiseSearchSetting`: []OnPremiseSearchSetting
	fmt.Fprintf(os.Stdout, "Response from `OnPremiseSearchSettingsAPI.GetSystemOnPremiseSearchSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemOnPremiseSearchSettingRequest struct via the builder pattern


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

[**[]OnPremiseSearchSetting**](OnPremiseSearchSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemOnPremiseSearchSettingById

> OnPremiseSearchSetting GetSystemOnPremiseSearchSettingById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get OnPremiseSearchSettings

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
	id := int32(56) // int32 | OnPremiseSearchSettingId
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
	resp, r, err := apiClient.OnPremiseSearchSettingsAPI.GetSystemOnPremiseSearchSettingById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremiseSearchSettingsAPI.GetSystemOnPremiseSearchSettingById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemOnPremiseSearchSettingById`: OnPremiseSearchSetting
	fmt.Fprintf(os.Stdout, "Response from `OnPremiseSearchSettingsAPI.GetSystemOnPremiseSearchSettingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | OnPremiseSearchSettingId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemOnPremiseSearchSettingByIdRequest struct via the builder pattern


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

[**OnPremiseSearchSetting**](OnPremiseSearchSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemOnPremiseSearchSettingCount

> Count GetSystemOnPremiseSearchSettingCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of OnPremiseSearchSettings

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
	resp, r, err := apiClient.OnPremiseSearchSettingsAPI.GetSystemOnPremiseSearchSettingCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremiseSearchSettingsAPI.GetSystemOnPremiseSearchSettingCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemOnPremiseSearchSettingCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `OnPremiseSearchSettingsAPI.GetSystemOnPremiseSearchSettingCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemOnPremiseSearchSettingCountRequest struct via the builder pattern


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


## PatchSystemOnPremiseSearchSettingById

> OnPremiseSearchSetting PatchSystemOnPremiseSearchSettingById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch OnPremiseSearchSettings

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
	id := int32(56) // int32 | OnPremiseSearchSettingId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnPremiseSearchSettingsAPI.PatchSystemOnPremiseSearchSettingById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremiseSearchSettingsAPI.PatchSystemOnPremiseSearchSettingById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemOnPremiseSearchSettingById`: OnPremiseSearchSetting
	fmt.Fprintf(os.Stdout, "Response from `OnPremiseSearchSettingsAPI.PatchSystemOnPremiseSearchSettingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | OnPremiseSearchSettingId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemOnPremiseSearchSettingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**OnPremiseSearchSetting**](OnPremiseSearchSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemOnPremiseSearchSettingById

> OnPremiseSearchSetting PutSystemOnPremiseSearchSettingById(ctx, id).ClientId(clientId).OnPremiseSearchSetting(onPremiseSearchSetting).Execute()

Put OnPremiseSearchSettings              This does not update Solr. This allows you to set Manage to the Solr password.

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
	id := int32(56) // int32 | OnPremiseSearchSettingId
	clientId := "clientId_example" // string | 
	onPremiseSearchSetting := *openapiclient.NewOnPremiseSearchSetting("Password_example") // OnPremiseSearchSetting | onPremiseSearchSetting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnPremiseSearchSettingsAPI.PutSystemOnPremiseSearchSettingById(context.Background(), id).ClientId(clientId).OnPremiseSearchSetting(onPremiseSearchSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnPremiseSearchSettingsAPI.PutSystemOnPremiseSearchSettingById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemOnPremiseSearchSettingById`: OnPremiseSearchSetting
	fmt.Fprintf(os.Stdout, "Response from `OnPremiseSearchSettingsAPI.PutSystemOnPremiseSearchSettingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | OnPremiseSearchSettingId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemOnPremiseSearchSettingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **onPremiseSearchSetting** | [**OnPremiseSearchSetting**](OnPremiseSearchSetting.md) | onPremiseSearchSetting | 

### Return type

[**OnPremiseSearchSetting**](OnPremiseSearchSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

