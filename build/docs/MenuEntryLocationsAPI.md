# \MenuEntryLocationsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemMenuEntriesByParentIdLocationsById**](MenuEntryLocationsAPI.md#DeleteSystemMenuEntriesByParentIdLocationsById) | **Delete** /system/menuEntries/{parentId}/locations/{id} | Delete MenuEntryLocation
[**GetSystemMenuEntriesByParentIdLocations**](MenuEntryLocationsAPI.md#GetSystemMenuEntriesByParentIdLocations) | **Get** /system/menuEntries/{parentId}/locations | Get List of MenuEntryLocation
[**GetSystemMenuEntriesByParentIdLocationsById**](MenuEntryLocationsAPI.md#GetSystemMenuEntriesByParentIdLocationsById) | **Get** /system/menuEntries/{parentId}/locations/{id} | Get MenuEntryLocation
[**GetSystemMenuEntriesByParentIdLocationsCount**](MenuEntryLocationsAPI.md#GetSystemMenuEntriesByParentIdLocationsCount) | **Get** /system/menuEntries/{parentId}/locations/count | Get Count of MenuEntryLocation
[**PostSystemMenuEntriesByParentIdLocations**](MenuEntryLocationsAPI.md#PostSystemMenuEntriesByParentIdLocations) | **Post** /system/menuEntries/{parentId}/locations | Post MenuEntryLocation



## DeleteSystemMenuEntriesByParentIdLocationsById

> DeleteSystemMenuEntriesByParentIdLocationsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete MenuEntryLocation

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
	id := int32(56) // int32 | locationId
	parentId := int32(56) // int32 | menuEntryId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MenuEntryLocationsAPI.DeleteSystemMenuEntriesByParentIdLocationsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MenuEntryLocationsAPI.DeleteSystemMenuEntriesByParentIdLocationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | locationId | 
**parentId** | **int32** | menuEntryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemMenuEntriesByParentIdLocationsByIdRequest struct via the builder pattern


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


## GetSystemMenuEntriesByParentIdLocations

> []MenuEntryLocation GetSystemMenuEntriesByParentIdLocations(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of MenuEntryLocation

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
	parentId := int32(56) // int32 | menuEntryId
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
	resp, r, err := apiClient.MenuEntryLocationsAPI.GetSystemMenuEntriesByParentIdLocations(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MenuEntryLocationsAPI.GetSystemMenuEntriesByParentIdLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMenuEntriesByParentIdLocations`: []MenuEntryLocation
	fmt.Fprintf(os.Stdout, "Response from `MenuEntryLocationsAPI.GetSystemMenuEntriesByParentIdLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | menuEntryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMenuEntriesByParentIdLocationsRequest struct via the builder pattern


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

[**[]MenuEntryLocation**](MenuEntryLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemMenuEntriesByParentIdLocationsById

> MenuEntryLocation GetSystemMenuEntriesByParentIdLocationsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get MenuEntryLocation

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
	id := int32(56) // int32 | locationId
	parentId := int32(56) // int32 | menuEntryId
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
	resp, r, err := apiClient.MenuEntryLocationsAPI.GetSystemMenuEntriesByParentIdLocationsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MenuEntryLocationsAPI.GetSystemMenuEntriesByParentIdLocationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMenuEntriesByParentIdLocationsById`: MenuEntryLocation
	fmt.Fprintf(os.Stdout, "Response from `MenuEntryLocationsAPI.GetSystemMenuEntriesByParentIdLocationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | locationId | 
**parentId** | **int32** | menuEntryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMenuEntriesByParentIdLocationsByIdRequest struct via the builder pattern


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

[**MenuEntryLocation**](MenuEntryLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemMenuEntriesByParentIdLocationsCount

> Count GetSystemMenuEntriesByParentIdLocationsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of MenuEntryLocation

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
	parentId := int32(56) // int32 | menuEntryId
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
	resp, r, err := apiClient.MenuEntryLocationsAPI.GetSystemMenuEntriesByParentIdLocationsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MenuEntryLocationsAPI.GetSystemMenuEntriesByParentIdLocationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMenuEntriesByParentIdLocationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `MenuEntryLocationsAPI.GetSystemMenuEntriesByParentIdLocationsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | menuEntryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMenuEntriesByParentIdLocationsCountRequest struct via the builder pattern


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


## PostSystemMenuEntriesByParentIdLocations

> MenuEntryLocation PostSystemMenuEntriesByParentIdLocations(ctx, parentId).ClientId(clientId).MenuEntryLocation(menuEntryLocation).Execute()

Post MenuEntryLocation

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
	parentId := int32(56) // int32 | menuEntryId
	clientId := "clientId_example" // string | 
	menuEntryLocation := *openapiclient.NewMenuEntryLocation(*openapiclient.NewSystemLocationReference()) // MenuEntryLocation | menuEntryLocation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MenuEntryLocationsAPI.PostSystemMenuEntriesByParentIdLocations(context.Background(), parentId).ClientId(clientId).MenuEntryLocation(menuEntryLocation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MenuEntryLocationsAPI.PostSystemMenuEntriesByParentIdLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemMenuEntriesByParentIdLocations`: MenuEntryLocation
	fmt.Fprintf(os.Stdout, "Response from `MenuEntryLocationsAPI.PostSystemMenuEntriesByParentIdLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | menuEntryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemMenuEntriesByParentIdLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **menuEntryLocation** | [**MenuEntryLocation**](MenuEntryLocation.md) | menuEntryLocation | 

### Return type

[**MenuEntryLocation**](MenuEntryLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

