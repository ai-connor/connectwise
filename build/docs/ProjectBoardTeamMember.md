# ProjectBoardTeamMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Member** | [**MemberReference**](MemberReference.md) |  | 
**ProjectRole** | [**ProjectRoleReference**](ProjectRoleReference.md) |  | 
**WorkRole** | Pointer to [**WorkRoleReference**](WorkRoleReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewProjectBoardTeamMember

`func NewProjectBoardTeamMember(member MemberReference, projectRole ProjectRoleReference, ) *ProjectBoardTeamMember`

NewProjectBoardTeamMember instantiates a new ProjectBoardTeamMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectBoardTeamMemberWithDefaults

`func NewProjectBoardTeamMemberWithDefaults() *ProjectBoardTeamMember`

NewProjectBoardTeamMemberWithDefaults instantiates a new ProjectBoardTeamMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectBoardTeamMember) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectBoardTeamMember) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectBoardTeamMember) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectBoardTeamMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMember

`func (o *ProjectBoardTeamMember) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ProjectBoardTeamMember) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ProjectBoardTeamMember) SetMember(v MemberReference)`

SetMember sets Member field to given value.


### GetProjectRole

`func (o *ProjectBoardTeamMember) GetProjectRole() ProjectRoleReference`

GetProjectRole returns the ProjectRole field if non-nil, zero value otherwise.

### GetProjectRoleOk

`func (o *ProjectBoardTeamMember) GetProjectRoleOk() (*ProjectRoleReference, bool)`

GetProjectRoleOk returns a tuple with the ProjectRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectRole

`func (o *ProjectBoardTeamMember) SetProjectRole(v ProjectRoleReference)`

SetProjectRole sets ProjectRole field to given value.


### GetWorkRole

`func (o *ProjectBoardTeamMember) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *ProjectBoardTeamMember) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *ProjectBoardTeamMember) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *ProjectBoardTeamMember) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetInfo

`func (o *ProjectBoardTeamMember) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProjectBoardTeamMember) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProjectBoardTeamMember) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProjectBoardTeamMember) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


