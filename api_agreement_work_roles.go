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


// AgreementWorkRolesAPIService AgreementWorkRolesAPI service
type AgreementWorkRolesAPIService service

type ApiDeleteFinanceAgreementsByParentIdWorkrolesByIdRequest struct {
	ctx context.Context
	ApiService *AgreementWorkRolesAPIService
	id int32
	parentId int32
	clientId *string
}

// 
func (r ApiDeleteFinanceAgreementsByParentIdWorkrolesByIdRequest) ClientId(clientId string) ApiDeleteFinanceAgreementsByParentIdWorkrolesByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiDeleteFinanceAgreementsByParentIdWorkrolesByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteFinanceAgreementsByParentIdWorkrolesByIdExecute(r)
}

/*
DeleteFinanceAgreementsByParentIdWorkrolesById Delete AgreementWorkRole

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id workroleId
 @param parentId agreementId
 @return ApiDeleteFinanceAgreementsByParentIdWorkrolesByIdRequest
*/
func (a *AgreementWorkRolesAPIService) DeleteFinanceAgreementsByParentIdWorkrolesById(ctx context.Context, id int32, parentId int32) ApiDeleteFinanceAgreementsByParentIdWorkrolesByIdRequest {
	return ApiDeleteFinanceAgreementsByParentIdWorkrolesByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
func (a *AgreementWorkRolesAPIService) DeleteFinanceAgreementsByParentIdWorkrolesByIdExecute(r ApiDeleteFinanceAgreementsByParentIdWorkrolesByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgreementWorkRolesAPIService.DeleteFinanceAgreementsByParentIdWorkrolesById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finance/agreements/{parentId}/workroles/{id}"
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

type ApiGetFinanceAgreementsByParentIdWorkrolesRequest struct {
	ctx context.Context
	ApiService *AgreementWorkRolesAPIService
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
func (r ApiGetFinanceAgreementsByParentIdWorkrolesRequest) Conditions(conditions string) ApiGetFinanceAgreementsByParentIdWorkrolesRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesRequest) ChildConditions(childConditions string) ApiGetFinanceAgreementsByParentIdWorkrolesRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesRequest) CustomFieldConditions(customFieldConditions string) ApiGetFinanceAgreementsByParentIdWorkrolesRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesRequest) OrderBy(orderBy string) ApiGetFinanceAgreementsByParentIdWorkrolesRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesRequest) Fields(fields string) ApiGetFinanceAgreementsByParentIdWorkrolesRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesRequest) Page(page int32) ApiGetFinanceAgreementsByParentIdWorkrolesRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesRequest) PageSize(pageSize int32) ApiGetFinanceAgreementsByParentIdWorkrolesRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesRequest) PageId(pageId int32) ApiGetFinanceAgreementsByParentIdWorkrolesRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesRequest) ClientId(clientId string) ApiGetFinanceAgreementsByParentIdWorkrolesRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetFinanceAgreementsByParentIdWorkrolesRequest) Execute() ([]AgreementWorkRole, *http.Response, error) {
	return r.ApiService.GetFinanceAgreementsByParentIdWorkrolesExecute(r)
}

/*
GetFinanceAgreementsByParentIdWorkroles Get List of AgreementWorkRole

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId agreementId
 @return ApiGetFinanceAgreementsByParentIdWorkrolesRequest
*/
func (a *AgreementWorkRolesAPIService) GetFinanceAgreementsByParentIdWorkroles(ctx context.Context, parentId int32) ApiGetFinanceAgreementsByParentIdWorkrolesRequest {
	return ApiGetFinanceAgreementsByParentIdWorkrolesRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return []AgreementWorkRole
func (a *AgreementWorkRolesAPIService) GetFinanceAgreementsByParentIdWorkrolesExecute(r ApiGetFinanceAgreementsByParentIdWorkrolesRequest) ([]AgreementWorkRole, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AgreementWorkRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgreementWorkRolesAPIService.GetFinanceAgreementsByParentIdWorkroles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finance/agreements/{parentId}/workroles"
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

type ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest struct {
	ctx context.Context
	ApiService *AgreementWorkRolesAPIService
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
func (r ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest) Conditions(conditions string) ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest) ChildConditions(childConditions string) ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest) OrderBy(orderBy string) ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest) Fields(fields string) ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest) Page(page int32) ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest) PageSize(pageSize int32) ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest) PageId(pageId int32) ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest) ClientId(clientId string) ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest) Execute() (*AgreementWorkRole, *http.Response, error) {
	return r.ApiService.GetFinanceAgreementsByParentIdWorkrolesByIdExecute(r)
}

/*
GetFinanceAgreementsByParentIdWorkrolesById Get AgreementWorkRole

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id workroleId
 @param parentId agreementId
 @return ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest
*/
func (a *AgreementWorkRolesAPIService) GetFinanceAgreementsByParentIdWorkrolesById(ctx context.Context, id int32, parentId int32) ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest {
	return ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return AgreementWorkRole
func (a *AgreementWorkRolesAPIService) GetFinanceAgreementsByParentIdWorkrolesByIdExecute(r ApiGetFinanceAgreementsByParentIdWorkrolesByIdRequest) (*AgreementWorkRole, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgreementWorkRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgreementWorkRolesAPIService.GetFinanceAgreementsByParentIdWorkrolesById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finance/agreements/{parentId}/workroles/{id}"
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

type ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest struct {
	ctx context.Context
	ApiService *AgreementWorkRolesAPIService
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
func (r ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest) Conditions(conditions string) ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest) ChildConditions(childConditions string) ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest) OrderBy(orderBy string) ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest) Fields(fields string) ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest) Page(page int32) ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest) PageSize(pageSize int32) ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest) PageId(pageId int32) ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest) ClientId(clientId string) ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetFinanceAgreementsByParentIdWorkrolesCountExecute(r)
}

/*
GetFinanceAgreementsByParentIdWorkrolesCount Get Count of AgreementWorkRole

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId agreementId
 @return ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest
*/
func (a *AgreementWorkRolesAPIService) GetFinanceAgreementsByParentIdWorkrolesCount(ctx context.Context, parentId int32) ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest {
	return ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return Count
func (a *AgreementWorkRolesAPIService) GetFinanceAgreementsByParentIdWorkrolesCountExecute(r ApiGetFinanceAgreementsByParentIdWorkrolesCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgreementWorkRolesAPIService.GetFinanceAgreementsByParentIdWorkrolesCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finance/agreements/{parentId}/workroles/count"
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

type ApiPatchFinanceAgreementsByParentIdWorkrolesByIdRequest struct {
	ctx context.Context
	ApiService *AgreementWorkRolesAPIService
	id int32
	parentId int32
	patchOperation *[]PatchOperation
	clientId *string
}

// List of PatchOperation
func (r ApiPatchFinanceAgreementsByParentIdWorkrolesByIdRequest) PatchOperation(patchOperation []PatchOperation) ApiPatchFinanceAgreementsByParentIdWorkrolesByIdRequest {
	r.patchOperation = &patchOperation
	return r
}

// 
func (r ApiPatchFinanceAgreementsByParentIdWorkrolesByIdRequest) ClientId(clientId string) ApiPatchFinanceAgreementsByParentIdWorkrolesByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPatchFinanceAgreementsByParentIdWorkrolesByIdRequest) Execute() (*AgreementWorkRole, *http.Response, error) {
	return r.ApiService.PatchFinanceAgreementsByParentIdWorkrolesByIdExecute(r)
}

/*
PatchFinanceAgreementsByParentIdWorkrolesById Patch AgreementWorkRole

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id workroleId
 @param parentId agreementId
 @return ApiPatchFinanceAgreementsByParentIdWorkrolesByIdRequest
*/
func (a *AgreementWorkRolesAPIService) PatchFinanceAgreementsByParentIdWorkrolesById(ctx context.Context, id int32, parentId int32) ApiPatchFinanceAgreementsByParentIdWorkrolesByIdRequest {
	return ApiPatchFinanceAgreementsByParentIdWorkrolesByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return AgreementWorkRole
func (a *AgreementWorkRolesAPIService) PatchFinanceAgreementsByParentIdWorkrolesByIdExecute(r ApiPatchFinanceAgreementsByParentIdWorkrolesByIdRequest) (*AgreementWorkRole, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgreementWorkRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgreementWorkRolesAPIService.PatchFinanceAgreementsByParentIdWorkrolesById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finance/agreements/{parentId}/workroles/{id}"
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

type ApiPostFinanceAgreementsByParentIdWorkrolesRequest struct {
	ctx context.Context
	ApiService *AgreementWorkRolesAPIService
	parentId int32
	agreementWorkRole *AgreementWorkRole
	clientId *string
}

// workRole
func (r ApiPostFinanceAgreementsByParentIdWorkrolesRequest) AgreementWorkRole(agreementWorkRole AgreementWorkRole) ApiPostFinanceAgreementsByParentIdWorkrolesRequest {
	r.agreementWorkRole = &agreementWorkRole
	return r
}

// 
func (r ApiPostFinanceAgreementsByParentIdWorkrolesRequest) ClientId(clientId string) ApiPostFinanceAgreementsByParentIdWorkrolesRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostFinanceAgreementsByParentIdWorkrolesRequest) Execute() (*AgreementWorkRole, *http.Response, error) {
	return r.ApiService.PostFinanceAgreementsByParentIdWorkrolesExecute(r)
}

/*
PostFinanceAgreementsByParentIdWorkroles Post AgreementWorkRole

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId agreementId
 @return ApiPostFinanceAgreementsByParentIdWorkrolesRequest
*/
func (a *AgreementWorkRolesAPIService) PostFinanceAgreementsByParentIdWorkroles(ctx context.Context, parentId int32) ApiPostFinanceAgreementsByParentIdWorkrolesRequest {
	return ApiPostFinanceAgreementsByParentIdWorkrolesRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return AgreementWorkRole
func (a *AgreementWorkRolesAPIService) PostFinanceAgreementsByParentIdWorkrolesExecute(r ApiPostFinanceAgreementsByParentIdWorkrolesRequest) (*AgreementWorkRole, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgreementWorkRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgreementWorkRolesAPIService.PostFinanceAgreementsByParentIdWorkroles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finance/agreements/{parentId}/workroles"
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agreementWorkRole == nil {
		return localVarReturnValue, nil, reportError("agreementWorkRole is required and must be specified")
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
	localVarPostBody = r.agreementWorkRole
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

type ApiPutFinanceAgreementsByParentIdWorkrolesByIdRequest struct {
	ctx context.Context
	ApiService *AgreementWorkRolesAPIService
	id int32
	parentId int32
	agreementWorkRole *AgreementWorkRole
	clientId *string
}

// workRole
func (r ApiPutFinanceAgreementsByParentIdWorkrolesByIdRequest) AgreementWorkRole(agreementWorkRole AgreementWorkRole) ApiPutFinanceAgreementsByParentIdWorkrolesByIdRequest {
	r.agreementWorkRole = &agreementWorkRole
	return r
}

// 
func (r ApiPutFinanceAgreementsByParentIdWorkrolesByIdRequest) ClientId(clientId string) ApiPutFinanceAgreementsByParentIdWorkrolesByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPutFinanceAgreementsByParentIdWorkrolesByIdRequest) Execute() (*AgreementWorkRole, *http.Response, error) {
	return r.ApiService.PutFinanceAgreementsByParentIdWorkrolesByIdExecute(r)
}

/*
PutFinanceAgreementsByParentIdWorkrolesById Put AgreementWorkRole

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id workroleId
 @param parentId agreementId
 @return ApiPutFinanceAgreementsByParentIdWorkrolesByIdRequest
*/
func (a *AgreementWorkRolesAPIService) PutFinanceAgreementsByParentIdWorkrolesById(ctx context.Context, id int32, parentId int32) ApiPutFinanceAgreementsByParentIdWorkrolesByIdRequest {
	return ApiPutFinanceAgreementsByParentIdWorkrolesByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return AgreementWorkRole
func (a *AgreementWorkRolesAPIService) PutFinanceAgreementsByParentIdWorkrolesByIdExecute(r ApiPutFinanceAgreementsByParentIdWorkrolesByIdRequest) (*AgreementWorkRole, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgreementWorkRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgreementWorkRolesAPIService.PutFinanceAgreementsByParentIdWorkrolesById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finance/agreements/{parentId}/workroles/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agreementWorkRole == nil {
		return localVarReturnValue, nil, reportError("agreementWorkRole is required and must be specified")
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
	localVarPostBody = r.agreementWorkRole
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
