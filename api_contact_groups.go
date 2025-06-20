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


// ContactGroupsAPIService ContactGroupsAPI service
type ContactGroupsAPIService service

type ApiDeleteCompanyContactsByParentIdGroupsByIdRequest struct {
	ctx context.Context
	ApiService *ContactGroupsAPIService
	id int32
	parentId int32
	clientId *string
}

// 
func (r ApiDeleteCompanyContactsByParentIdGroupsByIdRequest) ClientId(clientId string) ApiDeleteCompanyContactsByParentIdGroupsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiDeleteCompanyContactsByParentIdGroupsByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCompanyContactsByParentIdGroupsByIdExecute(r)
}

/*
DeleteCompanyContactsByParentIdGroupsById Delete ContactGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id groupId
 @param parentId contactId
 @return ApiDeleteCompanyContactsByParentIdGroupsByIdRequest
*/
func (a *ContactGroupsAPIService) DeleteCompanyContactsByParentIdGroupsById(ctx context.Context, id int32, parentId int32) ApiDeleteCompanyContactsByParentIdGroupsByIdRequest {
	return ApiDeleteCompanyContactsByParentIdGroupsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
func (a *ContactGroupsAPIService) DeleteCompanyContactsByParentIdGroupsByIdExecute(r ApiDeleteCompanyContactsByParentIdGroupsByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContactGroupsAPIService.DeleteCompanyContactsByParentIdGroupsById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/contacts/{parentId}/groups/{id}"
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

type ApiGetCompanyContactsByParentIdGroupsRequest struct {
	ctx context.Context
	ApiService *ContactGroupsAPIService
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
func (r ApiGetCompanyContactsByParentIdGroupsRequest) Conditions(conditions string) ApiGetCompanyContactsByParentIdGroupsRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsRequest) ChildConditions(childConditions string) ApiGetCompanyContactsByParentIdGroupsRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsRequest) CustomFieldConditions(customFieldConditions string) ApiGetCompanyContactsByParentIdGroupsRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsRequest) OrderBy(orderBy string) ApiGetCompanyContactsByParentIdGroupsRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsRequest) Fields(fields string) ApiGetCompanyContactsByParentIdGroupsRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsRequest) Page(page int32) ApiGetCompanyContactsByParentIdGroupsRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsRequest) PageSize(pageSize int32) ApiGetCompanyContactsByParentIdGroupsRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsRequest) PageId(pageId int32) ApiGetCompanyContactsByParentIdGroupsRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsRequest) ClientId(clientId string) ApiGetCompanyContactsByParentIdGroupsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetCompanyContactsByParentIdGroupsRequest) Execute() ([]ContactGroup, *http.Response, error) {
	return r.ApiService.GetCompanyContactsByParentIdGroupsExecute(r)
}

/*
GetCompanyContactsByParentIdGroups Get List of ContactGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId contactId
 @return ApiGetCompanyContactsByParentIdGroupsRequest
*/
func (a *ContactGroupsAPIService) GetCompanyContactsByParentIdGroups(ctx context.Context, parentId int32) ApiGetCompanyContactsByParentIdGroupsRequest {
	return ApiGetCompanyContactsByParentIdGroupsRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return []ContactGroup
func (a *ContactGroupsAPIService) GetCompanyContactsByParentIdGroupsExecute(r ApiGetCompanyContactsByParentIdGroupsRequest) ([]ContactGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ContactGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContactGroupsAPIService.GetCompanyContactsByParentIdGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/contacts/{parentId}/groups"
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

type ApiGetCompanyContactsByParentIdGroupsByIdRequest struct {
	ctx context.Context
	ApiService *ContactGroupsAPIService
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
func (r ApiGetCompanyContactsByParentIdGroupsByIdRequest) Conditions(conditions string) ApiGetCompanyContactsByParentIdGroupsByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsByIdRequest) ChildConditions(childConditions string) ApiGetCompanyContactsByParentIdGroupsByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetCompanyContactsByParentIdGroupsByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsByIdRequest) OrderBy(orderBy string) ApiGetCompanyContactsByParentIdGroupsByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsByIdRequest) Fields(fields string) ApiGetCompanyContactsByParentIdGroupsByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsByIdRequest) Page(page int32) ApiGetCompanyContactsByParentIdGroupsByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsByIdRequest) PageSize(pageSize int32) ApiGetCompanyContactsByParentIdGroupsByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsByIdRequest) PageId(pageId int32) ApiGetCompanyContactsByParentIdGroupsByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsByIdRequest) ClientId(clientId string) ApiGetCompanyContactsByParentIdGroupsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetCompanyContactsByParentIdGroupsByIdRequest) Execute() (*ContactGroup, *http.Response, error) {
	return r.ApiService.GetCompanyContactsByParentIdGroupsByIdExecute(r)
}

/*
GetCompanyContactsByParentIdGroupsById Get ContactGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id groupId
 @param parentId contactId
 @return ApiGetCompanyContactsByParentIdGroupsByIdRequest
*/
func (a *ContactGroupsAPIService) GetCompanyContactsByParentIdGroupsById(ctx context.Context, id int32, parentId int32) ApiGetCompanyContactsByParentIdGroupsByIdRequest {
	return ApiGetCompanyContactsByParentIdGroupsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return ContactGroup
func (a *ContactGroupsAPIService) GetCompanyContactsByParentIdGroupsByIdExecute(r ApiGetCompanyContactsByParentIdGroupsByIdRequest) (*ContactGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContactGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContactGroupsAPIService.GetCompanyContactsByParentIdGroupsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/contacts/{parentId}/groups/{id}"
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

type ApiGetCompanyContactsByParentIdGroupsCountRequest struct {
	ctx context.Context
	ApiService *ContactGroupsAPIService
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
func (r ApiGetCompanyContactsByParentIdGroupsCountRequest) Conditions(conditions string) ApiGetCompanyContactsByParentIdGroupsCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsCountRequest) ChildConditions(childConditions string) ApiGetCompanyContactsByParentIdGroupsCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetCompanyContactsByParentIdGroupsCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsCountRequest) OrderBy(orderBy string) ApiGetCompanyContactsByParentIdGroupsCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsCountRequest) Fields(fields string) ApiGetCompanyContactsByParentIdGroupsCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsCountRequest) Page(page int32) ApiGetCompanyContactsByParentIdGroupsCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsCountRequest) PageSize(pageSize int32) ApiGetCompanyContactsByParentIdGroupsCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsCountRequest) PageId(pageId int32) ApiGetCompanyContactsByParentIdGroupsCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetCompanyContactsByParentIdGroupsCountRequest) ClientId(clientId string) ApiGetCompanyContactsByParentIdGroupsCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetCompanyContactsByParentIdGroupsCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetCompanyContactsByParentIdGroupsCountExecute(r)
}

/*
GetCompanyContactsByParentIdGroupsCount Get Count of ContactGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId contactId
 @return ApiGetCompanyContactsByParentIdGroupsCountRequest
*/
func (a *ContactGroupsAPIService) GetCompanyContactsByParentIdGroupsCount(ctx context.Context, parentId int32) ApiGetCompanyContactsByParentIdGroupsCountRequest {
	return ApiGetCompanyContactsByParentIdGroupsCountRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return Count
func (a *ContactGroupsAPIService) GetCompanyContactsByParentIdGroupsCountExecute(r ApiGetCompanyContactsByParentIdGroupsCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContactGroupsAPIService.GetCompanyContactsByParentIdGroupsCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/contacts/{parentId}/groups/count"
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

type ApiPatchCompanyContactsByParentIdGroupsByIdRequest struct {
	ctx context.Context
	ApiService *ContactGroupsAPIService
	id int32
	parentId int32
	patchOperation *[]PatchOperation
	clientId *string
}

// List of PatchOperation
func (r ApiPatchCompanyContactsByParentIdGroupsByIdRequest) PatchOperation(patchOperation []PatchOperation) ApiPatchCompanyContactsByParentIdGroupsByIdRequest {
	r.patchOperation = &patchOperation
	return r
}

// 
func (r ApiPatchCompanyContactsByParentIdGroupsByIdRequest) ClientId(clientId string) ApiPatchCompanyContactsByParentIdGroupsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPatchCompanyContactsByParentIdGroupsByIdRequest) Execute() (*ContactGroup, *http.Response, error) {
	return r.ApiService.PatchCompanyContactsByParentIdGroupsByIdExecute(r)
}

/*
PatchCompanyContactsByParentIdGroupsById Patch ContactGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id groupId
 @param parentId contactId
 @return ApiPatchCompanyContactsByParentIdGroupsByIdRequest
*/
func (a *ContactGroupsAPIService) PatchCompanyContactsByParentIdGroupsById(ctx context.Context, id int32, parentId int32) ApiPatchCompanyContactsByParentIdGroupsByIdRequest {
	return ApiPatchCompanyContactsByParentIdGroupsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return ContactGroup
func (a *ContactGroupsAPIService) PatchCompanyContactsByParentIdGroupsByIdExecute(r ApiPatchCompanyContactsByParentIdGroupsByIdRequest) (*ContactGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContactGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContactGroupsAPIService.PatchCompanyContactsByParentIdGroupsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/contacts/{parentId}/groups/{id}"
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

type ApiPostCompanyContactsByParentIdGroupsRequest struct {
	ctx context.Context
	ApiService *ContactGroupsAPIService
	parentId int32
	contactGroup *ContactGroup
	clientId *string
}

// contactGroup
func (r ApiPostCompanyContactsByParentIdGroupsRequest) ContactGroup(contactGroup ContactGroup) ApiPostCompanyContactsByParentIdGroupsRequest {
	r.contactGroup = &contactGroup
	return r
}

// 
func (r ApiPostCompanyContactsByParentIdGroupsRequest) ClientId(clientId string) ApiPostCompanyContactsByParentIdGroupsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostCompanyContactsByParentIdGroupsRequest) Execute() (*ContactGroup, *http.Response, error) {
	return r.ApiService.PostCompanyContactsByParentIdGroupsExecute(r)
}

/*
PostCompanyContactsByParentIdGroups Post ContactGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId contactId
 @return ApiPostCompanyContactsByParentIdGroupsRequest
*/
func (a *ContactGroupsAPIService) PostCompanyContactsByParentIdGroups(ctx context.Context, parentId int32) ApiPostCompanyContactsByParentIdGroupsRequest {
	return ApiPostCompanyContactsByParentIdGroupsRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return ContactGroup
func (a *ContactGroupsAPIService) PostCompanyContactsByParentIdGroupsExecute(r ApiPostCompanyContactsByParentIdGroupsRequest) (*ContactGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContactGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContactGroupsAPIService.PostCompanyContactsByParentIdGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/contacts/{parentId}/groups"
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.contactGroup == nil {
		return localVarReturnValue, nil, reportError("contactGroup is required and must be specified")
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
	localVarPostBody = r.contactGroup
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

type ApiPutCompanyContactsByParentIdGroupsByIdRequest struct {
	ctx context.Context
	ApiService *ContactGroupsAPIService
	id int32
	parentId int32
	contactGroup *ContactGroup
	clientId *string
}

// contactGroup
func (r ApiPutCompanyContactsByParentIdGroupsByIdRequest) ContactGroup(contactGroup ContactGroup) ApiPutCompanyContactsByParentIdGroupsByIdRequest {
	r.contactGroup = &contactGroup
	return r
}

// 
func (r ApiPutCompanyContactsByParentIdGroupsByIdRequest) ClientId(clientId string) ApiPutCompanyContactsByParentIdGroupsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPutCompanyContactsByParentIdGroupsByIdRequest) Execute() (*ContactGroup, *http.Response, error) {
	return r.ApiService.PutCompanyContactsByParentIdGroupsByIdExecute(r)
}

/*
PutCompanyContactsByParentIdGroupsById Put ContactGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id groupId
 @param parentId contactId
 @return ApiPutCompanyContactsByParentIdGroupsByIdRequest
*/
func (a *ContactGroupsAPIService) PutCompanyContactsByParentIdGroupsById(ctx context.Context, id int32, parentId int32) ApiPutCompanyContactsByParentIdGroupsByIdRequest {
	return ApiPutCompanyContactsByParentIdGroupsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return ContactGroup
func (a *ContactGroupsAPIService) PutCompanyContactsByParentIdGroupsByIdExecute(r ApiPutCompanyContactsByParentIdGroupsByIdRequest) (*ContactGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContactGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContactGroupsAPIService.PutCompanyContactsByParentIdGroupsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/contacts/{parentId}/groups/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.contactGroup == nil {
		return localVarReturnValue, nil, reportError("contactGroup is required and must be specified")
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
	localVarPostBody = r.contactGroup
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
