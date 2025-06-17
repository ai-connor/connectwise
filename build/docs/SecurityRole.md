# SecurityRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 30; | 
**RoleType** | Pointer to **string** |  Max length: 30; | [optional] 
**AdminFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSecurityRole

`func NewSecurityRole(name string, ) *SecurityRole`

NewSecurityRole instantiates a new SecurityRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityRoleWithDefaults

`func NewSecurityRoleWithDefaults() *SecurityRole`

NewSecurityRoleWithDefaults instantiates a new SecurityRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityRole) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityRole) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityRole) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SecurityRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityRole) SetName(v string)`

SetName sets Name field to given value.


### GetRoleType

`func (o *SecurityRole) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *SecurityRole) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *SecurityRole) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *SecurityRole) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetAdminFlag

`func (o *SecurityRole) GetAdminFlag() bool`

GetAdminFlag returns the AdminFlag field if non-nil, zero value otherwise.

### GetAdminFlagOk

`func (o *SecurityRole) GetAdminFlagOk() (*bool, bool)`

GetAdminFlagOk returns a tuple with the AdminFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFlag

`func (o *SecurityRole) SetAdminFlag(v bool)`

SetAdminFlag sets AdminFlag field to given value.

### HasAdminFlag

`func (o *SecurityRole) HasAdminFlag() bool`

HasAdminFlag returns a boolean if a field has been set.

### SetAdminFlagNil

`func (o *SecurityRole) SetAdminFlagNil(b bool)`

 SetAdminFlagNil sets the value for AdminFlag to be an explicit nil

### UnsetAdminFlag
`func (o *SecurityRole) UnsetAdminFlag()`

UnsetAdminFlag ensures that no value is present for AdminFlag, not even an explicit nil
### GetInactiveFlag

`func (o *SecurityRole) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *SecurityRole) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *SecurityRole) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *SecurityRole) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *SecurityRole) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *SecurityRole) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetInfo

`func (o *SecurityRole) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SecurityRole) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SecurityRole) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SecurityRole) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


