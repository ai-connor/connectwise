# \OpportunitiesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSalesOpportunitiesById**](OpportunitiesAPI.md#DeleteSalesOpportunitiesById) | **Delete** /sales/opportunities/{id} | Delete ConnectWise.Apis.v3_0.v2015_3.Sales.Opportunity.Opportunity
[**GetSalesOpportunities**](OpportunitiesAPI.md#GetSalesOpportunities) | **Get** /sales/opportunities | Get List of ConnectWise.Apis.v3_0.v2015_3.Sales.Opportunity.Opportunity
[**GetSalesOpportunitiesById**](OpportunitiesAPI.md#GetSalesOpportunitiesById) | **Get** /sales/opportunities/{id} | Get ConnectWise.Apis.v3_0.v2015_3.Sales.Opportunity.Opportunity
[**GetSalesOpportunitiesConversionsById**](OpportunitiesAPI.md#GetSalesOpportunitiesConversionsById) | **Get** /sales/opportunities/conversions/{id} | Get Conversion
[**GetSalesOpportunitiesCount**](OpportunitiesAPI.md#GetSalesOpportunitiesCount) | **Get** /sales/opportunities/count | Get Count of ConnectWise.Apis.v3_0.v2015_3.Sales.Opportunity.Opportunity
[**GetSalesOpportunitiesDefault**](OpportunitiesAPI.md#GetSalesOpportunitiesDefault) | **Get** /sales/opportunities/default | Get ConnectWise.Apis.v3_0.v2015_3.Sales.Opportunity.Opportunity
[**PatchSalesOpportunitiesById**](OpportunitiesAPI.md#PatchSalesOpportunitiesById) | **Patch** /sales/opportunities/{id} | Patch ConnectWise.Apis.v3_0.v2015_3.Sales.Opportunity.Opportunity
[**PostSalesOpportunities**](OpportunitiesAPI.md#PostSalesOpportunities) | **Post** /sales/opportunities | Post ConnectWise.Apis.v3_0.v2015_3.Sales.Opportunity.Opportunity
[**PostSalesOpportunitiesByIdConvertToAgreement**](OpportunitiesAPI.md#PostSalesOpportunitiesByIdConvertToAgreement) | **Post** /sales/opportunities/{id}/convertToAgreement | Post ApiAgreement
[**PostSalesOpportunitiesByIdConvertToProject**](OpportunitiesAPI.md#PostSalesOpportunitiesByIdConvertToProject) | **Post** /sales/opportunities/{id}/convertToProject | Post ApiProject
[**PostSalesOpportunitiesByIdConvertToSalesOrder**](OpportunitiesAPI.md#PostSalesOpportunitiesByIdConvertToSalesOrder) | **Post** /sales/opportunities/{id}/convertToSalesOrder | Post ApiSalesOrder
[**PostSalesOpportunitiesByIdConvertToServiceTicket**](OpportunitiesAPI.md#PostSalesOpportunitiesByIdConvertToServiceTicket) | **Post** /sales/opportunities/{id}/convertToServiceTicket | Post ApiTicket
[**PutSalesOpportunitiesById**](OpportunitiesAPI.md#PutSalesOpportunitiesById) | **Put** /sales/opportunities/{id} | Put ConnectWise.Apis.v3_0.v2015_3.Sales.Opportunity.Opportunity



## DeleteSalesOpportunitiesById

> DeleteSalesOpportunitiesById(ctx, id).ClientId(clientId).Execute()

Delete ConnectWise.Apis.v3_0.v2015_3.Sales.Opportunity.Opportunity

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
	id := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OpportunitiesAPI.DeleteSalesOpportunitiesById(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunitiesAPI.DeleteSalesOpportunitiesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSalesOpportunitiesByIdRequest struct via the builder pattern


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


## GetSalesOpportunities

> []Opportunity GetSalesOpportunities(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of ConnectWise.Apis.v3_0.v2015_3.Sales.Opportunity.Opportunity

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
	resp, r, err := apiClient.OpportunitiesAPI.GetSalesOpportunities(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunitiesAPI.GetSalesOpportunities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesOpportunities`: []Opportunity
	fmt.Fprintf(os.Stdout, "Response from `OpportunitiesAPI.GetSalesOpportunities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesOpportunitiesRequest struct via the builder pattern


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

[**[]Opportunity**](Opportunity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalesOpportunitiesById

> Opportunity GetSalesOpportunitiesById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ConnectWise.Apis.v3_0.v2015_3.Sales.Opportunity.Opportunity

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
	id := int32(56) // int32 | opportunityId
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
	resp, r, err := apiClient.OpportunitiesAPI.GetSalesOpportunitiesById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunitiesAPI.GetSalesOpportunitiesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesOpportunitiesById`: Opportunity
	fmt.Fprintf(os.Stdout, "Response from `OpportunitiesAPI.GetSalesOpportunitiesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesOpportunitiesByIdRequest struct via the builder pattern


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

[**Opportunity**](Opportunity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalesOpportunitiesConversionsById

> []SalesConversion GetSalesOpportunitiesConversionsById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Conversion

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
	id := int32(56) // int32 | opportunityId
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
	resp, r, err := apiClient.OpportunitiesAPI.GetSalesOpportunitiesConversionsById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunitiesAPI.GetSalesOpportunitiesConversionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesOpportunitiesConversionsById`: []SalesConversion
	fmt.Fprintf(os.Stdout, "Response from `OpportunitiesAPI.GetSalesOpportunitiesConversionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesOpportunitiesConversionsByIdRequest struct via the builder pattern


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

[**[]SalesConversion**](SalesConversion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalesOpportunitiesCount

> Count GetSalesOpportunitiesCount(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of ConnectWise.Apis.v3_0.v2015_3.Sales.Opportunity.Opportunity

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
	resp, r, err := apiClient.OpportunitiesAPI.GetSalesOpportunitiesCount(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunitiesAPI.GetSalesOpportunitiesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesOpportunitiesCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `OpportunitiesAPI.GetSalesOpportunitiesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesOpportunitiesCountRequest struct via the builder pattern


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


## GetSalesOpportunitiesDefault

> Opportunity GetSalesOpportunitiesDefault(ctx).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get ConnectWise.Apis.v3_0.v2015_3.Sales.Opportunity.Opportunity

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
	resp, r, err := apiClient.OpportunitiesAPI.GetSalesOpportunitiesDefault(context.Background()).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunitiesAPI.GetSalesOpportunitiesDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesOpportunitiesDefault`: Opportunity
	fmt.Fprintf(os.Stdout, "Response from `OpportunitiesAPI.GetSalesOpportunitiesDefault`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesOpportunitiesDefaultRequest struct via the builder pattern


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

[**Opportunity**](Opportunity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSalesOpportunitiesById

> Opportunity PatchSalesOpportunitiesById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch ConnectWise.Apis.v3_0.v2015_3.Sales.Opportunity.Opportunity

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
	id := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpportunitiesAPI.PatchSalesOpportunitiesById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunitiesAPI.PatchSalesOpportunitiesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSalesOpportunitiesById`: Opportunity
	fmt.Fprintf(os.Stdout, "Response from `OpportunitiesAPI.PatchSalesOpportunitiesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSalesOpportunitiesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**Opportunity**](Opportunity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSalesOpportunities

> Opportunity PostSalesOpportunities(ctx).ClientId(clientId).Opportunity(opportunity).Execute()

Post ConnectWise.Apis.v3_0.v2015_3.Sales.Opportunity.Opportunity

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
	opportunity := *openapiclient.NewOpportunity("Name_example", *openapiclient.NewMemberReference(), *openapiclient.NewCompanyReference(), *openapiclient.NewContactReference()) // Opportunity | opportunity

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpportunitiesAPI.PostSalesOpportunities(context.Background()).ClientId(clientId).Opportunity(opportunity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunitiesAPI.PostSalesOpportunities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSalesOpportunities`: Opportunity
	fmt.Fprintf(os.Stdout, "Response from `OpportunitiesAPI.PostSalesOpportunities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSalesOpportunitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **opportunity** | [**Opportunity**](Opportunity.md) | opportunity | 

### Return type

[**Opportunity**](Opportunity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSalesOpportunitiesByIdConvertToAgreement

> Agreement PostSalesOpportunitiesByIdConvertToAgreement(ctx, id).ClientId(clientId).OpportunityToAgreementConversion(opportunityToAgreementConversion).Execute()

Post ApiAgreement

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
	id := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 
	opportunityToAgreementConversion := *openapiclient.NewOpportunityToAgreementConversion() // OpportunityToAgreementConversion | conversion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpportunitiesAPI.PostSalesOpportunitiesByIdConvertToAgreement(context.Background(), id).ClientId(clientId).OpportunityToAgreementConversion(opportunityToAgreementConversion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunitiesAPI.PostSalesOpportunitiesByIdConvertToAgreement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSalesOpportunitiesByIdConvertToAgreement`: Agreement
	fmt.Fprintf(os.Stdout, "Response from `OpportunitiesAPI.PostSalesOpportunitiesByIdConvertToAgreement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSalesOpportunitiesByIdConvertToAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **opportunityToAgreementConversion** | [**OpportunityToAgreementConversion**](OpportunityToAgreementConversion.md) | conversion | 

### Return type

[**Agreement**](Agreement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSalesOpportunitiesByIdConvertToProject

> Project PostSalesOpportunitiesByIdConvertToProject(ctx, id).ClientId(clientId).OpportunityToProjectConversion(opportunityToProjectConversion).Execute()

Post ApiProject

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
	id := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 
	opportunityToProjectConversion := *openapiclient.NewOpportunityToProjectConversion() // OpportunityToProjectConversion | conversion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpportunitiesAPI.PostSalesOpportunitiesByIdConvertToProject(context.Background(), id).ClientId(clientId).OpportunityToProjectConversion(opportunityToProjectConversion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunitiesAPI.PostSalesOpportunitiesByIdConvertToProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSalesOpportunitiesByIdConvertToProject`: Project
	fmt.Fprintf(os.Stdout, "Response from `OpportunitiesAPI.PostSalesOpportunitiesByIdConvertToProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSalesOpportunitiesByIdConvertToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **opportunityToProjectConversion** | [**OpportunityToProjectConversion**](OpportunityToProjectConversion.md) | conversion | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSalesOpportunitiesByIdConvertToSalesOrder

> Order PostSalesOpportunitiesByIdConvertToSalesOrder(ctx, id).ClientId(clientId).OpportunityToSalesOrderConversion(opportunityToSalesOrderConversion).Execute()

Post ApiSalesOrder

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
	id := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 
	opportunityToSalesOrderConversion := *openapiclient.NewOpportunityToSalesOrderConversion() // OpportunityToSalesOrderConversion | conversion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpportunitiesAPI.PostSalesOpportunitiesByIdConvertToSalesOrder(context.Background(), id).ClientId(clientId).OpportunityToSalesOrderConversion(opportunityToSalesOrderConversion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunitiesAPI.PostSalesOpportunitiesByIdConvertToSalesOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSalesOpportunitiesByIdConvertToSalesOrder`: Order
	fmt.Fprintf(os.Stdout, "Response from `OpportunitiesAPI.PostSalesOpportunitiesByIdConvertToSalesOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSalesOpportunitiesByIdConvertToSalesOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **opportunityToSalesOrderConversion** | [**OpportunityToSalesOrderConversion**](OpportunityToSalesOrderConversion.md) | conversion | 

### Return type

[**Order**](Order.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSalesOpportunitiesByIdConvertToServiceTicket

> Ticket PostSalesOpportunitiesByIdConvertToServiceTicket(ctx, id).ClientId(clientId).OpportunityToServiceTicketConversion(opportunityToServiceTicketConversion).Execute()

Post ApiTicket

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
	id := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 
	opportunityToServiceTicketConversion := *openapiclient.NewOpportunityToServiceTicketConversion() // OpportunityToServiceTicketConversion | conversion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpportunitiesAPI.PostSalesOpportunitiesByIdConvertToServiceTicket(context.Background(), id).ClientId(clientId).OpportunityToServiceTicketConversion(opportunityToServiceTicketConversion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunitiesAPI.PostSalesOpportunitiesByIdConvertToServiceTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSalesOpportunitiesByIdConvertToServiceTicket`: Ticket
	fmt.Fprintf(os.Stdout, "Response from `OpportunitiesAPI.PostSalesOpportunitiesByIdConvertToServiceTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSalesOpportunitiesByIdConvertToServiceTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **opportunityToServiceTicketConversion** | [**OpportunityToServiceTicketConversion**](OpportunityToServiceTicketConversion.md) | conversion | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSalesOpportunitiesById

> Opportunity PutSalesOpportunitiesById(ctx, id).ClientId(clientId).Opportunity(opportunity).Execute()

Put ConnectWise.Apis.v3_0.v2015_3.Sales.Opportunity.Opportunity

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
	id := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 
	opportunity := *openapiclient.NewOpportunity("Name_example", *openapiclient.NewMemberReference(), *openapiclient.NewCompanyReference(), *openapiclient.NewContactReference()) // Opportunity | opportunity

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpportunitiesAPI.PutSalesOpportunitiesById(context.Background(), id).ClientId(clientId).Opportunity(opportunity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunitiesAPI.PutSalesOpportunitiesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSalesOpportunitiesById`: Opportunity
	fmt.Fprintf(os.Stdout, "Response from `OpportunitiesAPI.PutSalesOpportunitiesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSalesOpportunitiesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **opportunity** | [**Opportunity**](Opportunity.md) | opportunity | 

### Return type

[**Opportunity**](Opportunity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

