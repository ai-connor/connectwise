# \ManagedDeviceAccountsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemMembersByParentIdManagedDeviceAccountsBulk**](ManagedDeviceAccountsAPI.md#DeleteSystemMembersByParentIdManagedDeviceAccountsBulk) | **Delete** /system/members/{parentId}/managedDeviceAccounts/bulk | Delete BulkResult
[**GetSystemMembersByParentIdManagedDeviceAccounts**](ManagedDeviceAccountsAPI.md#GetSystemMembersByParentIdManagedDeviceAccounts) | **Get** /system/members/{parentId}/managedDeviceAccounts | Get List of ManagedDeviceAccount
[**PutSystemMembersByParentIdManagedDeviceAccountsBulk**](ManagedDeviceAccountsAPI.md#PutSystemMembersByParentIdManagedDeviceAccountsBulk) | **Put** /system/members/{parentId}/managedDeviceAccounts/bulk | Put BulkResult



## DeleteSystemMembersByParentIdManagedDeviceAccountsBulk

> BulkResult DeleteSystemMembersByParentIdManagedDeviceAccountsBulk(ctx, parentId).ClientId(clientId).IdCollection(idCollection).Execute()

Delete BulkResult

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
	idCollection := *openapiclient.NewIdCollection() // IdCollection | managedDeviceAccounts

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedDeviceAccountsAPI.DeleteSystemMembersByParentIdManagedDeviceAccountsBulk(context.Background(), parentId).ClientId(clientId).IdCollection(idCollection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDeviceAccountsAPI.DeleteSystemMembersByParentIdManagedDeviceAccountsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSystemMembersByParentIdManagedDeviceAccountsBulk`: BulkResult
	fmt.Fprintf(os.Stdout, "Response from `ManagedDeviceAccountsAPI.DeleteSystemMembersByParentIdManagedDeviceAccountsBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemMembersByParentIdManagedDeviceAccountsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **idCollection** | [**IdCollection**](IdCollection.md) | managedDeviceAccounts | 

### Return type

[**BulkResult**](BulkResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemMembersByParentIdManagedDeviceAccounts

> []ManagedDeviceAccount GetSystemMembersByParentIdManagedDeviceAccounts(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ManagedDeviceAccount

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
	resp, r, err := apiClient.ManagedDeviceAccountsAPI.GetSystemMembersByParentIdManagedDeviceAccounts(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDeviceAccountsAPI.GetSystemMembersByParentIdManagedDeviceAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMembersByParentIdManagedDeviceAccounts`: []ManagedDeviceAccount
	fmt.Fprintf(os.Stdout, "Response from `ManagedDeviceAccountsAPI.GetSystemMembersByParentIdManagedDeviceAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMembersByParentIdManagedDeviceAccountsRequest struct via the builder pattern


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

[**[]ManagedDeviceAccount**](ManagedDeviceAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemMembersByParentIdManagedDeviceAccountsBulk

> BulkResult PutSystemMembersByParentIdManagedDeviceAccountsBulk(ctx, parentId).ClientId(clientId).ManagedDeviceAccount(managedDeviceAccount).Execute()

Put BulkResult

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
	managedDeviceAccount := []openapiclient.ManagedDeviceAccount{*openapiclient.NewManagedDeviceAccount()} // []ManagedDeviceAccount | List of ManagedDeviceAccount

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedDeviceAccountsAPI.PutSystemMembersByParentIdManagedDeviceAccountsBulk(context.Background(), parentId).ClientId(clientId).ManagedDeviceAccount(managedDeviceAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedDeviceAccountsAPI.PutSystemMembersByParentIdManagedDeviceAccountsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemMembersByParentIdManagedDeviceAccountsBulk`: BulkResult
	fmt.Fprintf(os.Stdout, "Response from `ManagedDeviceAccountsAPI.PutSystemMembersByParentIdManagedDeviceAccountsBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemMembersByParentIdManagedDeviceAccountsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **managedDeviceAccount** | [**[]ManagedDeviceAccount**](ManagedDeviceAccount.md) | List of ManagedDeviceAccount | 

### Return type

[**BulkResult**](BulkResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

