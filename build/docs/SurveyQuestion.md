# SurveyQuestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Survey** | Pointer to [**SurveyReference**](SurveyReference.md) |  | [optional] 
**FieldType** | **NullableString** |  | 
**EntryType** | **NullableString** |  | 
**SequenceNumber** | **NullableFloat64** |  | 
**Question** | **string** |  Max length: 1000; | 
**NumberOfDecimals** | Pointer to **NullableInt32** |  | [optional] 
**RequiredFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSurveyQuestion

`func NewSurveyQuestion(fieldType NullableString, entryType NullableString, sequenceNumber NullableFloat64, question string, ) *SurveyQuestion`

NewSurveyQuestion instantiates a new SurveyQuestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSurveyQuestionWithDefaults

`func NewSurveyQuestionWithDefaults() *SurveyQuestion`

NewSurveyQuestionWithDefaults instantiates a new SurveyQuestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SurveyQuestion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SurveyQuestion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SurveyQuestion) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SurveyQuestion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSurvey

`func (o *SurveyQuestion) GetSurvey() SurveyReference`

GetSurvey returns the Survey field if non-nil, zero value otherwise.

### GetSurveyOk

`func (o *SurveyQuestion) GetSurveyOk() (*SurveyReference, bool)`

GetSurveyOk returns a tuple with the Survey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurvey

`func (o *SurveyQuestion) SetSurvey(v SurveyReference)`

SetSurvey sets Survey field to given value.

### HasSurvey

`func (o *SurveyQuestion) HasSurvey() bool`

HasSurvey returns a boolean if a field has been set.

### GetFieldType

`func (o *SurveyQuestion) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *SurveyQuestion) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *SurveyQuestion) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.


### SetFieldTypeNil

`func (o *SurveyQuestion) SetFieldTypeNil(b bool)`

 SetFieldTypeNil sets the value for FieldType to be an explicit nil

### UnsetFieldType
`func (o *SurveyQuestion) UnsetFieldType()`

UnsetFieldType ensures that no value is present for FieldType, not even an explicit nil
### GetEntryType

`func (o *SurveyQuestion) GetEntryType() string`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *SurveyQuestion) GetEntryTypeOk() (*string, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *SurveyQuestion) SetEntryType(v string)`

SetEntryType sets EntryType field to given value.


### SetEntryTypeNil

`func (o *SurveyQuestion) SetEntryTypeNil(b bool)`

 SetEntryTypeNil sets the value for EntryType to be an explicit nil

### UnsetEntryType
`func (o *SurveyQuestion) UnsetEntryType()`

UnsetEntryType ensures that no value is present for EntryType, not even an explicit nil
### GetSequenceNumber

`func (o *SurveyQuestion) GetSequenceNumber() float64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *SurveyQuestion) GetSequenceNumberOk() (*float64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *SurveyQuestion) SetSequenceNumber(v float64)`

SetSequenceNumber sets SequenceNumber field to given value.


### SetSequenceNumberNil

`func (o *SurveyQuestion) SetSequenceNumberNil(b bool)`

 SetSequenceNumberNil sets the value for SequenceNumber to be an explicit nil

### UnsetSequenceNumber
`func (o *SurveyQuestion) UnsetSequenceNumber()`

UnsetSequenceNumber ensures that no value is present for SequenceNumber, not even an explicit nil
### GetQuestion

`func (o *SurveyQuestion) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *SurveyQuestion) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *SurveyQuestion) SetQuestion(v string)`

SetQuestion sets Question field to given value.


### GetNumberOfDecimals

`func (o *SurveyQuestion) GetNumberOfDecimals() int32`

GetNumberOfDecimals returns the NumberOfDecimals field if non-nil, zero value otherwise.

### GetNumberOfDecimalsOk

`func (o *SurveyQuestion) GetNumberOfDecimalsOk() (*int32, bool)`

GetNumberOfDecimalsOk returns a tuple with the NumberOfDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDecimals

`func (o *SurveyQuestion) SetNumberOfDecimals(v int32)`

SetNumberOfDecimals sets NumberOfDecimals field to given value.

### HasNumberOfDecimals

`func (o *SurveyQuestion) HasNumberOfDecimals() bool`

HasNumberOfDecimals returns a boolean if a field has been set.

### SetNumberOfDecimalsNil

`func (o *SurveyQuestion) SetNumberOfDecimalsNil(b bool)`

 SetNumberOfDecimalsNil sets the value for NumberOfDecimals to be an explicit nil

### UnsetNumberOfDecimals
`func (o *SurveyQuestion) UnsetNumberOfDecimals()`

UnsetNumberOfDecimals ensures that no value is present for NumberOfDecimals, not even an explicit nil
### GetRequiredFlag

`func (o *SurveyQuestion) GetRequiredFlag() bool`

GetRequiredFlag returns the RequiredFlag field if non-nil, zero value otherwise.

### GetRequiredFlagOk

`func (o *SurveyQuestion) GetRequiredFlagOk() (*bool, bool)`

GetRequiredFlagOk returns a tuple with the RequiredFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFlag

`func (o *SurveyQuestion) SetRequiredFlag(v bool)`

SetRequiredFlag sets RequiredFlag field to given value.

### HasRequiredFlag

`func (o *SurveyQuestion) HasRequiredFlag() bool`

HasRequiredFlag returns a boolean if a field has been set.

### SetRequiredFlagNil

`func (o *SurveyQuestion) SetRequiredFlagNil(b bool)`

 SetRequiredFlagNil sets the value for RequiredFlag to be an explicit nil

### UnsetRequiredFlag
`func (o *SurveyQuestion) UnsetRequiredFlag()`

UnsetRequiredFlag ensures that no value is present for RequiredFlag, not even an explicit nil
### GetInactiveFlag

`func (o *SurveyQuestion) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *SurveyQuestion) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *SurveyQuestion) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *SurveyQuestion) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *SurveyQuestion) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *SurveyQuestion) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetInfo

`func (o *SurveyQuestion) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SurveyQuestion) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SurveyQuestion) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SurveyQuestion) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


