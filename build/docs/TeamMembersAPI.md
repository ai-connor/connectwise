# \TeamMembersAPI

All URIs are relative to *http://na.myconnectwise.net/v4_6_release/apis/3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostServiceTeamMembers**](TeamMembersAPI.md#PostServiceTeamMembers) | **Post** /service/teamMembers | Post TeamMember



## PostServiceTeamMembers

> TeamMember PostServiceTeamMembers(ctx).ClientId(clientId).TeamMember(teamMember).Execute()

Post TeamMember

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
	teamMember := *openapiclient.NewTeamMember(*openapiclient.NewServiceTeamReference(), *openapiclient.NewMemberReference()) // TeamMember | teamMember

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamMembersAPI.PostServiceTeamMembers(context.Background()).ClientId(clientId).TeamMember(teamMember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamMembersAPI.PostServiceTeamMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostServiceTeamMembers`: TeamMember
	fmt.Fprintf(os.Stdout, "Response from `TeamMembersAPI.PostServiceTeamMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostServiceTeamMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **teamMember** | [**TeamMember**](TeamMember.md) | teamMember | 

### Return type

[**TeamMember**](TeamMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.connectwise.com+json; version=2025.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

