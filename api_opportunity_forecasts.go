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


// OpportunityForecastsAPIService OpportunityForecastsAPI service
type OpportunityForecastsAPIService service

type ApiDeleteSalesOpportunitiesByParentIdForecastRequest struct {
	ctx context.Context
	ApiService *OpportunityForecastsAPIService
	parentId int32
	clientId *string
}

// 
func (r ApiDeleteSalesOpportunitiesByParentIdForecastRequest) ClientId(clientId string) ApiDeleteSalesOpportunitiesByParentIdForecastRequest {
	r.clientId = &clientId
	return r
}

func (r ApiDeleteSalesOpportunitiesByParentIdForecastRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSalesOpportunitiesByParentIdForecastExecute(r)
}

/*
DeleteSalesOpportunitiesByParentIdForecast Delete Forecast

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId opportunityId
 @return ApiDeleteSalesOpportunitiesByParentIdForecastRequest
*/
func (a *OpportunityForecastsAPIService) DeleteSalesOpportunitiesByParentIdForecast(ctx context.Context, parentId int32) ApiDeleteSalesOpportunitiesByParentIdForecastRequest {
	return ApiDeleteSalesOpportunitiesByParentIdForecastRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
func (a *OpportunityForecastsAPIService) DeleteSalesOpportunitiesByParentIdForecastExecute(r ApiDeleteSalesOpportunitiesByParentIdForecastRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpportunityForecastsAPIService.DeleteSalesOpportunitiesByParentIdForecast")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales/opportunities/{parentId}/forecast/"
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

type ApiGetSalesOpportunitiesByParentIdForecastRequest struct {
	ctx context.Context
	ApiService *OpportunityForecastsAPIService
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
func (r ApiGetSalesOpportunitiesByParentIdForecastRequest) Conditions(conditions string) ApiGetSalesOpportunitiesByParentIdForecastRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSalesOpportunitiesByParentIdForecastRequest) ChildConditions(childConditions string) ApiGetSalesOpportunitiesByParentIdForecastRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSalesOpportunitiesByParentIdForecastRequest) CustomFieldConditions(customFieldConditions string) ApiGetSalesOpportunitiesByParentIdForecastRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSalesOpportunitiesByParentIdForecastRequest) OrderBy(orderBy string) ApiGetSalesOpportunitiesByParentIdForecastRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSalesOpportunitiesByParentIdForecastRequest) Fields(fields string) ApiGetSalesOpportunitiesByParentIdForecastRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSalesOpportunitiesByParentIdForecastRequest) Page(page int32) ApiGetSalesOpportunitiesByParentIdForecastRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSalesOpportunitiesByParentIdForecastRequest) PageSize(pageSize int32) ApiGetSalesOpportunitiesByParentIdForecastRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSalesOpportunitiesByParentIdForecastRequest) PageId(pageId int32) ApiGetSalesOpportunitiesByParentIdForecastRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSalesOpportunitiesByParentIdForecastRequest) ClientId(clientId string) ApiGetSalesOpportunitiesByParentIdForecastRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSalesOpportunitiesByParentIdForecastRequest) Execute() ([]Forecast, *http.Response, error) {
	return r.ApiService.GetSalesOpportunitiesByParentIdForecastExecute(r)
}

/*
GetSalesOpportunitiesByParentIdForecast Get List of Forecast

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId opportunityId
 @return ApiGetSalesOpportunitiesByParentIdForecastRequest
*/
func (a *OpportunityForecastsAPIService) GetSalesOpportunitiesByParentIdForecast(ctx context.Context, parentId int32) ApiGetSalesOpportunitiesByParentIdForecastRequest {
	return ApiGetSalesOpportunitiesByParentIdForecastRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return []Forecast
func (a *OpportunityForecastsAPIService) GetSalesOpportunitiesByParentIdForecastExecute(r ApiGetSalesOpportunitiesByParentIdForecastRequest) ([]Forecast, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Forecast
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpportunityForecastsAPIService.GetSalesOpportunitiesByParentIdForecast")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales/opportunities/{parentId}/forecast"
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

type ApiGetSalesOpportunitiesByParentIdForecastCountRequest struct {
	ctx context.Context
	ApiService *OpportunityForecastsAPIService
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
func (r ApiGetSalesOpportunitiesByParentIdForecastCountRequest) Conditions(conditions string) ApiGetSalesOpportunitiesByParentIdForecastCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetSalesOpportunitiesByParentIdForecastCountRequest) ChildConditions(childConditions string) ApiGetSalesOpportunitiesByParentIdForecastCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetSalesOpportunitiesByParentIdForecastCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetSalesOpportunitiesByParentIdForecastCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetSalesOpportunitiesByParentIdForecastCountRequest) OrderBy(orderBy string) ApiGetSalesOpportunitiesByParentIdForecastCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetSalesOpportunitiesByParentIdForecastCountRequest) Fields(fields string) ApiGetSalesOpportunitiesByParentIdForecastCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetSalesOpportunitiesByParentIdForecastCountRequest) Page(page int32) ApiGetSalesOpportunitiesByParentIdForecastCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetSalesOpportunitiesByParentIdForecastCountRequest) PageSize(pageSize int32) ApiGetSalesOpportunitiesByParentIdForecastCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetSalesOpportunitiesByParentIdForecastCountRequest) PageId(pageId int32) ApiGetSalesOpportunitiesByParentIdForecastCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetSalesOpportunitiesByParentIdForecastCountRequest) ClientId(clientId string) ApiGetSalesOpportunitiesByParentIdForecastCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetSalesOpportunitiesByParentIdForecastCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetSalesOpportunitiesByParentIdForecastCountExecute(r)
}

/*
GetSalesOpportunitiesByParentIdForecastCount Get Count of Forecast

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId opportunityId
 @return ApiGetSalesOpportunitiesByParentIdForecastCountRequest
*/
func (a *OpportunityForecastsAPIService) GetSalesOpportunitiesByParentIdForecastCount(ctx context.Context, parentId int32) ApiGetSalesOpportunitiesByParentIdForecastCountRequest {
	return ApiGetSalesOpportunitiesByParentIdForecastCountRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return Count
func (a *OpportunityForecastsAPIService) GetSalesOpportunitiesByParentIdForecastCountExecute(r ApiGetSalesOpportunitiesByParentIdForecastCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpportunityForecastsAPIService.GetSalesOpportunitiesByParentIdForecastCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales/opportunities/{parentId}/forecast/count"
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

type ApiPatchSalesOpportunitiesByParentIdForecastRequest struct {
	ctx context.Context
	ApiService *OpportunityForecastsAPIService
	parentId int32
	patchOperation *[]PatchOperation
	clientId *string
}

// List of PatchOperation
func (r ApiPatchSalesOpportunitiesByParentIdForecastRequest) PatchOperation(patchOperation []PatchOperation) ApiPatchSalesOpportunitiesByParentIdForecastRequest {
	r.patchOperation = &patchOperation
	return r
}

// 
func (r ApiPatchSalesOpportunitiesByParentIdForecastRequest) ClientId(clientId string) ApiPatchSalesOpportunitiesByParentIdForecastRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPatchSalesOpportunitiesByParentIdForecastRequest) Execute() (*Forecast, *http.Response, error) {
	return r.ApiService.PatchSalesOpportunitiesByParentIdForecastExecute(r)
}

/*
PatchSalesOpportunitiesByParentIdForecast Patch Forecast

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId opportunityId
 @return ApiPatchSalesOpportunitiesByParentIdForecastRequest
*/
func (a *OpportunityForecastsAPIService) PatchSalesOpportunitiesByParentIdForecast(ctx context.Context, parentId int32) ApiPatchSalesOpportunitiesByParentIdForecastRequest {
	return ApiPatchSalesOpportunitiesByParentIdForecastRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return Forecast
func (a *OpportunityForecastsAPIService) PatchSalesOpportunitiesByParentIdForecastExecute(r ApiPatchSalesOpportunitiesByParentIdForecastRequest) (*Forecast, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Forecast
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpportunityForecastsAPIService.PatchSalesOpportunitiesByParentIdForecast")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales/opportunities/{parentId}/forecast/"
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

type ApiPostSalesOpportunitiesByParentIdForecastRequest struct {
	ctx context.Context
	ApiService *OpportunityForecastsAPIService
	parentId int32
	forecast *Forecast
	clientId *string
}

// forecast
func (r ApiPostSalesOpportunitiesByParentIdForecastRequest) Forecast(forecast Forecast) ApiPostSalesOpportunitiesByParentIdForecastRequest {
	r.forecast = &forecast
	return r
}

// 
func (r ApiPostSalesOpportunitiesByParentIdForecastRequest) ClientId(clientId string) ApiPostSalesOpportunitiesByParentIdForecastRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostSalesOpportunitiesByParentIdForecastRequest) Execute() (*Forecast, *http.Response, error) {
	return r.ApiService.PostSalesOpportunitiesByParentIdForecastExecute(r)
}

/*
PostSalesOpportunitiesByParentIdForecast Post Forecast

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId opportunityId
 @return ApiPostSalesOpportunitiesByParentIdForecastRequest
*/
func (a *OpportunityForecastsAPIService) PostSalesOpportunitiesByParentIdForecast(ctx context.Context, parentId int32) ApiPostSalesOpportunitiesByParentIdForecastRequest {
	return ApiPostSalesOpportunitiesByParentIdForecastRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return Forecast
func (a *OpportunityForecastsAPIService) PostSalesOpportunitiesByParentIdForecastExecute(r ApiPostSalesOpportunitiesByParentIdForecastRequest) (*Forecast, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Forecast
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpportunityForecastsAPIService.PostSalesOpportunitiesByParentIdForecast")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales/opportunities/{parentId}/forecast"
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.forecast == nil {
		return localVarReturnValue, nil, reportError("forecast is required and must be specified")
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
	localVarPostBody = r.forecast
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

type ApiPostSalesOpportunitiesByParentIdForecastCopyByIdRequest struct {
	ctx context.Context
	ApiService *OpportunityForecastsAPIService
	id int32
	parentId int32
	clientId *string
}

// 
func (r ApiPostSalesOpportunitiesByParentIdForecastCopyByIdRequest) ClientId(clientId string) ApiPostSalesOpportunitiesByParentIdForecastCopyByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostSalesOpportunitiesByParentIdForecastCopyByIdRequest) Execute() (*SuccessResponse, *http.Response, error) {
	return r.ApiService.PostSalesOpportunitiesByParentIdForecastCopyByIdExecute(r)
}

/*
PostSalesOpportunitiesByParentIdForecastCopyById Post SuccessResponse

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id copyId
 @param parentId opportunityId
 @return ApiPostSalesOpportunitiesByParentIdForecastCopyByIdRequest
*/
func (a *OpportunityForecastsAPIService) PostSalesOpportunitiesByParentIdForecastCopyById(ctx context.Context, id int32, parentId int32) ApiPostSalesOpportunitiesByParentIdForecastCopyByIdRequest {
	return ApiPostSalesOpportunitiesByParentIdForecastCopyByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return SuccessResponse
func (a *OpportunityForecastsAPIService) PostSalesOpportunitiesByParentIdForecastCopyByIdExecute(r ApiPostSalesOpportunitiesByParentIdForecastCopyByIdRequest) (*SuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpportunityForecastsAPIService.PostSalesOpportunitiesByParentIdForecastCopyById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales/opportunities/{parentId}/forecast/copy/{id}"
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

type ApiPutSalesOpportunitiesByParentIdForecastRequest struct {
	ctx context.Context
	ApiService *OpportunityForecastsAPIService
	parentId int32
	forecast *Forecast
	clientId *string
}

// forecast
func (r ApiPutSalesOpportunitiesByParentIdForecastRequest) Forecast(forecast Forecast) ApiPutSalesOpportunitiesByParentIdForecastRequest {
	r.forecast = &forecast
	return r
}

// 
func (r ApiPutSalesOpportunitiesByParentIdForecastRequest) ClientId(clientId string) ApiPutSalesOpportunitiesByParentIdForecastRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPutSalesOpportunitiesByParentIdForecastRequest) Execute() (*Forecast, *http.Response, error) {
	return r.ApiService.PutSalesOpportunitiesByParentIdForecastExecute(r)
}

/*
PutSalesOpportunitiesByParentIdForecast Put Forecast

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId opportunityId
 @return ApiPutSalesOpportunitiesByParentIdForecastRequest
*/
func (a *OpportunityForecastsAPIService) PutSalesOpportunitiesByParentIdForecast(ctx context.Context, parentId int32) ApiPutSalesOpportunitiesByParentIdForecastRequest {
	return ApiPutSalesOpportunitiesByParentIdForecastRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return Forecast
func (a *OpportunityForecastsAPIService) PutSalesOpportunitiesByParentIdForecastExecute(r ApiPutSalesOpportunitiesByParentIdForecastRequest) (*Forecast, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Forecast
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpportunityForecastsAPIService.PutSalesOpportunitiesByParentIdForecast")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sales/opportunities/{parentId}/forecast/"
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.forecast == nil {
		return localVarReturnValue, nil, reportError("forecast is required and must be specified")
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
	localVarPostBody = r.forecast
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
