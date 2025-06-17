# ServiceTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Leader** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**DeleteNotifyFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewServiceTeam

`func NewServiceTeam() *ServiceTeam`

NewServiceTeam instantiates a new ServiceTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTeamWithDefaults

`func NewServiceTeamWithDefaults() *ServiceTeam`

NewServiceTeamWithDefaults instantiates a new ServiceTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceTeam) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceTeam) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceTeam) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceTeam) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceTeam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceTeam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceTeam) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceTeam) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLeader

`func (o *ServiceTeam) GetLeader() MemberReference`

GetLeader returns the Leader field if non-nil, zero value otherwise.

### GetLeaderOk

`func (o *ServiceTeam) GetLeaderOk() (*MemberReference, bool)`

GetLeaderOk returns a tuple with the Leader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeader

`func (o *ServiceTeam) SetLeader(v MemberReference)`

SetLeader sets Leader field to given value.

### HasLeader

`func (o *ServiceTeam) HasLeader() bool`

HasLeader returns a boolean if a field has been set.

### GetLocation

`func (o *ServiceTeam) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ServiceTeam) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ServiceTeam) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ServiceTeam) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartment

`func (o *ServiceTeam) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *ServiceTeam) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *ServiceTeam) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *ServiceTeam) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetDeleteNotifyFlag

`func (o *ServiceTeam) GetDeleteNotifyFlag() bool`

GetDeleteNotifyFlag returns the DeleteNotifyFlag field if non-nil, zero value otherwise.

### GetDeleteNotifyFlagOk

`func (o *ServiceTeam) GetDeleteNotifyFlagOk() (*bool, bool)`

GetDeleteNotifyFlagOk returns a tuple with the DeleteNotifyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteNotifyFlag

`func (o *ServiceTeam) SetDeleteNotifyFlag(v bool)`

SetDeleteNotifyFlag sets DeleteNotifyFlag field to given value.

### HasDeleteNotifyFlag

`func (o *ServiceTeam) HasDeleteNotifyFlag() bool`

HasDeleteNotifyFlag returns a boolean if a field has been set.

### SetDeleteNotifyFlagNil

`func (o *ServiceTeam) SetDeleteNotifyFlagNil(b bool)`

 SetDeleteNotifyFlagNil sets the value for DeleteNotifyFlag to be an explicit nil

### UnsetDeleteNotifyFlag
`func (o *ServiceTeam) UnsetDeleteNotifyFlag()`

UnsetDeleteNotifyFlag ensures that no value is present for DeleteNotifyFlag, not even an explicit nil
### GetInfo

`func (o *ServiceTeam) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ServiceTeam) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ServiceTeam) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ServiceTeam) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


