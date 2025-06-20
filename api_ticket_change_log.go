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
)


// TicketChangeLogAPIService TicketChangeLogAPI service
type TicketChangeLogAPIService service

type ApiDeleteServiceTicketsChangelogsRequest struct {
	ctx context.Context
	ApiService *TicketChangeLogAPIService
	clientId *string
}

// 
func (r ApiDeleteServiceTicketsChangelogsRequest) ClientId(clientId string) ApiDeleteServiceTicketsChangelogsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiDeleteServiceTicketsChangelogsRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteServiceTicketsChangelogsExecute(r)
}

/*
DeleteServiceTicketsChangelogs Delete Ticket Change Logs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteServiceTicketsChangelogsRequest
*/
func (a *TicketChangeLogAPIService) DeleteServiceTicketsChangelogs(ctx context.Context) ApiDeleteServiceTicketsChangelogsRequest {
	return ApiDeleteServiceTicketsChangelogsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *TicketChangeLogAPIService) DeleteServiceTicketsChangelogsExecute(r ApiDeleteServiceTicketsChangelogsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TicketChangeLogAPIService.DeleteServiceTicketsChangelogs")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/tickets/changelogs"

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

type ApiGetServiceTicketsChangelogsRequest struct {
	ctx context.Context
	ApiService *TicketChangeLogAPIService
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
func (r ApiGetServiceTicketsChangelogsRequest) Conditions(conditions string) ApiGetServiceTicketsChangelogsRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetServiceTicketsChangelogsRequest) ChildConditions(childConditions string) ApiGetServiceTicketsChangelogsRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetServiceTicketsChangelogsRequest) CustomFieldConditions(customFieldConditions string) ApiGetServiceTicketsChangelogsRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetServiceTicketsChangelogsRequest) OrderBy(orderBy string) ApiGetServiceTicketsChangelogsRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetServiceTicketsChangelogsRequest) Fields(fields string) ApiGetServiceTicketsChangelogsRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetServiceTicketsChangelogsRequest) Page(page int32) ApiGetServiceTicketsChangelogsRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetServiceTicketsChangelogsRequest) PageSize(pageSize int32) ApiGetServiceTicketsChangelogsRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetServiceTicketsChangelogsRequest) PageId(pageId int32) ApiGetServiceTicketsChangelogsRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetServiceTicketsChangelogsRequest) ClientId(clientId string) ApiGetServiceTicketsChangelogsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetServiceTicketsChangelogsRequest) Execute() ([]TicketChangeLog, *http.Response, error) {
	return r.ApiService.GetServiceTicketsChangelogsExecute(r)
}

/*
GetServiceTicketsChangelogs Get List of Ticket Change Log

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetServiceTicketsChangelogsRequest
*/
func (a *TicketChangeLogAPIService) GetServiceTicketsChangelogs(ctx context.Context) ApiGetServiceTicketsChangelogsRequest {
	return ApiGetServiceTicketsChangelogsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []TicketChangeLog
func (a *TicketChangeLogAPIService) GetServiceTicketsChangelogsExecute(r ApiGetServiceTicketsChangelogsRequest) ([]TicketChangeLog, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []TicketChangeLog
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TicketChangeLogAPIService.GetServiceTicketsChangelogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/tickets/changelogs"

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
