# ConfigurationTypeQuestionValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ConfigurationType** | Pointer to [**ConfigurationTypeReference**](ConfigurationTypeReference.md) |  | [optional] 
**Question** | Pointer to [**ConfigurationTypeQuestionReference**](ConfigurationTypeQuestionReference.md) |  | [optional] 
**Value** | **string** |  Max length: 1000; | 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewConfigurationTypeQuestionValue

`func NewConfigurationTypeQuestionValue(value string, ) *ConfigurationTypeQuestionValue`

NewConfigurationTypeQuestionValue instantiates a new ConfigurationTypeQuestionValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationTypeQuestionValueWithDefaults

`func NewConfigurationTypeQuestionValueWithDefaults() *ConfigurationTypeQuestionValue`

NewConfigurationTypeQuestionValueWithDefaults instantiates a new ConfigurationTypeQuestionValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigurationTypeQuestionValue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigurationTypeQuestionValue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigurationTypeQuestionValue) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigurationTypeQuestionValue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConfigurationType

`func (o *ConfigurationTypeQuestionValue) GetConfigurationType() ConfigurationTypeReference`

GetConfigurationType returns the ConfigurationType field if non-nil, zero value otherwise.

### GetConfigurationTypeOk

`func (o *ConfigurationTypeQuestionValue) GetConfigurationTypeOk() (*ConfigurationTypeReference, bool)`

GetConfigurationTypeOk returns a tuple with the ConfigurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationType

`func (o *ConfigurationTypeQuestionValue) SetConfigurationType(v ConfigurationTypeReference)`

SetConfigurationType sets ConfigurationType field to given value.

### HasConfigurationType

`func (o *ConfigurationTypeQuestionValue) HasConfigurationType() bool`

HasConfigurationType returns a boolean if a field has been set.

### GetQuestion

`func (o *ConfigurationTypeQuestionValue) GetQuestion() ConfigurationTypeQuestionReference`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *ConfigurationTypeQuestionValue) GetQuestionOk() (*ConfigurationTypeQuestionReference, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *ConfigurationTypeQuestionValue) SetQuestion(v ConfigurationTypeQuestionReference)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *ConfigurationTypeQuestionValue) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetValue

`func (o *ConfigurationTypeQuestionValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConfigurationTypeQuestionValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConfigurationTypeQuestionValue) SetValue(v string)`

SetValue sets Value field to given value.


### GetDefaultFlag

`func (o *ConfigurationTypeQuestionValue) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *ConfigurationTypeQuestionValue) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *ConfigurationTypeQuestionValue) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *ConfigurationTypeQuestionValue) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *ConfigurationTypeQuestionValue) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *ConfigurationTypeQuestionValue) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetInactiveFlag

`func (o *ConfigurationTypeQuestionValue) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *ConfigurationTypeQuestionValue) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *ConfigurationTypeQuestionValue) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *ConfigurationTypeQuestionValue) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *ConfigurationTypeQuestionValue) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *ConfigurationTypeQuestionValue) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetInfo

`func (o *ConfigurationTypeQuestionValue) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ConfigurationTypeQuestionValue) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ConfigurationTypeQuestionValue) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ConfigurationTypeQuestionValue) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


