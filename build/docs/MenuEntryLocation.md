# MenuEntryLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Location** | [**SystemLocationReference**](SystemLocationReference.md) |  | 
**MenuEntry** | Pointer to [**SystemMenuEntryReference**](SystemMenuEntryReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMenuEntryLocation

`func NewMenuEntryLocation(location SystemLocationReference, ) *MenuEntryLocation`

NewMenuEntryLocation instantiates a new MenuEntryLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuEntryLocationWithDefaults

`func NewMenuEntryLocationWithDefaults() *MenuEntryLocation`

NewMenuEntryLocationWithDefaults instantiates a new MenuEntryLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MenuEntryLocation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MenuEntryLocation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MenuEntryLocation) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MenuEntryLocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *MenuEntryLocation) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MenuEntryLocation) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MenuEntryLocation) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.


### GetMenuEntry

`func (o *MenuEntryLocation) GetMenuEntry() SystemMenuEntryReference`

GetMenuEntry returns the MenuEntry field if non-nil, zero value otherwise.

### GetMenuEntryOk

`func (o *MenuEntryLocation) GetMenuEntryOk() (*SystemMenuEntryReference, bool)`

GetMenuEntryOk returns a tuple with the MenuEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenuEntry

`func (o *MenuEntryLocation) SetMenuEntry(v SystemMenuEntryReference)`

SetMenuEntry sets MenuEntry field to given value.

### HasMenuEntry

`func (o *MenuEntryLocation) HasMenuEntry() bool`

HasMenuEntry returns a boolean if a field has been set.

### GetInfo

`func (o *MenuEntryLocation) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *MenuEntryLocation) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *MenuEntryLocation) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *MenuEntryLocation) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


