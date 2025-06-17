# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**SrNotify** | **NullableString** |  | 
**ScheduleSpan** | **string** |  | 
**HideDelimiterFlag** | Pointer to **NullableBool** |  | [optional] 
**AllowCCFlag** | Pointer to **NullableBool** |  | [optional] 
**AllowTOFlag** | Pointer to **NullableBool** |  | [optional] 
**HeaderColor** | Pointer to **string** |  Max length: 50; | [optional] 
**MemberColor** | Pointer to **string** |  Max length: 50; | [optional] 
**ContactColor** | Pointer to **string** |  Max length: 50; | [optional] 
**UnknownColor** | Pointer to **string** |  Max length: 50; | [optional] 
**CalendarSetup** | Pointer to [**CalendarSetupReference**](CalendarSetupReference.md) |  | [optional] 
**HeaderColorDisableFlag** | Pointer to **NullableBool** |  | [optional] 
**MemberColorDisableFlag** | Pointer to **NullableBool** |  | [optional] 
**ContactColorDisableFlag** | Pointer to **NullableBool** |  | [optional] 
**UnknownColorDisableFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewService

`func NewService(srNotify NullableString, scheduleSpan string, ) *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Service) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Service) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Service) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSrNotify

`func (o *Service) GetSrNotify() string`

GetSrNotify returns the SrNotify field if non-nil, zero value otherwise.

### GetSrNotifyOk

`func (o *Service) GetSrNotifyOk() (*string, bool)`

GetSrNotifyOk returns a tuple with the SrNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrNotify

`func (o *Service) SetSrNotify(v string)`

SetSrNotify sets SrNotify field to given value.


### SetSrNotifyNil

`func (o *Service) SetSrNotifyNil(b bool)`

 SetSrNotifyNil sets the value for SrNotify to be an explicit nil

### UnsetSrNotify
`func (o *Service) UnsetSrNotify()`

UnsetSrNotify ensures that no value is present for SrNotify, not even an explicit nil
### GetScheduleSpan

`func (o *Service) GetScheduleSpan() string`

GetScheduleSpan returns the ScheduleSpan field if non-nil, zero value otherwise.

### GetScheduleSpanOk

`func (o *Service) GetScheduleSpanOk() (*string, bool)`

GetScheduleSpanOk returns a tuple with the ScheduleSpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleSpan

`func (o *Service) SetScheduleSpan(v string)`

SetScheduleSpan sets ScheduleSpan field to given value.


### GetHideDelimiterFlag

`func (o *Service) GetHideDelimiterFlag() bool`

GetHideDelimiterFlag returns the HideDelimiterFlag field if non-nil, zero value otherwise.

### GetHideDelimiterFlagOk

`func (o *Service) GetHideDelimiterFlagOk() (*bool, bool)`

GetHideDelimiterFlagOk returns a tuple with the HideDelimiterFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideDelimiterFlag

`func (o *Service) SetHideDelimiterFlag(v bool)`

SetHideDelimiterFlag sets HideDelimiterFlag field to given value.

### HasHideDelimiterFlag

`func (o *Service) HasHideDelimiterFlag() bool`

HasHideDelimiterFlag returns a boolean if a field has been set.

### SetHideDelimiterFlagNil

`func (o *Service) SetHideDelimiterFlagNil(b bool)`

 SetHideDelimiterFlagNil sets the value for HideDelimiterFlag to be an explicit nil

### UnsetHideDelimiterFlag
`func (o *Service) UnsetHideDelimiterFlag()`

UnsetHideDelimiterFlag ensures that no value is present for HideDelimiterFlag, not even an explicit nil
### GetAllowCCFlag

`func (o *Service) GetAllowCCFlag() bool`

GetAllowCCFlag returns the AllowCCFlag field if non-nil, zero value otherwise.

### GetAllowCCFlagOk

`func (o *Service) GetAllowCCFlagOk() (*bool, bool)`

GetAllowCCFlagOk returns a tuple with the AllowCCFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCCFlag

`func (o *Service) SetAllowCCFlag(v bool)`

SetAllowCCFlag sets AllowCCFlag field to given value.

### HasAllowCCFlag

`func (o *Service) HasAllowCCFlag() bool`

HasAllowCCFlag returns a boolean if a field has been set.

### SetAllowCCFlagNil

`func (o *Service) SetAllowCCFlagNil(b bool)`

 SetAllowCCFlagNil sets the value for AllowCCFlag to be an explicit nil

### UnsetAllowCCFlag
`func (o *Service) UnsetAllowCCFlag()`

UnsetAllowCCFlag ensures that no value is present for AllowCCFlag, not even an explicit nil
### GetAllowTOFlag

`func (o *Service) GetAllowTOFlag() bool`

GetAllowTOFlag returns the AllowTOFlag field if non-nil, zero value otherwise.

### GetAllowTOFlagOk

`func (o *Service) GetAllowTOFlagOk() (*bool, bool)`

GetAllowTOFlagOk returns a tuple with the AllowTOFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTOFlag

`func (o *Service) SetAllowTOFlag(v bool)`

SetAllowTOFlag sets AllowTOFlag field to given value.

### HasAllowTOFlag

`func (o *Service) HasAllowTOFlag() bool`

HasAllowTOFlag returns a boolean if a field has been set.

### SetAllowTOFlagNil

`func (o *Service) SetAllowTOFlagNil(b bool)`

 SetAllowTOFlagNil sets the value for AllowTOFlag to be an explicit nil

### UnsetAllowTOFlag
`func (o *Service) UnsetAllowTOFlag()`

UnsetAllowTOFlag ensures that no value is present for AllowTOFlag, not even an explicit nil
### GetHeaderColor

`func (o *Service) GetHeaderColor() string`

GetHeaderColor returns the HeaderColor field if non-nil, zero value otherwise.

### GetHeaderColorOk

`func (o *Service) GetHeaderColorOk() (*string, bool)`

GetHeaderColorOk returns a tuple with the HeaderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderColor

`func (o *Service) SetHeaderColor(v string)`

SetHeaderColor sets HeaderColor field to given value.

### HasHeaderColor

`func (o *Service) HasHeaderColor() bool`

HasHeaderColor returns a boolean if a field has been set.

### GetMemberColor

`func (o *Service) GetMemberColor() string`

GetMemberColor returns the MemberColor field if non-nil, zero value otherwise.

### GetMemberColorOk

`func (o *Service) GetMemberColorOk() (*string, bool)`

GetMemberColorOk returns a tuple with the MemberColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberColor

`func (o *Service) SetMemberColor(v string)`

SetMemberColor sets MemberColor field to given value.

### HasMemberColor

`func (o *Service) HasMemberColor() bool`

HasMemberColor returns a boolean if a field has been set.

### GetContactColor

`func (o *Service) GetContactColor() string`

GetContactColor returns the ContactColor field if non-nil, zero value otherwise.

### GetContactColorOk

`func (o *Service) GetContactColorOk() (*string, bool)`

GetContactColorOk returns a tuple with the ContactColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactColor

`func (o *Service) SetContactColor(v string)`

SetContactColor sets ContactColor field to given value.

### HasContactColor

`func (o *Service) HasContactColor() bool`

HasContactColor returns a boolean if a field has been set.

### GetUnknownColor

`func (o *Service) GetUnknownColor() string`

GetUnknownColor returns the UnknownColor field if non-nil, zero value otherwise.

### GetUnknownColorOk

`func (o *Service) GetUnknownColorOk() (*string, bool)`

GetUnknownColorOk returns a tuple with the UnknownColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownColor

`func (o *Service) SetUnknownColor(v string)`

SetUnknownColor sets UnknownColor field to given value.

### HasUnknownColor

`func (o *Service) HasUnknownColor() bool`

HasUnknownColor returns a boolean if a field has been set.

### GetCalendarSetup

`func (o *Service) GetCalendarSetup() CalendarSetupReference`

GetCalendarSetup returns the CalendarSetup field if non-nil, zero value otherwise.

### GetCalendarSetupOk

`func (o *Service) GetCalendarSetupOk() (*CalendarSetupReference, bool)`

GetCalendarSetupOk returns a tuple with the CalendarSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarSetup

`func (o *Service) SetCalendarSetup(v CalendarSetupReference)`

SetCalendarSetup sets CalendarSetup field to given value.

### HasCalendarSetup

`func (o *Service) HasCalendarSetup() bool`

HasCalendarSetup returns a boolean if a field has been set.

### GetHeaderColorDisableFlag

`func (o *Service) GetHeaderColorDisableFlag() bool`

GetHeaderColorDisableFlag returns the HeaderColorDisableFlag field if non-nil, zero value otherwise.

### GetHeaderColorDisableFlagOk

`func (o *Service) GetHeaderColorDisableFlagOk() (*bool, bool)`

GetHeaderColorDisableFlagOk returns a tuple with the HeaderColorDisableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderColorDisableFlag

`func (o *Service) SetHeaderColorDisableFlag(v bool)`

SetHeaderColorDisableFlag sets HeaderColorDisableFlag field to given value.

### HasHeaderColorDisableFlag

`func (o *Service) HasHeaderColorDisableFlag() bool`

HasHeaderColorDisableFlag returns a boolean if a field has been set.

### SetHeaderColorDisableFlagNil

`func (o *Service) SetHeaderColorDisableFlagNil(b bool)`

 SetHeaderColorDisableFlagNil sets the value for HeaderColorDisableFlag to be an explicit nil

### UnsetHeaderColorDisableFlag
`func (o *Service) UnsetHeaderColorDisableFlag()`

UnsetHeaderColorDisableFlag ensures that no value is present for HeaderColorDisableFlag, not even an explicit nil
### GetMemberColorDisableFlag

`func (o *Service) GetMemberColorDisableFlag() bool`

GetMemberColorDisableFlag returns the MemberColorDisableFlag field if non-nil, zero value otherwise.

### GetMemberColorDisableFlagOk

`func (o *Service) GetMemberColorDisableFlagOk() (*bool, bool)`

GetMemberColorDisableFlagOk returns a tuple with the MemberColorDisableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberColorDisableFlag

`func (o *Service) SetMemberColorDisableFlag(v bool)`

SetMemberColorDisableFlag sets MemberColorDisableFlag field to given value.

### HasMemberColorDisableFlag

`func (o *Service) HasMemberColorDisableFlag() bool`

HasMemberColorDisableFlag returns a boolean if a field has been set.

### SetMemberColorDisableFlagNil

`func (o *Service) SetMemberColorDisableFlagNil(b bool)`

 SetMemberColorDisableFlagNil sets the value for MemberColorDisableFlag to be an explicit nil

### UnsetMemberColorDisableFlag
`func (o *Service) UnsetMemberColorDisableFlag()`

UnsetMemberColorDisableFlag ensures that no value is present for MemberColorDisableFlag, not even an explicit nil
### GetContactColorDisableFlag

`func (o *Service) GetContactColorDisableFlag() bool`

GetContactColorDisableFlag returns the ContactColorDisableFlag field if non-nil, zero value otherwise.

### GetContactColorDisableFlagOk

`func (o *Service) GetContactColorDisableFlagOk() (*bool, bool)`

GetContactColorDisableFlagOk returns a tuple with the ContactColorDisableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactColorDisableFlag

`func (o *Service) SetContactColorDisableFlag(v bool)`

SetContactColorDisableFlag sets ContactColorDisableFlag field to given value.

### HasContactColorDisableFlag

`func (o *Service) HasContactColorDisableFlag() bool`

HasContactColorDisableFlag returns a boolean if a field has been set.

### SetContactColorDisableFlagNil

`func (o *Service) SetContactColorDisableFlagNil(b bool)`

 SetContactColorDisableFlagNil sets the value for ContactColorDisableFlag to be an explicit nil

### UnsetContactColorDisableFlag
`func (o *Service) UnsetContactColorDisableFlag()`

UnsetContactColorDisableFlag ensures that no value is present for ContactColorDisableFlag, not even an explicit nil
### GetUnknownColorDisableFlag

`func (o *Service) GetUnknownColorDisableFlag() bool`

GetUnknownColorDisableFlag returns the UnknownColorDisableFlag field if non-nil, zero value otherwise.

### GetUnknownColorDisableFlagOk

`func (o *Service) GetUnknownColorDisableFlagOk() (*bool, bool)`

GetUnknownColorDisableFlagOk returns a tuple with the UnknownColorDisableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownColorDisableFlag

`func (o *Service) SetUnknownColorDisableFlag(v bool)`

SetUnknownColorDisableFlag sets UnknownColorDisableFlag field to given value.

### HasUnknownColorDisableFlag

`func (o *Service) HasUnknownColorDisableFlag() bool`

HasUnknownColorDisableFlag returns a boolean if a field has been set.

### SetUnknownColorDisableFlagNil

`func (o *Service) SetUnknownColorDisableFlagNil(b bool)`

 SetUnknownColorDisableFlagNil sets the value for UnknownColorDisableFlag to be an explicit nil

### UnsetUnknownColorDisableFlag
`func (o *Service) UnsetUnknownColorDisableFlag()`

UnsetUnknownColorDisableFlag ensures that no value is present for UnknownColorDisableFlag, not even an explicit nil
### GetInfo

`func (o *Service) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Service) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Service) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Service) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


