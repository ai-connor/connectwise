# \CompanyPickerItemsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyCompanyPickerItemsById**](CompanyPickerItemsAPI.md#DeleteCompanyCompanyPickerItemsById) | **Delete** /company/companyPickerItems/{id} | Delete CompanyPickerItem
[**GetCompanyCompanyPickerItems**](CompanyPickerItemsAPI.md#GetCompanyCompanyPickerItems) | **Get** /company/companyPickerItems | Get List of CompanyPickerItem
[**GetCompanyCompanyPickerItemsById**](CompanyPickerItemsAPI.md#GetCompanyCompanyPickerItemsById) | **Get** /company/companyPickerItems/{id} | Get CompanyPickerItem
[**GetCompanyCompanyPickerItemsCount**](CompanyPickerItemsAPI.md#GetCompanyCompanyPickerItemsCount) | **Get** /company/companyPickerItems/count | Get Count of CompanyPickerItem
[**PostCompanyCompanyPickerItems**](CompanyPickerItemsAPI.md#PostCompanyCompanyPickerItems) | **Post** /company/companyPickerItems | Post CompanyPickerItem
[**PostCompanyCompanyPickerItemsClear**](CompanyPickerItemsAPI.md#PostCompanyCompanyPickerItemsClear) | **Post** /company/companyPickerItems/clear | Post ClearPickerRequest



## DeleteCompanyCompanyPickerItemsById

> DeleteCompanyCompanyPickerItemsById(ctx, id).ClientId(clientId).Execute()

Delete CompanyPickerItem

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
	id := int32(56) // int32 | companyPickerItemId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CompanyPickerItemsAPI.DeleteCompanyCompanyPickerItemsById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyPickerItemsAPI.DeleteCompanyCompanyPickerItemsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | companyPickerItemId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyCompanyPickerItemsByIdRequest struct via the builder pattern


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


## GetCompanyCompanyPickerItems

> []CompanyPickerItem GetCompanyCompanyPickerItems(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of CompanyPickerItem

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
	resp, r, err := apiClient.CompanyPickerItemsAPI.GetCompanyCompanyPickerItems(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyPickerItemsAPI.GetCompanyCompanyPickerItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCompanyPickerItems`: []CompanyPickerItem
	fmt.Fprintf(os.Stdout, "Response from `CompanyPickerItemsAPI.GetCompanyCompanyPickerItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCompanyPickerItemsRequest struct via the builder pattern


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

[**[]CompanyPickerItem**](CompanyPickerItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyCompanyPickerItemsById

> CompanyPickerItem GetCompanyCompanyPickerItemsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get CompanyPickerItem

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
	id := int32(56) // int32 | companyPickerItemId
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
	resp, r, err := apiClient.CompanyPickerItemsAPI.GetCompanyCompanyPickerItemsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyPickerItemsAPI.GetCompanyCompanyPickerItemsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCompanyPickerItemsById`: CompanyPickerItem
	fmt.Fprintf(os.Stdout, "Response from `CompanyPickerItemsAPI.GetCompanyCompanyPickerItemsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | companyPickerItemId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCompanyPickerItemsByIdRequest struct via the builder pattern


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

[**CompanyPickerItem**](CompanyPickerItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyCompanyPickerItemsCount

> Count GetCompanyCompanyPickerItemsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of CompanyPickerItem

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
	resp, r, err := apiClient.CompanyPickerItemsAPI.GetCompanyCompanyPickerItemsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyPickerItemsAPI.GetCompanyCompanyPickerItemsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCompanyPickerItemsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `CompanyPickerItemsAPI.GetCompanyCompanyPickerItemsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCompanyPickerItemsCountRequest struct via the builder pattern


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


## PostCompanyCompanyPickerItems

> CompanyPickerItem PostCompanyCompanyPickerItems(ctx).ClientId(clientId).CompanyPickerItem(companyPickerItem).Execute()

Post CompanyPickerItem

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
	companyPickerItem := *openapiclient.NewCompanyPickerItem(*openapiclient.NewCompanyReference()) // CompanyPickerItem | companyPickerItem

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyPickerItemsAPI.PostCompanyCompanyPickerItems(context.Background()).ClientId(clientId).CompanyPickerItem(companyPickerItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyPickerItemsAPI.PostCompanyCompanyPickerItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyCompanyPickerItems`: CompanyPickerItem
	fmt.Fprintf(os.Stdout, "Response from `CompanyPickerItemsAPI.PostCompanyCompanyPickerItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyCompanyPickerItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **companyPickerItem** | [**CompanyPickerItem**](CompanyPickerItem.md) | companyPickerItem | 

### Return type

[**CompanyPickerItem**](CompanyPickerItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyCompanyPickerItemsClear

> ClearPickerRequest PostCompanyCompanyPickerItemsClear(ctx, clearPickerRequest).ClientId(clientId).Execute()

Post ClearPickerRequest

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
	clearPickerRequest := *openapiclient.NewClearPickerRequest() // ClearPickerRequest | clearPickerRequest

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyPickerItemsAPI.PostCompanyCompanyPickerItemsClear(context.Background(), clearPickerRequest).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyPickerItemsAPI.PostCompanyCompanyPickerItemsClear``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyCompanyPickerItemsClear`: ClearPickerRequest
	fmt.Fprintf(os.Stdout, "Response from `CompanyPickerItemsAPI.PostCompanyCompanyPickerItemsClear`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clearPickerRequest** | [**ClearPickerRequest**](.md) | clearPickerRequest | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyCompanyPickerItemsClearRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 


### Return type

[**ClearPickerRequest**](ClearPickerRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

