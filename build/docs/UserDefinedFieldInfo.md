# UserDefinedFieldInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the custom user defined field | [optional] 
**PodId** | Pointer to **NullableInt32** | Id of the Pod where the custom field will be placed | [optional] 
**Caption** | Pointer to **string** | Field caption | [optional] 
**SequenceNumber** | Pointer to **NullableInt32** | Must be between 1 and 500.  This defines the order in which the custom fields will appear | [optional] 
**HelpText** | Pointer to **string** | Help text to accompany the custom field | [optional] 
**FieldTypeIdentifier** | Pointer to **NullableString** |  | [optional] 
**NumberDecimals** | Pointer to **NullableInt32** | Only valid for Number or percent | [optional] 
**EntryTypeIdentifier** | Pointer to **NullableString** |  | [optional] 
**RequiredFlag** | Pointer to **NullableBool** |  | [optional] 
**DisplayOnScreenFlag** | Pointer to **NullableBool** |  | [optional] 
**ReadOnlyFlag** | Pointer to **NullableBool** |  | [optional] 
**ListViewFlag** | Pointer to **NullableBool** | Denotes that this custom field is included on a list view | [optional] 
**ButtonUrl** | Pointer to **string** | Only available with Button Field Type. Required when entryTypeIdentifier is button | [optional] 
**Options** | Pointer to [**[]UserDefinedFieldOption**](UserDefinedFieldOption.md) |  | [optional] 
**BusinessUnitIds** | Pointer to **[]int32** | List of business unit ids using custom field | [optional] 
**LocationIds** | Pointer to **[]int32** | List of locations ids using custom field | [optional] 
**DateCreated** | Pointer to **string** | Date in UTC the custom field was created | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewUserDefinedFieldInfo

`func NewUserDefinedFieldInfo() *UserDefinedFieldInfo`

NewUserDefinedFieldInfo instantiates a new UserDefinedFieldInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDefinedFieldInfoWithDefaults

`func NewUserDefinedFieldInfoWithDefaults() *UserDefinedFieldInfo`

NewUserDefinedFieldInfoWithDefaults instantiates a new UserDefinedFieldInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserDefinedFieldInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDefinedFieldInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDefinedFieldInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserDefinedFieldInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPodId

`func (o *UserDefinedFieldInfo) GetPodId() int32`

GetPodId returns the PodId field if non-nil, zero value otherwise.

### GetPodIdOk

`func (o *UserDefinedFieldInfo) GetPodIdOk() (*int32, bool)`

GetPodIdOk returns a tuple with the PodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodId

`func (o *UserDefinedFieldInfo) SetPodId(v int32)`

SetPodId sets PodId field to given value.

### HasPodId

`func (o *UserDefinedFieldInfo) HasPodId() bool`

HasPodId returns a boolean if a field has been set.

### SetPodIdNil

`func (o *UserDefinedFieldInfo) SetPodIdNil(b bool)`

 SetPodIdNil sets the value for PodId to be an explicit nil

### UnsetPodId
`func (o *UserDefinedFieldInfo) UnsetPodId()`

UnsetPodId ensures that no value is present for PodId, not even an explicit nil
### GetCaption

`func (o *UserDefinedFieldInfo) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *UserDefinedFieldInfo) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *UserDefinedFieldInfo) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *UserDefinedFieldInfo) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *UserDefinedFieldInfo) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *UserDefinedFieldInfo) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *UserDefinedFieldInfo) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *UserDefinedFieldInfo) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### SetSequenceNumberNil

`func (o *UserDefinedFieldInfo) SetSequenceNumberNil(b bool)`

 SetSequenceNumberNil sets the value for SequenceNumber to be an explicit nil

### UnsetSequenceNumber
`func (o *UserDefinedFieldInfo) UnsetSequenceNumber()`

UnsetSequenceNumber ensures that no value is present for SequenceNumber, not even an explicit nil
### GetHelpText

`func (o *UserDefinedFieldInfo) GetHelpText() string`

GetHelpText returns the HelpText field if non-nil, zero value otherwise.

### GetHelpTextOk

`func (o *UserDefinedFieldInfo) GetHelpTextOk() (*string, bool)`

GetHelpTextOk returns a tuple with the HelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpText

`func (o *UserDefinedFieldInfo) SetHelpText(v string)`

SetHelpText sets HelpText field to given value.

### HasHelpText

`func (o *UserDefinedFieldInfo) HasHelpText() bool`

HasHelpText returns a boolean if a field has been set.

### GetFieldTypeIdentifier

`func (o *UserDefinedFieldInfo) GetFieldTypeIdentifier() string`

GetFieldTypeIdentifier returns the FieldTypeIdentifier field if non-nil, zero value otherwise.

### GetFieldTypeIdentifierOk

`func (o *UserDefinedFieldInfo) GetFieldTypeIdentifierOk() (*string, bool)`

GetFieldTypeIdentifierOk returns a tuple with the FieldTypeIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTypeIdentifier

`func (o *UserDefinedFieldInfo) SetFieldTypeIdentifier(v string)`

SetFieldTypeIdentifier sets FieldTypeIdentifier field to given value.

### HasFieldTypeIdentifier

`func (o *UserDefinedFieldInfo) HasFieldTypeIdentifier() bool`

HasFieldTypeIdentifier returns a boolean if a field has been set.

### SetFieldTypeIdentifierNil

`func (o *UserDefinedFieldInfo) SetFieldTypeIdentifierNil(b bool)`

 SetFieldTypeIdentifierNil sets the value for FieldTypeIdentifier to be an explicit nil

### UnsetFieldTypeIdentifier
`func (o *UserDefinedFieldInfo) UnsetFieldTypeIdentifier()`

UnsetFieldTypeIdentifier ensures that no value is present for FieldTypeIdentifier, not even an explicit nil
### GetNumberDecimals

`func (o *UserDefinedFieldInfo) GetNumberDecimals() int32`

GetNumberDecimals returns the NumberDecimals field if non-nil, zero value otherwise.

### GetNumberDecimalsOk

`func (o *UserDefinedFieldInfo) GetNumberDecimalsOk() (*int32, bool)`

GetNumberDecimalsOk returns a tuple with the NumberDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberDecimals

`func (o *UserDefinedFieldInfo) SetNumberDecimals(v int32)`

SetNumberDecimals sets NumberDecimals field to given value.

### HasNumberDecimals

`func (o *UserDefinedFieldInfo) HasNumberDecimals() bool`

HasNumberDecimals returns a boolean if a field has been set.

### SetNumberDecimalsNil

`func (o *UserDefinedFieldInfo) SetNumberDecimalsNil(b bool)`

 SetNumberDecimalsNil sets the value for NumberDecimals to be an explicit nil

### UnsetNumberDecimals
`func (o *UserDefinedFieldInfo) UnsetNumberDecimals()`

UnsetNumberDecimals ensures that no value is present for NumberDecimals, not even an explicit nil
### GetEntryTypeIdentifier

`func (o *UserDefinedFieldInfo) GetEntryTypeIdentifier() string`

GetEntryTypeIdentifier returns the EntryTypeIdentifier field if non-nil, zero value otherwise.

### GetEntryTypeIdentifierOk

`func (o *UserDefinedFieldInfo) GetEntryTypeIdentifierOk() (*string, bool)`

GetEntryTypeIdentifierOk returns a tuple with the EntryTypeIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryTypeIdentifier

`func (o *UserDefinedFieldInfo) SetEntryTypeIdentifier(v string)`

SetEntryTypeIdentifier sets EntryTypeIdentifier field to given value.

### HasEntryTypeIdentifier

`func (o *UserDefinedFieldInfo) HasEntryTypeIdentifier() bool`

HasEntryTypeIdentifier returns a boolean if a field has been set.

### SetEntryTypeIdentifierNil

`func (o *UserDefinedFieldInfo) SetEntryTypeIdentifierNil(b bool)`

 SetEntryTypeIdentifierNil sets the value for EntryTypeIdentifier to be an explicit nil

### UnsetEntryTypeIdentifier
`func (o *UserDefinedFieldInfo) UnsetEntryTypeIdentifier()`

UnsetEntryTypeIdentifier ensures that no value is present for EntryTypeIdentifier, not even an explicit nil
### GetRequiredFlag

`func (o *UserDefinedFieldInfo) GetRequiredFlag() bool`

GetRequiredFlag returns the RequiredFlag field if non-nil, zero value otherwise.

### GetRequiredFlagOk

`func (o *UserDefinedFieldInfo) GetRequiredFlagOk() (*bool, bool)`

GetRequiredFlagOk returns a tuple with the RequiredFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFlag

`func (o *UserDefinedFieldInfo) SetRequiredFlag(v bool)`

SetRequiredFlag sets RequiredFlag field to given value.

### HasRequiredFlag

`func (o *UserDefinedFieldInfo) HasRequiredFlag() bool`

HasRequiredFlag returns a boolean if a field has been set.

### SetRequiredFlagNil

`func (o *UserDefinedFieldInfo) SetRequiredFlagNil(b bool)`

 SetRequiredFlagNil sets the value for RequiredFlag to be an explicit nil

### UnsetRequiredFlag
`func (o *UserDefinedFieldInfo) UnsetRequiredFlag()`

UnsetRequiredFlag ensures that no value is present for RequiredFlag, not even an explicit nil
### GetDisplayOnScreenFlag

`func (o *UserDefinedFieldInfo) GetDisplayOnScreenFlag() bool`

GetDisplayOnScreenFlag returns the DisplayOnScreenFlag field if non-nil, zero value otherwise.

### GetDisplayOnScreenFlagOk

`func (o *UserDefinedFieldInfo) GetDisplayOnScreenFlagOk() (*bool, bool)`

GetDisplayOnScreenFlagOk returns a tuple with the DisplayOnScreenFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOnScreenFlag

`func (o *UserDefinedFieldInfo) SetDisplayOnScreenFlag(v bool)`

SetDisplayOnScreenFlag sets DisplayOnScreenFlag field to given value.

### HasDisplayOnScreenFlag

`func (o *UserDefinedFieldInfo) HasDisplayOnScreenFlag() bool`

HasDisplayOnScreenFlag returns a boolean if a field has been set.

### SetDisplayOnScreenFlagNil

`func (o *UserDefinedFieldInfo) SetDisplayOnScreenFlagNil(b bool)`

 SetDisplayOnScreenFlagNil sets the value for DisplayOnScreenFlag to be an explicit nil

### UnsetDisplayOnScreenFlag
`func (o *UserDefinedFieldInfo) UnsetDisplayOnScreenFlag()`

UnsetDisplayOnScreenFlag ensures that no value is present for DisplayOnScreenFlag, not even an explicit nil
### GetReadOnlyFlag

`func (o *UserDefinedFieldInfo) GetReadOnlyFlag() bool`

GetReadOnlyFlag returns the ReadOnlyFlag field if non-nil, zero value otherwise.

### GetReadOnlyFlagOk

`func (o *UserDefinedFieldInfo) GetReadOnlyFlagOk() (*bool, bool)`

GetReadOnlyFlagOk returns a tuple with the ReadOnlyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyFlag

`func (o *UserDefinedFieldInfo) SetReadOnlyFlag(v bool)`

SetReadOnlyFlag sets ReadOnlyFlag field to given value.

### HasReadOnlyFlag

`func (o *UserDefinedFieldInfo) HasReadOnlyFlag() bool`

HasReadOnlyFlag returns a boolean if a field has been set.

### SetReadOnlyFlagNil

`func (o *UserDefinedFieldInfo) SetReadOnlyFlagNil(b bool)`

 SetReadOnlyFlagNil sets the value for ReadOnlyFlag to be an explicit nil

### UnsetReadOnlyFlag
`func (o *UserDefinedFieldInfo) UnsetReadOnlyFlag()`

UnsetReadOnlyFlag ensures that no value is present for ReadOnlyFlag, not even an explicit nil
### GetListViewFlag

`func (o *UserDefinedFieldInfo) GetListViewFlag() bool`

GetListViewFlag returns the ListViewFlag field if non-nil, zero value otherwise.

### GetListViewFlagOk

`func (o *UserDefinedFieldInfo) GetListViewFlagOk() (*bool, bool)`

GetListViewFlagOk returns a tuple with the ListViewFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListViewFlag

`func (o *UserDefinedFieldInfo) SetListViewFlag(v bool)`

SetListViewFlag sets ListViewFlag field to given value.

### HasListViewFlag

`func (o *UserDefinedFieldInfo) HasListViewFlag() bool`

HasListViewFlag returns a boolean if a field has been set.

### SetListViewFlagNil

`func (o *UserDefinedFieldInfo) SetListViewFlagNil(b bool)`

 SetListViewFlagNil sets the value for ListViewFlag to be an explicit nil

### UnsetListViewFlag
`func (o *UserDefinedFieldInfo) UnsetListViewFlag()`

UnsetListViewFlag ensures that no value is present for ListViewFlag, not even an explicit nil
### GetButtonUrl

`func (o *UserDefinedFieldInfo) GetButtonUrl() string`

GetButtonUrl returns the ButtonUrl field if non-nil, zero value otherwise.

### GetButtonUrlOk

`func (o *UserDefinedFieldInfo) GetButtonUrlOk() (*string, bool)`

GetButtonUrlOk returns a tuple with the ButtonUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonUrl

`func (o *UserDefinedFieldInfo) SetButtonUrl(v string)`

SetButtonUrl sets ButtonUrl field to given value.

### HasButtonUrl

`func (o *UserDefinedFieldInfo) HasButtonUrl() bool`

HasButtonUrl returns a boolean if a field has been set.

### GetOptions

`func (o *UserDefinedFieldInfo) GetOptions() []UserDefinedFieldOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UserDefinedFieldInfo) GetOptionsOk() (*[]UserDefinedFieldOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UserDefinedFieldInfo) SetOptions(v []UserDefinedFieldOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UserDefinedFieldInfo) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetBusinessUnitIds

`func (o *UserDefinedFieldInfo) GetBusinessUnitIds() []int32`

GetBusinessUnitIds returns the BusinessUnitIds field if non-nil, zero value otherwise.

### GetBusinessUnitIdsOk

`func (o *UserDefinedFieldInfo) GetBusinessUnitIdsOk() (*[]int32, bool)`

GetBusinessUnitIdsOk returns a tuple with the BusinessUnitIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitIds

`func (o *UserDefinedFieldInfo) SetBusinessUnitIds(v []int32)`

SetBusinessUnitIds sets BusinessUnitIds field to given value.

### HasBusinessUnitIds

`func (o *UserDefinedFieldInfo) HasBusinessUnitIds() bool`

HasBusinessUnitIds returns a boolean if a field has been set.

### GetLocationIds

`func (o *UserDefinedFieldInfo) GetLocationIds() []int32`

GetLocationIds returns the LocationIds field if non-nil, zero value otherwise.

### GetLocationIdsOk

`func (o *UserDefinedFieldInfo) GetLocationIdsOk() (*[]int32, bool)`

GetLocationIdsOk returns a tuple with the LocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationIds

`func (o *UserDefinedFieldInfo) SetLocationIds(v []int32)`

SetLocationIds sets LocationIds field to given value.

### HasLocationIds

`func (o *UserDefinedFieldInfo) HasLocationIds() bool`

HasLocationIds returns a boolean if a field has been set.

### GetDateCreated

`func (o *UserDefinedFieldInfo) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UserDefinedFieldInfo) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UserDefinedFieldInfo) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UserDefinedFieldInfo) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetInfo

`func (o *UserDefinedFieldInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *UserDefinedFieldInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *UserDefinedFieldInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *UserDefinedFieldInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


