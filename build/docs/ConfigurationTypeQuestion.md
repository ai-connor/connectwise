# ConfigurationTypeQuestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ConfigurationType** | Pointer to [**ConfigurationTypeReference**](ConfigurationTypeReference.md) |  | [optional] 
**FieldType** | **NullableString** |  | 
**EntryType** | **NullableString** |  | 
**SequenceNumber** | **NullableFloat64** |  | 
**Question** | **string** |  Max length: 1000; | 
**NumberOfDecimals** | Pointer to **NullableInt32** |  | [optional] 
**RequiredFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewConfigurationTypeQuestion

`func NewConfigurationTypeQuestion(fieldType NullableString, entryType NullableString, sequenceNumber NullableFloat64, question string, ) *ConfigurationTypeQuestion`

NewConfigurationTypeQuestion instantiates a new ConfigurationTypeQuestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationTypeQuestionWithDefaults

`func NewConfigurationTypeQuestionWithDefaults() *ConfigurationTypeQuestion`

NewConfigurationTypeQuestionWithDefaults instantiates a new ConfigurationTypeQuestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigurationTypeQuestion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigurationTypeQuestion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigurationTypeQuestion) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigurationTypeQuestion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConfigurationType

`func (o *ConfigurationTypeQuestion) GetConfigurationType() ConfigurationTypeReference`

GetConfigurationType returns the ConfigurationType field if non-nil, zero value otherwise.

### GetConfigurationTypeOk

`func (o *ConfigurationTypeQuestion) GetConfigurationTypeOk() (*ConfigurationTypeReference, bool)`

GetConfigurationTypeOk returns a tuple with the ConfigurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationType

`func (o *ConfigurationTypeQuestion) SetConfigurationType(v ConfigurationTypeReference)`

SetConfigurationType sets ConfigurationType field to given value.

### HasConfigurationType

`func (o *ConfigurationTypeQuestion) HasConfigurationType() bool`

HasConfigurationType returns a boolean if a field has been set.

### GetFieldType

`func (o *ConfigurationTypeQuestion) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *ConfigurationTypeQuestion) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *ConfigurationTypeQuestion) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.


### SetFieldTypeNil

`func (o *ConfigurationTypeQuestion) SetFieldTypeNil(b bool)`

 SetFieldTypeNil sets the value for FieldType to be an explicit nil

### UnsetFieldType
`func (o *ConfigurationTypeQuestion) UnsetFieldType()`

UnsetFieldType ensures that no value is present for FieldType, not even an explicit nil
### GetEntryType

`func (o *ConfigurationTypeQuestion) GetEntryType() string`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *ConfigurationTypeQuestion) GetEntryTypeOk() (*string, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *ConfigurationTypeQuestion) SetEntryType(v string)`

SetEntryType sets EntryType field to given value.


### SetEntryTypeNil

`func (o *ConfigurationTypeQuestion) SetEntryTypeNil(b bool)`

 SetEntryTypeNil sets the value for EntryType to be an explicit nil

### UnsetEntryType
`func (o *ConfigurationTypeQuestion) UnsetEntryType()`

UnsetEntryType ensures that no value is present for EntryType, not even an explicit nil
### GetSequenceNumber

`func (o *ConfigurationTypeQuestion) GetSequenceNumber() float64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *ConfigurationTypeQuestion) GetSequenceNumberOk() (*float64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *ConfigurationTypeQuestion) SetSequenceNumber(v float64)`

SetSequenceNumber sets SequenceNumber field to given value.


### SetSequenceNumberNil

`func (o *ConfigurationTypeQuestion) SetSequenceNumberNil(b bool)`

 SetSequenceNumberNil sets the value for SequenceNumber to be an explicit nil

### UnsetSequenceNumber
`func (o *ConfigurationTypeQuestion) UnsetSequenceNumber()`

UnsetSequenceNumber ensures that no value is present for SequenceNumber, not even an explicit nil
### GetQuestion

`func (o *ConfigurationTypeQuestion) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *ConfigurationTypeQuestion) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *ConfigurationTypeQuestion) SetQuestion(v string)`

SetQuestion sets Question field to given value.


### GetNumberOfDecimals

`func (o *ConfigurationTypeQuestion) GetNumberOfDecimals() int32`

GetNumberOfDecimals returns the NumberOfDecimals field if non-nil, zero value otherwise.

### GetNumberOfDecimalsOk

`func (o *ConfigurationTypeQuestion) GetNumberOfDecimalsOk() (*int32, bool)`

GetNumberOfDecimalsOk returns a tuple with the NumberOfDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDecimals

`func (o *ConfigurationTypeQuestion) SetNumberOfDecimals(v int32)`

SetNumberOfDecimals sets NumberOfDecimals field to given value.

### HasNumberOfDecimals

`func (o *ConfigurationTypeQuestion) HasNumberOfDecimals() bool`

HasNumberOfDecimals returns a boolean if a field has been set.

### SetNumberOfDecimalsNil

`func (o *ConfigurationTypeQuestion) SetNumberOfDecimalsNil(b bool)`

 SetNumberOfDecimalsNil sets the value for NumberOfDecimals to be an explicit nil

### UnsetNumberOfDecimals
`func (o *ConfigurationTypeQuestion) UnsetNumberOfDecimals()`

UnsetNumberOfDecimals ensures that no value is present for NumberOfDecimals, not even an explicit nil
### GetRequiredFlag

`func (o *ConfigurationTypeQuestion) GetRequiredFlag() bool`

GetRequiredFlag returns the RequiredFlag field if non-nil, zero value otherwise.

### GetRequiredFlagOk

`func (o *ConfigurationTypeQuestion) GetRequiredFlagOk() (*bool, bool)`

GetRequiredFlagOk returns a tuple with the RequiredFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFlag

`func (o *ConfigurationTypeQuestion) SetRequiredFlag(v bool)`

SetRequiredFlag sets RequiredFlag field to given value.

### HasRequiredFlag

`func (o *ConfigurationTypeQuestion) HasRequiredFlag() bool`

HasRequiredFlag returns a boolean if a field has been set.

### SetRequiredFlagNil

`func (o *ConfigurationTypeQuestion) SetRequiredFlagNil(b bool)`

 SetRequiredFlagNil sets the value for RequiredFlag to be an explicit nil

### UnsetRequiredFlag
`func (o *ConfigurationTypeQuestion) UnsetRequiredFlag()`

UnsetRequiredFlag ensures that no value is present for RequiredFlag, not even an explicit nil
### GetInactiveFlag

`func (o *ConfigurationTypeQuestion) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *ConfigurationTypeQuestion) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *ConfigurationTypeQuestion) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *ConfigurationTypeQuestion) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *ConfigurationTypeQuestion) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *ConfigurationTypeQuestion) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetInfo

`func (o *ConfigurationTypeQuestion) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ConfigurationTypeQuestion) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ConfigurationTypeQuestion) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ConfigurationTypeQuestion) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


