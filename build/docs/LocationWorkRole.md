# LocationWorkRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**WorkRole** | Pointer to [**WorkRoleReference**](WorkRoleReference.md) |  | [optional] 
**WorkRoleInactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewLocationWorkRole

`func NewLocationWorkRole() *LocationWorkRole`

NewLocationWorkRole instantiates a new LocationWorkRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWorkRoleWithDefaults

`func NewLocationWorkRoleWithDefaults() *LocationWorkRole`

NewLocationWorkRoleWithDefaults instantiates a new LocationWorkRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LocationWorkRole) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocationWorkRole) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocationWorkRole) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LocationWorkRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *LocationWorkRole) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *LocationWorkRole) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *LocationWorkRole) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *LocationWorkRole) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetWorkRole

`func (o *LocationWorkRole) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *LocationWorkRole) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *LocationWorkRole) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *LocationWorkRole) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetWorkRoleInactiveFlag

`func (o *LocationWorkRole) GetWorkRoleInactiveFlag() bool`

GetWorkRoleInactiveFlag returns the WorkRoleInactiveFlag field if non-nil, zero value otherwise.

### GetWorkRoleInactiveFlagOk

`func (o *LocationWorkRole) GetWorkRoleInactiveFlagOk() (*bool, bool)`

GetWorkRoleInactiveFlagOk returns a tuple with the WorkRoleInactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRoleInactiveFlag

`func (o *LocationWorkRole) SetWorkRoleInactiveFlag(v bool)`

SetWorkRoleInactiveFlag sets WorkRoleInactiveFlag field to given value.

### HasWorkRoleInactiveFlag

`func (o *LocationWorkRole) HasWorkRoleInactiveFlag() bool`

HasWorkRoleInactiveFlag returns a boolean if a field has been set.

### SetWorkRoleInactiveFlagNil

`func (o *LocationWorkRole) SetWorkRoleInactiveFlagNil(b bool)`

 SetWorkRoleInactiveFlagNil sets the value for WorkRoleInactiveFlag to be an explicit nil

### UnsetWorkRoleInactiveFlag
`func (o *LocationWorkRole) UnsetWorkRoleInactiveFlag()`

UnsetWorkRoleInactiveFlag ensures that no value is present for WorkRoleInactiveFlag, not even an explicit nil
### GetInfo

`func (o *LocationWorkRole) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *LocationWorkRole) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *LocationWorkRole) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *LocationWorkRole) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


