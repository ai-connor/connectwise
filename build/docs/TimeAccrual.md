# TimeAccrual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**VacationFlag** | Pointer to **NullableBool** | if vacationFlag is set to false, system will clear out or ingore the values of vacationAvailableType, vacationCarryoverAllowedFlag, vacationCarryoverLimit | [optional] 
**VacationAvailableType** | Pointer to **NullableString** |  | [optional] 
**VacationCarryoverAllowedFlag** | Pointer to **NullableBool** |  | [optional] 
**VacationCarryoverLimit** | Pointer to **NullableFloat64** |  | [optional] 
**SickFlag** | Pointer to **NullableBool** | if sickFlag is set to false, system will clear out or ignore the values of sickAvailableType, sickCarryoverAllowedFlag, sickCarryoverLimit | [optional] 
**SickAvailableType** | Pointer to **NullableString** |  | [optional] 
**SickCarryoverAllowedFlag** | Pointer to **NullableBool** |  | [optional] 
**SickCarryoverLimit** | Pointer to **NullableFloat64** |  | [optional] 
**PtoFlag** | Pointer to **NullableBool** | if ptoFlag is set to false, system will clear out or ignore the values of ptoAvailableType, ptoCarryoverAllowedFlag, ptoCarryoverLimit | [optional] 
**PtoAvailableType** | Pointer to **NullableString** |  | [optional] 
**PtoCarryoverAllowedFlag** | Pointer to **NullableBool** |  | [optional] 
**PtoCarryoverLimit** | Pointer to **NullableFloat64** |  | [optional] 
**HolidayFlag** | Pointer to **NullableBool** | if holidayFlag is set to false, system will clear out or ignore the values of holidayAvailableType, holidayCarryoverAllowedFlag, holidayCarryoverLimit | [optional] 
**HolidayAvailableType** | Pointer to **NullableString** |  | [optional] 
**HolidayCarryoverAllowedFlag** | Pointer to **NullableBool** |  | [optional] 
**HolidayCarryoverLimit** | Pointer to **NullableFloat64** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTimeAccrual

`func NewTimeAccrual() *TimeAccrual`

NewTimeAccrual instantiates a new TimeAccrual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeAccrualWithDefaults

`func NewTimeAccrualWithDefaults() *TimeAccrual`

NewTimeAccrualWithDefaults instantiates a new TimeAccrual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimeAccrual) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimeAccrual) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimeAccrual) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TimeAccrual) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *TimeAccrual) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *TimeAccrual) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *TimeAccrual) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *TimeAccrual) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetVacationFlag

`func (o *TimeAccrual) GetVacationFlag() bool`

GetVacationFlag returns the VacationFlag field if non-nil, zero value otherwise.

### GetVacationFlagOk

`func (o *TimeAccrual) GetVacationFlagOk() (*bool, bool)`

GetVacationFlagOk returns a tuple with the VacationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacationFlag

`func (o *TimeAccrual) SetVacationFlag(v bool)`

SetVacationFlag sets VacationFlag field to given value.

### HasVacationFlag

`func (o *TimeAccrual) HasVacationFlag() bool`

HasVacationFlag returns a boolean if a field has been set.

### SetVacationFlagNil

`func (o *TimeAccrual) SetVacationFlagNil(b bool)`

 SetVacationFlagNil sets the value for VacationFlag to be an explicit nil

### UnsetVacationFlag
`func (o *TimeAccrual) UnsetVacationFlag()`

UnsetVacationFlag ensures that no value is present for VacationFlag, not even an explicit nil
### GetVacationAvailableType

`func (o *TimeAccrual) GetVacationAvailableType() string`

GetVacationAvailableType returns the VacationAvailableType field if non-nil, zero value otherwise.

### GetVacationAvailableTypeOk

`func (o *TimeAccrual) GetVacationAvailableTypeOk() (*string, bool)`

GetVacationAvailableTypeOk returns a tuple with the VacationAvailableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacationAvailableType

`func (o *TimeAccrual) SetVacationAvailableType(v string)`

SetVacationAvailableType sets VacationAvailableType field to given value.

### HasVacationAvailableType

`func (o *TimeAccrual) HasVacationAvailableType() bool`

HasVacationAvailableType returns a boolean if a field has been set.

### SetVacationAvailableTypeNil

`func (o *TimeAccrual) SetVacationAvailableTypeNil(b bool)`

 SetVacationAvailableTypeNil sets the value for VacationAvailableType to be an explicit nil

### UnsetVacationAvailableType
`func (o *TimeAccrual) UnsetVacationAvailableType()`

UnsetVacationAvailableType ensures that no value is present for VacationAvailableType, not even an explicit nil
### GetVacationCarryoverAllowedFlag

`func (o *TimeAccrual) GetVacationCarryoverAllowedFlag() bool`

GetVacationCarryoverAllowedFlag returns the VacationCarryoverAllowedFlag field if non-nil, zero value otherwise.

### GetVacationCarryoverAllowedFlagOk

`func (o *TimeAccrual) GetVacationCarryoverAllowedFlagOk() (*bool, bool)`

GetVacationCarryoverAllowedFlagOk returns a tuple with the VacationCarryoverAllowedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacationCarryoverAllowedFlag

`func (o *TimeAccrual) SetVacationCarryoverAllowedFlag(v bool)`

SetVacationCarryoverAllowedFlag sets VacationCarryoverAllowedFlag field to given value.

### HasVacationCarryoverAllowedFlag

`func (o *TimeAccrual) HasVacationCarryoverAllowedFlag() bool`

HasVacationCarryoverAllowedFlag returns a boolean if a field has been set.

### SetVacationCarryoverAllowedFlagNil

`func (o *TimeAccrual) SetVacationCarryoverAllowedFlagNil(b bool)`

 SetVacationCarryoverAllowedFlagNil sets the value for VacationCarryoverAllowedFlag to be an explicit nil

### UnsetVacationCarryoverAllowedFlag
`func (o *TimeAccrual) UnsetVacationCarryoverAllowedFlag()`

UnsetVacationCarryoverAllowedFlag ensures that no value is present for VacationCarryoverAllowedFlag, not even an explicit nil
### GetVacationCarryoverLimit

`func (o *TimeAccrual) GetVacationCarryoverLimit() float64`

GetVacationCarryoverLimit returns the VacationCarryoverLimit field if non-nil, zero value otherwise.

### GetVacationCarryoverLimitOk

`func (o *TimeAccrual) GetVacationCarryoverLimitOk() (*float64, bool)`

GetVacationCarryoverLimitOk returns a tuple with the VacationCarryoverLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacationCarryoverLimit

`func (o *TimeAccrual) SetVacationCarryoverLimit(v float64)`

SetVacationCarryoverLimit sets VacationCarryoverLimit field to given value.

### HasVacationCarryoverLimit

`func (o *TimeAccrual) HasVacationCarryoverLimit() bool`

HasVacationCarryoverLimit returns a boolean if a field has been set.

### SetVacationCarryoverLimitNil

`func (o *TimeAccrual) SetVacationCarryoverLimitNil(b bool)`

 SetVacationCarryoverLimitNil sets the value for VacationCarryoverLimit to be an explicit nil

### UnsetVacationCarryoverLimit
`func (o *TimeAccrual) UnsetVacationCarryoverLimit()`

UnsetVacationCarryoverLimit ensures that no value is present for VacationCarryoverLimit, not even an explicit nil
### GetSickFlag

`func (o *TimeAccrual) GetSickFlag() bool`

GetSickFlag returns the SickFlag field if non-nil, zero value otherwise.

### GetSickFlagOk

`func (o *TimeAccrual) GetSickFlagOk() (*bool, bool)`

GetSickFlagOk returns a tuple with the SickFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSickFlag

`func (o *TimeAccrual) SetSickFlag(v bool)`

SetSickFlag sets SickFlag field to given value.

### HasSickFlag

`func (o *TimeAccrual) HasSickFlag() bool`

HasSickFlag returns a boolean if a field has been set.

### SetSickFlagNil

`func (o *TimeAccrual) SetSickFlagNil(b bool)`

 SetSickFlagNil sets the value for SickFlag to be an explicit nil

### UnsetSickFlag
`func (o *TimeAccrual) UnsetSickFlag()`

UnsetSickFlag ensures that no value is present for SickFlag, not even an explicit nil
### GetSickAvailableType

`func (o *TimeAccrual) GetSickAvailableType() string`

GetSickAvailableType returns the SickAvailableType field if non-nil, zero value otherwise.

### GetSickAvailableTypeOk

`func (o *TimeAccrual) GetSickAvailableTypeOk() (*string, bool)`

GetSickAvailableTypeOk returns a tuple with the SickAvailableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSickAvailableType

`func (o *TimeAccrual) SetSickAvailableType(v string)`

SetSickAvailableType sets SickAvailableType field to given value.

### HasSickAvailableType

`func (o *TimeAccrual) HasSickAvailableType() bool`

HasSickAvailableType returns a boolean if a field has been set.

### SetSickAvailableTypeNil

`func (o *TimeAccrual) SetSickAvailableTypeNil(b bool)`

 SetSickAvailableTypeNil sets the value for SickAvailableType to be an explicit nil

### UnsetSickAvailableType
`func (o *TimeAccrual) UnsetSickAvailableType()`

UnsetSickAvailableType ensures that no value is present for SickAvailableType, not even an explicit nil
### GetSickCarryoverAllowedFlag

`func (o *TimeAccrual) GetSickCarryoverAllowedFlag() bool`

GetSickCarryoverAllowedFlag returns the SickCarryoverAllowedFlag field if non-nil, zero value otherwise.

### GetSickCarryoverAllowedFlagOk

`func (o *TimeAccrual) GetSickCarryoverAllowedFlagOk() (*bool, bool)`

GetSickCarryoverAllowedFlagOk returns a tuple with the SickCarryoverAllowedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSickCarryoverAllowedFlag

`func (o *TimeAccrual) SetSickCarryoverAllowedFlag(v bool)`

SetSickCarryoverAllowedFlag sets SickCarryoverAllowedFlag field to given value.

### HasSickCarryoverAllowedFlag

`func (o *TimeAccrual) HasSickCarryoverAllowedFlag() bool`

HasSickCarryoverAllowedFlag returns a boolean if a field has been set.

### SetSickCarryoverAllowedFlagNil

`func (o *TimeAccrual) SetSickCarryoverAllowedFlagNil(b bool)`

 SetSickCarryoverAllowedFlagNil sets the value for SickCarryoverAllowedFlag to be an explicit nil

### UnsetSickCarryoverAllowedFlag
`func (o *TimeAccrual) UnsetSickCarryoverAllowedFlag()`

UnsetSickCarryoverAllowedFlag ensures that no value is present for SickCarryoverAllowedFlag, not even an explicit nil
### GetSickCarryoverLimit

`func (o *TimeAccrual) GetSickCarryoverLimit() float64`

GetSickCarryoverLimit returns the SickCarryoverLimit field if non-nil, zero value otherwise.

### GetSickCarryoverLimitOk

`func (o *TimeAccrual) GetSickCarryoverLimitOk() (*float64, bool)`

GetSickCarryoverLimitOk returns a tuple with the SickCarryoverLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSickCarryoverLimit

`func (o *TimeAccrual) SetSickCarryoverLimit(v float64)`

SetSickCarryoverLimit sets SickCarryoverLimit field to given value.

### HasSickCarryoverLimit

`func (o *TimeAccrual) HasSickCarryoverLimit() bool`

HasSickCarryoverLimit returns a boolean if a field has been set.

### SetSickCarryoverLimitNil

`func (o *TimeAccrual) SetSickCarryoverLimitNil(b bool)`

 SetSickCarryoverLimitNil sets the value for SickCarryoverLimit to be an explicit nil

### UnsetSickCarryoverLimit
`func (o *TimeAccrual) UnsetSickCarryoverLimit()`

UnsetSickCarryoverLimit ensures that no value is present for SickCarryoverLimit, not even an explicit nil
### GetPtoFlag

`func (o *TimeAccrual) GetPtoFlag() bool`

GetPtoFlag returns the PtoFlag field if non-nil, zero value otherwise.

### GetPtoFlagOk

`func (o *TimeAccrual) GetPtoFlagOk() (*bool, bool)`

GetPtoFlagOk returns a tuple with the PtoFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtoFlag

`func (o *TimeAccrual) SetPtoFlag(v bool)`

SetPtoFlag sets PtoFlag field to given value.

### HasPtoFlag

`func (o *TimeAccrual) HasPtoFlag() bool`

HasPtoFlag returns a boolean if a field has been set.

### SetPtoFlagNil

`func (o *TimeAccrual) SetPtoFlagNil(b bool)`

 SetPtoFlagNil sets the value for PtoFlag to be an explicit nil

### UnsetPtoFlag
`func (o *TimeAccrual) UnsetPtoFlag()`

UnsetPtoFlag ensures that no value is present for PtoFlag, not even an explicit nil
### GetPtoAvailableType

`func (o *TimeAccrual) GetPtoAvailableType() string`

GetPtoAvailableType returns the PtoAvailableType field if non-nil, zero value otherwise.

### GetPtoAvailableTypeOk

`func (o *TimeAccrual) GetPtoAvailableTypeOk() (*string, bool)`

GetPtoAvailableTypeOk returns a tuple with the PtoAvailableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtoAvailableType

`func (o *TimeAccrual) SetPtoAvailableType(v string)`

SetPtoAvailableType sets PtoAvailableType field to given value.

### HasPtoAvailableType

`func (o *TimeAccrual) HasPtoAvailableType() bool`

HasPtoAvailableType returns a boolean if a field has been set.

### SetPtoAvailableTypeNil

`func (o *TimeAccrual) SetPtoAvailableTypeNil(b bool)`

 SetPtoAvailableTypeNil sets the value for PtoAvailableType to be an explicit nil

### UnsetPtoAvailableType
`func (o *TimeAccrual) UnsetPtoAvailableType()`

UnsetPtoAvailableType ensures that no value is present for PtoAvailableType, not even an explicit nil
### GetPtoCarryoverAllowedFlag

`func (o *TimeAccrual) GetPtoCarryoverAllowedFlag() bool`

GetPtoCarryoverAllowedFlag returns the PtoCarryoverAllowedFlag field if non-nil, zero value otherwise.

### GetPtoCarryoverAllowedFlagOk

`func (o *TimeAccrual) GetPtoCarryoverAllowedFlagOk() (*bool, bool)`

GetPtoCarryoverAllowedFlagOk returns a tuple with the PtoCarryoverAllowedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtoCarryoverAllowedFlag

`func (o *TimeAccrual) SetPtoCarryoverAllowedFlag(v bool)`

SetPtoCarryoverAllowedFlag sets PtoCarryoverAllowedFlag field to given value.

### HasPtoCarryoverAllowedFlag

`func (o *TimeAccrual) HasPtoCarryoverAllowedFlag() bool`

HasPtoCarryoverAllowedFlag returns a boolean if a field has been set.

### SetPtoCarryoverAllowedFlagNil

`func (o *TimeAccrual) SetPtoCarryoverAllowedFlagNil(b bool)`

 SetPtoCarryoverAllowedFlagNil sets the value for PtoCarryoverAllowedFlag to be an explicit nil

### UnsetPtoCarryoverAllowedFlag
`func (o *TimeAccrual) UnsetPtoCarryoverAllowedFlag()`

UnsetPtoCarryoverAllowedFlag ensures that no value is present for PtoCarryoverAllowedFlag, not even an explicit nil
### GetPtoCarryoverLimit

`func (o *TimeAccrual) GetPtoCarryoverLimit() float64`

GetPtoCarryoverLimit returns the PtoCarryoverLimit field if non-nil, zero value otherwise.

### GetPtoCarryoverLimitOk

`func (o *TimeAccrual) GetPtoCarryoverLimitOk() (*float64, bool)`

GetPtoCarryoverLimitOk returns a tuple with the PtoCarryoverLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtoCarryoverLimit

`func (o *TimeAccrual) SetPtoCarryoverLimit(v float64)`

SetPtoCarryoverLimit sets PtoCarryoverLimit field to given value.

### HasPtoCarryoverLimit

`func (o *TimeAccrual) HasPtoCarryoverLimit() bool`

HasPtoCarryoverLimit returns a boolean if a field has been set.

### SetPtoCarryoverLimitNil

`func (o *TimeAccrual) SetPtoCarryoverLimitNil(b bool)`

 SetPtoCarryoverLimitNil sets the value for PtoCarryoverLimit to be an explicit nil

### UnsetPtoCarryoverLimit
`func (o *TimeAccrual) UnsetPtoCarryoverLimit()`

UnsetPtoCarryoverLimit ensures that no value is present for PtoCarryoverLimit, not even an explicit nil
### GetHolidayFlag

`func (o *TimeAccrual) GetHolidayFlag() bool`

GetHolidayFlag returns the HolidayFlag field if non-nil, zero value otherwise.

### GetHolidayFlagOk

`func (o *TimeAccrual) GetHolidayFlagOk() (*bool, bool)`

GetHolidayFlagOk returns a tuple with the HolidayFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidayFlag

`func (o *TimeAccrual) SetHolidayFlag(v bool)`

SetHolidayFlag sets HolidayFlag field to given value.

### HasHolidayFlag

`func (o *TimeAccrual) HasHolidayFlag() bool`

HasHolidayFlag returns a boolean if a field has been set.

### SetHolidayFlagNil

`func (o *TimeAccrual) SetHolidayFlagNil(b bool)`

 SetHolidayFlagNil sets the value for HolidayFlag to be an explicit nil

### UnsetHolidayFlag
`func (o *TimeAccrual) UnsetHolidayFlag()`

UnsetHolidayFlag ensures that no value is present for HolidayFlag, not even an explicit nil
### GetHolidayAvailableType

`func (o *TimeAccrual) GetHolidayAvailableType() string`

GetHolidayAvailableType returns the HolidayAvailableType field if non-nil, zero value otherwise.

### GetHolidayAvailableTypeOk

`func (o *TimeAccrual) GetHolidayAvailableTypeOk() (*string, bool)`

GetHolidayAvailableTypeOk returns a tuple with the HolidayAvailableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidayAvailableType

`func (o *TimeAccrual) SetHolidayAvailableType(v string)`

SetHolidayAvailableType sets HolidayAvailableType field to given value.

### HasHolidayAvailableType

`func (o *TimeAccrual) HasHolidayAvailableType() bool`

HasHolidayAvailableType returns a boolean if a field has been set.

### SetHolidayAvailableTypeNil

`func (o *TimeAccrual) SetHolidayAvailableTypeNil(b bool)`

 SetHolidayAvailableTypeNil sets the value for HolidayAvailableType to be an explicit nil

### UnsetHolidayAvailableType
`func (o *TimeAccrual) UnsetHolidayAvailableType()`

UnsetHolidayAvailableType ensures that no value is present for HolidayAvailableType, not even an explicit nil
### GetHolidayCarryoverAllowedFlag

`func (o *TimeAccrual) GetHolidayCarryoverAllowedFlag() bool`

GetHolidayCarryoverAllowedFlag returns the HolidayCarryoverAllowedFlag field if non-nil, zero value otherwise.

### GetHolidayCarryoverAllowedFlagOk

`func (o *TimeAccrual) GetHolidayCarryoverAllowedFlagOk() (*bool, bool)`

GetHolidayCarryoverAllowedFlagOk returns a tuple with the HolidayCarryoverAllowedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidayCarryoverAllowedFlag

`func (o *TimeAccrual) SetHolidayCarryoverAllowedFlag(v bool)`

SetHolidayCarryoverAllowedFlag sets HolidayCarryoverAllowedFlag field to given value.

### HasHolidayCarryoverAllowedFlag

`func (o *TimeAccrual) HasHolidayCarryoverAllowedFlag() bool`

HasHolidayCarryoverAllowedFlag returns a boolean if a field has been set.

### SetHolidayCarryoverAllowedFlagNil

`func (o *TimeAccrual) SetHolidayCarryoverAllowedFlagNil(b bool)`

 SetHolidayCarryoverAllowedFlagNil sets the value for HolidayCarryoverAllowedFlag to be an explicit nil

### UnsetHolidayCarryoverAllowedFlag
`func (o *TimeAccrual) UnsetHolidayCarryoverAllowedFlag()`

UnsetHolidayCarryoverAllowedFlag ensures that no value is present for HolidayCarryoverAllowedFlag, not even an explicit nil
### GetHolidayCarryoverLimit

`func (o *TimeAccrual) GetHolidayCarryoverLimit() float64`

GetHolidayCarryoverLimit returns the HolidayCarryoverLimit field if non-nil, zero value otherwise.

### GetHolidayCarryoverLimitOk

`func (o *TimeAccrual) GetHolidayCarryoverLimitOk() (*float64, bool)`

GetHolidayCarryoverLimitOk returns a tuple with the HolidayCarryoverLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidayCarryoverLimit

`func (o *TimeAccrual) SetHolidayCarryoverLimit(v float64)`

SetHolidayCarryoverLimit sets HolidayCarryoverLimit field to given value.

### HasHolidayCarryoverLimit

`func (o *TimeAccrual) HasHolidayCarryoverLimit() bool`

HasHolidayCarryoverLimit returns a boolean if a field has been set.

### SetHolidayCarryoverLimitNil

`func (o *TimeAccrual) SetHolidayCarryoverLimitNil(b bool)`

 SetHolidayCarryoverLimitNil sets the value for HolidayCarryoverLimit to be an explicit nil

### UnsetHolidayCarryoverLimit
`func (o *TimeAccrual) UnsetHolidayCarryoverLimit()`

UnsetHolidayCarryoverLimit ensures that no value is present for HolidayCarryoverLimit, not even an explicit nil
### GetInfo

`func (o *TimeAccrual) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TimeAccrual) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TimeAccrual) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TimeAccrual) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


