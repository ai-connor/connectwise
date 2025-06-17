# PriorityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **NullableString** |  | [optional] 
**SortOrder** | Pointer to **NullableInt32** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**ImageLink** | Pointer to **string** |  | [optional] 
**UrgencySortOrder** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **NullableString** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPriorityInfo

`func NewPriorityInfo() *PriorityInfo`

NewPriorityInfo instantiates a new PriorityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriorityInfoWithDefaults

`func NewPriorityInfoWithDefaults() *PriorityInfo`

NewPriorityInfoWithDefaults instantiates a new PriorityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PriorityInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PriorityInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PriorityInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PriorityInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PriorityInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PriorityInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PriorityInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PriorityInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetColor

`func (o *PriorityInfo) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PriorityInfo) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PriorityInfo) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *PriorityInfo) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *PriorityInfo) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *PriorityInfo) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetSortOrder

`func (o *PriorityInfo) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *PriorityInfo) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *PriorityInfo) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *PriorityInfo) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *PriorityInfo) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *PriorityInfo) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetDefaultFlag

`func (o *PriorityInfo) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *PriorityInfo) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *PriorityInfo) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *PriorityInfo) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *PriorityInfo) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *PriorityInfo) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetImageLink

`func (o *PriorityInfo) GetImageLink() string`

GetImageLink returns the ImageLink field if non-nil, zero value otherwise.

### GetImageLinkOk

`func (o *PriorityInfo) GetImageLinkOk() (*string, bool)`

GetImageLinkOk returns a tuple with the ImageLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageLink

`func (o *PriorityInfo) SetImageLink(v string)`

SetImageLink sets ImageLink field to given value.

### HasImageLink

`func (o *PriorityInfo) HasImageLink() bool`

HasImageLink returns a boolean if a field has been set.

### GetUrgencySortOrder

`func (o *PriorityInfo) GetUrgencySortOrder() string`

GetUrgencySortOrder returns the UrgencySortOrder field if non-nil, zero value otherwise.

### GetUrgencySortOrderOk

`func (o *PriorityInfo) GetUrgencySortOrderOk() (*string, bool)`

GetUrgencySortOrderOk returns a tuple with the UrgencySortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrgencySortOrder

`func (o *PriorityInfo) SetUrgencySortOrder(v string)`

SetUrgencySortOrder sets UrgencySortOrder field to given value.

### HasUrgencySortOrder

`func (o *PriorityInfo) HasUrgencySortOrder() bool`

HasUrgencySortOrder returns a boolean if a field has been set.

### GetLevel

`func (o *PriorityInfo) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *PriorityInfo) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *PriorityInfo) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *PriorityInfo) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### SetLevelNil

`func (o *PriorityInfo) SetLevelNil(b bool)`

 SetLevelNil sets the value for Level to be an explicit nil

### UnsetLevel
`func (o *PriorityInfo) UnsetLevel()`

UnsetLevel ensures that no value is present for Level, not even an explicit nil
### GetInfo

`func (o *PriorityInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PriorityInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PriorityInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PriorityInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


