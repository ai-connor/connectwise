# \MyAccountsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemMyAccountById**](MyAccountsAPI.md#GetSystemMyAccountById) | **Get** /system/myAccount/{id} | Get MyAccount
[**PatchSystemMyAccountById**](MyAccountsAPI.md#PatchSystemMyAccountById) | **Patch** /system/myAccount/{id} | Patch MyAccount
[**PutSystemMyAccountById**](MyAccountsAPI.md#PutSystemMyAccountById) | **Put** /system/myAccount/{id} | Put MyAccount



## GetSystemMyAccountById

> MyAccount GetSystemMyAccountById(ctx, id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get MyAccount

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
	id := int32(56) // int32 | memberId
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
	resp, r, err := apiClient.MyAccountsAPI.GetSystemMyAccountById(context.Background(), id).ClientId(clientId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountsAPI.GetSystemMyAccountById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemMyAccountById`: MyAccount
	fmt.Fprintf(os.Stdout, "Response from `MyAccountsAPI.GetSystemMyAccountById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemMyAccountByIdRequest struct via the builder pattern


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

[**MyAccount**](MyAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSystemMyAccountById

> MyAccount PatchSystemMyAccountById(ctx, id).ClientId(clientId).PatchOperation(patchOperation).Execute()

Patch MyAccount

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
	id := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation | List of PatchOperation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyAccountsAPI.PatchSystemMyAccountById(context.Background(), id).ClientId(clientId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountsAPI.PatchSystemMyAccountById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemMyAccountById`: MyAccount
	fmt.Fprintf(os.Stdout, "Response from `MyAccountsAPI.PatchSystemMyAccountById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemMyAccountByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) | List of PatchOperation | 

### Return type

[**MyAccount**](MyAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemMyAccountById

> MyAccount PutSystemMyAccountById(ctx, id).ClientId(clientId).MyAccount(myAccount).Execute()

Put MyAccount

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/ai-connor/connectwise"
)

func main() {
	id := int32(56) // int32 | memberId
	clientId := "clientId_example" // string | 
	myAccount := *openapiclient.NewMyAccount("Identifier_example", "FirstName_example", "LastName_example", "LicenseClass_example", *openapiclient.NewTimeZoneSetupReference(), "DefaultEmail_example", "DefaultPhone_example", *openapiclient.NewSystemLocationReference(), *openapiclient.NewSystemDepartmentReference(), *openapiclient.NewWorkRoleReference(), *openapiclient.NewMemberReference(), *openapiclient.NewMemberReference(), time.Now(), *openapiclient.NewSystemLocationReference(), "CompanyActivityTabFormat_example", "InvoiceTimeTabFormat_example", "InvoiceScreenDefaultTabFormat_example", "InvoicingDisplayOptions_example", "AgreementInvoicingDisplayOptions_example") // MyAccount | myAccount

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyAccountsAPI.PutSystemMyAccountById(context.Background(), id).ClientId(clientId).MyAccount(myAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccountsAPI.PutSystemMyAccountById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSystemMyAccountById`: MyAccount
	fmt.Fprintf(os.Stdout, "Response from `MyAccountsAPI.PutSystemMyAccountById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | memberId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemMyAccountByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 
 **myAccount** | [**MyAccount**](MyAccount.md) | myAccount | 

### Return type

[**MyAccount**](MyAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

