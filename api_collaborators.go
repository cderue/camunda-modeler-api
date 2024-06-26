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


// CollaboratorsAPIService CollaboratorsAPI service
type CollaboratorsAPIService service

type ApiDeleteCollaboratorRequest struct {
	ctx context.Context
	ApiService *CollaboratorsAPIService
	projectId string
	email string
}

func (r ApiDeleteCollaboratorRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCollaboratorExecute(r)
}

/*
DeleteCollaborator Method for DeleteCollaborator

Deletion of resources is recursive and cannot be undone.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId
 @param email
 @return ApiDeleteCollaboratorRequest
*/
func (a *CollaboratorsAPIService) DeleteCollaborator(ctx context.Context, projectId string, email string) ApiDeleteCollaboratorRequest {
	return ApiDeleteCollaboratorRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		email: email,
	}
}

// Execute executes the request
func (a *CollaboratorsAPIService) DeleteCollaboratorExecute(r ApiDeleteCollaboratorRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CollaboratorsAPIService.DeleteCollaborator")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/projects/{projectId}/collaborators/{email}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"email"+"}", url.PathEscape(parameterValueToString(r.email, "email")), -1)

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

type ApiModifyCollaboratorRequest struct {
	ctx context.Context
	ApiService *CollaboratorsAPIService
	createCollaboratorDto *CreateCollaboratorDto
}

func (r ApiModifyCollaboratorRequest) CreateCollaboratorDto(createCollaboratorDto CreateCollaboratorDto) ApiModifyCollaboratorRequest {
	r.createCollaboratorDto = &createCollaboratorDto
	return r
}

func (r ApiModifyCollaboratorRequest) Execute() (*http.Response, error) {
	return r.ApiService.ModifyCollaboratorExecute(r)
}

/*
ModifyCollaborator Method for ModifyCollaborator

Adds a new collaborator to a project or modifies the permission level of an existing collaborator.<br/><br/>
<strong>Note:</strong> Only users that are part of the authorized organization
(see <code>GET /api/v1/info</code>) <i>and</i> logged in to Web Modeler at least once can be added to a
project.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiModifyCollaboratorRequest
*/
func (a *CollaboratorsAPIService) ModifyCollaborator(ctx context.Context) ApiModifyCollaboratorRequest {
	return ApiModifyCollaboratorRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *CollaboratorsAPIService) ModifyCollaboratorExecute(r ApiModifyCollaboratorRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CollaboratorsAPIService.ModifyCollaborator")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/collaborators"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createCollaboratorDto == nil {
		return nil, reportError("createCollaboratorDto is required and must be specified")
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
	localVarPostBody = r.createCollaboratorDto
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

type ApiSearchCollaboratorsRequest struct {
	ctx context.Context
	ApiService *CollaboratorsAPIService
	pubSearchDtoProjectCollaboratorDto *PubSearchDtoProjectCollaboratorDto
}

func (r ApiSearchCollaboratorsRequest) PubSearchDtoProjectCollaboratorDto(pubSearchDtoProjectCollaboratorDto PubSearchDtoProjectCollaboratorDto) ApiSearchCollaboratorsRequest {
	r.pubSearchDtoProjectCollaboratorDto = &pubSearchDtoProjectCollaboratorDto
	return r
}

func (r ApiSearchCollaboratorsRequest) Execute() (*PubSearchResultDtoProjectCollaboratorDto, *http.Response, error) {
	return r.ApiService.SearchCollaboratorsExecute(r)
}

/*
SearchCollaborators Method for SearchCollaborators

Searches for collaborators.<br/>
<ul>
  <li><p><em>filter</em> specifies which fields should match. Only items that match the given fields will be
   returned.</p></li>
  <li><p><em>sort</em> specifies by which fields and direction (<code>ASC</code>/<code>DESC</code>) the result
   should be sorted.</p></li>
  <p>
    <strong>Note:</strong> Sorting by <code>role</code> is not supported.
  </p>
  <li><p><em>page</em> specifies the page number to return.</p></li>
  <li><p><em>size</em> specifies the number of items per page. The default value is 10.</p></li>
</ul>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchCollaboratorsRequest
*/
func (a *CollaboratorsAPIService) SearchCollaborators(ctx context.Context) ApiSearchCollaboratorsRequest {
	return ApiSearchCollaboratorsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PubSearchResultDtoProjectCollaboratorDto
func (a *CollaboratorsAPIService) SearchCollaboratorsExecute(r ApiSearchCollaboratorsRequest) (*PubSearchResultDtoProjectCollaboratorDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PubSearchResultDtoProjectCollaboratorDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CollaboratorsAPIService.SearchCollaborators")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/collaborators/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pubSearchDtoProjectCollaboratorDto == nil {
		return localVarReturnValue, nil, reportError("pubSearchDtoProjectCollaboratorDto is required and must be specified")
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
	localVarPostBody = r.pubSearchDtoProjectCollaboratorDto
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
