# \AgreementTypeWorkRolesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceAgreementTypesByParentIdWorkrolesById**](AgreementTypeWorkRolesAPI.md#DeleteFinanceAgreementTypesByParentIdWorkrolesById) | **Delete** /finance/agreementTypes/{parentId}/workroles/{id} | Delete AgreementTypeWorkRole
[**GetFinanceAgreementTypesByParentIdWorkroles**](AgreementTypeWorkRolesAPI.md#GetFinanceAgreementTypesByParentIdWorkroles) | **Get** /finance/agreementTypes/{parentId}/workroles | Get List of AgreementTypeWorkRole
[**GetFinanceAgreementTypesByParentIdWorkrolesById**](AgreementTypeWorkRolesAPI.md#GetFinanceAgreementTypesByParentIdWorkrolesById) | **Get** /finance/agreementTypes/{parentId}/workroles/{id} | Get AgreementTypeWorkRole
[**GetFinanceAgreementTypesByParentIdWorkrolesCount**](AgreementTypeWorkRolesAPI.md#GetFinanceAgreementTypesByParentIdWorkrolesCount) | **Get** /finance/agreementTypes/{parentId}/workroles/count | Get Count of AgreementTypeWorkRole
[**GetFinanceAgreementTypesByParentIdWorkrolesInfo**](AgreementTypeWorkRolesAPI.md#GetFinanceAgreementTypesByParentIdWorkrolesInfo) | **Get** /finance/agreementTypes/{parentId}/workroles/info | Get List of AgreementTypeWorkRole
[**GetFinanceAgreementTypesByParentIdWorkrolesInfoById**](AgreementTypeWorkRolesAPI.md#GetFinanceAgreementTypesByParentIdWorkrolesInfoById) | **Get** /finance/agreementTypes/{parentId}/workroles/info/{id} | Get AgreementTypeWorkRoleInfo
[**GetFinanceAgreementTypesByParentIdWorkrolesInfoCount**](AgreementTypeWorkRolesAPI.md#GetFinanceAgreementTypesByParentIdWorkrolesInfoCount) | **Get** /finance/agreementTypes/{parentId}/workroles/info/count | Get Count of AgreementTypeWorkRoleInfo
[**PatchFinanceAgreementTypesByParentIdWorkrolesById**](AgreementTypeWorkRolesAPI.md#PatchFinanceAgreementTypesByParentIdWorkrolesById) | **Patch** /finance/agreementTypes/{parentId}/workroles/{id} | Patch AgreementTypeWorkRole
[**PostFinanceAgreementTypesByParentIdWorkroles**](AgreementTypeWorkRolesAPI.md#PostFinanceAgreementTypesByParentIdWorkroles) | **Post** /finance/agreementTypes/{parentId}/workroles | Post AgreementTypeWorkRole
[**PutFinanceAgreementTypesByParentIdWorkrolesById**](AgreementTypeWorkRolesAPI.md#PutFinanceAgreementTypesByParentIdWorkrolesById) | **Put** /finance/agreementTypes/{parentId}/workroles/{id} | Put AgreementTypeWorkRole



## DeleteFinanceAgreementTypesByParentIdWorkrolesById

> DeleteFinanceAgreementTypesByParentIdWorkrolesById(ctx, id, parentId).ClientId(clientId).Execute()

Delete AgreementTypeWorkRole

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
	parentId := int32(56) // int32 | agreementTypeId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgreementTypeWorkRolesAPI.DeleteFinanceAgreementTypesByParentIdWorkrolesById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkRolesAPI.DeleteFinanceAgreementTypesByParentIdWorkrolesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | workroleId | 
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceAgreementTypesByParentIdWorkrolesByIdRequest struct via the builder pattern


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


## GetFinanceAgreementTypesByParentIdWorkroles

> []AgreementTypeWorkRole GetFinanceAgreementTypesByParentIdWorkroles(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of AgreementTypeWorkRole

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
	parentId := int32(56) // int32 | agreementTypeId
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
	resp, r, err := apiClient.AgreementTypeWorkRolesAPI.GetFinanceAgreementTypesByParentIdWorkroles(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkRolesAPI.GetFinanceAgreementTypesByParentIdWorkroles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementTypesByParentIdWorkroles`: []AgreementTypeWorkRole
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeWorkRolesAPI.GetFinanceAgreementTypesByParentIdWorkroles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementTypesByParentIdWorkrolesRequest struct via the builder pattern


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

[**[]AgreementTypeWorkRole**](AgreementTypeWorkRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAgreementTypesByParentIdWorkrolesById

> AgreementTypeWorkRole GetFinanceAgreementTypesByParentIdWorkrolesById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get AgreementTypeWorkRole

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
	parentId := int32(56) // int32 | agreementTypeId
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
	resp, r, err := apiClient.AgreementTypeWorkRolesAPI.GetFinanceAgreementTypesByParentIdWorkrolesById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkRolesAPI.GetFinanceAgreementTypesByParentIdWorkrolesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementTypesByParentIdWorkrolesById`: AgreementTypeWorkRole
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeWorkRolesAPI.GetFinanceAgreementTypesByParentIdWorkrolesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | workroleId | 
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementTypesByParentIdWorkrolesByIdRequest struct via the builder pattern


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

[**AgreementTypeWorkRole**](AgreementTypeWorkRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAgreementTypesByParentIdWorkrolesCount

> Count GetFinanceAgreementTypesByParentIdWorkrolesCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of AgreementTypeWorkRole

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
	parentId := int32(56) // int32 | agreementTypeId
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
	resp, r, err := apiClient.AgreementTypeWorkRolesAPI.GetFinanceAgreementTypesByParentIdWorkrolesCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkRolesAPI.GetFinanceAgreementTypesByParentIdWorkrolesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementTypesByParentIdWorkrolesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeWorkRolesAPI.GetFinanceAgreementTypesByParentIdWorkrolesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementTypesByParentIdWorkrolesCountRequest struct via the builder pattern


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


## GetFinanceAgreementTypesByParentIdWorkrolesInfo

> []AgreementTypeWorkRoleInfo GetFinanceAgreementTypesByParentIdWorkrolesInfo(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of AgreementTypeWorkRole

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
	parentId := int32(56) // int32 | agreementTypeId
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
	resp, r, err := apiClient.AgreementTypeWorkRolesAPI.GetFinanceAgreementTypesByParentIdWorkrolesInfo(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkRolesAPI.GetFinanceAgreementTypesByParentIdWorkrolesInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementTypesByParentIdWorkrolesInfo`: []AgreementTypeWorkRoleInfo
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeWorkRolesAPI.GetFinanceAgreementTypesByParentIdWorkrolesInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementTypesByParentIdWorkrolesInfoRequest struct via the builder pattern


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

[**[]AgreementTypeWorkRoleInfo**](AgreementTypeWorkRoleInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAgreementTypesByParentIdWorkrolesInfoById

> AgreementTypeWorkRoleInfo GetFinanceAgreementTypesByParentIdWorkrolesInfoById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get AgreementTypeWorkRoleInfo

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
	parentId := int32(56) // int32 | agreementTypeId
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
	resp, r, err := apiClient.AgreementTypeWorkRolesAPI.GetFinanceAgreementTypesByParentIdWorkrolesInfoById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkRolesAPI.GetFinanceAgreementTypesByParentIdWorkrolesInfoById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementTypesByParentIdWorkrolesInfoById`: AgreementTypeWorkRoleInfo
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeWorkRolesAPI.GetFinanceAgreementTypesByParentIdWorkrolesInfoById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | workroleId | 
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementTypesByParentIdWorkrolesInfoByIdRequest struct via the builder pattern


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

[**AgreementTypeWorkRoleInfo**](AgreementTypeWorkRoleInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAgreementTypesByParentIdWorkrolesInfoCount

> Count GetFinanceAgreementTypesByParentIdWorkrolesInfoCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of AgreementTypeWorkRoleInfo

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
	parentId := int32(56) // int32 | agreementTypeId
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
	resp, r, err := apiClient.AgreementTypeWorkRolesAPI.GetFinanceAgreementTypesByParentIdWorkrolesInfoCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkRolesAPI.GetFinanceAgreementTypesByParentIdWorkrolesInfoCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementTypesByParentIdWorkrolesInfoCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeWorkRolesAPI.GetFinanceAgreementTypesByParentIdWorkrolesInfoCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementTypesByParentIdWorkrolesInfoCountRequest struct via the builder pattern


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


## PatchFinanceAgreementTypesByParentIdWorkrolesById

> AgreementTypeWorkRole PatchFinanceAgreementTypesByParentIdWorkrolesById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch AgreementTypeWorkRole

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
	parentId := int32(56) // int32 | agreementTypeId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgreementTypeWorkRolesAPI.PatchFinanceAgreementTypesByParentIdWorkrolesById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkRolesAPI.PatchFinanceAgreementTypesByParentIdWorkrolesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceAgreementTypesByParentIdWorkrolesById`: AgreementTypeWorkRole
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeWorkRolesAPI.PatchFinanceAgreementTypesByParentIdWorkrolesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | workroleId | 
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceAgreementTypesByParentIdWorkrolesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**AgreementTypeWorkRole**](AgreementTypeWorkRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceAgreementTypesByParentIdWorkroles

> AgreementTypeWorkRole PostFinanceAgreementTypesByParentIdWorkroles(ctx, parentId).ClientId(clientId).AgreementTypeWorkRole(agreementTypeWorkRole).Execute()

Post AgreementTypeWorkRole

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
	parentId := int32(56) // int32 | agreementTypeId
	clientId := "clientId_example" // string | 
	agreementTypeWorkRole := *openapiclient.NewAgreementTypeWorkRole("RateType_example") // AgreementTypeWorkRole | workRole

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgreementTypeWorkRolesAPI.PostFinanceAgreementTypesByParentIdWorkroles(context.Background(), parentId).ClientId(clientId).AgreementTypeWorkRole(agreementTypeWorkRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkRolesAPI.PostFinanceAgreementTypesByParentIdWorkroles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceAgreementTypesByParentIdWorkroles`: AgreementTypeWorkRole
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeWorkRolesAPI.PostFinanceAgreementTypesByParentIdWorkroles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceAgreementTypesByParentIdWorkrolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **agreementTypeWorkRole** | [**AgreementTypeWorkRole**](AgreementTypeWorkRole.md) | workRole | 

### Return type

[**AgreementTypeWorkRole**](AgreementTypeWorkRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceAgreementTypesByParentIdWorkrolesById

> AgreementTypeWorkRole PutFinanceAgreementTypesByParentIdWorkrolesById(ctx, id, parentId).ClientId(clientId).AgreementTypeWorkRole(agreementTypeWorkRole).Execute()

Put AgreementTypeWorkRole

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
	parentId := int32(56) // int32 | agreementTypeId
	clientId := "clientId_example" // string | 
	agreementTypeWorkRole := *openapiclient.NewAgreementTypeWorkRole("RateType_example") // AgreementTypeWorkRole | workRole

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgreementTypeWorkRolesAPI.PutFinanceAgreementTypesByParentIdWorkrolesById(context.Background(), id, parentId).ClientId(clientId).AgreementTypeWorkRole(agreementTypeWorkRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkRolesAPI.PutFinanceAgreementTypesByParentIdWorkrolesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceAgreementTypesByParentIdWorkrolesById`: AgreementTypeWorkRole
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeWorkRolesAPI.PutFinanceAgreementTypesByParentIdWorkrolesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | workroleId | 
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceAgreementTypesByParentIdWorkrolesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **agreementTypeWorkRole** | [**AgreementTypeWorkRole**](AgreementTypeWorkRole.md) | workRole | 

### Return type

[**AgreementTypeWorkRole**](AgreementTypeWorkRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

