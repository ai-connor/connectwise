# \DepartmentLocationsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemDepartmentsByParentIdLocationsById**](DepartmentLocationsAPI.md#DeleteSystemDepartmentsByParentIdLocationsById) | **Delete** /system/departments/{parentId}/locations/{id} | Delete DepartmentLocation
[**GetSystemDepartmentsByParentIdLocations**](DepartmentLocationsAPI.md#GetSystemDepartmentsByParentIdLocations) | **Get** /system/departments/{parentId}/locations | Get List of DepartmentLocation
[**GetSystemDepartmentsByParentIdLocationsById**](DepartmentLocationsAPI.md#GetSystemDepartmentsByParentIdLocationsById) | **Get** /system/departments/{parentId}/locations/{id} | Get DepartmentLocation
[**GetSystemDepartmentsByParentIdLocationsCount**](DepartmentLocationsAPI.md#GetSystemDepartmentsByParentIdLocationsCount) | **Get** /system/departments/{parentId}/locations/count | Get Count of DepartmentLocation
[**PatchSystemDepartmentsByParentIdLocationsById**](DepartmentLocationsAPI.md#PatchSystemDepartmentsByParentIdLocationsById) | **Patch** /system/departments/{parentId}/locations/{id} | Patch DepartmentLocation
[**PostSystemDepartmentsByParentIdLocations**](DepartmentLocationsAPI.md#PostSystemDepartmentsByParentIdLocations) | **Post** /system/departments/{parentId}/locations | Post DepartmentLocation
[**PutSystemDepartmentsByParentIdLocationsById**](DepartmentLocationsAPI.md#PutSystemDepartmentsByParentIdLocationsById) | **Put** /system/departments/{parentId}/locations/{id} | Put DepartmentLocation



## DeleteSystemDepartmentsByParentIdLocationsById

> DeleteSystemDepartmentsByParentIdLocationsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete DepartmentLocation

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
	parentId := int32(56) // int32 | departmentId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DepartmentLocationsAPI.DeleteSystemDepartmentsByParentIdLocationsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentLocationsAPI.DeleteSystemDepartmentsByParentIdLocationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | locationId | 
**parentId** | **int32** | departmentId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemDepartmentsByParentIdLocationsByIdRequest struct via the builder pattern


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


## GetSystemDepartmentsByParentIdLocations

> []DepartmentLocation GetSystemDepartmentsByParentIdLocations(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of DepartmentLocation

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
	parentId := int32(56) // int32 | departmentId
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
	resp, r, err := apiClient.DepartmentLocationsAPI.GetSystemDepartmentsByParentIdLocations(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentLocationsAPI.GetSystemDepartmentsByParentIdLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemDepartmentsByParentIdLocations`: []DepartmentLocation
	fmt.Fprintf(os.Stdout, "Response from `DepartmentLocationsAPI.GetSystemDepartmentsByParentIdLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | departmentId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemDepartmentsByParentIdLocationsRequest struct via the builder pattern


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

[**[]DepartmentLocation**](DepartmentLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemDepartmentsByParentIdLocationsById

> DepartmentLocation GetSystemDepartmentsByParentIdLocationsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get DepartmentLocation

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
	parentId := int32(56) // int32 | departmentId
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
	resp, r, err := apiClient.DepartmentLocationsAPI.GetSystemDepartmentsByParentIdLocationsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentLocationsAPI.GetSystemDepartmentsByParentIdLocationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemDepartmentsByParentIdLocationsById`: DepartmentLocation
	fmt.Fprintf(os.Stdout, "Response from `DepartmentLocationsAPI.GetSystemDepartmentsByParentIdLocationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | locationId | 
**parentId** | **int32** | departmentId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemDepartmentsByParentIdLocationsByIdRequest struct via the builder pattern


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

[**DepartmentLocation**](DepartmentLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemDepartmentsByParentIdLocationsCount

> Count GetSystemDepartmentsByParentIdLocationsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of DepartmentLocation

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
	parentId := int32(56) // int32 | departmentId
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
	resp, r, err := apiClient.DepartmentLocationsAPI.GetSystemDepartmentsByParentIdLocationsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentLocationsAPI.GetSystemDepartmentsByParentIdLocationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemDepartmentsByParentIdLocationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `DepartmentLocationsAPI.GetSystemDepartmentsByParentIdLocationsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | departmentId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemDepartmentsByParentIdLocationsCountRequest struct via the builder pattern


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


## PatchSystemDepartmentsByParentIdLocationsById

> DepartmentLocation PatchSystemDepartmentsByParentIdLocationsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch DepartmentLocation

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
	parentId := int32(56) // int32 | departmentId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DepartmentLocationsAPI.PatchSystemDepartmentsByParentIdLocationsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentLocationsAPI.PatchSystemDepartmentsByParentIdLocationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemDepartmentsByParentIdLocationsById`: DepartmentLocation
	fmt.Fprintf(os.Stdout, "Response from `DepartmentLocationsAPI.PatchSystemDepartmentsByParentIdLocationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | locationId | 
**parentId** | **int32** | departmentId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemDepartmentsByParentIdLocationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**DepartmentLocation**](DepartmentLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemDepartmentsByParentIdLocations

> DepartmentLocation PostSystemDepartmentsByParentIdLocations(ctx, parentId).ClientId(clientId).DepartmentLocation(departmentLocation).Execute()

Post DepartmentLocation

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
	parentId := int32(56) // int32 | departmentId
	clientId := "clientId_example" // string | 
	departmentLocation := *openapiclient.NewDepartmentLocation(*openapiclient.NewSystemLocationReference()) // DepartmentLocation | departmentLocation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DepartmentLocationsAPI.PostSystemDepartmentsByParentIdLocations(context.Background(), parentId).ClientId(clientId).DepartmentLocation(departmentLocation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentLocationsAPI.PostSystemDepartmentsByParentIdLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemDepartmentsByParentIdLocations`: DepartmentLocation
	fmt.Fprintf(os.Stdout, "Response from `DepartmentLocationsAPI.PostSystemDepartmentsByParentIdLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | departmentId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemDepartmentsByParentIdLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **departmentLocation** | [**DepartmentLocation**](DepartmentLocation.md) | departmentLocation | 

### Return type

[**DepartmentLocation**](DepartmentLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemDepartmentsByParentIdLocationsById

> DepartmentLocation PutSystemDepartmentsByParentIdLocationsById(ctx, id, parentId).ClientId(clientId).DepartmentLocation(departmentLocation).Execute()

Put DepartmentLocation

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
	parentId := int32(56) // int32 | departmentId
	clientId := "clientId_example" // string | 
	departmentLocation := *openapiclient.NewDepartmentLocation(*openapiclient.NewSystemLocationReference()) // DepartmentLocation | departmentLocation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DepartmentLocationsAPI.PutSystemDepartmentsByParentIdLocationsById(context.Background(), id, parentId).ClientId(clientId).DepartmentLocation(departmentLocation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentLocationsAPI.PutSystemDepartmentsByParentIdLocationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemDepartmentsByParentIdLocationsById`: DepartmentLocation
	fmt.Fprintf(os.Stdout, "Response from `DepartmentLocationsAPI.PutSystemDepartmentsByParentIdLocationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | locationId | 
**parentId** | **int32** | departmentId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemDepartmentsByParentIdLocationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **departmentLocation** | [**DepartmentLocation**](DepartmentLocation.md) | departmentLocation | 

### Return type

[**DepartmentLocation**](DepartmentLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

