# TeamMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Board** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**Team** | [**ServiceTeamReference**](ServiceTeamReference.md) |  | 
**Member** | [**MemberReference**](MemberReference.md) |  | 
**TeamLeaderFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTeamMember

`func NewTeamMember(team ServiceTeamReference, member MemberReference, ) *TeamMember`

NewTeamMember instantiates a new TeamMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamMemberWithDefaults

`func NewTeamMemberWithDefaults() *TeamMember`

NewTeamMemberWithDefaults instantiates a new TeamMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamMember) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamMember) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamMember) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TeamMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBoard

`func (o *TeamMember) GetBoard() BoardReference`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *TeamMember) GetBoardOk() (*BoardReference, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *TeamMember) SetBoard(v BoardReference)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *TeamMember) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetTeam

`func (o *TeamMember) GetTeam() ServiceTeamReference`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *TeamMember) GetTeamOk() (*ServiceTeamReference, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *TeamMember) SetTeam(v ServiceTeamReference)`

SetTeam sets Team field to given value.


### GetMember

`func (o *TeamMember) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *TeamMember) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *TeamMember) SetMember(v MemberReference)`

SetMember sets Member field to given value.


### GetTeamLeaderFlag

`func (o *TeamMember) GetTeamLeaderFlag() bool`

GetTeamLeaderFlag returns the TeamLeaderFlag field if non-nil, zero value otherwise.

### GetTeamLeaderFlagOk

`func (o *TeamMember) GetTeamLeaderFlagOk() (*bool, bool)`

GetTeamLeaderFlagOk returns a tuple with the TeamLeaderFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamLeaderFlag

`func (o *TeamMember) SetTeamLeaderFlag(v bool)`

SetTeamLeaderFlag sets TeamLeaderFlag field to given value.

### HasTeamLeaderFlag

`func (o *TeamMember) HasTeamLeaderFlag() bool`

HasTeamLeaderFlag returns a boolean if a field has been set.

### SetTeamLeaderFlagNil

`func (o *TeamMember) SetTeamLeaderFlagNil(b bool)`

 SetTeamLeaderFlagNil sets the value for TeamLeaderFlag to be an explicit nil

### UnsetTeamLeaderFlag
`func (o *TeamMember) UnsetTeamLeaderFlag()`

UnsetTeamLeaderFlag ensures that no value is present for TeamLeaderFlag, not even an explicit nil
### GetInfo

`func (o *TeamMember) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TeamMember) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TeamMember) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TeamMember) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


