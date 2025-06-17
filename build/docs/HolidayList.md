# HolidayList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewHolidayList

`func NewHolidayList(name string, ) *HolidayList`

NewHolidayList instantiates a new HolidayList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHolidayListWithDefaults

`func NewHolidayListWithDefaults() *HolidayList`

NewHolidayListWithDefaults instantiates a new HolidayList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HolidayList) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HolidayList) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HolidayList) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *HolidayList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *HolidayList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HolidayList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HolidayList) SetName(v string)`

SetName sets Name field to given value.


### GetInfo

`func (o *HolidayList) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *HolidayList) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *HolidayList) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *HolidayList) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


