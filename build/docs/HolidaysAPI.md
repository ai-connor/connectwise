# \HolidaysAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteScheduleHolidayListsByParentIdHolidaysById**](HolidaysAPI.md#DeleteScheduleHolidayListsByParentIdHolidaysById) | **Delete** /schedule/holidayLists/{parentId}/holidays/{id} | Delete Holiday
[**GetScheduleHolidayListsByParentIdHolidays**](HolidaysAPI.md#GetScheduleHolidayListsByParentIdHolidays) | **Get** /schedule/holidayLists/{parentId}/holidays | Get List of Holiday
[**GetScheduleHolidayListsByParentIdHolidaysById**](HolidaysAPI.md#GetScheduleHolidayListsByParentIdHolidaysById) | **Get** /schedule/holidayLists/{parentId}/holidays/{id} | Get Holiday
[**GetScheduleHolidayListsByParentIdHolidaysCount**](HolidaysAPI.md#GetScheduleHolidayListsByParentIdHolidaysCount) | **Get** /schedule/holidayLists/{parentId}/holidays/count | Get Count of Holiday
[**PatchScheduleHolidayListsByParentIdHolidaysById**](HolidaysAPI.md#PatchScheduleHolidayListsByParentIdHolidaysById) | **Patch** /schedule/holidayLists/{parentId}/holidays/{id} | Patch Holiday
[**PostScheduleHolidayListsByParentIdHolidays**](HolidaysAPI.md#PostScheduleHolidayListsByParentIdHolidays) | **Post** /schedule/holidayLists/{parentId}/holidays | Post Holiday
[**PutScheduleHolidayListsByParentIdHolidaysById**](HolidaysAPI.md#PutScheduleHolidayListsByParentIdHolidaysById) | **Put** /schedule/holidayLists/{parentId}/holidays/{id} | Put Holiday



## DeleteScheduleHolidayListsByParentIdHolidaysById

> DeleteScheduleHolidayListsByParentIdHolidaysById(ctx, id, parentId).ClientId(clientId).Execute()

Delete Holiday

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
	id := int32(56) // int32 | holidayId
	parentId := int32(56) // int32 | holidayListId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HolidaysAPI.DeleteScheduleHolidayListsByParentIdHolidaysById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.DeleteScheduleHolidayListsByParentIdHolidaysById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | holidayId | 
**parentId** | **int32** | holidayListId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScheduleHolidayListsByParentIdHolidaysByIdRequest struct via the builder pattern


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


## GetScheduleHolidayListsByParentIdHolidays

> []Holiday GetScheduleHolidayListsByParentIdHolidays(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of Holiday

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
	parentId := int32(56) // int32 | holidayListId
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
	resp, r, err := apiClient.HolidaysAPI.GetScheduleHolidayListsByParentIdHolidays(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.GetScheduleHolidayListsByParentIdHolidays``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScheduleHolidayListsByParentIdHolidays`: []Holiday
	fmt.Fprintf(os.Stdout, "Response from `HolidaysAPI.GetScheduleHolidayListsByParentIdHolidays`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | holidayListId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduleHolidayListsByParentIdHolidaysRequest struct via the builder pattern


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

[**[]Holiday**](Holiday.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScheduleHolidayListsByParentIdHolidaysById

> Holiday GetScheduleHolidayListsByParentIdHolidaysById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Holiday

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
	id := int32(56) // int32 | holidayId
	parentId := int32(56) // int32 | holidayListId
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
	resp, r, err := apiClient.HolidaysAPI.GetScheduleHolidayListsByParentIdHolidaysById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.GetScheduleHolidayListsByParentIdHolidaysById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScheduleHolidayListsByParentIdHolidaysById`: Holiday
	fmt.Fprintf(os.Stdout, "Response from `HolidaysAPI.GetScheduleHolidayListsByParentIdHolidaysById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | holidayId | 
**parentId** | **int32** | holidayListId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduleHolidayListsByParentIdHolidaysByIdRequest struct via the builder pattern


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

[**Holiday**](Holiday.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScheduleHolidayListsByParentIdHolidaysCount

> Count GetScheduleHolidayListsByParentIdHolidaysCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of Holiday

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
	parentId := int32(56) // int32 | holidayListId
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
	resp, r, err := apiClient.HolidaysAPI.GetScheduleHolidayListsByParentIdHolidaysCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.GetScheduleHolidayListsByParentIdHolidaysCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScheduleHolidayListsByParentIdHolidaysCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `HolidaysAPI.GetScheduleHolidayListsByParentIdHolidaysCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | holidayListId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduleHolidayListsByParentIdHolidaysCountRequest struct via the builder pattern


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


## PatchScheduleHolidayListsByParentIdHolidaysById

> Holiday PatchScheduleHolidayListsByParentIdHolidaysById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch Holiday

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
	id := int32(56) // int32 | holidayId
	parentId := int32(56) // int32 | holidayListId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HolidaysAPI.PatchScheduleHolidayListsByParentIdHolidaysById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.PatchScheduleHolidayListsByParentIdHolidaysById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchScheduleHolidayListsByParentIdHolidaysById`: Holiday
	fmt.Fprintf(os.Stdout, "Response from `HolidaysAPI.PatchScheduleHolidayListsByParentIdHolidaysById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | holidayId | 
**parentId** | **int32** | holidayListId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchScheduleHolidayListsByParentIdHolidaysByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**Holiday**](Holiday.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostScheduleHolidayListsByParentIdHolidays

> Holiday PostScheduleHolidayListsByParentIdHolidays(ctx, parentId).ClientId(clientId).Holiday(holiday).Execute()

Post Holiday

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
	parentId := int32(56) // int32 | holidayListId
	clientId := "clientId_example" // string | 
	holiday := *openapiclient.NewHoliday("Name_example", time.Now()) // Holiday | holiday

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HolidaysAPI.PostScheduleHolidayListsByParentIdHolidays(context.Background(), parentId).ClientId(clientId).Holiday(holiday).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.PostScheduleHolidayListsByParentIdHolidays``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostScheduleHolidayListsByParentIdHolidays`: Holiday
	fmt.Fprintf(os.Stdout, "Response from `HolidaysAPI.PostScheduleHolidayListsByParentIdHolidays`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | holidayListId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostScheduleHolidayListsByParentIdHolidaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **holiday** | [**Holiday**](Holiday.md) | holiday | 

### Return type

[**Holiday**](Holiday.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutScheduleHolidayListsByParentIdHolidaysById

> Holiday PutScheduleHolidayListsByParentIdHolidaysById(ctx, id, parentId).ClientId(clientId).Holiday(holiday).Execute()

Put Holiday

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
	id := int32(56) // int32 | holidayId
	parentId := int32(56) // int32 | holidayListId
	clientId := "clientId_example" // string | 
	holiday := *openapiclient.NewHoliday("Name_example", time.Now()) // Holiday | holiday

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HolidaysAPI.PutScheduleHolidayListsByParentIdHolidaysById(context.Background(), id, parentId).ClientId(clientId).Holiday(holiday).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HolidaysAPI.PutScheduleHolidayListsByParentIdHolidaysById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutScheduleHolidayListsByParentIdHolidaysById`: Holiday
	fmt.Fprintf(os.Stdout, "Response from `HolidaysAPI.PutScheduleHolidayListsByParentIdHolidaysById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | holidayId | 
**parentId** | **int32** | holidayListId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutScheduleHolidayListsByParentIdHolidaysByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **holiday** | [**Holiday**](Holiday.md) | holiday | 

### Return type

[**Holiday**](Holiday.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

