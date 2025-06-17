# TicketInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTicketInfo

`func NewTicketInfo() *TicketInfo`

NewTicketInfo instantiates a new TicketInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketInfoWithDefaults

`func NewTicketInfoWithDefaults() *TicketInfo`

NewTicketInfoWithDefaults instantiates a new TicketInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TicketInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TicketInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TicketInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TicketInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSummary

`func (o *TicketInfo) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *TicketInfo) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *TicketInfo) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *TicketInfo) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetCompany

`func (o *TicketInfo) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *TicketInfo) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *TicketInfo) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *TicketInfo) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetInfo

`func (o *TicketInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TicketInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TicketInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TicketInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


