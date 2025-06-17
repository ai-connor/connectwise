# TimeZoneSetup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**TimeZone** | [**TimeZoneReference**](TimeZoneReference.md) |  | 
**Offset** | Pointer to **NullableFloat64** | The hours offset from UTC (+/-) | [optional] 
**DefaultFlag** | Pointer to **NullableBool** | Identifies the default system time zone setup | [optional] 
**DaylightSavingsFlag** | Pointer to **NullableBool** | Determined based on system library value for specified timeZone.             Not able to be used in query params at this time | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTimeZoneSetup

`func NewTimeZoneSetup(name string, timeZone TimeZoneReference, ) *TimeZoneSetup`

NewTimeZoneSetup instantiates a new TimeZoneSetup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeZoneSetupWithDefaults

`func NewTimeZoneSetupWithDefaults() *TimeZoneSetup`

NewTimeZoneSetupWithDefaults instantiates a new TimeZoneSetup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimeZoneSetup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimeZoneSetup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimeZoneSetup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TimeZoneSetup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TimeZoneSetup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TimeZoneSetup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TimeZoneSetup) SetName(v string)`

SetName sets Name field to given value.


### GetTimeZone

`func (o *TimeZoneSetup) GetTimeZone() TimeZoneReference`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *TimeZoneSetup) GetTimeZoneOk() (*TimeZoneReference, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *TimeZoneSetup) SetTimeZone(v TimeZoneReference)`

SetTimeZone sets TimeZone field to given value.


### GetOffset

`func (o *TimeZoneSetup) GetOffset() float64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TimeZoneSetup) GetOffsetOk() (*float64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TimeZoneSetup) SetOffset(v float64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TimeZoneSetup) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *TimeZoneSetup) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *TimeZoneSetup) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetDefaultFlag

`func (o *TimeZoneSetup) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *TimeZoneSetup) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *TimeZoneSetup) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *TimeZoneSetup) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *TimeZoneSetup) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *TimeZoneSetup) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetDaylightSavingsFlag

`func (o *TimeZoneSetup) GetDaylightSavingsFlag() bool`

GetDaylightSavingsFlag returns the DaylightSavingsFlag field if non-nil, zero value otherwise.

### GetDaylightSavingsFlagOk

`func (o *TimeZoneSetup) GetDaylightSavingsFlagOk() (*bool, bool)`

GetDaylightSavingsFlagOk returns a tuple with the DaylightSavingsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaylightSavingsFlag

`func (o *TimeZoneSetup) SetDaylightSavingsFlag(v bool)`

SetDaylightSavingsFlag sets DaylightSavingsFlag field to given value.

### HasDaylightSavingsFlag

`func (o *TimeZoneSetup) HasDaylightSavingsFlag() bool`

HasDaylightSavingsFlag returns a boolean if a field has been set.

### SetDaylightSavingsFlagNil

`func (o *TimeZoneSetup) SetDaylightSavingsFlagNil(b bool)`

 SetDaylightSavingsFlagNil sets the value for DaylightSavingsFlag to be an explicit nil

### UnsetDaylightSavingsFlag
`func (o *TimeZoneSetup) UnsetDaylightSavingsFlag()`

UnsetDaylightSavingsFlag ensures that no value is present for DaylightSavingsFlag, not even an explicit nil
### GetInfo

`func (o *TimeZoneSetup) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TimeZoneSetup) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TimeZoneSetup) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TimeZoneSetup) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


