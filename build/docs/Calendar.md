# Calendar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**HolidayList** | Pointer to [**HolidayListReference**](HolidayListReference.md) |  | [optional] 
**MondayStartTime** | Pointer to **string** |  | [optional] 
**MondayEndTime** | Pointer to **string** |  | [optional] 
**TuesdayStartTime** | Pointer to **string** |  | [optional] 
**TuesdayEndTime** | Pointer to **string** |  | [optional] 
**WednesdayStartTime** | Pointer to **string** |  | [optional] 
**WednesdayEndTime** | Pointer to **string** |  | [optional] 
**ThursdayStartTime** | Pointer to **string** |  | [optional] 
**ThursdayEndTime** | Pointer to **string** |  | [optional] 
**FridayStartTime** | Pointer to **string** |  | [optional] 
**FridayEndTime** | Pointer to **string** |  | [optional] 
**SaturdayStartTime** | Pointer to **string** |  | [optional] 
**SaturdayEndTime** | Pointer to **string** |  | [optional] 
**SundayStartTime** | Pointer to **string** |  | [optional] 
**SundayEndTime** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCalendar

`func NewCalendar(name string, ) *Calendar`

NewCalendar instantiates a new Calendar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarWithDefaults

`func NewCalendarWithDefaults() *Calendar`

NewCalendarWithDefaults instantiates a new Calendar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Calendar) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Calendar) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Calendar) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Calendar) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Calendar) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Calendar) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Calendar) SetName(v string)`

SetName sets Name field to given value.


### GetHolidayList

`func (o *Calendar) GetHolidayList() HolidayListReference`

GetHolidayList returns the HolidayList field if non-nil, zero value otherwise.

### GetHolidayListOk

`func (o *Calendar) GetHolidayListOk() (*HolidayListReference, bool)`

GetHolidayListOk returns a tuple with the HolidayList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidayList

`func (o *Calendar) SetHolidayList(v HolidayListReference)`

SetHolidayList sets HolidayList field to given value.

### HasHolidayList

`func (o *Calendar) HasHolidayList() bool`

HasHolidayList returns a boolean if a field has been set.

### GetMondayStartTime

`func (o *Calendar) GetMondayStartTime() string`

GetMondayStartTime returns the MondayStartTime field if non-nil, zero value otherwise.

### GetMondayStartTimeOk

`func (o *Calendar) GetMondayStartTimeOk() (*string, bool)`

GetMondayStartTimeOk returns a tuple with the MondayStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMondayStartTime

`func (o *Calendar) SetMondayStartTime(v string)`

SetMondayStartTime sets MondayStartTime field to given value.

### HasMondayStartTime

`func (o *Calendar) HasMondayStartTime() bool`

HasMondayStartTime returns a boolean if a field has been set.

### GetMondayEndTime

`func (o *Calendar) GetMondayEndTime() string`

GetMondayEndTime returns the MondayEndTime field if non-nil, zero value otherwise.

### GetMondayEndTimeOk

`func (o *Calendar) GetMondayEndTimeOk() (*string, bool)`

GetMondayEndTimeOk returns a tuple with the MondayEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMondayEndTime

`func (o *Calendar) SetMondayEndTime(v string)`

SetMondayEndTime sets MondayEndTime field to given value.

### HasMondayEndTime

`func (o *Calendar) HasMondayEndTime() bool`

HasMondayEndTime returns a boolean if a field has been set.

### GetTuesdayStartTime

`func (o *Calendar) GetTuesdayStartTime() string`

GetTuesdayStartTime returns the TuesdayStartTime field if non-nil, zero value otherwise.

### GetTuesdayStartTimeOk

`func (o *Calendar) GetTuesdayStartTimeOk() (*string, bool)`

GetTuesdayStartTimeOk returns a tuple with the TuesdayStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesdayStartTime

`func (o *Calendar) SetTuesdayStartTime(v string)`

SetTuesdayStartTime sets TuesdayStartTime field to given value.

### HasTuesdayStartTime

`func (o *Calendar) HasTuesdayStartTime() bool`

HasTuesdayStartTime returns a boolean if a field has been set.

### GetTuesdayEndTime

`func (o *Calendar) GetTuesdayEndTime() string`

GetTuesdayEndTime returns the TuesdayEndTime field if non-nil, zero value otherwise.

### GetTuesdayEndTimeOk

`func (o *Calendar) GetTuesdayEndTimeOk() (*string, bool)`

GetTuesdayEndTimeOk returns a tuple with the TuesdayEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesdayEndTime

`func (o *Calendar) SetTuesdayEndTime(v string)`

SetTuesdayEndTime sets TuesdayEndTime field to given value.

### HasTuesdayEndTime

`func (o *Calendar) HasTuesdayEndTime() bool`

HasTuesdayEndTime returns a boolean if a field has been set.

### GetWednesdayStartTime

`func (o *Calendar) GetWednesdayStartTime() string`

GetWednesdayStartTime returns the WednesdayStartTime field if non-nil, zero value otherwise.

### GetWednesdayStartTimeOk

`func (o *Calendar) GetWednesdayStartTimeOk() (*string, bool)`

GetWednesdayStartTimeOk returns a tuple with the WednesdayStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesdayStartTime

`func (o *Calendar) SetWednesdayStartTime(v string)`

SetWednesdayStartTime sets WednesdayStartTime field to given value.

### HasWednesdayStartTime

`func (o *Calendar) HasWednesdayStartTime() bool`

HasWednesdayStartTime returns a boolean if a field has been set.

### GetWednesdayEndTime

`func (o *Calendar) GetWednesdayEndTime() string`

GetWednesdayEndTime returns the WednesdayEndTime field if non-nil, zero value otherwise.

### GetWednesdayEndTimeOk

`func (o *Calendar) GetWednesdayEndTimeOk() (*string, bool)`

GetWednesdayEndTimeOk returns a tuple with the WednesdayEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesdayEndTime

`func (o *Calendar) SetWednesdayEndTime(v string)`

SetWednesdayEndTime sets WednesdayEndTime field to given value.

### HasWednesdayEndTime

`func (o *Calendar) HasWednesdayEndTime() bool`

HasWednesdayEndTime returns a boolean if a field has been set.

### GetThursdayStartTime

`func (o *Calendar) GetThursdayStartTime() string`

GetThursdayStartTime returns the ThursdayStartTime field if non-nil, zero value otherwise.

### GetThursdayStartTimeOk

`func (o *Calendar) GetThursdayStartTimeOk() (*string, bool)`

GetThursdayStartTimeOk returns a tuple with the ThursdayStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursdayStartTime

`func (o *Calendar) SetThursdayStartTime(v string)`

SetThursdayStartTime sets ThursdayStartTime field to given value.

### HasThursdayStartTime

`func (o *Calendar) HasThursdayStartTime() bool`

HasThursdayStartTime returns a boolean if a field has been set.

### GetThursdayEndTime

`func (o *Calendar) GetThursdayEndTime() string`

GetThursdayEndTime returns the ThursdayEndTime field if non-nil, zero value otherwise.

### GetThursdayEndTimeOk

`func (o *Calendar) GetThursdayEndTimeOk() (*string, bool)`

GetThursdayEndTimeOk returns a tuple with the ThursdayEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursdayEndTime

`func (o *Calendar) SetThursdayEndTime(v string)`

SetThursdayEndTime sets ThursdayEndTime field to given value.

### HasThursdayEndTime

`func (o *Calendar) HasThursdayEndTime() bool`

HasThursdayEndTime returns a boolean if a field has been set.

### GetFridayStartTime

`func (o *Calendar) GetFridayStartTime() string`

GetFridayStartTime returns the FridayStartTime field if non-nil, zero value otherwise.

### GetFridayStartTimeOk

`func (o *Calendar) GetFridayStartTimeOk() (*string, bool)`

GetFridayStartTimeOk returns a tuple with the FridayStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFridayStartTime

`func (o *Calendar) SetFridayStartTime(v string)`

SetFridayStartTime sets FridayStartTime field to given value.

### HasFridayStartTime

`func (o *Calendar) HasFridayStartTime() bool`

HasFridayStartTime returns a boolean if a field has been set.

### GetFridayEndTime

`func (o *Calendar) GetFridayEndTime() string`

GetFridayEndTime returns the FridayEndTime field if non-nil, zero value otherwise.

### GetFridayEndTimeOk

`func (o *Calendar) GetFridayEndTimeOk() (*string, bool)`

GetFridayEndTimeOk returns a tuple with the FridayEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFridayEndTime

`func (o *Calendar) SetFridayEndTime(v string)`

SetFridayEndTime sets FridayEndTime field to given value.

### HasFridayEndTime

`func (o *Calendar) HasFridayEndTime() bool`

HasFridayEndTime returns a boolean if a field has been set.

### GetSaturdayStartTime

`func (o *Calendar) GetSaturdayStartTime() string`

GetSaturdayStartTime returns the SaturdayStartTime field if non-nil, zero value otherwise.

### GetSaturdayStartTimeOk

`func (o *Calendar) GetSaturdayStartTimeOk() (*string, bool)`

GetSaturdayStartTimeOk returns a tuple with the SaturdayStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturdayStartTime

`func (o *Calendar) SetSaturdayStartTime(v string)`

SetSaturdayStartTime sets SaturdayStartTime field to given value.

### HasSaturdayStartTime

`func (o *Calendar) HasSaturdayStartTime() bool`

HasSaturdayStartTime returns a boolean if a field has been set.

### GetSaturdayEndTime

`func (o *Calendar) GetSaturdayEndTime() string`

GetSaturdayEndTime returns the SaturdayEndTime field if non-nil, zero value otherwise.

### GetSaturdayEndTimeOk

`func (o *Calendar) GetSaturdayEndTimeOk() (*string, bool)`

GetSaturdayEndTimeOk returns a tuple with the SaturdayEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturdayEndTime

`func (o *Calendar) SetSaturdayEndTime(v string)`

SetSaturdayEndTime sets SaturdayEndTime field to given value.

### HasSaturdayEndTime

`func (o *Calendar) HasSaturdayEndTime() bool`

HasSaturdayEndTime returns a boolean if a field has been set.

### GetSundayStartTime

`func (o *Calendar) GetSundayStartTime() string`

GetSundayStartTime returns the SundayStartTime field if non-nil, zero value otherwise.

### GetSundayStartTimeOk

`func (o *Calendar) GetSundayStartTimeOk() (*string, bool)`

GetSundayStartTimeOk returns a tuple with the SundayStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSundayStartTime

`func (o *Calendar) SetSundayStartTime(v string)`

SetSundayStartTime sets SundayStartTime field to given value.

### HasSundayStartTime

`func (o *Calendar) HasSundayStartTime() bool`

HasSundayStartTime returns a boolean if a field has been set.

### GetSundayEndTime

`func (o *Calendar) GetSundayEndTime() string`

GetSundayEndTime returns the SundayEndTime field if non-nil, zero value otherwise.

### GetSundayEndTimeOk

`func (o *Calendar) GetSundayEndTimeOk() (*string, bool)`

GetSundayEndTimeOk returns a tuple with the SundayEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSundayEndTime

`func (o *Calendar) SetSundayEndTime(v string)`

SetSundayEndTime sets SundayEndTime field to given value.

### HasSundayEndTime

`func (o *Calendar) HasSundayEndTime() bool`

HasSundayEndTime returns a boolean if a field has been set.

### GetInfo

`func (o *Calendar) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Calendar) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Calendar) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Calendar) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


