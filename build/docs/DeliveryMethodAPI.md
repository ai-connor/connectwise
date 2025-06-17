# \DeliveryMethodAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFinanceDeliveryMethodsById**](DeliveryMethodAPI.md#DeleteFinanceDeliveryMethodsById) | **Delete** /finance/deliveryMethods/{id} | Delete DeliveryMethod
[**GetFinanceDeliveryMethods**](DeliveryMethodAPI.md#GetFinanceDeliveryMethods) | **Get** /finance/deliveryMethods | Get List of DeliveryMethod
[**GetFinanceDeliveryMethodsById**](DeliveryMethodAPI.md#GetFinanceDeliveryMethodsById) | **Get** /finance/deliveryMethods/{id} | Get DeliveryMethod
[**GetFinanceDeliveryMethodsCount**](DeliveryMethodAPI.md#GetFinanceDeliveryMethodsCount) | **Get** /finance/deliveryMethods/count | Get Count of DeliveryMethod
[**PatchFinanceDeliveryMethodsById**](DeliveryMethodAPI.md#PatchFinanceDeliveryMethodsById) | **Patch** /finance/deliveryMethods/{id} | Patch DeliveryMethod
[**PostFinanceDeliveryMethods**](DeliveryMethodAPI.md#PostFinanceDeliveryMethods) | **Post** /finance/deliveryMethods | Post DeliveryMethod
[**PutFinanceDeliveryMethodsById**](DeliveryMethodAPI.md#PutFinanceDeliveryMethodsById) | **Put** /finance/deliveryMethods/{id} | Put DeliveryMethod



## DeleteFinanceDeliveryMethodsById

> DeleteFinanceDeliveryMethodsById(ctx, id).ClientId(clientId).Execute()

Delete DeliveryMethod

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
	id := int32(56) // int32 | deliveryMethodId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeliveryMethodAPI.DeleteFinanceDeliveryMethodsById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeliveryMethodAPI.DeleteFinanceDeliveryMethodsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | deliveryMethodId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinanceDeliveryMethodsByIdRequest struct via the builder pattern


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


## GetFinanceDeliveryMethods

> []DeliveryMethod GetFinanceDeliveryMethods(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of DeliveryMethod

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
	resp, r, err := apiClient.DeliveryMethodAPI.GetFinanceDeliveryMethods(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeliveryMethodAPI.GetFinanceDeliveryMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceDeliveryMethods`: []DeliveryMethod
	fmt.Fprintf(os.Stdout, "Response from `DeliveryMethodAPI.GetFinanceDeliveryMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceDeliveryMethodsRequest struct via the builder pattern


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

[**[]DeliveryMethod**](DeliveryMethod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceDeliveryMethodsById

> DeliveryMethod GetFinanceDeliveryMethodsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get DeliveryMethod

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
	id := int32(56) // int32 | deliveryMethodId
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
	resp, r, err := apiClient.DeliveryMethodAPI.GetFinanceDeliveryMethodsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeliveryMethodAPI.GetFinanceDeliveryMethodsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceDeliveryMethodsById`: DeliveryMethod
	fmt.Fprintf(os.Stdout, "Response from `DeliveryMethodAPI.GetFinanceDeliveryMethodsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | deliveryMethodId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceDeliveryMethodsByIdRequest struct via the builder pattern


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

[**DeliveryMethod**](DeliveryMethod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceDeliveryMethodsCount

> Count GetFinanceDeliveryMethodsCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of DeliveryMethod

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
	resp, r, err := apiClient.DeliveryMethodAPI.GetFinanceDeliveryMethodsCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeliveryMethodAPI.GetFinanceDeliveryMethodsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceDeliveryMethodsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `DeliveryMethodAPI.GetFinanceDeliveryMethodsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceDeliveryMethodsCountRequest struct via the builder pattern


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


## PatchFinanceDeliveryMethodsById

> DeliveryMethod PatchFinanceDeliveryMethodsById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch DeliveryMethod

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
	id := int32(56) // int32 | deliveryMethodId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeliveryMethodAPI.PatchFinanceDeliveryMethodsById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeliveryMethodAPI.PatchFinanceDeliveryMethodsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFinanceDeliveryMethodsById`: DeliveryMethod
	fmt.Fprintf(os.Stdout, "Response from `DeliveryMethodAPI.PatchFinanceDeliveryMethodsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | deliveryMethodId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFinanceDeliveryMethodsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**DeliveryMethod**](DeliveryMethod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFinanceDeliveryMethods

> DeliveryMethod PostFinanceDeliveryMethods(ctx).ClientId(clientId).DeliveryMethod(deliveryMethod).Execute()

Post DeliveryMethod

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
	deliveryMethod := *openapiclient.NewDeliveryMethod("Name_example") // DeliveryMethod | deliveryMethod

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeliveryMethodAPI.PostFinanceDeliveryMethods(context.Background()).ClientId(clientId).DeliveryMethod(deliveryMethod).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeliveryMethodAPI.PostFinanceDeliveryMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFinanceDeliveryMethods`: DeliveryMethod
	fmt.Fprintf(os.Stdout, "Response from `DeliveryMethodAPI.PostFinanceDeliveryMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFinanceDeliveryMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **deliveryMethod** | [**DeliveryMethod**](DeliveryMethod.md) | deliveryMethod | 

### Return type

[**DeliveryMethod**](DeliveryMethod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFinanceDeliveryMethodsById

> DeliveryMethod PutFinanceDeliveryMethodsById(ctx, id).ClientId(clientId).DeliveryMethod(deliveryMethod).Execute()

Put DeliveryMethod

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
	id := int32(56) // int32 | deliveryMethodId
	clientId := "clientId_example" // string | 
	deliveryMethod := *openapiclient.NewDeliveryMethod("Name_example") // DeliveryMethod | deliveryMethod

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeliveryMethodAPI.PutFinanceDeliveryMethodsById(context.Background(), id).ClientId(clientId).DeliveryMethod(deliveryMethod).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeliveryMethodAPI.PutFinanceDeliveryMethodsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFinanceDeliveryMethodsById`: DeliveryMethod
	fmt.Fprintf(os.Stdout, "Response from `DeliveryMethodAPI.PutFinanceDeliveryMethodsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | deliveryMethodId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFinanceDeliveryMethodsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **deliveryMethod** | [**DeliveryMethod**](DeliveryMethod.md) | deliveryMethod | 

### Return type

[**DeliveryMethod**](DeliveryMethod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

