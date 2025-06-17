# WorkflowActionUserDefinedField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**EventId** | Pointer to **int32** |  | [optional] 
**ActionId** | Pointer to **int32** |  | [optional] 
**Caption** | Pointer to **string** |  | [optional] 
**UserDefinedFieldId** | Pointer to **int32** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**OverwriteFlag** | Pointer to **NullableBool** |  | [optional] 
**PodDescription** | Pointer to **string** |  | [optional] 
**FieldTypeId** | Pointer to **string** |  | [optional] 
**EntryTypeId** | Pointer to **string** |  | [optional] 
**RequiredFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**ConnectWiseID** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **NullableInt32** | WF_NotifyActions_RecID | [optional] 
**GrandParentId** | Pointer to **NullableInt32** | WF_NotifyEvents_RecID | [optional] 
**ParentConnectWiseId** | Pointer to **string** |  | [optional] 
**GrandParentConnectWiseId** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewWorkflowActionUserDefinedField

`func NewWorkflowActionUserDefinedField() *WorkflowActionUserDefinedField`

NewWorkflowActionUserDefinedField instantiates a new WorkflowActionUserDefinedField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowActionUserDefinedFieldWithDefaults

`func NewWorkflowActionUserDefinedFieldWithDefaults() *WorkflowActionUserDefinedField`

NewWorkflowActionUserDefinedFieldWithDefaults instantiates a new WorkflowActionUserDefinedField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowActionUserDefinedField) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowActionUserDefinedField) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowActionUserDefinedField) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowActionUserDefinedField) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEventId

`func (o *WorkflowActionUserDefinedField) GetEventId() int32`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *WorkflowActionUserDefinedField) GetEventIdOk() (*int32, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *WorkflowActionUserDefinedField) SetEventId(v int32)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *WorkflowActionUserDefinedField) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetActionId

`func (o *WorkflowActionUserDefinedField) GetActionId() int32`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *WorkflowActionUserDefinedField) GetActionIdOk() (*int32, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *WorkflowActionUserDefinedField) SetActionId(v int32)`

SetActionId sets ActionId field to given value.

### HasActionId

`func (o *WorkflowActionUserDefinedField) HasActionId() bool`

HasActionId returns a boolean if a field has been set.

### GetCaption

`func (o *WorkflowActionUserDefinedField) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *WorkflowActionUserDefinedField) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *WorkflowActionUserDefinedField) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *WorkflowActionUserDefinedField) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetUserDefinedFieldId

`func (o *WorkflowActionUserDefinedField) GetUserDefinedFieldId() int32`

GetUserDefinedFieldId returns the UserDefinedFieldId field if non-nil, zero value otherwise.

### GetUserDefinedFieldIdOk

`func (o *WorkflowActionUserDefinedField) GetUserDefinedFieldIdOk() (*int32, bool)`

GetUserDefinedFieldIdOk returns a tuple with the UserDefinedFieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFieldId

`func (o *WorkflowActionUserDefinedField) SetUserDefinedFieldId(v int32)`

SetUserDefinedFieldId sets UserDefinedFieldId field to given value.

### HasUserDefinedFieldId

`func (o *WorkflowActionUserDefinedField) HasUserDefinedFieldId() bool`

HasUserDefinedFieldId returns a boolean if a field has been set.

### GetValue

`func (o *WorkflowActionUserDefinedField) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WorkflowActionUserDefinedField) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WorkflowActionUserDefinedField) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *WorkflowActionUserDefinedField) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOverwriteFlag

`func (o *WorkflowActionUserDefinedField) GetOverwriteFlag() bool`

GetOverwriteFlag returns the OverwriteFlag field if non-nil, zero value otherwise.

### GetOverwriteFlagOk

`func (o *WorkflowActionUserDefinedField) GetOverwriteFlagOk() (*bool, bool)`

GetOverwriteFlagOk returns a tuple with the OverwriteFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteFlag

`func (o *WorkflowActionUserDefinedField) SetOverwriteFlag(v bool)`

SetOverwriteFlag sets OverwriteFlag field to given value.

### HasOverwriteFlag

`func (o *WorkflowActionUserDefinedField) HasOverwriteFlag() bool`

HasOverwriteFlag returns a boolean if a field has been set.

### SetOverwriteFlagNil

`func (o *WorkflowActionUserDefinedField) SetOverwriteFlagNil(b bool)`

 SetOverwriteFlagNil sets the value for OverwriteFlag to be an explicit nil

### UnsetOverwriteFlag
`func (o *WorkflowActionUserDefinedField) UnsetOverwriteFlag()`

UnsetOverwriteFlag ensures that no value is present for OverwriteFlag, not even an explicit nil
### GetPodDescription

`func (o *WorkflowActionUserDefinedField) GetPodDescription() string`

GetPodDescription returns the PodDescription field if non-nil, zero value otherwise.

### GetPodDescriptionOk

`func (o *WorkflowActionUserDefinedField) GetPodDescriptionOk() (*string, bool)`

GetPodDescriptionOk returns a tuple with the PodDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodDescription

`func (o *WorkflowActionUserDefinedField) SetPodDescription(v string)`

SetPodDescription sets PodDescription field to given value.

### HasPodDescription

`func (o *WorkflowActionUserDefinedField) HasPodDescription() bool`

HasPodDescription returns a boolean if a field has been set.

### GetFieldTypeId

`func (o *WorkflowActionUserDefinedField) GetFieldTypeId() string`

GetFieldTypeId returns the FieldTypeId field if non-nil, zero value otherwise.

### GetFieldTypeIdOk

`func (o *WorkflowActionUserDefinedField) GetFieldTypeIdOk() (*string, bool)`

GetFieldTypeIdOk returns a tuple with the FieldTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTypeId

`func (o *WorkflowActionUserDefinedField) SetFieldTypeId(v string)`

SetFieldTypeId sets FieldTypeId field to given value.

### HasFieldTypeId

`func (o *WorkflowActionUserDefinedField) HasFieldTypeId() bool`

HasFieldTypeId returns a boolean if a field has been set.

### GetEntryTypeId

`func (o *WorkflowActionUserDefinedField) GetEntryTypeId() string`

GetEntryTypeId returns the EntryTypeId field if non-nil, zero value otherwise.

### GetEntryTypeIdOk

`func (o *WorkflowActionUserDefinedField) GetEntryTypeIdOk() (*string, bool)`

GetEntryTypeIdOk returns a tuple with the EntryTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryTypeId

`func (o *WorkflowActionUserDefinedField) SetEntryTypeId(v string)`

SetEntryTypeId sets EntryTypeId field to given value.

### HasEntryTypeId

`func (o *WorkflowActionUserDefinedField) HasEntryTypeId() bool`

HasEntryTypeId returns a boolean if a field has been set.

### GetRequiredFlag

`func (o *WorkflowActionUserDefinedField) GetRequiredFlag() bool`

GetRequiredFlag returns the RequiredFlag field if non-nil, zero value otherwise.

### GetRequiredFlagOk

`func (o *WorkflowActionUserDefinedField) GetRequiredFlagOk() (*bool, bool)`

GetRequiredFlagOk returns a tuple with the RequiredFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFlag

`func (o *WorkflowActionUserDefinedField) SetRequiredFlag(v bool)`

SetRequiredFlag sets RequiredFlag field to given value.

### HasRequiredFlag

`func (o *WorkflowActionUserDefinedField) HasRequiredFlag() bool`

HasRequiredFlag returns a boolean if a field has been set.

### SetRequiredFlagNil

`func (o *WorkflowActionUserDefinedField) SetRequiredFlagNil(b bool)`

 SetRequiredFlagNil sets the value for RequiredFlag to be an explicit nil

### UnsetRequiredFlag
`func (o *WorkflowActionUserDefinedField) UnsetRequiredFlag()`

UnsetRequiredFlag ensures that no value is present for RequiredFlag, not even an explicit nil
### GetInactiveFlag

`func (o *WorkflowActionUserDefinedField) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *WorkflowActionUserDefinedField) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *WorkflowActionUserDefinedField) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *WorkflowActionUserDefinedField) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *WorkflowActionUserDefinedField) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *WorkflowActionUserDefinedField) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetConnectWiseID

`func (o *WorkflowActionUserDefinedField) GetConnectWiseID() string`

GetConnectWiseID returns the ConnectWiseID field if non-nil, zero value otherwise.

### GetConnectWiseIDOk

`func (o *WorkflowActionUserDefinedField) GetConnectWiseIDOk() (*string, bool)`

GetConnectWiseIDOk returns a tuple with the ConnectWiseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectWiseID

`func (o *WorkflowActionUserDefinedField) SetConnectWiseID(v string)`

SetConnectWiseID sets ConnectWiseID field to given value.

### HasConnectWiseID

`func (o *WorkflowActionUserDefinedField) HasConnectWiseID() bool`

HasConnectWiseID returns a boolean if a field has been set.

### GetParentId

`func (o *WorkflowActionUserDefinedField) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *WorkflowActionUserDefinedField) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *WorkflowActionUserDefinedField) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *WorkflowActionUserDefinedField) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *WorkflowActionUserDefinedField) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *WorkflowActionUserDefinedField) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetGrandParentId

`func (o *WorkflowActionUserDefinedField) GetGrandParentId() int32`

GetGrandParentId returns the GrandParentId field if non-nil, zero value otherwise.

### GetGrandParentIdOk

`func (o *WorkflowActionUserDefinedField) GetGrandParentIdOk() (*int32, bool)`

GetGrandParentIdOk returns a tuple with the GrandParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrandParentId

`func (o *WorkflowActionUserDefinedField) SetGrandParentId(v int32)`

SetGrandParentId sets GrandParentId field to given value.

### HasGrandParentId

`func (o *WorkflowActionUserDefinedField) HasGrandParentId() bool`

HasGrandParentId returns a boolean if a field has been set.

### SetGrandParentIdNil

`func (o *WorkflowActionUserDefinedField) SetGrandParentIdNil(b bool)`

 SetGrandParentIdNil sets the value for GrandParentId to be an explicit nil

### UnsetGrandParentId
`func (o *WorkflowActionUserDefinedField) UnsetGrandParentId()`

UnsetGrandParentId ensures that no value is present for GrandParentId, not even an explicit nil
### GetParentConnectWiseId

`func (o *WorkflowActionUserDefinedField) GetParentConnectWiseId() string`

GetParentConnectWiseId returns the ParentConnectWiseId field if non-nil, zero value otherwise.

### GetParentConnectWiseIdOk

`func (o *WorkflowActionUserDefinedField) GetParentConnectWiseIdOk() (*string, bool)`

GetParentConnectWiseIdOk returns a tuple with the ParentConnectWiseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConnectWiseId

`func (o *WorkflowActionUserDefinedField) SetParentConnectWiseId(v string)`

SetParentConnectWiseId sets ParentConnectWiseId field to given value.

### HasParentConnectWiseId

`func (o *WorkflowActionUserDefinedField) HasParentConnectWiseId() bool`

HasParentConnectWiseId returns a boolean if a field has been set.

### GetGrandParentConnectWiseId

`func (o *WorkflowActionUserDefinedField) GetGrandParentConnectWiseId() string`

GetGrandParentConnectWiseId returns the GrandParentConnectWiseId field if non-nil, zero value otherwise.

### GetGrandParentConnectWiseIdOk

`func (o *WorkflowActionUserDefinedField) GetGrandParentConnectWiseIdOk() (*string, bool)`

GetGrandParentConnectWiseIdOk returns a tuple with the GrandParentConnectWiseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrandParentConnectWiseId

`func (o *WorkflowActionUserDefinedField) SetGrandParentConnectWiseId(v string)`

SetGrandParentConnectWiseId sets GrandParentConnectWiseId field to given value.

### HasGrandParentConnectWiseId

`func (o *WorkflowActionUserDefinedField) HasGrandParentConnectWiseId() bool`

HasGrandParentConnectWiseId returns a boolean if a field has been set.

### GetInfo

`func (o *WorkflowActionUserDefinedField) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *WorkflowActionUserDefinedField) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *WorkflowActionUserDefinedField) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *WorkflowActionUserDefinedField) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


