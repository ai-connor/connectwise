# PortalCalendar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**WeekStart** | **NullableString** |  | 
**Adjust1Start** | Pointer to **string** |  | [optional] 
**Adjust1End** | Pointer to **string** |  | [optional] 
**Adjust1Hours** | Pointer to **NullableFloat64** |  | [optional] 
**Adjust2Start** | Pointer to **string** |  | [optional] 
**Adjust2End** | Pointer to **string** |  | [optional] 
**Adjust2Hours** | Pointer to **NullableFloat64** |  | [optional] 
**Adjust3Start** | Pointer to **string** |  | [optional] 
**Adjust3End** | Pointer to **string** |  | [optional] 
**Adjust3Hours** | Pointer to **NullableFloat64** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPortalCalendar

`func NewPortalCalendar(weekStart NullableString, ) *PortalCalendar`

NewPortalCalendar instantiates a new PortalCalendar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalCalendarWithDefaults

`func NewPortalCalendarWithDefaults() *PortalCalendar`

NewPortalCalendarWithDefaults instantiates a new PortalCalendar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortalCalendar) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortalCalendar) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortalCalendar) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PortalCalendar) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWeekStart

`func (o *PortalCalendar) GetWeekStart() string`

GetWeekStart returns the WeekStart field if non-nil, zero value otherwise.

### GetWeekStartOk

`func (o *PortalCalendar) GetWeekStartOk() (*string, bool)`

GetWeekStartOk returns a tuple with the WeekStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekStart

`func (o *PortalCalendar) SetWeekStart(v string)`

SetWeekStart sets WeekStart field to given value.


### SetWeekStartNil

`func (o *PortalCalendar) SetWeekStartNil(b bool)`

 SetWeekStartNil sets the value for WeekStart to be an explicit nil

### UnsetWeekStart
`func (o *PortalCalendar) UnsetWeekStart()`

UnsetWeekStart ensures that no value is present for WeekStart, not even an explicit nil
### GetAdjust1Start

`func (o *PortalCalendar) GetAdjust1Start() string`

GetAdjust1Start returns the Adjust1Start field if non-nil, zero value otherwise.

### GetAdjust1StartOk

`func (o *PortalCalendar) GetAdjust1StartOk() (*string, bool)`

GetAdjust1StartOk returns a tuple with the Adjust1Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjust1Start

`func (o *PortalCalendar) SetAdjust1Start(v string)`

SetAdjust1Start sets Adjust1Start field to given value.

### HasAdjust1Start

`func (o *PortalCalendar) HasAdjust1Start() bool`

HasAdjust1Start returns a boolean if a field has been set.

### GetAdjust1End

`func (o *PortalCalendar) GetAdjust1End() string`

GetAdjust1End returns the Adjust1End field if non-nil, zero value otherwise.

### GetAdjust1EndOk

`func (o *PortalCalendar) GetAdjust1EndOk() (*string, bool)`

GetAdjust1EndOk returns a tuple with the Adjust1End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjust1End

`func (o *PortalCalendar) SetAdjust1End(v string)`

SetAdjust1End sets Adjust1End field to given value.

### HasAdjust1End

`func (o *PortalCalendar) HasAdjust1End() bool`

HasAdjust1End returns a boolean if a field has been set.

### GetAdjust1Hours

`func (o *PortalCalendar) GetAdjust1Hours() float64`

GetAdjust1Hours returns the Adjust1Hours field if non-nil, zero value otherwise.

### GetAdjust1HoursOk

`func (o *PortalCalendar) GetAdjust1HoursOk() (*float64, bool)`

GetAdjust1HoursOk returns a tuple with the Adjust1Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjust1Hours

`func (o *PortalCalendar) SetAdjust1Hours(v float64)`

SetAdjust1Hours sets Adjust1Hours field to given value.

### HasAdjust1Hours

`func (o *PortalCalendar) HasAdjust1Hours() bool`

HasAdjust1Hours returns a boolean if a field has been set.

### SetAdjust1HoursNil

`func (o *PortalCalendar) SetAdjust1HoursNil(b bool)`

 SetAdjust1HoursNil sets the value for Adjust1Hours to be an explicit nil

### UnsetAdjust1Hours
`func (o *PortalCalendar) UnsetAdjust1Hours()`

UnsetAdjust1Hours ensures that no value is present for Adjust1Hours, not even an explicit nil
### GetAdjust2Start

`func (o *PortalCalendar) GetAdjust2Start() string`

GetAdjust2Start returns the Adjust2Start field if non-nil, zero value otherwise.

### GetAdjust2StartOk

`func (o *PortalCalendar) GetAdjust2StartOk() (*string, bool)`

GetAdjust2StartOk returns a tuple with the Adjust2Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjust2Start

`func (o *PortalCalendar) SetAdjust2Start(v string)`

SetAdjust2Start sets Adjust2Start field to given value.

### HasAdjust2Start

`func (o *PortalCalendar) HasAdjust2Start() bool`

HasAdjust2Start returns a boolean if a field has been set.

### GetAdjust2End

`func (o *PortalCalendar) GetAdjust2End() string`

GetAdjust2End returns the Adjust2End field if non-nil, zero value otherwise.

### GetAdjust2EndOk

`func (o *PortalCalendar) GetAdjust2EndOk() (*string, bool)`

GetAdjust2EndOk returns a tuple with the Adjust2End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjust2End

`func (o *PortalCalendar) SetAdjust2End(v string)`

SetAdjust2End sets Adjust2End field to given value.

### HasAdjust2End

`func (o *PortalCalendar) HasAdjust2End() bool`

HasAdjust2End returns a boolean if a field has been set.

### GetAdjust2Hours

`func (o *PortalCalendar) GetAdjust2Hours() float64`

GetAdjust2Hours returns the Adjust2Hours field if non-nil, zero value otherwise.

### GetAdjust2HoursOk

`func (o *PortalCalendar) GetAdjust2HoursOk() (*float64, bool)`

GetAdjust2HoursOk returns a tuple with the Adjust2Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjust2Hours

`func (o *PortalCalendar) SetAdjust2Hours(v float64)`

SetAdjust2Hours sets Adjust2Hours field to given value.

### HasAdjust2Hours

`func (o *PortalCalendar) HasAdjust2Hours() bool`

HasAdjust2Hours returns a boolean if a field has been set.

### SetAdjust2HoursNil

`func (o *PortalCalendar) SetAdjust2HoursNil(b bool)`

 SetAdjust2HoursNil sets the value for Adjust2Hours to be an explicit nil

### UnsetAdjust2Hours
`func (o *PortalCalendar) UnsetAdjust2Hours()`

UnsetAdjust2Hours ensures that no value is present for Adjust2Hours, not even an explicit nil
### GetAdjust3Start

`func (o *PortalCalendar) GetAdjust3Start() string`

GetAdjust3Start returns the Adjust3Start field if non-nil, zero value otherwise.

### GetAdjust3StartOk

`func (o *PortalCalendar) GetAdjust3StartOk() (*string, bool)`

GetAdjust3StartOk returns a tuple with the Adjust3Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjust3Start

`func (o *PortalCalendar) SetAdjust3Start(v string)`

SetAdjust3Start sets Adjust3Start field to given value.

### HasAdjust3Start

`func (o *PortalCalendar) HasAdjust3Start() bool`

HasAdjust3Start returns a boolean if a field has been set.

### GetAdjust3End

`func (o *PortalCalendar) GetAdjust3End() string`

GetAdjust3End returns the Adjust3End field if non-nil, zero value otherwise.

### GetAdjust3EndOk

`func (o *PortalCalendar) GetAdjust3EndOk() (*string, bool)`

GetAdjust3EndOk returns a tuple with the Adjust3End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjust3End

`func (o *PortalCalendar) SetAdjust3End(v string)`

SetAdjust3End sets Adjust3End field to given value.

### HasAdjust3End

`func (o *PortalCalendar) HasAdjust3End() bool`

HasAdjust3End returns a boolean if a field has been set.

### GetAdjust3Hours

`func (o *PortalCalendar) GetAdjust3Hours() float64`

GetAdjust3Hours returns the Adjust3Hours field if non-nil, zero value otherwise.

### GetAdjust3HoursOk

`func (o *PortalCalendar) GetAdjust3HoursOk() (*float64, bool)`

GetAdjust3HoursOk returns a tuple with the Adjust3Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjust3Hours

`func (o *PortalCalendar) SetAdjust3Hours(v float64)`

SetAdjust3Hours sets Adjust3Hours field to given value.

### HasAdjust3Hours

`func (o *PortalCalendar) HasAdjust3Hours() bool`

HasAdjust3Hours returns a boolean if a field has been set.

### SetAdjust3HoursNil

`func (o *PortalCalendar) SetAdjust3HoursNil(b bool)`

 SetAdjust3HoursNil sets the value for Adjust3Hours to be an explicit nil

### UnsetAdjust3Hours
`func (o *PortalCalendar) UnsetAdjust3Hours()`

UnsetAdjust3Hours ensures that no value is present for Adjust3Hours, not even an explicit nil
### GetInfo

`func (o *PortalCalendar) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PortalCalendar) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PortalCalendar) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PortalCalendar) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


