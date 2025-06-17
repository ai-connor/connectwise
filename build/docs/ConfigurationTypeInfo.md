# ConfigurationTypeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**SystemFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewConfigurationTypeInfo

`func NewConfigurationTypeInfo() *ConfigurationTypeInfo`

NewConfigurationTypeInfo instantiates a new ConfigurationTypeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationTypeInfoWithDefaults

`func NewConfigurationTypeInfoWithDefaults() *ConfigurationTypeInfo`

NewConfigurationTypeInfoWithDefaults instantiates a new ConfigurationTypeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigurationTypeInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigurationTypeInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigurationTypeInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigurationTypeInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ConfigurationTypeInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigurationTypeInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigurationTypeInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigurationTypeInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInactiveFlag

`func (o *ConfigurationTypeInfo) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *ConfigurationTypeInfo) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *ConfigurationTypeInfo) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *ConfigurationTypeInfo) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *ConfigurationTypeInfo) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *ConfigurationTypeInfo) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetSystemFlag

`func (o *ConfigurationTypeInfo) GetSystemFlag() bool`

GetSystemFlag returns the SystemFlag field if non-nil, zero value otherwise.

### GetSystemFlagOk

`func (o *ConfigurationTypeInfo) GetSystemFlagOk() (*bool, bool)`

GetSystemFlagOk returns a tuple with the SystemFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemFlag

`func (o *ConfigurationTypeInfo) SetSystemFlag(v bool)`

SetSystemFlag sets SystemFlag field to given value.

### HasSystemFlag

`func (o *ConfigurationTypeInfo) HasSystemFlag() bool`

HasSystemFlag returns a boolean if a field has been set.

### SetSystemFlagNil

`func (o *ConfigurationTypeInfo) SetSystemFlagNil(b bool)`

 SetSystemFlagNil sets the value for SystemFlag to be an explicit nil

### UnsetSystemFlag
`func (o *ConfigurationTypeInfo) UnsetSystemFlag()`

UnsetSystemFlag ensures that no value is present for SystemFlag, not even an explicit nil
### GetInfo

`func (o *ConfigurationTypeInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ConfigurationTypeInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ConfigurationTypeInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ConfigurationTypeInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


