# UserDefinedField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the custom user defined field | [optional] 
**PodId** | **NullableInt32** | Id of the Pod where the custom field will be placed | 
**Caption** | **string** | Field caption Max length: 25; | 
**SequenceNumber** | **NullableInt32** | Must be between 1 and 500.  This defines the order in which the custom fields will appear | 
**ScreenId** | Pointer to **string** | Field ScreenID Max length: 25; | [optional] 
**HelpText** | Pointer to **string** | Help text to accompany the custom field Max length: 1000; | [optional] 
**FieldTypeIdentifier** | **NullableString** |  | 
**NumberDecimals** | Pointer to **NullableInt32** | Only valid for Number or percent | [optional] 
**EntryTypeIdentifier** | Pointer to **NullableString** |  | [optional] 
**RequiredFlag** | Pointer to **NullableBool** |  | [optional] 
**DisplayOnScreenFlag** | Pointer to **NullableBool** |  | [optional] 
**ReadOnlyFlag** | Pointer to **NullableBool** |  | [optional] 
**ListViewFlag** | Pointer to **NullableBool** | Denotes that this custom field is included on a list view | [optional] 
**ButtonUrl** | Pointer to **string** | Only available with Button Field Type. Required when entryTypeIdentifier is button Max length: 1000; | [optional] 
**Options** | Pointer to [**[]UserDefinedFieldOption**](UserDefinedFieldOption.md) |  | [optional] 
**BusinessUnitIds** | Pointer to **[]int32** |  | [optional] 
**LocationIds** | Pointer to **[]int32** |  | [optional] 
**AddAllBusinessUnits** | Pointer to **NullableBool** |  | [optional] 
**RemoveAllBusinessUnits** | Pointer to **NullableBool** |  | [optional] 
**AddAllLocations** | Pointer to **NullableBool** |  | [optional] 
**RemoveAllLocations** | Pointer to **NullableBool** |  | [optional] 
**ConnectWiseID** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** | Date in UTC the custom field was created | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewUserDefinedField

`func NewUserDefinedField(podId NullableInt32, caption string, sequenceNumber NullableInt32, fieldTypeIdentifier NullableString, ) *UserDefinedField`

NewUserDefinedField instantiates a new UserDefinedField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDefinedFieldWithDefaults

`func NewUserDefinedFieldWithDefaults() *UserDefinedField`

NewUserDefinedFieldWithDefaults instantiates a new UserDefinedField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserDefinedField) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDefinedField) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDefinedField) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserDefinedField) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPodId

`func (o *UserDefinedField) GetPodId() int32`

GetPodId returns the PodId field if non-nil, zero value otherwise.

### GetPodIdOk

`func (o *UserDefinedField) GetPodIdOk() (*int32, bool)`

GetPodIdOk returns a tuple with the PodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodId

`func (o *UserDefinedField) SetPodId(v int32)`

SetPodId sets PodId field to given value.


### SetPodIdNil

`func (o *UserDefinedField) SetPodIdNil(b bool)`

 SetPodIdNil sets the value for PodId to be an explicit nil

### UnsetPodId
`func (o *UserDefinedField) UnsetPodId()`

UnsetPodId ensures that no value is present for PodId, not even an explicit nil
### GetCaption

`func (o *UserDefinedField) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *UserDefinedField) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *UserDefinedField) SetCaption(v string)`

SetCaption sets Caption field to given value.


### GetSequenceNumber

`func (o *UserDefinedField) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *UserDefinedField) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *UserDefinedField) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.


### SetSequenceNumberNil

`func (o *UserDefinedField) SetSequenceNumberNil(b bool)`

 SetSequenceNumberNil sets the value for SequenceNumber to be an explicit nil

### UnsetSequenceNumber
`func (o *UserDefinedField) UnsetSequenceNumber()`

UnsetSequenceNumber ensures that no value is present for SequenceNumber, not even an explicit nil
### GetScreenId

`func (o *UserDefinedField) GetScreenId() string`

GetScreenId returns the ScreenId field if non-nil, zero value otherwise.

### GetScreenIdOk

`func (o *UserDefinedField) GetScreenIdOk() (*string, bool)`

GetScreenIdOk returns a tuple with the ScreenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenId

`func (o *UserDefinedField) SetScreenId(v string)`

SetScreenId sets ScreenId field to given value.

### HasScreenId

`func (o *UserDefinedField) HasScreenId() bool`

HasScreenId returns a boolean if a field has been set.

### GetHelpText

`func (o *UserDefinedField) GetHelpText() string`

GetHelpText returns the HelpText field if non-nil, zero value otherwise.

### GetHelpTextOk

`func (o *UserDefinedField) GetHelpTextOk() (*string, bool)`

GetHelpTextOk returns a tuple with the HelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpText

`func (o *UserDefinedField) SetHelpText(v string)`

SetHelpText sets HelpText field to given value.

### HasHelpText

`func (o *UserDefinedField) HasHelpText() bool`

HasHelpText returns a boolean if a field has been set.

### GetFieldTypeIdentifier

`func (o *UserDefinedField) GetFieldTypeIdentifier() string`

GetFieldTypeIdentifier returns the FieldTypeIdentifier field if non-nil, zero value otherwise.

### GetFieldTypeIdentifierOk

`func (o *UserDefinedField) GetFieldTypeIdentifierOk() (*string, bool)`

GetFieldTypeIdentifierOk returns a tuple with the FieldTypeIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTypeIdentifier

`func (o *UserDefinedField) SetFieldTypeIdentifier(v string)`

SetFieldTypeIdentifier sets FieldTypeIdentifier field to given value.


### SetFieldTypeIdentifierNil

`func (o *UserDefinedField) SetFieldTypeIdentifierNil(b bool)`

 SetFieldTypeIdentifierNil sets the value for FieldTypeIdentifier to be an explicit nil

### UnsetFieldTypeIdentifier
`func (o *UserDefinedField) UnsetFieldTypeIdentifier()`

UnsetFieldTypeIdentifier ensures that no value is present for FieldTypeIdentifier, not even an explicit nil
### GetNumberDecimals

`func (o *UserDefinedField) GetNumberDecimals() int32`

GetNumberDecimals returns the NumberDecimals field if non-nil, zero value otherwise.

### GetNumberDecimalsOk

`func (o *UserDefinedField) GetNumberDecimalsOk() (*int32, bool)`

GetNumberDecimalsOk returns a tuple with the NumberDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberDecimals

`func (o *UserDefinedField) SetNumberDecimals(v int32)`

SetNumberDecimals sets NumberDecimals field to given value.

### HasNumberDecimals

`func (o *UserDefinedField) HasNumberDecimals() bool`

HasNumberDecimals returns a boolean if a field has been set.

### SetNumberDecimalsNil

`func (o *UserDefinedField) SetNumberDecimalsNil(b bool)`

 SetNumberDecimalsNil sets the value for NumberDecimals to be an explicit nil

### UnsetNumberDecimals
`func (o *UserDefinedField) UnsetNumberDecimals()`

UnsetNumberDecimals ensures that no value is present for NumberDecimals, not even an explicit nil
### GetEntryTypeIdentifier

`func (o *UserDefinedField) GetEntryTypeIdentifier() string`

GetEntryTypeIdentifier returns the EntryTypeIdentifier field if non-nil, zero value otherwise.

### GetEntryTypeIdentifierOk

`func (o *UserDefinedField) GetEntryTypeIdentifierOk() (*string, bool)`

GetEntryTypeIdentifierOk returns a tuple with the EntryTypeIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryTypeIdentifier

`func (o *UserDefinedField) SetEntryTypeIdentifier(v string)`

SetEntryTypeIdentifier sets EntryTypeIdentifier field to given value.

### HasEntryTypeIdentifier

`func (o *UserDefinedField) HasEntryTypeIdentifier() bool`

HasEntryTypeIdentifier returns a boolean if a field has been set.

### SetEntryTypeIdentifierNil

`func (o *UserDefinedField) SetEntryTypeIdentifierNil(b bool)`

 SetEntryTypeIdentifierNil sets the value for EntryTypeIdentifier to be an explicit nil

### UnsetEntryTypeIdentifier
`func (o *UserDefinedField) UnsetEntryTypeIdentifier()`

UnsetEntryTypeIdentifier ensures that no value is present for EntryTypeIdentifier, not even an explicit nil
### GetRequiredFlag

`func (o *UserDefinedField) GetRequiredFlag() bool`

GetRequiredFlag returns the RequiredFlag field if non-nil, zero value otherwise.

### GetRequiredFlagOk

`func (o *UserDefinedField) GetRequiredFlagOk() (*bool, bool)`

GetRequiredFlagOk returns a tuple with the RequiredFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFlag

`func (o *UserDefinedField) SetRequiredFlag(v bool)`

SetRequiredFlag sets RequiredFlag field to given value.

### HasRequiredFlag

`func (o *UserDefinedField) HasRequiredFlag() bool`

HasRequiredFlag returns a boolean if a field has been set.

### SetRequiredFlagNil

`func (o *UserDefinedField) SetRequiredFlagNil(b bool)`

 SetRequiredFlagNil sets the value for RequiredFlag to be an explicit nil

### UnsetRequiredFlag
`func (o *UserDefinedField) UnsetRequiredFlag()`

UnsetRequiredFlag ensures that no value is present for RequiredFlag, not even an explicit nil
### GetDisplayOnScreenFlag

`func (o *UserDefinedField) GetDisplayOnScreenFlag() bool`

GetDisplayOnScreenFlag returns the DisplayOnScreenFlag field if non-nil, zero value otherwise.

### GetDisplayOnScreenFlagOk

`func (o *UserDefinedField) GetDisplayOnScreenFlagOk() (*bool, bool)`

GetDisplayOnScreenFlagOk returns a tuple with the DisplayOnScreenFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOnScreenFlag

`func (o *UserDefinedField) SetDisplayOnScreenFlag(v bool)`

SetDisplayOnScreenFlag sets DisplayOnScreenFlag field to given value.

### HasDisplayOnScreenFlag

`func (o *UserDefinedField) HasDisplayOnScreenFlag() bool`

HasDisplayOnScreenFlag returns a boolean if a field has been set.

### SetDisplayOnScreenFlagNil

`func (o *UserDefinedField) SetDisplayOnScreenFlagNil(b bool)`

 SetDisplayOnScreenFlagNil sets the value for DisplayOnScreenFlag to be an explicit nil

### UnsetDisplayOnScreenFlag
`func (o *UserDefinedField) UnsetDisplayOnScreenFlag()`

UnsetDisplayOnScreenFlag ensures that no value is present for DisplayOnScreenFlag, not even an explicit nil
### GetReadOnlyFlag

`func (o *UserDefinedField) GetReadOnlyFlag() bool`

GetReadOnlyFlag returns the ReadOnlyFlag field if non-nil, zero value otherwise.

### GetReadOnlyFlagOk

`func (o *UserDefinedField) GetReadOnlyFlagOk() (*bool, bool)`

GetReadOnlyFlagOk returns a tuple with the ReadOnlyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyFlag

`func (o *UserDefinedField) SetReadOnlyFlag(v bool)`

SetReadOnlyFlag sets ReadOnlyFlag field to given value.

### HasReadOnlyFlag

`func (o *UserDefinedField) HasReadOnlyFlag() bool`

HasReadOnlyFlag returns a boolean if a field has been set.

### SetReadOnlyFlagNil

`func (o *UserDefinedField) SetReadOnlyFlagNil(b bool)`

 SetReadOnlyFlagNil sets the value for ReadOnlyFlag to be an explicit nil

### UnsetReadOnlyFlag
`func (o *UserDefinedField) UnsetReadOnlyFlag()`

UnsetReadOnlyFlag ensures that no value is present for ReadOnlyFlag, not even an explicit nil
### GetListViewFlag

`func (o *UserDefinedField) GetListViewFlag() bool`

GetListViewFlag returns the ListViewFlag field if non-nil, zero value otherwise.

### GetListViewFlagOk

`func (o *UserDefinedField) GetListViewFlagOk() (*bool, bool)`

GetListViewFlagOk returns a tuple with the ListViewFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListViewFlag

`func (o *UserDefinedField) SetListViewFlag(v bool)`

SetListViewFlag sets ListViewFlag field to given value.

### HasListViewFlag

`func (o *UserDefinedField) HasListViewFlag() bool`

HasListViewFlag returns a boolean if a field has been set.

### SetListViewFlagNil

`func (o *UserDefinedField) SetListViewFlagNil(b bool)`

 SetListViewFlagNil sets the value for ListViewFlag to be an explicit nil

### UnsetListViewFlag
`func (o *UserDefinedField) UnsetListViewFlag()`

UnsetListViewFlag ensures that no value is present for ListViewFlag, not even an explicit nil
### GetButtonUrl

`func (o *UserDefinedField) GetButtonUrl() string`

GetButtonUrl returns the ButtonUrl field if non-nil, zero value otherwise.

### GetButtonUrlOk

`func (o *UserDefinedField) GetButtonUrlOk() (*string, bool)`

GetButtonUrlOk returns a tuple with the ButtonUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonUrl

`func (o *UserDefinedField) SetButtonUrl(v string)`

SetButtonUrl sets ButtonUrl field to given value.

### HasButtonUrl

`func (o *UserDefinedField) HasButtonUrl() bool`

HasButtonUrl returns a boolean if a field has been set.

### GetOptions

`func (o *UserDefinedField) GetOptions() []UserDefinedFieldOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UserDefinedField) GetOptionsOk() (*[]UserDefinedFieldOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UserDefinedField) SetOptions(v []UserDefinedFieldOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UserDefinedField) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetBusinessUnitIds

`func (o *UserDefinedField) GetBusinessUnitIds() []int32`

GetBusinessUnitIds returns the BusinessUnitIds field if non-nil, zero value otherwise.

### GetBusinessUnitIdsOk

`func (o *UserDefinedField) GetBusinessUnitIdsOk() (*[]int32, bool)`

GetBusinessUnitIdsOk returns a tuple with the BusinessUnitIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitIds

`func (o *UserDefinedField) SetBusinessUnitIds(v []int32)`

SetBusinessUnitIds sets BusinessUnitIds field to given value.

### HasBusinessUnitIds

`func (o *UserDefinedField) HasBusinessUnitIds() bool`

HasBusinessUnitIds returns a boolean if a field has been set.

### GetLocationIds

`func (o *UserDefinedField) GetLocationIds() []int32`

GetLocationIds returns the LocationIds field if non-nil, zero value otherwise.

### GetLocationIdsOk

`func (o *UserDefinedField) GetLocationIdsOk() (*[]int32, bool)`

GetLocationIdsOk returns a tuple with the LocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationIds

`func (o *UserDefinedField) SetLocationIds(v []int32)`

SetLocationIds sets LocationIds field to given value.

### HasLocationIds

`func (o *UserDefinedField) HasLocationIds() bool`

HasLocationIds returns a boolean if a field has been set.

### GetAddAllBusinessUnits

`func (o *UserDefinedField) GetAddAllBusinessUnits() bool`

GetAddAllBusinessUnits returns the AddAllBusinessUnits field if non-nil, zero value otherwise.

### GetAddAllBusinessUnitsOk

`func (o *UserDefinedField) GetAddAllBusinessUnitsOk() (*bool, bool)`

GetAddAllBusinessUnitsOk returns a tuple with the AddAllBusinessUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAllBusinessUnits

`func (o *UserDefinedField) SetAddAllBusinessUnits(v bool)`

SetAddAllBusinessUnits sets AddAllBusinessUnits field to given value.

### HasAddAllBusinessUnits

`func (o *UserDefinedField) HasAddAllBusinessUnits() bool`

HasAddAllBusinessUnits returns a boolean if a field has been set.

### SetAddAllBusinessUnitsNil

`func (o *UserDefinedField) SetAddAllBusinessUnitsNil(b bool)`

 SetAddAllBusinessUnitsNil sets the value for AddAllBusinessUnits to be an explicit nil

### UnsetAddAllBusinessUnits
`func (o *UserDefinedField) UnsetAddAllBusinessUnits()`

UnsetAddAllBusinessUnits ensures that no value is present for AddAllBusinessUnits, not even an explicit nil
### GetRemoveAllBusinessUnits

`func (o *UserDefinedField) GetRemoveAllBusinessUnits() bool`

GetRemoveAllBusinessUnits returns the RemoveAllBusinessUnits field if non-nil, zero value otherwise.

### GetRemoveAllBusinessUnitsOk

`func (o *UserDefinedField) GetRemoveAllBusinessUnitsOk() (*bool, bool)`

GetRemoveAllBusinessUnitsOk returns a tuple with the RemoveAllBusinessUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAllBusinessUnits

`func (o *UserDefinedField) SetRemoveAllBusinessUnits(v bool)`

SetRemoveAllBusinessUnits sets RemoveAllBusinessUnits field to given value.

### HasRemoveAllBusinessUnits

`func (o *UserDefinedField) HasRemoveAllBusinessUnits() bool`

HasRemoveAllBusinessUnits returns a boolean if a field has been set.

### SetRemoveAllBusinessUnitsNil

`func (o *UserDefinedField) SetRemoveAllBusinessUnitsNil(b bool)`

 SetRemoveAllBusinessUnitsNil sets the value for RemoveAllBusinessUnits to be an explicit nil

### UnsetRemoveAllBusinessUnits
`func (o *UserDefinedField) UnsetRemoveAllBusinessUnits()`

UnsetRemoveAllBusinessUnits ensures that no value is present for RemoveAllBusinessUnits, not even an explicit nil
### GetAddAllLocations

`func (o *UserDefinedField) GetAddAllLocations() bool`

GetAddAllLocations returns the AddAllLocations field if non-nil, zero value otherwise.

### GetAddAllLocationsOk

`func (o *UserDefinedField) GetAddAllLocationsOk() (*bool, bool)`

GetAddAllLocationsOk returns a tuple with the AddAllLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAllLocations

`func (o *UserDefinedField) SetAddAllLocations(v bool)`

SetAddAllLocations sets AddAllLocations field to given value.

### HasAddAllLocations

`func (o *UserDefinedField) HasAddAllLocations() bool`

HasAddAllLocations returns a boolean if a field has been set.

### SetAddAllLocationsNil

`func (o *UserDefinedField) SetAddAllLocationsNil(b bool)`

 SetAddAllLocationsNil sets the value for AddAllLocations to be an explicit nil

### UnsetAddAllLocations
`func (o *UserDefinedField) UnsetAddAllLocations()`

UnsetAddAllLocations ensures that no value is present for AddAllLocations, not even an explicit nil
### GetRemoveAllLocations

`func (o *UserDefinedField) GetRemoveAllLocations() bool`

GetRemoveAllLocations returns the RemoveAllLocations field if non-nil, zero value otherwise.

### GetRemoveAllLocationsOk

`func (o *UserDefinedField) GetRemoveAllLocationsOk() (*bool, bool)`

GetRemoveAllLocationsOk returns a tuple with the RemoveAllLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAllLocations

`func (o *UserDefinedField) SetRemoveAllLocations(v bool)`

SetRemoveAllLocations sets RemoveAllLocations field to given value.

### HasRemoveAllLocations

`func (o *UserDefinedField) HasRemoveAllLocations() bool`

HasRemoveAllLocations returns a boolean if a field has been set.

### SetRemoveAllLocationsNil

`func (o *UserDefinedField) SetRemoveAllLocationsNil(b bool)`

 SetRemoveAllLocationsNil sets the value for RemoveAllLocations to be an explicit nil

### UnsetRemoveAllLocations
`func (o *UserDefinedField) UnsetRemoveAllLocations()`

UnsetRemoveAllLocations ensures that no value is present for RemoveAllLocations, not even an explicit nil
### GetConnectWiseID

`func (o *UserDefinedField) GetConnectWiseID() string`

GetConnectWiseID returns the ConnectWiseID field if non-nil, zero value otherwise.

### GetConnectWiseIDOk

`func (o *UserDefinedField) GetConnectWiseIDOk() (*string, bool)`

GetConnectWiseIDOk returns a tuple with the ConnectWiseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectWiseID

`func (o *UserDefinedField) SetConnectWiseID(v string)`

SetConnectWiseID sets ConnectWiseID field to given value.

### HasConnectWiseID

`func (o *UserDefinedField) HasConnectWiseID() bool`

HasConnectWiseID returns a boolean if a field has been set.

### GetDateCreated

`func (o *UserDefinedField) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UserDefinedField) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UserDefinedField) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UserDefinedField) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetInfo

`func (o *UserDefinedField) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *UserDefinedField) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *UserDefinedField) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *UserDefinedField) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


