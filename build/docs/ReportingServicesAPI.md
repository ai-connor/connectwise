# \ReportingServicesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemMycompanyReportingServices**](ReportingServicesAPI.md#GetSystemMycompanyReportingServices) | **Get** /system/mycompany/reportingServices | Get List of ReportingService
[**GetSystemMycompanyReportingServicesById**](ReportingServicesAPI.md#GetSystemMycompanyReportingServicesById) | **Get** /system/mycompany/reportingServices/{id} | Get ReportingService
[**PatchSystemMycompanyReportingServicesById**](ReportingServicesAPI.md#PatchSystemMycompanyReportingServicesById) | **Patch** /system/mycompany/reportingServices/{id} | Patch ReportingService
[**PostSystemMycompanyReportingServicesByIdTestConnection**](ReportingServicesAPI.md#PostSystemMycompanyReportingServicesByIdTestConnection) | **Post** /system/mycompany/reportingServices/{id}/testConnection | Post SuccessResponse
[**PutSystemMycompanyReportingServicesById**](ReportingServicesAPI.md#PutSystemMycompanyReportingServicesById) | **Put** /system/mycompany/reportingServices/{id} | Put ReportingService



## GetSystemMycompanyReportingServices

> []ReportingService GetSystemMycompanyReportingServices(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ReportingService

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
	resp, r, err := apiClient.ReportingServicesAPI.GetSystemMycompanyReportingServices(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingServicesAPI.GetSystemMycompanyReportingServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMycompanyReportingServices`: []ReportingService
	fmt.Fprintf(os.Stdout, "Response from `ReportingServicesAPI.GetSystemMycompanyReportingServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMycompanyReportingServicesRequest struct via the builder pattern


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

[**[]ReportingService**](ReportingService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemMycompanyReportingServicesById

> ReportingService GetSystemMycompanyReportingServicesById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ReportingService

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
	id := int32(56) // int32 | reportingServiceId
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
	resp, r, err := apiClient.ReportingServicesAPI.GetSystemMycompanyReportingServicesById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingServicesAPI.GetSystemMycompanyReportingServicesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMycompanyReportingServicesById`: ReportingService
	fmt.Fprintf(os.Stdout, "Response from `ReportingServicesAPI.GetSystemMycompanyReportingServicesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | reportingServiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMycompanyReportingServicesByIdRequest struct via the builder pattern


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

[**ReportingService**](ReportingService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSystemMycompanyReportingServicesById

> ReportingService PatchSystemMycompanyReportingServicesById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ReportingService

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
	id := int32(56) // int32 | reportingServiceId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportingServicesAPI.PatchSystemMycompanyReportingServicesById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingServicesAPI.PatchSystemMycompanyReportingServicesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemMycompanyReportingServicesById`: ReportingService
	fmt.Fprintf(os.Stdout, "Response from `ReportingServicesAPI.PatchSystemMycompanyReportingServicesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | reportingServiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemMycompanyReportingServicesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**ReportingService**](ReportingService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemMycompanyReportingServicesByIdTestConnection

> SuccessResponse PostSystemMycompanyReportingServicesByIdTestConnection(ctx, id).ClientId(clientId).Execute()

Post SuccessResponse

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
	id := int32(56) // int32 | reportingServiceId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportingServicesAPI.PostSystemMycompanyReportingServicesByIdTestConnection(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingServicesAPI.PostSystemMycompanyReportingServicesByIdTestConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemMycompanyReportingServicesByIdTestConnection`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportingServicesAPI.PostSystemMycompanyReportingServicesByIdTestConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | reportingServiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemMycompanyReportingServicesByIdTestConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemMycompanyReportingServicesById

> ReportingService PutSystemMycompanyReportingServicesById(ctx, id).ClientId(clientId).ReportingService(reportingService).Execute()

Put ReportingService

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
	id := int32(56) // int32 | reportingServiceId
	clientId := "clientId_example" // string | 
	reportingService := *openapiclient.NewReportingService() // ReportingService | service

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportingServicesAPI.PutSystemMycompanyReportingServicesById(context.Background(), id).ClientId(clientId).ReportingService(reportingService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportingServicesAPI.PutSystemMycompanyReportingServicesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemMycompanyReportingServicesById`: ReportingService
	fmt.Fprintf(os.Stdout, "Response from `ReportingServicesAPI.PutSystemMycompanyReportingServicesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | reportingServiceId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemMycompanyReportingServicesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **reportingService** | [**ReportingService**](ReportingService.md) | service | 

### Return type

[**ReportingService**](ReportingService.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

