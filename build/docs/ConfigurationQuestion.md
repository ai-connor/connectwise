# ConfigurationQuestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnswerId** | Pointer to **NullableInt32** |  | [optional] 
**QuestionId** | Pointer to **NullableInt32** |  | [optional] 
**Question** | Pointer to **string** |  | [optional] 
**Answer** | Pointer to **map[string]interface{}** |  | [optional] 
**SequenceNumber** | Pointer to **NullableFloat64** |  | [optional] 
**NumberOfDecimals** | Pointer to **NullableInt32** |  | [optional] 
**FieldType** | Pointer to **NullableString** |  | [optional] 
**RequiredFlag** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewConfigurationQuestion

`func NewConfigurationQuestion() *ConfigurationQuestion`

NewConfigurationQuestion instantiates a new ConfigurationQuestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationQuestionWithDefaults

`func NewConfigurationQuestionWithDefaults() *ConfigurationQuestion`

NewConfigurationQuestionWithDefaults instantiates a new ConfigurationQuestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswerId

`func (o *ConfigurationQuestion) GetAnswerId() int32`

GetAnswerId returns the AnswerId field if non-nil, zero value otherwise.

### GetAnswerIdOk

`func (o *ConfigurationQuestion) GetAnswerIdOk() (*int32, bool)`

GetAnswerIdOk returns a tuple with the AnswerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerId

`func (o *ConfigurationQuestion) SetAnswerId(v int32)`

SetAnswerId sets AnswerId field to given value.

### HasAnswerId

`func (o *ConfigurationQuestion) HasAnswerId() bool`

HasAnswerId returns a boolean if a field has been set.

### SetAnswerIdNil

`func (o *ConfigurationQuestion) SetAnswerIdNil(b bool)`

 SetAnswerIdNil sets the value for AnswerId to be an explicit nil

### UnsetAnswerId
`func (o *ConfigurationQuestion) UnsetAnswerId()`

UnsetAnswerId ensures that no value is present for AnswerId, not even an explicit nil
### GetQuestionId

`func (o *ConfigurationQuestion) GetQuestionId() int32`

GetQuestionId returns the QuestionId field if non-nil, zero value otherwise.

### GetQuestionIdOk

`func (o *ConfigurationQuestion) GetQuestionIdOk() (*int32, bool)`

GetQuestionIdOk returns a tuple with the QuestionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionId

`func (o *ConfigurationQuestion) SetQuestionId(v int32)`

SetQuestionId sets QuestionId field to given value.

### HasQuestionId

`func (o *ConfigurationQuestion) HasQuestionId() bool`

HasQuestionId returns a boolean if a field has been set.

### SetQuestionIdNil

`func (o *ConfigurationQuestion) SetQuestionIdNil(b bool)`

 SetQuestionIdNil sets the value for QuestionId to be an explicit nil

### UnsetQuestionId
`func (o *ConfigurationQuestion) UnsetQuestionId()`

UnsetQuestionId ensures that no value is present for QuestionId, not even an explicit nil
### GetQuestion

`func (o *ConfigurationQuestion) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *ConfigurationQuestion) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *ConfigurationQuestion) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *ConfigurationQuestion) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetAnswer

`func (o *ConfigurationQuestion) GetAnswer() map[string]interface{}`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *ConfigurationQuestion) GetAnswerOk() (*map[string]interface{}, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *ConfigurationQuestion) SetAnswer(v map[string]interface{})`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *ConfigurationQuestion) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *ConfigurationQuestion) GetSequenceNumber() float64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *ConfigurationQuestion) GetSequenceNumberOk() (*float64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *ConfigurationQuestion) SetSequenceNumber(v float64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *ConfigurationQuestion) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### SetSequenceNumberNil

`func (o *ConfigurationQuestion) SetSequenceNumberNil(b bool)`

 SetSequenceNumberNil sets the value for SequenceNumber to be an explicit nil

### UnsetSequenceNumber
`func (o *ConfigurationQuestion) UnsetSequenceNumber()`

UnsetSequenceNumber ensures that no value is present for SequenceNumber, not even an explicit nil
### GetNumberOfDecimals

`func (o *ConfigurationQuestion) GetNumberOfDecimals() int32`

GetNumberOfDecimals returns the NumberOfDecimals field if non-nil, zero value otherwise.

### GetNumberOfDecimalsOk

`func (o *ConfigurationQuestion) GetNumberOfDecimalsOk() (*int32, bool)`

GetNumberOfDecimalsOk returns a tuple with the NumberOfDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDecimals

`func (o *ConfigurationQuestion) SetNumberOfDecimals(v int32)`

SetNumberOfDecimals sets NumberOfDecimals field to given value.

### HasNumberOfDecimals

`func (o *ConfigurationQuestion) HasNumberOfDecimals() bool`

HasNumberOfDecimals returns a boolean if a field has been set.

### SetNumberOfDecimalsNil

`func (o *ConfigurationQuestion) SetNumberOfDecimalsNil(b bool)`

 SetNumberOfDecimalsNil sets the value for NumberOfDecimals to be an explicit nil

### UnsetNumberOfDecimals
`func (o *ConfigurationQuestion) UnsetNumberOfDecimals()`

UnsetNumberOfDecimals ensures that no value is present for NumberOfDecimals, not even an explicit nil
### GetFieldType

`func (o *ConfigurationQuestion) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *ConfigurationQuestion) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *ConfigurationQuestion) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *ConfigurationQuestion) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### SetFieldTypeNil

`func (o *ConfigurationQuestion) SetFieldTypeNil(b bool)`

 SetFieldTypeNil sets the value for FieldType to be an explicit nil

### UnsetFieldType
`func (o *ConfigurationQuestion) UnsetFieldType()`

UnsetFieldType ensures that no value is present for FieldType, not even an explicit nil
### GetRequiredFlag

`func (o *ConfigurationQuestion) GetRequiredFlag() bool`

GetRequiredFlag returns the RequiredFlag field if non-nil, zero value otherwise.

### GetRequiredFlagOk

`func (o *ConfigurationQuestion) GetRequiredFlagOk() (*bool, bool)`

GetRequiredFlagOk returns a tuple with the RequiredFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFlag

`func (o *ConfigurationQuestion) SetRequiredFlag(v bool)`

SetRequiredFlag sets RequiredFlag field to given value.

### HasRequiredFlag

`func (o *ConfigurationQuestion) HasRequiredFlag() bool`

HasRequiredFlag returns a boolean if a field has been set.

### SetRequiredFlagNil

`func (o *ConfigurationQuestion) SetRequiredFlagNil(b bool)`

 SetRequiredFlagNil sets the value for RequiredFlag to be an explicit nil

### UnsetRequiredFlag
`func (o *ConfigurationQuestion) UnsetRequiredFlag()`

UnsetRequiredFlag ensures that no value is present for RequiredFlag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


