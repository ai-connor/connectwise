# ServiceNote

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
**DateCreated** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**InternalFlag** | Pointer to **NullableBool** |  | [optional] 
**ExternalFlag** | Pointer to **NullableBool** |  | [optional] 
**SentimentScore** | Pointer to **NullableFloat64** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewServiceNote

`func NewServiceNote() *ServiceNote`

NewServiceNote instantiates a new ServiceNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceNoteWithDefaults

`func NewServiceNoteWithDefaults() *ServiceNote`

NewServiceNoteWithDefaults instantiates a new ServiceNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceNote) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceNote) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceNote) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceNote) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTicketId

`func (o *ServiceNote) GetTicketId() int32`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *ServiceNote) GetTicketIdOk() (*int32, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *ServiceNote) SetTicketId(v int32)`

SetTicketId sets TicketId field to given value.

### HasTicketId

`func (o *ServiceNote) HasTicketId() bool`

HasTicketId returns a boolean if a field has been set.

### SetTicketIdNil

`func (o *ServiceNote) SetTicketIdNil(b bool)`

 SetTicketIdNil sets the value for TicketId to be an explicit nil

### UnsetTicketId
`func (o *ServiceNote) UnsetTicketId()`

UnsetTicketId ensures that no value is present for TicketId, not even an explicit nil
### GetText

`func (o *ServiceNote) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ServiceNote) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ServiceNote) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ServiceNote) HasText() bool`

HasText returns a boolean if a field has been set.

### GetDetailDescriptionFlag

`func (o *ServiceNote) GetDetailDescriptionFlag() bool`

GetDetailDescriptionFlag returns the DetailDescriptionFlag field if non-nil, zero value otherwise.

### GetDetailDescriptionFlagOk

`func (o *ServiceNote) GetDetailDescriptionFlagOk() (*bool, bool)`

GetDetailDescriptionFlagOk returns a tuple with the DetailDescriptionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailDescriptionFlag

`func (o *ServiceNote) SetDetailDescriptionFlag(v bool)`

SetDetailDescriptionFlag sets DetailDescriptionFlag field to given value.

### HasDetailDescriptionFlag

`func (o *ServiceNote) HasDetailDescriptionFlag() bool`

HasDetailDescriptionFlag returns a boolean if a field has been set.

### SetDetailDescriptionFlagNil

`func (o *ServiceNote) SetDetailDescriptionFlagNil(b bool)`

 SetDetailDescriptionFlagNil sets the value for DetailDescriptionFlag to be an explicit nil

### UnsetDetailDescriptionFlag
`func (o *ServiceNote) UnsetDetailDescriptionFlag()`

UnsetDetailDescriptionFlag ensures that no value is present for DetailDescriptionFlag, not even an explicit nil
### GetInternalAnalysisFlag

`func (o *ServiceNote) GetInternalAnalysisFlag() bool`

GetInternalAnalysisFlag returns the InternalAnalysisFlag field if non-nil, zero value otherwise.

### GetInternalAnalysisFlagOk

`func (o *ServiceNote) GetInternalAnalysisFlagOk() (*bool, bool)`

GetInternalAnalysisFlagOk returns a tuple with the InternalAnalysisFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAnalysisFlag

`func (o *ServiceNote) SetInternalAnalysisFlag(v bool)`

SetInternalAnalysisFlag sets InternalAnalysisFlag field to given value.

### HasInternalAnalysisFlag

`func (o *ServiceNote) HasInternalAnalysisFlag() bool`

HasInternalAnalysisFlag returns a boolean if a field has been set.

### SetInternalAnalysisFlagNil

`func (o *ServiceNote) SetInternalAnalysisFlagNil(b bool)`

 SetInternalAnalysisFlagNil sets the value for InternalAnalysisFlag to be an explicit nil

### UnsetInternalAnalysisFlag
`func (o *ServiceNote) UnsetInternalAnalysisFlag()`

UnsetInternalAnalysisFlag ensures that no value is present for InternalAnalysisFlag, not even an explicit nil
### GetResolutionFlag

`func (o *ServiceNote) GetResolutionFlag() bool`

GetResolutionFlag returns the ResolutionFlag field if non-nil, zero value otherwise.

### GetResolutionFlagOk

`func (o *ServiceNote) GetResolutionFlagOk() (*bool, bool)`

GetResolutionFlagOk returns a tuple with the ResolutionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionFlag

`func (o *ServiceNote) SetResolutionFlag(v bool)`

SetResolutionFlag sets ResolutionFlag field to given value.

### HasResolutionFlag

`func (o *ServiceNote) HasResolutionFlag() bool`

HasResolutionFlag returns a boolean if a field has been set.

### SetResolutionFlagNil

`func (o *ServiceNote) SetResolutionFlagNil(b bool)`

 SetResolutionFlagNil sets the value for ResolutionFlag to be an explicit nil

### UnsetResolutionFlag
`func (o *ServiceNote) UnsetResolutionFlag()`

UnsetResolutionFlag ensures that no value is present for ResolutionFlag, not even an explicit nil
### GetIssueFlag

`func (o *ServiceNote) GetIssueFlag() bool`

GetIssueFlag returns the IssueFlag field if non-nil, zero value otherwise.

### GetIssueFlagOk

`func (o *ServiceNote) GetIssueFlagOk() (*bool, bool)`

GetIssueFlagOk returns a tuple with the IssueFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueFlag

`func (o *ServiceNote) SetIssueFlag(v bool)`

SetIssueFlag sets IssueFlag field to given value.

### HasIssueFlag

`func (o *ServiceNote) HasIssueFlag() bool`

HasIssueFlag returns a boolean if a field has been set.

### SetIssueFlagNil

`func (o *ServiceNote) SetIssueFlagNil(b bool)`

 SetIssueFlagNil sets the value for IssueFlag to be an explicit nil

### UnsetIssueFlag
`func (o *ServiceNote) UnsetIssueFlag()`

UnsetIssueFlag ensures that no value is present for IssueFlag, not even an explicit nil
### GetMember

`func (o *ServiceNote) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ServiceNote) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ServiceNote) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *ServiceNote) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetContact

`func (o *ServiceNote) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ServiceNote) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ServiceNote) SetContact(v ContactReference)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ServiceNote) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetCustomerUpdatedFlag

`func (o *ServiceNote) GetCustomerUpdatedFlag() bool`

GetCustomerUpdatedFlag returns the CustomerUpdatedFlag field if non-nil, zero value otherwise.

### GetCustomerUpdatedFlagOk

`func (o *ServiceNote) GetCustomerUpdatedFlagOk() (*bool, bool)`

GetCustomerUpdatedFlagOk returns a tuple with the CustomerUpdatedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUpdatedFlag

`func (o *ServiceNote) SetCustomerUpdatedFlag(v bool)`

SetCustomerUpdatedFlag sets CustomerUpdatedFlag field to given value.

### HasCustomerUpdatedFlag

`func (o *ServiceNote) HasCustomerUpdatedFlag() bool`

HasCustomerUpdatedFlag returns a boolean if a field has been set.

### SetCustomerUpdatedFlagNil

`func (o *ServiceNote) SetCustomerUpdatedFlagNil(b bool)`

 SetCustomerUpdatedFlagNil sets the value for CustomerUpdatedFlag to be an explicit nil

### UnsetCustomerUpdatedFlag
`func (o *ServiceNote) UnsetCustomerUpdatedFlag()`

UnsetCustomerUpdatedFlag ensures that no value is present for CustomerUpdatedFlag, not even an explicit nil
### GetProcessNotifications

`func (o *ServiceNote) GetProcessNotifications() bool`

GetProcessNotifications returns the ProcessNotifications field if non-nil, zero value otherwise.

### GetProcessNotificationsOk

`func (o *ServiceNote) GetProcessNotificationsOk() (*bool, bool)`

GetProcessNotificationsOk returns a tuple with the ProcessNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessNotifications

`func (o *ServiceNote) SetProcessNotifications(v bool)`

SetProcessNotifications sets ProcessNotifications field to given value.

### HasProcessNotifications

`func (o *ServiceNote) HasProcessNotifications() bool`

HasProcessNotifications returns a boolean if a field has been set.

### SetProcessNotificationsNil

`func (o *ServiceNote) SetProcessNotificationsNil(b bool)`

 SetProcessNotificationsNil sets the value for ProcessNotifications to be an explicit nil

### UnsetProcessNotifications
`func (o *ServiceNote) UnsetProcessNotifications()`

UnsetProcessNotifications ensures that no value is present for ProcessNotifications, not even an explicit nil
### GetDateCreated

`func (o *ServiceNote) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ServiceNote) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ServiceNote) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ServiceNote) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ServiceNote) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ServiceNote) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ServiceNote) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ServiceNote) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetInternalFlag

`func (o *ServiceNote) GetInternalFlag() bool`

GetInternalFlag returns the InternalFlag field if non-nil, zero value otherwise.

### GetInternalFlagOk

`func (o *ServiceNote) GetInternalFlagOk() (*bool, bool)`

GetInternalFlagOk returns a tuple with the InternalFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalFlag

`func (o *ServiceNote) SetInternalFlag(v bool)`

SetInternalFlag sets InternalFlag field to given value.

### HasInternalFlag

`func (o *ServiceNote) HasInternalFlag() bool`

HasInternalFlag returns a boolean if a field has been set.

### SetInternalFlagNil

`func (o *ServiceNote) SetInternalFlagNil(b bool)`

 SetInternalFlagNil sets the value for InternalFlag to be an explicit nil

### UnsetInternalFlag
`func (o *ServiceNote) UnsetInternalFlag()`

UnsetInternalFlag ensures that no value is present for InternalFlag, not even an explicit nil
### GetExternalFlag

`func (o *ServiceNote) GetExternalFlag() bool`

GetExternalFlag returns the ExternalFlag field if non-nil, zero value otherwise.

### GetExternalFlagOk

`func (o *ServiceNote) GetExternalFlagOk() (*bool, bool)`

GetExternalFlagOk returns a tuple with the ExternalFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFlag

`func (o *ServiceNote) SetExternalFlag(v bool)`

SetExternalFlag sets ExternalFlag field to given value.

### HasExternalFlag

`func (o *ServiceNote) HasExternalFlag() bool`

HasExternalFlag returns a boolean if a field has been set.

### SetExternalFlagNil

`func (o *ServiceNote) SetExternalFlagNil(b bool)`

 SetExternalFlagNil sets the value for ExternalFlag to be an explicit nil

### UnsetExternalFlag
`func (o *ServiceNote) UnsetExternalFlag()`

UnsetExternalFlag ensures that no value is present for ExternalFlag, not even an explicit nil
### GetSentimentScore

`func (o *ServiceNote) GetSentimentScore() float64`

GetSentimentScore returns the SentimentScore field if non-nil, zero value otherwise.

### GetSentimentScoreOk

`func (o *ServiceNote) GetSentimentScoreOk() (*float64, bool)`

GetSentimentScoreOk returns a tuple with the SentimentScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentimentScore

`func (o *ServiceNote) SetSentimentScore(v float64)`

SetSentimentScore sets SentimentScore field to given value.

### HasSentimentScore

`func (o *ServiceNote) HasSentimentScore() bool`

HasSentimentScore returns a boolean if a field has been set.

### SetSentimentScoreNil

`func (o *ServiceNote) SetSentimentScoreNil(b bool)`

 SetSentimentScoreNil sets the value for SentimentScore to be an explicit nil

### UnsetSentimentScore
`func (o *ServiceNote) UnsetSentimentScore()`

UnsetSentimentScore ensures that no value is present for SentimentScore, not even an explicit nil
### GetInfo

`func (o *ServiceNote) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ServiceNote) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ServiceNote) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ServiceNote) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


