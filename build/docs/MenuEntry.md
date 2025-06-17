# MenuEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**MenuLocation** | [**MenuLocationReference**](MenuLocationReference.md) |  | 
**Caption** | **string** |  Max length: 50; | 
**Link** | **string** |  Max length: 2000; | 
**NewWindowFlag** | **NullableBool** |  | 
**LocationIds** | Pointer to **[]int32** |  | [optional] 
**Origin** | Pointer to **string** |  Max length: 2000; | [optional] 
**ClientId** | Pointer to **string** | Only required if not already set Max length: 36; | [optional] 
**AddAllLocations** | Pointer to **NullableBool** |  | [optional] 
**RemoveAllLocations** | Pointer to **NullableBool** |  | [optional] 
**SmallMenuIconId** | Pointer to **NullableInt32** |  | [optional] 
**LargeMenuIconId** | Pointer to **NullableInt32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMenuEntry

`func NewMenuEntry(menuLocation MenuLocationReference, caption string, link string, newWindowFlag NullableBool, ) *MenuEntry`

NewMenuEntry instantiates a new MenuEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuEntryWithDefaults

`func NewMenuEntryWithDefaults() *MenuEntry`

NewMenuEntryWithDefaults instantiates a new MenuEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MenuEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MenuEntry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MenuEntry) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MenuEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMenuLocation

`func (o *MenuEntry) GetMenuLocation() MenuLocationReference`

GetMenuLocation returns the MenuLocation field if non-nil, zero value otherwise.

### GetMenuLocationOk

`func (o *MenuEntry) GetMenuLocationOk() (*MenuLocationReference, bool)`

GetMenuLocationOk returns a tuple with the MenuLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenuLocation

`func (o *MenuEntry) SetMenuLocation(v MenuLocationReference)`

SetMenuLocation sets MenuLocation field to given value.


### GetCaption

`func (o *MenuEntry) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *MenuEntry) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *MenuEntry) SetCaption(v string)`

SetCaption sets Caption field to given value.


### GetLink

`func (o *MenuEntry) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *MenuEntry) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *MenuEntry) SetLink(v string)`

SetLink sets Link field to given value.


### GetNewWindowFlag

`func (o *MenuEntry) GetNewWindowFlag() bool`

GetNewWindowFlag returns the NewWindowFlag field if non-nil, zero value otherwise.

### GetNewWindowFlagOk

`func (o *MenuEntry) GetNewWindowFlagOk() (*bool, bool)`

GetNewWindowFlagOk returns a tuple with the NewWindowFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewWindowFlag

`func (o *MenuEntry) SetNewWindowFlag(v bool)`

SetNewWindowFlag sets NewWindowFlag field to given value.


### SetNewWindowFlagNil

`func (o *MenuEntry) SetNewWindowFlagNil(b bool)`

 SetNewWindowFlagNil sets the value for NewWindowFlag to be an explicit nil

### UnsetNewWindowFlag
`func (o *MenuEntry) UnsetNewWindowFlag()`

UnsetNewWindowFlag ensures that no value is present for NewWindowFlag, not even an explicit nil
### GetLocationIds

`func (o *MenuEntry) GetLocationIds() []int32`

GetLocationIds returns the LocationIds field if non-nil, zero value otherwise.

### GetLocationIdsOk

`func (o *MenuEntry) GetLocationIdsOk() (*[]int32, bool)`

GetLocationIdsOk returns a tuple with the LocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationIds

`func (o *MenuEntry) SetLocationIds(v []int32)`

SetLocationIds sets LocationIds field to given value.

### HasLocationIds

`func (o *MenuEntry) HasLocationIds() bool`

HasLocationIds returns a boolean if a field has been set.

### GetOrigin

`func (o *MenuEntry) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *MenuEntry) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *MenuEntry) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *MenuEntry) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetClientId

`func (o *MenuEntry) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *MenuEntry) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *MenuEntry) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *MenuEntry) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetAddAllLocations

`func (o *MenuEntry) GetAddAllLocations() bool`

GetAddAllLocations returns the AddAllLocations field if non-nil, zero value otherwise.

### GetAddAllLocationsOk

`func (o *MenuEntry) GetAddAllLocationsOk() (*bool, bool)`

GetAddAllLocationsOk returns a tuple with the AddAllLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAllLocations

`func (o *MenuEntry) SetAddAllLocations(v bool)`

SetAddAllLocations sets AddAllLocations field to given value.

### HasAddAllLocations

`func (o *MenuEntry) HasAddAllLocations() bool`

HasAddAllLocations returns a boolean if a field has been set.

### SetAddAllLocationsNil

`func (o *MenuEntry) SetAddAllLocationsNil(b bool)`

 SetAddAllLocationsNil sets the value for AddAllLocations to be an explicit nil

### UnsetAddAllLocations
`func (o *MenuEntry) UnsetAddAllLocations()`

UnsetAddAllLocations ensures that no value is present for AddAllLocations, not even an explicit nil
### GetRemoveAllLocations

`func (o *MenuEntry) GetRemoveAllLocations() bool`

GetRemoveAllLocations returns the RemoveAllLocations field if non-nil, zero value otherwise.

### GetRemoveAllLocationsOk

`func (o *MenuEntry) GetRemoveAllLocationsOk() (*bool, bool)`

GetRemoveAllLocationsOk returns a tuple with the RemoveAllLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAllLocations

`func (o *MenuEntry) SetRemoveAllLocations(v bool)`

SetRemoveAllLocations sets RemoveAllLocations field to given value.

### HasRemoveAllLocations

`func (o *MenuEntry) HasRemoveAllLocations() bool`

HasRemoveAllLocations returns a boolean if a field has been set.

### SetRemoveAllLocationsNil

`func (o *MenuEntry) SetRemoveAllLocationsNil(b bool)`

 SetRemoveAllLocationsNil sets the value for RemoveAllLocations to be an explicit nil

### UnsetRemoveAllLocations
`func (o *MenuEntry) UnsetRemoveAllLocations()`

UnsetRemoveAllLocations ensures that no value is present for RemoveAllLocations, not even an explicit nil
### GetSmallMenuIconId

`func (o *MenuEntry) GetSmallMenuIconId() int32`

GetSmallMenuIconId returns the SmallMenuIconId field if non-nil, zero value otherwise.

### GetSmallMenuIconIdOk

`func (o *MenuEntry) GetSmallMenuIconIdOk() (*int32, bool)`

GetSmallMenuIconIdOk returns a tuple with the SmallMenuIconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallMenuIconId

`func (o *MenuEntry) SetSmallMenuIconId(v int32)`

SetSmallMenuIconId sets SmallMenuIconId field to given value.

### HasSmallMenuIconId

`func (o *MenuEntry) HasSmallMenuIconId() bool`

HasSmallMenuIconId returns a boolean if a field has been set.

### SetSmallMenuIconIdNil

`func (o *MenuEntry) SetSmallMenuIconIdNil(b bool)`

 SetSmallMenuIconIdNil sets the value for SmallMenuIconId to be an explicit nil

### UnsetSmallMenuIconId
`func (o *MenuEntry) UnsetSmallMenuIconId()`

UnsetSmallMenuIconId ensures that no value is present for SmallMenuIconId, not even an explicit nil
### GetLargeMenuIconId

`func (o *MenuEntry) GetLargeMenuIconId() int32`

GetLargeMenuIconId returns the LargeMenuIconId field if non-nil, zero value otherwise.

### GetLargeMenuIconIdOk

`func (o *MenuEntry) GetLargeMenuIconIdOk() (*int32, bool)`

GetLargeMenuIconIdOk returns a tuple with the LargeMenuIconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeMenuIconId

`func (o *MenuEntry) SetLargeMenuIconId(v int32)`

SetLargeMenuIconId sets LargeMenuIconId field to given value.

### HasLargeMenuIconId

`func (o *MenuEntry) HasLargeMenuIconId() bool`

HasLargeMenuIconId returns a boolean if a field has been set.

### SetLargeMenuIconIdNil

`func (o *MenuEntry) SetLargeMenuIconIdNil(b bool)`

 SetLargeMenuIconIdNil sets the value for LargeMenuIconId to be an explicit nil

### UnsetLargeMenuIconId
`func (o *MenuEntry) UnsetLargeMenuIconId()`

UnsetLargeMenuIconId ensures that no value is present for LargeMenuIconId, not even an explicit nil
### GetInfo

`func (o *MenuEntry) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *MenuEntry) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *MenuEntry) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *MenuEntry) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


