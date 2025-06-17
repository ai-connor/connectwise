# \ManagementReportNotificationsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyCompaniesByParentIdManagementReportNotificationsById**](ManagementReportNotificationsAPI.md#DeleteCompanyCompaniesByParentIdManagementReportNotificationsById) | **Delete** /company/companies/{parentId}/managementReportNotifications/{id} | Delete ManagementReportNotification
[**DeleteCompanyManagementByParentIdManagementReportNotificationsById**](ManagementReportNotificationsAPI.md#DeleteCompanyManagementByParentIdManagementReportNotificationsById) | **Delete** /company/management/{parentId}/managementReportNotifications/{id} | Delete ManagementReportNotification
[**GetCompanyCompaniesByParentIdManagementReportNotifications**](ManagementReportNotificationsAPI.md#GetCompanyCompaniesByParentIdManagementReportNotifications) | **Get** /company/companies/{parentId}/managementReportNotifications | Get List of ManagementReportNotification
[**GetCompanyCompaniesByParentIdManagementReportNotificationsById**](ManagementReportNotificationsAPI.md#GetCompanyCompaniesByParentIdManagementReportNotificationsById) | **Get** /company/companies/{parentId}/managementReportNotifications/{id} | Get ManagementReportNotification
[**GetCompanyCompaniesByParentIdManagementReportNotificationsCount**](ManagementReportNotificationsAPI.md#GetCompanyCompaniesByParentIdManagementReportNotificationsCount) | **Get** /company/companies/{parentId}/managementReportNotifications/count | Get Count of ManagementReportNotification
[**GetCompanyManagementByParentIdManagementReportNotifications**](ManagementReportNotificationsAPI.md#GetCompanyManagementByParentIdManagementReportNotifications) | **Get** /company/management/{parentId}/managementReportNotifications | Get List of ManagementReportNotification
[**GetCompanyManagementByParentIdManagementReportNotificationsById**](ManagementReportNotificationsAPI.md#GetCompanyManagementByParentIdManagementReportNotificationsById) | **Get** /company/management/{parentId}/managementReportNotifications/{id} | Get ManagementReportNotification
[**GetCompanyManagementByParentIdManagementReportNotificationsCount**](ManagementReportNotificationsAPI.md#GetCompanyManagementByParentIdManagementReportNotificationsCount) | **Get** /company/management/{parentId}/managementReportNotifications/count | Get Count of ManagementReportNotification
[**PatchCompanyCompaniesByParentIdManagementReportNotificationsById**](ManagementReportNotificationsAPI.md#PatchCompanyCompaniesByParentIdManagementReportNotificationsById) | **Patch** /company/companies/{parentId}/managementReportNotifications/{id} | Patch ManagementReportNotification
[**PatchCompanyManagementByParentIdManagementReportNotificationsById**](ManagementReportNotificationsAPI.md#PatchCompanyManagementByParentIdManagementReportNotificationsById) | **Patch** /company/management/{parentId}/managementReportNotifications/{id} | Patch ManagementReportNotification
[**PostCompanyCompaniesByParentIdManagementReportNotifications**](ManagementReportNotificationsAPI.md#PostCompanyCompaniesByParentIdManagementReportNotifications) | **Post** /company/companies/{parentId}/managementReportNotifications | Post ManagementReportNotification
[**PostCompanyManagementByParentIdManagementReportNotifications**](ManagementReportNotificationsAPI.md#PostCompanyManagementByParentIdManagementReportNotifications) | **Post** /company/management/{parentId}/managementReportNotifications | Post ManagementReportNotification
[**PutCompanyCompaniesByParentIdManagementReportNotificationsById**](ManagementReportNotificationsAPI.md#PutCompanyCompaniesByParentIdManagementReportNotificationsById) | **Put** /company/companies/{parentId}/managementReportNotifications/{id} | Put ManagementReportNotification
[**PutCompanyManagementByParentIdManagementReportNotificationsById**](ManagementReportNotificationsAPI.md#PutCompanyManagementByParentIdManagementReportNotificationsById) | **Put** /company/management/{parentId}/managementReportNotifications/{id} | Put ManagementReportNotification



## DeleteCompanyCompaniesByParentIdManagementReportNotificationsById

> DeleteCompanyCompaniesByParentIdManagementReportNotificationsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ManagementReportNotification

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
	id := int32(56) // int32 | managementReportNotificationId
	parentId := int32(56) // int32 | companyId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementReportNotificationsAPI.DeleteCompanyCompaniesByParentIdManagementReportNotificationsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementReportNotificationsAPI.DeleteCompanyCompaniesByParentIdManagementReportNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementReportNotificationId | 
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyCompaniesByParentIdManagementReportNotificationsByIdRequest struct via the builder pattern


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


## DeleteCompanyManagementByParentIdManagementReportNotificationsById

> DeleteCompanyManagementByParentIdManagementReportNotificationsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete ManagementReportNotification

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
	id := int32(56) // int32 | managementReportNotificationId
	parentId := int32(56) // int32 | managementId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementReportNotificationsAPI.DeleteCompanyManagementByParentIdManagementReportNotificationsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementReportNotificationsAPI.DeleteCompanyManagementByParentIdManagementReportNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementReportNotificationId | 
**parentId** | **int32** | managementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyManagementByParentIdManagementReportNotificationsByIdRequest struct via the builder pattern


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


## GetCompanyCompaniesByParentIdManagementReportNotifications

> []ManagementReportNotification GetCompanyCompaniesByParentIdManagementReportNotifications(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ManagementReportNotification

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
	parentId := int32(56) // int32 | companyId
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
	resp, r, err := apiClient.ManagementReportNotificationsAPI.GetCompanyCompaniesByParentIdManagementReportNotifications(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementReportNotificationsAPI.GetCompanyCompaniesByParentIdManagementReportNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCompaniesByParentIdManagementReportNotifications`: []ManagementReportNotification
	fmt.Fprintf(os.Stdout, "Response from `ManagementReportNotificationsAPI.GetCompanyCompaniesByParentIdManagementReportNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCompaniesByParentIdManagementReportNotificationsRequest struct via the builder pattern


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

[**[]ManagementReportNotification**](ManagementReportNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyCompaniesByParentIdManagementReportNotificationsById

> ManagementReportNotification GetCompanyCompaniesByParentIdManagementReportNotificationsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ManagementReportNotification

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
	id := int32(56) // int32 | managementReportNotificationId
	parentId := int32(56) // int32 | companyId
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
	resp, r, err := apiClient.ManagementReportNotificationsAPI.GetCompanyCompaniesByParentIdManagementReportNotificationsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementReportNotificationsAPI.GetCompanyCompaniesByParentIdManagementReportNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCompaniesByParentIdManagementReportNotificationsById`: ManagementReportNotification
	fmt.Fprintf(os.Stdout, "Response from `ManagementReportNotificationsAPI.GetCompanyCompaniesByParentIdManagementReportNotificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementReportNotificationId | 
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCompaniesByParentIdManagementReportNotificationsByIdRequest struct via the builder pattern


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

[**ManagementReportNotification**](ManagementReportNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyCompaniesByParentIdManagementReportNotificationsCount

> Count GetCompanyCompaniesByParentIdManagementReportNotificationsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ManagementReportNotification

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
	parentId := int32(56) // int32 | companyId
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
	resp, r, err := apiClient.ManagementReportNotificationsAPI.GetCompanyCompaniesByParentIdManagementReportNotificationsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementReportNotificationsAPI.GetCompanyCompaniesByParentIdManagementReportNotificationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCompaniesByParentIdManagementReportNotificationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ManagementReportNotificationsAPI.GetCompanyCompaniesByParentIdManagementReportNotificationsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCompaniesByParentIdManagementReportNotificationsCountRequest struct via the builder pattern


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


## GetCompanyManagementByParentIdManagementReportNotifications

> []ManagementReportNotification GetCompanyManagementByParentIdManagementReportNotifications(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ManagementReportNotification

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
	parentId := int32(56) // int32 | managementId
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
	resp, r, err := apiClient.ManagementReportNotificationsAPI.GetCompanyManagementByParentIdManagementReportNotifications(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementReportNotificationsAPI.GetCompanyManagementByParentIdManagementReportNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyManagementByParentIdManagementReportNotifications`: []ManagementReportNotification
	fmt.Fprintf(os.Stdout, "Response from `ManagementReportNotificationsAPI.GetCompanyManagementByParentIdManagementReportNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | managementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyManagementByParentIdManagementReportNotificationsRequest struct via the builder pattern


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

[**[]ManagementReportNotification**](ManagementReportNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyManagementByParentIdManagementReportNotificationsById

> ManagementReportNotification GetCompanyManagementByParentIdManagementReportNotificationsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ManagementReportNotification

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
	id := int32(56) // int32 | managementReportNotificationId
	parentId := int32(56) // int32 | managementId
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
	resp, r, err := apiClient.ManagementReportNotificationsAPI.GetCompanyManagementByParentIdManagementReportNotificationsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementReportNotificationsAPI.GetCompanyManagementByParentIdManagementReportNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyManagementByParentIdManagementReportNotificationsById`: ManagementReportNotification
	fmt.Fprintf(os.Stdout, "Response from `ManagementReportNotificationsAPI.GetCompanyManagementByParentIdManagementReportNotificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementReportNotificationId | 
**parentId** | **int32** | managementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyManagementByParentIdManagementReportNotificationsByIdRequest struct via the builder pattern


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

[**ManagementReportNotification**](ManagementReportNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyManagementByParentIdManagementReportNotificationsCount

> Count GetCompanyManagementByParentIdManagementReportNotificationsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ManagementReportNotification

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
	parentId := int32(56) // int32 | managementId
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
	resp, r, err := apiClient.ManagementReportNotificationsAPI.GetCompanyManagementByParentIdManagementReportNotificationsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementReportNotificationsAPI.GetCompanyManagementByParentIdManagementReportNotificationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyManagementByParentIdManagementReportNotificationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `ManagementReportNotificationsAPI.GetCompanyManagementByParentIdManagementReportNotificationsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | managementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyManagementByParentIdManagementReportNotificationsCountRequest struct via the builder pattern


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


## PatchCompanyCompaniesByParentIdManagementReportNotificationsById

> ManagementReportNotification PatchCompanyCompaniesByParentIdManagementReportNotificationsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ManagementReportNotification

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
	id := int32(56) // int32 | managementReportNotificationId
	parentId := int32(56) // int32 | companyId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementReportNotificationsAPI.PatchCompanyCompaniesByParentIdManagementReportNotificationsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementReportNotificationsAPI.PatchCompanyCompaniesByParentIdManagementReportNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyCompaniesByParentIdManagementReportNotificationsById`: ManagementReportNotification
	fmt.Fprintf(os.Stdout, "Response from `ManagementReportNotificationsAPI.PatchCompanyCompaniesByParentIdManagementReportNotificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementReportNotificationId | 
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyCompaniesByParentIdManagementReportNotificationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ManagementReportNotification**](ManagementReportNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCompanyManagementByParentIdManagementReportNotificationsById

> ManagementReportNotification PatchCompanyManagementByParentIdManagementReportNotificationsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ManagementReportNotification

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
	id := int32(56) // int32 | managementReportNotificationId
	parentId := int32(56) // int32 | managementId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementReportNotificationsAPI.PatchCompanyManagementByParentIdManagementReportNotificationsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementReportNotificationsAPI.PatchCompanyManagementByParentIdManagementReportNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyManagementByParentIdManagementReportNotificationsById`: ManagementReportNotification
	fmt.Fprintf(os.Stdout, "Response from `ManagementReportNotificationsAPI.PatchCompanyManagementByParentIdManagementReportNotificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementReportNotificationId | 
**parentId** | **int32** | managementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyManagementByParentIdManagementReportNotificationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ManagementReportNotification**](ManagementReportNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyCompaniesByParentIdManagementReportNotifications

> ManagementReportNotification PostCompanyCompaniesByParentIdManagementReportNotifications(ctx, parentId).ClientId(clientId).ManagementReportNotification(managementReportNotification).Execute()

Post ManagementReportNotification

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
	parentId := int32(56) // int32 | companyId
	clientId := "clientId_example" // string | 
	managementReportNotification := *openapiclient.NewManagementReportNotification(*openapiclient.NewNotificationRecipientReference()) // ManagementReportNotification | managementReportNotification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementReportNotificationsAPI.PostCompanyCompaniesByParentIdManagementReportNotifications(context.Background(), parentId).ClientId(clientId).ManagementReportNotification(managementReportNotification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementReportNotificationsAPI.PostCompanyCompaniesByParentIdManagementReportNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyCompaniesByParentIdManagementReportNotifications`: ManagementReportNotification
	fmt.Fprintf(os.Stdout, "Response from `ManagementReportNotificationsAPI.PostCompanyCompaniesByParentIdManagementReportNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyCompaniesByParentIdManagementReportNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **managementReportNotification** | [**ManagementReportNotification**](ManagementReportNotification.md) | managementReportNotification | 

### Return type

[**ManagementReportNotification**](ManagementReportNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyManagementByParentIdManagementReportNotifications

> ManagementReportNotification PostCompanyManagementByParentIdManagementReportNotifications(ctx, parentId).ClientId(clientId).ManagementReportNotification(managementReportNotification).Execute()

Post ManagementReportNotification

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
	parentId := int32(56) // int32 | managementId
	clientId := "clientId_example" // string | 
	managementReportNotification := *openapiclient.NewManagementReportNotification(*openapiclient.NewNotificationRecipientReference()) // ManagementReportNotification | managementReportNotification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementReportNotificationsAPI.PostCompanyManagementByParentIdManagementReportNotifications(context.Background(), parentId).ClientId(clientId).ManagementReportNotification(managementReportNotification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementReportNotificationsAPI.PostCompanyManagementByParentIdManagementReportNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyManagementByParentIdManagementReportNotifications`: ManagementReportNotification
	fmt.Fprintf(os.Stdout, "Response from `ManagementReportNotificationsAPI.PostCompanyManagementByParentIdManagementReportNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | managementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyManagementByParentIdManagementReportNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **managementReportNotification** | [**ManagementReportNotification**](ManagementReportNotification.md) | managementReportNotification | 

### Return type

[**ManagementReportNotification**](ManagementReportNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyCompaniesByParentIdManagementReportNotificationsById

> ManagementReportNotification PutCompanyCompaniesByParentIdManagementReportNotificationsById(ctx, id, parentId).ClientId(clientId).ManagementReportNotification(managementReportNotification).Execute()

Put ManagementReportNotification

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
	id := int32(56) // int32 | managementReportNotificationId
	parentId := int32(56) // int32 | companyId
	clientId := "clientId_example" // string | 
	managementReportNotification := *openapiclient.NewManagementReportNotification(*openapiclient.NewNotificationRecipientReference()) // ManagementReportNotification | managementReportNotification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementReportNotificationsAPI.PutCompanyCompaniesByParentIdManagementReportNotificationsById(context.Background(), id, parentId).ClientId(clientId).ManagementReportNotification(managementReportNotification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementReportNotificationsAPI.PutCompanyCompaniesByParentIdManagementReportNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyCompaniesByParentIdManagementReportNotificationsById`: ManagementReportNotification
	fmt.Fprintf(os.Stdout, "Response from `ManagementReportNotificationsAPI.PutCompanyCompaniesByParentIdManagementReportNotificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementReportNotificationId | 
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyCompaniesByParentIdManagementReportNotificationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **managementReportNotification** | [**ManagementReportNotification**](ManagementReportNotification.md) | managementReportNotification | 

### Return type

[**ManagementReportNotification**](ManagementReportNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyManagementByParentIdManagementReportNotificationsById

> ManagementReportNotification PutCompanyManagementByParentIdManagementReportNotificationsById(ctx, id, parentId).ClientId(clientId).ManagementReportNotification(managementReportNotification).Execute()

Put ManagementReportNotification

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
	id := int32(56) // int32 | managementReportNotificationId
	parentId := int32(56) // int32 | managementId
	clientId := "clientId_example" // string | 
	managementReportNotification := *openapiclient.NewManagementReportNotification(*openapiclient.NewNotificationRecipientReference()) // ManagementReportNotification | managementReportNotification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementReportNotificationsAPI.PutCompanyManagementByParentIdManagementReportNotificationsById(context.Background(), id, parentId).ClientId(clientId).ManagementReportNotification(managementReportNotification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementReportNotificationsAPI.PutCompanyManagementByParentIdManagementReportNotificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyManagementByParentIdManagementReportNotificationsById`: ManagementReportNotification
	fmt.Fprintf(os.Stdout, "Response from `ManagementReportNotificationsAPI.PutCompanyManagementByParentIdManagementReportNotificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementReportNotificationId | 
**parentId** | **int32** | managementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyManagementByParentIdManagementReportNotificationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **managementReportNotification** | [**ManagementReportNotification**](ManagementReportNotification.md) | managementReportNotification | 

### Return type

[**ManagementReportNotification**](ManagementReportNotification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

