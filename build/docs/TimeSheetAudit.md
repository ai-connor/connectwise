# TimeSheetAudit

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

### NewTimeSheetAudit

`func NewTimeSheetAudit() *TimeSheetAudit`

NewTimeSheetAudit instantiates a new TimeSheetAudit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSheetAuditWithDefaults

`func NewTimeSheetAuditWithDefaults() *TimeSheetAudit`

NewTimeSheetAuditWithDefaults instantiates a new TimeSheetAudit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimeSheetAudit) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimeSheetAudit) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimeSheetAudit) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TimeSheetAudit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMember

`func (o *TimeSheetAudit) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *TimeSheetAudit) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *TimeSheetAudit) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *TimeSheetAudit) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetSource

`func (o *TimeSheetAudit) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TimeSheetAudit) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TimeSheetAudit) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *TimeSheetAudit) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *TimeSheetAudit) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *TimeSheetAudit) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetType

`func (o *TimeSheetAudit) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimeSheetAudit) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimeSheetAudit) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TimeSheetAudit) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *TimeSheetAudit) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TimeSheetAudit) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetMessage

`func (o *TimeSheetAudit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TimeSheetAudit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TimeSheetAudit) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TimeSheetAudit) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOldValue

`func (o *TimeSheetAudit) GetOldValue() string`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *TimeSheetAudit) GetOldValueOk() (*string, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *TimeSheetAudit) SetOldValue(v string)`

SetOldValue sets OldValue field to given value.

### HasOldValue

`func (o *TimeSheetAudit) HasOldValue() bool`

HasOldValue returns a boolean if a field has been set.

### GetNewValue

`func (o *TimeSheetAudit) GetNewValue() string`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *TimeSheetAudit) GetNewValueOk() (*string, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *TimeSheetAudit) SetNewValue(v string)`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *TimeSheetAudit) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### GetValue

`func (o *TimeSheetAudit) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TimeSheetAudit) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TimeSheetAudit) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TimeSheetAudit) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetInfo

`func (o *TimeSheetAudit) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TimeSheetAudit) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TimeSheetAudit) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TimeSheetAudit) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


