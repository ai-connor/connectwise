# SurveyQuestionValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Survey** | Pointer to [**SurveyReference**](SurveyReference.md) |  | [optional] 
**Question** | Pointer to [**SurveyQuestionReference**](SurveyQuestionReference.md) |  | [optional] 
**Value** | **string** |  Max length: 1000; | 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**PointValue** | Pointer to **NullableInt32** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSurveyQuestionValue

`func NewSurveyQuestionValue(value string, ) *SurveyQuestionValue`

NewSurveyQuestionValue instantiates a new SurveyQuestionValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSurveyQuestionValueWithDefaults

`func NewSurveyQuestionValueWithDefaults() *SurveyQuestionValue`

NewSurveyQuestionValueWithDefaults instantiates a new SurveyQuestionValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SurveyQuestionValue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SurveyQuestionValue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SurveyQuestionValue) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SurveyQuestionValue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSurvey

`func (o *SurveyQuestionValue) GetSurvey() SurveyReference`

GetSurvey returns the Survey field if non-nil, zero value otherwise.

### GetSurveyOk

`func (o *SurveyQuestionValue) GetSurveyOk() (*SurveyReference, bool)`

GetSurveyOk returns a tuple with the Survey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurvey

`func (o *SurveyQuestionValue) SetSurvey(v SurveyReference)`

SetSurvey sets Survey field to given value.

### HasSurvey

`func (o *SurveyQuestionValue) HasSurvey() bool`

HasSurvey returns a boolean if a field has been set.

### GetQuestion

`func (o *SurveyQuestionValue) GetQuestion() SurveyQuestionReference`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *SurveyQuestionValue) GetQuestionOk() (*SurveyQuestionReference, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *SurveyQuestionValue) SetQuestion(v SurveyQuestionReference)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *SurveyQuestionValue) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetValue

`func (o *SurveyQuestionValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SurveyQuestionValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SurveyQuestionValue) SetValue(v string)`

SetValue sets Value field to given value.


### GetDefaultFlag

`func (o *SurveyQuestionValue) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *SurveyQuestionValue) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *SurveyQuestionValue) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *SurveyQuestionValue) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *SurveyQuestionValue) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *SurveyQuestionValue) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetPointValue

`func (o *SurveyQuestionValue) GetPointValue() int32`

GetPointValue returns the PointValue field if non-nil, zero value otherwise.

### GetPointValueOk

`func (o *SurveyQuestionValue) GetPointValueOk() (*int32, bool)`

GetPointValueOk returns a tuple with the PointValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointValue

`func (o *SurveyQuestionValue) SetPointValue(v int32)`

SetPointValue sets PointValue field to given value.

### HasPointValue

`func (o *SurveyQuestionValue) HasPointValue() bool`

HasPointValue returns a boolean if a field has been set.

### SetPointValueNil

`func (o *SurveyQuestionValue) SetPointValueNil(b bool)`

 SetPointValueNil sets the value for PointValue to be an explicit nil

### UnsetPointValue
`func (o *SurveyQuestionValue) UnsetPointValue()`

UnsetPointValue ensures that no value is present for PointValue, not even an explicit nil
### GetInactiveFlag

`func (o *SurveyQuestionValue) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *SurveyQuestionValue) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *SurveyQuestionValue) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *SurveyQuestionValue) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *SurveyQuestionValue) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *SurveyQuestionValue) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetInfo

`func (o *SurveyQuestionValue) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SurveyQuestionValue) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SurveyQuestionValue) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SurveyQuestionValue) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


