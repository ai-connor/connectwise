# \BundlesAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostSystemBundles**](BundlesAPI.md#PostSystemBundles) | **Post** /system/bundles | Post BundleResultsCollection
[**PostSystemBundlesCount**](BundlesAPI.md#PostSystemBundlesCount) | **Post** /system/bundles/count | Post BundleResultsCollection



## PostSystemBundles

> BundleResultsCollection PostSystemBundles(ctx).ClientId(clientId).BundleRequestsCollection(bundleRequestsCollection).Execute()

Post BundleResultsCollection

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
	bundleRequestsCollection := *openapiclient.NewBundleRequestsCollection([]openapiclient.BundleRequest{*openapiclient.NewBundleRequest()}) // BundleRequestsCollection | requests

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BundlesAPI.PostSystemBundles(context.Background()).ClientId(clientId).BundleRequestsCollection(bundleRequestsCollection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundlesAPI.PostSystemBundles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemBundles`: BundleResultsCollection
	fmt.Fprintf(os.Stdout, "Response from `BundlesAPI.PostSystemBundles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemBundlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **bundleRequestsCollection** | [**BundleRequestsCollection**](BundleRequestsCollection.md) | requests | 

### Return type

[**BundleResultsCollection**](BundleResultsCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemBundlesCount

> BundleResultsCollection PostSystemBundlesCount(ctx).ClientId(clientId).BundleRequestsCollection(bundleRequestsCollection).Execute()

Post BundleResultsCollection

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
	bundleRequestsCollection := *openapiclient.NewBundleRequestsCollection([]openapiclient.BundleRequest{*openapiclient.NewBundleRequest()}) // BundleRequestsCollection | requests

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BundlesAPI.PostSystemBundlesCount(context.Background()).ClientId(clientId).BundleRequestsCollection(bundleRequestsCollection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundlesAPI.PostSystemBundlesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSystemBundlesCount`: BundleResultsCollection
	fmt.Fprintf(os.Stdout, "Response from `BundlesAPI.PostSystemBundlesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSystemBundlesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **bundleRequestsCollection** | [**BundleRequestsCollection**](BundleRequestsCollection.md) | requests | 

### Return type

[**BundleResultsCollection**](BundleResultsCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

