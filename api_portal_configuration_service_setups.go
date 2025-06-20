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


// PortalConfigurationServiceSetupsAPIService PortalConfigurationServiceSetupsAPI service
type PortalConfigurationServiceSetupsAPIService service

type ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest struct {
	ctx context.Context
	ApiService *PortalConfigurationServiceSetupsAPIService
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
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest) Conditions(conditions string) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest) ChildConditions(childConditions string) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest) CustomFieldConditions(customFieldConditions string) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest) OrderBy(orderBy string) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest) Fields(fields string) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest) Page(page int32) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest) PageSize(pageSize int32) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest) PageId(pageId int32) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest) ClientId(clientId string) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest) Execute() ([]PortalConfigurationServiceSetup, *http.Response, error) {
	return r.ApiService.GetCompanyPortalConfigurationsByParentIdServiceSetupsExecute(r)
}

/*
GetCompanyPortalConfigurationsByParentIdServiceSetups Get List of PortalConfigurationServiceSetup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId portalConfigurationId
 @return ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest
*/
func (a *PortalConfigurationServiceSetupsAPIService) GetCompanyPortalConfigurationsByParentIdServiceSetups(ctx context.Context, parentId int32) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest {
	return ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return []PortalConfigurationServiceSetup
func (a *PortalConfigurationServiceSetupsAPIService) GetCompanyPortalConfigurationsByParentIdServiceSetupsExecute(r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsRequest) ([]PortalConfigurationServiceSetup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []PortalConfigurationServiceSetup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalConfigurationServiceSetupsAPIService.GetCompanyPortalConfigurationsByParentIdServiceSetups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/portalConfigurations/{parentId}/serviceSetups"
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

type ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest struct {
	ctx context.Context
	ApiService *PortalConfigurationServiceSetupsAPIService
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
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest) Conditions(conditions string) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest) ChildConditions(childConditions string) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest) OrderBy(orderBy string) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest) Fields(fields string) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest) Page(page int32) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest) PageSize(pageSize int32) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest) PageId(pageId int32) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest) ClientId(clientId string) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest) Execute() (*PortalConfigurationServiceSetup, *http.Response, error) {
	return r.ApiService.GetCompanyPortalConfigurationsByParentIdServiceSetupsByIdExecute(r)
}

/*
GetCompanyPortalConfigurationsByParentIdServiceSetupsById Get PortalConfigurationServiceSetup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id serviceSetupId
 @param parentId portalConfigurationId
 @return ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest
*/
func (a *PortalConfigurationServiceSetupsAPIService) GetCompanyPortalConfigurationsByParentIdServiceSetupsById(ctx context.Context, id int32, parentId int32) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest {
	return ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return PortalConfigurationServiceSetup
func (a *PortalConfigurationServiceSetupsAPIService) GetCompanyPortalConfigurationsByParentIdServiceSetupsByIdExecute(r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest) (*PortalConfigurationServiceSetup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PortalConfigurationServiceSetup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalConfigurationServiceSetupsAPIService.GetCompanyPortalConfigurationsByParentIdServiceSetupsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/portalConfigurations/{parentId}/serviceSetups/{id}"
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

type ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest struct {
	ctx context.Context
	ApiService *PortalConfigurationServiceSetupsAPIService
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
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest) Conditions(conditions string) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest) ChildConditions(childConditions string) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest) OrderBy(orderBy string) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest) Fields(fields string) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest) Page(page int32) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest) PageSize(pageSize int32) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest) PageId(pageId int32) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest) ClientId(clientId string) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetCompanyPortalConfigurationsByParentIdServiceSetupsCountExecute(r)
}

/*
GetCompanyPortalConfigurationsByParentIdServiceSetupsCount Get Count of PortalConfigurationServiceSetup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId portalConfigurationId
 @return ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest
*/
func (a *PortalConfigurationServiceSetupsAPIService) GetCompanyPortalConfigurationsByParentIdServiceSetupsCount(ctx context.Context, parentId int32) ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest {
	return ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return Count
func (a *PortalConfigurationServiceSetupsAPIService) GetCompanyPortalConfigurationsByParentIdServiceSetupsCountExecute(r ApiGetCompanyPortalConfigurationsByParentIdServiceSetupsCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalConfigurationServiceSetupsAPIService.GetCompanyPortalConfigurationsByParentIdServiceSetupsCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/portalConfigurations/{parentId}/serviceSetups/count"
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

type ApiPatchCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest struct {
	ctx context.Context
	ApiService *PortalConfigurationServiceSetupsAPIService
	id int32
	parentId int32
	patchOperation *[]PatchOperation
	clientId *string
}

// List of PatchOperation
func (r ApiPatchCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest) PatchOperation(patchOperation []PatchOperation) ApiPatchCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest {
	r.patchOperation = &patchOperation
	return r
}

// 
func (r ApiPatchCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest) ClientId(clientId string) ApiPatchCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPatchCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest) Execute() (*PortalConfigurationServiceSetup, *http.Response, error) {
	return r.ApiService.PatchCompanyPortalConfigurationsByParentIdServiceSetupsByIdExecute(r)
}

/*
PatchCompanyPortalConfigurationsByParentIdServiceSetupsById Patch PortalConfigurationServiceSetup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id serviceSetupId
 @param parentId portalConfigurationId
 @return ApiPatchCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest
*/
func (a *PortalConfigurationServiceSetupsAPIService) PatchCompanyPortalConfigurationsByParentIdServiceSetupsById(ctx context.Context, id int32, parentId int32) ApiPatchCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest {
	return ApiPatchCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return PortalConfigurationServiceSetup
func (a *PortalConfigurationServiceSetupsAPIService) PatchCompanyPortalConfigurationsByParentIdServiceSetupsByIdExecute(r ApiPatchCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest) (*PortalConfigurationServiceSetup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PortalConfigurationServiceSetup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalConfigurationServiceSetupsAPIService.PatchCompanyPortalConfigurationsByParentIdServiceSetupsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/portalConfigurations/{parentId}/serviceSetups/{id}"
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

type ApiPutCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest struct {
	ctx context.Context
	ApiService *PortalConfigurationServiceSetupsAPIService
	id int32
	parentId int32
	portalConfigurationServiceSetup *PortalConfigurationServiceSetup
	clientId *string
}

// portalConfigurationServiceSetup
func (r ApiPutCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest) PortalConfigurationServiceSetup(portalConfigurationServiceSetup PortalConfigurationServiceSetup) ApiPutCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest {
	r.portalConfigurationServiceSetup = &portalConfigurationServiceSetup
	return r
}

// 
func (r ApiPutCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest) ClientId(clientId string) ApiPutCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPutCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest) Execute() (*PortalConfigurationServiceSetup, *http.Response, error) {
	return r.ApiService.PutCompanyPortalConfigurationsByParentIdServiceSetupsByIdExecute(r)
}

/*
PutCompanyPortalConfigurationsByParentIdServiceSetupsById Put PortalConfigurationServiceSetup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id serviceSetupId
 @param parentId portalConfigurationId
 @return ApiPutCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest
*/
func (a *PortalConfigurationServiceSetupsAPIService) PutCompanyPortalConfigurationsByParentIdServiceSetupsById(ctx context.Context, id int32, parentId int32) ApiPutCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest {
	return ApiPutCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return PortalConfigurationServiceSetup
func (a *PortalConfigurationServiceSetupsAPIService) PutCompanyPortalConfigurationsByParentIdServiceSetupsByIdExecute(r ApiPutCompanyPortalConfigurationsByParentIdServiceSetupsByIdRequest) (*PortalConfigurationServiceSetup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PortalConfigurationServiceSetup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalConfigurationServiceSetupsAPIService.PutCompanyPortalConfigurationsByParentIdServiceSetupsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/portalConfigurations/{parentId}/serviceSetups/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.portalConfigurationServiceSetup == nil {
		return localVarReturnValue, nil, reportError("portalConfigurationServiceSetup is required and must be specified")
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
	localVarPostBody = r.portalConfigurationServiceSetup
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
