# HolidayInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AllDayFlag** | Pointer to **NullableBool** | Can be set to false to set a holiday for specific hours (Defaults to True). | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**TimeStart** | Pointer to **string** |  | [optional] 
**TimeEnd** | Pointer to **string** |  | [optional] 
**HolidayList** | Pointer to [**HolidayListReference**](HolidayListReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewHolidayInfo

`func NewHolidayInfo() *HolidayInfo`

NewHolidayInfo instantiates a new HolidayInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHolidayInfoWithDefaults

`func NewHolidayInfoWithDefaults() *HolidayInfo`

NewHolidayInfoWithDefaults instantiates a new HolidayInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HolidayInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HolidayInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HolidayInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *HolidayInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *HolidayInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HolidayInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HolidayInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HolidayInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAllDayFlag

`func (o *HolidayInfo) GetAllDayFlag() bool`

GetAllDayFlag returns the AllDayFlag field if non-nil, zero value otherwise.

### GetAllDayFlagOk

`func (o *HolidayInfo) GetAllDayFlagOk() (*bool, bool)`

GetAllDayFlagOk returns a tuple with the AllDayFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDayFlag

`func (o *HolidayInfo) SetAllDayFlag(v bool)`

SetAllDayFlag sets AllDayFlag field to given value.

### HasAllDayFlag

`func (o *HolidayInfo) HasAllDayFlag() bool`

HasAllDayFlag returns a boolean if a field has been set.

### SetAllDayFlagNil

`func (o *HolidayInfo) SetAllDayFlagNil(b bool)`

 SetAllDayFlagNil sets the value for AllDayFlag to be an explicit nil

### UnsetAllDayFlag
`func (o *HolidayInfo) UnsetAllDayFlag()`

UnsetAllDayFlag ensures that no value is present for AllDayFlag, not even an explicit nil
### GetDate

`func (o *HolidayInfo) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *HolidayInfo) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *HolidayInfo) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *HolidayInfo) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetTimeStart

`func (o *HolidayInfo) GetTimeStart() string`

GetTimeStart returns the TimeStart field if non-nil, zero value otherwise.

### GetTimeStartOk

`func (o *HolidayInfo) GetTimeStartOk() (*string, bool)`

GetTimeStartOk returns a tuple with the TimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStart

`func (o *HolidayInfo) SetTimeStart(v string)`

SetTimeStart sets TimeStart field to given value.

### HasTimeStart

`func (o *HolidayInfo) HasTimeStart() bool`

HasTimeStart returns a boolean if a field has been set.

### GetTimeEnd

`func (o *HolidayInfo) GetTimeEnd() string`

GetTimeEnd returns the TimeEnd field if non-nil, zero value otherwise.

### GetTimeEndOk

`func (o *HolidayInfo) GetTimeEndOk() (*string, bool)`

GetTimeEndOk returns a tuple with the TimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEnd

`func (o *HolidayInfo) SetTimeEnd(v string)`

SetTimeEnd sets TimeEnd field to given value.

### HasTimeEnd

`func (o *HolidayInfo) HasTimeEnd() bool`

HasTimeEnd returns a boolean if a field has been set.

### GetHolidayList

`func (o *HolidayInfo) GetHolidayList() HolidayListReference`

GetHolidayList returns the HolidayList field if non-nil, zero value otherwise.

### GetHolidayListOk

`func (o *HolidayInfo) GetHolidayListOk() (*HolidayListReference, bool)`

GetHolidayListOk returns a tuple with the HolidayList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidayList

`func (o *HolidayInfo) SetHolidayList(v HolidayListReference)`

SetHolidayList sets HolidayList field to given value.

### HasHolidayList

`func (o *HolidayInfo) HasHolidayList() bool`

HasHolidayList returns a boolean if a field has been set.

### GetInfo

`func (o *HolidayInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *HolidayInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *HolidayInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *HolidayInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


