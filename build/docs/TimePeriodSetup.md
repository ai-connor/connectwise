# TimePeriodSetup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**PeriodApplyTo** | **NullableString** |  | 
**Year** | **NullableInt32** |  | 
**NumberFuturePeriods** | **NullableInt32** |  | 
**Type** | **NullableString** |  | 
**Description** | Pointer to **string** |  Max length: 100; | [optional] 
**FirstPeriodEndDate** | **string** |  | 
**MonthlyPeriodEnds** | Pointer to **NullableInt32** | Only needed when type is monthly | [optional] 
**SemiMonthlyFirstPeriod** | Pointer to **NullableInt32** | Only needed when type is semi-monthly | [optional] 
**SemiMonthlySecondPeriod** | Pointer to **NullableInt32** | Only needed when type is semi-monthly | [optional] 
**SemiMonthlyLastDayFlag** | Pointer to **NullableBool** |  | [optional] 
**LastDayFlag** | Pointer to **NullableBool** | Only needed when type is monthly | [optional] 
**DaysPastEndDate** | **NullableInt32** |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTimePeriodSetup

`func NewTimePeriodSetup(periodApplyTo NullableString, year NullableInt32, numberFuturePeriods NullableInt32, type_ NullableString, firstPeriodEndDate string, daysPastEndDate NullableInt32, ) *TimePeriodSetup`

NewTimePeriodSetup instantiates a new TimePeriodSetup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimePeriodSetupWithDefaults

`func NewTimePeriodSetupWithDefaults() *TimePeriodSetup`

NewTimePeriodSetupWithDefaults instantiates a new TimePeriodSetup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimePeriodSetup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimePeriodSetup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimePeriodSetup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TimePeriodSetup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPeriodApplyTo

`func (o *TimePeriodSetup) GetPeriodApplyTo() string`

GetPeriodApplyTo returns the PeriodApplyTo field if non-nil, zero value otherwise.

### GetPeriodApplyToOk

`func (o *TimePeriodSetup) GetPeriodApplyToOk() (*string, bool)`

GetPeriodApplyToOk returns a tuple with the PeriodApplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodApplyTo

`func (o *TimePeriodSetup) SetPeriodApplyTo(v string)`

SetPeriodApplyTo sets PeriodApplyTo field to given value.


### SetPeriodApplyToNil

`func (o *TimePeriodSetup) SetPeriodApplyToNil(b bool)`

 SetPeriodApplyToNil sets the value for PeriodApplyTo to be an explicit nil

### UnsetPeriodApplyTo
`func (o *TimePeriodSetup) UnsetPeriodApplyTo()`

UnsetPeriodApplyTo ensures that no value is present for PeriodApplyTo, not even an explicit nil
### GetYear

`func (o *TimePeriodSetup) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *TimePeriodSetup) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *TimePeriodSetup) SetYear(v int32)`

SetYear sets Year field to given value.


### SetYearNil

`func (o *TimePeriodSetup) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *TimePeriodSetup) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetNumberFuturePeriods

`func (o *TimePeriodSetup) GetNumberFuturePeriods() int32`

GetNumberFuturePeriods returns the NumberFuturePeriods field if non-nil, zero value otherwise.

### GetNumberFuturePeriodsOk

`func (o *TimePeriodSetup) GetNumberFuturePeriodsOk() (*int32, bool)`

GetNumberFuturePeriodsOk returns a tuple with the NumberFuturePeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberFuturePeriods

`func (o *TimePeriodSetup) SetNumberFuturePeriods(v int32)`

SetNumberFuturePeriods sets NumberFuturePeriods field to given value.


### SetNumberFuturePeriodsNil

`func (o *TimePeriodSetup) SetNumberFuturePeriodsNil(b bool)`

 SetNumberFuturePeriodsNil sets the value for NumberFuturePeriods to be an explicit nil

### UnsetNumberFuturePeriods
`func (o *TimePeriodSetup) UnsetNumberFuturePeriods()`

UnsetNumberFuturePeriods ensures that no value is present for NumberFuturePeriods, not even an explicit nil
### GetType

`func (o *TimePeriodSetup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimePeriodSetup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimePeriodSetup) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *TimePeriodSetup) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TimePeriodSetup) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetDescription

`func (o *TimePeriodSetup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TimePeriodSetup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TimePeriodSetup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TimePeriodSetup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFirstPeriodEndDate

`func (o *TimePeriodSetup) GetFirstPeriodEndDate() string`

GetFirstPeriodEndDate returns the FirstPeriodEndDate field if non-nil, zero value otherwise.

### GetFirstPeriodEndDateOk

`func (o *TimePeriodSetup) GetFirstPeriodEndDateOk() (*string, bool)`

GetFirstPeriodEndDateOk returns a tuple with the FirstPeriodEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPeriodEndDate

`func (o *TimePeriodSetup) SetFirstPeriodEndDate(v string)`

SetFirstPeriodEndDate sets FirstPeriodEndDate field to given value.


### GetMonthlyPeriodEnds

`func (o *TimePeriodSetup) GetMonthlyPeriodEnds() int32`

GetMonthlyPeriodEnds returns the MonthlyPeriodEnds field if non-nil, zero value otherwise.

### GetMonthlyPeriodEndsOk

`func (o *TimePeriodSetup) GetMonthlyPeriodEndsOk() (*int32, bool)`

GetMonthlyPeriodEndsOk returns a tuple with the MonthlyPeriodEnds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPeriodEnds

`func (o *TimePeriodSetup) SetMonthlyPeriodEnds(v int32)`

SetMonthlyPeriodEnds sets MonthlyPeriodEnds field to given value.

### HasMonthlyPeriodEnds

`func (o *TimePeriodSetup) HasMonthlyPeriodEnds() bool`

HasMonthlyPeriodEnds returns a boolean if a field has been set.

### SetMonthlyPeriodEndsNil

`func (o *TimePeriodSetup) SetMonthlyPeriodEndsNil(b bool)`

 SetMonthlyPeriodEndsNil sets the value for MonthlyPeriodEnds to be an explicit nil

### UnsetMonthlyPeriodEnds
`func (o *TimePeriodSetup) UnsetMonthlyPeriodEnds()`

UnsetMonthlyPeriodEnds ensures that no value is present for MonthlyPeriodEnds, not even an explicit nil
### GetSemiMonthlyFirstPeriod

`func (o *TimePeriodSetup) GetSemiMonthlyFirstPeriod() int32`

GetSemiMonthlyFirstPeriod returns the SemiMonthlyFirstPeriod field if non-nil, zero value otherwise.

### GetSemiMonthlyFirstPeriodOk

`func (o *TimePeriodSetup) GetSemiMonthlyFirstPeriodOk() (*int32, bool)`

GetSemiMonthlyFirstPeriodOk returns a tuple with the SemiMonthlyFirstPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemiMonthlyFirstPeriod

`func (o *TimePeriodSetup) SetSemiMonthlyFirstPeriod(v int32)`

SetSemiMonthlyFirstPeriod sets SemiMonthlyFirstPeriod field to given value.

### HasSemiMonthlyFirstPeriod

`func (o *TimePeriodSetup) HasSemiMonthlyFirstPeriod() bool`

HasSemiMonthlyFirstPeriod returns a boolean if a field has been set.

### SetSemiMonthlyFirstPeriodNil

`func (o *TimePeriodSetup) SetSemiMonthlyFirstPeriodNil(b bool)`

 SetSemiMonthlyFirstPeriodNil sets the value for SemiMonthlyFirstPeriod to be an explicit nil

### UnsetSemiMonthlyFirstPeriod
`func (o *TimePeriodSetup) UnsetSemiMonthlyFirstPeriod()`

UnsetSemiMonthlyFirstPeriod ensures that no value is present for SemiMonthlyFirstPeriod, not even an explicit nil
### GetSemiMonthlySecondPeriod

`func (o *TimePeriodSetup) GetSemiMonthlySecondPeriod() int32`

GetSemiMonthlySecondPeriod returns the SemiMonthlySecondPeriod field if non-nil, zero value otherwise.

### GetSemiMonthlySecondPeriodOk

`func (o *TimePeriodSetup) GetSemiMonthlySecondPeriodOk() (*int32, bool)`

GetSemiMonthlySecondPeriodOk returns a tuple with the SemiMonthlySecondPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemiMonthlySecondPeriod

`func (o *TimePeriodSetup) SetSemiMonthlySecondPeriod(v int32)`

SetSemiMonthlySecondPeriod sets SemiMonthlySecondPeriod field to given value.

### HasSemiMonthlySecondPeriod

`func (o *TimePeriodSetup) HasSemiMonthlySecondPeriod() bool`

HasSemiMonthlySecondPeriod returns a boolean if a field has been set.

### SetSemiMonthlySecondPeriodNil

`func (o *TimePeriodSetup) SetSemiMonthlySecondPeriodNil(b bool)`

 SetSemiMonthlySecondPeriodNil sets the value for SemiMonthlySecondPeriod to be an explicit nil

### UnsetSemiMonthlySecondPeriod
`func (o *TimePeriodSetup) UnsetSemiMonthlySecondPeriod()`

UnsetSemiMonthlySecondPeriod ensures that no value is present for SemiMonthlySecondPeriod, not even an explicit nil
### GetSemiMonthlyLastDayFlag

`func (o *TimePeriodSetup) GetSemiMonthlyLastDayFlag() bool`

GetSemiMonthlyLastDayFlag returns the SemiMonthlyLastDayFlag field if non-nil, zero value otherwise.

### GetSemiMonthlyLastDayFlagOk

`func (o *TimePeriodSetup) GetSemiMonthlyLastDayFlagOk() (*bool, bool)`

GetSemiMonthlyLastDayFlagOk returns a tuple with the SemiMonthlyLastDayFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemiMonthlyLastDayFlag

`func (o *TimePeriodSetup) SetSemiMonthlyLastDayFlag(v bool)`

SetSemiMonthlyLastDayFlag sets SemiMonthlyLastDayFlag field to given value.

### HasSemiMonthlyLastDayFlag

`func (o *TimePeriodSetup) HasSemiMonthlyLastDayFlag() bool`

HasSemiMonthlyLastDayFlag returns a boolean if a field has been set.

### SetSemiMonthlyLastDayFlagNil

`func (o *TimePeriodSetup) SetSemiMonthlyLastDayFlagNil(b bool)`

 SetSemiMonthlyLastDayFlagNil sets the value for SemiMonthlyLastDayFlag to be an explicit nil

### UnsetSemiMonthlyLastDayFlag
`func (o *TimePeriodSetup) UnsetSemiMonthlyLastDayFlag()`

UnsetSemiMonthlyLastDayFlag ensures that no value is present for SemiMonthlyLastDayFlag, not even an explicit nil
### GetLastDayFlag

`func (o *TimePeriodSetup) GetLastDayFlag() bool`

GetLastDayFlag returns the LastDayFlag field if non-nil, zero value otherwise.

### GetLastDayFlagOk

`func (o *TimePeriodSetup) GetLastDayFlagOk() (*bool, bool)`

GetLastDayFlagOk returns a tuple with the LastDayFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDayFlag

`func (o *TimePeriodSetup) SetLastDayFlag(v bool)`

SetLastDayFlag sets LastDayFlag field to given value.

### HasLastDayFlag

`func (o *TimePeriodSetup) HasLastDayFlag() bool`

HasLastDayFlag returns a boolean if a field has been set.

### SetLastDayFlagNil

`func (o *TimePeriodSetup) SetLastDayFlagNil(b bool)`

 SetLastDayFlagNil sets the value for LastDayFlag to be an explicit nil

### UnsetLastDayFlag
`func (o *TimePeriodSetup) UnsetLastDayFlag()`

UnsetLastDayFlag ensures that no value is present for LastDayFlag, not even an explicit nil
### GetDaysPastEndDate

`func (o *TimePeriodSetup) GetDaysPastEndDate() int32`

GetDaysPastEndDate returns the DaysPastEndDate field if non-nil, zero value otherwise.

### GetDaysPastEndDateOk

`func (o *TimePeriodSetup) GetDaysPastEndDateOk() (*int32, bool)`

GetDaysPastEndDateOk returns a tuple with the DaysPastEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysPastEndDate

`func (o *TimePeriodSetup) SetDaysPastEndDate(v int32)`

SetDaysPastEndDate sets DaysPastEndDate field to given value.


### SetDaysPastEndDateNil

`func (o *TimePeriodSetup) SetDaysPastEndDateNil(b bool)`

 SetDaysPastEndDateNil sets the value for DaysPastEndDate to be an explicit nil

### UnsetDaysPastEndDate
`func (o *TimePeriodSetup) UnsetDaysPastEndDate()`

UnsetDaysPastEndDate ensures that no value is present for DaysPastEndDate, not even an explicit nil
### GetInfo

`func (o *TimePeriodSetup) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TimePeriodSetup) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TimePeriodSetup) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TimePeriodSetup) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


