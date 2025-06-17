# WorkTypeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**ActivityDefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewWorkTypeInfo

`func NewWorkTypeInfo() *WorkTypeInfo`

NewWorkTypeInfo instantiates a new WorkTypeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkTypeInfoWithDefaults

`func NewWorkTypeInfoWithDefaults() *WorkTypeInfo`

NewWorkTypeInfoWithDefaults instantiates a new WorkTypeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkTypeInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkTypeInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkTypeInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WorkTypeInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WorkTypeInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkTypeInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkTypeInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkTypeInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultFlag

`func (o *WorkTypeInfo) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *WorkTypeInfo) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *WorkTypeInfo) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *WorkTypeInfo) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *WorkTypeInfo) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *WorkTypeInfo) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetInactiveFlag

`func (o *WorkTypeInfo) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *WorkTypeInfo) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *WorkTypeInfo) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *WorkTypeInfo) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *WorkTypeInfo) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *WorkTypeInfo) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetActivityDefaultFlag

`func (o *WorkTypeInfo) GetActivityDefaultFlag() bool`

GetActivityDefaultFlag returns the ActivityDefaultFlag field if non-nil, zero value otherwise.

### GetActivityDefaultFlagOk

`func (o *WorkTypeInfo) GetActivityDefaultFlagOk() (*bool, bool)`

GetActivityDefaultFlagOk returns a tuple with the ActivityDefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityDefaultFlag

`func (o *WorkTypeInfo) SetActivityDefaultFlag(v bool)`

SetActivityDefaultFlag sets ActivityDefaultFlag field to given value.

### HasActivityDefaultFlag

`func (o *WorkTypeInfo) HasActivityDefaultFlag() bool`

HasActivityDefaultFlag returns a boolean if a field has been set.

### SetActivityDefaultFlagNil

`func (o *WorkTypeInfo) SetActivityDefaultFlagNil(b bool)`

 SetActivityDefaultFlagNil sets the value for ActivityDefaultFlag to be an explicit nil

### UnsetActivityDefaultFlag
`func (o *WorkTypeInfo) UnsetActivityDefaultFlag()`

UnsetActivityDefaultFlag ensures that no value is present for ActivityDefaultFlag, not even an explicit nil
### GetInfo

`func (o *WorkTypeInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *WorkTypeInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *WorkTypeInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *WorkTypeInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


