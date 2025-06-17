# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**PublicDescription** | Pointer to **string** |  Max length: 255; | [optional] 
**PublicFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewGroup

`func NewGroup(name string, ) *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Group) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Group) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.


### GetPublicDescription

`func (o *Group) GetPublicDescription() string`

GetPublicDescription returns the PublicDescription field if non-nil, zero value otherwise.

### GetPublicDescriptionOk

`func (o *Group) GetPublicDescriptionOk() (*string, bool)`

GetPublicDescriptionOk returns a tuple with the PublicDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicDescription

`func (o *Group) SetPublicDescription(v string)`

SetPublicDescription sets PublicDescription field to given value.

### HasPublicDescription

`func (o *Group) HasPublicDescription() bool`

HasPublicDescription returns a boolean if a field has been set.

### GetPublicFlag

`func (o *Group) GetPublicFlag() bool`

GetPublicFlag returns the PublicFlag field if non-nil, zero value otherwise.

### GetPublicFlagOk

`func (o *Group) GetPublicFlagOk() (*bool, bool)`

GetPublicFlagOk returns a tuple with the PublicFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicFlag

`func (o *Group) SetPublicFlag(v bool)`

SetPublicFlag sets PublicFlag field to given value.

### HasPublicFlag

`func (o *Group) HasPublicFlag() bool`

HasPublicFlag returns a boolean if a field has been set.

### SetPublicFlagNil

`func (o *Group) SetPublicFlagNil(b bool)`

 SetPublicFlagNil sets the value for PublicFlag to be an explicit nil

### UnsetPublicFlag
`func (o *Group) UnsetPublicFlag()`

UnsetPublicFlag ensures that no value is present for PublicFlag, not even an explicit nil
### GetInactiveFlag

`func (o *Group) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *Group) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *Group) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *Group) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *Group) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *Group) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetInfo

`func (o *Group) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Group) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Group) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Group) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


