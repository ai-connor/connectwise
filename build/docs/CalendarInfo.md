# CalendarInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
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

### NewCalendarInfo

`func NewCalendarInfo() *CalendarInfo`

NewCalendarInfo instantiates a new CalendarInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarInfoWithDefaults

`func NewCalendarInfoWithDefaults() *CalendarInfo`

NewCalendarInfoWithDefaults instantiates a new CalendarInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CalendarInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CalendarInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CalendarInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CalendarInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CalendarInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CalendarInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CalendarInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CalendarInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHolidayList

`func (o *CalendarInfo) GetHolidayList() HolidayListReference`

GetHolidayList returns the HolidayList field if non-nil, zero value otherwise.

### GetHolidayListOk

`func (o *CalendarInfo) GetHolidayListOk() (*HolidayListReference, bool)`

GetHolidayListOk returns a tuple with the HolidayList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidayList

`func (o *CalendarInfo) SetHolidayList(v HolidayListReference)`

SetHolidayList sets HolidayList field to given value.

### HasHolidayList

`func (o *CalendarInfo) HasHolidayList() bool`

HasHolidayList returns a boolean if a field has been set.

### GetMondayStartTime

`func (o *CalendarInfo) GetMondayStartTime() string`

GetMondayStartTime returns the MondayStartTime field if non-nil, zero value otherwise.

### GetMondayStartTimeOk

`func (o *CalendarInfo) GetMondayStartTimeOk() (*string, bool)`

GetMondayStartTimeOk returns a tuple with the MondayStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMondayStartTime

`func (o *CalendarInfo) SetMondayStartTime(v string)`

SetMondayStartTime sets MondayStartTime field to given value.

### HasMondayStartTime

`func (o *CalendarInfo) HasMondayStartTime() bool`

HasMondayStartTime returns a boolean if a field has been set.

### GetMondayEndTime

`func (o *CalendarInfo) GetMondayEndTime() string`

GetMondayEndTime returns the MondayEndTime field if non-nil, zero value otherwise.

### GetMondayEndTimeOk

`func (o *CalendarInfo) GetMondayEndTimeOk() (*string, bool)`

GetMondayEndTimeOk returns a tuple with the MondayEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMondayEndTime

`func (o *CalendarInfo) SetMondayEndTime(v string)`

SetMondayEndTime sets MondayEndTime field to given value.

### HasMondayEndTime

`func (o *CalendarInfo) HasMondayEndTime() bool`

HasMondayEndTime returns a boolean if a field has been set.

### GetTuesdayStartTime

`func (o *CalendarInfo) GetTuesdayStartTime() string`

GetTuesdayStartTime returns the TuesdayStartTime field if non-nil, zero value otherwise.

### GetTuesdayStartTimeOk

`func (o *CalendarInfo) GetTuesdayStartTimeOk() (*string, bool)`

GetTuesdayStartTimeOk returns a tuple with the TuesdayStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesdayStartTime

`func (o *CalendarInfo) SetTuesdayStartTime(v string)`

SetTuesdayStartTime sets TuesdayStartTime field to given value.

### HasTuesdayStartTime

`func (o *CalendarInfo) HasTuesdayStartTime() bool`

HasTuesdayStartTime returns a boolean if a field has been set.

### GetTuesdayEndTime

`func (o *CalendarInfo) GetTuesdayEndTime() string`

GetTuesdayEndTime returns the TuesdayEndTime field if non-nil, zero value otherwise.

### GetTuesdayEndTimeOk

`func (o *CalendarInfo) GetTuesdayEndTimeOk() (*string, bool)`

GetTuesdayEndTimeOk returns a tuple with the TuesdayEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesdayEndTime

`func (o *CalendarInfo) SetTuesdayEndTime(v string)`

SetTuesdayEndTime sets TuesdayEndTime field to given value.

### HasTuesdayEndTime

`func (o *CalendarInfo) HasTuesdayEndTime() bool`

HasTuesdayEndTime returns a boolean if a field has been set.

### GetWednesdayStartTime

`func (o *CalendarInfo) GetWednesdayStartTime() string`

GetWednesdayStartTime returns the WednesdayStartTime field if non-nil, zero value otherwise.

### GetWednesdayStartTimeOk

`func (o *CalendarInfo) GetWednesdayStartTimeOk() (*string, bool)`

GetWednesdayStartTimeOk returns a tuple with the WednesdayStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesdayStartTime

`func (o *CalendarInfo) SetWednesdayStartTime(v string)`

SetWednesdayStartTime sets WednesdayStartTime field to given value.

### HasWednesdayStartTime

`func (o *CalendarInfo) HasWednesdayStartTime() bool`

HasWednesdayStartTime returns a boolean if a field has been set.

### GetWednesdayEndTime

`func (o *CalendarInfo) GetWednesdayEndTime() string`

GetWednesdayEndTime returns the WednesdayEndTime field if non-nil, zero value otherwise.

### GetWednesdayEndTimeOk

`func (o *CalendarInfo) GetWednesdayEndTimeOk() (*string, bool)`

GetWednesdayEndTimeOk returns a tuple with the WednesdayEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesdayEndTime

`func (o *CalendarInfo) SetWednesdayEndTime(v string)`

SetWednesdayEndTime sets WednesdayEndTime field to given value.

### HasWednesdayEndTime

`func (o *CalendarInfo) HasWednesdayEndTime() bool`

HasWednesdayEndTime returns a boolean if a field has been set.

### GetThursdayStartTime

`func (o *CalendarInfo) GetThursdayStartTime() string`

GetThursdayStartTime returns the ThursdayStartTime field if non-nil, zero value otherwise.

### GetThursdayStartTimeOk

`func (o *CalendarInfo) GetThursdayStartTimeOk() (*string, bool)`

GetThursdayStartTimeOk returns a tuple with the ThursdayStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursdayStartTime

`func (o *CalendarInfo) SetThursdayStartTime(v string)`

SetThursdayStartTime sets ThursdayStartTime field to given value.

### HasThursdayStartTime

`func (o *CalendarInfo) HasThursdayStartTime() bool`

HasThursdayStartTime returns a boolean if a field has been set.

### GetThursdayEndTime

`func (o *CalendarInfo) GetThursdayEndTime() string`

GetThursdayEndTime returns the ThursdayEndTime field if non-nil, zero value otherwise.

### GetThursdayEndTimeOk

`func (o *CalendarInfo) GetThursdayEndTimeOk() (*string, bool)`

GetThursdayEndTimeOk returns a tuple with the ThursdayEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursdayEndTime

`func (o *CalendarInfo) SetThursdayEndTime(v string)`

SetThursdayEndTime sets ThursdayEndTime field to given value.

### HasThursdayEndTime

`func (o *CalendarInfo) HasThursdayEndTime() bool`

HasThursdayEndTime returns a boolean if a field has been set.

### GetFridayStartTime

`func (o *CalendarInfo) GetFridayStartTime() string`

GetFridayStartTime returns the FridayStartTime field if non-nil, zero value otherwise.

### GetFridayStartTimeOk

`func (o *CalendarInfo) GetFridayStartTimeOk() (*string, bool)`

GetFridayStartTimeOk returns a tuple with the FridayStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFridayStartTime

`func (o *CalendarInfo) SetFridayStartTime(v string)`

SetFridayStartTime sets FridayStartTime field to given value.

### HasFridayStartTime

`func (o *CalendarInfo) HasFridayStartTime() bool`

HasFridayStartTime returns a boolean if a field has been set.

### GetFridayEndTime

`func (o *CalendarInfo) GetFridayEndTime() string`

GetFridayEndTime returns the FridayEndTime field if non-nil, zero value otherwise.

### GetFridayEndTimeOk

`func (o *CalendarInfo) GetFridayEndTimeOk() (*string, bool)`

GetFridayEndTimeOk returns a tuple with the FridayEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFridayEndTime

`func (o *CalendarInfo) SetFridayEndTime(v string)`

SetFridayEndTime sets FridayEndTime field to given value.

### HasFridayEndTime

`func (o *CalendarInfo) HasFridayEndTime() bool`

HasFridayEndTime returns a boolean if a field has been set.

### GetSaturdayStartTime

`func (o *CalendarInfo) GetSaturdayStartTime() string`

GetSaturdayStartTime returns the SaturdayStartTime field if non-nil, zero value otherwise.

### GetSaturdayStartTimeOk

`func (o *CalendarInfo) GetSaturdayStartTimeOk() (*string, bool)`

GetSaturdayStartTimeOk returns a tuple with the SaturdayStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturdayStartTime

`func (o *CalendarInfo) SetSaturdayStartTime(v string)`

SetSaturdayStartTime sets SaturdayStartTime field to given value.

### HasSaturdayStartTime

`func (o *CalendarInfo) HasSaturdayStartTime() bool`

HasSaturdayStartTime returns a boolean if a field has been set.

### GetSaturdayEndTime

`func (o *CalendarInfo) GetSaturdayEndTime() string`

GetSaturdayEndTime returns the SaturdayEndTime field if non-nil, zero value otherwise.

### GetSaturdayEndTimeOk

`func (o *CalendarInfo) GetSaturdayEndTimeOk() (*string, bool)`

GetSaturdayEndTimeOk returns a tuple with the SaturdayEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturdayEndTime

`func (o *CalendarInfo) SetSaturdayEndTime(v string)`

SetSaturdayEndTime sets SaturdayEndTime field to given value.

### HasSaturdayEndTime

`func (o *CalendarInfo) HasSaturdayEndTime() bool`

HasSaturdayEndTime returns a boolean if a field has been set.

### GetSundayStartTime

`func (o *CalendarInfo) GetSundayStartTime() string`

GetSundayStartTime returns the SundayStartTime field if non-nil, zero value otherwise.

### GetSundayStartTimeOk

`func (o *CalendarInfo) GetSundayStartTimeOk() (*string, bool)`

GetSundayStartTimeOk returns a tuple with the SundayStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSundayStartTime

`func (o *CalendarInfo) SetSundayStartTime(v string)`

SetSundayStartTime sets SundayStartTime field to given value.

### HasSundayStartTime

`func (o *CalendarInfo) HasSundayStartTime() bool`

HasSundayStartTime returns a boolean if a field has been set.

### GetSundayEndTime

`func (o *CalendarInfo) GetSundayEndTime() string`

GetSundayEndTime returns the SundayEndTime field if non-nil, zero value otherwise.

### GetSundayEndTimeOk

`func (o *CalendarInfo) GetSundayEndTimeOk() (*string, bool)`

GetSundayEndTimeOk returns a tuple with the SundayEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSundayEndTime

`func (o *CalendarInfo) SetSundayEndTime(v string)`

SetSundayEndTime sets SundayEndTime field to given value.

### HasSundayEndTime

`func (o *CalendarInfo) HasSundayEndTime() bool`

HasSundayEndTime returns a boolean if a field has been set.

### GetInfo

`func (o *CalendarInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CalendarInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CalendarInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CalendarInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


