# \OpportunityTeamsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSalesOpportunitiesByParentIdTeamById**](OpportunityTeamsAPI.md#DeleteSalesOpportunitiesByParentIdTeamById) | **Delete** /sales/opportunities/{parentId}/team/{id} | Delete Team
[**GetSalesOpportunitiesByParentIdTeam**](OpportunityTeamsAPI.md#GetSalesOpportunitiesByParentIdTeam) | **Get** /sales/opportunities/{parentId}/team | Get List of Team
[**GetSalesOpportunitiesByParentIdTeamById**](OpportunityTeamsAPI.md#GetSalesOpportunitiesByParentIdTeamById) | **Get** /sales/opportunities/{parentId}/team/{id} | Get Team
[**GetSalesOpportunitiesByParentIdTeamCount**](OpportunityTeamsAPI.md#GetSalesOpportunitiesByParentIdTeamCount) | **Get** /sales/opportunities/{parentId}/team/count | Get Count of Team
[**PatchSalesOpportunitiesByParentIdTeamById**](OpportunityTeamsAPI.md#PatchSalesOpportunitiesByParentIdTeamById) | **Patch** /sales/opportunities/{parentId}/team/{id} | Patch Team
[**PostSalesOpportunitiesByParentIdTeam**](OpportunityTeamsAPI.md#PostSalesOpportunitiesByParentIdTeam) | **Post** /sales/opportunities/{parentId}/team | Post Team
[**PutSalesOpportunitiesByParentIdTeamById**](OpportunityTeamsAPI.md#PutSalesOpportunitiesByParentIdTeamById) | **Put** /sales/opportunities/{parentId}/team/{id} | Put Team



## DeleteSalesOpportunitiesByParentIdTeamById

> DeleteSalesOpportunitiesByParentIdTeamById(ctx, id, parentId).ClientId(clientId).Execute()

Delete Team

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
	id := int32(56) // int32 | teamId
	parentId := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OpportunityTeamsAPI.DeleteSalesOpportunitiesByParentIdTeamById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityTeamsAPI.DeleteSalesOpportunitiesByParentIdTeamById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | teamId | 
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSalesOpportunitiesByParentIdTeamByIdRequest struct via the builder pattern


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


## GetSalesOpportunitiesByParentIdTeam

> []Team GetSalesOpportunitiesByParentIdTeam(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of Team

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
	parentId := int32(56) // int32 | opportunityId
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
	resp, r, err := apiClient.OpportunityTeamsAPI.GetSalesOpportunitiesByParentIdTeam(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityTeamsAPI.GetSalesOpportunitiesByParentIdTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesOpportunitiesByParentIdTeam`: []Team
	fmt.Fprintf(os.Stdout, "Response from `OpportunityTeamsAPI.GetSalesOpportunitiesByParentIdTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesOpportunitiesByParentIdTeamRequest struct via the builder pattern


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

[**[]Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalesOpportunitiesByParentIdTeamById

> Team GetSalesOpportunitiesByParentIdTeamById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Team

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
	id := int32(56) // int32 | teamId
	parentId := int32(56) // int32 | opportunityId
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
	resp, r, err := apiClient.OpportunityTeamsAPI.GetSalesOpportunitiesByParentIdTeamById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityTeamsAPI.GetSalesOpportunitiesByParentIdTeamById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesOpportunitiesByParentIdTeamById`: Team
	fmt.Fprintf(os.Stdout, "Response from `OpportunityTeamsAPI.GetSalesOpportunitiesByParentIdTeamById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | teamId | 
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesOpportunitiesByParentIdTeamByIdRequest struct via the builder pattern


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

[**Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalesOpportunitiesByParentIdTeamCount

> Count GetSalesOpportunitiesByParentIdTeamCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of Team

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
	parentId := int32(56) // int32 | opportunityId
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
	resp, r, err := apiClient.OpportunityTeamsAPI.GetSalesOpportunitiesByParentIdTeamCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityTeamsAPI.GetSalesOpportunitiesByParentIdTeamCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesOpportunitiesByParentIdTeamCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `OpportunityTeamsAPI.GetSalesOpportunitiesByParentIdTeamCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesOpportunitiesByParentIdTeamCountRequest struct via the builder pattern


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


## PatchSalesOpportunitiesByParentIdTeamById

> Team PatchSalesOpportunitiesByParentIdTeamById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch Team

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
	id := int32(56) // int32 | teamId
	parentId := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpportunityTeamsAPI.PatchSalesOpportunitiesByParentIdTeamById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityTeamsAPI.PatchSalesOpportunitiesByParentIdTeamById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSalesOpportunitiesByParentIdTeamById`: Team
	fmt.Fprintf(os.Stdout, "Response from `OpportunityTeamsAPI.PatchSalesOpportunitiesByParentIdTeamById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | teamId | 
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSalesOpportunitiesByParentIdTeamByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSalesOpportunitiesByParentIdTeam

> Team PostSalesOpportunitiesByParentIdTeam(ctx, parentId).ClientId(clientId).Team(team).Execute()

Post Team

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
	parentId := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 
	team := *openapiclient.NewTeam("Type_example") // Team | team

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpportunityTeamsAPI.PostSalesOpportunitiesByParentIdTeam(context.Background(), parentId).ClientId(clientId).Team(team).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityTeamsAPI.PostSalesOpportunitiesByParentIdTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSalesOpportunitiesByParentIdTeam`: Team
	fmt.Fprintf(os.Stdout, "Response from `OpportunityTeamsAPI.PostSalesOpportunitiesByParentIdTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSalesOpportunitiesByParentIdTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **team** | [**Team**](Team.md) | team | 

### Return type

[**Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSalesOpportunitiesByParentIdTeamById

> Team PutSalesOpportunitiesByParentIdTeamById(ctx, id, parentId).ClientId(clientId).Team(team).Execute()

Put Team

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
	id := int32(56) // int32 | teamId
	parentId := int32(56) // int32 | opportunityId
	clientId := "clientId_example" // string | 
	team := *openapiclient.NewTeam("Type_example") // Team | team

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpportunityTeamsAPI.PutSalesOpportunitiesByParentIdTeamById(context.Background(), id, parentId).ClientId(clientId).Team(team).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpportunityTeamsAPI.PutSalesOpportunitiesByParentIdTeamById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSalesOpportunitiesByParentIdTeamById`: Team
	fmt.Fprintf(os.Stdout, "Response from `OpportunityTeamsAPI.PutSalesOpportunitiesByParentIdTeamById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | teamId | 
**parentId** | **int32** | opportunityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSalesOpportunitiesByParentIdTeamByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **team** | [**Team**](Team.md) | team | 

### Return type

[**Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

