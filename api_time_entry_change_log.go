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


// TimeEntryChangeLogAPIService TimeEntryChangeLogAPI service
type TimeEntryChangeLogAPIService service

type ApiGetTimeChangelogsRequest struct {
	ctx context.Context
	ApiService *TimeEntryChangeLogAPIService
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
func (r ApiGetTimeChangelogsRequest) Conditions(conditions string) ApiGetTimeChangelogsRequest {
	r.conditions = &conditions
	return r
}

// 
func (r ApiGetTimeChangelogsRequest) ChildConditions(childConditions string) ApiGetTimeChangelogsRequest {
	r.childConditions = &childConditions
	return r
}

// 
func (r ApiGetTimeChangelogsRequest) CustomFieldConditions(customFieldConditions string) ApiGetTimeChangelogsRequest {
	r.customFieldConditions = &customFieldConditions
	return r
}

// 
func (r ApiGetTimeChangelogsRequest) OrderBy(orderBy string) ApiGetTimeChangelogsRequest {
	r.orderBy = &orderBy
	return r
}

// 
func (r ApiGetTimeChangelogsRequest) Fields(fields string) ApiGetTimeChangelogsRequest {
	r.fields = &fields
	return r
}

// 
func (r ApiGetTimeChangelogsRequest) Page(page int32) ApiGetTimeChangelogsRequest {
	r.page = &page
	return r
}

// 
func (r ApiGetTimeChangelogsRequest) PageSize(pageSize int32) ApiGetTimeChangelogsRequest {
	r.pageSize = &pageSize
	return r
}

// 
func (r ApiGetTimeChangelogsRequest) PageId(pageId int32) ApiGetTimeChangelogsRequest {
	r.pageId = &pageId
	return r
}

// 
func (r ApiGetTimeChangelogsRequest) ClientId(clientId string) ApiGetTimeChangelogsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetTimeChangelogsRequest) Execute() ([]TimeEntryChangeLog, *http.Response, error) {
	return r.ApiService.GetTimeChangelogsExecute(r)
}

/*
GetTimeChangelogs Get List of Time Entry Change Logs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetTimeChangelogsRequest
*/
func (a *TimeEntryChangeLogAPIService) GetTimeChangelogs(ctx context.Context) ApiGetTimeChangelogsRequest {
	return ApiGetTimeChangelogsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []TimeEntryChangeLog
func (a *TimeEntryChangeLogAPIService) GetTimeChangelogsExecute(r ApiGetTimeChangelogsRequest) ([]TimeEntryChangeLog, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []TimeEntryChangeLog
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TimeEntryChangeLogAPIService.GetTimeChangelogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/time/changelogs"

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
