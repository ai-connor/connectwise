# TicketStopwatch

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
**ServiceStatus** | Pointer to [**ServiceStatusReference**](ServiceStatusReference.md) |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**Status** | **NullableString** |  | 
**Ticket** | [**TicketReference**](TicketReference.md) |  | 
**TicketMobileGuid** | Pointer to **NullableString** |  | [optional] 
**TotalPauseTime** | Pointer to **NullableInt32** |  | [optional] 
**WorkRole** | Pointer to [**WorkRoleReference**](WorkRoleReference.md) |  | [optional] 
**WorkType** | Pointer to [**WorkTypeReference**](WorkTypeReference.md) |  | [optional] 
**ShowNotesInDiscussionFlag** | Pointer to **NullableBool** |  | [optional] 
**ShowNotesInInternalFlag** | Pointer to **NullableBool** |  | [optional] 
**ShowNotesInResolutionFlag** | Pointer to **NullableBool** |  | [optional] 
**EmailNotesToContactFlag** | Pointer to **NullableBool** |  | [optional] 
**EmailNotesToResourcesFlag** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewTicketStopwatch

`func NewTicketStopwatch(member MemberReference, status NullableString, ticket TicketReference, ) *TicketStopwatch`

NewTicketStopwatch instantiates a new TicketStopwatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketStopwatchWithDefaults

`func NewTicketStopwatchWithDefaults() *TicketStopwatch`

NewTicketStopwatchWithDefaults instantiates a new TicketStopwatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *TicketStopwatch) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TicketStopwatch) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TicketStopwatch) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TicketStopwatch) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetAgreement

`func (o *TicketStopwatch) GetAgreement() AgreementReference`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *TicketStopwatch) GetAgreementOk() (*AgreementReference, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *TicketStopwatch) SetAgreement(v AgreementReference)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *TicketStopwatch) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetBillableOption

`func (o *TicketStopwatch) GetBillableOption() string`

GetBillableOption returns the BillableOption field if non-nil, zero value otherwise.

### GetBillableOptionOk

`func (o *TicketStopwatch) GetBillableOptionOk() (*string, bool)`

GetBillableOptionOk returns a tuple with the BillableOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableOption

`func (o *TicketStopwatch) SetBillableOption(v string)`

SetBillableOption sets BillableOption field to given value.

### HasBillableOption

`func (o *TicketStopwatch) HasBillableOption() bool`

HasBillableOption returns a boolean if a field has been set.

### SetBillableOptionNil

`func (o *TicketStopwatch) SetBillableOptionNil(b bool)`

 SetBillableOptionNil sets the value for BillableOption to be an explicit nil

### UnsetBillableOption
`func (o *TicketStopwatch) UnsetBillableOption()`

UnsetBillableOption ensures that no value is present for BillableOption, not even an explicit nil
### GetBusinessUnitId

`func (o *TicketStopwatch) GetBusinessUnitId() int32`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *TicketStopwatch) GetBusinessUnitIdOk() (*int32, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *TicketStopwatch) SetBusinessUnitId(v int32)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *TicketStopwatch) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### SetBusinessUnitIdNil

`func (o *TicketStopwatch) SetBusinessUnitIdNil(b bool)`

 SetBusinessUnitIdNil sets the value for BusinessUnitId to be an explicit nil

### UnsetBusinessUnitId
`func (o *TicketStopwatch) UnsetBusinessUnitId()`

UnsetBusinessUnitId ensures that no value is present for BusinessUnitId, not even an explicit nil
### GetDateEntered

`func (o *TicketStopwatch) GetDateEntered() time.Time`

GetDateEntered returns the DateEntered field if non-nil, zero value otherwise.

### GetDateEnteredOk

`func (o *TicketStopwatch) GetDateEnteredOk() (*time.Time, bool)`

GetDateEnteredOk returns a tuple with the DateEntered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEntered

`func (o *TicketStopwatch) SetDateEntered(v time.Time)`

SetDateEntered sets DateEntered field to given value.

### HasDateEntered

`func (o *TicketStopwatch) HasDateEntered() bool`

HasDateEntered returns a boolean if a field has been set.

### GetEndTime

`func (o *TicketStopwatch) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TicketStopwatch) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TicketStopwatch) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TicketStopwatch) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetId

`func (o *TicketStopwatch) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TicketStopwatch) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TicketStopwatch) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TicketStopwatch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalNotes

`func (o *TicketStopwatch) GetInternalNotes() string`

GetInternalNotes returns the InternalNotes field if non-nil, zero value otherwise.

### GetInternalNotesOk

`func (o *TicketStopwatch) GetInternalNotesOk() (*string, bool)`

GetInternalNotesOk returns a tuple with the InternalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNotes

`func (o *TicketStopwatch) SetInternalNotes(v string)`

SetInternalNotes sets InternalNotes field to given value.

### HasInternalNotes

`func (o *TicketStopwatch) HasInternalNotes() bool`

HasInternalNotes returns a boolean if a field has been set.

### GetLocationId

`func (o *TicketStopwatch) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *TicketStopwatch) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *TicketStopwatch) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *TicketStopwatch) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *TicketStopwatch) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *TicketStopwatch) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetMember

`func (o *TicketStopwatch) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *TicketStopwatch) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *TicketStopwatch) SetMember(v MemberReference)`

SetMember sets Member field to given value.


### GetMobileGuid

`func (o *TicketStopwatch) GetMobileGuid() string`

GetMobileGuid returns the MobileGuid field if non-nil, zero value otherwise.

### GetMobileGuidOk

`func (o *TicketStopwatch) GetMobileGuidOk() (*string, bool)`

GetMobileGuidOk returns a tuple with the MobileGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileGuid

`func (o *TicketStopwatch) SetMobileGuid(v string)`

SetMobileGuid sets MobileGuid field to given value.

### HasMobileGuid

`func (o *TicketStopwatch) HasMobileGuid() bool`

HasMobileGuid returns a boolean if a field has been set.

### SetMobileGuidNil

`func (o *TicketStopwatch) SetMobileGuidNil(b bool)`

 SetMobileGuidNil sets the value for MobileGuid to be an explicit nil

### UnsetMobileGuid
`func (o *TicketStopwatch) UnsetMobileGuid()`

UnsetMobileGuid ensures that no value is present for MobileGuid, not even an explicit nil
### GetNotes

`func (o *TicketStopwatch) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TicketStopwatch) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TicketStopwatch) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TicketStopwatch) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetServiceStatus

`func (o *TicketStopwatch) GetServiceStatus() ServiceStatusReference`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *TicketStopwatch) GetServiceStatusOk() (*ServiceStatusReference, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *TicketStopwatch) SetServiceStatus(v ServiceStatusReference)`

SetServiceStatus sets ServiceStatus field to given value.

### HasServiceStatus

`func (o *TicketStopwatch) HasServiceStatus() bool`

HasServiceStatus returns a boolean if a field has been set.

### GetStartTime

`func (o *TicketStopwatch) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TicketStopwatch) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TicketStopwatch) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TicketStopwatch) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *TicketStopwatch) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TicketStopwatch) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TicketStopwatch) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *TicketStopwatch) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TicketStopwatch) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTicket

`func (o *TicketStopwatch) GetTicket() TicketReference`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *TicketStopwatch) GetTicketOk() (*TicketReference, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *TicketStopwatch) SetTicket(v TicketReference)`

SetTicket sets Ticket field to given value.


### GetTicketMobileGuid

`func (o *TicketStopwatch) GetTicketMobileGuid() string`

GetTicketMobileGuid returns the TicketMobileGuid field if non-nil, zero value otherwise.

### GetTicketMobileGuidOk

`func (o *TicketStopwatch) GetTicketMobileGuidOk() (*string, bool)`

GetTicketMobileGuidOk returns a tuple with the TicketMobileGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketMobileGuid

`func (o *TicketStopwatch) SetTicketMobileGuid(v string)`

SetTicketMobileGuid sets TicketMobileGuid field to given value.

### HasTicketMobileGuid

`func (o *TicketStopwatch) HasTicketMobileGuid() bool`

HasTicketMobileGuid returns a boolean if a field has been set.

### SetTicketMobileGuidNil

`func (o *TicketStopwatch) SetTicketMobileGuidNil(b bool)`

 SetTicketMobileGuidNil sets the value for TicketMobileGuid to be an explicit nil

### UnsetTicketMobileGuid
`func (o *TicketStopwatch) UnsetTicketMobileGuid()`

UnsetTicketMobileGuid ensures that no value is present for TicketMobileGuid, not even an explicit nil
### GetTotalPauseTime

`func (o *TicketStopwatch) GetTotalPauseTime() int32`

GetTotalPauseTime returns the TotalPauseTime field if non-nil, zero value otherwise.

### GetTotalPauseTimeOk

`func (o *TicketStopwatch) GetTotalPauseTimeOk() (*int32, bool)`

GetTotalPauseTimeOk returns a tuple with the TotalPauseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPauseTime

`func (o *TicketStopwatch) SetTotalPauseTime(v int32)`

SetTotalPauseTime sets TotalPauseTime field to given value.

### HasTotalPauseTime

`func (o *TicketStopwatch) HasTotalPauseTime() bool`

HasTotalPauseTime returns a boolean if a field has been set.

### SetTotalPauseTimeNil

`func (o *TicketStopwatch) SetTotalPauseTimeNil(b bool)`

 SetTotalPauseTimeNil sets the value for TotalPauseTime to be an explicit nil

### UnsetTotalPauseTime
`func (o *TicketStopwatch) UnsetTotalPauseTime()`

UnsetTotalPauseTime ensures that no value is present for TotalPauseTime, not even an explicit nil
### GetWorkRole

`func (o *TicketStopwatch) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *TicketStopwatch) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *TicketStopwatch) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *TicketStopwatch) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetWorkType

`func (o *TicketStopwatch) GetWorkType() WorkTypeReference`

GetWorkType returns the WorkType field if non-nil, zero value otherwise.

### GetWorkTypeOk

`func (o *TicketStopwatch) GetWorkTypeOk() (*WorkTypeReference, bool)`

GetWorkTypeOk returns a tuple with the WorkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkType

`func (o *TicketStopwatch) SetWorkType(v WorkTypeReference)`

SetWorkType sets WorkType field to given value.

### HasWorkType

`func (o *TicketStopwatch) HasWorkType() bool`

HasWorkType returns a boolean if a field has been set.

### GetShowNotesInDiscussionFlag

`func (o *TicketStopwatch) GetShowNotesInDiscussionFlag() bool`

GetShowNotesInDiscussionFlag returns the ShowNotesInDiscussionFlag field if non-nil, zero value otherwise.

### GetShowNotesInDiscussionFlagOk

`func (o *TicketStopwatch) GetShowNotesInDiscussionFlagOk() (*bool, bool)`

GetShowNotesInDiscussionFlagOk returns a tuple with the ShowNotesInDiscussionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNotesInDiscussionFlag

`func (o *TicketStopwatch) SetShowNotesInDiscussionFlag(v bool)`

SetShowNotesInDiscussionFlag sets ShowNotesInDiscussionFlag field to given value.

### HasShowNotesInDiscussionFlag

`func (o *TicketStopwatch) HasShowNotesInDiscussionFlag() bool`

HasShowNotesInDiscussionFlag returns a boolean if a field has been set.

### SetShowNotesInDiscussionFlagNil

`func (o *TicketStopwatch) SetShowNotesInDiscussionFlagNil(b bool)`

 SetShowNotesInDiscussionFlagNil sets the value for ShowNotesInDiscussionFlag to be an explicit nil

### UnsetShowNotesInDiscussionFlag
`func (o *TicketStopwatch) UnsetShowNotesInDiscussionFlag()`

UnsetShowNotesInDiscussionFlag ensures that no value is present for ShowNotesInDiscussionFlag, not even an explicit nil
### GetShowNotesInInternalFlag

`func (o *TicketStopwatch) GetShowNotesInInternalFlag() bool`

GetShowNotesInInternalFlag returns the ShowNotesInInternalFlag field if non-nil, zero value otherwise.

### GetShowNotesInInternalFlagOk

`func (o *TicketStopwatch) GetShowNotesInInternalFlagOk() (*bool, bool)`

GetShowNotesInInternalFlagOk returns a tuple with the ShowNotesInInternalFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNotesInInternalFlag

`func (o *TicketStopwatch) SetShowNotesInInternalFlag(v bool)`

SetShowNotesInInternalFlag sets ShowNotesInInternalFlag field to given value.

### HasShowNotesInInternalFlag

`func (o *TicketStopwatch) HasShowNotesInInternalFlag() bool`

HasShowNotesInInternalFlag returns a boolean if a field has been set.

### SetShowNotesInInternalFlagNil

`func (o *TicketStopwatch) SetShowNotesInInternalFlagNil(b bool)`

 SetShowNotesInInternalFlagNil sets the value for ShowNotesInInternalFlag to be an explicit nil

### UnsetShowNotesInInternalFlag
`func (o *TicketStopwatch) UnsetShowNotesInInternalFlag()`

UnsetShowNotesInInternalFlag ensures that no value is present for ShowNotesInInternalFlag, not even an explicit nil
### GetShowNotesInResolutionFlag

`func (o *TicketStopwatch) GetShowNotesInResolutionFlag() bool`

GetShowNotesInResolutionFlag returns the ShowNotesInResolutionFlag field if non-nil, zero value otherwise.

### GetShowNotesInResolutionFlagOk

`func (o *TicketStopwatch) GetShowNotesInResolutionFlagOk() (*bool, bool)`

GetShowNotesInResolutionFlagOk returns a tuple with the ShowNotesInResolutionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNotesInResolutionFlag

`func (o *TicketStopwatch) SetShowNotesInResolutionFlag(v bool)`

SetShowNotesInResolutionFlag sets ShowNotesInResolutionFlag field to given value.

### HasShowNotesInResolutionFlag

`func (o *TicketStopwatch) HasShowNotesInResolutionFlag() bool`

HasShowNotesInResolutionFlag returns a boolean if a field has been set.

### SetShowNotesInResolutionFlagNil

`func (o *TicketStopwatch) SetShowNotesInResolutionFlagNil(b bool)`

 SetShowNotesInResolutionFlagNil sets the value for ShowNotesInResolutionFlag to be an explicit nil

### UnsetShowNotesInResolutionFlag
`func (o *TicketStopwatch) UnsetShowNotesInResolutionFlag()`

UnsetShowNotesInResolutionFlag ensures that no value is present for ShowNotesInResolutionFlag, not even an explicit nil
### GetEmailNotesToContactFlag

`func (o *TicketStopwatch) GetEmailNotesToContactFlag() bool`

GetEmailNotesToContactFlag returns the EmailNotesToContactFlag field if non-nil, zero value otherwise.

### GetEmailNotesToContactFlagOk

`func (o *TicketStopwatch) GetEmailNotesToContactFlagOk() (*bool, bool)`

GetEmailNotesToContactFlagOk returns a tuple with the EmailNotesToContactFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotesToContactFlag

`func (o *TicketStopwatch) SetEmailNotesToContactFlag(v bool)`

SetEmailNotesToContactFlag sets EmailNotesToContactFlag field to given value.

### HasEmailNotesToContactFlag

`func (o *TicketStopwatch) HasEmailNotesToContactFlag() bool`

HasEmailNotesToContactFlag returns a boolean if a field has been set.

### SetEmailNotesToContactFlagNil

`func (o *TicketStopwatch) SetEmailNotesToContactFlagNil(b bool)`

 SetEmailNotesToContactFlagNil sets the value for EmailNotesToContactFlag to be an explicit nil

### UnsetEmailNotesToContactFlag
`func (o *TicketStopwatch) UnsetEmailNotesToContactFlag()`

UnsetEmailNotesToContactFlag ensures that no value is present for EmailNotesToContactFlag, not even an explicit nil
### GetEmailNotesToResourcesFlag

`func (o *TicketStopwatch) GetEmailNotesToResourcesFlag() bool`

GetEmailNotesToResourcesFlag returns the EmailNotesToResourcesFlag field if non-nil, zero value otherwise.

### GetEmailNotesToResourcesFlagOk

`func (o *TicketStopwatch) GetEmailNotesToResourcesFlagOk() (*bool, bool)`

GetEmailNotesToResourcesFlagOk returns a tuple with the EmailNotesToResourcesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotesToResourcesFlag

`func (o *TicketStopwatch) SetEmailNotesToResourcesFlag(v bool)`

SetEmailNotesToResourcesFlag sets EmailNotesToResourcesFlag field to given value.

### HasEmailNotesToResourcesFlag

`func (o *TicketStopwatch) HasEmailNotesToResourcesFlag() bool`

HasEmailNotesToResourcesFlag returns a boolean if a field has been set.

### SetEmailNotesToResourcesFlagNil

`func (o *TicketStopwatch) SetEmailNotesToResourcesFlagNil(b bool)`

 SetEmailNotesToResourcesFlagNil sets the value for EmailNotesToResourcesFlag to be an explicit nil

### UnsetEmailNotesToResourcesFlag
`func (o *TicketStopwatch) UnsetEmailNotesToResourcesFlag()`

UnsetEmailNotesToResourcesFlag ensures that no value is present for EmailNotesToResourcesFlag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


