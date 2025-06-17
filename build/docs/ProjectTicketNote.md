# ProjectTicketNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**NoteType** | Pointer to **NullableString** |  | [optional] 
**Ticket** | Pointer to [**TicketReference**](TicketReference.md) |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**DetailDescriptionFlag** | Pointer to **NullableBool** |  | [optional] 
**InternalAnalysisFlag** | Pointer to **NullableBool** |  | [optional] 
**ResolutionFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeStart** | Pointer to **time.Time** |  | [optional] 
**TimeEnd** | Pointer to **time.Time** |  | [optional] 
**BundledFlag** | Pointer to **NullableBool** |  | [optional] 
**MergedFlag** | Pointer to **NullableBool** |  | [optional] 
**IssueFlag** | Pointer to **NullableBool** |  | [optional] 
**OriginalAuthor** | Pointer to **string** |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Contact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewProjectTicketNote

`func NewProjectTicketNote() *ProjectTicketNote`

NewProjectTicketNote instantiates a new ProjectTicketNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectTicketNoteWithDefaults

`func NewProjectTicketNoteWithDefaults() *ProjectTicketNote`

NewProjectTicketNoteWithDefaults instantiates a new ProjectTicketNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectTicketNote) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectTicketNote) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectTicketNote) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectTicketNote) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNoteType

`func (o *ProjectTicketNote) GetNoteType() string`

GetNoteType returns the NoteType field if non-nil, zero value otherwise.

### GetNoteTypeOk

`func (o *ProjectTicketNote) GetNoteTypeOk() (*string, bool)`

GetNoteTypeOk returns a tuple with the NoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteType

`func (o *ProjectTicketNote) SetNoteType(v string)`

SetNoteType sets NoteType field to given value.

### HasNoteType

`func (o *ProjectTicketNote) HasNoteType() bool`

HasNoteType returns a boolean if a field has been set.

### SetNoteTypeNil

`func (o *ProjectTicketNote) SetNoteTypeNil(b bool)`

 SetNoteTypeNil sets the value for NoteType to be an explicit nil

### UnsetNoteType
`func (o *ProjectTicketNote) UnsetNoteType()`

UnsetNoteType ensures that no value is present for NoteType, not even an explicit nil
### GetTicket

`func (o *ProjectTicketNote) GetTicket() TicketReference`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *ProjectTicketNote) GetTicketOk() (*TicketReference, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *ProjectTicketNote) SetTicket(v TicketReference)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *ProjectTicketNote) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetText

`func (o *ProjectTicketNote) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ProjectTicketNote) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ProjectTicketNote) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ProjectTicketNote) HasText() bool`

HasText returns a boolean if a field has been set.

### GetDetailDescriptionFlag

`func (o *ProjectTicketNote) GetDetailDescriptionFlag() bool`

GetDetailDescriptionFlag returns the DetailDescriptionFlag field if non-nil, zero value otherwise.

### GetDetailDescriptionFlagOk

`func (o *ProjectTicketNote) GetDetailDescriptionFlagOk() (*bool, bool)`

GetDetailDescriptionFlagOk returns a tuple with the DetailDescriptionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailDescriptionFlag

`func (o *ProjectTicketNote) SetDetailDescriptionFlag(v bool)`

SetDetailDescriptionFlag sets DetailDescriptionFlag field to given value.

### HasDetailDescriptionFlag

`func (o *ProjectTicketNote) HasDetailDescriptionFlag() bool`

HasDetailDescriptionFlag returns a boolean if a field has been set.

### SetDetailDescriptionFlagNil

`func (o *ProjectTicketNote) SetDetailDescriptionFlagNil(b bool)`

 SetDetailDescriptionFlagNil sets the value for DetailDescriptionFlag to be an explicit nil

### UnsetDetailDescriptionFlag
`func (o *ProjectTicketNote) UnsetDetailDescriptionFlag()`

UnsetDetailDescriptionFlag ensures that no value is present for DetailDescriptionFlag, not even an explicit nil
### GetInternalAnalysisFlag

`func (o *ProjectTicketNote) GetInternalAnalysisFlag() bool`

GetInternalAnalysisFlag returns the InternalAnalysisFlag field if non-nil, zero value otherwise.

### GetInternalAnalysisFlagOk

`func (o *ProjectTicketNote) GetInternalAnalysisFlagOk() (*bool, bool)`

GetInternalAnalysisFlagOk returns a tuple with the InternalAnalysisFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAnalysisFlag

`func (o *ProjectTicketNote) SetInternalAnalysisFlag(v bool)`

SetInternalAnalysisFlag sets InternalAnalysisFlag field to given value.

### HasInternalAnalysisFlag

`func (o *ProjectTicketNote) HasInternalAnalysisFlag() bool`

HasInternalAnalysisFlag returns a boolean if a field has been set.

### SetInternalAnalysisFlagNil

`func (o *ProjectTicketNote) SetInternalAnalysisFlagNil(b bool)`

 SetInternalAnalysisFlagNil sets the value for InternalAnalysisFlag to be an explicit nil

### UnsetInternalAnalysisFlag
`func (o *ProjectTicketNote) UnsetInternalAnalysisFlag()`

UnsetInternalAnalysisFlag ensures that no value is present for InternalAnalysisFlag, not even an explicit nil
### GetResolutionFlag

`func (o *ProjectTicketNote) GetResolutionFlag() bool`

GetResolutionFlag returns the ResolutionFlag field if non-nil, zero value otherwise.

### GetResolutionFlagOk

`func (o *ProjectTicketNote) GetResolutionFlagOk() (*bool, bool)`

GetResolutionFlagOk returns a tuple with the ResolutionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionFlag

`func (o *ProjectTicketNote) SetResolutionFlag(v bool)`

SetResolutionFlag sets ResolutionFlag field to given value.

### HasResolutionFlag

`func (o *ProjectTicketNote) HasResolutionFlag() bool`

HasResolutionFlag returns a boolean if a field has been set.

### SetResolutionFlagNil

`func (o *ProjectTicketNote) SetResolutionFlagNil(b bool)`

 SetResolutionFlagNil sets the value for ResolutionFlag to be an explicit nil

### UnsetResolutionFlag
`func (o *ProjectTicketNote) UnsetResolutionFlag()`

UnsetResolutionFlag ensures that no value is present for ResolutionFlag, not even an explicit nil
### GetTimeStart

`func (o *ProjectTicketNote) GetTimeStart() time.Time`

GetTimeStart returns the TimeStart field if non-nil, zero value otherwise.

### GetTimeStartOk

`func (o *ProjectTicketNote) GetTimeStartOk() (*time.Time, bool)`

GetTimeStartOk returns a tuple with the TimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStart

`func (o *ProjectTicketNote) SetTimeStart(v time.Time)`

SetTimeStart sets TimeStart field to given value.

### HasTimeStart

`func (o *ProjectTicketNote) HasTimeStart() bool`

HasTimeStart returns a boolean if a field has been set.

### GetTimeEnd

`func (o *ProjectTicketNote) GetTimeEnd() time.Time`

GetTimeEnd returns the TimeEnd field if non-nil, zero value otherwise.

### GetTimeEndOk

`func (o *ProjectTicketNote) GetTimeEndOk() (*time.Time, bool)`

GetTimeEndOk returns a tuple with the TimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEnd

`func (o *ProjectTicketNote) SetTimeEnd(v time.Time)`

SetTimeEnd sets TimeEnd field to given value.

### HasTimeEnd

`func (o *ProjectTicketNote) HasTimeEnd() bool`

HasTimeEnd returns a boolean if a field has been set.

### GetBundledFlag

`func (o *ProjectTicketNote) GetBundledFlag() bool`

GetBundledFlag returns the BundledFlag field if non-nil, zero value otherwise.

### GetBundledFlagOk

`func (o *ProjectTicketNote) GetBundledFlagOk() (*bool, bool)`

GetBundledFlagOk returns a tuple with the BundledFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundledFlag

`func (o *ProjectTicketNote) SetBundledFlag(v bool)`

SetBundledFlag sets BundledFlag field to given value.

### HasBundledFlag

`func (o *ProjectTicketNote) HasBundledFlag() bool`

HasBundledFlag returns a boolean if a field has been set.

### SetBundledFlagNil

`func (o *ProjectTicketNote) SetBundledFlagNil(b bool)`

 SetBundledFlagNil sets the value for BundledFlag to be an explicit nil

### UnsetBundledFlag
`func (o *ProjectTicketNote) UnsetBundledFlag()`

UnsetBundledFlag ensures that no value is present for BundledFlag, not even an explicit nil
### GetMergedFlag

`func (o *ProjectTicketNote) GetMergedFlag() bool`

GetMergedFlag returns the MergedFlag field if non-nil, zero value otherwise.

### GetMergedFlagOk

`func (o *ProjectTicketNote) GetMergedFlagOk() (*bool, bool)`

GetMergedFlagOk returns a tuple with the MergedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedFlag

`func (o *ProjectTicketNote) SetMergedFlag(v bool)`

SetMergedFlag sets MergedFlag field to given value.

### HasMergedFlag

`func (o *ProjectTicketNote) HasMergedFlag() bool`

HasMergedFlag returns a boolean if a field has been set.

### SetMergedFlagNil

`func (o *ProjectTicketNote) SetMergedFlagNil(b bool)`

 SetMergedFlagNil sets the value for MergedFlag to be an explicit nil

### UnsetMergedFlag
`func (o *ProjectTicketNote) UnsetMergedFlag()`

UnsetMergedFlag ensures that no value is present for MergedFlag, not even an explicit nil
### GetIssueFlag

`func (o *ProjectTicketNote) GetIssueFlag() bool`

GetIssueFlag returns the IssueFlag field if non-nil, zero value otherwise.

### GetIssueFlagOk

`func (o *ProjectTicketNote) GetIssueFlagOk() (*bool, bool)`

GetIssueFlagOk returns a tuple with the IssueFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueFlag

`func (o *ProjectTicketNote) SetIssueFlag(v bool)`

SetIssueFlag sets IssueFlag field to given value.

### HasIssueFlag

`func (o *ProjectTicketNote) HasIssueFlag() bool`

HasIssueFlag returns a boolean if a field has been set.

### SetIssueFlagNil

`func (o *ProjectTicketNote) SetIssueFlagNil(b bool)`

 SetIssueFlagNil sets the value for IssueFlag to be an explicit nil

### UnsetIssueFlag
`func (o *ProjectTicketNote) UnsetIssueFlag()`

UnsetIssueFlag ensures that no value is present for IssueFlag, not even an explicit nil
### GetOriginalAuthor

`func (o *ProjectTicketNote) GetOriginalAuthor() string`

GetOriginalAuthor returns the OriginalAuthor field if non-nil, zero value otherwise.

### GetOriginalAuthorOk

`func (o *ProjectTicketNote) GetOriginalAuthorOk() (*string, bool)`

GetOriginalAuthorOk returns a tuple with the OriginalAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAuthor

`func (o *ProjectTicketNote) SetOriginalAuthor(v string)`

SetOriginalAuthor sets OriginalAuthor field to given value.

### HasOriginalAuthor

`func (o *ProjectTicketNote) HasOriginalAuthor() bool`

HasOriginalAuthor returns a boolean if a field has been set.

### GetMember

`func (o *ProjectTicketNote) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ProjectTicketNote) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ProjectTicketNote) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *ProjectTicketNote) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetContact

`func (o *ProjectTicketNote) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ProjectTicketNote) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ProjectTicketNote) SetContact(v ContactReference)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ProjectTicketNote) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetInfo

`func (o *ProjectTicketNote) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProjectTicketNote) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProjectTicketNote) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProjectTicketNote) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


