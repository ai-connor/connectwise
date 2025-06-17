# BoardTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 30; | 
**TeamLeader** | [**MemberReference**](MemberReference.md) |  | 
**Members** | Pointer to **[]int32** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**NotifyOnTicketDelete** | Pointer to **NullableBool** |  | [optional] 
**DefaultRoundRobinFlag** | Pointer to **NullableBool** |  | [optional] 
**RoundRobinFlag** | Pointer to **NullableBool** |  | [optional] 
**BoardId** | Pointer to **NullableInt32** |  | [optional] 
**LocationId** | Pointer to **NullableInt32** |  | [optional] 
**BusinessUnitId** | Pointer to **NullableInt32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBoardTeam

`func NewBoardTeam(name string, teamLeader MemberReference, ) *BoardTeam`

NewBoardTeam instantiates a new BoardTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardTeamWithDefaults

`func NewBoardTeamWithDefaults() *BoardTeam`

NewBoardTeamWithDefaults instantiates a new BoardTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BoardTeam) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoardTeam) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoardTeam) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BoardTeam) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BoardTeam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BoardTeam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BoardTeam) SetName(v string)`

SetName sets Name field to given value.


### GetTeamLeader

`func (o *BoardTeam) GetTeamLeader() MemberReference`

GetTeamLeader returns the TeamLeader field if non-nil, zero value otherwise.

### GetTeamLeaderOk

`func (o *BoardTeam) GetTeamLeaderOk() (*MemberReference, bool)`

GetTeamLeaderOk returns a tuple with the TeamLeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamLeader

`func (o *BoardTeam) SetTeamLeader(v MemberReference)`

SetTeamLeader sets TeamLeader field to given value.


### GetMembers

`func (o *BoardTeam) GetMembers() []int32`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *BoardTeam) GetMembersOk() (*[]int32, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *BoardTeam) SetMembers(v []int32)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *BoardTeam) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetDefaultFlag

`func (o *BoardTeam) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *BoardTeam) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *BoardTeam) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *BoardTeam) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *BoardTeam) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *BoardTeam) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetNotifyOnTicketDelete

`func (o *BoardTeam) GetNotifyOnTicketDelete() bool`

GetNotifyOnTicketDelete returns the NotifyOnTicketDelete field if non-nil, zero value otherwise.

### GetNotifyOnTicketDeleteOk

`func (o *BoardTeam) GetNotifyOnTicketDeleteOk() (*bool, bool)`

GetNotifyOnTicketDeleteOk returns a tuple with the NotifyOnTicketDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnTicketDelete

`func (o *BoardTeam) SetNotifyOnTicketDelete(v bool)`

SetNotifyOnTicketDelete sets NotifyOnTicketDelete field to given value.

### HasNotifyOnTicketDelete

`func (o *BoardTeam) HasNotifyOnTicketDelete() bool`

HasNotifyOnTicketDelete returns a boolean if a field has been set.

### SetNotifyOnTicketDeleteNil

`func (o *BoardTeam) SetNotifyOnTicketDeleteNil(b bool)`

 SetNotifyOnTicketDeleteNil sets the value for NotifyOnTicketDelete to be an explicit nil

### UnsetNotifyOnTicketDelete
`func (o *BoardTeam) UnsetNotifyOnTicketDelete()`

UnsetNotifyOnTicketDelete ensures that no value is present for NotifyOnTicketDelete, not even an explicit nil
### GetDefaultRoundRobinFlag

`func (o *BoardTeam) GetDefaultRoundRobinFlag() bool`

GetDefaultRoundRobinFlag returns the DefaultRoundRobinFlag field if non-nil, zero value otherwise.

### GetDefaultRoundRobinFlagOk

`func (o *BoardTeam) GetDefaultRoundRobinFlagOk() (*bool, bool)`

GetDefaultRoundRobinFlagOk returns a tuple with the DefaultRoundRobinFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoundRobinFlag

`func (o *BoardTeam) SetDefaultRoundRobinFlag(v bool)`

SetDefaultRoundRobinFlag sets DefaultRoundRobinFlag field to given value.

### HasDefaultRoundRobinFlag

`func (o *BoardTeam) HasDefaultRoundRobinFlag() bool`

HasDefaultRoundRobinFlag returns a boolean if a field has been set.

### SetDefaultRoundRobinFlagNil

`func (o *BoardTeam) SetDefaultRoundRobinFlagNil(b bool)`

 SetDefaultRoundRobinFlagNil sets the value for DefaultRoundRobinFlag to be an explicit nil

### UnsetDefaultRoundRobinFlag
`func (o *BoardTeam) UnsetDefaultRoundRobinFlag()`

UnsetDefaultRoundRobinFlag ensures that no value is present for DefaultRoundRobinFlag, not even an explicit nil
### GetRoundRobinFlag

`func (o *BoardTeam) GetRoundRobinFlag() bool`

GetRoundRobinFlag returns the RoundRobinFlag field if non-nil, zero value otherwise.

### GetRoundRobinFlagOk

`func (o *BoardTeam) GetRoundRobinFlagOk() (*bool, bool)`

GetRoundRobinFlagOk returns a tuple with the RoundRobinFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundRobinFlag

`func (o *BoardTeam) SetRoundRobinFlag(v bool)`

SetRoundRobinFlag sets RoundRobinFlag field to given value.

### HasRoundRobinFlag

`func (o *BoardTeam) HasRoundRobinFlag() bool`

HasRoundRobinFlag returns a boolean if a field has been set.

### SetRoundRobinFlagNil

`func (o *BoardTeam) SetRoundRobinFlagNil(b bool)`

 SetRoundRobinFlagNil sets the value for RoundRobinFlag to be an explicit nil

### UnsetRoundRobinFlag
`func (o *BoardTeam) UnsetRoundRobinFlag()`

UnsetRoundRobinFlag ensures that no value is present for RoundRobinFlag, not even an explicit nil
### GetBoardId

`func (o *BoardTeam) GetBoardId() int32`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *BoardTeam) GetBoardIdOk() (*int32, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *BoardTeam) SetBoardId(v int32)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *BoardTeam) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### SetBoardIdNil

`func (o *BoardTeam) SetBoardIdNil(b bool)`

 SetBoardIdNil sets the value for BoardId to be an explicit nil

### UnsetBoardId
`func (o *BoardTeam) UnsetBoardId()`

UnsetBoardId ensures that no value is present for BoardId, not even an explicit nil
### GetLocationId

`func (o *BoardTeam) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *BoardTeam) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *BoardTeam) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *BoardTeam) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *BoardTeam) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *BoardTeam) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetBusinessUnitId

`func (o *BoardTeam) GetBusinessUnitId() int32`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *BoardTeam) GetBusinessUnitIdOk() (*int32, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *BoardTeam) SetBusinessUnitId(v int32)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *BoardTeam) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### SetBusinessUnitIdNil

`func (o *BoardTeam) SetBusinessUnitIdNil(b bool)`

 SetBusinessUnitIdNil sets the value for BusinessUnitId to be an explicit nil

### UnsetBusinessUnitId
`func (o *BoardTeam) UnsetBusinessUnitId()`

UnsetBusinessUnitId ensures that no value is present for BusinessUnitId, not even an explicit nil
### GetInfo

`func (o *BoardTeam) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BoardTeam) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BoardTeam) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BoardTeam) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


