# \CompanyTeamsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCompanyCompaniesByParentIdTeamsById**](CompanyTeamsAPI.md#DeleteCompanyCompaniesByParentIdTeamsById) | **Delete** /company/companies/{parentId}/teams/{id} | Delete CompanyTeam
[**GetCompanyCompaniesByParentIdTeams**](CompanyTeamsAPI.md#GetCompanyCompaniesByParentIdTeams) | **Get** /company/companies/{parentId}/teams | Get List of CompanyTeam
[**GetCompanyCompaniesByParentIdTeamsById**](CompanyTeamsAPI.md#GetCompanyCompaniesByParentIdTeamsById) | **Get** /company/companies/{parentId}/teams/{id} | Get CompanyTeam
[**GetCompanyCompaniesByParentIdTeamsCount**](CompanyTeamsAPI.md#GetCompanyCompaniesByParentIdTeamsCount) | **Get** /company/companies/{parentId}/teams/count | Get Count of CompanyTeam
[**PatchCompanyCompaniesByParentIdTeamsById**](CompanyTeamsAPI.md#PatchCompanyCompaniesByParentIdTeamsById) | **Patch** /company/companies/{parentId}/teams/{id} | Patch CompanyTeam
[**PostCompanyCompaniesByParentIdTeams**](CompanyTeamsAPI.md#PostCompanyCompaniesByParentIdTeams) | **Post** /company/companies/{parentId}/teams | Post CompanyTeam
[**PutCompanyCompaniesByParentIdTeamsById**](CompanyTeamsAPI.md#PutCompanyCompaniesByParentIdTeamsById) | **Put** /company/companies/{parentId}/teams/{id} | Put CompanyTeam



## DeleteCompanyCompaniesByParentIdTeamsById

> DeleteCompanyCompaniesByParentIdTeamsById(ctx, id, parentId).ClientId(clientId).Execute()

Delete CompanyTeam

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
	parentId := int32(56) // int32 | companyId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CompanyTeamsAPI.DeleteCompanyCompaniesByParentIdTeamsById(context.Background(), id, parentId).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyTeamsAPI.DeleteCompanyCompaniesByParentIdTeamsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | teamId | 
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyCompaniesByParentIdTeamsByIdRequest struct via the builder pattern


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


## GetCompanyCompaniesByParentIdTeams

> []CompanyTeam GetCompanyCompaniesByParentIdTeams(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of CompanyTeam

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
	parentId := int32(56) // int32 | companyId
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
	resp, r, err := apiClient.CompanyTeamsAPI.GetCompanyCompaniesByParentIdTeams(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyTeamsAPI.GetCompanyCompaniesByParentIdTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCompaniesByParentIdTeams`: []CompanyTeam
	fmt.Fprintf(os.Stdout, "Response from `CompanyTeamsAPI.GetCompanyCompaniesByParentIdTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCompaniesByParentIdTeamsRequest struct via the builder pattern


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

[**[]CompanyTeam**](CompanyTeam.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyCompaniesByParentIdTeamsById

> CompanyTeam GetCompanyCompaniesByParentIdTeamsById(ctx, id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get CompanyTeam

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
	parentId := int32(56) // int32 | companyId
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
	resp, r, err := apiClient.CompanyTeamsAPI.GetCompanyCompaniesByParentIdTeamsById(context.Background(), id, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyTeamsAPI.GetCompanyCompaniesByParentIdTeamsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCompaniesByParentIdTeamsById`: CompanyTeam
	fmt.Fprintf(os.Stdout, "Response from `CompanyTeamsAPI.GetCompanyCompaniesByParentIdTeamsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | teamId | 
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCompaniesByParentIdTeamsByIdRequest struct via the builder pattern


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

[**CompanyTeam**](CompanyTeam.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyCompaniesByParentIdTeamsCount

> Count GetCompanyCompaniesByParentIdTeamsCount(ctx, parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of CompanyTeam

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
	parentId := int32(56) // int32 | companyId
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
	resp, r, err := apiClient.CompanyTeamsAPI.GetCompanyCompaniesByParentIdTeamsCount(context.Background(), parentId).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyTeamsAPI.GetCompanyCompaniesByParentIdTeamsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyCompaniesByParentIdTeamsCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `CompanyTeamsAPI.GetCompanyCompaniesByParentIdTeamsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyCompaniesByParentIdTeamsCountRequest struct via the builder pattern


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


## PatchCompanyCompaniesByParentIdTeamsById

> CompanyTeam PatchCompanyCompaniesByParentIdTeamsById(ctx, id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch CompanyTeam

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
	parentId := int32(56) // int32 | companyId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyTeamsAPI.PatchCompanyCompaniesByParentIdTeamsById(context.Background(), id, parentId).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyTeamsAPI.PatchCompanyCompaniesByParentIdTeamsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCompanyCompaniesByParentIdTeamsById`: CompanyTeam
	fmt.Fprintf(os.Stdout, "Response from `CompanyTeamsAPI.PatchCompanyCompaniesByParentIdTeamsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | teamId | 
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCompanyCompaniesByParentIdTeamsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**CompanyTeam**](CompanyTeam.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCompanyCompaniesByParentIdTeams

> CompanyTeam PostCompanyCompaniesByParentIdTeams(ctx, parentId).ClientId(clientId).CompanyTeam(companyTeam).Execute()

Post CompanyTeam

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
	parentId := int32(56) // int32 | companyId
	clientId := "clientId_example" // string | 
	companyTeam := *openapiclient.NewCompanyTeam(*openapiclient.NewTeamRoleReference()) // CompanyTeam | companyTeam

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyTeamsAPI.PostCompanyCompaniesByParentIdTeams(context.Background(), parentId).ClientId(clientId).CompanyTeam(companyTeam).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyTeamsAPI.PostCompanyCompaniesByParentIdTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyCompaniesByParentIdTeams`: CompanyTeam
	fmt.Fprintf(os.Stdout, "Response from `CompanyTeamsAPI.PostCompanyCompaniesByParentIdTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyCompaniesByParentIdTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **companyTeam** | [**CompanyTeam**](CompanyTeam.md) | companyTeam | 

### Return type

[**CompanyTeam**](CompanyTeam.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCompanyCompaniesByParentIdTeamsById

> CompanyTeam PutCompanyCompaniesByParentIdTeamsById(ctx, id, parentId).ClientId(clientId).CompanyTeam(companyTeam).Execute()

Put CompanyTeam

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
	parentId := int32(56) // int32 | companyId
	clientId := "clientId_example" // string | 
	companyTeam := *openapiclient.NewCompanyTeam(*openapiclient.NewTeamRoleReference()) // CompanyTeam | companyTeam

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompanyTeamsAPI.PutCompanyCompaniesByParentIdTeamsById(context.Background(), id, parentId).ClientId(clientId).CompanyTeam(companyTeam).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyTeamsAPI.PutCompanyCompaniesByParentIdTeamsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCompanyCompaniesByParentIdTeamsById`: CompanyTeam
	fmt.Fprintf(os.Stdout, "Response from `CompanyTeamsAPI.PutCompanyCompaniesByParentIdTeamsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | teamId | 
**parentId** | **int32** | companyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCompanyCompaniesByParentIdTeamsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** |  | 
 **companyTeam** | [**CompanyTeam**](CompanyTeam.md) | companyTeam | 

### Return type

[**CompanyTeam**](CompanyTeam.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

