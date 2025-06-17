# TimeEntryAudit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**OldValue** | Pointer to **string** |  | [optional] 
**NewValue** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTimeEntryAudit

`func NewTimeEntryAudit() *TimeEntryAudit`

NewTimeEntryAudit instantiates a new TimeEntryAudit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeEntryAuditWithDefaults

`func NewTimeEntryAuditWithDefaults() *TimeEntryAudit`

NewTimeEntryAuditWithDefaults instantiates a new TimeEntryAudit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimeEntryAudit) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimeEntryAudit) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimeEntryAudit) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TimeEntryAudit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMember

`func (o *TimeEntryAudit) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *TimeEntryAudit) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *TimeEntryAudit) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *TimeEntryAudit) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetSource

`func (o *TimeEntryAudit) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TimeEntryAudit) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TimeEntryAudit) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *TimeEntryAudit) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *TimeEntryAudit) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *TimeEntryAudit) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetType

`func (o *TimeEntryAudit) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimeEntryAudit) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimeEntryAudit) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TimeEntryAudit) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *TimeEntryAudit) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TimeEntryAudit) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetMessage

`func (o *TimeEntryAudit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TimeEntryAudit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TimeEntryAudit) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TimeEntryAudit) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOldValue

`func (o *TimeEntryAudit) GetOldValue() string`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *TimeEntryAudit) GetOldValueOk() (*string, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *TimeEntryAudit) SetOldValue(v string)`

SetOldValue sets OldValue field to given value.

### HasOldValue

`func (o *TimeEntryAudit) HasOldValue() bool`

HasOldValue returns a boolean if a field has been set.

### GetNewValue

`func (o *TimeEntryAudit) GetNewValue() string`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *TimeEntryAudit) GetNewValueOk() (*string, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *TimeEntryAudit) SetNewValue(v string)`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *TimeEntryAudit) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### GetValue

`func (o *TimeEntryAudit) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TimeEntryAudit) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TimeEntryAudit) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TimeEntryAudit) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetInfo

`func (o *TimeEntryAudit) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TimeEntryAudit) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TimeEntryAudit) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TimeEntryAudit) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


