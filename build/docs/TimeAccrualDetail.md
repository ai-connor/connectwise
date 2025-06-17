# TimeAccrualDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AccrualType** | **NullableString** | Available types are: Holiday, PTO, Sick and Vacation. | 
**StartYear** | **NullableInt32** |  | 
**EndYear** | **NullableInt32** |  | 
**Hours** | **NullableFloat64** |  | 
**TimeAccrual** | Pointer to [**TimeAccrualReference**](TimeAccrualReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTimeAccrualDetail

`func NewTimeAccrualDetail(accrualType NullableString, startYear NullableInt32, endYear NullableInt32, hours NullableFloat64, ) *TimeAccrualDetail`

NewTimeAccrualDetail instantiates a new TimeAccrualDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeAccrualDetailWithDefaults

`func NewTimeAccrualDetailWithDefaults() *TimeAccrualDetail`

NewTimeAccrualDetailWithDefaults instantiates a new TimeAccrualDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimeAccrualDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimeAccrualDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimeAccrualDetail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TimeAccrualDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccrualType

`func (o *TimeAccrualDetail) GetAccrualType() string`

GetAccrualType returns the AccrualType field if non-nil, zero value otherwise.

### GetAccrualTypeOk

`func (o *TimeAccrualDetail) GetAccrualTypeOk() (*string, bool)`

GetAccrualTypeOk returns a tuple with the AccrualType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualType

`func (o *TimeAccrualDetail) SetAccrualType(v string)`

SetAccrualType sets AccrualType field to given value.


### SetAccrualTypeNil

`func (o *TimeAccrualDetail) SetAccrualTypeNil(b bool)`

 SetAccrualTypeNil sets the value for AccrualType to be an explicit nil

### UnsetAccrualType
`func (o *TimeAccrualDetail) UnsetAccrualType()`

UnsetAccrualType ensures that no value is present for AccrualType, not even an explicit nil
### GetStartYear

`func (o *TimeAccrualDetail) GetStartYear() int32`

GetStartYear returns the StartYear field if non-nil, zero value otherwise.

### GetStartYearOk

`func (o *TimeAccrualDetail) GetStartYearOk() (*int32, bool)`

GetStartYearOk returns a tuple with the StartYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartYear

`func (o *TimeAccrualDetail) SetStartYear(v int32)`

SetStartYear sets StartYear field to given value.


### SetStartYearNil

`func (o *TimeAccrualDetail) SetStartYearNil(b bool)`

 SetStartYearNil sets the value for StartYear to be an explicit nil

### UnsetStartYear
`func (o *TimeAccrualDetail) UnsetStartYear()`

UnsetStartYear ensures that no value is present for StartYear, not even an explicit nil
### GetEndYear

`func (o *TimeAccrualDetail) GetEndYear() int32`

GetEndYear returns the EndYear field if non-nil, zero value otherwise.

### GetEndYearOk

`func (o *TimeAccrualDetail) GetEndYearOk() (*int32, bool)`

GetEndYearOk returns a tuple with the EndYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndYear

`func (o *TimeAccrualDetail) SetEndYear(v int32)`

SetEndYear sets EndYear field to given value.


### SetEndYearNil

`func (o *TimeAccrualDetail) SetEndYearNil(b bool)`

 SetEndYearNil sets the value for EndYear to be an explicit nil

### UnsetEndYear
`func (o *TimeAccrualDetail) UnsetEndYear()`

UnsetEndYear ensures that no value is present for EndYear, not even an explicit nil
### GetHours

`func (o *TimeAccrualDetail) GetHours() float64`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *TimeAccrualDetail) GetHoursOk() (*float64, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *TimeAccrualDetail) SetHours(v float64)`

SetHours sets Hours field to given value.


### SetHoursNil

`func (o *TimeAccrualDetail) SetHoursNil(b bool)`

 SetHoursNil sets the value for Hours to be an explicit nil

### UnsetHours
`func (o *TimeAccrualDetail) UnsetHours()`

UnsetHours ensures that no value is present for Hours, not even an explicit nil
### GetTimeAccrual

`func (o *TimeAccrualDetail) GetTimeAccrual() TimeAccrualReference`

GetTimeAccrual returns the TimeAccrual field if non-nil, zero value otherwise.

### GetTimeAccrualOk

`func (o *TimeAccrualDetail) GetTimeAccrualOk() (*TimeAccrualReference, bool)`

GetTimeAccrualOk returns a tuple with the TimeAccrual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAccrual

`func (o *TimeAccrualDetail) SetTimeAccrual(v TimeAccrualReference)`

SetTimeAccrual sets TimeAccrual field to given value.

### HasTimeAccrual

`func (o *TimeAccrualDetail) HasTimeAccrual() bool`

HasTimeAccrual returns a boolean if a field has been set.

### GetInfo

`func (o *TimeAccrualDetail) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TimeAccrualDetail) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TimeAccrualDetail) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TimeAccrualDetail) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


