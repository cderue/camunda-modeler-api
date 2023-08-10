# \MilestonesAPI

All URIs are relative to *https://modeler.cloud.camunda.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompareMilestones**](MilestonesAPI.md#CompareMilestones) | **Get** /api/v1/milestones/compare/{milestone1Id}...{milestone2Id} | 
[**CreateMilestone**](MilestonesAPI.md#CreateMilestone) | **Post** /api/v1/milestones | 
[**DeleteMilestone**](MilestonesAPI.md#DeleteMilestone) | **Delete** /api/v1/milestones/{milestoneId} | 
[**GetMilestone**](MilestonesAPI.md#GetMilestone) | **Get** /api/v1/milestones/{milestoneId} | 
[**PatchMilestone**](MilestonesAPI.md#PatchMilestone) | **Patch** /api/v1/milestones/{milestoneId} | 
[**RestoreMilestone**](MilestonesAPI.md#RestoreMilestone) | **Post** /api/v1/milestones/{milestoneId}/restore | 
[**SearchMilestones**](MilestonesAPI.md#SearchMilestones) | **Post** /api/v1/milestones/search | 



## CompareMilestones

> MilestoneComparisonDto CompareMilestones(ctx, milestone1Id, milestone2Id).Execute()





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
	milestone1Id := "7e39bfc3-32b2-46a7-b9d0-95f6549cd85e" // string | 
	milestone2Id := "7e39bfc3-32b2-46a7-b9d0-95f6549cd85e" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MilestonesAPI.CompareMilestones(context.Background(), milestone1Id, milestone2Id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MilestonesAPI.CompareMilestones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompareMilestones`: MilestoneComparisonDto
	fmt.Fprintf(os.Stdout, "Response from `MilestonesAPI.CompareMilestones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**milestone1Id** | **string** |  | 
**milestone2Id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompareMilestonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MilestoneComparisonDto**](MilestoneComparisonDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMilestone

> MilestoneMetadataDto CreateMilestone(ctx).CreateMilestoneDto(createMilestoneDto).Execute()



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
	createMilestoneDto := *openapiclient.NewCreateMilestoneDto("Name_example", "FileId_example") // CreateMilestoneDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MilestonesAPI.CreateMilestone(context.Background()).CreateMilestoneDto(createMilestoneDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MilestonesAPI.CreateMilestone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMilestone`: MilestoneMetadataDto
	fmt.Fprintf(os.Stdout, "Response from `MilestonesAPI.CreateMilestone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMilestoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMilestoneDto** | [**CreateMilestoneDto**](CreateMilestoneDto.md) |  | 

### Return type

[**MilestoneMetadataDto**](MilestoneMetadataDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMilestone

> DeleteMilestone(ctx, milestoneId).Execute()





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
	milestoneId := "7e39bfc3-32b2-46a7-b9d0-95f6549cd85e" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MilestonesAPI.DeleteMilestone(context.Background(), milestoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MilestonesAPI.DeleteMilestone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**milestoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMilestoneRequest struct via the builder pattern


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


## GetMilestone

> MilestoneDto GetMilestone(ctx, milestoneId).Execute()



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
	milestoneId := "7e39bfc3-32b2-46a7-b9d0-95f6549cd85e" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MilestonesAPI.GetMilestone(context.Background(), milestoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MilestonesAPI.GetMilestone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMilestone`: MilestoneDto
	fmt.Fprintf(os.Stdout, "Response from `MilestonesAPI.GetMilestone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**milestoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMilestoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MilestoneDto**](MilestoneDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMilestone

> MilestoneMetadataDto PatchMilestone(ctx, milestoneId).UpdateMilestoneDto(updateMilestoneDto).Execute()



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
	milestoneId := "7e39bfc3-32b2-46a7-b9d0-95f6549cd85e" // string | 
	updateMilestoneDto := *openapiclient.NewUpdateMilestoneDto("Name_example") // UpdateMilestoneDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MilestonesAPI.PatchMilestone(context.Background(), milestoneId).UpdateMilestoneDto(updateMilestoneDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MilestonesAPI.PatchMilestone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMilestone`: MilestoneMetadataDto
	fmt.Fprintf(os.Stdout, "Response from `MilestonesAPI.PatchMilestone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**milestoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMilestoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMilestoneDto** | [**UpdateMilestoneDto**](UpdateMilestoneDto.md) |  | 

### Return type

[**MilestoneMetadataDto**](MilestoneMetadataDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreMilestone

> MilestoneMetadataDto RestoreMilestone(ctx, milestoneId).Execute()



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
	milestoneId := "7e39bfc3-32b2-46a7-b9d0-95f6549cd85e" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MilestonesAPI.RestoreMilestone(context.Background(), milestoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MilestonesAPI.RestoreMilestone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreMilestone`: MilestoneMetadataDto
	fmt.Fprintf(os.Stdout, "Response from `MilestonesAPI.RestoreMilestone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**milestoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreMilestoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MilestoneMetadataDto**](MilestoneMetadataDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchMilestones

> PubSearchResultDtoMilestoneMetadataDto SearchMilestones(ctx).PubSearchDtoMilestoneMetadataDto(pubSearchDtoMilestoneMetadataDto).Execute()





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
	pubSearchDtoMilestoneMetadataDto := *openapiclient.NewPubSearchDtoMilestoneMetadataDto() // PubSearchDtoMilestoneMetadataDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MilestonesAPI.SearchMilestones(context.Background()).PubSearchDtoMilestoneMetadataDto(pubSearchDtoMilestoneMetadataDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MilestonesAPI.SearchMilestones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchMilestones`: PubSearchResultDtoMilestoneMetadataDto
	fmt.Fprintf(os.Stdout, "Response from `MilestonesAPI.SearchMilestones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchMilestonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pubSearchDtoMilestoneMetadataDto** | [**PubSearchDtoMilestoneMetadataDto**](PubSearchDtoMilestoneMetadataDto.md) |  | 

### Return type

[**PubSearchResultDtoMilestoneMetadataDto**](PubSearchResultDtoMilestoneMetadataDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

