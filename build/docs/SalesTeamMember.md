# SalesTeamMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Member** | [**MemberReference**](MemberReference.md) |  | 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**AllowAccessFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSalesTeamMember

`func NewSalesTeamMember(member MemberReference, ) *SalesTeamMember`

NewSalesTeamMember instantiates a new SalesTeamMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesTeamMemberWithDefaults

`func NewSalesTeamMemberWithDefaults() *SalesTeamMember`

NewSalesTeamMemberWithDefaults instantiates a new SalesTeamMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SalesTeamMember) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SalesTeamMember) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SalesTeamMember) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SalesTeamMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMember

`func (o *SalesTeamMember) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *SalesTeamMember) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *SalesTeamMember) SetMember(v MemberReference)`

SetMember sets Member field to given value.


### GetLocation

`func (o *SalesTeamMember) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SalesTeamMember) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SalesTeamMember) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SalesTeamMember) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartment

`func (o *SalesTeamMember) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *SalesTeamMember) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *SalesTeamMember) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *SalesTeamMember) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetAllowAccessFlag

`func (o *SalesTeamMember) GetAllowAccessFlag() bool`

GetAllowAccessFlag returns the AllowAccessFlag field if non-nil, zero value otherwise.

### GetAllowAccessFlagOk

`func (o *SalesTeamMember) GetAllowAccessFlagOk() (*bool, bool)`

GetAllowAccessFlagOk returns a tuple with the AllowAccessFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAccessFlag

`func (o *SalesTeamMember) SetAllowAccessFlag(v bool)`

SetAllowAccessFlag sets AllowAccessFlag field to given value.

### HasAllowAccessFlag

`func (o *SalesTeamMember) HasAllowAccessFlag() bool`

HasAllowAccessFlag returns a boolean if a field has been set.

### SetAllowAccessFlagNil

`func (o *SalesTeamMember) SetAllowAccessFlagNil(b bool)`

 SetAllowAccessFlagNil sets the value for AllowAccessFlag to be an explicit nil

### UnsetAllowAccessFlag
`func (o *SalesTeamMember) UnsetAllowAccessFlag()`

UnsetAllowAccessFlag ensures that no value is present for AllowAccessFlag, not even an explicit nil
### GetInfo

`func (o *SalesTeamMember) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SalesTeamMember) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SalesTeamMember) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SalesTeamMember) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


