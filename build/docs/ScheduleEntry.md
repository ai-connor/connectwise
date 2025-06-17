# ScheduleEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ObjectId** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **string** |  Max length: 250; | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Where** | Pointer to [**ServiceLocationReference**](ServiceLocationReference.md) |  | [optional] 
**DateStart** | Pointer to **time.Time** |  | [optional] 
**DateEnd** | Pointer to **time.Time** |  | [optional] 
**Reminder** | Pointer to [**ReminderReference**](ReminderReference.md) |  | [optional] 
**Status** | Pointer to [**ScheduleStatusReference**](ScheduleStatusReference.md) |  | [optional] 
**Type** | [**ScheduleTypeReference**](ScheduleTypeReference.md) |  | 
**Span** | Pointer to [**ScheduleSpanReference**](ScheduleSpanReference.md) |  | [optional] 
**DoneFlag** | Pointer to **NullableBool** |  | [optional] 
**AcknowledgedFlag** | Pointer to **NullableBool** |  | [optional] 
**OwnerFlag** | Pointer to **NullableBool** |  | [optional] 
**MeetingFlag** | Pointer to **NullableBool** |  | [optional] 
**AllowScheduleConflictsFlag** | Pointer to **NullableBool** |  | [optional] 
**AddMemberToProjectFlag** | Pointer to **NullableBool** |  | [optional] 
**ProjectRoleId** | Pointer to **NullableInt32** |  | [optional] 
**MobileGuid** | Pointer to **NullableString** |  | [optional] 
**AcknowledgedDate** | Pointer to **time.Time** |  | [optional] 
**CloseDate** | Pointer to **time.Time** |  | [optional] 
**NotifyResource** | Pointer to **NullableBool** |  | [optional] 
**NotificationSent** | Pointer to **NullableBool** |  | [optional] 
**NotificationResponse** | Pointer to **string** |  | [optional] 
**Hours** | Pointer to **NullableFloat64** |  | [optional] 
**StartTimeSet** | Pointer to **NullableBool** |  | [optional] 
**EndTimeSet** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewScheduleEntry

`func NewScheduleEntry(type_ ScheduleTypeReference, ) *ScheduleEntry`

NewScheduleEntry instantiates a new ScheduleEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleEntryWithDefaults

`func NewScheduleEntryWithDefaults() *ScheduleEntry`

NewScheduleEntryWithDefaults instantiates a new ScheduleEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduleEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleEntry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleEntry) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ScheduleEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectId

`func (o *ScheduleEntry) GetObjectId() int32`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ScheduleEntry) GetObjectIdOk() (*int32, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ScheduleEntry) SetObjectId(v int32)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *ScheduleEntry) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### SetObjectIdNil

`func (o *ScheduleEntry) SetObjectIdNil(b bool)`

 SetObjectIdNil sets the value for ObjectId to be an explicit nil

### UnsetObjectId
`func (o *ScheduleEntry) UnsetObjectId()`

UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil
### GetName

`func (o *ScheduleEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduleEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduleEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScheduleEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMember

`func (o *ScheduleEntry) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ScheduleEntry) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ScheduleEntry) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *ScheduleEntry) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetWhere

`func (o *ScheduleEntry) GetWhere() ServiceLocationReference`

GetWhere returns the Where field if non-nil, zero value otherwise.

### GetWhereOk

`func (o *ScheduleEntry) GetWhereOk() (*ServiceLocationReference, bool)`

GetWhereOk returns a tuple with the Where field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhere

`func (o *ScheduleEntry) SetWhere(v ServiceLocationReference)`

SetWhere sets Where field to given value.

### HasWhere

`func (o *ScheduleEntry) HasWhere() bool`

HasWhere returns a boolean if a field has been set.

### GetDateStart

`func (o *ScheduleEntry) GetDateStart() time.Time`

GetDateStart returns the DateStart field if non-nil, zero value otherwise.

### GetDateStartOk

`func (o *ScheduleEntry) GetDateStartOk() (*time.Time, bool)`

GetDateStartOk returns a tuple with the DateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStart

`func (o *ScheduleEntry) SetDateStart(v time.Time)`

SetDateStart sets DateStart field to given value.

### HasDateStart

`func (o *ScheduleEntry) HasDateStart() bool`

HasDateStart returns a boolean if a field has been set.

### GetDateEnd

`func (o *ScheduleEntry) GetDateEnd() time.Time`

GetDateEnd returns the DateEnd field if non-nil, zero value otherwise.

### GetDateEndOk

`func (o *ScheduleEntry) GetDateEndOk() (*time.Time, bool)`

GetDateEndOk returns a tuple with the DateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnd

`func (o *ScheduleEntry) SetDateEnd(v time.Time)`

SetDateEnd sets DateEnd field to given value.

### HasDateEnd

`func (o *ScheduleEntry) HasDateEnd() bool`

HasDateEnd returns a boolean if a field has been set.

### GetReminder

`func (o *ScheduleEntry) GetReminder() ReminderReference`

GetReminder returns the Reminder field if non-nil, zero value otherwise.

### GetReminderOk

`func (o *ScheduleEntry) GetReminderOk() (*ReminderReference, bool)`

GetReminderOk returns a tuple with the Reminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminder

`func (o *ScheduleEntry) SetReminder(v ReminderReference)`

SetReminder sets Reminder field to given value.

### HasReminder

`func (o *ScheduleEntry) HasReminder() bool`

HasReminder returns a boolean if a field has been set.

### GetStatus

`func (o *ScheduleEntry) GetStatus() ScheduleStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScheduleEntry) GetStatusOk() (*ScheduleStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScheduleEntry) SetStatus(v ScheduleStatusReference)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ScheduleEntry) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *ScheduleEntry) GetType() ScheduleTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScheduleEntry) GetTypeOk() (*ScheduleTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScheduleEntry) SetType(v ScheduleTypeReference)`

SetType sets Type field to given value.


### GetSpan

`func (o *ScheduleEntry) GetSpan() ScheduleSpanReference`

GetSpan returns the Span field if non-nil, zero value otherwise.

### GetSpanOk

`func (o *ScheduleEntry) GetSpanOk() (*ScheduleSpanReference, bool)`

GetSpanOk returns a tuple with the Span field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpan

`func (o *ScheduleEntry) SetSpan(v ScheduleSpanReference)`

SetSpan sets Span field to given value.

### HasSpan

`func (o *ScheduleEntry) HasSpan() bool`

HasSpan returns a boolean if a field has been set.

### GetDoneFlag

`func (o *ScheduleEntry) GetDoneFlag() bool`

GetDoneFlag returns the DoneFlag field if non-nil, zero value otherwise.

### GetDoneFlagOk

`func (o *ScheduleEntry) GetDoneFlagOk() (*bool, bool)`

GetDoneFlagOk returns a tuple with the DoneFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoneFlag

`func (o *ScheduleEntry) SetDoneFlag(v bool)`

SetDoneFlag sets DoneFlag field to given value.

### HasDoneFlag

`func (o *ScheduleEntry) HasDoneFlag() bool`

HasDoneFlag returns a boolean if a field has been set.

### SetDoneFlagNil

`func (o *ScheduleEntry) SetDoneFlagNil(b bool)`

 SetDoneFlagNil sets the value for DoneFlag to be an explicit nil

### UnsetDoneFlag
`func (o *ScheduleEntry) UnsetDoneFlag()`

UnsetDoneFlag ensures that no value is present for DoneFlag, not even an explicit nil
### GetAcknowledgedFlag

`func (o *ScheduleEntry) GetAcknowledgedFlag() bool`

GetAcknowledgedFlag returns the AcknowledgedFlag field if non-nil, zero value otherwise.

### GetAcknowledgedFlagOk

`func (o *ScheduleEntry) GetAcknowledgedFlagOk() (*bool, bool)`

GetAcknowledgedFlagOk returns a tuple with the AcknowledgedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedFlag

`func (o *ScheduleEntry) SetAcknowledgedFlag(v bool)`

SetAcknowledgedFlag sets AcknowledgedFlag field to given value.

### HasAcknowledgedFlag

`func (o *ScheduleEntry) HasAcknowledgedFlag() bool`

HasAcknowledgedFlag returns a boolean if a field has been set.

### SetAcknowledgedFlagNil

`func (o *ScheduleEntry) SetAcknowledgedFlagNil(b bool)`

 SetAcknowledgedFlagNil sets the value for AcknowledgedFlag to be an explicit nil

### UnsetAcknowledgedFlag
`func (o *ScheduleEntry) UnsetAcknowledgedFlag()`

UnsetAcknowledgedFlag ensures that no value is present for AcknowledgedFlag, not even an explicit nil
### GetOwnerFlag

`func (o *ScheduleEntry) GetOwnerFlag() bool`

GetOwnerFlag returns the OwnerFlag field if non-nil, zero value otherwise.

### GetOwnerFlagOk

`func (o *ScheduleEntry) GetOwnerFlagOk() (*bool, bool)`

GetOwnerFlagOk returns a tuple with the OwnerFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerFlag

`func (o *ScheduleEntry) SetOwnerFlag(v bool)`

SetOwnerFlag sets OwnerFlag field to given value.

### HasOwnerFlag

`func (o *ScheduleEntry) HasOwnerFlag() bool`

HasOwnerFlag returns a boolean if a field has been set.

### SetOwnerFlagNil

`func (o *ScheduleEntry) SetOwnerFlagNil(b bool)`

 SetOwnerFlagNil sets the value for OwnerFlag to be an explicit nil

### UnsetOwnerFlag
`func (o *ScheduleEntry) UnsetOwnerFlag()`

UnsetOwnerFlag ensures that no value is present for OwnerFlag, not even an explicit nil
### GetMeetingFlag

`func (o *ScheduleEntry) GetMeetingFlag() bool`

GetMeetingFlag returns the MeetingFlag field if non-nil, zero value otherwise.

### GetMeetingFlagOk

`func (o *ScheduleEntry) GetMeetingFlagOk() (*bool, bool)`

GetMeetingFlagOk returns a tuple with the MeetingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeetingFlag

`func (o *ScheduleEntry) SetMeetingFlag(v bool)`

SetMeetingFlag sets MeetingFlag field to given value.

### HasMeetingFlag

`func (o *ScheduleEntry) HasMeetingFlag() bool`

HasMeetingFlag returns a boolean if a field has been set.

### SetMeetingFlagNil

`func (o *ScheduleEntry) SetMeetingFlagNil(b bool)`

 SetMeetingFlagNil sets the value for MeetingFlag to be an explicit nil

### UnsetMeetingFlag
`func (o *ScheduleEntry) UnsetMeetingFlag()`

UnsetMeetingFlag ensures that no value is present for MeetingFlag, not even an explicit nil
### GetAllowScheduleConflictsFlag

`func (o *ScheduleEntry) GetAllowScheduleConflictsFlag() bool`

GetAllowScheduleConflictsFlag returns the AllowScheduleConflictsFlag field if non-nil, zero value otherwise.

### GetAllowScheduleConflictsFlagOk

`func (o *ScheduleEntry) GetAllowScheduleConflictsFlagOk() (*bool, bool)`

GetAllowScheduleConflictsFlagOk returns a tuple with the AllowScheduleConflictsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowScheduleConflictsFlag

`func (o *ScheduleEntry) SetAllowScheduleConflictsFlag(v bool)`

SetAllowScheduleConflictsFlag sets AllowScheduleConflictsFlag field to given value.

### HasAllowScheduleConflictsFlag

`func (o *ScheduleEntry) HasAllowScheduleConflictsFlag() bool`

HasAllowScheduleConflictsFlag returns a boolean if a field has been set.

### SetAllowScheduleConflictsFlagNil

`func (o *ScheduleEntry) SetAllowScheduleConflictsFlagNil(b bool)`

 SetAllowScheduleConflictsFlagNil sets the value for AllowScheduleConflictsFlag to be an explicit nil

### UnsetAllowScheduleConflictsFlag
`func (o *ScheduleEntry) UnsetAllowScheduleConflictsFlag()`

UnsetAllowScheduleConflictsFlag ensures that no value is present for AllowScheduleConflictsFlag, not even an explicit nil
### GetAddMemberToProjectFlag

`func (o *ScheduleEntry) GetAddMemberToProjectFlag() bool`

GetAddMemberToProjectFlag returns the AddMemberToProjectFlag field if non-nil, zero value otherwise.

### GetAddMemberToProjectFlagOk

`func (o *ScheduleEntry) GetAddMemberToProjectFlagOk() (*bool, bool)`

GetAddMemberToProjectFlagOk returns a tuple with the AddMemberToProjectFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddMemberToProjectFlag

`func (o *ScheduleEntry) SetAddMemberToProjectFlag(v bool)`

SetAddMemberToProjectFlag sets AddMemberToProjectFlag field to given value.

### HasAddMemberToProjectFlag

`func (o *ScheduleEntry) HasAddMemberToProjectFlag() bool`

HasAddMemberToProjectFlag returns a boolean if a field has been set.

### SetAddMemberToProjectFlagNil

`func (o *ScheduleEntry) SetAddMemberToProjectFlagNil(b bool)`

 SetAddMemberToProjectFlagNil sets the value for AddMemberToProjectFlag to be an explicit nil

### UnsetAddMemberToProjectFlag
`func (o *ScheduleEntry) UnsetAddMemberToProjectFlag()`

UnsetAddMemberToProjectFlag ensures that no value is present for AddMemberToProjectFlag, not even an explicit nil
### GetProjectRoleId

`func (o *ScheduleEntry) GetProjectRoleId() int32`

GetProjectRoleId returns the ProjectRoleId field if non-nil, zero value otherwise.

### GetProjectRoleIdOk

`func (o *ScheduleEntry) GetProjectRoleIdOk() (*int32, bool)`

GetProjectRoleIdOk returns a tuple with the ProjectRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectRoleId

`func (o *ScheduleEntry) SetProjectRoleId(v int32)`

SetProjectRoleId sets ProjectRoleId field to given value.

### HasProjectRoleId

`func (o *ScheduleEntry) HasProjectRoleId() bool`

HasProjectRoleId returns a boolean if a field has been set.

### SetProjectRoleIdNil

`func (o *ScheduleEntry) SetProjectRoleIdNil(b bool)`

 SetProjectRoleIdNil sets the value for ProjectRoleId to be an explicit nil

### UnsetProjectRoleId
`func (o *ScheduleEntry) UnsetProjectRoleId()`

UnsetProjectRoleId ensures that no value is present for ProjectRoleId, not even an explicit nil
### GetMobileGuid

`func (o *ScheduleEntry) GetMobileGuid() string`

GetMobileGuid returns the MobileGuid field if non-nil, zero value otherwise.

### GetMobileGuidOk

`func (o *ScheduleEntry) GetMobileGuidOk() (*string, bool)`

GetMobileGuidOk returns a tuple with the MobileGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileGuid

`func (o *ScheduleEntry) SetMobileGuid(v string)`

SetMobileGuid sets MobileGuid field to given value.

### HasMobileGuid

`func (o *ScheduleEntry) HasMobileGuid() bool`

HasMobileGuid returns a boolean if a field has been set.

### SetMobileGuidNil

`func (o *ScheduleEntry) SetMobileGuidNil(b bool)`

 SetMobileGuidNil sets the value for MobileGuid to be an explicit nil

### UnsetMobileGuid
`func (o *ScheduleEntry) UnsetMobileGuid()`

UnsetMobileGuid ensures that no value is present for MobileGuid, not even an explicit nil
### GetAcknowledgedDate

`func (o *ScheduleEntry) GetAcknowledgedDate() time.Time`

GetAcknowledgedDate returns the AcknowledgedDate field if non-nil, zero value otherwise.

### GetAcknowledgedDateOk

`func (o *ScheduleEntry) GetAcknowledgedDateOk() (*time.Time, bool)`

GetAcknowledgedDateOk returns a tuple with the AcknowledgedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedDate

`func (o *ScheduleEntry) SetAcknowledgedDate(v time.Time)`

SetAcknowledgedDate sets AcknowledgedDate field to given value.

### HasAcknowledgedDate

`func (o *ScheduleEntry) HasAcknowledgedDate() bool`

HasAcknowledgedDate returns a boolean if a field has been set.

### GetCloseDate

`func (o *ScheduleEntry) GetCloseDate() time.Time`

GetCloseDate returns the CloseDate field if non-nil, zero value otherwise.

### GetCloseDateOk

`func (o *ScheduleEntry) GetCloseDateOk() (*time.Time, bool)`

GetCloseDateOk returns a tuple with the CloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseDate

`func (o *ScheduleEntry) SetCloseDate(v time.Time)`

SetCloseDate sets CloseDate field to given value.

### HasCloseDate

`func (o *ScheduleEntry) HasCloseDate() bool`

HasCloseDate returns a boolean if a field has been set.

### GetNotifyResource

`func (o *ScheduleEntry) GetNotifyResource() bool`

GetNotifyResource returns the NotifyResource field if non-nil, zero value otherwise.

### GetNotifyResourceOk

`func (o *ScheduleEntry) GetNotifyResourceOk() (*bool, bool)`

GetNotifyResourceOk returns a tuple with the NotifyResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyResource

`func (o *ScheduleEntry) SetNotifyResource(v bool)`

SetNotifyResource sets NotifyResource field to given value.

### HasNotifyResource

`func (o *ScheduleEntry) HasNotifyResource() bool`

HasNotifyResource returns a boolean if a field has been set.

### SetNotifyResourceNil

`func (o *ScheduleEntry) SetNotifyResourceNil(b bool)`

 SetNotifyResourceNil sets the value for NotifyResource to be an explicit nil

### UnsetNotifyResource
`func (o *ScheduleEntry) UnsetNotifyResource()`

UnsetNotifyResource ensures that no value is present for NotifyResource, not even an explicit nil
### GetNotificationSent

`func (o *ScheduleEntry) GetNotificationSent() bool`

GetNotificationSent returns the NotificationSent field if non-nil, zero value otherwise.

### GetNotificationSentOk

`func (o *ScheduleEntry) GetNotificationSentOk() (*bool, bool)`

GetNotificationSentOk returns a tuple with the NotificationSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSent

`func (o *ScheduleEntry) SetNotificationSent(v bool)`

SetNotificationSent sets NotificationSent field to given value.

### HasNotificationSent

`func (o *ScheduleEntry) HasNotificationSent() bool`

HasNotificationSent returns a boolean if a field has been set.

### SetNotificationSentNil

`func (o *ScheduleEntry) SetNotificationSentNil(b bool)`

 SetNotificationSentNil sets the value for NotificationSent to be an explicit nil

### UnsetNotificationSent
`func (o *ScheduleEntry) UnsetNotificationSent()`

UnsetNotificationSent ensures that no value is present for NotificationSent, not even an explicit nil
### GetNotificationResponse

`func (o *ScheduleEntry) GetNotificationResponse() string`

GetNotificationResponse returns the NotificationResponse field if non-nil, zero value otherwise.

### GetNotificationResponseOk

`func (o *ScheduleEntry) GetNotificationResponseOk() (*string, bool)`

GetNotificationResponseOk returns a tuple with the NotificationResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationResponse

`func (o *ScheduleEntry) SetNotificationResponse(v string)`

SetNotificationResponse sets NotificationResponse field to given value.

### HasNotificationResponse

`func (o *ScheduleEntry) HasNotificationResponse() bool`

HasNotificationResponse returns a boolean if a field has been set.

### GetHours

`func (o *ScheduleEntry) GetHours() float64`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *ScheduleEntry) GetHoursOk() (*float64, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *ScheduleEntry) SetHours(v float64)`

SetHours sets Hours field to given value.

### HasHours

`func (o *ScheduleEntry) HasHours() bool`

HasHours returns a boolean if a field has been set.

### SetHoursNil

`func (o *ScheduleEntry) SetHoursNil(b bool)`

 SetHoursNil sets the value for Hours to be an explicit nil

### UnsetHours
`func (o *ScheduleEntry) UnsetHours()`

UnsetHours ensures that no value is present for Hours, not even an explicit nil
### GetStartTimeSet

`func (o *ScheduleEntry) GetStartTimeSet() bool`

GetStartTimeSet returns the StartTimeSet field if non-nil, zero value otherwise.

### GetStartTimeSetOk

`func (o *ScheduleEntry) GetStartTimeSetOk() (*bool, bool)`

GetStartTimeSetOk returns a tuple with the StartTimeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeSet

`func (o *ScheduleEntry) SetStartTimeSet(v bool)`

SetStartTimeSet sets StartTimeSet field to given value.

### HasStartTimeSet

`func (o *ScheduleEntry) HasStartTimeSet() bool`

HasStartTimeSet returns a boolean if a field has been set.

### SetStartTimeSetNil

`func (o *ScheduleEntry) SetStartTimeSetNil(b bool)`

 SetStartTimeSetNil sets the value for StartTimeSet to be an explicit nil

### UnsetStartTimeSet
`func (o *ScheduleEntry) UnsetStartTimeSet()`

UnsetStartTimeSet ensures that no value is present for StartTimeSet, not even an explicit nil
### GetEndTimeSet

`func (o *ScheduleEntry) GetEndTimeSet() bool`

GetEndTimeSet returns the EndTimeSet field if non-nil, zero value otherwise.

### GetEndTimeSetOk

`func (o *ScheduleEntry) GetEndTimeSetOk() (*bool, bool)`

GetEndTimeSetOk returns a tuple with the EndTimeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeSet

`func (o *ScheduleEntry) SetEndTimeSet(v bool)`

SetEndTimeSet sets EndTimeSet field to given value.

### HasEndTimeSet

`func (o *ScheduleEntry) HasEndTimeSet() bool`

HasEndTimeSet returns a boolean if a field has been set.

### SetEndTimeSetNil

`func (o *ScheduleEntry) SetEndTimeSetNil(b bool)`

 SetEndTimeSetNil sets the value for EndTimeSet to be an explicit nil

### UnsetEndTimeSet
`func (o *ScheduleEntry) UnsetEndTimeSet()`

UnsetEndTimeSet ensures that no value is present for EndTimeSet, not even an explicit nil
### GetInfo

`func (o *ScheduleEntry) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ScheduleEntry) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ScheduleEntry) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ScheduleEntry) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


