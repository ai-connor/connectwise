# TrackAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**NotifyType** | **string** |  | 
**ServiceTemplate** | Pointer to [**ServiceTemplateReference**](ServiceTemplateReference.md) |  | [optional] 
**SpecificMemberTo** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**EmailRecipient** | Pointer to **string** |  Max length: 250; | [optional] 
**SpecificMemberFrom** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**EmailFrom** | Pointer to **string** |  Max length: 250; | [optional] 
**Subject** | Pointer to **string** |  Max length: 100; | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**ActivityType** | Pointer to [**ActivityTypeReference**](ActivityTypeReference.md) |  | [optional] 
**ActivityStatus** | Pointer to [**ActivityStatusReference**](ActivityStatusReference.md) |  | [optional] 
**CompanyStatus** | Pointer to [**CompanyStatusReference**](CompanyStatusReference.md) |  | [optional] 
**Track** | Pointer to [**TrackReference**](TrackReference.md) |  | [optional] 
**AttachedTrack** | Pointer to [**TrackReference**](TrackReference.md) |  | [optional] 
**Group** | Pointer to [**GroupReference**](GroupReference.md) |  | [optional] 
**CcContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**BccContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**DaysToExecute** | Pointer to **NullableInt32** |  | [optional] 
**NotifyWho** | Pointer to [**NotificationRecipientReference**](NotificationRecipientReference.md) |  | [optional] 
**NotifyFrom** | Pointer to [**NotificationRecipientReference**](NotificationRecipientReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTrackAction

`func NewTrackAction(notifyType string, ) *TrackAction`

NewTrackAction instantiates a new TrackAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackActionWithDefaults

`func NewTrackActionWithDefaults() *TrackAction`

NewTrackActionWithDefaults instantiates a new TrackAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrackAction) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrackAction) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrackAction) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TrackAction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotifyType

`func (o *TrackAction) GetNotifyType() string`

GetNotifyType returns the NotifyType field if non-nil, zero value otherwise.

### GetNotifyTypeOk

`func (o *TrackAction) GetNotifyTypeOk() (*string, bool)`

GetNotifyTypeOk returns a tuple with the NotifyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyType

`func (o *TrackAction) SetNotifyType(v string)`

SetNotifyType sets NotifyType field to given value.


### GetServiceTemplate

`func (o *TrackAction) GetServiceTemplate() ServiceTemplateReference`

GetServiceTemplate returns the ServiceTemplate field if non-nil, zero value otherwise.

### GetServiceTemplateOk

`func (o *TrackAction) GetServiceTemplateOk() (*ServiceTemplateReference, bool)`

GetServiceTemplateOk returns a tuple with the ServiceTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTemplate

`func (o *TrackAction) SetServiceTemplate(v ServiceTemplateReference)`

SetServiceTemplate sets ServiceTemplate field to given value.

### HasServiceTemplate

`func (o *TrackAction) HasServiceTemplate() bool`

HasServiceTemplate returns a boolean if a field has been set.

### GetSpecificMemberTo

`func (o *TrackAction) GetSpecificMemberTo() MemberReference`

GetSpecificMemberTo returns the SpecificMemberTo field if non-nil, zero value otherwise.

### GetSpecificMemberToOk

`func (o *TrackAction) GetSpecificMemberToOk() (*MemberReference, bool)`

GetSpecificMemberToOk returns a tuple with the SpecificMemberTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificMemberTo

`func (o *TrackAction) SetSpecificMemberTo(v MemberReference)`

SetSpecificMemberTo sets SpecificMemberTo field to given value.

### HasSpecificMemberTo

`func (o *TrackAction) HasSpecificMemberTo() bool`

HasSpecificMemberTo returns a boolean if a field has been set.

### GetEmailRecipient

`func (o *TrackAction) GetEmailRecipient() string`

GetEmailRecipient returns the EmailRecipient field if non-nil, zero value otherwise.

### GetEmailRecipientOk

`func (o *TrackAction) GetEmailRecipientOk() (*string, bool)`

GetEmailRecipientOk returns a tuple with the EmailRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailRecipient

`func (o *TrackAction) SetEmailRecipient(v string)`

SetEmailRecipient sets EmailRecipient field to given value.

### HasEmailRecipient

`func (o *TrackAction) HasEmailRecipient() bool`

HasEmailRecipient returns a boolean if a field has been set.

### GetSpecificMemberFrom

`func (o *TrackAction) GetSpecificMemberFrom() MemberReference`

GetSpecificMemberFrom returns the SpecificMemberFrom field if non-nil, zero value otherwise.

### GetSpecificMemberFromOk

`func (o *TrackAction) GetSpecificMemberFromOk() (*MemberReference, bool)`

GetSpecificMemberFromOk returns a tuple with the SpecificMemberFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificMemberFrom

`func (o *TrackAction) SetSpecificMemberFrom(v MemberReference)`

SetSpecificMemberFrom sets SpecificMemberFrom field to given value.

### HasSpecificMemberFrom

`func (o *TrackAction) HasSpecificMemberFrom() bool`

HasSpecificMemberFrom returns a boolean if a field has been set.

### GetEmailFrom

`func (o *TrackAction) GetEmailFrom() string`

GetEmailFrom returns the EmailFrom field if non-nil, zero value otherwise.

### GetEmailFromOk

`func (o *TrackAction) GetEmailFromOk() (*string, bool)`

GetEmailFromOk returns a tuple with the EmailFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFrom

`func (o *TrackAction) SetEmailFrom(v string)`

SetEmailFrom sets EmailFrom field to given value.

### HasEmailFrom

`func (o *TrackAction) HasEmailFrom() bool`

HasEmailFrom returns a boolean if a field has been set.

### GetSubject

`func (o *TrackAction) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TrackAction) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TrackAction) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TrackAction) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetNotes

`func (o *TrackAction) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TrackAction) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TrackAction) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TrackAction) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetActivityType

`func (o *TrackAction) GetActivityType() ActivityTypeReference`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *TrackAction) GetActivityTypeOk() (*ActivityTypeReference, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *TrackAction) SetActivityType(v ActivityTypeReference)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *TrackAction) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### GetActivityStatus

`func (o *TrackAction) GetActivityStatus() ActivityStatusReference`

GetActivityStatus returns the ActivityStatus field if non-nil, zero value otherwise.

### GetActivityStatusOk

`func (o *TrackAction) GetActivityStatusOk() (*ActivityStatusReference, bool)`

GetActivityStatusOk returns a tuple with the ActivityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityStatus

`func (o *TrackAction) SetActivityStatus(v ActivityStatusReference)`

SetActivityStatus sets ActivityStatus field to given value.

### HasActivityStatus

`func (o *TrackAction) HasActivityStatus() bool`

HasActivityStatus returns a boolean if a field has been set.

### GetCompanyStatus

`func (o *TrackAction) GetCompanyStatus() CompanyStatusReference`

GetCompanyStatus returns the CompanyStatus field if non-nil, zero value otherwise.

### GetCompanyStatusOk

`func (o *TrackAction) GetCompanyStatusOk() (*CompanyStatusReference, bool)`

GetCompanyStatusOk returns a tuple with the CompanyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyStatus

`func (o *TrackAction) SetCompanyStatus(v CompanyStatusReference)`

SetCompanyStatus sets CompanyStatus field to given value.

### HasCompanyStatus

`func (o *TrackAction) HasCompanyStatus() bool`

HasCompanyStatus returns a boolean if a field has been set.

### GetTrack

`func (o *TrackAction) GetTrack() TrackReference`

GetTrack returns the Track field if non-nil, zero value otherwise.

### GetTrackOk

`func (o *TrackAction) GetTrackOk() (*TrackReference, bool)`

GetTrackOk returns a tuple with the Track field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack

`func (o *TrackAction) SetTrack(v TrackReference)`

SetTrack sets Track field to given value.

### HasTrack

`func (o *TrackAction) HasTrack() bool`

HasTrack returns a boolean if a field has been set.

### GetAttachedTrack

`func (o *TrackAction) GetAttachedTrack() TrackReference`

GetAttachedTrack returns the AttachedTrack field if non-nil, zero value otherwise.

### GetAttachedTrackOk

`func (o *TrackAction) GetAttachedTrackOk() (*TrackReference, bool)`

GetAttachedTrackOk returns a tuple with the AttachedTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedTrack

`func (o *TrackAction) SetAttachedTrack(v TrackReference)`

SetAttachedTrack sets AttachedTrack field to given value.

### HasAttachedTrack

`func (o *TrackAction) HasAttachedTrack() bool`

HasAttachedTrack returns a boolean if a field has been set.

### GetGroup

`func (o *TrackAction) GetGroup() GroupReference`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *TrackAction) GetGroupOk() (*GroupReference, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *TrackAction) SetGroup(v GroupReference)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *TrackAction) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetCcContact

`func (o *TrackAction) GetCcContact() ContactReference`

GetCcContact returns the CcContact field if non-nil, zero value otherwise.

### GetCcContactOk

`func (o *TrackAction) GetCcContactOk() (*ContactReference, bool)`

GetCcContactOk returns a tuple with the CcContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcContact

`func (o *TrackAction) SetCcContact(v ContactReference)`

SetCcContact sets CcContact field to given value.

### HasCcContact

`func (o *TrackAction) HasCcContact() bool`

HasCcContact returns a boolean if a field has been set.

### GetBccContact

`func (o *TrackAction) GetBccContact() ContactReference`

GetBccContact returns the BccContact field if non-nil, zero value otherwise.

### GetBccContactOk

`func (o *TrackAction) GetBccContactOk() (*ContactReference, bool)`

GetBccContactOk returns a tuple with the BccContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBccContact

`func (o *TrackAction) SetBccContact(v ContactReference)`

SetBccContact sets BccContact field to given value.

### HasBccContact

`func (o *TrackAction) HasBccContact() bool`

HasBccContact returns a boolean if a field has been set.

### GetDaysToExecute

`func (o *TrackAction) GetDaysToExecute() int32`

GetDaysToExecute returns the DaysToExecute field if non-nil, zero value otherwise.

### GetDaysToExecuteOk

`func (o *TrackAction) GetDaysToExecuteOk() (*int32, bool)`

GetDaysToExecuteOk returns a tuple with the DaysToExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToExecute

`func (o *TrackAction) SetDaysToExecute(v int32)`

SetDaysToExecute sets DaysToExecute field to given value.

### HasDaysToExecute

`func (o *TrackAction) HasDaysToExecute() bool`

HasDaysToExecute returns a boolean if a field has been set.

### SetDaysToExecuteNil

`func (o *TrackAction) SetDaysToExecuteNil(b bool)`

 SetDaysToExecuteNil sets the value for DaysToExecute to be an explicit nil

### UnsetDaysToExecute
`func (o *TrackAction) UnsetDaysToExecute()`

UnsetDaysToExecute ensures that no value is present for DaysToExecute, not even an explicit nil
### GetNotifyWho

`func (o *TrackAction) GetNotifyWho() NotificationRecipientReference`

GetNotifyWho returns the NotifyWho field if non-nil, zero value otherwise.

### GetNotifyWhoOk

`func (o *TrackAction) GetNotifyWhoOk() (*NotificationRecipientReference, bool)`

GetNotifyWhoOk returns a tuple with the NotifyWho field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyWho

`func (o *TrackAction) SetNotifyWho(v NotificationRecipientReference)`

SetNotifyWho sets NotifyWho field to given value.

### HasNotifyWho

`func (o *TrackAction) HasNotifyWho() bool`

HasNotifyWho returns a boolean if a field has been set.

### GetNotifyFrom

`func (o *TrackAction) GetNotifyFrom() NotificationRecipientReference`

GetNotifyFrom returns the NotifyFrom field if non-nil, zero value otherwise.

### GetNotifyFromOk

`func (o *TrackAction) GetNotifyFromOk() (*NotificationRecipientReference, bool)`

GetNotifyFromOk returns a tuple with the NotifyFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyFrom

`func (o *TrackAction) SetNotifyFrom(v NotificationRecipientReference)`

SetNotifyFrom sets NotifyFrom field to given value.

### HasNotifyFrom

`func (o *TrackAction) HasNotifyFrom() bool`

HasNotifyFrom returns a boolean if a field has been set.

### GetInfo

`func (o *TrackAction) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TrackAction) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TrackAction) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TrackAction) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


