# ConfigurationType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**SystemFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewConfigurationType

`func NewConfigurationType(name string, ) *ConfigurationType`

NewConfigurationType instantiates a new ConfigurationType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationTypeWithDefaults

`func NewConfigurationTypeWithDefaults() *ConfigurationType`

NewConfigurationTypeWithDefaults instantiates a new ConfigurationType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigurationType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigurationType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigurationType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigurationType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ConfigurationType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigurationType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigurationType) SetName(v string)`

SetName sets Name field to given value.


### GetInactiveFlag

`func (o *ConfigurationType) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *ConfigurationType) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *ConfigurationType) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *ConfigurationType) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *ConfigurationType) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *ConfigurationType) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetSystemFlag

`func (o *ConfigurationType) GetSystemFlag() bool`

GetSystemFlag returns the SystemFlag field if non-nil, zero value otherwise.

### GetSystemFlagOk

`func (o *ConfigurationType) GetSystemFlagOk() (*bool, bool)`

GetSystemFlagOk returns a tuple with the SystemFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemFlag

`func (o *ConfigurationType) SetSystemFlag(v bool)`

SetSystemFlag sets SystemFlag field to given value.

### HasSystemFlag

`func (o *ConfigurationType) HasSystemFlag() bool`

HasSystemFlag returns a boolean if a field has been set.

### SetSystemFlagNil

`func (o *ConfigurationType) SetSystemFlagNil(b bool)`

 SetSystemFlagNil sets the value for SystemFlag to be an explicit nil

### UnsetSystemFlag
`func (o *ConfigurationType) UnsetSystemFlag()`

UnsetSystemFlag ensures that no value is present for SystemFlag, not even an explicit nil
### GetInfo

`func (o *ConfigurationType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ConfigurationType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ConfigurationType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ConfigurationType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


