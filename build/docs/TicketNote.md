# TicketNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**TicketId** | Pointer to **NullableInt32** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**DetailDescriptionFlag** | Pointer to **NullableBool** |  | [optional] 
**InternalAnalysisFlag** | Pointer to **NullableBool** |  | [optional] 
**ResolutionFlag** | Pointer to **NullableBool** |  | [optional] 
**IssueFlag** | Pointer to **NullableBool** |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Contact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**CustomerUpdatedFlag** | Pointer to **NullableBool** |  | [optional] 
**ProcessNotifications** | Pointer to **NullableBool** |  | [optional] 
**InternalFlag** | Pointer to **NullableBool** |  | [optional] 
**ExternalFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTicketNote

`func NewTicketNote() *TicketNote`

NewTicketNote instantiates a new TicketNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketNoteWithDefaults

`func NewTicketNoteWithDefaults() *TicketNote`

NewTicketNoteWithDefaults instantiates a new TicketNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TicketNote) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TicketNote) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TicketNote) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TicketNote) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTicketId

`func (o *TicketNote) GetTicketId() int32`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *TicketNote) GetTicketIdOk() (*int32, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *TicketNote) SetTicketId(v int32)`

SetTicketId sets TicketId field to given value.

### HasTicketId

`func (o *TicketNote) HasTicketId() bool`

HasTicketId returns a boolean if a field has been set.

### SetTicketIdNil

`func (o *TicketNote) SetTicketIdNil(b bool)`

 SetTicketIdNil sets the value for TicketId to be an explicit nil

### UnsetTicketId
`func (o *TicketNote) UnsetTicketId()`

UnsetTicketId ensures that no value is present for TicketId, not even an explicit nil
### GetText

`func (o *TicketNote) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TicketNote) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TicketNote) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TicketNote) HasText() bool`

HasText returns a boolean if a field has been set.

### GetDetailDescriptionFlag

`func (o *TicketNote) GetDetailDescriptionFlag() bool`

GetDetailDescriptionFlag returns the DetailDescriptionFlag field if non-nil, zero value otherwise.

### GetDetailDescriptionFlagOk

`func (o *TicketNote) GetDetailDescriptionFlagOk() (*bool, bool)`

GetDetailDescriptionFlagOk returns a tuple with the DetailDescriptionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailDescriptionFlag

`func (o *TicketNote) SetDetailDescriptionFlag(v bool)`

SetDetailDescriptionFlag sets DetailDescriptionFlag field to given value.

### HasDetailDescriptionFlag

`func (o *TicketNote) HasDetailDescriptionFlag() bool`

HasDetailDescriptionFlag returns a boolean if a field has been set.

### SetDetailDescriptionFlagNil

`func (o *TicketNote) SetDetailDescriptionFlagNil(b bool)`

 SetDetailDescriptionFlagNil sets the value for DetailDescriptionFlag to be an explicit nil

### UnsetDetailDescriptionFlag
`func (o *TicketNote) UnsetDetailDescriptionFlag()`

UnsetDetailDescriptionFlag ensures that no value is present for DetailDescriptionFlag, not even an explicit nil
### GetInternalAnalysisFlag

`func (o *TicketNote) GetInternalAnalysisFlag() bool`

GetInternalAnalysisFlag returns the InternalAnalysisFlag field if non-nil, zero value otherwise.

### GetInternalAnalysisFlagOk

`func (o *TicketNote) GetInternalAnalysisFlagOk() (*bool, bool)`

GetInternalAnalysisFlagOk returns a tuple with the InternalAnalysisFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAnalysisFlag

`func (o *TicketNote) SetInternalAnalysisFlag(v bool)`

SetInternalAnalysisFlag sets InternalAnalysisFlag field to given value.

### HasInternalAnalysisFlag

`func (o *TicketNote) HasInternalAnalysisFlag() bool`

HasInternalAnalysisFlag returns a boolean if a field has been set.

### SetInternalAnalysisFlagNil

`func (o *TicketNote) SetInternalAnalysisFlagNil(b bool)`

 SetInternalAnalysisFlagNil sets the value for InternalAnalysisFlag to be an explicit nil

### UnsetInternalAnalysisFlag
`func (o *TicketNote) UnsetInternalAnalysisFlag()`

UnsetInternalAnalysisFlag ensures that no value is present for InternalAnalysisFlag, not even an explicit nil
### GetResolutionFlag

`func (o *TicketNote) GetResolutionFlag() bool`

GetResolutionFlag returns the ResolutionFlag field if non-nil, zero value otherwise.

### GetResolutionFlagOk

`func (o *TicketNote) GetResolutionFlagOk() (*bool, bool)`

GetResolutionFlagOk returns a tuple with the ResolutionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionFlag

`func (o *TicketNote) SetResolutionFlag(v bool)`

SetResolutionFlag sets ResolutionFlag field to given value.

### HasResolutionFlag

`func (o *TicketNote) HasResolutionFlag() bool`

HasResolutionFlag returns a boolean if a field has been set.

### SetResolutionFlagNil

`func (o *TicketNote) SetResolutionFlagNil(b bool)`

 SetResolutionFlagNil sets the value for ResolutionFlag to be an explicit nil

### UnsetResolutionFlag
`func (o *TicketNote) UnsetResolutionFlag()`

UnsetResolutionFlag ensures that no value is present for ResolutionFlag, not even an explicit nil
### GetIssueFlag

`func (o *TicketNote) GetIssueFlag() bool`

GetIssueFlag returns the IssueFlag field if non-nil, zero value otherwise.

### GetIssueFlagOk

`func (o *TicketNote) GetIssueFlagOk() (*bool, bool)`

GetIssueFlagOk returns a tuple with the IssueFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueFlag

`func (o *TicketNote) SetIssueFlag(v bool)`

SetIssueFlag sets IssueFlag field to given value.

### HasIssueFlag

`func (o *TicketNote) HasIssueFlag() bool`

HasIssueFlag returns a boolean if a field has been set.

### SetIssueFlagNil

`func (o *TicketNote) SetIssueFlagNil(b bool)`

 SetIssueFlagNil sets the value for IssueFlag to be an explicit nil

### UnsetIssueFlag
`func (o *TicketNote) UnsetIssueFlag()`

UnsetIssueFlag ensures that no value is present for IssueFlag, not even an explicit nil
### GetMember

`func (o *TicketNote) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *TicketNote) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *TicketNote) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *TicketNote) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetContact

`func (o *TicketNote) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *TicketNote) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *TicketNote) SetContact(v ContactReference)`

SetContact sets Contact field to given value.

### HasContact

`func (o *TicketNote) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetCustomerUpdatedFlag

`func (o *TicketNote) GetCustomerUpdatedFlag() bool`

GetCustomerUpdatedFlag returns the CustomerUpdatedFlag field if non-nil, zero value otherwise.

### GetCustomerUpdatedFlagOk

`func (o *TicketNote) GetCustomerUpdatedFlagOk() (*bool, bool)`

GetCustomerUpdatedFlagOk returns a tuple with the CustomerUpdatedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUpdatedFlag

`func (o *TicketNote) SetCustomerUpdatedFlag(v bool)`

SetCustomerUpdatedFlag sets CustomerUpdatedFlag field to given value.

### HasCustomerUpdatedFlag

`func (o *TicketNote) HasCustomerUpdatedFlag() bool`

HasCustomerUpdatedFlag returns a boolean if a field has been set.

### SetCustomerUpdatedFlagNil

`func (o *TicketNote) SetCustomerUpdatedFlagNil(b bool)`

 SetCustomerUpdatedFlagNil sets the value for CustomerUpdatedFlag to be an explicit nil

### UnsetCustomerUpdatedFlag
`func (o *TicketNote) UnsetCustomerUpdatedFlag()`

UnsetCustomerUpdatedFlag ensures that no value is present for CustomerUpdatedFlag, not even an explicit nil
### GetProcessNotifications

`func (o *TicketNote) GetProcessNotifications() bool`

GetProcessNotifications returns the ProcessNotifications field if non-nil, zero value otherwise.

### GetProcessNotificationsOk

`func (o *TicketNote) GetProcessNotificationsOk() (*bool, bool)`

GetProcessNotificationsOk returns a tuple with the ProcessNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessNotifications

`func (o *TicketNote) SetProcessNotifications(v bool)`

SetProcessNotifications sets ProcessNotifications field to given value.

### HasProcessNotifications

`func (o *TicketNote) HasProcessNotifications() bool`

HasProcessNotifications returns a boolean if a field has been set.

### SetProcessNotificationsNil

`func (o *TicketNote) SetProcessNotificationsNil(b bool)`

 SetProcessNotificationsNil sets the value for ProcessNotifications to be an explicit nil

### UnsetProcessNotifications
`func (o *TicketNote) UnsetProcessNotifications()`

UnsetProcessNotifications ensures that no value is present for ProcessNotifications, not even an explicit nil
### GetInternalFlag

`func (o *TicketNote) GetInternalFlag() bool`

GetInternalFlag returns the InternalFlag field if non-nil, zero value otherwise.

### GetInternalFlagOk

`func (o *TicketNote) GetInternalFlagOk() (*bool, bool)`

GetInternalFlagOk returns a tuple with the InternalFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalFlag

`func (o *TicketNote) SetInternalFlag(v bool)`

SetInternalFlag sets InternalFlag field to given value.

### HasInternalFlag

`func (o *TicketNote) HasInternalFlag() bool`

HasInternalFlag returns a boolean if a field has been set.

### SetInternalFlagNil

`func (o *TicketNote) SetInternalFlagNil(b bool)`

 SetInternalFlagNil sets the value for InternalFlag to be an explicit nil

### UnsetInternalFlag
`func (o *TicketNote) UnsetInternalFlag()`

UnsetInternalFlag ensures that no value is present for InternalFlag, not even an explicit nil
### GetExternalFlag

`func (o *TicketNote) GetExternalFlag() bool`

GetExternalFlag returns the ExternalFlag field if non-nil, zero value otherwise.

### GetExternalFlagOk

`func (o *TicketNote) GetExternalFlagOk() (*bool, bool)`

GetExternalFlagOk returns a tuple with the ExternalFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFlag

`func (o *TicketNote) SetExternalFlag(v bool)`

SetExternalFlag sets ExternalFlag field to given value.

### HasExternalFlag

`func (o *TicketNote) HasExternalFlag() bool`

HasExternalFlag returns a boolean if a field has been set.

### SetExternalFlagNil

`func (o *TicketNote) SetExternalFlagNil(b bool)`

 SetExternalFlagNil sets the value for ExternalFlag to be an explicit nil

### UnsetExternalFlag
`func (o *TicketNote) UnsetExternalFlag()`

UnsetExternalFlag ensures that no value is present for ExternalFlag, not even an explicit nil
### GetInfo

`func (o *TicketNote) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TicketNote) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TicketNote) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TicketNote) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


