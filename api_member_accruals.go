/*
Connectwise Manage Public Endpoints

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2025.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cwapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// MemberAccrualsAPIService MemberAccrualsAPI service
type MemberAccrualsAPIService service

type ApiDeleteSystemMembersByParentIdAccrualsByIdRequest struct {
	ctx context.Context
	ApiService *MemberAccrualsAPIService
	id int32
	parentId int32
	clientId *string
}

// 
func (r ApiDeleteSystemMembersByParentIdAccrualsByIdRequest) ClientId(clientId string) ApiDeleteSystemMembersByParentIdAccrualsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiDeleteSystemMembersByParentIdAccrualsByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSystemMembersByParentIdAccrualsByIdExecute(r)
}

/*
DeleteSystemMembersByParentIdAccrualsById Delete MemberAccrual

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id accrualId
 @param parentId memberId
 @return ApiDeleteSystemMembersByParentIdAccrualsByIdRequest
*/
func (a *MemberAccrualsAPIService) DeleteSystemMembersByParentIdAccrualsById(ctx context.Context, id int32, parentId int32) ApiDeleteSystemMembersByParentIdAccrualsByIdRequest {
	return ApiDeleteSystemMembersByParentIdAccrualsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
func (a *MemberAccrualsAPIService) DeleteSystemMembersByParentIdAccrualsByIdExecute(r ApiDeleteSystemMembersByParentIdAccrualsByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MemberAccrualsAPIService.DeleteSystemMembersByParentIdAccrualsById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/members/{parentId}/accruals/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.clientId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "clientId", r.clientId, "simple", "")
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

type ApiGetSystemMembersByParentIdAccrualsRequest struct {
	ctx context.Context
	ApiService *MemberAccrualsAPIService
	parentId int32
	conditions *string
	childConditions *string
	customFieldConditions *string
	orderBy *string
	fields *string
	page *int32
	pageSize *int32
	pageId *int32
	clientId *string
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsRequest) Conditions(conditions string) ApiGetSystemMembersByParentIdAccrualsRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsRequest) ChildConditions(childConditions string) ApiGetSystemMembersByParentIdAccrualsRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemMembersByParentIdAccrualsRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsRequest) OrderBy(orderBy string) ApiGetSystemMembersByParentIdAccrualsRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsRequest) Fields(fields string) ApiGetSystemMembersByParentIdAccrualsRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsRequest) Page(page int32) ApiGetSystemMembersByParentIdAccrualsRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsRequest) PageSize(pageSize int32) ApiGetSystemMembersByParentIdAccrualsRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsRequest) PageId(pageId int32) ApiGetSystemMembersByParentIdAccrualsRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsRequest) ClientId(clientId string) ApiGetSystemMembersByParentIdAccrualsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemMembersByParentIdAccrualsRequest) Execute() ([]MemberAccrual, *http.Response, error) {
	return r.ApiService.GetSystemMembersByParentIdAccrualsExecute(r)
}

/*
GetSystemMembersByParentIdAccruals Get List of MemberAccrual

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId memberId
 @return ApiGetSystemMembersByParentIdAccrualsRequest
*/
func (a *MemberAccrualsAPIService) GetSystemMembersByParentIdAccruals(ctx context.Context, parentId int32) ApiGetSystemMembersByParentIdAccrualsRequest {
	return ApiGetSystemMembersByParentIdAccrualsRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return []MemberAccrual
func (a *MemberAccrualsAPIService) GetSystemMembersByParentIdAccrualsExecute(r ApiGetSystemMembersByParentIdAccrualsRequest) ([]MemberAccrual, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []MemberAccrual
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MemberAccrualsAPIService.GetSystemMembersByParentIdAccruals")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/members/{parentId}/accruals"
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.conditions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conditions", r.conditions, "form", "")
	}
	if r.childConditions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "childConditions", r.childConditions, "form", "")
	}
	if r.customFieldConditions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "customFieldConditions", r.customFieldConditions, "form", "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderBy", r.orderBy, "form", "")
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.pageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageId", r.pageId, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.connectwise.com+json; version=2025.1"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.clientId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "clientId", r.clientId, "simple", "")
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

type ApiGetSystemMembersByParentIdAccrualsByIdRequest struct {
	ctx context.Context
	ApiService *MemberAccrualsAPIService
	id int32
	parentId int32
	conditions *string
	childConditions *string
	customFieldConditions *string
	orderBy *string
	fields *string
	page *int32
	pageSize *int32
	pageId *int32
	clientId *string
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsByIdRequest) Conditions(conditions string) ApiGetSystemMembersByParentIdAccrualsByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsByIdRequest) ChildConditions(childConditions string) ApiGetSystemMembersByParentIdAccrualsByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemMembersByParentIdAccrualsByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsByIdRequest) OrderBy(orderBy string) ApiGetSystemMembersByParentIdAccrualsByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsByIdRequest) Fields(fields string) ApiGetSystemMembersByParentIdAccrualsByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsByIdRequest) Page(page int32) ApiGetSystemMembersByParentIdAccrualsByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsByIdRequest) PageSize(pageSize int32) ApiGetSystemMembersByParentIdAccrualsByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsByIdRequest) PageId(pageId int32) ApiGetSystemMembersByParentIdAccrualsByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsByIdRequest) ClientId(clientId string) ApiGetSystemMembersByParentIdAccrualsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemMembersByParentIdAccrualsByIdRequest) Execute() (*MemberAccrual, *http.Response, error) {
	return r.ApiService.GetSystemMembersByParentIdAccrualsByIdExecute(r)
}

/*
GetSystemMembersByParentIdAccrualsById Get MemberAccrual

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id accrualId
 @param parentId memberId
 @return ApiGetSystemMembersByParentIdAccrualsByIdRequest
*/
func (a *MemberAccrualsAPIService) GetSystemMembersByParentIdAccrualsById(ctx context.Context, id int32, parentId int32) ApiGetSystemMembersByParentIdAccrualsByIdRequest {
	return ApiGetSystemMembersByParentIdAccrualsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return MemberAccrual
func (a *MemberAccrualsAPIService) GetSystemMembersByParentIdAccrualsByIdExecute(r ApiGetSystemMembersByParentIdAccrualsByIdRequest) (*MemberAccrual, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MemberAccrual
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MemberAccrualsAPIService.GetSystemMembersByParentIdAccrualsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/members/{parentId}/accruals/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.conditions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conditions", r.conditions, "form", "")
	}
	if r.childConditions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "childConditions", r.childConditions, "form", "")
	}
	if r.customFieldConditions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "customFieldConditions", r.customFieldConditions, "form", "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderBy", r.orderBy, "form", "")
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.pageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageId", r.pageId, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.connectwise.com+json; version=2025.1"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.clientId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "clientId", r.clientId, "simple", "")
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

type ApiGetSystemMembersByParentIdAccrualsCountRequest struct {
	ctx context.Context
	ApiService *MemberAccrualsAPIService
	parentId int32
	conditions *string
	childConditions *string
	customFieldConditions *string
	orderBy *string
	fields *string
	page *int32
	pageSize *int32
	pageId *int32
	clientId *string
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsCountRequest) Conditions(conditions string) ApiGetSystemMembersByParentIdAccrualsCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsCountRequest) ChildConditions(childConditions string) ApiGetSystemMembersByParentIdAccrualsCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetSystemMembersByParentIdAccrualsCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsCountRequest) OrderBy(orderBy string) ApiGetSystemMembersByParentIdAccrualsCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsCountRequest) Fields(fields string) ApiGetSystemMembersByParentIdAccrualsCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsCountRequest) Page(page int32) ApiGetSystemMembersByParentIdAccrualsCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsCountRequest) PageSize(pageSize int32) ApiGetSystemMembersByParentIdAccrualsCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsCountRequest) PageId(pageId int32) ApiGetSystemMembersByParentIdAccrualsCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSystemMembersByParentIdAccrualsCountRequest) ClientId(clientId string) ApiGetSystemMembersByParentIdAccrualsCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSystemMembersByParentIdAccrualsCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetSystemMembersByParentIdAccrualsCountExecute(r)
}

/*
GetSystemMembersByParentIdAccrualsCount Get Count of MemberAccrual

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId memberId
 @return ApiGetSystemMembersByParentIdAccrualsCountRequest
*/
func (a *MemberAccrualsAPIService) GetSystemMembersByParentIdAccrualsCount(ctx context.Context, parentId int32) ApiGetSystemMembersByParentIdAccrualsCountRequest {
	return ApiGetSystemMembersByParentIdAccrualsCountRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return Count
func (a *MemberAccrualsAPIService) GetSystemMembersByParentIdAccrualsCountExecute(r ApiGetSystemMembersByParentIdAccrualsCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MemberAccrualsAPIService.GetSystemMembersByParentIdAccrualsCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/members/{parentId}/accruals/count"
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.conditions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conditions", r.conditions, "form", "")
	}
	if r.childConditions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "childConditions", r.childConditions, "form", "")
	}
	if r.customFieldConditions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "customFieldConditions", r.customFieldConditions, "form", "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderBy", r.orderBy, "form", "")
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.pageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageId", r.pageId, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.connectwise.com+json; version=2025.1"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.clientId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "clientId", r.clientId, "simple", "")
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

type ApiPatchSystemMembersByParentIdAccrualsByIdRequest struct {
	ctx context.Context
	ApiService *MemberAccrualsAPIService
	id int32
	parentId int32
	patchOperation *[]PatchOperation
	clientId *string
}

// List of PatchOperation
func (r ApiPatchSystemMembersByParentIdAccrualsByIdRequest) PatchOperation(patchOperation []PatchOperation) ApiPatchSystemMembersByParentIdAccrualsByIdRequest {
	r.patchOperation = &patchOperation
	return r
}

// 
func (r ApiPatchSystemMembersByParentIdAccrualsByIdRequest) ClientId(clientId string) ApiPatchSystemMembersByParentIdAccrualsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPatchSystemMembersByParentIdAccrualsByIdRequest) Execute() (*MemberAccrual, *http.Response, error) {
	return r.ApiService.PatchSystemMembersByParentIdAccrualsByIdExecute(r)
}

/*
PatchSystemMembersByParentIdAccrualsById Patch MemberAccrual

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id accrualId
 @param parentId memberId
 @return ApiPatchSystemMembersByParentIdAccrualsByIdRequest
*/
func (a *MemberAccrualsAPIService) PatchSystemMembersByParentIdAccrualsById(ctx context.Context, id int32, parentId int32) ApiPatchSystemMembersByParentIdAccrualsByIdRequest {
	return ApiPatchSystemMembersByParentIdAccrualsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return MemberAccrual
func (a *MemberAccrualsAPIService) PatchSystemMembersByParentIdAccrualsByIdExecute(r ApiPatchSystemMembersByParentIdAccrualsByIdRequest) (*MemberAccrual, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MemberAccrual
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MemberAccrualsAPIService.PatchSystemMembersByParentIdAccrualsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/members/{parentId}/accruals/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchOperation == nil {
		return localVarReturnValue, nil, reportError("patchOperation is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.connectwise.com+json; version=2025.1"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.clientId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "clientId", r.clientId, "simple", "")
	}
	// body params
	localVarPostBody = r.patchOperation
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

type ApiPostSystemMembersByParentIdAccrualsRequest struct {
	ctx context.Context
	ApiService *MemberAccrualsAPIService
	parentId int32
	memberAccrual *MemberAccrual
	clientId *string
}

// memberAccrual
func (r ApiPostSystemMembersByParentIdAccrualsRequest) MemberAccrual(memberAccrual MemberAccrual) ApiPostSystemMembersByParentIdAccrualsRequest {
	r.memberAccrual = &memberAccrual
	return r
}

// 
func (r ApiPostSystemMembersByParentIdAccrualsRequest) ClientId(clientId string) ApiPostSystemMembersByParentIdAccrualsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostSystemMembersByParentIdAccrualsRequest) Execute() (*MemberAccrual, *http.Response, error) {
	return r.ApiService.PostSystemMembersByParentIdAccrualsExecute(r)
}

/*
PostSystemMembersByParentIdAccruals Post MemberAccrual

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId memberId
 @return ApiPostSystemMembersByParentIdAccrualsRequest
*/
func (a *MemberAccrualsAPIService) PostSystemMembersByParentIdAccruals(ctx context.Context, parentId int32) ApiPostSystemMembersByParentIdAccrualsRequest {
	return ApiPostSystemMembersByParentIdAccrualsRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return MemberAccrual
func (a *MemberAccrualsAPIService) PostSystemMembersByParentIdAccrualsExecute(r ApiPostSystemMembersByParentIdAccrualsRequest) (*MemberAccrual, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MemberAccrual
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MemberAccrualsAPIService.PostSystemMembersByParentIdAccruals")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/members/{parentId}/accruals"
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.memberAccrual == nil {
		return localVarReturnValue, nil, reportError("memberAccrual is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.connectwise.com+json; version=2025.1"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.clientId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "clientId", r.clientId, "simple", "")
	}
	// body params
	localVarPostBody = r.memberAccrual
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

type ApiPutSystemMembersByParentIdAccrualsByIdRequest struct {
	ctx context.Context
	ApiService *MemberAccrualsAPIService
	id int32
	parentId int32
	memberAccrual *MemberAccrual
	clientId *string
}

// memberAccrual
func (r ApiPutSystemMembersByParentIdAccrualsByIdRequest) MemberAccrual(memberAccrual MemberAccrual) ApiPutSystemMembersByParentIdAccrualsByIdRequest {
	r.memberAccrual = &memberAccrual
	return r
}

// 
func (r ApiPutSystemMembersByParentIdAccrualsByIdRequest) ClientId(clientId string) ApiPutSystemMembersByParentIdAccrualsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPutSystemMembersByParentIdAccrualsByIdRequest) Execute() (*MemberAccrual, *http.Response, error) {
	return r.ApiService.PutSystemMembersByParentIdAccrualsByIdExecute(r)
}

/*
PutSystemMembersByParentIdAccrualsById Put MemberAccrual

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id accrualId
 @param parentId memberId
 @return ApiPutSystemMembersByParentIdAccrualsByIdRequest
*/
func (a *MemberAccrualsAPIService) PutSystemMembersByParentIdAccrualsById(ctx context.Context, id int32, parentId int32) ApiPutSystemMembersByParentIdAccrualsByIdRequest {
	return ApiPutSystemMembersByParentIdAccrualsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return MemberAccrual
func (a *MemberAccrualsAPIService) PutSystemMembersByParentIdAccrualsByIdExecute(r ApiPutSystemMembersByParentIdAccrualsByIdRequest) (*MemberAccrual, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MemberAccrual
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MemberAccrualsAPIService.PutSystemMembersByParentIdAccrualsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/members/{parentId}/accruals/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.memberAccrual == nil {
		return localVarReturnValue, nil, reportError("memberAccrual is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.connectwise.com+json; version=2025.1"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.clientId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "clientId", r.clientId, "simple", "")
	}
	// body params
	localVarPostBody = r.memberAccrual
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
