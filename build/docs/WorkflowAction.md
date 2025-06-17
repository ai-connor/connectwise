# WorkflowAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**NotifyType** | [**NotifyTypeReference**](NotifyTypeReference.md) |  | 
**NotifyWho** | Pointer to [**NotificationRecipientReference**](NotificationRecipientReference.md) |  | [optional] 
**SpecificMemberTo** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**EmailRecipient** | Pointer to **string** | Required when notifyWho is set to: \&quot;Email Address\&quot; Max length: 250; | [optional] 
**NotifyFrom** | Pointer to [**NotificationRecipientReference**](NotificationRecipientReference.md) |  | [optional] 
**SpecificMemberFrom** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**EmailFrom** | Pointer to **string** | Required when notifyFrom is set to: \&quot;Email Address\&quot; Max length: 250; | [optional] 
**CcContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**BccContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**Subject** | Pointer to **string** | Required when notifyType is set to: \&quot;Create Activity\&quot;, \&quot;Send Email\&quot;, \&quot;Assign Resource\&quot; Max length: 100; | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**ActivityStatus** | Pointer to [**ActivityStatusReference**](ActivityStatusReference.md) |  | [optional] 
**ActivityType** | Pointer to [**ActivityTypeReference**](ActivityTypeReference.md) |  | [optional] 
**AttachedTrack** | Pointer to [**TrackReference**](TrackReference.md) |  | [optional] 
**DaysToExecute** | Pointer to **NullableInt32** |  | [optional] 
**Board** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**BoardStatus** | Pointer to [**ServiceStatusReference**](ServiceStatusReference.md) |  | [optional] 
**ServiceType** | Pointer to [**ServiceTypeReference**](ServiceTypeReference.md) |  | [optional] 
**ServiceSubType** | Pointer to [**ServiceSubTypeReference**](ServiceSubTypeReference.md) |  | [optional] 
**ServiceItem** | Pointer to [**ServiceItemReference**](ServiceItemReference.md) |  | [optional] 
**Group** | Pointer to [**GroupReference**](GroupReference.md) |  | [optional] 
**ServiceTemplate** | Pointer to [**ServiceTemplateReference**](ServiceTemplateReference.md) |  | [optional] 
**InvoiceMinDays** | Pointer to **NullableInt32** |  | [optional] 
**AutomateScript** | Pointer to [**AutomateScriptReference**](AutomateScriptReference.md) |  | [optional] 
**ScriptSuccessStatus** | Pointer to [**ServiceStatusReference**](ServiceStatusReference.md) |  | [optional] 
**ScriptFailStatus** | Pointer to [**ServiceStatusReference**](ServiceStatusReference.md) |  | [optional] 
**DetailNotesFlag** | Pointer to **NullableBool** |  | [optional] 
**InternalNotesFlag** | Pointer to **NullableBool** |  | [optional] 
**AuditNotesFlag** | Pointer to **NullableBool** |  | [optional] 
**ServicePriority** | Pointer to [**PriorityReference**](PriorityReference.md) |  | [optional] 
**UpdateOwnerFlag** | Pointer to **NullableBool** |  | [optional] 
**SalesOrderStatus** | Pointer to [**OrderStatusReference**](OrderStatusReference.md) |  | [optional] 
**ProjectStatus** | Pointer to [**ProjectStatusReference**](ProjectStatusReference.md) |  | [optional] 
**CompanyStatus** | Pointer to [**CompanyStatusReference**](CompanyStatusReference.md) |  | [optional] 
**Attachments** | Pointer to **[]int32** |  | [optional] 
**ServiceSurvey** | Pointer to [**ServiceSurveyReference**](ServiceSurveyReference.md) |  | [optional] 
**SpecificTeamTo** | Pointer to [**GenericBoardTeamReference**](GenericBoardTeamReference.md) |  | [optional] 
**AttachConfigurationsFor** | Pointer to **NullableString** | Required when notifyType is set to: \&quot;Attach Configuration\&quot; | [optional] 
**ConfigurationType** | Pointer to [**ConfigurationTypeReference**](ConfigurationTypeReference.md) |  | [optional] 
**ConfigurationStatus** | Pointer to [**ConfigurationStatusReference**](ConfigurationStatusReference.md) |  | [optional] 
**ConnectWiseID** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **NullableInt32** | WF_NotifyEvents_RecID | [optional] 
**GrandParentId** | Pointer to **NullableInt32** | WF_NotifyHeader_RecID | [optional] 
**ParentConnectWiseId** | Pointer to **string** |  | [optional] 
**GrandParentConnectWiseId** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewWorkflowAction

`func NewWorkflowAction(notifyType NotifyTypeReference, ) *WorkflowAction`

NewWorkflowAction instantiates a new WorkflowAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowActionWithDefaults

`func NewWorkflowActionWithDefaults() *WorkflowAction`

NewWorkflowActionWithDefaults instantiates a new WorkflowAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowAction) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowAction) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowAction) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowAction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotifyType

`func (o *WorkflowAction) GetNotifyType() NotifyTypeReference`

GetNotifyType returns the NotifyType field if non-nil, zero value otherwise.

### GetNotifyTypeOk

`func (o *WorkflowAction) GetNotifyTypeOk() (*NotifyTypeReference, bool)`

GetNotifyTypeOk returns a tuple with the NotifyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyType

`func (o *WorkflowAction) SetNotifyType(v NotifyTypeReference)`

SetNotifyType sets NotifyType field to given value.


### GetNotifyWho

`func (o *WorkflowAction) GetNotifyWho() NotificationRecipientReference`

GetNotifyWho returns the NotifyWho field if non-nil, zero value otherwise.

### GetNotifyWhoOk

`func (o *WorkflowAction) GetNotifyWhoOk() (*NotificationRecipientReference, bool)`

GetNotifyWhoOk returns a tuple with the NotifyWho field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyWho

`func (o *WorkflowAction) SetNotifyWho(v NotificationRecipientReference)`

SetNotifyWho sets NotifyWho field to given value.

### HasNotifyWho

`func (o *WorkflowAction) HasNotifyWho() bool`

HasNotifyWho returns a boolean if a field has been set.

### GetSpecificMemberTo

`func (o *WorkflowAction) GetSpecificMemberTo() MemberReference`

GetSpecificMemberTo returns the SpecificMemberTo field if non-nil, zero value otherwise.

### GetSpecificMemberToOk

`func (o *WorkflowAction) GetSpecificMemberToOk() (*MemberReference, bool)`

GetSpecificMemberToOk returns a tuple with the SpecificMemberTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificMemberTo

`func (o *WorkflowAction) SetSpecificMemberTo(v MemberReference)`

SetSpecificMemberTo sets SpecificMemberTo field to given value.

### HasSpecificMemberTo

`func (o *WorkflowAction) HasSpecificMemberTo() bool`

HasSpecificMemberTo returns a boolean if a field has been set.

### GetEmailRecipient

`func (o *WorkflowAction) GetEmailRecipient() string`

GetEmailRecipient returns the EmailRecipient field if non-nil, zero value otherwise.

### GetEmailRecipientOk

`func (o *WorkflowAction) GetEmailRecipientOk() (*string, bool)`

GetEmailRecipientOk returns a tuple with the EmailRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailRecipient

`func (o *WorkflowAction) SetEmailRecipient(v string)`

SetEmailRecipient sets EmailRecipient field to given value.

### HasEmailRecipient

`func (o *WorkflowAction) HasEmailRecipient() bool`

HasEmailRecipient returns a boolean if a field has been set.

### GetNotifyFrom

`func (o *WorkflowAction) GetNotifyFrom() NotificationRecipientReference`

GetNotifyFrom returns the NotifyFrom field if non-nil, zero value otherwise.

### GetNotifyFromOk

`func (o *WorkflowAction) GetNotifyFromOk() (*NotificationRecipientReference, bool)`

GetNotifyFromOk returns a tuple with the NotifyFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyFrom

`func (o *WorkflowAction) SetNotifyFrom(v NotificationRecipientReference)`

SetNotifyFrom sets NotifyFrom field to given value.

### HasNotifyFrom

`func (o *WorkflowAction) HasNotifyFrom() bool`

HasNotifyFrom returns a boolean if a field has been set.

### GetSpecificMemberFrom

`func (o *WorkflowAction) GetSpecificMemberFrom() MemberReference`

GetSpecificMemberFrom returns the SpecificMemberFrom field if non-nil, zero value otherwise.

### GetSpecificMemberFromOk

`func (o *WorkflowAction) GetSpecificMemberFromOk() (*MemberReference, bool)`

GetSpecificMemberFromOk returns a tuple with the SpecificMemberFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificMemberFrom

`func (o *WorkflowAction) SetSpecificMemberFrom(v MemberReference)`

SetSpecificMemberFrom sets SpecificMemberFrom field to given value.

### HasSpecificMemberFrom

`func (o *WorkflowAction) HasSpecificMemberFrom() bool`

HasSpecificMemberFrom returns a boolean if a field has been set.

### GetEmailFrom

`func (o *WorkflowAction) GetEmailFrom() string`

GetEmailFrom returns the EmailFrom field if non-nil, zero value otherwise.

### GetEmailFromOk

`func (o *WorkflowAction) GetEmailFromOk() (*string, bool)`

GetEmailFromOk returns a tuple with the EmailFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFrom

`func (o *WorkflowAction) SetEmailFrom(v string)`

SetEmailFrom sets EmailFrom field to given value.

### HasEmailFrom

`func (o *WorkflowAction) HasEmailFrom() bool`

HasEmailFrom returns a boolean if a field has been set.

### GetCcContact

`func (o *WorkflowAction) GetCcContact() ContactReference`

GetCcContact returns the CcContact field if non-nil, zero value otherwise.

### GetCcContactOk

`func (o *WorkflowAction) GetCcContactOk() (*ContactReference, bool)`

GetCcContactOk returns a tuple with the CcContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcContact

`func (o *WorkflowAction) SetCcContact(v ContactReference)`

SetCcContact sets CcContact field to given value.

### HasCcContact

`func (o *WorkflowAction) HasCcContact() bool`

HasCcContact returns a boolean if a field has been set.

### GetBccContact

`func (o *WorkflowAction) GetBccContact() ContactReference`

GetBccContact returns the BccContact field if non-nil, zero value otherwise.

### GetBccContactOk

`func (o *WorkflowAction) GetBccContactOk() (*ContactReference, bool)`

GetBccContactOk returns a tuple with the BccContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBccContact

`func (o *WorkflowAction) SetBccContact(v ContactReference)`

SetBccContact sets BccContact field to given value.

### HasBccContact

`func (o *WorkflowAction) HasBccContact() bool`

HasBccContact returns a boolean if a field has been set.

### GetSubject

`func (o *WorkflowAction) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *WorkflowAction) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *WorkflowAction) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *WorkflowAction) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetNotes

`func (o *WorkflowAction) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *WorkflowAction) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *WorkflowAction) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *WorkflowAction) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetActivityStatus

`func (o *WorkflowAction) GetActivityStatus() ActivityStatusReference`

GetActivityStatus returns the ActivityStatus field if non-nil, zero value otherwise.

### GetActivityStatusOk

`func (o *WorkflowAction) GetActivityStatusOk() (*ActivityStatusReference, bool)`

GetActivityStatusOk returns a tuple with the ActivityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityStatus

`func (o *WorkflowAction) SetActivityStatus(v ActivityStatusReference)`

SetActivityStatus sets ActivityStatus field to given value.

### HasActivityStatus

`func (o *WorkflowAction) HasActivityStatus() bool`

HasActivityStatus returns a boolean if a field has been set.

### GetActivityType

`func (o *WorkflowAction) GetActivityType() ActivityTypeReference`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *WorkflowAction) GetActivityTypeOk() (*ActivityTypeReference, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *WorkflowAction) SetActivityType(v ActivityTypeReference)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *WorkflowAction) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### GetAttachedTrack

`func (o *WorkflowAction) GetAttachedTrack() TrackReference`

GetAttachedTrack returns the AttachedTrack field if non-nil, zero value otherwise.

### GetAttachedTrackOk

`func (o *WorkflowAction) GetAttachedTrackOk() (*TrackReference, bool)`

GetAttachedTrackOk returns a tuple with the AttachedTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedTrack

`func (o *WorkflowAction) SetAttachedTrack(v TrackReference)`

SetAttachedTrack sets AttachedTrack field to given value.

### HasAttachedTrack

`func (o *WorkflowAction) HasAttachedTrack() bool`

HasAttachedTrack returns a boolean if a field has been set.

### GetDaysToExecute

`func (o *WorkflowAction) GetDaysToExecute() int32`

GetDaysToExecute returns the DaysToExecute field if non-nil, zero value otherwise.

### GetDaysToExecuteOk

`func (o *WorkflowAction) GetDaysToExecuteOk() (*int32, bool)`

GetDaysToExecuteOk returns a tuple with the DaysToExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToExecute

`func (o *WorkflowAction) SetDaysToExecute(v int32)`

SetDaysToExecute sets DaysToExecute field to given value.

### HasDaysToExecute

`func (o *WorkflowAction) HasDaysToExecute() bool`

HasDaysToExecute returns a boolean if a field has been set.

### SetDaysToExecuteNil

`func (o *WorkflowAction) SetDaysToExecuteNil(b bool)`

 SetDaysToExecuteNil sets the value for DaysToExecute to be an explicit nil

### UnsetDaysToExecute
`func (o *WorkflowAction) UnsetDaysToExecute()`

UnsetDaysToExecute ensures that no value is present for DaysToExecute, not even an explicit nil
### GetBoard

`func (o *WorkflowAction) GetBoard() BoardReference`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *WorkflowAction) GetBoardOk() (*BoardReference, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *WorkflowAction) SetBoard(v BoardReference)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *WorkflowAction) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetBoardStatus

`func (o *WorkflowAction) GetBoardStatus() ServiceStatusReference`

GetBoardStatus returns the BoardStatus field if non-nil, zero value otherwise.

### GetBoardStatusOk

`func (o *WorkflowAction) GetBoardStatusOk() (*ServiceStatusReference, bool)`

GetBoardStatusOk returns a tuple with the BoardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardStatus

`func (o *WorkflowAction) SetBoardStatus(v ServiceStatusReference)`

SetBoardStatus sets BoardStatus field to given value.

### HasBoardStatus

`func (o *WorkflowAction) HasBoardStatus() bool`

HasBoardStatus returns a boolean if a field has been set.

### GetServiceType

`func (o *WorkflowAction) GetServiceType() ServiceTypeReference`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *WorkflowAction) GetServiceTypeOk() (*ServiceTypeReference, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *WorkflowAction) SetServiceType(v ServiceTypeReference)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *WorkflowAction) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetServiceSubType

`func (o *WorkflowAction) GetServiceSubType() ServiceSubTypeReference`

GetServiceSubType returns the ServiceSubType field if non-nil, zero value otherwise.

### GetServiceSubTypeOk

`func (o *WorkflowAction) GetServiceSubTypeOk() (*ServiceSubTypeReference, bool)`

GetServiceSubTypeOk returns a tuple with the ServiceSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSubType

`func (o *WorkflowAction) SetServiceSubType(v ServiceSubTypeReference)`

SetServiceSubType sets ServiceSubType field to given value.

### HasServiceSubType

`func (o *WorkflowAction) HasServiceSubType() bool`

HasServiceSubType returns a boolean if a field has been set.

### GetServiceItem

`func (o *WorkflowAction) GetServiceItem() ServiceItemReference`

GetServiceItem returns the ServiceItem field if non-nil, zero value otherwise.

### GetServiceItemOk

`func (o *WorkflowAction) GetServiceItemOk() (*ServiceItemReference, bool)`

GetServiceItemOk returns a tuple with the ServiceItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItem

`func (o *WorkflowAction) SetServiceItem(v ServiceItemReference)`

SetServiceItem sets ServiceItem field to given value.

### HasServiceItem

`func (o *WorkflowAction) HasServiceItem() bool`

HasServiceItem returns a boolean if a field has been set.

### GetGroup

`func (o *WorkflowAction) GetGroup() GroupReference`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *WorkflowAction) GetGroupOk() (*GroupReference, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *WorkflowAction) SetGroup(v GroupReference)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *WorkflowAction) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetServiceTemplate

`func (o *WorkflowAction) GetServiceTemplate() ServiceTemplateReference`

GetServiceTemplate returns the ServiceTemplate field if non-nil, zero value otherwise.

### GetServiceTemplateOk

`func (o *WorkflowAction) GetServiceTemplateOk() (*ServiceTemplateReference, bool)`

GetServiceTemplateOk returns a tuple with the ServiceTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTemplate

`func (o *WorkflowAction) SetServiceTemplate(v ServiceTemplateReference)`

SetServiceTemplate sets ServiceTemplate field to given value.

### HasServiceTemplate

`func (o *WorkflowAction) HasServiceTemplate() bool`

HasServiceTemplate returns a boolean if a field has been set.

### GetInvoiceMinDays

`func (o *WorkflowAction) GetInvoiceMinDays() int32`

GetInvoiceMinDays returns the InvoiceMinDays field if non-nil, zero value otherwise.

### GetInvoiceMinDaysOk

`func (o *WorkflowAction) GetInvoiceMinDaysOk() (*int32, bool)`

GetInvoiceMinDaysOk returns a tuple with the InvoiceMinDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceMinDays

`func (o *WorkflowAction) SetInvoiceMinDays(v int32)`

SetInvoiceMinDays sets InvoiceMinDays field to given value.

### HasInvoiceMinDays

`func (o *WorkflowAction) HasInvoiceMinDays() bool`

HasInvoiceMinDays returns a boolean if a field has been set.

### SetInvoiceMinDaysNil

`func (o *WorkflowAction) SetInvoiceMinDaysNil(b bool)`

 SetInvoiceMinDaysNil sets the value for InvoiceMinDays to be an explicit nil

### UnsetInvoiceMinDays
`func (o *WorkflowAction) UnsetInvoiceMinDays()`

UnsetInvoiceMinDays ensures that no value is present for InvoiceMinDays, not even an explicit nil
### GetAutomateScript

`func (o *WorkflowAction) GetAutomateScript() AutomateScriptReference`

GetAutomateScript returns the AutomateScript field if non-nil, zero value otherwise.

### GetAutomateScriptOk

`func (o *WorkflowAction) GetAutomateScriptOk() (*AutomateScriptReference, bool)`

GetAutomateScriptOk returns a tuple with the AutomateScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomateScript

`func (o *WorkflowAction) SetAutomateScript(v AutomateScriptReference)`

SetAutomateScript sets AutomateScript field to given value.

### HasAutomateScript

`func (o *WorkflowAction) HasAutomateScript() bool`

HasAutomateScript returns a boolean if a field has been set.

### GetScriptSuccessStatus

`func (o *WorkflowAction) GetScriptSuccessStatus() ServiceStatusReference`

GetScriptSuccessStatus returns the ScriptSuccessStatus field if non-nil, zero value otherwise.

### GetScriptSuccessStatusOk

`func (o *WorkflowAction) GetScriptSuccessStatusOk() (*ServiceStatusReference, bool)`

GetScriptSuccessStatusOk returns a tuple with the ScriptSuccessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSuccessStatus

`func (o *WorkflowAction) SetScriptSuccessStatus(v ServiceStatusReference)`

SetScriptSuccessStatus sets ScriptSuccessStatus field to given value.

### HasScriptSuccessStatus

`func (o *WorkflowAction) HasScriptSuccessStatus() bool`

HasScriptSuccessStatus returns a boolean if a field has been set.

### GetScriptFailStatus

`func (o *WorkflowAction) GetScriptFailStatus() ServiceStatusReference`

GetScriptFailStatus returns the ScriptFailStatus field if non-nil, zero value otherwise.

### GetScriptFailStatusOk

`func (o *WorkflowAction) GetScriptFailStatusOk() (*ServiceStatusReference, bool)`

GetScriptFailStatusOk returns a tuple with the ScriptFailStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptFailStatus

`func (o *WorkflowAction) SetScriptFailStatus(v ServiceStatusReference)`

SetScriptFailStatus sets ScriptFailStatus field to given value.

### HasScriptFailStatus

`func (o *WorkflowAction) HasScriptFailStatus() bool`

HasScriptFailStatus returns a boolean if a field has been set.

### GetDetailNotesFlag

`func (o *WorkflowAction) GetDetailNotesFlag() bool`

GetDetailNotesFlag returns the DetailNotesFlag field if non-nil, zero value otherwise.

### GetDetailNotesFlagOk

`func (o *WorkflowAction) GetDetailNotesFlagOk() (*bool, bool)`

GetDetailNotesFlagOk returns a tuple with the DetailNotesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailNotesFlag

`func (o *WorkflowAction) SetDetailNotesFlag(v bool)`

SetDetailNotesFlag sets DetailNotesFlag field to given value.

### HasDetailNotesFlag

`func (o *WorkflowAction) HasDetailNotesFlag() bool`

HasDetailNotesFlag returns a boolean if a field has been set.

### SetDetailNotesFlagNil

`func (o *WorkflowAction) SetDetailNotesFlagNil(b bool)`

 SetDetailNotesFlagNil sets the value for DetailNotesFlag to be an explicit nil

### UnsetDetailNotesFlag
`func (o *WorkflowAction) UnsetDetailNotesFlag()`

UnsetDetailNotesFlag ensures that no value is present for DetailNotesFlag, not even an explicit nil
### GetInternalNotesFlag

`func (o *WorkflowAction) GetInternalNotesFlag() bool`

GetInternalNotesFlag returns the InternalNotesFlag field if non-nil, zero value otherwise.

### GetInternalNotesFlagOk

`func (o *WorkflowAction) GetInternalNotesFlagOk() (*bool, bool)`

GetInternalNotesFlagOk returns a tuple with the InternalNotesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNotesFlag

`func (o *WorkflowAction) SetInternalNotesFlag(v bool)`

SetInternalNotesFlag sets InternalNotesFlag field to given value.

### HasInternalNotesFlag

`func (o *WorkflowAction) HasInternalNotesFlag() bool`

HasInternalNotesFlag returns a boolean if a field has been set.

### SetInternalNotesFlagNil

`func (o *WorkflowAction) SetInternalNotesFlagNil(b bool)`

 SetInternalNotesFlagNil sets the value for InternalNotesFlag to be an explicit nil

### UnsetInternalNotesFlag
`func (o *WorkflowAction) UnsetInternalNotesFlag()`

UnsetInternalNotesFlag ensures that no value is present for InternalNotesFlag, not even an explicit nil
### GetAuditNotesFlag

`func (o *WorkflowAction) GetAuditNotesFlag() bool`

GetAuditNotesFlag returns the AuditNotesFlag field if non-nil, zero value otherwise.

### GetAuditNotesFlagOk

`func (o *WorkflowAction) GetAuditNotesFlagOk() (*bool, bool)`

GetAuditNotesFlagOk returns a tuple with the AuditNotesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditNotesFlag

`func (o *WorkflowAction) SetAuditNotesFlag(v bool)`

SetAuditNotesFlag sets AuditNotesFlag field to given value.

### HasAuditNotesFlag

`func (o *WorkflowAction) HasAuditNotesFlag() bool`

HasAuditNotesFlag returns a boolean if a field has been set.

### SetAuditNotesFlagNil

`func (o *WorkflowAction) SetAuditNotesFlagNil(b bool)`

 SetAuditNotesFlagNil sets the value for AuditNotesFlag to be an explicit nil

### UnsetAuditNotesFlag
`func (o *WorkflowAction) UnsetAuditNotesFlag()`

UnsetAuditNotesFlag ensures that no value is present for AuditNotesFlag, not even an explicit nil
### GetServicePriority

`func (o *WorkflowAction) GetServicePriority() PriorityReference`

GetServicePriority returns the ServicePriority field if non-nil, zero value otherwise.

### GetServicePriorityOk

`func (o *WorkflowAction) GetServicePriorityOk() (*PriorityReference, bool)`

GetServicePriorityOk returns a tuple with the ServicePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePriority

`func (o *WorkflowAction) SetServicePriority(v PriorityReference)`

SetServicePriority sets ServicePriority field to given value.

### HasServicePriority

`func (o *WorkflowAction) HasServicePriority() bool`

HasServicePriority returns a boolean if a field has been set.

### GetUpdateOwnerFlag

`func (o *WorkflowAction) GetUpdateOwnerFlag() bool`

GetUpdateOwnerFlag returns the UpdateOwnerFlag field if non-nil, zero value otherwise.

### GetUpdateOwnerFlagOk

`func (o *WorkflowAction) GetUpdateOwnerFlagOk() (*bool, bool)`

GetUpdateOwnerFlagOk returns a tuple with the UpdateOwnerFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateOwnerFlag

`func (o *WorkflowAction) SetUpdateOwnerFlag(v bool)`

SetUpdateOwnerFlag sets UpdateOwnerFlag field to given value.

### HasUpdateOwnerFlag

`func (o *WorkflowAction) HasUpdateOwnerFlag() bool`

HasUpdateOwnerFlag returns a boolean if a field has been set.

### SetUpdateOwnerFlagNil

`func (o *WorkflowAction) SetUpdateOwnerFlagNil(b bool)`

 SetUpdateOwnerFlagNil sets the value for UpdateOwnerFlag to be an explicit nil

### UnsetUpdateOwnerFlag
`func (o *WorkflowAction) UnsetUpdateOwnerFlag()`

UnsetUpdateOwnerFlag ensures that no value is present for UpdateOwnerFlag, not even an explicit nil
### GetSalesOrderStatus

`func (o *WorkflowAction) GetSalesOrderStatus() OrderStatusReference`

GetSalesOrderStatus returns the SalesOrderStatus field if non-nil, zero value otherwise.

### GetSalesOrderStatusOk

`func (o *WorkflowAction) GetSalesOrderStatusOk() (*OrderStatusReference, bool)`

GetSalesOrderStatusOk returns a tuple with the SalesOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrderStatus

`func (o *WorkflowAction) SetSalesOrderStatus(v OrderStatusReference)`

SetSalesOrderStatus sets SalesOrderStatus field to given value.

### HasSalesOrderStatus

`func (o *WorkflowAction) HasSalesOrderStatus() bool`

HasSalesOrderStatus returns a boolean if a field has been set.

### GetProjectStatus

`func (o *WorkflowAction) GetProjectStatus() ProjectStatusReference`

GetProjectStatus returns the ProjectStatus field if non-nil, zero value otherwise.

### GetProjectStatusOk

`func (o *WorkflowAction) GetProjectStatusOk() (*ProjectStatusReference, bool)`

GetProjectStatusOk returns a tuple with the ProjectStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectStatus

`func (o *WorkflowAction) SetProjectStatus(v ProjectStatusReference)`

SetProjectStatus sets ProjectStatus field to given value.

### HasProjectStatus

`func (o *WorkflowAction) HasProjectStatus() bool`

HasProjectStatus returns a boolean if a field has been set.

### GetCompanyStatus

`func (o *WorkflowAction) GetCompanyStatus() CompanyStatusReference`

GetCompanyStatus returns the CompanyStatus field if non-nil, zero value otherwise.

### GetCompanyStatusOk

`func (o *WorkflowAction) GetCompanyStatusOk() (*CompanyStatusReference, bool)`

GetCompanyStatusOk returns a tuple with the CompanyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyStatus

`func (o *WorkflowAction) SetCompanyStatus(v CompanyStatusReference)`

SetCompanyStatus sets CompanyStatus field to given value.

### HasCompanyStatus

`func (o *WorkflowAction) HasCompanyStatus() bool`

HasCompanyStatus returns a boolean if a field has been set.

### GetAttachments

`func (o *WorkflowAction) GetAttachments() []int32`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *WorkflowAction) GetAttachmentsOk() (*[]int32, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *WorkflowAction) SetAttachments(v []int32)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *WorkflowAction) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetServiceSurvey

`func (o *WorkflowAction) GetServiceSurvey() ServiceSurveyReference`

GetServiceSurvey returns the ServiceSurvey field if non-nil, zero value otherwise.

### GetServiceSurveyOk

`func (o *WorkflowAction) GetServiceSurveyOk() (*ServiceSurveyReference, bool)`

GetServiceSurveyOk returns a tuple with the ServiceSurvey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSurvey

`func (o *WorkflowAction) SetServiceSurvey(v ServiceSurveyReference)`

SetServiceSurvey sets ServiceSurvey field to given value.

### HasServiceSurvey

`func (o *WorkflowAction) HasServiceSurvey() bool`

HasServiceSurvey returns a boolean if a field has been set.

### GetSpecificTeamTo

`func (o *WorkflowAction) GetSpecificTeamTo() GenericBoardTeamReference`

GetSpecificTeamTo returns the SpecificTeamTo field if non-nil, zero value otherwise.

### GetSpecificTeamToOk

`func (o *WorkflowAction) GetSpecificTeamToOk() (*GenericBoardTeamReference, bool)`

GetSpecificTeamToOk returns a tuple with the SpecificTeamTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificTeamTo

`func (o *WorkflowAction) SetSpecificTeamTo(v GenericBoardTeamReference)`

SetSpecificTeamTo sets SpecificTeamTo field to given value.

### HasSpecificTeamTo

`func (o *WorkflowAction) HasSpecificTeamTo() bool`

HasSpecificTeamTo returns a boolean if a field has been set.

### GetAttachConfigurationsFor

`func (o *WorkflowAction) GetAttachConfigurationsFor() string`

GetAttachConfigurationsFor returns the AttachConfigurationsFor field if non-nil, zero value otherwise.

### GetAttachConfigurationsForOk

`func (o *WorkflowAction) GetAttachConfigurationsForOk() (*string, bool)`

GetAttachConfigurationsForOk returns a tuple with the AttachConfigurationsFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachConfigurationsFor

`func (o *WorkflowAction) SetAttachConfigurationsFor(v string)`

SetAttachConfigurationsFor sets AttachConfigurationsFor field to given value.

### HasAttachConfigurationsFor

`func (o *WorkflowAction) HasAttachConfigurationsFor() bool`

HasAttachConfigurationsFor returns a boolean if a field has been set.

### SetAttachConfigurationsForNil

`func (o *WorkflowAction) SetAttachConfigurationsForNil(b bool)`

 SetAttachConfigurationsForNil sets the value for AttachConfigurationsFor to be an explicit nil

### UnsetAttachConfigurationsFor
`func (o *WorkflowAction) UnsetAttachConfigurationsFor()`

UnsetAttachConfigurationsFor ensures that no value is present for AttachConfigurationsFor, not even an explicit nil
### GetConfigurationType

`func (o *WorkflowAction) GetConfigurationType() ConfigurationTypeReference`

GetConfigurationType returns the ConfigurationType field if non-nil, zero value otherwise.

### GetConfigurationTypeOk

`func (o *WorkflowAction) GetConfigurationTypeOk() (*ConfigurationTypeReference, bool)`

GetConfigurationTypeOk returns a tuple with the ConfigurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationType

`func (o *WorkflowAction) SetConfigurationType(v ConfigurationTypeReference)`

SetConfigurationType sets ConfigurationType field to given value.

### HasConfigurationType

`func (o *WorkflowAction) HasConfigurationType() bool`

HasConfigurationType returns a boolean if a field has been set.

### GetConfigurationStatus

`func (o *WorkflowAction) GetConfigurationStatus() ConfigurationStatusReference`

GetConfigurationStatus returns the ConfigurationStatus field if non-nil, zero value otherwise.

### GetConfigurationStatusOk

`func (o *WorkflowAction) GetConfigurationStatusOk() (*ConfigurationStatusReference, bool)`

GetConfigurationStatusOk returns a tuple with the ConfigurationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationStatus

`func (o *WorkflowAction) SetConfigurationStatus(v ConfigurationStatusReference)`

SetConfigurationStatus sets ConfigurationStatus field to given value.

### HasConfigurationStatus

`func (o *WorkflowAction) HasConfigurationStatus() bool`

HasConfigurationStatus returns a boolean if a field has been set.

### GetConnectWiseID

`func (o *WorkflowAction) GetConnectWiseID() string`

GetConnectWiseID returns the ConnectWiseID field if non-nil, zero value otherwise.

### GetConnectWiseIDOk

`func (o *WorkflowAction) GetConnectWiseIDOk() (*string, bool)`

GetConnectWiseIDOk returns a tuple with the ConnectWiseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectWiseID

`func (o *WorkflowAction) SetConnectWiseID(v string)`

SetConnectWiseID sets ConnectWiseID field to given value.

### HasConnectWiseID

`func (o *WorkflowAction) HasConnectWiseID() bool`

HasConnectWiseID returns a boolean if a field has been set.

### GetParentId

`func (o *WorkflowAction) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *WorkflowAction) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *WorkflowAction) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *WorkflowAction) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *WorkflowAction) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *WorkflowAction) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetGrandParentId

`func (o *WorkflowAction) GetGrandParentId() int32`

GetGrandParentId returns the GrandParentId field if non-nil, zero value otherwise.

### GetGrandParentIdOk

`func (o *WorkflowAction) GetGrandParentIdOk() (*int32, bool)`

GetGrandParentIdOk returns a tuple with the GrandParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrandParentId

`func (o *WorkflowAction) SetGrandParentId(v int32)`

SetGrandParentId sets GrandParentId field to given value.

### HasGrandParentId

`func (o *WorkflowAction) HasGrandParentId() bool`

HasGrandParentId returns a boolean if a field has been set.

### SetGrandParentIdNil

`func (o *WorkflowAction) SetGrandParentIdNil(b bool)`

 SetGrandParentIdNil sets the value for GrandParentId to be an explicit nil

### UnsetGrandParentId
`func (o *WorkflowAction) UnsetGrandParentId()`

UnsetGrandParentId ensures that no value is present for GrandParentId, not even an explicit nil
### GetParentConnectWiseId

`func (o *WorkflowAction) GetParentConnectWiseId() string`

GetParentConnectWiseId returns the ParentConnectWiseId field if non-nil, zero value otherwise.

### GetParentConnectWiseIdOk

`func (o *WorkflowAction) GetParentConnectWiseIdOk() (*string, bool)`

GetParentConnectWiseIdOk returns a tuple with the ParentConnectWiseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConnectWiseId

`func (o *WorkflowAction) SetParentConnectWiseId(v string)`

SetParentConnectWiseId sets ParentConnectWiseId field to given value.

### HasParentConnectWiseId

`func (o *WorkflowAction) HasParentConnectWiseId() bool`

HasParentConnectWiseId returns a boolean if a field has been set.

### GetGrandParentConnectWiseId

`func (o *WorkflowAction) GetGrandParentConnectWiseId() string`

GetGrandParentConnectWiseId returns the GrandParentConnectWiseId field if non-nil, zero value otherwise.

### GetGrandParentConnectWiseIdOk

`func (o *WorkflowAction) GetGrandParentConnectWiseIdOk() (*string, bool)`

GetGrandParentConnectWiseIdOk returns a tuple with the GrandParentConnectWiseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrandParentConnectWiseId

`func (o *WorkflowAction) SetGrandParentConnectWiseId(v string)`

SetGrandParentConnectWiseId sets GrandParentConnectWiseId field to given value.

### HasGrandParentConnectWiseId

`func (o *WorkflowAction) HasGrandParentConnectWiseId() bool`

HasGrandParentConnectWiseId returns a boolean if a field has been set.

### GetInfo

`func (o *WorkflowAction) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *WorkflowAction) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *WorkflowAction) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *WorkflowAction) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


