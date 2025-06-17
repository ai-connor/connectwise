# \CorporateStructuresAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemMyCompanyCorporateStructure**](CorporateStructuresAPI.md#GetSystemMyCompanyCorporateStructure) | **Get** /system/myCompany/corporateStructure | Get List of CorporateStructure
[**GetSystemMyCompanyCorporateStructureById**](CorporateStructuresAPI.md#GetSystemMyCompanyCorporateStructureById) | **Get** /system/myCompany/corporateStructure/{id} | Get CorporateStructure
[**GetSystemMyCompanyCorporateStructureCount**](CorporateStructuresAPI.md#GetSystemMyCompanyCorporateStructureCount) | **Get** /system/myCompany/corporateStructure/count | Get Count of CorporateStructure
[**PatchSystemMyCompanyCorporateStructureById**](CorporateStructuresAPI.md#PatchSystemMyCompanyCorporateStructureById) | **Patch** /system/myCompany/corporateStructure/{id} | Patch CorporateStructure
[**PutSystemMyCompanyCorporateStructureById**](CorporateStructuresAPI.md#PutSystemMyCompanyCorporateStructureById) | **Put** /system/myCompany/corporateStructure/{id} | Put CorporateStructure



## GetSystemMyCompanyCorporateStructure

> []CorporateStructure GetSystemMyCompanyCorporateStructure(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of CorporateStructure

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
	resp, r, err := apiClient.CorporateStructuresAPI.GetSystemMyCompanyCorporateStructure(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CorporateStructuresAPI.GetSystemMyCompanyCorporateStructure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMyCompanyCorporateStructure`: []CorporateStructure
	fmt.Fprintf(os.Stdout, "Response from `CorporateStructuresAPI.GetSystemMyCompanyCorporateStructure`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMyCompanyCorporateStructureRequest struct via the builder pattern


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

[**[]CorporateStructure**](CorporateStructure.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemMyCompanyCorporateStructureById

> CorporateStructure GetSystemMyCompanyCorporateStructureById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get CorporateStructure

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
	id := int32(56) // int32 | corporateStructureId
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
	resp, r, err := apiClient.CorporateStructuresAPI.GetSystemMyCompanyCorporateStructureById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CorporateStructuresAPI.GetSystemMyCompanyCorporateStructureById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMyCompanyCorporateStructureById`: CorporateStructure
	fmt.Fprintf(os.Stdout, "Response from `CorporateStructuresAPI.GetSystemMyCompanyCorporateStructureById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | corporateStructureId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMyCompanyCorporateStructureByIdRequest struct via the builder pattern


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

[**CorporateStructure**](CorporateStructure.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemMyCompanyCorporateStructureCount

> Count GetSystemMyCompanyCorporateStructureCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of CorporateStructure

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
	resp, r, err := apiClient.CorporateStructuresAPI.GetSystemMyCompanyCorporateStructureCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CorporateStructuresAPI.GetSystemMyCompanyCorporateStructureCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMyCompanyCorporateStructureCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `CorporateStructuresAPI.GetSystemMyCompanyCorporateStructureCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMyCompanyCorporateStructureCountRequest struct via the builder pattern


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


## PatchSystemMyCompanyCorporateStructureById

> CorporateStructure PatchSystemMyCompanyCorporateStructureById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch CorporateStructure

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
	id := int32(56) // int32 | corporateStructureId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CorporateStructuresAPI.PatchSystemMyCompanyCorporateStructureById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CorporateStructuresAPI.PatchSystemMyCompanyCorporateStructureById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemMyCompanyCorporateStructureById`: CorporateStructure
	fmt.Fprintf(os.Stdout, "Response from `CorporateStructuresAPI.PatchSystemMyCompanyCorporateStructureById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | corporateStructureId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemMyCompanyCorporateStructureByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**CorporateStructure**](CorporateStructure.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemMyCompanyCorporateStructureById

> CorporateStructure PutSystemMyCompanyCorporateStructureById(ctx, id).ClientId(clientId).CorporateStructure(corporateStructure).Execute()

Put CorporateStructure

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
	id := int32(56) // int32 | corporateStructureId
	clientId := "clientId_example" // string | 
	corporateStructure := *openapiclient.NewCorporateStructure("FiscalYearStart_example", "LocationCaption_example", "GroupCaption_example", *openapiclient.NewCurrencyReference()) // CorporateStructure | corporateStructure

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CorporateStructuresAPI.PutSystemMyCompanyCorporateStructureById(context.Background(), id).ClientId(clientId).CorporateStructure(corporateStructure).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CorporateStructuresAPI.PutSystemMyCompanyCorporateStructureById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemMyCompanyCorporateStructureById`: CorporateStructure
	fmt.Fprintf(os.Stdout, "Response from `CorporateStructuresAPI.PutSystemMyCompanyCorporateStructureById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | corporateStructureId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemMyCompanyCorporateStructureByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **corporateStructure** | [**CorporateStructure**](CorporateStructure.md) | corporateStructure | 

### Return type

[**CorporateStructure**](CorporateStructure.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

