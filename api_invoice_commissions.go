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


// InvoiceCommissionsAPIService InvoiceCommissionsAPI service
type InvoiceCommissionsAPIService service

type ApiGetFinanceInvoicesByParentIdCommissionsRequest struct {
	ctx context.Context
	ApiService *InvoiceCommissionsAPIService
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
func (r ApiGetFinanceInvoicesByParentIdCommissionsRequest) Conditions(conditions string) ApiGetFinanceInvoicesByParentIdCommissionsRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetFinanceInvoicesByParentIdCommissionsRequest) ChildConditions(childConditions string) ApiGetFinanceInvoicesByParentIdCommissionsRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetFinanceInvoicesByParentIdCommissionsRequest) CustomFieldConditions(customFieldConditions string) ApiGetFinanceInvoicesByParentIdCommissionsRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetFinanceInvoicesByParentIdCommissionsRequest) OrderBy(orderBy string) ApiGetFinanceInvoicesByParentIdCommissionsRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetFinanceInvoicesByParentIdCommissionsRequest) Fields(fields string) ApiGetFinanceInvoicesByParentIdCommissionsRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetFinanceInvoicesByParentIdCommissionsRequest) Page(page int32) ApiGetFinanceInvoicesByParentIdCommissionsRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetFinanceInvoicesByParentIdCommissionsRequest) PageSize(pageSize int32) ApiGetFinanceInvoicesByParentIdCommissionsRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetFinanceInvoicesByParentIdCommissionsRequest) PageId(pageId int32) ApiGetFinanceInvoicesByParentIdCommissionsRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetFinanceInvoicesByParentIdCommissionsRequest) ClientId(clientId string) ApiGetFinanceInvoicesByParentIdCommissionsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetFinanceInvoicesByParentIdCommissionsRequest) Execute() ([]InvoiceCommission, *http.Response, error) {
	return r.ApiService.GetFinanceInvoicesByParentIdCommissionsExecute(r)
}

/*
GetFinanceInvoicesByParentIdCommissions Get List of InvoiceCommissions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId invoiceId
 @return ApiGetFinanceInvoicesByParentIdCommissionsRequest
*/
func (a *InvoiceCommissionsAPIService) GetFinanceInvoicesByParentIdCommissions(ctx context.Context, parentId int32) ApiGetFinanceInvoicesByParentIdCommissionsRequest {
	return ApiGetFinanceInvoicesByParentIdCommissionsRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return []InvoiceCommission
func (a *InvoiceCommissionsAPIService) GetFinanceInvoicesByParentIdCommissionsExecute(r ApiGetFinanceInvoicesByParentIdCommissionsRequest) ([]InvoiceCommission, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InvoiceCommission
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoiceCommissionsAPIService.GetFinanceInvoicesByParentIdCommissions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finance/invoices/{parentId}/commissions"
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

type ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest struct {
	ctx context.Context
	ApiService *InvoiceCommissionsAPIService
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
func (r ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest) Conditions(conditions string) ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest) ChildConditions(childConditions string) ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest) CustomFieldConditions(customFieldConditions string) ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest) OrderBy(orderBy string) ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest) Fields(fields string) ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest) Page(page int32) ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest) PageSize(pageSize int32) ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest) PageId(pageId int32) ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest) ClientId(clientId string) ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest) Execute() (*InvoiceCommission, *http.Response, error) {
	return r.ApiService.GetFinanceInvoicesByParentIdCommissionsByIdExecute(r)
}

/*
GetFinanceInvoicesByParentIdCommissionsById Get InvoiceCommissions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id paymentId
 @param parentId invoiceId
 @return ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest
*/
func (a *InvoiceCommissionsAPIService) GetFinanceInvoicesByParentIdCommissionsById(ctx context.Context, id int32, parentId int32) ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest {
	return ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return InvoiceCommission
func (a *InvoiceCommissionsAPIService) GetFinanceInvoicesByParentIdCommissionsByIdExecute(r ApiGetFinanceInvoicesByParentIdCommissionsByIdRequest) (*InvoiceCommission, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvoiceCommission
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoiceCommissionsAPIService.GetFinanceInvoicesByParentIdCommissionsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finance/invoices/{parentId}/commissions/{id}"
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

type ApiPatchFinanceInvoicesByParentIdCommissionsByIdRequest struct {
	ctx context.Context
	ApiService *InvoiceCommissionsAPIService
	id int32
	parentId int32
	patchOperation *[]PatchOperation
	clientId *string
}

// List of PatchOperation
func (r ApiPatchFinanceInvoicesByParentIdCommissionsByIdRequest) PatchOperation(patchOperation []PatchOperation) ApiPatchFinanceInvoicesByParentIdCommissionsByIdRequest {
	r.patchOperation = &patchOperation
	return r
}

// 
func (r ApiPatchFinanceInvoicesByParentIdCommissionsByIdRequest) ClientId(clientId string) ApiPatchFinanceInvoicesByParentIdCommissionsByIdRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPatchFinanceInvoicesByParentIdCommissionsByIdRequest) Execute() (*InvoiceCommission, *http.Response, error) {
	return r.ApiService.PatchFinanceInvoicesByParentIdCommissionsByIdExecute(r)
}

/*
PatchFinanceInvoicesByParentIdCommissionsById Patch InvoiceCommissions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id InvoiceCommissionId
 @param parentId invoiceId
 @return ApiPatchFinanceInvoicesByParentIdCommissionsByIdRequest
*/
func (a *InvoiceCommissionsAPIService) PatchFinanceInvoicesByParentIdCommissionsById(ctx context.Context, id int32, parentId int32) ApiPatchFinanceInvoicesByParentIdCommissionsByIdRequest {
	return ApiPatchFinanceInvoicesByParentIdCommissionsByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		parentId: parentId,
	}
}

// Execute executes the request
//  @return InvoiceCommission
func (a *InvoiceCommissionsAPIService) PatchFinanceInvoicesByParentIdCommissionsByIdExecute(r ApiPatchFinanceInvoicesByParentIdCommissionsByIdRequest) (*InvoiceCommission, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvoiceCommission
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoiceCommissionsAPIService.PatchFinanceInvoicesByParentIdCommissionsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finance/invoices/{parentId}/commissions/{id}"
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

type ApiPostFinanceInvoicesByParentIdCommissionsRecalculateRequest struct {
	ctx context.Context
	ApiService *InvoiceCommissionsAPIService
	parentId int32
	clientId *string
}

// 
func (r ApiPostFinanceInvoicesByParentIdCommissionsRecalculateRequest) ClientId(clientId string) ApiPostFinanceInvoicesByParentIdCommissionsRecalculateRequest {
	r.clientId = &clientId
	return r
}

func (r ApiPostFinanceInvoicesByParentIdCommissionsRecalculateRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostFinanceInvoicesByParentIdCommissionsRecalculateExecute(r)
}

/*
PostFinanceInvoicesByParentIdCommissionsRecalculate Recalculate InvoiceCommissions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentId invoiceId
 @return ApiPostFinanceInvoicesByParentIdCommissionsRecalculateRequest
*/
func (a *InvoiceCommissionsAPIService) PostFinanceInvoicesByParentIdCommissionsRecalculate(ctx context.Context, parentId int32) ApiPostFinanceInvoicesByParentIdCommissionsRecalculateRequest {
	return ApiPostFinanceInvoicesByParentIdCommissionsRecalculateRequest{
		ApiService: a,
		ctx: ctx,
		parentId: parentId,
	}
}

// Execute executes the request
func (a *InvoiceCommissionsAPIService) PostFinanceInvoicesByParentIdCommissionsRecalculateExecute(r ApiPostFinanceInvoicesByParentIdCommissionsRecalculateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoiceCommissionsAPIService.PostFinanceInvoicesByParentIdCommissionsRecalculate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/finance/invoices/{parentId}/commissions/recalculate"
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
