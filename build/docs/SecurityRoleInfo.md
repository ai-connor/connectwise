# SecurityRoleInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RoleType** | Pointer to **string** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSecurityRoleInfo

`func NewSecurityRoleInfo() *SecurityRoleInfo`

NewSecurityRoleInfo instantiates a new SecurityRoleInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityRoleInfoWithDefaults

`func NewSecurityRoleInfoWithDefaults() *SecurityRoleInfo`

NewSecurityRoleInfoWithDefaults instantiates a new SecurityRoleInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityRoleInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityRoleInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityRoleInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityRoleInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SecurityRoleInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityRoleInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityRoleInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityRoleInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoleType

`func (o *SecurityRoleInfo) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *SecurityRoleInfo) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *SecurityRoleInfo) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *SecurityRoleInfo) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetInactiveFlag

`func (o *SecurityRoleInfo) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *SecurityRoleInfo) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *SecurityRoleInfo) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *SecurityRoleInfo) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *SecurityRoleInfo) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *SecurityRoleInfo) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetInfo

`func (o *SecurityRoleInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SecurityRoleInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SecurityRoleInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SecurityRoleInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


