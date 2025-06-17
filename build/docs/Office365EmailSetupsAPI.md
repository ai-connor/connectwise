# \Office365EmailSetupsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemOffice365EmailSetupsById**](Office365EmailSetupsAPI.md#DeleteSystemOffice365EmailSetupsById) | **Delete** /system/office365/emailSetups/{id} | Delete Office365EmailSetup
[**GetSystemOffice365EmailSetups**](Office365EmailSetupsAPI.md#GetSystemOffice365EmailSetups) | **Get** /system/office365/emailSetups | Get List of Office365EmailSetup
[**GetSystemOffice365EmailSetupsById**](Office365EmailSetupsAPI.md#GetSystemOffice365EmailSetupsById) | **Get** /system/office365/emailSetups/{id} | Get Office365EmailSetup
[**GetSystemOffice365EmailSetupsByIdGetEmails**](Office365EmailSetupsAPI.md#GetSystemOffice365EmailSetupsByIdGetEmails) | **Get** /system/office365/emailSetups/{id}/getEmails/ | Get List of UserEmails from inbound ticket service
[**GetSystemOffice365EmailSetupsCount**](Office365EmailSetupsAPI.md#GetSystemOffice365EmailSetupsCount) | **Get** /system/office365/emailSetups/count | Get Count of Office365EmailSetup
[**PatchSystemOffice365EmailSetupsById**](Office365EmailSetupsAPI.md#PatchSystemOffice365EmailSetupsById) | **Patch** /system/office365/emailSetups/{id} | Patch Office365EmailSetup
[**PostSystemOffice365EmailSetups**](Office365EmailSetupsAPI.md#PostSystemOffice365EmailSetups) | **Post** /system/office365/emailSetups | Post Office365EmailSetup
[**PostSystemOffice365EmailSetupsByIdAuthorize**](Office365EmailSetupsAPI.md#PostSystemOffice365EmailSetupsByIdAuthorize) | **Post** /system/office365/emailSetups/{id}/authorize | Post SuccessResponse
[**PostSystemOffice365EmailSetupsByIdTestConnection**](Office365EmailSetupsAPI.md#PostSystemOffice365EmailSetupsByIdTestConnection) | **Post** /system/office365/emailSetups/{id}/testConnection | Post SuccessResponse
[**PutSystemOffice365EmailSetupsById**](Office365EmailSetupsAPI.md#PutSystemOffice365EmailSetupsById) | **Put** /system/office365/emailSetups/{id} | Put Office365EmailSetup



## DeleteSystemOffice365EmailSetupsById

> DeleteSystemOffice365EmailSetupsById(ctx, id).ClientId(clientId).Execute()

Delete Office365EmailSetup

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
	id := int32(56) // int32 | emailSetupId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.Office365EmailSetupsAPI.DeleteSystemOffice365EmailSetupsById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Office365EmailSetupsAPI.DeleteSystemOffice365EmailSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemOffice365EmailSetupsByIdRequest struct via the builder pattern


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


## GetSystemOffice365EmailSetups

> []Office365EmailSetup GetSystemOffice365EmailSetups(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of Office365EmailSetup

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
	resp, r, err := apiClient.Office365EmailSetupsAPI.GetSystemOffice365EmailSetups(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Office365EmailSetupsAPI.GetSystemOffice365EmailSetups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemOffice365EmailSetups`: []Office365EmailSetup
	fmt.Fprintf(os.Stdout, "Response from `Office365EmailSetupsAPI.GetSystemOffice365EmailSetups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemOffice365EmailSetupsRequest struct via the builder pattern


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

[**[]Office365EmailSetup**](Office365EmailSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemOffice365EmailSetupsById

> Office365EmailSetup GetSystemOffice365EmailSetupsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Office365EmailSetup

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
	id := int32(56) // int32 | emailSetupId
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
	resp, r, err := apiClient.Office365EmailSetupsAPI.GetSystemOffice365EmailSetupsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Office365EmailSetupsAPI.GetSystemOffice365EmailSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemOffice365EmailSetupsById`: Office365EmailSetup
	fmt.Fprintf(os.Stdout, "Response from `Office365EmailSetupsAPI.GetSystemOffice365EmailSetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemOffice365EmailSetupsByIdRequest struct via the builder pattern


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

[**Office365EmailSetup**](Office365EmailSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemOffice365EmailSetupsByIdGetEmails

> []UserEmail GetSystemOffice365EmailSetupsByIdGetEmails(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of UserEmails from inbound ticket service

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
	id := int32(56) // int32 | emailSetupId
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
	resp, r, err := apiClient.Office365EmailSetupsAPI.GetSystemOffice365EmailSetupsByIdGetEmails(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Office365EmailSetupsAPI.GetSystemOffice365EmailSetupsByIdGetEmails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemOffice365EmailSetupsByIdGetEmails`: []UserEmail
	fmt.Fprintf(os.Stdout, "Response from `Office365EmailSetupsAPI.GetSystemOffice365EmailSetupsByIdGetEmails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemOffice365EmailSetupsByIdGetEmailsRequest struct via the builder pattern


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

[**[]UserEmail**](UserEmail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemOffice365EmailSetupsCount

> Count GetSystemOffice365EmailSetupsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of Office365EmailSetup

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
	resp, r, err := apiClient.Office365EmailSetupsAPI.GetSystemOffice365EmailSetupsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Office365EmailSetupsAPI.GetSystemOffice365EmailSetupsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemOffice365EmailSetupsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `Office365EmailSetupsAPI.GetSystemOffice365EmailSetupsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemOffice365EmailSetupsCountRequest struct via the builder pattern


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


## PatchSystemOffice365EmailSetupsById

> Office365EmailSetup PatchSystemOffice365EmailSetupsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch Office365EmailSetup

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
	id := int32(56) // int32 | emailSetupId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Office365EmailSetupsAPI.PatchSystemOffice365EmailSetupsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Office365EmailSetupsAPI.PatchSystemOffice365EmailSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemOffice365EmailSetupsById`: Office365EmailSetup
	fmt.Fprintf(os.Stdout, "Response from `Office365EmailSetupsAPI.PatchSystemOffice365EmailSetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemOffice365EmailSetupsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**Office365EmailSetup**](Office365EmailSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemOffice365EmailSetups

> Office365EmailSetup PostSystemOffice365EmailSetups(ctx).ClientId(clientId).Office365EmailSetup(office365EmailSetup).Execute()

Post Office365EmailSetup

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
	office365EmailSetup := *openapiclient.NewOffice365EmailSetup("Name_example", "InboxFolder_example", "ProcessedFolder_example", "FailedFolder_example") // Office365EmailSetup | entity

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Office365EmailSetupsAPI.PostSystemOffice365EmailSetups(context.Background()).ClientId(clientId).Office365EmailSetup(office365EmailSetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Office365EmailSetupsAPI.PostSystemOffice365EmailSetups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemOffice365EmailSetups`: Office365EmailSetup
	fmt.Fprintf(os.Stdout, "Response from `Office365EmailSetupsAPI.PostSystemOffice365EmailSetups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemOffice365EmailSetupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **office365EmailSetup** | [**Office365EmailSetup**](Office365EmailSetup.md) | entity | 

### Return type

[**Office365EmailSetup**](Office365EmailSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemOffice365EmailSetupsByIdAuthorize

> SuccessResponse PostSystemOffice365EmailSetupsByIdAuthorize(ctx, id).ClientId(clientId).Execute()

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
	id := int32(56) // int32 | emailSetupId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Office365EmailSetupsAPI.PostSystemOffice365EmailSetupsByIdAuthorize(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Office365EmailSetupsAPI.PostSystemOffice365EmailSetupsByIdAuthorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemOffice365EmailSetupsByIdAuthorize`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `Office365EmailSetupsAPI.PostSystemOffice365EmailSetupsByIdAuthorize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemOffice365EmailSetupsByIdAuthorizeRequest struct via the builder pattern


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


## PostSystemOffice365EmailSetupsByIdTestConnection

> SuccessResponse PostSystemOffice365EmailSetupsByIdTestConnection(ctx, id).ClientId(clientId).Execute()

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
	id := int32(56) // int32 | emailSetupId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Office365EmailSetupsAPI.PostSystemOffice365EmailSetupsByIdTestConnection(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Office365EmailSetupsAPI.PostSystemOffice365EmailSetupsByIdTestConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemOffice365EmailSetupsByIdTestConnection`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `Office365EmailSetupsAPI.PostSystemOffice365EmailSetupsByIdTestConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemOffice365EmailSetupsByIdTestConnectionRequest struct via the builder pattern


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


## PutSystemOffice365EmailSetupsById

> Office365EmailSetup PutSystemOffice365EmailSetupsById(ctx, id).ClientId(clientId).Office365EmailSetup(office365EmailSetup).Execute()

Put Office365EmailSetup

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
	id := int32(56) // int32 | emailSetupId
	clientId := "clientId_example" // string | 
	office365EmailSetup := *openapiclient.NewOffice365EmailSetup("Name_example", "InboxFolder_example", "ProcessedFolder_example", "FailedFolder_example") // Office365EmailSetup | entity

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Office365EmailSetupsAPI.PutSystemOffice365EmailSetupsById(context.Background(), id).ClientId(clientId).Office365EmailSetup(office365EmailSetup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Office365EmailSetupsAPI.PutSystemOffice365EmailSetupsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemOffice365EmailSetupsById`: Office365EmailSetup
	fmt.Fprintf(os.Stdout, "Response from `Office365EmailSetupsAPI.PutSystemOffice365EmailSetupsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | emailSetupId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemOffice365EmailSetupsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **office365EmailSetup** | [**Office365EmailSetup**](Office365EmailSetup.md) | entity | 

### Return type

[**Office365EmailSetup**](Office365EmailSetup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

