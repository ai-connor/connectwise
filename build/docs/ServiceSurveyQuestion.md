# ServiceSurveyQuestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**SequenceNumber** | Pointer to **NullableInt32** |  | [optional] 
**Type** | **NullableString** |  | 
**Question** | **string** |  Max length: 1000; | 
**Options** | Pointer to [**[]ServiceSurveyQuestionOption**](ServiceSurveyQuestionOption.md) |  | [optional] 
**IncludeFlag** | Pointer to **NullableBool** |  | [optional] 
**RequiredFlag** | Pointer to **NullableBool** |  | [optional] 
**NoAnswerPoints** | Pointer to **NullableInt32** |  | [optional] 
**SurveyId** | Pointer to **NullableInt32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewServiceSurveyQuestion

`func NewServiceSurveyQuestion(type_ NullableString, question string, ) *ServiceSurveyQuestion`

NewServiceSurveyQuestion instantiates a new ServiceSurveyQuestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceSurveyQuestionWithDefaults

`func NewServiceSurveyQuestionWithDefaults() *ServiceSurveyQuestion`

NewServiceSurveyQuestionWithDefaults instantiates a new ServiceSurveyQuestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceSurveyQuestion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceSurveyQuestion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceSurveyQuestion) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceSurveyQuestion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *ServiceSurveyQuestion) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *ServiceSurveyQuestion) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *ServiceSurveyQuestion) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *ServiceSurveyQuestion) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### SetSequenceNumberNil

`func (o *ServiceSurveyQuestion) SetSequenceNumberNil(b bool)`

 SetSequenceNumberNil sets the value for SequenceNumber to be an explicit nil

### UnsetSequenceNumber
`func (o *ServiceSurveyQuestion) UnsetSequenceNumber()`

UnsetSequenceNumber ensures that no value is present for SequenceNumber, not even an explicit nil
### GetType

`func (o *ServiceSurveyQuestion) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceSurveyQuestion) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceSurveyQuestion) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ServiceSurveyQuestion) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ServiceSurveyQuestion) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetQuestion

`func (o *ServiceSurveyQuestion) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *ServiceSurveyQuestion) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *ServiceSurveyQuestion) SetQuestion(v string)`

SetQuestion sets Question field to given value.


### GetOptions

`func (o *ServiceSurveyQuestion) GetOptions() []ServiceSurveyQuestionOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ServiceSurveyQuestion) GetOptionsOk() (*[]ServiceSurveyQuestionOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ServiceSurveyQuestion) SetOptions(v []ServiceSurveyQuestionOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ServiceSurveyQuestion) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetIncludeFlag

`func (o *ServiceSurveyQuestion) GetIncludeFlag() bool`

GetIncludeFlag returns the IncludeFlag field if non-nil, zero value otherwise.

### GetIncludeFlagOk

`func (o *ServiceSurveyQuestion) GetIncludeFlagOk() (*bool, bool)`

GetIncludeFlagOk returns a tuple with the IncludeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFlag

`func (o *ServiceSurveyQuestion) SetIncludeFlag(v bool)`

SetIncludeFlag sets IncludeFlag field to given value.

### HasIncludeFlag

`func (o *ServiceSurveyQuestion) HasIncludeFlag() bool`

HasIncludeFlag returns a boolean if a field has been set.

### SetIncludeFlagNil

`func (o *ServiceSurveyQuestion) SetIncludeFlagNil(b bool)`

 SetIncludeFlagNil sets the value for IncludeFlag to be an explicit nil

### UnsetIncludeFlag
`func (o *ServiceSurveyQuestion) UnsetIncludeFlag()`

UnsetIncludeFlag ensures that no value is present for IncludeFlag, not even an explicit nil
### GetRequiredFlag

`func (o *ServiceSurveyQuestion) GetRequiredFlag() bool`

GetRequiredFlag returns the RequiredFlag field if non-nil, zero value otherwise.

### GetRequiredFlagOk

`func (o *ServiceSurveyQuestion) GetRequiredFlagOk() (*bool, bool)`

GetRequiredFlagOk returns a tuple with the RequiredFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFlag

`func (o *ServiceSurveyQuestion) SetRequiredFlag(v bool)`

SetRequiredFlag sets RequiredFlag field to given value.

### HasRequiredFlag

`func (o *ServiceSurveyQuestion) HasRequiredFlag() bool`

HasRequiredFlag returns a boolean if a field has been set.

### SetRequiredFlagNil

`func (o *ServiceSurveyQuestion) SetRequiredFlagNil(b bool)`

 SetRequiredFlagNil sets the value for RequiredFlag to be an explicit nil

### UnsetRequiredFlag
`func (o *ServiceSurveyQuestion) UnsetRequiredFlag()`

UnsetRequiredFlag ensures that no value is present for RequiredFlag, not even an explicit nil
### GetNoAnswerPoints

`func (o *ServiceSurveyQuestion) GetNoAnswerPoints() int32`

GetNoAnswerPoints returns the NoAnswerPoints field if non-nil, zero value otherwise.

### GetNoAnswerPointsOk

`func (o *ServiceSurveyQuestion) GetNoAnswerPointsOk() (*int32, bool)`

GetNoAnswerPointsOk returns a tuple with the NoAnswerPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAnswerPoints

`func (o *ServiceSurveyQuestion) SetNoAnswerPoints(v int32)`

SetNoAnswerPoints sets NoAnswerPoints field to given value.

### HasNoAnswerPoints

`func (o *ServiceSurveyQuestion) HasNoAnswerPoints() bool`

HasNoAnswerPoints returns a boolean if a field has been set.

### SetNoAnswerPointsNil

`func (o *ServiceSurveyQuestion) SetNoAnswerPointsNil(b bool)`

 SetNoAnswerPointsNil sets the value for NoAnswerPoints to be an explicit nil

### UnsetNoAnswerPoints
`func (o *ServiceSurveyQuestion) UnsetNoAnswerPoints()`

UnsetNoAnswerPoints ensures that no value is present for NoAnswerPoints, not even an explicit nil
### GetSurveyId

`func (o *ServiceSurveyQuestion) GetSurveyId() int32`

GetSurveyId returns the SurveyId field if non-nil, zero value otherwise.

### GetSurveyIdOk

`func (o *ServiceSurveyQuestion) GetSurveyIdOk() (*int32, bool)`

GetSurveyIdOk returns a tuple with the SurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyId

`func (o *ServiceSurveyQuestion) SetSurveyId(v int32)`

SetSurveyId sets SurveyId field to given value.

### HasSurveyId

`func (o *ServiceSurveyQuestion) HasSurveyId() bool`

HasSurveyId returns a boolean if a field has been set.

### SetSurveyIdNil

`func (o *ServiceSurveyQuestion) SetSurveyIdNil(b bool)`

 SetSurveyIdNil sets the value for SurveyId to be an explicit nil

### UnsetSurveyId
`func (o *ServiceSurveyQuestion) UnsetSurveyId()`

UnsetSurveyId ensures that no value is present for SurveyId, not even an explicit nil
### GetInfo

`func (o *ServiceSurveyQuestion) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ServiceSurveyQuestion) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ServiceSurveyQuestion) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ServiceSurveyQuestion) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


