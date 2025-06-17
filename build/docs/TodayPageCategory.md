# TodayPageCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 30; | 
**SortOrder** | **NullableInt32** |  | 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTodayPageCategory

`func NewTodayPageCategory(name string, sortOrder NullableInt32, ) *TodayPageCategory`

NewTodayPageCategory instantiates a new TodayPageCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTodayPageCategoryWithDefaults

`func NewTodayPageCategoryWithDefaults() *TodayPageCategory`

NewTodayPageCategoryWithDefaults instantiates a new TodayPageCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TodayPageCategory) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TodayPageCategory) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TodayPageCategory) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TodayPageCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TodayPageCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TodayPageCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TodayPageCategory) SetName(v string)`

SetName sets Name field to given value.


### GetSortOrder

`func (o *TodayPageCategory) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *TodayPageCategory) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *TodayPageCategory) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.


### SetSortOrderNil

`func (o *TodayPageCategory) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *TodayPageCategory) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetLocation

`func (o *TodayPageCategory) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *TodayPageCategory) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *TodayPageCategory) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *TodayPageCategory) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetInfo

`func (o *TodayPageCategory) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TodayPageCategory) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TodayPageCategory) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TodayPageCategory) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


