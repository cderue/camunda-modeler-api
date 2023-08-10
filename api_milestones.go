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


// MilestonesAPIService MilestonesAPI service
type MilestonesAPIService service

type ApiCompareMilestonesRequest struct {
	ctx context.Context
	ApiService *MilestonesAPIService
	milestone1Id string
	milestone2Id string
}

func (r ApiCompareMilestonesRequest) Execute() (*MilestoneComparisonDto, *http.Response, error) {
	return r.ApiService.CompareMilestonesExecute(r)
}

/*
CompareMilestones Method for CompareMilestones

Returns a link to a visual comparison between two milestones where the milestone referenced by
<code>milestone1Id</code> acts as a baseline to compare the milestone referenced by <code>milestone2Id</code>
against.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param milestone1Id
 @param milestone2Id
 @return ApiCompareMilestonesRequest
*/
func (a *MilestonesAPIService) CompareMilestones(ctx context.Context, milestone1Id string, milestone2Id string) ApiCompareMilestonesRequest {
	return ApiCompareMilestonesRequest{
		ApiService: a,
		ctx: ctx,
		milestone1Id: milestone1Id,
		milestone2Id: milestone2Id,
	}
}

// Execute executes the request
//  @return MilestoneComparisonDto
func (a *MilestonesAPIService) CompareMilestonesExecute(r ApiCompareMilestonesRequest) (*MilestoneComparisonDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MilestoneComparisonDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MilestonesAPIService.CompareMilestones")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/milestones/compare/{milestone1Id}...{milestone2Id}"
	localVarPath = strings.Replace(localVarPath, "{"+"milestone1Id"+"}", url.PathEscape(parameterValueToString(r.milestone1Id, "milestone1Id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"milestone2Id"+"}", url.PathEscape(parameterValueToString(r.milestone2Id, "milestone2Id")), -1)

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

type ApiCreateMilestoneRequest struct {
	ctx context.Context
	ApiService *MilestonesAPIService
	createMilestoneDto *CreateMilestoneDto
}

func (r ApiCreateMilestoneRequest) CreateMilestoneDto(createMilestoneDto CreateMilestoneDto) ApiCreateMilestoneRequest {
	r.createMilestoneDto = &createMilestoneDto
	return r
}

func (r ApiCreateMilestoneRequest) Execute() (*MilestoneMetadataDto, *http.Response, error) {
	return r.ApiService.CreateMilestoneExecute(r)
}

/*
CreateMilestone Method for CreateMilestone

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateMilestoneRequest
*/
func (a *MilestonesAPIService) CreateMilestone(ctx context.Context) ApiCreateMilestoneRequest {
	return ApiCreateMilestoneRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MilestoneMetadataDto
func (a *MilestonesAPIService) CreateMilestoneExecute(r ApiCreateMilestoneRequest) (*MilestoneMetadataDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MilestoneMetadataDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MilestonesAPIService.CreateMilestone")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/milestones"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createMilestoneDto == nil {
		return localVarReturnValue, nil, reportError("createMilestoneDto is required and must be specified")
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
	localVarPostBody = r.createMilestoneDto
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

type ApiDeleteMilestoneRequest struct {
	ctx context.Context
	ApiService *MilestonesAPIService
	milestoneId string
}

func (r ApiDeleteMilestoneRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteMilestoneExecute(r)
}

/*
DeleteMilestone Method for DeleteMilestone

Deletion of resources is recursive and cannot be undone.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param milestoneId
 @return ApiDeleteMilestoneRequest
*/
func (a *MilestonesAPIService) DeleteMilestone(ctx context.Context, milestoneId string) ApiDeleteMilestoneRequest {
	return ApiDeleteMilestoneRequest{
		ApiService: a,
		ctx: ctx,
		milestoneId: milestoneId,
	}
}

// Execute executes the request
func (a *MilestonesAPIService) DeleteMilestoneExecute(r ApiDeleteMilestoneRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MilestonesAPIService.DeleteMilestone")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/milestones/{milestoneId}"
	localVarPath = strings.Replace(localVarPath, "{"+"milestoneId"+"}", url.PathEscape(parameterValueToString(r.milestoneId, "milestoneId")), -1)

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

type ApiGetMilestoneRequest struct {
	ctx context.Context
	ApiService *MilestonesAPIService
	milestoneId string
}

func (r ApiGetMilestoneRequest) Execute() (*MilestoneDto, *http.Response, error) {
	return r.ApiService.GetMilestoneExecute(r)
}

/*
GetMilestone Method for GetMilestone

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param milestoneId
 @return ApiGetMilestoneRequest
*/
func (a *MilestonesAPIService) GetMilestone(ctx context.Context, milestoneId string) ApiGetMilestoneRequest {
	return ApiGetMilestoneRequest{
		ApiService: a,
		ctx: ctx,
		milestoneId: milestoneId,
	}
}

// Execute executes the request
//  @return MilestoneDto
func (a *MilestonesAPIService) GetMilestoneExecute(r ApiGetMilestoneRequest) (*MilestoneDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MilestoneDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MilestonesAPIService.GetMilestone")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/milestones/{milestoneId}"
	localVarPath = strings.Replace(localVarPath, "{"+"milestoneId"+"}", url.PathEscape(parameterValueToString(r.milestoneId, "milestoneId")), -1)

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

type ApiPatchMilestoneRequest struct {
	ctx context.Context
	ApiService *MilestonesAPIService
	milestoneId string
	updateMilestoneDto *UpdateMilestoneDto
}

func (r ApiPatchMilestoneRequest) UpdateMilestoneDto(updateMilestoneDto UpdateMilestoneDto) ApiPatchMilestoneRequest {
	r.updateMilestoneDto = &updateMilestoneDto
	return r
}

func (r ApiPatchMilestoneRequest) Execute() (*MilestoneMetadataDto, *http.Response, error) {
	return r.ApiService.PatchMilestoneExecute(r)
}

/*
PatchMilestone Method for PatchMilestone

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param milestoneId
 @return ApiPatchMilestoneRequest
*/
func (a *MilestonesAPIService) PatchMilestone(ctx context.Context, milestoneId string) ApiPatchMilestoneRequest {
	return ApiPatchMilestoneRequest{
		ApiService: a,
		ctx: ctx,
		milestoneId: milestoneId,
	}
}

// Execute executes the request
//  @return MilestoneMetadataDto
func (a *MilestonesAPIService) PatchMilestoneExecute(r ApiPatchMilestoneRequest) (*MilestoneMetadataDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MilestoneMetadataDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MilestonesAPIService.PatchMilestone")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/milestones/{milestoneId}"
	localVarPath = strings.Replace(localVarPath, "{"+"milestoneId"+"}", url.PathEscape(parameterValueToString(r.milestoneId, "milestoneId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateMilestoneDto == nil {
		return localVarReturnValue, nil, reportError("updateMilestoneDto is required and must be specified")
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
	localVarPostBody = r.updateMilestoneDto
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

type ApiRestoreMilestoneRequest struct {
	ctx context.Context
	ApiService *MilestonesAPIService
	milestoneId string
}

func (r ApiRestoreMilestoneRequest) Execute() (*MilestoneMetadataDto, *http.Response, error) {
	return r.ApiService.RestoreMilestoneExecute(r)
}

/*
RestoreMilestone Method for RestoreMilestone

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param milestoneId
 @return ApiRestoreMilestoneRequest
*/
func (a *MilestonesAPIService) RestoreMilestone(ctx context.Context, milestoneId string) ApiRestoreMilestoneRequest {
	return ApiRestoreMilestoneRequest{
		ApiService: a,
		ctx: ctx,
		milestoneId: milestoneId,
	}
}

// Execute executes the request
//  @return MilestoneMetadataDto
func (a *MilestonesAPIService) RestoreMilestoneExecute(r ApiRestoreMilestoneRequest) (*MilestoneMetadataDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MilestoneMetadataDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MilestonesAPIService.RestoreMilestone")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/milestones/{milestoneId}/restore"
	localVarPath = strings.Replace(localVarPath, "{"+"milestoneId"+"}", url.PathEscape(parameterValueToString(r.milestoneId, "milestoneId")), -1)

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

type ApiSearchMilestonesRequest struct {
	ctx context.Context
	ApiService *MilestonesAPIService
	pubSearchDtoMilestoneMetadataDto *PubSearchDtoMilestoneMetadataDto
}

func (r ApiSearchMilestonesRequest) PubSearchDtoMilestoneMetadataDto(pubSearchDtoMilestoneMetadataDto PubSearchDtoMilestoneMetadataDto) ApiSearchMilestonesRequest {
	r.pubSearchDtoMilestoneMetadataDto = &pubSearchDtoMilestoneMetadataDto
	return r
}

func (r ApiSearchMilestonesRequest) Execute() (*PubSearchResultDtoMilestoneMetadataDto, *http.Response, error) {
	return r.ApiService.SearchMilestonesExecute(r)
}

/*
SearchMilestones Method for SearchMilestones

Searches for milestones.<br/>
<ul>
  <li><p><em>filter</em> specifies which fields should match. Only items that match the given fields will be
   returned.</p>
   <p/><strong>Note:</strong> Date fields need to be specified in a format compatible with
   <code>java.time.ZonedDateTime</code>; for example <code>2023-09-20T11:31:20.206801604Z</code>.<br/><br/>
   You can use suffixes to match date ranges:
   <table>
    <thead>
     <tr><th>Modifier</th><th>Description</th></tr>
    </thead>
    <tbody>
     <tr><td>||/y</td><td>Within the year of the provided date</td></tr>
     <tr><td>||/M</td><td>Within the month of the provided date</td></tr>
     <tr><td>||/w</td><td>Within the week of the provided date</td></tr>
     <tr><td>||/d</td><td>Within the day of the provided date</td></tr>
     <tr><td>||/h</td><td>Within the hour of the provided date</td></tr>
     <tr><td>||/m</td><td>Within the minute of the provided date</td></tr>
     <tr><td>||/s</td><td>Within the second of the provided date</td></tr>
    </tbody>
   </table>
   <br/>
   <p>You can also use the following prefixes to filter before or after the date:</p>
    <table>
    <thead>
     <tr><th>Operator</th><th>Description</th></tr>
    </thead>
    <tbody>
     <tr><td>&lt;</td><td>Before the provided date</td></tr>
     <tr><td>&lt;=</td><td>On or before the provided date</td></tr>
     <tr><td>&gt;</td><td>After the provided date</td></tr>
     <tr><td>&gt;=</td><td>On or after the provided date</td></tr>
    </tbody>
    </table>
    <p>Operator and modifier can be combined. Example:
      <code>&gt;=2023-09-20T11:31:20.206801604Z||/y</code> returns all milestones within or after 2023</p>
  </p></li>
  <li><p><em>sort</em> specifies by which fields and direction (<code>ASC</code>/<code>DESC</code>) the result
   should be sorted.</p></li>
  <li><p><em>page</em> specifies the page number to return.</p></li>
  <li><p><em>size</em> specifies the number of items per page. The default value is 10.</p></li>
</ul>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchMilestonesRequest
*/
func (a *MilestonesAPIService) SearchMilestones(ctx context.Context) ApiSearchMilestonesRequest {
	return ApiSearchMilestonesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PubSearchResultDtoMilestoneMetadataDto
func (a *MilestonesAPIService) SearchMilestonesExecute(r ApiSearchMilestonesRequest) (*PubSearchResultDtoMilestoneMetadataDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PubSearchResultDtoMilestoneMetadataDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MilestonesAPIService.SearchMilestones")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/milestones/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pubSearchDtoMilestoneMetadataDto == nil {
		return localVarReturnValue, nil, reportError("pubSearchDtoMilestoneMetadataDto is required and must be specified")
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
	localVarPostBody = r.pubSearchDtoMilestoneMetadataDto
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
