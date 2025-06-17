# \WorkRoleLocationsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTimeWorkRolesByParentIdLocationsById**](WorkRoleLocationsAPI.md#DeleteTimeWorkRolesByParentIdLocationsById) | **Delete** /time/workRoles/{parentId}/locations/{id} | Delete WorkRoleLocation
[**GetTimeWorkRolesByParentIdLocations**](WorkRoleLocationsAPI.md#GetTimeWorkRolesByParentIdLocations) | **Get** /time/workRoles/{parentId}/locations | Get List of WorkRoleLocation
[**GetTimeWorkRolesByParentIdLocationsById**](WorkRoleLocationsAPI.md#GetTimeWorkRolesByParentIdLocationsById) | **Get** /time/workRoles/{parentId}/locations/{id} | Get WorkRoleLocation
[**GetTimeWorkRolesByParentIdLocationsCount**](WorkRoleLocationsAPI.md#GetTimeWorkRolesByParentIdLocationsCount) | **Get** /time/workRoles/{parentId}/locations/count | Get Count of WorkRoleLocation
[**PatchTimeWorkRolesByParentIdLocationsById**](WorkRoleLocationsAPI.md#PatchTimeWorkRolesByParentIdLocationsById) | **Patch** /time/workRoles/{parentId}/locations/{id} | Patch WorkRoleLocation
[**PostTimeWorkRolesByParentIdLocations**](WorkRoleLocationsAPI.md#PostTimeWorkRolesByParentIdLocations) | **Post** /time/workRoles/{parentId}/locations | Post WorkRoleLocation
[**PutTimeWorkRolesByParentIdLocationsById**](WorkRoleLocationsAPI.md#PutTimeWorkRolesByParentIdLocationsById) | **Put** /time/workRoles/{parentId}/locations/{id} | Put WorkRoleLocation



## DeleteTimeWorkRolesByParentIdLocationsById

> DeleteTimeWorkRolesByParentIdLocationsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete WorkRoleLocation

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
	parentId := int32(56) // int32 | workRoleId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkRoleLocationsAPI.DeleteTimeWorkRolesByParentIdLocationsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkRoleLocationsAPI.DeleteTimeWorkRolesByParentIdLocationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | locationId | 
**parentId** | **int32** | workRoleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTimeWorkRolesByParentIdLocationsByIdRequest struct via the builder pattern


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


## GetTimeWorkRolesByParentIdLocations

> []WorkRoleLocation GetTimeWorkRolesByParentIdLocations(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of WorkRoleLocation

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
	parentId := int32(56) // int32 | workRoleId
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
	resp, r, err := apiClient.WorkRoleLocationsAPI.GetTimeWorkRolesByParentIdLocations(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkRoleLocationsAPI.GetTimeWorkRolesByParentIdLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeWorkRolesByParentIdLocations`: []WorkRoleLocation
	fmt.Fprintf(os.Stdout, "Response from `WorkRoleLocationsAPI.GetTimeWorkRolesByParentIdLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | workRoleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeWorkRolesByParentIdLocationsRequest struct via the builder pattern


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

[**[]WorkRoleLocation**](WorkRoleLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeWorkRolesByParentIdLocationsById

> WorkRoleLocation GetTimeWorkRolesByParentIdLocationsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get WorkRoleLocation

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
	parentId := int32(56) // int32 | workRoleId
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
	resp, r, err := apiClient.WorkRoleLocationsAPI.GetTimeWorkRolesByParentIdLocationsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkRoleLocationsAPI.GetTimeWorkRolesByParentIdLocationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeWorkRolesByParentIdLocationsById`: WorkRoleLocation
	fmt.Fprintf(os.Stdout, "Response from `WorkRoleLocationsAPI.GetTimeWorkRolesByParentIdLocationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | locationId | 
**parentId** | **int32** | workRoleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeWorkRolesByParentIdLocationsByIdRequest struct via the builder pattern


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

[**WorkRoleLocation**](WorkRoleLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeWorkRolesByParentIdLocationsCount

> Count GetTimeWorkRolesByParentIdLocationsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of WorkRoleLocation

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
	parentId := int32(56) // int32 | workRoleId
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
	resp, r, err := apiClient.WorkRoleLocationsAPI.GetTimeWorkRolesByParentIdLocationsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkRoleLocationsAPI.GetTimeWorkRolesByParentIdLocationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeWorkRolesByParentIdLocationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `WorkRoleLocationsAPI.GetTimeWorkRolesByParentIdLocationsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | workRoleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeWorkRolesByParentIdLocationsCountRequest struct via the builder pattern


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


## PatchTimeWorkRolesByParentIdLocationsById

> WorkRoleLocation PatchTimeWorkRolesByParentIdLocationsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch WorkRoleLocation

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
	parentId := int32(56) // int32 | workRoleId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkRoleLocationsAPI.PatchTimeWorkRolesByParentIdLocationsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkRoleLocationsAPI.PatchTimeWorkRolesByParentIdLocationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTimeWorkRolesByParentIdLocationsById`: WorkRoleLocation
	fmt.Fprintf(os.Stdout, "Response from `WorkRoleLocationsAPI.PatchTimeWorkRolesByParentIdLocationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | locationId | 
**parentId** | **int32** | workRoleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTimeWorkRolesByParentIdLocationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**WorkRoleLocation**](WorkRoleLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTimeWorkRolesByParentIdLocations

> WorkRoleLocation PostTimeWorkRolesByParentIdLocations(ctx, parentId).ClientId(clientId).WorkRoleLocation(workRoleLocation).Execute()

Post WorkRoleLocation

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
	parentId := int32(56) // int32 | workRoleId
	clientId := "clientId_example" // string | 
	workRoleLocation := *openapiclient.NewWorkRoleLocation(*openapiclient.NewSystemLocationReference()) // WorkRoleLocation | workRoleLocation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkRoleLocationsAPI.PostTimeWorkRolesByParentIdLocations(context.Background(), parentId).ClientId(clientId).WorkRoleLocation(workRoleLocation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkRoleLocationsAPI.PostTimeWorkRolesByParentIdLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTimeWorkRolesByParentIdLocations`: WorkRoleLocation
	fmt.Fprintf(os.Stdout, "Response from `WorkRoleLocationsAPI.PostTimeWorkRolesByParentIdLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | workRoleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTimeWorkRolesByParentIdLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **workRoleLocation** | [**WorkRoleLocation**](WorkRoleLocation.md) | workRoleLocation | 

### Return type

[**WorkRoleLocation**](WorkRoleLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTimeWorkRolesByParentIdLocationsById

> WorkRoleLocation PutTimeWorkRolesByParentIdLocationsById(ctx, id, parentId).ClientId(clientId).WorkRoleLocation(workRoleLocation).Execute()

Put WorkRoleLocation

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
	parentId := int32(56) // int32 | workRoleId
	clientId := "clientId_example" // string | 
	workRoleLocation := *openapiclient.NewWorkRoleLocation(*openapiclient.NewSystemLocationReference()) // WorkRoleLocation | workRoleLocation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkRoleLocationsAPI.PutTimeWorkRolesByParentIdLocationsById(context.Background(), id, parentId).ClientId(clientId).WorkRoleLocation(workRoleLocation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkRoleLocationsAPI.PutTimeWorkRolesByParentIdLocationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTimeWorkRolesByParentIdLocationsById`: WorkRoleLocation
	fmt.Fprintf(os.Stdout, "Response from `WorkRoleLocationsAPI.PutTimeWorkRolesByParentIdLocationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | locationId | 
**parentId** | **int32** | workRoleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTimeWorkRolesByParentIdLocationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **workRoleLocation** | [**WorkRoleLocation**](WorkRoleLocation.md) | workRoleLocation | 

### Return type

[**WorkRoleLocation**](WorkRoleLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

