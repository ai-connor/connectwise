# \EmailConnectorParsingRulesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById**](EmailConnectorParsingRulesAPI.md#DeleteSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById) | **Delete** /system/emailConnectors/{grandparentId}/parsingStyles/{parentId}/parsingRules/{id} | Delete EmailConnectorParsingRule
[**GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRules**](EmailConnectorParsingRulesAPI.md#GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRules) | **Get** /system/emailConnectors/{grandparentId}/parsingStyles/{parentId}/parsingRules | Get List of EmailConnectorParsingRule
[**GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById**](EmailConnectorParsingRulesAPI.md#GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById) | **Get** /system/emailConnectors/{grandparentId}/parsingStyles/{parentId}/parsingRules/{id} | Get EmailConnectorParsingRule
[**GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesCount**](EmailConnectorParsingRulesAPI.md#GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesCount) | **Get** /system/emailConnectors/{grandparentId}/parsingStyles/{parentId}/parsingRules/count | Get Count of EmailConnectorParsingRule
[**PatchSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById**](EmailConnectorParsingRulesAPI.md#PatchSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById) | **Patch** /system/emailConnectors/{grandparentId}/parsingStyles/{parentId}/parsingRules/{id} | Patch EmailConnectorParsingRule
[**PostSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRules**](EmailConnectorParsingRulesAPI.md#PostSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRules) | **Post** /system/emailConnectors/{grandparentId}/parsingStyles/{parentId}/parsingRules | Post EmailConnectorParsingRule
[**PutSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById**](EmailConnectorParsingRulesAPI.md#PutSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById) | **Put** /system/emailConnectors/{grandparentId}/parsingStyles/{parentId}/parsingRules/{id} | Put EmailConnectorParsingRule



## DeleteSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById

> DeleteSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById(ctx, id, parentId, grandparentId).ClientId(clientId).Execute()

Delete EmailConnectorParsingRule

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
	id := int32(56) // int32 | parsingRuleId
	parentId := int32(56) // int32 | parsingStyleId
	grandparentId := int32(56) // int32 | emailConnectorId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmailConnectorParsingRulesAPI.DeleteSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById(context.Background(), id, parentId, grandparentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailConnectorParsingRulesAPI.DeleteSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | parsingRuleId | 
**parentId** | **int32** | parsingStyleId | 
**grandparentId** | **int32** | emailConnectorId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesByIdRequest struct via the builder pattern


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


## GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRules

> []EmailConnectorParsingRule GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRules(ctx, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of EmailConnectorParsingRule

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
	parentId := int32(56) // int32 | parsingStyleId
	grandparentId := int32(56) // int32 | emailConnectorId
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
	resp, r, err := apiClient.EmailConnectorParsingRulesAPI.GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRules(context.Background(), parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailConnectorParsingRulesAPI.GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRules`: []EmailConnectorParsingRule
	fmt.Fprintf(os.Stdout, "Response from `EmailConnectorParsingRulesAPI.GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | parsingStyleId | 
**grandparentId** | **int32** | emailConnectorId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesRequest struct via the builder pattern


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

[**[]EmailConnectorParsingRule**](EmailConnectorParsingRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById

> EmailConnectorParsingRule GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById(ctx, id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get EmailConnectorParsingRule

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
	id := int32(56) // int32 | parsingRuleId
	parentId := int32(56) // int32 | parsingStyleId
	grandparentId := int32(56) // int32 | emailConnectorId
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
	resp, r, err := apiClient.EmailConnectorParsingRulesAPI.GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById(context.Background(), id, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailConnectorParsingRulesAPI.GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById`: EmailConnectorParsingRule
	fmt.Fprintf(os.Stdout, "Response from `EmailConnectorParsingRulesAPI.GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | parsingRuleId | 
**parentId** | **int32** | parsingStyleId | 
**grandparentId** | **int32** | emailConnectorId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesByIdRequest struct via the builder pattern


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

[**EmailConnectorParsingRule**](EmailConnectorParsingRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesCount

> Count GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesCount(ctx, parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of EmailConnectorParsingRule

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
	parentId := int32(56) // int32 | parsingStyleId
	grandparentId := int32(56) // int32 | emailConnectorId
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
	resp, r, err := apiClient.EmailConnectorParsingRulesAPI.GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesCount(context.Background(), parentId, grandparentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailConnectorParsingRulesAPI.GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `EmailConnectorParsingRulesAPI.GetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | parsingStyleId | 
**grandparentId** | **int32** | emailConnectorId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesCountRequest struct via the builder pattern


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


## PatchSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById

> EmailConnectorParsingRule PatchSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById(ctx, id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch EmailConnectorParsingRule

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
	id := int32(56) // int32 | parsingRuleId
	parentId := int32(56) // int32 | parsingStyleId
	grandparentId := int32(56) // int32 | emailConnectorId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailConnectorParsingRulesAPI.PatchSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById(context.Background(), id, parentId, grandparentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailConnectorParsingRulesAPI.PatchSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById`: EmailConnectorParsingRule
	fmt.Fprintf(os.Stdout, "Response from `EmailConnectorParsingRulesAPI.PatchSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | parsingRuleId | 
**parentId** | **int32** | parsingStyleId | 
**grandparentId** | **int32** | emailConnectorId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**EmailConnectorParsingRule**](EmailConnectorParsingRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRules

> EmailConnectorParsingRule PostSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRules(ctx, parentId, grandparentId).ClientId(clientId).EmailConnectorParsingRule(emailConnectorParsingRule).Execute()

Post EmailConnectorParsingRule

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
	parentId := int32(56) // int32 | parsingStyleId
	grandparentId := int32(56) // int32 | emailConnectorId
	clientId := "clientId_example" // string | 
	emailConnectorParsingRule := *openapiclient.NewEmailConnectorParsingRule(NullableInt32(123), *openapiclient.NewEmailConnectorParsingVariableReference(), "SearchTerm_example") // EmailConnectorParsingRule | emailConnectorParsingRule

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailConnectorParsingRulesAPI.PostSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRules(context.Background(), parentId, grandparentId).ClientId(clientId).EmailConnectorParsingRule(emailConnectorParsingRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailConnectorParsingRulesAPI.PostSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRules`: EmailConnectorParsingRule
	fmt.Fprintf(os.Stdout, "Response from `EmailConnectorParsingRulesAPI.PostSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | parsingStyleId | 
**grandparentId** | **int32** | emailConnectorId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **emailConnectorParsingRule** | [**EmailConnectorParsingRule**](EmailConnectorParsingRule.md) | emailConnectorParsingRule | 

### Return type

[**EmailConnectorParsingRule**](EmailConnectorParsingRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById

> EmailConnectorParsingRule PutSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById(ctx, id, parentId, grandparentId).ClientId(clientId).EmailConnectorParsingRule(emailConnectorParsingRule).Execute()

Put EmailConnectorParsingRule

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
	id := int32(56) // int32 | parsingRuleId
	parentId := int32(56) // int32 | parsingStyleId
	grandparentId := int32(56) // int32 | emailConnectorId
	clientId := "clientId_example" // string | 
	emailConnectorParsingRule := *openapiclient.NewEmailConnectorParsingRule(NullableInt32(123), *openapiclient.NewEmailConnectorParsingVariableReference(), "SearchTerm_example") // EmailConnectorParsingRule | emailConnectorParsingRule

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailConnectorParsingRulesAPI.PutSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById(context.Background(), id, parentId, grandparentId).ClientId(clientId).EmailConnectorParsingRule(emailConnectorParsingRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailConnectorParsingRulesAPI.PutSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById`: EmailConnectorParsingRule
	fmt.Fprintf(os.Stdout, "Response from `EmailConnectorParsingRulesAPI.PutSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | parsingRuleId | 
**parentId** | **int32** | parsingStyleId | 
**grandparentId** | **int32** | emailConnectorId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemEmailConnectorsByGrandparentIdParsingStylesByParentIdParsingRulesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **string** |  | 
 **emailConnectorParsingRule** | [**EmailConnectorParsingRule**](EmailConnectorParsingRule.md) | emailConnectorParsingRule | 

### Return type

[**EmailConnectorParsingRule**](EmailConnectorParsingRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

