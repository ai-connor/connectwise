# ConfigurationTypeQuestionReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Question** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewConfigurationTypeQuestionReference

`func NewConfigurationTypeQuestionReference() *ConfigurationTypeQuestionReference`

NewConfigurationTypeQuestionReference instantiates a new ConfigurationTypeQuestionReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationTypeQuestionReferenceWithDefaults

`func NewConfigurationTypeQuestionReferenceWithDefaults() *ConfigurationTypeQuestionReference`

NewConfigurationTypeQuestionReferenceWithDefaults instantiates a new ConfigurationTypeQuestionReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigurationTypeQuestionReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigurationTypeQuestionReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigurationTypeQuestionReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigurationTypeQuestionReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ConfigurationTypeQuestionReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ConfigurationTypeQuestionReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetQuestion

`func (o *ConfigurationTypeQuestionReference) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *ConfigurationTypeQuestionReference) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *ConfigurationTypeQuestionReference) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *ConfigurationTypeQuestionReference) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetInfo

`func (o *ConfigurationTypeQuestionReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ConfigurationTypeQuestionReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ConfigurationTypeQuestionReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ConfigurationTypeQuestionReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


