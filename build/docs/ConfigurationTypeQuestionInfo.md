# ConfigurationTypeQuestionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ConfigurationType** | Pointer to [**ConfigurationTypeReference**](ConfigurationTypeReference.md) |  | [optional] 
**FieldType** | Pointer to **NullableString** |  | [optional] 
**EntryType** | Pointer to **NullableString** |  | [optional] 
**SequenceNumber** | Pointer to **NullableFloat64** |  | [optional] 
**Question** | Pointer to **string** |  | [optional] 
**NumberOfDecimals** | Pointer to **NullableInt32** |  | [optional] 
**RequiredFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewConfigurationTypeQuestionInfo

`func NewConfigurationTypeQuestionInfo() *ConfigurationTypeQuestionInfo`

NewConfigurationTypeQuestionInfo instantiates a new ConfigurationTypeQuestionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationTypeQuestionInfoWithDefaults

`func NewConfigurationTypeQuestionInfoWithDefaults() *ConfigurationTypeQuestionInfo`

NewConfigurationTypeQuestionInfoWithDefaults instantiates a new ConfigurationTypeQuestionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigurationTypeQuestionInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigurationTypeQuestionInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigurationTypeQuestionInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigurationTypeQuestionInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConfigurationType

`func (o *ConfigurationTypeQuestionInfo) GetConfigurationType() ConfigurationTypeReference`

GetConfigurationType returns the ConfigurationType field if non-nil, zero value otherwise.

### GetConfigurationTypeOk

`func (o *ConfigurationTypeQuestionInfo) GetConfigurationTypeOk() (*ConfigurationTypeReference, bool)`

GetConfigurationTypeOk returns a tuple with the ConfigurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationType

`func (o *ConfigurationTypeQuestionInfo) SetConfigurationType(v ConfigurationTypeReference)`

SetConfigurationType sets ConfigurationType field to given value.

### HasConfigurationType

`func (o *ConfigurationTypeQuestionInfo) HasConfigurationType() bool`

HasConfigurationType returns a boolean if a field has been set.

### GetFieldType

`func (o *ConfigurationTypeQuestionInfo) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *ConfigurationTypeQuestionInfo) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *ConfigurationTypeQuestionInfo) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *ConfigurationTypeQuestionInfo) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### SetFieldTypeNil

`func (o *ConfigurationTypeQuestionInfo) SetFieldTypeNil(b bool)`

 SetFieldTypeNil sets the value for FieldType to be an explicit nil

### UnsetFieldType
`func (o *ConfigurationTypeQuestionInfo) UnsetFieldType()`

UnsetFieldType ensures that no value is present for FieldType, not even an explicit nil
### GetEntryType

`func (o *ConfigurationTypeQuestionInfo) GetEntryType() string`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *ConfigurationTypeQuestionInfo) GetEntryTypeOk() (*string, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *ConfigurationTypeQuestionInfo) SetEntryType(v string)`

SetEntryType sets EntryType field to given value.

### HasEntryType

`func (o *ConfigurationTypeQuestionInfo) HasEntryType() bool`

HasEntryType returns a boolean if a field has been set.

### SetEntryTypeNil

`func (o *ConfigurationTypeQuestionInfo) SetEntryTypeNil(b bool)`

 SetEntryTypeNil sets the value for EntryType to be an explicit nil

### UnsetEntryType
`func (o *ConfigurationTypeQuestionInfo) UnsetEntryType()`

UnsetEntryType ensures that no value is present for EntryType, not even an explicit nil
### GetSequenceNumber

`func (o *ConfigurationTypeQuestionInfo) GetSequenceNumber() float64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *ConfigurationTypeQuestionInfo) GetSequenceNumberOk() (*float64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *ConfigurationTypeQuestionInfo) SetSequenceNumber(v float64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *ConfigurationTypeQuestionInfo) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### SetSequenceNumberNil

`func (o *ConfigurationTypeQuestionInfo) SetSequenceNumberNil(b bool)`

 SetSequenceNumberNil sets the value for SequenceNumber to be an explicit nil

### UnsetSequenceNumber
`func (o *ConfigurationTypeQuestionInfo) UnsetSequenceNumber()`

UnsetSequenceNumber ensures that no value is present for SequenceNumber, not even an explicit nil
### GetQuestion

`func (o *ConfigurationTypeQuestionInfo) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *ConfigurationTypeQuestionInfo) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *ConfigurationTypeQuestionInfo) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *ConfigurationTypeQuestionInfo) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetNumberOfDecimals

`func (o *ConfigurationTypeQuestionInfo) GetNumberOfDecimals() int32`

GetNumberOfDecimals returns the NumberOfDecimals field if non-nil, zero value otherwise.

### GetNumberOfDecimalsOk

`func (o *ConfigurationTypeQuestionInfo) GetNumberOfDecimalsOk() (*int32, bool)`

GetNumberOfDecimalsOk returns a tuple with the NumberOfDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDecimals

`func (o *ConfigurationTypeQuestionInfo) SetNumberOfDecimals(v int32)`

SetNumberOfDecimals sets NumberOfDecimals field to given value.

### HasNumberOfDecimals

`func (o *ConfigurationTypeQuestionInfo) HasNumberOfDecimals() bool`

HasNumberOfDecimals returns a boolean if a field has been set.

### SetNumberOfDecimalsNil

`func (o *ConfigurationTypeQuestionInfo) SetNumberOfDecimalsNil(b bool)`

 SetNumberOfDecimalsNil sets the value for NumberOfDecimals to be an explicit nil

### UnsetNumberOfDecimals
`func (o *ConfigurationTypeQuestionInfo) UnsetNumberOfDecimals()`

UnsetNumberOfDecimals ensures that no value is present for NumberOfDecimals, not even an explicit nil
### GetRequiredFlag

`func (o *ConfigurationTypeQuestionInfo) GetRequiredFlag() bool`

GetRequiredFlag returns the RequiredFlag field if non-nil, zero value otherwise.

### GetRequiredFlagOk

`func (o *ConfigurationTypeQuestionInfo) GetRequiredFlagOk() (*bool, bool)`

GetRequiredFlagOk returns a tuple with the RequiredFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFlag

`func (o *ConfigurationTypeQuestionInfo) SetRequiredFlag(v bool)`

SetRequiredFlag sets RequiredFlag field to given value.

### HasRequiredFlag

`func (o *ConfigurationTypeQuestionInfo) HasRequiredFlag() bool`

HasRequiredFlag returns a boolean if a field has been set.

### SetRequiredFlagNil

`func (o *ConfigurationTypeQuestionInfo) SetRequiredFlagNil(b bool)`

 SetRequiredFlagNil sets the value for RequiredFlag to be an explicit nil

### UnsetRequiredFlag
`func (o *ConfigurationTypeQuestionInfo) UnsetRequiredFlag()`

UnsetRequiredFlag ensures that no value is present for RequiredFlag, not even an explicit nil
### GetInactiveFlag

`func (o *ConfigurationTypeQuestionInfo) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *ConfigurationTypeQuestionInfo) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *ConfigurationTypeQuestionInfo) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *ConfigurationTypeQuestionInfo) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *ConfigurationTypeQuestionInfo) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *ConfigurationTypeQuestionInfo) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetInfo

`func (o *ConfigurationTypeQuestionInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ConfigurationTypeQuestionInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ConfigurationTypeQuestionInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ConfigurationTypeQuestionInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


