# \AuditTrailAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemAudittrail**](AuditTrailAPI.md#GetSystemAudittrail) | **Get** /system/audittrail | Get List of AuditTrailEntry
[**GetSystemAudittrailCount**](AuditTrailAPI.md#GetSystemAudittrailCount) | **Get** /system/audittrail/count | Get Count of AuditTrailEntry



## GetSystemAudittrail

> []AuditTrailEntry GetSystemAudittrail(ctx).ClientId(clientId).Type_(type_).Id(id).DeviceIdentifier(deviceIdentifier).XrefRecId(xrefRecId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get List of AuditTrailEntry

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
	type_ := "type__example" // string | type (optional)
	id := int32(56) // int32 | id (optional)
	deviceIdentifier := "deviceIdentifier_example" // string | deviceIdentifier (optional)
	xrefRecId := int32(56) // int32 | xrefRecId (optional)
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
	resp, r, err := apiClient.AuditTrailAPI.GetSystemAudittrail(context.Background()).ClientId(clientId).Type_(type_).Id(id).DeviceIdentifier(deviceIdentifier).XrefRecId(xrefRecId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditTrailAPI.GetSystemAudittrail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemAudittrail`: []AuditTrailEntry
	fmt.Fprintf(os.Stdout, "Response from `AuditTrailAPI.GetSystemAudittrail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemAudittrailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **type_** | **string** | type | 
 **id** | **int32** | id | 
 **deviceIdentifier** | **string** | deviceIdentifier | 
 **xrefRecId** | **int32** | xrefRecId | 
 **conditions** | **string** |  | 
 **childConditions** | **string** |  | 
 **customFieldConditions** | **string** |  | 
 **orderBy** | **string** |  | 
 **fields** | **string** |  | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **pageId** | **int32** |  | 

### Return type

[**[]AuditTrailEntry**](AuditTrailEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemAudittrailCount

> Count GetSystemAudittrailCount(ctx).ClientId(clientId).Type_(type_).Id(id).DeviceIdentifier(deviceIdentifier).XrefRecId(xrefRecId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()

Get Count of AuditTrailEntry

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
	type_ := "type__example" // string | type (optional)
	id := int32(56) // int32 | id (optional)
	deviceIdentifier := "deviceIdentifier_example" // string | deviceIdentifier (optional)
	xrefRecId := int32(56) // int32 | xrefRecId (optional)
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
	resp, r, err := apiClient.AuditTrailAPI.GetSystemAudittrailCount(context.Background()).ClientId(clientId).Type_(type_).Id(id).DeviceIdentifier(deviceIdentifier).XrefRecId(xrefRecId).Conditions(conditions).ChildConditions(childConditions).CustomFieldConditions(customFieldConditions).OrderBy(orderBy).Fields(fields).Page(page).PageSize(pageSize).PageId(pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditTrailAPI.GetSystemAudittrailCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemAudittrailCount`: Count
	fmt.Fprintf(os.Stdout, "Response from `AuditTrailAPI.GetSystemAudittrailCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemAudittrailCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **type_** | **string** | type | 
 **id** | **int32** | id | 
 **deviceIdentifier** | **string** | deviceIdentifier | 
 **xrefRecId** | **int32** | xrefRecId | 
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

