# \AgreementTypeWorkRoleExclusionsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceAgreementTypesByParentIdWorkRoleExclusionsById**](AgreementTypeWorkRoleExclusionsAPI.md#DeleteFinanceAgreementTypesByParentIdWorkRoleExclusionsById) | **Delete** /finance/agreementTypes/{parentId}/workRoleExclusions/{id} | Delete AgreementTypeWorkRoleExclusion
[**GetFinanceAgreementTypesByParentIdWorkRoleExclusions**](AgreementTypeWorkRoleExclusionsAPI.md#GetFinanceAgreementTypesByParentIdWorkRoleExclusions) | **Get** /finance/agreementTypes/{parentId}/workRoleExclusions | Get List of AgreementTypeWorkRoleExclusion
[**GetFinanceAgreementTypesByParentIdWorkRoleExclusionsById**](AgreementTypeWorkRoleExclusionsAPI.md#GetFinanceAgreementTypesByParentIdWorkRoleExclusionsById) | **Get** /finance/agreementTypes/{parentId}/workRoleExclusions/{id} | Get AgreementTypeWorkRoleExclusion
[**GetFinanceAgreementTypesByParentIdWorkRoleExclusionsCount**](AgreementTypeWorkRoleExclusionsAPI.md#GetFinanceAgreementTypesByParentIdWorkRoleExclusionsCount) | **Get** /finance/agreementTypes/{parentId}/workRoleExclusions/count | Get Count of AgreementTypeWorkRoleExclusion
[**PostFinanceAgreementTypesByParentIdWorkRoleExclusions**](AgreementTypeWorkRoleExclusionsAPI.md#PostFinanceAgreementTypesByParentIdWorkRoleExclusions) | **Post** /finance/agreementTypes/{parentId}/workRoleExclusions | Post AgreementTypeWorkRoleExclusion



## DeleteFinanceAgreementTypesByParentIdWorkRoleExclusionsById

> DeleteFinanceAgreementTypesByParentIdWorkRoleExclusionsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete AgreementTypeWorkRoleExclusion

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
	id := int32(56) // int32 | workRoleExclusionId
	parentId := int32(56) // int32 | agreementTypeId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgreementTypeWorkRoleExclusionsAPI.DeleteFinanceAgreementTypesByParentIdWorkRoleExclusionsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkRoleExclusionsAPI.DeleteFinanceAgreementTypesByParentIdWorkRoleExclusionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | workRoleExclusionId | 
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceAgreementTypesByParentIdWorkRoleExclusionsByIdRequest struct via the builder pattern


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


## GetFinanceAgreementTypesByParentIdWorkRoleExclusions

> []AgreementTypeWorkRoleExclusion GetFinanceAgreementTypesByParentIdWorkRoleExclusions(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of AgreementTypeWorkRoleExclusion

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
	resp, r, err := apiClient.AgreementTypeWorkRoleExclusionsAPI.GetFinanceAgreementTypesByParentIdWorkRoleExclusions(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkRoleExclusionsAPI.GetFinanceAgreementTypesByParentIdWorkRoleExclusions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementTypesByParentIdWorkRoleExclusions`: []AgreementTypeWorkRoleExclusion
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeWorkRoleExclusionsAPI.GetFinanceAgreementTypesByParentIdWorkRoleExclusions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementTypesByParentIdWorkRoleExclusionsRequest struct via the builder pattern


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

[**[]AgreementTypeWorkRoleExclusion**](AgreementTypeWorkRoleExclusion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAgreementTypesByParentIdWorkRoleExclusionsById

> AgreementTypeWorkRoleExclusion GetFinanceAgreementTypesByParentIdWorkRoleExclusionsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get AgreementTypeWorkRoleExclusion

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
	id := int32(56) // int32 | workRoleExclusionId
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
	resp, r, err := apiClient.AgreementTypeWorkRoleExclusionsAPI.GetFinanceAgreementTypesByParentIdWorkRoleExclusionsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkRoleExclusionsAPI.GetFinanceAgreementTypesByParentIdWorkRoleExclusionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementTypesByParentIdWorkRoleExclusionsById`: AgreementTypeWorkRoleExclusion
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeWorkRoleExclusionsAPI.GetFinanceAgreementTypesByParentIdWorkRoleExclusionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | workRoleExclusionId | 
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementTypesByParentIdWorkRoleExclusionsByIdRequest struct via the builder pattern


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

[**AgreementTypeWorkRoleExclusion**](AgreementTypeWorkRoleExclusion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceAgreementTypesByParentIdWorkRoleExclusionsCount

> Count GetFinanceAgreementTypesByParentIdWorkRoleExclusionsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of AgreementTypeWorkRoleExclusion

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
	resp, r, err := apiClient.AgreementTypeWorkRoleExclusionsAPI.GetFinanceAgreementTypesByParentIdWorkRoleExclusionsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkRoleExclusionsAPI.GetFinanceAgreementTypesByParentIdWorkRoleExclusionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceAgreementTypesByParentIdWorkRoleExclusionsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeWorkRoleExclusionsAPI.GetFinanceAgreementTypesByParentIdWorkRoleExclusionsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceAgreementTypesByParentIdWorkRoleExclusionsCountRequest struct via the builder pattern


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


## PostFinanceAgreementTypesByParentIdWorkRoleExclusions

> AgreementTypeWorkRoleExclusion PostFinanceAgreementTypesByParentIdWorkRoleExclusions(ctx, parentId).ClientId(clientId).AgreementTypeWorkRoleExclusion(agreementTypeWorkRoleExclusion).Execute()

Post AgreementTypeWorkRoleExclusion

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
	agreementTypeWorkRoleExclusion := *openapiclient.NewAgreementTypeWorkRoleExclusion(*openapiclient.NewWorkRoleReference()) // AgreementTypeWorkRoleExclusion | workRoleExclusion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgreementTypeWorkRoleExclusionsAPI.PostFinanceAgreementTypesByParentIdWorkRoleExclusions(context.Background(), parentId).ClientId(clientId).AgreementTypeWorkRoleExclusion(agreementTypeWorkRoleExclusion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgreementTypeWorkRoleExclusionsAPI.PostFinanceAgreementTypesByParentIdWorkRoleExclusions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceAgreementTypesByParentIdWorkRoleExclusions`: AgreementTypeWorkRoleExclusion
	fmt.Fprintf(os.Stdout, "Response from `AgreementTypeWorkRoleExclusionsAPI.PostFinanceAgreementTypesByParentIdWorkRoleExclusions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | agreementTypeId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceAgreementTypesByParentIdWorkRoleExclusionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **agreementTypeWorkRoleExclusion** | [**AgreementTypeWorkRoleExclusion**](AgreementTypeWorkRoleExclusion.md) | workRoleExclusion | 

### Return type

[**AgreementTypeWorkRoleExclusion**](AgreementTypeWorkRoleExclusion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

