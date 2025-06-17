# ProjectTeamMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **NullableInt32** |  | [optional] 
**Hours** | Pointer to **NullableFloat64** |  | [optional] 
**Member** | [**MemberReference**](MemberReference.md) |  | 
**ProjectRole** | [**ProjectRoleReference**](ProjectRoleReference.md) |  | 
**WorkRole** | Pointer to [**WorkRoleReference**](WorkRoleReference.md) |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewProjectTeamMember

`func NewProjectTeamMember(member MemberReference, projectRole ProjectRoleReference, ) *ProjectTeamMember`

NewProjectTeamMember instantiates a new ProjectTeamMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectTeamMemberWithDefaults

`func NewProjectTeamMemberWithDefaults() *ProjectTeamMember`

NewProjectTeamMemberWithDefaults instantiates a new ProjectTeamMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectTeamMember) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectTeamMember) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectTeamMember) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectTeamMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectId

`func (o *ProjectTeamMember) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectTeamMember) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectTeamMember) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProjectTeamMember) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *ProjectTeamMember) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *ProjectTeamMember) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetHours

`func (o *ProjectTeamMember) GetHours() float64`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *ProjectTeamMember) GetHoursOk() (*float64, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *ProjectTeamMember) SetHours(v float64)`

SetHours sets Hours field to given value.

### HasHours

`func (o *ProjectTeamMember) HasHours() bool`

HasHours returns a boolean if a field has been set.

### SetHoursNil

`func (o *ProjectTeamMember) SetHoursNil(b bool)`

 SetHoursNil sets the value for Hours to be an explicit nil

### UnsetHours
`func (o *ProjectTeamMember) UnsetHours()`

UnsetHours ensures that no value is present for Hours, not even an explicit nil
### GetMember

`func (o *ProjectTeamMember) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ProjectTeamMember) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ProjectTeamMember) SetMember(v MemberReference)`

SetMember sets Member field to given value.


### GetProjectRole

`func (o *ProjectTeamMember) GetProjectRole() ProjectRoleReference`

GetProjectRole returns the ProjectRole field if non-nil, zero value otherwise.

### GetProjectRoleOk

`func (o *ProjectTeamMember) GetProjectRoleOk() (*ProjectRoleReference, bool)`

GetProjectRoleOk returns a tuple with the ProjectRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectRole

`func (o *ProjectTeamMember) SetProjectRole(v ProjectRoleReference)`

SetProjectRole sets ProjectRole field to given value.


### GetWorkRole

`func (o *ProjectTeamMember) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *ProjectTeamMember) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *ProjectTeamMember) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *ProjectTeamMember) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetStartDate

`func (o *ProjectTeamMember) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ProjectTeamMember) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ProjectTeamMember) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ProjectTeamMember) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ProjectTeamMember) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ProjectTeamMember) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ProjectTeamMember) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ProjectTeamMember) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetInfo

`func (o *ProjectTeamMember) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProjectTeamMember) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProjectTeamMember) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProjectTeamMember) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


