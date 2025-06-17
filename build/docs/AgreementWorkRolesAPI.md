# \AgreementWorkRolesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceAgreementsByParentIdWorkrolesById**](AgreementWorkRolesAPI.md#DeleteFinanceAgreementsByParentIdWorkrolesById) | **Delete** /finance/agreements/{parentId}/workroles/{id} | Delete AgreementWorkRole
[**GetFinanceAgreementsByParentIdWorkroles**](AgreementWorkRolesAPI.md#GetFinanceAgreementsByParentIdWorkroles) | **Get** /finance/agreements/{parentId}/workroles | Get List of AgreementWorkRole
[**GetFinanceAgreementsByParentIdWorkrolesById**](AgreementWorkRolesAPI.md#GetFinanceAgreementsByParentIdWorkrolesById) | **Get** /finance/agreements/{parentId}/workroles/{id} | Get AgreementWorkRole
[**GetFinanceAgreementsByParentIdWorkrolesCount**](AgreementWorkRolesAPI.md#GetFinanceAgreementsByParentIdWorkrolesCount) | **Get** /finance/agreements/{parentId}/workroles/count | Get Count of AgreementWorkRole
[**PatchFinanceAgreementsByParentIdWorkrolesById**](AgreementWorkRolesAPI.md#PatchFinanceAgreementsByParentIdWorkrolesById) | **Patch** /finance/agreements/{parentId}/workroles/{id} | Patch AgreementWorkRole
[**PostFinanceAgreementsByParentIdWorkroles**](AgreementWorkRolesAPI.md#PostFinanceAgreementsByParentIdWorkroles) | **Post** /finance/agreements/{parentId}/workroles | Post AgreementWorkRole
[**PutFinanceAgreementsByParentIdWorkrolesById**](AgreementWorkRolesAPI.md#PutFinanceAgreementsByParentIdWorkrolesById) | **Put** /finance/agreements/{parentId}/workroles/{id} | Put AgreementWorkRole



## DeleteFinanceAgreementsByParentIdWorkrolesById

> DeleteFinanceAgreementsByParentIdWorkrolesById(ctx, id, parentId).ClientId(clientId).Execute()

Delete AgreementWorkRole

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
	id := int32(56) // int32 | workroleId
	parentId := int32(56) // int32 | agreementId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgreementWorkRolesAPI.DeleteFinanceAgreementsByParentIdWorkrolesById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementWorkRolesAPI.DeleteFinanceAgreementsByParentIdWorkrolesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | workroleId | 
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceAgreementsByParentIdWorkrolesByIdRequest struct via the builder pattern


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


## GetFinanceAgreementsByParentIdWorkroles

> []AgreementWorkRole GetFinanceAgreementsByParentIdWorkroles(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of AgreementWorkRole

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
	parentId := int32(56) // int32 | agreementId
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
	resp, r, err := apiClient.AgreementWorkRolesAPI.GetFinanceAgreementsByParentIdWorkroles(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementWorkRolesAPI.GetFinanceAgreementsByParentIdWorkroles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementsByParentIdWorkroles`: []AgreementWorkRole
	fmt.Fprintf(os.Stdout, "Response from `AgreementWorkRolesAPI.GetFinanceAgreementsByParentIdWorkroles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementsByParentIdWorkrolesRequest struct via the builder pattern


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

[**[]AgreementWorkRole**](AgreementWorkRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAgreementsByParentIdWorkrolesById

> AgreementWorkRole GetFinanceAgreementsByParentIdWorkrolesById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get AgreementWorkRole

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
	id := int32(56) // int32 | workroleId
	parentId := int32(56) // int32 | agreementId
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
	resp, r, err := apiClient.AgreementWorkRolesAPI.GetFinanceAgreementsByParentIdWorkrolesById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementWorkRolesAPI.GetFinanceAgreementsByParentIdWorkrolesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementsByParentIdWorkrolesById`: AgreementWorkRole
	fmt.Fprintf(os.Stdout, "Response from `AgreementWorkRolesAPI.GetFinanceAgreementsByParentIdWorkrolesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | workroleId | 
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementsByParentIdWorkrolesByIdRequest struct via the builder pattern


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

[**AgreementWorkRole**](AgreementWorkRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAgreementsByParentIdWorkrolesCount

> Count GetFinanceAgreementsByParentIdWorkrolesCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of AgreementWorkRole

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
	parentId := int32(56) // int32 | agreementId
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
	resp, r, err := apiClient.AgreementWorkRolesAPI.GetFinanceAgreementsByParentIdWorkrolesCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementWorkRolesAPI.GetFinanceAgreementsByParentIdWorkrolesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementsByParentIdWorkrolesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `AgreementWorkRolesAPI.GetFinanceAgreementsByParentIdWorkrolesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementsByParentIdWorkrolesCountRequest struct via the builder pattern


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


## PatchFinanceAgreementsByParentIdWorkrolesById

> AgreementWorkRole PatchFinanceAgreementsByParentIdWorkrolesById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch AgreementWorkRole

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
	id := int32(56) // int32 | workroleId
	parentId := int32(56) // int32 | agreementId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgreementWorkRolesAPI.PatchFinanceAgreementsByParentIdWorkrolesById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementWorkRolesAPI.PatchFinanceAgreementsByParentIdWorkrolesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceAgreementsByParentIdWorkrolesById`: AgreementWorkRole
	fmt.Fprintf(os.Stdout, "Response from `AgreementWorkRolesAPI.PatchFinanceAgreementsByParentIdWorkrolesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | workroleId | 
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceAgreementsByParentIdWorkrolesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**AgreementWorkRole**](AgreementWorkRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceAgreementsByParentIdWorkroles

> AgreementWorkRole PostFinanceAgreementsByParentIdWorkroles(ctx, parentId).ClientId(clientId).AgreementWorkRole(agreementWorkRole).Execute()

Post AgreementWorkRole

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
	parentId := int32(56) // int32 | agreementId
	clientId := "clientId_example" // string | 
	agreementWorkRole := *openapiclient.NewAgreementWorkRole("RateType_example") // AgreementWorkRole | workRole

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgreementWorkRolesAPI.PostFinanceAgreementsByParentIdWorkroles(context.Background(), parentId).ClientId(clientId).AgreementWorkRole(agreementWorkRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementWorkRolesAPI.PostFinanceAgreementsByParentIdWorkroles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceAgreementsByParentIdWorkroles`: AgreementWorkRole
	fmt.Fprintf(os.Stdout, "Response from `AgreementWorkRolesAPI.PostFinanceAgreementsByParentIdWorkroles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceAgreementsByParentIdWorkrolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **agreementWorkRole** | [**AgreementWorkRole**](AgreementWorkRole.md) | workRole | 

### Return type

[**AgreementWorkRole**](AgreementWorkRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceAgreementsByParentIdWorkrolesById

> AgreementWorkRole PutFinanceAgreementsByParentIdWorkrolesById(ctx, id, parentId).ClientId(clientId).AgreementWorkRole(agreementWorkRole).Execute()

Put AgreementWorkRole

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
	id := int32(56) // int32 | workroleId
	parentId := int32(56) // int32 | agreementId
	clientId := "clientId_example" // string | 
	agreementWorkRole := *openapiclient.NewAgreementWorkRole("RateType_example") // AgreementWorkRole | workRole

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgreementWorkRolesAPI.PutFinanceAgreementsByParentIdWorkrolesById(context.Background(), id, parentId).ClientId(clientId).AgreementWorkRole(agreementWorkRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementWorkRolesAPI.PutFinanceAgreementsByParentIdWorkrolesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceAgreementsByParentIdWorkrolesById`: AgreementWorkRole
	fmt.Fprintf(os.Stdout, "Response from `AgreementWorkRolesAPI.PutFinanceAgreementsByParentIdWorkrolesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | workroleId | 
**parentId** | **int32** | agreementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceAgreementsByParentIdWorkrolesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **agreementWorkRole** | [**AgreementWorkRole**](AgreementWorkRole.md) | workRole | 

### Return type

[**AgreementWorkRole**](AgreementWorkRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

