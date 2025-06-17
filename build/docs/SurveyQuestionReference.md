# SurveyQuestionReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Question** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSurveyQuestionReference

`func NewSurveyQuestionReference() *SurveyQuestionReference`

NewSurveyQuestionReference instantiates a new SurveyQuestionReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSurveyQuestionReferenceWithDefaults

`func NewSurveyQuestionReferenceWithDefaults() *SurveyQuestionReference`

NewSurveyQuestionReferenceWithDefaults instantiates a new SurveyQuestionReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SurveyQuestionReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SurveyQuestionReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SurveyQuestionReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SurveyQuestionReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SurveyQuestionReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SurveyQuestionReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetQuestion

`func (o *SurveyQuestionReference) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *SurveyQuestionReference) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *SurveyQuestionReference) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *SurveyQuestionReference) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetInfo

`func (o *SurveyQuestionReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SurveyQuestionReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SurveyQuestionReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SurveyQuestionReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


