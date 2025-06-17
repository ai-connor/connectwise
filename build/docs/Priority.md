# Priority

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**Color** | **NullableString** |  | 
**SortOrder** | Pointer to **NullableInt32** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**ImageLink** | Pointer to **string** |  | [optional] 
**UrgencySortOrder** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **NullableString** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPriority

`func NewPriority(name string, color NullableString, ) *Priority`

NewPriority instantiates a new Priority object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriorityWithDefaults

`func NewPriorityWithDefaults() *Priority`

NewPriorityWithDefaults instantiates a new Priority object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Priority) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Priority) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Priority) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Priority) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Priority) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Priority) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Priority) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *Priority) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Priority) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Priority) SetColor(v string)`

SetColor sets Color field to given value.


### SetColorNil

`func (o *Priority) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *Priority) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetSortOrder

`func (o *Priority) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *Priority) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *Priority) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *Priority) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *Priority) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *Priority) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetDefaultFlag

`func (o *Priority) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *Priority) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *Priority) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *Priority) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *Priority) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *Priority) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetImageLink

`func (o *Priority) GetImageLink() string`

GetImageLink returns the ImageLink field if non-nil, zero value otherwise.

### GetImageLinkOk

`func (o *Priority) GetImageLinkOk() (*string, bool)`

GetImageLinkOk returns a tuple with the ImageLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageLink

`func (o *Priority) SetImageLink(v string)`

SetImageLink sets ImageLink field to given value.

### HasImageLink

`func (o *Priority) HasImageLink() bool`

HasImageLink returns a boolean if a field has been set.

### GetUrgencySortOrder

`func (o *Priority) GetUrgencySortOrder() string`

GetUrgencySortOrder returns the UrgencySortOrder field if non-nil, zero value otherwise.

### GetUrgencySortOrderOk

`func (o *Priority) GetUrgencySortOrderOk() (*string, bool)`

GetUrgencySortOrderOk returns a tuple with the UrgencySortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrgencySortOrder

`func (o *Priority) SetUrgencySortOrder(v string)`

SetUrgencySortOrder sets UrgencySortOrder field to given value.

### HasUrgencySortOrder

`func (o *Priority) HasUrgencySortOrder() bool`

HasUrgencySortOrder returns a boolean if a field has been set.

### GetLevel

`func (o *Priority) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *Priority) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *Priority) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *Priority) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### SetLevelNil

`func (o *Priority) SetLevelNil(b bool)`

 SetLevelNil sets the value for Level to be an explicit nil

### UnsetLevel
`func (o *Priority) UnsetLevel()`

UnsetLevel ensures that no value is present for Level, not even an explicit nil
### GetInfo

`func (o *Priority) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Priority) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Priority) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Priority) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


