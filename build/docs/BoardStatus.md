# BoardStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**Board** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**SortOrder** | Pointer to **NullableInt32** |  | [optional] 
**DisplayOnBoard** | Pointer to **NullableBool** |  | [optional] 
**Inactive** | Pointer to **NullableBool** |  | [optional] 
**ClosedStatus** | Pointer to **NullableBool** |  | [optional] 
**TimeEntryNotAllowed** | Pointer to **NullableBool** |  | [optional] 
**RoundRobinCatchall** | Pointer to **NullableBool** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**EscalationStatus** | Pointer to **NullableString** |  | [optional] 
**CustomerPortalDescription** | Pointer to **string** |  Max length: 500; | [optional] 
**CustomerPortalFlag** | Pointer to **NullableBool** |  | [optional] 
**EmailTemplate** | Pointer to [**ServiceEmailTemplateReference**](ServiceEmailTemplateReference.md) |  | [optional] 
**StatusIndicator** | Pointer to [**StatusIndicatorReference**](StatusIndicatorReference.md) |  | [optional] 
**CustomStatusIndicatorName** | Pointer to **string** |  Max length: 30; | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**SaveTimeAsNote** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewBoardStatus

`func NewBoardStatus(name string, ) *BoardStatus`

NewBoardStatus instantiates a new BoardStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardStatusWithDefaults

`func NewBoardStatusWithDefaults() *BoardStatus`

NewBoardStatusWithDefaults instantiates a new BoardStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BoardStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoardStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoardStatus) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BoardStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BoardStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BoardStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BoardStatus) SetName(v string)`

SetName sets Name field to given value.


### GetBoard

`func (o *BoardStatus) GetBoard() BoardReference`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *BoardStatus) GetBoardOk() (*BoardReference, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *BoardStatus) SetBoard(v BoardReference)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *BoardStatus) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetSortOrder

`func (o *BoardStatus) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *BoardStatus) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *BoardStatus) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *BoardStatus) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *BoardStatus) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *BoardStatus) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetDisplayOnBoard

`func (o *BoardStatus) GetDisplayOnBoard() bool`

GetDisplayOnBoard returns the DisplayOnBoard field if non-nil, zero value otherwise.

### GetDisplayOnBoardOk

`func (o *BoardStatus) GetDisplayOnBoardOk() (*bool, bool)`

GetDisplayOnBoardOk returns a tuple with the DisplayOnBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOnBoard

`func (o *BoardStatus) SetDisplayOnBoard(v bool)`

SetDisplayOnBoard sets DisplayOnBoard field to given value.

### HasDisplayOnBoard

`func (o *BoardStatus) HasDisplayOnBoard() bool`

HasDisplayOnBoard returns a boolean if a field has been set.

### SetDisplayOnBoardNil

`func (o *BoardStatus) SetDisplayOnBoardNil(b bool)`

 SetDisplayOnBoardNil sets the value for DisplayOnBoard to be an explicit nil

### UnsetDisplayOnBoard
`func (o *BoardStatus) UnsetDisplayOnBoard()`

UnsetDisplayOnBoard ensures that no value is present for DisplayOnBoard, not even an explicit nil
### GetInactive

`func (o *BoardStatus) GetInactive() bool`

GetInactive returns the Inactive field if non-nil, zero value otherwise.

### GetInactiveOk

`func (o *BoardStatus) GetInactiveOk() (*bool, bool)`

GetInactiveOk returns a tuple with the Inactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactive

`func (o *BoardStatus) SetInactive(v bool)`

SetInactive sets Inactive field to given value.

### HasInactive

`func (o *BoardStatus) HasInactive() bool`

HasInactive returns a boolean if a field has been set.

### SetInactiveNil

`func (o *BoardStatus) SetInactiveNil(b bool)`

 SetInactiveNil sets the value for Inactive to be an explicit nil

### UnsetInactive
`func (o *BoardStatus) UnsetInactive()`

UnsetInactive ensures that no value is present for Inactive, not even an explicit nil
### GetClosedStatus

`func (o *BoardStatus) GetClosedStatus() bool`

GetClosedStatus returns the ClosedStatus field if non-nil, zero value otherwise.

### GetClosedStatusOk

`func (o *BoardStatus) GetClosedStatusOk() (*bool, bool)`

GetClosedStatusOk returns a tuple with the ClosedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedStatus

`func (o *BoardStatus) SetClosedStatus(v bool)`

SetClosedStatus sets ClosedStatus field to given value.

### HasClosedStatus

`func (o *BoardStatus) HasClosedStatus() bool`

HasClosedStatus returns a boolean if a field has been set.

### SetClosedStatusNil

`func (o *BoardStatus) SetClosedStatusNil(b bool)`

 SetClosedStatusNil sets the value for ClosedStatus to be an explicit nil

### UnsetClosedStatus
`func (o *BoardStatus) UnsetClosedStatus()`

UnsetClosedStatus ensures that no value is present for ClosedStatus, not even an explicit nil
### GetTimeEntryNotAllowed

`func (o *BoardStatus) GetTimeEntryNotAllowed() bool`

GetTimeEntryNotAllowed returns the TimeEntryNotAllowed field if non-nil, zero value otherwise.

### GetTimeEntryNotAllowedOk

`func (o *BoardStatus) GetTimeEntryNotAllowedOk() (*bool, bool)`

GetTimeEntryNotAllowedOk returns a tuple with the TimeEntryNotAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEntryNotAllowed

`func (o *BoardStatus) SetTimeEntryNotAllowed(v bool)`

SetTimeEntryNotAllowed sets TimeEntryNotAllowed field to given value.

### HasTimeEntryNotAllowed

`func (o *BoardStatus) HasTimeEntryNotAllowed() bool`

HasTimeEntryNotAllowed returns a boolean if a field has been set.

### SetTimeEntryNotAllowedNil

`func (o *BoardStatus) SetTimeEntryNotAllowedNil(b bool)`

 SetTimeEntryNotAllowedNil sets the value for TimeEntryNotAllowed to be an explicit nil

### UnsetTimeEntryNotAllowed
`func (o *BoardStatus) UnsetTimeEntryNotAllowed()`

UnsetTimeEntryNotAllowed ensures that no value is present for TimeEntryNotAllowed, not even an explicit nil
### GetRoundRobinCatchall

`func (o *BoardStatus) GetRoundRobinCatchall() bool`

GetRoundRobinCatchall returns the RoundRobinCatchall field if non-nil, zero value otherwise.

### GetRoundRobinCatchallOk

`func (o *BoardStatus) GetRoundRobinCatchallOk() (*bool, bool)`

GetRoundRobinCatchallOk returns a tuple with the RoundRobinCatchall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundRobinCatchall

`func (o *BoardStatus) SetRoundRobinCatchall(v bool)`

SetRoundRobinCatchall sets RoundRobinCatchall field to given value.

### HasRoundRobinCatchall

`func (o *BoardStatus) HasRoundRobinCatchall() bool`

HasRoundRobinCatchall returns a boolean if a field has been set.

### SetRoundRobinCatchallNil

`func (o *BoardStatus) SetRoundRobinCatchallNil(b bool)`

 SetRoundRobinCatchallNil sets the value for RoundRobinCatchall to be an explicit nil

### UnsetRoundRobinCatchall
`func (o *BoardStatus) UnsetRoundRobinCatchall()`

UnsetRoundRobinCatchall ensures that no value is present for RoundRobinCatchall, not even an explicit nil
### GetDefaultFlag

`func (o *BoardStatus) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *BoardStatus) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *BoardStatus) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *BoardStatus) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *BoardStatus) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *BoardStatus) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetEscalationStatus

`func (o *BoardStatus) GetEscalationStatus() string`

GetEscalationStatus returns the EscalationStatus field if non-nil, zero value otherwise.

### GetEscalationStatusOk

`func (o *BoardStatus) GetEscalationStatusOk() (*string, bool)`

GetEscalationStatusOk returns a tuple with the EscalationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscalationStatus

`func (o *BoardStatus) SetEscalationStatus(v string)`

SetEscalationStatus sets EscalationStatus field to given value.

### HasEscalationStatus

`func (o *BoardStatus) HasEscalationStatus() bool`

HasEscalationStatus returns a boolean if a field has been set.

### SetEscalationStatusNil

`func (o *BoardStatus) SetEscalationStatusNil(b bool)`

 SetEscalationStatusNil sets the value for EscalationStatus to be an explicit nil

### UnsetEscalationStatus
`func (o *BoardStatus) UnsetEscalationStatus()`

UnsetEscalationStatus ensures that no value is present for EscalationStatus, not even an explicit nil
### GetCustomerPortalDescription

`func (o *BoardStatus) GetCustomerPortalDescription() string`

GetCustomerPortalDescription returns the CustomerPortalDescription field if non-nil, zero value otherwise.

### GetCustomerPortalDescriptionOk

`func (o *BoardStatus) GetCustomerPortalDescriptionOk() (*string, bool)`

GetCustomerPortalDescriptionOk returns a tuple with the CustomerPortalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPortalDescription

`func (o *BoardStatus) SetCustomerPortalDescription(v string)`

SetCustomerPortalDescription sets CustomerPortalDescription field to given value.

### HasCustomerPortalDescription

`func (o *BoardStatus) HasCustomerPortalDescription() bool`

HasCustomerPortalDescription returns a boolean if a field has been set.

### GetCustomerPortalFlag

`func (o *BoardStatus) GetCustomerPortalFlag() bool`

GetCustomerPortalFlag returns the CustomerPortalFlag field if non-nil, zero value otherwise.

### GetCustomerPortalFlagOk

`func (o *BoardStatus) GetCustomerPortalFlagOk() (*bool, bool)`

GetCustomerPortalFlagOk returns a tuple with the CustomerPortalFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPortalFlag

`func (o *BoardStatus) SetCustomerPortalFlag(v bool)`

SetCustomerPortalFlag sets CustomerPortalFlag field to given value.

### HasCustomerPortalFlag

`func (o *BoardStatus) HasCustomerPortalFlag() bool`

HasCustomerPortalFlag returns a boolean if a field has been set.

### SetCustomerPortalFlagNil

`func (o *BoardStatus) SetCustomerPortalFlagNil(b bool)`

 SetCustomerPortalFlagNil sets the value for CustomerPortalFlag to be an explicit nil

### UnsetCustomerPortalFlag
`func (o *BoardStatus) UnsetCustomerPortalFlag()`

UnsetCustomerPortalFlag ensures that no value is present for CustomerPortalFlag, not even an explicit nil
### GetEmailTemplate

`func (o *BoardStatus) GetEmailTemplate() ServiceEmailTemplateReference`

GetEmailTemplate returns the EmailTemplate field if non-nil, zero value otherwise.

### GetEmailTemplateOk

`func (o *BoardStatus) GetEmailTemplateOk() (*ServiceEmailTemplateReference, bool)`

GetEmailTemplateOk returns a tuple with the EmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplate

`func (o *BoardStatus) SetEmailTemplate(v ServiceEmailTemplateReference)`

SetEmailTemplate sets EmailTemplate field to given value.

### HasEmailTemplate

`func (o *BoardStatus) HasEmailTemplate() bool`

HasEmailTemplate returns a boolean if a field has been set.

### GetStatusIndicator

`func (o *BoardStatus) GetStatusIndicator() StatusIndicatorReference`

GetStatusIndicator returns the StatusIndicator field if non-nil, zero value otherwise.

### GetStatusIndicatorOk

`func (o *BoardStatus) GetStatusIndicatorOk() (*StatusIndicatorReference, bool)`

GetStatusIndicatorOk returns a tuple with the StatusIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusIndicator

`func (o *BoardStatus) SetStatusIndicator(v StatusIndicatorReference)`

SetStatusIndicator sets StatusIndicator field to given value.

### HasStatusIndicator

`func (o *BoardStatus) HasStatusIndicator() bool`

HasStatusIndicator returns a boolean if a field has been set.

### GetCustomStatusIndicatorName

`func (o *BoardStatus) GetCustomStatusIndicatorName() string`

GetCustomStatusIndicatorName returns the CustomStatusIndicatorName field if non-nil, zero value otherwise.

### GetCustomStatusIndicatorNameOk

`func (o *BoardStatus) GetCustomStatusIndicatorNameOk() (*string, bool)`

GetCustomStatusIndicatorNameOk returns a tuple with the CustomStatusIndicatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStatusIndicatorName

`func (o *BoardStatus) SetCustomStatusIndicatorName(v string)`

SetCustomStatusIndicatorName sets CustomStatusIndicatorName field to given value.

### HasCustomStatusIndicatorName

`func (o *BoardStatus) HasCustomStatusIndicatorName() bool`

HasCustomStatusIndicatorName returns a boolean if a field has been set.

### GetInfo

`func (o *BoardStatus) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BoardStatus) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BoardStatus) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BoardStatus) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetSaveTimeAsNote

`func (o *BoardStatus) GetSaveTimeAsNote() bool`

GetSaveTimeAsNote returns the SaveTimeAsNote field if non-nil, zero value otherwise.

### GetSaveTimeAsNoteOk

`func (o *BoardStatus) GetSaveTimeAsNoteOk() (*bool, bool)`

GetSaveTimeAsNoteOk returns a tuple with the SaveTimeAsNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveTimeAsNote

`func (o *BoardStatus) SetSaveTimeAsNote(v bool)`

SetSaveTimeAsNote sets SaveTimeAsNote field to given value.

### HasSaveTimeAsNote

`func (o *BoardStatus) HasSaveTimeAsNote() bool`

HasSaveTimeAsNote returns a boolean if a field has been set.

### SetSaveTimeAsNoteNil

`func (o *BoardStatus) SetSaveTimeAsNoteNil(b bool)`

 SetSaveTimeAsNoteNil sets the value for SaveTimeAsNote to be an explicit nil

### UnsetSaveTimeAsNote
`func (o *BoardStatus) UnsetSaveTimeAsNote()`

UnsetSaveTimeAsNote ensures that no value is present for SaveTimeAsNote, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


