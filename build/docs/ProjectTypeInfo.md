# ProjectTypeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewProjectTypeInfo

`func NewProjectTypeInfo() *ProjectTypeInfo`

NewProjectTypeInfo instantiates a new ProjectTypeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectTypeInfoWithDefaults

`func NewProjectTypeInfoWithDefaults() *ProjectTypeInfo`

NewProjectTypeInfoWithDefaults instantiates a new ProjectTypeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectTypeInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectTypeInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectTypeInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectTypeInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProjectTypeInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectTypeInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectTypeInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectTypeInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultFlag

`func (o *ProjectTypeInfo) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *ProjectTypeInfo) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *ProjectTypeInfo) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *ProjectTypeInfo) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *ProjectTypeInfo) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *ProjectTypeInfo) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetInactiveFlag

`func (o *ProjectTypeInfo) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *ProjectTypeInfo) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *ProjectTypeInfo) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *ProjectTypeInfo) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *ProjectTypeInfo) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *ProjectTypeInfo) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetInfo

`func (o *ProjectTypeInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProjectTypeInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProjectTypeInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProjectTypeInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


