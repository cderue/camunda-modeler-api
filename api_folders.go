/*
Web Modeler REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 41f4f56
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// FoldersAPIService FoldersAPI service
type FoldersAPIService service

type ApiCreateFolderRequest struct {
	ctx context.Context
	ApiService *FoldersAPIService
	createFolderDto *CreateFolderDto
}

func (r ApiCreateFolderRequest) CreateFolderDto(createFolderDto CreateFolderDto) ApiCreateFolderRequest {
	r.createFolderDto = &createFolderDto
	return r
}

func (r ApiCreateFolderRequest) Execute() (*FolderMetadataDto, *http.Response, error) {
	return r.ApiService.CreateFolderExecute(r)
}

/*
CreateFolder Method for CreateFolder

Creates a new folder.<br/>
<ul>
  <li><p>When only <em>parentId</em> is given, the folder will be created in that
      folder. The folder can be in any project of the same organization.</p></li>
  <li><p>When <em>projectId</em> is given and <em>parentId</em> is either
      null or omitted altogether, the folder will be created in the root of the project.</p></li>
  <li><p>When <em>projectId</em> and <em>parentId</em> are both given,
      they must be consistent - i.e. the parent folder is in the project.</p></li>
</ul>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateFolderRequest
*/
func (a *FoldersAPIService) CreateFolder(ctx context.Context) ApiCreateFolderRequest {
	return ApiCreateFolderRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FolderMetadataDto
func (a *FoldersAPIService) CreateFolderExecute(r ApiCreateFolderRequest) (*FolderMetadataDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FolderMetadataDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FoldersAPIService.CreateFolder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/folders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createFolderDto == nil {
		return localVarReturnValue, nil, reportError("createFolderDto is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createFolderDto
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteFolderRequest struct {
	ctx context.Context
	ApiService *FoldersAPIService
	folderId string
}

func (r ApiDeleteFolderRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteFolderExecute(r)
}

/*
DeleteFolder Method for DeleteFolder

This endpoint deletes an empty folder. A folder is considered empty if there are no files in it.
 Deletion of resources is recursive and cannot be undone.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param folderId
 @return ApiDeleteFolderRequest
*/
func (a *FoldersAPIService) DeleteFolder(ctx context.Context, folderId string) ApiDeleteFolderRequest {
	return ApiDeleteFolderRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
	}
}

// Execute executes the request
func (a *FoldersAPIService) DeleteFolderExecute(r ApiDeleteFolderRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FoldersAPIService.DeleteFolder")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/folders/{folderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetFolderRequest struct {
	ctx context.Context
	ApiService *FoldersAPIService
	folderId string
}

func (r ApiGetFolderRequest) Execute() (*FolderDto, *http.Response, error) {
	return r.ApiService.GetFolderExecute(r)
}

/*
GetFolder Method for GetFolder

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param folderId
 @return ApiGetFolderRequest
*/
func (a *FoldersAPIService) GetFolder(ctx context.Context, folderId string) ApiGetFolderRequest {
	return ApiGetFolderRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
	}
}

// Execute executes the request
//  @return FolderDto
func (a *FoldersAPIService) GetFolderExecute(r ApiGetFolderRequest) (*FolderDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FolderDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FoldersAPIService.GetFolder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/folders/{folderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPatchFolderRequest struct {
	ctx context.Context
	ApiService *FoldersAPIService
	folderId string
	updateFolderDto *UpdateFolderDto
}

func (r ApiPatchFolderRequest) UpdateFolderDto(updateFolderDto UpdateFolderDto) ApiPatchFolderRequest {
	r.updateFolderDto = &updateFolderDto
	return r
}

func (r ApiPatchFolderRequest) Execute() (*FolderMetadataDto, *http.Response, error) {
	return r.ApiService.PatchFolderExecute(r)
}

/*
PatchFolder Method for PatchFolder

This endpoint updates the name or location of a folder, or both at the same time.<br/>
<br/>
To move a folder, specify <em>projectId</em> and/or <em>parentId</em>:
<ul>
  <li><p>When only <em>parentId</em> is given, the file will be moved to that
      folder. The folder must keep in the same organization.</p></li>
  <li><p>When <em>projectId</em> is given and <em>parentId</em> is either
      null or omitted altogether, the file will be moved to the root of the project.</p></li>
  <li><p>When <em>projectId</em> and <em>parentId</em> are both given,
      they must be consistent - i.e. the new parent folder is in the new project.</p></li>
</ul>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param folderId
 @return ApiPatchFolderRequest
*/
func (a *FoldersAPIService) PatchFolder(ctx context.Context, folderId string) ApiPatchFolderRequest {
	return ApiPatchFolderRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
	}
}

// Execute executes the request
//  @return FolderMetadataDto
func (a *FoldersAPIService) PatchFolderExecute(r ApiPatchFolderRequest) (*FolderMetadataDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FolderMetadataDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FoldersAPIService.PatchFolder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/folders/{folderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateFolderDto == nil {
		return localVarReturnValue, nil, reportError("updateFolderDto is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateFolderDto
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}