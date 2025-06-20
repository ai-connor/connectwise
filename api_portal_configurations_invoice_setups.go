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


// PortalConfigurationsInvoiceSetupsAPIService PortalConfigurationsInvoiceSetupsAPI service
type PortalConfigurationsInvoiceSetupsAPIService service

type ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest struct {
	ctx context.Context
	ApiService *PortalConfigurationsInvoiceSetupsAPIService
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
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest) Conditions(conditions string) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest) ChildConditions(childConditions string) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest) CustomFieldConditions(customFieldConditions string) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest) OrderBy(orderBy string) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest) Fields(fields string) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest) Page(page int32) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest) PageSize(pageSize int32) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest) PageId(pageId int32) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest) ClientId(clientId string) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest) Execute() ([]PortalConfigurationInvoiceSetup, *http.Response, error) {
	return r.ApiService.GetCompanyPortalConfigurationsByParentIdInvoiceSetupsExecute(r)
}

/*
GetCompanyPortalConfigurationsByParentIdInvoiceSetups Get List of PortalConfigurationInvoiceSetup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId portalConfigurationId
 @return ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest
*/
func (a *PortalConfigurationsInvoiceSetupsAPIService) GetCompanyPortalConfigurationsByParentIdInvoiceSetups(ctx context.Context, parentId int32) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest {
	return ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return []PortalConfigurationInvoiceSetup
func (a *PortalConfigurationsInvoiceSetupsAPIService) GetCompanyPortalConfigurationsByParentIdInvoiceSetupsExecute(r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsRequest) ([]PortalConfigurationInvoiceSetup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []PortalConfigurationInvoiceSetup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalConfigurationsInvoiceSetupsAPIService.GetCompanyPortalConfigurationsByParentIdInvoiceSetups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/portalConfigurations/{parentId}/invoiceSetups"
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

type ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest struct {
	ctx context.Context
	ApiService *PortalConfigurationsInvoiceSetupsAPIService
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
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest) Conditions(conditions string) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest) ChildConditions(childConditions string) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest) OrderBy(orderBy string) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest) Fields(fields string) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest) Page(page int32) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest) PageSize(pageSize int32) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest) PageId(pageId int32) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest) ClientId(clientId string) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest) Execute() (*PortalConfigurationInvoiceSetup, *http.Response, error) {
	return r.ApiService.GetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdExecute(r)
}

/*
GetCompanyPortalConfigurationsByParentIdInvoiceSetupsById Get PortalConfigurationInvoiceSetup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id invoiceSetupId
 @param parentId portalConfigurationId
 @return ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest
*/
func (a *PortalConfigurationsInvoiceSetupsAPIService) GetCompanyPortalConfigurationsByParentIdInvoiceSetupsById(ctx context.Context, id int32, parentId int32) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest {
	return ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return PortalConfigurationInvoiceSetup
func (a *PortalConfigurationsInvoiceSetupsAPIService) GetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdExecute(r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest) (*PortalConfigurationInvoiceSetup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PortalConfigurationInvoiceSetup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalConfigurationsInvoiceSetupsAPIService.GetCompanyPortalConfigurationsByParentIdInvoiceSetupsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/portalConfigurations/{parentId}/invoiceSetups/{id}"
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

type ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest struct {
	ctx context.Context
	ApiService *PortalConfigurationsInvoiceSetupsAPIService
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
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest) Conditions(conditions string) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest) ChildConditions(childConditions string) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest) CustomFieldConditions(customFieldConditions string) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest) OrderBy(orderBy string) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest) Fields(fields string) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest) Page(page int32) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest) PageSize(pageSize int32) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest) PageId(pageId int32) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest) ClientId(clientId string) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest) Execute() (*Count, *http.Response, error) {
	return r.ApiService.GetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountExecute(r)
}

/*
GetCompanyPortalConfigurationsByParentIdInvoiceSetupsCount Get Count of PortalConfigurationInvoiceSetup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId portalConfigurationId
 @return ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest
*/
func (a *PortalConfigurationsInvoiceSetupsAPIService) GetCompanyPortalConfigurationsByParentIdInvoiceSetupsCount(ctx context.Context, parentId int32) ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest {
	return ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return Count
func (a *PortalConfigurationsInvoiceSetupsAPIService) GetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountExecute(r ApiGetCompanyPortalConfigurationsByParentIdInvoiceSetupsCountRequest) (*Count, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Count
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalConfigurationsInvoiceSetupsAPIService.GetCompanyPortalConfigurationsByParentIdInvoiceSetupsCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/portalConfigurations/{parentId}/invoiceSetups/count"
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

type ApiPatchCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest struct {
	ctx context.Context
	ApiService *PortalConfigurationsInvoiceSetupsAPIService
	id int32
	parentId int32
	patchOperation *[]PatchOperation
	clientId *string
}

// List of PatchOperation
func (r ApiPatchCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest) PatchOperation(patchOperation []PatchOperation) ApiPatchCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest {
	r.patchOperation = &patchOperation
	return r
}

// 
func (r ApiPatchCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest) ClientId(clientId string) ApiPatchCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPatchCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest) Execute() (*PortalConfigurationInvoiceSetup, *http.Response, error) {
	return r.ApiService.PatchCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdExecute(r)
}

/*
PatchCompanyPortalConfigurationsByParentIdInvoiceSetupsById Patch PortalConfigurationInvoiceSetup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id invoiceSetupId
 @param parentId portalConfigurationId
 @return ApiPatchCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest
*/
func (a *PortalConfigurationsInvoiceSetupsAPIService) PatchCompanyPortalConfigurationsByParentIdInvoiceSetupsById(ctx context.Context, id int32, parentId int32) ApiPatchCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest {
	return ApiPatchCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return PortalConfigurationInvoiceSetup
func (a *PortalConfigurationsInvoiceSetupsAPIService) PatchCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdExecute(r ApiPatchCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest) (*PortalConfigurationInvoiceSetup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PortalConfigurationInvoiceSetup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalConfigurationsInvoiceSetupsAPIService.PatchCompanyPortalConfigurationsByParentIdInvoiceSetupsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/portalConfigurations/{parentId}/invoiceSetups/{id}"
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

type ApiPostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransactionRequest struct {
	ctx context.Context
	ApiService *PortalConfigurationsInvoiceSetupsAPIService
	id int32
	parentId int32
	portalConfigurationInvoiceSetup *PortalConfigurationInvoiceSetup
	clientId *string
}

// portalConfigurationInvoiceSetup
func (r ApiPostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransactionRequest) PortalConfigurationInvoiceSetup(portalConfigurationInvoiceSetup PortalConfigurationInvoiceSetup) ApiPostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransactionRequest {
	r.portalConfigurationInvoiceSetup = &portalConfigurationInvoiceSetup
	return r
}

// 
func (r ApiPostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransactionRequest) ClientId(clientId string) ApiPostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransactionRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransactionRequest) Execute() (*SuccessResponse, *http.Response, error) {
	return r.ApiService.PostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransactionExecute(r)
}

/*
PostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransaction Post SuccessResponse

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id invoiceSetupId
 @param parentId portalConfigurationId
 @return ApiPostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransactionRequest
*/
func (a *PortalConfigurationsInvoiceSetupsAPIService) PostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransaction(ctx context.Context, id int32, parentId int32) ApiPostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransactionRequest {
	return ApiPostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransactionRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return SuccessResponse
func (a *PortalConfigurationsInvoiceSetupsAPIService) PostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransactionExecute(r ApiPostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransactionRequest) (*SuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalConfigurationsInvoiceSetupsAPIService.PostCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdTestTransaction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/portalConfigurations/{parentId}/invoiceSetups/{id}/testTransaction"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.portalConfigurationInvoiceSetup == nil {
		return localVarReturnValue, nil, reportError("portalConfigurationInvoiceSetup is required and must be specified")
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
	localVarPostBody = r.portalConfigurationInvoiceSetup
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

type ApiPutCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest struct {
	ctx context.Context
	ApiService *PortalConfigurationsInvoiceSetupsAPIService
	id int32
	parentId int32
	portalConfigurationInvoiceSetup *PortalConfigurationInvoiceSetup
	clientId *string
}

// portalConfigurationInvoiceSetup
func (r ApiPutCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest) PortalConfigurationInvoiceSetup(portalConfigurationInvoiceSetup PortalConfigurationInvoiceSetup) ApiPutCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest {
	r.portalConfigurationInvoiceSetup = &portalConfigurationInvoiceSetup
	return r
}

// 
func (r ApiPutCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest) ClientId(clientId string) ApiPutCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPutCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest) Execute() (*PortalConfigurationInvoiceSetup, *http.Response, error) {
	return r.ApiService.PutCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdExecute(r)
}

/*
PutCompanyPortalConfigurationsByParentIdInvoiceSetupsById Put PortalConfigurationInvoiceSetup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id invoiceSetupId
 @param parentId portalConfigurationId
 @return ApiPutCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest
*/
func (a *PortalConfigurationsInvoiceSetupsAPIService) PutCompanyPortalConfigurationsByParentIdInvoiceSetupsById(ctx context.Context, id int32, parentId int32) ApiPutCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest {
	return ApiPutCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return PortalConfigurationInvoiceSetup
func (a *PortalConfigurationsInvoiceSetupsAPIService) PutCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdExecute(r ApiPutCompanyPortalConfigurationsByParentIdInvoiceSetupsByIdRequest) (*PortalConfigurationInvoiceSetup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PortalConfigurationInvoiceSetup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalConfigurationsInvoiceSetupsAPIService.PutCompanyPortalConfigurationsByParentIdInvoiceSetupsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/company/portalConfigurations/{parentId}/invoiceSetups/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"parentId"+"}", url.PathEscape(parameterValueToString(r.parentId, "parentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.portalConfigurationInvoiceSetup == nil {
		return localVarReturnValue, nil, reportError("portalConfigurationInvoiceSetup is required and must be specified")
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
	localVarPostBody = r.portalConfigurationInvoiceSetup
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
