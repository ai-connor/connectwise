# SurveyResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**TicketId** | **NullableInt32** |  | 
**EmailAddress** | Pointer to **string** |  | [optional] 
**FooterResponse** | Pointer to **string** |  | [optional] 
**ContactMeFlag** | Pointer to **NullableBool** |  | [optional] 
**Contact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**Results** | Pointer to [**[]SurveyResultDetail**](SurveyResultDetail.md) |  | [optional] 
**TotalPoints** | Pointer to **NullableInt32** |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**SurveyId** | Pointer to **NullableInt32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSurveyResult

`func NewSurveyResult(ticketId NullableInt32, ) *SurveyResult`

NewSurveyResult instantiates a new SurveyResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSurveyResultWithDefaults

`func NewSurveyResultWithDefaults() *SurveyResult`

NewSurveyResultWithDefaults instantiates a new SurveyResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SurveyResult) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SurveyResult) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SurveyResult) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SurveyResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTicketId

`func (o *SurveyResult) GetTicketId() int32`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *SurveyResult) GetTicketIdOk() (*int32, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *SurveyResult) SetTicketId(v int32)`

SetTicketId sets TicketId field to given value.


### SetTicketIdNil

`func (o *SurveyResult) SetTicketIdNil(b bool)`

 SetTicketIdNil sets the value for TicketId to be an explicit nil

### UnsetTicketId
`func (o *SurveyResult) UnsetTicketId()`

UnsetTicketId ensures that no value is present for TicketId, not even an explicit nil
### GetEmailAddress

`func (o *SurveyResult) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *SurveyResult) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *SurveyResult) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *SurveyResult) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetFooterResponse

`func (o *SurveyResult) GetFooterResponse() string`

GetFooterResponse returns the FooterResponse field if non-nil, zero value otherwise.

### GetFooterResponseOk

`func (o *SurveyResult) GetFooterResponseOk() (*string, bool)`

GetFooterResponseOk returns a tuple with the FooterResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterResponse

`func (o *SurveyResult) SetFooterResponse(v string)`

SetFooterResponse sets FooterResponse field to given value.

### HasFooterResponse

`func (o *SurveyResult) HasFooterResponse() bool`

HasFooterResponse returns a boolean if a field has been set.

### GetContactMeFlag

`func (o *SurveyResult) GetContactMeFlag() bool`

GetContactMeFlag returns the ContactMeFlag field if non-nil, zero value otherwise.

### GetContactMeFlagOk

`func (o *SurveyResult) GetContactMeFlagOk() (*bool, bool)`

GetContactMeFlagOk returns a tuple with the ContactMeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactMeFlag

`func (o *SurveyResult) SetContactMeFlag(v bool)`

SetContactMeFlag sets ContactMeFlag field to given value.

### HasContactMeFlag

`func (o *SurveyResult) HasContactMeFlag() bool`

HasContactMeFlag returns a boolean if a field has been set.

### SetContactMeFlagNil

`func (o *SurveyResult) SetContactMeFlagNil(b bool)`

 SetContactMeFlagNil sets the value for ContactMeFlag to be an explicit nil

### UnsetContactMeFlag
`func (o *SurveyResult) UnsetContactMeFlag()`

UnsetContactMeFlag ensures that no value is present for ContactMeFlag, not even an explicit nil
### GetContact

`func (o *SurveyResult) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *SurveyResult) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *SurveyResult) SetContact(v ContactReference)`

SetContact sets Contact field to given value.

### HasContact

`func (o *SurveyResult) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetResults

`func (o *SurveyResult) GetResults() []SurveyResultDetail`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SurveyResult) GetResultsOk() (*[]SurveyResultDetail, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SurveyResult) SetResults(v []SurveyResultDetail)`

SetResults sets Results field to given value.

### HasResults

`func (o *SurveyResult) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotalPoints

`func (o *SurveyResult) GetTotalPoints() int32`

GetTotalPoints returns the TotalPoints field if non-nil, zero value otherwise.

### GetTotalPointsOk

`func (o *SurveyResult) GetTotalPointsOk() (*int32, bool)`

GetTotalPointsOk returns a tuple with the TotalPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPoints

`func (o *SurveyResult) SetTotalPoints(v int32)`

SetTotalPoints sets TotalPoints field to given value.

### HasTotalPoints

`func (o *SurveyResult) HasTotalPoints() bool`

HasTotalPoints returns a boolean if a field has been set.

### SetTotalPointsNil

`func (o *SurveyResult) SetTotalPointsNil(b bool)`

 SetTotalPointsNil sets the value for TotalPoints to be an explicit nil

### UnsetTotalPoints
`func (o *SurveyResult) UnsetTotalPoints()`

UnsetTotalPoints ensures that no value is present for TotalPoints, not even an explicit nil
### GetCompany

`func (o *SurveyResult) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *SurveyResult) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *SurveyResult) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *SurveyResult) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetSurveyId

`func (o *SurveyResult) GetSurveyId() int32`

GetSurveyId returns the SurveyId field if non-nil, zero value otherwise.

### GetSurveyIdOk

`func (o *SurveyResult) GetSurveyIdOk() (*int32, bool)`

GetSurveyIdOk returns a tuple with the SurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyId

`func (o *SurveyResult) SetSurveyId(v int32)`

SetSurveyId sets SurveyId field to given value.

### HasSurveyId

`func (o *SurveyResult) HasSurveyId() bool`

HasSurveyId returns a boolean if a field has been set.

### SetSurveyIdNil

`func (o *SurveyResult) SetSurveyIdNil(b bool)`

 SetSurveyIdNil sets the value for SurveyId to be an explicit nil

### UnsetSurveyId
`func (o *SurveyResult) UnsetSurveyId()`

UnsetSurveyId ensures that no value is present for SurveyId, not even an explicit nil
### GetInfo

`func (o *SurveyResult) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SurveyResult) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SurveyResult) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SurveyResult) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


