# ScheduleEntryDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ScheduleEntry** | Pointer to [**ScheduleEntryReference**](ScheduleEntryReference.md) |  | [optional] 
**DateStart** | Pointer to **string** |  | [optional] 
**DateEnd** | Pointer to **string** |  | [optional] 
**HoursScheduled** | Pointer to **float64** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewScheduleEntryDetail

`func NewScheduleEntryDetail() *ScheduleEntryDetail`

NewScheduleEntryDetail instantiates a new ScheduleEntryDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleEntryDetailWithDefaults

`func NewScheduleEntryDetailWithDefaults() *ScheduleEntryDetail`

NewScheduleEntryDetailWithDefaults instantiates a new ScheduleEntryDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduleEntryDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleEntryDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleEntryDetail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ScheduleEntryDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScheduleEntry

`func (o *ScheduleEntryDetail) GetScheduleEntry() ScheduleEntryReference`

GetScheduleEntry returns the ScheduleEntry field if non-nil, zero value otherwise.

### GetScheduleEntryOk

`func (o *ScheduleEntryDetail) GetScheduleEntryOk() (*ScheduleEntryReference, bool)`

GetScheduleEntryOk returns a tuple with the ScheduleEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEntry

`func (o *ScheduleEntryDetail) SetScheduleEntry(v ScheduleEntryReference)`

SetScheduleEntry sets ScheduleEntry field to given value.

### HasScheduleEntry

`func (o *ScheduleEntryDetail) HasScheduleEntry() bool`

HasScheduleEntry returns a boolean if a field has been set.

### GetDateStart

`func (o *ScheduleEntryDetail) GetDateStart() string`

GetDateStart returns the DateStart field if non-nil, zero value otherwise.

### GetDateStartOk

`func (o *ScheduleEntryDetail) GetDateStartOk() (*string, bool)`

GetDateStartOk returns a tuple with the DateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStart

`func (o *ScheduleEntryDetail) SetDateStart(v string)`

SetDateStart sets DateStart field to given value.

### HasDateStart

`func (o *ScheduleEntryDetail) HasDateStart() bool`

HasDateStart returns a boolean if a field has been set.

### GetDateEnd

`func (o *ScheduleEntryDetail) GetDateEnd() string`

GetDateEnd returns the DateEnd field if non-nil, zero value otherwise.

### GetDateEndOk

`func (o *ScheduleEntryDetail) GetDateEndOk() (*string, bool)`

GetDateEndOk returns a tuple with the DateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnd

`func (o *ScheduleEntryDetail) SetDateEnd(v string)`

SetDateEnd sets DateEnd field to given value.

### HasDateEnd

`func (o *ScheduleEntryDetail) HasDateEnd() bool`

HasDateEnd returns a boolean if a field has been set.

### GetHoursScheduled

`func (o *ScheduleEntryDetail) GetHoursScheduled() float64`

GetHoursScheduled returns the HoursScheduled field if non-nil, zero value otherwise.

### GetHoursScheduledOk

`func (o *ScheduleEntryDetail) GetHoursScheduledOk() (*float64, bool)`

GetHoursScheduledOk returns a tuple with the HoursScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursScheduled

`func (o *ScheduleEntryDetail) SetHoursScheduled(v float64)`

SetHoursScheduled sets HoursScheduled field to given value.

### HasHoursScheduled

`func (o *ScheduleEntryDetail) HasHoursScheduled() bool`

HasHoursScheduled returns a boolean if a field has been set.

### GetInfo

`func (o *ScheduleEntryDetail) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ScheduleEntryDetail) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ScheduleEntryDetail) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ScheduleEntryDetail) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


