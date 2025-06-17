# \MemberCertificationsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemMembersByParentIdCertificationsById**](MemberCertificationsAPI.md#DeleteSystemMembersByParentIdCertificationsById) | **Delete** /system/members/{parentId}/certifications/{id} | Delete MemberCertification
[**GetSystemMembersByParentIdCertifications**](MemberCertificationsAPI.md#GetSystemMembersByParentIdCertifications) | **Get** /system/members/{parentId}/certifications | Get List of MemberCertification
[**GetSystemMembersByParentIdCertificationsById**](MemberCertificationsAPI.md#GetSystemMembersByParentIdCertificationsById) | **Get** /system/members/{parentId}/certifications/{id} | Get MemberCertification
[**GetSystemMembersByParentIdCertificationsCount**](MemberCertificationsAPI.md#GetSystemMembersByParentIdCertificationsCount) | **Get** /system/members/{parentId}/certifications/count | Get Count of MemberCertification
[**PatchSystemMembersByParentIdCertificationsById**](MemberCertificationsAPI.md#PatchSystemMembersByParentIdCertificationsById) | **Patch** /system/members/{parentId}/certifications/{id} | Patch MemberCertification
[**PostSystemMembersByParentIdCertifications**](MemberCertificationsAPI.md#PostSystemMembersByParentIdCertifications) | **Post** /system/members/{parentId}/certifications | Post MemberCertification
[**PutSystemMembersByParentIdCertificationsById**](MemberCertificationsAPI.md#PutSystemMembersByParentIdCertificationsById) | **Put** /system/members/{parentId}/certifications/{id} | Put MemberCertification



## DeleteSystemMembersByParentIdCertificationsById

> DeleteSystemMembersByParentIdCertificationsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete MemberCertification

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
	id := int32(56) // int32 | certificationId
	parentId := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MemberCertificationsAPI.DeleteSystemMembersByParentIdCertificationsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberCertificationsAPI.DeleteSystemMembersByParentIdCertificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | certificationId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemMembersByParentIdCertificationsByIdRequest struct via the builder pattern


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


## GetSystemMembersByParentIdCertifications

> []MemberCertification GetSystemMembersByParentIdCertifications(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of MemberCertification

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
	parentId := int32(56) // int32 | memberId
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
	resp, r, err := apiClient.MemberCertificationsAPI.GetSystemMembersByParentIdCertifications(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberCertificationsAPI.GetSystemMembersByParentIdCertifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMembersByParentIdCertifications`: []MemberCertification
	fmt.Fprintf(os.Stdout, "Response from `MemberCertificationsAPI.GetSystemMembersByParentIdCertifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMembersByParentIdCertificationsRequest struct via the builder pattern


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

[**[]MemberCertification**](MemberCertification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemMembersByParentIdCertificationsById

> MemberCertification GetSystemMembersByParentIdCertificationsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get MemberCertification

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
	id := int32(56) // int32 | certificationId
	parentId := int32(56) // int32 | memberId
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
	resp, r, err := apiClient.MemberCertificationsAPI.GetSystemMembersByParentIdCertificationsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberCertificationsAPI.GetSystemMembersByParentIdCertificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMembersByParentIdCertificationsById`: MemberCertification
	fmt.Fprintf(os.Stdout, "Response from `MemberCertificationsAPI.GetSystemMembersByParentIdCertificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | certificationId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMembersByParentIdCertificationsByIdRequest struct via the builder pattern


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

[**MemberCertification**](MemberCertification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemMembersByParentIdCertificationsCount

> Count GetSystemMembersByParentIdCertificationsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of MemberCertification

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
	parentId := int32(56) // int32 | memberId
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
	resp, r, err := apiClient.MemberCertificationsAPI.GetSystemMembersByParentIdCertificationsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberCertificationsAPI.GetSystemMembersByParentIdCertificationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMembersByParentIdCertificationsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `MemberCertificationsAPI.GetSystemMembersByParentIdCertificationsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMembersByParentIdCertificationsCountRequest struct via the builder pattern


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


## PatchSystemMembersByParentIdCertificationsById

> MemberCertification PatchSystemMembersByParentIdCertificationsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch MemberCertification

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
	id := int32(56) // int32 | certificationId
	parentId := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberCertificationsAPI.PatchSystemMembersByParentIdCertificationsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberCertificationsAPI.PatchSystemMembersByParentIdCertificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemMembersByParentIdCertificationsById`: MemberCertification
	fmt.Fprintf(os.Stdout, "Response from `MemberCertificationsAPI.PatchSystemMembersByParentIdCertificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | certificationId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemMembersByParentIdCertificationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**MemberCertification**](MemberCertification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemMembersByParentIdCertifications

> MemberCertification PostSystemMembersByParentIdCertifications(ctx, parentId).ClientId(clientId).MemberCertification(memberCertification).Execute()

Post MemberCertification

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
	parentId := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 
	memberCertification := *openapiclient.NewMemberCertification(*openapiclient.NewCertificationReference()) // MemberCertification | memberCertification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberCertificationsAPI.PostSystemMembersByParentIdCertifications(context.Background(), parentId).ClientId(clientId).MemberCertification(memberCertification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberCertificationsAPI.PostSystemMembersByParentIdCertifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemMembersByParentIdCertifications`: MemberCertification
	fmt.Fprintf(os.Stdout, "Response from `MemberCertificationsAPI.PostSystemMembersByParentIdCertifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemMembersByParentIdCertificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **memberCertification** | [**MemberCertification**](MemberCertification.md) | memberCertification | 

### Return type

[**MemberCertification**](MemberCertification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemMembersByParentIdCertificationsById

> MemberCertification PutSystemMembersByParentIdCertificationsById(ctx, id, parentId).ClientId(clientId).MemberCertification(memberCertification).Execute()

Put MemberCertification

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
	id := int32(56) // int32 | certificationId
	parentId := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 
	memberCertification := *openapiclient.NewMemberCertification(*openapiclient.NewCertificationReference()) // MemberCertification | memberCertification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberCertificationsAPI.PutSystemMembersByParentIdCertificationsById(context.Background(), id, parentId).ClientId(clientId).MemberCertification(memberCertification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberCertificationsAPI.PutSystemMembersByParentIdCertificationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemMembersByParentIdCertificationsById`: MemberCertification
	fmt.Fprintf(os.Stdout, "Response from `MemberCertificationsAPI.PutSystemMembersByParentIdCertificationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | certificationId | 
**parentId** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemMembersByParentIdCertificationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **memberCertification** | [**MemberCertification**](MemberCertification.md) | memberCertification | 

### Return type

[**MemberCertification**](MemberCertification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

