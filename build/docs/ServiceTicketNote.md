# ServiceTicketNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**NoteType** | Pointer to **NullableString** |  | [optional] 
**Ticket** | Pointer to [**TicketReference**](TicketReference.md) |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**IsMarkdownFlag** | Pointer to **NullableBool** |  | [optional] 
**DetailDescriptionFlag** | Pointer to **NullableBool** |  | [optional] 
**InternalAnalysisFlag** | Pointer to **NullableBool** |  | [optional] 
**ResolutionFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeStart** | Pointer to **string** |  | [optional] 
**TimeEnd** | Pointer to **string** |  | [optional] 
**BundledFlag** | Pointer to **NullableBool** |  | [optional] 
**MergedFlag** | Pointer to **NullableBool** |  | [optional] 
**IssueFlag** | Pointer to **NullableBool** |  | [optional] 
**OriginalAuthor** | Pointer to **string** |  | [optional] 
**CreatedByParentFlag** | Pointer to **NullableBool** |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Contact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewServiceTicketNote

`func NewServiceTicketNote() *ServiceTicketNote`

NewServiceTicketNote instantiates a new ServiceTicketNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTicketNoteWithDefaults

`func NewServiceTicketNoteWithDefaults() *ServiceTicketNote`

NewServiceTicketNoteWithDefaults instantiates a new ServiceTicketNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceTicketNote) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceTicketNote) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceTicketNote) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceTicketNote) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNoteType

`func (o *ServiceTicketNote) GetNoteType() string`

GetNoteType returns the NoteType field if non-nil, zero value otherwise.

### GetNoteTypeOk

`func (o *ServiceTicketNote) GetNoteTypeOk() (*string, bool)`

GetNoteTypeOk returns a tuple with the NoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteType

`func (o *ServiceTicketNote) SetNoteType(v string)`

SetNoteType sets NoteType field to given value.

### HasNoteType

`func (o *ServiceTicketNote) HasNoteType() bool`

HasNoteType returns a boolean if a field has been set.

### SetNoteTypeNil

`func (o *ServiceTicketNote) SetNoteTypeNil(b bool)`

 SetNoteTypeNil sets the value for NoteType to be an explicit nil

### UnsetNoteType
`func (o *ServiceTicketNote) UnsetNoteType()`

UnsetNoteType ensures that no value is present for NoteType, not even an explicit nil
### GetTicket

`func (o *ServiceTicketNote) GetTicket() TicketReference`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *ServiceTicketNote) GetTicketOk() (*TicketReference, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *ServiceTicketNote) SetTicket(v TicketReference)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *ServiceTicketNote) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetText

`func (o *ServiceTicketNote) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ServiceTicketNote) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ServiceTicketNote) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ServiceTicketNote) HasText() bool`

HasText returns a boolean if a field has been set.

### GetIsMarkdownFlag

`func (o *ServiceTicketNote) GetIsMarkdownFlag() bool`

GetIsMarkdownFlag returns the IsMarkdownFlag field if non-nil, zero value otherwise.

### GetIsMarkdownFlagOk

`func (o *ServiceTicketNote) GetIsMarkdownFlagOk() (*bool, bool)`

GetIsMarkdownFlagOk returns a tuple with the IsMarkdownFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarkdownFlag

`func (o *ServiceTicketNote) SetIsMarkdownFlag(v bool)`

SetIsMarkdownFlag sets IsMarkdownFlag field to given value.

### HasIsMarkdownFlag

`func (o *ServiceTicketNote) HasIsMarkdownFlag() bool`

HasIsMarkdownFlag returns a boolean if a field has been set.

### SetIsMarkdownFlagNil

`func (o *ServiceTicketNote) SetIsMarkdownFlagNil(b bool)`

 SetIsMarkdownFlagNil sets the value for IsMarkdownFlag to be an explicit nil

### UnsetIsMarkdownFlag
`func (o *ServiceTicketNote) UnsetIsMarkdownFlag()`

UnsetIsMarkdownFlag ensures that no value is present for IsMarkdownFlag, not even an explicit nil
### GetDetailDescriptionFlag

`func (o *ServiceTicketNote) GetDetailDescriptionFlag() bool`

GetDetailDescriptionFlag returns the DetailDescriptionFlag field if non-nil, zero value otherwise.

### GetDetailDescriptionFlagOk

`func (o *ServiceTicketNote) GetDetailDescriptionFlagOk() (*bool, bool)`

GetDetailDescriptionFlagOk returns a tuple with the DetailDescriptionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailDescriptionFlag

`func (o *ServiceTicketNote) SetDetailDescriptionFlag(v bool)`

SetDetailDescriptionFlag sets DetailDescriptionFlag field to given value.

### HasDetailDescriptionFlag

`func (o *ServiceTicketNote) HasDetailDescriptionFlag() bool`

HasDetailDescriptionFlag returns a boolean if a field has been set.

### SetDetailDescriptionFlagNil

`func (o *ServiceTicketNote) SetDetailDescriptionFlagNil(b bool)`

 SetDetailDescriptionFlagNil sets the value for DetailDescriptionFlag to be an explicit nil

### UnsetDetailDescriptionFlag
`func (o *ServiceTicketNote) UnsetDetailDescriptionFlag()`

UnsetDetailDescriptionFlag ensures that no value is present for DetailDescriptionFlag, not even an explicit nil
### GetInternalAnalysisFlag

`func (o *ServiceTicketNote) GetInternalAnalysisFlag() bool`

GetInternalAnalysisFlag returns the InternalAnalysisFlag field if non-nil, zero value otherwise.

### GetInternalAnalysisFlagOk

`func (o *ServiceTicketNote) GetInternalAnalysisFlagOk() (*bool, bool)`

GetInternalAnalysisFlagOk returns a tuple with the InternalAnalysisFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAnalysisFlag

`func (o *ServiceTicketNote) SetInternalAnalysisFlag(v bool)`

SetInternalAnalysisFlag sets InternalAnalysisFlag field to given value.

### HasInternalAnalysisFlag

`func (o *ServiceTicketNote) HasInternalAnalysisFlag() bool`

HasInternalAnalysisFlag returns a boolean if a field has been set.

### SetInternalAnalysisFlagNil

`func (o *ServiceTicketNote) SetInternalAnalysisFlagNil(b bool)`

 SetInternalAnalysisFlagNil sets the value for InternalAnalysisFlag to be an explicit nil

### UnsetInternalAnalysisFlag
`func (o *ServiceTicketNote) UnsetInternalAnalysisFlag()`

UnsetInternalAnalysisFlag ensures that no value is present for InternalAnalysisFlag, not even an explicit nil
### GetResolutionFlag

`func (o *ServiceTicketNote) GetResolutionFlag() bool`

GetResolutionFlag returns the ResolutionFlag field if non-nil, zero value otherwise.

### GetResolutionFlagOk

`func (o *ServiceTicketNote) GetResolutionFlagOk() (*bool, bool)`

GetResolutionFlagOk returns a tuple with the ResolutionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionFlag

`func (o *ServiceTicketNote) SetResolutionFlag(v bool)`

SetResolutionFlag sets ResolutionFlag field to given value.

### HasResolutionFlag

`func (o *ServiceTicketNote) HasResolutionFlag() bool`

HasResolutionFlag returns a boolean if a field has been set.

### SetResolutionFlagNil

`func (o *ServiceTicketNote) SetResolutionFlagNil(b bool)`

 SetResolutionFlagNil sets the value for ResolutionFlag to be an explicit nil

### UnsetResolutionFlag
`func (o *ServiceTicketNote) UnsetResolutionFlag()`

UnsetResolutionFlag ensures that no value is present for ResolutionFlag, not even an explicit nil
### GetTimeStart

`func (o *ServiceTicketNote) GetTimeStart() string`

GetTimeStart returns the TimeStart field if non-nil, zero value otherwise.

### GetTimeStartOk

`func (o *ServiceTicketNote) GetTimeStartOk() (*string, bool)`

GetTimeStartOk returns a tuple with the TimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStart

`func (o *ServiceTicketNote) SetTimeStart(v string)`

SetTimeStart sets TimeStart field to given value.

### HasTimeStart

`func (o *ServiceTicketNote) HasTimeStart() bool`

HasTimeStart returns a boolean if a field has been set.

### GetTimeEnd

`func (o *ServiceTicketNote) GetTimeEnd() string`

GetTimeEnd returns the TimeEnd field if non-nil, zero value otherwise.

### GetTimeEndOk

`func (o *ServiceTicketNote) GetTimeEndOk() (*string, bool)`

GetTimeEndOk returns a tuple with the TimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEnd

`func (o *ServiceTicketNote) SetTimeEnd(v string)`

SetTimeEnd sets TimeEnd field to given value.

### HasTimeEnd

`func (o *ServiceTicketNote) HasTimeEnd() bool`

HasTimeEnd returns a boolean if a field has been set.

### GetBundledFlag

`func (o *ServiceTicketNote) GetBundledFlag() bool`

GetBundledFlag returns the BundledFlag field if non-nil, zero value otherwise.

### GetBundledFlagOk

`func (o *ServiceTicketNote) GetBundledFlagOk() (*bool, bool)`

GetBundledFlagOk returns a tuple with the BundledFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundledFlag

`func (o *ServiceTicketNote) SetBundledFlag(v bool)`

SetBundledFlag sets BundledFlag field to given value.

### HasBundledFlag

`func (o *ServiceTicketNote) HasBundledFlag() bool`

HasBundledFlag returns a boolean if a field has been set.

### SetBundledFlagNil

`func (o *ServiceTicketNote) SetBundledFlagNil(b bool)`

 SetBundledFlagNil sets the value for BundledFlag to be an explicit nil

### UnsetBundledFlag
`func (o *ServiceTicketNote) UnsetBundledFlag()`

UnsetBundledFlag ensures that no value is present for BundledFlag, not even an explicit nil
### GetMergedFlag

`func (o *ServiceTicketNote) GetMergedFlag() bool`

GetMergedFlag returns the MergedFlag field if non-nil, zero value otherwise.

### GetMergedFlagOk

`func (o *ServiceTicketNote) GetMergedFlagOk() (*bool, bool)`

GetMergedFlagOk returns a tuple with the MergedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedFlag

`func (o *ServiceTicketNote) SetMergedFlag(v bool)`

SetMergedFlag sets MergedFlag field to given value.

### HasMergedFlag

`func (o *ServiceTicketNote) HasMergedFlag() bool`

HasMergedFlag returns a boolean if a field has been set.

### SetMergedFlagNil

`func (o *ServiceTicketNote) SetMergedFlagNil(b bool)`

 SetMergedFlagNil sets the value for MergedFlag to be an explicit nil

### UnsetMergedFlag
`func (o *ServiceTicketNote) UnsetMergedFlag()`

UnsetMergedFlag ensures that no value is present for MergedFlag, not even an explicit nil
### GetIssueFlag

`func (o *ServiceTicketNote) GetIssueFlag() bool`

GetIssueFlag returns the IssueFlag field if non-nil, zero value otherwise.

### GetIssueFlagOk

`func (o *ServiceTicketNote) GetIssueFlagOk() (*bool, bool)`

GetIssueFlagOk returns a tuple with the IssueFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueFlag

`func (o *ServiceTicketNote) SetIssueFlag(v bool)`

SetIssueFlag sets IssueFlag field to given value.

### HasIssueFlag

`func (o *ServiceTicketNote) HasIssueFlag() bool`

HasIssueFlag returns a boolean if a field has been set.

### SetIssueFlagNil

`func (o *ServiceTicketNote) SetIssueFlagNil(b bool)`

 SetIssueFlagNil sets the value for IssueFlag to be an explicit nil

### UnsetIssueFlag
`func (o *ServiceTicketNote) UnsetIssueFlag()`

UnsetIssueFlag ensures that no value is present for IssueFlag, not even an explicit nil
### GetOriginalAuthor

`func (o *ServiceTicketNote) GetOriginalAuthor() string`

GetOriginalAuthor returns the OriginalAuthor field if non-nil, zero value otherwise.

### GetOriginalAuthorOk

`func (o *ServiceTicketNote) GetOriginalAuthorOk() (*string, bool)`

GetOriginalAuthorOk returns a tuple with the OriginalAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAuthor

`func (o *ServiceTicketNote) SetOriginalAuthor(v string)`

SetOriginalAuthor sets OriginalAuthor field to given value.

### HasOriginalAuthor

`func (o *ServiceTicketNote) HasOriginalAuthor() bool`

HasOriginalAuthor returns a boolean if a field has been set.

### GetCreatedByParentFlag

`func (o *ServiceTicketNote) GetCreatedByParentFlag() bool`

GetCreatedByParentFlag returns the CreatedByParentFlag field if non-nil, zero value otherwise.

### GetCreatedByParentFlagOk

`func (o *ServiceTicketNote) GetCreatedByParentFlagOk() (*bool, bool)`

GetCreatedByParentFlagOk returns a tuple with the CreatedByParentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByParentFlag

`func (o *ServiceTicketNote) SetCreatedByParentFlag(v bool)`

SetCreatedByParentFlag sets CreatedByParentFlag field to given value.

### HasCreatedByParentFlag

`func (o *ServiceTicketNote) HasCreatedByParentFlag() bool`

HasCreatedByParentFlag returns a boolean if a field has been set.

### SetCreatedByParentFlagNil

`func (o *ServiceTicketNote) SetCreatedByParentFlagNil(b bool)`

 SetCreatedByParentFlagNil sets the value for CreatedByParentFlag to be an explicit nil

### UnsetCreatedByParentFlag
`func (o *ServiceTicketNote) UnsetCreatedByParentFlag()`

UnsetCreatedByParentFlag ensures that no value is present for CreatedByParentFlag, not even an explicit nil
### GetMember

`func (o *ServiceTicketNote) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ServiceTicketNote) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ServiceTicketNote) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *ServiceTicketNote) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetContact

`func (o *ServiceTicketNote) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ServiceTicketNote) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ServiceTicketNote) SetContact(v ContactReference)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ServiceTicketNote) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetInfo

`func (o *ServiceTicketNote) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ServiceTicketNote) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ServiceTicketNote) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ServiceTicketNote) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


