# Holiday

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**AllDayFlag** | Pointer to **NullableBool** | Can be set to false to set a holiday for specific hours (Defaults to True). | [optional] 
**Date** | **string** |  | 
**TimeStart** | Pointer to **string** |  | [optional] 
**TimeEnd** | Pointer to **string** |  | [optional] 
**HolidayList** | Pointer to [**HolidayListReference**](HolidayListReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewHoliday

`func NewHoliday(name string, date string, ) *Holiday`

NewHoliday instantiates a new Holiday object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHolidayWithDefaults

`func NewHolidayWithDefaults() *Holiday`

NewHolidayWithDefaults instantiates a new Holiday object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Holiday) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Holiday) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Holiday) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Holiday) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Holiday) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Holiday) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Holiday) SetName(v string)`

SetName sets Name field to given value.


### GetAllDayFlag

`func (o *Holiday) GetAllDayFlag() bool`

GetAllDayFlag returns the AllDayFlag field if non-nil, zero value otherwise.

### GetAllDayFlagOk

`func (o *Holiday) GetAllDayFlagOk() (*bool, bool)`

GetAllDayFlagOk returns a tuple with the AllDayFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDayFlag

`func (o *Holiday) SetAllDayFlag(v bool)`

SetAllDayFlag sets AllDayFlag field to given value.

### HasAllDayFlag

`func (o *Holiday) HasAllDayFlag() bool`

HasAllDayFlag returns a boolean if a field has been set.

### SetAllDayFlagNil

`func (o *Holiday) SetAllDayFlagNil(b bool)`

 SetAllDayFlagNil sets the value for AllDayFlag to be an explicit nil

### UnsetAllDayFlag
`func (o *Holiday) UnsetAllDayFlag()`

UnsetAllDayFlag ensures that no value is present for AllDayFlag, not even an explicit nil
### GetDate

`func (o *Holiday) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Holiday) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Holiday) SetDate(v string)`

SetDate sets Date field to given value.


### GetTimeStart

`func (o *Holiday) GetTimeStart() string`

GetTimeStart returns the TimeStart field if non-nil, zero value otherwise.

### GetTimeStartOk

`func (o *Holiday) GetTimeStartOk() (*string, bool)`

GetTimeStartOk returns a tuple with the TimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStart

`func (o *Holiday) SetTimeStart(v string)`

SetTimeStart sets TimeStart field to given value.

### HasTimeStart

`func (o *Holiday) HasTimeStart() bool`

HasTimeStart returns a boolean if a field has been set.

### GetTimeEnd

`func (o *Holiday) GetTimeEnd() string`

GetTimeEnd returns the TimeEnd field if non-nil, zero value otherwise.

### GetTimeEndOk

`func (o *Holiday) GetTimeEndOk() (*string, bool)`

GetTimeEndOk returns a tuple with the TimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEnd

`func (o *Holiday) SetTimeEnd(v string)`

SetTimeEnd sets TimeEnd field to given value.

### HasTimeEnd

`func (o *Holiday) HasTimeEnd() bool`

HasTimeEnd returns a boolean if a field has been set.

### GetHolidayList

`func (o *Holiday) GetHolidayList() HolidayListReference`

GetHolidayList returns the HolidayList field if non-nil, zero value otherwise.

### GetHolidayListOk

`func (o *Holiday) GetHolidayListOk() (*HolidayListReference, bool)`

GetHolidayListOk returns a tuple with the HolidayList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidayList

`func (o *Holiday) SetHolidayList(v HolidayListReference)`

SetHolidayList sets HolidayList field to given value.

### HasHolidayList

`func (o *Holiday) HasHolidayList() bool`

HasHolidayList returns a boolean if a field has been set.

### GetInfo

`func (o *Holiday) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Holiday) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Holiday) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Holiday) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


