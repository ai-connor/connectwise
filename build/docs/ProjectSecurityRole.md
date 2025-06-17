# ProjectSecurityRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 30; | 
**ManagerRoleFlag** | Pointer to **NullableBool** |  | [optional] 
**DefaultContactFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewProjectSecurityRole

`func NewProjectSecurityRole(name string, ) *ProjectSecurityRole`

NewProjectSecurityRole instantiates a new ProjectSecurityRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectSecurityRoleWithDefaults

`func NewProjectSecurityRoleWithDefaults() *ProjectSecurityRole`

NewProjectSecurityRoleWithDefaults instantiates a new ProjectSecurityRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectSecurityRole) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectSecurityRole) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectSecurityRole) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectSecurityRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProjectSecurityRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectSecurityRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectSecurityRole) SetName(v string)`

SetName sets Name field to given value.


### GetManagerRoleFlag

`func (o *ProjectSecurityRole) GetManagerRoleFlag() bool`

GetManagerRoleFlag returns the ManagerRoleFlag field if non-nil, zero value otherwise.

### GetManagerRoleFlagOk

`func (o *ProjectSecurityRole) GetManagerRoleFlagOk() (*bool, bool)`

GetManagerRoleFlagOk returns a tuple with the ManagerRoleFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerRoleFlag

`func (o *ProjectSecurityRole) SetManagerRoleFlag(v bool)`

SetManagerRoleFlag sets ManagerRoleFlag field to given value.

### HasManagerRoleFlag

`func (o *ProjectSecurityRole) HasManagerRoleFlag() bool`

HasManagerRoleFlag returns a boolean if a field has been set.

### SetManagerRoleFlagNil

`func (o *ProjectSecurityRole) SetManagerRoleFlagNil(b bool)`

 SetManagerRoleFlagNil sets the value for ManagerRoleFlag to be an explicit nil

### UnsetManagerRoleFlag
`func (o *ProjectSecurityRole) UnsetManagerRoleFlag()`

UnsetManagerRoleFlag ensures that no value is present for ManagerRoleFlag, not even an explicit nil
### GetDefaultContactFlag

`func (o *ProjectSecurityRole) GetDefaultContactFlag() bool`

GetDefaultContactFlag returns the DefaultContactFlag field if non-nil, zero value otherwise.

### GetDefaultContactFlagOk

`func (o *ProjectSecurityRole) GetDefaultContactFlagOk() (*bool, bool)`

GetDefaultContactFlagOk returns a tuple with the DefaultContactFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultContactFlag

`func (o *ProjectSecurityRole) SetDefaultContactFlag(v bool)`

SetDefaultContactFlag sets DefaultContactFlag field to given value.

### HasDefaultContactFlag

`func (o *ProjectSecurityRole) HasDefaultContactFlag() bool`

HasDefaultContactFlag returns a boolean if a field has been set.

### SetDefaultContactFlagNil

`func (o *ProjectSecurityRole) SetDefaultContactFlagNil(b bool)`

 SetDefaultContactFlagNil sets the value for DefaultContactFlag to be an explicit nil

### UnsetDefaultContactFlag
`func (o *ProjectSecurityRole) UnsetDefaultContactFlag()`

UnsetDefaultContactFlag ensures that no value is present for DefaultContactFlag, not even an explicit nil
### GetInfo

`func (o *ProjectSecurityRole) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProjectSecurityRole) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProjectSecurityRole) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProjectSecurityRole) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


