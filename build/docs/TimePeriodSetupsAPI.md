# \TimePeriodSetupsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTimeTimePeriodSetupsById**](TimePeriodSetupsAPI.md#DeleteTimeTimePeriodSetupsById) | **Delete** /time/timePeriodSetups/{id} | Delete TimePeriodSetup
[**GetTimeTimePeriodSetups**](TimePeriodSetupsAPI.md#GetTimeTimePeriodSetups) | **Get** /time/timePeriodSetups | Get List of TimePeriodSetup
[**GetTimeTimePeriodSetupsById**](TimePeriodSetupsAPI.md#GetTimeTimePeriodSetupsById) | **Get** /time/timePeriodSetups/{id} | Get TimePeriodSetup
[**GetTimeTimePeriodSetupsCount**](TimePeriodSetupsAPI.md#GetTimeTimePeriodSetupsCount) | **Get** /time/timePeriodSetups/count | Get Count of TimePeriodSetup
[**GetTimeTimePeriodSetupsDefault**](TimePeriodSetupsAPI.md#GetTimeTimePeriodSetupsDefault) | **Get** /time/timePeriodSetups/default | Get TimePeriodSetupDefaults
[**PatchTimeTimePeriodSetupsById**](TimePeriodSetupsAPI.md#PatchTimeTimePeriodSetupsById) | **Patch** /time/timePeriodSetups/{id} | Patch TimePeriodSetup
[**PostTimeTimePeriodSetups**](TimePeriodSetupsAPI.md#PostTimeTimePeriodSetups) | **Post** /time/timePeriodSetups | Post TimePeriodSetup
[**PutTimeTimePeriodSetupsById**](TimePeriodSetupsAPI.md#PutTimeTimePeriodSetupsById) | **Put** /time/timePeriodSetups/{id} | Put TimePeriodSetup



## DeleteTimeTimePeriodSetupsById

> DeleteTimeTimePeriodSetupsById(ctx, id).ClientId(clientId).Execute()

Delete TimePeriodSetup

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
	id := int32(56) // int32 | timePeriodSetupId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TimePeriodSetupsAPI.DeleteTimeTimePeriodSetupsById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimePeriodSetupsAPI.DeleteTimeTimePeriodSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | timePeriodSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTimeTimePeriodSetupsByIdRequest struct via the builder pattern


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


## GetTimeTimePeriodSetups

> []TimePeriodSetup GetTimeTimePeriodSetups(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of TimePeriodSetup

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
	resp, r, err := apiClient.TimePeriodSetupsAPI.GetTimeTimePeriodSetups(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimePeriodSetupsAPI.GetTimeTimePeriodSetups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeTimePeriodSetups`: []TimePeriodSetup
	fmt.Fprintf(os.Stdout, "Response from `TimePeriodSetupsAPI.GetTimeTimePeriodSetups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeTimePeriodSetupsRequest struct via the builder pattern


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

[**[]TimePeriodSetup**](TimePeriodSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeTimePeriodSetupsById

> TimePeriodSetup GetTimeTimePeriodSetupsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get TimePeriodSetup

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
	id := int32(56) // int32 | timePeriodSetupId
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
	resp, r, err := apiClient.TimePeriodSetupsAPI.GetTimeTimePeriodSetupsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimePeriodSetupsAPI.GetTimeTimePeriodSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeTimePeriodSetupsById`: TimePeriodSetup
	fmt.Fprintf(os.Stdout, "Response from `TimePeriodSetupsAPI.GetTimeTimePeriodSetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | timePeriodSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeTimePeriodSetupsByIdRequest struct via the builder pattern


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

[**TimePeriodSetup**](TimePeriodSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeTimePeriodSetupsCount

> Count GetTimeTimePeriodSetupsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of TimePeriodSetup

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
	resp, r, err := apiClient.TimePeriodSetupsAPI.GetTimeTimePeriodSetupsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimePeriodSetupsAPI.GetTimeTimePeriodSetupsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeTimePeriodSetupsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `TimePeriodSetupsAPI.GetTimeTimePeriodSetupsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeTimePeriodSetupsCountRequest struct via the builder pattern


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


## GetTimeTimePeriodSetupsDefault

> map[string]interface{} GetTimeTimePeriodSetupsDefault(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get TimePeriodSetupDefaults

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
	resp, r, err := apiClient.TimePeriodSetupsAPI.GetTimeTimePeriodSetupsDefault(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimePeriodSetupsAPI.GetTimeTimePeriodSetupsDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeTimePeriodSetupsDefault`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TimePeriodSetupsAPI.GetTimeTimePeriodSetupsDefault`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeTimePeriodSetupsDefaultRequest struct via the builder pattern


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

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTimeTimePeriodSetupsById

> TimePeriodSetup PatchTimeTimePeriodSetupsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch TimePeriodSetup

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
	id := int32(56) // int32 | timePeriodSetupId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimePeriodSetupsAPI.PatchTimeTimePeriodSetupsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimePeriodSetupsAPI.PatchTimeTimePeriodSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTimeTimePeriodSetupsById`: TimePeriodSetup
	fmt.Fprintf(os.Stdout, "Response from `TimePeriodSetupsAPI.PatchTimeTimePeriodSetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | timePeriodSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTimeTimePeriodSetupsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**TimePeriodSetup**](TimePeriodSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTimeTimePeriodSetups

> TimePeriodSetup PostTimeTimePeriodSetups(ctx).ClientId(clientId).TimePeriodSetup(timePeriodSetup).Execute()

Post TimePeriodSetup

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ai-connor/connectwise"
)

func main() {
	clientId := "clientId_example" // string | 
	timePeriodSetup := *openapiclient.NewTimePeriodSetup("PeriodApplyTo_example", NullableInt32(123), NullableInt32(123), "Type_example", time.Now(), NullableInt32(123)) // TimePeriodSetup | timePeriodSetup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimePeriodSetupsAPI.PostTimeTimePeriodSetups(context.Background()).ClientId(clientId).TimePeriodSetup(timePeriodSetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimePeriodSetupsAPI.PostTimeTimePeriodSetups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTimeTimePeriodSetups`: TimePeriodSetup
	fmt.Fprintf(os.Stdout, "Response from `TimePeriodSetupsAPI.PostTimeTimePeriodSetups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTimeTimePeriodSetupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **timePeriodSetup** | [**TimePeriodSetup**](TimePeriodSetup.md) | timePeriodSetup | 

### Return type

[**TimePeriodSetup**](TimePeriodSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTimeTimePeriodSetupsById

> TimePeriodSetup PutTimeTimePeriodSetupsById(ctx, id).ClientId(clientId).TimePeriodSetup(timePeriodSetup).Execute()

Put TimePeriodSetup

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ai-connor/connectwise"
)

func main() {
	id := int32(56) // int32 | timePeriodSetupId
	clientId := "clientId_example" // string | 
	timePeriodSetup := *openapiclient.NewTimePeriodSetup("PeriodApplyTo_example", NullableInt32(123), NullableInt32(123), "Type_example", time.Now(), NullableInt32(123)) // TimePeriodSetup | timePeriodSetup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimePeriodSetupsAPI.PutTimeTimePeriodSetupsById(context.Background(), id).ClientId(clientId).TimePeriodSetup(timePeriodSetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimePeriodSetupsAPI.PutTimeTimePeriodSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTimeTimePeriodSetupsById`: TimePeriodSetup
	fmt.Fprintf(os.Stdout, "Response from `TimePeriodSetupsAPI.PutTimeTimePeriodSetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | timePeriodSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTimeTimePeriodSetupsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **timePeriodSetup** | [**TimePeriodSetup**](TimePeriodSetup.md) | timePeriodSetup | 

### Return type

[**TimePeriodSetup**](TimePeriodSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

