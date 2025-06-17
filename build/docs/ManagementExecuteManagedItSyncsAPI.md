# \ManagementExecuteManagedItSyncsAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCompanyManagementByIdExecuteManagedItSync**](ManagementExecuteManagedItSyncsAPI.md#PostCompanyManagementByIdExecuteManagedItSync) | **Post** /company/management/{id}/executeManagedItSync | Post SuccessResponse



## PostCompanyManagementByIdExecuteManagedItSync

> SuccessResponse PostCompanyManagementByIdExecuteManagedItSync(ctx, id).ClientId(clientId).Execute()

Post SuccessResponse

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
	id := int32(56) // int32 | managementId
	clientId := "clientId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementExecuteManagedItSyncsAPI.PostCompanyManagementByIdExecuteManagedItSync(context.Background(), id).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementExecuteManagedItSyncsAPI.PostCompanyManagementByIdExecuteManagedItSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCompanyManagementByIdExecuteManagedItSync`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagementExecuteManagedItSyncsAPI.PostCompanyManagementByIdExecuteManagedItSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | managementId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCompanyManagementByIdExecuteManagedItSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** |  | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

