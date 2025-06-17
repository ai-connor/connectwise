# \UnitOfMeasureConversionsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProcurementUnitOfMeasuresByParentIdConversionsById**](UnitOfMeasureConversionsAPI.md#DeleteProcurementUnitOfMeasuresByParentIdConversionsById) | **Delete** /procurement/unitOfMeasures/{parentId}/conversions/{id} | Delete Conversion
[**GetProcurementUnitOfMeasuresByParentIdConversions**](UnitOfMeasureConversionsAPI.md#GetProcurementUnitOfMeasuresByParentIdConversions) | **Get** /procurement/unitOfMeasures/{parentId}/conversions | Get List of Conversion
[**GetProcurementUnitOfMeasuresByParentIdConversionsById**](UnitOfMeasureConversionsAPI.md#GetProcurementUnitOfMeasuresByParentIdConversionsById) | **Get** /procurement/unitOfMeasures/{parentId}/conversions/{id} | Get Conversion
[**GetProcurementUnitOfMeasuresByParentIdConversionsCount**](UnitOfMeasureConversionsAPI.md#GetProcurementUnitOfMeasuresByParentIdConversionsCount) | **Get** /procurement/unitOfMeasures/{parentId}/conversions/count | Get Count of Conversion
[**PatchProcurementUnitOfMeasuresByParentIdConversionsById**](UnitOfMeasureConversionsAPI.md#PatchProcurementUnitOfMeasuresByParentIdConversionsById) | **Patch** /procurement/unitOfMeasures/{parentId}/conversions/{id} | Patch Conversion
[**PostProcurementUnitOfMeasuresByParentIdConversions**](UnitOfMeasureConversionsAPI.md#PostProcurementUnitOfMeasuresByParentIdConversions) | **Post** /procurement/unitOfMeasures/{parentId}/conversions | Post Conversion
[**PutProcurementUnitOfMeasuresByParentIdConversionsById**](UnitOfMeasureConversionsAPI.md#PutProcurementUnitOfMeasuresByParentIdConversionsById) | **Put** /procurement/unitOfMeasures/{parentId}/conversions/{id} | Put Conversion



## DeleteProcurementUnitOfMeasuresByParentIdConversionsById

> DeleteProcurementUnitOfMeasuresByParentIdConversionsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete Conversion

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
	id := int32(56) // int32 | conversionId
	parentId := int32(56) // int32 | unitOfMeasureId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UnitOfMeasureConversionsAPI.DeleteProcurementUnitOfMeasuresByParentIdConversionsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnitOfMeasureConversionsAPI.DeleteProcurementUnitOfMeasuresByParentIdConversionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | conversionId | 
**parentId** | **int32** | unitOfMeasureId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcurementUnitOfMeasuresByParentIdConversionsByIdRequest struct via the builder pattern


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


## GetProcurementUnitOfMeasuresByParentIdConversions

> []Conversion GetProcurementUnitOfMeasuresByParentIdConversions(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of Conversion

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
	parentId := int32(56) // int32 | unitOfMeasureId
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
	resp, r, err := apiClient.UnitOfMeasureConversionsAPI.GetProcurementUnitOfMeasuresByParentIdConversions(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnitOfMeasureConversionsAPI.GetProcurementUnitOfMeasuresByParentIdConversions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementUnitOfMeasuresByParentIdConversions`: []Conversion
	fmt.Fprintf(os.Stdout, "Response from `UnitOfMeasureConversionsAPI.GetProcurementUnitOfMeasuresByParentIdConversions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | unitOfMeasureId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementUnitOfMeasuresByParentIdConversionsRequest struct via the builder pattern


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

[**[]Conversion**](Conversion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementUnitOfMeasuresByParentIdConversionsById

> Conversion GetProcurementUnitOfMeasuresByParentIdConversionsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Conversion

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
	id := int32(56) // int32 | conversionId
	parentId := int32(56) // int32 | unitOfMeasureId
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
	resp, r, err := apiClient.UnitOfMeasureConversionsAPI.GetProcurementUnitOfMeasuresByParentIdConversionsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnitOfMeasureConversionsAPI.GetProcurementUnitOfMeasuresByParentIdConversionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementUnitOfMeasuresByParentIdConversionsById`: Conversion
	fmt.Fprintf(os.Stdout, "Response from `UnitOfMeasureConversionsAPI.GetProcurementUnitOfMeasuresByParentIdConversionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | conversionId | 
**parentId** | **int32** | unitOfMeasureId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementUnitOfMeasuresByParentIdConversionsByIdRequest struct via the builder pattern


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

[**Conversion**](Conversion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProcurementUnitOfMeasuresByParentIdConversionsCount

> Count GetProcurementUnitOfMeasuresByParentIdConversionsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of Conversion

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
	parentId := int32(56) // int32 | unitOfMeasureId
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
	resp, r, err := apiClient.UnitOfMeasureConversionsAPI.GetProcurementUnitOfMeasuresByParentIdConversionsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnitOfMeasureConversionsAPI.GetProcurementUnitOfMeasuresByParentIdConversionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProcurementUnitOfMeasuresByParentIdConversionsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `UnitOfMeasureConversionsAPI.GetProcurementUnitOfMeasuresByParentIdConversionsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | unitOfMeasureId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcurementUnitOfMeasuresByParentIdConversionsCountRequest struct via the builder pattern


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


## PatchProcurementUnitOfMeasuresByParentIdConversionsById

> Conversion PatchProcurementUnitOfMeasuresByParentIdConversionsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch Conversion

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
	id := int32(56) // int32 | conversionId
	parentId := int32(56) // int32 | unitOfMeasureId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UnitOfMeasureConversionsAPI.PatchProcurementUnitOfMeasuresByParentIdConversionsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnitOfMeasureConversionsAPI.PatchProcurementUnitOfMeasuresByParentIdConversionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProcurementUnitOfMeasuresByParentIdConversionsById`: Conversion
	fmt.Fprintf(os.Stdout, "Response from `UnitOfMeasureConversionsAPI.PatchProcurementUnitOfMeasuresByParentIdConversionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | conversionId | 
**parentId** | **int32** | unitOfMeasureId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProcurementUnitOfMeasuresByParentIdConversionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**Conversion**](Conversion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProcurementUnitOfMeasuresByParentIdConversions

> Conversion PostProcurementUnitOfMeasuresByParentIdConversions(ctx, parentId).ClientId(clientId).Conversion(conversion).Execute()

Post Conversion

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
	parentId := int32(56) // int32 | unitOfMeasureId
	clientId := "clientId_example" // string | 
	conversion := *openapiclient.NewConversion(*openapiclient.NewUnitOfMeasureReference()) // Conversion | conversion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UnitOfMeasureConversionsAPI.PostProcurementUnitOfMeasuresByParentIdConversions(context.Background(), parentId).ClientId(clientId).Conversion(conversion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnitOfMeasureConversionsAPI.PostProcurementUnitOfMeasuresByParentIdConversions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProcurementUnitOfMeasuresByParentIdConversions`: Conversion
	fmt.Fprintf(os.Stdout, "Response from `UnitOfMeasureConversionsAPI.PostProcurementUnitOfMeasuresByParentIdConversions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | unitOfMeasureId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProcurementUnitOfMeasuresByParentIdConversionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **conversion** | [**Conversion**](Conversion.md) | conversion | 

### Return type

[**Conversion**](Conversion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProcurementUnitOfMeasuresByParentIdConversionsById

> Conversion PutProcurementUnitOfMeasuresByParentIdConversionsById(ctx, id, parentId).ClientId(clientId).Conversion(conversion).Execute()

Put Conversion

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
	id := int32(56) // int32 | conversionId
	parentId := int32(56) // int32 | unitOfMeasureId
	clientId := "clientId_example" // string | 
	conversion := *openapiclient.NewConversion(*openapiclient.NewUnitOfMeasureReference()) // Conversion | conversion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UnitOfMeasureConversionsAPI.PutProcurementUnitOfMeasuresByParentIdConversionsById(context.Background(), id, parentId).ClientId(clientId).Conversion(conversion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UnitOfMeasureConversionsAPI.PutProcurementUnitOfMeasuresByParentIdConversionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProcurementUnitOfMeasuresByParentIdConversionsById`: Conversion
	fmt.Fprintf(os.Stdout, "Response from `UnitOfMeasureConversionsAPI.PutProcurementUnitOfMeasuresByParentIdConversionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | conversionId | 
**parentId** | **int32** | unitOfMeasureId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProcurementUnitOfMeasuresByParentIdConversionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **conversion** | [**Conversion**](Conversion.md) | conversion | 

### Return type

[**Conversion**](Conversion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

