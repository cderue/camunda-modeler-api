# \CollaboratorsAPI

All URIs are relative to *https://modeler.cloud.camunda.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCollaborator**](CollaboratorsAPI.md#DeleteCollaborator) | **Delete** /api/v1/projects/{projectId}/collaborators/{email} | 
[**ModifyCollaborator**](CollaboratorsAPI.md#ModifyCollaborator) | **Put** /api/v1/collaborators | 
[**SearchCollaborators**](CollaboratorsAPI.md#SearchCollaborators) | **Post** /api/v1/collaborators/search | 



## DeleteCollaborator

> DeleteCollaborator(ctx, projectId, email).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	projectId := "7e39bfc3-32b2-46a7-b9d0-95f6549cd85e" // string | 
	email := "email_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollaboratorsAPI.DeleteCollaborator(context.Background(), projectId, email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollaboratorsAPI.DeleteCollaborator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCollaboratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyCollaborator

> ModifyCollaborator(ctx).CreateCollaboratorDto(createCollaboratorDto).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createCollaboratorDto := *openapiclient.NewCreateCollaboratorDto("Email_example", "ProjectId_example", "Role_example") // CreateCollaboratorDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollaboratorsAPI.ModifyCollaborator(context.Background()).CreateCollaboratorDto(createCollaboratorDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollaboratorsAPI.ModifyCollaborator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyCollaboratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCollaboratorDto** | [**CreateCollaboratorDto**](CreateCollaboratorDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCollaborators

> PubSearchResultDtoProjectCollaboratorDto SearchCollaborators(ctx).PubSearchDtoProjectCollaboratorDto(pubSearchDtoProjectCollaboratorDto).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	pubSearchDtoProjectCollaboratorDto := *openapiclient.NewPubSearchDtoProjectCollaboratorDto() // PubSearchDtoProjectCollaboratorDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollaboratorsAPI.SearchCollaborators(context.Background()).PubSearchDtoProjectCollaboratorDto(pubSearchDtoProjectCollaboratorDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollaboratorsAPI.SearchCollaborators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCollaborators`: PubSearchResultDtoProjectCollaboratorDto
	fmt.Fprintf(os.Stdout, "Response from `CollaboratorsAPI.SearchCollaborators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCollaboratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pubSearchDtoProjectCollaboratorDto** | [**PubSearchDtoProjectCollaboratorDto**](PubSearchDtoProjectCollaboratorDto.md) |  | 

### Return type

[**PubSearchResultDtoProjectCollaboratorDto**](PubSearchResultDtoProjectCollaboratorDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

