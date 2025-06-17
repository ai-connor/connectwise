# \TrackActionsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyTracksByParentIdActionsById**](TrackActionsAPI.md#DeleteCompanyTracksByParentIdActionsById) | **Delete** /company/tracks/{parentId}/actions/{id} | Delete TrackAction
[**GetCompanyTracksByParentIdActions**](TrackActionsAPI.md#GetCompanyTracksByParentIdActions) | **Get** /company/tracks/{parentId}/actions | Get List of TrackAction
[**GetCompanyTracksByParentIdActionsById**](TrackActionsAPI.md#GetCompanyTracksByParentIdActionsById) | **Get** /company/tracks/{parentId}/actions/{id} | Get TrackAction
[**GetCompanyTracksByParentIdActionsCount**](TrackActionsAPI.md#GetCompanyTracksByParentIdActionsCount) | **Get** /company/tracks/{parentId}/actions/count | Get Count of TrackAction
[**PatchCompanyTracksByParentIdActionsById**](TrackActionsAPI.md#PatchCompanyTracksByParentIdActionsById) | **Patch** /company/tracks/{parentId}/actions/{id} | Patch TrackAction
[**PostCompanyTracksByParentIdActions**](TrackActionsAPI.md#PostCompanyTracksByParentIdActions) | **Post** /company/tracks/{parentId}/actions | Post TrackAction
[**PutCompanyTracksByParentIdActionsById**](TrackActionsAPI.md#PutCompanyTracksByParentIdActionsById) | **Put** /company/tracks/{parentId}/actions/{id} | Put TrackAction



## DeleteCompanyTracksByParentIdActionsById

> DeleteCompanyTracksByParentIdActionsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete TrackAction

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
	id := int32(56) // int32 | actionId
	parentId := int32(56) // int32 | trackId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TrackActionsAPI.DeleteCompanyTracksByParentIdActionsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackActionsAPI.DeleteCompanyTracksByParentIdActionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | actionId | 
**parentId** | **int32** | trackId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyTracksByParentIdActionsByIdRequest struct via the builder pattern


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


## GetCompanyTracksByParentIdActions

> []TrackAction GetCompanyTracksByParentIdActions(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of TrackAction

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
	parentId := int32(56) // int32 | trackId
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
	resp, r, err := apiClient.TrackActionsAPI.GetCompanyTracksByParentIdActions(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackActionsAPI.GetCompanyTracksByParentIdActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyTracksByParentIdActions`: []TrackAction
	fmt.Fprintf(os.Stdout, "Response from `TrackActionsAPI.GetCompanyTracksByParentIdActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | trackId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyTracksByParentIdActionsRequest struct via the builder pattern


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

[**[]TrackAction**](TrackAction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyTracksByParentIdActionsById

> TrackAction GetCompanyTracksByParentIdActionsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get TrackAction

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
	id := int32(56) // int32 | actionId
	parentId := int32(56) // int32 | trackId
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
	resp, r, err := apiClient.TrackActionsAPI.GetCompanyTracksByParentIdActionsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackActionsAPI.GetCompanyTracksByParentIdActionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyTracksByParentIdActionsById`: TrackAction
	fmt.Fprintf(os.Stdout, "Response from `TrackActionsAPI.GetCompanyTracksByParentIdActionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | actionId | 
**parentId** | **int32** | trackId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyTracksByParentIdActionsByIdRequest struct via the builder pattern


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

[**TrackAction**](TrackAction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyTracksByParentIdActionsCount

> Count GetCompanyTracksByParentIdActionsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of TrackAction

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
	parentId := int32(56) // int32 | trackId
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
	resp, r, err := apiClient.TrackActionsAPI.GetCompanyTracksByParentIdActionsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackActionsAPI.GetCompanyTracksByParentIdActionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyTracksByParentIdActionsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `TrackActionsAPI.GetCompanyTracksByParentIdActionsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | trackId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyTracksByParentIdActionsCountRequest struct via the builder pattern


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


## PatchCompanyTracksByParentIdActionsById

> TrackAction PatchCompanyTracksByParentIdActionsById(ctx, id, parentId).ClientId(clientId).Execute()

Patch TrackAction

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
	id := int32(56) // int32 | actionId
	parentId := int32(56) // int32 | trackId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrackActionsAPI.PatchCompanyTracksByParentIdActionsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackActionsAPI.PatchCompanyTracksByParentIdActionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyTracksByParentIdActionsById`: TrackAction
	fmt.Fprintf(os.Stdout, "Response from `TrackActionsAPI.PatchCompanyTracksByParentIdActionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | actionId | 
**parentId** | **int32** | trackId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyTracksByParentIdActionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 

### Return type

[**TrackAction**](TrackAction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyTracksByParentIdActions

> TrackAction PostCompanyTracksByParentIdActions(ctx, parentId).ClientId(clientId).TrackAction(trackAction).Execute()

Post TrackAction

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
	parentId := int32(56) // int32 | trackId
	clientId := "clientId_example" // string | 
	trackAction := *openapiclient.NewTrackAction("NotifyType_example") // TrackAction | trackAction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrackActionsAPI.PostCompanyTracksByParentIdActions(context.Background(), parentId).ClientId(clientId).TrackAction(trackAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackActionsAPI.PostCompanyTracksByParentIdActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyTracksByParentIdActions`: TrackAction
	fmt.Fprintf(os.Stdout, "Response from `TrackActionsAPI.PostCompanyTracksByParentIdActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | trackId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyTracksByParentIdActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **trackAction** | [**TrackAction**](TrackAction.md) | trackAction | 

### Return type

[**TrackAction**](TrackAction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyTracksByParentIdActionsById

> TrackAction PutCompanyTracksByParentIdActionsById(ctx, id, parentId).ClientId(clientId).TrackAction(trackAction).Execute()

Put TrackAction

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
	id := int32(56) // int32 | actionId
	parentId := int32(56) // int32 | trackId
	clientId := "clientId_example" // string | 
	trackAction := *openapiclient.NewTrackAction("NotifyType_example") // TrackAction | trackAction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrackActionsAPI.PutCompanyTracksByParentIdActionsById(context.Background(), id, parentId).ClientId(clientId).TrackAction(trackAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackActionsAPI.PutCompanyTracksByParentIdActionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyTracksByParentIdActionsById`: TrackAction
	fmt.Fprintf(os.Stdout, "Response from `TrackActionsAPI.PutCompanyTracksByParentIdActionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | actionId | 
**parentId** | **int32** | trackId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyTracksByParentIdActionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **trackAction** | [**TrackAction**](TrackAction.md) | trackAction | 

### Return type

[**TrackAction**](TrackAction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

