# ScheduleReminderTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Time** | Pointer to **NullableInt32** | Time is calculated in minutes. | [optional] 
**Name** | Pointer to **string** |  Max length: 10; | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewScheduleReminderTime

`func NewScheduleReminderTime() *ScheduleReminderTime`

NewScheduleReminderTime instantiates a new ScheduleReminderTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleReminderTimeWithDefaults

`func NewScheduleReminderTimeWithDefaults() *ScheduleReminderTime`

NewScheduleReminderTimeWithDefaults instantiates a new ScheduleReminderTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduleReminderTime) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleReminderTime) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleReminderTime) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ScheduleReminderTime) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTime

`func (o *ScheduleReminderTime) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ScheduleReminderTime) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ScheduleReminderTime) SetTime(v int32)`

SetTime sets Time field to given value.

### HasTime

`func (o *ScheduleReminderTime) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTimeNil

`func (o *ScheduleReminderTime) SetTimeNil(b bool)`

 SetTimeNil sets the value for Time to be an explicit nil

### UnsetTime
`func (o *ScheduleReminderTime) UnsetTime()`

UnsetTime ensures that no value is present for Time, not even an explicit nil
### GetName

`func (o *ScheduleReminderTime) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduleReminderTime) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduleReminderTime) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScheduleReminderTime) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultFlag

`func (o *ScheduleReminderTime) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *ScheduleReminderTime) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *ScheduleReminderTime) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *ScheduleReminderTime) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *ScheduleReminderTime) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *ScheduleReminderTime) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetInfo

`func (o *ScheduleReminderTime) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ScheduleReminderTime) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ScheduleReminderTime) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ScheduleReminderTime) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


