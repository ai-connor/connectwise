# CategoryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCategoryInfo

`func NewCategoryInfo() *CategoryInfo`

NewCategoryInfo instantiates a new CategoryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryInfoWithDefaults

`func NewCategoryInfoWithDefaults() *CategoryInfo`

NewCategoryInfoWithDefaults instantiates a new CategoryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CategoryInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CategoryInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CategoryInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CategoryInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CategoryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CategoryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CategoryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CategoryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInactiveFlag

`func (o *CategoryInfo) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *CategoryInfo) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *CategoryInfo) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *CategoryInfo) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *CategoryInfo) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *CategoryInfo) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetDefaultFlag

`func (o *CategoryInfo) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *CategoryInfo) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *CategoryInfo) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *CategoryInfo) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *CategoryInfo) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *CategoryInfo) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetInfo

`func (o *CategoryInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CategoryInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CategoryInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CategoryInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


