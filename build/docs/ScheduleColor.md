# ScheduleColor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**StartPercent** | Pointer to **NullableInt32** | A startPercent (0 or higher) is required if endPercent has value. | [optional] 
**EndPercent** | Pointer to **NullableInt32** | A endPercent is required if startPercent has value. | [optional] 
**Color** | **string** | Must be a valid Hexadecimal Color Code. | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewScheduleColor

`func NewScheduleColor(color string, ) *ScheduleColor`

NewScheduleColor instantiates a new ScheduleColor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleColorWithDefaults

`func NewScheduleColorWithDefaults() *ScheduleColor`

NewScheduleColorWithDefaults instantiates a new ScheduleColor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduleColor) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleColor) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleColor) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ScheduleColor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStartPercent

`func (o *ScheduleColor) GetStartPercent() int32`

GetStartPercent returns the StartPercent field if non-nil, zero value otherwise.

### GetStartPercentOk

`func (o *ScheduleColor) GetStartPercentOk() (*int32, bool)`

GetStartPercentOk returns a tuple with the StartPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPercent

`func (o *ScheduleColor) SetStartPercent(v int32)`

SetStartPercent sets StartPercent field to given value.

### HasStartPercent

`func (o *ScheduleColor) HasStartPercent() bool`

HasStartPercent returns a boolean if a field has been set.

### SetStartPercentNil

`func (o *ScheduleColor) SetStartPercentNil(b bool)`

 SetStartPercentNil sets the value for StartPercent to be an explicit nil

### UnsetStartPercent
`func (o *ScheduleColor) UnsetStartPercent()`

UnsetStartPercent ensures that no value is present for StartPercent, not even an explicit nil
### GetEndPercent

`func (o *ScheduleColor) GetEndPercent() int32`

GetEndPercent returns the EndPercent field if non-nil, zero value otherwise.

### GetEndPercentOk

`func (o *ScheduleColor) GetEndPercentOk() (*int32, bool)`

GetEndPercentOk returns a tuple with the EndPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPercent

`func (o *ScheduleColor) SetEndPercent(v int32)`

SetEndPercent sets EndPercent field to given value.

### HasEndPercent

`func (o *ScheduleColor) HasEndPercent() bool`

HasEndPercent returns a boolean if a field has been set.

### SetEndPercentNil

`func (o *ScheduleColor) SetEndPercentNil(b bool)`

 SetEndPercentNil sets the value for EndPercent to be an explicit nil

### UnsetEndPercent
`func (o *ScheduleColor) UnsetEndPercent()`

UnsetEndPercent ensures that no value is present for EndPercent, not even an explicit nil
### GetColor

`func (o *ScheduleColor) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ScheduleColor) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ScheduleColor) SetColor(v string)`

SetColor sets Color field to given value.


### GetInfo

`func (o *ScheduleColor) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ScheduleColor) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ScheduleColor) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ScheduleColor) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


