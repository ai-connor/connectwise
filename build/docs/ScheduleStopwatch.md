# ScheduleStopwatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to **map[string]string** |  | [optional] 
**Agreement** | Pointer to [**AgreementReference**](AgreementReference.md) |  | [optional] 
**BillableOption** | Pointer to **NullableString** |  | [optional] 
**BusinessUnitId** | Pointer to **NullableInt32** |  | [optional] 
**DateEntered** | Pointer to **time.Time** |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**InternalNotes** | Pointer to **string** |  | [optional] 
**LocationId** | Pointer to **NullableInt32** |  | [optional] 
**Member** | [**MemberReference**](MemberReference.md) |  | 
**MobileGuid** | Pointer to **NullableString** |  | [optional] 
**Notes** | Pointer to **string** |  Max length: 4000; | [optional] 
**ScheduleId** | **NullableInt32** |  | 
**ScheduleMobileGuid** | Pointer to **NullableString** |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**Status** | **NullableString** |  | 
**TotalPauseTime** | Pointer to **NullableInt32** |  | [optional] 
**WorkRole** | Pointer to [**WorkRoleReference**](WorkRoleReference.md) |  | [optional] 
**WorkType** | Pointer to [**WorkTypeReference**](WorkTypeReference.md) |  | [optional] 

## Methods

### NewScheduleStopwatch

`func NewScheduleStopwatch(member MemberReference, scheduleId NullableInt32, status NullableString, ) *ScheduleStopwatch`

NewScheduleStopwatch instantiates a new ScheduleStopwatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleStopwatchWithDefaults

`func NewScheduleStopwatchWithDefaults() *ScheduleStopwatch`

NewScheduleStopwatchWithDefaults instantiates a new ScheduleStopwatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *ScheduleStopwatch) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ScheduleStopwatch) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ScheduleStopwatch) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ScheduleStopwatch) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetAgreement

`func (o *ScheduleStopwatch) GetAgreement() AgreementReference`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *ScheduleStopwatch) GetAgreementOk() (*AgreementReference, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *ScheduleStopwatch) SetAgreement(v AgreementReference)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *ScheduleStopwatch) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetBillableOption

`func (o *ScheduleStopwatch) GetBillableOption() string`

GetBillableOption returns the BillableOption field if non-nil, zero value otherwise.

### GetBillableOptionOk

`func (o *ScheduleStopwatch) GetBillableOptionOk() (*string, bool)`

GetBillableOptionOk returns a tuple with the BillableOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableOption

`func (o *ScheduleStopwatch) SetBillableOption(v string)`

SetBillableOption sets BillableOption field to given value.

### HasBillableOption

`func (o *ScheduleStopwatch) HasBillableOption() bool`

HasBillableOption returns a boolean if a field has been set.

### SetBillableOptionNil

`func (o *ScheduleStopwatch) SetBillableOptionNil(b bool)`

 SetBillableOptionNil sets the value for BillableOption to be an explicit nil

### UnsetBillableOption
`func (o *ScheduleStopwatch) UnsetBillableOption()`

UnsetBillableOption ensures that no value is present for BillableOption, not even an explicit nil
### GetBusinessUnitId

`func (o *ScheduleStopwatch) GetBusinessUnitId() int32`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *ScheduleStopwatch) GetBusinessUnitIdOk() (*int32, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *ScheduleStopwatch) SetBusinessUnitId(v int32)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *ScheduleStopwatch) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### SetBusinessUnitIdNil

`func (o *ScheduleStopwatch) SetBusinessUnitIdNil(b bool)`

 SetBusinessUnitIdNil sets the value for BusinessUnitId to be an explicit nil

### UnsetBusinessUnitId
`func (o *ScheduleStopwatch) UnsetBusinessUnitId()`

UnsetBusinessUnitId ensures that no value is present for BusinessUnitId, not even an explicit nil
### GetDateEntered

`func (o *ScheduleStopwatch) GetDateEntered() time.Time`

GetDateEntered returns the DateEntered field if non-nil, zero value otherwise.

### GetDateEnteredOk

`func (o *ScheduleStopwatch) GetDateEnteredOk() (*time.Time, bool)`

GetDateEnteredOk returns a tuple with the DateEntered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEntered

`func (o *ScheduleStopwatch) SetDateEntered(v time.Time)`

SetDateEntered sets DateEntered field to given value.

### HasDateEntered

`func (o *ScheduleStopwatch) HasDateEntered() bool`

HasDateEntered returns a boolean if a field has been set.

### GetEndTime

`func (o *ScheduleStopwatch) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ScheduleStopwatch) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ScheduleStopwatch) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ScheduleStopwatch) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetId

`func (o *ScheduleStopwatch) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleStopwatch) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleStopwatch) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ScheduleStopwatch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalNotes

`func (o *ScheduleStopwatch) GetInternalNotes() string`

GetInternalNotes returns the InternalNotes field if non-nil, zero value otherwise.

### GetInternalNotesOk

`func (o *ScheduleStopwatch) GetInternalNotesOk() (*string, bool)`

GetInternalNotesOk returns a tuple with the InternalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNotes

`func (o *ScheduleStopwatch) SetInternalNotes(v string)`

SetInternalNotes sets InternalNotes field to given value.

### HasInternalNotes

`func (o *ScheduleStopwatch) HasInternalNotes() bool`

HasInternalNotes returns a boolean if a field has been set.

### GetLocationId

`func (o *ScheduleStopwatch) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *ScheduleStopwatch) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *ScheduleStopwatch) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *ScheduleStopwatch) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *ScheduleStopwatch) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *ScheduleStopwatch) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetMember

`func (o *ScheduleStopwatch) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ScheduleStopwatch) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ScheduleStopwatch) SetMember(v MemberReference)`

SetMember sets Member field to given value.


### GetMobileGuid

`func (o *ScheduleStopwatch) GetMobileGuid() string`

GetMobileGuid returns the MobileGuid field if non-nil, zero value otherwise.

### GetMobileGuidOk

`func (o *ScheduleStopwatch) GetMobileGuidOk() (*string, bool)`

GetMobileGuidOk returns a tuple with the MobileGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileGuid

`func (o *ScheduleStopwatch) SetMobileGuid(v string)`

SetMobileGuid sets MobileGuid field to given value.

### HasMobileGuid

`func (o *ScheduleStopwatch) HasMobileGuid() bool`

HasMobileGuid returns a boolean if a field has been set.

### SetMobileGuidNil

`func (o *ScheduleStopwatch) SetMobileGuidNil(b bool)`

 SetMobileGuidNil sets the value for MobileGuid to be an explicit nil

### UnsetMobileGuid
`func (o *ScheduleStopwatch) UnsetMobileGuid()`

UnsetMobileGuid ensures that no value is present for MobileGuid, not even an explicit nil
### GetNotes

`func (o *ScheduleStopwatch) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ScheduleStopwatch) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ScheduleStopwatch) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ScheduleStopwatch) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetScheduleId

`func (o *ScheduleStopwatch) GetScheduleId() int32`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *ScheduleStopwatch) GetScheduleIdOk() (*int32, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *ScheduleStopwatch) SetScheduleId(v int32)`

SetScheduleId sets ScheduleId field to given value.


### SetScheduleIdNil

`func (o *ScheduleStopwatch) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *ScheduleStopwatch) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetScheduleMobileGuid

`func (o *ScheduleStopwatch) GetScheduleMobileGuid() string`

GetScheduleMobileGuid returns the ScheduleMobileGuid field if non-nil, zero value otherwise.

### GetScheduleMobileGuidOk

`func (o *ScheduleStopwatch) GetScheduleMobileGuidOk() (*string, bool)`

GetScheduleMobileGuidOk returns a tuple with the ScheduleMobileGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleMobileGuid

`func (o *ScheduleStopwatch) SetScheduleMobileGuid(v string)`

SetScheduleMobileGuid sets ScheduleMobileGuid field to given value.

### HasScheduleMobileGuid

`func (o *ScheduleStopwatch) HasScheduleMobileGuid() bool`

HasScheduleMobileGuid returns a boolean if a field has been set.

### SetScheduleMobileGuidNil

`func (o *ScheduleStopwatch) SetScheduleMobileGuidNil(b bool)`

 SetScheduleMobileGuidNil sets the value for ScheduleMobileGuid to be an explicit nil

### UnsetScheduleMobileGuid
`func (o *ScheduleStopwatch) UnsetScheduleMobileGuid()`

UnsetScheduleMobileGuid ensures that no value is present for ScheduleMobileGuid, not even an explicit nil
### GetStartTime

`func (o *ScheduleStopwatch) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ScheduleStopwatch) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ScheduleStopwatch) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ScheduleStopwatch) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *ScheduleStopwatch) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScheduleStopwatch) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScheduleStopwatch) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *ScheduleStopwatch) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ScheduleStopwatch) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTotalPauseTime

`func (o *ScheduleStopwatch) GetTotalPauseTime() int32`

GetTotalPauseTime returns the TotalPauseTime field if non-nil, zero value otherwise.

### GetTotalPauseTimeOk

`func (o *ScheduleStopwatch) GetTotalPauseTimeOk() (*int32, bool)`

GetTotalPauseTimeOk returns a tuple with the TotalPauseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPauseTime

`func (o *ScheduleStopwatch) SetTotalPauseTime(v int32)`

SetTotalPauseTime sets TotalPauseTime field to given value.

### HasTotalPauseTime

`func (o *ScheduleStopwatch) HasTotalPauseTime() bool`

HasTotalPauseTime returns a boolean if a field has been set.

### SetTotalPauseTimeNil

`func (o *ScheduleStopwatch) SetTotalPauseTimeNil(b bool)`

 SetTotalPauseTimeNil sets the value for TotalPauseTime to be an explicit nil

### UnsetTotalPauseTime
`func (o *ScheduleStopwatch) UnsetTotalPauseTime()`

UnsetTotalPauseTime ensures that no value is present for TotalPauseTime, not even an explicit nil
### GetWorkRole

`func (o *ScheduleStopwatch) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *ScheduleStopwatch) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *ScheduleStopwatch) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *ScheduleStopwatch) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetWorkType

`func (o *ScheduleStopwatch) GetWorkType() WorkTypeReference`

GetWorkType returns the WorkType field if non-nil, zero value otherwise.

### GetWorkTypeOk

`func (o *ScheduleStopwatch) GetWorkTypeOk() (*WorkTypeReference, bool)`

GetWorkTypeOk returns a tuple with the WorkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkType

`func (o *ScheduleStopwatch) SetWorkType(v WorkTypeReference)`

SetWorkType sets WorkType field to given value.

### HasWorkType

`func (o *ScheduleStopwatch) HasWorkType() bool`

HasWorkType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


