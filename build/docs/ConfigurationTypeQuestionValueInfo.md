# ConfigurationTypeQuestionValueInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ConfigurationType** | Pointer to [**ConfigurationTypeReference**](ConfigurationTypeReference.md) |  | [optional] 
**Question** | Pointer to [**ConfigurationTypeQuestionReference**](ConfigurationTypeQuestionReference.md) |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewConfigurationTypeQuestionValueInfo

`func NewConfigurationTypeQuestionValueInfo() *ConfigurationTypeQuestionValueInfo`

NewConfigurationTypeQuestionValueInfo instantiates a new ConfigurationTypeQuestionValueInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationTypeQuestionValueInfoWithDefaults

`func NewConfigurationTypeQuestionValueInfoWithDefaults() *ConfigurationTypeQuestionValueInfo`

NewConfigurationTypeQuestionValueInfoWithDefaults instantiates a new ConfigurationTypeQuestionValueInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigurationTypeQuestionValueInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigurationTypeQuestionValueInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigurationTypeQuestionValueInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigurationTypeQuestionValueInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConfigurationType

`func (o *ConfigurationTypeQuestionValueInfo) GetConfigurationType() ConfigurationTypeReference`

GetConfigurationType returns the ConfigurationType field if non-nil, zero value otherwise.

### GetConfigurationTypeOk

`func (o *ConfigurationTypeQuestionValueInfo) GetConfigurationTypeOk() (*ConfigurationTypeReference, bool)`

GetConfigurationTypeOk returns a tuple with the ConfigurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationType

`func (o *ConfigurationTypeQuestionValueInfo) SetConfigurationType(v ConfigurationTypeReference)`

SetConfigurationType sets ConfigurationType field to given value.

### HasConfigurationType

`func (o *ConfigurationTypeQuestionValueInfo) HasConfigurationType() bool`

HasConfigurationType returns a boolean if a field has been set.

### GetQuestion

`func (o *ConfigurationTypeQuestionValueInfo) GetQuestion() ConfigurationTypeQuestionReference`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *ConfigurationTypeQuestionValueInfo) GetQuestionOk() (*ConfigurationTypeQuestionReference, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *ConfigurationTypeQuestionValueInfo) SetQuestion(v ConfigurationTypeQuestionReference)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *ConfigurationTypeQuestionValueInfo) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetValue

`func (o *ConfigurationTypeQuestionValueInfo) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConfigurationTypeQuestionValueInfo) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConfigurationTypeQuestionValueInfo) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ConfigurationTypeQuestionValueInfo) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDefaultFlag

`func (o *ConfigurationTypeQuestionValueInfo) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *ConfigurationTypeQuestionValueInfo) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *ConfigurationTypeQuestionValueInfo) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *ConfigurationTypeQuestionValueInfo) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *ConfigurationTypeQuestionValueInfo) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *ConfigurationTypeQuestionValueInfo) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetInactiveFlag

`func (o *ConfigurationTypeQuestionValueInfo) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *ConfigurationTypeQuestionValueInfo) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *ConfigurationTypeQuestionValueInfo) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *ConfigurationTypeQuestionValueInfo) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *ConfigurationTypeQuestionValueInfo) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *ConfigurationTypeQuestionValueInfo) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetInfo

`func (o *ConfigurationTypeQuestionValueInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ConfigurationTypeQuestionValueInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ConfigurationTypeQuestionValueInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ConfigurationTypeQuestionValueInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


