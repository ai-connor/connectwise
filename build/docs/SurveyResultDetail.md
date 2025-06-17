# SurveyResultDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuestionId** | Pointer to **NullableInt32** |  | [optional] 
**Answer** | Pointer to **map[string]interface{}** | If question type is Selection, this should be the option array index. | [optional] 

## Methods

### NewSurveyResultDetail

`func NewSurveyResultDetail() *SurveyResultDetail`

NewSurveyResultDetail instantiates a new SurveyResultDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSurveyResultDetailWithDefaults

`func NewSurveyResultDetailWithDefaults() *SurveyResultDetail`

NewSurveyResultDetailWithDefaults instantiates a new SurveyResultDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestionId

`func (o *SurveyResultDetail) GetQuestionId() int32`

GetQuestionId returns the QuestionId field if non-nil, zero value otherwise.

### GetQuestionIdOk

`func (o *SurveyResultDetail) GetQuestionIdOk() (*int32, bool)`

GetQuestionIdOk returns a tuple with the QuestionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionId

`func (o *SurveyResultDetail) SetQuestionId(v int32)`

SetQuestionId sets QuestionId field to given value.

### HasQuestionId

`func (o *SurveyResultDetail) HasQuestionId() bool`

HasQuestionId returns a boolean if a field has been set.

### SetQuestionIdNil

`func (o *SurveyResultDetail) SetQuestionIdNil(b bool)`

 SetQuestionIdNil sets the value for QuestionId to be an explicit nil

### UnsetQuestionId
`func (o *SurveyResultDetail) UnsetQuestionId()`

UnsetQuestionId ensures that no value is present for QuestionId, not even an explicit nil
### GetAnswer

`func (o *SurveyResultDetail) GetAnswer() map[string]interface{}`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *SurveyResultDetail) GetAnswerOk() (*map[string]interface{}, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *SurveyResultDetail) SetAnswer(v map[string]interface{})`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *SurveyResultDetail) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


