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


// AgreementAdjustmentsAPIService AgreementAdjustmentsAPI service
type AgreementAdjustmentsAPIService service

type ApiDeleteFinanceAgreementsByParentIdAdjustmentsByIdRequest struct {
	ctx context.Context
	ApiService *AgreementAdjustmentsAPIService
	id int32
	parentId int32
	clientId *string
}

// 
func (r ApiDeleteFinanceAgreementsByParentIdAdjustmentsByIdRequest) ClientId(clientId string) ApiDeleteFinanceAgreementsByParentIdAdjustmentsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiDeleteFinanceAgreementsByParentIdAdjustmentsByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteFinanceAgreementsByParentIdAdjustmentsByIdExecute(r)
}

/*
DeleteFinanceAgreementsByParentIdAdjustmentsById Delete Adjustment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id adjustmentId
 @param parentId agreementId
 @return ApiDeleteFinanceAgreementsByParentIdAdjustmentsByIdRequest
*/
func (a *AgreementAdjustmentsAPIService) DeleteFinanceAgreementsByParentIdAdjustmentsById(ctx context.Context, id int32, parentId int32) ApiDeleteFinanceAgreementsByParentIdAdjustmentsByIdRequest {
	return ApiDeleteFinanceAgreementsByParentIdAdjustmentsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
func (a *AgreementAdjustmentsAPIService) DeleteFinanceAgreementsByParentIdAdjustmentsByIdExecute(r ApiDeleteFinanceAgreementsByParentIdAdjustmentsByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgreementAdjustmentsAPIService.DeleteFinanceAgreementsByParentIdAdjustmentsById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finance/agreements/{parentId}/adjustments/{id}"
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

type ApiGetFinanceAgreementsByParentIdAdjustmentsRequest struct {
	ctx context.Context
	ApiService *AgreementAdjustmentsAPIService
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
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsRequest) Conditions(conditions string) ApiGetFinanceAgreementsByParentIdAdjustmentsRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsRequest) ChildConditions(childConditions string) ApiGetFinanceAgreementsByParentIdAdjustmentsRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsRequest) CustomFieldConditions(customFieldConditions string) ApiGetFinanceAgreementsByParentIdAdjustmentsRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsRequest) OrderBy(orderBy string) ApiGetFinanceAgreementsByParentIdAdjustmentsRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsRequest) Fields(fields string) ApiGetFinanceAgreementsByParentIdAdjustmentsRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsRequest) Page(page int32) ApiGetFinanceAgreementsByParentIdAdjustmentsRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsRequest) PageSize(pageSize int32) ApiGetFinanceAgreementsByParentIdAdjustmentsRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsRequest) PageId(pageId int32) ApiGetFinanceAgreementsByParentIdAdjustmentsRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsRequest) ClientId(clientId string) ApiGetFinanceAgreementsByParentIdAdjustmentsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetFinanceAgreementsByParentIdAdjustmentsRequest) Execute() ([]AgreementAdjustment, *http.Response, error) {
	return r.ApiService.GetFinanceAgreementsByParentIdAdjustmentsExecute(r)
}

/*
GetFinanceAgreementsByParentIdAdjustments Get List of Adjustment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId agreementId
 @return ApiGetFinanceAgreementsByParentIdAdjustmentsRequest
*/
func (a *AgreementAdjustmentsAPIService) GetFinanceAgreementsByParentIdAdjustments(ctx context.Context, parentId int32) ApiGetFinanceAgreementsByParentIdAdjustmentsRequest {
	return ApiGetFinanceAgreementsByParentIdAdjustmentsRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return []AgreementAdjustment
func (a *AgreementAdjustmentsAPIService) GetFinanceAgreementsByParentIdAdjustmentsExecute(r ApiGetFinanceAgreementsByParentIdAdjustmentsRequest) ([]AgreementAdjustment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AgreementAdjustment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgreementAdjustmentsAPIService.GetFinanceAgreementsByParentIdAdjustments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finance/agreements/{parentId}/adjustments"
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

type ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest struct {
	ctx context.Context
	ApiService *AgreementAdjustmentsAPIService
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
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest) Conditions(conditions string) ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest) ChildConditions(childConditions string) ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest) OrderBy(orderBy string) ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest) Fields(fields string) ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest) Page(page int32) ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest) PageSize(pageSize int32) ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest) PageId(pageId int32) ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest) ClientId(clientId string) ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest) Execute() (*AgreementAdjustment, *http.Response, error) {
	return r.ApiService.GetFinanceAgreementsByParentIdAdjustmentsByIdExecute(r)
}

/*
GetFinanceAgreementsByParentIdAdjustmentsById Get Adjustment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id adjustmentId
 @param parentId agreementId
 @return ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest
*/
func (a *AgreementAdjustmentsAPIService) GetFinanceAgreementsByParentIdAdjustmentsById(ctx context.Context, id int32, parentId int32) ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest {
	return ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return AgreementAdjustment
func (a *AgreementAdjustmentsAPIService) GetFinanceAgreementsByParentIdAdjustmentsByIdExecute(r ApiGetFinanceAgreementsByParentIdAdjustmentsByIdRequest) (*AgreementAdjustment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgreementAdjustment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgreementAdjustmentsAPIService.GetFinanceAgreementsByParentIdAdjustmentsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finance/agreements/{parentId}/adjustments/{id}"
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

type ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest struct {
	ctx context.Context
	ApiService *AgreementAdjustmentsAPIService
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
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest) Conditions(conditions string) ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest) ChildConditions(childConditions string) ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest) OrderBy(orderBy string) ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest) Fields(fields string) ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest) Page(page int32) ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest) PageSize(pageSize int32) ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest) PageId(pageId int32) ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest) ClientId(clientId string) ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetFinanceAgreementsByParentIdAdjustmentsCountExecute(r)
}

/*
GetFinanceAgreementsByParentIdAdjustmentsCount Get Count of Adjustment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId agreementId
 @return ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest
*/
func (a *AgreementAdjustmentsAPIService) GetFinanceAgreementsByParentIdAdjustmentsCount(ctx context.Context, parentId int32) ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest {
	return ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return Count
func (a *AgreementAdjustmentsAPIService) GetFinanceAgreementsByParentIdAdjustmentsCountExecute(r ApiGetFinanceAgreementsByParentIdAdjustmentsCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgreementAdjustmentsAPIService.GetFinanceAgreementsByParentIdAdjustmentsCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finance/agreements/{parentId}/adjustments/count"
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

type ApiPatchFinanceAgreementsByParentIdAdjustmentsByIdRequest struct {
	ctx context.Context
	ApiService *AgreementAdjustmentsAPIService
	id int32
	parentId int32
	patchOperation *[]PatchOperation
	clientId *string
}

// List of PatchOperation
func (r ApiPatchFinanceAgreementsByParentIdAdjustmentsByIdRequest) PatchOperation(patchOperation []PatchOperation) ApiPatchFinanceAgreementsByParentIdAdjustmentsByIdRequest {
	r.patchOperation = &patchOperation
	return r
}

// 
func (r ApiPatchFinanceAgreementsByParentIdAdjustmentsByIdRequest) ClientId(clientId string) ApiPatchFinanceAgreementsByParentIdAdjustmentsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPatchFinanceAgreementsByParentIdAdjustmentsByIdRequest) Execute() (*AgreementAdjustment, *http.Response, error) {
	return r.ApiService.PatchFinanceAgreementsByParentIdAdjustmentsByIdExecute(r)
}

/*
PatchFinanceAgreementsByParentIdAdjustmentsById Patch Adjustment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id adjustmentId
 @param parentId agreementId
 @return ApiPatchFinanceAgreementsByParentIdAdjustmentsByIdRequest
*/
func (a *AgreementAdjustmentsAPIService) PatchFinanceAgreementsByParentIdAdjustmentsById(ctx context.Context, id int32, parentId int32) ApiPatchFinanceAgreementsByParentIdAdjustmentsByIdRequest {
	return ApiPatchFinanceAgreementsByParentIdAdjustmentsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return AgreementAdjustment
func (a *AgreementAdjustmentsAPIService) PatchFinanceAgreementsByParentIdAdjustmentsByIdExecute(r ApiPatchFinanceAgreementsByParentIdAdjustmentsByIdRequest) (*AgreementAdjustment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgreementAdjustment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgreementAdjustmentsAPIService.PatchFinanceAgreementsByParentIdAdjustmentsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finance/agreements/{parentId}/adjustments/{id}"
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

type ApiPostFinanceAgreementsByParentIdAdjustmentsRequest struct {
	ctx context.Context
	ApiService *AgreementAdjustmentsAPIService
	parentId int32
	agreementAdjustment *AgreementAdjustment
	clientId *string
}

// adjustment
func (r ApiPostFinanceAgreementsByParentIdAdjustmentsRequest) AgreementAdjustment(agreementAdjustment AgreementAdjustment) ApiPostFinanceAgreementsByParentIdAdjustmentsRequest {
	r.agreementAdjustment = &agreementAdjustment
	return r
}

// 
func (r ApiPostFinanceAgreementsByParentIdAdjustmentsRequest) ClientId(clientId string) ApiPostFinanceAgreementsByParentIdAdjustmentsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostFinanceAgreementsByParentIdAdjustmentsRequest) Execute() (*AgreementAdjustment, *http.Response, error) {
	return r.ApiService.PostFinanceAgreementsByParentIdAdjustmentsExecute(r)
}

/*
PostFinanceAgreementsByParentIdAdjustments Post Adjustment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId agreementId
 @return ApiPostFinanceAgreementsByParentIdAdjustmentsRequest
*/
func (a *AgreementAdjustmentsAPIService) PostFinanceAgreementsByParentIdAdjustments(ctx context.Context, parentId int32) ApiPostFinanceAgreementsByParentIdAdjustmentsRequest {
	return ApiPostFinanceAgreementsByParentIdAdjustmentsRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return AgreementAdjustment
func (a *AgreementAdjustmentsAPIService) PostFinanceAgreementsByParentIdAdjustmentsExecute(r ApiPostFinanceAgreementsByParentIdAdjustmentsRequest) (*AgreementAdjustment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgreementAdjustment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgreementAdjustmentsAPIService.PostFinanceAgreementsByParentIdAdjustments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finance/agreements/{parentId}/adjustments"
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agreementAdjustment == nil {
		return localVarReturnValue, nil, reportError("agreementAdjustment is required and must be specified")
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
	localVarPostBody = r.agreementAdjustment
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

type ApiPutFinanceAgreementsByParentIdAdjustmentsByIdRequest struct {
	ctx context.Context
	ApiService *AgreementAdjustmentsAPIService
	id int32
	parentId int32
	agreementAdjustment *AgreementAdjustment
	clientId *string
}

// adjustment
func (r ApiPutFinanceAgreementsByParentIdAdjustmentsByIdRequest) AgreementAdjustment(agreementAdjustment AgreementAdjustment) ApiPutFinanceAgreementsByParentIdAdjustmentsByIdRequest {
	r.agreementAdjustment = &agreementAdjustment
	return r
}

// 
func (r ApiPutFinanceAgreementsByParentIdAdjustmentsByIdRequest) ClientId(clientId string) ApiPutFinanceAgreementsByParentIdAdjustmentsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPutFinanceAgreementsByParentIdAdjustmentsByIdRequest) Execute() (*AgreementAdjustment, *http.Response, error) {
	return r.ApiService.PutFinanceAgreementsByParentIdAdjustmentsByIdExecute(r)
}

/*
PutFinanceAgreementsByParentIdAdjustmentsById Put Adjustment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id adjustmentId
 @param parentId agreementId
 @return ApiPutFinanceAgreementsByParentIdAdjustmentsByIdRequest
*/
func (a *AgreementAdjustmentsAPIService) PutFinanceAgreementsByParentIdAdjustmentsById(ctx context.Context, id int32, parentId int32) ApiPutFinanceAgreementsByParentIdAdjustmentsByIdRequest {
	return ApiPutFinanceAgreementsByParentIdAdjustmentsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return AgreementAdjustment
func (a *AgreementAdjustmentsAPIService) PutFinanceAgreementsByParentIdAdjustmentsByIdExecute(r ApiPutFinanceAgreementsByParentIdAdjustmentsByIdRequest) (*AgreementAdjustment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgreementAdjustment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgreementAdjustmentsAPIService.PutFinanceAgreementsByParentIdAdjustmentsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finance/agreements/{parentId}/adjustments/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agreementAdjustment == nil {
		return localVarReturnValue, nil, reportError("agreementAdjustment is required and must be specified")
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
	localVarPostBody = r.agreementAdjustment
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
