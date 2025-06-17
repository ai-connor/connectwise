# \MemberNotificationSettingsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemMembersByParentIdNotificationSettingsById**](MemberNotificationSettingsAPI.md#DeleteSystemMembersByParentIdNotificationSettingsById) | **Delete** /system/members/{parentId}/notificationSettings/{id} | Delete MemberNotificationSetting
[**GetSystemMembersByParentIdNotificationSettings**](MemberNotificationSettingsAPI.md#GetSystemMembersByParentIdNotificationSettings) | **Get** /system/members/{parentId}/notificationSettings | Get List of MemberNotificationSetting
[**GetSystemMembersByParentIdNotificationSettingsById**](MemberNotificationSettingsAPI.md#GetSystemMembersByParentIdNotificationSettingsById) | **Get** /system/members/{parentId}/notificationSettings/{id} | Get MemberNotificationSetting
[**GetSystemMembersByParentIdNotificationSettingsCount**](MemberNotificationSettingsAPI.md#GetSystemMembersByParentIdNotificationSettingsCount) | **Get** /system/members/{parentId}/notificationSettings/count | Get Count of MemberNotificationSetting
[**PatchSystemMembersByParentIdNotificationSettingsById**](MemberNotificationSettingsAPI.md#PatchSystemMembersByParentIdNotificationSettingsById) | **Patch** /system/members/{parentId}/notificationSettings/{id} | Patch MemberNotificationSetting
[**PostSystemMembersByParentIdNotificationSettings**](MemberNotificationSettingsAPI.md#PostSystemMembersByParentIdNotificationSettings) | **Post** /system/members/{parentId}/notificationSettings | Post MemberNotificationSetting
[**PutSystemMembersByParentIdNotificationSettingsById**](MemberNotificationSettingsAPI.md#PutSystemMembersByParentIdNotificationSettingsById) | **Put** /system/members/{parentId}/notificationSettings/{id} | Put MemberNotificationSetting



## DeleteSystemMembersByParentIdNotificationSettingsById

> DeleteSystemMembersByParentIdNotificationSettingsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete MemberNotificationSetting

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
	id := int32(56) // int32 | notificationSettingId
	parentId := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MemberNotificationSettingsAPI.DeleteSystemMembersByParentIdNotificationSettingsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberNotificationSettingsAPI.DeleteSystemMembersByParentIdNotificationSettingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | notificationSettingId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemMembersByParentIdNotificationSettingsByIdRequest struct via the builder pattern


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


## GetSystemMembersByParentIdNotificationSettings

> []MemberNotificationSetting GetSystemMembersByParentIdNotificationSettings(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of MemberNotificationSetting

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
	parentId := int32(56) // int32 | memberId
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
	resp, r, err := apiClient.MemberNotificationSettingsAPI.GetSystemMembersByParentIdNotificationSettings(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberNotificationSettingsAPI.GetSystemMembersByParentIdNotificationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMembersByParentIdNotificationSettings`: []MemberNotificationSetting
	fmt.Fprintf(os.Stdout, "Response from `MemberNotificationSettingsAPI.GetSystemMembersByParentIdNotificationSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMembersByParentIdNotificationSettingsRequest struct via the builder pattern


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

[**[]MemberNotificationSetting**](MemberNotificationSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemMembersByParentIdNotificationSettingsById

> MemberNotificationSetting GetSystemMembersByParentIdNotificationSettingsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get MemberNotificationSetting

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
	id := int32(56) // int32 | notificationSettingId
	parentId := int32(56) // int32 | memberId
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
	resp, r, err := apiClient.MemberNotificationSettingsAPI.GetSystemMembersByParentIdNotificationSettingsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberNotificationSettingsAPI.GetSystemMembersByParentIdNotificationSettingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMembersByParentIdNotificationSettingsById`: MemberNotificationSetting
	fmt.Fprintf(os.Stdout, "Response from `MemberNotificationSettingsAPI.GetSystemMembersByParentIdNotificationSettingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | notificationSettingId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMembersByParentIdNotificationSettingsByIdRequest struct via the builder pattern


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

[**MemberNotificationSetting**](MemberNotificationSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemMembersByParentIdNotificationSettingsCount

> Count GetSystemMembersByParentIdNotificationSettingsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of MemberNotificationSetting

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
	parentId := int32(56) // int32 | memberId
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
	resp, r, err := apiClient.MemberNotificationSettingsAPI.GetSystemMembersByParentIdNotificationSettingsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberNotificationSettingsAPI.GetSystemMembersByParentIdNotificationSettingsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMembersByParentIdNotificationSettingsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `MemberNotificationSettingsAPI.GetSystemMembersByParentIdNotificationSettingsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMembersByParentIdNotificationSettingsCountRequest struct via the builder pattern


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


## PatchSystemMembersByParentIdNotificationSettingsById

> MemberNotificationSetting PatchSystemMembersByParentIdNotificationSettingsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch MemberNotificationSetting

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
	id := int32(56) // int32 | notificationSettingId
	parentId := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberNotificationSettingsAPI.PatchSystemMembersByParentIdNotificationSettingsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberNotificationSettingsAPI.PatchSystemMembersByParentIdNotificationSettingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemMembersByParentIdNotificationSettingsById`: MemberNotificationSetting
	fmt.Fprintf(os.Stdout, "Response from `MemberNotificationSettingsAPI.PatchSystemMembersByParentIdNotificationSettingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | notificationSettingId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemMembersByParentIdNotificationSettingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**MemberNotificationSetting**](MemberNotificationSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemMembersByParentIdNotificationSettings

> MemberNotificationSetting PostSystemMembersByParentIdNotificationSettings(ctx, parentId).ClientId(clientId).MemberNotificationSetting(memberNotificationSetting).Execute()

Post MemberNotificationSetting

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
	parentId := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 
	memberNotificationSetting := *openapiclient.NewMemberNotificationSetting("NotificationType_example", "NotificationTrigger_example") // MemberNotificationSetting | memberNotificationSetting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberNotificationSettingsAPI.PostSystemMembersByParentIdNotificationSettings(context.Background(), parentId).ClientId(clientId).MemberNotificationSetting(memberNotificationSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberNotificationSettingsAPI.PostSystemMembersByParentIdNotificationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemMembersByParentIdNotificationSettings`: MemberNotificationSetting
	fmt.Fprintf(os.Stdout, "Response from `MemberNotificationSettingsAPI.PostSystemMembersByParentIdNotificationSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemMembersByParentIdNotificationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **memberNotificationSetting** | [**MemberNotificationSetting**](MemberNotificationSetting.md) | memberNotificationSetting | 

### Return type

[**MemberNotificationSetting**](MemberNotificationSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemMembersByParentIdNotificationSettingsById

> MemberNotificationSetting PutSystemMembersByParentIdNotificationSettingsById(ctx, id, parentId).ClientId(clientId).MemberNotificationSetting(memberNotificationSetting).Execute()

Put MemberNotificationSetting

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
	id := int32(56) // int32 | notificationSettingId
	parentId := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 
	memberNotificationSetting := *openapiclient.NewMemberNotificationSetting("NotificationType_example", "NotificationTrigger_example") // MemberNotificationSetting | memberNotificationSetting

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberNotificationSettingsAPI.PutSystemMembersByParentIdNotificationSettingsById(context.Background(), id, parentId).ClientId(clientId).MemberNotificationSetting(memberNotificationSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberNotificationSettingsAPI.PutSystemMembersByParentIdNotificationSettingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemMembersByParentIdNotificationSettingsById`: MemberNotificationSetting
	fmt.Fprintf(os.Stdout, "Response from `MemberNotificationSettingsAPI.PutSystemMembersByParentIdNotificationSettingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | notificationSettingId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemMembersByParentIdNotificationSettingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **memberNotificationSetting** | [**MemberNotificationSetting**](MemberNotificationSetting.md) | memberNotificationSetting | 

### Return type

[**MemberNotificationSetting**](MemberNotificationSetting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

