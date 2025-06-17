# MemberAccrual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AccrualType** | **NullableString** |  | 
**Year** | **NullableInt32** |  | 
**Hours** | **NullableFloat64** |  | 
**Reason** | **string** |  | 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMemberAccrual

`func NewMemberAccrual(accrualType NullableString, year NullableInt32, hours NullableFloat64, reason string, ) *MemberAccrual`

NewMemberAccrual instantiates a new MemberAccrual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberAccrualWithDefaults

`func NewMemberAccrualWithDefaults() *MemberAccrual`

NewMemberAccrualWithDefaults instantiates a new MemberAccrual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MemberAccrual) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberAccrual) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberAccrual) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MemberAccrual) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccrualType

`func (o *MemberAccrual) GetAccrualType() string`

GetAccrualType returns the AccrualType field if non-nil, zero value otherwise.

### GetAccrualTypeOk

`func (o *MemberAccrual) GetAccrualTypeOk() (*string, bool)`

GetAccrualTypeOk returns a tuple with the AccrualType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualType

`func (o *MemberAccrual) SetAccrualType(v string)`

SetAccrualType sets AccrualType field to given value.


### SetAccrualTypeNil

`func (o *MemberAccrual) SetAccrualTypeNil(b bool)`

 SetAccrualTypeNil sets the value for AccrualType to be an explicit nil

### UnsetAccrualType
`func (o *MemberAccrual) UnsetAccrualType()`

UnsetAccrualType ensures that no value is present for AccrualType, not even an explicit nil
### GetYear

`func (o *MemberAccrual) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *MemberAccrual) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *MemberAccrual) SetYear(v int32)`

SetYear sets Year field to given value.


### SetYearNil

`func (o *MemberAccrual) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *MemberAccrual) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetHours

`func (o *MemberAccrual) GetHours() float64`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *MemberAccrual) GetHoursOk() (*float64, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *MemberAccrual) SetHours(v float64)`

SetHours sets Hours field to given value.


### SetHoursNil

`func (o *MemberAccrual) SetHoursNil(b bool)`

 SetHoursNil sets the value for Hours to be an explicit nil

### UnsetHours
`func (o *MemberAccrual) UnsetHours()`

UnsetHours ensures that no value is present for Hours, not even an explicit nil
### GetReason

`func (o *MemberAccrual) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MemberAccrual) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MemberAccrual) SetReason(v string)`

SetReason sets Reason field to given value.


### GetMember

`func (o *MemberAccrual) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *MemberAccrual) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *MemberAccrual) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *MemberAccrual) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetInfo

`func (o *MemberAccrual) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *MemberAccrual) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *MemberAccrual) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *MemberAccrual) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


