# Activity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 100; | 
**Type** | Pointer to [**ActivityTypeReference**](ActivityTypeReference.md) |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Contact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**PhoneNumber** | Pointer to **string** |  Max length: 30; | [optional] 
**Email** | Pointer to **string** |  Max length: 250; | [optional] 
**Status** | Pointer to [**ActivityStatusReference**](ActivityStatusReference.md) |  | [optional] 
**Opportunity** | Pointer to [**OpportunityReference**](OpportunityReference.md) |  | [optional] 
**Ticket** | Pointer to [**TicketReference**](TicketReference.md) |  | [optional] 
**Agreement** | Pointer to [**AgreementReference**](AgreementReference.md) |  | [optional] 
**Campaign** | Pointer to [**CampaignReference**](CampaignReference.md) |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**DateStart** | Pointer to **time.Time** |  | [optional] 
**DateEnd** | Pointer to **time.Time** |  | [optional] 
**AssignedBy** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**AssignTo** | [**MemberReference**](MemberReference.md) |  | 
**ScheduleStatus** | Pointer to [**ScheduleStatusReference**](ScheduleStatusReference.md) |  | [optional] 
**Reminder** | Pointer to [**ReminderReference**](ReminderReference.md) |  | [optional] 
**Where** | Pointer to [**ServiceLocationReference**](ServiceLocationReference.md) |  | [optional] 
**NotifyFlag** | Pointer to **NullableBool** |  | [optional] 
**MobileGuid** | Pointer to **NullableString** |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewActivity

`func NewActivity(name string, assignTo MemberReference, ) *Activity`

NewActivity instantiates a new Activity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityWithDefaults

`func NewActivityWithDefaults() *Activity`

NewActivityWithDefaults instantiates a new Activity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Activity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Activity) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Activity) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Activity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Activity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Activity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Activity) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Activity) GetType() ActivityTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Activity) GetTypeOk() (*ActivityTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Activity) SetType(v ActivityTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *Activity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCompany

`func (o *Activity) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Activity) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Activity) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Activity) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetContact

`func (o *Activity) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Activity) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Activity) SetContact(v ContactReference)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Activity) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *Activity) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Activity) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Activity) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Activity) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetEmail

`func (o *Activity) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Activity) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Activity) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Activity) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetStatus

`func (o *Activity) GetStatus() ActivityStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Activity) GetStatusOk() (*ActivityStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Activity) SetStatus(v ActivityStatusReference)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Activity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOpportunity

`func (o *Activity) GetOpportunity() OpportunityReference`

GetOpportunity returns the Opportunity field if non-nil, zero value otherwise.

### GetOpportunityOk

`func (o *Activity) GetOpportunityOk() (*OpportunityReference, bool)`

GetOpportunityOk returns a tuple with the Opportunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunity

`func (o *Activity) SetOpportunity(v OpportunityReference)`

SetOpportunity sets Opportunity field to given value.

### HasOpportunity

`func (o *Activity) HasOpportunity() bool`

HasOpportunity returns a boolean if a field has been set.

### GetTicket

`func (o *Activity) GetTicket() TicketReference`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *Activity) GetTicketOk() (*TicketReference, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *Activity) SetTicket(v TicketReference)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *Activity) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetAgreement

`func (o *Activity) GetAgreement() AgreementReference`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *Activity) GetAgreementOk() (*AgreementReference, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *Activity) SetAgreement(v AgreementReference)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *Activity) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetCampaign

`func (o *Activity) GetCampaign() CampaignReference`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *Activity) GetCampaignOk() (*CampaignReference, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *Activity) SetCampaign(v CampaignReference)`

SetCampaign sets Campaign field to given value.

### HasCampaign

`func (o *Activity) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### GetNotes

`func (o *Activity) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Activity) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Activity) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Activity) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetDateStart

`func (o *Activity) GetDateStart() time.Time`

GetDateStart returns the DateStart field if non-nil, zero value otherwise.

### GetDateStartOk

`func (o *Activity) GetDateStartOk() (*time.Time, bool)`

GetDateStartOk returns a tuple with the DateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStart

`func (o *Activity) SetDateStart(v time.Time)`

SetDateStart sets DateStart field to given value.

### HasDateStart

`func (o *Activity) HasDateStart() bool`

HasDateStart returns a boolean if a field has been set.

### GetDateEnd

`func (o *Activity) GetDateEnd() time.Time`

GetDateEnd returns the DateEnd field if non-nil, zero value otherwise.

### GetDateEndOk

`func (o *Activity) GetDateEndOk() (*time.Time, bool)`

GetDateEndOk returns a tuple with the DateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnd

`func (o *Activity) SetDateEnd(v time.Time)`

SetDateEnd sets DateEnd field to given value.

### HasDateEnd

`func (o *Activity) HasDateEnd() bool`

HasDateEnd returns a boolean if a field has been set.

### GetAssignedBy

`func (o *Activity) GetAssignedBy() MemberReference`

GetAssignedBy returns the AssignedBy field if non-nil, zero value otherwise.

### GetAssignedByOk

`func (o *Activity) GetAssignedByOk() (*MemberReference, bool)`

GetAssignedByOk returns a tuple with the AssignedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedBy

`func (o *Activity) SetAssignedBy(v MemberReference)`

SetAssignedBy sets AssignedBy field to given value.

### HasAssignedBy

`func (o *Activity) HasAssignedBy() bool`

HasAssignedBy returns a boolean if a field has been set.

### GetAssignTo

`func (o *Activity) GetAssignTo() MemberReference`

GetAssignTo returns the AssignTo field if non-nil, zero value otherwise.

### GetAssignToOk

`func (o *Activity) GetAssignToOk() (*MemberReference, bool)`

GetAssignToOk returns a tuple with the AssignTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignTo

`func (o *Activity) SetAssignTo(v MemberReference)`

SetAssignTo sets AssignTo field to given value.


### GetScheduleStatus

`func (o *Activity) GetScheduleStatus() ScheduleStatusReference`

GetScheduleStatus returns the ScheduleStatus field if non-nil, zero value otherwise.

### GetScheduleStatusOk

`func (o *Activity) GetScheduleStatusOk() (*ScheduleStatusReference, bool)`

GetScheduleStatusOk returns a tuple with the ScheduleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleStatus

`func (o *Activity) SetScheduleStatus(v ScheduleStatusReference)`

SetScheduleStatus sets ScheduleStatus field to given value.

### HasScheduleStatus

`func (o *Activity) HasScheduleStatus() bool`

HasScheduleStatus returns a boolean if a field has been set.

### GetReminder

`func (o *Activity) GetReminder() ReminderReference`

GetReminder returns the Reminder field if non-nil, zero value otherwise.

### GetReminderOk

`func (o *Activity) GetReminderOk() (*ReminderReference, bool)`

GetReminderOk returns a tuple with the Reminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminder

`func (o *Activity) SetReminder(v ReminderReference)`

SetReminder sets Reminder field to given value.

### HasReminder

`func (o *Activity) HasReminder() bool`

HasReminder returns a boolean if a field has been set.

### GetWhere

`func (o *Activity) GetWhere() ServiceLocationReference`

GetWhere returns the Where field if non-nil, zero value otherwise.

### GetWhereOk

`func (o *Activity) GetWhereOk() (*ServiceLocationReference, bool)`

GetWhereOk returns a tuple with the Where field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhere

`func (o *Activity) SetWhere(v ServiceLocationReference)`

SetWhere sets Where field to given value.

### HasWhere

`func (o *Activity) HasWhere() bool`

HasWhere returns a boolean if a field has been set.

### GetNotifyFlag

`func (o *Activity) GetNotifyFlag() bool`

GetNotifyFlag returns the NotifyFlag field if non-nil, zero value otherwise.

### GetNotifyFlagOk

`func (o *Activity) GetNotifyFlagOk() (*bool, bool)`

GetNotifyFlagOk returns a tuple with the NotifyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyFlag

`func (o *Activity) SetNotifyFlag(v bool)`

SetNotifyFlag sets NotifyFlag field to given value.

### HasNotifyFlag

`func (o *Activity) HasNotifyFlag() bool`

HasNotifyFlag returns a boolean if a field has been set.

### SetNotifyFlagNil

`func (o *Activity) SetNotifyFlagNil(b bool)`

 SetNotifyFlagNil sets the value for NotifyFlag to be an explicit nil

### UnsetNotifyFlag
`func (o *Activity) UnsetNotifyFlag()`

UnsetNotifyFlag ensures that no value is present for NotifyFlag, not even an explicit nil
### GetMobileGuid

`func (o *Activity) GetMobileGuid() string`

GetMobileGuid returns the MobileGuid field if non-nil, zero value otherwise.

### GetMobileGuidOk

`func (o *Activity) GetMobileGuidOk() (*string, bool)`

GetMobileGuidOk returns a tuple with the MobileGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileGuid

`func (o *Activity) SetMobileGuid(v string)`

SetMobileGuid sets MobileGuid field to given value.

### HasMobileGuid

`func (o *Activity) HasMobileGuid() bool`

HasMobileGuid returns a boolean if a field has been set.

### SetMobileGuidNil

`func (o *Activity) SetMobileGuidNil(b bool)`

 SetMobileGuidNil sets the value for MobileGuid to be an explicit nil

### UnsetMobileGuid
`func (o *Activity) UnsetMobileGuid()`

UnsetMobileGuid ensures that no value is present for MobileGuid, not even an explicit nil
### GetCurrency

`func (o *Activity) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Activity) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Activity) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Activity) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInfo

`func (o *Activity) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Activity) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Activity) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Activity) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCustomFields

`func (o *Activity) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Activity) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Activity) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Activity) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


