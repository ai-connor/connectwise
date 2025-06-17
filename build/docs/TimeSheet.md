# TimeSheet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Year** | Pointer to **NullableInt32** |  | [optional] 
**Period** | Pointer to **NullableInt32** |  | [optional] 
**DateStart** | Pointer to **string** |  | [optional] 
**DateEnd** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Hours** | Pointer to **NullableFloat64** |  | [optional] 
**Deadline** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTimeSheet

`func NewTimeSheet() *TimeSheet`

NewTimeSheet instantiates a new TimeSheet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSheetWithDefaults

`func NewTimeSheetWithDefaults() *TimeSheet`

NewTimeSheetWithDefaults instantiates a new TimeSheet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimeSheet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimeSheet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimeSheet) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TimeSheet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMember

`func (o *TimeSheet) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *TimeSheet) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *TimeSheet) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *TimeSheet) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetYear

`func (o *TimeSheet) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *TimeSheet) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *TimeSheet) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *TimeSheet) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *TimeSheet) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *TimeSheet) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetPeriod

`func (o *TimeSheet) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *TimeSheet) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *TimeSheet) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *TimeSheet) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### SetPeriodNil

`func (o *TimeSheet) SetPeriodNil(b bool)`

 SetPeriodNil sets the value for Period to be an explicit nil

### UnsetPeriod
`func (o *TimeSheet) UnsetPeriod()`

UnsetPeriod ensures that no value is present for Period, not even an explicit nil
### GetDateStart

`func (o *TimeSheet) GetDateStart() string`

GetDateStart returns the DateStart field if non-nil, zero value otherwise.

### GetDateStartOk

`func (o *TimeSheet) GetDateStartOk() (*string, bool)`

GetDateStartOk returns a tuple with the DateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStart

`func (o *TimeSheet) SetDateStart(v string)`

SetDateStart sets DateStart field to given value.

### HasDateStart

`func (o *TimeSheet) HasDateStart() bool`

HasDateStart returns a boolean if a field has been set.

### GetDateEnd

`func (o *TimeSheet) GetDateEnd() string`

GetDateEnd returns the DateEnd field if non-nil, zero value otherwise.

### GetDateEndOk

`func (o *TimeSheet) GetDateEndOk() (*string, bool)`

GetDateEndOk returns a tuple with the DateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnd

`func (o *TimeSheet) SetDateEnd(v string)`

SetDateEnd sets DateEnd field to given value.

### HasDateEnd

`func (o *TimeSheet) HasDateEnd() bool`

HasDateEnd returns a boolean if a field has been set.

### GetStatus

`func (o *TimeSheet) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TimeSheet) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TimeSheet) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TimeSheet) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TimeSheet) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TimeSheet) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetHours

`func (o *TimeSheet) GetHours() float64`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *TimeSheet) GetHoursOk() (*float64, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *TimeSheet) SetHours(v float64)`

SetHours sets Hours field to given value.

### HasHours

`func (o *TimeSheet) HasHours() bool`

HasHours returns a boolean if a field has been set.

### SetHoursNil

`func (o *TimeSheet) SetHoursNil(b bool)`

 SetHoursNil sets the value for Hours to be an explicit nil

### UnsetHours
`func (o *TimeSheet) UnsetHours()`

UnsetHours ensures that no value is present for Hours, not even an explicit nil
### GetDeadline

`func (o *TimeSheet) GetDeadline() string`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *TimeSheet) GetDeadlineOk() (*string, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *TimeSheet) SetDeadline(v string)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *TimeSheet) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetInfo

`func (o *TimeSheet) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TimeSheet) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TimeSheet) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TimeSheet) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


