# \M365ContactSyncPropertiesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyM365contactsyncProperty**](M365ContactSyncPropertiesAPI.md#DeleteCompanyM365contactsyncProperty) | **Delete** /company/m365contactsync/property/ | Delete M365ContactSyncProperty
[**GetCompanyM365contactsyncByIdProperty**](M365ContactSyncPropertiesAPI.md#GetCompanyM365contactsyncByIdProperty) | **Get** /company/m365contactsync/{id}/property | Get M365ContactSyncProperties
[**GetCompanyM365contactsyncPropertyCount**](M365ContactSyncPropertiesAPI.md#GetCompanyM365contactsyncPropertyCount) | **Get** /company/m365contactsync/property/count | Get Count of M365ContactSyncProperty
[**GetCompanyM365contactsyncPropertyExcluded**](M365ContactSyncPropertiesAPI.md#GetCompanyM365contactsyncPropertyExcluded) | **Get** /company/m365contactsync/property/excluded | Get List of M365ContactSyncPropertiesExcluded
[**GetCompanyM365contactsyncPropertyIncluded**](M365ContactSyncPropertiesAPI.md#GetCompanyM365contactsyncPropertyIncluded) | **Get** /company/m365contactsync/property/included | Get List of M365ContactSyncPropertiesIncluded
[**PostCompanyM365contactsyncProperty**](M365ContactSyncPropertiesAPI.md#PostCompanyM365contactsyncProperty) | **Post** /company/m365contactsync/property | Create M365ContactSyncProperty



## DeleteCompanyM365contactsyncProperty

> DeleteCompanyM365contactsyncProperty(ctx).ClientId(clientId).Execute()

Delete M365ContactSyncProperty

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.M365ContactSyncPropertiesAPI.DeleteCompanyM365contactsyncProperty(context.Background()).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `M365ContactSyncPropertiesAPI.DeleteCompanyM365contactsyncProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyM365contactsyncPropertyRequest struct via the builder pattern


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


## GetCompanyM365contactsyncByIdProperty

> M365ContactSyncProperty GetCompanyM365contactsyncByIdProperty(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get M365ContactSyncProperties

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
	id := int32(56) // int32 | M365ContactSyncPropertyId
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
	resp, r, err := apiClient.M365ContactSyncPropertiesAPI.GetCompanyM365contactsyncByIdProperty(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `M365ContactSyncPropertiesAPI.GetCompanyM365contactsyncByIdProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyM365contactsyncByIdProperty`: M365ContactSyncProperty
	fmt.Fprintf(os.Stdout, "Response from `M365ContactSyncPropertiesAPI.GetCompanyM365contactsyncByIdProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | M365ContactSyncPropertyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyM365contactsyncByIdPropertyRequest struct via the builder pattern


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

[**M365ContactSyncProperty**](M365ContactSyncProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyM365contactsyncPropertyCount

> Count GetCompanyM365contactsyncPropertyCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of M365ContactSyncProperty

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
	resp, r, err := apiClient.M365ContactSyncPropertiesAPI.GetCompanyM365contactsyncPropertyCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `M365ContactSyncPropertiesAPI.GetCompanyM365contactsyncPropertyCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyM365contactsyncPropertyCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `M365ContactSyncPropertiesAPI.GetCompanyM365contactsyncPropertyCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyM365contactsyncPropertyCountRequest struct via the builder pattern


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


## GetCompanyM365contactsyncPropertyExcluded

> []M365ContactSyncProperty GetCompanyM365contactsyncPropertyExcluded(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of M365ContactSyncPropertiesExcluded

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
	id := int32(56) // int32 | companyRecId
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
	resp, r, err := apiClient.M365ContactSyncPropertiesAPI.GetCompanyM365contactsyncPropertyExcluded(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `M365ContactSyncPropertiesAPI.GetCompanyM365contactsyncPropertyExcluded``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyM365contactsyncPropertyExcluded`: []M365ContactSyncProperty
	fmt.Fprintf(os.Stdout, "Response from `M365ContactSyncPropertiesAPI.GetCompanyM365contactsyncPropertyExcluded`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | companyRecId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyM365contactsyncPropertyExcludedRequest struct via the builder pattern


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

[**[]M365ContactSyncProperty**](M365ContactSyncProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyM365contactsyncPropertyIncluded

> []M365ContactSyncProperty GetCompanyM365contactsyncPropertyIncluded(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of M365ContactSyncPropertiesIncluded

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
	id := int32(56) // int32 | companyRecId
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
	resp, r, err := apiClient.M365ContactSyncPropertiesAPI.GetCompanyM365contactsyncPropertyIncluded(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `M365ContactSyncPropertiesAPI.GetCompanyM365contactsyncPropertyIncluded``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyM365contactsyncPropertyIncluded`: []M365ContactSyncProperty
	fmt.Fprintf(os.Stdout, "Response from `M365ContactSyncPropertiesAPI.GetCompanyM365contactsyncPropertyIncluded`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | companyRecId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyM365contactsyncPropertyIncludedRequest struct via the builder pattern


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

[**[]M365ContactSyncProperty**](M365ContactSyncProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyM365contactsyncProperty

> M365ContactSyncProperty PostCompanyM365contactsyncProperty(ctx).ClientId(clientId).M365ContactSyncProperty(m365ContactSyncProperty).Execute()

Create M365ContactSyncProperty

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
	m365ContactSyncProperty := *openapiclient.NewM365ContactSyncProperty() // M365ContactSyncProperty | country

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.M365ContactSyncPropertiesAPI.PostCompanyM365contactsyncProperty(context.Background()).ClientId(clientId).M365ContactSyncProperty(m365ContactSyncProperty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `M365ContactSyncPropertiesAPI.PostCompanyM365contactsyncProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyM365contactsyncProperty`: M365ContactSyncProperty
	fmt.Fprintf(os.Stdout, "Response from `M365ContactSyncPropertiesAPI.PostCompanyM365contactsyncProperty`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyM365contactsyncPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **m365ContactSyncProperty** | [**M365ContactSyncProperty**](M365ContactSyncProperty.md) | country | 

### Return type

[**M365ContactSyncProperty**](M365ContactSyncProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

