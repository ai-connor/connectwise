# ActivityStopwatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityId** | **NullableInt32** |  | 
**ActivityMobileGuid** | Pointer to **NullableString** |  | [optional] 
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
**StartTime** | Pointer to **time.Time** |  | [optional] 
**Status** | **NullableString** |  | 
**TotalPauseTime** | Pointer to **NullableInt32** |  | [optional] 
**WorkRole** | Pointer to [**WorkRoleReference**](WorkRoleReference.md) |  | [optional] 
**WorkType** | Pointer to [**WorkTypeReference**](WorkTypeReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewActivityStopwatch

`func NewActivityStopwatch(activityId NullableInt32, member MemberReference, status NullableString, ) *ActivityStopwatch`

NewActivityStopwatch instantiates a new ActivityStopwatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityStopwatchWithDefaults

`func NewActivityStopwatchWithDefaults() *ActivityStopwatch`

NewActivityStopwatchWithDefaults instantiates a new ActivityStopwatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityId

`func (o *ActivityStopwatch) GetActivityId() int32`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *ActivityStopwatch) GetActivityIdOk() (*int32, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *ActivityStopwatch) SetActivityId(v int32)`

SetActivityId sets ActivityId field to given value.


### SetActivityIdNil

`func (o *ActivityStopwatch) SetActivityIdNil(b bool)`

 SetActivityIdNil sets the value for ActivityId to be an explicit nil

### UnsetActivityId
`func (o *ActivityStopwatch) UnsetActivityId()`

UnsetActivityId ensures that no value is present for ActivityId, not even an explicit nil
### GetActivityMobileGuid

`func (o *ActivityStopwatch) GetActivityMobileGuid() string`

GetActivityMobileGuid returns the ActivityMobileGuid field if non-nil, zero value otherwise.

### GetActivityMobileGuidOk

`func (o *ActivityStopwatch) GetActivityMobileGuidOk() (*string, bool)`

GetActivityMobileGuidOk returns a tuple with the ActivityMobileGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityMobileGuid

`func (o *ActivityStopwatch) SetActivityMobileGuid(v string)`

SetActivityMobileGuid sets ActivityMobileGuid field to given value.

### HasActivityMobileGuid

`func (o *ActivityStopwatch) HasActivityMobileGuid() bool`

HasActivityMobileGuid returns a boolean if a field has been set.

### SetActivityMobileGuidNil

`func (o *ActivityStopwatch) SetActivityMobileGuidNil(b bool)`

 SetActivityMobileGuidNil sets the value for ActivityMobileGuid to be an explicit nil

### UnsetActivityMobileGuid
`func (o *ActivityStopwatch) UnsetActivityMobileGuid()`

UnsetActivityMobileGuid ensures that no value is present for ActivityMobileGuid, not even an explicit nil
### GetAgreement

`func (o *ActivityStopwatch) GetAgreement() AgreementReference`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *ActivityStopwatch) GetAgreementOk() (*AgreementReference, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *ActivityStopwatch) SetAgreement(v AgreementReference)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *ActivityStopwatch) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetBillableOption

`func (o *ActivityStopwatch) GetBillableOption() string`

GetBillableOption returns the BillableOption field if non-nil, zero value otherwise.

### GetBillableOptionOk

`func (o *ActivityStopwatch) GetBillableOptionOk() (*string, bool)`

GetBillableOptionOk returns a tuple with the BillableOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableOption

`func (o *ActivityStopwatch) SetBillableOption(v string)`

SetBillableOption sets BillableOption field to given value.

### HasBillableOption

`func (o *ActivityStopwatch) HasBillableOption() bool`

HasBillableOption returns a boolean if a field has been set.

### SetBillableOptionNil

`func (o *ActivityStopwatch) SetBillableOptionNil(b bool)`

 SetBillableOptionNil sets the value for BillableOption to be an explicit nil

### UnsetBillableOption
`func (o *ActivityStopwatch) UnsetBillableOption()`

UnsetBillableOption ensures that no value is present for BillableOption, not even an explicit nil
### GetBusinessUnitId

`func (o *ActivityStopwatch) GetBusinessUnitId() int32`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *ActivityStopwatch) GetBusinessUnitIdOk() (*int32, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *ActivityStopwatch) SetBusinessUnitId(v int32)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *ActivityStopwatch) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### SetBusinessUnitIdNil

`func (o *ActivityStopwatch) SetBusinessUnitIdNil(b bool)`

 SetBusinessUnitIdNil sets the value for BusinessUnitId to be an explicit nil

### UnsetBusinessUnitId
`func (o *ActivityStopwatch) UnsetBusinessUnitId()`

UnsetBusinessUnitId ensures that no value is present for BusinessUnitId, not even an explicit nil
### GetDateEntered

`func (o *ActivityStopwatch) GetDateEntered() time.Time`

GetDateEntered returns the DateEntered field if non-nil, zero value otherwise.

### GetDateEnteredOk

`func (o *ActivityStopwatch) GetDateEnteredOk() (*time.Time, bool)`

GetDateEnteredOk returns a tuple with the DateEntered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEntered

`func (o *ActivityStopwatch) SetDateEntered(v time.Time)`

SetDateEntered sets DateEntered field to given value.

### HasDateEntered

`func (o *ActivityStopwatch) HasDateEntered() bool`

HasDateEntered returns a boolean if a field has been set.

### GetEndTime

`func (o *ActivityStopwatch) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ActivityStopwatch) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ActivityStopwatch) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ActivityStopwatch) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetId

`func (o *ActivityStopwatch) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityStopwatch) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityStopwatch) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityStopwatch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalNotes

`func (o *ActivityStopwatch) GetInternalNotes() string`

GetInternalNotes returns the InternalNotes field if non-nil, zero value otherwise.

### GetInternalNotesOk

`func (o *ActivityStopwatch) GetInternalNotesOk() (*string, bool)`

GetInternalNotesOk returns a tuple with the InternalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNotes

`func (o *ActivityStopwatch) SetInternalNotes(v string)`

SetInternalNotes sets InternalNotes field to given value.

### HasInternalNotes

`func (o *ActivityStopwatch) HasInternalNotes() bool`

HasInternalNotes returns a boolean if a field has been set.

### GetLocationId

`func (o *ActivityStopwatch) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *ActivityStopwatch) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *ActivityStopwatch) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *ActivityStopwatch) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *ActivityStopwatch) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *ActivityStopwatch) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetMember

`func (o *ActivityStopwatch) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ActivityStopwatch) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ActivityStopwatch) SetMember(v MemberReference)`

SetMember sets Member field to given value.


### GetMobileGuid

`func (o *ActivityStopwatch) GetMobileGuid() string`

GetMobileGuid returns the MobileGuid field if non-nil, zero value otherwise.

### GetMobileGuidOk

`func (o *ActivityStopwatch) GetMobileGuidOk() (*string, bool)`

GetMobileGuidOk returns a tuple with the MobileGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileGuid

`func (o *ActivityStopwatch) SetMobileGuid(v string)`

SetMobileGuid sets MobileGuid field to given value.

### HasMobileGuid

`func (o *ActivityStopwatch) HasMobileGuid() bool`

HasMobileGuid returns a boolean if a field has been set.

### SetMobileGuidNil

`func (o *ActivityStopwatch) SetMobileGuidNil(b bool)`

 SetMobileGuidNil sets the value for MobileGuid to be an explicit nil

### UnsetMobileGuid
`func (o *ActivityStopwatch) UnsetMobileGuid()`

UnsetMobileGuid ensures that no value is present for MobileGuid, not even an explicit nil
### GetNotes

`func (o *ActivityStopwatch) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ActivityStopwatch) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ActivityStopwatch) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ActivityStopwatch) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetStartTime

`func (o *ActivityStopwatch) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ActivityStopwatch) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ActivityStopwatch) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ActivityStopwatch) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *ActivityStopwatch) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActivityStopwatch) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActivityStopwatch) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *ActivityStopwatch) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ActivityStopwatch) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTotalPauseTime

`func (o *ActivityStopwatch) GetTotalPauseTime() int32`

GetTotalPauseTime returns the TotalPauseTime field if non-nil, zero value otherwise.

### GetTotalPauseTimeOk

`func (o *ActivityStopwatch) GetTotalPauseTimeOk() (*int32, bool)`

GetTotalPauseTimeOk returns a tuple with the TotalPauseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPauseTime

`func (o *ActivityStopwatch) SetTotalPauseTime(v int32)`

SetTotalPauseTime sets TotalPauseTime field to given value.

### HasTotalPauseTime

`func (o *ActivityStopwatch) HasTotalPauseTime() bool`

HasTotalPauseTime returns a boolean if a field has been set.

### SetTotalPauseTimeNil

`func (o *ActivityStopwatch) SetTotalPauseTimeNil(b bool)`

 SetTotalPauseTimeNil sets the value for TotalPauseTime to be an explicit nil

### UnsetTotalPauseTime
`func (o *ActivityStopwatch) UnsetTotalPauseTime()`

UnsetTotalPauseTime ensures that no value is present for TotalPauseTime, not even an explicit nil
### GetWorkRole

`func (o *ActivityStopwatch) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *ActivityStopwatch) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *ActivityStopwatch) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *ActivityStopwatch) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetWorkType

`func (o *ActivityStopwatch) GetWorkType() WorkTypeReference`

GetWorkType returns the WorkType field if non-nil, zero value otherwise.

### GetWorkTypeOk

`func (o *ActivityStopwatch) GetWorkTypeOk() (*WorkTypeReference, bool)`

GetWorkTypeOk returns a tuple with the WorkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkType

`func (o *ActivityStopwatch) SetWorkType(v WorkTypeReference)`

SetWorkType sets WorkType field to given value.

### HasWorkType

`func (o *ActivityStopwatch) HasWorkType() bool`

HasWorkType returns a boolean if a field has been set.

### GetInfo

`func (o *ActivityStopwatch) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ActivityStopwatch) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ActivityStopwatch) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ActivityStopwatch) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


